package gitlab

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "gitlab_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) The ID of the group.`,
				},
				resource.Attribute{
					Name:        "full_path",
					Description: `(Optional) The full path of the group. >`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID assigned to the group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this group.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path of the group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the group.`,
				},
				resource.Attribute{
					Name:        "lfs_enabled",
					Description: `Boolean, is LFS enabled for projects in this group.`,
				},
				resource.Attribute{
					Name:        "request_access_enabled",
					Description: `Boolean, is request for access enabled to the group.`,
				},
				resource.Attribute{
					Name:        "visibility_level",
					Description: `Visibility level of the group. Possible values are ` + "`" + `private` + "`" + `, ` + "`" + `internal` + "`" + `, ` + "`" + `public` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Integer, ID of the parent group.`,
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
					Description: `The group level registration token to use during runner setup. [doc]: https://docs.gitlab.com/ee/api/groups.html#details-of-a-group`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID assigned to the group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this group.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path of the group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the group.`,
				},
				resource.Attribute{
					Name:        "lfs_enabled",
					Description: `Boolean, is LFS enabled for projects in this group.`,
				},
				resource.Attribute{
					Name:        "request_access_enabled",
					Description: `Boolean, is request for access enabled to the group.`,
				},
				resource.Attribute{
					Name:        "visibility_level",
					Description: `Visibility level of the group. Possible values are ` + "`" + `private` + "`" + `, ` + "`" + `internal` + "`" + `, ` + "`" + `public` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Integer, ID of the parent group.`,
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
					Description: `The group level registration token to use during runner setup. [doc]: https://docs.gitlab.com/ee/api/groups.html#details-of-a-group`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_group_membership",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) The ID of the group.`,
				},
				resource.Attribute{
					Name:        "full_path",
					Description: `(Optional) The full path of the group.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `(Optional) Only return members with the desired access level. Acceptable values are: ` + "`" + `guest` + "`" + `, ` + "`" + `reporter` + "`" + `, ` + "`" + `developer` + "`" + `, ` + "`" + `maintainer` + "`" + `, ` + "`" + `owner` + "`" + `. >`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `The list of group members.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the user by the gitlab server.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Whether the user is active or blocked.`,
				},
				resource.Attribute{
					Name:        "avatar_url",
					Description: `The avatar URL of the user.`,
				},
				resource.Attribute{
					Name:        "web_url",
					Description: `User's website URL.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `One of five levels of access to the group.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `Expiration date for the group membership.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "members",
					Description: `The list of group members.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the user by the gitlab server.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Whether the user is active or blocked.`,
				},
				resource.Attribute{
					Name:        "avatar_url",
					Description: `The avatar URL of the user.`,
				},
				resource.Attribute{
					Name:        "web_url",
					Description: `User's website URL.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `One of five levels of access to the group.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `Expiration date for the group membership.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_project",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The integer or path with namespace that uniquely identifies the project within the gitlab install. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path of the repository.`,
				},
				resource.Attribute{
					Name:        "path_with_namespace",
					Description: `The path of the repository with namespace.`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `The namespace (group or user) of the project. Defaults to your user. See [` + "`" + `gitlab_group` + "`" + `](../resources/group) for an example.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the project.`,
				},
				resource.Attribute{
					Name:        "default_branch",
					Description: `The default branch for the project.`,
				},
				resource.Attribute{
					Name:        "request_access_enabled",
					Description: `Allow users to request member access.`,
				},
				resource.Attribute{
					Name:        "issues_enabled",
					Description: `Enable issue tracking for the project.`,
				},
				resource.Attribute{
					Name:        "merge_requests_enabled",
					Description: `Enable merge requests for the project.`,
				},
				resource.Attribute{
					Name:        "pipelines_enabled",
					Description: `Enable pipelines for the project.`,
				},
				resource.Attribute{
					Name:        "wiki_enabled",
					Description: `Enable wiki for the project.`,
				},
				resource.Attribute{
					Name:        "snippets_enabled",
					Description: `Enable snippets for the project.`,
				},
				resource.Attribute{
					Name:        "lfs_enabled",
					Description: `Enable LFS for the project.`,
				},
				resource.Attribute{
					Name:        "visibility_level",
					Description: `Repositories are created as private by default.`,
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
					Name:        "archived",
					Description: `Whether the project is in read-only mode (archived).`,
				},
				resource.Attribute{
					Name:        "remove_source_branch_after_merge",
					Description: `Enable ` + "`" + `Delete source branch` + "`" + ` option by default for all new merge requests`,
				},
				resource.Attribute{
					Name:        "packages_enabled",
					Description: `Enable packages repository for the project.`,
				},
				resource.Attribute{
					Name:        "author_email_regex",
					Description: `All commit author emails must match this regex, e.g. ` + "`" + `@my-company.com$` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "branch_name_regex",
					Description: `All branch names must match this regex, e.g. ` + "`" + `(feature|hotfix)\/`,
				},
				resource.Attribute{
					Name:        "commit_message_regex",
					Description: `All commit messages must match this regex, e.g. ` + "`" + `Fixed \d+\..`,
				},
				resource.Attribute{
					Name:        "commit_message_negative_regex",
					Description: `No commit message is allowed to match this regex, for example ` + "`" + `ssh\:\/\/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_name_regex",
					Description: `All commited filenames must not match this regex, e.g. ` + "`" + `(jar|exe)$` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "commit_committer_check",
					Description: `Users can only push commits to this repository that were committed with one of their own verified emails.`,
				},
				resource.Attribute{
					Name:        "deny_delete_tag",
					Description: `Deny deleting a tag.`,
				},
				resource.Attribute{
					Name:        "member_check",
					Description: `Restrict commits by author (email) to existing GitLab users.`,
				},
				resource.Attribute{
					Name:        "prevent_secrets",
					Description: `GitLab will reject any files that are likely to contain secrets.`,
				},
				resource.Attribute{
					Name:        "reject_unsigned_commits",
					Description: `Reject commit when it’s not signed through GPG.`,
				},
				resource.Attribute{
					Name:        "max_file_size",
					Description: `Maximum file size (MB).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `The path of the repository.`,
				},
				resource.Attribute{
					Name:        "path_with_namespace",
					Description: `The path of the repository with namespace.`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `The namespace (group or user) of the project. Defaults to your user. See [` + "`" + `gitlab_group` + "`" + `](../resources/group) for an example.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the project.`,
				},
				resource.Attribute{
					Name:        "default_branch",
					Description: `The default branch for the project.`,
				},
				resource.Attribute{
					Name:        "request_access_enabled",
					Description: `Allow users to request member access.`,
				},
				resource.Attribute{
					Name:        "issues_enabled",
					Description: `Enable issue tracking for the project.`,
				},
				resource.Attribute{
					Name:        "merge_requests_enabled",
					Description: `Enable merge requests for the project.`,
				},
				resource.Attribute{
					Name:        "pipelines_enabled",
					Description: `Enable pipelines for the project.`,
				},
				resource.Attribute{
					Name:        "wiki_enabled",
					Description: `Enable wiki for the project.`,
				},
				resource.Attribute{
					Name:        "snippets_enabled",
					Description: `Enable snippets for the project.`,
				},
				resource.Attribute{
					Name:        "lfs_enabled",
					Description: `Enable LFS for the project.`,
				},
				resource.Attribute{
					Name:        "visibility_level",
					Description: `Repositories are created as private by default.`,
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
					Name:        "archived",
					Description: `Whether the project is in read-only mode (archived).`,
				},
				resource.Attribute{
					Name:        "remove_source_branch_after_merge",
					Description: `Enable ` + "`" + `Delete source branch` + "`" + ` option by default for all new merge requests`,
				},
				resource.Attribute{
					Name:        "packages_enabled",
					Description: `Enable packages repository for the project.`,
				},
				resource.Attribute{
					Name:        "author_email_regex",
					Description: `All commit author emails must match this regex, e.g. ` + "`" + `@my-company.com$` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "branch_name_regex",
					Description: `All branch names must match this regex, e.g. ` + "`" + `(feature|hotfix)\/`,
				},
				resource.Attribute{
					Name:        "commit_message_regex",
					Description: `All commit messages must match this regex, e.g. ` + "`" + `Fixed \d+\..`,
				},
				resource.Attribute{
					Name:        "commit_message_negative_regex",
					Description: `No commit message is allowed to match this regex, for example ` + "`" + `ssh\:\/\/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_name_regex",
					Description: `All commited filenames must not match this regex, e.g. ` + "`" + `(jar|exe)$` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "commit_committer_check",
					Description: `Users can only push commits to this repository that were committed with one of their own verified emails.`,
				},
				resource.Attribute{
					Name:        "deny_delete_tag",
					Description: `Deny deleting a tag.`,
				},
				resource.Attribute{
					Name:        "member_check",
					Description: `Restrict commits by author (email) to existing GitLab users.`,
				},
				resource.Attribute{
					Name:        "prevent_secrets",
					Description: `GitLab will reject any files that are likely to contain secrets.`,
				},
				resource.Attribute{
					Name:        "reject_unsigned_commits",
					Description: `Reject commit when it’s not signed through GPG.`,
				},
				resource.Attribute{
					Name:        "max_file_size",
					Description: `Maximum file size (MB).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_projects",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) The ID of the group owned by the authenticated user to look projects for within. Cannot be used with ` + "`" + `min_access_level` + "`" + `, ` + "`" + `with_programming_language` + "`" + ` or ` + "`" + `statistics` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "archived",
					Description: `(Optional) Limit by archived status.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) Limit by visibility ` + "`" + `public` + "`" + `, ` + "`" + `internal` + "`" + `, or ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "order_by",
					Description: `(Optional) Return projects ordered by ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `path` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `updated_at` + "`" + `, or ` + "`" + `last_activity_at` + "`" + ` fields. Default is ` + "`" + `created_at` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Return projects sorted in ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + ` order. Default is ` + "`" + `desc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "search",
					Description: `(Optional) Return list of authorized projects matching the search criteria.`,
				},
				resource.Attribute{
					Name:        "simple",
					Description: `(Optional) Return only the ID, URL, name, and path of each project.`,
				},
				resource.Attribute{
					Name:        "owned",
					Description: `(Optional) Limit by projects owned by the current user.`,
				},
				resource.Attribute{
					Name:        "starred",
					Description: `(Optional) Limit by projects starred by the current user.`,
				},
				resource.Attribute{
					Name:        "with_issues_enabled",
					Description: `(Optional) Limit by projects with issues feature enabled. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "with_merge_requests_enabled",
					Description: `(Optional) Limit by projects with merge requests feature enabled. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "with_shared",
					Description: `(Optional) Include projects shared to this group. Default is ` + "`" + `true` + "`" + `. Needs ` + "`" + `group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "include_subgroups",
					Description: `(Optional) Include projects in subgroups of this group. Default is ` + "`" + `false` + "`" + `. Needs ` + "`" + `group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_access_level",
					Description: `(Optional) Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with ` + "`" + `group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "with_custom_attributes",
					Description: `(Optional) Include custom attributes in response _(admins only)_.`,
				},
				resource.Attribute{
					Name:        "membership",
					Description: `(Optional) Limit by projects that the current user is a member of.`,
				},
				resource.Attribute{
					Name:        "statistics",
					Description: `(Optional) Include project statistics. Cannot be used with ` + "`" + `group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "with_programming_language",
					Description: `(Optional) Limit by projects which use the given programming language. Cannot be used with ` + "`" + `group_id` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "projects",
					Description: `A list containing the projects matching the supplied arguments Projects items have the following fields:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the project.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the project.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Whether the project is public.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `The visibility of the project.`,
				},
				resource.Attribute{
					Name:        "ssh_url_to_repo",
					Description: `The SSH clone URL of the project.`,
				},
				resource.Attribute{
					Name:        "http_url_to_repo",
					Description: `The HTTP clone URL of the project.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `A set of the project topics (formerly called "project tags").`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The owner of the project, due to Terraform aggregate types limitations, this field's attributes are accessed with the ` + "`" + `owner.0` + "`" + ` prefix. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "name_with_namespace",
					Description: `In ` + "`" + `group / subgroup / project` + "`" + ` or ` + "`" + `user / project` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "path_with_namespace",
					Description: `In ` + "`" + `group/subgroup/project` + "`" + ` or ` + "`" + `user/project` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "approvals_before_merge",
					Description: `The numbers of approvals needed in a merge requests.`,
				},
				resource.Attribute{
					Name:        "jobs_enabled",
					Description: `Whether pipelines are enabled for the project.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "projects",
					Description: `A list containing the projects matching the supplied arguments Projects items have the following fields:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the project.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the project.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Whether the project is public.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `The visibility of the project.`,
				},
				resource.Attribute{
					Name:        "ssh_url_to_repo",
					Description: `The SSH clone URL of the project.`,
				},
				resource.Attribute{
					Name:        "http_url_to_repo",
					Description: `The HTTP clone URL of the project.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `A set of the project topics (formerly called "project tags").`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The owner of the project, due to Terraform aggregate types limitations, this field's attributes are accessed with the ` + "`" + `owner.0` + "`" + ` prefix. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "name_with_namespace",
					Description: `In ` + "`" + `group / subgroup / project` + "`" + ` or ` + "`" + `user / project` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "path_with_namespace",
					Description: `In ` + "`" + `group/subgroup/project` + "`" + ` or ` + "`" + `user/project` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "approvals_before_merge",
					Description: `The numbers of approvals needed in a merge requests.`,
				},
				resource.Attribute{
					Name:        "jobs_enabled",
					Description: `Whether pipelines are enabled for the project.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) The e-mail address of the user. (Requires administrator privileges)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The username of the user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Optional) The ID of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the user by the gitlab server.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The e-mail address of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "is_admin",
					Description: `Whether the user is an admin.`,
				},
				resource.Attribute{
					Name:        "can_create_group",
					Description: `Whether the user can create groups.`,
				},
				resource.Attribute{
					Name:        "can_create_project",
					Description: `Whether the user can create projects.`,
				},
				resource.Attribute{
					Name:        "projects_limit",
					Description: `Number of projects the user can create.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date the user was created at.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Whether the user is active or blocked.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `Whether the user is external.`,
				},
				resource.Attribute{
					Name:        "extern_uid",
					Description: `The external UID of the user.`,
				},
				resource.Attribute{
					Name:        "user_provider",
					Description: `The UID provider of the user.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `The organization of the user.`,
				},
				resource.Attribute{
					Name:        "two_factor_enabled",
					Description: `Whether user's two factor auth is enabled.`,
				},
				resource.Attribute{
					Name:        "avatar_url",
					Description: `The avatar URL of the user.`,
				},
				resource.Attribute{
					Name:        "bio",
					Description: `The bio of the user.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the user.`,
				},
				resource.Attribute{
					Name:        "skype",
					Description: `Skype username of the user.`,
				},
				resource.Attribute{
					Name:        "linkedin",
					Description: `Linkedin profile of the user.`,
				},
				resource.Attribute{
					Name:        "twitter",
					Description: `Twitter username of the user.`,
				},
				resource.Attribute{
					Name:        "website_url",
					Description: `User's website URL.`,
				},
				resource.Attribute{
					Name:        "theme_id",
					Description: `User's theme ID.`,
				},
				resource.Attribute{
					Name:        "color_scheme_id",
					Description: `User's color scheme ID.`,
				},
				resource.Attribute{
					Name:        "last_sign_in_at",
					Description: `Last user's sign-in date.`,
				},
				resource.Attribute{
					Name:        "current_sign_in_at",
					Description: `Current user's sign-in date.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the user by the gitlab server.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The e-mail address of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "is_admin",
					Description: `Whether the user is an admin.`,
				},
				resource.Attribute{
					Name:        "can_create_group",
					Description: `Whether the user can create groups.`,
				},
				resource.Attribute{
					Name:        "can_create_project",
					Description: `Whether the user can create projects.`,
				},
				resource.Attribute{
					Name:        "projects_limit",
					Description: `Number of projects the user can create.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date the user was created at.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Whether the user is active or blocked.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `Whether the user is external.`,
				},
				resource.Attribute{
					Name:        "extern_uid",
					Description: `The external UID of the user.`,
				},
				resource.Attribute{
					Name:        "user_provider",
					Description: `The UID provider of the user.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `The organization of the user.`,
				},
				resource.Attribute{
					Name:        "two_factor_enabled",
					Description: `Whether user's two factor auth is enabled.`,
				},
				resource.Attribute{
					Name:        "avatar_url",
					Description: `The avatar URL of the user.`,
				},
				resource.Attribute{
					Name:        "bio",
					Description: `The bio of the user.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the user.`,
				},
				resource.Attribute{
					Name:        "skype",
					Description: `Skype username of the user.`,
				},
				resource.Attribute{
					Name:        "linkedin",
					Description: `Linkedin profile of the user.`,
				},
				resource.Attribute{
					Name:        "twitter",
					Description: `Twitter username of the user.`,
				},
				resource.Attribute{
					Name:        "website_url",
					Description: `User's website URL.`,
				},
				resource.Attribute{
					Name:        "theme_id",
					Description: `User's theme ID.`,
				},
				resource.Attribute{
					Name:        "color_scheme_id",
					Description: `User's color scheme ID.`,
				},
				resource.Attribute{
					Name:        "last_sign_in_at",
					Description: `Last user's sign-in date.`,
				},
				resource.Attribute{
					Name:        "current_sign_in_at",
					Description: `Current user's sign-in date.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_users",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "search",
					Description: `(Optional) Search users by username, name or email.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Filter users that are active.`,
				},
				resource.Attribute{
					Name:        "blocked",
					Description: `(Optional) Filter users that are blocked.`,
				},
				resource.Attribute{
					Name:        "order_by",
					Description: `(Optional) Order the users' list by ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `username` + "`" + `, ` + "`" + `created_at` + "`" + ` or ` + "`" + `updated_at` + "`" + `. (Requires administrator privileges)`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort users' list in asc or desc order. (Requires administrator privileges)`,
				},
				resource.Attribute{
					Name:        "extern_uid",
					Description: `(Optional) Lookup users by external UID. (Requires administrator privileges)`,
				},
				resource.Attribute{
					Name:        "extern_provider",
					Description: `(Optional) Lookup users by external provider. (Requires administrator privileges)`,
				},
				resource.Attribute{
					Name:        "created_before",
					Description: `(Optional) Search for users created before a specific date. (Requires administrator privileges)`,
				},
				resource.Attribute{
					Name:        "created_after",
					Description: `(Optional) Search for users created after a specific date. (Requires administrator privileges) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `The list of users.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the user by the gitlab server.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The e-mail address of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "is_admin",
					Description: `Whether the user is an admin.`,
				},
				resource.Attribute{
					Name:        "can_create_group",
					Description: `Whether the user can create groups.`,
				},
				resource.Attribute{
					Name:        "can_create_project",
					Description: `Whether the user can create projects.`,
				},
				resource.Attribute{
					Name:        "projects_limit",
					Description: `Number of projects the user can create.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date the user was created at.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Whether the user is active or blocked.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `Whether the user is external.`,
				},
				resource.Attribute{
					Name:        "extern_uid",
					Description: `The external UID of the user.`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `The UID provider of the user.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `The organization of the user.`,
				},
				resource.Attribute{
					Name:        "two_factor_enabled",
					Description: `Whether user's two-factor auth is enabled.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `The note associated to the user.`,
				},
				resource.Attribute{
					Name:        "avatar_url",
					Description: `The avatar URL of the user.`,
				},
				resource.Attribute{
					Name:        "bio",
					Description: `The bio of the user.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the user.`,
				},
				resource.Attribute{
					Name:        "skype",
					Description: `Skype username of the user.`,
				},
				resource.Attribute{
					Name:        "linkedin",
					Description: `LinkedIn profile of the user.`,
				},
				resource.Attribute{
					Name:        "twitter",
					Description: `Twitter username of the user.`,
				},
				resource.Attribute{
					Name:        "website_url",
					Description: `User's website URL.`,
				},
				resource.Attribute{
					Name:        "theme_id",
					Description: `User's theme ID.`,
				},
				resource.Attribute{
					Name:        "color_scheme_id",
					Description: `User's color scheme ID.`,
				},
				resource.Attribute{
					Name:        "last_sign_in_at",
					Description: `Last user's sign-in date.`,
				},
				resource.Attribute{
					Name:        "current_sign_in_at",
					Description: `Current user's sign-in date. [users_for_admins]: https://docs.gitlab.com/ce/api/users.html#for-admins`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "users",
					Description: `The list of users.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique id assigned to the user by the gitlab server.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The e-mail address of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "is_admin",
					Description: `Whether the user is an admin.`,
				},
				resource.Attribute{
					Name:        "can_create_group",
					Description: `Whether the user can create groups.`,
				},
				resource.Attribute{
					Name:        "can_create_project",
					Description: `Whether the user can create projects.`,
				},
				resource.Attribute{
					Name:        "projects_limit",
					Description: `Number of projects the user can create.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date the user was created at.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Whether the user is active or blocked.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `Whether the user is external.`,
				},
				resource.Attribute{
					Name:        "extern_uid",
					Description: `The external UID of the user.`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `The UID provider of the user.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `The organization of the user.`,
				},
				resource.Attribute{
					Name:        "two_factor_enabled",
					Description: `Whether user's two-factor auth is enabled.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `The note associated to the user.`,
				},
				resource.Attribute{
					Name:        "avatar_url",
					Description: `The avatar URL of the user.`,
				},
				resource.Attribute{
					Name:        "bio",
					Description: `The bio of the user.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the user.`,
				},
				resource.Attribute{
					Name:        "skype",
					Description: `Skype username of the user.`,
				},
				resource.Attribute{
					Name:        "linkedin",
					Description: `LinkedIn profile of the user.`,
				},
				resource.Attribute{
					Name:        "twitter",
					Description: `Twitter username of the user.`,
				},
				resource.Attribute{
					Name:        "website_url",
					Description: `User's website URL.`,
				},
				resource.Attribute{
					Name:        "theme_id",
					Description: `User's theme ID.`,
				},
				resource.Attribute{
					Name:        "color_scheme_id",
					Description: `User's color scheme ID.`,
				},
				resource.Attribute{
					Name:        "last_sign_in_at",
					Description: `Last user's sign-in date.`,
				},
				resource.Attribute{
					Name:        "current_sign_in_at",
					Description: `Current user's sign-in date. [users_for_admins]: https://docs.gitlab.com/ce/api/users.html#for-admins`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"gitlab_group":            0,
		"gitlab_group_membership": 1,
		"gitlab_project":          2,
		"gitlab_projects":         3,
		"gitlab_user":             4,
		"gitlab_users":            5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
