// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package networksecurity_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccNetworkSecurityMirroringEndpointGroup_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityMirroringEndpointGroup_basic(context),
			},
			{
				ResourceName:            "google_network_security_mirroring_endpoint_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
			{
				Config: testAccNetworkSecurityMirroringEndpointGroup_update(context),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("google_network_security_mirroring_endpoint_group.default", plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				ResourceName:            "google_network_security_mirroring_endpoint_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"update_time", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecurityMirroringEndpointGroup_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "network" {
  provider                = google-beta
  name                    = "tf-test-example-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_network_security_mirroring_deployment_group" "deployment_group" {
  provider                      = google-beta
  mirroring_deployment_group_id = "tf-test-example-dg%{random_suffix}"
  location                      = "global"
  network                       = google_compute_network.network.id
}

resource "google_network_security_mirroring_endpoint_group" "default" {
  provider                      = google-beta
  mirroring_endpoint_group_id   = "tf-test-example-eg%{random_suffix}"
  location                      = "global"
  mirroring_deployment_group    = google_network_security_mirroring_deployment_group.deployment_group.id
  description                   = "initial description"
  labels = {
    foo = "bar"
  }
}
`, context)
}

func testAccNetworkSecurityMirroringEndpointGroup_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "network" {
  provider                = google-beta
  name                    = "tf-test-example-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_network_security_mirroring_deployment_group" "deployment_group" {
  provider                      = google-beta
  mirroring_deployment_group_id = "tf-test-example-dg%{random_suffix}"
  location                      = "global"
  network                       = google_compute_network.network.id
}

resource "google_network_security_mirroring_endpoint_group" "default" {
  provider                      = google-beta
  mirroring_endpoint_group_id   = "tf-test-example-eg%{random_suffix}"
  location                      = "global"
  mirroring_deployment_group    = google_network_security_mirroring_deployment_group.deployment_group.id
  description                   = "updated description"
  labels = {
    foo = "goo"
  }
}
`, context)
}
