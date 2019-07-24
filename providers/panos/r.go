package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "panos_address_group",
			Category:         "Firewall Resources",
			ShortDescription: `Manages address groups.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"address",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The address group's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the address group into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_addresses",
					Description: `(Optional) The address objects to include in this statically defined address group.`,
				},
				resource.Attribute{
					Name:        "dynamic_match",
					Description: `(Optional) The IP tags to include in this DAG.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The address group's description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_address_object",
			Category:         "Firewall Resources",
			ShortDescription: `Manages address objects.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"address",
				"object",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The address object's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the address object into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of address object. This can be ` + "`" + `ip-netmask` + "`" + ` (default), ` + "`" + `ip-range` + "`" + `, or ` + "`" + `fqdn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The address object's value. This can take various forms depending on what type of address object this is, but can be something like ` + "`" + `192.168.80.150` + "`" + ` or ` + "`" + `192.168.80.0/24` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The address object's description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_administrative_tag",
			Category:         "Firewall Resources",
			ShortDescription: `Manages administrative tags.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"administrative",
				"tag",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The administrative tag's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the administrative tag into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) The tag's color. This should be either an empty string (no color) or a string such as ` + "`" + `color1` + "`" + ` or ` + "`" + `color15` + "`" + `. Note that for maximum portability, you should limit color usage to ` + "`" + `color16` + "`" + `, which was available in PAN-OS 6.1. PAN-OS 8.1's colors go up to ` + "`" + `color42` + "`" + `. The value ` + "`" + `color18` + "`" + ` is reserved internally by PAN-OS and thus not available for use.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The administrative tag's description.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bfd_profile",
			Category:         "Firewall Resources",
			ShortDescription: `Manages BFD profiles.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bfd",
				"profile",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The BBFD profile's name.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) BFD operation mode. Valid values are ` + "`" + `active` + "`" + ` (default) or ` + "`" + `passive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minimum_tx_interval",
					Description: `(Optional, int) Desired minimum TX interval in ms. Default is ` + "`" + `1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minimum_rx_interval",
					Description: `(Optional, int) Required minimum RX interval in ms. Default is ` + "`" + `1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "detection_multiplier",
					Description: `(Optional, int) Multiplier sent to remote system. Default is ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hold_time",
					Description: `(Optional, int) Delay transmission and reception of control packets in ms.`,
				},
				resource.Attribute{
					Name:        "minimum_rx_ttl",
					Description: `(Optional, int) Minimum accepted ttl on received BFD packet.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a virtual router's BGP configuration.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bgp",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP configuration to.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable BGP or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Optional) Router ID of this BGP instance.`,
				},
				resource.Attribute{
					Name:        "as_number",
					Description: `(Optional) Local AS number.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `(Optional, PAN-OS 7.1+) BFD configuration.`,
				},
				resource.Attribute{
					Name:        "reject_default_route",
					Description: `(Optional, bool) Do not learn default route from BGP (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "install_route",
					Description: `(Optional, bool) Populate BGP learned route to global route table.`,
				},
				resource.Attribute{
					Name:        "aggregate_med",
					Description: `(Optional, bool) Aggregate route only if they have same MED attributes (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "default_local_preference",
					Description: `(Optional) Default local preference (default: ` + "`" + `"100"` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_format",
					Description: `(Optional) AS format. Valid values are ` + "`" + `"2-byte"` + "`" + ` (default) or ` + "`" + `"4-byte"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "always_compare_med",
					Description: `(Optional, bool) Always compare MEDs.`,
				},
				resource.Attribute{
					Name:        "deterministic_med_comparison",
					Description: `(Optional, bool) Deterministic MED comparison (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ecmp_multi_as",
					Description: `(Optional, bool) Support multiple AS in ECMP.`,
				},
				resource.Attribute{
					Name:        "enforce_first_as",
					Description: `(Optional, bool) Enforce First AS for EBGP (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_graceful_restart",
					Description: `(Optional, bool) Enable graceful restart (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "stale_route_time",
					Description: `(Optional, int) Time to remove stale routes after peer restart, in seconds (default: ` + "`" + `120` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "local_restart_time",
					Description: `(Optional, int) Local restart time to advertise to peer, in seconds (default: ` + "`" + `120` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "max_peer_restart_time",
					Description: `(Optional, int) Maximum of peer restart time accepted, in seconds (default: ` + "`" + `120` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "reflector_cluster_id",
					Description: `(Optional) Route reflector cluster ID.`,
				},
				resource.Attribute{
					Name:        "confederation_member_as",
					Description: `(Optional) Confederation requires member-AS number.`,
				},
				resource.Attribute{
					Name:        "allow_redistribute_default_route",
					Description: `(Optional, bool) Allow redistribute default route to BGP.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_aggregate",
			Category:         "Firewall Resources",
			ShortDescription: `Manages BGP address aggregation rules.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bgp",
				"aggregate",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to put the rule into.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The rule name.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Aggregating address prefix.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this rule (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "as_set",
					Description: `(Optional, bool) Generate AS-set attribute.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `(Optional, bool) Summarize route.`,
				},
				resource.Attribute{
					Name:        "local_preference",
					Description: `(Optional) New local preference value.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) New MED value.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional, int) New weight value.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop address.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) New route origin. Valid values are ` + "`" + `incomplete` + "`" + ` (default), ` + "`" + `igp` + "`" + `, or ` + "`" + `egp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_limit",
					Description: `(Optional, int) Add AS path limit attribute if it does not exist.`,
				},
				resource.Attribute{
					Name:        "as_path_type",
					Description: `(Optional) AS path update options. Valid values are ` + "`" + `none` + "`" + ` (default) or ` + "`" + `prepend` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_value",
					Description: `(Optional) For ` + "`" + `as_path_type` + "`" + ` of ` + "`" + `prepend` + "`" + `, the value to prepend.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community update options. Valid values are ` + "`" + `none` + "`" + ` (default), ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_value",
					Description: `(Optional) If ` + "`" + `community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. For the ` + "`" + `append` + "`" + ` and ` + "`" + `overwrite` + "`" + ` types specifically, valid values are ` + "`" + `no-export` + "`" + `, ` + "`" + `no-advertise` + "`" + `, ` + "`" + `local-as` + "`" + `, or ` + "`" + `nopeer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_type",
					Description: `(Optional) Extended community update options. Valid values are ` + "`" + `none` + "`" + ` (default), ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_vaule",
					Description: `(Optional) If ` + "`" + `extended_community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_aggregate_advertise_filter",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a route advertise filter for a BGP address aggregation rule.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bgp",
				"aggregate",
				"advertise",
				"filter",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_aggregate",
					Description: `(Required) The BGP address aggregation rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below).`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry. Each ` + "`" + `address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_aggregate_suppress_filter",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a route suppression filter for a BGP address aggregation rule.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bgp",
				"aggregate",
				"suppress",
				"filter",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_aggregate",
					Description: `(Required) The BGP address aggregation rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below).`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry. Each ` + "`" + `address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_auth_profile",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a BGP auth profile.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bgp",
				"auth",
				"profile",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP auth profile to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Shared secret for the TCP MD5 authentication.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_conditional_adv",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a BGP conditional advertisement.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bgp",
				"conditional",
				"adv",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP conditional advertisement to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "used_by",
					Description: `(Optional) List of BGP peer groups that use this rule.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_conditional_adv_advertise_filter",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a advertise filter for a BGP conditional advertisement.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bgp",
				"conditional",
				"adv",
				"advertise",
				"filter",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_conditional_adv",
					Description: `(Required) The BGP conditional advertisement to add this filter to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "address_prefixes",
					Description: `(Optional) List of matching address prefixes.`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_conditional_adv_non_exist_filter",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a non-exist filter for a BGP conditional advertisement.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bgp",
				"conditional",
				"adv",
				"non",
				"exist",
				"filter",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_conditional_adv",
					Description: `(Required) The BGP conditional advertisement to add this filter to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "address_prefixes",
					Description: `(Optional) List of matching address prefixes.`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_dampening_profile",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a BGP dampening profile.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bgp",
				"dampening",
				"profile",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP dampening profile to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "cutoff",
					Description: `(Optional, float) Cutoff threshold value (default: ` + "`" + `1.25` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "reuse",
					Description: `(Optional, float) Reuse threshold value (default: ` + "`" + `0.5` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "max_hold_time",
					Description: `(Optional, int) Maximum hold-down time, in seconds (default: ` + "`" + `900` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "decay_half_life_reachable",
					Description: `(Optional, int) Decay half-life while reachable, in seconds (default: ` + "`" + `300` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "decay_half_life_unreachable",
					Description: `(Optional, int) Decay half-life while unreachable, in seconds (default: ` + "`" + `900` + "`" + `).`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_export_rule_group",
			Category:         "Firewall Resources",
			ShortDescription: `Manages BGP export rule groups.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bgp",
				"export",
				"rule",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to put the rule into.`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `(Optional) A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `(Optional) Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The export rule definition (see below). The export rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The security rule name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this export rule (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "used_by",
					Description: `(Optional) List of auth profiles.`,
				},
				resource.Attribute{
					Name:        "match_as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "match_community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "match_extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "match_med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "match_route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "match_address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below). below for the params for this section.`,
				},
				resource.Attribute{
					Name:        "match_next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "match_from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Rule action. Valid values are ` + "`" + `allow` + "`" + ` (default) or ` + "`" + `deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dampening",
					Description: `(Optional) Route flap dampening profile.`,
				},
				resource.Attribute{
					Name:        "local_preference",
					Description: `(Optional) New local preference value.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) New MED value.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional, int) New weight value.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop address.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) New route origin. Valid values are ` + "`" + `igp` + "`" + `, ` + "`" + `egp` + "`" + `, or ` + "`" + `incomplete` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_limit",
					Description: `(Optional, int) Add AS path limit attribute if it does not exist.`,
				},
				resource.Attribute{
					Name:        "as_path_type",
					Description: `(Optional) AS path update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove` + "`" + `, ` + "`" + `prepend` + "`" + `, or ` + "`" + `remove-and-prepend` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_value",
					Description: `(Optional) If ` + "`" + `as_path_type` + "`" + ` is ` + "`" + `prepend` + "`" + ` or ` + "`" + `remove-and-prepend` + "`" + `, the value to prepend.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_value",
					Description: `(Optional) If ` + "`" + `community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. For the ` + "`" + `append` + "`" + ` and ` + "`" + `overwrite` + "`" + ` types specifically, valid values for ` + "`" + `community_value` + "`" + ` are ` + "`" + `no-export` + "`" + `, ` + "`" + `no-advertise` + "`" + `, ` + "`" + `local-as` + "`" + `, or ` + "`" + `nopeer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_type",
					Description: `(Optional) Extended community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_vaule",
					Description: `(Optional) If ` + "`" + `extended_community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. Each ` + "`" + `match_address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_import_rule_group",
			Category:         "Firewall Resources",
			ShortDescription: `Manages BGP import rule groups.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bgp",
				"import",
				"rule",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to put the rule into.`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `(Optional) A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `(Optional) Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The import rule definition (see below). The import rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The security rule name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this import rule (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "used_by",
					Description: `(Optional) List of auth profiles.`,
				},
				resource.Attribute{
					Name:        "match_as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "match_community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "match_extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "match_med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "match_route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "match_address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below). below for the params for this section.`,
				},
				resource.Attribute{
					Name:        "match_next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "match_from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Rule action. Valid values are ` + "`" + `allow` + "`" + ` (default) or ` + "`" + `deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dampening",
					Description: `(Optional) Route flap dampening profile.`,
				},
				resource.Attribute{
					Name:        "local_preference",
					Description: `(Optional) New local preference value.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) New MED value.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional, int) New weight value.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop address.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) New route origin. Valid values are ` + "`" + `igp` + "`" + `, ` + "`" + `egp` + "`" + `, or ` + "`" + `incomplete` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_limit",
					Description: `(Optional, int) Add AS path limit attribute if it does not exist.`,
				},
				resource.Attribute{
					Name:        "as_path_type",
					Description: `(Optional) AS path update options. Valid values are ` + "`" + `none` + "`" + ` or ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_value",
					Description: `(Optional) If ` + "`" + `community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. For the ` + "`" + `append` + "`" + ` and ` + "`" + `overwrite` + "`" + ` types specifically, valid values for ` + "`" + `community_value` + "`" + ` are ` + "`" + `no-export` + "`" + `, ` + "`" + `no-advertise` + "`" + `, ` + "`" + `local-as` + "`" + `, or ` + "`" + `nopeer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_type",
					Description: `(Optional) Extended community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_vaule",
					Description: `(Optional) If ` + "`" + `extended_community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. Each ` + "`" + `match_address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_peer",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a BGP peer.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bgp",
				"peer",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP peer to.`,
				},
				resource.Attribute{
					Name:        "bgp_peer_group",
					Description: `(Required) The BGP peer group to put this peer into.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "peer_as",
					Description: `(Optional) Peer AS number.`,
				},
				resource.Attribute{
					Name:        "local_address_interface",
					Description: `(Required) Interface to accept BGP session.`,
				},
				resource.Attribute{
					Name:        "local_address_ip",
					Description: `(Optional) Specify exact IP address if interface has multiple addresses.`,
				},
				resource.Attribute{
					Name:        "peer_address_ip",
					Description: `(Required) Peer IP address configuration.`,
				},
				resource.Attribute{
					Name:        "reflector_client",
					Description: `(Optional) This peer is reflector client. Valid values are ` + "`" + `non-client` + "`" + `, ` + "`" + `client` + "`" + `, or ` + "`" + `meshed-client` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "peering_type",
					Description: `(Optional) Peering type that affects NOPEER community value handling. Valid values are ` + "`" + `unspecified` + "`" + ` (default) or ` + "`" + `bilateral` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_prefixes",
					Description: `(Optional) Maximum of prefixes to receive from the peer. This can be a number such as ` + "`" + `"5000"` + "`" + ` (default) or ` + "`" + `unlimited` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_profile",
					Description: `(Optional) Auth profile.`,
				},
				resource.Attribute{
					Name:        "keep_alive_interval",
					Description: `(Optional, int) Keep alive interval, in seconds (default: ` + "`" + `30` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "multi_hop",
					Description: `(Optional, int) IP TTL value used for sending BGP packet.`,
				},
				resource.Attribute{
					Name:        "open_delay_time",
					Description: `(Optional, int) Open delay time, in seconds.`,
				},
				resource.Attribute{
					Name:        "hold_time",
					Description: `(Optional, int) Hold time, in seconds.`,
				},
				resource.Attribute{
					Name:        "idle_hold_time",
					Description: `(Optional, int) Idle hold time, in seconds.`,
				},
				resource.Attribute{
					Name:        "allow_incoming_connections",
					Description: `(Optional, bool) Allow incoming connections (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "incoming_connections_remote_port",
					Description: `(Optional, int) Restrict remote port for incoming BGP connections.`,
				},
				resource.Attribute{
					Name:        "allow_outgoing_connections",
					Description: `(Optional, bool) Allow outgoing connections (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "outgoing_connections_local_port",
					Description: `(Optional, int) Use specific local port for outgoing BGP connections.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `(Optional, PAN-OS 7.1+) BFD profile. This can be a specific BFD profile name, ` + "`" + `None` + "`" + ` (disables BFD), or ` + "`" + `Inherit-vr-global-setting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_mp_bgp",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable MP BGP.`,
				},
				resource.Attribute{
					Name:        "address_family_type",
					Description: `(Optional, PAN-OS 8.0+) Set the AFI for this peer. Valid values are ` + "`" + `ipv4` + "`" + ` or ` + "`" + `ipv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subsequent_address_family_unicast",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable unicast subsequent address family for this peer.`,
				},
				resource.Attribute{
					Name:        "subsequent_address_family_multicast",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable multicast subsequent address family for this peer.`,
				},
				resource.Attribute{
					Name:        "enable_sender_side_loop_detection",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable sender side loop detection.`,
				},
				resource.Attribute{
					Name:        "min_route_advertisement_interval",
					Description: `(Optional, int, PAN-OS 8.1+) Minimum route advertisement interval, in seconds.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_peer_group",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a BGP peer group.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bgp",
				"peer",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP peer group to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "aggregated_confed_as_path",
					Description: `(Optional, bool) The peers understand aggregated confederation AS path (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "soft_reset_with_stored_info",
					Description: `(Optional, bool) Soft reset with stored info.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Peer group type. Valid options are ` + "`" + `ebgp` + "`" + ` (default), ` + "`" + `ebgp-confed` + "`" + `, ` + "`" + `ibgp` + "`" + `, or ` + "`" + `ibgp-confed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "export_next_hop",
					Description: `(Optional) Export next hop. Valid values are ` + "`" + `original` + "`" + `, ` + "`" + `use-self` + "`" + `, or ` + "`" + `resolve` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "import_next_hop",
					Description: `(Optional) Import next hop. Valid values are ` + "`" + `original` + "`" + `, ` + "`" + `use-peer` + "`" + `, or the empty string.`,
				},
				resource.Attribute{
					Name:        "remove_private_as",
					Description: `(Optional, bool) Remove private AS when exporting route. Only available for ` + "`" + `type=ebgp` + "`" + `.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_redist_rule",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a BGP redistribution rule.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"bgp",
				"redist",
				"rule",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP redist rule to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A subnet or a redistribution profile.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this rule or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `(Optional) The address family. Valid values are ` + "`" + `ipv4` + "`" + ` (default) or ` + "`" + `ipv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional, int) Metric value.`,
				},
				resource.Attribute{
					Name:        "set_origin",
					Description: `(Optional) Add the origin path attribute. Valid values are ` + "`" + `incomplete` + "`" + ` (default), ` + "`" + `igp` + "`" + `, or ` + "`" + `egp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "set_med",
					Description: `(Optional) Add the MULTI_EXIT_DISC path attribute.`,
				},
				resource.Attribute{
					Name:        "set_local_preference",
					Description: `(Optional) Add the LOCAL_PREF path attribute.`,
				},
				resource.Attribute{
					Name:        "set_as_path_limit",
					Description: `(Optional, int) Add the AS_PATHLIMIT path attribute.`,
				},
				resource.Attribute{
					Name:        "set_communities",
					Description: `(Optional) List of COMMUNITY path attributes to add.`,
				},
				resource.Attribute{
					Name:        "set_extended_communities",
					Description: `(Optional) List of EXTENDED COMMUNITY path attributes to add.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_dag_tags",
			Category:         "Firewall Resources",
			ShortDescription: `Manages Dynamic address group tags.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"dag",
				"tags",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the DAG tags in (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "register",
					Description: `(Required) A set that includes ` + "`" + `ip` + "`" + `, the IP address to be tagged and ` + "`" + `tags` + "`" + `, a list of tags to associate with the given IP.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_edl",
			Category:         "Firewall Resources",
			ShortDescription: `Manages external dynamic lists.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"edl",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the object into (default: ` + "`" + `vsys1` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of EDL. This can be ` + "`" + `ip` + "`" + ` (the default; and the only valid value for PAN-OS 6.1 - 7.0), ` + "`" + `domain` + "`" + `, ` + "`" + `url` + "`" + `, or ` + "`" + `predefined` + "`" + ` (PAN-OS 8.0+)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The object's description.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) The EDL source URL`,
				},
				resource.Attribute{
					Name:        "certificate_profile",
					Description: `(Optional) Profile for authenticating client certificates`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) EDL username`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) EDL password`,
				},
				resource.Attribute{
					Name:        "repeat",
					Description: `(Optional) How often to retrieve the EDL. This can be ` + "`" + `hourly` + "`" + ` (the default), ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `, ` + "`" + `monthly` + "`" + `, or ` + "`" + `every five minutes` + "`" + ` (valid for PAN-OS 7.1+)`,
				},
				resource.Attribute{
					Name:        "repeat_at",
					Description: `(Optional) The time at which to retrieve the EDL. Please refer to the section above for how to set this value properly.`,
				},
				resource.Attribute{
					Name:        "repeat_day_of_week",
					Description: `(Optional) If ` + "`" + `repeat` + "`" + ` is ` + "`" + `weekly` + "`" + `, then this should be set to the desired day of the week. Valid values are ` + "`" + `sunday` + "`" + `, ` + "`" + `monday` + "`" + `, ` + "`" + `tuesday` + "`" + `, ` + "`" + `wednesday` + "`" + `, ` + "`" + `thursday` + "`" + `, ` + "`" + `friday` + "`" + `, ` + "`" + `saturday` + "`" + `, and ` + "`" + `sunday` + "`" + ``,
				},
				resource.Attribute{
					Name:        "repeat_day_of_month",
					Description: `(Optional, int) If ` + "`" + `repeat` + "`" + ` is ` + "`" + `monthly` + "`" + `, then this should be set to the desired day of the month.`,
				},
				resource.Attribute{
					Name:        "exceptions",
					Description: `(Optional, list) Provide a list of exception entries.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ethernet_interface",
			Category:         "Firewall Resources",
			ShortDescription: `Manages ethernet interfaces.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"ethernet",
				"interface",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The ethernet interface's name. This should be something like ` + "`" + `ethernet1/X` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Required) The vsys that will use this interface. This should be something like ` + "`" + `vsys1` + "`" + ` or ` + "`" + `vsys3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The interface mode. This can be any of the following values: ` + "`" + `layer3` + "`" + `, ` + "`" + `layer2` + "`" + `, ` + "`" + `virtual-wire` + "`" + `, ` + "`" + `tap` + "`" + `, ` + "`" + `ha` + "`" + `, ` + "`" + `decrypt-mirror` + "`" + `, or ` + "`" + `aggregate-group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable DHCP on this interface.`,
				},
				resource.Attribute{
					Name:        "create_dhcp_default_route",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to create a DHCP default route.`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route_metric",
					Description: `(Optional) The metric for the DHCP default route.`,
				},
				resource.Attribute{
					Name:        "ipv6_enabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable IPv6.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "lldp_enabled",
					Description: `(Optional) Enable LLDP (default: false).`,
				},
				resource.Attribute{
					Name:        "lldp_profile",
					Description: `(Optional) LLDP profile.`,
				},
				resource.Attribute{
					Name:        "link_speed",
					Description: `(Optional) Link speed. This can be any of the following: ` + "`" + `10` + "`" + `, ` + "`" + `100` + "`" + `, ` + "`" + `1000` + "`" + `, or ` + "`" + `auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "link_duplex",
					Description: `(Optional) Link duplex setting. This can be ` + "`" + `full` + "`" + `, ` + "`" + `half` + "`" + `, or ` + "`" + `auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "link_state",
					Description: `(Optional) The link state. This can be ` + "`" + `up` + "`" + `, ` + "`" + `down` + "`" + `, or ` + "`" + `auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aggregate_group",
					Description: `(Optional) The aggregate group (applicable for physical firewalls only).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv6 MSS adjust value.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_general_settings",
			Category:         "Firewall Resources",
			ShortDescription: `Manages general device settings.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"general",
				"settings",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "hostname",
					Description: `Firewall hostname.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `The timezone (e.g. - ` + "`" + `US/Pacific` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain.`,
				},
				resource.Attribute{
					Name:        "update_server",
					Description: `The update server (Default: ` + "`" + `updates.paloaltonetworks.com` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "verify_update_server",
					Description: `Verify update server identity (Default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "proxy_server",
					Description: `(1.5+) Specify a proxy server.`,
				},
				resource.Attribute{
					Name:        "proxy_port",
					Description: `(int, 1.5+) Proxy's port number.`,
				},
				resource.Attribute{
					Name:        "proxy_username",
					Description: `(1.5+) Proxy's username.`,
				},
				resource.Attribute{
					Name:        "proxy_password",
					Description: `(1.5+) Proxy's password.`,
				},
				resource.Attribute{
					Name:        "dns_primary",
					Description: `Primary DNS server.`,
				},
				resource.Attribute{
					Name:        "dns_secondary",
					Description: `Secondary DNS server.`,
				},
				resource.Attribute{
					Name:        "ntp_primary_address",
					Description: `Primary NTP server.`,
				},
				resource.Attribute{
					Name:        "ntp_primary_auth_type",
					Description: `Primary NTP auth type. This can be ` + "`" + `none` + "`" + `, ` + "`" + `autokey` + "`" + `, or ` + "`" + `symmetric-key` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ntp_primary_key_id",
					Description: `Primary NTP ` + "`" + `symmetric-key` + "`" + ` key ID.`,
				},
				resource.Attribute{
					Name:        "ntp_primary_algorithm",
					Description: `Primary NTP ` + "`" + `symmetric-key` + "`" + ` algorithm. This can be ` + "`" + `sha1` + "`" + ` or ` + "`" + `md5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ntp_primary_auth_key",
					Description: `Primary NTP ` + "`" + `symmetric-key` + "`" + ` auth key. This is the SHA1 hash if the algorithm is ` + "`" + `sha1` + "`" + `, or the md5sum if the algorithm is ` + "`" + `md5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ntp_secondary_address",
					Description: `Secondary NTP server.`,
				},
				resource.Attribute{
					Name:        "ntp_secondary_auth_type",
					Description: `Secondary NTP auth type. This can be ` + "`" + `none` + "`" + `, ` + "`" + `autokey` + "`" + `, or ` + "`" + `symmetric-key` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ntp_secondary_key_id",
					Description: `Secondary NTP ` + "`" + `symmetric-key` + "`" + ` key ID.`,
				},
				resource.Attribute{
					Name:        "ntp_secondary_algorithm",
					Description: `Secondary NTP ` + "`" + `symmetric-key` + "`" + ` algorithm. This can be ` + "`" + `sha1` + "`" + ` or ` + "`" + `md5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ntp_secondary_auth_key",
					Description: `Secondary NTP ` + "`" + `symmetric-key` + "`" + ` auth key. This is the SHA1 hash if the algorithm is ` + "`" + `sha1` + "`" + `, or the md5sum if the algorithm is ` + "`" + `md5` + "`" + `.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ike_crypto_profile",
			Category:         "Firewall Resources",
			ShortDescription: `Manages IKE crypto profiles.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"ike",
				"crypto",
				"profile",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "dh_groups",
					Description: `(Required, list) List of DH Group entries. Values should have a prefix if ` + "`" + `group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "authentications",
					Description: `(Required, list) List of authentication types. This c`,
				},
				resource.Attribute{
					Name:        "encryptions",
					Description: `(Required, list) List of encryption types. Valid values are ` + "`" + `des` + "`" + `, ` + "`" + `3des` + "`" + `, ` + "`" + `aes-128-cbc` + "`" + `, ` + "`" + `aes-192-cbc` + "`" + `, and ` + "`" + `aes-256-cbc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifetime_type",
					Description: `(Optional) The lifetime type. Valid values are ` + "`" + `seconds` + "`" + `, ` + "`" + `minutes` + "`" + `, ` + "`" + `hours` + "`" + ` (the default), and ` + "`" + `days` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifetime_value",
					Description: `(Optional, int) The lifetime value.`,
				},
				resource.Attribute{
					Name:        "authentication_multiple",
					Description: `(Optional, PAN-OS 7.0+, int) IKEv2 SA reauthentication interval equals authetication-multiple`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ike_gateway",
			Category:         "Firewall Resources",
			ShortDescription: `Manages IKE gateways.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"ike",
				"gateway",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional, PAN-OS 7.0+) The IKE gateway version. Valid values are ` + "`" + `ikev1` + "`" + `, (the default), ` + "`" + `ikev2` + "`" + `, or ` + "`" + `ikev2-preferred` + "`" + `. For PAN-OS 6.1, only ` + "`" + `ikev1` + "`" + ` is acceptable.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `(Optional, PAN-OS 7.0+, bool) Enable IPv6 or not.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to disable.`,
				},
				resource.Attribute{
					Name:        "peer_ip_type",
					Description: `(Optional) The peer IP type. Valid values are ` + "`" + `ip` + "`" + `, ` + "`" + `dynamic` + "`" + `, and ` + "`" + `fqdn` + "`" + ` (PANOS 8.1+).`,
				},
				resource.Attribute{
					Name:        "peer_ip_value",
					Description: `(Optional) The peer IP value.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface.`,
				},
				resource.Attribute{
					Name:        "local_ip_address_type",
					Description: `(Optional) The local IP address type. Valid values for this are ` + "`" + `ip` + "`" + `, ` + "`" + `floating-ip` + "`" + `, or an empty string (the default) which is ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_ip_address_value",
					Description: `(Optional) The IP address if ` + "`" + `local_ip_address_type` + "`" + ` is set to ` + "`" + `ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) The auth type. Valid values are ` + "`" + `pre-shared-key` + "`" + ` (the default), or ` + "`" + `certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(Optional) The pre-shared key value.`,
				},
				resource.Attribute{
					Name:        "local_id_type",
					Description: `(Optional) The local ID type. Valid values are ` + "`" + `ipaddr` + "`" + `, ` + "`" + `fqdn` + "`" + `, ` + "`" + `ufqdn` + "`" + `, ` + "`" + `keyid` + "`" + `, or ` + "`" + `dn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_id_value",
					Description: `(Optional) The local ID value.`,
				},
				resource.Attribute{
					Name:        "peer_id_type",
					Description: `(Optional) The peer ID type. Valid values are ` + "`" + `ipaddr` + "`" + `, ` + "`" + `fqdn` + "`" + `, ` + "`" + `ufqdn` + "`" + `, ` + "`" + `keyid` + "`" + `, or ` + "`" + `dn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "peer_id_value",
					Description: `(Optional) The peer ID value.`,
				},
				resource.Attribute{
					Name:        "peer_id_check",
					Description: `(Optional) Enable peer ID wildcard match for certificate authentication. Valid values are ` + "`" + `exact` + "`" + ` or ` + "`" + `wildcard` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_cert",
					Description: `(Optional) The local certificate name.`,
				},
				resource.Attribute{
					Name:        "cert_enable_hash_and_url",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to use hash-and-url for local certificate.`,
				},
				resource.Attribute{
					Name:        "cert_base_url",
					Description: `(Optional) The host and directory part of URL for local certificates.`,
				},
				resource.Attribute{
					Name:        "cert_use_management_as_source",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to use management interface IP as source to retrieve http certificates`,
				},
				resource.Attribute{
					Name:        "cert_permit_payload_mismatch",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to permit peer identification and certificate payload identification mismatch.`,
				},
				resource.Attribute{
					Name:        "cert_profile",
					Description: `(Optional) Profile for certificate valdiation during IKE negotiation.`,
				},
				resource.Attribute{
					Name:        "cert_enable_strict_validation",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable strict validation of peer's extended key use.`,
				},
				resource.Attribute{
					Name:        "enable_passive_mode",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable passive mode (responder only).`,
				},
				resource.Attribute{
					Name:        "enable_nat_traversal",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable NAT traversal.`,
				},
				resource.Attribute{
					Name:        "nat_traversal_keep_alive",
					Description: `(Optional, int) Sending interval for NAT keep-alive packets (in seconds). For versions 6.1 - 8.1, this param, if specified, should be a multiple of 10 between 10 and 3600 to be valid.`,
				},
				resource.Attribute{
					Name:        "nat_traversal_enable_udp_checksum",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable NAT traversal UDP checksum.`,
				},
				resource.Attribute{
					Name:        "enable_fragmentation",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable fragmentation.`,
				},
				resource.Attribute{
					Name:        "ikev1_exchange_mode",
					Description: `(Optional) The IKEv1 exchange mode.`,
				},
				resource.Attribute{
					Name:        "ikev1_crypto_profile",
					Description: `(Optional) IKEv1 crypto profile.`,
				},
				resource.Attribute{
					Name:        "enable_dead_peer_detection",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable dead peer detection.`,
				},
				resource.Attribute{
					Name:        "dead_peer_detection_interval",
					Description: `(Optional, int) The dead peer detection interval.`,
				},
				resource.Attribute{
					Name:        "dead_peer_detection_retry",
					Description: `(Optional, int) Number of retries before disconnection.`,
				},
				resource.Attribute{
					Name:        "ikev2_crypto_profile",
					Description: `(Optional, PAN-OS 7.0+) IKEv2 crypto profile.`,
				},
				resource.Attribute{
					Name:        "ikev2_cookie_validation",
					Description: `(Optional, PAN-OS 7.0+) Set to ` + "`" + `true` + "`" + ` to require cookie.`,
				},
				resource.Attribute{
					Name:        "enable_liveness_check",
					Description: `(Optional, , PAN-OS 7.0+bool) Set to ` + "`" + `true` + "`" + ` to enable sending empty information liveness check message.`,
				},
				resource.Attribute{
					Name:        "liveness_check_interval",
					Description: `(Optional, , PAN-OS 7.0+int) Delay interval before sending probing packets (in seconds).`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ipsec_crypto_profile",
			Category:         "Firewall Resources",
			ShortDescription: `Manages IPSec crypto profiles.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"ipsec",
				"crypto",
				"profile",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol. Valid values are ` + "`" + `esp` + "`" + ` (the default) or ` + "`" + `ah` + "`" + ``,
				},
				resource.Attribute{
					Name:        "authentications",
					Description: `(Required, list) - List of authentication types.`,
				},
				resource.Attribute{
					Name:        "encryptions",
					Description: `(Required, list) - List of encryption types. Valid values are ` + "`" + `des` + "`" + `, ` + "`" + `3des` + "`" + `, ` + "`" + `aes-128-cbc` + "`" + `, ` + "`" + `aes-192-cbc` + "`" + `, ` + "`" + `aes-256-cbc` + "`" + `, ` + "`" + `aes-128-gcm` + "`" + `, ` + "`" + `aes-256-gcm` + "`" + `, and ` + "`" + `null` + "`" + `. Note that the "gcm" values are only available in PAN-OS 7.0+.`,
				},
				resource.Attribute{
					Name:        "dh_group",
					Description: `(Optional) The DH group value. Valid values should start with the string ` + "`" + `group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifetime_type",
					Description: `(Optional) The lifetime type. Valid values are ` + "`" + `seconds` + "`" + `, ` + "`" + `minutes` + "`" + `, ` + "`" + `hours` + "`" + ` (the default), or ` + "`" + `days` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifetime_value",
					Description: `(Optional, int) The lifetime value.`,
				},
				resource.Attribute{
					Name:        "lifesize_type",
					Description: `(Optional) The lifesize type. Valid values are ` + "`" + `kb` + "`" + `, ` + "`" + `mb` + "`" + `, ` + "`" + `gb` + "`" + `, or ` + "`" + `tb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifesize_value",
					Description: `(Optional, int) the lifesize value.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ipsec_tunnel",
			Category:         "Firewall Resources",
			ShortDescription: `Manages IPSec tunnels.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"ipsec",
				"tunnel",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "tunnel_interface",
					Description: `(Required) The tunnel interface.`,
				},
				resource.Attribute{
					Name:        "anti_replay",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable Anti-Replay check on this tunnel.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to enable IPv6.`,
				},
				resource.Attribute{
					Name:        "copy_tos",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to copy IP TOS bits from inner packet to IPSec packet (not recommended).`,
				},
				resource.Attribute{
					Name:        "copy_flow_label",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to copy IPv6 flow label for 6in6 tunnel from inner packet to IPSec packet (not recommended).`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to disable this IPSec tunnel.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type. Valid values are ` + "`" + `auto-key` + "`" + ` (the default), ` + "`" + `manual-key` + "`" + `, or ` + "`" + `global-protect-satellite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ak_ike_gateway",
					Description: `(Optional) IKE gateway name.`,
				},
				resource.Attribute{
					Name:        "ak_ipsec_crypto_profile",
					Description: `(Optional) IPSec crypto profile name.`,
				},
				resource.Attribute{
					Name:        "mk_local_spi",
					Description: `(Optional) Outbound SPI, hex format.`,
				},
				resource.Attribute{
					Name:        "mk_remote_spi",
					Description: `(Optional) Inbound SPI, hex format.`,
				},
				resource.Attribute{
					Name:        "mk_local_address_ip",
					Description: `(Optional) Specify exact IP address if interface has multiple addresses.`,
				},
				resource.Attribute{
					Name:        "mk_local_address_floating_ip",
					Description: `(Optional) Floating IP address in HA Active-Active configuration.`,
				},
				resource.Attribute{
					Name:        "mk_protocol",
					Description: `(Optional) Manual key protocol. Valid valies are ` + "`" + `esp` + "`" + ` or ` + "`" + `ah` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mk_auth_type",
					Description: `(Optional) Authentication algorithm. Valid values are ` + "`" + `md5` + "`" + `, ` + "`" + `sha1` + "`" + `, ` + "`" + `sha256` + "`" + `, ` + "`" + `sha384` + "`" + `, ` + "`" + `sha512` + "`" + `, or ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mk_auth_key",
					Description: `(Optional) The auth key for the given auth type.`,
				},
				resource.Attribute{
					Name:        "mk_esp_encryption_type",
					Description: `(Optional) The encryption algorithm. Valid values are ` + "`" + `des` + "`" + `, ` + "`" + `3des` + "`" + `, ` + "`" + `aes-128-cbc` + "`" + `, ` + "`" + `aes-192-cbc` + "`" + `, ` + "`" + `aes-256-cbc` + "`" + `, or ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mk_esp_encryption_key",
					Description: `(Optional) The encryption key.`,
				},
				resource.Attribute{
					Name:        "gps_interface",
					Description: `(Optional) Interface to communicate with portal.`,
				},
				resource.Attribute{
					Name:        "gps_portal_address",
					Description: `(Optional) GlobalProtect portal address.`,
				},
				resource.Attribute{
					Name:        "gps_prefer_ipv6",
					Description: `(Optional, PAN-OS 8.0+, bool) Prefer to register the portal in IPv6. Only applicable to FQDN portal-address.`,
				},
				resource.Attribute{
					Name:        "gps_interface_ip_ipv4",
					Description: `(Optional) specify exact IP address if interface has multiple addresses (IPv4).`,
				},
				resource.Attribute{
					Name:        "gps_interface_ip_ipv6",
					Description: `(Optional, PAN-OS 8.0+) specify exact IP address if interface has multiple addresses (IPv6).`,
				},
				resource.Attribute{
					Name:        "gps_interface_floating_ip_ipv4",
					Description: `(Optional, PAN-OS 7.0+) Floating IPv4 address in HA Active-Active configuration.`,
				},
				resource.Attribute{
					Name:        "gps_interface_floating_ip_ipv6",
					Description: `(Optional, PAN-OS 8.0+) Floating IPv6 address in HA Active-Active configuration.`,
				},
				resource.Attribute{
					Name:        "gps_publish_connected_routes",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to to publish connected and static routes.`,
				},
				resource.Attribute{
					Name:        "gps_publish_routes",
					Description: `(Optional, list) Specify list of routes to publish to Global Protect Gateway.`,
				},
				resource.Attribute{
					Name:        "gps_local_certificate",
					Description: `(Optional) GlobalProtect satellite certificate file name.`,
				},
				resource.Attribute{
					Name:        "gps_certificate_profile",
					Description: `(Optional) Profile for authenticating GlobalProtect gateway certificates.`,
				},
				resource.Attribute{
					Name:        "enable_tunnel_monitor",
					Description: `(Optional, bool) Enable tunnel monitoring on this tunnel.`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_destination_ip",
					Description: `(Optional) Destination IP to send ICMP probe.`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_source_ip",
					Description: `(Optional) Source IP to send ICMP probe`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_profile",
					Description: `(Optional) Tunnel monitor profile.`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_proxy_id",
					Description: `(Optional, PAN-OS 7.0+) Which proxy-id (or proxy-id-v6) the monitoring traffic will use.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ipsec_tunnel_proxy_id_ipv4",
			Category:         "Firewall Resources",
			ShortDescription: `Manages IPv4 proxy IDs for auto key IPSec tunnels.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"ipsec",
				"tunnel",
				"proxy",
				"id",
				"ipv4",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "ipsec_tunnel",
					Description: `(Required) The auto key IPSec tunnel to attach this proxy ID to.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) IP subnet or IP address represents local network.`,
				},
				resource.Attribute{
					Name:        "remote",
					Description: `(Optional) IP subnet or IP address represents remote network.`,
				},
				resource.Attribute{
					Name:        "protocol_any",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` for any IP protocol.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `(Optional, int) IP protocol number.`,
				},
				resource.Attribute{
					Name:        "protocol_tcp_local",
					Description: `(Optional, int) Local TCP port number.`,
				},
				resource.Attribute{
					Name:        "protocol_tcp_remote",
					Description: `(Optional, int) Remote TCP port number.`,
				},
				resource.Attribute{
					Name:        "protocol_udp_local",
					Description: `(Optional, int) Local UDP port number.`,
				},
				resource.Attribute{
					Name:        "protocol_udp_remote",
					Description: `(Optional, int) Remote UDP port number.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_license_api_key",
			Category:         "Firewall Resources",
			ShortDescription: `Manages the licensing API key.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"license",
				"api",
				"key",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The licensing API key.`,
				},
				resource.Attribute{
					Name:        "retain_key",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to retain the licensing API key even after the deletion of this resource (recommended).`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_licensing",
			Category:         "Firewall Resources",
			ShortDescription: `Manages PAN-OS licensing.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"licensing",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "auth_codes",
					Description: `(Required) The list of auth codes to install.`,
				},
				resource.Attribute{
					Name:        "delicense",
					Description: `(Optional, bool) Leave as ` + "`" + `true` + "`" + ` if you want to delicense the firewall when this resource is removed, otherwise set to ` + "`" + `false` + "`" + ` to prevent firewall delicensing. Delicensing requires that the licensing API key has been installed.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) For ` + "`" + `delicense` + "`" + ` of ` + "`" + `true` + "`" + `, the type of delicensing to perform. Right now, only ` + "`" + `auto` + "`" + ` is supported (no manual delicensing). ## Attribute Reference The following attributes are available after read operations:`,
				},
				resource.Attribute{
					Name:        "licenses",
					Description: `List of licenses. Licenses have the following attributes:`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `The feature name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `License description.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial number.`,
				},
				resource.Attribute{
					Name:        "issued",
					Description: `When the license was issued.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `When the license expires.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `If the license has expired or not.`,
				},
				resource.Attribute{
					Name:        "auth_code",
					Description: `Associated auth code (if applicable).`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "licenses",
					Description: `List of licenses. Licenses have the following attributes:`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `The feature name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `License description.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial number.`,
				},
				resource.Attribute{
					Name:        "issued",
					Description: `When the license was issued.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `When the license expires.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `If the license has expired or not.`,
				},
				resource.Attribute{
					Name:        "auth_code",
					Description: `Associated auth code (if applicable).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_loopback_interface",
			Category:         "Firewall Resources",
			ShortDescription: `Manages loopback interfaces.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"loopback",
				"interface",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name. This must start with ` + "`" + `loopback.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys that will use this interface (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional, bool) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv6 MSS adjust value.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_management_profile",
			Category:         "Firewall Resources",
			ShortDescription: `Manages interface management profiles.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"management",
				"profile",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The management profile's name.`,
				},
				resource.Attribute{
					Name:        "ping",
					Description: `(Optional) Allow ping.`,
				},
				resource.Attribute{
					Name:        "telnet",
					Description: `(Optional) Allow telnet.`,
				},
				resource.Attribute{
					Name:        "ssh",
					Description: `(Optional) Allow SSH.`,
				},
				resource.Attribute{
					Name:        "http",
					Description: `(Optional) Allow HTTP.`,
				},
				resource.Attribute{
					Name:        "http_ocsp",
					Description: `(Optional) Allow HTTP OCSP.`,
				},
				resource.Attribute{
					Name:        "https",
					Description: `(Optional) Allow HTTPS.`,
				},
				resource.Attribute{
					Name:        "snmp",
					Description: `(Optional) Allow SNMP.`,
				},
				resource.Attribute{
					Name:        "response_pages",
					Description: `(Optional) Allow response pages.`,
				},
				resource.Attribute{
					Name:        "userid_service",
					Description: `(Optional) Allow User ID service.`,
				},
				resource.Attribute{
					Name:        "userid_syslog_listener_ssl",
					Description: `(Optional) Allow User ID syslog listener for SSL.`,
				},
				resource.Attribute{
					Name:        "userid_syslog_listener_udp",
					Description: `(Optional) Allow User ID syslog listener for UDP.`,
				},
				resource.Attribute{
					Name:        "permitted_ips",
					Description: `(Optional) The list of permitted IP addresses or address ranges for this management profile.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_nat_rule",
			Category:         "Firewall Resources",
			ShortDescription: `Manages NAT rules.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"nat",
				"rule",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The NAT rule's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the NAT rule into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rulebase",
					Description: `(Optional, Deprecated) The rulebase. For firewalls, there is only the ` + "`" + `rulebase` + "`" + ` value (default), but on Panorama, there is also ` + "`" + `pre-rulebase` + "`" + ` and ` + "`" + `post-rulebase` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional). NAT type. This can be ` + "`" + `ipv4` + "`" + ` (default), ` + "`" + `nat64` + "`" + `, or ` + "`" + `nptv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `(Required) The list of source zone(s).`,
				},
				resource.Attribute{
					Name:        "destination_zone",
					Description: `(Required) The destination zone.`,
				},
				resource.Attribute{
					Name:        "to_interface",
					Description: `(Optional) Egress interface from route lookup (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Service (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Required) List of source address(es).`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Required) List of destination address(es).`,
				},
				resource.Attribute{
					Name:        "sat_type",
					Description: `(Optional) Type of source address translation. This can be ` + "`" + `none` + "`" + ` (default), ` + "`" + `dynamic-ip-and-port` + "`" + `, ` + "`" + `dynamic-ip` + "`" + `, or ` + "`" + `static-ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_address_type",
					Description: `(Optional) Source address translation address type. This can be ` + "`" + `interface-address` + "`" + ` or ` + "`" + `translated-address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_translated_addresses",
					Description: `(Optional) Source address translation list of translated addresses.`,
				},
				resource.Attribute{
					Name:        "sat_interface",
					Description: `(Optional) Source address translation interface.`,
				},
				resource.Attribute{
					Name:        "sat_ip_address",
					Description: `(Optional) Source address translation IP address.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_type",
					Description: `(Optional) Source address translation fallback type. This can be ` + "`" + `none` + "`" + `, ` + "`" + `interface-address` + "`" + `, or ` + "`" + `translated-address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_translated_addresses",
					Description: `(Optional) Source address translation list of fallback translated addresses.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_interface",
					Description: `(Optional) Source address translation fallback interface.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_ip_type",
					Description: `(Optional) Source address translation fallback IP type. This can be ` + "`" + `ip` + "`" + ` or ` + "`" + `floating` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_ip_address",
					Description: `(Optional) The source address translation fallback IP address.`,
				},
				resource.Attribute{
					Name:        "sat_static_translated_address",
					Description: `(Optional) The statically translated source address.`,
				},
				resource.Attribute{
					Name:        "sat_static_bi_directional",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable bi-directional source address translation.`,
				},
				resource.Attribute{
					Name:        "dat_type",
					Description: `(Optional) Destination address translation type. This should be either ` + "`" + `static` + "`" + ` or ` + "`" + `dynamic` + "`" + `. The ` + "`" + `dynamic` + "`" + ` option is only available on PAN-OS 8.1+.`,
				},
				resource.Attribute{
					Name:        "dat_address",
					Description: `(Optional) Destination address translation's address. Requires ` + "`" + `dat_type` + "`" + ` be set to "static" or "dynamic".`,
				},
				resource.Attribute{
					Name:        "dat_port",
					Description: `(Optional) Destination address translation's port number. Requires ` + "`" + `dat_type` + "`" + ` be set to "static" or "dynamic".`,
				},
				resource.Attribute{
					Name:        "dat_dynamic_distribution",
					Description: `(Optional, PAN-OS 8.1+) Distribution algorithm for destination address pool. The PAN-OS 8.1 GUI doesn't seem to set this anywhere, but this is added here for completeness' sake. Requires ` + "`" + `dat_type` + "`" + ` of "dynamic".`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to disable this rule.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_nat_rule_group",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a group of NAT rules.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"nat",
				"rule",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the NAT rule group into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `(Optional) A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `(Optional) Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Repeatable) The rule definition (see below). The rule ordering will match how they appear in the terraform plan file. Each ` + "`" + `rule` + "`" + ` defined supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The NAT rule's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional). NAT type. This can be ` + "`" + `ipv4` + "`" + ` (default), ` + "`" + `nat64` + "`" + `, or ` + "`" + `nptv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to disable this rule.`,
				},
				resource.Attribute{
					Name:        "original_packet",
					Description: `(Required) The original packet specification (see below).`,
				},
				resource.Attribute{
					Name:        "translated_packet",
					Description: `(Required) The translated packet spec (see below). ` + "`" + `original_packet` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `(Required) The list of source zone(s).`,
				},
				resource.Attribute{
					Name:        "destination_zone",
					Description: `(Required) The destination zone.`,
				},
				resource.Attribute{
					Name:        "destination_interface",
					Description: `(Optional) Egress interface from route lookup (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Service (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Required) List of source address(es).`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Required) List of destination address(es). ` + "`" + `translated_packet` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The source spec (see below). Leave this empty for a destination NAT of "none".`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) The destination spec (see below). Leave this empty for a destination NAT of "none". ` + "`" + `translated_packet.source` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "dynamic_ip_and_port",
					Description: `(Optional) Dynamic IP and port source translation spec (see below).`,
				},
				resource.Attribute{
					Name:        "dynamic_ip",
					Description: `(Optional) Dynamic IP source translation spec (see below).`,
				},
				resource.Attribute{
					Name:        "static_ip",
					Description: `(Optional) Static IP source translation spec (see below). ` + "`" + `translated_packet.source.dynamic_ip_and_port` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `(Optional) Translated address source translation type spec (see below).`,
				},
				resource.Attribute{
					Name:        "interface_address",
					Description: `(Optional) Interface address source translation type spec (see below). ` + "`" + `translated_packet.source.dynamic_ip_and_port.translated_address` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `(Required) List of translated addresses. ` + "`" + `translated_packet.source.dynamic_ip_and_port.interface_address` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address. ` + "`" + `translated_packet.source.dynamic_ip` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `(Optional) The list of translated addresses.`,
				},
				resource.Attribute{
					Name:        "fallback",
					Description: `(Optional) The fallback spec (see below). Leaving this empty or omiting it means a fallback of "None". ` + "`" + `translated_packet.source.dynamic_ip.fallback` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `(Optional) The translated address fallback spec (see below).`,
				},
				resource.Attribute{
					Name:        "interface_address",
					Description: `(Optional) The interface address fallback spec (see below). ` + "`" + `translated_packet.source.dynamic_ip.fallback.translated_address` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `(Optional) List of source address translation fallback translated addresses. ` + "`" + `translated_packet.source.dynamic_ip.fallback.interface_address` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) Source address translation fallback interface.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of interface fallback. Valid values are ` + "`" + `ip` + "`" + ` (default) or ` + "`" + `floating` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP address of the fallback interface. ` + "`" + `translated_packet.source.static_ip` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `(Required) The statically translated source address.`,
				},
				resource.Attribute{
					Name:        "bi_directional",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable bi-directional source address translation. ` + "`" + `translated_packet.destination` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "static",
					Description: `(Optional) Specifies a static destination NAT (see below).`,
				},
				resource.Attribute{
					Name:        "dynamic",
					Description: `(Optional, PAN-OS 8.1+) Specify a dynamic destination NAT (see below). ` + "`" + `translated_packet.destination.static` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Destination address translation address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, int) Destination address translation port number. ` + "`" + `translated_packet.destination.dynamic` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Destination address translation address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, int) Destination address translation port number.`,
				},
				resource.Attribute{
					Name:        "distribution",
					Description: `(Optional, PAN-OS 8.1+) Distribution algorithm for destination address pool. The PAN-OS 8.1 GUI doesn't seem to set this anywhere, but this is added here for completeness' sake. The GUI sets this to ` + "`" + `round-robin` + "`" + ` currently.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_address_group",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama address groups.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"address",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The address group's name.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group to put the address group into (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_addresses",
					Description: `(Optional) The address objects to include in this statically defined address group.`,
				},
				resource.Attribute{
					Name:        "dynamic_match",
					Description: `(Optional) The IP tags to include in this DAG.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The address group's description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_address_object",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama address objects.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"address",
				"object",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The address object's name.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group to put the address object into (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of address object. This can be ` + "`" + `ip-netmask` + "`" + ` (default), ` + "`" + `ip-range` + "`" + `, or ` + "`" + `fqdn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The address object's value. This can take various forms depending on what type of address object this is, but can be something like ` + "`" + `192.168.80.150` + "`" + ` or ` + "`" + `192.168.80.0/24` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The address object's description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_administrative_tag",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama administrative tags.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"administrative",
				"tag",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The administrative tag's name.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group to put the administrative tag into (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) The tag's color. This should be either an empty string (no color) or a string such as ` + "`" + `color1` + "`" + ` or ` + "`" + `color15` + "`" + `. Note that for maximum portability, you should limit color usage to ` + "`" + `color16` + "`" + `, which was available in PAN-OS 6.1. PAN-OS 8.1's colors go up to ` + "`" + `color42` + "`" + `. The value ` + "`" + `color18` + "`" + ` is reserved internally by PAN-OS and thus not available for use.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The administrative tag's description.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bfd_profile",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama BFD profiles.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bfd",
				"profile",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The BBFD profile's name.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) BFD operation mode. Valid values are ` + "`" + `active` + "`" + ` (default) or ` + "`" + `passive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minimum_tx_interval",
					Description: `(Optional, int) Desired minimum TX interval in ms. Default is ` + "`" + `1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minimum_rx_interval",
					Description: `(Optional, int) Required minimum RX interval in ms. Default is ` + "`" + `1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "detection_multiplier",
					Description: `(Optional, int) Multiplier sent to remote system. Default is ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hold_time",
					Description: `(Optional, int) Delay transmission and reception of control packets in ms.`,
				},
				resource.Attribute{
					Name:        "minimum_rx_ttl",
					Description: `(Optional, int) Minimum accepted ttl on received BFD packet.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a Panorama virtual router's BGP configuration.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bgp",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP configuration to.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable BGP or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Optional) Router ID of this BGP instance.`,
				},
				resource.Attribute{
					Name:        "as_number",
					Description: `(Optional) Local AS number.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `(Optional, PAN-OS 7.1+) BFD configuration.`,
				},
				resource.Attribute{
					Name:        "reject_default_route",
					Description: `(Optional, bool) Do not learn default route from BGP (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "install_route",
					Description: `(Optional, bool) Populate BGP learned route to global route table.`,
				},
				resource.Attribute{
					Name:        "aggregate_med",
					Description: `(Optional, bool) Aggregate route only if they have same MED attributes (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "default_local_preference",
					Description: `(Optional) Default local preference (default: ` + "`" + `"100"` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_format",
					Description: `(Optional) AS format. Valid values are ` + "`" + `"2-byte"` + "`" + ` (default) or ` + "`" + `"4-byte"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "always_compare_med",
					Description: `(Optional, bool) Always compare MEDs.`,
				},
				resource.Attribute{
					Name:        "deterministic_med_comparison",
					Description: `(Optional, bool) Deterministic MED comparison (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ecmp_multi_as",
					Description: `(Optional, bool) Support multiple AS in ECMP.`,
				},
				resource.Attribute{
					Name:        "enforce_first_as",
					Description: `(Optional, bool) Enforce First AS for EBGP (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_graceful_restart",
					Description: `(Optional, bool) Enable graceful restart (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "stale_route_time",
					Description: `(Optional, int) Time to remove stale routes after peer restart, in seconds (default: ` + "`" + `120` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "local_restart_time",
					Description: `(Optional, int) Local restart time to advertise to peer, in seconds (default: ` + "`" + `120` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "max_peer_restart_time",
					Description: `(Optional, int) Maximum of peer restart time accepted, in seconds (default: ` + "`" + `120` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "reflector_cluster_id",
					Description: `(Optional) Route reflector cluster ID.`,
				},
				resource.Attribute{
					Name:        "confederation_member_as",
					Description: `(Optional) Confederation requires member-AS number.`,
				},
				resource.Attribute{
					Name:        "allow_redistribute_default_route",
					Description: `(Optional, bool) Allow redistribute default route to BGP.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_aggregate",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama BGP address aggregation rules.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bgp",
				"aggregate",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to put the rule into.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The rule name.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Aggregating address prefix.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this rule (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "as_set",
					Description: `(Optional, bool) Generate AS-set attribute.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `(Optional, bool) Summarize route.`,
				},
				resource.Attribute{
					Name:        "local_preference",
					Description: `(Optional) New local preference value.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) New MED value.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional, int) New weight value.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop address.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) New route origin. Valid values are ` + "`" + `incomplete` + "`" + ` (default), ` + "`" + `igp` + "`" + `, or ` + "`" + `egp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_limit",
					Description: `(Optional, int) Add AS path limit attribute if it does not exist.`,
				},
				resource.Attribute{
					Name:        "as_path_type",
					Description: `(Optional) AS path update options. Valid values are ` + "`" + `none` + "`" + ` (default) or ` + "`" + `prepend` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_value",
					Description: `(Optional) For ` + "`" + `as_path_type` + "`" + ` of ` + "`" + `prepend` + "`" + `, the value to prepend.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community update options. Valid values are ` + "`" + `none` + "`" + ` (default), ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_value",
					Description: `(Optional) If ` + "`" + `community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. For the ` + "`" + `append` + "`" + ` and ` + "`" + `overwrite` + "`" + ` types specifically, valid values are ` + "`" + `no-export` + "`" + `, ` + "`" + `no-advertise` + "`" + `, ` + "`" + `local-as` + "`" + `, or ` + "`" + `nopeer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_type",
					Description: `(Optional) Extended community update options. Valid values are ` + "`" + `none` + "`" + ` (default), ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_vaule",
					Description: `(Optional) If ` + "`" + `extended_community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_aggregate_advertise_filter",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a Panorama route advertise filter for a BGP address aggregation rule.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bgp",
				"aggregate",
				"advertise",
				"filter",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_aggregate",
					Description: `(Required) The BGP address aggregation rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below).`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry. Each ` + "`" + `address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_aggregate_suppress_filter",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a Panorama route suppression filter for a BGP address aggregation rule.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bgp",
				"aggregate",
				"suppress",
				"filter",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_aggregate",
					Description: `(Required) The BGP address aggregation rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below).`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry. Each ` + "`" + `address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_auth_profile",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a Panorama BGP auth profile.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bgp",
				"auth",
				"profile",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP auth profile to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Shared secret for the TCP MD5 authentication.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_conditional_adv",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a Panorama BGP conditional advertisement.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bgp",
				"conditional",
				"adv",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP conditional advertisement to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "used_by",
					Description: `(Optional) List of BGP peer groups that use this rule.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_conditional_adv_advertise_filter",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a Panorama advertise filter for a BGP conditional advertisement.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bgp",
				"conditional",
				"adv",
				"advertise",
				"filter",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_conditional_adv",
					Description: `(Required) The BGP conditional advertisement to add this filter to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "address_prefixes",
					Description: `(Optional) List of matching address prefixes.`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_conditional_adv_non_exist_filter",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a Panorama non-exist filter for a BGP conditional advertisement.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bgp",
				"conditional",
				"adv",
				"non",
				"exist",
				"filter",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_conditional_adv",
					Description: `(Required) The BGP conditional advertisement to add this filter to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "address_prefixes",
					Description: `(Optional) List of matching address prefixes.`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_dampening_profile",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a Panorama BGP dampening profile.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bgp",
				"dampening",
				"profile",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP dampening profile to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "cutoff",
					Description: `(Optional, float) Cutoff threshold value (default: ` + "`" + `1.25` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "reuse",
					Description: `(Optional, float) Reuse threshold value (default: ` + "`" + `0.5` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "max_hold_time",
					Description: `(Optional, int) Maximum hold-down time, in seconds (default: ` + "`" + `900` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "decay_half_life_reachable",
					Description: `(Optional, int) Decay half-life while reachable, in seconds (default: ` + "`" + `300` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "decay_half_life_unreachable",
					Description: `(Optional, int) Decay half-life while unreachable, in seconds (default: ` + "`" + `900` + "`" + `).`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_export_rule_group",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama BGP export rule groups.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bgp",
				"export",
				"rule",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to put the rule into.`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `(Optional) A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `(Optional) Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The export rule definition (see below). The export rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The security rule name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this export rule (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "used_by",
					Description: `(Optional) List of auth profiles.`,
				},
				resource.Attribute{
					Name:        "match_as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "match_community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "match_extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "match_med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "match_route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "match_address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below). below for the params for this section.`,
				},
				resource.Attribute{
					Name:        "match_next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "match_from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Rule action. Valid values are ` + "`" + `allow` + "`" + ` (default) or ` + "`" + `deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dampening",
					Description: `(Optional) Route flap dampening profile.`,
				},
				resource.Attribute{
					Name:        "local_preference",
					Description: `(Optional) New local preference value.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) New MED value.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional, int) New weight value.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop address.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) New route origin. Valid values are ` + "`" + `igp` + "`" + `, ` + "`" + `egp` + "`" + `, or ` + "`" + `incomplete` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_limit",
					Description: `(Optional, int) Add AS path limit attribute if it does not exist.`,
				},
				resource.Attribute{
					Name:        "as_path_type",
					Description: `(Optional) AS path update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove` + "`" + `, ` + "`" + `prepend` + "`" + ` or ` + "`" + `remove-and-prepend` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_value",
					Description: `(Optional) If ` + "`" + `as_path_type` + "`" + ` is ` + "`" + `prepend` + "`" + ` or ` + "`" + `remove-and-prepend` + "`" + `, the value to prepend.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_value",
					Description: `(Optional) If ` + "`" + `community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. For the ` + "`" + `append` + "`" + ` and ` + "`" + `overwrite` + "`" + ` types specifically, valid values for ` + "`" + `community_value` + "`" + ` are ` + "`" + `no-export` + "`" + `, ` + "`" + `no-advertise` + "`" + `, ` + "`" + `local-as` + "`" + `, or ` + "`" + `nopeer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_type",
					Description: `(Optional) Extended community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_vaule",
					Description: `(Optional) If ` + "`" + `extended_community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. Each ` + "`" + `match_address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_import_rule_group",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama BGP import rule groups.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bgp",
				"import",
				"rule",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to put the rule into.`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `(Optional) A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `(Optional) Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The import rule definition (see below). The import rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The security rule name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this import rule (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "used_by",
					Description: `(Optional) List of auth profiles.`,
				},
				resource.Attribute{
					Name:        "match_as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "match_community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "match_extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "match_med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "match_route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "match_address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below). below for the params for this section.`,
				},
				resource.Attribute{
					Name:        "match_next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "match_from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Rule action. Valid values are ` + "`" + `allow` + "`" + ` (default) or ` + "`" + `deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dampening",
					Description: `(Optional) Route flap dampening profile.`,
				},
				resource.Attribute{
					Name:        "local_preference",
					Description: `(Optional) New local preference value.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) New MED value.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional, int) New weight value.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop address.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) New route origin. Valid values are ` + "`" + `igp` + "`" + `, ` + "`" + `egp` + "`" + `, or ` + "`" + `incomplete` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_limit",
					Description: `(Optional, int) Add AS path limit attribute if it does not exist.`,
				},
				resource.Attribute{
					Name:        "as_path_type",
					Description: `(Optional) AS path update options. Valid values are ` + "`" + `none` + "`" + ` or ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_value",
					Description: `(Optional) If ` + "`" + `community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. For the ` + "`" + `append` + "`" + ` and ` + "`" + `overwrite` + "`" + ` types specifically, valid values for ` + "`" + `community_value` + "`" + ` are ` + "`" + `no-export` + "`" + `, ` + "`" + `no-advertise` + "`" + `, ` + "`" + `local-as` + "`" + `, or ` + "`" + `nopeer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_type",
					Description: `(Optional) Extended community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_vaule",
					Description: `(Optional) If ` + "`" + `extended_community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. Each ` + "`" + `match_address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_peer",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a Panorama BGP peer.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bgp",
				"peer",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP peer to.`,
				},
				resource.Attribute{
					Name:        "bgp_peer_group",
					Description: `(Required) The BGP peer group to put this peer into.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "peer_as",
					Description: `(Optional) Peer AS number.`,
				},
				resource.Attribute{
					Name:        "local_address_interface",
					Description: `(Required) Interface to accept BGP session.`,
				},
				resource.Attribute{
					Name:        "local_address_ip",
					Description: `(Optional) Specify exact IP address if interface has multiple addresses.`,
				},
				resource.Attribute{
					Name:        "peer_address_ip",
					Description: `(Required) Peer IP address configuration.`,
				},
				resource.Attribute{
					Name:        "reflector_client",
					Description: `(Optional) This peer is reflector client. Valid values are ` + "`" + `non-client` + "`" + `, ` + "`" + `client` + "`" + `, or ` + "`" + `meshed-client` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "peering_type",
					Description: `(Optional) Peering type that affects NOPEER community value handling. Valid values are ` + "`" + `unspecified` + "`" + ` (default) or ` + "`" + `bilateral` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_prefixes",
					Description: `(Optional) Maximum of prefixes to receive from the peer. This can be a number such as ` + "`" + `"5000"` + "`" + ` (default) or ` + "`" + `unlimited` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_profile",
					Description: `(Optional) Auth profile.`,
				},
				resource.Attribute{
					Name:        "keep_alive_interval",
					Description: `(Optional, int) Keep alive interval, in seconds (default: ` + "`" + `30` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "multi_hop",
					Description: `(Optional, int) IP TTL value used for sending BGP packet.`,
				},
				resource.Attribute{
					Name:        "open_delay_time",
					Description: `(Optional, int) Open delay time, in seconds.`,
				},
				resource.Attribute{
					Name:        "hold_time",
					Description: `(Optional, int) Hold time, in seconds.`,
				},
				resource.Attribute{
					Name:        "idle_hold_time",
					Description: `(Optional, int) Idle hold time, in seconds.`,
				},
				resource.Attribute{
					Name:        "allow_incoming_connections",
					Description: `(Optional, bool) Allow incoming connections (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "incoming_connections_remote_port",
					Description: `(Optional, int) Restrict remote port for incoming BGP connections.`,
				},
				resource.Attribute{
					Name:        "allow_outgoing_connections",
					Description: `(Optional, bool) Allow outgoing connections (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "outgoing_connections_local_port",
					Description: `(Optional, int) Use specific local port for outgoing BGP connections.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `(Optional, PAN-OS 7.1+) BFD profile. This can be a specific BFD profile name, ` + "`" + `None` + "`" + ` (disables BFD), or ` + "`" + `Inherit-vr-global-setting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_mp_bgp",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable MP BGP.`,
				},
				resource.Attribute{
					Name:        "address_family_type",
					Description: `(Optional, PAN-OS 8.0+) Set the AFI for this peer. Valid values are ` + "`" + `ipv4` + "`" + ` or ` + "`" + `ipv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subsequent_address_family_unicast",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable unicast subsequent address family for this peer.`,
				},
				resource.Attribute{
					Name:        "subsequent_address_family_multicast",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable multicast subsequent address family for this peer.`,
				},
				resource.Attribute{
					Name:        "enable_sender_side_loop_detection",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable sender side loop detection.`,
				},
				resource.Attribute{
					Name:        "min_route_advertisement_interval",
					Description: `(Optional, int, PAN-OS 8.1+) Minimum route advertisement interval, in seconds.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_peer_group",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a Panorama BGP peer group.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bgp",
				"peer",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP peer group to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "aggregated_confed_as_path",
					Description: `(Optional, bool) The peers understand aggregated confederation AS path (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "soft_reset_with_stored_info",
					Description: `(Optional, bool) Soft reset with stored info.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Peer group type. Valid options are ` + "`" + `ebgp` + "`" + ` (default), ` + "`" + `ebgp-confed` + "`" + `, ` + "`" + `ibgp` + "`" + `, or ` + "`" + `ibgp-confed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "export_next_hop",
					Description: `(Optional) Export next hop. Valid values are ` + "`" + `original` + "`" + `, ` + "`" + `use-self` + "`" + `, or ` + "`" + `resolve` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "import_next_hop",
					Description: `(Optional) Import next hop. Valid values are ` + "`" + `original` + "`" + `, ` + "`" + `use-peer` + "`" + `, or the empty string.`,
				},
				resource.Attribute{
					Name:        "remove_private_as",
					Description: `(Optional, bool) Remove private AS when exporting route. Only available for ` + "`" + `type=ebgp` + "`" + `.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_redist_rule",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a Panorama BGP redistribution rule.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"bgp",
				"redist",
				"rule",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP redist rule to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A subnet or a redistribution profile.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this rule or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `(Optional) The address family. Valid values are ` + "`" + `ipv4` + "`" + ` (default) or ` + "`" + `ipv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional, int) Metric value.`,
				},
				resource.Attribute{
					Name:        "set_origin",
					Description: `(Optional) Add the origin path attribute. Valid values are ` + "`" + `incomplete` + "`" + ` (default), ` + "`" + `igp` + "`" + `, or ` + "`" + `egp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "set_med",
					Description: `(Optional) Add the MULTI_EXIT_DISC path attribute.`,
				},
				resource.Attribute{
					Name:        "set_local_preference",
					Description: `(Optional) Add the LOCAL_PREF path attribute.`,
				},
				resource.Attribute{
					Name:        "set_as_path_limit",
					Description: `(Optional, int) Add the AS_PATHLIMIT path attribute.`,
				},
				resource.Attribute{
					Name:        "set_communities",
					Description: `(Optional) List of COMMUNITY path attributes to add.`,
				},
				resource.Attribute{
					Name:        "set_extended_communities",
					Description: `(Optional) List of EXTENDED COMMUNITY path attributes to add.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_device_group",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama device groups.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"device",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The device group's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The device group's description.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `The device definition (see below). The following arguments are valid for each ` + "`" + `device` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `(Optional) A subset of all available vsys on the firewall that should be in this device group. If the firewall is a virtual firewall, then this parameter should just be omitted.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_device_group_entry",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a specific device in a Panorama device group.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"device",
				"group",
				"entry",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "device_group",
					Description: `(Required) The device group's name.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `(Optional) A subset of all available vsys on the firewall that should be in this device group. If the firewall is a virtual firewall, then this parameter should just be omitted.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_edl",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama external dynamic lists.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"edl",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group (default: ` + "`" + `shared` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of EDL. This can be ` + "`" + `ip` + "`" + ` (the default; and the only valid value for PAN-OS 6.1 - 7.0), ` + "`" + `domain` + "`" + `, ` + "`" + `url` + "`" + `, or ` + "`" + `predefined` + "`" + ` (PAN-OS 8.0+)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The object's description.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) The EDL source URL`,
				},
				resource.Attribute{
					Name:        "certificate_profile",
					Description: `(Optional) Profile for authenticating client certificates`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) EDL username`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) EDL password`,
				},
				resource.Attribute{
					Name:        "repeat",
					Description: `(Optional) How often to retrieve the EDL. This can be ` + "`" + `hourly` + "`" + ` (the default), ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `, ` + "`" + `monthly` + "`" + `, or ` + "`" + `every five minutes` + "`" + ` (valid for PAN-OS 7.1+)`,
				},
				resource.Attribute{
					Name:        "repeat_at",
					Description: `(Optional) The time at which to retrieve the EDL. Please refer to the section above for how to set this value properly.`,
				},
				resource.Attribute{
					Name:        "repeat_day_of_week",
					Description: `(Optional) If ` + "`" + `repeat` + "`" + ` is ` + "`" + `weekly` + "`" + `, then this should be set to the desired day of the week. Valid values are ` + "`" + `sunday` + "`" + `, ` + "`" + `monday` + "`" + `, ` + "`" + `tuesday` + "`" + `, ` + "`" + `wednesday` + "`" + `, ` + "`" + `thursday` + "`" + `, ` + "`" + `friday` + "`" + `, ` + "`" + `saturday` + "`" + `, and ` + "`" + `sunday` + "`" + ``,
				},
				resource.Attribute{
					Name:        "repeat_day_of_month",
					Description: `(Optional, int) If ` + "`" + `repeat` + "`" + ` is ` + "`" + `monthly` + "`" + `, then this should be set to the desired day of the month.`,
				},
				resource.Attribute{
					Name:        "exceptions",
					Description: `(Optional, list) Provide a list of exception entries.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_ethernet_interface",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama ethernet interfaces.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"ethernet",
				"interface",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The ethernet interface's name. This should be something like ` + "`" + `ethernet1/X` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys that will use this interface (default: ` + "`" + `vsys1` + "`" + `). This should be something like ` + "`" + `vsys1` + "`" + ` or ` + "`" + `vsys3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The interface mode. This can be any of the following values: ` + "`" + `layer3` + "`" + `, ` + "`" + `layer2` + "`" + `, ` + "`" + `virtual-wire` + "`" + `, ` + "`" + `tap` + "`" + `, ` + "`" + `ha` + "`" + `, ` + "`" + `decrypt-mirror` + "`" + `, or ` + "`" + `aggregate-group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable DHCP on this interface.`,
				},
				resource.Attribute{
					Name:        "create_dhcp_default_route",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to create a DHCP default route.`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route_metric",
					Description: `(Optional) The metric for the DHCP default route.`,
				},
				resource.Attribute{
					Name:        "ipv6_enabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable IPv6.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "lldp_enabled",
					Description: `(Optional) Enable LLDP (default: false).`,
				},
				resource.Attribute{
					Name:        "lldp_profile",
					Description: `(Optional) LLDP profile.`,
				},
				resource.Attribute{
					Name:        "link_speed",
					Description: `(Optional) Link speed. This can be any of the following: ` + "`" + `10` + "`" + `, ` + "`" + `100` + "`" + `, ` + "`" + `1000` + "`" + `, or ` + "`" + `auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "link_duplex",
					Description: `(Optional) Link duplex setting. This can be ` + "`" + `full` + "`" + `, ` + "`" + `half` + "`" + `, or ` + "`" + `auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "link_state",
					Description: `(Optional) The link state. This can be ` + "`" + `up` + "`" + `, ` + "`" + `down` + "`" + `, or ` + "`" + `auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aggregate_group",
					Description: `(Optional) The aggregate group (applicable for physical firewalls only).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv6 MSS adjust value.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_ike_crypto_profile",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama IKE crypto profiles.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"ike",
				"crypto",
				"profile",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "dh_groups",
					Description: `(Required, list) List of DH Group entries. Values should have a prefix if ` + "`" + `group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "authentications",
					Description: `(Required, list) List of authentication types. This c`,
				},
				resource.Attribute{
					Name:        "encryptions",
					Description: `(Required, list) List of encryption types. Valid values are ` + "`" + `des` + "`" + `, ` + "`" + `3des` + "`" + `, ` + "`" + `aes-128-cbc` + "`" + `, ` + "`" + `aes-192-cbc` + "`" + `, and ` + "`" + `aes-256-cbc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifetime_type",
					Description: `(Optional) The lifetime type. Valid values are ` + "`" + `seconds` + "`" + `, ` + "`" + `minutes` + "`" + `, ` + "`" + `hours` + "`" + ` (the default), and ` + "`" + `days` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifetime_value",
					Description: `(Optional, int) The lifetime value.`,
				},
				resource.Attribute{
					Name:        "authentication_multiple",
					Description: `(Optional, PAN-OS 7.0+, int) IKEv2 SA reauthentication interval equals authetication-multiple`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_ike_gateway",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama IKE gateways.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"ike",
				"gateway",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional, PAN-OS 7.0+) The IKE gateway version. Valid values are ` + "`" + `ikev1` + "`" + `, (the default), ` + "`" + `ikev2` + "`" + `, or ` + "`" + `ikev2-preferred` + "`" + `. For PAN-OS 6.1, only ` + "`" + `ikev1` + "`" + ` is acceptable.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `(Optional, PAN-OS 7.0+, bool) Enable IPv6 or not.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to disable.`,
				},
				resource.Attribute{
					Name:        "peer_ip_type",
					Description: `(Optional) The peer IP type. Valid values are ` + "`" + `ip` + "`" + `, ` + "`" + `dynamic` + "`" + `, and ` + "`" + `fqdn` + "`" + ` (PANOS 8.1+).`,
				},
				resource.Attribute{
					Name:        "peer_ip_value",
					Description: `(Optional) The peer IP value.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface.`,
				},
				resource.Attribute{
					Name:        "local_ip_address_type",
					Description: `(Optional) The local IP address type. Valid values for this are ` + "`" + `ip` + "`" + `, ` + "`" + `floating-ip` + "`" + `, or an empty string (the default) which is ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_ip_address_value",
					Description: `(Optional) The IP address if ` + "`" + `local_ip_address_type` + "`" + ` is set to ` + "`" + `ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) The auth type. Valid values are ` + "`" + `pre-shared-key` + "`" + ` (the default), or ` + "`" + `certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(Optional) The pre-shared key value.`,
				},
				resource.Attribute{
					Name:        "local_id_type",
					Description: `(Optional) The local ID type. Valid values are ` + "`" + `ipaddr` + "`" + `, ` + "`" + `fqdn` + "`" + `, ` + "`" + `ufqdn` + "`" + `, ` + "`" + `keyid` + "`" + `, or ` + "`" + `dn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_id_value",
					Description: `(Optional) The local ID value.`,
				},
				resource.Attribute{
					Name:        "peer_id_type",
					Description: `(Optional) The peer ID type. Valid values are ` + "`" + `ipaddr` + "`" + `, ` + "`" + `fqdn` + "`" + `, ` + "`" + `ufqdn` + "`" + `, ` + "`" + `keyid` + "`" + `, or ` + "`" + `dn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "peer_id_value",
					Description: `(Optional) The peer ID value.`,
				},
				resource.Attribute{
					Name:        "peer_id_check",
					Description: `(Optional) Enable peer ID wildcard match for certificate authentication. Valid values are ` + "`" + `exact` + "`" + ` or ` + "`" + `wildcard` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_cert",
					Description: `(Optional) The local certificate name.`,
				},
				resource.Attribute{
					Name:        "cert_enable_hash_and_url",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to use hash-and-url for local certificate.`,
				},
				resource.Attribute{
					Name:        "cert_base_url",
					Description: `(Optional) The host and directory part of URL for local certificates.`,
				},
				resource.Attribute{
					Name:        "cert_use_management_as_source",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to use management interface IP as source to retrieve http certificates`,
				},
				resource.Attribute{
					Name:        "cert_permit_payload_mismatch",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to permit peer identification and certificate payload identification mismatch.`,
				},
				resource.Attribute{
					Name:        "cert_profile",
					Description: `(Optional) Profile for certificate valdiation during IKE negotiation.`,
				},
				resource.Attribute{
					Name:        "cert_enable_strict_validation",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable strict validation of peer's extended key use.`,
				},
				resource.Attribute{
					Name:        "enable_passive_mode",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable passive mode (responder only).`,
				},
				resource.Attribute{
					Name:        "enable_nat_traversal",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable NAT traversal.`,
				},
				resource.Attribute{
					Name:        "nat_traversal_keep_alive",
					Description: `(Optional, int) Sending interval for NAT keep-alive packets (in seconds). For versions 6.1 - 8.1, this param, if specified, should be a multiple of 10 between 10 and 3600 to be valid.`,
				},
				resource.Attribute{
					Name:        "nat_traversal_enable_udp_checksum",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable NAT traversal UDP checksum.`,
				},
				resource.Attribute{
					Name:        "enable_fragmentation",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable fragmentation.`,
				},
				resource.Attribute{
					Name:        "ikev1_exchange_mode",
					Description: `(Optional) The IKEv1 exchange mode.`,
				},
				resource.Attribute{
					Name:        "ikev1_crypto_profile",
					Description: `(Optional) IKEv1 crypto profile.`,
				},
				resource.Attribute{
					Name:        "enable_dead_peer_detection",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable dead peer detection.`,
				},
				resource.Attribute{
					Name:        "dead_peer_detection_interval",
					Description: `(Optional, int) The dead peer detection interval.`,
				},
				resource.Attribute{
					Name:        "dead_peer_detection_retry",
					Description: `(Optional, int) Number of retries before disconnection.`,
				},
				resource.Attribute{
					Name:        "ikev2_crypto_profile",
					Description: `(Optional, PAN-OS 7.0+) IKEv2 crypto profile.`,
				},
				resource.Attribute{
					Name:        "ikev2_cookie_validation",
					Description: `(Optional, PAN-OS 7.0+) Set to ` + "`" + `true` + "`" + ` to require cookie.`,
				},
				resource.Attribute{
					Name:        "enable_liveness_check",
					Description: `(Optional, , PAN-OS 7.0+bool) Set to ` + "`" + `true` + "`" + ` to enable sending empty information liveness check message.`,
				},
				resource.Attribute{
					Name:        "liveness_check_interval",
					Description: `(Optional, , PAN-OS 7.0+int) Delay interval before sending probing packets (in seconds).`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_ipsec_crypto_profile",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama IPSec crypto profiles.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"ipsec",
				"crypto",
				"profile",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol. Valid values are ` + "`" + `esp` + "`" + ` (the default) or ` + "`" + `ah` + "`" + ``,
				},
				resource.Attribute{
					Name:        "authentications",
					Description: `(Required, list) - List of authentication types.`,
				},
				resource.Attribute{
					Name:        "encryptions",
					Description: `(Required, list) - List of encryption types. Valid values are ` + "`" + `des` + "`" + `, ` + "`" + `3des` + "`" + `, ` + "`" + `aes-128-cbc` + "`" + `, ` + "`" + `aes-192-cbc` + "`" + `, ` + "`" + `aes-256-cbc` + "`" + `, ` + "`" + `aes-128-gcm` + "`" + `, ` + "`" + `aes-256-gcm` + "`" + `, and ` + "`" + `null` + "`" + `. Note that the "gcm" values are only available in PAN-OS 7.0+.`,
				},
				resource.Attribute{
					Name:        "dh_group",
					Description: `(Optional) The DH group value. Valid values should start with the string ` + "`" + `group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifetime_type",
					Description: `(Optional) The lifetime type. Valid values are ` + "`" + `seconds` + "`" + `, ` + "`" + `minutes` + "`" + `, ` + "`" + `hours` + "`" + ` (the default), or ` + "`" + `days` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifetime_value",
					Description: `(Optional, int) The lifetime value.`,
				},
				resource.Attribute{
					Name:        "lifesize_type",
					Description: `(Optional) The lifesize type. Valid values are ` + "`" + `kb` + "`" + `, ` + "`" + `mb` + "`" + `, ` + "`" + `gb` + "`" + `, or ` + "`" + `tb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifesize_value",
					Description: `(Optional, int) the lifesize value.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_ipsec_tunnel",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama IPSec tunnels.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"ipsec",
				"tunnel",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "tunnel_interface",
					Description: `(Required) The tunnel interface.`,
				},
				resource.Attribute{
					Name:        "anti_replay",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable Anti-Replay check on this tunnel.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to enable IPv6.`,
				},
				resource.Attribute{
					Name:        "copy_tos",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to copy IP TOS bits from inner packet to IPSec packet (not recommended).`,
				},
				resource.Attribute{
					Name:        "copy_flow_label",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to copy IPv6 flow label for 6in6 tunnel from inner packet to IPSec packet (not recommended).`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to disable this IPSec tunnel.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type. Valid values are ` + "`" + `auto-key` + "`" + ` (the default), ` + "`" + `manual-key` + "`" + `, or ` + "`" + `global-protect-satellite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ak_ike_gateway",
					Description: `(Optional) IKE gateway name.`,
				},
				resource.Attribute{
					Name:        "ak_ipsec_crypto_profile",
					Description: `(Optional) IPSec crypto profile name.`,
				},
				resource.Attribute{
					Name:        "mk_local_spi",
					Description: `(Optional) Outbound SPI, hex format.`,
				},
				resource.Attribute{
					Name:        "mk_remote_spi",
					Description: `(Optional) Inbound SPI, hex format.`,
				},
				resource.Attribute{
					Name:        "mk_local_address_ip",
					Description: `(Optional) Specify exact IP address if interface has multiple addresses.`,
				},
				resource.Attribute{
					Name:        "mk_local_address_floating_ip",
					Description: `(Optional) Floating IP address in HA Active-Active configuration.`,
				},
				resource.Attribute{
					Name:        "mk_protocol",
					Description: `(Optional) Manual key protocol. Valid valies are ` + "`" + `esp` + "`" + ` or ` + "`" + `ah` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mk_auth_type",
					Description: `(Optional) Authentication algorithm. Valid values are ` + "`" + `md5` + "`" + `, ` + "`" + `sha1` + "`" + `, ` + "`" + `sha256` + "`" + `, ` + "`" + `sha384` + "`" + `, ` + "`" + `sha512` + "`" + `, or ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mk_auth_key",
					Description: `(Optional) The auth key for the given auth type.`,
				},
				resource.Attribute{
					Name:        "mk_esp_encryption_type",
					Description: `(Optional) The encryption algorithm. Valid values are ` + "`" + `des` + "`" + `, ` + "`" + `3des` + "`" + `, ` + "`" + `aes-128-cbc` + "`" + `, ` + "`" + `aes-192-cbc` + "`" + `, ` + "`" + `aes-256-cbc` + "`" + `, or ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mk_esp_encryption_key",
					Description: `(Optional) The encryption key.`,
				},
				resource.Attribute{
					Name:        "gps_interface",
					Description: `(Optional) Interface to communicate with portal.`,
				},
				resource.Attribute{
					Name:        "gps_portal_address",
					Description: `(Optional) GlobalProtect portal address.`,
				},
				resource.Attribute{
					Name:        "gps_prefer_ipv6",
					Description: `(Optional, PAN-OS 8.0+, bool) Prefer to register the portal in IPv6. Only applicable to FQDN portal-address.`,
				},
				resource.Attribute{
					Name:        "gps_interface_ip_ipv4",
					Description: `(Optional) specify exact IP address if interface has multiple addresses (IPv4).`,
				},
				resource.Attribute{
					Name:        "gps_interface_ip_ipv6",
					Description: `(Optional, PAN-OS 8.0+) specify exact IP address if interface has multiple addresses (IPv6).`,
				},
				resource.Attribute{
					Name:        "gps_interface_floating_ip_ipv4",
					Description: `(Optional, PAN-OS 7.0+) Floating IPv4 address in HA Active-Active configuration.`,
				},
				resource.Attribute{
					Name:        "gps_interface_floating_ip_ipv6",
					Description: `(Optional, PAN-OS 8.0+) Floating IPv6 address in HA Active-Active configuration.`,
				},
				resource.Attribute{
					Name:        "gps_publish_connected_routes",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to to publish connected and static routes.`,
				},
				resource.Attribute{
					Name:        "gps_publish_routes",
					Description: `(Optional, list) Specify list of routes to publish to Global Protect Gateway.`,
				},
				resource.Attribute{
					Name:        "gps_local_certificate",
					Description: `(Optional) GlobalProtect satellite certificate file name.`,
				},
				resource.Attribute{
					Name:        "gps_certificate_profile",
					Description: `(Optional) Profile for authenticating GlobalProtect gateway certificates.`,
				},
				resource.Attribute{
					Name:        "enable_tunnel_monitor",
					Description: `(Optional, bool) Enable tunnel monitoring on this tunnel.`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_destination_ip",
					Description: `(Optional) Destination IP to send ICMP probe.`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_source_ip",
					Description: `(Optional) Source IP to send ICMP probe`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_profile",
					Description: `(Optional) Tunnel monitor profile.`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_proxy_id",
					Description: `(Optional, PAN-OS 7.0+) Which proxy-id (or proxy-id-v6) the monitoring traffic will use.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_ipsec_tunnel_proxy_id_ipv4",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama IPv4 proxy IDs for auto key IPSec tunnels.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"ipsec",
				"tunnel",
				"proxy",
				"id",
				"ipv4",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "ipsec_tunnel",
					Description: `(Required) The auto key IPSec tunnel to attach this proxy ID to.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) IP subnet or IP address represents local network.`,
				},
				resource.Attribute{
					Name:        "remote",
					Description: `(Optional) IP subnet or IP address represents remote network.`,
				},
				resource.Attribute{
					Name:        "protocol_any",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` for any IP protocol.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `(Optional, int) IP protocol number.`,
				},
				resource.Attribute{
					Name:        "protocol_tcp_local",
					Description: `(Optional, int) Local TCP port number.`,
				},
				resource.Attribute{
					Name:        "protocol_tcp_remote",
					Description: `(Optional, int) Remote TCP port number.`,
				},
				resource.Attribute{
					Name:        "protocol_udp_local",
					Description: `(Optional, int) Local UDP port number.`,
				},
				resource.Attribute{
					Name:        "protocol_udp_remote",
					Description: `(Optional, int) Remote UDP port number.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_loopback_interface",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama loopback interfaces.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"loopback",
				"interface",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name. This must start with ` + "`" + `loopback.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys that will use this interface (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional, bool) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv6 MSS adjust value.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_management_profile",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama interface management profiles.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"management",
				"profile",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The management profile's name.`,
				},
				resource.Attribute{
					Name:        "ping",
					Description: `(Optional) Allow ping.`,
				},
				resource.Attribute{
					Name:        "telnet",
					Description: `(Optional) Allow telnet.`,
				},
				resource.Attribute{
					Name:        "ssh",
					Description: `(Optional) Allow SSH.`,
				},
				resource.Attribute{
					Name:        "http",
					Description: `(Optional) Allow HTTP.`,
				},
				resource.Attribute{
					Name:        "http_ocsp",
					Description: `(Optional) Allow HTTP OCSP.`,
				},
				resource.Attribute{
					Name:        "https",
					Description: `(Optional) Allow HTTPS.`,
				},
				resource.Attribute{
					Name:        "snmp",
					Description: `(Optional) Allow SNMP.`,
				},
				resource.Attribute{
					Name:        "response_pages",
					Description: `(Optional) Allow response pages.`,
				},
				resource.Attribute{
					Name:        "userid_service",
					Description: `(Optional) Allow User ID service.`,
				},
				resource.Attribute{
					Name:        "userid_syslog_listener_ssl",
					Description: `(Optional) Allow User ID syslog listener for SSL.`,
				},
				resource.Attribute{
					Name:        "userid_syslog_listener_udp",
					Description: `(Optional) Allow User ID syslog listener for UDP.`,
				},
				resource.Attribute{
					Name:        "permitted_ips",
					Description: `(Optional) The list of permitted IP addresses or address ranges for this management profile.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_nat_rule",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama NAT rules.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"nat",
				"rule",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The NAT rule's name.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group to put the NAT rule into (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rulebase",
					Description: `(Optional) The rulebase. This can be ` + "`" + `pre-rulebase` + "`" + ` (default), ` + "`" + `post-rulebase` + "`" + `, or ` + "`" + `rulebase` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional). NAT type. This can be ` + "`" + `ipv4` + "`" + ` (default), ` + "`" + `nat64` + "`" + `, or ` + "`" + `nptv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `(Required) The list of source zone(s).`,
				},
				resource.Attribute{
					Name:        "destination_zone",
					Description: `(Required) The destination zone.`,
				},
				resource.Attribute{
					Name:        "to_interface",
					Description: `(Optional) Egress interface from route lookup (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Service (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Required) List of source address(es).`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Required) List of destination address(es).`,
				},
				resource.Attribute{
					Name:        "sat_type",
					Description: `(Optional) Type of source address translation. This can be ` + "`" + `none` + "`" + ` (default), ` + "`" + `dynamic-ip-and-port` + "`" + `, ` + "`" + `dynamic-ip` + "`" + `, or ` + "`" + `static-ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_address_type",
					Description: `(Optional) Source address translation address type. This can be ` + "`" + `interface-address` + "`" + ` or ` + "`" + `translated-address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_translated_addresses",
					Description: `(Optional) Source address translation list of translated addresses.`,
				},
				resource.Attribute{
					Name:        "sat_interface",
					Description: `(Optional) Source address translation interface.`,
				},
				resource.Attribute{
					Name:        "sat_ip_address",
					Description: `(Optional) Source address translation IP address.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_type",
					Description: `(Optional) Source address translation fallback type. This can be ` + "`" + `none` + "`" + `, ` + "`" + `interface-address` + "`" + `, or ` + "`" + `translated-address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_translated_addresses",
					Description: `(Optional) Source address translation list of fallback translated addresses.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_interface",
					Description: `(Optional) Source address translation fallback interface.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_ip_type",
					Description: `(Optional) Source address translation fallback IP type. This can be ` + "`" + `ip` + "`" + ` or ` + "`" + `floating` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_ip_address",
					Description: `(Optional) The source address translation fallback IP address.`,
				},
				resource.Attribute{
					Name:        "sat_static_translated_address",
					Description: `(Optional) The statically translated source address.`,
				},
				resource.Attribute{
					Name:        "sat_static_bi_directional",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable bi-directional source address translation.`,
				},
				resource.Attribute{
					Name:        "dat_type",
					Description: `(Optional) Destination address translation type. This should be either ` + "`" + `static` + "`" + ` or ` + "`" + `dynamic` + "`" + `. The ` + "`" + `dynamic` + "`" + ` option is only available on PAN-OS 8.1+.`,
				},
				resource.Attribute{
					Name:        "dat_address",
					Description: `(Optional) Destination address translation's address. Requires ` + "`" + `dat_type` + "`" + ` be set to "static" or "dynamic".`,
				},
				resource.Attribute{
					Name:        "dat_port",
					Description: `(Optional) Destination address translation's port number. Requires ` + "`" + `dat_type` + "`" + ` be set to "static" or "dynamic".`,
				},
				resource.Attribute{
					Name:        "dat_dynamic_distribution",
					Description: `(Optional, PAN-OS 8.1+) Distribution algorithm for destination address pool. The PAN-OS 8.1 GUI doesn't seem to set this anywhere, but this is added here for completeness' sake. Requires ` + "`" + `dat_type` + "`" + ` of "dynamic".`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to disable this rule.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) A target definition (see below). If there are no target sections, then the rule will apply to every vsys of every device in the device group.`,
				},
				resource.Attribute{
					Name:        "negate_target",
					Description: `(Optional, bool) Instead of applying the rule for the given serial numbers, apply it to everything except them. The following arguments are valid for each ` + "`" + `target` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `(Optional) A subset of all available vsys on the firewall that should be in this device group. If the firewall is a virtual firewall, then this parameter should just be omitted.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_nat_rule_group",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a group of Panorama NAT rules.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"nat",
				"rule",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the NAT rule group into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) Device group the NAT rules should be put into (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rulebase",
					Description: `(Optional) The rulebase the NAT rules should be put into. Valid values are ` + "`" + `pre-rulebase` + "`" + ` (default), ` + "`" + `rulebase` + "`" + `, or ` + "`" + `post-rulebase` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `(Optional) A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `(Optional) Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Repeatable) The rule definition (see below). The rule ordering will match how they appear in the terraform plan file. Each ` + "`" + `rule` + "`" + ` defined supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The NAT rule's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional). NAT type. This can be ` + "`" + `ipv4` + "`" + ` (default), ` + "`" + `nat64` + "`" + `, or ` + "`" + `nptv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to disable this rule.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional, repeatable) A target definition (see below). If there are no target sections, then the rule will apply to every vsys of every device in the device group.`,
				},
				resource.Attribute{
					Name:        "negate_target",
					Description: `(Optional, bool) Instead of applying the rule for the given serial numbers, apply it to everything except them.`,
				},
				resource.Attribute{
					Name:        "original_packet",
					Description: `(Required) The original packet specification (see below).`,
				},
				resource.Attribute{
					Name:        "translated_packet",
					Description: `(Required) The translated packet spec (see below). ` + "`" + `target` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `(Optional) A subset of all available vsys on the firewall that should be in this device group. If the firewall is a virtual firewall, then this parameter should just be omitted. ` + "`" + `original_packet` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `(Required) The list of source zone(s).`,
				},
				resource.Attribute{
					Name:        "destination_zone",
					Description: `(Required) The destination zone.`,
				},
				resource.Attribute{
					Name:        "destination_interface",
					Description: `(Optional) Egress interface from route lookup (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Service (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Required) List of source address(es).`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Required) List of destination address(es). ` + "`" + `translated_packet` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The source spec (see below). Leave this empty for a destination NAT of "none".`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) The destination spec (see below). Leave this empty for a destination NAT of "none". ` + "`" + `translated_packet.source` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "dynamic_ip_and_port",
					Description: `(Optional) Dynamic IP and port source translation spec (see below).`,
				},
				resource.Attribute{
					Name:        "dynamic_ip",
					Description: `(Optional) Dynamic IP source translation spec (see below).`,
				},
				resource.Attribute{
					Name:        "static_ip",
					Description: `(Optional) Static IP source translation spec (see below). ` + "`" + `translated_packet.source.dynamic_ip_and_port` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `(Optional) Translated address source translation type spec (see below).`,
				},
				resource.Attribute{
					Name:        "interface_address",
					Description: `(Optional) Interface address source translation type spec (see below). ` + "`" + `translated_packet.source.dynamic_ip_and_port.translated_address` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `(Required) List of translated addresses. ` + "`" + `translated_packet.source.dynamic_ip_and_port.interface_address` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address. ` + "`" + `translated_packet.source.dynamic_ip` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `(Optional) The list of translated addresses.`,
				},
				resource.Attribute{
					Name:        "fallback",
					Description: `(Optional) The fallback spec (see below). Leaving this empty or omiting it means a fallback of "None". ` + "`" + `translated_packet.source.dynamic_ip.fallback` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `(Optional) The translated address fallback spec (see below).`,
				},
				resource.Attribute{
					Name:        "interface_address",
					Description: `(Optional) The interface address fallback spec (see below). ` + "`" + `translated_packet.source.dynamic_ip.fallback.translated_address` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `(Optional) List of source address translation fallback translated addresses. ` + "`" + `translated_packet.source.dynamic_ip.fallback.interface_address` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) Source address translation fallback interface.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of interface fallback. Valid values are ` + "`" + `ip` + "`" + ` (default) or ` + "`" + `floating` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP address of the fallback interface. ` + "`" + `translated_packet.source.static_ip` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `(Required) The statically translated source address.`,
				},
				resource.Attribute{
					Name:        "bi_directional",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable bi-directional source address translation. ` + "`" + `translated_packet.destination` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "static",
					Description: `(Optional) Specifies a static destination NAT (see below).`,
				},
				resource.Attribute{
					Name:        "dynamic",
					Description: `(Optional, PAN-OS 8.1+) Specify a dynamic destination NAT (see below). ` + "`" + `translated_packet.destination.static` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Destination address translation address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, int) Destination address translation port number. ` + "`" + `translated_packet.destination.dynamic` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Destination address translation address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, int) Destination address translation port number.`,
				},
				resource.Attribute{
					Name:        "distribution",
					Description: `(Optional, PAN-OS 8.1+) Distribution algorithm for destination address pool. The PAN-OS 8.1 GUI doesn't seem to set this anywhere, but this is added here for completeness' sake. The GUI sets this to ` + "`" + `round-robin` + "`" + ` currently.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_redistribution_profile_ipv4",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama redistribution profiles.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"redistribution",
				"profile",
				"ipv4",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The redistribution profile's name.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add the redistribution profile to.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required, int) The priority, integer from 1 to 255.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The action. Valid values are ` + "`" + `redist` + "`" + ` (default) or ` + "`" + `no-redist` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `(Optional) The source types. Valid values are ` + "`" + `bgp` + "`" + `, ` + "`" + `connect` + "`" + `, ` + "`" + `ospf` + "`" + `, ` + "`" + `rip` + "`" + `, and ` + "`" + `static` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Specify candidate routes.`,
				},
				resource.Attribute{
					Name:        "destinations",
					Description: `(Optional) Specify candidate routes' next-hop addresses (subnet match).`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) Specify candidate routes' next-hop addresses (subnet match).`,
				},
				resource.Attribute{
					Name:        "ospf_path_types",
					Description: `(Optional) OSPF path types. Valid values are ` + "`" + `intra-area` + "`" + `, ` + "`" + `inter-area` + "`" + `, ` + "`" + `ext-1` + "`" + `, and ` + "`" + `ext-2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ospf_areas",
					Description: `(Optional) OSPF areas.`,
				},
				resource.Attribute{
					Name:        "ospf_tags",
					Description: `(Optional) OSPF tags.`,
				},
				resource.Attribute{
					Name:        "bgp_communities",
					Description: `(Optional) BGP communities.`,
				},
				resource.Attribute{
					Name:        "bgp_extended_communities",
					Description: `(Optional) BGP extended communities.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_security_policy",
			Category:         "Panorama Resources",
			ShortDescription: `Manages the full Panorama security policy.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"security",
				"policy",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group to put the security policy into (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rulebase",
					Description: `(Optional) The rulebase. This can be ` + "`" + `pre-rulebase` + "`" + ` (default), ` + "`" + `post-rulebase` + "`" + `, or ` + "`" + `rulebase` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The security rule definition (see below). The security rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The security rule name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Rule type. This can be ` + "`" + `universal` + "`" + ` (default), ` + "`" + `interzone` + "`" + `, or ` + "`" + `intrazone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags for this security rule.`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `(Required) List of source zones.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Required) List of source addresses.`,
				},
				resource.Attribute{
					Name:        "negate_source",
					Description: `(Optional, bool) If the source should be negated.`,
				},
				resource.Attribute{
					Name:        "source_users",
					Description: `(Required) List of source users.`,
				},
				resource.Attribute{
					Name:        "hip_profiles",
					Description: `(Required) List of HIP profiles.`,
				},
				resource.Attribute{
					Name:        "destination_zones",
					Description: `(Required) List of destination zones.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Required) List of destination addresses.`,
				},
				resource.Attribute{
					Name:        "negate_destination",
					Description: `(Optional, bool) If the destination should be negated.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(Required) List of applications.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) List of services.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `(Required) List of categories.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action for the matched traffic. This can be ` + "`" + `allow` + "`" + ` (default), ` + "`" + `deny` + "`" + `, ` + "`" + `drop` + "`" + `, ` + "`" + `reset-client` + "`" + `, ` + "`" + `reset-server` + "`" + `, or ` + "`" + `reset-both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_setting",
					Description: `(Optional) Log forwarding profile.`,
				},
				resource.Attribute{
					Name:        "log_start",
					Description: `(Optional, bool) Log the start of the traffic flow.`,
				},
				resource.Attribute{
					Name:        "log_end",
					Description: `(Optional, bool) Log the end of the traffic flow (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to disable this rule.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Optional) The security rule schedule.`,
				},
				resource.Attribute{
					Name:        "icmp_unreachable",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable ICMP unreachable.`,
				},
				resource.Attribute{
					Name:        "disable_server_response_inspection",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to disable server response inspection.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Profile Setting: ` + "`" + `Group` + "`" + ` - The group profile name.`,
				},
				resource.Attribute{
					Name:        "virus",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The antivirus setting.`,
				},
				resource.Attribute{
					Name:        "spyware",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The anti-spyware setting.`,
				},
				resource.Attribute{
					Name:        "vulnerability",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The Vulnerability Protection setting.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The URL filtering setting.`,
				},
				resource.Attribute{
					Name:        "file_blocking",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The file blocking setting.`,
				},
				resource.Attribute{
					Name:        "wildfire_analysis",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The WildFire Analysis setting.`,
				},
				resource.Attribute{
					Name:        "data_filtering",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The Data Filtering setting.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) A target definition (see below). If there are no target sections, then the rule will apply to every vsys of every device in the device group.`,
				},
				resource.Attribute{
					Name:        "negate_target",
					Description: `(Optional, bool) Instead of applying the rule for the given serial numbers, apply it to everything except them. The following arguments are valid for each ` + "`" + `target` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `(Optional) A subset of all available vsys on the firewall that should be in this device group. If the firewall is a virtual firewall, then this parameter should just be omitted.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_security_rule_group",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama security rule groups.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"security",
				"rule",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group to put the security rules into (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rulebase",
					Description: `(Optional) The rulebase. This can be ` + "`" + `pre-rulebase` + "`" + ` (default), ` + "`" + `post-rulebase` + "`" + `, or ` + "`" + `rulebase` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `(Optional) A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `(Optional) Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The security rule definition (see below). The security rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The security rule name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Rule type. This can be ` + "`" + `universal` + "`" + ` (default), ` + "`" + `interzone` + "`" + `, or ` + "`" + `intrazone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags for this security rule.`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `(Required) List of source zones.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Required) List of source addresses.`,
				},
				resource.Attribute{
					Name:        "negate_source",
					Description: `(Optional, bool) If the source should be negated.`,
				},
				resource.Attribute{
					Name:        "source_users",
					Description: `(Required) List of source users.`,
				},
				resource.Attribute{
					Name:        "hip_profiles",
					Description: `(Required) List of HIP profiles.`,
				},
				resource.Attribute{
					Name:        "destination_zones",
					Description: `(Required) List of destination zones.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Required) List of destination addresses.`,
				},
				resource.Attribute{
					Name:        "negate_destination",
					Description: `(Optional, bool) If the destination should be negated.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(Required) List of applications.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) List of services.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `(Required) List of categories.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action for the matched traffic. This can be ` + "`" + `allow` + "`" + ` (default), ` + "`" + `deny` + "`" + `, ` + "`" + `drop` + "`" + `, ` + "`" + `reset-client` + "`" + `, ` + "`" + `reset-server` + "`" + `, or ` + "`" + `reset-both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_setting",
					Description: `(Optional) Log forwarding profile.`,
				},
				resource.Attribute{
					Name:        "log_start",
					Description: `(Optional, bool) Log the start of the traffic flow.`,
				},
				resource.Attribute{
					Name:        "log_end",
					Description: `(Optional, bool) Log the end of the traffic flow (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to disable this rule.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Optional) The security rule schedule.`,
				},
				resource.Attribute{
					Name:        "icmp_unreachable",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable ICMP unreachable.`,
				},
				resource.Attribute{
					Name:        "disable_server_response_inspection",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to disable server response inspection.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Profile Setting: ` + "`" + `Group` + "`" + ` - The group profile name.`,
				},
				resource.Attribute{
					Name:        "virus",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The antivirus setting.`,
				},
				resource.Attribute{
					Name:        "spyware",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The anti-spyware setting.`,
				},
				resource.Attribute{
					Name:        "vulnerability",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The Vulnerability Protection setting.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The URL filtering setting.`,
				},
				resource.Attribute{
					Name:        "file_blocking",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The file blocking setting.`,
				},
				resource.Attribute{
					Name:        "wildfire_analysis",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The WildFire Analysis setting.`,
				},
				resource.Attribute{
					Name:        "data_filtering",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The Data Filtering setting.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) A target definition (see below). If there are no target sections, then the rule will apply to every vsys of every device in the device group.`,
				},
				resource.Attribute{
					Name:        "negate_target",
					Description: `(Optional, bool) Instead of applying the rule for the given serial numbers, apply it to everything except them. The following arguments are valid for each ` + "`" + `target` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `(Optional) A subset of all available vsys on the firewall that should be in this device group. If the firewall is a virtual firewall, then this parameter should just be omitted.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_service_group",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama service groups.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"service",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The service group's name.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group to put the service group into (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) List of services to put in this service group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_service_object",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama service objects.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"service",
				"object",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The service object's name.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group to put the service object into (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The service object's description.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The service's protocol. This should be ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) The source port. This can be a single port number, range (1-65535), or comma separated (80,8080,443).`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Required) The destination port. This can be a single port number, range (1-65535), or comma separated (80,8080,443).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_static_route_ipv4",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama IPv4 static routes.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"static",
				"route",
				"ipv4",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The static route's name.`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add the static route to.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Destination IP address / prefix.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) Interface to use.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The next hop type. Valid values are ` + "`" + `ip-address` + "`" + ` (the default), ` + "`" + `discard` + "`" + `, ` + "`" + `next-vr` + "`" + `, or an empty string for ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) The value for the ` + "`" + `type` + "`" + ` setting.`,
				},
				resource.Attribute{
					Name:        "admin_distance",
					Description: `(Optional) The admin distance.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional, int) Metric value / path cost (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional) Target routing table to install the route. Valid values are ` + "`" + `unicast` + "`" + ` (the default), ` + "`" + `no install` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `(Optional, PAN-OS 7.1+) BFD configuration.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_template",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama templates.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"template",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The template's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The template's description.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `The device definition (see below). The following arguments are valid for each ` + "`" + `device` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `(Optional) A subset of all available vsys on the firewall that should be in this template. If the firewall is a virtual firewall, then this parameter should just be omitted.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_template_entry",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a specific device in a Panorama template.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"template",
				"entry",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `(Optional) A subset of all available vsys on the firewall that should be in this template. If the firewall is a virtual firewall, then this parameter should just be omitted.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_template_stack",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama template stacks.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"template",
				"stack",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The stack's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The stack's description.`,
				},
				resource.Attribute{
					Name:        "default_vsys",
					Description: `(Optional) The default virtual system template configuration pushed to firewalls with a single virtual system.`,
				},
				resource.Attribute{
					Name:        "templates",
					Description: `(Optional) List of templates in this stack.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `(Optional) List of serial numbers to include in this stack.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_template_stack_entry",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a specific device in a Panorama template stack.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"template",
				"stack",
				"entry",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template_stack",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Required) The serial number of the device to add.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_template_variable",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama template variables.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"template",
				"variable",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The template's name. This must start with a dollar sign ($).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The variable type. Valid values are ` + "`" + `ip-netmask` + "`" + ` (default), ` + "`" + `ip-range` + "`" + `, ` + "`" + `fqdn` + "`" + `, ` + "`" + `group-id` + "`" + `, or ` + "`" + `interface` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The variable value.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_tunnel_interface",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama tunnel interfaces.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"tunnel",
				"interface",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name. This must start with ` + "`" + `tunnel.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys that will use this interface (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_virtual_router",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama virtual routers.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"virtual",
				"router",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The virtual router's name.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Required) The vsys that will use this virtual router. This should be something like ` + "`" + `vsys1` + "`" + ` or ` + "`" + `vsys3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) List of interfaces that should use this virtual router.`,
				},
				resource.Attribute{
					Name:        "static_dist",
					Description: `(Optional) Admin distance - Static (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_ipv6_dist",
					Description: `(Optional) Admin distance - Static IPv6 (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ospf_int_dist",
					Description: `(Optional) Admin distance - OSPF Int (default: ` + "`" + `30` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ospf_ext_dist",
					Description: `(Optional) Admin distance - OSPF Ext (default: ` + "`" + `110` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ospfv3_int_dist",
					Description: `(Optional) Admin distance - OSPFv3 Int (default: ` + "`" + `30` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ospfv3_ext_dist",
					Description: `(Optional) Admin distance - OSPFv3 Ext (default: ` + "`" + `110` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ibgp_dist",
					Description: `(Optional) Admin distance - IBGP (default: ` + "`" + `200` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ebgp_dist",
					Description: `(Optional) Admin distance - EBGP (default: ` + "`" + `20` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rip_dist",
					Description: `(Optional) Admin distance - RIP (default: ` + "`" + `120` + "`" + `).`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_virtual_router_entry",
			Category:         "Panorama Resources",
			ShortDescription: `Manages an interface in a Panorama virtual router.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"virtual",
				"router",
				"entry",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router's name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface to import into the virtual router.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_vlan_interface",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama VLAN interfaces.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"vlan",
				"interface",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name. Must start with ` + "`" + `vlan.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys that will use this interface (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable DHCP on this interface.`,
				},
				resource.Attribute{
					Name:        "create_dhcp_default_route",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to create a DHCP default route.`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route_metric",
					Description: `(Optional) The metric for the DHCP default route.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv6 MSS adjust value.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_zone",
			Category:         "Panorama Resources",
			ShortDescription: `Manages Panorama Zone objects.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"zone",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The zone's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the zone into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The zone's mode. This can be ` + "`" + `layer3` + "`" + `, ` + "`" + `layer2` + "`" + `, ` + "`" + `virtual-wire` + "`" + `, ` + "`" + `tap` + "`" + `, or ` + "`" + `tunnel` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_profile",
					Description: `(Optional) The zone protection profile.`,
				},
				resource.Attribute{
					Name:        "log_setting",
					Description: `(Optional) Log setting.`,
				},
				resource.Attribute{
					Name:        "enable_user_id",
					Description: `(Optional) Boolean to enable user identification.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) List of interfaces to associated with this zone.`,
				},
				resource.Attribute{
					Name:        "include_acls",
					Description: `(Optional) Users from these addresses/subnets will be identified. This can be an address object, an address group, a single IP address, or an IP address subnet.`,
				},
				resource.Attribute{
					Name:        "exclude_acls",
					Description: `(Optional) Users from these addresses/subnets will not be identified. This can be an address object, an address group, a single IP address, or an IP address subnet.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_zone_entry",
			Category:         "Panorama Resources",
			ShortDescription: `Manages a specific interface in a Panorama zone.`,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"zone",
				"entry",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The zone's name.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The mode. Can be ` + "`" + `layer3` + "`" + ` (default), ` + "`" + `layer2` + "`" + `, ` + "`" + `virtual-wire` + "`" + `, ` + "`" + `tap` + "`" + `, or ` + "`" + `external` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface's name.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_redistribution_profile_ipv4",
			Category:         "Firewall Resources",
			ShortDescription: `Manages redistribution profiles.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"redistribution",
				"profile",
				"ipv4",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The redistribution profile's name.`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add the redistribution profile to.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required, int) The priority, integer from 1 to 255.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The action. Valid values are ` + "`" + `redist` + "`" + ` (default) or ` + "`" + `no-redist` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `(Optional) The source types. Valid values are ` + "`" + `bgp` + "`" + `, ` + "`" + `connect` + "`" + `, ` + "`" + `ospf` + "`" + `, ` + "`" + `rip` + "`" + `, and ` + "`" + `static` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Specify candidate routes.`,
				},
				resource.Attribute{
					Name:        "destinations",
					Description: `(Optional) Specify candidate routes' next-hop addresses (subnet match).`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) Specify candidate routes' next-hop addresses (subnet match).`,
				},
				resource.Attribute{
					Name:        "ospf_path_types",
					Description: `(Optional) OSPF path types. Valid values are ` + "`" + `intra-area` + "`" + `, ` + "`" + `inter-area` + "`" + `, ` + "`" + `ext-1` + "`" + `, and ` + "`" + `ext-2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ospf_areas",
					Description: `(Optional) OSPF areas.`,
				},
				resource.Attribute{
					Name:        "ospf_tags",
					Description: `(Optional) OSPF tags.`,
				},
				resource.Attribute{
					Name:        "bgp_communities",
					Description: `(Optional) BGP communities.`,
				},
				resource.Attribute{
					Name:        "bgp_extended_communities",
					Description: `(Optional) BGP extended communities.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_security_policy",
			Category:         "Firewall Resources",
			ShortDescription: `Manages the full security policy.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"security",
				"policy",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the security policy into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rulebase",
					Description: `(Optional, Deprecated) The rulebase. For firewalls, there is only the ` + "`" + `rulebase` + "`" + ` value (default), but on Panorama, there is also ` + "`" + `pre-rulebase` + "`" + ` and ` + "`" + `post-rulebase` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `A security rule definition (see below). The security rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The security rule name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Rule type. This can be ` + "`" + `universal` + "`" + ` (default), ` + "`" + `interzone` + "`" + `, or ` + "`" + `intrazone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags for this security rule.`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `(Required) List of source zones.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Required) List of source addresses.`,
				},
				resource.Attribute{
					Name:        "negate_source",
					Description: `(Optional, bool) If the source should be negated.`,
				},
				resource.Attribute{
					Name:        "source_users",
					Description: `(Required) List of source users.`,
				},
				resource.Attribute{
					Name:        "hip_profiles",
					Description: `(Required) List of HIP profiles.`,
				},
				resource.Attribute{
					Name:        "destination_zones",
					Description: `(Required) List of destination zones.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Required) List of destination addresses.`,
				},
				resource.Attribute{
					Name:        "negate_destination",
					Description: `(Optional, bool) If the destination should be negated.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(Required) List of applications.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) List of services.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `(Required) List of categories.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action for the matched traffic. This can be ` + "`" + `allow` + "`" + ` (default), ` + "`" + `deny` + "`" + `, ` + "`" + `drop` + "`" + `, ` + "`" + `reset-client` + "`" + `, ` + "`" + `reset-server` + "`" + `, or ` + "`" + `reset-both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_setting",
					Description: `(Optional) Log forwarding profile.`,
				},
				resource.Attribute{
					Name:        "log_start",
					Description: `(Optional, bool) Log the start of the traffic flow.`,
				},
				resource.Attribute{
					Name:        "log_end",
					Description: `(Optional, bool) Log the end of the traffic flow (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to disable this rule.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Optional) The security policy schedule.`,
				},
				resource.Attribute{
					Name:        "icmp_unreachable",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable ICMP unreachable.`,
				},
				resource.Attribute{
					Name:        "disable_server_response_inspection",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to disable server response inspection.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Profile Setting: ` + "`" + `Group` + "`" + ` - The group profile name.`,
				},
				resource.Attribute{
					Name:        "virus",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The antivirus setting.`,
				},
				resource.Attribute{
					Name:        "spyware",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The anti-spyware setting.`,
				},
				resource.Attribute{
					Name:        "vulnerability",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The Vulnerability Protection setting.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The URL filtering setting.`,
				},
				resource.Attribute{
					Name:        "file_blocking",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The file blocking setting.`,
				},
				resource.Attribute{
					Name:        "wildfire_analysis",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The WildFire Analysis setting.`,
				},
				resource.Attribute{
					Name:        "data_filtering",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The Data Filtering setting.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_security_rule_group",
			Category:         "Firewall Resources",
			ShortDescription: `Manages security rule groups.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"security",
				"rule",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the security rule into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `(Optional) A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `(Optional) Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The security rule definition (see below). The security rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The security rule name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Rule type. This can be ` + "`" + `universal` + "`" + ` (default), ` + "`" + `interzone` + "`" + `, or ` + "`" + `intrazone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of tags for this security rule.`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `(Required) List of source zones.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Required) List of source addresses.`,
				},
				resource.Attribute{
					Name:        "negate_source",
					Description: `(Optional, bool) If the source should be negated.`,
				},
				resource.Attribute{
					Name:        "source_users",
					Description: `(Required) List of source users.`,
				},
				resource.Attribute{
					Name:        "hip_profiles",
					Description: `(Required) List of HIP profiles.`,
				},
				resource.Attribute{
					Name:        "destination_zones",
					Description: `(Required) List of destination zones.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Required) List of destination addresses.`,
				},
				resource.Attribute{
					Name:        "negate_destination",
					Description: `(Optional, bool) If the destination should be negated.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(Required) List of applications.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) List of services.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `(Required) List of categories.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action for the matched traffic. This can be ` + "`" + `allow` + "`" + ` (default), ` + "`" + `deny` + "`" + `, ` + "`" + `drop` + "`" + `, ` + "`" + `reset-client` + "`" + `, ` + "`" + `reset-server` + "`" + `, or ` + "`" + `reset-both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_setting",
					Description: `(Optional) Log forwarding profile.`,
				},
				resource.Attribute{
					Name:        "log_start",
					Description: `(Optional, bool) Log the start of the traffic flow.`,
				},
				resource.Attribute{
					Name:        "log_end",
					Description: `(Optional, bool) Log the end of the traffic flow (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to disable this rule.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Optional) The security rule schedule.`,
				},
				resource.Attribute{
					Name:        "icmp_unreachable",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable ICMP unreachable.`,
				},
				resource.Attribute{
					Name:        "disable_server_response_inspection",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to disable server response inspection.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Profile Setting: ` + "`" + `Group` + "`" + ` - The group profile name.`,
				},
				resource.Attribute{
					Name:        "virus",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The antivirus setting.`,
				},
				resource.Attribute{
					Name:        "spyware",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The anti-spyware setting.`,
				},
				resource.Attribute{
					Name:        "vulnerability",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The Vulnerability Protection setting.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The URL filtering setting.`,
				},
				resource.Attribute{
					Name:        "file_blocking",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The file blocking setting.`,
				},
				resource.Attribute{
					Name:        "wildfire_analysis",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The WildFire Analysis setting.`,
				},
				resource.Attribute{
					Name:        "data_filtering",
					Description: `(Optional) Profile Setting: ` + "`" + `Profiles` + "`" + ` - The Data Filtering setting.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_service_group",
			Category:         "Firewall Resources",
			ShortDescription: `Manages service groups.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"service",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The service group's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the service group into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) List of services to put in this service group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_service_object",
			Category:         "Firewall Resources",
			ShortDescription: `Manages service objects.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"service",
				"object",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The service object's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the service object into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The service object's description.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The service's protocol. This should be ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) The source port. This can be a single port number, range (1-65535), or comma separated (80,8080,443).`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Required) The destination port. This can be a single port number, range (1-65535), or comma separated (80,8080,443).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_static_route_ipv4",
			Category:         "Firewall Resources",
			ShortDescription: `Manages IPv4 static routes.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"static",
				"route",
				"ipv4",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The static route's name.`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add the static route to.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Destination IP address / prefix.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) Interface to use.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The next hop type. Valid values are ` + "`" + `ip-address` + "`" + ` (the default), ` + "`" + `discard` + "`" + `, ` + "`" + `next-vr` + "`" + `, or an empty string for ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) The value for the ` + "`" + `type` + "`" + ` setting.`,
				},
				resource.Attribute{
					Name:        "admin_distance",
					Description: `(Optional) The admin distance.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional, int) Metric value / path cost (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional) Target routing table to install the route. Valid values are ` + "`" + `unicast` + "`" + ` (the default), ` + "`" + `no install` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `(Optional, PAN-OS 7.1+) BFD configuration.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_telemetry",
			Category:         "Firewall Resources",
			ShortDescription: `Manages anonymous statistics sharing.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"telemetry",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "application_reports",
					Description: `(Bool, optional) Application reports.`,
				},
				resource.Attribute{
					Name:        "threat_prevention_reports",
					Description: `(Bool, optional) Threat reports.`,
				},
				resource.Attribute{
					Name:        "url_reports",
					Description: `(Bool, optional) URL reports.`,
				},
				resource.Attribute{
					Name:        "file_type_identification_reports",
					Description: `(Bool, optional) File type identification reports.`,
				},
				resource.Attribute{
					Name:        "threat_prevention_data",
					Description: `(Bool, optional) Threat prevention data.`,
				},
				resource.Attribute{
					Name:        "threat_prevention_packet_captures",
					Description: `(Bool, optional) Enable sending packet- captures with threat prevention information. This requires that ` + "`" + `threat_prevention_data` + "`" + ` also be enabled.`,
				},
				resource.Attribute{
					Name:        "product_usage_stats",
					Description: `(Bool, optional) Health and performance reports.`,
				},
				resource.Attribute{
					Name:        "passive_dns_monitoring",
					Description: `(Bool, optional) Passive DNS monitoring.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_tunnel_interface",
			Category:         "Firewall Resources",
			ShortDescription: `Manages tunnel interfaces.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"tunnel",
				"interface",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name. This must start with ` + "`" + `tunnel.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys that will use this interface (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_virtual_router",
			Category:         "Firewall Resources",
			ShortDescription: `Manages virtual routers.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"virtual",
				"router",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The virtual router's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Required) The vsys that will use this virtual router. This should be something like ` + "`" + `vsys1` + "`" + ` or ` + "`" + `vsys3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) List of interfaces that should use this virtual router.`,
				},
				resource.Attribute{
					Name:        "static_dist",
					Description: `(Optional) Admin distance - Static (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_ipv6_dist",
					Description: `(Optional) Admin distance - Static IPv6 (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ospf_int_dist",
					Description: `(Optional) Admin distance - OSPF Int (default: ` + "`" + `30` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ospf_ext_dist",
					Description: `(Optional) Admin distance - OSPF Ext (default: ` + "`" + `110` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ospfv3_int_dist",
					Description: `(Optional) Admin distance - OSPFv3 Int (default: ` + "`" + `30` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ospfv3_ext_dist",
					Description: `(Optional) Admin distance - OSPFv3 Ext (default: ` + "`" + `110` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ibgp_dist",
					Description: `(Optional) Admin distance - IBGP (default: ` + "`" + `200` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ebgp_dist",
					Description: `(Optional) Admin distance - EBGP (default: ` + "`" + `20` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rip_dist",
					Description: `(Optional) Admin distance - RIP (default: ` + "`" + `120` + "`" + `).`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_virtual_router_entry",
			Category:         "Firewall Resources",
			ShortDescription: `Manages an interface in a virtual router.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"virtual",
				"router",
				"entry",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router's name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface to import into the virtual router.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_vlan_interface",
			Category:         "Firewall Resources",
			ShortDescription: `Manages vlan interfaces.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"vlan",
				"interface",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name. Must start with ` + "`" + `vlan.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys that will use this interface (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable DHCP on this interface.`,
				},
				resource.Attribute{
					Name:        "create_dhcp_default_route",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to create a DHCP default route.`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route_metric",
					Description: `(Optional) The metric for the DHCP default route.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv6 MSS adjust value.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_zone",
			Category:         "Firewall Resources",
			ShortDescription: `Manages Zone objects.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"zone",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The zone's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the zone into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The zone's mode. This can be ` + "`" + `layer3` + "`" + `, ` + "`" + `layer2` + "`" + `, ` + "`" + `virtual-wire` + "`" + `, ` + "`" + `tap` + "`" + `, or ` + "`" + `tunnel` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_profile",
					Description: `(Optional) The zone protection profile.`,
				},
				resource.Attribute{
					Name:        "log_setting",
					Description: `(Optional) Log setting.`,
				},
				resource.Attribute{
					Name:        "enable_user_id",
					Description: `(Optional) Boolean to enable user identification.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) List of interfaces to associated with this zone.`,
				},
				resource.Attribute{
					Name:        "include_acls",
					Description: `(Optional) Users from these addresses/subnets will be identified. This can be an address object, an address group, a single IP address, or an IP address subnet.`,
				},
				resource.Attribute{
					Name:        "exclude_acls",
					Description: `(Optional) Users from these addresses/subnets will not be identified. This can be an address object, an address group, a single IP address, or an IP address subnet.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_zone_entry",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a specific interface in a zone.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"zone",
				"entry",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The zone's name.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The mode. Can be ` + "`" + `layer3` + "`" + ` (default), ` + "`" + `layer2` + "`" + `, ` + "`" + `virtual-wire` + "`" + `, ` + "`" + `tap` + "`" + `, or ` + "`" + `external` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface's name.`,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"panos_address_group":                                 0,
		"panos_address_object":                                1,
		"panos_administrative_tag":                            2,
		"panos_bfd_profile":                                   3,
		"panos_bgp":                                           4,
		"panos_bgp_aggregate":                                 5,
		"panos_bgp_aggregate_advertise_filter":                6,
		"panos_bgp_aggregate_suppress_filter":                 7,
		"panos_bgp_auth_profile":                              8,
		"panos_bgp_conditional_adv":                           9,
		"panos_bgp_conditional_adv_advertise_filter":          10,
		"panos_bgp_conditional_adv_non_exist_filter":          11,
		"panos_bgp_dampening_profile":                         12,
		"panos_bgp_export_rule_group":                         13,
		"panos_bgp_import_rule_group":                         14,
		"panos_bgp_peer":                                      15,
		"panos_bgp_peer_group":                                16,
		"panos_bgp_redist_rule":                               17,
		"panos_dag_tags":                                      18,
		"panos_edl":                                           19,
		"panos_ethernet_interface":                            20,
		"panos_general_settings":                              21,
		"panos_ike_crypto_profile":                            22,
		"panos_ike_gateway":                                   23,
		"panos_ipsec_crypto_profile":                          24,
		"panos_ipsec_tunnel":                                  25,
		"panos_ipsec_tunnel_proxy_id_ipv4":                    26,
		"panos_license_api_key":                               27,
		"panos_licensing":                                     28,
		"panos_loopback_interface":                            29,
		"panos_management_profile":                            30,
		"panos_nat_rule":                                      31,
		"panos_nat_rule_group":                                32,
		"panos_panorama_address_group":                        33,
		"panos_panorama_address_object":                       34,
		"panos_panorama_administrative_tag":                   35,
		"panos_panorama_bfd_profile":                          36,
		"panos_panorama_bgp":                                  37,
		"panos_panorama_bgp_aggregate":                        38,
		"panos_panorama_bgp_aggregate_advertise_filter":       39,
		"panos_panorama_bgp_aggregate_suppress_filter":        40,
		"panos_panorama_bgp_auth_profile":                     41,
		"panos_panorama_bgp_conditional_adv":                  42,
		"panos_panorama_bgp_conditional_adv_advertise_filter": 43,
		"panos_panorama_bgp_conditional_adv_non_exist_filter": 44,
		"panos_panorama_bgp_dampening_profile":                45,
		"panos_panorama_bgp_export_rule_group":                46,
		"panos_panorama_bgp_import_rule_group":                47,
		"panos_panorama_bgp_peer":                             48,
		"panos_panorama_bgp_peer_group":                       49,
		"panos_panorama_bgp_redist_rule":                      50,
		"panos_panorama_device_group":                         51,
		"panos_panorama_device_group_entry":                   52,
		"panos_panorama_edl":                                  53,
		"panos_panorama_ethernet_interface":                   54,
		"panos_panorama_ike_crypto_profile":                   55,
		"panos_panorama_ike_gateway":                          56,
		"panos_panorama_ipsec_crypto_profile":                 57,
		"panos_panorama_ipsec_tunnel":                         58,
		"panos_panorama_ipsec_tunnel_proxy_id_ipv4":           59,
		"panos_panorama_loopback_interface":                   60,
		"panos_panorama_management_profile":                   61,
		"panos_panorama_nat_rule":                             62,
		"panos_panorama_nat_rule_group":                       63,
		"panos_panorama_redistribution_profile_ipv4":          64,
		"panos_panorama_security_policy":                      65,
		"panos_panorama_security_rule_group":                  66,
		"panos_panorama_service_group":                        67,
		"panos_panorama_service_object":                       68,
		"panos_panorama_static_route_ipv4":                    69,
		"panos_panorama_template":                             70,
		"panos_panorama_template_entry":                       71,
		"panos_panorama_template_stack":                       72,
		"panos_panorama_template_stack_entry":                 73,
		"panos_panorama_template_variable":                    74,
		"panos_panorama_tunnel_interface":                     75,
		"panos_panorama_virtual_router":                       76,
		"panos_panorama_virtual_router_entry":                 77,
		"panos_panorama_vlan_interface":                       78,
		"panos_panorama_zone":                                 79,
		"panos_panorama_zone_entry":                           80,
		"panos_redistribution_profile_ipv4":                   81,
		"panos_security_policy":                               82,
		"panos_security_rule_group":                           83,
		"panos_service_group":                                 84,
		"panos_service_object":                                85,
		"panos_static_route_ipv4":                             86,
		"panos_telemetry":                                     87,
		"panos_tunnel_interface":                              88,
		"panos_virtual_router":                                89,
		"panos_virtual_router_entry":                          90,
		"panos_vlan_interface":                                91,
		"panos_zone":                                          92,
		"panos_zone_entry":                                    93,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
