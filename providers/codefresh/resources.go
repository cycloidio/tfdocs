package codefresh

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "codefresh_account",
			Category:         "Resources",
			ShortDescription: `By creating different accounts for different teams within the same company a customer can achieve complete segregation of assets between the teams.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_account_admins",
			Category:         "Resources",
			ShortDescription: `Use this resource to set a list of admins for any account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_api_key",
			Category:         "Resources",
			ShortDescription: `Manages an API Key tied to an Account and a User.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_context",
			Category:         "Resources",
			ShortDescription: `A Context is an authentication/configuration that is used by the Codefresh system and engine.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_idp_accounts",
			Category:         "Resources",
			ShortDescription: `This resource adds the list of provided account IDs to the IDP.Because of the current Codefresh API limitation it's impossible to remove account from IDP, thus deletion is not supported.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_permission",
			Category:         "Resources",
			ShortDescription: `Permission are used to setup access control and allow to define which teams have access to which clusters and pipelines based on tags.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_pipeline",
			Category:         "Resources",
			ShortDescription: `The central component of the Codefresh Platform. Pipelines are workflows that contain individual steps. Each step is responsible for a specific action in the process.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_pipeline_cron_trigger",
			Category:         "Resources",
			ShortDescription: `This resource is used to create cron-based triggers for pipeilnes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_project",
			Category:         "Resources",
			ShortDescription: `The top-level concept in Codefresh. You can create projects to group pipelines that are related. In most cases a single project will be a single application (that itself contains many micro-services). You are free to use projects as you see fit. For example, you could create a project for a specific Kubernetes cluster or a specific team/department.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_registry",
			Category:         "Resources",
			ShortDescription: `Registry is the configuration that Codefresh uses to push/pull container images.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_step_types",
			Category:         "Resources",
			ShortDescription: `This resource allows to create your own typed step and manage all of its published versions. The resource allows to handle the life-cycle of the version by allowing specifying multiple blocks 'version' where the user provides a version number and the yaml file representing the plugin.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_team",
			Category:         "Resources",
			ShortDescription: `Teams are groups of users that are used to enforce access control.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "codefresh_user",
			Category:         "Resources",
			ShortDescription: `This resource is used to manage a Codefresh user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"codefresh_account":               0,
		"codefresh_account_admins":        1,
		"codefresh_api_key":               2,
		"codefresh_context":               3,
		"codefresh_idp_accounts":          4,
		"codefresh_permission":            5,
		"codefresh_pipeline":              6,
		"codefresh_pipeline_cron_trigger": 7,
		"codefresh_project":               8,
		"codefresh_registry":              9,
		"codefresh_step_types":            10,
		"codefresh_team":                  11,
		"codefresh_user":                  12,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
