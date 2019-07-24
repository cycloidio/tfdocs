package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "logicmonitor_collector",
			Category:         "Resources",
			ShortDescription: `Provides a LogicMonitor collector resource. This can be used to create and manage LogicMonitor collectors`,
			Description:      ``,
			Keywords: []string{
				"collector",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backup_collector_id",
					Description: `(Optional) The Id of the failover Collector configured for this Collector`,
				},
				resource.Attribute{
					Name:        "collector_group_id",
					Description: `(Optional) The Id of the group the Collector is in`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Collector's description`,
				},
				resource.Attribute{
					Name:        "enable_failback",
					Description: `(Optional) Whether or not automatic failback is enabled for the Collector`,
				},
				resource.Attribute{
					Name:        "enable_collector_device_failover",
					Description: `(Optional) Whether or not the device the Collector is installed on is enabled for fail over`,
				},
				resource.Attribute{
					Name:        "escalation_chain_id",
					Description: `(Optional) The Id of the escalation chain associated with this Collector`,
				},
				resource.Attribute{
					Name:        "resend_interval",
					Description: `(Optional) The interval, in minutes, after which alert notifications for the Collector will be resent`,
				},
				resource.Attribute{
					Name:        "suppress_alert_clear",
					Description: `(Optional) Whether alert clear notifications are suppressed for the Collector`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logicmonitor_collectorgroup",
			Category:         "Resources",
			ShortDescription: `Provides a LogicMonitor collector group resource. This can be used to create and manage LogicMonitor collector groups`,
			Description:      ``,
			Keywords: []string{
				"collectorgroup",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of collector group`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Set description of collector group`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logicmonitor_device",
			Category:         "Resources",
			ShortDescription: `Provides a LogicMonitor device resource. This can be used to create and manage LogicMonitor devices`,
			Description:      ``,
			Keywords: []string{
				"device",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "ip_addr",
					Description: `(Required) Ip Address/Hostname of device`,
				},
				resource.Attribute{
					Name:        "collector",
					Description: `(required) The id of the collector that will monitoring the device`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name of device, (default is ip_addr)`,
				},
				resource.Attribute{
					Name:        "disable_alerting",
					Description: `(Optional) The host is created with alerting disabled (default is true)`,
				},
				resource.Attribute{
					Name:        "hostgroup_id",
					Description: `(Optional) The host group id that specifies which group the device belongs to (multiple host group ids can be added, represented by a comma separated string)`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) The properties associated with this device group. Any string value pair will work (see example). ## Import Devices can be imported using their device id or ip address/dns name ` + "`" + `` + "`" + `` + "`" + ` $ terraform import logicmonitor_device.host 751 $ terraform import logicmonitor_device.host server01.us-east-1.logicmonitor.net ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logicmonitor_devicegroup",
			Category:         "Resources",
			ShortDescription: `Provides a LogicMonitor device group resource. This can be used to create and manage LogicMonitor device groups`,
			Description:      ``,
			Keywords: []string{
				"devicegroup",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of device group`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of device group`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Optional) The id of the parent group for this device group (the root device group has an Id of 1)`,
				},
				resource.Attribute{
					Name:        "applies_to",
					Description: `(Optional) The Applies to custom query for this group. Setting this field will make this a dynamic group.`,
				},
				resource.Attribute{
					Name:        "disable_alerting",
					Description: `(Optional) Indicates whether alerting is disabled (true) or enabled (false) for this device group`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) The properties associated with this device group. Any string value pair will work (see example). ## Import Device Groups can be imported using their group id or full path ` + "`" + `` + "`" + `` + "`" + ` $ terraform import logicmonitor_device_group.group1 451 $ terraform import logicmonitor_device_group.group1 Production/SBA/Linux ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"logicmonitor_collector":      0,
		"logicmonitor_collectorgroup": 1,
		"logicmonitor_device":         2,
		"logicmonitor_devicegroup":    3,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
