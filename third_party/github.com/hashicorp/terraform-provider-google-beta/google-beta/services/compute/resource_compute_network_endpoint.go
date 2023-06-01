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

func ResourceComputeNetworkEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNetworkEndpointCreate,
		Read:   resourceComputeNetworkEndpointRead,
		Delete: resourceComputeNetworkEndpointDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeNetworkEndpointImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `IPv4 address of network endpoint. The IP address must belong
to a VM in GCE (either the primary IP or as part of an aliased IP
range).`,
			},
			"network_endpoint_group": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description:      `The network endpoint group this endpoint is part of.`,
			},
			"instance": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `The name for a specific VM instance that the IP address belongs to.
This is required for network endpoints of type GCE_VM_IP_PORT.
The instance must be in the same zone of network endpoint group.`,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Description: `Port number of network endpoint.
**Note** 'port' is required unless the Network Endpoint Group is created
with the type of 'GCE_VM_IP'`,
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Zone where the containing network endpoint group is located.`,
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

func resourceComputeNetworkEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	instanceProp, err := expandNestedComputeNetworkEndpointInstance(d.Get("instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("instance"); !tpgresource.IsEmptyValue(reflect.ValueOf(instanceProp)) && (ok || !reflect.DeepEqual(v, instanceProp)) {
		obj["instance"] = instanceProp
	}
	portProp, err := expandNestedComputeNetworkEndpointPort(d.Get("port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port"); !tpgresource.IsEmptyValue(reflect.ValueOf(portProp)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	ipAddressProp, err := expandNestedComputeNetworkEndpointIpAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_address"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipAddressProp)) && (ok || !reflect.DeepEqual(v, ipAddressProp)) {
		obj["ipAddress"] = ipAddressProp
	}

	obj, err = resourceComputeNetworkEndpointEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "networkEndpoint/{{project}}/{{zone}}/{{network_endpoint_group}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/attachNetworkEndpoints")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NetworkEndpoint: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkEndpoint: %s", err)
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
		return fmt.Errorf("Error creating NetworkEndpoint: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating NetworkEndpoint", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create NetworkEndpoint: %s", err)
	}

	log.Printf("[DEBUG] Finished creating NetworkEndpoint %q: %#v", d.Id(), res)

	return resourceComputeNetworkEndpointRead(d, meta)
}

func resourceComputeNetworkEndpointRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/listNetworkEndpoints")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkEndpoint: %s", err)
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
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeNetworkEndpoint %q", d.Id()))
	}

	res, err = flattenNestedComputeNetworkEndpoint(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeNetworkEndpoint because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	res, err = resourceComputeNetworkEndpointDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ComputeNetworkEndpoint because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoint: %s", err)
	}

	if err := d.Set("instance", flattenNestedComputeNetworkEndpointInstance(res["instance"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoint: %s", err)
	}
	if err := d.Set("port", flattenNestedComputeNetworkEndpointPort(res["port"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoint: %s", err)
	}
	if err := d.Set("ip_address", flattenNestedComputeNetworkEndpointIpAddress(res["ipAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoint: %s", err)
	}

	return nil
}

func resourceComputeNetworkEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkEndpoint: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "networkEndpoint/{{project}}/{{zone}}/{{network_endpoint_group}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/detachNetworkEndpoints")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	toDelete := make(map[string]interface{})
	instanceProp, err := expandNestedComputeNetworkEndpointInstance(d.Get("instance"), d, config)
	if err != nil {
		return err
	}
	if instanceProp != "" {
		toDelete["instance"] = instanceProp
	}

	portProp, err := expandNestedComputeNetworkEndpointPort(d.Get("port"), d, config)
	if err != nil {
		return err
	}
	if portProp != 0 {
		toDelete["port"] = portProp
	}

	ipAddressProp, err := expandNestedComputeNetworkEndpointIpAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return err
	}
	toDelete["ipAddress"] = ipAddressProp

	obj = map[string]interface{}{
		"networkEndpoints": []map[string]interface{}{toDelete},
	}
	log.Printf("[DEBUG] Deleting NetworkEndpoint %q", d.Id())

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
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "NetworkEndpoint")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting NetworkEndpoint", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting NetworkEndpoint %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeNetworkEndpointImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	// instance is optional, so use * instead of + when reading the import id
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/networkEndpointGroups/(?P<network_endpoint_group>[^/]+)/(?P<instance>[^/]*)/(?P<ip_address>[^/]+)/(?P<port>[^/]+)",
		"(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<network_endpoint_group>[^/]+)/(?P<instance>[^/]*)/(?P<ip_address>[^/]+)/(?P<port>[^/]+)",
		"(?P<zone>[^/]+)/(?P<network_endpoint_group>[^/]+)/(?P<instance>[^/]*)/(?P<ip_address>[^/]+)/(?P<port>[^/]+)",
		"(?P<network_endpoint_group>[^/]+)/(?P<instance>[^/]*)/(?P<ip_address>[^/]+)/(?P<port>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedComputeNetworkEndpointInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenNestedComputeNetworkEndpointPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles int given in float64 format
	if floatVal, ok := v.(float64); ok {
		return int(floatVal)
	}
	return v
}

func flattenNestedComputeNetworkEndpointIpAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNestedComputeNetworkEndpointInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.GetResourceNameFromSelfLink(v.(string)), nil
}

func expandNestedComputeNetworkEndpointPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeNetworkEndpointIpAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceComputeNetworkEndpointEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Network Endpoint Group is a URL parameter only, so replace self-link/path with resource name only.
	if err := d.Set("network_endpoint_group", tpgresource.GetResourceNameFromSelfLink(d.Get("network_endpoint_group").(string))); err != nil {
		return nil, fmt.Errorf("Error setting network_endpoint_group: %s", err)
	}

	wrappedReq := map[string]interface{}{
		"networkEndpoints": []interface{}{obj},
	}
	return wrappedReq, nil
}

func flattenNestedComputeNetworkEndpoint(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["items"]
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
		return nil, fmt.Errorf("expected list or map for value items. Actual value: %v", v)
	}

	_, item, err := resourceComputeNetworkEndpointFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeNetworkEndpointFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedInstance, err := expandNestedComputeNetworkEndpointInstance(d.Get("instance"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedInstance := flattenNestedComputeNetworkEndpointInstance(expectedInstance, d, meta.(*transport_tpg.Config))
	expectedIpAddress, err := expandNestedComputeNetworkEndpointIpAddress(d.Get("ip_address"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedIpAddress := flattenNestedComputeNetworkEndpointIpAddress(expectedIpAddress, d, meta.(*transport_tpg.Config))
	expectedPort, err := expandNestedComputeNetworkEndpointPort(d.Get("port"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedPort := flattenNestedComputeNetworkEndpointPort(expectedPort, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		// Decode list item before comparing.
		item, err := resourceComputeNetworkEndpointDecoder(d, meta, item)
		if err != nil {
			return -1, nil, err
		}

		itemInstance := flattenNestedComputeNetworkEndpointInstance(item["instance"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemInstance)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedInstance))) && !reflect.DeepEqual(itemInstance, expectedFlattenedInstance) {
			log.Printf("[DEBUG] Skipping item with instance= %#v, looking for %#v)", itemInstance, expectedFlattenedInstance)
			continue
		}
		itemIpAddress := flattenNestedComputeNetworkEndpointIpAddress(item["ipAddress"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemIpAddress)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedIpAddress))) && !reflect.DeepEqual(itemIpAddress, expectedFlattenedIpAddress) {
			log.Printf("[DEBUG] Skipping item with ipAddress= %#v, looking for %#v)", itemIpAddress, expectedFlattenedIpAddress)
			continue
		}
		itemPort := flattenNestedComputeNetworkEndpointPort(item["port"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemPort)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedPort))) && !reflect.DeepEqual(itemPort, expectedFlattenedPort) {
			log.Printf("[DEBUG] Skipping item with port= %#v, looking for %#v)", itemPort, expectedFlattenedPort)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
func resourceComputeNetworkEndpointDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	v, ok := res["networkEndpoint"]
	if !ok || v == nil {
		return res, nil
	}

	return v.(map[string]interface{}), nil
}
