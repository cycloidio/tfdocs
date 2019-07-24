package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "bigip_cm_device",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip device`,
			Description:      ``,
			Keywords: []string{
				"cm",
				"device",
			},
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_cm_devicegroup",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip device_name`,
			Description:      ``,
			Keywords: []string{
				"cm",
				"devicegroup",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "bigip_cm_devicegroup",
					Description: `Is the resource used to configure new device group on the BIG-IP.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Is the name of the device Group`,
				},
				resource.Attribute{
					Name:        "auto_sync",
					Description: `Specifies if the device-group will automatically sync configuration data to its members`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies if the device-group will be used for failover or resource syncing`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Name of the device to be included in device group, this need to be configured before using devicegroup resource`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_datagroup",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_datagroup resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"datagroup",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the datagroup`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) datagroup type (applies to the ` + "`" + `name` + "`" + ` field of the record), supports: ` + "`" + `string` + "`" + `, ` + "`" + `ip` + "`" + ` or ` + "`" + `integer` + "`" + ``,
				},
				resource.Attribute{
					Name:        "record",
					Description: `(Optional) a set of ` + "`" + `name` + "`" + ` and ` + "`" + `data` + "`" + ` attributes, name must be of type specified by the ` + "`" + `type` + "`" + ` attributed (` + "`" + `string` + "`" + `, ` + "`" + `ip` + "`" + ` and ` + "`" + `integer` + "`" + `), data is optional and can take any value, multiple ` + "`" + `record` + "`" + ` sets can be specified as needed.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required if ` + "`" + `record` + "`" + ` defined), sets the value of the record's ` + "`" + `name` + "`" + ` attribute, must be of type defined in ` + "`" + `type` + "`" + ` attribute`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Optional if ` + "`" + `record` + "`" + ` defined), sets the value of the record's ` + "`" + `data` + "`" + ` attribute, specifying a value here will create a record in the form of ` + "`" + `name := data` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_irule",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_irule resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"irule",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the iRule`,
				},
				resource.Attribute{
					Name:        "irule",
					Description: `(Required) Body of the iRule`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_monitor",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_monitor resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"monitor",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "parent",
					Description: `(Required) Existing LTM monitor to inherit from`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Check interval in seconds`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout in seconds`,
				},
				resource.Attribute{
					Name:        "send",
					Description: `(Optional) Request string to send`,
				},
				resource.Attribute{
					Name:        "receive",
					Description: `(Optional) Expected response string`,
				},
				resource.Attribute{
					Name:        "receive_disable",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "transparent",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "manual_resume",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ip_dscp",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "time_until_up",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Optional) Specify an alias address for monitoring`,
				},
				resource.Attribute{
					Name:        "compatibility",
					Description: `(Optional) Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `(Optional) Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_node",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_node resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"node",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the node`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) IP or hostname of the node`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `(Optional) Specifies the maximum number of connections allowed for the node or node address.`,
				},
				resource.Attribute{
					Name:        "dynamic_ratio",
					Description: `(Optional) Specifies the fixed ratio value used for a node during ratio load balancing.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Optional) specifies the name of the monitor or monitor rule that you want to associate with the node.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Default is "user-up" you can set to "user-down" if you want to disable`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Specifies the amount of time before sending the next DNS query. Default is 3600. This needs to be specified inside the fqdn (fully qualified domain name).`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `(Optional) Specifies the node's address family. The default is 'unspecified', or IP-agnostic. This needs to be specified inside the fqdn (fully qualified domain name).`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_persistence_profile_cookie",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_persistence_profile_cookie resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"persistence",
				"profile",
				"cookie",
			},
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_persistence_profile_dstaddr",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_persistence_profile_dstaddr resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"persistence",
				"profile",
				"dstaddr",
			},
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_persistence_profile_srcaddr",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_persistence_profile_srcaddr resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"persistence",
				"profile",
				"srcaddr",
			},
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_persistence_profile_ssl",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_persistence_profile_ssl resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"persistence",
				"profile",
				"ssl",
			},
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_policy",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_policy resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"policy",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional) Specifies the match strategy`,
				},
				resource.Attribute{
					Name:        "requires",
					Description: `(Optional) Specifies the protocol`,
				},
				resource.Attribute{
					Name:        "published_copy",
					Description: `(Optional) If you want to publish the policy else it will be deployed in Drafts mode.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Rules can be applied using the policy`,
				},
				resource.Attribute{
					Name:        "tm_name",
					Description: `(Required) If Rule is used then you need to provide the tm_name it can be any value`,
				},
				resource.Attribute{
					Name:        "forward",
					Description: `(Optional) This action will affect forwarding.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `(Optional ) This action will direct the stream to this pool.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_pool",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_pool resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"pool",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the pool`,
				},
				resource.Attribute{
					Name:        "monitors",
					Description: `(Optional) List of monitor names to associate with the pool`,
				},
				resource.Attribute{
					Name:        "allow_nat",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "allow_snat",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "load_balancing_mode",
					Description: `(Optional, Default = round-robin)`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_pool_attachment",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_pool_attachment resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"pool",
				"attachment",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "pool",
					Description: `(Required) Name of the pool in /Partition/Name format`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `(Required) Node to add to the pool in /Partition/NodeName:Port format (e.g. /Common/Node01:80)`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_profile_fasthttp",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_profile_fasthttp resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"profile",
				"fasthttp",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Optional) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.`,
				},
				resource.Attribute{
					Name:        "connpoolidle_timeoutoverride",
					Description: `(Optional) Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.`,
				},
				resource.Attribute{
					Name:        "connpool_maxreuse",
					Description: `(Optional) Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).`,
				},
				resource.Attribute{
					Name:        "connpool_maxsize",
					Description: `(Optional) Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.`,
				},
				resource.Attribute{
					Name:        "connpool_replenish",
					Description: `(Optional) The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional) Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.`,
				},
				resource.Attribute{
					Name:        "connpool_minsize",
					Description: `(Optional) Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.`,
				},
				resource.Attribute{
					Name:        "forcehttp_10response",
					Description: `(Optional) Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.`,
				},
				resource.Attribute{
					Name:        "maxheader_size",
					Description: `(Optional) Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_profile_fastl4",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_profile_fastl4 resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"profile",
				"fastl4",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Optional) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) Displays the administrative partition within which this profile resides`,
				},
				resource.Attribute{
					Name:        "client_timeout",
					Description: `(Optional) Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.`,
				},
				resource.Attribute{
					Name:        "explicitflow_migration",
					Description: `(Optional) Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.`,
				},
				resource.Attribute{
					Name:        "hardware_syncookie",
					Description: `(Optional) Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional) Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.`,
				},
				resource.Attribute{
					Name:        "iptos_toclient",
					Description: `(Optional) Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.`,
				},
				resource.Attribute{
					Name:        "keepalive_interval",
					Description: `(Optional) Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_profile_http",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_profile_http resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"profile",
				"http",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Required) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.`,
				},
				resource.Attribute{
					Name:        "fallback_host",
					Description: `(Optional) Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number`,
				},
				resource.Attribute{
					Name:        "fallback_status_codes",
					Description: `(Optional) Specifies one or more three-digit status codes that can be returned by an HTTP server.`,
				},
				resource.Attribute{
					Name:        "oneconnect_transformations",
					Description: `(Optional) Enables the system to perform HTTP header transformations for the purpose of keeping server-side connections open. This feature requires configuration of a OneConnect profile`,
				},
				resource.Attribute{
					Name:        "basic_auth_realm",
					Description: `(Optional) Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none`,
				},
				resource.Attribute{
					Name:        "head_insert",
					Description: `(Optional) Specifies a quoted header string that you want to insert into an HTTP request`,
				},
				resource.Attribute{
					Name:        "insert_xforwarded_for",
					Description: `(Optional) When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_profile_http2",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_profile_http2 resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"profile",
				"http2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Required) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.`,
				},
				resource.Attribute{
					Name:        "concurrent_streams_per_connection",
					Description: `(Optional) Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.`,
				},
				resource.Attribute{
					Name:        "connection_idle_timeout",
					Description: `(Optional) Specifies the number of seconds that a connection is idle before the connection is eligible for deletion..`,
				},
				resource.Attribute{
					Name:        "connpool_maxsize",
					Description: `(Optional) Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.`,
				},
				resource.Attribute{
					Name:        "activation_modes",
					Description: `(Optional) Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default values npn and alpn specify that the TLS next-protocol-negotiation and application-layer-protocol-negotiation extensions will be used.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_profile_httpcompress",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_profile_httpcompress resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"profile",
				"httpcompress",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Optional) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.`,
				},
				resource.Attribute{
					Name:        "content_type_include",
					Description: `(Optional) Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.`,
				},
				resource.Attribute{
					Name:        "content_type_exclude",
					Description: `(Optional) Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_profile_oneconnect",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_profile_oneconnect resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"profile",
				"oneconnect",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) Displays the administrative partition within which this profile resides`,
				},
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Optional) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.`,
				},
				resource.Attribute{
					Name:        "idle_timeout_override",
					Description: `(Optional) Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are disabled, indefinite, or a numeric value that you specify. The default value is disabled.`,
				},
				resource.Attribute{
					Name:        "share_pools",
					Description: `(Optional) Specify if you want to share the pool, default value is "disabled"`,
				},
				resource.Attribute{
					Name:        "max_age",
					Description: `(Optional) Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is 86400.`,
				},
				resource.Attribute{
					Name:        "max_reuse",
					Description: `(Optional) Specifies the maximum number of times that a server-side connection can be reused. The default value is 1000.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Optional) Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is 10000.`,
				},
				resource.Attribute{
					Name:        "source_mask",
					Description: `(Optional) Specifies a source IP mask. The default value is 0.0.0.0. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_profile_tcp",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_profile_tcp resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"profile",
				"tcp",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) Displays the administrative partition within which this profile resides`,
				},
				resource.Attribute{
					Name:        "defaults_from",
					Description: `(Optional) Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional) Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.`,
				},
				resource.Attribute{
					Name:        "close_wait_timeout",
					Description: `(Optional) Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.`,
				},
				resource.Attribute{
					Name:        "finwait_timeout",
					Description: `(Optional) Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.`,
				},
				resource.Attribute{
					Name:        "finwait_2timeout",
					Description: `(Optional) Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).`,
				},
				resource.Attribute{
					Name:        "keepalive_interval",
					Description: `(Optional) Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.`,
				},
				resource.Attribute{
					Name:        "fast_open",
					Description: `(Optional) When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet.`,
				},
				resource.Attribute{
					Name:        "deferred_accept",
					Description: `(Optional) Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_snat",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_snat resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"snat",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the snat`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) Displays the administrative partition within which this profile resides`,
				},
				resource.Attribute{
					Name:        "origins",
					Description: `(Optional) IP or hostname of the snat`,
				},
				resource.Attribute{
					Name:        "snatpool",
					Description: `(Optional) Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.`,
				},
				resource.Attribute{
					Name:        "mirror",
					Description: `(Optional) Enables or disables mirroring of SNAT connections.`,
				},
				resource.Attribute{
					Name:        "sourceport",
					Description: `(Optional) Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.`,
				},
				resource.Attribute{
					Name:        "translation",
					Description: `(Optional) Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.`,
				},
				resource.Attribute{
					Name:        "vlansdisabled",
					Description: `(Optional) Disables the SNAT on all VLANs.`,
				},
				resource.Attribute{
					Name:        "vlans",
					Description: `(Optional) Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_snatpool",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_snatpool resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"snatpool",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the snatpool`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_virtual_address",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_virtual_address resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"virtual",
				"address",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the virtual address`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the virtual address`,
				},
				resource.Attribute{
					Name:        "advertize_route",
					Description: `(Optional) Enabled dynamic routing of the address`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `(Optional, Default=0) Max number of connections for virtual address`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, Default=true) Enable or disable the virtual address`,
				},
				resource.Attribute{
					Name:        "arp",
					Description: `(Optional, Default=true) Enable or disable ARP for the virtual address`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `(Optional, Default=true) Automatically delete the virtual address with the virtual server`,
				},
				resource.Attribute{
					Name:        "icmp_echo",
					Description: `(Optional, Default=true) Enable/Disable ICMP response to the virtual address`,
				},
				resource.Attribute{
					Name:        "traffic_group",
					Description: `(Optional, Default=/Common/traffic-group-1) Specify the partition and traffic group`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_virtual_server",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_ltm_virtual_server resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"virtual",
				"server",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Listen port for the virtual server`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Destination IP`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `(Optional) Default pool name`,
				},
				resource.Attribute{
					Name:        "mask",
					Description: `(Optional) Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0`,
				},
				resource.Attribute{
					Name:        "source_address_translation",
					Description: `(Optional) Can be either omitted for none or the values automap or snat`,
				},
				resource.Attribute{
					Name:        "translate_address",
					Description: `Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.`,
				},
				resource.Attribute{
					Name:        "translate_port",
					Description: `Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service`,
				},
				resource.Attribute{
					Name:        "profiles",
					Description: `(Optional) List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.`,
				},
				resource.Attribute{
					Name:        "client_profiles",
					Description: `(Optional) List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles`,
				},
				resource.Attribute{
					Name:        "server_profiles",
					Description: `(Optional) List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Specifies an IP address or network from which the virtual server will accept traffic.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Optional) The iRules you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.`,
				},
				resource.Attribute{
					Name:        "snatpool",
					Description: `(Optional) Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation`,
				},
				resource.Attribute{
					Name:        "vlans",
					Description: `(Optional) The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.`,
				},
				resource.Attribute{
					Name:        "vlans_enabled",
					Description: `(Optional Bool) Enables the virtual server on the VLANs specified by the VLANs option.`,
				},
				resource.Attribute{
					Name:        "vlans_disabled",
					Description: `(Optional Bool) Disables the virtual server on the VLANs specified by the VLANs option.`,
				},
				resource.Attribute{
					Name:        "persistence_profiles",
					Description: `(Optional) List of persistence profiles associated with the Virtual Server.`,
				},
				resource.Attribute{
					Name:        "fallback_persistence_profile",
					Description: `(Optional) Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_net_route",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_net_route resource`,
			Description:      ``,
			Keywords: []string{
				"net",
				"route",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the route`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) The destination subnet and netmask for the route.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) Specifies a gateway address for the route.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_net_selfip",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_net_selfip resource`,
			Description:      ``,
			Keywords: []string{
				"net",
				"selfip",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the selfip`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The Self IP's address and netmask.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Required) Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.`,
				},
				resource.Attribute{
					Name:        "traffic_group",
					Description: `(Optional) Specifies the traffic group, defaults to ` + "`" + `traffic-group-local-only` + "`" + ` if not specified.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_net_vlan",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_net_vlan resource`,
			Description:      ``,
			Keywords: []string{
				"net",
				"vlan",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the vlan`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Specifies a number that the system adds into the header of any frame passing through the VLAN.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Specifies which interfaces you want this VLAN to use for traffic management.`,
				},
				resource.Attribute{
					Name:        "vlanport",
					Description: `Physical or virtual port used for traffic`,
				},
				resource.Attribute{
					Name:        "tagged",
					Description: `Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_ltm_dns",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip_sys_dns resource`,
			Description:      ``,
			Keywords: []string{
				"ltm",
				"dns",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name_servers",
					Description: `Name or IP address of the DNS server`,
				},
				resource.Attribute{
					Name:        "number_of_dots",
					Description: `Configures the number of dots needed in a name before an initial absolute query will be made.`,
				},
				resource.Attribute{
					Name:        "search",
					Description: `Specify what domains you want to search`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_sys_iapp",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip iapp resource for BIG-IP`,
			Description:      ``,
			Keywords: []string{
				"sys",
				"iapp",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the iApp.`,
				},
				resource.Attribute{
					Name:        "jsonfile",
					Description: `Refer to the Json file which will be deployed on F5 BIG-IP. ## Example Usage of Json file ` + "`" + `` + "`" + `` + "`" + ` { "fullPath":"/Common/simplehttp.app/simplehttp", "generation":222, "inheritedDevicegroup":"true", "inheritedTrafficGroup":"true", "kind":"tm:sys:application:service:servicestate", "name":"simplehttp", "partition":"Common", "selfLink":"https://localhost/mgmt/tm/sys/application/service/~Common~simplehttp.app~simplehttp?ver=13.0.0", "strictUpdates":"enabled", "subPath":"simplehttp.app", "tables":[ { "name":"basic__snatpool_members" }, { "name":"net__snatpool_members" }, { "name":"optimizations__hosts" }, { "columnNames":[ "name" ], "name":"pool__hosts", "rows":[ { "row":[ "f5.cisco.com" ] } ] }, { "columnNames":[ "addr", "port", "connection_limit" ], "name":"pool__members", "rows":[ { "row":[ "10.0.2.167", "80", "0" ] }, { "row":[ "10.0.2.168", "80", "0" ] } ] }, { "name":"server_pools__servers" } ], "template":"/Common/f5.http", "templateModified":"no", "templateReference":{ "link":"https://localhost/mgmt/tm/sys/application/template/~Common~f5.http?ver=13.0.0" }, "trafficGroup":"/Common/traffic-group-1", "trafficGroupReference":{ "link":"https://localhost/mgmt/tm/cm/traffic-group/~Common~traffic-group-1?ver=13.0.0" }, "variables":[ { "encrypted":"no", "name":"client__http_compression", "value":"/#create_new#" }, { "encrypted":"no", "name":"monitor__monitor", "value":"/Common/http" }, { "encrypted":"no", "name":"net__client_mode", "value":"wan" }, { "encrypted":"no", "name":"net__server_mode", "value":"lan" }, { "encrypted":"no", "name":"net__v13_tcp", "value":"warn" }, { "encrypted":"no", "name":"pool__addr", "value":"10.0.1.100" }, { "encrypted":"no", "name":"pool__pool_to_use", "value":"/#create_new#" }, { "encrypted":"no", "name":"pool__port", "value":"80" }, { "encrypted":"no", "name":"ssl__mode", "value":"no_ssl" }, { "encrypted":"no", "name":"ssl_encryption_questions__advanced", "value":"no" }, { "encrypted":"no", "name":"ssl_encryption_questions__help", "value":"hide" } ] } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `User defined description.`,
				},
				resource.Attribute{
					Name:        "deviceGroup",
					Description: `The name of the device group that the application service is assigned to.`,
				},
				resource.Attribute{
					Name:        "executeAction",
					Description: `Run the specified template action associated with the application.`,
				},
				resource.Attribute{
					Name:        "inheritedTrafficGroup",
					Description: `Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `Displays the administrative partition within which the application resides.`,
				},
				resource.Attribute{
					Name:        "strictUpdates",
					Description: `Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.`,
				},
				resource.Attribute{
					Name:        "templateModified",
					Description: `Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.`,
				},
				resource.Attribute{
					Name:        "templatePrerequisiteErrors",
					Description: `Indicates any missing prerequisites associated with the template that defines this application.`,
				},
				resource.Attribute{
					Name:        "trafficGroup",
					Description: `The name of the traffic group that the application service is assigned to.`,
				},
				resource.Attribute{
					Name:        "lists",
					Description: `string values`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `User defined generic data for the application service. It is a name and value pair.`,
				},
				resource.Attribute{
					Name:        "tables",
					Description: `Values provided like pool name, nodes etc.`,
				},
				resource.Attribute{
					Name:        "variables",
					Description: `Name, values, encrypted or not`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_sys_ntp",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip ntp`,
			Description:      ``,
			Keywords: []string{
				"sys",
				"ntp",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "bigip_sys_ntp",
					Description: `Is the resource is used to configure ntp server on the BIG-IP.`,
				},
				resource.Attribute{
					Name:        "/Common/NTP1",
					Description: `Is the description of the NTP server in the main or common partition of BIG-IP.`,
				},
				resource.Attribute{
					Name:        "time.facebook.com",
					Description: `Is the NTP server configured on the BIG-IP.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `(Optional) Adds NTP servers to or deletes NTP servers from the BIG-IP system.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) Specifies the time zone that you want to use for the system time.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_sys_provision",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip provision resource for BIG-IP`,
			Description:      ``,
			Keywords: []string{
				"sys",
				"provision",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "bigip_sys_provision",
					Description: `Is the resource which is used to provision big-ip modules like asm, afm, ilx etc`,
				},
				resource.Attribute{
					Name:        "Common/ilx",
					Description: `Common is the partition and ilx is the module being enabled it could be asm, afm apm etc.`,
				},
				resource.Attribute{
					Name:        "cpuRatio",
					Description: `how much cpu resources you need for this resource`,
				},
				resource.Attribute{
					Name:        "diskRatio",
					Description: `how much disk space you want to allocate for this resource.`,
				},
				resource.Attribute{
					Name:        "memoryRatio",
					Description: `how much memory you want to deidcate for this resource`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_sys_snmp",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip snmp resource for BIG-IP`,
			Description:      ``,
			Keywords: []string{
				"sys",
				"snmp",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "sys_contact",
					Description: `(Optional) Specifies the contact information for the system administrator.`,
				},
				resource.Attribute{
					Name:        "sys_location",
					Description: `Describes the system's physical location.`,
				},
				resource.Attribute{
					Name:        "allowedaddresses",
					Description: `Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bigip_sys_snmp_traps",
			Category:         "Resources",
			ShortDescription: `Provides details about bigip snmp_traps resource for BIG-IP`,
			Description:      ``,
			Keywords: []string{
				"sys",
				"snmp",
				"traps",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the snmp trap.`,
				},
				resource.Attribute{
					Name:        "community",
					Description: `(Optional) Specifies the community string used for this trap.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host the trap will be sent to.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The port that the trap will be sent to.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) User defined description.`,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"bigip_cm_device":                       0,
		"bigip_cm_devicegroup":                  1,
		"bigip_ltm_datagroup":                   2,
		"bigip_ltm_irule":                       3,
		"bigip_ltm_monitor":                     4,
		"bigip_ltm_node":                        5,
		"bigip_ltm_persistence_profile_cookie":  6,
		"bigip_ltm_persistence_profile_dstaddr": 7,
		"bigip_ltm_persistence_profile_srcaddr": 8,
		"bigip_ltm_persistence_profile_ssl":     9,
		"bigip_ltm_policy":                      10,
		"bigip_ltm_pool":                        11,
		"bigip_ltm_pool_attachment":             12,
		"bigip_ltm_profile_fasthttp":            13,
		"bigip_ltm_profile_fastl4":              14,
		"bigip_ltm_profile_http":                15,
		"bigip_ltm_profile_http2":               16,
		"bigip_ltm_profile_httpcompress":        17,
		"bigip_ltm_profile_oneconnect":          18,
		"bigip_ltm_profile_tcp":                 19,
		"bigip_ltm_snat":                        20,
		"bigip_ltm_snatpool":                    21,
		"bigip_ltm_virtual_address":             22,
		"bigip_ltm_virtual_server":              23,
		"bigip_net_route":                       24,
		"bigip_net_selfip":                      25,
		"bigip_net_vlan":                        26,
		"bigip_ltm_dns":                         27,
		"bigip_sys_iapp":                        28,
		"bigip_sys_ntp":                         29,
		"bigip_sys_provision":                   30,
		"bigip_sys_snmp":                        31,
		"bigip_sys_snmp_traps":                  32,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
