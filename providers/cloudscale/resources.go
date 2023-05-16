package cloudscale

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_custom_image",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "import_url",
					Description: `(Required) The URL used to download the image.`,
				},
				resource.Attribute{
					Name:        "import_source_format",
					Description: `(Optional, Ignored) Deprecated: this field no longer needs to be specified and will be ignored. The image format is detected automatically.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of the custom image.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Optional) A string identifying the custom image for use within the API.`,
				},
				resource.Attribute{
					Name:        "user_data_handling",
					Description: `(Required) How user_data will be handled when creating a server. Options include ` + "`" + `pass-through` + "`" + ` and ` + "`" + `extend-cloud-config` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "firmware_type",
					Description: `(Optional) The firmware type that will be used for servers created with the custom image. Options include ` + "`" + `bios` + "`" + ` and ` + "`" + `uefi` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_slugs",
					Description: `(Required) Specify the zones in which the custom image will be available. Options include ` + "`" + `lpg1` + "`" + ` and ` + "`" + `rma1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timeouts",
					Description: `(Optional) Specify how long certain operations are allowed to take before being considered to have failed. Currently, only the ` + "`" + `create` + "`" + ` timeout can be specified. Takes a string representation of a duration, such as ` + "`" + `20m` + "`" + ` for 20 minutes (default), ` + "`" + `10s` + "`" + ` for ten seconds, or ` + "`" + `2h` + "`" + ` for two hours.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags allow you to assign custom metadata to resources: ` + "`" + `` + "`" + `` + "`" + ` tags = { foo = "bar" } ` + "`" + `` + "`" + `` + "`" + ` Tags are always strings (both keys and values). ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size in GB of the custom image.`,
				},
				resource.Attribute{
					Name:        "checksums",
					Description: `The checksums of the custom image as map.`,
				},
				resource.Attribute{
					Name:        "import_href",
					Description: `The cloudscale.ch API URL of the custom image import.`,
				},
				resource.Attribute{
					Name:        "import_uuid",
					Description: `The UUID of the custom image import.`,
				},
				resource.Attribute{
					Name:        "import_status",
					Description: `The status of the custom image import. Options include ` + "`" + `started` + "`" + `, ` + "`" + `in_progress` + "`" + `, ` + "`" + `failed` + "`" + `, ` + "`" + `success` + "`" + `. ## Import Custom images can currently not be imported, please use a data source.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size in GB of the custom image.`,
				},
				resource.Attribute{
					Name:        "checksums",
					Description: `The checksums of the custom image as map.`,
				},
				resource.Attribute{
					Name:        "import_href",
					Description: `The cloudscale.ch API URL of the custom image import.`,
				},
				resource.Attribute{
					Name:        "import_uuid",
					Description: `The UUID of the custom image import.`,
				},
				resource.Attribute{
					Name:        "import_status",
					Description: `The status of the custom image import. Options include ` + "`" + `started` + "`" + `, ` + "`" + `in_progress` + "`" + `, ` + "`" + `failed` + "`" + `, ` + "`" + `success` + "`" + `. ## Import Custom images can currently not be imported, please use a data source.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_floating_ip",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server",
					Description: `(Optional) Assign the Floating IP to this server (UUID).`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `(Optional) Assign the Floating IP to this load balancer (UUID).`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Required) ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `, for an IPv4 or IPv6 address or network respectively.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Optional) If you want to assign an entire network instead of a single IP address to your server, you must specify the prefix length. Currently, there is only support for ` + "`" + `ip_version=6` + "`" + ` and ` + "`" + `prefix_length=56` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) You can specify the type. Options include ` + "`" + `regional` + "`" + ` (default) and ` + "`" + `global` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region_slug",
					Description: `(Optional) The slug of the region in which the new Regional Floating IP will be created. Options include ` + "`" + `lpg` + "`" + ` and ` + "`" + `rma` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reverse_ptr",
					Description: `(Optional) You can specify the PTR record (reverse DNS pointer) in case of a single Floating IP address.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags allow you to assign custom metadata to resources: ` + "`" + `` + "`" + `` + "`" + ` tags = { foo = "bar" } ` + "`" + `` + "`" + `` + "`" + ` Tags are always strings (both keys and values). The following arguments are supported when updating Floating IPs:`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Optional) (Re-)Assign the Floating IP to this server (UUID).`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `(Optional) (Re-)Assign the Floating IP to this load balancer (UUID).`,
				},
				resource.Attribute{
					Name:        "reverse_ptr",
					Description: `(Optional) You can specify the new PTR record (reverse DNS pointer) in case of a single Floating IP address.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Change tags (see documentation above) ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The CIDR notation of the Floating IP address or network, e.g. ` + "`" + `192.0.2.123/32` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `The IP address of the server or load balancer that your Floating IP is currently assigned to. ## Import Floating IPs can be imported using the Floating IP's network IP: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_floating_ip.floating_ip 192.0.2.24 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The CIDR notation of the Floating IP address or network, e.g. ` + "`" + `192.0.2.123/32` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `The IP address of the server or load balancer that your Floating IP is currently assigned to. ## Import Floating IPs can be imported using the Floating IP's network IP: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_floating_ip.floating_ip 192.0.2.24 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_load_balancer",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the new load balancer.`,
				},
				resource.Attribute{
					Name:        "flavor_slug",
					Description: `(Required) The slug (name) of the flavor to use for the new load balancer. Possible values can be found in our [API documentation](https://www.cloudscale.ch/en/api/v1#load-balancer-flavors).`,
				},
				resource.Attribute{
					Name:        "zone_slug",
					Description: `(Required) The slug of the zone in which the new load balancer will be created. Options include ` + "`" + `lpg1` + "`" + ` and ` + "`" + `rma1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vip_addresses",
					Description: `(Optional) A list of VIP address objects. This attributes needs to be specified if the load balancer should be assigned a VIP address in a subnet on a private network. If the VIP address should be created on the public network, this attribute should be omitted. Each VIP address object has the following attributes:`,
				},
				resource.Attribute{
					Name:        "subnet_uuid",
					Description: `(Optional) The UUID of the subnet this VIP address should be part of.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) An VIP address that has been assigned to this load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags allow you to assign custom metadata to resources: ` + "`" + `` + "`" + `` + "`" + ` tags = { foo = "bar" } ` + "`" + `` + "`" + `` + "`" + ` Tags are always strings (both keys and values). The following arguments are supported when updating load balancers:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `New name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Change tags (see documentation above)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this load balancer.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the load balancer.`,
				},
				resource.Attribute{
					Name:        "vip_addresses",
					Description: `A list of VIP address objects. Each VIP address object has the following attributes:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The IP version, either ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `The cidr of the subnet the VIP address is part of.`,
				},
				resource.Attribute{
					Name:        "subnet_href",
					Description: `The cloudscale.ch API URL of the subnet the VIP address is part of. ## Import Load balancer can be imported using the load balancer's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_load_balancer.lb 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this load balancer.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the load balancer.`,
				},
				resource.Attribute{
					Name:        "vip_addresses",
					Description: `A list of VIP address objects. Each VIP address object has the following attributes:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The IP version, either ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `The cidr of the subnet the VIP address is part of.`,
				},
				resource.Attribute{
					Name:        "subnet_href",
					Description: `The cloudscale.ch API URL of the subnet the VIP address is part of. ## Import Load balancer can be imported using the load balancer's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_load_balancer.lb 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_load_balancer_health_monitor",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the new load balancer health monitor.`,
				},
				resource.Attribute{
					Name:        "pool_uuid",
					Description: `(Required) The pool of the health monitor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the health monitor. Options include: ` + "`" + `"ping"` + "`" + `, ` + "`" + `"tcp"` + "`" + `, ` + "`" + `"http"` + "`" + `, ` + "`" + `"https"` + "`" + ` and ` + "`" + `"tls-hello"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delay_s",
					Description: `(Optional) The delay between two successive checks in seconds. Default is ` + "`" + `2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timeout_s",
					Description: `(Optional) The maximum time allowed for an individual check in seconds. Default is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "up_threshold",
					Description: `(Optional) The number of checks that need to be successful before the ` + "`" + `monitor_status` + "`" + ` of a pool member changes to ` + "`" + `"up"` + "`" + `. Default is ` + "`" + `2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "down_threshold",
					Description: `(Optional) The number of checks that need to fail before the ` + "`" + `monitor_status` + "`" + ` of a pool member changes to ` + "`" + `"down"` + "`" + `. Default is ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_expected_codes",
					Description: `(Optional) The HTTP status codes allowed for a check to be considered successful. Can either be a list of status codes, for example ` + "`" + `["200", "202"]` + "`" + `, or a list containing a single range, for example ` + "`" + `["200-204"]` + "`" + `. Default is ` + "`" + `["200"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `(Optional) The HTTP method used for the check. Options include ` + "`" + `"CONNECT"` + "`" + `, ` + "`" + `"DELETE"` + "`" + `, ` + "`" + `"GET"` + "`" + `, ` + "`" + `"HEAD"` + "`" + `, ` + "`" + `"OPTIONS"` + "`" + `, ` + "`" + `"PATCH"` + "`" + `, ` + "`" + `"POST"` + "`" + `, ` + "`" + `"PUT"` + "`" + ` and ` + "`" + `"TRACE"` + "`" + `. Default is ` + "`" + `"GET"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_url_path",
					Description: `(Optional) The URL used for the check. Default is ` + "`" + `"/"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_version",
					Description: `(Optional) The HTTP version used for the check. Options include ` + "`" + `"1.0"` + "`" + ` and ` + "`" + `"1.1"` + "`" + `. Default is ` + "`" + `"1.1"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_host",
					Description: `(Optional) The server name in the HTTP Host: header used for the check. Requires version to be set to ` + "`" + `"1.1"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags allow you to assign custom metadata to resources: ` + "`" + `` + "`" + `` + "`" + ` tags = { foo = "bar" } ` + "`" + `` + "`" + `` + "`" + ` Tags are always strings (both keys and values). The following arguments are supported when updating load balancer health monitor:`,
				},
				resource.Attribute{
					Name:        "delay_s",
					Description: `The delay between two successive checks in seconds. Default is ` + "`" + `2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timeout_s",
					Description: `The maximum time allowed for an individual check in seconds. Default is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "up_threshold",
					Description: `The number of checks that need to be successful before the ` + "`" + `monitor_status` + "`" + ` of a pool member changes to ` + "`" + `"up"` + "`" + `. Default is ` + "`" + `2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "down_threshold",
					Description: `The number of checks that need to fail before the ` + "`" + `monitor_status` + "`" + ` of a pool member changes to ` + "`" + `"down"` + "`" + `. Default is ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_expected_codes",
					Description: `The HTTP status codes allowed for a check to be considered successful. Can either be a list of status codes, for example ` + "`" + `["200", "202"]` + "`" + `, or a list containing a single range, for example ` + "`" + `["200-204"]` + "`" + `. Default is ` + "`" + `["200"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `The HTTP method used for the check. Options include ` + "`" + `"CONNECT"` + "`" + `, ` + "`" + `"DELETE"` + "`" + `, ` + "`" + `"GET"` + "`" + `, ` + "`" + `"HEAD"` + "`" + `, ` + "`" + `"OPTIONS"` + "`" + `, ` + "`" + `"PATCH"` + "`" + `, ` + "`" + `"POST"` + "`" + `, ` + "`" + `"PUT"` + "`" + ` and ` + "`" + `"TRACE"` + "`" + `. Default is ` + "`" + `"GET"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_url_path",
					Description: `The URL used for the check. Default is ` + "`" + `"/"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_host",
					Description: `The server name in the HTTP Host: header used for the check. Requires version to be set to ` + "`" + `"1.1"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Change tags (see documentation above) ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this load balancer health monitor.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `The load balancer pool name of the health monitor.`,
				},
				resource.Attribute{
					Name:        "pool_href",
					Description: `The cloudscale.ch API URL of the health monitor's load balancer pool. ## Import Load balancer health monitor can be imported using the load balancer health monitor's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_load_balancer_health_monitor.lb1-health-monitor 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this load balancer health monitor.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `The load balancer pool name of the health monitor.`,
				},
				resource.Attribute{
					Name:        "pool_href",
					Description: `The cloudscale.ch API URL of the health monitor's load balancer pool. ## Import Load balancer health monitor can be imported using the load balancer health monitor's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_load_balancer_health_monitor.lb1-health-monitor 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_load_balancer_listener",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the new load balancer listener.`,
				},
				resource.Attribute{
					Name:        "pool_uuid",
					Description: `(Required) The pool of the listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol used for receiving traffic. Options include ` + "`" + `"tcp"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required) The port on which traffic is received.`,
				},
				resource.Attribute{
					Name:        "timeout_client_data_ms",
					Description: `(Optional) Client inactivity timeout in milliseconds.`,
				},
				resource.Attribute{
					Name:        "timeout_member_connect_ms",
					Description: `(Optional) Pool member connection timeout in milliseconds.`,
				},
				resource.Attribute{
					Name:        "timeout_member_data_ms",
					Description: `(Optional) Pool member inactivity timeout in milliseconds.`,
				},
				resource.Attribute{
					Name:        "allowed_cidrs",
					Description: `(Optional) Restrict the allowed source IPs for this listener. ` + "`" + `[]` + "`" + ` means that any source IP is allowed. If the list is non-empty, traffic from source IPs not included is denied.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags allow you to assign custom metadata to resources: ` + "`" + `` + "`" + `` + "`" + ` tags = { foo = "bar" } ` + "`" + `` + "`" + `` + "`" + ` Tags are always strings (both keys and values). The following arguments are supported when updating load balancer listener:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `New name of the load balancer listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol used for receiving traffic. Options include ` + "`" + `"tcp"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `The port on which traffic is received.`,
				},
				resource.Attribute{
					Name:        "timeout_client_data_ms",
					Description: `Client inactivity timeout in milliseconds.`,
				},
				resource.Attribute{
					Name:        "timeout_member_connect_ms",
					Description: `Pool member connection timeout in milliseconds.`,
				},
				resource.Attribute{
					Name:        "timeout_member_data_ms",
					Description: `Pool member inactivity timeout in milliseconds.`,
				},
				resource.Attribute{
					Name:        "allowed_cidrs",
					Description: `Restrict the allowed source IPs for this listener. ` + "`" + `[]` + "`" + ` means that any source IP is allowed. If the list is non-empty, traffic from source IPs not included is denied.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Change tags (see documentation above) ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this load balancer listner.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `The load balancer pool name of the listener.`,
				},
				resource.Attribute{
					Name:        "pool_href",
					Description: `The cloudscale.ch API URL of the listener's load balancer pool. ## Import Load balancer listener can be imported using the load balancer listener's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_load_balancer_listener.lb1-listener 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this load balancer listner.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `The load balancer pool name of the listener.`,
				},
				resource.Attribute{
					Name:        "pool_href",
					Description: `The cloudscale.ch API URL of the listener's load balancer pool. ## Import Load balancer listener can be imported using the load balancer listener's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_load_balancer_listener.lb1-listener 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_load_balancer_pool",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the new load balancer pool.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Required) The algorithm according to which the incoming traffic is distributed between the pool members. Options include ` + "`" + `"round_robin"` + "`" + `, ` + "`" + `"least_connections"` + "`" + ` and ` + "`" + `"source_ip"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol used for traffic between the load balancer and the pool members. Options include: ` + "`" + `"tcp"` + "`" + `, ` + "`" + `"proxy"` + "`" + ` and ` + "`" + `"proxyv2"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "load_balancer_uuid",
					Description: `(Required) The load balancer of the pool.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags allow you to assign custom metadata to resources: ` + "`" + `` + "`" + `` + "`" + ` tags = { foo = "bar" } ` + "`" + `` + "`" + `` + "`" + ` Tags are always strings (both keys and values). The following arguments are supported when updating load balancer pools:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `New name of the load balancer pool.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Change tags (see documentation above) ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this load balancer pool.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The load balancer name of the pool.`,
				},
				resource.Attribute{
					Name:        "load_balancer_href",
					Description: `The cloudscale.ch API URL of the pool's load balancer. ## Import Load balancer pools can be imported using the load balancer pool's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_load_balancer_pool.lb1-pool 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this load balancer pool.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `The load balancer name of the pool.`,
				},
				resource.Attribute{
					Name:        "load_balancer_href",
					Description: `The cloudscale.ch API URL of the pool's load balancer. ## Import Load balancer pools can be imported using the load balancer pool's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_load_balancer_pool.lb1-pool 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_load_balancer_pool_member",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the new load balancer pool member.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Pool member will not receive traffic if ` + "`" + `false` + "`" + `. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pool_uuid",
					Description: `(Required) The load balancer pool of the member.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required) The port to which actual traffic is sent.`,
				},
				resource.Attribute{
					Name:        "monitor_port",
					Description: `(Optional) The port to which health monitor checks are sent. If not specified, ` + "`" + `protocol_port` + "`" + ` will be used.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The IP address to which traffic is sent.`,
				},
				resource.Attribute{
					Name:        "subnet_uuid",
					Description: `(Required) The subnet UUID of the address must be specified here.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags allow you to assign custom metadata to resources: ` + "`" + `` + "`" + `` + "`" + ` tags = { foo = "bar" } ` + "`" + `` + "`" + `` + "`" + ` Tags are always strings (both keys and values). The following arguments are supported when updating load balancer pool members:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `New name of the load balancer pool.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Pool member will not receive traffic if ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Change tags (see documentation above) ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this load balancer pool member.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "monitor_status",
					Description: `The status of the pool's health monitor check for this member. Can be ` + "`" + `"up"` + "`" + `, ` + "`" + `"down"` + "`" + `, ` + "`" + `"changing"` + "`" + `, ` + "`" + `"no_monitor"` + "`" + ` and ` + "`" + `"unknown"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `The load balancer pool name of the member.`,
				},
				resource.Attribute{
					Name:        "pool_href",
					Description: `The cloudscale.ch API URL of the member's load balancer pool.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `The CIDR of the member's address subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_href",
					Description: `The cloudscale.ch API URL of the member's address subnet. ## Import Load balancer pool members can be imported using the load balancer pool member's UUID and the pool UUID using this schema ` + "`" + `{pool_uuid}.{member_uuid}` + "`" + `: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_load_balancer_pool_member.lb1-pool-member 48151623-42aa-aaaa-bbbb-caffeeeeeeee.6a18a377-9977-4cd0-b1fa-70908356efaa ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this load balancer pool member.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "monitor_status",
					Description: `The status of the pool's health monitor check for this member. Can be ` + "`" + `"up"` + "`" + `, ` + "`" + `"down"` + "`" + `, ` + "`" + `"changing"` + "`" + `, ` + "`" + `"no_monitor"` + "`" + ` and ` + "`" + `"unknown"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `The load balancer pool name of the member.`,
				},
				resource.Attribute{
					Name:        "pool_href",
					Description: `The cloudscale.ch API URL of the member's load balancer pool.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `The CIDR of the member's address subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_href",
					Description: `The cloudscale.ch API URL of the member's address subnet. ## Import Load balancer pool members can be imported using the load balancer pool member's UUID and the pool UUID using this schema ` + "`" + `{pool_uuid}.{member_uuid}` + "`" + `: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_load_balancer_pool_member.lb1-pool-member 48151623-42aa-aaaa-bbbb-caffeeeeeeee.6a18a377-9977-4cd0-b1fa-70908356efaa ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_network",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the network.`,
				},
				resource.Attribute{
					Name:        "zone_slug",
					Description: `(Optional) The slug of the zone in which the new network will be created. Options include ` + "`" + `lpg1` + "`" + ` and ` + "`" + `rma1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) You can specify the MTU size for the network, defaults to 9000.`,
				},
				resource.Attribute{
					Name:        "auto_create_ipv4_subnet",
					Description: `(Optional) Automatically create an IPv4 Subnet on the network. Can be ` + "`" + `true` + "`" + ` (default) or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags allow you to assign custom metadata to resources: ` + "`" + `` + "`" + `` + "`" + ` tags = { foo = "bar" } ` + "`" + `` + "`" + `` + "`" + ` Tags are always strings (both keys and values). ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current network.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `A list of subnet objects that are used in this network. Each subnet object has the following attributes:`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The CIDR notation of the subnet.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of this subnet.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of this subnet. ## Import Networks can be imported using the network's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_network.network 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current network.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `A list of subnet objects that are used in this network. Each subnet object has the following attributes:`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The CIDR notation of the subnet.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of this subnet.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of this subnet. ## Import Networks can be imported using the network's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_network.network 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_objects_user",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The display name of the Objects User.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags allow you to assign custom metadata to resources: ` + "`" + `` + "`" + `` + "`" + ` tags = { foo = "bar" } ` + "`" + `` + "`" + `` + "`" + ` Tags are always strings (both keys and values). The following arguments are supported when updating Objects Users:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The new display name of the Objects User.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Change tags (see documentation above) ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The unique identifier of the Objects User.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `A list of key objects containing the access and secret key associated with the Objects User. Currently, only one key object is returned. Each key object has the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `The S3 access key of the Objects User.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `The S3 secret key of the Objects User. ## Import Objects Users can be imported using the Objects User's ID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_objects_user.objects_user 192f95401a23ef307d42e4ba0fdc475e9630db45132a5b499d1dd2425c28a0ca ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The unique identifier of the Objects User.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `A list of key objects containing the access and secret key associated with the Objects User. Currently, only one key object is returned. Each key object has the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `The S3 access key of the Objects User.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `The S3 secret key of the Objects User. ## Import Objects Users can be imported using the Objects User's ID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_objects_user.objects_user 192f95401a23ef307d42e4ba0fdc475e9630db45132a5b499d1dd2425c28a0ca ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_server",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the new server. The name has to be a valid host name or a fully qualified domain name (FQDN).`,
				},
				resource.Attribute{
					Name:        "flavor_slug",
					Description: `(Required) The slug (name) of the flavor to use for the new server. Possible values can be found in our [API documentation](https://www.cloudscale.ch/en/api/v1#flavors).`,
				},
				resource.Attribute{
					Name:        "image_slug",
					Description: `(Required, if ` + "`" + `image_uuid` + "`" + ` not set) The slug (name) of the image (or custom image) to use for the new server. Possible values can be found in our [API documentation](https://www.cloudscale.ch/en/api/v1#images).`,
				},
				resource.Attribute{
					Name:        "image_uuid",
					Description: `(Required, if ` + "`" + `image_slug` + "`" + ` not set) The UUID of the custom image to use for the new server.`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `(Optional) A list of SSH public keys. Use the full content of your \`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password of the default user of the new server. When omitted, no password will be set.`,
				},
				resource.Attribute{
					Name:        "zone_slug",
					Description: `(Optional) The slug of the zone in which the new server will be created. Options include ` + "`" + `lpg1` + "`" + ` and ` + "`" + `rma1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "volume_size_gb",
					Description: `(Optional) The size in GB of the SSD root volume of the new server. If this parameter is not specified, the value will be set to 10. The minimum value is 10.`,
				},
				resource.Attribute{
					Name:        "bulk_volume_size_gb",
					Description: `(Optional, Deprecated) The size in GB of the bulk storage volume of the new server. If this parameter is not specified, no bulk storage volume will be attached to the server. Valid values are multiples of 100.`,
				},
				resource.Attribute{
					Name:        "use_public_network",
					Description: `(Optional) Attach the public network interface to the new server. Can be ` + "`" + `true` + "`" + ` (default) or ` + "`" + `false` + "`" + `. Use [` + "`" + `interfaces` + "`" + `](#interfaces) option for advanced setups.`,
				},
				resource.Attribute{
					Name:        "use_private_network",
					Description: `(Optional) Attach the ` + "`" + `default` + "`" + ` private network interface to the new server. Can be ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + ` (default). Use [` + "`" + `interfaces` + "`" + `](#interfaces) option for advanced setups.`,
				},
				resource.Attribute{
					Name:        "use_ipv6",
					Description: `(Optional) Enable/disable IPv6 on the public network interface of the new server. Can be ` + "`" + `true` + "`" + ` (default) or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) A list of interface configuration objects (see [example](network.html)). Each interface object has the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the interface. Can be ` + "`" + `public` + "`" + ` or ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `(Optional, can be set only for ` + "`" + `private` + "`" + ` interfaces) The UUID of the private network this interface should be attached to. Must be compatible with ` + "`" + `subnet_uuid` + "`" + ` if both are specified.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `(Optional, can be set only for ` + "`" + `private` + "`" + ` interfaces) A list of address objects:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) An IP address that has been assigned to this server.`,
				},
				resource.Attribute{
					Name:        "subnet_uuid",
					Description: `(Optional) The UUID of the subnet this address should be part of. Must be compatible with ` + "`" + `network_uuid` + "`" + ` if both are specified.`,
				},
				resource.Attribute{
					Name:        "no_address",
					Description: `(Optional, can be set only for ` + "`" + `private` + "`" + ` interfaces) You need to set this to ` + "`" + `true` + "`" + ` if no address should be configured, e.g. if you want to attach to a network without a subnet.`,
				},
				resource.Attribute{
					Name:        "server_group_ids",
					Description: `(Optional) A list of server group UUIDs to which the server should be added. Default to an empty list.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) User data (custom cloud-config settings) to use for the new server. Needs to be valid YAML. A default configuration will be used if this parameter is not specified or set to null. Use only if you are an advanced user with knowledge of cloud-config and cloud-init.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The desired state of a server. Can be ` + "`" + `running` + "`" + ` (default) or ` + "`" + `stopped` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_stopping_for_update",
					Description: `(Optional) If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires stopping the instance without setting this field, the update will fail.`,
				},
				resource.Attribute{
					Name:        "skip_waiting_for_ssh_host_keys",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, do not wait until SSH host keys become available.`,
				},
				resource.Attribute{
					Name:        "timeouts",
					Description: `(Optional) Specify how long certain operations are allowed to take before being considered to have failed. Currently, only the ` + "`" + `create` + "`" + ` timeout can be specified. Takes a string representation of a duration such as ` + "`" + `5m` + "`" + ` for 5 minutes (default), ` + "`" + `10s` + "`" + ` for ten seconds, or ` + "`" + `2h` + "`" + ` for two hours.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags allow you to assign custom metadata to resources: ` + "`" + `` + "`" + `` + "`" + ` tags = { foo = "bar" } ` + "`" + `` + "`" + `` + "`" + ` Tags are always strings (both keys and values). The following arguments are supported when updating servers:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `New name of the server. The name has to be a valid host name or a fully qualified domain name (FQDN).`,
				},
				resource.Attribute{
					Name:        "volume_size_gb",
					Description: `The size in GB of the SSD root volume of the new server.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `A list of interface configuration objects. Each interface object has the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the interface. Can be ` + "`" + `public` + "`" + ` or ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The desired state of a server. Can be ` + "`" + `running` + "`" + ` (default) or ` + "`" + `stopped` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Change tags (see documentation above) ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this server.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "ssh_fingerprints",
					Description: `A list of SSH host key fingerprints (strings) of this server.`,
				},
				resource.Attribute{
					Name:        "ssh_host_keys",
					Description: `A list of SSH host keys (strings) of this server.`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `A list of volume objects attached to this server. Each volume object has three attributes:`,
				},
				resource.Attribute{
					Name:        "device_path",
					Description: `The path (string) to the volume on your server (e.g. ` + "`" + `/dev/vda` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size (int) of the volume in GB. Typically matches ` + "`" + `volume_size_gb` + "`" + ` or ` + "`" + `bulk_volume_size_gb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `A string. Either ` + "`" + `ssd` + "`" + ` or ` + "`" + `bulk` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ipv4_address",
					Description: `The first ` + "`" + `public` + "`" + ` IPv4 address of this server. The returned IP address may be ` + "`" + `""` + "`" + ` if the server does not have a public IPv4.`,
				},
				resource.Attribute{
					Name:        "private_ipv4_address",
					Description: `The first ` + "`" + `private` + "`" + ` IPv4 address of this server. The returned IP address may be ` + "`" + `""` + "`" + ` if the server does not have private networking enabled.`,
				},
				resource.Attribute{
					Name:        "public_ipv6_address",
					Description: `The first ` + "`" + `public` + "`" + ` IPv6 address of this server. The returned IP address may be ` + "`" + `""` + "`" + ` if the server does not have a public IPv6.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `A list of interface objects attached to this server. Each interface object has the following attributes:`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `The name of the network the interface is attached to.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `The cloudscale.ch API URL of the network the interface is attached to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Either ` + "`" + `public` + "`" + ` or ` + "`" + `private` + "`" + `. Public interfaces are connected to the Internet, while private interfaces are not.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `A list of address objects:`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `An IPv4 or IPv6 address that represents the default gateway for this interface.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `The prefix length for this IP address, typically 24 for IPv4 and 128 for IPv6.`,
				},
				resource.Attribute{
					Name:        "reverse_ptr",
					Description: `The PTR record (reverse DNS pointer) for this IP address. If you use an FQDN as your server name it will automatically be used here.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The IP version, either ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `The cidr of the subnet the address is part of.`,
				},
				resource.Attribute{
					Name:        "subnet_href",
					Description: `The cloudscale.ch API URL of the subnet the address is part of. ## Import Servers can be imported using the server's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_server.server 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of this server.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "ssh_fingerprints",
					Description: `A list of SSH host key fingerprints (strings) of this server.`,
				},
				resource.Attribute{
					Name:        "ssh_host_keys",
					Description: `A list of SSH host keys (strings) of this server.`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `A list of volume objects attached to this server. Each volume object has three attributes:`,
				},
				resource.Attribute{
					Name:        "device_path",
					Description: `The path (string) to the volume on your server (e.g. ` + "`" + `/dev/vda` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size (int) of the volume in GB. Typically matches ` + "`" + `volume_size_gb` + "`" + ` or ` + "`" + `bulk_volume_size_gb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `A string. Either ` + "`" + `ssd` + "`" + ` or ` + "`" + `bulk` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ipv4_address",
					Description: `The first ` + "`" + `public` + "`" + ` IPv4 address of this server. The returned IP address may be ` + "`" + `""` + "`" + ` if the server does not have a public IPv4.`,
				},
				resource.Attribute{
					Name:        "private_ipv4_address",
					Description: `The first ` + "`" + `private` + "`" + ` IPv4 address of this server. The returned IP address may be ` + "`" + `""` + "`" + ` if the server does not have private networking enabled.`,
				},
				resource.Attribute{
					Name:        "public_ipv6_address",
					Description: `The first ` + "`" + `public` + "`" + ` IPv6 address of this server. The returned IP address may be ` + "`" + `""` + "`" + ` if the server does not have a public IPv6.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `A list of interface objects attached to this server. Each interface object has the following attributes:`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `The name of the network the interface is attached to.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `The cloudscale.ch API URL of the network the interface is attached to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Either ` + "`" + `public` + "`" + ` or ` + "`" + `private` + "`" + `. Public interfaces are connected to the Internet, while private interfaces are not.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `A list of address objects:`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `An IPv4 or IPv6 address that represents the default gateway for this interface.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `The prefix length for this IP address, typically 24 for IPv4 and 128 for IPv6.`,
				},
				resource.Attribute{
					Name:        "reverse_ptr",
					Description: `The PTR record (reverse DNS pointer) for this IP address. If you use an FQDN as your server name it will automatically be used here.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The IP version, either ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `The cidr of the subnet the address is part of.`,
				},
				resource.Attribute{
					Name:        "subnet_href",
					Description: `The cloudscale.ch API URL of the subnet the address is part of. ## Import Servers can be imported using the server's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_server.server 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_server_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the new server group.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the server group can currently only be ` + "`" + `"anti-affinity"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_slug",
					Description: `(Optional) The slug of the zone in which the new server group will be created. Options include ` + "`" + `lpg1` + "`" + ` and ` + "`" + `rma1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags allow you to assign custom metadata to resources: ` + "`" + `` + "`" + `` + "`" + ` tags = { foo = "bar" } ` + "`" + `` + "`" + `` + "`" + ` Tags are always strings (both keys and values). The following arguments are supported when updating server groups:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The new name of the server group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Change tags (see documentation above) ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource. ## Import Server groups can be imported using the server group's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_server_group.server_group 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource. ## Import Server groups can be imported using the server group's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_server_group.server_group 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_subnet",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The address range in CIDR notation. Must be at least /24.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `(Required) The network of the subnet.`,
				},
				resource.Attribute{
					Name:        "gateway_address",
					Description: `(Optional) The gateway address of the subnet.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `(Optional) A list of DNS resolver IP addresses, that act as DNS servers.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags allow you to assign custom metadata to resources: ` + "`" + `` + "`" + `` + "`" + ` tags = { foo = "bar" } ` + "`" + `` + "`" + `` + "`" + ` Tags are always strings (both keys and values). ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current subnet.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `The network name of the subnet.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `The cloudscale.ch API URL of the subnet's network. ## Import Subnets can be imported using the subnet's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_subnet.subnet 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current subnet.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `The network name of the subnet.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `The cloudscale.ch API URL of the subnet's network. ## Import Subnets can be imported using the subnet's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_subnet.subnet 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_volume",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the new volume.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `(Required) The volume size in GB. Valid values are multiples of 1 for type "ssd" and multiples of 100 for type "bulk".`,
				},
				resource.Attribute{
					Name:        "zone_slug",
					Description: `(Optional) The slug of the zone in which the new volume will be created. Options include ` + "`" + `lpg1` + "`" + ` and ` + "`" + `rma1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) For SSD/NVMe volumes specify "ssd" (default) or use "bulk" for our HDD cluster with NVMe caching. This is the only attribute that cannot be altered.`,
				},
				resource.Attribute{
					Name:        "server_uuids",
					Description: `(Optional) A list of server UUIDs. Default to an empty list. Currently a volume can only be attached to one server UUID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags allow you to assign custom metadata to resources: ` + "`" + `` + "`" + `` + "`" + ` tags = { foo = "bar" } ` + "`" + `` + "`" + `` + "`" + ` Tags are always strings (both keys and values). ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource. ## Import Volumes can be imported using the volume's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_volume.volume 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource. ## Import Volumes can be imported using the volume's UUID: ` + "`" + `` + "`" + `` + "`" + ` terraform import cloudscale_volume.volume 48151623-42aa-aaaa-bbbb-caffeeeeeeee ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"cloudscale_custom_image":                 0,
		"cloudscale_floating_ip":                  1,
		"cloudscale_load_balancer":                2,
		"cloudscale_load_balancer_health_monitor": 3,
		"cloudscale_load_balancer_listener":       4,
		"cloudscale_load_balancer_pool":           5,
		"cloudscale_load_balancer_pool_member":    6,
		"cloudscale_network":                      7,
		"cloudscale_objects_user":                 8,
		"cloudscale_server":                       9,
		"cloudscale_server_group":                 10,
		"cloudscale_subnet":                       11,
		"cloudscale_volume":                       12,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
