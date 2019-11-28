package opennebula

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "opennebula_group",
			Category:         "Data Sources",
			ShortDescription: `Get the group information for a given name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The OpenNebula group to retrieve information for.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_image",
			Category:         "Data Sources",
			ShortDescription: `Get the image information for a given name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The OpenNebula image to retrieve information for.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_security group",
			Category:         "Data Sources",
			ShortDescription: `Get the security group information for a given name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The OpenNebula security group to retrieve information for.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_template",
			Category:         "Data Sources",
			ShortDescription: `Get the template information for a given name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The OpenNebula template to retrieve information for.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_virtual data center",
			Category:         "Data Sources",
			ShortDescription: `Get the virtual data center information for a given name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The OpenNebula virtual data center to retrieve information for.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_virtual network",
			Category:         "Data Sources",
			ShortDescription: `Get the virtual network information for a given name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The OpenNebula virtual network to retrieve information for.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"opennebula_group":               0,
		"opennebula_image":               1,
		"opennebula_security group":      2,
		"opennebula_template":            3,
		"opennebula_virtual data center": 4,
		"opennebula_virtual network":     5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
