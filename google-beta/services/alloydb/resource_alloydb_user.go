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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/alloydb/User.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package alloydb

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

func ResourceAlloydbUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlloydbUserCreate,
		Read:   resourceAlloydbUserRead,
		Update: resourceAlloydbUserUpdate,
		Delete: resourceAlloydbUserDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAlloydbUserImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `Identifies the alloydb cluster. Must be in the format
'projects/{project}/locations/{location}/clusters/{cluster_id}'`,
			},
			"user_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The database role name of the user.`,
			},
			"user_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ALLOYDB_BUILT_IN", "ALLOYDB_IAM_USER"}),
				Description:  `The type of this user. Possible values: ["ALLOYDB_BUILT_IN", "ALLOYDB_IAM_USER"]`,
			},
			"database_roles": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `List of database roles this database user has.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Password for this database user.`,
				Sensitive:   true,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Name of the resource in the form of projects/{project}/locations/{location}/clusters/{cluster}/users/{user}.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAlloydbUserCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	passwordProp, err := expandAlloydbUserPassword(d.Get("password"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("password"); !tpgresource.IsEmptyValue(reflect.ValueOf(passwordProp)) && (ok || !reflect.DeepEqual(v, passwordProp)) {
		obj["password"] = passwordProp
	}
	databaseRolesProp, err := expandAlloydbUserDatabaseRoles(d.Get("database_roles"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("database_roles"); !tpgresource.IsEmptyValue(reflect.ValueOf(databaseRolesProp)) && (ok || !reflect.DeepEqual(v, databaseRolesProp)) {
		obj["databaseRoles"] = databaseRolesProp
	}
	userTypeProp, err := expandAlloydbUserUserType(d.Get("user_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(userTypeProp)) && (ok || !reflect.DeepEqual(v, userTypeProp)) {
		obj["userType"] = userTypeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/users?userId={{user_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new User: %#v", obj)
	billingProject := ""

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
		return fmt.Errorf("Error creating User: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{cluster}}/users/{{user_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating User %q: %#v", d.Id(), res)

	return resourceAlloydbUserRead(d, meta)
}

func resourceAlloydbUserRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/users/{{user_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AlloydbUser %q", d.Id()))
	}

	if err := d.Set("name", flattenAlloydbUserName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading User: %s", err)
	}
	if err := d.Set("database_roles", flattenAlloydbUserDatabaseRoles(res["databaseRoles"], d, config)); err != nil {
		return fmt.Errorf("Error reading User: %s", err)
	}
	if err := d.Set("user_type", flattenAlloydbUserUserType(res["userType"], d, config)); err != nil {
		return fmt.Errorf("Error reading User: %s", err)
	}

	return nil
}

func resourceAlloydbUserUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	passwordProp, err := expandAlloydbUserPassword(d.Get("password"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("password"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, passwordProp)) {
		obj["password"] = passwordProp
	}
	databaseRolesProp, err := expandAlloydbUserDatabaseRoles(d.Get("database_roles"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("database_roles"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, databaseRolesProp)) {
		obj["databaseRoles"] = databaseRolesProp
	}
	userTypeProp, err := expandAlloydbUserUserType(d.Get("user_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, userTypeProp)) {
		obj["userType"] = userTypeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/users?userId={{user_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating User %q: %#v", d.Id(), obj)
	headers := make(http.Header)

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
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating User %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating User %q: %#v", d.Id(), res)
	}

	return resourceAlloydbUserRead(d, meta)
}

func resourceAlloydbUserDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}{{cluster}}/users/{{user_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting User %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "User")
	}

	log.Printf("[DEBUG] Finished deleting User %q: %#v", d.Id(), res)
	return nil
}

func resourceAlloydbUserImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{
		"(?P<cluster>.+)/users/(?P<user_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{cluster}}/users/{{user_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAlloydbUserName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbUserDatabaseRoles(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbUserUserType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandAlloydbUserPassword(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbUserDatabaseRoles(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbUserUserType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
