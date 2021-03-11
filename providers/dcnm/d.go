package dcnm

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dcnm_interface",
			Category:         "Data Sources",
			ShortDescription: `Data source for DCNM interface module`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "serial_number",
					Description: `(Required) Dn for the interface module.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the interface.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) type of the interface. Allowed values are "loopback", "port-channel", "vpc", "sub-interface", "ethernet". ## Common Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "fabric_name",
					Description: `fabric name under which interface is created.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `policy name for the interface.`,
				},
				resource.Attribute{
					Name:        "admin_state",
					Description: `administrative state for the interface.`,
				},
				resource.Attribute{
					Name:        "deploy",
					Description: `deploy flag for the deployment of interface.`,
				},
				resource.Attribute{
					Name:        "switch_name_1",
					Description: `name of the switch which is associated to the interface. ## Attribute Reference for loopback Interface ##`,
				},
				resource.Attribute{
					Name:        "vrf",
					Description: `vrf name for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `ipv4 address for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `ipv6 address for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_tag",
					Description: `tag for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_routing_tag",
					Description: `routing tag for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_ls_routing",
					Description: `link state routing protocol for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_router_id",
					Description: `router id for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_replication_mode",
					Description: `replication mode for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `configuration for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description for the loopback interface. ## Attribute Reference for port-channel Interface ##`,
				},
				resource.Attribute{
					Name:        "pc_interface",
					Description: `list of port channel member interface for port-channel interface.`,
				},
				resource.Attribute{
					Name:        "access_vlans",
					Description: `access vlans for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `mode for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "bpdu_gaurd_flag",
					Description: `BPDU flag for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "port_fast_flag",
					Description: `port type fast flag for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `mtu for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "allowed_vlans",
					Description: `allowed vlans for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `configuration for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description for the port-channel interface. ## Attribute Reference for vPC Interface ##`,
				},
				resource.Attribute{
					Name:        "switch_name_2",
					Description: `name of the second switch with which vpc is associated.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_id",
					Description: `peer1 port-channel id for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_id",
					Description: `peer2 port-channel id for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_interface",
					Description: `list of peer1 member interface for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_interface",
					Description: `list of peer2 member interface for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `mode for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "bpdu_gaurd_flag",
					Description: `BPDU flag for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "port_fast_flag",
					Description: `port type fast flag for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `mtu for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_allowed_vlans",
					Description: `peer1 allowed vlans for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_allowed_vlans",
					Description: `peer2 allowed vlans for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_access_vlans",
					Description: `peer1 access vlans for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_access_vlans",
					Description: `peer2 access vlans for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_desc",
					Description: `peer1 description for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_desc",
					Description: `peer2 description for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_conf",
					Description: `peer1 configuration for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_conf",
					Description: `peer2 configuration for the vPC interface. ## Attribute Reference for sub-interface Interface ##`,
				},
				resource.Attribute{
					Name:        "subinterface_vlan",
					Description: `vlan for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "vrf",
					Description: `vrf for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `ipv4 address for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `ipv6 address for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "ipv6_prefix",
					Description: `ipv6 prefic for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "ipv4_prefix",
					Description: `ipv4 prefix for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "subinterface_mtu",
					Description: `mtu for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `configuration for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description for the sub-interface. ## Attribute Reference for ethernet Interface ##`,
				},
				resource.Attribute{
					Name:        "bpdu_gaurd_flag",
					Description: `BPDU flag for the ethernet interface.`,
				},
				resource.Attribute{
					Name:        "port_fast_flag",
					Description: `port type fast flag for the ethernet interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `mtu for the ethernet interface.`,
				},
				resource.Attribute{
					Name:        "ethernet_speed",
					Description: `speed of the ethernet.`,
				},
				resource.Attribute{
					Name:        "allowed_vlans",
					Description: `allowed vlans for the ethernet interface.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `configuration for the ethernet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description for the ethernet.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `ipv4 address for the ethernet.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `ipv6 address for the ethernet.`,
				},
				resource.Attribute{
					Name:        "ipv6_prefix",
					Description: `ipv6 prefic for the ethernet.`,
				},
				resource.Attribute{
					Name:        "ipv4_prefix",
					Description: `ipv4 prefix for the ethernet.`,
				},
				resource.Attribute{
					Name:        "access_vlans",
					Description: `access vlans for the ethernet interface.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vrf",
					Description: `vrf name for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `ipv4 address for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `ipv6 address for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_tag",
					Description: `tag for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_routing_tag",
					Description: `routing tag for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_ls_routing",
					Description: `link state routing protocol for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_router_id",
					Description: `router id for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "loopback_replication_mode",
					Description: `replication mode for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `configuration for the loopback interface.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description for the loopback interface. ## Attribute Reference for port-channel Interface ##`,
				},
				resource.Attribute{
					Name:        "pc_interface",
					Description: `list of port channel member interface for port-channel interface.`,
				},
				resource.Attribute{
					Name:        "access_vlans",
					Description: `access vlans for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `mode for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "bpdu_gaurd_flag",
					Description: `BPDU flag for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "port_fast_flag",
					Description: `port type fast flag for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `mtu for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "allowed_vlans",
					Description: `allowed vlans for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `configuration for the port-channel interface.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description for the port-channel interface. ## Attribute Reference for vPC Interface ##`,
				},
				resource.Attribute{
					Name:        "switch_name_2",
					Description: `name of the second switch with which vpc is associated.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_id",
					Description: `peer1 port-channel id for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_id",
					Description: `peer2 port-channel id for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_interface",
					Description: `list of peer1 member interface for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_interface",
					Description: `list of peer2 member interface for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `mode for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "bpdu_gaurd_flag",
					Description: `BPDU flag for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "port_fast_flag",
					Description: `port type fast flag for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `mtu for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_allowed_vlans",
					Description: `peer1 allowed vlans for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_allowed_vlans",
					Description: `peer2 allowed vlans for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_access_vlans",
					Description: `peer1 access vlans for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_access_vlans",
					Description: `peer2 access vlans for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_desc",
					Description: `peer1 description for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_desc",
					Description: `peer2 description for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer1_conf",
					Description: `peer1 configuration for the vPC interface.`,
				},
				resource.Attribute{
					Name:        "vpc_peer2_conf",
					Description: `peer2 configuration for the vPC interface. ## Attribute Reference for sub-interface Interface ##`,
				},
				resource.Attribute{
					Name:        "subinterface_vlan",
					Description: `vlan for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "vrf",
					Description: `vrf for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `ipv4 address for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `ipv6 address for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "ipv6_prefix",
					Description: `ipv6 prefic for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "ipv4_prefix",
					Description: `ipv4 prefix for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "subinterface_mtu",
					Description: `mtu for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `configuration for the sub-interface.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description for the sub-interface. ## Attribute Reference for ethernet Interface ##`,
				},
				resource.Attribute{
					Name:        "bpdu_gaurd_flag",
					Description: `BPDU flag for the ethernet interface.`,
				},
				resource.Attribute{
					Name:        "port_fast_flag",
					Description: `port type fast flag for the ethernet interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `mtu for the ethernet interface.`,
				},
				resource.Attribute{
					Name:        "ethernet_speed",
					Description: `speed of the ethernet.`,
				},
				resource.Attribute{
					Name:        "allowed_vlans",
					Description: `allowed vlans for the ethernet interface.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `configuration for the ethernet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description for the ethernet.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `ipv4 address for the ethernet.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `ipv6 address for the ethernet.`,
				},
				resource.Attribute{
					Name:        "ipv6_prefix",
					Description: `ipv6 prefic for the ethernet.`,
				},
				resource.Attribute{
					Name:        "ipv4_prefix",
					Description: `ipv4 prefix for the ethernet.`,
				},
				resource.Attribute{
					Name:        "access_vlans",
					Description: `access vlans for the ethernet interface.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dcnm_inventory",
			Category:         "Data Sources",
			ShortDescription: `Data source for DCNM inventory module`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fabric_name",
					Description: `(Required) fabric name under which inventory should be created.`,
				},
				resource.Attribute{
					Name:        "switch_name",
					Description: `(Required) name of switch. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Dn for the switch inventory.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Ip address of the switch.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Role of the switch.`,
				},
				resource.Attribute{
					Name:        "switch_db_id",
					Description: `Db id for the switch.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `Serial number of the switch.`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `Model name of the switch.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Mode of the switch.`,
				},
				resource.Attribute{
					Name:        "deploy",
					Description: `(Optional) deploy flag for the switch.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Dn for the switch inventory.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Ip address of the switch.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Role of the switch.`,
				},
				resource.Attribute{
					Name:        "switch_db_id",
					Description: `Db id for the switch.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `Serial number of the switch.`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `Model name of the switch.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Mode of the switch.`,
				},
				resource.Attribute{
					Name:        "deploy",
					Description: `(Optional) deploy flag for the switch.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dcnm_network",
			Category:         "Data Sources",
			ShortDescription: `Data source for DCNM network`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of network object.`,
				},
				resource.Attribute{
					Name:        "fabric_name",
					Description: `(Required) fabric name under which network exists. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the network.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `display name for the network object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description for the network.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `name of the vrf which should be associated with the network.`,
				},
				resource.Attribute{
					Name:        "l2_only_flag",
					Description: `layer 2 only flag for the network.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `vlan number for the network.`,
				},
				resource.Attribute{
					Name:        "vlan_name",
					Description: `vlan name for the network.`,
				},
				resource.Attribute{
					Name:        "ipv4_gateway",
					Description: `ipv4 address of gateway for the network.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway",
					Description: `ipv6 address of gateway for the network.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `mtu value for the network.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `tag for the Network.`,
				},
				resource.Attribute{
					Name:        "secondary_gw_1",
					Description: `ipv4 secondary gateway 1 for the network.`,
				},
				resource.Attribute{
					Name:        "secondary_gw_2",
					Description: `ipv4 secondary gateway 2 for the network.`,
				},
				resource.Attribute{
					Name:        "arp_supp_flag",
					Description: `arp suppression flag for the network.`,
				},
				resource.Attribute{
					Name:        "ir_enable_flag",
					Description: `ingress replication flag for the network.`,
				},
				resource.Attribute{
					Name:        "mcast_group",
					Description: `multicast group address for the network.`,
				},
				resource.Attribute{
					Name:        "dhcp_1",
					Description: `ipv4 address of DHCP server 1 for the network.`,
				},
				resource.Attribute{
					Name:        "dhcp_2",
					Description: `ipv4 address of DHCP server 2 for the network.`,
				},
				resource.Attribute{
					Name:        "dhcp_vrf",
					Description: `vrf name of DHCP server for the network.`,
				},
				resource.Attribute{
					Name:        "loopback_id",
					Description: `loopback id for the network.`,
				},
				resource.Attribute{
					Name:        "rt_both_flag",
					Description: `l2 VNI route-target both enable flag for the network.`,
				},
				resource.Attribute{
					Name:        "trm_enable_flag",
					Description: `TRM enable flag for the network.`,
				},
				resource.Attribute{
					Name:        "l3_gateway_flag",
					Description: `enable L3 gateway on border flag for the network.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `template name for the network. Default is "Default_VRF_Universal".`,
				},
				resource.Attribute{
					Name:        "extension_template",
					Description: `extension Template name for the network. Default is "Default_Network_Extension_Universal".`,
				},
				resource.Attribute{
					Name:        "service_template",
					Description: `service template name for the network.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `source for the network.`,
				},
				resource.Attribute{
					Name:        "deploy",
					Description: `deploy flag, used to deploy the network.`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `attachment block, have information regarding the switches which should be attached or detached to/from network.`,
				},
				resource.Attribute{
					Name:        "attachments.serial_number",
					Description: `serial number of the switch.`,
				},
				resource.Attribute{
					Name:        "attachments.vlan_id",
					Description: `vlan ID for the switch associated with network.`,
				},
				resource.Attribute{
					Name:        "attachments.attach",
					Description: `attach flag for switch.`,
				},
				resource.Attribute{
					Name:        "attachments.switch_ports",
					Description: `list of port name(i.e. interface names) for switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.untagged",
					Description: `untagged flag for switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.free_form_config",
					Description: `free form configuration for the switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.extension_values",
					Description: `extension values for switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.instance_values",
					Description: `instance values for switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.dot1_qvlan",
					Description: `dot1 qvlan for switch attachment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the network.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `display name for the network object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description for the network.`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `name of the vrf which should be associated with the network.`,
				},
				resource.Attribute{
					Name:        "l2_only_flag",
					Description: `layer 2 only flag for the network.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `vlan number for the network.`,
				},
				resource.Attribute{
					Name:        "vlan_name",
					Description: `vlan name for the network.`,
				},
				resource.Attribute{
					Name:        "ipv4_gateway",
					Description: `ipv4 address of gateway for the network.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway",
					Description: `ipv6 address of gateway for the network.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `mtu value for the network.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `tag for the Network.`,
				},
				resource.Attribute{
					Name:        "secondary_gw_1",
					Description: `ipv4 secondary gateway 1 for the network.`,
				},
				resource.Attribute{
					Name:        "secondary_gw_2",
					Description: `ipv4 secondary gateway 2 for the network.`,
				},
				resource.Attribute{
					Name:        "arp_supp_flag",
					Description: `arp suppression flag for the network.`,
				},
				resource.Attribute{
					Name:        "ir_enable_flag",
					Description: `ingress replication flag for the network.`,
				},
				resource.Attribute{
					Name:        "mcast_group",
					Description: `multicast group address for the network.`,
				},
				resource.Attribute{
					Name:        "dhcp_1",
					Description: `ipv4 address of DHCP server 1 for the network.`,
				},
				resource.Attribute{
					Name:        "dhcp_2",
					Description: `ipv4 address of DHCP server 2 for the network.`,
				},
				resource.Attribute{
					Name:        "dhcp_vrf",
					Description: `vrf name of DHCP server for the network.`,
				},
				resource.Attribute{
					Name:        "loopback_id",
					Description: `loopback id for the network.`,
				},
				resource.Attribute{
					Name:        "rt_both_flag",
					Description: `l2 VNI route-target both enable flag for the network.`,
				},
				resource.Attribute{
					Name:        "trm_enable_flag",
					Description: `TRM enable flag for the network.`,
				},
				resource.Attribute{
					Name:        "l3_gateway_flag",
					Description: `enable L3 gateway on border flag for the network.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `template name for the network. Default is "Default_VRF_Universal".`,
				},
				resource.Attribute{
					Name:        "extension_template",
					Description: `extension Template name for the network. Default is "Default_Network_Extension_Universal".`,
				},
				resource.Attribute{
					Name:        "service_template",
					Description: `service template name for the network.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `source for the network.`,
				},
				resource.Attribute{
					Name:        "deploy",
					Description: `deploy flag, used to deploy the network.`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `attachment block, have information regarding the switches which should be attached or detached to/from network.`,
				},
				resource.Attribute{
					Name:        "attachments.serial_number",
					Description: `serial number of the switch.`,
				},
				resource.Attribute{
					Name:        "attachments.vlan_id",
					Description: `vlan ID for the switch associated with network.`,
				},
				resource.Attribute{
					Name:        "attachments.attach",
					Description: `attach flag for switch.`,
				},
				resource.Attribute{
					Name:        "attachments.switch_ports",
					Description: `list of port name(i.e. interface names) for switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.untagged",
					Description: `untagged flag for switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.free_form_config",
					Description: `free form configuration for the switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.extension_values",
					Description: `extension values for switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.instance_values",
					Description: `instance values for switch attachment.`,
				},
				resource.Attribute{
					Name:        "attachments.dot1_qvlan",
					Description: `dot1 qvlan for switch attachment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dcnm_vrf",
			Category:         "Data Sources",
			ShortDescription: `Data source for DCNM VRF`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object VRF.`,
				},
				resource.Attribute{
					Name:        "fabric_name",
					Description: `(Required) fabric name under which VRF exists. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `attribute id set to the Dn of the VRF.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `vlan Id for the VRF.`,
				},
				resource.Attribute{
					Name:        "vlan_name",
					Description: `vlan name for the VRF.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description for the VRF.`,
				},
				resource.Attribute{
					Name:        "intf_description",
					Description: `intf desscription for the VRF.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `tag for the VRF.`,
				},
				resource.Attribute{
					Name:        "max_bgp_path",
					Description: `maximum BGP path value for the VRF.`,
				},
				resource.Attribute{
					Name:        "max_ibgp_path",
					Description: `maximum iBGP path value for the VRF.`,
				},
				resource.Attribute{
					Name:        "trm_enable",
					Description: `trm enable flag for the VRF. Allowed values are "true" and "false".`,
				},
				resource.Attribute{
					Name:        "rp_external_flag",
					Description: `rp external flag for the VRF. Allowed values are "true" and "false".`,
				},
				resource.Attribute{
					Name:        "rp_address",
					Description: `rp address for the VRF.`,
				},
				resource.Attribute{
					Name:        "loopback_id",
					Description: `loopback ip address for the VRF.`,
				},
				resource.Attribute{
					Name:        "mutlicast_group",
					Description: `multicast group address for the VRF.`,
				},
				resource.Attribute{
					Name:        "mutlicast_address",
					Description: `multicast address for the VRF.`,
				},
				resource.Attribute{
					Name:        "ipv6_link_local_flag",
					Description: `ipv6 link local enable flag for the VRF. Allowed values are "true" and "false".`,
				},
				resource.Attribute{
					Name:        "trm_bgw_msite_flag",
					Description: `trm bgw multisite enable flag for the VRF. Allowed values are "true" and "false".`,
				},
				resource.Attribute{
					Name:        "advertise_host_route",
					Description: `advertise host route enable flag for the VRF. Allowed values are "true" and "false".`,
				},
				resource.Attribute{
					Name:        "advertise_default_route",
					Description: `advertise default route enable flag for the VRF. Allowed values are "true" and "false".`,
				},
				resource.Attribute{
					Name:        "static_default_route",
					Description: `configure static default route enable flag for the VRF. Allowed values are "true" and "false".`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `template name for the VRF. Values allowed "Default_VRF_Universal". Default is "Default_VRF_Universal".`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `mtu value for the VRF. Ranginf from 68 to 9216.`,
				},
				resource.Attribute{
					Name:        "extension_template",
					Description: `extension Template name for the VRF. Values allowed are "Default_VRF_Extension_Universal". Default is "Default_VRF_Extension_Universal".`,
				},
				resource.Attribute{
					Name:        "service_template",
					Description: `service template name for the VRF.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `source for the VRF.`,
				},
				resource.Attribute{
					Name:        "deploy",
					Description: `deploy flag, used to deploy the VRF. Default value is "true".`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `attachment Block, have information regarding the switches which should be attached or detached to/from VRF.`,
				},
				resource.Attribute{
					Name:        "attachments.serial_number",
					Description: `serial number of the switch.`,
				},
				resource.Attribute{
					Name:        "attachments.vlan_id",
					Description: `vlan ID for the switch associated with VRF.`,
				},
				resource.Attribute{
					Name:        "attachments.attach",
					Description: `attach flag for switch. Default value is "true".`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `attribute id set to the Dn of the VRF.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `vlan Id for the VRF.`,
				},
				resource.Attribute{
					Name:        "vlan_name",
					Description: `vlan name for the VRF.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description for the VRF.`,
				},
				resource.Attribute{
					Name:        "intf_description",
					Description: `intf desscription for the VRF.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `tag for the VRF.`,
				},
				resource.Attribute{
					Name:        "max_bgp_path",
					Description: `maximum BGP path value for the VRF.`,
				},
				resource.Attribute{
					Name:        "max_ibgp_path",
					Description: `maximum iBGP path value for the VRF.`,
				},
				resource.Attribute{
					Name:        "trm_enable",
					Description: `trm enable flag for the VRF. Allowed values are "true" and "false".`,
				},
				resource.Attribute{
					Name:        "rp_external_flag",
					Description: `rp external flag for the VRF. Allowed values are "true" and "false".`,
				},
				resource.Attribute{
					Name:        "rp_address",
					Description: `rp address for the VRF.`,
				},
				resource.Attribute{
					Name:        "loopback_id",
					Description: `loopback ip address for the VRF.`,
				},
				resource.Attribute{
					Name:        "mutlicast_group",
					Description: `multicast group address for the VRF.`,
				},
				resource.Attribute{
					Name:        "mutlicast_address",
					Description: `multicast address for the VRF.`,
				},
				resource.Attribute{
					Name:        "ipv6_link_local_flag",
					Description: `ipv6 link local enable flag for the VRF. Allowed values are "true" and "false".`,
				},
				resource.Attribute{
					Name:        "trm_bgw_msite_flag",
					Description: `trm bgw multisite enable flag for the VRF. Allowed values are "true" and "false".`,
				},
				resource.Attribute{
					Name:        "advertise_host_route",
					Description: `advertise host route enable flag for the VRF. Allowed values are "true" and "false".`,
				},
				resource.Attribute{
					Name:        "advertise_default_route",
					Description: `advertise default route enable flag for the VRF. Allowed values are "true" and "false".`,
				},
				resource.Attribute{
					Name:        "static_default_route",
					Description: `configure static default route enable flag for the VRF. Allowed values are "true" and "false".`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `template name for the VRF. Values allowed "Default_VRF_Universal". Default is "Default_VRF_Universal".`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `mtu value for the VRF. Ranginf from 68 to 9216.`,
				},
				resource.Attribute{
					Name:        "extension_template",
					Description: `extension Template name for the VRF. Values allowed are "Default_VRF_Extension_Universal". Default is "Default_VRF_Extension_Universal".`,
				},
				resource.Attribute{
					Name:        "service_template",
					Description: `service template name for the VRF.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `source for the VRF.`,
				},
				resource.Attribute{
					Name:        "deploy",
					Description: `deploy flag, used to deploy the VRF. Default value is "true".`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `attachment Block, have information regarding the switches which should be attached or detached to/from VRF.`,
				},
				resource.Attribute{
					Name:        "attachments.serial_number",
					Description: `serial number of the switch.`,
				},
				resource.Attribute{
					Name:        "attachments.vlan_id",
					Description: `vlan ID for the switch associated with VRF.`,
				},
				resource.Attribute{
					Name:        "attachments.attach",
					Description: `attach flag for switch. Default value is "true".`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"dcnm_interface": 0,
		"dcnm_inventory": 1,
		"dcnm_network":   2,
		"dcnm_vrf":       3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
