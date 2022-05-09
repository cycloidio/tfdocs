package dcnm

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dcnm_interface",
			Category:         "Resources",
			ShortDescription: `Manages DCNM interface modules`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrf",
					Description: `(Optional) vrf name for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) ipv4 address for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) ipv6 address for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_tag",
					Description: `(Optional) tag for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_routing_tag",
					Description: `(Optional) routing tag for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_ls_routing",
					Description: `(Optional) link state routing protocol for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_router_id",
					Description: `(Optional) router id for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_replication_mode",
					Description: `(Optional) replication mode for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) configuration for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description for the loopback interface. ## Argument Reference for port-channel Interface ##`,
				},
				resource.Attribute{
					Name:        "pc_interface",
					Description: `(Optional) list of port channel member interface for port-channel interface.`,
				},
				resource.Attribute{
					Name:        "access_vlans",
					Description: `(Optional) access vlans for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) mode for the port-channel interface. Allowed values are "on", "active" and "passive".`,
				},
				resource.Attribute{
					Name:        "bpdu_guard_flag",
					Description: `(Optional) BPDU flag for the port-channel interface. Allowed values are "true", "false" and "no".`,
				},
				resource.Attribute{
					Name:        "port_fast_flag",
					Description: `(Optional) port type fast flag for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) mtu for the port-channel interface. Allowed values are "jumbo" and "default".`,
				},
				resource.Attribute{
					Name:        "allowed_vlans",
					Description: `(Optional) allowed vlans for the port-channel interface. Allowed values are "none", "all" or vlan ranges(1-200,500-2000,3000)`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) configuration for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description for the port-channel interface. ## Argument Reference for vPC Interface ##`,
				},
				resource.Attribute{
					Name:        "switch_name_2",
					Description: `(Optional) name of the second switch with which interface should be associated.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_id",
					Description: `(Optional) peer1 port-channel id for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_id",
					Description: `(Optional) peer2 port-channel id for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_interface",
					Description: `(Optional) list of peer1 member interface for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_interface",
					Description: `(Optional) list of peer2 member interface for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) mode for the vPC interface. Allowed values are "on", "active" and "passive".`,
				},
				resource.Attribute{
					Name:        "bpdu_guard_flag",
					Description: `(Optional) BPDU flag for the vPC interface. Allowed values are "true", "false" and "no".`,
				},
				resource.Attribute{
					Name:        "port_fast_flag",
					Description: `(Optional) port type fast flag for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) mtu for the vPC interface. Allowed values are "jumbo" and "default".`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_allowed_vlans",
					Description: `(Optional) peer1 allowed vlans for the vPC interface. Allowed values are "none", "all" or vlan ranges(1-200,500-2000,3000)`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_allowed_vlans",
					Description: `(Optional) peer2 allowed vlans for the vPC interface. Allowed values are "none", "all" or vlan ranges(1-200,500-2000,3000)`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_access_vlans",
					Description: `(Optional) peer1 access vlans for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_access_vlans",
					Description: `(Optional) peer2 access vlans for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_desc",
					Description: `(Optional) peer1 description for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_desc",
					Description: `(Optional) peer2 description for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_conf",
					Description: `(Optional) peer1 configuration for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_conf",
					Description: `(Optional) peer2 configuration for the vPC interface. ## Argument Reference for sub-interface Interface ##`,
				},
				resource.Attribute{
					Name:        "subinterface_vlan",
					Description: `(Optional) vlan for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "vrf",
					Description: `(Optional) vrf for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) ipv4 address for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) ipv6 address for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "ipv6_prefix",
					Description: `(Optional) ipv6 prefic for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "ipv4_prefix",
					Description: `(Optional) ipv4 prefix for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "subinterface_mtu",
					Description: `(Optional) mtu for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) configuration for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description for the sub-interface. ## Argument Reference for ethernet Interface ##`,
				},
				resource.Attribute{
					Name:        "vrf",
					Description: `(Optional) vrf name for the ethernet interface.`,
				},
				resource.Attribute{
					Name:        "bpdu_guard_flag",
					Description: `(Optional) BPDU flag for the ethernet interface. Allowed values are "true", "false" and "no".`,
				},
				resource.Attribute{
					Name:        "port_fast_flag",
					Description: `(Optional) port type fast flag for the ethernet interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) mtu for the ethernet interface. Allowed values are "jumbo" and "default". If ` + "`" + `policy` + "`" + ` is configured as "epl_routed_intf" or "int_routed_host_11_1", then allowed value range is from 576 to 9216.`,
				},
				resource.Attribute{
					Name:        "ethernet_speed",
					Description: `(Optional) speed of the ethernet. Allowed values are "Auto", "100Mb", "1Gb", "10Gb", "25Gb", "40Gb" and "100Gb".`,
				},
				resource.Attribute{
					Name:        "allowed_vlans",
					Description: `(Optional) allowed vlans for the ethernet interface. Allowed values are "none", "all" or vlan ranges(1-200,500-2000,3000)`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) configuration for the ethernet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description for the ethernet.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) ipv4 address for the ethernet.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) ipv6 address for the ethernet.`,
				},
				resource.Attribute{
					Name:        "ipv6_prefix",
					Description: `(Optional) ipv6 prefix for the ethernet.`,
				},
				resource.Attribute{
					Name:        "ipv4_prefix",
					Description: `(Optional) ipv4 prefix for the ethernet.`,
				},
				resource.Attribute{
					Name:        "access_vlans",
					Description: `(Optional) access vlans for the ethernet interface. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `Dn for the interface module. ## Importing ## An existing interface can be [imported][docs-import] into this resource via its serial number, type and name, using the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import dcnm_interface.example <type>:<serial_number>:<name>:<fabric_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "serial_number",
					Description: `Dn for the interface module. ## Importing ## An existing interface can be [imported][docs-import] into this resource via its serial number, type and name, using the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import dcnm_interface.example <type>:<serial_number>:<name>:<fabric_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dcnm_inventory",
			Category:         "Resources",
			ShortDescription: `Manages DCNM inventory modules`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fabric_name",
					Description: `(Required) Fabric name under which inventory should be created.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username for the the switch.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password for the the switch.`,
				},
				resource.Attribute{
					Name:        "max_hops",
					Description: `(Optional) Maximum number hops for switch. Ranging from 0 to 10, default value is 0.`,
				},
				resource.Attribute{
					Name:        "auth_protocol",
					Description: `(Optional) Authentication protocol for switch. Mapping is as ` + "`" + `0 : "MD5", 1: "SHA", 2 : "MD5_DES", 3 : "MD5_AES", 4 : "SHA_DES", 5 : "SHA_AES"` + "`" + `. Default value is 0.`,
				},
				resource.Attribute{
					Name:        "preserve_config",
					Description: `(Optional) Flag to preserve the configuration of switch. Default value is "false".`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `(Optional) Platform name for the switch. Default value is null.`,
				},
				resource.Attribute{
					Name:        "second_timeout",
					Description: `(Optional) Second timeout value for switch. Default value is null.`,
				},
				resource.Attribute{
					Name:        "config_timeout",
					Description: `(Optional) Configuration timeout value in minutes. Default value is "5".`,
				},
				resource.Attribute{
					Name:        "switch_config",
					Description: `(Required) Switch configuration block for inventory resource. It consists of the information regarding switches.`,
				},
				resource.Attribute{
					Name:        "switch_config.ip",
					Description: `(Required) Ip Address of switch.`,
				},
				resource.Attribute{
					Name:        "switch_config.role",
					Description: `(Optional) Role of the switch. Allowed values are "leaf", "spine", "border", "border_spine", "border_gateway", "border_gateway_spine", "super_spine", "border_super_spine", "border_gateway_super_spine". Default value is "leaf"`,
				},
				resource.Attribute{
					Name:        "deploy",
					Description: `(Optional) Deploy flag for the switch. Default value is "true". ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Dn for the switch inventory.`,
				},
				resource.Attribute{
					Name:        "switch_config",
					Description: `Switch configuration block for inventory.`,
				},
				resource.Attribute{
					Name:        "switch_config.switch_name",
					Description: `Name of the switch.`,
				},
				resource.Attribute{
					Name:        "switch_config.switch_db_id",
					Description: `DB ID for the switch.`,
				},
				resource.Attribute{
					Name:        "switch_config.serial_number",
					Description: `Serial number of the switch.`,
				},
				resource.Attribute{
					Name:        "switch_config.model",
					Description: `Model name of the switch.`,
				},
				resource.Attribute{
					Name:        "switch_config.mode",
					Description: `Mode of the switch. ## Importing ## An existing switch inventory can be [imported][docs-import] into this resource via its fabric and name, using the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import dcnm_inventory.example <fabric_name>:<switch_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Dn for the switch inventory.`,
				},
				resource.Attribute{
					Name:        "switch_config",
					Description: `Switch configuration block for inventory.`,
				},
				resource.Attribute{
					Name:        "switch_config.switch_name",
					Description: `Name of the switch.`,
				},
				resource.Attribute{
					Name:        "switch_config.switch_db_id",
					Description: `DB ID for the switch.`,
				},
				resource.Attribute{
					Name:        "switch_config.serial_number",
					Description: `Serial number of the switch.`,
				},
				resource.Attribute{
					Name:        "switch_config.model",
					Description: `Model name of the switch.`,
				},
				resource.Attribute{
					Name:        "switch_config.mode",
					Description: `Mode of the switch. ## Importing ## An existing switch inventory can be [imported][docs-import] into this resource via its fabric and name, using the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import dcnm_inventory.example <fabric_name>:<switch_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dcnm_network",
			Category:         "Resources",
			ShortDescription: `Manages DCNM Network`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of network object.`,
				},
				resource.Attribute{
					Name:        "fabric_name",
					Description: `(Required) fabric name under which network should be created.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) Network id to be associated with the network to be created. Pass this value while creating multiple networks in single plan to avoid conflict of autogenerating ids. Id will be computed by DCNM if not provided. <strong>Note: </strong> For auto-generation of network-id while creating multiple networks in the same plan, use the depends on functionality of terraform to avoid any network-id conflicts.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) display name for the network object. If not mentioned, then ` + "`" + `name` + "`" + ` will be considered as ` + "`" + `display_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description for the network.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Optional) name of the vrf which should be associated with the network. If not given then will be configured as "NA" with ` + "`" + `l2_only_flag` + "`" + ` as "true".`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(Optional) vlan number for the network.`,
				},
				resource.Attribute{
					Name:        "vlan_name",
					Description: `(Optional) vlan name for the network.`,
				},
				resource.Attribute{
					Name:        "ipv4_gateway",
					Description: `(Optional) ipv4 address of gateway for the network.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway",
					Description: `(Optional) ipv6 address of gateway for the network.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) mtu value for the network. Ranging from 68 to 9216.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) tag for the Network. Ranging from 0 to 4294967295.`,
				},
				resource.Attribute{
					Name:        "secondary_gw_1",
					Description: `(Optional) ipv4 secondary gateway 1 for the network.`,
				},
				resource.Attribute{
					Name:        "secondary_gw_2",
					Description: `(Optional) ipv4 secondary gateway 2 for the network.`,
				},
				resource.Attribute{
					Name:        "secondary_gw_3",
					Description: `(Optional) ipv4 secondary gateway 3 for the network.`,
				},
				resource.Attribute{
					Name:        "secondary_gw_4",
					Description: `(Optional) ipv4 secondary gateway 4 for the network.`,
				},
				resource.Attribute{
					Name:        "arp_supp_flag",
					Description: `(Optional) arp suppression flag for the network.`,
				},
				resource.Attribute{
					Name:        "ir_enable_flag",
					Description: `(Optional) ingress replication flag for the network.`,
				},
				resource.Attribute{
					Name:        "mcast_group",
					Description: `(Optional) multicast group address for the network.`,
				},
				resource.Attribute{
					Name:        "dhcp_1",
					Description: `(Optional) ipv4 address of DHCP server 1 for the network.`,
				},
				resource.Attribute{
					Name:        "dhcp_2",
					Description: `(Optional) ipv4 address of DHCP server 2 for the network.`,
				},
				resource.Attribute{
					Name:        "dhcp_3",
					Description: `(Optional) ipv4 address of DHCP server 3 for the network.`,
				},
				resource.Attribute{
					Name:        "dhcp_vrf",
					Description: `(Optional) vrf name of DHCP server for the network.`,
				},
				resource.Attribute{
					Name:        "dhcp_vrf_2",
					Description: `(Optional) vrf name of DHCP server 2 for the network.`,
				},
				resource.Attribute{
					Name:        "dhcp_vrf_3",
					Description: `(Optional) vrf name of DHCP server 3 for the network.`,
				},
				resource.Attribute{
					Name:        "loopback_id",
					Description: `(Optional) loopback id for the network. Ranging from 0 to 1023.`,
				},
				resource.Attribute{
					Name:        "rt_both_flag",
					Description: `(Optional) l2 VNI route-target both enable flag for the network.`,
				},
				resource.Attribute{
					Name:        "trm_enable_flag",
					Description: `(Optional) TRM enable flag for the network.`,
				},
				resource.Attribute{
					Name:        "l3_gateway_flag",
					Description: `(Optional) enable L3 gateway on border flag for the network.`,
				},
				resource.Attribute{
					Name:        "netflow_flag",
					Description: `(Optional) enable Netflow flag for the network. default is "false".`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) template name for the network. Values allowed "Default_Network_Universal" and "Service_Network_Universal". Default is "Default_Network_Universal".`,
				},
				resource.Attribute{
					Name:        "extension_template",
					Description: `(Optional) extension Template name for the network. Values allowed are "Default_Network_Extension_Universal". Default is "Default_Network_Extension_Universal".`,
				},
				resource.Attribute{
					Name:        "service_template",
					Description: `(Optional) service template name for the network.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) source for the network.`,
				},
				resource.Attribute{
					Name:        "svi_netflow_monitor",
					Description: `(Optional) SVI netflow monitor for the network.`,
				},
				resource.Attribute{
					Name:        "vlan_netflow_monitor",
					Description: `(Optional) VLAN netflow monitor for the network.`,
				},
				resource.Attribute{
					Name:        "nve_id",
					Description: `(Optional) NVE-Id of the network. Default value is 1.`,
				},
				resource.Attribute{
					Name:        "deploy",
					Description: `(Optional) deploy flag, used to deploy the network. Default value is "true".`,
				},
				resource.Attribute{
					Name:        "deploy_timeout",
					Description: `(Optional) deployment timeout, used as the limiter for the deployment status check for network resource. It is in the unit of seconds and default value is "300".`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `(Optional) attachment block, have information regarding the switches which should be attached or detached to/from network. If ` + "`" + `deploy` + "`" + ` is "true", then atleast one attachment must be configured.`,
				},
				resource.Attribute{
					Name:        "attachments.serial_number",
					Description: `(Required) serial number of the switch.`,
				},
				resource.Attribute{
					Name:        "attachments.vlan_id",
					Description: `(Optional) vlan ID for the switch associated with network. If not mentioned then network's default vlan id will be used for attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.attach",
					Description: `(Optional) attach flag for switch. Default value is "true".`,
				},
				resource.Attribute{
					Name:        "attachments.dot1_qvlan",
					Description: `(Optional) dot1 qvlan for switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.switch_ports",
					Description: `(Optional) list of port name(i.e. interface names) for switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.untagged",
					Description: `(Optional) untagged flag for switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.free_form_config",
					Description: `(Optional) free form configuration for the switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.extension_values",
					Description: `(Optional) extension values for switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.instance_values",
					Description: `(Optional) instance values for switch attachment. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Dn for the network.`,
				},
				resource.Attribute{
					Name:        "l2_only_flag",
					Description: `Layer 2 only flag. If VRF is not set then ` + "`" + `l2_only_flag` + "`" + ` will be set to true. ## Importing ## An existing network can be [imported][docs-import] into this resource via its fabric and name, using the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import dcnm_network.example <fabric_name>:<network_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Dn for the network.`,
				},
				resource.Attribute{
					Name:        "l2_only_flag",
					Description: `Layer 2 only flag. If VRF is not set then ` + "`" + `l2_only_flag` + "`" + ` will be set to true. ## Importing ## An existing network can be [imported][docs-import] into this resource via its fabric and name, using the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import dcnm_network.example <fabric_name>:<network_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dcnm_policy",
			Category:         "Resources",
			ShortDescription: `Manages DCNM policy modules`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dcnm_rest",
			Category:         "Resources",
			ShortDescription: `Manages DCNM rest modules`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) DCNM REST endpoint, where the data is being sent.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) HTTP method. Allowed values are "GET", "PUT", "POST", "DELETE".`,
				},
				resource.Attribute{
					Name:        "payload",
					Description: `(Optional) JSON/TEXT payload data.`,
				},
				resource.Attribute{
					Name:        "payload_type",
					Description: `(Optional) Encoding type for payload. Allowed values are "json" and "text". Default value is "json". NOTE: This resource will not work well in the case of Terraform destroy if there is a change in the terraform configuration required to destroy the object from the DCNM, as Destroy only has the access to the data in the state file. To destroy the objects created via dcnm_rest in such cases modify the payload and method and use the Terraform apply instead. ## Attribute Reference No attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dcnm_route_peering",
			Category:         "Resources",
			ShortDescription: `Manages DCNM Route Peering modules`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of route peering.`,
				},
				resource.Attribute{
					Name:        "attached_fabric",
					Description: `(Required) Name of the target fabric for route peering operations.`,
				},
				resource.Attribute{
					Name:        "deployment_mode",
					Description: `(Required) Type of service node.Allowed values are "IntraTenantFW","InterTenantFW","OneArmADC","TwoArmADC","OneArmVNF".`,
				},
				resource.Attribute{
					Name:        "service_fabric",
					Description: `(Required) Name of the target fabric for route peering operations.`,
				},
				resource.Attribute{
					Name:        "next_hop_ip",
					Description: `(Optional) Nexthop IPV4 information.NOTE: This object is applicable only when 'deploy_mode' is 'IntraTenantFW'`,
				},
				resource.Attribute{
					Name:        "option",
					Description: `(Required) Specifies the type of peering.Allowed values are "StaticPeering","EBGPDynamicPeering","None".`,
				},
				resource.Attribute{
					Name:        "service_networks",
					Description: `(Required) List of network under which peering will be created.`,
				},
				resource.Attribute{
					Name:        "service_networks.network_name",
					Description: `(Required) Network name.`,
				},
				resource.Attribute{
					Name:        "service_networks.network_type",
					Description: `(Required) Type of network.Allowed values are "InsideNetworkFW"(service node = Firewall),"OutsideNetworkFW"(service node = Firewall),"ArmOneADC"(service node = ADC),"ArmTwoADC"(service node = ADC),"ArmOneVNF"(service node= VNF).`,
				},
				resource.Attribute{
					Name:        "service_networks.template_name",
					Description: `(Required) Name of template.`,
				},
				resource.Attribute{
					Name:        "service_networks.vrf_name",
					Description: `(Required) VRF name under which network is created.`,
				},
				resource.Attribute{
					Name:        "service_networks.vlan_id",
					Description: `(Required) VLan Id of network.`,
				},
				resource.Attribute{
					Name:        "service_networks.gateway_ip_address",
					Description: `(Required) IPV4 gateway information including the mask e.g. 192.168.1.1/24.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `(Optional) Routing configuration.`,
				},
				resource.Attribute{
					Name:        "routes.template_name",
					Description: `(Optional) Template name for routing.`,
				},
				resource.Attribute{
					Name:        "routes.route_parmas",
					Description: `(Optional) NVPair map for routing. The value for predefined route parameters depens upon deployment mode.`,
				},
				resource.Attribute{
					Name:        "routes.vrf_name",
					Description: `(Optional) VRF name for routing.`,
				},
				resource.Attribute{
					Name:        "deploy",
					Description: `(Optional) A flag specifying if a route peering is to be deployed on the switches. Default value is "true".`,
				},
				resource.Attribute{
					Name:        "deploy_timeout",
					Description: `(Optional) Timeout seconds for deployment. Default value is 300s.`,
				},
				resource.Attribute{
					Name:        "service_node_type",
					Description: `(Required) Type of service node.Allowed values are "Firewall","VNF","ADC". ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Route peering deployment status. ## Importing ## An existing route peering can be [imported][docs-import] into this resource via its fabric and name, using the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import dcnm_route_peering.example <peering_name>:<external_fabric>:<service_node>:<attached_fabric_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Route peering deployment status. ## Importing ## An existing route peering can be [imported][docs-import] into this resource via its fabric and name, using the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import dcnm_route_peering.example <peering_name>:<external_fabric>:<service_node>:<attached_fabric_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dcnm_service_node",
			Category:         "Resources",
			ShortDescription: `Manages DCNM Service Node`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dcnm_service_policy",
			Category:         "Resources",
			ShortDescription: `Manages DCNM Service Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dcnm_template",
			Category:         "Resources",
			ShortDescription: `Manages DCNM Template`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Template.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) Content of file or file name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of template.`,
				},
				resource.Attribute{
					Name:        "supported_platforms",
					Description: `(Optional) Platform supported by the template.`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `(Optional) Type of template.`,
				},
				resource.Attribute{
					Name:        "template_content_type",
					Description: `(Optional) Content type of template.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tag of template.`,
				},
				resource.Attribute{
					Name:        "template_sub_type",
					Description: `(Optional) Sub type of template. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the template. ## Importing ## An existing Template can be [imported][docs-import] into this resource via template name, using the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import dcnm_template.example <template_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dcnm_vrf",
			Category:         "Resources",
			ShortDescription: `Manages DCNM VRF`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"dcnm_interface":      0,
		"dcnm_inventory":      1,
		"dcnm_network":        2,
		"dcnm_policy":         3,
		"dcnm_rest":           4,
		"dcnm_route_peering":  5,
		"dcnm_service_node":   6,
		"dcnm_service_policy": 7,
		"dcnm_template":       8,
		"dcnm_vrf":            9,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
