package cloudsmith

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudsmith_namespace",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The slug identifies the namespace in URIs. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A descriptive name for the namespace.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug identifies the namespace in URIs.`,
				},
				resource.Attribute{
					Name:        "slug_perm",
					Description: `The slug_perm immutably identifies the namespace. It will never change once a namespace has been created.`,
				},
				resource.Attribute{
					Name:        "type_name",
					Description: `Is this a user or an organization namespace?.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `A descriptive name for the namespace.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug identifies the namespace in URIs.`,
				},
				resource.Attribute{
					Name:        "slug_perm",
					Description: `The slug_perm immutably identifies the namespace. It will never change once a namespace has been created.`,
				},
				resource.Attribute{
					Name:        "type_name",
					Description: `Is this a user or an organization namespace?.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"cloudsmith_namespace": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
