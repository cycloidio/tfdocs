package nsxt

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nsxt_algorithm_type_ns_service",
			Category:         "Manager NS Services Resources",
			ShortDescription: `A resource that can be used to configure a networking and security service in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ns",
				"services",
				"algorithm",
				"type",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Required) a single destination port.`,
				},
				resource.Attribute{
					Name:        "source_ports",
					Description: `(Optional) Set of source ports/ranges.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Required) Algorithm one of "ORACLE_TNS", "FTP", "SUN_RPC_TCP", "SUN_RPC_UDP", "MS_RPC_TCP", "MS_RPC_UDP", "NBNS_BROADCAST", "NBDG_BROADCAST", "TFTP"`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this service. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS service.`,
				},
				resource.Attribute{
					Name:        "default_service",
					Description: `The default NSServices are created in the system by default. These NSServices can't be modified/deleted.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing Algorithm type NS service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_algorithm_type_ns_service.ns_service_alg UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the algorithm based networking and security service named ` + "`" + `ns_service_alg` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS service.`,
				},
				resource.Attribute{
					Name:        "default_service",
					Description: `The default NSServices are created in the system by default. These NSServices can't be modified/deleted.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing Algorithm type NS service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_algorithm_type_ns_service.ns_service_alg UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the algorithm based networking and security service named ` + "`" + `ns_service_alg` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_dhcp_relay_profile",
			Category:         "Manager Resources",
			ShortDescription: `A resource that can be used to configure a DHCP relay profile in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"dhcp",
				"relay",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this DHCP relay profile.`,
				},
				resource.Attribute{
					Name:        "server_addresses",
					Description: `(Required) IP addresses of the DHCP relay servers. Maximum allowed amount is 2. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the DHCP relay profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing DHCP Relay profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_dhcp_relay_profile.dr_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the DHCP relay profile named ` + "`" + `dr_profile` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the DHCP relay profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing DHCP Relay profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_dhcp_relay_profile.dr_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the DHCP relay profile named ` + "`" + `dr_profile` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_dhcp_relay_service",
			Category:         "Manager Resources",
			ShortDescription: `A resource that can be used to configure a DHCP relay service in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"dhcp",
				"relay",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this dhcp_relay_service.`,
				},
				resource.Attribute{
					Name:        "dhcp_relay_profile_id",
					Description: `(Required) DHCP relay profile referenced by the DHCP relay service. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the DHCP relay service.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing DHCP Relay service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_dhcp_relay_service.dr_service UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the DHCP relay service named ` + "`" + `dr_service` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the DHCP relay service.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing DHCP Relay service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_dhcp_relay_service.dr_service UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the DHCP relay service named ` + "`" + `dr_service` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_dhcp_server_ip_pool",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure IP Pool for logical DHCP server on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"dhcp",
				"server",
				"ip",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "logical_dhcp_server_id",
					Description: `(Required) DHCP server uuid. Changing this would force new pool to be created.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Gateway IP.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `(Required) IP Ranges to be used within this pool.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) IP address that indicates range start.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) IP address that indicates range end.`,
				},
				resource.Attribute{
					Name:        "lease_time",
					Description: `(Optional) Lease time in seconds. Default is 86400.`,
				},
				resource.Attribute{
					Name:        "error_threshold",
					Description: `(Optional) Error threshold in percent. Valid values are from 80 to 100, default is 100.`,
				},
				resource.Attribute{
					Name:        "warning_threshold",
					Description: `(Optional) Warning threshold in percent. Valid values are from 50 to 80, default is 80.`,
				},
				resource.Attribute{
					Name:        "dhcp_option_121",
					Description: `(Optional) DHCP classless static routes. If specified, overrides DHCP server settings.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Destination in cidr format.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required) IP address of next hop.`,
				},
				resource.Attribute{
					Name:        "dhcp_generic_option",
					Description: `(Optional) Generic DHCP options. If specified, overrides DHCP server settings.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(Required) DHCP option code. Valid values are from 0 to 255.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) List of DHCP option values.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this logical DHCP server. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the DHCP server IP pool.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing DHCP server IP Pool can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_dhcp_server_ip_pool.ip_pool DHCP_SERVER_UUID POOL_UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the IP pool named ` + "`" + `ip pool` + "`" + ` for dhcp server with nsx ID ` + "`" + `DHCP_SERVER_UUID` + "`" + ` and pool nsx id ` + "`" + `POOL_UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the DHCP server IP pool.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing DHCP server IP Pool can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_dhcp_server_ip_pool.ip_pool DHCP_SERVER_UUID POOL_UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the IP pool named ` + "`" + `ip pool` + "`" + ` for dhcp server with nsx ID ` + "`" + `DHCP_SERVER_UUID` + "`" + ` and pool nsx id ` + "`" + `POOL_UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_dhcp_server_profile",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure DHCP server profile on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"dhcp",
				"server",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `(Required) Edge cluster uuid.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_member_indexes",
					Description: `(Optional) Up to 2 edge nodes from the given cluster. If none is provided, the NSX will auto-select two edge-nodes from the given edge cluster. If user provides only one edge node, there will be no HA support.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this DHCP profile. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the DHCP server profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing DHCP profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_dhcp_server_profile.dhcp_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the DHCP server profile named ` + "`" + `dhcp_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the DHCP server profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing DHCP profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_dhcp_server_profile.dhcp_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the DHCP server profile named ` + "`" + `dhcp_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ether_type_ns_service",
			Category:         "Manager NS Services Resources",
			ShortDescription: `A resource that can be used to configure a layer 2 ethernet networking and security service in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ns",
				"services",
				"ether",
				"type",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description.`,
				},
				resource.Attribute{
					Name:        "ether_type",
					Description: `(Required) Type of the encapsulated protocol.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this service. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS service.`,
				},
				resource.Attribute{
					Name:        "default_service",
					Description: `The default NSServices are created in the system by default. These NSServices can't be modified/deleted.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing Ethernet type NS service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ether_type_ns_service.etns UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the ethernet type networking and security service named ` + "`" + `etns` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS service.`,
				},
				resource.Attribute{
					Name:        "default_service",
					Description: `The default NSServices are created in the system by default. These NSServices can't be modified/deleted.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing Ethernet type NS service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ether_type_ns_service.etns UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the ethernet type networking and security service named ` + "`" + `etns` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_firewall_section",
			Category:         "Manager Resources",
			ShortDescription: `A resource that can be used to configure a firewall section in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"firewall",
				"section",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this firewall section. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this firewall section.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this firewall section.`,
				},
				resource.Attribute{
					Name:        "applied_to",
					Description: `(Optional) List of objects where the rules in this section will be enforced. This will take precedence over rule level applied_to. [Supported target types: "LogicalPort", "LogicalSwitch", "NSGroup", "LogicalRouter"]`,
				},
				resource.Attribute{
					Name:        "section_type",
					Description: `(Required) Type of the rules which a section can contain. Either LAYER2 or LAYER3. Only homogeneous sections are supported.`,
				},
				resource.Attribute{
					Name:        "stateful",
					Description: `(Required) Stateful or Stateless nature of firewall section is enforced on all rules inside the section. Layer3 sections can be stateful or stateless. Layer2 sections can only be stateless.`,
				},
				resource.Attribute{
					Name:        "insert_before",
					Description: `(Optional) Firewall section id that should come immediately after this one. It is user responsibility to use this attribute in consistent manner (for example, if same value would be set in two separate sections, the outcome would depend on order of creation). Changing this attribute would force recreation of the firewall section.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) A list of rules to be applied in this section. each rule has the following arguments:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this rule. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action enforced on the packets which matches the firewall rule. [Allowed values: "ALLOW", "DROP", "REJECT"]`,
				},
				resource.Attribute{
					Name:        "applied_to",
					Description: `(Optional) List of objects where rule will be enforced. The section level field overrides this one. Null will be treated as any. [Supported target types: "LogicalPort", "LogicalSwitch", "NSGroup", "LogicalRouterPort"]`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Optional) List of the destinations. Null will be treated as any. [Allowed target types: "IPSet", "LogicalPort", "LogicalSwitch", "NSGroup", "MACSet" (depending on the section type)]`,
				},
				resource.Attribute{
					Name:        "destinations_excluded",
					Description: `(Optional) When this boolean flag is set to true, the rule destinations will be negated.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Rule direction in case of stateless firewall rules. This will only considered if section level parameter is set to stateless. Default to IN_OUT if not specified. [Allowed values: "IN", "OUT", "IN_OUT"]`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Flag to disable rule. Disabled will only be persisted but never provisioned/realized.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) Type of IP packet that should be matched while enforcing the rule. [allowed values: "IPV4", "IPV6", "IPV4_IPV6"]`,
				},
				resource.Attribute{
					Name:        "logged",
					Description: `(Optional) Flag to enable packet logging. Default is disabled.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) User notes specific to the rule.`,
				},
				resource.Attribute{
					Name:        "rule_tag",
					Description: `(Optional) User level field which will be printed in CLI and packet logs.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) List of the services. Null will be treated as any. [Allowed target types: "NSService", "NSServiceGroup"]`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) List of sources. Null will be treated as any. [Allowed target types: "IPSet", "LogicalPort", "LogicalSwitch", "NSGroup", "MACSet" (depending on the section type)]`,
				},
				resource.Attribute{
					Name:        "sources_excluded",
					Description: `(Optional) When this boolean flag is set to true, the rule sources will be negated. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the firewall section.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `A boolean flag which reflects whether a firewall section is default section or not. Each Layer 3 and Layer 2 section will have at least and at most one default section. ## Importing An existing Firewall section can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_firewall_section.firewall_sect UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the firewall section named ` + "`" + `firewall_sect` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the firewall section.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `A boolean flag which reflects whether a firewall section is default section or not. Each Layer 3 and Layer 2 section will have at least and at most one default section. ## Importing An existing Firewall section can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_firewall_section.firewall_sect UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the firewall section named ` + "`" + `firewall_sect` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_icmp_type_ns_service",
			Category:         "Manager NS Services Resources",
			ShortDescription: `A resource that can be used to configure an ICMP based networking and security service in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ns",
				"services",
				"icmp",
				"type",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Version of ICMP protocol ICMPv4 or ICMPv6.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `(Optional) ICMP message type.`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `(Optional) ICMP message code`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this service. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS service.`,
				},
				resource.Attribute{
					Name:        "default_service",
					Description: `The default NSServices are created in the system by default. These NSServices can't be modified/deleted.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing ICMP type NS Service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_icmp_type_ns_service.x id ` + "`" + `` + "`" + `` + "`" + ` The above service imports the ICMP type network and security service named ` + "`" + `x` + "`" + ` with the NSX id ` + "`" + `id` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS service.`,
				},
				resource.Attribute{
					Name:        "default_service",
					Description: `The default NSServices are created in the system by default. These NSServices can't be modified/deleted.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing ICMP type NS Service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_icmp_type_ns_service.x id ` + "`" + `` + "`" + `` + "`" + ` The above service imports the ICMP type network and security service named ` + "`" + `x` + "`" + ` with the NSX id ` + "`" + `id` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_igmp_type_ns_service",
			Category:         "Manager NS Services Resources",
			ShortDescription: `A resource that can be used to configure an IGMP based networking and security service in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ns",
				"services",
				"igmp",
				"type",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this service. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS service.`,
				},
				resource.Attribute{
					Name:        "default_service",
					Description: `The default NSServices are created in the system by default. These NSServices can't be modified/deleted.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing IGMP type NS Service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_igmp_type_ns_service.ns_service_igmp UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the IGMP based networking and security service named ` + "`" + `ns_service_igmp` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS service.`,
				},
				resource.Attribute{
					Name:        "default_service",
					Description: `The default NSServices are created in the system by default. These NSServices can't be modified/deleted.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing IGMP type NS Service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_igmp_type_ns_service.ns_service_igmp UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the IGMP based networking and security service named ` + "`" + `ns_service_igmp` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ip_block",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure IP block on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ip",
				"block",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) Represents network address and the prefix length which will be associated with a layer-2 broadcast domain.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this IP block. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP block.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing IP block can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ip_block.ip_block UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the IP block named ` + "`" + `ip_block` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP block.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing IP block can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ip_block.ip_block UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the IP block named ` + "`" + `ip_block` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ip_block_subnet",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure IP block subnet on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ip",
				"block",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "block_id",
					Description: `(Required) Block id for which the subnet is created.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Represents the size or number of IP addresses in the subnet.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this IP block subnet. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP block subnet.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "allocation_range",
					Description: `A collection of IPv4 IP ranges used for IP allocation.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Represents the size or number of IP addresses in the subnet. All subnets of the same block must have the same size, which must be a power of 2. ## Importing An existing IP block subnet can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ip_block_subnet.ip_block_subnet UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the IP block subnet named ` + "`" + `ip_block_subnet` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP block subnet.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "allocation_range",
					Description: `A collection of IPv4 IP ranges used for IP allocation.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Represents the size or number of IP addresses in the subnet. All subnets of the same block must have the same size, which must be a power of 2. ## Importing An existing IP block subnet can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ip_block_subnet.ip_block_subnet UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the IP block subnet named ` + "`" + `ip_block_subnet` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ip_discovery_switching_profile",
			Category:         "Manager Switching Profiles Resources",
			ShortDescription: `Provides a resource to configure IP discovery switching profile on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"switching",
				"profiles",
				"ip",
				"discovery",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this IP discovery switching profile.`,
				},
				resource.Attribute{
					Name:        "arp_snooping_enabled",
					Description: `(Optional) A boolean flag iIndicates whether ARP snooping is enabled.`,
				},
				resource.Attribute{
					Name:        "vm_tools_enabled",
					Description: `(Optional) A boolean flag iIndicates whether VM tools will be enabled. This option is only supported on ESX where vm-tools is installed.`,
				},
				resource.Attribute{
					Name:        "dhcp_snooping_enabled",
					Description: `(Optional) A boolean flag iIndicates whether DHCP snooping is enabled.`,
				},
				resource.Attribute{
					Name:        "arp_bindings_limit",
					Description: `(Optional) Limit for the amount of ARP bindings. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP discovery switching profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing IP discovery switching profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ip_discovery_switching_profile.ip_discovery_switching_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the IP discovery switching profile named ` + "`" + `ip_discovery_switching_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP discovery switching profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing IP discovery switching profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ip_discovery_switching_profile.ip_discovery_switching_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the IP discovery switching profile named ` + "`" + `ip_discovery_switching_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ip_pool",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure IP pool on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ip",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this IP pool.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional) Subnets can be IPv4 or IPv6 and they should not overlap. The maximum number will not exceed 5 subnets. Each subnet has the following arguments:`,
				},
				resource.Attribute{
					Name:        "allocation_ranges",
					Description: `(Required) A collection of IPv4 Pool Ranges`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) Network address and the prefix length which will be associated with a layer-2 broadcast domainIPv4 Pool Ranges`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `(Optional) A collection of up to 3 DNS servers for the subnet`,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `(Optional) The DNS suffix for the DNS server`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) The default gateway address on a layer-3 router ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP pool.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing IP pool can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ip_pool.ip_pool UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the IP pool named ` + "`" + `ip_pool` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP pool.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing IP pool can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ip_pool.ip_pool UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the IP pool named ` + "`" + `ip_pool` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ip_pool_allocation_ip_address",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to allocate an IP address from an IP pool on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ip",
				"pool",
				"allocation",
				"address",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_pool_id",
					Description: `(Required) Ip Pool ID from which the IP address will be allocated. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP pool allocation IP address (currently identical to ` + "`" + `allocation_ip` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "allocation_ip",
					Description: `Allocation IP address. ## Importing An existing IP pool allocation address can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ip_pool_allocation_ip_address.ip1 POOL-UUID/UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the IP pool allocation address named ` + "`" + `ip_pool` + "`" + ` with the nsx ID ` + "`" + `UUID` + "`" + `, from IP Pool with nsx ID ` + "`" + `POOL-UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP pool allocation IP address (currently identical to ` + "`" + `allocation_ip` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "allocation_ip",
					Description: `Allocation IP address. ## Importing An existing IP pool allocation address can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ip_pool_allocation_ip_address.ip1 POOL-UUID/UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the IP pool allocation address named ` + "`" + `ip_pool` + "`" + ` with the nsx ID ` + "`" + `UUID` + "`" + `, from IP Pool with nsx ID ` + "`" + `POOL-UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ip_protocol_ns_service",
			Category:         "Manager NS Services Resources",
			ShortDescription: `A resource that can be used to configure an IP protocol based networking and security service in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ns",
				"services",
				"ip",
				"protocol",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) IP protocol number (0-255)`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this service. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS service.`,
				},
				resource.Attribute{
					Name:        "default_service",
					Description: `The default NSServices are created in the system by default. These NSServices can't be modified/deleted.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing IP protocol NS service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ip_protocol_ns_service.ns_service_ip UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the IP protocol based networking and security service named ` + "`" + `ns_service_ip` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS service.`,
				},
				resource.Attribute{
					Name:        "default_service",
					Description: `The default NSServices are created in the system by default. These NSServices can't be modified/deleted.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing IP protocol NS service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ip_protocol_ns_service.ns_service_ip UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the IP protocol based networking and security service named ` + "`" + `ns_service_ip` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ip_set",
			Category:         "Manager Resources",
			ShortDescription: `A resource that can be used to configure an IP set in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ip",
				"set",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this IP set.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(Optional) IP addresses. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP set.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing IP set can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ip_set.ip_set1 UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the IP set named ` + "`" + `ip_set1` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP set.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing IP set can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ip_set.ip_set1 UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the IP set named ` + "`" + `ip_set1` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_l4_port_set_ns_service",
			Category:         "Manager NS Services Resources",
			ShortDescription: `A resource that can be used to configure a layer 4 networking and security service with ports in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ns",
				"services",
				"l4",
				"port",
				"set",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "destination_ports",
					Description: `(Optional) Set of destination ports.`,
				},
				resource.Attribute{
					Name:        "source_ports",
					Description: `(Optional) Set of source ports.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) L4 protocol. Accepted values - 'TCP' or 'UDP'.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this service. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS service.`,
				},
				resource.Attribute{
					Name:        "default_service",
					Description: `The default NSServices are created in the system by default. These NSServices can't be modified/deleted.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing L4 port set NS service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_l4_port_set_ns_service.ns_service_l4 UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the layer 4 port based networking and security service named ` + "`" + `ns_service_l4` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS service.`,
				},
				resource.Attribute{
					Name:        "default_service",
					Description: `The default NSServices are created in the system by default. These NSServices can't be modified/deleted.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing L4 port set NS service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_l4_port_set_ns_service.ns_service_l4 UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the layer 4 port based networking and security service named ` + "`" + `ns_service_l4` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_client_ssl_profile",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure lb client ssl profile on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"client",
				"ssl",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb client ssl profile.`,
				},
				resource.Attribute{
					Name:        "prefer_server_ciphers",
					Description: `(Optional) During SSL handshake as part of the SSL client Hello client sends an ordered list of ciphers that it can support (or prefers) and typically server selects the first one from the top of that list it can also support. For Perfect Forward Secrecy(PFS), server could override the client's preference. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "ciphers",
					Description: `(Optional) supported SSL cipher list to client side. The supported ciphers can contain: TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA, TLS_ECDH_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA, TLS_RSA_WITH_3DES_EDE_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384, TLS_RSA_WITH_AES_128_CBC_SHA256, TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256, TLS_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256, TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384, TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA256, TLS_ECDH_ECDSA_WITH_AES_128_GCM_SHA256, TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA384, TLS_ECDH_ECDSA_WITH_AES_256_GCM_SHA384, TLS_ECDH_RSA_WITH_AES_128_CBC_SHA, TLS_ECDH_RSA_WITH_AES_128_CBC_SHA256, TLS_ECDH_RSA_WITH_AES_128_GCM_SHA256, TLS_ECDH_RSA_WITH_AES_256_CBC_SHA384, TLS_ECDH_RSA_WITH_AES_256_GCM_SHA384.`,
				},
				resource.Attribute{
					Name:        "prefer_server_ciphers",
					Description: `(Optional) During SSL handshake as part of the SSL client Hello client sends an ordered list of ciphers that it can support (or prefers) and typically server selects the first one from the top of that list it can also support. For Perfect Forward Secrecy(PFS), server could override the client's preference. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "protocols",
					Description: `(Optional) SSL versions TLS_V1_1 and TLS_V1_2 are supported and enabled by default. SSL_V2, SSL_V3, and TLS_V1 are supported, but disabled by default.`,
				},
				resource.Attribute{
					Name:        "session_cache_enabled",
					Description: `(Optional) SSL session caching allows SSL client and server to reuse previously negotiated security parameters avoiding the expensive public key operation during handshake. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "session_cache_timeout",
					Description: `(Optional) Session cache timeout specifies how long the SSL session parameters are held on to and can be reused. Default value is 300. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb client ssl profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "is_secure",
					Description: `This flag is set to true when all the ciphers and protocols are secure. It is set to false when one of the ciphers or protocols is insecure. ## Importing An existing lb client ssl profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_client_ssl_profile.lb_client_ssl_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb client ssl profile named ` + "`" + `lb_client_ssl_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb client ssl profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "is_secure",
					Description: `This flag is set to true when all the ciphers and protocols are secure. It is set to false when one of the ciphers or protocols is insecure. ## Importing An existing lb client ssl profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_client_ssl_profile.lb_client_ssl_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb client ssl profile named ` + "`" + `lb_client_ssl_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_cookie_persistence_profile",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure lb cookie persistence profile on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"cookie",
				"persistence",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "cookie_mode",
					Description: `(Optional) The cookie persistence mode. Accepted values: PREFIX, REWRITE and INSERT which is the default.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `(Required) cookie name.`,
				},
				resource.Attribute{
					Name:        "persistence_shared",
					Description: `(Optional) A boolean flag which reflects whether the cookie persistence is private or shared. When false (which is the default value), the cookie persistence is private to each virtual server and is qualified by the pool. If set to true, in cookie insert mode, cookie persistence could be shared across multiple virtual servers that are bound to the same pools.`,
				},
				resource.Attribute{
					Name:        "cookie_fallback",
					Description: `(Optional) A boolean flag which reflects whether once the server points by this cookie is down, a new server is selected, or the requests will be rejected.`,
				},
				resource.Attribute{
					Name:        "cookie_garble",
					Description: `(Optional) A boolean flag which reflects whether the cookie value (server IP and port) would be encrypted or in plain text.`,
				},
				resource.Attribute{
					Name:        "insert_mode_params",
					Description: `(Optional) Additional parameters for the INSERT cookie mode:`,
				},
				resource.Attribute{
					Name:        "cookie_domain",
					Description: `(Optional) HTTP cookie domain (for INSERT mode only).`,
				},
				resource.Attribute{
					Name:        "cookie_path",
					Description: `(Optional) HTTP cookie path (for INSERT mode only).`,
				},
				resource.Attribute{
					Name:        "cookie_expiry_type",
					Description: `(Optional) Type of cookie expiration timing (for INSERT mode only). Accepted values: SESSION_COOKIE_TIME for session cookie time setting and PERSISTENCE_COOKIE_TIME for persistence cookie time setting.`,
				},
				resource.Attribute{
					Name:        "max_idle_time",
					Description: `(Required if cookie_expiry_type is set) Maximum interval the cookie is valid for from the last time it was seen in a request.`,
				},
				resource.Attribute{
					Name:        "max_life_time",
					Description: `(Required for INSERT mode with SESSION_COOKIE_TIME expiration) Maximum interval the cookie is valid for from the first time the cookie was seen in a request.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb cookie persistence profile. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb cookie persistence profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb cookie persistence profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_cookie_persistence_profile.lb_cookie_persistence_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb cookie persistence profile named ` + "`" + `lb_cookie_persistence_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb cookie persistence profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb cookie persistence profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_cookie_persistence_profile.lb_cookie_persistence_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb cookie persistence profile named ` + "`" + `lb_cookie_persistence_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_fast_tcp_application_profile",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure LB fast TCP application profile on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"fast",
				"tcp",
				"application",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "close_timeout",
					Description: `(Optional) Timeout in seconds to specify how long a closed TCP connection should be kept for this application before cleaning up the connection. Value can range between 1-60, with a default of 8 seconds.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional) Timeout in seconds to specify how long an idle TCP connection in ESTABLISHED state should be kept for this application before cleaning up. The default value will be 1800 seconds`,
				},
				resource.Attribute{
					Name:        "ha_flow_mirroring",
					Description: `(Optional) A boolean flag which reflects whether flow mirroring is enabled, and all the flows to the bounded virtual server are mirrored to the standby node. By default this is disabled.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb fast tcp profile. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb fast tcp profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb fast tcp profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_fast_tcp_application_profile.lb_fast_tcp_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the LB fast TCP application profile named ` + "`" + `lb_fast_tcp_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb fast tcp profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb fast tcp profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_fast_tcp_application_profile.lb_fast_tcp_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the LB fast TCP application profile named ` + "`" + `lb_fast_tcp_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_fast_udp_application_profile",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure LB fast UDP application profile on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"fast",
				"udp",
				"application",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional) Timeout in seconds to specify how long an idle UDP connection in ESTABLISHED state should be kept for this application before cleaning up. The default value will be 300 seconds`,
				},
				resource.Attribute{
					Name:        "ha_flow_mirroring",
					Description: `(Optional) A boolean flag which reflects whether flow mirroring is enabled, and all the flows to the bounded virtual server are mirrored to the standby node. By default this is disabled.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb fast udp profile. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb fast udp profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb fast udp profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_fast_udp_application_profile.lb_fast_udp_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the LB fast UDP application profile named ` + "`" + `lb_fast_udp_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb fast udp profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb fast udp profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_fast_udp_application_profile.lb_fast_udp_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the LB fast UDP application profile named ` + "`" + `lb_fast_udp_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_http_application_profile",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure LB HTTP application profile on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"http",
				"application",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "http_redirect_to",
					Description: `(Optional) A URL that incoming requests for that virtual server can be temporarily redirected to, If a website is temporarily down or has moved. When set, http_redirect_to_https should be false.`,
				},
				resource.Attribute{
					Name:        "http_redirect_to_https",
					Description: `(Optional) A boolean flag which reflects whether the client will automatically be redirected to use SSL. When true, the http_redirect_to should not be specified.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional) Timeout in seconds to specify how long an HTTP application can remain idle. Defaults to 15 seconds.`,
				},
				resource.Attribute{
					Name:        "ntlm",
					Description: `(Optional) A boolean flag which reflects whether NTLM challenge/response methodology will be used over HTTP. Can be set to true only if http_redirect_to_https is false.`,
				},
				resource.Attribute{
					Name:        "request_body_size",
					Description: `(Optional) Maximum request body size in bytes. If it is not specified, it means that request body size is unlimited.`,
				},
				resource.Attribute{
					Name:        "request_header_size",
					Description: `(Optional) Maximum request header size in bytes. Requests with larger header size will be processed as best effort whereas a request with header below this specified size is guaranteed to be processed. Defaults to 1024 bytes.`,
				},
				resource.Attribute{
					Name:        "response_timeout",
					Description: `(Optional) Number of seconds waiting for the server response before the connection is closed. Defaults to 60 seconds.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for",
					Description: `(Optional) When this value is set, the x_forwarded_for header in the incoming request will be inserted or replaced. Supported values are "INSERT" and "REPLACE".`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb http profile. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb http application profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb http profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_http_application_profile.lb_http_application_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the LB HTTP application profile named ` + "`" + `lb_http_application_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb http application profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb http profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_http_application_profile.lb_http_application_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the LB HTTP application profile named ` + "`" + `lb_http_application_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_http_forwarding_rule",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure lb rule on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"http",
				"forwarding",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb rule.`,
				},
				resource.Attribute{
					Name:        "match_strategy",
					Description: `(Required) Strategy to define how load balancer rule is considered a match when multiple match conditions are specified in one rule. If set to ALL, then load balancer rule is considered a match only if all the conditions match. If set to ANY, then load balancer rule is considered a match if any one of the conditions match.`,
				},
				resource.Attribute{
					Name:        "body_condition",
					Description: `(Optional) Set of match conditions used to match http request body:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value to look for in the body.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Defines how value field is used to match the body of HTTP requests. Accepted values are STARTS_WITH, ENDS_WITH, CONTAINS, EQUALS, REGEX.`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) If true, case is significant in the match. Default is true.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "header_condition",
					Description: `(Optional) Set of match conditions used to match http request header:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of HTTP header to match.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of HTTP header to match.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Defines how value field is used to match the header value of HTTP requests. Accepted values are STARTS_WITH, ENDS_WITH, CONTAINS, EQUALS, REGEX. Header name field does not support match types.`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) If true, case is significant in the match. Default is true.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "cookie_condition",
					Description: `(Optional) Set of match conditions used to match http request cookie:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of cookie to match.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of cookie to match.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Defines how value field is used to match the cookie. Accepted values are STARTS_WITH, ENDS_WITH, CONTAINS, EQUALS, REGEX.`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) If true, case is significant in the match. Default is true.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "method_condition",
					Description: `(Optional) Set of match conditions used to match http request method:`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) One of GET, HEAD, POST, PUT, OPTIONS.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "version_condition",
					Description: `(Optional) Match condition used to match http version of the request:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) One of HTTP_VERSION_1_0, HTTP_VERSION_1_1.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "ip_condition",
					Description: `(Optional) Set of match conditions used to match IP header values of HTTP request:`,
				},
				resource.Attribute{
					Name:        "source_address",
					Description: `(Required) The value source IP address to match.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "uri_condition",
					Description: `(Optional) Set of match conditions used to match http request URI:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Required) The value of URI to match.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Defines how value field is used to match the URI. Accepted values are STARTS_WITH, ENDS_WITH, CONTAINS, EQUALS, REGEX.`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) If true, case is significant in the match. Default is true.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "http_reject_action",
					Description: `(At least one action is required) Set of http reject actions to be executed when load balancer rule matches:`,
				},
				resource.Attribute{
					Name:        "reply_status",
					Description: `(Required) The HTTP reply status.`,
				},
				resource.Attribute{
					Name:        "reply_message",
					Description: `(Required) The HTTP reply message.`,
				},
				resource.Attribute{
					Name:        "http_redirect_action",
					Description: `(At least one action is required) Set of http redirect actions to be executed when load balancer rule matches:`,
				},
				resource.Attribute{
					Name:        "redirect_status",
					Description: `(Required) The HTTP reply status.`,
				},
				resource.Attribute{
					Name:        "redirect_url",
					Description: `(Required) The URL to redirect to.`,
				},
				resource.Attribute{
					Name:        "select_pool_action",
					Description: `(At least one action is required) Set of pool selection actions to be executed when load balancer rule matches:`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Required) The loadbalancer pool the request will be forwarded to. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb rule.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb rule can be [imported][docs-import] into this resource, via the following command: } } [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_http_forwarding_rule.lb_rule UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb rule named ` + "`" + `lb_rule` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb rule.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb rule can be [imported][docs-import] into this resource, via the following command: } } [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_http_forwarding_rule.lb_rule UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb rule named ` + "`" + `lb_rule` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_http_monitor",
			Category:         "Manager Load Balancer Monitor Resources",
			ShortDescription: `Provides a resource to configure lb http monitor on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"load",
				"balancer",
				"monitor",
				"lb",
				"http",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb http monitor.`,
				},
				resource.Attribute{
					Name:        "fall_count",
					Description: `(Optional) Number of consecutive checks that must fail before marking it down.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) The frequency at which the system issues the monitor check (in seconds).`,
				},
				resource.Attribute{
					Name:        "monitor_port",
					Description: `(Optional) If the monitor port is specified, it would override pool member port setting for healthcheck. A port range is not supported.`,
				},
				resource.Attribute{
					Name:        "rise_count",
					Description: `(Optional) Number of consecutive checks that must pass before marking it up.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Number of seconds the target has to respond to the monitor request.`,
				},
				resource.Attribute{
					Name:        "request_body",
					Description: `(Optional) String to send as HTTP health check request body. Valid only for certain HTTP methods like POST.`,
				},
				resource.Attribute{
					Name:        "request_header",
					Description: `(Optional) HTTP request headers.`,
				},
				resource.Attribute{
					Name:        "request_method",
					Description: `(Optional) Health check method for HTTP monitor type. Valid values are GET, HEAD, PUT, POST and OPTIONS.`,
				},
				resource.Attribute{
					Name:        "request_url",
					Description: `(Optional) URL used for HTTP monitor.`,
				},
				resource.Attribute{
					Name:        "request_version",
					Description: `(Optional) HTTP request version. Valid values are HTTP_VERSION_1_0 and HTTP_VERSION_1_1.`,
				},
				resource.Attribute{
					Name:        "response_body",
					Description: `(Optional) If response body is specified, healthcheck HTTP response body is matched against the specified string and server is considered healthy only if there is a match (regular expressions not supported). If response body string is not specified, HTTP healthcheck is considered successful if the HTTP response status code is among configured values.`,
				},
				resource.Attribute{
					Name:        "response_status_codes",
					Description: `(Optional) HTTP response status code should be a valid HTTP status code. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb_http_monitor.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb http monitor can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_http_monitor.lb_http_monitor UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb http monitor named ` + "`" + `lb_http_monitor` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb_http_monitor.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb http monitor can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_http_monitor.lb_http_monitor UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb http monitor named ` + "`" + `lb_http_monitor` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_http_request_rewrite_rule",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure lb rule on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"http",
				"request",
				"rewrite",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb rule.`,
				},
				resource.Attribute{
					Name:        "match_strategy",
					Description: `(Required) Strategy to define how load balancer rule is considered a match when multiple match conditions are specified in one rule. If set to ALL, then load balancer rule is considered a match only if all the conditions match. If set to ANY, then load balancer rule is considered a match if any one of the conditions match.`,
				},
				resource.Attribute{
					Name:        "body_condition",
					Description: `(Optional) Set of match conditions used to match http request body:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value to look for in the body.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Defines how value field is used to match the body of HTTP requests. Accepted values are STARTS_WITH, ENDS_WITH, CONTAINS, EQUALS, REGEX.`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) If true, case is significant in the match. Default is true.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "header_condition",
					Description: `(Optional) Set of match conditions used to match http request header:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of HTTP header to match.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of HTTP header to match.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Defines how value field is used to match the header value of HTTP requests. Accepted values are STARTS_WITH, ENDS_WITH, CONTAINS, EQUALS, REGEX. Header name field does not support match types.`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) If true, case is significant in the match. Default is true.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "cookie_condition",
					Description: `(Optional) Set of match conditions used to match http request cookie:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of cookie to match.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of cookie to match.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Defines how value field is used to match the cookie. Accepted values are STARTS_WITH, ENDS_WITH, CONTAINS, EQUALS, REGEX.`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) If true, case is significant in the match. Default is true.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "method_condition",
					Description: `(Optional) Set of match conditions used to match http request method:`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) One of GET, HEAD, POST, PUT, OPTIONS.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "version_condition",
					Description: `(Optional) Match condition used to match http version of the request:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) One of HTTP_VERSION_1_0, HTTP_VERSION_1_1.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "uri_condition",
					Description: `(Optional) Set of match conditions used to match http request URI:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Required) The value of URI to match.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Defines how value field is used to match the URI. Accepted values are STARTS_WITH, ENDS_WITH, CONTAINS, EQUALS, REGEX.`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) If true, case is significant in the match. Default is true.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "uri_arguments_condition",
					Description: `(Optional) Set of match conditions used to match http request URI arguments (query string):`,
				},
				resource.Attribute{
					Name:        "uri_arguments",
					Description: `(Required) Query string of URI, typically contains key value pairs.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Defines how value field is used to match the URI. Accepted values are STARTS_WITH, ENDS_WITH, CONTAINS, EQUALS, REGEX.`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) If true, case is significant in the match. Default is true.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "ip_condition",
					Description: `(Optional) Set of match conditions used to match IP header values of HTTP request:`,
				},
				resource.Attribute{
					Name:        "source_address",
					Description: `(Required) The value source IP address to match.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "header_rewrite_action",
					Description: `(At least one action is required) Set of header rewrite actions to be executed when load balancer rule matches:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of HTTP header to be rewritten.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The new value of HTTP header.`,
				},
				resource.Attribute{
					Name:        "uri_rewrite_action",
					Description: `(At least one action is required) Set of URI rewrite actions to be executed when load balancer rule matches:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Required) The new URI for the HTTP request.`,
				},
				resource.Attribute{
					Name:        "uri_arguments",
					Description: `(Required) The new URI arguments(query string) for the HTTP request. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb rule.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb rule can be [imported][docs-import] into this resource, via the following command: } } [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_http_request_rewrite_rule.lb_rule UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb rule named ` + "`" + `lb_rule` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb rule.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb rule can be [imported][docs-import] into this resource, via the following command: } } [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_http_request_rewrite_rule.lb_rule UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb rule named ` + "`" + `lb_rule` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_http_response_rewrite_rule",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure lb rule on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"http",
				"response",
				"rewrite",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb rule.`,
				},
				resource.Attribute{
					Name:        "match_strategy",
					Description: `(Required) Strategy to define how load balancer rule is considered a match when multiple match conditions are specified in one rule. If set to ALL, then load balancer rule is considered a match only if all the conditions match. If set to ANY, then load balancer rule is considered a match if any one of the conditions match.`,
				},
				resource.Attribute{
					Name:        "request_header_condition",
					Description: `(Optional) Set of match conditions used to match http request header:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of HTTP header to match.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of HTTP header to match.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Defines how value field is used to match the header value of HTTP request. Accepted values are STARTS_WITH, ENDS_WITH, CONTAINS, EQUALS, REGEX. Header name field does not support match types.`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) If true, case is significant in the match. Default is true.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "response_header_condition",
					Description: `(Optional) Set of match conditions used to match http response header:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of HTTP header to match.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of HTTP header to match.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Defines how value field is used to match the header value of HTTP response. Accepted values are STARTS_WITH, ENDS_WITH, CONTAINS, EQUALS, REGEX. Header name field does not support match types.`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) If true, case is significant in the match. Default is true.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "cookie_condition",
					Description: `(Optional) Set of match conditions used to match http request cookie:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of cookie to match.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of cookie to match.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Defines how value field is used to match the cookie. Accepted values are STARTS_WITH, ENDS_WITH, CONTAINS, EQUALS, REGEX.`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) If true, case is significant in the match. Default is true.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "method_condition",
					Description: `(Optional) Set of match conditions used to match http request method:`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) One of GET, HEAD, POST, PUT, OPTIONS.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "version_condition",
					Description: `(Optional) Match condition used to match http version of the request:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) One of HTTP_VERSION_1_0, HTTP_VERSION_1_1.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "uri_condition",
					Description: `(Optional) Set of match conditions used to match http request URI:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Required) The value of URI to match.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Defines how value field is used to match the URI. Accepted values are STARTS_WITH, ENDS_WITH, CONTAINS, EQUALS, REGEX.`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) If true, case is significant in the match. Default is true.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "uri_arguments_condition",
					Description: `(Optional) Set of match conditions used to match http request URI arguments (query string):`,
				},
				resource.Attribute{
					Name:        "uri_arguments",
					Description: `(Required) Query string of URI, typically contains key value pairs.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Defines how value field is used to match the URI. Accepted values are STARTS_WITH, ENDS_WITH, CONTAINS, EQUALS, REGEX.`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `(Optional) If true, case is significant in the match. Default is true.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "ip_condition",
					Description: `(Optional) Set of match conditions used to match IP header values of HTTP message:`,
				},
				resource.Attribute{
					Name:        "source_address",
					Description: `(Required) The value source IP address to match.`,
				},
				resource.Attribute{
					Name:        "inverse",
					Description: `(Optional) A flag to indicate whether reverse the match result of this condition. Default is false.`,
				},
				resource.Attribute{
					Name:        "header_rewrite_action",
					Description: `(Required) Set of header rewrite actions to be executed on the outgoing response when load balancer rule matches:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of HTTP header to be rewritten.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The new value of HTTP header. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb rule.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb rule can be [imported][docs-import] into this resource, via the following command: } } [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_http_response_rewrite_rule.lb_rule UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb rule named ` + "`" + `lb_rule` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb rule.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb rule can be [imported][docs-import] into this resource, via the following command: } } [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_http_response_rewrite_rule.lb_rule UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb rule named ` + "`" + `lb_rule` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_http_virtual_server",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure lb http virtual server on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"http",
				"virtual",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the virtual server is enabled. Default is true.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) Virtual server IP address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Virtual server port.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb http virtual server.`,
				},
				resource.Attribute{
					Name:        "access_log_enabled",
					Description: `(Optional) Whether access log is enabled. Default is false.`,
				},
				resource.Attribute{
					Name:        "application_profile_id",
					Description: `(Required) The application profile defines the application protocol characteristics.`,
				},
				resource.Attribute{
					Name:        "default_pool_member_port",
					Description: `(Optional) Default pool member port.`,
				},
				resource.Attribute{
					Name:        "max_concurrent_connections",
					Description: `(Optional) To ensure one virtual server does not over consume resources, affecting other applications hosted on the same LBS, connections to a virtual server can be capped. If it is not specified, it means that connections are unlimited.`,
				},
				resource.Attribute{
					Name:        "max_new_connection_rate",
					Description: `(Optional) To ensure one virtual server does not over consume resources, connections to a member can be rate limited. If it is not specified, it means that connection rate is unlimited.`,
				},
				resource.Attribute{
					Name:        "persistence_profile_id",
					Description: `(Optional) Persistence profile is used to allow related client connections to be sent to the same backend server.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Optional) Pool of backend servers. Server pool consists of one or more servers, also referred to as pool members, that are similarly configured and are running the same application.`,
				},
				resource.Attribute{
					Name:        "sorry_pool_id",
					Description: `(Optional) When load balancer can not select a backend server to serve the request in default pool or pool in rules, the request would be served by sorry server pool.`,
				},
				resource.Attribute{
					Name:        "rule_ids",
					Description: `(Optional) List of load balancer rules that provide customization of load balancing behavior using match/action rules.`,
				},
				resource.Attribute{
					Name:        "client_ssl",
					Description: `(Optional) Client side SSL customization.`,
				},
				resource.Attribute{
					Name:        "client_ssl_profile_id",
					Description: `(Required) Id of client SSL profile that defines reusable properties.`,
				},
				resource.Attribute{
					Name:        "default_certificate_id",
					Description: `(Required) Id of certificate that will be used if the server does not host multiple hostnames on the same IP address or if the client does not support SNI extension.`,
				},
				resource.Attribute{
					Name:        "certificate_chain_depth",
					Description: `(Optional) Allowed depth of certificate chain. Default is 3.`,
				},
				resource.Attribute{
					Name:        "client_auth",
					Description: `(Optional) Whether client authentication is mandatory. Default is false.`,
				},
				resource.Attribute{
					Name:        "ca_ids",
					Description: `(Optional) List of CA certificate ids for client authentication.`,
				},
				resource.Attribute{
					Name:        "crl_ids",
					Description: `(Optional) List of CRL certificate ids for client authentication.`,
				},
				resource.Attribute{
					Name:        "sni_certificate_ids",
					Description: `(Optional) List of certificates to serve different hostnames.`,
				},
				resource.Attribute{
					Name:        "server_ssl",
					Description: `(Optional) Server side SSL customization.`,
				},
				resource.Attribute{
					Name:        "server_ssl_profile_id",
					Description: `(Required) Id of server SSL profile that defines reusable properties.`,
				},
				resource.Attribute{
					Name:        "server_auth",
					Description: `(Optional) Whether server authentication is needed. Default is False. If true, ca_ids should be provided.`,
				},
				resource.Attribute{
					Name:        "certificate_chain_depth",
					Description: `(Optional) Allowed depth of certificate chain. Default is 3.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(Optional) Whether server authentication is required. Default is false.`,
				},
				resource.Attribute{
					Name:        "ca_ids",
					Description: `(Optional) List of CA certificate ids for server authentication.`,
				},
				resource.Attribute{
					Name:        "crl_ids",
					Description: `(Optional) List of CRL certificate ids for server authentication. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb http virtual server.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb http virtual server can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_http_virtual_server.lb_http_virtual_server UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb http virtual server named ` + "`" + `lb_http_virtual_server` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb http virtual server.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb http virtual server can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_http_virtual_server.lb_http_virtual_server UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb http virtual server named ` + "`" + `lb_http_virtual_server` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_https_monitor",
			Category:         "Manager Load Balancer Monitor Resources",
			ShortDescription: `Provides a resource to configure lb https monitor on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"load",
				"balancer",
				"monitor",
				"lb",
				"https",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb https monitor.`,
				},
				resource.Attribute{
					Name:        "fall_count",
					Description: `(Optional) Number of consecutive checks that must fail before marking it down.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) The frequency at which the system issues the monitor check (in seconds).`,
				},
				resource.Attribute{
					Name:        "monitor_port",
					Description: `(Optional) If the monitor port is specified, it would override pool member port setting for healthcheck. A port range is not supported.`,
				},
				resource.Attribute{
					Name:        "rise_count",
					Description: `(Optional) Number of consecutive checks that must pass before marking it up.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Number of seconds the target has to respond to the monitor request.`,
				},
				resource.Attribute{
					Name:        "certificate_chain_depth",
					Description: `(Optional) Authentication depth is used to set the verification depth in the server certificates chain.`,
				},
				resource.Attribute{
					Name:        "ciphers",
					Description: `(Optional) List of supported SSL ciphers.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(Optional) Client certificate can be specified to support client authentication.`,
				},
				resource.Attribute{
					Name:        "protocols",
					Description: `(Optional) SSL versions TLS1.1 and TLS1.2 are supported and enabled by default. SSLv2, SSLv3, and TLS1.0 are supported, but disabled by default.`,
				},
				resource.Attribute{
					Name:        "request_body",
					Description: `(Optional) String to send as HTTP health check request body. Valid only for certain HTTP methods like POST.`,
				},
				resource.Attribute{
					Name:        "request_header",
					Description: `(Optional) HTTP request headers.`,
				},
				resource.Attribute{
					Name:        "request_method",
					Description: `(Optional) Health check method for HTTP monitor type. Valid values are GET, HEAD, PUT, POST and OPTIONS.`,
				},
				resource.Attribute{
					Name:        "request_url",
					Description: `(Optional) URL used for HTTP monitor.`,
				},
				resource.Attribute{
					Name:        "request_version",
					Description: `(Optional) HTTP request version. Valid values are HTTP_VERSION_1_0 and HTTP_VERSION_1_1.`,
				},
				resource.Attribute{
					Name:        "response_body",
					Description: `(Optional) If response body is specified, healthcheck HTTP response body is matched against the specified string and server is considered healthy only if there is a match (regular expressions not supported). If response body string is not specified, HTTP healthcheck is considered successful if the HTTP response status code is among configured values.`,
				},
				resource.Attribute{
					Name:        "response_status_codes",
					Description: `(Optional) HTTP response status code should be a valid HTTP status code.`,
				},
				resource.Attribute{
					Name:        "server_auth",
					Description: `(Optional) Server authentication mode - REQUIRED or IGNORE.`,
				},
				resource.Attribute{
					Name:        "server_auth_ca_ids",
					Description: `(Optional) If server auth type is REQUIRED, server certificate must be signed by one of the trusted Certificate Authorities (CAs), also referred to as root CAs, whose self signed certificates are specified.`,
				},
				resource.Attribute{
					Name:        "server_auth_crl_ids",
					Description: `(Optional) A Certificate Revocation List (CRL) can be specified in the server-side SSL profile binding to disallow compromised server certificates. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb_https_monitor.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "is_secure",
					Description: `This flag is set to true when all the ciphers and protocols are secure. It is set to false when one of the ciphers or protocols is insecure. ## Importing An existing lb https monitor can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_https_monitor.lb_https_monitor UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb https monitor named ` + "`" + `lb_https_monitor` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb_https_monitor.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "is_secure",
					Description: `This flag is set to true when all the ciphers and protocols are secure. It is set to false when one of the ciphers or protocols is insecure. ## Importing An existing lb https monitor can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_https_monitor.lb_https_monitor UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb https monitor named ` + "`" + `lb_https_monitor` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_icmp_monitor",
			Category:         "Manager Load Balancer Monitor Resources",
			ShortDescription: `Provides a resource to configure load balancer icmp monitor on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"load",
				"balancer",
				"monitor",
				"lb",
				"icmp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb icmp monitor.`,
				},
				resource.Attribute{
					Name:        "fall_count",
					Description: `(Optional) Number of consecutive checks must fail before marking it down.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) The frequency at which the system issues the monitor check (in seconds).`,
				},
				resource.Attribute{
					Name:        "monitor_port",
					Description: `(Optional) If the monitor port is specified, it would override pool member port setting for healthcheck. Port range is not supported.`,
				},
				resource.Attribute{
					Name:        "rise_count",
					Description: `(Optional) Number of consecutive checks must pass before marking it up.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Number of seconds the target has in which to respond to the monitor request.`,
				},
				resource.Attribute{
					Name:        "data_length",
					Description: `(Optional) The data size (in bytes) of the ICMP healthcheck packet. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb_icmp_monitor.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb icmp monitor can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_icmp_monitor.lb_icmp_monitor UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb icmp monitor named ` + "`" + `lb_icmp_monitor` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb_icmp_monitor.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb icmp monitor can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_icmp_monitor.lb_icmp_monitor UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb icmp monitor named ` + "`" + `lb_icmp_monitor` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_passive_monitor",
			Category:         "Manager Load Balancer Monitor Resources",
			ShortDescription: `Provides a resource to configure load balancer passive monitor on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"load",
				"balancer",
				"monitor",
				"lb",
				"passive",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb passive monitor.`,
				},
				resource.Attribute{
					Name:        "max_fails",
					Description: `(Optional) When consecutive failures reach this value, the member is considered temporarily unavailable for a configurable period.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) After this timeout period, the member is probed again. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb_passive_monitor.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb passive monitor can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_passive_monitor.lb_passive_monitor UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb passive monitor named ` + "`" + `lb_passive_monitor` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb_passive_monitor.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb passive monitor can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_passive_monitor.lb_passive_monitor UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb passive monitor named ` + "`" + `lb_passive_monitor` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_pool",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure lb pool on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "active_monitor_id",
					Description: `(Optional) Active health monitor Id. If one is not set, the active healthchecks will be disabled.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Optional) Load balancing algorithm controls how the incoming connections are distributed among the members. Supported algorithms are: ROUND_ROBIN, WEIGHTED_ROUND_ROBIN, LEAST_CONNECTION, WEIGHTED_LEAST_CONNECTION, IP_HASH.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Optional) Server pool consists of one or more pool members. Each pool member is identified, typically, by an IP address and a port. Each member has the following arguments:`,
				},
				resource.Attribute{
					Name:        "admin_state",
					Description: `(Optional) Pool member admin state. Possible values: ENABLED, DISABLED and GRACEFUL_DISABLED`,
				},
				resource.Attribute{
					Name:        "backup_member",
					Description: `(Optional) A boolean flag which reflects whether this is a backup pool member. Backup servers are typically configured with a sorry page indicating to the user that the application is currently unavailable. While the pool is active (a specified minimum number of pool members are active) BACKUP members are skipped during server selection. When the pool is inactive, incoming connections are sent to only the BACKUP member(s).`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. pool member name.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) Pool member IP address.`,
				},
				resource.Attribute{
					Name:        "max_concurrent_connections",
					Description: `(Optional) To ensure members are not overloaded, connections to a member can be capped by the load balancer. When a member reaches this limit, it is skipped during server selection. If it is not specified, it means that connections are unlimited.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) If port is specified, all connections will be sent to this port. Only single port is supported. If unset, the same port the client connected to will be used, it could be overrode by default_pool_member_port setting in virtual server. The port should not specified for port range case.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Pool member weight is used for WEIGHTED_ROUND_ROBIN balancing algorithm. The weight value would be ignored in other algorithms.`,
				},
				resource.Attribute{
					Name:        "member_group",
					Description: `(Optional) Dynamic pool members for the loadbalancing pool. When member group is defined, members setting should not be specified. The member_group has the following arguments:`,
				},
				resource.Attribute{
					Name:        "grouping_object",
					Description: `(Required) Grouping object of type NSGroup which will be used as dynamic pool members. The IP list of the grouping object would be used as pool member IP setting.`,
				},
				resource.Attribute{
					Name:        "ip_version_filter",
					Description: `(Optional) Ip version filter is used to filter IPv4 or IPv6 addresses from the grouping object. If the filter is not specified, both IPv4 and IPv6 addresses would be used as server IPs. Supported filtering is "IPV4" and "IPV6" ("IPV4" is the default one)`,
				},
				resource.Attribute{
					Name:        "limit_ip_list_size",
					Description: `(Optional) Limits the max number of pool members. If false, allows the dynamic pool to grow up to the load balancer max pool member capacity.`,
				},
				resource.Attribute{
					Name:        "max_ip_list_size",
					Description: `(Optional) Should only be specified if limit_ip_list_size is set to true. Limits the max number of pool members to the specified value.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) If port is specified, all connections will be sent to this port. If unset, the same port the client connected to will be used, it could be overridden by default_pool_member_ports setting in virtual server. The port should not specified for multiple ports case.`,
				},
				resource.Attribute{
					Name:        "min_active_members",
					Description: `(Optional) The minimum number of members for the pool to be considered active. This value is 1 by default.`,
				},
				resource.Attribute{
					Name:        "passive_monitor_id",
					Description: `(Optional) Passive health monitor Id. If one is not set, the passive healthchecks will be disabled.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of SNAT performed to ensure reverse traffic from the server can be received and processed by the loadbalancer. Supported types are: SNAT_AUTO_MAP, SNAT_IP_POOL and TRANSPARENT`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required for snat_translation of type SNAT_IP_POOL) Ip address or Ip range for SNAT of type SNAT_IP_POOL.`,
				},
				resource.Attribute{
					Name:        "tcp_multiplexing_enabled",
					Description: `(Optional) TCP multiplexing allows the same TCP connection between load balancer and the backend server to be used for sending multiple client requests from different client TCP connections. Disabled by default.`,
				},
				resource.Attribute{
					Name:        "tcp_multiplexing_number",
					Description: `(Optional) The maximum number of TCP connections per pool that are idly kept alive for sending future client requests. The default value for this is 6.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb pool. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb pool.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb pool can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_pool.lb_pool UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb pool named ` + "`" + `lb_pool` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb pool.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb pool can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_pool.lb_pool UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb pool named ` + "`" + `lb_pool` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_server_ssl_profile",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure lb server ssl profile on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"server",
				"ssl",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb server ssl profile.`,
				},
				resource.Attribute{
					Name:        "ciphers",
					Description: `(Optional) supported SSL cipher list to client side. The supported ciphers can contain: TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA, TLS_ECDH_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA, TLS_RSA_WITH_3DES_EDE_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384, TLS_RSA_WITH_AES_128_CBC_SHA256, TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256, TLS_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256, TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384, TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA256, TLS_ECDH_ECDSA_WITH_AES_128_GCM_SHA256, TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA384, TLS_ECDH_ECDSA_WITH_AES_256_GCM_SHA384, TLS_ECDH_RSA_WITH_AES_128_CBC_SHA, TLS_ECDH_RSA_WITH_AES_128_CBC_SHA256, TLS_ECDH_RSA_WITH_AES_128_GCM_SHA256, TLS_ECDH_RSA_WITH_AES_256_CBC_SHA384, TLS_ECDH_RSA_WITH_AES_256_GCM_SHA384.`,
				},
				resource.Attribute{
					Name:        "prefer_server_ciphers",
					Description: `(Optional) During SSL handshake as part of the SSL client Hello client sends an ordered list of ciphers that it can support (or prefers) and typically server selects the first one from the top of that list it can also support. For Perfect Forward Secrecy(PFS), server could override the client's preference. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "protocols",
					Description: `(Optional) SSL versions TLS_V1_1 and TLS_V1_2 are supported and enabled by default. SSL_V2, SSL_V3, and TLS_V1 are supported, but disabled by default.`,
				},
				resource.Attribute{
					Name:        "session_cache_enabled",
					Description: `(Optional) SSL session caching allows SSL server and server to reuse previously negotiated security parameters avoiding the expensive public key operation during handshake. Defaults to true. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb server ssl profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "is_secure",
					Description: `This flag is set to true when all the ciphers and protocols are secure. It is set to false when one of the ciphers or protocols is insecure. ## Importing An existing lb server ssl profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_server_ssl_profile.lb_server_ssl_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb server ssl profile named ` + "`" + `lb_server_ssl_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb server ssl profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "is_secure",
					Description: `This flag is set to true when all the ciphers and protocols are secure. It is set to false when one of the ciphers or protocols is insecure. ## Importing An existing lb server ssl profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_server_ssl_profile.lb_server_ssl_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb server ssl profile named ` + "`" + `lb_server_ssl_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_service",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure lb service on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb service.`,
				},
				resource.Attribute{
					Name:        "logical_router_id",
					Description: `(Required) Tier1 logical router this service is attached to. Note that this router needs to have edge cluster configured, and have an uplink port or CSP (centralized service port).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) whether the load balancer service is enabled.`,
				},
				resource.Attribute{
					Name:        "error_log_level",
					Description: `(Optional) Load balancer engine writes information about encountered issues of different severity levels to the error log. This setting is used to define the severity level of the error log.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Size of load balancer service. Accepted values are SMALL/MEDIUM/LARGE.`,
				},
				resource.Attribute{
					Name:        "virtual_server_ids",
					Description: `(Optional) Virtual servers associated with this Load Balancer. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb_service.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_service.lb_service UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb service named ` + "`" + `lb_service` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb_service.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_service.lb_service UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb service named ` + "`" + `lb_service` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_source__persistence_profile",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure lb source ip persistence profile on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"source",
				"persistence",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb source ip persistence profile.`,
				},
				resource.Attribute{
					Name:        "persistence_shared",
					Description: `(Optional) A boolean flag which reflects whether the cookie persistence is private or shared.`,
				},
				resource.Attribute{
					Name:        "ha_persistence_mirroring",
					Description: `(Optional) A boolean flag which reflects whether persistence entries will be synchronized to the HA peer.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Persistence expiration time in seconds, counted from the time all the connections are completed. Defaults to 300 seconds.`,
				},
				resource.Attribute{
					Name:        "purge_when_full",
					Description: `(Optional) A boolean flag which reflects whether entries will be purged when the persistence table is full. Defaults to true. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb source ip persistence profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb source ip persistence profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_source_ip_persistence_profile.lb_source_ip_persistence_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb source ip persistence profile named ` + "`" + `lb_source_ip_persistence_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb source ip persistence profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb source ip persistence profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_source_ip_persistence_profile.lb_source_ip_persistence_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb source ip persistence profile named ` + "`" + `lb_source_ip_persistence_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_tcp_monitor",
			Category:         "Manager Load Balancer Monitor Resources",
			ShortDescription: `Provides a resource to configure load balancer tcp monitor on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"load",
				"balancer",
				"monitor",
				"lb",
				"tcp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb tcp monitor.`,
				},
				resource.Attribute{
					Name:        "fall_count",
					Description: `(Optional) Number of consecutive checks must fail before marking it down.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) The frequency at which the system issues the monitor check (in seconds).`,
				},
				resource.Attribute{
					Name:        "monitor_port",
					Description: `(Optional) If the monitor port is specified, it would override pool member port setting for healthcheck. Port range is not supported.`,
				},
				resource.Attribute{
					Name:        "rise_count",
					Description: `(Optional) Number of consecutive checks must pass before marking it up.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Number of seconds the target has in which to respond to the monitor request.`,
				},
				resource.Attribute{
					Name:        "receive",
					Description: `(Optional) Expected data, if specified, can be anywhere in the response and it has to be a string, regular expressions are not supported.`,
				},
				resource.Attribute{
					Name:        "send",
					Description: `(Optional) Payload to send out to the monitored server. If both send and receive are not specified, then just a TCP connection is established (3-way handshake) to validate server is healthy, no data is sent. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb_tcp_monitor.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb tcp monitor can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_tcp_monitor.lb_tcp_monitor UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb tcp monitor named ` + "`" + `lb_tcp_monitor` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb_tcp_monitor.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb tcp monitor can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_tcp_monitor.lb_tcp_monitor UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb tcp monitor named ` + "`" + `lb_tcp_monitor` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_tcp_virtual_server",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure lb tcp virtual server on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"tcp",
				"virtual",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the virtual server is enabled. Default is true.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) Virtual server IP address.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Required) List of virtual server ports.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb tcp virtual server.`,
				},
				resource.Attribute{
					Name:        "access_log_enabled",
					Description: `(Optional) Whether access log is enabled. Default is false.`,
				},
				resource.Attribute{
					Name:        "application_profile_id",
					Description: `(Required) The application profile defines the application protocol characteristics.`,
				},
				resource.Attribute{
					Name:        "default_pool_member_ports",
					Description: `(Optional) List of default pool member ports.`,
				},
				resource.Attribute{
					Name:        "max_concurrent_connections",
					Description: `(Optional) To ensure one virtual server does not over consume resources, affecting other applications hosted on the same LBS, connections to a virtual server can be capped. If it is not specified, it means that connections are unlimited.`,
				},
				resource.Attribute{
					Name:        "max_new_connection_rate",
					Description: `(Optional) To ensure one virtual server does not over consume resources, connections to a member can be rate limited. If it is not specified, it means that connection rate is unlimited.`,
				},
				resource.Attribute{
					Name:        "persistence_profile_id",
					Description: `(Optional) Persistence profile is used to allow related client connections to be sent to the same backend server. Only source ip persistance profile is accepted.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Optional) Pool of backend servers. Server pool consists of one or more servers, also referred to as pool members, that are similarly configured and are running the same application.`,
				},
				resource.Attribute{
					Name:        "sorry_pool_id",
					Description: `(Optional) When load balancer can not select a backend server to serve the request in default pool or pool in rules, the request would be served by sorry server pool. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb tcp virtual server.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb tcp virtual server can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_tcp_virtual_server.lb_tcp_virtual_server UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb tcp virtual server named ` + "`" + `lb_tcp_virtual_server` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb tcp virtual server.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb tcp virtual server can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_tcp_virtual_server.lb_tcp_virtual_server UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb tcp virtual server named ` + "`" + `lb_tcp_virtual_server` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_udp_monitor",
			Category:         "Manager Load Balancer Monitor Resources",
			ShortDescription: `Provides a resource to configure load balancer udp monitor on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"load",
				"balancer",
				"monitor",
				"lb",
				"udp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb udp monitor.`,
				},
				resource.Attribute{
					Name:        "fall_count",
					Description: `(Optional) Number of consecutive checks must fail before marking it down.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) The frequency at which the system issues the monitor check (in seconds).`,
				},
				resource.Attribute{
					Name:        "monitor_port",
					Description: `(Optional) If the monitor port is specified, it would override pool member port setting for healthcheck. Port range is not supported.`,
				},
				resource.Attribute{
					Name:        "rise_count",
					Description: `(Optional) Number of consecutive checks must pass before marking it up.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Number of seconds the target has in which to respond to the monitor request.`,
				},
				resource.Attribute{
					Name:        "receive",
					Description: `(Required) Expected data, if specified, can be anywhere in the response and it has to be a string, regular expressions are not supported.`,
				},
				resource.Attribute{
					Name:        "send",
					Description: `(Required) Payload to send out to the monitored server. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb_udp_monitor.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb udp monitor can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_udp_monitor.lb_udp_monitor UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb udp monitor named ` + "`" + `lb_udp_monitor` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb_udp_monitor.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb udp monitor can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_udp_monitor.lb_udp_monitor UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb udp monitor named ` + "`" + `lb_udp_monitor` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_lb_udp_virtual_server",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure lb udp virtual server on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"lb",
				"udp",
				"virtual",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the virtual server is enabled. Default is true.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) Virtual server IP address.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Required) List of virtual server port.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this lb udp virtual server.`,
				},
				resource.Attribute{
					Name:        "access_log_enabled",
					Description: `(Optional) Whether access log is enabled. Default is false.`,
				},
				resource.Attribute{
					Name:        "application_profile_id",
					Description: `(Required) The application profile defines the application protocol characteristics.`,
				},
				resource.Attribute{
					Name:        "default_pool_member_ports",
					Description: `(Optional) List of default pool member ports.`,
				},
				resource.Attribute{
					Name:        "max_concurrent_connections",
					Description: `(Optional) To ensure one virtual server does not over consume resources, affecting other applications hosted on the same LBS, connections to a virtual server can be capped. If it is not specified, it means that connections are unlimited.`,
				},
				resource.Attribute{
					Name:        "max_new_connection_rate",
					Description: `(Optional) To ensure one virtual server does not over consume resources, connections to a member can be rate limited. If it is not specified, it means that connection rate is unlimited.`,
				},
				resource.Attribute{
					Name:        "persistence_profile_id",
					Description: `(Optional) Persistence profile is used to allow related client connections to be sent to the same backend server. Only source ip persistence profile is accepted.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Optional) Pool of backend servers. Server pool consists of one or more servers, also referred to as pool members, that are similarly configured and are running the same application.`,
				},
				resource.Attribute{
					Name:        "sorry_pool_id",
					Description: `(Optional) When load balancer can not select a backend server to serve the request in default pool or pool in rules, the request would be served by sorry server pool. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb udp virtual server.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb udp virtual server can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_udp_virtual_server.lb_udp_virtual_server UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb udp virtual server named ` + "`" + `lb_udp_virtual_server` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lb udp virtual server.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing lb udp virtual server can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_lb_udp_virtual_server.lb_udp_virtual_server UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the lb udp virtual server named ` + "`" + `lb_udp_virtual_server` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_dhcp_port",
			Category:         "Manager Resources",
			ShortDescription: `A resource that can be used to configure a Logical DHCP Port in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"logical",
				"dhcp",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "logical_switch_id",
					Description: `(Required) Logical switch ID for the logical port.`,
				},
				resource.Attribute{
					Name:        "dhcp_server_id",
					Description: `(Required) Logical DHCP server ID for the logical port.`,
				},
				resource.Attribute{
					Name:        "admin_state",
					Description: `(Optional) Admin state for the logical port. Accepted values - 'UP' or 'DOWN'. The default value is 'UP'.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this logical port. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical DHCP port.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing DHCP Logical Port can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_dhcp_port.dhcp_port UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical DHCP port named ` + "`" + `dhcp_port` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical DHCP port.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing DHCP Logical Port can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_dhcp_port.dhcp_port UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical DHCP port named ` + "`" + `dhcp_port` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_dhcp_server",
			Category:         "Manager Resources",
			ShortDescription: `Provides a resource to configure logical DHCP server on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"logical",
				"dhcp",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "dhcp_profile_id",
					Description: `(Required) DHCP profile uuid.`,
				},
				resource.Attribute{
					Name:        "dhcp_server_ip",
					Description: `(Required) DHCP server IP in cidr format.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Gateway IP.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional) Domain name.`,
				},
				resource.Attribute{
					Name:        "dns_name_servers",
					Description: `(Optional) DNS IPs.`,
				},
				resource.Attribute{
					Name:        "dhcp_option_121",
					Description: `(Optional) DHCP classless static routes.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Destination in cidr format.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required) IP address of next hop.`,
				},
				resource.Attribute{
					Name:        "dhcp_generic_option",
					Description: `(Optional) Generic DHCP options.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(Required) DHCP option code. Valid values are from 0 to 255.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) List of DHCP option values.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this logical DHCP server. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical DHCP server.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "attached_logical_port_id",
					Description: `ID of the attached logical port. ## Importing An existing logical DHCP server can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_dhcp_server.logical_dhcp_server UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the logical DHCP server named ` + "`" + `logical_dhcp_server` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical DHCP server.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "attached_logical_port_id",
					Description: `ID of the attached logical port. ## Importing An existing logical DHCP server can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_dhcp_server.logical_dhcp_server UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the logical DHCP server named ` + "`" + `logical_dhcp_server` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_port",
			Category:         "Manager Resources",
			ShortDescription: `A resource that can be used to configure a Logical Port in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"logical",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "logical_switch_id",
					Description: `(Required) Logical switch ID for the logical port.`,
				},
				resource.Attribute{
					Name:        "admin_state",
					Description: `(Optional) Admin state for the logical port. Accepted values - 'UP' or 'DOWN'. The default value is 'UP'.`,
				},
				resource.Attribute{
					Name:        "switching_profile_id",
					Description: `(Optional) List of IDs of switching profiles (of various types) to be associated with this switch. Default switching profiles will be used if not specified.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this logical port. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical port.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing Logical Port can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_port.logical_port UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical port named ` + "`" + `logical_port` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical port.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing Logical Port can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_port.logical_port UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical port named ` + "`" + `logical_port` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_router_centralized_service_port",
			Category:         "Manager Resources",
			ShortDescription: `A resource that can be used to configure logical router centralized service port in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"logical",
				"router",
				"centralized",
				"service",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_router_id",
					Description: `(Required) Identifier for logical Tier-0 or Tier-1 router on which this port is created`,
				},
				resource.Attribute{
					Name:        "linked_logical_switch_port_id",
					Description: `(Required) Identifier for port on logical switch to connect to`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) Logical router port subnet (ip_address / prefix length)`,
				},
				resource.Attribute{
					Name:        "urpf_mode",
					Description: `(Optional) Unicast Reverse Path Forwarding mode. Accepted values are "NONE" and "STRICT" which is the default value.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this port. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical router centralized service port.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing logical router centralized service port can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_router_centralized_service_port.cs_port UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical router centralized service port named ` + "`" + `cs_port` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical router centralized service port.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing logical router centralized service port can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_router_centralized_service_port.cs_port UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical router centralized service port named ` + "`" + `cs_port` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_router_downlink_port",
			Category:         "Manager Resources",
			ShortDescription: `A resource that can be used to configure logical router downlink port in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"logical",
				"router",
				"downlink",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_router_id",
					Description: `(Required) Identifier for logical Tier-1 router on which this port is created`,
				},
				resource.Attribute{
					Name:        "linked_logical_switch_port_id",
					Description: `(Required) Identifier for port on logical switch to connect to`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) Logical router port subnet (ip_address / prefix length)`,
				},
				resource.Attribute{
					Name:        "urpf_mode",
					Description: `(Optional) Unicast Reverse Path Forwarding mode. Accepted values are "NONE" and "STRICT" which is the default value.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this port.`,
				},
				resource.Attribute{
					Name:        "service_binding",
					Description: `(Optional) A list of services for this port. Currently only "LogicalService" is supported as a target_type, and a DHCP relay service ID as target_id ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical router downlink port.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address assigned to this port ## Importing An existing logical router downlink port can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_router_downlink_port.downlink_port UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical router downlink port named ` + "`" + `downlink_port` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical router downlink port.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address assigned to this port ## Importing An existing logical router downlink port can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_router_downlink_port.downlink_port UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical router downlink port named ` + "`" + `downlink_port` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_router_link_port_on_tier0",
			Category:         "Manager Resources",
			ShortDescription: `A resource that can be used to configure a logical router link port on Tier-0 router in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"logical",
				"router",
				"link",
				"port",
				"on",
				"tier0",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_router_id",
					Description: `(Required) Identifier for logical Tier0 router on which this port is created.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this port. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical router link port.`,
				},
				resource.Attribute{
					Name:        "linked_logical_switch_port_id",
					Description: `Identifier for port on logical router to connect to.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing logical router link port on Tier-0 can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_router_link_port_on_tier0.link_port_tier0 UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical router link port on the tier 0 logical router named ` + "`" + `link_port_tier0` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical router link port.`,
				},
				resource.Attribute{
					Name:        "linked_logical_switch_port_id",
					Description: `Identifier for port on logical router to connect to.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing logical router link port on Tier-0 can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_router_link_port_on_tier0.link_port_tier0 UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical router link port on the tier 0 logical router named ` + "`" + `link_port_tier0` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_router_link_port_on_tier1",
			Category:         "Manager Resources",
			ShortDescription: `A resource to configure a logical router link port on a Tier-1 router in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"logical",
				"router",
				"link",
				"port",
				"on",
				"tier1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_router_id",
					Description: `(Required) Identifier for logical tier-1 router on which this port is created.`,
				},
				resource.Attribute{
					Name:        "linked_logical_switch_port_id",
					Description: `(Required) Identifier for port on logical Tier-0 router to connect to.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this port. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical router link port.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing logical router link port on Tier-1 can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_router_link_port_on_tier1.link_port_tier1 UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical router link port on the tier 1 router named ` + "`" + `link_port_tier1` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical router link port.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing logical router link port on Tier-1 can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_router_link_port_on_tier1.link_port_tier1 UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical router link port on the tier 1 router named ` + "`" + `link_port_tier1` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_switch",
			Category:         "Manager Resources",
			ShortDescription: `A resource to configure overlay logical switch in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"logical",
				"switch",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "transport_zone_id",
					Description: `(Required) Transport Zone ID for the logical switch.`,
				},
				resource.Attribute{
					Name:        "admin_state",
					Description: `(Optional) Admin state for the logical switch. Accepted values - 'UP' or 'DOWN'. The default value is 'UP'.`,
				},
				resource.Attribute{
					Name:        "replication_mode",
					Description: `(Optional) Replication mode of the Logical Switch. Accepted values - 'MTEP' (Hierarchical Two-Tier replication) and 'SOURCE' (Head Replication), with 'MTEP' being the default value. Applies to overlay logical switches.`,
				},
				resource.Attribute{
					Name:        "switching_profile_id",
					Description: `(Optional) List of IDs of switching profiles (of various types) to be associated with this switch. Default switching profiles will be used if not specified.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "ip_pool_id",
					Description: `(Optional) Ip Pool ID to be associated with the logical switch.`,
				},
				resource.Attribute{
					Name:        "mac_pool_id",
					Description: `(Optional) Mac Pool ID to be associated with the logical switch.`,
				},
				resource.Attribute{
					Name:        "address_binding",
					Description: `(Optional) A list address bindings for this logical switch`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) IP Address`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Required) MAC Address`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Optional) Vlan`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Deprecated, Optional) Vlan for vlan logical switch. This attribute is deprecated, please use nsxt_vlan_logical_switch resource to manage vlan logical switches.`,
				},
				resource.Attribute{
					Name:        "vni",
					Description: `(Optional, Readonly) Vni for the logical switch.`,
				},
				resource.Attribute{
					Name:        "address_binding",
					Description: `(Optional) List of Address Bindings for the logical switch. This setting allows to provide bindings between IP address, mac Address and vlan.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this logical switch. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical switch.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing X can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_switch.switch1 UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical switch named ` + "`" + `switch1` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical switch.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing X can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_switch.switch1 UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical switch named ` + "`" + `switch1` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_tier0_router",
			Category:         "Manager Resources",
			ShortDescription: `A resource to configure a logical Tier0 router in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"logical",
				"tier0",
				"router",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `(Required) Edge Cluster ID for the logical Tier0 router. Changing this setting on existing router will re-create the router.`,
				},
				resource.Attribute{
					Name:        "failover_mode",
					Description: `(Optional) Failover mode which determines whether the preferred service router instance for given logical router will preempt the peer. Accepted values are PREEMPTIVE/NON_PREEMPTIVE. This setting is relevant only for ACTIVE_STANDBY high availability mode.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this logical Tier0 router.`,
				},
				resource.Attribute{
					Name:        "high_availability_mode",
					Description: `(Optional) High availability mode "ACTIVE_ACTIVE"/"ACTIVE_STANDBY". Changing this setting on existing router will re-create the router. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical Tier0 router.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "firewall_sections",
					Description: `(Optional) The list of firewall sections for this router ## Importing An existing logical tier0 router can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_tier0_router.tier0_router UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical tier 0 router named ` + "`" + `tier0_router` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical Tier0 router.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "firewall_sections",
					Description: `(Optional) The list of firewall sections for this router ## Importing An existing logical tier0 router can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_tier0_router.tier0_router UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical tier 0 router named ` + "`" + `tier0_router` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_tier1_router",
			Category:         "Manager Resources",
			ShortDescription: `A resource to configure a logical Tier1 router in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"logical",
				"tier1",
				"router",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `(Optional) Edge Cluster ID for the logical Tier1 router.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this logical Tier1 router.`,
				},
				resource.Attribute{
					Name:        "failover_mode",
					Description: `(Optional) This failover mode determines, whether the preferred service router instance for given logical router will preempt the peer. Note - It can be specified if and only if logical router is ACTIVE_STANDBY and NON_PREEMPTIVE mode is supported only for a Tier1 logical router. For ACTIVE_ACTIVE logical routers, this field must not be populated`,
				},
				resource.Attribute{
					Name:        "enable_router_advertisement",
					Description: `(Optional) Enable the router advertisement`,
				},
				resource.Attribute{
					Name:        "advertise_connected_routes",
					Description: `(Optional) Enable the router advertisement for all NSX connected routes`,
				},
				resource.Attribute{
					Name:        "advertise_static_routes",
					Description: `(Optional) Enable the router advertisement for static routes`,
				},
				resource.Attribute{
					Name:        "advertise_nat_routes",
					Description: `(Optional) Enable the router advertisement for NAT routes`,
				},
				resource.Attribute{
					Name:        "advertise_lb_vip_routes",
					Description: `(Optional) Enable the router advertisement for LB VIP routes`,
				},
				resource.Attribute{
					Name:        "advertise_lb_snat_ip_routes",
					Description: `(Optional) Enable the router advertisement for LB SNAT IP routes ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical Tier1 router.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "advertise_config_revision",
					Description: `Indicates current revision number of the advertisement configuration object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "firewall_sections",
					Description: `(Optional) The list of firewall sections for this router ## Importing An existing logical tier1 router can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_tier1_router.tier1_router UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical tier 1 router named ` + "`" + `tier1_router` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical Tier1 router.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "advertise_config_revision",
					Description: `Indicates current revision number of the advertisement configuration object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "firewall_sections",
					Description: `(Optional) The list of firewall sections for this router ## Importing An existing logical tier1 router can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_logical_tier1_router.tier1_router UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical tier 1 router named ` + "`" + `tier1_router` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_mac_management_switching_profile",
			Category:         "Manager Switching Profiles Resources",
			ShortDescription: `Provides a resource to configure MAC management switching profile on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"switching",
				"profiles",
				"mac",
				"management",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this MAC management switching profile.`,
				},
				resource.Attribute{
					Name:        "mac_change_allowed",
					Description: `(Optional) A boolean flag indicating allowing source MAC address change.`,
				},
				resource.Attribute{
					Name:        "mac_learning",
					Description: `(Optional) Mac learning configuration:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) A boolean flag indicating allowing source MAC address learning.`,
				},
				resource.Attribute{
					Name:        "unicast_flooding_allowed",
					Description: `(Optional) A boolean flag indicating allowing flooding for unlearned MAC for ingress traffic. Can be True only if mac_learning is enabled.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) The maximum number of MAC addresses that can be learned on this port.`,
				},
				resource.Attribute{
					Name:        "limit_policy",
					Description: `(Optional) The policy after MAC Limit is exceeded: ALLOW/DROP. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the MAC management switching profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing MAC management switching profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_mac_management_switching_profile.mac_management_switching_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the MAC management switching profile named ` + "`" + `mac_management_switching_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the MAC management switching profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing MAC management switching profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_mac_management_switching_profile.mac_management_switching_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the MAC management switching profile named ` + "`" + `mac_management_switching_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_nat_rule",
			Category:         "Manager Resources",
			ShortDescription: `A resource to configure a NAT rule in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"nat",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_router_id",
					Description: `(Required) ID of the logical router.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this NAT rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) NAT rule action type. Valid actions are: SNAT, DNAT, NO_NAT and REFLEXIVE. All rules in a logical router are either stateless or stateful. Mix is not supported. SNAT and DNAT are stateful, and can NOT be supported when the logical router is running at active-active HA mode. The REFLEXIVE action is stateless. The NO_NAT action has no translated_fields, only match fields.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) enable/disable the rule.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Optional) enable/disable the logging of rule.`,
				},
				resource.Attribute{
					Name:        "match_destination_network",
					Description: `(Required for action=DNAT, not allowed for action=REFLEXIVE) IP Address | CIDR. Omitting this field implies Any.`,
				},
				resource.Attribute{
					Name:        "match_source_network",
					Description: `(Required for action=NO_NAT or REFLEXIVE, Optional for the other actions) IP Address | CIDR. Omitting this field implies Any.`,
				},
				resource.Attribute{
					Name:        "nat_pass",
					Description: `(Optional) Enable/disable to bypass following firewall stage. The default is true, meaning that the following firewall stage will be skipped. Please note, if action is NO_NAT, then nat_pass must be set to true or omitted.`,
				},
				resource.Attribute{
					Name:        "translated_network",
					Description: `(Required for action=DNAT or SNAT) IP Address | IP Range | CIDR.`,
				},
				resource.Attribute{
					Name:        "translated_ports",
					Description: `(Optional) port number or port range. Allowed only when action=DNAT.`,
				},
				resource.Attribute{
					Name:        "rule_priority",
					Description: `The priority of the rule which is ascending, valid range [0-2147483647]. If multiple rules have the same priority, evaluation sequence is undefined. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NAT rule.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing NAT rule can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_nat_rule.rule1 logical-router-uuid/nat-rule-num ` + "`" + `` + "`" + `` + "`" + ` The above command imports the NAT rule named ` + "`" + `rule1` + "`" + ` with the number id ` + "`" + `nat-rule-num` + "`" + ` that belongs to the tier 1 logical router with the NSX id ` + "`" + `logical-router-uuid` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NAT rule.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing NAT rule can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_nat_rule.rule1 logical-router-uuid/nat-rule-num ` + "`" + `` + "`" + `` + "`" + ` The above command imports the NAT rule named ` + "`" + `rule1` + "`" + ` with the number id ` + "`" + `nat-rule-num` + "`" + ` that belongs to the tier 1 logical router with the NSX id ` + "`" + `logical-router-uuid` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ns_group",
			Category:         "Manager Resources",
			ShortDescription: `A resource to configure a networking and security group in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ns",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this NS group.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Optional) Reference to the direct/static members of the NSGroup. Can be ID based expressions only. VirtualMachine cannot be added as a static member.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `(Required) Static member type, one of: NSGroup, IPSet, LogicalPort, LogicalSwitch, MACSet`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Member ID`,
				},
				resource.Attribute{
					Name:        "membership_criteria",
					Description: `(Optional) List of tag or ID expressions which define the membership criteria for this NSGroup. An object must satisfy at least one of these expressions to qualify as a member of this group.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `(Required) Dynamic member type, one of: LogicalPort, LogicalSwitch, VirtualMachine.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Tag scope for matching dynamic members.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Tag value for matching dynamic members. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS group.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing networking and security group can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ns_group.group2 UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the networking and security group named ` + "`" + `group2` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS group.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing networking and security group can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ns_group.group2 UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the networking and security group named ` + "`" + `group2` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ns_service_group",
			Category:         "Manager NS Services Resources",
			ShortDescription: `Provides a resource to configure NS service group on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"ns",
				"services",
				"service",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this NS service group.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) List of NSServices IDs that can be added as members to an NSServiceGroup. All members should be of the same L2 type: Ethernet, or Non Ethernet. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS service group.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing ns service group can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ns_service_group.ns_service_group UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the NS service group named ` + "`" + `ns_service_group` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NS service group.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing ns service group can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_ns_service_group.ns_service_group UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the NS service group named ` + "`" + `ns_service_group` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_bgp_neighbor",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure a BGP Neighbor.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"bgp",
				"neighbor",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this resource.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.`,
				},
				resource.Attribute{
					Name:        "bgp_path",
					Description: `(Required) The policy path to the BGP configuration for this neighbor.`,
				},
				resource.Attribute{
					Name:        "allow_as_in",
					Description: `(Optional) Flag to enable allowas_in option for BGP neighbor. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "graceful_restart_mode",
					Description: `(Optional) BGP Graceful Restart Configuration Mode. One of ` + "`" + `DISABLE` + "`" + `, ` + "`" + `GR_AND_HELPER` + "`" + ` or ` + "`" + `HELPER_ONLY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hold_down_time",
					Description: `(Optional) Wait time in seconds before declaring peer dead. Defaults to ` + "`" + `180` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "keep_alive_time",
					Description: `(Optional) Interval between keep alive messages sent to peer. Defaults to ` + "`" + `60` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "maximum_hop_limit",
					Description: `(Optional) Maximum number of hops allowed to reach BGP neighbor. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "neighbor_address",
					Description: `(Required) Neighbor IP Address.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password for BGP neighbor authentication. Set to the empty string to clear out the password.`,
				},
				resource.Attribute{
					Name:        "remote_as_num",
					Description: `(Required) 4 Byte ASN of the neighbor in ASPLAIN Format.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Optional) A list of up to 8 source IP Addresses for BGP peering. ` + "`" + `ip_addresses` + "`" + ` field of an existing ` + "`" + `nsxt_policy_tier0_gateway_interface` + "`" + ` can be used here.`,
				},
				resource.Attribute{
					Name:        "bfd_config",
					Description: `(Optional) The BFD configuration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) A boolean flag to enable/disable BFD. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Time interval between heartbeat packets in milliseconds. Defaults to ` + "`" + `500` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "multiple",
					Description: `(Optional) Number of times heartbeat packet is missed before BFD declares the neighbor is down. Defaults to ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_filtering",
					Description: `(Optional) Up to 2 route filters for the neighbor. Note that prior to NSX version 3.0.0, only 1 element is supported.`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `(Required) Address family type. Must be one of ` + "`" + `EVPN` + "`" + `, ` + "`" + `IPV4` + "`" + ` or ` + "`" + `IPV6` + "`" + `. Note the ` + "`" + `EVPN` + "`" + ` property is only available starting with NSX version 3.0.0.`,
				},
				resource.Attribute{
					Name:        "maximum_routes",
					Description: `(Optional) Maximum number of routes for the address family. Note this property is only available starting with NSX version 3.0.0. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing BGP Neighbor can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_bgp_neighbor.test T0_ID/LOCALE_SERVICE_ID/NEIGHBOR_ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports BGP Neighbor named ` + "`" + `test` + "`" + ` with the NSX BGP Neighbor ID ` + "`" + `NEIGHBOR_ID` + "`" + ` from the Tier-0 ` + "`" + `T0_ID` + "`" + ` and Locale Service ` + "`" + `LOCALE_SERVICE_ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing BGP Neighbor can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_bgp_neighbor.test T0_ID/LOCALE_SERVICE_ID/NEIGHBOR_ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports BGP Neighbor named ` + "`" + `test` + "`" + ` with the NSX BGP Neighbor ID ` + "`" + `NEIGHBOR_ID` + "`" + ` from the Tier-0 ` + "`" + `T0_ID` + "`" + ` and Locale Service ` + "`" + `LOCALE_SERVICE_ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_dhcp_relay",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure a Dhcp Relay.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"dhcp",
				"relay",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this resource.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.`,
				},
				resource.Attribute{
					Name:        "server_addresses",
					Description: `(Required) List of DHCP server addresses. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing object can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_dhcp_relay.test ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports Dhcp Relay named ` + "`" + `test` + "`" + ` with the NSX Dhcp Relay ID ` + "`" + `ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing object can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_dhcp_relay.test ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports Dhcp Relay named ` + "`" + `test` + "`" + ` with the NSX Dhcp Relay ID ` + "`" + `ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_dhcp_server",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure a DHCP Servers in NSX-T.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"dhcp",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this resource.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_path",
					Description: `(Optional) The Policy path to the edge cluster for this DHCP Server.`,
				},
				resource.Attribute{
					Name:        "lease_time",
					Description: `(Optional) IP address lease time in seconds. Valid values from ` + "`" + `60` + "`" + ` to ` + "`" + `4294967295` + "`" + `. Default is ` + "`" + `86400` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "preferred_edge_paths",
					Description: `(Optional) Policy paths to edge nodes. The first edge node is assigned as active edge, and second one as standby edge.`,
				},
				resource.Attribute{
					Name:        "server_addresses",
					Description: `(Optional) DHCP server address in CIDR format. At most 2 supported; one IPv4 and one IPv6 address. Server address can also be specified on segment subnet level. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing DHCP Server can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_dhcp_server.dhcp1 ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports a DHCP Server named ` + "`" + `dhcp1` + "`" + ` with the NSX DHCP Server ID ` + "`" + `ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing DHCP Server can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_dhcp_server.dhcp1 ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports a DHCP Server named ` + "`" + `dhcp1` + "`" + ` with the NSX DHCP Server ID ` + "`" + `ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_gateway_policy",
			Category:         "Policy Resources",
			ShortDescription: `A resource to Gateway security policies.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) The category to use for priority of this Gateway Policy. Must be one of: ` + "`" + `Emergency` + "`" + `, ` + "`" + `SystemRules` + "`" + `, ` + "`" + `SharedPreRules` + "`" + `, ` + "`" + `LocalGatewayRules` + "`" + `, ` + "`" + `AutoServiceRules` + "`" + ` and ` + "`" + `Default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain to use for the Gateway Policy. This domain must already exist. For VMware Cloud on AWS use ` + "`" + `cgw` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this Gateway Policy.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the Gateway Policy resource.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments for this Gateway Policy including lock/unlock comments.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `(Optional) A boolean value indicating if the policy is locked. If locked, no other users can update the resource.`,
				},
				resource.Attribute{
					Name:        "sequence_number",
					Description: `(Optional) An int value used to resolve conflicts between security policies across domains`,
				},
				resource.Attribute{
					Name:        "stateful",
					Description: `(Optional) A boolean value to indicate if this Policy is stateful. When it is stateful, the state of the network connects are tracked and a stateful packet inspection is performed.`,
				},
				resource.Attribute{
					Name:        "tcp_strict",
					Description: `(Optional) A boolean value to enable/disable a 3 way TCP handshake is done before the data packets are sent.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "destination_groups",
					Description: `(Optional) A list of destination group paths to use for the policy.`,
				},
				resource.Attribute{
					Name:        "destinations_excluded",
					Description: `(Optional) A boolean value indicating negation of destination groups.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) The traffic direction for the policy. Must be one of: ` + "`" + `IN` + "`" + `, ` + "`" + `OUT` + "`" + ` or ` + "`" + `IN_OUT` + "`" + `. Defaults to ` + "`" + `IN_OUT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) A boolean value to indicate the rule is disabled. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) The IP Protocol for the rule. Must be one of: ` + "`" + `IPV4` + "`" + `, ` + "`" + `IPV6` + "`" + ` or ` + "`" + `IPV4_IPV6` + "`" + `. Defaults to ` + "`" + `IPV4_IPV6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logged",
					Description: `(Optional) A boolean flag to enable packet logging.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Text for additional notes on changes for the rule.`,
				},
				resource.Attribute{
					Name:        "profiles",
					Description: `(Optional) A list of profiles for the rule.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Required) List of policy paths where the rule is applied.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional) List of services to match.`,
				},
				resource.Attribute{
					Name:        "source_groups",
					Description: `(Optional) A list of source group paths to use for the policy.`,
				},
				resource.Attribute{
					Name:        "source_excluded",
					Description: `(Optional) A boolean value indicating negation of source groups.`,
				},
				resource.Attribute{
					Name:        "log_label",
					Description: `(Optional) Additional information (string) which will be propagated to the rule syslog.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this Rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The action for the Rule. Must be one of: ` + "`" + `ALLOW` + "`" + `, ` + "`" + `DROP` + "`" + ` or ` + "`" + `REJECT` + "`" + `. Defaults to ` + "`" + `ALLOW` + "`" + `. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Secuirty Policy.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "sequence_number",
					Description: `Sequence number of the this rule, is defined by order of rules in the list.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Unique positive number that is assigned by the system and is useful for debugging. ## Importing An existing Gateway Policy can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_gateway_policy.gwpolicy1 ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy Gateway Policy named ` + "`" + `gwpolicy1` + "`" + ` with the NSX Policy id ` + "`" + `ID` + "`" + `. If the Policy to import isn't in the ` + "`" + `default` + "`" + ` domain, the domain name can be added to the ` + "`" + `ID` + "`" + ` before a slash. For example to import a Group with ` + "`" + `ID` + "`" + ` in the ` + "`" + `MyDomain` + "`" + ` domain: ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_gateway_policy.gwpolicy1 MyDomain/ID ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Secuirty Policy.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "sequence_number",
					Description: `Sequence number of the this rule, is defined by order of rules in the list.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Unique positive number that is assigned by the system and is useful for debugging. ## Importing An existing Gateway Policy can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_gateway_policy.gwpolicy1 ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy Gateway Policy named ` + "`" + `gwpolicy1` + "`" + ` with the NSX Policy id ` + "`" + `ID` + "`" + `. If the Policy to import isn't in the ` + "`" + `default` + "`" + ` domain, the domain name can be added to the ` + "`" + `ID` + "`" + ` before a slash. For example to import a Group with ` + "`" + `ID` + "`" + ` in the ` + "`" + `MyDomain` + "`" + ` domain: ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_gateway_policy.gwpolicy1 MyDomain/ID ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_group",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure a Group and its members.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain to use for the Group. This domain must already exist. For VMware Cloud on AWS use ` + "`" + `cgw` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this Group.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the group resource.`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `(Optional) A repeatable block to specify criteria for members of this Group. If more than 1 criteria block is specified, it must be separated by a ` + "`" + `conjunction` + "`" + `. In a ` + "`" + `criteria` + "`" + ` block the following membership selection expressions can be used:`,
				},
				resource.Attribute{
					Name:        "ipaddress_expression",
					Description: `(Optional) An expression block to specify individual IP Addresses, ranges of IP Addresses or subnets for this Group.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(Required for a ` + "`" + `ipaddress_expression` + "`" + `) This list can consist of a single IP address, IP address range or a subnet. Its type can be of either IPv4 or IPv6. Both IPv4 and IPv6 addresses within one expression is not allowed.`,
				},
				resource.Attribute{
					Name:        "path_expression",
					Description: `(Optional) An expression block to specify direct group members by policy path.`,
				},
				resource.Attribute{
					Name:        "member_paths",
					Description: `(Required for a ` + "`" + `path_expression` + "`" + `) List of policy paths for direct members for this Group (such as Segments, Segment ports, Groups etc).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Group.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing policy Group can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_group.group1 ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy Group named ` + "`" + `group` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + `. If the Group to import isn't in the ` + "`" + `default` + "`" + ` domain, the domain name can be added to the ` + "`" + `ID` + "`" + ` before a slash. For example to import a Group with ` + "`" + `ID` + "`" + ` in the ` + "`" + `MyDomain` + "`" + ` domain: ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_group.group1 MyDomain/ID ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Group.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing policy Group can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_group.group1 ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy Group named ` + "`" + `group` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + `. If the Group to import isn't in the ` + "`" + `default` + "`" + ` domain, the domain name can be added to the ` + "`" + `ID` + "`" + ` before a slash. For example to import a Group with ` + "`" + `ID` + "`" + ` in the ` + "`" + `MyDomain` + "`" + ` domain: ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_group.group1 MyDomain/ID ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_ip_address_allocation",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure a IP Address allocations.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"ip",
				"address",
				"allocation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this resource.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.`,
				},
				resource.Attribute{
					Name:        "allocation_ip",
					Description: `(Optional) The IP Address to allocate. If unspecified any free IP in the pool will be allocated.`,
				},
				resource.Attribute{
					Name:        "pool_path",
					Description: `(Required) The policy path to the IP Pool for this Allocation. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Allocation.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "allocation_ip",
					Description: `If the ` + "`" + `allocation_ip` + "`" + ` is not specified in the resource, any free IP is allocated and its value is exported on this attribute. ## Importing An existing IP Allocation can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_ip_address_allocation.test POOL-ID/ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports IpAddressAllocation named ` + "`" + `test` + "`" + ` with the NSX IpAddressAllocation ID ` + "`" + `ID` + "`" + ` in IP Pool ` + "`" + `POOL-ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Allocation.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "allocation_ip",
					Description: `If the ` + "`" + `allocation_ip` + "`" + ` is not specified in the resource, any free IP is allocated and its value is exported on this attribute. ## Importing An existing IP Allocation can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_ip_address_allocation.test POOL-ID/ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports IpAddressAllocation named ` + "`" + `test` + "`" + ` with the NSX IpAddressAllocation ID ` + "`" + `ID` + "`" + ` in IP Pool ` + "`" + `POOL-ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_ip_block",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure IP Blocks in NSX Policy.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"ip",
				"block",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name for the IP Block.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) Network address and the prefix length which will be associated with a layer-2 broadcast domain.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this IP Block. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP Block.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the resource. ## Importing An existing IP Block can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_ip_block.block1 ID ` + "`" + `` + "`" + `` + "`" + ` The above would import NSX IP Block as a resource named ` + "`" + `block1` + "`" + ` with the NSX id ` + "`" + `ID` + "`" + `, where ` + "`" + `ID` + "`" + ` is NSX ID of the IP Block.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP Block.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the resource. ## Importing An existing IP Block can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_ip_block.block1 ID ` + "`" + `` + "`" + `` + "`" + ` The above would import NSX IP Block as a resource named ` + "`" + `block1` + "`" + ` with the NSX id ` + "`" + `ID` + "`" + `, where ` + "`" + `ID` + "`" + ` is NSX ID of the IP Block.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_ip_pool",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure IP Pools in NSX Policy.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"ip",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name for the IP Pool.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this IP Pool. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP Pool.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the resource. ## Importing An existing IP Pool can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_ip_pool.pool1 ID ` + "`" + `` + "`" + `` + "`" + ` The above would import NSX IP Pool as a resource named ` + "`" + `pool1` + "`" + ` with the NSX ID ` + "`" + `ID` + "`" + `, where ` + "`" + `ID` + "`" + ` is NSX ID of the IP Pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the IP Pool.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the resource. ## Importing An existing IP Pool can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_ip_pool.pool1 ID ` + "`" + `` + "`" + `` + "`" + ` The above would import NSX IP Pool as a resource named ` + "`" + `pool1` + "`" + ` with the NSX ID ` + "`" + `ID` + "`" + `, where ` + "`" + `ID` + "`" + ` is NSX ID of the IP Pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_ip_pool_block_subnet",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure IP Pool Block Subnets in NSX Policy.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"ip",
				"pool",
				"block",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name for the Block Subnet.`,
				},
				resource.Attribute{
					Name:        "pool_path",
					Description: `(Required) The Policy path to the IP Pool for this Block Subnet.`,
				},
				resource.Attribute{
					Name:        "block_path",
					Description: `(Required) The Policy path to the IP Block for this Block Subnet.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of this Block Subnet. Must be a power of 2`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this Block Subnet.`,
				},
				resource.Attribute{
					Name:        "auto_assign_gateway",
					Description: `(Optional) A boolean flag to toggle auto-assignment of the Gateway IP for this Subnet ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of this Block Subnet.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the resource. ## Importing An existing Block can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_ip_pool_block_subnet.block_subnet1 pool-id/subnet-id ` + "`" + `` + "`" + `` + "`" + ` The above would import NSX Block Subnet as a resource named ` + "`" + `block_subnet1` + "`" + ` with the NSX ID ` + "`" + `subnet-id` + "`" + ` in the IP Pool ` + "`" + `pool-id` + "`" + `, where ` + "`" + `subnet-id` + "`" + ` is NSX ID of Block Subnet and ` + "`" + `pool-id` + "`" + ` is the IP Pool ID the Subnet is in.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of this Block Subnet.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the resource. ## Importing An existing Block can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_ip_pool_block_subnet.block_subnet1 pool-id/subnet-id ` + "`" + `` + "`" + `` + "`" + ` The above would import NSX Block Subnet as a resource named ` + "`" + `block_subnet1` + "`" + ` with the NSX ID ` + "`" + `subnet-id` + "`" + ` in the IP Pool ` + "`" + `pool-id` + "`" + `, where ` + "`" + `subnet-id` + "`" + ` is NSX ID of Block Subnet and ` + "`" + `pool-id` + "`" + ` is the IP Pool ID the Subnet is in.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_ip_pool_static_subnet",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure IP Pool Static Subnets in NSX Policy.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"ip",
				"pool",
				"static",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name for the Static Subnet.`,
				},
				resource.Attribute{
					Name:        "pool_path",
					Description: `(Required) The Policy path to the IP Pool for this Static Subnet.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The network CIDR`,
				},
				resource.Attribute{
					Name:        "allocation_range",
					Description: `(Required) One or more IP allocation ranges for the Subnet.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) The start IP address for the allocation range.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) The end IP address for the allocation range.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this Static Subnet.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `(Optional) A list of up to 3 DNS nameservers for the Subnet.`,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `(Optional) The DNS suffix for the Subnet.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Optional) The gateway IP for the Subnet. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of this Static Subnet.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the resource. ## Importing An existing Static can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_ip_pool_static_subnet.static_subnet1 pool-id/subnet-id ` + "`" + `` + "`" + `` + "`" + ` The above would import NSX Static Subnet as a resource named ` + "`" + `static_subnet1` + "`" + ` with the NSX ID ` + "`" + `subnet-id` + "`" + ` in the IP Pool ` + "`" + `pool-id` + "`" + `, where ` + "`" + `subnet-id` + "`" + ` is ID of Static Subnet and ` + "`" + `pool-id` + "`" + ` is the IP Pool ID the Subnet is in.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of this Static Subnet.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the resource. ## Importing An existing Static can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_ip_pool_static_subnet.static_subnet1 pool-id/subnet-id ` + "`" + `` + "`" + `` + "`" + ` The above would import NSX Static Subnet as a resource named ` + "`" + `static_subnet1` + "`" + ` with the NSX ID ` + "`" + `subnet-id` + "`" + ` in the IP Pool ` + "`" + `pool-id` + "`" + `, where ` + "`" + `subnet-id` + "`" + ` is ID of Static Subnet and ` + "`" + `pool-id` + "`" + ` is the IP Pool ID the Subnet is in.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_lb_pool",
			Category:         "Policy Load Balancing Resources",
			ShortDescription: `A resource to configure a LBPool.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"load",
				"balancing",
				"lb",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this resource.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Optional) Load balancing algorithm, one of ` + "`" + `ROUND_ROBIN` + "`" + `, ` + "`" + `WEIGHTED_ROUND_ROBIN` + "`" + `, ` + "`" + `LEAST_CONNECTION` + "`" + `, ` + "`" + `WEIGHTED_LEAST_CONNECTION` + "`" + `, ` + "`" + `IP_HASH` + "`" + `. Default is ` + "`" + `ROUND_ROBIN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "member_group",
					Description: `(Optional) Grouping specification for pool members. When ` + "`" + `member_group` + "`" + ` is set, ` + "`" + `member` + "`" + ` should not be specified.`,
				},
				resource.Attribute{
					Name:        "group_path",
					Description: `(Required) Path for policy group.`,
				},
				resource.Attribute{
					Name:        "allow_ipv4",
					Description: `(Optional) Use IPv4 addresses from the grouping object, default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_ipv6",
					Description: `(Optional) Use IPv6 addresses from the grouping object, default is ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "max_ip_list_size",
					Description: `(Optional) Maximum number of IPs to use from the grouping object.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) If port is specified, all connections will be redirected to this port.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) Member IP address.`,
				},
				resource.Attribute{
					Name:        "admin_state",
					Description: `(Optional) One of ` + "`" + `ENABLED` + "`" + `, ` + "`" + `DISABLED` + "`" + `, ` + "`" + `GRACEFUL_DISABLED` + "`" + `. Default is ` + "`" + `ENABLED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_member",
					Description: `(Optional) Whether this member is a backup member.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name of the member.`,
				},
				resource.Attribute{
					Name:        "max_concurrent_connections",
					Description: `(Optional) To ensure members are not overloaded, connections to a member can be capped by this setting.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) If port is specified, all connections will be redirected to this port.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Pool member weight is used for WEIGHTED algorithms.`,
				},
				resource.Attribute{
					Name:        "min_active_members",
					Description: `(Optional) A pool is considered active if there are at least certain minimum number of members.`,
				},
				resource.Attribute{
					Name:        "active_monitor_path",
					Description: `(Optional) Active monitor to be associated with this pool.`,
				},
				resource.Attribute{
					Name:        "passive_monitor_path",
					Description: `(Optional) Passive monitor to be associated with this pool.`,
				},
				resource.Attribute{
					Name:        "snat",
					Description: `(Optional) Source NAT may be required to ensure traffic from the server destined to the client is received by the load balancer.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) SNAT type, one of 'AUTOMAP` + "`" + `, ` + "`" + `DISABLED` + "`" + `, ` + "`" + `IPPOOL` + "`" + `. Default is ` + "`" + `AUTOMAP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_pool_addresses",
					Description: `(Optional) List of IP ranges or IP CIDRs to use for IPPOOL SNAT type.`,
				},
				resource.Attribute{
					Name:        "tcp_multiplexing_enabled",
					Description: `(Optional) Enable TCP multiplexing within the pool.`,
				},
				resource.Attribute{
					Name:        "tcp_multiplexing_number",
					Description: `(Optional) The maximum number of TCP connections per pool that are idly kept alive for sending future client requests. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Secuirty Policy.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing pool can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_lb_pool.test ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports LBPool named ` + "`" + `test` + "`" + ` with the NSX LBPool ID ` + "`" + `ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Secuirty Policy.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing pool can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_lb_pool.test ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports LBPool named ` + "`" + `test` + "`" + ` with the NSX LBPool ID ` + "`" + `ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_lb_service",
			Category:         "Policy Load Balancing Resources",
			ShortDescription: `A resource to configure a Load Balancer Service.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"load",
				"balancing",
				"lb",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Load Balancer Service size, one of ` + "`" + `SMALL` + "`" + `, ` + "`" + `MEDIUM` + "`" + `, ` + "`" + `LARGE` + "`" + `, ` + "`" + `XLARGE` + "`" + `, ` + "`" + `DLB` + "`" + `. Default is ` + "`" + `SMALL` + "`" + `. Please note that XLARGE is only supported since NSX 3.0.0`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this resource.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.`,
				},
				resource.Attribute{
					Name:        "connectivity_path",
					Description: `(Optional) Tier1 Gateway where this service is instantiated. In future, other objects will be supported.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Flag to enable the service.`,
				},
				resource.Attribute{
					Name:        "error_log_level",
					Description: `(Optional) Log level for the service, one of ` + "`" + `DEBUG` + "`" + `, ` + "`" + `INFO` + "`" + `, ` + "`" + `WARNING` + "`" + `, ` + "`" + `ERROR` + "`" + `, ` + "`" + `CRITICAL` + "`" + `, ` + "`" + `ALERT` + "`" + `, ` + "`" + `EMERGENCY` + "`" + `. Default is ` + "`" + `INFO` + "`" + `. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_lb_service.test ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports LBService named ` + "`" + `test` + "`" + ` with the NSX Load Balancer Service ID ` + "`" + `ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_lb_service.test ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports LBService named ` + "`" + `test` + "`" + ` with the NSX Load Balancer Service ID ` + "`" + `ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_lb_virtual_server",
			Category:         "Policy Load Balancing Resources",
			ShortDescription: `A resource to configure a Load Balancer Virtual Server.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"load",
				"balancing",
				"lb",
				"virtual",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this resource.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.`,
				},
				resource.Attribute{
					Name:        "application_profile_path",
					Description: `(Required) Application profile path for this virtual server.`,
				},
				resource.Attribute{
					Name:        "access_log_enabled",
					Description: `(Optional) If set, all connections/requests sent to the virtual server are logged to access log.`,
				},
				resource.Attribute{
					Name:        "default_pool_member_ports",
					Description: `(Optional) Default pool member ports to use when member port is not defined on the pool.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Flag to enable this Virtual Server.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) Virtual Server IP address.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Required) Virtual Server Ports.`,
				},
				resource.Attribute{
					Name:        "persistence_profile_path",
					Description: `(Optional) Path to persistance profile allowing related client connections to be sent to the same backend server.`,
				},
				resource.Attribute{
					Name:        "service_path",
					Description: `(Optional) Virtual Server can be associated with Load Balancer Service.`,
				},
				resource.Attribute{
					Name:        "max_concurrent_connections",
					Description: `(Optional) To ensure one virtual server does not over consume resources, connections to Virtual Server can be capped.`,
				},
				resource.Attribute{
					Name:        "max_new_connection_rate",
					Description: `(Optional) To ensure one virtual server does not over consume resources, connections to a member can be rate limited.`,
				},
				resource.Attribute{
					Name:        "pool_path",
					Description: `(Optional)Path for Load Balancer Pool.`,
				},
				resource.Attribute{
					Name:        "sorry_pool_path",
					Description: `(Optional) When load balancer can not select server in pool, the request would be served by sorry pool`,
				},
				resource.Attribute{
					Name:        "server_ssl",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "server_auth",
					Description: `(Optional) Server Authentication Mode, one of ` + "`" + `REQUIRED` + "`" + `, ` + "`" + `IGNORE` + "`" + `, ` + "`" + `AUTO_APPLY` + "`" + `. Default is ` + "`" + `AUTO_APPLY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_chain_depth",
					Description: `(Optional) Allowed certificate chain depth.`,
				},
				resource.Attribute{
					Name:        "client_certificate_path",
					Description: `(Optional) Client certificat path for client authentication against the server.`,
				},
				resource.Attribute{
					Name:        "ca_paths",
					Description: `(Optional) If server auth type is REQUIRED, client certificate must be signed by one Certificate Authorities provided here.`,
				},
				resource.Attribute{
					Name:        "crl_paths",
					Description: `(Optional) Certificate Revocation Lists can be specified to disallow compromised certificate.`,
				},
				resource.Attribute{
					Name:        "ssl_profile_path",
					Description: `(Optional) Server SSL profile path.`,
				},
				resource.Attribute{
					Name:        "client_ssl",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "client_auth",
					Description: `(Optional) Client Authentication Mode, one of ` + "`" + `REQUIRED` + "`" + `, ` + "`" + `IGNORE` + "`" + `. Default is ` + "`" + `IGNORE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_chain_depth",
					Description: `(Optional) Allowed certificate chain depth.`,
				},
				resource.Attribute{
					Name:        "ca_paths",
					Description: `(Optional) If client auth type is REQUIRED, client certificate must be signed by one Certificate Authorities provided here.`,
				},
				resource.Attribute{
					Name:        "crl_paths",
					Description: `(Optional) Certificate Revocation Lists can be specified to disallow compromised client certificate.`,
				},
				resource.Attribute{
					Name:        "default_certificate_path",
					Description: `(Optional) Default Certificate Path. Must be specified if client_auth is set to ` + "`" + `REQUIRED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sni_paths",
					Description: `(Optional) This setting allows multiple certificates(for different hostnames) to be bound to the same virtual server.`,
				},
				resource.Attribute{
					Name:        "ssl_profile_path",
					Description: `(Optional) Client SSL profile path.`,
				},
				resource.Attribute{
					Name:        "log_significant_event_only",
					Description: `(Optional) If true, significant events are logged in access log. This flag is supported since NSX 3.0.0.`,
				},
				resource.Attribute{
					Name:        "access_list_control",
					Description: `(Optional) Specifies the access list control to define how to filter client connections.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action for connections matching the grouping object.`,
				},
				resource.Attribute{
					Name:        "group_path",
					Description: `(Required) The path of grouping object which defines the IP addresses or ranges to match client IP.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Indicates whether to enable access list control option. Default is true. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing Virtual Server can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_lb_virtual_server.test ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports Load Balancer Virtual Server named ` + "`" + `test` + "`" + ` with the NSX Load Balancer Virtual Server ID ` + "`" + `ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing Virtual Server can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_lb_virtual_server.test ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports Load Balancer Virtual Server named ` + "`" + `test` + "`" + ` with the NSX Load Balancer Virtual Server ID ` + "`" + `ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_nat_rule",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure NAT Ruels in NSX Policy manager.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"nat",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this NAT Rule.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the policy resource.`,
				},
				resource.Attribute{
					Name:        "gateway_path",
					Description: `(Required) The NSX Policy path to the Tier0 or Tier1 Gateway for this NAT Rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action for the NAT Rule. One of ` + "`" + `SNAT` + "`" + `, ` + "`" + `DNAT` + "`" + `, ` + "`" + `REFLEXIVE` + "`" + `, ` + "`" + `NO_SNAT` + "`" + `, ` + "`" + `NO_DNAT` + "`" + `, ` + "`" + `NAT64` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_networks",
					Description: `(Optional) A list of destination network IP addresses or CIDR.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable/disable the Rule. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "firewall_match",
					Description: `(Optional) Firewall match flag. One of ` + "`" + `MATCH_EXTERNAL_ADDRESS` + "`" + `, ` + "`" + `MATCH_INTERNAL_ADDRESS` + "`" + `, ` + "`" + `BYPASS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Optional) Enable/disable rule logging. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule_priority",
					Description: `(Optional) The priority of the rule. Valid values between 0 to 2147483647. Defaults to ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Policy path of Service on which the NAT rule will be applied.`,
				},
				resource.Attribute{
					Name:        "source_networks",
					Description: `(Optional) A list of source network IP addresses or CIDR.`,
				},
				resource.Attribute{
					Name:        "translated_networks",
					Description: `(Optional) A list of translated network IP addresses or CIDR.`,
				},
				resource.Attribute{
					Name:        "translated_ports",
					Description: `(Optional) Port number or port range. For use with ` + "`" + `DNAT` + "`" + ` action only.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) A list of paths to interfaces and/or labels where the NAT Rule is enforced. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing policy NAT Rule can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_nat_rule.rule1 GWID/ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy NAT Rule named ` + "`" + `rule1` + "`" + ` for the NSX Tier0 or Tier1 Gateway ` + "`" + `GWID` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing policy NAT Rule can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_nat_rule.rule1 GWID/ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy NAT Rule named ` + "`" + `rule1` + "`" + ` for the NSX Tier0 or Tier1 Gateway ` + "`" + `GWID` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_security_policy",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure a Security Group and its rules.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"security",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain to use for the resource. This domain must already exist. For VMware Cloud on AWS use ` + "`" + `cgw` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this policy.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) Category of this policy, one of ` + "`" + `Ethernet` + "`" + `, ` + "`" + `Emergency` + "`" + `, ` + "`" + `Infrastructure` + "`" + `, ` + "`" + `Environment` + "`" + `, ` + "`" + `Application` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `(Optional) Comments for security policy lock/unlock.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `(Optional) Indicates whether a security policy should be locked. If locked by a user, no other user would be able to modify this policy.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The list of policy object paths where the rules in this policy will get applied.`,
				},
				resource.Attribute{
					Name:        "sequence_number",
					Description: `(Optional) This field is used to resolve conflicts between security policies across domains.`,
				},
				resource.Attribute{
					Name:        "stateful",
					Description: `(Optional) If true, state of the network connects are tracked and a stateful packet inspection is performed. Default is true.`,
				},
				resource.Attribute{
					Name:        "tcp_strict",
					Description: `(Optional) Ensures that a 3 way TCP handshake is done before the data packets are sent. Default is false.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) A repeatable block to specify rules for the Security Policy. Each rule includes the following fields:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Rule action, one of ` + "`" + `ALLOW` + "`" + `, ` + "`" + `DROP` + "`" + `, ` + "`" + `REJECT` + "`" + `. Default is ` + "`" + `ALLOW` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination_groups",
					Description: `(Optional) Set of group paths that serve as destination for this rule.`,
				},
				resource.Attribute{
					Name:        "source_groups",
					Description: `(Optional) Set of group paths that serve as source for this rule.`,
				},
				resource.Attribute{
					Name:        "destinations_excluded",
					Description: `(Optional) Negation of destination groups.`,
				},
				resource.Attribute{
					Name:        "sources_excluded",
					Description: `(Optional) Negation of source groups.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Traffic direction, one of ` + "`" + `IN` + "`" + `, ` + "`" + `OUT` + "`" + ` or ` + "`" + `IN_OUT` + "`" + `. Default is ` + "`" + `IN_OUT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Flag to disable this rule. Default is false.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) Version of IP protocol, one of ` + "`" + `IPV4` + "`" + `, ` + "`" + `IPV6` + "`" + `, ` + "`" + `IPV4_IPV6` + "`" + `. Default is ` + "`" + `IPV4_IPV6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logged",
					Description: `(Optional) Flag to enable packet logging. Default is false.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Additional notes on changes.`,
				},
				resource.Attribute{
					Name:        "profiles",
					Description: `(Optional) Set of profile paths relevant for this rule.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Set of policy object paths where the rule is applied.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional) Set of service paths to match.`,
				},
				resource.Attribute{
					Name:        "log_label",
					Description: `(Optional) Additional information (string) which will be propagated to the rule syslog.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this Rule. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Security Policy.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "sequence_number",
					Description: `Sequence number of the this rule, is defined by order of rules in the list.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Unique positive number that is assigned by the system and is useful for debugging. ## Importing An existing security policy can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_security_policy.policy1 domain/ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the security policy named ` + "`" + `policy1` + "`" + ` under NSX domain ` + "`" + `domain` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Security Policy.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "sequence_number",
					Description: `Sequence number of the this rule, is defined by order of rules in the list.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Unique positive number that is assigned by the system and is useful for debugging. ## Importing An existing security policy can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_security_policy.policy1 domain/ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the security policy named ` + "`" + `policy1` + "`" + ` under NSX domain ` + "`" + `domain` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_segment",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure a network Segment.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"segment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this policy.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.`,
				},
				resource.Attribute{
					Name:        "connectivity_path",
					Description: `(Optional) Policy path to the connecting Tier-0 or Tier-1.`,
				},
				resource.Attribute{
					Name:        "overlay_id",
					Description: `(Optional) Overlay connectivity ID for this Segment.`,
				},
				resource.Attribute{
					Name:        "transport_zone_path",
					Description: `(Required) Policy path to the Overlay transport zone. This property is required if more than one overlay transport zone is defined, and none is marked as default.`,
				},
				resource.Attribute{
					Name:        "dhcp_config_path",
					Description: `(Optional) Policy path to DHCP server or relay configuration to use for subnets configured on this segment. This attribute is supported with NSX 3.0.0 onwards.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) Subnet configuration block.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) Gateway IP address CIDR. This argument can not be changed if DHCP is enabled for the subnet.`,
				},
				resource.Attribute{
					Name:        "dhcp_ranges",
					Description: `(Optional) List of DHCP address ranges for dynamic IP allocation.`,
				},
				resource.Attribute{
					Name:        "dhcp_v4_config",
					Description: `(Optional) DHCPv4 config for IPv4 subnet. This clause is supported with NSX 3.0.0 onwards.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Optional) IP address of the DHCP server in CIDR format. This attribute is required if segment has provided dhcp_config_path and it represents a DHCP server config.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `(Optional) List of IP addresses of DNS servers for the subnet.`,
				},
				resource.Attribute{
					Name:        "dhcp_option_121",
					Description: `(Optional) DHCP classless static routes.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Destination in cidr format.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required) IP address of next hop.`,
				},
				resource.Attribute{
					Name:        "dhcp_generic_option",
					Description: `(Optional) Generic DHCP options.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(Required) DHCP option code. Valid values are from 0 to 255.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) List of DHCP option values.`,
				},
				resource.Attribute{
					Name:        "dhcp_v6_config",
					Description: `(Optional) DHCPv6 config for IPv6 subnet. This clause is supported with NSX 3.0.0 onwards.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Optional) IP address of the DHCP server in CIDR format. This attribute is required if segment has provided dhcp_config_path and it represents a DHCP server config.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `(Optional) List of IP addresses of DNS servers for the subnet.`,
				},
				resource.Attribute{
					Name:        "preferred_time",
					Description: `(Optional) The time interval in seconds, in which the prefix is advertised as preferred.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(Optional) List of domain names for this subnet.`,
				},
				resource.Attribute{
					Name:        "excluded_range",
					Description: `(Optional) List of excluded address ranges to define dynamic ip allocation ranges.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) IPv6 address that marks beginning of the range.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) IPv6 address that marks end of the range.`,
				},
				resource.Attribute{
					Name:        "sntp_servers",
					Description: `(Optional) IPv6 address of SNTP servers for the subnet.`,
				},
				resource.Attribute{
					Name:        "l2_extension",
					Description: `(Optional) Configuration for extending Segment through L2 VPN.`,
				},
				resource.Attribute{
					Name:        "l2vpn_paths",
					Description: `(Optional) Policy paths of associated L2 VPN sessions.`,
				},
				resource.Attribute{
					Name:        "tunnel_id",
					Description: `(Optional) The Tunnel ID that's a int value between 1 and 4093.`,
				},
				resource.Attribute{
					Name:        "advanced_config",
					Description: `(Optional) Advanced Segment configuration.`,
				},
				resource.Attribute{
					Name:        "address_pool_paths",
					Description: `(Optional) List of Policy path to IP address pools.`,
				},
				resource.Attribute{
					Name:        "connectivity",
					Description: `(Optional) Connectivity configuration to manually connect (ON) or disconnect (OFF).`,
				},
				resource.Attribute{
					Name:        "hybrid",
					Description: `(Optional) Boolean flag to identify a hybrid logical switch.`,
				},
				resource.Attribute{
					Name:        "local_egress",
					Description: `(Optional) Boolean flag to enable local egress when used in conjunction with L2VPN.`,
				},
				resource.Attribute{
					Name:        "uplink_teaming_policy",
					Description: `(Optional) The name of the switching uplink teaming policy for the bridge endpoint. This name corresponds to one of the switching uplink teaming policy names listed in the transport zone. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Secuirty Policy.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Secuirty Policy.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
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
			Category:         "Policy Resources",
			ShortDescription: `A resource that can be used to configure a networking and security service in NSX Policy.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this resource.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the policy resource. The service must contain at least 1 entry (of at least one of the types), and possibly more.`,
				},
				resource.Attribute{
					Name:        "icmp_entry",
					Description: `(Optional) Set of ICMP type service entries. Each with the following attributes:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name of the service entry.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the service entry.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Version of ICMP protocol ICMPv4 or ICMPv6.`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `(Optional) ICMP message code.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `(Optional) ICMP message type.`,
				},
				resource.Attribute{
					Name:        "l4_port_set_entry",
					Description: `(Optional) Set of L4 ports set service entries. Each with the following attributes:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name of the service entry.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the service entry.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) L4 protocol. Accepted values - 'TCP' or 'UDP'.`,
				},
				resource.Attribute{
					Name:        "destination_ports",
					Description: `(Optional) Set of destination ports.`,
				},
				resource.Attribute{
					Name:        "source_ports",
					Description: `(Optional) Set of source ports.`,
				},
				resource.Attribute{
					Name:        "igmp_entry",
					Description: `(Optional) Set of IGMP type service entries. Each with the following attributes:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name of the service entry.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the service entry.`,
				},
				resource.Attribute{
					Name:        "ether_type_entry",
					Description: `(Optional) Set of Ether type service entries. Each with the following attributes:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name of the service entry.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the service entry.`,
				},
				resource.Attribute{
					Name:        "ether_type",
					Description: `(Required) Type of the encapsulated protocol.`,
				},
				resource.Attribute{
					Name:        "ip_protocol_entry",
					Description: `(Optional) Set of IP Protocol type service entries. Each with the following attributes:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name of the service entry.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the service entry.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) IP protocol number.`,
				},
				resource.Attribute{
					Name:        "algorithm_entry",
					Description: `(Optional) Set of Algorithm type service entries. Each with the following attributes:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name of the service entry.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the service entry.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Required) a single destination port.`,
				},
				resource.Attribute{
					Name:        "source_ports",
					Description: `(Optional) Set of source ports/ranges.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Required) Algorithm one of "ORACLE_TNS", "FTP", "SUN_RPC_TCP", "SUN_RPC_UDP", "MS_RPC_TCP", "MS_RPC_UDP", "NBNS_BROADCAST", "NBDG_BROADCAST", "TFTP" ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the service.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_service.service_icmp ID ` + "`" + `` + "`" + `` + "`" + ` The above service imports the service named ` + "`" + `service_icmp` + "`" + ` with the NSX ID ` + "`" + `ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the service.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing service can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_service.service_icmp ID ` + "`" + `` + "`" + `` + "`" + ` The above service imports the service named ` + "`" + `service_icmp` + "`" + ` with the NSX ID ` + "`" + `ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_static_route",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure Static Routes in NSX Policy manager.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"static",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this Tier-0 gateway.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the policy resource.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) The network address in CIDR format for the route.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required) One or more next hops for the static route.`,
				},
				resource.Attribute{
					Name:        "admin_distance",
					Description: `(Optional) The cost associated with the next hop. Valid values are 1 - 255 and the default is 1.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The gateway address of the next hop.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) The policy path to the interface associated with the static route. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing policy Static Route can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_static_route.route1 GWID/ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy Static Route named ` + "`" + `route1` + "`" + ` for the NSX Tier0 or Tier1 Gateway ` + "`" + `GWID` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing policy Static Route can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_static_route.route1 GWID/ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy Static Route named ` + "`" + `route1` + "`" + ` for the NSX Tier0 or Tier1 Gateway ` + "`" + `GWID` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_tier0_gateway",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure a Tier-0 gateway in NSX Policy manager.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"tier0",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this Tier-0 gateway.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the policy resource.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_path",
					Description: `(Optional) The path of the edge cluster where the Tier-0 is placed. Must be specified when ` + "`" + `bgp_config` + "`" + ` is enabled.`,
				},
				resource.Attribute{
					Name:        "failover_mode",
					Description: `(Optional) This failover mode determines, whether the preferred service router instance for given logical router will preempt the peer. Accepted values are PREEMPTIVE/NON_PREEMPTIVE.`,
				},
				resource.Attribute{
					Name:        "default_rule_logging",
					Description: `(Optional) Boolean flag indicating if the default rule logging will be enabled or not. The default value is false.`,
				},
				resource.Attribute{
					Name:        "enable_firewall",
					Description: `(Optional) Boolean flag indicating if the edge firewall will be enabled or not. The default value is true.`,
				},
				resource.Attribute{
					Name:        "force_whitelisting",
					Description: `(Optional) Boolean flag indicating if white-listing will be forced or not. The default value is false.`,
				},
				resource.Attribute{
					Name:        "ipv6_ndra_profile_path",
					Description: `(Optional) Policy path to IPv6 NDRA profile.`,
				},
				resource.Attribute{
					Name:        "ipv6_dad_profile_path",
					Description: `(Optional) Policy path to IPv6 DAD profile.`,
				},
				resource.Attribute{
					Name:        "ha_mode",
					Description: `(Optional) High-availability Mode for Tier-0. Valid values are ` + "`" + `ACTIVE_ACTIVE` + "`" + ` and ` + "`" + `ACTIVE_STANDBY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internal_transit_subnets",
					Description: `(Optional) Internal transit subnets in CIDR format. At most 1 CIDR.`,
				},
				resource.Attribute{
					Name:        "transit_subnets",
					Description: `(Optional) Transit subnets in CIDR format.`,
				},
				resource.Attribute{
					Name:        "dhcp_config_path",
					Description: `(Optional) Policy path to DHCP server or relay configuration to use for this gateway.`,
				},
				resource.Attribute{
					Name:        "bgp_config",
					Description: `(Optional) The BGP configuration for the Tier-0 gateway. When enabled a valid ` + "`" + `edge_cluster_path` + "`" + ` must be set on the Tier-0 gateway.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this Tier-0 gateway's BGP configuration.`,
				},
				resource.Attribute{
					Name:        "ecmp",
					Description: `(Optional) A boolean flag to enable/disable ECMP. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) A boolean flag to enable/disable BGP. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "inter_sr_ibgp",
					Description: `(Optional) A boolean flag to enable/disable inter SR IBGP configuration. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_as_num",
					Description: `(Optional) BGP AS number in ASPLAIN/ASDOT Format. Default is ` + "`" + `65000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "multipath_relax",
					Description: `(Optional) A boolean flag to enable/disable multipath relax for BGP. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "graceful_restart_mode",
					Description: `(Optional) Setting to control BGP graceful restart mode, one of ` + "`" + `DISABLE` + "`" + `, ` + "`" + `GR_AND_HELPER` + "`" + `, ` + "`" + `HELPER_ONLY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "graceful_restart_timer",
					Description: `(Optional) BGP graceful restart timer. Default is ` + "`" + `180` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "graceful_restart_stale_route_timer",
					Description: `(Optional) BGP stale route timer. Default is ` + "`" + `600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) CIDR of aggregate address.`,
				},
				resource.Attribute{
					Name:        "summary_only",
					Description: `(Optional) A boolean flag to enable/disable summarized route info. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vrf_config",
					Description: `(Optional) VRF config for VRF Tier0. This clause is supported with NSX 3.0.0 onwards.`,
				},
				resource.Attribute{
					Name:        "gateway_path",
					Description: `(Required) Default Tier0 path. Cannot be modified after realization.`,
				},
				resource.Attribute{
					Name:        "evpn_transit_vni",
					Description: `(Optional) L3 VNI associated with the VRF for overlay traffic. VNI must be unique and belong to configured VNI pool.`,
				},
				resource.Attribute{
					Name:        "route_distinguisher",
					Description: `(Optional) Route distinguisher. Format: <ASN>:<number> or <IPAddress>:<number>.`,
				},
				resource.Attribute{
					Name:        "route_target",
					Description: `(Optional) Only one target is supported.`,
				},
				resource.Attribute{
					Name:        "auto_mode",
					Description: `(Optional) When true, import and export targets should not be specified.`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `(Optional) Address family, currently only ` + "`" + `L2VPN_EVPN` + "`" + ` is supported, which is the default.`,
				},
				resource.Attribute{
					Name:        "import_targets",
					Description: `(Optional) List of import route targets. Format: <ASN>:<number>.`,
				},
				resource.Attribute{
					Name:        "export_targets",
					Description: `(Optional) List of export route targets. Format: <ASN>:<number>. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Tier-0 gateway.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "bgp_config",
					Description: `The following attributes are exported for ` + "`" + `bgp_config` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing policy Tier-0 gateway can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_tier0_gateway.tier0_gw ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy Tier-0 gateway named ` + "`" + `tier0_gw` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Tier-0 gateway.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "bgp_config",
					Description: `The following attributes are exported for ` + "`" + `bgp_config` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing policy Tier-0 gateway can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_tier0_gateway.tier0_gw ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy Tier-0 gateway named ` + "`" + `tier0_gw` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_tier0_gateway_interface",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure an Interface on Tier-0 gateway in NSX Policy manager.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"tier0",
				"gateway",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this resource.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the policy resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of this interface, one of ` + "`" + `SERVICE` + "`" + `, ` + "`" + `EXTERNAL` + "`" + `, ` + "`" + `LOOPBACK` + "`" + `. Default is ` + "`" + `EXTERNAL` + "`" + ``,
				},
				resource.Attribute{
					Name:        "gateway_path",
					Description: `(Required) Policy path for the Tier-0 Gateway.`,
				},
				resource.Attribute{
					Name:        "segment_path",
					Description: `(Optional) Policy path for segment to be connected with this Tier1 Gateway. This argemnt is required for interfaces of type ` + "`" + `SERVICE` + "`" + ` and ` + "`" + `EXTERNAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Required) list of Ip Addresses/Prefixes in CIDR format, to be associated with this interface.`,
				},
				resource.Attribute{
					Name:        "edge_node_path",
					Description: `(Optional) Path of edge node for this interface, relevant for interfaces of type ` + "`" + `EXTERNAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) Maximum Transmission Unit for this interface.`,
				},
				resource.Attribute{
					Name:        "ipv6_ndra_profile_path",
					Description: `(Optional) IPv6 NDRA profile to be associated with this interface.`,
				},
				resource.Attribute{
					Name:        "enable_pim",
					Description: `(Optional) Flag to enable Protocol Independent Multicast, relevant only for interfaces of type ` + "`" + `EXTERNAL` + "`" + `. This attribute is supported with NSX 3.0.0 onwards.`,
				},
				resource.Attribute{
					Name:        "urpf_mode",
					Description: `(Optional) Unicast Reverse Path Forwarding mode, one of ` + "`" + `NONE` + "`" + `, ` + "`" + `STRICT` + "`" + `. Default is ` + "`" + `STRICT` + "`" + `. This attribute is supported with NSX 3.0.0 onwards. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `list of Ip Addresses picked from each subnet in ` + "`" + `subnets` + "`" + ` field. This attribute can serve as ` + "`" + `source_addresses` + "`" + ` field of ` + "`" + `nsxt_policy_bgp_neighbor` + "`" + ` resource. ## Importing An existing policy Tier-0 Gateway Interface can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_tier0_gateway_interface.interface1 GW-ID/LOCALE-SERVICE-ID/ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy Tier-0 gateway interface named ` + "`" + `interface1` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + ` on Tier0 Gateway ` + "`" + `GW-ID` + "`" + `, under locale service ` + "`" + `LOCALE-SERVICE-ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `list of Ip Addresses picked from each subnet in ` + "`" + `subnets` + "`" + ` field. This attribute can serve as ` + "`" + `source_addresses` + "`" + ` field of ` + "`" + `nsxt_policy_bgp_neighbor` + "`" + ` resource. ## Importing An existing policy Tier-0 Gateway Interface can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_tier0_gateway_interface.interface1 GW-ID/LOCALE-SERVICE-ID/ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy Tier-0 gateway interface named ` + "`" + `interface1` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + ` on Tier0 Gateway ` + "`" + `GW-ID` + "`" + `, under locale service ` + "`" + `LOCALE-SERVICE-ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_tier1_gateway",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure a Tier-1 gateway in NSX Policy manager.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"tier1",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this Tier-1 gateway.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the policy resource.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_path",
					Description: `(Optional) The path of the edge cluster where the Tier-1 is placed.`,
				},
				resource.Attribute{
					Name:        "failover_mode",
					Description: `(Optional) This failover mode determines, whether the preferred service router instance for given logical router will preempt the peer. Accepted values are PREEMPTIVE/NON_PREEMPTIVE.`,
				},
				resource.Attribute{
					Name:        "default_rule_logging",
					Description: `(Optional) Boolean flag indicating if the default rule logging will be enabled or not. The default value is false.`,
				},
				resource.Attribute{
					Name:        "enable_firewall",
					Description: `(Optional) Boolean flag indicating if the edge firewall will be enabled or not. The default value is true.`,
				},
				resource.Attribute{
					Name:        "enable_standby_relocation",
					Description: `(Optional) Boolean flag indicating if the standby relocation will be enabled or not. The default value is false.`,
				},
				resource.Attribute{
					Name:        "force_whitelisting",
					Description: `(Optional) Boolean flag indicating if white-listing will be forced or not. The default value is false. Setting it to ` + "`" + `true` + "`" + ` will create a base deny entry rule on Tier-1 firewall.`,
				},
				resource.Attribute{
					Name:        "tier0_path",
					Description: `(Optional) The path of the connected Tier0.`,
				},
				resource.Attribute{
					Name:        "route_advertisement_types",
					Description: `(Optional) Enable different types of route advertisements: TIER1_STATIC_ROUTES, TIER1_CONNECTED, TIER1_NAT, TIER1_LB_VIP, TIER1_LB_SNAT, TIER1_DNS_FORWARDER_IP, TIER1_IPSEC_LOCAL_ENDPOINT.`,
				},
				resource.Attribute{
					Name:        "ipv6_ndra_profile_path",
					Description: `(Optional) Policy path to IPv6 NDRA profile.`,
				},
				resource.Attribute{
					Name:        "ipv6_dad_profile_path",
					Description: `(Optional) Policy path to IPv6 DAD profile.`,
				},
				resource.Attribute{
					Name:        "dhcp_config_path",
					Description: `(Optional) Policy path to DHCP server or relay configuration to use for this gateway.`,
				},
				resource.Attribute{
					Name:        "pool_allocation",
					Description: `(Optional) Size of edge node allocation at for routing and load balancer service to meet performance and scalability requirements, one of ` + "`" + `ROUTING` + "`" + `, ` + "`" + `LB_SMALL` + "`" + `, ` + "`" + `LB_MEDIUM` + "`" + `, ` + "`" + `LB_LARGE` + "`" + `, ` + "`" + `LB_XLARGE` + "`" + `. Default is ` + "`" + `ROUTING` + "`" + `. Changing this attribute would force new resource.`,
				},
				resource.Attribute{
					Name:        "route_advertisement_rule",
					Description: `(Optional) List of rules for routes advertisement:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action to advertise filtered routes to the connected Tier0 gateway. PERMIT (which is the default): Enables the advertisement, DENY: Disables the advertisement.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Required) list of network CIDRs to be routed.`,
				},
				resource.Attribute{
					Name:        "prefix_operator",
					Description: `(Optional) Prefix operator to apply on subnets. GE prefix operator (which is the default|) filters all the routes having network subset of any of the networks configured in Advertise rule. EQ prefix operator filter all the routes having network equal to any of the network configured in Advertise rule.The name of the rule.`,
				},
				resource.Attribute{
					Name:        "route_advertisement_types",
					Description: `(Optional) List of desired types of route advertisements, supported values: ` + "`" + `TIER1_STATIC_ROUTES` + "`" + `, ` + "`" + `TIER1_CONNECTED` + "`" + `, ` + "`" + `TIER1_NAT` + "`" + `, ` + "`" + `TIER1_LB_VIP` + "`" + `, ` + "`" + `TIER1_LB_SNAT` + "`" + `, ` + "`" + `TIER1_DNS_FORWARDER_IP` + "`" + `, ` + "`" + `TIER1_IPSEC_LOCAL_ENDPOINT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ingress_qos_profile_path",
					Description: `(Optional) QoS Profile path for ingress traffic on link connected to Tier0 gateway.`,
				},
				resource.Attribute{
					Name:        "egress_qos_profile_path",
					Description: `(Optional) QoS Profile path for egress traffic on link connected to Tier0 gateway. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Tier-1 gateway.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing policy Tier-1 gateway can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_tier1_gateway.tier1_gw ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy Tier-1 gateway named ` + "`" + `tier1_gw` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Tier-1 gateway.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing policy Tier-1 gateway can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_tier1_gateway.tier1_gw ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy Tier-1 gateway named ` + "`" + `tier1_gw` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_tier1_gateway_interface",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure an Interface on Tier-1 gateway in NSX Policy manager.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"tier1",
				"gateway",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this resource.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the policy resource.`,
				},
				resource.Attribute{
					Name:        "gateway_path",
					Description: `(Required) Policy path for the Tier-1 Gateway.`,
				},
				resource.Attribute{
					Name:        "segment_path",
					Description: `(Required) Policy path for segment to be connected with this Tier1 Gateway.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Required) list of Ip Addresses/Prefixes in CIDR format, to be associated with this interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) Maximum Transmission Unit for this interface.`,
				},
				resource.Attribute{
					Name:        "ipv6_ndra_profile_path",
					Description: `(Optional) IPv6 NDRA profile to be associated with this interface.`,
				},
				resource.Attribute{
					Name:        "urpf_mode",
					Description: `(Optional) Unicast Reverse Path Forwarding mode, one of ` + "`" + `NONE` + "`" + `, ` + "`" + `STRICT` + "`" + `. Default is ` + "`" + `STRICT` + "`" + `. This attribute is supported with NSX 3.0.0 onwards. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing policy Tier-1 Gateway Interface can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_tier1_gateway_interface.interface1 GW-ID/LOCALE-SERVICE-ID/ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy Tier-1 gateway interface named ` + "`" + `interface1` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + ` on Tier1 Gateway ` + "`" + `GW-ID` + "`" + `, under locale service ` + "`" + `LOCALE-SERVICE-ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing policy Tier-1 Gateway Interface can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_tier1_gateway_interface.interface1 GW-ID/LOCALE-SERVICE-ID/ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the policy Tier-1 gateway interface named ` + "`" + `interface1` + "`" + ` with the NSX Policy ID ` + "`" + `ID` + "`" + ` on Tier1 Gateway ` + "`" + `GW-ID` + "`" + `, under locale service ` + "`" + `LOCALE-SERVICE-ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_vlan_segment",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure a VLAN backed network Segment.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"vlan",
				"segment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this policy.`,
				},
				resource.Attribute{
					Name:        "nsx_id",
					Description: `(Optional) The NSX ID of this resource. If set, this ID will be used to create the resource.`,
				},
				resource.Attribute{
					Name:        "transport_zone_path",
					Description: `(Required) Policy path to the VLAN backed transport zone.`,
				},
				resource.Attribute{
					Name:        "vlan_ids",
					Description: `(Optional) List of VLAN IDs or VLAN ranges.`,
				},
				resource.Attribute{
					Name:        "dhcp_config_path",
					Description: `(Optional) Policy path to DHCP server or relay configuration to use for subnets configured on this segment. This attribute is supported with NSX 3.0.0 onwards.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) Subnet configuration block.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) Gateway IP address CIDR.`,
				},
				resource.Attribute{
					Name:        "dhcp_ranges",
					Description: `(Optional) List of DHCP address ranges for dynamic IP allocation.`,
				},
				resource.Attribute{
					Name:        "dhcp_v4_config",
					Description: `(Optional) DHCPv4 config for IPv4 subnet. This attribute is supported with NSX 3.0.0 onwards.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Optional) IP address of the DHCP server in CIDR format. This attribute is required if segment has provided dhcp_config_path and it represents a DHCP server config.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `(Optional) List of IP addresses of DNS servers for the subnet.`,
				},
				resource.Attribute{
					Name:        "dhcp_option_121",
					Description: `(Optional) DHCP classless static routes.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Destination in cidr format.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required) IP address of next hop.`,
				},
				resource.Attribute{
					Name:        "dhcp_generic_option",
					Description: `(Optional) Generic DHCP options.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(Required) DHCP option code. Valid values are from 0 to 255.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) List of DHCP option values.`,
				},
				resource.Attribute{
					Name:        "dhcp_v6_config",
					Description: `(Optional) DHCPv6 config for IPv6 subnet. This attribute is supported with NSX 3.0.0 onwards.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Optional) IP address of the DHCP server in CIDR format. This attribute is required if segment has provided dhcp_config_path and it represents a DHCP server config.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `(Optional) List of IP addresses of DNS servers for the subnet.`,
				},
				resource.Attribute{
					Name:        "preferred_time",
					Description: `(Optional) The time interval in seconds, in which the prefix is advertised as preferred.`,
				},
				resource.Attribute{
					Name:        "domain_names",
					Description: `(Optional) List of domain names for this subnet.`,
				},
				resource.Attribute{
					Name:        "excluded_range",
					Description: `(Optional) List of excluded address ranges to define dynamic ip allocation ranges.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) IPv6 address that marks beginning of the range.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) IPv6 address that marks end of the range.`,
				},
				resource.Attribute{
					Name:        "sntp_servers",
					Description: `(Optional) IPv6 address of SNTP servers for the subnet.`,
				},
				resource.Attribute{
					Name:        "l2_extension",
					Description: `(Optional) Configuration for extending Segment through L2 VPN.`,
				},
				resource.Attribute{
					Name:        "l2vpn_paths",
					Description: `(Optional) Policy paths of associated L2 VPN sessions.`,
				},
				resource.Attribute{
					Name:        "tunnel_id",
					Description: `(Optional) The Tunnel ID that's a int value between 1 and 4093.`,
				},
				resource.Attribute{
					Name:        "advanced_config",
					Description: `(Optional) Advanced Segment configuration.`,
				},
				resource.Attribute{
					Name:        "address_pool_path",
					Description: `(Optional) Policy path to IP address pool.`,
				},
				resource.Attribute{
					Name:        "connectivity",
					Description: `(Optional) Connectivity configuration to manually connect (ON) or disconnect (OFF).`,
				},
				resource.Attribute{
					Name:        "hybrid",
					Description: `(Optional) Boolean flag to identify a hybrid logical switch.`,
				},
				resource.Attribute{
					Name:        "local_egress",
					Description: `(Optional) Boolean flag to enable local egress.`,
				},
				resource.Attribute{
					Name:        "uplink_teaming_policy",
					Description: `(Optional) The name of the switching uplink teaming policy for the bridge endpoint. This name corresponds to one of the switching uplink teaming policy names listed in the transport zone. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Secuirty Policy.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing segment can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_vlan_segment.segment1 ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the VLAN backed segment named ` + "`" + `segment1` + "`" + ` with the NSX Segment ID ` + "`" + `ID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Secuirty Policy.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The NSX path of the policy resource. ## Importing An existing segment can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_vlan_segment.segment1 ID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the VLAN backed segment named ` + "`" + `segment1` + "`" + ` with the NSX Segment ID ` + "`" + `ID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_policy_vm_tags",
			Category:         "Policy Resources",
			ShortDescription: `A resource to configure tags for a Virtual Machine in NSX Policy.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"vm",
				"tags",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the Virtual Machine. Can be the instance UUID or BIOS UUID.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this Virtual Machine. ## Importing An existing Tags collection can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_policy_vm_tags.vm1_tags ID ` + "`" + `` + "`" + `` + "`" + ` The above would import NSX Virtual Machine tags as a resource named ` + "`" + `vm1_tags` + "`" + ` with the NSX ID ` + "`" + `ID` + "`" + `, where ID is external ID of the Virtual Machine.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_qos_switching_profile",
			Category:         "Manager Switching Profiles Resources",
			ShortDescription: `Provides a resource to configure QoS switching profile on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"switching",
				"profiles",
				"qos",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this qos switching profile.`,
				},
				resource.Attribute{
					Name:        "class_of_service",
					Description: `(Optional) Class of service.`,
				},
				resource.Attribute{
					Name:        "dscp_trusted",
					Description: `(Optional) Trust mode for DSCP (False by default)`,
				},
				resource.Attribute{
					Name:        "dscp_priority",
					Description: `(Optional) DSCP Priority (0-63)`,
				},
				resource.Attribute{
					Name:        "ingress_rate_shaper",
					Description: `(Optional) Ingress rate shaper configuration:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether this rate shaper is enabled.`,
				},
				resource.Attribute{
					Name:        "average_bw_mbps",
					Description: `(Optional) Average Bandwidth in MBPS.`,
				},
				resource.Attribute{
					Name:        "peak_bw_mbps",
					Description: `(Optional) Peak Bandwidth in MBPS.`,
				},
				resource.Attribute{
					Name:        "burst_size",
					Description: `(Optional) Burst size in bytes.`,
				},
				resource.Attribute{
					Name:        "egress_rate_shaper",
					Description: `(Optional) Egress rate shaper configuration:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether this rate shaper is enabled.`,
				},
				resource.Attribute{
					Name:        "average_bw_mbps",
					Description: `(Optional) Average Bandwidth in MBPS.`,
				},
				resource.Attribute{
					Name:        "peak_bw_mbps",
					Description: `(Optional) Peak Bandwidth in MBPS.`,
				},
				resource.Attribute{
					Name:        "burst_size",
					Description: `(Optional) Burst size in bytes.`,
				},
				resource.Attribute{
					Name:        "ingress_broadcast_rate_shaper",
					Description: `(Optional) Ingress rate shaper configuration:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether this rate shaper is enabled.`,
				},
				resource.Attribute{
					Name:        "average_bw_kbps",
					Description: `(Optional) Average Bandwidth in KBPS.`,
				},
				resource.Attribute{
					Name:        "peak_bw_kbps",
					Description: `(Optional) Peak Bandwidth in KBPS.`,
				},
				resource.Attribute{
					Name:        "burst_size",
					Description: `(Optional) Burst size in bytes. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the QoS switching profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing qos switching profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_qos_switching_profile.qos_switching_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the Qos switching profile named ` + "`" + `qos_switching_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the QoS switching profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing qos switching profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_qos_switching_profile.qos_switching_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the Qos switching profile named ` + "`" + `qos_switching_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_spoofguard_switching_profile",
			Category:         "Manager Switching Profiles Resources",
			ShortDescription: `Provides a resource to configure spoofguard switching profile on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"switching",
				"profiles",
				"spoofguard",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this spoofguard switching profile.`,
				},
				resource.Attribute{
					Name:        "address_binding_whitelist_enabled",
					Description: `(Optional) A boolean flag indicating whether this profile overrides the default system wide settings for Spoof Guard when assigned to ports. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the spoofguard switching profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing spoofguard switching profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_spoofguard_switching_profile.spoofguard_switching_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the spoofguard switching profile named ` + "`" + `spoofguard_switching_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the spoofguard switching profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing spoofguard switching profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_spoofguard_switching_profile.spoofguard_switching_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import the spoofguard switching profile named ` + "`" + `spoofguard_switching_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_static_route",
			Category:         "Manager Resources",
			ShortDescription: `A resource to configure a static route in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"static",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this static route.`,
				},
				resource.Attribute{
					Name:        "logical_router_id",
					Description: `(Required) Logical router id.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) CIDR.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required) List of Next Hops, each with those arguments:`,
				},
				resource.Attribute{
					Name:        "administrative_distance",
					Description: `(Optional) Administrative Distance for the next hop IP.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Next Hop IP.`,
				},
				resource.Attribute{
					Name:        "logical_router_port_id",
					Description: `(Optional) Reference of logical router port to be used for next hop. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the static route.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "bfd_enabled",
					Description: `Status of bfd for this next hop where bfd_enabled = true indicate bfd is enabled for this next hop and bfd_enabled = false indicate bfd peer is disabled or not configured for this next hop.`,
				},
				resource.Attribute{
					Name:        "blackhole_action",
					Description: `Action to be taken on matching packets for NULL routes. ## Importing An existing static route can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_static_route.static_route logical-router-uuid/static-route-num ` + "`" + `` + "`" + `` + "`" + ` The above command imports the static route named ` + "`" + `static_route` + "`" + ` with the number ` + "`" + `static-route-num` + "`" + ` that belongs to the tier 1 logical router with the NSX id ` + "`" + `logical-router-uuid` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the static route.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging.`,
				},
				resource.Attribute{
					Name:        "bfd_enabled",
					Description: `Status of bfd for this next hop where bfd_enabled = true indicate bfd is enabled for this next hop and bfd_enabled = false indicate bfd peer is disabled or not configured for this next hop.`,
				},
				resource.Attribute{
					Name:        "blackhole_action",
					Description: `Action to be taken on matching packets for NULL routes. ## Importing An existing static route can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_static_route.static_route logical-router-uuid/static-route-num ` + "`" + `` + "`" + `` + "`" + ` The above command imports the static route named ` + "`" + `static_route` + "`" + ` with the number ` + "`" + `static-route-num` + "`" + ` that belongs to the tier 1 logical router with the NSX id ` + "`" + `logical-router-uuid` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_switch_security_switching_profile",
			Category:         "Manager Switching Profiles Resources",
			ShortDescription: `Provides a resource to configure switch security switching profile on NSX-T manager`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"switching",
				"profiles",
				"switch",
				"security",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this resource.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of this resource. Defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this qos switching profile.`,
				},
				resource.Attribute{
					Name:        "block_non_ip",
					Description: `(Optional) Indicates whether blocking of all traffic except IP/(G)ARP/BPDU is enabled.`,
				},
				resource.Attribute{
					Name:        "block_client_dhcp",
					Description: `(Optional) Indicates whether DHCP client blocking is enabled`,
				},
				resource.Attribute{
					Name:        "block_server_dhcp",
					Description: `(Optional) Indicates whether DHCP server blocking is enabled`,
				},
				resource.Attribute{
					Name:        "bpdu_filter_enabled",
					Description: `(Optional) Indicates whether BPDU filter is enabled`,
				},
				resource.Attribute{
					Name:        "bpdu_filter_whitelist",
					Description: `(Optional) Set of allowed MAC addresses to be excluded from BPDU filtering, if enabled.`,
				},
				resource.Attribute{
					Name:        "rate_limits",
					Description: `(Optional) Rate limit definitions for broadcast and multicast traffic.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether rate limitimg is enabled.`,
				},
				resource.Attribute{
					Name:        "rx_broadcast",
					Description: `(Optional) Incoming broadcast traffic limit in packets per second.`,
				},
				resource.Attribute{
					Name:        "rx_multicast",
					Description: `(Optional) Incoming multicast traffic limit in packets per second.`,
				},
				resource.Attribute{
					Name:        "tx_broadcast",
					Description: `(Optional) Outgoing broadcast traffic limit in packets per second.`,
				},
				resource.Attribute{
					Name:        "tx_multicast",
					Description: `(Optional) Outgoing multicast traffic limit in packets per second. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the switch security switching profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing switch security switching profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_switch_security_switching_profile.switch_security_switching_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import switching profile named ` + "`" + `switch_security_switching_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the switch security switching profile.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing switch security switching profile can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_switch_security_switching_profile.switch_security_switching_profile UUID ` + "`" + `` + "`" + `` + "`" + ` The above would import switching profile named ` + "`" + `switch_security_switching_profile` + "`" + ` with the nsx id ` + "`" + `UUID` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_vlan_logical_switch",
			Category:         "Manager Resources",
			ShortDescription: `A resource to configure vlan logical switch in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"vlan",
				"logical",
				"switch",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "transport_zone_id",
					Description: `(Required) Transport Zone ID for the logical switch.`,
				},
				resource.Attribute{
					Name:        "admin_state",
					Description: `(Optional) Admin state for the logical switch. Accepted values - 'UP' or 'DOWN'. The default value is 'UP'.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Required) Vlan for the logical switch.`,
				},
				resource.Attribute{
					Name:        "switching_profile_id",
					Description: `(Optional) List of IDs of switching profiles (of various types) to be associated with this switch. Default switching profiles will be used if not specified.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name, defaults to ID if not set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "ip_pool_id",
					Description: `(Optional) Ip Pool ID to be associated with the logical switch.`,
				},
				resource.Attribute{
					Name:        "mac_pool_id",
					Description: `(Optional) Mac Pool ID to be associated with the logical switch.`,
				},
				resource.Attribute{
					Name:        "address_binding",
					Description: `(Optional) List of Address Bindings for the logical switch. This setting allows to provide bindings between IP address, mac Address and vlan.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this logical switch. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical switch.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing X can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_vlan_logical_switch.switch1 UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical switch named ` + "`" + `switch1` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the logical switch.`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `Indicates current revision number of the object as seen by NSX-T API server. This attribute can be useful for debugging. ## Importing An existing X can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_vlan_logical_switch.switch1 UUID ` + "`" + `` + "`" + `` + "`" + ` The above command imports the logical switch named ` + "`" + `switch1` + "`" + ` with the NSX id ` + "`" + `UUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_vm_tags",
			Category:         "Manager Resources",
			ShortDescription: `A resource to configure tags for a virtual machine in NSX.`,
			Description:      ``,
			Keywords: []string{
				"manager",
				"vm",
				"tags",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) BIOS Id of the Virtual Machine.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A list of scope + tag pairs to associate with this VM.`,
				},
				resource.Attribute{
					Name:        "logical_port_tag",
					Description: `(Optional) A list of scope + tag pairs to associate with all logical ports that are automatically created for this VM. ## Importing An existing Tags collection can be [imported][docs-import] into this resource, via the following command: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import nsxt_vm_tags.vm1_tags id ` + "`" + `` + "`" + `` + "`" + ` The above would import NSX virtual machine tags as a resource named ` + "`" + `vm1_tags` + "`" + ` with the NSX id ` + "`" + `id` + "`" + `, where id is external ID (not the BIOS id) of the virtual machine.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"nsxt_algorithm_type_ns_service":               0,
		"nsxt_dhcp_relay_profile":                      1,
		"nsxt_dhcp_relay_service":                      2,
		"nsxt_dhcp_server_ip_pool":                     3,
		"nsxt_dhcp_server_profile":                     4,
		"nsxt_ether_type_ns_service":                   5,
		"nsxt_firewall_section":                        6,
		"nsxt_icmp_type_ns_service":                    7,
		"nsxt_igmp_type_ns_service":                    8,
		"nsxt_ip_block":                                9,
		"nsxt_ip_block_subnet":                         10,
		"nsxt_ip_discovery_switching_profile":          11,
		"nsxt_ip_pool":                                 12,
		"nsxt_ip_pool_allocation_ip_address":           13,
		"nsxt_ip_protocol_ns_service":                  14,
		"nsxt_ip_set":                                  15,
		"nsxt_l4_port_set_ns_service":                  16,
		"nsxt_lb_client_ssl_profile":                   17,
		"nsxt_lb_cookie_persistence_profile":           18,
		"nsxt_lb_fast_tcp_application_profile":         19,
		"nsxt_lb_fast_udp_application_profile":         20,
		"nsxt_lb_http_application_profile":             21,
		"nsxt_lb_http_forwarding_rule":                 22,
		"nsxt_lb_http_monitor":                         23,
		"nsxt_lb_http_request_rewrite_rule":            24,
		"nsxt_lb_http_response_rewrite_rule":           25,
		"nsxt_lb_http_virtual_server":                  26,
		"nsxt_lb_https_monitor":                        27,
		"nsxt_lb_icmp_monitor":                         28,
		"nsxt_lb_passive_monitor":                      29,
		"nsxt_lb_pool":                                 30,
		"nsxt_lb_server_ssl_profile":                   31,
		"nsxt_lb_service":                              32,
		"nsxt_lb_source__persistence_profile":          33,
		"nsxt_lb_tcp_monitor":                          34,
		"nsxt_lb_tcp_virtual_server":                   35,
		"nsxt_lb_udp_monitor":                          36,
		"nsxt_lb_udp_virtual_server":                   37,
		"nsxt_logical_dhcp_port":                       38,
		"nsxt_logical_dhcp_server":                     39,
		"nsxt_logical_port":                            40,
		"nsxt_logical_router_centralized_service_port": 41,
		"nsxt_logical_router_downlink_port":            42,
		"nsxt_logical_router_link_port_on_tier0":       43,
		"nsxt_logical_router_link_port_on_tier1":       44,
		"nsxt_logical_switch":                          45,
		"nsxt_logical_tier0_router":                    46,
		"nsxt_logical_tier1_router":                    47,
		"nsxt_mac_management_switching_profile":        48,
		"nsxt_nat_rule":                                49,
		"nsxt_ns_group":                                50,
		"nsxt_ns_service_group":                        51,
		"nsxt_policy_bgp_neighbor":                     52,
		"nsxt_policy_dhcp_relay":                       53,
		"nsxt_policy_dhcp_server":                      54,
		"nsxt_policy_gateway_policy":                   55,
		"nsxt_policy_group":                            56,
		"nsxt_policy_ip_address_allocation":            57,
		"nsxt_policy_ip_block":                         58,
		"nsxt_policy_ip_pool":                          59,
		"nsxt_policy_ip_pool_block_subnet":             60,
		"nsxt_policy_ip_pool_static_subnet":            61,
		"nsxt_policy_lb_pool":                          62,
		"nsxt_policy_lb_service":                       63,
		"nsxt_policy_lb_virtual_server":                64,
		"nsxt_policy_nat_rule":                         65,
		"nsxt_policy_security_policy":                  66,
		"nsxt_policy_segment":                          67,
		"nsxt_policy_service":                          68,
		"nsxt_policy_static_route":                     69,
		"nsxt_policy_tier0_gateway":                    70,
		"nsxt_policy_tier0_gateway_interface":          71,
		"nsxt_policy_tier1_gateway":                    72,
		"nsxt_policy_tier1_gateway_interface":          73,
		"nsxt_policy_vlan_segment":                     74,
		"nsxt_policy_vm_tags":                          75,
		"nsxt_qos_switching_profile":                   76,
		"nsxt_spoofguard_switching_profile":            77,
		"nsxt_static_route":                            78,
		"nsxt_switch_security_switching_profile":       79,
		"nsxt_vlan_logical_switch":                     80,
		"nsxt_vm_tags":                                 81,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
