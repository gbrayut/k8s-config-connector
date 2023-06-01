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

package firebase

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceFirebaseProjectLocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseProjectLocationCreate,
		Read:   resourceFirebaseProjectLocationRead,
		Delete: resourceFirebaseProjectLocationDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseProjectLocationImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID of the default GCP resource location for the Project. The location must be one of the available GCP
resource locations.`,
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

func resourceFirebaseProjectLocationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	locationIdProp, err := expandNestedFirebaseProjectLocationLocationId(d.Get("location_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("location_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(locationIdProp)) && (ok || !reflect.DeepEqual(v, locationIdProp)) {
		obj["locationId"] = locationIdProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}/defaultLocation:finalize")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ProjectLocation: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectLocation: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating ProjectLocation: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = FirebaseOperationWaitTime(
		config, res, project, "Creating ProjectLocation", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ProjectLocation: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ProjectLocation %q: %#v", d.Id(), res)

	return resourceFirebaseProjectLocationRead(d, meta)
}

func resourceFirebaseProjectLocationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseBasePath}}projects/{{project}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ProjectLocation: %s", err)
	}
	billingProject = project

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("FirebaseProjectLocation %q", d.Id()))
	}

	res, err = flattenNestedFirebaseProjectLocation(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing FirebaseProjectLocation because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ProjectLocation: %s", err)
	}

	if err := d.Set("location_id", flattenNestedFirebaseProjectLocationLocationId(res["locationId"], d, config)); err != nil {
		return fmt.Errorf("Error reading ProjectLocation: %s", err)
	}

	return nil
}

func resourceFirebaseProjectLocationDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Firebase ProjectLocation resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceFirebaseProjectLocationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)",
		"(?P<project>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedFirebaseProjectLocationLocationId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNestedFirebaseProjectLocationLocationId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func flattenNestedFirebaseProjectLocation(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["resources"]
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
		return nil, fmt.Errorf("expected list or map for value resources. Actual value: %v", v)
	}

	_, item, err := resourceFirebaseProjectLocationFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceFirebaseProjectLocationFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
