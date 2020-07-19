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
					Description: `(Required) IP set name for identifying the exact IP set ## Attribute Reference All the attributes defined in [` + "`" + `vcd_nsxv_ip_set` + "`" + `](/docs/providers/vcd/r/ipset.html) resource are available.`,
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
			ShortDescription: `Provides a vCloud Director Org VDC Network attached to an external one. This can be used to reference internal networks for vApps to connect.`,
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
			ShortDescription: `Provides a vCloud Director Org VDC isolated Network. This can be used to reference internal networks for vApps to connect.`,
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
			Type:             "vcd_network_routed",
			Category:         "Data Sources",
			ShortDescription: `Provides a vCloud Director Org VDC routed Network. This can be used to reference internal networks for vApps to connect.`,
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
			ShortDescription: `Provides a vCloud Director DNAT data source for advanced edge gateways (NSX-V). This can be used to read existing rule by ID and use its attributes in other resources.`,
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
			ShortDescription: `Provides a vCloud Director firewall rule data source for advanced edge gateways (NSX-V). This can be used to read existing rules by ID and use its attributes in other resources.`,
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
			Type:             "vcd_nsxv_snat",
			Category:         "Data Sources",
			ShortDescription: `Provides a vCloud Director SNAT data source for advanced edge gateways (NSX-V). This can be used to read existing rule by ID and use its attributes in other resources.`,
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
			Type:             "vcd_vapp",
			Category:         "Data Sources",
			ShortDescription: `Provides a vCloud Director vApp data source. This can be used to reference vApps.`,
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
			ShortDescription: `Provides a vCloud Director vApp network data source. This can be used to access a vApp network.`,
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
			ShortDescription: `Provides a data source for vCloud director Org network attached to vApp. This can be used to access vApp Org network.`,
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
			ShortDescription: `Provides a vCloud Director VM data source. This can be used to access VMs within a vApp.`,
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
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vm_affinity_rule",
			Category:         "Data Sources",
			ShortDescription: `Provides a vCloud Director VM affinity rule data source. This can be used to read VM affinity and anti-affinity rules.`,
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
	}

	dataSourcesMap = map[string]int{

		"vcd_catalog":            0,
		"vcd_catalog_item":       1,
		"vcd_catalog_media":      2,
		"vcd_edgegateway":        3,
		"vcd_external_network":   4,
		"vcd_independent_disk":   5,
		"vcd_nsxv_ip_set":        6,
		"vcd_lb_app_profile":     7,
		"vcd_lb_app_rule":        8,
		"vcd_lb_server_pool":     9,
		"vcd_lb_service_monitor": 10,
		"vcd_lb_virtual_server":  11,
		"vcd_network_direct":     12,
		"vcd_network_isolated":   13,
		"vcd_network_routed":     14,
		"vcd_nsxv_dhcp_relay":    15,
		"vcd_nsxv_dnat":          16,
		"vcd_nsxv_firewall_rule": 17,
		"vcd_nsxv_snat":          18,
		"vcd_org":                19,
		"vcd_org_vdc":            20,
		"vcd_vapp":               21,
		"vcd_vapp_network":       22,
		"vcd_vapp_org_network":   23,
		"vcd_vapp_vm":            24,
		"vcd_vm_affinity_rule":   25,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
