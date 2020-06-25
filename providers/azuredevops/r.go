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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the agent pool.`,
				},
				resource.Attribute{
					Name:        "auto_provision",
					Description: `(Optional) Specifies whether or not a queue should be automatically provisioned for each project collection. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pool_type",
					Description: `(Optional) Specifies whether the agent pool type is Automation or Deployment. Defaults to ` + "`" + `automation` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the agent pool. ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the agent pool. ## Relevant Links`,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project in which to create the resource.`,
				},
				resource.Attribute{
					Name:        "agent_pool_id",
					Description: `(Required) The ID of the organization agent pool. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the agent queue reference. ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the agent queue reference. ## Relevant Links`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_azure_git_repository",
			Category:         "Resources",
			ShortDescription: `Manages a git repository within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"azure",
				"git",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project ID or project name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the git repository.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Optional) The ID of a Git project from which a fork is to be created.`,
				},
				resource.Attribute{
					Name:        "initialization",
					Description: `(Optional) An ` + "`" + `initialization` + "`" + ` block as documented below. ` + "`" + `initialization` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "init_type",
					Description: `(Required) The type of repository to create. Valid values: ` + "`" + `Uninitialized` + "`" + `, ` + "`" + `Clean` + "`" + `, or ` + "`" + `Import` + "`" + `. Defaults to ` + "`" + `Uninitialized` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Optional) Type type of the source repository. Used if the ` + "`" + `init_type` + "`" + ` is ` + "`" + `Import` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_url",
					Description: `(Optional) The URL of the source repository. Used if the ` + "`" + `init_type` + "`" + ` is ` + "`" + `Import` + "`" + `. ## Attributes Reference In addition to all arguments above, except ` + "`" + `initialization` + "`" + `, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Git repository.`,
				},
				resource.Attribute{
					Name:        "default_branch",
					Description: `The ref of the default branch.`,
				},
				resource.Attribute{
					Name:        "is_fork",
					Description: `True if the repository was created as a fork.`,
				},
				resource.Attribute{
					Name:        "remote_url",
					Description: `Git HTTPS URL of the repository`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size in bytes.`,
				},
				resource.Attribute{
					Name:        "ssh_url",
					Description: `Git SSH URL of the repository.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `REST API URL of the repository.`,
				},
				resource.Attribute{
					Name:        "web_url",
					Description: `Web link to the repository. ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Git repository.`,
				},
				resource.Attribute{
					Name:        "default_branch",
					Description: `The ref of the default branch.`,
				},
				resource.Attribute{
					Name:        "is_fork",
					Description: `True if the repository was created as a fork.`,
				},
				resource.Attribute{
					Name:        "remote_url",
					Description: `Git HTTPS URL of the repository`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size in bytes.`,
				},
				resource.Attribute{
					Name:        "ssh_url",
					Description: `Git SSH URL of the repository.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `REST API URL of the repository.`,
				},
				resource.Attribute{
					Name:        "web_url",
					Description: `Web link to the repository. ## Relevant Links`,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project in which the policy will be created.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) A flag indicating if the policy should be enabled. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "blocking",
					Description: `(Optional) A flag indicating if the policy should be blocking. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `(Required) Configuration for the policy. This block must be defined exactly once. A ` + "`" + `settings` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "build_definition_id",
					Description: `(Required) The ID of the build to monitor for the policy.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name for the policy.`,
				},
				resource.Attribute{
					Name:        "manual_queue_only",
					Description: `(Optional) If set to true, the build will need to be manually queued. Defaults to ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "queue_on_source_update_only",
					Description: `(Optional) True if the build should queue on source updates only. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "valid_duration",
					Description: `(Optional) The number of minutes for which the build is valid. If ` + "`" + `0` + "`" + `, the build will not expire. Defaults to ` + "`" + `720` + "`" + ` (12 hours).`,
				},
				resource.Attribute{
					Name:        "repository_id",
					Description: `(Optional) The repository ID. Needed only if the scope of the policy will be limited to a single repository.`,
				},
				resource.Attribute{
					Name:        "repository_ref",
					Description: `(Optional) The ref pattern to use for the match. If ` + "`" + `match_type` + "`" + ` is ` + "`" + `Exact` + "`" + `, this should be a qualified ref such as ` + "`" + `refs/heads/master` + "`" + `. If ` + "`" + `match_type` + "`" + ` is ` + "`" + `Prefix` + "`" + `, this should be a ref path such as ` + "`" + `refs/heads/releases` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of branch policy configuration. ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of branch policy configuration. ## Relevant Links`,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project in which the policy will be created.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) A flag indicating if the policy should be enabled. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "blocking",
					Description: `(Optional) A flag indicating if the policy should be blocking. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `(Required) Configuration for the policy. This block must be defined exactly once. A ` + "`" + `settings` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "reviewer_count",
					Description: `(Required) The number of reviewrs needed to approve.`,
				},
				resource.Attribute{
					Name:        "submitter_can_vote",
					Description: `(Optional) Controls whether or not the submitter's vote counts. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "repository_id",
					Description: `(Optional) The repository ID. Needed only if the scope of the policy will be limited to a single repository.`,
				},
				resource.Attribute{
					Name:        "repository_ref",
					Description: `(Optional) The ref pattern to use for the match. If ` + "`" + `match_type` + "`" + ` is ` + "`" + `Exact` + "`" + `, this should be a qualified ref such as ` + "`" + `refs/heads/master` + "`" + `. If ` + "`" + `match_type` + "`" + ` is ` + "`" + `Prefix` + "`" + `, this should be a ref path such as ` + "`" + `refs/heads/releases` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of branch policy configuration. ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of branch policy configuration. ## Relevant Links`,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project ID or project name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the build definition.`,
				},
				resource.Attribute{
					Name:        "agent_pool_name",
					Description: `(Optional) The agent pool that should execute the build. Defaults to ` + "`" + `Hosted Ubuntu 1604` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) A ` + "`" + `repository` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "ci_trigger",
					Description: `(Optional) Continuous Integration Integration trigger.`,
				},
				resource.Attribute{
					Name:        "pull_request_trigger",
					Description: `(Optional) Pull Request Integration Integration trigger.`,
				},
				resource.Attribute{
					Name:        "variable_groups",
					Description: `(Optional) A list of variable group IDs (integers) to link to the build definition.`,
				},
				resource.Attribute{
					Name:        "variable",
					Description: `(Optional) A list of ` + "`" + `variable` + "`" + ` blocks, as documented below. ` + "`" + `variable` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the variable.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value of the variable.`,
				},
				resource.Attribute{
					Name:        "secret_value",
					Description: `(Optional) The secret value of the variable. Used when ` + "`" + `is_secret` + "`" + ` set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_secret",
					Description: `(Optional) True if the variable is a secret. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_override",
					Description: `(Optional) True if the variable can be overridden. Defaults to ` + "`" + `true` + "`" + `. ` + "`" + `repository` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "branch_name",
					Description: `(Optional) The branch name for which builds are triggered. Defaults to ` + "`" + `master` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "repo_id",
					Description: `(Required) The id of the repository. For ` + "`" + `TfsGit` + "`" + ` repos, this is simply the ID of the repository. For ` + "`" + `Github` + "`" + ` repos, this will take the form of ` + "`" + `<GitHub Org>/<Repo Name>` + "`" + `. For ` + "`" + `Bitbucket` + "`" + ` repos, this will take the form of ` + "`" + `<Workspace ID>/<Repo Name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "repo_type",
					Description: `(Optional) The repository type. Valid values: ` + "`" + `GitHub` + "`" + ` or ` + "`" + `TfsGit` + "`" + ` or ` + "`" + `Bitbucket` + "`" + `. Defaults to ` + "`" + `Github` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_connection_id",
					Description: `(Optional) The service connection ID. Used if the ` + "`" + `repo_type` + "`" + ` is ` + "`" + `GitHub` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "yml_path",
					Description: `(Required) The path of the Yaml file describing the build definition. ` + "`" + `ci_trigger` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "use_yaml",
					Description: `(Optional) Use the azure-pipeline file for the build configuration. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "override",
					Description: `(Optional) Override the azure-pipeline file and use a this configuration for all builds. ` + "`" + `ci_trigger` + "`" + ` ` + "`" + `override` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "batch",
					Description: `(Optional) If you set batch to true, when a pipeline is running, the system waits until the run is completed, then starts another run with all changes that have not yet been built. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "branch_filter",
					Description: `(Optional) The branches to include and exclude from the trigger.`,
				},
				resource.Attribute{
					Name:        "path_filter",
					Description: `(Optional) Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.`,
				},
				resource.Attribute{
					Name:        "max_concurrent_builds_per_branch",
					Description: `(Optional) The number of max builds per branch. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "polling_interval",
					Description: `(Optional) How often the external repository is polled. Defaults to ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "polling_job_id",
					Description: `(Computed) This is the ID of the polling job that polls the external repository. Once the build definition is saved/updated, this value is set. ` + "`" + `pull_request_trigger` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "use_yaml",
					Description: `(Optional) Use the azure-pipeline file for the build configuration. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "initial_branch",
					Description: `(Optional) When use_yaml is true set this to the name of the branch that the azure-pipelines.yml exists on. Defaults to ` + "`" + `Managed by Terraform` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "forks",
					Description: `(Required) Set permissions for Forked repositories.`,
				},
				resource.Attribute{
					Name:        "override",
					Description: `(Optional) Override the azure-pipeline file and use a this configuration for all builds. ` + "`" + `forks` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Build pull requests form forms of this repository.`,
				},
				resource.Attribute{
					Name:        "share_secrets",
					Description: `(Required) Make secrets available to builds of forks. ` + "`" + `pull_request_trigger` + "`" + ` ` + "`" + `override` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_cancel",
					Description: `(Optional) . Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "branch_filter",
					Description: `(Optional) The branches to include and exclude from the trigger.`,
				},
				resource.Attribute{
					Name:        "path_filter",
					Description: `(Optional) Specify file paths to include or exclude. Note that the wildcard syntax is different between branches/tags and file paths.`,
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
					Description: `(Optional) List of path patterns to exclude. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the build definition`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `The revision of the build definition ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the build definition`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `The revision of the build definition ## Relevant Links`,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the group. A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization.x`,
				},
				resource.Attribute{
					Name:        "origin_id",
					Description: `(Optional) The OriginID as a reference to a group from an external AD or AAD backed provider. The ` + "`" + `scope` + "`" + `, ` + "`" + `mail` + "`" + ` and ` + "`" + `display_name` + "`" + ` arguments cannot be used simultaneously with ` + "`" + `origin_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mail",
					Description: `(Optional) The mail address as a reference to an existing group from an external AD or AAD backed provider. The ` + "`" + `scope` + "`" + `, ` + "`" + `origin_id` + "`" + ` and ` + "`" + `display_name` + "`" + ` arguments cannot be used simultaneously with ` + "`" + `mail` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The name of a new Azure DevOps group that is not backed by an external provider. The ` + "`" + `origin_id` + "`" + ` and ` + "`" + `mail` + "`" + ` arguments cannot be used simultaneously with ` + "`" + `display_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Description of the Project.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) > NOTE: It's possible to define group members both within the ` + "`" + `azuredevops_group` + "`" + ` resource via the members block and by using the ` + "`" + `azuredevops_group_membership` + "`" + ` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Group.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `This url is the full route to the source resource of this graph subject.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `The type of source provider for the origin identifier (ex:AD, AAD, MSA)`,
				},
				resource.Attribute{
					Name:        "subject_kind",
					Description: `This field identifies the type of the graph subject (ex: Group, Scope, User).`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `This represents the name of the container of origin for a graph member.`,
				},
				resource.Attribute{
					Name:        "principal_name",
					Description: `This is the PrincipalName of this graph member from the source provider.`,
				},
				resource.Attribute{
					Name:        "descriptor",
					Description: `The identity (subject) descriptor of the Group. ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Group.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `This url is the full route to the source resource of this graph subject.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `The type of source provider for the origin identifier (ex:AD, AAD, MSA)`,
				},
				resource.Attribute{
					Name:        "subject_kind",
					Description: `This field identifies the type of the graph subject (ex: Group, Scope, User).`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `This represents the name of the container of origin for a graph member.`,
				},
				resource.Attribute{
					Name:        "principal_name",
					Description: `This is the PrincipalName of this graph member from the source provider.`,
				},
				resource.Attribute{
					Name:        "descriptor",
					Description: `The identity (subject) descriptor of the Group. ## Relevant Links`,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group",
					Description: `(Required) The descriptor of the group being managed.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) A list of user or group descriptors that will become members of the group. > NOTE: It's possible to define group members both within the ` + "`" + `azuredevops_group_membership resource` + "`" + ` via the members block and by using the ` + "`" + `azuredevops_group` + "`" + ` resource. However it's not possible to use both methods to manage group members, since there'll be conflicts.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The mode how the resource manages group members.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A random ID for this resource. There is no "natural" ID, so a random one is assigned. ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A random ID for this resource. There is no "natural" ID, so a random one is assigned. ## Relevant Links`,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_name",
					Description: `(Required) The Project Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Description of the Project.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) Specifies the visibility of the Project. Valid values: ` + "`" + `private` + "`" + ` or ` + "`" + `public` + "`" + `. Defaults to ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version_control",
					Description: `(Optional) Specifies the version control system. Valid values: ` + "`" + `Git` + "`" + ` or ` + "`" + `Tfvc` + "`" + `. Defaults to ` + "`" + `Git` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "work_item_template",
					Description: `(Optional) Specifies the work item template. Defaults to ` + "`" + `Agile` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `(Optional) Defines the status (` + "`" + `enabled` + "`" + `, ` + "`" + `disabled` + "`" + `) of the project features. Valid features ` + "`" + `boards` + "`" + `, ` + "`" + `repositories` + "`" + `, ` + "`" + `pipelines` + "`" + `, ` + "`" + `testplans` + "`" + `, ` + "`" + `artifacts` + "`" + ` >`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Project ID of the Project.`,
				},
				resource.Attribute{
					Name:        "process_template_id",
					Description: `The Process Template ID used by the Project. ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Project ID of the Project.`,
				},
				resource.Attribute{
					Name:        "process_template_id",
					Description: `The Process Template ID used by the Project. ## Relevant Links`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azuredevops_resource_authorization",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"resource",
				"authorization",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project ID or project name. Type: string.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The ID of the resource to authorize. Type: string.`,
				},
				resource.Attribute{
					Name:        "authorized",
					Description: `(Required) Set to true to allow public access in the project. Type: boolean.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the resource to authorize. Type: string. Valid values: ` + "`" + `endpoint` + "`" + `, ` + "`" + `queue` + "`" + `. Default value: ` + "`" + `endpoint` + "`" + `. ## Attributes Reference The following attributes are exported: n/a ## Relevant Links`,
				},
			},
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
					Name:        "azurerm_spn_tenantid",
					Description: `(Required) The tenant id if the service principal.`,
				},
				resource.Attribute{
					Name:        "azurerm_subscription_id",
					Description: `(Required) The subscription Id of the Azure targets.`,
				},
				resource.Attribute{
					Name:        "azurerm_subscription_name",
					Description: `(Required) The subscription Name of the targets.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Optional) A ` + "`" + `credentials` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Optional) The resource group used for scope of automatic service endpoint. --- A ` + "`" + `credentials` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "serviceprincipalid",
					Description: `(Required) The service principal application Id`,
				},
				resource.Attribute{
					Name:        "serviceprincipalkey",
					Description: `(Required) The service principal secret. ## Attributes Reference The following attributes are exported:`,
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
			Type:             "azuredevops_serviceendpoint_bitbucket",
			Category:         "Resources",
			ShortDescription: `Manages a Bitbucket service endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"bitbucket",
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
					Name:        "username",
					Description: `(Required) Bitbucket account username.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Bitbucket account password.`,
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
			Type:             "azuredevops_serviceendpoint_dockerregistry",
			Category:         "Resources",
			ShortDescription: `Manages a Docker Registry service endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"dockerregistry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project ID or project name.`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `(Required) The name you will use to refer to this service connection in task inputs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The name you will use to refer to this service connection in task inputs.`,
				},
				resource.Attribute{
					Name:        "docker_registry",
					Description: `(Optional) The URL of the Docker registry. (Default: "https://index.docker.io/v1/")`,
				},
				resource.Attribute{
					Name:        "docker_username",
					Description: `(Optional) The identifier of the Docker account user.`,
				},
				resource.Attribute{
					Name:        "docker_email",
					Description: `(Optional) The email for Docker account user.`,
				},
				resource.Attribute{
					Name:        "docker_password",
					Description: `(Optional) The password for the account user identified above.`,
				},
				resource.Attribute{
					Name:        "registry_type",
					Description: `(Optional) Can be "DockerHub" or "Others" (Default "DockerHub") ## Attributes Reference The following attributes are exported:`,
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
			Type:             "azuredevops_serviceendpoint_github",
			Category:         "Resources",
			ShortDescription: `Manages a GitHub service endpoint within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"github",
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
					Name:        "description",
					Description: `(Optional) The Service Endpoint description. Defaults to ` + "`" + `Managed by Terraform` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_personal",
					Description: `(Optional) An ` + "`" + `auth_personal` + "`" + ` block as documented below. Allows connecting using a personal access token.`,
				},
				resource.Attribute{
					Name:        "auth_oauth",
					Description: `(Optional) An ` + "`" + `auth_oauth` + "`" + ` block as documented below. Allows connecting using an Oauth token.`,
				},
				resource.Attribute{
					Name:        "personal_access_token",
					Description: `(Required) The Personal Access Token for Github. ` + "`" + `auth_oauth` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "oauth_configuration_id",
					Description: `(Required)`,
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
			Type:             "azuredevops_serviceendpoint_kubernetes",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"serviceendpoint",
				"kubernetes",
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
					Name:        "apiserver_url",
					Description: `(Required) The Service Endpoint description.`,
				},
				resource.Attribute{
					Name:        "authorization_type",
					Description: `(Required) The authentication method used to authenticate on the Kubernetes cluster. The value should be one of AzureSubscription, Kubeconfig, ServiceAccount.`,
				},
				resource.Attribute{
					Name:        "azure_subscription",
					Description: `(Optional) The configuration for authorization_type="AzureSubscription".`,
				},
				resource.Attribute{
					Name:        "azure_environment",
					Description: `(Optional) Azure environment refers to whether the public cloud offering or domestic (government) clouds are being used. Currently, only the public cloud is supported. The value must be AzureCloud. This is also the default-value.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required) The id of the Azure subscription.`,
				},
				resource.Attribute{
					Name:        "subscription_name",
					Description: `(Required) The name of the Azure subscription.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) The id of the tenant used by the subscription.`,
				},
				resource.Attribute{
					Name:        "resourcegroup_id",
					Description: `(Required) The resource group id, to which the Kubernetes cluster is deployed.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The Kubernetes namespace. Default value is "default".`,
				},
				resource.Attribute{
					Name:        "kubeconfig",
					Description: `(Optional) The configuration for authorization_type="Kubeconfig".`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `(Required) The content of the kubeconfig in yaml notation to be used to communicate with the API-Server of Kubernetes.`,
				},
				resource.Attribute{
					Name:        "accept_untrusted_certs",
					Description: `(Optional) Set this option to allow clients to accept a self-signed certificate.`,
				},
				resource.Attribute{
					Name:        "cluster_context",
					Description: `(Optional) Context within the kubeconfig file that is to be used for identifying the cluster. Default value is the current-context set in kubeconfig.`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `(Optional) The configuration for authorization_type="ServiceAccount". This type uses the credentials of a service account currently deployed to the cluster.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required) The token from a Kubernetes secret object.`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `(Required) The certificate from a Kubernetes secret object. ## Attributes Reference The following attributes are exported:`,
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
			Type:             "azuredevops_user_entitlement",
			Category:         "Resources",
			ShortDescription: `Manages a user entitlement within Azure DevOps organization.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"entitlement",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "principal_name",
					Description: `(Optional) The principal name is the PrincipalName of a graph member from the source provider. Usually, e-mail address.`,
				},
				resource.Attribute{
					Name:        "origin_id",
					Description: `(Optional) The unique identifier from the system of origin. Typically a sid, object id or Guid. e.g. Used for member of other tenant on Azure Active Directory.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) The type of source provider for the origin identifier.`,
				},
				resource.Attribute{
					Name:        "account_license_type",
					Description: `(Optional) Type of Account License. Valid values: ` + "`" + `advanced` + "`" + `, ` + "`" + `earlyAdopter` + "`" + `, ` + "`" + `express` + "`" + `, ` + "`" + `none` + "`" + `, ` + "`" + `professional` + "`" + `, or ` + "`" + `stakeholder` + "`" + `. Defaults to ` + "`" + `express` + "`" + `. In addition the value ` + "`" + `basic` + "`" + ` is allowed which is an alias for ` + "`" + `express` + "`" + ` and reflects the name of the ` + "`" + `express` + "`" + ` license used in the Azure DevOps web interface.`,
				},
				resource.Attribute{
					Name:        "licensing_source",
					Description: `(Optional) The source of the licensing (e.g. Account. MSDN etc.) Valid values: ` + "`" + `account` + "`" + ` (Default), ` + "`" + `auto` + "`" + `, ` + "`" + `msdn` + "`" + `, ` + "`" + `none` + "`" + `, ` + "`" + `profile` + "`" + `, ` + "`" + `trail` + "`" + ` >`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the entitlement.`,
				},
				resource.Attribute{
					Name:        "descriptor",
					Description: `The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject. ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the entitlement.`,
				},
				resource.Attribute{
					Name:        "descriptor",
					Description: `The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the user graph subject. ## Relevant Links`,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project ID or project name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Variable Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the Variable Group.`,
				},
				resource.Attribute{
					Name:        "allow_access",
					Description: `(Required) Boolean that indicate if this variable group is shared by all pipelines of this project.`,
				},
				resource.Attribute{
					Name:        "variable",
					Description: `(Optional) One or more ` + "`" + `variable` + "`" + ` blocks as documented below. A ` + "`" + `variable` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The key value used for the variable. Must be unique within the Variable Group.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value of the variable. If omitted, it will default to empty string.`,
				},
				resource.Attribute{
					Name:        "secret_value",
					Description: `(Optional) The secret value of the variable. If omitted, it will default to empty string. Used when ` + "`" + `is_secret` + "`" + ` set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_secret",
					Description: `(Optional) A boolean flag describing if the variable value is sensitive. Defaults to ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Variable Group returned after creation in Azure DevOps. ## Relevant Links`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Variable Group returned after creation in Azure DevOps. ## Relevant Links`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"azuredevops_agent_pool":                     0,
		"azuredevops_agent_queue":                    1,
		"azuredevops_azure_git_repository":           2,
		"azuredevops_branch_policy_build_validation": 3,
		"azuredevops_branch_policy_min_reviewers":    4,
		"azuredevops_build_definition":               5,
		"azuredevops_group":                          6,
		"azuredevops_group_membership":               7,
		"azuredevops_project":                        8,
		"azuredevops_resource_authorization":         9,
		"azuredevops_serviceendpoint_azurerm":        10,
		"azuredevops_serviceendpoint_bitbucket":      11,
		"azuredevops_serviceendpoint_dockerregistry": 12,
		"azuredevops_serviceendpoint_github":         13,
		"azuredevops_serviceendpoint_kubernetes":     14,
		"azuredevops_user_entitlement":               15,
		"azuredevops_variable_group":                 16,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
