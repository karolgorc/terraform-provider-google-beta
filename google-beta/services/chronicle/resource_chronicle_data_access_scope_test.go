// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package chronicle_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
)

func TestAccChronicleDataAccessScope_chronicleDataaccessscopeBasicExample_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"chronicle_id":  envvar.GetTestChronicleInstanceIdFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckChronicleDataAccessScopeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccChronicleDataAccessScope_chronicleDataaccessscopeBasicExample_full(context),
			},
			{
				ResourceName:            "google_chronicle_data_access_scope.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"data_access_scope_id", "instance", "location"},
			},

			{
				Config: testAccChronicleDataAccessScope_chronicleDataaccessscopeBasicExample_update(context),
			},
			{
				ResourceName:            "google_chronicle_data_access_scope.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"data_access_scope_id", "instance", "location"},
			},
		},
	})
}

func testAccChronicleDataAccessScope_chronicleDataaccessscopeBasicExample_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_chronicle_data_access_label" "custom_data_access_label" {
  provider = "google-beta"
  location = "us"
  instance = "%{chronicle_id}"
  data_access_label_id = "tf-test-label-id%{random_suffix}"
  udm_query = "principal.hostname=\"google.com\""
}

resource "google_chronicle_data_access_scope" "example" {
  provider = "google-beta"
  location = "us"
  instance = "%{chronicle_id}"
  data_access_scope_id = "tf-test-scope-id%{random_suffix}"
  description = "tf-test-scope-description%{random_suffix}"
  allow_all = false
  allowed_data_access_labels {
    log_type = "GITHUB"
  }
  denied_data_access_labels {
    log_type = "GCP_CLOUDAUDIT"
  }
  denied_data_access_labels {
    data_access_label = resource.google_chronicle_data_access_label.custom_data_access_label.data_access_label_id
  }
  denied_data_access_labels {
    ingestion_label {
	    ingestion_label_key = "ingestion_key"
      ingestion_label_value = "ingestion_value"
    }
  }
  denied_data_access_labels {
    asset_namespace = "my-namespace"
  }
}
`, context)
}

func testAccChronicleDataAccessScope_chronicleDataaccessscopeBasicExample_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_chronicle_data_access_label" "custom_data_access_label" {
  provider = "google-beta"
  location = "us"
  instance = "%{chronicle_id}"
  data_access_label_id = "tf-test-label-id%{random_suffix}"
  udm_query = "principal.hostname=\"google-updated.com\""
}

resource "google_chronicle_data_access_scope" "example" {
  provider = "google-beta"
  location = "us"
  instance = "%{chronicle_id}"
  data_access_scope_id = "tf-test-scope-id%{random_suffix}"
  description = "tf-test-scope-description-updated%{random_suffix}"
  allow_all = true
  denied_data_access_labels {
    log_type = "GITHUB"
  }
  denied_data_access_labels {
    data_access_label = resource.google_chronicle_data_access_label.custom_data_access_label.data_access_label_id
  }
  denied_data_access_labels {
    ingestion_label {
	    ingestion_label_key = "ingestion_key-updated"
      ingestion_label_value = "ingestion_value-updated"
    }
  }
  denied_data_access_labels {
    asset_namespace = "my-namespace-updated"
  }
}
`, context)
}