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

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccDataCatalogPolicyTag_dataCatalogTaxonomiesPolicyTagBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataCatalogPolicyTagDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataCatalogPolicyTag_dataCatalogTaxonomiesPolicyTagBasicExample(context),
			},
			{
				ResourceName:            "google_data_catalog_policy_tag.basic_policy_tag",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"taxonomy"},
			},
		},
	})
}

func testAccDataCatalogPolicyTag_dataCatalogTaxonomiesPolicyTagBasicExample(context map[string]interface{}) string {
	return tpgresource.Nprintf(`
resource "google_data_catalog_policy_tag" "basic_policy_tag" {
  taxonomy = google_data_catalog_taxonomy.my_taxonomy.id
  display_name = "Low security"
  description = "A policy tag normally associated with low security items"
}

resource "google_data_catalog_taxonomy" "my_taxonomy" {
  display_name =  "tf_test_taxonomy_display_name%{random_suffix}"
  description = "A collection of policy tags"
  activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
}
`, context)
}

func TestAccDataCatalogPolicyTag_dataCatalogTaxonomiesPolicyTagChildPoliciesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataCatalogPolicyTagDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataCatalogPolicyTag_dataCatalogTaxonomiesPolicyTagChildPoliciesExample(context),
			},
			{
				ResourceName:            "google_data_catalog_policy_tag.child_policy",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"taxonomy"},
			},
		},
	})
}

func testAccDataCatalogPolicyTag_dataCatalogTaxonomiesPolicyTagChildPoliciesExample(context map[string]interface{}) string {
	return tpgresource.Nprintf(`
resource "google_data_catalog_policy_tag" "parent_policy" {
  taxonomy = google_data_catalog_taxonomy.my_taxonomy.id
  display_name = "High"
  description = "A policy tag category used for high security access"
}

resource "google_data_catalog_policy_tag" "child_policy" {
  taxonomy = google_data_catalog_taxonomy.my_taxonomy.id
  display_name = "ssn"
  description = "A hash of the users ssn"
  parent_policy_tag = google_data_catalog_policy_tag.parent_policy.id
}

resource "google_data_catalog_policy_tag" "child_policy2" {
  taxonomy = google_data_catalog_taxonomy.my_taxonomy.id
  display_name = "dob"
  description = "The users date of birth"
  parent_policy_tag = google_data_catalog_policy_tag.parent_policy.id
  // depends_on to avoid concurrent delete issues
  depends_on = [google_data_catalog_policy_tag.child_policy]
}

resource "google_data_catalog_taxonomy" "my_taxonomy" {
  display_name =  "tf_test_taxonomy_display_name%{random_suffix}"
  description = "A collection of policy tags"
  activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
}
`, context)
}

func testAccCheckDataCatalogPolicyTagDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_data_catalog_policy_tag" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataCatalogBasePath}}{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("DataCatalogPolicyTag still exists at %s", url)
			}
		}

		return nil
	}
}
