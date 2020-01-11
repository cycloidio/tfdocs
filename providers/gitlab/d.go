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
			ShortDescription: `Looks up a gitlab group`,
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
			Type:             "gitlab_project",
			Category:         "Data Sources",
			ShortDescription: `View information about a project`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The integer that uniquely identifies the project within the gitlab install. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path of the repository.`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `The namespace (group or user) of the project. Defaults to your user. See [` + "`" + `gitlab_group` + "`" + `](../r/group.html) for an example.`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `The path of the repository.`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `The namespace (group or user) of the project. Defaults to your user. See [` + "`" + `gitlab_group` + "`" + `](../r/group.html) for an example.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gitlab_user",
			Category:         "Data Sources",
			ShortDescription: `Looks up a gitlab user`,
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
			ShortDescription: `Looks up gitlab users`,
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
					Description: `Current user's sign-in date. [users_for_admins]: https://docs.gitlab.com/ce/api/users.html#for-admins`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"gitlab_group":   0,
		"gitlab_project": 1,
		"gitlab_user":    2,
		"gitlab_users":   3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
