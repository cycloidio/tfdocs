package github

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

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
					Description: `The id of the collaborator.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The github api url for the collaborator.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `The github html url for the collaborator.`,
				},
				resource.Attribute{
					Name:        "followers_url",
					Description: `The github api url for the collaborator's followers.`,
				},
				resource.Attribute{
					Name:        "following_url",
					Description: `The github api url for those following the collaborator.`,
				},
				resource.Attribute{
					Name:        "gists_url",
					Description: `The github api url for the collaborator's gists.`,
				},
				resource.Attribute{
					Name:        "starred_url",
					Description: `The github api url for the collaborator's starred repositories.`,
				},
				resource.Attribute{
					Name:        "subscriptions_url",
					Description: `The github api url for the collaborator's subscribed repositories.`,
				},
				resource.Attribute{
					Name:        "organizations_url",
					Description: `The github api url for the collaborator's organizations.`,
				},
				resource.Attribute{
					Name:        "repos_url",
					Description: `The github api url for the collaborator's repositories.`,
				},
				resource.Attribute{
					Name:        "events_url",
					Description: `The github api url for the collaborator's events.`,
				},
				resource.Attribute{
					Name:        "received_events_url",
					Description: `The github api url for the collaborator's received events.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the collaborator (ex. ` + "`" + `User` + "`" + `).`,
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
			ShortDescription: `Get information on a GitHub's IP addresses.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
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
			},
			Attributes: []resource.Attribute{
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `the ID of the team.`,
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
					Description: `(Required) The username. ## Attributes Reference`,
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
					Description: `list of user's GPG keys`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `list of user's SSH keys`,
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
					Description: `list of user's GPG keys`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `list of user's SSH keys`,
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

		"github_collaborators": 0,
		"github_ip_ranges":     1,
		"github_repositories":  2,
		"github_repository":    3,
		"github_team":          4,
		"github_user":          5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
