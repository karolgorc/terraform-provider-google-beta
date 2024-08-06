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

package dns_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccDNSManagedZone_dnsManagedZoneQuickstartExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"dns_name":      "m-z.gcp.tfacc.hashicorptest.com.",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZoneQuickstartExample(context),
			},
			{
				ResourceName:            "google_dns_managed_zone.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_destroy", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZoneQuickstartExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
# to setup a web-server
resource "google_compute_instance" "default" {
  name         = "tf-test-dns-compute-instance%{random_suffix}"
  machine_type = "g1-small"
  zone         = "us-central1-b"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral public IP
    }
  }
  metadata_startup_script = <<-EOF
  sudo apt-get update && \
  sudo apt-get install apache2 -y && \
  echo "<!doctype html><html><body><h1>Hello World!</h1></body></html>" > /var/www/html/index.html
  EOF
}

# to allow http traffic
resource "google_compute_firewall" "default" {
  name     = "tf-test-allow-http-traffic%{random_suffix}"
  network  = "default"
  allow {
    ports    = ["80"]
    protocol = "tcp"
  }
  source_ranges = ["0.0.0.0/0"]
}

# to create a DNS zone
resource "google_dns_managed_zone" "default" {
  name          = "tf-test-example-zone-googlecloudexample%{random_suffix}"
  dns_name      = "%{dns_name}"
  description   = "Example DNS zone"
  force_destroy = "true"
}

# to register web-server's ip address in DNS
resource "google_dns_record_set" "default" {
  name         = google_dns_managed_zone.default.dns_name
  managed_zone = google_dns_managed_zone.default.name
  type         = "A"
  ttl          = 300
  rrdatas = [
    google_compute_instance.default.network_interface.0.access_config.0.nat_ip
  ]
}
`, context)
}

func TestAccDNSManagedZone_dnsRecordSetBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsRecordSetBasicExample(context),
			},
			{
				ResourceName:            "google_dns_managed_zone.parent-zone",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccDNSManagedZone_dnsRecordSetBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dns_managed_zone" "parent-zone" {
  name        = "tf-test-sample-zone%{random_suffix}"
  dns_name    = "tf-test-sample-zone%{random_suffix}.hashicorptest.com."
  description = "Test Description"
}

resource "google_dns_record_set" "default" {
  managed_zone = google_dns_managed_zone.parent-zone.name
  name         = "test-record.tf-test-sample-zone%{random_suffix}.hashicorptest.com."
  type         = "A"
  rrdatas      = ["10.0.0.1", "10.1.0.1"]
  ttl          = 86400
}
`, context)
}

func TestAccDNSManagedZone_dnsManagedZoneBasicExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZoneBasicExample(context),
			},
			{
				ResourceName:            "google_dns_managed_zone.example-zone",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZoneBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dns_managed_zone" "example-zone" {
  name        = "example-zone"
  dns_name    = "example-${random_id.rnd.hex}.com."
  description = "Example DNS zone"
  labels = {
    foo = "bar"
  }
}

resource "random_id" "rnd" {
  byte_length = 4
}
`, context)
}

func TestAccDNSManagedZone_dnsManagedZonePrivateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZonePrivateExample(context),
			},
			{
				ResourceName:            "google_dns_managed_zone.private-zone",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZonePrivateExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dns_managed_zone" "private-zone" {
  name        = "tf-test-private-zone%{random_suffix}"
  dns_name    = "private.example.com."
  description = "Example private DNS zone"
  labels = {
    foo = "bar"
  }

  visibility = "private"

  private_visibility_config {
    networks {
      network_url = google_compute_network.network-1.id
    }
    networks {
      network_url = google_compute_network.network-2.id
    }
  }
}

resource "google_compute_network" "network-1" {
  name                    = "tf-test-network-1%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_network" "network-2" {
  name                    = "tf-test-network-2%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccDNSManagedZone_dnsManagedZonePrivateMultiprojectExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZonePrivateMultiprojectExample(context),
			},
			{
				ResourceName:            "google_dns_managed_zone.private-zone",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZonePrivateMultiprojectExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dns_managed_zone" "private-zone" {
  name        = "tf-test-private-zone%{random_suffix}"
  dns_name    = "multiproject.private.example.com."
  description = "Example private DNS zone"
  labels = {
    foo = "bar"
  }

  visibility = "private"

  private_visibility_config {
    networks {
      network_url = google_compute_network.network_1_project_1.id
    }
    networks {
      network_url = google_compute_network.network_2_project_1.id
    }
    networks {
      network_url = google_compute_network.network_1_project_2.id
    }
    networks {
      network_url = google_compute_network.network_2_project_2.id
    }
  }

  depends_on = [
    google_project_service.compute_project_1,
    google_project_service.dns_project_1,
    google_project_service.compute_project_2,
    google_project_service.dns_project_2,
  ]
}

resource "google_project" "project_1" {
  name            = "tf-test-project-1%{random_suffix}"
  project_id      = "tf-test-project-1%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project" "project_2" {
  name            = "tf-test-project-2%{random_suffix}"
  project_id      = "tf-test-project-2%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_compute_network" "network_1_project_1" {
  name                    = "tf-test-network-1%{random_suffix}"
  project                 = google_project.project_1.project_id
  auto_create_subnetworks = false
  depends_on              = [ 
    google_project_service.compute_project_1,
    google_project_service.dns_project_1,
  ]
}

resource "google_compute_network" "network_2_project_1" {
  name                    = "tf-test-network-2%{random_suffix}"
  project                 = google_project.project_1.project_id
  auto_create_subnetworks = false
  depends_on              = [ 
    google_project_service.compute_project_1,
    google_project_service.dns_project_1,
  ]
}

resource "google_compute_network" "network_1_project_2" {
  name                    = "tf-test-network-1%{random_suffix}"
  project                 = google_project.project_2.project_id
  auto_create_subnetworks = false
  depends_on              = [ 
    google_project_service.compute_project_2,
    google_project_service.dns_project_2,
  ]
}

resource "google_compute_network" "network_2_project_2" {
  name                    = "tf-test-network-2%{random_suffix}"
  project                 = google_project.project_2.project_id
  auto_create_subnetworks = false
  depends_on              = [ 
    google_project_service.compute_project_2,
    google_project_service.dns_project_2,
  ]
}

resource "google_project_service" "compute_project_1" {
  project    = google_project.project_1.project_id
  service    = "compute.googleapis.com"
  depends_on = [
    google_project.project_1,
  ]
}

resource "google_project_service" "compute_project_2" {
  project    = google_project.project_2.project_id
  service    = "compute.googleapis.com"
  depends_on = [
    google_project_service.dns_project_1
  ]
}

resource "google_project_service" "dns_project_1" {
  project    = google_project.project_1.project_id
  service    = "dns.googleapis.com"
  depends_on = [
    google_project_service.compute_project_1
  ]
}

resource "google_project_service" "dns_project_2" {
  project    = google_project.project_2.project_id
  service    = "dns.googleapis.com"
  depends_on = [
    google_project_service.compute_project_2,
  ]
}
`, context)
}

func TestAccDNSManagedZone_dnsManagedZonePrivateGkeExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZonePrivateGkeExample(context),
			},
			{
				ResourceName:            "google_dns_managed_zone.private-zone-gke",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZonePrivateGkeExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dns_managed_zone" "private-zone-gke" {
  name        = "tf-test-private-zone%{random_suffix}"
  dns_name    = "private.example.com."
  description = "Example private DNS zone"
  labels = {
    foo = "bar"
  }

  visibility = "private"

  private_visibility_config {
    gke_clusters {
      gke_cluster_name = google_container_cluster.cluster-1.id
    }
  }
}

resource "google_compute_network" "network-1" {
  name                    = "tf-test-network-1%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnetwork-1" {
  name                     = google_compute_network.network-1.name
  network                  = google_compute_network.network-1.name
  ip_cidr_range            = "10.0.36.0/24"
  region                   = "us-central1"
  private_ip_google_access = true

  secondary_ip_range {
    range_name    = "pod"
    ip_cidr_range = "10.0.0.0/19"
  }

  secondary_ip_range {
    range_name    = "svc"
    ip_cidr_range = "10.0.32.0/22"
  }
}

resource "google_container_cluster" "cluster-1" {
  name               = "tf-test-cluster-1%{random_suffix}"
  location           = "us-central1-c"
  initial_node_count = 1

  networking_mode = "VPC_NATIVE"
  default_snat_status {
    disabled = true
  }
  network    = google_compute_network.network-1.name
  subnetwork = google_compute_subnetwork.subnetwork-1.name

  private_cluster_config {
    enable_private_endpoint = true
    enable_private_nodes    = true
    master_ipv4_cidr_block  = "10.42.0.0/28"
    master_global_access_config {
      enabled = true
	}
  }
  master_authorized_networks_config {
  }
  ip_allocation_policy {
    cluster_secondary_range_name  = google_compute_subnetwork.subnetwork-1.secondary_ip_range[0].range_name
    services_secondary_range_name = google_compute_subnetwork.subnetwork-1.secondary_ip_range[1].range_name
  }
  deletion_protection  = "%{deletion_protection}"
}
`, context)
}

func TestAccDNSManagedZone_dnsManagedZonePrivatePeeringExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZonePrivatePeeringExample(context),
			},
			{
				ResourceName:            "google_dns_managed_zone.peering-zone",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZonePrivatePeeringExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dns_managed_zone" "peering-zone" {
  name        = "tf-test-peering-zone%{random_suffix}"
  dns_name    = "peering.example.com."
  description = "Example private DNS peering zone"

  visibility = "private"

  private_visibility_config {
    networks {
      network_url = google_compute_network.network-source.id
    }
  }

  peering_config {
    target_network {
      network_url = google_compute_network.network-target.id
    }
  }
}

resource "google_compute_network" "network-source" {
  name                    = "tf-test-network-source%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_network" "network-target" {
  name                    = "tf-test-network-target%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccDNSManagedZone_dnsManagedZoneServiceDirectoryExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZoneServiceDirectoryExample(context),
			},
			{
				ResourceName:            "google_dns_managed_zone.sd-zone",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZoneServiceDirectoryExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dns_managed_zone" "sd-zone" {
  provider = google-beta

  name        = "tf-test-peering-zone%{random_suffix}"
  dns_name    = "services.example.com."
  description = "Example private DNS Service Directory zone"

  visibility = "private"

  service_directory_config {
    namespace {
      namespace_url = google_service_directory_namespace.example.id
    }
  }
}

resource "google_service_directory_namespace" "example" {
  provider = google-beta

  namespace_id = "example"
  location     = "us-central1"
}

resource "google_compute_network" "network" {
  provider = google-beta

  name                    = "network%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccDNSManagedZone_dnsManagedZoneCloudLoggingExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDNSManagedZoneDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSManagedZone_dnsManagedZoneCloudLoggingExample(context),
			},
			{
				ResourceName:            "google_dns_managed_zone.cloud-logging-enabled-zone",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "terraform_labels"},
			},
		},
	})
}

func testAccDNSManagedZone_dnsManagedZoneCloudLoggingExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dns_managed_zone" "cloud-logging-enabled-zone" {
  name        = "tf-test-cloud-logging-enabled-zone%{random_suffix}"
  dns_name    = "services.example.com."
  description = "Example cloud logging enabled DNS zone"
  labels = {
    foo = "bar"
  }

  cloud_logging_config {
    enable_logging = true
  }
}
`, context)
}

func testAccCheckDNSManagedZoneDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dns_managed_zone" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DNSBasePath}}projects/{{project}}/managedZones/{{name}}")
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
				return fmt.Errorf("DNSManagedZone still exists at %s", url)
			}
		}

		return nil
	}
}
