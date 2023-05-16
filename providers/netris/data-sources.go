package netris

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "netris_bgp_object",
			Category:         "Data Sources",
			ShortDescription: `Data Source: BGP Objects`,
			Description: `

You can define various BGP objects referenced from a route-map to declare a dynamic BGP policy.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "netris_dhcp_option_set",
			Category:         "Data Sources",
			ShortDescription: `Data Source: dhcp option sets`,
			Description: `

Get dhcp option set resource in Netris.
`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "netris_inventory_profile",
			Category:         "Data Sources",
			ShortDescription: `Data Source: inventory profiles`,
			Description: `

Inventory profiles allow security hardening of inventory devices. By default all traffic flow destined to switch/SoftGate is allowed. As soon as the inventory profile is attached to a device it denies all traffic destined to the device except Netris-defined and user-defined custom flows.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "netris_network_interface",
			Category:         "Data Sources",
			ShortDescription: `Data Source: Network Interface`,
			Description: `

Network Interface data.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "netris_routemap",
			Category:         "Data Sources",
			ShortDescription: `Data Source: BGP Route-maps`,
			Description: `

Route-maps section, you can define route-map policies, which can be associated with the BGP neighbors inbound or outbound.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "netris_site",
			Category:         "Data Sources",
			ShortDescription: `Data Source: Sites`,
			Description: `

Each separate deployment (each data center) should be defined as a Site. All network units and resources are attached to a site. Netris Controller comes with a “default” site preconfigured. Site entry defines global attributes such as; AS numbers, default ACL policy, and Site Mesh (site to site VPN) type.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "netris_tenant",
			Category:         "Data Sources",
			ShortDescription: `Data Source: Tenants`,
			Description: `

IP addresses and Switch Ports are network resources that can be assigned to different Tenants to have under their management. Admin is the default tenant, and by default, it owns all the resources. The concept of Tenants can be used for sharing and delegation of control over the network resources, typically used by network teams to grant access to other teams for requesting & managing network services using the Netris Controller as a self service portal or programmatically (with Kubernetes CRDs) as part of DevOps/NetOps pipeline.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "netris_vnet",
			Category:         "Data Sources",
			ShortDescription: `Data Source: Vnets`,
			Description: `

V-Net is a virtual networking service that provide a Layer-2 (unrouted) or Layer-3 (routed) virtual network segments on switch ports anywhere on the switch fabric. V-NETs can be created and managed by a single tenant (single team) or they can be created and managed collaboratively by multiple tenants (different teams inside and/or outside the organization). Netris automatically configures a VXLAN with an EVPN control plane over an unnumbered BGP Layer-3 underlay network and organize the high availability for the default gateway behind the scenes.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"netris_bgp_object":        0,
		"netris_dhcp_option_set":   1,
		"netris_inventory_profile": 2,
		"netris_network_interface": 3,
		"netris_routemap":          4,
		"netris_site":              5,
		"netris_tenant":            6,
		"netris_vnet":              7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
