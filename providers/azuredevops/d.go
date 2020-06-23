package azuredevops

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_agent_pool",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about an existing Agent Pool within Azure DevOps.`,
			Description: `

Use this data source to access information about an existing Agent Pool within Azure DevOps.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_agent_pools",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about existing Agent Pools within Azure DevOps.`,
			Description: `

Use this data source to access information about existing Agent Pools within Azure DevOps.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_client_config",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about the Azure DevOps organization configured for the provider.`,
			Description: `

Use this data source to access information about the Azure DevOps organization configured for the provider.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_git_repositories",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about an existing Projects within Azure DevOps.`,
			Description: `

Use this data source to access information about an existing Git Repositories within Azure DevOps.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_group",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about an existing Group within Azure DevOps.`,
			Description: `
Use this data source to access information about an existing Group within Azure DevOps

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The Project Id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Group Name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID for this resource is the group descriptor. See below.`,
				},
				resource.Attribute{
					Name:        "descriptor",
					Description: `The Descriptor is the primary way to reference the graph subject. This field will uniquely identify the same graph subject across both Accounts and Organizations. ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID for this resource is the group descriptor. See below.`,
				},
				resource.Attribute{
					Name:        "descriptor",
					Description: `The Descriptor is the primary way to reference the graph subject. This field will uniquely identify the same graph subject across both Accounts and Organizations. ## Relevant Links`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_project",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about an existing Project within Azure DevOps.`,
			Description: `

Use this data source to access information about an existing Project within Azure DevOps.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_projects",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about a existing Projects within Azure DevOps.`,
			Description: `

Use this data source to access information about existing Projects within Azure DevOps.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_users",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about an existing users within Azure DevOps.`,
			Description: `

Use this data source to access information about an existing users within Azure DevOps.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"azuredevops_agent_pool":       0,
		"azuredevops_agent_pools":      1,
		"azuredevops_client_config":    2,
		"azuredevops_git_repositories": 3,
		"azuredevops_group":            4,
		"azuredevops_project":          5,
		"azuredevops_projects":         6,
		"azuredevops_users":            7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
