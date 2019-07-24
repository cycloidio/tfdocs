package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "brightbox_cloudip",
			Category:         "Resources",
			ShortDescription: `Provides a Brightbox CloudIP resource.`,
			Description:      ``,
			Keywords: []string{
				"cloudip",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) a label to assign to the CloudIP`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `(Optional) The reverse DNS entry for the CloudIP`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) The CloudIP mapping target. This is the interface id from a server, or the id of a load balancer, server group or cloud sql resource.`,
				},
				resource.Attribute{
					Name:        "port_translator",
					Description: `(Optional) An array of port translator blocks. The Port Translator block is descibed below Note that the default group for each account cannot be used as the target for a cloud ip. Port Translator (` + "`" + `port_translator` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "incoming",
					Description: `(Required) The Port number traffic is coming in on the network`,
				},
				resource.Attribute{
					Name:        "outgoing",
					Description: `(Required) The Port number traffic is received at the mapped device`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol of the port translator. Either ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + ` ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the CloudIP`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name of the CloudIP`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `the public IPV4 address of the CloudIP`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current state of the CloudIP: ` + "`" + `mapped` + "`" + ` or ` + "`" + `unmapped` + "`" + ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username used to log onto the server ## Import CloudIPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_cloudip.mycloudip cip-vsalc ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_cloudip` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Mapping Cloud IPs - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Unmapping Cloud IPs`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the CloudIP`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name of the CloudIP`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `the public IPV4 address of the CloudIP`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current state of the CloudIP: ` + "`" + `mapped` + "`" + ` or ` + "`" + `unmapped` + "`" + ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username used to log onto the server ## Import CloudIPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_cloudip.mycloudip cip-vsalc ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_cloudip` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Mapping Cloud IPs - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Unmapping Cloud IPs`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_database_server",
			Category:         "Resources",
			ShortDescription: `Provides a Brightbox Database Server resource. This can be used to create, modify, and delete Database Servers.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"server",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A label assigned to the Database Server`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A further description of the Database Server`,
				},
				resource.Attribute{
					Name:        "maintenance_weekday",
					Description: `(Optional) Numerical index of weekday (0 is Sunday, 1 is Monday...) to set when automatic updates may be performed. Default is 0 (Sunday).`,
				},
				resource.Attribute{
					Name:        "maintenance_hour",
					Description: `(Optional) Number representing 24hr time start of maintenance window hour for x:00-x:59 (0-23). Default is 6`,
				},
				resource.Attribute{
					Name:        "snapshots_schedule",
					Description: `(Optional) A crontab pattern to determine approximately when scheduled snapshots will run (must be at least hourly)`,
				},
				resource.Attribute{
					Name:        "database_engine",
					Description: `(Optional) Database engine to request. Default is mysql.`,
				},
				resource.Attribute{
					Name:        "database_version",
					Description: `(Optional) Database version to request. Default is 8.0.`,
				},
				resource.Attribute{
					Name:        "database_type",
					Description: `(Optional) ID of the Database Type required.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The handle of the zone required (` + "`" + `gb1-a` + "`" + `, ` + "`" + `gb1-b` + "`" + `) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Database Server`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `The username used to log onto the database`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `The password used to log onto the database`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current state of the database server, usually ` + "`" + `active` + "`" + ` or ` + "`" + `deleted` + "`" + ``,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `True if database server has been set to locked and cannot be deleted`,
				},
				resource.Attribute{
					Name:        "snapshots_schedule_next_at",
					Description: `The approximate UTC time when the next snapshot is scheduled ## Import Database Servers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_database_server.mydatabase dbs-qwert ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_database_server` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Creating Databases - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Deleting Databases`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Database Server`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `The username used to log onto the database`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `The password used to log onto the database`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current state of the database server, usually ` + "`" + `active` + "`" + ` or ` + "`" + `deleted` + "`" + ``,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `True if database server has been set to locked and cannot be deleted`,
				},
				resource.Attribute{
					Name:        "snapshots_schedule_next_at",
					Description: `The approximate UTC time when the next snapshot is scheduled ## Import Database Servers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_database_server.mydatabase dbs-qwert ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_database_server` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Creating Databases - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Deleting Databases`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_firewall_policy",
			Category:         "Resources",
			ShortDescription: `Provides a Brightbox Firewall Policy resource.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"policy",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "server_group",
					Description: `(Optional) The ID of the Server Group the policy will be applied to`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A label to assign to the Firewall Policy`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A further description of the Firewall Policy ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Firewall Policy ## Import Firewall Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_firewall_policy.mypolicy fwp-zxcvb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Firewall Policy ## Import Firewall Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_firewall_policy.mypolicy fwp-zxcvb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_firewall_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Brightbox Firewall Rule resource.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"rule",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "firewall_policy",
					Description: `(Required) The ID of the firewall policy this rule belongs to`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol Number or one of ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + ``,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Subnet, ServerGroup or ServerID. ` + "`" + `any` + "`" + `,` + "`" + `10.1.1.23/32` + "`" + ` or ` + "`" + `srv-4ktk4` + "`" + ``,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) single port, multiple ports or range separated by ` + "`" + `-` + "`" + ` or ` + "`" + `:` + "`" + `; upto 255 characters. Example - ` + "`" + `80` + "`" + `, ` + "`" + `80,443,21` + "`" + ` or ` + "`" + `3000-3999` + "`" + ``,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Optional) Subnet, ServerGroup or ServerID. ` + "`" + `any` + "`" + `,` + "`" + `10.1.1.23/32` + "`" + ` or ` + "`" + `srv-4ktk4` + "`" + ``,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Optional) single port, multiple ports or range separated by ` + "`" + `-` + "`" + ` or ` + "`" + `:` + "`" + `; upto 255 characters. Example - ` + "`" + `80` + "`" + `, ` + "`" + `80,443,21` + "`" + ` or ` + "`" + `3000-3999` + "`" + ``,
				},
				resource.Attribute{
					Name:        "icmp_type_name",
					Description: `(Optional) ICMP type name. ` + "`" + `echo-request` + "`" + `, ` + "`" + `echo-reply` + "`" + `. Only allowed if protocol is ` + "`" + `icmp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A further description of the Firewall Rule ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Firewall Rule ## Import Firewall Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_firewall_rule.myrule fwr-ghjkl ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Firewall Rule ## Import Firewall Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_firewall_rule.myrule fwr-ghjkl ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_load_balancer",
			Category:         "Resources",
			ShortDescription: `Provides a Brightbox Load Balancer resource. This can be used to create, modify, and delete Load Balancers.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A label assigned to the Load Balancer`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) Method of load balancing to use, either ` + "`" + `least-connections` + "`" + ` or ` + "`" + `round-robin` + "`" + ``,
				},
				resource.Attribute{
					Name:        "certificate_pem",
					Description: `(Optional) A X509 SSL certificate in PEM format. Must be included along with ` + "`" + `certificate_key` + "`" + `. If intermediate certificates are required they should be concatenated after the main certificate`,
				},
				resource.Attribute{
					Name:        "certificate_private_key",
					Description: `(Optional) The RSA private key used to sign the certificate in PEM format. Must be included along with ` + "`" + `certificate_pem` + "`" + ``,
				},
				resource.Attribute{
					Name:        "sslv3",
					Description: `(Optional) Allow SSL v3 to be used. Default is ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "buffer_size",
					Description: `(Optional) Buffer size in bytes`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `(Optional) An array of Server IDs`,
				},
				resource.Attribute{
					Name:        "listener",
					Description: `(Required) An array of listener blocks. The Listener block is described below`,
				},
				resource.Attribute{
					Name:        "healthcheck",
					Description: `(Required) A healthcheck block. The Healthcheck block is described below Listener (` + "`" + `listener` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol of the listener. One of ` + "`" + `tcp` + "`" + `, ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `http+ws` + "`" + `, ` + "`" + `https+wss` + "`" + ``,
				},
				resource.Attribute{
					Name:        "in",
					Description: `(Required) Port to listen on`,
				},
				resource.Attribute{
					Name:        "out",
					Description: `(Required) Port to pass through to`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout of connection in milliseconds. Default is 50000 Health Check (` + "`" + `healthcheck` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of health check required: ` + "`" + `tcp` + "`" + ` or ` + "`" + `http` + "`" + ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port to connect to to check health`,
				},
				resource.Attribute{
					Name:        "request",
					Description: `(Optional) Path used for HTTP check`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Frequency of checks in milliseconds`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout of health check in milliseconds`,
				},
				resource.Attribute{
					Name:        "threshold_up",
					Description: `(Optional) Number of checks that must pass before connection is considered healthy`,
				},
				resource.Attribute{
					Name:        "threshold_down",
					Description: `(Optional) Number of checks that must fail before connection is considered unhealthy ## Attributes Reference The following attributes are exported`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current state of the load balancer. Usually ` + "`" + `creating` + "`" + ` or ` + "`" + `active` + "`" + ``,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `True if the database server has been set to locked and cannot be deleted ## Import Load Balancers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_load_balancer.mylba lba-12345 ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_load_balancer` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Creating Load Balancers - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Deleting Load Balancers`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current state of the load balancer. Usually ` + "`" + `creating` + "`" + ` or ` + "`" + `active` + "`" + ``,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `True if the database server has been set to locked and cannot be deleted ## Import Load Balancers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_load_balancer.mylba lba-12345 ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_load_balancer` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Creating Load Balancers - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Deleting Load Balancers`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_orbit_container",
			Category:         "Resources",
			ShortDescription: `Provides a Brightbox Orbit Container resource. This can be used to create, modify, and delete Containers in Orbit.`,
			Description:      ``,
			Keywords: []string{
				"orbit",
				"container",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A label assigned to the Orbit container`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A dictionary of metadata key/value items. The key must be in lower case with no underscores or spaces`,
				},
				resource.Attribute{
					Name:        "object_count",
					Description: `The number of items in the Orbit Container`,
				},
				resource.Attribute{
					Name:        "bytes_used",
					Description: `The total size of the items in the Orbit Container`,
				},
				resource.Attribute{
					Name:        "storage_policy",
					Description: `The storage policy in place for this container. Always 'Policy-0' at present`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time the container was created ## Import Orbit Containers can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_orbit_container.myorbitcontainer initial ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "object_count",
					Description: `The number of items in the Orbit Container`,
				},
				resource.Attribute{
					Name:        "bytes_used",
					Description: `The total size of the items in the Orbit Container`,
				},
				resource.Attribute{
					Name:        "storage_policy",
					Description: `The storage policy in place for this container. Always 'Policy-0' at present`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time the container was created ## Import Orbit Containers can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_orbit_container.myorbitcontainer initial ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_server",
			Category:         "Resources",
			ShortDescription: `Provides a Brightbox Server resource. This can be used to create, modify, and delete Servers. Servers also support provisioning.`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "image",
					Description: `(Required) The Server image ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Server name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The handle of the server type required (` + "`" + `1gb.ssd` + "`" + `, etc)`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The handle of the zone required (` + "`" + `gb1-a` + "`" + `, ` + "`" + `gb1-b` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Server`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name of server`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `short name of server, usually the same as the ` + "`" + `id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `the id reference of the network interface. Used to target cloudips.`,
				},
				resource.Attribute{
					Name:        "ipv4_address_private",
					Description: `The RFC 1912 address of the server`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `the IPv6 address of the server`,
				},
				resource.Attribute{
					Name:        "ipv6_hostname",
					Description: `the FQDN of the IPv6 address`,
				},
				resource.Attribute{
					Name:        "public_hostname",
					Description: `the FQDN of the public IPv4 address. Appears if a cloud ip is mapped`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `the public IPV4 address of the server. Appears if a cloud ip is mapped`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `True if server has been set to locked and cannot be deleted`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current state of the server, usually ` + "`" + `active` + "`" + `, ` + "`" + `inactive` + "`" + ` or ` + "`" + `deleted` + "`" + ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username used to log onto the server ## Import Servers can be imported using the server ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_server.myserver srv-ojy3o ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_server` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Creating Servers - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Deleting Servers`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Server`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name of server`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `short name of server, usually the same as the ` + "`" + `id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `the id reference of the network interface. Used to target cloudips.`,
				},
				resource.Attribute{
					Name:        "ipv4_address_private",
					Description: `The RFC 1912 address of the server`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `the IPv6 address of the server`,
				},
				resource.Attribute{
					Name:        "ipv6_hostname",
					Description: `the FQDN of the IPv6 address`,
				},
				resource.Attribute{
					Name:        "public_hostname",
					Description: `the FQDN of the public IPv4 address. Appears if a cloud ip is mapped`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `the public IPV4 address of the server. Appears if a cloud ip is mapped`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `True if server has been set to locked and cannot be deleted`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current state of the server, usually ` + "`" + `active` + "`" + `, ` + "`" + `inactive` + "`" + ` or ` + "`" + `deleted` + "`" + ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username used to log onto the server ## Import Servers can be imported using the server ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_server.myserver srv-ojy3o ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_server` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Creating Servers - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Deleting Servers`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_server_group",
			Category:         "Resources",
			ShortDescription: `Provides a Brightbox Server Group resource. This can be used to create, modify, and delete Server Groups.`,
			Description:      ``,
			Keywords: []string{
				"server",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A label assigned to the Server Group`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A further description of the Server Group ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Server ## Import Server Groups can be imported using the server group ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_server_group.default grp-ok8vw ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Server ## Import Server Groups can be imported using the server group ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_server_group.default grp-ok8vw ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"brightbox_cloudip":         0,
		"brightbox_database_server": 1,
		"brightbox_firewall_policy": 2,
		"brightbox_firewall_rule":   3,
		"brightbox_load_balancer":   4,
		"brightbox_orbit_container": 5,
		"brightbox_server":          6,
		"brightbox_server_group":    7,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
