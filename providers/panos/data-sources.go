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
					Description: `The device group (default: ` + "`" + `shared` + "`" + `). The following NGFW arguments are supported:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The object's vsys (default: ` + "`" + `vsys1` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "defaults",
					Description: `The application's defaults spec (defined below). This section is omitted for a "defaults" of ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The category.`,
				},
				resource.Attribute{
					Name:        "subcategory",
					Description: `The subcategory.`,
				},
				resource.Attribute{
					Name:        "technology",
					Description: `The technology.`,
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
					Description: `(bool) No appid caching. ` + "`" + `defaults` + "`" + ` supports the following attributes:`,
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
					Description: `The ICMP6 spec (defined below) ` + "`" + `defaults.port` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `List of ports. ` + "`" + `defaults.ip_protocol` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The IP protocol value. ` + "`" + `defaults.icmp` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(int) The type.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(int) The code. ` + "`" + `defaults.icmp6` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(int) The type.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(int) The code. ` + "`" + `timeout_settings` + "`" + ` supports the following attributes:`,
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
					Description: `(int) TCP time wait timeout. ` + "`" + `scanning` + "`" + ` supports the following attributes:`,
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "defaults",
					Description: `The application's defaults spec (defined below). This section is omitted for a "defaults" of ` + "`" + `None` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The category.`,
				},
				resource.Attribute{
					Name:        "subcategory",
					Description: `The subcategory.`,
				},
				resource.Attribute{
					Name:        "technology",
					Description: `The technology.`,
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
					Description: `(bool) No appid caching. ` + "`" + `defaults` + "`" + ` supports the following attributes:`,
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
					Description: `The ICMP6 spec (defined below) ` + "`" + `defaults.port` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `List of ports. ` + "`" + `defaults.ip_protocol` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The IP protocol value. ` + "`" + `defaults.icmp` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(int) The type.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(int) The code. ` + "`" + `defaults.icmp6` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(int) The type.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(int) The code. ` + "`" + `timeout_settings` + "`" + ` supports the following attributes:`,
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
					Description: `(int) TCP time wait timeout. ` + "`" + `scanning` + "`" + ` supports the following attributes:`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_application_objects",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
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
			Type:             "panos_audit_comment_history",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"audit",
				"comment",
				"history",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys. (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `The device group location (default: ` + "`" + `shared` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "rulebase",
					Description: `The rulebase. Valid values are ` + "`" + `pre-rulebase` + "`" + ` (default), ` + "`" + `rulebase` + "`" + `, or ` + "`" + `post-rulebase` + "`" + `. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `The rule type. Valid values are ` + "`" + `security` + "`" + ` (default), ` + "`" + `pbf` + "`" + `, ` + "`" + `nat` + "`" + `, and ` + "`" + `decryption` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The rule's name.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `Valid values are ` + "`" + `backward` + "`" + ` (default) to see newest logs first, or ` + "`" + `forward` + "`" + ` to see oldest first.`,
				},
				resource.Attribute{
					Name:        "nlogs",
					Description: `(int) Number of audit comments to return, max 5000 (default: ` + "`" + `100` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "skip",
					Description: `(int) Specify the number of logs to skip when doing log retrieval. This is useful when retrieving logs in batches to skip previously retrieved logs. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `(repeated) A log entry spec, defined below. Each ` + "`" + `log` + "`" + ` section has the following attributes:`,
				},
				resource.Attribute{
					Name:        "admin",
					Description: `The admin who made the change.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `The audit comment.`,
				},
				resource.Attribute{
					Name:        "config_version",
					Description: `The config version.`,
				},
				resource.Attribute{
					Name:        "time_generated",
					Description: `The time generated, as reported by PAN-OS.`,
				},
				resource.Attribute{
					Name:        "time_generated_rfc3339",
					Description: `An opportunistic representation of ` + "`" + `time_generated` + "`" + ` in RFC3339. This is created by combining the ` + "`" + `time_generated` + "`" + ` with the timezone information of PAN-OS.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "log",
					Description: `(repeated) A log entry spec, defined below. Each ` + "`" + `log` + "`" + ` section has the following attributes:`,
				},
				resource.Attribute{
					Name:        "admin",
					Description: `The admin who made the change.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `The audit comment.`,
				},
				resource.Attribute{
					Name:        "config_version",
					Description: `The config version.`,
				},
				resource.Attribute{
					Name:        "time_generated",
					Description: `The time generated, as reported by PAN-OS.`,
				},
				resource.Attribute{
					Name:        "time_generated_rfc3339",
					Description: `An opportunistic representation of ` + "`" + `time_generated` + "`" + ` in RFC3339. This is created by combining the ` + "`" + `time_generated` + "`" + ` with the timezone information of PAN-OS.`,
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
					Description: `(Required) The name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "username_field",
					Description: `Username field. Valid values are an empty string for ` + "`" + `None` + "`" + `, ` + "`" + `subject` + "`" + `, or ` + "`" + `subject-alt` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username_field_value",
					Description: `The value.`,
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
					Description: `(int) CRL receive timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "ocsp_receive_timeout",
					Description: `(int) OCSP receive timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "certificate_status_timeout",
					Description: `(int) Certificate status timeout in seconds.`,
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
					Description: `(repeated) List of CA certificates, defined below. ` + "`" + `certificate` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name.`,
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "username_field",
					Description: `Username field. Valid values are an empty string for ` + "`" + `None` + "`" + `, ` + "`" + `subject` + "`" + `, or ` + "`" + `subject-alt` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username_field_value",
					Description: `The value.`,
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
					Description: `(int) CRL receive timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "ocsp_receive_timeout",
					Description: `(int) OCSP receive timeout in seconds.`,
				},
				resource.Attribute{
					Name:        "certificate_status_timeout",
					Description: `(int) Certificate status timeout in seconds.`,
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
					Description: `(repeated) List of CA certificates, defined below. ` + "`" + `certificate` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_certificate_profiles",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"certificate",
				"profiles",
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
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). ## Attribute Reference The following attributes are supported:`,
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
			Type:             "panos_custom_url_categories",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"custom",
				"url",
				"categories",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `The device group location (default: ` + "`" + `shared` + "`" + `) ## Attribute Reference The following attributes are supported:`,
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
					Description: `(Optional) The vsys (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
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
					Name:        "sites",
					Description: `(list) The site list.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The custom URL category type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "sites",
					Description: `(list) The site list.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The custom URL category type.`,
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
			Type:             "panos_decryption_rule",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"decryption",
				"rule",
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
					Name:        "name",
					Description: `(Required) The rule name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The rule definition (see below). ` + "`" + `rule` + "`" + ` has the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The rule name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The PAN-OS UUID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `List of source zones.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `List of source addresses.`,
				},
				resource.Attribute{
					Name:        "negate_source",
					Description: `(bool) If the source should be negated.`,
				},
				resource.Attribute{
					Name:        "source_users",
					Description: `List of source users.`,
				},
				resource.Attribute{
					Name:        "destination_zones",
					Description: `List of destination zones.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `List of destination addresses.`,
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
					Description: `List of services.`,
				},
				resource.Attribute{
					Name:        "url_categories",
					Description: `List of URL categories.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action to take.`,
				},
				resource.Attribute{
					Name:        "decryption_type",
					Description: `The decryption type.`,
				},
				resource.Attribute{
					Name:        "ssl_certificate",
					Description: `The SSL certificate.`,
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
					Description: `A listing of vsys to apply this rule to. If ` + "`" + `serial` + "`" + ` is a VM, then this parameter should just be omitted.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rule",
					Description: `The rule definition (see below). ` + "`" + `rule` + "`" + ` has the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The rule name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The PAN-OS UUID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description.`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `List of source zones.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `List of source addresses.`,
				},
				resource.Attribute{
					Name:        "negate_source",
					Description: `(bool) If the source should be negated.`,
				},
				resource.Attribute{
					Name:        "source_users",
					Description: `List of source users.`,
				},
				resource.Attribute{
					Name:        "destination_zones",
					Description: `List of destination zones.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `List of destination addresses.`,
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
					Description: `List of services.`,
				},
				resource.Attribute{
					Name:        "url_categories",
					Description: `List of URL categories.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action to take.`,
				},
				resource.Attribute{
					Name:        "decryption_type",
					Description: `The decryption type.`,
				},
				resource.Attribute{
					Name:        "ssl_certificate",
					Description: `The SSL certificate.`,
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
					Description: `A listing of vsys to apply this rule to. If ` + "`" + `serial` + "`" + ` is a VM, then this parameter should just be omitted.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_decryption_rules",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"decryption",
				"rules",
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
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). ## Attribute Reference The following attributes are supported:`,
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
					Description: `(Required) The device group's name. ## Attribute Reference The following attributes are supported:`,
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
					Description: `The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `A subset of all available vsys on the firewall that should be in this device group.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `A subset of all available vsys on the firewall that should be in this device group.`,
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
			Type:             "panos_device_groups",
			Category:         "Panorama",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"panorama",
				"device",
				"groups",
			},
			Arguments: []resource.Attribute{},
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
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `The device group location (default: ` + "`" + `shared` + "`" + `) The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name. ## Attribute Reference The following attributes are supported:`,
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
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). Panorama:`,
				},
				resource.Attribute{
					Name:        "device_group",
					Description: `The device group location (default: ` + "`" + `shared` + "`" + `) ## Attribute Reference The following attributes are supported:`,
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
					Description: `(Optional) The device group (default: ` + "`" + `shared` + "`" + `). NGFW specific arguments:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The object's name ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of EDL.`,
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
					Description: `How often to retrieve the EDL.`,
				},
				resource.Attribute{
					Name:        "repeat_at",
					Description: `The time at which to retrieve the EDL.`,
				},
				resource.Attribute{
					Name:        "repeat_day_of_week",
					Description: `Repeat day of week.`,
				},
				resource.Attribute{
					Name:        "repeat_day_of_month",
					Description: `(int) If ` + "`" + `repeat` + "`" + ` is ` + "`" + `monthly` + "`" + `, then the repeat day of month.`,
				},
				resource.Attribute{
					Name:        "exceptions",
					Description: `(list) List of exceptions.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `The type of EDL.`,
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
					Description: `How often to retrieve the EDL.`,
				},
				resource.Attribute{
					Name:        "repeat_at",
					Description: `The time at which to retrieve the EDL.`,
				},
				resource.Attribute{
					Name:        "repeat_day_of_week",
					Description: `Repeat day of week.`,
				},
				resource.Attribute{
					Name:        "repeat_day_of_month",
					Description: `(int) If ` + "`" + `repeat` + "`" + ` is ` + "`" + `monthly` + "`" + `, then the repeat day of month.`,
				},
				resource.Attribute{
					Name:        "exceptions",
					Description: `(list) List of exceptions.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_edls",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"edls",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group (default: ` + "`" + `shared` + "`" + `) NGFW:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys (default: ` + "`" + `vsys1` + "`" + `). ## Attribute Reference The following attributes are supported:`,
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
					Description: `(Optional) The vsys (default: ` + "`" + `shared` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `List of users in this group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "users",
					Description: `List of users in this group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_local_user_db_groups",
			Category:         "Device",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"device",
				"local",
				"user",
				"db",
				"groups",
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
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). ## Attribute Reference The following attributes are supported:`,
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
					Name:        "name",
					Description: `(Required) The rule name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The rule definition (see below). Each ` + "`" + `rule` + "`" + ` defined supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The NAT rule's name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(PAN-OS 9.0+) The PAN-OS UUID.`,
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
					Description: `NAT type.`,
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
					Description: `The original packet specification (see below).`,
				},
				resource.Attribute{
					Name:        "translated_packet",
					Description: `The translated packet spec (see below). ` + "`" + `original_packet` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `The list of source zone(s).`,
				},
				resource.Attribute{
					Name:        "destination_zone",
					Description: `The destination zone.`,
				},
				resource.Attribute{
					Name:        "destination_interface",
					Description: `Egress interface from route lookup.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `List of source address(es).`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `List of destination address(es). ` + "`" + `translated_packet` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The source spec (see below).`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `The destination spec (see below). ` + "`" + `translated_packet.source` + "`" + ` supports the following attributes:`,
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
					Description: `Static IP source translation spec (see below). ` + "`" + `translated_packet.source.dynamic_ip_and_port` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `Translated address source translation type spec (see below).`,
				},
				resource.Attribute{
					Name:        "interface_address",
					Description: `Interface address source translation type spec (see below). ` + "`" + `translated_packet.source.dynamic_ip_and_port.translated_address` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `List of translated addresses. ` + "`" + `translated_packet.source.dynamic_ip_and_port.interface_address` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `The interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address. ` + "`" + `translated_packet.source.dynamic_ip` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `The list of translated addresses.`,
				},
				resource.Attribute{
					Name:        "fallback",
					Description: `The fallback spec (see below). ` + "`" + `translated_packet.source.dynamic_ip.fallback` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `The translated address fallback spec (see below).`,
				},
				resource.Attribute{
					Name:        "interface_address",
					Description: `The interface address fallback spec (see below). ` + "`" + `translated_packet.source.dynamic_ip.fallback.translated_address` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `List of source address translation fallback translated addresses. ` + "`" + `translated_packet.source.dynamic_ip.fallback.interface_address` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Source address translation fallback interface.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of interface fallback.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address of the fallback interface. ` + "`" + `translated_packet.source.static_ip` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `The statically translated source address.`,
				},
				resource.Attribute{
					Name:        "bi_directional",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to enable bi-directional source address translation. ` + "`" + `translated_packet.destination` + "`" + ` supports the following attributes:`,
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
					Description: `Destination address translation address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(int) Destination address translation port number. ` + "`" + `translated_packet.destination.dynamic` + "`" + ` and ` + "`" + `translated_packet.destination.dynamic_translation` + "`" + ` support the following attributes:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Destination address translation address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(int) Destination address translation port number.`,
				},
				resource.Attribute{
					Name:        "distribution",
					Description: `(PAN-OS 8.1+) Distribution algorithm for destination address pool. The PAN-OS 8.1 GUI doesn't seem to set this anywhere, but this is added here for completeness' sake. The GUI sets this to ` + "`" + `round-robin` + "`" + ` currently.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rule",
					Description: `The rule definition (see below). Each ` + "`" + `rule` + "`" + ` defined supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The NAT rule's name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(PAN-OS 9.0+) The PAN-OS UUID.`,
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
					Description: `NAT type.`,
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
					Description: `The original packet specification (see below).`,
				},
				resource.Attribute{
					Name:        "translated_packet",
					Description: `The translated packet spec (see below). ` + "`" + `original_packet` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "source_zones",
					Description: `The list of source zone(s).`,
				},
				resource.Attribute{
					Name:        "destination_zone",
					Description: `The destination zone.`,
				},
				resource.Attribute{
					Name:        "destination_interface",
					Description: `Egress interface from route lookup.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `List of source address(es).`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `List of destination address(es). ` + "`" + `translated_packet` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `The source spec (see below).`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `The destination spec (see below). ` + "`" + `translated_packet.source` + "`" + ` supports the following attributes:`,
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
					Description: `Static IP source translation spec (see below). ` + "`" + `translated_packet.source.dynamic_ip_and_port` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `Translated address source translation type spec (see below).`,
				},
				resource.Attribute{
					Name:        "interface_address",
					Description: `Interface address source translation type spec (see below). ` + "`" + `translated_packet.source.dynamic_ip_and_port.translated_address` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `List of translated addresses. ` + "`" + `translated_packet.source.dynamic_ip_and_port.interface_address` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `The interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address. ` + "`" + `translated_packet.source.dynamic_ip` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `The list of translated addresses.`,
				},
				resource.Attribute{
					Name:        "fallback",
					Description: `The fallback spec (see below). ` + "`" + `translated_packet.source.dynamic_ip.fallback` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `The translated address fallback spec (see below).`,
				},
				resource.Attribute{
					Name:        "interface_address",
					Description: `The interface address fallback spec (see below). ` + "`" + `translated_packet.source.dynamic_ip.fallback.translated_address` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "translated_addresses",
					Description: `List of source address translation fallback translated addresses. ` + "`" + `translated_packet.source.dynamic_ip.fallback.interface_address` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Source address translation fallback interface.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of interface fallback.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address of the fallback interface. ` + "`" + `translated_packet.source.static_ip` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "translated_address",
					Description: `The statically translated source address.`,
				},
				resource.Attribute{
					Name:        "bi_directional",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to enable bi-directional source address translation. ` + "`" + `translated_packet.destination` + "`" + ` supports the following attributes:`,
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
					Description: `Destination address translation address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(int) Destination address translation port number. ` + "`" + `translated_packet.destination.dynamic` + "`" + ` and ` + "`" + `translated_packet.destination.dynamic_translation` + "`" + ` support the following attributes:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Destination address translation address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(int) Destination address translation port number.`,
				},
				resource.Attribute{
					Name:        "distribution",
					Description: `(PAN-OS 8.1+) Distribution algorithm for destination address pool. The PAN-OS 8.1 GUI doesn't seem to set this anywhere, but this is added here for completeness' sake. The GUI sets this to ` + "`" + `round-robin` + "`" + ` currently.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_nat_rules",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"nat",
				"rules",
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
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). ## Attribute Reference The following attributes are supported:`,
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
			Type:             "panos_pbf_rule",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"pbf",
				"rule",
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
					Name:        "name",
					Description: `(Required) The rule name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The rule definition (see below). The ` + "`" + `rule` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The rule name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(PAN-OS 9.0+) The PAN-OS UUID.`,
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
					Description: `The source spec (defined below).`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `The destination spec (defined below).`,
				},
				resource.Attribute{
					Name:        "forwarding",
					Description: `The forwarding spec (defined below). ` + "`" + `rule.source` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `List of source zones.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `List of source interfaces.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `List of source IP addresses.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `List of source users.`,
				},
				resource.Attribute{
					Name:        "negate",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to negate the source. ` + "`" + `rule.destination` + "`" + ` supports the following attributes:`,
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
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to negate the destination. ` + "`" + `rule.forwarding` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action to take.`,
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
					Description: `If ` + "`" + `action=forward` + "`" + `, the next hop type.`,
				},
				resource.Attribute{
					Name:        "next_hop_value",
					Description: `If ` + "`" + `action=forward` + "`" + ` and ` + "`" + `next_hop_type` + "`" + ` is defined, then the next hop address.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `The monitor spec (defined below) if monitoring is enabled.`,
				},
				resource.Attribute{
					Name:        "symmetric_return",
					Description: `The symmetric return spec (defined below) if it's enforced. ` + "`" + `rule.forwarding.monitor` + "`" + ` supports the following attributes:`,
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
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to disable this rule if nexthop/monitor IP is unreachable. ` + "`" + `rule.forwarding.symmetric_return` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to enforce symmetric return.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `List of next hop addresses.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rule",
					Description: `The rule definition (see below). The ` + "`" + `rule` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The rule name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(PAN-OS 9.0+) The PAN-OS UUID.`,
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
					Description: `The source spec (defined below).`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `The destination spec (defined below).`,
				},
				resource.Attribute{
					Name:        "forwarding",
					Description: `The forwarding spec (defined below). ` + "`" + `rule.source` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `List of source zones.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `List of source interfaces.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `List of source IP addresses.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `List of source users.`,
				},
				resource.Attribute{
					Name:        "negate",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to negate the source. ` + "`" + `rule.destination` + "`" + ` supports the following attributes:`,
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
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to negate the destination. ` + "`" + `rule.forwarding` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `The action to take.`,
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
					Description: `If ` + "`" + `action=forward` + "`" + `, the next hop type.`,
				},
				resource.Attribute{
					Name:        "next_hop_value",
					Description: `If ` + "`" + `action=forward` + "`" + ` and ` + "`" + `next_hop_type` + "`" + ` is defined, then the next hop address.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `The monitor spec (defined below) if monitoring is enabled.`,
				},
				resource.Attribute{
					Name:        "symmetric_return",
					Description: `The symmetric return spec (defined below) if it's enforced. ` + "`" + `rule.forwarding.monitor` + "`" + ` supports the following attributes:`,
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
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to disable this rule if nexthop/monitor IP is unreachable. ` + "`" + `rule.forwarding.symmetric_return` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(bool) Set to ` + "`" + `true` + "`" + ` to enforce symmetric return.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `List of next hop addresses.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_pbf_rules",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"pbf",
				"rules",
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
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). ## Attribute Reference The following attributes are supported:`,
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
					Description: `(Optional) The device group (default: ` + "`" + `shared` + "`" + `) NGFW:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys (default: ` + "`" + `vsys1` + "`" + `). The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name. ## Attribute Reference The following attributes are supported:`,
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
			Attributes: []resource.Attribute{
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_security_profile_groups",
			Category:         "Objects",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"objects",
				"security",
				"profile",
				"groups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_group",
					Description: `(Optional) The device group location (default: ` + "`" + `shared` + "`" + `) NGFW:`,
				},
				resource.Attribute{
					Name:        "vsys",
					Description: `(Optional) The vsys (default: ` + "`" + `vsys1` + "`" + `). ## Attribute Reference The following attributes are supported:`,
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
			Type:             "panos_security_rule",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"security",
				"rule",
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
					Name:        "name",
					Description: `(Required) The rule name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The security rule definition (see below). ` + "`" + `rule` + "`" + ` has the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The security rule name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The PAN-OS UUID.`,
				},
				resource.Attribute{
					Name:        "group_tag",
					Description: `The group tag.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Rule type.`,
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
					Description: `List of source zones.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `List of source addresses.`,
				},
				resource.Attribute{
					Name:        "negate_source",
					Description: `(bool) If the source should be negated.`,
				},
				resource.Attribute{
					Name:        "source_users",
					Description: `List of source users.`,
				},
				resource.Attribute{
					Name:        "hip_profiles",
					Description: `List of HIP profiles.`,
				},
				resource.Attribute{
					Name:        "destination_zones",
					Description: `List of destination zones.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `List of destination addresses.`,
				},
				resource.Attribute{
					Name:        "negate_destination",
					Description: `(bool) If the destination should be negated.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `List of applications.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `List of services.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `List of categories.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action for the matched traffic.`,
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
					Description: `(bool) Log the end of the traffic flow.`,
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
					Description: `(bool) ICMP unreachable.`,
				},
				resource.Attribute{
					Name:        "disable_server_response_inspection",
					Description: `(bool) If server response inspection is disabled.`,
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
					Description: `(bool, Panorama only) Instead of applying the rule for the given serial numbers, apply it to everything except them. ` + "`" + `rule.target` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `A listing of vsys to apply this rule to. If ` + "`" + `serial` + "`" + ` is a VM, then this parameter should just be omitted.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rule",
					Description: `The security rule definition (see below). ` + "`" + `rule` + "`" + ` has the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The security rule name.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The PAN-OS UUID.`,
				},
				resource.Attribute{
					Name:        "group_tag",
					Description: `The group tag.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Rule type.`,
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
					Description: `List of source zones.`,
				},
				resource.Attribute{
					Name:        "source_addresses",
					Description: `List of source addresses.`,
				},
				resource.Attribute{
					Name:        "negate_source",
					Description: `(bool) If the source should be negated.`,
				},
				resource.Attribute{
					Name:        "source_users",
					Description: `List of source users.`,
				},
				resource.Attribute{
					Name:        "hip_profiles",
					Description: `List of HIP profiles.`,
				},
				resource.Attribute{
					Name:        "destination_zones",
					Description: `List of destination zones.`,
				},
				resource.Attribute{
					Name:        "destination_addresses",
					Description: `List of destination addresses.`,
				},
				resource.Attribute{
					Name:        "negate_destination",
					Description: `(bool) If the destination should be negated.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `List of applications.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `List of services.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `List of categories.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action for the matched traffic.`,
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
					Description: `(bool) Log the end of the traffic flow.`,
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
					Description: `(bool) ICMP unreachable.`,
				},
				resource.Attribute{
					Name:        "disable_server_response_inspection",
					Description: `(bool) If server response inspection is disabled.`,
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
					Description: `(bool, Panorama only) Instead of applying the rule for the given serial numbers, apply it to everything except them. ` + "`" + `rule.target` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Required) The serial number of the firewall.`,
				},
				resource.Attribute{
					Name:        "vsys_list",
					Description: `A listing of vsys to apply this rule to. If ` + "`" + `serial` + "`" + ` is a VM, then this parameter should just be omitted.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_security_rules",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"security",
				"rules",
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
					Description: `The vsys (default: ` + "`" + `vsys1` + "`" + `). ## Attribute Reference The following attributes are supported:`,
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
					Description: `The vsys (default: ` + "`" + `shared` + "`" + `). ## Attribute Reference The following attributes are supported:`,
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
					Description: `(repeatable) List of SSL decrypt exclude certificates specs (specified below). ` + "`" + `ssl_decrypt_exclude_certificate` + "`" + ` sections support the following attributes:`,
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
					Name:        "exclude",
					Description: `(bool) Exclude or not.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `(repeatable) List of SSL decrypt exclude certificates specs (specified below). ` + "`" + `ssl_decrypt_exclude_certificate` + "`" + ` sections support the following attributes:`,
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
					Name:        "exclude",
					Description: `(bool) Exclude or not.`,
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
			Type:             "panos_tech_support_file",
			Category:         "Operational State",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"operational",
				"state",
				"tech",
				"support",
				"file",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "timeout",
					Description: `(int) Timeout for retrieving the tech support file in seconds (default: ` + "`" + `600` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "save_to_local_file_system",
					Description: `(bool) Save the tech support file to the local file system where Terraform is running.`,
				},
				resource.Attribute{
					Name:        "file_system_path",
					Description: `When ` + "`" + `save_to_local_file_system=true` + "`" + `, the file system path to place the tech support file.`,
				},
				resource.Attribute{
					Name:        "save_to_state",
					Description: `(bool) Save the tech support file to the state. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `The tech support file filename.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `For ` + "`" + `save_to_state=true` + "`" + `, the content of the .tgz tech support file.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "filename",
					Description: `The tech support file filename.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `For ` + "`" + `save_to_state=true` + "`" + `, the content of the .tgz tech support file.`,
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
					Description: `The template stack. The following arguments are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The virtual router's name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `List of interfaces.`,
				},
				resource.Attribute{
					Name:        "static_dist",
					Description: `(int) Admin distance - Static.`,
				},
				resource.Attribute{
					Name:        "static_ipv6_dist",
					Description: `(int) Admin distance - Static IPv6.`,
				},
				resource.Attribute{
					Name:        "ospf_int_dist",
					Description: `(int) Admin distance - OSPF Int.`,
				},
				resource.Attribute{
					Name:        "ospf_ext_dist",
					Description: `(int) Admin distance - OSPF Ext.`,
				},
				resource.Attribute{
					Name:        "ospfv3_int_dist",
					Description: `(int) Admin distance - OSPFv3 Int.`,
				},
				resource.Attribute{
					Name:        "ospfv3_ext_dist",
					Description: `(int) Admin distance - OSPFv3 Ext.`,
				},
				resource.Attribute{
					Name:        "ibgp_dist",
					Description: `(int) Admin distance - IBGP.`,
				},
				resource.Attribute{
					Name:        "ebgp_dist",
					Description: `(int) Admin distance - EBGP.`,
				},
				resource.Attribute{
					Name:        "rip_dist",
					Description: `(int) Admin distance - RIP.`,
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
					Description: `Load balancing algorithm.`,
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "interfaces",
					Description: `List of interfaces.`,
				},
				resource.Attribute{
					Name:        "static_dist",
					Description: `(int) Admin distance - Static.`,
				},
				resource.Attribute{
					Name:        "static_ipv6_dist",
					Description: `(int) Admin distance - Static IPv6.`,
				},
				resource.Attribute{
					Name:        "ospf_int_dist",
					Description: `(int) Admin distance - OSPF Int.`,
				},
				resource.Attribute{
					Name:        "ospf_ext_dist",
					Description: `(int) Admin distance - OSPF Ext.`,
				},
				resource.Attribute{
					Name:        "ospfv3_int_dist",
					Description: `(int) Admin distance - OSPFv3 Int.`,
				},
				resource.Attribute{
					Name:        "ospfv3_ext_dist",
					Description: `(int) Admin distance - OSPFv3 Ext.`,
				},
				resource.Attribute{
					Name:        "ibgp_dist",
					Description: `(int) Admin distance - IBGP.`,
				},
				resource.Attribute{
					Name:        "ebgp_dist",
					Description: `(int) Admin distance - EBGP.`,
				},
				resource.Attribute{
					Name:        "rip_dist",
					Description: `(int) Admin distance - RIP.`,
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
					Description: `Load balancing algorithm.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_virtual_routers",
			Category:         "Network",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"virtual",
				"routers",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "template",
					Description: `The template.`,
				},
				resource.Attribute{
					Name:        "template_stack",
					Description: `The template stack. ## Attribute Reference The following attributes are supported:`,
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
		"panos_application_object":                  7,
		"panos_application_objects":                 8,
		"panos_arp":                                 9,
		"panos_arps":                                10,
		"panos_audit_comment_history":               11,
		"panos_certificate_profile":                 12,
		"panos_certificate_profiles":                13,
		"panos_custom_data_pattern_object":          14,
		"panos_custom_data_pattern_objects":         15,
		"panos_custom_url_categories":               16,
		"panos_custom_url_category":                 17,
		"panos_data_filtering_security_profile":     18,
		"panos_data_filtering_security_profiles":    19,
		"panos_decryption_rule":                     20,
		"panos_decryption_rules":                    21,
		"panos_device_group":                        22,
		"panos_device_group_parent":                 23,
		"panos_device_groups":                       24,
		"panos_dhcp_interface_info":                 25,
		"panos_dos_protection_profile":              26,
		"panos_dos_protection_profiles":             27,
		"panos_dynamic_user_group":                  28,
		"panos_dynamic_user_groups":                 29,
		"panos_edl":                                 30,
		"panos_edls":                                31,
		"panos_file_blocking_security_profile":      32,
		"panos_file_blocking_security_profiles":     33,
		"panos_ip_tag":                              34,
		"panos_local_user_db_group":                 35,
		"panos_local_user_db_groups":                36,
		"panos_nat_rule":                            37,
		"panos_nat_rules":                           38,
		"panos_ospf":                                39,
		"panos_ospf_area":                           40,
		"panos_ospf_area_interface":                 41,
		"panos_ospf_area_interfaces":                42,
		"panos_ospf_area_virtual_link":              43,
		"panos_ospf_area_virtual_links":             44,
		"panos_ospf_areas":                          45,
		"panos_ospf_auth_profiles":                  46,
		"panos_ospf_export":                         47,
		"panos_ospf_exports":                        48,
		"panos_pbf_rule":                            49,
		"panos_pbf_rules":                           50,
		"panos_plugin":                              51,
		"panos_predefined_dlp_file_type":            52,
		"panos_predefined_tdb_file_type":            53,
		"panos_predefined_threat":                   54,
		"panos_security_profile_group":              55,
		"panos_security_profile_groups":             56,
		"panos_security_rule":                       57,
		"panos_security_rules":                      58,
		"panos_ssl_decrypt":                         59,
		"panos_system_info":                         60,
		"panos_tech_support_file":                   61,
		"panos_url_filtering_security_profile":      62,
		"panos_url_filtering_security_profiles":     63,
		"panos_user_tag":                            64,
		"panos_virtual_router":                      65,
		"panos_virtual_routers":                     66,
		"panos_vm_auth_key":                         67,
		"panos_vulnerability_security_profile":      68,
		"panos_vulnerability_security_profiles":     69,
		"panos_wildfire_analysis_security_profile":  70,
		"panos_wildfire_analysis_security_profiles": 71,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
