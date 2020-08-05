package vthunder

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vthunder_configure_sync",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder configure sync resource for A10`,
			Description:      ``,
			Keywords: []string{
				"configure",
				"sync",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all_partitions",
					Description: `All partition configurations`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `Use private key for authentication`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_dns_primary",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder dns primary resource for A10`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"primary",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_v4_addr",
					Description: `IP v4 address`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ethernet",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ethernet resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ethernet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ifnum",
					Description: `Index of the ethernet interface`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `flag to indicate if dhcp is enabled or disabled`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `ipv4 address to be assigned to the ethernet interface`,
				},
				resource.Attribute{
					Name:        "ipv4_netmask",
					Description: `ipv4 subnet mask to be assigned to the ethernet interface`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `flag to indicate if interface is enabled or disabled`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_glm",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder GLM resource for A10`,
			Description:      ``,
			Keywords: []string{
				"glm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "token",
					Description: `License entitlement token`,
				},
				resource.Attribute{
					Name:        "use_mgmt_port",
					Description: `Use management port to connect to GLM`,
				},
				resource.Attribute{
					Name:        "enable_requests",
					Description: `Turn on periodic GLM license requests (default license retrieval interval is every 24 hours)`,
				},
				resource.Attribute{
					Name:        "allocate_bandwidth",
					Description: `Enter the requested bandwidth in Mbps for Capacity Pool`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_harmony_controller_profile",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder harmony-controller profile resource for A10`,
			Description:      ``,
			Keywords: []string{
				"harmony",
				"controller",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "host",
					Description: `Set harmony controller host adddress`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Set port for remote Harmony Controller, default is 8443`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `user-name for the tenant`,
				},
				resource.Attribute{
					Name:        "secret_value",
					Description: `Specify the password for the user`,
				},
				resource.Attribute{
					Name:        "provider2",
					Description: `provider for the harmony-controller`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `'register’: Register the device to the controller; 'deregister’: Deregister the device from controller;`,
				},
				resource.Attribute{
					Name:        "use_mgmt_port",
					Description: `Use management port for connections`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `region of the thunder-device`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `availablity zone of the thunder-device`,
				},
				resource.Attribute{
					Name:        "thunder_mgmt_ip",
					Description: ``,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address (IPv4 address)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_import",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder import resource for A10`,
			Description:      ``,
			Keywords: []string{
				"import",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "password",
					Description: `password for the remote site`,
				},
				resource.Attribute{
					Name:        "cloud_creds",
					Description: `Cloud Credentials File`,
				},
				resource.Attribute{
					Name:        "use_mgmt_port",
					Description: `Use management port as source port`,
				},
				resource.Attribute{
					Name:        "overwrite",
					Description: `Overwrite existing file`,
				},
				resource.Attribute{
					Name:        "remote_file",
					Description: `profile name for remote url`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_interface_ethernet",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder interface ethernet resource for A10`,
			Description:      ``,
			Keywords: []string{
				"interface",
				"ethernet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `‘enable’: Enable Router Advertisements on this interface; ‘disable’: Disable Router Advertisements on this interface;`,
				},
				resource.Attribute{
					Name:        "auto_neg_enable",
					Description: `enable auto-negotiation`,
				},
				resource.Attribute{
					Name:        "cpu_process",
					Description: `All Packets to this port are processed by CPU`,
				},
				resource.Attribute{
					Name:        "cpu_process_dir",
					Description: `‘primary’: Primary board; ‘blade’: blade board; ‘hash-dip’: Hash based on the Destination IP; ‘hash-sip’: Hash based on the Source IP; ‘hash-dmac’: Hash based on the Destination MAC; ‘hash-smac’: Hash based on the Source MAC;`,
				},
				resource.Attribute{
					Name:        "duplexity",
					Description: `‘Full’: Full; ‘Half’: Half; ‘auto’: auto;`,
				},
				resource.Attribute{
					Name:        "fec_forced_off",
					Description: `turn off the FEC`,
				},
				resource.Attribute{
					Name:        "fec_forced_on",
					Description: `turn on the FEC`,
				},
				resource.Attribute{
					Name:        "flow_control",
					Description: `Enable 802.3x flow control on full duplex port`,
				},
				resource.Attribute{
					Name:        "ifnum",
					Description: `Ethernet interface number`,
				},
				resource.Attribute{
					Name:        "load_interval",
					Description: `Configure Load Interval (Seconds (5-300, Multiple of 5), default 300)`,
				},
				resource.Attribute{
					Name:        "media_type_copper",
					Description: `Set the media type to copper`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `OSPF interface MTU (MTU size)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name for the interface`,
				},
				resource.Attribute{
					Name:        "remove_vlan_tag",
					Description: `Remove the vlan tag for egressing packets`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `‘10’: 10; ‘100’: 100; ‘1000’: 1000; ‘auto’: auto;`,
				},
				resource.Attribute{
					Name:        "speed_forced_40g",
					Description: `force the speed to be 40G on 100G link`,
				},
				resource.Attribute{
					Name:        "traffic_distribution_mode",
					Description: `‘sip’: sip; ‘dip’: dip; ‘primary’: primary; ‘blade’: blade; ‘l4-src-port’: l4-src-port; ‘l4-dst-port’: l4-dst-port;`,
				},
				resource.Attribute{
					Name:        "trap_source",
					Description: `The trap source`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "allow_promiscuous_vip",
					Description: `Allow traffic to be associated with promiscuous VIP`,
				},
				resource.Attribute{
					Name:        "cache_spoofing_port",
					Description: `This interface connects to spoofing cache`,
				},
				resource.Attribute{
					Name:        "client",
					Description: `Client facing interface for IPv4/v6 traffic`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Use DHCP to configure IP address`,
				},
				resource.Attribute{
					Name:        "generate_membership_query",
					Description: `Enable Membership Query`,
				},
				resource.Attribute{
					Name:        "inside",
					Description: `Configure LW-4over6 outside interface`,
				},
				resource.Attribute{
					Name:        "max_resp_time",
					Description: `Maximum Response Time (Max Response Time (Default is 100))`,
				},
				resource.Attribute{
					Name:        "outside",
					Description: `Configure LW-4over6 inside interface`,
				},
				resource.Attribute{
					Name:        "query_interval",
					Description: `1 - 255 (Default is 125)`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Server facing interface for IPv4/v6 traffic`,
				},
				resource.Attribute{
					Name:        "slb_partition_redirect",
					Description: `Redirect SLB traffic across partition`,
				},
				resource.Attribute{
					Name:        "ttl_ignore",
					Description: `Ignore TTL decrement for a received packet before sending out`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `ISO routing area tag`,
				},
				resource.Attribute{
					Name:        "helper_address",
					Description: `Helper address for DHCP packets (IP address)`,
				},
				resource.Attribute{
					Name:        "access_list",
					Description: `Access-list for traffic from the outside`,
				},
				resource.Attribute{
					Name:        "acl_id",
					Description: `ACL id`,
				},
				resource.Attribute{
					Name:        "class_list",
					Description: `Class List (Class List Name)`,
				},
				resource.Attribute{
					Name:        "receive_packet",
					Description: `Enable receiving packet through the specified interface`,
				},
				resource.Attribute{
					Name:        "send_packet",
					Description: `Enable sending packets through the specified interface`,
				},
				resource.Attribute{
					Name:        "receive",
					Description: `Advertisement reception`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `‘1’: RIP version 1; ‘2’: RIP version 2; ‘1-compatible’: RIPv1-compatible; ‘1-2’: RIP version 1 & 2;`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `‘poisoned’: Perform split horizon with poisoned reverse; ‘disable’: Disable split horizon; ‘enable’: Perform split horizon without poisoned reverse;`,
				},
				resource.Attribute{
					Name:        "key_chain",
					Description: `Authentication key-chain (Name of key-chain)`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `‘md5’: Keyed message digest;`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `The RIP authentication string`,
				},
				resource.Attribute{
					Name:        "send",
					Description: `Advertisement transmission`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `IP address`,
				},
				resource.Attribute{
					Name:        "ipv4_netmask",
					Description: `IP subnet mask`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Enable authentication`,
				},
				resource.Attribute{
					Name:        "authentication_key",
					Description: `Authentication password (key) (The OSPF password (key))`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `Interface cost`,
				},
				resource.Attribute{
					Name:        "database_filter",
					Description: `‘all’: Filter all LSA;`,
				},
				resource.Attribute{
					Name:        "dead_interval",
					Description: `Interval after which a neighbor is declared dead (Seconds)`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `Time between HELLO packets (Seconds)`,
				},
				resource.Attribute{
					Name:        "ip_addr",
					Description: `Address of interface`,
				},
				resource.Attribute{
					Name:        "mtu_ignore",
					Description: `Ignores the MTU in DBD packets`,
				},
				resource.Attribute{
					Name:        "out",
					Description: `Outgoing LSA`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Router priority`,
				},
				resource.Attribute{
					Name:        "retransmit_interval",
					Description: `Time between retransmitting lost link state advertisements (Seconds)`,
				},
				resource.Attribute{
					Name:        "transmit_delay",
					Description: `Link state transmit delay (Seconds)`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Mesh group number`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "md5_value",
					Description: `The OSPF password (1-16)`,
				},
				resource.Attribute{
					Name:        "message_digest_key",
					Description: `Message digest authentication password (key) (Key id)`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `Disable BFD`,
				},
				resource.Attribute{
					Name:        "broadcast",
					Description: `Specify OSPF broadcast multi-access network`,
				},
				resource.Attribute{
					Name:        "non_broadcast",
					Description: `Specify OSPF NBMA network`,
				},
				resource.Attribute{
					Name:        "p2mp_nbma",
					Description: `Specify non-broadcast point-to-multipoint network`,
				},
				resource.Attribute{
					Name:        "point_to_multipoint",
					Description: `Specify OSPF point-to-multipoint network`,
				},
				resource.Attribute{
					Name:        "point_to_point",
					Description: `Specify OSPF point-to-point network`,
				},
				resource.Attribute{
					Name:        "bfd",
					Description: `Bidirectional Forwarding Detection (BFD)`,
				},
				resource.Attribute{
					Name:        "acl_name",
					Description: `Access-list Name`,
				},
				resource.Attribute{
					Name:        "lockup",
					Description: `Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)`,
				},
				resource.Attribute{
					Name:        "lockup_period",
					Description: `Lockup period (second)`,
				},
				resource.Attribute{
					Name:        "normal",
					Description: `Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit`,
				},
				resource.Attribute{
					Name:        "link_aggregation",
					Description: `Interface link aggregation information`,
				},
				resource.Attribute{
					Name:        "tx_dot1_tlvs",
					Description: `Interface lldp tx IEEE 802.1 Organizationally specific TLVs configuration`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `Interface vlan information`,
				},
				resource.Attribute{
					Name:        "notif_enable",
					Description: `Interface lldp notification enable`,
				},
				resource.Attribute{
					Name:        "notification",
					Description: `Interface lldp notification configuration`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `Configure which TLVs excluded. All basic TLVs will be included by default`,
				},
				resource.Attribute{
					Name:        "management_address",
					Description: `Interface lldp management address`,
				},
				resource.Attribute{
					Name:        "port_description",
					Description: `Interface lldp port description`,
				},
				resource.Attribute{
					Name:        "system_capabilities",
					Description: `Interface lldp system capabilities`,
				},
				resource.Attribute{
					Name:        "system_description",
					Description: `Interface lldp system description`,
				},
				resource.Attribute{
					Name:        "system_name",
					Description: `Interface lldp system name`,
				},
				resource.Attribute{
					Name:        "tx_tlvs",
					Description: `Interface lldp tx TLVs configuration`,
				},
				resource.Attribute{
					Name:        "rt_enable",
					Description: `Interface lldp enable/disable`,
				},
				resource.Attribute{
					Name:        "rx",
					Description: `Enable lldp rx`,
				},
				resource.Attribute{
					Name:        "tx",
					Description: `Enable lldp tx`,
				},
				resource.Attribute{
					Name:        "demand",
					Description: `Demand mode`,
				},
				resource.Attribute{
					Name:        "echo",
					Description: `Enable BFD Echo`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `Key ID`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `‘md5’: Keyed MD5; ‘meticulous-md5’: Meticulous Keyed MD5; ‘meticulous-sha1’: Meticulous Keyed SHA1; ‘sha1’: Keyed SHA1; ‘simple’: Simple Password;`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Configure the authentication password for interface`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Transmit interval between BFD packets (Milliseconds)`,
				},
				resource.Attribute{
					Name:        "min_rx",
					Description: `Minimum receive interval capability (Milliseconds)`,
				},
				resource.Attribute{
					Name:        "multiplier",
					Description: `Multiplier value used to compute holddown (value used to multiply the interval)`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘packets_input’: Input packets; ‘bytes_input’: Input bytes; ‘received_broadcasts’: Received broadcasts; ‘received_multicasts’: Received multicasts; ‘received_unicasts’: Received unicasts; ‘input_errors’: Input errors; ‘crc’: CRC; ‘frame’: Frames; ‘runts’: Runts; ‘giants’: Giants; ‘packets_output’: Output packets; ‘bytes_output’: Output bytes; ‘transmitted_broadcasts’: Transmitted braodcasts; ‘transmitted_multicasts’: Transmitted multicasts; ‘transmitted_unicasts’: Transmitted unicasts; ‘output_errors’: Output errors; ‘collisions’: Collisions;`,
				},
				resource.Attribute{
					Name:        "map_t_inside",
					Description: `Configure MAP inside interface (connected to MAP domains)`,
				},
				resource.Attribute{
					Name:        "map_t_outside",
					Description: `Configure MAP outside interface`,
				},
				resource.Attribute{
					Name:        "admin_key",
					Description: `LACP admin key (Admin key value)`,
				},
				resource.Attribute{
					Name:        "port_priority",
					Description: `Set LACP priority for a port (LACP port priority)`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `‘long’: Set LACP long timeout (default); ‘short’: Set LACP short timeout;`,
				},
				resource.Attribute{
					Name:        "trunk_number",
					Description: `Trunk Number`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `‘static’: Static (default); ‘lacp’: lacp; ‘lacp-udld’: lacp-udld;`,
				},
				resource.Attribute{
					Name:        "fast",
					Description: `fast timeout in unit of milli-seconds(default 1000)`,
				},
				resource.Attribute{
					Name:        "slow",
					Description: `slow timeout in unit of seconds`,
				},
				resource.Attribute{
					Name:        "bind_type",
					Description: `‘inside’: This interface is connected to NPTv6 inside networks; ‘outside’: This interface is connected to NPTv6 outside networks;`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `NPTv6 domain name`,
				},
				resource.Attribute{
					Name:        "circuit_type",
					Description: `‘level-1’: Level-1 only adjacencies are formed; ‘level-1-2’: Level-1-2 adjacencies are formed; ‘level-2-only’: Level-2 only adjacencies are formed;`,
				},
				resource.Attribute{
					Name:        "lsp_interval",
					Description: `Set LSP transmission interval (LSP transmission interval (milliseconds))`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `‘broadcast’: Specify IS-IS broadcast multi-access network; ‘point-to-point’: Specify IS-IS point-to-point network;`,
				},
				resource.Attribute{
					Name:        "padding",
					Description: `Add padding to IS-IS hello packets`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `‘level-1’: Speficy interval for level-1 CSNPs; ‘level-2’: Specify interval for level-2 CSNPs;`,
				},
				resource.Attribute{
					Name:        "hello_interval_minimal",
					Description: `Set Hello holdtime 1 second, interval depends on multiplier`,
				},
				resource.Attribute{
					Name:        "blocked",
					Description: `Block LSPs on this interface`,
				},
				resource.Attribute{
					Name:        "send_only",
					Description: `Authentication send-only`,
				},
				resource.Attribute{
					Name:        "wide_metric",
					Description: `Configure the wide metric for interface`,
				},
				resource.Attribute{
					Name:        "hello_multiplier",
					Description: `Set multiplier for Hello holding time (Hello multiplier value)`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `Configure the metric for interface (Default metric)`,
				},
				resource.Attribute{
					Name:        "csnp_interval",
					Description: `Set CSNP interval in seconds (CSNP interval value)`,
				},
				resource.Attribute{
					Name:        "lockup_period_v6",
					Description: `Lockup period (second)`,
				},
				resource.Attribute{
					Name:        "lockup_v6",
					Description: `Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)`,
				},
				resource.Attribute{
					Name:        "normal_v6",
					Description: `Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit`,
				},
				resource.Attribute{
					Name:        "ipv6_enable",
					Description: `Enable IPv6 processing`,
				},
				resource.Attribute{
					Name:        "address_type",
					Description: `‘anycast’: Configure an IPv6 anycast address; ‘link-local’: Configure an IPv6 link local address;`,
				},
				resource.Attribute{
					Name:        "ipv6_addr",
					Description: `Set the IPv6 address of an interface`,
				},
				resource.Attribute{
					Name:        "rip",
					Description: `RIP Routing for IPv6`,
				},
				resource.Attribute{
					Name:        "area_id_addr",
					Description: `OSPF area ID in IP address format`,
				},
				resource.Attribute{
					Name:        "area_id_num",
					Description: `OSPF area ID as a decimal value`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `ACL applied on incoming packets to this interface`,
				},
				resource.Attribute{
					Name:        "v6_acl_name",
					Description: `Apply ACL rules to incoming packets on this interface (Named Access List)`,
				},
				resource.Attribute{
					Name:        "broadcast_type",
					Description: `‘broadcast’: Specify OSPF broadcast multi-access network; ‘non-broadcast’: Specify OSPF NBMA network; ‘point-to-point’: Specify OSPF point-to-point network; ‘point-to-multipoint’: Specify OSPF point-to-multipoint network;`,
				},
				resource.Attribute{
					Name:        "network_instance_id",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "neig_inst",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "neighbor",
					Description: `OSPFv3 neighbor (Neighbor IPv6 address)`,
				},
				resource.Attribute{
					Name:        "neighbor_cost",
					Description: `OSPF cost for point-to-multipoint neighbor (metric)`,
				},
				resource.Attribute{
					Name:        "neighbor_poll_interval",
					Description: `OSPF dead-router polling interval (Seconds)`,
				},
				resource.Attribute{
					Name:        "neighbor_priority",
					Description: `OSPF priority of non-broadcast neighbor`,
				},
				resource.Attribute{
					Name:        "adver_mtu",
					Description: `Set Router Advertisement MTU Option`,
				},
				resource.Attribute{
					Name:        "adver_mtu_disable",
					Description: `Disable Router Advertisement MTU Option`,
				},
				resource.Attribute{
					Name:        "adver_vrid",
					Description: `Specify ha VRRP-A vrid`,
				},
				resource.Attribute{
					Name:        "adver_vrid_default",
					Description: `Default VRRP-A vrid`,
				},
				resource.Attribute{
					Name:        "default_lifetime",
					Description: `Set Router Advertisement Default Lifetime (default: 1800) (Default Lifetime (seconds))`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "floating_ip_default_vrid",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "hop_limit",
					Description: `Set Router Advertisement Hop Limit (default: 255)`,
				},
				resource.Attribute{
					Name:        "managed_config_action",
					Description: `‘enable’: Enable the Managed Address Configuration flag; ‘disable’: Disable the Managed Address Configuration flag (default);`,
				},
				resource.Attribute{
					Name:        "max_interval",
					Description: `Set Router Advertisement Max Interval (default: 600) (Max Router Advertisement Interval (seconds))`,
				},
				resource.Attribute{
					Name:        "min_interval",
					Description: `Set Router Advertisement Min Interval (default: 200) (Min Router Advertisement Interval (seconds))`,
				},
				resource.Attribute{
					Name:        "other_config_action",
					Description: `‘enable’: Enable the Other Stateful Configuration flag; ‘disable’: Disable the Other Stateful Configuration flag (default);`,
				},
				resource.Attribute{
					Name:        "rate_limit",
					Description: `Rate Limit the processing of incoming Router Solicitations (Max Number of Router Solicitations to process per second)`,
				},
				resource.Attribute{
					Name:        "reachable_time",
					Description: `Set Router Advertisement Reachable ime (default: 0) (Reachable Time (milliseconds))`,
				},
				resource.Attribute{
					Name:        "retransmit_timer",
					Description: `Set Router Advertisement Retransmit Timer (default: 0)`,
				},
				resource.Attribute{
					Name:        "use_floating_ip",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "use_floating_ip_default_vrid",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "not_autonomous",
					Description: `Specify that the Prefix is not usable for autoconfiguration (default:autonomous)`,
				},
				resource.Attribute{
					Name:        "not_on_link",
					Description: `Specify that the Prefix is not On-Link (default: on-link)`,
				},
				resource.Attribute{
					Name:        "preferred_lifetime",
					Description: `Specify Prefix Preferred Lifetime (default:604800) (Prefix Advertised Preferred Lifetime (default: 604800))`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `Set Router Advertisement On-Link Prefix (IPv6 On-Link Prefix)`,
				},
				resource.Attribute{
					Name:        "valid_lifetime",
					Description: `Specify Valid Lifetime (default:2592000) (Prefix Advertised Valid Lifetime (default: 2592000))`,
				},
				resource.Attribute{
					Name:        "mirror_index",
					Description: `Mirror index`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `‘input’: Incoming packets; ‘output’: Outgoing packets; ‘both’: Both incoming and outgoing packets;`,
				},
				resource.Attribute{
					Name:        "monitor_vlan",
					Description: `VLAN number`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_interface_ethernet_bfd",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder interface ethernet bfd resource for A10`,
			Description:      ``,
			Keywords: []string{
				"interface",
				"ethernet",
				"bfd",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ifnum",
					Description: `Interface no.`,
				},
				resource.Attribute{
					Name:        "demand",
					Description: `Demand mode`,
				},
				resource.Attribute{
					Name:        "echo",
					Description: `Enable BFD Echo`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `Key ID`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `‘md5’: Keyed MD5; ‘meticulous-md5’: Meticulous Keyed MD5; ‘meticulous-sha1’: Meticulous Keyed SHA1; ‘sha1’: Keyed SHA1; ‘simple’: Simple Password;`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Key String`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Transmit interval between BFD packets (Milliseconds)`,
				},
				resource.Attribute{
					Name:        "min_rx",
					Description: `Minimum receive interval capability (Milliseconds)`,
				},
				resource.Attribute{
					Name:        "multiplier",
					Description: `Multiplier value used to compute holddown (value used to multiply the interval)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_interface_ethernet_ipv6",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder interface ethernet ipv6 resource for A10`,
			Description:      ``,
			Keywords: []string{
				"interface",
				"ethernet",
				"ipv6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "inside",
					Description: `Inside (private) interface for stateful firewall`,
				},
				resource.Attribute{
					Name:        "ipv6_enable",
					Description: `Enable IPv6 processing`,
				},
				resource.Attribute{
					Name:        "outside",
					Description: `Outside (public) interface for stateful firewall`,
				},
				resource.Attribute{
					Name:        "ttl_ignore",
					Description: `Ignore TTL decrement for a received packet before sending out`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "address_type",
					Description: `‘anycast’: Configure an IPv6 anycast address; ‘link-local’: Configure an IPv6 link local address;`,
				},
				resource.Attribute{
					Name:        "ipv6_addr",
					Description: `Set the IPv6 address of an interface`,
				},
				resource.Attribute{
					Name:        "rip",
					Description: `RIP Routing for IPv6`,
				},
				resource.Attribute{
					Name:        "area_id_addr",
					Description: `OSPF area ID in IP address format`,
				},
				resource.Attribute{
					Name:        "area_id_num",
					Description: `OSPF area ID as a decimal value`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `ISO routing area tag`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `‘poisoned’: Perform split horizon with poisoned reverse; ‘disable’: Disable split horizon; ‘enable’: Perform split horizon without poisoned reverse;`,
				},
				resource.Attribute{
					Name:        "access_list",
					Description: `Access-list for traffic from the outside`,
				},
				resource.Attribute{
					Name:        "acl_name",
					Description: `Access-list Name`,
				},
				resource.Attribute{
					Name:        "class_list",
					Description: `Class List (Class List Name)`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `ACL applied on incoming packets to this interface`,
				},
				resource.Attribute{
					Name:        "v6_acl_name",
					Description: `Apply ACL rules to incoming packets on this interface (Named Access List)`,
				},
				resource.Attribute{
					Name:        "bfd",
					Description: `Bidirectional Forwarding Detection (BFD)`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `Disable BFD`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `Interface cost`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `Time between HELLO packets (Seconds)`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Router priority`,
				},
				resource.Attribute{
					Name:        "mtu_ignore",
					Description: `Ignores the MTU in DBD packets`,
				},
				resource.Attribute{
					Name:        "retransmit_interval",
					Description: `Time between retransmitting lost link state advertisements (Seconds)`,
				},
				resource.Attribute{
					Name:        "broadcast_type",
					Description: `‘broadcast’: Specify OSPF broadcast multi-access network; ‘non-broadcast’: Specify OSPF NBMA network; ‘point-to-point’: Specify OSPF point-to-point network; ‘point-to-multipoint’: Specify OSPF point-to-multipoint network;`,
				},
				resource.Attribute{
					Name:        "network_instance_id",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "p2mp_nbma",
					Description: `Specify non-broadcast point-to-multipoint network`,
				},
				resource.Attribute{
					Name:        "transmit_delay",
					Description: `Link state transmit delay (Seconds)`,
				},
				resource.Attribute{
					Name:        "neig_inst",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "neighbor",
					Description: `OSPFv3 neighbor (Neighbor IPv6 address)`,
				},
				resource.Attribute{
					Name:        "neighbor_cost",
					Description: `OSPF cost for point-to-multipoint neighbor (metric)`,
				},
				resource.Attribute{
					Name:        "neighbor_poll_interval",
					Description: `OSPF dead-router polling interval (Seconds)`,
				},
				resource.Attribute{
					Name:        "neighbor_priority",
					Description: `OSPF priority of non-broadcast neighbor`,
				},
				resource.Attribute{
					Name:        "dead_interval",
					Description: `Interval after which a neighbor is declared dead (Seconds)`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `‘enable’: Enable Router Advertisements on this interface; ‘disable’: Disable Router Advertisements on this interface;`,
				},
				resource.Attribute{
					Name:        "adver_mtu",
					Description: `Set Router Advertisement MTU Option`,
				},
				resource.Attribute{
					Name:        "adver_mtu_disable",
					Description: `Disable Router Advertisement MTU Option`,
				},
				resource.Attribute{
					Name:        "adver_vrid",
					Description: `Specify ha VRRP-A vrid`,
				},
				resource.Attribute{
					Name:        "adver_vrid_default",
					Description: `Default VRRP-A vrid`,
				},
				resource.Attribute{
					Name:        "default_lifetime",
					Description: `Set Router Advertisement Default Lifetime (default: 1800) (Default Lifetime (seconds))`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "floating_ip_default_vrid",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "hop_limit",
					Description: `Set Router Advertisement Hop Limit (default: 255)`,
				},
				resource.Attribute{
					Name:        "managed_config_action",
					Description: `‘enable’: Enable the Managed Address Configuration flag; ‘disable’: Disable the Managed Address Configuration flag (default);`,
				},
				resource.Attribute{
					Name:        "max_interval",
					Description: `Set Router Advertisement Max Interval (default: 600) (Max Router Advertisement Interval (seconds))`,
				},
				resource.Attribute{
					Name:        "min_interval",
					Description: `Set Router Advertisement Min Interval (default: 200) (Min Router Advertisement Interval (seconds))`,
				},
				resource.Attribute{
					Name:        "other_config_action",
					Description: `‘enable’: Enable the Other Stateful Configuration flag; ‘disable’: Disable the Other Stateful Configuration flag (default);`,
				},
				resource.Attribute{
					Name:        "rate_limit",
					Description: `Rate Limit the processing of incoming Router Solicitations (Max Number of Router Solicitations to process per second)`,
				},
				resource.Attribute{
					Name:        "reachable_time",
					Description: `Set Router Advertisement Reachable ime (default: 0) (Reachable Time (milliseconds))`,
				},
				resource.Attribute{
					Name:        "retransmit_timer",
					Description: `Set Router Advertisement Retransmit Timer (default: 0)`,
				},
				resource.Attribute{
					Name:        "use_floating_ip",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "use_floating_ip_default_vrid",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "not_autonomous",
					Description: `Specify that the Prefix is not usable for autoconfiguration (default:autonomous)`,
				},
				resource.Attribute{
					Name:        "not_on_link",
					Description: `Specify that the Prefix is not On-Link (default: on-link)`,
				},
				resource.Attribute{
					Name:        "preferred_lifetime",
					Description: `Specify Prefix Preferred Lifetime (default:604800) (Prefix Advertised Preferred Lifetime (default: 604800))`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `Set Router Advertisement On-Link Prefix (IPv6 On-Link Prefix)`,
				},
				resource.Attribute{
					Name:        "valid_lifetime",
					Description: `Specify Valid Lifetime (default:2592000) (Prefix Advertised Valid Lifetime (default: 2592000))`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_interface_ethernet_lldp",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder interface ethernet lldp resource for A10`,
			Description:      ``,
			Keywords: []string{
				"interface",
				"ethernet",
				"lldp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ifnum",
					Description: `Interface no.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "link_aggregation",
					Description: `Interface link aggregation information`,
				},
				resource.Attribute{
					Name:        "tx_dot1_tlvs",
					Description: `Interface lldp tx IEEE 802.1 Organizationally specific TLVs configuration`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `Interface vlan information`,
				},
				resource.Attribute{
					Name:        "notif_enable",
					Description: `Interface lldp notification enable`,
				},
				resource.Attribute{
					Name:        "notification",
					Description: `Interface lldp notification configuration`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `Configure which TLVs excluded. All basic TLVs will be included by default`,
				},
				resource.Attribute{
					Name:        "management_address",
					Description: `Interface lldp management address`,
				},
				resource.Attribute{
					Name:        "port_description",
					Description: `Interface lldp port description`,
				},
				resource.Attribute{
					Name:        "system_capabilities",
					Description: `Interface lldp system capabilities`,
				},
				resource.Attribute{
					Name:        "system_description",
					Description: `Interface lldp system description`,
				},
				resource.Attribute{
					Name:        "system_name",
					Description: `Interface lldp system name`,
				},
				resource.Attribute{
					Name:        "tx_tlvs",
					Description: `Interface lldp tx TLVs configuration`,
				},
				resource.Attribute{
					Name:        "rt_enable",
					Description: `Interface lldp enable/disable`,
				},
				resource.Attribute{
					Name:        "rx",
					Description: `Enable lldp rx`,
				},
				resource.Attribute{
					Name:        "tx",
					Description: `Enable lldp tx`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_interface_ethernet_trunk_group",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder interface ethernet trunk group resource for A10`,
			Description:      ``,
			Keywords: []string{
				"interface",
				"ethernet",
				"trunk",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_key",
					Description: `LACP admin key (Admin key value)`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `‘active’: enable initiation of LACP negotiation on a port(default); ‘passive’: disable initiation of LACP negotiation on a port;`,
				},
				resource.Attribute{
					Name:        "port_priority",
					Description: `Set LACP priority for a port (LACP port priority)`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `‘long’: Set LACP long timeout (default); ‘short’: Set LACP short timeout;`,
				},
				resource.Attribute{
					Name:        "trunk_number",
					Description: `Trunk Number`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `‘static’: Static (default); ‘lacp’: lacp; ‘lacp-udld’: lacp-udld;`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "fast",
					Description: `fast timeout in unit of milli-seconds(default 1000)`,
				},
				resource.Attribute{
					Name:        "slow",
					Description: `slow timeout in unit of seconds`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_interface_management",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder interface management resource for A10`,
			Description:      ``,
			Keywords: []string{
				"interface",
				"management",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `‘enable’: Enable Management Port; ‘disable’: Disable Management Port;`,
				},
				resource.Attribute{
					Name:        "duplexity",
					Description: `‘Full’: Full; ‘Half’: Half; ‘auto’: Auto;`,
				},
				resource.Attribute{
					Name:        "flow_control",
					Description: `Enable 802.3x flow control on full duplex port`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `‘10’: 10 Mbs/sec; ‘100’: 100 Mbs/sec; ‘1000’: 1 Gb/sec; ‘auto’: Auto Negotiate Speed; (Interface Speed)`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "link_aggregation",
					Description: `Interface link aggregation information`,
				},
				resource.Attribute{
					Name:        "tx_dot1_tlvs",
					Description: `Interface lldp tx IEEE 802.1 Organizationally specific TLVs configuration`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `Interface vlan information`,
				},
				resource.Attribute{
					Name:        "notif_enable",
					Description: `Interface lldp notification enable`,
				},
				resource.Attribute{
					Name:        "notification",
					Description: `Interface lldp notification configuration`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `Configure which TLVs excluded. All basic TLVs will be included by default`,
				},
				resource.Attribute{
					Name:        "management_address",
					Description: `Interface lldp management address`,
				},
				resource.Attribute{
					Name:        "port_description",
					Description: `Interface lldp port description`,
				},
				resource.Attribute{
					Name:        "system_capabilities",
					Description: `Interface lldp system capabilities`,
				},
				resource.Attribute{
					Name:        "system_description",
					Description: `Interface lldp system description`,
				},
				resource.Attribute{
					Name:        "system_name",
					Description: `Interface lldp system name`,
				},
				resource.Attribute{
					Name:        "tx_tlvs",
					Description: `Interface lldp tx TLVs configuration`,
				},
				resource.Attribute{
					Name:        "rt_enable",
					Description: `Interface lldp enable/disable`,
				},
				resource.Attribute{
					Name:        "rx",
					Description: `Enable lldp rx`,
				},
				resource.Attribute{
					Name:        "tx",
					Description: `Enable lldp tx`,
				},
				resource.Attribute{
					Name:        "bcast_rate_limit_enable",
					Description: `Rate limit the l2 broadcast packet on mgmt port`,
				},
				resource.Attribute{
					Name:        "rate",
					Description: `packets per second. Default is 500. (packets per second. Please specify an even number. Default is 500)`,
				},
				resource.Attribute{
					Name:        "address_type",
					Description: `‘link-local’: Configure an IPv6 link local address;`,
				},
				resource.Attribute{
					Name:        "default_ipv6_gateway",
					Description: `Set default gateway (Default gateway address)`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `ACL applied on incoming packets to this interface`,
				},
				resource.Attribute{
					Name:        "ipv6_addr",
					Description: `Set the IPv6 address of an interface`,
				},
				resource.Attribute{
					Name:        "v6_acl_name",
					Description: `Apply ACL rules to incoming packets on this interface (Named Access List)`,
				},
				resource.Attribute{
					Name:        "control_apps_use_mgmt_port",
					Description: `Control applications use management port`,
				},
				resource.Attribute{
					Name:        "default_gateway",
					Description: `Set default gateway (Default gateway address)`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Use DHCP to configure IP address`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `IP address`,
				},
				resource.Attribute{
					Name:        "ipv4_netmask",
					Description: `IP subnet mask`,
				},
				resource.Attribute{
					Name:        "secondary_ip",
					Description: `Global IP configuration subcommands`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘packets_input’: Input packets; ‘bytes_input’: Input bytes; ‘received_broadcasts’: Received broadcasts; ‘received_multicasts’: Received multicasts; ‘received_unicasts’: Received unicasts; ‘input_errors’: Input errors; ‘crc’: CRC; ‘frame’: Frames; ‘input_err_short’: Runts; ‘input_err_long’: Giants; ‘packets_output’: Output packets; ‘bytes_output’: Output bytes; ‘transmitted_broadcasts’: Transmitted broadcasts; ‘transmitted_multicasts’: Transmitted multicasts; ‘transmitted_unicasts’: Transmitted unicasts; ‘output_errors’: Output errors; ‘collisions’: Collisions;`,
				},
				resource.Attribute{
					Name:        "acl_id",
					Description: `ACL id`,
				},
				resource.Attribute{
					Name:        "acl_name",
					Description: `Apply an access list (Named Access List)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_interface_ve",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder interface ve resource for A10`,
			Description:      ``,
			Keywords: []string{
				"interface",
				"ve",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `‘enable’: Enable Router Advertisements on this interface; ‘disable’: Disable Router Advertisements on this interface;`,
				},
				resource.Attribute{
					Name:        "ifnum",
					Description: `Virtual ethernet interface number`,
				},
				resource.Attribute{
					Name:        "l3_vlan_fwd_disable",
					Description: `Disable L3 forwarding between VLANs for incoming packets on this interface`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `OSPF interface MTU (MTU size)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name for the interface`,
				},
				resource.Attribute{
					Name:        "trap_source",
					Description: `The trap source`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "inside",
					Description: `DDoS inside (trusted) or outside (untrusted) interface`,
				},
				resource.Attribute{
					Name:        "map_t_inside",
					Description: `Configure MAP inside interface (connected to MAP domains)`,
				},
				resource.Attribute{
					Name:        "map_t_outside",
					Description: `Configure MAP outside interface`,
				},
				resource.Attribute{
					Name:        "outside",
					Description: `DDoS inside (trusted) or outside (untrusted) interface`,
				},
				resource.Attribute{
					Name:        "lockup",
					Description: `Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)`,
				},
				resource.Attribute{
					Name:        "lockup_period",
					Description: `Lockup period (second)`,
				},
				resource.Attribute{
					Name:        "normal",
					Description: `Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit`,
				},
				resource.Attribute{
					Name:        "circuit_type",
					Description: `‘level-1’: Level-1 only adjacencies are formed; ‘level-1-2’: Level-1-2 adjacencies are formed; ‘level-2-only’: Level-2 only adjacencies are formed;`,
				},
				resource.Attribute{
					Name:        "lsp_interval",
					Description: `Set LSP transmission interval (LSP transmission interval (milliseconds))`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `‘broadcast’: Specify IS-IS broadcast multi-access network; ‘point-to-point’: Specify IS-IS point-to-point network;`,
				},
				resource.Attribute{
					Name:        "padding",
					Description: `Add padding to IS-IS hello packets`,
				},
				resource.Attribute{
					Name:        "retransmit_interval",
					Description: `Time between retransmitting lost link state advertisements (Seconds)`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `‘level-1’: Speficy interval for level-1 CSNPs; ‘level-2’: Specify interval for level-2 CSNPs;`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Router priority`,
				},
				resource.Attribute{
					Name:        "hello_interval_minimal",
					Description: `Set Hello holdtime 1 second, interval depends on multiplier`,
				},
				resource.Attribute{
					Name:        "blocked",
					Description: `Block LSPs on this interface`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `‘message-digest’: Use message-digest authentication; ‘null’: Use no authentication;`,
				},
				resource.Attribute{
					Name:        "bfd",
					Description: `Bidirectional Forwarding Detection (BFD)`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `Disable BFD`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Key String`,
				},
				resource.Attribute{
					Name:        "key_chain",
					Description: `Authentication key-chain (Name of key-chain)`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `‘md5’: Keyed message digest; ‘text’: Clear text authentication;`,
				},
				resource.Attribute{
					Name:        "send_only",
					Description: `Authentication send-only`,
				},
				resource.Attribute{
					Name:        "wide_metric",
					Description: `Configure the wide metric for interface`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `Time between HELLO packets (Seconds)`,
				},
				resource.Attribute{
					Name:        "hello_multiplier",
					Description: `Set multiplier for Hello holding time (Hello multiplier value)`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `Configure the metric for interface (Default metric)`,
				},
				resource.Attribute{
					Name:        "csnp_interval",
					Description: `Set CSNP interval in seconds (CSNP interval value)`,
				},
				resource.Attribute{
					Name:        "demand",
					Description: `Demand mode`,
				},
				resource.Attribute{
					Name:        "echo",
					Description: `Enable BFD Echo`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `Key ID`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `‘md5’: Keyed MD5; ‘meticulous-md5’: Meticulous Keyed MD5; ‘meticulous-sha1’: Meticulous Keyed SHA1; ‘sha1’: Keyed SHA1; ‘simple’: Simple Password;`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Transmit interval between BFD packets (Milliseconds)`,
				},
				resource.Attribute{
					Name:        "min_rx",
					Description: `Minimum receive interval capability (Milliseconds)`,
				},
				resource.Attribute{
					Name:        "multiplier",
					Description: `Multiplier value used to compute holddown (value used to multiply the interval)`,
				},
				resource.Attribute{
					Name:        "allow_promiscuous_vip",
					Description: `Allow traffic to be associated with promiscuous VIP`,
				},
				resource.Attribute{
					Name:        "client",
					Description: `Client facing interface for IPv4/v6 traffic`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Use DHCP to configure IP address`,
				},
				resource.Attribute{
					Name:        "generate_membership_query",
					Description: `Enable Membership Query`,
				},
				resource.Attribute{
					Name:        "max_resp_time",
					Description: `Maximum Response Time (Max Response Time (Default is 100))`,
				},
				resource.Attribute{
					Name:        "query_interval",
					Description: `1 - 255 (Default is 125)`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Server facing interface for IPv4/v6 traffic`,
				},
				resource.Attribute{
					Name:        "slb_partition_redirect",
					Description: `Redirect SLB traffic across partition`,
				},
				resource.Attribute{
					Name:        "ttl_ignore",
					Description: `Ignore TTL decrement for a received packet`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `IP address`,
				},
				resource.Attribute{
					Name:        "ipv4_netmask",
					Description: `IP subnet mask`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `ISO routing area tag`,
				},
				resource.Attribute{
					Name:        "helper_address",
					Description: `Helper address for DHCP packets (IP address)`,
				},
				resource.Attribute{
					Name:        "access_list",
					Description: `Access-list for traffic from the outside`,
				},
				resource.Attribute{
					Name:        "acl_id",
					Description: `ACL id`,
				},
				resource.Attribute{
					Name:        "class_list",
					Description: `Class List (Class List Name)`,
				},
				resource.Attribute{
					Name:        "receive_packet",
					Description: `Enable receiving packet through the specified interface`,
				},
				resource.Attribute{
					Name:        "send_packet",
					Description: `Enable sending packets through the specified interface`,
				},
				resource.Attribute{
					Name:        "receive",
					Description: `Advertisement reception`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `‘1’: RIP version 1; ‘2’: RIP version 2; ‘1-compatible’: RIPv1-compatible; ‘1-2’: RIP version 1 & 2;`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `‘poisoned’: Perform split horizon with poisoned reverse; ‘disable’: Disable split horizon; ‘enable’: Perform split horizon without poisoned reverse;`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `The RIP authentication string`,
				},
				resource.Attribute{
					Name:        "send",
					Description: `Advertisement transmission`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Enable authentication`,
				},
				resource.Attribute{
					Name:        "authentication_key",
					Description: `Authentication password (key) (The OSPF password (key))`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `Interface cost`,
				},
				resource.Attribute{
					Name:        "database_filter",
					Description: `‘all’: Filter all LSA;`,
				},
				resource.Attribute{
					Name:        "dead_interval",
					Description: `Interval after which a neighbor is declared dead (Seconds)`,
				},
				resource.Attribute{
					Name:        "ip_addr",
					Description: `Address of interface`,
				},
				resource.Attribute{
					Name:        "mtu_ignore",
					Description: `Ignores the MTU in DBD packets`,
				},
				resource.Attribute{
					Name:        "out",
					Description: `Outgoing LSA`,
				},
				resource.Attribute{
					Name:        "transmit_delay",
					Description: `Link state transmit delay (Seconds)`,
				},
				resource.Attribute{
					Name:        "md5_value",
					Description: `The OSPF password (1-16)`,
				},
				resource.Attribute{
					Name:        "message_digest_key",
					Description: `Message digest authentication password (key) (Key id)`,
				},
				resource.Attribute{
					Name:        "broadcast",
					Description: `Specify OSPF broadcast multi-access network`,
				},
				resource.Attribute{
					Name:        "non_broadcast",
					Description: `Specify OSPF NBMA network`,
				},
				resource.Attribute{
					Name:        "p2mp_nbma",
					Description: `Specify non-broadcast point-to-multipoint network`,
				},
				resource.Attribute{
					Name:        "point_to_multipoint",
					Description: `Specify OSPF point-to-multipoint network`,
				},
				resource.Attribute{
					Name:        "point_to_point",
					Description: `Specify OSPF point-to-point network`,
				},
				resource.Attribute{
					Name:        "bind_type",
					Description: `‘inside’: This interface is connected to NPTv6 inside networks; ‘outside’: This interface is connected to NPTv6 outside networks;`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `NPTv6 domain name`,
				},
				resource.Attribute{
					Name:        "lockup_period_v6",
					Description: `Lockup period (second)`,
				},
				resource.Attribute{
					Name:        "lockup_v6",
					Description: `Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)`,
				},
				resource.Attribute{
					Name:        "normal_v6",
					Description: `Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit`,
				},
				resource.Attribute{
					Name:        "acl_name",
					Description: `Access-list Name`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘num_pkts’: Input packets; ‘num_total_bytes’: Input bytes; ‘num_unicast_pkts’: Received unicasts; ‘num_broadcast_pkts’: Received braodcasts; ‘num_multicast_pkts’: Received multicasts; ‘num_tx_pkts’: Transmitted packtes; ‘num_total_tx_bytes’: Transmitte bytes; ‘num_unicast_tx_pkts’: Trasnmitted unicasts; ‘num_broadcast_tx_pkts’: Transmitted broadcasts; ‘num_multicast_tx_pkts’: Transmitted multicasts; ‘rate_pkt_sent’: Packet sent rate packets/sec; ‘rate_byte_sent’: Byte sent rate bits/sec; ‘rate_pkt_rcvd’: Packet received rate packets/sec; ‘rate_byte_rcvd’: Byte received rate bits/sec; ‘load_interval’: Load Interval;`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `ACL applied on incoming packets to this interface`,
				},
				resource.Attribute{
					Name:        "ipv6_enable",
					Description: `Enable IPv6 processing`,
				},
				resource.Attribute{
					Name:        "v6_acl_name",
					Description: `Apply ACL rules to incoming packets on this interface (Named Access List)`,
				},
				resource.Attribute{
					Name:        "address_type",
					Description: `‘anycast’: Configure an IPv6 anycast address; ‘link-local’: Configure an IPv6 link local address;`,
				},
				resource.Attribute{
					Name:        "ipv6_addr",
					Description: `Set the IPv6 address of an interface`,
				},
				resource.Attribute{
					Name:        "rip",
					Description: `RIP Routing for IPv6`,
				},
				resource.Attribute{
					Name:        "area_id_addr",
					Description: `OSPF area ID in IP address format`,
				},
				resource.Attribute{
					Name:        "area_id_num",
					Description: `OSPF area ID as a decimal value`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "broadcast_type",
					Description: `‘broadcast’: Specify OSPF broadcast multi-access network; ‘non-broadcast’: Specify OSPF NBMA network; ‘point-to-point’: Specify OSPF point-to-point network; ‘point-to-multipoint’: Specify OSPF point-to-multipoint network;`,
				},
				resource.Attribute{
					Name:        "network_instance_id",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "neig_inst",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "neighbor",
					Description: `OSPFv3 neighbor (Neighbor IPv6 address)`,
				},
				resource.Attribute{
					Name:        "neighbor_cost",
					Description: `OSPF cost for point-to-multipoint neighbor (metric)`,
				},
				resource.Attribute{
					Name:        "neighbor_poll_interval",
					Description: `OSPF dead-router polling interval (Seconds)`,
				},
				resource.Attribute{
					Name:        "neighbor_priority",
					Description: `OSPF priority of non-broadcast neighbor`,
				},
				resource.Attribute{
					Name:        "adver_mtu",
					Description: `Set Router Advertisement MTU Option`,
				},
				resource.Attribute{
					Name:        "adver_mtu_disable",
					Description: `Disable Router Advertisement MTU Option`,
				},
				resource.Attribute{
					Name:        "adver_vrid",
					Description: `Vrid`,
				},
				resource.Attribute{
					Name:        "adver_vrid_default",
					Description: `Default VRRP-A vrid`,
				},
				resource.Attribute{
					Name:        "default_lifetime",
					Description: `Set Router Advertisement Default Lifetime (default: 1800) (Default Lifetime (seconds))`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "floating_ip_default_vrid",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "hop_limit",
					Description: `Set Router Advertisement Hop Limit (default: 255)`,
				},
				resource.Attribute{
					Name:        "managed_config_action",
					Description: `‘enable’: Enable the Managed Address Configuration flag; ‘disable’: Disable the Managed Address Configuration flag (default);`,
				},
				resource.Attribute{
					Name:        "max_interval",
					Description: `Set Router Advertisement Max Interval (default: 600) (Max Router Advertisement Interval (seconds))`,
				},
				resource.Attribute{
					Name:        "min_interval",
					Description: `Set Router Advertisement Min Interval (default: 200) (Min Router Advertisement Interval (seconds))`,
				},
				resource.Attribute{
					Name:        "other_config_action",
					Description: `‘enable’: Enable the Other Stateful Configuration flag; ‘disable’: Disable the Other Stateful Configuration flag (default);`,
				},
				resource.Attribute{
					Name:        "rate_limit",
					Description: `Rate Limit the processing of incoming Router Solicitations (Max Number of Router Solicitations to process per second)`,
				},
				resource.Attribute{
					Name:        "reachable_time",
					Description: `Set Router Advertisement Reachable ime (default: 0) (Reachable Time (milliseconds))`,
				},
				resource.Attribute{
					Name:        "retransmit_timer",
					Description: `Set Router Advertisement Retransmit Timer (default: 0)`,
				},
				resource.Attribute{
					Name:        "use_floating_ip",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "use_floating_ip_default_vrid",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "not_autonomous",
					Description: `Specify that the Prefix is not usable for autoconfiguration (default:autonomous)`,
				},
				resource.Attribute{
					Name:        "not_on_link",
					Description: `Specify that the Prefix is not On-Link (default: on-link)`,
				},
				resource.Attribute{
					Name:        "preferred_lifetime",
					Description: `Specify Prefix Preferred Lifetime (default:604800) (Prefix Advertised Preferred Lifetime (default: 604800))`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `Set Router Advertisement On-Link Prefix (IPv6 On-Link Prefix)`,
				},
				resource.Attribute{
					Name:        "valid_lifetime",
					Description: `Specify Valid Lifetime (default:2592000) (Prefix Advertised Valid Lifetime (default: 2592000))`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_interface_ve_bfd",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder interface ve bfd resource for A10`,
			Description:      ``,
			Keywords: []string{
				"interface",
				"ve",
				"bfd",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "demand",
					Description: `Demand mode`,
				},
				resource.Attribute{
					Name:        "echo",
					Description: `Enable BFD Echo`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `Key ID`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `‘md5’: Keyed MD5; ‘meticulous-md5’: Meticulous Keyed MD5; ‘meticulous-sha1’: Meticulous Keyed SHA1; ‘sha1’: Keyed SHA1; ‘simple’: Simple Password;`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Key String`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Transmit interval between BFD packets (Milliseconds)`,
				},
				resource.Attribute{
					Name:        "min_rx",
					Description: `Minimum receive interval capability (Milliseconds)`,
				},
				resource.Attribute{
					Name:        "multiplier",
					Description: `Multiplier value used to compute holddown (value used to multiply the interval)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_interface_ve_ip",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder interface ve ip resource for A10`,
			Description:      ``,
			Keywords: []string{
				"interface",
				"ve",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allow_promiscuous_vip",
					Description: `Allow traffic to be associated with promiscuous VIP`,
				},
				resource.Attribute{
					Name:        "client",
					Description: `Client facing interface for IPv4/v6 traffic`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Use DHCP to configure IP address`,
				},
				resource.Attribute{
					Name:        "generate_membership_query",
					Description: `Enable Membership Query`,
				},
				resource.Attribute{
					Name:        "inside",
					Description: `Inside (private) interface for stateful firewall`,
				},
				resource.Attribute{
					Name:        "max_resp_time",
					Description: `Maximum Response Time (Max Response Time (Default is 100))`,
				},
				resource.Attribute{
					Name:        "outside",
					Description: `Outside (public) interface for stateful firewall`,
				},
				resource.Attribute{
					Name:        "query_interval",
					Description: `1 - 255 (Default is 125)`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Server facing interface for IPv4/v6 traffic`,
				},
				resource.Attribute{
					Name:        "slb_partition_redirect",
					Description: `Redirect SLB traffic across partition`,
				},
				resource.Attribute{
					Name:        "ttl_ignore",
					Description: `Ignore TTL decrement for a received packet`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `IP address`,
				},
				resource.Attribute{
					Name:        "ipv4_netmask",
					Description: `IP subnet mask`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `ISO routing area tag`,
				},
				resource.Attribute{
					Name:        "helper_address",
					Description: `Helper address for DHCP packets (IP address)`,
				},
				resource.Attribute{
					Name:        "access_list",
					Description: `Access-list for traffic from the outside`,
				},
				resource.Attribute{
					Name:        "acl_id",
					Description: `ACL id`,
				},
				resource.Attribute{
					Name:        "class_list",
					Description: `Class List (Class List Name)`,
				},
				resource.Attribute{
					Name:        "receive_packet",
					Description: `Enable receiving packet through the specified interface`,
				},
				resource.Attribute{
					Name:        "send_packet",
					Description: `Enable sending packets through the specified interface`,
				},
				resource.Attribute{
					Name:        "receive",
					Description: `Advertisement reception`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `‘1’: RIP version 1; ‘2’: RIP version 2; ‘1-compatible’: RIPv1-compatible; ‘1-2’: RIP version 1 & 2;`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `‘poisoned’: Perform split horizon with poisoned reverse; ‘disable’: Disable split horizon; ‘enable’: Perform split horizon without poisoned reverse;`,
				},
				resource.Attribute{
					Name:        "key_chain",
					Description: `Authentication key-chain (Name of key-chain)`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `‘md5’: Keyed message digest; ‘text’: Clear text authentication;`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `The RIP authentication string`,
				},
				resource.Attribute{
					Name:        "send",
					Description: `Advertisement transmission`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Enable authentication`,
				},
				resource.Attribute{
					Name:        "authentication_key",
					Description: `Authentication password (key) (The OSPF password (key))`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `Interface cost`,
				},
				resource.Attribute{
					Name:        "database_filter",
					Description: `‘all’: Filter all LSA;`,
				},
				resource.Attribute{
					Name:        "dead_interval",
					Description: `Interval after which a neighbor is declared dead (Seconds)`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `Time between HELLO packets (Seconds)`,
				},
				resource.Attribute{
					Name:        "ip_addr",
					Description: `Address of interface`,
				},
				resource.Attribute{
					Name:        "mtu_ignore",
					Description: `Ignores the MTU in DBD packets`,
				},
				resource.Attribute{
					Name:        "out",
					Description: `Outgoing LSA`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Router priority`,
				},
				resource.Attribute{
					Name:        "retransmit_interval",
					Description: `Time between retransmitting lost link state advertisements (Seconds)`,
				},
				resource.Attribute{
					Name:        "transmit_delay",
					Description: `Link state transmit delay (Seconds)`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `‘message-digest’: Use message-digest authentication; ‘null’: Use no authentication;`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "md5_value",
					Description: `The OSPF password (1-16)`,
				},
				resource.Attribute{
					Name:        "message_digest_key",
					Description: `Message digest authentication password (key) (Key id)`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `Disable BFD`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `OSPF interface MTU (MTU size)`,
				},
				resource.Attribute{
					Name:        "broadcast",
					Description: `Specify OSPF broadcast multi-access network`,
				},
				resource.Attribute{
					Name:        "non_broadcast",
					Description: `Specify OSPF NBMA network`,
				},
				resource.Attribute{
					Name:        "p2mp_nbma",
					Description: `Specify non-broadcast point-to-multipoint network`,
				},
				resource.Attribute{
					Name:        "point_to_multipoint",
					Description: `Specify OSPF point-to-multipoint network`,
				},
				resource.Attribute{
					Name:        "point_to_point",
					Description: `Specify OSPF point-to-point network`,
				},
				resource.Attribute{
					Name:        "bfd",
					Description: `Bidirectional Forwarding Detection (BFD)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_interface_ve_ipv6",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder interface ve ipv6 resource for A10`,
			Description:      ``,
			Keywords: []string{
				"interface",
				"ve",
				"ipv6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ifnum",
					Description: `Interface no.`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `ACL applied on incoming packets to this interface`,
				},
				resource.Attribute{
					Name:        "inside",
					Description: `Inside (private) interface for stateful firewall`,
				},
				resource.Attribute{
					Name:        "ipv6_enable",
					Description: `Enable IPv6 processing`,
				},
				resource.Attribute{
					Name:        "outside",
					Description: `Outside (public) interface for stateful firewall`,
				},
				resource.Attribute{
					Name:        "ttl_ignore",
					Description: `Ignore TTL decrement for a received packet`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "v6_acl_name",
					Description: `Apply ACL rules to incoming packets on this interface (Named Access List)`,
				},
				resource.Attribute{
					Name:        "address_type",
					Description: `‘anycast’: Configure an IPv6 anycast address; ‘link-local’: Configure an IPv6 link local address;`,
				},
				resource.Attribute{
					Name:        "ipv6_addr",
					Description: `Set the IPv6 address of an interface`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `‘poisoned’: Perform split horizon with poisoned reverse; ‘disable’: Disable split horizon; ‘enable’: Perform split horizon without poisoned reverse;`,
				},
				resource.Attribute{
					Name:        "access_list",
					Description: `Access-list for traffic from the outside`,
				},
				resource.Attribute{
					Name:        "acl_name",
					Description: `Access-list Name`,
				},
				resource.Attribute{
					Name:        "class_list",
					Description: `Class List (Class List Name)`,
				},
				resource.Attribute{
					Name:        "rip",
					Description: `RIP Routing for IPv6`,
				},
				resource.Attribute{
					Name:        "area_id_addr",
					Description: `OSPF area ID in IP address format`,
				},
				resource.Attribute{
					Name:        "area_id_num",
					Description: `OSPF area ID as a decimal value`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `ISO routing area tag`,
				},
				resource.Attribute{
					Name:        "bfd",
					Description: `Bidirectional Forwarding Detection (BFD)`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `Disable BFD`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `Interface cost`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `Time between HELLO packets (Seconds)`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Router priority`,
				},
				resource.Attribute{
					Name:        "mtu_ignore",
					Description: `Ignores the MTU in DBD packets`,
				},
				resource.Attribute{
					Name:        "retransmit_interval",
					Description: `Time between retransmitting lost link state advertisements (Seconds)`,
				},
				resource.Attribute{
					Name:        "broadcast_type",
					Description: `‘broadcast’: Specify OSPF broadcast multi-access network; ‘non-broadcast’: Specify OSPF NBMA network; ‘point-to-point’: Specify OSPF point-to-point network; ‘point-to-multipoint’: Specify OSPF point-to-multipoint network;`,
				},
				resource.Attribute{
					Name:        "network_instance_id",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "p2mp_nbma",
					Description: `Specify non-broadcast point-to-multipoint network`,
				},
				resource.Attribute{
					Name:        "transmit_delay",
					Description: `Link state transmit delay (Seconds)`,
				},
				resource.Attribute{
					Name:        "neig_inst",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "neighbor",
					Description: `OSPFv3 neighbor (Neighbor IPv6 address)`,
				},
				resource.Attribute{
					Name:        "neighbor_cost",
					Description: `OSPF cost for point-to-multipoint neighbor (metric)`,
				},
				resource.Attribute{
					Name:        "neighbor_poll_interval",
					Description: `OSPF dead-router polling interval (Seconds)`,
				},
				resource.Attribute{
					Name:        "neighbor_priority",
					Description: `OSPF priority of non-broadcast neighbor`,
				},
				resource.Attribute{
					Name:        "dead_interval",
					Description: `Interval after which a neighbor is declared dead (Seconds)`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `‘enable’: Enable Router Advertisements on this interface; ‘disable’: Disable Router Advertisements on this interface;`,
				},
				resource.Attribute{
					Name:        "adver_mtu",
					Description: `Set Router Advertisement MTU Option`,
				},
				resource.Attribute{
					Name:        "adver_mtu_disable",
					Description: `Disable Router Advertisement MTU Option`,
				},
				resource.Attribute{
					Name:        "adver_vrid",
					Description: `Vrid`,
				},
				resource.Attribute{
					Name:        "adver_vrid_default",
					Description: `Default VRRP-A vrid`,
				},
				resource.Attribute{
					Name:        "default_lifetime",
					Description: `Set Router Advertisement Default Lifetime (default: 1800) (Default Lifetime (seconds))`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "floating_ip_default_vrid",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "hop_limit",
					Description: `Set Router Advertisement Hop Limit (default: 255)`,
				},
				resource.Attribute{
					Name:        "managed_config_action",
					Description: `‘enable’: Enable the Managed Address Configuration flag; ‘disable’: Disable the Managed Address Configuration flag (default);`,
				},
				resource.Attribute{
					Name:        "max_interval",
					Description: `Set Router Advertisement Max Interval (default: 600) (Max Router Advertisement Interval (seconds))`,
				},
				resource.Attribute{
					Name:        "min_interval",
					Description: `Set Router Advertisement Min Interval (default: 200) (Min Router Advertisement Interval (seconds))`,
				},
				resource.Attribute{
					Name:        "other_config_action",
					Description: `‘enable’: Enable the Other Stateful Configuration flag; ‘disable’: Disable the Other Stateful Configuration flag (default);`,
				},
				resource.Attribute{
					Name:        "rate_limit",
					Description: `Rate Limit the processing of incoming Router Solicitations (Max Number of Router Solicitations to process per second)`,
				},
				resource.Attribute{
					Name:        "reachable_time",
					Description: `Set Router Advertisement Reachable ime (default: 0) (Reachable Time (milliseconds))`,
				},
				resource.Attribute{
					Name:        "retransmit_timer",
					Description: `Set Router Advertisement Retransmit Timer (default: 0)`,
				},
				resource.Attribute{
					Name:        "use_floating_ip",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "use_floating_ip_default_vrid",
					Description: `Use a floating IP as the source address for Router advertisements`,
				},
				resource.Attribute{
					Name:        "not_autonomous",
					Description: `Specify that the Prefix is not usable for autoconfiguration (default:autonomous)`,
				},
				resource.Attribute{
					Name:        "not_on_link",
					Description: `Specify that the Prefix is not On-Link (default: on-link)`,
				},
				resource.Attribute{
					Name:        "preferred_lifetime",
					Description: `Specify Prefix Preferred Lifetime (default:604800) (Prefix Advertised Preferred Lifetime (default: 604800))`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `Set Router Advertisement On-Link Prefix (IPv6 On-Link Prefix)`,
				},
				resource.Attribute{
					Name:        "valid_lifetime",
					Description: `Specify Valid Lifetime (default:2592000) (Prefix Advertised Valid Lifetime (default: 2592000))`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ip_address",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ip address resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"address",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_addr",
					Description: `IP address`,
				},
				resource.Attribute{
					Name:        "ip_mask",
					Description: `IP subnet mask`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ip_dns_primary",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ip dns primary resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"dns",
				"primary",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_v4_addr",
					Description: `DNS server address`,
				},
				resource.Attribute{
					Name:        "ip_v6_addr",
					Description: `DNS server address`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ip_dns_secondary",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ip dns secondary resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"dns",
				"secondary",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_v4_addr",
					Description: `DNS server address`,
				},
				resource.Attribute{
					Name:        "ip_v6_addr",
					Description: `DNS server address`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ip_dns_suffix",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ip dns suffix resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"dns",
				"suffix",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name",
					Description: `DNS suffix`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ip_frag",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ip frag resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"frag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "buff",
					Description: `Max buff used for fragmentation (Buffer Value(10000-3000000))`,
				},
				resource.Attribute{
					Name:        "max_packets_per_reassembly",
					Description: `Max number of fragmented packets allowed per reassembly(0 is unlimited) (default 0)`,
				},
				resource.Attribute{
					Name:        "max_reassembly_sessions",
					Description: `Max number of pending reassembly sessions allowed (default 100000)`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Fragmentation timeout (in milliseconds 4 - 16000 (default is 1000))`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘session-inserted’: Session Inserted; ‘session-expired’: Session Expired; ‘icmp-rcv’: ICMP Received; ‘icmpv6-rcv’: ICMPv6 Received; ‘udp-rcv’: UDP Received; ‘tcp-rcv’: TCP Received; ‘ipip-rcv’: IP-in-IP Received; ‘ipv6ip-rcv’: IPv6-in-IP Received; ‘other-rcv’: Other Received; ‘icmp-dropped’: ICMP Dropped; ‘icmpv6-dropped’: ICMPv6 Dropped; ‘udp-dropped’: UDP Dropped; ‘tcp-dropped’: TCP Dropped; ‘ipip-dropped’: IP-in-IP Dropped; ‘ipv6ip-dropped’: IPv6-in-IP Dropped; ‘other-dropped’: Other Dropped; ‘overlap-error’: Overlapping Fragment Dropped; ‘bad-ip-len’: Bad IP Length; ‘too-small’: Fragment Too Small Drop; ‘first-tcp-too-small’: First TCP Fragment Too Small Drop; ‘first-l4-too-small’: First L4 Fragment Too Small Drop; ‘total-sessions-exceeded’: Total Sessions Exceeded Drop; ‘no-session-memory’: Out of Session Memory; ‘fast-aging-set’: Fragmentation Fast Aging Set; ‘fast-aging-unset’: Fragmentation Fast Aging Unset; ‘fragment-queue-success’: Fragment Queue Success; ‘unaligned-len’: Payload Length Unaligned; ‘exceeded-len’: Payload Length Out of Bounds; ‘duplicate-first-frag’: Duplicate First Fragment; ‘duplicate-last-frag’: Duplicate Last Fragment; ‘total-fragments-exceeded’: Total Queued Fragments Exceeded; ‘fragment-queue-failure’: Fragment Queue Failure; ‘reassembly-success’: Fragment Reassembly Success; ‘max-len-exceeded’: Fragment Max Data Length Exceeded; ‘reassembly-failure’: Fragment Reassembly Failure; ‘policy-drop’: MTU Exceeded Policy Drop; ‘error-drop’: Fragment Processing Drop; ‘max-packets-exceeded’: Too Many Packets Per Reassembly Drop; ‘session-packets-exceeded’: Session Max Packets Exceeded; ‘frag-session-count’: Fragmentation Session Count; ‘high-cpu-threshold’: High CPU Threshold Reached; ‘low-cpu-threshold’: Low CPU Threshold Reached; ‘cpu-threshold-drop’: High CPU Drop; ‘ipd-entry-drop’: DDoS Protection Drop; ‘sctp-rcv’: SCTP Received; ‘sctp-dropped’: SCTP Dropped;`,
				},
				resource.Attribute{
					Name:        "high",
					Description: `When CPU usage reaches this value, it will stop processing fragments (default: 75%)`,
				},
				resource.Attribute{
					Name:        "low",
					Description: `When CPU usage remains under this value, it will resume processing fragments (default: 60%)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ip_icmp",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ip icmp resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"icmp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "redirect",
					Description: `Disable outbound ICMP redirect messages`,
				},
				resource.Attribute{
					Name:        "unreachable",
					Description: `Disable outbound ICMP unreachable messages`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ip_nat_alg_pptp",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ip nat alg pptp resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"nat",
				"alg",
				"pptp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pptp",
					Description: `‘disable’: Disable PPTP NAT ALG; ‘enable’: Enable PPTP NAT ALG;`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘current-smp-sessions’: current-smp-sessions; ‘current-gre-sessions’: current-gre-sessions; ‘smp-session-creation-failure’: smp-session-creation-failure; ‘truncated-pns-message’: truncated-pns-message; ‘truncated-pac-message’: truncated-pac-message; ‘mismatched-pns-call-id’: mismatched-pns-call-id; ‘mismatched-pac-call-id’: mismatched-pac-call-id; ‘retransmitted-pns-message’: retransmitted-pns-message; ‘retransmitted-pac-message’: retransmitted-pac-message; ‘truncated-gre-packet’: truncated-gre-packet; ‘unknown-gre-version’: unknown-gre-version; ‘no-matching-gre-session’: no-matching-gre-session;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ip_nat_global",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ip nat global resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"nat",
				"global",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "reset_idle_tcp_conn",
					Description: `Reset Idle TCP Connections`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ip_nat_icmp",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ip nat icmp resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"nat",
				"icmp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "always_source_nat_errors",
					Description: `Source NAT intermediate routers’ IPs for ICMP errors (default: disabled)`,
				},
				resource.Attribute{
					Name:        "respond_to_ping",
					Description: `Respond to ICMP echo requests to NAT pool IPs (default: disabled)`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ip_prefix_list",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ip prefix list resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"prefix",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of a prefix list`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `‘deny’: Specify packets to reject; ‘permit’: Specify packets to forward;`,
				},
				resource.Attribute{
					Name:        "any",
					Description: `Any prefix match. Same as “0.0.0.0/0 le 32”`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Prefix-list specific description (Up to 80 characters describing this prefix-list)`,
				},
				resource.Attribute{
					Name:        "ge",
					Description: `Minimum prefix length to be matched`,
				},
				resource.Attribute{
					Name:        "ipaddr",
					Description: `IP prefix, e.g., 35.0.0.0/8`,
				},
				resource.Attribute{
					Name:        "le",
					Description: `Maximum prefix length to be matched`,
				},
				resource.Attribute{
					Name:        "seq",
					Description: `Sequence number of an entry`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ip_reroute",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ip reroute resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"reroute",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "ebgp",
					Description: `EBGP`,
				},
				resource.Attribute{
					Name:        "ibgp",
					Description: `IBGP`,
				},
				resource.Attribute{
					Name:        "isis",
					Description: `ISIS`,
				},
				resource.Attribute{
					Name:        "ospf",
					Description: `OSPF`,
				},
				resource.Attribute{
					Name:        "rip",
					Description: `RIP`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ip_route_static_bfd",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ip route static bfd resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"route",
				"static",
				"bfd",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "local_ip",
					Description: `Local IP address`,
				},
				resource.Attribute{
					Name:        "nexthop_ip",
					Description: `Nexthop IP address`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ip_tcp",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ip tcp resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"tcp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `SYN cookie expire threshold (seconds (default is 4))`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ipv6_frag",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ipv6 frag resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ipv6",
				"frag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "frag_timeout",
					Description: `in milliseconds 4 - 16000 (default is 1000)`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘session-inserted’: Session Inserted; ‘session-expired’: Session Expired; ‘icmp-rcv’: ICMP Received; ‘icmpv6-rcv’: ICMPv6 Received; ‘udp-rcv’: UDP Received; ‘tcp-rcv’: TCP Received; ‘ipip-rcv’: IP-in-IP Received; ‘ipv6ip-rcv’: IPv6-in-IP Received; ‘other-rcv’: Other Received; ‘icmp-dropped’: ICMP Dropped; ‘icmpv6-dropped’: ICMPv6 Dropped; ‘udp-dropped’: UDP Dropped; ‘tcp-dropped’: TCP Dropped; ‘ipip-dropped’: IP-in-IP Dropped; ‘ipv6ip-dropped’: IPv6-in-IP Dropped; ‘other-dropped’: Other Dropped; ‘overlap-error’: Overlapping Fragment Dropped; ‘bad-ip-len’: Bad IP Length; ‘too-small’: Fragment Too Small Drop; ‘first-tcp-too-small’: First TCP Fragment Too Small Drop; ‘first-l4-too-small’: First L4 Fragment Too Small Drop; ‘total-sessions-exceeded’: Total Sessions Exceeded Drop; ‘no-session-memory’: Out of Session Memory; ‘fast-aging-set’: Fragmentation Fast Aging Set; ‘fast-aging-unset’: Fragmentation Fast Aging Unset; ‘fragment-queue-success’: Fragment Queue Success; ‘unaligned-len’: Payload Length Unaligned; ‘exceeded-len’: Payload Length Out of Bounds; ‘duplicate-first-frag’: Duplicate First Fragment; ‘duplicate-last-frag’: Duplicate Last Fragment; ‘total-fragments-exceeded’: Total Queued Fragments Exceeded; ‘fragment-queue-failure’: Fragment Queue Failure; ‘reassembly-success’: Fragment Reassembly Success; ‘max-len-exceeded’: Fragment Max Data Length Exceeded; ‘reassembly-failure’: Fragment Reassembly Failure; ‘policy-drop’: MTU Exceeded Policy Drop; ‘error-drop’: Fragment Processing Drop; ‘max-packets-exceeded’: Too Many Packets Per Reassembly Drop; ‘session-packets-exceeded’: Session Max Packets Exceeded; ‘frag-session-count’: Fragmentation Session Count; ‘high-cpu-threshold’: High CPU Threshold Reached; ‘low-cpu-threshold’: Low CPU Threshold Reached; ‘cpu-threshold-drop’: High CPU Drop; ‘ipd-entry-drop’: DDoS Protection Drop; ‘sctp-rcv’: SCTP Received; ‘sctp-dropped’: SCTP Dropped;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ipv6_icmpv6",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ipv6 icmpv6 resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ipv6",
				"icmpv6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "redirect",
					Description: `Disable outbound ICMPv6 redirect messages`,
				},
				resource.Attribute{
					Name:        "unreachable",
					Description: `Disable outbound ICMPv6 unreachable messages`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_ipv6_nat_icmpv6",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder ipv6 nat icmpv6 resource for A10`,
			Description:      ``,
			Keywords: []string{
				"ipv6",
				"nat",
				"icmpv6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "respond_to_ping",
					Description: `Respond to ICMPv6 echo requests to NAT pool IPs (default: disabled)`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_overlay_tunnel_options",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder overlay tunnel options resource for A10`,
			Description:      ``,
			Keywords: []string{
				"overlay",
				"tunnel",
				"options",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tcp_mss_adjust_disable",
					Description: `Disable TCP MSS adjustment in SYN packet for tunnels`,
				},
				resource.Attribute{
					Name:        "nvgre_disable_flow_id",
					Description: `Disable Flow-ID computation for NVGRE`,
				},
				resource.Attribute{
					Name:        "ip_dscp_preserve",
					Description: `Copy DSCP bits from inner IP to outer IP header`,
				},
				resource.Attribute{
					Name:        "vxlan_dest_port",
					Description: `VXLAN UDP Destination Port (UDP Port Number (default 4789))`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "gateway_mac",
					Description: `MAC to be used with Gateway segment Id (MAC Address for the Gateway segment)`,
				},
				resource.Attribute{
					Name:        "nvgre_key_mode_lower24",
					Description: `Use the lower 24-bits of the GRE key as the VSID`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_overlay_tunnel_vtep",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder overlay tunnel vtep resource for A10`,
			Description:      ``,
			Keywords: []string{
				"overlay",
				"tunnel",
				"vtep",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `Source Tunnel End Point IPv4 address`,
				},
				resource.Attribute{
					Name:        "segment",
					Description: `VNI configured for the remote VTEP`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `Logical interface binding the Provider Partition to the User Partition (logical interface number)`,
				},
				resource.Attribute{
					Name:        "lif",
					Description: `Name of the Partition with the L2 segment being extended (Name of the User Partition with the L2 segment being extended)`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `'nvgre’: Tunnel Encapsulation Type is NVGRE; 'vxlan’: Tunnel Encapsulation Type is VXLAN;`,
				},
				resource.Attribute{
					Name:        "ip_addr",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "destination_vtep",
					Description: `Configure the VTEP IP address (IPv4 address of the VTEP for the remote host)`,
				},
				resource.Attribute{
					Name:        "overlay_mac_addr",
					Description: `MAC Address of the overlay host`,
				},
				resource.Attribute{
					Name:        "vni",
					Description: `Configure the segment id ( VNI of the remote host)`,
				},
				resource.Attribute{
					Name:        "id2",
					Description: `Disable TCP MSS adjustment in SYN packet for tunnels`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Disable Flow-ID computation for NVGRE`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IPv4 address of the overlay host`,
				},
				resource.Attribute{
					Name:        "segment",
					Description: `Configure the segment id ( VNI of the remote host)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_partition",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder partition resource for A10`,
			Description:      ``,
			Keywords: []string{
				"partition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "partition_name",
					Description: `Object partition name`,
				},
				resource.Attribute{
					Name:        "application_type",
					Description: `'adc’: Application type ADC; 'cgnv6’: Application type CGNv6;`,
				},
				resource.Attribute{
					Name:        "id2",
					Description: `Specify unique Partition id`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_reboot",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder reboot resource for A10`,
			Description:      ``,
			Keywords: []string{
				"reboot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all",
					Description: `Reboot all devices when VCS is enabled, or only this device itself if VCS is not enabled`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_rib_route",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder rib route resource for A10`,
			Description:      ``,
			Keywords: []string{
				"rib",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_dest_addr",
					Description: `Destination ip address`,
				},
				resource.Attribute{
					Name:        "ip_mask",
					Description: `Ip address subnet mask`,
				},
				resource.Attribute{
					Name:        "ip_next_hop",
					Description: `Ip next hop`,
				},
				resource.Attribute{
					Name:        "distance_nexthop_ip",
					Description: `Distance of next hop ip address`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_server",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder server resource for A10`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "health_check_disable",
					Description: `Disable configured health check configuration`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Server name`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Server host IP Address`,
				},
				resource.Attribute{
					Name:        "port_number",
					Description: `Port number`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `protocol`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_service_group",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder service group resource for A10`,
			Description:      ``,
			Keywords: []string{
				"service",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of service group`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `protocol`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_aflow",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB aflow resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"aflow",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'hits’: Cache hits; 'miss’: Cache misses; 'bytes_served’: Bytes served from cache; 'total_req’: Total requests received; 'caching_req’: Total requests to cache; 'nc_req_header’: nc_req_header; 'nc_res_header’: nc_res_header; 'rv_success’: rv_success; 'rv_failure’: rv_failure; 'ims_request’: ims_request; 'nm_response’: nm_response; 'rsp_type_CL’: rsp_type_CL; 'rsp_type_CE’: rsp_type_CE; 'rsp_type_304’: rsp_type_304; 'rsp_type_other’: rsp_type_other; 'rsp_no_compress’: rsp_no_compress; 'rsp_gzip’: rsp_gzip; 'rsp_deflate’: rsp_deflate; 'rsp_other’: rsp_other; 'nocache_match’: nocache_match; 'match’: match; 'invalidate_match’: invalidate_match; 'content_toobig’: content_toobig; 'content_toosmall’: content_toosmall; 'entry_create_failures’: entry_create_failures; 'mem_size’: mem_size; 'entry_num’: entry_num; 'replaced_entry’: replaced_entry; 'aging_entry’: aging_entry; 'cleaned_entry’: cleaned_entry;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_common",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB common resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"common",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "low_latency",
					Description: `Enable low latency mode`,
				},
				resource.Attribute{
					Name:        "use_mss_tab",
					Description: `Use MSS based on internal table for SLB processing`,
				},
				resource.Attribute{
					Name:        "stats_data_disable",
					Description: `Disable global slb data statistics`,
				},
				resource.Attribute{
					Name:        "compress_block_size",
					Description: `Set compression block size (Compression block size in bytes)`,
				},
				resource.Attribute{
					Name:        "player_id_check_enable",
					Description: `Enable the Player id check`,
				},
				resource.Attribute{
					Name:        "msl_time",
					Description: `Configure maximum session life, default is 2 seconds (1-40 seconds, default is 2 seconds)`,
				},
				resource.Attribute{
					Name:        "graceful_shutdown_enable",
					Description: `Enable graceful shutdown`,
				},
				resource.Attribute{
					Name:        "hw_syn_rr",
					Description: `Configure hardware SYN round robin (range 1-500000)`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit",
					Description: ``,
				},
				resource.Attribute{
					Name:        "src_ip_list",
					Description: ``,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `'tcp’: Set TCP connection rate limit; 'udp’: Set UDP packet rate limit;`,
				},
				resource.Attribute{
					Name:        "limit_period",
					Description: `'100’: 100 ms; '1000’: 1000 ms;`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `Set max connections per period`,
				},
				resource.Attribute{
					Name:        "exceed_action",
					Description: `Set action if threshold exceeded`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Set threshold shared amongst all virtual ports`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_common_conn_rate_limit_src_ip",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB common conn-rate-limit src-ip resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"common",
				"conn",
				"rate",
				"limit",
				"src",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "protocol",
					Description: `'tcp’: Set TCP connection rate limit; 'udp’: Set UDP packet rate limit;`,
				},
				resource.Attribute{
					Name:        "limit_period",
					Description: `'100’: 100 ms; '1000’: 1000 ms;`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `Set max connections per period`,
				},
				resource.Attribute{
					Name:        "exceed_action",
					Description: `Set action if threshold exceeded`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Set threshold shared amongst all virtual ports`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_dns",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB dns resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"dns",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'hits’: Cache hits; 'miss’: Cache misses; 'bytes_served’: Bytes served from cache; 'total_req’: Total requests received; 'caching_req’: Total requests to cache; 'nc_req_header’: nc_req_header; 'nc_res_header’: nc_res_header; 'rv_success’: rv_success; 'rv_failure’: rv_failure; 'ims_request’: ims_request; 'nm_response’: nm_response; 'rsp_type_CL’: rsp_type_CL; 'rsp_type_CE’: rsp_type_CE; 'rsp_type_304’: rsp_type_304; 'rsp_type_other’: rsp_type_other; 'rsp_no_compress’: rsp_no_compress; 'rsp_gzip’: rsp_gzip; 'rsp_deflate’: rsp_deflate; 'rsp_other’: rsp_other; 'nocache_match’: nocache_match; 'match’: match; 'invalidate_match’: invalidate_match; 'content_toobig’: content_toobig; 'content_toosmall’: content_toosmall; 'entry_create_failures’: entry_create_failures; 'mem_size’: mem_size; 'entry_num’: entry_num; 'replaced_entry’: replaced_entry; 'aging_entry’: aging_entry; 'cleaned_entry’: cleaned_entry;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_dns_cache",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB dns-cache resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"dns",
				"cache",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'hits’: Cache hits; 'miss’: Cache misses; 'bytes_served’: Bytes served from cache; 'total_req’: Total requests received; 'caching_req’: Total requests to cache; 'nc_req_header’: nc_req_header; 'nc_res_header’: nc_res_header; 'rv_success’: rv_success; 'rv_failure’: rv_failure; 'ims_request’: ims_request; 'nm_response’: nm_response; 'rsp_type_CL’: rsp_type_CL; 'rsp_type_CE’: rsp_type_CE; 'rsp_type_304’: rsp_type_304; 'rsp_type_other’: rsp_type_other; 'rsp_no_compress’: rsp_no_compress; 'rsp_gzip’: rsp_gzip; 'rsp_deflate’: rsp_deflate; 'rsp_other’: rsp_other; 'nocache_match’: nocache_match; 'match’: match; 'invalidate_match’: invalidate_match; 'content_toobig’: content_toobig; 'content_toosmall’: content_toosmall; 'entry_create_failures’: entry_create_failures; 'mem_size’: mem_size; 'entry_num’: entry_num; 'replaced_entry’: replaced_entry; 'aging_entry’: aging_entry; 'cleaned_entry’: cleaned_entry;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_dns_response_rate_limiting",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB dns-response-rate-limiting resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"dns",
				"response",
				"rate",
				"limiting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'hits’: Cache hits; 'miss’: Cache misses; 'bytes_served’: Bytes served from cache; 'total_req’: Total requests received; 'caching_req’: Total requests to cache; 'nc_req_header’: nc_req_header; 'nc_res_header’: nc_res_header; 'rv_success’: rv_success; 'rv_failure’: rv_failure; 'ims_request’: ims_request; 'nm_response’: nm_response; 'rsp_type_CL’: rsp_type_CL; 'rsp_type_CE’: rsp_type_CE; 'rsp_type_304’: rsp_type_304; 'rsp_type_other’: rsp_type_other; 'rsp_no_compress’: rsp_no_compress; 'rsp_gzip’: rsp_gzip; 'rsp_deflate’: rsp_deflate; 'rsp_other’: rsp_other; 'nocache_match’: nocache_match; 'match’: match; 'invalidate_match’: invalidate_match; 'content_toobig’: content_toobig; 'content_toosmall’: content_toosmall; 'entry_create_failures’: entry_create_failures; 'mem_size’: mem_size; 'entry_num’: entry_num; 'replaced_entry’: replaced_entry; 'aging_entry’: aging_entry; 'cleaned_entry’: cleaned_entry;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_fast_http_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB fast-http-proxy resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"fast",
				"http",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'num’: Num; 'curr_proxy’: Curr Proxy Conns; 'total_proxy’: Total Proxy Conns; 'req’: HTTP requests; 'req_succ’: HTTP requests(succ); 'noproxy’: No proxy error; 'client_rst’: Client RST; 'server_rst’: Server RST; 'notuple’: No tuple error; 'parsereq_fail’: Parse req fail; 'svrsel_fail’: Server selection fail; 'fwdreq_fail’: Fwd req fail; 'fwdreq_fail_buff’: Fwd req fail - buff; 'fwdreq_fail_rport’: Fwd req fail - rport; 'fwdreq_fail_route’: Fwd req fail - route; 'fwdreq_fail_persist’: Fwd req fail - persist; 'fwdreq_fail_server’: Fwd req fail - server; 'fwdreq_fail_tuple’: Fwd req fail - tuple; 'fwdreqdata_fail’: Fwd req data fail; 'req_retran’: Packets retrans; 'req_ofo’: Packets ofo; 'server_resel’: Server reselection; 'svr_prem_close’: Server premature close; 'new_svrconn’: Server conn made; 'snat_fail’: Source NAT failure; 'tcpoutrst’: Out RSTs; 'full_proxy’: Full proxy tot; 'full_proxy_post’: Full proxy POST; 'full_proxy_pipeline’: Full proxy pipeline; 'full_proxy_fpga_err’: Full proxy fpga err; 'req_over_limit’: Request over limit; 'req_rate_over_limit’: Request rate over limit; 'l4_switching’: L4 switching; 'cookie_switching’: Cookie switching; 'aflex_switching’: aFleX switching; 'http_policy_switching’: HTTP Policy switching; 'url_switching’: URL switching; 'host_switching’: Host switching; 'lb_switching’: Normal LB switching; 'l4_switching_ok’: L4 switching (succ); 'cookie_switching_ok’: Cookie switching (succ); 'aflex_switching_ok’: aFleX switching (succ); 'http_policy_switching_ok’: HTTP Policy switching (succ); 'url_switching_ok’: URL switching (succ); 'host_switching_ok’: Host switching (succ); 'lb_switching_ok’: Normal LB switch. (succ); 'l4_switching_enqueue’: L4 switching (enQ); 'cookie_switching_enqueue’: Cookie switching (enQ); 'aflex_switching_enqueue’: aFleX switching (enQ); 'http_policy_switching_enqueue’: HTTP Policy switching (enQ); 'url_switching_enqueue’: URL switching (enQ); 'host_switching_enqueue’: Host switching (enQ); 'lb_switching_enqueue’: Normal LB switch. (enQ); 'retry_503’: Retry on 503; 'aflex_retry’: aFleX http retry; 'aflex_lb_reselect’: aFleX lb reselect; 'aflex_lb_reselect_ok’: aFleX lb reselect (succ); 'client_rst_request’: Client RST - request; 'client_rst_connecting’: Client RST - connecting; 'client_rst_connected’: Client RST - connected; 'client_rst_response’: Client RST - response; 'server_rst_request’: Server RST - request; 'server_rst_connecting’: Server RST - connecting; 'server_rst_connected’: Server RST - connected; 'server_rst_response’: Server RST - response; 'invalid_header’: Invalid header; 'too_many_headers’: Too many headers; 'line_too_long’: Line too long; 'header_name_too_long’: Header name too long; 'wrong_resp_header’: Wrong response header; 'header_insert’: Header insert; 'header_delete’: Header delete; 'insert_client_ip’: Insert client IP; 'negative_req_remain’: Negative request remain; 'negative_resp_remain’: Negative response remain; 'large_cookie’: Large cookies; 'large_cookie_header’: Large cookie headers; 'huge_cookie’: Huge cookies; 'huge_cookie_header’: Huge cookie headers; 'parse_cookie_fail’: Parse cookie fail; 'parse_setcookie_fail’: Parse set-cookie fail; 'asm_cookie_fail’: Assemble cookie fail; 'asm_cookie_header_fail’: Asm cookie header fail; 'asm_setcookie_fail’: Assemble set-cookie fail; 'asm_setcookie_header_fail’: Asm set-cookie hdr fail; 'client_req_unexp_flag’: Client req unexp flags; 'connecting_fin’: Connecting FIN; 'connecting_fin_retrans’: Connecting FIN retran; 'connecting_fin_ofo’: Connecting FIN ofo; 'connecting_rst’: Connecting RST; 'connecting_rst_retrans’: Connecting RST retran; 'connecting_rst_ofo’: Connecting RST ofo; 'connecting_ack’: Connecting ACK; 'pkts_ofo’: Packets ofo; 'pkts_retrans’: Packets retrans; 'pkts_retrans_ack_finwait’: retrans ACK FWAIT; 'pkts_retrans_fin’: retrans FIN; 'pkts_retrans_rst’: retrans RST; 'pkts_retrans_push’: retrans PSH; 'stale_sess’: Stale sess; 'server_resel_failed’: Server re-select failed; 'compression_before’: Tot data before compress; 'compression_after’: Tot data after compress; 'response_1xx’: Status code 1XX; 'response_100’: Status code 100; 'response_101’: Status code 101; 'response_102’: Status code 102; 'response_2xx’: Status code 2XX; 'response_200’: Status code 200; 'response_201’: Status code 201; 'response_202’: Status code 202; 'response_203’: Status code 203; 'response_204’: Status code 204; 'response_205’: Status code 205; 'response_206’: Status code 206; 'response_207’: Status code 207; 'response_3xx’: Status code 3XX; 'response_300’: Status code 300; 'response_301’: Status code 301; 'response_302’: Status code 302; 'response_303’: Status code 303; 'response_304’: Status code 304; 'response_305’: Status code 305; 'response_306’: Status code 306; 'response_307’: Status code 307; 'response_4xx’: Status code 4XX; 'response_400’: Status code 400; 'response_401’: Status code 401; 'response_402’: Status code 402; 'response_403’: Status code 403; 'response_404’: Status code 404; 'response_405’: Status code 405; 'response_406’: Status code 406; 'response_407’: Status code 407; 'response_408’: Status code 408; 'response_409’: Status code 409; 'response_410’: Status code 410; 'response_411’: Status code 411; 'response_412’: Status code 412; 'response_413’: Status code 413; 'response_414’: Status code 414; 'response_415’: Status code 415; 'response_416’: Status code 416; 'response_417’: Status code 417; 'response_418’: Status code 418; 'response_422’: Status code 422; 'response_423’: Status code 423; 'response_424’: Status code 424; 'response_425’: Status code 425; 'response_426’: Status code 426; 'response_449’: Status code 449; 'response_450’: Status code 450; 'response_5xx’: Status code 5XX; 'response_500’: Status code 500; 'response_501’: Status code 501; 'response_502’: Status code 502; 'response_503’: Status code 503; 'response_504’: Status code 504; 'response_505’: Status code 505; 'response_506’: Status code 506; 'response_507’: Status code 507; 'response_508’: Status code 508; 'response_509’: Status code 509; 'response_510’: Status code 510; 'response_6xx’: Status code 6XX; 'response_unknown’: Status code unknown; 'req_http10’: Request 1.0; 'req_http11’: Request 1.1; 'response_http10’: Resp 1.0; 'response_http11’: Resp 1.1; 'req_get’: Method GET; 'req_head’: Method HEAD; 'req_put’: Method PUT; 'req_post’: Method POST; 'req_trace’: Method TRACE; 'req_options’: Method OPTIONS; 'req_connect’: Method CONNECT; 'req_delete’: Method DELETE; 'req_unknown’: Method UNKNOWN; 'req_content_len’: Req content len; 'rsp_content_len’: Resp content len; 'rsp_chunk’: Resp chunk encoding; 'req_chunk’: Req chunk encoding; 'compress_rsp’: Compress req; 'compress_del_accept_enc’: Compress del accept enc; 'compress_resp_already_compressed’: Resp already compressed; 'compress_content_type_excluded’: Compress cont type excl; 'compress_no_content_type’: Compress no cont type; 'compress_resp_lt_min’: Compress resp less than min; 'compress_resp_no_cl_or_ce’: Compress resp no CL/CE; 'compress_ratio_too_high’: Compress ratio too high; 'cache_rsp’: HTTP req (cache succ); 'close_on_ddos’: Close on DDoS; 'req_http10_keepalive’: 1.0 Keepalive; 'req_sz_1k’: Req less than equal to 1K; 'req_sz_2k’: Req less than equal to 2K;`,
				},
				resource.Attribute{
					Name:        "counters2",
					Description: `'req_sz_4k’: Req less than equal to 4K; 'req_sz_8k’: Req less than equal to 8K; 'req_sz_16k’: Req less than equal to 16K; 'req_sz_32k’: Req less than equal to 32K; 'req_sz_64k’: Req less than equal to 64K; 'req_sz_256k’: Req less than equal to 256K; 'req_sz_gt_256k’: Req greater than 256K; 'rsp_sz_1k’: Resp less than equal to 1K; 'rsp_sz_2k’: Resp less than equal to 2K; 'rsp_sz_4k’: Resp less than equal to 4K; 'rsp_sz_8k’: Resp less than equal to 8K; 'rsp_sz_16k’: Resp less than equal to 16K; 'rsp_sz_32k’: Resp less than equal to 32K; 'rsp_sz_64k’: Resp less than equal to 64K; 'rsp_sz_256k’: Resp less than equal to 256K; 'rsp_sz_gt_256k’: Resp greater than 256K; 'chunk_sz_512’: Chunk less than equal to 512; 'chunk_sz_1k’: Chunk less than equal to 1K; 'chunk_sz_2k’: Chunk less than equal to 2K; 'chunk_sz_4k’: Chunk less than equal to 4K; 'chunk_sz_gt_4k’: Chunk greater than 4K; 'pconn_connecting’: pconn connecting; 'pconn_connected’: pconn connected; 'pconn_connecting_failed’: pconn conn failed; 'chunk_bad’: Bad Chunk; 'req_10u’: Rsp time less than 10u; 'req_20u’: Rsp time less than 20u; 'req_50u’: Rsp time less than 50u; 'req_100u’: Rsp time less than 100u; 'req_200u’: Rsp time less than 200u; 'req_500u’: Rsp time less than 500u; 'req_1m’: Rsp time less than 1m; 'req_2m’: Rsp time less than 2m; 'req_5m’: Rsp time less than 5m; 'req_10m’: Rsp time less than 10m; 'req_20m’: Rsp time less than 20m; 'req_50m’: Rsp time less than 50m; 'req_100m’: Rsp time less than 100m; 'req_200m’: Rsp time less than 200m; 'req_500m’: Rsp time less than 500m; 'req_1s’: Rsp time less than 1s; 'req_2s’: Rsp time less than 2s; 'req_5s’: Rsp time less than 5s; 'req_over_5s’: Rsp time greater than equal to 5s; 'insert_client_port’: Insert client Port; 'req_track’: Method TRACK; 'full_proxy_put’: Full proxy PUT; 'non_http_bypass’: Non-HTTP bypass; 'skip_insert_client_ip’: Skip Insert Client IP; 'skip_insert_client_port’: Skip Insert Client Port; 'decompression_before’: Tot data before decompress; 'decompression_after’: Tot data after decompress; 'http_pkts_in_seq’: Tot In-seq fHTTP packets; 'http_pkts_retx’: Tot Re-Tx fHTTP packets; 'http_client_retx’: Client Re-Tx fHTTP packets; 'http_server_retx’: Server Re-Tx fHTTP packets; 'http_pkts_ofo’: fHTTP Out of Order packets;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_fix",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB fix resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"fix",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'hits’: Cache hits; 'miss’: Cache misses; 'bytes_served’: Bytes served from cache; 'total_req’: Total requests received; 'caching_req’: Total requests to cache; 'nc_req_header’: nc_req_header; 'nc_res_header’: nc_res_header; 'rv_success’: rv_success; 'rv_failure’: rv_failure; 'ims_request’: ims_request; 'nm_response’: nm_response; 'rsp_type_CL’: rsp_type_CL; 'rsp_type_CE’: rsp_type_CE; 'rsp_type_304’: rsp_type_304; 'rsp_type_other’: rsp_type_other; 'rsp_no_compress’: rsp_no_compress; 'rsp_gzip’: rsp_gzip; 'rsp_deflate’: rsp_deflate; 'rsp_other’: rsp_other; 'nocache_match’: nocache_match; 'match’: match; 'invalidate_match’: invalidate_match; 'content_toobig’: content_toobig; 'content_toosmall’: content_toosmall; 'entry_create_failures’: entry_create_failures; 'mem_size’: mem_size; 'entry_num’: entry_num; 'replaced_entry’: replaced_entry; 'aging_entry’: aging_entry; 'cleaned_entry’: cleaned_entry;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_ftp_ctl",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB ftp-ctl resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"ftp",
				"ctl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'hits’: Cache hits; 'miss’: Cache misses; 'bytes_served’: Bytes served from cache; 'total_req’: Total requests received; 'caching_req’: Total requests to cache; 'nc_req_header’: nc_req_header; 'nc_res_header’: nc_res_header; 'rv_success’: rv_success; 'rv_failure’: rv_failure; 'ims_request’: ims_request; 'nm_response’: nm_response; 'rsp_type_CL’: rsp_type_CL; 'rsp_type_CE’: rsp_type_CE; 'rsp_type_304’: rsp_type_304; 'rsp_type_other’: rsp_type_other; 'rsp_no_compress’: rsp_no_compress; 'rsp_gzip’: rsp_gzip; 'rsp_deflate’: rsp_deflate; 'rsp_other’: rsp_other; 'nocache_match’: nocache_match; 'match’: match; 'invalidate_match’: invalidate_match; 'content_toobig’: content_toobig; 'content_toosmall’: content_toosmall; 'entry_create_failures’: entry_create_failures; 'mem_size’: mem_size; 'entry_num’: entry_num; 'replaced_entry’: replaced_entry; 'aging_entry’: aging_entry; 'cleaned_entry’: cleaned_entry;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_ftp_data",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB ftp-data resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"ftp",
				"data",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'hits’: Cache hits; 'miss’: Cache misses; 'bytes_served’: Bytes served from cache; 'total_req’: Total requests received; 'caching_req’: Total requests to cache; 'nc_req_header’: nc_req_header; 'nc_res_header’: nc_res_header; 'rv_success’: rv_success; 'rv_failure’: rv_failure; 'ims_request’: ims_request; 'nm_response’: nm_response; 'rsp_type_CL’: rsp_type_CL; 'rsp_type_CE’: rsp_type_CE; 'rsp_type_304’: rsp_type_304; 'rsp_type_other’: rsp_type_other; 'rsp_no_compress’: rsp_no_compress; 'rsp_gzip’: rsp_gzip; 'rsp_deflate’: rsp_deflate; 'rsp_other’: rsp_other; 'nocache_match’: nocache_match; 'match’: match; 'invalidate_match’: invalidate_match; 'content_toobig’: content_toobig; 'content_toosmall’: content_toosmall; 'entry_create_failures’: entry_create_failures; 'mem_size’: mem_size; 'entry_num’: entry_num; 'replaced_entry’: replaced_entry; 'aging_entry’: aging_entry; 'cleaned_entry’: cleaned_entry;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_ftp_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB ftp-proxy resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"ftp",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'hits’: Cache hits; 'miss’: Cache misses; 'bytes_served’: Bytes served from cache; 'total_req’: Total requests received; 'caching_req’: Total requests to cache; 'nc_req_header’: nc_req_header; 'nc_res_header’: nc_res_header; 'rv_success’: rv_success; 'rv_failure’: rv_failure; 'ims_request’: ims_request; 'nm_response’: nm_response; 'rsp_type_CL’: rsp_type_CL; 'rsp_type_CE’: rsp_type_CE; 'rsp_type_304’: rsp_type_304; 'rsp_type_other’: rsp_type_other; 'rsp_no_compress’: rsp_no_compress; 'rsp_gzip’: rsp_gzip; 'rsp_deflate’: rsp_deflate; 'rsp_other’: rsp_other; 'nocache_match’: nocache_match; 'match’: match; 'invalidate_match’: invalidate_match; 'content_toobig’: content_toobig; 'content_toosmall’: content_toosmall; 'entry_create_failures’: entry_create_failures; 'mem_size’: mem_size; 'entry_num’: entry_num; 'replaced_entry’: replaced_entry; 'aging_entry’: aging_entry; 'cleaned_entry’: cleaned_entry;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_generic_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB generic-proxy resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"generic",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'hits’: Cache hits; 'miss’: Cache misses; 'bytes_served’: Bytes served from cache; 'total_req’: Total requests received; 'caching_req’: Total requests to cache; 'nc_req_header’: nc_req_header; 'nc_res_header’: nc_res_header; 'rv_success’: rv_success; 'rv_failure’: rv_failure; 'ims_request’: ims_request; 'nm_response’: nm_response; 'rsp_type_CL’: rsp_type_CL; 'rsp_type_CE’: rsp_type_CE; 'rsp_type_304’: rsp_type_304; 'rsp_type_other’: rsp_type_other; 'rsp_no_compress’: rsp_no_compress; 'rsp_gzip’: rsp_gzip; 'rsp_deflate’: rsp_deflate; 'rsp_other’: rsp_other; 'nocache_match’: nocache_match; 'match’: match; 'invalidate_match’: invalidate_match; 'content_toobig’: content_toobig; 'content_toosmall’: content_toosmall; 'entry_create_failures’: entry_create_failures; 'mem_size’: mem_size; 'entry_num’: entry_num; 'replaced_entry’: replaced_entry; 'aging_entry’: aging_entry; 'cleaned_entry’: cleaned_entry;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_health_gateway",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB health-gateway resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"health",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'total_sent’: Number of Total health-check sent; 'total_retry_sent’: Number of Total health-check retry sent; 'total_timeout’: Number of Total health-check timeout;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_health_stat",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB health-stat resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"health",
				"stat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'num_burst’: Number of burst; 'max_jiffie’: Maximum number of jiffies; 'min_jiffie’: Minimum number of jiffies; 'avg_jiffie’: Average number of jiffies; 'open_socket’: Number of open sockets; 'open_socket_failed’: Number of failed open sockets; 'close_socket’: Number of closed sockets; 'connect_failed’: Number of failed connections; 'send_packet’: Number of packets sent; 'send_packet_failed’: Number of packet send failures; 'recv_packet’: Number of received packets; 'recv_packet_failed’: Number of failed packet receives; 'retry_times’: Retry times; 'timeout’: Timouet value; 'unexpected_error’: Number of unexpected errors; 'conn_imdt_succ’: Number of connection immediete success; 'sock_close_before_17’: Number of sockets closed before l7; 'sock_close_without_notify’: Number of sockets closed without notify; 'curr_health_rate’: Current health rate; 'ext_health_rate’: External health rate; 'ext_health_rate_val’: External health rate value; 'total_number’: Total number; 'status_up’: Number of status ups; 'status_down’: Number of status downs; 'status_unkn’: Number of status unknowns; 'status_other’: Number of other status; 'running_time’: Running time; 'config_health_rate’: Config health rate;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_http2",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB http2 resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"http2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'curr_proxy’: Curr Proxy Conns; 'total_proxy’: Total Proxy Conns; 'connection_preface_rcvd’: Connection preface rcvd; 'control_frame’: Control Frame Rcvd; 'headers_frame’: HEADERS Frame Rcvd; 'continuation_frame’: CONTINUATION Frame Rcvd; 'rst_frame_rcvd’: RST_STREAM Frame Rcvd; 'settings_frame’: SETTINGS Frame Rcvd; 'window_update_frame’: WINDOW_UPDATE Frame Rcvd; 'ping_frame’: PING Frame Rcvd; 'goaway_frame’: GOAWAY Frame Rcvd; 'priority_frame’: PRIORITY Frame Rcvd; 'data_frame’: DATA Frame Recvd; 'unknown_frame’: Unknown Frame Recvd; 'connection_preface_sent’: Connection preface sent; 'settings_frame_sent’: SETTINGS Frame Sent; 'settings_ack_sent’: SETTINGS ACK Frame Sent; 'empty_settings_sent’: Empty SETTINGS Frame Sent; 'ping_frame_sent’: PING Frame Sent; 'window_update_frame_sent’: WINDOW_UPDATE Frame Sent; 'rst_frame_sent’: RST_STREAM Frame Sent; 'goaway_frame_sent’: GOAWAY Frame Sent; 'header_to_app’: HEADER Frame to HTTP; 'data_to_app’: DATA Frame to HTTP; 'protocol_error’: Protocol Error; 'internal_error’: Internal Error; 'proxy_alloc_error’: HTTP2 Proxy alloc Error; 'split_buff_fail’: Splitting Buffer Failed; 'invalid_frame_size’: Invalid Frame Size Rcvd; 'error_max_invalid_stream’: Max Invalid Stream Rcvd; 'data_no_stream’: DATA Frame Rcvd on non-existent stream; 'flow_control_error’: Flow Control Error; 'settings_timeout’: Settings Timeout; 'frame_size_error’: Frame Size Error; 'refused_stream’: Refused Stream; 'cancel’: cancel; 'compression_error’: compression error; 'connect_error’: connect error; 'enhance_your_calm’: enhance your calm error; 'inadequate_security’: inadequate security; 'http_1_1_required’: HTTP1.1 Required; 'deflate_alloc_fail’: deflate alloc fail; 'inflate_alloc_fail’: inflate alloc fail; 'inflate_header_fail’: Inflate Header Fail; 'bad_connection_preface’: Bad Connection Preface; 'cant_allocate_control_frame’: Cant allocate control frame; 'cant_allocate_settings_frame’: Cant allocate SETTINGS frame; 'bad_frame_type_for_stream_state’: Bad frame type for stream state; 'wrong_stream_state’: Wrong Stream State; 'data_queue_alloc_error’: Data Queue Alloc Error; 'buff_alloc_error’: Buff alloc error; 'cant_allocate_rst_frame’: Cant allocate RST_STREAM frame; 'cant_allocate_goaway_frame’: Cant allocate GOAWAY frame; 'cant_allocate_ping_frame’: Cant allocate PING frame; 'cant_allocate_stream’: Cant allocate stream; 'cant_allocate_window_frame’: Cant allocate WINDOW_UPDATE frame; 'header_no_stream’: header no stream; 'header_padlen_gt_frame_payload’: Header padlen greater than frame payload size; 'streams_gt_max_concur_streams’: Streams greater than max allowed concurrent streams; 'idle_state_unexpected_frame’: Unxpected frame received in idle state; 'reserved_local_state_unexpected_frame’: Unexpected frame received in reserved local state; 'reserved_remote_state_unexpected_frame’: Unexpected frame received in reserved remote state; 'half_closed_remote_state_unexpected_frame’: Unexpected frame received in half closed remote state; 'closed_state_unexpected_frame’: Unexpected frame received in closed state; 'zero_window_size_on_stream’: Window Update with zero increment rcvd; 'exceeds_max_window_size_stream’: Window Update with increment that results in exceeding max window; 'stream_closed’: stream closed; 'continuation_before_headers’: CONTINUATION frame with no headers frame; 'invalid_frame_during_headers’: frame before headers were complete; 'headers_after_continuation’: headers frame before CONTINUATION was complete; 'push_promise_frame_sent’: Push Promise Frame Sent; 'invalid_push_promise’: unexpected PUSH_PROMISE frame; 'invalid_stream_id’: received invalid stream ID; 'headers_interleaved’: headers interleaved on streams; 'trailers_no_end_stream’: trailers not marked as end-of-stream; 'invalid_setting_value’: invalid setting-frame value; 'invalid_window_update’: window-update value out of range; 'frame_header_bytes_received’: frame header bytes received; 'frame_header_bytes_sent’: frame header bytes sent; 'control_bytes_received’: HTTP/2 control frame bytes received; 'control_bytes_sent’: HTTP/2 control frame bytes sent; 'header_bytes_received’: HTTP/2 header bytes received; 'header_bytes_sent’: HTTP/2 header bytes sent; 'data_bytes_received’: HTTP/2 data bytes received; 'data_bytes_sent’: HTTP/2 data bytes sent; 'total_bytes_received’: HTTP/2 total bytes received; 'total_bytes_sent’: HTTP/2 total bytes sent; 'peak_proxy’: Peak Proxy Conns; 'control_frame_sent’: Control Frame Sent; 'continuation_frame_sent’: CONTINUATION Frame Sent; 'data_frame_sent’: DATA Frame Sent; 'headers_frame_sent’: HEADERS Frame Sent; 'priority_frame_sent’: PRIORITY Frame Sent; 'settings_ack_rcvd’: SETTINGS ACK Frame Rcvd; 'empty_settings_rcvd’: Empty SETTINGS Frame Rcvd; 'alloc_fail_total’: Alloc Fail - Total; 'err_rcvd_total’: Error Rcvd - Total; 'err_sent_total’: Error Rent - Total; 'err_sent_proto_err’: Error Sent - PROTOCOL_ERROR; 'err_sent_internal_err’: Error Sent - INTERNAL_ERROR; 'err_sent_flow_control’: Error Sent - FLOW_CONTROL_ERROR; 'err_sent_setting_timeout’: Error Sent - SETTINGS_TIMEOUT; 'err_sent_stream_closed’: Error Sent - STREAM_CLOSED; 'err_sent_frame_size_err’: Error Sent - FRAME_SIZE_ERROR; 'err_sent_refused_stream’: Error Sent - REFUSED_STREAM; 'err_sent_cancel’: Error Sent - CANCEL; 'err_sent_compression_err’: Error Sent - COMPRESSION_ERROR; 'err_sent_connect_err’: Error Sent - CONNECT_ERROR; 'err_sent_your_calm’: Error Sent - ENHANCE_YOUR_CALM; 'err_sent_inadequate_security’: Error Sent - INADEQUATE_SECURITY; 'err_sent_http11_required’: Error Sent - HTTP_1_1_REQUIRED;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_http_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB http-proxy resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"http",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'num’: Num; 'curr_proxy’: Curr Proxy Conns; 'total_proxy’: Total Proxy Conns; 'req’: HTTP requests; 'req_succ’: HTTP requests(succ); 'noproxy’: No proxy error; 'client_rst’: Client RST; 'server_rst’: Server RST; 'notuple’: No tuple error; 'parsereq_fail’: Parse req fail; 'svrsel_fail’: Server selection fail; 'fwdreq_fail’: Fwd req fail; 'fwdreq_fail_buff’: Fwd req fail - buff; 'fwdreq_fail_rport’: Fwd req fail - rport; 'fwdreq_fail_route’: Fwd req fail - route; 'fwdreq_fail_persist’: Fwd req fail - persist; 'fwdreq_fail_server’: Fwd req fail - server; 'fwdreq_fail_tuple’: Fwd req fail - tuple; 'fwdreqdata_fail’: fwdreqdata_fail; 'req_retran’: Packets retrans; 'req_ofo’: Packets ofo; 'server_resel’: Server reselection; 'svr_prem_close’: Server premature close; 'new_svrconn’: Server conn made; 'snat_fail’: Source NAT failure; 'tcpoutrst’: Out RSTs; 'full_proxy’: Full proxy tot; 'full_proxy_post’: Full proxy POST; 'full_proxy_pipeline’: Full proxy pipeline; 'full_proxy_fpga_err’: Full proxy fpga err; 'req_over_limit’: Request over limit; 'req_rate_over_limit’: Request rate over limit; 'l4_switching’: L4 switching; 'cookie_switching’: Cookie switching; 'aflex_switching’: aFleX switching; 'http_policy_switching’: HTTP Policy switching; 'url_switching’: URL switching; 'host_switching’: Host switching; 'lb_switching’: Normal LB switching; 'l4_switching_ok’: L4 switching (succ); 'cookie_switching_ok’: Cookie switching (succ); 'aflex_switching_ok’: aFleX switching (succ); 'http_policy_switching_ok’: HTTP Policy switching (succ); 'url_switching_ok’: URL switching (succ); 'host_switching_ok’: Host switching (succ); 'lb_switching_ok’: Normal LB switch. (succ); 'l4_switching_enqueue’: L4 switching (enQ); 'cookie_switching_enqueue’: Cookie switching (enQ); 'aflex_switching_enqueue’: aFleX switching (enQ); 'http_policy_switching_enqueue’: HTTP Policy switching (enQ); 'url_switching_enqueue’: URL switching (enQ); 'host_switching_enqueue’: Host switching (enQ); 'lb_switching_enqueue’: Normal LB switch. (enQ); 'retry_503’: Retry on 503; 'aflex_retry’: aFleX http retry; 'aflex_lb_reselect’: aFleX lb reselect; 'aflex_lb_reselect_ok’: aFleX lb reselect (succ); 'client_rst_request’: Client RST - request; 'client_rst_connecting’: Client RST - connecting; 'client_rst_connected’: Client RST - connected; 'client_rst_response’: Client RST - response; 'server_rst_request’: Server RST - request; 'server_rst_connecting’: Server RST - connecting; 'server_rst_connected’: Server RST - connected; 'server_rst_response’: Server RST - response; 'invalid_header’: Invalid header; 'too_many_headers’: Too many headers; 'line_too_long’: Line too long; 'header_name_too_long’: Header name too long; 'wrong_resp_header’: Wrong response header; 'header_insert’: Header insert; 'header_delete’: Header delete; 'insert_client_ip’: Insert client IP; 'negative_req_remain’: Negative request remain; 'negative_resp_remain’: Negative response remain; 'large_cookie’: Large cookies; 'large_cookie_header’: Large cookie headers; 'huge_cookie’: Huge cookies; 'huge_cookie_header’: Huge cookie headers; 'parse_cookie_fail’: Parse cookie fail; 'parse_setcookie_fail’: Parse set-cookie fail; 'asm_cookie_fail’: Assemble cookie fail; 'asm_cookie_header_fail’: Asm cookie header fail; 'asm_setcookie_fail’: Assemble set-cookie fail; 'asm_setcookie_header_fail’: Asm set-cookie hdr fail; 'client_req_unexp_flag’: Client req unexp flags; 'connecting_fin’: Connecting FIN; 'connecting_fin_retrans’: Connecting FIN retran; 'connecting_fin_ofo’: Connecting FIN ofo; 'connecting_rst’: Connecting RST; 'connecting_rst_retrans’: Connecting RST retran; 'connecting_rst_ofo’: Connecting RST ofo; 'connecting_ack’: Connecting ACK; 'pkts_ofo’: Packets ofo; 'pkts_retrans’: Packets retrans; 'pkts_retrans_ack_finwait’: retrans ACK FWAIT; 'pkts_retrans_fin’: retrans FIN; 'pkts_retrans_rst’: retrans RST; 'pkts_retrans_push’: retrans PSH; 'stale_sess’: Stale sess; 'server_resel_failed’: Server re-select failed; 'compression_before’: Tot data before compress; 'compression_after’: Tot data after compress; 'response_1xx’: Status code 1XX; 'response_100’: Status code 100; 'response_101’: Status code 101; 'response_102’: Status code 102; 'response_2xx’: Status code 2XX; 'response_200’: Status code 200; 'response_201’: Status code 201; 'response_202’: Status code 202; 'response_203’: Status code 203; 'response_204’: Status code 204; 'response_205’: Status code 205; 'response_206’: Status code 206; 'response_207’: Status code 207; 'response_3xx’: Status code 3XX; 'response_300’: Status code 300; 'response_301’: Status code 301; 'response_302’: Status code 302; 'response_303’: Status code 303; 'response_304’: Status code 304; 'response_305’: Status code 305; 'response_306’: Status code 306; 'response_307’: Status code 307; 'response_4xx’: Status code 4XX; 'response_400’: Status code 400; 'response_401’: Status code 401; 'response_402’: Status code 402; 'response_403’: Status code 403; 'response_404’: Status code 404; 'response_405’: Status code 405; 'response_406’: Status code 406; 'response_407’: Status code 407; 'response_408’: Status code 408; 'response_409’: Status code 409; 'response_410’: Status code 410; 'response_411’: Status code 411; 'response_412’: Status code 412; 'response_413’: Status code 413; 'response_414’: Status code 414; 'response_415’: Status code 415; 'response_416’: Status code 416; 'response_417’: Status code 417; 'response_418’: Status code 418; 'response_422’: Status code 422; 'response_423’: Status code 423; 'response_424’: Status code 424; 'response_425’: Status code 425; 'response_426’: Status code 426; 'response_449’: Status code 449; 'response_450’: Status code 450; 'response_5xx’: Status code 5XX; 'response_500’: Status code 500; 'response_501’: Status code 501; 'response_502’: Status code 502; 'response_503’: Status code 503; 'response_504’: Status code 504; 'response_505’: Status code 505; 'response_506’: Status code 506; 'response_507’: Status code 507; 'response_508’: Status code 508; 'response_509’: Status code 509; 'response_510’: Status code 510; 'response_6xx’: Status code 6XX; 'response_unknown’: Status code unknown; 'req_http10’: Request 1.0; 'req_http11’: Request 1.1; 'response_http10’: Resp 1.0; 'response_http11’: Resp 1.1; 'req_get’: Method GET; 'req_head’: Method HEAD; 'req_put’: Method PUT; 'req_post’: Method POST; 'req_trace’: Method TRACE; 'req_options’: Method OPTIONS; 'req_connect’: Method CONNECT; 'req_delete’: Method DELETE; 'req_unknown’: Method UNKNOWN; 'req_content_len’: Req content len; 'rsp_content_len’: Resp content len; 'rsp_chunk’: Resp chunk encoding; 'req_chunk’: Req chunk encoding; 'compress_rsp’: Compress req; 'compress_del_accept_enc’: Compress del accept enc; 'compress_resp_already_compressed’: Resp already compressed; 'compress_content_type_excluded’: Compress cont type excl; 'compress_no_content_type’: Compress no cont type; 'compress_resp_lt_min’: Compress resp less than min; 'compress_resp_no_cl_or_ce’: Compress resp no CL/CE; 'compress_ratio_too_high’: Compress ratio too high; 'cache_rsp’: HTTP req (cache succ); 'close_on_ddos’: Close on DDoS; 'req_http10_keepalive’: 1.0 Keepalive; 'req_sz_1k’: Req less than equal to 1K; 'req_sz_2k’: Req less than equal to 2K; 'req_sz_4k’: Req less than equal to 4K;`,
				},
				resource.Attribute{
					Name:        "counters2",
					Description: `'req_sz_8k’: Req less than equal to 8K; 'req_sz_16k’: Req less than equal to 16K; 'req_sz_32k’: Req less than equal to 32K; 'req_sz_64k’: Req less than equal to 64K; 'req_sz_256k’: Req less than equal to 256K; 'req_sz_gt_256k’: Req greater than 256K; 'rsp_sz_1k’: Resp less than equal to 1K; 'rsp_sz_2k’: Resp less than equal to 2K; 'rsp_sz_4k’: Resp less than equal to 4K; 'rsp_sz_8k’: Resp less than equal to 8K; 'rsp_sz_16k’: Resp less than equal to 16K; 'rsp_sz_32k’: Resp less than equal to 32K; 'rsp_sz_64k’: Resp less than equal to 64K; 'rsp_sz_256k’: Resp less than equal to 256K; 'rsp_sz_gt_256k’: Resp greater than 256K; 'chunk_sz_512’: Chunk less than equal to 512; 'chunk_sz_1k’: Chunk less than equal to 1K; 'chunk_sz_2k’: Chunk less than equal to 2K; 'chunk_sz_4k’: Chunk less than equal to 4K; 'chunk_sz_gt_4k’: Chunk greater than 4K; 'pconn_connecting’: pconn connecting; 'pconn_connected’: pconn connected; 'pconn_connecting_failed’: pconn conn failed; 'chunk_bad’: Bad Chunk; 'req_10u’: Rsp time less than 10u; 'req_20u’: Rsp time less than 20u; 'req_50u’: Rsp time less than 50u; 'req_100u’: Rsp time less than 100u; 'req_200u’: Rsp time less than 200u; 'req_500u’: Rsp time less than 500u; 'req_1m’: Rsp time less than 1m; 'req_2m’: Rsp time less than 2m; 'req_5m’: Rsp time less than 5m; 'req_10m’: Rsp time less than 10m; 'req_20m’: Rsp time less than 20m; 'req_50m’: Rsp time less than 50m; 'req_100m’: Rsp time less than 100m; 'req_200m’: Rsp time less than 200m; 'req_500m’: Rsp time less than 500m; 'req_1s’: Rsp time less than 1s; 'req_2s’: Rsp time less than 2s; 'req_5s’: Rsp time less than 5s; 'req_over_5s’: Rsp time greater than equal to 5s; 'insert_client_port’: Insert client Port; 'req_track’: Method TRACK; 'connect_req’: Total HTTP CONNECT requests; 'req_enter_ssli’: Total HTTP requests enter SSLi; 'non_http_bypass’: Non-HTTP bypass; 'decompression_before’: Tot data before decompress; 'decompression_after’: Tot data after decompress; 'req_http2’: Request 2.0; 'response_http2’: Resp 2.0; 'req_timeout_retry’: Retry on Req Timeout; 'req_timeout_close’: Close on Req Timeout; 'doh_req’: DoH Requests; 'doh_req_get’: DoH GET Requests; 'doh_req_post’: DoH POST Requests; 'doh_non_doh_req’: DoH non DoH Requests; 'doh_non_doh_req_get’: DoH non DoH GET Requests; 'doh_non_doh_req_post’: DoH non DoH POST Requests; 'doh_resp’: DoH Responses; 'doh_tc_resp’: DoH TC Responses; 'doh_udp_dns_req’: DoH UDP DNS Requests; 'doh_udp_dns_resp’: DoH UDP DNS Responses; 'doh_tcp_dns_req’: DoH TCP DNS Requests; 'doh_tcp_dns_resp’: DoH TCP DNS Responses; 'doh_req_send_failed’: DoH Request Send Failed; 'doh_resp_send_failed’: DoH Response Send Failed; 'doh_malloc_fail’: DoH Memory alloc failed; 'doh_req_udp_retry’: DoH UDP Retry; 'doh_req_udp_retry_fail’: DoH UDP Retry failed; 'doh_req_tcp_retry’: DoH TCP Retry; 'doh_req_tcp_retry_fail’: DoH TCP Retry failed; 'doh_snat_failed’: DoH Source NAT failed; 'doh_path_not_found’: DoH URI Path not found; 'doh_get_dns_arg_failed’: DoH GET dns arg not found in uri; 'doh_get_base64_decode_failed’: DoH GET base64url decode failed; 'doh_post_content_type_mismatch’: DoH POST content-type not found; 'doh_post_payload_not_found’: DoH POST payload not found; 'doh_post_payload_extract_failed’: DoH POST payload extract failed; 'doh_non_doh_method’: DoH Non DoH HTTP request method rcvd; 'doh_tcp_send_failed’: DoH serv TCP DNS send failed; 'doh_udp_send_failed’: DoH serv UDP DNS send failed; 'doh_query_time_out’: DoH serv Query timed out; 'doh_dns_query_type_a’: DoH Query type A; 'doh_dns_query_type_aaaa’: DoH Query type AAAA; 'doh_dns_query_type_ns’: DoH Query type NS; 'doh_dns_query_type_cname’: DoH Query type CNAME; 'doh_dns_query_type_any’: DoH Query type ANY; 'doh_dns_query_type_srv’: DoH Query type SRV; 'doh_dns_query_type_mx’: DoH Query type MX; 'doh_dns_query_type_soa’: DoH Query type SOA; 'doh_dns_query_type_others’: DoH Query type Others; 'doh_resp_setup_failed’: DoH Response setup failed; 'doh_resp_header_alloc_failed’: DoH Resp hdr alloc failed; 'doh_resp_que_failed’: DoH Resp queue failed; 'doh_resp_udp_frags’: DoH UDP Frags Rcvd; 'doh_resp_tcp_frags’: DoH TCP Frags Rcvd; 'doh_serv_sel_failed’: DoH Server Select Failed; 'doh_retry_w_tcp’: DoH Retry with TCP SG;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_hw_compress",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB hw-compress resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"hw",
				"compress",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sampling_enable",
					Description: ``,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'request_count’: Total request count; 'submit_count’: Total submit count; 'response_count’: Total response count; 'failure_count’: Total failure count; 'failure_code’: Last failure code; 'ring_full_count’: Compression queue full; 'max_outstanding_request_count’: Max queued request count; 'max_outstanding_submit_count’: Max queued submit count;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_icap",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB icap resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"icap",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'reqmod_request’: Reqmod Request Stats; 'respmod_request’: Respmod Request Stats; 'reqmod_request_after_100’: Reqmod Request Sent After 100 Cont Stats; 'respmod_request_after_100’: Respmod Request Sent After 100 Cont Stats; 'reqmod_response’: Reqmod Response Stats; 'respmod_response’: Respmod Response Stats; 'reqmod_response_after_100’: Reqmod Response After 100 Cont Stats; 'respmod_response_after_100’: Respmod Response After 100 Cont Stats; 'chunk_no_allow_204’: Chunk so no Allow 204 Stats; 'len_exceed_no_allow_204’: Length Exceeded so no Allow 204 Stats; 'result_continue’: Result Continue Stats; 'result_icap_response’: Result ICAP Response Stats; 'result_100_continue’: Result 100 Continue Stats; 'result_other’: Result Other Stats; 'status_2xx’: Status 2xx Stats; 'status_200’: Status 200 Stats; 'status_201’: Status 201 Stats; 'status_202’: Status 202 Stats; 'status_203’: Status 203 Stats; 'status_204’: Status 204 Stats; 'status_205’: Status 205 Stats; 'status_206’: Status 206 Stats; 'status_207’: Status 207 Stats; 'status_1xx’: Status 1xx Stats; 'status_100’: Status 100 Stats; 'status_101’: Status 101 Stats; 'status_102’: Status 102 Stats; 'status_3xx’: Status 3xx Stats; 'status_300’: Status 300 Stats; 'status_301’: Status 301 Stats; 'status_302’: Status 302 Stats; 'status_303’: Status 303 Stats; 'status_304’: Status 304 Stats; 'status_305’: Status 305 Stats; 'status_306’: Status 306 Stats; 'status_307’: Status 307 Stats; 'status_4xx’: Status 4xx Stats; 'status_400’: Status 400 Stats; 'status_401’: Status 401 Stats; 'status_402’: Status 402 Stats; 'status_403’: Status 403 Stats; 'status_404’: Status 404 Stats; 'status_405’: Status 405 Stats; 'status_406’: Status 406 Stats; 'status_407’: Status 407 Stats; 'status_408’: Status 408 Stats; 'status_409’: Status 409 Stats; 'status_410’: Status 410 Stats; 'status_411’: Status 411 Stats; 'status_412’: Status 412 Stats; 'status_413’: Status 413 Stats; 'status_414’: Status 414 Stats; 'status_415’: Status 415 Stats; 'status_416’: Status 416 Stats; 'status_417’: Status 417 Stats; 'status_418’: Status 418 Stats; 'status_419’: Status 419 Stats; 'status_420’: Status 420 Stats; 'status_422’: Status 422 Stats; 'status_423’: Status 423 Stats; 'status_424’: Status 424 Stats; 'status_425’: Status 425 Stats; 'status_426’: Status 426 Stats; 'status_449’: Status 449 Stats; 'status_450’: Status 450 Stats; 'status_5xx’: Status 5xx Stats; 'status_500’: Status 500 Stats; 'status_501’: Status 501 Stats; 'status_502’: Status 502 Stats; 'status_503’: Status 503 Stats; 'status_504’: Status 504 Stats; 'status_505’: Status 505 Stats; 'status_506’: Status 506 Stats; 'status_507’: Status 507 Stats; 'status_508’: Status 508 Stats; 'status_509’: Status 509 Stats; 'status_510’: Status 510 Stats; 'status_6xx’: Status 6xx Stats; 'status_unknown’: Status Unknown Stats; 'send_option_req’: Send Option Req Stats; 'app_serv_conn_no_pcb_err’: App Server Conn no ES PCB Err Stats; 'app_serv_conn_err’: App Server Conn Err Stats; 'chunk1_hdr_err’: Chunk Hdr Err1 Stats; 'chunk2_hdr_err’: Chunk Hdr Err2 Stats; 'chunk_bad_trail_err’: Chunk Bad Trail Err Stats; 'no_payload_next_buff_err’: No Payload In Next Buff Err Stats; 'no_payload_buff_err’: No Payload Buff Err Stats; 'resp_hdr_incomplete_err’: Resp Hdr Incomplete Err Stats; 'serv_sel_fail_err’: Server Select Fail Err Stats; 'start_icap_conn_fail_err’: Start ICAP conn fail Stats; 'prep_req_fail_err’: Prepare ICAP req fail Err Stats; 'icap_ver_err’: ICAP Ver Err Stats; 'icap_line_err’: ICAP Line Err Stats; 'encap_hdr_incomplete_err’: Encap HDR Incomplete Err Stats; 'no_icap_resp_err’: No ICAP Resp Err Stats; 'resp_line_read_err’: Resp Line Read Err Stats; 'resp_line_parse_err’: Resp Line Parse Err Stats; 'resp_hdr_err’: Resp Hdr Err Stats; 'req_hdr_incomplete_err’: Req Hdr Incomplete Err Stats; 'no_status_code_err’: No Status Code Err Stats; 'http_resp_line_read_err’: HTTP Response Line Read Err Stats; 'http_resp_line_parse_err’: HTTP Response Line Parse Err Stats; 'http_resp_hdr_err’: HTTP Resp Hdr Err Stats; 'recv_option_resp’: Send Option Req Stats;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_icap_http",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB icap http resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"icap",
				"http",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'status_200’: Status code 200; 'status_201’: Status code 201; 'status_202’: Status code 202; 'status_203’: Status code 203; 'status_204’: Status code 204; 'status_205’: Status code 205; 'status_206’: Status code 206; 'status_207’: Status code 207; 'status_100’: Status code 100; 'status_101’: Status code 101; 'status_102’: Status code 102; 'status_300’: Status code 300; 'status_301’: Status code 301; 'status_302’: Status code 302; 'status_303’: Status code 303; 'status_304’: Status code 304; 'status_305’: Status code 305; 'status_306’: Status code 306; 'status_307’: Status code 307; 'status_400’: Status code 400; 'status_401’: Status code 401; 'status_402’: Status code 402; 'status_403’: Status code 403; 'status_404’: Status code 404; 'status_405’: Status code 405; 'status_406’: Status code 406; 'status_407’: Status code 407; 'status_408’: Status code 408; 'status_409’: Status code 409; 'status_410’: Status code 410; 'status_411’: Status code 411; 'status_412’: Status code 412; 'status_413’: Status code 413; 'status_414’: Status code 414; 'status_415’: Status code 415; 'status_416’: Status code 416; 'status_417’: Status code 417; 'status_418’: Status code 418; 'status_422’: Status code 422; 'status_423’: Status code 423; 'status_424’: Status code 424; 'status_425’: Status code 425; 'status_426’: Status code 426; 'status_449’: Status code 449; 'status_450’: Status code 450; 'status_500’: Status code 500; 'status_501’: Status code 501; 'status_502’: Status code 502; 'status_503’: Status code 503; 'status_504’: Status code 504; 'status_505’: Status code 505; 'status_506’: Status code 506; 'status_507’: Status code 507; 'status_508’: Status code 508; 'status_509’: Status code 509; 'status_510’: Status code 510; 'status_1xx’: status code 1XX; 'status_2xx’: status code 2XX; 'status_3xx’: status code 3XX; 'status_4xx’: status code 4XX; 'status_5xx’: status code 5XX; 'status_6xx’: status code 6XX; 'status_unknown’: Status code unknown;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_imapproxy",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB imapproxy resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"imapproxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'num’: Num; 'curr’: Current proxy conns; 'total’: Total proxy conns; 'svrsel_fail’: Server selection failure; 'no_route’: no route failure; 'snat_fail’: source nat failure; 'feat’: feat packet; 'cc’: clear ctrl port packet; 'data_ssl’: data ssl force; 'line_too_long’: line too long; 'line_mem_freed’: request line freed; 'invalid_start_line’: invalid start line; 'auth_tls’: auth tls cmd; 'prot’: prot cmd; 'pbsz’: pbsz cmd; 'pasv’: pasv cmd; 'port’: port cmd; 'request_dont_care’: other cmd; 'client_auth_tls’: client auth tls; 'cant_find_pasv’: cant find pasv; 'pasv_addr_ne_server’: psv addr not equal to svr; 'smp_create_fail’: smp create fail; 'data_server_conn_fail’: data svr conn fail; 'data_send_fail’: data send fail; 'epsv’: epsv command; 'cant_find_epsv’: cant find epsv; 'data_curr’: Current Data Proxy; 'data_total’: Total Data Proxy; 'auth_unsupported’: Unsupported auth; 'adat’: adat cmd; 'unsupported_pbsz_value’: Unsupported PBSZ; 'unsupported_prot_value’: Unsupported PROT; 'unsupported_command’: Unsupported cmd; 'control_to_clear’: Control chn clear txt; 'control_to_ssl’: Control chn ssl; 'bad_sequence’: Bad Sequence; 'rsv_persist_conn_fail’: Serv Sel Persist fail; 'smp_v6_fail’: Serv Sel SMPv6 fail; 'smp_v4_fail’: Serv Sel SMPv4 fail; 'insert_tuple_fail’: Serv Sel insert tuple fail; 'cl_est_err’: Client EST state erro; 'ser_connecting_err’: Serv CTNG state error; 'server_response_err’: Serv RESP state error; 'cl_request_err’: Client RQ state error; 'data_conn_start_err’: Data Start state error; 'data_serv_connecting_err’: Data Serv CTNG error; 'data_serv_connected_err’: Data Serv CTED error; 'request’: Total FTP Request; 'capability’: Capability cmd; 'start_tls’: Total Start TLS cmd; 'login’: Total Login cmd;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_l4",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB l4 resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"l4",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'intcp’: TCP received; 'synreceived’: TCP SYN received; 'tcp_fwd_last_ack’: L4 rcv fwd last ACK; 'tcp_rev_last_ack’: L4 rcv rev last ACK; 'tcp_rev_fin’: L4 rcv rev FIN; 'tcp_fwd_fin’: L4 rcv fwd FIN; 'tcp_fwd_ackfin’: L4 rcv fwd FIN|ACK; 'inudp’: UDP received; 'syncookiessent’: TCP SYN cookie snt; 'syncookiessent_ts’: TCP SYN cookie snt ts; 'syncookiessentfailed’: TCP SYN cookie snt fail; 'outrst’: TCP out RST; 'outrst_nosyn’: TCP out RST no SYN; 'outrst_broker’: TCP out RST L4 proxy; 'outrst_ack_attack’: TCP out RST ACK attack; 'outrst_aflex’: TCP out RST aFleX; 'outrst_stale_sess’: TCP out RST stale sess; 'syn_stale_sess’: SYN stale sess drop; 'outrst_tcpproxy’: TCP out RST TCP proxy; 'svrselfail’: Server sel failure; 'noroute’: IP out noroute; 'snat_fail’: Source NAT failure; 'snat_no_fwd_route’: Source NAT no fwd route; 'snat_no_rev_route’: Source NAT no rev route; 'snat_icmp_error_process’: Source NAT ICMP Process; 'snat_icmp_no_match’: Source NAT ICMP No Match; 'smart_nat_id_mismatch’: Auto NAT id mismatch; 'syncookiescheckfailed’: TCP SYN cookie failed; 'novport_drop’: NAT no session drops; 'no_vport_drop’: vport not matching drops; 'nosyn_drop’: No SYN pkt drops; 'nosyn_drop_fin’: No SYN pkt drops - FIN; 'nosyn_drop_rst’: No SYN pkt drops - RST; 'nosyn_drop_ack’: No SYN pkt drops - ACK; 'connlimit_drop’: Conn Limit drops; 'connlimit_reset’: Conn Limit resets; 'conn_rate_limit_drop’: Conn rate limit drops; 'conn_rate_limit_reset’: Conn rate limit resets; 'proxy_nosock_drop’: Proxy no sock drops; 'drop_aflex’: aFleX drops; 'sess_aged_out’: Session aged out; 'tcp_sess_aged_out’: TCP Session aged out; 'udp_sess_aged_out’: UDP Session aged out; 'other_sess_aged_out’: Other Session aged out; 'tcp_no_slb’: TCP no SLB; 'udp_no_slb’: UDP no SLB; 'throttle_syn’: SYN Throttle; 'drop_gslb’: Drop GSLB; 'inband_hm_retry’: Inband HM retry; 'inband_hm_reassign’: Inband HM reassign; 'auto_reassign’: Auto-reselect server; 'fast_aging_set’: Fast aging set; 'fast_aging_reset’: Fast aging reset; 'dns_policy_drop’: DNS Policy Drop; 'tcp_invalid_drop’: TCP invalid drop; 'anomaly_out_seq’: Anomaly out of sequence; 'anomaly_zero_win’: Anomaly zero window; 'anomaly_bad_content’: Anomaly bad content; 'anomaly_pbslb_drop’: Anomaly pbslb drop; 'no_resourse_drop’: No resource drop; 'reset_unknown_conn’: Reset unknown conn; 'reset_l7_on_failover’: RST L7 on failover; 'ignore_msl’: ignore msl; 'l2_dsr’: L2 DSR received; 'l3_dsr’: L3 DSR received; 'port_preserve_attempt’: NAT Port Preserve Try; 'port_preserve_succ’: NAT Port Preserve Succ; 'tcpsyndata_drop’: TCP SYN With Data Drop; 'tcpotherflags_drop’: TCP SYN Other Flags Drop; 'bw_rate_limit_exceed’: BW-Limit Exceed drop; 'bw_watermark_drop’: BW-Watermark drop; 'l4_cps_exceed’: L4 CPS exceed drop; 'nat_cps_exceed’: NAT CPS exceed drop; 'l7_cps_exceed’: L7 CPS exceed drop; 'ssl_cps_exceed’: SSL CPS exceed drop; 'ssl_tpt_exceed’: SSL TPT exceed drop; 'ssl_watermark_drop’: SSL TPT-Watermark drop; 'concurrent_conn_exceed’: L3V Conn Limit Drop; 'svr_syn_handshake_fail’: L4 server handshake fail; 'stateless_conn_timeout’: L4 stateless Conn TO; 'tcp_ax_rexmit_syn’: L4 AX re-xmit SYN; 'tcp_syn_rcv_ack’: L4 rcv ACK on SYN; 'tcp_syn_rcv_rst’: L4 rcv RST on SYN; 'tcp_sess_noest_aged_out’: TCP no-Est Sess aged out; 'tcp_sess_noest_csyn_rcv_aged_out’: no-Est CSYN rcv aged out; 'tcp_sess_noest_ssyn_xmit_aged_out’: no-Est SSYN snt aged out; 'tcp_rexmit_syn’: L4 rcv rexmit SYN; 'tcp_rexmit_syn_delq’: L4 rcv rexmit SYN (delq); 'tcp_rexmit_synack’: L4 rcv rexmit SYN|ACK; 'tcp_rexmit_synack_delq’: L4 rcv rexmit SYN|ACK DQ; 'tcp_fwd_fin_dup’: L4 rcv fwd FIN dup; 'tcp_rev_fin_dup’: L4 rcv rev FIN dup; 'tcp_rev_ackfin’: L4 rcv rev FIN|ACK; 'tcp_fwd_rst’: L4 rcv fwd RST; 'tcp_rev_rst’: L4 rcv rev RST; 'udp_req_oneplus_no_resp’: L4 UDP reqs no rsp; 'udp_req_one_oneplus_resp’: L4 UDP req rsps; 'udp_req_resp_notmatch’: L4 UDP req/rsp not match; 'udp_req_more_resp’: L4 UDP req greater than rsps; 'udp_resp_more_req’: L4 UDP rsps greater than reqs; 'udp_req_oneplus’: L4 UDP reqs; 'udp_resp_oneplus’: L4 UDP rsps; 'out_seq_ack_drop’: Out of sequence ACK drop; 'tcp_est’: L4 TCP Established; 'synattack’: L4 SYN attack; 'syn_rate’: TCP SYN rate per sec; 'syncookie_buff_drop’: TCP SYN cookie buff drop; 'syncookie_buff_queue’: TCP SYN cookie buff queue; 'skip_insert_client_ip’: Skip Insert-client-ip; 'synreceived_hw’: TCP SYN (HW SYN cookie); 'dns_id_switch’: DNS query id switch; 'server_down_del’: Server Down Del switch; 'dnssec_switch’: DNSSEC SG switch; 'rate_drop_reset_unkn’: Rate Drop reset; 'tcp_connections_closed’: TCP Connections Closed;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_l7session",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB l7session resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"l7session",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'start_server_conn_succ’: Start Server Conn Success; 'conn_not_exist’: Conn does not exist; 'data_event’: Data event from TCP; 'client_fin’: FIN from client; 'server_fin’: FIN from server; 'wbuf_event’: Wbuf event from TCP; 'wbuf_cb_failed’: Wbuf event callback failed; 'err_event’: Err event from TCP; 'err_cb_failed’: Err event callback failed; 'server_conn_failed’: Server connection failed; 'client_rst’: RST from client; 'server_rst’: RST from server; 'client_rst_req’: RST from client - request; 'client_rst_connecting’: RST from client - connecting; 'client_rst_connected’: RST from client - connected; 'client_rst_rsp’: RST from client - response; 'server_rst_req’: RST from server - request; 'server_rst_connecting’: RST from server - connecting; 'server_rst_connected’: RST from server - connected; 'server_rst_rsp’: RST from server - response; 'proxy_v1_connection’: counter for Proxy v1 connection; 'proxy_v2_connection’: counter for Proxy v2 connection; 'curr_proxy’: Curr proxy conn; 'curr_proxy_client’: Curr proxy conn - client; 'curr_proxy_server’: Curr proxy conn - server; 'curr_proxy_es’: Curr proxy conn - ES; 'total_proxy’: Total proxy conn; 'total_proxy_client’: Total proxy conn - client; 'total_proxy_server’: Total proxy conn - server; 'total_proxy_es’: Total proxy conn - ES; 'server_select_fail’: Server selection fail; 'est_event’: Est event from TCP; 'est_cb_failed’: Est event callback fail; 'data_cb_failed’: Data event callback fail; 'hps_fwdreq_fail’: Fwd req fail; 'hps_fwdreq_fail_buff’: Fwd req fail - buff; 'hps_fwdreq_fail_rport’: Fwd req fail - rport; 'hps_fwdreq_fail_route’: Fwd req fail - route; 'hps_fwdreq_fail_persist’: Fwd req fail - persist; 'hps_fwdreq_fail_server’: Fwd req fail - server; 'hps_fwdreq_fail_tuple’: Fwd req fail - tuple;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_mlb",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB mlb resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"mlb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'client_msg_sent’: Client message sent; 'server_msg_received’: Server message received; 'server_conn_created’: Server connection created; 'server_conn_rst’: Server connection reset; 'server_conn_failed’: Server connection failed; 'server_conn_closed’: Server connection closed; 'client_conn_created’: Client connection created; 'client_conn_closed’: Client connection closed; 'client_conn_not_found’: Client connection not found;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_mssql",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB mssql resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"mssql",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'curr_proxy’: Curr Proxy Conns; 'total_proxy’: Total Proxy Conns; 'curr_be_enc’: Curr BE Encryption Conns; 'total_be_enc’: Total BE Encryption Conns; 'curr_fe_enc’: Curr FE Encryption Conns; 'total_fe_enc’: Total FE Encryption Conns; 'client_fin’: Client FIN; 'server_fin’: Server FIN; 'session_err’: Session err; 'queries’: DB Queries; 'commands’: DB commands reply; 'auth_success’: Authentication Success; 'auth_failure’: Authentication Failure;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_mysql",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB mysql resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"mysql",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'curr_proxy’: Curr Proxy Conns; 'total_proxy’: Total Proxy Conns; 'curr_be_enc’: Curr BE Encryption Conns; 'total_be_enc’: Total BE Encryption Conns; 'curr_fe_enc’: Curr FE Encryption Conns; 'total_fe_enc’: Total FE Encryption Conns; 'client_fin’: Client FIN; 'server_fin’: Server FIN; 'session_err’: Session err; 'queries’: DB Queries; 'commands’: DB commands reply;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_passthrough",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB passthrough resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"passthrough",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'curr_conn’: Current connections; 'total_conn’: Total connections; 'total_fwd_bytes’: Forward bytes; 'total_fwd_packets’: Forward packets; 'total_rev_bytes’: Reverse bytes; 'total_rev_packets’: Reverse packets; 'curr_pconn’: Persistent connections;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_perf",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB perf resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"perf",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'total-throughput-bits-per-sec’: Total Throughput in bits/sec; 'l4-conns-per-sec’: L4 Connections/sec; 'l7-conns-per-sec’: L7 Connections/sec; 'l7-trans-per-sec’: L7 Transactions/sec; 'ssl-conns-per-sec’: SSL Connections/sec; 'ip-nat-conns-per-sec’: IP NAT Connections/sec; 'total-new-conns-per-sec’: Total New Connections Established/sec; 'total-curr-conns’: Total Current Established Connections; 'l4-bandwidth’: L4 Bandwidth in bits/sec; 'l7-bandwidth’: L7 Bandwidth in bits/sec; 'serv-ssl-conns-per-sec’: Server SSL Connections/sec; 'fw-conns-per-sec’: FW Connections/sec; 'gifw-conns-per-sec’: GiFW Connections/sec; 'l7-proxy-conns-per-sec’: L7 Proxy Connections/sec; 'l7-proxy-trans-per-sec’: L7 Proxy Transactions/sec;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_persist",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB persist resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"persist",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'hash_tbl_trylock_fail’: Hash tbl lock fail; 'hash_tbl_create_ok’: Hash tbl create ok; 'hash_tbl_create_fail’: Hash tbl create fail; 'hash_tbl_free’: Hash tbl free; 'hash_tbl_rst_updown’: Hash tbl reset (up/down); 'hash_tbl_rst_adddel’: Hash tbl reset (add/del); 'url_hash_pri’: URL hash persist (pri); 'url_hash_enqueue’: URL hash persist (enQ); 'url_hash_sec’: URL hash persist (sec); 'url_hash_fail’: URL hash persist fail; 'header_hash_pri’: Header hash persist(pri); 'header_hash_enqueue’: Header hash persist(enQ); 'header_hash_sec’: Header hash persist(sec); 'header_hash_fail’: Header hash persist fail; 'src_ip’: SRC IP persist ok; 'src_ip_enqueue’: SRC IP persist enqueue; 'src_ip_fail’: SRC IP persist fail; 'src_ip_new_sess_cache’: SRC IP new sess (cache); 'src_ip_new_sess_cache_fail’: SRC IP new sess fail ©; 'src_ip_new_sess_sel’: SRC IP new sess (select); 'src_ip_new_sess_sel_fail’: SRC IP new sess fail (s); 'src_ip_hash_pri’: SRC IP hash persist(pri); 'src_ip_hash_enqueue’: SRC IP hash persist(enQ); 'src_ip_hash_sec’: SRC IP hash persist(sec); 'src_ip_hash_fail’: SRC IP hash persist fail; 'src_ip_enforce’: Enforce higher priority; 'dst_ip’: DST IP persist ok; 'dst_ip_enqueue’: DST IP persist enqueue; 'dst_ip_fail’: DST IP persist fail; 'dst_ip_new_sess_cache’: DST IP new sess (cache); 'dst_ip_new_sess_cache_fail’: DST IP new sess fail ©; 'dst_ip_new_sess_sel’: DST IP new sess (select); 'dst_ip_new_sess_sel_fail’: DST IP new sess fail (s); 'dst_ip_hash_pri’: DST IP hash persist(pri); 'dst_ip_hash_enqueue’: DST IP hash persist(enQ); 'dst_ip_hash_sec’: DST IP hash persist(sec); 'dst_ip_hash_fail’: DST IP hash persist fail; 'cssl_sid_not_found’: Client SSL SID not found; 'cssl_sid_match’: Client SSL SID match; 'cssl_sid_not_match’: Client SSL SID not match; 'sssl_sid_not_found’: Server SSL SID not found; 'sssl_sid_reset’: Server SSL SID reset; 'sssl_sid_match’: Server SSL SID match; 'sssl_sid_not_match’: Server SSL SID not match; 'ssl_sid_persist_ok’: SSL SID persist ok; 'ssl_sid_persist_fail’: SSL SID persist fail; 'ssl_sid_session_ok’: Create SSL SID ok; 'ssl_sid_session_fail’: Create SSL SID fail; 'cookie_persist_ok’: Cookie persist ok; 'cookie_persist_fail’: Cookie persist fail; 'cookie_not_found’: Persist cookie not found; 'cookie_pass_thru’: Persist cookie Pass-thru; 'cookie_invalid’: Invalid persist cookie;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_player_id_global",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB player id global resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"player",
				"id",
				"global",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enforcement_timer",
					Description: `Time to playerid enforcement after bootup (default 480 seconds) (Time to playerid enforcement in seconds, default 480)`,
				},
				resource.Attribute{
					Name:        "abs_max_expiration",
					Description: `Absolute max record expiration value (default 10 minutes) (Absolute max record expiration time in minutes, default 10)`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'total_playerids_created’: Playerid records created; 'total_playerids_deleted’: Playerid records deleted; 'total_abs_max_age_outs’: Playerid records max time aged out; 'total_pkt_activity_age_outs’: Playerid records idle timeout; 'total_invalid_playerid_pkts’: Invalid playerid packets; 'total_invalid_playerid_drops’: Invalid playerid packet drops; 'total_valid_playerid_pkts’: Valid playerid packets;`,
				},
				resource.Attribute{
					Name:        "force_passive",
					Description: `Forces the device to be in passive mode (Only stats and no packet drops)`,
				},
				resource.Attribute{
					Name:        "pkt_activity_expiration",
					Description: `Packet activity record expiration value (default 5 minutes) (Packet activity record expiration time in minutes, default 5)`,
				},
				resource.Attribute{
					Name:        "min_expiration",
					Description: `Minimum record expiration value (default 1 min) (Min record expiration time in minutes, default 1)`,
				},
				resource.Attribute{
					Name:        "enable_64bit_player_id",
					Description: `Enable 64 bit player id check. Default is 32 bit`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_pop3_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB pop3 proxy resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"pop3",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'num’: Num; 'curr’: Current proxy conns; 'total’: Total proxy conns; 'svrsel_fail’: Server selection failure; 'no_route’: no route failure; 'snat_fail’: source nat failure; 'line_too_long’: line too long; 'line_mem_freed’: request line freed; 'invalid_start_line’: invalid start line; 'stls’: stls cmd; 'request_dont_care’: other cmd; 'unsupported_command’: Unsupported cmd; 'bad_sequence’: Bad Sequence; 'rsv_persist_conn_fail’: Serv Sel Persist fail; 'smp_v6_fail’: Serv Sel SMPv6 fail; 'smp_v4_fail’: Serv Sel SMPv4 fail; 'insert_tuple_fail’: Serv Sel insert tuple fail; 'cl_est_err’: Client EST state erro; 'ser_connecting_err’: Serv CTNG state error; 'server_response_err’: Serv RESP state error; 'cl_request_err’: Client RQ state error; 'request’: Total POP3 Request; 'control_to_ssl’: Control chn ssl;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB proxy resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'num’: Num; 'tcp_event’: TCP stack event; 'est_event’: Connection established; 'data_event’: Data received; 'client_fin’: Client FIN; 'server_fin’: Server FIN; 'wbuf_event’: Ready to send data; 'err_event’: Error occured; 'no_mem’: No memory; 'client_rst’: Client RST; 'server_rst’: Server RST; 'queue_depth_over_limit’: Queue depth over limit; 'event_failed’: Event failed; 'conn_not_exist’: Conn not exist; 'service_alloc_cb’: Service alloc callback; 'service_alloc_cb_failed’: Service alloc callback failed; 'service_free_cb’: Service free callback; 'service_free_cb_failed’: Service free callback failed; 'est_cb_failed’: App EST callback failed; 'data_cb_failed’: App DATA callback failed; 'wbuf_cb_failed’: App WBUF callback failed; 'err_cb_failed’: App ERR callback failed; 'start_server_conn’: Start server conn; 'start_server_conn_succ’: Success; 'start_server_conn_no_route’: No route to server; 'start_server_conn_fail_mem’: No memory; 'start_server_conn_fail_snat’: Failed Source NAT; 'start_server_conn_fail_persist’: Fail Persistence; 'start_server_conn_fail_server’: Fail Server issue; 'start_server_conn_fail_tuple’: Fail Tuple Issue; 'line_too_long’: Line too long;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_rate_limit_log",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB rate limit log resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"rate",
				"limit",
				"log",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'total_log_times’: Total log times; 'total_log_msg’: Total log messages; 'local_log_msg’: Local log messages; 'remote_log_msg’: Remote log messages; 'local_log_rate’: Local rate (per sec); 'remote_log_rate’: Remote rate (per sec); 'msg_too_big’: Log message too big; 'buff_alloc_fail’: Buffer alloc fail; 'no_route’: No route; 'buff_send_fail’: Buffer send fail; 'alloc_conn’: Log-session alloc; 'free_conn’: Log-session free; 'conn_alloc_fail’: Log-session alloc fail; 'no_repeat_msg’: No repeat message; 'local_log_dropped’: Local log dropped due to rate-limit;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_rc_cache_global",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB rc cache global resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"rc",
				"cache",
				"global",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'hits’: Cache Hits; 'miss’: Cache Misses; 'bytes_served’: Bytes Served; 'total_req’: Total Requests; 'caching_req’: Cacheable Requests; 'nc_req_header’: No-cache Request; 'nc_res_header’: Not cacheable; 'rv_success’: Revalidation Successes; 'rv_failure’: Revalidation Failures; 'ims_request’: IMS Requests; 'nm_response’: Responses from cache 304 Not Modified; 'rsp_type_CL’: Responses from server 200 OK - Cont Len; 'rsp_type_CE’: Responses from server 200 OK - Chnk Enc; 'rsp_type_304’: Responses from server 304 Not Modified; 'rsp_type_other’: Responses from server 200 OK - Other; 'rsp_no_compress’: Responses from cache 200 OK - No Comp; 'rsp_gzip’: Responses from cache 200 OK - Gzip; 'rsp_deflate’: Responses from cache 200 OK - Deflate; 'rsp_other’: Responses from cache Other; 'nocache_match’: Policy URI nocache; 'match’: Policy URI cache; 'invalidate_match’: Policy URI invalidate; 'content_toobig’: Policy Content Too Big; 'content_toosmall’: Policy Content Too Small; 'entry_create_failures’: Entry Create failures; 'mem_size’: Memory Used; 'entry_num’: Entry Cached; 'replaced_entry’: Entry Replaced; 'aging_entry’: Entry Aged Out; 'cleaned_entry’: Entry Cleaned;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_resource_usage",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB resource usage resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"resource",
				"usage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "real_server_count",
					Description: `Total Real Servers in the System`,
				},
				resource.Attribute{
					Name:        "stream_template_count",
					Description: `Total configurable Streaming media in the System`,
				},
				resource.Attribute{
					Name:        "proxy_template_count",
					Description: `Total configurable Proxy Templates in the System`,
				},
				resource.Attribute{
					Name:        "http_template_count",
					Description: `Total configurable HTTP Templates in the System`,
				},
				resource.Attribute{
					Name:        "persist_srcip_template_count",
					Description: `Total configurable Source IP Persistent Templates in the System`,
				},
				resource.Attribute{
					Name:        "server_ssl_template_count",
					Description: `Total configurable Server SSL Templates in the System`,
				},
				resource.Attribute{
					Name:        "service_group_count",
					Description: `Total Service Groups in the System`,
				},
				resource.Attribute{
					Name:        "persist_cookie_template_count",
					Description: `Total configurable Persistent cookie Templates in the System`,
				},
				resource.Attribute{
					Name:        "virtual_server_count",
					Description: `Total Virtual Servers in the System`,
				},
				resource.Attribute{
					Name:        "fast_udp_template_count",
					Description: `Total configurable Fast UDP Templates in the System`,
				},
				resource.Attribute{
					Name:        "fast_tcp_template_count",
					Description: `Total configurable Fast TCP Templates in the System`,
				},
				resource.Attribute{
					Name:        "virtual_port_count",
					Description: `Total Virtual Server Ports in the System`,
				},
				resource.Attribute{
					Name:        "conn_reuse_template_count",
					Description: `Total configurable Connection reuse Templates in the System`,
				},
				resource.Attribute{
					Name:        "real_port_count",
					Description: `Total Real Server Ports in the System`,
				},
				resource.Attribute{
					Name:        "client_ssl_template_count",
					Description: `Total configurable Client SSL Templates in the System`,
				},
				resource.Attribute{
					Name:        "nat_pool_addr_count",
					Description: `Total configurable NAT Pool addresses in the System (deprecated)`,
				},
				resource.Attribute{
					Name:        "pbslb_subnet_count",
					Description: `Total PBSLB Subnets in the System`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_sip",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB sip resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"sip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'msg_proxy_current’: Number of current sip proxy connections; 'msg_proxy_total’: Total number of sip proxy connections; 'msg_proxy_mem_allocd’: msg_proxy_mem_allocd; 'msg_proxy_mem_cached’: msg_proxy_mem_cached; 'msg_proxy_mem_freed’: msg_proxy_mem_freed; 'msg_proxy_client_recv’: Number of SIP messages received from client; 'msg_proxy_client_send_success’: Number of SIP messages received from client and forwarded to server; 'msg_proxy_client_incomplete’: Number of packet which contains incomplete message; 'msg_proxy_client_drop’: Number of AX drop; 'msg_proxy_client_connection’: Connecting server; 'msg_proxy_client_fail’: Number of SIP messages received from client but failed to forward to server; 'msg_proxy_client_fail_parse’: msg_proxy_client_fail_parse; 'msg_proxy_client_fail_process’: msg_proxy_client_fail_process; 'msg_proxy_client_fail_snat’: msg_proxy_client_fail_snat; 'msg_proxy_client_exceed_tmp_buff’: msg_proxy_client_exceed_tmp_buff; 'msg_proxy_client_fail_send_pkt’: msg_proxy_client_fail_send_pkt; 'msg_proxy_client_fail_start_server_Conn’: msg_proxy_client_fail_start_server_Conn; 'msg_proxy_server_recv’: Number of SIP messages received from server; 'msg_proxy_server_send_success’: Number of SIP messages received from server and forwarded to client; 'msg_proxy_server_incomplete’: Number of packet which contains incomplete message; 'msg_proxy_server_drop’: Number of AX drop; 'msg_proxy_server_fail’: Number of SIP messages received from server but failed to forward to client; 'msg_proxy_server_fail_parse’: msg_proxy_server_fail_parse; 'msg_proxy_server_fail_process’: msg_proxy_server_fail_process; 'msg_proxy_server_fail_selec_connt’: msg_proxy_server_fail_selec_connt; 'msg_proxy_server_fail_snat’: msg_proxy_server_fail_snat; 'msg_proxy_server_exceed_tmp_buff’: msg_proxy_server_exceed_tmp_buff; 'msg_proxy_server_fail_send_pkt’: msg_proxy_server_fail_send_pkt; 'msg_proxy_create_server_conn’: Number of server connection system tries to create; 'msg_proxy_start_server_conn’: Number of server connection created successfully; 'msg_proxy_fail_start_server_conn’: Number of server connection create failed; 'msg_proxy_server_conn_fail_snat’: msg_proxy_server_conn_fail_snat; 'msg_proxy_fail_construct_server_conn’: msg_proxy_fail_construct_server_conn; 'msg_proxy_fail_reserve_pconn’: msg_proxy_fail_reserve_pconn; 'msg_proxy_start_server_conn_failed’: msg_proxy_start_server_conn_failed; 'msg_proxy_server_conn_already_exists’: msg_proxy_server_conn_already_exists; 'msg_proxy_fail_insert_server_conn’: msg_proxy_fail_insert_server_conn; 'msg_proxy_parse_msg_fail’: msg_proxy_parse_msg_fail; 'msg_proxy_process_msg_fail’: msg_proxy_process_msg_fail; 'msg_proxy_no_vport’: msg_proxy_no_vport; 'msg_proxy_fail_select_server’: msg_proxy_fail_select_server; 'msg_proxy_fail_alloc_mem’: msg_proxy_fail_alloc_mem; 'msg_proxy_unexpected_err’: msg_proxy_unexpected_err; 'msg_proxy_l7_cpu_failed’: msg_proxy_l7_cpu_failed; 'msg_proxy_l4_to_l7’: msg_proxy_l4_to_l7; 'msg_proxy_l4_from_l7’: msg_proxy_l4_from_l7; 'msg_proxy_to_l4_send_pkt’: msg_proxy_to_l4_send_pkt; 'msg_proxy_l4_from_l4_send’: msg_proxy_l4_from_l4_send; 'msg_proxy_l7_to_L4’: msg_proxy_l7_to_L4; 'msg_proxy_mag_back’: msg_proxy_mag_back; 'msg_proxy_fail_dcmsg’: msg_proxy_fail_dcmsg; 'msg_proxy_deprecated_conn’: msg_proxy_deprecated_conn; 'msg_proxy_hold_msg’: msg_proxy_hold_msg; 'msg_proxy_split_pkt’: msg_proxy_split_pkt; 'msg_proxy_pipline_msg’: msg_proxy_pipline_msg; 'msg_proxy_client_reset’: msg_proxy_client_reset; 'msg_proxy_server_reset’: msg_proxy_server_reset; 'session_created’: SIP Session created; 'session_freed’: SIP Session freed; 'session_in_rml’: session_in_rml; 'session_invalid’: session_invalid; 'conn_allocd’: conn_allocd; 'conn_freed’: conn_freed; 'session_callid_allocd’: session_callid_allocd; 'session_callid_freed’: session_callid_freed; 'line_mem_allocd’: line_mem_allocd; 'line_mem_freed’: line_mem_freed; 'table_mem_allocd’: table_mem_allocd; 'table_mem_freed’: table_mem_freed; 'cmsg_no_uri_header’: cmsg_no_uri_header; 'cmsg_no_uri_session’: cmsg_no_uri_session; 'sg_no_uri_header’: sg_no_uri_header; 'smsg_no_uri_session’: smsg_no_uri_session; 'line_too_long’: line_too_long; 'fail_read_start_line’: fail_read_start_line; 'fail_parse_start_line’: fail_parse_start_line; 'invalid_start_line’: invalid_start_line; 'request_unknown_version’: request_unknown_version; 'response_unknown_version’: response_unknown_version; 'request_unknown’: request_unknown; 'fail_parse_headers’: fail_parse_headers; 'too_many_headers’: too_many_headers; 'invalid_header’: invalid_header; 'header_name_too_long’: header_name_too_long; 'body_too_big’: body_too_big; 'fail_get_counter’: fail_get_counter; 'msg_no_call_id’: msg_no_call_id; 'identify_dir_failed’: identify_dir_failed; 'no_sip_request’: no_sip_request; 'deprecated_msg’: deprecated_msg; 'fail_insert_callid_session’: fail_insert_callid_session; 'fail_insert_uri_session’: fail_insert_uri_session; 'fail_insert_header’: fail_insert_header; 'select_server_conn’: select_server_conn; 'select_server_conn_by_callid’: select_server_conn_by_callid; 'select_server_conn_by_uri’: select_server_conn_by_uri; 'select_server_conn_by_rev_tuple’: select_server_conn_by_rev_tuple; 'select_server_conn_failed’: select_server_conn_failed; 'select_client_conn’: select_client_conn; 'X_forward_for_select_client’: X_forward_for_select_client; 'call_id_select_client’: call_id_select_client; 'uri_select_client’: uri_select_client; 'client_select_failed’: client_select_failed; 'acl_denied’: acl_denied; 'assemble_frag_failed’: assemble_frag_failed; 'wrong_ip_version’: wrong_ip_version; 'size_too_large’: size_too_large; 'fail_split_fragment’: fail_split_fragment; 'client_keepalive_received’: client_keepalive_received; 'server_keepalive_received’: server_keepalive_received; 'client_keepalive_send’: client_keepalive_send; 'server_keepalive_send’: server_keepalive_send; 'ax_health_check_received’: ax_health_check_received; 'client_request’: client_request; 'client_request_ok’: client_request_ok; 'concatenate_msg’: concatenate_msg; 'save_uri’: save_uri; 'save_uri_ok’: save_uri_ok; 'save_call_id’: save_call_id; 'save_call_id_ok’: save_call_id_ok; 'msg_translation’: msg_translation; 'msg_translation_fail’: msg_translation_fail; 'msg_trans_start_line’: msg_trans_start_line; 'msg_trans_start_headers’: msg_trans_start_headers; 'msg_trans_body’: msg_trans_body; 'request_register’: request_register; 'request_invite’: request_invite; 'request_ack’: request_ack; 'request_cancel’: request_cancel; 'request_bye’: request_bye; 'request_options’: request_options; 'request_prack’: request_prack; 'request_subscribe’: request_subscribe; 'request_notify’: request_notify; 'request_publish’: request_publish; 'request_info’: request_info; 'request_refer’: request_refer; 'request_message’: request_message; 'request_update’: request_update; 'response_unknown’: response_unknown; 'response_1XX’: response_1XX; 'response_2XX’: response_2XX; 'response_3XX’: response_3XX; 'response_4XX’: response_4XX; 'response_5XX’: response_5XX; 'response_6XX’: response_6XX; 'ha_send_sip_session’: ha_send_sip_session; 'ha_send_sip_session_ok’: ha_send_sip_session_ok; 'ha_fail_get_msg_header’: ha_fail_get_msg_header; 'ha_recv_sip_session’: ha_recv_sip_session; 'ha_insert_sip_session_ok’: ha_insert_sip_session_ok; 'ha_update_sip_session_ok’: ha_update_sip_session_ok; 'ha_invalid_pkt’: ha_invalid_pkt; 'ha_fail_alloc_sip_session’: ha_fail_alloc_sip_session; 'ha_fail_alloc_call_id’: ha_fail_alloc_call_id; 'ha_fail_clone_sip_session’: ha_fail_clone_sip_session; 'save_smp_call_id_rtp’: save_smp_call_id_rtp; 'update_smp_call_id_rtp’: update_smp_call_id_rtp; 'smp_call_id_rtp_session_match’: smp_call_id_rtp_session_match; 'smp_call_id_rtp_session_not_match’: smp_call_id_rtp_session_not_match; 'process_error_when_message_switch’: process_error_when_message_switch;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_smpp",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB smpp resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"smpp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'msg_proxy_current’: Curr SMPP Proxy; 'msg_proxy_total’: Total SMPP Proxy; 'msg_proxy_mem_allocd’: msg_proxy_mem_allocd; 'msg_proxy_mem_cached’: msg_proxy_mem_cached; 'msg_proxy_mem_freed’: msg_proxy_mem_freed; 'msg_proxy_client_recv’: Client message rcvd; 'msg_proxy_client_send_success’: Sent to server; 'msg_proxy_client_incomplete’: Incomplete; 'msg_proxy_client_drop’: AX responds directly; 'msg_proxy_client_connection’: Connecting server; 'msg_proxy_client_fail’: Number of SMPP messages received from client but failed to forward to server; 'msg_proxy_client_fail_parse’: msg_proxy_client_fail_parse; 'msg_proxy_client_fail_process’: msg_proxy_client_fail_process; 'msg_proxy_client_fail_snat’: msg_proxy_client_fail_snat; 'msg_proxy_client_exceed_tmp_buff’: msg_proxy_client_exceed_tmp_buff; 'msg_proxy_client_fail_send_pkt’: msg_proxy_client_fail_send_pkt; 'msg_proxy_client_fail_start_server_Conn’: msg_proxy_client_fail_start_server_Conn; 'msg_proxy_server_recv’: Server message rcvd; 'msg_proxy_server_send_success’: Sent to client; 'msg_proxy_server_incomplete’: Incomplete; 'msg_proxy_server_drop’: Number of the packet AX drop; 'msg_proxy_server_fail’: Number of SMPP messages received from server but failed to forward to client; 'msg_proxy_server_fail_parse’: msg_proxy_server_fail_parse; 'msg_proxy_server_fail_process’: msg_proxy_server_fail_process; 'msg_proxy_server_fail_selec_connt’: msg_proxy_server_fail_selec_connt; 'msg_proxy_server_fail_snat’: msg_proxy_server_fail_snat; 'msg_proxy_server_exceed_tmp_buff’: msg_proxy_server_exceed_tmp_buff; 'msg_proxy_server_fail_send_pkt’: msg_proxy_server_fail_send_pkt; 'msg_proxy_create_server_conn’: Server conn created; 'msg_proxy_start_server_conn’: Number of server connection created successfully; 'msg_proxy_fail_start_server_conn’: Number of server connection created failed; 'msg_proxy_server_conn_fail_snat’: msg_proxy_server_conn_fail_snat; 'msg_proxy_fail_construct_server_conn’: msg_proxy_fail_construct_server_conn; 'msg_proxy_fail_reserve_pconn’: msg_proxy_fail_reserve_pconn; 'msg_proxy_start_server_conn_failed’: msg_proxy_start_server_conn_failed; 'msg_proxy_server_conn_already_exists’: msg_proxy_server_conn_already_exists; 'msg_proxy_fail_insert_server_conn’: msg_proxy_fail_insert_server_conn; 'msg_proxy_parse_msg_fail’: msg_proxy_parse_msg_fail; 'msg_proxy_process_msg_fail’: msg_proxy_process_msg_fail; 'msg_proxy_no_vport’: msg_proxy_no_vport; 'msg_proxy_fail_select_server’: msg_proxy_fail_select_server; 'msg_proxy_fail_alloc_mem’: msg_proxy_fail_alloc_mem; 'msg_proxy_unexpected_err’: msg_proxy_unexpected_err; 'msg_proxy_l7_cpu_failed’: msg_proxy_l7_cpu_failed; 'msg_proxy_l4_to_l7’: msg_proxy_l4_to_l7; 'msg_proxy_l4_from_l7’: msg_proxy_l4_from_l7; 'msg_proxy_to_l4_send_pkt’: msg_proxy_to_l4_send_pkt; 'msg_proxy_l4_from_l4_send’: msg_proxy_l4_from_l4_send; 'msg_proxy_l7_to_L4’: msg_proxy_l7_to_L4; 'msg_proxy_mag_back’: msg_proxy_mag_back; 'msg_proxy_fail_dcmsg’: msg_proxy_fail_dcmsg; 'msg_proxy_deprecated_conn’: msg_proxy_deprecated_conn; 'msg_proxy_hold_msg’: msg_proxy_hold_msg; 'msg_proxy_split_pkt’: msg_proxy_split_pkt; 'msg_proxy_pipline_msg’: msg_proxy_pipline_msg; 'msg_proxy_client_reset’: msg_proxy_client_reset; 'msg_proxy_server_reset’: msg_proxy_server_reset; 'payload_allocd’: payload_allocd; 'payload_freed’: payload_freed; 'pkt_too_small’: pkt_too_small; 'invalid_seq’: invalid_seq; 'AX_response_directly’: Number of packet which AX responds directly; 'select_client_conn’: Client conn selection; 'select_client_by_req’: Select by request; 'select_client_from_list’: Select by roundbin; 'select_client_by_conn’: Select by conn; 'select_client_fail’: Select failed; 'select_server_conn’: Server conn selection; 'select_server_by_req’: Select by request; 'select_server_from_list’: Select by roundbin; 'select_server_by_conn’: Select server conn by client conn; 'select_server_fail’: Fail to select server conn; 'bind_conn’: bind_conn; 'unbind_conn’: unbind_conn; 'enquire_link_recv’: enquire_link_recv; 'enquire_link_resp_recv’: enquire_link_resp_recv; 'enquire_link_send’: enquire_link_send; 'enquire_link_resp_send’: enquire_link_resp_send; 'client_conn_put_in_list’: client_conn_put_in_list; 'client_conn_get_from_list’: client_conn_get_from_list; 'server_conn_put_in_list’: server_conn_put_in_list; 'server_conn_get_from_list’: server_conn_get_from_list; 'server_conn_fail_bind’: server_conn_fail_bind; 'single_msg’: single_msg; 'fail_bind_msg’: fail_bind_msg;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_smtp",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB smtp resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"smtp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'curr_proxy’: Current proxy conns; 'total_proxy’: Total proxy conns; 'request’: SMTP requests; 'request_success’: SMTP requests (success); 'no_proxy’: No proxy error; 'client_reset’: Client reset; 'server_reset’: Server reset; 'no_tuple’: No tuple error; 'parse_req_fail’: Parse request failure; 'server_select_fail’: Server selection failure; 'forward_req_fail’: Forward request failure; 'forward_req_data_fail’: Forward REQ data failure; 'req_retran’: Request retransmit; 'req_ofo’: Request pkt out-of-order; 'server_reselect’: Server reselection; 'server_prem_close’: Server premature close; 'new_server_conn’: Server connection made; 'snat_fail’: Source NAT failure; 'tcp_out_reset’: TCP out reset; 'recv_client_command_EHLO’: Recv client EHLO; 'recv_client_command_HELO’: Recv client HELO; 'recv_client_command_MAIL’: Recv client MAIL; 'recv_client_command_RCPT’: Recv client RCPT; 'recv_client_command_DATA’: Recv client DATA; 'recv_client_command_RSET’: Recv client RSET; 'recv_client_command_VRFY’: Recv client VRFY; 'recv_client_command_EXPN’: Recv client EXPN; 'recv_client_command_HELP’: Recv client HELP; 'recv_client_command_NOOP’: Recv client NOOP; 'recv_client_command_QUIT’: Recv client QUIT; 'recv_client_command_STARTTLS’: Recv client STARTTLS; 'recv_client_command_others’: Recv client other cmds; 'recv_server_service_not_ready’: Recv server serv-not-rdy; 'recv_server_unknow_reply_code’: Recv server unknown-code; 'send_client_service_ready’: Sent client serv-rdy; 'send_client_service_not_ready’: Sent client serv-not-rdy; 'send_client_close_connection’: Sent client close-conn; 'send_client_go_ahead’: Sent client go-ahead; 'send_client_start_TLS_first’: Sent client STARTTLS-1st; 'send_client_TLS_not_available’: Sent client TLS-not-aval; 'send_client_no_command’: Sent client no-such-cmd; 'send_server_cmd_reset’: Sent server RSET; 'TLS_established’: SSL session established; 'L4_switch’: L4 switching; 'Aflex_switch’: aFleX switching; 'Aflex_switch_ok’: aFleX switching (succ); 'client_domain_switch’: Client domain switching; 'client_domain_switch_ok’: Client domain sw (succ); 'LB_switch’: LB switching; 'LB_switch_ok’: LB switching (succ); 'read_request_line_fail’: Read request line fail; 'get_all_headers_fail’: Get all headers fail; 'too_many_headers’: Too many headers; 'line_too_long’: Line too long; 'line_across_packet’: Line across packets; 'line_extend’: Line extend; 'line_extend_fail’: Line extend fail; 'line_table_extend’: Table extend; 'line_table_extend_fail’: Table extend fail; 'parse_request_line_fail’: Parse request line fail; 'insert_resonse_line_fail’: Ins response line fail; 'remove_resonse_line_fail’: Del response line fail; 'parse_resonse_line_fail’: Parse response line fail; 'Aflex_lb_reselect’: aFleX lb reselect; 'Aflex_lb_reselect_ok’: aFleX lb reselect (succ); 'server_STARTTLS_init’: Init server side STARTTLS; 'server_STARTTLS_fail’: Server side STARTTLS fail; 'rserver_STARTTLS_disable’: real server not support STARTTLS; 'recv_client_command_TURN’: Recv client TURN; 'recv_client_command_ETRN’: Recv client ETRN;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_spdy_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB spdy proxy resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"spdy",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'curr_proxy’: Curr Proxy Conns; 'total_proxy’: Total Proxy Conns; 'curr_http_proxy’: Curr HTTP Proxy Conns; 'total_http_proxy’: Total HTTP Proxy Conns; 'total_v2_proxy’: Version 2 Streams; 'total_v3_proxy’: Version 3 Streams; 'curr_stream’: Curr Streams; 'total_stream’: Total Streams; 'total_stream_succ’: Streams(succ); 'client_rst’: client_rst; 'server_rst’: Server RST sent; 'client_goaway’: client_goaway; 'server_goaway’: Server GOAWAY sent; 'tcp_err’: TCP sock error; 'inflate_ctx’: Inflate context; 'deflate_ctx’: Deflate context; 'ping_sent’: PING sent; 'stream_not_found’: STREAM not found; 'client_fin’: Client FIN; 'server_fin’: Server FIN; 'stream_close’: Stream close; 'stream_err’: Stream err; 'session_err’: Session err; 'control_frame’: Control frame received; 'syn_frame’: SYN stream frame received; 'syn_reply_frame’: SYN reply frame received; 'headers_frame’: Headers frame received; 'settings_frame’: Setting frame received; 'window_frame’: Window update frame received; 'ping_frame’: Ping frame received; 'data_frame’: Data frame received; 'data_no_stream’: Data no stream found; 'data_no_stream_no_goaway’: Data no stream and no goaway; 'data_no_stream_goaway_close’: Data no stream and no goaway and close session; 'est_cb_no_tuple’: Est callback no tuple; 'data_cb_no_tuple’: Data callback no tuple; 'ctx_alloc_fail’: Context alloc fail; 'fin_close_session’: FIN close session; 'server_rst_close_stream’: Server RST close stream; 'stream_found’: Stream found; 'close_stream_session_not_found’: Close stream session not found; 'close_stream_stream_not_found’: Close stream stream not found; 'close_stream_already_closed’: Closing closed stream; 'close_stream_session_close’: Stream close session close; 'close_session_already_closed’: Closing closed session; 'max_concurrent_stream_limit’: Max concurrent stream limit; 'stream_alloc_fail’: Stream alloc fail; 'http_conn_alloc_fail’: HTTP connection allocation fail; 'request_header_alloc_fail’: Request/Header allocation fail; 'name_value_total_len_ex’: Name value total length exceeded; 'name_value_zero_len’: Name value zero name length; 'name_value_invalid_http_ver’: Name value invalid http version; 'name_value_connection’: Name value connection; 'name_value_keepalive’: Name value keep alive; 'name_value_proxy_conn’: Name value proxy-connection; 'name_value_trasnfer_encod’: Name value transfer encoding; 'name_value_no_must_have’: Name value no must have; 'decompress_fail’: Decompress fail; 'syn_after_goaway’: SYN after goaway; 'stream_lt_prev’: Stream id less than previous; 'syn_stream_exist_or_even’: Stream already exists; 'syn_unidir’: Unidirectional SYN; 'syn_reply_alr_rcvd’: SYN reply already received; 'client_rst_nostream’: Close RST stream not found; 'window_no_stream’: Window update no stream found; 'invalid_window_size’: Invalid window size; 'unknown_control_frame’: Unknown control frame; 'data_on_closed_stream’: Data on closed stream; 'invalid_frame_size’: Invalid frame size; 'invalid_version’: Invalid version; 'header_after_session_close’: Header after session close; 'compress_ctx_alloc_fail’: Compression context allocation fail; 'header_compress_fail’: Header compress fail; 'http_data_session_close’: HTTP data session close; 'http_data_stream_not_found’: HTTP data stream not found; 'close_stream_not_http_proxy’: Close Stream not http-proxy; 'session_needs_requeue’: Session needs requeue; 'new_stream_session_del’: New Stream after Session delete; 'fin_stream_closed’: HTTP FIN stream already closed; 'http_close_stream_closed’: HTTP close stream already closed; 'http_err_stream_closed’: HTTP error stream already closed; 'http_hdr_stream_close’: HTTP header stream already closed; 'http_data_stream_close’: HTTP data stream already closed; 'session_close’: Session close;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_sport_rate_limit",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB sport rate limit resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"sport",
				"rate",
				"limit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'alloc_sport’: Alloc’d src port entry; 'alloc_sportip’: Alloc’d src port-ip entry; 'freed_sport’: Freed src port entry; 'freed_sportip’: Freed src port-ip entry; 'total_drop’: Total rate exceed drop; 'total_reset’: Total rate exceed reset; 'total_log’: Total log sent;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_ssl_cert_revoke",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB ssl cert revoke resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"ssl",
				"cert",
				"revoke",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'ocsp_stapling_response_good’: OCSP stapling response good; 'ocsp_chain_status_good’: Certificate chain status good; 'ocsp_chain_status_revoked’: Certificate chain status revoked; 'ocsp_chain_status_unknown’: Certificate chain status unknown; 'ocsp_request’: OCSP requests; 'ocsp_response’: OCSP responses; 'ocsp_connection_error’: OCSP connection error; 'ocsp_uri_not_found’: OCSP URI not found; 'ocsp_uri_https’: Log OCSP URI https; 'ocsp_uri_unsupported’: OCSP URI unsupported; 'ocsp_response_status_good’: OCSP response status good; 'ocsp_response_status_revoked’: OCSP response status revoked; 'ocsp_response_status_unknown’: OCSP response status unknown; 'ocsp_cache_status_good’: OCSP cache status good; 'ocsp_cache_status_revoked’: OCSP cache status revoked; 'ocsp_cache_miss’: OCSP cache miss; 'ocsp_cache_expired’: OCSP cache expired; 'ocsp_other_error’: Log OCSP other errors; 'ocsp_response_no_nonce’: Log OCSP other errors; 'ocsp_response_nonce_error’: Log OCSP other errors; 'crl_request’: CRL requests; 'crl_response’: CRL responses; 'crl_connection_error’: CRL connection errors; 'crl_uri_not_found’: CRL URI not found; 'crl_uri_https’: CRL URI https; 'crl_uri_unsupported’: CRL URI unsupported; 'crl_response_status_good’: CRL response status good; 'crl_response_status_revoked’: CRL response status revoked; 'crl_response_status_unknown’: CRL response status unknown; 'crl_cache_status_good’: CRL cache status good; 'crl_cache_status_revoked’: CRL cache status revoked; 'crl_other_error’: CRL other errors;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_ssl_expire_check",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB ssl expire check resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"ssl",
				"expire",
				"check",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ssl_expire_email_address",
					Description: `Config Email address for certificate expiration check`,
				},
				resource.Attribute{
					Name:        "before",
					Description: `The number of days in advance notice before expiration, default is 5`,
				},
				resource.Attribute{
					Name:        "interval_days",
					Description: `The interval of days notice after expiration, default is 2`,
				},
				resource.Attribute{
					Name:        "expire_address1",
					Description: `Email address for certificate expiration check`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_ssl_forward_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB ssl forward proxy resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"ssl",
				"forward",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'cert_create’: Certificates created; 'cert_expr’: Certificates expired; 'cert_hit’: Certificate cache hits; 'cert_miss’: Certificate cache miss; 'conn_bypass’: Connections bypassed; 'conn_inspect’: Connections inspected; 'bypass-failsafe-ssl-sessions’: Bypass Failsafe SSL sessions; 'bypass-sni-sessions’: Bypass SNI sessions; 'bypass-client-auth-sessions’: Bypass Client Auth sessions; 'failed-in-ssl-handshakes’: Failed in SSL handshakes; 'failed-in-crypto-operations’: Failed in crypto operations; 'failed-in-tcp’: Failed in TCP; 'failed-in-certificate-verification’: Failed in Certificate verification; 'failed-in-certificate-signing’: Failed in Certificate signing; 'invalid-ocsp-stapling-response’: Invalid OCSP Stapling Response; 'revoked-ocsp-response’: Revoked OCSP Response; 'unsupported-ssl-version’: Unsupported SSL version; 'certificates-in-cache’: Certificates in cache; 'connections-failed’: Connections failed; 'aflex-bypass’: Bypass triggered by aFleX; 'bypass-cert-subject-sessions’: Bypass Cert Subject sessions; 'bypass-cert-issuer-sessions’: Bypass Cert issuer sessions; 'bypass-cert-san-sessions’: Bypass Cert SAN sessions; 'bypass-no-sni-sessions’: Bypass NO SNI sessions; 'reset-no-sni-sessions’: Reset No SNI sessions; 'bypass-username-sessions’: Bypass Username sessions; 'bypass-ad-group-sessions’: Bypass AD-group sessions; 'cert_in_cache’: Certificates in cache; 'tot_conn_in_buff’: Total buffered async connections; 'curr_conn_in_buff’: Current buffered async connections;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_svm_source_nat",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB svm source nat resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"svm",
				"source",
				"nat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pool",
					Description: `Specify NAT pool or pool group`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_switch",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB switch resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"switch",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'fwlb’: FWLB; 'licexpire_drop’: License Expire Drop; 'bwl_drop’: BW Limit Drop; 'rx_kernel’: Received kernel; 'rx_arp_req’: ARP REQ Rcvd; 'rx_arp_resp’: ARP RESP Rcvd; 'vlan_flood’: VLAN Flood; 'l2_def_vlan_drop’: L2 Default Vlan FWD Drop; 'ipv4_noroute_drop’: IPv4 No Route Drop; 'ipv6_noroute_drop’: IPv6 No Route Drop; 'prot_down_drop’: Prot Down Drop; 'l2_forward’: L2 Forward; 'l3_forward_ip’: L3 IP Forward; 'l3_forward_ipv6’: L3 IPv6 Forward; 'l4_process’: L4 Process; 'unknown_prot_drop’: Unknown Prot Drop; 'ttl_exceeded_drop’: TTL Exceeded Drop; 'linkdown_drop’: Link Down Drop; 'sport_drop’: SPORT Drop; 'incorrect_len_drop’: Incorrect Length Drop; 'ip_defrag’: IP Defrag; 'acl_deny’: ACL Denys; 'ipfrag_tcp’: IP(TCP) Fragment Rcvd; 'ipfrag_overlap’: IP Fragment Overlap; 'ipfrag_timeout’: IP Fragment Timeout; 'ipfrag_overload’: IP Frag Overload Drops; 'ipfrag_reasmoks’: IP Fragment Reasm OKs; 'ipfrag_reasmfails’: IP Fragment Reasm Fails; 'land_drop’: Anomaly Land Attack Drop; 'ipoptions_drop’: Anomaly IP OPT Drops; 'badpkt_drop’: Bad Pkt Drop; 'pingofdeath_drop’: Anomaly PingDeath Drop; 'allfrag_drop’: Anomaly All Frag Drop; 'tcpnoflag_drop’: Anomaly TCP noFlag Drop; 'tcpsynfrag_drop’: Anomaly SYN Frag Drop; 'tcpsynfin_drop’: Anomaly TCP SYNFIN Drop; 'ipsec_drop’: IPSec Drop; 'bpdu_rcvd’: BPDUs Received; 'bpdu_sent’: BPDUs Sent; 'ctrl_syn_rate_drop’: SYN rate exceeded Drop; 'ip_defrag_invalid_len’: IP Invalid Length Frag; 'ipv4_frag_6rd_ok’: IPv4 Frag 6RD OK; 'ipv4_frag_6rd_drop’: IPv4 Frag 6RD Dropped; 'no_ip_drop’: No IP Drop; 'ipv6frag_udp’: IPv6 Frag UDP; 'ipv6frag_udp_dropped’: IPv6 Frag UDP Dropped; 'ipv6frag_tcp_dropped’: IPv6 Frag TCP Dropped; 'ipv6frag_ipip_ok’: IPv6 Frag IPIP OKs; 'ipv6frag_ipip_dropped’: IPv6 Frag IPIP Drop; 'ip_frag_oversize’: IP Fragment oversize; 'ip_frag_too_many’: IP Fragment too many; 'ipv4_novlanfwd_drop’: IPv4 No L3 VLAN FWD Drop; 'ipv6_novlanfwd_drop’: IPv6 No L3 VLAN FWD Drop; 'fpga_error_pkt1’: FPGA Error PKT1; 'fpga_error_pkt2’: FPGA Error PKT2; 'max_arp_drop’: Max ARP Drop; 'ipv6frag_tcp’: IPv6 Frag TCP; 'ipv6frag_icmp’: IPv6 Frag ICMP; 'ipv6frag_ospf’: IPv6 Frag OSPF; 'ipv6frag_esp’: IPv6 Frag ESP; 'l4_in_ctrl_cpu’: L4 In Ctrl CPU; 'mgmt_svc_drop’: Management Service Drop; 'jumbo_frag_drop’: Jumbo Frag Drop; 'ipv6_jumbo_frag_drop’: IPv6 Jumbo Frag Drop; 'ipipv6_jumbo_frag_drop’: IPIPv6 Jumbo Frag Drop; 'ipv6_ndisc_dad_solicits’: IPv6 DAD on Solicits; 'ipv6_ndisc_dad_adverts’: IPv6 DAD on Adverts; 'ipv6_ndisc_mac_changes’: IPv6 DAD MAC Changed; 'ipv6_ndisc_out_of_memory’: IPv6 DAD Out-of-memory; 'sp_non_ctrl_pkt_drop’: Shared IP mode non ctrl packet to linux drop; 'urpf_pkt_drop’: URPF check packet drop; 'fw_smp_zone_mismatch’: FW SMP Zone Mismatch; 'ipfrag_udp’: IP(UDP) Fragment Rcvd; 'ipfrag_icmp’: IP(ICMP) Fragment Rcvd; 'ipfrag_ospf’: IP(OSPF) Fragment Rcvd; 'ipfrag_esp’: IP(ESP) Fragment Rcvd; 'ipfrag_tcp_dropped’: IP Frag TCP Dropped; 'ipfrag_udp_dropped’: IP Frag UDP Dropped; 'ipfrag_ipip_dropped’: IP Frag IPIP Drop; 'redirect_fwd_fail’: Redirect failed in the fwd direction; 'redirect_fwd_sent’: Redirect succeeded in the fwd direction; 'redirect_rev_fail’: Redirect failed in the rev direction; 'redirect_rev_sent’: Redirect succeeded in the rev direction; 'redirect_setup_fail’: Redirect connection setup failed; 'ip_frag_sent’: IP frag sent; 'invalid_rx_arp_pkt’: Invalid ARP PKT Rcvd;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_cache",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template cache resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"cache",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specify cache template name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "accept_reload_req",
					Description: `Accept reload requests via cache-control directives in HTTP headers`,
				},
				resource.Attribute{
					Name:        "default_policy_nocache",
					Description: `Specify default policy to be to not cache`,
				},
				resource.Attribute{
					Name:        "age",
					Description: `Specify duration in seconds cached content valid, default is 3600 seconds (seconds that the cached content is valid (default 3600 seconds))`,
				},
				resource.Attribute{
					Name:        "disable_insert_via",
					Description: `Disable insertion of via header in response served from RAM cache`,
				},
				resource.Attribute{
					Name:        "replacement_policy",
					Description: `'LFU’: LFU;`,
				},
				resource.Attribute{
					Name:        "disable_insert_age",
					Description: `Disable insertion of age header in response served from RAM cache`,
				},
				resource.Attribute{
					Name:        "max_content_size",
					Description: `Maximum size (bytes) of response that can be cached - default 81920 (80KB)`,
				},
				resource.Attribute{
					Name:        "max_cache_size",
					Description: `Specify maximum cache size in megabytes, default is 80MB (RAM cache size in megabytes (default 80MB))`,
				},
				resource.Attribute{
					Name:        "remove_cookies",
					Description: `Remove cookies in response and cache`,
				},
				resource.Attribute{
					Name:        "verify_host",
					Description: `Verify request using host before sending response from RAM cache`,
				},
				resource.Attribute{
					Name:        "min_content_size",
					Description: `Minimum size (bytes) of response that can be cached - default 512`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all’: all; 'hits’: Cache hits; 'miss’: Cache misses; 'bytes_served’: Bytes served from cache; 'total_req’: Total requests received; 'caching_req’: Total requests to cache; 'nc_req_header’: nc_req_header; 'nc_res_header’: nc_res_header; 'rv_success’: rv_success; 'rv_failure’: rv_failure; 'ims_request’: ims_request; 'nm_response’: nm_response; 'rsp_type_CL’: rsp_type_CL; 'rsp_type_CE’: rsp_type_CE; 'rsp_type_304’: rsp_type_304; 'rsp_type_other’: rsp_type_other; 'rsp_no_compress’: rsp_no_compress; 'rsp_gzip’: rsp_gzip; 'rsp_deflate’: rsp_deflate; 'rsp_other’: rsp_other; 'nocache_match’: nocache_match; 'match’: match; 'invalidate_match’: invalidate_match; 'content_toobig’: content_toobig; 'content_toosmall’: content_toosmall; 'entry_create_failures’: entry_create_failures; 'mem_size’: mem_size; 'entry_num’: entry_num; 'replaced_entry’: replaced_entry; 'aging_entry’: aging_entry; 'cleaned_entry’: cleaned_entry;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_cipher",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template cipher resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"cipher",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Cipher Template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Cipher priority (Cipher priority (default 1))`,
				},
				resource.Attribute{
					Name:        "cipher_suite",
					Description: `'SSL3_RSA_DES_192_CBC3_SHA’: SSL3_RSA_DES_192_CBC3_SHA; 'SSL3_RSA_RC4_128_MD5’: SSL3_RSA_RC4_128_MD5; 'SSL3_RSA_RC4_128_SHA’: SSL3_RSA_RC4_128_SHA; 'TLS1_RSA_AES_128_SHA’: TLS1_RSA_AES_128_SHA; 'TLS1_RSA_AES_256_SHA’: TLS1_RSA_AES_256_SHA; 'TLS1_RSA_AES_128_SHA256’: TLS1_RSA_AES_128_SHA256; 'TLS1_RSA_AES_256_SHA256’: TLS1_RSA_AES_256_SHA256; 'TLS1_DHE_RSA_AES_128_GCM_SHA256’: TLS1_DHE_RSA_AES_128_GCM_SHA256; 'TLS1_DHE_RSA_AES_128_SHA’: TLS1_DHE_RSA_AES_128_SHA; 'TLS1_DHE_RSA_AES_128_SHA256’: TLS1_DHE_RSA_AES_128_SHA256; 'TLS1_DHE_RSA_AES_256_GCM_SHA384’: TLS1_DHE_RSA_AES_256_GCM_SHA384; 'TLS1_DHE_RSA_AES_256_SHA’: TLS1_DHE_RSA_AES_256_SHA; 'TLS1_DHE_RSA_AES_256_SHA256’: TLS1_DHE_RSA_AES_256_SHA256; 'TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256’: TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256; 'TLS1_ECDHE_ECDSA_AES_128_SHA’: TLS1_ECDHE_ECDSA_AES_128_SHA; 'TLS1_ECDHE_ECDSA_AES_128_SHA256’: TLS1_ECDHE_ECDSA_AES_128_SHA256; 'TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384’: TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384; 'TLS1_ECDHE_ECDSA_AES_256_SHA’: TLS1_ECDHE_ECDSA_AES_256_SHA; 'TLS1_ECDHE_RSA_AES_128_GCM_SHA256’: TLS1_ECDHE_RSA_AES_128_GCM_SHA256; 'TLS1_ECDHE_RSA_AES_128_SHA’: TLS1_ECDHE_RSA_AES_128_SHA; 'TLS1_ECDHE_RSA_AES_128_SHA256’: TLS1_ECDHE_RSA_AES_128_SHA256; 'TLS1_ECDHE_RSA_AES_256_GCM_SHA384’: TLS1_ECDHE_RSA_AES_256_GCM_SHA384; 'TLS1_ECDHE_RSA_AES_256_SHA’: TLS1_ECDHE_RSA_AES_256_SHA; 'TLS1_RSA_AES_128_GCM_SHA256’: TLS1_RSA_AES_128_GCM_SHA256; 'TLS1_RSA_AES_256_GCM_SHA384’: TLS1_RSA_AES_256_GCM_SHA384; 'TLS1_ECDHE_RSA_AES_256_SHA384’: TLS1_ECDHE_RSA_AES_256_SHA384; 'TLS1_ECDHE_ECDSA_AES_256_SHA384’: TLS1_ECDHE_ECDSA_AES_256_SHA384; 'TLS1_ECDHE_RSA_CHACHA20_POLY1305_SHA256’: TLS1_ECDHE_RSA_CHACHA20_POLY1305_SHA256; 'TLS1_ECDHE_ECDSA_CHACHA20_POLY1305_SHA256’: TLS1_ECDHE_ECDSA_CHACHA20_POLY1305_SHA256; 'TLS1_DHE_RSA_CHACHA20_POLY1305_SHA256’: TLS1_DHE_RSA_CHACHA20_POLY1305_SHA256;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_client_ssh",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template client-ssh resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"client",
				"ssh",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Client SSH Template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "forward_proxy_enable",
					Description: `Enable SSH forward proxy`,
				},
				resource.Attribute{
					Name:        "forward_proxy_hostkey",
					Description: `Specify private-key`,
				},
				resource.Attribute{
					Name:        "passphrase",
					Description: `Password Phrase`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_client_ssl",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template client-ssl resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"client",
				"ssl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Client SSL Template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "forward_proxy_ssl_version",
					Description: `TLS/SSL version, default is TLS1.2 (TLS/SSL version: 31-TLSv1.0, 32-TLSv1.1 and 33-TLSv1.2)`,
				},
				resource.Attribute{
					Name:        "forward_proxy_crl_disable",
					Description: `Disable Certificate Revocation List checking for forward proxy`,
				},
				resource.Attribute{
					Name:        "fp_cert_ext_aia_ocsp",
					Description: `OCSP (Authority Information Access URI)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_connection_reuse",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template connection-reuse resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"connection",
				"reuse",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Connection Reuse Template Name`,
				},
				resource.Attribute{
					Name:        "keep_alive_conn",
					Description: `Keep a number of server connections open`,
				},
				resource.Attribute{
					Name:        "limit_per_server",
					Description: `Max Server Connections allowed (Connections per Server Port (default 1000))`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout in seconds. Multiple of 60 (default 2400)`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_csv",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template csv resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"csv",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "csv_name",
					Description: `Specify name of csv template`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "ipv6_enable",
					Description: `Support IPv6 IP ranges`,
				},
				resource.Attribute{
					Name:        "delim_num",
					Description: `enter a delimiter number, default 44 (“,”)`,
				},
				resource.Attribute{
					Name:        "multiple_fields",
					Description: ``,
				},
				resource.Attribute{
					Name:        "field",
					Description: `Field index number (Index of Field)`,
				},
				resource.Attribute{
					Name:        "csv_type",
					Description: `'ip-from’: Beginning address of IP range or subnet; 'ip-to-mask’: Ending address of IP range or Mask; 'continent’: Continent; 'country’: Country; 'state’: State or province; 'city’: City;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_dblb",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template dblb resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"dblb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `BDLB Template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "server_version",
					Description: `'MSSQL2008’: MSSQL server 2008 or 2008 R2; 'MSSQL2012’: MSSQL server 2012; 'MySQL’: MySQL server (any version)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_diameter",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template diameter resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"diameter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `diameter template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "terminate_on_cca_t",
					Description: `remove diameter session when receiving CCA-T message`,
				},
				resource.Attribute{
					Name:        "message_code_list",
					Description: ``,
				},
				resource.Attribute{
					Name:        "message_code",
					Description: `minimum: 1, maximum: 2147483647`,
				},
				resource.Attribute{
					Name:        "avp_list",
					Description: ``,
				},
				resource.Attribute{
					Name:        "int32",
					Description: `32 bits integer`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `mandatory avp`,
				},
				resource.Attribute{
					Name:        "avp",
					Description: `customize avps for cer to the server (avp number)`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `user sesison idle timeout (in minutes, default is 5)`,
				},
				resource.Attribute{
					Name:        "customize_cea",
					Description: `customizing cea response`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `product name avp`,
				},
				resource.Attribute{
					Name:        "dwr_up_retry",
					Description: `number of successful dwr health-check before declaring target up`,
				},
				resource.Attribute{
					Name:        "forward_unknown_session_id",
					Description: `Forward server message even it has unknown session id`,
				},
				resource.Attribute{
					Name:        "vendor_id",
					Description: `vendor-id avp (Vendor Id)`,
				},
				resource.Attribute{
					Name:        "session_age",
					Description: `user session age allowed (default 10), this is not idle-time (in minutes)`,
				},
				resource.Attribute{
					Name:        "load_balance_on_session_id",
					Description: `Load balance based on the session id`,
				},
				resource.Attribute{
					Name:        "dwr_time",
					Description: `dwr health-check timer interval (in 100 milli second unit, default is 100, 0 means unset this option)`,
				},
				resource.Attribute{
					Name:        "origin_realm",
					Description: `origin-realm name avp`,
				},
				resource.Attribute{
					Name:        "origin_host",
					Description: ``,
				},
				resource.Attribute{
					Name:        "origin_host_name",
					Description: `origin-host name avp`,
				},
				resource.Attribute{
					Name:        "multiple_origin_host",
					Description: `allowing multiple origin-host to a single server`,
				},
				resource.Attribute{
					Name:        "forward_to_latest_server",
					Description: `Forward client message to the latest server that sends message with the same session id`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_dns",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template DNS resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"dns",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `DNS Template Name`,
				},
				resource.Attribute{
					Name:        "enable_log",
					Description: `Use private key for authentication`,
				},
				resource.Attribute{
					Name:        "filter_response_rate",
					Description: `Maximum allowed request rate for the filter. This should match average traffic. (default 20 per two seconds)`,
				},
				resource.Attribute{
					Name:        "slip_rate",
					Description: `Every n’th response that would be rate-limited will be let through instead`,
				},
				resource.Attribute{
					Name:        "response_rate",
					Description: `Responses exceeding this rate within the window will be dropped (default 5 per second)`,
				},
				resource.Attribute{
					Name:        "window",
					Description: `Rate-Limiting Interval in Seconds (default is one)`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `'log-only’: Only log rate-limiting, do not actually rate limit. Requires enable-log configuration; 'rate-limit’: Rate-Limit based on configuration (Default); 'whitelist’: Whitelist, disable rate-limiting;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_dynamic_service",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template dynamic-service resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"dynamic",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Dynamic Service Template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "dns_server",
					Description: `blade parameter priority`,
				},
				resource.Attribute{
					Name:        "ipv4_dns_server",
					Description: `DNS Server IPv4 Address`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_external_service",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template external-service resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"external",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `External Service Template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `'skyfire-icap’: Skyfire ICAP service; 'url-filter’: URL filtering service;`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `'continue’: Continue; 'drop’: Drop; 'reset’: Reset;`,
				},
				resource.Attribute{
					Name:        "failure_action",
					Description: `'continue’: Continue; 'drop’: Drop; 'reset’: Reset;`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout value 1 - 200 in units of 200ms, default is 5 (default is 1000ms) (1 - 200 in units of 200ms, default is 5 (1000ms))`,
				},
				resource.Attribute{
					Name:        "request_header_forward_list",
					Description: ``,
				},
				resource.Attribute{
					Name:        "request_header_forward",
					Description: `Request header to be forwarded to external service (Header Name)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_fix",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template fix resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"fix",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `FIX Template Name`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `'init’: init only log; 'term’: termination only log; 'both’: both initial and termination log`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Create a Service Group comprising Servers (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "switching_type",
					Description: `'sender-comp-id’: Select service group based on SenderCompID; 'target-comp-id’: Select service group based on TargetComp`,
				},
				resource.Attribute{
					Name:        "insert_client_ip",
					Description: `Insert client ip to tag 11447`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_ftp",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template ftp resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"ftp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `FTP Template Name`,
				},
				resource.Attribute{
					Name:        "active_mode_port",
					Description: `Non-Standard FTP Active mode port`,
				},
				resource.Attribute{
					Name:        "active_mode_port_val",
					Description: `Non-Standard FTP Active mode port`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `End range of FTP Active mode port`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_http",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template http resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"http",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `HTTP Template Name`,
				},
				resource.Attribute{
					Name:        "keep_client_alive",
					Description: `Keep client alive`,
				},
				resource.Attribute{
					Name:        "req_hdr_wait_time",
					Description: `HTTP request header wait time before abort connection`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_http_policy",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template http-policy resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"http",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `http-policy template name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `name of cookie to match (Cookie Name)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_imap_pop3",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template imap-pop3 resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"imap",
				"pop3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `IMAP-POP3 Template Name`,
				},
				resource.Attribute{
					Name:        "logindisabled",
					Description: `Disable Login before STARTTLS.Works only for imap`,
				},
				resource.Attribute{
					Name:        "starttls",
					Description: `'disabled’: Disable STARTTLS; 'optional’: STARTTLS is optional requirement; 'enforced’: Must issue STARTTLS command before imap transaction;`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_logging",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template logging resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"logging",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Logging Template Name`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Specify a format string for web logging (format string(less than 250 characters) for web logging)`,
				},
				resource.Attribute{
					Name:        "local_logging",
					Description: `1 to enable local logging (1 to enable local logging, default 0)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_monitor",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template monitor resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"monitor",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "monitor_relation",
					Description: `'monitor-and’: Configures the monitors in current template to work with AND logic; 'monitor-or’: Configures the monitors in current template to work with OR logic;`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "dis_sequence",
					Description: `Sequence number (Specify the physical port number)`,
				},
				resource.Attribute{
					Name:        "diseth",
					Description: `Specify the physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "id2",
					Description: `Monitor template ID Number`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_mqtt",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template mqtt resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"mqtt",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `MQTT Template Name`,
				},
				resource.Attribute{
					Name:        "clientid_hash_persist",
					Description: `Immediate Removal after Transaction`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_policy",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template policy resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Policy template name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "use_destination_ip",
					Description: `Use destination IP to match the policy`,
				},
				resource.Attribute{
					Name:        "over_limit",
					Description: `Specify operation in case over limit`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Define timeout value of PBSLB dynamic entry (Timeout value (minute, default is 5))`,
				},
				resource.Attribute{
					Name:        "over_limit_lockup",
					Description: `Don’t accept any new connection for certain time (Lockup duration (minute))`,
				},
				resource.Attribute{
					Name:        "over_limit_reset",
					Description: `Reset the connection when it exceeds limit`,
				},
				resource.Attribute{
					Name:        "overlap",
					Description: `Use overlap mode for geo-location to do longest match`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_port",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template port resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `PORT Template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "slow_start",
					Description: `Slowly ramp up the connection number after port is up`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `Connection limit`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight (port weight)`,
				},
				resource.Attribute{
					Name:        "extended_stats",
					Description: `Enable extended statistics on real server port`,
				},
				resource.Attribute{
					Name:        "del_session_on_server_down",
					Description: `Delete session if the server/port goes down (either disabled/hm down)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_reqmod_icap",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template reqmod-icap resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"reqmod",
				"icap",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Reqmod ICAP Template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "min_payload_size",
					Description: `min-payload-size value 0 - 65535, default is 0`,
				},
				resource.Attribute{
					Name:        "allowed_http_methods",
					Description: `List of allowed HTTP methods. Default is "Allow All". (List of HTTP methods allowed (default “Allow All”))`,
				},
				resource.Attribute{
					Name:        "service_url",
					Description: `URL to send to ICAP server (Service URL Name)`,
				},
				resource.Attribute{
					Name:        "preview",
					Description: `Preview value 1 - 32768, default is 32768`,
				},
				resource.Attribute{
					Name:        "disable_http_server_reset",
					Description: `Don’t reset http server`,
				},
				resource.Attribute{
					Name:        "fail_close",
					Description: `When template sg is down mark vport down`,
				},
				resource.Attribute{
					Name:        "include_protocol_in_uri",
					Description: `Include protocol and port in HTTP URI`,
				},
				resource.Attribute{
					Name:        "log_only_allowed_method",
					Description: `Only log allowed HTTP method`,
				},
				resource.Attribute{
					Name:        "cylance",
					Description: `cylance external server`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_respmod_icap",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template respmod-icap resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"respmod",
				"icap",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Respmod ICAP Template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "min_payload_size",
					Description: `min-payload-size value 0 - 65535, default is 0`,
				},
				resource.Attribute{
					Name:        "preview",
					Description: `Preview value 1 - 32768, default is 32768`,
				},
				resource.Attribute{
					Name:        "disable_http_server_reset",
					Description: `Don’t reset http server`,
				},
				resource.Attribute{
					Name:        "fail_close",
					Description: `When template sg is down mark vport down`,
				},
				resource.Attribute{
					Name:        "include_protocol_in_uri",
					Description: `Include protocol and port in HTTP URI`,
				},
				resource.Attribute{
					Name:        "log_only_allowed_method",
					Description: `Only log allowed HTTP method`,
				},
				resource.Attribute{
					Name:        "cylance",
					Description: `cylance external server`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_server",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template server resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Server template name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "stats_data_action",
					Description: `'stats-data-enable’: Enable statistical data collection for real server; 'stats-data-disable’: Disable statistical data collection for real server;`,
				},
				resource.Attribute{
					Name:        "slow_start",
					Description: `Slowly ramp up the connection number after server is up`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight for the Real Servers (Connection Weight (default is 1))`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit",
					Description: `Configure bandwidth rate limit on real server (Bandwidth rate limit in Kbps)`,
				},
				resource.Attribute{
					Name:        "spoofing_cache",
					Description: `Servers under the template are spoofing cache`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `Connection limit`,
				},
				resource.Attribute{
					Name:        "resume",
					Description: `Resume accepting new connection after connection number drops below threshold (Connection resume threshold)`,
				},
				resource.Attribute{
					Name:        "max_dynamic_server",
					Description: `Maximum dynamic server number (Maximum dynamic server number (default is 255))`,
				},
				resource.Attribute{
					Name:        "rate_interval",
					Description: `'100ms’: Use 100 ms as sampling interval; 'second’: Use 1 second as sampling interval;`,
				},
				resource.Attribute{
					Name:        "min_ttl_ratio",
					Description: `Minimum TTL to DNS query interval ratio (Minimum TTL ratio (default is 2))`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_no_logging",
					Description: `Do not log bandwidth rate limit related state transitions`,
				},
				resource.Attribute{
					Name:        "dynamic_server_prefix",
					Description: `Prefix of dynamic server (Prefix of dynamic server (default is “DRS”))`,
				},
				resource.Attribute{
					Name:        "conn_limit_no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "extended_stats",
					Description: `Enable extended statistics on real server`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit_no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_duration",
					Description: `Duration in seconds the observed rate needs to honor`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_resume",
					Description: `Resume server selection after bandwidth drops below this threshold (in Kbps) (Bandwidth rate limit resume threshold (in Kbps))`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_acct",
					Description: `'to-server-only’: Only account for traffic sent to server; 'from-server-only’: Only account for traffic received from server; 'all’: Account for all traffic sent to and received from server;`,
				},
				resource.Attribute{
					Name:        "log_selection_failure",
					Description: `Enable real-time logging for server selection failure event`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit",
					Description: `Connection rate limit`,
				},
				resource.Attribute{
					Name:        "dns_query_interval",
					Description: `The interval to query DNS server for the hostname (DNS query interval (in minute, default is 10))`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_server_ssh",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template server-ssh resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"server",
				"ssh",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Server SSH Template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "forward_proxy_enable",
					Description: `Enable SSH forward proxy`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_server_ssl",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template server-ssl resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"server",
				"ssl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Server SSL Template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "sslilogging",
					Description: `'disable’: Disable all logging; 'all’: enable all logging(error, info);`,
				},
				resource.Attribute{
					Name:        "ocsp_stapling",
					Description: `Enable ocsp-stapling support`,
				},
				resource.Attribute{
					Name:        "ssli_logging",
					Description: `SSLi logging level, default is error logging only`,
				},
				resource.Attribute{
					Name:        "session_cache_size",
					Description: `Session Cache Size (Maximum cache size. Default value 0 (Session ID reuse disabled))`,
				},
				resource.Attribute{
					Name:        "handshake_logging_enable",
					Description: `Enable SSL handshake logging`,
				},
				resource.Attribute{
					Name:        "close_notify",
					Description: `Send close notification when terminate connection`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_sip",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template sip resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"sip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `SIP Template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "smp_call_id_rtp_session",
					Description: `Create the across cpu call-id rtp session`,
				},
				resource.Attribute{
					Name:        "client_keep_alive",
					Description: `Respond client keep-alive packet directly instead of forwarding to server`,
				},
				resource.Attribute{
					Name:        "alg_source_nat",
					Description: `Translate source IP to NAT IP in SIP message when source NAT is used`,
				},
				resource.Attribute{
					Name:        "alg_dest_nat",
					Description: `Translate VIP to real server IP in SIP message when destination NAT is used`,
				},
				resource.Attribute{
					Name:        "server_keep_alive",
					Description: `Send server keep-alive packet for every persist connection when enable conn-reuse`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `The interval of keep-alive packet for each persist connection (second)`,
				},
				resource.Attribute{
					Name:        "dialog_aware",
					Description: `Permit system processes dialog session`,
				},
				resource.Attribute{
					Name:        "failed_server_selection",
					Description: `Define action when select server fail`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Time in minutes`,
				},
				resource.Attribute{
					Name:        "drop_when_server_fail",
					Description: `Drop current SIP message when select server fail`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_smpp",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template smpp resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"smpp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `SMPP Template Name`,
				},
				resource.Attribute{
					Name:        "server_enquire_link",
					Description: `Send server ENQUIRE_LINK packet for every persist connection when enable conn-reuse`,
				},
				resource.Attribute{
					Name:        "server_selection_per_request",
					Description: `Force server selection on every SMPP request when enable conn-reuse`,
				},
				resource.Attribute{
					Name:        "client_enquire_link",
					Description: `Respond client ENQUIRE_LINK packet directly instead of forwarding to server`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "server_enquire_link_val",
					Description: `Set interval of keep-alive packet for each persistent connection (second, default is 30)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Configure the user to bind (The name used to bind)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Configure the password used to bind`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_smtp",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template smtp resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"smtp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `SMTP Template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "server_domain",
					Description: `Config the domain of the email servers (Server’s domain, default is “mail-server-domain”)`,
				},
				resource.Attribute{
					Name:        "client_starttls_type",
					Description: `'optional’: STARTTLS is optional requirement; 'enforced’: Must issue STARTTLS command before mail transaction`,
				},
				resource.Attribute{
					Name:        "disable_type",
					Description: `'optional’: STARTTLS is optional requirement; 'enforced’: Must issue STARTTLS command before mail transaction`,
				},
				resource.Attribute{
					Name:        "server_starttls_type",
					Description: `'optional’: STARTTLS is optional requirement; 'enforced’: Must issue STARTTLS command before mail transaction;`,
				},
				resource.Attribute{
					Name:        "service_ready_msg",
					Description: `Set SMTP service ready message (SMTP service ready message, default is “ESMTP mail service ready”)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_snmp",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template snmp resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"snmp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "priv_proto",
					Description: `'aes’: AES; 'des’: DES;`,
				},
				resource.Attribute{
					Name:        "context_name",
					Description: `Specify context name`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Specify interval, default is 3 (Interval, unit: second, default is 3)`,
				},
				resource.Attribute{
					Name:        "security_level",
					Description: `'no-auth’: No authentication; 'auth-no-priv’: Authentication, but no privacy; 'auth-priv’: Authentication and privacy;`,
				},
				resource.Attribute{
					Name:        "community",
					Description: `Specify community for version 2c (Community name)`,
				},
				resource.Attribute{
					Name:        "auth_proto",
					Description: `'sha’: SHA; 'md5’: MD5;`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `'v1’: Version 1; 'v2c’: Version 2c; 'v3’: Version 3;`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Specify Interface ID`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Specify port, default is 161 (Port Number, default is 161)`,
				},
				resource.Attribute{
					Name:        "snmp_name",
					Description: `Specify name of snmp template`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_ssli",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template ssli resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"ssli",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `SSLi Template Name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `'http’: HTTP service; 'xmpp’: XMPP service; 'smtp’: SMTP service; 'pop’: POP service;`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_tcp",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template tcp resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"tcp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `TCP Template Name`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Idle Timeout value (Interval of 60 seconds), default 120 seconds (idle timeout in second, default 120)`,
				},
				resource.Attribute{
					Name:        "insert_client_ip",
					Description: `Insert client ip into TCP option`,
				},
				resource.Attribute{
					Name:        "reset_fwd",
					Description: `send reset to server if error happens`,
				},
				resource.Attribute{
					Name:        "reset_rev",
					Description: `send reset to client if error happens`,
				},
				resource.Attribute{
					Name:        "del_session_on_server_down",
					Description: `Delete session if the server/port goes down (either disabled/hm down)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_tcp_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template tcp-proxy resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"tcp",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `TCP Proxy Template Name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `QOS level (number)`,
				},
				resource.Attribute{
					Name:        "init_cwnd",
					Description: `The initial congestion control window size (packets), default is 10 (init-cwnd in packets, default 10)`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Idle Timeout (Interval of 60 seconds), default is 600 (idle timeout in second, default 600)`,
				},
				resource.Attribute{
					Name:        "fin_timeout",
					Description: `FIN timeout (sec), default is disabled (number)`,
				},
				resource.Attribute{
					Name:        "half_open_idle_timeout",
					Description: `TCP Half Open Idle Timeout (sec), default is off (number)`,
				},
				resource.Attribute{
					Name:        "reno",
					Description: `Enable Reno Congestion Control Algorithm`,
				},
				resource.Attribute{
					Name:        "early_retransmit",
					Description: `Configure the Early-Retransmit Algorithm (RFC 5827) (Early-Retransmit is disabled by default)`,
				},
				resource.Attribute{
					Name:        "server_down_action",
					Description: `'FIN’: FIN Connection; 'RST’: Reset Connection;`,
				},
				resource.Attribute{
					Name:        "timewait",
					Description: `Timewait Threshold (sec), default 5 (number)`,
				},
				resource.Attribute{
					Name:        "min_rto",
					Description: `The minmum retransmission timeout, default is 200ms (number)`,
				},
				resource.Attribute{
					Name:        "dynamic_buffer_allocation",
					Description: `Optimally adjust the transmit and receive buffer sizes of TCP proxy while keeping their sum constant`,
				},
				resource.Attribute{
					Name:        "limited_slowstart",
					Description: `RFC 3742 Limited Slow-Start for TCP with Large Congestion Windows (number)`,
				},
				resource.Attribute{
					Name:        "disable_sack",
					Description: `disable Selective Ack Option`,
				},
				resource.Attribute{
					Name:        "disable_window_scale",
					Description: `disable TCP Window-Scale Option`,
				},
				resource.Attribute{
					Name:        "alive_if_active",
					Description: `keep connection alive if active traffic`,
				},
				resource.Attribute{
					Name:        "mss",
					Description: `Responding MSS to use if client MSS is large, default is off (number)`,
				},
				resource.Attribute{
					Name:        "keepalive_interval",
					Description: `Interval between keepalive probes (sec), default is off (number (seconds))`,
				},
				resource.Attribute{
					Name:        "retransmit_retries",
					Description: `Number of Retries for Retransmit, default is 5`,
				},
				resource.Attribute{
					Name:        "insert_client_ip",
					Description: `Insert client ip into TCP option`,
				},
				resource.Attribute{
					Name:        "transmit_buffer",
					Description: `TCP Transmit Buffer (default 200k) (number default 200000 bytes)`,
				},
				resource.Attribute{
					Name:        "nagle",
					Description: `Enable Nagle Algorithm`,
				},
				resource.Attribute{
					Name:        "initial_window_size",
					Description: `Set the initial window size, default is off (number)`,
				},
				resource.Attribute{
					Name:        "keepalive_probes",
					Description: `Number of keepalive probes sent, default is off`,
				},
				resource.Attribute{
					Name:        "psh_flag_optimization",
					Description: `Enable Optimized PSH Flag Use`,
				},
				resource.Attribute{
					Name:        "ack_aggressiveness",
					Description: `'low’: Delayed ACK; 'medium’: Delayed ACK, with ACK on each packet with PUSH flag; 'high’: ACK on each packet;`,
				},
				resource.Attribute{
					Name:        "backend_wscale",
					Description: `The TCP window scale used for the server side, default is off (number)`,
				},
				resource.Attribute{
					Name:        "reset_rev",
					Description: `send reset to client if error happens`,
				},
				resource.Attribute{
					Name:        "maxburst",
					Description: `The max packet count sent per transmission event (number)`,
				},
				resource.Attribute{
					Name:        "receive_buffer",
					Description: `TCP Receive Buffer (default 200k) (number default 200000 bytes)`,
				},
				resource.Attribute{
					Name:        "del_session_on_server_down",
					Description: `Delete session if the server/port goes down (either disabled/hm down)`,
				},
				resource.Attribute{
					Name:        "reassembly_timeout",
					Description: `The reassembly timeout, default is 30sec (number)`,
				},
				resource.Attribute{
					Name:        "reset_fwd",
					Description: `send reset to server if error happens`,
				},
				resource.Attribute{
					Name:        "disable_tcp_timestamps",
					Description: `disable TCP Timestamps Option`,
				},
				resource.Attribute{
					Name:        "syn_retries",
					Description: `SYN Retry Numbers, default is 5`,
				},
				resource.Attribute{
					Name:        "force_delete_timeout",
					Description: `The maximum time that a session can stay in the system before being deleted, default is off (number (second))`,
				},
				resource.Attribute{
					Name:        "reassembly_limit",
					Description: `The reassembly queuing limit, default is 25 segments (number)`,
				},
				resource.Attribute{
					Name:        "invalid_rate_limit",
					Description: `Invalid Packet Response Rate Limit (ms), default is 500 (number default 500 challenges)`,
				},
				resource.Attribute{
					Name:        "disable_abc",
					Description: `Appropriate Byte Counting RFC 3465 Disabled, default is enabled (Appropriate Byte Counting (ABC) is enabled by default)`,
				},
				resource.Attribute{
					Name:        "half_close_idle_timeout",
					Description: `TCP Half Close Idle Timeout (sec), default is off (cmd is deprecated, use fin-timeout instead) (number)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_udp",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB template udp resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"udp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "qos",
					Description: `QOS level (number)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Fast UDP Template Name`,
				},
				resource.Attribute{
					Name:        "stateless_conn_timeout",
					Description: `Stateless Current Connection Timeout value (5 - 120 seconds) (idle timeout in second, default 120)`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Idle Timeout value (Interval of 60 seconds), default 120 seconds (idle timeout in second, default 120)`,
				},
				resource.Attribute{
					Name:        "re_select_if_server_down",
					Description: `re-select another server if service port is down`,
				},
				resource.Attribute{
					Name:        "immediate",
					Description: `Immediate Removal after Transaction`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_virtual_port",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template virtual-port resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"virtual",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Virtual port template name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "reset_unknown_conn",
					Description: `Send reset back if receives TCP packet without SYN or RST flag and it does not belong to any existing connections`,
				},
				resource.Attribute{
					Name:        "ignore_tcp_msl",
					Description: `reclaim TCP resource immediately without MSL`,
				},
				resource.Attribute{
					Name:        "rate",
					Description: `Source IP and port rate limit (Packet rate limit)`,
				},
				resource.Attribute{
					Name:        "snat_msl",
					Description: `Source NAT MSL (Source NAT MSL value (seconds))`,
				},
				resource.Attribute{
					Name:        "allow_syn_otherflags",
					Description: `Allow initial SYN packet with other flags`,
				},
				resource.Attribute{
					Name:        "aflow",
					Description: `Use aFlow to eliminate the traffic surge`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `Connection limit`,
				},
				resource.Attribute{
					Name:        "drop_unknown_conn",
					Description: `Drop conection if receives TCP packet without SYN or RST flag and it does not belong to any existing connections`,
				},
				resource.Attribute{
					Name:        "reset_l7_on_failover",
					Description: `Send reset to L7 client and server connection upon a failover`,
				},
				resource.Attribute{
					Name:        "pkt_rate_type",
					Description: `'src-ip-port’: Source IP and port rate limit; 'src-port’: Source port rate limit;`,
				},
				resource.Attribute{
					Name:        "rate_interval",
					Description: `'100ms’: Use 100 ms as sampling interval; 'second’: Use 1 second as sampling interval;`,
				},
				resource.Attribute{
					Name:        "snat_port_preserve",
					Description: `Source NAT Port Preservation`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit_reset",
					Description: `Send client reset when connection rate over limit`,
				},
				resource.Attribute{
					Name:        "when_rr_enable",
					Description: `Only do rate limit if CPU RR triggered`,
				},
				resource.Attribute{
					Name:        "non_syn_initiation",
					Description: `Allow initial TCP packet to be non-SYN`,
				},
				resource.Attribute{
					Name:        "conn_limit_reset",
					Description: `Send client reset when connection over limit`,
				},
				resource.Attribute{
					Name:        "dscp",
					Description: `Differentiated Services Code Point (DSCP to Real Server IP Mapping Value)`,
				},
				resource.Attribute{
					Name:        "pkt_rate_limit_reset",
					Description: `send client-side reset (reset after packet limit)`,
				},
				resource.Attribute{
					Name:        "conn_limit_no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit_no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "log_options",
					Description: `'no-logging’: Do not log over limit event; 'no-repeat-logging’: log once for over limit event. Default is log once per minute;`,
				},
				resource.Attribute{
					Name:        "allow_vip_to_rport_mapping",
					Description: `Allow mapping of VIP to real port`,
				},
				resource.Attribute{
					Name:        "pkt_rate_interval",
					Description: `'100ms’: Source IP and port rate limit per 100ms; 'second’: Source IP and port rate limit per second (default);`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit",
					Description: `Connection rate limit`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_template_virtual_server",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder slb template virtual-server resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"template",
				"virtual",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Virtual server template name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `Connection limit`,
				},
				resource.Attribute{
					Name:        "conn_limit_reset",
					Description: `Send client reset when connection over limit`,
				},
				resource.Attribute{
					Name:        "icmpv6_rate_limit",
					Description: `ICMPv6 rate limit (Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit)`,
				},
				resource.Attribute{
					Name:        "subnet_gratuitous_arp",
					Description: `Send gratuitous ARP for every IP in the subnet virtual server`,
				},
				resource.Attribute{
					Name:        "tcp_stack_tfo_backoff_time",
					Description: `The time tcp stack will wait before allowing new fast-open requests after security condition, default 600 seconds (number)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_transparent_acl_template",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB transparent acl template resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"transparent",
				"acl",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specify template name`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_slb_transparent_tcp_template",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder SLB transparent tcp template resource for A10`,
			Description:      ``,
			Keywords: []string{
				"slb",
				"transparent",
				"tcp",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specify template name`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_virtual_server",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder virtual server resource for A10`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of virtual server`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `Configure auto NAT for the vport`,
				},
				resource.Attribute{
					Name:        "port_number",
					Description: `Port number`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `protocol`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Service group name`,
				},
				resource.Attribute{
					Name:        "snat_on_vip",
					Description: `Enable source NAT traffic against VIP`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_vrrp_common",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder vrrp common resource for A10`,
			Description:      ``,
			Keywords: []string{
				"vrrp",
				"common",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "set_id",
					Description: `Set-ID for HA configuration (Set id from 1 to 15)`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `Unique ID for each VRRP-A box (Device-id number)`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `'enable'= enable vrrp-a; 'disable'= disable vrrp-a;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_vrrp_peer_group",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder vrrp peer group resource for A10`,
			Description:      ``,
			Keywords: []string{
				"vrrp",
				"peer",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_peer_address",
					Description: `IP address of peer`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_vrrp_session_sync",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder vrrp session sync resource for A10`,
			Description:      ``,
			Keywords: []string{
				"vrrp",
				"session",
				"sync",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `'enable'= enable vrrp-a session sync (default); 'disable'= disable vrrp-a session sync; "`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_vrrp_vrid",
			Category:         "Resources",
			ShortDescription: `Provides details about vthunder vrrp vrid resource for A10`,
			Description:      ``,
			Keywords: []string{
				"vrrp",
				"vrid",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrid_val",
					Description: `"Specify ha VRRP-A vrid"`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address field`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `blade parameter priority`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"vthunder_configure_sync":                    0,
		"vthunder_dns_primary":                       1,
		"vthunder_ethernet":                          2,
		"vthunder_glm":                               3,
		"vthunder_harmony_controller_profile":        4,
		"vthunder_import":                            5,
		"vthunder_interface_ethernet":                6,
		"vthunder_interface_ethernet_bfd":            7,
		"vthunder_interface_ethernet_ipv6":           8,
		"vthunder_interface_ethernet_lldp":           9,
		"vthunder_interface_ethernet_trunk_group":    10,
		"vthunder_interface_management":              11,
		"vthunder_interface_ve":                      12,
		"vthunder_interface_ve_bfd":                  13,
		"vthunder_interface_ve_ip":                   14,
		"vthunder_interface_ve_ipv6":                 15,
		"vthunder_ip_address":                        16,
		"vthunder_ip_dns_primary":                    17,
		"vthunder_ip_dns_secondary":                  18,
		"vthunder_ip_dns_suffix":                     19,
		"vthunder_ip_frag":                           20,
		"vthunder_ip_icmp":                           21,
		"vthunder_ip_nat_alg_pptp":                   22,
		"vthunder_ip_nat_global":                     23,
		"vthunder_ip_nat_icmp":                       24,
		"vthunder_ip_prefix_list":                    25,
		"vthunder_ip_reroute":                        26,
		"vthunder_ip_route_static_bfd":               27,
		"vthunder_ip_tcp":                            28,
		"vthunder_ipv6_frag":                         29,
		"vthunder_ipv6_icmpv6":                       30,
		"vthunder_ipv6_nat_icmpv6":                   31,
		"vthunder_overlay_tunnel_options":            32,
		"vthunder_overlay_tunnel_vtep":               33,
		"vthunder_partition":                         34,
		"vthunder_reboot":                            35,
		"vthunder_rib_route":                         36,
		"vthunder_server":                            37,
		"vthunder_service_group":                     38,
		"vthunder_slb_aflow":                         39,
		"vthunder_slb_common":                        40,
		"vthunder_slb_common_conn_rate_limit_src_ip": 41,
		"vthunder_slb_dns":                           42,
		"vthunder_slb_dns_cache":                     43,
		"vthunder_slb_dns_response_rate_limiting":    44,
		"vthunder_slb_fast_http_proxy":               45,
		"vthunder_slb_fix":                           46,
		"vthunder_slb_ftp_ctl":                       47,
		"vthunder_slb_ftp_data":                      48,
		"vthunder_slb_ftp_proxy":                     49,
		"vthunder_slb_generic_proxy":                 50,
		"vthunder_slb_health_gateway":                51,
		"vthunder_slb_health_stat":                   52,
		"vthunder_slb_http2":                         53,
		"vthunder_slb_http_proxy":                    54,
		"vthunder_slb_hw_compress":                   55,
		"vthunder_slb_icap":                          56,
		"vthunder_slb_icap_http":                     57,
		"vthunder_slb_imapproxy":                     58,
		"vthunder_slb_l4":                            59,
		"vthunder_slb_l7session":                     60,
		"vthunder_slb_mlb":                           61,
		"vthunder_slb_mssql":                         62,
		"vthunder_slb_mysql":                         63,
		"vthunder_slb_passthrough":                   64,
		"vthunder_slb_perf":                          65,
		"vthunder_slb_persist":                       66,
		"vthunder_slb_player_id_global":              67,
		"vthunder_slb_pop3_proxy":                    68,
		"vthunder_slb_proxy":                         69,
		"vthunder_slb_rate_limit_log":                70,
		"vthunder_slb_rc_cache_global":               71,
		"vthunder_slb_resource_usage":                72,
		"vthunder_slb_sip":                           73,
		"vthunder_slb_smpp":                          74,
		"vthunder_slb_smtp":                          75,
		"vthunder_slb_spdy_proxy":                    76,
		"vthunder_slb_sport_rate_limit":              77,
		"vthunder_slb_ssl_cert_revoke":               78,
		"vthunder_slb_ssl_expire_check":              79,
		"vthunder_slb_ssl_forward_proxy":             80,
		"vthunder_slb_svm_source_nat":                81,
		"vthunder_slb_switch":                        82,
		"vthunder_slb_template_cache":                83,
		"vthunder_slb_template_cipher":               84,
		"vthunder_slb_template_client_ssh":           85,
		"vthunder_slb_template_client_ssl":           86,
		"vthunder_slb_template_connection_reuse":     87,
		"vthunder_slb_template_csv":                  88,
		"vthunder_slb_template_dblb":                 89,
		"vthunder_slb_template_diameter":             90,
		"vthunder_slb_template_dns":                  91,
		"vthunder_slb_template_dynamic_service":      92,
		"vthunder_slb_template_external_service":     93,
		"vthunder_slb_template_fix":                  94,
		"vthunder_slb_template_ftp":                  95,
		"vthunder_slb_template_http":                 96,
		"vthunder_slb_template_http_policy":          97,
		"vthunder_slb_template_imap_pop3":            98,
		"vthunder_slb_template_logging":              99,
		"vthunder_slb_template_monitor":              100,
		"vthunder_slb_template_mqtt":                 101,
		"vthunder_slb_template_policy":               102,
		"vthunder_slb_template_port":                 103,
		"vthunder_slb_template_reqmod_icap":          104,
		"vthunder_slb_template_respmod_icap":         105,
		"vthunder_slb_template_server":               106,
		"vthunder_slb_template_server_ssh":           107,
		"vthunder_slb_template_server_ssl":           108,
		"vthunder_slb_template_sip":                  109,
		"vthunder_slb_template_smpp":                 110,
		"vthunder_slb_template_smtp":                 111,
		"vthunder_slb_template_snmp":                 112,
		"vthunder_slb_template_ssli":                 113,
		"vthunder_slb_template_tcp":                  114,
		"vthunder_slb_template_tcp_proxy":            115,
		"vthunder_slb_template_udp":                  116,
		"vthunder_slb_template_virtual_port":         117,
		"vthunder_slb_template_virtual_server":       118,
		"vthunder_slb_transparent_acl_template":      119,
		"vthunder_slb_transparent_tcp_template":      120,
		"vthunder_virtual_server":                    121,
		"vthunder_vrrp_common":                       122,
		"vthunder_vrrp_peer_group":                   123,
		"vthunder_vrrp_session_sync":                 124,
		"vthunder_vrrp_vrid":                         125,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
