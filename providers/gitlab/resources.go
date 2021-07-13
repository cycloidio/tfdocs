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
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Optional) One of five levels of access to the project. Valid values are: ` + "`" + `no one` + "`" + `, ` + "`" + `developer` + "`" + `, ` + "`" + `maintainer` + "`" + `, ` + "`" + `admin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "merge_access_level",
					Description: `(Optional) One of five levels of access to the project. Valid values are: ` + "`" + `no one` + "`" + `, ` + "`" + `developer` + "`" + `, ` + "`" + `maintainer` + "`" + `, ` + "`" + `admin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with ` + "`" + `group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with ` + "`" + `user_id` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_deploy_key",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "gitlab_deploy_key_enable",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, string) The name or id of the project to add the deploy key to.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Required, string) The Gitlab key id for the pre-existing deploy key ## Import GitLab enabled deploy keys can be imported using an id made up of ` + "`" + `{project_id}:{deploy_key_id}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_deploy_key_enable.example 12345:67890 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_deploy_token",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, string) The name or id of the project to add the deploy token to. Either ` + "`" + `project` + "`" + ` or ` + "`" + `group` + "`" + ` must be set.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Required, string) The name or id of the group to add the deploy token to. Either ` + "`" + `project` + "`" + ` or ` + "`" + `group` + "`" + ` must be set.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) A name to describe the deploy token with.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional, string) A username for the deploy token. Default is ` + "`" + `gitlab+deploy-token-{n}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `(Optional, string) Time the token will expire it, RFC3339 format. Will not expire per default.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Required, set of strings) Valid values: ` + "`" + `read_repository` + "`" + `, ` + "`" + `read_registry` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The secret token. This is only populated when creating a new deploy token.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "token",
					Description: `The secret token. This is only populated when creating a new deploy token.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Optional) The group's visibility. Can be ` + "`" + `private` + "`" + `, ` + "`" + `internal` + "`" + `, or ` + "`" + `public` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "share_with_group_lock",
					Description: `(Optional) Boolean, defaults to false. Prevent sharing a project with another group within this group.`,
				},
				resource.Attribute{
					Name:        "project_creation_level",
					Description: `(Optional), defaults to Maintainer. Determine if developers can create projects in the group. Can be noone (No one), maintainer (Maintainers), or developer (Developers + Maintainers).`,
				},
				resource.Attribute{
					Name:        "auto_devops_enabled",
					Description: `(Optional) Boolean, defaults to false. Default to Auto DevOps pipeline for all projects within this group.`,
				},
				resource.Attribute{
					Name:        "emails_disabled",
					Description: `(Optional) Boolean, defaults to false. Disable email notifications`,
				},
				resource.Attribute{
					Name:        "mentions_disabled",
					Description: `(Optional) Boolean, defaults to false. Disable the capability of a group from getting mentioned`,
				},
				resource.Attribute{
					Name:        "subgroup_creation_level",
					Description: `(Optional), defaults to Owner. Allowed to create subgroups. Can be owner (Owners), or maintainer (Maintainers).`,
				},
				resource.Attribute{
					Name:        "require_two_factor_authentication",
					Description: `(Optional) Boolean, defaults to false. equire all users in this group to setup Two-factor authentication.`,
				},
				resource.Attribute{
					Name:        "two_factor_grace_period",
					Description: `(Optional) Int, defaults to 48. Time before Two-factor authentication is enforced (in hours).`,
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
					Description: `Web URL of the group.`,
				},
				resource.Attribute{
					Name:        "runners_token",
					Description: `The group level registration token to use during runner setup. ## Importing groups You can import a group state using ` + "`" + `terraform import <resource> <id>` + "`" + `. The ` + "`" + `id` + "`" + ` can be whatever the [details of a group][details_of_a_group] api takes for its ` + "`" + `:id` + "`" + ` value, so for example: terraform import gitlab_group.example example [details_of_a_group]: https://docs.gitlab.com/ee/api/groups.html#details-of-a-group`,
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
					Description: `Web URL of the group.`,
				},
				resource.Attribute{
					Name:        "runners_token",
					Description: `The group level registration token to use during runner setup. ## Importing groups You can import a group state using ` + "`" + `terraform import <resource> <id>` + "`" + `. The ` + "`" + `id` + "`" + ` can be whatever the [details of a group][details_of_a_group] api takes for its ` + "`" + `:id` + "`" + ` value, so for example: terraform import gitlab_group.example example [details_of_a_group]: https://docs.gitlab.com/ee/api/groups.html#details-of-a-group`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_group_cluster",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group",
					Description: `(Required, string) The id of the group to add the cluster to.`,
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
					Name:        "kubernetes_authorization_type",
					Description: `(Optional, string) The cluster authorization type. Valid values are ` + "`" + `rbac` + "`" + `, ` + "`" + `abac` + "`" + `, ` + "`" + `unknown_authorization` + "`" + `. Defaults to ` + "`" + `rbac` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "environment_scope",
					Description: `(Optional, string) The associated environment to the cluster. Defaults to ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "management_project_id",
					Description: `(Optional, string) The ID of the management project for the cluster. ## Import GitLab group clusters can be imported using an id made up of ` + "`" + `groupid:clusterid` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_group_cluster.bar 123:321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_group_label",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group",
					Description: `(Required) The name or id of the group to add the label to.`,
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
					Description: `The unique id assigned to the label by the GitLab server (the name of the label). ## Import Gitlab group labels can be imported using an id made up of ` + "`" + `{group_id}:{group_label_id}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_group_label.example 12345:fixme ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the label by the GitLab server (the name of the label). ## Import Gitlab group labels can be imported using an id made up of ` + "`" + `{group_id}:{group_label_id}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_group_label.example 12345:fixme ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_group_ldap_link",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The id of the GitLab group.`,
				},
				resource.Attribute{
					Name:        "cn",
					Description: `(Required) The CN of the LDAP group to link with.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `(Required) Acceptable values are: guest, reporter, developer, maintainer, owner.`,
				},
				resource.Attribute{
					Name:        "ldap_provider",
					Description: `(Required) The name of the LDAP provider as stored in the GitLab database. ## Import GitLab group ldap links can be imported using an id made up of ` + "`" + `ldap_provider:cn` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_group_ldap_link.test "ldapmain:testuser" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_group_membership",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Required) Acceptable values are: guest, reporter, developer, maintainer, owner.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `(Optional) Expiration date for the group membership. Format: ` + "`" + `YYYY-MM-DD` + "`" + ` ## Import GitLab group membership can be imported using an id made up of ` + "`" + `group_id:user_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_group_membership.test "12345:1337" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_group_share_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The id of the main group.`,
				},
				resource.Attribute{
					Name:        "share_group_id",
					Description: `(Required) The id of an additional group which will be shared with the main group.`,
				},
				resource.Attribute{
					Name:        "group_access",
					Description: `(Required) One of five levels of access to the group.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `(Optional) Share expiration date. Format: ` + "`" + `YYYY-MM-DD` + "`" + ` ## Import GitLab group shares can be imported using an id made up of ` + "`" + `mainGroupId:shareGroupId` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_group_share_group.test 12345:1337 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_group_variable",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Name:        "variable_type",
					Description: `(Optional, string) The type of a variable. Available types are: env_var (default) and file.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `(Optional, boolean) If set to ` + "`" + `true` + "`" + `, the variable will be passed only to pipelines running on protected branches and tags. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "masked",
					Description: `(Optional, boolean) If set to ` + "`" + `true` + "`" + `, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variables). Defaults to ` + "`" + `false` + "`" + `. ## Import GitLab group variables can be imported using an id made up of ` + "`" + `groupid:variablename` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_group_variable.example 12345:group_variable_key ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_instance_cluster",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_instance_variable",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required, string) The name of the variable.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required, string) The value of the variable.`,
				},
				resource.Attribute{
					Name:        "variable_type",
					Description: `(Optional, string) The type of a variable. Available types are: env_var (default) and file.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `(Optional, boolean) If set to ` + "`" + `true` + "`" + `, the variable will be passed only to pipelines running on protected branches and tags. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "masked",
					Description: `(Optional, boolean) If set to ` + "`" + `true` + "`" + `, the value of the variable will be hidden in job logs. The value must meet the [masking requirements](https://docs.gitlab.com/ee/ci/variables/#masked-variable-requirements). Defaults to ` + "`" + `false` + "`" + `. ## Import GitLab instance variables can be imported using an id made up of ` + "`" + `variablename` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `console $ terraform import gitlab_instance_variable.example instance_variable_key ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_label",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Optional, bool) The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially. ## Import GitLab pipeline schedules can be imported using an id made up of ` + "`" + `{project_id}:{pipeline_schedule_id}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_pipeline_schedule.test 1:3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_pipeline_schedule_variable",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, string) The id of the project to add the schedule to.`,
				},
				resource.Attribute{
					Name:        "pipeline_schedule_id",
					Description: `(Required, string) The id of the pipeline schedule.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required, string) Name of the variable.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required, string) Value of the variable.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_pipeline_trigger",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, string) The name or id of the project to add the trigger to.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required, string) The description of the pipeline trigger. ## Import GitLab pipeline triggers can be imported using an id made up of ` + "`" + `{project_id}:{pipeline_trigger_id}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_pipeline_trigger.test 1:3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Name:        "import_url",
					Description: `(Optional) Git URL to a repository to be imported.`,
				},
				resource.Attribute{
					Name:        "request_access_enabled",
					Description: `Allow users to request member access.`,
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
					Name:        "pipelines_enabled",
					Description: `(Optional) Enable pipelines for the project.`,
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
					Name:        "lfs_enabled",
					Description: `(Optional) Enable LFS for the project.`,
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
					Name:        "archived",
					Description: `(Optional) Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.`,
				},
				resource.Attribute{
					Name:        "initialize_with_readme",
					Description: `(Optional) Create main branch with first commit containing a README.md file.`,
				},
				resource.Attribute{
					Name:        "packages_enabled",
					Description: `(Optional) Enable packages repository for the project.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) When used without use_custom_template, name of a built-in project template. When used with use_custom_template, name of a custom project template. This option is mutually exclusive with ` + "`" + `template_project_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template_project_id",
					Description: `(Optional) When used with use_custom_template, project ID of a custom project template. This is preferable to using template_name since template_name may be ambiguous (enterprise edition). This option is mutually exclusive with ` + "`" + `template_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_custom_template",
					Description: `(Optional) Use either custom instance or group (with group_with_project_templates_id) project template (enterprise edition).`,
				},
				resource.Attribute{
					Name:        "group_with_project_templates_id",
					Description: `(Optional) For group-level custom templates, specifies ID of group from which all the custom project templates are sourced. Leave empty for instance-level templates. Requires use_custom_template to be true (enterprise edition).`,
				},
				resource.Attribute{
					Name:        "pages_access_level",
					Description: `(Optional) Enable pages access control Valid values are ` + "`" + `disabled` + "`" + `, ` + "`" + `private` + "`" + `, ` + "`" + `enabled` + "`" + `, ` + "`" + `public` + "`" + `. ` + "`" + `private` + "`" + ` is the default. ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Integer that uniquely identifies the project within the gitlab install.`,
				},
				resource.Attribute{
					Name:        "path_with_namespace",
					Description: `The path of the repository with namespace.`,
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
					Name:        "remove_source_branch_after_merge",
					Description: `Enable ` + "`" + `Delete source branch` + "`" + ` option by default for all new merge requests. ## Nested Blocks ### push_rules For information on push rules, consult the [GitLab documentation](https://docs.gitlab.com/ce/push_rules/push_rules.html#push-rules). #### Arguments`,
				},
				resource.Attribute{
					Name:        "author_email_regex",
					Description: `(Optional) All commit author emails must match this regex, e.g. ` + "`" + `@my-company.com$` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "branch_name_regex",
					Description: `(Optional) All branch names must match this regex, e.g. ` + "`" + `(feature|hotfix)\/`,
				},
				resource.Attribute{
					Name:        "commit_message_regex",
					Description: `(Optional) All commit messages must match this regex, e.g. ` + "`" + `Fixed \d+\..`,
				},
				resource.Attribute{
					Name:        "commit_message_negative_regex",
					Description: `(Optional) No commit message is allowed to match this regex, for example ` + "`" + `ssh\:\/\/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_name_regex",
					Description: `(Optional) All commited filenames must not match this regex, e.g. ` + "`" + `(jar|exe)$` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "commit_committer_check",
					Description: `(Optional, bool) Users can only push commits to this repository that were committed with one of their own verified emails.`,
				},
				resource.Attribute{
					Name:        "deny_delete_tag",
					Description: `(Optional, bool) Deny deleting a tag.`,
				},
				resource.Attribute{
					Name:        "member_check",
					Description: `(Optional, bool) Restrict commits by author (email) to existing GitLab users.`,
				},
				resource.Attribute{
					Name:        "prevent_secrets",
					Description: `(Optional, bool) GitLab will reject any files that are likely to contain secrets.`,
				},
				resource.Attribute{
					Name:        "reject_unsigned_commits",
					Description: `(Optional, bool) Reject commit when it’s not signed through GPG.`,
				},
				resource.Attribute{
					Name:        "max_file_size",
					Description: `(Optional, int) Maximum file size (MB). ## Importing projects You can import a project state using ` + "`" + `terraform import <resource> <id>` + "`" + `. The ` + "`" + `id` + "`" + ` can be whatever the [get single project api][get_single_project] takes for its ` + "`" + `:id` + "`" + ` value, so for example: terraform import gitlab_project.example richardc/example [get_single_project]: https://docs.gitlab.com/ee/api/projects.html#get-single-project [group_members_permissions]: https://docs.gitlab.com/ce/user/permissions.html#group-members-permissions`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Integer that uniquely identifies the project within the gitlab install.`,
				},
				resource.Attribute{
					Name:        "path_with_namespace",
					Description: `The path of the repository with namespace.`,
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
					Name:        "remove_source_branch_after_merge",
					Description: `Enable ` + "`" + `Delete source branch` + "`" + ` option by default for all new merge requests. ## Nested Blocks ### push_rules For information on push rules, consult the [GitLab documentation](https://docs.gitlab.com/ce/push_rules/push_rules.html#push-rules). #### Arguments`,
				},
				resource.Attribute{
					Name:        "author_email_regex",
					Description: `(Optional) All commit author emails must match this regex, e.g. ` + "`" + `@my-company.com$` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "branch_name_regex",
					Description: `(Optional) All branch names must match this regex, e.g. ` + "`" + `(feature|hotfix)\/`,
				},
				resource.Attribute{
					Name:        "commit_message_regex",
					Description: `(Optional) All commit messages must match this regex, e.g. ` + "`" + `Fixed \d+\..`,
				},
				resource.Attribute{
					Name:        "commit_message_negative_regex",
					Description: `(Optional) No commit message is allowed to match this regex, for example ` + "`" + `ssh\:\/\/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_name_regex",
					Description: `(Optional) All commited filenames must not match this regex, e.g. ` + "`" + `(jar|exe)$` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "commit_committer_check",
					Description: `(Optional, bool) Users can only push commits to this repository that were committed with one of their own verified emails.`,
				},
				resource.Attribute{
					Name:        "deny_delete_tag",
					Description: `(Optional, bool) Deny deleting a tag.`,
				},
				resource.Attribute{
					Name:        "member_check",
					Description: `(Optional, bool) Restrict commits by author (email) to existing GitLab users.`,
				},
				resource.Attribute{
					Name:        "prevent_secrets",
					Description: `(Optional, bool) GitLab will reject any files that are likely to contain secrets.`,
				},
				resource.Attribute{
					Name:        "reject_unsigned_commits",
					Description: `(Optional, bool) Reject commit when it’s not signed through GPG.`,
				},
				resource.Attribute{
					Name:        "max_file_size",
					Description: `(Optional, int) Maximum file size (MB). ## Importing projects You can import a project state using ` + "`" + `terraform import <resource> <id>` + "`" + `. The ` + "`" + `id` + "`" + ` can be whatever the [get single project api][get_single_project] takes for its ` + "`" + `:id` + "`" + ` value, so for example: terraform import gitlab_project.example richardc/example [get_single_project]: https://docs.gitlab.com/ee/api/projects.html#get-single-project [group_members_permissions]: https://docs.gitlab.com/ce/user/permissions.html#group-members-permissions`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project_approval_rule",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, string) The name or id of the project to add the approval rules.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the approval rule.`,
				},
				resource.Attribute{
					Name:        "approvals_required",
					Description: `(Required) The number of approvals required for this rule.`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `(Optional) A list of specific User IDs to add to the list of approvers.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `(Optional) A list of group IDs who's members can approve of the merge request ## Import GitLab project approval rules can be imported using an id consisting of ` + "`" + `project-id:rule-id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_project_approval_rule.example "12345:6" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project_cluster",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
				resource.Attribute{
					Name:        "management_project_id",
					Description: `(Optional, string) The ID of the management project for the cluster. ## Import GitLab project clusters can be imported using an id made up of ` + "`" + `projectid:clusterid` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_project_cluster.bar 123:321 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project_freeze_period",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required, string) The id of the project to add the schedule to.`,
				},
				resource.Attribute{
					Name:        "freeze_start",
					Description: `(Required,string) Start of the Freeze Period in cron format (e.g. ` + "`" + `0 1`,
				},
				resource.Attribute{
					Name:        "freeze_end",
					Description: `(Required, string) End of the Freeze Period in cron format (e.g. ` + "`" + `0 2`,
				},
				resource.Attribute{
					Name:        "cron_timezone",
					Description: `(Optional, string) The timezone. ## Import GitLab project freeze periods can be imported using an id made up of ` + "`" + `project_id:freeze_period_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_project_freeze_period.schedule "12345:1337" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project_hook",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Name:        "push_events_branch_filter",
					Description: `(Optional) Invoke the hook for push events on matching branches only.`,
				},
				resource.Attribute{
					Name:        "issues_events",
					Description: `(Optional) Invoke the hook for issues events.`,
				},
				resource.Attribute{
					Name:        "confidential_issues_events",
					Description: `(Optional) Invoke the hook for confidential issues events.`,
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
					Name:        "confidential_note_events",
					Description: `(Optional) Invoke the hook for confidential notes events.`,
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
					Description: `(Optional) Invoke the hook for wiki page events.`,
				},
				resource.Attribute{
					Name:        "deployment_events",
					Description: `(Optional) Invoke the hook for deployment events. ## Attributes Reference The resource exports the following attributes:`,
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
			Type:             "gitlab_project_level_mr_approvals",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project to change MR approval configuration.`,
				},
				resource.Attribute{
					Name:        "reset_approvals_on_push",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_overriding_approvers_per_merge_request",
					Description: `(Optional) By default, users are able to edit the approval rules in merge requests. If set to true, the approval rules for all new merge requests will be determined by the default approval rules. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "merge_requests_author_approval",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` if you want to allow merge request authors to self-approve merge requests. Authors also need to be included in the approvers list in order to be able to approve their merge request. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "merge_requests_disable_committers_approval",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` if you want to prevent approval of merge requests by merge request committers. Default is ` + "`" + `false` + "`" + `. ## Importing approval configuration You can import an approval configuration state using ` + "`" + `terraform import <resource> <project_id>` + "`" + `. For example: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import gitlab_project_level_mr_approvals.foo 53 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project_membership",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Required) One of five levels of access to the project. ## Import GitLab group membership can be imported using an id made up of ` + "`" + `group_id:user_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_project_membership.test "12345:1337" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project_mirror",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required) The id of the project.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL of the remote repository to be mirrored.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Determines if the mirror is enabled.`,
				},
				resource.Attribute{
					Name:        "only_protected_branches",
					Description: `Determines if only protected branches are mirrored.`,
				},
				resource.Attribute{
					Name:        "keep_divergent_refs",
					Description: `Determines if divergent refs are skipped. ## Import GitLab project mirror can be imported using an id made up of ` + "`" + `project_id:mirror_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_project_mirror.foo "12345:1337" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project_share_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the project.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The id of the group.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `(Required) One of five levels of access to the project. ## Import GitLab project group shares can be imported using an id made up of ` + "`" + `projectid:groupid` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gitlab_project_share_group.test 12345:1337 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project_variable",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Name:        "variable_type",
					Description: `(Optional, string) The type of a variable. Available types are: env_var (default) and file.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `(Optional, boolean) If set to ` + "`" + `true` + "`" + `, the variable will be passed only to pipelines running on protected branches and tags. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "masked",
					Description: `(Optional, boolean) If set to ` + "`" + `true` + "`" + `, the variable will be masked if it would have been written to the logs. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "environment_scope",
					Description: `(Optional, string) The environment_scope of the variable. Defaults to ` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_service_github",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required) ID of the project you want to activate integration on.`,
				},
				resource.Attribute{
					Name:        "repository_url",
					Description: `(Required) The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required) A GitHub personal access token with at least ` + "`" + `repo:status` + "`" + ` scope.`,
				},
				resource.Attribute{
					Name:        "static_context",
					Description: `(Optional) Append instance name instead of branch to the status. Must enable to set a GitLab status check as _required_ in GitHub. See [Static / dynamic status check names] to learn more. ## Importing GitHub service You can import a service_github state using ` + "`" + `terraform import <resource> <project_id>` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import gitlab_service_github.github 1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_service_jira",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Optional) The ID of a transition that moves issues to a closed state. You can find this number under the JIRA workflow administration (Administration > Issues > Workflows) by selecting View under Operations of the desired workflow of your project. By default, this ID is set to 2.`,
				},
				resource.Attribute{
					Name:        "commit_events",
					Description: `(Optional) Enable notifications for commit events`,
				},
				resource.Attribute{
					Name:        "merge_requests_events",
					Description: `(Optional) Enable notifications for merge request events`,
				},
				resource.Attribute{
					Name:        "comment_on_event_enabled",
					Description: `(Optional) Enable comments inside Jira issues on each GitLab event (commit / merge request) ## Importing Jira service You can import a service_jira state using ` + "`" + `terraform import <resource> <project_id>` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import gitlab_service_jira.jira 1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_service_pipelines_email",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required, string) ID of the project you want to activate integration on.`,
				},
				resource.Attribute{
					Name:        "recipients",
					Description: `(Required, set(string)) email addresses where notifications are sent.`,
				},
				resource.Attribute{
					Name:        "notify_only_broken_pipelines",
					Description: `(Optional, bool) Notify only broken pipelines. Default is true.`,
				},
				resource.Attribute{
					Name:        "branches_to_be_notified",
					Description: `(Optional, string) Branches to send notifications for. Valid options are ` + "`" + `all` + "`" + `, ` + "`" + `default` + "`" + `, ` + "`" + `protected` + "`" + `, and ` + "`" + `default_and_protected` + "`" + `. Default is ` + "`" + `default` + "`" + ` ## Importing Pipelines email service You can import a service_pipelines_email state using ` + "`" + `terraform import <resource> <project_id>` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import gitlab_service_pipelines_email.email 1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_service_slack",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Optional) DEPRECATED: This parameter has been replaced with ` + "`" + `branches_to_be_notified` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "branches_to_be_notified",
					Description: `(Optional) Branches to send notifications for. Valid options are "all", "default", "protected", and "default_and_protected".`,
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
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Name:        "email",
					Description: `(Required) The e-mail address of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password of the user.`,
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
					Description: `(Optional) Boolean, defaults to false. Whether a user has access only to some internal or private projects. External users can only access projects to which they are explicitly granted access.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) The note associated to the user.`,
				},
				resource.Attribute{
					Name:        "reset_password",
					Description: `(Optional) Boolean, defaults to false. Send user password reset link. ## Attributes Reference The resource exports the following attributes:`,
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

		"gitlab_branch_protection":          0,
		"gitlab_deploy_key":                 1,
		"gitlab_deploy_key_enable":          2,
		"gitlab_deploy_token":               3,
		"gitlab_group":                      4,
		"gitlab_group_cluster":              5,
		"gitlab_group_label":                6,
		"gitlab_group_ldap_link":            7,
		"gitlab_group_membership":           8,
		"gitlab_group_share_group":          9,
		"gitlab_group_variable":             10,
		"gitlab_instance_cluster":           11,
		"gitlab_instance_variable":          12,
		"gitlab_label":                      13,
		"gitlab_pipeline_schedule":          14,
		"gitlab_pipeline_schedule_variable": 15,
		"gitlab_pipeline_trigger":           16,
		"gitlab_project":                    17,
		"gitlab_project_approval_rule":      18,
		"gitlab_project_cluster":            19,
		"gitlab_project_freeze_period":      20,
		"gitlab_project_hook":               21,
		"gitlab_project_level_mr_approvals": 22,
		"gitlab_project_membership":         23,
		"gitlab_project_mirror":             24,
		"gitlab_project_share_group":        25,
		"gitlab_project_variable":           26,
		"gitlab_service_github":             27,
		"gitlab_service_jira":               28,
		"gitlab_service_pipelines_email":    29,
		"gitlab_service_slack":              30,
		"gitlab_tag_protection":             31,
		"gitlab_user":                       32,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
