package cloudscale

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_custom_image",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The UUID of a custom image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The human-readable name of a custom image.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Optional) A string identifying a custom image. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
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
					Name:        "user_data_handling",
					Description: `How user_data will be handled when creating a server. Options include ` + "`" + `pass-through` + "`" + ` and ` + "`" + `extend-cloud-config` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "firmware_type",
					Description: `The firmware type that will be used for servers created with the custom image. Options include ` + "`" + `bios` + "`" + ` and ` + "`" + `uefi` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_slugs",
					Description: `The zones in which the custom image will be available. Options include ` + "`" + `lpg1` + "`" + ` and ` + "`" + `rma1` + "`" + `.`,
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
					Name:        "user_data_handling",
					Description: `How user_data will be handled when creating a server. Options include ` + "`" + `pass-through` + "`" + ` and ` + "`" + `extend-cloud-config` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "firmware_type",
					Description: `The firmware type that will be used for servers created with the custom image. Options include ` + "`" + `bios` + "`" + ` and ` + "`" + `uefi` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_slugs",
					Description: `The zones in which the custom image will be available. Options include ` + "`" + `lpg1` + "`" + ` and ` + "`" + `rma1` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_floating_ip",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The network IP of the floating IP, e.g. ` + "`" + `192.0.2.0` + "`" + ` of the network ` + "`" + `192.0.2.0/24` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) The CIDR notation of the Floating IP address or network, e.g. ` + "`" + `192.0.2.123/32` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reverse_ptr",
					Description: `(Optional) The PTR record (reverse DNS pointer) in case of a single Floating IP address.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `, for an IPv4 or IPv6 address or network respectively.`,
				},
				resource.Attribute{
					Name:        "region_slug",
					Description: `(Optional) The slug of the region in which a Regional Floating IP is assigned.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Options include ` + "`" + `regional` + "`" + ` and ` + "`" + `global` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `The IP address of the server that your Floating IP is currently assigned to.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The UUID of the server that your Floating IP is currently assigned to.`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `The UUID of the load balancer that your Floating IP is currently assigned to.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `The prefix length of a Floating IP (e.g. /128 or /56, as an integer).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `The IP address of the server that your Floating IP is currently assigned to.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The UUID of the server that your Floating IP is currently assigned to.`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `The UUID of the load balancer that your Floating IP is currently assigned to.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `The prefix length of a Floating IP (e.g. /128 or /56, as an integer).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_load_balancer",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The UUID of the load balancer.`,
				},
				resource.Attribute{
					Name:        "zone_slug",
					Description: `(Optional) The slug of the zone in which the load balancer exists. Options include ` + "`" + `lpg1` + "`" + ` and ` + "`" + `rma1` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "flavor_slug",
					Description: `The slug (name) of the flavor to use for the new load balancer. Possible values can be found in our [API documentation](https://www.cloudscale.ch/en/api/v1#load-balancer-flavors).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the new load balancer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the load balancer.`,
				},
				resource.Attribute{
					Name:        "vip_addresses",
					Description: `A list of VIP address objects. This attributes needs to be specified if the load balancer should be assigned a VIP address in a subnet on a private network. If the VIP address should be created on the public network, this attribute should be omitted. Each VIP address object has the following attributes:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The IP version, either ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_href",
					Description: `The cloudscale.ch API URL of the subnet the VIP address is part of.`,
				},
				resource.Attribute{
					Name:        "subnet_uuid",
					Description: `The UUID of the subnet this VIP address should be part of.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `The cidr of the subnet the VIP address is part of.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `An VIP address that has been assigned to this load balancer.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "flavor_slug",
					Description: `The slug (name) of the flavor to use for the new load balancer. Possible values can be found in our [API documentation](https://www.cloudscale.ch/en/api/v1#load-balancer-flavors).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the new load balancer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the load balancer.`,
				},
				resource.Attribute{
					Name:        "vip_addresses",
					Description: `A list of VIP address objects. This attributes needs to be specified if the load balancer should be assigned a VIP address in a subnet on a private network. If the VIP address should be created on the public network, this attribute should be omitted. Each VIP address object has the following attributes:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The IP version, either ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_href",
					Description: `The cloudscale.ch API URL of the subnet the VIP address is part of.`,
				},
				resource.Attribute{
					Name:        "subnet_uuid",
					Description: `The UUID of the subnet this VIP address should be part of.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `The cidr of the subnet the VIP address is part of.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `An VIP address that has been assigned to this load balancer.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_load_balancer_pool",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The UUID of the load balancer pool.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the load balancer pool.`,
				},
				resource.Attribute{
					Name:        "load_balancer_uuid",
					Description: `(Optional) The load balancer of the pool. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
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
					Description: `The cloudscale.ch API URL of the pool's load balancer.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The algorithm according to which the incoming traffic is distributed between the pool members. Options include ` + "`" + `"round_robin"` + "`" + `, ` + "`" + `"least_connections"` + "`" + ` and ` + "`" + `"source_ip"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol used for traffic between the load balancer and the pool members. Options include: ` + "`" + `"tcp"` + "`" + `, ` + "`" + `"proxy"` + "`" + ` and ` + "`" + `"proxyv2"` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The cloudscale.ch API URL of the pool's load balancer.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The algorithm according to which the incoming traffic is distributed between the pool members. Options include ` + "`" + `"round_robin"` + "`" + `, ` + "`" + `"least_connections"` + "`" + ` and ` + "`" + `"source_ip"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol used for traffic between the load balancer and the pool members. Options include: ` + "`" + `"tcp"` + "`" + `, ` + "`" + `"proxy"` + "`" + ` and ` + "`" + `"proxyv2"` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_load_balancer_pool_health_monitor",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The UUID of the load balancer health monitor.`,
				},
				resource.Attribute{
					Name:        "pool_uuid",
					Description: `(Optional) The UUID of the pool this health monitor belongs to. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "delay_s",
					Description: `The delay between two successive checks in seconds.`,
				},
				resource.Attribute{
					Name:        "timeout_s",
					Description: `The maximum time allowed for an individual check in seconds.`,
				},
				resource.Attribute{
					Name:        "up_threshold",
					Description: `The number of checks that need to be successful before the ` + "`" + `monitor_status` + "`" + ` of a pool member changes to ` + "`" + `"up"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "down_threshold",
					Description: `The number of checks that need to fail before the ` + "`" + `monitor_status` + "`" + ` of a pool member changes to ` + "`" + `"down"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the health monitor. Options include: ` + "`" + `"ping"` + "`" + `, ` + "`" + `"tcp"` + "`" + `, ` + "`" + `"http"` + "`" + `, ` + "`" + `"https"` + "`" + ` and ` + "`" + `"tls-hello"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_expected_codes",
					Description: `The HTTP status codes allowed for a check to be considered successful. Can either be a list of status codes, for example ` + "`" + `["200", "202"]` + "`" + `, or a list containing a single range, for example ` + "`" + `["200-204"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `The HTTP method used for the check. Options include ` + "`" + `"CONNECT"` + "`" + `, ` + "`" + `"DELETE"` + "`" + `, ` + "`" + `"GET"` + "`" + `, ` + "`" + `"HEAD"` + "`" + `, ` + "`" + `"OPTIONS"` + "`" + `, ` + "`" + `"PATCH"` + "`" + `, ` + "`" + `"POST"` + "`" + `, ` + "`" + `"PUT"` + "`" + ` and ` + "`" + `"TRACE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_url_path",
					Description: `The URL used for the check.`,
				},
				resource.Attribute{
					Name:        "http_version",
					Description: `The HTTP version used for the check. Options include ` + "`" + `"1.0"` + "`" + ` and ` + "`" + `"1.1"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_host",
					Description: `The server name in the HTTP Host: header used for the check.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "delay_s",
					Description: `The delay between two successive checks in seconds.`,
				},
				resource.Attribute{
					Name:        "timeout_s",
					Description: `The maximum time allowed for an individual check in seconds.`,
				},
				resource.Attribute{
					Name:        "up_threshold",
					Description: `The number of checks that need to be successful before the ` + "`" + `monitor_status` + "`" + ` of a pool member changes to ` + "`" + `"up"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "down_threshold",
					Description: `The number of checks that need to fail before the ` + "`" + `monitor_status` + "`" + ` of a pool member changes to ` + "`" + `"down"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the health monitor. Options include: ` + "`" + `"ping"` + "`" + `, ` + "`" + `"tcp"` + "`" + `, ` + "`" + `"http"` + "`" + `, ` + "`" + `"https"` + "`" + ` and ` + "`" + `"tls-hello"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_expected_codes",
					Description: `The HTTP status codes allowed for a check to be considered successful. Can either be a list of status codes, for example ` + "`" + `["200", "202"]` + "`" + `, or a list containing a single range, for example ` + "`" + `["200-204"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `The HTTP method used for the check. Options include ` + "`" + `"CONNECT"` + "`" + `, ` + "`" + `"DELETE"` + "`" + `, ` + "`" + `"GET"` + "`" + `, ` + "`" + `"HEAD"` + "`" + `, ` + "`" + `"OPTIONS"` + "`" + `, ` + "`" + `"PATCH"` + "`" + `, ` + "`" + `"POST"` + "`" + `, ` + "`" + `"PUT"` + "`" + ` and ` + "`" + `"TRACE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_url_path",
					Description: `The URL used for the check.`,
				},
				resource.Attribute{
					Name:        "http_version",
					Description: `The HTTP version used for the check. Options include ` + "`" + `"1.0"` + "`" + ` and ` + "`" + `"1.1"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_host",
					Description: `The server name in the HTTP Host: header used for the check.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_load_balancer_pool_listener",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The UUID of the load balancer listener.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the load balancer listener.`,
				},
				resource.Attribute{
					Name:        "pool_uuid",
					Description: `(Optional) The UUID of the pool this listener belongs to. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_load_balancer_pool_member",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pool_uuid",
					Description: `(Required) The UUID of the pool this member belongs to.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The UUID of the load balancer pool. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Pool member will not receive traffic if ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `The load balancer name of the pool.`,
				},
				resource.Attribute{
					Name:        "pool_href",
					Description: `The cloudscale.ch API URL of the pool's load balancer.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `The port to which actual traffic is sent.`,
				},
				resource.Attribute{
					Name:        "monitor_port",
					Description: `The port to which health monitor checks are sent. If not specified, ` + "`" + `protocol_port` + "`" + ` will be used.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP address to which traffic is sent.`,
				},
				resource.Attribute{
					Name:        "subnet_uuid",
					Description: `The subnet UUID of the address must be specified here.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `The CIDR of the member's address subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_href",
					Description: `The cloudscale.ch API URL of the member's address subnet.`,
				},
				resource.Attribute{
					Name:        "monitor_status",
					Description: `The status of the pool's health monitor check for this member. Can be ` + "`" + `"up"` + "`" + `, ` + "`" + `"down"` + "`" + `, ` + "`" + `"changing"` + "`" + `, ` + "`" + `"no_monitor"` + "`" + ` and ` + "`" + `"unknown"` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Pool member will not receive traffic if ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pool_name",
					Description: `The load balancer name of the pool.`,
				},
				resource.Attribute{
					Name:        "pool_href",
					Description: `The cloudscale.ch API URL of the pool's load balancer.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `The port to which actual traffic is sent.`,
				},
				resource.Attribute{
					Name:        "monitor_port",
					Description: `The port to which health monitor checks are sent. If not specified, ` + "`" + `protocol_port` + "`" + ` will be used.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP address to which traffic is sent.`,
				},
				resource.Attribute{
					Name:        "subnet_uuid",
					Description: `The subnet UUID of the address must be specified here.`,
				},
				resource.Attribute{
					Name:        "subnet_cidr",
					Description: `The CIDR of the member's address subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_href",
					Description: `The cloudscale.ch API URL of the member's address subnet.`,
				},
				resource.Attribute{
					Name:        "monitor_status",
					Description: `The status of the pool's health monitor check for this member. Can be ` + "`" + `"up"` + "`" + `, ` + "`" + `"down"` + "`" + `, ` + "`" + `"changing"` + "`" + `, ` + "`" + `"no_monitor"` + "`" + ` and ` + "`" + `"unknown"` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_network",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The UUID of a network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of a network.`,
				},
				resource.Attribute{
					Name:        "zone_slug",
					Description: `(Optional) The zone slug of a network. Options include ` + "`" + `lpg1` + "`" + ` and ` + "`" + `rma1` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current network.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `The MTU size for the network.`,
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
					Description: `The UUID of this subnet.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current network.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `The MTU size for the network.`,
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
					Description: `The UUID of this subnet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_objects_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The unique identifier of the Objects User.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name of the Objects User.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Optional) The unique identifier of the Objects User. (Exactly the same as ` + "`" + `id` + "`" + `) ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the resource.`,
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
					Description: `The S3 secret key of the Objects User.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the resource.`,
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
					Description: `The S3 secret key of the Objects User.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_server",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The UUID of a server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the server.`,
				},
				resource.Attribute{
					Name:        "zone_slug",
					Description: `(Optional) The slug of the zone in which the server exists. Options include ` + "`" + `lpg1` + "`" + ` and ` + "`" + `rma1` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
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
					Name:        "flavor_slug",
					Description: `The slug (name) of the flavor used by this server.`,
				},
				resource.Attribute{
					Name:        "image_slug",
					Description: `The slug (name) of the image (or custom image) used by the server.`,
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
					Description: `The cloudscale.ch API URL of the subnet the address is part of.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The state of a server.`,
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
					Name:        "flavor_slug",
					Description: `The slug (name) of the flavor used by this server.`,
				},
				resource.Attribute{
					Name:        "image_slug",
					Description: `The slug (name) of the image (or custom image) used by the server.`,
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
					Description: `The cloudscale.ch API URL of the subnet the address is part of.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The state of a server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_server_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The UUID of a server group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the server group.`,
				},
				resource.Attribute{
					Name:        "zone_slug",
					Description: `(Optional) The slug of the zone in which the server group exists. Options include ` + "`" + `lpg1` + "`" + ` and ` + "`" + `rma1` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the server group can currently only be ` + "`" + `"anti-affinity"` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the server group can currently only be ` + "`" + `"anti-affinity"` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_subnet",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The UUID of the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) The address range in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "network_uuid",
					Description: `(Optional) The network UUID of the subnet.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `(Optional) The network name of the subnet.`,
				},
				resource.Attribute{
					Name:        "gateway_address",
					Description: `(Optional) The gateway address of the subnet. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current subnet.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `A list of DNS resolver IP addresses, that act as DNS servers.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `The cloudscale.ch API URL of the subnet's network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current subnet.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `A list of DNS resolver IP addresses, that act as DNS servers.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `The cloudscale.ch API URL of the subnet's network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudscale_volume",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The UUID of a volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Name of the volume.`,
				},
				resource.Attribute{
					Name:        "zone_slug",
					Description: `(Optional) The slug of the zone in which the new volume will be created. Options include ` + "`" + `lpg1` + "`" + ` and ` + "`" + `rma1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) For SSD/NVMe volumes "ssd" (default); or "bulk" for our HDD cluster with NVMe caching. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The volume size in GB. Valid values are multiples of 1 for type "ssd" and multiples of 100 for type "bulk".`,
				},
				resource.Attribute{
					Name:        "server_uuids",
					Description: `(Optional) A list of server UUIDs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The cloudscale.ch API URL of the current resource.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The volume size in GB. Valid values are multiples of 1 for type "ssd" and multiples of 100 for type "bulk".`,
				},
				resource.Attribute{
					Name:        "server_uuids",
					Description: `(Optional) A list of server UUIDs.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"cloudscale_custom_image":                      0,
		"cloudscale_floating_ip":                       1,
		"cloudscale_load_balancer":                     2,
		"cloudscale_load_balancer_pool":                3,
		"cloudscale_load_balancer_pool_health_monitor": 4,
		"cloudscale_load_balancer_pool_listener":       5,
		"cloudscale_load_balancer_pool_member":         6,
		"cloudscale_network":                           7,
		"cloudscale_objects_user":                      8,
		"cloudscale_server":                            9,
		"cloudscale_server_group":                      10,
		"cloudscale_subnet":                            11,
		"cloudscale_volume":                            12,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
