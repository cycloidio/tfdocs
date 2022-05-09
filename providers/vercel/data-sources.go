package vercel

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vercel_file",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a file on disk. This will read a single file, providing metadata for use with a vercel_deployment.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vercel_project",
			Category:         "Data Sources",
			ShortDescription: `Provides information about an existing project within Vercel. A Project groups deployments and custom domains. To deploy on Vercel, you need a Project. For more detailed information, please see the Vercel documentation https://vercel.com/docs/concepts/projects/overview.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vercel_project_directory",
			Category:         "Data Sources",
			ShortDescription: `Provides information about files within a directory on disk. This will recursively read files, providing metadata for use with a vercel_deployment. -> If you want to prevent files from being included, this can be done with a vercelignore file https://vercel.com/guides/prevent-uploading-sourcepaths-with-vercelignore.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"vercel_file":              0,
		"vercel_project":           1,
		"vercel_project_directory": 2,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
