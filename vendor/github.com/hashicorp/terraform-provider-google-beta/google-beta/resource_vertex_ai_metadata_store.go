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

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVertexAIMetadataStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAIMetadataStoreCreate,
		Read:   resourceVertexAIMetadataStoreRead,
		Delete: resourceVertexAIMetadataStoreDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAIMetadataStoreImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Description of the MetadataStore.`,
			},
			"encryption_spec": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Customer-managed encryption key spec for a MetadataStore. If set, this MetadataStore and all sub-resources of this MetadataStore will be secured by this key.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Description: `Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.`,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The name of the MetadataStore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.`,
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of the Metadata Store. eg us-central1`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the MetadataStore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"state": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `State information of the MetadataStore.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disk_utilization_bytes": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The disk utilization of the MetadataStore in bytes.`,
						},
					},
				},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the MetadataStore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
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

func resourceVertexAIMetadataStoreCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandVertexAIMetadataStoreDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	encryptionSpecProp, err := expandVertexAIMetadataStoreEncryptionSpec(d.Get("encryption_spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("encryption_spec"); !isEmptyValue(reflect.ValueOf(encryptionSpecProp)) && (ok || !reflect.DeepEqual(v, encryptionSpecProp)) {
		obj["encryptionSpec"] = encryptionSpecProp
	}

	url, err := replaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/metadataStores?metadataStoreId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new MetadataStore: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MetadataStore: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating MetadataStore: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = vertexAIOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating MetadataStore", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create MetadataStore: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating MetadataStore %q: %#v", d.Id(), res)

	return resourceVertexAIMetadataStoreRead(d, meta)
}

func resourceVertexAIMetadataStoreRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/metadataStores/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MetadataStore: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("VertexAIMetadataStore %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading MetadataStore: %s", err)
	}

	if err := d.Set("description", flattenVertexAIMetadataStoreDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading MetadataStore: %s", err)
	}
	if err := d.Set("create_time", flattenVertexAIMetadataStoreCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading MetadataStore: %s", err)
	}
	if err := d.Set("update_time", flattenVertexAIMetadataStoreUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading MetadataStore: %s", err)
	}
	if err := d.Set("encryption_spec", flattenVertexAIMetadataStoreEncryptionSpec(res["encryptionSpec"], d, config)); err != nil {
		return fmt.Errorf("Error reading MetadataStore: %s", err)
	}
	if err := d.Set("state", flattenVertexAIMetadataStoreState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading MetadataStore: %s", err)
	}

	return nil
}

func resourceVertexAIMetadataStoreDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for MetadataStore: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/metadataStores/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting MetadataStore %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "MetadataStore")
	}

	err = vertexAIOperationWaitTime(
		config, res, project, "Deleting MetadataStore", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting MetadataStore %q: %#v", d.Id(), res)
	return nil
}

func resourceVertexAIMetadataStoreImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/metadataStores/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenVertexAIMetadataStoreDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIMetadataStoreCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIMetadataStoreUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIMetadataStoreEncryptionSpec(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_name"] =
		flattenVertexAIMetadataStoreEncryptionSpecKmsKeyName(original["kmsKeyName"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIMetadataStoreEncryptionSpecKmsKeyName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenVertexAIMetadataStoreState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["disk_utilization_bytes"] =
		flattenVertexAIMetadataStoreStateDiskUtilizationBytes(original["diskUtilizationBytes"], d, config)
	return []interface{}{transformed}
}
func flattenVertexAIMetadataStoreStateDiskUtilizationBytes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandVertexAIMetadataStoreDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIMetadataStoreEncryptionSpec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandVertexAIMetadataStoreEncryptionSpecKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !isEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandVertexAIMetadataStoreEncryptionSpecKmsKeyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
