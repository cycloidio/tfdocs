package vercel

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vercel_deployment",
			Category:         "Resources",
			ShortDescription: `Provides a Deployment resource. A Deployment is the result of building your Project and making it available through a live URL. When making deployments, the Project will be uploaded and transformed into a production-ready output through the use of a Build Step. Once the build step has completed successfully, a new, immutable deployment will be made available at the preview URL. Deployments are retained indefinitely unless deleted manually. -> In order to provide files to a deployment, you'll need to use the vercel_file or vercel_project_directory data sources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vercel_project",
			Category:         "Resources",
			ShortDescription: `Provides a Project resource. A Project groups deployments and custom domains. To deploy on Vercel, you need to create a Project. For more detailed information, please see the Vercel documentation https://vercel.com/docs/concepts/projects/overview. ~> If you are creating Deployments through terraform and intend to use both preview and production deployments, you may wish to 'layer' your terraform, creating the Project with a different set of terraform to your Deployment.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vercel_project_domain",
			Category:         "Resources",
			ShortDescription: `Provides a Project Domain resource. A Project Domain is used to associate a domain name with a vercel_project. By default, Project Domains will be automatically applied to any production deployments.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"vercel_deployment":     0,
		"vercel_project":        1,
		"vercel_project_domain": 2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
