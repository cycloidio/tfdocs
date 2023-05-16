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
			Type:             "azuredevops_build_definition",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing Build Definition.`,
			Description: `

Use this data source to access information about an existing Build Definition.

`,
			Keywords:  []string{},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Build Definition.`,
				},
				resource.Attribute{
					Name:        "agent_pool_name",
					Description: `The agent pool that should execute the build.`,
				},
				resource.Attribute{
					Name:        "ci_trigger",
					Description: `A ` + "`" + `ci_trigger` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "pull_request_trigger",
					Description: `A ` + "`" + `pull_request_trigger` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `A ` + "`" + `repository` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `The revision of the build definition.`,
				},
				resource.Attribute{
					Name:        "schedules",
					Description: `A ` + "`" + `schedules` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "variable",
					Description: `A ` + "`" + `variable` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "variable_groups",
					Description: `A list of variable group IDs. --- A ` + "`" + `branch_filter` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `A ` + "`" + `exclude` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "include",
					Description: `A ` + "`" + `include` + "`" + ` block as defined below. --- A ` + "`" + `ci_trigger` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "override",
					Description: `A ` + "`" + `override` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "use_yaml",
					Description: `Use the azure-pipeline file for the build configuration. --- A ` + "`" + `ci_trigger` + "`" + ` ` + "`" + `override` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "batch",
					Description: `If batch is true, when a pipeline is running, the system waits until the run is completed, then starts another run with all changes that have not yet been built.`,
				},
				resource.Attribute{
					Name:        "branch_filter",
					Description: `The branches to include and exclude from the trigger.`,
				},
				resource.Attribute{
					Name:        "path_filter",
					Description: `Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.`,
				},
				resource.Attribute{
					Name:        "max_concurrent_builds_per_branch",
					Description: `The number of max builds per branch.`,
				},
				resource.Attribute{
					Name:        "polling_interval",
					Description: `How often the external repository is polled.`,
				},
				resource.Attribute{
					Name:        "polling_job_id",
					Description: `This is the ID of the polling job that polls the external repository. Once the build definition is saved/updated, this value is set.`,
				},
				resource.Attribute{
					Name:        "include",
					Description: `(Optional) List of branch patterns to include.`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `(Optional) List of branch patterns to exclude.`,
				},
				resource.Attribute{
					Name:        "include",
					Description: `(Optional) List of path patterns to include.`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `(Optional) List of path patterns to exclude. --- A ` + "`" + `pull_request_trigger` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "comment_required",
					Description: `Is a comment required on the PR?`,
				},
				resource.Attribute{
					Name:        "forks",
					Description: `A ` + "`" + `forks` + "`" + ` block as defined above.`,
				},
				resource.Attribute{
					Name:        "initial_branch",
					Description: `When use_yaml is true set this to the name of the branch that the azure-pipelines.yml exists on.`,
				},
				resource.Attribute{
					Name:        "override",
					Description: `A ` + "`" + `override` + "`" + ` block as defined below.`,
				},
				resource.Attribute{
					Name:        "use_yaml",
					Description: `Use the azure-pipeline file for the build configuration. --- A ` + "`" + `forks` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Build pull requests from forks of this repository.`,
				},
				resource.Attribute{
					Name:        "share_secrets",
					Description: `Make secrets available to builds of forks. --- A ` + "`" + `pull_request_trigger` + "`" + ` ` + "`" + `override` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "branch_filter",
					Description: `The branches to include and exclude from the trigger.`,
				},
				resource.Attribute{
					Name:        "path_filter",
					Description: `The file paths to include or exclude.`,
				},
				resource.Attribute{
					Name:        "include",
					Description: `(Optional) List of branch patterns to include.`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `(Optional) List of branch patterns to exclude.`,
				},
				resource.Attribute{
					Name:        "include",
					Description: `(Optional) List of path patterns to include.`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `(Optional) List of path patterns to exclude. --- A ` + "`" + `repository` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "branch_name",
					Description: `The branch name for which builds are triggered.`,
				},
				resource.Attribute{
					Name:        "github_enterprise_url",
					Description: `The Github Enterprise URL.`,
				},
				resource.Attribute{
					Name:        "repo_id",
					Description: `The id of the repository.`,
				},
				resource.Attribute{
					Name:        "repo_type",
					Description: `The repository type.`,
				},
				resource.Attribute{
					Name:        "report_build_status",
					Description: `Report build status.`,
				},
				resource.Attribute{
					Name:        "service_connection_id",
					Description: `The service connection ID.`,
				},
				resource.Attribute{
					Name:        "yml_path",
					Description: `The path of the Yaml file describing the build definition. --- A ` + "`" + `schedules` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "branch_filter",
					Description: `A ` + "`" + `branch_filter` + "`" + ` block as defined above.`,
				},
				resource.Attribute{
					Name:        "days_to_build",
					Description: `A list of days to build on.`,
				},
				resource.Attribute{
					Name:        "schedule_job_id",
					Description: `The ID of the schedule job.`,
				},
				resource.Attribute{
					Name:        "schedule_only_with_changes",
					Description: `Schedule builds if the source or pipeline has changed.`,
				},
				resource.Attribute{
					Name:        "start_hours",
					Description: `Build start hour.`,
				},
				resource.Attribute{
					Name:        "start_minutes",
					Description: `Build start minute.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `Build time zone. --- A ` + "`" + `variable` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "allow_override",
					Description: `` + "`" + `true` + "`" + ` if the variable can be overridden.`,
				},
				resource.Attribute{
					Name:        "is_secret",
					Description: `` + "`" + `true` + "`" + ` if the variable is a secret.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the variable.`,
				},
				resource.Attribute{
					Name:        "secret_value",
					Description: `The secret value of the variable.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the variable.`,
				},
			},
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
			Type:             "azuredevops_serviceendpoint_azurerm",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing AzureRM Service Endpoint.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "authorization",
					Description: `Specifies the Authorization Scheme Map.`,
				},
				resource.Attribute{
					Name:        "azurerm_management_group_id",
					Description: `Specified the Management Group ID of the Service Endpoint is target, if available.`,
				},
				resource.Attribute{
					Name:        "azurerm_management_group_name",
					Description: `Specified the Management Group Name of the Service Endpoint target, if available.`,
				},
				resource.Attribute{
					Name:        "azurerm_subscription_id",
					Description: `Specifies the Subscription ID of the Service Endpoint target, if available.`,
				},
				resource.Attribute{
					Name:        "azurerm_subscription_name",
					Description: `Specifies the Subscription Name of the Service Endpoint target, if available.`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `Specifies the Resource Group of the Service Endpoint target, if available.`,
				},
				resource.Attribute{
					Name:        "azurerm_spn_tenantid",
					Description: `Specifies the Tenant ID of the Azure targets.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Specifies the description of the Service Endpoint.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_github",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an existing GitHub Service Endpoint.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "authorization",
					Description: `Specifies the Authorization Scheme Map.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Specifies the description of the Service Endpoint.`,
				},
			},
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

		"azuredevops_agent_pool":              0,
		"azuredevops_agent_pools":             1,
		"azuredevops_agent_queue":             2,
		"azuredevops_area":                    3,
		"azuredevops_build_definition":        4,
		"azuredevops_client_config":           5,
		"azuredevops_team":                    6,
		"azuredevops_teams":                   7,
		"azuredevops_git_repositories":        8,
		"azuredevops_git_repository":          9,
		"azuredevops_group":                   10,
		"azuredevops_groups":                  11,
		"azuredevops_iteration":               12,
		"azuredevops_project":                 13,
		"azuredevops_projects":                14,
		"azuredevops_serviceendpoint_azurerm": 15,
		"azuredevops_serviceendpoint_github":  16,
		"azuredevops_users":                   17,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
