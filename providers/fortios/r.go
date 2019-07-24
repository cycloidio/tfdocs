package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_address",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure firewall addresses used in firewall policies of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"object",
				"address",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Address name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of address(Support ipmask, iprange, fqdn and geography).`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IP address and subnet mask of address.`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `First IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `Final IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name address.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `IP addresses associated to a specific country.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the address item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Address name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of address(Support ipmask, iprange, fqdn and geography).`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IP address and subnet mask of address.`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `First IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `Final IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name address.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `IP addresses associated to a specific country.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the address item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Address name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of address(Support ipmask, iprange, fqdn and geography).`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IP address and subnet mask of address.`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `First IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `Final IP address (inclusive) in the range for the address.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name address.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `IP addresses associated to a specific country.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_addressgroup",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure firewall address group used in firewall policies of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"object",
				"addressgroup",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Address group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) Address objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall address group item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Address group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Address objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall address group item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Address group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Address objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_ippool",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure IPv4 IP address pools of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"object",
				"ippool",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) IP pool name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) IP pool type(Support overload and one-to-one).`,
				},
				resource.Attribute{
					Name:        "startip",
					Description: `(Required) First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "endip",
					Description: `(Required) Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable/disable replying to ARP requests when an IP Pool is added to a policy.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IP pool item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IP pool name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `IP pool type(Support overload and one-to-one).`,
				},
				resource.Attribute{
					Name:        "startip",
					Description: `First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "endip",
					Description: `Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable/disable replying to ARP requests when an IP Pool is added to a policy.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IP pool item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IP pool name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `IP pool type(Support overload and one-to-one).`,
				},
				resource.Attribute{
					Name:        "startip",
					Description: `First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "endip",
					Description: `Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx).`,
				},
				resource.Attribute{
					Name:        "arp_reply",
					Description: `Enable/disable replying to ARP requests when an IP Pool is added to a policy.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_service",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure firewall service of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"object",
				"service",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) Service category.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol type based on IANA numbers.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "iprange",
					Description: `Start and end of the IP range associated with service.`,
				},
				resource.Attribute{
					Name:        "tcp_portrange",
					Description: `Multiple TCP port ranges.`,
				},
				resource.Attribute{
					Name:        "udp_portrange",
					Description: `Multiple UDP port ranges.`,
				},
				resource.Attribute{
					Name:        "sctp_portrange",
					Description: `Multiple SCTP port ranges.`,
				},
				resource.Attribute{
					Name:        "icmptype",
					Description: `ICMP type.`,
				},
				resource.Attribute{
					Name:        "icmpcode",
					Description: `ICMP code.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `IP protocol number.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Service category.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol type based on IANA numbers.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "iprange",
					Description: `Start and end of the IP range associated with service.`,
				},
				resource.Attribute{
					Name:        "tcp_portrange",
					Description: `Multiple TCP port ranges.`,
				},
				resource.Attribute{
					Name:        "udp_portrange",
					Description: `Multiple UDP port ranges.`,
				},
				resource.Attribute{
					Name:        "sctp_portrange",
					Description: `Multiple SCTP port ranges.`,
				},
				resource.Attribute{
					Name:        "icmptype",
					Description: `ICMP type.`,
				},
				resource.Attribute{
					Name:        "icmpcode",
					Description: `ICMP code.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `IP protocol number.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Service category.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol type based on IANA numbers.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "iprange",
					Description: `Start and end of the IP range associated with service.`,
				},
				resource.Attribute{
					Name:        "tcp_portrange",
					Description: `Multiple TCP port ranges.`,
				},
				resource.Attribute{
					Name:        "udp_portrange",
					Description: `Multiple UDP port ranges.`,
				},
				resource.Attribute{
					Name:        "sctp_portrange",
					Description: `Multiple SCTP port ranges.`,
				},
				resource.Attribute{
					Name:        "icmptype",
					Description: `ICMP type.`,
				},
				resource.Attribute{
					Name:        "icmpcode",
					Description: `ICMP code.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `IP protocol number.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_servicegroup",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure firewall service group of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"object",
				"servicegroup",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Service group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) Service objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service group item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Service group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Service objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall service group item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Service group name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Service objects contained within the group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_vip",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure firewall virtual IPs (VIPs) of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"object",
				"vip",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Virtual IP name.`,
				},
				resource.Attribute{
					Name:        "extip",
					Description: `(Required) IP address or address range on the external interface that you want to map to an address or address range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedip",
					Description: `(Required) IP address or address range on the destination network to which the external IP address is mapped.`,
				},
				resource.Attribute{
					Name:        "extintf",
					Description: `Interface connected to the source network that receives the packets that will be forwarded to the destination network.`,
				},
				resource.Attribute{
					Name:        "portforward",
					Description: `Enable/disable port forwarding.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol to use when forwarding packets.`,
				},
				resource.Attribute{
					Name:        "extport",
					Description: `Incoming port number range that you want to map to a port number range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedport",
					Description: `Port number range on the destination network to which the external port number range is mapped.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall virtual IPs item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Virtual IP name.`,
				},
				resource.Attribute{
					Name:        "extip",
					Description: `IP address or address range on the external interface that you want to map to an address or address range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedip",
					Description: `IP address or address range on the destination network to which the external IP address is mapped.`,
				},
				resource.Attribute{
					Name:        "extintf",
					Description: `Interface connected to the source network that receives the packets that will be forwarded to the destination network.`,
				},
				resource.Attribute{
					Name:        "portforward",
					Description: `Enable/disable port forwarding.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol to use when forwarding packets.`,
				},
				resource.Attribute{
					Name:        "extport",
					Description: `Incoming port number range that you want to map to a port number range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedport",
					Description: `Port number range on the destination network to which the external port number range is mapped.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall virtual IPs item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Virtual IP name.`,
				},
				resource.Attribute{
					Name:        "extip",
					Description: `IP address or address range on the external interface that you want to map to an address or address range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedip",
					Description: `IP address or address range on the destination network to which the external IP address is mapped.`,
				},
				resource.Attribute{
					Name:        "extintf",
					Description: `Interface connected to the source network that receives the packets that will be forwarded to the destination network.`,
				},
				resource.Attribute{
					Name:        "portforward",
					Description: `Enable/disable port forwarding.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol to use when forwarding packets.`,
				},
				resource.Attribute{
					Name:        "extport",
					Description: `Incoming port number range that you want to map to a port number range on the destination network.`,
				},
				resource.Attribute{
					Name:        "mappedport",
					Description: `Port number range on the destination network to which the external port number range is mapped.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_object_vipgroup",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure virtual IP groups of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"object",
				"vipgroup",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) VIP group name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) Member VIP objects of the group.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual IP groups item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VIP group name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Member VIP objects of the group.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the virtual IP groups item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VIP group name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Member VIP objects of the group.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_firewall_security_policy",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure firewall policies of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"security",
				"policy",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Policy name.`,
				},
				resource.Attribute{
					Name:        "srcintf",
					Description: `(Required) Incoming (ingress) interface.`,
				},
				resource.Attribute{
					Name:        "dstintf",
					Description: `(Required) Outgoing (egress) interface.`,
				},
				resource.Attribute{
					Name:        "srcaddr",
					Description: `(Required) Source address and address group names.`,
				},
				resource.Attribute{
					Name:        "dstaddr",
					Description: `(Required) Destination address and address group names.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Policy action.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Required) Schedule name.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) Service and service group names..`,
				},
				resource.Attribute{
					Name:        "utm_status",
					Description: `Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.`,
				},
				resource.Attribute{
					Name:        "logtraffic",
					Description: `Enable or disable logging. Log all sessions or security profile sessions.`,
				},
				resource.Attribute{
					Name:        "logtraffic_start",
					Description: `Record logs when a session starts and ends.`,
				},
				resource.Attribute{
					Name:        "capture_packet",
					Description: `Enable/disable capture packets.`,
				},
				resource.Attribute{
					Name:        "ippool",
					Description: `Enable to use IP Pools for source NAT.`,
				},
				resource.Attribute{
					Name:        "poolname",
					Description: `IP Pool names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Names of user groups that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `Device type category.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "av_profile",
					Description: `Name of an existing Antivirus profile.`,
				},
				resource.Attribute{
					Name:        "webfilter_profile",
					Description: `Name of an existing Web filter profile.`,
				},
				resource.Attribute{
					Name:        "dnsfilter_profile",
					Description: `Name of an existing DNS filter profile.`,
				},
				resource.Attribute{
					Name:        "ips_sensor",
					Description: `Name of an existing IPS sensor.`,
				},
				resource.Attribute{
					Name:        "application_list",
					Description: `Name of an existing Application list.`,
				},
				resource.Attribute{
					Name:        "ssl_ssh_profile",
					Description: `Name of an existing SSL SSH profile.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Enable/disable source NAT.`,
				},
				resource.Attribute{
					Name:        "internet_service_src",
					Description: `Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_src_id",
					Description: `Internet Service source ID.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Names of individual users that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable or disable this policy.`,
				},
				resource.Attribute{
					Name:        "profile_protocol_options",
					Description: `Name of an existing Protocol options profile. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall policy item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Policy name.`,
				},
				resource.Attribute{
					Name:        "srcintf",
					Description: `Incoming (ingress) interface.`,
				},
				resource.Attribute{
					Name:        "dstintf",
					Description: `Outgoing (egress) interface.`,
				},
				resource.Attribute{
					Name:        "srcaddr",
					Description: `Source address and address group names.`,
				},
				resource.Attribute{
					Name:        "dstaddr",
					Description: `Destination address and address group names.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Policy action.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `Schedule name.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service and service group names..`,
				},
				resource.Attribute{
					Name:        "utm_status",
					Description: `Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.`,
				},
				resource.Attribute{
					Name:        "logtraffic",
					Description: `Enable or disable logging. Log all sessions or security profile sessions.`,
				},
				resource.Attribute{
					Name:        "logtraffic_start",
					Description: `Record logs when a session starts and ends.`,
				},
				resource.Attribute{
					Name:        "capture_packet",
					Description: `Enable/disable capture packets.`,
				},
				resource.Attribute{
					Name:        "ippool",
					Description: `Enable to use IP Pools for source NAT.`,
				},
				resource.Attribute{
					Name:        "poolname",
					Description: `IP Pool names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Names of user groups that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `Device type category.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "av_profile",
					Description: `Name of an existing Antivirus profile.`,
				},
				resource.Attribute{
					Name:        "webfilter_profile",
					Description: `Name of an existing Web filter profile.`,
				},
				resource.Attribute{
					Name:        "dnsfilter_profile",
					Description: `Name of an existing DNS filter profile.`,
				},
				resource.Attribute{
					Name:        "ips_sensor",
					Description: `Name of an existing IPS sensor.`,
				},
				resource.Attribute{
					Name:        "application_list",
					Description: `Name of an existing Application list.`,
				},
				resource.Attribute{
					Name:        "ssl_ssh_profile",
					Description: `Name of an existing SSL SSH profile.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Enable/disable source NAT.`,
				},
				resource.Attribute{
					Name:        "internet_service_src",
					Description: `Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_src_id",
					Description: `Internet Service source ID.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Names of individual users that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable or disable this policy.`,
				},
				resource.Attribute{
					Name:        "profile_protocol_options",
					Description: `Name of an existing Protocol options profile.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the firewall policy item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Policy name.`,
				},
				resource.Attribute{
					Name:        "srcintf",
					Description: `Incoming (ingress) interface.`,
				},
				resource.Attribute{
					Name:        "dstintf",
					Description: `Outgoing (egress) interface.`,
				},
				resource.Attribute{
					Name:        "srcaddr",
					Description: `Source address and address group names.`,
				},
				resource.Attribute{
					Name:        "dstaddr",
					Description: `Destination address and address group names.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_id",
					Description: `Internet Service ID.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Policy action.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `Schedule name.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service and service group names..`,
				},
				resource.Attribute{
					Name:        "utm_status",
					Description: `Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.`,
				},
				resource.Attribute{
					Name:        "logtraffic",
					Description: `Enable or disable logging. Log all sessions or security profile sessions.`,
				},
				resource.Attribute{
					Name:        "logtraffic_start",
					Description: `Record logs when a session starts and ends.`,
				},
				resource.Attribute{
					Name:        "capture_packet",
					Description: `Enable/disable capture packets.`,
				},
				resource.Attribute{
					Name:        "ippool",
					Description: `Enable to use IP Pools for source NAT.`,
				},
				resource.Attribute{
					Name:        "poolname",
					Description: `IP Pool names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Names of user groups that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `Device type category.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "av_profile",
					Description: `Name of an existing Antivirus profile.`,
				},
				resource.Attribute{
					Name:        "webfilter_profile",
					Description: `Name of an existing Web filter profile.`,
				},
				resource.Attribute{
					Name:        "dnsfilter_profile",
					Description: `Name of an existing DNS filter profile.`,
				},
				resource.Attribute{
					Name:        "ips_sensor",
					Description: `Name of an existing IPS sensor.`,
				},
				resource.Attribute{
					Name:        "application_list",
					Description: `Name of an existing Application list.`,
				},
				resource.Attribute{
					Name:        "ssl_ssh_profile",
					Description: `Name of an existing SSL SSH profile.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Enable/disable source NAT.`,
				},
				resource.Attribute{
					Name:        "internet_service_src",
					Description: `Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.`,
				},
				resource.Attribute{
					Name:        "internet_service_src_id",
					Description: `Internet Service source ID.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Names of individual users that can authenticate with this policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable or disable this policy.`,
				},
				resource.Attribute{
					Name:        "profile_protocol_options",
					Description: `Name of an existing Protocol options profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_log_fortianalyzer_setting",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure configure logging to FortiAnalyzer log management devices.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"fortianalyzer",
				"setting",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Enable/disable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The remote FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "upload_option",
					Description: `Enable/disable logging to hard disk and then uploading to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "hmac_algorithm",
					Description: `FortiAnalyzer IPsec tunnel HMAC algorithm.`,
				},
				resource.Attribute{
					Name:        "enc_algorithm",
					Description: `Enable/disable sending FortiAnalyzer log data with SSL encryption. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The remote FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "upload_option",
					Description: `Enable/disable logging to hard disk and then uploading to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "hmac_algorithm",
					Description: `FortiAnalyzer IPsec tunnel HMAC algorithm.`,
				},
				resource.Attribute{
					Name:        "enc_algorithm",
					Description: `Enable/disable sending FortiAnalyzer log data with SSL encryption.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The remote FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "upload_option",
					Description: `Enable/disable logging to hard disk and then uploading to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable logging to FortiAnalyzer.`,
				},
				resource.Attribute{
					Name:        "hmac_algorithm",
					Description: `FortiAnalyzer IPsec tunnel HMAC algorithm.`,
				},
				resource.Attribute{
					Name:        "enc_algorithm",
					Description: `Enable/disable sending FortiAnalyzer log data with SSL encryption.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_log_syslog_setting",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure logging to remote Syslog logging servers.`,
			Description:      ``,
			Keywords: []string{
				"log",
				"syslog",
				"setting",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Enable/disable remote syslog logging.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Address of remote syslog server.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Remote syslog logging over UDP/Reliable TCP.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Server listen port.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Remote syslog facility.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IP address of syslog.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Log format. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable remote syslog logging.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Address of remote syslog server.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Remote syslog logging over UDP/Reliable TCP.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Server listen port.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Remote syslog facility.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IP address of syslog.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Log format.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable remote syslog logging.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Address of remote syslog server.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Remote syslog logging over UDP/Reliable TCP.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Server listen port.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Remote syslog facility.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IP address of syslog.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `Log format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_networking_interface_port",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure interface settings of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"interface",
				"port",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) If the interface is physical, the argument is the name of the interface.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Interface type (support physical, vlan, loopback).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Interface IPv4 address and subnet mask, syntax` + "`" + ` - X.X.X.X/24.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `Alias will be displayed with the interface name to make it easier to distinguish.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Bring the interface up or shut the interface down.`,
				},
				resource.Attribute{
					Name:        "device_identification",
					Description: `Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.`,
				},
				resource.Attribute{
					Name:        "tcp_mss",
					Description: `TCP maximum segment size. 0 means do not change segment size.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Interface speed. The default setting and the options available depend on the interface hardware.`,
				},
				resource.Attribute{
					Name:        "mtu_override",
					Description: `Enable to set a custom MTU for this interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `MTU value for this interface.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Interface role.`,
				},
				resource.Attribute{
					Name:        "allowaccess",
					Description: `Permitted types of management access to this interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) Addressing mode.`,
				},
				resource.Attribute{
					Name:        "dns_server_override",
					Description: `Enable/disable use DNS acquired by DHCP or PPPoE.`,
				},
				resource.Attribute{
					Name:        "defaultgw",
					Description: `Enable to get the gateway IP from the DHCP or PPPoE server.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Interface is in this virtual domain (VDOM).`,
				},
				resource.Attribute{
					Name:        "vlanid",
					Description: `VLAN ID. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Name of the interface.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Interface IPv4 address and subnet mask, syntax` + "`" + ` - X.X.X.X/24.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `Alias will be displayed with the interface name to make it easier to distinguish.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Bring the interface up or shut the interface down.`,
				},
				resource.Attribute{
					Name:        "device_identification",
					Description: `Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.`,
				},
				resource.Attribute{
					Name:        "tcp_mss",
					Description: `TCP maximum segment size. 0 means do not change segment size.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Interface speed. The default setting and the options available depend on the interface hardware.`,
				},
				resource.Attribute{
					Name:        "mtu_override",
					Description: `Enable to set a custom MTU for this interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `MTU value for this interface.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Interface role.`,
				},
				resource.Attribute{
					Name:        "allowaccess",
					Description: `Permitted types of management access to this interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Addressing mode.`,
				},
				resource.Attribute{
					Name:        "dns_server_override",
					Description: `Enable/disable use DNS acquired by DHCP or PPPoE.`,
				},
				resource.Attribute{
					Name:        "defaultgw",
					Description: `Enable to get the gateway IP from the DHCP or PPPoE server.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Interface type (support physical, vlan, loopback).`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Interface is in this virtual domain (VDOM).`,
				},
				resource.Attribute{
					Name:        "vlanid",
					Description: `VLAN ID.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The Name of the interface.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Interface IPv4 address and subnet mask, syntax` + "`" + ` - X.X.X.X/24.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `Alias will be displayed with the interface name to make it easier to distinguish.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Bring the interface up or shut the interface down.`,
				},
				resource.Attribute{
					Name:        "device_identification",
					Description: `Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.`,
				},
				resource.Attribute{
					Name:        "tcp_mss",
					Description: `TCP maximum segment size. 0 means do not change segment size.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Interface speed. The default setting and the options available depend on the interface hardware.`,
				},
				resource.Attribute{
					Name:        "mtu_override",
					Description: `Enable to set a custom MTU for this interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `MTU value for this interface.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Interface role.`,
				},
				resource.Attribute{
					Name:        "allowaccess",
					Description: `Permitted types of management access to this interface.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Addressing mode.`,
				},
				resource.Attribute{
					Name:        "dns_server_override",
					Description: `Enable/disable use DNS acquired by DHCP or PPPoE.`,
				},
				resource.Attribute{
					Name:        "defaultgw",
					Description: `Enable to get the gateway IP from the DHCP or PPPoE server.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Interface type (support physical, vlan, loopback).`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Interface is in this virtual domain (VDOM).`,
				},
				resource.Attribute{
					Name:        "vlanid",
					Description: `VLAN ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_networking_route_static",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure static route of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"route",
				"static",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "dst",
					Description: `(Required) Destination IP and mask for this route.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Required) Gateway IP for this route.`,
				},
				resource.Attribute{
					Name:        "blackhole",
					Description: `Enable/disable black hole.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Administrative distance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Administrative weight.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Administrative priority.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Required) Gateway out interface or tunnel.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Optional comments. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the static route item.`,
				},
				resource.Attribute{
					Name:        "dst",
					Description: `Destination IP and mask for this route.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Gateway IP for this route.`,
				},
				resource.Attribute{
					Name:        "blackhole",
					Description: `Enable/disable black hole.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Administrative distance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Administrative weight.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Administrative priority.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Gateway out interface or tunnel.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Optional comments.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the static route item.`,
				},
				resource.Attribute{
					Name:        "dst",
					Description: `Destination IP and mask for this route.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Gateway IP for this route.`,
				},
				resource.Attribute{
					Name:        "blackhole",
					Description: `Enable/disable black hole.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `Administrative distance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Administrative weight.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Administrative priority.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Gateway out interface or tunnel.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Optional comments.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_admin_administrator",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure administrator accounts of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"admin",
				"administrator",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Admin user password.`,
				},
				resource.Attribute{
					Name:        "trusthostN",
					Description: `Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domain(s) that the administrator can access.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `Access profile for this administrator. Access profiles control administrator access to FortiGate features.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the administrator account item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Admin user password.`,
				},
				resource.Attribute{
					Name:        "trusthostN",
					Description: `Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domain(s) that the administrator can access.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `Access profile for this administrator. Access profiles control administrator access to FortiGate features.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the administrator account item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Admin user password.`,
				},
				resource.Attribute{
					Name:        "trusthostN",
					Description: `Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domain(s) that the administrator can access.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `Access profile for this administrator. Access profiles control administrator access to FortiGate features.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_admin_profiles",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure access profiles of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"admin",
				"profiles",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Profile name.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `Scope of admin access.`,
				},
				resource.Attribute{
					Name:        "secfabgrp",
					Description: `Security Fabric.`,
				},
				resource.Attribute{
					Name:        "ftviewgrp",
					Description: `FortiView.`,
				},
				resource.Attribute{
					Name:        "authgrp",
					Description: `Administrator access to Users and Devices.`,
				},
				resource.Attribute{
					Name:        "sysgrp",
					Description: `System Configuration.`,
				},
				resource.Attribute{
					Name:        "netgrp",
					Description: `Network Configuration.`,
				},
				resource.Attribute{
					Name:        "loggrp",
					Description: `Administrator access to Logging and Reporting including viewing log messages.`,
				},
				resource.Attribute{
					Name:        "fwgrp",
					Description: `Administrator access to the Firewall configuration.`,
				},
				resource.Attribute{
					Name:        "vpngrp",
					Description: `Administrator access to IPsec, SSL, PPTP, and L2TP VPN.`,
				},
				resource.Attribute{
					Name:        "utmgrp",
					Description: `Administrator access to Security Profiles.`,
				},
				resource.Attribute{
					Name:        "wanoptgrp",
					Description: `Administrator access to WAN Opt & Cache.`,
				},
				resource.Attribute{
					Name:        "wifi",
					Description: `Administrator access to the WiFi controller and Switch controller.`,
				},
				resource.Attribute{
					Name:        "admintimeout_override",
					Description: `Enable/disable overriding the global administrator idle timeout.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the access profile item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Profile name.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `Scope of admin access.`,
				},
				resource.Attribute{
					Name:        "secfabgrp",
					Description: `Security Fabric.`,
				},
				resource.Attribute{
					Name:        "ftviewgrp",
					Description: `FortiView.`,
				},
				resource.Attribute{
					Name:        "authgrp",
					Description: `Administrator access to Users and Devices.`,
				},
				resource.Attribute{
					Name:        "sysgrp",
					Description: `System Configuration.`,
				},
				resource.Attribute{
					Name:        "netgrp",
					Description: `Network Configuration.`,
				},
				resource.Attribute{
					Name:        "loggrp",
					Description: `Administrator access to Logging and Reporting including viewing log messages.`,
				},
				resource.Attribute{
					Name:        "fwgrp",
					Description: `Administrator access to the Firewall configuration.`,
				},
				resource.Attribute{
					Name:        "vpngrp",
					Description: `Administrator access to IPsec, SSL, PPTP, and L2TP VPN.`,
				},
				resource.Attribute{
					Name:        "utmgrp",
					Description: `Administrator access to Security Profiles.`,
				},
				resource.Attribute{
					Name:        "wanoptgrp",
					Description: `Administrator access to WAN Opt & Cache.`,
				},
				resource.Attribute{
					Name:        "wifi",
					Description: `Administrator access to the WiFi controller and Switch controller.`,
				},
				resource.Attribute{
					Name:        "admintimeout_override",
					Description: `Enable/disable overriding the global administrator idle timeout.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the access profile item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Profile name.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `Scope of admin access.`,
				},
				resource.Attribute{
					Name:        "secfabgrp",
					Description: `Security Fabric.`,
				},
				resource.Attribute{
					Name:        "ftviewgrp",
					Description: `FortiView.`,
				},
				resource.Attribute{
					Name:        "authgrp",
					Description: `Administrator access to Users and Devices.`,
				},
				resource.Attribute{
					Name:        "sysgrp",
					Description: `System Configuration.`,
				},
				resource.Attribute{
					Name:        "netgrp",
					Description: `Network Configuration.`,
				},
				resource.Attribute{
					Name:        "loggrp",
					Description: `Administrator access to Logging and Reporting including viewing log messages.`,
				},
				resource.Attribute{
					Name:        "fwgrp",
					Description: `Administrator access to the Firewall configuration.`,
				},
				resource.Attribute{
					Name:        "vpngrp",
					Description: `Administrator access to IPsec, SSL, PPTP, and L2TP VPN.`,
				},
				resource.Attribute{
					Name:        "utmgrp",
					Description: `Administrator access to Security Profiles.`,
				},
				resource.Attribute{
					Name:        "wanoptgrp",
					Description: `Administrator access to WAN Opt & Cache.`,
				},
				resource.Attribute{
					Name:        "wifi",
					Description: `Administrator access to the WiFi controller and Switch controller.`,
				},
				resource.Attribute{
					Name:        "admintimeout_override",
					Description: `Enable/disable overriding the global administrator idle timeout.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_apiuser_setting",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure API users of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"apiuser",
				"setting",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User name.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `(Required) Admin user access profile.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `(Required) Virtual domains.`,
				},
				resource.Attribute{
					Name:        "trusthost-Type",
					Description: `(Required) Trusthost type.`,
				},
				resource.Attribute{
					Name:        "trusthost-ipv4_trusthost",
					Description: `(Required) IPv4 trusted host address.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `Admin user access profile.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domains.`,
				},
				resource.Attribute{
					Name:        "trusthost-Type",
					Description: `Trusthost type.`,
				},
				resource.Attribute{
					Name:        "trusthost-ipv4_trusthost",
					Description: `IPv4 trusted host address.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "accprofile",
					Description: `Admin user access profile.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual domains.`,
				},
				resource.Attribute{
					Name:        "trusthost-Type",
					Description: `Trusthost type.`,
				},
				resource.Attribute{
					Name:        "trusthost-ipv4_trusthost",
					Description: `IPv4 trusted host address.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_license_forticare",
			Category:         "Resources",
			ShortDescription: `Provides a resource to add a FortiCare license for FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"license",
				"forticare",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "registration_code",
					Description: `(Required) Registration code.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_license_vdom",
			Category:         "Resources",
			ShortDescription: `Provides a resource to add a VDOM license for FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"license",
				"vdom",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "license",
					Description: `(Required) Registration code.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_license_vm",
			Category:         "Resources",
			ShortDescription: `Provides a resource to update VM license using uploaded file for FortiOS. Reboots immediately if successful.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"license",
				"vm",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "file_content",
					Description: `(Required) The license file, it needs to be base64 encoded, must not contain whitespace or other invalid base64 characters, and must be included in HTTP body.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_setting_dns",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure DNS of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"setting",
				"dns",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "primary",
					Description: `Primary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `Secondary DNS server IP address. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Primary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `Secondary DNS server IP address.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "primary",
					Description: `Primary DNS server IP address.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `Secondary DNS server IP address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_setting_global",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure options related to the overall operation of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"setting",
				"global",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) FortiGate unit's hostname.`,
				},
				resource.Attribute{
					Name:        "admintimeout",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `Number corresponding to your time zone from 00 to 86.`,
				},
				resource.Attribute{
					Name:        "admin_sport",
					Description: `Administrative access port for HTTPS.`,
				},
				resource.Attribute{
					Name:        "admin_ssh_port",
					Description: `Administrative access port for SSH. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "admintimeout",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `Number corresponding to your time zone from 00 to 86.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `FortiGate unit's hostname.`,
				},
				resource.Attribute{
					Name:        "admin_sport",
					Description: `Administrative access port for HTTPS.`,
				},
				resource.Attribute{
					Name:        "admin_ssh_port",
					Description: `Administrative access port for SSH.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "admintimeout",
					Description: `Number of minutes before an idle administrator session time out.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `Number corresponding to your time zone from 00 to 86.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `FortiGate unit's hostname.`,
				},
				resource.Attribute{
					Name:        "admin_sport",
					Description: `Administrative access port for HTTPS.`,
				},
				resource.Attribute{
					Name:        "admin_ssh_port",
					Description: `Administrative access port for SSH.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_setting_ntp",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure Network Time Protocol (NTP) servers of FortiOS.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"setting",
				"ntp",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Use the FortiGuard NTP server or any other available NTP Server.`,
				},
				resource.Attribute{
					Name:        "ntpserver",
					Description: `Configure the FortiGate to connect to any available third-party NTP server.`,
				},
				resource.Attribute{
					Name:        "ntpsync",
					Description: `Enable/disable setting the FortiGate system time by synchronizing with an NTP Server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Use the FortiGuard NTP server or any other available NTP Server.`,
				},
				resource.Attribute{
					Name:        "ntpserver",
					Description: `Configure the FortiGate to connect to any available third-party NTP server.`,
				},
				resource.Attribute{
					Name:        "ntpsync",
					Description: `Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "type",
					Description: `Use the FortiGuard NTP server or any other available NTP Server.`,
				},
				resource.Attribute{
					Name:        "ntpserver",
					Description: `Configure the FortiGate to connect to any available third-party NTP server.`,
				},
				resource.Attribute{
					Name:        "ntpsync",
					Description: `Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_system_vdom_setting",
			Category:         "Resources",
			ShortDescription: `Provides a resource to configure VDOM of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"vdom",
				"setting",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) VDOM name.`,
				},
				resource.Attribute{
					Name:        "short_name",
					Description: `VDOM short name.`,
				},
				resource.Attribute{
					Name:        "temporary",
					Description: `Temporary. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VDOM.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VDOM name.`,
				},
				resource.Attribute{
					Name:        "short_name",
					Description: `VDOM short name.`,
				},
				resource.Attribute{
					Name:        "temporary",
					Description: `Temporary.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the VDOM.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VDOM name.`,
				},
				resource.Attribute{
					Name:        "short_name",
					Description: `VDOM short name.`,
				},
				resource.Attribute{
					Name:        "temporary",
					Description: `Temporary.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_vpn_ipsec_phase1interface",
			Category:         "Resources",
			ShortDescription: `Provides a resource to use phase1-interface to define a phase 1 definition for a route-based (interface mode) IPsec VPN tunnel that generates authentication and encryption keys automatically.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"ipsec",
				"phase1interface",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) IPsec remote gateway name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Remote gateway type.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) Local physical, aggregate, or VLAN outgoing interface.`,
				},
				resource.Attribute{
					Name:        "peertype",
					Description: `Accept this peer type.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase1 proposal.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "wizard_type",
					Description: `GUI VPN Wizard Type.`,
				},
				resource.Attribute{
					Name:        "remote_gw",
					Description: `(Required) IPv4 address of the remote gateway's external interface.`,
				},
				resource.Attribute{
					Name:        "psksecret",
					Description: `(Required) Pre-shared secret for PSK authentication.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `Names of signed personal certificates.`,
				},
				resource.Attribute{
					Name:        "peerid",
					Description: `Accept this peer identity.`,
				},
				resource.Attribute{
					Name:        "peer",
					Description: `Accept this peer certificate.`,
				},
				resource.Attribute{
					Name:        "peergrp",
					Description: `Accept this peer certificate group.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_include",
					Description: `IPv4 split-include subnets.`,
				},
				resource.Attribute{
					Name:        "split_include_service",
					Description: `Split-include services.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_exclude",
					Description: `IPv4 subnets that should not be sent over the IPsec tunnel.`,
				},
				resource.Attribute{
					Name:        "authmethod",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "authmethod_remote",
					Description: `Authentication method (remote side).`,
				},
				resource.Attribute{
					Name:        "mode_cfg",
					Description: `Enable/disable configuration method. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the phase1-interface item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IPsec remote gateway name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Remote gateway type.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Local physical, aggregate, or VLAN outgoing interface.`,
				},
				resource.Attribute{
					Name:        "peertype",
					Description: `Accept this peer type.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase1 proposal.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "wizard_type",
					Description: `GUI VPN Wizard Type.`,
				},
				resource.Attribute{
					Name:        "remote_gw",
					Description: `IPv4 address of the remote gateway's external interface.`,
				},
				resource.Attribute{
					Name:        "psksecret",
					Description: `Pre-shared secret for PSK authentication.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `Names of signed personal certificates.`,
				},
				resource.Attribute{
					Name:        "peerid",
					Description: `Accept this peer identity.`,
				},
				resource.Attribute{
					Name:        "peer",
					Description: `Accept this peer certificate.`,
				},
				resource.Attribute{
					Name:        "peergrp",
					Description: `Accept this peer certificate group.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_include",
					Description: `IPv4 split-include subnets.`,
				},
				resource.Attribute{
					Name:        "split_include_service",
					Description: `Split-include services.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_exclude",
					Description: `IPv4 subnets that should not be sent over the IPsec tunnel.`,
				},
				resource.Attribute{
					Name:        "authmethod",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "authmethod_remote",
					Description: `Authentication method (remote side).`,
				},
				resource.Attribute{
					Name:        "mode_cfg",
					Description: `Enable/disable configuration method.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the phase1-interface item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IPsec remote gateway name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Remote gateway type.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Local physical, aggregate, or VLAN outgoing interface.`,
				},
				resource.Attribute{
					Name:        "peertype",
					Description: `Accept this peer type.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase1 proposal.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "wizard_type",
					Description: `GUI VPN Wizard Type.`,
				},
				resource.Attribute{
					Name:        "remote_gw",
					Description: `IPv4 address of the remote gateway's external interface.`,
				},
				resource.Attribute{
					Name:        "psksecret",
					Description: `Pre-shared secret for PSK authentication.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `Names of signed personal certificates.`,
				},
				resource.Attribute{
					Name:        "peerid",
					Description: `Accept this peer identity.`,
				},
				resource.Attribute{
					Name:        "peer",
					Description: `Accept this peer certificate.`,
				},
				resource.Attribute{
					Name:        "peergrp",
					Description: `Accept this peer certificate group.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_include",
					Description: `IPv4 split-include subnets.`,
				},
				resource.Attribute{
					Name:        "split_include_service",
					Description: `Split-include services.`,
				},
				resource.Attribute{
					Name:        "ipv4_split_exclude",
					Description: `IPv4 subnets that should not be sent over the IPsec tunnel.`,
				},
				resource.Attribute{
					Name:        "authmethod",
					Description: `Authentication method.`,
				},
				resource.Attribute{
					Name:        "authmethod_remote",
					Description: `Authentication method (remote side).`,
				},
				resource.Attribute{
					Name:        "mode_cfg",
					Description: `Enable/disable configuration method.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortios_vpn_ipsec_phase2interface",
			Category:         "Resources",
			ShortDescription: `Provides a resource to use phase2-interface to add or edit a phase 2 configuration on a route-based (interface mode) IPsec tunnel.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"ipsec",
				"phase2interface",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) IPsec tunnel name.`,
				},
				resource.Attribute{
					Name:        "phase1name",
					Description: `(Required) Phase 1 determines the options required for phase 2.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase2 proposal.`,
				},
				resource.Attribute{
					Name:        "src_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_start_ip",
					Description: `Local proxy ID start.`,
				},
				resource.Attribute{
					Name:        "src_end_ip",
					Description: `Local proxy ID end.`,
				},
				resource.Attribute{
					Name:        "src_subnet",
					Description: `Local proxy ID subnet.`,
				},
				resource.Attribute{
					Name:        "dst_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_name",
					Description: `Local proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_name",
					Description: `Remote proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_start_ip",
					Description: `Remote proxy ID IPv4 start.`,
				},
				resource.Attribute{
					Name:        "dst_end_ip",
					Description: `Remote proxy ID IPv4 end.`,
				},
				resource.Attribute{
					Name:        "dst_subnet",
					Description: `Remote proxy ID IPv4 subnet.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the phase2-interface.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IPsec tunnel name.`,
				},
				resource.Attribute{
					Name:        "phase1name",
					Description: `Phase 1 determines the options required for phase 2.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase2 proposal.`,
				},
				resource.Attribute{
					Name:        "src_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_start_ip",
					Description: `Local proxy ID start.`,
				},
				resource.Attribute{
					Name:        "src_end_ip",
					Description: `Local proxy ID end.`,
				},
				resource.Attribute{
					Name:        "src_subnet",
					Description: `Local proxy ID subnet.`,
				},
				resource.Attribute{
					Name:        "dst_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_name",
					Description: `Local proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_name",
					Description: `Remote proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_start_ip",
					Description: `Remote proxy ID IPv4 start.`,
				},
				resource.Attribute{
					Name:        "dst_end_ip",
					Description: `Remote proxy ID IPv4 end.`,
				},
				resource.Attribute{
					Name:        "dst_subnet",
					Description: `Remote proxy ID IPv4 subnet.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the phase2-interface.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IPsec tunnel name.`,
				},
				resource.Attribute{
					Name:        "phase1name",
					Description: `Phase 1 determines the options required for phase 2.`,
				},
				resource.Attribute{
					Name:        "proposal",
					Description: `Phase2 proposal.`,
				},
				resource.Attribute{
					Name:        "src_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_start_ip",
					Description: `Local proxy ID start.`,
				},
				resource.Attribute{
					Name:        "src_end_ip",
					Description: `Local proxy ID end.`,
				},
				resource.Attribute{
					Name:        "src_subnet",
					Description: `Local proxy ID subnet.`,
				},
				resource.Attribute{
					Name:        "dst_addr_type",
					Description: `Local proxy ID type.`,
				},
				resource.Attribute{
					Name:        "src_name",
					Description: `Local proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_name",
					Description: `Remote proxy ID name.`,
				},
				resource.Attribute{
					Name:        "dst_start_ip",
					Description: `Remote proxy ID IPv4 start.`,
				},
				resource.Attribute{
					Name:        "dst_end_ip",
					Description: `Remote proxy ID IPv4 end.`,
				},
				resource.Attribute{
					Name:        "dst_subnet",
					Description: `Remote proxy ID IPv4 subnet.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comment.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"fortios_firewall_object_address":      0,
		"fortios_firewall_object_addressgroup": 1,
		"fortios_firewall_object_ippool":       2,
		"fortios_firewall_object_service":      3,
		"fortios_firewall_object_servicegroup": 4,
		"fortios_firewall_object_vip":          5,
		"fortios_firewall_object_vipgroup":     6,
		"fortios_firewall_security_policy":     7,
		"fortios_log_fortianalyzer_setting":    8,
		"fortios_log_syslog_setting":           9,
		"fortios_networking_interface_port":    10,
		"fortios_networking_route_static":      11,
		"fortios_system_admin_administrator":   12,
		"fortios_system_admin_profiles":        13,
		"fortios_system_apiuser_setting":       14,
		"fortios_system_license_forticare":     15,
		"fortios_system_license_vdom":          16,
		"fortios_system_license_vm":            17,
		"fortios_system_setting_dns":           18,
		"fortios_system_setting_global":        19,
		"fortios_system_setting_ntp":           20,
		"fortios_system_vdom_setting":          21,
		"fortios_vpn_ipsec_phase1interface":    22,
		"fortios_vpn_ipsec_phase2interface":    23,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
