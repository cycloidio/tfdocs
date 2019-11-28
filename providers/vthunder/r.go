package vthunder

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ethernet",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ethernet resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ethernet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ifnum",
					Description: `Index of the ethernet interface`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `flag to indicate if dhcp is enabled or disabled`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `ipv4 address to be assigned to the ethernet interface`,
				},
				resource.Attribute{
					Name:        "ipv4_netmask",
					Description: `ipv4 subnet mask to be assigned to the ethernet interface`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `flag to indicate if interface is enabled or disabled`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_rib_route",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder rib route resource for A10`,
			Description:      ``,
			Keywords: []string{
				"rib",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_dest_addr",
					Description: `Destination ip address`,
				},
				resource.Attribute{
					Name:        "ip_mask",
					Description: `Ip address subnet mask`,
				},
				resource.Attribute{
					Name:        "ip_next_hop",
					Description: `Ip next hop`,
				},
				resource.Attribute{
					Name:        "distance_nexthop_ip",
					Description: `Distance of next hop ip address`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_server",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder server resource for A10`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "health_check_disable",
					Description: `Disable configured health check configuration`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Server name`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Server host IP Address`,
				},
				resource.Attribute{
					Name:        "port_number",
					Description: `Port number`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `protocol`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_service_group",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder service group resource for A10`,
			Description:      ``,
			Keywords: []string{
				"service",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of service group`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `protocol`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_virtual_server",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder virtual server resource for A10`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of virtual server`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `Configure auto NAT for the vport`,
				},
				resource.Attribute{
					Name:        "port_number",
					Description: `Port number`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `protocol`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Service group name`,
				},
				resource.Attribute{
					Name:        "snat_on_vip",
					Description: `Enable source NAT traffic against VIP`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"vthunder_ethernet":       0,
		"vthunder_rib_route":      1,
		"vthunder_server":         2,
		"vthunder_service_group":  3,
		"vthunder_virtual_server": 4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
