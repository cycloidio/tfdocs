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
					Description: `(Required) Name of object aaa domain. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the aaa domain.`,
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
					Description: `(Optional) Name alias for object aaa domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the aaa domain.`,
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
					Description: `(Optional) Name alias for object aaa domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_aaa_domain_relationship",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI AAA Domain Relationship for Parent Object`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "parent_dn",
					Description: `(Required) Distinguished name of parent object.`,
				},
				resource.Attribute{
					Name:        "aaa_domain_dn",
					Description: `(Required) Distinguished name of the AAA Security Domain for Parent Object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the AAA Security Domain for Parent Object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object AAA Security Domain for Parent Object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object AAA Security Domain for Parent Object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the AAA Security Domain for Parent Object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object AAA Security Domain for Parent Object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object AAA Security Domain for Parent Object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_aaep_to_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Attachable Access Entity Profile (AAEP) to domain (VMM, Physical or External domain) relationships.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "attachable_access_entity_profile_dn",
					Description: `(Required) Distinguished name of the parent Attachable Access Entity Profile object.`,
				},
				resource.Attribute{
					Name:        "domain_dn",
					Description: `(Required) The Distinguished name of the domain object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Attachable AccessEntity Profile to Domain Relationship object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Attachable AccessEntity Profile to Domain Relationship object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Attachable AccessEntity Profile to Domain Relationship object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Attachable AccessEntity Profile to Domain Relationship object.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Access Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_port_block",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Access Port Block`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_port_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Access Port Selector`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Description: `(Required) Distinguished name of parent Access Port Selector or Spine Access Port Selector object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object access sub port block. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Access Sub Port Block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object access sub port block.`,
				},
				resource.Attribute{
					Name:        "from_card",
					Description: `(Optional) From card`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `(Optional) Port block from port`,
				},
				resource.Attribute{
					Name:        "from_sub_port",
					Description: `(Optional) From sub port for object access sub port block.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object access sub port block.`,
				},
				resource.Attribute{
					Name:        "to_card",
					Description: `(Optional) To card`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `(Optional) To port`,
				},
				resource.Attribute{
					Name:        "to_sub_port",
					Description: `(Optional) To sub port for object access sub port block.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Access Sub Port Block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object access sub port block.`,
				},
				resource.Attribute{
					Name:        "from_card",
					Description: `(Optional) From card`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `(Optional) Port block from port`,
				},
				resource.Attribute{
					Name:        "from_sub_port",
					Description: `(Optional) From sub port for object access sub port block.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object access sub port block.`,
				},
				resource.Attribute{
					Name:        "to_card",
					Description: `(Optional) To card`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `(Optional) To port`,
				},
				resource.Attribute{
					Name:        "to_sub_port",
					Description: `(Optional) To sub port for object access sub port block.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_access_switch_policy_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Access Switch Policy Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object Access Switch Policy Group. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Access Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Access Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Access Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Access Switch Policy Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Access Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Access Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Access Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Access Switch Policy Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_action_rule_additional_communities",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Action Rule Profile Set Additional Communities`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action_rule_profile_dn",
					Description: `(Required) Distinguished name of the parent action rule profile object.`,
				},
				resource.Attribute{
					Name:        "community",
					Description: `(Required) The community value of the set action rule profile additional communities object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the additional communities object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the additional communities object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the set action rule profile additional communities object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the additional communities object.`,
				},
				resource.Attribute{
					Name:        "set_criteria",
					Description: `(Optional) The criteria for setting the (extended) community attribute for a BGP route update.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the additional communities object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the additional communities object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the set action rule profile additional communities object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the additional communities object.`,
				},
				resource.Attribute{
					Name:        "set_criteria",
					Description: `(Optional) The criteria for setting the (extended) community attribute for a BGP route update.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_action_rule_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for the ACI Action Rule Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of the parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Action Rule Profile object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_route_tag",
					Description: `(Optional) Set Route Tag of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_preference",
					Description: `(Optional) Set Preference of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_weight",
					Description: `(Optional) Set Weight of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_metric",
					Description: `(Optional) Set Metric of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_metric_type",
					Description: `(Optional) Set Metric Type of the Action Rule Profile object.`,
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
					Description: `(Optional) Criteria of the Set Communities object.`,
				},
				resource.Attribute{
					Name:        "community",
					Description: `(Optional) Community of the Set Communities object.`,
				},
				resource.Attribute{
					Name:        "next_hop_propagation",
					Description: `(Optional) Next Hop Propagation of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "multipath",
					Description: `(Optional) Multipath of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_as_path_prepend_last_as",
					Description: `(Optional) Number of ASN prepended to AS Path of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_as_path_prepend_as",
					Description: `(Optional) A block representing ASNs configured as Set As Path - Prepend AS of the Action Rule Profile object. Type: Block.`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `ASN prepended to Set AS Path.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Order in which the ASN is prepended to Set AS Path.`,
				},
				resource.Attribute{
					Name:        "set_dampening",
					Description: `(Optional) A block representing the attributes of Set Dampening object. Type: Block.`,
				},
				resource.Attribute{
					Name:        "half_life",
					Description: `Half Life of the Set Dampening object.`,
				},
				resource.Attribute{
					Name:        "reuse",
					Description: `Reuse Limit of the Set Dampening object.`,
				},
				resource.Attribute{
					Name:        "suppress",
					Description: `Suppress Limit of the Set Dampening object.`,
				},
				resource.Attribute{
					Name:        "max_suppress_time",
					Description: `Max Suppress Time of the Set Dampening object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_route_tag",
					Description: `(Optional) Set Route Tag of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_preference",
					Description: `(Optional) Set Preference of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_weight",
					Description: `(Optional) Set Weight of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_metric",
					Description: `(Optional) Set Metric of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_metric_type",
					Description: `(Optional) Set Metric Type of the Action Rule Profile object.`,
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
					Description: `(Optional) Criteria of the Set Communities object.`,
				},
				resource.Attribute{
					Name:        "community",
					Description: `(Optional) Community of the Set Communities object.`,
				},
				resource.Attribute{
					Name:        "next_hop_propagation",
					Description: `(Optional) Next Hop Propagation of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "multipath",
					Description: `(Optional) Multipath of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_as_path_prepend_last_as",
					Description: `(Optional) Number of ASN prepended to AS Path of the Action Rule Profile object.`,
				},
				resource.Attribute{
					Name:        "set_as_path_prepend_as",
					Description: `(Optional) A block representing ASNs configured as Set As Path - Prepend AS of the Action Rule Profile object. Type: Block.`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `ASN prepended to Set AS Path.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Order in which the ASN is prepended to Set AS Path.`,
				},
				resource.Attribute{
					Name:        "set_dampening",
					Description: `(Optional) A block representing the attributes of Set Dampening object. Type: Block.`,
				},
				resource.Attribute{
					Name:        "half_life",
					Description: `Half Life of the Set Dampening object.`,
				},
				resource.Attribute{
					Name:        "reuse",
					Description: `Reuse Limit of the Set Dampening object.`,
				},
				resource.Attribute{
					Name:        "suppress",
					Description: `Suppress Limit of the Set Dampening object.`,
				},
				resource.Attribute{
					Name:        "max_suppress_time",
					Description: `Max Suppress Time of the Set Dampening object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_annotation",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Annotation`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "parent_dn",
					Description: `(Required) Distinguished name of the object to which the Annotation is attached to.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Key of the Annotation. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Annotation.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value of the Annotation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Annotation.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value of the Annotation.`,
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
					Description: `(Required) Distinguished name of the parent VRF object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Any object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Any object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Any object.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) Represents the provider label match criteria.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of the Any object.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPgs can be divided in a the context can be divided into two subgroups.`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_any_to_cons",
					Description: `(Optional) Relation to Consumed Contracts (vzBrCP class)`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_any_to_cons_if",
					Description: `(Optional) Relation to Consumed Contract Interfaces (vzCPIf class)`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_any_to_prov",
					Description: `(Optional) Relation to Provided Contracts (vzBrCP class).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Any object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Any object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Any object.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) Represents the provider label match criteria.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of the Any object.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPgs can be divided in a the context can be divided into two subgroups.`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_any_to_cons",
					Description: `(Optional) Relation to Consumed Contracts (vzBrCP class)`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_any_to_cons_if",
					Description: `(Optional) Relation to Consumed Contract Interfaces (vzCPIf class)`,
				},
				resource.Attribute{
					Name:        "relation_vz_rs_any_to_prov",
					Description: `(Optional) Relation to Provided Contracts (vzBrCP class).`,
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
					Description: `(Required) Name of Object application epg. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Application EPG.`,
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
					Description: `(Optional) Name alias for object application epg.`,
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
					Description: `(Optional) QoS priority class id.`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `(Optional) Shutdown for object application epg.`,
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
					Description: `(Optional) Name alias for object application epg.`,
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
					Description: `(Optional) QoS priority class id.`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `(Optional) Shutdown for object application epg.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Description: `(Required) Name of Object attachable access entity profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the attachable access entity profile.`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the attachable access entity profile.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_authentication_properties",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI AAA Authentication Properties and Default Radius Authentication Settings`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the AAA Authentication.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of objects AAA Authentication and Default Radius Authentication Settings.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of objects AAA Authentication and Default Radius Authentication Settings.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of objects AAA Authentication and Default Radius Authentication Settings.`,
				},
				resource.Attribute{
					Name:        "def_role_policy",
					Description: `(Optional) The default role policy of object AAA Authentication.`,
				},
				resource.Attribute{
					Name:        "ping_check",
					Description: `(Optional) Heart bit ping checks for RADIUS/TACACS/LDAP/SAML/RSA server reachability.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) The number of attempts that the authentication method is tried.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The amount of time between authentication attempts.`,
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
					Description: `(Optional) Annotation for object Autonomous System Profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Autonomous System Profile.`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `(Optional) A number that uniquely identifies an autonomous system.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Autonomous System Profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Autonomous System Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Autonomous System Profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Autonomous System Profile.`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `(Optional) A number that uniquely identifies an autonomous system.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Autonomous System Profile.`,
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
			Type:             "aci_bfd_interface_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI BFD Interface Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object BFD Interface Policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Administrative state of the BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) Control state of object BFD Interface Policy`,
				},
				resource.Attribute{
					Name:        "detect_mult",
					Description: `(Optional) Detection multiplier for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "echo_admin_st",
					Description: `(Optional) Echo mode indicator for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "echo_rx_intvl",
					Description: `(Optional) Echo rx interval for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "min_rx_intvl",
					Description: `(Optional) Required minimum rx interval for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "min_tx_intvl",
					Description: `(Optional) Desired minimum tx interval for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object BFD Interface Policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Administrative state of the BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) Control state of object BFD Interface Policy`,
				},
				resource.Attribute{
					Name:        "detect_mult",
					Description: `(Optional) Detection multiplier for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "echo_admin_st",
					Description: `(Optional) Echo mode indicator for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "echo_rx_intvl",
					Description: `(Optional) Echo rx interval for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "min_rx_intvl",
					Description: `(Optional) Required minimum rx interval for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "min_tx_intvl",
					Description: `(Optional) Desired minimum tx interval for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object BFD Interface Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object BFD Interface Policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bgp_address_family_context",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI BGP Address Family Context`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bgp_best_path_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI BGP Best Path Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bgp_peer_connectivity_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI BGP Peer Connectivity Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bgp_peer_prefix",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI BGP Peer Prefix`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of BGP peer prefix object. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action when the maximum prefix limit is reached for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `Annotation for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "max_pfx",
					Description: `Maximum number of prefixes allowed from the peer for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `Name alias for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "restart_time",
					Description: `Time before restarting peer for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "thresh",
					Description: `Threshold for a maximum number of prefixes for BGP peer prefix object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action when the maximum prefix limit is reached for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `Annotation for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "max_pfx",
					Description: `Maximum number of prefixes allowed from the peer for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `Name alias for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "restart_time",
					Description: `Time before restarting peer for BGP peer prefix object.`,
				},
				resource.Attribute{
					Name:        "thresh",
					Description: `Threshold for a maximum number of prefixes for BGP peer prefix object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bgp_route_control_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI BGP Route Control Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bgp_route_summarization",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI BGP Route Summarization`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_bgp_timers",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI BGP Timers`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Description: `(Required) Name of Object cdp interface policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the CDP Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Administrative state`,
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
					Description: `(Optional) Description for object cdp interface policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the CDP Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Administrative state`,
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
					Description: `(Optional) Description for object cdp interface policy.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_account",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Account`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of the parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) ID of the Cloud Account object.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `(Required) Vendor of the Cloud Account object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Cloud Account object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Cloud Account object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Cloud Account object.`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Optional) Authentication type for the Cloud Account (managed=no credentials required (IAM), credentials=using accessKeys).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Cloud Account object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Cloud Account object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Cloud Account object.`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Optional) Authentication type for the Cloud Account (managed=no credentials required (IAM), credentials=using accessKeys).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_ad",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Active Directory`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of the parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Active Directory object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Active Directory.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Active Directory object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Active Directory object.`,
				},
				resource.Attribute{
					Name:        "active_directory_id",
					Description: `(Optional) Id of the Azure Active Directory.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Active Directory.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Active Directory object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Active Directory object.`,
				},
				resource.Attribute{
					Name:        "active_directory_id",
					Description: `(Optional) Id of the Azure Active Directory.`,
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
					Name:        "description",
					Description: `(Optional) Description for object cloud applicationcontainer.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object cloud applicationcontainer.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object cloud applicationcontainer.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Application container.`,
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
					Description: `(Optional) Name alias for object cloud applicationcontainer.`,
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
					Description: `(Required) Name of Object cloud availability zone. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Availability Zone.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object cloud availability zone.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object cloud availability zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Availability Zone.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object cloud availability zone.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object cloud availability zone.`,
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
					Description: `(Optional) access_key_id for the AWS account provided in the account id field.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) AWS account-id to manage with cloud APIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object cloud aws provider.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object cloud aws provider.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Email address of the local user.`,
				},
				resource.Attribute{
					Name:        "http_proxy",
					Description: `(Optional) Http proxy for object cloud aws provider.`,
				},
				resource.Attribute{
					Name:        "is_account_in_org",
					Description: `(Optional) Flag to decide whether the account is in the organization or not.`,
				},
				resource.Attribute{
					Name:        "is_trusted",
					Description: `(Optional) Whether the account is trusted with Tenant infra account.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object cloud aws provider.`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `(Optional) Provider id for object cloud aws provider.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Which AWS region to manage. \[Supported only in Cloud-APIC 4.2(x) or earlier\]`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Optional) Secret access key for the AWS account provided in the account id field.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud AWS Provider.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `(Optional) access_key_id for the AWS account provided in the account id field.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) AWS account-id to manage with cloud APIC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object cloud aws provider.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object cloud aws provider.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Email address of the local user.`,
				},
				resource.Attribute{
					Name:        "http_proxy",
					Description: `(Optional) Http proxy for object cloud aws provider.`,
				},
				resource.Attribute{
					Name:        "is_account_in_org",
					Description: `(Optional) Flag to decide whether the account is in the organization or not.`,
				},
				resource.Attribute{
					Name:        "is_trusted",
					Description: `(Optional) Whether the account is trusted with Tenant infra account.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object cloud aws provider.`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `(Optional) Provider id for object cloud aws provider.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Which AWS region to manage. \[Supported only in Cloud-APIC 4.2(x) or earlier\]`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Optional) Secret access key for the AWS account provided in the account id field.`,
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
					Description: `(Optional) This will represent whether CIDR is primary CIDR or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud CIDR Pool.`,
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
					Description: `(Required) Distinguished name of the parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Cloud Context Profile object. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Dn of the Cloud Context Profile object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Cloud Context Profile object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `Annotation of the Cloud Context Profile object.`,
				},
				resource.Attribute{
					Name:        "primary_cidr",
					Description: `Primary CIDR block of the Cloud Context Profile object.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of the Cloud Context Profile object.`,
				},
				resource.Attribute{
					Name:        "cloud_vendor",
					Description: `Name of the vendor.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `Name alias of the Cloud Context Profile object.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the Cloud Context Profile object.`,
				},
				resource.Attribute{
					Name:        "hub_network",
					Description: `Hub Network Dn which enables Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_ctx_to_flow_log",
					Description: `Relation to a AWS Flow Log Policy (class cloudAwsFlowLogPol).`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_to_ctx",
					Description: `Relation to a VRF (class fvCtx).`,
				},
				resource.Attribute{
					Name:        "cloud_brownfield",
					Description: `ID of imported brownfield virtual network.`,
				},
				resource.Attribute{
					Name:        "access_policy_type",
					Description: `Type of cloud context access policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Dn of the Cloud Context Profile object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Cloud Context Profile object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `Annotation of the Cloud Context Profile object.`,
				},
				resource.Attribute{
					Name:        "primary_cidr",
					Description: `Primary CIDR block of the Cloud Context Profile object.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of the Cloud Context Profile object.`,
				},
				resource.Attribute{
					Name:        "cloud_vendor",
					Description: `Name of the vendor.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `Name alias of the Cloud Context Profile object.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the Cloud Context Profile object.`,
				},
				resource.Attribute{
					Name:        "hub_network",
					Description: `Hub Network Dn which enables Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_ctx_to_flow_log",
					Description: `Relation to a AWS Flow Log Policy (class cloudAwsFlowLogPol).`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_to_ctx",
					Description: `Relation to a VRF (class fvCtx).`,
				},
				resource.Attribute{
					Name:        "cloud_brownfield",
					Description: `ID of imported brownfield virtual network.`,
				},
				resource.Attribute{
					Name:        "access_policy_type",
					Description: `Type of cloud context access policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_credentials",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Credential to manage the cloud resources`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of the parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Cloud Credential object used to manage the cloud resources. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Credential to manage the cloud resources.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Cloud Credential object to manage the cloud resources.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Cloud Credential object to manage the cloud resources.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) Client ID of the Cloud Credential object.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) email address of the locally-authenticated user.`,
				},
				resource.Attribute{
					Name:        "http_proxy",
					Description: `(Optional) HTTP Proxy to connect to the cloud provider.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The Secret key or password used to uniquely identify the cloud resource configuration object.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) The Access key ID of the cloud resource.`,
				},
				resource.Attribute{
					Name:        "rsa_private_key",
					Description: `(Optional) RSA Secret Key of the cloud resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Credential to manage the cloud resources.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Cloud Credential object to manage the cloud resources.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Cloud Credential object to manage the cloud resources.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) Client ID of the Cloud Credential object.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) email address of the locally-authenticated user.`,
				},
				resource.Attribute{
					Name:        "http_proxy",
					Description: `(Optional) HTTP Proxy to connect to the cloud provider.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The Secret key or password used to uniquely identify the cloud resource configuration object.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) The Access key ID of the cloud resource.`,
				},
				resource.Attribute{
					Name:        "rsa_private_key",
					Description: `(Optional) RSA Secret Key of the cloud resource.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_endpoint_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Endpoint Selector`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_endpoint_selectorfor_external_epgs",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Endpoint Selector for External EPgs`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_epg",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud EPg`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Description: `(Required) Distinguished name of parent Cloud Application container object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object Cloud External EPg. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud External EPg.`,
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
					Description: `(Optional) Exception-tag for object Cloud External EPg.`,
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
					Description: `(Optional) Name alias for object Cloud External EPg.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) QOS priority class id.`,
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
					Name:        "description",
					Description: `(Optional) Description for object Cloud External EPg.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Cloud External EPg.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) Exception-tag for object Cloud External EPg.`,
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
					Description: `(Optional) Name alias for object Cloud External EPg.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if EPg is part of a group that does not a contract for communication.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) QOS priority class id.`,
				},
				resource.Attribute{
					Name:        "route_reachability",
					Description: `(Optional) Route reachability for this EPG.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_external_network",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud External Network`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Cloud External Network. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud External Network.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Cloud External Network.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Cloud External Network.`,
				},
				resource.Attribute{
					Name:        "vrf_dn",
					Description: `(Optional) Distinguished name of the VRF. Note that the VRF has to be created under the infra tenant.`,
				},
				resource.Attribute{
					Name:        "hub_network_name",
					Description: `(Optional) Hub Network name of the Cloud External Network.`,
				},
				resource.Attribute{
					Name:        "vpn_router_name",
					Description: `(Optional) VPN Router name of the Cloud External Network.`,
				},
				resource.Attribute{
					Name:        "host_router_name",
					Description: `(Optional) Host Router name of the Cloud External Network.`,
				},
				resource.Attribute{
					Name:        "all_regions",
					Description: `(Optional) Selects all regions available to the Cloud External Network. This option is always set to "yes" for Azure and AWS cAPICs and "no" in GCP cAPIC.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Optional) Manually adds the regions to the Cloud External Network. This option is only available in GCP cAPICs.`,
				},
				resource.Attribute{
					Name:        "router_type",
					Description: `(Optional) Router type. Allowed values are "c8kv", "tgw". (Available only for AWS cAPIC).`,
				},
				resource.Attribute{
					Name:        "cloud_vendor",
					Description: `(Optional) Name of the vendor. Allowed values are "aws", "azure", "gcp".`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud External Network.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Cloud External Network.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Cloud External Network.`,
				},
				resource.Attribute{
					Name:        "vrf_dn",
					Description: `(Optional) Distinguished name of the VRF. Note that the VRF has to be created under the infra tenant.`,
				},
				resource.Attribute{
					Name:        "hub_network_name",
					Description: `(Optional) Hub Network name of the Cloud External Network.`,
				},
				resource.Attribute{
					Name:        "vpn_router_name",
					Description: `(Optional) VPN Router name of the Cloud External Network.`,
				},
				resource.Attribute{
					Name:        "host_router_name",
					Description: `(Optional) Host Router name of the Cloud External Network.`,
				},
				resource.Attribute{
					Name:        "all_regions",
					Description: `(Optional) Selects all regions available to the Cloud External Network. This option is always set to "yes" for Azure and AWS cAPICs and "no" in GCP cAPIC.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Optional) Manually adds the regions to the Cloud External Network. This option is only available in GCP cAPICs.`,
				},
				resource.Attribute{
					Name:        "router_type",
					Description: `(Optional) Router type. Allowed values are "c8kv", "tgw". (Available only for AWS cAPIC).`,
				},
				resource.Attribute{
					Name:        "cloud_vendor",
					Description: `(Optional) Name of the vendor. Allowed values are "aws", "azure", "gcp".`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_ipsec_tunnel_subnet_pool",
			Category:         "Data Sources",
			ShortDescription: `Data source for the ACI Cloud Subnet Pool for IPsec Tunnels`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "subnet_pool",
					Description: `(Required) Subnet of the Subnet Pool for IPsec Tunnels object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Subnet Pool for IPsec Tunnels object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Subnet Pool Name of the Subnet Pool for IPsec Tunnels object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Subnet Pool for IPsec Tunnels object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Subnet Pool Name of the Subnet Pool for IPsec Tunnels object.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Description: `(Required) Name of Object cloud providers region. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Providers Region.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Administrative state of the object or policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object cloud providers region.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object cloud providers region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Providers Region.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Administrative state of the object or policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object cloud providers region.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object cloud providers region.`,
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
					Description: `(Required) Distinguished name of the Cloud CIDR Pool parent object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) CIDR block of Object cloud subnet. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Dn of the Cloud Subnet object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Cloud Subnet object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Cloud Subnet object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Cloud Subnet object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of the Cloud Subnet object.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) List of domain applicable to the capability. Allowed values are "public", "private" and "shared". Default is ["private"].`,
				},
				resource.Attribute{
					Name:        "usage",
					Description: `(Optional) The usage of the port. This property shows how the port is used.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Relation to a Cloud Resource Zone (class cloudRsZoneAttach). It is only applicable to the AWS vendor.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_subnet_to_flow_log",
					Description: `(Optional) Relation to the AWS Flow Log Policy (class cloudAwsFlowLogPol).`,
				},
				resource.Attribute{
					Name:        "subnet_group_label",
					Description: `(Optional) Subnet Group Label of the Cloud Subnet object. It is only applicable to the GCP vendor.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Dn of the Cloud Subnet object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Cloud Subnet object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Cloud Subnet object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Cloud Subnet object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of the Cloud Subnet object.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) List of domain applicable to the capability. Allowed values are "public", "private" and "shared". Default is ["private"].`,
				},
				resource.Attribute{
					Name:        "usage",
					Description: `(Optional) The usage of the port. This property shows how the port is used.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Relation to a Cloud Resource Zone (class cloudRsZoneAttach). It is only applicable to the AWS vendor.`,
				},
				resource.Attribute{
					Name:        "relation_cloud_rs_subnet_to_flow_log",
					Description: `(Optional) Relation to the AWS Flow Log Policy (class cloudAwsFlowLogPol).`,
				},
				resource.Attribute{
					Name:        "subnet_group_label",
					Description: `(Optional) Subnet Group Label of the Cloud Subnet object. It is only applicable to the GCP vendor.`,
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
					Description: `(Required) Name of Object Cloud Router Profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud Router Profile.`,
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
					Description: `(Optional) Name alias for object Cloud Router Profile.`,
				},
				resource.Attribute{
					Name:        "num_instances",
					Description: `(Optional) Num instances for object Cloud Router Profile.`,
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
					Name:        "description",
					Description: `(Optional) Description for object Cloud Router Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Cloud Router Profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Cloud Router Profile.`,
				},
				resource.Attribute{
					Name:        "num_instances",
					Description: `(Optional) Num instances for object Cloud Router Profile.`,
				},
				resource.Attribute{
					Name:        "cloud_router_profile_type",
					Description: `(Optional) Component type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_vpn_network",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Cloud Template for VPN Network`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aci_cloud_external_network_dn",
					Description: `(Required) Distinguished name of parent TemplateforExternalNetwork object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Cloud VPN Network object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud VPN Network.`,
				},
				resource.Attribute{
					Name:        "remote_site_id",
					Description: `(Optional) Remote Site ID.`,
				},
				resource.Attribute{
					Name:        "remote_site_name",
					Description: `(Optional) Name of the Remote Site.`,
				},
				resource.Attribute{
					Name:        "ipsec_tunnel",
					Description: `(Optional) IPsec tunnel destination (cloudtemplateIpSecTunnelSourceInterface class). Type: Block.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `(Required) IKE version. Allowed values are "ikev1", "ikev2", and default value is "ikev2".`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `(Required) Peer address of the Cloud IPsec tunnel object.`,
				},
				resource.Attribute{
					Name:        "subnet_pool_name",
					Description: `(Required) Subnet Pool Name.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(Optional) Pre Shared Key for all tunnels to this peer address.`,
				},
				resource.Attribute{
					Name:        "bgp_peer_asn",
					Description: `(Required) BGP ASN Number. A number that uniquely identifies an autonomous system.`,
				},
				resource.Attribute{
					Name:        "source_interfaces",
					Description: `(Optional) Source Interface Ids of the object for IPsec tunnel Source Interface. It is available only on Azure cAPIC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Cloud VPN Network.`,
				},
				resource.Attribute{
					Name:        "remote_site_id",
					Description: `(Optional) Remote Site ID.`,
				},
				resource.Attribute{
					Name:        "remote_site_name",
					Description: `(Optional) Name of the Remote Site.`,
				},
				resource.Attribute{
					Name:        "ipsec_tunnel",
					Description: `(Optional) IPsec tunnel destination (cloudtemplateIpSecTunnelSourceInterface class). Type: Block.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `(Required) IKE version. Allowed values are "ikev1", "ikev2", and default value is "ikev2".`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `(Required) Peer address of the Cloud IPsec tunnel object.`,
				},
				resource.Attribute{
					Name:        "subnet_pool_name",
					Description: `(Required) Subnet Pool Name.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(Optional) Pre Shared Key for all tunnels to this peer address.`,
				},
				resource.Attribute{
					Name:        "bgp_peer_asn",
					Description: `(Required) BGP ASN Number. A number that uniquely identifies an autonomous system.`,
				},
				resource.Attribute{
					Name:        "source_interfaces",
					Description: `(Optional) Source Interface Ids of the object for IPsec tunnel Source Interface. It is available only on Azure cAPIC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_cloud_vrf_leak_routes",
			Category:         "Data Sources",
			ShortDescription: `Data source for Cloud ACI Inter-VRF Leaked Routes`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrf_dn",
					Description: `(Required) Distinguished name of the parent VRF object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Inter-VRF Leaked Route object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Inter-VRF Leaked Route object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Inter-VRF Leaked Route object.`,
				},
				resource.Attribute{
					Name:        "leak_to",
					Description: `(Optional) A block representing the attributes of ` + "`" + `Leak Routes` + "`" + ` for the Inter-VRF Leaked Route object. Type: Block.`,
				},
				resource.Attribute{
					Name:        "vrf_dn",
					Description: `Distinguished name of the destination VRF object, which is mapped to the Inter-VRF Leaked Route object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Inter-VRF Leaked Route object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Inter-VRF Leaked Route object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Inter-VRF Leaked Route object.`,
				},
				resource.Attribute{
					Name:        "leak_to",
					Description: `(Optional) A block representing the attributes of ` + "`" + `Leak Routes` + "`" + ` for the Inter-VRF Leaked Route object. Type: Block.`,
				},
				resource.Attribute{
					Name:        "vrf_dn",
					Description: `Distinguished name of the destination VRF object, which is mapped to the Inter-VRF Leaked Route object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_concrete_device",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Concrete Device`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "l4-l7_device_dn",
					Description: `(Required) Distinguished name of the parent L4-L7 Device object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Concrete Device object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Concrete Device.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Concrete Device object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Concrete Device object.`,
				},
				resource.Attribute{
					Name:        "vmm_controller_dn",
					Description: `(Optional) Distinguished name of the VMM controller object. Type: String.`,
				},
				resource.Attribute{
					Name:        "vm_name",
					Description: `(Optional) The name of the Virtual Machine (VM) in the vCenter on which the device is hosted in the L4-L7 device cluster. It uniquely identifies the VM. Type: String.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Concrete Device.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Concrete Device object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Concrete Device object.`,
				},
				resource.Attribute{
					Name:        "vmm_controller_dn",
					Description: `(Optional) Distinguished name of the VMM controller object. Type: String.`,
				},
				resource.Attribute{
					Name:        "vm_name",
					Description: `(Optional) The name of the Virtual Machine (VM) in the vCenter on which the device is hosted in the L4-L7 device cluster. It uniquely identifies the VM. Type: String.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_concrete_interface",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Concrete Interface`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "concrete_device_dn",
					Description: `(Required) Distinguished name of the parent Concrete Device object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Concrete Interface object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Concrete Interface.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Concrete Interface object.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) The port encapsulation. Type: String.`,
				},
				resource.Attribute{
					Name:        "vnic_name",
					Description: `(Optional) The virtual NIC (vNIC) name of the L4-L7 Device VM represented by the concrete interface. Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_vns_rs_c_if_path_att",
					Description: `(Optional) Represents a relation from the Concrete Interface to the Physical Port on the Leaf (class fabricPathEp). Type: String.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Concrete Interface.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Concrete Interface object.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) The port encapsulation. Type: String.`,
				},
				resource.Attribute{
					Name:        "vnic_name",
					Description: `(Optional) The virtual NIC (vNIC) name of the L4-L7 Device VM represented by the concrete interface. Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_vns_rs_c_if_path_att",
					Description: `(Optional) Represents a relation from the Concrete Interface to the Physical Port on the Leaf (class fabricPathEp). Type: String.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_configuration_import_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Configuration Import Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_connection",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Connection`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_console_authentication",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Console Authentication`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Console Authentication.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Console Authentication.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Console Authentication.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Console Authentication.`,
				},
				resource.Attribute{
					Name:        "provider_group",
					Description: `(Optional) Provider Group. An AAA configuration provider group is a group of remote servers supporting the same AAA protocol that will be used for authentication and authorization. When a provider group is specified, only the servers within that group will be used for authentication and authorization. If no provider group is specified, all servers supporting the realm of AAA protocols will be used for authentication and authorization.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `(Optional) Realm. The security method for processing authentication and authorization requests. The realm allows the protected resources on the associated server to be partitioned into a set of protection spaces, each with its own authentication authorization database. This is an abstract class and cannot be instantiated.`,
				},
				resource.Attribute{
					Name:        "realm_sub_type",
					Description: `(Optional) Realm subtype that can be default or Duo.`,
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
					Description: `(Optional) Priority level of the service contract.`,
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
					Description: `(Optional) Priority level of the service contract.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_contract_subject_filter",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Contract Subject Filter`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "contract_subject_dn",
					Description: `(Required) Distinguished name of parent Contract Subject object.`,
				},
				resource.Attribute{
					Name:        "filter_dn",
					Description: `(Required) Distinguished name of the Filter object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Subject Filter.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Contract Subject Filter object.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The action required when the condition is met.`,
				},
				resource.Attribute{
					Name:        "directives",
					Description: `(Optional) Directives of the Contract Subject Filter object.`,
				},
				resource.Attribute{
					Name:        "priority_override",
					Description: `(Optional) Priority Override of the Contract Subject Filter object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Subject Filter.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Contract Subject Filter object.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The action required when the condition is met.`,
				},
				resource.Attribute{
					Name:        "directives",
					Description: `(Optional) Directives of the Contract Subject Filter object.`,
				},
				resource.Attribute{
					Name:        "priority_override",
					Description: `(Optional) Priority Override of the Contract Subject Filter object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_contract_subject_one_way_filter",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI One Way Filter`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "contract_subject_dn",
					Description: `(Required) Distinguished name of the parent Contract Subject object.`,
				},
				resource.Attribute{
					Name:        "filter_dn",
					Description: `(Required) Distinguished name of the Filter object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Filter.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Filter object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Filter object.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The action required when the condition is met. Allowed values are "deny", "permit", and the default value is "permit". Type: String.`,
				},
				resource.Attribute{
					Name:        "directives",
					Description: `(Optional) Directives of the Contract Subject Filter object. Allowed values are "log", "no_stats", "none", and the default value is "none". Type: List.`,
				},
				resource.Attribute{
					Name:        "priority_override",
					Description: `(Optional) Priority override of the Filter object. It is only used when action is set to deny. Allowed values are "default", "level1", "level2", "level3", and the default value is "default". Type: String.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Filter.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Filter object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Filter object.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The action required when the condition is met. Allowed values are "deny", "permit", and the default value is "permit". Type: String.`,
				},
				resource.Attribute{
					Name:        "directives",
					Description: `(Optional) Directives of the Contract Subject Filter object. Allowed values are "log", "no_stats", "none", and the default value is "none". Type: List.`,
				},
				resource.Attribute{
					Name:        "priority_override",
					Description: `(Optional) Priority override of the Filter object. It is only used when action is set to deny. Allowed values are "default", "level1", "level2", "level3", and the default value is "default". Type: String.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_coop_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI COOP Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the COOP Group Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object COOP Group Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object COOP Group Policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Authentication type. The specific type of the object or component.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object COOP Group Policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_default_authentication",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Default Authentication Method for all Logins`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Default Authentication Method for all Logins.`,
				},
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
					Description: `(Optional) The parameter to disable fallback in case there are active servers in the default auth type.`,
				},
				resource.Attribute{
					Name:        "provider_group",
					Description: `(Optional) The group of remote servers supporting the same AAA protocol that will be used for authentication and authorization.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `(Optional) The security method for processing authentication and authorization requests. The realm allows the protected resources on the associated server to be partitioned into a set of protection spaces, each with its own authentication authorization database.`,
				},
				resource.Attribute{
					Name:        "realm_sub_type",
					Description: `(Optional) Realm subtype.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_duo_provider_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Duo Provider Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object Duo Provider Group. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Duo Provider Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Duo Provider Group.`,
				},
				resource.Attribute{
					Name:        "auth_choice",
					Description: `(Optional) Authentication choice of object Duo Provider Group.`,
				},
				resource.Attribute{
					Name:        "provider_type",
					Description: `(Optional) Type of the Auth Provider.`,
				},
				resource.Attribute{
					Name:        "ldap_group_map_ref",
					Description: `(Optional) Reference to LDAP Group Map containing user's group membership info.`,
				},
				resource.Attribute{
					Name:        "sec_fac_auth_methods",
					Description: `(Optional) Second factor authentication methods.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of object Duo Provider Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Duo Provider Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Duo Provider Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Duo Provider Group.`,
				},
				resource.Attribute{
					Name:        "auth_choice",
					Description: `(Optional) Authentication choice of object Duo Provider Group.`,
				},
				resource.Attribute{
					Name:        "provider_type",
					Description: `(Optional) Type of the Auth Provider.`,
				},
				resource.Attribute{
					Name:        "ldap_group_map_ref",
					Description: `(Optional) Reference to LDAP Group Map containing user's group membership info.`,
				},
				resource.Attribute{
					Name:        "sec_fac_auth_methods",
					Description: `(Optional) Second factor authentication methods.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of object Duo Provider Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Duo Provider Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_encryption_key",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI AES Encryption Passphrase and Keys for Config Export and Import`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the AES Encryption Passphrase and Keys for Config Export (and Import).`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object AES Encryption Passphrase and Keys for Config Export and Import.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object AES Encryption Passphrase and Keys for Config Export and Import.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object AES Encryption Passphrase and Keys for Config Export and Import.`,
				},
				resource.Attribute{
					Name:        "passphrase_key_derivation_version",
					Description: `(Optional) Version of the algorithm used for forward compatibility.`,
				},
				resource.Attribute{
					Name:        "strong_encryption_enabled",
					Description: `(Optional) Parameter indicating whether encryption is weak or strong.`,
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
					Name:        "descripton",
					Description: `(Optional) Descripton for object end point retention policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object end point retention policy.`,
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
					Description: `(Optional) Name alias for object end point retention policy.`,
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
					Name:        "descripton",
					Description: `(Optional) Descripton for object end point retention policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object end point retention policy.`,
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
					Description: `(Optional) Name alias for object end point retention policy.`,
				},
				resource.Attribute{
					Name:        "remote_ep_age_intvl",
					Description: `(Optional) The aging interval for all remote endpoints learned in this bridge domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_endpoint_controls",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Endpoint Control`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Endpoint Control.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint Control.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Endpoint Control.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) The administrative state of object Endpoint Control.`,
				},
				resource.Attribute{
					Name:        "hold_intvl",
					Description: `(Optional) The period of time before declaring that the neighbor is down of object Endpoint Control.`,
				},
				resource.Attribute{
					Name:        "rogue_ep_detect_intvl",
					Description: `(Optional) Rogue Endpoint Detection Interval of object Endpoint Control.`,
				},
				resource.Attribute{
					Name:        "rogue_ep_detect_mult",
					Description: `(Optional) Rogue Endpoint Detection Multiplication Factor of object Endpoint Control.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Endpoint Control.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_endpoint_ip_aging_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Endpoint IP Aging Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Endpoint IP Aging Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint IP Aging Profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Endpoint IP Aging Profile.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) The administrative state of the object Endpoint IP Aging Profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Endpoint IP Aging Profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_endpoint_loop_protection",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Endpoint Loop Protection`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Endpoint Loop Protection Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint Loop Protection Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Endpoint Loop Protection Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Endpoint Loop Protection.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action. Sets the action to take when a loop is detected.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Admin State. The administrative state of the object or policy.`,
				},
				resource.Attribute{
					Name:        "loop_detect_intvl",
					Description: `(Optional) Loop Detection Interval. Sets the loop detection interval, which specifies the time to detect a loop.`,
				},
				resource.Attribute{
					Name:        "loop_detect_mult",
					Description: `(Optional) Loop Detection Multiplier. Sets the loop detection multiplication factor, which is the number of times a single Endpoint moves between ports within the Detection interval.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_endpoint_security_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Endpoint Security Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_profile_dn",
					Description: `(Required) Distinguished name of parent Application Profile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object Endpoint Security Group. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Endpoint Security Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint Security Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Endpoint Security Group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Endpoint Security Group.`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Handles L2 Multicast/Broadcast and Link-Layer traffic at EPG level. It represents Control at EPG level and decides if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP, or based on bridge-domain settings.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria.`,
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
					Description: `(Optional) The QoS priority class identifier.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Endpoint Security Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint Security Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Endpoint Security Group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Endpoint Security Group.`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Handles L2 Multicast/Broadcast and Link-Layer traffic at EPG level. It represents Control at EPG level and decides if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP, or based on bridge-domain settings.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria.`,
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
					Description: `(Optional) The QoS priority class identifier.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_endpoint_security_group_epg_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Endpoint Security Group EPG Selector`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_security_group_dn",
					Description: `(Required) Distinguished name of parent Endpoint Security Group object.`,
				},
				resource.Attribute{
					Name:        "match_epg_dn",
					Description: `(Required) EPG Dn to be associated. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Endpoint Security Group EPG Selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint Security Group EPG Selector.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Endpoint Security Group EPG Selector.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Endpoint Security Group EPG Selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint Security Group EPG Selector.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Endpoint Security Group EPG Selector.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_endpoint_security_group_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Endpoint Security Group Selector`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_security_group_dn",
					Description: `(Required) Distinguished name of parent Endpoint Security Group object.`,
				},
				resource.Attribute{
					Name:        "match_expression",
					Description: `(Required) Expression used to define matching tags. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Endpoint Security Group Selector.`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Endpoint Security Group Selector.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_endpoint_security_group_tag_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Endpoint Security Group Tag Selector`,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Required) Match value of object Endpoint Security Group Tag Selector. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Endpoint Security Group Tag Selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint Security Group Tag Selector.`,
				},
				resource.Attribute{
					Name:        "match_key",
					Description: `(Optional) Key of Tag to be associated with.`,
				},
				resource.Attribute{
					Name:        "match_value",
					Description: `(Optional) Value of Tag to be associated with.`,
				},
				resource.Attribute{
					Name:        "value_operator",
					Description: `(Optional) Match Value Operator.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Endpoint Security Group Tag Selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Endpoint Security Group Tag Selector.`,
				},
				resource.Attribute{
					Name:        "match_key",
					Description: `(Optional) Key of Tag to be associated with.`,
				},
				resource.Attribute{
					Name:        "match_value",
					Description: `(Optional) Value of Tag to be associated with.`,
				},
				resource.Attribute{
					Name:        "value_operator",
					Description: `(Optional) Match Value Operator.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_epg_to_contract_interface",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Contract Interface Relationship`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_epg_dn",
					Description: `(Required) Distinguished name of the parent Application EPG object.`,
				},
				resource.Attribute{
					Name:        "contract_interface_dn",
					Description: `(Required) Distinguished name of the Contract Interface object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Contract Interface Relationship object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Contract Interface Relationship object.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The Contract Interface Relationship priority.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Contract Interface Relationship object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Contract Interface Relationship object.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The Contract Interface Relationship priority.`,
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
					Description: `(Required) Vmm domain instance. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Domain.`,
				},
				resource.Attribute{
					Name:        "binding_type",
					Description: `(Optional) Binding type for object Domain.`,
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
					Description: `(Optional) Encap mode for object Domain.`,
				},
				resource.Attribute{
					Name:        "epg_cos",
					Description: `(Optional) Epg cos for object Domain.`,
				},
				resource.Attribute{
					Name:        "epg_cos_pref",
					Description: `(Optional) Epg cos pref for object Domain.`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) Determines when policies are pushed to cam.`,
				},
				resource.Attribute{
					Name:        "netflow_dir",
					Description: `(Optional) Netflow dir for object Domain.`,
				},
				resource.Attribute{
					Name:        "netflow_pref",
					Description: `(Optional) Netflow pref for object Domain.`,
				},
				resource.Attribute{
					Name:        "num_ports",
					Description: `(Optional) Number of ports existing operationally in module`,
				},
				resource.Attribute{
					Name:        "port_allocation",
					Description: `(Optional) Port allocation for object Domain.`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) Primary encap for object Domain.`,
				},
				resource.Attribute{
					Name:        "primary_encap_inner",
					Description: `(Optional) Primary encap inner for object Domain.`,
				},
				resource.Attribute{
					Name:        "res_imedcy",
					Description: `(Optional) Policy resolution`,
				},
				resource.Attribute{
					Name:        "secondary_encap_inner",
					Description: `(Optional) Secondary encap inner for object Domain.`,
				},
				resource.Attribute{
					Name:        "switching_mode",
					Description: `(Optional) Switching mode for object Domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Domain.`,
				},
				resource.Attribute{
					Name:        "binding_type",
					Description: `(Optional) Binding type for object Domain.`,
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
					Description: `(Optional) Encap mode for object Domain.`,
				},
				resource.Attribute{
					Name:        "epg_cos",
					Description: `(Optional) Epg cos for object Domain.`,
				},
				resource.Attribute{
					Name:        "epg_cos_pref",
					Description: `(Optional) Epg cos pref for object Domain.`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) Determines when policies are pushed to cam.`,
				},
				resource.Attribute{
					Name:        "netflow_dir",
					Description: `(Optional) Netflow dir for object Domain.`,
				},
				resource.Attribute{
					Name:        "netflow_pref",
					Description: `(Optional) Netflow pref for object Domain.`,
				},
				resource.Attribute{
					Name:        "num_ports",
					Description: `(Optional) Number of ports existing operationally in module`,
				},
				resource.Attribute{
					Name:        "port_allocation",
					Description: `(Optional) Port allocation for object Domain.`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) Primary encap for object Domain.`,
				},
				resource.Attribute{
					Name:        "primary_encap_inner",
					Description: `(Optional) Primary encap inner for object Domain.`,
				},
				resource.Attribute{
					Name:        "res_imedcy",
					Description: `(Optional) Policy resolution`,
				},
				resource.Attribute{
					Name:        "secondary_encap_inner",
					Description: `(Optional) Secondary encap inner for object Domain.`,
				},
				resource.Attribute{
					Name:        "switching_mode",
					Description: `(Optional) Switching mode for object Domain.`,
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
					Name:        "tdn",
					Description: `(Required) tdn of Object Static Path. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Static Path.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Static Path.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) Encapsulation of the Static Path.`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) Immediacy of the Static Path.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Mode of the static association with the path.`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) Primary encap for object Static Path.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Static Path.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Static Path.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) Encapsulation of the Static Path.`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) Immediacy of the Static Path.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Mode of the static association with the path.`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) Primary encap for object Static Path.`,
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
					Description: `(Required) tDn of Object EPGs Using Function. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the EPGs Using Function.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object EPGs Using Function.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) Vlan number encap`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) Instrumentation immediacy`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Bgp domain mode`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) Primary encap for object EPGs Using Function.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the EPGs Using Function.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object EPGs Using Function.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) Vlan number encap`,
				},
				resource.Attribute{
					Name:        "instr_imedcy",
					Description: `(Optional) Instrumentation immediacy`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Bgp domain mode`,
				},
				resource.Attribute{
					Name:        "primary_encap",
					Description: `(Optional) Primary encap for object EPGs Using Function.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_error_disable_recovery",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Error Disable Recovery`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Error Disable Recovery.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Error Disable Recovery.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Error Disable Recovery.`,
				},
				resource.Attribute{
					Name:        "err_dis_recov_intvl",
					Description: `(Optional) Error Disable Recovery Interval. Sets the error disable recovery interval, which specifies the time to recover from an error-disabled state.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Error Disable Recovery.`,
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
					Description: `(Required) Distinguished name of the parent L3Outside object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the External Network Instance Profile object. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the External Network Instance Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the External Network Instance Profile object.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) Exception tag of the External Network Instance Profile object.`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria of the External Network Instance Profile object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of the External Network Instance Profile object.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if an External EPG is part of a group that does not require a contract for communication.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The QoS priority class identifier of the External Network Instance Profile object.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_sec_inherited",
					Description: `(Optional) Relation to EPGs to be used as Contract Masters (class fvEPg). Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prov",
					Description: `(Optional) Relation to Provided Contracts (class vzBrCP). Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons_if",
					Description: `(Optional) Relation to Provided Contract Interfaces (class vzCPIf). Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_inst_p_to_profile",
					Description: `(Optional) Relation to Route Control Profiles (class rtctrlProfile). Cardinality - N_TO_M. Type: Block.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons",
					Description: `(Optional) Relation to Consumed Contracts (class vzBrCP). Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prot_by",
					Description: `(Optional) Relation to Taboo Contracts (vzTaboo). Cardinality - N_TO_M. Type - Set of String.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the External Network Instance Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the External Network Instance Profile object.`,
				},
				resource.Attribute{
					Name:        "exception_tag",
					Description: `(Optional) Exception tag of the External Network Instance Profile object.`,
				},
				resource.Attribute{
					Name:        "flood_on_encap",
					Description: `(Optional) Control at EPG level if the traffic L2 Multicast/Broadcast and Link Local Layer should be flooded only on ENCAP or based on bridg-domain settings.`,
				},
				resource.Attribute{
					Name:        "match_t",
					Description: `(Optional) The provider label match criteria of the External Network Instance Profile object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of the External Network Instance Profile object.`,
				},
				resource.Attribute{
					Name:        "pref_gr_memb",
					Description: `(Optional) Represents parameter used to determine if an External EPG is part of a group that does not require a contract for communication.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `(Optional) The QoS priority class identifier of the External Network Instance Profile object.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_sec_inherited",
					Description: `(Optional) Relation to EPGs to be used as Contract Masters (class fvEPg). Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prov",
					Description: `(Optional) Relation to Provided Contracts (class vzBrCP). Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons_if",
					Description: `(Optional) Relation to Provided Contract Interfaces (class vzCPIf). Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_inst_p_to_profile",
					Description: `(Optional) Relation to Route Control Profiles (class rtctrlProfile). Cardinality - N_TO_M. Type: Block.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_cons",
					Description: `(Optional) Relation to Consumed Contracts (class vzBrCP). Cardinality - N_TO_M. Type - Set of String.`,
				},
				resource.Attribute{
					Name:        "relation_fv_rs_prot_by",
					Description: `(Optional) Relation to Taboo Contracts (vzTaboo). Cardinality - N_TO_M. Type - Set of String.`,
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
					Description: `(Required) Name of object fabric if pol. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the fabric if pol.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "auto_neg",
					Description: `(Optional) Policy auto-negotiation for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "fec_mode",
					Description: `(Optional) Forwarding error correction for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "link_debounce",
					Description: `(Optional) Link debounce interval for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) Port speed for object fabric if pol.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the fabric if pol.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "auto_neg",
					Description: `(Optional) Policy auto-negotiation for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "fec_mode",
					Description: `(Optional) Forwarding error correction for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "link_debounce",
					Description: `(Optional) Link debounce interval for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object fabric if pol.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) Port speed for object fabric if pol.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fabric_node",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Fabric Node`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fabric_pod_dn",
					Description: `(Required) Distinguished name of parent Fabric Pod object.`,
				},
				resource.Attribute{
					Name:        "fabric_node_id",
					Description: `(Required) fabric_node_id of object Fabric Node. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Fabric Node.`,
				},
				resource.Attribute{
					Name:        "ad_st",
					Description: `(Optional) The administrative state of object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "apic_type",
					Description: `(Optional) The APIC type for object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "fabric_st",
					Description: `(Optional) Fabric state for object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) IP address of object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Optional) Fabric Node type of object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) Fabric Node role of object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Fabric Node.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Fabric Node.`,
				},
				resource.Attribute{
					Name:        "ad_st",
					Description: `(Optional) The administrative state of object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "apic_type",
					Description: `(Optional) The APIC type for object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "fabric_st",
					Description: `(Optional) Fabric state for object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) IP address of object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Optional) Fabric Node type of object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) Fabric Node role of object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Fabric Node.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fabric_node_control",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Fabric Node Control`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object Fabric Node Control. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Fabric Node Control.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Fabric Node Control.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Fabric Node Control.`,
				},
				resource.Attribute{
					Name:        "control",
					Description: `(Optional) Fabric node control bitmask of object Fabric Node Control.`,
				},
				resource.Attribute{
					Name:        "feature_sel",
					Description: `(Optional) Feature Selection of object Fabric Node Control.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Fabric Node Control.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Fabric Node Control.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Fabric Node Control.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Fabric Node Control.`,
				},
				resource.Attribute{
					Name:        "control",
					Description: `(Optional) Fabric node control bitmask of object Fabric Node Control.`,
				},
				resource.Attribute{
					Name:        "feature_sel",
					Description: `(Optional) Feature Selection of object Fabric Node Control.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Fabric Node Control.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Description: `(Required) name of Object fabric path endpoint.`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `(Optional) Boolean, set to true if path is for a vPC interface ## Attribute Reference`,
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
			Type:             "aci_fabric_wide_settings",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Fabric-Wide Settings Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Fabric-Wide Settings Policy.`,
				},
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
					Description: `(Optional) Disable Ep Dampening knob of object Fabric-Wide Settings Policy.`,
				},
				resource.Attribute{
					Name:        "domain_validation",
					Description: `(Optional) Validate that the correct physical domain is added before associating a new static path to an EPG.`,
				},
				resource.Attribute{
					Name:        "enable_mo_streaming",
					Description: `(Optional) Enable MO streaming of object Fabric-Wide Settings Policy.`,
				},
				resource.Attribute{
					Name:        "enable_remote_leaf_direct",
					Description: `(Optional) Enable remote leaf direct communication of object Fabric-Wide Settings Policy.`,
				},
				resource.Attribute{
					Name:        "enforce_subnet_check",
					Description: `(Optional) Enforce subnet check of object Fabric-Wide Settings Policy.`,
				},
				resource.Attribute{
					Name:        "leaf_opflexp_authenticate_clients",
					Description: `(Optional) Require Opflexp Client Certificates for authentication for Leaf.`,
				},
				resource.Attribute{
					Name:        "leaf_opflexp_use_ssl",
					Description: `(Optional) Require SSL transport for Opflexp for Leaf.`,
				},
				resource.Attribute{
					Name:        "opflexp_authenticate_clients",
					Description: `(Optional) Opflexp Client Certificates for authentication of object Fabric-Wide Settings Policy.`,
				},
				resource.Attribute{
					Name:        "opflexp_ssl_protocols",
					Description: `(Optional) SSL Opflex versions.`,
				},
				resource.Attribute{
					Name:        "opflexp_use_ssl",
					Description: `(Optional) SSL transport for Opflexp indicator of object Fabric-Wide Settings Policy.`,
				},
				resource.Attribute{
					Name:        "policy_sync_node_bringup",
					Description: `(Optional) Blacklist the Leaf frontpanel port until policy download during first time bringup.`,
				},
				resource.Attribute{
					Name:        "reallocate_gipo",
					Description: `(Optional) Reallocate gipo such that stretched and non stretched BDs have non overlapping gipos.`,
				},
				resource.Attribute{
					Name:        "restrict_infra_vlan_traffic",
					Description: `(Optional) Intra Leaf Communication traffic indicator of object Fabric-Wide Settings Policy.`,
				},
				resource.Attribute{
					Name:        "unicast_xr_ep_learn_disable",
					Description: `(Optional) Disable xrLeanrs indicator of object Fabric-Wide Settings Policy.`,
				},
				resource.Attribute{
					Name:        "validate_overlapping_vlans",
					Description: `(Optional) Validate Overlapping VLANS indicator of object Fabric-Wide Settings Policy.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fex_bundle_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Fex Bundle Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_fex_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI FEX Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_file_remote_path",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Remote Path of a File`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object File Remote Path. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the File Remote Path.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) Authentication Type Choice.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Hostname or IP for export destination of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Transfer protocol to be used for data export of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "remote_path",
					Description: `(Optional) Path where data will reside in the destination of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "remote_port",
					Description: `(Optional) Remote port for data export destination of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional) Username to be used to transfer data to destination of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object File Remote Path.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the File Remote Path.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) Authentication Type Choice.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Hostname or IP for export destination of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Transfer protocol to be used for data export of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "remote_path",
					Description: `(Optional) Path where data will reside in the destination of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "remote_port",
					Description: `(Optional) Remote port for data export destination of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional) Username to be used to transfer data to destination of object File Remote Path.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object File Remote Path.`,
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
					Description: `(Required) Name of Object filter. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Filter.`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Filter.`,
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
					Description: `(Required) Name of Object filter entry. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Filter Entry.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object filter entry.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object filter_entry.`,
				},
				resource.Attribute{
					Name:        "apply_to_frag",
					Description: `(Optional) Flag to determine whether to apply changes to fragment.`,
				},
				resource.Attribute{
					Name:        "arp_opc",
					Description: `(Optional) Open peripheral codes.`,
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
					Description: `(Optional) Ether type for the entry.`,
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
					Description: `(Optional) Name alias for object filter_entry.`,
				},
				resource.Attribute{
					Name:        "prot",
					Description: `(Optional) Level 3 ip protocol.`,
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
					Name:        "description",
					Description: `(Optional) Description for object filter entry.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object filter_entry.`,
				},
				resource.Attribute{
					Name:        "apply_to_frag",
					Description: `(Optional) Flag to determine whether to apply changes to fragment.`,
				},
				resource.Attribute{
					Name:        "arp_opc",
					Description: `(Optional) Open peripheral codes.`,
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
					Description: `(Optional) Ether type for the entry.`,
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
					Description: `(Optional) Name alias for object filter_entry.`,
				},
				resource.Attribute{
					Name:        "prot",
					Description: `(Optional) Level 3 ip protocol.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Description: `(Required) Name of Object firmware_group. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Firmware Group.`,
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
					Description: `(Optional) Name alias for object Firmware Group.`,
				},
				resource.Attribute{
					Name:        "firmware_group_type",
					Description: `(Optional) Component type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Firmware Group.`,
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
					Description: `(Optional) Name alias for object Firmware Group.`,
				},
				resource.Attribute{
					Name:        "firmware_group_type",
					Description: `(Optional) Component type.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
			Type:             "aci_global_security",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Global Security`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Global Security.`,
				},
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
					Description: `(Optional) Session Recording Options.Enables or disables a refresh in the session records. Allowed values are "login", "logout", "refresh", and default value is "7". Type: List.`,
				},
				resource.Attribute{
					Name:        "ui_idle_timeout_seconds",
					Description: `(Optional) GUI Idle Timeout in Seconds.The maximum interval time the GUI remains idle before login needs to be refreshed. Allowed range is 60-65525 and default value is "1200".`,
				},
				resource.Attribute{
					Name:        "webtoken_timeout_seconds",
					Description: `(Optional) Timeout in Seconds.The web token timeout interval. Allowed range is 300-9600 and default value is "600".`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_hsrp_group_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI HSRP Group Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_hsrp_interface_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI HSRP Interface Policy`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_interface_blacklist",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Out of Service Fabric Path`,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Required) The interface name of the interface that need to be disabled. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Out of Service Fabric Path.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Out of Service Fabric Path.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Out of Service Fabric Path.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Out of Service Fabric Path.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_interface_config",
			Category:         "Data Sources",
			ShortDescription: `Manages ACI Access and Fabric Ports is only supported for ACI 5.2(5)+`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "node",
					Description: `(Required) The Node ID of the Port Configuration object. Type: Integer.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The Interface address of the Port Configuration object. The format of the interface value should be 1/1/1 (card/port_id/sub_port) or 1/1 (card/port_id). Type: String.`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `(Optional) The Type of the Port Configuration object. Allowed values are "access", "fabric". Default value is "access". Type: String. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Attribute ID set to the Dn of the Port Configuration.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The Role of the Port Configuration object. Type: String.`,
				},
				resource.Attribute{
					Name:        "policy_group",
					Description: `(Optional) The Distinguished Name of the Policy Group being associated with the Port Configuration object. Type: String.`,
				},
				resource.Attribute{
					Name:        "breakout",
					Description: `(Optional) The Breakout Map of the Port Configuration object. Type: String.`,
				},
				resource.Attribute{
					Name:        "admin_state",
					Description: `(Optional) The Admin State of the Port Configuration object. Type: String.`,
				},
				resource.Attribute{
					Name:        "pc_member",
					Description: `(Optional) The Distinguished Name of the Port Channel Member being associated with the Port Configuration object. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Description of the Port Configuration object. Type: String.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) The Annotation of the Port Configuration object. Type: String.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) The Name Alias of the Port Configuration object. Type: String.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Attribute ID set to the Dn of the Port Configuration.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The Role of the Port Configuration object. Type: String.`,
				},
				resource.Attribute{
					Name:        "policy_group",
					Description: `(Optional) The Distinguished Name of the Policy Group being associated with the Port Configuration object. Type: String.`,
				},
				resource.Attribute{
					Name:        "breakout",
					Description: `(Optional) The Breakout Map of the Port Configuration object. Type: String.`,
				},
				resource.Attribute{
					Name:        "admin_state",
					Description: `(Optional) The Admin State of the Port Configuration object. Type: String.`,
				},
				resource.Attribute{
					Name:        "pc_member",
					Description: `(Optional) The Distinguished Name of the Port Channel Member being associated with the Port Configuration object. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Description of the Port Configuration object. Type: String.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) The Annotation of the Port Configuration object. Type: String.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) The Name Alias of the Port Configuration object. Type: String.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ip_sla_monitoring_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI IP SLA Monitoring Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of the parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the IP SLA Monitoring Policy object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the IP SLA Monitoring Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the IP SLA Monitoring Policy object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the IP SLA Monitoring Policy object.`,
				},
				resource.Attribute{
					Name:        "http_version",
					Description: `(Optional) HTTP Version used for probing. Type: String.`,
				},
				resource.Attribute{
					Name:        "type_of_service",
					Description: `(Optional) Type of Service value for IPv4 packets which provides an indication of the desired Quality of Service (QoS). Allowed range is 0-255 and default value is "0". Type: String.`,
				},
				resource.Attribute{
					Name:        "traffic_class_value",
					Description: `(Optional) Traffic Class Value indicates class or priority of IPv6 packet. Type: String.`,
				},
				resource.Attribute{
					Name:        "request_data_size",
					Description: `(Optional) Minimum size of the IP SLA packet. Type: String.`,
				},
				resource.Attribute{
					Name:        "sla_detect_multiplier",
					Description: `(Optional) Detect Multiplier value for number of missed probes. Type: String.`,
				},
				resource.Attribute{
					Name:        "sla_frequency",
					Description: `(Optional) The SLA frequency value for forwarding packets. Type: String.`,
				},
				resource.Attribute{
					Name:        "sla_port",
					Description: `(Optional) The SLA destination port number. Type: String.`,
				},
				resource.Attribute{
					Name:        "sla_type",
					Description: `(Optional) The IP SLA protocol type. Type: String.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Optional) The threshold value at which the SLA is considered as failed. Type: String.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The amount of time between authentication attempts. Type: String.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the IP SLA Monitoring Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the IP SLA Monitoring Policy object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the IP SLA Monitoring Policy object.`,
				},
				resource.Attribute{
					Name:        "http_version",
					Description: `(Optional) HTTP Version used for probing. Type: String.`,
				},
				resource.Attribute{
					Name:        "type_of_service",
					Description: `(Optional) Type of Service value for IPv4 packets which provides an indication of the desired Quality of Service (QoS). Allowed range is 0-255 and default value is "0". Type: String.`,
				},
				resource.Attribute{
					Name:        "traffic_class_value",
					Description: `(Optional) Traffic Class Value indicates class or priority of IPv6 packet. Type: String.`,
				},
				resource.Attribute{
					Name:        "request_data_size",
					Description: `(Optional) Minimum size of the IP SLA packet. Type: String.`,
				},
				resource.Attribute{
					Name:        "sla_detect_multiplier",
					Description: `(Optional) Detect Multiplier value for number of missed probes. Type: String.`,
				},
				resource.Attribute{
					Name:        "sla_frequency",
					Description: `(Optional) The SLA frequency value for forwarding packets. Type: String.`,
				},
				resource.Attribute{
					Name:        "sla_port",
					Description: `(Optional) The SLA destination port number. Type: String.`,
				},
				resource.Attribute{
					Name:        "sla_type",
					Description: `(Optional) The IP SLA protocol type. Type: String.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Optional) The threshold value at which the SLA is considered as failed. Type: String.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The amount of time between authentication attempts. Type: String.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_isis_domain_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI ISIS Domain Policy and ISIS Level`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the ISIS Domain Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object ISIS Domain Policy and ISIS Level.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object ISIS Domain Policy and ISIS Level.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) Maximum Transmission Unit. The IS-IS Domain policy LSP MTU.`,
				},
				resource.Attribute{
					Name:        "redistrib_metric",
					Description: `(Optional) Metric. Metric used for redistributed routes.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object ISIS Domain Policy and ISIS Level.`,
				},
				resource.Attribute{
					Name:        "lsp_fast_flood",
					Description: `(Optional) The IS-IS Fast-Flooding of LSPs improves Intermediate System-to-Intermediate System (IS-IS) convergence time when new link-state packets (LSPs) are generated in the network and shortest path first (SPF) is triggered by the new LSPs. Allowed values are "disabled" and "enabled".`,
				},
				resource.Attribute{
					Name:        "lsp_gen_init_intvl",
					Description: `(Optional) The LSP generation initial wait interval.`,
				},
				resource.Attribute{
					Name:        "lsp_gen_max_intvl",
					Description: `(Optional) The LSP generation maximum wait interval.`,
				},
				resource.Attribute{
					Name:        "lsp_gen_sec_intvl",
					Description: `(Optional) The LSP generation second wait interval.`,
				},
				resource.Attribute{
					Name:        "spf_comp_init_intvl",
					Description: `(Optional) The SPF computation frequency initial wait interval.`,
				},
				resource.Attribute{
					Name:        "spf_comp_max_intvl",
					Description: `(Optional) The SPF computation frequency maximum wait interval.`,
				},
				resource.Attribute{
					Name:        "spf_comp_sec_intvl",
					Description: `(Optional) The SPF computation frequency second wait interval.`,
				},
				resource.Attribute{
					Name:        "isis_level_name",
					Description: `(Optional) The name of ISIS Level object.`,
				},
				resource.Attribute{
					Name:        "isis_level_type",
					Description: `(Optional) The type of ISIS Level object.`,
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
					Description: `(Required) Name of object L2 Domain. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L2 Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object L2 Domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object L2 Domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L2 Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object L2 Domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object L2 Domain.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Description: `(Required) Name of Object l3 domain profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L3 Domain Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object l3 domain profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object l3 domain profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L3 Domain Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object l3 domain profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object l3 domain profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3_ext_subnet",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI External EPG Subnet`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "external_network_instance_profile_dn",
					Description: `(Required) Distinguished name of the parent External Network Instance Profile object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) IP address of the External EPG Subnet object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the External EPG Subnet object.`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `(Optional) Aggregate Routes of the External EPG Subnet object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the External EPG Subnet object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the External EPG Subnet object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of the External EPG Subnet object.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The list of domain applicable to the capability.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_subnet_to_profile",
					Description: `(Optional) Relation to Route Control Profile (class rtctrlProfile). Type: Block.`,
				},
				resource.Attribute{
					Name:        "tn_rtctrl_profile_dn",
					Description: `(Optional) Associates the External EPGs with the Route Control Profiles.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Relation to configure route map for each BGP peer in the inbound and outbound directions.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_subnet_to_rt_summ",
					Description: `(Optional) Relation to a Route Summarization Policy (class rtsumARtSummPol).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the External EPG Subnet object.`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `(Optional) Aggregate Routes of the External EPG Subnet object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the External EPG Subnet object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the External EPG Subnet object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of the External EPG Subnet object.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The list of domain applicable to the capability.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_subnet_to_profile",
					Description: `(Optional) Relation to Route Control Profile (class rtctrlProfile). Type: Block.`,
				},
				resource.Attribute{
					Name:        "tn_rtctrl_profile_dn",
					Description: `(Optional) Associates the External EPGs with the Route Control Profiles.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Relation to configure route map for each BGP peer in the inbound and outbound directions.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_subnet_to_rt_summ",
					Description: `(Optional) Relation to a Route Summarization Policy (class rtsumARtSummPol).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3_interface_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3 Interface Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object L3 Interface Policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L3 Interface Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object L3 Interface Policy.`,
				},
				resource.Attribute{
					Name:        "bfd_isis",
					Description: `(Optional) BFD ISIS Configuration for object L3 Interface Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object L3 Interface Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object L3 Interface Policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L3 Interface Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object L3 Interface Policy.`,
				},
				resource.Attribute{
					Name:        "bfd_isis",
					Description: `(Optional) BFD ISIS Configuration for object L3 Interface Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object L3 Interface Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object L3 Interface Policy.`,
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
					Description: `(Required) Distinguished name of the parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the L3 Outside object. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L3 Outside.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the L3 Outside object.`,
				},
				resource.Attribute{
					Name:        "enforce_rtctrl",
					Description: `(Optional) Enforce route control type of the L3 Outside object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of the L3 Outside object.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the L3 Outside object.`,
				},
				resource.Attribute{
					Name:        "mpls_enabled",
					Description: `(Optional) Indiscate whether MPLS is enabled or not.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_dampening_pol",
					Description: `(Optional) Relation to Route Control Profile for Dampening Policies (class rtctrlProfile). Cardinality - N_TO_M. Type - Block.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_ectx",
					Description: `(Optional) Relation to VRF (class fvCtx). Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_interleak_pol",
					Description: `(Optional) Relation to Route Profile for Interleak (class rtctrlProfile). Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_l3_dom_att",
					Description: `(Optional) Relation to a L3 Domain (class extnwDomP). Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_l3extrs_redistribute_pol",
					Description: `(Optional) A block representing the relation to a Route Profile for Redistribution (class rtctrlProfile). Cardinality - N_TO_M. Type: Block.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Route Map Source for the Route Profile for Redistribution.`,
				},
				resource.Attribute{
					Name:        "target_dn",
					Description: `(Optional) Distinguished name of the Route Control Profile for the Route Profile for Redistribution.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L3 Outside.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the L3 Outside object.`,
				},
				resource.Attribute{
					Name:        "enforce_rtctrl",
					Description: `(Optional) Enforce route control type of the L3 Outside object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias of the L3 Outside object.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) The target differentiated services code point (DSCP) of the path attached to the L3 Outside object.`,
				},
				resource.Attribute{
					Name:        "mpls_enabled",
					Description: `(Optional) Indiscate whether MPLS is enabled or not.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_dampening_pol",
					Description: `(Optional) Relation to Route Control Profile for Dampening Policies (class rtctrlProfile). Cardinality - N_TO_M. Type - Block.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_ectx",
					Description: `(Optional) Relation to VRF (class fvCtx). Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_interleak_pol",
					Description: `(Optional) Relation to Route Profile for Interleak (class rtctrlProfile). Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_l3_dom_att",
					Description: `(Optional) Relation to a L3 Domain (class extnwDomP). Cardinality - N_TO_ONE. Type - String.`,
				},
				resource.Attribute{
					Name:        "relation_l3extrs_redistribute_pol",
					Description: `(Optional) A block representing the relation to a Route Profile for Redistribution (class rtctrlProfile). Cardinality - N_TO_M. Type: Block.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Route Map Source for the Route Profile for Redistribution.`,
				},
				resource.Attribute{
					Name:        "target_dn",
					Description: `(Optional) Distinguished name of the Route Control Profile for the Route Profile for Redistribution.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_bfd_interface_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3out BFD Interface Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_bgp_external_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3-out BGP External Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_bgp_protocol_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3out BGP Protocol Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_floating_svi",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3out Floating SVI`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_interface_profile_dn",
					Description: `(Required) Distinguished name of the parent logical interface profile object.`,
				},
				resource.Attribute{
					Name:        "node_dn",
					Description: `(Required) Node DN of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Required) Port encapsulation of the L3out floating SVI object. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Logical Interface Profile.`,
				},
				resource.Attribute{
					Name:        "addr",
					Description: `(Optional) Peer address of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "autostate",
					Description: `(Optional) Autostate of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "encap_scope",
					Description: `(Optional) Encap scope of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "if_inst_t",
					Description: `(Optional) Interface type of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "ipv6_dad",
					Description: `(Optional) IPv6 dad of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "ll_addr",
					Description: `(Optional) Link local address address of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Optional) MAC address of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) BGP domain mode of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) Target DSCP of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_dyn_path_att",
					Description: `(Optional) A block representing the relation to a Domain (class infraDomP or vmmDomP). Type: Block.`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Required) The distinguished name of the target.`,
				},
				resource.Attribute{
					Name:        "floating_address",
					Description: `(Optional) Floating address of the target.`,
				},
				resource.Attribute{
					Name:        "forged_transmit",
					Description: `(Optional) A configuration option that allows virtual machines to send frames with a mac address that is different from the one specified in the virtual-nic adapter configuration.`,
				},
				resource.Attribute{
					Name:        "mac_change",
					Description: `(Optional) The status of the mac address change support of the port groups in an external VMM controller, such as a vCenter.`,
				},
				resource.Attribute{
					Name:        "promiscuous_mode",
					Description: `(Optional) The status of the promiscuous mode support status of the port groups in an external VMM controller, such as a vCenter. This needs to be turned on only for service devices in the cloud, not for Enterprise AVE service deployments.`,
				},
				resource.Attribute{
					Name:        "enhanced_lag_policy_dn",
					Description: `(Optional) The distinguished name of the target enhanced lag policy (class lacpEnhancedLagPol).`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) Access port encapsulation (VLAN) of the target (format: vlan-101).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Logical Interface Profile.`,
				},
				resource.Attribute{
					Name:        "addr",
					Description: `(Optional) Peer address of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "autostate",
					Description: `(Optional) Autostate of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "encap_scope",
					Description: `(Optional) Encap scope of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "if_inst_t",
					Description: `(Optional) Interface type of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "ipv6_dad",
					Description: `(Optional) IPv6 dad of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "ll_addr",
					Description: `(Optional) Link local address address of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Optional) MAC address of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) BGP domain mode of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "target_dscp",
					Description: `(Optional) Target DSCP of the L3out floating SVI object.`,
				},
				resource.Attribute{
					Name:        "relation_l3ext_rs_dyn_path_att",
					Description: `(Optional) A block representing the relation to a Domain (class infraDomP or vmmDomP). Type: Block.`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Required) The distinguished name of the target.`,
				},
				resource.Attribute{
					Name:        "floating_address",
					Description: `(Optional) Floating address of the target.`,
				},
				resource.Attribute{
					Name:        "forged_transmit",
					Description: `(Optional) A configuration option that allows virtual machines to send frames with a mac address that is different from the one specified in the virtual-nic adapter configuration.`,
				},
				resource.Attribute{
					Name:        "mac_change",
					Description: `(Optional) The status of the mac address change support of the port groups in an external VMM controller, such as a vCenter.`,
				},
				resource.Attribute{
					Name:        "promiscuous_mode",
					Description: `(Optional) The status of the promiscuous mode support status of the port groups in an external VMM controller, such as a vCenter. This needs to be turned on only for service devices in the cloud, not for Enterprise AVE service deployments.`,
				},
				resource.Attribute{
					Name:        "enhanced_lag_policy_dn",
					Description: `(Optional) The distinguished name of the target enhanced lag policy (class lacpEnhancedLagPol).`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) Access port encapsulation (VLAN) of the target (format: vlan-101).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_hsrp_interface_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI HSRP Interface Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_hsrp_interface_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3-out HSRP interface profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_hsrp_secondary_vip",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3out HSRP Secondary VIP`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_loopback_interface_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Loop Back Interface Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fabric_node_dn",
					Description: `(Required) Distinguished name of parent Fabric Node object.`,
				},
				resource.Attribute{
					Name:        "addr",
					Description: `(Required) Address of L3out lookback interface profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L3out lookback interface profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for L3out lookback interface profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `Annotation for L3out lookback interface profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `Name alias for L3out lookback interface profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L3out lookback interface profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for L3out lookback interface profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `Annotation for L3out lookback interface profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `Name alias for L3out lookback interface profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_ospf_external_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3-out OSPF External Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_ospf_interface_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3out OSPF Interface Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_path_attachment",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3-out Path Attachment`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_path_attachment_secondary_ip",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3-out Path Attachment Secondary IP`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_route_tag_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3out Route Tag Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_static_route",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3out Static Route`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_static_route_next_hop",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3out Static Route Next Hop`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l3out_vpc_member",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L3out VPC Member`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l4_l7_deployed_graph_connector_vlan",
			Category:         "Data Sources",
			ShortDescription: `Data Source for ACI L4-L7 Deployed Graph Connector VLAN`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_context_dn",
					Description: `(Required) Distinguished name of the Logical Interface Context.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of EPgDef.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the EPgDef object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `Annotation of the EPgDef object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `Name Alias of the EPgDef object.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `The VLAN encapsulation tag.`,
				},
				resource.Attribute{
					Name:        "fabric_encap",
					Description: `The VXLAN encapsulation tag.`,
				},
				resource.Attribute{
					Name:        "delete_pbr_scenario",
					Description: `The boolean value for deleting Policy Based Routing.`,
				},
				resource.Attribute{
					Name:        "member_type",
					Description: `The type of member for the EPgDef object.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `The IP address of the routing device.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of EPgDef.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the EPgDef object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `Annotation of the EPgDef object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `Name Alias of the EPgDef object.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `The VLAN encapsulation tag.`,
				},
				resource.Attribute{
					Name:        "fabric_encap",
					Description: `The VXLAN encapsulation tag.`,
				},
				resource.Attribute{
					Name:        "delete_pbr_scenario",
					Description: `The boolean value for deleting Policy Based Routing.`,
				},
				resource.Attribute{
					Name:        "member_type",
					Description: `The type of member for the EPgDef object.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `The IP address of the routing device.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l4_l7_devices",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L4-L7 Device`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of the parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the L4-L7 Device. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L4-L7 Device.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the L4-L7 Device object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the L4-L7 Device object.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Enables L4-L7 device cluster to operate in active/active mode. Allowed values are "no", "yes", and default value is "no". Type: String.`,
				},
				resource.Attribute{
					Name:        "context_aware",
					Description: `(Optional) Determines if the L4-L7 device cluster supports multiple contexts (VRFs). Allowed values are "multi-Context", "single-Context", and default value is "single-Context". Type: String.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `(Optional) Device Type. Allowed values are "CLOUD", "PHYSICAL", "VIRTUAL", and default value is "PHYSICAL". Type: String.`,
				},
				resource.Attribute{
					Name:        "function_type",
					Description: `(Optional) Function Type of the L4-L7 device cluster. Allowed values are "GoThrough", "GoTo", "L1", "L2", "None", and default value is "GoTo". Type: String.`,
				},
				resource.Attribute{
					Name:        "is_copy",
					Description: `(Optional) Sets device as a copy device. Allowed values are "no", "yes", and default value is "no". Type: String.`,
				},
				resource.Attribute{
					Name:        "promiscuous_mode",
					Description: `(Optional) Enabling Promiscuous Mode to allow all the traffic in a port group to reach a VM attached to a promiscuous port. Allowed values are "no", "yes", and default value is "no". Type: String.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Optional) The type of service the L4-L7 device performs. Allowed values are "ADC", "COPY", "FW", "NATIVELB", "OTHERS", and default value is "OTHERS". Type: String.`,
				},
				resource.Attribute{
					Name:        "trunking",
					Description: `(Optional) Configures the device port group for trunking of virtual devices. Allowed values are "no", "yes", and default value is "no". Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_vns_rs_al_dev_to_dom_p",
					Description: `(Optional) Represents a relation from L4-L7 Device to a VMM Domain Profile (class vmmDomP). Type: Block.`,
				},
				resource.Attribute{
					Name:        "domain_dn",
					Description: `(Optional) Distinguished name of the VMM Domain in which the VM is deployed. Type: String.`,
				},
				resource.Attribute{
					Name:        "switching_mode",
					Description: `(Optional) Port group switching mode. Allowed values are "native", "AVE", and default value is "native". The value "AVE" is not supported with non-AVE VMM Domain. Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_vns_rs_al_dev_to_phys_dom_p",
					Description: `(Optional) Represents a relation from L4-L7 Device to a Physical Domain Profile (class physDomP). Type: String.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L4-L7 Device.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the L4-L7 Device object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the L4-L7 Device object.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Enables L4-L7 device cluster to operate in active/active mode. Allowed values are "no", "yes", and default value is "no". Type: String.`,
				},
				resource.Attribute{
					Name:        "context_aware",
					Description: `(Optional) Determines if the L4-L7 device cluster supports multiple contexts (VRFs). Allowed values are "multi-Context", "single-Context", and default value is "single-Context". Type: String.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `(Optional) Device Type. Allowed values are "CLOUD", "PHYSICAL", "VIRTUAL", and default value is "PHYSICAL". Type: String.`,
				},
				resource.Attribute{
					Name:        "function_type",
					Description: `(Optional) Function Type of the L4-L7 device cluster. Allowed values are "GoThrough", "GoTo", "L1", "L2", "None", and default value is "GoTo". Type: String.`,
				},
				resource.Attribute{
					Name:        "is_copy",
					Description: `(Optional) Sets device as a copy device. Allowed values are "no", "yes", and default value is "no". Type: String.`,
				},
				resource.Attribute{
					Name:        "promiscuous_mode",
					Description: `(Optional) Enabling Promiscuous Mode to allow all the traffic in a port group to reach a VM attached to a promiscuous port. Allowed values are "no", "yes", and default value is "no". Type: String.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Optional) The type of service the L4-L7 device performs. Allowed values are "ADC", "COPY", "FW", "NATIVELB", "OTHERS", and default value is "OTHERS". Type: String.`,
				},
				resource.Attribute{
					Name:        "trunking",
					Description: `(Optional) Configures the device port group for trunking of virtual devices. Allowed values are "no", "yes", and default value is "no". Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_vns_rs_al_dev_to_dom_p",
					Description: `(Optional) Represents a relation from L4-L7 Device to a VMM Domain Profile (class vmmDomP). Type: Block.`,
				},
				resource.Attribute{
					Name:        "domain_dn",
					Description: `(Optional) Distinguished name of the VMM Domain in which the VM is deployed. Type: String.`,
				},
				resource.Attribute{
					Name:        "switching_mode",
					Description: `(Optional) Port group switching mode. Allowed values are "native", "AVE", and default value is "native". The value "AVE" is not supported with non-AVE VMM Domain. Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_vns_rs_al_dev_to_phys_dom_p",
					Description: `(Optional) Represents a relation from L4-L7 Device to a Physical Domain Profile (class physDomP). Type: String.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l4_l7_logical_interface",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L4-L7 Logical Interface`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "l4-l7_devices_dn",
					Description: `(Required) Distinguished name of the parent L4-L7 Device object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Logical Interface object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Logical Interface.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Logical Interface object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Logical Interface object.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) The port encapsulation to be used with the device. Type: String.`,
				},
				resource.Attribute{
					Name:        "lag_policy_name",
					Description: `(Optional) Name of the enhanced Lag policy. Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_vns_rs_c_if_att_n",
					Description: `(Optional) Represents the relation between a set of Concrete Interfaces and the Device Cluster (class vnsCIf). Type: List.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Logical Interface.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Logical Interface object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Logical Interface object.`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `(Optional) The port encapsulation to be used with the device. Type: String.`,
				},
				resource.Attribute{
					Name:        "lag_policy_name",
					Description: `(Optional) Name of the enhanced Lag policy. Type: String.`,
				},
				resource.Attribute{
					Name:        "relation_vns_rs_c_if_att_n",
					Description: `(Optional) Represents the relation between a set of Concrete Interfaces and the Device Cluster (class vnsCIf). Type: List.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_l4_l7_redirect_health_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI L4-L7 Redirect Health Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of the parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the L4-L7 Redirect Health Group object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L4-L7 Redirect Health Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the L4-L7 Redirect Health Group object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the L4-L7 Redirect Health Group object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the L4-L7 Redirect Health Group object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the L4-L7 Redirect Health Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the L4-L7 Redirect Health Group object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the L4-L7 Redirect Health Group object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the L4-L7 Redirect Health Group object.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_lacp_member_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI LACP Member Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of LACP Member Policy object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LACP Member Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the LACP Member Policy object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the LACP Member Policy object.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Port priority - LACP uses the port priority to decide which ports should be put in standby mode when there is a limitation that prevents all compatible ports from aggregating and which ports should be put into active mode. A higher port priority value means a lower priority for LACP.`,
				},
				resource.Attribute{
					Name:        "transmit_rate",
					Description: `(Optional) Transmission Rate. The configured transmit rate of the LACP packets.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LACP Member Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the LACP Member Policy object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the LACP Member Policy object.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Port priority - LACP uses the port priority to decide which ports should be put in standby mode when there is a limitation that prevents all compatible ports from aggregating and which ports should be put into active mode. A higher port priority value means a lower priority for LACP.`,
				},
				resource.Attribute{
					Name:        "transmit_rate",
					Description: `(Optional) Transmission Rate. The configured transmit rate of the LACP packets.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ldap_group_map",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI LDAP Group Map`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object LDAP Group Map.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of object LDAP Group Map. Allowed values are "duo" and "ldap". ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LDAP Group Map.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object LDAP Group Map.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object LDAP Group Map.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object LDAP Group Map.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LDAP Group Map.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object LDAP Group Map.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object LDAP Group Map.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object LDAP Group Map.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ldap_group_map_rule",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI LDAP Group Map Rule`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object LDAP Group Map Rule.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of object LDAP Group MAp Rule. Allowed Values are "duo" and "ldap". ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LDAP Group Map Rule.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object LDAP Group Map Rule.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object LDAP Group Map Rule.`,
				},
				resource.Attribute{
					Name:        "groupdn",
					Description: `(Optional) LDAP Group DN to compare with LDAP search query for user's membership.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object LDAP Group Map Rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LDAP Group Map Rule.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object LDAP Group Map Rule.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object LDAP Group Map Rule.`,
				},
				resource.Attribute{
					Name:        "groupdn",
					Description: `(Optional) LDAP Group DN to compare with LDAP search query for user's membership.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object LDAP Group Map Rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ldap_group_map_rule_to_group_map",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI LDAP Group Map Rule to Group Map Ref`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ldap_group_map_dn",
					Description: `(Required) Distinguished name of parent LDAP Group Map object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object LDAP Group Map Rule to Group Map Ref. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LDAP Group Map Rule to Group Map Ref.`,
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
					Description: `(Optional) Description of object LDAP Group Map Rule to Group Map Ref.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LDAP Group Map Rule to Group Map Ref.`,
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
					Description: `(Optional) Description of object LDAP Group Map Rule to Group Map Ref.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ldap_provider",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI LDAP Provider`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Host name or IP address of object LDAP Provider.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of LDAP Provider. Allowed values are "ldap" and "duo". ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LDAP Provider.`,
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
					Description: `(Optional) The LDAP Server SSL Certificate validation level.`,
				},
				resource.Attribute{
					Name:        "attribute",
					Description: `(Optional) The attribute to be downloaded that contains user role and domain information.`,
				},
				resource.Attribute{
					Name:        "basedn",
					Description: `(Optional) The LDAP base DN to be used in a user search.`,
				},
				resource.Attribute{
					Name:        "enable_ssl",
					Description: `(Optional) A property for enabling an SSL connection with the LDAP provider.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) The LDAP filter to be used in a user search.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The service port number for the LDAP service.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Retry count of object LDAP Provider.`,
				},
				resource.Attribute{
					Name:        "rootdn",
					Description: `(Optional) The root DN or bind DN of the LDAP provider.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The timeout for communication with an LDAP provider server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LDAP Provider.`,
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
					Description: `(Optional) The LDAP Server SSL Certificate validation level.`,
				},
				resource.Attribute{
					Name:        "attribute",
					Description: `(Optional) The attribute to be downloaded that contains user role and domain information.`,
				},
				resource.Attribute{
					Name:        "basedn",
					Description: `(Optional) The LDAP base DN to be used in a user search.`,
				},
				resource.Attribute{
					Name:        "enable_ssl",
					Description: `(Optional) A property for enabling an SSL connection with the LDAP provider.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) The LDAP filter to be used in a user search.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The service port number for the LDAP service.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Retry count of object LDAP Provider.`,
				},
				resource.Attribute{
					Name:        "rootdn",
					Description: `(Optional) The root DN or bind DN of the LDAP provider.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The timeout for communication with an LDAP provider server.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_access_bundle_policy_sub_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Override Policy Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "leaf_access_bundle_policy_group_dn",
					Description: `(Required) Distinguished name of the parent infraAccBndlGrp object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Override Policy Group object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Override Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Override Policy Group object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Override Policy Group object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Override Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Override Policy Group object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Override Policy Group object.`,
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
					Description: `(Required) Name of Object leaf_access_port_policy_group. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Leaf Access Port Policy Group.`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Leaf Access Port Policy Group.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_leaf_breakout_port_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Leaf Breakout Port Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Description: `(Required) Name of Object Leaf Interface Profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Leaf Interface Profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Leaf Interface Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Leaf Interface Profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Leaf Interface Profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Leaf Interface Profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Leaf Interface Profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Leaf Interface Profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object Leaf Interface Profile.`,
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
					Name:        "description",
					Description: `(Optional) Description for object leaf profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object leaf profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object leaf profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Leaf Profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object leaf profile.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object leaf profile.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object leaf profile.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Description: `(Required) Name of Object LLDP Interface Policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LLDP Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_rx_st",
					Description: `(Optional) Admin receive state.`,
				},
				resource.Attribute{
					Name:        "admin_tx_st",
					Description: `(Optional) Admin transmit state.`,
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
					Description: `(Optional) Name alias for object LLDP Interface Policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the LLDP Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_rx_st",
					Description: `(Optional) Admin receive state.`,
				},
				resource.Attribute{
					Name:        "admin_tx_st",
					Description: `(Optional) Admin transmit state.`,
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
					Description: `(Optional) Name alias for object LLDP Interface Policy.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_logical_node_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Logical Node Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Name:        "tdn",
					Description: `(Required) Tdn of Object Fabric Node. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Fabric Node.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "config_issues",
					Description: `(Optional) Configuration issues`,
				},
				resource.Attribute{
					Name:        "rtr_id",
					Description: `(Optional) Router identifier`,
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
					Description: `(Optional) Annotation for object Fabric Node.`,
				},
				resource.Attribute{
					Name:        "config_issues",
					Description: `(Optional) Configuration issues`,
				},
				resource.Attribute{
					Name:        "rtr_id",
					Description: `(Optional) Router identifier`,
				},
				resource.Attribute{
					Name:        "rtr_id_loop_back",
					Description: `(Optional)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_login_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Login Domain`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object Login Domain. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Login Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Login Domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Login Domain.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Login Domain.`,
				},
				resource.Attribute{
					Name:        "provider_group",
					Description: `(Optional) Provider Group. An AAA configuration provider group is a group of remote servers supporting the same AAA protocol that will be used for authentication and authorization. When a provider group is specified, only the servers within that group will be used for authentication and authorization. If no provider group is specified, all servers supporting the realm of AAA protocols will be used for authentication and authorization.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `(Optional) Realm. The security method for processing authentication requests. The realm allows the protected resources on the associated server to be partitioned into a set of protection spaces, each with its own authentication authorization database.`,
				},
				resource.Attribute{
					Name:        "realm_sub_type",
					Description: `(Optional) Realm subtype of object Login Domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Login Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Login Domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Login Domain.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Login Domain.`,
				},
				resource.Attribute{
					Name:        "provider_group",
					Description: `(Optional) Provider Group. An AAA configuration provider group is a group of remote servers supporting the same AAA protocol that will be used for authentication and authorization. When a provider group is specified, only the servers within that group will be used for authentication and authorization. If no provider group is specified, all servers supporting the realm of AAA protocols will be used for authentication and authorization.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `(Optional) Realm. The security method for processing authentication requests. The realm allows the protected resources on the associated server to be partitioned into a set of protection spaces, each with its own authentication authorization database.`,
				},
				resource.Attribute{
					Name:        "realm_sub_type",
					Description: `(Optional) Realm subtype of object Login Domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_login_domain_provider",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Login Domain Provider`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "parent_dn",
					Description: `(Required) Distinguished name of parent.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object Login Domain Provider. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Login Domain Provider.`,
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
					Description: `(Optional) Order in which Providers are Tried. The relative priority in which the AAA provider will be contacted within the provider group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Login Domain Provider.`,
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
					Description: `(Optional) Order in which Providers are Tried. The relative priority in which the AAA provider will be contacted within the provider group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_maintenance_group_node",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Maintenance Group Node`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pod_maintenance_group_dn",
					Description: `(Required) Distinguished name of parent POD Maintenance Group Object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Maintenance Group Node Object. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the dn of the Maintenance Group Node Object.`,
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
					Description: `(Optional) From for Maintenance Group Node Object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for Maintenance Group Node Object.`,
				},
				resource.Attribute{
					Name:        "to_",
					Description: `(Optional) To for Maintenance Group Node Object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the dn of the Maintenance Group Node Object.`,
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
					Description: `(Optional) From for Maintenance Group Node Object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for Maintenance Group Node Object.`,
				},
				resource.Attribute{
					Name:        "to_",
					Description: `(Optional) To for Maintenance Group Node Object.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_managed_node_connectivity_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Managed Node Connectivity Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object Managed Node Connectivity Group. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Managed Node Connectivity Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Managed Node Connectivity Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Managed Node Connectivity Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Managed Node Connectivity Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_match_community_terms",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Match Community Term`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "match_rule_dn",
					Description: `(Required) Distinguished name of the parent Match Rule object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Match Community Term object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Match Community Term.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Match Community Term object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Match Community Term object.`,
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
					Name:        "community",
					Description: `(Optional) The community of the Match Community Factor object. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the Match Community Factor object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Match Community Term.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Match Community Term object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Match Community Term object.`,
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
					Name:        "community",
					Description: `(Optional) The community of the Match Community Factor object. Type: String.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the Match Community Factor object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_match_regex_community_terms",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Match Rule Based on Community Regular Expression`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "match_rule_dn",
					Description: `(Required) Distinguished name of parent Match Rule object.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Required) Community Type of the Match Rule Based on Community Regular Expression object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Match Rule Based on Community Regular Expression.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Match Rule Based on Community Regular Expression object.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community Type of the Match Rule Based on Community Regular Expression object.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) Regular Expression. A regular expression used to specify a pattern to match against the community string.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Match Rule Based on Community Regular Expression.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Match Rule Based on Community Regular Expression object.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community Type of the Match Rule Based on Community Regular Expression object.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) Regular Expression. A regular expression used to specify a pattern to match against the community string.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_match_route_destination_rule",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Match Route Destination Rule`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "match_rule_dn",
					Description: `(Required) Distinguished name of parent MatchRule object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) ip of object Match Route Destination Rule. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Match Route Destination Rule.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Match Route Destination Rule.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Match Route Destination Rule.`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `(Optional) Aggregated Route. Aggregated Route`,
				},
				resource.Attribute{
					Name:        "greater_than_mask",
					Description: `(Optional) Start of Prefix Length. Prefix list range`,
				},
				resource.Attribute{
					Name:        "less_than_mask",
					Description: `(Optional) End of Prefix Length.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Match Route Destination Rule.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Match Route Destination Rule.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Match Route Destination Rule.`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `(Optional) Aggregated Route. Aggregated Route`,
				},
				resource.Attribute{
					Name:        "greater_than_mask",
					Description: `(Optional) Start of Prefix Length. Prefix list range`,
				},
				resource.Attribute{
					Name:        "less_than_mask",
					Description: `(Optional) End of Prefix Length.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_match_rule",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Match Rule`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object Match Rule. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Match Rule.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Match Rule.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Match Rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Match Rule.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Match Rule.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Match Rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_mcp_instance_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI MCP Instance Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the MCP Instance Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object MCP Instance Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object MCP Instance Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of Object MCP Instance Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Admin State. The administrative state of the object or policy.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) Controls. The control state.`,
				},
				resource.Attribute{
					Name:        "init_delay_time",
					Description: `(Optional) Init Delay Time.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(optional) Secret Key. The key or password used to uniquely identify this configuration object.`,
				},
				resource.Attribute{
					Name:        "loop_detect_mult",
					Description: `(Optional) Loop Detection Multiplier.`,
				},
				resource.Attribute{
					Name:        "loop_protect_act",
					Description: `(Optional) Loop Protection Action.`,
				},
				resource.Attribute{
					Name:        "tx_freq",
					Description: `(Optional) Transmission Frequency. Sets the transmission frequency of the instance advertisements.`,
				},
				resource.Attribute{
					Name:        "tx_freq_msec",
					Description: `(Optional) Transmission Frequency. Sets the transmission frequency of mcp advertisements in milliseconds`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_mgmt_preference",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Mgmt Preference`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Mgmt preference.`,
				},
				resource.Attribute{
					Name:        "interface_pref",
					Description: `(Optional) Management interface that has to be used.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Mgmt preference.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Mgmt preference.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Mgmt Preference.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_mgmt_zone",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Management Zone`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Description: `(Optional) Administrative state of the object or policy.`,
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
					Description: `(Optional) Name alias for object miscabling protocol interface policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Mis-cabling Protocol Interface Policy.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Administrative state of the object or policy.`,
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
					Description: `(Optional) Name alias for object miscabling protocol interface policy.`,
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
					Description: `(Required) Tenant dn of object monitoring policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object monitoring policy. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the object monitoring policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object monitoring policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object monitoring policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object monitoring policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the object monitoring policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object monitoring policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object monitoring policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object monitoring policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_multicast_pool",
			Category:         "Data Sources",
			ShortDescription: `Data source for the ACI Multicast Address Pool`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Multicast Address Pool.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Multicast Address Pool.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Multicast Address Pool.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Multicast Address Pool.`,
				},
				resource.Attribute{
					Name:        "multicast_address_block",
					Description: `(Optional) Multicast Address Pool Block Configuration. Type: block.`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `(Optional) First multicast IP of the Multicast Address Pool Block.`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `(Optional) Last multicast IP of the Multicast Address Pool Block.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name Alias of the Multicast Address Pool Block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Multicast Address Pool Block.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Multicast Address Pool Block.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Multicast Address Pool Block. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Multicast Address Pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Multicast Address Pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_multicast_pool_block",
			Category:         "Data Sources",
			ShortDescription: `Data source for the ACI Multicast Address Pool Block`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "multicast_pool_dn",
					Description: `(Required) Distinguished name of the parent Multicast Pool object.`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `(Required) First multicast IP of the Multicast Address Pool Block.`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `(Required) Last multicast IP of the Multicast Address Pool Block.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Multicast Address Pool Block.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Multicast Address Pool Block.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Multicast Address Pool Block. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Multicast Address Pool Block.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Multicast Address Pool Block.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Description: `(Required) Distinguished name of parent Firmware Group object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object Node Block. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Node Block.`,
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
					Description: `(Optional) From value for Object Node Block.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for Object Node Block..`,
				},
				resource.Attribute{
					Name:        "to_",
					Description: `(Optional) To value for Object Node Block.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Node Block.`,
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
					Description: `(Optional) From value for Object Node Block.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for Object Node Block..`,
				},
				resource.Attribute{
					Name:        "to_",
					Description: `(Optional) To value for Object Node Block.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ospf_route_summarization",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI OSPF Route Summarization`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ospf_timers",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI OSPF Timers`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_pbr_l1_l2_destination",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Destination of L1/L2 Redirected Traffic`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_based_redirect_dn",
					Description: `(Required) Distinguished name of the parent Policy-Based (Redirect or Redirect Backup) object.`,
				},
				resource.Attribute{
					Name:        "destination_name",
					Description: `(Required) Destination Name of the destination of the L1/L2 Redirected Traffic. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the destination of the L1/L2 Redirected Traffic.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the destination of the L1/L2 Redirected Traffic.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the destination of the L1/L2 Redirected Traffic.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the destination of the L1/L2 Redirected Traffic.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Optional) MAC Address of the destination of the L1/L2 Redirected Traffic.`,
				},
				resource.Attribute{
					Name:        "pod_id",
					Description: `(Optional) Pod Id of the destination of the L1/L2 Redirected Traffic.`,
				},
				resource.Attribute{
					Name:        "relation_vns_rs_l1_l2_redirect_health_group",
					Description: `(Optional) Represents the relation to a L4-L7 Redirect Health Group (class vnsRedirectHealthGroup).`,
				},
				resource.Attribute{
					Name:        "relation_vns_rs_to_c_if",
					Description: `(Optional) Represents the relation to a Concrete Interface (class vnsCIf).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the destination of the L1/L2 Redirected Traffic.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the destination of the L1/L2 Redirected Traffic.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the destination of the L1/L2 Redirected Traffic.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the destination of the L1/L2 Redirected Traffic.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Optional) MAC Address of the destination of the L1/L2 Redirected Traffic.`,
				},
				resource.Attribute{
					Name:        "pod_id",
					Description: `(Optional) Pod Id of the destination of the L1/L2 Redirected Traffic.`,
				},
				resource.Attribute{
					Name:        "relation_vns_rs_l1_l2_redirect_health_group",
					Description: `(Optional) Represents the relation to a L4-L7 Redirect Health Group (class vnsRedirectHealthGroup).`,
				},
				resource.Attribute{
					Name:        "relation_vns_rs_to_c_if",
					Description: `(Optional) Represents the relation to a Concrete Interface (class vnsCIf).`,
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
					Description: `(Required) Name of object physical domain. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Physical Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object physical domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object physical domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Physical Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for object physical domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object physical domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_pod_maintenance_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI POD Maintenance Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_port_security_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Port Security Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_port_tracking",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Port Tracking`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Port Tracking.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Port Tracking.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Port Tracking.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Port Tracking State. The administrative state of the object or policy.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Optional) Delay Timeout. The administrative port delay.`,
				},
				resource.Attribute{
					Name:        "include_apic_ports",
					Description: `(Optional) Include APIC Ports when port tracking is triggered.`,
				},
				resource.Attribute{
					Name:        "minlinks",
					Description: `(Optional) Minimum links left up before trigger.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Port Tracking.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Port Tracking.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Port Tracking.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Port Tracking.`,
				},
				resource.Attribute{
					Name:        "admin_st",
					Description: `(Optional) Port Tracking State. The administrative state of the object or policy.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Optional) Delay Timeout. The administrative port delay.`,
				},
				resource.Attribute{
					Name:        "include_apic_ports",
					Description: `(Optional) Include APIC Ports when port tracking is triggered.`,
				},
				resource.Attribute{
					Name:        "minlinks",
					Description: `(Optional) Minimum links left up before trigger.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Port Tracking.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_qos_instance_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI QOS Instance Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the QOS Instance Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object QOS Instance Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object QOS Instance Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object QOS Instance Policy`,
				},
				resource.Attribute{
					Name:        "etrap_age_timer",
					Description: `(Optional) E-trap flow age out timer.`,
				},
				resource.Attribute{
					Name:        "etrap_bw_thresh",
					Description: `(Optional) Track activeness of elephant flow.`,
				},
				resource.Attribute{
					Name:        "etrap_byte_ct",
					Description: `(Optional) E-trap elephant flow identifier.`,
				},
				resource.Attribute{
					Name:        "etrap_st",
					Description: `(Optional) E-trap enable knob. E-trap parameters`,
				},
				resource.Attribute{
					Name:        "fabric_flush_interval",
					Description: `(Optional) Fabric Flush Interval in ms.`,
				},
				resource.Attribute{
					Name:        "fabric_flush_st",
					Description: `(Optional) Fabric PFC Flush enable knob. Fabric Flush parameters`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) Global Control Settings. The control state.`,
				},
				resource.Attribute{
					Name:        "uburst_spine_queues",
					Description: `(Optional) Micro burst spine queues percent.`,
				},
				resource.Attribute{
					Name:        "uburst_tor_queues",
					Description: `(Optional) Micro burst tor queues percent.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_radius_provider",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI RADIUS Provider`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Host name or IP address of object RADIUS Provider.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of object RADIUS Provider. Allowed values are "duo" and "radius". ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the RADIUS Provider.`,
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
					Description: `(Optional) The service port number for RADIUS service.`,
				},
				resource.Attribute{
					Name:        "auth_protocol",
					Description: `(Optional) The RADIUS authentication protocol.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Number of retries for a for communication with a RADIUS provider server.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The timeout for communication with a RADIUS provider server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the RADIUS Provider.`,
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
					Description: `(Optional) The service port number for RADIUS service.`,
				},
				resource.Attribute{
					Name:        "auth_protocol",
					Description: `(Optional) The RADIUS authentication protocol.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Number of retries for a for communication with a RADIUS provider server.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The timeout for communication with a RADIUS provider server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_radius_provider_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI RADIUS Provider Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object RADIUS Provider Group. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the RADIUS Provider Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object RADIUS Provider Group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object RADIUS Provider Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object RADIUS Provider Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the RADIUS Provider Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object RADIUS Provider Group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object RADIUS Provider Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object RADIUS Provider Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_ranges",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI VLAN Pool Ranges`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vlan_pool_dn",
					Description: `(Required) Distinguished name of parent VLAN Pool object.`,
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
					Description: `(Optional) Annotation for object ranges.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object ranges.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object ranges.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) System role type`,
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
					Description: `(Optional) Annotation for object ranges.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object ranges.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object ranges.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) System role type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_recurring_window",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Recurring Window`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scheduler_dn",
					Description: `(Required) Distinguished name of parent Scheduler object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object Recurring Window. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Recurring Window.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Recurring Window.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Recurring Window.`,
				},
				resource.Attribute{
					Name:        "concur_cap",
					Description: `(Optional) Maximum Concurrent Tasks. The concurrency capacity limit. This is the maximum number of tasks that can be processed concurrently.`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `(Optional) Recurring Window Schedule Day. The day of the week that the recurring window begins.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `(Optional) Schedule Hour. The hour that the recurring window begins.`,
				},
				resource.Attribute{
					Name:        "minute",
					Description: `(Optional) Schedule Minute. The minute that the recurring window begins.`,
				},
				resource.Attribute{
					Name:        "node_upg_interval",
					Description: `(Optional) Delay between node upgrades. Delay between node upgrades in seconds.`,
				},
				resource.Attribute{
					Name:        "proc_break",
					Description: `(Optional) procBreak. A period of time taken between processing of items within the concurrency cap.`,
				},
				resource.Attribute{
					Name:        "proc_cap",
					Description: `(Optional) procCap. Processing size capacity limitation specification. Indicates the limit of items to be processed within this window.`,
				},
				resource.Attribute{
					Name:        "time_cap",
					Description: `(Optional) Maximum Running Time. The processing time capacity limit. This is the maximum duration of the window.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Recurring Window.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Recurring Window.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Recurring Window.`,
				},
				resource.Attribute{
					Name:        "concur_cap",
					Description: `(Optional) Maximum Concurrent Tasks. The concurrency capacity limit. This is the maximum number of tasks that can be processed concurrently.`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `(Optional) Recurring Window Schedule Day. The day of the week that the recurring window begins.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `(Optional) Schedule Hour. The hour that the recurring window begins.`,
				},
				resource.Attribute{
					Name:        "minute",
					Description: `(Optional) Schedule Minute. The minute that the recurring window begins.`,
				},
				resource.Attribute{
					Name:        "node_upg_interval",
					Description: `(Optional) Delay between node upgrades. Delay between node upgrades in seconds.`,
				},
				resource.Attribute{
					Name:        "proc_break",
					Description: `(Optional) procBreak. A period of time taken between processing of items within the concurrency cap.`,
				},
				resource.Attribute{
					Name:        "proc_cap",
					Description: `(Optional) procCap. Processing size capacity limitation specification. Indicates the limit of items to be processed within this window.`,
				},
				resource.Attribute{
					Name:        "time_cap",
					Description: `(Optional) Maximum Running Time. The processing time capacity limit. This is the maximum duration of the window.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_rest",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Rest`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_rest_managed",
			Category:         "Data Sources",
			ShortDescription: `This data source can read one ACI object and its children.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_route_control_context",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Route Control Context`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route_control_profile_dn",
					Description: `(Required) Distinguished name of parent Route Control Profile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object Route Control Context. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Route Control Context.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Route Control Context.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Route Control Context.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action. The action required when the condition is met.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Local Order. The order of the policy context.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Route Control Context.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Route Control Context.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Route Control Context.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action. The action required when the condition is met.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Local Order. The order of the policy context.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_route_control_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Route Control Profile`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_rsa_provider",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI RSA Provider`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object RSA Provider. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the RSA Provider.`,
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
					Description: `(Optional) Port.`,
				},
				resource.Attribute{
					Name:        "auth_protocol",
					Description: `(Optional) Authentication Protocol.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Key. A password for the AAA provider database.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring.`,
				},
				resource.Attribute{
					Name:        "monitoring_password",
					Description: `(Optional) Periodic Server Monitoring Password.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Retries. null`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout in Seconds. The amount of time between authentication attempts.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the RSA Provider.`,
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
					Description: `(Optional) Port.`,
				},
				resource.Attribute{
					Name:        "auth_protocol",
					Description: `(Optional) Authentication Protocol.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Key. A password for the AAA provider database.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring.`,
				},
				resource.Attribute{
					Name:        "monitoring_password",
					Description: `(Optional) Periodic Server Monitoring Password.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Retries. null`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout in Seconds. The amount of time between authentication attempts.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_saml_certificate",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI SAML Encryption Certificate`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Key pair for SAML Encryption Certificate.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object SAML Encryption Certificate.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object SAML Encryption Certificate.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object SAML Encryption Certificate.`,
				},
				resource.Attribute{
					Name:        "regenerate",
					Description: `(Optional) Regenerate Encryption Key Pair.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional) Certificate of SAML Encryption Key.`,
				},
				resource.Attribute{
					Name:        "certificate_validty",
					Description: `(Optional) Certificate validity of SAML Encryption Certificate.`,
				},
				resource.Attribute{
					Name:        "expiry_status",
					Description: `(Optional) Expiry status of SAML Encryption Certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_saml_provider",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI SAML Provider`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object SAML Provider. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SAML Provider.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object SAML Provider.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object SAML Provider.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object SAML Provider.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `(Optional) Entity ID.`,
				},
				resource.Attribute{
					Name:        "gui_banner_message",
					Description: `(Optional) Gui Redirect Banner Message.`,
				},
				resource.Attribute{
					Name:        "https_proxy",
					Description: `(Optional) Https Proxy to reach IDP's Metadata URL.`,
				},
				resource.Attribute{
					Name:        "id_p",
					Description: `(Optional) Identity Provider.`,
				},
				resource.Attribute{
					Name:        "metadata_url",
					Description: `(Optional) Metadata Url provided by IDP.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Retries. null`,
				},
				resource.Attribute{
					Name:        "sig_alg",
					Description: `(Optional) Signature Algorithm.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout in Seconds. The amount of time between authentication attempts.`,
				},
				resource.Attribute{
					Name:        "tp",
					Description: `(Optional) Certificate Authority.`,
				},
				resource.Attribute{
					Name:        "want_assertions_encrypted",
					Description: `(Optional) Want Encrypted SAML Assertions.`,
				},
				resource.Attribute{
					Name:        "want_assertions_signed",
					Description: `(Optional) Want Assertions in SAML Response Signed.`,
				},
				resource.Attribute{
					Name:        "want_requests_signed",
					Description: `(Optional) Want SAML Auth Requests Signed.`,
				},
				resource.Attribute{
					Name:        "want_response_signed",
					Description: `(Optional) Want SAML Response Message Signed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SAML Provider.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object SAML Provider.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object SAML Provider.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object SAML Provider.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `(Optional) Entity ID.`,
				},
				resource.Attribute{
					Name:        "gui_banner_message",
					Description: `(Optional) Gui Redirect Banner Message.`,
				},
				resource.Attribute{
					Name:        "https_proxy",
					Description: `(Optional) Https Proxy to reach IDP's Metadata URL.`,
				},
				resource.Attribute{
					Name:        "id_p",
					Description: `(Optional) Identity Provider.`,
				},
				resource.Attribute{
					Name:        "metadata_url",
					Description: `(Optional) Metadata Url provided by IDP.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Retries. null`,
				},
				resource.Attribute{
					Name:        "sig_alg",
					Description: `(Optional) Signature Algorithm.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout in Seconds. The amount of time between authentication attempts.`,
				},
				resource.Attribute{
					Name:        "tp",
					Description: `(Optional) Certificate Authority.`,
				},
				resource.Attribute{
					Name:        "want_assertions_encrypted",
					Description: `(Optional) Want Encrypted SAML Assertions.`,
				},
				resource.Attribute{
					Name:        "want_assertions_signed",
					Description: `(Optional) Want Assertions in SAML Response Signed.`,
				},
				resource.Attribute{
					Name:        "want_requests_signed",
					Description: `(Optional) Want SAML Auth Requests Signed.`,
				},
				resource.Attribute{
					Name:        "want_response_signed",
					Description: `(Optional) Want SAML Response Message Signed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_saml_provider_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI SAML Provider Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object SAML Provider Group. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SAML Provider Group.`,
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
					Description: `(Optional) Description of object SAML Provider Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SAML Provider Group.`,
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
					Description: `(Optional) Description of object SAML Provider Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_service_redirect_backup_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI PBR Backup Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of the parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the PBR Backup Policy object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the PBR Backup Policy object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the PBR Backup Policy object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the PBR Backup Policy object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the PBR Backup Policy object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the PBR Backup Policy object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the PBR Backup Policy object.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_snmp_community",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI SNMP Community`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "parent_dn",
					Description: `(Required) Distinguished name of the parent object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object SNMP Community. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SNMP Community.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the SNMP Community.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the SNMP Community.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SNMP Community.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the SNMP Community.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the SNMP Community.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_span_source_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI SPAN Source Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
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
					Description: `(Required) Name of Object SPAN Source Group object. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SPAN Source-destination Group Match Label.`,
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
					Description: `(Optional) label color`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spanning_tree_interface_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Spanning Tree Interface Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spanning Tree Interface Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Spanning Tree Interface Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Spanning Tree Interface Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Spanning Tree Interface Policy.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) Interface controls.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spanning Tree Interface Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Spanning Tree Interface Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Spanning Tree Interface Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object Spanning Tree Interface Policy.`,
				},
				resource.Attribute{
					Name:        "ctrl",
					Description: `(Optional) Interface controls.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_access_port_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Spine Access Port Selector`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "spine_interface_profile_dn",
					Description: `(Required) Distinguished name of the parent Spine Interface Profile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Spine Access Port Selector.`,
				},
				resource.Attribute{
					Name:        "spine_access_port_selector_type",
					Description: `(Required) The type of Spine Access Port Selector. Allowed values are "ALL" and "range". Default is "ALL". ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Access Port Selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Spine Access Port Selector.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Spine Access Port Selector.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Access Port Selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Spine Access Port Selector.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Spine Access Port Selector.`,
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
					Description: `(Required) Name of Object spine_interface_profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Object Spine interface profile.`,
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
					Description: `(Optional) Name alias for Object Spine interface profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Object Spine interface profile.`,
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
					Description: `(Optional) Name alias for Object Spine interface profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_interface_profile_selector",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Spine Interface Profile Selector`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "spine_profile_dn",
					Description: `(Required) Distinguished name of parent Spine Profile.`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Required) tDn of the Spine Interface Profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Interface Profile selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for Spine Interface Profile selector.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Interface Profile selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for Spine Interface Profile selector.`,
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
					Description: `(Required) Name of Object Spine Port Policy Group. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Port Policy Group.`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Port Policy Group.`,
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
					Description: `(Required) Distinguished name of parent Spine Profile.`,
				},
				resource.Attribute{
					Name:        "tdn",
					Description: `(Required) tDn of the Spine Interface Profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Port selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for Spine Port selector.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Port selector.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation for Spine Port selector.`,
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
					Description: `(Required) Name of Object Spine Profile. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Profile.`,
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
					Description: `(Optional) Name alias for object Spine Profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Profile.`,
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
					Description: `(Optional) Name alias for object Spine Profile.`,
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
					Description: `(Required) Distinguished name of parent Spine Profile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Object Spine Switch Association.`,
				},
				resource.Attribute{
					Name:        "spine_switch_association_type",
					Description: `(Required) Spine association type of Object Spine Switch Association. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Switch Association.`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Switch Association.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_spine_switch_policy_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Spine Switch Policy Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object Spine Switch Policy Group. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Spine Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Spine Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Spine Switch Policy Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Spine Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object Spine Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object Spine Switch Policy Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for object Spine Switch Policy Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_static_node_mgmt_address",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Management Static Node`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "management_epg_dn",
					Description: `(Required) Distinguished name of parent Management static node object.`,
				},
				resource.Attribute{
					Name:        "t_dn",
					Description: `(Required) Target dn of Management static node object.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type for Management static node object. Allowed values are "in_band" and "out_of_band". Note := for "in_band", ` + "`" + `management_epg_dn` + "`" + ` should be of type "in_band" and for "out_of_band", ` + "`" + `management_epg_dn` + "`" + ` should be of type "out_of_band". ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of Management static node object.`,
				},
				resource.Attribute{
					Name:        "addr",
					Description: `Peer address for Management static node object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `Annotation for Management static node object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for Management static node object.`,
				},
				resource.Attribute{
					Name:        "gw",
					Description: `Gateway IP address for Management static node object.`,
				},
				resource.Attribute{
					Name:        "v6_addr",
					Description: `V6 address for Management static node object.`,
				},
				resource.Attribute{
					Name:        "v6_gw",
					Description: `V6 gw for Management static node object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of Management static node object.`,
				},
				resource.Attribute{
					Name:        "addr",
					Description: `Peer address for Management static node object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `Annotation for Management static node object.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for Management static node object.`,
				},
				resource.Attribute{
					Name:        "gw",
					Description: `Gateway IP address for Management static node object.`,
				},
				resource.Attribute{
					Name:        "v6_addr",
					Description: `V6 address for Management static node object.`,
				},
				resource.Attribute{
					Name:        "v6_gw",
					Description: `V6 gw for Management static node object.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_system",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI System`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_taboo_contract",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Taboo Contract`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tacacs_accounting",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI TACACS Accounting`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object TACACS Accounting. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the TACACS Accounting.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object TACACS Accounting.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object TACACS Accounting.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object TACACS Accounting.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the TACACS Accounting.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object TACACS Accounting.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object TACACS Accounting.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object TACACS Accounting.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tacacs_accounting_destination",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI TACACS Accounting Destination`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tacacs_accounting_dn",
					Description: `(Required) Distinguished name of parent TACACS Accounting object.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) Host or IP address of object TACACS Accounting Destination.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port of object TACACS Accounting Destination. Allowed Range: "1" - "65535". Default value: "49". ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the TACACS Destination.`,
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
					Description: `(Optional) Authentication Protocol of object TACACS Accounting Destination.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object TACACS Accounting Destination.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the TACACS Destination.`,
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
					Description: `(Optional) Authentication Protocol of object TACACS Accounting Destination.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object TACACS Accounting Destination.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tacacs_provider",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI TACACS Provider`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object TACACS Provider. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the TACACS Provider.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object TACACS Provider.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object TACACS Provider.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object TACACS Provider.`,
				},
				resource.Attribute{
					Name:        "auth_protocol",
					Description: `(Optional) TACACS Authentication Protocol. The TACACS authentication protocol.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port. The service port number for the TACACS service.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Retries. Null.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout in Seconds. The timeout for communication with the TACACS+ provider server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the TACACS Provider.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object TACACS Provider.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object TACACS Provider.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object TACACS Provider.`,
				},
				resource.Attribute{
					Name:        "auth_protocol",
					Description: `(Optional) TACACS Authentication Protocol. The TACACS authentication protocol.`,
				},
				resource.Attribute{
					Name:        "monitor_server",
					Description: `(Optional) Periodic Server Monitoring.`,
				},
				resource.Attribute{
					Name:        "monitoring_user",
					Description: `(Optional) Periodic Server Monitoring Username.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port. The service port number for the TACACS service.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) Retries. Null.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout in Seconds. The timeout for communication with the TACACS+ provider server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tacacs_provider_group",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI TACACS + Provider Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object TACACS + Provider Group.er Group. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the TACACS+ Provider Group.`,
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
					Description: `(Optional) Name alias of object TACACS + Provider Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the TACACS+ Provider Group.`,
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
					Description: `(Optional) Name alias of object TACACS + Provider Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tacacs_source",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI TACACS Source`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object TACACS Source. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "parent_dn",
					Description: `(Optional) Distinguished name of parent object of TACACS Source.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the TACACS Source.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object TACACS Source.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object TACACS Source.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object TACACS Source.`,
				},
				resource.Attribute{
					Name:        "incl",
					Description: `(Optional) Include Action. The information to include for the call home source.`,
				},
				resource.Attribute{
					Name:        "min_sev",
					Description: `(Optional) minSev.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "parent_dn",
					Description: `(Optional) Distinguished name of parent object of TACACS Source.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the TACACS Source.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object TACACS Source.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object TACACS Source.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object TACACS Source.`,
				},
				resource.Attribute{
					Name:        "incl",
					Description: `(Optional) Include Action. The information to include for the call home source.`,
				},
				resource.Attribute{
					Name:        "min_sev",
					Description: `(Optional) minSev.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tag",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Tag`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "parent_dn",
					Description: `(Required) Distinguished name of the object to which the Tag is attached to.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Key of the Tag. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Tag.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value of the Tag.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Tag.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value of the Tag.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_tenant_to_cloud_account",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Tenant to Cloud Account association`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_dn",
					Description: `(Required) Distinguished name of the parent Tenant object.`,
				},
				resource.Attribute{
					Name:        "cloud_account_dn",
					Description: `(Optional) The distinguished name of the target Cloud Account object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to Dn of the Tenant to Cloud Account association object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Tenant to Cloud Account association object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Tenant to Cloud Account association object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to Dn of the Tenant to Cloud Account association object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Tenant to Cloud Account association object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Tenant to Cloud Account association object.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_user_security_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI User Security Domain`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "local_user_dn",
					Description: `(Required) Distinguished name of parent LocalUser object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object User Domain. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the User Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object User Security Domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object User SecDomain.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object User Security Domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the User Domain.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object User Security Domain.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object User SecDomain.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object User Security Domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_user_security_domain_role",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI User Security Domain Role`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_domain_dn",
					Description: `(Required) Distinguished name of parent UserDomain object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object User Security Domain Role. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the User Security Domain Role.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object User Security Domain Role.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object User Security Domain Role.`,
				},
				resource.Attribute{
					Name:        "priv_type",
					Description: `(Optional) Privilege Type. The privilege type for a user role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object User Security Domain Role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the User Security Domain Role.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object User Security Domain Role.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object User Security Domain Role.`,
				},
				resource.Attribute{
					Name:        "priv_type",
					Description: `(Optional) Privilege Type. The privilege type for a user role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object User Security Domain Role.`,
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
					Description: `(Optional) Annotation for object vlan encapsulation for vxlan traffic.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name alias for object vlan encapsulation for vxlan traffic.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Vlan Encapsulation for Vxlan Traffic.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vlan_pool",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI VLAN Pool`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vmm_controller",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI VMM Controller`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vmm_domain_dn",
					Description: `(Required) Distinguished name of parent VMM Domain object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object VMM Controller. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VMM Controller.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object VMM Controller.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object VMM Controller.`,
				},
				resource.Attribute{
					Name:        "dvs_version",
					Description: `(Optional) Dvs Version.`,
				},
				resource.Attribute{
					Name:        "host_or_ip",
					Description: `(Optional) Hostname or IP Address.`,
				},
				resource.Attribute{
					Name:        "inventory_trig_st",
					Description: `(Optional) Triggered Inventory Sync Status.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The mode of operation.`,
				},
				resource.Attribute{
					Name:        "msft_config_err_msg",
					Description: `(Optional) Deployment Error Message of Mirosoft Plugin SCVM Controller. It captures error message encountered in SCVMM Controller plugin. This error message represents specific details for bitmask based msftConfigIssues fault.`,
				},
				resource.Attribute{
					Name:        "msft_config_issues",
					Description: `(Optional) msftConfigIssues.`,
				},
				resource.Attribute{
					Name:        "n1kv_stats_mode",
					Description: `(Optional) n1kv statistics enable.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port. Port`,
				},
				resource.Attribute{
					Name:        "root_cont_name",
					Description: `Top level container name.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The VMM control policy scope.`,
				},
				resource.Attribute{
					Name:        "seq_num",
					Description: `(Optional) An ISIS link-state packet sequence number.`,
				},
				resource.Attribute{
					Name:        "stats_mode",
					Description: `(Optional) The statistics mode.`,
				},
				resource.Attribute{
					Name:        "vxlan_depl_pref",
					Description: `(Optional) VxLAN Deployment Preference.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VMM Controller.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object VMM Controller.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object VMM Controller.`,
				},
				resource.Attribute{
					Name:        "dvs_version",
					Description: `(Optional) Dvs Version.`,
				},
				resource.Attribute{
					Name:        "host_or_ip",
					Description: `(Optional) Hostname or IP Address.`,
				},
				resource.Attribute{
					Name:        "inventory_trig_st",
					Description: `(Optional) Triggered Inventory Sync Status.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The mode of operation.`,
				},
				resource.Attribute{
					Name:        "msft_config_err_msg",
					Description: `(Optional) Deployment Error Message of Mirosoft Plugin SCVM Controller. It captures error message encountered in SCVMM Controller plugin. This error message represents specific details for bitmask based msftConfigIssues fault.`,
				},
				resource.Attribute{
					Name:        "msft_config_issues",
					Description: `(Optional) msftConfigIssues.`,
				},
				resource.Attribute{
					Name:        "n1kv_stats_mode",
					Description: `(Optional) n1kv statistics enable.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port. Port`,
				},
				resource.Attribute{
					Name:        "root_cont_name",
					Description: `Top level container name.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The VMM control policy scope.`,
				},
				resource.Attribute{
					Name:        "seq_num",
					Description: `(Optional) An ISIS link-state packet sequence number.`,
				},
				resource.Attribute{
					Name:        "stats_mode",
					Description: `(Optional) The statistics mode.`,
				},
				resource.Attribute{
					Name:        "vxlan_depl_pref",
					Description: `(Optional) VxLAN Deployment Preference.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vmm_credential",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI VMM Credential`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vmm_domain_dn",
					Description: `(Required) Distinguished name of parent VMM Domain object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of object VMM Credential. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VMM Credential.`,
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
					Description: `(Optional) Name Alias of object VMM Credential.`,
				},
				resource.Attribute{
					Name:        "pwd",
					Description: `(Optional) Password.`,
				},
				resource.Attribute{
					Name:        "usr",
					Description: `(Optional) Username.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VMM Credential.`,
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
					Description: `(Optional) Name Alias of object VMM Credential.`,
				},
				resource.Attribute{
					Name:        "pwd",
					Description: `(Optional) Password.`,
				},
				resource.Attribute{
					Name:        "usr",
					Description: `(Optional) Username.`,
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
					Description: `(Required) Distinguished name of parent Provider Profile object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of Object vmm_domain. ## Argument Reference - ` + "`" + `provider_profile_dn` + "`" + ` - (Required) Distinguished name of parent ProviderProfile object. - ` + "`" + `name` + "`" + ` - (Required) name of Object vmm_domain. ## Attribute Reference - ` + "`" + `id` + "`" + ` - Attribute id set to the Dn of the VMM Domain. - ` + "`" + `access_mode` + "`" + ` - (Optional) Access mode for object vmm domain. - ` + "`" + `annotation` + "`" + ` - (Optional) Annotation for object vmm domain. - ` + "`" + `arp_learning` + "`" + ` - (Optional) Enable/Disable arp learning for AVS Domain. - ` + "`" + `ave_time_out` + "`" + ` - (Optional) ACI Virtual Edge time-out for object vmm domain. - ` + "`" + `config_infra_pg` + "`" + ` - (Optional) Flag to enable configure infra port groups for object vmm domain. - ` + "`" + `ctrl_knob` + "`" + ` - (Optional) Type pf control knob to use. Allowed values are "none" and "epDpVerify". - ` + "`" + `delimiter` + "`" + ` - (Optional) Delimiter for object vmm domain. - ` + "`" + `enable_ave` + "`" + ` - (Optional) Flag to enable ACI Virtual Edge for object vmm domain. - ` + "`" + `enable_tag` + "`" + ` - (Optional) Flag enable tagging for object vmm domain. - ` + "`" + `enable_vm_foler` + "`" + ` - (Optional) Flag enable vm folder for object vmm domain. - ` + "`" + `encap_mode` + "`" + ` - (Optional) The layer 2 encapsulation protocol to use with the virtual switch. - ` + "`" + `enf_pref` + "`" + ` - (Optional) The switching enforcement preference. This determines whether switches can be done within the virtual switch (Local Switching) or whether all switched traffic must go through the fabric (No Local Switching). - ` + "`" + `ep_inventory_type` + "`" + ` - (Optional) Determines which end point inventory type to use for object VMM domain. - ` + "`" + `ep_ret_time` + "`" + ` - (Optional) End point retention time for object vmm domain. Allowed value range is "0" - "600". Default value is "0". - ` + "`" + `hv_avail_monitor` + "`" + ` - (Optional) Flag to enable host availability monitor for object VMM domain. - ` + "`" + `mcast_addr` + "`" + ` - (Optional) The multicast address of the VMM domain profile. - ` + "`" + `mode` + "`" + ` - (Optional) The switch to be used for the domain profile. - ` + "`" + `name_alias` + "`" + ` - (Optional) Name alias for object VMM domain. - ` + "`" + `pref_encap_mode` + "`" + ` - (Optional) The preferred encapsulation mode for object VMM domain.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vpc_domain_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI VPC Domain Policy`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object VPC Domain Policy. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VPC Domain Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object VPC Domain Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object VPC Domain Policy.`,
				},
				resource.Attribute{
					Name:        "dead_intvl",
					Description: `(Optional) The VPC peer dead interval time of object VPC Domain Policy`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object VPC Domain Policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VPC Domain Policy.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object VPC Domain Policy.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object VPC Domain Policy.`,
				},
				resource.Attribute{
					Name:        "dead_intvl",
					Description: `(Optional) The VPC peer dead interval time of object VPC Domain Policy`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object VPC Domain Policy.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vrf",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI VRF`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vrf_leak_epg_bd_subnet",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI Inter-VRF Leaked EPG/BD Subnet`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrf_dn",
					Description: `(Required) Distinguished name of the parent VRF object.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) IP of the Inter-VRF Leaked EPG/BD Subnet object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Inter-VRF Leaked EPG/BD Subnet object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Inter-VRF Leaked EPG/BD Subnet object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Inter-VRF Leaked EPG/BD Subnet object.`,
				},
				resource.Attribute{
					Name:        "allow_l3out_advertisement",
					Description: `(Optional) Visibility of the Inter-VRF Leaked EPG/BD Subnet object.`,
				},
				resource.Attribute{
					Name:        "leak_to",
					Description: `(Optional) A block representing the attributes of ` + "`" + `Tenant and VRF Destinations` + "`" + ` for the Inter-VRF Leaked Routes object. Type: Block.`,
				},
				resource.Attribute{
					Name:        "vrf_dn",
					Description: `Distinguished name of the destination VRF object, which is mapped to the ` + "`" + `Tenant and VRF Destinations` + "`" + ` object.`,
				},
				resource.Attribute{
					Name:        "allow_l3out_advertisement",
					Description: `Scope of the destination VRF object, which is mapped to the ` + "`" + `Tenant and VRF Destinations` + "`" + ` object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the Inter-VRF Leaked EPG/BD Subnet object.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of the Inter-VRF Leaked EPG/BD Subnet object.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of the Inter-VRF Leaked EPG/BD Subnet object.`,
				},
				resource.Attribute{
					Name:        "allow_l3out_advertisement",
					Description: `(Optional) Visibility of the Inter-VRF Leaked EPG/BD Subnet object.`,
				},
				resource.Attribute{
					Name:        "leak_to",
					Description: `(Optional) A block representing the attributes of ` + "`" + `Tenant and VRF Destinations` + "`" + ` for the Inter-VRF Leaked Routes object. Type: Block.`,
				},
				resource.Attribute{
					Name:        "vrf_dn",
					Description: `Distinguished name of the destination VRF object, which is mapped to the ` + "`" + `Tenant and VRF Destinations` + "`" + ` object.`,
				},
				resource.Attribute{
					Name:        "allow_l3out_advertisement",
					Description: `Scope of the destination VRF object, which is mapped to the ` + "`" + `Tenant and VRF Destinations` + "`" + ` object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vrf_snmp_context",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI VRF SNMP Context`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrf_dn",
					Description: `(Required) Distinguished name of parent VRF object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VRF SNMP Context.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of object VRF SNMP Context`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object VRF SNMP Context.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object VRF SNMP Context.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VRF SNMP Context.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of object VRF SNMP Context`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object VRF SNMP Context.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object VRF SNMP Context.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vrf_snmp_context_community",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI SNMP Community`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrf_snmp_context_dn",
					Description: `(Required) Distinguished name of parent VRF SNMP Context object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of object SNMP Community. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SNMP Community.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object SNMP Community.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object SNMP Community.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object SNMP Community.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the SNMP Community.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object SNMP Community.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object SNMP Community.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object SNMP Community.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vrf_to_bgp_address_family_context",
			Category:         "Data Sources",
			ShortDescription: `Data source for the ACI Relationship object between VRF and the BGP Address Family Context Policy`,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Required) The BGP address family. Allowed values are "ipv4-ucast", "ipv6-ucast", and default value is "ipv4-ucast". Type: String. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the BGP address family context policy relationship.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `Annotation of object BGP address family context policy relationship.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the BGP address family context policy relationship.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `Annotation of object BGP address family context policy relationship.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aci_vswitch_policy",
			Category:         "Data Sources",
			ShortDescription: `Data source for ACI VSwitch Policy Group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vmm_domain_dn",
					Description: `(Required) Distinguished name of parent VMM Domain object. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VSwitch Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object VSwitch Policy Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object VSwitch Policy Group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object VSwitch Policy Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VSwitch Policy Group.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) Annotation of object VSwitch Policy Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of object VSwitch Policy Group.`,
				},
				resource.Attribute{
					Name:        "name_alias",
					Description: `(Optional) Name Alias of object VSwitch Policy Group.`,
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
					Description: `(Required) Name of Object vxlan pool. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VXLAN Pool.`,
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
					Description: `(Optional) Name alias for object vxlan pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Attribute id set to the Dn of the VXLAN Pool.`,
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
					Description: `(Optional) Name alias for object vxlan pool.`,
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
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

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
		"aci_autonomous_system_profile":                17,
		"aci_bd_dhcp_label":                            18,
		"aci_bfd_interface_policy":                     19,
		"aci_bgp_address_family_context":               20,
		"aci_bgp_best_path_policy":                     21,
		"aci_bgp_peer_connectivity_profile":            22,
		"aci_bgp_peer_prefix":                          23,
		"aci_bgp_route_control_profile":                24,
		"aci_bgp_route_summarization":                  25,
		"aci_bgp_timers":                               26,
		"aci_bridge_domain":                            27,
		"aci_cdp_interface_policy":                     28,
		"aci_client_end_point":                         29,
		"aci_cloud_account":                            30,
		"aci_cloud_ad":                                 31,
		"aci_cloud_applicationcontainer":               32,
		"aci_cloud_availability_zone":                  33,
		"aci_cloud_aws_provider":                       34,
		"aci_cloud_cidr_pool":                          35,
		"aci_cloud_context_profile":                    36,
		"aci_cloud_credentials":                        37,
		"aci_cloud_domain_profile":                     38,
		"aci_cloud_endpoint_selector":                  39,
		"aci_cloud_endpoint_selectorfor_external_epgs": 40,
		"aci_cloud_epg":                                41,
		"aci_cloud_external_epg":                       42,
		"aci_cloud_external_network":                   43,
		"aci_cloud_ipsec_tunnel_subnet_pool":           44,
		"aci_cloud_provider_profile":                   45,
		"aci_cloud_providers_region":                   46,
		"aci_cloud_subnet":                             47,
		"aci_cloud_vpn_gateway":                        48,
		"aci_cloud_vpn_network":                        49,
		"aci_cloud_vrf_leak_routes":                    50,
		"aci_concrete_device":                          51,
		"aci_concrete_interface":                       52,
		"aci_configuration_export_policy":              53,
		"aci_configuration_import_policy":              54,
		"aci_connection":                               55,
		"aci_console_authentication":                   56,
		"aci_contract":                                 57,
		"aci_contract_subject":                         58,
		"aci_contract_subject_filter":                  59,
		"aci_contract_subject_one_way_filter":          60,
		"aci_coop_policy":                              61,
		"aci_default_authentication":                   62,
		"aci_destination_of_redirected_traffic":        63,
		"aci_dhcp_option":                              64,
		"aci_dhcp_option_policy":                       65,
		"aci_dhcp_relay_policy":                        66,
		"aci_duo_provider_group":                       67,
		"aci_encryption_key":                           68,
		"aci_end_point_retention_policy":               69,
		"aci_endpoint_controls":                        70,
		"aci_endpoint_ip_aging_profile":                71,
		"aci_endpoint_loop_protection":                 72,
		"aci_endpoint_security_group":                  73,
		"aci_endpoint_security_group_epg_selector":     74,
		"aci_endpoint_security_group_selector":         75,
		"aci_endpoint_security_group_tag_selector":     76,
		"aci_epg_to_contract":                          77,
		"aci_epg_to_contract_interface":                78,
		"aci_epg_to_domain":                            79,
		"aci_epg_to_static_path":                       80,
		"aci_epgs_using_function":                      81,
		"aci_error_disable_recovery":                   82,
		"aci_external_network_instance_profile":        83,
		"aci_fabric_if_pol":                            84,
		"aci_fabric_node":                              85,
		"aci_fabric_node_control":                      86,
		"aci_fabric_node_member":                       87,
		"aci_fabric_path_ep":                           88,
		"aci_fabric_wide_settings":                     89,
		"aci_fc_domain":                                90,
		"aci_fex_bundle_group":                         91,
		"aci_fex_profile":                              92,
		"aci_file_remote_path":                         93,
		"aci_filter":                                   94,
		"aci_filter_entry":                             95,
		"aci_firmware_download_task":                   96,
		"aci_firmware_group":                           97,
		"aci_firmware_policy":                          98,
		"aci_function_node":                            99,
		"aci_global_security":                          100,
		"aci_hsrp_group_policy":                        101,
		"aci_hsrp_interface_policy":                    102,
		"aci_imported_contract":                        103,
		"aci_interface_blacklist":                      104,
		"aci_interface_config":                         105,
		"aci_interface_fc_policy":                      106,
		"aci_ip_sla_monitoring_policy":                 107,
		"aci_isis_domain_policy":                       108,
		"aci_l2_domain":                                109,
		"aci_l2_interface_policy":                      110,
		"aci_l2_outside":                               111,
		"aci_l2out_extepg":                             112,
		"aci_l3_domain_profile":                        113,
		"aci_l3_ext_subnet":                            114,
		"aci_l3_interface_policy":                      115,
		"aci_l3_outside":                               116,
		"aci_l3out_bfd_interface_profile":              117,
		"aci_l3out_bgp_external_policy":                118,
		"aci_l3out_bgp_protocol_profile":               119,
		"aci_l3out_floating_svi":                       120,
		"aci_l3out_hsrp_interface_group":               121,
		"aci_l3out_hsrp_interface_profile":             122,
		"aci_l3out_hsrp_secondary_vip":                 123,
		"aci_l3out_loopback_interface_profile":         124,
		"aci_l3out_ospf_external_policy":               125,
		"aci_l3out_ospf_interface_profile":             126,
		"aci_l3out_path_attachment":                    127,
		"aci_l3out_path_attachment_secondary_ip":       128,
		"aci_l3out_route_tag_policy":                   129,
		"aci_l3out_static_route":                       130,
		"aci_l3out_static_route_next_hop":              131,
		"aci_l3out_vpc_member":                         132,
		"aci_l4_l7_deployed_graph_connector_vlan":      133,
		"aci_l4_l7_devices":                            134,
		"aci_l4_l7_logical_interface":                  135,
		"aci_l4_l7_redirect_health_group":              136,
		"aci_l4_l7_service_graph_template":             137,
		"aci_lacp_member_policy":                       138,
		"aci_lacp_policy":                              139,
		"aci_ldap_group_map":                           140,
		"aci_ldap_group_map_rule":                      141,
		"aci_ldap_group_map_rule_to_group_map":         142,
		"aci_ldap_provider":                            143,
		"aci_leaf_access_bundle_policy_group":          144,
		"aci_leaf_access_bundle_policy_sub_group":      145,
		"aci_leaf_access_port_policy_group":            146,
		"aci_leaf_breakout_port_group":                 147,
		"aci_leaf_interface_profile":                   148,
		"aci_leaf_profile":                             149,
		"aci_leaf_selector":                            150,
		"aci_lldp_interface_policy":                    151,
		"aci_local_user":                               152,
		"aci_logical_device_context":                   153,
		"aci_logical_interface_context":                154,
		"aci_logical_interface_profile":                155,
		"aci_logical_node_profile":                     156,
		"aci_logical_node_to_fabric_node":              157,
		"aci_login_domain":                             158,
		"aci_login_domain_provider":                    159,
		"aci_maintenance_group_node":                   160,
		"aci_maintenance_policy":                       161,
		"aci_managed_node_connectivity_group":          162,
		"aci_match_community_terms":                    163,
		"aci_match_regex_community_terms":              164,
		"aci_match_route_destination_rule":             165,
		"aci_match_rule":                               166,
		"aci_mcp_instance_policy":                      167,
		"aci_mgmt_preference":                          168,
		"aci_mgmt_zone":                                169,
		"aci_miscabling_protocol_interface_policy":     170,
		"aci_monitoring_policy":                        171,
		"aci_multicast_pool":                           172,
		"aci_multicast_pool_block":                     173,
		"aci_node_block":                               174,
		"aci_node_block_firmware":                      175,
		"aci_node_mgmt_epg":                            176,
		"aci_ospf_interface_policy":                    177,
		"aci_ospf_route_summarization":                 178,
		"aci_ospf_timers":                              179,
		"aci_pbr_l1_l2_destination":                    180,
		"aci_physical_domain":                          181,
		"aci_pod_maintenance_group":                    182,
		"aci_port_security_policy":                     183,
		"aci_port_tracking":                            184,
		"aci_qos_instance_policy":                      185,
		"aci_radius_provider":                          186,
		"aci_radius_provider_group":                    187,
		"aci_ranges":                                   188,
		"aci_recurring_window":                         189,
		"aci_rest":                                     190,
		"aci_rest_managed":                             191,
		"aci_route_control_context":                    192,
		"aci_route_control_profile":                    193,
		"aci_rsa_provider":                             194,
		"aci_saml_certificate":                         195,
		"aci_saml_provider":                            196,
		"aci_saml_provider_group":                      197,
		"aci_service_redirect_backup_policy":           198,
		"aci_service_redirect_policy":                  199,
		"aci_snmp_community":                           200,
		"aci_span_destination_group":                   201,
		"aci_span_source_group":                        202,
		"aci_span_sourcedestination_group_match_label": 203,
		"aci_spanning_tree_interface_policy":           204,
		"aci_spine_access_port_selector":               205,
		"aci_spine_interface_profile":                  206,
		"aci_spine_interface_profile_selector":         207,
		"aci_spine_port_policy_group":                  208,
		"aci_spine_port_selector":                      209,
		"aci_spine_profile":                            210,
		"aci_spine_switch_association":                 211,
		"aci_spine_switch_policy_group":                212,
		"aci_static_node_mgmt_address":                 213,
		"aci_subnet":                                   214,
		"aci_system":                                   215,
		"aci_taboo_contract":                           216,
		"aci_tacacs_accounting":                        217,
		"aci_tacacs_accounting_destination":            218,
		"aci_tacacs_provider":                          219,
		"aci_tacacs_provider_group":                    220,
		"aci_tacacs_source":                            221,
		"aci_tag":                                      222,
		"aci_tenant":                                   223,
		"aci_tenant_to_cloud_account":                  224,
		"aci_trigger_scheduler":                        225,
		"aci_user_security_domain":                     226,
		"aci_user_security_domain_role":                227,
		"aci_vlan_encapsulationfor_vxlan_traffic":      228,
		"aci_vlan_pool":                                229,
		"aci_vmm_controller":                           230,
		"aci_vmm_credential":                           231,
		"aci_vmm_domain":                               232,
		"aci_vpc_domain_policy":                        233,
		"aci_vpc_explicit_protection_group":            234,
		"aci_vrf":                                      235,
		"aci_vrf_leak_epg_bd_subnet":                   236,
		"aci_vrf_snmp_context":                         237,
		"aci_vrf_snmp_context_community":               238,
		"aci_vrf_to_bgp_address_family_context":        239,
		"aci_vsan_pool":                                240,
		"aci_vswitch_policy":                           241,
		"aci_vxlan_pool":                               242,
		"aci_x509_certificate":                         243,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
