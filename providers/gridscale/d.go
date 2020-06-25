package gridscale

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "gridscale_firewall",
			Category:         "Data Sources",
			ShortDescription: `Gets data of a firewall by its UUID.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The UUID of the firewall. ## Attributes Reference The following attributes are exported:`,
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
			Type:             "gridscale_ip",
			Category:         "Data Sources",
			ShortDescription: `Gets data of an IP address.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The UUID of the IP address. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the ip.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Defines the IP Address (v4 or v6) the ip.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The IP prefix of the ip.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location, that helps to identify which datacenter an object belongs to.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `failover mode of this ip. If true, then this IP is no longer available for DHCP and can no longer be related to any server..`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the ip.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `The reverse DNS of the ip.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country of the ip.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location of the ip.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the ip was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last ip change.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `Defines if the ip is administratively blocked.`,
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
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the ip.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Defines the IP Address (v4 or v6) the ip.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The IP prefix of the ip.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location, that helps to identify which datacenter an object belongs to.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `failover mode of this ip. If true, then this IP is no longer available for DHCP and can no longer be related to any server..`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the ip.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `The reverse DNS of the ip.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country of the ip.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location of the ip.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the ip was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last ip change.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `Defines if the ip is administratively blocked.`,
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
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_isoimage",
			Category:         "Data Sources",
			ShortDescription: `Gets data of an ISO Image by its UUID.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The UUID of the ISO Image. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "Data Sources",
			ShortDescription: `Get the data of a Load Balancer.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The UUID of the loadbalancer. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the loadbalancer.`,
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the loadbalancer.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_network",
			Category:         "Data Sources",
			ShortDescription: `Get the data of a network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The UUID of the network. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The UUID of the network.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location, that helps to identify which datacenter the network belongs to.`,
				},
				resource.Attribute{
					Name:        "l2security",
					Description: `Defines information about MAC spoofing protection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the network.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of the network.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country where the network is located.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location where the network is located.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `Defines if the network is administratively blocked.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the network was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last network change.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The UUID of the network.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location, that helps to identify which datacenter the network belongs to.`,
				},
				resource.Attribute{
					Name:        "l2security",
					Description: `Defines information about MAC spoofing protection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the network.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of the network.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country where the network is located.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location where the network is located.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `Defines if the network is administratively blocked.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the network was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last network change.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_object_storage_accesskey",
			Category:         "Data Sources",
			ShortDescription: `Gets the data of an access key of an object storage.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) ID of a resource (access key of an object storage). ## Attributes Reference The following attributes are exported:`,
				},
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
			Type:             "gridscale_paas",
			Category:         "Data Sources",
			ShortDescription: `Gets the data of a PaaS based on given UUID.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The UUID of the PaaS service. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The human-readable name of the object.`,
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
					Description: `The UUID of the security zone that the service is running in.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `Network UUID containing security zone.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `The template used to create the service.`,
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
					Description: `Contains the service parameters for the service.`,
				},
				resource.Attribute{
					Name:        "param",
					Description: `Name of parameter.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of the corresponding parameter.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Primitive type of the parameter.`,
				},
				resource.Attribute{
					Name:        "resource_limit",
					Description: `A list of service resource limits.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `The name of the resource you would like to cap.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `The maximum number of the specific resource your service can use.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels in the format [ "label1", "label2" ].`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The human-readable name of the object.`,
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
					Description: `The UUID of the security zone that the service is running in.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `Network UUID containing security zone.`,
				},
				resource.Attribute{
					Name:        "service_template_uuid",
					Description: `The template used to create the service.`,
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
					Description: `Contains the service parameters for the service.`,
				},
				resource.Attribute{
					Name:        "param",
					Description: `Name of parameter.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of the corresponding parameter.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Primitive type of the parameter.`,
				},
				resource.Attribute{
					Name:        "resource_limit",
					Description: `A list of service resource limits.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `The name of the resource you would like to cap.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `The maximum number of the specific resource your service can use.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels in the format [ "label1", "label2" ].`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_public_network",
			Category:         "Data Sources",
			ShortDescription: `Gets data of a public network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The UUID of the public network. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The UUID of the network.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location, that helps to identify which datacenter the network belongs to.`,
				},
				resource.Attribute{
					Name:        "l2security",
					Description: `Defines information about MAC spoofing protection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the network.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of the network.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country where the network is located.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location where the network is located.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `Defines if the network is administratively blocked.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the network was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last network change.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The UUID of the network.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location, that helps to identify which datacenter the network belongs to.`,
				},
				resource.Attribute{
					Name:        "l2security",
					Description: `Defines information about MAC spoofing protection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the network.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of the network.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country where the network is located.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location where the network is located.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `Defines if the network is administratively blocked.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the network was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last network change.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_paas_securityzone",
			Category:         "Data Sources",
			ShortDescription: `Get data of a security zone.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The UUID of the security zone. ## Attributes Reference The following attributes are exported:`,
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
			Attributes: []resource.Attribute{
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_server",
			Category:         "Data Sources",
			ShortDescription: `Gets data of a server by its UUID.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The UUID of the firewall. ## Attributes Reference This resource exports the following attributes:`,
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
			Attributes: []resource.Attribute{
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Gets data of a storage snapshot.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) ID of a resource (UUID of snapshot).`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `(Required) UUID of the storage used to create this snapshot. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `The UUID of the storage used to create this snapshot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the snapshot.`,
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
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `The UUID of the storage used to create this snapshot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the snapshot.`,
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
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_snapshotschedule",
			Category:         "Data Sources",
			ShortDescription: `Gets data of a storage snapshot schedule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) UUID of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `(Required) UUID of the storage that the snapshot schedule belongs to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `UUID of the storage that the snapshot schedule belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The human-readable name of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "next_runtime",
					Description: `The date and time that the snapshot schedule will be run.`,
				},
				resource.Attribute{
					Name:        "keep_snapshots",
					Description: `The amount of Snapshots to keep before overwriting the last created Snapshot.`,
				},
				resource.Attribute{
					Name:        "run_interval",
					Description: `The interval at which the schedule will run (in minutes).`,
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
					Description: `The list of labels.`,
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
					Description: `UUID of the storage that the snapshot schedule belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The human-readable name of the snapshot schedule.`,
				},
				resource.Attribute{
					Name:        "next_runtime",
					Description: `The date and time that the snapshot schedule will be run.`,
				},
				resource.Attribute{
					Name:        "keep_snapshots",
					Description: `The amount of Snapshots to keep before overwriting the last created Snapshot.`,
				},
				resource.Attribute{
					Name:        "run_interval",
					Description: `The interval at which the schedule will run (in minutes).`,
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
					Description: `The list of labels.`,
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
			Category:         "Data Sources",
			ShortDescription: `Gets data of an sshkey.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The UUID of the SSH key. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the sshkey.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The human-readable name of the sshkey.`,
				},
				resource.Attribute{
					Name:        "sshkey",
					Description: `The OpenSSH public key string of the sshkey.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the sshkey.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time of the sshkey was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last sshkey change.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the sshkey.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The human-readable name of the sshkey.`,
				},
				resource.Attribute{
					Name:        "sshkey",
					Description: `The OpenSSH public key string of the sshkey.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the sshkey.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time of the sshkey was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last sshkey change.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_storage",
			Category:         "Data Sources",
			ShortDescription: `Gets data of a storage.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The UUID of the storage. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the storage.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last storage change.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code of the location where storage locates.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the storage.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `The license key (e.g. Windows Servers), if the template used by the storage requires.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country where the storage locates.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the the storage has been running.`,
				},
				resource.Attribute{
					Name:        "last_used_template",
					Description: `The UUID of the last used template on the storage.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity (GB) of the storage.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location where the storage locates.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `The type of the storage.`,
				},
				resource.Attribute{
					Name:        "parent_uuid",
					Description: `The UUID of the parent of the storage.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The human-readable name of the storage.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location where the storage locates.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the storage was initially created.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the storage.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last storage change.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code of the location where storage locates.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the storage.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `The license key (e.g. Windows Servers), if the template used by the storage requires.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country where the storage locates.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the the storage has been running.`,
				},
				resource.Attribute{
					Name:        "last_used_template",
					Description: `The UUID of the last used template on the storage.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity (GB) of the storage.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location where the storage locates.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `The type of the storage.`,
				},
				resource.Attribute{
					Name:        "parent_uuid",
					Description: `The UUID of the parent of the storage.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The human-readable name of the storage.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location where the storage locates.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the storage was initially created.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_template",
			Category:         "Data Sources",
			ShortDescription: `Gets data of a template by name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The exact name of the template as show in [the expert panel of gridscale](https://my.gridscale.io/Expert/Template). ## Attributes Reference The following attributes are exported:`,
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

	dataSourcesMap = map[string]int{

		"gridscale_firewall":                 0,
		"gridscale_ip":                       1,
		"gridscale_isoimage":                 2,
		"gridscale_loadbalancer":             3,
		"gridscale_network":                  4,
		"gridscale_object_storage_accesskey": 5,
		"gridscale_paas":                     6,
		"gridscale_public_network":           7,
		"gridscale_paas_securityzone":        8,
		"gridscale_server":                   9,
		"gridscale_snapshot":                 10,
		"gridscale_snapshotschedule":         11,
		"gridscale_sshkey":                   12,
		"gridscale_storage":                  13,
		"gridscale_template":                 14,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
