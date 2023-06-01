// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package appengine

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceAppEngineApplicationUrlDispatchRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppEngineApplicationUrlDispatchRulesCreate,
		Read:   resourceAppEngineApplicationUrlDispatchRulesRead,
		Update: resourceAppEngineApplicationUrlDispatchRulesUpdate,
		Delete: resourceAppEngineApplicationUrlDispatchRulesDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAppEngineApplicationUrlDispatchRulesImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dispatch_rules": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Rules to match an HTTP request and dispatch that request to a service.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
The sum of the lengths of the domain and path may not exceed 100 characters.`,
						},
						"service": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
The sum of the lengths of the domain and path may not exceed 100 characters.`,
						},
						"domain": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Domain name to match against. The wildcard "*" is supported if specified before a period: "*.".
Defaults to matching all domains: "*".`,
							Default: "*",
						},
					},
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAppEngineApplicationUrlDispatchRulesCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	dispatchRulesProp, err := expandAppEngineApplicationUrlDispatchRulesDispatchRules(d.Get("dispatch_rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dispatch_rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(dispatchRulesProp)) && (ok || !reflect.DeepEqual(v, dispatchRulesProp)) {
		obj["dispatchRules"] = dispatchRulesProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}?updateMask=dispatch_rules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ApplicationUrlDispatchRules: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ApplicationUrlDispatchRules: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "PATCH",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutCreate),
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsAppEngineRetryableError},
	})
	if err != nil {
		return fmt.Errorf("Error creating ApplicationUrlDispatchRules: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = AppEngineOperationWaitTime(
		config, res, project, "Creating ApplicationUrlDispatchRules", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ApplicationUrlDispatchRules: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ApplicationUrlDispatchRules %q: %#v", d.Id(), res)

	return resourceAppEngineApplicationUrlDispatchRulesRead(d, meta)
}

func resourceAppEngineApplicationUrlDispatchRulesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ApplicationUrlDispatchRules: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "GET",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsAppEngineRetryableError},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AppEngineApplicationUrlDispatchRules %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ApplicationUrlDispatchRules: %s", err)
	}

	if err := d.Set("dispatch_rules", flattenAppEngineApplicationUrlDispatchRulesDispatchRules(res["dispatchRules"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApplicationUrlDispatchRules: %s", err)
	}

	return nil
}

func resourceAppEngineApplicationUrlDispatchRulesUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ApplicationUrlDispatchRules: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	dispatchRulesProp, err := expandAppEngineApplicationUrlDispatchRulesDispatchRules(d.Get("dispatch_rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dispatch_rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, dispatchRulesProp)) {
		obj["dispatchRules"] = dispatchRulesProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}?updateMask=dispatch_rules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ApplicationUrlDispatchRules %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "PATCH",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutUpdate),
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsAppEngineRetryableError},
	})

	if err != nil {
		return fmt.Errorf("Error updating ApplicationUrlDispatchRules %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ApplicationUrlDispatchRules %q: %#v", d.Id(), res)
	}

	err = AppEngineOperationWaitTime(
		config, res, project, "Updating ApplicationUrlDispatchRules", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceAppEngineApplicationUrlDispatchRulesRead(d, meta)
}

func resourceAppEngineApplicationUrlDispatchRulesDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ApplicationUrlDispatchRules: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}?updateMask=dispatch_rules")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ApplicationUrlDispatchRules %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "PATCH",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutDelete),
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsAppEngineRetryableError},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ApplicationUrlDispatchRules")
	}

	err = AppEngineOperationWaitTime(
		config, res, project, "Deleting ApplicationUrlDispatchRules", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ApplicationUrlDispatchRules %q: %#v", d.Id(), res)
	return nil
}

func resourceAppEngineApplicationUrlDispatchRulesImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"(?P<project>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAppEngineApplicationUrlDispatchRulesDispatchRules(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"domain":  flattenAppEngineApplicationUrlDispatchRulesDispatchRulesDomain(original["domain"], d, config),
			"path":    flattenAppEngineApplicationUrlDispatchRulesDispatchRulesPath(original["path"], d, config),
			"service": flattenAppEngineApplicationUrlDispatchRulesDispatchRulesService(original["service"], d, config),
		})
	}
	return transformed
}
func flattenAppEngineApplicationUrlDispatchRulesDispatchRulesDomain(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAppEngineApplicationUrlDispatchRulesDispatchRulesPath(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAppEngineApplicationUrlDispatchRulesDispatchRulesService(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDomain, err := expandAppEngineApplicationUrlDispatchRulesDispatchRulesDomain(original["domain"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDomain); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["domain"] = transformedDomain
		}

		transformedPath, err := expandAppEngineApplicationUrlDispatchRulesDispatchRulesPath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		transformedService, err := expandAppEngineApplicationUrlDispatchRulesDispatchRulesService(original["service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["service"] = transformedService
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRulesDomain(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRulesPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRulesService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
