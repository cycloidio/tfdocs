package azuredevops

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_agent_pool",
			Category:         "Resources",
			ShortDescription: `Manages an agent pool within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"agent",
				"pool",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_agent_queue",
			Category:         "Resources",
			ShortDescription: `Manages an agent queue within Azure DevOps project.`,
			Description:      ``,
			Keywords: []string{
				"agent",
				"queue",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_area_permissions",
			Category:         "Resources",
			ShortDescription: `Manages permissions for a AzureDevOps Area (Component)`,
			Description:      ``,
			Keywords: []string{
				"area",
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project to assign the permissions.`,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `(Required) The`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) the permissions to assign. The following permissions are available.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The name of the branch to assign the permissions.`,
				},
				resource.Attribute{
					Name:        "replace",
					Description: `(Optional) Replace (` + "`" + `true` + "`" + `) or merge (` + "`" + `false` + "`" + `) the permissions. Default: ` + "`" + `true` + "`" + `. | Permission | Description | |--------------------|--------------------------------| | GENERIC_READ | View permissions for this node | | GENERIC_WRITE | Edit this node | | CREATE_CHILDREN | Create child nodes | | DELETE | Delete this node | | WORK_ITEM_READ | View work items in this node | | WORK_ITEM_WRITE | Edit work items in this node | | MANAGE_TEST_PLANS | Manage test plans | | MANAGE_TEST_SUITES | Manage test suites | ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_branch_policy_auto_reviewers",
			Category:         "Resources",
			ShortDescription: `Manages required reviewer policy branch policy within Azure DevOps project.`,
			Description:      ``,
			Keywords: []string{
				"branch",
				"policy",
				"auto",
				"reviewers",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_branch_policy_build_validation",
			Category:         "Resources",
			ShortDescription: `Manages a build validation branch policy within Azure DevOps project.`,
			Description:      ``,
			Keywords: []string{
				"branch",
				"policy",
				"build",
				"validation",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_branch_policy_comment_resolution",
			Category:         "Resources",
			ShortDescription: `Configure a comment resolution policy for your branch within Azure DevOps project.`,
			Description:      ``,
			Keywords: []string{
				"branch",
				"policy",
				"comment",
				"resolution",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_branch_policy_merge_types",
			Category:         "Resources",
			ShortDescription: `Enforces the merge types allowed on a branch.`,
			Description:      ``,
			Keywords: []string{
				"branch",
				"policy",
				"merge",
				"types",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_branch_policy_min_reviewers",
			Category:         "Resources",
			ShortDescription: `Manages a minimum reviewer branch policy within Azure DevOps project.`,
			Description:      ``,
			Keywords: []string{
				"branch",
				"policy",
				"min",
				"reviewers",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_branch_policy_status_check",
			Category:         "Resources",
			ShortDescription: `Manages status check branch policy within Azure DevOps project.`,
			Description:      ``,
			Keywords: []string{
				"branch",
				"policy",
				"status",
				"check",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_branch_policy_work_item_linking",
			Category:         "Resources",
			ShortDescription: `Require associations between branches and a work item within Azure DevOps.`,
			Description:      ``,
			Keywords: []string{
				"branch",
				"policy",
				"work",
				"item",
				"linking",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_build_definition",
			Category:         "Resources",
			ShortDescription: `Manages a Build Definition within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"build",
				"definition",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_build_definition_permissions",
			Category:         "Resources",
			ShortDescription: `Manages permissions for a AzureDevOps Build Definition`,
			Description:      ``,
			Keywords: []string{
				"build",
				"definition",
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project to assign the permissions.`,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `(Required) The`,
				},
				resource.Attribute{
					Name:        "build_definition_id",
					Description: `(Required) The id of the build definition to assign the permissions.`,
				},
				resource.Attribute{
					Name:        "replace",
					Description: `(Optional) Replace (` + "`" + `true` + "`" + `) or merge (` + "`" + `false` + "`" + `) the permissions. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) the permissions to assign. The following permissions are available. | Permission | Description | |--------------------------------|---------------------------------------| | ViewBuilds | View builds | | EditBuildQuality | Edit build quality | | RetainIndefinitely | Retain indefinitely | | DeleteBuilds | Delete builds | | ManageBuildQualities | Manage build qualities | | DestroyBuilds | Destroy builds | | UpdateBuildInformation | Update build information | | QueueBuilds | Queue builds | | ManageBuildQueue | Manage build queue | | StopBuilds | Stop builds | | ViewBuildDefinition | View build pipeline | | EditBuildDefinition | Edit build pipeline | | DeleteBuildDefinition | Delete build pipeline | | OverrideBuildCheckInValidation | Override check-in validation by build | | AdministerBuildPermissions | Administer build permissions | ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_build_folder",
			Category:         "Resources",
			ShortDescription: `Manages a Build Folder.`,
			Description:      ``,
			Keywords: []string{
				"build",
				"folder",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_build_folder_permissions",
			Category:         "Resources",
			ShortDescription: `Manages permissions for a AzureDevOps Build Folder`,
			Description:      ``,
			Keywords: []string{
				"build",
				"folder",
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project to assign the permissions.`,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `(Required) The`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The folder path to assign the permissions.`,
				},
				resource.Attribute{
					Name:        "replace",
					Description: `(Optional) Replace (` + "`" + `true` + "`" + `) or merge (` + "`" + `false` + "`" + `) the permissions. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) the permissions to assign. The following permissions are available. | Permission | Description | |--------------------------------|---------------------------------------| | ViewBuilds | View builds | | EditBuildQuality | Edit build quality | | RetainIndefinitely | Retain indefinitely | | DeleteBuilds | Delete builds | | ManageBuildQualities | Manage build qualities | | DestroyBuilds | Destroy builds | | UpdateBuildInformation | Update build information | | QueueBuilds | Queue builds | | ManageBuildQueue | Manage build queue | | StopBuilds | Stop builds | | ViewBuildDefinition | View build pipeline | | EditBuildDefinition | Edit build pipeline | | DeleteBuildDefinition | Delete build pipeline | | OverrideBuildCheckInValidation | Override check-in validation by build | | AdministerBuildPermissions | Administer build permissions | ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_check_branch_control",
			Category:         "Resources",
			ShortDescription: `Manages a branch control check.`,
			Description:      ``,
			Keywords: []string{
				"check",
				"branch",
				"control",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project ID.`,
				},
				resource.Attribute{
					Name:        "target_resource_id",
					Description: `(Required) The ID of the resource being protected by the check.`,
				},
				resource.Attribute{
					Name:        "target_resource_type",
					Description: `(Required) The type of resource being protected by the check. Valid values: ` + "`" + `endpoint` + "`" + `, ` + "`" + `environment` + "`" + `, ` + "`" + `queue` + "`" + `, ` + "`" + `repository` + "`" + `, ` + "`" + `securefile` + "`" + `, ` + "`" + `variablegroup` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The name of the branch control check displayed in the web UI.`,
				},
				resource.Attribute{
					Name:        "allowed_branches",
					Description: `(Optional) The branches allowed to use the resource. Specify a comma separated list of allowed branches in ` + "`" + `refs/heads/branch_name` + "`" + ` format. To allow deployments from all branches, specify ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "verify_branch_protection",
					Description: `(Optional) Validate the branches being deployed are protected. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ignore_unknown_protection_status",
					Description: `(Optional) Allow deployment from branches for which protection status could not be obtained. Only relevant when verify_branch_protection is ` + "`" + `true` + "`" + `. Defaults to ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to all arguments above the following attributes are exported: - ` + "`" + `id` + "`" + ` - The ID of the check. ## Relevant Links - [Define approvals and checks](https://learn.microsoft.com/en-us/azure/devops/pipelines/process/approvals?view=azure-devops&tabs=check-pass) ## Import Importing this resource is not supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_check_business_hours",
			Category:         "Resources",
			ShortDescription: `Manages a branch control check.`,
			Description:      ``,
			Keywords: []string{
				"check",
				"business",
				"hours",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project ID.`,
				},
				resource.Attribute{
					Name:        "target_resource_id",
					Description: `(Required) The ID of the resource being protected by the check.`,
				},
				resource.Attribute{
					Name:        "target_resource_type",
					Description: `(Required) The type of resource being protected by the check. Valid values: ` + "`" + `endpoint` + "`" + `, ` + "`" + `environment` + "`" + `, ` + "`" + `queue` + "`" + `, ` + "`" + `repository` + "`" + `, ` + "`" + `securefile` + "`" + `, ` + "`" + `variablegroup` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The name of the business hours check displayed in the web UI.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) The beginning of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Required) The end of the time period that this check will be allowed to pass, specified as 24-hour time with leading zeros.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Required) The time zone this check will be evaluated in. See below for supported values.`,
				},
				resource.Attribute{
					Name:        "monday",
					Description: `(Optional) This check will pass on Mondays. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tuesday",
					Description: `(Optional) This check will pass on Tuesday. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "wednesday",
					Description: `(Optional) This check will pass on Wednesdays. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "thursday",
					Description: `(Optional) This check will pass on Thursdays. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "friday",
					Description: `(Optional) This check will pass on Fridays. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "saturday",
					Description: `(Optional) This check will pass on Saturdays. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sunday",
					Description: `(Optional) This check will pass on Sundays. Defaults to ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported: - ` + "`" + `id` + "`" + ` - The ID of the check. ## Relevant Links - [Define approvals and checks](https://learn.microsoft.com/en-us/azure/devops/pipelines/process/approvals?view=azure-devops&tabs=check-pass) ## Import Importing this resource is not supported. ## Supported Time Zones - AUS Central Standard Time - AUS Eastern Standard Time - Afghanistan Standard Time - Alaskan Standard Time - Aleutian Standard Time - Altai Standard Time - Arab Standard Time - Arabian Standard Time - Arabic Standard Time - Argentina Standard Time - Astrakhan Standard Time - Atlantic Standard Time - Aus Central W. Standard Time - Azerbaijan Standard Time - Azores Standard Time - Bahia Standard Time - Bangladesh Standard Time - Belarus Standard Time - Bougainville Standard Time - Canada Central Standard Time - Cape Verde Standard Time - Caucasus Standard Time - Cen. Australia Standard Time - Central America Standard Time - Central Asia Standard Time - Central Brazilian Standard Time - Central Europe Standard Time - Central European Standard Time - Central Pacific Standard Time - Central Standard Time (Mexico) - Central Standard Time - Chatham Islands Standard Time - China Standard Time - Cuba Standard Time - Dateline Standard Time - E. Africa Standard Time - E. Australia Standard Time - E. Europe Standard Time - E. South America Standard Time - Easter Island Standard Time - Eastern Standard Time (Mexico) - Eastern Standard Time - Egypt Standard Time - Ekaterinburg Standard Time - FLE Standard Time - Fiji Standard Time - GMT Standard Time - GTB Standard Time - Georgian Standard Time - Greenland Standard Time - Greenwich Standard Time - Haiti Standard Time - Hawaiian Standard Time - India Standard Time - Iran Standard Time - Israel Standard Time - Jordan Standard Time - Kaliningrad Standard Time - Kamchatka Standard Time - Korea Standard Time - Libya Standard Time - Line Islands Standard Time - Lord Howe Standard Time - Magadan Standard Time - Magallanes Standard Time - Marquesas Standard Time - Mauritius Standard Time - Mid-Atlantic Standard Time - Middle East Standard Time - Montevideo Standard Time - Morocco Standard Time - Mountain Standard Time (Mexico) - Mountain Standard Time - Myanmar Standard Time - N. Central Asia Standard Time - Namibia Standard Time - Nepal Standard Time - New Zealand Standard Time - Newfoundland Standard Time - Norfolk Standard Time - North Asia East Standard Time - North Asia Standard Time - North Korea Standard Time - Omsk Standard Time - Pacific SA Standard Time - Pacific Standard Time (Mexico) - Pacific Standard Time - Pakistan Standard Time - Paraguay Standard Time - Qyzylorda Standard Time - Romance Standard Time - Russia Time Zone 10 - Russia Time Zone 11 - Russia Time Zone 3 - Russian Standard Time - SA Eastern Standard Time - SA Pacific Standard Time - SA Western Standard Time - SE Asia Standard Time - Saint Pierre Standard Time - Sakhalin Standard Time - Samoa Standard Time - Sao Tome Standard Time - Saratov Standard Time - Singapore Standard Time - South Africa Standard Time - South Sudan Standard Time - Sri Lanka Standard Time - Sudan Standard Time - Syria Standard Time - Taipei Standard Time - Tasmania Standard Time - Tocantins Standard Time - Tokyo Standard Time - Tomsk Standard Time - Tonga Standard Time - Transbaikal Standard Time - Turkey Standard Time - Turks And Caicos Standard Time - US Eastern Standard Time - US Mountain Standard Time - UTC - UTC+12 - UTC+13 - UTC-02 - UTC-08 - UTC-09 - UTC-11 - Ulaanbaatar Standard Time - Venezuela Standard Time - Vladivostok Standard Time - Volgograd Standard Time - W. Australia Standard Time - W. Central Africa Standard Time - W. Europe Standard Time - W. Mongolia Standard Time - West Asia Standard Time - West Bank Standard Time - West Pacific Standard Time - Yakutsk Standard Time - Yukon Standard Time`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_git_permissions",
			Category:         "Resources",
			ShortDescription: `Manages permissions for Git repositories`,
			Description:      ``,
			Keywords: []string{
				"git",
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project to assign the permissions.`,
				},
				resource.Attribute{
					Name:        "repository_id",
					Description: `(Optional) The ID of the GIT repository to assign the permissions`,
				},
				resource.Attribute{
					Name:        "branch_name",
					Description: `(Optional) The name of the branch to assign the permissions. ~>`,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `(Required) The`,
				},
				resource.Attribute{
					Name:        "replace",
					Description: `(Optional) Replace (` + "`" + `true` + "`" + `) or merge (` + "`" + `false` + "`" + `) the permissions. Default: ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) the permissions to assign. The follwing permissions are available | Permissions | Description | |-------------------------|--------------------------------------------------------| | Administer | Administer | | GenericRead | Read | | GenericContribute | Contribute | | ForcePush | Force push (rewrite history, delete branches and tags) | | CreateBranch | Create branch | | CreateTag | Create tag | | ManageNote | Manage notes | | PolicyExempt | Bypass policies when pushing | | CreateRepository | Create repository | | DeleteRepository | Delete repository | | RenameRepository | Rename repository | | EditPolicies | Edit policies | | RemoveOthersLocks | Remove others' locks | | ManagePermissions | Manage permissions | | PullRequestContribute | Contribute to pull requests | | PullRequestBypassPolicy | Bypass policies when completing pull requests | ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_git_repository",
			Category:         "Resources",
			ShortDescription: `Manages a git repository within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"git",
				"repository",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_git_repository_branch",
			Category:         "Resources",
			ShortDescription: `Manages a Git Repository Branch.`,
			Description:      ``,
			Keywords: []string{
				"git",
				"repository",
				"branch",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_git_repository_file",
			Category:         "Resources",
			ShortDescription: `Manage files within an Azure DevOps Git repository.`,
			Description:      ``,
			Keywords: []string{
				"git",
				"repository",
				"file",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_group",
			Category:         "Resources",
			ShortDescription: `Manages a group within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_group_membership",
			Category:         "Resources",
			ShortDescription: `Manages group membership within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"group",
				"membership",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_iteration_permissions",
			Category:         "Resources",
			ShortDescription: `Manages permissions for a AzureDevOps Iteration (Sprint)`,
			Description:      ``,
			Keywords: []string{
				"iteration",
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project to assign the permissions.`,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `(Required) The`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) the permissions to assign. The following permissions are available.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The name of the branch to assign the permissions.`,
				},
				resource.Attribute{
					Name:        "replace",
					Description: `(Optional) Replace (` + "`" + `true` + "`" + `) or merge (` + "`" + `false` + "`" + `) the permissions. Default: ` + "`" + `true` + "`" + ` | Permission | Description | |-----------------|--------------------------------| | GENERIC_READ | View permissions for this node | | GENERIC_WRITE | Edit this node | | CREATE_CHILDREN | Create child nodes | | DELETE | Delete this node | ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_project",
			Category:         "Resources",
			ShortDescription: `Manages a project within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"project",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_project_features",
			Category:         "Resources",
			ShortDescription: `Manages features for Azure DevOps projects.`,
			Description:      ``,
			Keywords: []string{
				"project",
				"features",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_project_permissions",
			Category:         "Resources",
			ShortDescription: `Manages permissions for a AzureDevOps project`,
			Description:      ``,
			Keywords: []string{
				"project",
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project to assign the permissions.`,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `(Required) The`,
				},
				resource.Attribute{
					Name:        "replace",
					Description: `(Optional) Replace (` + "`" + `true` + "`" + `) or merge (` + "`" + `false` + "`" + `) the permissions. Default: ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) the permissions to assign. The following permissions are available | Permission | Description | |------------------------------|----------------------------------------------| | GENERIC_READ | View project-level information | | GENERIC_WRITE | Edit project-level information | | DELETE | Delete team project | | PUBLISH_TEST_RESULTS | Create test runs | | ADMINISTER_BUILD | Administer a build | | START_BUILD | Start a build | | EDIT_BUILD_STATUS | Edit build quality | | UPDATE_BUILD | Write to build operational store | | DELETE_TEST_RESULTS | Delete test runs | | VIEW_TEST_RESULTS | View test runs | | MANAGE_TEST_ENVIRONMENTS | Manage test environments | | MANAGE_TEST_CONFIGURATIONS | Manage test configurations | | WORK_ITEM_DELETE | Delete and restore work items | | WORK_ITEM_MOVE | Move work items out of this project | | WORK_ITEM_PERMANENTLY_DELETE | Permanently delete work items | | RENAME | Rename team project | | MANAGE_PROPERTIES | Manage project properties | | MANAGE_SYSTEM_PROPERTIES | Manage system project properties | | BYPASS_PROPERTY_CACHE | Bypass project property cache | | BYPASS_RULES | Bypass rules on work item updates | | SUPPRESS_NOTIFICATIONS | Suppress notifications for work item updates | | UPDATE_VISIBILITY | Update project visibility | | CHANGE_PROCESS | Change process of team project. | | AGILETOOLS_BACKLOG | Agile backlog management. | | AGILETOOLS_PLANS | Agile plans. | ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_project_pipeline_settings",
			Category:         "Resources",
			ShortDescription: `Manages Pipeline Settings for Azure DevOps projects.`,
			Description:      ``,
			Keywords: []string{
				"project",
				"pipeline",
				"settings",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_repository_policy_author_email_pattern",
			Category:         "Resources",
			ShortDescription: `Manages author email pattern repository policy within Azure DevOps project.`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"policy",
				"author",
				"email",
				"pattern",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_repository_policy_case_enforcement",
			Category:         "Resources",
			ShortDescription: `Manages a case enforcement repository policy within Azure DevOps project.`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"policy",
				"case",
				"enforcement",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_repository_policy_check_credentials",
			Category:         "Resources",
			ShortDescription: `Manage a credentials check repository policy within Azure DevOps project.`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"policy",
				"check",
				"credentials",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_repository_policy_file_path_pattern",
			Category:         "Resources",
			ShortDescription: `Manages a file path pattern repository policy within Azure DevOps project.`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"policy",
				"file",
				"path",
				"pattern",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_repository_policy_max_file_size",
			Category:         "Resources",
			ShortDescription: `Manages a max file size repository policy within Azure DevOps project.`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"policy",
				"max",
				"file",
				"size",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_repository_policy_max_path_length",
			Category:         "Resources",
			ShortDescription: `Manages a max path length repository policy within Azure DevOps project.`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"policy",
				"max",
				"path",
				"length",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_repository_policy_reserved_names",
			Category:         "Resources",
			ShortDescription: `Manage a reserved names repository policy within Azure DevOps project.`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"policy",
				"reserved",
				"names",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_resource_authorization",
			Category:         "Resources",
			ShortDescription: `Manages authorization of resources within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"authorization",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_argocd",
			Category:         "Resources",
			ShortDescription: `Manages a ArgoCD server endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"argocd",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_artifactory",
			Category:         "Resources",
			ShortDescription: `Manages an Artifactory server endpoint within an Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"artifactory",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project.`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `(Required) The Service Endpoint name.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) URL of the Artifactory server to connect with. _Note: URL should not end in a slash character._`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `Authentication Token generated through Artifactory.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Artifactory Username.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Artifactory Password.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Service Endpoint description. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service endpoint.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project.`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `The Service Endpoint name. ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service endpoint.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project.`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `The Service Endpoint name. ## Relevant Links`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_aws",
			Category:         "Resources",
			ShortDescription: `Manages a AWS service endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"aws",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project.`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `(Required) The Service Endpoint name.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `(Required) The AWS access key ID for signing programmatic requests.`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Required) The AWS secret access key for signing programmatic requests.`,
				},
				resource.Attribute{
					Name:        "session_token",
					Description: `(Optional) The AWS session token for signing programmatic requests.`,
				},
				resource.Attribute{
					Name:        "role_to_assume",
					Description: `(Optional) The Amazon Resource Name (ARN) of the role to assume.`,
				},
				resource.Attribute{
					Name:        "role_session_name",
					Description: `(Optional) Optional identifier for the assumed role session.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Optional) A unique identifier that is used by third parties when assuming roles in their customers' accounts, aka cross-account role access.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Service Endpoint description. Defaults to ` + "`" + `Managed by Terraform` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service endpoint.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project.`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `The Service Endpoint name. ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service endpoint.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project.`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `The Service Endpoint name. ## Relevant Links`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_azuredevops",
			Category:         "Resources",
			ShortDescription: `Manages a Azure DevOps service endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_azurerm",
			Category:         "Resources",
			ShortDescription: `Manages a AzureRM service endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"azurerm",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_bitbucket",
			Category:         "Resources",
			ShortDescription: `Manages a Bitbucket service endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"bitbucket",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_dockerregistry",
			Category:         "Resources",
			ShortDescription: `Manages a Docker Registry service endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"dockerregistry",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_externaltfs",
			Category:         "Resources",
			ShortDescription: `Manages an Azure Repos/Team Foundation Server service endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"externaltfs",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_generic",
			Category:         "Resources",
			ShortDescription: `Manages a generic service endpoint within Azure DevOps, which can be used to authenticate to any external server using basic authentication via a username and password.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"generic",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_generic_git",
			Category:         "Resources",
			ShortDescription: `Manages a generic service endpoint within Azure DevOps, which can be used to authenticate to any external git service using basic authentication via a username and password.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"generic",
				"git",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_github",
			Category:         "Resources",
			ShortDescription: `Manages a GitHub service endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"github",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_github_enterprise",
			Category:         "Resources",
			ShortDescription: `Manages a GitHub-Enterprise-Server service endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"github",
				"enterprise",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_incomingwebhook",
			Category:         "Resources",
			ShortDescription: `Manages a Service Connection Incoming WebHook.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"incomingwebhook",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service endpoint.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 2 minutes) Used when creating the Service Connection Incoming WebHook.`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 1 minute) Used when retrieving the Service Connection Incoming WebHook.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 2 minutes) Used when updating the Service Connection Incoming WebHook.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 2 minutes) Used when deleting the Service Connection Incoming WebHook. ## Import Azure DevOps Service Endpoint Incoming WebHook can be imported using`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_kubernetes",
			Category:         "Resources",
			ShortDescription: `Manages a Kubernetes service endpoint Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"kubernetes",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_npm",
			Category:         "Resources",
			ShortDescription: `Manages a npm server endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"npm",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_octopusdeploy",
			Category:         "Resources",
			ShortDescription: `Manages an Octopus Deploy service endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"octopusdeploy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_permissions",
			Category:         "Resources",
			ShortDescription: `Manages permissions for a AzureDevOps Service Endpoint`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project.`,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `(Required) The`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) the permissions to assign. The following permissions are available.`,
				},
				resource.Attribute{
					Name:        "serviceendpoint_id",
					Description: `(Optional) The id of the service endpoint to assign the permissions.`,
				},
				resource.Attribute{
					Name:        "replace",
					Description: `(Optional) Replace (` + "`" + `true` + "`" + `) or merge (` + "`" + `false` + "`" + `) the permissions. Default: ` + "`" + `true` + "`" + ` | Permission | Description | | ----------------- | ----------------------------------- | | Use | Use service endpoint | | Administer | Full control over service endpoints | | Create | Create service endpoints | | ViewAuthorization | View authorizations | | ViewEndpoint | View service endpoint properties | ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_runpipeline",
			Category:         "Resources",
			ShortDescription: `Manages a Azure DevOps plugin RunPipeline.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"runpipeline",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_servicefabric",
			Category:         "Resources",
			ShortDescription: `Manages a Service Fabric service endpoint Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"servicefabric",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_sonarcloud",
			Category:         "Resources",
			ShortDescription: `Manages the SonarCloud service endpoint within an Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"sonarcloud",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project.`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `(Required) The Service Endpoint name.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required) Authentication Token generated through SonarCloud (go to ` + "`" + `My Account > Security > Generate Tokens` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Service Endpoint description. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service endpoint.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project.`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `The Service Endpoint name. ## Relevant Links - [Azure DevOps Service REST API 6.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0) - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml) - [SonarCloud User Token](https://docs.sonarcloud.io/advanced-setup/user-accounts/) ## Import Azure DevOps Service Endpoint SonarCloud can be imported using the`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service endpoint.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project.`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `The Service Endpoint name. ## Relevant Links - [Azure DevOps Service REST API 6.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0) - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml) - [SonarCloud User Token](https://docs.sonarcloud.io/advanced-setup/user-accounts/) ## Import Azure DevOps Service Endpoint SonarCloud can be imported using the`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_sonarqube",
			Category:         "Resources",
			ShortDescription: `Manages a SonarQube server endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"sonarqube",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project.`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `(Required) The Service Endpoint name.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) URL of the SonarQube server to connect with.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required) Authentication Token generated through SonarQube (go to My Account > Security > Generate Tokens).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Service Endpoint description. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service endpoint.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project.`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `The Service Endpoint name. ## Relevant Links - [Azure DevOps Service REST API 6.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0) - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml) - [SonarQube User Token](https://docs.sonarqube.org/latest/user-guide/user-token/) ## Import Azure DevOps Service Endpoint SonarQube can be imported using the`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service endpoint.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project.`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `The Service Endpoint name. ## Relevant Links - [Azure DevOps Service REST API 6.0 - Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-6.0) - [Azure DevOps Service Connections](https://docs.microsoft.com/en-us/azure/devops/pipelines/library/service-endpoints?view=azure-devops&tabs=yaml) - [SonarQube User Token](https://docs.sonarqube.org/latest/user-guide/user-token/) ## Import Azure DevOps Service Endpoint SonarQube can be imported using the`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_serviceendpoint_ssh",
			Category:         "Resources",
			ShortDescription: `Manages a SSH service endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"ssh",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_servicehook_permissions",
			Category:         "Resources",
			ShortDescription: `Manages permissions for AzureDevOps service hooks`,
			Description:      ``,
			Keywords: []string{
				"servicehook",
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(optional) The ID of the project.`,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `(Required) The`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) the permissions to assign. The following permissions are available.`,
				},
				resource.Attribute{
					Name:        "replace",
					Description: `(Optional) Replace (` + "`" + `true` + "`" + `) or merge (` + "`" + `false` + "`" + `) the permissions. Default: ` + "`" + `true` + "`" + ` | Name | Permission Description | | ------------------ | ------------------------ | | ViewSubscriptions | View Subscriptions | | EditSubscriptions | Edit Subscription | | DeleteSubscriptions| Delete Subscriptions | | PublishEvents | Publish Events | ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_tagging_permissions",
			Category:         "Resources",
			ShortDescription: `Manages permissions for AzureDevOps Tagging`,
			Description:      ``,
			Keywords: []string{
				"tagging",
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The ID of the project to assign the permissions. If omitted, organization wide permissions for tagging are managed.`,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `(Required) The`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) the permissions to assign. The following permissions are available.`,
				},
				resource.Attribute{
					Name:        "replace",
					Description: `(Optional) Replace (` + "`" + `true` + "`" + `) or merge (` + "`" + `false` + "`" + `) the permissions. Default: ` + "`" + `true` + "`" + ` | Name | Permission Description | | ------------------ | -------------------------- | | Enumerate | Enumerate tag definitions | | Create | Create tag definition | | Update | Update tag definition | | Delete | Delete tag definition | ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_team",
			Category:         "Resources",
			ShortDescription: `Manages a team within a project in a Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"team",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_team_administrators",
			Category:         "Resources",
			ShortDescription: `Manages administrators of a team within a project in a Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"team",
				"administrators",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_team_members",
			Category:         "Resources",
			ShortDescription: `Manages members of a team within a project in a Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"team",
				"members",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_user_entitlement",
			Category:         "Resources",
			ShortDescription: `Manages a user entitlement within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"entitlement",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_variable_group",
			Category:         "Resources",
			ShortDescription: `Manages variable groups within Azure DevOps project.`,
			Description:      ``,
			Keywords: []string{
				"variable",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_workitem",
			Category:         "Resources",
			ShortDescription: `Manages a Work Item in Azure Devops.`,
			Description:      ``,
			Keywords: []string{
				"workitem",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Work Item. ## Import Work Item resource does not support import.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_workitemquery_permissions",
			Category:         "Resources",
			ShortDescription: `Manages permissions for Work Item Queries`,
			Description:      ``,
			Keywords: []string{
				"workitemquery",
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project to assign the permissions.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path to a query or folder beneath ` + "`" + `Shared Queries` + "`" + ``,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `(Required) The`,
				},
				resource.Attribute{
					Name:        "replace",
					Description: `(Optional) Replace (` + "`" + `true` + "`" + `) or merge (` + "`" + `false` + "`" + `) the permissions. Default: ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) the permissions to assign. The following permissions are available | Permissions | Description | |--------------------------|------------------------------------| | Read | Read | | Contribute | Contribute | | Delete | Delete | | ManagePermissions | Manage Permissions | ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"azuredevops_agent_pool":                             0,
		"azuredevops_agent_queue":                            1,
		"azuredevops_area_permissions":                       2,
		"azuredevops_branch_policy_auto_reviewers":           3,
		"azuredevops_branch_policy_build_validation":         4,
		"azuredevops_branch_policy_comment_resolution":       5,
		"azuredevops_branch_policy_merge_types":              6,
		"azuredevops_branch_policy_min_reviewers":            7,
		"azuredevops_branch_policy_status_check":             8,
		"azuredevops_branch_policy_work_item_linking":        9,
		"azuredevops_build_definition":                       10,
		"azuredevops_build_definition_permissions":           11,
		"azuredevops_build_folder":                           12,
		"azuredevops_build_folder_permissions":               13,
		"azuredevops_check_branch_control":                   14,
		"azuredevops_check_business_hours":                   15,
		"azuredevops_git_permissions":                        16,
		"azuredevops_git_repository":                         17,
		"azuredevops_git_repository_branch":                  18,
		"azuredevops_git_repository_file":                    19,
		"azuredevops_group":                                  20,
		"azuredevops_group_membership":                       21,
		"azuredevops_iteration_permissions":                  22,
		"azuredevops_project":                                23,
		"azuredevops_project_features":                       24,
		"azuredevops_project_permissions":                    25,
		"azuredevops_project_pipeline_settings":              26,
		"azuredevops_repository_policy_author_email_pattern": 27,
		"azuredevops_repository_policy_case_enforcement":     28,
		"azuredevops_repository_policy_check_credentials":    29,
		"azuredevops_repository_policy_file_path_pattern":    30,
		"azuredevops_repository_policy_max_file_size":        31,
		"azuredevops_repository_policy_max_path_length":      32,
		"azuredevops_repository_policy_reserved_names":       33,
		"azuredevops_resource_authorization":                 34,
		"azuredevops_serviceendpoint_argocd":                 35,
		"azuredevops_serviceendpoint_artifactory":            36,
		"azuredevops_serviceendpoint_aws":                    37,
		"azuredevops_serviceendpoint_azuredevops":            38,
		"azuredevops_serviceendpoint_azurerm":                39,
		"azuredevops_serviceendpoint_bitbucket":              40,
		"azuredevops_serviceendpoint_dockerregistry":         41,
		"azuredevops_serviceendpoint_externaltfs":            42,
		"azuredevops_serviceendpoint_generic":                43,
		"azuredevops_serviceendpoint_generic_git":            44,
		"azuredevops_serviceendpoint_github":                 45,
		"azuredevops_serviceendpoint_github_enterprise":      46,
		"azuredevops_serviceendpoint_incomingwebhook":        47,
		"azuredevops_serviceendpoint_kubernetes":             48,
		"azuredevops_serviceendpoint_npm":                    49,
		"azuredevops_serviceendpoint_octopusdeploy":          50,
		"azuredevops_serviceendpoint_permissions":            51,
		"azuredevops_serviceendpoint_runpipeline":            52,
		"azuredevops_serviceendpoint_servicefabric":          53,
		"azuredevops_serviceendpoint_sonarcloud":             54,
		"azuredevops_serviceendpoint_sonarqube":              55,
		"azuredevops_serviceendpoint_ssh":                    56,
		"azuredevops_servicehook_permissions":                57,
		"azuredevops_tagging_permissions":                    58,
		"azuredevops_team":                                   59,
		"azuredevops_team_administrators":                    60,
		"azuredevops_team_members":                           61,
		"azuredevops_user_entitlement":                       62,
		"azuredevops_variable_group":                         63,
		"azuredevops_workitem":                               64,
		"azuredevops_workitemquery_permissions":              65,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
