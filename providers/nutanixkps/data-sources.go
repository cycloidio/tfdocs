package nutanixkps

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_node",
			Category:         "Data Sources",
			ShortDescription: `Describes a Karbon Platform Services Service Domain Node. A Service Domain Node is the baremetal/virtual machine that is being managed by Karbon Platform Services.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_nodes",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_serviceclass",
			Category:         "Data Sources",
			ShortDescription: `Describes a Karbon Platform Services Service Class.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_serviceclasses",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_servicedomains",
			Category:         "Data Sources",
			ShortDescription: `Describes a Nutanix KPS Service Domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_user",
			Category:         "Data Sources",
			ShortDescription: `Describes a Karbon Platform Services User.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_users",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_vm_config",
			Category:         "Data Sources",
			ShortDescription: `Describes a Karbon Platform Services Virtual Machine.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"nutanixkps_node":           0,
		"nutanixkps_nodes":          1,
		"nutanixkps_serviceclass":   2,
		"nutanixkps_serviceclasses": 3,
		"nutanixkps_servicedomains": 4,
		"nutanixkps_user":           5,
		"nutanixkps_users":          6,
		"nutanixkps_vm_config":      7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
