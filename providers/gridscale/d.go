package gridscale

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "gridscale_ip",
			Category:         "Data Sources",
			ShortDescription: `Gets the id of an ip.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the ip.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_network",
			Category:         "Data Sources",
			ShortDescription: `Gets the id of a network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_sshkey",
			Category:         "Data Sources",
			ShortDescription: `Gets the id of an sshkey.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the sshkey.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_storage",
			Category:         "Data Sources",
			ShortDescription: `Gets the id of a storage.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the storage.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_template",
			Category:         "Data Sources",
			ShortDescription: `Gets the id of a template by name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The exact name of the template as show in [the expert panel of gridscale](https://my.gridscale.io/Expert/Template). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the template.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the template.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"gridscale_ip":       0,
		"gridscale_network":  1,
		"gridscale_sshkey":   2,
		"gridscale_storage":  3,
		"gridscale_template": 4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
