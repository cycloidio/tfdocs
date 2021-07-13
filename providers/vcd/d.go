package vcd

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vcd_catalog",
			Category:         "Data Sources",
			ShortDescription: `Provides a catalog data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional, but required if not set at provider level) Org name`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Catalog name (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Catalog description. ## Filter arguments (Supported in provider`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Catalog description. ## Filter arguments (Supported in provider`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_catalog_item",
			Category:         "Data Sources",
			ShortDescription: `Provides a catalog item data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional, but required if not set at provider level) Org name`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) Catalog name`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Catalog Item name (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Catalog item description.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Key value map of metadata. ## Filter arguments (Supported in provider`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Catalog item description.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Key value map of metadata. ## Filter arguments (Supported in provider`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_catalog_media",
			Category:         "Data Sources",
			ShortDescription: `Provides a catalog media data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) The name of the catalog where media file is`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Media name in catalog (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_edgegateway",
			Category:         "Data Sources",
			ShortDescription: `Provides an edge gateway data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the edge gateway (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the VDC belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC that owns the edge gateway. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_external_network",
			Category:         "Data Sources",
			ShortDescription: `Provides an external network data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) external network name ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Network friendly description`,
				},
				resource.Attribute{
					Name:        "ip_scope",
					Description: `A list of IP scopes for the network. See [IP Scope](/docs/providers/vcd/r/external_network.html#ipscope) for details.`,
				},
				resource.Attribute{
					Name:        "vsphere_network",
					Description: `A list of DV_PORTGROUP or NETWORK objects names that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server registered with the system. See [vSphere Network](/docs/providers/vcd/r/external_network.html#vspherenetwork) for details.`,
				},
				resource.Attribute{
					Name:        "retain_net_info_across_deployments",
					Description: `Specifies whether the network resources such as IP/MAC of router will be retained across deployments.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Network friendly description`,
				},
				resource.Attribute{
					Name:        "ip_scope",
					Description: `A list of IP scopes for the network. See [IP Scope](/docs/providers/vcd/r/external_network.html#ipscope) for details.`,
				},
				resource.Attribute{
					Name:        "vsphere_network",
					Description: `A list of DV_PORTGROUP or NETWORK objects names that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server registered with the system. See [vSphere Network](/docs/providers/vcd/r/external_network.html#vspherenetwork) for details.`,
				},
				resource.Attribute{
					Name:        "retain_net_info_across_deployments",
					Description: `Specifies whether the network resources such as IP/MAC of router will be retained across deployments.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_external_network_v2",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director External Network data source (version 2). New version of this data source uses new VCD API and is capable of creating NSX-T backed external networks as well as port group backed ones.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) external network name ## Attribute Reference All properties defined in [vcd_external_network_v2](/docs/providers/vcd/r/external_network_v2.html) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_global_role",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director global role data source . This can be used to read global roles.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the global role. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the global role`,
				},
				resource.Attribute{
					Name:        "bundle_key",
					Description: `Key used for internationalization.`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `List of rights assigned to this role`,
				},
				resource.Attribute{
					Name:        "publish_to_all_tenants",
					Description: `When true, publishes the global role to all tenants`,
				},
				resource.Attribute{
					Name:        "tenants",
					Description: `List of tenants to which this global role gets published. Ignored if ` + "`" + `publish_to_all_tenants` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether this global role is read-only ## More information See [Roles management](/docs/providers/vcd/guides/roles_management.html) for a broader description of how global roles and rights work together.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `A description of the global role`,
				},
				resource.Attribute{
					Name:        "bundle_key",
					Description: `Key used for internationalization.`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `List of rights assigned to this role`,
				},
				resource.Attribute{
					Name:        "publish_to_all_tenants",
					Description: `When true, publishes the global role to all tenants`,
				},
				resource.Attribute{
					Name:        "tenants",
					Description: `List of tenants to which this global role gets published. Ignored if ` + "`" + `publish_to_all_tenants` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether this global role is read-only ## More information See [Roles management](/docs/providers/vcd/guides/roles_management.html) for a broader description of how global roles and rights work together.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_independent_disk",
			Category:         "Data Sources",
			ShortDescription: `Provides a independent disk data source.`,
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
					Name:        "id",
					Description: `(Optional) Disk id or name is required. If both provided - Id is used. Id can be found by using import function [Listing independent disk IDs](/docs/providers/vcd/r/independent_disk.html#listing-independent-disk-ids)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Disk name.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
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
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_direct",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC Network attached to an external one. This can be used to reference internal networks for vApps to connect.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "external_network",
					Description: `The name of the external network.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Defines if this network is shared between multiple vDCs in the vOrg. ## Filter arguments (Supported in provider`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "external_network",
					Description: `The name of the external network.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Defines if this network is shared between multiple vDCs in the vOrg. ## Filter arguments (Supported in provider`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_isolated",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC isolated Network. This can be used to reference internal networks for vApps to connect.`,
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
					Name:        "name",
					Description: `(Required) A unique name for the network (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_isolated_v2",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC isolated Network data source to read data or reference existing network.`,
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
					Name:        "name",
					Description: `(Required) A unique name for the network (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Retrieves the data source using one or more filter parameters ## Attribute reference All attributes defined in [isolated network resource](/docs/providers/vcd/r/network_isolated_v2.html#attribute-reference) are supported. ## Filter arguments`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_routed",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC routed Network. This can be used to reference internal networks for vApps to connect.`,
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
					Name:        "name",
					Description: `(Required) A unique name for the network (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_routed_v2",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC routed Network data source to read data or reference existing network (backed by NSX-T or NSX-V).`,
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
					Name:        "name",
					Description: `(Required) A unique name for the network (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Retrieves the data source using one or more filter parameters ## Attribute reference All attributes defined in [routed network resource](/docs/providers/vcd/r/network_routed_v2.html#attribute-reference) are supported. ## Filter arguments`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_app_port_profile",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T Application Port Profiles. Application Port Profiles include a combination of a protocol and a port, or a group of ports, that is used for Firewall and NAT services on the Edge Gateway.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Unique name of existing Security Group.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Required) - ` + "`" + `SYSTEM` + "`" + `, ` + "`" + `PROVIDER` + "`" + `, or ` + "`" + `TENANT` + "`" + `. ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_app_port_profile` + "`" + `](/docs/providers/vcd/r/nsxt_app_port_profile.html) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_edge_cluster",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for available NSX-T Edge Clusters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) NSX-T Edge Cluster name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Edge Cluster description in NSX-T manager.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of nodes in Edge Cluster.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `Type of nodes in Edge Cluster.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `Deployment type of Edge Cluster.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_edgegateway",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director NSX-T edge gateway data source. This can be used to read NSX-T edge gateway configurations.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the edge gatweway belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC that owns the edge gateway. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) NSX-T Edge Gateway name. ## Attribute reference All properties defined in [vcd_nsxt_edgegateway](/docs/providers/vcd/r/nsxt_edgegateway.html) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_firewall",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T Firewall configuration of an Edge Gateway. Firewalls allow user to control the incoming and outgoing network traffic to and from an NSX-T Data Center Edge Gateway.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) The ID of the Edge Gateway (NSX-T only). Can be looked up using ` + "`" + `vcd_nsxt_edgegateway` + "`" + ` data source ## Attribute reference All properties defined in [vcd_nsxt_firewall](/docs/providers/vcd/r/nsxt_firewall.html) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_ip_set",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T IP Set. IP sets are groups of objects to which the firewall rules apply. Combining multiple objects into IP sets helps reduce the total number of firewall rules to be created.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) The ID of the edge gateway (NSX-T only). Can be looked up using`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Unique name of existing IP Set. ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_ip_set` + "`" + `](/docs/providers/vcd/r/nsxt_ip_set.html) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_ipsec_vpn_tunnel",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T IPsec VPN Tunnel. You can configure site-to-site connectivity between an NSX-T Data Center Edge Gateway and remote sites. The remote sites must use NSX-T Data Center, have third-party hardware routers, or VPN gateways that support IPSec.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) The ID of the edge gateway (NSX-T only). Can be looked up using ` + "`" + `vcd_nsxt_edgegateway` + "`" + ` data source`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Name of existing IPsec VPN Tunnel ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_ipsec_vpn_tunnel` + "`" + `](/docs/providers/vcd/r/nsxt_ipsec_vpn_tunnel.html) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_manager",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for available NSX-T manager.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) NSX-T manager name ## Attribute reference Only ID is set to be able and reference in other resources or data sources.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_nat_rule",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T NAT rules. Source NAT (SNAT) rules change the source IP address from a private to a public IP address. Destination NAT (DNAT) rules change the destination IP address from a public to a private IP address.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) The ID of the Edge Gateway (NSX-T only). Can be looked up using`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Name of existing NAT Rule. -> Name uniqueness is not enforced in NSX-T NAT rules, but for this data source to work properly names should be unique so that they can be distinguished. ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_nat_rule` + "`" + `](/docs/providers/vcd/r/nsxt_nat_rule.html) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_network_dhcp",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read DHCP pools for NSX-T Org VDC Routed network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "org_network_id",
					Description: `(Required) ID of parent Org VDC Routed network ## Attribute Reference All the attributes defined in [` + "`" + `vcd_nsxt_network_dhcp` + "`" + `](/docs/providers/vcd/r/nsxt_network_dhcp.html) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_network_imported",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC NSX-T Imported Network data source to read data or reference existing network.`,
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
					Name:        "name",
					Description: `(Required) A unique name for the network (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Retrieves the data source using one or more filter parameters ## Attribute reference All attributes defined in [imported network resource](/docs/providers/vcd/r/nsxt_network_imported.html#attribute-reference) are supported. ## Filter arguments`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_security_group",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to access NSX-T Security Group configuration. Security groups are groups of data center group networks to which distributed firewall rules apply. Grouping networks helps you to reduce the total number of distributed firewall rules to be created.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) The ID of the edge gateway (NSX-T only). Can be looked up using`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Unique name of existing Security Group. ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_security_group` + "`" + `](/docs/providers/vcd/r/nsxt_security_group.html) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_tier0_router",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for available NSX-T Tier-0 routers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) NSX-T Tier-0 router name.`,
				},
				resource.Attribute{
					Name:        "nsxt_manager_id",
					Description: `(Required) NSX-T manager should be referenced. ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "is_assigned",
					Description: `Boolean value reflecting if Tier-0 router is already consumed by external network.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_dhcp_relay",
			Category:         "Data Sources",
			ShortDescription: `Provides an NSX edge gateway DHCP relay configuration data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which DHCP relay is to be configured. ## Attribute Reference All the attributes defined in [` + "`" + `vcd_nsxv_dhcp_relay` + "`" + `](/docs/providers/vcd/r/nsxv_dhcp_relay.html) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_dnat",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director DNAT data source for advanced edge gateways (NSX-V). This can be used to read existing rule by ID and use its attributes in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which to apply the DNAT rule.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Required) ID of DNAT rule as shown in the UI. ## Attribute Reference All the attributes defined in ` + "`" + `vcd_nsxv_dnat` + "`" + ` resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_firewall_rule",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director firewall rule data source for advanced edge gateways (NSX-V). This can be used to read existing rules by ID and use its attributes in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which to apply the DNAT rule.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Required) ID of firewall rule (not UI number). See more information about firewall rule ID in ` + "`" + `vcd_nsxv_firewall_rule` + "`" + ` [import section](/docs/providers/vcd/r/nsxv_firewall_rule.html#listing-real-firewall-rule-ids). ## Attribute Reference All the attributes defined in [` + "`" + `vcd_nsxv_firewall_rule` + "`" + `](/docs/providers/vcd/r/nsxv_firewall_rule.html) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_ip_set",
			Category:         "Data Sources",
			ShortDescription: `Provides an IP set data source.`,
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
					Name:        "name",
					Description: `(Required) IP set name for identifying the exact IP set ## Attribute Reference All the attributes defined in [` + "`" + `vcd_nsxv_ip_set` + "`" + `](/docs/providers/vcd/r/nsxv_ip_set.html) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_snat",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director SNAT data source for advanced edge gateways (NSX-V). This can be used to read existing rule by ID and use its attributes in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which to apply the SNAT rule.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Required) ID of SNAT rule as shown in the UI. ## Attribute Reference All the attributes defined in ` + "`" + `vcd_nsxv_snat` + "`" + ` resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_org",
			Category:         "Data Sources",
			ShortDescription: `Provides an organization data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Org name ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `Org full name`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `True if this organization is enabled (allows login and all other operations).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Org description.`,
				},
				resource.Attribute{
					Name:        "deployed_vm_quota",
					Description: `Maximum number of virtual machines that can be deployed simultaneously by a member of this organization.`,
				},
				resource.Attribute{
					Name:        "stored_vm_quota",
					Description: `Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization.`,
				},
				resource.Attribute{
					Name:        "can_publish_catalogs",
					Description: `True if this organization is allowed to share catalogs.`,
				},
				resource.Attribute{
					Name:        "delay_after_power_on_seconds",
					Description: `Specifies this organization's default for virtual machine boot delay after power on.`,
				},
				resource.Attribute{
					Name:        "vapp_lease",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "vapp_template_lease",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "maximum_runtime_lease_in_sec",
					Description: `How long vApps can run before they are automatically stopped (in seconds)`,
				},
				resource.Attribute{
					Name:        "power_off_on_runtime_lease_expiration",
					Description: `When true, vApps are powered off when the runtime lease expires. When false, vApps are suspended when the runtime lease expires.`,
				},
				resource.Attribute{
					Name:        "maximum_storage_lease_in_sec",
					Description: `How long stopped vApps are available before being automatically cleaned up (in seconds)`,
				},
				resource.Attribute{
					Name:        "delete_on_storage_lease_expiration",
					Description: `If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted. <a id="vapp-template-lease"></a> ## vApp Template Lease The ` + "`" + `vapp_template_lease` + "`" + ` section contains lease parameters for vApp templates created in the current organization, as defined below:`,
				},
				resource.Attribute{
					Name:        "maximum_storage_lease_in_sec",
					Description: `How long vApp templates are available before being automatically cleaned up (in seconds)`,
				},
				resource.Attribute{
					Name:        "delete_on_storage_lease_expiration",
					Description: `If true, storage for a vAppTemplate is deleted when the vAppTemplate lease expires. If false, the storage is flagged for deletion, but not deleted`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "full_name",
					Description: `Org full name`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `True if this organization is enabled (allows login and all other operations).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Org description.`,
				},
				resource.Attribute{
					Name:        "deployed_vm_quota",
					Description: `Maximum number of virtual machines that can be deployed simultaneously by a member of this organization.`,
				},
				resource.Attribute{
					Name:        "stored_vm_quota",
					Description: `Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization.`,
				},
				resource.Attribute{
					Name:        "can_publish_catalogs",
					Description: `True if this organization is allowed to share catalogs.`,
				},
				resource.Attribute{
					Name:        "delay_after_power_on_seconds",
					Description: `Specifies this organization's default for virtual machine boot delay after power on.`,
				},
				resource.Attribute{
					Name:        "vapp_lease",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "vapp_template_lease",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "maximum_runtime_lease_in_sec",
					Description: `How long vApps can run before they are automatically stopped (in seconds)`,
				},
				resource.Attribute{
					Name:        "power_off_on_runtime_lease_expiration",
					Description: `When true, vApps are powered off when the runtime lease expires. When false, vApps are suspended when the runtime lease expires.`,
				},
				resource.Attribute{
					Name:        "maximum_storage_lease_in_sec",
					Description: `How long stopped vApps are available before being automatically cleaned up (in seconds)`,
				},
				resource.Attribute{
					Name:        "delete_on_storage_lease_expiration",
					Description: `If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted. <a id="vapp-template-lease"></a> ## vApp Template Lease The ` + "`" + `vapp_template_lease` + "`" + ` section contains lease parameters for vApp templates created in the current organization, as defined below:`,
				},
				resource.Attribute{
					Name:        "maximum_storage_lease_in_sec",
					Description: `How long vApp templates are available before being automatically cleaned up (in seconds)`,
				},
				resource.Attribute{
					Name:        "delete_on_storage_lease_expiration",
					Description: `If true, storage for a vAppTemplate is deleted when the vAppTemplate lease expires. If false, the storage is flagged for deletion, but not deleted`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_org_user",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Organization user data source. This can be used to read organization users.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the user belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the user. Required if ` + "`" + `user_id` + "`" + ` is not set.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Optional) The ID of the user. Required if ` + "`" + `name` + "`" + ` is not set. ## Attribute reference ~>`,
				},
				resource.Attribute{
					Name:        "provider_type",
					Description: `Identity provider type for this user.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `The role of the user.`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `The full name of the user.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `An optional description of the user.`,
				},
				resource.Attribute{
					Name:        "telephone",
					Description: `The Org User telephone number.`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `The Org User email address.`,
				},
				resource.Attribute{
					Name:        "instant_messaging",
					Description: `The Org User instant messaging.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `True if the user is enabled and can log in.`,
				},
				resource.Attribute{
					Name:        "is_group_role",
					Description: `True if this user has a group role.`,
				},
				resource.Attribute{
					Name:        "is_locked",
					Description: `If the user account has been locked due to too many invalid login attempts, the value will be true.`,
				},
				resource.Attribute{
					Name:        "deployed_vm_quota",
					Description: `Quota of vApps that this user can deploy. A value of 0 specifies an unlimited quota.`,
				},
				resource.Attribute{
					Name:        "stored_vm_quota",
					Description: `Quota of vApps that this user can store. A value of 0 specifies an unlimited quota.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Organization user`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_org_vdc",
			Category:         "Data Sources",
			ShortDescription: `Provides an organization VDC data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional, but required if not set at provider level) Org name`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Organization VDC name ## Attribute reference All attributes defined in [organization VDC resource](/docs/providers/vcd/r/org_vdc.html#attribute-reference) are supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_portgroup",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for available vCenter Port Groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Organization VDC name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) ` + "`" + `NETWORK` + "`" + ` for vSwitch port group or ` + "`" + `DV_PORTGROUP` + "`" + ` for distributed port group. ## Attribute reference Only ID is set to be able and reference in other resources or data sources.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_resource_list",
			Category:         "Data Sources",
			ShortDescription: `Provides lists of VCD resources`,
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
					Name:        "name",
					Description: `(Required) An unique name to identify the data source`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `(Computed) The list of requested resources in the chosen format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `(Computed) The list of requested resources in the chosen format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_resource_schema",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a vCD resource structure`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) An unique name to identify the data source`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Computed) The list of attributes for the resource. Each attribute is made of:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the attribute`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `an optional description of the attribute`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `whether the attribute is required`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `whether the attribute is optional`,
				},
				resource.Attribute{
					Name:        "computed",
					Description: `whether the attribute is computed`,
				},
				resource.Attribute{
					Name:        "sensitive",
					Description: `whether the attribute is sensitive`,
				},
				resource.Attribute{
					Name:        "block_attributes",
					Description: `(Computed) The list of compound attributes Each bock attribute is made of:`,
				},
				resource.Attribute{
					Name:        "nesting_type",
					Description: `(Computed) How the block is organized (one of ` + "`" + `NestingSet` + "`" + `, ` + "`" + `NestingList` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Computed) Same composition of the simple ` + "`" + `attributes` + "`" + ` above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "attributes",
					Description: `(Computed) The list of attributes for the resource. Each attribute is made of:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the attribute`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `an optional description of the attribute`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `whether the attribute is required`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `whether the attribute is optional`,
				},
				resource.Attribute{
					Name:        "computed",
					Description: `whether the attribute is computed`,
				},
				resource.Attribute{
					Name:        "sensitive",
					Description: `whether the attribute is sensitive`,
				},
				resource.Attribute{
					Name:        "block_attributes",
					Description: `(Computed) The list of compound attributes Each bock attribute is made of:`,
				},
				resource.Attribute{
					Name:        "nesting_type",
					Description: `(Computed) How the block is organized (one of ` + "`" + `NestingSet` + "`" + `, ` + "`" + `NestingList` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Computed) Same composition of the simple ` + "`" + `attributes` + "`" + ` above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_right",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Organization Right data source. This can be used to read existing rights`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the right. ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the right`,
				},
				resource.Attribute{
					Name:        "category_id",
					Description: `The ID of the category for this right`,
				},
				resource.Attribute{
					Name:        "bundle_key",
					Description: `Key used for internationalization`,
				},
				resource.Attribute{
					Name:        "right type",
					Description: `Type of the right (VIEW or MODIFY)`,
				},
				resource.Attribute{
					Name:        "implied_rights",
					Description: `List of rights that are implied with this one ## More information See [Roles management](/docs/providers/vcd/guides/roles_management.html) for a broader description of how roles and rights work together.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_rights_bundle",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director rights bundle data source. This can be used to read rights bundles.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the rights bundle. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the rights bundle`,
				},
				resource.Attribute{
					Name:        "bundle_key",
					Description: `Key used for internationalization.`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `List of rights assigned to this role`,
				},
				resource.Attribute{
					Name:        "publish_to_all_tenants",
					Description: `When true, publishes the rights bundle to all tenants`,
				},
				resource.Attribute{
					Name:        "tenants",
					Description: `List of tenants to which this rights bundle gets published. Ignored if ` + "`" + `publish_to_all_tenants` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether this rights bundle is read-only ## More information See [Roles management](/docs/providers/vcd/guides/roles_management.html) for a broader description of how rights bundles and rights work together.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `A description of the rights bundle`,
				},
				resource.Attribute{
					Name:        "bundle_key",
					Description: `Key used for internationalization.`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `List of rights assigned to this role`,
				},
				resource.Attribute{
					Name:        "publish_to_all_tenants",
					Description: `When true, publishes the rights bundle to all tenants`,
				},
				resource.Attribute{
					Name:        "tenants",
					Description: `List of tenants to which this rights bundle gets published. Ignored if ` + "`" + `publish_to_all_tenants` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether this rights bundle is read-only ## More information See [Roles management](/docs/providers/vcd/guides/roles_management.html) for a broader description of how rights bundles and rights work together.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_role",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director role. This can be used to read roles.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the role. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether this role is read-only`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the role`,
				},
				resource.Attribute{
					Name:        "bundle_key",
					Description: `Key used for internationalization.`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `List of rights assigned to this role ## More information See [Roles management](/docs/providers/vcd/guides/roles_management.html) for a broader description of how roles and rights work together.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether this role is read-only`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the role`,
				},
				resource.Attribute{
					Name:        "bundle_key",
					Description: `Key used for internationalization.`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `List of rights assigned to this role ## More information See [Roles management](/docs/providers/vcd/guides/roles_management.html) for a broader description of how roles and rights work together.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_storage_profile",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for VDC storage profile.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Storage profile name. ## Attribute reference This data source exports only ` + "`" + `id` + "`" + ` field.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director vApp data source. This can be used to reference vApps.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the vApp`,
				},
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The vApp Hyper Reference`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Key value map of metadata to assign to this vApp. Key and value can be any string.`,
				},
				resource.Attribute{
					Name:        "guest_properties",
					Description: `Key value map of vApp guest properties.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The vApp status as a numeric code`,
				},
				resource.Attribute{
					Name:        "status_text",
					Description: `The vApp status as text.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_network",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director vApp network data source. This can be used to access a vApp network.`,
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
					Name:        "vapp_name",
					Description: `(Required) The vApp name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the vApp network, unique within the vApp ## Attribute reference All attributes defined in [` + "`" + `vcd_vapp_network` + "`" + `](/docs/providers/vcd/r/vapp_network.html#attribute-reference) are supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_org_network",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for VMware Cloud Director Org network attached to vApp. This can be used to access vApp Org network.`,
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
					Name:        "vapp_name",
					Description: `(Required) The vApp name.`,
				},
				resource.Attribute{
					Name:        "org_network_name",
					Description: `(Required) A name for the vApp Org network, unique within the vApp. ## Attribute reference All attributes defined in [` + "`" + `vcd_vapp_org_network` + "`" + `](/docs/providers/vcd/r/vapp_org_network.html#attribute-reference) are supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_vm",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director VM data source. This can be used to access VMs within a vApp.`,
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
					Name:        "vapp_name",
					Description: `(Required) The vApp this VM belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the VM, unique within the vApp`,
				},
				resource.Attribute{
					Name:        "network_dhcp_wait_seconds",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "computer_name",
					Description: `Computer name to assign to this virtual machine.`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `The catalog name in which to find the given vApp Template`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `The name of the vApp Template to use`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of RAM (in MB) allocated to the VM`,
				},
				resource.Attribute{
					Name:        "cpus",
					Description: `The number of virtual CPUs allocated to the VM`,
				},
				resource.Attribute{
					Name:        "cpu_cores",
					Description: `The number of cores per socket`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Key value map of metadata assigned to this VM`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Independent disk attachment configuration.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `A block defining a network interface. Multiple can be used.`,
				},
				resource.Attribute{
					Name:        "guest_properties",
					Description: `Key value map of guest properties`,
				},
				resource.Attribute{
					Name:        "expose_hardware_virtualization",
					Description: `Expose hardware-assisted CPU virtualization to guest OS`,
				},
				resource.Attribute{
					Name:        "internal_disk",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "hardware_version",
					Description: `(`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vcenter",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for vCenter server attached to VCD.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) vCenter name ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "connection_status",
					Description: `vCenter connection status (e.g. ` + "`" + `CONNECTED` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `Boolean value if vCenter is enabled.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `vCenter status (e.g. ` + "`" + `READY` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vcenter_host",
					Description: `Hostname of configured vCenter.`,
				},
				resource.Attribute{
					Name:        "vcenter_version",
					Description: `vCenter version (e.g. ` + "`" + `6.7.0` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vm",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director VM data source. This can be used to access standalone VMs.`,
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
					Name:        "name",
					Description: `(Required) A name or ID for the standalone VM in VDC ## Attributes reference This data source provides all attributes available for ` + "`" + `vcd_vapp_vm` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vm_affinity_rule",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director VM affinity rule data source. This can be used to read VM affinity and anti-affinity rules.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of VM affinity rule. Needed if we don't provide ` + "`" + `rule_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Is the ID of the affinity rule. It's the preferred way to retrieve the affinity rule, especially if the rule name could have duplicates ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "polarity",
					Description: `One of ` + "`" + `Affinity` + "`" + ` or ` + "`" + `Anti-Affinity` + "`" + `. This property cannot be changed. Once created, if we need to change polarity, we need to remove the rule and create a new one.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vm_sizing_policy",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director VM sizing policy data source. This can be used to read VM sizing policy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name VM sizing policy All arguments defined in [` + "`" + `vcd_vm_sizing_policy` + "`" + `](/docs/providers/vcd/r/vm_sizing_policy.html#argument-reference) are supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"vcd_catalog":               0,
		"vcd_catalog_item":          1,
		"vcd_catalog_media":         2,
		"vcd_edgegateway":           3,
		"vcd_external_network":      4,
		"vcd_external_network_v2":   5,
		"vcd_global_role":           6,
		"vcd_independent_disk":      7,
		"vcd_lb_app_profile":        8,
		"vcd_lb_app_rule":           9,
		"vcd_lb_server_pool":        10,
		"vcd_lb_service_monitor":    11,
		"vcd_lb_virtual_server":     12,
		"vcd_network_direct":        13,
		"vcd_network_isolated":      14,
		"vcd_network_isolated_v2":   15,
		"vcd_network_routed":        16,
		"vcd_network_routed_v2":     17,
		"vcd_nsxt_app_port_profile": 18,
		"vcd_nsxt_edge_cluster":     19,
		"vcd_nsxt_edgegateway":      20,
		"vcd_nsxt_firewall":         21,
		"vcd_nsxt_ip_set":           22,
		"vcd_nsxt_ipsec_vpn_tunnel": 23,
		"vcd_nsxt_manager":          24,
		"vcd_nsxt_nat_rule":         25,
		"vcd_nsxt_network_dhcp":     26,
		"vcd_nsxt_network_imported": 27,
		"vcd_nsxt_security_group":   28,
		"vcd_nsxt_tier0_router":     29,
		"vcd_nsxv_dhcp_relay":       30,
		"vcd_nsxv_dnat":             31,
		"vcd_nsxv_firewall_rule":    32,
		"vcd_nsxv_ip_set":           33,
		"vcd_nsxv_snat":             34,
		"vcd_org":                   35,
		"vcd_org_user":              36,
		"vcd_org_vdc":               37,
		"vcd_portgroup":             38,
		"vcd_resource_list":         39,
		"vcd_resource_schema":       40,
		"vcd_right":                 41,
		"vcd_rights_bundle":         42,
		"vcd_role":                  43,
		"vcd_storage_profile":       44,
		"vcd_vapp":                  45,
		"vcd_vapp_network":          46,
		"vcd_vapp_org_network":      47,
		"vcd_vapp_vm":               48,
		"vcd_vcenter":               49,
		"vcd_vm":                    50,
		"vcd_vm_affinity_rule":      51,
		"vcd_vm_sizing_policy":      52,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
