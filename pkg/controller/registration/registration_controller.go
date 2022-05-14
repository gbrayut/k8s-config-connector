// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registration

import (
	"context"
	"fmt"
	"sync"
	"time"

	dclcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/deletiondefender"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/gsakeysecretgenerator"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/auditconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/partialpolicy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/policy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/policymember"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/tf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	klog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const controllerName = "registration-controller"
const serviceAccountKeyAPIGroup = "iam.cnrm.cloud.google.com"
const serviceAccountKeyKind = "IAMServiceAccountKey"

var logger = klog.Log.WithName(controllerName)

// Add creates a new registration Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager, p *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader, dclConfig *dcl.Config, dclConverter *conversion.Converter, regFunc registrationFunc) error {
	r := &ReconcileRegistration{
		Client:           mgr.GetClient(),
		provider:         p,
		smLoader:         smLoader,
		dclConfig:        dclConfig,
		dclConverter:     dclConverter,
		mgr:              mgr,
		controllers:      make(map[string]map[string]bool),
		registrationFunc: regFunc,
	}
	c, err := controller.New(controllerName, mgr,
		controller.Options{
			Reconciler:              r,
			MaxConcurrentReconciles: k8s.ControllerMaxConcurrentReconciles,
		})
	if err != nil {
		return err
	}
	return c.Watch(&source.Kind{Type: &apiextensions.CustomResourceDefinition{}}, &handler.EnqueueRequestForObject{}, ManagedByKCCPredicate{})
}

var _ reconcile.Reconciler = &ReconcileRegistration{}

// ReconcileRegistration reconciles a CRD owned by KCC
type ReconcileRegistration struct {
	client.Client
	provider         *tfschema.Provider
	smLoader         *servicemappingloader.ServiceMappingLoader
	dclConfig        *dcl.Config
	dclConverter     *conversion.Converter
	mgr              manager.Manager
	controllers      map[string]map[string]bool
	registrationFunc registrationFunc
	mu               sync.Mutex
}

// RegistrationFunc is the function that handles the registration of a controller for the given CRD
type registrationFunc func(*ReconcileRegistration, *apiextensions.CustomResourceDefinition, schema.GroupVersionKind) error

func (r *ReconcileRegistration) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	// Fetch the TypeProvider tp
	crd := &apiextensions.CustomResourceDefinition{}
	err := r.Get(ctx, request.NamespacedName, crd)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	logger.Info("Waiting to obtain lock...", "kind", crd.Spec.Names.Kind)
	start := time.Now()
	r.mu.Lock()
	logger.Info("Obtained lock", "kind", crd.Spec.Names.Kind, "elapsed (μs)", time.Since(start).Microseconds())
	defer func() {
		logger.Info("Releasing lock...", "kind", crd.Spec.Names.Kind)
		r.mu.Unlock()
	}()
	gvk := schema.GroupVersionKind{
		Group:   crd.Spec.Group,
		Version: k8s.GetVersionFromCRD(crd),
		Kind:    crd.Spec.Names.Kind,
	}
	if kindMapForGroup, exists := r.controllers[gvk.Group]; exists {
		if kindMapForGroup[gvk.Kind] {
			logger.Info("controller already registered for kind in API group", "group", gvk.Group, "version", gvk.Version, "kind", gvk.Kind)
			return reconcile.Result{}, nil
		}
	} else {
		r.controllers[gvk.Group] = make(map[string]bool)
	}

	if err := r.registrationFunc(r, crd, gvk); err != nil {
		return reconcile.Result{}, fmt.Errorf("error registering controller: %w", err)
	}

	r.controllers[gvk.Group][gvk.Kind] = true
	return reconcile.Result{}, nil
}

func isServiceAccountKeyCRD(crd *apiextensions.CustomResourceDefinition) bool {
	return crd.Spec.Group == serviceAccountKeyAPIGroup && crd.Spec.Names.Kind == serviceAccountKeyKind
}

func RegisterDefaultController(r *ReconcileRegistration, crd *apiextensions.CustomResourceDefinition, gvk schema.GroupVersionKind) error {
	// Depending on which resource it is, we need to register a different controller.
	switch gvk.Kind {
	case "IAMPolicy":
		if err := policy.Add(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig); err != nil {
			return err
		}
	case "IAMPartialPolicy":
		if err := partialpolicy.Add(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig); err != nil {
			return err
		}
	case "IAMPolicyMember":
		if err := policymember.Add(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig); err != nil {
			return err
		}
	case "IAMAuditConfig":
		if err := auditconfig.Add(r.mgr, r.provider, r.smLoader, r.dclConverter, r.dclConfig); err != nil {
			return err
		}
	default:
		// register controllers for dcl-based CRDs
		if val, ok := crd.Labels[k8s.DCL2CRDLabel]; ok && val == "true" {
			if err := dclcontroller.Add(r.mgr, crd, r.dclConverter, r.dclConfig, r.smLoader); err != nil {
				return fmt.Errorf("error adding dcl controller for %v to a manager: %v", crd.Spec.Names.Kind, err)
			}
			return nil
		}
		// register controllers for tf-based CRDs
		if val, ok := crd.Labels[crdgeneration.TF2CRDLabel]; !ok || val != "true" {
			logger.Info("unrecognized CRD; skipping controller registration", "group", gvk.Group, "version", gvk.Version, "kind", gvk.Kind)
			return nil
		}
		if err := tf.Add(r.mgr, crd, r.provider, r.smLoader); err != nil {
			return fmt.Errorf("error adding terraform controller for %v to a manager: %v", crd.Spec.Names.Kind, err)
		}
		// register the controller to automatically create secrets for GSA keys
		if isServiceAccountKeyCRD(crd) {
			logger.Info("registering the GSA-Key-to-Secret generation controller")
			if err := gsakeysecretgenerator.Add(r.mgr, crd); err != nil {
				return fmt.Errorf("error adding the gsa-to-secret generator for %v to a manager: %v", crd.Spec.Names.Kind, err)
			}
		}
	}
	return nil
}

func RegisterDeletionDefenderController(r *ReconcileRegistration, crd *apiextensions.CustomResourceDefinition, gvk schema.GroupVersionKind) error {
	if crd.Spec.Names.Kind == "ServiceMapping" {
		// ServiceMapping is a special resource type that does not make a call to an underlying GCP API
		return nil
	}
	if err := deletiondefender.Add(r.mgr, crd); err != nil {
		return fmt.Errorf("error registering deletion defender controller for '%v': %w", crd.GetName(), err)
	}
	return nil
}