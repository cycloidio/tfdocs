package ome

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ome_configuration_report_info",
			Category:         "Data Sources",
			ShortDescription: `Data source to list the compliance configuration report of a baseline from OpenManage Enterprise.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ome_groupdevices_info",
			Category:         "Data Sources",
			ShortDescription: `Data source to list the devices in the group from OpenManage Enterprise.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ome_template_info",
			Category:         "Data Sources",
			ShortDescription: `Data Source to list the Template details from OpenManage Enterprise`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ome_vlannetworks_info",
			Category:         "Data Sources",
			ShortDescription: `Data source to list the vlan networks from OpenManage Enterprise.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"ome_configuration_report_info": 0,
		"ome_groupdevices_info":         1,
		"ome_template_info":             2,
		"ome_vlannetworks_info":         3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
