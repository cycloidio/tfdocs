package hpegl

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "hpegl_vmaas_instance",
			Category:         "Resources",
			ShortDescription: `Instance resource facilitates creating, updating and deleting virtual machines. It is recommend to use the Vmware type for provisioning.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hpegl_vmaas_instance_clone",
			Category:         "Resources",
			ShortDescription: `Instance clone resource facilitates creating, updating and deleting cloned virtual machines. For creating an instance clone, provide a unique name and all the Mandatory(Required) parameters. All optional parameters will be inherits from parent resource if not provided.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hpegl_vmaas_network",
			Category:         "Resources",
			ShortDescription: `Network resource facilitates creating, updating and deleting NSX-T Networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hpegl_vmaas_router",
			Category:         "Resources",
			ShortDescription: `Router resource facilitates creating, updating and deleting NSX-T Tier0/Tier1 Network Routers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hpegl_vmaas_router_bgp_neighbor",
			Category:         "Resources",
			ShortDescription: `Router Bgp Neighbor resource facilitates creating, updating and deleting NSX-T Network Router BGP Neighbors.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hpegl_vmaas_router_firewall_rule_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hpegl_vmaas_router_nat_rule",
			Category:         "Resources",
			ShortDescription: `Router NAT rule resource facilitates creating, updating and deleting NSX-T Network Router NAT rules.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hpegl_vmaas_router_route",
			Category:         "Resources",
			ShortDescription: `Router route resource facilitates creating, updating and deleting NSX-T Network Router routes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"hpegl_vmaas_instance":                   0,
		"hpegl_vmaas_instance_clone":             1,
		"hpegl_vmaas_network":                    2,
		"hpegl_vmaas_router":                     3,
		"hpegl_vmaas_router_bgp_neighbor":        4,
		"hpegl_vmaas_router_firewall_rule_group": 5,
		"hpegl_vmaas_router_nat_rule":            6,
		"hpegl_vmaas_router_route":               7,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
