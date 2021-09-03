package britive

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "britive_application",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the application. ## Attribute Reference In addition to the above argument, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An identifier for the application.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An identifier for the application.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "britive_identity_provider",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the identity provider. ## Attribute Reference In addition to the above argument, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An identifier for the identity provider.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An identifier for the identity provider.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"britive_application":       0,
		"britive_identity_provider": 1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
