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
)

func TestAccCloudTasksQueue_queueBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudTasksQueueDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudTasksQueue_queueBasicExample(context),
			},
			{
				ResourceName:            "google_cloud_tasks_queue.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccCloudTasksQueue_queueBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_tasks_queue" "default" {
  name = "tf-test-cloud-tasks-queue-test%{random_suffix}"
  location = "us-central1"
}
`, context)
}

func TestAccCloudTasksQueue_cloudTasksQueueAdvancedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudTasksQueueDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudTasksQueue_cloudTasksQueueAdvancedExample(context),
			},
			{
				ResourceName:            "google_cloud_tasks_queue.advanced_configuration",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "app_engine_routing_override.0.service", "app_engine_routing_override.0.version", "app_engine_routing_override.0.instance"},
			},
		},
	})
}

func testAccCloudTasksQueue_cloudTasksQueueAdvancedExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_tasks_queue" "advanced_configuration" {
  name = "tf-test-instance-name%{random_suffix}"
  location = "us-central1"

  app_engine_routing_override {
    service = "worker"
    version = "1.0"
    instance = "test"
  }

  rate_limits {
    max_concurrent_dispatches = 3
    max_dispatches_per_second = 2
  }

  retry_config {
    max_attempts = 5
    max_retry_duration = "4s"
    max_backoff = "3s"
    min_backoff = "2s"
    max_doublings = 1
  }

  stackdriver_logging_config {
    sampling_ratio = 0.9
  }
}
`, context)
}

func testAccCheckCloudTasksQueueDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloud_tasks_queue" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{CloudTasksBasePath}}projects/{{project}}/locations/{{location}}/queues/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("CloudTasksQueue still exists at %s", url)
			}
		}

		return nil
	}
}