package panos

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

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
					Description: `(Required) The address object's name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of address object.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The address object's value.`,
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of address object.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The address object's value.`,
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
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
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
					Description: `(Required) The name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `Packet capture setting. Valid values are ` + "`" + `disable` + "`" + `, ` + "`" + `single-packet` + "`" + `, or ` + "`" + `extended-capture` + "`" + `.`,
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
					Description: `(repeatable) Except list spec, as defined below. ` + "`" + `botnet_list` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action to take. Valid values are ` + "`" + `alert` + "`" + `, ` + "`" + `default` + "`" + `, ` + "`" + `allow` + "`" + `, ` + "`" + `block` + "`" + `, or ` + "`" + `sinkhole` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `(PAN-OS 9.0+) Packet capture setting. Valid values are ` + "`" + `disable` + "`" + `, ` + "`" + `single-packet` + "`" + `, or ` + "`" + `extended-capture` + "`" + `. ` + "`" + `dns_category` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
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
					Description: `Packet capture setting. Valid values are ` + "`" + `disable` + "`" + `, ` + "`" + `single-packet` + "`" + `, or ` + "`" + `extended-capture` + "`" + `. ` + "`" + `white_list` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description ` + "`" + `rule` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "threat_name",
					Description: `Threat name.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Category.`,
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
					Description: `(action=` + "`" + `block-ip` + "`" + `, int) The duration. ` + "`" + `exception` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Threat name. You can use the ` + "`" + `panos_predefined_threat` + "`" + ` data source to discover the various phone home names available to use.`,
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `Packet capture setting. Valid values are ` + "`" + `disable` + "`" + `, ` + "`" + `single-packet` + "`" + `, or ` + "`" + `extended-capture` + "`" + `.`,
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
					Description: `(repeatable) Except list spec, as defined below. ` + "`" + `botnet_list` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action to take. Valid values are ` + "`" + `alert` + "`" + `, ` + "`" + `default` + "`" + `, ` + "`" + `allow` + "`" + `, ` + "`" + `block` + "`" + `, or ` + "`" + `sinkhole` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `(PAN-OS 9.0+) Packet capture setting. Valid values are ` + "`" + `disable` + "`" + `, ` + "`" + `single-packet` + "`" + `, or ` + "`" + `extended-capture` + "`" + `. ` + "`" + `dns_category` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
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
					Description: `Packet capture setting. Valid values are ` + "`" + `disable` + "`" + `, ` + "`" + `single-packet` + "`" + `, or ` + "`" + `extended-capture` + "`" + `. ` + "`" + `white_list` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description ` + "`" + `rule` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "threat_name",
					Description: `Threat name.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Category.`,
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
					Description: `(action=` + "`" + `block-ip` + "`" + `, int) The duration. ` + "`" + `exception` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Threat name. You can use the ` + "`" + `panos_predefined_threat` + "`" + ` data source to discover the various phone home names available to use.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_anti_spyware_security_profiles",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"anti",
				"spyware",
				"security",
				"profiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
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
					Description: `(Required) The name. ## Attribute Reference The following attributes are supported:`,
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
					Description: `(Repeatable) A machine learning exception spec, as defined below. ` + "`" + `decoder` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Decoder name.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Decoder action.`,
				},
				resource.Attribute{
					Name:        "wildfire_action",
					Description: `Wildfire action.`,
				},
				resource.Attribute{
					Name:        "machine_learning_action",
					Description: `(PAN-OS 10.0+) Machine learning action. ` + "`" + `application_exception` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `The application name`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action. ` + "`" + `machine_learning_model` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `The model.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action. ` + "`" + `machine_learning_exception` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name.`,
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
			Attributes: []resource.Attribute{
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
					Description: `(Repeatable) A machine learning exception spec, as defined below. ` + "`" + `decoder` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Decoder name.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Decoder action.`,
				},
				resource.Attribute{
					Name:        "wildfire_action",
					Description: `Wildfire action.`,
				},
				resource.Attribute{
					Name:        "machine_learning_action",
					Description: `(PAN-OS 10.0+) Machine learning action. ` + "`" + `application_exception` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `The application name`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action. ` + "`" + `machine_learning_model` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `The model.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action. ` + "`" + `machine_learning_exception` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_antivirus_security_profiles",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"antivirus",
				"security",
				"profiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_api_key",
			Category:         "Operational State",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"operational",
				"state",
				"api",
				"key",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_key",
					Description: `The API key`,
				},
			},
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
					Description: `(Required) The IP address. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(` + "`" + `interface_type` + "`" + `=` + "`" + `vlan` + "`" + `) The interface.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(` + "`" + `interface_type` + "`" + `=` + "`" + `vlan` + "`" + `) The interface.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_arps",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"arps",
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
					Description: `The subinterface name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
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
					Description: `(Required) The name. ## Attribute Reference The following attributes are available:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type.`,
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
					Description: `The name.`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `(list) List of file types. ` + "`" + `regex` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `(list) List of file types.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `The regex. ` + "`" + `file_property` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "file_type",
					Description: `The file type.`,
				},
				resource.Attribute{
					Name:        "file_property",
					Description: `File property.`,
				},
				resource.Attribute{
					Name:        "property_value",
					Description: `Property value.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type.`,
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
					Description: `The name.`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `(list) List of file types. ` + "`" + `regex` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `(list) List of file types.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `The regex. ` + "`" + `file_property` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "file_type",
					Description: `The file type.`,
				},
				resource.Attribute{
					Name:        "file_property",
					Description: `File property.`,
				},
				resource.Attribute{
					Name:        "property_value",
					Description: `Property value.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_custom_data_pattern_objects",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"custom",
				"data",
				"pattern",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
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
					Description: `(Required) The name. ## Attribute Reference The following attributes are supported:`,
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
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "data_pattern",
					Description: `The data pattern name.`,
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
					Description: `Direction.`,
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
					Description: `Log severity.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "data_pattern",
					Description: `The data pattern name.`,
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
					Description: `Direction.`,
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
					Description: `Log severity.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_data_filtering_security_profiles",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"data",
				"filtering",
				"security",
				"profiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
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
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of entries (device groups).`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `(map of strings) Map of strings where the key is the device group name and the value is the parent for that device group. An empty string for the value means that the parent is the "shared" device group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_dhcp_interface_info",
			Category:         "Operational State",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"operational",
				"state",
				"dhcp",
				"interface",
				"info",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The data interface to get DHCP information for. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The interface's state.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `DHCP IP address.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The default gateway assigned.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The DHCP server IP`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `DHCP server ID`,
				},
				resource.Attribute{
					Name:        "primary_dns",
					Description: `Primary DNS server`,
				},
				resource.Attribute{
					Name:        "secondary_dns",
					Description: `Secondary DNS server`,
				},
				resource.Attribute{
					Name:        "primary_wins",
					Description: `Primary WINS server`,
				},
				resource.Attribute{
					Name:        "secondary_wins",
					Description: `Secondary WINS`,
				},
				resource.Attribute{
					Name:        "primary_nis",
					Description: `Primary NIS`,
				},
				resource.Attribute{
					Name:        "secondary_nis",
					Description: `Secondary NIS`,
				},
				resource.Attribute{
					Name:        "primary_ntp",
					Description: `Primary NTP`,
				},
				resource.Attribute{
					Name:        "secondary_ntp",
					Description: `Secondary NTP`,
				},
				resource.Attribute{
					Name:        "pop3_server",
					Description: `POP3 Server`,
				},
				resource.Attribute{
					Name:        "smtp_server",
					Description: `SMTP Server`,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `DNS Suffix`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `The interface's state.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `DHCP IP address.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The default gateway assigned.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The DHCP server IP`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `DHCP server ID`,
				},
				resource.Attribute{
					Name:        "primary_dns",
					Description: `Primary DNS server`,
				},
				resource.Attribute{
					Name:        "secondary_dns",
					Description: `Secondary DNS server`,
				},
				resource.Attribute{
					Name:        "primary_wins",
					Description: `Primary WINS server`,
				},
				resource.Attribute{
					Name:        "secondary_wins",
					Description: `Secondary WINS`,
				},
				resource.Attribute{
					Name:        "primary_nis",
					Description: `Primary NIS`,
				},
				resource.Attribute{
					Name:        "secondary_nis",
					Description: `Secondary NIS`,
				},
				resource.Attribute{
					Name:        "primary_ntp",
					Description: `Primary NTP`,
				},
				resource.Attribute{
					Name:        "secondary_ntp",
					Description: `Secondary NTP`,
				},
				resource.Attribute{
					Name:        "pop3_server",
					Description: `POP3 Server`,
				},
				resource.Attribute{
					Name:        "smtp_server",
					Description: `SMTP Server`,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `DNS Suffix`,
				},
			},
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
					Description: `(Required) The name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The profile type.`,
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
					Description: `Optional other IP flood protection spec, as defined below. ` + "`" + `syn` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(bool) Enable`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `SYN protection action.`,
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
					Description: `(int) Block duration. ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `icmpv6` + "`" + `, and ` + "`" + `other` + "`" + ` all support the following attributes:`,
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The profile type.`,
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
					Description: `Optional other IP flood protection spec, as defined below. ` + "`" + `syn` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(bool) Enable`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `SYN protection action.`,
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
					Description: `(int) Block duration. ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `icmpv6` + "`" + `, and ` + "`" + `other` + "`" + ` all support the following attributes:`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_dos_protection_profiles",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"dos",
				"protection",
				"profiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
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
					Description: `(Required) Name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `The filter.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(list) List of administrative tags.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `The filter.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(list) List of administrative tags.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_dynamic_user_groups",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"dynamic",
				"user",
				"groups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys to put the address object into (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
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
					Description: `(Required) The name. ## Attribute Reference The following attributes are supported:`,
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
					Description: `Name.`,
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
					Description: `The direction.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action to take.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `Name.`,
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
					Description: `The direction.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action to take.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_file_blocking_security_profiles",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"file",
				"blocking",
				"security",
				"profiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ip_tag",
			Category:         "NGFW User-ID",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"ngfw",
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
					Description: `Filter on a specific IP address.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Filter on a specific tag. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of entries.`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `List of entries specs, as defined below. ` + "`" + `entries` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IP address.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(list) List of tags.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of entries.`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `List of entries specs, as defined below. ` + "`" + `entries` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IP address.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(list) List of tags.`,
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
					Description: `(Required) The virtual router name. ## Attribute Reference The following attributes are supported:`,
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
			Attributes: []resource.Attribute{
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
					Description: `(Required) Name. ## Attribute Reference The following attributes are available:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Area type.`,
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
					Description: `(nssa) Advertise type.`,
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
					Description: `Action.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `Area type.`,
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
					Description: `(nssa) Advertise type.`,
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
					Description: `Action.`,
				},
			},
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
					Description: `(Required) Interface name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(bool) Enable.`,
				},
				resource.Attribute{
					Name:        "passive",
					Description: `(bool) Passive.`,
				},
				resource.Attribute{
					Name:        "link_type",
					Description: `Link type.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(int) Metric.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(int) Priority.`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `(int) Hello interval in seconds.`,
				},
				resource.Attribute{
					Name:        "dead_counts",
					Description: `(int) Dead counts.`,
				},
				resource.Attribute{
					Name:        "retransmit_interval",
					Description: `(int) Retransmit interval in seconds.`,
				},
				resource.Attribute{
					Name:        "transit_delay",
					Description: `(int) Transit delay in seconds.`,
				},
				resource.Attribute{
					Name:        "grace_restart_delay",
					Description: `(int) Graceful restart hello delay in seconds.`,
				},
				resource.Attribute{
					Name:        "auth_profile",
					Description: `Auth profile.`,
				},
				resource.Attribute{
					Name:        "neighbors",
					Description: `(list) List of neighbors.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `BFD profile.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "enable",
					Description: `(bool) Enable.`,
				},
				resource.Attribute{
					Name:        "passive",
					Description: `(bool) Passive.`,
				},
				resource.Attribute{
					Name:        "link_type",
					Description: `Link type.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(int) Metric.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(int) Priority.`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `(int) Hello interval in seconds.`,
				},
				resource.Attribute{
					Name:        "dead_counts",
					Description: `(int) Dead counts.`,
				},
				resource.Attribute{
					Name:        "retransmit_interval",
					Description: `(int) Retransmit interval in seconds.`,
				},
				resource.Attribute{
					Name:        "transit_delay",
					Description: `(int) Transit delay in seconds.`,
				},
				resource.Attribute{
					Name:        "grace_restart_delay",
					Description: `(int) Graceful restart hello delay in seconds.`,
				},
				resource.Attribute{
					Name:        "auth_profile",
					Description: `Auth profile.`,
				},
				resource.Attribute{
					Name:        "neighbors",
					Description: `(list) List of neighbors.`,
				},
				resource.Attribute{
					Name:        "bfd_profile",
					Description: `BFD profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ospf_area_interfaces",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ospf",
				"area",
				"interfaces",
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
					Description: `(Required) OSPF area name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
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
					Description: `(Required) Interface name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(bool) Enable.`,
				},
				resource.Attribute{
					Name:        "neighbor_id",
					Description: `Neighbor ID.`,
				},
				resource.Attribute{
					Name:        "transit_area_id",
					Description: `Transit area ID.`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `(int) Hello interval in seconds.`,
				},
				resource.Attribute{
					Name:        "dead_counts",
					Description: `(int) Dead counts.`,
				},
				resource.Attribute{
					Name:        "retransmit_interval",
					Description: `(int) Retransmit interval in seconds.`,
				},
				resource.Attribute{
					Name:        "transit_delay",
					Description: `(int) Transit delay in seconds.`,
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "enable",
					Description: `(bool) Enable.`,
				},
				resource.Attribute{
					Name:        "neighbor_id",
					Description: `Neighbor ID.`,
				},
				resource.Attribute{
					Name:        "transit_area_id",
					Description: `Transit area ID.`,
				},
				resource.Attribute{
					Name:        "hello_interval",
					Description: `(int) Hello interval in seconds.`,
				},
				resource.Attribute{
					Name:        "dead_counts",
					Description: `(int) Dead counts.`,
				},
				resource.Attribute{
					Name:        "retransmit_interval",
					Description: `(int) Retransmit interval in seconds.`,
				},
				resource.Attribute{
					Name:        "transit_delay",
					Description: `(int) Transit delay in seconds.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ospf_area_virtual_links",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ospf",
				"area",
				"virtual",
				"links",
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
					Description: `(Required) OSPF area name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ospf_areas",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ospf",
				"areas",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Optional, but Required for Panorama) The template location. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router name. ## Attribute Reference The following attributes are available:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ospf_auth_profiles",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ospf",
				"auth",
				"profiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Optional, but Required for Panorama) The template location. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router name. ## Attribute Reference The following attributes are available:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
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
					Description: `(Required) The export rule name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "path_type",
					Description: `Path type.`,
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "path_type",
					Description: `Path type.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_ospf_exports",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"ospf",
				"exports",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `(Optional, but Required for Panorama) The template location. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "virtual_router",
					Description: `(Required) The virtual router name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_plugin",
			Category:         "Operational State",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"operational",
				"state",
				"plugin",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "installed",
					Description: `A list of installed plugins.`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of plugins, installed or not.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `A list of maps (see below). The following attributes are present in each ` + "`" + `details` + "`" + ` entry:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version number.`,
				},
				resource.Attribute{
					Name:        "release_date",
					Description: `Release date.`,
				},
				resource.Attribute{
					Name:        "release_note_url",
					Description: `Release note URL.`,
				},
				resource.Attribute{
					Name:        "package_file",
					Description: `The package file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `Platform.`,
				},
				resource.Attribute{
					Name:        "installed",
					Description: `If the package is installed or not.`,
				},
				resource.Attribute{
					Name:        "downloaded",
					Description: `If the package is downloaded or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_predefined_dlp_file_type",
			Category:         "Predefined",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"predefined",
				"dlp",
				"file",
				"type",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Filter on this specific name.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `A specific label to filter on. ## Attribute Reference The following attributes are supported.`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of file types.`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `List of file types structs, as defined below. ` + "`" + `file_types` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The file type.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `List of property specs, as defined below. ` + "`" + `file_types.properties` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The DLP property name.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The DLP property label.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of file types.`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `List of file types structs, as defined below. ` + "`" + `file_types` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The file type.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `List of property specs, as defined below. ` + "`" + `file_types.properties` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The DLP property name.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The DLP property label.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_predefined_tdb_file_type",
			Category:         "Predefined",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"predefined",
				"tdb",
				"file",
				"type",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Filter on this specific file type.`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `The full name.`,
				},
				resource.Attribute{
					Name:        "full_name_regex",
					Description: `A regex to match against the full name.`,
				},
				resource.Attribute{
					Name:        "data_ident_only",
					Description: `(bool) Limit results to those with data_ident=` + "`" + `true` + "`" + `. ## Attribute Reference The following attributes are supported.`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of file types.`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `List of file types structs, as defined below. ` + "`" + `file_types` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The file type.`,
				},
				resource.Attribute{
					Name:        "file_type_id",
					Description: `(int) The ID`,
				},
				resource.Attribute{
					Name:        "threat_name",
					Description: `The threat name`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `The full name.`,
				},
				resource.Attribute{
					Name:        "data_ident",
					Description: `(bool) Data ident`,
				},
				resource.Attribute{
					Name:        "file_type_ident",
					Description: `(bool) File type ident`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of file types.`,
				},
				resource.Attribute{
					Name:        "file_types",
					Description: `List of file types structs, as defined below. ` + "`" + `file_types` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The file type.`,
				},
				resource.Attribute{
					Name:        "file_type_id",
					Description: `(int) The ID`,
				},
				resource.Attribute{
					Name:        "threat_name",
					Description: `The threat name`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `The full name.`,
				},
				resource.Attribute{
					Name:        "data_ident",
					Description: `(bool) Data ident`,
				},
				resource.Attribute{
					Name:        "file_type_ident",
					Description: `(bool) File type ident`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_predefined_threat",
			Category:         "Predefined",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"predefined",
				"threat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `A specific threat ID / name.`,
				},
				resource.Attribute{
					Name:        "threat_regex",
					Description: `A regex to apply to the threat name.`,
				},
				resource.Attribute{
					Name:        "threat_name",
					Description: `An exact match against the threat name.`,
				},
				resource.Attribute{
					Name:        "threat_type",
					Description: `The threat type. Valid values are ` + "`" + `phone-home` + "`" + ` (default) or ` + "`" + `vulnerability` + "`" + `. ## Attribute Reference The following attributes are supported.`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of file types.`,
				},
				resource.Attribute{
					Name:        "threats",
					Description: `List of matched threats specs, as defined below. ` + "`" + `threats` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The threat name / ID.`,
				},
				resource.Attribute{
					Name:        "threat_name",
					Description: `The threat name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of file types.`,
				},
				resource.Attribute{
					Name:        "threats",
					Description: `List of matched threats specs, as defined below. ` + "`" + `threats` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The threat name / ID.`,
				},
				resource.Attribute{
					Name:        "threat_name",
					Description: `The threat name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_system_info",
			Category:         "Operational State",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"operational",
				"state",
				"system",
				"info",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "info",
					Description: `a map containing the contents of ` + "`" + `show system info` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version_major",
					Description: `Major version number.`,
				},
				resource.Attribute{
					Name:        "version_minor",
					Description: `Minor version number.`,
				},
				resource.Attribute{
					Name:        "version_patch",
					Description: `Patch version number.`,
				},
			},
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
					Description: `(Required) The name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "dynamic_url",
					Description: `(bool) Dynamic URL.`,
				},
				resource.Attribute{
					Name:        "expired_license_action",
					Description: `(bool) Expired license action.`,
				},
				resource.Attribute{
					Name:        "block_list_action",
					Description: `Block list action.`,
				},
				resource.Attribute{
					Name:        "block_list",
					Description: `(list) Block list.`,
				},
				resource.Attribute{
					Name:        "allow_list",
					Description: `(list) Allow list.`,
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
					Description: `User credential detection mode.`,
				},
				resource.Attribute{
					Name:        "ucd_mode_group_mapping",
					Description: `(` + "`" + `ucd_mode` + "`" + `=` + "`" + `group-mapping` + "`" + `) User credential detection: the group mapping settings.`,
				},
				resource.Attribute{
					Name:        "ucd_log_severity",
					Description: `User credential detection: valid username detected log severity.`,
				},
				resource.Attribute{
					Name:        "ucd_allow_categories",
					Description: `(list) Categories allowed with user credential submission.`,
				},
				resource.Attribute{
					Name:        "ucd_alert_categories",
					Description: `(list) Categories alerted on with user credential submission.`,
				},
				resource.Attribute{
					Name:        "ucd_block_categories",
					Description: `(list) Categories blocked with user credential submission.`,
				},
				resource.Attribute{
					Name:        "ucd_continue_categories",
					Description: `(list) Categories continued with user credential submission.`,
				},
				resource.Attribute{
					Name:        "http_header_insertion",
					Description: `(repeatable) List of HTTP header insertion specs, as defined below.`,
				},
				resource.Attribute{
					Name:        "machine_learning_model",
					Description: `(repeatable) List of machine learning specs, as defined below.`,
				},
				resource.Attribute{
					Name:        "machine_learning_exceptions",
					Description: `(list) List of machine learning exceptions. ` + "`" + `http_header_insertion` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
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
					Description: `(repeatable) List of HTTP header specs, as defined below. ` + "`" + `http_header_insertion.http_header` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `The header.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the header.`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `(bool) Logging of this header. ` + "`" + `machine_learning_model` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `Machine learning model.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Model action.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "dynamic_url",
					Description: `(bool) Dynamic URL.`,
				},
				resource.Attribute{
					Name:        "expired_license_action",
					Description: `(bool) Expired license action.`,
				},
				resource.Attribute{
					Name:        "block_list_action",
					Description: `Block list action.`,
				},
				resource.Attribute{
					Name:        "block_list",
					Description: `(list) Block list.`,
				},
				resource.Attribute{
					Name:        "allow_list",
					Description: `(list) Allow list.`,
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
					Description: `User credential detection mode.`,
				},
				resource.Attribute{
					Name:        "ucd_mode_group_mapping",
					Description: `(` + "`" + `ucd_mode` + "`" + `=` + "`" + `group-mapping` + "`" + `) User credential detection: the group mapping settings.`,
				},
				resource.Attribute{
					Name:        "ucd_log_severity",
					Description: `User credential detection: valid username detected log severity.`,
				},
				resource.Attribute{
					Name:        "ucd_allow_categories",
					Description: `(list) Categories allowed with user credential submission.`,
				},
				resource.Attribute{
					Name:        "ucd_alert_categories",
					Description: `(list) Categories alerted on with user credential submission.`,
				},
				resource.Attribute{
					Name:        "ucd_block_categories",
					Description: `(list) Categories blocked with user credential submission.`,
				},
				resource.Attribute{
					Name:        "ucd_continue_categories",
					Description: `(list) Categories continued with user credential submission.`,
				},
				resource.Attribute{
					Name:        "http_header_insertion",
					Description: `(repeatable) List of HTTP header insertion specs, as defined below.`,
				},
				resource.Attribute{
					Name:        "machine_learning_model",
					Description: `(repeatable) List of machine learning specs, as defined below.`,
				},
				resource.Attribute{
					Name:        "machine_learning_exceptions",
					Description: `(list) List of machine learning exceptions. ` + "`" + `http_header_insertion` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
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
					Description: `(repeatable) List of HTTP header specs, as defined below. ` + "`" + `http_header_insertion.http_header` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `The header.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the header.`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `(bool) Logging of this header. ` + "`" + `machine_learning_model` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `Machine learning model.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Model action.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_url_filtering_security_profiles",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"url",
				"filtering",
				"security",
				"profiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_user_tag",
			Category:         "NGFW User-ID",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"ngfw",
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
					Description: `The user. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `A list of entries specs, as defined below. ` + "`" + `entries` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `The user.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(list) List of tags.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "entries",
					Description: `A list of entries specs, as defined below. ` + "`" + `entries` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `The user.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(list) List of tags.`,
				},
			},
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
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of entries.`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `List of entry structs, as defined below. ` + "`" + `entries` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "auth_key",
					Description: `The VM auth key.`,
				},
				resource.Attribute{
					Name:        "expiry",
					Description: `The expiry time as reported by PAN-OS`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `(bool) If the VM auth key is still valid or not.`,
				},
			},
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
					Description: `(Required) The name. ## Attribute Reference The following attributes are supported:`,
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
					Description: `Name.`,
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
					Description: `The host.`,
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
					Description: `The category.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action.`,
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
					Description: `Packet capture setting. ` + "`" + `exception` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Threat name.`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `Packet capture setting.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action.`,
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
					Description: `Time track by setting.`,
				},
				resource.Attribute{
					Name:        "exempt_ips",
					Description: `(list) List of exempt IPs.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `Name.`,
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
					Description: `The host.`,
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
					Description: `The category.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action.`,
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
					Description: `Packet capture setting. ` + "`" + `exception` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Threat name.`,
				},
				resource.Attribute{
					Name:        "packet_capture",
					Description: `Packet capture setting.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action.`,
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
					Description: `Time track by setting.`,
				},
				resource.Attribute{
					Name:        "exempt_ips",
					Description: `(list) List of exempt IPs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_vulnerability_security_profiles",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"vulnerability",
				"security",
				"profiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
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
			Type:             "panos_wildfire_analysis_security_profiles",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"wildfire",
				"analysis",
				"security",
				"profiles",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys location (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) The number of items present.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `(list) A list of the items present.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"panos_address_object":                      0,
		"panos_address_objects":                     1,
		"panos_anti_spyware_security_profile":       2,
		"panos_anti_spyware_security_profiles":      3,
		"panos_antivirus_security_profile":          4,
		"panos_antivirus_security_profiles":         5,
		"panos_api_key":                             6,
		"panos_arp":                                 7,
		"panos_arps":                                8,
		"panos_custom_data_pattern_object":          9,
		"panos_custom_data_pattern_objects":         10,
		"panos_data_filtering_security_profile":     11,
		"panos_data_filtering_security_profiles":    12,
		"panos_device_group_parent":                 13,
		"panos_dhcp_interface_info":                 14,
		"panos_dos_protection_profile":              15,
		"panos_dos_protection_profiles":             16,
		"panos_dynamic_user_group":                  17,
		"panos_dynamic_user_groups":                 18,
		"panos_file_blocking_security_profile":      19,
		"panos_file_blocking_security_profiles":     20,
		"panos_ip_tag":                              21,
		"panos_ospf":                                22,
		"panos_ospf_area":                           23,
		"panos_ospf_area_interface":                 24,
		"panos_ospf_area_interfaces":                25,
		"panos_ospf_area_virtual_link":              26,
		"panos_ospf_area_virtual_links":             27,
		"panos_ospf_areas":                          28,
		"panos_ospf_auth_profiles":                  29,
		"panos_ospf_export":                         30,
		"panos_ospf_exports":                        31,
		"panos_plugin":                              32,
		"panos_predefined_dlp_file_type":            33,
		"panos_predefined_tdb_file_type":            34,
		"panos_predefined_threat":                   35,
		"panos_system_info":                         36,
		"panos_url_filtering_security_profile":      37,
		"panos_url_filtering_security_profiles":     38,
		"panos_user_tag":                            39,
		"panos_vm_auth_key":                         40,
		"panos_vulnerability_security_profile":      41,
		"panos_vulnerability_security_profiles":     42,
		"panos_wildfire_analysis_security_profile":  43,
		"panos_wildfire_analysis_security_profiles": 44,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
