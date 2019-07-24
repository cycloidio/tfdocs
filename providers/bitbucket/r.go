package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{},
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{},
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
			Arguments: []resource.Argument{
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
					Description: `(Optional) What the description of the repo is. ## Computed Arguments The following arguments are computed. You can access both ` + "`" + `clone_ssh` + "`" + ` and ` + "`" + `clone_https` + "`" + ` for getting a clone URL. ## Import Repositories can be imported using their ` + "`" + `owner/name` + "`" + ` ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import bitbucket_repository.my-repo my-account/my-repo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"bitbucket_default_reviewers": 0,
		"bitbucket_hook":              1,
		"bitbucket_repository":        2,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
