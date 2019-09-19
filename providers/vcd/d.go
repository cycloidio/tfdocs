package vcd

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_app_profile",
			Category:         "Data Sources",
			ShortDescription: `Provides an NSX edge gateway load balancer application profile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the service monitor is defined`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Application profile name for identifying the exact application profile ## Attribute Reference All the attributes defined in ` + "`" + `vcd_lb_app_profile` + "`" + ` resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_app_rule",
			Category:         "Data Sources",
			ShortDescription: `Provides an NSX edge gateway load balancer application rule data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the service monitor is defined`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Application rule name for identifying the exact application rule ## Attribute Reference All the attributes defined in ` + "`" + `vcd_lb_app_rule` + "`" + ` resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_server_pool",
			Category:         "Data Sources",
			ShortDescription: `Provides an NSX edge gateway load balancer server pool data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the server pool is defined`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Server Pool name for identifying the exact server pool ## Attribute Reference All the attributes defined in ` + "`" + `vcd_lb_server_pool` + "`" + ` resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_service_monitor",
			Category:         "Data Sources",
			ShortDescription: `Provides an NSX edge gateway load balancer service monitor data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the service monitor is defined`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Service Monitor name for identifying the exact service monitor ## Attribute Reference All the attributes defined in ` + "`" + `vcd_lb_service_monitor` + "`" + ` resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_virtual_server",
			Category:         "Data Sources",
			ShortDescription: `Provides an NSX edge gateway load balancer virtual server data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the virtual server is defined`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for identifying the exact virtual server ## Attribute Reference All the attributes defined in ` + "`" + `vcd_lb_virtual_server` + "`" + ` resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"vcd_lb_app_profile":     0,
		"vcd_lb_app_rule":        1,
		"vcd_lb_server_pool":     2,
		"vcd_lb_service_monitor": 3,
		"vcd_lb_virtual_server":  4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
