package nsxt

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nsxt_certificate",
			Category:         "Data Sources",
			ShortDescription: `A certificate data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Certificate to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name of the Certificate to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Certificate.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_edge_cluster",
			Category:         "Data Sources",
			ShortDescription: `An Edge Cluster data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Edge Cluster to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Edge Cluster to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the edge cluster.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `This field could show deployment_type of members. It would return UNKNOWN if there is no members, and return VIRTUAL_MACHINE|PHYSICAL_MACHINE if all Edge members are VIRTUAL_MACHINE|PHYSICAL_MACHINE.`,
				},
				resource.Attribute{
					Name:        "member_node_type",
					Description: `An Edge cluster is homogeneous collection of NSX transport nodes used for north/south connectivity between NSX logical networking and physical networking. Hence all transport nodes of the cluster must be of same type. This field shows the type of transport node,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the edge cluster.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `This field could show deployment_type of members. It would return UNKNOWN if there is no members, and return VIRTUAL_MACHINE|PHYSICAL_MACHINE if all Edge members are VIRTUAL_MACHINE|PHYSICAL_MACHINE.`,
				},
				resource.Attribute{
					Name:        "member_node_type",
					Description: `An Edge cluster is homogeneous collection of NSX transport nodes used for north/south connectivity between NSX logical networking and physical networking. Hence all transport nodes of the cluster must be of same type. This field shows the type of transport node,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_firewall_section",
			Category:         "Data Sources",
			ShortDescription: `A firewall section data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of resource to retrieve`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name of resource to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ip_pool",
			Category:         "Data Sources",
			ShortDescription: `A IP pool data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of IP pool to retrieve`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name of the IP pool to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the IP pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the IP pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_tier0_router",
			Category:         "Data Sources",
			ShortDescription: `A logical Tier 0 router data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Logical Tier 0 Router to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Logical Tier 0 Router to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the logical Tier 0 router.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `The id of the Edge cluster where this logical router is placed.`,
				},
				resource.Attribute{
					Name:        "high_availability_mode",
					Description: `The high availability mode of this logical router.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the logical Tier 0 router.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `The id of the Edge cluster where this logical router is placed.`,
				},
				resource.Attribute{
					Name:        "high_availability_mode",
					Description: `The high availability mode of this logical router.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_tier1_router",
			Category:         "Data Sources",
			ShortDescription: `A logical Tier 1 router data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Logical Tier 1 Router to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Logical Tier 1 Router to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the logical Tier 0 router.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `The id of the Edge cluster where this logical router is placed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the logical Tier 0 router.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `The id of the Edge cluster where this logical router is placed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_mac_pool",
			Category:         "Data Sources",
			ShortDescription: `A MAC pool data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of MAC pool to retrieve`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name of the MAC pool to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the MAC pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the MAC pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ns_group",
			Category:         "Data Sources",
			ShortDescription: `A networking and security group data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of NS group to retrieve`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name of the NS group to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NS group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NS group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ns_service",
			Category:         "Data Sources",
			ShortDescription: `A networking and security service data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of NS service to retrieve`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name of the NS service to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NS service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NS service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_certificate",
			Category:         "Data Sources",
			ShortDescription: `Policy Certificate data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Certificate to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Certificate to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_edge_cluster",
			Category:         "Data Sources",
			ShortDescription: `A policy Edge Cluster data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the edge cluster to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the edge cluster to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_edge_node",
			Category:         "Data Sources",
			ShortDescription: `A policy Edge Node data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "edge_cluster_path",
					Description: `(Required) The path of edge cluster where to which this node belongs.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the edge node to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the edge node to retrieve.`,
				},
				resource.Attribute{
					Name:        "member_index",
					Description: `(Optional) Member index of the node in edge cluster. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_gateway_qos_profile",
			Category:         "Data Sources",
			ShortDescription: `Policy GatewayQosProfile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of GatewayQosProfile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the GatewayQosProfile to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_ip_block",
			Category:         "Data Sources",
			ShortDescription: `Policy IP Block Config data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of IP Block Config to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the IP Block Config to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_ip_discovery_profile",
			Category:         "Data Sources",
			ShortDescription: `Policy IpDiscoveryProfile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of IpDiscoveryProfile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the IpDiscoveryProfile to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_ip_pool",
			Category:         "Data Sources",
			ShortDescription: `Policy IP Pool Config data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of IP Pool Config to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the IP Pool Config to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_ipv6_dad_profile",
			Category:         "Data Sources",
			ShortDescription: `Policy Ipv6DadProfile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Ipv6DadProfile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Ipv6DadProfile to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_ipv6_ndra_profile",
			Category:         "Data Sources",
			ShortDescription: `Policy Ipv6NdraProfile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Ipv6NdraProfile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Ipv6NdraProfile to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_lb_app_profile",
			Category:         "Data Sources",
			ShortDescription: `Policy Load Balancer Application Profile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Profile to retrieve.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of Profile to retrieve, one of ` + "`" + `HTTP` + "`" + `, ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ANY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Profile to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_lb_client_ssl_profile",
			Category:         "Data Sources",
			ShortDescription: `Policy Load Balancer Client SSL Profile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Profile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Profile to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_lb_monitor",
			Category:         "Data Sources",
			ShortDescription: `Policy Load Balancer Monitor data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Monitor to retrieve.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of Monitor to retrieve, one of ` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + `, ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + `, ` + "`" + `PASSIVE` + "`" + `, ` + "`" + `ANY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of Monitor to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_lb_persistence_profile",
			Category:         "Data Sources",
			ShortDescription: `Policy Load Balancer Persistence Profile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Load Balanacer Persistence Profile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Load Balancer Persistence Profile to retrieve.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The Load Balancer Persistence Profile type. One of ` + "`" + `ANY` + "`" + `, ` + "`" + `SOURCE_IP` + "`" + `, ` + "`" + `COOKIE` + "`" + ` or ` + "`" + `GENERIC` + "`" + `. Defaults to ` + "`" + `ANY` + "`" + `. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_lb_server_ssl_profile",
			Category:         "Data Sources",
			ShortDescription: `Policy Load Balancer Server SSL Profile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Profile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Profile to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_mac_discovery_profile",
			Category:         "Data Sources",
			ShortDescription: `Policy MacDiscoveryProfile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of MacDiscoveryProfile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the MacDiscoveryProfile to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_qos_profile",
			Category:         "Data Sources",
			ShortDescription: `Policy QosProfile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of QosProfile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the QosProfile to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_realization_info",
			Category:         "Data Sources",
			ShortDescription: `A policy resource realization information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The policy path of the resource.`,
				},
				resource.Attribute{
					Name:        "entity_type",
					Description: `(Optional) The entity type of realized resource. If not set, on of the realized resources of the policy resource will be retrieved. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The realization state of the resource: "REALIZED", "UNKNOWN", "UNREALIZED" or "ERROR".`,
				},
				resource.Attribute{
					Name:        "realized_id",
					Description: `The id of the realized object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `The realization state of the resource: "REALIZED", "UNKNOWN", "UNREALIZED" or "ERROR".`,
				},
				resource.Attribute{
					Name:        "realized_id",
					Description: `The id of the realized object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_segment_realization",
			Category:         "Data Sources",
			ShortDescription: `State of segment realization on hypervisors.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The policy path of the segment. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The realization state of the resource: ` + "`" + `success` + "`" + `, ` + "`" + `partial_success` + "`" + `, ` + "`" + `orphaned` + "`" + `, ` + "`" + `failed` + "`" + ` or ` + "`" + `error` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Network name on the hypervisor. This attribute can be used in vsphere provider in order to ensure implicit dependency on segment realization.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `The realization state of the resource: ` + "`" + `success` + "`" + `, ` + "`" + `partial_success` + "`" + `, ` + "`" + `orphaned` + "`" + `, ` + "`" + `failed` + "`" + ` or ` + "`" + `error` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Network name on the hypervisor. This attribute can be used in vsphere provider in order to ensure implicit dependency on segment realization.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_segment_security_profile",
			Category:         "Data Sources",
			ShortDescription: `Policy SegmentSecurityProfile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of SegmentSecurityProfile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the SegmentSecurityProfile to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_service",
			Category:         "Data Sources",
			ShortDescription: `A policy service data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of service to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the service to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_spoofguard_profile",
			Category:         "Data Sources",
			ShortDescription: `Policy SpoofGuardProfile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of SpoofGuardProfile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the SpoofGuardProfile to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_tier0_gateway",
			Category:         "Data Sources",
			ShortDescription: `A policy Tier-0 gateway data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Tier-0 gateway to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Tier-0 gateway to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_path",
					Description: `The path of the Edge cluster where this Tier-0 gateway is placed.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_path",
					Description: `The path of the Edge cluster where this Tier-0 gateway is placed.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_tier1_gateway",
			Category:         "Data Sources",
			ShortDescription: `A policy Tier-1 gateway data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Tier-1 gateway to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Tier-1 gateway to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_path",
					Description: `The path of the Edge cluster where this Tier-1 gateway is placed.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_path",
					Description: `The path of the Edge cluster where this Tier-1 gateway is placed.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_transport_zone",
			Category:         "Data Sources",
			ShortDescription: `A Policy Transport Zone data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Transport Zone to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Transport Zone to retrieve.`,
				},
				resource.Attribute{
					Name:        "transport_type",
					Description: `(Optional) Transport type of requested Transport Zone, one of ` + "`" + `OVERLAY_STANDARD` + "`" + `, ` + "`" + `OVERLAY_ENS` + "`" + `, ` + "`" + `VLAN_BACKED` + "`" + ` and ` + "`" + `UNKNOWN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) May be set together with ` + "`" + `transport_type` + "`" + ` in order to retrieve default Transport Zone for for this transport type. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Transport Zone.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `A boolean flag indicating if this Transport Zone is the default.`,
				},
				resource.Attribute{
					Name:        "transport_type",
					Description: `The transport type of this transport zone.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Transport Zone.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `A boolean flag indicating if this Transport Zone is the default.`,
				},
				resource.Attribute{
					Name:        "transport_type",
					Description: `The transport type of this transport zone.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_vm",
			Category:         "Data Sources",
			ShortDescription: `A Policy Virtual Machine ID data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Virtual Machine to retrieve.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Optional) The external ID of the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "bios_id",
					Description: `(Optional) The BIOS UUID of the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) The instance UUID of the Virtual Machine. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Virtual Machine.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Virtual Machine.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_vni_pool",
			Category:         "Data Sources",
			ShortDescription: `Policy VNI Pool Config data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of VNI Pool Config to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the VNI Pool Config to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `The start range of VNI Pool.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `The end range of VNI Pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `The start range of VNI Pool.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `The end range of VNI Pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_switching_profile",
			Category:         "Data Sources",
			ShortDescription: `A switching profile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Switching Profile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name of the Switching Profile to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `The resource type representing the specific type of this switching profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the switching profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_type",
					Description: `The resource type representing the specific type of this switching profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the switching profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_transport_zone",
			Category:         "Data Sources",
			ShortDescription: `A transport zone data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Transport Zone to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Transport Zone to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Transport Zone.`,
				},
				resource.Attribute{
					Name:        "host_switch_name",
					Description: `The name of the N-VDS (host switch) on all Transport Nodes in this Transport Zone that will be used to run NSX network traffic.`,
				},
				resource.Attribute{
					Name:        "transport_type",
					Description: `The transport type of this transport zone (OVERLAY or VLAN).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Transport Zone.`,
				},
				resource.Attribute{
					Name:        "host_switch_name",
					Description: `The name of the N-VDS (host switch) on all Transport Nodes in this Transport Zone that will be used to run NSX network traffic.`,
				},
				resource.Attribute{
					Name:        "transport_type",
					Description: `The transport type of this transport zone (OVERLAY or VLAN).`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"nsxt_certificate":                     0,
		"nsxt_edge_cluster":                    1,
		"nsxt_firewall_section":                2,
		"nsxt_ip_pool":                         3,
		"nsxt_logical_tier0_router":            4,
		"nsxt_logical_tier1_router":            5,
		"nsxt_mac_pool":                        6,
		"nsxt_ns_group":                        7,
		"nsxt_ns_service":                      8,
		"nsxt_policy_certificate":              9,
		"nsxt_policy_edge_cluster":             10,
		"nsxt_policy_edge_node":                11,
		"nsxt_policy_gateway_qos_profile":      12,
		"nsxt_policy_ip_block":                 13,
		"nsxt_policy_ip_discovery_profile":     14,
		"nsxt_policy_ip_pool":                  15,
		"nsxt_policy_ipv6_dad_profile":         16,
		"nsxt_policy_ipv6_ndra_profile":        17,
		"nsxt_policy_lb_app_profile":           18,
		"nsxt_policy_lb_client_ssl_profile":    19,
		"nsxt_policy_lb_monitor":               20,
		"nsxt_policy_lb_persistence_profile":   21,
		"nsxt_policy_lb_server_ssl_profile":    22,
		"nsxt_policy_mac_discovery_profile":    23,
		"nsxt_policy_qos_profile":              24,
		"nsxt_policy_realization_info":         25,
		"nsxt_policy_segment_realization":      26,
		"nsxt_policy_segment_security_profile": 27,
		"nsxt_policy_service":                  28,
		"nsxt_policy_spoofguard_profile":       29,
		"nsxt_policy_tier0_gateway":            30,
		"nsxt_policy_tier1_gateway":            31,
		"nsxt_policy_transport_zone":           32,
		"nsxt_policy_vm":                       33,
		"nsxt_policy_vni_pool":                 34,
		"nsxt_switching_profile":               35,
		"nsxt_transport_zone":                  36,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
