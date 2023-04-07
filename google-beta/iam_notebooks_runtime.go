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

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var NotebooksRuntimeIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"location": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"runtime_name": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type NotebooksRuntimeIamUpdater struct {
	project     string
	location    string
	runtimeName string
	d           TerraformResourceData
	Config      *Config
}

func NotebooksRuntimeIamUpdaterProducer(d TerraformResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	location, _ := getLocation(d, config)
	if location != "" {
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	values["location"] = location
	if v, ok := d.GetOk("runtime_name"); ok {
		values["runtime_name"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/runtimes/(?P<runtime_name>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<runtime_name>[^/]+)", "(?P<location>[^/]+)/(?P<runtime_name>[^/]+)", "(?P<runtime_name>[^/]+)"}, d, config, d.Get("runtime_name").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &NotebooksRuntimeIamUpdater{
		project:     values["project"],
		location:    values["location"],
		runtimeName: values["runtime_name"],
		d:           d,
		Config:      config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("runtime_name", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting runtime_name: %s", err)
	}

	return u, nil
}

func NotebooksRuntimeIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := getLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/runtimes/(?P<runtime_name>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<runtime_name>[^/]+)", "(?P<location>[^/]+)/(?P<runtime_name>[^/]+)", "(?P<runtime_name>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &NotebooksRuntimeIamUpdater{
		project:     values["project"],
		location:    values["location"],
		runtimeName: values["runtime_name"],
		d:           d,
		Config:      config,
	}
	if err := d.Set("runtime_name", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting runtime_name: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *NotebooksRuntimeIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyRuntimeUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := SendRequest(u.Config, "GET", project, url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *NotebooksRuntimeIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyRuntimeUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = SendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *NotebooksRuntimeIamUpdater) qualifyRuntimeUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{NotebooksBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/runtimes/%s", u.project, u.location, u.runtimeName), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *NotebooksRuntimeIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/runtimes/%s", u.project, u.location, u.runtimeName)
}

func (u *NotebooksRuntimeIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-notebooks-runtime-%s", u.GetResourceId())
}

func (u *NotebooksRuntimeIamUpdater) DescribeResource() string {
	return fmt.Sprintf("notebooks runtime %q", u.GetResourceId())
}
