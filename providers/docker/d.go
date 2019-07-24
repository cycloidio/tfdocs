package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "docker_registry_image",
			Category:         "Data Sources",
			ShortDescription: `Finds the latest available sha256 digest for a docker image/tag from a registry.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The name of the Docker image, including any tags. e.g. ` + "`" + `alpine:latest` + "`" + ` ## Attributes Reference The following attributes are exported in addition to the above configuration:`,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	dataSourcesMap = map[string]Resource{

		"docker_registry_image": 0,
	}
)

func GetDataSource(r string) (*resouce.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs]
}
