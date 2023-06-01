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

package accesscontextmanager

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceAccessContextManagerServicePerimeterEgressPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessContextManagerServicePerimeterEgressPolicyCreate,
		Read:   resourceAccessContextManagerServicePerimeterEgressPolicyRead,
		Update: resourceAccessContextManagerServicePerimeterEgressPolicyUpdate,
		Delete: resourceAccessContextManagerServicePerimeterEgressPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessContextManagerServicePerimeterEgressPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"perimeter": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The name of the Service Perimeter to add this resource to.`,
			},
			"egress_from": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Defines conditions on the source of a request causing this 'EgressPolicy' to apply.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identities": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `A list of identities that are allowed access through this 'EgressPolicy'.
Should be in the format of email address. The email address should
represent individual user or service account only.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"identity_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT", ""}),
							Description: `Specifies the type of identities that are allowed access to outside the
perimeter. If left unspecified, then members of 'identities' field will
be allowed access. Possible values: ["ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]`,
						},
					},
				},
			},
			"egress_to": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Defines the conditions on the 'ApiOperation' and destination resources that
cause this 'EgressPolicy' to apply.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external_resources": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `A list of external resources that are allowed to be accessed. A request
matches if it contains an external resource in this list (Example:
s3://bucket/path). Currently '*' is not allowed.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"operations": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `A list of 'ApiOperations' that this egress rule applies to. A request matches
if it contains an operation/service in this list.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"method_selectors": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `API methods or permissions to allow. Method or permission must belong
to the service specified by 'serviceName' field. A single MethodSelector
entry with '*' specified for the 'method' field will allow all methods
AND permissions for the service specified in 'serviceName'.`,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"method": {
													Type:     schema.TypeString,
													Optional: true,
													Description: `Value for 'method' should be a valid method name for the corresponding
'serviceName' in 'ApiOperation'. If '*' used as value for method,
then ALL methods and permissions are allowed.`,
												},
												"permission": {
													Type:     schema.TypeString,
													Optional: true,
													Description: `Value for permission should be a valid Cloud IAM permission for the
corresponding 'serviceName' in 'ApiOperation'.`,
												},
											},
										},
									},
									"service_name": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `The name of the API whose methods or permissions the 'IngressPolicy' or
'EgressPolicy' want to allow. A single 'ApiOperation' with serviceName
field set to '*' will allow all methods AND permissions for all services.`,
									},
								},
							},
						},
						"resources": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `A list of resources, currently only projects in the form
'projects/<projectnumber>', that match this to stanza. A request matches
if it contains a resource in this list. If * is specified for resources,
then this 'EgressTo' rule will authorize access to all resources outside
the perimeter.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAccessContextManagerServicePerimeterEgressPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	egressFromProp, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressFrom(d.Get("egress_from"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("egress_from"); !tpgresource.IsEmptyValue(reflect.ValueOf(egressFromProp)) && (ok || !reflect.DeepEqual(v, egressFromProp)) {
		obj["egressFrom"] = egressFromProp
	}
	egressToProp, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressTo(d.Get("egress_to"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("egress_to"); !tpgresource.IsEmptyValue(reflect.ValueOf(egressToProp)) && (ok || !reflect.DeepEqual(v, egressToProp)) {
		obj["egressTo"] = egressToProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "{{perimeter}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{perimeter}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServicePerimeterEgressPolicy: %#v", obj)

	obj, err = resourceAccessContextManagerServicePerimeterEgressPolicyPatchCreateEncoder(d, meta, obj)
	if err != nil {
		return err
	}
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": "status.egressPolicies"})
	if err != nil {
		return err
	}
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating ServicePerimeterEgressPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{perimeter}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = AccessContextManagerOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating ServicePerimeterEgressPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create ServicePerimeterEgressPolicy: %s", err)
	}

	if _, ok := opRes["status"]; ok {
		opRes, err = flattenNestedAccessContextManagerServicePerimeterEgressPolicy(d, meta, opRes)
		if err != nil {
			return fmt.Errorf("Error getting nested object from operation response: %s", err)
		}
		if opRes == nil {
			// Object isn't there any more - remove it from the state.
			return fmt.Errorf("Error decoding response from operation, could not find nested object")
		}
	}
	if err := d.Set("egress_from", flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFrom(opRes["egressFrom"], d, config)); err != nil {
		return err
	}
	if err := d.Set("egress_to", flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressTo(opRes["egressTo"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "{{perimeter}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ServicePerimeterEgressPolicy %q: %#v", d.Id(), res)

	return resourceAccessContextManagerServicePerimeterEgressPolicyRead(d, meta)
}

func resourceAccessContextManagerServicePerimeterEgressPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{perimeter}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AccessContextManagerServicePerimeterEgressPolicy %q", d.Id()))
	}

	res, err = flattenNestedAccessContextManagerServicePerimeterEgressPolicy(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing AccessContextManagerServicePerimeterEgressPolicy because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("egress_from", flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFrom(res["egressFrom"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeterEgressPolicy: %s", err)
	}
	if err := d.Set("egress_to", flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressTo(res["egressTo"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeterEgressPolicy: %s", err)
	}

	return nil
}

func resourceAccessContextManagerServicePerimeterEgressPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	egressFromProp, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressFrom(d.Get("egress_from"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("egress_from"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, egressFromProp)) {
		obj["egressFrom"] = egressFromProp
	}
	egressToProp, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressTo(d.Get("egress_to"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("egress_to"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, egressToProp)) {
		obj["egressTo"] = egressToProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "{{perimeter}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{perimeter}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ServicePerimeterEgressPolicy %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("egress_from") {
		updateMask = append(updateMask, "egressFrom")
	}

	if d.HasChange("egress_to") {
		updateMask = append(updateMask, "egressTo")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	obj, err = resourceAccessContextManagerServicePerimeterEgressPolicyPatchUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating ServicePerimeterEgressPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ServicePerimeterEgressPolicy %q: %#v", d.Id(), res)
	}

	err = AccessContextManagerOperationWaitTime(
		config, res, "Updating ServicePerimeterEgressPolicy", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceAccessContextManagerServicePerimeterEgressPolicyRead(d, meta)
}

func resourceAccessContextManagerServicePerimeterEgressPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	lockName, err := tpgresource.ReplaceVars(d, config, "{{perimeter}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{perimeter}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	obj, err = resourceAccessContextManagerServicePerimeterEgressPolicyPatchDeleteEncoder(d, meta, obj)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ServicePerimeterEgressPolicy")
	}
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": "status.egressPolicies"})
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Deleting ServicePerimeterEgressPolicy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ServicePerimeterEgressPolicy")
	}

	err = AccessContextManagerOperationWaitTime(
		config, res, "Deleting ServicePerimeterEgressPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ServicePerimeterEgressPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceAccessContextManagerServicePerimeterEgressPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	parts, err := tpgresource.GetImportIdQualifiers([]string{"accessPolicies/(?P<accessPolicy>[^/]+)/servicePerimeters/(?P<perimeter>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return nil, err
	}

	if err := d.Set("perimeter", fmt.Sprintf("accessPolicies/%s/servicePerimeters/%s", parts["accessPolicy"], parts["perimeter"])); err != nil {
		return nil, fmt.Errorf("Error setting perimeter: %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFrom(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["identity_type"] =
		flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromIdentityType(original["identityType"], d, config)
	transformed["identities"] =
		flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromIdentities(original["identities"], d, config)
	return []interface{}{transformed}
}
func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromIdentityType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromIdentities(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressTo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["resources"] =
		flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToResources(original["resources"], d, config)
	transformed["external_resources"] =
		flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToExternalResources(original["externalResources"], d, config)
	transformed["operations"] =
		flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperations(original["operations"], d, config)
	return []interface{}{transformed}
}
func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToResources(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToExternalResources(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"service_name":     flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsServiceName(original["serviceName"], d, config),
			"method_selectors": flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectors(original["methodSelectors"], d, config),
		})
	}
	return transformed
}
func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsServiceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectors(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"method":     flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsMethod(original["method"], d, config),
			"permission": flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsPermission(original["permission"], d, config),
		})
	}
	return transformed
}
func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsMethod(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsPermission(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressFrom(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIdentityType, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromIdentityType(original["identity_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdentityType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["identityType"] = transformedIdentityType
	}

	transformedIdentities, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromIdentities(original["identities"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdentities); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["identities"] = transformedIdentities
	}

	return transformed, nil
}

func expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromIdentityType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressFromIdentities(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressTo(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResources, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressToResources(original["resources"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResources); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resources"] = transformedResources
	}

	transformedExternalResources, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressToExternalResources(original["external_resources"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExternalResources); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["externalResources"] = transformedExternalResources
	}

	transformedOperations, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperations(original["operations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOperations); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["operations"] = transformedOperations
	}

	return transformed, nil
}

func expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressToResources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressToExternalResources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedServiceName, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsServiceName(original["service_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedServiceName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["serviceName"] = transformedServiceName
		}

		transformedMethodSelectors, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectors(original["method_selectors"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMethodSelectors); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["methodSelectors"] = transformedMethodSelectors
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsServiceName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectors(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMethod, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsMethod(original["method"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMethod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["method"] = transformedMethod
		}

		transformedPermission, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsPermission(original["permission"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPermission); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["permission"] = transformedPermission
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsMethod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsPermission(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func flattenNestedAccessContextManagerServicePerimeterEgressPolicy(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["status"]
	if !ok || v == nil {
		return nil, nil
	}
	res = v.(map[string]interface{})

	v, ok = res["egressPolicies"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value status.egressPolicies. Actual value: %v", v)
	}

	_, item, err := resourceAccessContextManagerServicePerimeterEgressPolicyFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceAccessContextManagerServicePerimeterEgressPolicyFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedEgressFrom, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressFrom(d.Get("egress_from"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedEgressFrom := flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFrom(expectedEgressFrom, d, meta.(*transport_tpg.Config))
	expectedEgressTo, err := expandNestedAccessContextManagerServicePerimeterEgressPolicyEgressTo(d.Get("egress_to"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedEgressTo := flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressTo(expectedEgressTo, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemEgressFrom := flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressFrom(item["egressFrom"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemEgressFrom)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedEgressFrom))) && !reflect.DeepEqual(itemEgressFrom, expectedFlattenedEgressFrom) {
			log.Printf("[DEBUG] Skipping item with egressFrom= %#v, looking for %#v)", itemEgressFrom, expectedFlattenedEgressFrom)
			continue
		}
		itemEgressTo := flattenNestedAccessContextManagerServicePerimeterEgressPolicyEgressTo(item["egressTo"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemEgressTo)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedEgressTo))) && !reflect.DeepEqual(itemEgressTo, expectedFlattenedEgressTo) {
			log.Printf("[DEBUG] Skipping item with egressTo= %#v, looking for %#v)", itemEgressTo, expectedFlattenedEgressTo)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}

// PatchCreateEncoder handles creating request data to PATCH parent resource
// with list including new object.
func resourceAccessContextManagerServicePerimeterEgressPolicyPatchCreateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceAccessContextManagerServicePerimeterEgressPolicyListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	_, found, err := resourceAccessContextManagerServicePerimeterEgressPolicyFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}

	// Return error if item already created.
	if found != nil {
		return nil, fmt.Errorf("Unable to create ServicePerimeterEgressPolicy, existing object already found: %+v", found)
	}

	// Return list with the resource to create appended
	res := map[string]interface{}{
		"egressPolicies": append(currItems, obj),
	}
	wrapped := map[string]interface{}{
		"status": res,
	}
	res = wrapped

	return res, nil
}

// PatchUpdateEncoder handles creating request data to PATCH parent resource
// with list including updated object.
func resourceAccessContextManagerServicePerimeterEgressPolicyPatchUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	items, err := resourceAccessContextManagerServicePerimeterEgressPolicyListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, item, err := resourceAccessContextManagerServicePerimeterEgressPolicyFindNestedObjectInList(d, meta, items)
	if err != nil {
		return nil, err
	}

	// Return error if item to update does not exist.
	if item == nil {
		return nil, fmt.Errorf("Unable to update ServicePerimeterEgressPolicy %q - not found in list", d.Id())
	}

	// Merge new object into old.
	for k, v := range obj {
		item[k] = v
	}
	items[idx] = item

	// Return list with new item added
	res := map[string]interface{}{
		"egressPolicies": items,
	}
	wrapped := map[string]interface{}{
		"status": res,
	}
	res = wrapped

	return res, nil
}

// PatchDeleteEncoder handles creating request data to PATCH parent resource
// with list excluding object to delete.
func resourceAccessContextManagerServicePerimeterEgressPolicyPatchDeleteEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceAccessContextManagerServicePerimeterEgressPolicyListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, item, err := resourceAccessContextManagerServicePerimeterEgressPolicyFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}
	if item == nil {
		// Spoof 404 error for proper handling by Delete (i.e. no-op)
		return nil, tpgresource.Fake404("nested", "AccessContextManagerServicePerimeterEgressPolicy")
	}

	updatedItems := append(currItems[:idx], currItems[idx+1:]...)
	res := map[string]interface{}{
		"egressPolicies": updatedItems,
	}
	wrapped := map[string]interface{}{
		"status": res,
	}
	res = wrapped

	return res, nil
}

// ListForPatch handles making API request to get parent resource and
// extracting list of objects.
func resourceAccessContextManagerServicePerimeterEgressPolicyListForPatch(d *schema.ResourceData, meta interface{}) ([]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	url, err := tpgresource.ReplaceVars(d, config, "{{AccessContextManagerBasePath}}{{perimeter}}")
	if err != nil {
		return nil, err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return nil, err
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return nil, err
	}

	var v interface{}
	var ok bool
	if v, ok = res["status"]; ok && v != nil {
		res = v.(map[string]interface{})
	} else {
		return nil, nil
	}

	v, ok = res["egressPolicies"]
	if ok && v != nil {
		ls, lsOk := v.([]interface{})
		if !lsOk {
			return nil, fmt.Errorf(`expected list for nested field "egressPolicies"`)
		}
		return ls, nil
	}
	return nil, nil
}
