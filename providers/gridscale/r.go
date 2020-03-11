package gridscale

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "gridscale_firewall",
			Category:         "Resources",
			ShortDescription: `Manages a firewall in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "rules_v4_in",
					Description: `(Optional`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v4_out",
					Description: `(Optional`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_in",
					Description: `(Optional`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_out",
					Description: `(Optional`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the firewall.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the firewall.`,
				},
				resource.Attribute{
					Name:        "rules_v4_in",
					Description: `Firewall template rules for inbound traffic - covers ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v4_out",
					Description: `Firewall template rules for outbound traffic - covers ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_in",
					Description: `Firewall template rules for inbound traffic - covers ipv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_out",
					Description: `Firewall template rules for outbound traffic - covers ipv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The information about networks which are related to this firewall.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `The object UUID or id of the firewall.`,
				},
				resource.Attribute{
					Name:        "object_name",
					Description: `Name of the firewall.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `The object UUID or id of the network.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the network.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `The object is private, the value will be true. Otherwise the value will be false.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the firewall.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the firewall.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the firewall.`,
				},
				resource.Attribute{
					Name:        "rules_v4_in",
					Description: `Firewall template rules for inbound traffic - covers ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v4_out",
					Description: `Firewall template rules for outbound traffic - covers ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_in",
					Description: `Firewall template rules for inbound traffic - covers ipv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_out",
					Description: `Firewall template rules for outbound traffic - covers ipv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The information about networks which are related to this firewall.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `The object UUID or id of the firewall.`,
				},
				resource.Attribute{
					Name:        "object_name",
					Description: `Name of the firewall.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `The object UUID or id of the network.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the network.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `The object is private, the value will be true. Otherwise the value will be false.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the firewall.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_ipv4",
			Category:         "Resources",
			ShortDescription: `Manages an IPv4 address in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"ipv4",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `(Optional) Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `(Optional) Defines the reverse DNS entry for the IP Address (PTR Resource Record).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Defines the IP Address.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The network address and the subnet.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Formatted by the 2 digit country code (ISO 3166-2) of the host country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `Defines if the object is administratively blocked. If true, it can not be deleted by the user.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `The amount of minutes the IP address has been in use.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_ipv6",
			Category:         "Resources",
			ShortDescription: `Manages an IPv6 address in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"ipv6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `(Optional) Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `(Optional) Defines the reverse DNS entry for the IP Address (PTR Resource Record).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `Helps to identify which datacenter an object belongs to. The location of the resource depends on the location of the project.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Defines the IP Address.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The network address and the subnet.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Formatted by the 2 digit country code (ISO 3166-2) of the host country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `Defines if the object is administratively blocked. If true, it can not be deleted by the user.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `The amount of minutes the IP address has been in use.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_isoimage",
			Category:         "Resources",
			ShortDescription: `Manages an ISO Image in Gridscale.`,
			Description:      ``,
			Keywords: []string{
				"isoimage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "source_url",
					Description: `(Required) Contains the source URL of the ISO Image that it was originally fetched from.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ISO Image.`,
				},
				resource.Attribute{
					Name:        "source_url",
					Description: `Contains the source URL of the ISO Image that it was originally fetched from.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The information about servers which are related to this ISO Image.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `The object UUID or id of the server.`,
				},
				resource.Attribute{
					Name:        "object_name",
					Description: `Name of the server.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "bootdevice",
					Description: `True if the ISO Image is a boot device of this server.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the ISO Image.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `Helps to identify which datacenter an object belongs to.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Formatted by the 2 digit country code (ISO 3166-2) of the host country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the ISO Image.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `The object is private, the value will be true. Otherwise the value will be false.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Template.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the object has been running.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of a storage/ISO Image/ISO Image/snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `Defines the price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ISO Image.`,
				},
				resource.Attribute{
					Name:        "source_url",
					Description: `Contains the source URL of the ISO Image that it was originally fetched from.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The information about servers which are related to this ISO Image.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `The object UUID or id of the server.`,
				},
				resource.Attribute{
					Name:        "object_name",
					Description: `Name of the server.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "bootdevice",
					Description: `True if the ISO Image is a boot device of this server.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the ISO Image.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `Helps to identify which datacenter an object belongs to.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Formatted by the 2 digit country code (ISO 3166-2) of the host country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the ISO Image.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `The object is private, the value will be true. Otherwise the value will be false.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Template.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the object has been running.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of a storage/ISO Image/ISO Image/snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `Defines the price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_loadbalancer",
			Category:         "Resources",
			ShortDescription: `Manage a loadbalancer in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"loadbalancer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "redirect_http_to_https",
					Description: `(Required) Whether the loadbalancer is forced to redirect requests from HTTP to HTTPS.`,
				},
				resource.Attribute{
					Name:        "listen_ipv4_uuid",
					Description: `(Required) The UUID of the IPv4 address the loadbalancer will listen to for incoming requests.`,
				},
				resource.Attribute{
					Name:        "listen_ipv6_uuid",
					Description: `(Required) The UUID of the IPv6 address the loadbalancer will listen to for incoming requests.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Required) The algorithm used to process requests. Accepted values: roundrobin/leastconn.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the loadbalancer.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `Helps to identify which datacenter an object belongs to. The location of the resource depends on the location of the project.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The human-readable name of the loadbalancer.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The algorithm used to process requests.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the loadbalancer.`,
				},
				resource.Attribute{
					Name:        "redirect_http_to_https",
					Description: `Whether the Load balancer is forced to redirect requests from HTTP to HTTPS.`,
				},
				resource.Attribute{
					Name:        "listen_ipv4_uuid",
					Description: `The UUID of the IPv4 address the loadbalancer will listen to for incoming requests.`,
				},
				resource.Attribute{
					Name:        "listen_ipv6_uuid",
					Description: `The UUID of the IPv6 address the loadbalancer will listen to for incoming requests.`,
				},
				resource.Attribute{
					Name:        "forwarding_rule",
					Description: `The forwarding rules of the loadbalancer.`,
				},
				resource.Attribute{
					Name:        "backend_server",
					Description: `The servers that the loadbalancer can communicate with.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_network",
			Category:         "Resources",
			ShortDescription: `Manages a network in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "l2security",
					Description: `(Optional) Defines information about MAC spoofing protection (filters layer2 and ARP traffic based on MAC source). It can only be (de-)activated on a private network - the public network always has l2security enabled. It will be true if the network is public, and false if the network is private.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `Helps to identify which datacenter an object belongs to. The location of the resource depends on the location of the project.`,
				},
				resource.Attribute{
					Name:        "l2security",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of this network, can be mpls, breakout or network.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Formatted by the 2 digit country code (ISO 3166-2) of the host country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "public_net",
					Description: `Is the network public or not.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `If deleting this network is allowed.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_object_storage_accesskey",
			Category:         "Resources",
			ShortDescription: `Manages an access key of an object storage in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"accesskey",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The access key of the object storage.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `Access key of an object storage.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Secret key of an object storage.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_paas_securityzone",
			Category:         "Resources",
			ShortDescription: `Manages a PaaS in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"paas",
				"securityzone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `(Required) The template used to create the service.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: `(Optional) The UUID of the security zone that the service is running in.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) Contains the service parameters for the service.`,
				},
				resource.Attribute{
					Name:        "param",
					Description: `(Required) Name of parameter.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value of the corresponding parameter.`,
				},
				resource.Attribute{
					Name:        "resource_limit",
					Description: `(Optional) A list of service resource limits..`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) The name of the resource you would like to cap.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Required) The maximum number of the specific resource your service can use.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Primitive type of the parameter: bool, int (better use float for int case), float, string. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for PaaS service.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for PaaS service.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `Ports that PaaS service listens to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a port.`,
				},
				resource.Attribute{
					Name:        "listen_port",
					Description: `Port number.`,
				},
				resource.Attribute{
					Name:        "security_zone_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `Network UUID containing security zone.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "usage_in_minute",
					Description: `Number of minutes that PaaS service is in use.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `Current price of PaaS service.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Time of the last change.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time of the creation.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of PaaS service.`,
				},
				resource.Attribute{
					Name:        "parameter",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "param",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "resource_limit",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_paas",
			Category:         "Resources",
			ShortDescription: `Manages a security zone in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"paas",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `(Optional) Helps to identify which datacenter an object belongs to. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the security zone.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `Helps to identify which datacenter an object belongs to.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the location's country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels.`,
				},
				resource.Attribute{
					Name:        "relations",
					Description: `List of PaaS services' UUIDs relating to the security zone.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_server",
			Category:         "Resources",
			ShortDescription: `Manages a server in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `(Required) The number of server cores.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) The amount of server memory in GB.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "auto_recovery",
					Description: `(Optional) If the server should be auto-started in case of a failure (default=true).`,
				},
				resource.Attribute{
					Name:        "hardware_profile",
					Description: `(Optional, ForceNew) The hardware profile of the Server. Options are default, legacy, nested, cisco_csr, sophos_utm, f5_bigip and q35 at the moment of writing. Check the`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) The UUID of the IPv4 address of the server. (`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) The UUID of the IPv6 address of the server. (`,
				},
				resource.Attribute{
					Name:        "isoimage",
					Description: `(Optional) The UUID of an ISO image in gridscale. The server will automatically boot from the ISO if one was added. The UUIDs of ISO images can be found in [the expert panel](https://my.gridscale.io/Expert/ISOImage).`,
				},
				resource.Attribute{
					Name:        "power",
					Description: `(Optional, Computed) The power state of the server. Set this to true to will boot the server, false will shut it down.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, Computed) Defines which Availability-Zone the Server is placed.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Optional) Connects a storage to the server.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `(Required) The object UUID or id of the storage.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) Connects a network to the server.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `(Required) The object UUID or id of the network.`,
				},
				resource.Attribute{
					Name:        "bootdevice",
					Description: `(Optional, Computed) Make this network the boot device. This can only be set for one network.`,
				},
				resource.Attribute{
					Name:        "firewall_template_uuid",
					Description: `(Optional) The UUID of firewall template.`,
				},
				resource.Attribute{
					Name:        "rules_v4_in",
					Description: `(Optional) Firewall template rules for inbound traffic - covers ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v4_out",
					Description: `(Optional) Firewall template rules for outbound traffic - covers ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_in",
					Description: `(Optional) Firewall template rules for inbound traffic - covers ipv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_out",
					Description: `(Optional) Firewall template rules for outbound traffic - covers ipv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the server.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `The number of server cores.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of server memory in GB.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `Helps to identify which datacenter an object belongs to. The location of the resource depends on the location of the project.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "hardware_profile",
					Description: `The hardware profile of the Server.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `Connects a storage to the server.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `The object UUID or id of the storage.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Indicates the speed of the storage. This may be (storage, storage_high or storage_insane).`,
				},
				resource.Attribute{
					Name:        "bootdevice",
					Description: `True is the storage is a boot device.`,
				},
				resource.Attribute{
					Name:        "object_name",
					Description: `Name of the storage.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Capacity of the storage (GB).`,
				},
				resource.Attribute{
					Name:        "controller",
					Description: `Defines the SCSI controller id. The SCSI defines transmission routes such as Serial Attached SCSI (SAS), Fibre Channel and iSCSI.`,
				},
				resource.Attribute{
					Name:        "bus",
					Description: `The SCSI bus id. The SCSI defines transmission routes like Serial Attached SCSI (SAS), Fibre Channel and iSCSI. Each SCSI device is addressed via a specific number. Each SCSI bus can have multiple SCSI devices connected to it.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Defines the SCSI target ID. The target ID is a device (e.g. disk).`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `Is the common SCSI abbreviation of the Logical Unit Number. A lun is a unique identifier for a single disk or a composite of disks.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "last_used_template",
					Description: `Indicates the UUID of the last used template on this storage (inherited from snapshots).`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `Connects a network to the server.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `The object UUID or id of the network.`,
				},
				resource.Attribute{
					Name:        "bootdevice",
					Description: `Make this network the boot device. This can only be set for one network.`,
				},
				resource.Attribute{
					Name:        "object_name",
					Description: `Name of the network.`,
				},
				resource.Attribute{
					Name:        "ordering",
					Description: `Defines the ordering of the network interfaces. Lower numbers have lower PCI-IDs.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `One of network, network_high, network_insane.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `network_mac defines the MAC address of the network interface.`,
				},
				resource.Attribute{
					Name:        "firewall_template_uuid",
					Description: `The UUID of firewall template.`,
				},
				resource.Attribute{
					Name:        "rules_v4_in",
					Description: `Firewall template rules for inbound traffic - covers ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v4_out",
					Description: `Firewall template rules for outbound traffic - covers ipv4 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_in",
					Description: `Firewall template rules for inbound traffic - covers ipv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "rules_v6_out",
					Description: `Firewall template rules for outbound traffic - covers ipv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order at which the firewall will compare packets against its rules, a packet will be compared against the first rule, it will either allow it to pass or block it and it won t be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `This defines what the firewall will do. Either accept or drop.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Either 'udp' or 'tcp'.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `A Number between 1 and 65535, port ranges are separated by a colon for FTP.`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `The UUID of the IPv4 address of the server.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `The UUID of the IPv6 address of the server.`,
				},
				resource.Attribute{
					Name:        "isoimage",
					Description: `The UUID of an ISO image in gridscale.`,
				},
				resource.Attribute{
					Name:        "power",
					Description: `The power state of the server.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Defines which Availability-Zone the Server is placed.`,
				},
				resource.Attribute{
					Name:        "auto_recovery",
					Description: `If the server should be auto-started in case of a failure.`,
				},
				resource.Attribute{
					Name:        "console_token",
					Description: `The token used by the panel to open the websocket VNC connection to the server console.`,
				},
				resource.Attribute{
					Name:        "legacy",
					Description: `Legacy-Hardware emulation instead of virtio hardware. If enabled, hotplugging cores, memory, storage, network, etc. will not work, but the server will most likely run every x86 compatible operating system. This mode comes with a performance penalty, as emulated hardware does not benefit from the virtio driver infrastructure.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes_memory",
					Description: `Total minutes of memory used.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes_cores",
					Description: `Total minutes of cores used.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_snapshot",
			Category:         "Resources",
			ShortDescription: `Manages a storage snapshot in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `(Required) UUID of the storage used to create this snapshot.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) The list of labels.`,
				},
				resource.Attribute{
					Name:        "rollback",
					Description: `(Optional) Returns a storage to the state of the selected Snapshot.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID of the rollback request. It can be any string value. Each rollback request has to have a UNIQUE id. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location, that helps to identify which datacenter an object belongs to.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country of the snapshot.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location of the snapshot.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the ip was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last snapshot change.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the ip has been running.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of the snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "rollback",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rollback_time",
					Description: `The time when rollback request is fulfilled.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the rollback request.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location, that helps to identify which datacenter an object belongs to.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country of the snapshot.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location of the snapshot.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the ip was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last snapshot change.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the ip has been running.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of the snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "rollback",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rollback_time",
					Description: `The time when rollback request is fulfilled.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the rollback request.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_snapshotschedule",
			Category:         "Resources",
			ShortDescription: `Manages a storage snapshot schedule.`,
			Description:      ``,
			Keywords: []string{
				"snapshotschedule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) UUID of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `(Required) UUID of the storage that the snapshot schedule belongs to.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) The list of labels.`,
				},
				resource.Attribute{
					Name:        "next_runtime",
					Description: `(Optional) The date and time that the snapshot schedule will be run.`,
				},
				resource.Attribute{
					Name:        "keep_snapshots",
					Description: `(Required) The amount of Snapshots to keep before overwriting the last created Snapshot (>=1).`,
				},
				resource.Attribute{
					Name:        "run_interval",
					Description: `(Required) The interval at which the schedule will run (in minutes, >=60). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_runtime",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "keep_snapshots",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "run_interval",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the snapshot schedule was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last snapshot schedule change.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `Related snapshots.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `UUID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the snapshot was initially created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_runtime",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "keep_snapshots",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "run_interval",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the snapshot schedule was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last snapshot schedule change.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `Related snapshots.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "object_uuid",
					Description: `UUID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the snapshot was initially created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_sshkey",
			Category:         "Resources",
			ShortDescription: `Manages an SSH public key in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"sshkey",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "sshkey",
					Description: `(Required) This is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ]. ## Attributes This resource exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sshkey",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_storage",
			Category:         "Resources",
			ShortDescription: `Manages a storage in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"storage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the object. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Required) required (integer - minimum: 1 - maximum: 4096).`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) (one of storage, storage_high, storage_insane).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) List of labels in the format [ "label1", "label2" ].`,
				},
				resource.Attribute{
					Name:        "template_uuid",
					Description: `(Required) The UUID of a template. This can be found in the [expert panel](https://my.gridscale.io/Expert/Template) by clicking more on the template or by using a gridscale_template datasource.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root (Linux) or Administrator (Windows) password to set for the installed storage. Valid only for public templates. The password has to be either plain-text or a crypt string (modular crypt format - MCF).`,
				},
				resource.Attribute{
					Name:        "password_type",
					Description: `(Optional) (one of plain, crypt) Required if password is set (ignored for private templates and public Windows templates).`,
				},
				resource.Attribute{
					Name:        "sshkeys",
					Description: `(Optional) (array of any - minItems: 0) Public Linux templates only! The UUIDs of SSH keys to be added for the root user.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The hostname of the installed server (ignored for private templates and public windows templates). ~>`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `Helps to identify which datacenter an object belongs to. The location of the resource depends on the location of the project.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the object was created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Formatted by the 2 digit country code (ISO 3166-2) of the host country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The location name.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "last_used_template",
					Description: `Indicates the UUID of the last used template on this storage (inherited from snapshots).`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `The amount of minutes the IP address has been in use.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_template",
			Category:         "Resources",
			ShortDescription: `Manages a template in gridscale.`,
			Description:      ``,
			Keywords: []string{
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The exact name of the template as show in [the expert panel of gridscale](https://my.gridscale.io/Expert/Template).`,
				},
				resource.Attribute{
					Name:        "snapshot_uuid",
					Description: `(Required) Snapshot uuid for template.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) List of labels. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the template.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the template.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `Helps to identify which datacenter an object belongs to.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Formatted by the 2 digit country code (ISO 3166-2) of the host country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "ostype",
					Description: `The operating system installed in the template.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the template.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `The object is private, the value will be true. Otherwise the value will be false.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "distro",
					Description: `The OS distrobution that the Template contains.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Template.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the object has been running.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of a storage/ISO Image/template/snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `Defines the price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the template.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the template.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `Helps to identify which datacenter an object belongs to.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `Formatted by the 2 digit country code (ISO 3166-2) of the host country.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `Uses IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location. It supports the full UTF-8 charset, with a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status indicates the status of the object.`,
				},
				resource.Attribute{
					Name:        "ostype",
					Description: `The operating system installed in the template.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the template.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `The object is private, the value will be true. Otherwise the value will be false.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the object was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last object change.`,
				},
				resource.Attribute{
					Name:        "distro",
					Description: `The OS distrobution that the Template contains.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Template.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the object has been running.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of a storage/ISO Image/template/snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `Defines the price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"gridscale_firewall":                 0,
		"gridscale_ipv4":                     1,
		"gridscale_ipv6":                     2,
		"gridscale_isoimage":                 3,
		"gridscale_loadbalancer":             4,
		"gridscale_network":                  5,
		"gridscale_object_storage_accesskey": 6,
		"gridscale_paas_securityzone":        7,
		"gridscale_paas":                     8,
		"gridscale_server":                   9,
		"gridscale_snapshot":                 10,
		"gridscale_snapshotschedule":         11,
		"gridscale_sshkey":                   12,
		"gridscale_storage":                  13,
		"gridscale_template":                 14,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
