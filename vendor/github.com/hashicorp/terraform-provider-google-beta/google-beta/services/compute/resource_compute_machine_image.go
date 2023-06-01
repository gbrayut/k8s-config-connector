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

package compute

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceComputeMachineImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeMachineImageCreate,
		Read:   resourceComputeMachineImageRead,
		Delete: resourceComputeMachineImageDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeMachineImageImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the resource.`,
			},
			"source_instance": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `A text description of the resource.`,
			},
			"guest_flush": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `Specify this to create an application consistent machine image by informing the OS to prepare for the snapshot process.
Currently only supported on Windows instances using the Volume Shadow Copy Service (VSS).`,
			},
			"machine_image_encryption_key": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Encrypts the machine image using a customer-supplied encryption key.

After you encrypt a machine image with a customer-supplied key, you must
provide the same key if you use the machine image later (e.g. to create a
instance from the image)`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_name": {
							Type:             schema.TypeString,
							Optional:         true,
							ForceNew:         true,
							DiffSuppressFunc: tpgresource.CompareCryptoKeyVersions,
							Description:      `The name of the encryption key that is stored in Google Cloud KMS.`,
						},
						"kms_key_service_account": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `The service account used for the encryption request for the given KMS key.
If absent, the Compute Engine Service Agent service account is used.`,
						},
						"raw_key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Specifies a 256-bit customer-supplied encryption key, encoded in
RFC 4648 base64 to either encrypt or decrypt this resource.`,
						},
						"sha256": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The RFC 4648 base64 encoded SHA-256 hash of the
customer-supplied encryption key that protects this resource.`,
						},
					},
				},
			},
			"storage_locations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The regional or multi-regional Cloud Storage bucket location where the machine image is stored.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeMachineImageCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeMachineImageName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeMachineImageDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sourceInstanceProp, err := expandComputeMachineImageSourceInstance(d.Get("source_instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_instance"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceInstanceProp)) && (ok || !reflect.DeepEqual(v, sourceInstanceProp)) {
		obj["sourceInstance"] = sourceInstanceProp
	}
	guestFlushProp, err := expandComputeMachineImageGuestFlush(d.Get("guest_flush"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("guest_flush"); !tpgresource.IsEmptyValue(reflect.ValueOf(guestFlushProp)) && (ok || !reflect.DeepEqual(v, guestFlushProp)) {
		obj["guestFlush"] = guestFlushProp
	}
	machineImageEncryptionKeyProp, err := expandComputeMachineImageMachineImageEncryptionKey(d.Get("machine_image_encryption_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("machine_image_encryption_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(machineImageEncryptionKeyProp)) && (ok || !reflect.DeepEqual(v, machineImageEncryptionKeyProp)) {
		obj["machineImageEncryptionKey"] = machineImageEncryptionKeyProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/machineImages")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new MachineImage: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MachineImage: %s", err)
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
		return fmt.Errorf("Error creating MachineImage: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/machineImages/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating MachineImage", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create MachineImage: %s", err)
	}

	log.Printf("[DEBUG] Finished creating MachineImage %q: %#v", d.Id(), res)

	return resourceComputeMachineImageRead(d, meta)
}

func resourceComputeMachineImageRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/machineImages/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MachineImage: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeMachineImage %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading MachineImage: %s", err)
	}

	if err := d.Set("name", flattenComputeMachineImageName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading MachineImage: %s", err)
	}
	if err := d.Set("description", flattenComputeMachineImageDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading MachineImage: %s", err)
	}
	if err := d.Set("source_instance", flattenComputeMachineImageSourceInstance(res["sourceInstance"], d, config)); err != nil {
		return fmt.Errorf("Error reading MachineImage: %s", err)
	}
	if err := d.Set("storage_locations", flattenComputeMachineImageStorageLocations(res["storageLocations"], d, config)); err != nil {
		return fmt.Errorf("Error reading MachineImage: %s", err)
	}
	if err := d.Set("guest_flush", flattenComputeMachineImageGuestFlush(res["guestFlush"], d, config)); err != nil {
		return fmt.Errorf("Error reading MachineImage: %s", err)
	}
	if err := d.Set("machine_image_encryption_key", flattenComputeMachineImageMachineImageEncryptionKey(res["machineImageEncryptionKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading MachineImage: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading MachineImage: %s", err)
	}

	return nil
}

func resourceComputeMachineImageDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MachineImage: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/machineImages/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting MachineImage %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "MachineImage")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting MachineImage", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting MachineImage %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeMachineImageImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/global/machineImages/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/global/machineImages/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeMachineImageName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeMachineImageDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeMachineImageSourceInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeMachineImageStorageLocations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeMachineImageGuestFlush(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeMachineImageMachineImageEncryptionKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["raw_key"] =
		flattenComputeMachineImageMachineImageEncryptionKeyRawKey(original["rawKey"], d, config)
	transformed["sha256"] =
		flattenComputeMachineImageMachineImageEncryptionKeySha256(original["sha256"], d, config)
	transformed["kms_key_name"] =
		flattenComputeMachineImageMachineImageEncryptionKeyKmsKeyName(original["kmsKeyName"], d, config)
	transformed["kms_key_service_account"] =
		flattenComputeMachineImageMachineImageEncryptionKeyKmsKeyServiceAccount(original["kmsKeyServiceAccount"], d, config)
	return []interface{}{transformed}
}
func flattenComputeMachineImageMachineImageEncryptionKeyRawKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeMachineImageMachineImageEncryptionKeySha256(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeMachineImageMachineImageEncryptionKeyKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeMachineImageMachineImageEncryptionKeyKmsKeyServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandComputeMachineImageName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeMachineImageDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeMachineImageSourceInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseZonalFieldValue("instances", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for source_instance: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeMachineImageGuestFlush(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeMachineImageMachineImageEncryptionKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRawKey, err := expandComputeMachineImageMachineImageEncryptionKeyRawKey(original["raw_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRawKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rawKey"] = transformedRawKey
	}

	transformedSha256, err := expandComputeMachineImageMachineImageEncryptionKeySha256(original["sha256"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSha256); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sha256"] = transformedSha256
	}

	transformedKmsKeyName, err := expandComputeMachineImageMachineImageEncryptionKeyKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	transformedKmsKeyServiceAccount, err := expandComputeMachineImageMachineImageEncryptionKeyKmsKeyServiceAccount(original["kms_key_service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyServiceAccount"] = transformedKmsKeyServiceAccount
	}

	return transformed, nil
}

func expandComputeMachineImageMachineImageEncryptionKeyRawKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeMachineImageMachineImageEncryptionKeySha256(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeMachineImageMachineImageEncryptionKeyKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeMachineImageMachineImageEncryptionKeyKmsKeyServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
