// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dataproc/AutoscalingPolicy.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dataproc

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceDataprocAutoscalingPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataprocAutoscalingPolicyCreate,
		Read:   resourceDataprocAutoscalingPolicyRead,
		Update: resourceDataprocAutoscalingPolicyUpdate,
		Delete: resourceDataprocAutoscalingPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataprocAutoscalingPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"policy_id": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
3 and 50 characters.`,
			},
			"basic_algorithm": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Basic algorithm for autoscaling.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"yarn_config": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `YARN autoscaling configuration.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"graceful_decommission_timeout": {
										Type:     schema.TypeString,
										Required: true,
										Description: `Timeout for YARN graceful decommissioning of Node Managers. Specifies the
duration to wait for jobs to complete before forcefully removing workers
(and potentially interrupting jobs). Only applicable to downscaling operations.

Bounds: [0s, 1d].`,
									},
									"scale_down_factor": {
										Type:     schema.TypeFloat,
										Required: true,
										Description: `Fraction of average pending memory in the last cooldown period for which to
remove workers. A scale-down factor of 1 will result in scaling down so that there
is no available memory remaining after the update (more aggressive scaling).
A scale-down factor of 0 disables removing workers, which can be beneficial for
autoscaling a single job.

Bounds: [0.0, 1.0].`,
									},
									"scale_up_factor": {
										Type:     schema.TypeFloat,
										Required: true,
										Description: `Fraction of average pending memory in the last cooldown period for which to
add workers. A scale-up factor of 1.0 will result in scaling up so that there
is no pending memory remaining after the update (more aggressive scaling).
A scale-up factor closer to 0 will result in a smaller magnitude of scaling up
(less aggressive scaling).

Bounds: [0.0, 1.0].`,
									},
									"scale_down_min_worker_fraction": {
										Type:     schema.TypeFloat,
										Optional: true,
										Description: `Minimum scale-down threshold as a fraction of total cluster size before scaling occurs.
For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler must
recommend at least a 2 worker scale-down for the cluster to scale. A threshold of 0
means the autoscaler will scale down on any recommended change.

Bounds: [0.0, 1.0]. Default: 0.0.`,
										Default: 0.0,
									},
									"scale_up_min_worker_fraction": {
										Type:     schema.TypeFloat,
										Optional: true,
										Description: `Minimum scale-up threshold as a fraction of total cluster size before scaling
occurs. For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler
must recommend at least a 2-worker scale-up for the cluster to scale. A threshold of
0 means the autoscaler will scale up on any recommended change.

Bounds: [0.0, 1.0]. Default: 0.0.`,
										Default: 0.0,
									},
								},
							},
						},
						"cooldown_period": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Duration between scaling events. A scaling period starts after the
update operation from the previous event has completed.

Bounds: [2m, 1d]. Default: 2m.`,
							Default: "120s",
						},
					},
				},
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The  location where the autoscaling policy should reside.
The default value is 'global'.`,
				Default: "global",
			},
			"secondary_worker_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Describes how the autoscaler will operate for secondary workers.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_instances": {
							Type:     schema.TypeInt,
							Optional: true,
							Description: `Maximum number of instances for this group. Note that by default, clusters will not use
secondary workers. Required for secondary workers if the minimum secondary instances is set.
Bounds: [minInstances, ). Defaults to 0.`,
							Default:      0,
							AtLeastOneOf: []string{"secondary_worker_config.0.min_instances", "secondary_worker_config.0.max_instances", "secondary_worker_config.0.weight"},
						},
						"min_instances": {
							Type:         schema.TypeInt,
							Optional:     true,
							Description:  `Minimum number of instances for this group. Bounds: [0, maxInstances]. Defaults to 0.`,
							Default:      0,
							AtLeastOneOf: []string{"secondary_worker_config.0.min_instances", "secondary_worker_config.0.max_instances", "secondary_worker_config.0.weight"},
						},
						"weight": {
							Type:     schema.TypeInt,
							Optional: true,
							Description: `Weight for the instance group, which is used to determine the fraction of total workers
in the cluster from this instance group. For example, if primary workers have weight 2,
and secondary workers have weight 1, the cluster will have approximately 2 primary workers
for each secondary worker.

The cluster may not reach the specified balance if constrained by min/max bounds or other
autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
primary workers will be added. The cluster can also be out of balance when created.

If weight is not set on any instance group, the cluster will default to equal weight for
all groups: the cluster will attempt to maintain an equal number of workers in each group
within the configured size bounds for each group. If weight is set for one group only,
the cluster will default to zero weight on the unset group. For example if weight is set
only on primary workers, the cluster will use primary workers only and no secondary workers.`,
							Default:      1,
							AtLeastOneOf: []string{"secondary_worker_config.0.min_instances", "secondary_worker_config.0.max_instances", "secondary_worker_config.0.weight"},
						},
					},
				},
			},
			"worker_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Describes how the autoscaler will operate for primary workers.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_instances": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: `Maximum number of instances for this group.`,
						},
						"min_instances": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: `Minimum number of instances for this group. Bounds: [2, maxInstances]. Defaults to 2.`,
							Default:     2,
						},
						"weight": {
							Type:     schema.TypeInt,
							Optional: true,
							Description: `Weight for the instance group, which is used to determine the fraction of total workers
in the cluster from this instance group. For example, if primary workers have weight 2,
and secondary workers have weight 1, the cluster will have approximately 2 primary workers
for each secondary worker.

The cluster may not reach the specified balance if constrained by min/max bounds or other
autoscaling settings. For example, if maxInstances for secondary workers is 0, then only
primary workers will be added. The cluster can also be out of balance when created.

If weight is not set on any instance group, the cluster will default to equal weight for
all groups: the cluster will attempt to maintain an equal number of workers in each group
within the configured size bounds for each group. If weight is set for one group only,
the cluster will default to zero weight on the unset group. For example if weight is set
only on primary workers, the cluster will use primary workers only and no secondary workers.`,
							Default: 1,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The "resource name" of the autoscaling policy.`,
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

func resourceDataprocAutoscalingPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	idProp, err := expandDataprocAutoscalingPolicyPolicyId(d.Get("policy_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("policy_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	workerConfigProp, err := expandDataprocAutoscalingPolicyWorkerConfig(d.Get("worker_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("worker_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(workerConfigProp)) && (ok || !reflect.DeepEqual(v, workerConfigProp)) {
		obj["workerConfig"] = workerConfigProp
	}
	secondaryWorkerConfigProp, err := expandDataprocAutoscalingPolicySecondaryWorkerConfig(d.Get("secondary_worker_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("secondary_worker_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(secondaryWorkerConfigProp)) && (ok || !reflect.DeepEqual(v, secondaryWorkerConfigProp)) {
		obj["secondaryWorkerConfig"] = secondaryWorkerConfigProp
	}
	basicAlgorithmProp, err := expandDataprocAutoscalingPolicyBasicAlgorithm(d.Get("basic_algorithm"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("basic_algorithm"); !tpgresource.IsEmptyValue(reflect.ValueOf(basicAlgorithmProp)) && (ok || !reflect.DeepEqual(v, basicAlgorithmProp)) {
		obj["basicAlgorithm"] = basicAlgorithmProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocBasePath}}projects/{{project}}/locations/{{location}}/autoscalingPolicies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AutoscalingPolicy: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AutoscalingPolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating AutoscalingPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating AutoscalingPolicy %q: %#v", d.Id(), res)

	return resourceDataprocAutoscalingPolicyRead(d, meta)
}

func resourceDataprocAutoscalingPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocBasePath}}projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AutoscalingPolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DataprocAutoscalingPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading AutoscalingPolicy: %s", err)
	}

	if err := d.Set("policy_id", flattenDataprocAutoscalingPolicyPolicyId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading AutoscalingPolicy: %s", err)
	}
	if err := d.Set("name", flattenDataprocAutoscalingPolicyName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading AutoscalingPolicy: %s", err)
	}
	if err := d.Set("worker_config", flattenDataprocAutoscalingPolicyWorkerConfig(res["workerConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading AutoscalingPolicy: %s", err)
	}
	if err := d.Set("secondary_worker_config", flattenDataprocAutoscalingPolicySecondaryWorkerConfig(res["secondaryWorkerConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading AutoscalingPolicy: %s", err)
	}
	if err := d.Set("basic_algorithm", flattenDataprocAutoscalingPolicyBasicAlgorithm(res["basicAlgorithm"], d, config)); err != nil {
		return fmt.Errorf("Error reading AutoscalingPolicy: %s", err)
	}

	return nil
}

func resourceDataprocAutoscalingPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AutoscalingPolicy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	idProp, err := expandDataprocAutoscalingPolicyPolicyId(d.Get("policy_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("policy_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	workerConfigProp, err := expandDataprocAutoscalingPolicyWorkerConfig(d.Get("worker_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("worker_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, workerConfigProp)) {
		obj["workerConfig"] = workerConfigProp
	}
	secondaryWorkerConfigProp, err := expandDataprocAutoscalingPolicySecondaryWorkerConfig(d.Get("secondary_worker_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("secondary_worker_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, secondaryWorkerConfigProp)) {
		obj["secondaryWorkerConfig"] = secondaryWorkerConfigProp
	}
	basicAlgorithmProp, err := expandDataprocAutoscalingPolicyBasicAlgorithm(d.Get("basic_algorithm"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("basic_algorithm"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, basicAlgorithmProp)) {
		obj["basicAlgorithm"] = basicAlgorithmProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocBasePath}}projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AutoscalingPolicy %q: %#v", d.Id(), obj)
	headers := make(http.Header)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating AutoscalingPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating AutoscalingPolicy %q: %#v", d.Id(), res)
	}

	return resourceDataprocAutoscalingPolicyRead(d, meta)
}

func resourceDataprocAutoscalingPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AutoscalingPolicy: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DataprocBasePath}}projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting AutoscalingPolicy %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "AutoscalingPolicy")
	}

	log.Printf("[DEBUG] Finished deleting AutoscalingPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceDataprocAutoscalingPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/autoscalingPolicies/(?P<policy_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<policy_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<policy_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDataprocAutoscalingPolicyPolicyId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocAutoscalingPolicyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocAutoscalingPolicyWorkerConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["min_instances"] =
		flattenDataprocAutoscalingPolicyWorkerConfigMinInstances(original["minInstances"], d, config)
	transformed["max_instances"] =
		flattenDataprocAutoscalingPolicyWorkerConfigMaxInstances(original["maxInstances"], d, config)
	transformed["weight"] =
		flattenDataprocAutoscalingPolicyWorkerConfigWeight(original["weight"], d, config)
	return []interface{}{transformed}
}
func flattenDataprocAutoscalingPolicyWorkerConfigMinInstances(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenDataprocAutoscalingPolicyWorkerConfigMaxInstances(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenDataprocAutoscalingPolicyWorkerConfigWeight(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenDataprocAutoscalingPolicySecondaryWorkerConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["min_instances"] =
		flattenDataprocAutoscalingPolicySecondaryWorkerConfigMinInstances(original["minInstances"], d, config)
	transformed["max_instances"] =
		flattenDataprocAutoscalingPolicySecondaryWorkerConfigMaxInstances(original["maxInstances"], d, config)
	transformed["weight"] =
		flattenDataprocAutoscalingPolicySecondaryWorkerConfigWeight(original["weight"], d, config)
	return []interface{}{transformed}
}
func flattenDataprocAutoscalingPolicySecondaryWorkerConfigMinInstances(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenDataprocAutoscalingPolicySecondaryWorkerConfigMaxInstances(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenDataprocAutoscalingPolicySecondaryWorkerConfigWeight(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenDataprocAutoscalingPolicyBasicAlgorithm(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["cooldown_period"] =
		flattenDataprocAutoscalingPolicyBasicAlgorithmCooldownPeriod(original["cooldownPeriod"], d, config)
	transformed["yarn_config"] =
		flattenDataprocAutoscalingPolicyBasicAlgorithmYarnConfig(original["yarnConfig"], d, config)
	return []interface{}{transformed}
}
func flattenDataprocAutoscalingPolicyBasicAlgorithmCooldownPeriod(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocAutoscalingPolicyBasicAlgorithmYarnConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["graceful_decommission_timeout"] =
		flattenDataprocAutoscalingPolicyBasicAlgorithmYarnConfigGracefulDecommissionTimeout(original["gracefulDecommissionTimeout"], d, config)
	transformed["scale_up_factor"] =
		flattenDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpFactor(original["scaleUpFactor"], d, config)
	transformed["scale_down_factor"] =
		flattenDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownFactor(original["scaleDownFactor"], d, config)
	transformed["scale_up_min_worker_fraction"] =
		flattenDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpMinWorkerFraction(original["scaleUpMinWorkerFraction"], d, config)
	transformed["scale_down_min_worker_fraction"] =
		flattenDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownMinWorkerFraction(original["scaleDownMinWorkerFraction"], d, config)
	return []interface{}{transformed}
}
func flattenDataprocAutoscalingPolicyBasicAlgorithmYarnConfigGracefulDecommissionTimeout(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpFactor(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownFactor(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpMinWorkerFraction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownMinWorkerFraction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDataprocAutoscalingPolicyPolicyId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyWorkerConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinInstances, err := expandDataprocAutoscalingPolicyWorkerConfigMinInstances(original["min_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinInstances); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minInstances"] = transformedMinInstances
	}

	transformedMaxInstances, err := expandDataprocAutoscalingPolicyWorkerConfigMaxInstances(original["max_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxInstances); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxInstances"] = transformedMaxInstances
	}

	transformedWeight, err := expandDataprocAutoscalingPolicyWorkerConfigWeight(original["weight"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeight); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["weight"] = transformedWeight
	}

	return transformed, nil
}

func expandDataprocAutoscalingPolicyWorkerConfigMinInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyWorkerConfigMaxInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyWorkerConfigWeight(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicySecondaryWorkerConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinInstances, err := expandDataprocAutoscalingPolicySecondaryWorkerConfigMinInstances(original["min_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinInstances); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minInstances"] = transformedMinInstances
	}

	transformedMaxInstances, err := expandDataprocAutoscalingPolicySecondaryWorkerConfigMaxInstances(original["max_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxInstances); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxInstances"] = transformedMaxInstances
	}

	transformedWeight, err := expandDataprocAutoscalingPolicySecondaryWorkerConfigWeight(original["weight"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeight); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["weight"] = transformedWeight
	}

	return transformed, nil
}

func expandDataprocAutoscalingPolicySecondaryWorkerConfigMinInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicySecondaryWorkerConfigMaxInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicySecondaryWorkerConfigWeight(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCooldownPeriod, err := expandDataprocAutoscalingPolicyBasicAlgorithmCooldownPeriod(original["cooldown_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCooldownPeriod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cooldownPeriod"] = transformedCooldownPeriod
	}

	transformedYarnConfig, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfig(original["yarn_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedYarnConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["yarnConfig"] = transformedYarnConfig
	}

	return transformed, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmCooldownPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGracefulDecommissionTimeout, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigGracefulDecommissionTimeout(original["graceful_decommission_timeout"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGracefulDecommissionTimeout); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gracefulDecommissionTimeout"] = transformedGracefulDecommissionTimeout
	}

	transformedScaleUpFactor, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpFactor(original["scale_up_factor"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleUpFactor); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scaleUpFactor"] = transformedScaleUpFactor
	}

	transformedScaleDownFactor, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownFactor(original["scale_down_factor"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleDownFactor); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scaleDownFactor"] = transformedScaleDownFactor
	}

	transformedScaleUpMinWorkerFraction, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpMinWorkerFraction(original["scale_up_min_worker_fraction"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleUpMinWorkerFraction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scaleUpMinWorkerFraction"] = transformedScaleUpMinWorkerFraction
	}

	transformedScaleDownMinWorkerFraction, err := expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownMinWorkerFraction(original["scale_down_min_worker_fraction"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScaleDownMinWorkerFraction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scaleDownMinWorkerFraction"] = transformedScaleDownMinWorkerFraction
	}

	return transformed, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigGracefulDecommissionTimeout(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpFactor(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownFactor(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleUpMinWorkerFraction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataprocAutoscalingPolicyBasicAlgorithmYarnConfigScaleDownMinWorkerFraction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
