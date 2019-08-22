// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataprocAutoscalingPolicy_dataprocAutoscalingPolicyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckDataprocAutoscalingPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocAutoscalingPolicy_dataprocAutoscalingPolicyExample(context),
			},
		},
	})
}

func testAccDataprocAutoscalingPolicy_dataprocAutoscalingPolicyExample(context map[string]interface{}) string {
	return Nprintf(`
provider "google-beta" {}

resource "google_dataproc_cluster" "basic" {
  provider = "google-beta"
  name   = "tf-dataproc-test-%{random_suffix}"
  region = "us-central1"

  cluster_config {
    autoscaling_config {
      policy_uri = google_dataproc_autoscaling_policy.asp.name
    }
  }
}

resource "google_dataproc_autoscaling_policy" "asp" {
  provider = "google-beta"
  policy_id = "tf-dataproc-test-%{random_suffix}"
  location = "us-central1"

  worker_config {
    max_instances = 3
  }

  basic_algorithm {
    yarn_config {
      graceful_decommission_timeout = "30s"

      scale_up_factor   = 0.5
      scale_down_factor = 0.5
    }
  }
}
`, context)
}

func testAccCheckDataprocAutoscalingPolicyDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_dataproc_autoscaling_policy" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{DataprocBasePath}}projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("DataprocAutoscalingPolicy still exists at %s", url)
		}
	}

	return nil
}
