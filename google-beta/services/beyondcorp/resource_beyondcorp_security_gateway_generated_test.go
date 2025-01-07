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

package beyondcorp_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccBeyondcorpSecurityGateway_beyondcorpSecurityGatewayBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBeyondcorpSecurityGatewayDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBeyondcorpSecurityGateway_beyondcorpSecurityGatewayBasicExample(context),
			},
			{
				ResourceName:            "google_beyondcorp_security_gateway.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "security_gateway_id"},
			},
		},
	})
}

func testAccBeyondcorpSecurityGateway_beyondcorpSecurityGatewayBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_beyondcorp_security_gateway" "example" {
  security_gateway_id = "default%{random_suffix}"
  location = "global"
  display_name = "My Security Gateway resource"
  hubs { region = "us-central1" }
}
`, context)
}

func testAccCheckBeyondcorpSecurityGatewayDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_beyondcorp_security_gateway" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{BeyondcorpBasePath}}projects/{{project}}/locations/{{location}}/securityGateways/{{security_gateway_id}}")
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
				return fmt.Errorf("BeyondcorpSecurityGateway still exists at %s", url)
			}
		}

		return nil
	}
}
