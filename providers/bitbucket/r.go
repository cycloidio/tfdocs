package bitbucket

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "bitbucket_branch_restriction",
			Category:         "Resources",
			ShortDescription: `Provides a Bitbucket Branch Restriction`,
			Description:      ``,
			Keywords: []string{
				"branch",
				"restriction",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "owner",
					Description: `(Required) The owner of this repository. Can be you or any team you have write access to.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The name of the repository.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The type of restriction that is being applied. List of possible stages is [here](https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Busername%7D/%7Brepo_slug%7D/branch-restrictions).`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Required) The pattern to determine which branches will be restricted.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) A list of users to use.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) A list of groups to use.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bitbucket_default_reviewers",
			Category:         "Resources",
			ShortDescription: `Provides support for setting up default reviews for bitbucket.`,
			Description:      ``,
			Keywords: []string{
				"default",
				"reviewers",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "owner",
					Description: `(Required) The owner of this repository. Can be you or any team you have write access to.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The name of the repository.`,
				},
				resource.Attribute{
					Name:        "reviewers",
					Description: `(Required) A list of reviewers to use.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bitbucket_hook",
			Category:         "Resources",
			ShortDescription: `Provides a Bitbucket Webhook`,
			Description:      ``,
			Keywords: []string{
				"hook",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "owner",
					Description: `(Required) The owner of this repository. Can be you or any team you have write access to.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The name of the repository.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Where to POST to.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The name / description to show in the UI.`,
				},
				resource.Attribute{
					Name:        "events",
					Description: `(Required) The event you want to react on.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bitbucket_project",
			Category:         "Resources",
			ShortDescription: `Create and manage a Bitbucket project`,
			Description:      ``,
			Keywords: []string{
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "owner",
					Description: `(Required) The owner of this project. Can be you or any team you have write access to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key used for this project`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the project`,
				},
				resource.Attribute{
					Name:        "is_private",
					Description: `(Optional) If you want to keep the project private - defaults to true`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bitbucket_repository",
			Category:         "Resources",
			ShortDescription: `Provides a Bitbucket Repository`,
			Description:      ``,
			Keywords: []string{
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "owner",
					Description: `(Required) The owner of this repository. Can be you or any team you have write access to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the repository.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Optional) The slug of the repository.`,
				},
				resource.Attribute{
					Name:        "scm",
					Description: `(Optional) What SCM you want to use. Valid options are hg or git. Defaults to git.`,
				},
				resource.Attribute{
					Name:        "is_private",
					Description: `(Optional) If this should be private or not. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `(Optional) URL of website associated with this repository.`,
				},
				resource.Attribute{
					Name:        "language",
					Description: `(Optional) What the language of this repository should be.`,
				},
				resource.Attribute{
					Name:        "has_issues",
					Description: `(Optional) If this should have issues turned on or not.`,
				},
				resource.Attribute{
					Name:        "has_wiki",
					Description: `(Optional) If this should have wiki turned on or not.`,
				},
				resource.Attribute{
					Name:        "project_key",
					Description: `(Optional) If you want to have this repo associated with a project.`,
				},
				resource.Attribute{
					Name:        "fork_policy",
					Description: `(Optional) What the fork policy should be. Defaults to allow_forks.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) What the description of the repo is.`,
				},
				resource.Attribute{
					Name:        "pipelines_enabled",
					Description: `(Optional) Turn on to enable pipelines support ## Computed Arguments The following arguments are computed. You can access both ` + "`" + `clone_ssh` + "`" + ` and ` + "`" + `clone_https` + "`" + ` for getting a clone URL. ## Import Repositories can be imported using their ` + "`" + `owner/name` + "`" + ` ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import bitbucket_repository.my-repo my-account/my-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bitbucket_repository_variable",
			Category:         "Resources",
			ShortDescription: `Manage your pipelines repository variables and configuration`,
			Description:      ``,
			Keywords: []string{
				"repository",
				"variable",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"bitbucket_branch_restriction":  0,
		"bitbucket_default_reviewers":   1,
		"bitbucket_hook":                2,
		"bitbucket_project":             3,
		"bitbucket_repository":          4,
		"bitbucket_repository_variable": 5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
