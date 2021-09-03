package vthunder

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_access_list_extended",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder access list extended resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "extended",
					Description: `Configure Extended Access List`,
				},
				resource.Attribute{
					Name:        "extd",
					Description: `IP extended access list`,
				},
				resource.Attribute{
					Name:        "extd-seq-num",
					Description: `Sequence number`,
				},
				resource.Attribute{
					Name:        "extd-remark",
					Description: `Access list entry comment (Notes for this ACL)`,
				},
				resource.Attribute{
					Name:        "extd-action",
					Description: `'deny': Deny; 'permit': Permit; 'l3-vlan-fwd-disable': Disable L3 forwarding between VLANs;`,
				},
				resource.Attribute{
					Name:        "icmp",
					Description: `Internet Control Message Protocol`,
				},
				resource.Attribute{
					Name:        "tcp",
					Description: `protocol TCP`,
				},
				resource.Attribute{
					Name:        "udp",
					Description: `protocol UDP`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Any Internet Protocol`,
				},
				resource.Attribute{
					Name:        "service-obj-group",
					Description: `Service object group (Source object group name)`,
				},
				resource.Attribute{
					Name:        "icmp-type",
					Description: `ICMP type number`,
				},
				resource.Attribute{
					Name:        "any-type",
					Description: `Any ICMP type`,
				},
				resource.Attribute{
					Name:        "special-type",
					Description: `'echo-reply': Type 0, echo reply; 'echo-request': Type 8, echo request; 'info-reply': Type 16, information reply; 'info-request': Type 15, information request; 'mask-reply': Type 18, address mask reply; 'mask-request': Type 17, address mask request; 'parameter-problem': Type 12, parameter problem; 'redirect': Type 5, redirect message; 'source-quench': Type 4, source quench; 'time-exceeded': Type 11, time exceeded; 'timestamp': Type 13, timestamp; 'timestamp-reply': Type 14, timestamp reply; 'dest-unreachable': Type 3, destination unreachable;`,
				},
				resource.Attribute{
					Name:        "any-code",
					Description: `Any ICMP code`,
				},
				resource.Attribute{
					Name:        "icmp-code",
					Description: `ICMP code number`,
				},
				resource.Attribute{
					Name:        "special-code",
					Description: `'frag-required': Code 4, fragmentation required; 'host-unreachable': Code 1, destination host unreachable; 'network-unreachable': Code 0, destination network unreachable; 'port-unreachable': Code 3, destination port unreachable; 'proto-unreachable': Code 2, destination protocol unreachable; 'route-failed': Code 5, source route failed;`,
				},
				resource.Attribute{
					Name:        "src-any",
					Description: `Any source host`,
				},
				resource.Attribute{
					Name:        "src-host",
					Description: `A single source host (Host address)`,
				},
				resource.Attribute{
					Name:        "src-subnet",
					Description: `Source Address`,
				},
				resource.Attribute{
					Name:        "src-mask",
					Description: `Source Mask 0=apply 255=ignore`,
				},
				resource.Attribute{
					Name:        "src-object-group",
					Description: `Network object group (Source network object group name)`,
				},
				resource.Attribute{
					Name:        "src-eq",
					Description: `Match only packets on a given source port (port number)`,
				},
				resource.Attribute{
					Name:        "src-gt",
					Description: `Match only packets with a greater port number`,
				},
				resource.Attribute{
					Name:        "src-lt",
					Description: `Match only packets with a lower port number`,
				},
				resource.Attribute{
					Name:        "src-range",
					Description: `match only packets in the range of port numbers (Starting Port Number)`,
				},
				resource.Attribute{
					Name:        "src-port-end",
					Description: `Ending Port Number`,
				},
				resource.Attribute{
					Name:        "dst-any",
					Description: `Any destination host`,
				},
				resource.Attribute{
					Name:        "dst-host",
					Description: `A single destination host (Host address)`,
				},
				resource.Attribute{
					Name:        "dst-subnet",
					Description: `Destination Address`,
				},
				resource.Attribute{
					Name:        "dst-mask",
					Description: `Destination Mask 0=apply 255=ignore`,
				},
				resource.Attribute{
					Name:        "dst-object-group",
					Description: `Destination network object group name`,
				},
				resource.Attribute{
					Name:        "dst-eq",
					Description: `Match only packets on a given destination port (port number)`,
				},
				resource.Attribute{
					Name:        "dst-gt",
					Description: `Match only packets with a greater port number`,
				},
				resource.Attribute{
					Name:        "dst-lt",
					Description: `Match only packets with a lesser port number`,
				},
				resource.Attribute{
					Name:        "dst-range",
					Description: `Match only packets in the range of port numbers (Starting Destination Port Number)`,
				},
				resource.Attribute{
					Name:        "dst-port-end",
					Description: `Edning Destination Port Number`,
				},
				resource.Attribute{
					Name:        "fragments",
					Description: `IP fragments`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `VLAN ID`,
				},
				resource.Attribute{
					Name:        "ethernet",
					Description: `Ethernet interface (Port number)`,
				},
				resource.Attribute{
					Name:        "trunk",
					Description: `Ethernet trunk (trunk number)`,
				},
				resource.Attribute{
					Name:        "dscp",
					Description: `DSCP`,
				},
				resource.Attribute{
					Name:        "established",
					Description: `TCP established`,
				},
				resource.Attribute{
					Name:        "acl-log",
					Description: `Log matches against this entry`,
				},
				resource.Attribute{
					Name:        "transparent-session-only",
					Description: `Only log transparent sessions`,
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
			Type:             "vthunder_thunder_access_list_standard",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder access list standard resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "standard",
					Description: `Configure Standard Access List`,
				},
				resource.Attribute{
					Name:        "std",
					Description: `IP standard access list`,
				},
				resource.Attribute{
					Name:        "seq-num",
					Description: `Sequence number`,
				},
				resource.Attribute{
					Name:        "std-remark",
					Description: `Access list entry comment (Notes for this ACL)`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `'deny': Deny; 'permit': Permit; 'l3-vlan-fwd-disable': Disable L3 forwarding between VLANs;`,
				},
				resource.Attribute{
					Name:        "any",
					Description: `Any source host`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `A single source host (Host address)`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `Source Address`,
				},
				resource.Attribute{
					Name:        "rev-subnet-mask",
					Description: `Network Mask 0=apply 255=ignore`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `Log matches against this entry`,
				},
				resource.Attribute{
					Name:        "transparent-session-only",
					Description: `Only log transparent sessions`,
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
			Type:             "vthunder_thunder_active_partition",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder active partition resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "active-partition",
					Description: `Change active partition`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Shared partition`,
				},
				resource.Attribute{
					Name:        "curr_part_name",
					Description: `current active patition`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_bgp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder bgp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bgp",
					Description: `Border Gateway Protocol (BGP)`,
				},
				resource.Attribute{
					Name:        "extended-asn-cap",
					Description: `Enable the router to send 4-octet ASN capabilities`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Enable the nexthop tracking functionality`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `Configure nexthop trigger delay time interval (Nexthop Trigger Delay time interval between 1 and 100)`,
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
			Type:             "vthunder_thunder_configure_sync",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder configure sync resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `Specify the destination ip address to sync`,
				},
				resource.Attribute{
					Name:        "all_partitions",
					Description: `All partition configurations`,
				},
				resource.Attribute{
					Name:        "auto_authentication",
					Description: `Authenticate with local username and password`,
				},
				resource.Attribute{
					Name:        "partition_name",
					Description: `Partition name`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `Use private key for authentication`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Shared partition`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `‘running’: Sync local running to peer’s running configuration; ‘all’: Sync local running to peer’s running configuration, and local startup to peer’s startup configuration;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_dns_primary",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder dns primary resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ethernet",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ethernet resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_file_ssl_cert",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder file ssl cert resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ssl-cert",
					Description: `ssl certificate file information and management commands`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `ssl certificate local file name`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `ssl certificate file size in byte`,
				},
				resource.Attribute{
					Name:        "file-handle",
					Description: `full path of the uploaded file`,
				},
				resource.Attribute{
					Name:        "certificate-type",
					Description: `'pem': pem; 'der': der; 'pfx': pfx; 'p7b': p7b;`,
				},
				resource.Attribute{
					Name:        "pfx-password",
					Description: `The password for certificate file (pfx type only)`,
				},
				resource.Attribute{
					Name:        "pfx-password-export",
					Description: `The password for exported certificate file (pfx type only)`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;`,
				},
				resource.Attribute{
					Name:        "dst-file",
					Description: `destination file name for copy and rename action`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "file_local_path",
					Description: `local directory path for pki ssl-cert to upload`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_active_rule_set",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw active rule set resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Rule set name`,
				},
				resource.Attribute{
					Name:        "override_nat_aging",
					Description: `Override NAT idle-timeout`,
				},
				resource.Attribute{
					Name:        "session_aging",
					Description: `Session Aging Template`,
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
			Type:             "vthunder_thunder_fw_alg_dns",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw alg dns resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default_port_disable",
					Description: `‘default-port-disable’: Disable DNS ALG default port 53;`,
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
			Type:             "vthunder_thunder_fw_alg_ftp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw alg ftp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default_port_disable",
					Description: `‘default-port-disable’: Disable FTP ALG default port 21;`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘client-port-request’: PORT Requests From Client; ‘client-eprt-request’: EPRT Requests From Client; ‘server-pasv-reply’: PASV Replies From Server; ‘server-epsv-reply’: EPSV Replies From Server; ‘port-retransmits’: PORT Retransmits; ‘pasv-retransmits’: PASV Retransmits; ‘smp-app-type-mismatch’: SMP App Type Mismatch; ‘retransmit-sanity-check-failure’: Retransmit Sanity Check Failure; ‘smp-conn-alloc-failure’: SMP Helper Conn Alloc Failure; ‘port-helper-created’: PORT Helper Created; ‘pasv-helper-created’: PASV Helper Created; ‘port-helper-acquire-in-del-q’: PORT Helper Acquire In Del Queue; ‘port-helper-acquire-already-used’: PORT Helper Acquire Already Used; ‘pasv-helper-acquire-in-del-q’: PASV Helper Acquire In Del Queue; ‘pasv-helper-acquire-already-used’: PASV Helper Acquire Already Used; ‘port-helper-freed-used’: PORT Helper Freed Used; ‘port-helper-freed-unused’: PORT Helper Freed Unused; ‘pasv-helper-freed-used’: PASV Helper Freed Used; ‘pasv-helper-freed-unused’: PASV Helper Freed Unused;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_alg_icmp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw alg icmp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "disable",
					Description: `‘disable’: Disable ICMP ALG which allows ICMP errors to pass the firewall;`,
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
			Type:             "vthunder_thunder_fw_alg_pptp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw alg pptp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default_port_disable",
					Description: `‘default-port-disable’: Disable PPTP ALG default port 1723;`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘calls-established’: Calls Established; ‘call-req-pns-call-id-mismatch’: Call ID Mismatch on Call Request; ‘call-reply-pns-call-id-mismatch’: Call ID Mismatch on Call Reply; ‘gre-session-created’: GRE Session Created; ‘gre-session-freed’: GRE Session Freed; ‘call-req-retransmit’: Call Request Retransmit; ‘call-req-new’: Call Request New; ‘call-req-ext-alloc-failure’: Call Request Ext Alloc Failure; ‘call-reply-call-id-unknown’: Call Reply Unknown Client Call ID; ‘call-reply-retransmit’: Call Reply Retransmit; ‘call-reply-ext-ext-alloc-failure’: Call Request Ext Alloc Failure; ‘smp-app-type-mismatch’: SMP App Type Mismatch; ‘smp-client-call-id-mismatch’: SMP Client Call ID Mismatch; ‘smp-sessions-created’: SMP Session Created; ‘smp-sessions-freed’: SMP Session Freed; ‘smp-alloc-failure’: SMP Session Alloc Failure; ‘gre-conn-creation-failure’: GRE Conn Alloc Failure; ‘gre-conn-ext-creation-failure’: GRE Conn Ext Alloc Failure; ‘gre-no-fwd-route’: GRE No Fwd Route; ‘gre-no-rev-route’: GRE No Rev Route; ‘gre-no-control-conn’: GRE No Control Conn; ‘gre-conn-already-exists’: GRE Conn Already Exists; ‘gre-free-no-ext’: GRE Free No Ext; ‘gre-free-no-smp’: GRE Free No SMP; ‘gre-free-smp-app-type-mismatch’: GRE Free SMP App Type Mismatch; ‘control-freed’: Control Session Freed; ‘control-free-no-ext’: Control Free No Ext; ‘control-free-no-smp’: Control Free No SMP; ‘control-free-smp-app-type-mismatch’: Control Free SMP App Type Mismatch;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_alg_rtsp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw alg rtsp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default_port_disable",
					Description: `‘default-port-disable’: Disable RTSP ALG default port 554;`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘transport-inserted’: Transport Created; ‘transport-freed’: Transport Freed; ‘transport-alloc-failure’: Transport Alloc Failure; ‘data-session-created’: Data Session Created; ‘data-session-freed’: Data Session Freed; ‘ext-creation-failure’: Extension Creation Failure; ‘transport-add-to-ext’: Transport Added to Extension; ‘transport-removed-from-ext’: Transport Removed from Extension; ‘transport-too-many’: Too Many Transports for Control Conn; ‘transport-already-in-ext’: Transport Already in Extension; ‘transport-exists’: Transport Already Exists; ‘transport-link-ext-failure-control’: Transport Link to Extension Failure Control; ‘transport-link-ext-data’: Transport Link to Extension Data; ‘transport-link-ext-failure-data’: Transport Link to Extension Failure Data; ‘transport-inserted-shadow’: Transport Inserted Shadow; ‘transport-creation-race’: Transport Create Race; ‘transport-alloc-failure-shadow’: Transport Alloc Failure Shadow; ‘transport-put-in-del-q’: Transport Put in Delete Queue; ‘transport-freed-shadow’: Transport Freed Shadow; ‘transport-acquired-from-control’: Transport Acquired Control; ‘transport-found-from-prev-control’: Transport Found From Prev Control; ‘transport-acquire-failure-from-control’: Transport Acquire Failure Control; ‘transport-released-from-control’: Transport Released Control; ‘transport-double-release-from-control’: Transport Double Release Control; ‘transport-acquired-from-data’: Transport Acquired Data; ‘transport-acquire-failure-from-data’: Transport Acquire Failure Data; ‘transport-released-from-data’: Transport Released Data; ‘transport-double-release-from-data’: Transport Double Release Data; ‘transport-retry-lookup-on-data-free’: Transport Retry Lookup Data; ‘transport-not-found-on-data-free’: Transport Not Found Data; ‘data-session-created-shadow’: Data Session Created Shadow; ‘data-session-freed-shadow’: Data Session Freed Shadow; ‘ha-control-ext-creation-failure’: HA Control Extension Creation Failure; ‘ha-control-session-created’: HA Control Session Created; ‘ha-data-session-created’: HA Data Session Created;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_alg_sip",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw alg sip resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default_port_disable",
					Description: `‘default-port-disable’: Disable SIP ALG default port 5060;`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘stat-request’: Request Received; ‘stat-response’: Response Received; ‘method-register’: Method REGISTER; ‘method-invite’: Method INVITE; ‘method-ack’: Method ACK; ‘method-cancel’: Method CANCEL; ‘method-bye’: Method BYE; ‘method-port-config’: Method OPTIONS; ‘method-prack’: Method PRACK; ‘method-subscribe’: Method SUBSCRIBE; ‘method-notify’: Method NOTIFY; ‘method-publish’: Method PUBLISH; ‘method-info’: Method INFO; ‘method-refer’: Method REFER; ‘method-message’: Method MESSAGE; ‘method-update’: Method UPDATE; ‘method-unknown’: Method Unknown; ‘parse-error’: Message Parse Error; ‘keep-alive’: Keep Alive; ‘contact-error’: Contact Process Error; ‘sdp-error’: SDP Process Error; ‘rtp-port-no-op’: RTP Port No Op; ‘rtp-rtcp-port-success’: RTP RTCP Port Success; ‘rtp-port-failure’: RTP Port Failure; ‘rtcp-port-failure’: RTCP Port Failure; ‘contact-port-no-op’: Contact Port No Op; ‘contact-port-success’: Contact Port Success; ‘contact-port-failure’: Contact Port Failure; ‘contact-new’: Contact Alloc; ‘contact-alloc-failure’: Contact Alloc Failure; ‘contact-eim’: Contact EIM; ‘contact-eim-set’: Contact EIM Set; ‘rtp-new’: RTP Alloc; ‘rtp-alloc-failure’: RTP Alloc Failure; ‘rtp-eim’: RTP EIM; ‘helper-found’: SMP Helper Conn Found; ‘helper-created’: SMP Helper Conn Created; ‘helper-deleted’: SMP Helper Conn Already Deleted; ‘helper-freed’: SMP Helper Conn Freed; ‘helper-failure’: SMP Helper Failure;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_alg_tftp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw alg tftp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default_port_disable",
					Description: `‘default-port-disable’: Disable TFTP ALG default port 69;`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘session-created’: TFTP Client Sessions Created; ‘helper-created’: TFTP Helper Sessions created; ‘helper-freed’: TFTP Helper Sessions freed; ‘helper-freed-used’: TFTP Helper Sessions freed used; ‘helper-freed-unused’: TFTP Helper Sessions freed unused; ‘helper-already-used’: TFTP Helper Session already used; ‘helper-in-rml’: TFTP Helper Session in Remove List;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_app",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw app resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘dummy’: Entry for a10countergen;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_apply_changes",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw apply changes resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "forced",
					Description: `Force recompile rule-set`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_clear_session_filter",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw clear session filter resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `‘disable’: Disable clear L4 session filter for fw (Default: disabled); ‘enable’: Enable clear L4 session filter for fw;`,
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
			Type:             "vthunder_thunder_fw_full_cone_session",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw full cone session resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_global",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw global resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alg_processing",
					Description: `‘honor-rule-set’: Honors firewall rule-sets; ‘override-rule-set’: Override firewall rule-sets;`,
				},
				resource.Attribute{
					Name:        "disable_ip_fw_sessions",
					Description: `disable create sessions for non TCP/UDP/ICMP`,
				},
				resource.Attribute{
					Name:        "extended_matching",
					Description: `‘disable’: Disable extended matching;`,
				},
				resource.Attribute{
					Name:        "listen_on_port_timeout",
					Description: `STUN timeout (default: 2 minutes)`,
				},
				resource.Attribute{
					Name:        "natip_ddos_protection",
					Description: `‘enable’: Enable; ‘disable’: Disable;`,
				},
				resource.Attribute{
					Name:        "permit_default_action",
					Description: `‘forward’: Forward; ‘next-service-mode’: Service to be applied chosen based on configuration;`,
				},
				resource.Attribute{
					Name:        "respond_to_user_mac",
					Description: `Use the user’s source MAC for the next hop rather than the routing table (default: off)`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘tcp_fullcone_created’: TCP Full-cone Created; ‘tcp_fullcone_freed’: TCP Full-cone Freed; ‘udp_fullcone_created’: UDP Full-cone Created; ‘udp_fullcone_freed’: UDP Full-cone Freed; ‘fullcone_creation_failure’: Full-Cone Creation Failure; ‘data_session_created’: Data Session Created; ‘data_session_freed’: Data Session Freed; ‘fullcone_in_del_q’: Full-cone session found in delete queue; ‘fullcone_retry_lookup’: Full-cone session retry look-up; ‘fullcone_not_found’: Full-cone session not found; ‘fullcone_overflow_eim’: Full-cone Session EIM Overflow; ‘fullcone_overflow_eif’: Full-cone Session EIF Overflow; ‘udp_fullcone_created_shadow’: Total UDP Full-cone sessions created; ‘tcp_fullcone_created_shadow’: Total TCP Full-cone sessions created; ‘udp_fullcone_freed_shadow’: Total UDP Full-cone sessions freed; ‘tcp_fullcone_freed_shadow’: Total TCP Full-cone sessions freed; ‘fullcone_created’: Total Full-cone sessions created; ‘fullcone_freed’: Total Full-cone sessions freed; ‘fullcone_ext_too_many’: Fullcone Extension Too Many; ‘fullcone_ext_mem_allocated’: Fullcone Extension Memory Allocated; ‘fullcone_ext_mem_alloc_failure’: Fullcone Extension Memory Allocate Failure; ‘fullcone_ext_mem_alloc_init_faulure’: Fullcone Extension Initialization Failure; ‘fullcone_ext_mem_freed’: Fullcone Extension Memory Freed; ‘fullcone_ext_added’: Fullcone Extension Added; ‘ha_fullcone_failure’: HA Full-cone Session Failure; ‘data_session_created_shadow’: Total Data Sessions Created; ‘data_session_freed_shadow’: Total Data Sessions Freed; ‘active_fullcone_session’: Total Active Full-cone sessions; ‘limit-entry-failure’: Limit Entry Creation Failure; ‘limit-entry-allocated’: Limit Entry Allocated; ‘limit-entry-mem-freed’: Limit Entry Freed; ‘limit-entry-created’: Limit Entry Created; ‘limit-entry-not-in-bucket’: Limit Entry Not in Bucket; ‘limit-entry-marked-deleted’: Limit Entry Marked Deleted; ‘invalid-lid-drop’: Invalid Lid Drop; ‘src-session-limit-exceeded’: Source Prefix Session Limit Exceeded; ‘limit-exceeded’: Per Second Limit Exceeded; ‘limit-entry-per-cpu-mem-allocated’: Limit Entry Memory Allocated; ‘limit-entry-per-cpu-mem-allocation-failed’: Limit Entry Memory Allocation Failed; ‘limit-entry-per-cpu-mem-freed’: Limit Entry Memory Freed; ‘alg_default_port_disable’: Total ALG packets matching Default Port Disable; ‘no_fwd_route’: No Forward Route; ‘no_rev_route’: No Reverse Route; ‘no_fwd_l2_dst’: No Forward Mac Entry; ‘no_rev_l2_dst’: No Reverse Mac Entry; ‘urpf_pkt_drop’: URPF check packet drop; ‘fwd_ingress_packets_tcp’: Forward Ingress Packets TCP; ‘fwd_egress_packets_tcp’: Forward Egress Packets TCP; ‘rev_ingress_packets_tcp’: Reverse Ingress Packets TCP; ‘rev_egress_packets_tcp’: Reverse Egress Packets TCP; ‘fwd_ingress_bytes_tcp’: Forward Ingress Bytes TCP; ‘fwd_egress_bytes_tcp’: Forward Egress Bytes TCP; ‘rev_ingress_bytes_tcp’: Reverse Ingress Bytes TCP; ‘rev_egress_bytes_tcp’: Reverse Egress Bytes TCP; ‘fwd_ingress_packets_udp’: Forward Ingress Packets UDP; ‘fwd_egress_packets_udp’: Forward Egress Packets UDP; ‘rev_ingress_packets_udp’: Reverse Ingress Packets UDP; ‘rev_egress_packets_udp’: Reverse Egress Packets UDP; ‘fwd_ingress_bytes_udp’: Forward Ingress Bytes UDP; ‘fwd_egress_bytes_udp’: Forward Egress Bytes UDP; ‘rev_ingress_bytes_udp’: Reverse Ingress Bytes UDP; ‘rev_egress_bytes_udp’: Reverse Egress Bytes UDP; ‘fwd_ingress_packets_icmp’: Forward Ingress Packets ICMP; ‘fwd_egress_packets_icmp’: Forward Egress Packets ICMP; ‘rev_ingress_packets_icmp’: Reverse Ingress Packets ICMP; ‘rev_egress_packets_icmp’: Reverse Egress Packets ICMP; ‘fwd_ingress_bytes_icmp’: Forward Ingress Bytes ICMP; ‘fwd_egress_bytes_icmp’: Forward Egress Bytes ICMP; ‘rev_ingress_bytes_icmp’: Reverse Ingress Bytes ICMP; ‘rev_egress_bytes_icmp’: Reverse Egress Bytes ICMP; ‘fwd_ingress_packets_others’: Forward Ingress Packets OTHERS; ‘fwd_egress_packets_others’: Forward Egress Packets OTHERS; ‘rev_ingress_packets_others’: Reverse Ingress Packets OTHERS; ‘rev_egress_packets_others’: Reverse Egress Packets OTHERS; ‘fwd_ingress_bytes_others’: Forward Ingress Bytes OTHERS; ‘fwd_egress_bytes_others’: Forward Egress Bytes OTHERS; ‘rev_ingress_bytes_others’: Reverse Ingress Bytes OTHERS; ‘rev_egress_bytes_others’: Reverse Egress Bytes OTHERS; ‘fwd_ingress_pkt_size_range1’: Forward Ingress Packet size between 0 and 200; ‘fwd_ingress_pkt_size_range2’: Forward Ingress Packet size between 201 and 800; ‘fwd_ingress_pkt_size_range3’: Forward Ingress Packet size between 801 and 1550; ‘fwd_ingress_pkt_size_range4’: Forward Ingress Packet size between 1551 and 9000; ‘fwd_egress_pkt_size_range1’: Forward Egress Packet size between 0 and 200; ‘fwd_egress_pkt_size_range2’: Forward Egress Packet size between 201 and 800; ‘fwd_egress_pkt_size_range3’: Forward Egress Packet size between 801 and 1550; ‘fwd_egress_pkt_size_range4’: Forward Egress Packet size between 1551 and 9000; ‘rev_ingress_pkt_size_range1’: Reverse Ingress Packet size between 0 and 200; ‘rev_ingress_pkt_size_range2’: Reverse Ingress Packet size between 201 and 800; ‘rev_ingress_pkt_size_range3’: Reverse Ingress Packet size between 801 and 1550; ‘rev_ingress_pkt_size_range4’: Reverse Ingress Packet size between 1551 and 9000; ‘rev_egress_pkt_size_range1’: Reverse Egress Packet size between 0 and 200; ‘rev_egress_pkt_size_range2’: Reverse Egress Packet size between 201 and 800; ‘rev_egress_pkt_size_range3’: Reverse Egress Packet size between 801 and 1550; ‘rev_egress_pkt_size_range4’: Reverse Egress Packet size between 1551 and 9000;`,
				},
				resource.Attribute{
					Name:        "disable_application_category",
					Description: `‘aaa’: Protocol/application used for AAA (Authentification, Authorization and Accounting) purposes.; ‘adult-content’: Adult content.; ‘advertising’: Advertising networks and applications.; ‘analytics-and-statistics’: user-analytics and statistics.; ‘anonymizers-and-proxies’: Traffic-anonymization protocol/application.; ‘audio-chat’: Protocol/application used for Audio Chat.; ‘basic’: Protocols required for basic classification, e.g., ARP, HTTP; ‘blog’: Blogging platform.; ‘cdn’: Protocol/application used for Content-Delivery Networks.; ‘chat’: Protocol/application used for Text Chat.; ‘classified-ads’: Protocol/application used for Classified ads.; ‘cloud-based-services’: SaaS and/or PaaS cloud based services.; ‘database’: Database-specific protocols.; ‘email’: Native email protocol.; ‘enterprise’: Protocol/application used in an enterprise network.; ‘file-management’: Protocol/application designed specifically for file management and exchange, e.g., Dropbox, SMB; ‘file-transfer’: Protocol that offers file transferring as a functionality as a secondary feature. e.g., Skype, Whatsapp; ‘forum’: Online forum.; ‘gaming’: Protocol/application used by games.; ‘instant-messaging-and-multimedia-conferencing’: Protocol/application used for Instant messaging or multiconferencing.; ‘internet-of-things’: Internet Of Things protocol/application.; ‘mobile’: Mobile-specific protocol/application.; ‘multimedia-streaming’: Protocol/application used for multimedia streaming.; ‘networking’: Protocol used for (inter) networking purpose.; ‘news-portal’: Protocol/application used for News Portals.; ‘peer-to-peer’: Protocol/application used for Peer-to-peer purposes.; ‘remote-access’: Protocol/application used for remote access.; ‘scada’: SCADA (Supervisory control and data acquisition) protocols, all generations.; ‘social-networks’: Social networking application.; ‘software-update’: Auto-update protocol.; ‘standards-based’: Protocol issued from standardized bodies such as IETF, ITU, IEEE, ETSI, OIF.; ‘video-chat’: Protocol/application used for Video Chat.; ‘voip’: Application used for Voice over IP.; ‘vpn-tunnels’: Protocol/application used for VPN or tunneling purposes.; ‘web’: Application based on HTTP/HTTPS.; ‘web-e-commerce’: Protocol/application used for E-commerce websites.; ‘web-search-engines’: Protocol/application used for Web search portals.; ‘web-websites’: Protocol/application used for Company Websites.; ‘webmails’: Web email application.; ‘web-ext-adult’: Web Extension Adult; ‘web-ext-auctions’: Web Extension Auctions; ‘web-ext-blogs’: Web Extension Blogs; ‘web-ext-business-and-economy’: Web Extension Business and Economy; ‘web-ext-cdns’: Web Extension CDNs; ‘web-ext-collaboration’: Web Extension Collaboration; ‘web-ext-computer-and-internet-info’: Web Extension Computer and Internet Info; ‘web-ext-computer-and-internet-security’: Web Extension Computer and Internet Security; ‘web-ext-dating’: Web Extension Dating; ‘web-ext-educational-institutions’: Web Extension Educational Institutions; ‘web-ext-entertainment-and-arts’: Web Extension Entertainment and Arts; ‘web-ext-fashion-and-beauty’: Web Extension Fashion and Beauty; ‘web-ext-file-share’: Web Extension File Share; ‘web-ext-financial-services’: Web Extension Financial Services; ‘web-ext-gambling’: Web Extension Gambling; ‘web-ext-games’: Web Extension Games; ‘web-ext-government’: Web Extension Government; ‘web-ext-health-and-medicine’: Web Extension Health and Medicine; ‘web-ext-individual-stock-advice-and-tools’: Web Extension Individual Stock Advice and Tools; ‘web-ext-internet-portals’: Web Extension Internet Portals; ‘web-ext-job-search’: Web Extension Job Search; ‘web-ext-local-information’: Web Extension Local Information; ‘web-ext-malware’: Web Extension Malware; ‘web-ext-motor-vehicles’: Web Extension Motor Vehicles; ‘web-ext-music’: Web Extension Music; ‘web-ext-news’: Web Extension News; ‘web-ext-p2p’: Web Extension P2P; ‘web-ext-parked-sites’: Web Extension Parked Sites; ‘web-ext-proxy-avoid-and-anonymizers’: Web Extension Proxy Avoid and Anonymizers; ‘web-ext-real-estate’: Web Extension Real Estate; ‘web-ext-reference-and-research’: Web Extension Reference and Research; ‘web-ext-search-engines’: Web Extension Search Engines; ‘web-ext-shopping’: Web Extension Shopping; ‘web-ext-social-network’: Web Extension Social Network; ‘web-ext-society’: Web Extension Society; ‘web-ext-software’: Web Extension Software; ‘web-ext-sports’: Web Extension Sports; ‘web-ext-streaming-media’: Web Extension Streaming Media; ‘web-ext-training-and-tools’: Web Extension Training and Tools; ‘web-ext-translation’: Web Extension Translation; ‘web-ext-travel’: Web Extension Travel; ‘web-ext-web-advertisements’: Web Extension Web Advertisements; ‘web-ext-web-based-email’: Web Extension Web based Email; ‘web-ext-web-hosting’: Web Extension Web Hosting; ‘web-ext-web-service’: Web Extension Web Service;`,
				},
				resource.Attribute{
					Name:        "disable_application_protocol",
					Description: `Disable specific application protocol`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_gtp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw gtp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gtp_value",
					Description: `‘enable’: Enable GTP Inspection;`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘create-session-request’: Create Session Request; ‘create-session-response’: Create Session Response; ‘path-management-message’: Path Management Message; ‘delete-session-request’: Delete Session Request; ‘delete-session-response’: Delete Session Response; ‘reserved-field-set-drop’: Reserved field set drop; ‘tunnel-id-flag-drop’: Tunnel ID Flag Incorrect; ‘message-filtering-drop’: Message Filtering Drop; ‘reserved-information-element-drop’: Resevered Information Element Field Drop; ‘mandatory-information-element-drop’: Mandatory Information Element Field Drop; ‘filter-list-drop’: APN IMSI Information Filtering Drop; ‘invalid-teid-drop’: Invalid TEID Drop; ‘out-of-state-drop’: Out Of State Drop; ‘message-length-drop’: Message Length Exceeded; ‘unsupported-message-type-v2’: GTP v2 message type is not supported; ‘fast-conn-setup’: Fast Conn Setup Attempt; ‘out-of-session-memory’: Out of Session Memory; ‘no-fwd-route’: No Forward Route; ‘no-rev-route’: NO Reverse Route; ‘invalid-key’: Invalid TEID Field; ‘create-session-request-retransmit’: Retransmitted Create Session Request; ‘delete-session-request-retransmit’: Retransmitted Delete Session Request; ‘response-cause-not-accepted-drop’: Response Cause indicates Request not Accepted; ‘invalid-imsi-len-drop’: Invalid IMSI Length Drop; ‘invalid-apn-len-drop’: Invalid APN Length Drop; ‘create-pdp-context-request-v1’: GTP v1 Create PDP Context Request; ‘create-pdp-context-response-v1’: GTP v1 Create PDP Context Response; ‘path-management-message-v1’: GTP v1 Path Management Message; ‘reserved-field-set-drop-v1’: GTP v1 Reserved field set drop; ‘message-filtering-drop-v1’: GTP v1 Message Filtering Drop; ‘reserved-information-element-drop-v1’: GTP v1 Reserved Information Element Field Drop; ‘mandatory-information-element-drop-v1’: GTP v1 Mandatory Information Element Field Drop; ‘filter-list-drop-v1’: GTP v1 APN IMSI Information Filtering Drop; ‘invalid-teid-drop-v1’: GTP v1 Invalid TEID Drop; ‘message-length-drop-v1’: GTP v1 Message Length Exceeded; ‘version-not-supported’: GTP version is not supported; ‘unsupported-message-type-v1’: GTP v1 message type is not supported; ‘delete-pdp-context-request-v1’: GTP v1 Delete Context PDP Request; ‘delete-pdp-context-response-v1’: GTP v1 Delete Context PDP Response; ‘create-pdp-context-request-v0’: GTP v0 Create PDP Context Request; ‘create-pdp-context-response-v0’: GTP v0 Create PDP Context Response; ‘delete-pdp-context-request-v0’: GTP v0 Delete Context PDP Request; ‘delete-pdp-context-response-v0’: GTP v0 Delete Context PDP Response; ‘path-management-message-v0’: GTP v0 Path Management Message; ‘message-filtering-drop-v0’: GTP v0 Message Filtering Drop; ‘unsupported-message-type-v0’: GTP v0 message type is not supported; ‘invalid-flow-label-drop-v0’: GTP v0 Invalid flow label drop; ‘invalid-tid-drop-v0’: GTP v0 Invalid tid drop; ‘message-length-drop-v0’: GTP v0 Message Length Exceeded; ‘mandatory-information-element-drop-v0’: GTP v0 Mandatory Information Element Field Drop; ‘filter-list-drop-v0’: GTP v0 APN IMSI Information Filtering Drop; ‘gtp-in-gtp-drop’: GTP in GTP Filtering Drop;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_gtp_in_gtp_filtering",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw gtp in gtp filtering resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gtp_in_gtp_value",
					Description: `‘disable’: Disable GTP in GTP filtering, (default:Enabled);`,
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
			Type:             "vthunder_thunder_fw_gtp_v0",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw gtp v0 resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gtpv0_value",
					Description: `‘enable’: Enable GTP v0 Inspection;`,
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
			Type:             "vthunder_thunder_fw_helper_sessions",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw helper sessions resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `helper-sessions idle-timeout time (Idle-timeout in minutes (default: 1 minute))`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `Limit number of helper-sessions (Limit helper-sessions number)`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `‘disable’: Disable helper-sessions;`,
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
			Type:             "vthunder_thunder_fw_limit_entry",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw limit entry resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_local_log",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw local log resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "local_logging",
					Description: `Enable local logging`,
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
			Type:             "vthunder_thunder_fw_logging",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw logging resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Logging Template Name`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘log_message_sent’: Log Packet Sent; ‘log_type_reset’: Log Event Type Reset; ‘log_type_deny’: Log Event Type Deny; ‘log_type_session_closed’: Log Event Type Session Close; ‘log_type_session_opened’: Log Event Type Session Open; ‘rule_not_logged’: Firewall Rule Not Logged; ‘log-dropped’: Log Packets Dropped; ‘tcp-session-created’: TCP Session Created; ‘tcp-session-deleted’: TCP Session Deleted; ‘udp-session-created’: UDP Session Created; ‘udp-session-deleted’: UDP Session Deleted; ‘icmp-session-deleted’: ICMP Session Deleted; ‘icmp-session-created’: ICMP Session Created; ‘icmpv6-session-deleted’: ICMPV6 Session Deleted; ‘icmpv6-session-created’: ICMPV6 Session Created; ‘other-session-deleted’: Other Session Deleted; ‘other-session-created’: Other Session Created; ‘http-request-logged’: HTTP Request Logged; ‘http-logging-invalid-format’: HTTP Logging Invalid Format Error; ‘dcmsg_permit’: Dcmsg Permit; ‘alg_override_permit’: Alg Override Permit; ‘template_error’: Template Error; ‘ipv4-frag-applied’: IPv4 Fragmentation Applied; ‘ipv4-frag-failed’: IPv4 Fragmentation Failed; ‘ipv6-frag-applied’: IPv6 Fragmentation Applied; ‘ipv6-frag-failed’: IPv6 Fragmentation Failed; ‘out-of-buffers’: Out of Buffers; ‘add-msg-failed’: Add Message to Buffer Failed; ‘tcp-logging-conn-established’: TCP Logging Conn Established; ‘tcp-logging-conn-create-failed’: TCP Logging Conn Create Failed; ‘tcp-logging-conn-dropped’: TCP Logging Conn Dropped; ‘log-message-too-long’: Log message too long; ‘http-out-of-order-dropped’: HTTP out-of-order dropped; ‘http-alloc-failed’: HTTP Request Info Allocation Failed; ‘sctp-session-created’: SCTP Session Created; ‘sctp-session-deleted’: SCTP Session Deleted; ‘log_type_sctp_inner_proto_filter’: Log Event Type SCTP Inner Proto Filter; ‘log_type_gtp_message_filtering’: Log Event Type GTP Message Filtering; ‘log_type_gtp_apn_filtering’: Log Event Type GTP Apn Filtering; ‘tcp-logging-port-allocated’: TCP Logging Port Allocated; ‘tcp-logging-port-freed’: TCP Logging Port Freed; ‘tcp-logging-port-allocation-failed’: TCP Logging Port Allocation Failed; ‘log_type_gtp_invalid_teid’: Log Event Type GTP Invalid TEID; ‘log_gtp_type_reserved_ie_present’: Log Event Type GTP Reserved Information Element Present; ‘log_type_gtp_mandatory_ie_missing’: Log Event Type GTP Mandatory Information Element Missing;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_radius_server",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw radius server resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accounting_interim_update",
					Description: `‘ignore’: Ignore (default); ‘append-entry’: Append the AVPs to existing entry; ‘replace-entry’: Replace the AVPs of existing entry;`,
				},
				resource.Attribute{
					Name:        "accounting_on",
					Description: `‘ignore’: Ignore (default); ‘delete-entries-using-attribute’: Delete entries matching attribute in RADIUS Table;`,
				},
				resource.Attribute{
					Name:        "accounting_start",
					Description: `‘ignore’: Ignore; ‘append-entry’: Append the AVPs to existing entry (default); ‘replace-entry’: Replace the AVPs of existing entry;`,
				},
				resource.Attribute{
					Name:        "accounting_stop",
					Description: `‘ignore’: Ignore; ‘delete-entry’: Delete the entry (default);`,
				},
				resource.Attribute{
					Name:        "attribute_name",
					Description: `‘msisdn’: Clear using MSISDN; ‘imei’: Clear using IMEI; ‘imsi’: Clear using IMSI;`,
				},
				resource.Attribute{
					Name:        "custom_attribute_name",
					Description: `Clear using customized attribute`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED secret string)`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `Configure the listen port of RADIUS server (Port number)`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `Configure shared secret`,
				},
				resource.Attribute{
					Name:        "secret_string",
					Description: `The RADIUS secret`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "vrid",
					Description: `Join a VRRP-A failover group`,
				},
				resource.Attribute{
					Name:        "ip_list_encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED secret string)`,
				},
				resource.Attribute{
					Name:        "ip_list_name",
					Description: `IP-list name`,
				},
				resource.Attribute{
					Name:        "ip_list_secret",
					Description: `Configure shared secret`,
				},
				resource.Attribute{
					Name:        "ip_list_secret_string",
					Description: `The RADIUS secret`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘msisdn-received’: MSISDN Received; ‘imei-received’: IMEI Received; ‘imsi-received’: IMSI Received; ‘custom-received’: Custom attribute Received; ‘radius-request-received’: RADIUS Request Received; ‘radius-request-dropped’: RADIUS Request Dropped (Malformed Packet); ‘request-bad-secret-dropped’: RADIUS Request Bad Secret Dropped; ‘request-no-key-vap-dropped’: RADIUS Request No Key Attribute Dropped; ‘request-malformed-dropped’: RADIUS Request Malformed Dropped; ‘request-ignored’: RADIUS Request Table Full Dropped; ‘radius-table-full’: RADIUS Request Dropped (Table Full); ‘secret-not-configured-dropped’: RADIUS Secret Not Configured Dropped; ‘ha-standby-dropped’: HA Standby Dropped; ‘ipv6-prefix-length-mismatch’: Framed IPV6 Prefix Length Mismatch; ‘invalid-key’: Radius Request has Invalid Key Field; ‘smp-mem-allocated’: RADIUS SMP Memory Allocated; ‘smp-mem-alloc-failed’: RADIUS SMP Memory Allocation Failed; ‘smp-mem-freed’: RADIUS SMP Memory Freed; ‘smp-created’: RADIUS SMP Created; ‘smp-in-rml’: RADIUS SMP in RML; ‘smp-deleted’: RADIUS SMP Deleted; ‘mem-allocated’: RADIUS Memory Allocated; ‘mem-alloc-failed’: RADIUS Memory Allocation Failed; ‘mem-freed’: RADIUS Memory Freed; ‘ha-sync-create-sent’: HA Record Sync Create Sent; ‘ha-sync-delete-sent’: HA Record Sync Delete Sent; ‘ha-sync-create-recv’: HA Record Sync Create Received; ‘ha-sync-delete-recv’: HA Record Sync Delete Received; ‘acct-on-filters-full’: RADIUS Acct On Request Ignored(Filters Full); ‘acct-on-dup-request’: Duplicate RADIUS Acct On Request; ‘ip-mismatch-delete’: Radius Entry IP Mismatch Delete; ‘ip-add-race-drop’: Radius Entry IP Add Race Drop; ‘ha-sync-no-key-vap-dropped’: HA Record Sync No key dropped; ‘inter-card-msg-fail-drop’: Inter-Card Message Fail Drop;`,
				},
				resource.Attribute{
					Name:        "attribute_value",
					Description: `‘inside-ipv6-prefix’: Framed IPv6 Prefix; ‘inside-ip’: Inside IP address; ‘inside-ipv6’: Inside IPv6 address; ‘imei’: International Mobile Equipment Identity (IMEI); ‘imsi’: International Mobile Subscriber Identity (IMSI); ‘msisdn’: Mobile Subscriber Integrated Services Digital Network-Number (MSISDN); ‘custom1’: Customized attribute 1; ‘custom2’: Customized attribute 2; ‘custom3’: Customized attribute 3;`,
				},
				resource.Attribute{
					Name:        "custom_number",
					Description: `RADIUS attribute number`,
				},
				resource.Attribute{
					Name:        "custom_vendor",
					Description: `RADIUS vendor attribute information (RADIUS vendor ID)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Customized attribute name`,
				},
				resource.Attribute{
					Name:        "number",
					Description: `RADIUS attribute number`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `‘32’: Prefix length 32; ‘48’: Prefix length 48; ‘64’: Prefix length 64; ‘80’: Prefix length 80; ‘96’: Prefix length 96; ‘112’: Prefix length 112;`,
				},
				resource.Attribute{
					Name:        "prefix_number",
					Description: `RADIUS attribute number`,
				},
				resource.Attribute{
					Name:        "prefix_vendor",
					Description: `RADIUS vendor attribute information (RADIUS vendor ID)`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `‘hexadecimal’: Type of attribute value is hexadecimal;`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `RADIUS vendor attribute information (RADIUS vendor ID)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_resource_usage",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw resource usage resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_server",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw server resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `‘enable’: enable; ‘disable’: disable;`,
				},
				resource.Attribute{
					Name:        "fqdn_name",
					Description: `Server hostname`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Health Check (Monitor Name)`,
				},
				resource.Attribute{
					Name:        "health_check_disable",
					Description: `Disable health check`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `IP Address`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Server Name`,
				},
				resource.Attribute{
					Name:        "resolve_as",
					Description: `‘resolve-to-ipv4’: Use A Query only to resolve FQDN; ‘resolve-to-ipv6’: Use AAAA Query only to resolve FQDN; ‘resolve-to-ipv4-and-ipv6’: Use A as well as AAAA Query to resolve FQDN;`,
				},
				resource.Attribute{
					Name:        "server_ipv6_addr",
					Description: `IPV6 address`,
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
					Name:        "port_number",
					Description: `Port Number`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `‘tcp’: TCP Port; ‘udp’: UDP Port;`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘curr-conn’: Current connections; ‘total-conn’: Total connections; ‘fwd-pkt’: Forward packets; ‘rev-pkt’: Reverse Packets; ‘peak-conn’: Peak connections;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_service_group",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw service group resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "health_check",
					Description: `Health Check (Monitor Name)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Member name`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `‘tcp’: TCP LB service; ‘udp’: UDP LB service;`,
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
					Name:        "counters1",
					Description: `‘all’: all; ‘curr_conn’: Current connections; ‘total_fwd_bytes’: Total forward bytes; ‘total_fwd_pkts’: Total forward packets; ‘total_rev_bytes’: Total reverse bytes; ‘total_rev_pkts’: Total reverse packets; ‘total_conn’: Total connections; ‘total_rev_pkts_inspected’: Total reverse packets inspected; ‘total_rev_pkts_inspected_status_code_2xx’: Total reverse packets inspected status code 2xx; ‘total_rev_pkts_inspected_status_code_non_5xx’: Total reverse packets inspected status code non 5xx; ‘curr_req’: Current requests; ‘total_req’: Total requests; ‘total_req_succ’: Total requests success; ‘peak_conn’: Peak connections; ‘response_time’: Response time; ‘fastest_rsp_time’: Fastest response time; ‘slowest_rsp_time’: Slowest response time;`,
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
			Type:             "vthunder_thunder_fw_status",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw status resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_system_status",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw system status resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_tap_monitor",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw tap monitor resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `‘enable’: Enable tap monitor mode; ‘disable’: Disable tap monitor mode (Default:Disable);`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "tap_eth",
					Description: `Ethernet interface number`,
				},
				resource.Attribute{
					Name:        "tap_vlan",
					Description: `Vlan number`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_tcp_mss_clamp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw tcp mss clamp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "min",
					Description: `Specify the min value allowed for the TCP MSS (Specify the min value allowed for the TCP MSS (default: ((576 - 60 - 60))))`,
				},
				resource.Attribute{
					Name:        "mss_clamp_type",
					Description: `‘fixed’: Specify a fixed max value for the TCP MSS; ‘subtract’: Specify the value to subtract from the TCP MSS;`,
				},
				resource.Attribute{
					Name:        "mss_subtract",
					Description: `Specify the value to subtract from the TCP MSS (default: not configured)`,
				},
				resource.Attribute{
					Name:        "mss_value",
					Description: `The max value allowed for the TCP MSS (default: not configured)}`,
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
			Type:             "vthunder_thunder_fw_tcp_reset_on_error",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw tcp reset on error resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enable",
					Description: `Enable send TCP reset on error`,
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
			Type:             "vthunder_thunder_fw_tcp_rst_close_immediate",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw tcp rst close immediate resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `‘enable’: Enable TCP RST close immediate (default); ‘disable’: Disable TCP RST close immediate;`,
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
			Type:             "vthunder_thunder_fw_tcp_window_check",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw tcp window check resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `‘enable’: Enable TCP window check (default); ‘disable’: Disable TCP window check;`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘outside-window’: packet dropped for outside of tcp window;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_template_logging",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw template logging resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "facility",
					Description: `‘kernel’: 0: Kernel; ‘user’: 1: User-level; ‘mail’: 2: Mail; ‘daemon’: 3: System daemons; ‘security-authorization’: 4: Security/authorization; ‘syslog’: 5: Syslog internal; ‘line-printer’: 6: Line printer; ‘news’: 7: Network news; ‘uucp’: 8: UUCP subsystem; ‘cron’: 9: Time-related; ‘security-authorization-private’: 10: Private security/authorization; ‘ftp’: 11: FTP; ‘ntp’: 12: NTP; ‘audit’: 13: Audit; ‘alert’: 14: Alert; ‘clock’: 15: Clock-related; ‘local0’: 16: Local use 0; ‘local1’: 17: Local use 1; ‘local2’: 18: Local use 2; ‘local3’: 19: Local use 3; ‘local4’: 20: Local use 4; ‘local5’: 21: Local use 5; ‘local6’: 22: Local use 6; ‘local7’: 23: Local use 7;`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `‘ascii’: A10 Text logging format (ASCII); ‘cef’: Common Event Format for logging (default);`,
				},
				resource.Attribute{
					Name:        "include_dest_fqdn",
					Description: `Include destination FQDN string`,
				},
				resource.Attribute{
					Name:        "merged_style",
					Description: `Merge creation and deletion of session logs to one`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Logging Template Name`,
				},
				resource.Attribute{
					Name:        "resolution",
					Description: `‘seconds’: Logging timestamp resolution in seconds (default); ‘10-milliseconds’: Logging timestamp resolution in 10s of milli-seconds;`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Bind a Service Group to the logging template (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `‘emergency’: 0: Emergency; ‘alert’: 1: Alert; ‘critical’: 2: Critical; ‘error’: 3: Error; ‘warning’: 4: Warning; ‘notice’: 5: Notice; ‘informational’: 6: Informational; ‘debug’: 7: Debug;`,
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
					Name:        "ip",
					Description: `Specify source IP address`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `Specify source IPv6 address`,
				},
				resource.Attribute{
					Name:        "framed_ipv6_prefix",
					Description: `Include radius attributes for the prefix`,
				},
				resource.Attribute{
					Name:        "insert_if_not_existing",
					Description: `Configure what string is to be inserted for custom RADIUS attributes`,
				},
				resource.Attribute{
					Name:        "no_quote",
					Description: `No quotation marks for RADIUS attributes in logs`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `‘32’: Prefix length 32; ‘48’: Prefix length 48; ‘64’: Prefix length 64; ‘80’: Prefix length 80; ‘96’: Prefix length 96; ‘112’: Prefix length 112;`,
				},
				resource.Attribute{
					Name:        "zero_in_custom_attr",
					Description: `Insert 0000 for standard and custom attributes in log string`,
				},
				resource.Attribute{
					Name:        "attr",
					Description: `‘imei’: Include IMEI; ‘imsi’: Include IMSI; ‘msisdn’: Include MSISDN; ‘custom1’: Customized attribute 1; ‘custom2’: Customized attribute 2; ‘custom3’: Customized attribute 3;`,
				},
				resource.Attribute{
					Name:        "attr_event",
					Description: `‘http-requests’: Include in HTTP request logs; ‘sessions’: Include in session logs;`,
				},
				resource.Attribute{
					Name:        "disable_sequence_check",
					Description: `Disable http packet sequence check and don’t drop out of order packets`,
				},
				resource.Attribute{
					Name:        "include_all_headers",
					Description: `Include all configured headers despite of absence in HTTP request`,
				},
				resource.Attribute{
					Name:        "log_every_http_request",
					Description: `Log every HTTP request in an HTTP 1.1 session (Default: Log the first HTTP request in a session)`,
				},
				resource.Attribute{
					Name:        "max_url_len",
					Description: `Max length of URL log (Max URL length (Default: 100 char))`,
				},
				resource.Attribute{
					Name:        "include_byte_count",
					Description: `Include the byte count of HTTP Request/Response in FW session deletion logs`,
				},
				resource.Attribute{
					Name:        "file_extension",
					Description: `HTTP file extension`,
				},
				resource.Attribute{
					Name:        "l4_session_info",
					Description: `Log the L4 session information of the HTTP request`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `Log the HTTP Request Method`,
				},
				resource.Attribute{
					Name:        "request_number",
					Description: `HTTP Request Number`,
				},
				resource.Attribute{
					Name:        "custom_header_name",
					Description: `Header name`,
				},
				resource.Attribute{
					Name:        "custom_max_length",
					Description: `Max length for a HTTP request log (Max header length (Default: 100 char))`,
				},
				resource.Attribute{
					Name:        "http_header",
					Description: `‘cookie’: Log HTTP Cookie Header; ‘referer’: Log HTTP Referer Header; ‘user-agent’: Log HTTP User-Agent Header; ‘header1’: Log HTTP Header 1; ‘header2’: Log HTTP Header 2; ‘header3’: Log HTTP Header 3;`,
				},
				resource.Attribute{
					Name:        "max_length",
					Description: `Max length for a HTTP request log (Max header length (Default: 100 char))`,
				},
				resource.Attribute{
					Name:        "http_requests",
					Description: `‘host’: Log the HTTP Host Header; ‘url’: Log the HTTP Request URL;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_top_k_rules",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw top k rules resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_fw_urpf",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw urpf resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `‘disabled’: Disable urpf (Default: disabled); ‘strict’: Perform Strict Check; ‘loose’: Perform Loose Check;`,
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
			Type:             "vthunder_thunder_fw_vrid",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder fw vrid resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "vrid",
					Description: `Vrrp group (VRRP-A vrid)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_glm",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder glm resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allocate_bandwidth",
					Description: `Enter the requested bandwidth in Mbps for Capacity Pool`,
				},
				resource.Attribute{
					Name:        "appliance_name",
					Description: `Helpful identifier for this appliance`,
				},
				resource.Attribute{
					Name:        "enable_requests",
					Description: `Enable connection to the GLM`,
				},
				resource.Attribute{
					Name:        "enterprise",
					Description: `Enter the ELM hostname or IP`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Enter the desired ping interval in hours`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Proxy server port`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `License entitlement token`,
				},
				resource.Attribute{
					Name:        "use_mgmt_port",
					Description: `Use management port to connect to GLM`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "license_request",
					Description: `Do a single GLM license retrieval`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED secret string)`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Proxy server hostname or IP address`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for proxy authentication`,
				},
				resource.Attribute{
					Name:        "secret_string",
					Description: `password value`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for proxy authentication`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_glm_send",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder glm send resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "license_request",
					Description: `Do a single GLM license retrieval`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_harmony_controller_profile",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder harmony-controller profile resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_hostname",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder hostname resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `This System’s Network Name`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_import",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder import resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aflex",
					Description: `aFleX Script Source File`,
				},
				resource.Attribute{
					Name:        "auth_jwks",
					Description: `JSON web key`,
				},
				resource.Attribute{
					Name:        "auth_portal",
					Description: `Portal file for http authentication`,
				},
				resource.Attribute{
					Name:        "auth_portal_image",
					Description: `Image file for default portal`,
				},
				resource.Attribute{
					Name:        "bw_list",
					Description: `Black white List File`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `CA Cert File(enter bulk when import an archive file)`,
				},
				resource.Attribute{
					Name:        "certificate_type",
					Description: `‘pem’: pem; ‘der’: der; ‘pfx’: pfx; ‘p7b’: p7b;`,
				},
				resource.Attribute{
					Name:        "class_list",
					Description: `Class List File`,
				},
				resource.Attribute{
					Name:        "class_list_convert",
					Description: `Convert Class List File to A10 format`,
				},
				resource.Attribute{
					Name:        "class_list_type",
					Description: `‘ac’: ac; ‘ipv4’: ipv4; ‘ipv6’: ipv6; ‘string’: string; ‘string-case-insensitive’: string-case-insensitive;`,
				},
				resource.Attribute{
					Name:        "dnssec_dnskey",
					Description: `DNSSEC DNSKEY(KSK) file for child zone`,
				},
				resource.Attribute{
					Name:        "dnssec_ds",
					Description: `DNSSEC DS file for child zone`,
				},
				resource.Attribute{
					Name:        "file_inspection_bw_list",
					Description: `Black white List File`,
				},
				resource.Attribute{
					Name:        "geo_location",
					Description: `Geo-location CSV File`,
				},
				resource.Attribute{
					Name:        "glm_cert",
					Description: `GLM certificate`,
				},
				resource.Attribute{
					Name:        "glm_license",
					Description: `License File`,
				},
				resource.Attribute{
					Name:        "ip_map_list",
					Description: `IP Map List File`,
				},
				resource.Attribute{
					Name:        "local_uri_file",
					Description: `Local URI files for http response`,
				},
				resource.Attribute{
					Name:        "lw_4o6",
					Description: `LW-4over6 Binding Table File`,
				},
				resource.Attribute{
					Name:        "overwrite",
					Description: `Overwrite existing file`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `password for the remote site`,
				},
				resource.Attribute{
					Name:        "pfx_password",
					Description: `The password for certificate file (pfx type only)`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `WAF policy File`,
				},
				resource.Attribute{
					Name:        "remote_file",
					Description: `Profile name for remote url`,
				},
				resource.Attribute{
					Name:        "ssl_cert",
					Description: `SSL Cert File(enter bulk when import an archive file)`,
				},
				resource.Attribute{
					Name:        "ssl_cert_key",
					Description: `‘bulk’: import an archive file;`,
				},
				resource.Attribute{
					Name:        "ssl_crl",
					Description: `SSL Crl File`,
				},
				resource.Attribute{
					Name:        "ssl_key",
					Description: `SSL Key File(enter bulk when import an archive file)`,
				},
				resource.Attribute{
					Name:        "store_name",
					Description: `Import store name`,
				},
				resource.Attribute{
					Name:        "terminal",
					Description: `terminal vi`,
				},
				resource.Attribute{
					Name:        "thales_kmdata",
					Description: `Thales Kmdata files`,
				},
				resource.Attribute{
					Name:        "thales_secworld",
					Description: `Thales security world files`,
				},
				resource.Attribute{
					Name:        "usb_license",
					Description: `USB License File`,
				},
				resource.Attribute{
					Name:        "use_mgmt_port",
					Description: `Use management port as source port`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "web_category_license",
					Description: `License file to enable web-category feature`,
				},
				resource.Attribute{
					Name:        "wsdl",
					Description: `Web Service Definition Language File`,
				},
				resource.Attribute{
					Name:        "xml_schema",
					Description: `XML-Schema File`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Describe the Program Function briefly`,
				},
				resource.Attribute{
					Name:        "externalfilename",
					Description: `Specify the Program Name`,
				},
				resource.Attribute{
					Name:        "postfilename",
					Description: `Specify the File Name`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Create an import store profile`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Delete an import store profile`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `profile name to store remote url`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device (Device ID)`,
				},
				resource.Attribute{
					Name:        "saml_idp_name",
					Description: `Metadata name`,
				},
				resource.Attribute{
					Name:        "verify_xml_signature",
					Description: `Verify metadata’s XML signature`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_interface_ethernet",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder interface ethernet resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_interface_ethernet_bfd",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder interface ethernet bfd resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_interface_ethernet_ipv6",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder interface ethernet ipv6 resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_interface_ethernet_lldp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder interface ethernet lldp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
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
			Type:             "vthunder_thunder_interface_ethernet_trunk_group",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder interface ethernet trunk group resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_interface_lif",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder interface lif resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "lif",
					Description: `Logical interface`,
				},
				resource.Attribute{
					Name:        "ifname",
					Description: `Lif interface name`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `OSPF interface MTU (MTU size)`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `'enable': Enable; 'disable': Disable;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name for the interface`,
				},
				resource.Attribute{
					Name:        "acl-id",
					Description: `ACL id`,
				},
				resource.Attribute{
					Name:        "acl-name",
					Description: `Apply an access list (Named Access List)`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "user-tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all': all; 'num_pkts': num_pkts; 'num_total_bytes': num_total_bytes; 'num_unicast_pkts': num_unicast_pkts; 'num_broadcast_pkts': num_broadcast_pkts; 'num_multicast_pkts': num_multicast_pkts; 'num_tx_pkts': num_tx_pkts; 'num_total_tx_bytes': num_total_tx_bytes; 'num_unicast_tx_pkts': num_unicast_tx_pkts; 'num_broadcast_tx_pkts': num_broadcast_tx_pkts; 'num_multicast_tx_pkts': num_multicast_tx_pkts; 'dropped_dis_rx_pkts': dropped_dis_rx_pkts; 'dropped_rx_pkts': dropped_rx_pkts; 'dropped_dis_tx_pkts': dropped_dis_tx_pkts; 'dropped_tx_pkts': dropped_tx_pkts; 'dropped_rx_pkts_gre_key': dropped_rx_pkts_gre_key;`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Use DHCP to configure IP address`,
				},
				resource.Attribute{
					Name:        "ipv6-addr",
					Description: `Set the IPv6 address of an interface`,
				},
				resource.Attribute{
					Name:        "anycast",
					Description: `Configure an IPv6 anycast address`,
				},
				resource.Attribute{
					Name:        "link-local",
					Description: `Configure an IPv6 link local address`,
				},
				resource.Attribute{
					Name:        "allow-promiscuous-vip",
					Description: `Allow traffic to be associated with promiscuous VIP`,
				},
				resource.Attribute{
					Name:        "cache-spoofing-port",
					Description: `This interface connects to spoofing cache`,
				},
				resource.Attribute{
					Name:        "inside",
					Description: `Configure interface as inside`,
				},
				resource.Attribute{
					Name:        "outside",
					Description: `Configure interface as outside`,
				},
				resource.Attribute{
					Name:        "generate-membership-query",
					Description: `Enable Membership Query`,
				},
				resource.Attribute{
					Name:        "query-interval",
					Description: `1 - 255 (Default is 125)`,
				},
				resource.Attribute{
					Name:        "max-resp-time",
					Description: `Maximum Response Time (Max Response Time (Default is 100))`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `ISO routing area tag`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `The RIP authentication string`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `'md5': Keyed message digest;`,
				},
				resource.Attribute{
					Name:        "key-chain",
					Description: `Authentication key-chain (Name of key-chain)`,
				},
				resource.Attribute{
					Name:        "send-packet",
					Description: `Enable sending packets through the specified interface`,
				},
				resource.Attribute{
					Name:        "receive-packet",
					Description: `Enable receiving packet through the specified interface`,
				},
				resource.Attribute{
					Name:        "send",
					Description: `Advertisement transmission`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `'1': RIP version 1; '2': RIP version 2; '1-2': RIP version 1 & 2;`,
				},
				resource.Attribute{
					Name:        "receive",
					Description: `Advertisement reception`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `'poisoned': Perform split horizon with poisoned reverse; 'disable': Disable split horizon; 'enable': Perform split horizon without poisoned reverse;`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Enable authentication`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Mesh group number`,
				},
				resource.Attribute{
					Name:        "authentication-key",
					Description: `Authentication password (key) (The OSPF password (key))`,
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
					Name:        "database-filter",
					Description: `'all': Filter all LSA;`,
				},
				resource.Attribute{
					Name:        "out",
					Description: `Outgoing LSA`,
				},
				resource.Attribute{
					Name:        "dead-interval",
					Description: `Interval after which a neighbor is declared dead (Seconds)`,
				},
				resource.Attribute{
					Name:        "hello-interval",
					Description: `Set Hello interval in seconds (Hello interval value)`,
				},
				resource.Attribute{
					Name:        "message-digest-key",
					Description: `Message digest authentication password (key) (Key id)`,
				},
				resource.Attribute{
					Name:        "md5-value",
					Description: `The OSPF password (1-16)`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "mtu-ignore",
					Description: `Ignores the MTU in DBD packets`,
				},
				resource.Attribute{
					Name:        "broadcast",
					Description: `Specify OSPF broadcast multi-access network`,
				},
				resource.Attribute{
					Name:        "non-broadcast",
					Description: `Specify OSPF NBMA network`,
				},
				resource.Attribute{
					Name:        "point-to-point",
					Description: `Specify OSPF point-to-point network`,
				},
				resource.Attribute{
					Name:        "point-to-multipoint",
					Description: `Specify OSPF point-to-multipoint network`,
				},
				resource.Attribute{
					Name:        "p2mp-nbma",
					Description: `Specify non-broadcast point-to-multipoint network`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Set priority for Designated Router election (Priority value)`,
				},
				resource.Attribute{
					Name:        "retransmit-interval",
					Description: `Set per-LSP retransmission interval (Interval between retransmissions of the same LSP (seconds))`,
				},
				resource.Attribute{
					Name:        "transmit-delay",
					Description: `Link state transmit delay (Seconds)`,
				},
				resource.Attribute{
					Name:        "ip-addr",
					Description: `Address of interface`,
				},
				resource.Attribute{
					Name:        "ipv6-enable",
					Description: `Enable IPv6 processing`,
				},
				resource.Attribute{
					Name:        "rip",
					Description: `RIP Routing for IPv6`,
				},
				resource.Attribute{
					Name:        "area-id-num",
					Description: `OSPF area ID as a decimal value`,
				},
				resource.Attribute{
					Name:        "area-id-addr",
					Description: `OSPF area ID in IP address format`,
				},
				resource.Attribute{
					Name:        "instance-id",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "broadcast-type",
					Description: `'broadcast': Specify OSPF broadcast multi-access network; 'non-broadcast': Specify OSPF NBMA network; 'point-to-point': Specify OSPF point-to-point network; 'point-to-multipoint': Specify OSPF point-to-multipoint network;`,
				},
				resource.Attribute{
					Name:        "network-instance-id",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "neighbor",
					Description: `OSPFv3 neighbor (Neighbor IPv6 address)`,
				},
				resource.Attribute{
					Name:        "neig-inst",
					Description: `Specify the interface instance ID`,
				},
				resource.Attribute{
					Name:        "neighbor-cost",
					Description: `OSPF cost for point-to-multipoint neighbor (metric)`,
				},
				resource.Attribute{
					Name:        "neighbor-poll-interval",
					Description: `OSPF dead-router polling interval (Seconds)`,
				},
				resource.Attribute{
					Name:        "neighbor-priority",
					Description: `OSPF priority of non-broadcast neighbor`,
				},
				resource.Attribute{
					Name:        "key-id",
					Description: `Key ID`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `'md5': Keyed MD5; 'meticulous-md5': Meticulous Keyed MD5; 'meticulous-sha1': Meticulous Keyed SHA1; 'sha1': Keyed SHA1; 'simple': Simple Password;`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Configure the authentication password for interface`,
				},
				resource.Attribute{
					Name:        "echo",
					Description: `Enable BFD Echo`,
				},
				resource.Attribute{
					Name:        "demand",
					Description: `Demand mode`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Transmit interval between BFD packets (Milliseconds)`,
				},
				resource.Attribute{
					Name:        "min-rx",
					Description: `Minimum receive interval capability (Milliseconds)`,
				},
				resource.Attribute{
					Name:        "multiplier",
					Description: `Multiplier value used to compute holddown (value used to multiply the interval)`,
				},
				resource.Attribute{
					Name:        "send-only",
					Description: `Authentication send-only`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `'level-1': Apply metric to level-1 links; 'level-2': Apply metric to level-2 links;`,
				},
				resource.Attribute{
					Name:        "circuit-type",
					Description: `'level-1': Level-1 only adjacencies are formed; 'level-1-2': Level-1-2 adjacencies are formed; 'level-2-only': Level-2 only adjacencies are formed;`,
				},
				resource.Attribute{
					Name:        "csnp-interval",
					Description: `Set CSNP interval in seconds (CSNP interval value)`,
				},
				resource.Attribute{
					Name:        "padding",
					Description: `Add padding to IS-IS hello packets`,
				},
				resource.Attribute{
					Name:        "hello-interval-minimal",
					Description: `Set Hello holdtime 1 second, interval depends on multiplier`,
				},
				resource.Attribute{
					Name:        "hello-multiplier",
					Description: `Set multiplier for Hello holding time (Hello multiplier value)`,
				},
				resource.Attribute{
					Name:        "lsp-interval",
					Description: `Set LSP transmission interval (LSP transmission interval (milliseconds))`,
				},
				resource.Attribute{
					Name:        "blocked",
					Description: `Block LSPs on this interface`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `Configure the metric for interface (Default metric)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `'broadcast': Specify IS-IS broadcast multi-access network; 'point-to-point': Specify IS-IS point-to-point network;`,
				},
				resource.Attribute{
					Name:        "wide-metric",
					Description: `Configure the wide metric for interface`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_interface_lif_ip",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder interface lif ip resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip",
					Description: `Global IP configuration subcommands`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Use DHCP to configure IP address`,
				},
				resource.Attribute{
					Name:        "ipv4-address",
					Description: `IP address`,
				},
				resource.Attribute{
					Name:        "ipv4-netmask",
					Description: `IP subnet mask`,
				},
				resource.Attribute{
					Name:        "allow-promiscuous-vip",
					Description: `Allow traffic to be associated with promiscuous VIP`,
				},
				resource.Attribute{
					Name:        "cache-spoofing-port",
					Description: `This interface connects to spoofing cache`,
				},
				resource.Attribute{
					Name:        "inside",
					Description: `Configure interface as inside`,
				},
				resource.Attribute{
					Name:        "outside",
					Description: `Configure interface as outside`,
				},
				resource.Attribute{
					Name:        "generate-membership-query",
					Description: `Enable Membership Query`,
				},
				resource.Attribute{
					Name:        "query-interval",
					Description: `1 - 255 (Default is 125)`,
				},
				resource.Attribute{
					Name:        "max-resp-time",
					Description: `Maximum Response Time (Max Response Time (Default is 100))`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `ISO routing area tag`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `The RIP authentication string`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `'md5': Keyed message digest; 'text': Clear text authentication;`,
				},
				resource.Attribute{
					Name:        "key-chain",
					Description: `Authentication key-chain (Name of key-chain)`,
				},
				resource.Attribute{
					Name:        "send-packet",
					Description: `Enable sending packets through the specified interface`,
				},
				resource.Attribute{
					Name:        "receive-packet",
					Description: `Enable receiving packet through the specified interface`,
				},
				resource.Attribute{
					Name:        "send",
					Description: `Advertisement transmission`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `'1': RIP version 1; '2': RIP version 2; '1-2': RIP version 1 & 2;`,
				},
				resource.Attribute{
					Name:        "receive",
					Description: `Advertisement reception`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `'poisoned': Perform split horizon with poisoned reverse; 'disable': Disable split horizon; 'enable': Perform split horizon without poisoned reverse;`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Enable authentication`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `'message-digest': Use message-digest authentication; 'null': Use no authentication;`,
				},
				resource.Attribute{
					Name:        "authentication-key",
					Description: `Authentication password (key) (The OSPF password (key))`,
				},
				resource.Attribute{
					Name:        "bfd",
					Description: `Bidirectional Forwarding Detection (BFD)`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `'all': All functionality;`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `Interface cost`,
				},
				resource.Attribute{
					Name:        "database-filter",
					Description: `'all': Filter all LSA;`,
				},
				resource.Attribute{
					Name:        "out",
					Description: `Outgoing LSA`,
				},
				resource.Attribute{
					Name:        "dead-interval",
					Description: `Interval after which a neighbor is declared dead (Seconds)`,
				},
				resource.Attribute{
					Name:        "hello-interval",
					Description: `Time between HELLO packets (Seconds)`,
				},
				resource.Attribute{
					Name:        "message-digest-key",
					Description: `Message digest authentication password (key) (Key id)`,
				},
				resource.Attribute{
					Name:        "md5-value",
					Description: `The OSPF password (1-16)`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `OSPF interface MTU (MTU size)`,
				},
				resource.Attribute{
					Name:        "mtu-ignore",
					Description: `Ignores the MTU in DBD packets`,
				},
				resource.Attribute{
					Name:        "broadcast",
					Description: `Specify OSPF broadcast multi-access network`,
				},
				resource.Attribute{
					Name:        "non-broadcast",
					Description: `Specify OSPF NBMA network`,
				},
				resource.Attribute{
					Name:        "point-to-point",
					Description: `Specify OSPF point-to-point network`,
				},
				resource.Attribute{
					Name:        "point-to-multipoint",
					Description: `Specify OSPF point-to-multipoint network`,
				},
				resource.Attribute{
					Name:        "p2mp-nbma",
					Description: `Specify non-broadcast point-to-multipoint network`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Router priority`,
				},
				resource.Attribute{
					Name:        "retransmit-interval",
					Description: `Time between retransmitting lost link state advertisements (Seconds)`,
				},
				resource.Attribute{
					Name:        "transmit-delay",
					Description: `Link state transmit delay (Seconds)`,
				},
				resource.Attribute{
					Name:        "ip-addr",
					Description: `Address of interface`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_interface_management",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder interface management resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_interface_ve",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder interface ve resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_interface_ve_bfd",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder interface ve bfd resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_interface_ve_ip",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder interface ve ip resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_interface_ve_ipv6",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder interface ve ipv6 resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ip_address",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip address resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ip_dns_primary",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip dns primary resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ip_dns_secondary",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip dns secondary resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ip_dns_suffix",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip dns suffix resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ip_frag",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip frag resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ip_icmp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip icmp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ip_nat_alg_pptp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip nat alg pptp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ip_nat_global",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip nat global resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ip_nat_icmp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip nat icmp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ip_nat_pool",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip nat pool resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "end_address",
					Description: `Configure end IP address of NAT pool`,
				},
				resource.Attribute{
					Name:        "ethernet",
					Description: `Ethernet interface`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Configure gateway IP`,
				},
				resource.Attribute{
					Name:        "ip_rr",
					Description: `Use IP address round-robin behavior`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `Configure mask for pool`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `Specify pool name or pool group`,
				},
				resource.Attribute{
					Name:        "scaleout_device_id",
					Description: `Configure Scaleout device id to which this NAT pool is to be bound (Specify Scaleout device id)`,
				},
				resource.Attribute{
					Name:        "start_address",
					Description: `Configure start IP address of NAT pool`,
				},
				resource.Attribute{
					Name:        "use_if_ip",
					Description: `Use Interface IP`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "vrid",
					Description: `Configure VRRP-A vrid (Specify ha VRRP-A vrid)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_ip_nat_pool_group",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip nat pool group resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pool_group_name",
					Description: `Specify pool group name`,
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
					Name:        "vrid",
					Description: `Specify VRRP-A vrid (Specify ha VRRP-A vrid)`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `Specify NAT pool name`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_ip_prefix_list",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip prefix list resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ip_reroute",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip reroute resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ip_route_static_bfd",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip route static bfd resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ip_tcp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ip tcp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ipv6_frag",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ipv6 frag resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ipv6_icmpv6",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ipv6 icmpv6 resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ipv6_nat_icmpv6",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ipv6 nat icmpv6 resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_ntp_auth_key",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ntp auth key resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alg_type",
					Description: `‘M’: encryption using MD5; ‘SHA’: encryption using SHA; ‘SHA1’: encryption using SHA1;`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "hex_encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `authentication key`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `‘ascii’: key string in ASCII form; ‘hex’: key string in hex form;`,
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
			Type:             "vthunder_thunder_ntp_server_hostname",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ntp server hostname resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `‘enable’: Enable this NTP server; ‘disable’: Disable this NTP server;`,
				},
				resource.Attribute{
					Name:        "host_servername",
					Description: `IPV4 address, IPV6 address or host name of NTP server(string1~63)`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Use trusted key to authenticate this NTP server (The trusted key number to use)`,
				},
				resource.Attribute{
					Name:        "prefer",
					Description: `Set this NTP server as preferred`,
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
			Type:             "vthunder_thunder_ntp_trusted_key",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder ntp trusted key resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `trusted key`,
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
			Type:             "vthunder_thunder_overlay_tunnel_options",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder overlay tunnel options resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_overlay_tunnel_vtep",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder overlay tunnel vtep resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vtep",
					Description: `Virtual Tunnel end point Configuration`,
				},
				resource.Attribute{
					Name:        "id1",
					Description: `VTEP Identifier`,
				},
				resource.Attribute{
					Name:        "encap",
					Description: `'nvgre': Tunnel Encapsulation Type is NVGRE; 'vxlan': Tunnel Encapsulation Type is VXLAN;`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "user-tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all': all; 'cfg_err_count': Config errors; 'flooded_pkt_count': Flooded packet count; 'encap_unresolved_count': Encap unresolved failures; 'unknown_encap_rx_pkt': Encap miss rx pkts; 'unknown_encap_tx_pkt': Encap miss tx pkts; 'arp_req_sent': Arp request sent; 'vtep_host_learned': Hosts learned; 'vtep_host_learn_error': Host learn error; 'invalid_lif_rx': Invalid Lif pkts in; 'invalid_lif_tx': Invalid Lif pkts out; 'unknown_vtep_tx': Vtep unknown tx; 'unknown_vtep_rx': Vtep Unkown rx; 'unhandled_pkt_rx': Unhandled packets in; 'unhandled_pkt_tx': Unhandled packets out; 'total_pkts_rx': Total packets out; 'total_bytes_rx': Total packet bytes in; 'unicast_pkt_rx': Total unicast packets in; 'bcast_pkt_rx': Total broadcast packets in; 'mcast_pkt_rx': Total multicast packets in; 'dropped_pkt_rx': Dropped received packets; 'encap_miss_pkts_rx': Encap missed in received packets; 'bad_chksum_pks_rx': Bad checksum in received packets; 'requeue_pkts_in': Requeued packets in; 'pkts_out': Packets out; 'total_bytes_tx': Packet bytes out; 'unicast_pkt_tx': Unicast packets out; 'bcast_pkt_tx': Broadcast packets out; 'mcast_pkt_tx': Multicast packets out; 'dropped_pkts_tx': Dropped packets out; 'large_pkts_rx': Too large packets in; 'dot1q_pkts_rx': Dot1q packets in; 'frag_pkts_tx': Frag packets out; 'reassembled_pkts_rx': Reassembled packets in; 'bad_inner_ipv4_len_rx': bad inner ipv4 packet len; 'bad_inner_ipv6_len_rx': Bad inner ipv6 packet len; 'frag_drop_pkts_tx': Frag dropped packets out; 'lif_un_init_rx': Lif uninitialized packets in;`,
				},
				resource.Attribute{
					Name:        "ip-address",
					Description: `IP Address of the remote VTEP`,
				},
				resource.Attribute{
					Name:        "segment",
					Description: `VNI configured for the remote VTEP`,
				},
				resource.Attribute{
					Name:        "ipv6-address",
					Description: `IPv6 Address of the remote VTEP`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `Name of the Partition with the L2 segment being extended (Name of the User Partition with the L2 segment being extended)`,
				},
				resource.Attribute{
					Name:        "lif",
					Description: `Logical interface binding the Provider Partition to the User Partition (logical interface name)`,
				},
				resource.Attribute{
					Name:        "retry-time",
					Description: `Keepalive retry interval in seconds`,
				},
				resource.Attribute{
					Name:        "retry-count",
					Description: `Keepalive multiplier`,
				},
				resource.Attribute{
					Name:        "gre-key",
					Description: `key`,
				},
				resource.Attribute{
					Name:        "ip-addr",
					Description: `IPv4 address of the overlay host`,
				},
				resource.Attribute{
					Name:        "overlay-mac-addr",
					Description: `MAC Address of the overlay host`,
				},
				resource.Attribute{
					Name:        "vni",
					Description: `Configure the segment id ( VNI of the remote host)`,
				},
				resource.Attribute{
					Name:        "remote-vtep",
					Description: `Configure the VTEP IP address (IPv4 address of the VTEP for the remote host)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_partition",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder partition resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
					Name:        "id1",
					Description: `Specify unique Partition id`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_reboot",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder reboot resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all",
					Description: `Reboot all devices when VCS is enabled, or only this device itself if VCS is not enabled`,
				},
				resource.Attribute{
					Name:        "at",
					Description: `Reboot at a Specific time/date`,
				},
				resource.Attribute{
					Name:        "cancel",
					Description: `Cancel Pending Reboot`,
				},
				resource.Attribute{
					Name:        "day_of_month",
					Description: `Day of the Month`,
				},
				resource.Attribute{
					Name:        "day_of_month_2",
					Description: `Day of the Month`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Reboot a specific device when VCS is enabled (device id)`,
				},
				resource.Attribute{
					Name:        "in",
					Description: `Reboot after a time interval (Time in hours and minutes)`,
				},
				resource.Attribute{
					Name:        "month",
					Description: `‘January’: Month of the year; ‘February’: Month of the year; ‘March’: Month of the year; ‘April’: Month of the year; ‘May’: Month of the year; ‘June’: Month of the year; ‘July’: Month of the year; ‘August’: Month of the year; ‘September’: Month of the year; ‘October’: Month of the year; ‘November’: Month of the year; ‘December’: Month of the year;`,
				},
				resource.Attribute{
					Name:        "month_2",
					Description: `‘January’: Month of the year; ‘February’: Month of the year; ‘March’: Month of the year; ‘April’: Month of the year; ‘May’: Month of the year; ‘June’: Month of the year; ‘July’: Month of the r; ‘August’: Month of the year; ‘September’: Month of the year; ‘October’: Month of the year; ‘November’: Month of the year; ‘December’: Month of the year;`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `Reason for Reboot`,
				},
				resource.Attribute{
					Name:        "reason_2",
					Description: `Reason for Reboot`,
				},
				resource.Attribute{
					Name:        "reason_3",
					Description: `Reason for Reboot`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `Time to Reboot (hh:mm)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_rib_route",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder rib route resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_dest_addr",
					Description: `Destination prefix`,
				},
				resource.Attribute{
					Name:        "ip_mask",
					Description: `Destination prefix mask`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "description_nexthop_lif",
					Description: `Description for static route`,
				},
				resource.Attribute{
					Name:        "lif",
					Description: `LIF Interface (Logical tunnel interface number)`,
				},
				resource.Attribute{
					Name:        "description_nexthop_ip",
					Description: `Description for static route`,
				},
				resource.Attribute{
					Name:        "distance_nexthop_ip",
					Description: `Distance value for this route`,
				},
				resource.Attribute{
					Name:        "ip_next_hop",
					Description: `Forwarding router’s address`,
				},
				resource.Attribute{
					Name:        "description_nexthop_tunnel",
					Description: `Description for static route`,
				},
				resource.Attribute{
					Name:        "distance_nexthop_tunnel",
					Description: `Distance value for this route`,
				},
				resource.Attribute{
					Name:        "ip_next_hop_tunnel",
					Description: `Forwarding router’s address`,
				},
				resource.Attribute{
					Name:        "tunnel",
					Description: `Tunnel interface (Tunnel interface number)`,
				},
				resource.Attribute{
					Name:        "description_nexthop_partition",
					Description: `Description for static route`,
				},
				resource.Attribute{
					Name:        "description_partition_vrid",
					Description: `Description for static route`,
				},
				resource.Attribute{
					Name:        "partition_name",
					Description: `Name of network partition`,
				},
				resource.Attribute{
					Name:        "vrid_num_in_partition",
					Description: `Specify ha VRRP-A vrid`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_route_map",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder route map resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route-map",
					Description: `Configure route-map`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Route map tag`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `'permit': Route map permits set operations; 'deny': Route map denies set operations;`,
				},
				resource.Attribute{
					Name:        "sequence",
					Description: `Sequence to insert to/delete from existing route-map entry`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "user-tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Community-list name`,
				},
				resource.Attribute{
					Name:        "exact-match",
					Description: `Do exact matching of ecommunities`,
				},
				resource.Attribute{
					Name:        "group-id",
					Description: `HA or VRRP-A group id`,
				},
				resource.Attribute{
					Name:        "ha-state",
					Description: `'active': HA or VRRP-A in Active State; 'standby': HA or VRRP-A in Standby State;`,
				},
				resource.Attribute{
					Name:        "cluster-id",
					Description: `Scaleout Cluster-id`,
				},
				resource.Attribute{
					Name:        "operational-state",
					Description: `'up': Scaleout is up and running; 'down': Scaleout is down or disabled;`,
				},
				resource.Attribute{
					Name:        "ethernet",
					Description: `Ethernet interface (Port number)`,
				},
				resource.Attribute{
					Name:        "loopback",
					Description: `Loopback interface (Port number)`,
				},
				resource.Attribute{
					Name:        "trunk",
					Description: `Trunk Interface (Trunk interface number)`,
				},
				resource.Attribute{
					Name:        "ve",
					Description: `Virtual ethernet interface (Virtual ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "tunnel",
					Description: `Tunnel interface (Tunnel interface number)`,
				},
				resource.Attribute{
					Name:        "val",
					Description: `Preference value`,
				},
				resource.Attribute{
					Name:        "egp",
					Description: `remote EGP`,
				},
				resource.Attribute{
					Name:        "igp",
					Description: `local IGP`,
				},
				resource.Attribute{
					Name:        "incomplete",
					Description: `unknown heritage`,
				},
				resource.Attribute{
					Name:        "acl1",
					Description: `IPv6 access-list number`,
				},
				resource.Attribute{
					Name:        "acl2",
					Description: `IP access-list number (expanded range)`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `Exact match a prefix (Prefix IPv6 address)`,
				},
				resource.Attribute{
					Name:        "reachable",
					Description: `IPv6 address is reachable`,
				},
				resource.Attribute{
					Name:        "unreachable",
					Description: `IPv6 address is unreachable`,
				},
				resource.Attribute{
					Name:        "next-hop-acl-name",
					Description: `IPv6 access-list name`,
				},
				resource.Attribute{
					Name:        "v6-addr",
					Description: `IPv6 address of next hop`,
				},
				resource.Attribute{
					Name:        "prefix-list-name",
					Description: `IPv6 prefix-list name`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `VPN extended community`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `IPv6 address of next hop`,
				},
				resource.Attribute{
					Name:        "class-list-name",
					Description: `Class-List Name`,
				},
				resource.Attribute{
					Name:        "class-list-cid",
					Description: `Class-List Cid`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Zone Name`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `AS number`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of aggregator`,
				},
				resource.Attribute{
					Name:        "prepend",
					Description: `Prepend to the as-path (AS number)`,
				},
				resource.Attribute{
					Name:        "num",
					Description: `AS number`,
				},
				resource.Attribute{
					Name:        "num2",
					Description: `AS number`,
				},
				resource.Attribute{
					Name:        "atomic-aggregate",
					Description: `BGP atomic aggregate attribute`,
				},
				resource.Attribute{
					Name:        "v-std",
					Description: `Community-list number (standard)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `Delete matching communities`,
				},
				resource.Attribute{
					Name:        "v-exp",
					Description: `Community-list number (expanded)`,
				},
				resource.Attribute{
					Name:        "v-exp-delete",
					Description: `Delete matching communities`,
				},
				resource.Attribute{
					Name:        "name-delete",
					Description: `Delete matching communities`,
				},
				resource.Attribute{
					Name:        "community",
					Description: `BGP community attribute`,
				},
				resource.Attribute{
					Name:        "dampening",
					Description: `Enable route-flap dampening`,
				},
				resource.Attribute{
					Name:        "dampening-half-time",
					Description: `Reachability Half-life time for the penalty(minutes)`,
				},
				resource.Attribute{
					Name:        "dampening-reuse",
					Description: `Value to start reusing a route`,
				},
				resource.Attribute{
					Name:        "dampening-supress",
					Description: `Value to start suppressing a route`,
				},
				resource.Attribute{
					Name:        "dampening-max-supress",
					Description: `Maximum duration to suppress a stable route(minutes)`,
				},
				resource.Attribute{
					Name:        "dampening-penalty",
					Description: `Un-reachability Half-life time for the penalty(minutes)`,
				},
				resource.Attribute{
					Name:        "originator-ip",
					Description: `IP address of originator`,
				},
				resource.Attribute{
					Name:        "weight-val",
					Description: `Weight value`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_router_bgp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bgp",
					Description: `Border Gateway Protocol (BGP)`,
				},
				resource.Attribute{
					Name:        "as-number",
					Description: `AS number`,
				},
				resource.Attribute{
					Name:        "aggregate-address",
					Description: `Configure BGP aggregate entries (Aggregate IPv6 prefix)`,
				},
				resource.Attribute{
					Name:        "as-set",
					Description: `Generate AS set path information`,
				},
				resource.Attribute{
					Name:        "summary-only",
					Description: `Filter more specific routes from updates`,
				},
				resource.Attribute{
					Name:        "always-compare-med",
					Description: `Allow comparing MED from different neighbors`,
				},
				resource.Attribute{
					Name:        "ignore",
					Description: `Ignore as-path length in selecting a route`,
				},
				resource.Attribute{
					Name:        "compare-routerid",
					Description: `Compare router-id for identical EBGP paths`,
				},
				resource.Attribute{
					Name:        "remove-recv-med",
					Description: `To remove rcvd MED attribute`,
				},
				resource.Attribute{
					Name:        "remove-send-med",
					Description: `To remove send MED attribute`,
				},
				resource.Attribute{
					Name:        "missing-as-worst",
					Description: `Treat missing MED as the least preferred one`,
				},
				resource.Attribute{
					Name:        "dampening",
					Description: `Enable route-flap dampening`,
				},
				resource.Attribute{
					Name:        "dampening-half-time",
					Description: `Reachability Half-life time for the penalty(minutes)`,
				},
				resource.Attribute{
					Name:        "dampening-reuse",
					Description: `Value to start reusing a route`,
				},
				resource.Attribute{
					Name:        "dampening-supress",
					Description: `Value to start suppressing a route`,
				},
				resource.Attribute{
					Name:        "dampening-max-supress",
					Description: `Maximum duration to suppress a stable route(minutes)`,
				},
				resource.Attribute{
					Name:        "dampening-penalty",
					Description: `Un-reachability Half-life time for the penalty(minutes)`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route map reference (Pointer to route-map entries)`,
				},
				resource.Attribute{
					Name:        "local-preference-value",
					Description: `Configure default local preference value`,
				},
				resource.Attribute{
					Name:        "deterministic-med",
					Description: `Pick the best-MED path among paths advertised from the neighboring AS`,
				},
				resource.Attribute{
					Name:        "enforce-first-as",
					Description: `Enforce the first AS for EBGP routes`,
				},
				resource.Attribute{
					Name:        "fast-external-failover",
					Description: `Immediately reset session if a link to a directly connected external peer goes down`,
				},
				resource.Attribute{
					Name:        "log-neighbor-changes",
					Description: `Log neighbor up/down and reset reason`,
				},
				resource.Attribute{
					Name:        "nexthop-trigger-count",
					Description: `BGP nexthop-tracking status (count)`,
				},
				resource.Attribute{
					Name:        "router-id",
					Description: `Override current router identifier (peers will reset) (Manually configured router identifier)`,
				},
				resource.Attribute{
					Name:        "override-validation",
					Description: `override router-id validation`,
				},
				resource.Attribute{
					Name:        "scan-time",
					Description: `Configure background scan interval (Scan interval (sec) [Default:60 Disable:0])`,
				},
				resource.Attribute{
					Name:        "graceful-restart",
					Description: `enable graceful-restart helper for this neighbor`,
				},
				resource.Attribute{
					Name:        "bgp-restart-time",
					Description: `BGP Peer Graceful Restart time in seconds (default 90)`,
				},
				resource.Attribute{
					Name:        "bgp-stalepath-time",
					Description: `BGP Graceful Restart Stalepath retention time in seconds (default 360)`,
				},
				resource.Attribute{
					Name:        "admin-distance",
					Description: `Define an administrative distance`,
				},
				resource.Attribute{
					Name:        "src-prefix",
					Description: `IP source prefix`,
				},
				resource.Attribute{
					Name:        "acl-str",
					Description: `Access list name`,
				},
				resource.Attribute{
					Name:        "ext-routes-dist",
					Description: `Distance for routes external to the AS`,
				},
				resource.Attribute{
					Name:        "int-routes-dist",
					Description: `Distance for routes internal to the AS`,
				},
				resource.Attribute{
					Name:        "local-routes-dist",
					Description: `Distance for local routes`,
				},
				resource.Attribute{
					Name:        "maximum-paths-value",
					Description: `Supported BGP multipath numbers`,
				},
				resource.Attribute{
					Name:        "originate",
					Description: `Distribute an IPv6 default route`,
				},
				resource.Attribute{
					Name:        "bgp-keepalive",
					Description: `Keepalive interval`,
				},
				resource.Attribute{
					Name:        "bgp-holdtime",
					Description: `Holdtime`,
				},
				resource.Attribute{
					Name:        "synchronization",
					Description: `Perform IGP synchronization`,
				},
				resource.Attribute{
					Name:        "auto-summary",
					Description: `Enable automatic network number summarization`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "user-tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "network-synchronization",
					Description: `Perform IGP synchronization`,
				},
				resource.Attribute{
					Name:        "network-ipv4-cidr",
					Description: `Specify network mask`,
				},
				resource.Attribute{
					Name:        "backdoor",
					Description: `Specify a BGP backdoor route`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Network specific description (Up to 80 characters describing this network)`,
				},
				resource.Attribute{
					Name:        "comm-value",
					Description: `community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export`,
				},
				resource.Attribute{
					Name:        "peer-group",
					Description: `Neighbor tag`,
				},
				resource.Attribute{
					Name:        "activate",
					Description: `Enable the Address Family for this Neighbor`,
				},
				resource.Attribute{
					Name:        "allowas-in",
					Description: `Accept as-path with my AS present in it`,
				},
				resource.Attribute{
					Name:        "allowas-in-count",
					Description: `Number of occurrences of AS number`,
				},
				resource.Attribute{
					Name:        "prefix-list-direction",
					Description: `'both': both; 'receive': receive; 'send': send;`,
				},
				resource.Attribute{
					Name:        "default-originate",
					Description: `Originate default route to this neighbor`,
				},
				resource.Attribute{
					Name:        "distribute-list",
					Description: `Filter updates to/from this neighbor (IP standard/extended/named access list)`,
				},
				resource.Attribute{
					Name:        "distribute-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "filter-list",
					Description: `Establish BGP filters (AS path access-list name)`,
				},
				resource.Attribute{
					Name:        "filter-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "maximum-prefix",
					Description: `Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))`,
				},
				resource.Attribute{
					Name:        "maximum-prefix-thres",
					Description: `threshold-value, 1 to 100 percent`,
				},
				resource.Attribute{
					Name:        "next-hop-self",
					Description: `Disable the next hop calculation for this neighbor`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list",
					Description: `Filter updates to/from this neighbor (Name of a prefix list)`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "remove-private-as",
					Description: `Remove private AS number from outbound updates`,
				},
				resource.Attribute{
					Name:        "nbr-route-map",
					Description: `Apply route map to neighbor (Name of route map)`,
				},
				resource.Attribute{
					Name:        "nbr-rmap-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "send-community-val",
					Description: `'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `Allow inbound soft reconfiguration for this neighbor`,
				},
				resource.Attribute{
					Name:        "unsuppress-map",
					Description: `Route-map to selectively unsuppress suppressed routes (Name of route map)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Set default weight for routes from this neighbor`,
				},
				resource.Attribute{
					Name:        "neighbor-ipv4",
					Description: `Neighbor address`,
				},
				resource.Attribute{
					Name:        "peer-group-name",
					Description: `Configure peer-group (peer-group name)`,
				},
				resource.Attribute{
					Name:        "neighbor-ipv6",
					Description: `Neighbor IPv6 address`,
				},
				resource.Attribute{
					Name:        "ethernet",
					Description: `Ethernet interface number`,
				},
				resource.Attribute{
					Name:        "ve",
					Description: `Virtual ethernet interface number`,
				},
				resource.Attribute{
					Name:        "trunk",
					Description: `Trunk interface number`,
				},
				resource.Attribute{
					Name:        "connected",
					Description: `Connected`,
				},
				resource.Attribute{
					Name:        "floating-ip",
					Description: `Floating IP`,
				},
				resource.Attribute{
					Name:        "lw4o6",
					Description: `LW4O6 Prefix`,
				},
				resource.Attribute{
					Name:        "static-nat",
					Description: `Static NAT Prefix`,
				},
				resource.Attribute{
					Name:        "ip-nat",
					Description: `IP NAT`,
				},
				resource.Attribute{
					Name:        "ip-nat-list",
					Description: `IP NAT list`,
				},
				resource.Attribute{
					Name:        "isis",
					Description: `ISO IS-IS`,
				},
				resource.Attribute{
					Name:        "ospf",
					Description: `Open Shortest Path First (OSPF)`,
				},
				resource.Attribute{
					Name:        "rip",
					Description: `Routing Information Protocol (RIP)`,
				},
				resource.Attribute{
					Name:        "static",
					Description: `Static routes`,
				},
				resource.Attribute{
					Name:        "nat-map",
					Description: `NAT MAP Prefix`,
				},
				resource.Attribute{
					Name:        "only-flagged",
					Description: `Selected Virtual IP (VIP)`,
				},
				resource.Attribute{
					Name:        "only-not-flagged",
					Description: `Only not flagged`,
				},
				resource.Attribute{
					Name:        "dampening-half",
					Description: `Reachability Half-life time for the penalty(minutes)`,
				},
				resource.Attribute{
					Name:        "dampening-start-reuse",
					Description: `Value to start reusing a route`,
				},
				resource.Attribute{
					Name:        "dampening-start-supress",
					Description: `Value to start suppressing a route`,
				},
				resource.Attribute{
					Name:        "dampening-unreachability",
					Description: `Un-reachability Half-life time for the penalty(minutes)`,
				},
				resource.Attribute{
					Name:        "distance-ext",
					Description: `Distance for routes external to the AS`,
				},
				resource.Attribute{
					Name:        "distance-int",
					Description: `Distance for routes internal to the AS`,
				},
				resource.Attribute{
					Name:        "distance-local",
					Description: `Distance for local routes`,
				},
				resource.Attribute{
					Name:        "network-ipv6",
					Description: `Specify a network to announce via BGP`,
				},
				resource.Attribute{
					Name:        "nat64",
					Description: `NAT64 Prefix`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_router_bgp_address_family_ipv6",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp address family ipv6 resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv6",
					Description: `ipv6 Address family`,
				},
				resource.Attribute{
					Name:        "dampening",
					Description: `Enable route-flap dampening`,
				},
				resource.Attribute{
					Name:        "dampening-half",
					Description: `Reachability Half-life time for the penalty(minutes)`,
				},
				resource.Attribute{
					Name:        "dampening-start-reuse",
					Description: `Value to start reusing a route`,
				},
				resource.Attribute{
					Name:        "dampening-start-supress",
					Description: `Value to start suppressing a route`,
				},
				resource.Attribute{
					Name:        "dampening-max-supress",
					Description: `Maximum duration to suppress a stable route(minutes)`,
				},
				resource.Attribute{
					Name:        "dampening-unreachability",
					Description: `Un-reachability Half-life time for the penalty(minutes)`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route map reference (Pointer to route-map entries)`,
				},
				resource.Attribute{
					Name:        "distance-ext",
					Description: `Distance for routes external to the AS`,
				},
				resource.Attribute{
					Name:        "distance-int",
					Description: `Distance for routes internal to the AS`,
				},
				resource.Attribute{
					Name:        "distance-local",
					Description: `Distance for local routes`,
				},
				resource.Attribute{
					Name:        "maximum-paths-value",
					Description: `Supported BGP multipath numbers`,
				},
				resource.Attribute{
					Name:        "originate",
					Description: `Distribute an IPv6 default route`,
				},
				resource.Attribute{
					Name:        "aggregate-address",
					Description: `Configure BGP aggregate entries (Aggregate IPv6 prefix)`,
				},
				resource.Attribute{
					Name:        "as-set",
					Description: `Generate AS set path information`,
				},
				resource.Attribute{
					Name:        "summary-only",
					Description: `Filter more specific routes from updates`,
				},
				resource.Attribute{
					Name:        "auto-summary",
					Description: `Enable automatic network number summarization`,
				},
				resource.Attribute{
					Name:        "synchronization",
					Description: `Perform IGP synchronization`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "network-synchronization",
					Description: `Perform IGP synchronization`,
				},
				resource.Attribute{
					Name:        "network-ipv6",
					Description: `Specify a network to announce via BGP`,
				},
				resource.Attribute{
					Name:        "backdoor",
					Description: `Specify a BGP backdoor route`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Network specific description (Up to 80 characters describing this network)`,
				},
				resource.Attribute{
					Name:        "comm-value",
					Description: `community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export`,
				},
				resource.Attribute{
					Name:        "peer-group",
					Description: `Neighbor tag`,
				},
				resource.Attribute{
					Name:        "activate",
					Description: `Enable the Address Family for this Neighbor`,
				},
				resource.Attribute{
					Name:        "allowas-in",
					Description: `Accept as-path with my AS present in it`,
				},
				resource.Attribute{
					Name:        "allowas-in-count",
					Description: `Number of occurrences of AS number`,
				},
				resource.Attribute{
					Name:        "prefix-list-direction",
					Description: `'both': both; 'receive': receive; 'send': send;`,
				},
				resource.Attribute{
					Name:        "default-originate",
					Description: `Originate default route to this neighbor`,
				},
				resource.Attribute{
					Name:        "distribute-list",
					Description: `Filter updates to/from this neighbor (IP standard/extended/named access list)`,
				},
				resource.Attribute{
					Name:        "distribute-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "filter-list",
					Description: `Establish BGP filters (AS path access-list name)`,
				},
				resource.Attribute{
					Name:        "filter-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "maximum-prefix",
					Description: `Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))`,
				},
				resource.Attribute{
					Name:        "maximum-prefix-thres",
					Description: `threshold-value, 1 to 100 percent`,
				},
				resource.Attribute{
					Name:        "next-hop-self",
					Description: `Disable the next hop calculation for this neighbor`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list",
					Description: `Filter updates to/from this neighbor (Name of a prefix list)`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "remove-private-as",
					Description: `Remove private AS number from outbound updates`,
				},
				resource.Attribute{
					Name:        "nbr-route-map",
					Description: `Apply route map to neighbor (Name of route map)`,
				},
				resource.Attribute{
					Name:        "nbr-rmap-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "send-community-val",
					Description: `'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `Allow inbound soft reconfiguration for this neighbor`,
				},
				resource.Attribute{
					Name:        "unsuppress-map",
					Description: `Route-map to selectively unsuppress suppressed routes (Name of route map)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Set default weight for routes from this neighbor`,
				},
				resource.Attribute{
					Name:        "neighbor-ipv4",
					Description: `Neighbor address`,
				},
				resource.Attribute{
					Name:        "peer-group-name",
					Description: `Configure peer-group (peer-group name)`,
				},
				resource.Attribute{
					Name:        "graceful-restart",
					Description: `enable graceful-restart helper for this neighbor`,
				},
				resource.Attribute{
					Name:        "neighbor-ipv6",
					Description: `Neighbor IPv6 address`,
				},
				resource.Attribute{
					Name:        "ethernet",
					Description: `Ethernet interface number`,
				},
				resource.Attribute{
					Name:        "ve",
					Description: `Virtual ethernet interface number`,
				},
				resource.Attribute{
					Name:        "trunk",
					Description: `Trunk interface number`,
				},
				resource.Attribute{
					Name:        "connected",
					Description: `Connected`,
				},
				resource.Attribute{
					Name:        "floating-ip",
					Description: `Floating IP`,
				},
				resource.Attribute{
					Name:        "nat64",
					Description: `NAT64 Prefix`,
				},
				resource.Attribute{
					Name:        "nat-map",
					Description: `NAT MAP Prefix`,
				},
				resource.Attribute{
					Name:        "lw4o6",
					Description: `LW4O6 Prefix`,
				},
				resource.Attribute{
					Name:        "static-nat",
					Description: `Static NAT Prefix`,
				},
				resource.Attribute{
					Name:        "ip-nat",
					Description: `IP NAT`,
				},
				resource.Attribute{
					Name:        "ip-nat-list",
					Description: `IP NAT list`,
				},
				resource.Attribute{
					Name:        "isis",
					Description: `ISO IS-IS`,
				},
				resource.Attribute{
					Name:        "ospf",
					Description: `Open Shortest Path First (OSPF)`,
				},
				resource.Attribute{
					Name:        "rip",
					Description: `Routing Information Protocol (RIP)`,
				},
				resource.Attribute{
					Name:        "static",
					Description: `Static routes`,
				},
				resource.Attribute{
					Name:        "only-flagged",
					Description: `Selected Virtual IP (VIP)`,
				},
				resource.Attribute{
					Name:        "only-not-flagged",
					Description: `Only not flagged`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_router_bgp_address_family_ipv6_neighbor_ethernet_neighbor_ipv6",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp address family ipv6 neighbor ethernet neighbor ipv6 resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ethernet-neighbor-ipv6",
					Description: `Specify an ethernet unnumbered neighbor`,
				},
				resource.Attribute{
					Name:        "ethernet",
					Description: `Ethernet interface number`,
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
			Type:             "vthunder_thunder_router_bgp_address_family_ipv6_neighbor_ipv4_neighbor",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp address family ipv6 neighbor ipv4 neighbor resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv4-neighbor",
					Description: `Specify a peer-group neighbor router`,
				},
				resource.Attribute{
					Name:        "neighbor-ipv4",
					Description: `Neighbor address`,
				},
				resource.Attribute{
					Name:        "peer-group-name",
					Description: `Configure peer-group (peer-group name)`,
				},
				resource.Attribute{
					Name:        "activate",
					Description: `Enable the Address Family for this Neighbor`,
				},
				resource.Attribute{
					Name:        "allowas-in",
					Description: `Accept as-path with my AS present in it`,
				},
				resource.Attribute{
					Name:        "allowas-in-count",
					Description: `Number of occurrences of AS number`,
				},
				resource.Attribute{
					Name:        "prefix-list-direction",
					Description: `'both': both; 'receive': receive; 'send': send;`,
				},
				resource.Attribute{
					Name:        "graceful-restart",
					Description: `enable graceful-restart helper for this neighbor`,
				},
				resource.Attribute{
					Name:        "default-originate",
					Description: `Originate default route to this neighbor`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route-map to specify criteria to originate default (route-map name)`,
				},
				resource.Attribute{
					Name:        "distribute-list",
					Description: `Filter updates to/from this neighbor (IP standard/extended/named access list)`,
				},
				resource.Attribute{
					Name:        "distribute-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "filter-list",
					Description: `Establish BGP filters (AS path access-list name)`,
				},
				resource.Attribute{
					Name:        "filter-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "maximum-prefix",
					Description: `Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))`,
				},
				resource.Attribute{
					Name:        "maximum-prefix-thres",
					Description: `threshold-value, 1 to 100 percent`,
				},
				resource.Attribute{
					Name:        "next-hop-self",
					Description: `Disable the next hop calculation for this neighbor`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list",
					Description: `Filter updates to/from this neighbor (Name of a prefix list)`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "remove-private-as",
					Description: `Remove private AS number from outbound updates`,
				},
				resource.Attribute{
					Name:        "nbr-route-map",
					Description: `Apply route map to neighbor (Name of route map)`,
				},
				resource.Attribute{
					Name:        "nbr-rmap-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "send-community-val",
					Description: `'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `Allow inbound soft reconfiguration for this neighbor`,
				},
				resource.Attribute{
					Name:        "unsuppress-map",
					Description: `Route-map to selectively unsuppress suppressed routes (Name of route map)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Set default weight for routes from this neighbor`,
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
			Type:             "vthunder_thunder_router_bgp_address_family_ipv6_neighbor_ipv6_neighbor",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp address family ipv6 neighbor ipv6 neighbor resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv6-neighbor",
					Description: `Specify a peer-group neighbor router`,
				},
				resource.Attribute{
					Name:        "neighbor-ipv6",
					Description: `Neighbor IPv6 address`,
				},
				resource.Attribute{
					Name:        "peer-group-name",
					Description: `Configure peer-group (peer-group name)`,
				},
				resource.Attribute{
					Name:        "activate",
					Description: `Enable the Address Family for this Neighbor`,
				},
				resource.Attribute{
					Name:        "allowas-in",
					Description: `Accept as-path with my AS present in it`,
				},
				resource.Attribute{
					Name:        "allowas-in-count",
					Description: `Number of occurrences of AS number`,
				},
				resource.Attribute{
					Name:        "prefix-list-direction",
					Description: `'both': both; 'receive': receive; 'send': send;`,
				},
				resource.Attribute{
					Name:        "graceful-restart",
					Description: `enable graceful-restart helper for this neighbor`,
				},
				resource.Attribute{
					Name:        "default-originate",
					Description: `Originate default route to this neighbor`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route-map to specify criteria to originate default (route-map name)`,
				},
				resource.Attribute{
					Name:        "distribute-list",
					Description: `Filter updates to/from this neighbor (IP standard/extended/named access list)`,
				},
				resource.Attribute{
					Name:        "distribute-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "filter-list",
					Description: `Establish BGP filters (AS path access-list name)`,
				},
				resource.Attribute{
					Name:        "filter-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "maximum-prefix",
					Description: `Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))`,
				},
				resource.Attribute{
					Name:        "maximum-prefix-thres",
					Description: `threshold-value, 1 to 100 percent`,
				},
				resource.Attribute{
					Name:        "next-hop-self",
					Description: `Disable the next hop calculation for this neighbor`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list",
					Description: `Filter updates to/from this neighbor (Name of a prefix list)`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "remove-private-as",
					Description: `Remove private AS number from outbound updates`,
				},
				resource.Attribute{
					Name:        "nbr-route-map",
					Description: `Apply route map to neighbor (Name of route map)`,
				},
				resource.Attribute{
					Name:        "nbr-rmap-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "send-community-val",
					Description: `'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `Allow inbound soft reconfiguration for this neighbor`,
				},
				resource.Attribute{
					Name:        "unsuppress-map",
					Description: `Route-map to selectively unsuppress suppressed routes (Name of route map)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Set default weight for routes from this neighbor`,
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
			Type:             "vthunder_thunder_router_bgp_address_family_ipv6_neighbor_peer_group_neighbor",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp address family ipv6 neighbor peer group neighbor resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "peer-group-neighbor",
					Description: `Specify a peer-group neighbor router`,
				},
				resource.Attribute{
					Name:        "peer-group",
					Description: `Neighbor tag`,
				},
				resource.Attribute{
					Name:        "activate",
					Description: `Enable the Address Family for this Neighbor`,
				},
				resource.Attribute{
					Name:        "allowas-in",
					Description: `Accept as-path with my AS present in it`,
				},
				resource.Attribute{
					Name:        "allowas-in-count",
					Description: `Number of occurrences of AS number`,
				},
				resource.Attribute{
					Name:        "prefix-list-direction",
					Description: `'both': both; 'receive': receive; 'send': send;`,
				},
				resource.Attribute{
					Name:        "default-originate",
					Description: `Originate default route to this neighbor`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route-map to specify criteria to originate default (route-map name)`,
				},
				resource.Attribute{
					Name:        "distribute-list",
					Description: `Filter updates to/from this neighbor (IP standard/extended/named access list)`,
				},
				resource.Attribute{
					Name:        "distribute-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "filter-list",
					Description: `Establish BGP filters (AS path access-list name)`,
				},
				resource.Attribute{
					Name:        "filter-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "maximum-prefix",
					Description: `Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))`,
				},
				resource.Attribute{
					Name:        "maximum-prefix-thres",
					Description: `threshold-value, 1 to 100 percent`,
				},
				resource.Attribute{
					Name:        "next-hop-self",
					Description: `Disable the next hop calculation for this neighbor`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list",
					Description: `Filter updates to/from this neighbor (Name of a prefix list)`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "remove-private-as",
					Description: `Remove private AS number from outbound updates`,
				},
				resource.Attribute{
					Name:        "nbr-route-map",
					Description: `Apply route map to neighbor (Name of route map)`,
				},
				resource.Attribute{
					Name:        "nbr-rmap-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "send-community-val",
					Description: `'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `Allow inbound soft reconfiguration for this neighbor`,
				},
				resource.Attribute{
					Name:        "unsuppress-map",
					Description: `Route-map to selectively unsuppress suppressed routes (Name of route map)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Set default weight for routes from this neighbor`,
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
			Type:             "vthunder_thunder_router_bgp_address_family_ipv6_neighbor_trunk_neighbor_ipv6",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp address family ipv6 neighbor trunk neighbor ipv6 resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "trunk-neighbor-ipv6",
					Description: `Specify a trunk unnumbered neighbor`,
				},
				resource.Attribute{
					Name:        "trunk",
					Description: `Trunk interface number`,
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
			Type:             "vthunder_thunder_router_bgp_address_family_ipv6_neighbor_ve_neighbor_ipv6",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp address family ipv6 neighbor ve neighbor ipv6 resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ve-neighbor-ipv6",
					Description: `Specify a VE unnumbered neighbor`,
				},
				resource.Attribute{
					Name:        "ve",
					Description: `Virtual ethernet interface number`,
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
			Type:             "vthunder_thunder_router_bgp_address_family_ipv6_network_ipv6_network",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp address family ipv6 network ipv6 network resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv6-network",
					Description: `Specify a ip address mask network to announce via BGP`,
				},
				resource.Attribute{
					Name:        "network-ipv6",
					Description: `Specify a network to announce via BGP`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route-map to modify the attributes (Name of the route map)`,
				},
				resource.Attribute{
					Name:        "backdoor",
					Description: `Specify a BGP backdoor route`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Network specific description (Up to 80 characters describing this network)`,
				},
				resource.Attribute{
					Name:        "comm-value",
					Description: `community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export`,
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
			Type:             "vthunder_thunder_router_bgp_address_family_ipv6_network_synchronization",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp address family ipv6 network synchronization resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "synchronization",
					Description: `help Perform IGP synchronization`,
				},
				resource.Attribute{
					Name:        "network-synchronization",
					Description: `Perform IGP synchronization`,
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
			Type:             "vthunder_thunder_router_bgp_address_family_ipv6_redistribute",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp address family ipv6 redistribute resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "redistribute",
					Description: `Redistribute information from another routing protocol`,
				},
				resource.Attribute{
					Name:        "connected",
					Description: `Connected`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route map reference (Pointer to route-map entries)`,
				},
				resource.Attribute{
					Name:        "floating-ip",
					Description: `Floating IP`,
				},
				resource.Attribute{
					Name:        "nat64",
					Description: `NAT64 Prefix`,
				},
				resource.Attribute{
					Name:        "nat-map",
					Description: `NAT MAP Prefix`,
				},
				resource.Attribute{
					Name:        "lw4o6",
					Description: `LW4O6 Prefix`,
				},
				resource.Attribute{
					Name:        "static-nat",
					Description: `Static NAT Prefix`,
				},
				resource.Attribute{
					Name:        "ip-nat",
					Description: `IP NAT`,
				},
				resource.Attribute{
					Name:        "ip-nat-list",
					Description: `IP NAT list`,
				},
				resource.Attribute{
					Name:        "isis",
					Description: `ISO IS-IS`,
				},
				resource.Attribute{
					Name:        "ospf",
					Description: `Open Shortest Path First (OSPF)`,
				},
				resource.Attribute{
					Name:        "rip",
					Description: `Routing Information Protocol (RIP)`,
				},
				resource.Attribute{
					Name:        "static",
					Description: `Static routes`,
				},
				resource.Attribute{
					Name:        "only-flagged",
					Description: `Selected Virtual IP (VIP)`,
				},
				resource.Attribute{
					Name:        "only-not-flagged",
					Description: `Only not flagged`,
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
			Type:             "vthunder_thunder_router_bgp_neighbor_ethernet_neighbor",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp neighbor ethernet neighbor resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ethernet-neighbor",
					Description: `Specify an ethernet unnumbered neighbor`,
				},
				resource.Attribute{
					Name:        "ethernet",
					Description: `Ethernet interface number`,
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
			Type:             "vthunder_thunder_router_bgp_neighbor_ipv4_neighbor",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp neighbor ipv4 neighbor resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv4-neighbor",
					Description: `Specify a ipv4 neighbor router`,
				},
				resource.Attribute{
					Name:        "neighbor-ipv4",
					Description: `Neighbor address`,
				},
				resource.Attribute{
					Name:        "nbr-remote-as",
					Description: `Specify AS number of BGP neighbor`,
				},
				resource.Attribute{
					Name:        "peer-group-name",
					Description: `Configure peer-group (peer-group name)`,
				},
				resource.Attribute{
					Name:        "activate",
					Description: `Enable the Address Family for this Neighbor`,
				},
				resource.Attribute{
					Name:        "advertisement-interval",
					Description: `Minimum interval between sending BGP routing updates (time in seconds)`,
				},
				resource.Attribute{
					Name:        "allowas-in",
					Description: `Accept as-path with my AS present in it`,
				},
				resource.Attribute{
					Name:        "allowas-in-count",
					Description: `Number of occurrences of AS number`,
				},
				resource.Attribute{
					Name:        "as-origination-interval",
					Description: `Minimum interval between sending AS-origination routing updates (time in seconds)`,
				},
				resource.Attribute{
					Name:        "dynamic",
					Description: `Advertise dynamic capability to this neighbor`,
				},
				resource.Attribute{
					Name:        "prefix-list-direction",
					Description: `'both': both; 'receive': receive; 'send': send;`,
				},
				resource.Attribute{
					Name:        "route-refresh",
					Description: `Advertise route-refresh capability to this neighbor`,
				},
				resource.Attribute{
					Name:        "graceful-restart",
					Description: `enable graceful-restart helper for this neighbor`,
				},
				resource.Attribute{
					Name:        "collide-established",
					Description: `Include Neighbor in Established State for Collision Detection`,
				},
				resource.Attribute{
					Name:        "default-originate",
					Description: `Originate default route to this neighbor`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route-map to specify criteria to originate default (route-map name)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Neighbor specific description (Up to 80 characters describing this neighbor)`,
				},
				resource.Attribute{
					Name:        "disallow-infinite-holdtime",
					Description: `BGP per neighbor disallow-infinite-holdtime`,
				},
				resource.Attribute{
					Name:        "distribute-list",
					Description: `Filter updates to/from this neighbor (IP standard/extended/named access list)`,
				},
				resource.Attribute{
					Name:        "distribute-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "acos-application-only",
					Description: `Send BGP update to ACOS application`,
				},
				resource.Attribute{
					Name:        "dont-capability-negotiate",
					Description: `Do not perform capability negotiation`,
				},
				resource.Attribute{
					Name:        "ebgp-multihop",
					Description: `Allow EBGP neighbors not on directly connected networks`,
				},
				resource.Attribute{
					Name:        "ebgp-multihop-hop-count",
					Description: `maximum hop count`,
				},
				resource.Attribute{
					Name:        "enforce-multihop",
					Description: `Enforce EBGP neighbors to perform multihop`,
				},
				resource.Attribute{
					Name:        "bfd",
					Description: `Bidirectional Forwarding Detection (BFD)`,
				},
				resource.Attribute{
					Name:        "multihop",
					Description: `Enable multihop`,
				},
				resource.Attribute{
					Name:        "key-id",
					Description: `Key ID`,
				},
				resource.Attribute{
					Name:        "key-type",
					Description: `'md5': md5; 'meticulous-md5': meticulous-md5; 'meticulous-sha1': meticulous-sha1; 'sha1': sha1; 'simple': simple; (Keyed MD5/Meticulous Keyed MD5/Meticulous Keyed SHA1/Keyed SHA1/Simple Password)`,
				},
				resource.Attribute{
					Name:        "bfd-value",
					Description: `Key String`,
				},
				resource.Attribute{
					Name:        "bfd-encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "filter-list",
					Description: `Establish BGP filters (AS path access-list name)`,
				},
				resource.Attribute{
					Name:        "filter-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "maximum-prefix",
					Description: `Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))`,
				},
				resource.Attribute{
					Name:        "maximum-prefix-thres",
					Description: `threshold-value, 1 to 100 percent`,
				},
				resource.Attribute{
					Name:        "next-hop-self",
					Description: `Disable the next hop calculation for this neighbor`,
				},
				resource.Attribute{
					Name:        "override-capability",
					Description: `Override capability negotiation result`,
				},
				resource.Attribute{
					Name:        "pass-value",
					Description: `Key String`,
				},
				resource.Attribute{
					Name:        "passive",
					Description: `Don't send open messages to this neighbor`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list",
					Description: `Filter updates to/from this neighbor (Name of a prefix list)`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "remove-private-as",
					Description: `Remove private AS number from outbound updates`,
				},
				resource.Attribute{
					Name:        "nbr-route-map",
					Description: `Apply route map to neighbor (Name of route map)`,
				},
				resource.Attribute{
					Name:        "nbr-rmap-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "send-community-val",
					Description: `'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `Allow inbound soft reconfiguration for this neighbor`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `Administratively shut down this neighbor`,
				},
				resource.Attribute{
					Name:        "strict-capability-match",
					Description: `Strict capability negotiation match`,
				},
				resource.Attribute{
					Name:        "timers-keepalive",
					Description: `Keepalive interval`,
				},
				resource.Attribute{
					Name:        "timers-holdtime",
					Description: `Holdtime`,
				},
				resource.Attribute{
					Name:        "connect",
					Description: `BGP connect timer`,
				},
				resource.Attribute{
					Name:        "unsuppress-map",
					Description: `Route-map to selectively unsuppress suppressed routes (Name of route map)`,
				},
				resource.Attribute{
					Name:        "update-source-ip",
					Description: `IP address`,
				},
				resource.Attribute{
					Name:        "update-source-ipv6",
					Description: `IPv6 address`,
				},
				resource.Attribute{
					Name:        "ethernet",
					Description: `Ethernet interface (Port number)`,
				},
				resource.Attribute{
					Name:        "loopback",
					Description: `Loopback interface (Port number)`,
				},
				resource.Attribute{
					Name:        "ve",
					Description: `Virtual ethernet interface (Virtual ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "trunk",
					Description: `Trunk interface (Trunk interface number)`,
				},
				resource.Attribute{
					Name:        "lif",
					Description: `Logical interface (Lif interface number)`,
				},
				resource.Attribute{
					Name:        "tunnel",
					Description: `Tunnel interface (Tunnel interface number)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Set default weight for routes from this neighbor`,
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
			Type:             "vthunder_thunder_router_bgp_neighbor_ipv6_neighbor",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp neighbor ipv6 neighbor resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv6-neighbor",
					Description: `Specify a ipv6 neighbor router`,
				},
				resource.Attribute{
					Name:        "neighbor-ipv6",
					Description: `Neighbor IPv6 address`,
				},
				resource.Attribute{
					Name:        "nbr-remote-as",
					Description: `Specify AS number of BGP neighbor`,
				},
				resource.Attribute{
					Name:        "peer-group-name",
					Description: `Configure peer-group (peer-group name)`,
				},
				resource.Attribute{
					Name:        "activate",
					Description: `Enable the Address Family for this Neighbor`,
				},
				resource.Attribute{
					Name:        "advertisement-interval",
					Description: `Minimum interval between sending BGP routing updates (time in seconds)`,
				},
				resource.Attribute{
					Name:        "allowas-in",
					Description: `Accept as-path with my AS present in it`,
				},
				resource.Attribute{
					Name:        "allowas-in-count",
					Description: `Number of occurrences of AS number`,
				},
				resource.Attribute{
					Name:        "as-origination-interval",
					Description: `Minimum interval between sending AS-origination routing updates (time in seconds)`,
				},
				resource.Attribute{
					Name:        "dynamic",
					Description: `Advertise dynamic capability to this neighbor`,
				},
				resource.Attribute{
					Name:        "prefix-list-direction",
					Description: `'both': both; 'receive': receive; 'send': send;`,
				},
				resource.Attribute{
					Name:        "route-refresh",
					Description: `Advertise route-refresh capability to this neighbor`,
				},
				resource.Attribute{
					Name:        "graceful-restart",
					Description: `enable graceful-restart helper for this neighbor`,
				},
				resource.Attribute{
					Name:        "extended-nexthop",
					Description: `Advertise extended-nexthop capability to this neighbor`,
				},
				resource.Attribute{
					Name:        "collide-established",
					Description: `Include Neighbor in Established State for Collision Detection`,
				},
				resource.Attribute{
					Name:        "default-originate",
					Description: `Originate default route to this neighbor`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route-map to specify criteria to originate default (route-map name)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Neighbor specific description (Up to 80 characters describing this neighbor)`,
				},
				resource.Attribute{
					Name:        "disallow-infinite-holdtime",
					Description: `BGP per neighbor disallow-infinite-holdtime`,
				},
				resource.Attribute{
					Name:        "distribute-list",
					Description: `Filter updates to/from this neighbor (IP standard/extended/named access list)`,
				},
				resource.Attribute{
					Name:        "distribute-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "acos-application-only",
					Description: `Send BGP update to ACOS application`,
				},
				resource.Attribute{
					Name:        "dont-capability-negotiate",
					Description: `Do not perform capability negotiation`,
				},
				resource.Attribute{
					Name:        "ebgp-multihop",
					Description: `Allow EBGP neighbors not on directly connected networks`,
				},
				resource.Attribute{
					Name:        "ebgp-multihop-hop-count",
					Description: `maximum hop count`,
				},
				resource.Attribute{
					Name:        "enforce-multihop",
					Description: `Enforce EBGP neighbors to perform multihop`,
				},
				resource.Attribute{
					Name:        "bfd",
					Description: `Bidirectional Forwarding Detection (BFD)`,
				},
				resource.Attribute{
					Name:        "multihop",
					Description: `Enable multihop`,
				},
				resource.Attribute{
					Name:        "key-id",
					Description: `Key ID`,
				},
				resource.Attribute{
					Name:        "key-type",
					Description: `'md5': md5; 'meticulous-md5': meticulous-md5; 'meticulous-sha1': meticulous-sha1; 'sha1': sha1; 'simple': simple; (Keyed MD5/Meticulous Keyed MD5/Meticulous Keyed SHA1/Keyed SHA1/Simple Password)`,
				},
				resource.Attribute{
					Name:        "bfd-value",
					Description: `Key String`,
				},
				resource.Attribute{
					Name:        "bfd-encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "filter-list",
					Description: `Establish BGP filters (AS path access-list name)`,
				},
				resource.Attribute{
					Name:        "filter-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "maximum-prefix",
					Description: `Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))`,
				},
				resource.Attribute{
					Name:        "maximum-prefix-thres",
					Description: `threshold-value, 1 to 100 percent`,
				},
				resource.Attribute{
					Name:        "next-hop-self",
					Description: `Disable the next hop calculation for this neighbor`,
				},
				resource.Attribute{
					Name:        "override-capability",
					Description: `Override capability negotiation result`,
				},
				resource.Attribute{
					Name:        "pass-value",
					Description: `Key String`,
				},
				resource.Attribute{
					Name:        "passive",
					Description: `Don't send open messages to this neighbor`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list",
					Description: `Filter updates to/from this neighbor (Name of a prefix list)`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "remove-private-as",
					Description: `Remove private AS number from outbound updates`,
				},
				resource.Attribute{
					Name:        "nbr-route-map",
					Description: `Apply route map to neighbor (Name of route map)`,
				},
				resource.Attribute{
					Name:        "nbr-rmap-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "send-community-val",
					Description: `'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `Allow inbound soft reconfiguration for this neighbor`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `Administratively shut down this neighbor`,
				},
				resource.Attribute{
					Name:        "strict-capability-match",
					Description: `Strict capability negotiation match`,
				},
				resource.Attribute{
					Name:        "timers-keepalive",
					Description: `Keepalive interval`,
				},
				resource.Attribute{
					Name:        "timers-holdtime",
					Description: `Holdtime`,
				},
				resource.Attribute{
					Name:        "connect",
					Description: `BGP connect timer`,
				},
				resource.Attribute{
					Name:        "unsuppress-map",
					Description: `Route-map to selectively unsuppress suppressed routes (Name of route map)`,
				},
				resource.Attribute{
					Name:        "update-source-ip",
					Description: `IP address`,
				},
				resource.Attribute{
					Name:        "update-source-ipv6",
					Description: `IPv6 address`,
				},
				resource.Attribute{
					Name:        "ethernet",
					Description: `Ethernet interface (Port number)`,
				},
				resource.Attribute{
					Name:        "loopback",
					Description: `Loopback interface (Port number)`,
				},
				resource.Attribute{
					Name:        "ve",
					Description: `Virtual ethernet interface (Virtual ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "trunk",
					Description: `Trunk interface (Trunk interface number)`,
				},
				resource.Attribute{
					Name:        "lif",
					Description: `Logical interface (Lif interface number)`,
				},
				resource.Attribute{
					Name:        "tunnel",
					Description: `Tunnel interface (Tunnel interface number)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Set default weight for routes from this neighbor`,
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
			Type:             "vthunder_thunder_router_bgp_neighbor_peer_group_neighbor",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp neighbor peer group neighbor resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "peer-group-neighbor",
					Description: `Specify a peer-group neighbor router`,
				},
				resource.Attribute{
					Name:        "peer-group",
					Description: `Neighbor tag`,
				},
				resource.Attribute{
					Name:        "peer-group-key",
					Description: `Configure peer-group`,
				},
				resource.Attribute{
					Name:        "peer-group-remote-as",
					Description: `Specify AS number of BGP neighbor`,
				},
				resource.Attribute{
					Name:        "activate",
					Description: `Enable the Address Family for this Neighbor`,
				},
				resource.Attribute{
					Name:        "advertisement-interval",
					Description: `Minimum interval between sending BGP routing updates (time in seconds)`,
				},
				resource.Attribute{
					Name:        "allowas-in",
					Description: `Accept as-path with my AS present in it`,
				},
				resource.Attribute{
					Name:        "allowas-in-count",
					Description: `Number of occurrences of AS number`,
				},
				resource.Attribute{
					Name:        "as-origination-interval",
					Description: `Minimum interval between sending AS-origination routing updates (time in seconds)`,
				},
				resource.Attribute{
					Name:        "dynamic",
					Description: `Advertise dynamic capability to this neighbor`,
				},
				resource.Attribute{
					Name:        "prefix-list-direction",
					Description: `'both': both; 'receive': receive; 'send': send;`,
				},
				resource.Attribute{
					Name:        "route-refresh",
					Description: `Advertise route-refresh capability to this neighbor`,
				},
				resource.Attribute{
					Name:        "extended-nexthop",
					Description: `Advertise extended-nexthop capability to this neighbor`,
				},
				resource.Attribute{
					Name:        "collide-established",
					Description: `Include Neighbor in Established State for Collision Detection`,
				},
				resource.Attribute{
					Name:        "default-originate",
					Description: `Originate default route to this neighbor`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route-map to specify criteria to originate default (route-map name)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Neighbor specific description (Up to 80 characters describing this neighbor)`,
				},
				resource.Attribute{
					Name:        "disallow-infinite-holdtime",
					Description: `BGP per neighbor disallow-infinite-holdtime`,
				},
				resource.Attribute{
					Name:        "distribute-list",
					Description: `Filter updates to/from this neighbor (IP standard/extended/named access list)`,
				},
				resource.Attribute{
					Name:        "distribute-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "dont-capability-negotiate",
					Description: `Do not perform capability negotiation`,
				},
				resource.Attribute{
					Name:        "ebgp-multihop",
					Description: `Allow EBGP neighbors not on directly connected networks`,
				},
				resource.Attribute{
					Name:        "ebgp-multihop-hop-count",
					Description: `maximum hop count`,
				},
				resource.Attribute{
					Name:        "enforce-multihop",
					Description: `Enforce EBGP neighbors to perform multihop`,
				},
				resource.Attribute{
					Name:        "bfd",
					Description: `Bidirectional Forwarding Detection (BFD)`,
				},
				resource.Attribute{
					Name:        "multihop",
					Description: `Enable multihop`,
				},
				resource.Attribute{
					Name:        "filter-list",
					Description: `Establish BGP filters (AS path access-list name)`,
				},
				resource.Attribute{
					Name:        "filter-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "maximum-prefix",
					Description: `Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))`,
				},
				resource.Attribute{
					Name:        "maximum-prefix-thres",
					Description: `threshold-value, 1 to 100 percent`,
				},
				resource.Attribute{
					Name:        "next-hop-self",
					Description: `Disable the next hop calculation for this neighbor`,
				},
				resource.Attribute{
					Name:        "override-capability",
					Description: `Override capability negotiation result`,
				},
				resource.Attribute{
					Name:        "pass-value",
					Description: `Key String`,
				},
				resource.Attribute{
					Name:        "passive",
					Description: `Don't send open messages to this neighbor`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list",
					Description: `Filter updates to/from this neighbor (Name of a prefix list)`,
				},
				resource.Attribute{
					Name:        "nbr-prefix-list-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "remove-private-as",
					Description: `Remove private AS number from outbound updates`,
				},
				resource.Attribute{
					Name:        "nbr-route-map",
					Description: `Apply route map to neighbor (Name of route map)`,
				},
				resource.Attribute{
					Name:        "nbr-rmap-direction",
					Description: `'in': in; 'out': out;`,
				},
				resource.Attribute{
					Name:        "send-community-val",
					Description: `'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `Allow inbound soft reconfiguration for this neighbor`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `Administratively shut down this neighbor`,
				},
				resource.Attribute{
					Name:        "strict-capability-match",
					Description: `Strict capability negotiation match`,
				},
				resource.Attribute{
					Name:        "timers-keepalive",
					Description: `Keepalive interval`,
				},
				resource.Attribute{
					Name:        "timers-holdtime",
					Description: `Holdtime`,
				},
				resource.Attribute{
					Name:        "connect",
					Description: `BGP connect timer`,
				},
				resource.Attribute{
					Name:        "unsuppress-map",
					Description: `Route-map to selectively unsuppress suppressed routes (Name of route map)`,
				},
				resource.Attribute{
					Name:        "update-source-ip",
					Description: `IP address`,
				},
				resource.Attribute{
					Name:        "update-source-ipv6",
					Description: `IPv6 address`,
				},
				resource.Attribute{
					Name:        "ethernet",
					Description: `Ethernet interface (Port number)`,
				},
				resource.Attribute{
					Name:        "loopback",
					Description: `Loopback interface (Port number)`,
				},
				resource.Attribute{
					Name:        "ve",
					Description: `Virtual ethernet interface (Virtual ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "trunk",
					Description: `Trunk interface (Trunk interface number)`,
				},
				resource.Attribute{
					Name:        "lif",
					Description: `Logical interface (Lif interface number)`,
				},
				resource.Attribute{
					Name:        "tunnel",
					Description: `Tunnel interface (Tunnel interface number)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Set default weight for routes from this neighbor`,
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
			Type:             "vthunder_thunder_router_bgp_neighbor_trunk_neighbor",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp neighbor trunk neighbor resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "trunk-neighbor",
					Description: `Specify a trunk unnumbered neighbor`,
				},
				resource.Attribute{
					Name:        "trunk",
					Description: `Trunk interface number`,
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
			Type:             "vthunder_thunder_router_bgp_neighbor_ve_neighbor",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp neighbor ve neighbor resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ve-neighbor",
					Description: `Specify a VE unnumbered neighbor`,
				},
				resource.Attribute{
					Name:        "ve",
					Description: `Virtual ethernet interface number`,
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
			Type:             "vthunder_thunder_router_bgp_network_ip_cidr",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp network ip cidr resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip-cidr",
					Description: `Specify a ip network to announce via BGP`,
				},
				resource.Attribute{
					Name:        "network-ipv4-cidr",
					Description: `Specify network mask`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route-map to modify the attributes (Name of the route map)`,
				},
				resource.Attribute{
					Name:        "backdoor",
					Description: `Specify a BGP backdoor route`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Network specific description (Up to 80 characters describing this network)`,
				},
				resource.Attribute{
					Name:        "comm-value",
					Description: `community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export`,
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
			Type:             "vthunder_thunder_router_bgp_network_synchronization",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp network synchronization resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "synchronization",
					Description: `help Perform IGP synchronization`,
				},
				resource.Attribute{
					Name:        "network-synchronization",
					Description: `Perform IGP synchronization`,
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
			Type:             "vthunder_thunder_router_bgp_redistribute",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router bgp redistribute resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "redistribute",
					Description: `Redistribute information from another routing protocol`,
				},
				resource.Attribute{
					Name:        "connected",
					Description: `Connected`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route map reference (Pointer to route-map entries)`,
				},
				resource.Attribute{
					Name:        "floating-ip",
					Description: `Floating IP`,
				},
				resource.Attribute{
					Name:        "lw4o6",
					Description: `LW4O6 Prefix`,
				},
				resource.Attribute{
					Name:        "static-nat",
					Description: `Static NAT Prefix`,
				},
				resource.Attribute{
					Name:        "ip-nat",
					Description: `IP NAT`,
				},
				resource.Attribute{
					Name:        "ip-nat-list",
					Description: `IP NAT list`,
				},
				resource.Attribute{
					Name:        "isis",
					Description: `ISO IS-IS`,
				},
				resource.Attribute{
					Name:        "ospf",
					Description: `Open Shortest Path First (OSPF)`,
				},
				resource.Attribute{
					Name:        "rip",
					Description: `Routing Information Protocol (RIP)`,
				},
				resource.Attribute{
					Name:        "static",
					Description: `Static routes`,
				},
				resource.Attribute{
					Name:        "nat-map",
					Description: `NAT MAP Prefix`,
				},
				resource.Attribute{
					Name:        "only-flagged",
					Description: `Selected Virtual IP (VIP)`,
				},
				resource.Attribute{
					Name:        "only-not-flagged",
					Description: `Only not flagged`,
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
			Type:             "vthunder_thunder_router_ospf",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router ospf resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ospf",
					Description: `Open Shortest Path First (OSPF)`,
				},
				resource.Attribute{
					Name:        "process-id",
					Description: `OSPF process ID`,
				},
				resource.Attribute{
					Name:        "auto-cost-reference-bandwidth",
					Description: `Use reference bandwidth method to assign OSPF cost (The reference bandwidth in terms of Mbits per second)`,
				},
				resource.Attribute{
					Name:        "bfd-all-interfaces",
					Description: `Enable BFD on all interfaces`,
				},
				resource.Attribute{
					Name:        "rfc1583-compatible",
					Description: `Compatible with RFC 1583`,
				},
				resource.Attribute{
					Name:        "default-metric",
					Description: `Set metric of redistributed routes (Default metric)`,
				},
				resource.Attribute{
					Name:        "distance-value",
					Description: `OSPF Administrative distance`,
				},
				resource.Attribute{
					Name:        "distance-external",
					Description: `External routes (Distance for external routes)`,
				},
				resource.Attribute{
					Name:        "distance-inter-area",
					Description: `Inter-area routes (Distance for inter-area routes)`,
				},
				resource.Attribute{
					Name:        "distance-intra-area",
					Description: `Intra-area routes (Distance for intra-area routes)`,
				},
				resource.Attribute{
					Name:        "di-type",
					Description: `'lw4o6': LW4O6 Prefix; 'floating-ip': Floating IP; 'ip-nat': IP NAT; 'ip-nat-list': IP NAT list; 'static-nat': Static NAT; 'vip': Only not flagged Virtual IP (VIP); 'vip-only-flagged': Selected Virtual IP (VIP);`,
				},
				resource.Attribute{
					Name:        "di-area-ipv4",
					Description: `OSPF area ID as a IP address format`,
				},
				resource.Attribute{
					Name:        "di-area-num",
					Description: `OSPF area ID as a decimal value`,
				},
				resource.Attribute{
					Name:        "di-cost",
					Description: `Cost of route`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `OSPF router-id in IPv4 address format`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `'in': Filter incoming routing updates; 'out': Filter outgoing routing updates;`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'lw4o6': LW4O6 Prefix; 'ip-nat': IP NAT; 'ip-nat-list': IP NAT list; 'static-nat': Static NAT; 'isis': ISO IS-IS; 'ospf': Open Shortest Path First (OSPF); 'rip': Routing Information Protocol (RIP); 'static': Static routes;`,
				},
				resource.Attribute{
					Name:        "ospf-id",
					Description: `OSPF process ID`,
				},
				resource.Attribute{
					Name:        "option",
					Description: `'advertise': Advertise this range (default); 'not-advertise': DoNotAdvertise this range;`,
				},
				resource.Attribute{
					Name:        "extra-cost",
					Description: `The extra cost value`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `Group (Group ID)`,
				},
				resource.Attribute{
					Name:        "host-address",
					Description: `Host address`,
				},
				resource.Attribute{
					Name:        "area-ipv4",
					Description: `OSPF area ID in IP address format`,
				},
				resource.Attribute{
					Name:        "area-num",
					Description: `OSPF area ID as a decimal value`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `OSPF cost for point-to-multipoint neighbor (Metric)`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `'detail': Log changes in adjacency state; 'disable': Disable logging;`,
				},
				resource.Attribute{
					Name:        "max-concurrent-dd",
					Description: `Maximum number allowed to process DD concurrently (Number of DD process)`,
				},
				resource.Attribute{
					Name:        "maximum-area",
					Description: `Maximum number of non-backbone areas (OSPF area limit)`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Neighbor address`,
				},
				resource.Attribute{
					Name:        "poll-interval",
					Description: `OSPF dead-router polling interval (Seconds)`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `OSPF priority of non-broadcast neighbor`,
				},
				resource.Attribute{
					Name:        "network-ipv4",
					Description: `Network number`,
				},
				resource.Attribute{
					Name:        "network-ipv4-mask",
					Description: `OSPF wild card bits`,
				},
				resource.Attribute{
					Name:        "network-ipv4-cidr",
					Description: `OSPF network prefix`,
				},
				resource.Attribute{
					Name:        "network-area-ipv4",
					Description: `OSPF area ID in IP address format`,
				},
				resource.Attribute{
					Name:        "network-area-num",
					Description: `OSPF area ID as a decimal value`,
				},
				resource.Attribute{
					Name:        "instance-value",
					Description: `Instance ID`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `Maximum number of LSAs`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `'hard': Hard limit: Instance will be shutdown if exceeded; 'soft': Soft limit: Warning will be given if exceeded;`,
				},
				resource.Attribute{
					Name:        "db-external",
					Description: `Maximum number of LSAs`,
				},
				resource.Attribute{
					Name:        "recovery-time",
					Description: `Time to recover (0 not recover)`,
				},
				resource.Attribute{
					Name:        "loopback",
					Description: `Loopback interface (Port number)`,
				},
				resource.Attribute{
					Name:        "loopback-address",
					Description: `Address of Interface`,
				},
				resource.Attribute{
					Name:        "trunk",
					Description: `Trunk interface (Trunk interface number)`,
				},
				resource.Attribute{
					Name:        "trunk-address",
					Description: `Address of Interface`,
				},
				resource.Attribute{
					Name:        "ve",
					Description: `Virtual ethernet interface (Virtual ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "ve-address",
					Description: `Address of Interface`,
				},
				resource.Attribute{
					Name:        "tunnel",
					Description: `Tunnel interface (Tunnel interface number)`,
				},
				resource.Attribute{
					Name:        "tunnel-address",
					Description: `Address of Interface`,
				},
				resource.Attribute{
					Name:        "lif",
					Description: `Logical interface (Lif interface number)`,
				},
				resource.Attribute{
					Name:        "lif-address",
					Description: `Address of Interface`,
				},
				resource.Attribute{
					Name:        "ethernet",
					Description: `Ethernet interface (Port number)`,
				},
				resource.Attribute{
					Name:        "eth-address",
					Description: `Address of Interface`,
				},
				resource.Attribute{
					Name:        "summary-address",
					Description: `Configure IP address summaries (Summary prefix)`,
				},
				resource.Attribute{
					Name:        "not-advertise",
					Description: `Suppress routes that match the prefix`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Set tag for routes redistributed into OSPF (32-bit tag value)`,
				},
				resource.Attribute{
					Name:        "min-delay",
					Description: `Minimum delay between receiving a change to SPF calculation in milliseconds`,
				},
				resource.Attribute{
					Name:        "max-delay",
					Description: `Maximum delay between receiving a change to SPF calculation in milliseconds`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "user-tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "originate",
					Description: `Distribute a default route`,
				},
				resource.Attribute{
					Name:        "always",
					Description: `Always advertise default route`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `OSPF default metric (OSPF metric)`,
				},
				resource.Attribute{
					Name:        "metric-type",
					Description: `'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route map reference (Pointer to route-map entries)`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Enable authentication`,
				},
				resource.Attribute{
					Name:        "message-digest",
					Description: `Use message-digest authentication`,
				},
				resource.Attribute{
					Name:        "filter-list",
					Description: `Filter networks between OSPF areas`,
				},
				resource.Attribute{
					Name:        "acl-name",
					Description: `Filter networks by access-list (Name of an access-list)`,
				},
				resource.Attribute{
					Name:        "acl-direction",
					Description: `'in': Filter networks sent to this area; 'out': Filter networks sent from this area;`,
				},
				resource.Attribute{
					Name:        "plist-name",
					Description: `Filter networks by prefix-list (Name of an IP prefix-list)`,
				},
				resource.Attribute{
					Name:        "plist-direction",
					Description: `'in': Filter networks sent to this area; 'out': Filter networks sent from this area;`,
				},
				resource.Attribute{
					Name:        "nssa",
					Description: `Specify a NSSA area`,
				},
				resource.Attribute{
					Name:        "no-redistribution",
					Description: `No redistribution into this NSSA area`,
				},
				resource.Attribute{
					Name:        "no-summary",
					Description: `Do not inject inter-area routes into area`,
				},
				resource.Attribute{
					Name:        "translator-role",
					Description: `'always': Translate always; 'candidate': Candidate for translator (default); 'never': Do not translate;`,
				},
				resource.Attribute{
					Name:        "default-information-originate",
					Description: `Originate Type 7 default into NSSA area`,
				},
				resource.Attribute{
					Name:        "default-cost",
					Description: `Set the summary-default cost of a NSSA or stub area (Stub's advertised default summary cost)`,
				},
				resource.Attribute{
					Name:        "area-range-prefix",
					Description: `Area range for IPv4 prefix`,
				},
				resource.Attribute{
					Name:        "shortcut",
					Description: `'default': Set default shortcutting behavior; 'disable': Disable shortcutting through the area; 'enable': Enable shortcutting through the area;`,
				},
				resource.Attribute{
					Name:        "stub",
					Description: `Configure OSPF area as stub`,
				},
				resource.Attribute{
					Name:        "virtual-link-ip-addr",
					Description: `ID (IP addr) associated with virtual link neighbor`,
				},
				resource.Attribute{
					Name:        "bfd",
					Description: `Bidirectional Forwarding Detection (BFD)`,
				},
				resource.Attribute{
					Name:        "hello-interval",
					Description: `Hello packet interval (Seconds)`,
				},
				resource.Attribute{
					Name:        "dead-interval",
					Description: `Dead router detection time (Seconds)`,
				},
				resource.Attribute{
					Name:        "retransmit-interval",
					Description: `LSA retransmit interval (Seconds)`,
				},
				resource.Attribute{
					Name:        "transmit-delay",
					Description: `LSA transmission delay (Seconds)`,
				},
				resource.Attribute{
					Name:        "virtual-link-authentication",
					Description: `Enable authentication`,
				},
				resource.Attribute{
					Name:        "virtual-link-auth-type",
					Description: `'message-digest': Use message-digest authentication; 'null': Use null authentication;`,
				},
				resource.Attribute{
					Name:        "authentication-key",
					Description: `Set authentication key (Authentication key (8 chars))`,
				},
				resource.Attribute{
					Name:        "message-digest-key",
					Description: `Set message digest key (Key ID)`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `Use MD5 algorithm (Authentication key (16 chars))`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'lw4o6': LW4O6 Prefix; 'nat-map': NAT MAP Prefix; 'static-nat': Static NAT; 'isis': ISO IS-IS; 'rip': Routing Information Protocol (RIP); 'static': Static routes;`,
				},
				resource.Attribute{
					Name:        "metric-ospf",
					Description: `OSPF default metric (OSPF metric)`,
				},
				resource.Attribute{
					Name:        "metric-type-ospf",
					Description: `'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;`,
				},
				resource.Attribute{
					Name:        "route-map-ospf",
					Description: `Route map reference (Pointer to route-map entries)`,
				},
				resource.Attribute{
					Name:        "tag-ospf",
					Description: `Set tag for routes redistributed into OSPF (32-bit tag value)`,
				},
				resource.Attribute{
					Name:        "ip-nat",
					Description: `IP-NAT`,
				},
				resource.Attribute{
					Name:        "metric-ip-nat",
					Description: `OSPF default metric (OSPF metric)`,
				},
				resource.Attribute{
					Name:        "metric-type-ip-nat",
					Description: `'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;`,
				},
				resource.Attribute{
					Name:        "route-map-ip-nat",
					Description: `Route map reference (Pointer to route-map entries)`,
				},
				resource.Attribute{
					Name:        "tag-ip-nat",
					Description: `Set tag for routes redistributed into OSPF (32-bit tag value)`,
				},
				resource.Attribute{
					Name:        "ip-nat-prefix",
					Description: `Address`,
				},
				resource.Attribute{
					Name:        "ip-nat-floating-IP-forward",
					Description: `Floating-IP as forward address`,
				},
				resource.Attribute{
					Name:        "type-vip",
					Description: `'only-flagged': Selected Virtual IP (VIP); 'only-not-flagged': Only not flagged;`,
				},
				resource.Attribute{
					Name:        "metric-vip",
					Description: `OSPF default metric (OSPF metric)`,
				},
				resource.Attribute{
					Name:        "metric-type-vip",
					Description: `'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;`,
				},
				resource.Attribute{
					Name:        "route-map-vip",
					Description: `Route map reference (Pointer to route-map entries)`,
				},
				resource.Attribute{
					Name:        "tag-vip",
					Description: `Set tag for routes redistributed into OSPF (32-bit tag value)`,
				},
				resource.Attribute{
					Name:        "vip-address",
					Description: `Address`,
				},
				resource.Attribute{
					Name:        "vip-floating-IP-forward",
					Description: `Floating-IP as forward address`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_router_ospf_area",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router ospf area resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "area",
					Description: `OSPF area parameters`,
				},
				resource.Attribute{
					Name:        "area-ipv4",
					Description: `OSPF area ID in IP address format`,
				},
				resource.Attribute{
					Name:        "area-num",
					Description: `OSPF area ID as a decimal value`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Enable authentication`,
				},
				resource.Attribute{
					Name:        "message-digest",
					Description: `Use message-digest authentication`,
				},
				resource.Attribute{
					Name:        "filter-list",
					Description: `Filter networks between OSPF areas`,
				},
				resource.Attribute{
					Name:        "acl-name",
					Description: `Filter networks by access-list (Name of an access-list)`,
				},
				resource.Attribute{
					Name:        "acl-direction",
					Description: `'in': Filter networks sent to this area; 'out': Filter networks sent from this area;`,
				},
				resource.Attribute{
					Name:        "plist-name",
					Description: `Filter networks by prefix-list (Name of an IP prefix-list)`,
				},
				resource.Attribute{
					Name:        "plist-direction",
					Description: `'in': Filter networks sent to this area; 'out': Filter networks sent from this area;`,
				},
				resource.Attribute{
					Name:        "nssa",
					Description: `Specify a NSSA area`,
				},
				resource.Attribute{
					Name:        "no-redistribution",
					Description: `No redistribution into this NSSA area`,
				},
				resource.Attribute{
					Name:        "no-summary",
					Description: `Do not inject inter-area routes into area`,
				},
				resource.Attribute{
					Name:        "translator-role",
					Description: `'always': Translate always; 'candidate': Candidate for translator (default); 'never': Do not translate;`,
				},
				resource.Attribute{
					Name:        "default-information-originate",
					Description: `Originate Type 7 default into NSSA area`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `OSPF default metric (OSPF metric)`,
				},
				resource.Attribute{
					Name:        "metric-type",
					Description: `OSPF metric type (OSPF metric type for default routes)`,
				},
				resource.Attribute{
					Name:        "default-cost",
					Description: `Set the summary-default cost of a NSSA or stub area (Stub's advertised default summary cost)`,
				},
				resource.Attribute{
					Name:        "area-range-prefix",
					Description: `Area range for IPv4 prefix`,
				},
				resource.Attribute{
					Name:        "option",
					Description: `'advertise': Advertise this range (default); 'not-advertise': DoNotAdvertise this range;`,
				},
				resource.Attribute{
					Name:        "shortcut",
					Description: `'default': Set default shortcutting behavior; 'disable': Disable shortcutting through the area; 'enable': Enable shortcutting through the area;`,
				},
				resource.Attribute{
					Name:        "stub",
					Description: `Configure OSPF area as stub`,
				},
				resource.Attribute{
					Name:        "virtual-link-ip-addr",
					Description: `ID (IP addr) associated with virtual link neighbor`,
				},
				resource.Attribute{
					Name:        "bfd",
					Description: `Bidirectional Forwarding Detection (BFD)`,
				},
				resource.Attribute{
					Name:        "hello-interval",
					Description: `Hello packet interval (Seconds)`,
				},
				resource.Attribute{
					Name:        "dead-interval",
					Description: `Dead router detection time (Seconds)`,
				},
				resource.Attribute{
					Name:        "retransmit-interval",
					Description: `LSA retransmit interval (Seconds)`,
				},
				resource.Attribute{
					Name:        "transmit-delay",
					Description: `LSA transmission delay (Seconds)`,
				},
				resource.Attribute{
					Name:        "virtual-link-authentication",
					Description: `Enable authentication`,
				},
				resource.Attribute{
					Name:        "virtual-link-auth-type",
					Description: `'message-digest': Use message-digest authentication; 'null': Use null authentication;`,
				},
				resource.Attribute{
					Name:        "authentication-key",
					Description: `Set authentication key (Authentication key (8 chars))`,
				},
				resource.Attribute{
					Name:        "message-digest-key",
					Description: `Set message digest key (Key ID)`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `Use MD5 algorithm (Authentication key (16 chars))`,
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
			Type:             "vthunder_thunder_router_ospf_default_information",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router ospf default information resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default-information",
					Description: `Control distribution of default information`,
				},
				resource.Attribute{
					Name:        "originate",
					Description: `Distribute a default route`,
				},
				resource.Attribute{
					Name:        "always",
					Description: `Always advertise default route`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `OSPF default metric (OSPF metric)`,
				},
				resource.Attribute{
					Name:        "metric-type",
					Description: `OSPF metric type for default routes`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route map reference (Pointer to route-map entries)`,
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
			Type:             "vthunder_thunder_router_ospf_redistribute",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder router ospf redistribute resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "redistribute",
					Description: `Redistribute information from another routing protocol`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'lw4o6': LW4O6 Prefix; 'nat-map': NAT MAP Prefix; 'static-nat': Static NAT; 'isis': ISO IS-IS; 'rip': Routing Information Protocol (RIP); 'static': Static routes;`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `OSPF default metric (OSPF metric)`,
				},
				resource.Attribute{
					Name:        "metric-type",
					Description: `'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;`,
				},
				resource.Attribute{
					Name:        "route-map",
					Description: `Route map reference (Pointer to route-map entries)`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Set tag for routes redistributed into OSPF (32-bit tag value)`,
				},
				resource.Attribute{
					Name:        "ospf",
					Description: `Open Shortest Path First (OSPF)`,
				},
				resource.Attribute{
					Name:        "process-id",
					Description: `OSPF process ID`,
				},
				resource.Attribute{
					Name:        "metric-ospf",
					Description: `OSPF default metric (OSPF metric)`,
				},
				resource.Attribute{
					Name:        "metric-type-ospf",
					Description: `'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;`,
				},
				resource.Attribute{
					Name:        "route-map-ospf",
					Description: `Route map reference (Pointer to route-map entries)`,
				},
				resource.Attribute{
					Name:        "tag-ospf",
					Description: `Set tag for routes redistributed into OSPF (32-bit tag value)`,
				},
				resource.Attribute{
					Name:        "ip-nat",
					Description: `IP-NAT`,
				},
				resource.Attribute{
					Name:        "metric-ip-nat",
					Description: `OSPF default metric (OSPF metric)`,
				},
				resource.Attribute{
					Name:        "metric-type-ip-nat",
					Description: `'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;`,
				},
				resource.Attribute{
					Name:        "route-map-ip-nat",
					Description: `Route map reference (Pointer to route-map entries)`,
				},
				resource.Attribute{
					Name:        "tag-ip-nat",
					Description: `Set tag for routes redistributed into OSPF (32-bit tag value)`,
				},
				resource.Attribute{
					Name:        "ip-nat-prefix",
					Description: `Address`,
				},
				resource.Attribute{
					Name:        "ip-nat-floating-IP-forward",
					Description: `Floating-IP as forward address`,
				},
				resource.Attribute{
					Name:        "type-vip",
					Description: `'only-flagged': Selected Virtual IP (VIP); 'only-not-flagged': Only not flagged;`,
				},
				resource.Attribute{
					Name:        "metric-vip",
					Description: `OSPF default metric (OSPF metric)`,
				},
				resource.Attribute{
					Name:        "metric-type-vip",
					Description: `'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;`,
				},
				resource.Attribute{
					Name:        "route-map-vip",
					Description: `Route map reference (Pointer to route-map entries)`,
				},
				resource.Attribute{
					Name:        "tag-vip",
					Description: `Set tag for routes redistributed into OSPF (32-bit tag value)`,
				},
				resource.Attribute{
					Name:        "vip-address",
					Description: `Address`,
				},
				resource.Attribute{
					Name:        "vip-floating-IP-forward",
					Description: `Floating-IP as forward address`,
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
			Type:             "vthunder_thunder_server",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder server resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `‘enable’: enable; ‘disable’: disable; ‘disable-with-health-check’: disable port, but health check work;`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `Connection Limit`,
				},
				resource.Attribute{
					Name:        "conn_resume",
					Description: `Connection Resume`,
				},
				resource.Attribute{
					Name:        "extended_stats",
					Description: `Enable extended statistics on real server port`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `External IP address for NAT of GSLB`,
				},
				resource.Attribute{
					Name:        "fqdn_name",
					Description: `Server hostname`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Health Check (Monitor Name)`,
				},
				resource.Attribute{
					Name:        "health_check_disable",
					Description: `Disable health check`,
				},
				resource.Attribute{
					Name:        "health_check_shared",
					Description: `Health Check Monitor (Health monitor name)`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `IP Address`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `IPv6 address Mapping of GSLB`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Server Name`,
				},
				resource.Attribute{
					Name:        "no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "resolve_as",
					Description: `‘resolve-to-ipv4’: Use A Query only to resolve FQDN; ‘resolve-to-ipv6’: Use AAAA Query only to resolve FQDN; ‘resolve-to-ipv4-and-ipv6’: Use A as well as AAAA Query to resolve FQDN;`,
				},
				resource.Attribute{
					Name:        "server_ipv6_addr",
					Description: `IPV6 address`,
				},
				resource.Attribute{
					Name:        "shared_partition_health_check",
					Description: `Reference a health-check from shared partition`,
				},
				resource.Attribute{
					Name:        "slow_start",
					Description: `Slowly ramp up the connection number after server is up (start from 128, then double every 10 sec till 4096)`,
				},
				resource.Attribute{
					Name:        "spoofing_cache",
					Description: `This server is a spoofing cache`,
				},
				resource.Attribute{
					Name:        "stats_data_action",
					Description: `‘stats-data-enable’: Enable statistical data collection for real server port; ‘stats-data-disable’: Disable statistical data collection for real server port;`,
				},
				resource.Attribute{
					Name:        "template_server",
					Description: `Server template (Server template name)`,
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
					Name:        "weight",
					Description: `Port Weight (Connection Weight)`,
				},
				resource.Attribute{
					Name:        "follow_port_protocol",
					Description: `‘tcp’: TCP Port; ‘udp’: UDP Port;`,
				},
				resource.Attribute{
					Name:        "health_check_follow_port",
					Description: `Specify which port to follow for health status (Port Number)`,
				},
				resource.Attribute{
					Name:        "no_ssl",
					Description: `No SSL`,
				},
				resource.Attribute{
					Name:        "port_number",
					Description: `Port Number`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `‘tcp’: TCP Port; ‘udp’: UDP Port;`,
				},
				resource.Attribute{
					Name:        "range",
					Description: `Port range (Port range value - used for vip-to-rport-mapping)`,
				},
				resource.Attribute{
					Name:        "rport_health_check_shared",
					Description: `Health Check (Monitor Name)`,
				},
				resource.Attribute{
					Name:        "shared_rport_health_check",
					Description: `Reference a health-check from shared partition`,
				},
				resource.Attribute{
					Name:        "support_http2",
					Description: `Starting HTTP/2 with Prior Knowledge`,
				},
				resource.Attribute{
					Name:        "template_port",
					Description: `Port template (Port template name)`,
				},
				resource.Attribute{
					Name:        "template_server_ssl",
					Description: `Server side SSL template (Server side SSL Name)`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘total-conn’: Total connections; ‘fwd-pkt’: Forward packets; ‘rev-pkt’: Reverse packets; ‘peak-conn’: Peak connections; ‘total_req’: Total Requests; ‘total_req_succ’: Total requests succ; ‘curr_ssl_conn’: Current SSL connections; ‘total_ssl_conn’: Total SSL connections;`,
				},
				resource.Attribute{
					Name:        "alternate",
					Description: `Alternate Server (Alternate Server Number)`,
				},
				resource.Attribute{
					Name:        "alternate_name",
					Description: `Alternate Name`,
				},
				resource.Attribute{
					Name:        "alternate_server_port",
					Description: `Port (Alternate Server Port Value)`,
				},
				resource.Attribute{
					Name:        "service_principal_name",
					Description: `Service Principal Name (Kerberos principal name)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_service_group",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder service group resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "backup_server_event_log",
					Description: `Send log info on back up server events`,
				},
				resource.Attribute{
					Name:        "conn_rate",
					Description: `Dynamically enable stateless method by conn-rate (Rate to trigger stateless method(conn/sec))`,
				},
				resource.Attribute{
					Name:        "conn_rate_duration",
					Description: `Period that trigger condition consistently happens(seconds)`,
				},
				resource.Attribute{
					Name:        "conn_rate_grace_period",
					Description: `Define the grace period during transition (Define the grace period during transition(seconds))`,
				},
				resource.Attribute{
					Name:        "conn_rate_log",
					Description: `Send log if transition happens`,
				},
				resource.Attribute{
					Name:        "conn_rate_revert_duration",
					Description: `Period that revert condition consistently happens(seconds)`,
				},
				resource.Attribute{
					Name:        "conn_revert_rate",
					Description: `Rate to revert to statelful method (conn/sec)`,
				},
				resource.Attribute{
					Name:        "extended_stats",
					Description: `Enable extended statistics on service group`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Health Check (Monitor Name)`,
				},
				resource.Attribute{
					Name:        "health_check_disable",
					Description: `Disable health check`,
				},
				resource.Attribute{
					Name:        "l4_session_revert_duration",
					Description: `Period that revert condition consistently happens(seconds)`,
				},
				resource.Attribute{
					Name:        "l4_session_usage",
					Description: `Dynamically enable stateless method by session usage (Usage to trigger stateless method)`,
				},
				resource.Attribute{
					Name:        "l4_session_usage_duration",
					Description: `Period that trigger condition consistently happens(seconds)`,
				},
				resource.Attribute{
					Name:        "l4_session_usage_grace_period",
					Description: `Define the grace period during transition (Define the grace period during transition(seconds))`,
				},
				resource.Attribute{
					Name:        "l4_session_usage_log",
					Description: `Send log if transition happens`,
				},
				resource.Attribute{
					Name:        "l4_session_usage_revert_rate",
					Description: `Usage to revert to statelful method`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `‘dst-ip-hash’: Load-balancing based on only Dst IP and Port hash; ‘dst-ip-only-hash’: Load-balancing based on only Dst IP hash; ‘fastest-response’: Fastest response time on service port level; ‘least-request’: Least request on service port level; ‘src-ip-hash’: Load-balancing based on only Src IP and Port hash; ‘src-ip-only-hash’: Load-balancing based on only Src IP hash; ‘weighted-rr’: Weighted round robin on server level; ‘round-robin’: Round robin on server level; ‘round-robin-strict’: Strict mode round robin on server level; ‘odd-even-hash’: odd/even hash based of client src-ip;`,
				},
				resource.Attribute{
					Name:        "lc_method",
					Description: `‘least-connection’: Least connection on server level; ‘service-least-connection’: Least connection on service port level; ‘weighted-least-connection’: Weighted least connection on server level; ‘service-weighted-least-connection’: Weighted least connection on service port level;`,
				},
				resource.Attribute{
					Name:        "min_active_member",
					Description: `Minimum Active Member Per Priority (Minimum Active Member before Action)`,
				},
				resource.Attribute{
					Name:        "min_active_member_action",
					Description: `‘dynamic-priority’: dynamic change member priority to met the min-active-member requirement; ‘skip-pri-set’: Skip Current Priority Set If Min not met;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Member name`,
				},
				resource.Attribute{
					Name:        "priority_affinity",
					Description: `Priority affinity. Persist to the same priority if possible.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `‘tcp’: TCP LB service; ‘udp’: UDP LB service;`,
				},
				resource.Attribute{
					Name:        "pseudo_round_robin",
					Description: `PRR, select the oldest node for sub-select`,
				},
				resource.Attribute{
					Name:        "report_delay",
					Description: `Reporting frequency (in minutes)`,
				},
				resource.Attribute{
					Name:        "reset_on_server_selection_fail",
					Description: `Send reset to client if server selection fails`,
				},
				resource.Attribute{
					Name:        "reset_priority_affinity",
					Description: `Reset`,
				},
				resource.Attribute{
					Name:        "rpt_ext_server",
					Description: `Report top 10 fastest/slowest servers`,
				},
				resource.Attribute{
					Name:        "sample_rsp_time",
					Description: `sample server response time`,
				},
				resource.Attribute{
					Name:        "shared_partition_policy_template",
					Description: `Reference a policy template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_svcgrp_health_check",
					Description: `Reference a health-check from shared partition`,
				},
				resource.Attribute{
					Name:        "stateless_auto_switch",
					Description: `Enable auto stateless method`,
				},
				resource.Attribute{
					Name:        "stateless_lb_method",
					Description: `‘stateless-dst-ip-hash’: Stateless load-balancing based on Dst IP and Dst port hash; ‘stateless-per-pkt-round-robin’: Stateless load-balancing using per-packet round-robin; ‘stateless-src-dst-ip-hash’: Stateless load-balancing based on IP and port hash for both Src and Dst; ‘stateless-src-dst-ip-only-hash’: Stateless load-balancing based on only IP hash for both Src and Dst; ‘stateless-src-ip-hash’: Stateless load-balancing based on Src IP and Src port hash; ‘stateless-src-ip-only-hash’: Stateless load-balancing based on only Src IP hash;`,
				},
				resource.Attribute{
					Name:        "stateless_lb_method2",
					Description: `‘stateless-dst-ip-hash’: Stateless load-balancing based on Dst IP and Dst port hash; ‘stateless-per-pkt-round-robin’: Stateless load-balancing using per-packet round-robin; ‘stateless-src-dst-ip-hash’: Stateless load-balancing based on IP and port hash for both Src and Dst; ‘stateless-src-dst-ip-only-hash’: Stateless load-balancing based on only IP hash for both Src and Dst; ‘stateless-src-ip-hash’: Stateless load-balancing based on Src IP and Src port hash; ‘stateless-src-ip-only-hash’: Stateless load-balancing based on only Src IP hash;`,
				},
				resource.Attribute{
					Name:        "stats_data_action",
					Description: `‘stats-data-enable’: Enable statistical data collection for service group; ‘stats-data-disable’: Disable statistical data collection for service group;`,
				},
				resource.Attribute{
					Name:        "strict_select",
					Description: `strict selection`,
				},
				resource.Attribute{
					Name:        "svcgrp_health_check_shared",
					Description: `Health Check (Monitor Name)`,
				},
				resource.Attribute{
					Name:        "template_policy",
					Description: `Policy template (Policy template name)`,
				},
				resource.Attribute{
					Name:        "template_policy_shared",
					Description: `Policy template`,
				},
				resource.Attribute{
					Name:        "template_port",
					Description: `Port template (Port template name)`,
				},
				resource.Attribute{
					Name:        "template_server",
					Description: `Server template (Server template name)`,
				},
				resource.Attribute{
					Name:        "top_fastest",
					Description: `Report top 10 fastest servers`,
				},
				resource.Attribute{
					Name:        "top_slowest",
					Description: `Report top 10 slowest servers`,
				},
				resource.Attribute{
					Name:        "traffic_replication_mirror",
					Description: `Mirror Bi-directional Packet`,
				},
				resource.Attribute{
					Name:        "traffic_replication_mirror_da_repl",
					Description: `Replace Destination MAC`,
				},
				resource.Attribute{
					Name:        "traffic_replication_mirror_ip_repl",
					Description: `Replaces IP with server-IP`,
				},
				resource.Attribute{
					Name:        "traffic_replication_mirror_sa_da_repl",
					Description: `Replace Source MAC and Destination MAC`,
				},
				resource.Attribute{
					Name:        "traffic_replication_mirror_sa_repl",
					Description: `Replace Source MAC`,
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
					Name:        "fqdn_name",
					Description: `Server hostname - Not applicable if real server is already defined`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `IP Address - Not applicable if real server is already defined`,
				},
				resource.Attribute{
					Name:        "member_priority",
					Description: `Priority of Port in the Group (Priority of Port in the Group, default is 1)`,
				},
				resource.Attribute{
					Name:        "member_state",
					Description: `‘enable’: Enable member service port; ‘disable’: Disable member service port; ‘disable-with-health-check’: disable member service port, but health check work;`,
				},
				resource.Attribute{
					Name:        "member_stats_data_disable",
					Description: `Disable statistical data collection`,
				},
				resource.Attribute{
					Name:        "member_template",
					Description: `Real server port template (Real server port template name)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number`,
				},
				resource.Attribute{
					Name:        "resolve_as",
					Description: `‘resolve-to-ipv4’: Use A Query only to resolve FQDN; ‘resolve-to-ipv6’: Use AAAA Query only to resolve FQDN; ‘resolve-to-ipv4-and-ipv6’: Use A as well as AAAA Query to resolve FQDN;`,
				},
				resource.Attribute{
					Name:        "server_ipv6_addr",
					Description: `IPV6 Address - Not applicable if real server is already defined`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘server_selection_fail_drop’: Service selection fail drop; ‘server_selection_fail_reset’: Service selection fail reset; ‘service_peak_conn’: Service peak connection;`,
				},
				resource.Attribute{
					Name:        "auto_switch",
					Description: `Reset auto stateless state`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Priority option. Define different action for each priority node. (Priority in the Group)`,
				},
				resource.Attribute{
					Name:        "priority_action",
					Description: `‘drop’: Drop request when all priority nodes fail; ‘drop-if-exceed-limit’: Drop request when connection over limit; ‘proceed’: Proceed to next priority when all priority nodes fail(default); ‘reset’: Send client reset when all priority nodes fail; ‘reset-if-exceed-limit’: Send client reset when connection over limit;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_aflow",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB aflow resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_common",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb common resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "after_disable",
					Description: `Graceful shutdown after disable server/port and/or virtual server/port`,
				},
				resource.Attribute{
					Name:        "auto_nat_no_ip_refresh",
					Description: `‘enable’: enable; ‘disable’: disable;`,
				},
				resource.Attribute{
					Name:        "buff_thresh",
					Description: `Set buffer threshold`,
				},
				resource.Attribute{
					Name:        "buff_thresh_hw_buff",
					Description: `Set hardware buffer threshold`,
				},
				resource.Attribute{
					Name:        "buff_thresh_relieve_thresh",
					Description: `Relieve threshold`,
				},
				resource.Attribute{
					Name:        "buff_thresh_sys_buff_high",
					Description: `Set high water mark of system buffer`,
				},
				resource.Attribute{
					Name:        "buff_thresh_sys_buff_low",
					Description: `Set low water mark of system buffer`,
				},
				resource.Attribute{
					Name:        "compress_block_size",
					Description: `Set compression block size (Compression block size in bytes)`,
				},
				resource.Attribute{
					Name:        "disable_adaptive_resource_check",
					Description: `Disable adaptive resource check based on buffer usage`,
				},
				resource.Attribute{
					Name:        "disable_server_auto_reselect",
					Description: `Disable auto reselection of server`,
				},
				resource.Attribute{
					Name:        "dns_cache_age",
					Description: `Set DNS cache entry age, default is 300 seconds (1-1000000 seconds, default is 300 seconds)`,
				},
				resource.Attribute{
					Name:        "dns_cache_enable",
					Description: `Enable DNS cache`,
				},
				resource.Attribute{
					Name:        "dns_cache_entry_size",
					Description: `Set DNS cache entry size, default is 256 bytes (1-4096 bytes, default is 256 bytes)`,
				},
				resource.Attribute{
					Name:        "dns_vip_stateless",
					Description: `Enable DNS VIP stateless mode`,
				},
				resource.Attribute{
					Name:        "drop_icmp_to_vip_when_vip_down",
					Description: `Drop ICMP to VIP when VIP down`,
				},
				resource.Attribute{
					Name:        "dsr_health_check_enable",
					Description: `Enable dsr-health-check (direct server return health check)`,
				},
				resource.Attribute{
					Name:        "enable_l7_req_acct",
					Description: `Enable L7 request accounting`,
				},
				resource.Attribute{
					Name:        "entity",
					Description: `‘server’: Graceful shutdown server/port only; ‘virtual-server’: Graceful shutdown virtual server/port only;`,
				},
				resource.Attribute{
					Name:        "exclude_destination",
					Description: `‘local’: Maximum local rate; ‘remote’: Maximum remote rate; (Maximum rates)`,
				},
				resource.Attribute{
					Name:        "extended_stats",
					Description: `Enable global slb extended statistics`,
				},
				resource.Attribute{
					Name:        "fast_path_disable",
					Description: `Disable fast path in SLB processing`,
				},
				resource.Attribute{
					Name:        "gateway_health_check",
					Description: `Enable gateway health check`,
				},
				resource.Attribute{
					Name:        "graceful_shutdown",
					Description: `1-65535, in unit of seconds`,
				},
				resource.Attribute{
					Name:        "graceful_shutdown_enable",
					Description: `Enable graceful shutdown`,
				},
				resource.Attribute{
					Name:        "honor_server_response_ttl",
					Description: `Honor the server reponse TTL`,
				},
				resource.Attribute{
					Name:        "hw_compression",
					Description: `Use hardware compression`,
				},
				resource.Attribute{
					Name:        "hw_syn_rr",
					Description: `Configure hardware SYN round robin (range 1-500000)`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Specify the healthcheck interval, default is 5 seconds (Interval Value, in seconds (default 5))`,
				},
				resource.Attribute{
					Name:        "l2l3_trunk_lb_disable",
					Description: `Disable L2/L3 trunk LB`,
				},
				resource.Attribute{
					Name:        "log_for_reset_unknown_conn",
					Description: `Log when rate exceed`,
				},
				resource.Attribute{
					Name:        "low_latency",
					Description: `Enable low latency mode`,
				},
				resource.Attribute{
					Name:        "max_buff_queued_per_conn",
					Description: `Set per connection buffer threshold (Buffer value range 128-4096)`,
				},
				resource.Attribute{
					Name:        "max_http_header_count",
					Description: `Set maximum number of HTTP headers allowed`,
				},
				resource.Attribute{
					Name:        "max_local_rate",
					Description: `Set maximum local rate`,
				},
				resource.Attribute{
					Name:        "max_remote_rate",
					Description: `Set maximum remote rate`,
				},
				resource.Attribute{
					Name:        "msl_time",
					Description: `Configure maximum session life, default is 2 seconds (1-40 seconds, default is 2 seconds)`,
				},
				resource.Attribute{
					Name:        "mss_table",
					Description: `Set MSS table (128-750, default is 536)`,
				},
				resource.Attribute{
					Name:        "no_auto_up_on_aflex",
					Description: `Don’t automatically mark vport up when aFleX is bound`,
				},
				resource.Attribute{
					Name:        "override_port",
					Description: `Enable override port in DSR health check mode`,
				},
				resource.Attribute{
					Name:        "player_id_check_enable",
					Description: `Enable the Player id check`,
				},
				resource.Attribute{
					Name:        "range",
					Description: `auto translate port range`,
				},
				resource.Attribute{
					Name:        "range_end",
					Description: `port range end`,
				},
				resource.Attribute{
					Name:        "range_start",
					Description: `port range start`,
				},
				resource.Attribute{
					Name:        "rate_limit_logging",
					Description: `Configure rate limit logging`,
				},
				resource.Attribute{
					Name:        "reset_stale_session",
					Description: `Send reset if session in delete queue receives a SYN packet`,
				},
				resource.Attribute{
					Name:        "resolve_port_conflict",
					Description: `Enable client port service port conflicts`,
				},
				resource.Attribute{
					Name:        "response_type",
					Description: `‘single-answer’: Only cache DNS response with single answer; ‘round-robin’: Round robin;`,
				},
				resource.Attribute{
					Name:        "scale_out",
					Description: `Enable SLB scale out`,
				},
				resource.Attribute{
					Name:        "snat_gwy_for_l3",
					Description: `Use source NAT gateway for L3 traffic`,
				},
				resource.Attribute{
					Name:        "snat_on_vip",
					Description: `Enable source NAT traffic against VIP`,
				},
				resource.Attribute{
					Name:        "software",
					Description: `Software`,
				},
				resource.Attribute{
					Name:        "sort_res",
					Description: `Enable SLB sorting of resource names`,
				},
				resource.Attribute{
					Name:        "ssli_sni_hash_enable",
					Description: `Enable SSLi SNI hash table`,
				},
				resource.Attribute{
					Name:        "stateless_sg_multi_binding",
					Description: `Enable stateless service groups to be assigned to multiple L2/L3 DSR VIPs`,
				},
				resource.Attribute{
					Name:        "stats_data_disable",
					Description: `Disable global slb data statistics`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Specify the healthcheck timeout value, default is 15 seconds (Timeout Value, in seconds (default 15))`,
				},
				resource.Attribute{
					Name:        "ttl_threshold",
					Description: `Only cache DNS response with longer TTL`,
				},
				resource.Attribute{
					Name:        "use_mss_tab",
					Description: `Use MSS based on internal table for SLB processing`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "max_table_entries",
					Description: `Maximum number of entries allowed`,
				},
				resource.Attribute{
					Name:        "ipd_enable_toggle",
					Description: `‘enable’: Enable SLB DDoS protection; ‘disable’: Disable SLB DDoS protection (default);`,
				},
				resource.Attribute{
					Name:        "ipd_tcp",
					Description: `Configure packets-per-second threshold per TCP port (default: 200)`,
				},
				resource.Attribute{
					Name:        "ipd_udp",
					Description: `Configure packets-per-second threshold per UDP port (default: 200)`,
				},
				resource.Attribute{
					Name:        "ipd_logging_toggle",
					Description: `‘enable’: Enable SLB DDoS protection logging (default); ‘disable’: Disable SLB DDoS protection logging;`,
				},
				resource.Attribute{
					Name:        "exceed_action",
					Description: `Set action if threshold exceeded`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `Set max connections per period`,
				},
				resource.Attribute{
					Name:        "limit_period",
					Description: `‘100’: 100 ms; ‘1000’: 1000 ms;`,
				},
				resource.Attribute{
					Name:        "lock_out",
					Description: `Set lockout period in seconds if threshold exceeded`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `Send log if threshold exceeded`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `‘tcp’: Set TCP connection rate limit; ‘udp’: Set UDP packet rate limit;`,
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
			Type:             "vthunder_thunder_slb_common_conn_rate_limit_src_ip",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb common conn rate limit src ip resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "exceed_action",
					Description: `Set action if threshold exceeded`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `Set max connections per period`,
				},
				resource.Attribute{
					Name:        "limit_period",
					Description: `‘100’: 100 ms; ‘1000’: 1000 ms;`,
				},
				resource.Attribute{
					Name:        "lock_out",
					Description: `Set lockout period in seconds if threshold exceeded`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `Send log if threshold exceeded`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `‘tcp’: Set TCP connection rate limit; ‘udp’: Set UDP packet rate limit;`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Set threshold shared amongst all virtual ports`,
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
			Type:             "vthunder_thunder_slb_connection_reuse",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB connection_reuse resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_crl_srcip",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB crl-srcip resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_dns",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB dns resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_dns_cache",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB dns-cache resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_dns_response_rate_limiting",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB dns-response-rate-limiting resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_fast_http_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB fast-http-proxy resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_fix",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB fix resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_ftp_ctl",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB ftp-ctl resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_ftp_data",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB ftp-data resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_ftp_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB ftp-proxy resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_generic_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB generic-proxy resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_health_gateway",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB health-gateway resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_health_stat",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB health-stat resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_http2",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB http2 resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_http_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB http-proxy resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_hw_compress",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB hw-compress resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_icap",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB icap resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_icap_http",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB icap http resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_imapproxy",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB imapproxy resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_l4",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB l4 resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_l7session",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB l7session resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_mlb",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB mlb resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_mssql",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB mssql resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_mysql",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB mysql resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_passthrough",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB passthrough resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_perf",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB perf resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_persist",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB persist resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_player_id_global",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB player id global resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_pop3_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB pop3 proxy resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB proxy resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_rate_limit_log",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB rate limit log resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_rc_cache_global",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB rc cache global resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_resource_usage",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb resource usage resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cache_template_count",
					Description: `Total configurable HTTP Cache Templates in the System`,
				},
				resource.Attribute{
					Name:        "client_ssl_template_count",
					Description: `Total configurable Client SSL Templates in the System`,
				},
				resource.Attribute{
					Name:        "conn_reuse_template_count",
					Description: `Total configurable Connection reuse Templates in the System`,
				},
				resource.Attribute{
					Name:        "fast_tcp_template_count",
					Description: `Total configurable Fast TCP Templates in the System`,
				},
				resource.Attribute{
					Name:        "fast_udp_template_count",
					Description: `Total configurable Fast UDP Templates in the System`,
				},
				resource.Attribute{
					Name:        "health_monitor_count",
					Description: `Total Health Monitors in the System`,
				},
				resource.Attribute{
					Name:        "http_template_count",
					Description: `Total configurable HTTP Templates in the System`,
				},
				resource.Attribute{
					Name:        "nat_pool_addr_count",
					Description: `Total configurable NAT Pool addresses in the System`,
				},
				resource.Attribute{
					Name:        "pbslb_subnet_count",
					Description: `Total PBSLB Subnets in the System`,
				},
				resource.Attribute{
					Name:        "persist_cookie_template_count",
					Description: `Total configurable Persistent cookie Templates in the System`,
				},
				resource.Attribute{
					Name:        "persist_srcip_template_count",
					Description: `Total configurable Source IP Persistent Templates in the System`,
				},
				resource.Attribute{
					Name:        "proxy_template_count",
					Description: `Total configurable Proxy Templates in the System`,
				},
				resource.Attribute{
					Name:        "real_port_count",
					Description: `Total Real Server Ports in the System`,
				},
				resource.Attribute{
					Name:        "real_server_count",
					Description: `Total Real Servers in the System`,
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
					Name:        "slb_threshold_res_usage_percent",
					Description: `Enter the threshold as a percentage (Threshold in percentage(default is 0%))`,
				},
				resource.Attribute{
					Name:        "stream_template_count",
					Description: `Total configurable Streaming media in the System`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "virtual_port_count",
					Description: `Total Virtual Server Ports in the System`,
				},
				resource.Attribute{
					Name:        "virtual_server_count",
					Description: `Total Virtual Servers in the System`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_server_port",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb server port resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `‘enable’: enable; ‘disable’: disable; ‘disable-with-health-check’: disable port, but health check work;`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `Connection Limit`,
				},
				resource.Attribute{
					Name:        "conn_resume",
					Description: `Connection Resume`,
				},
				resource.Attribute{
					Name:        "extended_stats",
					Description: `Enable extended statistics on real server port`,
				},
				resource.Attribute{
					Name:        "follow_port_protocol",
					Description: `‘tcp’: TCP Port; ‘udp’: UDP Port;`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Health Check (Monitor Name)`,
				},
				resource.Attribute{
					Name:        "health_check_disable",
					Description: `Disable health check`,
				},
				resource.Attribute{
					Name:        "health_check_follow_port",
					Description: `Specify which port to follow for health status (Port Number)`,
				},
				resource.Attribute{
					Name:        "no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "no_ssl",
					Description: `No SSL`,
				},
				resource.Attribute{
					Name:        "port_number",
					Description: `Port Number`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `‘tcp’: TCP Port; ‘udp’: UDP Port;`,
				},
				resource.Attribute{
					Name:        "range",
					Description: `Port range (Port range value - used for vip-to-rport-mapping)`,
				},
				resource.Attribute{
					Name:        "rport_health_check_shared",
					Description: `Health Check (Monitor Name)`,
				},
				resource.Attribute{
					Name:        "shared_rport_health_check",
					Description: `Reference a health-check from shared partition`,
				},
				resource.Attribute{
					Name:        "stats_data_action",
					Description: `‘stats-data-enable’: Enable statistical data collection for real server port; ‘stats-data-disable’: Disable statistical data collection for real server port;`,
				},
				resource.Attribute{
					Name:        "support_http2",
					Description: `Starting HTTP/2 with Prior Knowledge`,
				},
				resource.Attribute{
					Name:        "template_port",
					Description: `Port template (Port template name)`,
				},
				resource.Attribute{
					Name:        "template_server_ssl",
					Description: `Server side SSL template (Server side SSL Name)`,
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
					Name:        "weight",
					Description: `Port Weight (Connection Weight)`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘curr_req’: Current requests; ‘total_req’: Total Requests; ‘total_req_succ’: Total requests succ; ‘total_fwd_bytes’: Forward bytes; ‘total_fwd_pkts’: Forward packets; ‘total_rev_bytes’: Reverse bytes; ‘total_rev_pkts’: Reverse packets; ‘total_conn’: Total connections; ‘last_total_conn’: Last total connections; ‘peak_conn’: Peak connections; ‘es_resp_200’: Response status 200; ‘es_resp_300’: Response status 300; ‘es_resp_400’: Response status 400; ‘es_resp_500’: Response status 500; ‘es_resp_other’: Response status other; ‘es_req_count’: Total proxy request; ‘es_resp_count’: Total proxy response; ‘es_resp_invalid_http’: Total non-http response; ‘total_rev_pkts_inspected’: Total reverse packets inspected; ‘total_rev_pkts_inspected_good_status_code’: Total reverse packets with good status code inspected; ‘response_time’: Response time; ‘fastest_rsp_time’: Fastest response time; ‘slowest_rsp_time’: Slowest response time; ‘curr_ssl_conn’: Current SSL connections; ‘total_ssl_conn’: Total SSL connections;`,
				},
				resource.Attribute{
					Name:        "alternate",
					Description: `Alternate Server Number`,
				},
				resource.Attribute{
					Name:        "alternate_name",
					Description: `Alternate Name`,
				},
				resource.Attribute{
					Name:        "alternate_server_port",
					Description: `Port (Alternate Server Port Value)`,
				},
				resource.Attribute{
					Name:        "service_principal_name",
					Description: `Service Principal Name (Kerberos principal name)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_sip",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB sip resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_smpp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB smpp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_smtp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB smtp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_spdy_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB spdy proxy resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_sport_rate_limit",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB sport rate limit resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_ssl_cert_revoke",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB ssl cert revoke resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_ssl_expire_check",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb ssl expire check resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "before",
					Description: `The number of days in advance notice before expiration, default is 5`,
				},
				resource.Attribute{
					Name:        "expire_address1",
					Description: `Email address for certificate expiration check`,
				},
				resource.Attribute{
					Name:        "interval_days",
					Description: `The interval of days notice after expiration, default is 2`,
				},
				resource.Attribute{
					Name:        "ssl_expire_email_address",
					Description: `Config Email address for certificate expiration check`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `‘add’: Add an exception; ‘delete’: Delete an exception; ‘clean’: Delete all exception;`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `The certificate name`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_ssl_forward_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB ssl forward proxy resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_svm_source_nat",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB svm source nat resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_switch",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB switch resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_template_cache",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template cache resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accept_reload_req",
					Description: `Accept reload requests via cache-control directives in HTTP headers`,
				},
				resource.Attribute{
					Name:        "age",
					Description: `Specify duration in seconds cached content valid, default is 3600 seconds (seconds that the cached content is valid (default 3600 seconds))`,
				},
				resource.Attribute{
					Name:        "default_policy_nocache",
					Description: `Specify default policy to be to not cache`,
				},
				resource.Attribute{
					Name:        "disable_insert_age",
					Description: `Disable insertion of age header in response served from RAM cache`,
				},
				resource.Attribute{
					Name:        "disable_insert_via",
					Description: `Disable insertion of via header in response served from RAM cache`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `Specify logging template (Logging Config name)`,
				},
				resource.Attribute{
					Name:        "max_cache_size",
					Description: `Specify maximum cache size in megabytes, default is 80MB (RAM cache size in megabytes (default 80MB))`,
				},
				resource.Attribute{
					Name:        "max_content_size",
					Description: `Maximum size (bytes) of response that can be cached - default 81920 (80KB)`,
				},
				resource.Attribute{
					Name:        "min_content_size",
					Description: `Minimum size (bytes) of response that can be cached - default 512`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specify cache template name`,
				},
				resource.Attribute{
					Name:        "remove_cookies",
					Description: `Remove cookies in response and cache`,
				},
				resource.Attribute{
					Name:        "replacement_policy",
					Description: `‘LFU’: LFU;`,
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
					Name:        "verify_host",
					Description: `Verify request using host before sending response from RAM cache`,
				},
				resource.Attribute{
					Name:        "local_uri",
					Description: `Specify Local URI for caching (Specify URI pattern that the policy should be applied to, maximum 63 charaters)`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘hits’: hits; ‘miss’: miss; ‘bytes_served’: bytes_served; ‘total_req’: total_req; ‘caching_req’: caching_req; ‘nc_req_header’: nc_req_header; ‘nc_res_header’: nc_res_header; ‘rv_success’: rv_success; ‘rv_failure’: rv_failure; ‘ims_request’: ims_request; ‘nm_response’: nm_response; ‘rsp_type_CL’: rsp_type_CL; ‘rsp_type_CE’: rsp_type_CE; ‘rsp_type_304’: rsp_type_304; ‘rsp_type_other’: rsp_type_other; ‘rsp_no_compress’: rsp_no_compress; ‘rsp_gzip’: rsp_gzip; ‘rsp_deflate’: rsp_deflate; ‘rsp_other’: rsp_other; ‘nocache_match’: nocache_match; ‘match’: match; ‘invalidate_match’: invalidate_match; ‘content_toobig’: content_toobig; ‘content_toosmall’: content_toosmall; ‘entry_create_failures’: entry_create_failures; ‘mem_size’: mem_size; ‘entry_num’: entry_num; ‘replaced_entry’: replaced_entry; ‘aging_entry’: aging_entry; ‘cleaned_entry’: cleaned_entry;`,
				},
				resource.Attribute{
					Name:        "cache_action",
					Description: `‘cache’: Specify if certain URIs should be cached; ‘nocache’: Specify if certain URIs should not be cached;`,
				},
				resource.Attribute{
					Name:        "cache_value",
					Description: `Specify seconds that content should be cached, default is age specified in cache template`,
				},
				resource.Attribute{
					Name:        "invalidate",
					Description: `Specify if URI should invalidate cache entries matching pattern (pattern that would match entries to be invalidated (64 chars max))`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `Specify URI for cache policy (Specify URI pattern that the policy should be applied to, maximum 63 charaters)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_cipher",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template cipher resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "cipher_suite",
					Description: `‘SSL3_RSA_DES_192_CBC3_SHA’: SSL3_RSA_DES_192_CBC3_SHA; ‘SSL3_RSA_RC4_128_MD5’: SSL3_RSA_RC4_128_MD5; ‘SSL3_RSA_RC4_128_SHA’: SSL3_RSA_RC4_128_SHA; ‘TLS1_RSA_AES_128_SHA’: TLS1_RSA_AES_128_SHA; ‘TLS1_RSA_AES_256_SHA’: TLS1_RSA_AES_256_SHA; ‘TLS1_RSA_AES_128_SHA256’: TLS1_RSA_AES_128_SHA256; ‘TLS1_RSA_AES_256_SHA256’: TLS1_RSA_AES_256_SHA256; ‘TLS1_DHE_RSA_AES_128_GCM_SHA256’: TLS1_DHE_RSA_AES_128_GCM_SHA256; ‘TLS1_DHE_RSA_AES_128_SHA’: TLS1_DHE_RSA_AES_128_SHA; ‘TLS1_DHE_RSA_AES_128_SHA256’: TLS1_DHE_RSA_AES_128_SHA256; ‘TLS1_DHE_RSA_AES_256_GCM_SHA384’: TLS1_DHE_RSA_AES_256_GCM_SHA384; ‘TLS1_DHE_RSA_AES_256_SHA’: TLS1_DHE_RSA_AES_256_SHA; ‘TLS1_DHE_RSA_AES_256_SHA256’: TLS1_DHE_RSA_AES_256_SHA256; ‘TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256’: TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256; ‘TLS1_ECDHE_ECDSA_AES_128_SHA’: TLS1_ECDHE_ECDSA_AES_128_SHA; ‘TLS1_ECDHE_ECDSA_AES_128_SHA256’: TLS1_ECDHE_ECDSA_AES_128_SHA256; ‘TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384’: TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384; ‘TLS1_ECDHE_ECDSA_AES_256_SHA’: TLS1_ECDHE_ECDSA_AES_256_SHA; ‘TLS1_ECDHE_RSA_AES_128_GCM_SHA256’: TLS1_ECDHE_RSA_AES_128_GCM_SHA256; ‘TLS1_ECDHE_RSA_AES_128_SHA’: TLS1_ECDHE_RSA_AES_128_SHA; ‘TLS1_ECDHE_RSA_AES_128_SHA256’: TLS1_ECDHE_RSA_AES_128_SHA256; ‘TLS1_ECDHE_RSA_AES_256_GCM_SHA384’: TLS1_ECDHE_RSA_AES_256_GCM_SHA384; ‘TLS1_ECDHE_RSA_AES_256_SHA’: TLS1_ECDHE_RSA_AES_256_SHA; ‘TLS1_RSA_AES_128_GCM_SHA256’: TLS1_RSA_AES_128_GCM_SHA256; ‘TLS1_RSA_AES_256_GCM_SHA384’: TLS1_RSA_AES_256_GCM_SHA384;`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Cipher priority (Cipher priority (default 1))`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_client_ssh",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template client ssh resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "forward_proxy_enable",
					Description: `Enable SSH forward proxy`,
				},
				resource.Attribute{
					Name:        "forward_proxy_hostkey",
					Description: `Specify private-key (Key Name)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Client SSH Template Name`,
				},
				resource.Attribute{
					Name:        "passphrase",
					Description: `Password Phrase`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
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
			Type:             "vthunder_thunder_slb_template_client_ssl",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template client ssl resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ad_group_list",
					Description: `Forward proxy bypass if ad-group matches class-list`,
				},
				resource.Attribute{
					Name:        "alert_type",
					Description: `‘fatal’: Log fatal alerts;`,
				},
				resource.Attribute{
					Name:        "auth_sg",
					Description: `Specify authorization LDAP service group`,
				},
				resource.Attribute{
					Name:        "auth_sg_dn",
					Description: `Use Subject DN as LDAP search base DN`,
				},
				resource.Attribute{
					Name:        "auth_sg_filter",
					Description: `Specify LDAP search filter`,
				},
				resource.Attribute{
					Name:        "auth_username",
					Description: `Specify the Username Field in the Client Certificate(If multi-fields are specificed, prior one has higher priority)`,
				},
				resource.Attribute{
					Name:        "auth_username_attribute",
					Description: `Specify attribute name of username for client SSL authorization`,
				},
				resource.Attribute{
					Name:        "authen_name",
					Description: `Specify authorization LDAP server name`,
				},
				resource.Attribute{
					Name:        "authorization",
					Description: `Specify LDAP server for client SSL authorizaiton`,
				},
				resource.Attribute{
					Name:        "bypass_cert_issuer_class_list_name",
					Description: `Class List Name`,
				},
				resource.Attribute{
					Name:        "bypass_cert_san_class_list_name",
					Description: `Class List Name`,
				},
				resource.Attribute{
					Name:        "bypass_cert_subject_class_list_name",
					Description: `Class List Name`,
				},
				resource.Attribute{
					Name:        "cache_persistence_list_name",
					Description: `Class List Name`,
				},
				resource.Attribute{
					Name:        "case_insensitive",
					Description: `Case insensitive forward proxy bypass`,
				},
				resource.Attribute{
					Name:        "cert_alternate",
					Description: `Specify the second certificate (Certificate Name)`,
				},
				resource.Attribute{
					Name:        "cert_revoke_action",
					Description: `‘bypass’: bypass SSLi processing; ‘continue’: continue the connection; ‘drop’: close the connection;`,
				},
				resource.Attribute{
					Name:        "cert_shared_str",
					Description: `Certificate Name`,
				},
				resource.Attribute{
					Name:        "cert_str",
					Description: `Certificate Name`,
				},
				resource.Attribute{
					Name:        "cert_unknown_action",
					Description: `‘bypass’: bypass SSLi processing; ‘continue’: continue the connection; ‘drop’: close the connection;`,
				},
				resource.Attribute{
					Name:        "chain_cert",
					Description: `Chain Certificate (Chain Certificate Name)`,
				},
				resource.Attribute{
					Name:        "chain_cert_shared_str",
					Description: `Chain Certificate Name`,
				},
				resource.Attribute{
					Name:        "class_list_name",
					Description: `Class List Name`,
				},
				resource.Attribute{
					Name:        "client_auth_case_insensitive",
					Description: `Case insensitive forward proxy client auth bypass`,
				},
				resource.Attribute{
					Name:        "client_auth_class_list",
					Description: `Forward proxy client auth bypass if SNI string matches class-list (Class List Name)`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `‘Ignore’: Don’t request client certificate; ‘Require’: Require client certificate; ‘Request’: Request client certificate;`,
				},
				resource.Attribute{
					Name:        "close_notify",
					Description: `Send close notification when terminate connection`,
				},
				resource.Attribute{
					Name:        "dgversion",
					Description: `Lower TLS/SSL version can be downgraded`,
				},
				resource.Attribute{
					Name:        "dh_type",
					Description: `‘1024’: 1024; ‘1024-dsa’: 1024-dsa; ‘2048’: 2048;`,
				},
				resource.Attribute{
					Name:        "direct_client_server_auth",
					Description: `Let backend server does SSL client authentication directly`,
				},
				resource.Attribute{
					Name:        "disable_sslv3",
					Description: `Reject Client requests for SSL version 3`,
				},
				resource.Attribute{
					Name:        "enable_tls_alert_logging",
					Description: `Enable TLS alert logging`,
				},
				resource.Attribute{
					Name:        "exception_ad_group_list",
					Description: `Exceptions to forward proxy bypass if ad-group matches class-list`,
				},
				resource.Attribute{
					Name:        "exception_certificate_issuer_cl_name",
					Description: `Exceptions to forward-proxy-bypass`,
				},
				resource.Attribute{
					Name:        "exception_certificate_san_cl_name",
					Description: `Exceptions to forward-proxy-bypass`,
				},
				resource.Attribute{
					Name:        "exception_certificate_subject_cl_name",
					Description: `Exceptions to forward-proxy-bypass`,
				},
				resource.Attribute{
					Name:        "exception_sni_cl_name",
					Description: `Exceptions to forward-proxy-bypass`,
				},
				resource.Attribute{
					Name:        "exception_user_name_list",
					Description: `Exceptions to forward proxy bypass if user-name matches class-list`,
				},
				resource.Attribute{
					Name:        "expire_hours",
					Description: `Certificate lifetime in hours`,
				},
				resource.Attribute{
					Name:        "forward_encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "forward_passphrase",
					Description: `Password Phrase`,
				},
				resource.Attribute{
					Name:        "forward_proxy_alt_sign",
					Description: `Forward proxy alternate signing cert and key`,
				},
				resource.Attribute{
					Name:        "forward_proxy_block_message",
					Description: `Message to be included on the block page (Message, enclose in quotes if spaces are present)`,
				},
				resource.Attribute{
					Name:        "forward_proxy_ca_cert",
					Description: `CA Certificate for forward proxy (SSL forward proxy CA Certificate Name)`,
				},
				resource.Attribute{
					Name:        "forward_proxy_ca_key",
					Description: `CA Private Key for forward proxy (SSL forward proxy CA Key Name)`,
				},
				resource.Attribute{
					Name:        "forward_proxy_cert_cache_limit",
					Description: `Certificate cache size limit, default is 524288 (set to 0 for unlimited size)`,
				},
				resource.Attribute{
					Name:        "forward_proxy_cert_cache_timeout",
					Description: `Certificate cache timeout, default is 1 hour (seconds, set to 0 for never timeout)`,
				},
				resource.Attribute{
					Name:        "forward_proxy_cert_expiry",
					Description: `Adjust certificate expiry relative to the time when it is created on the device`,
				},
				resource.Attribute{
					Name:        "forward_proxy_cert_not_ready_action",
					Description: `‘bypass’: bypass the connection; ‘reset’: reset the connection;`,
				},
				resource.Attribute{
					Name:        "forward_proxy_cert_revoke_action",
					Description: `Action taken if a certificate is irreversibly revoked, bypass SSLi processing by default`,
				},
				resource.Attribute{
					Name:        "forward_proxy_cert_unknown_action",
					Description: `Action taken if a certificate revocation status is unknown, bypass SSLi processing by default`,
				},
				resource.Attribute{
					Name:        "forward_proxy_crl_disable",
					Description: `Disable Certificate Revocation List checking for forward proxy`,
				},
				resource.Attribute{
					Name:        "forward_proxy_decrypted_dscp",
					Description: `Apply a DSCP to decrypted and bypassed traffic (DSCP to apply to decrypted traffic)`,
				},
				resource.Attribute{
					Name:        "forward_proxy_decrypted_dscp_bypass",
					Description: `DSCP to apply to bypassed traffic`,
				},
				resource.Attribute{
					Name:        "forward_proxy_enable",
					Description: `Enable SSL forward proxy`,
				},
				resource.Attribute{
					Name:        "forward_proxy_failsafe_disable",
					Description: `Disable Failsafe for SSL forward proxy`,
				},
				resource.Attribute{
					Name:        "forward_proxy_log_disable",
					Description: `Disable SSL forward proxy logging`,
				},
				resource.Attribute{
					Name:        "forward_proxy_no_shared_cipher_action",
					Description: `Action taken if handshake fails due to no shared ciper, close the connection by default`,
				},
				resource.Attribute{
					Name:        "forward_proxy_no_sni_action",
					Description: `‘intercept’: intercept in no SNI case; ‘bypass’: bypass in no SNI case; ‘reset’: reset in no SNI case;`,
				},
				resource.Attribute{
					Name:        "forward_proxy_ocsp_disable",
					Description: `Disable ocsp-stapling for forward proxy`,
				},
				resource.Attribute{
					Name:        "forward_proxy_selfsign_redir",
					Description: `Redirect connections to pages with self signed certs to a warning page`,
				},
				resource.Attribute{
					Name:        "forward_proxy_ssl_version",
					Description: `TLS/SSL version, default is TLS1.2 (TLS/SSL version: 31-TLSv1.0, 32-TLSv1.1 and 33-TLSv1.2)`,
				},
				resource.Attribute{
					Name:        "forward_proxy_verify_cert_fail_action",
					Description: `Action taken if certificate verification fails, close the connection by default`,
				},
				resource.Attribute{
					Name:        "fp_alt_cert",
					Description: `CA Certificate for forward proxy alternate signing (Certificate name)`,
				},
				resource.Attribute{
					Name:        "fp_alt_encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "fp_alt_key",
					Description: `CA Private Key for forward proxy alternate signing (Key name)`,
				},
				resource.Attribute{
					Name:        "fp_alt_passphrase",
					Description: `Password Phrase`,
				},
				resource.Attribute{
					Name:        "fp_cert_ext_aia_ca_issuers",
					Description: `CA Issuers (Authority Information Access URI)`,
				},
				resource.Attribute{
					Name:        "fp_cert_ext_aia_ocsp",
					Description: `OCSP (Authority Information Access URI)`,
				},
				resource.Attribute{
					Name:        "fp_cert_ext_crldp",
					Description: `CRL Distribution Point (CRL Distribution Point URI)`,
				},
				resource.Attribute{
					Name:        "fp_cert_fetch_autonat",
					Description: `‘auto’: Configure auto NAT for server certificate fetching;`,
				},
				resource.Attribute{
					Name:        "fp_cert_fetch_autonat_precedence",
					Description: `Set this NAT pool as higher precedence than other source NAT like configued under template policy`,
				},
				resource.Attribute{
					Name:        "fp_cert_fetch_natpool_name",
					Description: `Specify NAT pool or pool group`,
				},
				resource.Attribute{
					Name:        "fp_cert_fetch_natpool_name_shared",
					Description: `Specify NAT pool or pool group`,
				},
				resource.Attribute{
					Name:        "fp_cert_fetch_natpool_precedence",
					Description: `Set this NAT pool as higher precedence than other source NAT like configued under template policy`,
				},
				resource.Attribute{
					Name:        "handshake_logging_enable",
					Description: `Enable SSL handshake logging`,
				},
				resource.Attribute{
					Name:        "hsm_type",
					Description: `‘thales-embed’: Thales embed key; ‘thales-hwcrhk’: Thales hwcrhk Key;`,
				},
				resource.Attribute{
					Name:        "inspect_certificate_issuer_cl_name",
					Description: `Forward proxy Inspect if Certificate issuer matches class-list`,
				},
				resource.Attribute{
					Name:        "inspect_certificate_san_cl_name",
					Description: `Forward proxy Inspect if Certificate Subject Alternative Name matches class-list`,
				},
				resource.Attribute{
					Name:        "inspect_certificate_subject_cl_name",
					Description: `Forward proxy Inspect if Certificate Subject matches class-list`,
				},
				resource.Attribute{
					Name:        "inspect_list_name",
					Description: `Class List Name`,
				},
				resource.Attribute{
					Name:        "key_alt_encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "key_alt_passphrase",
					Description: `Password Phrase`,
				},
				resource.Attribute{
					Name:        "key_alternate",
					Description: `Specify the second private key (Key Name)`,
				},
				resource.Attribute{
					Name:        "key_encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "key_passphrase",
					Description: `Password Phrase`,
				},
				resource.Attribute{
					Name:        "key_shared_encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "key_shared_passphrase",
					Description: `Password Phrase`,
				},
				resource.Attribute{
					Name:        "key_shared_str",
					Description: `Key Name`,
				},
				resource.Attribute{
					Name:        "key_str",
					Description: `Key Name`,
				},
				resource.Attribute{
					Name:        "ldap_base_dn_from_cert",
					Description: `Use Subject DN as LDAP search base DN`,
				},
				resource.Attribute{
					Name:        "ldap_search_filter",
					Description: `Specify LDAP search filter`,
				},
				resource.Attribute{
					Name:        "local_logging",
					Description: `Enable local logging`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Client SSL Template Name`,
				},
				resource.Attribute{
					Name:        "no_shared_cipher_action",
					Description: `‘bypass’: bypass SSLi processing; ‘drop’: close the connection;`,
				},
				resource.Attribute{
					Name:        "non_ssl_bypass_l4session",
					Description: `Handle the non-ssl session as L4 for performance optimization`,
				},
				resource.Attribute{
					Name:        "non_ssl_bypass_service_group",
					Description: `Service Group for Bypass non-ssl traffic (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "notafter",
					Description: `notAfter date`,
				},
				resource.Attribute{
					Name:        "notafterday",
					Description: `Day`,
				},
				resource.Attribute{
					Name:        "notaftermonth",
					Description: `Month`,
				},
				resource.Attribute{
					Name:        "notafteryear",
					Description: `Year`,
				},
				resource.Attribute{
					Name:        "notbefore",
					Description: `notBefore date`,
				},
				resource.Attribute{
					Name:        "notbeforeday",
					Description: `Day`,
				},
				resource.Attribute{
					Name:        "notbeforemonth",
					Description: `Month`,
				},
				resource.Attribute{
					Name:        "notbeforeyear",
					Description: `Year`,
				},
				resource.Attribute{
					Name:        "ocsp_stapling",
					Description: `Config OCSP stapling support`,
				},
				resource.Attribute{
					Name:        "ocspst_ca_cert",
					Description: `CA certificate`,
				},
				resource.Attribute{
					Name:        "ocspst_ocsp",
					Description: `Specify OCSP Authentication`,
				},
				resource.Attribute{
					Name:        "ocspst_sg",
					Description: `Specify authentication service group`,
				},
				resource.Attribute{
					Name:        "ocspst_sg_days",
					Description: `Specify update period, in days`,
				},
				resource.Attribute{
					Name:        "ocspst_sg_hours",
					Description: `Specify update period, in hours`,
				},
				resource.Attribute{
					Name:        "ocspst_sg_minutes",
					Description: `Specify update period, in minutes`,
				},
				resource.Attribute{
					Name:        "ocspst_sg_timeout",
					Description: `Specify retry timeout (Default is 30 mins)`,
				},
				resource.Attribute{
					Name:        "ocspst_srvr",
					Description: `Specify OCSP authentication server`,
				},
				resource.Attribute{
					Name:        "ocspst_srvr_days",
					Description: `Specify update period, in days`,
				},
				resource.Attribute{
					Name:        "ocspst_srvr_hours",
					Description: `Specify update period, in hours`,
				},
				resource.Attribute{
					Name:        "ocspst_srvr_minutes",
					Description: `Specify update period, in minutes`,
				},
				resource.Attribute{
					Name:        "ocspst_srvr_timeout",
					Description: `Specify retry timeout (Default is 30 mins)`,
				},
				resource.Attribute{
					Name:        "renegotiation_disable",
					Description: `Disable SSL renegotiation`,
				},
				resource.Attribute{
					Name:        "require_web_category",
					Description: `Wait for web category to be resolved before taking bypass decision`,
				},
				resource.Attribute{
					Name:        "server_name_auto_map",
					Description: `Enable automatic mapping of server name indication in Client hello extension`,
				},
				resource.Attribute{
					Name:        "session_cache_size",
					Description: `Session Cache Size (Specify 0 to disable Session ID reuse.)`,
				},
				resource.Attribute{
					Name:        "session_cache_timeout",
					Description: `Session Cache Timeout (Timeout value, in seconds)`,
				},
				resource.Attribute{
					Name:        "session_ticket_lifetime",
					Description: `Session ticket lieftime in seconds from stateless session resumption (Lifetime value in seconds)`,
				},
				resource.Attribute{
					Name:        "shared_partition_cipher_template",
					Description: `Reference a cipher template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_pool",
					Description: `Reference a NAT pool or pool group from shared partition`,
				},
				resource.Attribute{
					Name:        "sni_enable_log",
					Description: `Enable logging of sni-auto-map failures. Disable by default`,
				},
				resource.Attribute{
					Name:        "ssl_false_start_disable",
					Description: `disable SSL False Start`,
				},
				resource.Attribute{
					Name:        "ssli_logging",
					Description: `SSLi logging level, default is error logging only`,
				},
				resource.Attribute{
					Name:        "sslilogging",
					Description: `‘disable’: Disable all logging; ‘all’: enable all logging(error, info);`,
				},
				resource.Attribute{
					Name:        "sslv2_bypass_service_group",
					Description: `Service Group for Bypass SSLV2 (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "template_cipher",
					Description: `Cipher Template (Cipher Config Name)`,
				},
				resource.Attribute{
					Name:        "template_cipher_shared",
					Description: `Cipher Template Name`,
				},
				resource.Attribute{
					Name:        "template_hsm",
					Description: `HSM Template (HSM Template Name)`,
				},
				resource.Attribute{
					Name:        "user_name_list",
					Description: `Forward proxy bypass if user-name matches class-list`,
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
					Name:        "verify_cert_fail_action",
					Description: `‘bypass’: bypass SSLi processing; ‘continue’: continue the connection; ‘drop’: close the connection;`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `TLS/SSL version, default is the highest number supported (TLS/SSL version: 30-SSLv3.0, 31-TLSv1.0, 32-TLSv1.1 and 33-TLSv1.2)`,
				},
				resource.Attribute{
					Name:        "bypass_cert_subject_multi_class_list_name",
					Description: `Class List Name`,
				},
				resource.Attribute{
					Name:        "certificate_san_contains",
					Description: `Forward proxy bypass if Certificate SAN contains another string`,
				},
				resource.Attribute{
					Name:        "equals",
					Description: `Forward proxy bypass if SNI string equals another string`,
				},
				resource.Attribute{
					Name:        "forward_proxy_trusted_ca",
					Description: `Forward proxy trusted CA file (CA file name)`,
				},
				resource.Attribute{
					Name:        "ec",
					Description: `‘secp256r1’: X9_62_prime256v1; ‘secp384r1’: secp384r1;`,
				},
				resource.Attribute{
					Name:        "contains",
					Description: `Forward proxy bypass if SNI string contains another string`,
				},
				resource.Attribute{
					Name:        "ends_with",
					Description: `Forward proxy bypass if SNI string ends with another string`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `CA Certificate (CA Certificate Name)`,
				},
				resource.Attribute{
					Name:        "ca_shared",
					Description: `CA Certificate Partition Shared`,
				},
				resource.Attribute{
					Name:        "client_ocsp",
					Description: `Specify ocsp authentication server(s) for client certificate verification`,
				},
				resource.Attribute{
					Name:        "client_ocsp_sg",
					Description: `Specify service-group (Service group name)`,
				},
				resource.Attribute{
					Name:        "client_ocsp_srvr",
					Description: `Specify authentication server`,
				},
				resource.Attribute{
					Name:        "client_auth_contains",
					Description: `Forward proxy bypass if SNI string contains another string`,
				},
				resource.Attribute{
					Name:        "certificate_subject_contains",
					Description: `Forward proxy bypass if Certificate Subject contains another string`,
				},
				resource.Attribute{
					Name:        "client_certificate_Request_CA",
					Description: `Send CA lists in certificate request (CA Certificate Name)`,
				},
				resource.Attribute{
					Name:        "certificate_subject_starts",
					Description: `Forward proxy bypass if Certificate Subject starts with another string`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘real-estate’: real estate category; ‘computer-and-internet-security’: computer and internet security category; ‘financial-services’: financial services category; ‘business-and-economy’: business and economy category; ‘computer-and-internet-info’: computer and internet info category; ‘auctions’: auctions category; ‘shopping’: shopping category; ‘cult-and-occult’: cult and occult category; ‘travel’: travel category; ‘drugs’: drugs category; ‘adult-and-pornography’: adult and pornography category; ‘home-and-garden’: home and garden category; ‘military’: military category; ‘social-network’: social network category; ‘dead-sites’: dead sites category; ‘stock-advice-and-tools’: stock advice and tools category; ‘training-and-tools’: training and tools category; ‘dating’: dating category; ‘sex-education’: sex education category; ‘religion’: religion category; ‘entertainment-and-arts’: entertainment and arts category; ‘personal-sites-and-blogs’: personal sites and blogs category; ‘legal’: legal category; ‘local-information’: local information category; ‘streaming-media’: streaming media category; ‘job-search’: job search category; ‘gambling’: gambling category; ‘translation’: translation category; ‘reference-and-research’: reference and research category; ‘shareware-and-freeware’: shareware and freeware category; ‘peer-to-peer’: peer to peer category; ‘marijuana’: marijuana category; ‘hacking’: hacking category; ‘games’: games category; ‘philosophy-and-politics’: philosophy and politics category; ‘weapons’: weapons category; ‘pay-to-surf’: pay to surf category; ‘hunting-and-fishing’: hunting and fishing category; ‘society’: society category; ‘educational-institutions’: educational institutions category; ‘online-greeting-cards’: online greeting cards category; ‘sports’: sports category; ‘swimsuits-and-intimate-apparel’: swimsuits and intimate apparel category; ‘questionable’: questionable category; ‘kids’: kids category; ‘hate-and-racism’: hate and racism category; ‘personal-storage’: personal storage category; ‘violence’: violence category; ‘keyloggers-and-monitoring’: keyloggers and monitoring category; ‘search-engines’: search engines category; ‘internet-portals’: internet portals category; ‘web-advertisements’: web advertisements category; ‘cheating’: cheating category; ‘gross’: gross category; ‘web-based-email’: web based email category; ‘malware-sites’: malware sites category; ‘phishing-and-other-fraud’: phishing and other fraud category; ‘proxy-avoid-and-anonymizers’: proxy avoid and anonymizers category; ‘spyware-and-adware’: spyware and adware category; ‘music’: music category; ‘government’: government category; ‘nudity’: nudity category; ‘news-and-media’: news and media category; ‘illegal’: illegal category; ‘CDNs’: content delivery networks category; ‘internet-communications’: internet communications category; ‘bot-nets’: bot nets category; ‘abortion’: abortion category; ‘health-and-medicine’: health and medicine category; ‘confirmed-SPAM-sources’: confirmed SPAM sources category; ‘SPAM-URLs’: SPAM URLs category; ‘unconfirmed-SPAM-sources’: unconfirmed SPAM sources category; ‘open-HTTP-proxies’: open HTTP proxies category; ‘dynamic-comment’: dynamic comment category; ‘parked-domains’: parked domains category; ‘alcohol-and-tobacco’: alcohol and tobacco category; ‘private-IP-addresses’: private IP addresses category; ‘image-and-video-search’: image and video search category; ‘fashion-and-beauty’: fashion and beauty category; ‘recreation-and-hobbies’: recreation and hobbies category; ‘motor-vehicles’: motor vehicles category; ‘web-hosting-sites’: web hosting sites category; ‘food-and-dining’: food and dining category; ‘uncategorised’: uncategorised; ‘other-category’: other category;`,
				},
				resource.Attribute{
					Name:        "bypass_cert_issuer_multi_class_list_name",
					Description: `Class List Name`,
				},
				resource.Attribute{
					Name:        "client_auth_equals",
					Description: `Forward proxy bypass if SNI string equals another string`,
				},
				resource.Attribute{
					Name:        "certificate_issuer_equals",
					Description: `Forward proxy bypass if Certificate issuer equals another string`,
				},
				resource.Attribute{
					Name:        "certificate_san_ends_with",
					Description: `Forward proxy bypass if Certificate SAN ends with another string`,
				},
				resource.Attribute{
					Name:        "crl",
					Description: `Certificate Revocation Lists (Certificate Revocation Lists file name)`,
				},
				resource.Attribute{
					Name:        "crl_shared",
					Description: `Certificate Revocation Lists Partition Shared`,
				},
				resource.Attribute{
					Name:        "multi_clist_name",
					Description: `Class List Name`,
				},
				resource.Attribute{
					Name:        "certificate_issuer_ends_with",
					Description: `Forward proxy bypass if Certificate issuer ends with another string`,
				},
				resource.Attribute{
					Name:        "abortion",
					Description: `Category Abortion`,
				},
				resource.Attribute{
					Name:        "adult_and_pornography",
					Description: `Category Adult and Pornography`,
				},
				resource.Attribute{
					Name:        "alcohol_and_tobacco",
					Description: `Category Alcohol and Tobacco`,
				},
				resource.Attribute{
					Name:        "auctions",
					Description: `Category Auctions`,
				},
				resource.Attribute{
					Name:        "bot_nets",
					Description: `Category Bot Nets`,
				},
				resource.Attribute{
					Name:        "business_and_economy",
					Description: `Category Business and Economy`,
				},
				resource.Attribute{
					Name:        "cdns",
					Description: `Category CDNs`,
				},
				resource.Attribute{
					Name:        "cheating",
					Description: `Category Cheating`,
				},
				resource.Attribute{
					Name:        "computer_and_internet_info",
					Description: `Category Computer and Internet Info`,
				},
				resource.Attribute{
					Name:        "computer_and_internet_security",
					Description: `Category Computer and Internet Security`,
				},
				resource.Attribute{
					Name:        "confirmed_spam_sources",
					Description: `Category Confirmed SPAM Sources`,
				},
				resource.Attribute{
					Name:        "cult_and_occult",
					Description: `Category Cult and Occult`,
				},
				resource.Attribute{
					Name:        "dating",
					Description: `Category Dating`,
				},
				resource.Attribute{
					Name:        "dead_sites",
					Description: `Category Dead Sites (db Ops only)`,
				},
				resource.Attribute{
					Name:        "drugs",
					Description: `Category Abused Drugs`,
				},
				resource.Attribute{
					Name:        "dynamic_comment",
					Description: `Category Dynamic Comment`,
				},
				resource.Attribute{
					Name:        "educational_institutions",
					Description: `Category Educational Institutions`,
				},
				resource.Attribute{
					Name:        "entertainment_and_arts",
					Description: `Category Entertainment and Arts`,
				},
				resource.Attribute{
					Name:        "fashion_and_beauty",
					Description: `Category Fashion and Beauty`,
				},
				resource.Attribute{
					Name:        "financial_services",
					Description: `Category Financial Services`,
				},
				resource.Attribute{
					Name:        "food_and_dining",
					Description: `Category Food and Dining`,
				},
				resource.Attribute{
					Name:        "gambling",
					Description: `Category Gambling`,
				},
				resource.Attribute{
					Name:        "games",
					Description: `Category Games`,
				},
				resource.Attribute{
					Name:        "government",
					Description: `Category Government`,
				},
				resource.Attribute{
					Name:        "gross",
					Description: `Category Gross`,
				},
				resource.Attribute{
					Name:        "hacking",
					Description: `Category Hacking`,
				},
				resource.Attribute{
					Name:        "hate_and_racism",
					Description: `Category Hate and Racism`,
				},
				resource.Attribute{
					Name:        "health_and_medicine",
					Description: `Category Health and Medicine`,
				},
				resource.Attribute{
					Name:        "home_and_garden",
					Description: `Category Home and Garden`,
				},
				resource.Attribute{
					Name:        "hunting_and_fishing",
					Description: `Category Hunting and Fishing`,
				},
				resource.Attribute{
					Name:        "illegal",
					Description: `Category Illegal`,
				},
				resource.Attribute{
					Name:        "image_and_video_search",
					Description: `Category Image and Video Search`,
				},
				resource.Attribute{
					Name:        "internet_communications",
					Description: `Category Internet Communications`,
				},
				resource.Attribute{
					Name:        "internet_portals",
					Description: `Category Internet Portals`,
				},
				resource.Attribute{
					Name:        "job_search",
					Description: `Category Job Search`,
				},
				resource.Attribute{
					Name:        "keyloggers_and_monitoring",
					Description: `Category Keyloggers and Monitoring`,
				},
				resource.Attribute{
					Name:        "kids",
					Description: `Category Kids`,
				},
				resource.Attribute{
					Name:        "legal",
					Description: `Category Legal`,
				},
				resource.Attribute{
					Name:        "local_information",
					Description: `Category Local Information`,
				},
				resource.Attribute{
					Name:        "malware_sites",
					Description: `Category Malware Sites`,
				},
				resource.Attribute{
					Name:        "marijuana",
					Description: `Category Marijuana`,
				},
				resource.Attribute{
					Name:        "military",
					Description: `Category Military`,
				},
				resource.Attribute{
					Name:        "motor_vehicles",
					Description: `Category Motor Vehicles`,
				},
				resource.Attribute{
					Name:        "music",
					Description: `Category Music`,
				},
				resource.Attribute{
					Name:        "news_and_media",
					Description: `Category News and Media`,
				},
				resource.Attribute{
					Name:        "nudity",
					Description: `Category Nudity`,
				},
				resource.Attribute{
					Name:        "online_greeting_cards",
					Description: `Category Online Greeting cards`,
				},
				resource.Attribute{
					Name:        "open_http_proxies",
					Description: `Category Open HTTP Proxies`,
				},
				resource.Attribute{
					Name:        "parked_domains",
					Description: `Category Parked Domains`,
				},
				resource.Attribute{
					Name:        "pay_to_surf",
					Description: `Category Pay to Surf`,
				},
				resource.Attribute{
					Name:        "peer_to_peer",
					Description: `Category Peer to Peer`,
				},
				resource.Attribute{
					Name:        "personal_sites_and_blogs",
					Description: `Category Personal sites and Blogs`,
				},
				resource.Attribute{
					Name:        "personal_storage",
					Description: `Category Personal Storage`,
				},
				resource.Attribute{
					Name:        "philosophy_and_politics",
					Description: `Category Philosophy and Political Advocacy`,
				},
				resource.Attribute{
					Name:        "phishing_and_other_fraud",
					Description: `Category Phishing and Other Frauds`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `Category Private IP Addresses`,
				},
				resource.Attribute{
					Name:        "proxy_avoid_and_anonymizers",
					Description: `Category Proxy Avoid and Anonymizers`,
				},
				resource.Attribute{
					Name:        "questionable",
					Description: `Category Questionable`,
				},
				resource.Attribute{
					Name:        "real_estate",
					Description: `Category Real Estate`,
				},
				resource.Attribute{
					Name:        "recreation_and_hobbies",
					Description: `Category Recreation and Hobbies`,
				},
				resource.Attribute{
					Name:        "reference_and_research",
					Description: `Category Reference and Research`,
				},
				resource.Attribute{
					Name:        "religion",
					Description: `Category Religion`,
				},
				resource.Attribute{
					Name:        "search_engines",
					Description: `Category Search Engines`,
				},
				resource.Attribute{
					Name:        "sex_education",
					Description: `Category Sex Education`,
				},
				resource.Attribute{
					Name:        "shareware_and_freeware",
					Description: `Category Shareware and Freeware`,
				},
				resource.Attribute{
					Name:        "shopping",
					Description: `Category Shopping`,
				},
				resource.Attribute{
					Name:        "social_network",
					Description: `Category Social Network`,
				},
				resource.Attribute{
					Name:        "society",
					Description: `Category Society`,
				},
				resource.Attribute{
					Name:        "spam_urls",
					Description: `Category SPAM URLs`,
				},
				resource.Attribute{
					Name:        "sports",
					Description: `Category Sports`,
				},
				resource.Attribute{
					Name:        "spyware_and_adware",
					Description: `Category Spyware and Adware`,
				},
				resource.Attribute{
					Name:        "stock_advice_and_tools",
					Description: `Category Stock Advice and Tools`,
				},
				resource.Attribute{
					Name:        "streaming_media",
					Description: `Category Streaming Media`,
				},
				resource.Attribute{
					Name:        "swimsuits_and_intimate_apparel",
					Description: `Category Swimsuits and Intimate Apparel`,
				},
				resource.Attribute{
					Name:        "training_and_tools",
					Description: `Category Training and Tools`,
				},
				resource.Attribute{
					Name:        "translation",
					Description: `Category Translation`,
				},
				resource.Attribute{
					Name:        "travel",
					Description: `Category Travel`,
				},
				resource.Attribute{
					Name:        "uncategorized",
					Description: `Uncategorized URLs`,
				},
				resource.Attribute{
					Name:        "unconfirmed_spam_sources",
					Description: `Category Unconfirmed SPAM Sources`,
				},
				resource.Attribute{
					Name:        "violence",
					Description: `Category Violence`,
				},
				resource.Attribute{
					Name:        "weapons",
					Description: `Category Weapons`,
				},
				resource.Attribute{
					Name:        "web_advertisements",
					Description: `Category Web Advertisements`,
				},
				resource.Attribute{
					Name:        "web_based_email",
					Description: `Category Web based email`,
				},
				resource.Attribute{
					Name:        "web_hosting_sites",
					Description: `Category Web Hosting Sites`,
				},
				resource.Attribute{
					Name:        "certificate_san_equals",
					Description: `Forward proxy bypass if Certificate SAN equals another string`,
				},
				resource.Attribute{
					Name:        "certificate_issuer_contains",
					Description: `Forward proxy bypass if Certificate issuer contains another string (Certificate issuer)`,
				},
				resource.Attribute{
					Name:        "client_auth_starts_with",
					Description: `Forward proxy bypass if SNI string starts with another string`,
				},
				resource.Attribute{
					Name:        "certificate_subject_ends_with",
					Description: `Forward proxy bypass if Certificate Subject ends with another string`,
				},
				resource.Attribute{
					Name:        "bypass_cert_san_multi_class_list_name",
					Description: `Class List Name`,
				},
				resource.Attribute{
					Name:        "server_cert",
					Description: `Server Certificate associated to SNI (Server Certificate Name)`,
				},
				resource.Attribute{
					Name:        "server_cert_regex",
					Description: `Server Certificate associated to SNI regex (Server Certificate Name)`,
				},
				resource.Attribute{
					Name:        "server_chain",
					Description: `Server Certificate Chain associated to SNI (Server Certificate Chain Name)`,
				},
				resource.Attribute{
					Name:        "server_chain_regex",
					Description: `Server Certificate Chain associated to SNI regex (Server Certificate Chain Name)`,
				},
				resource.Attribute{
					Name:        "server_encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "server_encrypted_regex",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "server_key",
					Description: `Server Private Key associated to SNI (Server Private Key Name)`,
				},
				resource.Attribute{
					Name:        "server_key_regex",
					Description: `Server Private Key associated to SNI regex (Server Private Key Name)`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Server name indication in Client hello extension (Server name String)`,
				},
				resource.Attribute{
					Name:        "server_name_alternate",
					Description: `Specific the second certifcate`,
				},
				resource.Attribute{
					Name:        "server_name_regex",
					Description: `Server name indication in Client hello extension with regular expression (Server name String with regex)`,
				},
				resource.Attribute{
					Name:        "server_name_regex_alternate",
					Description: `Specific the second certifcate`,
				},
				resource.Attribute{
					Name:        "server_passphrase",
					Description: `help Password Phrase`,
				},
				resource.Attribute{
					Name:        "server_passphrase_regex",
					Description: `help Password Phrase`,
				},
				resource.Attribute{
					Name:        "server_shared",
					Description: `Server Name Partition Shared`,
				},
				resource.Attribute{
					Name:        "server_shared_regex",
					Description: `Server Name Partition Shared`,
				},
				resource.Attribute{
					Name:        "certificate_issuer_starts",
					Description: `Forward proxy bypass if Certificate issuer starts with another string`,
				},
				resource.Attribute{
					Name:        "certificate_san_starts",
					Description: `Forward proxy bypass if Certificate SAN starts with another string`,
				},
				resource.Attribute{
					Name:        "client_auth_ends_with",
					Description: `Forward proxy bypass if SNI string ends with another string`,
				},
				resource.Attribute{
					Name:        "certificate_subject_equals",
					Description: `Forward proxy bypass if Certificate Subject equals another string`,
				},
				resource.Attribute{
					Name:        "cipher_wo_prio",
					Description: `‘SSL3_RSA_DES_192_CBC3_SHA’: SSL3_RSA_DES_192_CBC3_SHA; ‘SSL3_RSA_RC4_128_MD5’: SSL3_RSA_RC4_128_MD5; ‘SSL3_RSA_RC4_128_SHA’: SSL3_RSA_RC4_128_SHA; ‘TLS1_RSA_AES_128_SHA’: TLS1_RSA_AES_128_SHA; ‘TLS1_RSA_AES_256_SHA’: TLS1_RSA_AES_256_SHA; ‘TLS1_RSA_AES_128_SHA256’: TLS1_RSA_AES_128_SHA256; ‘TLS1_RSA_AES_256_SHA256’: TLS1_RSA_AES_256_SHA256; ‘TLS1_DHE_RSA_AES_128_GCM_SHA256’: TLS1_DHE_RSA_AES_128_GCM_SHA256; ‘TLS1_DHE_RSA_AES_128_SHA’: TLS1_DHE_RSA_AES_128_SHA; ‘TLS1_DHE_RSA_AES_128_SHA256’: TLS1_DHE_RSA_AES_128_SHA256; ‘TLS1_DHE_RSA_AES_256_GCM_SHA384’: TLS1_DHE_RSA_AES_256_GCM_SHA384; ‘TLS1_DHE_RSA_AES_256_SHA’: TLS1_DHE_RSA_AES_256_SHA; ‘TLS1_DHE_RSA_AES_256_SHA256’: TLS1_DHE_RSA_AES_256_SHA256; ‘TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256’: TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256; ‘TLS1_ECDHE_ECDSA_AES_128_SHA’: TLS1_ECDHE_ECDSA_AES_128_SHA; ‘TLS1_ECDHE_ECDSA_AES_128_SHA256’: TLS1_ECDHE_ECDSA_AES_128_SHA256; ‘TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384’: TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384; ‘TLS1_ECDHE_ECDSA_AES_256_SHA’: TLS1_ECDHE_ECDSA_AES_256_SHA; ‘TLS1_ECDHE_RSA_AES_128_GCM_SHA256’: TLS1_ECDHE_RSA_AES_128_GCM_SHA256; ‘TLS1_ECDHE_RSA_AES_128_SHA’: TLS1_ECDHE_RSA_AES_128_SHA; ‘TLS1_ECDHE_RSA_AES_128_SHA256’: TLS1_ECDHE_RSA_AES_128_SHA256; ‘TLS1_ECDHE_RSA_AES_256_GCM_SHA384’: TLS1_ECDHE_RSA_AES_256_GCM_SHA384; ‘TLS1_ECDHE_RSA_AES_256_SHA’: TLS1_ECDHE_RSA_AES_256_SHA; ‘TLS1_RSA_AES_128_GCM_SHA256’: TLS1_RSA_AES_128_GCM_SHA256; ‘TLS1_RSA_AES_256_GCM_SHA384’: TLS1_RSA_AES_256_GCM_SHA384;`,
				},
				resource.Attribute{
					Name:        "starts_with",
					Description: `Forward proxy bypass if SNI string starts with another string`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_connection_reuse",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template connection reuse resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "keep_alive_conn",
					Description: `Keep a number of server connections open`,
				},
				resource.Attribute{
					Name:        "limit_per_server",
					Description: `Max Server Connections allowed (Connections per Server Port (default 1000))`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Connection Reuse Template Name`,
				},
				resource.Attribute{
					Name:        "num_conn_per_port",
					Description: `Connections per Server Port (default 100)`,
				},
				resource.Attribute{
					Name:        "preopen",
					Description: `Preopen server connection`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout in seconds. Multiple of 60 (def 2400)`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
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
			Type:             "vthunder_thunder_slb_template_csv",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template csv resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_template_dblb",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template dblb resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "class_list",
					Description: `Specify user/password string class list (Class list name)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `DBLB template name`,
				},
				resource.Attribute{
					Name:        "server_version",
					Description: `‘MSSQL2008’: MSSQL server 2008 or 2008 R2; ‘MSSQL2012’: MSSQL server 2012; ‘MySQL’: MySQL server (any version);`,
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
					Name:        "sha1_value",
					Description: `Cleartext password`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_diameter",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template diameter resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "avp_code",
					Description: `avp code`,
				},
				resource.Attribute{
					Name:        "avp_string",
					Description: `pattern to be matched in the avp string name, max length 127 bytes`,
				},
				resource.Attribute{
					Name:        "customize_cea",
					Description: `customizing cea response`,
				},
				resource.Attribute{
					Name:        "dwr_time",
					Description: `dwr health-check timer interval (in 100 milli second unit, default is 100, 0 means unset this option)`,
				},
				resource.Attribute{
					Name:        "dwr_up_retry",
					Description: `number of successful dwr health-check before declaring target up`,
				},
				resource.Attribute{
					Name:        "forward_to_latest_server",
					Description: `Forward client message to the latest server that sends message with the same session id`,
				},
				resource.Attribute{
					Name:        "forward_unknown_session_id",
					Description: `Forward server message even it has unknown session id`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `user sesison idle timeout (in minutes, default is 5)`,
				},
				resource.Attribute{
					Name:        "load_balance_on_session_id",
					Description: `Load balance based on the session id`,
				},
				resource.Attribute{
					Name:        "multiple_origin_host",
					Description: `allowing multiple origin-host to a single server`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `diameter template Name`,
				},
				resource.Attribute{
					Name:        "origin_realm",
					Description: `origin-realm name avp`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `product name avp`,
				},
				resource.Attribute{
					Name:        "service_group_name",
					Description: `service group name, this is the service group that the message needs to be copied to`,
				},
				resource.Attribute{
					Name:        "session_age",
					Description: `user session age allowed (default 10), this is not idle-time (in minutes)`,
				},
				resource.Attribute{
					Name:        "terminate_on_cca_t",
					Description: `remove diameter session when receiving CCA-T message`,
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
					Name:        "vendor_id",
					Description: `vendor-id avp (Vendon Id)`,
				},
				resource.Attribute{
					Name:        "avp",
					Description: `customize avps for cer to the server (avp number)`,
				},
				resource.Attribute{
					Name:        "int32",
					Description: `32 bits integer`,
				},
				resource.Attribute{
					Name:        "int64",
					Description: `64 bits integer`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `mandatory avp`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `String (string name, max length 127 bytes)`,
				},
				resource.Attribute{
					Name:        "origin_host_name",
					Description: `origin-host name avp`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_dns",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template dns resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default_policy",
					Description: `‘nocache’: Cache disable; ‘cache’: Cache enable;`,
				},
				resource.Attribute{
					Name:        "disable_dns_template",
					Description: `Disable DNS template`,
				},
				resource.Attribute{
					Name:        "dnssec_service_group",
					Description: `Use different service group if DNSSEC DO bit set (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "drop",
					Description: `Drop the malformed query`,
				},
				resource.Attribute{
					Name:        "enable_cache_sharing",
					Description: `Enable DNS cache sharing`,
				},
				resource.Attribute{
					Name:        "forward",
					Description: `Forward to service group (Service group name)`,
				},
				resource.Attribute{
					Name:        "max_cache_entry_size",
					Description: `Define maximum cache entry size (Maximum cache entry size per VIP)`,
				},
				resource.Attribute{
					Name:        "max_cache_size",
					Description: `Define maximum cache size (Maximum cache entry per VIP)`,
				},
				resource.Attribute{
					Name:        "max_query_length",
					Description: `Define Maximum DNS Query Length, default is unlimited (Specify Maximum Length)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specify a class list name`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Period in minutes`,
				},
				resource.Attribute{
					Name:        "query_id_switch",
					Description: `Use DNS query ID to create sesion`,
				},
				resource.Attribute{
					Name:        "redirect_to_tcp_port",
					Description: `Direct the client to retry with TCP for DNS UDP request`,
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
					Name:        "action_value",
					Description: `‘dns-cache-disable’: Disable DNS cache when it exceeds limit; ‘dns-cache-enable’: Enable DNS cache when it exceeds limit; ‘forward’: Forward the traffic even it exceeds limit;`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit",
					Description: `Connection rate limit`,
				},
				resource.Attribute{
					Name:        "lidnum",
					Description: `Specify a limit ID`,
				},
				resource.Attribute{
					Name:        "lockout",
					Description: `Don’t accept any new connection for certain time (Lockout duration in minutes)`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `Log a message`,
				},
				resource.Attribute{
					Name:        "log_interval",
					Description: `Log interval (minute, by default system will log every over limit instance)`,
				},
				resource.Attribute{
					Name:        "over_limit_action",
					Description: `Action when exceeds limit`,
				},
				resource.Attribute{
					Name:        "per",
					Description: `Per (Number of 100ms)`,
				},
				resource.Attribute{
					Name:        "cache_action",
					Description: `‘cache-disable’: Disable dns cache; ‘cache-enable’: Enable dns cache;`,
				},
				resource.Attribute{
					Name:        "honor_server_response_ttl",
					Description: `Honor the server reponse TTL`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `TTL for cache entry (TTL in seconds)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight for cache entry`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `‘log-only’: Only log rate-limiting, do not actually rate limit. Requires enable-log configuration; ‘rate-limit’: Rate-Limit based on configuration (Default); ‘whitelist’: Whitelist, disable rate-limiting;`,
				},
				resource.Attribute{
					Name:        "enable_log",
					Description: `Enable logging`,
				},
				resource.Attribute{
					Name:        "filter_response_rate",
					Description: `Maximum allowed request rate for the filter. This should match average traffic. (default 20 per two seconds)`,
				},
				resource.Attribute{
					Name:        "response_rate",
					Description: `Responses exceeding this rate within the window will be dropped (default 5 per second)`,
				},
				resource.Attribute{
					Name:        "slip_rate",
					Description: `Every n’th response that would be rate-limited will be let through instead`,
				},
				resource.Attribute{
					Name:        "window",
					Description: `Rate-Limiting Interval in Seconds (default is one)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_dynamic_service",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template dynamic service resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "ipv4_dns_server",
					Description: `DNS Server IPv4 Address`,
				},
				resource.Attribute{
					Name:        "ipv6_dns_server",
					Description: `DNS Server IPv6 Address`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_external_service",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template external service resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `‘continue’: Continue; ‘drop’: Drop; ‘reset’: Reset;`,
				},
				resource.Attribute{
					Name:        "failure_action",
					Description: `‘continue’: Continue; ‘drop’: Drop; ‘reset’: Reset;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `External Service Template Name`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Bind a Service Group to the template (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "shared_partition_persist_source_ip_template",
					Description: `Reference a persist source ip template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_tcp_proxy_template",
					Description: `Reference a TCP Proxy template from shared partition`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IP persistence template (Source IP persistence template name)`,
				},
				resource.Attribute{
					Name:        "tcp_proxy",
					Description: `TCP proxy template (TCP proxy template name)`,
				},
				resource.Attribute{
					Name:        "template_persist_source_ip_shared",
					Description: `Source IP Persistence Template Name`,
				},
				resource.Attribute{
					Name:        "template_tcp_proxy_shared",
					Description: `TCP Proxy Template name`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout value 1 - 200 in units of 200ms, default is 5 (default is 1000ms) (1 - 200 in units of 200ms, default is 5 (1000ms))`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `‘skyfire-icap’: Skyfire ICAP service; ‘url-filter’: URL filtering service;`,
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
					Name:        "bypass_ip",
					Description: `ip address to bypass external service`,
				},
				resource.Attribute{
					Name:        "mask",
					Description: `IP prefix mask`,
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
			Type:             "vthunder_thunder_slb_template_fix",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template fix resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "insert_client_ip",
					Description: `Insert client ip to tag 11447`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `‘init’: init only log; ‘term’: termination only log; ‘both’: both initial and termination log;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `FIX Template Name`,
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
					Name:        "equals",
					Description: `Equals (Tag String)`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Create a Service Group comprising Servers (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "switching_type",
					Description: `‘sender-comp-id’: Select service group based on SenderCompID; ‘target-comp-id’: Select service group based on TargetCompID;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_ftp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template ftp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "active_mode_port",
					Description: `Non-Standard FTP Active mode port`,
				},
				resource.Attribute{
					Name:        "active_mode_port_val",
					Description: `Non-Standard FTP Active mode port`,
				},
				resource.Attribute{
					Name:        "any",
					Description: `Allow any FTP Active mode port`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `FTP template name`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `End range of FTP Active mode port`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
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
			Type:             "vthunder_thunder_slb_template_http",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template http resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "100_cont_wait_for_req_complete",
					Description: `When REQ has Expect 100 and response is not 100, then wait for whole request to be sent`,
				},
				resource.Attribute{
					Name:        "bypass_sg",
					Description: `Select service group for non-http traffic (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "client_ip_hdr_replace",
					Description: `Replace the existing header`,
				},
				resource.Attribute{
					Name:        "client_port_hdr_replace",
					Description: `Replace the existing header`,
				},
				resource.Attribute{
					Name:        "compression_auto_disable_on_high_cpu",
					Description: `Auto-disable software compression on high cpu usage (Disable compression if cpu usage is above threshold. Default is off.)`,
				},
				resource.Attribute{
					Name:        "compression_enable",
					Description: `Enable Compression`,
				},
				resource.Attribute{
					Name:        "compression_keep_accept_encoding",
					Description: `Keep accept encoding`,
				},
				resource.Attribute{
					Name:        "compression_keep_accept_encoding_enable",
					Description: `Enable Server Accept Encoding`,
				},
				resource.Attribute{
					Name:        "compression_level",
					Description: `compression level, default 1 (compression level value, default is 1)`,
				},
				resource.Attribute{
					Name:        "compression_minimum_content_length",
					Description: `Minimum Content Length (Minimum content length for compression in bytes. Default is 120.)`,
				},
				resource.Attribute{
					Name:        "cookie_format",
					Description: `‘rfc6265’: Follow rfc6265;`,
				},
				resource.Attribute{
					Name:        "failover_url",
					Description: `Failover to this URL (Failover URL Name)`,
				},
				resource.Attribute{
					Name:        "insert_client_ip",
					Description: `Insert Client IP address into HTTP header`,
				},
				resource.Attribute{
					Name:        "insert_client_ip_header_name",
					Description: `HTTP Header Name for inserting Client IP`,
				},
				resource.Attribute{
					Name:        "insert_client_port",
					Description: `Insert Client Port address into HTTP header`,
				},
				resource.Attribute{
					Name:        "insert_client_port_header_name",
					Description: `HTTP Header Name for inserting Client Port`,
				},
				resource.Attribute{
					Name:        "keep_client_alive",
					Description: `Keep client alive`,
				},
				resource.Attribute{
					Name:        "log_retry",
					Description: `log when HTTP request retry`,
				},
				resource.Attribute{
					Name:        "max_concurrent_streams",
					Description: `(http2 only) Max concurrent streams, default 100`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `HTTP Template Name`,
				},
				resource.Attribute{
					Name:        "non_http_bypass",
					Description: `Bypass non-http traffic instead of dropping`,
				},
				resource.Attribute{
					Name:        "persist_on_401",
					Description: `Persist to the same server if the response code is 401`,
				},
				resource.Attribute{
					Name:        "rd_port",
					Description: `Port (Port Number)`,
				},
				resource.Attribute{
					Name:        "rd_resp_code",
					Description: `‘301’: Moved Permanently; ‘302’: Found; ‘303’: See Other; ‘307’: Temporary Redirect;`,
				},
				resource.Attribute{
					Name:        "rd_secure",
					Description: `Use HTTPS`,
				},
				resource.Attribute{
					Name:        "rd_simple_loc",
					Description: `Redirect location tag absolute URI string`,
				},
				resource.Attribute{
					Name:        "redirect",
					Description: `Automatically send a redirect response`,
				},
				resource.Attribute{
					Name:        "req_hdr_wait_time",
					Description: `HTTP request header wait time before abort connection`,
				},
				resource.Attribute{
					Name:        "req_hdr_wait_time_val",
					Description: `Number of seconds wait for client request header (default is 7)`,
				},
				resource.Attribute{
					Name:        "request_line_case_insensitive",
					Description: `Parse http request line as case insensitive`,
				},
				resource.Attribute{
					Name:        "request_timeout",
					Description: `Request timeout if response not received (timeout in seconds)`,
				},
				resource.Attribute{
					Name:        "retry_on_5xx",
					Description: `Retry http request on HTTP 5xx code`,
				},
				resource.Attribute{
					Name:        "retry_on_5xx_per_req",
					Description: `Retry http request on HTTP 5xx code for each request`,
				},
				resource.Attribute{
					Name:        "retry_on_5xx_per_req_val",
					Description: `Number of times to retry (default is 3)`,
				},
				resource.Attribute{
					Name:        "retry_on_5xx_val",
					Description: `Number of times to retry (default is 3)`,
				},
				resource.Attribute{
					Name:        "strict_transaction_switch",
					Description: `Force server selection on every HTTP request`,
				},
				resource.Attribute{
					Name:        "term_11client_hdr_conn_close",
					Description: `Terminate HTTP 1.1 client when req has Connection: close`,
				},
				resource.Attribute{
					Name:        "url_hash_first",
					Description: `Use the begining part of URL to calculate hash value (URL string length to calculate hash value)`,
				},
				resource.Attribute{
					Name:        "url_hash_last",
					Description: `Use the end part of URL to calculate hash value (URL string length to calculate hash value)`,
				},
				resource.Attribute{
					Name:        "url_hash_offset",
					Description: `Skip part of URL to calculate hash value (Offset of the URL string)`,
				},
				resource.Attribute{
					Name:        "url_hash_persist",
					Description: `Use URL’s hash value to select server`,
				},
				resource.Attribute{
					Name:        "use_server_status",
					Description: `Use Server-Status header to do URL hashing`,
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
					Name:        "request_header_erase",
					Description: `Erase a header from HTTP request (Header Name)`,
				},
				resource.Attribute{
					Name:        "redirect_secure",
					Description: `Use HTTPS`,
				},
				resource.Attribute{
					Name:        "redirect_secure_port",
					Description: `Port (Port Number)`,
				},
				resource.Attribute{
					Name:        "redirect_match",
					Description: `URL Matching (Pattern URL String)`,
				},
				resource.Attribute{
					Name:        "rewrite_to",
					Description: `Rewrite to Destination URL String`,
				},
				resource.Attribute{
					Name:        "response_header_insert",
					Description: `Insert a header into HTTP response (Header Content (Format: “[name]: [value]”))`,
				},
				resource.Attribute{
					Name:        "response_header_insert_type",
					Description: `‘insert-if-not-exist’: Only insert the header when it does not exist; ‘insert-always’: Always insert the header even when there is a header with the same name;`,
				},
				resource.Attribute{
					Name:        "response_header_erase",
					Description: `Erase a header from HTTP response (Header Name)`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `Logging template (Logging Config name)`,
				},
				resource.Attribute{
					Name:        "url_match_string",
					Description: `URL String`,
				},
				resource.Attribute{
					Name:        "url_service_group",
					Description: `Create a Service Group comprising Servers (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "url_switching_type",
					Description: `‘contains’: Select service group if URL string contains another string; ‘ends-with’: Select service group if URL string ends with another string; ‘equals’: Select service group if URL string equals another string; ‘starts-with’: Select service group if URL string starts with another string; ‘url-case-insensitive’: Case insensitive URL switching; ‘url-hits-enable’: Enables URL Hits;`,
				},
				resource.Attribute{
					Name:        "host_match_string",
					Description: `Hostname String`,
				},
				resource.Attribute{
					Name:        "host_service_group",
					Description: `Create a Service Group comprising Servers (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "host_switching_type",
					Description: `‘contains’: Select service group if hostname contains another string; ‘ends-with’: Select service group if hostname ends with another string; ‘equals’: Select service group if hostname equals another string; ‘starts-with’: Select service group if hostname starts with another string; ‘host-hits-enable’: Enables Host Hits counters;`,
				},
				resource.Attribute{
					Name:        "response_content_replace",
					Description: `replace the data from HTTP response content (String in the http content need to be replaced)`,
				},
				resource.Attribute{
					Name:        "response_new_string",
					Description: `String will be in the http content`,
				},
				resource.Attribute{
					Name:        "request_header_insert",
					Description: `Insert a header into HTTP request (Header Content (Format: “[name]: [value]”))`,
				},
				resource.Attribute{
					Name:        "request_header_insert_type",
					Description: `‘insert-if-not-exist’: Only insert the header when it does not exist; ‘insert-always’: Always insert the header even when there is a header with the same name;`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `Compression content-type`,
				},
				resource.Attribute{
					Name:        "exclude_uri",
					Description: `Compression exclude uri`,
				},
				resource.Attribute{
					Name:        "exclude_content_type",
					Description: `Compression exclude content-type (Compression exclude content type)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_http_policy",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template http policy resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cookie_name",
					Description: `name of cookie to match (Cookie Name)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `http-policy template name`,
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
					Name:        "match_string",
					Description: `URL String`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `‘contains’: Select service group if URL string contains another string; ‘ends-with’: Select service group if URL string ends with another string; ‘equals’: Select service group if URL string equals another string; ‘starts-with’: Select service group if URL string starts with another string;`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Service Group to be used (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `‘waf’: waf; (WAF template to be used)`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `WAF template to be used (Template Name)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `‘cookie’: cookie value match; ‘host’: hostname match; ‘url’: URL match;`,
				},
				resource.Attribute{
					Name:        "geo_location",
					Description: `Geolocation name`,
				},
				resource.Attribute{
					Name:        "geo_location_service_group",
					Description: `Service Group to be used (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "geo_location_template",
					Description: `‘waf’: waf; (WAF template to be used)`,
				},
				resource.Attribute{
					Name:        "geo_location_template_name",
					Description: `WAF template to be used (Template Name)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_imap_pop3",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template imap-pop3 resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_template_logging",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template logging resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auto",
					Description: `‘auto’: Configure auto NAT for logging, default is auto enabled;`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Specfiy a format string for web logging (format string(less than 250 characters) for web logging)`,
				},
				resource.Attribute{
					Name:        "keep_end",
					Description: `Number of unmasked characters at the end (default: 0)`,
				},
				resource.Attribute{
					Name:        "keep_start",
					Description: `Number of unmasked characters at the beginning (default: 0)`,
				},
				resource.Attribute{
					Name:        "local_logging",
					Description: `1 to enable local logging (1 to enable local logging, default 0)`,
				},
				resource.Attribute{
					Name:        "mask",
					Description: `Character to mask the matched pattern (default: X)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Logging Template Name`,
				},
				resource.Attribute{
					Name:        "pcre_mask",
					Description: `Mask matched PCRE pattern in the log`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `Specify NAT pool or pool group`,
				},
				resource.Attribute{
					Name:        "pool_shared",
					Description: `Specify NAT pool or pool group`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Bind a Service Group to the logging template (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "shared_partition_pool",
					Description: `Reference a NAT pool or pool group from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_tcp_proxy_template",
					Description: `Reference a TCP Proxy template from shared partition`,
				},
				resource.Attribute{
					Name:        "tcp_proxy",
					Description: `TCP proxy template (TCP Proxy Config name)`,
				},
				resource.Attribute{
					Name:        "template_tcp_proxy_shared",
					Description: `TCP Proxy Template name`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
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
			Type:             "vthunder_thunder_slb_template_monitor",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template monitor resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Monitor template ID Number`,
				},
				resource.Attribute{
					Name:        "monitor_relation",
					Description: `‘monitor-and’: Configures the monitors in current template to work with AND logic; ‘monitor-or’: Configures the monitors in current template to work with OR logic;`,
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
					Name:        "clear_all_sequence",
					Description: `Sequence number (Specify the port physical port number)`,
				},
				resource.Attribute{
					Name:        "clear_sequence",
					Description: `Specify the port physical port number`,
				},
				resource.Attribute{
					Name:        "sessions",
					Description: `‘all’: Clear all sessions; ‘sequence’: Sequence number;`,
				},
				resource.Attribute{
					Name:        "ena_sequence",
					Description: `Sequence number (Specify the physical port number)`,
				},
				resource.Attribute{
					Name:        "enaeth",
					Description: `Specify the physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "link_up_sequence1",
					Description: `Sequence number (Specify the port physical port number)`,
				},
				resource.Attribute{
					Name:        "link_up_sequence2",
					Description: `Sequence number (Specify the port physical port number)`,
				},
				resource.Attribute{
					Name:        "link_up_sequence3",
					Description: `Sequence number (Specify the port physical port number)`,
				},
				resource.Attribute{
					Name:        "linkup_ethernet1",
					Description: `Specify the port physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "linkup_ethernet2",
					Description: `Specify the port physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "linkup_ethernet3",
					Description: `Specify the port physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "link_down_sequence1",
					Description: `Sequence number (Specify the port physical port number)`,
				},
				resource.Attribute{
					Name:        "link_down_sequence2",
					Description: `Sequence number (Specify the port physical port number)`,
				},
				resource.Attribute{
					Name:        "link_down_sequence3",
					Description: `Sequence number (Specify the port physical port number)`,
				},
				resource.Attribute{
					Name:        "linkdown_ethernet1",
					Description: `Specify the port physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "linkdown_ethernet2",
					Description: `Specify the port physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "linkdown_ethernet3",
					Description: `Specify the port physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "dis_sequence",
					Description: `Sequence number (Specify the physical port number)`,
				},
				resource.Attribute{
					Name:        "diseth",
					Description: `Specify the physical port number (Ethernet interface number)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_mqtt",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB template mqtt resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_template_persist_cookie",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template persist cookie resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cookie_name",
					Description: `Set cookie name (Cookie name, default “sto-id”)`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Set cookie domain`,
				},
				resource.Attribute{
					Name:        "dont_honor_conn_rules",
					Description: `Do not observe connection rate rules`,
				},
				resource.Attribute{
					Name:        "encrypt_level",
					Description: `Encryption level for cookie name / value`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "expire",
					Description: `Set cookie expiration time (Expiration in seconds)`,
				},
				resource.Attribute{
					Name:        "httponly",
					Description: `Enable HttpOnly attribute`,
				},
				resource.Attribute{
					Name:        "insert_always",
					Description: `Insert persist cookie to every reponse`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `Persist for server, default is port`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Cookie persistence (Cookie persistence template name)`,
				},
				resource.Attribute{
					Name:        "pass_phrase",
					Description: `Set passphrase for encryption`,
				},
				resource.Attribute{
					Name:        "pass_thru",
					Description: `Pass thru mode - Server sends the persist cookie`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Set cookie path (Cookie path, default is “/”)`,
				},
				resource.Attribute{
					Name:        "scan_all_members",
					Description: `Persist within the same server SCAN`,
				},
				resource.Attribute{
					Name:        "secure",
					Description: `Enable secure attribute`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Persist to the same server, default is port`,
				},
				resource.Attribute{
					Name:        "server_service_group",
					Description: `Persist to the same server and within the same service group`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Persist within the same service group`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
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
			Type:             "vthunder_thunder_slb_template_persist_source_ip",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template persist source ip resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dont_honor_conn_rules",
					Description: `Do not observe connection rate rules`,
				},
				resource.Attribute{
					Name:        "enforce_higher_priority",
					Description: `Enforce to use high priority node if available`,
				},
				resource.Attribute{
					Name:        "hash_persist",
					Description: `Use hash value of source IP address`,
				},
				resource.Attribute{
					Name:        "incl_dst_ip",
					Description: `Include destination IP on the persist`,
				},
				resource.Attribute{
					Name:        "incl_sport",
					Description: `Include source port on the persist`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `Persistence type`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Source IP persistence template name`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `IP subnet mask`,
				},
				resource.Attribute{
					Name:        "netmask6",
					Description: `IPV6 subnet mask`,
				},
				resource.Attribute{
					Name:        "primary_port",
					Description: `Primary port to create the persist session`,
				},
				resource.Attribute{
					Name:        "scan_all_members",
					Description: `Persist with SCAN of all members`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Persist to the same server, default is port`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Persist within the same service group`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Persistence timeout (in minutes)`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
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
			Type:             "vthunder_thunder_slb_template_policy",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template policy resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bw_list_name",
					Description: `Specify a blacklist/whitelist name`,
				},
				resource.Attribute{
					Name:        "full_domain_tree",
					Description: `Share counters between geo-location and sub regions`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Specify log interval in minutes, by default system will log every over limit instance`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Class list name or geo-location-class-list name`,
				},
				resource.Attribute{
					Name:        "over_limit",
					Description: `Specify operation in case over limit`,
				},
				resource.Attribute{
					Name:        "over_limit_lockup",
					Description: `Don’t accept any new connection for certain time (Lockup duration (minute))`,
				},
				resource.Attribute{
					Name:        "over_limit_logging",
					Description: `Log a message`,
				},
				resource.Attribute{
					Name:        "over_limit_reset",
					Description: `Reset the connection when it exceeds limit`,
				},
				resource.Attribute{
					Name:        "overlap",
					Description: `Use overlap mode for geo-location to do longest match`,
				},
				resource.Attribute{
					Name:        "share",
					Description: `Share counters between virtual ports and virtual servers`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Define timeout value of PBSLB dynamic entry (Timeout value (minute, default is 5))`,
				},
				resource.Attribute{
					Name:        "use_destination_ip",
					Description: `Use destination IP to match the policy`,
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
					Name:        "acos_event_log",
					Description: `Enable acos event logging`,
				},
				resource.Attribute{
					Name:        "local_logging",
					Description: `Enable local logging`,
				},
				resource.Attribute{
					Name:        "no_client_conn_reuse",
					Description: `Inspects only first request of a connection`,
				},
				resource.Attribute{
					Name:        "require_web_category",
					Description: `Wait for web category to be resolved before taking proxy decision`,
				},
				resource.Attribute{
					Name:        "ssli_url_filtering",
					Description: `‘bypassed-sni-disable’: Disable SNI filtering for bypassed URL’s(enabled by default); ‘intercepted-sni-enable’: Enable SNI filtering for intercepted URL’s(disabled by default); ‘intercepted-http-disable’: Disable HTTP(host/URL) filtering for intercepted URL’s(enabled by default); ‘no-sni-allow’: Allow connection if SNI filtering is enabled and SNI header is not present(Drop by default);`,
				},
				resource.Attribute{
					Name:        "ssli_url_filtering_san",
					Description: `‘enable-san’: Enable SAN filtering(disabled by default); ‘bypassed-san-disable’: Disable SAN filtering for bypassed URL’s(enabled by default); ‘intercepted-san-enable’: Enable SAN filtering for intercepted URL’s(disabled by default); ‘no-san-allow’: Allow connection if SAN filtering is enabled and SAN field is not present(Drop by default);`,
				},
				resource.Attribute{
					Name:        "action1",
					Description: `‘forward-to-internet’: Forward request to Internet; ‘forward-to-service-group’: Forward request to service group; ‘drop’: Drop request;`,
				},
				resource.Attribute{
					Name:        "drop_message",
					Description: `drop-message sent to the client as webpage(html tags are included and quotation marks are required for white spaces)`,
				},
				resource.Attribute{
					Name:        "drop_redirect_url",
					Description: `Specify URL to which client request is redirected upon being dropped`,
				},
				resource.Attribute{
					Name:        "drop_response_code",
					Description: `Specify response code for drop action`,
				},
				resource.Attribute{
					Name:        "fake_sg",
					Description: `service group to forward the packets to Internet`,
				},
				resource.Attribute{
					Name:        "fall_back",
					Description: `Fallback service group for Internet`,
				},
				resource.Attribute{
					Name:        "fall_back_snat",
					Description: `Source NAT pool or pool group for fallback server`,
				},
				resource.Attribute{
					Name:        "forward_snat",
					Description: `Source NAT pool or pool group`,
				},
				resource.Attribute{
					Name:        "http_status_code",
					Description: `‘301’: Moved permanently; ‘302’: Found;`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `Log a message`,
				},
				resource.Attribute{
					Name:        "real_sg",
					Description: `service group to forward the packets`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘fwd-policy-dns-unresolved’: Forward-policy unresolved DNS queries; ‘fwd-policy-dns-outstanding’: Forward-policy current DNS outstanding requests; ‘fwd-policy-snat-fail’: Forward-policy source-nat translation failure; ‘fwd-policy-hits’: Number of forward-policy requests for this policy template; ‘fwd-policy-forward-to-internet’: Number of forward-policy requests forwarded to internet; ‘fwd-policy-forward-to-service-group’: Number of forward-policy requests forwarded to service group; ‘fwd-policy-policy-drop’: Number of forward-policy requests dropped; ‘fwd-policy-source-match-not-found’: Forward-policy requests without matching source rule; ‘exp_client_hello_not_found’: Expected Client HELLO requests not found;`,
				},
				resource.Attribute{
					Name:        "match_any",
					Description: `Match any source`,
				},
				resource.Attribute{
					Name:        "match_authorize_policy",
					Description: `Authorize-policy for user and group based policy`,
				},
				resource.Attribute{
					Name:        "match_class_list",
					Description: `Class List Name`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Priority value of the action(higher the number higher the priority)`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action to be performed`,
				},
				resource.Attribute{
					Name:        "dest_class_list",
					Description: `Destination Class List Name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `‘host’: Match hostname; ‘url’: match URL;`,
				},
				resource.Attribute{
					Name:        "web_category_list",
					Description: `Destination Class List Name`,
				},
				resource.Attribute{
					Name:        "client_ip_l3_dest",
					Description: `Use destination IP as client IP address`,
				},
				resource.Attribute{
					Name:        "client_ip_l7_header",
					Description: `Use extract client IP address from L7 header`,
				},
				resource.Attribute{
					Name:        "header_name",
					Description: `Specify L7 header name`,
				},
				resource.Attribute{
					Name:        "action_value",
					Description: `‘forward’: Forward the traffic even it exceeds limit; ‘reset’: Reset the connection when it exceeds limit;`,
				},
				resource.Attribute{
					Name:        "bw_per",
					Description: `Per (Specify interval in number of 100ms)`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit",
					Description: `Specify bandwidth rate limit (Bandwidth rate limit in bytes)`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `Connection limit`,
				},
				resource.Attribute{
					Name:        "conn_per",
					Description: `Per (Specify interval in number of 100ms)`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit",
					Description: `Specify connection rate limit`,
				},
				resource.Attribute{
					Name:        "direct_action",
					Description: `Set action when match the lid`,
				},
				resource.Attribute{
					Name:        "direct_action_interval",
					Description: `Specify logging interval in minute (default is 3)`,
				},
				resource.Attribute{
					Name:        "direct_action_value",
					Description: `‘drop’: drop the packet; ‘reset’: Send reset back;`,
				},
				resource.Attribute{
					Name:        "direct_fail",
					Description: `Only log unsuccessful connections`,
				},
				resource.Attribute{
					Name:        "direct_logging_drp_rst",
					Description: `Configure PBSLB logging`,
				},
				resource.Attribute{
					Name:        "direct_pbslb_interval",
					Description: `Specify logging interval in minutes(default is 3)`,
				},
				resource.Attribute{
					Name:        "direct_pbslb_logging",
					Description: `Configure PBSLB logging`,
				},
				resource.Attribute{
					Name:        "direct_service_group",
					Description: `Specify a service group (Specify the service group name)`,
				},
				resource.Attribute{
					Name:        "lidnum",
					Description: `Specify a limit ID`,
				},
				resource.Attribute{
					Name:        "lockout",
					Description: `Don’t accept any new connection for certain time (Lockout duration in minutes)`,
				},
				resource.Attribute{
					Name:        "over_limit_action",
					Description: `Set action when exceeds limit`,
				},
				resource.Attribute{
					Name:        "request_limit",
					Description: `Request limit (Specify request limit)`,
				},
				resource.Attribute{
					Name:        "request_per",
					Description: `Per (Specify interval in number of 100ms)`,
				},
				resource.Attribute{
					Name:        "request_rate_limit",
					Description: `Request rate limit (Specify request rate limit)`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `Disable`,
				},
				resource.Attribute{
					Name:        "exclusive_answer",
					Description: `Exclusive Answer in DNS Response`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `IPv6 prefix`,
				},
				resource.Attribute{
					Name:        "code_range_end",
					Description: `server response code range end`,
				},
				resource.Attribute{
					Name:        "code_range_start",
					Description: `server response code range start`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `seconds`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `the times of getting the response code`,
				},
				resource.Attribute{
					Name:        "action_interval",
					Description: `Specify logging interval in minute (default is 3)`,
				},
				resource.Attribute{
					Name:        "bw_list_action",
					Description: `‘drop’: drop the packet; ‘reset’: Send reset back;`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `Only log unsuccessful connections`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specify id that maps to service group (The id number)`,
				},
				resource.Attribute{
					Name:        "logging_drp_rst",
					Description: `Configure PBSLB logging`,
				},
				resource.Attribute{
					Name:        "pbslb_interval",
					Description: `Specify logging interval in minutes`,
				},
				resource.Attribute{
					Name:        "pbslb_logging",
					Description: `Configure PBSLB logging`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Specify a service group (Specify the service group name)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_port",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template port resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "add",
					Description: `Slow start connection limit add by a number every interval (Add by this number every interval)`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit",
					Description: `Configure bandwidth rate limit on real server port (Bandwidth rate limit in Kbps)`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_duration",
					Description: `Duration in seconds the observed rate needs to honor`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_no_logging",
					Description: `Do not log bandwidth rate limit related state transitions`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_resume",
					Description: `Resume server selection after bandwidth drops below this threshold (in Kbps) (Bandwidth rate limit resume threshold (in Kbps))`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `Connection limit`,
				},
				resource.Attribute{
					Name:        "conn_limit_no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit",
					Description: `Connection rate limit`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit_no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "dampening_flaps",
					Description: `service dampening flaps count (max-flaps allowed in flap period)`,
				},
				resource.Attribute{
					Name:        "decrement",
					Description: `Decrease after every round of DNS query (default is 0)`,
				},
				resource.Attribute{
					Name:        "del_session_on_server_down",
					Description: `Delete session if the server/port goes down (either disabled/hm down)`,
				},
				resource.Attribute{
					Name:        "dest_nat",
					Description: `Destination NAT`,
				},
				resource.Attribute{
					Name:        "down_grace_period",
					Description: `Port down grace period`,
				},
				resource.Attribute{
					Name:        "down_timer",
					Description: `The timer to bring the marked down server/port to up (default is 0, never bring up) (The timer to bring up server (in second, default is 0))`,
				},
				resource.Attribute{
					Name:        "dscp",
					Description: `Differentiated Services Code Point (DSCP to Real Server IP Mapping Value)`,
				},
				resource.Attribute{
					Name:        "dynamic_member_priority",
					Description: `Set dynamic member’s priority (Initial priority (default is 16))`,
				},
				resource.Attribute{
					Name:        "every",
					Description: `Slow start connection limit increment interval (default 10)`,
				},
				resource.Attribute{
					Name:        "extended_stats",
					Description: `Enable extended statistics on real server port`,
				},
				resource.Attribute{
					Name:        "flap_period",
					Description: `take service out of rotation if max-flaps exceeded within time in seconds`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Health Check Monitor (Health monitor name)`,
				},
				resource.Attribute{
					Name:        "health_check_disable",
					Description: `Disable configured health check configuration`,
				},
				resource.Attribute{
					Name:        "inband_health_check",
					Description: `Use inband traffic to detect port’s health status`,
				},
				resource.Attribute{
					Name:        "initial_slow_start",
					Description: `Initial slow start connection limit (default 128)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Port template name`,
				},
				resource.Attribute{
					Name:        "no_ssl",
					Description: `No SSL`,
				},
				resource.Attribute{
					Name:        "rate_interval",
					Description: `‘100ms’: Use 100 ms as sampling interval; ‘second’: Use 1 second as sampling interval;`,
				},
				resource.Attribute{
					Name:        "reassign",
					Description: `Maximum reassign times before declear the server/port down (default is 25) (The maximum reassign number)`,
				},
				resource.Attribute{
					Name:        "request_rate_interval",
					Description: `‘100ms’: Use 100 ms as sampling interval; ‘second’: Use 1 second as sampling interval;`,
				},
				resource.Attribute{
					Name:        "request_rate_limit",
					Description: `Request rate limit`,
				},
				resource.Attribute{
					Name:        "request_rate_no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "resel_on_reset",
					Description: `When receiving reset from server, do the server/port reselection (default is 0, don’t do reselection)`,
				},
				resource.Attribute{
					Name:        "reset",
					Description: `Send client reset when connection rate over limit`,
				},
				resource.Attribute{
					Name:        "restore_svc_time",
					Description: `put the service back to the rotation after time in seconds`,
				},
				resource.Attribute{
					Name:        "resume",
					Description: `Resume accepting new connection after connection number drops below threshold (Connection resume threshold)`,
				},
				resource.Attribute{
					Name:        "retry",
					Description: `Maximum retry times before reassign this connection to another server/port (default is 2) (The maximum retry number)`,
				},
				resource.Attribute{
					Name:        "shared_partition_pool",
					Description: `Reference a NAT pool or pool-group from shared partition`,
				},
				resource.Attribute{
					Name:        "slow_start",
					Description: `Slowly ramp up the connection number after port is up`,
				},
				resource.Attribute{
					Name:        "source_nat",
					Description: `Source NAT (IP NAT Pool or pool group name)`,
				},
				resource.Attribute{
					Name:        "stats_data_action",
					Description: `‘stats-data-enable’: Enable statistical data collection for real server port; ‘stats-data-disable’: Disable statistical data collection for real server port;`,
				},
				resource.Attribute{
					Name:        "sub_group",
					Description: `Divide service group members into different sub groups (Sub group ID (default is 0))`,
				},
				resource.Attribute{
					Name:        "template_port_pool_shared",
					Description: `Source NAT (IP NAT Pool or pool group name)`,
				},
				resource.Attribute{
					Name:        "till",
					Description: `Slow start ends when slow start connection limit reaches a number (default 4096) (Slow start ends when connection limit reaches this number)`,
				},
				resource.Attribute{
					Name:        "times",
					Description: `Slow start connection limit multiply by a number every interval (default 2) (Multiply by this number every interval)`,
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
					Name:        "weight",
					Description: `Weight (port weight)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_reqmod_icap",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template reqmod icap resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `‘continue’: Continue; ‘drop’: Drop; ‘reset’: Reset;`,
				},
				resource.Attribute{
					Name:        "allowed_http_methods",
					Description: `List of allowed HTTP methods. Default is “Allow All”. (List of HTTP methods allowed (default “Allow All”))`,
				},
				resource.Attribute{
					Name:        "cylance",
					Description: `cylance external server`,
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
					Name:        "logging",
					Description: `logging template (Logging template name)`,
				},
				resource.Attribute{
					Name:        "min_payload_size",
					Description: `min-payload-size value 1 - 65536, default is 4096`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Reqmod ICAP Template Name`,
				},
				resource.Attribute{
					Name:        "preview",
					Description: `Preview value 1 - 32768, default is 32768`,
				},
				resource.Attribute{
					Name:        "server_ssl",
					Description: `Server SSL template (Server SSL template name)`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Bind a Service Group to the template (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "service_url",
					Description: `URL to send to ICAP server (Service URL Name)`,
				},
				resource.Attribute{
					Name:        "shared_partition_persist_source_ip_template",
					Description: `Reference a persist source ip template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_tcp_proxy_template",
					Description: `Reference a TCP Proxy template from shared partition`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IP persistence template (Source IP persistence template name)`,
				},
				resource.Attribute{
					Name:        "tcp_proxy",
					Description: `TCP proxy template (TCP proxy template name)`,
				},
				resource.Attribute{
					Name:        "template_persist_source_ip_shared",
					Description: `Source IP Persistence Template Name`,
				},
				resource.Attribute{
					Name:        "template_tcp_proxy_shared",
					Description: `TCP Proxy Template name`,
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
					Name:        "bypass_ip",
					Description: `ip address to bypass reqmod-icap service`,
				},
				resource.Attribute{
					Name:        "mask",
					Description: `IP prefix mask`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_respmod_icap",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template respmod icap resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `‘continue’: Continue; ‘drop’: Drop; ‘reset’: Reset;`,
				},
				resource.Attribute{
					Name:        "cylance",
					Description: `cylance external server`,
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
					Name:        "logging",
					Description: `logging template (Logging template name)`,
				},
				resource.Attribute{
					Name:        "min_payload_size",
					Description: `min-payload-size value 1 - 65536, default is 4096`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Reqmod ICAP Template Name`,
				},
				resource.Attribute{
					Name:        "preview",
					Description: `Preview value 1 - 32768, default is 32768`,
				},
				resource.Attribute{
					Name:        "server_ssl",
					Description: `Server SSL template (Server SSL template name)`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Bind a Service Group to the template (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "service_url",
					Description: `URL to send to ICAP server (Service URL Name)`,
				},
				resource.Attribute{
					Name:        "shared_partition_persist_source_ip_template",
					Description: `Reference a persist source ip template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_tcp_proxy_template",
					Description: `Reference a TCP Proxy template from shared partition`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IP persistence template (Source IP persistence template name)`,
				},
				resource.Attribute{
					Name:        "tcp_proxy",
					Description: `TCP proxy template (TCP proxy template name)`,
				},
				resource.Attribute{
					Name:        "template_persist_source_ip_shared",
					Description: `Source IP Persistence Template Name`,
				},
				resource.Attribute{
					Name:        "template_tcp_proxy_shared",
					Description: `TCP Proxy Template name`,
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
					Name:        "bypass_ip",
					Description: `ip address to bypass respmod-icap service`,
				},
				resource.Attribute{
					Name:        "mask",
					Description: `IP prefix mask`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_server",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template server resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "add",
					Description: `Slow start connection limit add by a number every interval (Add by this number every interval)`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit",
					Description: `Configure bandwidth rate limit on real server (Bandwidth rate limit in Kbps)`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_acct",
					Description: `‘to-server-only’: Only account for traffic sent to server; ‘from-server-only’: Only account for traffic received from server; ‘all’: Account for all traffic sent to and received from server;`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_duration",
					Description: `Duration in seconds the observed rate needs to honor`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_no_logging",
					Description: `Do not log bandwidth rate limit related state transitions`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_resume",
					Description: `Resume server selection after bandwidth drops below this threshold (in Kbps) (Bandwidth rate limit resume threshold (in Kbps))`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `Connection limit`,
				},
				resource.Attribute{
					Name:        "conn_limit_no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit",
					Description: `Connection rate limit`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit_no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "dns_query_interval",
					Description: `The interval to query DNS server for the hostname (DNS query interval (in minute, default is 10))`,
				},
				resource.Attribute{
					Name:        "dynamic_server_prefix",
					Description: `Prefix of dynamic server (Prefix of dynamic server (default is “DRS”))`,
				},
				resource.Attribute{
					Name:        "every",
					Description: `Slow start connection limit increment interval (default 10)`,
				},
				resource.Attribute{
					Name:        "extended_stats",
					Description: `Enable extended statistics on real server`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Health Check Monitor (Health monitor name)`,
				},
				resource.Attribute{
					Name:        "health_check_disable",
					Description: `Disable configured health check configuration`,
				},
				resource.Attribute{
					Name:        "initial_slow_start",
					Description: `Initial slow start connection limit (default 128)`,
				},
				resource.Attribute{
					Name:        "log_selection_failure",
					Description: `Enable real-time logging for server selection failure event`,
				},
				resource.Attribute{
					Name:        "max_dynamic_server",
					Description: `Maximum dynamic server number (Maximum dynamic server number (default is 255))`,
				},
				resource.Attribute{
					Name:        "min_ttl_ratio",
					Description: `Minimum TTL to DNS query interval ratio (Minimum TTL ratio (default is 2))`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Server template name`,
				},
				resource.Attribute{
					Name:        "rate_interval",
					Description: `‘100ms’: Use 100 ms as sampling interval; ‘second’: Use 1 second as sampling interval;`,
				},
				resource.Attribute{
					Name:        "resume",
					Description: `Resume accepting new connection after connection number drops below threshold (Connection resume threshold)`,
				},
				resource.Attribute{
					Name:        "slow_start",
					Description: `Slowly ramp up the connection number after server is up`,
				},
				resource.Attribute{
					Name:        "spoofing_cache",
					Description: `Servers under the template are spoofing cache`,
				},
				resource.Attribute{
					Name:        "stats_data_action",
					Description: `‘stats-data-enable’: Enable statistical data collection for real server; ‘stats-data-disable’: Disable statistical data collection for real server;`,
				},
				resource.Attribute{
					Name:        "till",
					Description: `Slow start ends when slow start connection limit reaches a number (default 4096) (Slow start ends when connection limit reaches this number)`,
				},
				resource.Attribute{
					Name:        "times",
					Description: `Slow start connection limit multiply by a number every interval (default 2) (Multiply by this number every interval)`,
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
					Name:        "weight",
					Description: `Weight for the Real Servers (Connection Weight)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_server_ssh",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template server ssh resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "forward_proxy_enable",
					Description: `Enable SSH forward proxy`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Server SSH Template Name`,
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
					Name:        "counters1",
					Description: `‘all’: all; ‘successful_handshakes’: successful_handshakes; ‘failed_handshakes’: failed_handshakes; ‘forwarding_errors’: forwarding_errors;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_server_ssl",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template server-ssl resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alert_type",
					Description: `‘fatal’: Log fatal alerts;`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `Specify Client certificate (Certificate Name)`,
				},
				resource.Attribute{
					Name:        "cert_shared_str",
					Description: `Certificate Name`,
				},
				resource.Attribute{
					Name:        "cipher_template",
					Description: `Cipher Template (Cipher Config Name)`,
				},
				resource.Attribute{
					Name:        "close_notify",
					Description: `Send close notification when terminate connection`,
				},
				resource.Attribute{
					Name:        "dgversion",
					Description: `Lower TLS/SSL version can be downgraded`,
				},
				resource.Attribute{
					Name:        "dh_type",
					Description: `‘1024’: 1024; ‘1024-dsa’: 1024-dsa; ‘2048’: 2048;`,
				},
				resource.Attribute{
					Name:        "enable_tls_alert_logging",
					Description: `Enable TLS alert logging`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "forward_proxy_enable",
					Description: `Enable SSL forward proxy`,
				},
				resource.Attribute{
					Name:        "handshake_logging_enable",
					Description: `Enable SSL handshake logging`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Client private-key (Key Name)`,
				},
				resource.Attribute{
					Name:        "key_shared_encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)`,
				},
				resource.Attribute{
					Name:        "key_shared_passphrase",
					Description: `Password Phrase`,
				},
				resource.Attribute{
					Name:        "key_shared_str",
					Description: `Key Name`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Server SSL Template Name`,
				},
				resource.Attribute{
					Name:        "ocsp_stapling",
					Description: `Enable ocsp-stapling support`,
				},
				resource.Attribute{
					Name:        "passphrase",
					Description: `Password Phrase`,
				},
				resource.Attribute{
					Name:        "renegotiation_disable",
					Description: `Disable SSL renegotiation`,
				},
				resource.Attribute{
					Name:        "session_cache_size",
					Description: `Session Cache Size (Specify 0 to disable Session ID reuse.)`,
				},
				resource.Attribute{
					Name:        "session_cache_timeout",
					Description: `Session Cache Timeout (Timeout value, in seconds)`,
				},
				resource.Attribute{
					Name:        "session_ticket_enable",
					Description: `Enable server side session ticket support`,
				},
				resource.Attribute{
					Name:        "shared_partition_cipher_template",
					Description: `Reference a cipher template from shared partition`,
				},
				resource.Attribute{
					Name:        "ssli_logging",
					Description: `SSLi logging level, default is error logging only`,
				},
				resource.Attribute{
					Name:        "sslilogging",
					Description: `‘disable’: Disable all logging; ‘all’: enable all logging(error, info);`,
				},
				resource.Attribute{
					Name:        "template_cipher_shared",
					Description: `Cipher Template Name`,
				},
				resource.Attribute{
					Name:        "use_client_sni",
					Description: `use client SNI`,
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
					Name:        "version",
					Description: `TLS/SSL version, default is the highest number supported (TLS/SSL version: 30-SSLv3.0, 31-TLSv1.0, 32-TLSv1.1 and 33-TLSv1.2)`,
				},
				resource.Attribute{
					Name:        "crl",
					Description: `Certificate Revocation Lists (Certificate Revocation Lists file name)`,
				},
				resource.Attribute{
					Name:        "ec",
					Description: `‘secp256r1’: X9_62_prime256v1; ‘secp384r1’: secp384r1;`,
				},
				resource.Attribute{
					Name:        "error_type",
					Description: `‘email’: Notify the error via email; ‘ignore’: Ignore the error, which mean the connection can continue; ‘logging’: Log the error; ‘trap’: Notify the error by SNMP trap;`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `Specify CA certificate`,
				},
				resource.Attribute{
					Name:        "ca_cert_partition_shared",
					Description: `CA Certificate Partition Shared`,
				},
				resource.Attribute{
					Name:        "server_ocsp_sg",
					Description: `Specify service-group (Service group name)`,
				},
				resource.Attribute{
					Name:        "server_ocsp_srvr",
					Description: `Specify authentication server`,
				},
				resource.Attribute{
					Name:        "cipher_wo_prio",
					Description: `‘SSL3_RSA_DES_192_CBC3_SHA’: SSL3_RSA_DES_192_CBC3_SHA; ‘SSL3_RSA_RC4_128_MD5’: SSL3_RSA_RC4_128_MD5; ‘SSL3_RSA_RC4_128_SHA’: SSL3_RSA_RC4_128_SHA; ‘TLS1_RSA_AES_128_SHA’: TLS1_RSA_AES_128_SHA; ‘TLS1_RSA_AES_256_SHA’: TLS1_RSA_AES_256_SHA; ‘TLS1_RSA_AES_128_SHA256’: TLS1_RSA_AES_128_SHA256; ‘TLS1_RSA_AES_256_SHA256’: TLS1_RSA_AES_256_SHA256; ‘TLS1_DHE_RSA_AES_128_GCM_SHA256’: TLS1_DHE_RSA_AES_128_GCM_SHA256; ‘TLS1_DHE_RSA_AES_128_SHA’: TLS1_DHE_RSA_AES_128_SHA; ‘TLS1_DHE_RSA_AES_128_SHA256’: TLS1_DHE_RSA_AES_128_SHA256; ‘TLS1_DHE_RSA_AES_256_GCM_SHA384’: TLS1_DHE_RSA_AES_256_GCM_SHA384; ‘TLS1_DHE_RSA_AES_256_SHA’: TLS1_DHE_RSA_AES_256_SHA; ‘TLS1_DHE_RSA_AES_256_SHA256’: TLS1_DHE_RSA_AES_256_SHA256; ‘TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256’: TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256; ‘TLS1_ECDHE_ECDSA_AES_128_SHA’: TLS1_ECDHE_ECDSA_AES_128_SHA; ‘TLS1_ECDHE_ECDSA_AES_128_SHA256’: TLS1_ECDHE_ECDSA_AES_128_SHA256; ‘TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384’: TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384; ‘TLS1_ECDHE_ECDSA_AES_256_SHA’: TLS1_ECDHE_ECDSA_AES_256_SHA; ‘TLS1_ECDHE_RSA_AES_128_GCM_SHA256’: TLS1_ECDHE_RSA_AES_128_GCM_SHA256; ‘TLS1_ECDHE_RSA_AES_128_SHA’: TLS1_ECDHE_RSA_AES_128_SHA; ‘TLS1_ECDHE_RSA_AES_128_SHA256’: TLS1_ECDHE_RSA_AES_128_SHA256; ‘TLS1_ECDHE_RSA_AES_256_GCM_SHA384’: TLS1_ECDHE_RSA_AES_256_GCM_SHA384; ‘TLS1_ECDHE_RSA_AES_256_SHA’: TLS1_ECDHE_RSA_AES_256_SHA; ‘TLS1_RSA_AES_128_GCM_SHA256’: TLS1_RSA_AES_128_GCM_SHA256; ‘TLS1_RSA_AES_256_GCM_SHA384’: TLS1_RSA_AES_256_GCM_SHA384;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_sip",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template sip resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "acl_id",
					Description: `ACL id`,
				},
				resource.Attribute{
					Name:        "acl_name_value",
					Description: `IPv4 Access List Name`,
				},
				resource.Attribute{
					Name:        "alg_dest_nat",
					Description: `Translate VIP to real server IP in SIP message when destination NAT is used`,
				},
				resource.Attribute{
					Name:        "alg_source_nat",
					Description: `Translate source IP to NAT IP in SIP message when source NAT is used`,
				},
				resource.Attribute{
					Name:        "call_id_persist_disable",
					Description: `Disable call-ID persistence`,
				},
				resource.Attribute{
					Name:        "client_keep_alive",
					Description: `Respond client keep-alive packet directly instead of forwarding to server`,
				},
				resource.Attribute{
					Name:        "dialog_aware",
					Description: `Permit system processes dialog session`,
				},
				resource.Attribute{
					Name:        "drop_when_client_fail",
					Description: `Drop current SIP message when select client fail`,
				},
				resource.Attribute{
					Name:        "drop_when_server_fail",
					Description: `Drop current SIP message when select server fail`,
				},
				resource.Attribute{
					Name:        "failed_client_selection",
					Description: `Define action when select client fail`,
				},
				resource.Attribute{
					Name:        "failed_client_selection_message",
					Description: `Send SIP message (includs status code) to server when select client fail(Format: 3 digits(1XX~6XX) space reason)`,
				},
				resource.Attribute{
					Name:        "failed_server_selection",
					Description: `Define action when select server fail`,
				},
				resource.Attribute{
					Name:        "failed_server_selection_message",
					Description: `Send SIP message (includs status code) to client when select server fail(Format: 3 digits(1XX~6XX) space reason)`,
				},
				resource.Attribute{
					Name:        "insert_client_ip",
					Description: `Insert Client IP address into SIP header`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `The interval of keep-alive packet for each persist connection (second)`,
				},
				resource.Attribute{
					Name:        "keep_server_ip_if_match_acl",
					Description: `Use Real Server IP for addresses matching the ACL for a Call-Id`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `SIP Template Name`,
				},
				resource.Attribute{
					Name:        "pstn_gw",
					Description: `configure pstn gw host name for tel: uri translate to sip: uri (Hostname String, default is “pstn”)`,
				},
				resource.Attribute{
					Name:        "server_keep_alive",
					Description: `Send server keep-alive packet for every persist connection when enable conn-reuse`,
				},
				resource.Attribute{
					Name:        "server_selection_per_request",
					Description: `Force server selection on every SIP request`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `service group name`,
				},
				resource.Attribute{
					Name:        "smp_call_id_rtp_session",
					Description: `Create the across cpu call-id rtp session`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Time in minutes`,
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
					Name:        "insert_condition_server_request",
					Description: `‘insert-if-not-exist’: Only insert the header when it does not exist; ‘insert-always’: Always insert the header even when there is a header with the same name;`,
				},
				resource.Attribute{
					Name:        "server_request_erase_all",
					Description: `Erase all headers`,
				},
				resource.Attribute{
					Name:        "server_request_header_erase",
					Description: `Erase a SIP header (Header Name)`,
				},
				resource.Attribute{
					Name:        "server_request_header_insert",
					Description: `Insert a SIP header (Header Content (Format: “name: value”))`,
				},
				resource.Attribute{
					Name:        "insert_condition_server_response",
					Description: `‘insert-if-not-exist’: Only insert the header when it does not exist; ‘insert-always’: Always insert the header even when there is a header with the same name;`,
				},
				resource.Attribute{
					Name:        "server_response_erase_all",
					Description: `Erase all headers`,
				},
				resource.Attribute{
					Name:        "server_response_header_erase",
					Description: `Erase a SIP header (Header Name)`,
				},
				resource.Attribute{
					Name:        "server_response_header_insert",
					Description: `Insert a SIP header (Header Content (Format: “name: value”))`,
				},
				resource.Attribute{
					Name:        "client_request_erase_all",
					Description: `Erase all headers`,
				},
				resource.Attribute{
					Name:        "client_request_header_erase",
					Description: `Erase a SIP header (Header Name)`,
				},
				resource.Attribute{
					Name:        "client_request_header_insert",
					Description: `Insert a SIP header (Header Content (Format: “name: value”))`,
				},
				resource.Attribute{
					Name:        "insert_condition_client_request",
					Description: `‘insert-if-not-exist’: Only insert the header when it does not exist; ‘insert-always’: Always insert the header even when there is a header with the same name;`,
				},
				resource.Attribute{
					Name:        "client_response_erase_all",
					Description: `Erase all headers`,
				},
				resource.Attribute{
					Name:        "client_response_header_erase",
					Description: `Erase a SIP header (Header Name)`,
				},
				resource.Attribute{
					Name:        "client_response_header_insert",
					Description: `Insert a SIP header (Header Content (Format: “name: value”))`,
				},
				resource.Attribute{
					Name:        "insert_condition_client_response",
					Description: `‘insert-if-not-exist’: Only insert the header when it does not exist; ‘insert-always’: Always insert the header even when there is a header with the same name;`,
				},
				resource.Attribute{
					Name:        "header_string",
					Description: `SIP header name`,
				},
				resource.Attribute{
					Name:        "translation_value",
					Description: `‘start-line’: SIP request line or status line; ‘header’: SIP message headers; ‘body’: SIP message body;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_smpp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template smpp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "client_enquire_link",
					Description: `Respond client ENQUIRE_LINK packet directly instead of forwarding to server`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `SMPP Template Name`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Configure the password used to bind`,
				},
				resource.Attribute{
					Name:        "server_enquire_link",
					Description: `Send server ENQUIRE_LINK packet for every persist connection when enable conn-reuse`,
				},
				resource.Attribute{
					Name:        "server_enquire_link_val",
					Description: `Set interval of keep-alive packet for each persistent connection (second, default is 30)`,
				},
				resource.Attribute{
					Name:        "server_selection_per_request",
					Description: `Force server selection on every SMPP request when enable conn-reuse`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Configure the user to bind (The name used to bind)`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
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
			Type:             "vthunder_thunder_slb_template_smtp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template smtp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "client_starttls_type",
					Description: `‘optional’: STARTTLS is optional requirement; ‘enforced’: Must issue STARTTLS command before mail transaction;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `SMTP Template Name`,
				},
				resource.Attribute{
					Name:        "server_domain",
					Description: `Config the domain of the email servers (Server’s domain, default is “mail-server-domain”)`,
				},
				resource.Attribute{
					Name:        "server_starttls_type",
					Description: `‘optional’: STARTTLS is optional requirement; ‘enforced’: Must issue STARTTLS command before mail transaction;`,
				},
				resource.Attribute{
					Name:        "service_ready_msg",
					Description: `Set SMTP service ready message (SMTP service ready message, default is “ESMTP mail service ready”)`,
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
					Name:        "match_string",
					Description: `Domain name string`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Select service group (Service group name)`,
				},
				resource.Attribute{
					Name:        "switching_type",
					Description: `‘contains’: Specify domain name string if domain contains another string; ‘ends-with’: Specify domain name string if domain ends with another string; ‘starts-with’: Specify domain string if domain starts with another string;`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `Logging template (Logging Config name)`,
				},
				resource.Attribute{
					Name:        "disable_type",
					Description: `‘expn’: Disable SMTP EXPN commands; ‘turn’: Disable SMTP TURN commands; ‘vrfy’: Disable SMTP VRFY commands;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_snmp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template snmp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_template_ssli",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template ssli resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_template_tcp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template tcp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alive_if_active",
					Description: `keep connection alive if active traffic`,
				},
				resource.Attribute{
					Name:        "del_session_on_server_down",
					Description: `Delete session if the server/port goes down (either disabled/hm down)`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `send reset to client when server is disabled`,
				},
				resource.Attribute{
					Name:        "down",
					Description: `send reset to client when server is down`,
				},
				resource.Attribute{
					Name:        "force_delete_timeout",
					Description: `The maximum time that a session can stay in the system before being delete (number (second))`,
				},
				resource.Attribute{
					Name:        "force_delete_timeout_100ms",
					Description: `The maximum time that a session can stay in the system before being delete (number in 100ms)`,
				},
				resource.Attribute{
					Name:        "half_close_idle_timeout",
					Description: `TCP Half Close Idle Timeout (sec), default off (half close idle timeout in second, default off)`,
				},
				resource.Attribute{
					Name:        "half_open_idle_timeout",
					Description: `TCP Half Open Idle Timeout (sec), default off (half open idle timeout in second, default off)`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Idle Timeout value (Interval of 60 seconds), default 120 seconds (idle timeout in second, default 120)`,
				},
				resource.Attribute{
					Name:        "initial_window_size",
					Description: `Set the initial window size (number)`,
				},
				resource.Attribute{
					Name:        "insert_client_ip",
					Description: `Insert client ip into TCP option`,
				},
				resource.Attribute{
					Name:        "lan_fast_ack",
					Description: `Enable fast TCP ack on LAN`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `‘init’: init only log; ‘term’: termination only log; ‘both’: both initial and termination log;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Fast TCP Template Name`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `QOS level (number)`,
				},
				resource.Attribute{
					Name:        "reset_follow_fin",
					Description: `send reset to client or server upon receiving first fin`,
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
					Name:        "user_tag",
					Description: `Customized tag`,
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
			Type:             "vthunder_thunder_slb_template_tcp_proxy",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template tcp proxy resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ack_aggressiveness",
					Description: `‘low’: Delayed ACK; ‘medium’: Delayed ACK, with ACK on each packet with PUSH flag; ‘high’: ACK on each packet;`,
				},
				resource.Attribute{
					Name:        "alive_if_active",
					Description: `keep connection alive if active traffic`,
				},
				resource.Attribute{
					Name:        "backend_wscale",
					Description: `The TCP window scale used for the server side, default is off (number)`,
				},
				resource.Attribute{
					Name:        "del_session_on_server_down",
					Description: `Delete session if the server/port goes down (either disabled/hm down)`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `send reset to client when server is disabled`,
				},
				resource.Attribute{
					Name:        "disable_abc",
					Description: `Appropriate Byte Counting RFC 3465 Disabled, default is enabled (Appropriate Byte Counting (ABC) is enabled by default)`,
				},
				resource.Attribute{
					Name:        "disable_sack",
					Description: `disable Selective Ack Option`,
				},
				resource.Attribute{
					Name:        "disable_tcp_timestamps",
					Description: `disable TCP Timestamps Option`,
				},
				resource.Attribute{
					Name:        "disable_window_scale",
					Description: `disable TCP Window-Scale Option`,
				},
				resource.Attribute{
					Name:        "down",
					Description: `send reset to client when server is down`,
				},
				resource.Attribute{
					Name:        "dynamic_buffer_allocation",
					Description: `Optimally adjust the transmit and receive buffer sizes of TCP proxy while keeping their sum constant`,
				},
				resource.Attribute{
					Name:        "early_retransmit",
					Description: `Configure the Early-Retransmit Algorithm (RFC 5827) (Early-Retransmit is disabled by default)`,
				},
				resource.Attribute{
					Name:        "fin_timeout",
					Description: `FIN timeout (sec), default is 5 (number)`,
				},
				resource.Attribute{
					Name:        "force_delete_timeout",
					Description: `The maximum time that a session can stay in the system before being deleted, default is off (number (second))`,
				},
				resource.Attribute{
					Name:        "force_delete_timeout_100ms",
					Description: `The maximum time that a session can stay in the system before being deleted, default is off (number in 100ms)`,
				},
				resource.Attribute{
					Name:        "half_close_idle_timeout",
					Description: `TCP Half Close Idle Timeout (sec), default is off (number)`,
				},
				resource.Attribute{
					Name:        "half_open_idle_timeout",
					Description: `TCP Half Open Idle Timeout (sec), default is off (number)`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Idle Timeout (Interval of 60 seconds), default is 600 (idle timeout in second, default 600)`,
				},
				resource.Attribute{
					Name:        "init_cwnd",
					Description: `The initial congestion control window size (packets), default is 10 (number)`,
				},
				resource.Attribute{
					Name:        "initial_window_size",
					Description: `Set the initial window size, default is off (number)`,
				},
				resource.Attribute{
					Name:        "insert_client_ip",
					Description: `Insert client ip into TCP option`,
				},
				resource.Attribute{
					Name:        "invalid_rate_limit",
					Description: `Invalid Packet Response Rate Limit (ms), default is 500 (number default 500 challenges)`,
				},
				resource.Attribute{
					Name:        "keepalive_interval",
					Description: `Interval between keepalive probes (sec), default is off (number)`,
				},
				resource.Attribute{
					Name:        "keepalive_probes",
					Description: `Number of keepalive probes sent, default is off`,
				},
				resource.Attribute{
					Name:        "limited_slowstart",
					Description: `RFC 3742 Limited Slow-Start for TCP with Large Congestion Windows (number)`,
				},
				resource.Attribute{
					Name:        "maxburst",
					Description: `The max packet count sent per transmission event (number)`,
				},
				resource.Attribute{
					Name:        "min_rto",
					Description: `The minmum retransmission timeout, default is 200ms (number)`,
				},
				resource.Attribute{
					Name:        "mss",
					Description: `Responding MSS to use if client MSS is large, default is off (number)`,
				},
				resource.Attribute{
					Name:        "nagle",
					Description: `Enable Nagle Algorithm`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `TCP Proxy Template Name`,
				},
				resource.Attribute{
					Name:        "psh_flag_optimization",
					Description: `Enable Optimized PSH Flag Use`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `QOS level (number)`,
				},
				resource.Attribute{
					Name:        "reassembly_limit",
					Description: `The reassembly queuing limit, default is 25 segments (number)`,
				},
				resource.Attribute{
					Name:        "reassembly_timeout",
					Description: `The reassembly timeout, default is 30sec (number)`,
				},
				resource.Attribute{
					Name:        "receive_buffer",
					Description: `TCP Receive Buffer (default 200k) (number)`,
				},
				resource.Attribute{
					Name:        "reno",
					Description: `Enable Reno Congestion Control Algorithm`,
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
					Name:        "retransmit_retries",
					Description: `Number of Retries for Retransmit, default is 5`,
				},
				resource.Attribute{
					Name:        "server_down_action",
					Description: `‘FIN’: FIN Connection; ‘RST’: Reset Connection;`,
				},
				resource.Attribute{
					Name:        "syn_retries",
					Description: `SYN Retry Numbers, default is 5`,
				},
				resource.Attribute{
					Name:        "timewait",
					Description: `Timewait Threshold (sec), default 5 (number)`,
				},
				resource.Attribute{
					Name:        "transmit_buffer",
					Description: `TCP Transmit Buffer (default 200k) (number)`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
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
			Type:             "vthunder_thunder_slb_template_udp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template udp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "age",
					Description: `short age (in sec), default is 31`,
				},
				resource.Attribute{
					Name:        "disable_clear_session",
					Description: `Disable immediate clearing of session`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Idle Timeout value (Interval of 60 seconds), default 120 seconds (idle timeout in second, default 120)`,
				},
				resource.Attribute{
					Name:        "immediate",
					Description: `Immediate Removal after Transaction`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Fast UDP Template Name`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `QOS level (number)`,
				},
				resource.Attribute{
					Name:        "re_select_if_server_down",
					Description: `re-select another server if service port is down`,
				},
				resource.Attribute{
					Name:        "short",
					Description: `Short lived session`,
				},
				resource.Attribute{
					Name:        "stateless_conn_timeout",
					Description: `Stateless Current Connection Timeout value (5 - 120 seconds) (idle timeout in second, default 120)`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
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
			Type:             "vthunder_thunder_slb_template_virtual_port",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template virtual port resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aflow",
					Description: `Use aFlow to eliminate the traffic surge`,
				},
				resource.Attribute{
					Name:        "allow_syn_otherflags",
					Description: `Allow initial SYN packet with other flags`,
				},
				resource.Attribute{
					Name:        "allow_vip_to_rport_mapping",
					Description: `Allow mapping of VIP to real port`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `Connection limit`,
				},
				resource.Attribute{
					Name:        "conn_limit_no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "conn_limit_reset",
					Description: `Send client reset when connection over limit`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit",
					Description: `Connection rate limit`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit_no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit_reset",
					Description: `Send client reset when connection rate over limit`,
				},
				resource.Attribute{
					Name:        "drop_unknown_conn",
					Description: `Drop conection if receives TCP packet without SYN or RST flag and it does not belong to any existing connections`,
				},
				resource.Attribute{
					Name:        "dscp",
					Description: `Differentiated Services Code Point (DSCP to Real Server IP Mapping Value)`,
				},
				resource.Attribute{
					Name:        "ignore_tcp_msl",
					Description: `reclaim TCP resource immediately without MSL`,
				},
				resource.Attribute{
					Name:        "log_options",
					Description: `‘no-logging’: Do not log over limit event; ‘no-repeat-logging’: log once for over limit event. Default is log once per minute;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Virtual port template name`,
				},
				resource.Attribute{
					Name:        "non_syn_initiation",
					Description: `Allow initial TCP packet to be non-SYN`,
				},
				resource.Attribute{
					Name:        "pkt_rate_interval",
					Description: `‘100ms’: Source IP and port rate limit per 100ms; ‘second’: Source IP and port rate limit per second (default);`,
				},
				resource.Attribute{
					Name:        "pkt_rate_limit_reset",
					Description: `send client-side reset (reset after packet limit)`,
				},
				resource.Attribute{
					Name:        "pkt_rate_type",
					Description: `‘src-ip-port’: Source IP and port rate limit; ‘src-port’: Source port rate limit;`,
				},
				resource.Attribute{
					Name:        "rate",
					Description: `Source IP and port rate limit (Packet rate limit)`,
				},
				resource.Attribute{
					Name:        "rate_interval",
					Description: `‘100ms’: Use 100 ms as sampling interval; ‘second’: Use 1 second as sampling interval;`,
				},
				resource.Attribute{
					Name:        "reset_l7_on_failover",
					Description: `Send reset to L7 client and server connection upon a failover`,
				},
				resource.Attribute{
					Name:        "reset_unknown_conn",
					Description: `Send reset back if receives TCP packet without SYN or RST flag and it does not belong to any existing connections`,
				},
				resource.Attribute{
					Name:        "snat_msl",
					Description: `Source NAT MSL (Source NAT MSL value)`,
				},
				resource.Attribute{
					Name:        "snat_port_preserve",
					Description: `Source NAT Port Preservation`,
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
					Name:        "when_rr_enable",
					Description: `Only do rate limit if CPU RR triggered`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_slb_template_virtual_server",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb template virtual server resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "conn_limit",
					Description: `Connection limit`,
				},
				resource.Attribute{
					Name:        "conn_limit_no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "conn_limit_reset",
					Description: `Send client reset when connection over limit`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit",
					Description: `Connection rate limit`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit_no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "conn_rate_limit_reset",
					Description: `Send client reset when connection rate over limit`,
				},
				resource.Attribute{
					Name:        "icmp_lockup",
					Description: `Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)`,
				},
				resource.Attribute{
					Name:        "icmp_lockup_period",
					Description: `Lockup period (second)`,
				},
				resource.Attribute{
					Name:        "icmp_rate_limit",
					Description: `ICMP rate limit (Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit)`,
				},
				resource.Attribute{
					Name:        "icmpv6_lockup",
					Description: `Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)`,
				},
				resource.Attribute{
					Name:        "icmpv6_lockup_period",
					Description: `Lockup period (second)`,
				},
				resource.Attribute{
					Name:        "icmpv6_rate_limit",
					Description: `ICMPv6 rate limit (Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Virtual server template name`,
				},
				resource.Attribute{
					Name:        "rate_interval",
					Description: `‘100ms’: Use 100 ms as sampling interval; ‘second’: Use 1 second as sampling interval;`,
				},
				resource.Attribute{
					Name:        "subnet_gratuitous_arp",
					Description: `Send gratuitous ARP for every IP in the subnet virtual server`,
				},
				resource.Attribute{
					Name:        "tcp_stack_tfo_active_conn_limit",
					Description: `The allowed active layer 7 tcp fast-open connection limit, default is zero (number)`,
				},
				resource.Attribute{
					Name:        "tcp_stack_tfo_backoff_time",
					Description: `The time tcp stack will wait before allowing new fast-open requests after security condition, default 600 seconds (number)`,
				},
				resource.Attribute{
					Name:        "tcp_stack_tfo_cookie_time_limit",
					Description: `The time limit (in seconds) that a layer 7 tcp fast-open cookie is valid, default is 60 seconds (number)`,
				},
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
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
			Type:             "vthunder_thunder_slb_transparent_acl_template",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB transparent acl template resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_transparent_tcp_template",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder SLB transparent tcp template resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_slb_virtual_server_port",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder slb virtual server port resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `‘enable’: Enable; ‘disable’: Disable;`,
				},
				resource.Attribute{
					Name:        "alt_protocol1",
					Description: `‘http’: HTTP Port;`,
				},
				resource.Attribute{
					Name:        "alt_protocol2",
					Description: `‘tcp’: TCP LB service;`,
				},
				resource.Attribute{
					Name:        "alternate_port",
					Description: `Alternate Virtual Port`,
				},
				resource.Attribute{
					Name:        "alternate_port_number",
					Description: `Virtual Port`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `Configure auto NAT for the vport`,
				},
				resource.Attribute{
					Name:        "clientip_sticky_nat",
					Description: `Prefer to use same source NAT address for a client`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `Connection Limit`,
				},
				resource.Attribute{
					Name:        "def_selection_if_pref_failed",
					Description: `‘def-selection-if-pref-failed’: Use default server selection method if prefer method failed; ‘def-selection-if-pref-failed-disable’: Stop using default server selection method if prefer method failed;`,
				},
				resource.Attribute{
					Name:        "enable_playerid_check",
					Description: `Enable playerid checks on UDP packets once the AX is in active mode`,
				},
				resource.Attribute{
					Name:        "eth_fwd",
					Description: `Ethernet interface number`,
				},
				resource.Attribute{
					Name:        "eth_rev",
					Description: `Ethernet interface number`,
				},
				resource.Attribute{
					Name:        "expand",
					Description: `expand syn-cookie with timestamp and wscale`,
				},
				resource.Attribute{
					Name:        "extended_stats",
					Description: `Enable extended statistics on virtual port`,
				},
				resource.Attribute{
					Name:        "force_routing_mode",
					Description: `Force routing mode`,
				},
				resource.Attribute{
					Name:        "gslb_enable",
					Description: `Enable Global Server Load Balancing`,
				},
				resource.Attribute{
					Name:        "ha_conn_mirror",
					Description: `Enable for HA Conn sync`,
				},
				resource.Attribute{
					Name:        "ip_map_list",
					Description: `Enter name of IP Map List to be bound (IP Map List Name)`,
				},
				resource.Attribute{
					Name:        "ipinip",
					Description: `Enable IP in IP`,
				},
				resource.Attribute{
					Name:        "l7_hardware_assist",
					Description: `FPGA assist L7 packet parsing`,
				},
				resource.Attribute{
					Name:        "message_switching",
					Description: `Message switching`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `SLB Virtual Service Name`,
				},
				resource.Attribute{
					Name:        "no_auto_up_on_aflex",
					Description: `Don’t automatically mark vport up when aFleX is bound`,
				},
				resource.Attribute{
					Name:        "no_dest_nat",
					Description: `Disable destination NAT, this option only supports in wildcard VIP or when a connection is operated in SSLi + EP mode`,
				},
				resource.Attribute{
					Name:        "no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "on_syn",
					Description: `Enable for HA Conn sync for l4 tcp sessions on SYN`,
				},
				resource.Attribute{
					Name:        "optimization_level",
					Description: `‘0’: No optimization; ‘1’: Optimization level 1 (Experimental);`,
				},
				resource.Attribute{
					Name:        "persist_type",
					Description: `‘src-dst-ip-swap-persist’: Create persist session after source IP and destination IP swap; ‘use-src-ip-for-dst-persist’: Use the source IP to create a destination persist session; ‘use-dst-ip-for-src-persist’: Use the destination IP to create source IP persist session;`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `Specify NAT pool or pool group`,
				},
				resource.Attribute{
					Name:        "pool_shared",
					Description: `Specify NAT pool or pool group`,
				},
				resource.Attribute{
					Name:        "port_number",
					Description: `Port`,
				},
				resource.Attribute{
					Name:        "port_translation",
					Description: `Enable port translation under no-dest-nat`,
				},
				resource.Attribute{
					Name:        "precedence",
					Description: `Set auto NAT pool as higher precedence for source NAT`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `‘tcp’: TCP LB service; ‘udp’: UDP Port; ‘others’: for no tcp/udp protocol, do IP load balancing; ‘diameter’: diameter port; ‘dns-tcp’: DNS service over TCP; ‘dns-udp’: DNS service over UDP; ‘fast-http’: Fast HTTP Port; ‘fix’: FIX Port; ‘ftp’: File Transfer Protocol Port; ‘ftp-proxy’: ftp proxy port; ‘http’: HTTP Port; ‘https’: HTTPS port; ‘imap’: imap proxy port; ‘mlb’: Message based load balancing; ‘mms’: Microsoft Multimedia Service Port; ‘mysql’: mssql port; ‘mssql’: mssql; ‘pop3’: pop3 proxy port; ‘radius’: RADIUS Port; ‘rtsp’: Real Time Streaming Protocol Port; ‘sip’: Session initiation protocol over UDP; ‘sip-tcp’: Session initiation protocol over TCP; ‘sips’: Session initiation protocol over TLS; ‘smpp-tcp’: SMPP service over TCP; ‘spdy’: spdy port; ‘spdys’: spdys port; ‘smtp’: SMTP Port; ‘ssl-proxy’: Generic SSL proxy; ‘ssli’: SSL insight; ‘tcp-proxy’: Generic TCP proxy; ‘tftp’: TFTP Port;`,
				},
				resource.Attribute{
					Name:        "range",
					Description: `Virtual Port range (Virtual Port range value)`,
				},
				resource.Attribute{
					Name:        "rate",
					Description: `Specify the log message rate`,
				},
				resource.Attribute{
					Name:        "redirect_to_https",
					Description: `Redirect HTTP to HTTPS`,
				},
				resource.Attribute{
					Name:        "req_fail",
					Description: `Use alternate virtual port when L7 request fail`,
				},
				resource.Attribute{
					Name:        "reset",
					Description: `Send client reset when connection number over limit`,
				},
				resource.Attribute{
					Name:        "reset_on_server_selection_fail",
					Description: `Send client reset when server selection fails`,
				},
				resource.Attribute{
					Name:        "rtp_sip_call_id_match",
					Description: `rtp traffic try to match the real server of sip smp call-id session`,
				},
				resource.Attribute{
					Name:        "scaleout_bucket_count",
					Description: `Number of traffic buckets`,
				},
				resource.Attribute{
					Name:        "scaleout_device_group",
					Description: `Device group id`,
				},
				resource.Attribute{
					Name:        "secs",
					Description: `Specify the interval in seconds`,
				},
				resource.Attribute{
					Name:        "serv_sel_fail",
					Description: `Use alternate virtual port when server selection failure`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Bind a Service Group to this Virtual Server (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "shared_partition_cache_template",
					Description: `Reference a Cache template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_client_ssl_template",
					Description: `Reference a Client SSL template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_connection_reuse_template",
					Description: `Reference a connection reuse template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_diameter_template",
					Description: `Reference a Diameter template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_dns_template",
					Description: `Reference a dns template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_dynamic_service_template",
					Description: `Reference a dynamic service template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_external_service_template",
					Description: `Reference a external service template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_http_policy_template",
					Description: `Reference a http policy template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_http_template",
					Description: `Reference a HTTP template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_persist_cookie_template",
					Description: `Reference a persist cookie template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_persist_destination_ip_template",
					Description: `Reference a persist destination ip template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_persist_source_ip_template",
					Description: `Reference a persist source ip template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_persist_ssl_sid_template",
					Description: `Reference a persist SSL SID template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_policy_template",
					Description: `Reference a policy template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_pool",
					Description: `Specify NAT pool or pool group from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_server_ssl_template",
					Description: `Reference a SSL Server template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_tcp",
					Description: `Reference a tcp template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_tcp_proxy_template",
					Description: `Reference a TCP Proxy template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_udp",
					Description: `Reference a UDP template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_virtual_port_template",
					Description: `Reference a Virtual Port template from shared partition`,
				},
				resource.Attribute{
					Name:        "skip_rev_hash",
					Description: `Skip rev tuple hash insertion`,
				},
				resource.Attribute{
					Name:        "snat_on_vip",
					Description: `Enable source NAT traffic against VIP`,
				},
				resource.Attribute{
					Name:        "stats_data_action",
					Description: `‘stats-data-enable’: Enable statistical data collection for virtual port; ‘stats-data-disable’: Disable statistical data collection for virtual port;`,
				},
				resource.Attribute{
					Name:        "support_http2",
					Description: `Support HTTP2`,
				},
				resource.Attribute{
					Name:        "syn_cookie",
					Description: `Enable syn-cookie`,
				},
				resource.Attribute{
					Name:        "template_cache",
					Description: `RAM caching template (Cache Template Name)`,
				},
				resource.Attribute{
					Name:        "template_cache_shared",
					Description: `Cache Template Name`,
				},
				resource.Attribute{
					Name:        "template_client_ssh",
					Description: `Client SSH Template (Client SSH Config Name)`,
				},
				resource.Attribute{
					Name:        "template_client_ssl",
					Description: `Client SSL Template (Client SSL Config Name)`,
				},
				resource.Attribute{
					Name:        "template_client_ssl_shared",
					Description: `Client SSL Template Name`,
				},
				resource.Attribute{
					Name:        "template_connection_reuse",
					Description: `Connection Reuse Template (Connection Reuse Template Name)`,
				},
				resource.Attribute{
					Name:        "template_connection_reuse_shared",
					Description: `Connection Reuse Template Name`,
				},
				resource.Attribute{
					Name:        "template_dblb",
					Description: `DBLB Template (DBLB template name)`,
				},
				resource.Attribute{
					Name:        "template_diameter",
					Description: `Diameter Template (diameter template name)`,
				},
				resource.Attribute{
					Name:        "template_diameter_shared",
					Description: `Diameter Template Name`,
				},
				resource.Attribute{
					Name:        "template_dns",
					Description: `DNS template (DNS template name)`,
				},
				resource.Attribute{
					Name:        "template_dns_shared",
					Description: `DNS Template Name`,
				},
				resource.Attribute{
					Name:        "template_dynamic_service",
					Description: `Dynamic Service Template (dynamic-service template name)`,
				},
				resource.Attribute{
					Name:        "template_dynamic_service_shared",
					Description: `Dynamic Service Template Name`,
				},
				resource.Attribute{
					Name:        "template_external_service",
					Description: `External service template (external-service template name)`,
				},
				resource.Attribute{
					Name:        "template_external_service_shared",
					Description: `External Service Template Name`,
				},
				resource.Attribute{
					Name:        "template_file_inspection",
					Description: `File Inspection service template (file-inspection template name)`,
				},
				resource.Attribute{
					Name:        "template_fix",
					Description: `FIX template (FIX Template Name)`,
				},
				resource.Attribute{
					Name:        "template_ftp",
					Description: `FTP port template (Ftp template name)`,
				},
				resource.Attribute{
					Name:        "template_http",
					Description: `HTTP Template (HTTP Config Name)`,
				},
				resource.Attribute{
					Name:        "template_http_policy",
					Description: `http-policy template (http-policy template name)`,
				},
				resource.Attribute{
					Name:        "template_http_policy_shared",
					Description: `Http Policy Template Name`,
				},
				resource.Attribute{
					Name:        "template_http_shared",
					Description: `HTTP Template Name`,
				},
				resource.Attribute{
					Name:        "template_imap_pop3",
					Description: `IMAP/POP3 Template (IMAP/POP3 Config Name)`,
				},
				resource.Attribute{
					Name:        "template_persist_cookie",
					Description: `Cookie persistence (Cookie persistence template name)`,
				},
				resource.Attribute{
					Name:        "template_persist_cookie_shared",
					Description: `Cookie Persistence Template Name`,
				},
				resource.Attribute{
					Name:        "template_persist_destination_ip",
					Description: `Destination IP persistence (Destination IP persistence template name)`,
				},
				resource.Attribute{
					Name:        "template_persist_destination_ip_shared",
					Description: `Destination IP Persistence Template Name`,
				},
				resource.Attribute{
					Name:        "template_persist_source_ip",
					Description: `Source IP persistence (Source IP persistence template name)`,
				},
				resource.Attribute{
					Name:        "template_persist_source_ip_shared",
					Description: `Source IP Persistence Template Name`,
				},
				resource.Attribute{
					Name:        "template_persist_ssl_sid",
					Description: `SSL session ID persistence (Source IP Persistence Config name)`,
				},
				resource.Attribute{
					Name:        "template_persist_ssl_sid_shared",
					Description: `SSL SID Persistence Template Name`,
				},
				resource.Attribute{
					Name:        "template_policy",
					Description: `Policy Template (Policy template name)`,
				},
				resource.Attribute{
					Name:        "template_policy_shared",
					Description: `Policy Template Name`,
				},
				resource.Attribute{
					Name:        "template_reqmod_icap",
					Description: `ICAP reqmod template (reqmod-icap template name)`,
				},
				resource.Attribute{
					Name:        "template_respmod_icap",
					Description: `ICAP respmod service template (respmod-icap template name)`,
				},
				resource.Attribute{
					Name:        "template_scaleout",
					Description: `Scaleout template (Scaleout template name)`,
				},
				resource.Attribute{
					Name:        "template_server_ssh",
					Description: `Server SSH Template (Server SSH Config Name)`,
				},
				resource.Attribute{
					Name:        "template_server_ssl",
					Description: `Server Side SSL Template (Server SSL Name)`,
				},
				resource.Attribute{
					Name:        "template_server_ssl_shared",
					Description: `Server SSL Template Name`,
				},
				resource.Attribute{
					Name:        "template_sip",
					Description: `SIP template`,
				},
				resource.Attribute{
					Name:        "template_smpp",
					Description: `SMPP template`,
				},
				resource.Attribute{
					Name:        "template_smtp",
					Description: `SMTP Template (SMTP Config Name)`,
				},
				resource.Attribute{
					Name:        "template_ssli",
					Description: `SSLi template (SSLi Template Name)`,
				},
				resource.Attribute{
					Name:        "template_tcp",
					Description: `L4 TCP Template (TCP Config Name)`,
				},
				resource.Attribute{
					Name:        "template_tcp_proxy",
					Description: `TCP Proxy Template Name`,
				},
				resource.Attribute{
					Name:        "template_tcp_proxy_client",
					Description: `TCP Proxy Config Client (TCP Proxy Config name)`,
				},
				resource.Attribute{
					Name:        "template_tcp_proxy_server",
					Description: `TCP Proxy Config Server (TCP Proxy Config name)`,
				},
				resource.Attribute{
					Name:        "template_tcp_proxy_shared",
					Description: `TCP Proxy Template name`,
				},
				resource.Attribute{
					Name:        "template_tcp_shared",
					Description: `TCP Template Name`,
				},
				resource.Attribute{
					Name:        "template_udp",
					Description: `L4 UDP Template (UDP Config Name)`,
				},
				resource.Attribute{
					Name:        "template_udp_shared",
					Description: `UDP Template Name`,
				},
				resource.Attribute{
					Name:        "template_virtual_port",
					Description: `Virtual port template (Virtual port template name)`,
				},
				resource.Attribute{
					Name:        "template_virtual_port_shared",
					Description: `Virtual Port Template Name`,
				},
				resource.Attribute{
					Name:        "trunk_fwd",
					Description: `Trunk interface number`,
				},
				resource.Attribute{
					Name:        "trunk_rev",
					Description: `Trunk interface number`,
				},
				resource.Attribute{
					Name:        "use_alternate_port",
					Description: `Use alternate virtual port`,
				},
				resource.Attribute{
					Name:        "use_cgnv6",
					Description: `Follow CGNv6 source NAT configuration`,
				},
				resource.Attribute{
					Name:        "use_default_if_no_server",
					Description: `Use default forwarding if server selection failed`,
				},
				resource.Attribute{
					Name:        "use_rcv_hop_for_resp",
					Description: `Use receive hop for response to client(For packets on default-vlan, also config “vlan-global enable-def-vlan-l2-forwarding”.)`,
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
					Name:        "view",
					Description: `Specify a GSLB View (ID)`,
				},
				resource.Attribute{
					Name:        "waf_template",
					Description: `WAF template (WAF Template Name)`,
				},
				resource.Attribute{
					Name:        "when_down",
					Description: `Use alternate virtual port when down`,
				},
				resource.Attribute{
					Name:        "when_down_protocol2",
					Description: `Use alternate virtual port when down`,
				},
				resource.Attribute{
					Name:        "acl_name",
					Description: `Apply an access list name (Named Access List)`,
				},
				resource.Attribute{
					Name:        "acl_name_seq_num",
					Description: `Specify ACL precedence (sequence-number)`,
				},
				resource.Attribute{
					Name:        "acl_name_seq_num_shared",
					Description: `Specify ACL precedence (sequence-number)`,
				},
				resource.Attribute{
					Name:        "acl_name_src_nat_pool",
					Description: `Policy based Source NAT (NAT Pool or Pool Group)`,
				},
				resource.Attribute{
					Name:        "acl_name_src_nat_pool_shared",
					Description: `Policy based Source NAT (NAT Pool or Pool Group)`,
				},
				resource.Attribute{
					Name:        "shared_partition_pool_name",
					Description: `Policy based Source NAT from shared partition`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘curr_conn’: Current connections; ‘total_l4_conn’: Total L4 connections; ‘total_l7_conn’: Total L7 connections; ‘total_tcp_conn’: Total TCP connections; ‘total_conn’: Total connections; ‘total_fwd_bytes’: Total forward bytes; ‘total_fwd_pkts’: Total forward packets; ‘total_rev_bytes’: Total reverse bytes; ‘total_rev_pkts’: Total reverse packets; ‘total_dns_pkts’: Total DNS packets; ‘total_mf_dns_pkts’: Total MF DNS packets; ‘es_total_failure_actions’: Total failure actions; ‘compression_bytes_before’: Data into compression engine; ‘compression_bytes_after’: Data out of compression engine; ‘compression_hit’: Number of requests compressed; ‘compression_miss’: Number of requests NOT compressed; ‘compression_miss_no_client’: Compression miss no client; ‘compression_miss_template_exclusion’: Compression miss template exclusion; ‘curr_req’: Current requests; ‘total_req’: Total requests; ‘total_req_succ’: Total successful requests; ‘peak_conn’: Peak connections; ‘curr_conn_rate’: Current connection rate; ‘last_rsp_time’: Last response time; ‘fastest_rsp_time’: Fastest response time; ‘slowest_rsp_time’: Slowest response time; ‘loc_permit’: Permit number; ‘loc_deny’: Deny number; ‘loc_conn’: Connection number; ‘curr_ssl_conn’: Current SSL connections; ‘total_ssl_conn’: Total SSL connections;`,
				},
				resource.Attribute{
					Name:        "acl_id",
					Description: `ACL id VPORT`,
				},
				resource.Attribute{
					Name:        "acl_id_seq_num",
					Description: `Specify ACL precedence (sequence-number)`,
				},
				resource.Attribute{
					Name:        "acl_id_seq_num_shared",
					Description: `Specify ACL precedence (sequence-number)`,
				},
				resource.Attribute{
					Name:        "acl_id_src_nat_pool",
					Description: `Policy based Source NAT (NAT Pool or Pool Group)`,
				},
				resource.Attribute{
					Name:        "acl_id_src_nat_pool_shared",
					Description: `Policy based Source NAT (NAT Pool or Pool Group)`,
				},
				resource.Attribute{
					Name:        "shared_partition_pool_id",
					Description: `Policy based Source NAT from shared partition`,
				},
				resource.Attribute{
					Name:        "aflex",
					Description: `Bind aFleX Script to the Virtual Port (aFleX Script Name)`,
				},
				resource.Attribute{
					Name:        "aflex_shared",
					Description: `aFleX Script Name`,
				},
				resource.Attribute{
					Name:        "aaa_policy",
					Description: `Specify AAA policy name to bind to the virtual port`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_SNMPv1_v2c_user",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server SNMPv1 v2c user resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED community string)`,
				},
				resource.Attribute{
					Name:        "passwd",
					Description: `define value of community string`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `SNMPv1/v2c community configuration key`,
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
					Name:        "dns_host",
					Description: `DNS remote host`,
				},
				resource.Attribute{
					Name:        "ipv4_mask",
					Description: `IPV4 mask`,
				},
				resource.Attribute{
					Name:        "ipv4_host",
					Description: `IPV4 remote host`,
				},
				resource.Attribute{
					Name:        "ipv6_host",
					Description: `IPV6 remote host`,
				},
				resource.Attribute{
					Name:        "ipv6_mask",
					Description: `IPV6 mask`,
				},
				resource.Attribute{
					Name:        "oid_val",
					Description: `specific the oid (The oid value, object-key)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_SNMPv1_v2c_user_oid",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server SNMPv1 v2c user oid resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "oid_val",
					Description: `specific the oid (The oid value, object-key)`,
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
					Name:        "dns_host",
					Description: `DNS remote host`,
				},
				resource.Attribute{
					Name:        "ipv4_mask",
					Description: `IPV4 mask`,
				},
				resource.Attribute{
					Name:        "ipv4_host",
					Description: `IPV4 remote host`,
				},
				resource.Attribute{
					Name:        "ipv6_host",
					Description: `IPV6 remote host`,
				},
				resource.Attribute{
					Name:        "ipv6_mask",
					Description: `IPV6 mask`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_SNMPv3_user",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server SNMPv3 user resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auth_val",
					Description: `‘md5’: Use HMAC MD5 algorithm for authentication; ‘sha’: Use HMAC SHA algorithm for authentication;`,
				},
				resource.Attribute{
					Name:        "encpasswd",
					Description: `Passphrase for encryption`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `Group to which the user belongs`,
				},
				resource.Attribute{
					Name:        "passwd",
					Description: `Password of this user`,
				},
				resource.Attribute{
					Name:        "priv",
					Description: `‘des’: DES encryption alogrithm; ‘aes’: AES encryption alogrithm; (Encryption type)`,
				},
				resource.Attribute{
					Name:        "priv_pw_encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED passphrase string)`,
				},
				resource.Attribute{
					Name:        "pw_encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED passphrase string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Name of the user`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "v3",
					Description: `‘auth’: Using the authNoPriv Security Level; ‘noauth’: Using the noAuthNoPriv Security Level;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_community_read",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server community read resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user",
					Description: `SNMPv1/v2c community string`,
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
					Name:        "oid_val",
					Description: `specific the oid (The oid value, object-key)`,
				},
				resource.Attribute{
					Name:        "dns_host",
					Description: `DNS remote host`,
				},
				resource.Attribute{
					Name:        "ipv4_mask",
					Description: `IPV4 mask`,
				},
				resource.Attribute{
					Name:        "ipv4_host",
					Description: `IPV4 remote host`,
				},
				resource.Attribute{
					Name:        "ipv6_host",
					Description: `IPV6 remote host`,
				},
				resource.Attribute{
					Name:        "ipv6_mask",
					Description: `IPV6 mask`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_community_read_oid",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server community read oid resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "oid_val",
					Description: `specific the oid (The oid value, object-key)`,
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
					Name:        "dns_host",
					Description: `DNS remote host`,
				},
				resource.Attribute{
					Name:        "ipv4_mask",
					Description: `IPV4 mask`,
				},
				resource.Attribute{
					Name:        "ipv4_host",
					Description: `IPV4 remote host`,
				},
				resource.Attribute{
					Name:        "ipv6_host",
					Description: `IPV6 remote host`,
				},
				resource.Attribute{
					Name:        "ipv6_mask",
					Description: `IPV6 mask`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_contact",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server contact resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "contact_name",
					Description: `Identification of the contact person for this managed node`,
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
			Type:             "vthunder_thunder_snmp_server_disable_traps",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server disable traps resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all",
					Description: `Disable all traps on this partition`,
				},
				resource.Attribute{
					Name:        "gslb",
					Description: `Disable all gslb traps on this partition`,
				},
				resource.Attribute{
					Name:        "slb",
					Description: `Disable all slb traps on this partition`,
				},
				resource.Attribute{
					Name:        "slb_change",
					Description: `Disable all slb-change traps on this partition`,
				},
				resource.Attribute{
					Name:        "snmp",
					Description: `Disable all snmp traps on this partition`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "vrrp_a",
					Description: `Disable all vrrp-a on this partition`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_enable_traps",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server enable traps resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all",
					Description: `Enable all SLB traps`,
				},
				resource.Attribute{
					Name:        "lldp",
					Description: `Enable lldp traps`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "connection_resource_event",
					Description: `Enable system connection resource event trap`,
				},
				resource.Attribute{
					Name:        "resource_usage_warning",
					Description: `Enable partition resource usage warning trap`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Enable slb server create/delete trap`,
				},
				resource.Attribute{
					Name:        "server_port",
					Description: `Enable slb server port create/delete trap`,
				},
				resource.Attribute{
					Name:        "ssl_cert_change",
					Description: `Enable SSL certificate change trap`,
				},
				resource.Attribute{
					Name:        "ssl_cert_expire",
					Description: `Enable SSL certificate expiring trap`,
				},
				resource.Attribute{
					Name:        "system_threshold",
					Description: `Enable slb system threshold trap`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `Enable slb vip create/delete trap`,
				},
				resource.Attribute{
					Name:        "vip_port",
					Description: `Enable slb vip-port create/delete trap`,
				},
				resource.Attribute{
					Name:        "fixed_nat_port_mapping_file_change",
					Description: `Enable LSN trap when fixed nat port mapping file change`,
				},
				resource.Attribute{
					Name:        "max_ipport_threshold",
					Description: `Maximum threshold`,
				},
				resource.Attribute{
					Name:        "max_port_threshold",
					Description: `Maximum threshold`,
				},
				resource.Attribute{
					Name:        "per_ip_port_usage_threshold",
					Description: `Enable LSN trap when IP total port usage reaches the threshold (default 64512)`,
				},
				resource.Attribute{
					Name:        "total_port_usage_threshold",
					Description: `Enable LSN trap when NAT total port usage reaches the threshold (default 655350000)`,
				},
				resource.Attribute{
					Name:        "traffic_exceeded",
					Description: `Enable LSN trap when NAT pool reaches the threshold`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `Enable VRRP-A active trap`,
				},
				resource.Attribute{
					Name:        "standby",
					Description: `Enable VRRP-A standby trap`,
				},
				resource.Attribute{
					Name:        "linkdown",
					Description: `Enable SNMP link-down trap`,
				},
				resource.Attribute{
					Name:        "linkup",
					Description: `Enable SNMP link-up trap`,
				},
				resource.Attribute{
					Name:        "control_cpu_high",
					Description: `Enable control CPU usage high trap`,
				},
				resource.Attribute{
					Name:        "data_cpu_high",
					Description: `Enable data CPU usage high trap`,
				},
				resource.Attribute{
					Name:        "fan",
					Description: `Enable system fan trap`,
				},
				resource.Attribute{
					Name:        "file_sys_read_only",
					Description: `Enable file system read-only trap`,
				},
				resource.Attribute{
					Name:        "high_disk_use",
					Description: `Enable system high disk usage trap`,
				},
				resource.Attribute{
					Name:        "high_memory_use",
					Description: `Enable system high memory usage trap`,
				},
				resource.Attribute{
					Name:        "high_temp",
					Description: `Enable system high temperature trap`,
				},
				resource.Attribute{
					Name:        "license_management",
					Description: `Enable system license management traps`,
				},
				resource.Attribute{
					Name:        "low_temp",
					Description: `Enable system low temperature trap`,
				},
				resource.Attribute{
					Name:        "packet_drop",
					Description: `Enable system packet dropped trap`,
				},
				resource.Attribute{
					Name:        "power",
					Description: `Enable system power supply trap`,
				},
				resource.Attribute{
					Name:        "pri_disk",
					Description: `Enable system primary hard disk trap`,
				},
				resource.Attribute{
					Name:        "restart",
					Description: `Enable system restart trap`,
				},
				resource.Attribute{
					Name:        "sec_disk",
					Description: `Enable system secondary hard disk trap`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `Enable system shutdown trap`,
				},
				resource.Attribute{
					Name:        "smp_resource_event",
					Description: `Enable system smp resource event trap`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `Enable system start trap`,
				},
				resource.Attribute{
					Name:        "tacacs_server_up_down",
					Description: `Enable system TACACS monitor server up/down trap`,
				},
				resource.Attribute{
					Name:        "server_certificate_error",
					Description: `Enable SSL server certificate error trap`,
				},
				resource.Attribute{
					Name:        "state_change",
					Description: `Enable VCS state change trap`,
				},
				resource.Attribute{
					Name:        "bgp_backward_trans_notification",
					Description: `Enable bgpBackwardTransNotification traps`,
				},
				resource.Attribute{
					Name:        "bgp_established_notification",
					Description: `Enable bgp_established_notification traps`,
				},
				resource.Attribute{
					Name:        "ospf_if_auth_failure",
					Description: `Enable ospfIfAuthFailure traps`,
				},
				resource.Attribute{
					Name:        "ospf_if_config_error",
					Description: `Enable ospfIfConfigError traps`,
				},
				resource.Attribute{
					Name:        "ospf_if_rx_bad_packet",
					Description: `Enable ospfIfRxBadPacket traps`,
				},
				resource.Attribute{
					Name:        "ospf_if_state_change",
					Description: `Enable ospfIfStateChange traps`,
				},
				resource.Attribute{
					Name:        "ospf_lsdb_approaching_overflow",
					Description: `Enable ospfLsdbApproachingOverflow traps`,
				},
				resource.Attribute{
					Name:        "ospf_lsdb_overflow",
					Description: `Enable ospfLsdbOverflow traps`,
				},
				resource.Attribute{
					Name:        "ospf_max_age_lsa",
					Description: `Enable ospfMaxAgeLsa traps`,
				},
				resource.Attribute{
					Name:        "ospf_nbr_state_change",
					Description: `Enable ospfNbrStateChange traps`,
				},
				resource.Attribute{
					Name:        "ospf_originate_lsa",
					Description: `Enable ospfOriginateLsa traps`,
				},
				resource.Attribute{
					Name:        "ospf_tx_retransmit",
					Description: `Enable ospfTxRetransmit traps`,
				},
				resource.Attribute{
					Name:        "ospf_virt_if_auth_failure",
					Description: `Enable ospfVirtIfAuthFailure traps`,
				},
				resource.Attribute{
					Name:        "ospf_virt_if_config_error",
					Description: `Enable ospfVirtIfConfigError traps`,
				},
				resource.Attribute{
					Name:        "ospf_virt_if_rx_bad_packet",
					Description: `Enable ospfVirtIfRxBadPacket traps`,
				},
				resource.Attribute{
					Name:        "ospf_virt_if_state_change",
					Description: `Enable ospfVirtIfStateChange traps`,
				},
				resource.Attribute{
					Name:        "ospf_virt_if_tx_retransmit",
					Description: `Enable ospfVirtIfTxRetransmit traps`,
				},
				resource.Attribute{
					Name:        "ospf_virt_nbr_state_change",
					Description: `Enable ospfVirtNbrStateChange traps`,
				},
				resource.Attribute{
					Name:        "isis_adjacency_change",
					Description: `Enable isisAdjacencyChange traps`,
				},
				resource.Attribute{
					Name:        "isis_area_mismatch",
					Description: `Enable isisAreaMismatch traps`,
				},
				resource.Attribute{
					Name:        "isis_attempt_to_exceed_max_sequence",
					Description: `Enable isisAttemptToExceedMaxSequence traps`,
				},
				resource.Attribute{
					Name:        "isis_authentication_failure",
					Description: `Enable isisAuthenticationFailure traps`,
				},
				resource.Attribute{
					Name:        "isis_authentication_type_failure",
					Description: `Enable isisAuthenticationTypeFailure traps`,
				},
				resource.Attribute{
					Name:        "isis_corrupted_lsp_detected",
					Description: `Enable isisCorruptedLSPDetected traps`,
				},
				resource.Attribute{
					Name:        "isis_database_overload",
					Description: `Enable isisDatabaseOverload traps`,
				},
				resource.Attribute{
					Name:        "isis_id_len_mismatch",
					Description: `Enable isisIDLenMismatch traps`,
				},
				resource.Attribute{
					Name:        "isis_lsp_too_large_to_propagate",
					Description: `Enable isisLSPTooLargeToPropagate traps`,
				},
				resource.Attribute{
					Name:        "isis_manual_address_drops",
					Description: `Enable isisManualAddressDrops traps`,
				},
				resource.Attribute{
					Name:        "isis_max_area_addresses_mismatch",
					Description: `Enable isisMaxAreaAddressesMismatch traps`,
				},
				resource.Attribute{
					Name:        "isis_originating_lsp_buffer_size_mismatch",
					Description: `Enable isisOriginatingLSPBufferSizeMismatch traps`,
				},
				resource.Attribute{
					Name:        "isis_own_lsp_purge",
					Description: `Enable isisOwnLSPPurge traps`,
				},
				resource.Attribute{
					Name:        "isis_protocols_supported_mismatch",
					Description: `Enable isisProtocolsSupportedMismatch traps`,
				},
				resource.Attribute{
					Name:        "isis_rejected_adjacency",
					Description: `Enable isisRejectedAdjacency traps`,
				},
				resource.Attribute{
					Name:        "isis_sequence_number_skip",
					Description: `Enable isisSequenceNumberSkip traps`,
				},
				resource.Attribute{
					Name:        "isis_version_skew",
					Description: `Enable isisVersionSkew traps`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `Enable GSLB group related traps`,
				},
				resource.Attribute{
					Name:        "service_ip",
					Description: `Enable GSLB service-ip related traps`,
				},
				resource.Attribute{
					Name:        "site",
					Description: `Enable GSLB site related traps`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Enable GSLB zone related traps`,
				},
				resource.Attribute{
					Name:        "application_buffer_limit",
					Description: `Enable application buffer reach limit trap`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_exceed",
					Description: `Enable SLB server/port bandwidth rate limit exceed trap`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_resume",
					Description: `Enable SLB server/port bandwidth rate limit resume trap`,
				},
				resource.Attribute{
					Name:        "gateway_down",
					Description: `Enable SLB server gateway down trap`,
				},
				resource.Attribute{
					Name:        "gateway_up",
					Description: `Enable SLB server gateway up trap`,
				},
				resource.Attribute{
					Name:        "server_conn_limit",
					Description: `Enable SLB server connection limit trap`,
				},
				resource.Attribute{
					Name:        "server_conn_resume",
					Description: `Enable SLB server connection resume trap`,
				},
				resource.Attribute{
					Name:        "server_disabled",
					Description: `Enable SLB server-disabled trap`,
				},
				resource.Attribute{
					Name:        "server_down",
					Description: `Enable SLB server-down trap`,
				},
				resource.Attribute{
					Name:        "server_selection_failure",
					Description: `Enable SLB server selection failure trap`,
				},
				resource.Attribute{
					Name:        "server_up",
					Description: `Enable slb server up trap`,
				},
				resource.Attribute{
					Name:        "service_conn_limit",
					Description: `Enable SLB service connection limit trap`,
				},
				resource.Attribute{
					Name:        "service_conn_resume",
					Description: `Enable SLB service connection resume trap`,
				},
				resource.Attribute{
					Name:        "service_down",
					Description: `Enable SLB service-down trap`,
				},
				resource.Attribute{
					Name:        "service_group_down",
					Description: `Enable SLB service-group-down trap`,
				},
				resource.Attribute{
					Name:        "service_group_member_down",
					Description: `Enable SLB service-group-member-down trap`,
				},
				resource.Attribute{
					Name:        "service_group_member_up",
					Description: `Enable SLB service-group-member-up trap`,
				},
				resource.Attribute{
					Name:        "service_group_up",
					Description: `Enable SLB service-group-up trap`,
				},
				resource.Attribute{
					Name:        "service_up",
					Description: `Enable SLB service-up trap`,
				},
				resource.Attribute{
					Name:        "vip_connlimit",
					Description: `Enable the virtual server reach conn-limit trap`,
				},
				resource.Attribute{
					Name:        "vip_connratelimit",
					Description: `Enable the virtual server reach conn-rate-limit trap`,
				},
				resource.Attribute{
					Name:        "vip_down",
					Description: `Enable SLB virtual server down trap`,
				},
				resource.Attribute{
					Name:        "vip_port_connlimit",
					Description: `Enable the virtual port reach conn-limit trap`,
				},
				resource.Attribute{
					Name:        "vip_port_connratelimit",
					Description: `Enable the virtual port reach conn-rate-limit trap`,
				},
				resource.Attribute{
					Name:        "vip_port_down",
					Description: `Enable SLB virtual port down trap`,
				},
				resource.Attribute{
					Name:        "vip_port_up",
					Description: `Enable SLB virtual port up trap`,
				},
				resource.Attribute{
					Name:        "vip_up",
					Description: `Enable SLB virtual server up trap`,
				},
				resource.Attribute{
					Name:        "trunk_port_threshold",
					Description: `Enable network trunk-port-threshold trap`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_enable_traps_gslb",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server enable traps gslb resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all",
					Description: `Enable all GSLB traps`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `Enable GSLB group related traps`,
				},
				resource.Attribute{
					Name:        "service_ip",
					Description: `Enable GSLB service-ip related traps`,
				},
				resource.Attribute{
					Name:        "site",
					Description: `Enable GSLB site related traps`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Enable GSLB zone related traps`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_enable_traps_lsn",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server enable traps lsn resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all",
					Description: `Enable all LSN group traps`,
				},
				resource.Attribute{
					Name:        "fixed_nat_port_mapping_file_change",
					Description: `Enable LSN trap when fixed nat port mapping file change`,
				},
				resource.Attribute{
					Name:        "max_ipport_threshold",
					Description: `Maximum threshold`,
				},
				resource.Attribute{
					Name:        "max_port_threshold",
					Description: `Maximum threshold`,
				},
				resource.Attribute{
					Name:        "per_ip_port_usage_threshold",
					Description: `Enable LSN trap when IP total port usage reaches the threshold (default 64512)`,
				},
				resource.Attribute{
					Name:        "total_port_usage_threshold",
					Description: `Enable LSN trap when NAT total port usage reaches the threshold (default 655350000)`,
				},
				resource.Attribute{
					Name:        "traffic_exceeded",
					Description: `Enable LSN trap when NAT pool reaches the threshold`,
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
			Type:             "vthunder_thunder_snmp_server_enable_traps_network",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server enable traps network resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "trunk_port_threshold",
					Description: `Enable network trunk-port-threshold trap`,
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
			Type:             "vthunder_thunder_snmp_server_enable_traps_routing_bgp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server enable traps routing bgp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bgp_backward_trans_notification",
					Description: `Enable bgpBackwardTransNotification traps`,
				},
				resource.Attribute{
					Name:        "bgp_established_notification",
					Description: `Enable bgpEstablishedNotification traps`,
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
			Type:             "vthunder_thunder_snmp_server_enable_traps_routing_isis",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server enable traps routing isis resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "isis_adjacency_change",
					Description: `Enable isisAdjacencyChange traps`,
				},
				resource.Attribute{
					Name:        "isis_area_mismatch",
					Description: `Enable isisAreaMismatch traps`,
				},
				resource.Attribute{
					Name:        "isis_attempt_to_exceed_max_sequence",
					Description: `Enable isisAttemptToExceedMaxSequence traps`,
				},
				resource.Attribute{
					Name:        "isis_authentication_failure",
					Description: `Enable isisAuthenticationFailure traps`,
				},
				resource.Attribute{
					Name:        "isis_authentication_type_failure",
					Description: `Enable isisAuthenticationTypeFailure traps`,
				},
				resource.Attribute{
					Name:        "isis_corrupted_lsp_detected",
					Description: `Enable isisCorruptedLSPDetected traps`,
				},
				resource.Attribute{
					Name:        "isis_database_overload",
					Description: `Enable isisDatabaseOverload traps`,
				},
				resource.Attribute{
					Name:        "isis_id_len_mismatch",
					Description: `Enable isisIDLenMismatch traps`,
				},
				resource.Attribute{
					Name:        "isis_lsp_too_large_to_propagate",
					Description: `Enable isisLSPTooLargeToPropagate traps`,
				},
				resource.Attribute{
					Name:        "isis_manual_address_drops",
					Description: `Enable isisManualAddressDrops traps`,
				},
				resource.Attribute{
					Name:        "isis_max_area_addresses_mismatch",
					Description: `Enable isisMaxAreaAddressesMismatch traps`,
				},
				resource.Attribute{
					Name:        "isis_originating_lsp_buffer_size_mismatch",
					Description: `Enable isisOriginatingLSPBufferSizeMismatch traps`,
				},
				resource.Attribute{
					Name:        "isis_own_lsp_purge",
					Description: `Enable isisOwnLSPPurge traps`,
				},
				resource.Attribute{
					Name:        "isis_protocols_supported_mismatch",
					Description: `Enable isisProtocolsSupportedMismatch traps`,
				},
				resource.Attribute{
					Name:        "isis_rejected_adjacency",
					Description: `Enable isisRejectedAdjacency traps`,
				},
				resource.Attribute{
					Name:        "isis_sequence_number_skip",
					Description: `Enable isisSequenceNumberSkip traps`,
				},
				resource.Attribute{
					Name:        "isis_version_skew",
					Description: `Enable isisVersionSkew traps`,
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
			Type:             "vthunder_thunder_snmp_server_enable_traps_routing_ospf",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server enable traps routing ospf resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ospf_if_auth_failure",
					Description: `Enable ospfIfAuthFailure traps`,
				},
				resource.Attribute{
					Name:        "ospf_if_config_error",
					Description: `Enable ospfIfConfigError traps`,
				},
				resource.Attribute{
					Name:        "ospf_if_rx_bad_packet",
					Description: `Enable ospfIfRxBadPacket traps`,
				},
				resource.Attribute{
					Name:        "ospf_if_state_change",
					Description: `Enable ospfIfStateChange traps`,
				},
				resource.Attribute{
					Name:        "ospf_lsdb_approaching_overflow",
					Description: `Enable ospfLsdbApproachingOverflow traps`,
				},
				resource.Attribute{
					Name:        "ospf_lsdb_overflow",
					Description: `Enable ospfLsdbOverflow traps`,
				},
				resource.Attribute{
					Name:        "ospf_max_age_lsa",
					Description: `Enable ospfMaxAgeLsa traps`,
				},
				resource.Attribute{
					Name:        "ospf_nbr_state_change",
					Description: `Enable ospfNbrStateChange traps`,
				},
				resource.Attribute{
					Name:        "ospf_originate_lsa",
					Description: `Enable ospfOriginateLsa traps`,
				},
				resource.Attribute{
					Name:        "ospf_tx_retransmit",
					Description: `Enable ospfTxRetransmit traps`,
				},
				resource.Attribute{
					Name:        "ospf_virt_if_auth_failure",
					Description: `Enable ospfVirtIfAuthFailure traps`,
				},
				resource.Attribute{
					Name:        "ospf_virt_if_config_error",
					Description: `Enable ospfVirtIfConfigError traps`,
				},
				resource.Attribute{
					Name:        "ospf_virt_if_rx_bad_packet",
					Description: `Enable ospfVirtIfRxBadPacket traps`,
				},
				resource.Attribute{
					Name:        "ospf_virt_if_state_change",
					Description: `Enable ospfVirtIfStateChange traps`,
				},
				resource.Attribute{
					Name:        "ospf_virt_if_tx_retransmit",
					Description: `Enable ospfVirtIfTxRetransmit traps`,
				},
				resource.Attribute{
					Name:        "ospf_virt_nbr_state_change",
					Description: `Enable ospfVirtNbrStateChange traps`,
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
			Type:             "vthunder_thunder_snmp_server_enable_traps_slb",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server enable traps slb resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all",
					Description: `Enable all SLB traps`,
				},
				resource.Attribute{
					Name:        "application_buffer_limit",
					Description: `Enable application buffer reach limit trap`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_exceed",
					Description: `Enable SLB server/port bandwidth rate limit exceed trap`,
				},
				resource.Attribute{
					Name:        "bw_rate_limit_resume",
					Description: `Enable SLB server/port bandwidth rate limit resume trap`,
				},
				resource.Attribute{
					Name:        "gateway_down",
					Description: `Enable SLB server gateway down trap`,
				},
				resource.Attribute{
					Name:        "gateway_up",
					Description: `Enable SLB server gateway up trap`,
				},
				resource.Attribute{
					Name:        "server_conn_limit",
					Description: `Enable SLB server connection limit trap`,
				},
				resource.Attribute{
					Name:        "server_conn_resume",
					Description: `Enable SLB server connection resume trap`,
				},
				resource.Attribute{
					Name:        "server_disabled",
					Description: `Enable SLB server-disabled trap`,
				},
				resource.Attribute{
					Name:        "server_down",
					Description: `Enable SLB server-down trap`,
				},
				resource.Attribute{
					Name:        "server_selection_failure",
					Description: `Enable SLB server selection failure trap`,
				},
				resource.Attribute{
					Name:        "server_up",
					Description: `Enable slb server up trap`,
				},
				resource.Attribute{
					Name:        "service_conn_limit",
					Description: `Enable SLB service connection limit trap`,
				},
				resource.Attribute{
					Name:        "service_conn_resume",
					Description: `Enable SLB service connection resume trap`,
				},
				resource.Attribute{
					Name:        "service_down",
					Description: `Enable SLB service-down trap`,
				},
				resource.Attribute{
					Name:        "service_group_down",
					Description: `Enable SLB service-group-down trap`,
				},
				resource.Attribute{
					Name:        "service_group_member_down",
					Description: `Enable SLB service-group-member-down trap`,
				},
				resource.Attribute{
					Name:        "service_group_member_up",
					Description: `Enable SLB service-group-member-up trap`,
				},
				resource.Attribute{
					Name:        "service_group_up",
					Description: `Enable SLB service-group-up trap`,
				},
				resource.Attribute{
					Name:        "service_up",
					Description: `Enable SLB service-up trap`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "vip_connlimit",
					Description: `Enable the virtual server reach conn-limit trap`,
				},
				resource.Attribute{
					Name:        "vip_connratelimit",
					Description: `Enable the virtual server reach conn-rate-limit trap`,
				},
				resource.Attribute{
					Name:        "vip_down",
					Description: `Enable SLB virtual server down trap`,
				},
				resource.Attribute{
					Name:        "vip_port_connlimit",
					Description: `Enable the virtual port reach conn-limit trap`,
				},
				resource.Attribute{
					Name:        "vip_port_connratelimit",
					Description: `Enable the virtual port reach conn-rate-limit trap`,
				},
				resource.Attribute{
					Name:        "vip_port_down",
					Description: `Enable SLB virtual port down trap`,
				},
				resource.Attribute{
					Name:        "vip_port_up",
					Description: `Enable SLB virtual port up trap`,
				},
				resource.Attribute{
					Name:        "vip_up",
					Description: `Enable SLB virtual server up trap`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_enable_traps_slb_change",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server enable traps slb change resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all",
					Description: `Enable all system group traps`,
				},
				resource.Attribute{
					Name:        "connection_resource_event",
					Description: `Enable system connection resource event trap`,
				},
				resource.Attribute{
					Name:        "resource_usage_warning",
					Description: `Enable partition resource usage warning trap`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Enable slb server create/delete trap`,
				},
				resource.Attribute{
					Name:        "server_port",
					Description: `Enable slb server port create/delete trap`,
				},
				resource.Attribute{
					Name:        "ssl_cert_change",
					Description: `Enable SSL certificate change trap`,
				},
				resource.Attribute{
					Name:        "ssl_cert_expire",
					Description: `Enable SSL certificate expiring trap`,
				},
				resource.Attribute{
					Name:        "system_threshold",
					Description: `Enable slb system threshold trap`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `Enable slb vip create/delete trap`,
				},
				resource.Attribute{
					Name:        "vip_port",
					Description: `Enable slb vip-port create/delete trap`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_enable_traps_snmp",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server enable traps snmp resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all",
					Description: `Enable all SNMP group traps`,
				},
				resource.Attribute{
					Name:        "linkdown",
					Description: `Enable SNMP link-down trap`,
				},
				resource.Attribute{
					Name:        "linkup",
					Description: `Enable SNMP link-up trap`,
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
			Type:             "vthunder_thunder_snmp_server_enable_traps_ssl",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server enable traps ssl resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server_certificate_error",
					Description: `Enable SSL server certificate error trap`,
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
			Type:             "vthunder_thunder_snmp_server_enable_traps_system",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server enable traps system resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all",
					Description: `Enable all system group traps`,
				},
				resource.Attribute{
					Name:        "control_cpu_high",
					Description: `Enable control CPU usage high trap`,
				},
				resource.Attribute{
					Name:        "data_cpu_high",
					Description: `Enable data CPU usage high trap`,
				},
				resource.Attribute{
					Name:        "fan",
					Description: `Enable system fan trap`,
				},
				resource.Attribute{
					Name:        "file_sys_read_only",
					Description: `Enable file system read-only trap`,
				},
				resource.Attribute{
					Name:        "high_disk_use",
					Description: `Enable system high disk usage trap`,
				},
				resource.Attribute{
					Name:        "high_memory_use",
					Description: `Enable system high memory usage trap`,
				},
				resource.Attribute{
					Name:        "high_temp",
					Description: `Enable system high temperature trap`,
				},
				resource.Attribute{
					Name:        "license_management",
					Description: `Enable system license management traps`,
				},
				resource.Attribute{
					Name:        "low_temp",
					Description: `Enable system low temperature trap`,
				},
				resource.Attribute{
					Name:        "packet_drop",
					Description: `Enable system packet dropped trap`,
				},
				resource.Attribute{
					Name:        "power",
					Description: `Enable system power supply trap`,
				},
				resource.Attribute{
					Name:        "pri_disk",
					Description: `Enable system primary hard disk trap`,
				},
				resource.Attribute{
					Name:        "restart",
					Description: `Enable system restart trap`,
				},
				resource.Attribute{
					Name:        "sec_disk",
					Description: `Enable system secondary hard disk trap`,
				},
				resource.Attribute{
					Name:        "shutdown",
					Description: `Enable system shutdown trap`,
				},
				resource.Attribute{
					Name:        "smp_resource_event",
					Description: `Enable system smp resource event trap`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `Enable system start trap`,
				},
				resource.Attribute{
					Name:        "tacacs_server_up_down",
					Description: `Enable system TACACS monitor server up/down trap`,
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
			Type:             "vthunder_thunder_snmp_server_enable_traps_vcs",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server enable traps vcs resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "state_change",
					Description: `Enable VCS state change trap`,
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
			Type:             "vthunder_thunder_snmp_server_enable_traps_vrrp_a",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server enable traps vrrp a resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "active",
					Description: `Enable VRRP-A active trap`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `Enable all VRRP-A group traps`,
				},
				resource.Attribute{
					Name:        "standby",
					Description: `Enable VRRP-A standby trap`,
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
			Type:             "vthunder_thunder_snmp_server_engineID",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server engineID resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engId",
					Description: `Define local engineID HEX WORD`,
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
			Type:             "vthunder_thunder_snmp_server_group",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server group resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "groupname",
					Description: `Name of the group`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `specify a read view for the group (read view name)`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "v3",
					Description: `‘auth’: group using the authNoPriv Security Level; ‘noauth’: group using the noAuthNoPriv Security Level; ‘priv’: group using SNMPv3 authPriv security level;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_host_host_name",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server host host name resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname of SNMP trap host`,
				},
				resource.Attribute{
					Name:        "udp_port",
					Description: `The trap host’s UDP port number(default: 162)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `SNMPv3 user to send traps (User Name)`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "v1_v2c_comm",
					Description: `SNMPv1/v2c community string`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `‘v1’: Use SNMPv1; ‘v2c’: Use SNMPv2c; ‘v3’: User SNMPv3;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_host_ipv4_host",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server host ipv4 host resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv4_addr",
					Description: `help IPV4 address of SNMP trap host`,
				},
				resource.Attribute{
					Name:        "udp_port",
					Description: `The trap host’s UDP port number(default: 162)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `SNMPv3 user to send traps (User Name)`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "v1_v2c_comm",
					Description: `SNMPv1/v2c community string`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `‘v1’: Use SNMPv1; ‘v2c’: Use SNMPv2c; ‘v3’: User SNMPv3;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_host_ipv6_host",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server host ipv6 host resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv6_addr",
					Description: `help IPV6 address of SNMP trap host`,
				},
				resource.Attribute{
					Name:        "udp_port",
					Description: `The trap host’s UDP port number(default: 162)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `SNMPv3 user to send traps (User Name)`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "v1_v2c_comm",
					Description: `SNMPv1/v2c community string`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `‘v1’: Use SNMPv1; ‘v2c’: Use SNMPv2c; ‘v3’: User SNMPv3;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_location",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server location resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "loc",
					Description: `The physical location of this node`,
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
			Type:             "vthunder_thunder_snmp_server_management_index",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server management index resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mgmt_index",
					Description: `index number`,
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
			Type:             "vthunder_thunder_snmp_server_slb_data_cache_timeout",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server slb data cache timeout resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "slblimit",
					Description: `Cache time-out in seconds, default as 60 seconds`,
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
			Type:             "vthunder_thunder_snmp_server_user",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server user resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auth_val",
					Description: `‘md5’: Use HMAC MD5 algorithm for authentication; ‘sha’: Use HMAC SHA algorithm for authentication;`,
				},
				resource.Attribute{
					Name:        "encpasswd",
					Description: `Passphrase for encryption`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `Group to which the user belongs`,
				},
				resource.Attribute{
					Name:        "passwd",
					Description: `Password of this user`,
				},
				resource.Attribute{
					Name:        "priv",
					Description: `‘des’: DES encryption alogrithm; ‘aes’: AES encryption alogrithm; (Encryption type)`,
				},
				resource.Attribute{
					Name:        "priv_pw_encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED passphrase string)`,
				},
				resource.Attribute{
					Name:        "pw_encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED passphrase string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Name of the user`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "v3",
					Description: `‘auth’: Using the authNoPriv Security Level; ‘noauth’: Using the noAuthNoPriv Security Level;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_snmp_server_view",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder snmp server view resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mask",
					Description: `Hex octets, separated by ‘.’, mask of oid`,
				},
				resource.Attribute{
					Name:        "oid",
					Description: `MIB view family name or oid`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `‘included’: MIB family is included in the view; ‘excluded’: MIB family is excluded from the view;`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "viewname",
					Description: `Name of the view`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_system",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder system resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "system",
					Description: `Configure System Parameters`,
				},
				resource.Attribute{
					Name:        "anomaly-log",
					Description: `log system anomalies`,
				},
				resource.Attribute{
					Name:        "attack-log",
					Description: `log attack anomalies`,
				},
				resource.Attribute{
					Name:        "ddos-attack",
					Description: `System DDoS Attack`,
				},
				resource.Attribute{
					Name:        "ddos-log",
					Description: `log DDoS attack anomalies`,
				},
				resource.Attribute{
					Name:        "sockstress-disable",
					Description: `Disable sockstress protection`,
				},
				resource.Attribute{
					Name:        "promiscuous-mode",
					Description: `Run in promiscous mode settings`,
				},
				resource.Attribute{
					Name:        "glid",
					Description: `Apply limits to the whole system`,
				},
				resource.Attribute{
					Name:        "module-ctrl-cpu",
					Description: `'high': high cpu usage; 'low': low cpu usage; 'medium': medium cpu usage;`,
				},
				resource.Attribute{
					Name:        "src-ip-hash-enable",
					Description: `Enable source ip hash`,
				},
				resource.Attribute{
					Name:        "class-list-hitcount-enable",
					Description: `Enable class list hit count`,
				},
				resource.Attribute{
					Name:        "geo-db-hitcount-enable",
					Description: `Enable Geolocation database hit count`,
				},
				resource.Attribute{
					Name:        "domain-list-hitcount-enable",
					Description: `Enable class list hit count`,
				},
				resource.Attribute{
					Name:        "dynamic-service-dns-socket-pool",
					Description: `Enable socket pool for dynamic-service DNS`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "ftp",
					Description: `set timeout to stop ftp transfer in seconds, 0 is no limit`,
				},
				resource.Attribute{
					Name:        "scp",
					Description: `set timeout to stop scp transfer in seconds, 0 is no limit`,
				},
				resource.Attribute{
					Name:        "sftp",
					Description: `set timeout to stop sftp transfer in seconds, 0 is no limit`,
				},
				resource.Attribute{
					Name:        "tftp",
					Description: `set timeout to stop tftp transfer in seconds, 0 is no limit`,
				},
				resource.Attribute{
					Name:        "http",
					Description: `set timeout to stop http transfer in seconds, 0 is no limit`,
				},
				resource.Attribute{
					Name:        "https",
					Description: `set timeout to stop https transfer in seconds, 0 is no limit`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all': all; 'connattempt': Connect initiated; 'connects': Connect established; 'drops': Connect dropped; 'conndrops': Embryonic connect dropped; 'closed': Connect closed; 'segstimed': Segs to get RTT; 'rttupdated': Update RTT; 'delack': Delayed acks sent; 'timeoutdrop': Conn dropped in rxmt timeout; 'rexmttimeo': Retransmit timeout; 'persisttimeo': Persist timeout; 'keeptimeo': Keepalive timeout; 'keepprobe': Keepalive probe sent; 'keepdrops': Connect dropped in keepalive; 'sndtotal': Total packet sent; 'sndpack': Data packet sent; 'sndbyte': Data bytes sent; 'sndrexmitpack': Data packet retransmit; 'sndrexmitbyte': Data byte retransmit; 'sndrexmitbad': Unnecessary packet retransmit; 'sndacks': Ack packet sent; 'sndprobe': Window probe sent; 'sndurg': URG packet sent; 'sndwinup': Window update packet sent; 'sndctrl': SYN|FIN|RST packet sent; 'sndrst': RST packet sent; 'sndfin': FIN packet sent; 'sndsyn': SYN packet sent; 'rcvtotal': Total packet received; 'rcvpack': Packet received; 'rcvbyte': Bytes received; 'rcvbadoff': Packet received with bad offset; 'rcvmemdrop': Packet dropped for lack of memory; 'rcvduppack': Duplicate packet received; 'rcvdupbyte': Duplicate bytes received; 'rcvpartduppack': Packet with some duplicate data; 'rcvpartdupbyte': Dup. bytes in part-dup. packets; 'rcvoopack': Out-of-order packet received; 'rcvoobyte': Out-of-order bytes received; 'rcvpackafterwin': Packets with data after window; 'rcvbyteafterwin': Bytes rcvd after window; 'rcvwinprobe': Rcvd window probe packet; 'rcvdupack': Rcvd duplicate acks; 'rcvacktoomuch': Rcvd acks for unsent data; 'rcvackpack': Rcvd ack packets; 'rcvackbyte': Bytes acked by rcvd acks; 'rcvwinupd': Rcvd window update packets; 'pawsdrop': Segments dropped due to PAWS; 'predack': Hdr predict for acks; 'preddat': Hdr predict for data pkts; 'persistdrop': Timeout in persist state; 'badrst': Ignored RST; 'finwait2_drops': Drop FIN_WAIT_2 connection after time limit; 'sack_recovery_episode': SACK recovery episodes; 'sack_rexmits': SACK rexmit segments; 'sack_rexmit_bytes': SACK rexmit bytes; 'sack_rcv_blocks': SACK received; 'sack_send_blocks': SACK sent; 'sndcack': Challenge ACK sent; 'cacklim': Challenge ACK limited; 'reassmemdrop': Packet dropped during reassembly; 'reasstimeout': Reassembly Time Out; 'cc_idle': Congestion control window set do to idle; 'cc_reduce': Congestion control window reduced by event; 'rcvdsack': Rcvd DSACK packets; 'a2brcvwnd': ATCP to BTCP receive window; 'a2bsackpresent': ATCP to BTCP SACK options present; 'a2bdupack': ATCP to BTCP Dup/OO ACK; 'a2brxdata': ATCP to BTCP Rxmitted data; 'a2btcpoptions': ATCP to BTCP unsupported TCP options; 'a2boodata': ATCP to BTCP oo data received; 'a2bpartialack': ATCP to BTCP partial ack received; 'a2bfsmtransition': ATCP to BTCP state machine transition; 'a2btransitionnum': ATCP to BTCP total transitions; 'b2atransitionnum': ATCP to BTCP total transitions; 'bad_iochan': IO Channel Modified; 'atcpforward': Adaptive TCP forward; 'atcpsent': Adaptive TCP sent; 'atcprexmitsadrop': Adaptive TCP transmit SA drops; 'atcpsendbackack': Adaptive TCP sendback ACK; 'atcprexmit': Adaptive TCP retransmits; 'atcpbuffallocfail': Adaptive TCP buffer allocation fails; 'a2bappbuffering': Transition to full stack on when application buffers too much data; 'atcpsendfail': Adaptive TCP sent fails; 'earlyrexmit': Early Retransmission sent; 'mburstlim': Maxburst limited tx; 'a2bsndwnd': ATCP to BTCP send window;`,
				},
				resource.Attribute{
					Name:        "port-index",
					Description: `port index to be configured (Specify port index)`,
				},
				resource.Attribute{
					Name:        "mac-address",
					Description: `mac-address to be configured as mgmt port`,
				},
				resource.Attribute{
					Name:        "pci-address",
					Description: `pci-address to be configured as mgmt port`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Enable/Disable micro-burst traffic support`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `Disable IPMI on platform`,
				},
				resource.Attribute{
					Name:        "port-number",
					Description: `port number to be configured (Specify port number)`,
				},
				resource.Attribute{
					Name:        "core-index",
					Description: `core index to be deleted (Specify core index)`,
				},
				resource.Attribute{
					Name:        "max-cores",
					Description: `max number of IO cores (Specify number of cores)`,
				},
				resource.Attribute{
					Name:        "rxd-size",
					Description: `Set new rx-descriptor size`,
				},
				resource.Attribute{
					Name:        "txd-size",
					Description: `Set new tx-descriptor size`,
				},
				resource.Attribute{
					Name:        "rxq-size",
					Description: `Set number of new rx queues`,
				},
				resource.Attribute{
					Name:        "txq-size",
					Description: `Set number of new tx queues`,
				},
				resource.Attribute{
					Name:        "template-policy",
					Description: `Apply policy template to the whole system (Policy template name)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specify unique Partition id`,
				},
				resource.Attribute{
					Name:        "sessions",
					Description: `'all': Clear all sessions; 'sequence': Sequence number;`,
				},
				resource.Attribute{
					Name:        "clear-all-sequence",
					Description: `Sequence number (Specify the port physical port number)`,
				},
				resource.Attribute{
					Name:        "clear-sequence",
					Description: `Specify the port physical port number`,
				},
				resource.Attribute{
					Name:        "diseth",
					Description: `Specify the physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "dis-sequence",
					Description: `Sequence number (Specify the sequence number)`,
				},
				resource.Attribute{
					Name:        "enaeth",
					Description: `Specify the physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "ena-sequence",
					Description: `Sequence number (Specify the sequence number)`,
				},
				resource.Attribute{
					Name:        "monitor-relation",
					Description: `'monitor-and': Configures the monitors in current template to work with AND logic; 'monitor-or': Configures the monitors in current template to work with OR logic;`,
				},
				resource.Attribute{
					Name:        "linkup-ethernet1",
					Description: `Specify the port physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "link-up-sequence1",
					Description: `Sequence number (Specify the sequence number)`,
				},
				resource.Attribute{
					Name:        "linkup-ethernet2",
					Description: `Specify the port physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "link-up-sequence2",
					Description: `Sequence number (Specify the sequence number)`,
				},
				resource.Attribute{
					Name:        "linkup-ethernet3",
					Description: `Specify the port physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "link-up-sequence3",
					Description: `Sequence number (Specify the sequece number)`,
				},
				resource.Attribute{
					Name:        "linkdown-ethernet1",
					Description: `Specify the port physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "link-down-sequence1",
					Description: `Sequence number (Specify the sequence number)`,
				},
				resource.Attribute{
					Name:        "linkdown-ethernet2",
					Description: `Specify the port physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "link-down-sequence2",
					Description: `Sequence number (Specify the seqeuence number)`,
				},
				resource.Attribute{
					Name:        "linkdown-ethernet3",
					Description: `Specify the port physical port number (Ethernet interface number)`,
				},
				resource.Attribute{
					Name:        "link-down-sequence3",
					Description: `Sequence number (Specify the sequence number)`,
				},
				resource.Attribute{
					Name:        "user-tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "ssl-context-memory",
					Description: `Total SSL context memory needed in units of MB. Will be rounded to closest multiple of 2MB`,
				},
				resource.Attribute{
					Name:        "ssl-dma-memory",
					Description: `Total SSL DMA memory needed in units of MB. Will be rounded to closest multiple of 2MB`,
				},
				resource.Attribute{
					Name:        "nat-pool-addr-count",
					Description: `Total configurable NAT Pool addresses in the System`,
				},
				resource.Attribute{
					Name:        "l4-session-count",
					Description: `Total Sessions in the System`,
				},
				resource.Attribute{
					Name:        "auth-portal-html-file-size",
					Description: `Specify maximum html file size for each html page in auth portal (in KB)`,
				},
				resource.Attribute{
					Name:        "auth-portal-image-file-size",
					Description: `Specify maximum image file size for default portal (in KB)`,
				},
				resource.Attribute{
					Name:        "max-aflex-file-size",
					Description: `Set maximum aFleX file size (Maximum file size in KBytes, default is 32K)`,
				},
				resource.Attribute{
					Name:        "aflex-table-entry-count",
					Description: `Total aFleX table entry in the system (Total aFlex entry in the system)`,
				},
				resource.Attribute{
					Name:        "class-list-ipv6-addr-count",
					Description: `Total IPv6 addresses for class-list`,
				},
				resource.Attribute{
					Name:        "class-list-ac-entry-count",
					Description: `Total entries for AC class-list`,
				},
				resource.Attribute{
					Name:        "max-aflex-authz-collection-number",
					Description: `Specify the maximum number of collections supported by aFleX authorization`,
				},
				resource.Attribute{
					Name:        "radius-table-size",
					Description: `Total configurable CGNV6 RADIUS Table entries`,
				},
				resource.Attribute{
					Name:        "authz-policy-number",
					Description: `Specify the maximum number of authorization policies`,
				},
				resource.Attribute{
					Name:        "ipsec-sa-number",
					Description: `Specify the maximum number of IPsec SA`,
				},
				resource.Attribute{
					Name:        "ram-cache-memory-limit",
					Description: `Specify the maximum memory used by ram cache`,
				},
				resource.Attribute{
					Name:        "monitored-entity-count",
					Description: `Total number of monitored entities for visibility`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specify name of Geolocation list`,
				},
				resource.Attribute{
					Name:        "gslb-device-max",
					Description: `Enter the number of gslb-device allowed (gslb-device count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "gslb-device-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "gslb-geo-location-max",
					Description: `Enter the number of gslb-geo-location allowed (gslb-geo-location count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "gslb-geo-location-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "gslb-ip-list-max",
					Description: `Enter the number of gslb-ip-list allowed (gslb-ip-list count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "gslb-ip-list-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "gslb-policy-max",
					Description: `Enter the number of gslb-policy allowed (gslb-policy count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "gslb-policy-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "gslb-service-max",
					Description: `Enter the number of gslb-service allowed (gslb-service count (default is max-value)`,
				},
				resource.Attribute{
					Name:        "gslb-service-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "gslb-service-ip-max",
					Description: `Enter the number of gslb-service-ip allowed (gslb-service-ip count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "gslb-service-ip-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "gslb-service-port-max",
					Description: `Enter the number of gslb-service-port allowed ( gslb-service-port count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "gslb-service-port-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "gslb-site-max",
					Description: `Enter the number of gslb-site allowed (gslb-site count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "gslb-site-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "gslb-svc-group-max",
					Description: `Enter the number of gslb-svc-group allowed (gslb-svc-group count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "gslb-svc-group-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "gslb-template-max",
					Description: `Enter the number of gslb-template allowed (gslb-template count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "gslb-template-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "gslb-zone-max",
					Description: `Enter the number of gslb-zone allowed (gslb-zone count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "gslb-zone-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "health-monitor-max",
					Description: `Enter the number of health monitors allowed (health-monitor count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "health-monitor-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "real-port-max",
					Description: `Enter the number of real-ports allowed (real-port count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "real-port-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "real-server-max",
					Description: `Enter the number of real-servers allowed (real-server count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "real-server-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "service-group-max",
					Description: `Enter the number of service groups allowed (service-group count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "service-group-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "virtual-server-max",
					Description: `Enter the number of virtual-servers allowed (virtual-server count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "virtual-server-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "virtual-port-max",
					Description: `Enter the number of virtual-port allowed (virtual-port count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "virtual-port-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "cache-template-max",
					Description: `Enter the number of cache-template allowed (cache-template count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "cache-template-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "client-ssl-template-max",
					Description: `Enter the number of client-ssl-template allowed (client-ssl-template count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "client-ssl-template-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "conn-reuse-template-max",
					Description: `Enter the number of conn-reuse-template allowed (conn-reuse-template count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "conn-reuse-template-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "fast-tcp-template-max",
					Description: `Enter the number of fast-tcp-template allowed (fast-tcp-template count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "fast-tcp-template-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "fast-udp-template-max",
					Description: `Enter the number of fast-udp-template allowed (fast-udp-template count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "fast-udp-template-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "fix-template-max",
					Description: `Enter the number of fix-template allowed (fix-template count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "fix-template-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "http-template-max",
					Description: `Enter the number of http-template allowed (http-template count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "http-template-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "link-cost-template-max",
					Description: `Enter the number of link-cost-template allowed (link-cost-template count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "link-cost-template-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "persist-cookie-template-max",
					Description: `Enter the number of persist-cookie-template allowed (persist-cookie-template count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "persist-cookie-template-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "persist-srcip-template-max",
					Description: `Enter the number of persist-srcip-template allowed (persist-source-ip-template count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "persist-srcip-template-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "server-ssl-template-max",
					Description: `Enter the number of server-ssl-template allowed (server-ssl-template count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "server-ssl-template-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "proxy-template-max",
					Description: `Enter the number of proxy-template allowed (server-ssl-template count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "proxy-template-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "stream-template-max",
					Description: `Enter the number of stream-template allowed (server-ssl-template count (default is max-value))`,
				},
				resource.Attribute{
					Name:        "stream-template-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Enter the threshold as a percentage (Threshold in percentage(default is 100%))`,
				},
				resource.Attribute{
					Name:        "static-ipv4-route-max",
					Description: `Enter the number of static ipv4 routes allowed (Static ipv4 routes (default is max-value))`,
				},
				resource.Attribute{
					Name:        "static-ipv4-route-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "static-ipv6-route-max",
					Description: `Enter the number of static ipv6 routes allowed (Static ipv6 routes (default is max-value))`,
				},
				resource.Attribute{
					Name:        "static-ipv6-route-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "ipv4-acl-line-max",
					Description: `Enter the number of ACL lines allowed (IPV4 ACL lines (default is max-value))`,
				},
				resource.Attribute{
					Name:        "ipv4-acl-line-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "ipv6-acl-line-max",
					Description: `Enter the number of ACL lines allowed (IPV6 ACL lines (default is max-value))`,
				},
				resource.Attribute{
					Name:        "ipv6-acl-line-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "static-arp-max",
					Description: `Enter the number of static arp entries allowed (Static arp (default is max-value))`,
				},
				resource.Attribute{
					Name:        "static-arp-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "static-neighbor-max",
					Description: `Enter the number of static neighbor entries allowed (Static neighbors (default is max-value))`,
				},
				resource.Attribute{
					Name:        "static-neighbor-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "static-mac-max",
					Description: `Enter the number of static MAC entries allowed (Static MACs (default is max-value))`,
				},
				resource.Attribute{
					Name:        "static-mac-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "object-group-max",
					Description: `Enter the number of object groups allowed (Object group (default is max-value))`,
				},
				resource.Attribute{
					Name:        "object-group-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "object-group-clause-max",
					Description: `Enter the number of object group clauses allowed (Object group clauses (default is max-value))`,
				},
				resource.Attribute{
					Name:        "object-group-clause-min-guarantee",
					Description: `Minimum guaranteed value ( Minimum guaranteed value)`,
				},
				resource.Attribute{
					Name:        "bw-limit-max",
					Description: `Enter the bandwidth limit in mbps (Bandwidth limit in Mbit/s (no limits applied by default))`,
				},
				resource.Attribute{
					Name:        "bw-limit-watermark-disable",
					Description: `Disable watermark (90% drop, keep existing sessions, drop new sessions)`,
				},
				resource.Attribute{
					Name:        "concurrent-session-limit-max",
					Description: `Enter the Concurrent Session limit (cps) (Concurrent-Session cps limit (no limits applied by default))`,
				},
				resource.Attribute{
					Name:        "l4-session-limit-max",
					Description: `Enter the l4 session limit in % (0.01% to 99.99%) (Enter a number from 0.01 to 99.99 (up to 2 digits precision))`,
				},
				resource.Attribute{
					Name:        "l4-session-limit-min-guarantee",
					Description: `minimum guaranteed value in % (up to 2 digits precision) (Enter a number from 0 to 99.99)`,
				},
				resource.Attribute{
					Name:        "l4cps-limit-max",
					Description: `Enter the L4 cps limit (L4 cps limit (no limits applied by default))`,
				},
				resource.Attribute{
					Name:        "l7cps-limit-max",
					Description: `L7cps-limit (L7 cps limit (no limits applied by default))`,
				},
				resource.Attribute{
					Name:        "natcps-limit-max",
					Description: `Enter the Nat cps limit (NAT cps limit (no limits applied by default))`,
				},
				resource.Attribute{
					Name:        "fwcps-limit-max",
					Description: `Enter the Firewall cps limit (Firewall cps limit (no limits applied by default))`,
				},
				resource.Attribute{
					Name:        "ssl-throughput-limit-max",
					Description: `Enter the ssl throughput limit in mbps (SSL Througput limit in Mbit/s (no limits applied by default))`,
				},
				resource.Attribute{
					Name:        "ssl-throughput-limit-watermark-disable",
					Description: `Disable watermark (90% drop, keep existing sessions, drop new sessions)`,
				},
				resource.Attribute{
					Name:        "sslcps-limit-max",
					Description: `Enter the SSL cps limit (SSL cps limit (no limits applied by default))`,
				},
				resource.Attribute{
					Name:        "use-l3",
					Description: `Layer-3 Header based load balancing`,
				},
				resource.Attribute{
					Name:        "use-l4",
					Description: `Layer-3/4 Header based load balancing`,
				},
				resource.Attribute{
					Name:        "link-detection-interval",
					Description: `Link detection interval in msecs`,
				},
				resource.Attribute{
					Name:        "packet-round-robin",
					Description: `Enable packet round robin for IPsec packets`,
				},
				resource.Attribute{
					Name:        "crypto-core",
					Description: `Crypto cores assigned for IPsec processing`,
				},
				resource.Attribute{
					Name:        "crypto-mem",
					Description: `Crypto memory percentage assigned for IPsec processing (rounded to increments of 10)`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `'ipv4-only': Enable IPv4 HW forward entries only; 'ipv6-only': Enable IPv6 HW forward entries only; 'ipv4-ipv6': Enable Both IPv4/IPv6 HW forward entries (shared);`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Minimum packets-per-second threshold (per CPU) before redistribution will take effect (Minimum packets-per-second threshold (per CPU) before redistribution will take effect (default: 100000))`,
				},
				resource.Attribute{
					Name:        "low",
					Description: `CPU usage threshold (percentage) that will restore the normal packet distribution (default: 60)`,
				},
				resource.Attribute{
					Name:        "high",
					Description: `CPU usage threshold (percentage) that will trigger the redistribution (default: 75)`,
				},
				resource.Attribute{
					Name:        "bcast",
					Description: `broadcast packets (per second limit)`,
				},
				resource.Attribute{
					Name:        "ipmcast",
					Description: `IP multicast packets (per second limit)`,
				},
				resource.Attribute{
					Name:        "mcast",
					Description: `multicast packets (per second limit)`,
				},
				resource.Attribute{
					Name:        "unknown-ucast",
					Description: `unknown unicast packets (per second limit)`,
				},
				resource.Attribute{
					Name:        "ve-mac-scheme-val",
					Description: `'hash-based': Hash-based using the VE number; 'round-robin': Round Robin scheme; 'system-mac': Use system MAC address;`,
				},
				resource.Attribute{
					Name:        "nscan-limit",
					Description: `smp session scan limit (number of smp sessions per scan)`,
				},
				resource.Attribute{
					Name:        "scan-freq",
					Description: `smp session scan frequency (scan per second)`,
				},
				resource.Attribute{
					Name:        "reset",
					Description: `Reset IPMI Controller`,
				},
				resource.Attribute{
					Name:        "ipv4-address",
					Description: `IP address`,
				},
				resource.Attribute{
					Name:        "ipv4-netmask",
					Description: `IP subnet mask`,
				},
				resource.Attribute{
					Name:        "default-gateway",
					Description: `Default gateway address`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `IP addr obtained by BMC running DHCP`,
				},
				resource.Attribute{
					Name:        "static",
					Description: `Manually configured static IP address`,
				},
				resource.Attribute{
					Name:        "add",
					Description: `Add a new IPMI user (IPMI User Name)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password`,
				},
				resource.Attribute{
					Name:        "administrator",
					Description: `Full control`,
				},
				resource.Attribute{
					Name:        "callback",
					Description: `Lowest privilege level`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Most BMC commands are allowed`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Only 'benign' commands are allowed`,
				},
				resource.Attribute{
					Name:        "privilege",
					Description: `Change an existing IPMI user privilege (IPMI User Name)`,
				},
				resource.Attribute{
					Name:        "setname",
					Description: `Change User Name (Current IPMI User Name)`,
				},
				resource.Attribute{
					Name:        "newname",
					Description: `New IPMI User Name`,
				},
				resource.Attribute{
					Name:        "setpass",
					Description: `Change Password (IPMI User Name)`,
				},
				resource.Attribute{
					Name:        "newpass",
					Description: `New Password`,
				},
				resource.Attribute{
					Name:        "cmd",
					Description: `Command to execute in double quotes`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Set HW hash mode, default is 6 (1:dst-mac 2:src-mac 3:src-dst-mac 4:src-ip 5:dst-ip 6:rtag6 7:rtag7)`,
				},
				resource.Attribute{
					Name:        "source_name",
					Description: `bind source name`,
				},
				resource.Attribute{
					Name:        "dest_name",
					Description: `bind dest name`,
				},
				resource.Attribute{
					Name:        "log-session-on-established",
					Description: `Send TCP session creation log on completion of 3-way handshake`,
				},
				resource.Attribute{
					Name:        "msl-time",
					Description: `Configure maximum session life, default is 2 seconds (1-40 seconds, default is 2 seconds)`,
				},
				resource.Attribute{
					Name:        "application-mempool",
					Description: `Enable application memory pool`,
				},
				resource.Attribute{
					Name:        "application-flow",
					Description: `Number of flows`,
				},
				resource.Attribute{
					Name:        "basic-dpi-enable",
					Description: `Enable basic dpi`,
				},
				resource.Attribute{
					Name:        "complexity",
					Description: `'Strict': Strict: Min length:8, Min Lower Case:2, Min Upper Case:2, Min Numbers:2, Min Special Character:1; 'Medium': Medium: Min length:6, Min Lower Case:2, Min Upper Case:2, Min Numbers:1, Min Special Character:1; 'Simple': Simple: Min length:4, Min Lower Case:1, Min Upper Case:1, Min Numbers:1, Min Special Character:0;`,
				},
				resource.Attribute{
					Name:        "aging",
					Description: `'Strict': Strict: Max Age-60 Days; 'Medium': Medium: Max Age- 90 Days; 'Simple': Simple: Max Age-120 Days;`,
				},
				resource.Attribute{
					Name:        "history",
					Description: `'Strict': Strict: Does not allow upto 5 old passwords; 'Medium': Medium: Does not allow upto 4 old passwords; 'Simple': Simple: Does not allow upto 3 old passwords;`,
				},
				resource.Attribute{
					Name:        "min-pswd-len",
					Description: `Configure custom password length`,
				},
				resource.Attribute{
					Name:        "listen-port",
					Description: `Configure the listen port of RADIUS server (default 1813) (Port number)`,
				},
				resource.Attribute{
					Name:        "ip-list-name",
					Description: `IP-list name`,
				},
				resource.Attribute{
					Name:        "ip-list-secret",
					Description: `Configure shared secret`,
				},
				resource.Attribute{
					Name:        "ip-list-secret-string",
					Description: `The RADIUS secret`,
				},
				resource.Attribute{
					Name:        "ip-list-encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED secret string)`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `Configure shared secret`,
				},
				resource.Attribute{
					Name:        "secret-string",
					Description: `The RADIUS secret`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED secret string)`,
				},
				resource.Attribute{
					Name:        "vrid",
					Description: `Join a VRRP-A failover group`,
				},
				resource.Attribute{
					Name:        "attribute-value",
					Description: `'inside-ipv6-prefix': Framed IPv6 Prefix; 'inside-ip': Inside IP address; 'inside-ipv6': Inside IPv6 address; 'imei': International Mobile Equipment Identity (IMEI); 'imsi': International Mobile Subscriber Identity (IMSI); 'msisdn': Mobile Subscriber Integrated Services Digital Network-Number (MSISDN); 'custom1': Customized attribute 1; 'custom2': Customized attribute 2; 'custom3': Customized attribute 3;`,
				},
				resource.Attribute{
					Name:        "prefix-length",
					Description: `'32': Prefix length 32; '48': Prefix length 48; '64': Prefix length 64; '80': Prefix length 80; '96': Prefix length 96; '112': Prefix length 112;`,
				},
				resource.Attribute{
					Name:        "prefix-vendor",
					Description: `RADIUS vendor attribute information (RADIUS vendor ID)`,
				},
				resource.Attribute{
					Name:        "prefix-number",
					Description: `RADIUS attribute number`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `'hexadecimal': Type of attribute value is hexadecimal;`,
				},
				resource.Attribute{
					Name:        "custom-vendor",
					Description: `RADIUS vendor attribute information (RADIUS vendor ID)`,
				},
				resource.Attribute{
					Name:        "custom-number",
					Description: `RADIUS attribute number`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `RADIUS vendor attribute information (RADIUS vendor ID)`,
				},
				resource.Attribute{
					Name:        "number",
					Description: `RADIUS attribute number`,
				},
				resource.Attribute{
					Name:        "disable-reply",
					Description: `Toggle option for RADIUS reply packet(Default: Accounting response will be sent)`,
				},
				resource.Attribute{
					Name:        "accounting-start",
					Description: `'ignore': Ignore; 'append-entry': Append the AVPs to existing entry (default); 'replace-entry': Replace the AVPs of existing entry;`,
				},
				resource.Attribute{
					Name:        "accounting-stop",
					Description: `'ignore': Ignore; 'delete-entry': Delete the entry (default); 'delete-entry-and-sessions': Delete the entry and data sessions associated(CGN only);`,
				},
				resource.Attribute{
					Name:        "accounting-interim-update",
					Description: `'ignore': Ignore (default); 'append-entry': Append the AVPs to existing entry; 'replace-entry': Replace the AVPs of existing entry;`,
				},
				resource.Attribute{
					Name:        "accounting-on",
					Description: `'ignore': Ignore (default); 'delete-entries-using-attribute': Delete entries matching attribute in RADIUS Table;`,
				},
				resource.Attribute{
					Name:        "attribute-name",
					Description: `'msisdn': Clear using MSISDN; 'imei': Clear using IMEI; 'imsi': Clear using IMSI; 'custom1': Clear using CUSTOM1 attribute configured; 'custom2': Clear using CUSTOM2 attribute configured; 'custom3': Clear using CUSTOM3 attribute configured;`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Enable sharing with other partitions`,
				},
				resource.Attribute{
					Name:        "include-geoloc-name-val",
					Description: `Geolocation name to add`,
				},
				resource.Attribute{
					Name:        "exclude-geoloc-name-val",
					Description: `Geolocation name to exclude`,
				},
				resource.Attribute{
					Name:        "geo-location-iana",
					Description: `Load built-in IANA Database`,
				},
				resource.Attribute{
					Name:        "geo-location-geolite2-city",
					Description: `Load built-in Maxmind GeoLite2-City database. Database available from http://www.maxmind.com`,
				},
				resource.Attribute{
					Name:        "geolite2-city-include-ipv6",
					Description: `Include IPv6 address`,
				},
				resource.Attribute{
					Name:        "geo-location-geolite2-country",
					Description: `Load built-in Maxmind GeoLite2-Country database. Database available from http://www.maxmind.com`,
				},
				resource.Attribute{
					Name:        "geolite2-country-include-ipv6",
					Description: `Include IPv6 address`,
				},
				resource.Attribute{
					Name:        "geo-location-load-filename",
					Description: `Specify file to be loaded`,
				},
				resource.Attribute{
					Name:        "template-name",
					Description: `CSV template to load this file`,
				},
				resource.Attribute{
					Name:        "geo-locn-obj-name",
					Description: `Specify geo-location name, section range is (1-15)`,
				},
				resource.Attribute{
					Name:        "first-ip-address",
					Description: `Specify IP information (Specify IP address)`,
				},
				resource.Attribute{
					Name:        "geol-ipv4-mask",
					Description: `Specify IPv4 mask`,
				},
				resource.Attribute{
					Name:        "ip-addr2",
					Description: `Specify IP address range`,
				},
				resource.Attribute{
					Name:        "first-ipv6-address",
					Description: `Specify IPv6 address`,
				},
				resource.Attribute{
					Name:        "geol-ipv6-mask",
					Description: `Specify IPv6 mask`,
				},
				resource.Attribute{
					Name:        "ipv6-addr2",
					Description: `Specify IPv6 address range`,
				},
				resource.Attribute{
					Name:        "class-list",
					Description: `Bind class-list (class-list name)`,
				},
				resource.Attribute{
					Name:        "ip-threat-action-tmpl",
					Description: `Bind ip-threat-action Template (ip-threat-action Template number)`,
				},
				resource.Attribute{
					Name:        "monitor-disable",
					Description: `Disable FPGA Core CRC error monitoring and act on it`,
				},
				resource.Attribute{
					Name:        "reboot-enable",
					Description: `Enable system reboot if system encounters FPGA Core CRC error`,
				},
				resource.Attribute{
					Name:        "log-for-reset-unknown-conn",
					Description: `Log when rate exceed`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_system_ve_mac_scheme",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder system ve mac scheme resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ve-mac-scheme",
					Description: `VE MAC allocation scheme`,
				},
				resource.Attribute{
					Name:        "ve-mac-scheme-val",
					Description: `'hash-based': Hash-based using the VE number; 'round-robin': Round Robin scheme; 'system-mac': Use system MAC address;`,
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
			Type:             "vthunder_thunder_virtual_server",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder virtual server resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "acl_id",
					Description: `ACL id VPORT`,
				},
				resource.Attribute{
					Name:        "acl_id_shared",
					Description: `acl id`,
				},
				resource.Attribute{
					Name:        "acl_name",
					Description: `Apply an access list name (Named Access List)`,
				},
				resource.Attribute{
					Name:        "acl_name_shared",
					Description: `Access List name (IPv4 Access List Name)`,
				},
				resource.Attribute{
					Name:        "arp_disable",
					Description: `Disable Respond to Virtual Server ARP request`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Create a description for VIP`,
				},
				resource.Attribute{
					Name:        "disable_vip_adv",
					Description: `Disable virtual server GARP`,
				},
				resource.Attribute{
					Name:        "enable_disable_action",
					Description: `‘enable’: Enable Virtual Server (default); ‘disable’: Disable Virtual Server; ‘disable-when-all-ports-down’: Disable Virtual Server when all member ports are down; ‘disable-when-any-port-down’: Disable Virtual Server when any member port is down;`,
				},
				resource.Attribute{
					Name:        "ethernet",
					Description: `Ethernet interface`,
				},
				resource.Attribute{
					Name:        "extended_stats",
					Description: `Enable extended statistics on virtual port`,
				},
				resource.Attribute{
					Name:        "ha_dynamic",
					Description: `Dynamic failover based on vip status`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP Address`,
				},
				resource.Attribute{
					Name:        "ipv6_acl",
					Description: `ipv6 acl name`,
				},
				resource.Attribute{
					Name:        "ipv6_acl_shared",
					Description: `ipv6 acl name`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `IPV6 address`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `SLB Virtual Service Name`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `IP subnet mask`,
				},
				resource.Attribute{
					Name:        "redistribute_route_map",
					Description: `Route map reference (Name of route-map)`,
				},
				resource.Attribute{
					Name:        "redistribution_flagged",
					Description: `Flag VIP for special redistribution handling`,
				},
				resource.Attribute{
					Name:        "shared_partition_policy_template",
					Description: `Reference a policy template from shared partition`,
				},
				resource.Attribute{
					Name:        "stats_data_action",
					Description: `‘stats-data-enable’: Enable statistical data collection for virtual port; ‘stats-data-disable’: Disable statistical data collection for virtual port;`,
				},
				resource.Attribute{
					Name:        "template_logging",
					Description: `NAT Logging template (NAT Logging template name)`,
				},
				resource.Attribute{
					Name:        "template_policy",
					Description: `Policy Template (Policy template name)`,
				},
				resource.Attribute{
					Name:        "template_policy_shared",
					Description: `Policy Template Name`,
				},
				resource.Attribute{
					Name:        "template_scaleout",
					Description: `Scaleout template (Scaleout template name)`,
				},
				resource.Attribute{
					Name:        "template_virtual_server",
					Description: `Virtual server template (Virtual server template name)`,
				},
				resource.Attribute{
					Name:        "use_if_ip",
					Description: `Use Interface IP`,
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
					Name:        "vrid",
					Description: `Join a vrrp group (Specify ha VRRP-A vrid)`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `‘enable’: Enable; ‘disable’: Disable;`,
				},
				resource.Attribute{
					Name:        "alt_protocol1",
					Description: `‘http’: HTTP Port;`,
				},
				resource.Attribute{
					Name:        "alt_protocol2",
					Description: `‘tcp’: TCP LB service;`,
				},
				resource.Attribute{
					Name:        "alternate_port",
					Description: `Alternate Virtual Port`,
				},
				resource.Attribute{
					Name:        "alternate_port_number",
					Description: `Virtual Port`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `Configure auto NAT for the vport`,
				},
				resource.Attribute{
					Name:        "clientip_sticky_nat",
					Description: `Prefer to use same source NAT address for a client`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `Connection Limit`,
				},
				resource.Attribute{
					Name:        "def_selection_if_pref_failed",
					Description: `‘def-selection-if-pref-failed’: Use default server selection method if prefer method failed; ‘def-selection-if-pref-failed-disable’: Stop using default server selection method if prefer method failed;`,
				},
				resource.Attribute{
					Name:        "enable_playerid_check",
					Description: `Enable playerid checks on UDP packets once the AX is in active mode`,
				},
				resource.Attribute{
					Name:        "eth_fwd",
					Description: `Ethernet interface number`,
				},
				resource.Attribute{
					Name:        "eth_rev",
					Description: `Ethernet interface number`,
				},
				resource.Attribute{
					Name:        "expand",
					Description: `expand syn-cookie with timestamp and wscale`,
				},
				resource.Attribute{
					Name:        "force_routing_mode",
					Description: `Force routing mode`,
				},
				resource.Attribute{
					Name:        "gslb_enable",
					Description: `Enable Global Server Load Balancing`,
				},
				resource.Attribute{
					Name:        "ha_conn_mirror",
					Description: `Enable for HA Conn sync`,
				},
				resource.Attribute{
					Name:        "ip_map_list",
					Description: `Enter name of IP Map List to be bound (IP Map List Name)`,
				},
				resource.Attribute{
					Name:        "ipinip",
					Description: `Enable IP in IP`,
				},
				resource.Attribute{
					Name:        "l7_hardware_assist",
					Description: `FPGA assist L7 packet parsing`,
				},
				resource.Attribute{
					Name:        "message_switching",
					Description: `Message switching`,
				},
				resource.Attribute{
					Name:        "no_auto_up_on_aflex",
					Description: `Don’t automatically mark vport up when aFleX is bound`,
				},
				resource.Attribute{
					Name:        "no_dest_nat",
					Description: `Disable destination NAT, this option only supports in wildcard VIP or when a connection is operated in SSLi + EP mode`,
				},
				resource.Attribute{
					Name:        "no_logging",
					Description: `Do not log connection over limit event`,
				},
				resource.Attribute{
					Name:        "on_syn",
					Description: `Enable for HA Conn sync for l4 tcp sessions on SYN`,
				},
				resource.Attribute{
					Name:        "optimization_level",
					Description: `‘0’: No optimization; ‘1’: Optimization level 1 (Experimental);`,
				},
				resource.Attribute{
					Name:        "persist_type",
					Description: `‘src-dst-ip-swap-persist’: Create persist session after source IP and destination IP swap; ‘use-src-ip-for-dst-persist’: Use the source IP to create a destination persist session; ‘use-dst-ip-for-src-persist’: Use the destination IP to create source IP persist session;`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `Specify NAT pool or pool group`,
				},
				resource.Attribute{
					Name:        "pool_shared",
					Description: `Specify NAT pool or pool group`,
				},
				resource.Attribute{
					Name:        "port_number",
					Description: `Port`,
				},
				resource.Attribute{
					Name:        "port_translation",
					Description: `Enable port translation under no-dest-nat`,
				},
				resource.Attribute{
					Name:        "precedence",
					Description: `Set auto NAT pool as higher precedence for source NAT`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `‘tcp’: TCP LB service; ‘udp’: UDP Port; ‘others’: for no tcp/udp protocol, do IP load balancing; ‘diameter’: diameter port; ‘dns-tcp’: DNS service over TCP; ‘dns-udp’: DNS service over UDP; ‘fast-http’: Fast HTTP Port; ‘fix’: FIX Port; ‘ftp’: File Transfer Protocol Port; ‘ftp-proxy’: ftp proxy port; ‘http’: HTTP Port; ‘https’: HTTPS port; ‘imap’: imap proxy port; ‘mlb’: Message based load balancing; ‘mms’: Microsoft Multimedia Service Port; ‘mysql’: mssql port; ‘mssql’: mssql; ‘pop3’: pop3 proxy port; ‘radius’: RADIUS Port; ‘rtsp’: Real Time Streaming Protocol Port; ‘sip’: Session initiation protocol over UDP; ‘sip-tcp’: Session initiation protocol over TCP; ‘sips’: Session initiation protocol over TLS; ‘smpp-tcp’: SMPP service over TCP; ‘spdy’: spdy port; ‘spdys’: spdys port; ‘smtp’: SMTP Port; ‘ssl-proxy’: Generic SSL proxy; ‘ssli’: SSL insight; ‘ssh’: SSH Port; ‘tcp-proxy’: Generic TCP proxy; ‘tftp’: TFTP Port;`,
				},
				resource.Attribute{
					Name:        "range",
					Description: `Virtual Port range (Virtual Port range value)`,
				},
				resource.Attribute{
					Name:        "rate",
					Description: `Specify the log message rate`,
				},
				resource.Attribute{
					Name:        "redirect_to_https",
					Description: `Redirect HTTP to HTTPS`,
				},
				resource.Attribute{
					Name:        "req_fail",
					Description: `Use alternate virtual port when L7 request fail`,
				},
				resource.Attribute{
					Name:        "reset",
					Description: `Send client reset when connection number over limit`,
				},
				resource.Attribute{
					Name:        "reset_on_server_selection_fail",
					Description: `Send client reset when server selection fails`,
				},
				resource.Attribute{
					Name:        "rtp_sip_call_id_match",
					Description: `rtp traffic try to match the real server of sip smp call-id session`,
				},
				resource.Attribute{
					Name:        "scaleout_bucket_count",
					Description: `Number of traffic buckets`,
				},
				resource.Attribute{
					Name:        "scaleout_device_group",
					Description: `Device group id`,
				},
				resource.Attribute{
					Name:        "secs",
					Description: `Specify the interval in seconds`,
				},
				resource.Attribute{
					Name:        "serv_sel_fail",
					Description: `Use alternate virtual port when server selection failure`,
				},
				resource.Attribute{
					Name:        "service_group",
					Description: `Bind a Service Group to this Virtual Server (Service Group Name)`,
				},
				resource.Attribute{
					Name:        "shared_partition_cache_template",
					Description: `Reference a Cache template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_client_ssl_template",
					Description: `Reference a Client SSL template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_connection_reuse_template",
					Description: `Reference a connection reuse template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_diameter_template",
					Description: `Reference a Diameter template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_dns_template",
					Description: `Reference a dns template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_dynamic_service_template",
					Description: `Reference a dynamic service template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_external_service_template",
					Description: `Reference a external service template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_http_policy_template",
					Description: `Reference a http policy template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_http_template",
					Description: `Reference a HTTP template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_persist_cookie_template",
					Description: `Reference a persist cookie template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_persist_destination_ip_template",
					Description: `Reference a persist destination ip template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_persist_source_ip_template",
					Description: `Reference a persist source ip template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_persist_ssl_sid_template",
					Description: `Reference a persist SSL SID template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_pool",
					Description: `Specify NAT pool or pool group from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_server_ssl_template",
					Description: `Reference a SSL Server template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_tcp",
					Description: `Reference a tcp template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_tcp_proxy_template",
					Description: `Reference a TCP Proxy template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_udp",
					Description: `Reference a UDP template from shared partition`,
				},
				resource.Attribute{
					Name:        "shared_partition_virtual_port_template",
					Description: `Reference a Virtual Port template from shared partition`,
				},
				resource.Attribute{
					Name:        "skip_rev_hash",
					Description: `Skip rev tuple hash insertion`,
				},
				resource.Attribute{
					Name:        "snat_on_vip",
					Description: `Enable source NAT traffic against VIP`,
				},
				resource.Attribute{
					Name:        "support_http2",
					Description: `Support HTTP2`,
				},
				resource.Attribute{
					Name:        "syn_cookie",
					Description: `Enable syn-cookie`,
				},
				resource.Attribute{
					Name:        "template_cache",
					Description: `RAM caching template (Cache Template Name)`,
				},
				resource.Attribute{
					Name:        "template_cache_shared",
					Description: `Cache Template Name`,
				},
				resource.Attribute{
					Name:        "template_client_ssh",
					Description: `Client SSH Template (Client SSH Config Name)`,
				},
				resource.Attribute{
					Name:        "template_client_ssl",
					Description: `Client SSL Template (Client SSL Config Name)`,
				},
				resource.Attribute{
					Name:        "template_client_ssl_shared",
					Description: `Client SSL Template Name`,
				},
				resource.Attribute{
					Name:        "template_connection_reuse",
					Description: `Connection Reuse Template (Connection Reuse Template Name)`,
				},
				resource.Attribute{
					Name:        "template_connection_reuse_shared",
					Description: `Connection Reuse Template Name`,
				},
				resource.Attribute{
					Name:        "template_dblb",
					Description: `DBLB Template (DBLB template name)`,
				},
				resource.Attribute{
					Name:        "template_diameter",
					Description: `Diameter Template (diameter template name)`,
				},
				resource.Attribute{
					Name:        "template_diameter_shared",
					Description: `Diameter Template Name`,
				},
				resource.Attribute{
					Name:        "template_dns",
					Description: `DNS template (DNS template name)`,
				},
				resource.Attribute{
					Name:        "template_dns_shared",
					Description: `DNS Template Name`,
				},
				resource.Attribute{
					Name:        "template_dynamic_service",
					Description: `Dynamic Service Template (dynamic-service template name)`,
				},
				resource.Attribute{
					Name:        "template_dynamic_service_shared",
					Description: `Dynamic Service Template Name`,
				},
				resource.Attribute{
					Name:        "template_external_service",
					Description: `External service template (external-service template name)`,
				},
				resource.Attribute{
					Name:        "template_external_service_shared",
					Description: `External Service Template Name`,
				},
				resource.Attribute{
					Name:        "template_file_inspection",
					Description: `File Inspection service template (file-inspection template name)`,
				},
				resource.Attribute{
					Name:        "template_fix",
					Description: `FIX template (FIX Template Name)`,
				},
				resource.Attribute{
					Name:        "template_ftp",
					Description: `FTP port template (Ftp template name)`,
				},
				resource.Attribute{
					Name:        "template_http",
					Description: `HTTP Template (HTTP Config Name)`,
				},
				resource.Attribute{
					Name:        "template_http_policy",
					Description: `http-policy template (http-policy template name)`,
				},
				resource.Attribute{
					Name:        "template_http_policy_shared",
					Description: `Http Policy Template Name`,
				},
				resource.Attribute{
					Name:        "template_http_shared",
					Description: `HTTP Template Name`,
				},
				resource.Attribute{
					Name:        "template_imap_pop3",
					Description: `IMAP/POP3 Template (IMAP/POP3 Config Name)`,
				},
				resource.Attribute{
					Name:        "template_persist_cookie",
					Description: `Cookie persistence (Cookie persistence template name)`,
				},
				resource.Attribute{
					Name:        "template_persist_cookie_shared",
					Description: `Cookie Persistence Template Name`,
				},
				resource.Attribute{
					Name:        "template_persist_destination_ip",
					Description: `Destination IP persistence (Destination IP persistence template name)`,
				},
				resource.Attribute{
					Name:        "template_persist_destination_ip_shared",
					Description: `Destination IP Persistence Template Name`,
				},
				resource.Attribute{
					Name:        "template_persist_source_ip",
					Description: `Source IP persistence (Source IP persistence template name)`,
				},
				resource.Attribute{
					Name:        "template_persist_source_ip_shared",
					Description: `Source IP Persistence Template Name`,
				},
				resource.Attribute{
					Name:        "template_persist_ssl_sid",
					Description: `SSL session ID persistence (Source IP Persistence Config name)`,
				},
				resource.Attribute{
					Name:        "template_persist_ssl_sid_shared",
					Description: `SSL SID Persistence Template Name`,
				},
				resource.Attribute{
					Name:        "template_reqmod_icap",
					Description: `ICAP reqmod template (reqmod-icap template name)`,
				},
				resource.Attribute{
					Name:        "template_respmod_icap",
					Description: `ICAP respmod service template (respmod-icap template name)`,
				},
				resource.Attribute{
					Name:        "template_server_ssh",
					Description: `Server SSH Template (Server SSH Config Name)`,
				},
				resource.Attribute{
					Name:        "template_server_ssl",
					Description: `Server Side SSL Template (Server SSL Name)`,
				},
				resource.Attribute{
					Name:        "template_server_ssl_shared",
					Description: `Server SSL Template Name`,
				},
				resource.Attribute{
					Name:        "template_sip",
					Description: `SIP template`,
				},
				resource.Attribute{
					Name:        "template_smpp",
					Description: `SMPP template`,
				},
				resource.Attribute{
					Name:        "template_smtp",
					Description: `SMTP Template (SMTP Config Name)`,
				},
				resource.Attribute{
					Name:        "template_ssli",
					Description: `SSLi template (SSLi Template Name)`,
				},
				resource.Attribute{
					Name:        "template_tcp",
					Description: `L4 TCP Template (TCP Config Name)`,
				},
				resource.Attribute{
					Name:        "template_tcp_proxy",
					Description: `TCP Proxy Template Name`,
				},
				resource.Attribute{
					Name:        "template_tcp_proxy_client",
					Description: `TCP Proxy Config Client (TCP Proxy Config name)`,
				},
				resource.Attribute{
					Name:        "template_tcp_proxy_server",
					Description: `TCP Proxy Config Server (TCP Proxy Config name)`,
				},
				resource.Attribute{
					Name:        "template_tcp_proxy_shared",
					Description: `TCP Proxy Template name`,
				},
				resource.Attribute{
					Name:        "template_tcp_shared",
					Description: `TCP Template Name`,
				},
				resource.Attribute{
					Name:        "template_udp",
					Description: `L4 UDP Template (UDP Config Name)`,
				},
				resource.Attribute{
					Name:        "template_udp_shared",
					Description: `UDP Template Name`,
				},
				resource.Attribute{
					Name:        "template_virtual_port",
					Description: `Virtual port template (Virtual port template name)`,
				},
				resource.Attribute{
					Name:        "template_virtual_port_shared",
					Description: `Virtual Port Template Name`,
				},
				resource.Attribute{
					Name:        "trunk_fwd",
					Description: `Trunk interface number`,
				},
				resource.Attribute{
					Name:        "trunk_rev",
					Description: `Trunk interface number`,
				},
				resource.Attribute{
					Name:        "use_alternate_port",
					Description: `Use alternate virtual port`,
				},
				resource.Attribute{
					Name:        "use_cgnv6",
					Description: `Follow CGNv6 source NAT configuration`,
				},
				resource.Attribute{
					Name:        "use_default_if_no_server",
					Description: `Use default forwarding if server selection failed`,
				},
				resource.Attribute{
					Name:        "use_rcv_hop_for_resp",
					Description: `Use receive hop for response to client(For packets on default-vlan, also config “vlan-global enable-def-vlan-l2-forwarding”.)`,
				},
				resource.Attribute{
					Name:        "view",
					Description: `Specify a GSLB View (ID)`,
				},
				resource.Attribute{
					Name:        "waf_template",
					Description: `WAF template (WAF Template Name)`,
				},
				resource.Attribute{
					Name:        "when_down",
					Description: `Use alternate virtual port when down`,
				},
				resource.Attribute{
					Name:        "when_down_protocol2",
					Description: `Use alternate virtual port when down`,
				},
				resource.Attribute{
					Name:        "acl_name_seq_num",
					Description: `Specify ACL precedence (sequence-number)`,
				},
				resource.Attribute{
					Name:        "acl_name_seq_num_shared",
					Description: `Specify ACL precedence (sequence-number)`,
				},
				resource.Attribute{
					Name:        "acl_name_src_nat_pool",
					Description: `Policy based Source NAT (NAT Pool or Pool Group)`,
				},
				resource.Attribute{
					Name:        "acl_name_src_nat_pool_shared",
					Description: `Policy based Source NAT (NAT Pool or Pool Group)`,
				},
				resource.Attribute{
					Name:        "shared_partition_pool_name",
					Description: `Policy based Source NAT from shared partition`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `‘all’: all; ‘curr_conn’: Current connections; ‘total_l4_conn’: Total L4 connections; ‘total_l7_conn’: Total L7 connections; ‘total_tcp_conn’: Total TCP connections; ‘total_conn’: Total connections; ‘total_fwd_bytes’: Total forward bytes; ‘total_fwd_pkts’: Total forward packets; ‘total_rev_bytes’: Total reverse bytes; ‘total_rev_pkts’: Total reverse packets; ‘total_dns_pkts’: Total DNS packets; ‘total_mf_dns_pkts’: Total MF DNS packets; ‘es_total_failure_actions’: Total failure actions; ‘compression_bytes_before’: Data into compression engine; ‘compression_bytes_after’: Data out of compression engine; ‘compression_hit’: Number of requests compressed; ‘compression_miss’: Number of requests NOT compressed; ‘compression_miss_no_client’: Compression miss no client; ‘compression_miss_template_exclusion’: Compression miss template exclusion; ‘curr_req’: Current requests; ‘total_req’: Total requests; ‘total_req_succ’: Total successful requests; ‘peak_conn’: Peak connections; ‘curr_conn_rate’: Current connection rate; ‘last_rsp_time’: Last response time; ‘fastest_rsp_time’: Fastest response time; ‘slowest_rsp_time’: Slowest response time; ‘loc_permit’: Permit number; ‘loc_deny’: Deny number; ‘loc_conn’: Connection number; ‘curr_ssl_conn’: Current SSL connections; ‘total_ssl_conn’: Total SSL connections;`,
				},
				resource.Attribute{
					Name:        "acl_id_seq_num",
					Description: `Specify ACL precedence (sequence-number)`,
				},
				resource.Attribute{
					Name:        "acl_id_seq_num_shared",
					Description: `Specify ACL precedence (sequence-number)`,
				},
				resource.Attribute{
					Name:        "acl_id_src_nat_pool",
					Description: `Policy based Source NAT (NAT Pool or Pool Group)`,
				},
				resource.Attribute{
					Name:        "acl_id_src_nat_pool_shared",
					Description: `Policy based Source NAT (NAT Pool or Pool Group)`,
				},
				resource.Attribute{
					Name:        "shared_partition_pool_id",
					Description: `Policy based Source NAT from shared partition`,
				},
				resource.Attribute{
					Name:        "aflex",
					Description: `Bind aFleX Script to the Virtual Port (aFleX Script Name)`,
				},
				resource.Attribute{
					Name:        "aflex_shared",
					Description: `aFleX Script Name`,
				},
				resource.Attribute{
					Name:        "aaa_policy",
					Description: `Specify AAA policy name to bind to the virtual port`,
				},
				resource.Attribute{
					Name:        "cancel_migration",
					Description: `Cancel migration`,
				},
				resource.Attribute{
					Name:        "finish_migration",
					Description: `Complete the migration`,
				},
				resource.Attribute{
					Name:        "target_data_cpu",
					Description: `Number of CPUs on the target platform`,
				},
				resource.Attribute{
					Name:        "target_floating_ipv4",
					Description: `Specify IP address`,
				},
				resource.Attribute{
					Name:        "target_floating_ipv6",
					Description: `Specify IPv6 address`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_vrrp_common",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder vrrp common resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `‘enable’: enable vrrp-a; ‘disable’: disable vrrp-a;`,
				},
				resource.Attribute{
					Name:        "arp_retry",
					Description: `Number of additional gratuitous ARPs sent out after HA failover (1-255, default is 4)`,
				},
				resource.Attribute{
					Name:        "dead_timer",
					Description: `VRRP-A dead timer in terms of how many hello messages missed, default is 5 (2-255, default is 5)`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `Unique ID for each VRRP-A box (Device-id number)`,
				},
				resource.Attribute{
					Name:        "disable_default_vrid",
					Description: `Disable default vrid`,
				},
				resource.Attribute{
					Name:        "forward_l4_packet_on_standby",
					Description: `Enables Layer 2/3 forwarding of Layer 4 traffic on the Standby ACOS device`,
				},
				resource.Attribute{
					Name:        "get_ready_time",
					Description: `set get ready time after ax starting up (60-1200, in unit of 100millisec, default is 60)`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `VRRP-A Hello Interval (1-255, in unit of 100millisec, default is 2)`,
				},
				resource.Attribute{
					Name:        "preemption_delay",
					Description: `Delay before changing state from Active to Standby (1-255, in unit of 100millisec, default is 60)`,
				},
				resource.Attribute{
					Name:        "restart_time",
					Description: `Time between restarting ports on standby system after transition`,
				},
				resource.Attribute{
					Name:        "set_id",
					Description: `Set-ID for HA configuration (Set id from 1 to 15)`,
				},
				resource.Attribute{
					Name:        "track_event_delay",
					Description: `Delay before changing state after up/down event (Units of 100 milliseconds (default 30))`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "hostid_append_to_vrid_default",
					Description: `hostid append to vrid default`,
				},
				resource.Attribute{
					Name:        "hostid_append_to_vrid_value",
					Description: `hostid append to vrid num`,
				},
				resource.Attribute{
					Name:        "inline_mode",
					Description: `Enable Layer 2 Inline Hot Standby Mode`,
				},
				resource.Attribute{
					Name:        "preferred_port",
					Description: `Preferred ethernet Port`,
				},
				resource.Attribute{
					Name:        "preferred_trunk",
					Description: `Preferred trunk Port`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_vrrp_peer_group",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder vrrp peer group resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "ip_peer_address",
					Description: `IP Address`,
				},
				resource.Attribute{
					Name:        "ipv6_peer_address",
					Description: `IPV6 address`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_vrrp_session_sync",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder vrrp session sync resource for A10`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "vthunder_thunder_vrrp_vrid",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder vrrp vrid resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "vrid_val",
					Description: `Specify ha VRRP-A vrid`,
				},
				resource.Attribute{
					Name:        "fail_over_policy_template",
					Description: `Apply a fail over policy template (VRRP-A fail over policy template name)`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `VRRP-A priorty (Priority, default is 150)`,
				},
				resource.Attribute{
					Name:        "priority_cost",
					Description: `The amount the priority will decrease`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `VLAN tracking (VLAN id)`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Route’s administrative distance (default: match any)`,
				},
				resource.Attribute{
					Name:        "gatewayv6",
					Description: `Match the route’s gateway (next-hop) (default: match any)`,
				},
				resource.Attribute{
					Name:        "ipv6_destination",
					Description: `IPv6 Destination Prefix`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `‘any’: Match any routing protocol (default); ‘static’: Match only static routes (added by user); ‘dynamic’: Match routes added by dynamic routing protocols (e.g. OSPF);`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Match the route’s gateway (next-hop) (default: match any)`,
				},
				resource.Attribute{
					Name:        "ip_destination",
					Description: `Destination prefix`,
				},
				resource.Attribute{
					Name:        "mask",
					Description: `Destination prefix mask`,
				},
				resource.Attribute{
					Name:        "bgp_ipv4_address",
					Description: `bgp IP Address`,
				},
				resource.Attribute{
					Name:        "bgp_ipv6_address",
					Description: `IPV6 address`,
				},
				resource.Attribute{
					Name:        "ethernet",
					Description: `Ethernet (for link-local address only)`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP Address`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `IPV6 address`,
				},
				resource.Attribute{
					Name:        "per_port_pri",
					Description: `per port priority`,
				},
				resource.Attribute{
					Name:        "trunk",
					Description: `Trunk interface (for link-local address only)`,
				},
				resource.Attribute{
					Name:        "ipv6_address_partition",
					Description: `IPV6 address`,
				},
				resource.Attribute{
					Name:        "ve",
					Description: `VE interface (for link-local address only)`,
				},
				resource.Attribute{
					Name:        "ip_address_partition",
					Description: `IP Address`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `disable preemption`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `preemption threshold (preemption threshhold (0-255), default 0)`,
				},
				resource.Attribute{
					Name:        "vrid_lead",
					Description: `Define a VRRP-A VRID leader`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_web_category",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder web category resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "web-category",
					Description: `Web-Category Commands`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `BrightCloud Query Server`,
				},
				resource.Attribute{
					Name:        "database-server",
					Description: `BrightCloud Database Server`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `BrightCloud Query Server Listening Port(default 80)`,
				},
				resource.Attribute{
					Name:        "ssl-port",
					Description: `BrightCloud Servers SSL Port(default 443)`,
				},
				resource.Attribute{
					Name:        "server-timeout",
					Description: `BrightCloud Servers Timeout in seconds (default: 15s)`,
				},
				resource.Attribute{
					Name:        "cloud-query-disable",
					Description: `Disables cloud queries for URL's not present in local database(default enable)`,
				},
				resource.Attribute{
					Name:        "cloud-query-cache-size",
					Description: `Maximum cache size for storing cloud query results`,
				},
				resource.Attribute{
					Name:        "db-update-time",
					Description: `Time of day to update database (default: 00:00)`,
				},
				resource.Attribute{
					Name:        "rtu-update-disable",
					Description: `Disables real time updates(default enable)`,
				},
				resource.Attribute{
					Name:        "rtu-update-interval",
					Description: `Interval to check for real time updates if enabled in mins(default 60)`,
				},
				resource.Attribute{
					Name:        "rtu-cache-size",
					Description: `Maximum cache size for storing RTU updates`,
				},
				resource.Attribute{
					Name:        "use-mgmt-port",
					Description: `Use management interface for all communication with BrightCloud`,
				},
				resource.Attribute{
					Name:        "remote-syslog-enable",
					Description: `Enable data plane logging to a remote syslog server`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Enable BrightCloud SDK`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "proxy-host",
					Description: `Proxy server hostname or IP address`,
				},
				resource.Attribute{
					Name:        "http-port",
					Description: `Proxy server HTTP port`,
				},
				resource.Attribute{
					Name:        "https-port",
					Description: `Proxy server HTTPS port(HTTP port will be used if not configured)`,
				},
				resource.Attribute{
					Name:        "auth-type",
					Description: `'ntlm': NTLM authentication(default); 'basic': Basic authentication;`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Realm for NTLM authentication`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for proxy authentication`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for proxy authentication`,
				},
				resource.Attribute{
					Name:        "secret-string",
					Description: `password value`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED secret string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Reputation Scope name`,
				},
				resource.Attribute{
					Name:        "uncategorized",
					Description: `Uncategorized URLs`,
				},
				resource.Attribute{
					Name:        "real-estate",
					Description: `Category Real Estate`,
				},
				resource.Attribute{
					Name:        "computer-and-internet-security",
					Description: `Category Computer and Internet Security`,
				},
				resource.Attribute{
					Name:        "financial-services",
					Description: `Category Financial Services`,
				},
				resource.Attribute{
					Name:        "business-and-economy",
					Description: `Category Business and Economy`,
				},
				resource.Attribute{
					Name:        "computer-and-internet-info",
					Description: `Category Computer and Internet Info`,
				},
				resource.Attribute{
					Name:        "auctions",
					Description: `Category Auctions`,
				},
				resource.Attribute{
					Name:        "shopping",
					Description: `Category Shopping`,
				},
				resource.Attribute{
					Name:        "cult-and-occult",
					Description: `Category Cult and Occult`,
				},
				resource.Attribute{
					Name:        "travel",
					Description: `Category Travel`,
				},
				resource.Attribute{
					Name:        "drugs",
					Description: `Category Abused Drugs`,
				},
				resource.Attribute{
					Name:        "adult-and-pornography",
					Description: `Category Adult and Pornography`,
				},
				resource.Attribute{
					Name:        "home-and-garden",
					Description: `Category Home and Garden`,
				},
				resource.Attribute{
					Name:        "military",
					Description: `Category Military`,
				},
				resource.Attribute{
					Name:        "social-network",
					Description: `Category Social Network`,
				},
				resource.Attribute{
					Name:        "dead-sites",
					Description: `Category Dead Sites (db Ops only)`,
				},
				resource.Attribute{
					Name:        "stock-advice-and-tools",
					Description: `Category Stock Advice and Tools`,
				},
				resource.Attribute{
					Name:        "training-and-tools",
					Description: `Category Training and Tools`,
				},
				resource.Attribute{
					Name:        "dating",
					Description: `Category Dating`,
				},
				resource.Attribute{
					Name:        "sex-education",
					Description: `Category Sex Education`,
				},
				resource.Attribute{
					Name:        "religion",
					Description: `Category Religion`,
				},
				resource.Attribute{
					Name:        "entertainment-and-arts",
					Description: `Category Entertainment and Arts`,
				},
				resource.Attribute{
					Name:        "personal-sites-and-blogs",
					Description: `Category Personal sites and Blogs`,
				},
				resource.Attribute{
					Name:        "legal",
					Description: `Category Legal`,
				},
				resource.Attribute{
					Name:        "local-information",
					Description: `Category Local Information`,
				},
				resource.Attribute{
					Name:        "streaming-media",
					Description: `Category Streaming Media`,
				},
				resource.Attribute{
					Name:        "job-search",
					Description: `Category Job Search`,
				},
				resource.Attribute{
					Name:        "gambling",
					Description: `Category Gambling`,
				},
				resource.Attribute{
					Name:        "translation",
					Description: `Category Translation`,
				},
				resource.Attribute{
					Name:        "reference-and-research",
					Description: `Category Reference and Research`,
				},
				resource.Attribute{
					Name:        "shareware-and-freeware",
					Description: `Category Shareware and Freeware`,
				},
				resource.Attribute{
					Name:        "peer-to-peer",
					Description: `Category Peer to Peer`,
				},
				resource.Attribute{
					Name:        "marijuana",
					Description: `Category Marijuana`,
				},
				resource.Attribute{
					Name:        "hacking",
					Description: `Category Hacking`,
				},
				resource.Attribute{
					Name:        "games",
					Description: `Category Games`,
				},
				resource.Attribute{
					Name:        "philosophy-and-politics",
					Description: `Category Philosophy and Political Advocacy`,
				},
				resource.Attribute{
					Name:        "weapons",
					Description: `Category Weapons`,
				},
				resource.Attribute{
					Name:        "pay-to-surf",
					Description: `Category Pay to Surf`,
				},
				resource.Attribute{
					Name:        "hunting-and-fishing",
					Description: `Category Hunting and Fishing`,
				},
				resource.Attribute{
					Name:        "society",
					Description: `Category Society`,
				},
				resource.Attribute{
					Name:        "educational-institutions",
					Description: `Category Educational Institutions`,
				},
				resource.Attribute{
					Name:        "online-greeting-cards",
					Description: `Category Online Greeting cards`,
				},
				resource.Attribute{
					Name:        "sports",
					Description: `Category Sports`,
				},
				resource.Attribute{
					Name:        "swimsuits-and-intimate-apparel",
					Description: `Category Swimsuits and Intimate Apparel`,
				},
				resource.Attribute{
					Name:        "questionable",
					Description: `Category Questionable`,
				},
				resource.Attribute{
					Name:        "kids",
					Description: `Category Kids`,
				},
				resource.Attribute{
					Name:        "hate-and-racism",
					Description: `Category Hate and Racism`,
				},
				resource.Attribute{
					Name:        "personal-storage",
					Description: `Category Personal Storage`,
				},
				resource.Attribute{
					Name:        "violence",
					Description: `Category Violence`,
				},
				resource.Attribute{
					Name:        "keyloggers-and-monitoring",
					Description: `Category Keyloggers and Monitoring`,
				},
				resource.Attribute{
					Name:        "search-engines",
					Description: `Category Search Engines`,
				},
				resource.Attribute{
					Name:        "internet-portals",
					Description: `Category Internet Portals`,
				},
				resource.Attribute{
					Name:        "web-advertisements",
					Description: `Category Web Advertisements`,
				},
				resource.Attribute{
					Name:        "cheating",
					Description: `Category Cheating`,
				},
				resource.Attribute{
					Name:        "gross",
					Description: `Category Gross`,
				},
				resource.Attribute{
					Name:        "web-based-email",
					Description: `Category Web based email`,
				},
				resource.Attribute{
					Name:        "malware-sites",
					Description: `Category Malware Sites`,
				},
				resource.Attribute{
					Name:        "phishing-and-other-fraud",
					Description: `Category Phishing and Other Frauds`,
				},
				resource.Attribute{
					Name:        "proxy-avoid-and-anonymizers",
					Description: `Category Proxy Avoid and Anonymizers`,
				},
				resource.Attribute{
					Name:        "spyware-and-adware",
					Description: `Category Spyware and Adware`,
				},
				resource.Attribute{
					Name:        "music",
					Description: `Category Music`,
				},
				resource.Attribute{
					Name:        "government",
					Description: `Category Government`,
				},
				resource.Attribute{
					Name:        "nudity",
					Description: `Category Nudity`,
				},
				resource.Attribute{
					Name:        "news-and-media",
					Description: `Category News and Media`,
				},
				resource.Attribute{
					Name:        "illegal",
					Description: `Category Illegal`,
				},
				resource.Attribute{
					Name:        "cdns",
					Description: `Category CDNs`,
				},
				resource.Attribute{
					Name:        "internet-communications",
					Description: `Category Internet Communications`,
				},
				resource.Attribute{
					Name:        "bot-nets",
					Description: `Category Bot Nets`,
				},
				resource.Attribute{
					Name:        "abortion",
					Description: `Category Abortion`,
				},
				resource.Attribute{
					Name:        "health-and-medicine",
					Description: `Category Health and Medicine`,
				},
				resource.Attribute{
					Name:        "confirmed-spam-sources",
					Description: `Category Confirmed SPAM Sources`,
				},
				resource.Attribute{
					Name:        "spam-urls",
					Description: `Category SPAM URLs`,
				},
				resource.Attribute{
					Name:        "unconfirmed-spam-sources",
					Description: `Category Unconfirmed SPAM Sources`,
				},
				resource.Attribute{
					Name:        "open-http-proxies",
					Description: `Category Open HTTP Proxies`,
				},
				resource.Attribute{
					Name:        "dynamic-comment",
					Description: `Category Dynamic Comment`,
				},
				resource.Attribute{
					Name:        "parked-domains",
					Description: `Category Parked Domains`,
				},
				resource.Attribute{
					Name:        "alcohol-and-tobacco",
					Description: `Category Alcohol and Tobacco`,
				},
				resource.Attribute{
					Name:        "private-ip-addresses",
					Description: `Category Private IP Addresses`,
				},
				resource.Attribute{
					Name:        "image-and-video-search",
					Description: `Category Image and Video Search`,
				},
				resource.Attribute{
					Name:        "fashion-and-beauty",
					Description: `Category Fashion and Beauty`,
				},
				resource.Attribute{
					Name:        "recreation-and-hobbies",
					Description: `Category Recreation and Hobbies`,
				},
				resource.Attribute{
					Name:        "motor-vehicles",
					Description: `Category Motor Vehicles`,
				},
				resource.Attribute{
					Name:        "web-hosting-sites",
					Description: `Category Web Hosting Sites`,
				},
				resource.Attribute{
					Name:        "food-and-dining",
					Description: `Category Food and Dining`,
				},
				resource.Attribute{
					Name:        "user-tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all': all; 'trustworthy': Trustworthy level(81-100); 'low-risk': Low-risk level(61-80); 'moderate-risk': Moderate-risk level(41-60); 'suspicious': Suspicious level(21-40); 'malicious': Malicious level(1-20);`,
				},
				resource.Attribute{
					Name:        "greater-trustworthy",
					Description: `Reputation score is greater than or equal to 81`,
				},
				resource.Attribute{
					Name:        "greater-low-risk",
					Description: `Reputation score is greater than or equal to 61`,
				},
				resource.Attribute{
					Name:        "greater-moderate-risk",
					Description: `Reputation score is greater than or equal to 41`,
				},
				resource.Attribute{
					Name:        "greater-suspicious",
					Description: `Reputation score is greater than or equal to 21`,
				},
				resource.Attribute{
					Name:        "greater-malicious",
					Description: `Reputation score is greater than or equal to 1`,
				},
				resource.Attribute{
					Name:        "greater-threshold",
					Description: `Reputation score is greater than or equal to the customized score (1-100)`,
				},
				resource.Attribute{
					Name:        "less-trustworthy",
					Description: `Reputation score is less than or equal to 100`,
				},
				resource.Attribute{
					Name:        "less-low-risk",
					Description: `Reputation score is less than or equal to 80`,
				},
				resource.Attribute{
					Name:        "less-moderate-risk",
					Description: `Reputation score is less than or equal to 60`,
				},
				resource.Attribute{
					Name:        "less-suspicious",
					Description: `Reputation score is less than or equal to 40`,
				},
				resource.Attribute{
					Name:        "less-malicious",
					Description: `Reputation score is less than or equal to 20`,
				},
				resource.Attribute{
					Name:        "less-threshold",
					Description: `Reputation score is less than or equal to a customized value (1-100)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_web_category_category_list",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder web category category list resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "category-list",
					Description: `List of web categories`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Web Category List name`,
				},
				resource.Attribute{
					Name:        "uncategorized",
					Description: `Uncategorized URLs`,
				},
				resource.Attribute{
					Name:        "real-estate",
					Description: `Category Real Estate`,
				},
				resource.Attribute{
					Name:        "computer-and-internet-security",
					Description: `Category Computer and Internet Security`,
				},
				resource.Attribute{
					Name:        "financial-services",
					Description: `Category Financial Services`,
				},
				resource.Attribute{
					Name:        "business-and-economy",
					Description: `Category Business and Economy`,
				},
				resource.Attribute{
					Name:        "computer-and-internet-info",
					Description: `Category Computer and Internet Info`,
				},
				resource.Attribute{
					Name:        "auctions",
					Description: `Category Auctions`,
				},
				resource.Attribute{
					Name:        "shopping",
					Description: `Category Shopping`,
				},
				resource.Attribute{
					Name:        "cult-and-occult",
					Description: `Category Cult and Occult`,
				},
				resource.Attribute{
					Name:        "travel",
					Description: `Category Travel`,
				},
				resource.Attribute{
					Name:        "drugs",
					Description: `Category Abused Drugs`,
				},
				resource.Attribute{
					Name:        "adult-and-pornography",
					Description: `Category Adult and Pornography`,
				},
				resource.Attribute{
					Name:        "home-and-garden",
					Description: `Category Home and Garden`,
				},
				resource.Attribute{
					Name:        "military",
					Description: `Category Military`,
				},
				resource.Attribute{
					Name:        "social-network",
					Description: `Category Social Network`,
				},
				resource.Attribute{
					Name:        "dead-sites",
					Description: `Category Dead Sites (db Ops only)`,
				},
				resource.Attribute{
					Name:        "stock-advice-and-tools",
					Description: `Category Stock Advice and Tools`,
				},
				resource.Attribute{
					Name:        "training-and-tools",
					Description: `Category Training and Tools`,
				},
				resource.Attribute{
					Name:        "dating",
					Description: `Category Dating`,
				},
				resource.Attribute{
					Name:        "sex-education",
					Description: `Category Sex Education`,
				},
				resource.Attribute{
					Name:        "religion",
					Description: `Category Religion`,
				},
				resource.Attribute{
					Name:        "entertainment-and-arts",
					Description: `Category Entertainment and Arts`,
				},
				resource.Attribute{
					Name:        "personal-sites-and-blogs",
					Description: `Category Personal sites and Blogs`,
				},
				resource.Attribute{
					Name:        "legal",
					Description: `Category Legal`,
				},
				resource.Attribute{
					Name:        "local-information",
					Description: `Category Local Information`,
				},
				resource.Attribute{
					Name:        "streaming-media",
					Description: `Category Streaming Media`,
				},
				resource.Attribute{
					Name:        "job-search",
					Description: `Category Job Search`,
				},
				resource.Attribute{
					Name:        "gambling",
					Description: `Category Gambling`,
				},
				resource.Attribute{
					Name:        "translation",
					Description: `Category Translation`,
				},
				resource.Attribute{
					Name:        "reference-and-research",
					Description: `Category Reference and Research`,
				},
				resource.Attribute{
					Name:        "shareware-and-freeware",
					Description: `Category Shareware and Freeware`,
				},
				resource.Attribute{
					Name:        "peer-to-peer",
					Description: `Category Peer to Peer`,
				},
				resource.Attribute{
					Name:        "marijuana",
					Description: `Category Marijuana`,
				},
				resource.Attribute{
					Name:        "hacking",
					Description: `Category Hacking`,
				},
				resource.Attribute{
					Name:        "games",
					Description: `Category Games`,
				},
				resource.Attribute{
					Name:        "philosophy-and-politics",
					Description: `Category Philosophy and Political Advocacy`,
				},
				resource.Attribute{
					Name:        "weapons",
					Description: `Category Weapons`,
				},
				resource.Attribute{
					Name:        "pay-to-surf",
					Description: `Category Pay to Surf`,
				},
				resource.Attribute{
					Name:        "hunting-and-fishing",
					Description: `Category Hunting and Fishing`,
				},
				resource.Attribute{
					Name:        "society",
					Description: `Category Society`,
				},
				resource.Attribute{
					Name:        "educational-institutions",
					Description: `Category Educational Institutions`,
				},
				resource.Attribute{
					Name:        "online-greeting-cards",
					Description: `Category Online Greeting cards`,
				},
				resource.Attribute{
					Name:        "sports",
					Description: `Category Sports`,
				},
				resource.Attribute{
					Name:        "swimsuits-and-intimate-apparel",
					Description: `Category Swimsuits and Intimate Apparel`,
				},
				resource.Attribute{
					Name:        "questionable",
					Description: `Category Questionable`,
				},
				resource.Attribute{
					Name:        "kids",
					Description: `Category Kids`,
				},
				resource.Attribute{
					Name:        "hate-and-racism",
					Description: `Category Hate and Racism`,
				},
				resource.Attribute{
					Name:        "personal-storage",
					Description: `Category Personal Storage`,
				},
				resource.Attribute{
					Name:        "violence",
					Description: `Category Violence`,
				},
				resource.Attribute{
					Name:        "keyloggers-and-monitoring",
					Description: `Category Keyloggers and Monitoring`,
				},
				resource.Attribute{
					Name:        "search-engines",
					Description: `Category Search Engines`,
				},
				resource.Attribute{
					Name:        "internet-portals",
					Description: `Category Internet Portals`,
				},
				resource.Attribute{
					Name:        "web-advertisements",
					Description: `Category Web Advertisements`,
				},
				resource.Attribute{
					Name:        "cheating",
					Description: `Category Cheating`,
				},
				resource.Attribute{
					Name:        "gross",
					Description: `Category Gross`,
				},
				resource.Attribute{
					Name:        "web-based-email",
					Description: `Category Web based email`,
				},
				resource.Attribute{
					Name:        "malware-sites",
					Description: `Category Malware Sites`,
				},
				resource.Attribute{
					Name:        "phishing-and-other-fraud",
					Description: `Category Phishing and Other Frauds`,
				},
				resource.Attribute{
					Name:        "proxy-avoid-and-anonymizers",
					Description: `Category Proxy Avoid and Anonymizers`,
				},
				resource.Attribute{
					Name:        "spyware-and-adware",
					Description: `Category Spyware and Adware`,
				},
				resource.Attribute{
					Name:        "music",
					Description: `Category Music`,
				},
				resource.Attribute{
					Name:        "government",
					Description: `Category Government`,
				},
				resource.Attribute{
					Name:        "nudity",
					Description: `Category Nudity`,
				},
				resource.Attribute{
					Name:        "news-and-media",
					Description: `Category News and Media`,
				},
				resource.Attribute{
					Name:        "illegal",
					Description: `Category Illegal`,
				},
				resource.Attribute{
					Name:        "cdns",
					Description: `Category CDNs`,
				},
				resource.Attribute{
					Name:        "internet-communications",
					Description: `Category Internet Communications`,
				},
				resource.Attribute{
					Name:        "bot-nets",
					Description: `Category Bot Nets`,
				},
				resource.Attribute{
					Name:        "abortion",
					Description: `Category Abortion`,
				},
				resource.Attribute{
					Name:        "health-and-medicine",
					Description: `Category Health and Medicine`,
				},
				resource.Attribute{
					Name:        "confirmed-spam-sources",
					Description: `Category Confirmed SPAM Sources`,
				},
				resource.Attribute{
					Name:        "spam-urls",
					Description: `Category SPAM URLs`,
				},
				resource.Attribute{
					Name:        "unconfirmed-spam-sources",
					Description: `Category Unconfirmed SPAM Sources`,
				},
				resource.Attribute{
					Name:        "open-http-proxies",
					Description: `Category Open HTTP Proxies`,
				},
				resource.Attribute{
					Name:        "dynamic-comment",
					Description: `Category Dynamic Comment`,
				},
				resource.Attribute{
					Name:        "parked-domains",
					Description: `Category Parked Domains`,
				},
				resource.Attribute{
					Name:        "alcohol-and-tobacco",
					Description: `Category Alcohol and Tobacco`,
				},
				resource.Attribute{
					Name:        "private-ip-addresses",
					Description: `Category Private IP Addresses`,
				},
				resource.Attribute{
					Name:        "image-and-video-search",
					Description: `Category Image and Video Search`,
				},
				resource.Attribute{
					Name:        "fashion-and-beauty",
					Description: `Category Fashion and Beauty`,
				},
				resource.Attribute{
					Name:        "recreation-and-hobbies",
					Description: `Category Recreation and Hobbies`,
				},
				resource.Attribute{
					Name:        "motor-vehicles",
					Description: `Category Motor Vehicles`,
				},
				resource.Attribute{
					Name:        "web-hosting-sites",
					Description: `Category Web Hosting Sites`,
				},
				resource.Attribute{
					Name:        "food-and-dining",
					Description: `Category Food and Dining`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "user-tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all': all; 'uncategorized': uncategorized category; 'real-estate': real estate category; 'computer-and-internet-security': computer and internet security category; 'financial-services': financial services category; 'business-and-economy': business and economy category; 'computer-and-internet-info': computer and internet info category; 'auctions': auctions category; 'shopping': shopping category; 'cult-and-occult': cult and occult category; 'travel': travel category; 'drugs': drugs category; 'adult-and-pornography': adult and pornography category; 'home-and-garden': home and garden category; 'military': military category; 'social-network': social network category; 'dead-sites': dead sites category; 'stock-advice-and-tools': stock advice and tools category; 'training-and-tools': training and tools category; 'dating': dating category; 'sex-education': sex education category; 'religion': religion category; 'entertainment-and-arts': entertainment and arts category; 'personal-sites-and-blogs': personal sites and blogs category; 'legal': legal category; 'local-information': local information category; 'streaming-media': streaming media category; 'job-search': job search category; 'gambling': gambling category; 'translation': translation category; 'reference-and-research': reference and research category; 'shareware-and-freeware': shareware and freeware category; 'peer-to-peer': peer to peer category; 'marijuana': marijuana category; 'hacking': hacking category; 'games': games category; 'philosophy-and-politics': philosophy and politics category; 'weapons': weapons category; 'pay-to-surf': pay to surf category; 'hunting-and-fishing': hunting and fishing category; 'society': society category; 'educational-institutions': educational institutions category; 'online-greeting-cards': online greeting cards category; 'sports': sports category; 'swimsuits-and-intimate-apparel': swimsuits and intimate apparel category; 'questionable': questionable category; 'kids': kids category; 'hate-and-racism': hate and racism category; 'personal-storage': personal storage category; 'violence': violence category; 'keyloggers-and-monitoring': keyloggers and monitoring category; 'search-engines': search engines category; 'internet-portals': internet portals category; 'web-advertisements': web advertisements category; 'cheating': cheating category; 'gross': gross category; 'web-based-email': web based email category; 'malware-sites': malware sites category; 'phishing-and-other-fraud': phishing and other fraud category; 'proxy-avoid-and-anonymizers': proxy avoid and anonymizers category; 'spyware-and-adware': spyware and adware category; 'music': music category; 'government': government category; 'nudity': nudity category; 'news-and-media': news and media category; 'illegal': illegal category; 'CDNs': content delivery networks category; 'internet-communications': internet communications category; 'bot-nets': bot nets category; 'abortion': abortion category; 'health-and-medicine': health and medicine category; 'confirmed-SPAM-sources': confirmed SPAM sources category; 'SPAM-URLs': SPAM URLs category; 'unconfirmed-SPAM-sources': unconfirmed SPAM sources category; 'open-HTTP-proxies': open HTTP proxies category; 'dynamic-comment': dynamic comment category; 'parked-domains': parked domains category; 'alcohol-and-tobacco': alcohol and tobacco category; 'private-IP-addresses': private IP addresses category; 'image-and-video-search': image and video search category; 'fashion-and-beauty': fashion and beauty category; 'recreation-and-hobbies': recreation and hobbies category; 'motor-vehicles': motor vehicles category; 'web-hosting-sites': web hosting sites category; 'food-and-dining': food and dining category;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_web_category_proxy_server",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder web category proxy server resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "proxy-server",
					Description: `Commands to connect web-category through proxy server`,
				},
				resource.Attribute{
					Name:        "proxy-host",
					Description: `Proxy server hostname or IP address`,
				},
				resource.Attribute{
					Name:        "http-port",
					Description: `Proxy server HTTP port`,
				},
				resource.Attribute{
					Name:        "https-port",
					Description: `Proxy server HTTPS port(HTTP port will be used if not configured)`,
				},
				resource.Attribute{
					Name:        "auth-type",
					Description: `'ntlm': NTLM authentication(default); 'basic': Basic authentication;`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Realm for NTLM authentication`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for proxy authentication`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for proxy authentication`,
				},
				resource.Attribute{
					Name:        "secret-string",
					Description: `password value`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED secret string)`,
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
			Type:             "vthunder_thunder_web_category_reputation_scope",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder web category reputation scope resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "reputation-scope",
					Description: `Configure the scope of reputation score`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Reputation Scope name`,
				},
				resource.Attribute{
					Name:        "greater-trustworthy",
					Description: `Reputation score is greater than or equal to 81`,
				},
				resource.Attribute{
					Name:        "greater-low-risk",
					Description: `Reputation score is greater than or equal to 61`,
				},
				resource.Attribute{
					Name:        "greater-moderate-risk",
					Description: `Reputation score is greater than or equal to 41`,
				},
				resource.Attribute{
					Name:        "greater-suspicious",
					Description: `Reputation score is greater than or equal to 21`,
				},
				resource.Attribute{
					Name:        "greater-malicious",
					Description: `Reputation score is greater than or equal to 1`,
				},
				resource.Attribute{
					Name:        "greater-threshold",
					Description: `Reputation score is greater than or equal to the customized score (1-100)`,
				},
				resource.Attribute{
					Name:        "less-trustworthy",
					Description: `Reputation score is less than or equal to 100`,
				},
				resource.Attribute{
					Name:        "less-low-risk",
					Description: `Reputation score is less than or equal to 80`,
				},
				resource.Attribute{
					Name:        "less-moderate-risk",
					Description: `Reputation score is less than or equal to 60`,
				},
				resource.Attribute{
					Name:        "less-suspicious",
					Description: `Reputation score is less than or equal to 40`,
				},
				resource.Attribute{
					Name:        "less-malicious",
					Description: `Reputation score is less than or equal to 20`,
				},
				resource.Attribute{
					Name:        "less-threshold",
					Description: `Reputation score is less than or equal to a customized value (1-100)`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "user-tag",
					Description: `Customized tag`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all': all; 'trustworthy': Trustworthy level(81-100); 'low-risk': Low-risk level(61-80); 'moderate-risk': Moderate-risk level(41-60); 'suspicious': Suspicious level(21-40); 'malicious': Malicious level(1-20);`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_web_category_statistics",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder web category statistics resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "statistics",
					Description: `Statistics`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `uuid of the object`,
				},
				resource.Attribute{
					Name:        "counters1",
					Description: `'all': all; 'db-lookup': db-lookup; 'cloud-cache-lookup': cloud-cache-lookup; 'cloud-lookup': cloud-lookup; 'rtu-lookup': rtu-lookup; 'lookup-latency': lookup-latency; 'db-mem': db-mem; 'rtu-cache-mem': rtu-cache-mem; 'lookup-cache-mem': lookup-cache-mem;`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vthunder_thunder_web_category_web_reputation",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder web category web reputation resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "web-reputation",
					Description: `Used for Web-Reputation Show Commands`,
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
			Type:             "vthunder_thunder_write_memory",
			Category:         "Resources",
			ShortDescription: `Provides details about thunder write memory resource for A10`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "destination",
					Description: `‘primary’: Write to default Primary Configuration; ‘secondary’: Write to default Secondary Configuration; ‘local’: Local Configuration Profile Name;`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `‘all’: All partition configurations; ‘shared’: Shared partition; ‘specified’: Specified partition;`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `Local Configuration Profile Name`,
				},
				resource.Attribute{
					Name:        "specified_partition",
					Description: `Specified partition`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"vthunder_thunder_access_list_extended":                                           0,
		"vthunder_thunder_access_list_standard":                                           1,
		"vthunder_thunder_active_partition":                                               2,
		"vthunder_thunder_bgp":                                                            3,
		"vthunder_thunder_configure_sync":                                                 4,
		"vthunder_thunder_dns_primary":                                                    5,
		"vthunder_thunder_ethernet":                                                       6,
		"vthunder_thunder_file_ssl_cert":                                                  7,
		"vthunder_thunder_fw_active_rule_set":                                             8,
		"vthunder_thunder_fw_alg_dns":                                                     9,
		"vthunder_thunder_fw_alg_ftp":                                                     10,
		"vthunder_thunder_fw_alg_icmp":                                                    11,
		"vthunder_thunder_fw_alg_pptp":                                                    12,
		"vthunder_thunder_fw_alg_rtsp":                                                    13,
		"vthunder_thunder_fw_alg_sip":                                                     14,
		"vthunder_thunder_fw_alg_tftp":                                                    15,
		"vthunder_thunder_fw_app":                                                         16,
		"vthunder_thunder_fw_apply_changes":                                               17,
		"vthunder_thunder_fw_clear_session_filter":                                        18,
		"vthunder_thunder_fw_full_cone_session":                                           19,
		"vthunder_thunder_fw_global":                                                      20,
		"vthunder_thunder_fw_gtp":                                                         21,
		"vthunder_thunder_fw_gtp_in_gtp_filtering":                                        22,
		"vthunder_thunder_fw_gtp_v0":                                                      23,
		"vthunder_thunder_fw_helper_sessions":                                             24,
		"vthunder_thunder_fw_limit_entry":                                                 25,
		"vthunder_thunder_fw_local_log":                                                   26,
		"vthunder_thunder_fw_logging":                                                     27,
		"vthunder_thunder_fw_radius_server":                                               28,
		"vthunder_thunder_fw_resource_usage":                                              29,
		"vthunder_thunder_fw_server":                                                      30,
		"vthunder_thunder_fw_service_group":                                               31,
		"vthunder_thunder_fw_status":                                                      32,
		"vthunder_thunder_fw_system_status":                                               33,
		"vthunder_thunder_fw_tap_monitor":                                                 34,
		"vthunder_thunder_fw_tcp_mss_clamp":                                               35,
		"vthunder_thunder_fw_tcp_reset_on_error":                                          36,
		"vthunder_thunder_fw_tcp_rst_close_immediate":                                     37,
		"vthunder_thunder_fw_tcp_window_check":                                            38,
		"vthunder_thunder_fw_template_logging":                                            39,
		"vthunder_thunder_fw_top_k_rules":                                                 40,
		"vthunder_thunder_fw_urpf":                                                        41,
		"vthunder_thunder_fw_vrid":                                                        42,
		"vthunder_thunder_glm":                                                            43,
		"vthunder_thunder_glm_send":                                                       44,
		"vthunder_thunder_harmony_controller_profile":                                     45,
		"vthunder_thunder_hostname":                                                       46,
		"vthunder_thunder_import":                                                         47,
		"vthunder_thunder_interface_ethernet":                                             48,
		"vthunder_thunder_interface_ethernet_bfd":                                         49,
		"vthunder_thunder_interface_ethernet_ipv6":                                        50,
		"vthunder_thunder_interface_ethernet_lldp":                                        51,
		"vthunder_thunder_interface_ethernet_trunk_group":                                 52,
		"vthunder_thunder_interface_lif":                                                  53,
		"vthunder_thunder_interface_lif_ip":                                               54,
		"vthunder_thunder_interface_management":                                           55,
		"vthunder_thunder_interface_ve":                                                   56,
		"vthunder_thunder_interface_ve_bfd":                                               57,
		"vthunder_thunder_interface_ve_ip":                                                58,
		"vthunder_thunder_interface_ve_ipv6":                                              59,
		"vthunder_thunder_ip_address":                                                     60,
		"vthunder_thunder_ip_dns_primary":                                                 61,
		"vthunder_thunder_ip_dns_secondary":                                               62,
		"vthunder_thunder_ip_dns_suffix":                                                  63,
		"vthunder_thunder_ip_frag":                                                        64,
		"vthunder_thunder_ip_icmp":                                                        65,
		"vthunder_thunder_ip_nat_alg_pptp":                                                66,
		"vthunder_thunder_ip_nat_global":                                                  67,
		"vthunder_thunder_ip_nat_icmp":                                                    68,
		"vthunder_thunder_ip_nat_pool":                                                    69,
		"vthunder_thunder_ip_nat_pool_group":                                              70,
		"vthunder_thunder_ip_prefix_list":                                                 71,
		"vthunder_thunder_ip_reroute":                                                     72,
		"vthunder_thunder_ip_route_static_bfd":                                            73,
		"vthunder_thunder_ip_tcp":                                                         74,
		"vthunder_thunder_ipv6_frag":                                                      75,
		"vthunder_thunder_ipv6_icmpv6":                                                    76,
		"vthunder_thunder_ipv6_nat_icmpv6":                                                77,
		"vthunder_thunder_ntp_auth_key":                                                   78,
		"vthunder_thunder_ntp_server_hostname":                                            79,
		"vthunder_thunder_ntp_trusted_key":                                                80,
		"vthunder_thunder_overlay_tunnel_options":                                         81,
		"vthunder_thunder_overlay_tunnel_vtep":                                            82,
		"vthunder_thunder_partition":                                                      83,
		"vthunder_thunder_reboot":                                                         84,
		"vthunder_thunder_rib_route":                                                      85,
		"vthunder_thunder_route_map":                                                      86,
		"vthunder_thunder_router_bgp":                                                     87,
		"vthunder_thunder_router_bgp_address_family_ipv6":                                 88,
		"vthunder_thunder_router_bgp_address_family_ipv6_neighbor_ethernet_neighbor_ipv6": 89,
		"vthunder_thunder_router_bgp_address_family_ipv6_neighbor_ipv4_neighbor":          90,
		"vthunder_thunder_router_bgp_address_family_ipv6_neighbor_ipv6_neighbor":          91,
		"vthunder_thunder_router_bgp_address_family_ipv6_neighbor_peer_group_neighbor":    92,
		"vthunder_thunder_router_bgp_address_family_ipv6_neighbor_trunk_neighbor_ipv6":    93,
		"vthunder_thunder_router_bgp_address_family_ipv6_neighbor_ve_neighbor_ipv6":       94,
		"vthunder_thunder_router_bgp_address_family_ipv6_network_ipv6_network":            95,
		"vthunder_thunder_router_bgp_address_family_ipv6_network_synchronization":         96,
		"vthunder_thunder_router_bgp_address_family_ipv6_redistribute":                    97,
		"vthunder_thunder_router_bgp_neighbor_ethernet_neighbor":                          98,
		"vthunder_thunder_router_bgp_neighbor_ipv4_neighbor":                              99,
		"vthunder_thunder_router_bgp_neighbor_ipv6_neighbor":                              100,
		"vthunder_thunder_router_bgp_neighbor_peer_group_neighbor":                        101,
		"vthunder_thunder_router_bgp_neighbor_trunk_neighbor":                             102,
		"vthunder_thunder_router_bgp_neighbor_ve_neighbor":                                103,
		"vthunder_thunder_router_bgp_network_ip_cidr":                                     104,
		"vthunder_thunder_router_bgp_network_synchronization":                             105,
		"vthunder_thunder_router_bgp_redistribute":                                        106,
		"vthunder_thunder_router_ospf":                                                    107,
		"vthunder_thunder_router_ospf_area":                                               108,
		"vthunder_thunder_router_ospf_default_information":                                109,
		"vthunder_thunder_router_ospf_redistribute":                                       110,
		"vthunder_thunder_server":                                                         111,
		"vthunder_thunder_service_group":                                                  112,
		"vthunder_thunder_slb_aflow":                                                      113,
		"vthunder_thunder_slb_common":                                                     114,
		"vthunder_thunder_slb_common_conn_rate_limit_src_ip":                              115,
		"vthunder_thunder_slb_connection_reuse":                                           116,
		"vthunder_thunder_slb_crl_srcip":                                                  117,
		"vthunder_thunder_slb_dns":                                                        118,
		"vthunder_thunder_slb_dns_cache":                                                  119,
		"vthunder_thunder_slb_dns_response_rate_limiting":                                 120,
		"vthunder_thunder_slb_fast_http_proxy":                                            121,
		"vthunder_thunder_slb_fix":                                                        122,
		"vthunder_thunder_slb_ftp_ctl":                                                    123,
		"vthunder_thunder_slb_ftp_data":                                                   124,
		"vthunder_thunder_slb_ftp_proxy":                                                  125,
		"vthunder_thunder_slb_generic_proxy":                                              126,
		"vthunder_thunder_slb_health_gateway":                                             127,
		"vthunder_thunder_slb_health_stat":                                                128,
		"vthunder_thunder_slb_http2":                                                      129,
		"vthunder_thunder_slb_http_proxy":                                                 130,
		"vthunder_thunder_slb_hw_compress":                                                131,
		"vthunder_thunder_slb_icap":                                                       132,
		"vthunder_thunder_slb_icap_http":                                                  133,
		"vthunder_thunder_slb_imapproxy":                                                  134,
		"vthunder_thunder_slb_l4":                                                         135,
		"vthunder_thunder_slb_l7session":                                                  136,
		"vthunder_thunder_slb_mlb":                                                        137,
		"vthunder_thunder_slb_mssql":                                                      138,
		"vthunder_thunder_slb_mysql":                                                      139,
		"vthunder_thunder_slb_passthrough":                                                140,
		"vthunder_thunder_slb_perf":                                                       141,
		"vthunder_thunder_slb_persist":                                                    142,
		"vthunder_thunder_slb_player_id_global":                                           143,
		"vthunder_thunder_slb_pop3_proxy":                                                 144,
		"vthunder_thunder_slb_proxy":                                                      145,
		"vthunder_thunder_slb_rate_limit_log":                                             146,
		"vthunder_thunder_slb_rc_cache_global":                                            147,
		"vthunder_thunder_slb_resource_usage":                                             148,
		"vthunder_thunder_slb_server_port":                                                149,
		"vthunder_thunder_slb_sip":                                                        150,
		"vthunder_thunder_slb_smpp":                                                       151,
		"vthunder_thunder_slb_smtp":                                                       152,
		"vthunder_thunder_slb_spdy_proxy":                                                 153,
		"vthunder_thunder_slb_sport_rate_limit":                                           154,
		"vthunder_thunder_slb_ssl_cert_revoke":                                            155,
		"vthunder_thunder_slb_ssl_expire_check":                                           156,
		"vthunder_thunder_slb_ssl_forward_proxy":                                          157,
		"vthunder_thunder_slb_svm_source_nat":                                             158,
		"vthunder_thunder_slb_switch":                                                     159,
		"vthunder_thunder_slb_template_cache":                                             160,
		"vthunder_thunder_slb_template_cipher":                                            161,
		"vthunder_thunder_slb_template_client_ssh":                                        162,
		"vthunder_thunder_slb_template_client_ssl":                                        163,
		"vthunder_thunder_slb_template_connection_reuse":                                  164,
		"vthunder_thunder_slb_template_csv":                                               165,
		"vthunder_thunder_slb_template_dblb":                                              166,
		"vthunder_thunder_slb_template_diameter":                                          167,
		"vthunder_thunder_slb_template_dns":                                               168,
		"vthunder_thunder_slb_template_dynamic_service":                                   169,
		"vthunder_thunder_slb_template_external_service":                                  170,
		"vthunder_thunder_slb_template_fix":                                               171,
		"vthunder_thunder_slb_template_ftp":                                               172,
		"vthunder_thunder_slb_template_http":                                              173,
		"vthunder_thunder_slb_template_http_policy":                                       174,
		"vthunder_thunder_slb_template_imap_pop3":                                         175,
		"vthunder_thunder_slb_template_logging":                                           176,
		"vthunder_thunder_slb_template_monitor":                                           177,
		"vthunder_thunder_slb_template_mqtt":                                              178,
		"vthunder_thunder_slb_template_persist_cookie":                                    179,
		"vthunder_thunder_slb_template_persist_source_ip":                                 180,
		"vthunder_thunder_slb_template_policy":                                            181,
		"vthunder_thunder_slb_template_port":                                              182,
		"vthunder_thunder_slb_template_reqmod_icap":                                       183,
		"vthunder_thunder_slb_template_respmod_icap":                                      184,
		"vthunder_thunder_slb_template_server":                                            185,
		"vthunder_thunder_slb_template_server_ssh":                                        186,
		"vthunder_thunder_slb_template_server_ssl":                                        187,
		"vthunder_thunder_slb_template_sip":                                               188,
		"vthunder_thunder_slb_template_smpp":                                              189,
		"vthunder_thunder_slb_template_smtp":                                              190,
		"vthunder_thunder_slb_template_snmp":                                              191,
		"vthunder_thunder_slb_template_ssli":                                              192,
		"vthunder_thunder_slb_template_tcp":                                               193,
		"vthunder_thunder_slb_template_tcp_proxy":                                         194,
		"vthunder_thunder_slb_template_udp":                                               195,
		"vthunder_thunder_slb_template_virtual_port":                                      196,
		"vthunder_thunder_slb_template_virtual_server":                                    197,
		"vthunder_thunder_slb_transparent_acl_template":                                   198,
		"vthunder_thunder_slb_transparent_tcp_template":                                   199,
		"vthunder_thunder_slb_virtual_server_port":                                        200,
		"vthunder_thunder_snmp_server_SNMPv1_v2c_user":                                    201,
		"vthunder_thunder_snmp_server_SNMPv1_v2c_user_oid":                                202,
		"vthunder_thunder_snmp_server_SNMPv3_user":                                        203,
		"vthunder_thunder_snmp_server_community_read":                                     204,
		"vthunder_thunder_snmp_server_community_read_oid":                                 205,
		"vthunder_thunder_snmp_server_contact":                                            206,
		"vthunder_thunder_snmp_server_disable_traps":                                      207,
		"vthunder_thunder_snmp_server_enable_traps":                                       208,
		"vthunder_thunder_snmp_server_enable_traps_gslb":                                  209,
		"vthunder_thunder_snmp_server_enable_traps_lsn":                                   210,
		"vthunder_thunder_snmp_server_enable_traps_network":                               211,
		"vthunder_thunder_snmp_server_enable_traps_routing_bgp":                           212,
		"vthunder_thunder_snmp_server_enable_traps_routing_isis":                          213,
		"vthunder_thunder_snmp_server_enable_traps_routing_ospf":                          214,
		"vthunder_thunder_snmp_server_enable_traps_slb":                                   215,
		"vthunder_thunder_snmp_server_enable_traps_slb_change":                            216,
		"vthunder_thunder_snmp_server_enable_traps_snmp":                                  217,
		"vthunder_thunder_snmp_server_enable_traps_ssl":                                   218,
		"vthunder_thunder_snmp_server_enable_traps_system":                                219,
		"vthunder_thunder_snmp_server_enable_traps_vcs":                                   220,
		"vthunder_thunder_snmp_server_enable_traps_vrrp_a":                                221,
		"vthunder_thunder_snmp_server_engineID":                                           222,
		"vthunder_thunder_snmp_server_group":                                              223,
		"vthunder_thunder_snmp_server_host_host_name":                                     224,
		"vthunder_thunder_snmp_server_host_ipv4_host":                                     225,
		"vthunder_thunder_snmp_server_host_ipv6_host":                                     226,
		"vthunder_thunder_snmp_server_location":                                           227,
		"vthunder_thunder_snmp_server_management_index":                                   228,
		"vthunder_thunder_snmp_server_slb_data_cache_timeout":                             229,
		"vthunder_thunder_snmp_server_user":                                               230,
		"vthunder_thunder_snmp_server_view":                                               231,
		"vthunder_thunder_system":                                                         232,
		"vthunder_thunder_system_ve_mac_scheme":                                           233,
		"vthunder_thunder_virtual_server":                                                 234,
		"vthunder_thunder_vrrp_common":                                                    235,
		"vthunder_thunder_vrrp_peer_group":                                                236,
		"vthunder_thunder_vrrp_session_sync":                                              237,
		"vthunder_thunder_vrrp_vrid":                                                      238,
		"vthunder_thunder_web_category":                                                   239,
		"vthunder_thunder_web_category_category_list":                                     240,
		"vthunder_thunder_web_category_proxy_server":                                      241,
		"vthunder_thunder_web_category_reputation_scope":                                  242,
		"vthunder_thunder_web_category_statistics":                                        243,
		"vthunder_thunder_web_category_web_reputation":                                    244,
		"vthunder_thunder_write_memory":                                                   245,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
