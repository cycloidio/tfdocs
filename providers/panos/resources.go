package panos

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "panos_address_group",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"address",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The address group's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the address group into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_addresses",
					Description: `(Optional) The address objects to include in this statically defined address group.`,
				},
				resource.Attribute{
					Name:        "dynamic_match",
					Description: `(Optional) The IP tags to include in this DAG.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The address group's description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_address_object",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"address",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the address object into (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The address object's name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of address object. This can be ` + "`" + `ip-netmask` + "`" + ` (default), ` + "`" + `ip-range` + "`" + `, ` + "`" + `fqdn` + "`" + `, or ` + "`" + `ip-wildcard` + "`" + ` (PAN-OS 9.0+).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The address object's value. This can take various forms depending on what type of address object this is, but can be something like ` + "`" + `192.168.80.150` + "`" + ` or ` + "`" + `192.168.80.0/24` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The address object's description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(list) List of administrative tags.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_address_objects",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"address",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the address object into (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `(Required, repeatable) An ` + "`" + `object` + "`" + ` spec, as defined below. The ` + "`" + `object` + "`" + ` spec support the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The address object's name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of address object. This can be ` + "`" + `ip-netmask` + "`" + ` (default), ` + "`" + `ip-range` + "`" + `, ` + "`" + `fqdn` + "`" + `, or ` + "`" + `ip-wildcard` + "`" + ` (PAN-OS 9.0+).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The address object's value. This can take various forms depending on what type of address object this is, but can be something like ` + "`" + `192.168.80.150` + "`" + ` or ` + "`" + `192.168.80.0/24` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The address object's description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(list) List of administrative tags.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_administrative_tag",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"administrative",
				"tag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the administrative tag into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The administrative tag's name.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) The tag's color. This should be either an empty string (no color) or a string such as ` + "`" + `color1` + "`" + ` or ` + "`" + `color15` + "`" + `. Note that for maximum portability, you should limit color usage to ` + "`" + `color16` + "`" + `, which was available in PAN-OS 6.1. PAN-OS 8.1's colors go up to ` + "`" + `color42` + "`" + `. The value ` + "`" + `color18` + "`" + ` is reserved internally by PAN-OS and thus not available for use.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The administrative tag's description.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_aggregate_interface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"aggregate",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Required) The vsys that will use this interface. This should be something like ` + "`" + `vsys1` + "`" + ` or ` + "`" + `vsys3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The interface mode. Valid values are ` + "`" + `layer3` + "`" + ` (default), ` + "`" + `layer2` + "`" + `, ` + "`" + `virtual-wire` + "`" + `, ` + "`" + `ha` + "`" + `, or ` + "`" + `decrypt-mirror` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional) The IPv6 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "enable_untagged_subinterface",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable untagged subinterfaces.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses.`,
				},
				resource.Attribute{
					Name:        "ipv6_enabled",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable IPv6.`,
				},
				resource.Attribute{
					Name:        "ipv6_interface_id",
					Description: `(Optional) The IPv6 interface ID.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable DHCP.`,
				},
				resource.Attribute{
					Name:        "create_dhcp_default_route",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to create a DHCP default route.`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route_metric",
					Description: `(Optional) The metric for the DHCP default route.`,
				},
				resource.Attribute{
					Name:        "lacp_enable",
					Description: `(bool) Enable LACP.`,
				},
				resource.Attribute{
					Name:        "lacp_fast_failover",
					Description: `(bool) Enable LACP fast failover.`,
				},
				resource.Attribute{
					Name:        "lacp_mode",
					Description: `LACP mode. Valid values are ` + "`" + `active` + "`" + ` or ` + "`" + `passive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lacp_transmission_rate",
					Description: `LACP transmission rate. Valid values are ` + "`" + `fast` + "`" + ` or ` + "`" + `slow` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lacp_system_priority",
					Description: `(int) LACP system priority.`,
				},
				resource.Attribute{
					Name:        "lacp_max_ports",
					Description: `(int) LACP max ports.`,
				},
				resource.Attribute{
					Name:        "lacp_ha_passive_pre_negotiation",
					Description: `(bool) LACP HA passive pre-negotiation.`,
				},
				resource.Attribute{
					Name:        "lacp_ha_enable_same_system_mac",
					Description: `(bool) LACP HA enable same system MAC.`,
				},
				resource.Attribute{
					Name:        "lacp_ha_same_system_mac_address",
					Description: `LACP HA same system MAC address.`,
				},
				resource.Attribute{
					Name:        "lldp_enable",
					Description: `(bool) Enable LLDP.`,
				},
				resource.Attribute{
					Name:        "lldp_profile",
					Description: `LLDP profile name.`,
				},
				resource.Attribute{
					Name:        "lldp_ha_passive_pre_negotiation",
					Description: `(bool) LLDP HA passive pre-negotiation.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "decrypt_forward",
					Description: `(Optional, bool, PAN-OS 8.1+) Set to ` + "`" + `true` + "`" + ` to enable decrypt forward.`,
				},
				resource.Attribute{
					Name:        "dhcp_send_hostname_enable",
					Description: `(Optional, PAN-OS 9.0+) For DHCP layer3 interfaces: enable sending the firewall or a custom hostname to DHCP server`,
				},
				resource.Attribute{
					Name:        "dhcp_send_hostname_value",
					Description: `(Optional, PAN-OS 9.0+) For DHCP layer3 interfaces: the interface hostname. Leaving this unspecified with ` + "`" + `dhcp_send_hostname_enable` + "`" + ` set means to send the system hostname.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_anti_spyware_security_profile",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"anti",
				"spyware",
				"security",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `(PAN-OS 8.x only) Packet capture setting. Valid values are ` + "`" + `disable` + "`" + `, ` + "`" + `single-packet` + "`" + `, or ` + "`" + `extended-capture` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sinkhole_ipv4_address",
					Description: `IPv4 sinkhole address.`,
				},
				resource.Attribute{
					Name:        "sinkhole_ipv6_address",
					Description: `IPv6 sinkhole address.`,
				},
				resource.Attribute{
					Name:        "threat_exceptions",
					Description: `(list) List of threat exceptions.`,
				},
				resource.Attribute{
					Name:        "bonet_list",
					Description: `(repeatable) Botnet spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "dns_category",
					Description: `(repeatable, PAN-OS 10.0+) DNS category spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "white_list",
					Description: `(repeatable, PAN-OS 10.0+) White list spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(repeatable) Rule list spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "exception",
					Description: `(repeatable) Except list spec, as defined below. ` + "`" + `botnet_list` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action to take. Valid values are ` + "`" + `alert` + "`" + `, ` + "`" + `default` + "`" + `, ` + "`" + `allow` + "`" + `, ` + "`" + `block` + "`" + `, or ` + "`" + `sinkhole` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `(PAN-OS 9.0+) Packet capture setting. Valid values are ` + "`" + `disable` + "`" + `, ` + "`" + `single-packet` + "`" + `, or ` + "`" + `extended-capture` + "`" + `. ` + "`" + `dns_category` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action to take. Valid values are ` + "`" + `alert` + "`" + `, ` + "`" + `default` + "`" + `, ` + "`" + `allow` + "`" + `, ` + "`" + `block` + "`" + `, or ` + "`" + `sinkhole` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_level",
					Description: `Logging level. Valid values are ` + "`" + `default` + "`" + `, ` + "`" + `none` + "`" + `, ` + "`" + `low` + "`" + `, ` + "`" + `informational` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `high` + "`" + `, or ` + "`" + `critical` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `Packet capture setting. Valid values are ` + "`" + `disable` + "`" + `, ` + "`" + `single-packet` + "`" + `, or ` + "`" + `extended-capture` + "`" + `. ` + "`" + `white_list` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description ` + "`" + `rule` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name.`,
				},
				resource.Attribute{
					Name:        "threat_name",
					Description: `Threat name.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) Category.`,
				},
				resource.Attribute{
					Name:        "severities",
					Description: `(list) Severities.`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `Packet capture setting. Valid values are ` + "`" + `disable` + "`" + `, ` + "`" + `single-packet` + "`" + `, or ` + "`" + `extended-capture` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action. Valid values are ` + "`" + `default` + "`" + `, ` + "`" + `allow` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `drop` + "`" + `, ` + "`" + `reset-client` + "`" + `, ` + "`" + `reset-server` + "`" + `, ` + "`" + `reset-both` + "`" + `, or ` + "`" + `block-ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "block_ip_track_by",
					Description: `(action=` + "`" + `block-ip` + "`" + `) The track by setting.`,
				},
				resource.Attribute{
					Name:        "block_ip_duration",
					Description: `(action=` + "`" + `block-ip` + "`" + `, int) The duration. ` + "`" + `exception` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Threat name. You can use the ` + "`" + `panos_predefined_threat` + "`" + ` data source to discover the various phone home names available to use.`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `Packet capture setting. Valid values are ` + "`" + `disable` + "`" + `, ` + "`" + `single-packet` + "`" + `, or ` + "`" + `extended-capture` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action. Valid values are ` + "`" + `default` + "`" + `, ` + "`" + `allow` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `drop` + "`" + `, ` + "`" + `reset-client` + "`" + `, ` + "`" + `reset-server` + "`" + `, ` + "`" + `reset-both` + "`" + `, or ` + "`" + `block-ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "block_ip_track_by",
					Description: `(action=` + "`" + `block-ip` + "`" + `) The track by setting.`,
				},
				resource.Attribute{
					Name:        "block_ip_duration",
					Description: `(action=` + "`" + `block-ip` + "`" + `, int) The duration.`,
				},
				resource.Attribute{
					Name:        "exempt_ips",
					Description: `(list) List of exempt IPs.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_antivirus_security_profile",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"antivirus",
				"security",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to enable packet capture.`,
				},
				resource.Attribute{
					Name:        "threat_exceptions",
					Description: `(list) List of threat exceptions.`,
				},
				resource.Attribute{
					Name:        "decoder",
					Description: `(Repeatable) A decoder spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "application_exception",
					Description: `(Repeatable) An application exception spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "machine_learning_model",
					Description: `(Repeatable) A machine learning model spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "machine_learning_exception",
					Description: `(Repeatable) A machine learning exception spec, as defined below. ` + "`" + `decoder` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Decoder name.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Decoder action. Valid values are ` + "`" + `default` + "`" + ` (default), ` + "`" + `allow` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `block` + "`" + ` (PAN-OS 6.1 only), ` + "`" + `drop` + "`" + ` (PAN-OS 7.0+), ` + "`" + `reset-client` + "`" + ` (PAN-OS 7.0+), ` + "`" + `reset-server` + "`" + ` (PAN-OS 7.0+), or ` + "`" + `reset-both` + "`" + ` (PAN-OS 7.0+).`,
				},
				resource.Attribute{
					Name:        "wildfire_action",
					Description: `Wildfire action. Valid values are ` + "`" + `default` + "`" + ` (default), ` + "`" + `allow` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `block` + "`" + ` (PAN-OS 6.1 only), ` + "`" + `drop` + "`" + ` (PAN-OS 7.0+), ` + "`" + `reset-client` + "`" + ` (PAN-OS 7.0+), ` + "`" + `reset-server` + "`" + ` (PAN-OS 7.0+), or ` + "`" + `reset-both` + "`" + ` (PAN-OS 7.0+).`,
				},
				resource.Attribute{
					Name:        "machine_learning_action",
					Description: `(PAN-OS 10.0+) Machine learning action. ` + "`" + `application_exception` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Required) The application name`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action. Valid values are ` + "`" + `default` + "`" + `, ` + "`" + `allow` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `block` + "`" + ` (PAN-OS 6.1 only), ` + "`" + `drop` + "`" + ` (PAN-OS 7.0+), ` + "`" + `reset-client` + "`" + ` (PAN-OS 7.0+), ` + "`" + `reset-server` + "`" + ` (PAN-OS 7.0+), or ` + "`" + `reset-both` + "`" + ` (PAN-OS 7.0+). ` + "`" + `machine_learning_model` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `(Required) The model.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action. ` + "`" + `machine_learning_exception` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `Filename`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_application_group",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"application",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The group's vsys (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group's name.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(Optional) List of applications in this group.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_application_object",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"application",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_group",
					Description: `The device group (default: ` + "`" + `shared` + "`" + `). NGFW:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The object's vsys (default: ` + "`" + `vsys1` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name.`,
				},
				resource.Attribute{
					Name:        "defaults",
					Description: `The application's defaults spec (defined below). To have a "defaults" of ` + "`" + `None` + "`" + `, omit this section.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) The category.`,
				},
				resource.Attribute{
					Name:        "subcategory",
					Description: `(Required) The subcategory.`,
				},
				resource.Attribute{
					Name:        "technology",
					Description: `(Required) The technology.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The object's description.`,
				},
				resource.Attribute{
					Name:        "timeout_settings",
					Description: `The timeout spec (defined below).`,
				},
				resource.Attribute{
					Name:        "risk",
					Description: `(int) The risk (default: 1).`,
				},
				resource.Attribute{
					Name:        "parent_app",
					Description: `The parent application.`,
				},
				resource.Attribute{
					Name:        "able_to_file_transfer",
					Description: `(bool) Able to file transfer.`,
				},
				resource.Attribute{
					Name:        "excessive_bandwidth",
					Description: `(bool) Excessive bandwidth use.`,
				},
				resource.Attribute{
					Name:        "tunnels_other_applications",
					Description: `(bool) This application tunnels other apps.`,
				},
				resource.Attribute{
					Name:        "has_known_vulnerability",
					Description: `(bool) Has known vulnerabilities.`,
				},
				resource.Attribute{
					Name:        "used_by_malware",
					Description: `(bool) App is used by malware.`,
				},
				resource.Attribute{
					Name:        "evasive_behavior",
					Description: `(bool) App is evasive.`,
				},
				resource.Attribute{
					Name:        "pervasive_use",
					Description: `(bool) App is pervasive.`,
				},
				resource.Attribute{
					Name:        "prone_to_misuse",
					Description: `(bool) Prone to misuse.`,
				},
				resource.Attribute{
					Name:        "continue_scanning_for_other_applications",
					Description: `(bool) Continue scanning for other applications.`,
				},
				resource.Attribute{
					Name:        "scanning",
					Description: `The scanning spec (defined below).`,
				},
				resource.Attribute{
					Name:        "alg_disable_capability",
					Description: `The alg disable capability.`,
				},
				resource.Attribute{
					Name:        "no_app_id_caching",
					Description: `(bool) No appid caching. ` + "`" + `defaults` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port spec (defined below)`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The ip protocol spec (defined below)`,
				},
				resource.Attribute{
					Name:        "icmp",
					Description: `The ICMP spec (defined below)`,
				},
				resource.Attribute{
					Name:        "icmp6",
					Description: `The ICMP6 spec (defined below) ` + "`" + `defaults.port` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Required) List of ports. ` + "`" + `defaults.ip_protocol` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The IP protocol value. ` + "`" + `defaults.icmp` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, int) The type.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(int) The code. ` + "`" + `defaults.icmp6` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, int) The type.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(int) The code. ` + "`" + `timeout_settings` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(int) The timeout.`,
				},
				resource.Attribute{
					Name:        "tcp_timeout",
					Description: `(int) TCP timeout.`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `(int) UDP timeout.`,
				},
				resource.Attribute{
					Name:        "tcp_half_closed",
					Description: `(int) TCP half closed timeout.`,
				},
				resource.Attribute{
					Name:        "tcp_time_wait",
					Description: `(int) TCP time wait timeout. ` + "`" + `scanning` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `(bool) File type scanning.`,
				},
				resource.Attribute{
					Name:        "viruses",
					Description: `(bool) Virus scanning.`,
				},
				resource.Attribute{
					Name:        "data_patterns",
					Description: `(bool) Data pattern scanning.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_application_signature",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"application",
				"signature",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The signature's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The signature's vsys (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "application_object",
					Description: `(Required) The applciation object for this signature.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The signature's scope. Valid values are ` + "`" + `transaction` + "`" + ` (default) or ` + "`" + `session` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ordered_match",
					Description: `(Optional, bool) Set to ` + "`" + `false` + "`" + ` to disable ordered matching (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "and_condition",
					Description: `(Optional) The and condition spec (defined below). ` + "`" + `and_condition` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) And condition name, this is computed and cannot be configured.`,
				},
				resource.Attribute{
					Name:        "or_condition",
					Description: `(Required) The or condition spec (defined below). ` + "`" + `and_condition.or_condition` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) Or condition name, this is computed and cannot be configured.`,
				},
				resource.Attribute{
					Name:        "pattern_match",
					Description: `(Optional) The pattern match spec (defined below).`,
				},
				resource.Attribute{
					Name:        "greater_than",
					Description: `(Optional) The greater than spec (defined below).`,
				},
				resource.Attribute{
					Name:        "less_than",
					Description: `(Optional) the less than spec (defined below).`,
				},
				resource.Attribute{
					Name:        "equal_to",
					Description: `(Optional) The equal to spec (defined below). ` + "`" + `and_condition.or_condition.pattern_match` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "context",
					Description: `(Required) The context.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Required) The pattern.`,
				},
				resource.Attribute{
					Name:        "qualifiers",
					Description: `(Optional, map) The qualifiers. ` + "`" + `and_condition.or_condition.greater_than` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "context",
					Description: `(Required) The context.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value.`,
				},
				resource.Attribute{
					Name:        "qualifiers",
					Description: `(Optional, map) The qualifiers. ` + "`" + `and_condition.or_condition.less_than` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "context",
					Description: `(Required) The context.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value.`,
				},
				resource.Attribute{
					Name:        "qualifiers",
					Description: `(Optional, map) The qualifiers. ` + "`" + `and_condition.or_condition.equal_to` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "context",
					Description: `(Required) The context.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Optional) The position.`,
				},
				resource.Attribute{
					Name:        "mask",
					Description: `(Optional) The mask.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_arp",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"arp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Optional, but Required for Panorama) The template location. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "interface_type",
					Description: `The interface type. Valid values are ` + "`" + `ethernet` + "`" + ` (default), ` + "`" + `aggregate-ethernet` + "`" + `, or ` + "`" + `vlan` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interface_name",
					Description: `The interface name (leave this empty for VLAN interfaces).`,
				},
				resource.Attribute{
					Name:        "subinterface_name",
					Description: `The subinterface name.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP address.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Required) The MAC address.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(` + "`" + `interface_type` + "`" + `=` + "`" + `vlan` + "`" + `) The interface.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_authentication_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"authentication",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. NGFW / Panorama:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "lockout_time",
					Description: `Number of minutes to lock-out.`,
				},
				resource.Attribute{
					Name:        "allow_list",
					Description: `(List of string) List of allowed users or user groups.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "username_modifier",
					Description: `(PAN-OS 7.0+) Username modifier to handle user domain. Valid values are ` + "`" + `"%USERINPUT%"` + "`" + ` (default), ` + "`" + `"%USERINPUT%@%USERDOMAIN%"` + "`" + `, or ` + "`" + `"%USERDOMAIN%\%USERINPUT%"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_domain",
					Description: `(PAN-OS 7.0+) Domain name to be used for authentication.`,
				},
				resource.Attribute{
					Name:        "single_sign_on",
					Description: `(PAN-OS 7.0+) Kerberos SSO settings spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "multi_factor_authentication",
					Description: `(PAN-OS 8.0+) Specify MFA configuration spec, as defined below. ` + "`" + `type` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "none",
					Description: `(bool) No authentication.`,
				},
				resource.Attribute{
					Name:        "local_database",
					Description: `(bool) Local database authentication.`,
				},
				resource.Attribute{
					Name:        "radius",
					Description: `Radius authentication, as defined below.`,
				},
				resource.Attribute{
					Name:        "ldap",
					Description: `LDAP authenticatin, as defined below. ` + "`" + `type.radius` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "server_profile",
					Description: `(Required) Radius server profile object.`,
				},
				resource.Attribute{
					Name:        "retrieve_user_group",
					Description: `(bool, PAN-OS 7.0+) Retrieve user group from RADIUS. ` + "`" + `type.ldap` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "server_profile",
					Description: `(Required) LDAP server profile object.`,
				},
				resource.Attribute{
					Name:        "login_attribute",
					Description: `Login attribute in LDAP server to authenticate against.`,
				},
				resource.Attribute{
					Name:        "password_expiry_warning",
					Description: `Number of days prior to warning a user about password expiry (default: ` + "`" + `"7"` + "`" + `). ` + "`" + `type.kerberos` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "server_profile",
					Description: `(Required) Kerberos server profile object.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `(Required, PAN-OS 7.0+) Realm name to be used for authentication. ` + "`" + `type.tacacs_plus` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "server_profile",
					Description: `(Required) TACACS+ server profile object.`,
				},
				resource.Attribute{
					Name:        "retrieve_user_group",
					Description: `(bool, PAN-OS 8.0+) Retrieve user group from TACACS+. ` + "`" + `type.saml` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "server_profile",
					Description: `(Required) SAML IDP server profile object.`,
				},
				resource.Attribute{
					Name:        "enable_single_logout",
					Description: `(bool) Enable single logout.`,
				},
				resource.Attribute{
					Name:        "request_signing_certificate",
					Description: `(Signing certificate for SAML requests.`,
				},
				resource.Attribute{
					Name:        "certificate_profile",
					Description: `Certificate profile for IDP and SP.`,
				},
				resource.Attribute{
					Name:        "username_attribute",
					Description: `Attribute name for username to be extracted from SAML response (default: ` + "`" + `"username"` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "user_group_attribute",
					Description: `User group attribute.`,
				},
				resource.Attribute{
					Name:        "admin_role_attribute",
					Description: `Admin role attribute.`,
				},
				resource.Attribute{
					Name:        "access_domain_attribute",
					Description: `Access domain attribute. ` + "`" + `single_sign_on` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `Kerberos realm to be used for authentication.`,
				},
				resource.Attribute{
					Name:        "service_principal",
					Description: `Kerberos service principal.`,
				},
				resource.Attribute{
					Name:        "keytab",
					Description: `Kerberos keytab. ` + "`" + `multi_factor_authentication` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Enable additional authentication factors.`,
				},
				resource.Attribute{
					Name:        "factors",
					Description: `(List of strings) List of additional authentication factors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_aws_cloud_watch",
			Category:         "Plugins",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"plugins",
				"aws",
				"cloud",
				"watch",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled",
					Description: `Enable AWS CloudWatch setup (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace (default: ` + "`" + `VMseries` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `(int) Update time (in min) (default: ` + "`" + `5` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bfd_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bfd",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The BBFD profile's name.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) BFD operation mode. Valid values are ` + "`" + `active` + "`" + ` (default) or ` + "`" + `passive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minimum_tx_interval",
					Description: `(Optional, int) Desired minimum TX interval in ms. Default is ` + "`" + `1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minimum_rx_interval",
					Description: `(Optional, int) Required minimum RX interval in ms. Default is ` + "`" + `1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "detection_multiplier",
					Description: `(Optional, int) Multiplier sent to remote system. Default is ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hold_time",
					Description: `(Optional, int) Delay transmission and reception of control packets in ms.`,
				},
				resource.Attribute{
					Name:        "minimum_rx_ttl",
					Description: `(Optional, int) Minimum accepted ttl on received BFD packet.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bgp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP configuration to.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable BGP or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Optional) Router ID of this BGP instance.`,
				},
				resource.Attribute{
					Name:        "as_number",
					Description: `(Optional) Local AS number.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `(Optional, PAN-OS 7.1+) BFD configuration.`,
				},
				resource.Attribute{
					Name:        "reject_default_route",
					Description: `(Optional, bool) Do not learn default route from BGP (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "install_route",
					Description: `(Optional, bool) Populate BGP learned route to global route table.`,
				},
				resource.Attribute{
					Name:        "aggregate_med",
					Description: `(Optional, bool) Aggregate route only if they have same MED attributes (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "default_local_preference",
					Description: `(Optional) Default local preference (default: ` + "`" + `"100"` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_format",
					Description: `(Optional) AS format. Valid values are ` + "`" + `"2-byte"` + "`" + ` (default) or ` + "`" + `"4-byte"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "always_compare_med",
					Description: `(Optional, bool) Always compare MEDs.`,
				},
				resource.Attribute{
					Name:        "deterministic_med_comparison",
					Description: `(Optional, bool) Deterministic MED comparison (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ecmp_multi_as",
					Description: `(Optional, bool) Support multiple AS in ECMP.`,
				},
				resource.Attribute{
					Name:        "enforce_first_as",
					Description: `(Optional, bool) Enforce First AS for EBGP (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_graceful_restart",
					Description: `(Optional, bool) Enable graceful restart (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "stale_route_time",
					Description: `(Optional, int) Time to remove stale routes after peer restart, in seconds (default: ` + "`" + `120` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "local_restart_time",
					Description: `(Optional, int) Local restart time to advertise to peer, in seconds (default: ` + "`" + `120` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "max_peer_restart_time",
					Description: `(Optional, int) Maximum of peer restart time accepted, in seconds (default: ` + "`" + `120` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "reflector_cluster_id",
					Description: `(Optional) Route reflector cluster ID.`,
				},
				resource.Attribute{
					Name:        "confederation_member_as",
					Description: `(Optional) Confederation requires member-AS number.`,
				},
				resource.Attribute{
					Name:        "allow_redistribute_default_route",
					Description: `(Optional, bool) Allow redistribute default route to BGP.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_aggregate",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bgp",
				"aggregate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to put the rule into.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The rule name.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Aggregating address prefix.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this rule (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "as_set",
					Description: `(Optional, bool) Generate AS-set attribute.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `(Optional, bool) Summarize route.`,
				},
				resource.Attribute{
					Name:        "local_preference",
					Description: `(Optional) New local preference value.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) New MED value.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional, int) New weight value.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop address.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) New route origin. Valid values are ` + "`" + `incomplete` + "`" + ` (default), ` + "`" + `igp` + "`" + `, or ` + "`" + `egp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_limit",
					Description: `(Optional, int) Add AS path limit attribute if it does not exist.`,
				},
				resource.Attribute{
					Name:        "as_path_type",
					Description: `(Optional) AS path update options. Valid values are ` + "`" + `none` + "`" + ` (default) or ` + "`" + `prepend` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_value",
					Description: `(Optional) For ` + "`" + `as_path_type` + "`" + ` of ` + "`" + `prepend` + "`" + `, the value to prepend.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community update options. Valid values are ` + "`" + `none` + "`" + ` (default), ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_value",
					Description: `(Optional) If ` + "`" + `community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. For the ` + "`" + `append` + "`" + ` and ` + "`" + `overwrite` + "`" + ` types specifically, valid values are ` + "`" + `no-export` + "`" + `, ` + "`" + `no-advertise` + "`" + `, ` + "`" + `local-as` + "`" + `, or ` + "`" + `nopeer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_type",
					Description: `(Optional) Extended community update options. Valid values are ` + "`" + `none` + "`" + ` (default), ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_vaule",
					Description: `(Optional) If ` + "`" + `extended_community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_aggregate_advertise_filter",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bgp",
				"aggregate",
				"advertise",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_aggregate",
					Description: `(Required) The BGP address aggregation rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below).`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry. Each ` + "`" + `address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_aggregate_suppress_filter",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bgp",
				"aggregate",
				"suppress",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_aggregate",
					Description: `(Required) The BGP address aggregation rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below).`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry. Each ` + "`" + `address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_auth_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bgp",
				"auth",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP auth profile to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Shared secret for the TCP MD5 authentication.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_conditional_adv",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bgp",
				"conditional",
				"adv",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP conditional advertisement to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "used_by",
					Description: `(Optional) List of BGP peer groups that use this rule.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_conditional_adv_advertise_filter",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bgp",
				"conditional",
				"adv",
				"advertise",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_conditional_adv",
					Description: `(Required) The BGP conditional advertisement to add this filter to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "address_prefixes",
					Description: `(Optional) List of matching address prefixes.`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_conditional_adv_non_exist_filter",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bgp",
				"conditional",
				"adv",
				"non",
				"exist",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_conditional_adv",
					Description: `(Required) The BGP conditional advertisement to add this filter to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "address_prefixes",
					Description: `(Optional) List of matching address prefixes.`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_dampening_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bgp",
				"dampening",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP dampening profile to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "cutoff",
					Description: `(Optional, float) Cutoff threshold value (default: ` + "`" + `1.25` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "reuse",
					Description: `(Optional, float) Reuse threshold value (default: ` + "`" + `0.5` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "max_hold_time",
					Description: `(Optional, int) Maximum hold-down time, in seconds (default: ` + "`" + `900` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "decay_half_life_reachable",
					Description: `(Optional, int) Decay half-life while reachable, in seconds (default: ` + "`" + `300` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "decay_half_life_unreachable",
					Description: `(Optional, int) Decay half-life while unreachable, in seconds (default: ` + "`" + `900` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_export_rule_group",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bgp",
				"export",
				"rule",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to put the rule into.`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `(Optional) A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `(Optional) Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The export rule definition (see below). The export rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The security rule name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this export rule (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "used_by",
					Description: `(Optional) List of peer groups.`,
				},
				resource.Attribute{
					Name:        "match_as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "match_community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "match_extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "match_med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "match_route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "match_address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below). below for the params for this section.`,
				},
				resource.Attribute{
					Name:        "match_next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "match_from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Rule action. Valid values are ` + "`" + `allow` + "`" + ` (default) or ` + "`" + `deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dampening",
					Description: `(Optional) Route flap dampening profile.`,
				},
				resource.Attribute{
					Name:        "local_preference",
					Description: `(Optional) New local preference value.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) New MED value.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional, int) New weight value.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop address.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) New route origin. Valid values are ` + "`" + `igp` + "`" + `, ` + "`" + `egp` + "`" + `, or ` + "`" + `incomplete` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_limit",
					Description: `(Optional, int) Add AS path limit attribute if it does not exist.`,
				},
				resource.Attribute{
					Name:        "as_path_type",
					Description: `(Optional) AS path update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove` + "`" + `, ` + "`" + `prepend` + "`" + `, or ` + "`" + `remove-and-prepend` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_value",
					Description: `(Optional) If ` + "`" + `as_path_type` + "`" + ` is ` + "`" + `prepend` + "`" + ` or ` + "`" + `remove-and-prepend` + "`" + `, the value to prepend.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_value",
					Description: `(Optional) If ` + "`" + `community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. For the ` + "`" + `append` + "`" + ` and ` + "`" + `overwrite` + "`" + ` types specifically, valid values for ` + "`" + `community_value` + "`" + ` are ` + "`" + `no-export` + "`" + `, ` + "`" + `no-advertise` + "`" + `, ` + "`" + `local-as` + "`" + `, or ` + "`" + `nopeer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_type",
					Description: `(Optional) Extended community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_vaule",
					Description: `(Optional) If ` + "`" + `extended_community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. Each ` + "`" + `match_address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_import_rule_group",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bgp",
				"import",
				"rule",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to put the rule into.`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `(Optional) A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `(Optional) Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The import rule definition (see below). The import rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The security rule name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this import rule (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "used_by",
					Description: `(Optional) List of peer groups.`,
				},
				resource.Attribute{
					Name:        "match_as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "match_community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "match_extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "match_med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "match_route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "match_address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below). below for the params for this section.`,
				},
				resource.Attribute{
					Name:        "match_next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "match_from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Rule action. Valid values are ` + "`" + `allow` + "`" + ` (default) or ` + "`" + `deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dampening",
					Description: `(Optional) Route flap dampening profile.`,
				},
				resource.Attribute{
					Name:        "local_preference",
					Description: `(Optional) New local preference value.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) New MED value.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional, int) New weight value.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop address.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) New route origin. Valid values are ` + "`" + `igp` + "`" + `, ` + "`" + `egp` + "`" + `, or ` + "`" + `incomplete` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_limit",
					Description: `(Optional, int) Add AS path limit attribute if it does not exist.`,
				},
				resource.Attribute{
					Name:        "as_path_type",
					Description: `(Optional) AS path update options. Valid values are ` + "`" + `none` + "`" + ` or ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_value",
					Description: `(Optional) If ` + "`" + `community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. For the ` + "`" + `append` + "`" + ` and ` + "`" + `overwrite` + "`" + ` types specifically, valid values for ` + "`" + `community_value` + "`" + ` are ` + "`" + `no-export` + "`" + `, ` + "`" + `no-advertise` + "`" + `, ` + "`" + `local-as` + "`" + `, or ` + "`" + `nopeer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_type",
					Description: `(Optional) Extended community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_vaule",
					Description: `(Optional) If ` + "`" + `extended_community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. Each ` + "`" + `match_address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_peer",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bgp",
				"peer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP peer to.`,
				},
				resource.Attribute{
					Name:        "bgp_peer_group",
					Description: `(Required) The BGP peer group to put this peer into.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "peer_as",
					Description: `(Optional) Peer AS number.`,
				},
				resource.Attribute{
					Name:        "local_address_interface",
					Description: `(Required) Interface to accept BGP session.`,
				},
				resource.Attribute{
					Name:        "local_address_ip",
					Description: `(Optional) Specify exact IP address if interface has multiple addresses.`,
				},
				resource.Attribute{
					Name:        "peer_address_ip",
					Description: `(Required) Peer IP address configuration.`,
				},
				resource.Attribute{
					Name:        "reflector_client",
					Description: `(Optional) This peer is reflector client. Valid values are ` + "`" + `non-client` + "`" + `, ` + "`" + `client` + "`" + `, or ` + "`" + `meshed-client` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "peering_type",
					Description: `(Optional) Peering type that affects NOPEER community value handling. Valid values are ` + "`" + `unspecified` + "`" + ` (default) or ` + "`" + `bilateral` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_prefixes",
					Description: `(Optional) Maximum of prefixes to receive from the peer. This can be a number such as ` + "`" + `"5000"` + "`" + ` (default) or ` + "`" + `unlimited` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_profile",
					Description: `(Optional) Auth profile.`,
				},
				resource.Attribute{
					Name:        "keep_alive_interval",
					Description: `(Optional, int) Keep alive interval, in seconds (default: ` + "`" + `30` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "multi_hop",
					Description: `(Optional, int) IP TTL value used for sending BGP packet.`,
				},
				resource.Attribute{
					Name:        "open_delay_time",
					Description: `(Optional, int) Open delay time, in seconds.`,
				},
				resource.Attribute{
					Name:        "hold_time",
					Description: `(Optional, int) Hold time, in seconds.`,
				},
				resource.Attribute{
					Name:        "idle_hold_time",
					Description: `(Optional, int) Idle hold time, in seconds.`,
				},
				resource.Attribute{
					Name:        "allow_incoming_connections",
					Description: `(Optional, bool) Allow incoming connections (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "incoming_connections_remote_port",
					Description: `(Optional, int) Restrict remote port for incoming BGP connections.`,
				},
				resource.Attribute{
					Name:        "allow_outgoing_connections",
					Description: `(Optional, bool) Allow outgoing connections (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "outgoing_connections_local_port",
					Description: `(Optional, int) Use specific local port for outgoing BGP connections.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `(Optional, PAN-OS 7.1+) BFD profile. This can be a specific BFD profile name, ` + "`" + `None` + "`" + ` (disables BFD), or ` + "`" + `Inherit-vr-global-setting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_mp_bgp",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable MP BGP.`,
				},
				resource.Attribute{
					Name:        "address_family_type",
					Description: `(Optional, PAN-OS 8.0+) Set the AFI for this peer. Valid values are ` + "`" + `ipv4` + "`" + ` or ` + "`" + `ipv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subsequent_address_family_unicast",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable unicast subsequent address family for this peer.`,
				},
				resource.Attribute{
					Name:        "subsequent_address_family_multicast",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable multicast subsequent address family for this peer.`,
				},
				resource.Attribute{
					Name:        "enable_sender_side_loop_detection",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable sender side loop detection.`,
				},
				resource.Attribute{
					Name:        "min_route_advertisement_interval",
					Description: `(Optional, int, PAN-OS 8.1+) Minimum route advertisement interval, in seconds.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_peer_group",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bgp",
				"peer",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP peer group to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "aggregated_confed_as_path",
					Description: `(Optional, bool) The peers understand aggregated confederation AS path (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "soft_reset_with_stored_info",
					Description: `(Optional, bool) Soft reset with stored info.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Peer group type. Valid options are ` + "`" + `ebgp` + "`" + ` (default), ` + "`" + `ebgp-confed` + "`" + `, ` + "`" + `ibgp` + "`" + `, or ` + "`" + `ibgp-confed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "export_next_hop",
					Description: `(Optional) Export next hop. Valid values are ` + "`" + `original` + "`" + `, ` + "`" + `use-self` + "`" + `, or ` + "`" + `resolve` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "import_next_hop",
					Description: `(Optional) Import next hop. Valid values are ` + "`" + `original` + "`" + `, ` + "`" + `use-peer` + "`" + `, or the empty string.`,
				},
				resource.Attribute{
					Name:        "remove_private_as",
					Description: `(Optional, bool) Remove private AS when exporting route. Only available for ` + "`" + `type=ebgp` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_bgp_redist_rule",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"bgp",
				"redist",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP redist rule to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A subnet or a redistribution profile.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this rule or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `(Optional) The address family. Valid values are ` + "`" + `ipv4` + "`" + ` (default) or ` + "`" + `ipv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional, int) Metric value.`,
				},
				resource.Attribute{
					Name:        "set_origin",
					Description: `(Optional) Add the origin path attribute. Valid values are ` + "`" + `incomplete` + "`" + ` (default), ` + "`" + `igp` + "`" + `, or ` + "`" + `egp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "set_med",
					Description: `(Optional) Add the MULTI_EXIT_DISC path attribute.`,
				},
				resource.Attribute{
					Name:        "set_local_preference",
					Description: `(Optional) Add the LOCAL_PREF path attribute.`,
				},
				resource.Attribute{
					Name:        "set_as_path_limit",
					Description: `(Optional, int) Add the AS_PATHLIMIT path attribute.`,
				},
				resource.Attribute{
					Name:        "set_communities",
					Description: `(Optional) List of COMMUNITY path attributes to add.`,
				},
				resource.Attribute{
					Name:        "set_extended_communities",
					Description: `(Optional) List of EXTENDED COMMUNITY path attributes to add.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_certificate_import",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"certificate",
				"import",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template. NGFW / Panorama:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "pem",
					Description: `A PEM style certificate, as defined below. Conflicts with ` + "`" + `pkcs12` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pkcs12",
					Description: `A PKCS12 style certificate, as defined below. Conflicts with ` + "`" + `pem` + "`" + `. ` + "`" + `pem` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) The contents of the certificate file.`,
				},
				resource.Attribute{
					Name:        "certificate_filename",
					Description: `The certificate filename for uploading to PAN-OS (default: ` + "`" + `cert.pem` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The contents of the private key file.`,
				},
				resource.Attribute{
					Name:        "private_key_filename",
					Description: `The private key filename for uplaoding to PAN-OS (default: ` + "`" + `key.pem` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "passphrase",
					Description: `(Required) The private key file passphrase. ` + "`" + `pkcs12` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) The contents of the certificate file.`,
				},
				resource.Attribute{
					Name:        "certificate_filename",
					Description: `The certificate filename for uploading to PAN-OS (default: ` + "`" + `cert.pfx` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "passphrase",
					Description: `(Required) The private key file passphrase. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "cert_format",
					Description: `The certificate format.`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `The common name.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The algorithm.`,
				},
				resource.Attribute{
					Name:        "ca",
					Description: `The CA.`,
				},
				resource.Attribute{
					Name:        "not_valid_after",
					Description: `Certificate is not valid after this date.`,
				},
				resource.Attribute{
					Name:        "not_valid_before",
					Description: `Certificate is not valid before this date.`,
				},
				resource.Attribute{
					Name:        "expiry_epoch",
					Description: `Expiry ephoch.`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `Subject.`,
				},
				resource.Attribute{
					Name:        "subject_hash",
					Description: `The subject hash.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `Certificate issuer.`,
				},
				resource.Attribute{
					Name:        "issuer_hash",
					Description: `Hash of the issuer.`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `The CSR.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `Public key.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `Encrypted private key.`,
				},
				resource.Attribute{
					Name:        "private_key_on_hsm",
					Description: `(bool) Private key on HSM.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status.`,
				},
				resource.Attribute{
					Name:        "revoke_date_epoch",
					Description: `Revoke date epoch.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cert_format",
					Description: `The certificate format.`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `The common name.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The algorithm.`,
				},
				resource.Attribute{
					Name:        "ca",
					Description: `The CA.`,
				},
				resource.Attribute{
					Name:        "not_valid_after",
					Description: `Certificate is not valid after this date.`,
				},
				resource.Attribute{
					Name:        "not_valid_before",
					Description: `Certificate is not valid before this date.`,
				},
				resource.Attribute{
					Name:        "expiry_epoch",
					Description: `Expiry ephoch.`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `Subject.`,
				},
				resource.Attribute{
					Name:        "subject_hash",
					Description: `The subject hash.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `Certificate issuer.`,
				},
				resource.Attribute{
					Name:        "issuer_hash",
					Description: `Hash of the issuer.`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `The CSR.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `Public key.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `Encrypted private key.`,
				},
				resource.Attribute{
					Name:        "private_key_on_hsm",
					Description: `(bool) Private key on HSM.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status.`,
				},
				resource.Attribute{
					Name:        "revoke_date_epoch",
					Description: `Revoke date epoch.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_certificate_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"certificate",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. NGFW / Panorama:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "username_field",
					Description: `Username field. Valid values are an empty string for ` + "`" + `None` + "`" + ` (default), ` + "`" + `subject` + "`" + `, or ` + "`" + `subject-alt` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username_field_value",
					Description: `The value. Common settings are ` + "`" + `common-name` + "`" + ` for ` + "`" + `username_field="subject"` + "`" + `, or ` + "`" + `email` + "`" + ` or ` + "`" + `principal-name` + "`" + ` for ` + "`" + `username_field="subject-alt"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `User domain.`,
				},
				resource.Attribute{
					Name:        "use_crl",
					Description: `(bool) Use CRL.`,
				},
				resource.Attribute{
					Name:        "use_ocsp",
					Description: `(bool) Use OCSP.`,
				},
				resource.Attribute{
					Name:        "crl_receive_timeout",
					Description: `(int) CRL receive timeout in seconds (default: ` + "`" + `5` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ocsp_receive_timeout",
					Description: `(int) OCSP receive timeout in seconds (default: ` + "`" + `5` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "certificate_status_timeout",
					Description: `(int) Certificate status timeout in seconds (default: ` + "`" + `5` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "block_unknown_certificate",
					Description: `(bool) Block session if certificate status is unknown.`,
				},
				resource.Attribute{
					Name:        "block_certificate_timeout",
					Description: `(bool) Block session if certificate status cannot be retrieved within timeout.`,
				},
				resource.Attribute{
					Name:        "block_unauthenticated_certificate",
					Description: `(bool) Block session if the certificate was not issued to the authenticating device.`,
				},
				resource.Attribute{
					Name:        "block_expired_certificate",
					Description: `(bool) Block sessions with expired certificates.`,
				},
				resource.Attribute{
					Name:        "ocsp_exclude_nonce",
					Description: `(bool) OCSP exclude nonce.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(repeated) List of CA certificates, defined below. ` + "`" + `certificate` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "default_ocsp_url",
					Description: `Default OCSP URL.`,
				},
				resource.Attribute{
					Name:        "ocsp_verify_certificate",
					Description: `OCSP verify certificate.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `Template name/OID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_custom_data_pattern_object",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"custom",
				"data",
				"pattern",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type. Valid values are ` + "`" + `file-properties` + "`" + ` (default), ` + "`" + `predefined` + "`" + `, or ` + "`" + `regex` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "predefined_pattern",
					Description: `(` + "`" + `type` + "`" + `=` + "`" + `predefined` + "`" + `, repeatable) List of predefined pattern specs, as definded below.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(` + "`" + `type` + "`" + `=` + "`" + `regex` + "`" + `) List of regex specs, as defined below.`,
				},
				resource.Attribute{
					Name:        "file_property",
					Description: `(` + "`" + `type` + "`" + `=` + "`" + `file-properties` + "`" + `) List of file properties specs, as defined below. ` + "`" + `predefined_pattern` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `(list) List of file types. ` + "`" + `regex` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name.`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `(list) List of file types.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Required) The regex. ` + "`" + `file_property` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name.`,
				},
				resource.Attribute{
					Name:        "file_type",
					Description: `(Required) The file type.`,
				},
				resource.Attribute{
					Name:        "file_property",
					Description: `(Required) File property.`,
				},
				resource.Attribute{
					Name:        "property_value",
					Description: `(Required) Property value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_custom_url_category",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"custom",
				"url",
				"category",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `The device group (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "sites",
					Description: `(list) The site list. Leave this undefined if you intend to manage the site listing with [` + "`" + `panos_custom_url_category_entry` + "`" + `](custom_url_category_entry.html) resources.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(PAN-OS 9.0+) The custom URL category type.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_custom_url_category_entry",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"custom",
				"url",
				"category",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `The device group (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "custom_url_category",
					Description: `(Required) The custom URL category name.`,
				},
				resource.Attribute{
					Name:        "site",
					Description: `(Required) The site to add to the specified custom URL category.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_dag_tags",
			Category:         "User-ID",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"user",
				"id",
				"dag",
				"tags",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the DAG tags in (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "register",
					Description: `(Required) A set that includes ` + "`" + `ip` + "`" + `, the IP address to be tagged and ` + "`" + `tags` + "`" + `, a list of tags to associate with the given IP.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_data_filtering_security_profile",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"data",
				"filtering",
				"security",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "data_capture",
					Description: `(bool) Data capture.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(repeatable) Rule list spec, as defined below. ` + "`" + `rule` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) Name.`,
				},
				resource.Attribute{
					Name:        "data_pattern",
					Description: `(Required) The data pattern name.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(list) List of applications.`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `(list) List of file types.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `Direction. Valid values are ` + "`" + `both` + "`" + ` (default), ` + "`" + `download` + "`" + `, or ` + "`" + `upload` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "alert_threshold",
					Description: `(int) Alert threshold.`,
				},
				resource.Attribute{
					Name:        "block_threshold",
					Description: `(int) Block threshold.`,
				},
				resource.Attribute{
					Name:        "log_severity",
					Description: `(PAN-OS 8.0+) Log severity.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_decryption_rule_group",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"decryption",
				"rule",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_group",
					Description: `The device group (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rulebase",
					Description: `The rulebase. This can be ` + "`" + `pre-rulebase` + "`" + ` (default), ` + "`" + `post-rulebase` + "`" + `, or ` + "`" + `rulebase` + "`" + `. NGFW specific arguments:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The rule definition (see below). The rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The rule name.`,
				},
				resource.Attribute{
					Name:        "audit_comment",
					Description: `When this rule is created/updated, the audit comment to apply for this rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `(Required) List of source zones.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Required) List of source addresses.`,
				},
				resource.Attribute{
					Name:        "negate_source",
					Description: `(bool) If the source should be negated.`,
				},
				resource.Attribute{
					Name:        "source_users",
					Description: `(Required) List of source users.`,
				},
				resource.Attribute{
					Name:        "destination_zones",
					Description: `(Required) List of destination zones.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Required) List of destination addresses.`,
				},
				resource.Attribute{
					Name:        "negate_destination",
					Description: `(bool) Negate the destination addresses.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List of administrative tags.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(bool) Disable this rule.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) List of services.`,
				},
				resource.Attribute{
					Name:        "url_categories",
					Description: `(Required) List of URL categories.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action to take. Valid values are ` + "`" + `no-decrypt` + "`" + ` (default), ` + "`" + `decrypt` + "`" + `, or ` + "`" + `decrypt-and-forward` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "decryption_type",
					Description: `The decryption type. Valid values are ` + "`" + `ssl-forward-proxy` + "`" + `, ` + "`" + `ssh-proxy` + "`" + `, or ` + "`" + `ssl-inbound-inspection` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl_certificate",
					Description: `(PAN-OS 10.1 and below) The SSL certificate.`,
				},
				resource.Attribute{
					Name:        "ssl_certificates",
					Description: `(PAN-OS 10.2+) List of SSL certificates.`,
				},
				resource.Attribute{
					Name:        "decryption_profile",
					Description: `The decryption profile.`,
				},
				resource.Attribute{
					Name:        "forwarding_profile",
					Description: `Forwarding profile.`,
				},
				resource.Attribute{
					Name:        "group_tag",
					Description: `The group tag.`,
				},
				resource.Attribute{
					Name:        "source_hips",
					Description: `List of source HIP devices.`,
				},
				resource.Attribute{
					Name:        "destination_hips",
					Description: `List of destination HIP devices.`,
				},
				resource.Attribute{
					Name:        "log_successful_tls_handshakes",
					Description: `Log successful TLS handshakes.`,
				},
				resource.Attribute{
					Name:        "log_failed_tls_handshakes",
					Description: `Log failed TLS handshakes.`,
				},
				resource.Attribute{
					Name:        "log_setting",
					Description: `The log setting.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(repeatable, Panorama only) A target definition (see below). If there are no target sections, then the rule will apply to every vsys of every device in the device group.`,
				},
				resource.Attribute{
					Name:        "negate_target",
					Description: `(bool, Panorama only) Instead of applying the rule for the given serial numbers, apply it to everything except them. ` + "`" + `rule.target` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `A listing of vsys to apply this rule to. If ` + "`" + `serial` + "`" + ` is a VM, then this parameter should just be omitted. ## Attributes Each ` + "`" + `rule` + "`" + ` has the following attributes:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The PAN-OS UUID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_device_group",
			Category:         "Panorama",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"device",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The device group's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The device group's description.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `The device definition (see below). The following arguments are valid for each ` + "`" + `device` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `(Optional) A subset of all available vsys on the firewall that should be in this device group. If the firewall is a virtual firewall, then this parameter should just be omitted.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_device_group_entry",
			Category:         "Panorama",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"device",
				"group",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_group",
					Description: `(Required) The device group's name.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `(Optional) A subset of all available vsys on the firewall that should be in this device group. If the firewall is a virtual firewall, then this parameter should just be omitted.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_device_group_parent",
			Category:         "Panorama",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"device",
				"group",
				"parent",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_group",
					Description: `(Required) The device group whose parent you intent to set.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `The parent device group name. Leaving this empty / unspecified means to move this device group under the "shared" device group.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_dos_protection_profile",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"dos",
				"protection",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The profile type. Valid values are ` + "`" + `aggregate` + "`" + ` (default) or ` + "`" + `classified` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_sessions_protections",
					Description: `(bool) Enable sessions protections.`,
				},
				resource.Attribute{
					Name:        "max_concurrent_sessions",
					Description: `(int) Max concurrent sessions.`,
				},
				resource.Attribute{
					Name:        "syn",
					Description: `Optional syn spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "udp",
					Description: `Optional protection spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "icmp",
					Description: `Optional ICMP spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "icmpv6",
					Description: `Optional ICMPv6 spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "other",
					Description: `Optional other IP flood protection spec, as defined below. ` + "`" + `syn` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(bool) Enable`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `SYN protection action. Valid values are ` + "`" + `red` + "`" + ` (default) for "Random Early Drop" or ` + "`" + `syn-cookies` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "alarm_rate",
					Description: `(int) Alarm rate.`,
				},
				resource.Attribute{
					Name:        "activate_rate",
					Description: `(int) Activate rate.`,
				},
				resource.Attribute{
					Name:        "max_rate",
					Description: `(int) Max rate.`,
				},
				resource.Attribute{
					Name:        "block_duration",
					Description: `(int) Block duration. ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `icmpv6` + "`" + `, and ` + "`" + `other` + "`" + ` all support the following arguments:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(bool) Enable`,
				},
				resource.Attribute{
					Name:        "alarm_rate",
					Description: `(int) Alarm rate.`,
				},
				resource.Attribute{
					Name:        "activate_rate",
					Description: `(int) Activate rate.`,
				},
				resource.Attribute{
					Name:        "max_rate",
					Description: `(int) Max rate.`,
				},
				resource.Attribute{
					Name:        "block_duration",
					Description: `(int) Block duration.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_dynamic_user_group",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"dynamic",
				"user",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the address object into (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) The filter.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(list) List of administrative tags.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_edl",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"edl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_group",
					Description: `The device group (default: ` + "`" + `shared` + "`" + `). NGFW specific arguments:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of EDL. This can be ` + "`" + `ip` + "`" + ` (the default; and the only valid value for PAN-OS 6.1 - 7.0), ` + "`" + `domain` + "`" + `, ` + "`" + `url` + "`" + `, ` + "`" + `predefined-ip` + "`" + ` (PAN-OS 8.0+), or ` + "`" + `predefined-url` + "`" + ` (PAN-OS 9.0+).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The object's description.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The EDL source URL`,
				},
				resource.Attribute{
					Name:        "certificate_profile",
					Description: `Profile for authenticating client certificates`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `EDL username`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `EDL password`,
				},
				resource.Attribute{
					Name:        "repeat",
					Description: `How often to retrieve the EDL. This can be ` + "`" + `hourly` + "`" + ` (the default), ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `, ` + "`" + `monthly` + "`" + `, or ` + "`" + `every five minutes` + "`" + ` (valid for PAN-OS 7.1+)`,
				},
				resource.Attribute{
					Name:        "repeat_at",
					Description: `The time at which to retrieve the EDL. Please refer to the section above for how to set this value properly.`,
				},
				resource.Attribute{
					Name:        "repeat_day_of_week",
					Description: `If ` + "`" + `repeat` + "`" + ` is ` + "`" + `weekly` + "`" + `, then this should be set to the desired day of the week. Valid values are ` + "`" + `sunday` + "`" + `, ` + "`" + `monday` + "`" + `, ` + "`" + `tuesday` + "`" + `, ` + "`" + `wednesday` + "`" + `, ` + "`" + `thursday` + "`" + `, ` + "`" + `friday` + "`" + `, ` + "`" + `saturday` + "`" + `, and ` + "`" + `sunday` + "`" + ``,
				},
				resource.Attribute{
					Name:        "repeat_day_of_month",
					Description: `(int) If ` + "`" + `repeat` + "`" + ` is ` + "`" + `monthly` + "`" + `, then this should be set to the desired day of the month.`,
				},
				resource.Attribute{
					Name:        "exceptions",
					Description: `(list) Provide a list of exception entries.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_email_server_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"email",
				"server",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group's name.`,
				},
				resource.Attribute{
					Name:        "config_format",
					Description: `(Optional) Config format.`,
				},
				resource.Attribute{
					Name:        "system_format",
					Description: `(Optional) System format.`,
				},
				resource.Attribute{
					Name:        "threat_format",
					Description: `(Optional) Threat format.`,
				},
				resource.Attribute{
					Name:        "traffic_format",
					Description: `(Optional) Traffic format.`,
				},
				resource.Attribute{
					Name:        "hip_match_format",
					Description: `(Optional) HIP match format.`,
				},
				resource.Attribute{
					Name:        "url_format",
					Description: `(Optional) URL format.`,
				},
				resource.Attribute{
					Name:        "data_format",
					Description: `(Optional) Data format.`,
				},
				resource.Attribute{
					Name:        "wildfire_format",
					Description: `(Optional) Wildfire format.`,
				},
				resource.Attribute{
					Name:        "tunnel_format",
					Description: `(Optional) Tunnel format.`,
				},
				resource.Attribute{
					Name:        "user_id_format",
					Description: `(Optional) UserID format.`,
				},
				resource.Attribute{
					Name:        "gtp_format",
					Description: `(Optional) GTP format.`,
				},
				resource.Attribute{
					Name:        "auth_format",
					Description: `(Optional) Auth format.`,
				},
				resource.Attribute{
					Name:        "sctp_format",
					Description: `(Optional) SCTP format.`,
				},
				resource.Attribute{
					Name:        "iptag_format",
					Description: `(Optional) IP tag format.`,
				},
				resource.Attribute{
					Name:        "escaped_characters",
					Description: `(Optional) The escaped characters (as a string).`,
				},
				resource.Attribute{
					Name:        "escape_character",
					Description: `(Optional) The escape character.`,
				},
				resource.Attribute{
					Name:        "email_server",
					Description: `(Required, repeatable) The server spec (defined below). ` + "`" + `email_server` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Server name.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name.`,
				},
				resource.Attribute{
					Name:        "from_email",
					Description: `(Required) From email address.`,
				},
				resource.Attribute{
					Name:        "to_email",
					Description: `(Required) To email address.`,
				},
				resource.Attribute{
					Name:        "also_to_email",
					Description: `(Optional) Secondary to email address.`,
				},
				resource.Attribute{
					Name:        "email_gateway",
					Description: `(Required) The email server.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ethernet_interface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ethernet",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The ethernet interface's name. This should be something like ` + "`" + `ethernet1/X` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Required) The vsys that will use this interface. This should be something like ` + "`" + `vsys1` + "`" + ` or ` + "`" + `vsys3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The interface mode. This can be any of the following values: ` + "`" + `layer3` + "`" + `, ` + "`" + `layer2` + "`" + `, ` + "`" + `virtual-wire` + "`" + `, ` + "`" + `tap` + "`" + `, ` + "`" + `ha` + "`" + `, ` + "`" + `decrypt-mirror` + "`" + `, or ` + "`" + `aggregate-group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable DHCP on this interface.`,
				},
				resource.Attribute{
					Name:        "create_dhcp_default_route",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to create a DHCP default route.`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route_metric",
					Description: `(Optional) The metric for the DHCP default route.`,
				},
				resource.Attribute{
					Name:        "ipv6_enabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable IPv6.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "lldp_enabled",
					Description: `(Optional) Enable LLDP (default: false).`,
				},
				resource.Attribute{
					Name:        "lldp_profile",
					Description: `(Optional) LLDP profile.`,
				},
				resource.Attribute{
					Name:        "lldp_ha_passive_pre_negotiation",
					Description: `(bool) LLDP HA passive pre-negotiation.`,
				},
				resource.Attribute{
					Name:        "lacp_ha_passive_pre_negotiation",
					Description: `(bool) LACP HA passive pre-negotiation.`,
				},
				resource.Attribute{
					Name:        "link_speed",
					Description: `(Optional) Link speed. This can be any of the following: ` + "`" + `10` + "`" + `, ` + "`" + `100` + "`" + `, ` + "`" + `1000` + "`" + `, or ` + "`" + `auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "link_duplex",
					Description: `(Optional) Link duplex setting. This can be ` + "`" + `full` + "`" + `, ` + "`" + `half` + "`" + `, or ` + "`" + `auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "link_state",
					Description: `(Optional) The link state. This can be ` + "`" + `up` + "`" + `, ` + "`" + `down` + "`" + `, or ` + "`" + `auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aggregate_group",
					Description: `(Optional) The aggregate group (applicable for physical firewalls only).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "lacp_port_priority",
					Description: `(int) LACP port priority.`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional, PAN-OS 7.1+) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional, PAN-OS 7.1+) The IPv6 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "decrypt_forward",
					Description: `(Optional, PAN-OS 8.1+) Enable decrypt forwarding.`,
				},
				resource.Attribute{
					Name:        "rx_policing_rate",
					Description: `(Optional, PAN-OS 8.1+) Receive policing rate in Mbps.`,
				},
				resource.Attribute{
					Name:        "tx_policing_rate",
					Description: `(Optional, PAN-OS 8.1+) Transmit policing rate in Mbps.`,
				},
				resource.Attribute{
					Name:        "dhcp_send_hostname_enable",
					Description: `(Optional, PAN-OS 9.0+) For DHCP layer3 interfaces: enable sending the firewall or a custom hostname to DHCP server`,
				},
				resource.Attribute{
					Name:        "dhcp_send_hostname_value",
					Description: `(Optional, PAN-OS 9.0+) For DHCP layer3 interfaces: the interface hostname. Leaving this unspecified with ` + "`" + `dhcp_send_hostname_enable` + "`" + ` set means to send the system hostname.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_file_blocking_security_profile",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"file",
				"blocking",
				"security",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(repeatable) Rule list spec, as defined below. ` + "`" + `rule` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(list) List of applications.`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `(list) List of file types.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `The direction. Valid values are ` + "`" + `both` + "`" + ` (default), ` + "`" + `upload` + "`" + `, or ` + "`" + `download` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action to take. Valid values are ` + "`" + `alert` + "`" + ` (default), ` + "`" + `block` + "`" + `, or ` + "`" + `continue` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_general_settings",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"general",
				"settings",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `Firewall hostname.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `The timezone (e.g. - ` + "`" + `US/Pacific` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain.`,
				},
				resource.Attribute{
					Name:        "update_server",
					Description: `The update server (Default: ` + "`" + `updates.paloaltonetworks.com` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "verify_update_server",
					Description: `Verify update server identity (Default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "proxy_server",
					Description: `(1.5+) Specify a proxy server.`,
				},
				resource.Attribute{
					Name:        "proxy_port",
					Description: `(int, 1.5+) Proxy's port number.`,
				},
				resource.Attribute{
					Name:        "proxy_username",
					Description: `(1.5+) Proxy's username.`,
				},
				resource.Attribute{
					Name:        "proxy_password",
					Description: `(1.5+) Proxy's password.`,
				},
				resource.Attribute{
					Name:        "dns_primary",
					Description: `Primary DNS server.`,
				},
				resource.Attribute{
					Name:        "dns_secondary",
					Description: `Secondary DNS server.`,
				},
				resource.Attribute{
					Name:        "ntp_primary_address",
					Description: `Primary NTP server.`,
				},
				resource.Attribute{
					Name:        "ntp_primary_auth_type",
					Description: `Primary NTP auth type. This can be ` + "`" + `none` + "`" + `, ` + "`" + `autokey` + "`" + `, or ` + "`" + `symmetric-key` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ntp_primary_key_id",
					Description: `Primary NTP ` + "`" + `symmetric-key` + "`" + ` key ID.`,
				},
				resource.Attribute{
					Name:        "ntp_primary_algorithm",
					Description: `Primary NTP ` + "`" + `symmetric-key` + "`" + ` algorithm. This can be ` + "`" + `sha1` + "`" + ` or ` + "`" + `md5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ntp_primary_auth_key",
					Description: `Primary NTP ` + "`" + `symmetric-key` + "`" + ` auth key. This is the SHA1 hash if the algorithm is ` + "`" + `sha1` + "`" + `, or the md5sum if the algorithm is ` + "`" + `md5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ntp_secondary_address",
					Description: `Secondary NTP server.`,
				},
				resource.Attribute{
					Name:        "ntp_secondary_auth_type",
					Description: `Secondary NTP auth type. This can be ` + "`" + `none` + "`" + `, ` + "`" + `autokey` + "`" + `, or ` + "`" + `symmetric-key` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ntp_secondary_key_id",
					Description: `Secondary NTP ` + "`" + `symmetric-key` + "`" + ` key ID.`,
				},
				resource.Attribute{
					Name:        "ntp_secondary_algorithm",
					Description: `Secondary NTP ` + "`" + `symmetric-key` + "`" + ` algorithm. This can be ` + "`" + `sha1` + "`" + ` or ` + "`" + `md5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ntp_secondary_auth_key",
					Description: `Secondary NTP ` + "`" + `symmetric-key` + "`" + ` auth key. This is the SHA1 hash if the algorithm is ` + "`" + `sha1` + "`" + `, or the md5sum if the algorithm is ` + "`" + `md5` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_globalprotect_ipsec_crypto_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"globalprotect",
				"ipsec",
				"crypto",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "authentications",
					Description: `(List of string) The authentication algorithms. Valid values are ` + "`" + `"sha1"` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_gre_tunnel",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"gre",
				"tunnel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The GRE tunnel name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) Interface to terminate tunnel.`,
				},
				resource.Attribute{
					Name:        "local_address_type",
					Description: `(Optional) Type of local address. Valid values are ` + "`" + `ip` + "`" + ` (default) or ` + "`" + `floating-ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_address_value",
					Description: `(Required) IP address value.`,
				},
				resource.Attribute{
					Name:        "peer_address",
					Description: `(Required) Peer IP address.`,
				},
				resource.Attribute{
					Name:        "tunnel_interface",
					Description: `(Required) Tunnel interface to apply the GRE tunnel to.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional, int) Time to live.`,
				},
				resource.Attribute{
					Name:        "copy_tos",
					Description: `(Optional, bool) Copy IP TOS bits from inner packet to GRE packet.`,
				},
				resource.Attribute{
					Name:        "enable_keep_alive",
					Description: `(Optional, bool) Enable tunnel monitoring.`,
				},
				resource.Attribute{
					Name:        "keep_alive_interval",
					Description: `(Optional, int) Keep alive interval.`,
				},
				resource.Attribute{
					Name:        "keep_alive_retry",
					Description: `(Optional, int) Keep alive retry.`,
				},
				resource.Attribute{
					Name:        "keep_alive_hold_timer",
					Description: `(Optional, int) Keep alive hold timer.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional, bool) Disable the GRE tunnel.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_http_server_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"http",
				"server",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group's name.`,
				},
				resource.Attribute{
					Name:        "tag_registration",
					Description: `(Optional, bool) Perform tag registration.`,
				},
				resource.Attribute{
					Name:        "config_format",
					Description: `(Optional) A format folder spec for config (defined below).`,
				},
				resource.Attribute{
					Name:        "system_format",
					Description: `(Optional) A format folder spec for system (defined below).`,
				},
				resource.Attribute{
					Name:        "threat_format",
					Description: `(Optional) A format folder spec for threat (defined below).`,
				},
				resource.Attribute{
					Name:        "traffic_format",
					Description: `(Optional) A format folder spec for traffic (defined below).`,
				},
				resource.Attribute{
					Name:        "hip_match_format",
					Description: `(Optional) A format folder spec for HIP match (defined below).`,
				},
				resource.Attribute{
					Name:        "url_format",
					Description: `(Optional) A format folder spec for url (defined below).`,
				},
				resource.Attribute{
					Name:        "data_format",
					Description: `(Optional) A format folder spec for data (defined below).`,
				},
				resource.Attribute{
					Name:        "wildfire_format",
					Description: `(Optional) A format folder spec for wildfire (defined below).`,
				},
				resource.Attribute{
					Name:        "tunnel_format",
					Description: `(Optional) A format folder spec for tunnel (defined below).`,
				},
				resource.Attribute{
					Name:        "user_id_format",
					Description: `(Optional) A format folder spec for user ID (defined below).`,
				},
				resource.Attribute{
					Name:        "gtp_format",
					Description: `(Optional) A format folder spec for gtp (defined below).`,
				},
				resource.Attribute{
					Name:        "auth_format",
					Description: `(Optional) A format folder spec for auth (defined below).`,
				},
				resource.Attribute{
					Name:        "sctp_format",
					Description: `(Optional, PAN-OS 8.1+) A format folder spec for sctp (defined below).`,
				},
				resource.Attribute{
					Name:        "iptag_format",
					Description: `(Optional, PAN-OS 9.0+) A format folder spec for iptag (defined below).`,
				},
				resource.Attribute{
					Name:        "http_server",
					Description: `(Optional, repeatable) A server spec (defined below). All format folders support the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name.`,
				},
				resource.Attribute{
					Name:        "uri_format",
					Description: `(Optional) The URI format.`,
				},
				resource.Attribute{
					Name:        "payload",
					Description: `(Optional) The payload.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `(Optional, map) A map of HTTP headers and their values.`,
				},
				resource.Attribute{
					Name:        "params",
					Description: `(Optional, map) A map of HTTP params and their values. ` + "`" + `http_server` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The server name.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The server address.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol. Valid values are ` + "`" + `HTTPS` + "`" + ` (default) or ` + "`" + `HTTP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, int) The port number (default: 443).`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `(Optional) The HTTP method (default: ` + "`" + `POST` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The username.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password.`,
				},
				resource.Attribute{
					Name:        "tls_version",
					Description: `(Optional) The TLS version.`,
				},
				resource.Attribute{
					Name:        "certificate_profile",
					Description: `(Optional) The certificate profile.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ike_crypto_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ike",
				"crypto",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "dh_groups",
					Description: `(Required, list) List of DH Group entries. Values should have a prefix if ` + "`" + `group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "authentications",
					Description: `(Required, list) List of authentication types. This c`,
				},
				resource.Attribute{
					Name:        "encryptions",
					Description: `(Required, list) List of encryption types. Valid values are ` + "`" + `des` + "`" + `, ` + "`" + `3des` + "`" + `, ` + "`" + `aes-128-cbc` + "`" + `, ` + "`" + `aes-192-cbc` + "`" + `, ` + "`" + `aes-256-cbc` + "`" + `, ` + "`" + `aes-128-gcm` + "`" + ` (PAN-OS 10.0), and ` + "`" + `aes-256-gcm` + "`" + ` (PAN-OS 10.0).`,
				},
				resource.Attribute{
					Name:        "lifetime_type",
					Description: `The lifetime type. Valid values are ` + "`" + `seconds` + "`" + `, ` + "`" + `minutes` + "`" + `, ` + "`" + `hours` + "`" + ` (the default), and ` + "`" + `days` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifetime_value",
					Description: `(int) The lifetime value.`,
				},
				resource.Attribute{
					Name:        "authentication_multiple",
					Description: `(PAN-OS 7.0+, int) IKEv2 SA reauthentication interval equals authetication-multiple`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ike_gateway",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ike",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional, PAN-OS 7.0+) The IKE gateway version. Valid values are ` + "`" + `ikev1` + "`" + `, (the default), ` + "`" + `ikev2` + "`" + `, or ` + "`" + `ikev2-preferred` + "`" + `. For PAN-OS 6.1, only ` + "`" + `ikev1` + "`" + ` is acceptable.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `(Optional, PAN-OS 7.0+, bool) Enable IPv6 or not.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to disable.`,
				},
				resource.Attribute{
					Name:        "peer_ip_type",
					Description: `(Optional) The peer IP type. Valid values are ` + "`" + `ip` + "`" + `, ` + "`" + `dynamic` + "`" + `, and ` + "`" + `fqdn` + "`" + ` (PANOS 8.1+).`,
				},
				resource.Attribute{
					Name:        "peer_ip_value",
					Description: `(Optional) The peer IP value.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface.`,
				},
				resource.Attribute{
					Name:        "local_ip_address_type",
					Description: `(Optional) The local IP address type. Valid values for this are ` + "`" + `ip` + "`" + `, ` + "`" + `floating-ip` + "`" + `, or an empty string (the default) which is ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_ip_address_value",
					Description: `(Optional) The IP address if ` + "`" + `local_ip_address_type` + "`" + ` is set to ` + "`" + `ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) The auth type. Valid values are ` + "`" + `pre-shared-key` + "`" + ` (the default), or ` + "`" + `certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(Optional) The pre-shared key value.`,
				},
				resource.Attribute{
					Name:        "local_id_type",
					Description: `(Optional) The local ID type. Valid values are ` + "`" + `ipaddr` + "`" + `, ` + "`" + `fqdn` + "`" + `, ` + "`" + `ufqdn` + "`" + `, ` + "`" + `keyid` + "`" + `, or ` + "`" + `dn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_id_value",
					Description: `(Optional) The local ID value.`,
				},
				resource.Attribute{
					Name:        "peer_id_type",
					Description: `(Optional) The peer ID type. Valid values are ` + "`" + `ipaddr` + "`" + `, ` + "`" + `fqdn` + "`" + `, ` + "`" + `ufqdn` + "`" + `, ` + "`" + `keyid` + "`" + `, or ` + "`" + `dn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "peer_id_value",
					Description: `(Optional) The peer ID value.`,
				},
				resource.Attribute{
					Name:        "peer_id_check",
					Description: `(Optional) Enable peer ID wildcard match for certificate authentication. Valid values are ` + "`" + `exact` + "`" + ` or ` + "`" + `wildcard` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_cert",
					Description: `(Optional) The local certificate name.`,
				},
				resource.Attribute{
					Name:        "cert_enable_hash_and_url",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to use hash-and-url for local certificate.`,
				},
				resource.Attribute{
					Name:        "cert_base_url",
					Description: `(Optional) The host and directory part of URL for local certificates.`,
				},
				resource.Attribute{
					Name:        "cert_use_management_as_source",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to use management interface IP as source to retrieve http certificates`,
				},
				resource.Attribute{
					Name:        "cert_permit_payload_mismatch",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to permit peer identification and certificate payload identification mismatch.`,
				},
				resource.Attribute{
					Name:        "cert_profile",
					Description: `(Optional) Profile for certificate valdiation during IKE negotiation.`,
				},
				resource.Attribute{
					Name:        "cert_enable_strict_validation",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable strict validation of peer's extended key use.`,
				},
				resource.Attribute{
					Name:        "enable_passive_mode",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable passive mode (responder only).`,
				},
				resource.Attribute{
					Name:        "enable_nat_traversal",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable NAT traversal.`,
				},
				resource.Attribute{
					Name:        "nat_traversal_keep_alive",
					Description: `(Optional, int) Sending interval for NAT keep-alive packets (in seconds). For versions 6.1 - 8.1, this param, if specified, should be a multiple of 10 between 10 and 3600 to be valid.`,
				},
				resource.Attribute{
					Name:        "nat_traversal_enable_udp_checksum",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable NAT traversal UDP checksum.`,
				},
				resource.Attribute{
					Name:        "enable_fragmentation",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable fragmentation.`,
				},
				resource.Attribute{
					Name:        "ikev1_exchange_mode",
					Description: `(Optional) The IKEv1 exchange mode.`,
				},
				resource.Attribute{
					Name:        "ikev1_crypto_profile",
					Description: `(Optional) IKEv1 crypto profile.`,
				},
				resource.Attribute{
					Name:        "enable_dead_peer_detection",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable dead peer detection.`,
				},
				resource.Attribute{
					Name:        "dead_peer_detection_interval",
					Description: `(Optional, int) The dead peer detection interval.`,
				},
				resource.Attribute{
					Name:        "dead_peer_detection_retry",
					Description: `(Optional, int) Number of retries before disconnection.`,
				},
				resource.Attribute{
					Name:        "ikev2_crypto_profile",
					Description: `(Optional, PAN-OS 7.0+) IKEv2 crypto profile.`,
				},
				resource.Attribute{
					Name:        "ikev2_cookie_validation",
					Description: `(Optional, PAN-OS 7.0+) Set to ` + "`" + `true` + "`" + ` to require cookie.`,
				},
				resource.Attribute{
					Name:        "enable_liveness_check",
					Description: `(Optional, , PAN-OS 7.0+bool) Set to ` + "`" + `true` + "`" + ` to enable sending empty information liveness check message.`,
				},
				resource.Attribute{
					Name:        "liveness_check_interval",
					Description: `(Optional, , PAN-OS 7.0+int) Delay interval before sending probing packets (in seconds).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ip_tag",
			Category:         "User-ID",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"user",
				"id",
				"ip",
				"tag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys location (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP address.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(list) List of tags.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ipsec_crypto_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ipsec",
				"crypto",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol. Valid values are ` + "`" + `esp` + "`" + ` (the default) or ` + "`" + `ah` + "`" + ``,
				},
				resource.Attribute{
					Name:        "authentications",
					Description: `(Required, list) - List of authentication types.`,
				},
				resource.Attribute{
					Name:        "encryptions",
					Description: `(Required, list) - List of encryption types. Valid values are ` + "`" + `des` + "`" + `, ` + "`" + `3des` + "`" + `, ` + "`" + `aes-128-cbc` + "`" + `, ` + "`" + `aes-192-cbc` + "`" + `, ` + "`" + `aes-256-cbc` + "`" + `, ` + "`" + `aes-128-gcm` + "`" + `, ` + "`" + `aes-256-gcm` + "`" + `, and ` + "`" + `null` + "`" + `. Note that the "gcm" values are only available in PAN-OS 7.0+.`,
				},
				resource.Attribute{
					Name:        "dh_group",
					Description: `(Optional) The DH group value. Valid values should start with the string ` + "`" + `group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifetime_type",
					Description: `(Optional) The lifetime type. Valid values are ` + "`" + `seconds` + "`" + `, ` + "`" + `minutes` + "`" + `, ` + "`" + `hours` + "`" + ` (the default), or ` + "`" + `days` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifetime_value",
					Description: `(Optional, int) The lifetime value.`,
				},
				resource.Attribute{
					Name:        "lifesize_type",
					Description: `(Optional) The lifesize type. Valid values are ` + "`" + `kb` + "`" + `, ` + "`" + `mb` + "`" + `, ` + "`" + `gb` + "`" + `, or ` + "`" + `tb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifesize_value",
					Description: `(Optional, int) the lifesize value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ipsec_tunnel",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ipsec",
				"tunnel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "tunnel_interface",
					Description: `(Required) The tunnel interface.`,
				},
				resource.Attribute{
					Name:        "anti_replay",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable Anti-Replay check on this tunnel.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to enable IPv6.`,
				},
				resource.Attribute{
					Name:        "copy_tos",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to copy IP TOS bits from inner packet to IPSec packet (not recommended).`,
				},
				resource.Attribute{
					Name:        "copy_flow_label",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to copy IPv6 flow label for 6in6 tunnel from inner packet to IPSec packet (not recommended).`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to disable this IPSec tunnel.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type. Valid values are ` + "`" + `auto-key` + "`" + ` (the default), ` + "`" + `manual-key` + "`" + `, or ` + "`" + `global-protect-satellite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ak_ike_gateway",
					Description: `(Optional) IKE gateway name.`,
				},
				resource.Attribute{
					Name:        "ak_ipsec_crypto_profile",
					Description: `(Optional) IPSec crypto profile name.`,
				},
				resource.Attribute{
					Name:        "mk_local_spi",
					Description: `(Optional) Outbound SPI, hex format.`,
				},
				resource.Attribute{
					Name:        "mk_remote_spi",
					Description: `(Optional) Inbound SPI, hex format.`,
				},
				resource.Attribute{
					Name:        "mk_local_address_ip",
					Description: `(Optional) Specify exact IP address if interface has multiple addresses.`,
				},
				resource.Attribute{
					Name:        "mk_local_address_floating_ip",
					Description: `(Optional) Floating IP address in HA Active-Active configuration.`,
				},
				resource.Attribute{
					Name:        "mk_protocol",
					Description: `(Optional) Manual key protocol. Valid valies are ` + "`" + `esp` + "`" + ` or ` + "`" + `ah` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mk_auth_type",
					Description: `(Optional) Authentication algorithm. Valid values are ` + "`" + `md5` + "`" + `, ` + "`" + `sha1` + "`" + `, ` + "`" + `sha256` + "`" + `, ` + "`" + `sha384` + "`" + `, ` + "`" + `sha512` + "`" + `, or ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mk_auth_key",
					Description: `(Optional) The auth key for the given auth type.`,
				},
				resource.Attribute{
					Name:        "mk_esp_encryption_type",
					Description: `(Optional) The encryption algorithm. Valid values are ` + "`" + `des` + "`" + `, ` + "`" + `3des` + "`" + `, ` + "`" + `aes-128-cbc` + "`" + `, ` + "`" + `aes-192-cbc` + "`" + `, ` + "`" + `aes-256-cbc` + "`" + `, or ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mk_esp_encryption_key",
					Description: `(Optional) The encryption key.`,
				},
				resource.Attribute{
					Name:        "gps_interface",
					Description: `(Optional) Interface to communicate with portal.`,
				},
				resource.Attribute{
					Name:        "gps_portal_address",
					Description: `(Optional) GlobalProtect portal address.`,
				},
				resource.Attribute{
					Name:        "gps_prefer_ipv6",
					Description: `(Optional, PAN-OS 8.0+, bool) Prefer to register the portal in IPv6. Only applicable to FQDN portal-address.`,
				},
				resource.Attribute{
					Name:        "gps_interface_ip_ipv4",
					Description: `(Optional) specify exact IP address if interface has multiple addresses (IPv4).`,
				},
				resource.Attribute{
					Name:        "gps_interface_ip_ipv6",
					Description: `(Optional, PAN-OS 8.0+) specify exact IP address if interface has multiple addresses (IPv6).`,
				},
				resource.Attribute{
					Name:        "gps_interface_floating_ip_ipv4",
					Description: `(Optional, PAN-OS 7.0+) Floating IPv4 address in HA Active-Active configuration.`,
				},
				resource.Attribute{
					Name:        "gps_interface_floating_ip_ipv6",
					Description: `(Optional, PAN-OS 8.0+) Floating IPv6 address in HA Active-Active configuration.`,
				},
				resource.Attribute{
					Name:        "gps_publish_connected_routes",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to to publish connected and static routes.`,
				},
				resource.Attribute{
					Name:        "gps_publish_routes",
					Description: `(Optional, list) Specify list of routes to publish to Global Protect Gateway.`,
				},
				resource.Attribute{
					Name:        "gps_local_certificate",
					Description: `(Optional) GlobalProtect satellite certificate file name.`,
				},
				resource.Attribute{
					Name:        "gps_certificate_profile",
					Description: `(Optional) Profile for authenticating GlobalProtect gateway certificates.`,
				},
				resource.Attribute{
					Name:        "enable_tunnel_monitor",
					Description: `(Optional, bool) Enable tunnel monitoring on this tunnel.`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_destination_ip",
					Description: `(Optional) Destination IP to send ICMP probe.`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_source_ip",
					Description: `(Optional) Source IP to send ICMP probe`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_profile",
					Description: `(Optional) Tunnel monitor profile.`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_proxy_id",
					Description: `(Optional, PAN-OS 7.0+) Which proxy-id (or proxy-id-v6) the monitoring traffic will use.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ipsec_tunnel_proxy_id_ipv4",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ipsec",
				"tunnel",
				"proxy",
				"id",
				"ipv4",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "ipsec_tunnel",
					Description: `(Required) The auto key IPSec tunnel to attach this proxy ID to.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) IP subnet or IP address represents local network.`,
				},
				resource.Attribute{
					Name:        "remote",
					Description: `(Optional) IP subnet or IP address represents remote network.`,
				},
				resource.Attribute{
					Name:        "protocol_any",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` for any IP protocol.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `(Optional, int) IP protocol number.`,
				},
				resource.Attribute{
					Name:        "protocol_tcp_local",
					Description: `(Optional, int) Local TCP port number.`,
				},
				resource.Attribute{
					Name:        "protocol_tcp_remote",
					Description: `(Optional, int) Remote TCP port number.`,
				},
				resource.Attribute{
					Name:        "protocol_udp_local",
					Description: `(Optional, int) Local UDP port number.`,
				},
				resource.Attribute{
					Name:        "protocol_udp_remote",
					Description: `(Optional, int) Remote UDP port number.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_kerberos_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"kerberos",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. NGFW / Panorama:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "admin_use_only",
					Description: `(bool) Administrator use only.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `List of server specs, as defined below. ` + "`" + `server` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) Server hostname or IP address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(int) Kerberos server port number (default: ` + "`" + `88` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_layer2_subinterface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"layer2",
				"subinterface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "interface_type",
					Description: `(Optional) The interface type. Valid values are ` + "`" + `ethernet` + "`" + ` (default) or ` + "`" + `aggregate-ethernet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parent_interface",
					Description: `(Required) The name of the parent interface.`,
				},
				resource.Attribute{
					Name:        "parent_mode",
					Description: `(Optional) The parent's mode. Valid values are ` + "`" + `layer2` + "`" + ` (default) or ` + "`" + `virtual-wire` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Required) The vsys that will use this interface. This should be something like ` + "`" + `vsys1` + "`" + ` or ` + "`" + `vsys3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional, int) The interface's tag.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_layer3_subinterface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"layer3",
				"subinterface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "interface_type",
					Description: `(Optional) The interface type. Valid values are ` + "`" + `ethernet` + "`" + ` (default) or ` + "`" + `aggregate-ethernet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parent_interface",
					Description: `(Required) The name of the parent interface.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Required) The vsys that will use this interface. This should be something like ` + "`" + `vsys1` + "`" + ` or ` + "`" + `vsys3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional, int) The interface's tag.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses.`,
				},
				resource.Attribute{
					Name:        "ipv6_enabled",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable IPv6.`,
				},
				resource.Attribute{
					Name:        "ipv6_interface_id",
					Description: `(Optional) The IPv6 interface ID.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional) The IPv6 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable DHCP.`,
				},
				resource.Attribute{
					Name:        "create_dhcp_default_route",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to create a DHCP default route.`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route_metric",
					Description: `(Optional) The metric for the DHCP default route.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "decrypt_forward",
					Description: `(Optional, bool, PAN-OS 8.1+) Set to ` + "`" + `true` + "`" + ` to enable decrypt forward.`,
				},
				resource.Attribute{
					Name:        "dhcp_send_hostname_enable",
					Description: `(Optional, PAN-OS 9.0+) For DHCP layer3 interfaces: enable sending the firewall or a custom hostname to DHCP server`,
				},
				resource.Attribute{
					Name:        "dhcp_send_hostname_value",
					Description: `(Optional, PAN-OS 9.0+) For DHCP layer3 interfaces: the interface hostname. Leaving this unspecified with ` + "`" + `dhcp_send_hostname_enable` + "`" + ` set means to send the system hostname.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ldap_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"ldap",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. NGFW / Panorama:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "admin_use_only",
					Description: `(bool) Administrator use only.`,
				},
				resource.Attribute{
					Name:        "ldap_type",
					Description: `LDAP type. Valid values are ` + "`" + `"active-directory"` + "`" + `, ` + "`" + `"e-directory"` + "`" + `, ` + "`" + `"sun"` + "`" + `, or ` + "`" + `"other"` + "`" + ` (default).`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `(bool) SSL (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "verify_server_certificate",
					Description: `(bool) Verify server certificate for SSL sessions.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(bool) Disable this profile.`,
				},
				resource.Attribute{
					Name:        "base_dn",
					Description: `Default base distinguished name (DN) to use for searches.`,
				},
				resource.Attribute{
					Name:        "bind_dn",
					Description: `Bind distinguished name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Bind password.`,
				},
				resource.Attribute{
					Name:        "search_timeout",
					Description: `(int) Number of seconds to wait for performing searches (default: ` + "`" + `30` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "bind_timeout",
					Description: `(int) Number of seconds to use for connecting to servers (default: ` + "`" + `30` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "retry_interval",
					Description: `(int) Interval (in seconds) for reconnecting LDAP server (default: ` + "`" + `60` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `List of server specs, as defined below. ` + "`" + `server` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) Server hostname or IP address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(int) LDAP server port number (default: ` + "`" + `389` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_license_api_key",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"license",
				"api",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The licensing API key.`,
				},
				resource.Attribute{
					Name:        "retain_key",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to retain the licensing API key even after the deletion of this resource (recommended).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_licensing",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"licensing",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auth_codes",
					Description: `(Required) The list of auth codes to install.`,
				},
				resource.Attribute{
					Name:        "delicense",
					Description: `(Optional, bool) Leave as ` + "`" + `true` + "`" + ` if you want to delicense the firewall when this resource is removed, otherwise set to ` + "`" + `false` + "`" + ` to prevent firewall delicensing. Delicensing requires that the licensing API key has been installed.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) For ` + "`" + `delicense` + "`" + ` of ` + "`" + `true` + "`" + `, the type of delicensing to perform. Right now, only ` + "`" + `auto` + "`" + ` is supported (no manual delicensing). ## Attribute Reference The following attributes are available after read operations:`,
				},
				resource.Attribute{
					Name:        "licenses",
					Description: `List of licenses. Licenses have the following attributes:`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `The feature name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `License description.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial number.`,
				},
				resource.Attribute{
					Name:        "issued",
					Description: `When the license was issued.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `When the license expires.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `If the license has expired or not.`,
				},
				resource.Attribute{
					Name:        "auth_code",
					Description: `Associated auth code (if applicable).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "licenses",
					Description: `List of licenses. Licenses have the following attributes:`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `The feature name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `License description.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial number.`,
				},
				resource.Attribute{
					Name:        "issued",
					Description: `When the license was issued.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `When the license expires.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `If the license has expired or not.`,
				},
				resource.Attribute{
					Name:        "auth_code",
					Description: `Associated auth code (if applicable).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_local_user_db_group",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"local",
				"user",
				"db",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. NGFW:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `List of users in this group.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_local_user_db_user",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"local",
				"user",
				"db",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. NGFW:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(bool) If the user is disabled or not.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_log_forwarding_profile",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"log",
				"forwarding",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "enhanced_logging",
					Description: `(Optional, bool, PAN-OS 8.1+) Set to ` + "`" + `true` + "`" + ` to enable enhanced logging.`,
				},
				resource.Attribute{
					Name:        "match_list",
					Description: `(Optional, repeatable) The match list spec (defined below). ` + "`" + `match_list` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `(Optional) The log type. Valid values are ` + "`" + `traffic` + "`" + ` (default), ` + "`" + `threat` + "`" + `, ` + "`" + `wildfire` + "`" + `, ` + "`" + `url` + "`" + `, ` + "`" + `data` + "`" + `, ` + "`" + `gtp` + "`" + `, ` + "`" + `tunnel` + "`" + `, ` + "`" + `auth` + "`" + `, ` + "`" + `sctp` + "`" + `, or ` + "`" + `decryption` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) The filter (default: ` + "`" + `All Logs` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "send_to_panorama",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to send to Panorama.`,
				},
				resource.Attribute{
					Name:        "snmptrap_server_profiles",
					Description: `(Optional) List of SNMP server profiles.`,
				},
				resource.Attribute{
					Name:        "email_server_profiles",
					Description: `(Optional) List of email server profiles.`,
				},
				resource.Attribute{
					Name:        "syslog_server_profiles",
					Description: `(Optional) List of syslog server profiles.`,
				},
				resource.Attribute{
					Name:        "http_server_profiles",
					Description: `(Optional) List of http server profiles.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional, repeatable) Match list action spec (defined below). ` + "`" + `match_list.action` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "azure_integration",
					Description: `(Optional) The Azure integration spec (defined below). Mutually exclusive with ` + "`" + `tagging_integration` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tagging_integration",
					Description: `(Optional) The tagging integration spec (defined below). Mutually exclusive with ` + "`" + `azure_integration` + "`" + `. ` + "`" + `match_list.action.azure_integration` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "azure_integration",
					Description: `(Optional, bool) This param defaults to ` + "`" + `true` + "`" + ` and should not be changed. ` + "`" + `match_list.action.tagging_integration` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The action. Valid values are ` + "`" + `add-tag` + "`" + ` (default) or ` + "`" + `remove-tag` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) The target. Valid values are ` + "`" + `source-address` + "`" + ` (default) or ` + "`" + `destination-address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional, int) The timeout.`,
				},
				resource.Attribute{
					Name:        "local_registration",
					Description: `(Optional) The local registration spec (defined below). Only one of ` + "`" + `local_registration` + "`" + `, ` + "`" + `remote_registration` + "`" + `, or ` + "`" + `panorama_registration` + "`" + ` should be defined.`,
				},
				resource.Attribute{
					Name:        "remote_registration",
					Description: `(Optional) The remote registration spec (defined below). Only one of ` + "`" + `local_registration` + "`" + `, ` + "`" + `remote_registration` + "`" + `, or ` + "`" + `panorama_registration` + "`" + ` should be defined.`,
				},
				resource.Attribute{
					Name:        "panorama_registration",
					Description: `(Optional) The panorama registration spec (defined below). Only one of ` + "`" + `local_registration` + "`" + `, ` + "`" + `remote_registration` + "`" + `, or ` + "`" + `panorama_registration` + "`" + ` should be defined. ` + "`" + `match_list.action.tagging_integration.local_registration` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Required) List of administrative tags. ` + "`" + `match_list.action.tagging_integration.remote_registration` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Required) List of administrative tags.`,
				},
				resource.Attribute{
					Name:        "http_profile",
					Description: `(Required) The HTTP profile. ` + "`" + `match_list.action.tagging_integration.panorama_registration` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Required) List of administrative tags.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_loopback_interface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"loopback",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name. This must start with ` + "`" + `loopback.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys that will use this interface (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional, bool) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv6 MSS adjust value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_management_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"management",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The management profile's name.`,
				},
				resource.Attribute{
					Name:        "ping",
					Description: `(Optional) Allow ping.`,
				},
				resource.Attribute{
					Name:        "telnet",
					Description: `(Optional) Allow telnet.`,
				},
				resource.Attribute{
					Name:        "ssh",
					Description: `(Optional) Allow SSH.`,
				},
				resource.Attribute{
					Name:        "http",
					Description: `(Optional) Allow HTTP.`,
				},
				resource.Attribute{
					Name:        "http_ocsp",
					Description: `(Optional) Allow HTTP OCSP.`,
				},
				resource.Attribute{
					Name:        "https",
					Description: `(Optional) Allow HTTPS.`,
				},
				resource.Attribute{
					Name:        "snmp",
					Description: `(Optional) Allow SNMP.`,
				},
				resource.Attribute{
					Name:        "response_pages",
					Description: `(Optional) Allow response pages.`,
				},
				resource.Attribute{
					Name:        "userid_service",
					Description: `(Optional) Allow User ID service.`,
				},
				resource.Attribute{
					Name:        "userid_syslog_listener_ssl",
					Description: `(Optional) Allow User ID syslog listener for SSL.`,
				},
				resource.Attribute{
					Name:        "userid_syslog_listener_udp",
					Description: `(Optional) Allow User ID syslog listener for UDP.`,
				},
				resource.Attribute{
					Name:        "permitted_ips",
					Description: `(Optional) The list of permitted IP addresses or address ranges for this management profile.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_monitor_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"monitor",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The monitor profile name.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional, int) The probing interval in seconds.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Optional, int) The number of failed probes to determine that the tunnel is down.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action triggered when tunnel's status changes. Valid values are ` + "`" + `wait-recover` + "`" + ` (default) or ` + "`" + `fail-over` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_nat_rule",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"nat",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The NAT rule's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the NAT rule into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rulebase",
					Description: `(Optional, Deprecated) The rulebase. For firewalls, there is only the ` + "`" + `rulebase` + "`" + ` value (default), but on Panorama, there is also ` + "`" + `pre-rulebase` + "`" + ` and ` + "`" + `post-rulebase` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional). NAT type. This can be ` + "`" + `ipv4` + "`" + ` (default), ` + "`" + `nat64` + "`" + `, or ` + "`" + `nptv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `(Required) The list of source zone(s).`,
				},
				resource.Attribute{
					Name:        "destination_zone",
					Description: `(Required) The destination zone.`,
				},
				resource.Attribute{
					Name:        "to_interface",
					Description: `(Optional) Egress interface from route lookup (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Service (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Required) List of source address(es).`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Required) List of destination address(es).`,
				},
				resource.Attribute{
					Name:        "sat_type",
					Description: `(Optional) Type of source address translation. This can be ` + "`" + `none` + "`" + ` (default), ` + "`" + `dynamic-ip-and-port` + "`" + `, ` + "`" + `dynamic-ip` + "`" + `, or ` + "`" + `static-ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_address_type",
					Description: `(Optional) Source address translation address type. This can be ` + "`" + `interface-address` + "`" + ` or ` + "`" + `translated-address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_translated_addresses",
					Description: `(Optional) Source address translation list of translated addresses.`,
				},
				resource.Attribute{
					Name:        "sat_interface",
					Description: `(Optional) Source address translation interface.`,
				},
				resource.Attribute{
					Name:        "sat_ip_address",
					Description: `(Optional) Source address translation IP address.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_type",
					Description: `(Optional) Source address translation fallback type. This can be ` + "`" + `none` + "`" + `, ` + "`" + `interface-address` + "`" + `, or ` + "`" + `translated-address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_translated_addresses",
					Description: `(Optional) Source address translation list of fallback translated addresses.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_interface",
					Description: `(Optional) Source address translation fallback interface.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_ip_type",
					Description: `(Optional) Source address translation fallback IP type. This can be ` + "`" + `ip` + "`" + ` or ` + "`" + `floating` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_ip_address",
					Description: `(Optional) The source address translation fallback IP address.`,
				},
				resource.Attribute{
					Name:        "sat_static_translated_address",
					Description: `(Optional) The statically translated source address.`,
				},
				resource.Attribute{
					Name:        "sat_static_bi_directional",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable bi-directional source address translation.`,
				},
				resource.Attribute{
					Name:        "dat_type",
					Description: `(Optional) Destination address translation type. This should be either ` + "`" + `static` + "`" + ` or ` + "`" + `dynamic` + "`" + `. The ` + "`" + `dynamic` + "`" + ` option is only available on PAN-OS 8.1+.`,
				},
				resource.Attribute{
					Name:        "dat_address",
					Description: `(Optional) Destination address translation's address. Requires ` + "`" + `dat_type` + "`" + ` be set to "static" or "dynamic".`,
				},
				resource.Attribute{
					Name:        "dat_port",
					Description: `(Optional) Destination address translation's port number. Requires ` + "`" + `dat_type` + "`" + ` be set to "static" or "dynamic".`,
				},
				resource.Attribute{
					Name:        "dat_dynamic_distribution",
					Description: `(Optional, PAN-OS 8.1+) Distribution algorithm for destination address pool. The PAN-OS 8.1 GUI doesn't seem to set this anywhere, but this is added here for completeness' sake. Requires ` + "`" + `dat_type` + "`" + ` of "dynamic".`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to disable this rule.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_nat_rule_group",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"nat",
				"rule",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_group",
					Description: `The device group (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rulebase",
					Description: `The rulebase. This can be ` + "`" + `pre-rulebase` + "`" + ` (default), ` + "`" + `post-rulebase` + "`" + `, or ` + "`" + `rulebase` + "`" + `. NGFW specific arguments:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `(Optional) A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `(Optional) Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Repeatable) The rule definition (see below). The rule ordering will match how they appear in the terraform plan file. Each ` + "`" + `rule` + "`" + ` defined supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The NAT rule's name.`,
				},
				resource.Attribute{
					Name:        "audit_comment",
					Description: `When this rule is created/updated, the audit comment to apply for this rule.`,
				},
				resource.Attribute{
					Name:        "group_tag",
					Description: `(PAN-OS 9.0+) The group tag.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `NAT type. This can be ` + "`" + `ipv4` + "`" + ` (default), ` + "`" + `nat64` + "`" + `, or ` + "`" + `nptv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List of administrative tags.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to disable this rule.`,
				},
				resource.Attribute{
					Name:        "original_packet",
					Description: `(Required) The original packet specification (see below).`,
				},
				resource.Attribute{
					Name:        "translated_packet",
					Description: `(Required) The translated packet spec (see below).`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(repeatable, Panorama only) A target definition (see below). If there are no target sections, then the rule will apply to every vsys of every device in the device group.`,
				},
				resource.Attribute{
					Name:        "negate_target",
					Description: `(bool, Panorama only) Instead of applying the rule for the given serial numbers, apply it to everything except them. ` + "`" + `original_packet` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `(Required) The list of source zone(s).`,
				},
				resource.Attribute{
					Name:        "destination_zone",
					Description: `(Required) The destination zone.`,
				},
				resource.Attribute{
					Name:        "destination_interface",
					Description: `Egress interface from route lookup (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Required) List of source address(es).`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Required) List of destination address(es). ` + "`" + `translated_packet` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The source spec (see below). Leave this empty for a destination NAT of "none".`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) The destination spec (see below). Leave this empty for a destination NAT of "none". ` + "`" + `translated_packet.source` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "dynamic_ip_and_port",
					Description: `Dynamic IP and port source translation spec (see below).`,
				},
				resource.Attribute{
					Name:        "dynamic_ip",
					Description: `Dynamic IP source translation spec (see below).`,
				},
				resource.Attribute{
					Name:        "static_ip",
					Description: `Static IP source translation spec (see below). ` + "`" + `translated_packet.source.dynamic_ip_and_port` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `Translated address source translation type spec (see below).`,
				},
				resource.Attribute{
					Name:        "interface_address",
					Description: `Interface address source translation type spec (see below). ` + "`" + `translated_packet.source.dynamic_ip_and_port.translated_address` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `(Required) List of translated addresses. ` + "`" + `translated_packet.source.dynamic_ip_and_port.interface_address` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address. ` + "`" + `translated_packet.source.dynamic_ip` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `The list of translated addresses.`,
				},
				resource.Attribute{
					Name:        "fallback",
					Description: `The fallback spec (see below). Leaving this empty or omiting it means a fallback of "None". ` + "`" + `translated_packet.source.dynamic_ip.fallback` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `The translated address fallback spec (see below).`,
				},
				resource.Attribute{
					Name:        "interface_address",
					Description: `The interface address fallback spec (see below). ` + "`" + `translated_packet.source.dynamic_ip.fallback.translated_address` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `List of source address translation fallback translated addresses. ` + "`" + `translated_packet.source.dynamic_ip.fallback.interface_address` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) Source address translation fallback interface.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of interface fallback. Valid values are ` + "`" + `ip` + "`" + ` (default) or ` + "`" + `floating` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address of the fallback interface. ` + "`" + `translated_packet.source.static_ip` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `(Required) The statically translated source address.`,
				},
				resource.Attribute{
					Name:        "bi_directional",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to enable bi-directional source address translation. ` + "`" + `translated_packet.destination` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "static_translation",
					Description: `Specifies a static destination NAT (see below).`,
				},
				resource.Attribute{
					Name:        "dynamic_translation",
					Description: `(PAN-OS 8.1+) Specify a dynamic destination NAT (see below).`,
				},
				resource.Attribute{
					Name:        "static",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "dynamic",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Destination address translation address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(int) Destination address translation port number. ` + "`" + `translated_packet.destination.dynamic` + "`" + ` and ` + "`" + `translated_packet.destination.dynamic_translation` + "`" + ` support the following arguments:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Destination address translation address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(int) Destination address translation port number.`,
				},
				resource.Attribute{
					Name:        "distribution",
					Description: `(PAN-OS 8.1+) Distribution algorithm for destination address pool. The PAN-OS 8.1 GUI doesn't seem to set this anywhere, but this is added here for completeness' sake. The GUI sets this to ` + "`" + `round-robin` + "`" + ` currently. ` + "`" + `target` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `A listing of vsys to apply this rule to. If ` + "`" + `serial` + "`" + ` is a VM, then this parameter should just be omitted. ## Attribute Reference Each ` + "`" + `rule` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(PAN-OS 9.0+) The PAN-OS UUID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `(PAN-OS 9.0+) The PAN-OS UUID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ospf",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ospf",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Optional, but Required for Panorama) The template location. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(bool) Enable flag.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `Router ID.`,
				},
				resource.Attribute{
					Name:        "reject_default_route",
					Description: `(bool) Reject default route.`,
				},
				resource.Attribute{
					Name:        "allow_redistribute_default_route",
					Description: `(bool) Allow redistribute default route.`,
				},
				resource.Attribute{
					Name:        "rfc_1583",
					Description: `(bool) RFC 1583.`,
				},
				resource.Attribute{
					Name:        "spf_calculation_delay",
					Description: `(float) SPF calculation delay.`,
				},
				resource.Attribute{
					Name:        "lsa_interval",
					Description: `(float) LSA interval.`,
				},
				resource.Attribute{
					Name:        "enable_graceful_restart",
					Description: `(bool) Enable graceful restart.`,
				},
				resource.Attribute{
					Name:        "grace_period",
					Description: `(int) Grace period.`,
				},
				resource.Attribute{
					Name:        "helper_enable",
					Description: `(bool) Helper enable.`,
				},
				resource.Attribute{
					Name:        "strict_lsa_checking",
					Description: `(bool) Strict LSA checking.`,
				},
				resource.Attribute{
					Name:        "max_neighbor_restart_time",
					Description: `(int) Max neighbor restart time.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `BFD profile name.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ospf_area",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ospf",
				"area",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Optional, but Required for Panorama) The template location. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Area type. Valid values are ` + "`" + `normal` + "`" + ` (default), ` + "`" + `stub` + "`" + `, or ` + "`" + `nssa` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "accept_summary",
					Description: `(bool) (stub/nssa) Accept summary.`,
				},
				resource.Attribute{
					Name:        "default_route_advertise",
					Description: `(bool) (stub/nssa) Default route advertise.`,
				},
				resource.Attribute{
					Name:        "advertise_metric",
					Description: `(int) (stub/nssa) Advertise metric.`,
				},
				resource.Attribute{
					Name:        "advertise_type",
					Description: `(nssa) Advertise type. Valid values are ` + "`" + `ext-2` + "`" + ` or ` + "`" + `ext-1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ext_range",
					Description: `(repeatable) (nssa) EXT range spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "range",
					Description: `(repeatable) Range spec, as defined below. ` + "`" + `ext_range` + "`" + ` and ` + "`" + `range` + "`" + ` both support the following arguments:`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Network.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action. Valid values are ` + "`" + `advertise` + "`" + ` (default) or ` + "`" + `suppress` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ospf_area_interface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ospf",
				"area",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Optional, but Required for Panorama) The template location. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router name.`,
				},
				resource.Attribute{
					Name:        "ospf_area",
					Description: `(Required) OSPF area name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Interface name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(bool) Enable (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "passive",
					Description: `(bool) Passive.`,
				},
				resource.Attribute{
					Name:        "link_type",
					Description: `Link type. Valid values are ` + "`" + `broadcast` + "`" + ` (default), ` + "`" + `p2p` + "`" + `, or ` + "`" + `p2mp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(int) Metric (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(int) Priority (default: ` + "`" + `1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `(int) Hello interval in seconds (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "dead_counts",
					Description: `(int) Dead counts (default: ` + "`" + `4` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "retransmit_interval",
					Description: `(int) Retransmit interval in seconds (default: ` + "`" + `5` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "transit_delay",
					Description: `(int) Transit delay in seconds (default: ` + "`" + `1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "grace_restart_delay",
					Description: `(int) Graceful restart hello delay in seconds (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "auth_profile",
					Description: `Auth profile.`,
				},
				resource.Attribute{
					Name:        "neighbors",
					Description: `(list) (p2mp) List of neighbors.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `BFD profile.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ospf_area_virtual_link",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ospf",
				"area",
				"virtual",
				"link",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Optional, but Required for Panorama) The template location. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router name.`,
				},
				resource.Attribute{
					Name:        "ospf_area",
					Description: `(Required) OSPF area name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Interface name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(bool) Enable (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "neighbor_id",
					Description: `(Required) Neighbor ID.`,
				},
				resource.Attribute{
					Name:        "transit_area_id",
					Description: `(Required) Transit area ID.`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `(int) Hello interval in seconds (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "dead_counts",
					Description: `(int) Dead counts (default: ` + "`" + `4` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "retransmit_interval",
					Description: `(int) Retransmit interval in seconds (default: ` + "`" + `5` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "transit_delay",
					Description: `(int) Transit delay in seconds (default: ` + "`" + `1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "auth_profile",
					Description: `Auth profile.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `BFD profile.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ospf_auth_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ospf",
				"auth",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Optional, but Required for Panorama) The template location. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The export rule name.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `The auth type. Valid values are ` + "`" + `password` + "`" + ` (default) or ` + "`" + `md5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The simple password.`,
				},
				resource.Attribute{
					Name:        "md5_key",
					Description: `(list) List of md5_key specs, as defined below. ` + "`" + `md5_key` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Required, int) MD5 key ID.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) MD5 key.`,
				},
				resource.Attribute{
					Name:        "preferred",
					Description: `(bool) Preferred key. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "password_enc",
					Description: `Encrypted simple password.`,
				},
				resource.Attribute{
					Name:        "md5_keys_enc",
					Description: `(list) List of encrypted/unencrypted MD5 keys.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "password_enc",
					Description: `Encrypted simple password.`,
				},
				resource.Attribute{
					Name:        "md5_keys_enc",
					Description: `(list) List of encrypted/unencrypted MD5 keys.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ospf_export",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ospf",
				"export",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Optional, but Required for Panorama) The template location. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The export rule name.`,
				},
				resource.Attribute{
					Name:        "path_type",
					Description: `Path type. Valid values are ` + "`" + `ext-2` + "`" + ` (default) or ` + "`" + `ext-1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Tag.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(int) Metric.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_address_group",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"panorama",
				"address",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The address group's name.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group to put the address group into (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_addresses",
					Description: `(Optional) The address objects to include in this statically defined address group.`,
				},
				resource.Attribute{
					Name:        "dynamic_match",
					Description: `(Optional) The IP tags to include in this DAG.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The address group's description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_administrative_tag",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"panorama",
				"administrative",
				"tag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The administrative tag's name.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group to put the administrative tag into (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) The tag's color. This should be either an empty string (no color) or a string such as ` + "`" + `color1` + "`" + ` or ` + "`" + `color15` + "`" + `. Note that for maximum portability, you should limit color usage to ` + "`" + `color16` + "`" + `, which was available in PAN-OS 6.1. PAN-OS 8.1's colors go up to ` + "`" + `color42` + "`" + `. The value ` + "`" + `color18` + "`" + ` is reserved internally by PAN-OS and thus not available for use.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The administrative tag's description.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_aggregate_interface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"aggregate",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name.`,
				},
				resource.Attribute{
					Name:        "templaet",
					Description: `(Required) The template where the interface should be.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Required) The vsys that will use this interface. This should be something like ` + "`" + `vsys1` + "`" + ` or ` + "`" + `vsys3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The interface mode. Valid values are ` + "`" + `layer3` + "`" + ` (default), ` + "`" + `layer2` + "`" + `, ` + "`" + `virtual-wire` + "`" + `, ` + "`" + `ha` + "`" + `, or ` + "`" + `decrypt-mirror` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional) The IPv6 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "enable_untagged_subinterface",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable untagged subinterfaces.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses.`,
				},
				resource.Attribute{
					Name:        "ipv6_enabled",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable IPv6.`,
				},
				resource.Attribute{
					Name:        "ipv6_interface_id",
					Description: `(Optional) The IPv6 interface ID.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable DHCP.`,
				},
				resource.Attribute{
					Name:        "create_dhcp_default_route",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to create a DHCP default route.`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route_metric",
					Description: `(Optional) The metric for the DHCP default route.`,
				},
				resource.Attribute{
					Name:        "lacp_enable",
					Description: `(bool) Enable LACP.`,
				},
				resource.Attribute{
					Name:        "lacp_fast_failover",
					Description: `(bool) Enable LACP fast failover.`,
				},
				resource.Attribute{
					Name:        "lacp_mode",
					Description: `LACP mode. Valid values are ` + "`" + `active` + "`" + ` or ` + "`" + `passive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lacp_transmission_rate",
					Description: `LACP transmission rate. Valid values are ` + "`" + `fast` + "`" + ` or ` + "`" + `slow` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lacp_system_priority",
					Description: `(int) LACP system priority.`,
				},
				resource.Attribute{
					Name:        "lacp_max_ports",
					Description: `(int) LACP max ports.`,
				},
				resource.Attribute{
					Name:        "lacp_ha_passive_pre_negotiation",
					Description: `(bool) LACP HA passive pre-negotiation.`,
				},
				resource.Attribute{
					Name:        "lacp_ha_enable_same_system_mac",
					Description: `(bool) LACP HA enable same system MAC.`,
				},
				resource.Attribute{
					Name:        "lacp_ha_same_system_mac_address",
					Description: `LACP HA same system MAC address.`,
				},
				resource.Attribute{
					Name:        "lldp_enable",
					Description: `(bool) Enable LLDP.`,
				},
				resource.Attribute{
					Name:        "lldp_profile",
					Description: `LLDP profile name.`,
				},
				resource.Attribute{
					Name:        "lldp_ha_passive_pre_negotiation",
					Description: `(bool) LLDP HA passive pre-negotiation.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "decrypt_forward",
					Description: `(Optional, bool, PAN-OS 8.1+) Set to ` + "`" + `true` + "`" + ` to enable decrypt forward.`,
				},
				resource.Attribute{
					Name:        "dhcp_send_hostname_enable",
					Description: `(Optional, PAN-OS 9.0+) For DHCP layer3 interfaces: enable sending the firewall or a custom hostname to DHCP server`,
				},
				resource.Attribute{
					Name:        "dhcp_send_hostname_value",
					Description: `(Optional, PAN-OS 9.0+) For DHCP layer3 interfaces: the interface hostname. Leaving this unspecified with ` + "`" + `dhcp_send_hostname_enable` + "`" + ` set means to send the system hostname.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_application_group",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"panorama",
				"application",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The group's device group (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group's name.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(Optional) List of applications in this group.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_application_object",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"panorama",
				"application",
				"object",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_application_signature",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"panorama",
				"application",
				"signature",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The signature's name.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group (default: ` + "`" + `shared` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "application_object",
					Description: `(Required) The applciation object for this signature.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) The signature's scope. Valid values are ` + "`" + `transaction` + "`" + ` (default) or ` + "`" + `session` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ordered_match",
					Description: `(Optional, bool) Set to ` + "`" + `false` + "`" + ` to disable ordered matching (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "and_condition",
					Description: `(Optional) The and condition spec (defined below). ` + "`" + `and_condition` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) And condition name, this is computed and cannot be configured.`,
				},
				resource.Attribute{
					Name:        "or_condition",
					Description: `(Required) The or condition spec (defined below). ` + "`" + `and_condition.or_condition` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) Or condition name, this is computed and cannot be configured.`,
				},
				resource.Attribute{
					Name:        "pattern_match",
					Description: `(Optional) The pattern match spec (defined below).`,
				},
				resource.Attribute{
					Name:        "greater_than",
					Description: `(Optional) The greater than spec (defined below).`,
				},
				resource.Attribute{
					Name:        "less_than",
					Description: `(Optional) the less than spec (defined below).`,
				},
				resource.Attribute{
					Name:        "equal_to",
					Description: `(Optional) The equal to spec (defined below). ` + "`" + `and_condition.or_condition.pattern_match` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "context",
					Description: `(Required) The context.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Required) The pattern.`,
				},
				resource.Attribute{
					Name:        "qualifiers",
					Description: `(Optional, map) The qualifiers. ` + "`" + `and_condition.or_condition.greater_than` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "context",
					Description: `(Required) The context.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value.`,
				},
				resource.Attribute{
					Name:        "qualifiers",
					Description: `(Optional, map) The qualifiers. ` + "`" + `and_condition.or_condition.less_than` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "context",
					Description: `(Required) The context.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value.`,
				},
				resource.Attribute{
					Name:        "qualifiers",
					Description: `(Optional, map) The qualifiers. ` + "`" + `and_condition.or_condition.equal_to` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "context",
					Description: `(Required) The context.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Optional) The position.`,
				},
				resource.Attribute{
					Name:        "mask",
					Description: `(Optional) The mask.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bfd_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bfd",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The BBFD profile's name.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) BFD operation mode. Valid values are ` + "`" + `active` + "`" + ` (default) or ` + "`" + `passive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minimum_tx_interval",
					Description: `(Optional, int) Desired minimum TX interval in ms. Default is ` + "`" + `1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minimum_rx_interval",
					Description: `(Optional, int) Required minimum RX interval in ms. Default is ` + "`" + `1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "detection_multiplier",
					Description: `(Optional, int) Multiplier sent to remote system. Default is ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hold_time",
					Description: `(Optional, int) Delay transmission and reception of control packets in ms.`,
				},
				resource.Attribute{
					Name:        "minimum_rx_ttl",
					Description: `(Optional, int) Minimum accepted ttl on received BFD packet.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bgp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP configuration to.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable BGP or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Optional) Router ID of this BGP instance.`,
				},
				resource.Attribute{
					Name:        "as_number",
					Description: `(Optional) Local AS number.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `(Optional, PAN-OS 7.1+) BFD configuration.`,
				},
				resource.Attribute{
					Name:        "reject_default_route",
					Description: `(Optional, bool) Do not learn default route from BGP (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "install_route",
					Description: `(Optional, bool) Populate BGP learned route to global route table.`,
				},
				resource.Attribute{
					Name:        "aggregate_med",
					Description: `(Optional, bool) Aggregate route only if they have same MED attributes (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "default_local_preference",
					Description: `(Optional) Default local preference (default: ` + "`" + `"100"` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_format",
					Description: `(Optional) AS format. Valid values are ` + "`" + `"2-byte"` + "`" + ` (default) or ` + "`" + `"4-byte"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "always_compare_med",
					Description: `(Optional, bool) Always compare MEDs.`,
				},
				resource.Attribute{
					Name:        "deterministic_med_comparison",
					Description: `(Optional, bool) Deterministic MED comparison (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ecmp_multi_as",
					Description: `(Optional, bool) Support multiple AS in ECMP.`,
				},
				resource.Attribute{
					Name:        "enforce_first_as",
					Description: `(Optional, bool) Enforce First AS for EBGP (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_graceful_restart",
					Description: `(Optional, bool) Enable graceful restart (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "stale_route_time",
					Description: `(Optional, int) Time to remove stale routes after peer restart, in seconds (default: ` + "`" + `120` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "local_restart_time",
					Description: `(Optional, int) Local restart time to advertise to peer, in seconds (default: ` + "`" + `120` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "max_peer_restart_time",
					Description: `(Optional, int) Maximum of peer restart time accepted, in seconds (default: ` + "`" + `120` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "reflector_cluster_id",
					Description: `(Optional) Route reflector cluster ID.`,
				},
				resource.Attribute{
					Name:        "confederation_member_as",
					Description: `(Optional) Confederation requires member-AS number.`,
				},
				resource.Attribute{
					Name:        "allow_redistribute_default_route",
					Description: `(Optional, bool) Allow redistribute default route to BGP.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_aggregate",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bgp",
				"aggregate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to put the rule into.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The rule name.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Aggregating address prefix.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this rule (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "as_set",
					Description: `(Optional, bool) Generate AS-set attribute.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `(Optional, bool) Summarize route.`,
				},
				resource.Attribute{
					Name:        "local_preference",
					Description: `(Optional) New local preference value.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) New MED value.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional, int) New weight value.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop address.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) New route origin. Valid values are ` + "`" + `incomplete` + "`" + ` (default), ` + "`" + `igp` + "`" + `, or ` + "`" + `egp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_limit",
					Description: `(Optional, int) Add AS path limit attribute if it does not exist.`,
				},
				resource.Attribute{
					Name:        "as_path_type",
					Description: `(Optional) AS path update options. Valid values are ` + "`" + `none` + "`" + ` (default) or ` + "`" + `prepend` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_value",
					Description: `(Optional) For ` + "`" + `as_path_type` + "`" + ` of ` + "`" + `prepend` + "`" + `, the value to prepend.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community update options. Valid values are ` + "`" + `none` + "`" + ` (default), ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_value",
					Description: `(Optional) If ` + "`" + `community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. For the ` + "`" + `append` + "`" + ` and ` + "`" + `overwrite` + "`" + ` types specifically, valid values are ` + "`" + `no-export` + "`" + `, ` + "`" + `no-advertise` + "`" + `, ` + "`" + `local-as` + "`" + `, or ` + "`" + `nopeer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_type",
					Description: `(Optional) Extended community update options. Valid values are ` + "`" + `none` + "`" + ` (default), ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_vaule",
					Description: `(Optional) If ` + "`" + `extended_community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_aggregate_advertise_filter",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bgp",
				"aggregate",
				"advertise",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_aggregate",
					Description: `(Required) The BGP address aggregation rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below).`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry. Each ` + "`" + `address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_aggregate_suppress_filter",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bgp",
				"aggregate",
				"suppress",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_aggregate",
					Description: `(Required) The BGP address aggregation rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below).`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry. Each ` + "`" + `address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_auth_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bgp",
				"auth",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP auth profile to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Shared secret for the TCP MD5 authentication.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_conditional_adv",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bgp",
				"conditional",
				"adv",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP conditional advertisement to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "used_by",
					Description: `(Optional) List of BGP peer groups that use this rule.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_conditional_adv_advertise_filter",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bgp",
				"conditional",
				"adv",
				"advertise",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_conditional_adv",
					Description: `(Required) The BGP conditional advertisement to add this filter to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "address_prefixes",
					Description: `(Optional) List of matching address prefixes.`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_conditional_adv_non_exist_filter",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bgp",
				"conditional",
				"adv",
				"non",
				"exist",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this filter to.`,
				},
				resource.Attribute{
					Name:        "bgp_conditional_adv",
					Description: `(Required) The BGP conditional advertisement to add this filter to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "address_prefixes",
					Description: `(Optional) List of matching address prefixes.`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_dampening_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bgp",
				"dampening",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP dampening profile to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "cutoff",
					Description: `(Optional, float) Cutoff threshold value (default: ` + "`" + `1.25` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "reuse",
					Description: `(Optional, float) Reuse threshold value (default: ` + "`" + `0.5` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "max_hold_time",
					Description: `(Optional, int) Maximum hold-down time, in seconds (default: ` + "`" + `900` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "decay_half_life_reachable",
					Description: `(Optional, int) Decay half-life while reachable, in seconds (default: ` + "`" + `300` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "decay_half_life_unreachable",
					Description: `(Optional, int) Decay half-life while unreachable, in seconds (default: ` + "`" + `900` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_export_rule_group",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bgp",
				"export",
				"rule",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to put the rule into.`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `(Optional) A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `(Optional) Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The export rule definition (see below). The export rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The security rule name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this export rule (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "used_by",
					Description: `(Optional) List of auth profiles.`,
				},
				resource.Attribute{
					Name:        "match_as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "match_community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "match_extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "match_med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "match_route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "match_address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below). below for the params for this section.`,
				},
				resource.Attribute{
					Name:        "match_next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "match_from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Rule action. Valid values are ` + "`" + `allow` + "`" + ` (default) or ` + "`" + `deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dampening",
					Description: `(Optional) Route flap dampening profile.`,
				},
				resource.Attribute{
					Name:        "local_preference",
					Description: `(Optional) New local preference value.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) New MED value.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional, int) New weight value.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop address.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) New route origin. Valid values are ` + "`" + `igp` + "`" + `, ` + "`" + `egp` + "`" + `, or ` + "`" + `incomplete` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_limit",
					Description: `(Optional, int) Add AS path limit attribute if it does not exist.`,
				},
				resource.Attribute{
					Name:        "as_path_type",
					Description: `(Optional) AS path update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove` + "`" + `, ` + "`" + `prepend` + "`" + ` or ` + "`" + `remove-and-prepend` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_value",
					Description: `(Optional) If ` + "`" + `as_path_type` + "`" + ` is ` + "`" + `prepend` + "`" + ` or ` + "`" + `remove-and-prepend` + "`" + `, the value to prepend.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_value",
					Description: `(Optional) If ` + "`" + `community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. For the ` + "`" + `append` + "`" + ` and ` + "`" + `overwrite` + "`" + ` types specifically, valid values for ` + "`" + `community_value` + "`" + ` are ` + "`" + `no-export` + "`" + `, ` + "`" + `no-advertise` + "`" + `, ` + "`" + `local-as` + "`" + `, or ` + "`" + `nopeer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_type",
					Description: `(Optional) Extended community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_vaule",
					Description: `(Optional) If ` + "`" + `extended_community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. Each ` + "`" + `match_address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_import_rule_group",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bgp",
				"import",
				"rule",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to put the rule into.`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `(Optional) A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `(Optional) Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The import rule definition (see below). The import rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The security rule name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this import rule (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "used_by",
					Description: `(Optional) List of auth profiles.`,
				},
				resource.Attribute{
					Name:        "match_as_path_regex",
					Description: `(Optional) AS path to match.`,
				},
				resource.Attribute{
					Name:        "match_community_regex",
					Description: `(Optional) Community to match.`,
				},
				resource.Attribute{
					Name:        "match_extended_community_regex",
					Description: `(Optional) Extended community to match.`,
				},
				resource.Attribute{
					Name:        "match_med",
					Description: `(Optional) Match MED.`,
				},
				resource.Attribute{
					Name:        "match_route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "match_address_prefix",
					Description: `(Optional, repeatable) Matching address prefix definition (see below). below for the params for this section.`,
				},
				resource.Attribute{
					Name:        "match_next_hops",
					Description: `(Optional) List of next hop attributes.`,
				},
				resource.Attribute{
					Name:        "match_from_peers",
					Description: `(Optional) List of peers that advertised the route entry.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Rule action. Valid values are ` + "`" + `allow` + "`" + ` (default) or ` + "`" + `deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dampening",
					Description: `(Optional) Route flap dampening profile.`,
				},
				resource.Attribute{
					Name:        "local_preference",
					Description: `(Optional) New local preference value.`,
				},
				resource.Attribute{
					Name:        "med",
					Description: `(Optional) New MED value.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional, int) New weight value.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) Next hop address.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) New route origin. Valid values are ` + "`" + `igp` + "`" + `, ` + "`" + `egp` + "`" + `, or ` + "`" + `incomplete` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_path_limit",
					Description: `(Optional, int) Add AS path limit attribute if it does not exist.`,
				},
				resource.Attribute{
					Name:        "as_path_type",
					Description: `(Optional) AS path update options. Valid values are ` + "`" + `none` + "`" + ` or ` + "`" + `remove` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_type",
					Description: `(Optional) Community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "community_value",
					Description: `(Optional) If ` + "`" + `community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. For the ` + "`" + `append` + "`" + ` and ` + "`" + `overwrite` + "`" + ` types specifically, valid values for ` + "`" + `community_value` + "`" + ` are ` + "`" + `no-export` + "`" + `, ` + "`" + `no-advertise` + "`" + `, ` + "`" + `local-as` + "`" + `, or ` + "`" + `nopeer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_type",
					Description: `(Optional) Extended community update options. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `remove-all` + "`" + `, ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_community_vaule",
					Description: `(Optional) If ` + "`" + `extended_community_type` + "`" + ` is ` + "`" + `remove-regex` + "`" + `, ` + "`" + `append` + "`" + `, or ` + "`" + `overwrite` + "`" + `, the value associated with that setting. Each ` + "`" + `match_address_prefix` + "`" + ` section offers the following params:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Required) Address prefix.`,
				},
				resource.Attribute{
					Name:        "exact",
					Description: `(Optional, bool) Match exact prefix length.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_peer",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bgp",
				"peer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP peer to.`,
				},
				resource.Attribute{
					Name:        "bgp_peer_group",
					Description: `(Required) The BGP peer group to put this peer into.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "peer_as",
					Description: `(Optional) Peer AS number.`,
				},
				resource.Attribute{
					Name:        "local_address_interface",
					Description: `(Required) Interface to accept BGP session.`,
				},
				resource.Attribute{
					Name:        "local_address_ip",
					Description: `(Optional) Specify exact IP address if interface has multiple addresses.`,
				},
				resource.Attribute{
					Name:        "peer_address_ip",
					Description: `(Required) Peer IP address configuration.`,
				},
				resource.Attribute{
					Name:        "reflector_client",
					Description: `(Optional) This peer is reflector client. Valid values are ` + "`" + `non-client` + "`" + `, ` + "`" + `client` + "`" + `, or ` + "`" + `meshed-client` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "peering_type",
					Description: `(Optional) Peering type that affects NOPEER community value handling. Valid values are ` + "`" + `unspecified` + "`" + ` (default) or ` + "`" + `bilateral` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_prefixes",
					Description: `(Optional) Maximum of prefixes to receive from the peer. This can be a number such as ` + "`" + `"5000"` + "`" + ` (default) or ` + "`" + `unlimited` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_profile",
					Description: `(Optional) Auth profile.`,
				},
				resource.Attribute{
					Name:        "keep_alive_interval",
					Description: `(Optional, int) Keep alive interval, in seconds (default: ` + "`" + `30` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "multi_hop",
					Description: `(Optional, int) IP TTL value used for sending BGP packet.`,
				},
				resource.Attribute{
					Name:        "open_delay_time",
					Description: `(Optional, int) Open delay time, in seconds.`,
				},
				resource.Attribute{
					Name:        "hold_time",
					Description: `(Optional, int) Hold time, in seconds.`,
				},
				resource.Attribute{
					Name:        "idle_hold_time",
					Description: `(Optional, int) Idle hold time, in seconds.`,
				},
				resource.Attribute{
					Name:        "allow_incoming_connections",
					Description: `(Optional, bool) Allow incoming connections (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "incoming_connections_remote_port",
					Description: `(Optional, int) Restrict remote port for incoming BGP connections.`,
				},
				resource.Attribute{
					Name:        "allow_outgoing_connections",
					Description: `(Optional, bool) Allow outgoing connections (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "outgoing_connections_local_port",
					Description: `(Optional, int) Use specific local port for outgoing BGP connections.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `(Optional, PAN-OS 7.1+) BFD profile. This can be a specific BFD profile name, ` + "`" + `None` + "`" + ` (disables BFD), or ` + "`" + `Inherit-vr-global-setting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_mp_bgp",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable MP BGP.`,
				},
				resource.Attribute{
					Name:        "address_family_type",
					Description: `(Optional, PAN-OS 8.0+) Set the AFI for this peer. Valid values are ` + "`" + `ipv4` + "`" + ` or ` + "`" + `ipv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subsequent_address_family_unicast",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable unicast subsequent address family for this peer.`,
				},
				resource.Attribute{
					Name:        "subsequent_address_family_multicast",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable multicast subsequent address family for this peer.`,
				},
				resource.Attribute{
					Name:        "enable_sender_side_loop_detection",
					Description: `(Optional, bool, PAN-OS 8.0+) Enable sender side loop detection.`,
				},
				resource.Attribute{
					Name:        "min_route_advertisement_interval",
					Description: `(Optional, int, PAN-OS 8.1+) Minimum route advertisement interval, in seconds.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_peer_group",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bgp",
				"peer",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP peer group to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "aggregated_confed_as_path",
					Description: `(Optional, bool) The peers understand aggregated confederation AS path (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "soft_reset_with_stored_info",
					Description: `(Optional, bool) Soft reset with stored info.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Peer group type. Valid options are ` + "`" + `ebgp` + "`" + ` (default), ` + "`" + `ebgp-confed` + "`" + `, ` + "`" + `ibgp` + "`" + `, or ` + "`" + `ibgp-confed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "export_next_hop",
					Description: `(Optional) Export next hop. Valid values are ` + "`" + `original` + "`" + `, ` + "`" + `use-self` + "`" + `, or ` + "`" + `resolve` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "import_next_hop",
					Description: `(Optional) Import next hop. Valid values are ` + "`" + `original` + "`" + `, ` + "`" + `use-peer` + "`" + `, or the empty string.`,
				},
				resource.Attribute{
					Name:        "remove_private_as",
					Description: `(Optional, bool) Remove private AS when exporting route. Only available for ` + "`" + `type=ebgp` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_bgp_redist_rule",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"bgp",
				"redist",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add this BGP redist rule to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A subnet or a redistribution profile.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional, bool) Enable this rule or not (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `(Optional) The address family. Valid values are ` + "`" + `ipv4` + "`" + ` (default) or ` + "`" + `ipv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional, PAN-OS 8.0+) Route table to match rule. Valid values are ` + "`" + `unicast` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `. As of PAN-OS 8.1, there doesn't seem to be a way to configure this in the GUI, it is always set to ` + "`" + `unicast` + "`" + `. Thus, if you're running this resource against PAN-OS 8.0+, the appropriate thing to do is set this value to ` + "`" + `unicast` + "`" + ` as well to match the GUI functionality.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional, int) Metric value.`,
				},
				resource.Attribute{
					Name:        "set_origin",
					Description: `(Optional) Add the origin path attribute. Valid values are ` + "`" + `incomplete` + "`" + ` (default), ` + "`" + `igp` + "`" + `, or ` + "`" + `egp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "set_med",
					Description: `(Optional) Add the MULTI_EXIT_DISC path attribute.`,
				},
				resource.Attribute{
					Name:        "set_local_preference",
					Description: `(Optional) Add the LOCAL_PREF path attribute.`,
				},
				resource.Attribute{
					Name:        "set_as_path_limit",
					Description: `(Optional, int) Add the AS_PATHLIMIT path attribute.`,
				},
				resource.Attribute{
					Name:        "set_communities",
					Description: `(Optional) List of COMMUNITY path attributes to add.`,
				},
				resource.Attribute{
					Name:        "set_extended_communities",
					Description: `(Optional) List of EXTENDED COMMUNITY path attributes to add.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_device_group",
			Category:         "Panorama",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"device",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_device_group_entry",
			Category:         "Panorama",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"device",
				"group",
				"entry",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_edl",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"panorama",
				"edl",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_email_server_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"panorama",
				"email",
				"server",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) The template location. Mutually exclusive with ` + "`" + `template_stack` + "`" + ` and ` + "`" + `device_group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `(Optional) The template stack location. Mutually exclusive with ` + "`" + `template` + "`" + ` and ` + "`" + `device_group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location. Mutually exclusive with ` + "`" + `template` + "`" + ` and ` + "`" + `template_stack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys. This will likely be ` + "`" + `shared` + "`" + `, and it should be defined if you specified either ` + "`" + `template` + "`" + ` or ` + "`" + `template_stack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group's name.`,
				},
				resource.Attribute{
					Name:        "config_format",
					Description: `(Optional) Config format.`,
				},
				resource.Attribute{
					Name:        "system_format",
					Description: `(Optional) System format.`,
				},
				resource.Attribute{
					Name:        "threat_format",
					Description: `(Optional) Threat format.`,
				},
				resource.Attribute{
					Name:        "traffic_format",
					Description: `(Optional) Traffic format.`,
				},
				resource.Attribute{
					Name:        "hip_match_format",
					Description: `(Optional) HIP match format.`,
				},
				resource.Attribute{
					Name:        "url_format",
					Description: `(Optional) URL format.`,
				},
				resource.Attribute{
					Name:        "data_format",
					Description: `(Optional) Data format.`,
				},
				resource.Attribute{
					Name:        "wildfire_format",
					Description: `(Optional) Wildfire format.`,
				},
				resource.Attribute{
					Name:        "tunnel_format",
					Description: `(Optional) Tunnel format.`,
				},
				resource.Attribute{
					Name:        "user_id_format",
					Description: `(Optional) UserID format.`,
				},
				resource.Attribute{
					Name:        "gtp_format",
					Description: `(Optional) GTP format.`,
				},
				resource.Attribute{
					Name:        "auth_format",
					Description: `(Optional) Auth format.`,
				},
				resource.Attribute{
					Name:        "sctp_format",
					Description: `(Optional) SCTP format.`,
				},
				resource.Attribute{
					Name:        "iptag_format",
					Description: `(Optional) IP tag format.`,
				},
				resource.Attribute{
					Name:        "escaped_characters",
					Description: `(Optional) The escaped characters (as a string).`,
				},
				resource.Attribute{
					Name:        "escape_character",
					Description: `(Optional) The escape character.`,
				},
				resource.Attribute{
					Name:        "email_server",
					Description: `(Required, repeatable) The server spec (defined below). ` + "`" + `email_server` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Server name.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name.`,
				},
				resource.Attribute{
					Name:        "from_email",
					Description: `(Required) From email address.`,
				},
				resource.Attribute{
					Name:        "to_email",
					Description: `(Required) To email address.`,
				},
				resource.Attribute{
					Name:        "also_to_email",
					Description: `(Optional) Secondary to email address.`,
				},
				resource.Attribute{
					Name:        "email_gateway",
					Description: `(Required) The email server.`,
				},
				resource.Attribute{
					Name:        "email_server",
					Description: `(Required, repeatable) The server spec (defined below).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_ethernet_interface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"ethernet",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The ethernet interface's name. This should be something like ` + "`" + `ethernet1/X` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys that will use this interface (default: ` + "`" + `vsys1` + "`" + `). This should be something like ` + "`" + `vsys1` + "`" + ` or ` + "`" + `vsys3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The interface mode. This can be any of the following values: ` + "`" + `layer3` + "`" + `, ` + "`" + `layer2` + "`" + `, ` + "`" + `virtual-wire` + "`" + `, ` + "`" + `tap` + "`" + `, ` + "`" + `ha` + "`" + `, ` + "`" + `decrypt-mirror` + "`" + `, or ` + "`" + `aggregate-group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable DHCP on this interface.`,
				},
				resource.Attribute{
					Name:        "create_dhcp_default_route",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to create a DHCP default route.`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route_metric",
					Description: `(Optional) The metric for the DHCP default route.`,
				},
				resource.Attribute{
					Name:        "ipv6_enabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable IPv6.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "lldp_enabled",
					Description: `(Optional) Enable LLDP (default: false).`,
				},
				resource.Attribute{
					Name:        "lldp_profile",
					Description: `(Optional) LLDP profile.`,
				},
				resource.Attribute{
					Name:        "lldp_ha_passive_pre_negotiation",
					Description: `(bool) LLDP HA passive pre-negotiation.`,
				},
				resource.Attribute{
					Name:        "lacp_ha_passive_pre_negotiation",
					Description: `(bool) LACP HA passive pre-negotiation.`,
				},
				resource.Attribute{
					Name:        "link_speed",
					Description: `(Optional) Link speed. This can be any of the following: ` + "`" + `10` + "`" + `, ` + "`" + `100` + "`" + `, ` + "`" + `1000` + "`" + `, or ` + "`" + `auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "link_duplex",
					Description: `(Optional) Link duplex setting. This can be ` + "`" + `full` + "`" + `, ` + "`" + `half` + "`" + `, or ` + "`" + `auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "link_state",
					Description: `(Optional) The link state. This can be ` + "`" + `up` + "`" + `, ` + "`" + `down` + "`" + `, or ` + "`" + `auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aggregate_group",
					Description: `(Optional) The aggregate group (applicable for physical firewalls only).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "lacp_port_priority",
					Description: `(int) LACP port priority.`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional, PAN-OS 7.1+) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional, PAN-OS 7.1+) The IPv6 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "decrypt_forward",
					Description: `(Optional, PAN-OS 8.1+) Enable decrypt forwarding.`,
				},
				resource.Attribute{
					Name:        "rx_policing_rate",
					Description: `(Optional, PAN-OS 8.1+) Receive policing rate in Mbps.`,
				},
				resource.Attribute{
					Name:        "tx_policing_rate",
					Description: `(Optional, PAN-OS 8.1+) Transmit policing rate in Mbps.`,
				},
				resource.Attribute{
					Name:        "dhcp_send_hostname_enable",
					Description: `(Optional, PAN-OS 9.0+) For DHCP layer3 interfaces: enable sending the firewall or a custom hostname to DHCP server`,
				},
				resource.Attribute{
					Name:        "dhcp_send_hostname_value",
					Description: `(Optional, PAN-OS 9.0+) For DHCP layer3 interfaces: the interface hostname. Leaving this unspecified with ` + "`" + `dhcp_send_hostname_enable` + "`" + ` set means to send the system hostname.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_gcp_account",
			Category:         "Plugins",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"plugins",
				"panorama",
				"gcp",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The account's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Account description.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The GCP project ID.`,
				},
				resource.Attribute{
					Name:        "service_account_credential_type",
					Description: `(Optional) The service account credential type. Valid values are ` + "`" + `gcp` + "`" + ` (default) or ` + "`" + `gke` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "credential_file",
					Description: `(Required) The contents of a GCP credentials file; use the ` + "`" + `file()` + "`" + ` function to pass in the credentials file.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_gke_cluster",
			Category:         "Plugins",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"plugins",
				"panorama",
				"gke",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The cluster group's name.`,
				},
				resource.Attribute{
					Name:        "gke_cluster_group",
					Description: `(Required) The cluster group name.`,
				},
				resource.Attribute{
					Name:        "gcp_zone",
					Description: `(Required) The GCP zone.`,
				},
				resource.Attribute{
					Name:        "cluster_credential",
					Description: `(Required) The GKE account to use.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_gke_cluster_group",
			Category:         "Plugins",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"plugins",
				"panorama",
				"gke",
				"cluster",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The cluster group's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "gcp_project_credential",
					Description: `(Required) The GCP account to use.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Required) The device group.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `(Required) The template stack.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_gre_tunnel",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"gre",
				"tunnel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The GRE tunnel name.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) Interface to terminate tunnel.`,
				},
				resource.Attribute{
					Name:        "local_address_type",
					Description: `(Optional) Type of local address. Valid values are ` + "`" + `ip` + "`" + ` (default) or ` + "`" + `floating-ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_address_value",
					Description: `(Required) IP address value.`,
				},
				resource.Attribute{
					Name:        "peer_address",
					Description: `(Required) Peer IP address.`,
				},
				resource.Attribute{
					Name:        "tunnel_interface",
					Description: `(Required) Tunnel interface to apply the GRE tunnel to.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional, int) Time to live.`,
				},
				resource.Attribute{
					Name:        "copy_tos",
					Description: `(Optional, bool) Copy IP TOS bits from inner packet to GRE packet.`,
				},
				resource.Attribute{
					Name:        "enable_keep_alive",
					Description: `(Optional, bool) Enable tunnel monitoring.`,
				},
				resource.Attribute{
					Name:        "keep_alive_interval",
					Description: `(Optional, int) Keep alive interval.`,
				},
				resource.Attribute{
					Name:        "keep_alive_retry",
					Description: `(Optional, int) Keep alive retry.`,
				},
				resource.Attribute{
					Name:        "keep_alive_hold_timer",
					Description: `(Optional, int) Keep alive hold timer.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional, bool) Disable the GRE tunnel.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_http_server_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"panorama",
				"http",
				"server",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) The template location. Mutually exclusive with ` + "`" + `template_stack` + "`" + ` and ` + "`" + `device_group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `(Optional) The template stack location. Mutually exclusive with ` + "`" + `template` + "`" + ` and ` + "`" + `device_group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location. Mutually exclusive with ` + "`" + `template` + "`" + ` and ` + "`" + `template_stack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys. This will likely be ` + "`" + `shared` + "`" + `, and it should be defined if you specified either ` + "`" + `template` + "`" + ` or ` + "`" + `template_stack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group's name.`,
				},
				resource.Attribute{
					Name:        "tag_registration",
					Description: `(Optional, bool) Perform tag registration.`,
				},
				resource.Attribute{
					Name:        "config_format",
					Description: `(Optional) A format folder spec for config (defined below).`,
				},
				resource.Attribute{
					Name:        "system_format",
					Description: `(Optional) A format folder spec for system (defined below).`,
				},
				resource.Attribute{
					Name:        "threat_format",
					Description: `(Optional) A format folder spec for threat (defined below).`,
				},
				resource.Attribute{
					Name:        "traffic_format",
					Description: `(Optional) A format folder spec for traffic (defined below).`,
				},
				resource.Attribute{
					Name:        "hip_match_format",
					Description: `(Optional) A format folder spec for HIP match (defined below).`,
				},
				resource.Attribute{
					Name:        "url_format",
					Description: `(Optional) A format folder spec for url (defined below).`,
				},
				resource.Attribute{
					Name:        "data_format",
					Description: `(Optional) A format folder spec for data (defined below).`,
				},
				resource.Attribute{
					Name:        "wildfire_format",
					Description: `(Optional) A format folder spec for wildfire (defined below).`,
				},
				resource.Attribute{
					Name:        "tunnel_format",
					Description: `(Optional) A format folder spec for tunnel (defined below).`,
				},
				resource.Attribute{
					Name:        "user_id_format",
					Description: `(Optional) A format folder spec for user ID (defined below).`,
				},
				resource.Attribute{
					Name:        "gtp_format",
					Description: `(Optional) A format folder spec for gtp (defined below).`,
				},
				resource.Attribute{
					Name:        "auth_format",
					Description: `(Optional) A format folder spec for auth (defined below).`,
				},
				resource.Attribute{
					Name:        "sctp_format",
					Description: `(Optional, PAN-OS 8.1+) A format folder spec for sctp (defined below).`,
				},
				resource.Attribute{
					Name:        "iptag_format",
					Description: `(Optional, PAN-OS 9.0+) A format folder spec for iptag (defined below).`,
				},
				resource.Attribute{
					Name:        "http_server",
					Description: `(Optional, repeatable) A server spec (defined below). All format folders support the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name.`,
				},
				resource.Attribute{
					Name:        "uri_format",
					Description: `(Optional) The URI format.`,
				},
				resource.Attribute{
					Name:        "payload",
					Description: `(Optional) The payload.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `(Optional, map) A map of HTTP headers and their values.`,
				},
				resource.Attribute{
					Name:        "params",
					Description: `(Optional, map) A map of HTTP params and their values. ` + "`" + `http_server` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The server name.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The server address.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol. Valid values are ` + "`" + `HTTPS` + "`" + ` (default) or ` + "`" + `HTTP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, int) The port number (default: 443).`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `(Optional) The HTTP method (default: ` + "`" + `POST` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The username.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password.`,
				},
				resource.Attribute{
					Name:        "tls_version",
					Description: `(Optional) The TLS version.`,
				},
				resource.Attribute{
					Name:        "certificate_profile",
					Description: `(Optional) The certificate profile.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_ike_crypto_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"ike",
				"crypto",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_ike_gateway",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"ike",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional, PAN-OS 7.0+) The IKE gateway version. Valid values are ` + "`" + `ikev1` + "`" + `, (the default), ` + "`" + `ikev2` + "`" + `, or ` + "`" + `ikev2-preferred` + "`" + `. For PAN-OS 6.1, only ` + "`" + `ikev1` + "`" + ` is acceptable.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `(Optional, PAN-OS 7.0+, bool) Enable IPv6 or not.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to disable.`,
				},
				resource.Attribute{
					Name:        "peer_ip_type",
					Description: `(Optional) The peer IP type. Valid values are ` + "`" + `ip` + "`" + `, ` + "`" + `dynamic` + "`" + `, and ` + "`" + `fqdn` + "`" + ` (PANOS 8.1+).`,
				},
				resource.Attribute{
					Name:        "peer_ip_value",
					Description: `(Optional) The peer IP value.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface.`,
				},
				resource.Attribute{
					Name:        "local_ip_address_type",
					Description: `(Optional) The local IP address type. Valid values for this are ` + "`" + `ip` + "`" + `, ` + "`" + `floating-ip` + "`" + `, or an empty string (the default) which is ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_ip_address_value",
					Description: `(Optional) The IP address if ` + "`" + `local_ip_address_type` + "`" + ` is set to ` + "`" + `ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) The auth type. Valid values are ` + "`" + `pre-shared-key` + "`" + ` (the default), or ` + "`" + `certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(Optional) The pre-shared key value.`,
				},
				resource.Attribute{
					Name:        "local_id_type",
					Description: `(Optional) The local ID type. Valid values are ` + "`" + `ipaddr` + "`" + `, ` + "`" + `fqdn` + "`" + `, ` + "`" + `ufqdn` + "`" + `, ` + "`" + `keyid` + "`" + `, or ` + "`" + `dn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_id_value",
					Description: `(Optional) The local ID value.`,
				},
				resource.Attribute{
					Name:        "peer_id_type",
					Description: `(Optional) The peer ID type. Valid values are ` + "`" + `ipaddr` + "`" + `, ` + "`" + `fqdn` + "`" + `, ` + "`" + `ufqdn` + "`" + `, ` + "`" + `keyid` + "`" + `, or ` + "`" + `dn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "peer_id_value",
					Description: `(Optional) The peer ID value.`,
				},
				resource.Attribute{
					Name:        "peer_id_check",
					Description: `(Optional) Enable peer ID wildcard match for certificate authentication. Valid values are ` + "`" + `exact` + "`" + ` or ` + "`" + `wildcard` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_cert",
					Description: `(Optional) The local certificate name.`,
				},
				resource.Attribute{
					Name:        "cert_enable_hash_and_url",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to use hash-and-url for local certificate.`,
				},
				resource.Attribute{
					Name:        "cert_base_url",
					Description: `(Optional) The host and directory part of URL for local certificates.`,
				},
				resource.Attribute{
					Name:        "cert_use_management_as_source",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to use management interface IP as source to retrieve http certificates`,
				},
				resource.Attribute{
					Name:        "cert_permit_payload_mismatch",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to permit peer identification and certificate payload identification mismatch.`,
				},
				resource.Attribute{
					Name:        "cert_profile",
					Description: `(Optional) Profile for certificate valdiation during IKE negotiation.`,
				},
				resource.Attribute{
					Name:        "cert_enable_strict_validation",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable strict validation of peer's extended key use.`,
				},
				resource.Attribute{
					Name:        "enable_passive_mode",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable passive mode (responder only).`,
				},
				resource.Attribute{
					Name:        "enable_nat_traversal",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable NAT traversal.`,
				},
				resource.Attribute{
					Name:        "nat_traversal_keep_alive",
					Description: `(Optional, int) Sending interval for NAT keep-alive packets (in seconds). For versions 6.1 - 8.1, this param, if specified, should be a multiple of 10 between 10 and 3600 to be valid.`,
				},
				resource.Attribute{
					Name:        "nat_traversal_enable_udp_checksum",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable NAT traversal UDP checksum.`,
				},
				resource.Attribute{
					Name:        "enable_fragmentation",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable fragmentation.`,
				},
				resource.Attribute{
					Name:        "ikev1_exchange_mode",
					Description: `(Optional) The IKEv1 exchange mode.`,
				},
				resource.Attribute{
					Name:        "ikev1_crypto_profile",
					Description: `(Optional) IKEv1 crypto profile.`,
				},
				resource.Attribute{
					Name:        "enable_dead_peer_detection",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable dead peer detection.`,
				},
				resource.Attribute{
					Name:        "dead_peer_detection_interval",
					Description: `(Optional, int) The dead peer detection interval.`,
				},
				resource.Attribute{
					Name:        "dead_peer_detection_retry",
					Description: `(Optional, int) Number of retries before disconnection.`,
				},
				resource.Attribute{
					Name:        "ikev2_crypto_profile",
					Description: `(Optional, PAN-OS 7.0+) IKEv2 crypto profile.`,
				},
				resource.Attribute{
					Name:        "ikev2_cookie_validation",
					Description: `(Optional, PAN-OS 7.0+) Set to ` + "`" + `true` + "`" + ` to require cookie.`,
				},
				resource.Attribute{
					Name:        "enable_liveness_check",
					Description: `(Optional, , PAN-OS 7.0+bool) Set to ` + "`" + `true` + "`" + ` to enable sending empty information liveness check message.`,
				},
				resource.Attribute{
					Name:        "liveness_check_interval",
					Description: `(Optional, , PAN-OS 7.0+int) Delay interval before sending probing packets (in seconds).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_ipsec_crypto_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"ipsec",
				"crypto",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol. Valid values are ` + "`" + `esp` + "`" + ` (the default) or ` + "`" + `ah` + "`" + ``,
				},
				resource.Attribute{
					Name:        "authentications",
					Description: `(Required, list) - List of authentication types.`,
				},
				resource.Attribute{
					Name:        "encryptions",
					Description: `(Required, list) - List of encryption types. Valid values are ` + "`" + `des` + "`" + `, ` + "`" + `3des` + "`" + `, ` + "`" + `aes-128-cbc` + "`" + `, ` + "`" + `aes-192-cbc` + "`" + `, ` + "`" + `aes-256-cbc` + "`" + `, ` + "`" + `aes-128-gcm` + "`" + `, ` + "`" + `aes-256-gcm` + "`" + `, and ` + "`" + `null` + "`" + `. Note that the "gcm" values are only available in PAN-OS 7.0+.`,
				},
				resource.Attribute{
					Name:        "dh_group",
					Description: `(Optional) The DH group value. Valid values should start with the string ` + "`" + `group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifetime_type",
					Description: `(Optional) The lifetime type. Valid values are ` + "`" + `seconds` + "`" + `, ` + "`" + `minutes` + "`" + `, ` + "`" + `hours` + "`" + ` (the default), or ` + "`" + `days` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifetime_value",
					Description: `(Optional, int) The lifetime value.`,
				},
				resource.Attribute{
					Name:        "lifesize_type",
					Description: `(Optional) The lifesize type. Valid values are ` + "`" + `kb` + "`" + `, ` + "`" + `mb` + "`" + `, ` + "`" + `gb` + "`" + `, or ` + "`" + `tb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifesize_value",
					Description: `(Optional, int) the lifesize value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_ipsec_tunnel",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"ipsec",
				"tunnel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "tunnel_interface",
					Description: `(Required) The tunnel interface.`,
				},
				resource.Attribute{
					Name:        "anti_replay",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable Anti-Replay check on this tunnel.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to enable IPv6.`,
				},
				resource.Attribute{
					Name:        "copy_tos",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to copy IP TOS bits from inner packet to IPSec packet (not recommended).`,
				},
				resource.Attribute{
					Name:        "copy_flow_label",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to copy IPv6 flow label for 6in6 tunnel from inner packet to IPSec packet (not recommended).`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional, PAN-OS 7.0+, bool) Set to ` + "`" + `true` + "`" + ` to disable this IPSec tunnel.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type. Valid values are ` + "`" + `auto-key` + "`" + ` (the default), ` + "`" + `manual-key` + "`" + `, or ` + "`" + `global-protect-satellite` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ak_ike_gateway",
					Description: `(Optional) IKE gateway name.`,
				},
				resource.Attribute{
					Name:        "ak_ipsec_crypto_profile",
					Description: `(Optional) IPSec crypto profile name.`,
				},
				resource.Attribute{
					Name:        "mk_local_spi",
					Description: `(Optional) Outbound SPI, hex format.`,
				},
				resource.Attribute{
					Name:        "mk_remote_spi",
					Description: `(Optional) Inbound SPI, hex format.`,
				},
				resource.Attribute{
					Name:        "mk_local_address_ip",
					Description: `(Optional) Specify exact IP address if interface has multiple addresses.`,
				},
				resource.Attribute{
					Name:        "mk_local_address_floating_ip",
					Description: `(Optional) Floating IP address in HA Active-Active configuration.`,
				},
				resource.Attribute{
					Name:        "mk_protocol",
					Description: `(Optional) Manual key protocol. Valid valies are ` + "`" + `esp` + "`" + ` or ` + "`" + `ah` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mk_auth_type",
					Description: `(Optional) Authentication algorithm. Valid values are ` + "`" + `md5` + "`" + `, ` + "`" + `sha1` + "`" + `, ` + "`" + `sha256` + "`" + `, ` + "`" + `sha384` + "`" + `, ` + "`" + `sha512` + "`" + `, or ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mk_auth_key",
					Description: `(Optional) The auth key for the given auth type.`,
				},
				resource.Attribute{
					Name:        "mk_esp_encryption_type",
					Description: `(Optional) The encryption algorithm. Valid values are ` + "`" + `des` + "`" + `, ` + "`" + `3des` + "`" + `, ` + "`" + `aes-128-cbc` + "`" + `, ` + "`" + `aes-192-cbc` + "`" + `, ` + "`" + `aes-256-cbc` + "`" + `, or ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mk_esp_encryption_key",
					Description: `(Optional) The encryption key.`,
				},
				resource.Attribute{
					Name:        "gps_interface",
					Description: `(Optional) Interface to communicate with portal.`,
				},
				resource.Attribute{
					Name:        "gps_portal_address",
					Description: `(Optional) GlobalProtect portal address.`,
				},
				resource.Attribute{
					Name:        "gps_prefer_ipv6",
					Description: `(Optional, PAN-OS 8.0+, bool) Prefer to register the portal in IPv6. Only applicable to FQDN portal-address.`,
				},
				resource.Attribute{
					Name:        "gps_interface_ip_ipv4",
					Description: `(Optional) specify exact IP address if interface has multiple addresses (IPv4).`,
				},
				resource.Attribute{
					Name:        "gps_interface_ip_ipv6",
					Description: `(Optional, PAN-OS 8.0+) specify exact IP address if interface has multiple addresses (IPv6).`,
				},
				resource.Attribute{
					Name:        "gps_interface_floating_ip_ipv4",
					Description: `(Optional, PAN-OS 7.0+) Floating IPv4 address in HA Active-Active configuration.`,
				},
				resource.Attribute{
					Name:        "gps_interface_floating_ip_ipv6",
					Description: `(Optional, PAN-OS 8.0+) Floating IPv6 address in HA Active-Active configuration.`,
				},
				resource.Attribute{
					Name:        "gps_publish_connected_routes",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to to publish connected and static routes.`,
				},
				resource.Attribute{
					Name:        "gps_publish_routes",
					Description: `(Optional, list) Specify list of routes to publish to Global Protect Gateway.`,
				},
				resource.Attribute{
					Name:        "gps_local_certificate",
					Description: `(Optional) GlobalProtect satellite certificate file name.`,
				},
				resource.Attribute{
					Name:        "gps_certificate_profile",
					Description: `(Optional) Profile for authenticating GlobalProtect gateway certificates.`,
				},
				resource.Attribute{
					Name:        "enable_tunnel_monitor",
					Description: `(Optional, bool) Enable tunnel monitoring on this tunnel.`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_destination_ip",
					Description: `(Optional) Destination IP to send ICMP probe.`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_source_ip",
					Description: `(Optional) Source IP to send ICMP probe`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_profile",
					Description: `(Optional) Tunnel monitor profile.`,
				},
				resource.Attribute{
					Name:        "tunnel_monitor_proxy_id",
					Description: `(Optional, PAN-OS 7.0+) Which proxy-id (or proxy-id-v6) the monitoring traffic will use.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_ipsec_tunnel_proxy_id_ipv4",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"ipsec",
				"tunnel",
				"proxy",
				"id",
				"ipv4",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name`,
				},
				resource.Attribute{
					Name:        "ipsec_tunnel",
					Description: `(Required) The auto key IPSec tunnel to attach this proxy ID to.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) IP subnet or IP address represents local network.`,
				},
				resource.Attribute{
					Name:        "remote",
					Description: `(Optional) IP subnet or IP address represents remote network.`,
				},
				resource.Attribute{
					Name:        "protocol_any",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` for any IP protocol.`,
				},
				resource.Attribute{
					Name:        "protocol_number",
					Description: `(Optional, int) IP protocol number.`,
				},
				resource.Attribute{
					Name:        "protocol_tcp_local",
					Description: `(Optional, int) Local TCP port number.`,
				},
				resource.Attribute{
					Name:        "protocol_tcp_remote",
					Description: `(Optional, int) Remote TCP port number.`,
				},
				resource.Attribute{
					Name:        "protocol_udp_local",
					Description: `(Optional, int) Local UDP port number.`,
				},
				resource.Attribute{
					Name:        "protocol_udp_remote",
					Description: `(Optional, int) Remote UDP port number.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_layer2_subinterface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"layer2",
				"subinterface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "interface_type",
					Description: `(Optional) The interface type. Valid values are ` + "`" + `ethernet` + "`" + ` (default) or ` + "`" + `aggregate-ethernet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parent_interface",
					Description: `(Required) The name of the parent interface.`,
				},
				resource.Attribute{
					Name:        "parent_mode",
					Description: `(Optional) The parent's mode. Valid values are ` + "`" + `layer2` + "`" + ` (default) or ` + "`" + `virtual-wire` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Required) The vsys that will use this interface. This should be something like ` + "`" + `vsys1` + "`" + ` or ` + "`" + `vsys3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional, int) The interface's tag.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_layer3_subinterface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"layer3",
				"subinterface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "interface_type",
					Description: `(Optional) The interface type. Valid values are ` + "`" + `ethernet` + "`" + ` (default) or ` + "`" + `aggregate-ethernet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parent_interface",
					Description: `(Required) The name of the parent interface.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Required) The vsys that will use this interface. This should be something like ` + "`" + `vsys1` + "`" + ` or ` + "`" + `vsys3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional, int) The interface's tag.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses.`,
				},
				resource.Attribute{
					Name:        "ipv6_enabled",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable IPv6.`,
				},
				resource.Attribute{
					Name:        "ipv6_interface_id",
					Description: `(Optional) The IPv6 interface ID.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional) The IPv6 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to enable DHCP.`,
				},
				resource.Attribute{
					Name:        "create_dhcp_default_route",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to create a DHCP default route.`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route_metric",
					Description: `(Optional) The metric for the DHCP default route.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "decrypt_forward",
					Description: `(Optional, bool, PAN-OS 8.1+) Set to ` + "`" + `true` + "`" + ` to enable decrypt forward.`,
				},
				resource.Attribute{
					Name:        "dhcp_send_hostname_enable",
					Description: `(Optional, PAN-OS 9.0+) For DHCP layer3 interfaces: enable sending the firewall or a custom hostname to DHCP server`,
				},
				resource.Attribute{
					Name:        "dhcp_send_hostname_value",
					Description: `(Optional, PAN-OS 9.0+) For DHCP layer3 interfaces: the interface hostname. Leaving this unspecified with ` + "`" + `dhcp_send_hostname_enable` + "`" + ` set means to send the system hostname.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_log_forwarding_profile",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"panorama",
				"log",
				"forwarding",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "enhanced_logging",
					Description: `(Optional, bool, PAN-OS 8.1+) Set to ` + "`" + `true` + "`" + ` to enable enhanced logging.`,
				},
				resource.Attribute{
					Name:        "match_list",
					Description: `(Optional, repeatable) The match list spec (defined below). ` + "`" + `match_list` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `(Optional) The log type. Valid values are ` + "`" + `traffic` + "`" + ` (default), ` + "`" + `threat` + "`" + `, ` + "`" + `wildfire` + "`" + `, ` + "`" + `url` + "`" + `, ` + "`" + `data` + "`" + `, ` + "`" + `gtp` + "`" + `, ` + "`" + `tunnel` + "`" + `, ` + "`" + `auth` + "`" + `, ` + "`" + `sctp` + "`" + `, or ` + "`" + `decryption` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) The filter (default: ` + "`" + `All Logs` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "send_to_panorama",
					Description: `(Optional, bool) Set to ` + "`" + `true` + "`" + ` to send to Panorama.`,
				},
				resource.Attribute{
					Name:        "snmptrap_server_profiles",
					Description: `(Optional) List of SNMP server profiles.`,
				},
				resource.Attribute{
					Name:        "email_server_profiles",
					Description: `(Optional) List of email server profiles.`,
				},
				resource.Attribute{
					Name:        "syslog_server_profiles",
					Description: `(Optional) List of syslog server profiles.`,
				},
				resource.Attribute{
					Name:        "http_server_profiles",
					Description: `(Optional) List of http server profiles.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional, repeatable) Match list action spec (defined below). ` + "`" + `match_list.action` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "azure_integration",
					Description: `(Optional) The Azure integration spec (defined below). Mutually exclusive with ` + "`" + `tagging_integration` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tagging_integration",
					Description: `(Optional) The tagging integration spec (defined below). Mutually exclusive with ` + "`" + `azure_integration` + "`" + `. ` + "`" + `match_list.action.azure_integration` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "azure_integration",
					Description: `(Optional, bool) This param defaults to ` + "`" + `true` + "`" + ` and should not be changed. ` + "`" + `match_list.action.tagging_integration` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The action. Valid values are ` + "`" + `add-tag` + "`" + ` (default) or ` + "`" + `remove-tag` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) The target. Valid values are ` + "`" + `source-address` + "`" + ` (default) or ` + "`" + `destination-address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional, int) The timeout.`,
				},
				resource.Attribute{
					Name:        "local_registration",
					Description: `(Optional) The local registration spec (defined below). Only one of ` + "`" + `local_registration` + "`" + `, ` + "`" + `remote_registration` + "`" + `, or ` + "`" + `panorama_registration` + "`" + ` should be defined.`,
				},
				resource.Attribute{
					Name:        "remote_registration",
					Description: `(Optional) The remote registration spec (defined below). Only one of ` + "`" + `local_registration` + "`" + `, ` + "`" + `remote_registration` + "`" + `, or ` + "`" + `panorama_registration` + "`" + ` should be defined.`,
				},
				resource.Attribute{
					Name:        "panorama_registration",
					Description: `(Optional) The panorama registration spec (defined below). Only one of ` + "`" + `local_registration` + "`" + `, ` + "`" + `remote_registration` + "`" + `, or ` + "`" + `panorama_registration` + "`" + ` should be defined. ` + "`" + `match_list.action.tagging_integration.local_registration` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Required) List of administrative tags. ` + "`" + `match_list.action.tagging_integration.remote_registration` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Required) List of administrative tags.`,
				},
				resource.Attribute{
					Name:        "http_profile",
					Description: `(Required) The HTTP profile. ` + "`" + `match_list.action.tagging_integration.panorama_registration` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Required) List of administrative tags.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_loopback_interface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"loopback",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name. This must start with ` + "`" + `loopback.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys that will use this interface (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional, bool) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv6 MSS adjust value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_management_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"management",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The management profile's name.`,
				},
				resource.Attribute{
					Name:        "ping",
					Description: `(Optional) Allow ping.`,
				},
				resource.Attribute{
					Name:        "telnet",
					Description: `(Optional) Allow telnet.`,
				},
				resource.Attribute{
					Name:        "ssh",
					Description: `(Optional) Allow SSH.`,
				},
				resource.Attribute{
					Name:        "http",
					Description: `(Optional) Allow HTTP.`,
				},
				resource.Attribute{
					Name:        "http_ocsp",
					Description: `(Optional) Allow HTTP OCSP.`,
				},
				resource.Attribute{
					Name:        "https",
					Description: `(Optional) Allow HTTPS.`,
				},
				resource.Attribute{
					Name:        "snmp",
					Description: `(Optional) Allow SNMP.`,
				},
				resource.Attribute{
					Name:        "response_pages",
					Description: `(Optional) Allow response pages.`,
				},
				resource.Attribute{
					Name:        "userid_service",
					Description: `(Optional) Allow User ID service.`,
				},
				resource.Attribute{
					Name:        "userid_syslog_listener_ssl",
					Description: `(Optional) Allow User ID syslog listener for SSL.`,
				},
				resource.Attribute{
					Name:        "userid_syslog_listener_udp",
					Description: `(Optional) Allow User ID syslog listener for UDP.`,
				},
				resource.Attribute{
					Name:        "permitted_ips",
					Description: `(Optional) The list of permitted IP addresses or address ranges for this management profile.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_monitor_profile",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"monitor",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The monitor profile name.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional, int) The probing interval in seconds.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Optional, int) The number of failed probes to determine that the tunnel is down.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action triggered when tunnel's status changes. Valid values are ` + "`" + `wait-recover` + "`" + ` (default) or ` + "`" + `fail-over` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_nat_rule",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"panorama",
				"nat",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The NAT rule's name.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group to put the NAT rule into (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rulebase",
					Description: `(Optional) The rulebase. This can be ` + "`" + `pre-rulebase` + "`" + ` (default), ` + "`" + `post-rulebase` + "`" + `, or ` + "`" + `rulebase` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional). NAT type. This can be ` + "`" + `ipv4` + "`" + ` (default), ` + "`" + `nat64` + "`" + `, or ` + "`" + `nptv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `(Required) The list of source zone(s).`,
				},
				resource.Attribute{
					Name:        "destination_zone",
					Description: `(Required) The destination zone.`,
				},
				resource.Attribute{
					Name:        "to_interface",
					Description: `(Optional) Egress interface from route lookup (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Service (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Required) List of source address(es).`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Required) List of destination address(es).`,
				},
				resource.Attribute{
					Name:        "sat_type",
					Description: `(Optional) Type of source address translation. This can be ` + "`" + `none` + "`" + ` (default), ` + "`" + `dynamic-ip-and-port` + "`" + `, ` + "`" + `dynamic-ip` + "`" + `, or ` + "`" + `static-ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_address_type",
					Description: `(Optional) Source address translation address type. This can be ` + "`" + `interface-address` + "`" + ` or ` + "`" + `translated-address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_translated_addresses",
					Description: `(Optional) Source address translation list of translated addresses.`,
				},
				resource.Attribute{
					Name:        "sat_interface",
					Description: `(Optional) Source address translation interface.`,
				},
				resource.Attribute{
					Name:        "sat_ip_address",
					Description: `(Optional) Source address translation IP address.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_type",
					Description: `(Optional) Source address translation fallback type. This can be ` + "`" + `none` + "`" + `, ` + "`" + `interface-address` + "`" + `, or ` + "`" + `translated-address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_translated_addresses",
					Description: `(Optional) Source address translation list of fallback translated addresses.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_interface",
					Description: `(Optional) Source address translation fallback interface.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_ip_type",
					Description: `(Optional) Source address translation fallback IP type. This can be ` + "`" + `ip` + "`" + ` or ` + "`" + `floating` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sat_fallback_ip_address",
					Description: `(Optional) The source address translation fallback IP address.`,
				},
				resource.Attribute{
					Name:        "sat_static_translated_address",
					Description: `(Optional) The statically translated source address.`,
				},
				resource.Attribute{
					Name:        "sat_static_bi_directional",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable bi-directional source address translation.`,
				},
				resource.Attribute{
					Name:        "dat_type",
					Description: `(Optional) Destination address translation type. This should be either ` + "`" + `static` + "`" + ` or ` + "`" + `dynamic` + "`" + `. The ` + "`" + `dynamic` + "`" + ` option is only available on PAN-OS 8.1+.`,
				},
				resource.Attribute{
					Name:        "dat_address",
					Description: `(Optional) Destination address translation's address. Requires ` + "`" + `dat_type` + "`" + ` be set to "static" or "dynamic".`,
				},
				resource.Attribute{
					Name:        "dat_port",
					Description: `(Optional) Destination address translation's port number. Requires ` + "`" + `dat_type` + "`" + ` be set to "static" or "dynamic".`,
				},
				resource.Attribute{
					Name:        "dat_dynamic_distribution",
					Description: `(Optional, PAN-OS 8.1+) Distribution algorithm for destination address pool. The PAN-OS 8.1 GUI doesn't seem to set this anywhere, but this is added here for completeness' sake. Requires ` + "`" + `dat_type` + "`" + ` of "dynamic".`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to disable this rule.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) A target definition (see below). If there are no target sections, then the rule will apply to every vsys of every device in the device group.`,
				},
				resource.Attribute{
					Name:        "negate_target",
					Description: `(Optional, bool) Instead of applying the rule for the given serial numbers, apply it to everything except them. The following arguments are valid for each ` + "`" + `target` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `(Optional) A subset of all available vsys on the firewall that should be in this device group. If the firewall is a virtual firewall, then this parameter should just be omitted.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_nat_rule_group",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"panorama",
				"nat",
				"rule",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_pbf_rule_group",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"panorama",
				"pbf",
				"rule",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_redistribution_profile_ipv4",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"redistribution",
				"profile",
				"ipv4",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The redistribution profile's name.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add the redistribution profile to.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required, int) The priority, integer from 1 to 255.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The action. Valid values are ` + "`" + `redist` + "`" + ` (default) or ` + "`" + `no-redist` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `(Optional) The source types. Valid values are ` + "`" + `bgp` + "`" + `, ` + "`" + `connect` + "`" + `, ` + "`" + `ospf` + "`" + `, ` + "`" + `rip` + "`" + `, and ` + "`" + `static` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Specify candidate routes.`,
				},
				resource.Attribute{
					Name:        "destinations",
					Description: `(Optional) Specify candidate routes' next-hop addresses (subnet match).`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) Specify candidate routes' next-hop addresses (subnet match).`,
				},
				resource.Attribute{
					Name:        "ospf_path_types",
					Description: `(Optional) OSPF path types. Valid values are ` + "`" + `intra-area` + "`" + `, ` + "`" + `inter-area` + "`" + `, ` + "`" + `ext-1` + "`" + `, and ` + "`" + `ext-2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ospf_areas",
					Description: `(Optional) OSPF areas.`,
				},
				resource.Attribute{
					Name:        "ospf_tags",
					Description: `(Optional) OSPF tags.`,
				},
				resource.Attribute{
					Name:        "bgp_communities",
					Description: `(Optional) BGP communities.`,
				},
				resource.Attribute{
					Name:        "bgp_extended_communities",
					Description: `(Optional) BGP extended communities.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_security_policy",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"panorama",
				"security",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_security_rule_group",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"panorama",
				"security",
				"rule",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_service_group",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"panorama",
				"service",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The service group's name.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group to put the service group into (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) List of services to put in this service group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_service_object",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"panorama",
				"service",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The service object's name.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group to put the service object into (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The service object's description.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The service's protocol. This should be ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, or ` + "`" + `sctp` + "`" + ` (PAN-OS 8.1+).`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) The source port. This can be a single port number, range (1-65535), or comma separated (80,8080,443).`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Required) The destination port. This can be a single port number, range (1-65535), or comma separated (80,8080,443).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
				resource.Attribute{
					Name:        "override_session_timeout",
					Description: `(Optional, bool, PAN-OS 8.1+) Set to ` + "`" + `true` + "`" + ` to override the default application timeouts.`,
				},
				resource.Attribute{
					Name:        "override_timeout",
					Description: `(Optional, int, PAN-OS 8.1+) The overridden TCP timeout.`,
				},
				resource.Attribute{
					Name:        "override_half_closed_timeout",
					Description: `(Optional, int, PAN-OS 8.1+) The overridden TCP half closed timeout.`,
				},
				resource.Attribute{
					Name:        "override_time_wait_timeout",
					Description: `(Optional, int, PAN-OS 8.1+) The overridden TCP wait time.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_snmptrap_server_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"panorama",
				"snmptrap",
				"server",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) The template location. Mutually exclusive with ` + "`" + `template_stack` + "`" + ` and ` + "`" + `device_group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `(Optional) The template stack location. Mutually exclusive with ` + "`" + `template` + "`" + ` and ` + "`" + `device_group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location. Mutually exclusive with ` + "`" + `template` + "`" + ` and ` + "`" + `template_stack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys. This will likely be ` + "`" + `shared` + "`" + `, and it should be defined if you specified either ` + "`" + `template` + "`" + ` or ` + "`" + `template_stack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group's name.`,
				},
				resource.Attribute{
					Name:        "v2c_server",
					Description: `(Optional, repeatable) A v2c server (defined below).`,
				},
				resource.Attribute{
					Name:        "v3_server",
					Description: `(Optional, repeatable) A v3 server (defined below). ` + "`" + `v2c_server` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The server name.`,
				},
				resource.Attribute{
					Name:        "manager",
					Description: `(Required) The hostname.`,
				},
				resource.Attribute{
					Name:        "community",
					Description: `(Required) The SNMP community. ` + "`" + `v3_server` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The server name.`,
				},
				resource.Attribute{
					Name:        "manager",
					Description: `(Required) The hostname.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) Username.`,
				},
				resource.Attribute{
					Name:        "engine_id",
					Description: `(Optional) The engine ID.`,
				},
				resource.Attribute{
					Name:        "auth_password",
					Description: `(Required) SNMP auth password.`,
				},
				resource.Attribute{
					Name:        "priv_password",
					Description: `(Required) SNMP priv password.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_static_route_ipv4",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"static",
				"route",
				"ipv4",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The static route's name.`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add the static route to.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Destination IP address / prefix.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) Interface to use.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The next hop type. Valid values are ` + "`" + `ip-address` + "`" + ` (the default), ` + "`" + `discard` + "`" + `, ` + "`" + `next-vr` + "`" + `, or an empty string for ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) The value for the ` + "`" + `type` + "`" + ` setting.`,
				},
				resource.Attribute{
					Name:        "admin_distance",
					Description: `(Optional) The admin distance.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional, int) Metric value / path cost (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional) Target routing table to install the route. Valid values are ` + "`" + `unicast` + "`" + ` (the default), ` + "`" + `no install` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `(Optional, PAN-OS 7.1+) BFD configuration.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_syslog_server_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"panorama",
				"syslog",
				"server",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) The template location. Mutually exclusive with ` + "`" + `template_stack` + "`" + ` and ` + "`" + `device_group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `(Optional) The template stack location. Mutually exclusive with ` + "`" + `template` + "`" + ` and ` + "`" + `device_group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location. Mutually exclusive with ` + "`" + `template` + "`" + ` and ` + "`" + `template_stack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys. This will likely be ` + "`" + `shared` + "`" + `, and it should be defined if you specified either ` + "`" + `template` + "`" + ` or ` + "`" + `template_stack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group's name.`,
				},
				resource.Attribute{
					Name:        "config_format",
					Description: `(Optional) Config format.`,
				},
				resource.Attribute{
					Name:        "system_format",
					Description: `(Optional) System format.`,
				},
				resource.Attribute{
					Name:        "threat_format",
					Description: `(Optional) Threat format.`,
				},
				resource.Attribute{
					Name:        "traffic_format",
					Description: `(Optional) Traffic format.`,
				},
				resource.Attribute{
					Name:        "hip_match_format",
					Description: `(Optional) HIP match format.`,
				},
				resource.Attribute{
					Name:        "url_format",
					Description: `(Optional) URL format.`,
				},
				resource.Attribute{
					Name:        "data_format",
					Description: `(Optional) Data format.`,
				},
				resource.Attribute{
					Name:        "wildfire_format",
					Description: `(Optional) Wildfire format.`,
				},
				resource.Attribute{
					Name:        "tunnel_format",
					Description: `(Optional) Tunnel format.`,
				},
				resource.Attribute{
					Name:        "user_id_format",
					Description: `(Optional) UserID format.`,
				},
				resource.Attribute{
					Name:        "gtp_format",
					Description: `(Optional) GTP format.`,
				},
				resource.Attribute{
					Name:        "auth_format",
					Description: `(Optional) Auth format.`,
				},
				resource.Attribute{
					Name:        "sctp_format",
					Description: `(Optional) SCTP format.`,
				},
				resource.Attribute{
					Name:        "iptag_format",
					Description: `(Optional) IP tag format.`,
				},
				resource.Attribute{
					Name:        "escaped_characters",
					Description: `(Optional) The escaped characters (as a string).`,
				},
				resource.Attribute{
					Name:        "escape_character",
					Description: `(Optional) The escape character.`,
				},
				resource.Attribute{
					Name:        "syslog_server",
					Description: `(Required, repeatable) The server spec (defined below). ` + "`" + `syslog_server` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Server name.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) The hostname.`,
				},
				resource.Attribute{
					Name:        "transport",
					Description: `(Optional) The transport. Valid values are ` + "`" + `UDP` + "`" + ` (default), ` + "`" + `TCP` + "`" + `, or ` + "`" + `SSL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, int) The port number (default: 514).`,
				},
				resource.Attribute{
					Name:        "syslog_format",
					Description: `(Optional) The syslog format. Valid values are ` + "`" + `BSD` + "`" + ` (default) or ` + "`" + `IETF` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(Optional) The syslog facility. Valid values are ` + "`" + `LOG_USER` + "`" + ` (default), ` + "`" + `LOG_LOCAL0` + "`" + `, ` + "`" + `LOG_LOCAL1` + "`" + `, ` + "`" + `LOG_LOCAL2` + "`" + `, ` + "`" + `LOG_LOCAL3` + "`" + `, ` + "`" + `LOG_LOCAL4` + "`" + `, ` + "`" + `LOG_LOCAL5` + "`" + `, ` + "`" + `LOG_LOCAL6` + "`" + `, or ` + "`" + `LOG_LOCAL7` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_template",
			Category:         "Panorama",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The template's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The template's description.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `The device definition (see below). The following arguments are valid for each ` + "`" + `device` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `(Optional) A subset of all available vsys on the firewall that should be in this template. If the firewall is a virtual firewall, then this parameter should just be omitted.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_template_entry",
			Category:         "Panorama",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"template",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `(Optional) A subset of all available vsys on the firewall that should be in this template. If the firewall is a virtual firewall, then this parameter should just be omitted.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_template_stack",
			Category:         "Panorama",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"template",
				"stack",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The stack's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The stack's description.`,
				},
				resource.Attribute{
					Name:        "default_vsys",
					Description: `(Optional) The default virtual system template configuration pushed to firewalls with a single virtual system.`,
				},
				resource.Attribute{
					Name:        "templates",
					Description: `(Optional) List of templates in this stack.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `(Optional) List of serial numbers to include in this stack.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_template_stack_entry",
			Category:         "Panorama",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"template",
				"stack",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template_stack",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Required) The serial number of the device to add.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_template_variable",
			Category:         "Panorama",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"template",
				"variable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The template's name. This must start with a dollar sign ($).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The variable type. Valid values are ` + "`" + `ip-netmask` + "`" + ` (default), ` + "`" + `ip-range` + "`" + `, ` + "`" + `fqdn` + "`" + `, ` + "`" + `group-id` + "`" + `, or ` + "`" + `interface` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The variable value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_tunnel_interface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"tunnel",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name. This must start with ` + "`" + `tunnel.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys that will use this interface (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_virtual_router",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"virtual",
				"router",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_virtual_router_entry",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"virtual",
				"router",
				"entry",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_vlan",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"vlan",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the object into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "vlan_interface",
					Description: `(Optional) The VLAN interface.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional, computed) List of layer2 interfaces. You can also leave this blank and also use [panos_panorama_vlan_entry](./vlan_entry.html) for more control.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_vlan_entry",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"vlan",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vlan",
					Description: `(Required) The VLAN's name.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface's name.`,
				},
				resource.Attribute{
					Name:        "mac_addresses",
					Description: `(Optional) List of MAC addresses that should go with this entry.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_vlan_interface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"vlan",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name. Must start with ` + "`" + `vlan.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The template name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys that will use this interface (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable DHCP on this interface.`,
				},
				resource.Attribute{
					Name:        "create_dhcp_default_route",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to create a DHCP default route.`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route_metric",
					Description: `(Optional) The metric for the DHCP default route.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv6 MSS adjust value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_zone",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"zone",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_zone_entry",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"panorama",
				"zone",
				"entry",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_pbf_rule_group",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"pbf",
				"rule",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_group",
					Description: `The device group (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rulebase",
					Description: `The rulebase. This can be ` + "`" + `pre-rulebase` + "`" + ` (default), ` + "`" + `post-rulebase` + "`" + `, or ` + "`" + `rulebase` + "`" + `. NGFW specific arguments:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The rule definition (see below). The rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The rule name.`,
				},
				resource.Attribute{
					Name:        "audit_comment",
					Description: `When this rule is created/updated, the audit comment to apply for this rule.`,
				},
				resource.Attribute{
					Name:        "group_tag",
					Description: `(PAN-OS 9.0+) The group tag.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The rule description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List of tags for this rule.`,
				},
				resource.Attribute{
					Name:        "active_active_device_binding",
					Description: `The active-active device binding.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `The schedule.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to disable this rule.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The source spec (defined below).`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) The destination spec (defined below).`,
				},
				resource.Attribute{
					Name:        "forwarding",
					Description: `(Required) The forwarding spec (defined below).`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `A target definition (see below). If there are no target sections, then the rule will apply to every vsys of every device in the device group.`,
				},
				resource.Attribute{
					Name:        "negate_target",
					Description: `(bool) Instead of applying the rule for the given serial numbers, apply it to everything except them. ` + "`" + `rule.source` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `If you want a source type of "zone", then define this list with the desired source zones. Mutually exclusive with ` + "`" + `rule.interfaces` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `If you want a source type of "interface", then define this list with the desired source interfaces. Mutually exclusive with ` + "`" + `rule.zones` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `(Required) List of source IP addresses.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Required) List of source users.`,
				},
				resource.Attribute{
					Name:        "negate",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to negate the source. ` + "`" + `rule.destination` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `(Required) The list of destination addresses.`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Required) The list of applications.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) The list of services.`,
				},
				resource.Attribute{
					Name:        "negate",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to negate the destination. ` + "`" + `rule.forwarding` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action to take. Valid values are ` + "`" + `forward` + "`" + ` (default), ` + "`" + `forward-to-vsys` + "`" + `, ` + "`" + `discard` + "`" + `, or ` + "`" + `no-pbf` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `If ` + "`" + `action=forward-to-vsys` + "`" + `, the vsys to forward to.`,
				},
				resource.Attribute{
					Name:        "egress_interface",
					Description: `If ` + "`" + `action=forward` + "`" + `, the egress interface.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `If ` + "`" + `action=forward` + "`" + `, the next hop type. Valid values are ` + "`" + `ip-address` + "`" + `, ` + "`" + `fqdn` + "`" + `, or leaving this empty for a next hop type of None.`,
				},
				resource.Attribute{
					Name:        "next_hop_value",
					Description: `If ` + "`" + `action=forward` + "`" + ` and ` + "`" + `next_hop_type` + "`" + ` is defined, then the next hop address.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `The monitor spec (defined below). If you do not want to enable monitoring, then do not specify a ` + "`" + `monitor` + "`" + ` config block.`,
				},
				resource.Attribute{
					Name:        "symmetric_return",
					Description: `The symmetric return spec (defined below). If you do not want to enforce symmetric ` + "`" + `rule.forwarding.monitor` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `The montior profile to use.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The monitor IP address.`,
				},
				resource.Attribute{
					Name:        "disable_if_unreachable",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to disable this rule if nexthop/monitor IP is unreachable. ` + "`" + `rule.forwarding.symmetric_return` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to enforce symmetric return.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `List of next hop addresses. ` + "`" + `rule.target` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `A subset of all available vsys on the firewall that should be in this device group. If the firewall is a virtual firewall, then this parameter should just be omitted. ## Attribute Reference Each ` + "`" + `rule` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(PAN-OS 9.0+) The PAN-OS UUID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `(PAN-OS 9.0+) The PAN-OS UUID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_radius_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"radius",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. NGFW / Panorama:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "admin_use_only",
					Description: `(bool) Administrator use only.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(int) Timeout in seconds (default: ` + "`" + `3` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(int) Number of attempts before giving up authentication (default: ` + "`" + `3` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, PAN-OS 8.0+) Authentication protocol settings spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `List of server specs, as defined below. ` + "`" + `protocol` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "chap",
					Description: `(bool) CHAP.`,
				},
				resource.Attribute{
					Name:        "pap",
					Description: `(bool) PAP.`,
				},
				resource.Attribute{
					Name:        "auth",
					Description: `(bool, PAN-OS 8.0 only) Auto.`,
				},
				resource.Attribute{
					Name:        "peap_mschap_v2",
					Description: `PEAP-MSCHAPv2 spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "peap_with_gtc",
					Description: `PEAP with GTC spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "eap_ttls_with_pap",
					Description: `EAP-TTLS with PAP spec, as defined below. ` + "`" + `protocol.peap_mschap_v2` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "make_outer_identity_anonymous",
					Description: `(bool) Make outer identity anonymous.`,
				},
				resource.Attribute{
					Name:        "allow_expired_password_change",
					Description: `(bool) Allow users to change passwords after expiry.`,
				},
				resource.Attribute{
					Name:        "certificate_profile",
					Description: `(Required) Certificate profile for verifying the Radius server. ` + "`" + `protocol.peap_with_gtc` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "make_outer_identity_anonymous",
					Description: `(bool) Make outer identity anonymous.`,
				},
				resource.Attribute{
					Name:        "certificate_profile",
					Description: `(Required) Certificate profile for verifying the Radius server. ` + "`" + `protocol.eap_ttls_with_pap` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "make_outer_identity_anonymous",
					Description: `(bool) Make outer identity anonymous.`,
				},
				resource.Attribute{
					Name:        "certificate_profile",
					Description: `(Required) Certificate profile for verifying the Radius server. ` + "`" + `server` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) Server hostname or IP address.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Required) Shared secret for Radius communication.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(int) Radius server port number (default: ` + "`" + `1812` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_redistribution_profile_ipv4",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"redistribution",
				"profile",
				"ipv4",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The redistribution profile's name.`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add the redistribution profile to.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required, int) The priority, integer from 1 to 255.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The action. Valid values are ` + "`" + `redist` + "`" + ` (default) or ` + "`" + `no-redist` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `(Optional) The source types. Valid values are ` + "`" + `bgp` + "`" + `, ` + "`" + `connect` + "`" + `, ` + "`" + `ospf` + "`" + `, ` + "`" + `rip` + "`" + `, and ` + "`" + `static` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional) Specify candidate routes.`,
				},
				resource.Attribute{
					Name:        "destinations",
					Description: `(Optional) Specify candidate routes' next-hop addresses (subnet match).`,
				},
				resource.Attribute{
					Name:        "next_hops",
					Description: `(Optional) Specify candidate routes' next-hop addresses (subnet match).`,
				},
				resource.Attribute{
					Name:        "ospf_path_types",
					Description: `(Optional) OSPF path types. Valid values are ` + "`" + `intra-area` + "`" + `, ` + "`" + `inter-area` + "`" + `, ` + "`" + `ext-1` + "`" + `, and ` + "`" + `ext-2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ospf_areas",
					Description: `(Optional) OSPF areas.`,
				},
				resource.Attribute{
					Name:        "ospf_tags",
					Description: `(Optional) OSPF tags.`,
				},
				resource.Attribute{
					Name:        "bgp_communities",
					Description: `(Optional) BGP communities.`,
				},
				resource.Attribute{
					Name:        "bgp_extended_communities",
					Description: `(Optional) BGP extended communities.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_saml_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"saml",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. NGFW / Panorama:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "admin_use_only",
					Description: `(bool) Administrator use only.`,
				},
				resource.Attribute{
					Name:        "identity_provider_id",
					Description: `(Required) Unique identifier for SAML IdP.`,
				},
				resource.Attribute{
					Name:        "identity_provider_certificate",
					Description: `(Required) Object name of IdP signing certificate.`,
				},
				resource.Attribute{
					Name:        "sso_url",
					Description: `(Required) The single sign on service URL for the IdP server.`,
				},
				resource.Attribute{
					Name:        "sso_binding",
					Description: `SAML HTTP binding for SSO requests to IdP. Valid values are ` + "`" + `"post"` + "`" + ` (default) or ` + "`" + `"redirect"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "slo_url",
					Description: `The single logout service URL for the IdP server.`,
				},
				resource.Attribute{
					Name:        "slo_binding",
					Description: `SAML HTTP binding for SLO requests to IdP. Valid values are ` + "`" + `"post"` + "`" + ` (default) or ` + "`" + `"redirect"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "validate_identity_provider_certificate",
					Description: `(bool) Validate identity provider certificate (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "sign_saml_message",
					Description: `(bool) Sign SAML message to IdP (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "max_clock_skew",
					Description: `(int) Maximum allowed clock skew in seconds between SAML entities (default: ` + "`" + `60` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_security_policy",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"security",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_group",
					Description: `The device group (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rulebase",
					Description: `The rulebase. This can be either ` + "`" + `pre-rulebase` + "`" + ` (default), ` + "`" + `rulebase` + "`" + `, or ` + "`" + `post-rulebase` + "`" + `. NGFW specific arguments:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `A security rule definition (see below). The security rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The security rule name.`,
				},
				resource.Attribute{
					Name:        "audit_comment",
					Description: `When this rule is created/updated, the audit comment to apply for this rule.`,
				},
				resource.Attribute{
					Name:        "group_tag",
					Description: `(PAN-OS 9.0+) The group tag.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Rule type. This can be ` + "`" + `universal` + "`" + ` (default), ` + "`" + `interzone` + "`" + `, or ` + "`" + `intrazone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List of tags for this security rule.`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `(Required) List of source zones.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Required) List of source addresses.`,
				},
				resource.Attribute{
					Name:        "negate_source",
					Description: `(bool) If the source should be negated.`,
				},
				resource.Attribute{
					Name:        "source_users",
					Description: `(Required) List of source users.`,
				},
				resource.Attribute{
					Name:        "hip_profiles",
					Description: `List of HIP profiles.`,
				},
				resource.Attribute{
					Name:        "destination_zones",
					Description: `(Required) List of destination zones.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Required) List of destination addresses.`,
				},
				resource.Attribute{
					Name:        "negate_destination",
					Description: `(bool) If the destination should be negated.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(Required) List of applications.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) List of services.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `(Required) List of categories.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action for the matched traffic. This can be ` + "`" + `allow` + "`" + ` (default), ` + "`" + `deny` + "`" + `, ` + "`" + `drop` + "`" + `, ` + "`" + `reset-client` + "`" + `, ` + "`" + `reset-server` + "`" + `, or ` + "`" + `reset-both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_setting",
					Description: `Log forwarding profile.`,
				},
				resource.Attribute{
					Name:        "log_start",
					Description: `(bool) Log the start of the traffic flow.`,
				},
				resource.Attribute{
					Name:        "log_end",
					Description: `(bool) Log the end of the traffic flow (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to disable this rule.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `The security rule schedule.`,
				},
				resource.Attribute{
					Name:        "icmp_unreachable",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to enable ICMP unreachable.`,
				},
				resource.Attribute{
					Name:        "disable_server_response_inspection",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to disable server response inspection.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `Profile Setting: ` + "`" + `Group` + "`" + ` - The group profile name.`,
				},
				resource.Attribute{
					Name:        "virus",
					Description: `Profile Setting: ` + "`" + `Profiles` + "`" + ` - The antivirus setting.`,
				},
				resource.Attribute{
					Name:        "spyware",
					Description: `Profile Setting: ` + "`" + `Profiles` + "`" + ` - The anti-spyware setting.`,
				},
				resource.Attribute{
					Name:        "vulnerability",
					Description: `Profile Setting: ` + "`" + `Profiles` + "`" + ` - The Vulnerability Protection setting.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `Profile Setting: ` + "`" + `Profiles` + "`" + ` - The URL filtering setting.`,
				},
				resource.Attribute{
					Name:        "file_blocking",
					Description: `Profile Setting: ` + "`" + `Profiles` + "`" + ` - The file blocking setting.`,
				},
				resource.Attribute{
					Name:        "wildfire_analysis",
					Description: `Profile Setting: ` + "`" + `Profiles` + "`" + ` - The WildFire Analysis setting.`,
				},
				resource.Attribute{
					Name:        "data_filtering",
					Description: `Profile Setting: ` + "`" + `Profiles` + "`" + ` - The Data Filtering setting.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(repeatable, Panorama only) A target definition (see below). If there are no target sections, then the rule will apply to every vsys of every device in the device group.`,
				},
				resource.Attribute{
					Name:        "negate_target",
					Description: `(bool, Panorama only) Instead of applying the rule for the given serial numbers, apply it to everything except them. ` + "`" + `rule.target` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `A listing of vsys to apply this rule to. If ` + "`" + `serial` + "`" + ` is a VM, then this parameter should just be omitted. ## Attribute Reference Each ` + "`" + `rule` + "`" + ` section has the following attributes:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The PAN-OS UUID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `The PAN-OS UUID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_security_profile_group",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"security",
				"profile",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_group",
					Description: `The device group (default: ` + "`" + `shared` + "`" + `) NGFW:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "antivirus_profile",
					Description: `The AV profile name.`,
				},
				resource.Attribute{
					Name:        "anti_spyware_profile",
					Description: `Anti Spyware profile name.`,
				},
				resource.Attribute{
					Name:        "vulnerability_profile",
					Description: `Vulnerability profile name.`,
				},
				resource.Attribute{
					Name:        "url_filtering_profile",
					Description: `URL filtering profile name.`,
				},
				resource.Attribute{
					Name:        "file_blocking_profile",
					Description: `File blocking profile name.`,
				},
				resource.Attribute{
					Name:        "data_filtering_profile",
					Description: `Data filtering profile name.`,
				},
				resource.Attribute{
					Name:        "wildfire_analysis_profile",
					Description: `Wildfire analysis profile name.`,
				},
				resource.Attribute{
					Name:        "gtp_profile",
					Description: `GTP profile name.`,
				},
				resource.Attribute{
					Name:        "sctp_profile",
					Description: `SCTP profile name.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_security_rule_group",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"security",
				"rule",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rulebase",
					Description: `(Optional) The rulebase. This can be ` + "`" + `pre-rulebase` + "`" + ` (default), ` + "`" + `post-rulebase` + "`" + `, or ` + "`" + `rulebase` + "`" + `. NGFW specific arguments:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "position_keyword",
					Description: `(Optional) A positioning keyword for this group. This can be ` + "`" + `before` + "`" + `, ` + "`" + `directly before` + "`" + `, ` + "`" + `after` + "`" + `, ` + "`" + `directly after` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `bottom` + "`" + `, or left empty (the default) to have no particular placement. This param works in combination with the ` + "`" + `position_reference` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "position_reference",
					Description: `(Optional) Required if ` + "`" + `position_keyword` + "`" + ` is one of the "above" or "below" variants, this is the name of a non-group rule to use as a reference to place this group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The security rule definition (see below). The security rule ordering will match how they appear in the terraform plan file. The following arguments are valid for each ` + "`" + `rule` + "`" + ` section:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The security rule name.`,
				},
				resource.Attribute{
					Name:        "audit_comment",
					Description: `When this rule is created/updated, the audit comment to apply for this rule.`,
				},
				resource.Attribute{
					Name:        "group_tag",
					Description: `(PAN-OS 9.0+) The group tag.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Rule type. This can be ` + "`" + `universal` + "`" + ` (default), ` + "`" + `interzone` + "`" + `, or ` + "`" + `intrazone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List of tags for this security rule.`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `(Required) List of source zones.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `(Required) List of source addresses.`,
				},
				resource.Attribute{
					Name:        "negate_source",
					Description: `(bool) If the source should be negated.`,
				},
				resource.Attribute{
					Name:        "source_users",
					Description: `(Required) List of source users.`,
				},
				resource.Attribute{
					Name:        "hip_profiles",
					Description: `List of HIP profiles.`,
				},
				resource.Attribute{
					Name:        "destination_zones",
					Description: `(Required) List of destination zones.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `(Required) List of destination addresses.`,
				},
				resource.Attribute{
					Name:        "negate_destination",
					Description: `(bool) If the destination should be negated.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(Required) List of applications.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) List of services.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `(Required) List of categories.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action for the matched traffic. This can be ` + "`" + `allow` + "`" + ` (default), ` + "`" + `deny` + "`" + `, ` + "`" + `drop` + "`" + `, ` + "`" + `reset-client` + "`" + `, ` + "`" + `reset-server` + "`" + `, or ` + "`" + `reset-both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_setting",
					Description: `Log forwarding profile.`,
				},
				resource.Attribute{
					Name:        "log_start",
					Description: `(bool) Log the start of the traffic flow.`,
				},
				resource.Attribute{
					Name:        "log_end",
					Description: `(bool) Log the end of the traffic flow (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to disable this rule.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `The security rule schedule.`,
				},
				resource.Attribute{
					Name:        "icmp_unreachable",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to enable ICMP unreachable.`,
				},
				resource.Attribute{
					Name:        "disable_server_response_inspection",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to disable server response inspection.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `Profile Setting: ` + "`" + `Group` + "`" + ` - The group profile name.`,
				},
				resource.Attribute{
					Name:        "virus",
					Description: `Profile Setting: ` + "`" + `Profiles` + "`" + ` - The antivirus setting.`,
				},
				resource.Attribute{
					Name:        "spyware",
					Description: `Profile Setting: ` + "`" + `Profiles` + "`" + ` - The anti-spyware setting.`,
				},
				resource.Attribute{
					Name:        "vulnerability",
					Description: `Profile Setting: ` + "`" + `Profiles` + "`" + ` - The Vulnerability Protection setting.`,
				},
				resource.Attribute{
					Name:        "url_filtering",
					Description: `Profile Setting: ` + "`" + `Profiles` + "`" + ` - The URL filtering setting.`,
				},
				resource.Attribute{
					Name:        "file_blocking",
					Description: `Profile Setting: ` + "`" + `Profiles` + "`" + ` - The file blocking setting.`,
				},
				resource.Attribute{
					Name:        "wildfire_analysis",
					Description: `Profile Setting: ` + "`" + `Profiles` + "`" + ` - The WildFire Analysis setting.`,
				},
				resource.Attribute{
					Name:        "data_filtering",
					Description: `Profile Setting: ` + "`" + `Profiles` + "`" + ` - The Data Filtering setting.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(repeatable, Panorama only) A target definition (see below). If there are no target sections, then the rule will apply to every vsys of every device in the device group.`,
				},
				resource.Attribute{
					Name:        "negate_target",
					Description: `(bool, Panorama only) Instead of applying the rule for the given serial numbers, apply it to everything except them. ` + "`" + `rule.target` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `A listing of vsys to apply this rule to. If ` + "`" + `serial` + "`" + ` is a VM, then this parameter should just be omitted. ## Attribute Reference Each ` + "`" + `rule` + "`" + ` has the following attribute:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The PAN-OS UUID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `The PAN-OS UUID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_service_group",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"service",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The service group's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the service group into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) List of services to put in this service group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_service_object",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"service",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The service object's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the service object into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The service object's description.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The service's protocol. This should be ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, or ` + "`" + `sctp` + "`" + ` (PAN-OS 8.1+).`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) The source port. This can be a single port number, range (1-65535), or comma separated (80,8080,443).`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Required) The destination port. This can be a single port number, range (1-65535), or comma separated (80,8080,443).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of administrative tags.`,
				},
				resource.Attribute{
					Name:        "override_session_timeout",
					Description: `(Optional, bool, PAN-OS 8.1+) Set to ` + "`" + `true` + "`" + ` to override the default application timeouts.`,
				},
				resource.Attribute{
					Name:        "override_timeout",
					Description: `(Optional, int, PAN-OS 8.1+) The overridden TCP timeout.`,
				},
				resource.Attribute{
					Name:        "override_half_closed_timeout",
					Description: `(Optional, int, PAN-OS 8.1+) The overridden TCP half closed timeout.`,
				},
				resource.Attribute{
					Name:        "override_time_wait_timeout",
					Description: `(Optional, int, PAN-OS 8.1+) The overridden TCP wait time.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_snmptrap_server_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"snmptrap",
				"server",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group's name.`,
				},
				resource.Attribute{
					Name:        "v2c_server",
					Description: `(Optional, repeatable) A v2c server (defined below).`,
				},
				resource.Attribute{
					Name:        "v3_server",
					Description: `(Optional, repeatable) A v3 server (defined below). ` + "`" + `v2c_server` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The server name.`,
				},
				resource.Attribute{
					Name:        "manager",
					Description: `(Required) The hostname.`,
				},
				resource.Attribute{
					Name:        "community",
					Description: `(Required) The SNMP community. ` + "`" + `v3_server` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The server name.`,
				},
				resource.Attribute{
					Name:        "manager",
					Description: `(Required) The hostname.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) Username.`,
				},
				resource.Attribute{
					Name:        "engine_id",
					Description: `(Optional) The engine ID.`,
				},
				resource.Attribute{
					Name:        "auth_password",
					Description: `(Required) SNMP auth password.`,
				},
				resource.Attribute{
					Name:        "priv_password",
					Description: `(Required) SNMP priv password.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ssl_decrypt",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"ssl",
				"decrypt",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. NGFW / Panorama:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "forward_trust_certificate_rsa",
					Description: `Forward trust RSA certificate.`,
				},
				resource.Attribute{
					Name:        "forward_trust_certificate_ecdsa",
					Description: `Forward trust ECDSA certificate.`,
				},
				resource.Attribute{
					Name:        "forward_untrust_certificate_rsa",
					Description: `Forward untrust RSA certificate.`,
				},
				resource.Attribute{
					Name:        "forward_untrust_certificate_ecdsa",
					Description: `Forward untrust ECDSA certificate.`,
				},
				resource.Attribute{
					Name:        "root_ca_excludes",
					Description: `List of root CA excludes.`,
				},
				resource.Attribute{
					Name:        "trusted_root_cas",
					Description: `List of trusted root CAs.`,
				},
				resource.Attribute{
					Name:        "disabled_predefined_exclude_certificates",
					Description: `List of disabled predefined exclude certificates.`,
				},
				resource.Attribute{
					Name:        "ssl_decrypt_exclude_certificate",
					Description: `(repeatable) List of SSL decrypt exclude certificates specs (specified below). ` + "`" + `ssl_decrypt_exclude_certificate` + "`" + ` sections support the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `(bool) Exclude or not.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ssl_decrypt_exclude_certificate_entry",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"ssl",
				"decrypt",
				"exclude",
				"certificate",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. NGFW / Panorama:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "exclude",
					Description: `(bool) Exclude or not (default: ` + "`" + `true` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ssl_decrypt_trusted_root_ca_entry",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"ssl",
				"decrypt",
				"trusted",
				"root",
				"ca",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. NGFW / Panorama:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `(Required) The certificate name.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ssl_tls_service_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"ssl",
				"tls",
				"service",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. NGFW / Panorama:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) SSL certificate file name.`,
				},
				resource.Attribute{
					Name:        "min_version",
					Description: `Minimum TLS protocol version. Valid values are ` + "`" + `"tls1-0"` + "`" + ` (default), ` + "`" + `"tls1-1"` + "`" + `, and ` + "`" + `"tls1-2"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_version",
					Description: `Maximum TLS protocol version. Valid values are ` + "`" + `"tls1-0"` + "`" + `, ` + "`" + `"tls1-1"` + "`" + `, ` + "`" + `"tls1-2"` + "`" + `, and ` + "`" + `max` + "`" + ` (default).`,
				},
				resource.Attribute{
					Name:        "allow_algorithm_rsa",
					Description: `(bool) Allow algorithm RSA (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "allow_algorithm_dhe",
					Description: `(bool) Allow algorithm DHE (defualt: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "allow_algorithm_ecdhe",
					Description: `(bool) Allow algorithm ECDHE (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "allow_algorithm_3des",
					Description: `(bool) Allow algorithm 3DES (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "allow_algorithm_rc4",
					Description: `(bool) Allow algorithm RC4 (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "allow_algorithm_aes_128_cbc",
					Description: `(bool) Allow algorithm AES-128-CBC (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "allow_algorithm_aes_256_cbc",
					Description: `(bool) Allow algorithm AES-256-CBC (defualt: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "allow_algorithm_aes_128_gcm",
					Description: `(bool) Allow algorithm AES-128-GCM (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "allow_algorithm_aes_256_gcm",
					Description: `(bool) Allow algorithm AES-256-GCM (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "allow_authentication_sha1",
					Description: `(bool) Allow authentication SHA1 (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "allow_authentication_sha256",
					Description: `(bool) Allow authentication SHA256 (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "allow_authentication_sha384",
					Description: `(bool) Allow authentication SHA384 (default: ` + "`" + `true` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_static_route_ipv4",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"static",
				"route",
				"ipv4",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The static route's name.`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router to add the static route to.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Destination IP address / prefix.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) Interface to use.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The next hop type. Valid values are ` + "`" + `ip-address` + "`" + ` (the default), ` + "`" + `discard` + "`" + `, ` + "`" + `next-vr` + "`" + `, or an empty string for ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Optional) The value for the ` + "`" + `type` + "`" + ` setting.`,
				},
				resource.Attribute{
					Name:        "admin_distance",
					Description: `(Optional) The admin distance.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional, int) Metric value / path cost (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_table",
					Description: `(Optional) Target routing table to install the route. Valid values are ` + "`" + `unicast` + "`" + ` (the default), ` + "`" + `no install` + "`" + `, ` + "`" + `multicast` + "`" + `, or ` + "`" + `both` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `(Optional, PAN-OS 7.1+) BFD configuration.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_syslog_server_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"syslog",
				"server",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys (default: ` + "`" + `shared` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group's name.`,
				},
				resource.Attribute{
					Name:        "config_format",
					Description: `(Optional) Config format.`,
				},
				resource.Attribute{
					Name:        "system_format",
					Description: `(Optional) System format.`,
				},
				resource.Attribute{
					Name:        "threat_format",
					Description: `(Optional) Threat format.`,
				},
				resource.Attribute{
					Name:        "traffic_format",
					Description: `(Optional) Traffic format.`,
				},
				resource.Attribute{
					Name:        "hip_match_format",
					Description: `(Optional) HIP match format.`,
				},
				resource.Attribute{
					Name:        "url_format",
					Description: `(Optional) URL format.`,
				},
				resource.Attribute{
					Name:        "data_format",
					Description: `(Optional) Data format.`,
				},
				resource.Attribute{
					Name:        "wildfire_format",
					Description: `(Optional) Wildfire format.`,
				},
				resource.Attribute{
					Name:        "tunnel_format",
					Description: `(Optional) Tunnel format.`,
				},
				resource.Attribute{
					Name:        "user_id_format",
					Description: `(Optional) UserID format.`,
				},
				resource.Attribute{
					Name:        "gtp_format",
					Description: `(Optional) GTP format.`,
				},
				resource.Attribute{
					Name:        "auth_format",
					Description: `(Optional) Auth format.`,
				},
				resource.Attribute{
					Name:        "sctp_format",
					Description: `(Optional) SCTP format.`,
				},
				resource.Attribute{
					Name:        "iptag_format",
					Description: `(Optional) IP tag format.`,
				},
				resource.Attribute{
					Name:        "escaped_characters",
					Description: `(Optional) The escaped characters (as a string).`,
				},
				resource.Attribute{
					Name:        "escape_character",
					Description: `(Optional) The escape character.`,
				},
				resource.Attribute{
					Name:        "syslog_server",
					Description: `(Required, repeatable) The server spec (defined below). ` + "`" + `syslog_server` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Server name.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) The hostname.`,
				},
				resource.Attribute{
					Name:        "transport",
					Description: `(Optional) The transport. Valid values are ` + "`" + `UDP` + "`" + ` (default), ` + "`" + `TCP` + "`" + `, or ` + "`" + `SSL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, int) The port number (default: 514).`,
				},
				resource.Attribute{
					Name:        "syslog_format",
					Description: `(Optional) The syslog format. Valid values are ` + "`" + `BSD` + "`" + ` (default) or ` + "`" + `IETF` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(Optional) The syslog facility. Valid values are ` + "`" + `LOG_USER` + "`" + ` (default), ` + "`" + `LOG_LOCAL0` + "`" + `, ` + "`" + `LOG_LOCAL1` + "`" + `, ` + "`" + `LOG_LOCAL2` + "`" + `, ` + "`" + `LOG_LOCAL3` + "`" + `, ` + "`" + `LOG_LOCAL4` + "`" + `, ` + "`" + `LOG_LOCAL5` + "`" + `, ` + "`" + `LOG_LOCAL6` + "`" + `, or ` + "`" + `LOG_LOCAL7` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_tacacs_plus_profile",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"tacacs",
				"plus",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. NGFW / Panorama:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "admin_use_only",
					Description: `(bool) Administrator use only.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(int) Timeout in seconds (default: ` + "`" + `3` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "use_single_connection",
					Description: `(bool) Use single connectino for all authentication."`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `List of TACACS+ server specs, as defined below. ` + "`" + `protocol` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "chap",
					Description: `(bool) CHAP.`,
				},
				resource.Attribute{
					Name:        "pap",
					Description: `(bool) PAP.`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `(bool, PAN-OS 8.0 only) Auto. ` + "`" + `server` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) Server hostname or IP address.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Required) Shared secret for TACACS+ communication.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(int) TACACS+ server port number (default: ` + "`" + `49` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_telemetry",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"telemetry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_reports",
					Description: `(Bool, optional) Application reports.`,
				},
				resource.Attribute{
					Name:        "threat_prevention_reports",
					Description: `(Bool, optional) Threat reports.`,
				},
				resource.Attribute{
					Name:        "url_reports",
					Description: `(Bool, optional) URL reports.`,
				},
				resource.Attribute{
					Name:        "file_type_identification_reports",
					Description: `(Bool, optional) File type identification reports.`,
				},
				resource.Attribute{
					Name:        "threat_prevention_data",
					Description: `(Bool, optional) Threat prevention data.`,
				},
				resource.Attribute{
					Name:        "threat_prevention_packet_captures",
					Description: `(Bool, optional) Enable sending packet- captures with threat prevention information. This requires that ` + "`" + `threat_prevention_data` + "`" + ` also be enabled.`,
				},
				resource.Attribute{
					Name:        "product_usage_stats",
					Description: `(Bool, optional) Health and performance reports.`,
				},
				resource.Attribute{
					Name:        "passive_dns_monitoring",
					Description: `(Bool, optional) Passive DNS monitoring.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_tunnel_interface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"tunnel",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name. This must start with ` + "`" + `tunnel.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys that will use this interface (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_url_filtering_security_profile",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"url",
				"filtering",
				"security",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "dynamic_url",
					Description: `(bool, removed in PAN-OS 9.0) Dynamic URL.`,
				},
				resource.Attribute{
					Name:        "expired_license_action",
					Description: `(bool, removed in PAN-OS 9.0) Expired license action.`,
				},
				resource.Attribute{
					Name:        "block_list_action",
					Description: `(removed in PAN-OS 9.0) Block list action.`,
				},
				resource.Attribute{
					Name:        "block_list",
					Description: `(list, removed in PAN-OS 9.0) Block list.`,
				},
				resource.Attribute{
					Name:        "allow_list",
					Description: `(list, removed in PAN-OS 9.0) Allow list.`,
				},
				resource.Attribute{
					Name:        "allow_categories",
					Description: `(list) List of categories to allow.`,
				},
				resource.Attribute{
					Name:        "alert_categories",
					Description: `(list) List of categories to alert.`,
				},
				resource.Attribute{
					Name:        "block_categories",
					Description: `(list) List of categories to block.`,
				},
				resource.Attribute{
					Name:        "continue_categories",
					Description: `(list) List of categories to continue.`,
				},
				resource.Attribute{
					Name:        "override_categories",
					Description: `(list) List of categories to override.`,
				},
				resource.Attribute{
					Name:        "track_container_page",
					Description: `(bool) Track the container page.`,
				},
				resource.Attribute{
					Name:        "log_container_page_only",
					Description: `(bool) Log container page only.`,
				},
				resource.Attribute{
					Name:        "safe_search_enforcement",
					Description: `(bool) Safe search enforcement.`,
				},
				resource.Attribute{
					Name:        "log_http_header_xff",
					Description: `(bool) HTTP header logging: X-Forwarded-For.`,
				},
				resource.Attribute{
					Name:        "log_http_header_user_agent",
					Description: `(bool) HTTP header logging: User-Agent.`,
				},
				resource.Attribute{
					Name:        "log_http_header_referer",
					Description: `(bool) HTTP header logging: Referer.`,
				},
				resource.Attribute{
					Name:        "ucd_mode",
					Description: `(PAN-OS 8.0+) User credential detection mode. Valid values are ` + "`" + `disabled` + "`" + ` (default), ` + "`" + `ip-user` + "`" + `, ` + "`" + `domain-credentials` + "`" + `, or ` + "`" + `group-mapping` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ucd_mode_group_mapping",
					Description: `(` + "`" + `ucd_mode` + "`" + `=` + "`" + `group-mapping` + "`" + `, PAN-OS 8.0+) User credential detection: the group mapping settings.`,
				},
				resource.Attribute{
					Name:        "ucd_log_severity",
					Description: `(Optional, but Required for PAN-OS 8.0+) User credential detection: valid username detected log severity.`,
				},
				resource.Attribute{
					Name:        "ucd_allow_categories",
					Description: `(list, PAN-OS 8.0+) Categories allowed with user credential submission.`,
				},
				resource.Attribute{
					Name:        "ucd_alert_categories",
					Description: `(list, PAN-OS 8.0+) Categories alerted on with user credential submission.`,
				},
				resource.Attribute{
					Name:        "ucd_block_categories",
					Description: `(list, PAN-OS 8.0+) Categories blocked with user credential submission.`,
				},
				resource.Attribute{
					Name:        "ucd_continue_categories",
					Description: `(list, PAN-OS 8.0+) Categories continued with user credential submission.`,
				},
				resource.Attribute{
					Name:        "http_header_insertion",
					Description: `(repeatable, PAN-OS 8.1+) List of HTTP header insertion specs, as defined below.`,
				},
				resource.Attribute{
					Name:        "machine_learning_model",
					Description: `(repeatable, PAN-OS 10.0+) List of machine learning specs, as defined below.`,
				},
				resource.Attribute{
					Name:        "machine_learning_exceptions",
					Description: `(list) List of machine learning exceptions. ` + "`" + `http_header_insertion` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Header type.`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `(list) Header domains.`,
				},
				resource.Attribute{
					Name:        "http_header",
					Description: `(repeatable) List of HTTP header specs, as defined below. ` + "`" + `http_header_insertion.http_header` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed attribute) Name.`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `(Required) The header.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the header.`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `(bool) Logging of this header. ` + "`" + `machine_learning_model` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `(Required) Machine learning model.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Model action.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_user_tag",
			Category:         "User-ID",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"user",
				"id",
				"tag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys location (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) The user.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(list) List of tags.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_userid_login",
			Category:         "User-ID",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"user",
				"id",
				"userid",
				"login",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys location (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) The user.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP address.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_virtual_router",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"virtual",
				"router",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. Panorama / NGFW:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys to import the virtual router into. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The virtual router's name.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `List of interfaces that should use this virtual router. If you intend to use [` + "`" + `panos_virtual_router_entry` + "`" + `](virtual_router_entry.html) then leave this field undefined.`,
				},
				resource.Attribute{
					Name:        "static_dist",
					Description: `(int) Admin distance - Static (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "static_ipv6_dist",
					Description: `(int) Admin distance - Static IPv6 (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ospf_int_dist",
					Description: `(int) Admin distance - OSPF Int (default: ` + "`" + `30` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ospf_ext_dist",
					Description: `(int) Admin distance - OSPF Ext (default: ` + "`" + `110` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ospfv3_int_dist",
					Description: `(int) Admin distance - OSPFv3 Int (default: ` + "`" + `30` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ospfv3_ext_dist",
					Description: `(int) Admin distance - OSPFv3 Ext (default: ` + "`" + `110` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ibgp_dist",
					Description: `(int) Admin distance - IBGP (default: ` + "`" + `200` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ebgp_dist",
					Description: `(int) Admin distance - EBGP (default: ` + "`" + `20` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "rip_dist",
					Description: `(int) Admin distance - RIP (default: ` + "`" + `120` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_ecmp",
					Description: `(bool) Enable ECMP.`,
				},
				resource.Attribute{
					Name:        "ecmp_max_path",
					Description: `(int) Maximum number of ECMP paths supported.`,
				},
				resource.Attribute{
					Name:        "ecmp_symmetric_return",
					Description: `(bool) Allows return packets to egress out of the ingress interface of the flow.`,
				},
				resource.Attribute{
					Name:        "ecmp_strict_source_path",
					Description: `(bool) Force VPN traffic to exit interface that the source-ip belongs to.`,
				},
				resource.Attribute{
					Name:        "ecmp_load_balance_method",
					Description: `Load balancing algorithm. Valid values are ` + "`" + `ip-modulo` + "`" + `, ` + "`" + `ip-hash` + "`" + `, ` + "`" + `weighted-round-robin` + "`" + `, or ` + "`" + `balanced-round-robin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ecmp_hash_source_only",
					Description: `(bool) For ` + "`" + `ecmp_load_balance_method` + "`" + ` = ` + "`" + `ip-hash` + "`" + `: Only use source address for hash.`,
				},
				resource.Attribute{
					Name:        "ecmp_hash_use_port",
					Description: `(bool) For ` + "`" + `ecmp_load_balance_method` + "`" + ` = ` + "`" + `ip-hash` + "`" + `: Use source/destination port for hash.`,
				},
				resource.Attribute{
					Name:        "ecmp_hash_seed",
					Description: `(int) For ` + "`" + `ecmp_load_balance_method` + "`" + ` = ` + "`" + `ip-hash` + "`" + `: User-specified hash seed.`,
				},
				resource.Attribute{
					Name:        "ecmp_weighted_round_robin_interfaces",
					Description: `(Map of ints) For ` + "`" + `ecmp_load_balance_method` + "`" + ` = ` + "`" + `weighted-round-robin` + "`" + `: Interface weight used in weighted ECMP load balancing.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_virtual_router_entry",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"virtual",
				"router",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router's name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface to import into the virtual router.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_vlan",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"vlan",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the object into (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vlan_interface",
					Description: `(Optional) The VLAN interface.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Optional, computed) List of layer2 interfaces. You can also leave this blank and also use [panos_vlan_entry](vlan_entry.html) for more control.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_vlan_entry",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"vlan",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vlan",
					Description: `(Required) The VLAN's name.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface's name.`,
				},
				resource.Attribute{
					Name:        "mac_addresses",
					Description: `(Optional) List of MAC addresses that should go with this entry.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_vlan_interface",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"vlan",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The interface's name. Must start with ` + "`" + `vlan.` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys that will use this interface (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) The interface comment.`,
				},
				resource.Attribute{
					Name:        "netflow_profile",
					Description: `(Optional) The netflow profile.`,
				},
				resource.Attribute{
					Name:        "static_ips",
					Description: `(Optional) List of static IPv4 addresses to set for this data interface.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable DHCP on this interface.`,
				},
				resource.Attribute{
					Name:        "create_dhcp_default_route",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to create a DHCP default route.`,
				},
				resource.Attribute{
					Name:        "dhcp_default_route_metric",
					Description: `(Optional) The metric for the DHCP default route.`,
				},
				resource.Attribute{
					Name:        "management_profile",
					Description: `(Optional) The management profile.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The MTU.`,
				},
				resource.Attribute{
					Name:        "adjust_tcp_mss",
					Description: `(Optional) Adjust TCP MSS (default: false).`,
				},
				resource.Attribute{
					Name:        "ipv4_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv4 MSS adjust value.`,
				},
				resource.Attribute{
					Name:        "ipv6_mss_adjust",
					Description: `(Optional, PAN-OS 8.0+) The IPv6 MSS adjust value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_vm_auth_key",
			Category:         "Panorama",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"vm",
				"auth",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hours",
					Description: `(int) The VM auth key lifetime in hours.`,
				},
				resource.Attribute{
					Name:        "keepers",
					Description: `(map) Arbitrary map of values that, when changed, will trigger a new key to be generated. ## Attribute Reference The following attributes are supported.`,
				},
				resource.Attribute{
					Name:        "auth_key",
					Description: `The bootstrap VM auth key.`,
				},
				resource.Attribute{
					Name:        "expiry",
					Description: `The date as returned from PAN-OS for when the auth key expires.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `The expiration time as a RFC 3339 string.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `(bool) If the auth key is still valid based on the lifetime given.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "auth_key",
					Description: `The bootstrap VM auth key.`,
				},
				resource.Attribute{
					Name:        "expiry",
					Description: `The date as returned from PAN-OS for when the auth key expires.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `The expiration time as a RFC 3339 string.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `(bool) If the auth key is still valid based on the lifetime given.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_vm_information_source",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"vm",
				"information",
				"source",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. NGFW / Panorama:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The zone's name.`,
				},
				resource.Attribute{
					Name:        "aws_vpc",
					Description: `AWS VPC information source spec (see below).`,
				},
				resource.Attribute{
					Name:        "esxi",
					Description: `VMware ESXi information source spec (see below).`,
				},
				resource.Attribute{
					Name:        "vcenter",
					Description: `VMware vCenter information source spec (see below).`,
				},
				resource.Attribute{
					Name:        "google_compute",
					Description: `Google compute engine information source spec (see below). ` + "`" + `aws_vpc` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(bool) Disabled or not.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) IP address or name.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `(Required) AWS access key ID.`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Required) AWS secret access key.`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `(int) Time interval (in sec) for updates (default: ` + "`" + `60` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_timeout",
					Description: `(bool) Enable vm-info timeout when source is disconnected.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(int) The vm-info timeout value (in hours) when source is disconnected (default: ` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) AWS VPC name or ID. ` + "`" + `esxi` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(int) The port number (default: ` + "`" + `443` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(bool) Disabled or not.`,
				},
				resource.Attribute{
					Name:        "enable_timeout",
					Description: `(bool) Enable vm-info timeout when source is disconnected.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(int) The vm-info timeout value (in hours) when source is disconnected (default: ` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) IP address or source name for vm-info-source.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The vm-info-source login username.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The vm-info-source login password.`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `(int) Time interval (in sec) for updates (default: ` + "`" + `5` + "`" + `). ` + "`" + `vcenter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(int) The port number (default: ` + "`" + `443` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(bool) Disabled or not.`,
				},
				resource.Attribute{
					Name:        "enable_timeout",
					Description: `(bool) Enable vm-info timeout when source is disconnected.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(int) The vm-info timeout value (in hours) when source is disconnected (default: ` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) IP address or source name for vm-info-source.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The vm-info-source login username.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The vm-info-source login password.`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `(int) Time interval (in sec) for updates (default: ` + "`" + `5` + "`" + `). ` + "`" + `google_compute` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(bool) Disabled or not.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `The auth type. Valid values are ` + "`" + `service-in-gce` + "`" + ` (default) or ` + "`" + `service-account` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_account_credential",
					Description: `GCE service account JSON file.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Google Compute Engine Project-ID.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) Google Compute Engine project zone name.`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `(int) Time interval (in sec) for updates (default: ` + "`" + `5` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_timeout",
					Description: `(bool) Enable vm-info timeout when source is disconnected.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(int) The vm-info timeout value (in hours) when source is disconnected (default: ` + "`" + `2` + "`" + `).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_vulnerability_security_profile",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"vulnerability",
				"security",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(repeatable) Rule list spec, as defined below.`,
				},
				resource.Attribute{
					Name:        "exception",
					Description: `(repeatable) Except list spec, as defined below. ` + "`" + `rule` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name.`,
				},
				resource.Attribute{
					Name:        "threat_name",
					Description: `Threat name.`,
				},
				resource.Attribute{
					Name:        "cves",
					Description: `(list) List of CVEs.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host. Valid values are ` + "`" + `any` + "`" + ` (default), ` + "`" + `client` + "`" + ` or ` + "`" + `server` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vendor_ids",
					Description: `(list) List of vendor IDs.`,
				},
				resource.Attribute{
					Name:        "severities",
					Description: `(list) List of severities.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The category (default: ` + "`" + `any` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action. Valid values are ` + "`" + `default` + "`" + ` (default), ` + "`" + `allow` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `drop` + "`" + `, ` + "`" + `reset-client` + "`" + `, ` + "`" + `reset-server` + "`" + `, ` + "`" + `reset-both` + "`" + `, or ` + "`" + `block-ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "block_ip_track_by",
					Description: `(action=` + "`" + `block-ip` + "`" + `) The track by setting.`,
				},
				resource.Attribute{
					Name:        "block_ip_duration",
					Description: `(action=` + "`" + `block-ip` + "`" + `, int) The duration.`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `Packet capture setting. Valid values are ` + "`" + `disable` + "`" + ` (default), ` + "`" + `single-packet` + "`" + `, or ` + "`" + `extended-capture` + "`" + `. ` + "`" + `exception` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Threat name. You can use the ` + "`" + `panos_predefined_threat` + "`" + ` data source to discover the various vulnerability names available to use.`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `Packet capture setting. Valid values are ` + "`" + `disable` + "`" + ` (default), ` + "`" + `single-packet` + "`" + `, or ` + "`" + `extended-capture` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action. Valid values are ` + "`" + `default` + "`" + ` (default), ` + "`" + `allow` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `drop` + "`" + `, ` + "`" + `reset-client` + "`" + `, ` + "`" + `reset-server` + "`" + `, ` + "`" + `reset-both` + "`" + `, or ` + "`" + `block-ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "block_ip_track_by",
					Description: `(action=` + "`" + `block-ip` + "`" + `) The track by setting.`,
				},
				resource.Attribute{
					Name:        "block_ip_duration",
					Description: `(action=` + "`" + `block-ip` + "`" + `, int) The duration.`,
				},
				resource.Attribute{
					Name:        "time_interval",
					Description: `(int) Time interval.`,
				},
				resource.Attribute{
					Name:        "time_threshold",
					Description: `(int) Time threshold.`,
				},
				resource.Attribute{
					Name:        "time_track_by",
					Description: `Time track by setting. Valid values are ` + "`" + `source` + "`" + `, ` + "`" + `destination` + "`" + `, or ` + "`" + `source-and-destination` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "exempt_ips",
					Description: `(list) List of exempt IPs.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_wildfire_analysis_security_profile",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"wildfire",
				"analysis",
				"security",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(repeatable) Rule list spec, as defined below. ` + "`" + `rule` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(list) List of applications.`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `(list) List of file types.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `Direction. Valid values are ` + "`" + `both` + "`" + ` (default), ` + "`" + `upload` + "`" + `, or ` + "`" + `download` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "analysis",
					Description: `Analysis setting. Valid values are ` + "`" + `public-cloud` + "`" + ` (default) or ` + "`" + `private-cloud` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_zone",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. NGFW / Panorama:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The zone's name.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The zone's mode. This can be ` + "`" + `layer3` + "`" + `, ` + "`" + `layer2` + "`" + `, ` + "`" + `virtual-wire` + "`" + `, ` + "`" + `tap` + "`" + `, or ` + "`" + `tunnel` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_profile",
					Description: `The zone protection profile.`,
				},
				resource.Attribute{
					Name:        "log_setting",
					Description: `Log setting.`,
				},
				resource.Attribute{
					Name:        "enable_user_id",
					Description: `Boolean to enable user identification.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `List of interfaces to associated with this zone. Leave this undefined if you want to use [` + "`" + `panos_zone_entry` + "`" + `](zone_entry.html) resources.`,
				},
				resource.Attribute{
					Name:        "include_acls",
					Description: `Users from these addresses/subnets will be identified. This can be an address object, an address group, a single IP address, or an IP address subnet.`,
				},
				resource.Attribute{
					Name:        "exclude_acls",
					Description: `Users from these addresses/subnets will not be identified. This can be an address object, an address group, a single IP address, or an IP address subnet.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_zone_entry",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"zone",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template name.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack name. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys (default: ` + "`" + `vsys1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The zone's name.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The mode. Can be ` + "`" + `layer3` + "`" + ` (default), ` + "`" + `layer2` + "`" + `, ` + "`" + `virtual-wire` + "`" + `, ` + "`" + `tap` + "`" + `, or ` + "`" + `external` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The interface's name.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"panos_address_group":                                 0,
		"panos_address_object":                                1,
		"panos_address_objects":                               2,
		"panos_administrative_tag":                            3,
		"panos_aggregate_interface":                           4,
		"panos_anti_spyware_security_profile":                 5,
		"panos_antivirus_security_profile":                    6,
		"panos_application_group":                             7,
		"panos_application_object":                            8,
		"panos_application_signature":                         9,
		"panos_arp":                                           10,
		"panos_authentication_profile":                        11,
		"panos_aws_cloud_watch":                               12,
		"panos_bfd_profile":                                   13,
		"panos_bgp":                                           14,
		"panos_bgp_aggregate":                                 15,
		"panos_bgp_aggregate_advertise_filter":                16,
		"panos_bgp_aggregate_suppress_filter":                 17,
		"panos_bgp_auth_profile":                              18,
		"panos_bgp_conditional_adv":                           19,
		"panos_bgp_conditional_adv_advertise_filter":          20,
		"panos_bgp_conditional_adv_non_exist_filter":          21,
		"panos_bgp_dampening_profile":                         22,
		"panos_bgp_export_rule_group":                         23,
		"panos_bgp_import_rule_group":                         24,
		"panos_bgp_peer":                                      25,
		"panos_bgp_peer_group":                                26,
		"panos_bgp_redist_rule":                               27,
		"panos_certificate_import":                            28,
		"panos_certificate_profile":                           29,
		"panos_custom_data_pattern_object":                    30,
		"panos_custom_url_category":                           31,
		"panos_custom_url_category_entry":                     32,
		"panos_dag_tags":                                      33,
		"panos_data_filtering_security_profile":               34,
		"panos_decryption_rule_group":                         35,
		"panos_device_group":                                  36,
		"panos_device_group_entry":                            37,
		"panos_device_group_parent":                           38,
		"panos_dos_protection_profile":                        39,
		"panos_dynamic_user_group":                            40,
		"panos_edl":                                           41,
		"panos_email_server_profile":                          42,
		"panos_ethernet_interface":                            43,
		"panos_file_blocking_security_profile":                44,
		"panos_general_settings":                              45,
		"panos_globalprotect_ipsec_crypto_profile":            46,
		"panos_gre_tunnel":                                    47,
		"panos_http_server_profile":                           48,
		"panos_ike_crypto_profile":                            49,
		"panos_ike_gateway":                                   50,
		"panos_ip_tag":                                        51,
		"panos_ipsec_crypto_profile":                          52,
		"panos_ipsec_tunnel":                                  53,
		"panos_ipsec_tunnel_proxy_id_ipv4":                    54,
		"panos_kerberos_profile":                              55,
		"panos_layer2_subinterface":                           56,
		"panos_layer3_subinterface":                           57,
		"panos_ldap_profile":                                  58,
		"panos_license_api_key":                               59,
		"panos_licensing":                                     60,
		"panos_local_user_db_group":                           61,
		"panos_local_user_db_user":                            62,
		"panos_log_forwarding_profile":                        63,
		"panos_loopback_interface":                            64,
		"panos_management_profile":                            65,
		"panos_monitor_profile":                               66,
		"panos_nat_rule":                                      67,
		"panos_nat_rule_group":                                68,
		"panos_ospf":                                          69,
		"panos_ospf_area":                                     70,
		"panos_ospf_area_interface":                           71,
		"panos_ospf_area_virtual_link":                        72,
		"panos_ospf_auth_profile":                             73,
		"panos_ospf_export":                                   74,
		"panos_panorama_address_group":                        75,
		"panos_panorama_administrative_tag":                   76,
		"panos_panorama_aggregate_interface":                  77,
		"panos_panorama_application_group":                    78,
		"panos_panorama_application_object":                   79,
		"panos_panorama_application_signature":                80,
		"panos_panorama_bfd_profile":                          81,
		"panos_panorama_bgp":                                  82,
		"panos_panorama_bgp_aggregate":                        83,
		"panos_panorama_bgp_aggregate_advertise_filter":       84,
		"panos_panorama_bgp_aggregate_suppress_filter":        85,
		"panos_panorama_bgp_auth_profile":                     86,
		"panos_panorama_bgp_conditional_adv":                  87,
		"panos_panorama_bgp_conditional_adv_advertise_filter": 88,
		"panos_panorama_bgp_conditional_adv_non_exist_filter": 89,
		"panos_panorama_bgp_dampening_profile":                90,
		"panos_panorama_bgp_export_rule_group":                91,
		"panos_panorama_bgp_import_rule_group":                92,
		"panos_panorama_bgp_peer":                             93,
		"panos_panorama_bgp_peer_group":                       94,
		"panos_panorama_bgp_redist_rule":                      95,
		"panos_panorama_device_group":                         96,
		"panos_panorama_device_group_entry":                   97,
		"panos_panorama_edl":                                  98,
		"panos_panorama_email_server_profile":                 99,
		"panos_panorama_ethernet_interface":                   100,
		"panos_panorama_gcp_account":                          101,
		"panos_panorama_gke_cluster":                          102,
		"panos_panorama_gke_cluster_group":                    103,
		"panos_panorama_gre_tunnel":                           104,
		"panos_panorama_http_server_profile":                  105,
		"panos_panorama_ike_crypto_profile":                   106,
		"panos_panorama_ike_gateway":                          107,
		"panos_panorama_ipsec_crypto_profile":                 108,
		"panos_panorama_ipsec_tunnel":                         109,
		"panos_panorama_ipsec_tunnel_proxy_id_ipv4":           110,
		"panos_panorama_layer2_subinterface":                  111,
		"panos_panorama_layer3_subinterface":                  112,
		"panos_panorama_log_forwarding_profile":               113,
		"panos_panorama_loopback_interface":                   114,
		"panos_panorama_management_profile":                   115,
		"panos_panorama_monitor_profile":                      116,
		"panos_panorama_nat_rule":                             117,
		"panos_panorama_nat_rule_group":                       118,
		"panos_panorama_pbf_rule_group":                       119,
		"panos_panorama_redistribution_profile_ipv4":          120,
		"panos_panorama_security_policy":                      121,
		"panos_panorama_security_rule_group":                  122,
		"panos_panorama_service_group":                        123,
		"panos_panorama_service_object":                       124,
		"panos_panorama_snmptrap_server_profile":              125,
		"panos_panorama_static_route_ipv4":                    126,
		"panos_panorama_syslog_server_profile":                127,
		"panos_panorama_template":                             128,
		"panos_panorama_template_entry":                       129,
		"panos_panorama_template_stack":                       130,
		"panos_panorama_template_stack_entry":                 131,
		"panos_panorama_template_variable":                    132,
		"panos_panorama_tunnel_interface":                     133,
		"panos_panorama_virtual_router":                       134,
		"panos_panorama_virtual_router_entry":                 135,
		"panos_panorama_vlan":                                 136,
		"panos_panorama_vlan_entry":                           137,
		"panos_panorama_vlan_interface":                       138,
		"panos_panorama_zone":                                 139,
		"panos_panorama_zone_entry":                           140,
		"panos_pbf_rule_group":                                141,
		"panos_radius_profile":                                142,
		"panos_redistribution_profile_ipv4":                   143,
		"panos_saml_profile":                                  144,
		"panos_security_policy":                               145,
		"panos_security_profile_group":                        146,
		"panos_security_rule_group":                           147,
		"panos_service_group":                                 148,
		"panos_service_object":                                149,
		"panos_snmptrap_server_profile":                       150,
		"panos_ssl_decrypt":                                   151,
		"panos_ssl_decrypt_exclude_certificate_entry":         152,
		"panos_ssl_decrypt_trusted_root_ca_entry":             153,
		"panos_ssl_tls_service_profile":                       154,
		"panos_static_route_ipv4":                             155,
		"panos_syslog_server_profile":                         156,
		"panos_tacacs_plus_profile":                           157,
		"panos_telemetry":                                     158,
		"panos_tunnel_interface":                              159,
		"panos_url_filtering_security_profile":                160,
		"panos_user_tag":                                      161,
		"panos_userid_login":                                  162,
		"panos_virtual_router":                                163,
		"panos_virtual_router_entry":                          164,
		"panos_vlan":                                          165,
		"panos_vlan_entry":                                    166,
		"panos_vlan_interface":                                167,
		"panos_vm_auth_key":                                   168,
		"panos_vm_information_source":                         169,
		"panos_vulnerability_security_profile":                170,
		"panos_wildfire_analysis_security_profile":            171,
		"panos_zone":                                          172,
		"panos_zone_entry":                                    173,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
