package brightbox

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "brightbox_api_client",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A label to assign to the API Client`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A further description of the API Client`,
				},
				resource.Attribute{
					Name:        "permissions_group",
					Description: `(Optional) The type of API Client required, either ` + "`" + `full` + "`" + ` or ` + "`" + `storage` + "`" + `. The default is ` + "`" + `full` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Client`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `The initial secret key of the API Client`,
				},
				resource.Attribute{
					Name:        "account",
					Description: `The ID of the account the API Client is linked to`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the API Client`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `The initial secret key of the API Client`,
				},
				resource.Attribute{
					Name:        "account",
					Description: `The ID of the account the API Client is linked to`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_cloudip",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
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
					Name:        "mode",
					Description: `(Optional) Type of CloudIP required, either ` + "`" + `nat` + "`" + ` or ` + "`" + `route` + "`" + `.`,
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
					Name:        "public_ipv4",
					Description: `the public IPV4 address of the CloudIP`,
				},
				resource.Attribute{
					Name:        "public_ipv6",
					Description: `the public IPV6 address of the CloudIP`,
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the CloudIP`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `Fully Qualified Domain Name of the CloudIP`,
				},
				resource.Attribute{
					Name:        "public_ipv4",
					Description: `the public IPV4 address of the CloudIP`,
				},
				resource.Attribute{
					Name:        "public_ipv6",
					Description: `the public IPV6 address of the CloudIP`,
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
			Type:             "brightbox_config_map",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A label assigned to the Config Map`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Required) A key value map of strings ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Server ## Import Config Maps can be imported using the config map ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_config_map.default cfg-ok8vw ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Server ## Import Config Maps can be imported using the config map ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_config_map.default cfg-ok8vw ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_database_server",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
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
					Name:        "snapshots_retention",
					Description: `(Optional) Keep this number of scheduled snapshots. Keep all if unset.`,
				},
				resource.Attribute{
					Name:        "snapshots_schedule",
					Description: `(Optional) A crontab pattern to determine approximately when scheduled snapshots will run (must be at least hourly)`,
				},
				resource.Attribute{
					Name:        "database_engine",
					Description: `(Optional) Database engine to request. Default is mysql`,
				},
				resource.Attribute{
					Name:        "database_version",
					Description: `(Optional) Database version to request. Default is 8.0`,
				},
				resource.Attribute{
					Name:        "database_type",
					Description: `(Optional) ID of the Database Type required`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The handle of the zone required (` + "`" + `gb1-a` + "`" + `, ` + "`" + `gb1-b` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `(Optional) Set to true to stop the database server from being deleted ## Attributes Reference The following attributes are exported:`,
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
					Name:        "snapshots_schedule_next_at",
					Description: `The approximate UTC time when the next snapshot is scheduled ## Import Database Servers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_database_server.mydatabase dbs-qwert ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_database_server` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Creating Databases - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Deleting Databases`,
				},
			},
			Attributes: []resource.Attribute{
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
					Name:        "snapshots_schedule_next_at",
					Description: `The approximate UTC time when the next snapshot is scheduled ## Import Database Servers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_database_server.mydatabase dbs-qwert ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_database_server` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Creating Databases - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Deleting Databases`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_firewall_policy",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
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
					Name:        "https_redirect",
					Description: `(Optional) Redirect any requests on port 80 automatically to port 443`,
				},
				resource.Attribute{
					Name:        "ssl_minimum_version",
					Description: `(Optional) The minimum TLS/SSL version for the load balancer to accept. Supports ` + "`" + `TLSv1.0` + "`" + `, ` + "`" + `TLSv1.1` + "`" + `, ` + "`" + `TLSv1.2` + "`" + `, ` + "`" + `TLSv1.3` + "`" + ` and ` + "`" + `SSLv3` + "`" + ``,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `(Optional) Set to true to stop the load balancer from being deleted`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `(Optional) An array of Server IDs`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `(Optional) An array of domain names to attempt to register with ACME. Conflicts with ` + "`" + `certificate_pem` + "`" + ` and ` + "`" + `certificate_private_key` + "`" + ``,
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
					Description: `(Optional) Timeout of connection in milliseconds. Default is 50000`,
				},
				resource.Attribute{
					Name:        "proxy_protocol",
					Description: `(Optional) Proxy Protocol version supported by backend server. One of ` + "`" + `v1` + "`" + `, ` + "`" + `v2` + "`" + `, ` + "`" + `v2-ssl` + "`" + `, ` + "`" + `v2-ssl-cn` + "`" + `. Default is no Proxy. Health Check (` + "`" + `healthcheck` + "`" + `) supports the following:`,
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
					Description: `Current state of the load balancer. Usually ` + "`" + `creating` + "`" + ` or ` + "`" + `active` + "`" + ` ## Import Load Balancers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_load_balancer.mylba lba-12345 ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_load_balancer` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Creating Load Balancers - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Deleting Load Balancers`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current state of the load balancer. Usually ` + "`" + `creating` + "`" + ` or ` + "`" + `active` + "`" + ` ## Import Load Balancers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_load_balancer.mylba lba-12345 ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_load_balancer` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Creating Load Balancers - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Deleting Load Balancers`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_orbit_container",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) The Server image ID. One of image or volume must be specified.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Optional) The volume to be used to boot the server. One of image or volume must be specified.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Server name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The handle the server type required (` + "`" + `1gb.ssd` + "`" + `, etc), or a Server Type ID.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The handle of the zone required (` + "`" + `gb1-a` + "`" + `, ` + "`" + `gb1-b` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `(Optional) Set to true to stop the server from being deleted`,
				},
				resource.Attribute{
					Name:        "disk_encrypted",
					Description: `(Optional) Create a server where the data on disk is 'encrypted as rest' by the cloud.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) The desired size of the disk storage for the Server. Only usable with types using network block storage.`,
				},
				resource.Attribute{
					Name:        "snapshots_retention",
					Description: `(Optional) Keep this number of scheduled snapshots. Keep all if unset.`,
				},
				resource.Attribute{
					Name:        "snapshots_schedule",
					Description: `(Optional) Crontab pattern for scheduled snapshots. Must be no more frequent than hourly.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Server`,
				},
				resource.Attribute{
					Name:        "data_volumes",
					Description: `List of data volumes attached to server.`,
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
					Name:        "status",
					Description: `Current state of the server, usually ` + "`" + `active` + "`" + `, ` + "`" + `inactive` + "`" + ` or ` + "`" + `deleted` + "`" + ``,
				},
				resource.Attribute{
					Name:        "snapshot_schedule_next_at",
					Description: `Time in UTC of approximately when the next scheduled snapshot will run.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username used to log onto the server ## Import Servers can be imported using the server ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_server.myserver srv-ojy3o ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_server` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Creating Servers - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Deleting Servers`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Server`,
				},
				resource.Attribute{
					Name:        "data_volumes",
					Description: `List of data volumes attached to server.`,
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
					Name:        "status",
					Description: `Current state of the server, usually ` + "`" + `active` + "`" + `, ` + "`" + `inactive` + "`" + ` or ` + "`" + `deleted` + "`" + ``,
				},
				resource.Attribute{
					Name:        "snapshot_schedule_next_at",
					Description: `Time in UTC of approximately when the next scheduled snapshot will run.`,
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
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
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
					Description: `The ID of the Server`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Is this the default server group?`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The Fully Qualified Domain Name of the server group`,
				},
				resource.Attribute{
					Name:        "firewall_policy",
					Description: `The ID of the Firewall Policy associated with this group ## Import Server Groups can be imported using the server group ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_server_group.default grp-ok8vw ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Server`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Is this the default server group?`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The Fully Qualified Domain Name of the server group`,
				},
				resource.Attribute{
					Name:        "firewall_policy",
					Description: `The ID of the Firewall Policy associated with this group ## Import Server Groups can be imported using the server group ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_server_group.default grp-ok8vw ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_server_group_membership",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group",
					Description: `(Required) The name of the [Server Group.][1]`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `(Required) A list of [Servers][2] to add to the group. ## Attributes Reference No additional attributes are exported. [1]: server_group [2]: server ## Import Server group membership can be imported using the group id and server ids separated by ` + "`" + `/` + "`" + `. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import brightbox_server_group_membership.example1 grp-12345/srv-abcde/srv-fghij ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_volume",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A label assigned to the Volume`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Verbose Description of this volume`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional) True if the volume is encrypted`,
				},
				resource.Attribute{
					Name:        "filesystem_label",
					Description: `(Optional) Label given to the filesystem on the volume. Up to 12 characters.`,
				},
				resource.Attribute{
					Name:        "filesystem_type",
					Description: `(Optional) Format of the filesystem on the volume. Either ` + "`" + `ext4` + "`" + ` or ` + "`" + `xfs` + "`" + `. One of ` + "`" + `image` + "`" + `, ` + "`" + `filesystem_type` + "`" + ` or ` + "`" + `source` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) Image used to create the volume. One of ` + "`" + `image` + "`" + `, ` + "`" + `filesystem_type` + "`" + ` or ` + "`" + `source` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Optional) Volume Serial Number. Up to 20 characters.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Optional) The ID of the server this volume should be attached to.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Disk size in megabytes`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) The ID of the source volume for this image. Defaults to the blank disk. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Volume`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current state of the volume`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `Source type for this image. One of ` + "`" + `image` + "`" + `, ` + "`" + `volume` + "`" + ` or ` + "`" + `raw` + "`" + ``,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Storage type for this volume. Either ` + "`" + `local` + "`" + ` or ` + "`" + `network` + "`" + ` ## Import Volumes can be imported using the volume ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_volume.default vol-ok8vw ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_volume` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Creating Volumes - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Deleting Volumes`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Volume`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current state of the volume`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `Source type for this image. One of ` + "`" + `image` + "`" + `, ` + "`" + `volume` + "`" + ` or ` + "`" + `raw` + "`" + ``,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Storage type for this volume. Either ` + "`" + `local` + "`" + ` or ` + "`" + `network` + "`" + ` ## Import Volumes can be imported using the volume ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import brightbox_volume.default vol-ok8vw ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `brightbox_volume` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Creating Volumes - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `5 minutes` + "`" + `) Used for Deleting Volumes`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"brightbox_api_client":              0,
		"brightbox_cloudip":                 1,
		"brightbox_config_map":              2,
		"brightbox_database_server":         3,
		"brightbox_firewall_policy":         4,
		"brightbox_firewall_rule":           5,
		"brightbox_load_balancer":           6,
		"brightbox_orbit_container":         7,
		"brightbox_server":                  8,
		"brightbox_server_group":            9,
		"brightbox_server_group_membership": 10,
		"brightbox_volume":                  11,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
