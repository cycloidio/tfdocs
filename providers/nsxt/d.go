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
			Category:         "Manager",
			ShortDescription: `A certificate data source.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"certificate",
			},
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
			Category:         "Manager",
			ShortDescription: `An Edge Cluster data source.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"edge",
				"cluster",
			},
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
			Category:         "Manager",
			ShortDescription: `A firewall section data source.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"firewall",
				"section",
			},
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
			Category:         "Manager",
			ShortDescription: `A IP pool data source.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ip",
				"pool",
			},
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
			Category:         "Manager",
			ShortDescription: `A logical Tier 0 router data source.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"logical",
				"tier0",
				"router",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Logical Router to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of Logical Router to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Logical Router.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `The id of the Edge Cluster where this Logical Router is placed.`,
				},
				resource.Attribute{
					Name:        "high_availability_mode",
					Description: `The high availability mode of this Logical Router.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Logical Router.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `The id of the Edge Cluster where this Logical Router is placed.`,
				},
				resource.Attribute{
					Name:        "high_availability_mode",
					Description: `The high availability mode of this Logical Router.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_tier1_router",
			Category:         "Manager",
			ShortDescription: `A logical Tier 1 router data source.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"logical",
				"tier1",
				"router",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Logical Router to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of Logical Router to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Logical Router.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `The id of the Edge cluster where this Logical Router is placed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Logical Router.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `The id of the Edge cluster where this Logical Router is placed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_mac_pool",
			Category:         "Manager",
			ShortDescription: `A MAC pool data source.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"mac",
				"pool",
			},
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
			Type:             "nsxt_management_cluster",
			Category:         "Manager",
			ShortDescription: `A NSX-T management cluster data source.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"management",
				"cluster",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of this cluster.`,
				},
				resource.Attribute{
					Name:        "node_sha256_thumbprint",
					Description: `SHA256 of certificate thumbprint of this manager node.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ns_group",
			Category:         "Manager",
			ShortDescription: `A networking and security group data source.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ns",
				"group",
			},
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
			Category:         "Manager",
			ShortDescription: `A networking and security service data source.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ns",
				"service",
			},
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
			Type:             "nsxt_policy_bfd_profile",
			Category:         "Policy - Gateways and Routing",
			ShortDescription: `Policy BFD Profile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"gateways",
				"and",
				"routing",
				"bfd",
				"profile",
			},
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
			Type:             "nsxt_policy_certificate",
			Category:         "Policy - Certificates",
			ShortDescription: `Policy Certificate data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"certificates",
				"certificate",
			},
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
			Type:             "nsxt_policy_context_profile",
			Category:         "Policy - Firewall",
			ShortDescription: `Policy Context Profile Profile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"firewall",
				"context",
				"profile",
			},
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
			Type:             "nsxt_policy_dhcp_server",
			Category:         "Policy - DHCP",
			ShortDescription: `A policy DHCP server data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"dhcp",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of DHCP server to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of DHCP server to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
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
			Category:         "Policy - Fabric",
			ShortDescription: `A policy Edge Cluster data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"fabric",
				"edge",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the edge cluster to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the edge cluster to retrieve.`,
				},
				resource.Attribute{
					Name:        "site_path",
					Description: `(Optional) The path of the site which the Edge Cluster belongs to, this configuration is required for global manager only. ` + "`" + `path` + "`" + ` field of the existing ` + "`" + `nsxt_policy_site` + "`" + ` can be used here. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
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
			Category:         "Policy - Fabric",
			ShortDescription: `A policy Edge Node data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"fabric",
				"edge",
				"node",
			},
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
			Type:             "nsxt_policy_gateway_policy",
			Category:         "Policy - Firewall",
			ShortDescription: `A policy Gateway Policy data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"firewall",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the gateway policy to retrieve.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain of the policy, defaults to ` + "`" + `default` + "`" + `. Needs to be specified in VMC environment.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) Category of the policy to retrieve. May be useful to retrieve default policy.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the policy to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
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
			Category:         "Policy - Gateways and Routing",
			ShortDescription: `Policy GatewayQosProfile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"gateways",
				"and",
				"routing",
				"gateway",
				"qos",
				"profile",
			},
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
			Type:             "nsxt_policy_group",
			Category:         "Policy - Grouping and Tagging",
			ShortDescription: `Policy Group data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"grouping",
				"and",
				"tagging",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Group to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Group to retrieve.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain this Group belongs to. For VMware Cloud on AWS use ` + "`" + `cgw` + "`" + `. For Global Manager, please use site id for this field. If not specified, this field is default to ` + "`" + `default` + "`" + `. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
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
			Type:             "nsxt_policy_intrusion_service_profile",
			Category:         "Policy - Firewall",
			ShortDescription: `Policy Intrusion Service Profile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"firewall",
				"intrusion",
				"service",
				"profile",
			},
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
			Type:             "nsxt_policy_ip_block",
			Category:         "Policy - IPAM",
			ShortDescription: `Policy IP Block data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"ipam",
				"ip",
				"block",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of IP Block to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the IP Block to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
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
			Category:         "Policy - Segments",
			ShortDescription: `Policy IP Discovery Profile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"segments",
				"ip",
				"discovery",
				"profile",
			},
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
			Type:             "nsxt_policy_ip_pool",
			Category:         "Policy - IPAM",
			ShortDescription: `Policy IP Pool Config data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"ipam",
				"ip",
				"pool",
			},
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
			Category:         "Policy - Gateways and Routing",
			ShortDescription: `Policy IPv6 DAD Profile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"gateways",
				"and",
				"routing",
				"ipv6",
				"dad",
				"profile",
			},
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
			Type:             "nsxt_policy_ipv6_ndra_profile",
			Category:         "Policy - Gateways and Routing",
			ShortDescription: `Policy IPv6 NDRA Profile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"gateways",
				"and",
				"routing",
				"ipv6",
				"ndra",
				"profile",
			},
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
			Type:             "nsxt_policy_lb_app_profile",
			Category:         "Policy - Load Balancer",
			ShortDescription: `Policy Load Balancer Application Profile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"load",
				"balancer",
				"lb",
				"app",
				"profile",
			},
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
			Category:         "Policy - Load Balancer",
			ShortDescription: `Policy Load Balancer Client SSL Profile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"load",
				"balancer",
				"lb",
				"client",
				"ssl",
				"profile",
			},
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
			Category:         "Policy - Load Balancer",
			ShortDescription: `Policy Load Balancer Monitor data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"load",
				"balancer",
				"lb",
				"monitor",
			},
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
			Category:         "Policy - Load Balancer",
			ShortDescription: `Policy Load Balancer Persistence Profile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"load",
				"balancer",
				"lb",
				"persistence",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Profile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of Profile to retrieve.`,
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
			Category:         "Policy - Load Balancer",
			ShortDescription: `Policy Load Balancer Server SSL Profile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"load",
				"balancer",
				"lb",
				"server",
				"ssl",
				"profile",
			},
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
			Category:         "Policy - Segments",
			ShortDescription: `Policy MacDiscoveryProfile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"segments",
				"mac",
				"discovery",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Profile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of Profile to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
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
			Category:         "Policy - Segments",
			ShortDescription: `Policy QosProfile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"segments",
				"qos",
				"profile",
			},
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
			Type:             "nsxt_policy_realization_info",
			Category:         "Policy - Realization",
			ShortDescription: `A policy resource realization information.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"realization",
				"info",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The policy path of the resource.`,
				},
				resource.Attribute{
					Name:        "entity_type",
					Description: `(Optional) The entity type of realized resource. If not set, on of the realized resources of the policy resource will be retrieved.`,
				},
				resource.Attribute{
					Name:        "site_path",
					Description: `(Optional) The path of the site which the resource belongs to, this configuration is required for global manager only. ` + "`" + `path` + "`" + ` field of the existing ` + "`" + `nsxt_policy_site` + "`" + ` can be used here. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
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
			Type:             "nsxt_policy_security_policy",
			Category:         "Policy - Firewall",
			ShortDescription: `A policy Security Policy data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"firewall",
				"security",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Security Policy to retrieve.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Whether this is a default policy. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain of the policy, defaults to ` + "`" + `default` + "`" + `. Needs to be specified in VMC environment.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) Category of the policy to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name of the policy to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
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
			Type:             "nsxt_policy_segment_realization",
			Category:         "Policy - Realization",
			ShortDescription: `State of segment realization on hypervisors.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"realization",
				"segment",
			},
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
			Category:         "Policy - Segments",
			ShortDescription: `Policy SegmentSecurityProfile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"segments",
				"segment",
				"security",
				"profile",
			},
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
			Category:         "Policy - Firewall",
			ShortDescription: `A policy service data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"firewall",
				"service",
			},
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
			Type:             "nsxt_policy_site",
			Category:         "Policy - Fabric",
			ShortDescription: `Policy Site data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"fabric",
				"site",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Site to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Site to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. This attribute can serve as ` + "`" + `site_path` + "`" + ` field of ` + "`" + `nsxt_policy_transport_zone` + "`" + ` data source.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. This attribute can serve as ` + "`" + `site_path` + "`" + ` field of ` + "`" + `nsxt_policy_transport_zone` + "`" + ` data source.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_spoofguard_profile",
			Category:         "Policy - Segments",
			ShortDescription: `Policy SpoofGuardProfile data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"segments",
				"spoofguard",
				"profile",
			},
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
			Category:         "Policy - Gateways and Routing",
			ShortDescription: `A policy Tier-0 gateway data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"gateways",
				"and",
				"routing",
				"tier0",
				"gateway",
			},
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
					Description: `The path of the Edge cluster where this Tier-0 gateway is placed. This attribute is not set for NSX Global Manager, where gateway can spawn across multiple sites.`,
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
					Description: `The path of the Edge cluster where this Tier-0 gateway is placed. This attribute is not set for NSX Global Manager, where gateway can spawn across multiple sites.`,
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
			Category:         "Policy - Gateways and Routing",
			ShortDescription: `A policy Tier-1 gateway data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"gateways",
				"and",
				"routing",
				"tier1",
				"gateway",
			},
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
			Category:         "Policy - Fabric",
			ShortDescription: `A Policy Transport Zone data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"fabric",
				"transport",
				"zone",
			},
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
					Description: `(Optional) May be set together with ` + "`" + `transport_type` + "`" + ` in order to retrieve default Transport Zone for for this transport type.`,
				},
				resource.Attribute{
					Name:        "site_path",
					Description: `(Optional) The path of the site which the Transport Zone belongs to, this configuration is required for global manager only. ` + "`" + `path` + "`" + ` field of the existing ` + "`" + `nsxt_policy_site` + "`" + ` can be used here. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
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
			Category:         "Policy - Grouping and Tagging",
			ShortDescription: `A Discovered Policy Virtual Machine data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"grouping",
				"and",
				"tagging",
				"vm",
			},
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
			Category:         "Policy - Segments",
			ShortDescription: `Policy VNI Pool Config data source.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"segments",
				"vni",
				"pool",
			},
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
			Category:         "Manager",
			ShortDescription: `A switching profile data source.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"switching",
				"profile",
			},
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
			Category:         "Manager",
			ShortDescription: `A transport zone data source.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"transport",
				"zone",
			},
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

		"nsxt_certificate":                      0,
		"nsxt_edge_cluster":                     1,
		"nsxt_firewall_section":                 2,
		"nsxt_ip_pool":                          3,
		"nsxt_logical_tier0_router":             4,
		"nsxt_logical_tier1_router":             5,
		"nsxt_mac_pool":                         6,
		"nsxt_management_cluster":               7,
		"nsxt_ns_group":                         8,
		"nsxt_ns_service":                       9,
		"nsxt_policy_bfd_profile":               10,
		"nsxt_policy_certificate":               11,
		"nsxt_policy_context_profile":           12,
		"nsxt_policy_dhcp_server":               13,
		"nsxt_policy_edge_cluster":              14,
		"nsxt_policy_edge_node":                 15,
		"nsxt_policy_gateway_policy":            16,
		"nsxt_policy_gateway_qos_profile":       17,
		"nsxt_policy_group":                     18,
		"nsxt_policy_intrusion_service_profile": 19,
		"nsxt_policy_ip_block":                  20,
		"nsxt_policy_ip_discovery_profile":      21,
		"nsxt_policy_ip_pool":                   22,
		"nsxt_policy_ipv6_dad_profile":          23,
		"nsxt_policy_ipv6_ndra_profile":         24,
		"nsxt_policy_lb_app_profile":            25,
		"nsxt_policy_lb_client_ssl_profile":     26,
		"nsxt_policy_lb_monitor":                27,
		"nsxt_policy_lb_persistence_profile":    28,
		"nsxt_policy_lb_server_ssl_profile":     29,
		"nsxt_policy_mac_discovery_profile":     30,
		"nsxt_policy_qos_profile":               31,
		"nsxt_policy_realization_info":          32,
		"nsxt_policy_security_policy":           33,
		"nsxt_policy_segment_realization":       34,
		"nsxt_policy_segment_security_profile":  35,
		"nsxt_policy_service":                   36,
		"nsxt_policy_site":                      37,
		"nsxt_policy_spoofguard_profile":        38,
		"nsxt_policy_tier0_gateway":             39,
		"nsxt_policy_tier1_gateway":             40,
		"nsxt_policy_transport_zone":            41,
		"nsxt_policy_vm":                        42,
		"nsxt_policy_vni_pool":                  43,
		"nsxt_switching_profile":                44,
		"nsxt_transport_zone":                   45,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
