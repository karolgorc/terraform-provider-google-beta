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

package healthcare_test

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

func TestAccHealthcareDicomStore_healthcareDicomStoreBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckHealthcareDicomStoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareDicomStore_healthcareDicomStoreBasicExample(context),
			},
			{
				ResourceName:            "google_healthcare_dicom_store.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"dataset", "labels", "self_link", "terraform_labels"},
			},
		},
	})
}

func testAccHealthcareDicomStore_healthcareDicomStoreBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_healthcare_dicom_store" "default" {
  name    = "tf-test-example-dicom-store%{random_suffix}"
  dataset = google_healthcare_dataset.dataset.id

  notification_config {
    pubsub_topic = google_pubsub_topic.topic.id
  }

  labels = {
    label1 = "labelvalue1"
  }
}

resource "google_pubsub_topic" "topic" {
  name     = "tf-test-dicom-notifications%{random_suffix}"
}

resource "google_healthcare_dataset" "dataset" {
  name     = "tf-test-example-dataset%{random_suffix}"
  location = "us-central1"
}
`, context)
}

func TestAccHealthcareDicomStore_healthcareDicomStoreBqStreamExample(t *testing.T) {
	t.Parallel()
	acctest.BootstrapIamMembers(t, []acctest.IamMember{
		{
			Member: "serviceAccount:service-{project_number}@gcp-sa-healthcare.iam.gserviceaccount.com",
			Role:   "roles/bigquery.dataEditor",
		},
		{
			Member: "serviceAccount:service-{project_number}@gcp-sa-healthcare.iam.gserviceaccount.com",
			Role:   "roles/bigquery.jobUser",
		},
	})

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckHealthcareDicomStoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareDicomStore_healthcareDicomStoreBqStreamExample(context),
			},
			{
				ResourceName:            "google_healthcare_dicom_store.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"dataset", "labels", "self_link", "terraform_labels"},
			},
		},
	})
}

func testAccHealthcareDicomStore_healthcareDicomStoreBqStreamExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_healthcare_dicom_store" "default" {
  provider = google-beta

  name    = "tf-test-example-dicom-store%{random_suffix}"
  dataset = google_healthcare_dataset.dataset.id

  notification_config {
    pubsub_topic         = google_pubsub_topic.topic.id
    send_for_bulk_import = true
  }

  labels = {
    label1 = "labelvalue1"
  }

  stream_configs {
    bigquery_destination {
      table_uri = "bq://${google_bigquery_dataset.bq_dataset.project}.${google_bigquery_dataset.bq_dataset.dataset_id}.${google_bigquery_table.bq_table.table_id}"
    }
  }  
}

resource "google_pubsub_topic" "topic" {
  provider = google-beta

  name     = "tf-test-dicom-notifications%{random_suffix}"
}

resource "google_healthcare_dataset" "dataset" {
  provider = google-beta

  name     = "tf-test-example-dataset%{random_suffix}"
  location = "us-central1"
}

resource "google_bigquery_dataset" "bq_dataset" {
  provider = google-beta

  dataset_id    = "tf_test_dicom_bq_ds%{random_suffix}"
  friendly_name = "test"
  description   = "This is a test description"
  location      = "US"
  delete_contents_on_destroy = true
}

resource "google_bigquery_table" "bq_table" {
  provider = google-beta

  deletion_protection = false
  dataset_id = google_bigquery_dataset.bq_dataset.dataset_id
  table_id   = "tf_test_dicom_bq_tb%{random_suffix}"
}
`, context)
}

func testAccCheckHealthcareDicomStoreDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_healthcare_dicom_store" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{HealthcareBasePath}}{{dataset}}/dicomStores/{{name}}")
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
				return fmt.Errorf("HealthcareDicomStore still exists at %s", url)
			}
		}

		return nil
	}
}
