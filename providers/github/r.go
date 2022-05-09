package github

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "github_actions_organization_permissions",
			Category:         "Resources",
			ShortDescription: `Creates and manages Actions permissions within a GitHub organization`,
			Description:      ``,
			Keywords: []string{
				"actions",
				"organization",
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled_repositories_config",
					Description: `(Optional) Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when ` + "`" + `enabled_repositories` + "`" + ` = ` + "`" + `selected` + "`" + `. See [Enabled Repositories Config](#enabled-repositories-config) below for details. ### Allowed Actions Config The ` + "`" + `allowed_actions_config` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "github_owned_allowed",
					Description: `(Required) Whether GitHub-owned actions are allowed in the organization.`,
				},
				resource.Attribute{
					Name:        "patterns_allowed",
					Description: `(Optional) Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@`,
				},
				resource.Attribute{
					Name:        "verified_allowed",
					Description: `(Optional) Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators. ### Enabled Repositories Config The ` + "`" + `enabled_repositories_config` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "repository_ids",
					Description: `(Required) List of repository IDs to enable for GitHub Actions. ## Import This resource can be imported using the ID of the GitHub organization: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_actions_organization_permissions.test <github_organization_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_actions_organization_secret",
			Category:         "Resources",
			ShortDescription: `Creates and manages an Action Secret within a GitHub organization`,
			Description:      ``,
			Keywords: []string{
				"actions",
				"organization",
				"secret",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "selected_repository_ids",
					Description: `(Optional) An array of repository ids that can access the organization secret. ## Attributes Reference`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_actions_organization_secret_repositories",
			Category:         "Resources",
			ShortDescription: `Manages repository allow list for an Action Secret within a GitHub organization`,
			Description:      ``,
			Keywords: []string{
				"actions",
				"organization",
				"secret",
				"repositories",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "selected_repository_ids",
					Description: `(Required) An array of repository ids that can access the organization secret. ## Import This resource can be imported using an ID made up of the secret name: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_actions_organization_secret_repositories.test_secret_repos test_secret_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_actions_runner_group",
			Category:         "Resources",
			ShortDescription: `Creates and manages an Actions Runner Group within a GitHub organization`,
			Description:      ``,
			Keywords: []string{
				"actions",
				"runner",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "selected_repository_ids",
					Description: `(Optional) IDs of the repositories which should be added to the runner group`,
				},
				resource.Attribute{
					Name:        "allows_public_repositories",
					Description: `Whether public repositories can be added to the runner group`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allows_public_repositories",
					Description: `Whether public repositories can be added to the runner group`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_actions_secret",
			Category:         "Resources",
			ShortDescription: `Creates and manages an Action Secret within a GitHub repository`,
			Description:      ``,
			Keywords: []string{
				"actions",
				"secret",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "encrypted_value",
					Description: `(Optional) Encrypted value of the secret using the Github public key in Base64 format.`,
				},
				resource.Attribute{
					Name:        "plaintext_value",
					Description: `(Optional) Plaintext value of the secret to be encrypted ## Attributes Reference`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_app_installation_repository",
			Category:         "Resources",
			ShortDescription: `Manages the associations between app installations and repositories.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"installation",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "installation_id",
					Description: `(Required) The GitHub app installation id.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_branch",
			Category:         "Resources",
			ShortDescription: `Creates and manages branches within GitHub repositories.`,
			Description:      ``,
			Keywords: []string{
				"branch",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The GitHub repository name.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `(Required) The repository branch to create.`,
				},
				resource.Attribute{
					Name:        "source_branch",
					Description: `(Optional) The branch name to start from. Defaults to ` + "`" + `main` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_sha",
					Description: `(Optional) The commit hash to start from. Defaults to the tip of ` + "`" + `source_branch` + "`" + `. If provided, ` + "`" + `source_branch` + "`" + ` is ignored. ## Attribute Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "source_sha",
					Description: `A string storing the commit this branch was started from. Not populated when imported.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `An etag representing the Branch object.`,
				},
				resource.Attribute{
					Name:        "ref",
					Description: `A string representing a branch reference, in the form of ` + "`" + `refs/heads/<branch>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sha",
					Description: `A string storing the reference's ` + "`" + `HEAD` + "`" + ` commit's SHA1. ## Import GitHub Branch can be imported using an ID made up of ` + "`" + `repository:branch` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_branch.terraform terraform:main ` + "`" + `` + "`" + `` + "`" + ` Importing github branch into an instance object (when using a for each block to manage multiple branches) ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_branch.terraform["terraform"] terraform:main ` + "`" + `` + "`" + `` + "`" + ` Optionally, a source branch may be specified using an ID of ` + "`" + `repository:branch:source_branch` + "`" + `. This is useful for importing branches that do not branch directly off main. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_branch.terraform terraform:feature-branch:dev ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "source_sha",
					Description: `A string storing the commit this branch was started from. Not populated when imported.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `An etag representing the Branch object.`,
				},
				resource.Attribute{
					Name:        "ref",
					Description: `A string representing a branch reference, in the form of ` + "`" + `refs/heads/<branch>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sha",
					Description: `A string storing the reference's ` + "`" + `HEAD` + "`" + ` commit's SHA1. ## Import GitHub Branch can be imported using an ID made up of ` + "`" + `repository:branch` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_branch.terraform terraform:main ` + "`" + `` + "`" + `` + "`" + ` Importing github branch into an instance object (when using a for each block to manage multiple branches) ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_branch.terraform["terraform"] terraform:main ` + "`" + `` + "`" + `` + "`" + ` Optionally, a source branch may be specified using an ID of ` + "`" + `repository:branch:source_branch` + "`" + `. This is useful for importing branches that do not branch directly off main. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_branch.terraform terraform:feature-branch:dev ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_branch_protection",
			Category:         "Resources",
			ShortDescription: `Protects a GitHub branch.`,
			Description:      ``,
			Keywords: []string{
				"branch",
				"protection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository_id",
					Description: `(Required) The name or node ID of the repository associated with this branch protection rule.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Required) Identifies the protection rule pattern.`,
				},
				resource.Attribute{
					Name:        "enforce_admins",
					Description: `(Optional) Boolean, setting this to ` + "`" + `true` + "`" + ` enforces status checks for repository administrators.`,
				},
				resource.Attribute{
					Name:        "require_signed_commits",
					Description: `(Optional) Boolean, setting this to ` + "`" + `true` + "`" + ` requires all commits to be signed with GPG.`,
				},
				resource.Attribute{
					Name:        "required_linear_history",
					Description: `(Optional) Boolean, setting this to ` + "`" + `true` + "`" + ` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch`,
				},
				resource.Attribute{
					Name:        "require_conversation_resolution",
					Description: `(Optional) Boolean, setting this to ` + "`" + `true` + "`" + ` requires all conversations on code must be resolved before a pull request can be merged.`,
				},
				resource.Attribute{
					Name:        "required_status_checks",
					Description: `(Optional) Enforce restrictions for required status checks. See [Required Status Checks](#required-status-checks) below for details.`,
				},
				resource.Attribute{
					Name:        "required_pull_request_reviews",
					Description: `(Optional) Enforce restrictions for pull request reviews. See [Required Pull Request Reviews](#required-pull-request-reviews) below for details.`,
				},
				resource.Attribute{
					Name:        "push_restrictions",
					Description: `(Optional) The list of actor IDs that may push to the branch.`,
				},
				resource.Attribute{
					Name:        "allows_deletions",
					Description: `(Optional) Boolean, setting this to ` + "`" + `true` + "`" + ` to allow the branch to be deleted.`,
				},
				resource.Attribute{
					Name:        "allows_force_pushes",
					Description: `(Optional) Boolean, setting this to ` + "`" + `true` + "`" + ` to allow force pushes on the branch. ### Required Status Checks ` + "`" + `required_status_checks` + "`" + ` supports the following arguments:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_branch_protection_v3",
			Category:         "Resources",
			ShortDescription: `Protects a GitHub branch using the v3 / REST implementation. The ` + "`" + `github_branch_protection` + "`" + ` resource has moved to the GraphQL API, while this resource will continue to leverage the REST API`,
			Description:      ``,
			Keywords: []string{
				"branch",
				"protection",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The GitHub repository name.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `(Required) The Git branch to protect.`,
				},
				resource.Attribute{
					Name:        "enforce_admins",
					Description: `(Optional) Boolean, setting this to ` + "`" + `true` + "`" + ` enforces status checks for repository administrators.`,
				},
				resource.Attribute{
					Name:        "require_signed_commits",
					Description: `(Optional) Boolean, setting this to ` + "`" + `true` + "`" + ` requires all commits to be signed with GPG.`,
				},
				resource.Attribute{
					Name:        "require_conversation_resolution",
					Description: `(Optional) Boolean, setting this to ` + "`" + `true` + "`" + ` requires all conversations on code must be resolved before a pull request can be merged.`,
				},
				resource.Attribute{
					Name:        "required_status_checks",
					Description: `(Optional) Enforce restrictions for required status checks. See [Required Status Checks](#required-status-checks) below for details.`,
				},
				resource.Attribute{
					Name:        "required_pull_request_reviews",
					Description: `(Optional) Enforce restrictions for pull request reviews. See [Required Pull Request Reviews](#required-pull-request-reviews) below for details.`,
				},
				resource.Attribute{
					Name:        "restrictions",
					Description: `(Optional) Enforce restrictions for the users and teams that may push to the branch. See [Restrictions](#restrictions) below for details. ### Required Status Checks ` + "`" + `required_status_checks` + "`" + ` supports the following arguments:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_emu_group_mapping",
			Category:         "Resources",
			ShortDescription: `Manages mappings between external groups for enterprise managed users.`,
			Description:      ``,
			Keywords: []string{
				"emu",
				"group",
				"mapping",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team_slug",
					Description: `(Required) Slug of the GitHub team`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_issue",
			Category:         "Resources",
			ShortDescription: `Provides a GitHub issue resource.`,
			Description:      ``,
			Keywords: []string{
				"issue",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "owner",
					Description: `(Required) Owner of the repository the issue belongs to.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The GitHub repository name`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) Title of the issue.`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `(Optional) Title of the issue.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels to attach to the issue.`,
				},
				resource.Attribute{
					Name:        "assignees",
					Description: `(Optional) List of Logins to assign the to the issue.`,
				},
				resource.Attribute{
					Name:        "milestone_number",
					Description: `(Optional) Milestone number to assign to the issue ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "number",
					Description: `(Computed) - The issue number`,
				},
				resource.Attribute{
					Name:        "issue_id",
					Description: `(Computed) - The issue id ## Import GitHub Issues can be imported using an ID made up of ` + "`" + `repository:number` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_issue.issue_15 myrepo:15 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "number",
					Description: `(Computed) - The issue number`,
				},
				resource.Attribute{
					Name:        "issue_id",
					Description: `(Computed) - The issue id ## Import GitHub Issues can be imported using an ID made up of ` + "`" + `repository:number` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_issue.issue_15 myrepo:15 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_issue_label",
			Category:         "Resources",
			ShortDescription: `Provides a GitHub issue label resource.`,
			Description:      ``,
			Keywords: []string{
				"issue",
				"label",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The GitHub repository`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the label.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Required) A 6 character hex code,`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A short description of the label.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Computed) The URL to the issue label ## Import GitHub Issue Labels can be imported using an ID made up of ` + "`" + `repository:name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_issue_label.panic_label terraform:panic ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_membership",
			Category:         "Resources",
			ShortDescription: `Provides a GitHub membership resource.`,
			Description:      ``,
			Keywords: []string{
				"membership",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The user to add to the organization.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The role of the user within the organization. Must be one of ` + "`" + `member` + "`" + ` or ` + "`" + `admin` + "`" + `. Defaults to ` + "`" + `member` + "`" + `. ## Import GitHub Membership can be imported using an ID made up of ` + "`" + `organization:username` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_membership.member hashicorp:someuser ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_organization_block",
			Category:         "Resources",
			ShortDescription: `Creates and manages blocks for GitHub organizations`,
			Description:      ``,
			Keywords: []string{
				"organization",
				"block",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The name of the user to block.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_organization_project",
			Category:         "Resources",
			ShortDescription: `Creates and manages projects for GitHub organizations`,
			Description:      ``,
			Keywords: []string{
				"organization",
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project.`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `(Optional) The body of the project. ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL of the project`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "url",
					Description: `URL of the project`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_organization_webhook",
			Category:         "Resources",
			ShortDescription: `Creates and manages webhooks for GitHub organizations`,
			Description:      ``,
			Keywords: []string{
				"organization",
				"webhook",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "events",
					Description: `(Required) A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Required) key/value pair of configuration for this webhook. Available keys are ` + "`" + `url` + "`" + `, ` + "`" + `content_type` + "`" + `, ` + "`" + `secret` + "`" + ` and ` + "`" + `insecure_ssl` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Indicate of the webhook should receive events. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The type of the webhook. ` + "`" + `web` + "`" + ` is the default and the only option. ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL of the webhook ## Import Organization webhooks can be imported using the ` + "`" + `id` + "`" + ` of the webhook. The ` + "`" + `id` + "`" + ` of the webhook can be found in the URL of the webhook. For example, ` + "`" + `"https://github.com/organizations/foo-org/settings/hooks/123456789"` + "`" + `. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_organization_webhook.terraform 123456789 ` + "`" + `` + "`" + `` + "`" + ` If secret is populated in the webhook's configuration, the value will be imported as "`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "url",
					Description: `URL of the webhook ## Import Organization webhooks can be imported using the ` + "`" + `id` + "`" + ` of the webhook. The ` + "`" + `id` + "`" + ` of the webhook can be found in the URL of the webhook. For example, ` + "`" + `"https://github.com/organizations/foo-org/settings/hooks/123456789"` + "`" + `. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_organization_webhook.terraform 123456789 ` + "`" + `` + "`" + `` + "`" + ` If secret is populated in the webhook's configuration, the value will be imported as "`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_project_card",
			Category:         "Resources",
			ShortDescription: `Creates and manages project cards for GitHub projects`,
			Description:      ``,
			Keywords: []string{
				"project",
				"card",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "column_id",
					Description: `(Required) The ID of the card.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) The note contents of the card. Markdown supported.`,
				},
				resource.Attribute{
					Name:        "content_id",
					Description: `(Optional) [` + "`" + `github_issue.issue_id` + "`" + `](issue.html#argument-reference).`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) Must be either ` + "`" + `Issue` + "`" + ` or ` + "`" + `PullRequest` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_project_column",
			Category:         "Resources",
			ShortDescription: `Creates and manages project columns for GitHub projects`,
			Description:      ``,
			Keywords: []string{
				"project",
				"column",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of an existing project that the column will be created in.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the column.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_repository",
			Category:         "Resources",
			ShortDescription: `Creates and manages repositories within GitHub organizations or personal accounts`,
			Description:      ``,
			Keywords: []string{
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the repository.`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `(Optional) URL of a page describing the project.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to create a private repository. Repositories are created as public (e.g. open source) by default.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) Can be ` + "`" + `public` + "`" + ` or ` + "`" + `private` + "`" + `. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be ` + "`" + `internal` + "`" + `. The ` + "`" + `visibility` + "`" + ` parameter overrides the ` + "`" + `private` + "`" + ` parameter.`,
				},
				resource.Attribute{
					Name:        "has_issues",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable the GitHub Issues features on the repository.`,
				},
				resource.Attribute{
					Name:        "has_projects",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to ` + "`" + `false` + "`" + ` and will otherwise default to ` + "`" + `true` + "`" + `. If you specify ` + "`" + `true` + "`" + ` when it has been disabled it will return an error.`,
				},
				resource.Attribute{
					Name:        "has_wiki",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable the GitHub Wiki features on the repository.`,
				},
				resource.Attribute{
					Name:        "is_template",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to tell GitHub that this is a template repository.`,
				},
				resource.Attribute{
					Name:        "allow_merge_commit",
					Description: `(Optional) Set to ` + "`" + `false` + "`" + ` to disable merge commits on the repository.`,
				},
				resource.Attribute{
					Name:        "allow_squash_merge",
					Description: `(Optional) Set to ` + "`" + `false` + "`" + ` to disable squash merges on the repository.`,
				},
				resource.Attribute{
					Name:        "allow_rebase_merge",
					Description: `(Optional) Set to ` + "`" + `false` + "`" + ` to disable rebase merges on the repository.`,
				},
				resource.Attribute{
					Name:        "allow_auto_merge",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to allow auto-merging pull requests on the repository.`,
				},
				resource.Attribute{
					Name:        "delete_branch_on_merge",
					Description: `(Optional) Automatically delete head branch after a pull request is merged. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "has_downloads",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable the (deprecated) downloads features on the repository.`,
				},
				resource.Attribute{
					Name:        "auto_init",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to produce an initial commit in the repository.`,
				},
				resource.Attribute{
					Name:        "gitignore_template",
					Description: `(Optional) Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".`,
				},
				resource.Attribute{
					Name:        "license_template",
					Description: `(Optional) Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".`,
				},
				resource.Attribute{
					Name:        "default_branch",
					Description: `(Optional) (Deprecated: Use ` + "`" + `github_branch_default` + "`" + ` resource instead) The name of the default branch of the repository.`,
				},
				resource.Attribute{
					Name:        "archived",
					Description: `(Optional) Specifies if the repository should be archived. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "archive_on_destroy",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to archive the repository instead of deleting on destroy.`,
				},
				resource.Attribute{
					Name:        "pages",
					Description: `(Optional) The repository's GitHub Pages configuration. See [GitHub Pages Configuration](#github-pages-configuration) below for details.`,
				},
				resource.Attribute{
					Name:        "topics",
					Description: `(Optional) The list of topics of the repository.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) Use a template repository to create this resource. See [Template Repositories](#template-repositories) below for details.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The source branch and directory for the rendered Pages site. See [GitHub Pages Source](#github-pages-source) below for details.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `(Optional) The custom domain for the repository. This can only be set after the repository has been created. #### GitHub Pages Source #### The ` + "`" + `source` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `(Required) The repository branch used to publish the site's source files. (i.e. ` + "`" + `main` + "`" + ` or ` + "`" + `gh-pages` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The repository directory from which the site publishes (Default: ` + "`" + `/` + "`" + `). ### Template Repositories ` + "`" + `template` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `A string of the form "orgname/reponame".`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL to the repository on the web.`,
				},
				resource.Attribute{
					Name:        "ssh_clone_url",
					Description: `URL that can be provided to ` + "`" + `git clone` + "`" + ` to clone the repository via SSH.`,
				},
				resource.Attribute{
					Name:        "http_clone_url",
					Description: `URL that can be provided to ` + "`" + `git clone` + "`" + ` to clone the repository via HTTPS.`,
				},
				resource.Attribute{
					Name:        "git_clone_url",
					Description: `URL that can be provided to ` + "`" + `git clone` + "`" + ` to clone the repository anonymously via the git protocol.`,
				},
				resource.Attribute{
					Name:        "svn_url",
					Description: `URL that can be provided to ` + "`" + `svn checkout` + "`" + ` to check out the repository via GitHub's Subversion protocol emulation.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `GraphQL global node id for use with v4 API`,
				},
				resource.Attribute{
					Name:        "repo_id",
					Description: `GitHub ID for the repository`,
				},
				resource.Attribute{
					Name:        "pages",
					Description: `The block consisting of the repository's GitHub Pages configuration with the following additional attributes:`,
				},
				resource.Attribute{
					Name:        "custom_404",
					Description: `Whether the rendered GitHub Pages site has a custom 404 page.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `The absolute URL (including scheme) of the rendered GitHub Pages site e.g. ` + "`" + `https://username.github.io` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The GitHub Pages site's build status e.g. ` + "`" + `building` + "`" + ` or ` + "`" + `built` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "branches",
					Description: `The list of this repository's branches. Each element of ` + "`" + `branches` + "`" + ` has the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the branch.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `Whether the branch is protected. ## Import Repositories can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_repository.terraform terraform ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "full_name",
					Description: `A string of the form "orgname/reponame".`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL to the repository on the web.`,
				},
				resource.Attribute{
					Name:        "ssh_clone_url",
					Description: `URL that can be provided to ` + "`" + `git clone` + "`" + ` to clone the repository via SSH.`,
				},
				resource.Attribute{
					Name:        "http_clone_url",
					Description: `URL that can be provided to ` + "`" + `git clone` + "`" + ` to clone the repository via HTTPS.`,
				},
				resource.Attribute{
					Name:        "git_clone_url",
					Description: `URL that can be provided to ` + "`" + `git clone` + "`" + ` to clone the repository anonymously via the git protocol.`,
				},
				resource.Attribute{
					Name:        "svn_url",
					Description: `URL that can be provided to ` + "`" + `svn checkout` + "`" + ` to check out the repository via GitHub's Subversion protocol emulation.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `GraphQL global node id for use with v4 API`,
				},
				resource.Attribute{
					Name:        "repo_id",
					Description: `GitHub ID for the repository`,
				},
				resource.Attribute{
					Name:        "pages",
					Description: `The block consisting of the repository's GitHub Pages configuration with the following additional attributes:`,
				},
				resource.Attribute{
					Name:        "custom_404",
					Description: `Whether the rendered GitHub Pages site has a custom 404 page.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `The absolute URL (including scheme) of the rendered GitHub Pages site e.g. ` + "`" + `https://username.github.io` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The GitHub Pages site's build status e.g. ` + "`" + `building` + "`" + ` or ` + "`" + `built` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "branches",
					Description: `The list of this repository's branches. Each element of ` + "`" + `branches` + "`" + ` has the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the branch.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `Whether the branch is protected. ## Import Repositories can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_repository.terraform terraform ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_repository_autolink_reference",
			Category:         "Resources",
			ShortDescription: `Creates and manages autolink references for a single repository`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"autolink",
				"reference",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The repository of the autolink reference.`,
				},
				resource.Attribute{
					Name:        "key_prefix",
					Description: `(Required) This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.`,
				},
				resource.Attribute{
					Name:        "target_url_template",
					Description: `(Required) The template of the target URL used for the links; must be a valid URL and contain ` + "`" + `<num>` + "`" + ` for the reference number ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `An etag representing the autolink reference object. ## Import Autolink references can be imported using the ` + "`" + `name` + "`" + ` of the repository, combined with the ` + "`" + `id` + "`" + ` of the autolink reference and a ` + "`" + `/` + "`" + ` character for separating components, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import github_repository_autolink_reference.auto oof/123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `An etag representing the autolink reference object. ## Import Autolink references can be imported using the ` + "`" + `name` + "`" + ` of the repository, combined with the ` + "`" + `id` + "`" + ` of the autolink reference and a ` + "`" + `/` + "`" + ` character for separating components, e.g. ` + "`" + `` + "`" + `` + "`" + `sh terraform import github_repository_autolink_reference.auto oof/123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_repository_collaborator",
			Category:         "Resources",
			ShortDescription: `Provides a GitHub repository collaborator resource.`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"collaborator",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The GitHub repository`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The user to add to the repository as a collaborator.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Optional) The permission of the outside collaborator for the repository. Must be one of ` + "`" + `pull` + "`" + `, ` + "`" + `push` + "`" + `, ` + "`" + `maintain` + "`" + `, ` + "`" + `triage` + "`" + ` or ` + "`" + `admin` + "`" + ` for organization-owned repositories. Must be ` + "`" + `push` + "`" + ` for personal repositories. Defaults to ` + "`" + `push` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "permission_diff_suppression",
					Description: `(Optional) Suppress plan diffs for ` + "`" + `triage` + "`" + ` and ` + "`" + `maintain` + "`" + `. Defaults to ` + "`" + `false` + "`" + `. ## Attribute Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "invitation_id",
					Description: `ID of the invitation to be used in [` + "`" + `github_user_invitation_accepter` + "`" + `](./user_invitation_accepter.html) ## Import GitHub Repository Collaborators can be imported using an ID made up of ` + "`" + `repository:username` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_repository_collaborator.collaborator terraform:someuser ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "invitation_id",
					Description: `ID of the invitation to be used in [` + "`" + `github_user_invitation_accepter` + "`" + `](./user_invitation_accepter.html) ## Import GitHub Repository Collaborators can be imported using an ID made up of ` + "`" + `repository:username` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_repository_collaborator.collaborator terraform:someuser ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_repository_deploy_key",
			Category:         "Resources",
			ShortDescription: `Provides a GitHub repository deploy key resource.`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"deploy",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) A SSH key.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Required) A boolean qualifying the key to be either read only or read/write.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) Name of the GitHub repository.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) A title. Changing any of the fields forces re-creating the resource. ## Import Repository deploy keys can be imported using a colon-separated pair of repository name and GitHub's key id. The latter can be obtained by GitHub's SDKs and API. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_repository_deploy_key.foo test-repo:23824728 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_repository_environment",
			Category:         "Resources",
			ShortDescription: `Creates and manages environments for GitHub repositories`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"environment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "environment",
					Description: `(Required) The name of the environment.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The repository of the environment.`,
				},
				resource.Attribute{
					Name:        "wait_timer",
					Description: `(Optional) Amount of time to delay a job after the job is initially triggered. ### Reviewers The ` + "`" + `reviewers` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `(Optional) Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed. #### Deployment Branch Policy #### The ` + "`" + `deployment_branch_policy` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "protected_branches",
					Description: `(Required) Whether only branches with branch protection rules can deploy to this environment.`,
				},
				resource.Attribute{
					Name:        "custom_branch_policies",
					Description: `(Required) Whether only branches that match the specified name patterns can deploy to this environment. ## Import GitHub Repository Environment can be imported using an ID made up of ` + "`" + `name` + "`" + ` of the repository combined with the ` + "`" + `environment` + "`" + ` name of the environment, separated by a ` + "`" + `:` + "`" + ` character, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_repository_environment.daily terraform:daily ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_repository_file",
			Category:         "Resources",
			ShortDescription: `Creates and manages files within a GitHub repository`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"file",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The repository to create the file in.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `(Required) The path of the file to manage.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) The file content.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `(Optional) Git branch (defaults to ` + "`" + `main` + "`" + `). The branch must already exist, it will not be created if it does not already exist.`,
				},
				resource.Attribute{
					Name:        "commit_author",
					Description: `(Optional) Committer author name to use.`,
				},
				resource.Attribute{
					Name:        "commit_email",
					Description: `(Optional) Committer email address to use.`,
				},
				resource.Attribute{
					Name:        "commit_message",
					Description: `(Optional) Commit message when adding or updating the managed file.`,
				},
				resource.Attribute{
					Name:        "overwrite_on_create",
					Description: `(Optional) Enable overwriting existing files ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "commit_sha",
					Description: `The SHA of the commit that modified the file.`,
				},
				resource.Attribute{
					Name:        "sha",
					Description: `The SHA blob of the file. ## Import Repository files can be imported using a combination of the ` + "`" + `repo` + "`" + ` and ` + "`" + `file` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_repository_file.gitignore example/.gitignore ` + "`" + `` + "`" + `` + "`" + ` To import a file from a branch other than main, append ` + "`" + `:` + "`" + ` and the branch name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_repository_file.gitignore example/.gitignore:dev ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "commit_sha",
					Description: `The SHA of the commit that modified the file.`,
				},
				resource.Attribute{
					Name:        "sha",
					Description: `The SHA blob of the file. ## Import Repository files can be imported using a combination of the ` + "`" + `repo` + "`" + ` and ` + "`" + `file` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_repository_file.gitignore example/.gitignore ` + "`" + `` + "`" + `` + "`" + ` To import a file from a branch other than main, append ` + "`" + `:` + "`" + ` and the branch name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_repository_file.gitignore example/.gitignore:dev ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_repository_milestone",
			Category:         "Resources",
			ShortDescription: `Provides a GitHub repository milestone resource.`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"milestone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "owner",
					Description: `(Required) The owner of the GitHub Repository.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The name of the GitHub Repository.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) The title of the milestone.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the milestone.`,
				},
				resource.Attribute{
					Name:        "due_date",
					Description: `(Optional) The milestone due date. In ` + "`" + `yyyy-mm-dd` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The state of the milestone. Either ` + "`" + `open` + "`" + ` or ` + "`" + `closed` + "`" + `. Default: ` + "`" + `open` + "`" + ` ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "number",
					Description: `The number of the milestone. ## Import A GitHub Repository Milestone can be imported using an ID made up of ` + "`" + `owner/repository/number` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_repository_milestone.example example-owner/example-repository/1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "number",
					Description: `The number of the milestone. ## Import A GitHub Repository Milestone can be imported using an ID made up of ` + "`" + `owner/repository/number` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_repository_milestone.example example-owner/example-repository/1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_repository_project",
			Category:         "Resources",
			ShortDescription: `Creates and manages projects for GitHub repositories`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The repository of the project.`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `(Optional) The body of the project. ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL of the project`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "url",
					Description: `URL of the project`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_repository_webhook",
			Category:         "Resources",
			ShortDescription: `Creates and manages repository webhooks within GitHub organizations or personal accounts`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"webhook",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The repository of the webhook.`,
				},
				resource.Attribute{
					Name:        "events",
					Description: `(Required) A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/).`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Required) Configuration block for the webhook. [Detailed below.](#configuration)`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Indicate if the webhook should receive events. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The type of the webhook. ` + "`" + `web` + "`" + ` is the default and the only option. ### configuration`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL of the webhook.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Required) The content type for the payload. Valid values are either ` + "`" + `form` + "`" + ` or ` + "`" + `json` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) The shared secret for the webhook. [See API documentation](https://developer.github.com/v3/repos/hooks/#create-a-hook).`,
				},
				resource.Attribute{
					Name:        "insecure_ssl",
					Description: `(Optional) Insecure SSL boolean toggle. Defaults to ` + "`" + `false` + "`" + `. ## Attributes Reference The following additional attributes are exported:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL of the webhook. This is a sensitive attribute because it may include basic auth credentials. ## Import Repository webhooks can be imported using the ` + "`" + `name` + "`" + ` of the repository, combined with the ` + "`" + `id` + "`" + ` of the webhook, separated by a ` + "`" + `/` + "`" + ` character. The ` + "`" + `id` + "`" + ` of the webhook can be found in the URL of the webhook. For example: ` + "`" + `"https://github.com/foo-org/foo-repo/settings/hooks/14711452"` + "`" + `. Importing uses the name of the repository, as well as the ID of the webhook, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_repository_webhook.terraform terraform/11235813 ` + "`" + `` + "`" + `` + "`" + ` If secret is populated in the webhook's configuration, the value will be imported as "`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "url",
					Description: `URL of the webhook. This is a sensitive attribute because it may include basic auth credentials. ## Import Repository webhooks can be imported using the ` + "`" + `name` + "`" + ` of the repository, combined with the ` + "`" + `id` + "`" + ` of the webhook, separated by a ` + "`" + `/` + "`" + ` character. The ` + "`" + `id` + "`" + ` of the webhook can be found in the URL of the webhook. For example: ` + "`" + `"https://github.com/foo-org/foo-repo/settings/hooks/14711452"` + "`" + `. Importing uses the name of the repository, as well as the ID of the webhook, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_repository_webhook.terraform terraform/11235813 ` + "`" + `` + "`" + `` + "`" + ` If secret is populated in the webhook's configuration, the value will be imported as "`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_team",
			Category:         "Resources",
			ShortDescription: `Provides a GitHub team resource.`,
			Description:      ``,
			Keywords: []string{
				"team",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the team.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the team.`,
				},
				resource.Attribute{
					Name:        "privacy",
					Description: `(Optional) The level of privacy for the team. Must be one of ` + "`" + `secret` + "`" + ` or ` + "`" + `closed` + "`" + `. Defaults to ` + "`" + `secret` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parent_team_id",
					Description: `(Optional) The ID of the parent team, if this is a nested team.`,
				},
				resource.Attribute{
					Name:        "ldap_dn",
					Description: `(Optional) The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.`,
				},
				resource.Attribute{
					Name:        "create_default_maintainer",
					Description: `(Optional) Adds a default maintainer to the team. Defaults to ` + "`" + `false` + "`" + ` and adds the creating user to the team when ` + "`" + `true` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the created team.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `The Node ID of the created team.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug of the created team, which may or may not differ from ` + "`" + `name` + "`" + `, depending on whether ` + "`" + `name` + "`" + ` contains "URL-unsafe" characters. Useful when referencing the team in [` + "`" + `github_branch_protection` + "`" + `](/docs/providers/github/r/branch_protection.html). ## Import GitHub Teams can be imported using the GitHub team ID e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_team.core 1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the created team.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `The Node ID of the created team.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug of the created team, which may or may not differ from ` + "`" + `name` + "`" + `, depending on whether ` + "`" + `name` + "`" + ` contains "URL-unsafe" characters. Useful when referencing the team in [` + "`" + `github_branch_protection` + "`" + `](/docs/providers/github/r/branch_protection.html). ## Import GitHub Teams can be imported using the GitHub team ID e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_team.core 1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_team_members",
			Category:         "Resources",
			ShortDescription: `Provides an authoritative GitHub team members resource.`,
			Description:      ``,
			Keywords: []string{
				"team",
				"members",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) The GitHub team id`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) List of team members. See [Members](#members) below for details. ### Members ` + "`" + `members` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The user to add to the team.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The role of the user within the team. Must be one of ` + "`" + `member` + "`" + ` or ` + "`" + `maintainer` + "`" + `. Defaults to ` + "`" + `member` + "`" + `. ## Import GitHub Team Membership can be imported using the team ID ` + "`" + `teamid` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_team_member.some_team 1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_team_membership",
			Category:         "Resources",
			ShortDescription: `Provides a GitHub team membership resource.`,
			Description:      ``,
			Keywords: []string{
				"team",
				"membership",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) The GitHub team id`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The user to add to the team.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The role of the user within the team. Must be one of ` + "`" + `member` + "`" + ` or ` + "`" + `maintainer` + "`" + `. Defaults to ` + "`" + `member` + "`" + `. ## Import GitHub Team Membership can be imported using an ID made up of ` + "`" + `teamid:username` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_team_membership.member 1234567:someuser ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_team_repository",
			Category:         "Resources",
			ShortDescription: `Manages the associations between teams and repositories.`,
			Description:      ``,
			Keywords: []string{
				"team",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) The GitHub team id or the GitHub team slug`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The repository to add to the team.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Optional) The permissions of team members regarding the repository. Must be one of ` + "`" + `pull` + "`" + `, ` + "`" + `triage` + "`" + `, ` + "`" + `push` + "`" + `, ` + "`" + `maintain` + "`" + `, or ` + "`" + `admin` + "`" + `. Defaults to ` + "`" + `pull` + "`" + `. ## Import GitHub Team Repository can be imported using an ID made up of ` + "`" + `teamid:repository` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_team_repository.terraform_repo 1234567:terraform ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_team_sync_group_mapping",
			Category:         "Resources",
			ShortDescription: `Creates and manages the connections between a team and its IdP group(s).`,
			Description:      ``,
			Keywords: []string{
				"team",
				"sync",
				"group",
				"mapping",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `The ID of the IdP group.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the IdP group.`,
				},
				resource.Attribute{
					Name:        "group_description",
					Description: `The description of the IdP group. ## Import GitHub Team Sync Group Mappings can be imported using the GitHub team ` + "`" + `slug` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_team_sync_group_mapping.example some_team ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_user_gpg_key",
			Category:         "Resources",
			ShortDescription: `Provides a GitHub user's GPG key resource.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"gpg",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "armored_public_key",
					Description: `(Required) Your public GPG key, generated in ASCII-armored format. See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The GitHub ID of the GPG key, e.g. ` + "`" + `401586` + "`" + ``,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `The key ID of the GPG key, e.g. ` + "`" + `3262EFF25BA0D270` + "`" + ` ## Import GPG keys are not importable due to the fact that [API](https://developer.github.com/v3/users/gpg_keys/#gpg-keys) does not return previously uploaded GPG key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The GitHub ID of the GPG key, e.g. ` + "`" + `401586` + "`" + ``,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `The key ID of the GPG key, e.g. ` + "`" + `3262EFF25BA0D270` + "`" + ` ## Import GPG keys are not importable due to the fact that [API](https://developer.github.com/v3/users/gpg_keys/#gpg-keys) does not return previously uploaded GPG key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_user_invitation_accepter",
			Category:         "Resources",
			ShortDescription: `Provides a resource to manage GitHub repository collaborator invitations.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"invitation",
				"accepter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "invitation_id",
					Description: `(Required) ID of the invitation to accept`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_user_ssh_key",
			Category:         "Resources",
			ShortDescription: `Provides a GitHub user's SSH key resource.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"ssh",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "title",
					Description: `(Required) A descriptive name for the new key. e.g. ` + "`" + `Personal MacBook Air` + "`" + ``,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The public SSH key to add to your GitHub account. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSH key`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL of the SSH key ## Import SSH keys can be imported using their ID e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_user_ssh_key.example 1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSH key`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL of the SSH key ## Import SSH keys can be imported using their ID e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import github_user_ssh_key.example 1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"github_actions_organization_permissions":         0,
		"github_actions_organization_secret":              1,
		"github_actions_organization_secret_repositories": 2,
		"github_actions_runner_group":                     3,
		"github_actions_secret":                           4,
		"github_app_installation_repository":              5,
		"github_branch":                                   6,
		"github_branch_protection":                        7,
		"github_branch_protection_v3":                     8,
		"github_emu_group_mapping":                        9,
		"github_issue":                                    10,
		"github_issue_label":                              11,
		"github_membership":                               12,
		"github_organization_block":                       13,
		"github_organization_project":                     14,
		"github_organization_webhook":                     15,
		"github_project_card":                             16,
		"github_project_column":                           17,
		"github_repository":                               18,
		"github_repository_autolink_reference":            19,
		"github_repository_collaborator":                  20,
		"github_repository_deploy_key":                    21,
		"github_repository_environment":                   22,
		"github_repository_file":                          23,
		"github_repository_milestone":                     24,
		"github_repository_project":                       25,
		"github_repository_webhook":                       26,
		"github_team":                                     27,
		"github_team_members":                             28,
		"github_team_membership":                          29,
		"github_team_repository":                          30,
		"github_team_sync_group_mapping":                  31,
		"github_user_gpg_key":                             32,
		"github_user_invitation_accepter":                 33,
		"github_user_ssh_key":                             34,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
