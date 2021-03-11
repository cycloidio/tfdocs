package anxcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_core-locations",
			Category:         "Data Sources",
			ShortDescription: `The core locations data source allows you to get all available core locations.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_cpu-performance-types",
			Category:         "Data Sources",
			ShortDescription: `The cpu performance types data source allows you to get all available cpu performance types.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_disk-types",
			Category:         "Data Sources",
			ShortDescription: `The disk types data source allows you to retrieve information about available disk types for specified location.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_ip-addresses",
			Category:         "Data Sources",
			ShortDescription: `The IP addresses data source allows you to get all available IP addresses.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_nic-types",
			Category:         "Data Sources",
			ShortDescription: `The NIC types data source allows you to get all available network interface card types.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_tags",
			Category:         "Data Sources",
			ShortDescription: `The tags data source allows you to get all available tags.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_template",
			Category:         "Data Sources",
			ShortDescription: `The template data source allows you to retrieve information about available templates for specified location.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_vlans",
			Category:         "Data Sources",
			ShortDescription: `The VLANs data source allows you to get all available VLANs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_vsphere-locations",
			Category:         "Data Sources",
			ShortDescription: `The vsphere locations data source allows you to get all available vsphere locations.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"anxcloud_core-locations":        0,
		"anxcloud_cpu-performance-types": 1,
		"anxcloud_disk-types":            2,
		"anxcloud_ip-addresses":          3,
		"anxcloud_nic-types":             4,
		"anxcloud_tags":                  5,
		"anxcloud_template":              6,
		"anxcloud_vlans":                 7,
		"anxcloud_vsphere-locations":     8,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
