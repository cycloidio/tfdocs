package gitlab

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "gitlab_branch_protection",
			Category:         "Resources",
			ShortDescription: `Protects a branch by assigning access levels to it`,
			Description:      ``,
			Keywords: []string{
				"branch",
				"protection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required) The id of the project.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `(Required) Name of the branch.`,
				},
				resource.Attribute{
					Name:        "push_access_level",
					Description: `(Required) One of five levels of access to the project.`,
				},
				resource.Attribute{
					Name:        "merge_access_level",
					Description: `(Required) One of five levels of access to the project.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_deploy_key",
			Category:         "Resources",
			ShortDescription: `Creates and manages deploy keys for GitLab projects`,
			Description:      ``,
			Keywords: []string{
				"deploy",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, string) The name or id of the project to add the deploy key to.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required, string) A title to describe the deploy key with.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required, string) The public ssh key body.`,
				},
				resource.Attribute{
					Name:        "can_push",
					Description: `(Optional, boolean) Allow this deploy key to be used to push changes to the project. Defaults to ` + "`" + `false` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_group",
			Category:         "Resources",
			ShortDescription: `Creates and manages GitLab groups`,
			Description:      ``,
			Keywords: []string{
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of this group.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path of the group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the group.`,
				},
				resource.Attribute{
					Name:        "lfs_enabled",
					Description: `(Optional) Boolean, defaults to true. Whether to enable LFS support for projects in this group.`,
				},
				resource.Attribute{
					Name:        "request_access_enabled",
					Description: `(Optional) Boolean, defaults to false. Whether to enable users to request access to the group.`,
				},
				resource.Attribute{
					Name:        "visibility_level",
					Description: `(Optional) Set to ` + "`" + `public` + "`" + ` to create a public group. Valid values are ` + "`" + `private` + "`" + `, ` + "`" + `internal` + "`" + `, ` + "`" + `public` + "`" + `. Groups are created as private by default.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Optional) Integer, id of the parent group (creates a nested group). ## Attributes Reference The resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the group by the GitLab server. Serves as a namespace id where one is needed.`,
				},
				resource.Attribute{
					Name:        "full_path",
					Description: `The full path of the group.`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `The full name of the group.`,
				},
				resource.Attribute{
					Name:        "web_url",
					Description: `Web URL of the group. ## Importing groups You can import a group state using ` + "`" + `terraform import <resource> <id>` + "`" + `. The ` + "`" + `id` + "`" + ` can be whatever the [details of a group][details_of_a_group] api takes for its ` + "`" + `:id` + "`" + ` value, so for example: terraform import gitlab_group.example example [details_of_a_group]: https://docs.gitlab.com/ee/api/groups.html#details-of-a-group`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the group by the GitLab server. Serves as a namespace id where one is needed.`,
				},
				resource.Attribute{
					Name:        "full_path",
					Description: `The full path of the group.`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `The full name of the group.`,
				},
				resource.Attribute{
					Name:        "web_url",
					Description: `Web URL of the group. ## Importing groups You can import a group state using ` + "`" + `terraform import <resource> <id>` + "`" + `. The ` + "`" + `id` + "`" + ` can be whatever the [details of a group][details_of_a_group] api takes for its ` + "`" + `:id` + "`" + ` value, so for example: terraform import gitlab_group.example example [details_of_a_group]: https://docs.gitlab.com/ee/api/groups.html#details-of-a-group`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_group_membership",
			Category:         "Resources",
			ShortDescription: `Adds a user to a group as a member`,
			Description:      ``,
			Keywords: []string{
				"group",
				"membership",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The id of the group.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) The id of the user.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `(Required) Acceptable values are: guest, reporter, developer, master.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `(Optional) Expiration date for the group membership. Format: ` + "`" + `YYYY-MM-DD` + "`" + ` ## Import GitLab group membership can be imported using an id made up of ` + "`" + `groupid:username` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_group_membership.test 12345:1337 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_group_variable",
			Category:         "Resources",
			ShortDescription: `Creates and manages CI/CD variables for GitLab groups`,
			Description:      ``,
			Keywords: []string{
				"group",
				"variable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group",
					Description: `(Required, string) The name or id of the group to add the hook to.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required, string) The name of the variable.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required, string) The value of the variable.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `(Optional, boolean) If set to ` + "`" + `true` + "`" + `, the variable will be passed only to pipelines running on protected branches and tags. Defaults to ` + "`" + `false` + "`" + `. ## Import GitLab group variables can be imported using an id made up of ` + "`" + `groupid:variablename` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_group_variable.example 12345:group_variable_key ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_label",
			Category:         "Resources",
			ShortDescription: `Creates and manages labels for GitLab projects`,
			Description:      ``,
			Keywords: []string{
				"label",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required) The name or id of the project to add the label to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the label.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Required) The color of the label given in 6-digit hex notation with leading '#' sign (e.g. #FFAABB) or one of the [CSS color names](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value#Color_keywords).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the label. ## Attributes Reference The resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the label by the GitLab server (the name of the label).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the label by the GitLab server (the name of the label).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_pipeline_schedule",
			Category:         "Resources",
			ShortDescription: `Creates and manages pipeline schedules for GitLab projects`,
			Description:      ``,
			Keywords: []string{
				"pipeline",
				"schedule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, string) The name or id of the project to add the schedule to.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required, string) The description of the pipeline schedule.`,
				},
				resource.Attribute{
					Name:        "ref",
					Description: `(Required, string) The branch/tag name to be triggered.`,
				},
				resource.Attribute{
					Name:        "cron",
					Description: `(Required, string) The cron (e.g. ` + "`" + `0 1`,
				},
				resource.Attribute{
					Name:        "cron_timezone",
					Description: `(Optional, string) The timezone.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional, bool) The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_pipeline_trigger",
			Category:         "Resources",
			ShortDescription: `Creates and manages pipeline triggers for GitLab projects`,
			Description:      ``,
			Keywords: []string{
				"pipeline",
				"trigger",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, string) The name or id of the project to add the trigger to.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required, string) The description of the pipeline trigger.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project",
			Category:         "Resources",
			ShortDescription: `Creates and manages projects within GitLab groups or within your user`,
			Description:      ``,
			Keywords: []string{
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path of the repository.`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `(Optional) The namespace (group or user) of the project. Defaults to your user. See [` + "`" + `gitlab_group` + "`" + `](group.html) for an example.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the project.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags (topics) of the project.`,
				},
				resource.Attribute{
					Name:        "default_branch",
					Description: `(Optional) The default branch for the project.`,
				},
				resource.Attribute{
					Name:        "issues_enabled",
					Description: `(Optional) Enable issue tracking for the project.`,
				},
				resource.Attribute{
					Name:        "merge_requests_enabled",
					Description: `(Optional) Enable merge requests for the project.`,
				},
				resource.Attribute{
					Name:        "approvals_before_merge",
					Description: `(Optional) Number of merge request approvals required for merging. Default is 0.`,
				},
				resource.Attribute{
					Name:        "wiki_enabled",
					Description: `(Optional) Enable wiki for the project.`,
				},
				resource.Attribute{
					Name:        "snippets_enabled",
					Description: `(Optional) Enable snippets for the project.`,
				},
				resource.Attribute{
					Name:        "container_registry_enabled",
					Description: `(Optional) Enable container registry for the project.`,
				},
				resource.Attribute{
					Name:        "visibility_level",
					Description: `(Optional) Set to ` + "`" + `public` + "`" + ` to create a public project. Valid values are ` + "`" + `private` + "`" + `, ` + "`" + `internal` + "`" + `, ` + "`" + `public` + "`" + `. Repositories are created as private by default.`,
				},
				resource.Attribute{
					Name:        "merge_method",
					Description: `(Optional) Set to ` + "`" + `ff` + "`" + ` to create fast-forward merges Valid values are ` + "`" + `merge` + "`" + `, ` + "`" + `rebase_merge` + "`" + `, ` + "`" + `ff` + "`" + ` Repositories are created with ` + "`" + `merge` + "`" + ` by default`,
				},
				resource.Attribute{
					Name:        "only_allow_merge_if_pipeline_succeeds",
					Description: `(Optional) Set to true if you want allow merges only if a pipeline succeeds.`,
				},
				resource.Attribute{
					Name:        "only_allow_merge_if_all_discussions_are_resolved",
					Description: `(Optional) Set to true if you want allow merges only if all discussions are resolved.`,
				},
				resource.Attribute{
					Name:        "shared_runners_enabled",
					Description: `(Optional) Enable shared runners for this project.`,
				},
				resource.Attribute{
					Name:        "shared_with_groups",
					Description: `(Optional) Enable sharing the project with a list of groups (maps).`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Group id of the group you want to share the project with.`,
				},
				resource.Attribute{
					Name:        "group_access_level",
					Description: `(Required) Group's sharing permissions. See [group members permission][group_members_permissions] for more info. Valid values are ` + "`" + `guest` + "`" + `, ` + "`" + `reporter` + "`" + `, ` + "`" + `developer` + "`" + `, ` + "`" + `master` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "archived",
					Description: `(Optional) Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter. ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Integer that uniquely identifies the project within the gitlab install.`,
				},
				resource.Attribute{
					Name:        "ssh_url_to_repo",
					Description: `URL that can be provided to ` + "`" + `git clone` + "`" + ` to clone the repository via SSH.`,
				},
				resource.Attribute{
					Name:        "http_url_to_repo",
					Description: `URL that can be provided to ` + "`" + `git clone` + "`" + ` to clone the repository via HTTP.`,
				},
				resource.Attribute{
					Name:        "web_url",
					Description: `URL that can be used to find the project in a browser.`,
				},
				resource.Attribute{
					Name:        "runners_token",
					Description: `Registration token to use during runner setup.`,
				},
				resource.Attribute{
					Name:        "shared_with_groups",
					Description: `List of the groups the project is shared with.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `Group's name. ## Importing projects You can import a project state using ` + "`" + `terraform import <resource> <id>` + "`" + `. The ` + "`" + `id` + "`" + ` can be whatever the [get single project api][get_single_project] takes for its ` + "`" + `:id` + "`" + ` value, so for example: terraform import gitlab_project.example richardc/example [get_single_project]: https://docs.gitlab.com/ee/api/projects.html#get-single-project [group_members_permissions]: https://docs.gitlab.com/ce/user/permissions.html#group-members-permissions`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Integer that uniquely identifies the project within the gitlab install.`,
				},
				resource.Attribute{
					Name:        "ssh_url_to_repo",
					Description: `URL that can be provided to ` + "`" + `git clone` + "`" + ` to clone the repository via SSH.`,
				},
				resource.Attribute{
					Name:        "http_url_to_repo",
					Description: `URL that can be provided to ` + "`" + `git clone` + "`" + ` to clone the repository via HTTP.`,
				},
				resource.Attribute{
					Name:        "web_url",
					Description: `URL that can be used to find the project in a browser.`,
				},
				resource.Attribute{
					Name:        "runners_token",
					Description: `Registration token to use during runner setup.`,
				},
				resource.Attribute{
					Name:        "shared_with_groups",
					Description: `List of the groups the project is shared with.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `Group's name. ## Importing projects You can import a project state using ` + "`" + `terraform import <resource> <id>` + "`" + `. The ` + "`" + `id` + "`" + ` can be whatever the [get single project api][get_single_project] takes for its ` + "`" + `:id` + "`" + ` value, so for example: terraform import gitlab_project.example richardc/example [get_single_project]: https://docs.gitlab.com/ee/api/projects.html#get-single-project [group_members_permissions]: https://docs.gitlab.com/ce/user/permissions.html#group-members-permissions`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project_cluster",
			Category:         "Resources",
			ShortDescription: `Creates and manages project clusters for GitLab projects`,
			Description:      ``,
			Keywords: []string{
				"project",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, string) The id of the project to add the cluster to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The name of cluster.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional, string) The base domain of the cluster.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, boolean) Determines if cluster is active or not. Defaults to ` + "`" + `true` + "`" + `. This attribute cannot be read.`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `(Optional, boolean) Determines if cluster is managed by gitlab or not. Defaults to ` + "`" + `true` + "`" + `. This attribute cannot be read.`,
				},
				resource.Attribute{
					Name:        "kubernetes_api_url",
					Description: `(Required, string) The URL to access the Kubernetes API.`,
				},
				resource.Attribute{
					Name:        "kubernetes_token",
					Description: `(Required, string) The token to authenticate against Kubernetes.`,
				},
				resource.Attribute{
					Name:        "kubernetes_ca_cert",
					Description: `(Optional, string) TLS certificate (needed if API is using a self-signed TLS certificate).`,
				},
				resource.Attribute{
					Name:        "kubernetes_namespace",
					Description: `(Optional, string) The unique namespace related to the project.`,
				},
				resource.Attribute{
					Name:        "kubernetes_authorization_type",
					Description: `(Optional, string) The cluster authorization type. Valid values are ` + "`" + `rbac` + "`" + `, ` + "`" + `abac` + "`" + `, ` + "`" + `unknown_authorization` + "`" + `. Defaults to ` + "`" + `rbac` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "environment_scope",
					Description: `(Optional, string) The associated environment to the cluster. Defaults to ` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project_hook",
			Category:         "Resources",
			ShortDescription: `Creates and manages hooks for GitLab projects`,
			Description:      ``,
			Keywords: []string{
				"project",
				"hook",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required) The name or id of the project to add the hook to.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The url of the hook to invoke.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) A token to present when invoking the hook.`,
				},
				resource.Attribute{
					Name:        "enable_ssl_verification",
					Description: `(Optional) Enable ssl verification when invoking the hook.`,
				},
				resource.Attribute{
					Name:        "push_events",
					Description: `(Optional) Invoke the hook for push events.`,
				},
				resource.Attribute{
					Name:        "issues_events",
					Description: `(Optional) Invoke the hook for issues events.`,
				},
				resource.Attribute{
					Name:        "merge_requests_events",
					Description: `(Optional) Invoke the hook for merge requests.`,
				},
				resource.Attribute{
					Name:        "tag_push_events",
					Description: `(Optional) Invoke the hook for tag push events.`,
				},
				resource.Attribute{
					Name:        "note_events",
					Description: `(Optional) Invoke the hook for notes events.`,
				},
				resource.Attribute{
					Name:        "job_events",
					Description: `(Optional) Invoke the hook for job events.`,
				},
				resource.Attribute{
					Name:        "pipeline_events",
					Description: `(Optional) Invoke the hook for pipeline events.`,
				},
				resource.Attribute{
					Name:        "wiki_page_events",
					Description: `(Optional) Invoke the hook for wiki page events. ## Attributes Reference The resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the hook by the GitLab server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the hook by the GitLab server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project_membership",
			Category:         "Resources",
			ShortDescription: `Adds a user to a project as a member`,
			Description:      ``,
			Keywords: []string{
				"project",
				"membership",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the project.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) The id of the user.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `(Required) One of five levels of access to the project. ## Import GitLab group membership can be imported using an id made up of ` + "`" + `groupid:username` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_project_membership.test 12345:1337`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project_variable",
			Category:         "Resources",
			ShortDescription: `Creates and manages CI/CD variables for GitLab projects`,
			Description:      ``,
			Keywords: []string{
				"project",
				"variable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, string) The name or id of the project to add the hook to.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required, string) The name of the variable.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required, string) The value of the variable.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `(Optional, boolean) If set to ` + "`" + `true` + "`" + `, the variable will be passed only to pipelines running on protected branches and tags. Defaults to ` + "`" + `false` + "`" + `. ## Import GitLab project variables can be imported using an id made up of ` + "`" + `projectid:variablename` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_project_variable.example 12345:project_variable_key ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_service_jira",
			Category:         "Resources",
			ShortDescription: `Manage Jira integration that allows to receive event notifications in Jira`,
			Description:      ``,
			Keywords: []string{
				"service",
				"jira",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required) ID of the project you want to activate integration on.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL to the JIRA project which is being linked to this GitLab project. For example, https://jira.example.com.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username of the user created to be used with GitLab/JIRA.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password of the user created to be used with GitLab/JIRA.`,
				},
				resource.Attribute{
					Name:        "project_key",
					Description: `(Required) The short identifier for your JIRA project, all uppercase, e.g., PROJ.`,
				},
				resource.Attribute{
					Name:        "jira_issue_transition_id",
					Description: `(Optional) The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2. ## Importing Jira service You can import a service_jira state using ` + "`" + `terraform import <resource> <project_id>` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import gitlab_service_jira.jira 1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_service_slack",
			Category:         "Resources",
			ShortDescription: `Manage Slack notifications integration that allows to receive event notifications in Slack`,
			Description:      ``,
			Keywords: []string{
				"service",
				"slack",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required) ID of the project you want to activate integration on.`,
				},
				resource.Attribute{
					Name:        "webhook",
					Description: `(Required) Webhook URL (ex.: https://hooks.slack.com/services/...)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Username to use.`,
				},
				resource.Attribute{
					Name:        "notify_only_broken_pipelines",
					Description: `(Optional) Send notifications for broken pipelines.`,
				},
				resource.Attribute{
					Name:        "notify_only_default_branch",
					Description: `(Optional) Send notifications only for the default branch.`,
				},
				resource.Attribute{
					Name:        "push_events",
					Description: `(Optional) Enable notifications for push events.`,
				},
				resource.Attribute{
					Name:        "push_channel",
					Description: `(Optional) The name of the channel to receive push events notifications.`,
				},
				resource.Attribute{
					Name:        "issues_events",
					Description: `(Optional) Enable notifications for issues events.`,
				},
				resource.Attribute{
					Name:        "issue_channel",
					Description: `(Optional) The name of the channel to receive issue events notifications.`,
				},
				resource.Attribute{
					Name:        "confidential_issues_events",
					Description: `(Optional) Enable notifications for confidential issues events.`,
				},
				resource.Attribute{
					Name:        "confidential_issue_channel",
					Description: `(Optional) The name of the channel to receive confidential issue events notifications.`,
				},
				resource.Attribute{
					Name:        "merge_requests_events",
					Description: `(Optional) Enable notifications for merge requests events.`,
				},
				resource.Attribute{
					Name:        "merge_request_channel",
					Description: `(Optional) The name of the channel to receive merge request events notifications.`,
				},
				resource.Attribute{
					Name:        "tag_push_events",
					Description: `(Optional) Enable notifications for tag push events.`,
				},
				resource.Attribute{
					Name:        "tag_push_channel",
					Description: `(Optional) The name of the channel to receive tag push events notifications.`,
				},
				resource.Attribute{
					Name:        "note_events",
					Description: `(Optional) Enable notifications for note events.`,
				},
				resource.Attribute{
					Name:        "note_channel",
					Description: `(Optional) The name of the channel to receive note events notifications.`,
				},
				resource.Attribute{
					Name:        "confidential_note_events",
					Description: `(Optional) Enable notifications for confidential note events.`,
				},
				resource.Attribute{
					Name:        "pipeline_events",
					Description: `(Optional) Enable notifications for pipeline events.`,
				},
				resource.Attribute{
					Name:        "pipeline_channel",
					Description: `(Optional) The name of the channel to receive pipeline events notifications.`,
				},
				resource.Attribute{
					Name:        "wiki_page_events",
					Description: `(Optional) Enable notifications for wiki page events.`,
				},
				resource.Attribute{
					Name:        "wiki_page_channel",
					Description: `(Optional) The name of the channel to receive wiki page events notifications. ## Importing Slack service You can import a service_slack state using ` + "`" + `terraform import <resource> <project_id>` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import gitlab_service_slack.slack 1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_tag_protection",
			Category:         "Resources",
			ShortDescription: `Protects a tag by assigning access levels to it`,
			Description:      ``,
			Keywords: []string{
				"tag",
				"protection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required) The id of the project.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Required) Name of the tag or wildcard.`,
				},
				resource.Attribute{
					Name:        "create_access_level",
					Description: `(Required) One of five levels of access to the project.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_user",
			Category:         "Resources",
			ShortDescription: `Creates and manages GitLab users`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password of the user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The e-mail address of the user.`,
				},
				resource.Attribute{
					Name:        "is_admin",
					Description: `(Optional) Boolean, defaults to false. Whether to enable administrative priviledges for the user.`,
				},
				resource.Attribute{
					Name:        "projects_limit",
					Description: `(Optional) Integer, defaults to 0. Number of projects user can create.`,
				},
				resource.Attribute{
					Name:        "can_create_group",
					Description: `(Optional) Boolean, defaults to false. Whether to allow the user to create groups.`,
				},
				resource.Attribute{
					Name:        "skip_confirmation",
					Description: `(Optional) Boolean, defaults to true. Whether to skip confirmation.`,
				},
				resource.Attribute{
					Name:        "is_external",
					Description: `(Optional) Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access. ## Attributes Reference The resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the user by the GitLab server. ## Importing users You can import a user to terraform state using ` + "`" + `terraform import <resource> <id>` + "`" + `. The ` + "`" + `id` + "`" + ` must be an integer for the id of the user you want to import, for example: terraform import gitlab_user.example 42`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the user by the GitLab server. ## Importing users You can import a user to terraform state using ` + "`" + `terraform import <resource> <id>` + "`" + `. The ` + "`" + `id` + "`" + ` must be an integer for the id of the user you want to import, for example: terraform import gitlab_user.example 42`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"gitlab_branch_protection":  0,
		"gitlab_deploy_key":         1,
		"gitlab_group":              2,
		"gitlab_group_membership":   3,
		"gitlab_group_variable":     4,
		"gitlab_label":              5,
		"gitlab_pipeline_schedule":  6,
		"gitlab_pipeline_trigger":   7,
		"gitlab_project":            8,
		"gitlab_project_cluster":    9,
		"gitlab_project_hook":       10,
		"gitlab_project_membership": 11,
		"gitlab_project_variable":   12,
		"gitlab_service_jira":       13,
		"gitlab_service_slack":      14,
		"gitlab_tag_protection":     15,
		"gitlab_user":               16,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
