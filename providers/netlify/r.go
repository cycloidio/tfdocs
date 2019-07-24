package netlify

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "netlify_build_hook",
			Category:         "Resources",
			ShortDescription: `Provides an build hook resource.`,
			Description:      ``,
			Keywords: []string{
				"build",
				"hook",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) Your netlify site's unique id`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `(Required) branch to be built when the hook is triggered`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) name of the webhook - this is purely for organization and can be any name you want ## Attribute Reference The following additional attributes are exported:`,
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
			Type:             "netlify_deploy_key",
			Category:         "Resources",
			ShortDescription: `Provides a deploy key resource.`,
			Description:      ``,
			Keywords: []string{
				"deploy",
				"key",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "public_key",
					Description: `Public Key`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "netlify_hook",
			Category:         "Resources",
			ShortDescription: `Provides an hook resource.`,
			Description:      ``,
			Keywords: []string{
				"hook",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) - id of the site on netlify`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) - type of outgoing webhook, for example slack, email, github commit status, etc`,
				},
				resource.Attribute{
					Name:        "event",
					Description: `(Required) - when to send the data, for example on deploy create, succeed, fail, etc`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Required) object/hash of data to be sent along with the webhook. this varies depending on the ` + "`" + `type` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "netlify_site",
			Category:         "Resources",
			ShortDescription: `Provides an site resource.`,
			Description:      ``,
			Keywords: []string{
				"site",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Name of your site on Netlify (e.g.`,
				},
				resource.Attribute{
					Name:        "repo",
					Description: `(Required) - See [Repository](#repo)`,
				},
				resource.Attribute{
					Name:        "custom_domain",
					Description: `(Optional) - Custom domain of the site, must be configured using a CNAME in accordance with [Netlify's docs](https://www.netlify.com/docs/custom-domains). (e.g. ` + "`" + `www.example.com` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "deploy_url",
					Description: `(Optional) ### Repository ` + "`" + `repo` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Optional) - Shell command to run before deployment, typically used to build the site`,
				},
				resource.Attribute{
					Name:        "deploy_key_id",
					Description: `(Optional) - A deploy key id from the ` + "`" + `deploy_key` + "`" + ` resource`,
				},
				resource.Attribute{
					Name:        "dir",
					Description: `(Optional) - Directory to deploy, typically where the build puts the processed files`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Required) - Name of your VCS provider (e.g. ` + "`" + `github` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "repo_path",
					Description: `(Required) - path to your repo, typically ` + "`" + `username/reponame` + "`" + ``,
				},
				resource.Attribute{
					Name:        "repo_branch",
					Description: `(Required) - branch to be deployed`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"netlify_build_hook": 0,
		"netlify_deploy_key": 1,
		"netlify_hook":       2,
		"netlify_site":       3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
