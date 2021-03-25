package aci

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "aci_aaa_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI aaa Domain`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object aaa domain. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the aaa domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object aaa domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object aaa domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the aaa domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object aaa domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object aaa domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_generic",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Access Generic`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "attachable_access_entity_profile_dn",
					Description: `(Required) Distinguished name of parent AttachableAccessEntityProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object access_generic. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Access Generic.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object access_generic.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object access_generic.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Access Generic.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object access_generic.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object access_generic.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Access Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_port_selector_dn",
					Description: `(Required) Distinguished name of parent AccessPortSelector object. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Access Access Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object access_access_group.`,
				},
				resource.Attribute{
					Name:        "fex_id",
					Description: `(Optional) interface policy group fex id`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Optional) interface policy group's target rn`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Access Access Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object access_access_group.`,
				},
				resource.Attribute{
					Name:        "fex_id",
					Description: `(Optional) interface policy group fex id`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Optional) interface policy group's target rn`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_port_block",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Access Port Block`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_port_selector_dn",
					Description: `(Required) Distinguished name of parent AccessPortSelector object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object access_port_block. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Access Port Block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object access_port_block.`,
				},
				resource.Attribute{
					Name:        "from_card",
					Description: `(Optional) The beginning (from-range) of the card range block for the leaf access port block.`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `(Optional) The beginning (from-range) of the port range block for the leaf access port block.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object access_port_block.`,
				},
				resource.Attribute{
					Name:        "to_card",
					Description: `(Optional) The end (to-range) of the card range block for the leaf access port block.`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `(Optional) The end (to-range) of the port range block for the leaf access port block.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Access Port Block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object access_port_block.`,
				},
				resource.Attribute{
					Name:        "from_card",
					Description: `(Optional) The beginning (from-range) of the card range block for the leaf access port block.`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `(Optional) The beginning (from-range) of the port range block for the leaf access port block.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object access_port_block.`,
				},
				resource.Attribute{
					Name:        "to_card",
					Description: `(Optional) The end (to-range) of the card range block for the leaf access port block.`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `(Optional) The end (to-range) of the port range block for the leaf access port block.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_port_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Access Port Selector`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "leaf_interface_profile_dn",
					Description: `(Required) Distinguished name of parent LeafInterfaceProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object access_port_selector.`,
				},
				resource.Attribute{
					Name:        "access_port_selector_type",
					Description: `(Required) access_port_selector_type of Object access_port_selector. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Access Port Selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object access_port_selector.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object access_port_selector.`,
				},
				resource.Attribute{
					Name:        "access_port_selector_type",
					Description: `(Optional) host port selector type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Access Port Selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object access_port_selector.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object access_port_selector.`,
				},
				resource.Attribute{
					Name:        "access_port_selector_type",
					Description: `(Optional) host port selector type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_sub_port_block",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Access Sub Port Block`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_port_selector_dn",
					Description: `(Required) Distinguished name of parent AccessPortSelector object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object access_sub_port_block. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Access Sub Port Block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object access_sub_port_block.`,
				},
				resource.Attribute{
					Name:        "from_card",
					Description: `(Optional) from card`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `(Optional) port block from port`,
				},
				resource.Attribute{
					Name:        "from_sub_port",
					Description: `(Optional) from_sub_port for object access_sub_port_block.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object access_sub_port_block.`,
				},
				resource.Attribute{
					Name:        "to_card",
					Description: `(Optional) to card`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `(Optional) to port`,
				},
				resource.Attribute{
					Name:        "to_sub_port",
					Description: `(Optional) to_sub_port for object access_sub_port_block.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Access Sub Port Block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object access_sub_port_block.`,
				},
				resource.Attribute{
					Name:        "from_card",
					Description: `(Optional) from card`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `(Optional) port block from port`,
				},
				resource.Attribute{
					Name:        "from_sub_port",
					Description: `(Optional) from_sub_port for object access_sub_port_block.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object access_sub_port_block.`,
				},
				resource.Attribute{
					Name:        "to_card",
					Description: `(Optional) to card`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `(Optional) to port`,
				},
				resource.Attribute{
					Name:        "to_sub_port",
					Description: `(Optional) to_sub_port for object access_sub_port_block.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_action_rule_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Action Rule Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object action_rule_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Action Rule Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object action_rule_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object action_rule_profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Action Rule Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object action_rule_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object action_rule_profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_any",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Any`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrf_dn",
					Description: `(Required) Distinguished name of parent VRF object. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Any.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object any.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) Represents the provider label match criteria.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object any.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPgs can be divided in a the context can be divided in two subgroups.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Any.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object any.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) Represents the provider label match criteria.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object any.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPgs can be divided in a the context can be divided in two subgroups.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_application_epg",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Application EPG`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_profile_dn",
					Description: `(Required) Distinguished name of parent ApplicationProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object application_epg. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Application EPG.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object application_epg.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) exception_tag for object application_epg.`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings.`,
				},
				resource.Attribute{
					Name:        "fwd_ctrl",
					Description: `(Optional) Forwarding control at EPG level.`,
				},
				resource.Attribute{
					Name:        "has_mcast_source",
					Description: `(Optional) If the source for the EPG is multicast or not.`,
				},
				resource.Attribute{
					Name:        "is_attr_based_epg",
					Description: `(Optional) If the EPG is attribute based or not.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria for EPG.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object application_epg.`,
				},
				resource.Attribute{
					Name:        "pc_enf_pref",
					Description: `(Optional) The preferred policy control.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) qos priority class id`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `(Optional) shutdown for object application_epg.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_sec_inherited",
					Description: `(Optional) Relation to another epg (fvEPg class) to inherit contracts from. Named 'contract master' in GUI.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Application EPG.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object application_epg.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) exception_tag for object application_epg.`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings.`,
				},
				resource.Attribute{
					Name:        "fwd_ctrl",
					Description: `(Optional) Forwarding control at EPG level.`,
				},
				resource.Attribute{
					Name:        "has_mcast_source",
					Description: `(Optional) If the source for the EPG is multicast or not.`,
				},
				resource.Attribute{
					Name:        "is_attr_based_epg",
					Description: `(Optional) If the EPG is attribute based or not.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria for EPG.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object application_epg.`,
				},
				resource.Attribute{
					Name:        "pc_enf_pref",
					Description: `(Optional) The preferred policy control.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) qos priority class id`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `(Optional) shutdown for object application_epg.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_sec_inherited",
					Description: `(Optional) Relation to another epg (fvEPg class) to inherit contracts from. Named 'contract master' in GUI.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_application_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Application Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object application_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Application Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object application_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object application_profile.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) priority class id`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Application Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object application_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object application_profile.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) priority class id`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_attachable_access_entity_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Attachable Access Entity Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object attachable_access_entity_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Attachable Access Entity Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object attachable_access_entity_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object attachable_access_entity_profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Attachable Access Entity Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object attachable_access_entity_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object attachable_access_entity_profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_autonomous_system_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Autonomous System Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Autonomous System Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object autonomous_system_profile.`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `(Optional) A number that uniquely identifies an autonomous system.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object autonomous_system_profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Autonomous System Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object autonomous_system_profile.`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `(Optional) A number that uniquely identifies an autonomous system.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object autonomous_system_profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bd_dhcp_label",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI BD DHCP Label`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bridge_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Bridge Domain`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object bridge_domain. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Bridge Domain.`,
				},
				resource.Attribute{
					Name:        "optimize_wan_bandwidth",
					Description: `(Optional) Flag to enable OptimizeWanBandwidth between sites.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object bridge_domain.`,
				},
				resource.Attribute{
					Name:        "arp_flood",
					Description: `(Optional) A property to specify whether ARP flooding is enabled. If flooding is disabled, unicast routing will be performed on the target IP address.`,
				},
				resource.Attribute{
					Name:        "ep_clear",
					Description: `(Optional) Represents the parameter used by the node (i.e. Leaf) to clear all EPs in all leaves for this BD.`,
				},
				resource.Attribute{
					Name:        "ep_move_detect_mode",
					Description: `(Optional) The End Point move detection option uses the Gratuitous Address Resolution Protocol (GARP). A gratuitous ARP is an ARP broadcast-type of packet that is used to verify that no other device on the network has the same IP address as the sending device.`,
				},
				resource.Attribute{
					Name:        "host_based_routing",
					Description: `(Optional) Enables advertising host routes out of l3outs of this BD.`,
				},
				resource.Attribute{
					Name:        "intersite_bum_traffic_allow",
					Description: `(Optional) Control whether BUM traffic is allowed between sites.`,
				},
				resource.Attribute{
					Name:        "intersite_l2_stretch",
					Description: `(Optional) Flag to enable l2Stretch between sites.`,
				},
				resource.Attribute{
					Name:        "ip_learning",
					Description: `(Optional) Endpoint Dataplane Learning.`,
				},
				resource.Attribute{
					Name:        "ipv6_mcast_allow",
					Description: `(Optional) Flag to indicate multicast IpV6 is allowed or not.`,
				},
				resource.Attribute{
					Name:        "limit_ip_learn_to_subnets",
					Description: `(Optional) Limits IP address learning to the bridge domain subnets only. Every BD can have multiple subnets associated with it. By default, all IPs are learned.`,
				},
				resource.Attribute{
					Name:        "ll_addr",
					Description: `(Optional) override of system generated ipv6 link-local address.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Optional) The MAC address of the bridge domain (BD) or switched virtual interface (SVI). Every BD by default takes the fabric-wide default MAC address. You can override that address with a different one. By default the BD will take a 00:22:BD:F8:19:FF mac address.`,
				},
				resource.Attribute{
					Name:        "mcast_allow",
					Description: `(Optional) Flag to indicate if multicast is enabled for IpV4 addresses.`,
				},
				resource.Attribute{
					Name:        "multi_dst_pkt_act",
					Description: `(Optional) The multiple destination forwarding method for L2 Multicast, Broadcast, and Link Layer traffic types.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object bridge_domain.`,
				},
				resource.Attribute{
					Name:        "bridge_domain_type",
					Description: `(Optional) The specific type of the object or component.`,
				},
				resource.Attribute{
					Name:        "unicast_route",
					Description: `(Optional) The forwarding method based on predefined forwarding criteria (IP or MAC address).`,
				},
				resource.Attribute{
					Name:        "unk_mac_ucast_act",
					Description: `(Optional) The forwarding method for unknown layer 2 destinations.`,
				},
				resource.Attribute{
					Name:        "unk_mcast_act",
					Description: `(Optional) The parameter used by the node (i.e. a leaf) for forwarding data for an unknown multicast destination.`,
				},
				resource.Attribute{
					Name:        "v6unk_mcast_act",
					Description: `(Optional) v6unk_mcast_act for object bridge_domain.`,
				},
				resource.Attribute{
					Name:        "vmac",
					Description: `(Optional) Virtual MAC address of the BD/SVI. This is used when the BD is extended to multiple sites using l2 Outside.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Bridge Domain.`,
				},
				resource.Attribute{
					Name:        "optimize_wan_bandwidth",
					Description: `(Optional) Flag to enable OptimizeWanBandwidth between sites.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object bridge_domain.`,
				},
				resource.Attribute{
					Name:        "arp_flood",
					Description: `(Optional) A property to specify whether ARP flooding is enabled. If flooding is disabled, unicast routing will be performed on the target IP address.`,
				},
				resource.Attribute{
					Name:        "ep_clear",
					Description: `(Optional) Represents the parameter used by the node (i.e. Leaf) to clear all EPs in all leaves for this BD.`,
				},
				resource.Attribute{
					Name:        "ep_move_detect_mode",
					Description: `(Optional) The End Point move detection option uses the Gratuitous Address Resolution Protocol (GARP). A gratuitous ARP is an ARP broadcast-type of packet that is used to verify that no other device on the network has the same IP address as the sending device.`,
				},
				resource.Attribute{
					Name:        "host_based_routing",
					Description: `(Optional) Enables advertising host routes out of l3outs of this BD.`,
				},
				resource.Attribute{
					Name:        "intersite_bum_traffic_allow",
					Description: `(Optional) Control whether BUM traffic is allowed between sites.`,
				},
				resource.Attribute{
					Name:        "intersite_l2_stretch",
					Description: `(Optional) Flag to enable l2Stretch between sites.`,
				},
				resource.Attribute{
					Name:        "ip_learning",
					Description: `(Optional) Endpoint Dataplane Learning.`,
				},
				resource.Attribute{
					Name:        "ipv6_mcast_allow",
					Description: `(Optional) Flag to indicate multicast IpV6 is allowed or not.`,
				},
				resource.Attribute{
					Name:        "limit_ip_learn_to_subnets",
					Description: `(Optional) Limits IP address learning to the bridge domain subnets only. Every BD can have multiple subnets associated with it. By default, all IPs are learned.`,
				},
				resource.Attribute{
					Name:        "ll_addr",
					Description: `(Optional) override of system generated ipv6 link-local address.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Optional) The MAC address of the bridge domain (BD) or switched virtual interface (SVI). Every BD by default takes the fabric-wide default MAC address. You can override that address with a different one. By default the BD will take a 00:22:BD:F8:19:FF mac address.`,
				},
				resource.Attribute{
					Name:        "mcast_allow",
					Description: `(Optional) Flag to indicate if multicast is enabled for IpV4 addresses.`,
				},
				resource.Attribute{
					Name:        "multi_dst_pkt_act",
					Description: `(Optional) The multiple destination forwarding method for L2 Multicast, Broadcast, and Link Layer traffic types.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object bridge_domain.`,
				},
				resource.Attribute{
					Name:        "bridge_domain_type",
					Description: `(Optional) The specific type of the object or component.`,
				},
				resource.Attribute{
					Name:        "unicast_route",
					Description: `(Optional) The forwarding method based on predefined forwarding criteria (IP or MAC address).`,
				},
				resource.Attribute{
					Name:        "unk_mac_ucast_act",
					Description: `(Optional) The forwarding method for unknown layer 2 destinations.`,
				},
				resource.Attribute{
					Name:        "unk_mcast_act",
					Description: `(Optional) The parameter used by the node (i.e. a leaf) for forwarding data for an unknown multicast destination.`,
				},
				resource.Attribute{
					Name:        "v6unk_mcast_act",
					Description: `(Optional) v6unk_mcast_act for object bridge_domain.`,
				},
				resource.Attribute{
					Name:        "vmac",
					Description: `(Optional) Virtual MAC address of the BD/SVI. This is used when the BD is extended to multiple sites using l2 Outside.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cdp_interface_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI CDP Interface Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cdp_interface_policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the CDP Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) administrative state`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cdp_interface_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cdp_interface_policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the CDP Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) administrative state`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cdp_interface_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cdp_interface_policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_client_end_point",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Client End Point`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_epg_dn",
					Description: `(Required) Distinguished name of parent ApplicationEPG object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) name of Object client end point.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Optional) Mac address of the object client end point.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) ip address of the object client end point.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Optional) vlan for the object client end point. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set as all Dns for matching the Client End Point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects",
					Description: `list of all client end point objects which matched to the given filter attributes.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.name",
					Description: `Name of object client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.mac",
					Description: `Mac address of object client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.ip",
					Description: `IP address of object client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.vlan",
					Description: `vlan of client end point object.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.tenant_name",
					Description: `parent Tenant name for matched client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.vrf_name",
					Description: `parent vrf name for matched client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.application_profile_name",
					Description: `parent application profile name for matched client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.epg_name",
					Description: `parent epg name for matched client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.l2out_name",
					Description: `parent l2out name for matched client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.instance_profile_name",
					Description: `parent instance profile name for matched client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.endpoint_path",
					Description: `list of endpoint paths associated with client end point.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set as all Dns for matching the Client End Point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects",
					Description: `list of all client end point objects which matched to the given filter attributes.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.name",
					Description: `Name of object client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.mac",
					Description: `Mac address of object client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.ip",
					Description: `IP address of object client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.vlan",
					Description: `vlan of client end point object.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.tenant_name",
					Description: `parent Tenant name for matched client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.vrf_name",
					Description: `parent vrf name for matched client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.application_profile_name",
					Description: `parent application profile name for matched client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.epg_name",
					Description: `parent epg name for matched client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.l2out_name",
					Description: `parent l2out name for matched client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.instance_profile_name",
					Description: `parent instance profile name for matched client end point.`,
				},
				resource.Attribute{
					Name:        "fvcep_objects.endpoint_path",
					Description: `list of endpoint paths associated with client end point.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_applicationcontainer",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Application container`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_applicationcontainer. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Application container.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_applicationcontainer.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_applicationcontainer.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Application container.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_applicationcontainer.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_applicationcontainer.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_availability_zone",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Availability Zone`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_providers_region_dn",
					Description: `(Required) Distinguished name of parent CloudProvidersRegion object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_availability_zone. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Availability Zone.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_availability_zone.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_availability_zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Availability Zone.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_availability_zone.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_availability_zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_aws_provider",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud AWS Provider`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud AWS Provider.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `(Optional) access_key_id for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) account_id for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) email address of the local user`,
				},
				resource.Attribute{
					Name:        "http_proxy",
					Description: `(Optional) http_proxy for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "is_account_in_org",
					Description: `(Optional) is_account_in_org for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "is_trusted",
					Description: `(Optional) is_trusted for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `(Optional) provider_id for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) region for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Optional) secret_access_key for object cloud_aws_provider.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud AWS Provider.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `(Optional) access_key_id for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) account_id for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) email address of the local user`,
				},
				resource.Attribute{
					Name:        "http_proxy",
					Description: `(Optional) http_proxy for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "is_account_in_org",
					Description: `(Optional) is_account_in_org for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "is_trusted",
					Description: `(Optional) is_trusted for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `(Optional) provider_id for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) region for object cloud_aws_provider.`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Optional) secret_access_key for object cloud_aws_provider.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_cidr_pool",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud CIDR Pool`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_context_profile_dn",
					Description: `(Required) Distinguished name of parent CloudContextProfile object.`,
				},
				resource.Attribute{
					Name:        "addr",
					Description: `(Required) CIDR IPv4 block. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud CIDR Pool.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_cidr_pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_cidr_pool.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Optional) This will represent whether CIDR is primary CIDR or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud CIDR Pool.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_cidr_pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_cidr_pool.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Optional) This will represent whether CIDR is primary CIDR or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_context_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Context Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud-ctx-profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Context profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `annotation for object Cloud Context profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `name_alias for object Cloud Context Profile.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The specific type of the object or component.`,
				},
				resource.Attribute{
					Name:        "primary_cidr",
					Description: `Primary CIDR block of Cloud Context profile.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `AWS region in which profile is created.`,
				},
				resource.Attribute{
					Name:        "hub_network",
					Description: `hub network Dn which enables Transit Gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Context profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `annotation for object Cloud Context profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `name_alias for object Cloud Context Profile.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The specific type of the object or component.`,
				},
				resource.Attribute{
					Name:        "primary_cidr",
					Description: `Primary CIDR block of Cloud Context profile.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `AWS region in which profile is created.`,
				},
				resource.Attribute{
					Name:        "hub_network",
					Description: `hub network Dn which enables Transit Gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_domain_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Domain Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Domain Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_domain_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_domain_profile.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Optional) site_id for object cloud_domain_profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Domain Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_domain_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_domain_profile.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Optional) site_id for object cloud_domain_profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_endpoint_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Endpoint Selector`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_epg_dn",
					Description: `(Required) Distinguished name of parent CloudEPg object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_endpoint_selector. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Endpoint Selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_endpoint_selector.`,
				},
				resource.Attribute{
					Name:        "match_expression",
					Description: `(Optional) Match expression for the endpoint selector to select EP on criteria.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_endpoint_selector.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Endpoint Selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_endpoint_selector.`,
				},
				resource.Attribute{
					Name:        "match_expression",
					Description: `(Optional) Match expression for the endpoint selector to select EP on criteria.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_endpoint_selector.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_endpoint_selectorfor_external_epgs",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Endpoint Selector for External EPgs`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_external_epg_dn",
					Description: `(Required) Distinguished name of parent CloudExternalEPg object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_endpoint_selectorfor_external_epgs. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Endpoint Selector for External EPgs.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_endpoint_selectorfor_external_epgs.`,
				},
				resource.Attribute{
					Name:        "is_shared",
					Description: `(Optional) For Selectors set the shared route control.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_endpoint_selectorfor_external_epgs.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional) Subnet from which EP to select.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Endpoint Selector for External EPgs.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_endpoint_selectorfor_external_epgs.`,
				},
				resource.Attribute{
					Name:        "is_shared",
					Description: `(Optional) For Selectors set the shared route control.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_endpoint_selectorfor_external_epgs.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional) Subnet from which EP to select.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_epg",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud EPg`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_applicationcontainer_dn",
					Description: `(Required) Distinguished name of parent CloudApplicationcontainer object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_epg. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud EPg.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_epg.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) exception_tag for object cloud_epg.`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_epg.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) qos priority class id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud EPg.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_epg.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) exception_tag for object cloud_epg.`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_epg.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) qos priority class id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_external_epg",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud External EPg`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_applicationcontainer_dn",
					Description: `(Required) Distinguished name of parent CloudApplicationcontainer object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_external_epg. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud External EPg.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_external_epg.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) exception_tag for object cloud_external_epg.`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_external_epg.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) qos priority class id.`,
				},
				resource.Attribute{
					Name:        "route_reachability",
					Description: `(Optional) Route reachability for this EPG.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud External EPg.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_external_epg.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) exception_tag for object cloud_external_epg.`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_external_epg.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) qos priority class id.`,
				},
				resource.Attribute{
					Name:        "route_reachability",
					Description: `(Optional) Route reachability for this EPG.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_provider_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Provider Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vendor",
					Description: `(Required) vendor of Object cloud_provider_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Provider Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_provider_profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Provider Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_provider_profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_providers_region",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Providers Region`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_provider_profile_dn",
					Description: `(Required) Distinguished name of parent CloudProviderProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object cloud_providers_region. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Providers Region.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) administrative state of the object or policy`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_providers_region.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_providers_region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Providers Region.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) administrative state of the object or policy`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_providers_region.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_providers_region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_subnet",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Subnet`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_cidr_pool_dn",
					Description: `(Required) Distinguished name of parent CloudCIDRPool object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) CIDR block of Object cloud_subnet. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for object cloud subnet.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_subnet.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The domain applicable to the capability.`,
				},
				resource.Attribute{
					Name:        "usage",
					Description: `(Optional) The usage of the port. This property shows how the port is used.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for object cloud subnet.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_subnet.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The domain applicable to the capability.`,
				},
				resource.Attribute{
					Name:        "usage",
					Description: `(Optional) The usage of the port. This property shows how the port is used.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_vpn_gateway",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Vpn Gateway`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_context_profile_dn",
					Description: `(Required) Distinguished name of parent CloudContextProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object cloud_router_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Router Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object cloud_router_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_router_profile.`,
				},
				resource.Attribute{
					Name:        "num_instances",
					Description: `(Optional) num_instances for object cloud_router_profile.`,
				},
				resource.Attribute{
					Name:        "cloud_router_profile_type",
					Description: `(Optional) Component type`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Router Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object cloud_router_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object cloud_router_profile.`,
				},
				resource.Attribute{
					Name:        "num_instances",
					Description: `(Optional) num_instances for object cloud_router_profile.`,
				},
				resource.Attribute{
					Name:        "cloud_router_profile_type",
					Description: `(Optional) Component type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_configuration_export_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Configuration Export Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object configuration_export_policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Configuration Export Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) admin state of the export policy`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) export data format`,
				},
				resource.Attribute{
					Name:        "include_secure_fields",
					Description: `(Optional) include_secure_fields for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "max_snapshot_count",
					Description: `(Optional) max_snapshot_count for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `(Optional) snapshot for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "target_dn",
					Description: `(Optional) target export object`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Configuration Export Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) admin state of the export policy`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) export data format`,
				},
				resource.Attribute{
					Name:        "include_secure_fields",
					Description: `(Optional) include_secure_fields for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "max_snapshot_count",
					Description: `(Optional) max_snapshot_count for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `(Optional) snapshot for object configuration_export_policy.`,
				},
				resource.Attribute{
					Name:        "target_dn",
					Description: `(Optional) target export object`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_configuration_import_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Configuration Import Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object configuration_import_policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Configuration Import Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) admin state of the import`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object configuration_import_policy.`,
				},
				resource.Attribute{
					Name:        "fail_on_decrypt_errors",
					Description: `(Optional) fail_on_decrypt_errors for object configuration_import_policy.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `(Optional) import file name`,
				},
				resource.Attribute{
					Name:        "import_mode",
					Description: `(Optional) data import mode`,
				},
				resource.Attribute{
					Name:        "import_type",
					Description: `(Optional) data import type`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object configuration_import_policy.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `(Optional) snapshot for object configuration_import_policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Configuration Import Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) admin state of the import`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object configuration_import_policy.`,
				},
				resource.Attribute{
					Name:        "fail_on_decrypt_errors",
					Description: `(Optional) fail_on_decrypt_errors for object configuration_import_policy.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `(Optional) import file name`,
				},
				resource.Attribute{
					Name:        "import_mode",
					Description: `(Optional) data import mode`,
				},
				resource.Attribute{
					Name:        "import_type",
					Description: `(Optional) data import type`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object configuration_import_policy.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `(Optional) snapshot for object configuration_import_policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_connection",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Connection`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "l4_l7_service_graph_template_dn",
					Description: `(Required) distinguished name of parent L4-L7 Service Graph Template object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object connection. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Connection.`,
				},
				resource.Attribute{
					Name:        "adj_type",
					Description: `Connector adjacency type`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `Annotation for object connection.`,
				},
				resource.Attribute{
					Name:        "conn_dir",
					Description: `Connection dir for object connection.`,
				},
				resource.Attribute{
					Name:        "conn_type",
					Description: `Connection type for object connection.`,
				},
				resource.Attribute{
					Name:        "direct_connect",
					Description: `Direct connect for object connection.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `Name alias for object connection.`,
				},
				resource.Attribute{
					Name:        "unicast_route",
					Description: `Unicast route for object connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Connection.`,
				},
				resource.Attribute{
					Name:        "adj_type",
					Description: `Connector adjacency type`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `Annotation for object connection.`,
				},
				resource.Attribute{
					Name:        "conn_dir",
					Description: `Connection dir for object connection.`,
				},
				resource.Attribute{
					Name:        "conn_type",
					Description: `Connection type for object connection.`,
				},
				resource.Attribute{
					Name:        "direct_connect",
					Description: `Direct connect for object connection.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `Name alias for object connection.`,
				},
				resource.Attribute{
					Name:        "unicast_route",
					Description: `Unicast route for object connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_contract",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Contract`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object contract. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Contract.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object contract.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object contract.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) priority level of the service contract.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Represents the scope of this contract. If the scope is set as application-profile, the epg can only communicate with epgs in the same application-profile.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Contract.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object contract.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object contract.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) priority level of the service contract.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Represents the scope of this contract. If the scope is set as application-profile, the epg can only communicate with epgs in the same application-profile.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_contract_subject",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Contract Subject`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "contract_dn",
					Description: `(Required) Distinguished name of parent Contract object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object contract_subject. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Contract Subject.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object contract_subject.`,
				},
				resource.Attribute{
					Name:        "cons_match_t",
					Description: `(Optional) The subject match criteria across consumers.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object contract_subject.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The priority level of a sub application running behind an endpoint group, such as an Exchange server.`,
				},
				resource.Attribute{
					Name:        "prov_match_t",
					Description: `(Optional) The subject match criteria across consumers.`,
				},
				resource.Attribute{
					Name:        "rev_flt_ports",
					Description: `(Optional) enables filter to apply on ingress and egress traffic.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Contract Subject.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object contract_subject.`,
				},
				resource.Attribute{
					Name:        "cons_match_t",
					Description: `(Optional) The subject match criteria across consumers.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object contract_subject.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The priority level of a sub application running behind an endpoint group, such as an Exchange server.`,
				},
				resource.Attribute{
					Name:        "prov_match_t",
					Description: `(Optional) The subject match criteria across consumers.`,
				},
				resource.Attribute{
					Name:        "rev_flt_ports",
					Description: `(Optional) enables filter to apply on ingress and egress traffic.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_destination_of_redirected_traffic",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Destination of redirected traffic`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_redirect_policy_dn",
					Description: `(Required) Distinguished name of parent Service Redirect Policy object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) ip of Object destination of redirected traffic. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Destination of redirected traffic.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object destination of redirected traffic.`,
				},
				resource.Attribute{
					Name:        "dest_name",
					Description: `(Optional) destination name to which data was exported.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) ip address.`,
				},
				resource.Attribute{
					Name:        "ip2",
					Description: `(Optional) ip2 for object destination of redirected traffic.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Optional) mac address.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object destination of redirected traffic.`,
				},
				resource.Attribute{
					Name:        "pod_id",
					Description: `(Optional) pod id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Destination of redirected traffic.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object destination of redirected traffic.`,
				},
				resource.Attribute{
					Name:        "dest_name",
					Description: `(Optional) destination name to which data was exported.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) ip address.`,
				},
				resource.Attribute{
					Name:        "ip2",
					Description: `(Optional) ip2 for object destination of redirected traffic.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Optional) mac address.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object destination of redirected traffic.`,
				},
				resource.Attribute{
					Name:        "pod_id",
					Description: `(Optional) pod id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_dhcp_option",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI DHCP Option`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_dhcp_option_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI DHCP Option Policy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_dhcp_relay_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI DHCP Relay Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object DHCP Relay Policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the DHCP Relay Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object DHCP Relay Policy.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Dhcp Relay policy mode.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object DHCP Relay Policy.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) Owner of the target relay servers.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the DHCP Relay Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object DHCP Relay Policy.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Dhcp Relay policy mode.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object DHCP Relay Policy.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) Owner of the target relay servers.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_end_point_retention_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI End Point Retention Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object end_point_retention_policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the End Point Retention Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object end_point_retention_policy.`,
				},
				resource.Attribute{
					Name:        "bounce_age_intvl",
					Description: `(Optional) The aging interval for a bounce entry. When an endpoint (VM) migrates to another switch, the endpoint is marked as bouncing for the specified aging interval and is deleted afterwards.`,
				},
				resource.Attribute{
					Name:        "bounce_trig",
					Description: `(Optional) Specifies whether to install the bounce entry by RARP flood or by COOP protocol. Allowed values are "rarp-flood" and "protocol".`,
				},
				resource.Attribute{
					Name:        "hold_intvl",
					Description: `(Optional) A time period during which new endpoint learn events will not be honored. This interval is triggered when the maximum endpoint move frequency is exceeded.`,
				},
				resource.Attribute{
					Name:        "local_ep_age_intvl",
					Description: `(Optional) The aging interval for all local endpoints learned in this bridge domain. When 75% of the interval is reached, 3 ARP requests are sent to verify the existence of the endpoint. If no response is received, the endpoint is deleted.`,
				},
				resource.Attribute{
					Name:        "move_freq",
					Description: `(Optional) A maximum allowed number of endpoint moves per second. If the move frequency is exceeded, the hold interval is triggered, and new endpoint learn events will not be honored until after the hold interval expires.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object end_point_retention_policy.`,
				},
				resource.Attribute{
					Name:        "remote_ep_age_intvl",
					Description: `(Optional) The aging interval for all remote endpoints learned in this bridge domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the End Point Retention Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object end_point_retention_policy.`,
				},
				resource.Attribute{
					Name:        "bounce_age_intvl",
					Description: `(Optional) The aging interval for a bounce entry. When an endpoint (VM) migrates to another switch, the endpoint is marked as bouncing for the specified aging interval and is deleted afterwards.`,
				},
				resource.Attribute{
					Name:        "bounce_trig",
					Description: `(Optional) Specifies whether to install the bounce entry by RARP flood or by COOP protocol. Allowed values are "rarp-flood" and "protocol".`,
				},
				resource.Attribute{
					Name:        "hold_intvl",
					Description: `(Optional) A time period during which new endpoint learn events will not be honored. This interval is triggered when the maximum endpoint move frequency is exceeded.`,
				},
				resource.Attribute{
					Name:        "local_ep_age_intvl",
					Description: `(Optional) The aging interval for all local endpoints learned in this bridge domain. When 75% of the interval is reached, 3 ARP requests are sent to verify the existence of the endpoint. If no response is received, the endpoint is deleted.`,
				},
				resource.Attribute{
					Name:        "move_freq",
					Description: `(Optional) A maximum allowed number of endpoint moves per second. If the move frequency is exceeded, the hold interval is triggered, and new endpoint learn events will not be honored until after the hold interval expires.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object end_point_retention_policy.`,
				},
				resource.Attribute{
					Name:        "remote_ep_age_intvl",
					Description: `(Optional) The aging interval for all remote endpoints learned in this bridge domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_epg_to_contract",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI EPG to contract relationship.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_epg_dn",
					Description: `(Required) Distinguished name of Parent epg.`,
				},
				resource.Attribute{
					Name:        "contract_name",
					Description: `(Required) Name of the contract.`,
				},
				resource.Attribute{
					Name:        "contract_type",
					Description: `(Required) Type of relationship. Allowed values are ` + "`" + `consumer` + "`" + ` and ` + "`" + `provider` + "`" + `. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the provider/consumer contract.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) Provider matching criteria.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) Priority of relation object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the provider/consumer contract.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) Provider matching criteria.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) Priority of relation object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_epg_to_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI epg to Domain`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_epg_dn",
					Description: `(Required) Distinguished name of parent ApplicationEPG object.`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Required) vmm domain instance. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object domain.`,
				},
				resource.Attribute{
					Name:        "binding_type",
					Description: `(Optional) binding_type for object domain.`,
				},
				resource.Attribute{
					Name:        "allow_micro_seg",
					Description: `(Optional) boolean flag for allow micro segment. default value will be "false". "true" maps to class_pref="useg" and "false maps to class_pref="encap"`,
				},
				resource.Attribute{
					Name:        "delimiter",
					Description: `(Optional) delimiter for object domain.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) port encapsulation`,
				},
				resource.Attribute{
					Name:        "encap_mode",
					Description: `(Optional) encap_mode for object domain.`,
				},
				resource.Attribute{
					Name:        "epg_cos",
					Description: `(Optional) epg_cos for object domain.`,
				},
				resource.Attribute{
					Name:        "epg_cos_pref",
					Description: `(Optional) epg_cos_pref for object domain.`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) determines when policies are pushed to cam`,
				},
				resource.Attribute{
					Name:        "lag_policy_name",
					Description: `(Optional) lag_policy_name for object domain.`,
				},
				resource.Attribute{
					Name:        "netflow_dir",
					Description: `(Optional) netflow_dir for object domain.`,
				},
				resource.Attribute{
					Name:        "netflow_pref",
					Description: `(Optional) netflow_pref for object domain.`,
				},
				resource.Attribute{
					Name:        "num_ports",
					Description: `(Optional) number of ports existing operationally in module`,
				},
				resource.Attribute{
					Name:        "port_allocation",
					Description: `(Optional) port_allocation for object domain.`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) primary_encap for object domain.`,
				},
				resource.Attribute{
					Name:        "primary_encap_inner",
					Description: `(Optional) primary_encap_inner for object domain.`,
				},
				resource.Attribute{
					Name:        "res_imedcy",
					Description: `(Optional) policy resolution`,
				},
				resource.Attribute{
					Name:        "secondary_encap_inner",
					Description: `(Optional) secondary_encap_inner for object domain.`,
				},
				resource.Attribute{
					Name:        "switching_mode",
					Description: `(Optional) switching_mode for object domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object domain.`,
				},
				resource.Attribute{
					Name:        "binding_type",
					Description: `(Optional) binding_type for object domain.`,
				},
				resource.Attribute{
					Name:        "allow_micro_seg",
					Description: `(Optional) boolean flag for allow micro segment. default value will be "false". "true" maps to class_pref="useg" and "false maps to class_pref="encap"`,
				},
				resource.Attribute{
					Name:        "delimiter",
					Description: `(Optional) delimiter for object domain.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) port encapsulation`,
				},
				resource.Attribute{
					Name:        "encap_mode",
					Description: `(Optional) encap_mode for object domain.`,
				},
				resource.Attribute{
					Name:        "epg_cos",
					Description: `(Optional) epg_cos for object domain.`,
				},
				resource.Attribute{
					Name:        "epg_cos_pref",
					Description: `(Optional) epg_cos_pref for object domain.`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) determines when policies are pushed to cam`,
				},
				resource.Attribute{
					Name:        "lag_policy_name",
					Description: `(Optional) lag_policy_name for object domain.`,
				},
				resource.Attribute{
					Name:        "netflow_dir",
					Description: `(Optional) netflow_dir for object domain.`,
				},
				resource.Attribute{
					Name:        "netflow_pref",
					Description: `(Optional) netflow_pref for object domain.`,
				},
				resource.Attribute{
					Name:        "num_ports",
					Description: `(Optional) number of ports existing operationally in module`,
				},
				resource.Attribute{
					Name:        "port_allocation",
					Description: `(Optional) port_allocation for object domain.`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) primary_encap for object domain.`,
				},
				resource.Attribute{
					Name:        "primary_encap_inner",
					Description: `(Optional) primary_encap_inner for object domain.`,
				},
				resource.Attribute{
					Name:        "res_imedcy",
					Description: `(Optional) policy resolution`,
				},
				resource.Attribute{
					Name:        "secondary_encap_inner",
					Description: `(Optional) secondary_encap_inner for object domain.`,
				},
				resource.Attribute{
					Name:        "switching_mode",
					Description: `(Optional) switching_mode for object domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_epg_to_static_path",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI EPG to Static Path`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_epg_dn",
					Description: `(Required) Distinguished name of parent ApplicationEPG object.`,
				},
				resource.Attribute{
					Name:        "tDn",
					Description: `(Required) tDn of Object static_path. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Static Path.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object static_path.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) encapsulation`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) immediacy`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) mode of the static association with the path`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) primary_encap for object static_path.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Static Path.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object static_path.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) encapsulation`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) immediacy`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) mode of the static association with the path`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) primary_encap for object static_path.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_epgs_using_function",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI EPGs Using Function`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_generic_dn",
					Description: `(Required) Distinguished name of parent AccessGeneric object.`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Required) tDn of Object epgs_using_function. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the EPGs Using Function.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object epgs_using_function.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) vlan number encap`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) instrumentation immediacy`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) bgp domain mode`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) primary_encap for object epgs_using_function.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the EPGs Using Function.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object epgs_using_function.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) vlan number encap`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) instrumentation immediacy`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) bgp domain mode`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) primary_encap for object epgs_using_function.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_external_network_instance_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI External Network Instance Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "l3_outside_dn",
					Description: `(Required) Distinguished name of parent L3Outside object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object external_network_instance_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the External Network Instance Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object external_network_instance_profile.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) exception_tag for object external_network_instance_profile.`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object external_network_instance_profile.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The QoS priority class identifier.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the External Network Instance Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object external_network_instance_profile.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) exception_tag for object external_network_instance_profile.`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object external_network_instance_profile.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The QoS priority class identifier.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fabric_if_pol",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI fabric if pol`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object fabric if pol. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the fabric if pol.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "auto_neg",
					Description: `(Optional) policy auto-negotiation`,
				},
				resource.Attribute{
					Name:        "fec_mode",
					Description: `(Optional) forwarding error correction`,
				},
				resource.Attribute{
					Name:        "link_debounce",
					Description: `(Optional) link debounce interval`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name alias for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) port speed`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the fabric if pol.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "auto_neg",
					Description: `(Optional) policy auto-negotiation`,
				},
				resource.Attribute{
					Name:        "fec_mode",
					Description: `(Optional) forwarding error correction`,
				},
				resource.Attribute{
					Name:        "link_debounce",
					Description: `(Optional) link debounce interval`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name alias for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) port speed`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fabric_node_member",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Fabric Node Member`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) serial of Object fabric_node_member. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Fabric Node Member.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fabric_node_member.`,
				},
				resource.Attribute{
					Name:        "ext_pool_id",
					Description: `(Optional) ext_pool_id for object fabric_node_member.`,
				},
				resource.Attribute{
					Name:        "fabric_id",
					Description: `(Optional) place holder for a value`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object fabric_node_member.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `(Optional) node id`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Optional) node_type for object fabric_node_member.`,
				},
				resource.Attribute{
					Name:        "pod_id",
					Description: `(Optional) pod id`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) system role type`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Optional) serial number`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Fabric Node Member.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fabric_node_member.`,
				},
				resource.Attribute{
					Name:        "ext_pool_id",
					Description: `(Optional) ext_pool_id for object fabric_node_member.`,
				},
				resource.Attribute{
					Name:        "fabric_id",
					Description: `(Optional) place holder for a value`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object fabric_node_member.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `(Optional) node id`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Optional) node_type for object fabric_node_member.`,
				},
				resource.Attribute{
					Name:        "pod_id",
					Description: `(Optional) pod id`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) system role type`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Optional) serial number`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fabric_path_ep",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Fabric Path End point`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pod_id",
					Description: `(Required) pod ID for Object fabric path endpoint.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `(Required) node ID for Object fabric path endpoint.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object fabric path endpoint. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Fabric Path End-point.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Fabric Path End-point.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fc_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI FC Domain`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object fc_domain. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the FC Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fc_domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object fc_domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the FC Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fc_domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object fc_domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fex_bundle_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Fex Bundle Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fex_profile_dn",
					Description: `(Required) Distinguished name of parent FEXProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object fex_bundle_group. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Fex Bundle Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fex_bundle_group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object fex_bundle_group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Fex Bundle Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fex_bundle_group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object fex_bundle_group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fex_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI FEX Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object fex_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the FEX Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fex_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object fex_profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the FEX Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fex_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object fex_profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_filter",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Filter`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object filter. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Filter.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object filter.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object filter.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Filter.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object filter.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object filter.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_filter_entry",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Filter Entry`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter_dn",
					Description: `(Required) Distinguished name of parent Filter object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object filter_entry. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Filter Entry.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object filter_entry.`,
				},
				resource.Attribute{
					Name:        "apply_to_frag",
					Description: `(Optional) Flag to determine whether to apply changes to fragment.`,
				},
				resource.Attribute{
					Name:        "arp_opc",
					Description: `(Optional) open peripheral codes.`,
				},
				resource.Attribute{
					Name:        "d_from_port",
					Description: `(Optional) Destination From Port.`,
				},
				resource.Attribute{
					Name:        "d_to_port",
					Description: `(Optional) Destination To Port.`,
				},
				resource.Attribute{
					Name:        "ether_t",
					Description: `(Optional) ether type for the entry.`,
				},
				resource.Attribute{
					Name:        "icmpv4_t",
					Description: `(Optional) ICMPv4 message type; used when ip_protocol is icmp.`,
				},
				resource.Attribute{
					Name:        "icmpv6_t",
					Description: `(Optional) ICMPv6 message type; used when ip_protocol is icmpv6.`,
				},
				resource.Attribute{
					Name:        "match_dscp",
					Description: `(Optional) The matching differentiated services code point (DSCP) of the path attached to the layer 3 outside profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object filter_entry.`,
				},
				resource.Attribute{
					Name:        "prot",
					Description: `(Optional) level 3 ip protocol.`,
				},
				resource.Attribute{
					Name:        "s_from_port",
					Description: `(Optional) Source From Port.`,
				},
				resource.Attribute{
					Name:        "s_to_port",
					Description: `(Optional) Source To Port.`,
				},
				resource.Attribute{
					Name:        "stateful",
					Description: `(Optional) Determines if entry is stateful or not.`,
				},
				resource.Attribute{
					Name:        "tcp_rules",
					Description: `(Optional) TCP Session Rules.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Filter Entry.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object filter_entry.`,
				},
				resource.Attribute{
					Name:        "apply_to_frag",
					Description: `(Optional) Flag to determine whether to apply changes to fragment.`,
				},
				resource.Attribute{
					Name:        "arp_opc",
					Description: `(Optional) open peripheral codes.`,
				},
				resource.Attribute{
					Name:        "d_from_port",
					Description: `(Optional) Destination From Port.`,
				},
				resource.Attribute{
					Name:        "d_to_port",
					Description: `(Optional) Destination To Port.`,
				},
				resource.Attribute{
					Name:        "ether_t",
					Description: `(Optional) ether type for the entry.`,
				},
				resource.Attribute{
					Name:        "icmpv4_t",
					Description: `(Optional) ICMPv4 message type; used when ip_protocol is icmp.`,
				},
				resource.Attribute{
					Name:        "icmpv6_t",
					Description: `(Optional) ICMPv6 message type; used when ip_protocol is icmpv6.`,
				},
				resource.Attribute{
					Name:        "match_dscp",
					Description: `(Optional) The matching differentiated services code point (DSCP) of the path attached to the layer 3 outside profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object filter_entry.`,
				},
				resource.Attribute{
					Name:        "prot",
					Description: `(Optional) level 3 ip protocol.`,
				},
				resource.Attribute{
					Name:        "s_from_port",
					Description: `(Optional) Source From Port.`,
				},
				resource.Attribute{
					Name:        "s_to_port",
					Description: `(Optional) Source To Port.`,
				},
				resource.Attribute{
					Name:        "stateful",
					Description: `(Optional) Determines if entry is stateful or not.`,
				},
				resource.Attribute{
					Name:        "tcp_rules",
					Description: `(Optional) TCP Session Rules.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_firmware_download_task",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Firmware Download Task`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object firmware_download_task. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Firmware Download Task.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "auth_pass",
					Description: `(Optional) authentication type`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) ospf authentication type specifier`,
				},
				resource.Attribute{
					Name:        "dnld_task_flip",
					Description: `(Optional) dnld_task_flip for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "identity_private_key_contents",
					Description: `(Optional) identity_private_key_contents for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "identity_private_key_passphrase",
					Description: `(Optional) identity_private_key_passphrase for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "identity_public_key_contents",
					Description: `(Optional) identity_public_key_contents for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "load_catalog_if_exists_and_newer",
					Description: `(Optional) tracks to load the contained catalog or newer`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) password/key string`,
				},
				resource.Attribute{
					Name:        "polling_interval",
					Description: `(Optional) polling interval`,
				},
				resource.Attribute{
					Name:        "proto",
					Description: `(Optional) download protocol`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) URL of image of source`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) username for source`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Firmware Download Task.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "auth_pass",
					Description: `(Optional) authentication type`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) ospf authentication type specifier`,
				},
				resource.Attribute{
					Name:        "dnld_task_flip",
					Description: `(Optional) dnld_task_flip for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "identity_private_key_contents",
					Description: `(Optional) identity_private_key_contents for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "identity_private_key_passphrase",
					Description: `(Optional) identity_private_key_passphrase for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "identity_public_key_contents",
					Description: `(Optional) identity_public_key_contents for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "load_catalog_if_exists_and_newer",
					Description: `(Optional) tracks to load the contained catalog or newer`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object firmware_download_task.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) password/key string`,
				},
				resource.Attribute{
					Name:        "polling_interval",
					Description: `(Optional) polling interval`,
				},
				resource.Attribute{
					Name:        "proto",
					Description: `(Optional) download protocol`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) URL of image of source`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) username for source`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_firmware_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Firmware Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object firmware_group. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Firmware Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object firmware_group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object firmware_group.`,
				},
				resource.Attribute{
					Name:        "firmware_group_type",
					Description: `(Optional) component type`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Firmware Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object firmware_group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object firmware_group.`,
				},
				resource.Attribute{
					Name:        "firmware_group_type",
					Description: `(Optional) component type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_firmware_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Firmware Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object firmware_policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Firmware Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object firmware_policy.`,
				},
				resource.Attribute{
					Name:        "effective_on_reboot",
					Description: `(Optional) firmware version effective on reboot selection`,
				},
				resource.Attribute{
					Name:        "ignore_compat",
					Description: `(Optional) whether compatibility check required`,
				},
				resource.Attribute{
					Name:        "internal_label",
					Description: `(Optional) firmware label`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object firmware_policy.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) firmware version`,
				},
				resource.Attribute{
					Name:        "version_check_override",
					Description: `(Optional) version check override`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Firmware Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object firmware_policy.`,
				},
				resource.Attribute{
					Name:        "effective_on_reboot",
					Description: `(Optional) firmware version effective on reboot selection`,
				},
				resource.Attribute{
					Name:        "ignore_compat",
					Description: `(Optional) whether compatibility check required`,
				},
				resource.Attribute{
					Name:        "internal_label",
					Description: `(Optional) firmware label`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object firmware_policy.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) firmware version`,
				},
				resource.Attribute{
					Name:        "version_check_override",
					Description: `(Optional) version check override`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_function_node",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Function Node`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_imported_contract",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Imported Contract`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object imported_contract. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Imported Contract.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object imported_contract.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object imported_contract.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Imported Contract.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object imported_contract.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object imported_contract.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_interface_fc_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Interface FC Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object interface_fc_policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Interface FC Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object interface_fc_policy.`,
				},
				resource.Attribute{
					Name:        "automaxspeed",
					Description: `(Optional) automaxspeed for object interface_fc_policy.`,
				},
				resource.Attribute{
					Name:        "fill_pattern",
					Description: `(Optional) Fill Pattern for native FC ports.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object interface_fc_policy.`,
				},
				resource.Attribute{
					Name:        "port_mode",
					Description: `(Optional) In which mode Ports should be used.`,
				},
				resource.Attribute{
					Name:        "rx_bb_credit",
					Description: `(Optional) Receive buffer credits for native FC ports.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) cpu or port speed.`,
				},
				resource.Attribute{
					Name:        "trunk_mode",
					Description: `(Optional) Trunking on/off for native FC ports.Default value is OFF.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Interface FC Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object interface_fc_policy.`,
				},
				resource.Attribute{
					Name:        "automaxspeed",
					Description: `(Optional) automaxspeed for object interface_fc_policy.`,
				},
				resource.Attribute{
					Name:        "fill_pattern",
					Description: `(Optional) Fill Pattern for native FC ports.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object interface_fc_policy.`,
				},
				resource.Attribute{
					Name:        "port_mode",
					Description: `(Optional) In which mode Ports should be used.`,
				},
				resource.Attribute{
					Name:        "rx_bb_credit",
					Description: `(Optional) Receive buffer credits for native FC ports.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) cpu or port speed.`,
				},
				resource.Attribute{
					Name:        "trunk_mode",
					Description: `(Optional) Trunking on/off for native FC ports.Default value is OFF.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l2_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L2 Domain`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object L2 Domain. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L2 Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for L2 Domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for L2 Domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L2 Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for L2 Domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for L2 Domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l2_interface_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L2 Interface Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object l2_interface_policy. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L2 Interface Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object l2_interface_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object l2_interface_policy.`,
				},
				resource.Attribute{
					Name:        "qinq",
					Description: `(Optional) Determines if QinQ is disabled or if the port should be considered a core or edge port.`,
				},
				resource.Attribute{
					Name:        "vepa",
					Description: `(Optional) Determines if Virtual Ethernet Port Aggregator is disabled or enabled.`,
				},
				resource.Attribute{
					Name:        "vlan_scope",
					Description: `(Optional) The scope of the VLAN.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L2 Interface Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object l2_interface_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object l2_interface_policy.`,
				},
				resource.Attribute{
					Name:        "qinq",
					Description: `(Optional) Determines if QinQ is disabled or if the port should be considered a core or edge port.`,
				},
				resource.Attribute{
					Name:        "vepa",
					Description: `(Optional) Determines if Virtual Ethernet Port Aggregator is disabled or enabled.`,
				},
				resource.Attribute{
					Name:        "vlan_scope",
					Description: `(Optional) The scope of the VLAN.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l2_outside",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L2 Outside`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l2out_extepg",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L2-Out External EPg`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3_domain_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3 Domain Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object l3_domain_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L3 Domain Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object l3_domain_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object l3_domain_profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L3 Domain Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object l3_domain_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object l3_domain_profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3_ext_subnet",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI l3 extension subnet`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "external_network_instance_profile_dn",
					Description: `(Required) Distinguished name of parent ExternalNetworkInstanceProfile object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) ip of Object l3 extension subnet. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the l3 extension subnet.`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `(Optional) Aggregate Routes for l3 extension subnet.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object l3 extension subnet.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object l3 extension subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The list of domain applicable to the capability.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the l3 extension subnet.`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `(Optional) Aggregate Routes for l3 extension subnet.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object l3 extension subnet.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object l3 extension subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The list of domain applicable to the capability.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3_outside",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3 Outside`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object l3_outside. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L3 Outside.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object l3_outside.`,
				},
				resource.Attribute{
					Name:        "enforce_rtctrl",
					Description: `(Optional) enforce route control type`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object l3_outside.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L3 Outside.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object l3_outside.`,
				},
				resource.Attribute{
					Name:        "enforce_rtctrl",
					Description: `(Optional) enforce route control type`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object l3_outside.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l4_l7_service_graph_template",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L4-L7 Service Graph Template`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for L4-L7 service graph template object. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L4-L7 Service Graph Template.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for L4-L7 service graph template object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for L4-L7 service graph template object.`,
				},
				resource.Attribute{
					Name:        "l4_l7_service_graph_template_type",
					Description: `(Optional) Component type for L4-L7 service graph template object.`,
				},
				resource.Attribute{
					Name:        "ui_template_type",
					Description: `(Optional) UI template type for L4-L7 service graph template object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L4-L7 Service Graph Template.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for L4-L7 service graph template object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for L4-L7 service graph template object.`,
				},
				resource.Attribute{
					Name:        "l4_l7_service_graph_template_type",
					Description: `(Optional) Component type for L4-L7 service graph template object.`,
				},
				resource.Attribute{
					Name:        "ui_template_type",
					Description: `(Optional) UI template type for L4-L7 service graph template object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_lacp_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI LACP Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object lacp_policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LACP Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object lacp_policy.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) List of LAG control properties`,
				},
				resource.Attribute{
					Name:        "max_links",
					Description: `(Optional) Maximum number of links.`,
				},
				resource.Attribute{
					Name:        "min_links",
					Description: `(Optional) Minimum number of links in port channel.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Policy mode.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name_alias for object lacp_policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LACP Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object lacp_policy.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) List of LAG control properties`,
				},
				resource.Attribute{
					Name:        "max_links",
					Description: `(Optional) Maximum number of links.`,
				},
				resource.Attribute{
					Name:        "min_links",
					Description: `(Optional) Minimum number of links in port channel.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Policy mode.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name_alias for object lacp_policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_access_bundle_policy_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI leaf access bundle policy group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object leaf_access_bundle_policy_group. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the ACI leaf access bundle policy group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object aci_leaf_access_bundle_policy_group.`,
				},
				resource.Attribute{
					Name:        "lag_t",
					Description: `(Optional) The bundled ports group link aggregation type: port channel vs virtual port channel.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object aci_leaf_access_bundle_policy_group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the ACI leaf access bundle policy group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object aci_leaf_access_bundle_policy_group.`,
				},
				resource.Attribute{
					Name:        "lag_t",
					Description: `(Optional) The bundled ports group link aggregation type: port channel vs virtual port channel.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object aci_leaf_access_bundle_policy_group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_access_port_policy_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Leaf Access Port Policy Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object leaf_access_port_policy_group. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Leaf Access Port Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object leaf_access_port_policy_group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object leaf_access_port_policy_group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Leaf Access Port Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object leaf_access_port_policy_group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object leaf_access_port_policy_group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_breakout_port_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Leaf Breakout Port Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of leaf breakout port group object. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the leaf breakout port group object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for leaf breakout port group object.`,
				},
				resource.Attribute{
					Name:        "brkout_map",
					Description: `(Optional) breakout map for leaf breakout port group object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name alias for leaf breakout port group object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the leaf breakout port group object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for leaf breakout port group object.`,
				},
				resource.Attribute{
					Name:        "brkout_map",
					Description: `(Optional) breakout map for leaf breakout port group object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name alias for leaf breakout port group object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_interface_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Leaf Interface Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object leaf_interface_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Leaf Interface Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object leaf_interface_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object leaf_interface_profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Leaf Interface Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object leaf_interface_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object leaf_interface_profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Leaf Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object leaf_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Leaf Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object leaf_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object leaf_profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Leaf Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object leaf_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object leaf_profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Leaf Selector`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "leaf_profile_dn",
					Description: `(Required) Distinguished name of parent LeafProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object switch_association.`,
				},
				resource.Attribute{
					Name:        "switch_association_type",
					Description: `(Required) switch_association_type of Object switch_association. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Switch Association.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object switch_association.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object switch_association.`,
				},
				resource.Attribute{
					Name:        "switch_association_type",
					Description: `(Optional) leaf selector type`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Switch Association.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object switch_association.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object switch_association.`,
				},
				resource.Attribute{
					Name:        "switch_association_type",
					Description: `(Optional) leaf selector type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_lldp_interface_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI LLDP Interface Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object lldp_interface_policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LLDP Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_rx_st",
					Description: `(Optional) admin receive state.`,
				},
				resource.Attribute{
					Name:        "admin_tx_st",
					Description: `(Optional) admin transmit state.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object lldp_interface_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object lldp_interface_policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LLDP Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_rx_st",
					Description: `(Optional) admin receive state.`,
				},
				resource.Attribute{
					Name:        "admin_tx_st",
					Description: `(Optional) admin transmit state.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object lldp_interface_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object lldp_interface_policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_local_user",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Local User`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object local_user. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Local User.`,
				},
				resource.Attribute{
					Name:        "account_status",
					Description: `(Optional) local AAA user account status`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object local_user.`,
				},
				resource.Attribute{
					Name:        "cert_attribute",
					Description: `(Optional) cert_attribute for object local_user.`,
				},
				resource.Attribute{
					Name:        "clear_pwd_history",
					Description: `(Optional) clear password history of local user`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) email address of the local user`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `(Optional) local user account expiration date`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `(Optional) enables local user account expiration`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `(Optional) first name of the local user`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `(Optional) last name of the local user`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object local_user.`,
				},
				resource.Attribute{
					Name:        "otpenable",
					Description: `(Optional) otpenable for object local_user.`,
				},
				resource.Attribute{
					Name:        "otpkey",
					Description: `(Optional) otpkey for object local_user.`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `(Optional) phone number of the local user`,
				},
				resource.Attribute{
					Name:        "pwd",
					Description: `(Optional) system user password`,
				},
				resource.Attribute{
					Name:        "pwd_life_time",
					Description: `(Optional) lifetime of the local user password`,
				},
				resource.Attribute{
					Name:        "pwd_update_required",
					Description: `(Optional) pwd_update_required for object local_user.`,
				},
				resource.Attribute{
					Name:        "rbac_string",
					Description: `(Optional) rbac_string for object local_user.`,
				},
				resource.Attribute{
					Name:        "unix_user_id",
					Description: `(Optional) UNIX identifier of the local user`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Local User.`,
				},
				resource.Attribute{
					Name:        "account_status",
					Description: `(Optional) local AAA user account status`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object local_user.`,
				},
				resource.Attribute{
					Name:        "cert_attribute",
					Description: `(Optional) cert_attribute for object local_user.`,
				},
				resource.Attribute{
					Name:        "clear_pwd_history",
					Description: `(Optional) clear password history of local user`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) email address of the local user`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `(Optional) local user account expiration date`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `(Optional) enables local user account expiration`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `(Optional) first name of the local user`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `(Optional) last name of the local user`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object local_user.`,
				},
				resource.Attribute{
					Name:        "otpenable",
					Description: `(Optional) otpenable for object local_user.`,
				},
				resource.Attribute{
					Name:        "otpkey",
					Description: `(Optional) otpkey for object local_user.`,
				},
				resource.Attribute{
					Name:        "phone",
					Description: `(Optional) phone number of the local user`,
				},
				resource.Attribute{
					Name:        "pwd",
					Description: `(Optional) system user password`,
				},
				resource.Attribute{
					Name:        "pwd_life_time",
					Description: `(Optional) lifetime of the local user password`,
				},
				resource.Attribute{
					Name:        "pwd_update_required",
					Description: `(Optional) pwd_update_required for object local_user.`,
				},
				resource.Attribute{
					Name:        "rbac_string",
					Description: `(Optional) rbac_string for object local_user.`,
				},
				resource.Attribute{
					Name:        "unix_user_id",
					Description: `(Optional) UNIX identifier of the local user`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_logical_device_context",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Logical Device Context`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_logical_interface_context",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Logical Interface Context`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_logical_interface_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Logical Interface Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_node_profile_dn",
					Description: `(Required) Distinguished name of parent LogicalNodeProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object logical_interface_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Logical Interface Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object logical_interface_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object logical_interface_profile.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) qos priority class id`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) label color`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Logical Interface Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object logical_interface_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object logical_interface_profile.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) qos priority class id`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) label color`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_logical_node_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Logical Node Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "l3_outside_dn",
					Description: `(Required) Distinguished name of parent L3Outside object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object logical_node_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Logical Node Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object logical_node_profile.`,
				},
				resource.Attribute{
					Name:        "config_issues",
					Description: `(Optional) configuration issues`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object logical_node_profile.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) label color`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) target dscp`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Logical Node Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object logical_node_profile.`,
				},
				resource.Attribute{
					Name:        "config_issues",
					Description: `(Optional) configuration issues`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object logical_node_profile.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) label color`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) target dscp`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_logical_node_to_fabric_node",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Logical Node to Fabric Node`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_node_profile_dn",
					Description: `(Required) Distinguished name of parent LogicalNodeProfile object.`,
				},
				resource.Attribute{
					Name:        "tDn",
					Description: `(Required) tDn of Object fabric_node. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Fabric Node.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fabric_node.`,
				},
				resource.Attribute{
					Name:        "config_issues",
					Description: `(Optional) configuration issues`,
				},
				resource.Attribute{
					Name:        "rtr_id",
					Description: `(Optional) router identifier`,
				},
				resource.Attribute{
					Name:        "rtr_id_loop_back",
					Description: `(Optional)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Fabric Node.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object fabric_node.`,
				},
				resource.Attribute{
					Name:        "config_issues",
					Description: `(Optional) configuration issues`,
				},
				resource.Attribute{
					Name:        "rtr_id",
					Description: `(Optional) router identifier`,
				},
				resource.Attribute{
					Name:        "rtr_id_loop_back",
					Description: `(Optional)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_maintenance_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Maintenance Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object maintenance_policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Maintenance Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) maintenance policy admin state`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object maintenance_policy.`,
				},
				resource.Attribute{
					Name:        "graceful",
					Description: `(Optional) graceful for object maintenance_policy.`,
				},
				resource.Attribute{
					Name:        "ignore_compat",
					Description: `(Optional) whether compatibility check required`,
				},
				resource.Attribute{
					Name:        "internal_label",
					Description: `(Optional) firmware label`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object maintenance_policy.`,
				},
				resource.Attribute{
					Name:        "notif_cond",
					Description: `(Optional) when to send notifications to the admin`,
				},
				resource.Attribute{
					Name:        "run_mode",
					Description: `(Optional) maintenance policy run mode`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) compatibility catalog version`,
				},
				resource.Attribute{
					Name:        "version_check_override",
					Description: `(Optional) version check override`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Maintenance Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) maintenance policy admin state`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object maintenance_policy.`,
				},
				resource.Attribute{
					Name:        "graceful",
					Description: `(Optional) graceful for object maintenance_policy.`,
				},
				resource.Attribute{
					Name:        "ignore_compat",
					Description: `(Optional) whether compatibility check required`,
				},
				resource.Attribute{
					Name:        "internal_label",
					Description: `(Optional) firmware label`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object maintenance_policy.`,
				},
				resource.Attribute{
					Name:        "notif_cond",
					Description: `(Optional) when to send notifications to the admin`,
				},
				resource.Attribute{
					Name:        "run_mode",
					Description: `(Optional) maintenance policy run mode`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) compatibility catalog version`,
				},
				resource.Attribute{
					Name:        "version_check_override",
					Description: `(Optional) version check override`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_miscabling_protocol_interface_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Mis-cabling Protocol Interface Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object miscabling_protocol_interface_policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Mis-cabling Protocol Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) administrative state of the object or policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object miscabling_protocol_interface_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object miscabling_protocol_interface_policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Mis-cabling Protocol Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) administrative state of the object or policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object miscabling_protocol_interface_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object miscabling_protocol_interface_policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_monitoring_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Monitoring Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) tenant dn of object monitoring policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object monitoring policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the object monitoring policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object monitoring policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object monitoring policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the object monitoring policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object monitoring policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object monitoring policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_node_block",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Node Block`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "switch_association_dn",
					Description: `(Required) Distinguished name of parent SwitchAssociation object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object node_block. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Node Block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object node_block.`,
				},
				resource.Attribute{
					Name:        "from_",
					Description: `(Optional) from Node ID. Range from 101 to 110.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object node_block.`,
				},
				resource.Attribute{
					Name:        "to_",
					Description: `(Optional) to node ID. Range from 101 to 110.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Node Block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object node_block.`,
				},
				resource.Attribute{
					Name:        "from_",
					Description: `(Optional) from Node ID. Range from 101 to 110.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object node_block.`,
				},
				resource.Attribute{
					Name:        "to_",
					Description: `(Optional) to node ID. Range from 101 to 110.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_node_block_firmware",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Node Block`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "firmware_group_dn",
					Description: `(Required) Distinguished name of parent FirmwareGroup object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object node_block. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Node Block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object node_block.`,
				},
				resource.Attribute{
					Name:        "from_",
					Description: `(Optional) from`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object node_block.`,
				},
				resource.Attribute{
					Name:        "to_",
					Description: `(Optional) to`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Node Block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object node_block.`,
				},
				resource.Attribute{
					Name:        "from_",
					Description: `(Optional) from`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object node_block.`,
				},
				resource.Attribute{
					Name:        "to_",
					Description: `(Optional) to`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_node_mgmt_epg",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Node Management EPg`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ospf_interface_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI OSPF Interface Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object ospf_interface_policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the OSPF Interface Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object ospf_interface_policy.`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `(Optional) The OSPF cost for the interface. The cost (also called metric) of an interface in OSPF is an indication of the overhead required to send packets across a certain interface. The cost of an interface is inversely proportional to the bandwidth of that interface. A higher bandwidth indicates a lower cost. There is more overhead (higher cost) and time delays involved in crossing a 56k serial line than crossing a 10M ethernet line. The formula used to calculate the cost is: cost= 10000 0000/bandwidth in bps For example, it will cost 10 EXP8/10 EXP7 = 10 to cross a 10M Ethernet line and will cost 10 EXP8/1544000 = 64 to cross a T1 line. By default, the cost of an interface is calculated based on the bandwidth; you can force the cost of an interface with the ip ospf cost value interface sub-configuration mode command.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) interface policy controls`,
				},
				resource.Attribute{
					Name:        "dead_intvl",
					Description: `(Optional) The interval between hello packets from a neighbor before the router declares the neighbor as down. This value must be the same for all networking devices on a specific network. Specifying a smaller dead interval (seconds) will give faster detection of a neighbor being down and improve convergence, but might cause more routing instability.`,
				},
				resource.Attribute{
					Name:        "hello_intvl",
					Description: `(Optional) The interval between hello packets that OSPF sends on the interface. Note that the smaller the hello interval, the faster topological changes will be detected, but more routing traffic will ensue. This value must be the same for all routers and access servers on a specific network.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object ospf_interface_policy.`,
				},
				resource.Attribute{
					Name:        "nw_t",
					Description: `(Optional) The OSPF interface policy network type. OSPF supports point-to-point and broadcast.`,
				},
				resource.Attribute{
					Name:        "pfx_suppress",
					Description: `(Optional) pfx-suppression for object ospf_interface_policy.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The OSPF interface priority used to determine the designated router (DR) on a specific network. The router with the highest OSPF priority on a segment will become the DR for that segment. The same process is repeated for the backup designated router (BDR). In the case of a tie, the router with the highest RID will win. The default for the interface OSPF priority is one. Remember that the DR and BDR concepts are per multiaccess segment.`,
				},
				resource.Attribute{
					Name:        "rexmit_intvl",
					Description: `(Optional) The interval between LSA retransmissions. The retransmit interval occurs while the router is waiting for an acknowledgement from the neighbor router that it received the LSA. If no acknowlegment is received at the end of the interval, then the LSA is resent.`,
				},
				resource.Attribute{
					Name:        "xmit_delay",
					Description: `(Optional) The delay time needed to send an LSA update packet. OSPF increments the LSA age time by the transmit delay amount before transmitting the LSA update. You should take into account the transmission and propagation delays for the interface when you set this value.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the OSPF Interface Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object ospf_interface_policy.`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `(Optional) The OSPF cost for the interface. The cost (also called metric) of an interface in OSPF is an indication of the overhead required to send packets across a certain interface. The cost of an interface is inversely proportional to the bandwidth of that interface. A higher bandwidth indicates a lower cost. There is more overhead (higher cost) and time delays involved in crossing a 56k serial line than crossing a 10M ethernet line. The formula used to calculate the cost is: cost= 10000 0000/bandwidth in bps For example, it will cost 10 EXP8/10 EXP7 = 10 to cross a 10M Ethernet line and will cost 10 EXP8/1544000 = 64 to cross a T1 line. By default, the cost of an interface is calculated based on the bandwidth; you can force the cost of an interface with the ip ospf cost value interface sub-configuration mode command.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) interface policy controls`,
				},
				resource.Attribute{
					Name:        "dead_intvl",
					Description: `(Optional) The interval between hello packets from a neighbor before the router declares the neighbor as down. This value must be the same for all networking devices on a specific network. Specifying a smaller dead interval (seconds) will give faster detection of a neighbor being down and improve convergence, but might cause more routing instability.`,
				},
				resource.Attribute{
					Name:        "hello_intvl",
					Description: `(Optional) The interval between hello packets that OSPF sends on the interface. Note that the smaller the hello interval, the faster topological changes will be detected, but more routing traffic will ensue. This value must be the same for all routers and access servers on a specific network.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object ospf_interface_policy.`,
				},
				resource.Attribute{
					Name:        "nw_t",
					Description: `(Optional) The OSPF interface policy network type. OSPF supports point-to-point and broadcast.`,
				},
				resource.Attribute{
					Name:        "pfx_suppress",
					Description: `(Optional) pfx-suppression for object ospf_interface_policy.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The OSPF interface priority used to determine the designated router (DR) on a specific network. The router with the highest OSPF priority on a segment will become the DR for that segment. The same process is repeated for the backup designated router (BDR). In the case of a tie, the router with the highest RID will win. The default for the interface OSPF priority is one. Remember that the DR and BDR concepts are per multiaccess segment.`,
				},
				resource.Attribute{
					Name:        "rexmit_intvl",
					Description: `(Optional) The interval between LSA retransmissions. The retransmit interval occurs while the router is waiting for an acknowledgement from the neighbor router that it received the LSA. If no acknowlegment is received at the end of the interval, then the LSA is resent.`,
				},
				resource.Attribute{
					Name:        "xmit_delay",
					Description: `(Optional) The delay time needed to send an LSA update packet. OSPF increments the LSA age time by the transmit delay amount before transmitting the LSA update. You should take into account the transmission and propagation delays for the interface when you set this value.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_physical_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Physical Domain`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object physical_domain. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Physical Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object physical_domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object physical_domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Physical Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object physical_domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object physical_domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_port_security_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI POD Maintenance Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object pod_maintenance_group. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the POD Maintenance Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object pod_maintenance_group.`,
				},
				resource.Attribute{
					Name:        "fwtype",
					Description: `(Optional) fwtype for object pod_maintenance_group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object pod_maintenance_group.`,
				},
				resource.Attribute{
					Name:        "pod_maintenance_group_type",
					Description: `(Optional) component type`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the POD Maintenance Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object pod_maintenance_group.`,
				},
				resource.Attribute{
					Name:        "fwtype",
					Description: `(Optional) fwtype for object pod_maintenance_group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object pod_maintenance_group.`,
				},
				resource.Attribute{
					Name:        "pod_maintenance_group_type",
					Description: `(Optional) component type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ranges",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Ranges`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vlan_pool_dn",
					Description: `(Required) Distinguished name of parent VLANPool object.`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `(Required) from of Object ranges.`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `(Required) to of Object ranges. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Ranges.`,
				},
				resource.Attribute{
					Name:        "alloc_mode",
					Description: `(Optional) alloc_mode for object ranges.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object ranges.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object ranges.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) system role type`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Ranges.`,
				},
				resource.Attribute{
					Name:        "alloc_mode",
					Description: `(Optional) alloc_mode for object ranges.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object ranges.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object ranges.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) system role type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_service_redirect_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Service Redirect Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object service_redirect_policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Service Redirect Policy.`,
				},
				resource.Attribute{
					Name:        "anycast_enabled",
					Description: `(Optional) anycast_enabled for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "dest_type",
					Description: `(Optional) dest_type for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "hashing_algorithm",
					Description: `(Optional) hashing_algorithm for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "max_threshold_percent",
					Description: `(Optional) max_threshold_percent for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "min_threshold_percent",
					Description: `(Optional) min_threshold_percent for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "program_local_pod_only",
					Description: `(Optional) program_local_pod_only for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "resilient_hash_enabled",
					Description: `(Optional) resilient_hash_enabled for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "threshold_down_action",
					Description: `(Optional) threshold_down_action for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "threshold_enable",
					Description: `(Optional) threshold_enable for object service_redirect_policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Service Redirect Policy.`,
				},
				resource.Attribute{
					Name:        "anycast_enabled",
					Description: `(Optional) anycast_enabled for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "dest_type",
					Description: `(Optional) dest_type for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "hashing_algorithm",
					Description: `(Optional) hashing_algorithm for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "max_threshold_percent",
					Description: `(Optional) max_threshold_percent for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "min_threshold_percent",
					Description: `(Optional) min_threshold_percent for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "program_local_pod_only",
					Description: `(Optional) program_local_pod_only for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "resilient_hash_enabled",
					Description: `(Optional) resilient_hash_enabled for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "threshold_down_action",
					Description: `(Optional) threshold_down_action for object service_redirect_policy.`,
				},
				resource.Attribute{
					Name:        "threshold_enable",
					Description: `(Optional) threshold_enable for object service_redirect_policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_span_destination_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI SPAN Destination Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object span_destination_group. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SPAN Destination Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SPAN Destination Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_span_source_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI SPAN Source Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object span_source_group. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SPAN Source Group.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) administrative state of the object or policy`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SPAN Source Group.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) administrative state of the object or policy`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_span_sourcedestination_group_match_label",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI SPAN Source-destination Group Match Label`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "span_source_group_dn",
					Description: `(Required) Distinguished name of parent SPANSourceGroup object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object span_sourcedestination_group_match_label. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SPAN Source-destination Group Match Label.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) label color`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SPAN Source-destination Group Match Label.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) label color`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_interface_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Spine Interface Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object spine_interface_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Interface Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object spine_interface_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object spine_interface_profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Interface Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object spine_interface_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object spine_interface_profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_port_policy_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Spine Port Policy Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object aci_spine_port_policy_group. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Port Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object aci_spine_port_policy_group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object aci_spine_port_policy_group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Port Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object aci_spine_port_policy_group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object aci_spine_port_policy_group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_port_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Spine Port Selector`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "spine_profile_dn",
					Description: `(Required) Distinguished name of parent SpineProfile object.`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Required) tDn of Object interface_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the port selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object port selector.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the port selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object port selector.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Spine Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object spine_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object spine_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object spine_profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object spine_profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object spine_profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_switch_association",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Spine Switch Association`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "spine_profile_dn",
					Description: `(Required) Distinguished name of parent SpineProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object Spine Switch association.`,
				},
				resource.Attribute{
					Name:        "spine_switch_association_type",
					Description: `(Required) spine association type of Object Spine Switch association. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Switch Association.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object Spine Switch association.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name alias for object Spine Switch association.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Switch Association.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object Spine Switch association.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name alias for object Spine Switch association.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_subnet",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Subnet`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "parent_dn",
					Description: `(Required) Distinguished name of parent object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP address and mask of the default gateway. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Subnet.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object subnet.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) The list of subnet control state. The control can be specific protocols applied to the subnet such as IGMP Snooping.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object subnet.`,
				},
				resource.Attribute{
					Name:        "preferred",
					Description: `(Optional) Indicates if the subnet is preferred (primary) over the available alternatives. Only one preferred subnet is allowed.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The List of network visibility of the subnet.`,
				},
				resource.Attribute{
					Name:        "virtual",
					Description: `(Optional) Treated as virtual IP address. Used in case of BD extended to multiple sites.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Subnet.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object subnet.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) The list of subnet control state. The control can be specific protocols applied to the subnet such as IGMP Snooping.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object subnet.`,
				},
				resource.Attribute{
					Name:        "preferred",
					Description: `(Optional) Indicates if the subnet is preferred (primary) over the available alternatives. Only one preferred subnet is allowed.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The List of network visibility of the subnet.`,
				},
				resource.Attribute{
					Name:        "virtual",
					Description: `(Optional) Treated as virtual IP address. Used in case of BD extended to multiple sites.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_taboo_contract",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Taboo Contract`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object taboo_contract. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Taboo Contract.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object taboo_contract.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object taboo_contract.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Taboo Contract.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object taboo_contract.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object taboo_contract.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tenant",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Tenant`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object tenant. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Tenant.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object tenant.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object tenant.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Tenant.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object tenant.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object tenant.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_trigger_scheduler",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Trigger Scheduler`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object trigger_scheduler. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Trigger Scheduler.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object trigger_scheduler.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object trigger_scheduler.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Trigger Scheduler.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object trigger_scheduler.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object trigger_scheduler.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vlan_encapsulationfor_vxlan_traffic",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Vlan Encapsulation for Vxlan Traffic`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "attachable_access_entity_profile_dn",
					Description: `(Required) Distinguished name of parent AttachableAccessEntityProfile object. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Vlan Encapsulation for Vxlan Traffic.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vlan_encapsulationfor_vxlan_traffic.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vlan_encapsulationfor_vxlan_traffic.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Vlan Encapsulation for Vxlan Traffic.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vlan_encapsulationfor_vxlan_traffic.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vlan_encapsulationfor_vxlan_traffic.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vlan_pool",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI VLAN Pool`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object vlan_pool.`,
				},
				resource.Attribute{
					Name:        "allocMode",
					Description: `(Required) allocMode of Object vlan_pool. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VLAN Pool.`,
				},
				resource.Attribute{
					Name:        "alloc_mode",
					Description: `(Optional) allocation mode`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vlan_pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vlan_pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VLAN Pool.`,
				},
				resource.Attribute{
					Name:        "alloc_mode",
					Description: `(Optional) allocation mode`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vlan_pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vlan_pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vmm_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI VMM Domain`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "provider_profile_dn",
					Description: `(Required) Distinguished name of parent ProviderProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object vmm_domain. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VMM Domain.`,
				},
				resource.Attribute{
					Name:        "access_mode",
					Description: `(Optional) access_mode for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "arp_learning",
					Description: `(Optional) Enable/Disable arp learning for AVS Domain.`,
				},
				resource.Attribute{
					Name:        "ave_time_out",
					Description: `(Optional) ave_time_out for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "config_infra_pg",
					Description: `(Optional) Flag to enable config_infra_pg for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "ctrl_knob",
					Description: `(Optional) Type pf control knob to use.`,
				},
				resource.Attribute{
					Name:        "delimiter",
					Description: `(Optional) delimiter for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "enable_ave",
					Description: `(Optional) Flag to enable ave for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "enable_tag",
					Description: `(Optional) Flag enable tagging for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "encap_mode",
					Description: `(Optional) The layer 2 encapsulation protocol to use with the virtual switch.`,
				},
				resource.Attribute{
					Name:        "enf_pref",
					Description: `(Optional) The switching enforcement preference. This determines whether switches can be done within the virtual switch (Local Switching) or whether all switched traffic must go through the fabric (No Local Switching).`,
				},
				resource.Attribute{
					Name:        "ep_inventory_type",
					Description: `(Optional) Determines which end point inventory_type to use for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "ep_ret_time",
					Description: `(Optional) end point retention time for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "hv_avail_monitor",
					Description: `(Optional) Flag to enable hv_avail_monitor for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "mcast_addr",
					Description: `(Optional) The multicast address of the VMM domain profile.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The switch to be used for the domain profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "pref_encap_mode",
					Description: `(Optional) The preferred encapsulation mode for object vmm_domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VMM Domain.`,
				},
				resource.Attribute{
					Name:        "access_mode",
					Description: `(Optional) access_mode for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "arp_learning",
					Description: `(Optional) Enable/Disable arp learning for AVS Domain.`,
				},
				resource.Attribute{
					Name:        "ave_time_out",
					Description: `(Optional) ave_time_out for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "config_infra_pg",
					Description: `(Optional) Flag to enable config_infra_pg for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "ctrl_knob",
					Description: `(Optional) Type pf control knob to use.`,
				},
				resource.Attribute{
					Name:        "delimiter",
					Description: `(Optional) delimiter for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "enable_ave",
					Description: `(Optional) Flag to enable ave for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "enable_tag",
					Description: `(Optional) Flag enable tagging for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "encap_mode",
					Description: `(Optional) The layer 2 encapsulation protocol to use with the virtual switch.`,
				},
				resource.Attribute{
					Name:        "enf_pref",
					Description: `(Optional) The switching enforcement preference. This determines whether switches can be done within the virtual switch (Local Switching) or whether all switched traffic must go through the fabric (No Local Switching).`,
				},
				resource.Attribute{
					Name:        "ep_inventory_type",
					Description: `(Optional) Determines which end point inventory_type to use for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "ep_ret_time",
					Description: `(Optional) end point retention time for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "hv_avail_monitor",
					Description: `(Optional) Flag to enable hv_avail_monitor for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "mcast_addr",
					Description: `(Optional) The multicast address of the VMM domain profile.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The switch to be used for the domain profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vmm_domain.`,
				},
				resource.Attribute{
					Name:        "pref_encap_mode",
					Description: `(Optional) The preferred encapsulation mode for object vmm_domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vpc_explicit_protection_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI VPC Explicit Protection Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object vpc_explicit_protection_group. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VPC Explicit Protection Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vpc_explicit_protection_group.`,
				},
				resource.Attribute{
					Name:        "vpc_explicit_protection_group_id",
					Description: `(Optional) explicit protection group ID`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VPC Explicit Protection Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vpc_explicit_protection_group.`,
				},
				resource.Attribute{
					Name:        "vpc_explicit_protection_group_id",
					Description: `(Optional) explicit protection group ID`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vrf",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI VRF`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object vrf. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VRF.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation(tags) for object vrf.`,
				},
				resource.Attribute{
					Name:        "bd_enforced_enable",
					Description: `(Optional) Flag to enable/disable bd_enforced for VRF.`,
				},
				resource.Attribute{
					Name:        "ip_data_plane_learning",
					Description: `(Optional) iFlag to enable/disable ip-data-plane learning for VRF.`,
				},
				resource.Attribute{
					Name:        "knw_mcast_act",
					Description: `(Optional) specifies if known multicast traffic is forwarded.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vrf.`,
				},
				resource.Attribute{
					Name:        "pc_enf_dir",
					Description: `(Optional) Policy Control Enforcement Direction. It is used for defining policy enforcement direction for the traffic coming to or from an L3Out. Egress and Ingress directions are wrt L3Out. Default will be Ingress. But on the existing L3Outs during upgrade it will get set to Egress so that right after upgrade behavior doesn't change for them. This also means that there is no special upgrade sequence needed for upgrading to the release introducing this feature. After upgrade user would have to change the property value to Ingress. Once changed, system will reprogram the rules and prefix entry. Rules will get removed from the egress leaf and will get installed on the ingress leaf. Actrl prefix entry, if not already, will get installed on the ingress leaf. This feature will be ignored for the following cases: 1. Golf: Gets applied at Ingress by design. 2. Transit Rules get applied at Ingress by design. 4. vzAny 5. Taboo.`,
				},
				resource.Attribute{
					Name:        "pc_enf_pref",
					Description: `(Optional) Determines if the fabric should enforce contract policies to allow routing and packet forwarding.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VRF.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation(tags) for object vrf.`,
				},
				resource.Attribute{
					Name:        "bd_enforced_enable",
					Description: `(Optional) Flag to enable/disable bd_enforced for VRF.`,
				},
				resource.Attribute{
					Name:        "ip_data_plane_learning",
					Description: `(Optional) iFlag to enable/disable ip-data-plane learning for VRF.`,
				},
				resource.Attribute{
					Name:        "knw_mcast_act",
					Description: `(Optional) specifies if known multicast traffic is forwarded.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vrf.`,
				},
				resource.Attribute{
					Name:        "pc_enf_dir",
					Description: `(Optional) Policy Control Enforcement Direction. It is used for defining policy enforcement direction for the traffic coming to or from an L3Out. Egress and Ingress directions are wrt L3Out. Default will be Ingress. But on the existing L3Outs during upgrade it will get set to Egress so that right after upgrade behavior doesn't change for them. This also means that there is no special upgrade sequence needed for upgrading to the release introducing this feature. After upgrade user would have to change the property value to Ingress. Once changed, system will reprogram the rules and prefix entry. Rules will get removed from the egress leaf and will get installed on the ingress leaf. Actrl prefix entry, if not already, will get installed on the ingress leaf. This feature will be ignored for the following cases: 1. Golf: Gets applied at Ingress by design. 2. Transit Rules get applied at Ingress by design. 4. vzAny 5. Taboo.`,
				},
				resource.Attribute{
					Name:        "pc_enf_pref",
					Description: `(Optional) Determines if the fabric should enforce contract policies to allow routing and packet forwarding.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vsan_pool",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI VSAN Pool`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object vsan_pool.`,
				},
				resource.Attribute{
					Name:        "allocMode",
					Description: `(Required) allocMode of Object vsan_pool. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VSAN Pool.`,
				},
				resource.Attribute{
					Name:        "alloc_mode",
					Description: `(Optional) alloc_mode for object vsan_pool.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vsan_pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vsan_pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VSAN Pool.`,
				},
				resource.Attribute{
					Name:        "alloc_mode",
					Description: `(Optional) alloc_mode for object vsan_pool.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vsan_pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vsan_pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vxlan_pool",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI VXLAN Pool`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object vxlan_pool. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VXLAN Pool.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vxlan_pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vxlan_pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VXLAN Pool.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object vxlan_pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object vxlan_pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_x509_certificate",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI X509 Certificate`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "local_user_dn",
					Description: `(Required) Distinguished name of parent LocalUser object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object x509_certificate. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the X509 Certificate.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object x509_certificate.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Optional) data from the user certificate`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object x509_certificate.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the X509 Certificate.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object x509_certificate.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Optional) data from the user certificate`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object x509_certificate.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"aci_aaa_domain":                               0,
		"aci_access_generic":                           1,
		"aci_access_group":                             2,
		"aci_access_port_block":                        3,
		"aci_access_port_selector":                     4,
		"aci_access_sub_port_block":                    5,
		"aci_action_rule_profile":                      6,
		"aci_any":                                      7,
		"aci_application_epg":                          8,
		"aci_application_profile":                      9,
		"aci_attachable_access_entity_profile":         10,
		"aci_autonomous_system_profile":                11,
		"aci_bd_dhcp_label":                            12,
		"aci_bridge_domain":                            13,
		"aci_cdp_interface_policy":                     14,
		"aci_client_end_point":                         15,
		"aci_cloud_applicationcontainer":               16,
		"aci_cloud_availability_zone":                  17,
		"aci_cloud_aws_provider":                       18,
		"aci_cloud_cidr_pool":                          19,
		"aci_cloud_context_profile":                    20,
		"aci_cloud_domain_profile":                     21,
		"aci_cloud_endpoint_selector":                  22,
		"aci_cloud_endpoint_selectorfor_external_epgs": 23,
		"aci_cloud_epg":                                24,
		"aci_cloud_external_epg":                       25,
		"aci_cloud_provider_profile":                   26,
		"aci_cloud_providers_region":                   27,
		"aci_cloud_subnet":                             28,
		"aci_cloud_vpn_gateway":                        29,
		"aci_configuration_export_policy":              30,
		"aci_configuration_import_policy":              31,
		"aci_connection":                               32,
		"aci_contract":                                 33,
		"aci_contract_subject":                         34,
		"aci_destination_of_redirected_traffic":        35,
		"aci_dhcp_option":                              36,
		"aci_dhcp_option_policy":                       37,
		"aci_dhcp_relay_policy":                        38,
		"aci_end_point_retention_policy":               39,
		"aci_epg_to_contract":                          40,
		"aci_epg_to_domain":                            41,
		"aci_epg_to_static_path":                       42,
		"aci_epgs_using_function":                      43,
		"aci_external_network_instance_profile":        44,
		"aci_fabric_if_pol":                            45,
		"aci_fabric_node_member":                       46,
		"aci_fabric_path_ep":                           47,
		"aci_fc_domain":                                48,
		"aci_fex_bundle_group":                         49,
		"aci_fex_profile":                              50,
		"aci_filter":                                   51,
		"aci_filter_entry":                             52,
		"aci_firmware_download_task":                   53,
		"aci_firmware_group":                           54,
		"aci_firmware_policy":                          55,
		"aci_function_node":                            56,
		"aci_imported_contract":                        57,
		"aci_interface_fc_policy":                      58,
		"aci_l2_domain":                                59,
		"aci_l2_interface_policy":                      60,
		"aci_l2_outside":                               61,
		"aci_l2out_extepg":                             62,
		"aci_l3_domain_profile":                        63,
		"aci_l3_ext_subnet":                            64,
		"aci_l3_outside":                               65,
		"aci_l4_l7_service_graph_template":             66,
		"aci_lacp_policy":                              67,
		"aci_leaf_access_bundle_policy_group":          68,
		"aci_leaf_access_port_policy_group":            69,
		"aci_leaf_breakout_port_group":                 70,
		"aci_leaf_interface_profile":                   71,
		"aci_leaf_profile":                             72,
		"aci_leaf_selector":                            73,
		"aci_lldp_interface_policy":                    74,
		"aci_local_user":                               75,
		"aci_logical_device_context":                   76,
		"aci_logical_interface_context":                77,
		"aci_logical_interface_profile":                78,
		"aci_logical_node_profile":                     79,
		"aci_logical_node_to_fabric_node":              80,
		"aci_maintenance_policy":                       81,
		"aci_miscabling_protocol_interface_policy":     82,
		"aci_monitoring_policy":                        83,
		"aci_node_block":                               84,
		"aci_node_block_firmware":                      85,
		"aci_node_mgmt_epg":                            86,
		"aci_ospf_interface_policy":                    87,
		"aci_physical_domain":                          88,
		"aci_port_security_policy":                     89,
		"aci_ranges":                                   90,
		"aci_service_redirect_policy":                  91,
		"aci_span_destination_group":                   92,
		"aci_span_source_group":                        93,
		"aci_span_sourcedestination_group_match_label": 94,
		"aci_spine_interface_profile":                  95,
		"aci_spine_port_policy_group":                  96,
		"aci_spine_port_selector":                      97,
		"aci_spine_profile":                            98,
		"aci_spine_switch_association":                 99,
		"aci_subnet":                                   100,
		"aci_taboo_contract":                           101,
		"aci_tenant":                                   102,
		"aci_trigger_scheduler":                        103,
		"aci_vlan_encapsulationfor_vxlan_traffic":      104,
		"aci_vlan_pool":                                105,
		"aci_vmm_domain":                               106,
		"aci_vpc_explicit_protection_group":            107,
		"aci_vrf":                                      108,
		"aci_vsan_pool":                                109,
		"aci_vxlan_pool":                               110,
		"aci_x509_certificate":                         111,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
