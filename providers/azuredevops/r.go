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
					Description: `(Required) The project ID or project name.`,
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
					Description: `The project ID or project name.`,
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
					Description: `The project ID or project name.`,
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
					Description: `(Required) The project ID or project name.`,
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
					Description: `The project ID or project name.`,
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
					Description: `The project ID or project name.`,
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
					Description: `(Required) The project ID or project name.`,
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
					Description: `The project ID or project name.`,
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
					Description: `The project ID or project name.`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `The Service Endpoint name. ## Relevant Links`,
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

		"azuredevops_agent_pool":                        0,
		"azuredevops_agent_queue":                       1,
		"azuredevops_area_permissions":                  2,
		"azuredevops_branch_policy_auto_reviewers":      3,
		"azuredevops_branch_policy_build_validation":    4,
		"azuredevops_branch_policy_comment_resolution":  5,
		"azuredevops_branch_policy_merge_types":         6,
		"azuredevops_branch_policy_min_reviewers":       7,
		"azuredevops_branch_policy_status_check":        8,
		"azuredevops_branch_policy_work_item_linking":   9,
		"azuredevops_build_definition":                  10,
		"azuredevops_build_definition_permissions":      11,
		"azuredevops_git_permissions":                   12,
		"azuredevops_git_repository":                    13,
		"azuredevops_group":                             14,
		"azuredevops_group_membership":                  15,
		"azuredevops_iteration_permissions":             16,
		"azuredevops_project":                           17,
		"azuredevops_project_features":                  18,
		"azuredevops_project_permissions":               19,
		"azuredevops_resource_authorization":            20,
		"azuredevops_serviceendpoint_artifactory":       21,
		"azuredevops_serviceendpoint_aws":               22,
		"azuredevops_serviceendpoint_azuredevops":       23,
		"azuredevops_serviceendpoint_azurerm":           24,
		"azuredevops_serviceendpoint_bitbucket":         25,
		"azuredevops_serviceendpoint_dockerregistry":    26,
		"azuredevops_serviceendpoint_github":            27,
		"azuredevops_serviceendpoint_github_enterprise": 28,
		"azuredevops_serviceendpoint_kubernetes":        29,
		"azuredevops_serviceendpoint_npm":               30,
		"azuredevops_serviceendpoint_runpipeline":       31,
		"azuredevops_serviceendpoint_servicefabric":     32,
		"azuredevops_serviceendpoint_sonarqube":         33,
		"azuredevops_serviceendpoint_ssh":               34,
		"azuredevops_user_entitlement":                  35,
		"azuredevops_variable_group":                    36,
		"azuredevops_workitemquery_permissions":         37,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
