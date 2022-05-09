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
			Type:             "azuredevops_agent_queue",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about an existing Agent Queue within Azure DevOps.`,
			Description: `

Use this data source to access information about an existing Agent Queue within Azure DevOps.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_area",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about an existing Area (Component) within Azure DevOps.`,
			Description: `

Use this data source to access information about an existing Area (Component) within Azure DevOps.

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
			Type:             "azuredevops_team",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about an existing Team in a Project within Azure DevOps.`,
			Description: `

Use this data source to access information about an existing Team in a Project within Azure DevOps.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_teams",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about existing Teams in a Project or globally within an Azure DevOps organization`,
			Description: `

Use this data source to access information about existing Teams in a Project or globally within an Azure DevOps organization

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_git_repositories",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about existing Git Repositories within Azure DevOps.`,
			Description: `

Use this data source to access information about **multiple** existing Git Repositories within Azure DevOps.
To read informations about a **single** Git Repository use the data source [` + "`" + `azuredevops_git_repository` + "`" + `](data_git_repository.html)

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_git_repository",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about an existing Git Repository within Azure DevOps.`,
			Description: `

Use this data source to access information about a **single** (existing) Git Repository within Azure DevOps.
To read information about **multiple** Git Repositories use the data source [` + "`" + `azuredevops_git_repositories` + "`" + `](data_git_repositories.html)

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
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about existing Groups within Azure DevOps`,
			Description: `

Use this data source to access information about existing Groups within Azure DevOps

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_iteration",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to access information about an existing Iteration (Sprint) within Azure DevOps.`,
			Description: `

Use this data source to access information about an existing Iteration (Sprint) within Azure DevOps.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
		"azuredevops_agent_queue":      2,
		"azuredevops_area":             3,
		"azuredevops_client_config":    4,
		"azuredevops_team":             5,
		"azuredevops_teams":            6,
		"azuredevops_git_repositories": 7,
		"azuredevops_git_repository":   8,
		"azuredevops_group":            9,
		"azuredevops_groups":           10,
		"azuredevops_iteration":        11,
		"azuredevops_project":          12,
		"azuredevops_projects":         13,
		"azuredevops_users":            14,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
