package cloudsigma

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudsigma_ip",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to use as filters. ## Attributes Reference In addition to all above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "resource_uri",
					Description: `The resource URI of the IP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_uri",
					Description: `The resource URI of the IP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsigma_library_drive",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to use as filters. ## Attributes Reference In addition to all above arguments, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsigma_location",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to use as filters. ## Attributes Reference In addition to all above arguments, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsigma_vlan",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to use as filters. ## Attributes Reference In addition to all above arguments, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"cloudsigma_ip":            0,
		"cloudsigma_library_drive": 1,
		"cloudsigma_location":      2,
		"cloudsigma_vlan":          3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
