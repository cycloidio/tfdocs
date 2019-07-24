package vcd

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vcd_catalog",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director catalog resource. This can be used to create and delete a catalog.`,
			Description:      ``,
			Keywords: []string{
				"catalog",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Catalog name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) - Description of catalog`,
				},
				resource.Attribute{
					Name:        "delete_recursive",
					Description: `(Required) - When destroying use delete_recursive=True to remove the catalog and any objects it contains that are in a state that normally allows removal`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_catalog_item",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director catalog item resource. This can be used to upload and delete OVA file inside a catalog.`,
			Description:      ``,
			Keywords: []string{
				"catalog",
				"item",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) The name of the catalog where to upload OVA file`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Item name in catalog`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) - Description of item`,
				},
				resource.Attribute{
					Name:        "ova_path",
					Description: `(Required) - Absolute or relative path to file to upload`,
				},
				resource.Attribute{
					Name:        "upload_piece_size",
					Description: `(Optional) - Size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB.`,
				},
				resource.Attribute{
					Name:        "show_upload_progress",
					Description: `(Optional) - Default false. Allows to see upload progress`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_catalog_media",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director media resource. This can be used to upload and delete media (ISO) file inside a catalog.`,
			Description:      ``,
			Keywords: []string{
				"catalog",
				"media",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) The name of the catalog where to upload media file`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Media file name in catalog`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) - Description of media file`,
				},
				resource.Attribute{
					Name:        "media_path",
					Description: `(Required) - Absolute or relative path to file to upload`,
				},
				resource.Attribute{
					Name:        "upload_piece_size",
					Description: `(Optional) - size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB.`,
				},
				resource.Attribute{
					Name:        "show_upload_progress",
					Description: `(Optional) - Default false. Allows to see upload progress`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_dnat",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director DNAT resource. This can be used to create, modify, and delete destination NATs to map external IPs to a VM.`,
			Description:      ``,
			Keywords: []string{
				"dnat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which to apply the DNAT`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `(Required) One of the external IPs available on your Edge Gateway`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port number to map. -1 translates to "any"`,
				},
				resource.Attribute{
					Name:        "translated_port",
					Description: `(Optional) The port number to map`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `(Required) The IP of the VM to map to`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "icmp_sub_type",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_edgegateway_vpn",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director IPsec VPN. This can be used to create, modify, and delete VPN settings and rules.`,
			Description:      ``,
			Keywords: []string{
				"edgegateway",
				"vpn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which to apply the Firewall Rules`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VPN`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for the VPN`,
				},
				resource.Attribute{
					Name:        "encryption_protocol",
					Description: `(Required) - E.g. ` + "`" + `AES256` + "`" + ``,
				},
				resource.Attribute{
					Name:        "local_ip_address",
					Description: `(Required) - Local IP Address`,
				},
				resource.Attribute{
					Name:        "local_id",
					Description: `(Required) - Local ID`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Required) - The MTU setting`,
				},
				resource.Attribute{
					Name:        "peer_ip_address",
					Description: `(Required) - Peer IP Address`,
				},
				resource.Attribute{
					Name:        "peer_id",
					Description: `(Required) - Peer ID`,
				},
				resource.Attribute{
					Name:        "shared_secret",
					Description: `(Required) - Shared Secret`,
				},
				resource.Attribute{
					Name:        "local_subnets",
					Description: `(Required) - List of Local Subnets see [Local Subnets](#localsubnets) below for details.`,
				},
				resource.Attribute{
					Name:        "peer_subnets",
					Description: `(Required) - List of Peer Subnets see [Peer Subnets](#peersubnets) below for details.`,
				},
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "local_subnet_name",
					Description: `(Required) Name of the local subnet`,
				},
				resource.Attribute{
					Name:        "local_subnet_gateway",
					Description: `(Required) Gateway of the local subnet`,
				},
				resource.Attribute{
					Name:        "local_subnet_mask",
					Description: `(Required) Subnet mask of the local subnet <a id="peersubnets"></a> ## Peer Subnets Each Peer Subnet supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "peer_subnet_name",
					Description: `(Required) Name of the peer subnet`,
				},
				resource.Attribute{
					Name:        "peer_subnet_gateway",
					Description: `(Required) Gateway of the peer subnet`,
				},
				resource.Attribute{
					Name:        "peer_subnet_mask",
					Description: `(Required) Subnet mask of the peer subnet`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_external_network",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director external network resource. This can be used to create and delete external networks.`,
			Description:      ``,
			Keywords: []string{
				"external",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Network friendly description`,
				},
				resource.Attribute{
					Name:        "ip_scope",
					Description: `(Required) A list of IP scopes for the network. See [IP Scope](#ipscope) below for details.`,
				},
				resource.Attribute{
					Name:        "vsphere_network",
					Description: `(Required) A list of DV_PORTGROUP or NETWORK objects names that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server registered with the system. See [vSphere Network](#vspherenetwork) below for details.`,
				},
				resource.Attribute{
					Name:        "retain_net_info_across_deployments",
					Description: `(Optional) Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false. <a id="ipscope"></a> ## IP Scope`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Required) Gateway of the network`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Required) Network mask`,
				},
				resource.Attribute{
					Name:        "dns1",
					Description: `(Optional) Primary DNS server`,
				},
				resource.Attribute{
					Name:        "dns2",
					Description: `(Optional) Secondary DNS server`,
				},
				resource.Attribute{
					Name:        "static_ip_pool",
					Description: `(Required) IP ranges used for static pool allocation in the network. See [IP Pool](#ip-pool) below for details. <a id="ip-pool"></a> ## IP Pool`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) Start address of the IP range`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) End address of the IP range <a id="vspherenetwork"></a> ## vSphere Network`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Port group name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The vSphere type of the object. One of: DV_PORTGROUP (distributed virtual port group), NETWORK (standard switch port group)`,
				},
				resource.Attribute{
					Name:        "vcenter",
					Description: `(Required) The vCenter server name`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_firewall_rules",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Firewall resource. This can be used to create, modify, and delete firewall settings and rules.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"rules",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which to apply the Firewall Rules`,
				},
				resource.Attribute{
					Name:        "default_action",
					Description: `(Required) Either "allow" or "drop". Specifies what to do should none of the rules match`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Configures a firewall rule; see [Rules](#rules) below for details.`,
				},
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the fireall rule`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) Specifies what to do when this rule is matched. Either "allow" or "drop"`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol to match. One of "tcp", "udp", "icmp" or "any"`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Required) The destination port to match. Either a port number or "any"`,
				},
				resource.Attribute{
					Name:        "destination_ip",
					Description: `(Required) The destination IP to match. Either an IP address, IP range or "any"`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Required) The source port to match. Either a port number or "any"`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `(Required) The source IP to match. Either an IP address, IP range or "any"`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_independent_disk",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director independent disk resource. This can be used to create and delete independent disks.`,
			Description:      ``,
			Keywords: []string{
				"independent",
				"disk",
			},
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
					Description: `(Required) Disk name`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) - Size of disk in MB`,
				},
				resource.Attribute{
					Name:        "bus_type",
					Description: `(Optional) - Disk bus type. Values can be: IDE, SCSI, SATA`,
				},
				resource.Attribute{
					Name:        "bus_sub_type",
					Description: `(Optional) - Disk bus subtype. Values can be: "IDE" for IDE. buslogic, lsilogic, lsilogicsas, VirtualSCSI for SCSI and ahci for SATA`,
				},
				resource.Attribute{
					Name:        "storage_profile",
					Description: `(Optional) - The name of storage profile where disk will be created`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_inserted_media",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director resource for inserting or ejecting media (ISO) file for the VM. Create this resource for inserting the media, and destroy it for ejecting.`,
			Description:      ``,
			Keywords: []string{
				"inserted",
				"media",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) The name of the catalog where to find media file`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Media file name in catalog which will be inserted to VM`,
				},
				resource.Attribute{
					Name:        "vapp_name",
					Description: `(Required) - The name of vApp to find`,
				},
				resource.Attribute{
					Name:        "vm_name",
					Description: `(Required) - The name of VM to be used to insert media file`,
				},
				resource.Attribute{
					Name:        "eject_force",
					Description: `(Optional;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Org VDC Network. This can be used to create, modify, and delete internal networks for vApps to connect.`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) The netmask for the new network. Defaults to ` + "`" + `255.255.255.0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dns1",
					Description: `(Optional) First DNS server to use. Defaults to ` + "`" + `8.8.8.8` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dns2",
					Description: `(Optional) Second DNS server to use. Defaults to ` + "`" + `8.8.4.4` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `(Optional) A FQDN for the virtual machines on this network`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Defines if this network is shared between multiple vDCs in the vOrg. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dhcp_pool",
					Description: `(Optional) A range of IPs to issue to virtual machines that don't have a static IP; see [IP Pools](#ip-pools) below for details.`,
				},
				resource.Attribute{
					Name:        "static_ip_pool",
					Description: `(Optional) A range of IPs permitted to be used as static IPs for virtual machines; see [IP Pools](#ip-pools) below for details. <a id="ip-pools"></a> ## IP Pools Static IP Pools and DHCP Pools support the following attributes:`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) The first address in the IP Range`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) The final address in the IP Range DHCP Pools additionally support the following attributes:`,
				},
				resource.Attribute{
					Name:        "default_lease_time",
					Description: `(Optional) The default DHCP lease time to use. Defaults to ` + "`" + `3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_lease_time",
					Description: `(Optional) The maximum DHCP lease time to use. Defaults to ` + "`" + `7200` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_direct",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Org VDC Network attached to an external one. This can be used to create, modify, and delete internal networks for vApps to connect.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"direct",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network`,
				},
				resource.Attribute{
					Name:        "external_network",
					Description: `(Required) The name of the external network.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Defines if this network is shared between multiple vDCs in the vOrg. Defaults to ` + "`" + `false` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_isolated",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Org VDC isolated Network. This can be used to create, modify, and delete internal networks for vApps to connect.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"isolated",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) The netmask for the new network. Defaults to ` + "`" + `255.255.255.0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dns1",
					Description: `(Optional) First DNS server to use. Defaults to ` + "`" + `8.8.8.8` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dns2",
					Description: `(Optional) Second DNS server to use. Defaults to ` + "`" + `8.8.4.4` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `(Optional) A FQDN for the virtual machines on this network`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Defines if this network is shared between multiple vDCs in the vOrg. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dhcp_pool",
					Description: `(Optional) A range of IPs to issue to virtual machines that don't have a static IP; see [IP Pools](#ip-pools) below for details.`,
				},
				resource.Attribute{
					Name:        "static_ip_pool",
					Description: `(Optional) A range of IPs permitted to be used as static IPs for virtual machines; see [IP Pools](#ip-pools) below for details. <a id="ip-pools"></a> ## IP Pools Static IP Pools and DHCP Pools support the following attributes:`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) The first address in the IP Range`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) The final address in the IP Range DHCP Pools additionally support the following attributes:`,
				},
				resource.Attribute{
					Name:        "default_lease_time",
					Description: `(Optional) The default DHCP lease time to use. Defaults to ` + "`" + `3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_lease_time",
					Description: `(Optional) The maximum DHCP lease time to use. Defaults to ` + "`" + `7200` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_routed",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Org VDC routed Network. This can be used to create, modify, and delete internal networks for vApps to connect.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"routed",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) The netmask for the new network. Defaults to ` + "`" + `255.255.255.0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dns1",
					Description: `(Optional) First DNS server to use. Defaults to ` + "`" + `8.8.8.8` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dns2",
					Description: `(Optional) Second DNS server to use. Defaults to ` + "`" + `8.8.4.4` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `(Optional) A FQDN for the virtual machines on this network`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Defines if this network is shared between multiple vDCs in the vOrg. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dhcp_pool",
					Description: `(Optional) A range of IPs to issue to virtual machines that don't have a static IP; see [IP Pools](#ip-pools) below for details.`,
				},
				resource.Attribute{
					Name:        "static_ip_pool",
					Description: `(Optional) A range of IPs permitted to be used as static IPs for virtual machines; see [IP Pools](#ip-pools) below for details. <a id="ip-pools"></a> ## IP Pools Static IP Pools and DHCP Pools support the following attributes:`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) The first address in the IP Range`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) The final address in the IP Range DHCP Pools additionally support the following attributes:`,
				},
				resource.Attribute{
					Name:        "default_lease_time",
					Description: `(Optional) The default DHCP lease time to use. Defaults to ` + "`" + `3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_lease_time",
					Description: `(Optional) The maximum DHCP lease time to use. Defaults to ` + "`" + `7200` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_org",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Organization resource. This can be used to create and delete an organization.`,
			Description:      ``,
			Keywords: []string{
				"org",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Org name`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `(Required) Org full name`,
				},
				resource.Attribute{
					Name:        "delete_recursive",
					Description: `(Required) - pass ` + "`" + `delete_recursive` + "`" + `=true as query parameter to remove an organization or VDC and any objects it contains that are in a state that normally allows removal.`,
				},
				resource.Attribute{
					Name:        "delete_force",
					Description: `(Required) - pass ` + "`" + `delete_force=true` + "`" + ` and ` + "`" + `delete_recursive=true` + "`" + ` to remove an organization or VDC and any objects it contains, regardless of their state.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) - True if this organization is enabled (allows login and all other operations). Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) - Org description. Default is empty.`,
				},
				resource.Attribute{
					Name:        "deployed_vm_quota",
					Description: `(Optional) - Maximum number of virtual machines that can be deployed simultaneously by a member of this organization. Default is unlimited (-1)`,
				},
				resource.Attribute{
					Name:        "stored_vm_quota",
					Description: `(Optional) - Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization. Default is unlimited (-1)`,
				},
				resource.Attribute{
					Name:        "can_publish_catalogs",
					Description: `(Optional) - True if this organization is allowed to share catalogs. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delay_after_power_on_seconds",
					Description: `(Optional) - Specifies this organization's default for virtual machine boot delay after power on. Default is ` + "`" + `0` + "`" + `. ## Sources`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_org_vdc",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director Organization VDC resource. This can be used to create and delete a Organization VDC.`,
			Description:      ``,
			Keywords: []string{
				"org",
				"vdc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) Organization to create the VDC in, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) VDC name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) VDC friendly description`,
				},
				resource.Attribute{
					Name:        "provider_vdc_name",
					Description: `(Required) A name of the Provider VDC from which this organization VDC is provisioned.`,
				},
				resource.Attribute{
					Name:        "allocation_model",
					Description: `(Required) The allocation model used by this VDC; must be one of {AllocationVApp ("Pay as you go"), AllocationPool ("Allocation pool"), ReservationPool ("Reservation pool")}`,
				},
				resource.Attribute{
					Name:        "compute_capacity",
					Description: `(Required) The compute capacity allocated to this VDC. See [Compute Capacity](#computecapacity) below for details.`,
				},
				resource.Attribute{
					Name:        "nic_quota",
					Description: `(Optional) Maximum number of virtual NICs allowed in this VDC. Defaults to 0, which specifies an unlimited number.`,
				},
				resource.Attribute{
					Name:        "network_quota",
					Description: `(Optional) Maximum number of network objects that can be deployed in this VDC. Defaults to 0, which means no networks can be deployed.`,
				},
				resource.Attribute{
					Name:        "vm_quota",
					Description: `(Optional) The maximum number of VMs that can be created in this VDC. Includes deployed and undeployed VMs in vApps and vApp templates. Defaults to 0, which specifies an unlimited number.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) True if this VDC is enabled for use by the organization VDCs. Default is true.`,
				},
				resource.Attribute{
					Name:        "storage_profile",
					Description: `(Required) Storage profiles supported by this VDC. See [Storage Profile](#storageprofile) below for details.`,
				},
				resource.Attribute{
					Name:        "memory_guaranteed",
					Description: `(Optional) Percentage of allocated memory resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when AllocationModel is AllocationVApp or AllocationPool. If left empty, vCD sets a value.`,
				},
				resource.Attribute{
					Name:        "cpu_guaranteed",
					Description: `(Optional) Percentage of allocated CPU resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when AllocationModel is AllocationVApp or AllocationPool. If left empty, vCD sets a value.`,
				},
				resource.Attribute{
					Name:        "cpu_speed",
					Description: `(Optional) Specifies the clock frequency, in Megahertz, for any virtual CPU that is allocated to a VM. A VM with 2 vCPUs will consume twice as much of this value. Ignored for ReservationPool. Required when AllocationModel is AllocationVApp or AllocationPool, and may not be less than 256 MHz. Defaults to 1000 MHz if value isn't provided.`,
				},
				resource.Attribute{
					Name:        "enable_thin_provisioning",
					Description: `(Optional) Boolean to request thin provisioning. Request will be honored only if the underlying data store supports it. Thin provisioning saves storage space by committing it on demand. This allows over-allocation of storage.`,
				},
				resource.Attribute{
					Name:        "enable_fast_provisioning",
					Description: `(Optional) Request fast provisioning. Request will be honored only if the underlying datastore supports it. Fast provisioning can reduce the time it takes to create virtual machines by using vSphere linked clones. If you disable fast provisioning, all provisioning operations will result in full clones.`,
				},
				resource.Attribute{
					Name:        "network_pool_name",
					Description: `(Optional) Reference to a network pool in the Provider VDC. Required if this VDC will contain routed or isolated networks.`,
				},
				resource.Attribute{
					Name:        "allow_over_commit",
					Description: `(Optional) Set to false to disallow creation of the VDC if the AllocationModel is AllocationPool or ReservationPool and the ComputeCapacity you specified is greater than what the backing Provider VDC can supply. Default is true.`,
				},
				resource.Attribute{
					Name:        "enable_vm_discovery",
					Description: `(Optional) If true, discovery of vCenter VMs is enabled for resource pools backing this VDC. If false, discovery is disabled. If left unspecified, the actual behaviour depends on enablement at the organization level and at the system level.`,
				},
				resource.Attribute{
					Name:        "delete_force",
					Description: `(Required) When destroying use ` + "`" + `delete_force=True` + "`" + ` to remove a VDC and any objects it contains, regardless of their state.`,
				},
				resource.Attribute{
					Name:        "delete_recursive",
					Description: `(Required) When destroying use ` + "`" + `delete_recursive=True` + "`" + ` to remove the VDC and any objects it contains that are in a state that normally allows removal. <a id="storageprofile"></a> ## Storage Profile`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Provider VDC storage profile.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) True if this storage profile is enabled for use in the VDC. Default is true.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Required) Maximum number of MB allocated for this storage profile. A value of 0 specifies unlimited MB.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Required) True if this is default storage profile for this VDC. The default storage profile is used when an object that can specify a storage profile is created with no storage profile specified. <a id="computecapacity"></a> ## Compute Capacity Capacity must be specified twice, once for ` + "`" + `memory` + "`" + ` and another for ` + "`" + `cpu` + "`" + `. Each has the same structure:`,
				},
				resource.Attribute{
					Name:        "allocated",
					Description: `(Optional) Capacity that is committed to be available. Value in MB or MHz. Used with AllocationPool ("Allocation pool") and ReservationPool ("Reservation pool").`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Capacity limit relative to the value specified for Allocation. It must not be less than that value. If it is greater than that value, it implies over provisioning. A value of 0 specifies unlimited units. Value in MB or MHz. Used with AllocationVApp ("Pay as you go").`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_snat",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director SNAT resource. This can be used to create, modify, and delete source NATs to allow vApps to send external traffic.`,
			Description:      ``,
			Keywords: []string{
				"snat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which to apply the SNAT`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `(Required) One of the external IPs available on your Edge Gateway`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `(Required) The IP or IP Range of the VM(s) to map from`,
				},
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director vApp resource. This can be used to create, modify, and delete vApps.`,
			Description:      ``,
			Keywords: []string{
				"vapp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the vApp`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `(Optional) The catalog name in which to find the given vApp Template`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) The name of the vApp Template to use`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) The amount of RAM (in MB) to allocate to the vApp`,
				},
				resource.Attribute{
					Name:        "cpus",
					Description: `(Optional) The number of virtual CPUs to allocate to the vApp`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `(Optional) Name of the network this vApp should join`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) The IP to assign to this vApp. Must be an IP address or one of dhcp, allocated or none. If given the address must be within the ` + "`" + `static_ip_pool` + "`" + ` set for the network. If left blank, and the network has ` + "`" + `dhcp_pool` + "`" + ` set with at least one available IP then this will be set with DHCP.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Key value map of metadata to assign to this vApp. Key and value can be any string. (Since`,
				},
				resource.Attribute{
					Name:        "ovf",
					Description: `(Optional) Key value map of ovf parameters to assign to VM product section`,
				},
				resource.Attribute{
					Name:        "power_on",
					Description: `(Optional) A boolean value stating if this vApp should be powered on. Default is ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "accept_all_eulas",
					Description: `(Optional;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_network",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director vApp isolated Network. This can be used to create and delete internal networks for vApps to connect.`,
			Description:      ``,
			Keywords: []string{
				"vapp",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network.`,
				},
				resource.Attribute{
					Name:        "vapp_name",
					Description: `(Required) The vApp this VM should belong to.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) The netmask for the new network. Default is ` + "`" + `255.255.255.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dns1",
					Description: `(Optional) First DNS server to use. Default is ` + "`" + `8.8.8.8` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dns2",
					Description: `(Optional) Second DNS server to use. Default is ` + "`" + `8.8.4.4` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `(Optional) A FQDN for the virtual machines on this network.`,
				},
				resource.Attribute{
					Name:        "static_ip_pool",
					Description: `(Optional) A range of IPs permitted to be used as static IPs for virtual machines; see [IP Pools](#ip-pools) below for details.`,
				},
				resource.Attribute{
					Name:        "dhcp_pool",
					Description: `(Optional) A range of IPs to issue to virtual machines that don't have a static IP; see [IP Pools](#ip-pools) below for details. <a id="ip-pools"></a> ## IP Pools Static IP Pools and DHCP Pools support the following attributes:`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `(Required) The first address in the IP Range.`,
				},
				resource.Attribute{
					Name:        "end_address",
					Description: `(Required) The final address in the IP Range. DHCP Pools additionally support the following attributes:`,
				},
				resource.Attribute{
					Name:        "default_lease_time",
					Description: `(Optional) The default DHCP lease time to use. Defaults to ` + "`" + `3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_lease_time",
					Description: `(Optional) The maximum DHCP lease time to use. Defaults to ` + "`" + `7200` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Allows to enable or disable service. Default is true.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_vm",
			Category:         "Resources",
			ShortDescription: `Provides a vCloud Director VM resource. This can be used to create, modify, and delete VMs within a vApp.`,
			Description:      ``,
			Keywords: []string{
				"vapp",
				"vm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vapp_name",
					Description: `(Required) The vApp this VM should belong to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the VM`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `(Required) The catalog name in which to find the given vApp Template`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) The name of the vApp Template to use`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) The amount of RAM (in MB) to allocate to the VM`,
				},
				resource.Attribute{
					Name:        "cpus",
					Description: `(Optional) The number of virtual CPUs to allocate to the VM. Socket count is a result of: virtual logical processors/cores per socket`,
				},
				resource.Attribute{
					Name:        "cpu_cores",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vapp_network_name",
					Description: `(Optional; v2.1+;`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "power_on",
					Description: `(Optional) A boolean value stating if this vApp should be powered on. Default is ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "accept_all_eulas",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "expose_hardware_virtualization",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Independent disk name`,
				},
				resource.Attribute{
					Name:        "bus_number",
					Description: `(Required) Bus number on which to place the disk controller`,
				},
				resource.Attribute{
					Name:        "unit_number",
					Description: `(Required) Unit number (slot) on the bus specified by BusNumber. <a id="network"></a> ## Network`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Computed) Mac address of network interface.`,
				},
				resource.Attribute{
					Name:        "POOL",
					Description: `Static IP address is allocated automatically from defined static pool in network.`,
				},
				resource.Attribute{
					Name:        "DHCP",
					Description: `IP address is obtained from a DHCP service. Field ` + "`" + `ip` + "`" + ` is not guaranteed to be populated. Because of this it may appear after multiple ` + "`" + `terraform refresh` + "`" + ` operations.`,
				},
				resource.Attribute{
					Name:        "MANUAL",
					Description: `IP address is assigned manually in the ` + "`" + `ip` + "`" + ` field. Must be valid IP address from static pool.`,
				},
				resource.Attribute{
					Name:        "NONE",
					Description: `No IP address will be set because VM will have a NIC without network.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"vcd_catalog":          0,
		"vcd_catalog_item":     1,
		"vcd_catalog_media":    2,
		"vcd_dnat":             3,
		"vcd_edgegateway_vpn":  4,
		"vcd_external_network": 5,
		"vcd_firewall_rules":   6,
		"vcd_independent_disk": 7,
		"vcd_inserted_media":   8,
		"vcd_network":          9,
		"vcd_network_direct":   10,
		"vcd_network_isolated": 11,
		"vcd_network_routed":   12,
		"vcd_org":              13,
		"vcd_org_vdc":          14,
		"vcd_snat":             15,
		"vcd_vapp":             16,
		"vcd_vapp_network":     17,
		"vcd_vapp_vm":          18,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
