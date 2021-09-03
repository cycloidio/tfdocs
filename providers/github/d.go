package github

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "github_actions_public_key",
			Category:         "Data Sources",
			ShortDescription: `Get information on a GitHub Actions Public Key.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) Name of the repository to get public key from. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `ID of the key that has been retrieved.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_id",
					Description: `ID of the key that has been retrieved.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_branch",
			Category:         "Data Sources",
			ShortDescription: `Get information about a repository branch.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The GitHub repository name.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `(Required) The repository branch to create. ## Attribute Reference The following additional attributes are exported:`,
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
					Description: `A string storing the reference's ` + "`" + `HEAD` + "`" + ` commit's SHA1.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `A string storing the reference's ` + "`" + `HEAD` + "`" + ` commit's SHA1.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_collaborators",
			Category:         "Data Sources",
			ShortDescription: `Get the collaborators for a given repository.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "collaborator",
					Description: `An Array of GitHub collaborators. Each ` + "`" + `collaborator` + "`" + ` block consists of the fields documented below. ___ The ` + "`" + `collaborator` + "`" + ` block consists of:`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `The collaborator's login.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the collaborator.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The GitHub API URL for the collaborator.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `The GitHub HTML URL for the collaborator.`,
				},
				resource.Attribute{
					Name:        "followers_url",
					Description: `The GitHub API URL for the collaborator's followers.`,
				},
				resource.Attribute{
					Name:        "following_url",
					Description: `The GitHub API URL for those following the collaborator.`,
				},
				resource.Attribute{
					Name:        "gists_url",
					Description: `The GitHub API URL for the collaborator's gists.`,
				},
				resource.Attribute{
					Name:        "starred_url",
					Description: `The GitHub API URL for the collaborator's starred repositories.`,
				},
				resource.Attribute{
					Name:        "subscriptions_url",
					Description: `The GitHub API URL for the collaborator's subscribed repositories.`,
				},
				resource.Attribute{
					Name:        "organizations_url",
					Description: `The GitHub API URL for the collaborator's organizations.`,
				},
				resource.Attribute{
					Name:        "repos_url",
					Description: `The GitHub API URL for the collaborator's repositories.`,
				},
				resource.Attribute{
					Name:        "events_url",
					Description: `The GitHub API URL for the collaborator's events.`,
				},
				resource.Attribute{
					Name:        "received_events_url",
					Description: `The GitHub API URL for the collaborator's received events.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the collaborator (ex. ` + "`" + `user` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "site_admin",
					Description: `Whether the user is a GitHub admin.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `The permission of the collaborator.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: `Get information on GitHub's IP addresses.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "actions",
					Description: `An array of IP addresses in CIDR format specifying the addresses that incoming requests from GitHub actions will originate from.`,
				},
				resource.Attribute{
					Name:        "dependabot",
					Description: `An array of IP addresses in CIDR format specifying the A records for dependabot.`,
				},
				resource.Attribute{
					Name:        "hooks",
					Description: `An Array of IP addresses in CIDR format specifying the addresses that incoming service hooks will originate from.`,
				},
				resource.Attribute{
					Name:        "git",
					Description: `An Array of IP addresses in CIDR format specifying the Git servers.`,
				},
				resource.Attribute{
					Name:        "pages",
					Description: `An Array of IP addresses in CIDR format specifying the A records for GitHub Pages.`,
				},
				resource.Attribute{
					Name:        "importer",
					Description: `An Array of IP addresses in CIDR format specifying the A records for GitHub Importer.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_membership",
			Category:         "Data Sources",
			ShortDescription: `Get information on user membership in an organization.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username to lookup in the organization.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Optional) The organization to check for the above username. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `` + "`" + `admin` + "`" + ` or ` + "`" + `member` + "`" + ` -- the role the user has within the organization.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `An etag representing the membership object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `The username.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `` + "`" + `admin` + "`" + ` or ` + "`" + `member` + "`" + ` -- the role the user has within the organization.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `An etag representing the membership object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_organization",
			Category:         "Data Sources",
			ShortDescription: `Get an organization.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the organization account`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `The login of the organization account`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description the organization account`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The plan name for the organization account`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(` + "`" + `list` + "`" + `) A list with the repositories on the organization`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(` + "`" + `list` + "`" + `) A list with the members of the organization`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_organization_team_sync_groups",
			Category:         "Data Sources",
			ShortDescription: `Get the external identity provider (IdP) groups for an organization.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "groups",
					Description: `An Array of GitHub Identity Provider Groups. Each ` + "`" + `group` + "`" + ` block consists of the fields documented below. ___ The ` + "`" + `group` + "`" + ` block consists of:`,
				},
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
					Description: `The description of the IdP group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_organization_teams",
			Category:         "Data Sources",
			ShortDescription: `Get information on all GitHub teams of an organization.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "root_teams_only",
					Description: `Only return teams that are at the organization's root, i.e. no nested teams. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `An Array of GitHub Teams. Each ` + "`" + `team` + "`" + ` block consists of the fields documented below. ___ The ` + "`" + `team` + "`" + ` block consists of:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `the ID of the team.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `the Node ID of the team.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `the slug of the team.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the team's full name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the team's description.`,
				},
				resource.Attribute{
					Name:        "privacy",
					Description: `the team's privacy type.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `List of team members.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `List of team repositories.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_release",
			Category:         "Data Sources",
			ShortDescription: `Get information on a GitHub release.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "release_tag",
					Description: `Tag of release`,
				},
				resource.Attribute{
					Name:        "release_id",
					Description: `ID of release`,
				},
				resource.Attribute{
					Name:        "target_commitish",
					Description: `Commitish value that determines where the Git release is created from`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of release`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `Contents of the description (body) of a release`,
				},
				resource.Attribute{
					Name:        "draft",
					Description: `(` + "`" + `Boolean` + "`" + `) indicates whether the release is a draft`,
				},
				resource.Attribute{
					Name:        "prerelease",
					Description: `(` + "`" + `Boolean` + "`" + `) indicates whether the release is a prerelease`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of release creation`,
				},
				resource.Attribute{
					Name:        "published_at",
					Description: `Date of release publishing`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Base URL of the release`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL directing to detailed information on the release`,
				},
				resource.Attribute{
					Name:        "asserts_url",
					Description: `URL of any associated assets with the release`,
				},
				resource.Attribute{
					Name:        "upload_url",
					Description: `URL that can be used to upload Assets to the release`,
				},
				resource.Attribute{
					Name:        "zipball_url",
					Description: `Download URL of a specific release in ` + "`" + `zip` + "`" + ` format`,
				},
				resource.Attribute{
					Name:        "tarball_url",
					Description: `Download URL of a specific release in ` + "`" + `tar.gz` + "`" + ` format`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "release_tag",
					Description: `Tag of release`,
				},
				resource.Attribute{
					Name:        "release_id",
					Description: `ID of release`,
				},
				resource.Attribute{
					Name:        "target_commitish",
					Description: `Commitish value that determines where the Git release is created from`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of release`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `Contents of the description (body) of a release`,
				},
				resource.Attribute{
					Name:        "draft",
					Description: `(` + "`" + `Boolean` + "`" + `) indicates whether the release is a draft`,
				},
				resource.Attribute{
					Name:        "prerelease",
					Description: `(` + "`" + `Boolean` + "`" + `) indicates whether the release is a prerelease`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of release creation`,
				},
				resource.Attribute{
					Name:        "published_at",
					Description: `Date of release publishing`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Base URL of the release`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL directing to detailed information on the release`,
				},
				resource.Attribute{
					Name:        "asserts_url",
					Description: `URL of any associated assets with the release`,
				},
				resource.Attribute{
					Name:        "upload_url",
					Description: `URL that can be used to upload Assets to the release`,
				},
				resource.Attribute{
					Name:        "zipball_url",
					Description: `Download URL of a specific release in ` + "`" + `zip` + "`" + ` format`,
				},
				resource.Attribute{
					Name:        "tarball_url",
					Description: `Download URL of a specific release in ` + "`" + `tar.gz` + "`" + ` format`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_repositories",
			Category:         "Data Sources",
			ShortDescription: `Search for GitHub repositories`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "query",
					Description: `(Required) Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sorts the repositories returned by the specified attribute. Valid values include ` + "`" + `stars` + "`" + `, ` + "`" + `fork` + "`" + `, and ` + "`" + `updated` + "`" + `. Defaults to ` + "`" + `updated` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "full_names",
					Description: `A list of full names of found repositories (e.g. ` + "`" + `hashicorp/terraform` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of found repository names (e.g. ` + "`" + `terraform` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "full_names",
					Description: `A list of full names of found repositories (e.g. ` + "`" + `hashicorp/terraform` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of found repository names (e.g. ` + "`" + `terraform` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_repository",
			Category:         "Data Sources",
			ShortDescription: `Get details about GitHub repository`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the repository.`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `(Optional) Full name of the repository (in ` + "`" + `org/name` + "`" + ` format). ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `the Node ID of the repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the repository.`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `URL of a page describing the project.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `Whether the repository is private.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `Whether the repository is public, private or internal.`,
				},
				resource.Attribute{
					Name:        "has_issues",
					Description: `Whether the repository has GitHub Issues enabled.`,
				},
				resource.Attribute{
					Name:        "has_projects",
					Description: `Whether the repository has the GitHub Projects enabled.`,
				},
				resource.Attribute{
					Name:        "has_wiki",
					Description: `Whether the repository has the GitHub Wiki enabled.`,
				},
				resource.Attribute{
					Name:        "allow_merge_commit",
					Description: `Whether the repository allows merge commits.`,
				},
				resource.Attribute{
					Name:        "allow_squash_merge",
					Description: `Whether the repository allows squash merges.`,
				},
				resource.Attribute{
					Name:        "allow_rebase_merge",
					Description: `Whether the repository allows rebase merges.`,
				},
				resource.Attribute{
					Name:        "has_downloads",
					Description: `Whether the repository has Downloads feature enabled.`,
				},
				resource.Attribute{
					Name:        "default_branch",
					Description: `The name of the default branch of the repository.`,
				},
				resource.Attribute{
					Name:        "archived",
					Description: `Whether the repository is archived.`,
				},
				resource.Attribute{
					Name:        "pages",
					Description: `The repository's GitHub Pages configuration.`,
				},
				resource.Attribute{
					Name:        "topics",
					Description: `The list of topics of the repository.`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "node_id",
					Description: `the Node ID of the repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the repository.`,
				},
				resource.Attribute{
					Name:        "homepage_url",
					Description: `URL of a page describing the project.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `Whether the repository is private.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `Whether the repository is public, private or internal.`,
				},
				resource.Attribute{
					Name:        "has_issues",
					Description: `Whether the repository has GitHub Issues enabled.`,
				},
				resource.Attribute{
					Name:        "has_projects",
					Description: `Whether the repository has the GitHub Projects enabled.`,
				},
				resource.Attribute{
					Name:        "has_wiki",
					Description: `Whether the repository has the GitHub Wiki enabled.`,
				},
				resource.Attribute{
					Name:        "allow_merge_commit",
					Description: `Whether the repository allows merge commits.`,
				},
				resource.Attribute{
					Name:        "allow_squash_merge",
					Description: `Whether the repository allows squash merges.`,
				},
				resource.Attribute{
					Name:        "allow_rebase_merge",
					Description: `Whether the repository allows rebase merges.`,
				},
				resource.Attribute{
					Name:        "has_downloads",
					Description: `Whether the repository has Downloads feature enabled.`,
				},
				resource.Attribute{
					Name:        "default_branch",
					Description: `The name of the default branch of the repository.`,
				},
				resource.Attribute{
					Name:        "archived",
					Description: `Whether the repository is archived.`,
				},
				resource.Attribute{
					Name:        "pages",
					Description: `The repository's GitHub Pages configuration.`,
				},
				resource.Attribute{
					Name:        "topics",
					Description: `The list of topics of the repository.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_repository_milestone",
			Category:         "Data Sources",
			ShortDescription: `Get information on a GitHub Repository Milestone.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the milestone.`,
				},
				resource.Attribute{
					Name:        "due_date",
					Description: `The milestone due date (in ISO-8601 ` + "`" + `yyyy-mm-dd` + "`" + ` format).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the milestone.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Title of the milestone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of the milestone.`,
				},
				resource.Attribute{
					Name:        "due_date",
					Description: `The milestone due date (in ISO-8601 ` + "`" + `yyyy-mm-dd` + "`" + ` format).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the milestone.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Title of the milestone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_team",
			Category:         "Data Sources",
			ShortDescription: `Get information on a GitHub team.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The team slug. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `the ID of the team.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `the Node ID of the team.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the team's full name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the team's description.`,
				},
				resource.Attribute{
					Name:        "privacy",
					Description: `the team's privacy type.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `the team's permission level.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `List of team members`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `List of team repositories`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `the ID of the team.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `the Node ID of the team.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the team's full name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the team's description.`,
				},
				resource.Attribute{
					Name:        "privacy",
					Description: `the team's privacy type.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `the team's permission level.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `List of team members`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `List of team repositories`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "github_user",
			Category:         "Data Sources",
			ShortDescription: `Get information on a GitHub user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username. Use an empty string ` + "`" + `""` + "`" + ` to retrieve information about the currently authenticated user. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `the Node ID of the user.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `the user's login.`,
				},
				resource.Attribute{
					Name:        "avatar_url",
					Description: `the user's avatar URL.`,
				},
				resource.Attribute{
					Name:        "gravatar_id",
					Description: `the user's gravatar ID.`,
				},
				resource.Attribute{
					Name:        "site_admin",
					Description: `whether the user is a GitHub admin.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the user's full name.`,
				},
				resource.Attribute{
					Name:        "company",
					Description: `the user's company name.`,
				},
				resource.Attribute{
					Name:        "blog",
					Description: `the user's blog location.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `the user's location.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `the user's email.`,
				},
				resource.Attribute{
					Name:        "gpg_keys",
					Description: `list of user's GPG keys.`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `list of user's SSH keys.`,
				},
				resource.Attribute{
					Name:        "bio",
					Description: `the user's bio.`,
				},
				resource.Attribute{
					Name:        "public_repos",
					Description: `the number of public repositories.`,
				},
				resource.Attribute{
					Name:        "public_gists",
					Description: `the number of public gists.`,
				},
				resource.Attribute{
					Name:        "followers",
					Description: `the number of followers.`,
				},
				resource.Attribute{
					Name:        "following",
					Description: `the number of following users.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `the creation date.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `the update date.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "node_id",
					Description: `the Node ID of the user.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `the user's login.`,
				},
				resource.Attribute{
					Name:        "avatar_url",
					Description: `the user's avatar URL.`,
				},
				resource.Attribute{
					Name:        "gravatar_id",
					Description: `the user's gravatar ID.`,
				},
				resource.Attribute{
					Name:        "site_admin",
					Description: `whether the user is a GitHub admin.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the user's full name.`,
				},
				resource.Attribute{
					Name:        "company",
					Description: `the user's company name.`,
				},
				resource.Attribute{
					Name:        "blog",
					Description: `the user's blog location.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `the user's location.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `the user's email.`,
				},
				resource.Attribute{
					Name:        "gpg_keys",
					Description: `list of user's GPG keys.`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `list of user's SSH keys.`,
				},
				resource.Attribute{
					Name:        "bio",
					Description: `the user's bio.`,
				},
				resource.Attribute{
					Name:        "public_repos",
					Description: `the number of public repositories.`,
				},
				resource.Attribute{
					Name:        "public_gists",
					Description: `the number of public gists.`,
				},
				resource.Attribute{
					Name:        "followers",
					Description: `the number of followers.`,
				},
				resource.Attribute{
					Name:        "following",
					Description: `the number of following users.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `the creation date.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `the update date.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"github_actions_public_key":            0,
		"github_branch":                        1,
		"github_collaborators":                 2,
		"github_ip_ranges":                     3,
		"github_membership":                    4,
		"github_organization":                  5,
		"github_organization_team_sync_groups": 6,
		"github_organization_teams":            7,
		"github_release":                       8,
		"github_repositories":                  9,
		"github_repository":                    10,
		"github_repository_milestone":          11,
		"github_team":                          12,
		"github_user":                          13,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
