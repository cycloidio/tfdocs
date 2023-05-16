package vercel

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vercel_alias",
			Category:         "Resources",
			ShortDescription: `Provides an Alias resource. An Alias allows a vercel_deployment to be accessed through a different URL.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vercel_deployment",
			Category:         "Resources",
			ShortDescription: `Provides a Deployment resource. A Deployment is the result of building your Project and making it available through a live URL. When making deployments, the Project will be uploaded and transformed into a production-ready output through the use of a Build Step. Once the build step has completed successfully, a new, immutable deployment will be made available at the preview URL. Deployments are retained indefinitely unless deleted manually. -> In order to provide files to a deployment, you'll need to use the vercel_file or vercel_project_directory data sources. ~> If you are creating Deployments through terraform and intend to use both preview and production deployments, you may wish to 'layer' your terraform, creating the Project with a different set of terraform to your Deployment.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vercel_dns_record",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vercel_project",
			Category:         "Resources",
			ShortDescription: `Provides a Project resource. A Project groups deployments and custom domains. To deploy on Vercel, you need to create a Project. For more detailed information, please see the Vercel documentation https://vercel.com/docs/concepts/projects/overview. ~> Terraform currently provides both a standalone Project Environment Variable resource (a single Environment Variable), and a Project resource with Environment Variables defined in-line via the environment field. At this time you cannot use a Vercel Project resource with in-line environment in conjunction with any vercel_project_environment_variable resources. Doing so will cause a conflict of settings and will overwrite Environment Variables.`,
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
		&resource.Resource{
			Name:             "",
			Type:             "vercel_project_environment_variable",
			Category:         "Resources",
			ShortDescription: `Provides a Project Environment Variable resource. A Project Environment Variable resource defines an Environment Variable on a Vercel Project. For more detailed information, please see the Vercel documentation https://vercel.com/docs/concepts/projects/environment-variables. ~> Terraform currently provides both a standalone Project Environment Variable resource (a single Environment Variable), and a Project resource with Environment Variables defined in-line via the environment field. At this time you cannot use a Vercel Project resource with in-line environment in conjunction with any vercel_project_environment_variable resources. Doing so will cause a conflict of settings and will overwrite Environment Variables.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vercel_shared_environment_variable",
			Category:         "Resources",
			ShortDescription: `Provides a Shared Environment Variable resource. A Shared Environment Variable resource defines an Environment Variable that can be shared between multiple Vercel Projects. For more detailed information, please see the Vercel documentation https://vercel.com/docs/concepts/projects/environment-variables/shared-environment-variables.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"vercel_alias":                        0,
		"vercel_deployment":                   1,
		"vercel_dns_record":                   2,
		"vercel_project":                      3,
		"vercel_project_domain":               4,
		"vercel_project_environment_variable": 5,
		"vercel_shared_environment_variable":  6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
