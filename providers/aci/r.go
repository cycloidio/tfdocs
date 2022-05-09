package aci

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "aci_aaa_domain",
			Category:         "AAA",
			ShortDescription: `Manages ACI aaa Domain`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object aaa domain.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object aaa domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object aaa domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object aaa domain. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the aaa domain. ## Importing ## An existing aaa domain can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: <https://www.terraform.io/docs/import/index.html> ` + "`" + `` + "`" + `` + "`" + `bash terraform import aci_aaa_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_aaa_domain_relationship",
			Category:         "AAA",
			ShortDescription: `Manages ACI AAA Domain Relationship for Parent Object`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"domain",
				"relationship",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "parent_dn",
					Description: `(Required) Distinguished name of parent object.`,
				},
				resource.Attribute{
					Name:        "aaa_domain_dn",
					Description: `(Required) Distinguished name of the AAA Security Domain for Parent Object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the AAA Security Domain for Parent Object. ## Importing ## An existing AAA Security Domain Relationship object can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_aaa_domain_relationship.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_aaep_to_domain",
			Category:         "Access Policies",
			ShortDescription: `Manages the ACI Attachable Access Entity Profile (AAEP) to domain (VMM, Physical or External domain) relationship.`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"aaep",
				"to",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "attachable_access_entity_profile_dn",
					Description: `(Required) Distinguished name of the parent Attachable Access Entity Profile object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Attachable AccessEntity Profile to Domain Relationship object.`,
				},
				resource.Attribute{
					Name:        "domain_dn",
					Description: `(Required) The Distinguished name of the domain object. ## Importing ## An existing Attachable Access Entity Profile to Domain Relationship object can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_aaep_to_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_generic",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Access Generic`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"generic",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_group",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Access Group`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_port_block",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Access Port Block`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"port",
				"block",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_port_selector",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Access Port Selector`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"port",
				"selector",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_sub_port_block",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Access Sub Port Block`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"sub",
				"port",
				"block",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_port_selector_dn",
					Description: `(Required) Distinguished name of parent Access Port Selector or Spine Access Port Selector object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object access sub port block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object access sub port block.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object access sub port block.`,
				},
				resource.Attribute{
					Name:        "from_card",
					Description: `(Optional) From card. Allowed Values are between 1 to 100. Default Value is "1".`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `(Optional) Port block from port Allowed Values are between 1 to 127. Default Value is "1".`,
				},
				resource.Attribute{
					Name:        "from_sub_port",
					Description: `(Optional) From sub port for object access sub port block. Allowed Values are between 1 to 64. Default Value is "1".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object access sub port block.`,
				},
				resource.Attribute{
					Name:        "to_card",
					Description: `(Optional) To card. Allowed Values are between 1 to 100. Default Value is "1".`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `(Optional) To port. Allowed Values are between 1 to 127. Default Value is "1".`,
				},
				resource.Attribute{
					Name:        "to_sub_port",
					Description: `(Optional) To sub port for object access sub port block. Allowed Values are between 1 to 64. Default Value is "1". ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Access Sub Port Block. ## Importing ## An existing Access Sub Port Block can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_access_sub_port_block.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_switch_policy_group",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Access Switch Policy Group`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"switch",
				"policy",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object Access Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Access Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_bfd_ipv4_inst_pol",
					Description: `(Optional) Represents the relation to a BFD Ipv4 Instance Policy (class bfdIpv4InstPol). Relationship to BFD Ipv4 Instance Policy Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_bfd_ipv6_inst_pol",
					Description: `(Optional) Represents the relation to a BFD Ipv6 Instance Policy (class bfdIpv6InstPol). Relationship to BFD Ipv6 Instance Policy Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_bfd_mh_ipv4_inst_pol",
					Description: `(Optional) Represents the relation to a MH BFD Ipv4 Instance Policy (class bfdMhIpv4InstPol). Relationship to MH BFD Ipv4 Instance Policy Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_bfd_mh_ipv6_inst_pol",
					Description: `(Optional) Represents the relation to a MH BFD Ipv6 Instance Policy (class bfdMhIpv6InstPol). Relationship to MH BFD Ipv6 Instance Policy Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_equipment_flash_config_pol",
					Description: `(Optional) Represents the relation to a Flash Configuration Policy (class equipmentFlashConfigPol). Relation to equipmentFlashConfigInstPol Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_fc_fabric_pol",
					Description: `(Optional) Represents the relation to a Fibre Channel Fabric Level Policy (class fcFabricPol). Relation to fcInstPol Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_fc_inst_pol",
					Description: `(Optional) Represents the relation to a Fibre Channel Instance Policy (class fcInstPol). Relation to fcInstPol Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_iacl_leaf_profile",
					Description: `(Optional) Represents the relation to a CoPP Prefilter Profile for Leafs (class iaclLeafProfile). Relationship the CoPP Prefilter Leaf profile to be applied on leafs Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_l2_node_auth_pol",
					Description: `(Optional) Represents the relation to a Node Authentication (802.1x) policy (class l2NodeAuthPol). Relation to l2NodeAuthPol Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_leaf_copp_profile",
					Description: `(Optional) Represents the relation to a CoPP Profile for Leafs (class coppLeafProfile). Relationship the CoPP profile to be applied on leafs Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_leaf_p_grp_to_cdp_if_pol",
					Description: `(Optional) Represents the relation to a Relation to cdp interface policy for mgmt port (class cdpIfPol). Relationship to cdp interface policy for mgmt port Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_leaf_p_grp_to_lldp_if_pol",
					Description: `(Optional) Represents the relation to a Relation to lldp interface policy for mgmt port (class lldpIfPol). Relationship to lldp interface policy for mgmt port Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_mon_node_infra_pol",
					Description: `(Optional) Represents the relation to a Relation to Access Monitoring Policy (class monInfraPol). A source relation to the monitoring policy model. Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_mst_inst_pol",
					Description: `(Optional) Represents the relation to a MST Instance Policy (class stpInstPol). A source relation to a spanning tree protocol policy. Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_netflow_node_pol",
					Description: `(Optional) Represents the relation to a Netflow Node Policy (class netflowNodePol). Relationship to Netflow Node Policy Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_poe_inst_pol",
					Description: `(Optional) Represents the relation to a POE Node Policy (class poeInstPol). Relationship to POE Node Policy Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_topoctrl_fast_link_failover_inst_pol",
					Description: `(Optional) Represents the relation to a Fast Link Failover Instance Policy (class topoctrlFastLinkFailoverInstPol). Relation to topoctrlFastLinkFailoverPol Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_topoctrl_fwd_scale_prof_pol",
					Description: `(Optional) Represents the relation to a Forwarding Scale Profile Policy (class topoctrlFwdScaleProfilePol). Relation to topoctrlFwdScaleProfilePol Type: String.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Access Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Access Switch Policy Group. ## Importing ## An existing Access Switch Policy Group can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_access_switch_policy_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_action_rule_additional_communities",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI Action Rule Profile Set Additional Communities`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"action",
				"rule",
				"additional",
				"communities",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action_rule_profile_dn",
					Description: `(Required) Distinguished name of the parent action rule profile object.`,
				},
				resource.Attribute{
					Name:        "community",
					Description: `(Required) The community value of the set action rule profile additional communities object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the set action rule profile additional communities object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the action rule profile additional communities object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the additional communities object.`,
				},
				resource.Attribute{
					Name:        "set_criteria",
					Description: `(Optional) The criteria for setting the (extended) community attribute for a BGP route update. Allowed values are "append", "none", "replace", and default value is "append". Type: String. ## Importing ## An existing Action Rule Profile Additional Communities can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_action_rule_additional_communities.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_action_rule_profile",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI Action Rule Profile`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"action",
				"rule",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of the parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_route_tag",
					Description: `(Optional) Set Route Tag of the Action Rule Profile object. Can not be configured along with ` + "`" + `next_hop_propagation` + "`" + ` and ` + "`" + `multipath` + "`" + `. Type: Integer.`,
				},
				resource.Attribute{
					Name:        "set_preference",
					Description: `(Optional) Set Preference of the Action Rule Profile object. Type: Integer.`,
				},
				resource.Attribute{
					Name:        "set_weight",
					Description: `(Optional) Set Weight of the Action Rule Profile object. Type: Integer.`,
				},
				resource.Attribute{
					Name:        "set_metric",
					Description: `(Optional) Set Metric of the Action Rule Profile object. Type: Integer.`,
				},
				resource.Attribute{
					Name:        "set_metric_type",
					Description: `(Optional) Set Metric Type of the Action Rule Profile object. Allowed values are ` + "`" + `ospf-type1` + "`" + `, ` + "`" + `ospf-type2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "set_next_hop",
					Description: `(Optional) Set Next Hop of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_communities",
					Description: `(Optional) A block representing the attributes of Set Communities object. Type: Block.`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `(Optional) Criteria of the Set Communities object. Allowed values are ` + "`" + `append` + "`" + ` or ` + "`" + `replace` + "`" + `. Type: String.`,
				},
				resource.Attribute{
					Name:        "community",
					Description: `(Optional) Community of the Set Communities object. Allowed input formats are ` + "`" + `regular:as2-nn2:4:15` + "`" + `, ` + "`" + `extended:as4-nn2:5:16` + "`" + `, ` + "`" + `no-export` + "`" + ` and ` + "`" + `no-advertise` + "`" + `. Type: String.`,
				},
				resource.Attribute{
					Name:        "next_hop_propagation",
					Description: `(Optional) Next Hop Propagation of the Action Rule Profile object. Allowed values are ` + "`" + `yes` + "`" + ` or ` + "`" + `no` + "`" + `. Can not be configured along with ` + "`" + `set_route_tag` + "`" + `. Type: String.`,
				},
				resource.Attribute{
					Name:        "multipath",
					Description: `(Optional) Multipath of the Action Rule Profile object. Allowed values are ` + "`" + `yes` + "`" + ` or ` + "`" + `no` + "`" + `. Can not be configured along with ` + "`" + `set_route_tag` + "`" + `. Type: String.`,
				},
				resource.Attribute{
					Name:        "set_as_path_prepend_last_as",
					Description: `(Optional) Number of ASN to be prepended to AS Path of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_as_path_prepend_as",
					Description: `(Optional) A block representing ASNs to be configured as Set As Path - Prepend AS of the Action Rule Profile object. Type: Block.`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `ASN to be prepended to Set AS Path.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Order in which the ASN should be prepended to Set AS Path. ## Importing ## An existing Action Rule Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_action_rule_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_annotation",
			Category:         "Resources",
			ShortDescription: `Manages ACI Annotation`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "parent_dn",
					Description: `(Required) Distinguished name of the object to which the Annotation is attached to.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the Annotation.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the Annotation. ## Importing ## An existing Annotation can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_annotation.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_any",
			Category:         "Networking",
			ShortDescription: `Manages ACI Any`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"any",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_application_epg",
			Category:         "Application Management",
			ShortDescription: `Manages ACI Application EPG`,
			Description:      ``,
			Keywords: []string{
				"application",
				"management",
				"epg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_profile_dn",
					Description: `(Required) Distinguished Name of the parent application profile.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object application epg.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object application epg.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object application epg.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) Exception tag for object application epg. Range: "0" - "512" .`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings. Allowed values are "disabled" and "enabled". Default is "disabled".`,
				},
				resource.Attribute{
					Name:        "fwd_ctrl",
					Description: `(Optional) Forwarding control at EPG level. Allowed values are "none" and "proxy-arp". Default is "none".`,
				},
				resource.Attribute{
					Name:        "has_mcast_source",
					Description: `(Optional) If the source for the EPG is multicast or not. Allowed values are "yes" and "no". Default values is "no".`,
				},
				resource.Attribute{
					Name:        "is_attr_based_epg",
					Description: `(Optional) If the EPG is attribute based or not. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria for EPG. Allowed values are "All", "AtleastOne", "AtmostOne", "None". Default is "AtleastOne".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object application epg.`,
				},
				resource.Attribute{
					Name:        "pc_enf_pref",
					Description: `(Optional) The preferred policy control. Allowed values are "unenforced" and "enforced". Default is "unenforced".`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication. Allowed values are "exclude" and "include". Default is "exclude".`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) QoS priority class id. Allowed values are "unspecified", "level1", "level2", "level3", "level4","level5" and "level6". By default the value is inherited from the parent application profile.`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `(Optional) Shutdown for object application epg. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd",
					Description: `(Optional) Relation to the Bridge domain associated with EPG (Point to class fvBD). This attribute is optional because the ACI API does not mandate it`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cust_qos_pol",
					Description: `(Optional) Relation to Custom Quality of Service traffic policy name (Point to class qosCustomPol). Cardinality - N_TO_ONE. Type - String. <!-- tenant -> policies -> protocol -> Custom QoS -->`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_fc_path_att",
					Description: `(Optional) Relation to Fibre Channel (Paths) (Point to class fabricPathEp). Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prov",
					Description: `(Optional) Relation to Provided Contract (Point to class vzBrCP). Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons_if",
					Description: `(Optional) Relation to Imported Contract (Point to class vzCPIf). Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_sec_inherited",
					Description: `(Optional) Relation represents that the EPG is inheriting security configuration from other EPGs (Point to class fvEPg). Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_node_att",
					Description: `(Optional) Relation used to define a Static Leaf binding (Point to class fabricNode). Cardinality - N_TO_M. Type - Set of String. <!-- tenant -> Application Profile -> EPG ->Static Leaf -->`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_dpp_pol",
					Description: `(Optional) Relation to define a Data Plane Policing policy (Point to class qosDppPol). Cardinality - N_TO_ONE. Type - String. <!-- tenant -> policies -> protocol -> Data Plane Policing -->`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons",
					Description: `(Optional) Relation to Consumed Contract (Point to class vzBrCP). Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_trust_ctrl",
					Description: `(Optional) Relation to First Hop Security trust control (Point to class fhsTrustCtrlPol). Cardinality - N_TO_ONE. Type - String. <!-- tenant -> policies -> protocol -> First Hop Security -->`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prot_by",
					Description: `(Optional) Relation to Taboo Contract (Point to class vzTaboo). Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_aepg_mon_pol",
					Description: `(Optional) Relation to create a container for monitoring policies associated with the tenant. This allows you to apply tenant-specific policies (Point to class monEPGPol). Cardinality - N_TO_ONE. Type - String. <!-- tenant -> policies -> Monitoring -->`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_intra_epg",
					Description: `(Optional) Relation to Intra EPG Contract (Point to class vzBrCP). Cardinality - N_TO_M. Type - Set of String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Application EPG. ## Importing An existing Application EPG can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_application_epg.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_application_profile",
			Category:         "Application Management",
			ShortDescription: `Manages ACI Application Profile`,
			Description:      ``,
			Keywords: []string{
				"application",
				"management",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_attachable_access_entity_profile",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Attachable Access Entity Profile`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"attachable",
				"entity",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object attachable access entity profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object attachable access entity profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object attachable access entity profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object attachable access entity profile.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_dom_p",
					Description: `(Optional) Relation to class infraADomP. Cardinality - N_TO_M. Type - [Set of String]. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Attachable Access Entity Profile. ## Importing ## An existing Attachable Access Entity Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: <https://www.terraform.io/docs/import/index.html> ` + "`" + `` + "`" + `` + "`" + `bash terraform import aci_attachable_access_entity_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_authentication_properties",
			Category:         "AAA",
			ShortDescription: `Manages ACI AAA Authentication Properties and Default Radius Authentication Settings`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"authentication",
				"properties",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of objects AAA Authentication Properties and Default Radius Authentication Settings.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of objects AAA Authentication Properties and Default Radius Authentication Settings.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of objects AAA Authentication Properties and Default Radius Authentication Settings.`,
				},
				resource.Attribute{
					Name:        "def_role_policy",
					Description: `(Optional) The default role policy of remote user. Allowed values are "assign-default-role" and "no-login".`,
				},
				resource.Attribute{
					Name:        "ping_check",
					Description: `(Optional) Heart bit ping checks for RADIUS/TACACS/LDAP/SAML/RSA server reachability. Allowed values are "false" and "true".`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) The number of attempts that the authentication method is tried. Allowed range: "0" - "5".`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The amount of time between authentication attempts. Allowed range: "1" - "60". ## Importing ## An existing AAA Authentication Properties and Default Radius Authentication Settings can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_authentication_properties.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bd_dhcp_label",
			Category:         "Networking",
			ShortDescription: `Manages ACI BD DHCP Label`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"bd",
				"dhcp",
				"label",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bfd_interface_policy",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI BFD Interface Policy`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"bfd",
				"interface",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Administrative state of the object BFD Interface Policy. Allowed values are "disabled" and "enabled". Default is "enabled".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) Control state for object BFD Interface Policy. Allowed values are "opt-subif" and "none". Default is "none".`,
				},
				resource.Attribute{
					Name:        "detect_mult",
					Description: `(Optional) Detection multiplier for object BFD Interface Policy. Range: "1" - "50". Default value is "3".`,
				},
				resource.Attribute{
					Name:        "echo_admin_st",
					Description: `(Optional) Echo mode indicator for object BFD Interface Policy. Allowed values are "disabled" and "enabled". Default is "enabled".`,
				},
				resource.Attribute{
					Name:        "echo_rx_intvl",
					Description: `(Optional) Echo rx interval for object BFD Interface Policy. Range: "50" - "999". Default value is "50".`,
				},
				resource.Attribute{
					Name:        "min_rx_intvl",
					Description: `(Optional) Required minimum rx interval for boject BFD Interface Policy. Range: "50" - "999". Default value is "50".`,
				},
				resource.Attribute{
					Name:        "min_tx_intvl",
					Description: `(Optional) Desired minimum tx interval for object BFD Interface Policy. Range: "50" - "999". Default value is "50".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object BFD Interface Policy. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the BFD Interface Policy. ## Importing ## An existing BFD Interface Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_bfd_interface_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bgp_address_family_context",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI BGP Address Family Context`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"bgp",
				"address",
				"family",
				"context",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bgp_best_path_policy",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI BGP Best Path Policy`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"bgp",
				"best",
				"path",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bgp_peer_connectivity_profile",
			Category:         "L3Out",
			ShortDescription: `Manages ACI BGP Peer Connectivity Profile`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"bgp",
				"peer",
				"connectivity",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) The connector direction. Allowed values are "export", "import", and default value is "import". Type: String.`,
				},
				resource.Attribute{
					Name:        "target_dn",
					Description: `(Required) The distinguished name of the target Route Map for Route Control Profile. Type: String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the BGP Peer Connectivity Profile. ## Importing An existing BGP Peer Connectivity Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_bgp_peer_connectivity_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bgp_peer_prefix",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI BGP Peer Prefix`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"bgp",
				"peer",
				"prefix",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action when the maximum prefix limit is reached for BGP peer prefix object. Allowed values are "log", "reject", "restart" and "shut". Default value is "reject".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "max_pfx",
					Description: `(Optional) Maximum number of prefixes allowed from the peer for BGP peer prefix object. Default value is "20000".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "restart_time",
					Description: `(Optional) The period of time in minutes before restarting the peer when the prefix limit is reached for BGP peer prefix object. Default value is "infinite".`,
				},
				resource.Attribute{
					Name:        "thresh",
					Description: `(Optional) Threshold percentage of the maximum number of prefixes before a warning is issued for BGP peer prefix object. Default value is "75". ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the BGP Peer Prefix. ## Importing ## An existing BGP Peer Prefix can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_bgp_peer_prefix.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bgp_route_control_profile",
			Category:         "L3Out",
			ShortDescription: `Manages ACI BGP Route Control Profile`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"bgp",
				"route",
				"control",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bgp_route_summarization",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI BGP Route Summarization`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"bgp",
				"route",
				"summarization",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bgp_timers",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI BGP Timers`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"bgp",
				"timers",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bridge_domain",
			Category:         "Networking",
			ShortDescription: `Manages ACI Bridge Domain`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"bridge",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "relation_fv_rs_bd_to_profile",
					Description: `(Optional) Relation to L3Outs Route Map For Import and Export Route Control (Point to class rtctrlProfile). Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_mldsn",
					Description: `(Optional) Relation to MLD Snoop (Point to class mldSnoopPol). Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_abd_pol_mon_pol",
					Description: `(Optional) Relation to create a container for monitoring policies associated with the tenant (Point to class monEPGPol). Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd_to_nd_p",
					Description: `(Optional) Relation to ND Policy (Point to class ndIfPol). Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd_flood_to",
					Description: `(Optional) Relation to Contract Filters (Point to class vzFilter). Cardinality - N_TO_M. Type - Set of String. <!-- Tenants -> Contracts -> Filters -->`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd_to_fhs",
					Description: `(Optional) Relation First Hop Security Feature Policies (Point to class fhsBDPol). Requires unicast_route to be set to "yes". Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_bd_to_relay_p",
					Description: `(Optional) Relation to DHCP Relay policy (Point to class dhcpRelayP). Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_ctx",
					Description: `(Optional) Relation to VRF (Point to class fvCtx). Cardinality - N_TO_ONE. Type - String. Note: In the APIC GUI,a VRF (fvCtx) was called a "Context"or "PrivateNetwork." - ` + "`" + `relation_fv_rs_bd_to_netflow_monitor_pol` + "`" + ` - (Optional) Relation to Netflow Monitors policy (Point to class netflowMonitorPol). Cardinality - N_TO_M. Type - Set of Map. - ` + "`" + `relation_fv_rs_igmpsn` + "`" + ` - (Optional) Relation to IGMP Snoop policy (Point to class igmpSnoopPol). Cardinality - N_TO_ONE. Type - String. - ` + "`" + `relation_fv_rs_bd_to_ep_ret` + "`" + ` - (Optional) Relation to End Point Retention policy (Point to class fvEpRetPol). Cardinality - N_TO_ONE. Type - String. - ` + "`" + `relation_fv_rs_bd_to_out` + "`" + ` - (Optional) Relation to L3Outs (Point to class l3extOut). Cardinality - N_TO_M. Type - Set of String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Bridge Domain. ## Importing An existing Bridge Domain can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_bridge_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cdp_interface_policy",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI CDP Interface Policy`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"cdp",
				"interface",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object cdp interface policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Administrative state. Allowed values: "enabled", "disabled". Default value is "enabled".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object cdp interface policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object cdp interface policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object cdp interface policy. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the CDP Interface Policy. ## Importing ## An existing CDP Interface Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: <https://www.terraform.io/docs/import/index.html> ` + "`" + `` + "`" + `` + "`" + `bash terraform import aci_cdp_interface_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_applicationcontainer",
			Category:         "Cloud",
			ShortDescription: `Manages ACI Cloud Application container`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"applicationcontainer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object cloud applicationcontainer.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object cloud applicationcontainer.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object cloud applicationcontainer.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object cloud applicationcontainer. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud Application container. ## Importing ## An existing Cloud Application container can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_applicationcontainer.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_aws_provider",
			Category:         "Cloud",
			ShortDescription: `Manages ACI Cloud AWS Provider`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"aws",
				"provider",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_cidr_pool",
			Category:         "Cloud",
			ShortDescription: `Manages ACI Cloud CIDR Pool`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"cidr",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_context_profile_dn",
					Description: `(Required) Distinguished name of parent CloudContextProfile object.`,
				},
				resource.Attribute{
					Name:        "addr",
					Description: `(Required) CIDR IPv4 block.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Cloud CIDR Pool.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Cloud CIDR Pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Cloud CIDR Pool.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Optional) Flag to specify whether CIDR is primary CIDR or not. Allowed values are "yes" and "no". Default is "no". Only one primary CIDR is supported under a cloud context profile. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud CIDR Pool. ## Importing ## An existing Cloud CIDR Pool can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_cidr_pool.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_context_profile",
			Category:         "Cloud",
			ShortDescription: `Manages ACI Cloud Context Profile`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"context",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object Cloud Context profile.`,
				},
				resource.Attribute{
					Name:        "primary_cidr",
					Description: `(Required) Primary CIDR block of Cloud Context profile.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) AWS region in which profile is created.`,
				},
				resource.Attribute{
					Name:        "cloud_vendor",
					Description: `(Required) Name of the vendor. e.g. "aws", "azure".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Cloud Context profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Cloud Context profile.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The specific type of the object or component. Allowed values are "regular", "shadow", "hosted" and "container-overlay". Default is "regular".`,
				},
				resource.Attribute{
					Name:        "hub_network",
					Description: `(Optional) Hub network Dn which enables Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_ctx_to_flow_log",
					Description: `(Optional) Relation to class cloudAwsFlowLogPol. Cardinality - N TO ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_to_ctx",
					Description: `(Required) Relation to class fvCtx. Cardinality - N TO ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_ctx_profile_to_region",
					Description: `(Optional) Relation to class cloudRegion. Cardinality - N TO ONE. Type - String.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_domain_profile",
			Category:         "Cloud",
			ShortDescription: `Manages ACI Cloud Domain Profile`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"domain",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_endpoint_selector",
			Category:         "Cloud",
			ShortDescription: `Manages ACI Cloud Endpoint Selector`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"endpoint",
				"selector",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_endpoint_selectorfor_external_epgs",
			Category:         "Cloud",
			ShortDescription: `Manages ACI Cloud Endpoint Selector for External EPgs`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"endpoint",
				"selectorfor",
				"external",
				"epgs",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_epg",
			Category:         "Cloud",
			ShortDescription: `Manages ACI Cloud EPg`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"epg",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_external_epg",
			Category:         "Cloud",
			ShortDescription: `Manages ACI Cloud External EPg`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"external",
				"epg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_applicationcontainer_dn",
					Description: `(Required) Distinguished name of parent CloudApplicationcontainer object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object Cloud External EPg.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Cloud External EPg.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Cloud External EPg.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) Exception tag for object Cloud External EPg. Allowed value range is "0" to "512".`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings. Allowed values are "disabled" and "enabled". Default is "disabled".`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria. Allowed values are "All", "AtleastOne", "AtmostOne" and "None". Default values is "AtleastOne".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) name_alias for object Cloud External EPg.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication. Allowed values are "include" and "exclude". Default is "exclude".`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) QOS priority class id. Allowed values are "unspecified", "level1", "level2", "level3", "level4", "level5" and "level6". Default is "unspecified.`,
				},
				resource.Attribute{
					Name:        "route_reachability",
					Description: `(Optional) Route reachability for this EPG. Allowed values are "unspecified", "inter-site", "internet", "site-ext" and "inter-site-ext".Default is "inter-site".`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_sec_inherited",
					Description: `(Optional) Relation to class fvEPg. Cardinality - N TO M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prov",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N TO M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons_if",
					Description: `(Optional) Relation to class vzCPIf. Cardinality - N TO M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cust_qos_pol",
					Description: `(Optional) Relation to class qosCustomPol. Cardinality - N TO ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N TO M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_cloud_epg_ctx",
					Description: `(Optional) Relation to class fvCtx. Cardinality - N TO ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prot_by",
					Description: `(Optional) Relation to class vzTaboo. Cardinality - N TO M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_intra_epg",
					Description: `(Optional) Relation to class vzBrCP. Cardinality - N TO M. Type - Set of String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud External EPg. ## Importing ## An existing Cloud External EPg can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_external_epg.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_provider_profile",
			Category:         "Cloud",
			ShortDescription: `Manages ACI Cloud Provider Profile`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"provider",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vendor",
					Description: `(Required) vendor of Object cloud_provider_profile. Currently only supported vendor is "aws".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object cloud_provider_profile. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud Provider Profile. ## Importing ## An existing Cloud Provider Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_provider_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_subnet",
			Category:         "Cloud",
			ShortDescription: `Manages ACI Cloud Subnet`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_cidr_pool_dn",
					Description: `(Required) Distinguished name of parent CloudCIDRPool object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) CIDR block of Object cloud subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for object cloud subnet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object cloud subnet.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object cloud subnet.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object cloud subnet.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) List of domain applicable to the capability. Allowed values are "public", "private" and "shared". Default is ["private"].`,
				},
				resource.Attribute{
					Name:        "usage",
					Description: `(Optional) The usage of the port. This property shows how the port is used. Allowed values are "user", "gateway" and "infra-router". Default is "user". To make any subnet a Gateway subnet use ` + "`" + `usage` + "`" + ` = "gateway".`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) [AWS Only] Availability zone where the subnet must be deployed. This property can carry both the actual zone or the ACI logical zone name. In the former case, the driver directly uses the value of this property. In the latter case, the Connector has to first resolve the mapping from ACI logical zone to the actual AWS zone. This parameter is required in APIC v5.0 or higher`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_subnet_to_flow_log",
					Description: `(Optional) Relation to class cloudAwsFlowLogPol. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud Subnet. ## Importing ## An existing Cloud Subnet can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_subnet.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_vpn_gateway",
			Category:         "Cloud",
			ShortDescription: `Manages ACI Cloud Vpn Gateway`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"vpn",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_context_profile_dn",
					Description: `(Required) Distinguished name of parent CloudContextProfile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object Cloud Router Profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Cloud Router Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Cloud Router Profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name_alias for object Cloud Router Profile.`,
				},
				resource.Attribute{
					Name:        "num_instances",
					Description: `(Optional) Num instances for object Cloud Router Profile. Default value: "1"`,
				},
				resource.Attribute{
					Name:        "cloud_router_profile_type",
					Description: `(Optional) Component type Allowed values are "host-router" and "vpn-gw". Default value is "vpn-gw".`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_to_vpn_gw_pol",
					Description: `(Optional) Relation to class cloudVpnGwPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_to_direct_conn_pol",
					Description: `(Optional) Relation to class cloudDirectConnPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_to_host_router_pol",
					Description: `(Optional) Relation to class cloudHostRouterPol. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Cloud Router Profile. ## Importing An existing Cloud Router Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_cloud_vpn_gateway.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_configuration_export_policy",
			Category:         "Import/Export",
			ShortDescription: `Manages ACI Configuration Export Policy`,
			Description:      ``,
			Keywords: []string{
				"import",
				"export",
				"configuration",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_configuration_import_policy",
			Category:         "Import/Export",
			ShortDescription: `Manages ACI Configuration Import Policy`,
			Description:      ``,
			Keywords: []string{
				"import",
				"export",
				"configuration",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_connection",
			Category:         "L4-L7 Services",
			ShortDescription: `Manages ACI Connection`,
			Description:      ``,
			Keywords: []string{
				"l4",
				"l7",
				"services",
				"connection",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_console_authentication",
			Category:         "AAA",
			ShortDescription: `Manages ACI Console Authentication`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"console",
				"authentication",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Console Authentication Method.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Console Authentication. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Console Authentication. Type: String.`,
				},
				resource.Attribute{
					Name:        "provider_group",
					Description: `(Optional) Provider Group. An AAA configuration provider group is a group of remote servers supporting the same AAA protocol that will be used for authentication and authorization. When a provider group is specified, only the servers within that group will be used for authentication and authorization. If no provider group is specified, all servers supporting the realm of AAA protocols will be used for authentication and authorization.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `(Optional) Realm. The security method for processing authentication and authorization requests. The realm allows the protected resources on the associated server to be partitioned into a set of protection spaces, each with its own authentication authorization database. This is an abstract class and cannot be instantiated. Allowed values are "ldap", "local", "radius", "rsa", "saml", "tacacs". Type: String.`,
				},
				resource.Attribute{
					Name:        "realm_sub_type",
					Description: `(Optional) Realm subtype that can be default or Duo. Allowed values are "default", "duo". Type: String. (Note: attribute realm_sub_type is supported for version 5 and above of APIC) ## Importing ## An existing ConsoleAuthentication can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_console_authentication.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_contract",
			Category:         "Contract",
			ShortDescription: `Manages ACI Contract`,
			Description:      ``,
			Keywords: []string{
				"contract",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object contract.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object contract.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object contract.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object contract.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) Priority level of the service contract. Allowed values are "unspecified", "level1", "level2", "level3", "level4", "level5" and "level6". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Represents the scope of this contract. If the scope is set as application-profile, the epg can only communicate with epgs in the same application-profile. Allowed values are "global", "tenant", "application-profile" and "context". Default is "context".`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile. Allowed values are "CS0", "CS1", "AF11", "AF12", "AF13", "CS2", "AF21", "AF22", "AF23", "CS3", "AF31", "AF32", "AF33", "CS4", "AF41", "AF42", "AF43", "CS5", "VA", "EF", "CS6", "CS7" and "unspecified". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) To manage filters from the contract resource. It has the attributes like filter_name, annotation, description and name_alias.`,
				},
				resource.Attribute{
					Name:        "filter.filter_name",
					Description: `(Required) Name of the filter object.`,
				},
				resource.Attribute{
					Name:        "filter.description",
					Description: `(Optional) Description for the filter object.`,
				},
				resource.Attribute{
					Name:        "filter.annotation",
					Description: `(Optional) Annotation for filter object.`,
				},
				resource.Attribute{
					Name:        "filter.name_alias",
					Description: `(Optional) Name alias for filter object.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry",
					Description: `(Optional) To manage filter entries for particular filter from the contract resource. It has following attributes.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.filter_entry_name",
					Description: `(Required) Name of Object filter entry.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.entry_annotation",
					Description: `(Optional) Annotation for object filter entry.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.entry_description",
					Description: `(Optional) Description for the filter entry.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.apply_to_frag",
					Description: `(Optional) Flag to determine whether to apply changes to fragment. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.arp_opc",
					Description: `(Optional) Open peripheral codes. Allowed values are "unspecified", "req" and "reply". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.d_from_port",
					Description: `(Optional) Destination From Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified". Allowed values: "unspecified", "ftpData", "smtp", "dns", "http","pop3", "https", "rtsp"`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.d_to_port",
					Description: `(Optional) Destination To Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified". Allowed values: "unspecified", "ftpData", "smtp", "dns", "http","pop3", "https", "rtsp"`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.ether_t",
					Description: `(Optional) Ether type for the entry. Allowed values are "unspecified", "ipv4", "trill", "arp", "ipv6", "mpls_ucast", "mac_security", "fcoe" and "ip". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.icmpv4_t",
					Description: `(Optional) ICMPv4 message type; used when ip_protocol is icmp. Allowed values are "echo-rep", "dst-unreach", "src-quench", "echo", "time-exceeded" and "unspecified". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.icmpv6_t",
					Description: `(Optional) ICMPv6 message type; used when ip_protocol is icmpv6. Allowed values are "unspecified", "dst-unreach", "time-exceeded", "echo-req", "echo-rep", "nbr-solicit", "nbr-advert" and "redirect". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.match_dscp",
					Description: `(Optional) The matching differentiated services code point (DSCP) of the path attached to the layer 3 outside profile. Allowed values are "CS0", "CS1", "AF11", "AF12", "AF13", "CS2", "AF21", "AF22", "AF23", "CS3", "AF31", "AF32", "AF33", "CS4", "AF41", "AF42", "AF43", "CS5", "VA", "EF", "CS6", "CS7" and "unspecified". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.entry_name_alias",
					Description: `(Optional) Name alias for object filter entry.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.prot",
					Description: `(Optional) Level 3 ip protocol. Allowed values are "unspecified", "icmp", "igmp", "tcp", "egp", "igp", "udp", "icmpv6", "eigrp", "ospfigp", "pim" and "l2tp". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.s_from_port",
					Description: `(Optional) Source From Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified". Allowed values: "unspecified", "ftpData", "smtp", "dns", "http","pop3", "https", "rtsp"`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.s_to_port",
					Description: `(Optional) Source To Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified". Allowed values: "unspecified", "ftpData", "smtp", "dns", "http","pop3", "https", "rtsp"`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.stateful",
					Description: `(Optional) Determines if entry is stateful or not. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.tcp_rules",
					Description: `(Optional) TCP Session Rules. Allowed values are "unspecified", "est", "syn", "ack", "fin" and "rst". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_graph_att",
					Description: `(Optional) Relation to class vzRsGraphAtt. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Contract.`,
				},
				resource.Attribute{
					Name:        "filter.id",
					Description: `Exports this attribute for filter object. Set to the Dn for the filter managed by the contract.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.id",
					Description: `Exports this attribute for filter entry object of filter object. Set to the Dn for the filter entry managed by the contract. ## Importing ## An existing Contract can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_contract.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "filter.id",
					Description: `Exports this attribute for filter object. Set to the Dn for the filter managed by the contract.`,
				},
				resource.Attribute{
					Name:        "filter.filter_entry.id",
					Description: `Exports this attribute for filter entry object of filter object. Set to the Dn for the filter entry managed by the contract. ## Importing ## An existing Contract can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_contract.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_contract_subject",
			Category:         "Contract",
			ShortDescription: `Manages ACI Contract Subject`,
			Description:      ``,
			Keywords: []string{
				"contract",
				"subject",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_coop_policy",
			Category:         "System Settings",
			ShortDescription: `Manages ACI COOP Policy`,
			Description:      ``,
			Keywords: []string{
				"system",
				"settings",
				"coop",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object COOP Group Policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Authentication type.The specific type of object or component. Allowed values are "compatible", "strict". Type: String.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object COOP Group Policy. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object COOP Group Policy. Type: String. ## Importing ## An existing COOPGroupPolicy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_coop_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_default_authentication",
			Category:         "AAA",
			ShortDescription: `Manages ACI Default Authentication Method for all Logins`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"default",
				"authentication",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Default Authentication Method for all Logins.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Default Authentication Method for all Logins.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of object Default Authentication Method for all Logins.`,
				},
				resource.Attribute{
					Name:        "fallback_check",
					Description: `(Optional) The parameter to disable fallback in case there are active servers in the default auth type. Allowed values are "false" and "true". Type: String.`,
				},
				resource.Attribute{
					Name:        "provider_group",
					Description: `(Optional) The group of remote servers supporting the same AAA protocol that will be used for authentication and authorization.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `(Optional) The security method for processing authentication and authorization requests. The realm allows the protected resources on the associated server to be partitioned into a set of protection spaces, each with its own authentication authorization database. Allowed values are "ldap", "local", "radius", "rsa", "saml" and "tacacs". Type: String.`,
				},
				resource.Attribute{
					Name:        "realm_sub_type",
					Description: `(Optional) Realm subtype. Allowed values are "default" and "duo". Type: String. ## Importing ## An existing Default Authentication Method for all Logins can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_default_authentication.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_destination_of_redirected_traffic",
			Category:         "L4-L7 Services",
			ShortDescription: `Manages ACI Destination of redirected traffic`,
			Description:      ``,
			Keywords: []string{
				"l4",
				"l7",
				"services",
				"destination",
				"of",
				"redirected",
				"traffic",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_dhcp_option_policy",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI DHCP Option Policy`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"dhcp",
				"option",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_dhcp_relay_policy",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI DHCP Relay Policy`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"dhcp",
				"relay",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_duo_provider_group",
			Category:         "AAA",
			ShortDescription: `Manages ACI Duo Provider Group`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"duo",
				"provider",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object Duo Provider Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Duo Provider Group.`,
				},
				resource.Attribute{
					Name:        "auth_choice",
					Description: `(Optional) Authentication choice of object Duo Provider Group. Allowed values are "CiscoAVPair" and "LdapGroupMap". Default value is "CiscoAVPair". Type: String.`,
				},
				resource.Attribute{
					Name:        "provider_type",
					Description: `(Optional) Type of the Auth Provider. Allowed values are "ldap" and "radius". Default value is "radius". Type: String.`,
				},
				resource.Attribute{
					Name:        "ldap_group_map_ref",
					Description: `(Optional) Reference to LDAP Group Map containing user's group membership info.`,
				},
				resource.Attribute{
					Name:        "sec_fac_auth_methods",
					Description: `(Optional) Second factor authentication methods. Allowed values are "auto", "passcode", "phone" and "push". Default value is "auto". Type: List.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of object Duo Provider Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Duo Provider Group. ## Importing ## An existing Duo Provider Group can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_duo_provider_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_encryption_key",
			Category:         "System Settings",
			ShortDescription: `Manages ACI AES Encryption Passphrase and Keys for Config Export and Import`,
			Description:      ``,
			Keywords: []string{
				"system",
				"settings",
				"encryption",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object AES Encryption Passphrase and Keys for Config Export and Import.`,
				},
				resource.Attribute{
					Name:        "passphrase",
					Description: `(Optional) Parameter to set the passphrase of object AES Encryption Passphrase and Keys for Config Export and Import. Length of ` + "`" + `passphrase` + "`" + ` should be between 16 - 32 characters.`,
				},
				resource.Attribute{
					Name:        "strong_encryption_enabled",
					Description: `(Optional) Parameter indicating whether encryption is weak or strong. This parameter can be set if and only if the ` + "`" + `passphrase` + "`" + ` is set. Allowed values are "yes" and "no". Type: String.`,
				},
				resource.Attribute{
					Name:        "clear_encryption_key",
					Description: `(Optional) Parameter to clear the encryption key, if configured. Allowed values are "yes" and "no". Default value is "no". Type: String.`,
				},
				resource.Attribute{
					Name:        "passphrase_key_derivation_version",
					Description: `(Optional) Version of the algorithm used for forward compatibility. Allowed value is "v1". Default value is "v1".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object AES Encryption Passphrase and Keys for Config Export and Import.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object AES Encryption Passphrase and Keys for Config Export and Import. ## Importing ## An existing AES Encryption Passphrase and Keys for Config Export and Import can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_encryption_key.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_end_point_retention_policy",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI End Point Retention Policy`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"end",
				"point",
				"retention",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object end point retention policy.`,
				},
				resource.Attribute{
					Name:        "descripton",
					Description: `(Optional) Descripton for object end point retention policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object end point retention policy.`,
				},
				resource.Attribute{
					Name:        "bounce_age_intvl",
					Description: `(Optional) The aging interval for a bounce entry. When an endpoint (VM) migrates to another switch, the endpoint is marked as bouncing for the specified aging interval and is deleted afterwards. Allowed value range is "0" - "0xffff". Default is "630"."0" is treated as special value here. Providing interval as "0" is treated as infinite interval.`,
				},
				resource.Attribute{
					Name:        "bounce_trig",
					Description: `(Optional) Specifies whether to install the bounce entry by RARP flood or by COOP protocol. Allowed values are "rarp-flood" and "protocol". Default is "protocol".`,
				},
				resource.Attribute{
					Name:        "hold_intvl",
					Description: `(Optional) A time period during which new endpoint learn events will not be honored. This interval is triggered when the maximum endpoint move frequency is exceeded. Allowed value range is "5" - "0xffff". Default is "300".`,
				},
				resource.Attribute{
					Name:        "local_ep_age_intvl",
					Description: `(Optional) The aging interval for all local endpoints learned in this bridge domain. When 75% of the interval is reached, 3 ARP requests are sent to verify the existence of the endpoint. If no response is received, the endpoint is deleted. Allowed value range is "120" - "0xffff". Default is "900". "0" is treated as special value here. Providing interval as "0" is treated as infinite interval.`,
				},
				resource.Attribute{
					Name:        "move_freq",
					Description: `(Optional) A maximum allowed number of endpoint moves per second. If the move frequency is exceeded, the hold interval is triggered, and new endpoint learn events will not be honored until after the hold interval expires. Allowed value range is "0" - "0xffff". Default is "256"."0" is treated as special value here. Providing interval as "0" is treated as none.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object end point retention policy.`,
				},
				resource.Attribute{
					Name:        "remote_ep_age_intvl",
					Description: `(Optional) The aging interval for all remote endpoints learned in this bridge domain.Allowed value range is "120" - "0xffff". Default is "300". "0" is treated as special value here. Providing interval as "0" is treated as infinite interval. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the End Point Retention Policy. ## Importing ## An existing End Point Retention Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_end_point_retention_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_endpoint_controls",
			Category:         "System Settings",
			ShortDescription: `Manages ACI Endpoint Control`,
			Description:      ``,
			Keywords: []string{
				"system",
				"settings",
				"endpoint",
				"controls",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint Control.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) The administrative state of object Endpoint Control. Allowed values are "disabled" and "enabled".`,
				},
				resource.Attribute{
					Name:        "hold_intvl",
					Description: `(Optional) The period of time before declaring that the neighbor is down of object Endpoint Control. Allowed range: "300" - "3600".`,
				},
				resource.Attribute{
					Name:        "rogue_ep_detect_intvl",
					Description: `(Optional) Rogue Endpoint Detection Interval of object Endpoint Control. Allowed range: "30" - "3600".`,
				},
				resource.Attribute{
					Name:        "rogue_ep_detect_mult",
					Description: `(Optional) Rogue Endpoint Detection Multiplication Factor of object Endpoint Control. Allowed range is "2" - "65535".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Endpoint Control.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of object Endpoint Control. ## Importing ## An existing Endpoint Control can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_endpoint_controls.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_endpoint_ip_aging_profile",
			Category:         "System Settings",
			ShortDescription: `Manages ACI Endpoint IP Aging Profile`,
			Description:      ``,
			Keywords: []string{
				"system",
				"settings",
				"endpoint",
				"ip",
				"aging",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) The administrative state of the object Endpoint IP Aging Profile. Allowed values are "disabled" and "enabled".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint IP Aging Profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Endpoint IP Aging Profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Endpoint IP Aging Profile. ## Importing ## An existing Endpoint IP Aging Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_endpoint_ip_aging_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_endpoint_loop_protection",
			Category:         "System Settings",
			ShortDescription: `Manages ACI Endpoint Loop Protection`,
			Description:      ``,
			Keywords: []string{
				"system",
				"settings",
				"endpoint",
				"loop",
				"protection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint Loop Protection Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Endpoint Loop Protection. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Endpoint Loop Protection. Type: String.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action.Sets the action to take when a loop is detected. Allowed values are "bd-learn-disable", "port-disable". Type: List.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Admin State.The administrative state of the object or policy. Allowed values are "disabled", "enabled". Type: String.`,
				},
				resource.Attribute{
					Name:        "loop_detect_intvl",
					Description: `(Optional) Loop Detection Interval.Sets the loop detection interval, which specifies the time to detect a loop. Allowed range is "30"-"300". Type: String.`,
				},
				resource.Attribute{
					Name:        "loop_detect_mult",
					Description: `(Optional) Loop Detection Multiplier.Sets the loop detection multiplication factor, which is the number of times a single Endpoint moves between ports within the Detection interval. Allowed range is "1"-"255". Type: String ## Importing ## An existing EPLoopProtectionPolicy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_endpoint_loop_protection.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_endpoint_security_group",
			Category:         "Application Management",
			ShortDescription: `Manages ACI Endpoint Security Group`,
			Description:      ``,
			Keywords: []string{
				"application",
				"management",
				"endpoint",
				"security",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_endpoint_security_group_epg_selector",
			Category:         "Application Management",
			ShortDescription: `Manages ACI Endpoint Security Group EPG Selector`,
			Description:      ``,
			Keywords: []string{
				"application",
				"management",
				"endpoint",
				"security",
				"group",
				"epg",
				"selector",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_security_group_dn",
					Description: `(Required) Distinguished name of parent Endpoint Security Group object.`,
				},
				resource.Attribute{
					Name:        "match_epg_dn",
					Description: `(Required) EPG Dn to be associated.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint Security Group EPG Selector. ## Importing ## An existing Endpoint Security Group EPG Selector can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_endpoint_security_group_epg_selector.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_endpoint_security_group_selector",
			Category:         "Application Management",
			ShortDescription: `Manages ACI Endpoint Security Group Selector`,
			Description:      ``,
			Keywords: []string{
				"application",
				"management",
				"endpoint",
				"security",
				"group",
				"selector",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_security_group_dn",
					Description: `(Required) Distinguished name of parent Endpoint Security Group object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint Security Group Selector.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Endpoint Security Group Selector.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of object Endpoint Security Group Selector.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Endpoint Security Group Selector.`,
				},
				resource.Attribute{
					Name:        "match_expression",
					Description: `(Optional) Expression used to define matching tags. ## Importing ## An existing EndpointSecurityGroupSelector can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import endpoint_security_group_selector.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_endpoint_security_group_tag_selector",
			Category:         "Application Management",
			ShortDescription: `Manages ACI Endpoint Security Group Tag Selector`,
			Description:      ``,
			Keywords: []string{
				"application",
				"management",
				"endpoint",
				"security",
				"group",
				"tag",
				"selector",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_security_group_dn",
					Description: `(Required) Distinguished name of parent EndpointSecurityGroup object.`,
				},
				resource.Attribute{
					Name:        "match_key",
					Description: `(Required) Match key of object Endpoint Security Group Tag Selector.`,
				},
				resource.Attribute{
					Name:        "match_value",
					Description: `(Required) Match value of object Endpoint Security Group Tag Selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint Security Group Tag Selector.`,
				},
				resource.Attribute{
					Name:        "value_operator",
					Description: `(Optional) Match Value Operator. Allowed values are "contains", "equals", "regex", and default value is "equals". Type: String. ## Importing ## An existing EndpointSecurityGroupTagSelector can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_endpoint_security_group_tag_selector.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_epg_to_contract",
			Category:         "Application Management",
			ShortDescription: `Manages ACI EPG to contract relationship.`,
			Description:      ``,
			Keywords: []string{
				"application",
				"management",
				"epg",
				"to",
				"contract",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_epg_to_contract_interface",
			Category:         "Application Management",
			ShortDescription: `Manages ACI Contract Interface Relationship`,
			Description:      ``,
			Keywords: []string{
				"application",
				"management",
				"epg",
				"to",
				"contract",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_epg_dn",
					Description: `(Required) Distinguished name of the parent Application EPG object.`,
				},
				resource.Attribute{
					Name:        "contract_interface_dn",
					Description: `(Required) Distinguished name of the Contract Interface object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Contract Interface Relationship object.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The Contract Interface Relationship priority. Allowed values are "level1", "level2", "level3", "level4", "level5", "level6", "unspecified", and default value is "unspecified". Type: String. ## Importing ## An existing Contract Interface Relationship object can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_epg_to_contract_interface.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_epg_to_domain",
			Category:         "Application Management",
			ShortDescription: `Manages ACI epg to Domain`,
			Description:      ``,
			Keywords: []string{
				"application",
				"management",
				"epg",
				"to",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_epg_dn",
					Description: `(Required) Distinguished name of parent ApplicationEPG object.`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Required) Distinguished Name of Target Domain object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Domain.`,
				},
				resource.Attribute{
					Name:        "binding_type",
					Description: `(Optional) Binding type for object Domain. Allowed values: "none", "staticBinding", "dynamicBinding", "ephemeral". Default value: "none"`,
				},
				resource.Attribute{
					Name:        "allow_micro_seg",
					Description: `(Optional) Boolean flag for allow micro segment. default value will be "false". "true" maps to class_pref="useg" and "false maps to class_pref="encap"`,
				},
				resource.Attribute{
					Name:        "custom_epg_name",
					Description: `(Optional) Custom EPG name used as name of the VMM port group for the domain.`,
				},
				resource.Attribute{
					Name:        "enhanced_lag_policy",
					Description: `(Optional) Distinguished Name of the Enhanced LACP LAG Policy (class lacpEnhancedLagPol) to associate with the VMM domain.`,
				},
				resource.Attribute{
					Name:        "delimiter",
					Description: `(Optional) Delimiter for object Domain.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) Port encapsulation.`,
				},
				resource.Attribute{
					Name:        "encap_mode",
					Description: `(Optional) Encap mode for object Domain. Allowed values: "auto", "vlan", "vxlan". Default value: "auto"`,
				},
				resource.Attribute{
					Name:        "epg_cos",
					Description: `(Optional) Epg cos for object Domain. Allowed values: "Cos0", "Cos1", "Cos2", "Cos3", "Cos4", "Cos5", "Cos6", "Cos7". Default value: "Cos0"`,
				},
				resource.Attribute{
					Name:        "epg_cos_pref",
					Description: `(Optional) Epg cos pref for object Domain. Allowed values: "disabled", "enabled". Default value: "disabled"`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) Determines when policies are pushed to cam. Allowed values: "immediate", "lazy". Default value: "lazy"`,
				},
				resource.Attribute{
					Name:        "lag_policy_name",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "netflow_dir",
					Description: `(Optional) Netflow dir for object Domain. Allowed values: "ingress", "egress", "both". Default value: "both"`,
				},
				resource.Attribute{
					Name:        "netflow_pref",
					Description: `(Optional) Netflow pref for object Domain. Allowed values: "disabled", "enabled"`,
				},
				resource.Attribute{
					Name:        "num_ports",
					Description: `(Optional) Number of ports existing operationally in module. Default value: "0"`,
				},
				resource.Attribute{
					Name:        "port_allocation",
					Description: `(Optional) Port allocation for object Domain. Allowed values: "none", "elastic", "fixed". Default value: "none"`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) Primary encap for object Domain. Default value: "unknown".`,
				},
				resource.Attribute{
					Name:        "primary_encap_inner",
					Description: `(Optional) Primary encap inner for object Domain. Default value: "unknown".`,
				},
				resource.Attribute{
					Name:        "res_imedcy",
					Description: `(Optional) Policy resolution. Allowed values: "immediate", "lazy", "pre-provision". Default value: "lazy"`,
				},
				resource.Attribute{
					Name:        "secondary_encap_inner",
					Description: `(Optional) Secondary encap inner for object Domain.Default value: "unknown".`,
				},
				resource.Attribute{
					Name:        "switching_mode",
					Description: `(Optional) Switching mode for object domain. Allowed values: "native", "AVE". Default value: "native"`,
				},
				resource.Attribute{
					Name:        "vmm_allow_promiscuous",
					Description: `(Optional) Allow promiscuous for object Vmm security policy.`,
				},
				resource.Attribute{
					Name:        "vmm_forged_transmits",
					Description: `(Optional) Forged transmits for object Vmm security policy.`,
				},
				resource.Attribute{
					Name:        "vmm_mac_changes",
					Description: `(Optional) Mac changes for object Vmm security policy. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Domain.`,
				},
				resource.Attribute{
					Name:        "vmm_id",
					Description: `which is set to the Dn of the VMM Security Policy. ## Importing ## An existing Domain can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vmm_id",
					Description: `which is set to the Dn of the VMM Security Policy. ## Importing ## An existing Domain can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_epg_to_static_path",
			Category:         "Application Management",
			ShortDescription: `Manages ACI EPG to Static Path`,
			Description:      ``,
			Keywords: []string{
				"application",
				"management",
				"epg",
				"to",
				"static",
				"path",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_epg_dn",
					Description: `(Required) Distinguished name of parent ApplicationEPG object.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Required) Encapsulation of the Static Path.`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Required) tdn of Object Static Path.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Static Path.`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) Immediacy of the Static Path. Allowed values: "immediate", "lazy". Default value: "lazy"`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Mode of the static association with the path. Allowed values: "regular", "native", "untagged". Default value: "regular"`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) Primary encap for object Static Path. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Static Path. ## Importing ## An existing Static Path can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_epg_to_static_path.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_epgs_using_function",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI EPGs Using Function`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"epgs",
				"using",
				"function",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_generic_dn",
					Description: `(Required) Distinguished name of parent AccessGeneric object.`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Required) tDn of Object EPGs Using Function.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Required) Vlan number encap.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) annotation for object EPGs Using Function.`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) Instrumentation immediacy. Allowed values: "immediate", "lazy". Default value: "lazy".`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Bgp domain mode. Allowed values: "regular", "native", "untagged". Default value: "regular"`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) Primary encap for object EPGs Using Function. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the EPGs Using Function. ## Importing ## An existing EPGs Using Function can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_ep_gs_using_function.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_error_disable_recovery",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Error Disable Recovery`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"error",
				"disable",
				"recovery",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Error Disable Recovery. Type String.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Error Disable Recovery. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Error Disable Recovery. Type: String.`,
				},
				resource.Attribute{
					Name:        "err_dis_recov_intvl",
					Description: `(Optional) Error Disable Recovery Interval.Sets the error disable recovery interval, which specifies the time to recover from an error-disabled state. Allowed range is "30" - "65535". Type: String.`,
				},
				resource.Attribute{
					Name:        "edr_event",
					Description: `(Optional) To manage Error Disable Recovery Event from the Error Disable Recovery Policy resource.`,
				},
				resource.Attribute{
					Name:        "edr_event.event",
					Description: `(Required) Event of object Error Disabled Recovery. The error disable recovery event type. Allowed values are "event-arp-inspection", "event-bpduguard", "event-debug-1", "event-debug-2", "event-debug-3", "event-debug-4", "event-debug-5", "event-dhcp-rate-lim", "event-ep-move", "event-ethpm", "event-ip-addr-conflict", "event-ipqos-dcbxp-compat-failure", "event-ipqos-mgr-error", "event-link-flap", "event-loopback", "event-mcp-loop", "event-psec-violation", "event-sec-violation", "event-set-port-state-failed", "event-storm-ctrl", "event-stp-inconsist-vpc-peerlink", "event-syserr-based", "event-udld", "unknown". Type: String.`,
				},
				resource.Attribute{
					Name:        "edr_event.recover",
					Description: `(Optional) Enables or disables Error Disable Recovery. Allowed values are "no", "yes". Type: String.`,
				},
				resource.Attribute{
					Name:        "edr_event.name",
					Description: `(Optional) Name of object Error Disable Recovery Event. Type: String.`,
				},
				resource.Attribute{
					Name:        "edr_event.name_alias",
					Description: `(Optional) Name Alias of object Error Disable Recovery Event. Type: String.`,
				},
				resource.Attribute{
					Name:        "edr_event.description",
					Description: `(Optional) Description of object Error Disable Recovery Event. Type: String.`,
				},
				resource.Attribute{
					Name:        "edr_event.annotation",
					Description: `(Optional) Annotation of object Error Disable Recovery Event. Type String. ## Importing ## An existing Error Disable Recovery can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_error_disable_recovery.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_external_network_instance_profile",
			Category:         "L3Out",
			ShortDescription: `Manages ACI External Network Instance Profile`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"external",
				"network",
				"instance",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fabric_if_pol",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI fabric if pol`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"fabric",
				"if",
				"pol",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "auto_neg",
					Description: `(Optional) Policy auto negotiation for object fabric if pol. Allowed values: "on", "off". Default value is "on".`,
				},
				resource.Attribute{
					Name:        "fec_mode",
					Description: `(Optional) Forwarding error correction for object fabric if pol. Allowed values: "inherit", "cl91-rs-fec", "cl74-fc-fec", "ieee-rs-fec", "cons16-rs-fec", "kp-fec", "disable-fec". Default value is "inherit".`,
				},
				resource.Attribute{
					Name:        "link_debounce",
					Description: `(Optional) Link debounce interval for object fabric if pol. Range of allowed values: "0" to "5000". Default value is "100".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) Port speed for object fabric if pol. Allowed values: "unknown", "100M", "1G", "10G", "25G", "40G", "50G", "100G","200G", "400G", "inherit". Default value is "inherit". ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Link Level Policy. ## Importing ## An existing Link Level Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: <https://www.terraform.io/docs/import/index.html> ` + "`" + `` + "`" + `` + "`" + `bash terraform import aci_fabric_if_pol.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fabric_node_control",
			Category:         "Fabric Policies",
			ShortDescription: `Manages ACI Fabric Node Control`,
			Description:      ``,
			Keywords: []string{
				"fabric",
				"policies",
				"node",
				"control",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object Fabric Node Control.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Fabric Node Control.`,
				},
				resource.Attribute{
					Name:        "control",
					Description: `(Optional) Fabric node control bitmask of object Fabric Node Control. Allowed values are "Dom" and "None". Default value is "None".`,
				},
				resource.Attribute{
					Name:        "feature_sel",
					Description: `(Optional) Feature Selection of object Fabric Node Control. Allowed values are "analytics", "netflow" and "telemetry". Default value is "telemetry".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Fabric Node Control.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Fabric Node Control. ## Importing ## An existing Fabric Node Control can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_fabric_node_control.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fabric_node_member",
			Category:         "Fabric Inventory",
			ShortDescription: `Manages ACI Fabric Node Member`,
			Description:      ``,
			Keywords: []string{
				"fabric",
				"inventory",
				"node",
				"member",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fabric_wide_settings",
			Category:         "System Settings",
			ShortDescription: `Manages ACI Fabric-Wide Settings Policy`,
			Description:      ``,
			Keywords: []string{
				"system",
				"settings",
				"fabric",
				"wide",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of object Fabric-Wide Settings Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Fabric-Wide Settings Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Fabric-Wide Settings Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of object Fabric-Wide Settings Policy.`,
				},
				resource.Attribute{
					Name:        "disable_ep_dampening",
					Description: `(Optional) Disable Ep Dampening knob of object Fabric-Wide Settings Policy. Allowed values are "yes" and "no".`,
				},
				resource.Attribute{
					Name:        "enable_mo_streaming",
					Description: `(Optional) Enable MO streaming of object Fabric-Wide Settings Policy. Allowed values are "yes" and "no".`,
				},
				resource.Attribute{
					Name:        "enable_remote_leaf_direct",
					Description: `(Optional) Enable remote leaf direct communication of object Fabric-Wide Settings Policy. Allowed values are "yes" and "no".`,
				},
				resource.Attribute{
					Name:        "enforce_subnet_check",
					Description: `(Optional) Enforce subnet check of object Fabric-Wide Settings Policy. Allowed values are "yes" and "no".`,
				},
				resource.Attribute{
					Name:        "opflexp_authenticate_clients",
					Description: `(Optional) Opflexp Client Certificates for authentication of object Fabric-Wide Settings Policy. Allowed values are "yes" and "no".`,
				},
				resource.Attribute{
					Name:        "opflexp_use_ssl",
					Description: `(Optional) SSL transport for Opflexp indicator of object Fabric-Wide Settings Policy. Allowed values are "yes" and "no".`,
				},
				resource.Attribute{
					Name:        "restrict_infra_vlan_traffic",
					Description: `(Optional) Intra Leaf Communication traffic indicator of object Fabric-Wide Settings Policy. Allowed values are "yes" and "no". (Note: attribute restrict_infra_vlan_traffic is supported for version 5 and above of APIC)`,
				},
				resource.Attribute{
					Name:        "unicast_xr_ep_learn_disable",
					Description: `(Optional) Disable xrLeanrs indicator of object Fabric-Wide Settings Policy. Allowed values are "yes" and "no".`,
				},
				resource.Attribute{
					Name:        "validate_overlapping_vlans",
					Description: `(Optional) Validate Overlapping VLANS indicator of object Fabric-Wide Settings Policy. Allowed values are "yes" and "no". ## Importing ## An existing Fabric-WideSettings Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_fabric_wide_settings.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fc_domain",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI FC Domain`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"fc",
				"domain",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fex_bundle_group",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Fex Bundle Group`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"fex",
				"bundle",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fex_profile",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI FEX Profile`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"fex",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_file_remote_path",
			Category:         "Import/Export",
			ShortDescription: `Manages ACI Remote Path of a File`,
			Description:      ``,
			Keywords: []string{
				"import",
				"export",
				"file",
				"remote",
				"path",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) Hostname or IP for export destination of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) Authentication Type Choice. Allowed values are "usePassword" and "useSshKeyContents". Default value is "usePassword". Type: String.`,
				},
				resource.Attribute{
					Name:        "identity_private_key_contents",
					Description: `(Optional) SSH Private Key File contents for datatransfer. Must be set if ` + "`" + `auth_type` + "`" + ` is equal to "useSshKeyContents".`,
				},
				resource.Attribute{
					Name:        "identity_private_key_passphrase",
					Description: `(Optional) Passphrase given at the identity key creation. Should be set if and only if ` + "`" + `identity_private_key_contents` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Transfer protocol to be used for data export of object File Remote Path .Allowed values are "ftp", "scp" and "sftp". Default value is "sftp". Type: String. Value "ftp" cannot be set if ` + "`" + `auth_type` + "`" + ` is equal to "useSshKeyContents".`,
				},
				resource.Attribute{
					Name:        "remote_path",
					Description: `(Optional) Path where data will reside in the destination of object File Remote Path(The first character of remote_path should be '/').`,
				},
				resource.Attribute{
					Name:        "remote_port",
					Description: `(Optional) Remote port for data export destination of object File Remote Path. Range: "0" - "65535". Default value is "0".`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional) Username to be used to transfer data to destination of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "user_passwd",
					Description: `(Optional) Password to be used to transfer data to destination of object File Remote Path. Must be set if ` + "`" + `auth_type` + "`" + ` is equal to "usePassword".`,
				},
				resource.Attribute{
					Name:        "relation_file_rs_a_remote_host_to_epg",
					Description: `(Optional) Represents the relation to a Attachable Target Group (class fvATg). A source relation to the endpoint group through which the remote host is reachable. Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_file_rs_a_remote_host_to_epp",
					Description: `(Optional) Represents the relation to a Relation to Remote Host Reachability EPP (class fvAREpP). A source relation to the abstract representation of the resolvable endpoint profile. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object File Remote Path. ## Importing ## An existing File Remote Path can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_file_remote_path.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_filter",
			Category:         "Contract",
			ShortDescription: `Manages ACI Filter`,
			Description:      ``,
			Keywords: []string{
				"contract",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object filter.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object filter.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object filter.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object filter.`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_filt_graph_att",
					Description: `(Optional) Relation to class vnsInTerm. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_fwd_r_flt_p_att",
					Description: `(Optional) Relation to class vzAFilterableUnit. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_rev_r_flt_p_att",
					Description: `(Optional) Relation to class vzAFilterableUnit. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Filter. ## Importing ## An existing Filter can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_filter.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_filter_entry",
			Category:         "Contract",
			ShortDescription: `Manages ACI Filter Entry`,
			Description:      ``,
			Keywords: []string{
				"contract",
				"filter",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter_dn",
					Description: `(Required) Distinguished name of parent Filter object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object filter entry.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object filter entry.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object filter entry.`,
				},
				resource.Attribute{
					Name:        "apply_to_frag",
					Description: `(Optional) Flag to determine whether to apply changes to fragment. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "arp_opc",
					Description: `(Optional) Open peripheral codes. Allowed values are "unspecified", "req" and "reply". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "d_from_port",
					Description: `(Optional) Destination From Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified". Allowed values: "unspecified", "ftpData", "smtp", "dns", "http","pop3", "https", "rtsp"`,
				},
				resource.Attribute{
					Name:        "d_to_port",
					Description: `(Optional) Destination To Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified". Allowed values: "unspecified", "ftpData", "smtp", "dns", "http","pop3", "https", "rtsp"`,
				},
				resource.Attribute{
					Name:        "ether_t",
					Description: `(Optional) Ether type for the entry. Allowed values are "unspecified", "ipv4", "trill", "arp", "ipv6", "mpls_ucast", "mac_security", "fcoe" and "ip". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "icmpv4_t",
					Description: `(Optional) ICMPv4 message type; used when ip_protocol is icmp. Allowed values are "echo-rep", "dst-unreach", "src-quench", "echo", "time-exceeded" and "unspecified". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "icmpv6_t",
					Description: `(Optional) ICMPv6 message type; used when ip_protocol is icmpv6. Allowed values are "unspecified", "dst-unreach", "time-exceeded", "echo-req", "echo-rep", "nbr-solicit", "nbr-advert" and "redirect". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "match_dscp",
					Description: `(Optional) The matching differentiated services code point (DSCP) of the path attached to the layer 3 outside profile. Allowed values are "CS0", "CS1", "AF11", "AF12", "AF13", "CS2", "AF21", "AF22", "AF23", "CS3", "AF31", "AF32", "AF33", "CS4", "AF41", "AF42", "AF43", "CS5", "VA", "EF", "CS6", "CS7" and "unspecified". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object filter entry.`,
				},
				resource.Attribute{
					Name:        "prot",
					Description: `(Optional) Level 3 IP protocol. Allowed values are "unspecified", "icmp", "igmp", "tcp", "egp", "igp", "udp", "icmpv6", "eigrp", "ospfigp", "pim", "l2tp" and a number between "0" and "255". Default is "unspecified".`,
				},
				resource.Attribute{
					Name:        "s_from_port",
					Description: `(Optional) Source From Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified". Allowed values: "unspecified", "ftpData", "smtp", "dns", "http","pop3", "https", "rtsp"`,
				},
				resource.Attribute{
					Name:        "s_to_port",
					Description: `(Optional) Source To Port. Accepted values are any valid TCP/UDP port range. Default is "unspecified". Allowed values: "unspecified", "ftpData", "smtp", "dns", "http","pop3", "https", "rtsp"`,
				},
				resource.Attribute{
					Name:        "stateful",
					Description: `(Optional) Determines if entry is stateful or not. Allowed values are "yes" and "no". Default is "no".`,
				},
				resource.Attribute{
					Name:        "tcp_rules",
					Description: `(Optional) TCP Rules. TCP Session Rules Allowed values are "ack", "est", "fin", "rst", "syn", "unspecified", and default value is "unspecified". Type: List. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Filter Entry. ## Importing ## An existing Filter Entry can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_filter_entry.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_firmware_download_task",
			Category:         "Firmware",
			ShortDescription: `Manages ACI Firmware Download Task`,
			Description:      ``,
			Keywords: []string{
				"firmware",
				"download",
				"task",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_firmware_group",
			Category:         "Firmware",
			ShortDescription: `Manages ACI Firmware Group`,
			Description:      ``,
			Keywords: []string{
				"firmware",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object Firmware Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Firmware Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Firmware Group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name_alias for object Firmware Group.`,
				},
				resource.Attribute{
					Name:        "firmware_group_type",
					Description: `(Optional) Component type. DefaultValue : "range" Allowed values: "ALL", "range", "ALL_IN_POD"`,
				},
				resource.Attribute{
					Name:        "relation_firmware_rs_fwgrpp",
					Description: `(Optional) Relation to class firmwareFwP. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Firmware Group. ## Importing ## An existing Firmware Group can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_firmware_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_firmware_policy",
			Category:         "Firmware",
			ShortDescription: `Manages ACI Firmware Policy`,
			Description:      ``,
			Keywords: []string{
				"firmware",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_function_node",
			Category:         "L4-L7 Services",
			ShortDescription: `Manages ACI Function Node`,
			Description:      ``,
			Keywords: []string{
				"l4",
				"l7",
				"services",
				"function",
				"node",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_global_security",
			Category:         "AAA",
			ShortDescription: `Manages ACI Global Security`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"global",
				"security",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Global Security.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Global Security.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of object Global Security.`,
				},
				resource.Attribute{
					Name:        "pwd_strength_check",
					Description: `(Optional) Password Strength Check.The password strength check specifies if the system enforces the strength of the user password. Allowed values are "no", "yes", and default value is "yes". Type: String.`,
				},
				resource.Attribute{
					Name:        "change_count",
					Description: `(Optional) Number of Password Changes in Interval.The number of password changes allowed within the change interval. Allowed range is 0-10 and default value is "2".`,
				},
				resource.Attribute{
					Name:        "change_during_interval",
					Description: `(Optional) Password Policy.The change count/change interval policy selector. This property enables you to select an option for enforcing password change. Allowed values are "disable", "enable", and default value is "enable". Type: String.`,
				},
				resource.Attribute{
					Name:        "change_interval",
					Description: `(Optional) Change Interval in Hours.A time interval for limiting the number of password changes. Allowed range is 0-745 and default value is "48".`,
				},
				resource.Attribute{
					Name:        "expiration_warn_time",
					Description: `(Optional) Password Expiration Warn Time in Days.A warning period before password expiration. A warning will be displayed when a user logs in within this number of days of an impending password expiration. Allowed range is 0-30 and default value is "15".`,
				},
				resource.Attribute{
					Name:        "history_count",
					Description: `(Optional) Password History Count.How many retired passwords are stored in a user's password history. Allowed range is 0-15 and default value is "5".`,
				},
				resource.Attribute{
					Name:        "no_change_interval",
					Description: `(Optional) No Password Change Interval in Hours.A minimum period after a password change before the user can change the password again. Allowed range is 0-745 and default value is "24".`,
				},
				resource.Attribute{
					Name:        "block_duration",
					Description: `(Optional) Duration in minutes for which login should be blocked.Duration in minutes for which future logins should be blocked Allowed range is 1-1440 and default value is "60".`,
				},
				resource.Attribute{
					Name:        "enable_login_block",
					Description: `(Optional) Enable blocking of user logins after failed attempts. Allowed values are "disable", "enable", and default value is "disable". Type: String.`,
				},
				resource.Attribute{
					Name:        "max_failed_attempts",
					Description: `(Optional) Maximum continuous failed logins before blocking user.max failed login attempts before blocking user login Allowed range is 1-15 and default value is "5".`,
				},
				resource.Attribute{
					Name:        "max_failed_attempts_window",
					Description: `(Optional) Time period for maximum continuous failed logins.times in minutes for max login failures to occur before blocking the user Allowed range is 1-720 and default value is "5".`,
				},
				resource.Attribute{
					Name:        "maximum_validity_period",
					Description: `(Optional) Maximum Validity Period in hours.The maximum validity period for a webt oken. Allowed range is 4-24 and default value is "24".`,
				},
				resource.Attribute{
					Name:        "session_record_flags",
					Description: `(Optional) Session Recording Options.Enables or disables a refresh in the session records. Allowed values are "login", "logout", "refresh" and default value is ["login", "logout", "refresh"]. Type: List.`,
				},
				resource.Attribute{
					Name:        "ui_idle_timeout_seconds",
					Description: `(Optional) GUI Idle Timeout in Seconds.The maximum interval time the GUI remains idle before login needs to be refreshed. Allowed range is 60-65525 and default value is "1200".`,
				},
				resource.Attribute{
					Name:        "webtoken_timeout_seconds",
					Description: `(Optional) Timeout in Seconds.The web token timeout interval. Allowed range is 300-9600 and default value is "600".`,
				},
				resource.Attribute{
					Name:        "relation_aaa_rs_to_user_ep",
					Description: `(Optional) Represents the relation to a Global Security (class aaaUserEp). Type: String. ## Importing ## An existing UserManagement can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_global_security.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_hsrp_group_policy",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI HSRP Group Policy`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"hsrp",
				"group",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_hsrp_interface_policy",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI HSRP Interface Policy`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"hsrp",
				"interface",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of HSRP interface policy object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for HSRP interface policy object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for HSRP interface policy object.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) List of control states for HSRP interface policy object. Allowed values are "bia" and "bfd".`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Optional) Administrative port delay for HSRP interface policy object.Range: "0" to "10000". Default value is "0".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for HSRP interface policy object.`,
				},
				resource.Attribute{
					Name:        "reload_delay",
					Description: `(Optional) Reload delay for HSRP interface policy object.Range: "0" to "10000". Default value is "0". ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the HSRP Interface Policy. ## Importing ## An existing HSRP Interface Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_hsrp_interface_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_imported_contract",
			Category:         "Contract",
			ShortDescription: `Manages ACI Imported Contract`,
			Description:      ``,
			Keywords: []string{
				"contract",
				"imported",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_interface_blacklist",
			Category:         "Fabric Inventory",
			ShortDescription: `Manages ACI Out of Service Fabric Path`,
			Description:      ``,
			Keywords: []string{
				"fabric",
				"inventory",
				"interface",
				"blacklist",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pod_id",
					Description: `(Required) The Pod ID of the switch that own the interface that need to be disabled.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `(Required) The Node ID of the switch that own the interface that need to be disabled.`,
				},
				resource.Attribute{
					Name:        "fex_id",
					Description: `(Required) The FEX ID of the FEX that own the interface that need to be disabled.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface name of the interface that need to be disabled.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the blacklist object (fabricRsOosPath). ## Importing ## An existing blacklist can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_interface_blacklist.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_interface_fc_policy",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Interface FC Policy`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"interface",
				"fc",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_isis_domain_policy",
			Category:         "System Settings",
			ShortDescription: `Manages ACI ISIS Domain Policy and ISIS Level`,
			Description:      ``,
			Keywords: []string{
				"system",
				"settings",
				"isis",
				"domain",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object ISIS Domain Policy.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) Maximum Transmission Unit of object ISIS Domain Policy. Allowed range: "256" - "4352".`,
				},
				resource.Attribute{
					Name:        "redistrib_metric",
					Description: `(Optional) Metric used for redistributed routes. Allowed range: "1" - "63".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object ISIS Domain Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of object ISIS Domain Policy.`,
				},
				resource.Attribute{
					Name:        "lsp_fast_flood",
					Description: `(Optional) The IS-IS Fast-Flooding of LSPs improves Intermediate System-to-Intermediate System (IS-IS) convergence time when new link-state packets (LSPs) are generated in the network and shortest path first (SPF) is triggered by the new LSPs. Allowed values are "disabled" and "enabled".`,
				},
				resource.Attribute{
					Name:        "lsp_gen_init_intvl",
					Description: `(Optional) The LSP generation initial wait interval. Allowed range: "50" - "120000".`,
				},
				resource.Attribute{
					Name:        "lsp_gen_max_intvl",
					Description: `(Optional) The LSP generation maximum wait interval. Allowed range: "50" - "120000".`,
				},
				resource.Attribute{
					Name:        "lsp_gen_sec_intvl",
					Description: `(Optional) The LSP generation second wait interval. Allowed range: "50" - "120000".`,
				},
				resource.Attribute{
					Name:        "spf_comp_init_intvl",
					Description: `(Optional) The SPF computation frequency initial wait interval. Allowed range: "50" - "120000".`,
				},
				resource.Attribute{
					Name:        "spf_comp_max_intvl",
					Description: `(Optional) The SPF computation frequency maximum wait interval. Allowed range: "50" - "120000".`,
				},
				resource.Attribute{
					Name:        "spf_comp_sec_intvl",
					Description: `(Optional) The SPF computation frequency second wait interval. Allowed range: "50" - "120000".`,
				},
				resource.Attribute{
					Name:        "isis_level_name",
					Description: `(Optional) (Optional) The name of ISIS Level object.`,
				},
				resource.Attribute{
					Name:        "isis_level_type",
					Description: `(Optional) The type of ISIS Level object. Allowed values are "l1" and "l2". Default value is "l1". ## Importing ## An existing ISIS DomainPolicy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_isis_domain_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l2_domain",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI L2 Domain`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"l2",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object L2 Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object L2 Domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object L2 Domain.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vlan_ns",
					Description: `(Optional) Relation to class fvnsVlanInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vlan_ns_def",
					Description: `(Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vip_addr_ns",
					Description: `(Optional) Relation to class fvnsAddrInst. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_extnw_rs_out",
					Description: `(Optional) Relation to class infraAccGrp. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_dom_vxlan_ns_def",
					Description: `(Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the L2 Domain . ## Importing ## An existing L2 Domain can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_l2_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l2_interface_policy",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI L2 Interface Policy`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"l2",
				"interface",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l2_outside",
			Category:         "L2Out",
			ShortDescription: `Manages ACI L2 Outside`,
			Description:      ``,
			Keywords: []string{
				"l2out",
				"l2",
				"outside",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l2out_extepg",
			Category:         "L2Out",
			ShortDescription: `Manages ACI ACI L2-Out External EPG`,
			Description:      ``,
			Keywords: []string{
				"l2out",
				"extepg",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3_domain_profile",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI L3 Domain Profile`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"l3",
				"domain",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object l3 domain profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object l3 domain profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object l3 domain profile.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vlan_ns",
					Description: `(Optional) Relation to class fvnsVlanInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vlan_ns_def",
					Description: `(Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vip_addr_ns",
					Description: `(Optional) Relation to class fvnsAddrInst. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_extnw_rs_out",
					Description: `(Optional) Relation to class infraAccGrp. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_dom_vxlan_ns_def",
					Description: `(Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the L3 Domain Profile. ## Importing ## An existing L3 Domain Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_l3_domain_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3_ext_subnet",
			Category:         "L3Out",
			ShortDescription: `Manages ACI l3 extension subnet`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"l3",
				"ext",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "external_network_instance_profile_dn",
					Description: `(Required) Distinguished name of parent ExternalNetworkInstanceProfile object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) ip of Object l3 extension subnet.`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `(Optional) Aggregate Routes for l3 extension subnet. Allowed values are "import-rtctrl", "export-rtctrl","shared-rtctrl" and "none".`,
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
					Description: `(Optional) The list of domain applicable to the capability. Allowed values are "import-rtctrl", "export-rtctrl", "import-security", "shared-security" and "shared-rtctrl". Default is "import-security".`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_subnet_to_profile",
					Description: `(Optional) Relation to Route Control Profile (class rtctrlProfile). Cardinality - N_TO_ONE. Type - Set of Map.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_subnet_to_profile.tn_rtctrl_profile_dn",
					Description: `(Optional) Associates the external EPGs with the route control profiles.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_subnet_to_profile.direction",
					Description: `(Required) Relation to configure route map for each BGP peer in the inbound and outbound directions.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_subnet_to_profile",
					Description: `(Optional) Relation to class rtctrlProfile. Cardinality - N_TO_M. Type - Set of Map.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_subnet_to_rt_summ",
					Description: `(Optional) Relation to class rtsumARtSummPol. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the l3 extension subnet. ## Importing An existing Subnet can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_l3_ext_subnet.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3_interface_policy",
			Category:         "Fabric Policies",
			ShortDescription: `Manages ACI L3 Interface Policy`,
			Description:      ``,
			Keywords: []string{
				"fabric",
				"policies",
				"l3",
				"interface",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object L3 Interface Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object L3 Interface Policy.`,
				},
				resource.Attribute{
					Name:        "bfd_isis",
					Description: `(Optional) BFD ISIS Configuration for object L3 Interface Policy. Allowed values are "disabled" and "enabled". Default value is "disabled".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object L3 Interface Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object L3 Interface Policy. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the L3 Interface Policy. ## Importing ## An existing L3 Interface Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_l3_interface_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3_outside",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3 Outside`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"l3",
				"outside",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_bfd_interface_profile",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3out BFD Interface Profile`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"bfd",
				"interface",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_bgp_external_policy",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3-out BGP External Policy`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"bgp",
				"external",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_bgp_protocol_profile",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3out BGP Protocol Profile`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"bgp",
				"protocol",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_floating_svi",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3out Floating SVI`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"floating",
				"svi",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_hsrp_interface_group",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3out HSRP Interface Group`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"hsrp",
				"interface",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_hsrp_interface_profile",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3-out HSRP interface profile`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"hsrp",
				"interface",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_hsrp_secondary_vip",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3out HSRP Secondary VIP`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"hsrp",
				"secondary",
				"vip",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_loopback_interface_profile",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3out Loopback Interface Profile`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"loopback",
				"interface",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fabric_node_dn",
					Description: `(Required) Distinguished name of parent fabric node object.`,
				},
				resource.Attribute{
					Name:        "addr",
					Description: `(Required) Address of L3out lookback interface profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for L3out lookback interface profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for L3out lookback interface profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for L3out lookback interface profile. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the L3out Loopback Interface Profile. ## Importing ## An existing L3out Loopback Interface Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_l3out_loopback_interface_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_ospf_external_policy",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3-out OSPF External Policy`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"ospf",
				"external",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_ospf_interface_profile",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3out OSPF Interface Profile`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"ospf",
				"interface",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_path_attachment",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3-out Path Attachment`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"path",
				"attachment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_path_attachment_secondary_ip",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3-out Path Attachment Secondary IP`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"path",
				"attachment",
				"secondary",
				"ip",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_route_tag_policy",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI L3out Route Tag Policy`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"l3out",
				"route",
				"tag",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_static_route",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3out Static Route`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"static",
				"route",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_static_route_next_hop",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3out Static Route Next Hop`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"static",
				"route",
				"next",
				"hop",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_vpc_member",
			Category:         "L3Out",
			ShortDescription: `Manages ACI L3out VPC Member`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"vpc",
				"member",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l4_l7_service_graph_template",
			Category:         "L4-L7 Services",
			ShortDescription: `Manages ACI L4-L7 Service Graph Template`,
			Description:      ``,
			Keywords: []string{
				"l4",
				"l7",
				"services",
				"service",
				"graph",
				"template",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_lacp_policy",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI LACP Policy`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"lacp",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ldap_group_map",
			Category:         "AAA",
			ShortDescription: `Manages ACI LDAP Group Map`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"ldap",
				"group",
				"map",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object LDAP Group Map.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of object LDAP Group Map. Allowed values are "duo" and "ldap".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object LDAP Group Map.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object LDAP Group Map.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Annotation of object LDAP Group Map. ## Importing ## An existing LDAPGroupMap can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_ldap_group_map.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ldap_group_map_rule",
			Category:         "AAA",
			ShortDescription: `Manages ACI LDAP Group Map Rule`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"ldap",
				"group",
				"map",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object LDAP Group Map Rule.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of object LDAP Group Map Rule. Allowed values are "duo" and "ldap". Type: String.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object LDAP Group Map Rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object LDAP Group Map Rule. Type: String.`,
				},
				resource.Attribute{
					Name:        "groupdn",
					Description: `(Optional) LDAP Group DN to compare with LDAP search query for user's membership. Type: String.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object LDAP Group Map Rule. Type: String. ## Importing ## An existing LDAPGroupMapRule can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_ldap_group_map_rule.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ldap_group_map_rule_to_group_map",
			Category:         "AAA",
			ShortDescription: `Manages ACI LDAP Group Map Rule to Group Map Ref`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"ldap",
				"group",
				"map",
				"rule",
				"to",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ldap_group_map_dn",
					Description: `(Required) Distinguished name of parent LDAP Group Map object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object LDAP Group Map Rule to Group Map Ref.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object LDAP Group Map Rule to Group Map Ref.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object LDAP Group Map Rule to Group Map Ref.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object LDAP Group Map Rule to Group Map Ref. ## Importing ## An existing LDAP Group Map Rule to Group Map Ref can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_ldap_group_map_rule_to_group_map.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ldap_provider",
			Category:         "AAA",
			ShortDescription: `Manages ACI LDAP Provider`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"ldap",
				"provider",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Host name or IP address of object LDAP Provider.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of LDAP Provider. Allowed values are "ldap" and "duo".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object LDAP Provider.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object LDAP Provider.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of object LDAP Provider.`,
				},
				resource.Attribute{
					Name:        "ssl_validation_level",
					Description: `(Optional) The LDAP Server SSL Certificate validation level. Allowed values are "permissive" and "strict". Default value is "strict". Type: String.`,
				},
				resource.Attribute{
					Name:        "attribute",
					Description: `(Optional) The attribute to be downloaded that contains user role and domain information. Default value is "CiscoAVPair".`,
				},
				resource.Attribute{
					Name:        "basedn",
					Description: `(Optional) The LDAP base DN to be used in a user search.`,
				},
				resource.Attribute{
					Name:        "enable_ssl",
					Description: `(Optional) A property for enabling an SSL connection with the LDAP provider. Allowed values are "no" and "yes". Default value is "no". Type: String.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) The LDAP filter to be used in a user search. Default value is "sAMAccountName=$userid".`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) A password for the AAA provider database.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring. Allowed values are "disabled" and "enabled". Default value is "disabled". Type: String.`,
				},
				resource.Attribute{
					Name:        "monitoring_password",
					Description: `(Optional) Periodic Server Monitoring Password.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username. Default value is "default".`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The service port number for the LDAP service. Allowed range: "1" - "65535". Default value is "389".`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Retry count of object LDAP Provider. Allowed range: "1" - "5". Default value is "1".`,
				},
				resource.Attribute{
					Name:        "rootdn",
					Description: `(Optional) The root DN or bind DN of the LDAP provider.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The timeout for communication with an LDAP provider server. Allowed range: "5" - "60". Default value is "30". (NOTE: For "duo" LDAP providers, the value of timeout should be greater than or equal to "30".)`,
				},
				resource.Attribute{
					Name:        "relation_aaa_rs_prov_to_epp",
					Description: `(Optional) Represents the relation to a Relation to AProvider Reachability EPP (class fvAREpP). Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_aaa_rs_sec_prov_to_epg",
					Description: `(Optional) Represents the relation to a Attachable Target Group (class fvATg). A source relation to the endpoint group through which the provider server is reachable. Type: String. ## Importing ## An existing LDAP Provider can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_ldap_provider.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_access_bundle_policy_group",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI leaf access bundle policy group`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"leaf",
				"bundle",
				"policy",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_access_port_policy_group",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Leaf Access Port Policy Group`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"leaf",
				"port",
				"policy",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object leaf access port policy group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object leaf access port policy group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object leaf access port policy group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object leaf access port policy group.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_span_v_src_grp",
					Description: `(Optional) Relation to class spanVSrcGrp. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_stormctrl_if_pol",
					Description: `(Optional) Relation to class stormctrlIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_poe_if_pol",
					Description: `(Optional) Relation to class poeIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_lldp_if_pol",
					Description: `(Optional) Relation to class lldpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_macsec_if_pol",
					Description: `(Optional) Relation to class macsecIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_dpp_if_pol",
					Description: `(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_h_if_pol",
					Description: `(Optional) Relation to class fabricHIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_netflow_monitor_pol",
					Description: `(Optional) Relation to class netflowMonitorPol. Cardinality - N_TO_M. Type - Set of Map. - ` + "`" + `flt_type` + "`" + ` - (Required) Netflow IP filter type. Allowed values: "ce", "ipv4", "ipv6". - ` + "`" + `target_dn` + "`" + ` - (Required) Distinguished name of target Netflow Monitor object.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_l2_port_auth_pol",
					Description: `(Optional) Relation to class l2PortAuthPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_mcp_if_pol",
					Description: `(Optional) Relation to class mcpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_l2_port_security_pol",
					Description: `(Optional) Relation to class l2PortSecurityPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_copp_if_pol",
					Description: `(Optional) Relation to class coppIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_span_v_dest_grp",
					Description: `(Optional) Relation to class spanVDestGrp. Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_dwdm_if_pol",
					Description: `(Optional) Relation to class dwdmIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_pfc_if_pol",
					Description: `(Optional) Relation to class qosPfcIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_sd_if_pol",
					Description: `(Optional) Relation to class qosSdIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_mon_if_infra_pol",
					Description: `(Optional) Relation to class monInfraPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_fc_if_pol",
					Description: `(Optional) Relation to class fcIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_ingress_dpp_if_pol",
					Description: `(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_cdp_if_pol",
					Description: `(Optional) Relation to class cdpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_l2_if_pol",
					Description: `(Optional) Relation to class l2IfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_stp_if_pol",
					Description: `(Optional) Relation to class stpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_qos_egress_dpp_if_pol",
					Description: `(Optional) Relation to class qosDppPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_att_ent_p",
					Description: `(Optional) Relation to class infraAttEntityP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_l2_inst_pol",
					Description: `(Optional) Relation to class l2InstPol. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Leaf Access Port Policy Group. ## Importing ## An existing Leaf Access Port Policy Group can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_leaf_access_port_policy_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_breakout_port_group",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Leaf Breakout Port Group`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"leaf",
				"breakout",
				"port",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_interface_profile",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Leaf Interface Profile`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"leaf",
				"interface",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object leaf interface profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object leaf interface profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object leaf interface profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object leaf interface profile. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Leaf Interface Profile. ## Importing ## An existing Leaf Interface Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_leaf_interface_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_profile",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Leaf Profile`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"leaf",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_selector",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Leaf Selector`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"leaf",
				"selector",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_lldp_interface_policy",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI LLDP Interface Policy`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"lldp",
				"interface",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object LLDP Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_rx_st",
					Description: `(Optional) Admin receive state. Allowed values are "enabled" and "disabled". Default value is "enabled".`,
				},
				resource.Attribute{
					Name:        "admin_tx_st",
					Description: `(Optional) Admin transmit state. Allowed values are "enabled" and "disabled". Default value is "enabled".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object LLDP Interface Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object LLDP Interface Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object LLDP Interface Policy. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the LLDP Interface Policy. ## Importing ## An existing LLDP Interface Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: <https://www.terraform.io/docs/import/index.html> ` + "`" + `` + "`" + `` + "`" + `bash terraform import aci_lldp_interface_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_local_user",
			Category:         "AAA",
			ShortDescription: `Manages ACI Local User`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"local",
				"user",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_logical_device_context",
			Category:         "L4-L7 Services",
			ShortDescription: `Manages ACI Logical Device Context`,
			Description:      ``,
			Keywords: []string{
				"l4",
				"l7",
				"services",
				"logical",
				"device",
				"context",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_logical_interface_context",
			Category:         "L4-L7 Services",
			ShortDescription: `Manages ACI Logical Interface Context.`,
			Description:      ``,
			Keywords: []string{
				"l4",
				"l7",
				"services",
				"logical",
				"interface",
				"context",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_logical_interface_profile",
			Category:         "L3Out",
			ShortDescription: `Manages ACI Logical Interface Profile`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"logical",
				"interface",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_logical_node_profile",
			Category:         "L3Out",
			ShortDescription: `Manages ACI Logical Node Profile`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"logical",
				"node",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_logical_node_to_fabric_node",
			Category:         "L3Out",
			ShortDescription: `Manages ACI Logical Node to Fabric Node`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"logical",
				"node",
				"to",
				"fabric",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_node_profile_dn",
					Description: `(Required) Distinguished name of parent LogicalNodeProfile object.`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Required) Tdn of Object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "config_issues",
					Description: `(Optional) Configuration issues. Allowed values: "anchor-node-mismatch", "bd-profile-missmatch", "loopback-ip-missing", "missing-mpls-infra-l3out", "missing-rs-export-route-profile", "node-path-misconfig", "node-vlif-misconfig", "none", "routerid-not-changable-with-mcast", "subnet-mismatch". Default value: "none"`,
				},
				resource.Attribute{
					Name:        "rtr_id",
					Description: `(Optional) Router identifier`,
				},
				resource.Attribute{
					Name:        "rtr_id_loop_back",
					Description: `(Optional) Allowed values: "yes", "no". Default value: "yes" ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Fabric Node. ## Importing ## An existing Fabric Node can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_logical_node_to_fabric_node.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_login_domain",
			Category:         "AAA",
			ShortDescription: `Manages ACI Login Domain`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"login",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object Login Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Login Domain.`,
				},
				resource.Attribute{
					Name:        "provider_group",
					Description: `(Optional) Provider Group. An AAA configuration provider group is a group of remote servers supporting the same AAA protocol that will be used for authentication and authorization. When a provider group is specified, only the servers within that group will be used for authentication and authorization. If no provider group is specified, all servers supporting the realm of AAA protocols will be used for authentication and authorization. (Note: Attribute provider_group will be set only for the value "none" of attribute "realm", for other values server will not allow to set "provider_group" attribute.)`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `(Optional) Realm. The security method for processing authentication requests. The realm allows the protected resources on the associated server to be partitioned into a set of protection spaces, each with its own authentication authorization database. Allowed values are "ldap", "local", "none", "radius", "rsa", "saml", "tacacs". Default value is "local". Type: String.`,
				},
				resource.Attribute{
					Name:        "realm_sub_type",
					Description: `(Optional) Realm subtype that can be default or Duo. Allowed values are "default", "duo". Default value is "default". Type: String. (Note: attribute realm_sub_type is supported for version 5 and above of APIC)`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Login Domain. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Login Domain. Type: String. ## Importing ## An existing LoginDomain can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_login_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_login_domain_provider",
			Category:         "AAA",
			ShortDescription: `Manages ACI Login Domain Provider`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"login",
				"domain",
				"provider",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "parent_dn",
					Description: `(Required) Distinguished name of parent.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object Login Domain Provider.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Login Domain Provider.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Login Domain Provider.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Login Domain Provider.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Order in which Providers are Tried. The relative priority in which the AAA provider will be contacted within the provider group. Allowed Range: "0" - "16". Allowed value "lowest-available". Default value is "0". (NOTE: "lowest-available" will set lowest value of order and will be translated by postConfig code to the numeric order value of 0.) ## Importing ## An existing Login Domain Provider can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_login_domain_provider.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_maintenance_group_node",
			Category:         "Firmware",
			ShortDescription: `Manages ACI Maintenance Group Node`,
			Description:      ``,
			Keywords: []string{
				"firmware",
				"maintenance",
				"group",
				"node",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pod_maintenance_group_dn",
					Description: `(Required) Distinguished name of parent POD Maintenance Group Object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Maintenance Group Node Object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for Maintenance Group Node Object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for Maintenance Group Node Object.`,
				},
				resource.Attribute{
					Name:        "from_",
					Description: `(Optional) From value for Maintenance Group Node Object. Range : 1 - 16000. DefaultValue : "1"`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for Maintenance Group Node Object.`,
				},
				resource.Attribute{
					Name:        "to_",
					Description: `(Optional) To value for Maintenance Group Node Object. Range : 1 - 16000. DefaultValue : "1" ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Maintenance Group Node. ## Importing ## An existing Maintenance Group Node can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_maintenance_group_node.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_maintenance_policy",
			Category:         "Firmware",
			ShortDescription: `Manages ACI Maintenance Policy`,
			Description:      ``,
			Keywords: []string{
				"firmware",
				"maintenance",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_managed_node_connectivity_group",
			Category:         "Node Management",
			ShortDescription: `Manages ACI Managed Node Connectivity Group`,
			Description:      ``,
			Keywords: []string{
				"node",
				"management",
				"managed",
				"connectivity",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object Managed Node Connectivity Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Managed Node Connectivity Group. ## Importing ## An existing Managed Node Connectivity Group can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_managed_node_connectivity_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_match_community_terms",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI Match Community Term`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"match",
				"community",
				"terms",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "match_rule_dn",
					Description: `(Required) Distinguished name of the parent Match Rule object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Match Community Term object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Match Community Term object.`,
				},
				resource.Attribute{
					Name:        "match_community_factors",
					Description: `(Optional) Match Community Factor object.Type: Block.`,
				},
				resource.Attribute{
					Name:        "community",
					Description: `(Required) The community of the Match Community Factor object. Type: String.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The scope of the Match Community Factor object. Allowed values are "transitive", "non-transitive", and default value is "transitive". Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the Match Community Factor object. ## Importing ## An existing Match Community Term can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_match_community_terms.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_match_regex_community_terms",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI Match Rule Based on Community Regular Expression`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"match",
				"regex",
				"community",
				"terms",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "match_rule_dn",
					Description: `(Required) Distinguished name of the parent Match Rule object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Match Rule Based on Community Regular Expression object.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community Type of the Match Rule Based on Community Regular Expression object. Allowed values are "extended", "regular", and default value is "regular". Type: String.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) Regular Expression.A regular expression used to specify a pattern to match against the community string. ## Importing ## An existing Match Rule Based on Community Regular Expression can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_match_regex_community_terms.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_match_route_destination_rule",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI Match Route Destination Rule`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"match",
				"route",
				"destination",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "match_rule_dn",
					Description: `(Required) Distinguished name of parent Match Rule object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) Ip of object Match Route Destination Rule.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Match Route Destination Rule.`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `(Optional) Aggregated Route. Allowed values are "false", "true" and default value is "false". Type: String.`,
				},
				resource.Attribute{
					Name:        "greater_than_mask",
					Description: `(Optional) Start of Prefix Length. Allowed range is 0-128 and default value is "0".`,
				},
				resource.Attribute{
					Name:        "less_than_mask",
					Description: `(Optional) End of Prefix Length. Allowed range is 0-128 and default value is "0". ## Importing ## An existing Match Route Destination Rule can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_match_route_destination_rule.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_match_rule",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI Match Rule`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"match",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object Match Rule.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Match Rule. ## Importing ## An existing Match Rule can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_match_rule.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_mcp_instance_policy",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI MCP Instance Policy`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"mcp",
				"instance",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object MCP Instance Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Admin State. The administrative state of the object or policy. Allowed values are "disabled", "enabled". Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object MCP Instance Policy. Type: String.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias for object MCP Instance Policy. Type: String.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) Controls. The control state. Allowed values are "pdu-per-vlan", "stateful-ha". Type: List.`,
				},
				resource.Attribute{
					Name:        "init_delay_time",
					Description: `(Optional) Init Delay Time. Allowed range is "0"-"1800". Type: String.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Secret Key. The key or password used to uniquely identify this configuration object.`,
				},
				resource.Attribute{
					Name:        "loop_detect_mult",
					Description: `(Optional) Loop Detection Multiplier. Allowed range is "1"-"255". Type: String.`,
				},
				resource.Attribute{
					Name:        "loop_protect_act",
					Description: `(Optional) Loop Protection Action. Allowed values are "port-disable","none". Type: String.`,
				},
				resource.Attribute{
					Name:        "tx_freq",
					Description: `(Optional) Transmission Frequency. Sets the transmission frequency of the instance advertisements. Allowed range is "0"-"300". Type: String.(Note: For value less than "2", loop_protect_act attribute needs to be "port-disable", Accepted Range of tx_freq is 100ms to 300s so total value of tx_freq & tx_freq_msec should in Accepted range.)`,
				},
				resource.Attribute{
					Name:        "tx_freq_msec",
					Description: `(Optional) Transmission Frequency. Sets the transmission frequency of mcp advertisements in milliseconds Allowed range is "0"-"999". Type: String.(Note: For value "0" of tx_freq, range of tx_freq_msec is "100"-"999".) ## Importing ## An existing MiscablingProtocolInstancePolicy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_mcp_instance_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_mgmt_preference",
			Category:         "System Settings",
			ShortDescription: `Manages ACI Mgmt Preference`,
			Description:      ``,
			Keywords: []string{
				"system",
				"settings",
				"mgmt",
				"preference",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "interface_pref",
					Description: `(Optional) Management interface that has to be used. Allowed values are "inband" and "ooband".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Mgmt Preference.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Mgmt Preference.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Mgmt Preference. ## Importing ## An existing Mgmt Preference can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_mgmt_preference.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_mgmt_zone",
			Category:         "Node Management",
			ShortDescription: `Manages ACI Management Zone`,
			Description:      ``,
			Keywords: []string{
				"node",
				"management",
				"mgmt",
				"zone",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_miscabling_protocol_interface_policy",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Mis-cabling Protocol Interface Policy`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"miscabling",
				"protocol",
				"interface",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object miscabling protocol interface policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Administrative state of the object or policy. Allowed values are "enabled" and "disabled". Default is "enabled".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object miscabling protocol interface policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object miscabling protocol interface policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object miscabling protocol interface policy. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Mis-cabling Protocol Interface Policy. ## Importing ## An existing Mis-cabling Protocol Interface Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: <https://www.terraform.io/docs/import/index.html> ` + "`" + `` + "`" + `` + "`" + `bash terraform import aci_miscabling_protocol_interface_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_monitoring_policy",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI Monitoring Policy`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"monitoring",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object monitoring policy.`,
				},
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Tenant dn for monitoring policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for monitoring policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object monitoring policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object monitoring policy. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the monitoring Policy. ## Importing ## An existing monitoring Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_monitoring_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_node_block",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Node Block`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"node",
				"block",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_node_block_firmware",
			Category:         "Firmware",
			ShortDescription: `Manages ACI Node Block`,
			Description:      ``,
			Keywords: []string{
				"firmware",
				"node",
				"block",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "firmware_group_dn",
					Description: `(Required) Distinguished name of parent Firmware Group Object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object Node Block.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for Object Node Block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for Object Node Block.`,
				},
				resource.Attribute{
					Name:        "from_",
					Description: `(Optional) From value for Object Node Block. Range : 1 - 16000. DefaultValue : "1".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for Object Node Block.`,
				},
				resource.Attribute{
					Name:        "to_",
					Description: `(Optional) To value for Object Node Block.Range : 1 - 16000. DefaultValue : "1" ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Node Block. ## Importing ## An existing Node Block can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_node_block_firmware.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_node_mgmt_epg",
			Category:         "Node Management",
			ShortDescription: `Manages ACI Node Management EPg`,
			Description:      ``,
			Keywords: []string{
				"node",
				"management",
				"mgmt",
				"epg",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ospf_interface_policy",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI OSPF Interface Policy`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"ospf",
				"interface",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object OSPF interface policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object OSPF interface policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object OSPF interface policy.`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `(Optional) The OSPF cost for the interface. The cost (also called metric) of an interface in OSPF is an indication of the overhead required to send packets across a certain interface. The cost of an interface is inversely proportional to the bandwidth of that interface. A higher bandwidth indicates a lower cost. There is more overhead (higher cost) and time delays involved in crossing a 56k serial line than crossing a 10M ethernet line. The formula used to calculate the cost is: cost= 10000 0000/bandwidth in bps For example, it will cost 10 EXP8/10 EXP7 = 10 to cross a 10M Ethernet line and will cost 10 EXP8/1544000 = 64 to cross a T1 line. By default, the cost of an interface is calculated based on the bandwidth; you can force the cost of an interface with the ip OSPF cost value interface sub-configuration mode command. Allowed value range is "0" - "65535". Default is "unspecified(0)".`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) List of interface policy controls. Allowed values are "unspecified", "passive", "mtu-ignore", "advert-subnet" and "bfd". Default is ["unspecified"]. E.g., ["bfd", "passive"]. "unspecified" should not be added along with other attributes.`,
				},
				resource.Attribute{
					Name:        "dead_intvl",
					Description: `(Optional) The interval between hello packets from a neighbor before the router declares the neighbor as down. This value must be the same for all networking devices on a specific network. Specifying a smaller dead interval (seconds) will give faster detection of a neighbor being down and improve convergence, but might cause more routing instability. Allowed value range is "1" - "65535". Default value is "40".`,
				},
				resource.Attribute{
					Name:        "hello_intvl",
					Description: `(Optional) The interval between hello packets that OSPF sends on the interface. Note that the smaller the hello interval, the faster topological changes will be detected, but more routing traffic will ensue. This value must be the same for all routers and access servers on a specific network. Allowed value range is "1" - "65535". Default value is "10".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object OSPF interface policy.`,
				},
				resource.Attribute{
					Name:        "nw_t",
					Description: `(Optional) The OSPF interface policy network type. OSPF supports point-to-point and broadcast. Allowed values are "unspecified", "p2p" and "bcast". Default value is "unspecified".`,
				},
				resource.Attribute{
					Name:        "pfx_suppress",
					Description: `(Optional) PFX suppression for object OSPF interface policy. Allowed values are "inherit", "enable" and "disable". Default value is "inherit".`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The OSPF interface priority used to determine the designated router (DR) on a specific network. The router with the highest OSPF priority on a segment will become the DR for that segment. The same process is repeated for the backup designated router (BDR). In the case of a tie, the router with the highest RID will win. The default for the interface OSPF priority is one. Remember that the DR and BDR concepts are per multiaccess segment. Allowed value range is "0" - "255". Default value is "1".`,
				},
				resource.Attribute{
					Name:        "rexmit_intvl",
					Description: `(Optional) The interval between LSA retransmissions. The retransmit interval occurs while the router is waiting for an acknowledgement from the neighbor router that it received the LSA. If no acknowlegment is received at the end of the interval, then the LSA is resent. Allowed value range is "1" - "65535". Default value is "5".`,
				},
				resource.Attribute{
					Name:        "xmit_delay",
					Description: `(Optional) The delay time needed to send an LSA update packet. OSPF increments the LSA age time by the transmit delay amount before transmitting the LSA update. You should take into account the transmission and propagation delays for the interface when you set this value. Allowed value range is "1" - "450". Default is "1". ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the OSPF Interface Policy. ## Importing ## An existing OSPF Interface Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_ospf_interface_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ospf_route_summarization",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI OSPF Route Summarization`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"ospf",
				"route",
				"summarization",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ospf_timers",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI OSPF Timers`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"ospf",
				"timers",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_physical_domain",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Physical Domain`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"physical",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object physical domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object physical domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object physical domain.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vlan_ns",
					Description: `(Optional) Relation to class fvnsVlanInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vlan_ns_def",
					Description: `(Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_vip_addr_ns",
					Description: `(Optional) Relation to class fvnsAddrInst. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_dom_vxlan_ns_def",
					Description: `(Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Physical Domain. ## Importing ## An existing Physical Domain can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: <https://www.terraform.io/docs/import/index.html> ` + "`" + `` + "`" + `` + "`" + `bash terraform import aci_physical_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_pod_maintenance_group",
			Category:         "Firmware",
			ShortDescription: `Manages ACI POD Maintenance Group`,
			Description:      ``,
			Keywords: []string{
				"firmware",
				"pod",
				"maintenance",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_port_security_policy",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Port Security Policy`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"port",
				"security",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_port_tracking",
			Category:         "System Settings",
			ShortDescription: `Manages ACI Port Tracking`,
			Description:      ``,
			Keywords: []string{
				"system",
				"settings",
				"port",
				"tracking",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Port Tracking. Type: String.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Port Tracking. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Port Tracking. Type: String.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Port Tracking State.The administrative state of the object or policy. Allowed values are "off", "on". Type: String.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Optional) Delay Timeout.The administrative port delay. Allowed range is "1"-"300". Type: String.`,
				},
				resource.Attribute{
					Name:        "include_apic_ports",
					Description: `(Optional) Include APIC Ports when port tracking is triggered. Allowed values are "no", "yes". Type: String. (Note: attribute include_apic_ports is supported for version 5 and above of APIC)`,
				},
				resource.Attribute{
					Name:        "minlinks",
					Description: `(Optional) Minimum links left up before trigger. Allowed range is "0"-"48". Type: String. ## Importing ## An existing PortTracking can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_port_tracking.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_qos_instance_policy",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI QOS Instance Policy`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"qos",
				"instance",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object QOS Instance Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object QOS Instance Policy. Type: String.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias for object QOS Instance Policy. Type: String.`,
				},
				resource.Attribute{
					Name:        "etrap_age_timer",
					Description: `(Optional) E-trap flow age out timer. Min Allowed Value is "0".`,
				},
				resource.Attribute{
					Name:        "etrap_bw_thresh",
					Description: `(Optional) Track activeness of elephant flow. Min Allowed Value is "0".`,
				},
				resource.Attribute{
					Name:        "etrap_byte_ct",
					Description: `(Optional) E-trap elephant flow identifier. Min Allowed Value is "0".`,
				},
				resource.Attribute{
					Name:        "etrap_st",
					Description: `(Optional) E-trap enable knob. E-trap parameters. Allowed values are "no", "yes". Type: String.`,
				},
				resource.Attribute{
					Name:        "fabric_flush_interval",
					Description: `(Optional) Fabric Flush Interval in ms. Allowed range is "100"-"1000". Type: String.`,
				},
				resource.Attribute{
					Name:        "fabric_flush_st",
					Description: `(Optional) Fabric PFC Flush enable knob. Fabric Flush parameters Allowed values are "no", "yes". Type: String.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) Global Control Settings. The control state. Allowed values are "dot1p-preserve", "none". Type: String.`,
				},
				resource.Attribute{
					Name:        "uburst_spine_queues",
					Description: `(Optional) Micro burst spine queues percent. Allowed range is "0"-"100". Type: String. (Note: attribute uburst_spine_queues is supported for version 5 and above of APIC)`,
				},
				resource.Attribute{
					Name:        "uburst_tor_queues",
					Description: `(Optional) Micro burst tor queues percent. Allowed range is "0"-"100". Type: String. (Note: attribute uburst_tor_queues is supported for version 5 and above of APIC) ## Importing ## An existing QOSInstancePolicy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_qos_instance_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_radius_provider",
			Category:         "AAA",
			ShortDescription: `Manages ACI RADIUS Provider`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"radius",
				"provider",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Host name or IP address of object RADIUS Provider.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of object RADIUS Provider. Allowed values are "duo" and "radius".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object RADIUS Provider.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object RADIUS Provider.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object RADIUS Provider.`,
				},
				resource.Attribute{
					Name:        "auth_port",
					Description: `(Optional) The service port number for the RADIUS service. Allowed range: "1" - "65535". Default value is "1812".`,
				},
				resource.Attribute{
					Name:        "auth_protocol",
					Description: `(Optional) Authentication Protocol.The RADIUS authentication protocol. Allowed values are "chap", "mschap" and "pap". Default value is "pap". Type: String.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) A password for the AAA provider database.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring. Allowed values are "disabled" and "enabled". Default value is "disabled". Type: String.`,
				},
				resource.Attribute{
					Name:        "monitoring_password",
					Description: `(Optional) Periodic Server Monitoring Password.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username. Default value is "default".`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Number of retries for a for communication with a RADIUS provider server. Allowed range is "1" - "5". Default value is "1".`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The timeout for communication with a RADIUS provider server. Allowed range: "1" - "60". Default value is "5". (NOTE: For "duo" RADIUS providers, the value of timeout should be greater than or equal to "30".)`,
				},
				resource.Attribute{
					Name:        "relation_aaa_rs_prov_to_epp",
					Description: `(Optional) Represents the relation to a Relation to AProvider Reachability EPP (class fvAREpP). Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_aaa_rs_sec_prov_to_epg",
					Description: `(Optional) Represents the relation to a Attachable Target Group (class fvATg). A source relation to the endpoint group through which the provider server is reachable. Type: String. ## Importing ## An existing RADIUSProvider can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_radius_provider.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_radius_provider_group",
			Category:         "AAA",
			ShortDescription: `Manages ACI RADIUS Provider Group`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"radius",
				"provider",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object RADIUS Provider Group. Type: String.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object RADIUS Provider Group. Type: String.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object RADIUS Provider Group. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object RADIUS Provider Group. Type: String. ## Importing ## An existing RADIUSProviderGroup can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_radius_provider_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ranges",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI VLAN Pool Ranges`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"ranges",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vlan_pool_dn",
					Description: `(Required) Distinguished name of parent VLAN Pool object (from aci_vlan_pool resource/data source).`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `(Required) From of Object VLAN Pool ranges. Allowed value min: vlan-1, max: vlan-4094.`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `(Required) To of Object VLAN Pool ranges. Allowed value min: vlan-1, max: vlan-4094.`,
				},
				resource.Attribute{
					Name:        "alloc_mode",
					Description: `(Optional) Alloc mode for object VLAN Pool ranges. Allowed values: "dynamic", "static", "inherit". Default is "inherit".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object VLAN Pool ranges.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object VLAN Pool ranges.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) System role type. Allowed values: "external", "internal". Default is "external".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object VLAN Pool ranges. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Ranges. ## Importing ## An existing Ranges can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: <https://www.terraform.io/docs/import/index.html> ` + "`" + `` + "`" + `` + "`" + `bash terraform import aci_ranges.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_recurring_window",
			Category:         "Scheduler",
			ShortDescription: `Manages ACI Recurring Window`,
			Description:      ``,
			Keywords: []string{
				"scheduler",
				"recurring",
				"window",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scheduler_dn",
					Description: `(Required) Distinguished name of parent Scheduler object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object Recurring Window.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Recurring Window.`,
				},
				resource.Attribute{
					Name:        "concur_cap",
					Description: `(Optional) Maximum Concurrent Tasks. The concurrency capacity limit. This is the maximum number of tasks that can be processed concurrently. Range: "1" - "65535". Default value is "unlimited"(If user sets "0" as a value, provider will accept it but it'll set it as "unlimited"). Type: String.`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `(Optional) Recurring Window Schedule Day. The day of the week that the recurring window begins. Allowed values are "Friday", "Monday", "Saturday", "Sunday", "Thursday", "Tuesday", "Wednesday", "even-day", "every-day", "odd-day", and default value is "every-day". Type: String.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `(Optional) Schedule Hour. The hour that the recurring window begins. Range: "0" - "23". Default value is "0".`,
				},
				resource.Attribute{
					Name:        "minute",
					Description: `(Optional) Schedule Minute. The minute that the recurring window begins. Range: "0" - "59". Default value is "0".`,
				},
				resource.Attribute{
					Name:        "node_upg_interval",
					Description: `(Optional) Delay between node upgrades. Delay between node upgrades in seconds. Range: "0" - "18000". Default value is "0".`,
				},
				resource.Attribute{
					Name:        "proc_break",
					Description: `(Optional) procBreak. A period of time taken between processing of items within the concurrency cap. Allowed Min Value: "00:00:00:00.001"(Format is DD:HH:MM:SS.Milliseconds). Default value is "none" (If user sets "00:00:00:00.000" as a value, provider will accept it but it'll set it as "none"). Type: String.`,
				},
				resource.Attribute{
					Name:        "proc_cap",
					Description: `(Optional) procCap. Processing size capacity limitation specification. Indicates the limit of items to be processed within this window. Range: "1" - "65535". Default value is "unlimited" (If user sets "0" as a value, provider will accept it but it'll set it as "unlimited"). Type: String.`,
				},
				resource.Attribute{
					Name:        "time_cap",
					Description: `(Optional) Maximum Running Time. The processing time capacity limit. This is the maximum duration of the window. Allowed Range: "00:00:00:00.001" to "00:23:59:59.000"(Max input should be less than 24 Hours, Format is DD:HH:MM:SS.Milliseconds). Default value is "unlimited" (If user sets "00:00:00:00.000" as a value, provider will accept it but it'll set it as "unlimited"). Type: String. ## Importing ## An existing RecurringWindow can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_recurring_window.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_rest",
			Category:         "Resources",
			ShortDescription: `Manages ACI Model Objects via REST API calls. Any Model Object that is not supported by provider can be created/managed using this resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_rest_managed",
			Category:         "Resources",
			ShortDescription: `Manages ACI Model Objects via REST API calls. This resource can only manage a single API object and its direct children. It is able to read the state and therefore reconcile configuration drift.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_route_control_context",
			Category:         "Tenant Policies",
			ShortDescription: `Manages ACI Route Control Context`,
			Description:      ``,
			Keywords: []string{
				"tenant",
				"policies",
				"route",
				"control",
				"context",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route_control_profile_dn",
					Description: `(Required) Distinguished name of parent Route Control Profile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object Route Control Context.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Route Control Context.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action. The action required when the condition is met. Allowed values are "deny", "permit", and default value is "permit". Type: String.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Local Order.The order of the policy context. Allowed range is 0-9 and default value is "0".`,
				},
				resource.Attribute{
					Name:        "set_rule",
					Description: `(Optional) Represents the relation to an Attribute Profile (class rtctrlAttrP). Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_rtctrl_rs_ctx_p_to_subj_p",
					Description: `(Optional) Represents the relation to a Subject Profile (class rtctrlSubjP). Type: List. ## Importing ## An existing Route Control Context can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_route_control_context.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_route_control_profile",
			Category:         "L3Out",
			ShortDescription: `Manages ACI Route Control Profile`,
			Description:      ``,
			Keywords: []string{
				"l3out",
				"route",
				"control",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_rsa_provider",
			Category:         "AAA",
			ShortDescription: `Manages ACI RSA Provider`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"rsa",
				"provider",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object RSA Provider.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object RSA Provider.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object RSA Provider.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object RSA Provider.`,
				},
				resource.Attribute{
					Name:        "auth_port",
					Description: `(Optional) Port. Allowed range is "1"-"65535". Default value is "1812". Type: String.`,
				},
				resource.Attribute{
					Name:        "auth_protocol",
					Description: `(Optional) Authentication Protocol. Allowed values are "chap", "mschap", "pap". Default value is "pap". Type: String.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Key. A password for the AAA provider database.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring. Allowed values are "disabled", "enabled". Default value is "disabled". Type: String.`,
				},
				resource.Attribute{
					Name:        "monitoring_password",
					Description: `(Optional) Periodic Server Monitoring Password. Type: String.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username. Default value is "default". Type: String.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Retries. Allowed range is "1"-"5". Default value is "1". Type: String.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout in Seconds. The amount of time between authentication attempts. Allowed range is "1"-"60" and default value is "5". Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_aaa_rs_prov_to_epp",
					Description: `(Optional) Represents the relation to a Relation to AProvider Reachability EPP (class fvAREpP). Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_aaa_rs_sec_prov_to_epg",
					Description: `(Optional) Represents the relation to a Attachable Target Group (class fvATg). A source relation to the endpoint group through which the provider server is reachable. Type: String. ## Importing ## An existing RSAProvider can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_rsa_provider.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_saml_provider",
			Category:         "AAA",
			ShortDescription: `Manages ACI SAML Provider`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"saml",
				"provider",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object SAML Provider.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object SAML Provider.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object SAML Provider. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object SAML Provider. Type: String.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `(Optional) Entity ID. Type: String.`,
				},
				resource.Attribute{
					Name:        "gui_banner_message",
					Description: `(Optional) Gui Redirect Banner Message. Type: String. (Note: Value passed for "https_proxy" attribute should be a URL)`,
				},
				resource.Attribute{
					Name:        "https_proxy",
					Description: `(Optional) Https Proxy to reach IDP's Metadata URL. Type: String.`,
				},
				resource.Attribute{
					Name:        "id_p",
					Description: `(Optional) Identity Provider. Allowed values are "adfs", "okta", "ping identity". Default value is "adfs". Type: String.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Key. A password for the AAA provider database. Type: String.`,
				},
				resource.Attribute{
					Name:        "metadata_url",
					Description: `(Optional) Metadata Url provided by IDP. Type: String.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring. Allowed values are "disabled", "enabled". Default value is "disabled". Type: String.`,
				},
				resource.Attribute{
					Name:        "monitoring_password",
					Description: `(Optional) Periodic Server Monitoring Password. Type: String.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username. Default value is "default". Type: String.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Retries. Allowed range is "1"-"5" and default value is "1". Type: String.`,
				},
				resource.Attribute{
					Name:        "sig_alg",
					Description: `(Optional) Signature Algorithm. Allowed values are "SIG_RSA_SHA1", "SIG_RSA_SHA224", "SIG_RSA_SHA256", "SIG_RSA_SHA384", "SIG_RSA_SHA512". Default value is "SIG_RSA_SHA256". Type: String.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout in Seconds. The amount of time between authentication attempts. Allowed range is "5"-"60" and default value is "5". Type: String.`,
				},
				resource.Attribute{
					Name:        "tp",
					Description: `(Optional) Certificate Authority. Type: String.`,
				},
				resource.Attribute{
					Name:        "want_assertions_encrypted",
					Description: `(Optional) Want Encrypted SAML Assertions. Allowed values are "no" and "yes". Default value is "yes". Type: String.`,
				},
				resource.Attribute{
					Name:        "want_assertions_signed",
					Description: `(Optional) Want Assertions in SAML Response Signed. Allowed values are "no" and "yes". Default value is "yes". Type: String.`,
				},
				resource.Attribute{
					Name:        "want_requests_signed",
					Description: `(Optional) Want SAML Auth Requests Signed. Allowed values are "no" and "yes". Default value is "yes". Type: String.`,
				},
				resource.Attribute{
					Name:        "want_response_signed",
					Description: `(Optional) Want SAML Response Message Signed. Allowed values are "no" and "yes". Default value is "yes". Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_aaa_rs_prov_to_epp",
					Description: `(Optional) Represents the relation to a Relation to AProvider Reachability EPP (class fvAREpP). Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_aaa_rs_sec_prov_to_epg",
					Description: `(Optional) Represents the relation to a Attachable Target Group (class fvATg). A source relation to the endpoint group through which the provider server is reachable. Type: String. ## Importing ## An existing SAMLProvider can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_saml_provider.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_saml_provider_group",
			Category:         "AAA",
			ShortDescription: `Manages ACI SAML Provider Group`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"saml",
				"provider",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object SAML Provider Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object SAML Provider Group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object SAML Provider Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object SAML Provider Group. ## Importing ## An existing SAMLProviderGroup can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_saml_provider_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_service_redirect_policy",
			Category:         "L4-L7 Services",
			ShortDescription: `Manages ACI Service Redirect Policy`,
			Description:      ``,
			Keywords: []string{
				"l4",
				"l7",
				"services",
				"service",
				"redirect",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object Service Redirect Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of Object Service Redirect Policy.`,
				},
				resource.Attribute{
					Name:        "anycast_enabled",
					Description: `(Optional) Anycast enabled for object Service Redirect Policy. NOTE: ` + "`" + `anycast_enabled` + "`" + ` and ` + "`" + `program_local_pod_only` + "`" + ` cannot be "yes" simultaneously. Allowed values: "yes", "no". Default value: "no".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Service Redirect Policy.`,
				},
				resource.Attribute{
					Name:        "dest_type",
					Description: `(Optional) Dest type for object Service Redirect Policy. Allowed values: "L1", "L2", "L3". Default value: "L3".`,
				},
				resource.Attribute{
					Name:        "hashing_algorithm",
					Description: `(Optional) Hashing algorithm for object Service Redirect Policy. Allowed values: "sip", "dip", "sip-dip-prototype", Default value: "sip-dip-prototype".`,
				},
				resource.Attribute{
					Name:        "max_threshold_percent",
					Description: `(Optional) Max threshold_percent for object Service Redirect Policy. Range : 1-100. Default Value: 0.`,
				},
				resource.Attribute{
					Name:        "min_threshold_percent",
					Description: `(Optional) Min threshold_percent for object Service Redirect Policy. Range : 1-100. Default Value: 0.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Service Redirect Policy.`,
				},
				resource.Attribute{
					Name:        "program_local_pod_only",
					Description: `(Optional) Program local pod only for object Service Redirect Policy. Allowed values: "yes", "no". Default value: "no".`,
				},
				resource.Attribute{
					Name:        "resilient_hash_enabled",
					Description: `(Optional) Resilient hash enabled for object Service Redirect Policy. Allowed values: "yes", "no". Default value: "no".`,
				},
				resource.Attribute{
					Name:        "threshold_down_action",
					Description: `(Optional) Threshold down the action for object Service Redirect Policy. Allowed values: "bypass", "deny", "permit". Default value: "permit".`,
				},
				resource.Attribute{
					Name:        "threshold_enable",
					Description: `(Optional) Threshold enable for object Service Redirect Policy. Allowed values: "yes", "no". Default value: "no".`,
				},
				resource.Attribute{
					Name:        "relation_vns_rs_ipsla_monitoring_pol",
					Description: `(Optional) Relation to class fvIPSLAMonitoringPol. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Service Redirect Policy. ## Importing ## An existing Service Redirect Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_service_redirect_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_snmp_community",
			Category:         "Resources",
			ShortDescription: `Manages ACI SNMP Community`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "parent_dn",
					Description: `(Required) Distinguished name of the parent object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the SNMP Community.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the SNMP Community.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the SNMP Community. ## Importing ## An existing SNMPCommunity can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_snmp_community.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_span_destination_group",
			Category:         "Monitoring",
			ShortDescription: `Manages ACI SPAN Destination Group`,
			Description:      ``,
			Keywords: []string{
				"monitoring",
				"span",
				"destination",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_span_source_group",
			Category:         "Monitoring",
			ShortDescription: `Manages ACI SPAN Source Group`,
			Description:      ``,
			Keywords: []string{
				"monitoring",
				"span",
				"source",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_span_sourcedestination_group_match_label",
			Category:         "Monitoring",
			ShortDescription: `Manages ACI SPAN Source-destination Group Match Label`,
			Description:      ``,
			Keywords: []string{
				"monitoring",
				"span",
				"sourcedestination",
				"group",
				"match",
				"label",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "span_source_group_dn",
					Description: `(Required) Distinguished name of parent SPANSourceGroup object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object SPANSourceGroup object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object SPANSourceGroup object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object SPANSourceGroup object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias for object SPANSourceGroup object.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) label color. Allowed values: "black", "navy", "dark-blue", "medium-blue", "blue", "dark-green", "green", "teal", "dark-cyan", "deep-sky-blue", "dark-turquoise", "medium-spring-green", "lime", "spring-green", "aqua", "cyan", "midnight-blue", "dodger-blue", "light-sea-green", "forest-green", "sea-green", "dark-slate-gray", "lime-green", "medium-sea-green", "turquoise", "royal-blue", "steel-blue", "dark-slate-blue", "medium-turquoise", "indigo", "dark-olive-green", "cadet-blue", "cornflower-blue", "medium-aquamarine", "dim-gray", "slate-blue", "olive-drab", "slate-gray", "light-slate-gray", "medium-slate-blue", "lawn-green", "chartreuse", "aquamarine", "maroon", "purple", "olive", "gray", "sky-blue", "light-sky-blue", "blue-violet", "dark-red", "dark-magenta", "saddle-brown", "dark-sea-green", "light-green", "medium-purple", "dark-violet", "pale-green", "dark-orchid", "yellow-green", "sienna", "brown", "dark-gray", "light-blue", "green-yellow", "pale-turquoise", "light-steel-blue", "powder-blue", "fire-brick", "dark-goldenrod", "medium-orchid", "rosy-brown", "dark-khaki", "silver", "medium-violet-red", "indian-red", "peru", "chocolate", "tan", "light-gray", "thistle", "orchid", "goldenrod", "pale-violet-red", "crimson", "gainsboro", "plum", "burlywood", "light-cyan", "lavender", "dark-salmon", "violet", "pale-goldenrod", "light-coral", "khaki", "alice-blue", "honeydew", "azure", "sandy-brown", "wheat", "beige", "white-smoke", "mint-cream", "ghost-white", "salmon", "antique-white", "linen", "light-goldenrod-yellow", "old-lace", "red", "fuchsia", "magenta", "deep-pink", "orange-red", "tomato", "hot-pink", "coral", "dark-orange", "light-salmon", "orange", "light-pink", "pink", "gold", "peachpuff", "navajo-white", "moccasin", "bisque", "misty-rose", "blanched-almond", "papaya-whip", "lavender-blush", "seashell", "cornsilk", "lemon-chiffon", "floral-white", "snow", "yellow", "light-yellow", "ivory", "white" ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the SPAN Source-destination Group Match Label. ## Importing ## An existing SPAN Source-destination Group Match Label can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: <https://www.terraform.io/docs/import/index.html> ` + "`" + `` + "`" + `` + "`" + `bash terraform import aci_span_sourcedestination_group_match_label.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spanning_tree_interface_policy",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Spanning Tree Interface Policy`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"spanning",
				"tree",
				"interface",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_access_port_selector",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Spine Access Port Selector`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"spine",
				"port",
				"selector",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "spine_interface_profile_dn",
					Description: `(Required) Distinguished name of the parent Spine Interface Profile.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Spine Access Port Selector.`,
				},
				resource.Attribute{
					Name:        "spine_access_port_selector_type",
					Description: `(Required) The type of Spine Access Port Selector. Allowed values are "ALL" and "range". Default is "ALL". The "range" can be specified with the resource "aci_access_port_block".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Spine Access Port Selector.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Spine Access Port Selector.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_sp_acc_grp",
					Description: `(Optional) Represents the relation to a Spine Access Group (class infraSpAccGrp). Type: String. ## Importing ## An existing Spine Access Port Selector can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_spine_access_port_selector.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_interface_profile",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Spine Interface Profile`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"spine",
				"interface",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object Spine interface profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for Object Spine interface profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for Object Spine interface profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for Object Spine interface profile. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Spine Interface Profile. ## Importing ## An existing Spine Interface Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_spine_interface_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_interface_profile_selector",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Spine Interface Profile selector`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"spine",
				"interface",
				"profile",
				"selector",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "spine_profile_dn",
					Description: `(Required) Distinguished name of parent Spine Profile.`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Required) tDn of the Spine Interface Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for Spine Interface Profile selector. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Spine Interface Profile selector. ## Importing ## An existing Spine Interface Profile selector can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_spine_interface_profile_selector.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_port_policy_group",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Spine Port Policy Group`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"spine",
				"port",
				"policy",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object Spine Port Policy Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Spine Port Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Spine Port Policy Group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Spine Port Policy Group.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_h_if_pol",
					Description: `(Optional) Relation to class fabricHIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_cdp_if_pol",
					Description: `(Optional) Relation to class cdpIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_copp_if_pol",
					Description: `(Optional) Relation to class coppIfPol. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_att_ent_p",
					Description: `(Optional) Relation to class infraAttEntityP. Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_macsec_if_pol",
					Description: `(Optional) Relation to class macsecIfPol. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Spine Access Port Policy Group. ## Importing ## An existing Spine Access Port Policy Group can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_spine_port_policy_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_port_selector",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Spine Port selector`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"spine",
				"port",
				"selector",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "spine_profile_dn",
					Description: `(Required) Distinguished name of parent Spine Profile.`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Required) tDn of the Spine Interface Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for Spine Port selector. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Spine Port selector. ## Importing ## An existing Spine port selector can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_spine_interface_profile_selector.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_profile",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Spine Profile`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"spine",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object Spine Profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Spine Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Spine Profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Spine Profile. - ` + "`" + `spine_selector` + "`" + ` - (Optional) Spine Selector block to attach with the Spine Profile. - ` + "`" + `spine_selector.name` + "`" + ` - (Required) Name of the Spine Selector. - ` + "`" + `spine_selector.switch_association_type` + "`" + ` - (Required) Type of switch association. Allowed values: "ALL", "range", "ALL_IN_POD" - ` + "`" + `spine_selector.node_block` + "`" + ` - (Optional) Node block to attach with Spine Selector. - ` + "`" + `spine_selector.node_block.name` + "`" + ` - (Required) Name of the node block. - ` + "`" + `spine_selector.node_block.from_` + "`" + ` - (Optional) Start of Node Block range. Range from 1 to 16000. Default value is "1". - ` + "`" + `spine_selector.node_block.to_` + "`" + ` - (Optional) End of Node Block range. Range from 1 to 16000. Default value is "1".`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_sp_acc_port_p",
					Description: `(Optional) Relation to class infraSpAccPortP. Cardinality - N_TO_M. Type - Set of String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Spine Profile. ## Importing ## An existing Spine Profile can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_spine_profile.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_switch_association",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Spine Switch Association`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"spine",
				"switch",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "spine_profile_dn",
					Description: `(Required) Distinguished name of parent Spine Profile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object Spine Switch association.`,
				},
				resource.Attribute{
					Name:        "spine_switch_association_type",
					Description: `(Required) Spine association type of Object Spine Switch Association. Allowed values: "ALL", "range", "ALL_IN_POD"`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Spine Switch Association.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Spine Switch Association.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Spine Switch Association.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_spine_acc_node_p_grp",
					Description: `(Optional) Relation to class infraSpineAccNodePGrp. Cardinality - N_TO_ONE. Type - String. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the Spine Association. ## Importing ## An existing Spine Association can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_spine_association.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_switch_policy_group",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Spine Switch Policy Group`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"spine",
				"switch",
				"policy",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object Spine Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Spine Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_iacl_spine_profile",
					Description: `(Optional) Represents the relation to a CoPP Prefilter Profile for Spines (class iaclSpineProfile). Relationship the CoPP Prefilter Spine profile to be applied on spines Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_spine_bfd_ipv4_inst_pol",
					Description: `(Optional) Represents the relation to a BFD Ipv4 Instance Policy (class bfdIpv4InstPol). Relationship to BFD Ipv4 Instance Policy Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_spine_bfd_ipv6_inst_pol",
					Description: `(Optional) Represents the relation to a BFD Ipv6 Instance Policy (class bfdIpv6InstPol). Relationship to BFD Ipv6 Instance Policy Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_spine_copp_profile",
					Description: `(Optional) Represents the relation to a CoPP Profile for Spines (class coppSpineProfile). Relationship the CoPP profile to be applied on spines Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_spine_p_grp_to_cdp_if_pol",
					Description: `(Optional) Represents the relation to a Relation to cdp interface policy for mgmt port (class cdpIfPol). Relationship to cdp interface policy for mgmt port Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_infra_rs_spine_p_grp_to_lldp_if_pol",
					Description: `(Optional) Represents the relation to a Relation to lldp interface policy for mgmt port (class lldpIfPol). Relationship to lldp interface policy for mgmt port Type: String.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Spine Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Spine Switch Policy Group. ## Importing ## An existing Spine Switch Policy Group can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_spine_switch_policy_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_static_node_mgmt_address",
			Category:         "Node Management",
			ShortDescription: `Manages ACI Management Static Node`,
			Description:      ``,
			Keywords: []string{
				"node",
				"management",
				"static",
				"mgmt",
				"address",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_subnet",
			Category:         "Networking",
			ShortDescription: `Manages ACI Subnet`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"subnet",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_taboo_contract",
			Category:         "Contract",
			ShortDescription: `Manages ACI Taboo Contract`,
			Description:      ``,
			Keywords: []string{
				"contract",
				"taboo",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tacacs_accounting",
			Category:         "AAA",
			ShortDescription: `Manages ACI TACACS Accounting`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"tacacs",
				"accounting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object TACACS Accounting. Type: String.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object TACACS Accounting. Type: String.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object TACACS Accounting. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object TACACS Accounting. Type: String. ## Importing ## An existing TACACSAccounting can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_tacacs_accounting.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tacacs_accounting_destination",
			Category:         "AAA",
			ShortDescription: `Manages ACI TACACS Accounting Destination`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"tacacs",
				"accounting",
				"destination",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tacacs_accounting_dn",
					Description: `(Required) Distinguished name of parent TACACS Accounting object..`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) Host or IP address of object TACACS Accounting Destination.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port of object TACACS Accounting Destination. Allowed Range: "1" - "65535". Default value: "49".`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object TACACS Accounting Destination.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object TACACS Accounting Destination.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of object TACACS Accounting Destination.`,
				},
				resource.Attribute{
					Name:        "auth_protocol",
					Description: `(Optional) Authentication Protocol of object TACACS Accounting Destination. Allowed values are "chap", "mschap" and "pap". Default value is "pap". Type: String.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key or password used to uniquely identify object TACACS Accounting Destination.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object TACACS Accounting Destination.`,
				},
				resource.Attribute{
					Name:        "relation_file_rs_a_remote_host_to_epg",
					Description: `(Optional) Represents the relation to a Attachable Target Group (class fvATg). A source relation to the endpoint group through which the remote host is reachable. Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_file_rs_a_remote_host_to_epp",
					Description: `(Optional) Represents the relation to a Relation to Remote Host Reachability EPP (class fvAREpP). A source relation to the abstract representation of the resolvable endpoint profile. Type: String. ## Importing ## An existing TACACS Accounting Destination can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_tacacs_accounting_destination.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tacacs_provider",
			Category:         "AAA",
			ShortDescription: `Manages ACI TACACS Provider`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"tacacs",
				"provider",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object TACACS Provider.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object TACACS Provider.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object TACACS Provider. Type: String.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object TACACS Provider. Type: String.`,
				},
				resource.Attribute{
					Name:        "auth_protocol",
					Description: `(Optional) TACACS Authentication Protocol. The TACACS authentication protocol. Allowed values are "chap", "mschap", "pap". Default value is "pap". Type: String.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Key. A password for the AAA provider database. Type: String.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring. Allowed values are "disabled", "enabled". Default value is "disabled". Type: String.`,
				},
				resource.Attribute{
					Name:        "monitoring_password",
					Description: `(Optional) Periodic Server Monitoring Password. Type: String.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username. Default value is "default". Type: String.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port. The service port number for the TACACS service. Allowed range is "1"-"65535". Default value is "49". Type: String.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Retries. Allowed range is "1"-"5" and default value is "1". Type: String.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout in Seconds. The timeout for communication with the TACACS provider server. Allowed range is "1"-"60". Default value is "5". Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_aaa_rs_prov_to_epp",
					Description: `(Optional) Represents the relation to a Relation to AProvider Reachability EPP (class fvAREpP). Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_aaa_rs_sec_prov_to_epg",
					Description: `(Optional) Represents the relation to a Attachable Target Group (class fvATg). A source relation to the endpoint group through which the provider server is reachable. Type: String. ## Importing ## An existing TACACSProvider can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_tacacs_provider.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tacacs_provider_group",
			Category:         "AAA",
			ShortDescription: `Manages ACI TACACS + Provider Group`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"tacacs",
				"provider",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object TACACS + Provider Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object TACACS + Provider Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object TACACS + Provider Group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of object TACACS + Provider Group. ## Importing ## An existing TACACS + ProviderGroup can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_tacacs_provider_group.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tacacs_source",
			Category:         "AAA",
			ShortDescription: `Manages ACI TACACS Source`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"tacacs",
				"source",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object TACACS Source.`,
				},
				resource.Attribute{
					Name:        "parent_dn",
					Description: `(Optional) Distinguished name of parent object of TACACS Source. Default value is "uni/fabric/moncommon". Type: String.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object TACACS Source.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object TACACS Source. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object TACACS Source. Type: String.`,
				},
				resource.Attribute{
					Name:        "incl",
					Description: `(Optional) Include Action. The information to include for the call home source. Allowed values are "audit", "events", "faults" and "session". Default value is ["audit","session"]. Type: List.`,
				},
				resource.Attribute{
					Name:        "min_sev",
					Description: `(Optional) minSev. Allowed values are "cleared", "critical", "info", "major", "minor" and "warning". Default value is "info". Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_tacacs_rs_dest_group",
					Description: `(Optional) Represents the relation to a TACACS Destination Group (class tacacsGroup). Type: String. ## Importing ## An existing TACACSSource can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_tacacs_source.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tag",
			Category:         "Resources",
			ShortDescription: `Manages ACI Tag`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "parent_dn",
					Description: `(Required) Distinguished name of the object to which the Tag is attached to.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the Tag.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the Tag. ## Importing ## An existing Tag can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_tag.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tenant",
			Category:         "Application Management",
			ShortDescription: `Manages ACI Tenant`,
			Description:      ``,
			Keywords: []string{
				"application",
				"management",
				"tenant",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_trigger_scheduler",
			Category:         "Scheduler",
			ShortDescription: `Manages ACI Trigger Scheduler`,
			Description:      ``,
			Keywords: []string{
				"scheduler",
				"trigger",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_user_security_domain",
			Category:         "AAA",
			ShortDescription: `Manages ACI User Security Domain`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"user",
				"security",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "local_user_dn",
					Description: `(Required) Distinguished name of parent LocalUser object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object User Security Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object User Security Domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object User Security Domain.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object User Security Domain. ## Importing ## An existing User Security Domain can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_user_security_domain.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_user_security_domain_role",
			Category:         "AAA",
			ShortDescription: `Manages ACI User Security Domain Role`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"user",
				"security",
				"domain",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_domain_dn",
					Description: `(Required) Distinguished name of parent UserDomain object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object User Security Domain Role.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object User Security Domain Role. Type: String.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object User Security Domain Role. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object User Security Domain Role. Type: String.`,
				},
				resource.Attribute{
					Name:        "priv_type",
					Description: `(Optional) Privilege Type.The privilege type for a user role. Allowed values are "readPriv", "writePriv". Default value is "readPriv". Type: String. ## Importing ## An existing User Security Domain Role can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_user_security_domain_role.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vlan_encapsulationfor_vxlan_traffic",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI Vlan Encapsulation for Vxlan Traffic`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"vlan",
				"encapsulationfor",
				"vxlan",
				"traffic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "attachable_access_entity_profile_dn",
					Description: `(Required) Distinguished name of parent AttachableAccessEntityProfile object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object vlan encapsulation for vxlan traffic.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object vlan encapsulation for vxlan traffic.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vlan_pool",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI VLAN Pool`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"vlan",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object vlan pool.`,
				},
				resource.Attribute{
					Name:        "alloc_mode",
					Description: `(Required) Allocation mode for object vlan_pool. Allowed values: "dynamic", "static"`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object vlan pool.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object vlan pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object vlan pool. ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the VLAN Pool. ## Importing ## An existing VLAN Pool can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: <https://www.terraform.io/docs/import/index.html> ` + "`" + `` + "`" + `` + "`" + `bash terraform import aci_vlan_pool.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vmm_controller",
			Category:         "Virtual Networking",
			ShortDescription: `Manages ACI VMM Controller`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"networking",
				"vmm",
				"controller",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vmm_credential",
			Category:         "Virtual Networking",
			ShortDescription: `Manages ACI VMM Credential`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"networking",
				"vmm",
				"credential",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vmm_domain_dn",
					Description: `(Required) Distinguished name of parent VMM Domain object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object VMM Credential.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object VMM Credential.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object VMM Credential.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of object VMM Credential.`,
				},
				resource.Attribute{
					Name:        "pwd",
					Description: `(Optional) Password.`,
				},
				resource.Attribute{
					Name:        "usr",
					Description: `(Optional) Username. Min length is "0". Max length is "128". If any value is assigned to a username then it cannot be updated to an empty value. ## Importing ## An existing VMMCredential can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vmm_credential.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vmm_domain",
			Category:         "Virtual Networking",
			ShortDescription: `Manages ACI VMM Domain`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"networking",
				"vmm",
				"domain",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vpc_domain_policy",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI VPC Domain Policy`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"vpc",
				"domain",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object VPC Domain Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object VPC Domain Policy.`,
				},
				resource.Attribute{
					Name:        "dead_intvl",
					Description: `(Optional) The VPC peer dead interval time of object VPC Domain Policy. Range: "5" - "600". Default value is "200".`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object VPC Domain Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object VPC Domain Policy. ## Importing ## An existing VPC Domain Policy can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_vpc_domain_policy.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vpc_explicit_protection_group",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI VPC Explicit Protection Group`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"vpc",
				"explicit",
				"protection",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vrf",
			Category:         "Networking",
			ShortDescription: `Manages ACI VRF`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"vrf",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "af",
					Description: `(Required) The BGP address family. Allowed values are "ipv4-ucast", "ipv6-ucast", and default value is "ipv4-ucast". Type: String.`,
				},
				resource.Attribute{
					Name:        "tn_bgp_ctx_af_pol_name",
					Description: `(Required) Distinguished name (DN) of the BGP address family context policy object. Type: String. Note: In the APIC GUI,a VRF (fvCtx) was called a "Context"or "PrivateNetwork." ## Attribute Reference ## The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the VRF. ## Importing An existing VRF can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: <https://www.terraform.io/docs/import/index.html> ` + "`" + `` + "`" + `` + "`" + `bash terraform import aci_vrf.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vrf_snmp_context",
			Category:         "Networking",
			ShortDescription: `Manages ACI VRF SNMP Context`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"vrf",
				"snmp",
				"context",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrf_dn",
					Description: `(Required) Distinguished name of parent VRF object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object VRF SNMP Context`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object VRF SNMP Context.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object VRF SNMP Context. ## Importing ## An existing VRF SNMP Context can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_vrf_snmp_context.example <Dn> ` + "`" + `` + "`" + `` + "`" + ` ## NOTE ## User can create only one VRF SNMP Context under one VRF.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vrf_snmp_context_community",
			Category:         "Networking",
			ShortDescription: `Manages ACI SNMP Community`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"vrf",
				"snmp",
				"context",
				"community",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrf_snmp_context_dn",
					Description: `(Required) Distinguished name of parent VRF SNMP Context object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object SNMP Community.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object SNMP Community.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object SNMP Community.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of object SNMP Community. ## Importing ## An existing VRF SNMP Context Community can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_vrf_snmp_context_community.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vrf_to_bgp_address_family_context",
			Category:         "Networking",
			ShortDescription: `Manages the ACI Relationship object between VRF and the BGP Address Family Context Policy`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"vrf",
				"to",
				"bgp",
				"address",
				"family",
				"context",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrf_dn",
					Description: `(Required) Distinguished name (DN) of the parent VRF object.`,
				},
				resource.Attribute{
					Name:        "bgp_address_family_context_dn",
					Description: `(Required) Distinguished name (DN) of the BGP address family context policy object.`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `(Required) The BGP address family. Allowed values are "ipv4-ucast", "ipv6-ucast", and default value is "ipv4-ucast". Type: String.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the VRF to BGP address family context policy relationship object. ## Importing ## An existing VRF to BGP address family context policy relationship object can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_vrf_to_bgp_address_family_context.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vsan_pool",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI VSAN Pool`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"vsan",
				"pool",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vswitch_policy",
			Category:         "Virtual Networking",
			ShortDescription: `Manages ACI VSwitch Policy`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"networking",
				"vswitch",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vxlan_pool",
			Category:         "Access Policies",
			ShortDescription: `Manages ACI VXLAN Pool`,
			Description:      ``,
			Keywords: []string{
				"access",
				"policies",
				"vxlan",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object vxlan pool.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object vxlan pool.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object vxlan pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object vxlan pool. ## Attribute Reference The only attribute that this resource exports is the ` + "`" + `id` + "`" + `, which is set to the Dn of the VXLAN Pool. ## Importing ## An existing VXLAN Pool can be [imported][docs-import] into this resource via its Dn, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import aci_vxlan_pool.example <Dn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_x509_certificate",
			Category:         "AAA",
			ShortDescription: `Manages ACI X509 Certificate`,
			Description:      ``,
			Keywords: []string{
				"aaa",
				"x509",
				"certificate",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"aci_aaa_domain":                               0,
		"aci_aaa_domain_relationship":                  1,
		"aci_aaep_to_domain":                           2,
		"aci_access_generic":                           3,
		"aci_access_group":                             4,
		"aci_access_port_block":                        5,
		"aci_access_port_selector":                     6,
		"aci_access_sub_port_block":                    7,
		"aci_access_switch_policy_group":               8,
		"aci_action_rule_additional_communities":       9,
		"aci_action_rule_profile":                      10,
		"aci_annotation":                               11,
		"aci_any":                                      12,
		"aci_application_epg":                          13,
		"aci_application_profile":                      14,
		"aci_attachable_access_entity_profile":         15,
		"aci_authentication_properties":                16,
		"aci_bd_dhcp_label":                            17,
		"aci_bfd_interface_policy":                     18,
		"aci_bgp_address_family_context":               19,
		"aci_bgp_best_path_policy":                     20,
		"aci_bgp_peer_connectivity_profile":            21,
		"aci_bgp_peer_prefix":                          22,
		"aci_bgp_route_control_profile":                23,
		"aci_bgp_route_summarization":                  24,
		"aci_bgp_timers":                               25,
		"aci_bridge_domain":                            26,
		"aci_cdp_interface_policy":                     27,
		"aci_cloud_applicationcontainer":               28,
		"aci_cloud_aws_provider":                       29,
		"aci_cloud_cidr_pool":                          30,
		"aci_cloud_context_profile":                    31,
		"aci_cloud_domain_profile":                     32,
		"aci_cloud_endpoint_selector":                  33,
		"aci_cloud_endpoint_selectorfor_external_epgs": 34,
		"aci_cloud_epg":                                35,
		"aci_cloud_external_epg":                       36,
		"aci_cloud_provider_profile":                   37,
		"aci_cloud_subnet":                             38,
		"aci_cloud_vpn_gateway":                        39,
		"aci_configuration_export_policy":              40,
		"aci_configuration_import_policy":              41,
		"aci_connection":                               42,
		"aci_console_authentication":                   43,
		"aci_contract":                                 44,
		"aci_contract_subject":                         45,
		"aci_coop_policy":                              46,
		"aci_default_authentication":                   47,
		"aci_destination_of_redirected_traffic":        48,
		"aci_dhcp_option_policy":                       49,
		"aci_dhcp_relay_policy":                        50,
		"aci_duo_provider_group":                       51,
		"aci_encryption_key":                           52,
		"aci_end_point_retention_policy":               53,
		"aci_endpoint_controls":                        54,
		"aci_endpoint_ip_aging_profile":                55,
		"aci_endpoint_loop_protection":                 56,
		"aci_endpoint_security_group":                  57,
		"aci_endpoint_security_group_epg_selector":     58,
		"aci_endpoint_security_group_selector":         59,
		"aci_endpoint_security_group_tag_selector":     60,
		"aci_epg_to_contract":                          61,
		"aci_epg_to_contract_interface":                62,
		"aci_epg_to_domain":                            63,
		"aci_epg_to_static_path":                       64,
		"aci_epgs_using_function":                      65,
		"aci_error_disable_recovery":                   66,
		"aci_external_network_instance_profile":        67,
		"aci_fabric_if_pol":                            68,
		"aci_fabric_node_control":                      69,
		"aci_fabric_node_member":                       70,
		"aci_fabric_wide_settings":                     71,
		"aci_fc_domain":                                72,
		"aci_fex_bundle_group":                         73,
		"aci_fex_profile":                              74,
		"aci_file_remote_path":                         75,
		"aci_filter":                                   76,
		"aci_filter_entry":                             77,
		"aci_firmware_download_task":                   78,
		"aci_firmware_group":                           79,
		"aci_firmware_policy":                          80,
		"aci_function_node":                            81,
		"aci_global_security":                          82,
		"aci_hsrp_group_policy":                        83,
		"aci_hsrp_interface_policy":                    84,
		"aci_imported_contract":                        85,
		"aci_interface_blacklist":                      86,
		"aci_interface_fc_policy":                      87,
		"aci_isis_domain_policy":                       88,
		"aci_l2_domain":                                89,
		"aci_l2_interface_policy":                      90,
		"aci_l2_outside":                               91,
		"aci_l2out_extepg":                             92,
		"aci_l3_domain_profile":                        93,
		"aci_l3_ext_subnet":                            94,
		"aci_l3_interface_policy":                      95,
		"aci_l3_outside":                               96,
		"aci_l3out_bfd_interface_profile":              97,
		"aci_l3out_bgp_external_policy":                98,
		"aci_l3out_bgp_protocol_profile":               99,
		"aci_l3out_floating_svi":                       100,
		"aci_l3out_hsrp_interface_group":               101,
		"aci_l3out_hsrp_interface_profile":             102,
		"aci_l3out_hsrp_secondary_vip":                 103,
		"aci_l3out_loopback_interface_profile":         104,
		"aci_l3out_ospf_external_policy":               105,
		"aci_l3out_ospf_interface_profile":             106,
		"aci_l3out_path_attachment":                    107,
		"aci_l3out_path_attachment_secondary_ip":       108,
		"aci_l3out_route_tag_policy":                   109,
		"aci_l3out_static_route":                       110,
		"aci_l3out_static_route_next_hop":              111,
		"aci_l3out_vpc_member":                         112,
		"aci_l4_l7_service_graph_template":             113,
		"aci_lacp_policy":                              114,
		"aci_ldap_group_map":                           115,
		"aci_ldap_group_map_rule":                      116,
		"aci_ldap_group_map_rule_to_group_map":         117,
		"aci_ldap_provider":                            118,
		"aci_leaf_access_bundle_policy_group":          119,
		"aci_leaf_access_port_policy_group":            120,
		"aci_leaf_breakout_port_group":                 121,
		"aci_leaf_interface_profile":                   122,
		"aci_leaf_profile":                             123,
		"aci_leaf_selector":                            124,
		"aci_lldp_interface_policy":                    125,
		"aci_local_user":                               126,
		"aci_logical_device_context":                   127,
		"aci_logical_interface_context":                128,
		"aci_logical_interface_profile":                129,
		"aci_logical_node_profile":                     130,
		"aci_logical_node_to_fabric_node":              131,
		"aci_login_domain":                             132,
		"aci_login_domain_provider":                    133,
		"aci_maintenance_group_node":                   134,
		"aci_maintenance_policy":                       135,
		"aci_managed_node_connectivity_group":          136,
		"aci_match_community_terms":                    137,
		"aci_match_regex_community_terms":              138,
		"aci_match_route_destination_rule":             139,
		"aci_match_rule":                               140,
		"aci_mcp_instance_policy":                      141,
		"aci_mgmt_preference":                          142,
		"aci_mgmt_zone":                                143,
		"aci_miscabling_protocol_interface_policy":     144,
		"aci_monitoring_policy":                        145,
		"aci_node_block":                               146,
		"aci_node_block_firmware":                      147,
		"aci_node_mgmt_epg":                            148,
		"aci_ospf_interface_policy":                    149,
		"aci_ospf_route_summarization":                 150,
		"aci_ospf_timers":                              151,
		"aci_physical_domain":                          152,
		"aci_pod_maintenance_group":                    153,
		"aci_port_security_policy":                     154,
		"aci_port_tracking":                            155,
		"aci_qos_instance_policy":                      156,
		"aci_radius_provider":                          157,
		"aci_radius_provider_group":                    158,
		"aci_ranges":                                   159,
		"aci_recurring_window":                         160,
		"aci_rest":                                     161,
		"aci_rest_managed":                             162,
		"aci_route_control_context":                    163,
		"aci_route_control_profile":                    164,
		"aci_rsa_provider":                             165,
		"aci_saml_provider":                            166,
		"aci_saml_provider_group":                      167,
		"aci_service_redirect_policy":                  168,
		"aci_snmp_community":                           169,
		"aci_span_destination_group":                   170,
		"aci_span_source_group":                        171,
		"aci_span_sourcedestination_group_match_label": 172,
		"aci_spanning_tree_interface_policy":           173,
		"aci_spine_access_port_selector":               174,
		"aci_spine_interface_profile":                  175,
		"aci_spine_interface_profile_selector":         176,
		"aci_spine_port_policy_group":                  177,
		"aci_spine_port_selector":                      178,
		"aci_spine_profile":                            179,
		"aci_spine_switch_association":                 180,
		"aci_spine_switch_policy_group":                181,
		"aci_static_node_mgmt_address":                 182,
		"aci_subnet":                                   183,
		"aci_taboo_contract":                           184,
		"aci_tacacs_accounting":                        185,
		"aci_tacacs_accounting_destination":            186,
		"aci_tacacs_provider":                          187,
		"aci_tacacs_provider_group":                    188,
		"aci_tacacs_source":                            189,
		"aci_tag":                                      190,
		"aci_tenant":                                   191,
		"aci_trigger_scheduler":                        192,
		"aci_user_security_domain":                     193,
		"aci_user_security_domain_role":                194,
		"aci_vlan_encapsulationfor_vxlan_traffic":      195,
		"aci_vlan_pool":                                196,
		"aci_vmm_controller":                           197,
		"aci_vmm_credential":                           198,
		"aci_vmm_domain":                               199,
		"aci_vpc_domain_policy":                        200,
		"aci_vpc_explicit_protection_group":            201,
		"aci_vrf":                                      202,
		"aci_vrf_snmp_context":                         203,
		"aci_vrf_snmp_context_community":               204,
		"aci_vrf_to_bgp_address_family_context":        205,
		"aci_vsan_pool":                                206,
		"aci_vswitch_policy":                           207,
		"aci_vxlan_pool":                               208,
		"aci_x509_certificate":                         209,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
