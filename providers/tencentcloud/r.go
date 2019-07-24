package tencentcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_alb_server_attachment",
			Category:         "LB Resources",
			ShortDescription: `Provides an tencentcloud application load balancer servers attachment as a resource, to attach and detach instances from load balancer.`,
			Description:      ``,
			Keywords: []string{
				"lb",
				"alb",
				"server",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required, Forces new resource) loadbalancer ID.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, Forces new resource) listener ID.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `(Optional) location ID only support for layer 7 loadbalancer`,
				},
				resource.Attribute{
					Name:        "backends",
					Description: `(Required) list of backend server. Valid value range [1-100]. ### Block backends The backends mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) A list backend instance ID (CVM instance ID).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port used by the backend server. Valid value range: [1-65535].`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight of the backend server. Valid value range: [0-100]. Default to 10. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `loadbalancer ID.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `listener ID.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `location ID (only support for layer 7 loadbalancer)`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `http or tcp`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `loadbalancer ID.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `listener ID.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `location ID (only support for layer 7 loadbalancer)`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `http or tcp`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_attachment",
			Category:         "AS Resources",
			ShortDescription: `Provides a resource to attach or detach CVM instances to a specified scaling group.`,
			Description:      ``,
			Keywords: []string{
				"as",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_ids",
					Description: `(Required) ID list of CVM instances to be attached to the scaling group.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) ID of a scaling group.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_lifecycle_hook",
			Category:         "AS Resources",
			ShortDescription: `Provides a resource for an AS (Auto scaling) lifecycle hook.`,
			Description:      ``,
			Keywords: []string{
				"as",
				"lifecycle",
				"hook",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "lifecycle_hook_name",
					Description: `(Required) The name of the lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "lifecycle_transition",
					Description: `(Required) The instance state to which you want to attach the lifecycle hook. The valid values are INSTANCE_LAUNCHING and INSTANCE_TERMINATING.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) ID of a scaling group.`,
				},
				resource.Attribute{
					Name:        "default_result",
					Description: `(Optional) Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON. The default value is CONTINUE.`,
				},
				resource.Attribute{
					Name:        "heartbeat_timeout",
					Description: `(Optional) Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. The range is 30 to 3600, and default value is 300.`,
				},
				resource.Attribute{
					Name:        "notification_metadata",
					Description: `(Optional) Contains additional information that you want to include any time AS sends a message to the notification target.`,
				},
				resource.Attribute{
					Name:        "notification_queue_name",
					Description: `(Optional) For CMQ_QUEUE type, a name of queue must be set.`,
				},
				resource.Attribute{
					Name:        "notification_target_type",
					Description: `(Optional) Target type, which can be CMQ_QUEUE or CMQ_TOPIC.`,
				},
				resource.Attribute{
					Name:        "notification_topic_name",
					Description: `(Optional) For CMQ_TOPIC type, a name of topic must be set.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_notification",
			Category:         "AS Resources",
			ShortDescription: `Provides a resource for an AS (Auto scaling) notification.`,
			Description:      ``,
			Keywords: []string{
				"as",
				"notification",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "notification_types",
					Description: `(Required) A list of Notification Types that trigger notifications. Acceptable values are SCALE_OUT_FAILED, SCALE_IN_SUCCESSFUL, SCALE_IN_FAILED, REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL and REPLACE_UNHEALTHY_INSTANCE_FAILED.`,
				},
				resource.Attribute{
					Name:        "notification_user_group_ids",
					Description: `(Required) A group of user IDs to be notified.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) ID of a scaling group.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_config",
			Category:         "AS Resources",
			ShortDescription: `Provides a resource to create a configuration for an AS (Auto scaling) instance.`,
			Description:      ``,
			Keywords: []string{
				"as",
				"scaling",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "configuration_name",
					Description: `(Required) Name of a launch configuration.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) An available image ID for a cvm instance.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `(Required) Specified types of CVM instances.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "enhanced_monitor_service",
					Description: `(Optional) To specify whether to enable cloud monitor service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "enhanced_security_service",
					Description: `(Optional) To specify whether to enable cloud security service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "instance_tags",
					Description: `(Optional) A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional) Charge types for network traffic. Available values include TRAFFIC_POSTPAID_BY_HOUR.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional) Max bandwidth of Internet access in Mbps. Default is 0.`,
				},
				resource.Attribute{
					Name:        "keep_image_login",
					Description: `(Optional) Specify whether to keep original settings of a CVM image. And it can't be used with password or key_ids together.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `(Optional) ID list of keys.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password to access.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Specifys to which project the configuration belongs.`,
				},
				resource.Attribute{
					Name:        "public_ip_assigned",
					Description: `(Optional) Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) Security groups to which a CVM instance belongs.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Volume of system disk in GB. Default is 50.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional) Type of a CVM disk, and available values include CLOUD_PREMIUM and CLOUD_SSD. Default is CLOUD_PREMIUM`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) ase64-encoded User Data text, the length limit is 16KB. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) Volume of disk in GB. Default is 0.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional) Types of disk，available values: CLOUD_PREMIUM and CLOUD_SSD.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) Data disk snapshot ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the launch configuration was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current statues of a launch configuration. ## Import AutoScaling Configuration can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_as_scaling_config.scaling_config asc-n32ymck2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the launch configuration was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current statues of a launch configuration. ## Import AutoScaling Configuration can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_as_scaling_config.scaling_config asc-n32ymck2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_group",
			Category:         "AS Resources",
			ShortDescription: `Provides a resource to create a group of AS (Auto scaling) instances.`,
			Description:      ``,
			Keywords: []string{
				"as",
				"scaling",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "configuration_id",
					Description: `(Required) An available ID for a launch configuration.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required) Maximum number of CVM instances (0~2000).`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Required) Minimum number of CVM instances (0~2000).`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `(Required) Name of a scaling group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) ID of VPC network.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `(Optional) Default cooldown time in second, and default value is 300.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Optional) Desired volume of CVM instances, which is between max_size and min_size.`,
				},
				resource.Attribute{
					Name:        "forward_balancer_ids",
					Description: `(Optional) List of application load balancers, which can't be specified with load_balancer_ids together.`,
				},
				resource.Attribute{
					Name:        "load_balancer_ids",
					Description: `(Optional) ID list of traditional load balancers.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Specifys to which project the scaling group belongs.`,
				},
				resource.Attribute{
					Name:        "retry_policy",
					Description: `(Optional) Available values for retry policies include IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) ID list of subnet, and for VPC it is required.`,
				},
				resource.Attribute{
					Name:        "termination_policies",
					Description: `(Optional) Available values for termination policies include OLDEST_INSTANCE and NEWEST_INSTANCE.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `(Optional) List of available zones, for Basic network it is required. The ` + "`" + `forward_balancer_ids` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) Listener ID for application load balancers.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) ID of available load balancers.`,
				},
				resource.Attribute{
					Name:        "target_attribute",
					Description: `(Required) Attribute list of target rules.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `(Optional) ID of forwarding rules. The ` + "`" + `target_attribute` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port number.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Weight. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `The time when the AS group was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of a scaling group. ## Import AutoScaling Groups can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_as_scaling_group.scaling_group asg-n32ymck2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_count",
					Description: `The time when the AS group was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of a scaling group. ## Import AutoScaling Groups can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_as_scaling_group.scaling_group asg-n32ymck2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_policy",
			Category:         "AS Resources",
			ShortDescription: `Provides a resource for an AS (Auto scaling) policy.`,
			Description:      ``,
			Keywords: []string{
				"as",
				"scaling",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "adjustment_type",
					Description: `(Required) Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Available values include CHANGE_IN_CAPACITY, EXACT_CAPACITY and PERCENT_CHANGE_IN_CAPACITY.`,
				},
				resource.Attribute{
					Name:        "adjustment_value",
					Description: `(Required) Define the number of instances by which to scale.For CHANGE_IN_CAPACITY type or PERCENT_CHANGE_IN_CAPACITY, a positive increment adds to the current capacity and a negative value removes from the current capacity. For EXACT_CAPACITY type, it defines an absolute number of the existing Auto Scaling group size.`,
				},
				resource.Attribute{
					Name:        "comparison_operator",
					Description: `(Required) Comparison operator, of which valid values can be GREATER_THAN, GREATER_THAN_OR_EQUAL_TO, LESS_THAN, LESS_THAN_OR_EQUAL_TO, EQUAL_TO and NOT_EQUAL_TO.`,
				},
				resource.Attribute{
					Name:        "continuous_time",
					Description: `(Required) Retry times (1~10).`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) Name of an indicator, which can be CPU_UTILIZATION, MEM_UTILIZATION, LAN_TRAFFIC_OUT, LAN_TRAFFIC_IN, WAN_TRAFFIC_OUT and WAN_TRAFFIC_IN.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Required) Time period in second, of which valid values can be 60 and 300.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required) Name of a policy used to define a reaction when an alarm is triggered.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) ID of a scaling group.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) Alarm threshold.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `(Optional) Cooldwon time in second. Default is 300.`,
				},
				resource.Attribute{
					Name:        "notification_user_group_ids",
					Description: `(Optional) An ID group of users to be notified when an alarm is triggered.`,
				},
				resource.Attribute{
					Name:        "statistic",
					Description: `(Optional) Statistic types, include AVERAGE, MAXIMUM and MINIMUM. Default is AVERAGE.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_schedule",
			Category:         "AS Resources",
			ShortDescription: `Provides a resource for an AS (Auto scaling) schedule.`,
			Description:      ``,
			Keywords: []string{
				"as",
				"schedule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Required) The desired number of CVM instances that should be running in the group.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required) The maximum size for the Auto Scaling group.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Required) The minimum size for the Auto Scaling group.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required, ForceNew) ID of a scaling group.`,
				},
				resource.Attribute{
					Name:        "schedule_action_name",
					Description: `(Required) The name of this scaling action.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) The time for this action to start, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) The time for this action to end, in "YYYY-MM-DDThh:mm:ss+08:00" format (UTC+8).`,
				},
				resource.Attribute{
					Name:        "recurrence",
					Description: `(Optional) The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with end_time together.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_snapshot",
			Category:         "CBS Resources",
			ShortDescription: `Provide a resource to create a CBS snapshot.`,
			Description:      ``,
			Keywords: []string{
				"cbs",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `(Required) Name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `(Required, ForceNew) ID of the the CBS which this snapshot created from. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of snapshot.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Types of CBS which this snapshot created from.`,
				},
				resource.Attribute{
					Name:        "percent",
					Description: `Snapshot creation progress percentage. If the snapshot has created successfully, the constant value is 100.`,
				},
				resource.Attribute{
					Name:        "snapshot_status",
					Description: `Status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Volume of storage which this snapshot created from. ## Import CBS snapshot can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_snapshot.snapshot snap-3sa3f39b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of snapshot.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Types of CBS which this snapshot created from.`,
				},
				resource.Attribute{
					Name:        "percent",
					Description: `Snapshot creation progress percentage. If the snapshot has created successfully, the constant value is 100.`,
				},
				resource.Attribute{
					Name:        "snapshot_status",
					Description: `Status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Volume of storage which this snapshot created from. ## Import CBS snapshot can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_snapshot.snapshot snap-3sa3f39b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_snapshot_policy",
			Category:         "CBS Resources",
			ShortDescription: `Provides a snapshot policy resource.`,
			Description:      ``,
			Keywords: []string{
				"cbs",
				"snapshot",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repeat_hours",
					Description: `(Required) Trigger times of periodic snapshot, the available values are 0 to 23. The 0 means 00:00, and so on.`,
				},
				resource.Attribute{
					Name:        "repeat_weekdays",
					Description: `(Required) Periodic snapshot is enabled, the available values are [0, 1, 2, 3, 4, 5, 6]. 0 means Sunday, 1-6 means Monday to Saturday.`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_name",
					Description: `(Required) Name of snapshot policy. The maximum length can not exceed 60 bytes.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `(Optional) Retention days of the snapshot, and the default value is 7. ## Import CBS snapshot policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_snapshot_policy.snapshot_policy asp-jliex1tn ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_storage",
			Category:         "CBS Resources",
			ShortDescription: `Provide a resource to create a CBS.`,
			Description:      ``,
			Keywords: []string{
				"cbs",
				"storage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) The available zone that the CBS instance locates at.`,
				},
				resource.Attribute{
					Name:        "storage_name",
					Description: `(Required) Name of CBS. The maximum length can not exceed 60 bytes.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `(Required) Volume of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Required, ForceNew) Type of CBS medium, and available values include CLOUD_BASIC, CLOUD_PREMIUM and CLOUD_SSD.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `(Optional, ForceNew) Indicates whether CBS is encrypted.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The purchased usage period of CBS, and value range [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36].`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project to which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) ID of the snapshot. If specified, created the CBS by this snapshot.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The available tags within this CBS. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "attached",
					Description: `Indicates whether the CBS is mounted the CVM.`,
				},
				resource.Attribute{
					Name:        "storage_status",
					Description: `Status of CBS, and available values include UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING. ## Import CBS storage can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_storage.storage disk-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "attached",
					Description: `Indicates whether the CBS is mounted the CVM.`,
				},
				resource.Attribute{
					Name:        "storage_status",
					Description: `Status of CBS, and available values include UNATTACHED, ATTACHING, ATTACHED, DETACHING, EXPANDING, ROLLBACKING, TORECYCLE and DUMPING. ## Import CBS storage can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_storage.storage disk-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_storage_attachment",
			Category:         "CBS Resources",
			ShortDescription: `Provides a CBS storage attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"cbs",
				"storage",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the CVM instance.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `(Required, ForceNew) ID of the mounted CBS.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ccn",
			Category:         "CCN Resources",
			ShortDescription: `Provides a resource to create a CCN instance.`,
			Description:      ``,
			Keywords: []string{
				"ccn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the CCN to be queried, and maximum length does not exceed 60 bytes.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of CCN, and maximum length does not exceed 100 bytes.`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `(Optional, ForceNew) Service quality of CCN, and the available value include 'PT', 'AU', 'AG'. The default is 'AU'. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of attached instances.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance. The available value include 'ISOLATED'(arrears) and 'AVAILABLE'. ## Import Ccn instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_ccn.test ccn-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of attached instances.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance. The available value include 'ISOLATED'(arrears) and 'AVAILABLE'. ## Import Ccn instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_ccn.test ccn-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ccn_attachment",
			Category:         "CCN Resources",
			ShortDescription: `Provides a CCN attaching resource.`,
			Description:      ``,
			Keywords: []string{
				"ccn",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ccn_id",
					Description: `(Required, ForceNew) ID of the CCN`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of instance is attached.`,
				},
				resource.Attribute{
					Name:        "instance_region",
					Description: `(Required, ForceNew) The region that the instance locates at.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) Type of attached instance network, and available values include VPC, DIRECTCONNECT and BMVPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "attached_time",
					Description: `Time of attaching.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of the instance that is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance is attached, and available values include PENDING, ACTIVE, EXPIRED, REJECTED, DELETED, FAILED(asynchronous forced disassociation after 2 hours), ATTACHING, DETACHING and DETACHFAILED(asynchronous forced disassociation after 2 hours).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "attached_time",
					Description: `Time of attaching.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of the instance that is attached.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance is attached, and available values include PENDING, ACTIVE, EXPIRED, REJECTED, DELETED, FAILED(asynchronous forced disassociation after 2 hours), ATTACHING, DETACHING and DETACHFAILED(asynchronous forced disassociation after 2 hours).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ccn_bandwidth_limit",
			Category:         "CCN Resources",
			ShortDescription: `Provides a resource to limit CCN bandwidth.`,
			Description:      ``,
			Keywords: []string{
				"ccn",
				"bandwidth",
				"limit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ccn_id",
					Description: `(Required, ForceNew) ID of the CCN`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required, ForceNew) Limitation of region.`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit",
					Description: `(Optional) Limitation of bandwidth.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_container_cluster",
			Category:         "Container Cluster Resources",
			ShortDescription: `Provides a TencentCloud Container Cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the cluster.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Required) The cpu of the node.`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `(Required) The memory of the node.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `(Required) The system os name of the node.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The network bandwidth of the node.`,
				},
				resource.Attribute{
					Name:        "bandwidth_type",
					Description: `(Required) The network type of the node.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The subnet id which the node stays in.`,
				},
				resource.Attribute{
					Name:        "is_vpc_gateway",
					Description: `(Required) Describe whether the node enable the gateway capability.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `(Required) The size of the data volumn.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) The type of the data volumn. see more from CVM.`,
				},
				resource.Attribute{
					Name:        "root_size",
					Description: `(Required) The size of the root volumn.`,
				},
				resource.Attribute{
					Name:        "root_type",
					Description: `(Optional) The type of the root volumn. see more from CVM.`,
				},
				resource.Attribute{
					Name:        "goods_num",
					Description: `(Required) The node number is going to create in the cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specify vpc which the node(s) stay in.`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `(Required) The CIDR which the cluster is going to use.`,
				},
				resource.Attribute{
					Name:        "cluster_desc",
					Description: `(Optional) The description of the cluster.`,
				},
				resource.Attribute{
					Name:        "cvm_type",
					Description: `(Optional) The type of node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The puchase duration of the node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The zone which the node stays in.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The instance type of the node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "sg_id",
					Description: `(Optional) The safe-group id.`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional) The path which volumn is going to be mounted.`,
				},
				resource.Attribute{
					Name:        "docker_graph_path",
					Description: `(Optional) The docker graph path is going to mounted.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name ot node.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `(Optional) The kubernetes version of the cluster.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password of each node.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) The key_id of each node(if using key pair to access).`,
				},
				resource.Attribute{
					Name:        "require_wan_ip",
					Description: `(Optional) Indicate whether wan ip is needed.`,
				},
				resource.Attribute{
					Name:        "user_script",
					Description: `(Optional) User defined script in a base64-format. The script runs after the kubernetes component is ready on node. see more from CCS api documents. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `The kubernetes version of the cluster`,
				},
				resource.Attribute{
					Name:        "nodes_num",
					Description: `The node number of the cluster`,
				},
				resource.Attribute{
					Name:        "nodes_status",
					Description: `The node status of the cluster`,
				},
				resource.Attribute{
					Name:        "total_cpu",
					Description: `The total cpu of the cluster`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `The total memory of the cluster`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `The kubernetes version of the cluster`,
				},
				resource.Attribute{
					Name:        "nodes_num",
					Description: `The node number of the cluster`,
				},
				resource.Attribute{
					Name:        "nodes_status",
					Description: `The node status of the cluster`,
				},
				resource.Attribute{
					Name:        "total_cpu",
					Description: `The total cpu of the cluster`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `The total memory of the cluster`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_container_cluster_instance",
			Category:         "Container Cluster Resources",
			ShortDescription: `Provides a TencentCloud Container Cluster Instance resource.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"cluster",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The id of the cluster.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Required) The cpu of the node.`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `(Required) The memory of the node.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The network bandwidth of the node.`,
				},
				resource.Attribute{
					Name:        "bandwidth_type",
					Description: `(Required) The network type of the node.`,
				},
				resource.Attribute{
					Name:        "require_wan_ip",
					Description: `(Optional) Indicate whether wan ip is needed.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The subnet id which the node stays in.`,
				},
				resource.Attribute{
					Name:        "is_vpc_gateway",
					Description: `(Required) Describe whether the node enable the gateway capability.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `(Required) The size of the data volumn.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) The type of the data volumn. see more from CVM.`,
				},
				resource.Attribute{
					Name:        "root_size",
					Description: `(Required) The size of the root volumn.`,
				},
				resource.Attribute{
					Name:        "root_type",
					Description: `(Optional) The type of the root volumn. see more from CVM.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specify vpc which the node(s) stay in.`,
				},
				resource.Attribute{
					Name:        "cvm_type",
					Description: `(Optional) The type of node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The puchase duration of the node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The zone which the node stays in.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The instance type of the node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "sg_id",
					Description: `(Optional) The safe-group id.`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional) The path which volumn is going to be mounted.`,
				},
				resource.Attribute{
					Name:        "docker_graph_path",
					Description: `(Optional) The docker graph path is going to mounted.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password of each node.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) The key_id of each node(if using key pair to access).`,
				},
				resource.Attribute{
					Name:        "unschedulable",
					Description: `(Optional) Determine whether the node will be schedulable. 0 is the default meaning node will be schedulable. 1 for unschedulable.`,
				},
				resource.Attribute{
					Name:        "user_script",
					Description: `(Optional) User defined script in a base64-format. The script runs after the kubernetes component is ready on node. see more from CCS api documents. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "abnormal_reason",
					Description: `Describe the reason when node is in abnormal state(if it was).`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `An id identify the node, provided by cvm.`,
				},
				resource.Attribute{
					Name:        "is_normal",
					Description: `Describe whether the node is normal.`,
				},
				resource.Attribute{
					Name:        "wan_ip",
					Description: `Describe the wan ip of the node.`,
				},
				resource.Attribute{
					Name:        "lan_ip",
					Description: `Describe the lan ip of the node.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "abnormal_reason",
					Description: `Describe the reason when node is in abnormal state(if it was).`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `An id identify the node, provided by cvm.`,
				},
				resource.Attribute{
					Name:        "is_normal",
					Description: `Describe whether the node is normal.`,
				},
				resource.Attribute{
					Name:        "wan_ip",
					Description: `Describe the wan ip of the node.`,
				},
				resource.Attribute{
					Name:        "lan_ip",
					Description: `Describe the lan ip of the node.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cos_bucket",
			Category:         "COS Resources",
			ShortDescription: `Provides a COS resource to create a COS bucket and set its attributes.`,
			Description:      ``,
			Keywords: []string{
				"cos",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required, ForceNew) The name of a bucket to be created.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The canned ACL to apply. Available values include private, public-read, and public-read-write. Defaults to private.`,
				},
				resource.Attribute{
					Name:        "cors_rules",
					Description: `(Optional) A rule of Cross-Origin Resource Sharing (documented below).`,
				},
				resource.Attribute{
					Name:        "lifecycle_rules",
					Description: `(Optional) A configuration of object lifecycle management (documented below).`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `(Optional) A website object(documented below). The ` + "`" + `lifecycle_rules` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "filter_prefix",
					Description: `(Required) Object key prefix identifying one or more objects to which the rule applies.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `(Optional) Specifies a period in the object's expire (documented below).`,
				},
				resource.Attribute{
					Name:        "transition",
					Description: `(Optional) Specifies a period in the object's transitions (documented below). The ` + "`" + `transition` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Required) Specifies the storage class to which you want the object to transition. Available values include STANDARD, STANDARD_IA and ARCHIVE.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) Specifies the date after which you want the corresponding action to take effect.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Optional) Specifies the number of days after object creation when the specific rule action takes effect. The ` + "`" + `expiration` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) Specifies the date after which you want the corresponding action to take effect.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Optional) Specifies the number of days after object creation when the specific rule action takes effect. The ` + "`" + `website` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `(Optional) An absolute path to the document to return in case of a 4XX error.`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `(Optional) COS returns this index document when requests are made to the root domain or any of the subfolders. The ` + "`" + `cors_rules` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `(Required) Specifies which headers are allowed.`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `(Required) Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `(Required) Specifies which origins are allowed.`,
				},
				resource.Attribute{
					Name:        "expose_headers",
					Description: `(Optional) Specifies expose header in the response.`,
				},
				resource.Attribute{
					Name:        "max_age_seconds",
					Description: `(Optional) Specifies time in seconds that browser can cache the response for a preflight request. ## Import COS bucket can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cos_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cos_bucket_object",
			Category:         "COS Resources",
			ShortDescription: `Provides a COS object resource to put an object(content or file) to the bucket.`,
			Description:      ``,
			Keywords: []string{
				"cos",
				"bucket",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required, ForceNew) The name of a bucket to use.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required, ForceNew) The name of the object once it is in the bucket.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The canned ACL to apply. Available values include private, public-read, and public-read-write. Defaults to private.`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Optional) Specifies caching behavior along the request/reply chain.For further details，RFC2616 can be referred.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Optional) Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Optional) Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Optional) The ETag generated for the object (an MD5 sum of the object content).`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) The path to the source file being uploaded to the bucket.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optional) Object storage type, Available values include STANDARD, STANDARD_IA and ARCHIVE.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dc_gateway",
			Category:         "DCG Resources",
			ShortDescription: `Provides a resource to creating direct connect gateway instance.`,
			Description:      ``,
			Keywords: []string{
				"dcg",
				"dc",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the DCG.`,
				},
				resource.Attribute{
					Name:        "network_instance_id",
					Description: `(Required, ForceNew) If the 'network_type' value is 'VPC', the available value is VPC ID. But when the 'network_type' value is 'CCN', the available value is CCN instance ID.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Required, ForceNew) Type of associated network, the available value include 'VPC' and 'CCN'.`,
				},
				resource.Attribute{
					Name:        "gateway_type",
					Description: `(Optional, ForceNew) Type of the gateway, the available value include 'NORMAL' and 'NAT'. Default is 'NORMAL’. NOTES: CCN only supports 'NORMAL' and a vpc can create two DCGs, the one is NAT type and the other is non-NAT type. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cnn_route_type",
					Description: `Type of CCN route, the available value include 'BGP' and 'STATIC'. The property is available when the DCG type is CCN gateway and BGP enabled.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `Indicates whether the BGP is enabled. ## Import Direct connect gateway instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_dc_gateway.instance dcg-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cnn_route_type",
					Description: `Type of CCN route, the available value include 'BGP' and 'STATIC'. The property is available when the DCG type is CCN gateway and BGP enabled.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `Indicates whether the BGP is enabled. ## Import Direct connect gateway instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_dc_gateway.instance dcg-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dc_gateway_ccn_route",
			Category:         "DCG Resources",
			ShortDescription: `Provides a resource to creating direct connect gateway route entry.`,
			Description:      ``,
			Keywords: []string{
				"dcg",
				"dc",
				"gateway",
				"ccn",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, ForceNew) A network address segment of IDC.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `(Required, ForceNew) ID of the DCG ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "as_path",
					Description: `As_Path list of the BGP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "as_path",
					Description: `As_Path list of the BGP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dcx",
			Category:         "DC Resources",
			ShortDescription: `Provides a resource to creating dedicated tunnels instances.`,
			Description:      ``,
			Keywords: []string{
				"dc",
				"dcx",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dc_id",
					Description: `(Required, ForceNew) ID of the DC to be queried, application deployment offline.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `(Required, ForceNew) ID of the DC Gateway. Currently only new in the console.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the dedicated tunnel.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the VPC or BMVPC.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional, ForceNew) Bandwidth of the DC.`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `(Optional, ForceNew) BGP ASN of the user. A required field within BGP.`,
				},
				resource.Attribute{
					Name:        "bgp_auth_key",
					Description: `(Optional, ForceNew) BGP key of the user.`,
				},
				resource.Attribute{
					Name:        "customer_address",
					Description: `(Optional, ForceNew) Interconnect IP of the DC within client.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Optional, ForceNew) Type of the network, and available values include VPC, BMVPC and CCN. The default value is VPC.`,
				},
				resource.Attribute{
					Name:        "route_filter_prefixes",
					Description: `(Optional, ForceNew) Static route, the network address of the user IDC. It can be modified after setting but cannot be deleted. AN unable field within BGP.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `(Optional, ForceNew) Type of the route, and available values include BGP and STATIC. The default value is BGP.`,
				},
				resource.Attribute{
					Name:        "tencent_address",
					Description: `(Optional, ForceNew) Interconnect IP of the DC within Tencent.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Optional, ForceNew) Vlan of the dedicated tunnels, and the range of values is [0-3000]. '0' means that only one tunnel can be created for the physical connect. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the dedicated tunnels, and available values include PENDING, ALLOCATING, ALLOCATED, ALTERING, DELETING, DELETED, COMFIRMING and REJECTED.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the dedicated tunnels, and available values include PENDING, ALLOCATING, ALLOCATED, ALTERING, DELETING, DELETED, COMFIRMING and REJECTED.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dnat",
			Category:         "VPC Resources",
			ShortDescription: `Provides a port mapping/forwarding of destination network address translation (DNAT/DNAPT) resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"dnat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "nat_id",
					Description: `(Required, Forces new resource) The ID for the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, Forces new resource) The VPC ID for the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, Forces new resource) The ip protocol, valid value is tcp|udp.`,
				},
				resource.Attribute{
					Name:        "elastic_ip",
					Description: `(Required, Forces new resource) The elastic IP of NAT gateway association, must a [Elastic IP](eip.html).`,
				},
				resource.Attribute{
					Name:        "elastic_port",
					Description: `(Required, Forces new resource) The external port, valid value is 1~65535.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Required, Forces new resource) The internal ip, must a private ip (VPC IP).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eip",
			Category:         "CVM Resources",
			ShortDescription: `Provides a TencentCloud EIP resource.`,
			Description:      ``,
			Keywords: []string{
				"cvm",
				"eip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The eip's name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The EIP id, something like ` + "`" + `eip-xxxxxxx` + "`" + `, use this for EIP assocication.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The elastic ip address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The EIP current status. ## Import EIPs can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_eip.foo eip-nyvf60va ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The EIP id, something like ` + "`" + `eip-xxxxxxx` + "`" + `, use this for EIP assocication.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The elastic ip address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The EIP current status. ## Import EIPs can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_eip.foo eip-nyvf60va ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eip_association",
			Category:         "CVM Resources",
			ShortDescription: `Provides a TencentCloud EIP association.`,
			Description:      ``,
			Keywords: []string{
				"cvm",
				"eip",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "eip_id",
					Description: `(Required) The eip's id.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) The instance id going to bind with the EIP. This field is conflict with ` + "`" + `network_interface_id` + "`" + ` and ` + "`" + `private_ip` + "`" + ` fields.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `(Optional) Indicates the network interface id like ` + "`" + `eni-xxxxxx` + "`" + `. This field is conflict with ` + "`" + `instance_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) Indicates an IP belongs to the ` + "`" + `network_interface_id` + "`" + `. This field is conflict with ` + "`" + `instance_id` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The association id.`,
				},
				resource.Attribute{
					Name:        "eip_id",
					Description: `The id of the EIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance id of the EIP bound with.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `The network interface id.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) The IP belongs to the ` + "`" + `network_interface_id` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The association id.`,
				},
				resource.Attribute{
					Name:        "eip_id",
					Description: `The id of the EIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance id of the EIP bound with.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `The network interface id.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) The IP belongs to the ` + "`" + `network_interface_id` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_instance",
			Category:         "CVM Resources",
			ShortDescription: `Provides a CVM instance resource.`,
			Description:      ``,
			Keywords: []string{
				"cvm",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) The Image to use for the instance. CVM instance's image can be replaced via changing 'image_id'.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) The available zone that the CVM instance locates at.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name of the CVM. This instance_name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://. If not specified, Terraform will autogenerate a default name is ` + "`" + `CVM-Instance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) The type of instance to start.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The hostname of CVM.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The project CVM belongs to, default to 0.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_period",
					Description: `(Optional) The tenancy (time unit is month) of the prepaid instance,`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `(Optional) When enabled, the CVM instance will be renew automatically when it reach the end of the prepaid tenancy,`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional) Internet charge type of the instance, Valid values are ` + "`" + `BANDWIDTH_PREPAID` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `BANDWIDTH_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `. Default is ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional) Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bit per second). Value range: [0, 200], If this value is not specified, then automatically sets it to 0 Mbps.`,
				},
				resource.Attribute{
					Name:        "allocate_public_ip",
					Description: `(Optional) Associate a public ip address with an instance in a VPC or Classic. Boolean value, Default is false.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The id of a VPC network. If you want to create instances in VPC network, this parameter must be set.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The id of a VPC subnetwork. If you want to create instances in VPC network, this parameter must be set.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) The private ip to be assigned to this instance, must be in the provided subnet and available.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) A list of security group ids to associate with.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional) Valid values are ` + "`" + `LOCAL_BASIC` + "`" + `, ` + "`" + `LOCAL_SSD` + "`" + `, ` + "`" + `CLOUD_BASIC` + "`" + `, ` + "`" + `CLOUD_SSD` + "`" + ` and ` + "`" + `CLOUD_PREMIUM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Size of the system disk, value range: 50GB ~ 1TB. Default is 50GB.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `(Optional) Settings for data disk. In each disk, ` + "`" + `data_disk_type` + "`" + ` indicates the disk type, valid values are ` + "`" + `LOCAL_BASIC` + "`" + `, ` + "`" + `LOCAL_SSD` + "`" + `, ` + "`" + `CLOUD_BASIC` + "`" + `, ` + "`" + `CLOUD_SSD` + "`" + ` and ` + "`" + `CLOUD_PREMIUM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_security_service",
					Description: `(Optional) Disable enhance service for security, it is enabled by default. When this options is set, security agent won't be installed.`,
				},
				resource.Attribute{
					Name:        "disable_monitor_service",
					Description: `(Optional) Disable enhance service for monitor, it is enabled by default. When this options is set, monitor agent won't be installed.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional) The key pair to use for the instance, it looks like ` + "`" + `skey-16jig7tx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password to an instance. In order to take effect new password, the instance will be restarted after modifying the password.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to be specified into this instance. Must be encrypted in base64 format and limited in 16 KB.`,
				},
				resource.Attribute{
					Name:        "user_data_raw",
					Description: `(Optional) The user data to be specified into this instance, plain text. Conflicts with ` + "`" + `user_data` + "`" + `. Limited in 16 KB after encrypted in base64 format.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID, something looks like ` + "`" + `ins-xxxxxx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `The Status of the instance.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The Local IP Address of the instance.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The instance public ip.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC Id associated with the instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The Subnet Id associated with the instance.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `The system disk type on the instance.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `The system disk type on the instance.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `The data disks info. In each data disk, ` + "`" + `data_disk_type` + "`" + ` is the disk type. ` + "`" + `data_disk_size` + "`" + ` is the size of the disk.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `The key pair id of the instance. ## Import CVM instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_instance.foo ins-2qol3a80 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID, something looks like ` + "`" + `ins-xxxxxx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `The Status of the instance.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The Local IP Address of the instance.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The instance public ip.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC Id associated with the instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The Subnet Id associated with the instance.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `The system disk type on the instance.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `The system disk type on the instance.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `The data disks info. In each data disk, ` + "`" + `data_disk_type` + "`" + ` is the disk type. ` + "`" + `data_disk_size` + "`" + ` is the size of the disk.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `The key pair id of the instance. ## Import CVM instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_instance.foo ins-2qol3a80 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_key_pair",
			Category:         "CVM Resources",
			ShortDescription: `Provides a TencentCloud key pair resource.`,
			Description:      ``,
			Keywords: []string{
				"cvm",
				"key",
				"pair",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `(Force new resource) The key pair's name. It is the only in one TencentCloud account.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Force new resource) You can import an existing public key and using TencentCloud key pair to manage it. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the key pair, something like ` + "`" + `skey-xxxxxxx` + "`" + `, use this for instance creation and resetting. ## Import Key pairs can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_key_pair.foo skey-17634f05 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the key pair, something like ` + "`" + `skey-xxxxxxx` + "`" + `, use this for instance creation and resetting. ## Import Key pairs can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_key_pair.foo skey-17634f05 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_lb",
			Category:         "LB Resources",
			ShortDescription: `Provides a Load Balancer resource.`,
			Description:      ``,
			Keywords: []string{
				"lb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The network type of the LB, valid choices: "OPEN", "INTERNAL".`,
				},
				resource.Attribute{
					Name:        "forward",
					Description: `(Optional) The type of the LB, valid choices: "CLASSIC", "APPLICATION".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the LB.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The VPC ID of the LB, unspecified or 0 stands for CVM basic network.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The project id of the LB, unspecified or 0 stands for default project. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the LB.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of the LB.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_account",
			Category:         "MySQL Resources",
			ShortDescription: `Provides a MySQL account resource for database management. A MySQL instance supports multiple database account.`,
			Description:      ``,
			Keywords: []string{
				"mysql",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mysql_id",
					Description: `(Required, ForceNew) Instance ID to which the account belongs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Account name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Operation password.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Database description.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_account_privilege",
			Category:         "MySQL Resources",
			ShortDescription: `Provides a mysql account privilege resource to grant different access privilege to different database. A database can be granted by multiple account.`,
			Description:      ``,
			Keywords: []string{
				"mysql",
				"account",
				"privilege",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required, ForceNew) Account name.`,
				},
				resource.Attribute{
					Name:        "database_names",
					Description: `(Required) List of specified database name.`,
				},
				resource.Attribute{
					Name:        "mysql_id",
					Description: `(Required, ForceNew) Instance ID.`,
				},
				resource.Attribute{
					Name:        "privileges",
					Description: `(Optional) Database permissions. Available values for Privileges: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "REFERENCES", "INDEX", "ALTER", "CREATE TEMPORARY TABLES", "LOCK TABLES","EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE", "ALTER ROUTINE", "EVENT", and "TRIGGER".`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_backup_policy",
			Category:         "MySQL Resources",
			ShortDescription: `Provides a mysql policy resource to create a backup policy.`,
			Description:      ``,
			Keywords: []string{
				"mysql",
				"backup",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mysql_id",
					Description: `(Required, ForceNew) Instance ID to which policies will be applied.`,
				},
				resource.Attribute{
					Name:        "backup_model",
					Description: `(Optional) Backup method. Supported values include: physical - physical backup, and logical - logical backup.`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `(Optional) Instance backup time, in the format of "HH:mm-HH:mm". Time setting interval is four hours. Default to "02:00-06:00". The following value can be supported: 02:00\-06:00, 06:00\-10:00, 10:00\-14:00, 14:00\-18:00, 18:00\-22:00, and 22:00\-02:00.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `(Optional) Instance backup retention days. Valid values: [7-730]. And default value is 7. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "binlog_period",
					Description: `Retention period for binlog in days.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "binlog_period",
					Description: `Retention period for binlog in days.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_instance",
			Category:         "MySQL Resources",
			ShortDescription: `Provides a mysql instance resource to create master database instances.`,
			Description:      ``,
			Keywords: []string{
				"mysql",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) The name of a mysql instance.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `(Required) Memory size (in MB).`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `(Required) Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Required) Disk size (in GB).`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Indicates which availability zone will be used.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional, ForceNew) The version number of the database engine to use. Supported versions include 5.5/5.6/5.7, and default is 5.7.`,
				},
				resource.Attribute{
					Name:        "first_slave_zone",
					Description: `(Optional, ForceNew) Zone information about first slave instance.`,
				},
				resource.Attribute{
					Name:        "internet_service",
					Description: `(Optional) Indicates whether to enable the access to an instance from public network: 0 - No, 1 - Yes.`,
				},
				resource.Attribute{
					Name:        "intranet_port",
					Description: `(Optional) Public access port, rang form 1024 to 65535 and default value is 3306.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) List of parameters to use.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID, default value is 0.`,
				},
				resource.Attribute{
					Name:        "second_slave_zone",
					Description: `(Optional, ForceNew) Zone information about second slave instance.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) Security groups to use.`,
				},
				resource.Attribute{
					Name:        "slave_deploy_mode",
					Description: `(Optional, ForceNew) Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.`,
				},
				resource.Attribute{
					Name:        "slave_sync_mode",
					Description: `(Optional, ForceNew) Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Private network ID. If vpc_id is set, this value is required.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Instance tags.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of VPC, which can be modified once every 24 hours and can’t be removed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "gtid",
					Description: `Indicates whether GTID is enable. 0 - Not enabled; 1 - Enabled.`,
				},
				resource.Attribute{
					Name:        "internet_host",
					Description: `host for public access.`,
				},
				resource.Attribute{
					Name:        "internet_port",
					Description: `Access port for public access.`,
				},
				resource.Attribute{
					Name:        "intranet_ip",
					Description: `instance intranet IP.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Indicates whether the instance is locked. 0 - No; 1 - Yes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Available values: 0 - Creating; 1 - Running; 4 - Isolating; 5 – Isolated.`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Indicates which kind of operations is being executed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gtid",
					Description: `Indicates whether GTID is enable. 0 - Not enabled; 1 - Enabled.`,
				},
				resource.Attribute{
					Name:        "internet_host",
					Description: `host for public access.`,
				},
				resource.Attribute{
					Name:        "internet_port",
					Description: `Access port for public access.`,
				},
				resource.Attribute{
					Name:        "intranet_ip",
					Description: `instance intranet IP.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Indicates whether the instance is locked. 0 - No; 1 - Yes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Available values: 0 - Creating; 1 - Running; 4 - Isolating; 5 – Isolated.`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Indicates which kind of operations is being executed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_readonly_instance",
			Category:         "MySQL Resources",
			ShortDescription: `Provides a mysql instance resource to create read-only database instances.`,
			Description:      ``,
			Keywords: []string{
				"mysql",
				"readonly",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) The name of a mysql instance.`,
				},
				resource.Attribute{
					Name:        "master_instance_id",
					Description: `(Required, ForceNew) Indicates the master instance ID of recovery instances.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `(Required) Memory size (in MB).`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Required) Disk size (in GB).`,
				},
				resource.Attribute{
					Name:        "intranet_port",
					Description: `(Optional) Public access port, rang form 1024 to 65535 and default value is 3306.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) Security groups to use.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Private network ID. If vpc_id is set, this value is required.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Instance tags.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of VPC, which can be modified once every 24 hours and can’t be removed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "intranet_ip",
					Description: `instance intranet IP.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Indicates whether the instance is locked. 0 - No; 1 - Yes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Available values: 0 - Creating; 1 - Running; 4 - Isolating; 5 – Isolated.`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Indicates which kind of operations is being executed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "intranet_ip",
					Description: `instance intranet IP.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Indicates whether the instance is locked. 0 - No; 1 - Yes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Available values: 0 - Creating; 1 - Running; 4 - Isolating; 5 – Isolated.`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Indicates which kind of operations is being executed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_nat_gateway",
			Category:         "VPC Resources",
			ShortDescription: `Provides a resource to create a VPC NAT Gateway.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"nat",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, Forces new resource) The VPC ID.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `(Required) The upper limit of concurrent connection of NAT gateway, for example: 1000000, 3000000, 10000000. To learn more, please refer to [Virtual Private Cloud Gateway Description](https://intl.cloud.tencent.com/doc/product/215/1682).`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The maximum public network output bandwidth of the gateway (unit: Mbps), for example: 10, 20, 50, 100, 200, 500, 1000, 2000, 5000. For more information, please refer to [Virtual Private Cloud Gateway Description](https://intl.cloud.tencent.com/doc/product/215/1682).`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `(Required) Elastic IP arrays bound to the gateway, For more information on elastic IP, please refer to [Elastic IP](eip.html). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `The upper limit of concurrent connection of NAT gateway.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of the gateway (unit: Mbps).`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `Elastic IP arrays bound to the gateway`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `The upper limit of concurrent connection of NAT gateway.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of the gateway (unit: Mbps).`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `Elastic IP arrays bound to the gateway`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_redis_backup_config",
			Category:         "Redis Resources",
			ShortDescription: `Use this data source to query which instance types of Redis are available in a specific region.`,
			Description:      ``,
			Keywords: []string{
				"redis",
				"backup",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "backup_period",
					Description: `(Required) Specifys which day the backup action should take place. Supported values include: Monday，Tuesday, Wednesday, Thursday, Friday, Saturday and Sunday.`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `(Required) Specifys what time the backup action should take place.`,
				},
				resource.Attribute{
					Name:        "redis_id",
					Description: `(Required, ForceNew) ID of a Redis instance to which the policy will be applied. ## Import Redis backup config can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_redis_backup_config.redisconfig redis-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_redis_instance",
			Category:         "Redis Resources",
			ShortDescription: `Provides a resource to create a Redis instance and set its attributes.`,
			Description:      ``,
			Keywords: []string{
				"redis",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) The available zone ID of an instance to be created., refer to tencentcloud_redis_zone_config.list`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `(Required) The memory volume of an available instance(in MB), refer to tencentcloud_redis_zone_config.list[zone].mem_sizes`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password for a Redis user，which should be 8 to 16 characters.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Instance name.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, ForceNew) The port used to access a redis instance. The default value is 6379. And this value can't be changed after creation, or the Redis instance will be recreated.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Specifies which project the instance should belong to.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional, ForceNew) ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) Specifies which subnet the instance should belong to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, ForceNew) Instance type. Available values: master_slave_redis.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of the vpc with which the instance is to be associated. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the instance was created.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of an instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of an instance，maybe: init, processing, online, isolate and todelete. ## Import Redis instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_redis_instance.redislab redis-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the instance was created.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of an instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of an instance，maybe: init, processing, online, isolate and todelete. ## Import Redis instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_redis_instance.redislab redis-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_route_entry",
			Category:         "VPC Resources",
			ShortDescription: `Provides a resource to create a routing entry in a VPC routing table.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"route",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, Forces new resource) The VPC ID.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required, Forces new resource) The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, Forces new resource) The RouteEntry's target network segment.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `(Required, Forces new resource) The next hop type. Available value is ` + "`" + `public_gateway` + "`" + `、` + "`" + `vpn_gateway` + "`" + `、` + "`" + `sslvpn_gateway` + "`" + `、` + "`" + `dc_gateway` + "`" + `、` + "`" + `peering_connection` + "`" + `、` + "`" + `nat_gateway` + "`" + ` and ` + "`" + `instance` + "`" + `. ` + "`" + `instance` + "`" + ` points to CVM Instance.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `(Required, Forces new resource) The route entry's next hub. CVM instance ID or VPC router interface ID. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The RouteEntry's target network segment.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `The next hub type.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `The route entry's next hub.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The RouteEntry's target network segment.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `The next hub type.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `The route entry's next hub.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_route_table",
			Category:         "VPC Resources",
			ShortDescription: `Provides a resource to create a VPC routing table.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"route",
				"table",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of routing table.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of VPC to which the route table should be associated. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the routing table.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default routing table.`,
				},
				resource.Attribute{
					Name:        "route_entry_ids",
					Description: `ID list of the routing entries.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `ID list of the subnets associated with this route table. ## Import Vpc routetable instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_route_table.test route_table_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the routing table.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default routing table.`,
				},
				resource.Attribute{
					Name:        "route_entry_ids",
					Description: `ID list of the routing entries.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `ID list of the subnets associated with this route table. ## Import Vpc routetable instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_route_table.test route_table_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_route_table_entry",
			Category:         "VPC Resources",
			ShortDescription: `Provides a resource to create an entry of a routing table.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"route",
				"table",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "destination_cidr_block",
					Description: `(Required, ForceNew) Destination address block.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `(Required, ForceNew) ID of next-hop gateway. Note: when 'next_type' is EIP, GatewayId should be '0'.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `(Required, ForceNew) Type of next-hop, and available values include CVM, VPN, DIRECTCONNECT, PEERCONNECTION, SSLVPN, NAT, NORMAL_CVM, EIP and CCN.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required, ForceNew) ID of routing table to which this entry belongs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) Description of the routing table entry.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_security_group",
			Category:         "VPC Resources",
			ShortDescription: `Provides a security group resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the security group. Name should be unique in each project, and no more than 60 characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The security group's description, maximum length is 100 characters.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The security group's project, default is 0. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group. ## Import Security group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_security_group.foo sg-ey3wmiz1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group. ## Import Security group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_security_group.foo sg-ey3wmiz1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_security_group_rule",
			Category:         "VPC Resources",
			ShortDescription: `Provides an security group rule resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"security",
				"group",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required, Forces new resource) The security group to apply this rule to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, Forces new resource) The type of rule being created. Valid options are "ingress" (inbound) or "egress" (outbound).`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `(Optional, Forces new resource) can be IP, or CIDR block.`,
				},
				resource.Attribute{
					Name:        "source_sgid",
					Description: `(Optional, Forces new resource) The ID of a security group rule. Either ` + "`" + `cidr_ip` + "`" + ` or ` + "`" + `source_sgid` + "`" + ` must be specified, but it isn't supported simultaneously.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional, Forces new resource) Support "UDP"、"TCP"、"ICMP", Not configured means all protocols.`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `(Optional, Forces new resource) examples, Single port: "53"、Multiple ports: "80,8080,443"、Continuous port: "80-90", Not configured to represent all ports.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required, Forces new resource) Policy of rule, "accept" or "drop". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group rule.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of rule, "ingress" or "egress".`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `The source of rule, IP or CIDR block.`,
				},
				resource.Attribute{
					Name:        "source_sgid",
					Description: `The ID of a security group rule.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The policy of rule, "accept" or "drop".`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group rule.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of rule, "ingress" or "egress".`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `The source of rule, IP or CIDR block.`,
				},
				resource.Attribute{
					Name:        "source_sgid",
					Description: `The ID of a security group rule.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The policy of rule, "accept" or "drop".`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_subnet",
			Category:         "VPC Resources",
			ShortDescription: `Provide a resource to create a VPC subnet.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) The availability zone within which the subnet should be created.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, ForceNew) A network address block of the subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of subnet to be created.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the VPC to be associated.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `(Optional) Indicates whether multicast is enabled. The default value is 'true'.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Optional) ID of a routing table to which the subnet should be associated. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "available_ip_count",
					Description: `The number of available IPs.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of subnet resource.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default VPC for this region. ## Import Vpc subnet instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_subnet.test subnet_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "available_ip_count",
					Description: `The number of available IPs.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of subnet resource.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default VPC for this region. ## Import Vpc subnet instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_subnet.test subnet_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpc",
			Category:         "VPC Resources",
			ShortDescription: `Provide a resource to create a VPC.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, ForceNew) A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VPC.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `(Optional) The DNS server list of the VPC. And you can specify 0 to 5 servers to this list.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `(Optional) Indicates whether VPC multicast is enabled. The default value is 'true'. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default VPC for this region. ## Import Vpc instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_vpc.test vpc-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default VPC for this region. ## Import Vpc instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import tencentcloud_vpc.test vpc-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"tencentcloud_alb_server_attachment":      0,
		"tencentcloud_as_attachment":              1,
		"tencentcloud_as_lifecycle_hook":          2,
		"tencentcloud_as_notification":            3,
		"tencentcloud_as_scaling_config":          4,
		"tencentcloud_as_scaling_group":           5,
		"tencentcloud_as_scaling_policy":          6,
		"tencentcloud_as_schedule":                7,
		"tencentcloud_cbs_snapshot":               8,
		"tencentcloud_cbs_snapshot_policy":        9,
		"tencentcloud_cbs_storage":                10,
		"tencentcloud_cbs_storage_attachment":     11,
		"tencentcloud_ccn":                        12,
		"tencentcloud_ccn_attachment":             13,
		"tencentcloud_ccn_bandwidth_limit":        14,
		"tencentcloud_container_cluster":          15,
		"tencentcloud_container_cluster_instance": 16,
		"tencentcloud_cos_bucket":                 17,
		"tencentcloud_cos_bucket_object":          18,
		"tencentcloud_dc_gateway":                 19,
		"tencentcloud_dc_gateway_ccn_route":       20,
		"tencentcloud_dcx":                        21,
		"tencentcloud_dnat":                       22,
		"tencentcloud_eip":                        23,
		"tencentcloud_eip_association":            24,
		"tencentcloud_instance":                   25,
		"tencentcloud_key_pair":                   26,
		"tencentcloud_lb":                         27,
		"tencentcloud_mysql_account":              28,
		"tencentcloud_mysql_account_privilege":    29,
		"tencentcloud_mysql_backup_policy":        30,
		"tencentcloud_mysql_instance":             31,
		"tencentcloud_mysql_readonly_instance":    32,
		"tencentcloud_nat_gateway":                33,
		"tencentcloud_redis_backup_config":        34,
		"tencentcloud_redis_instance":             35,
		"tencentcloud_route_entry":                36,
		"tencentcloud_route_table":                37,
		"tencentcloud_route_table_entry":          38,
		"tencentcloud_security_group":             39,
		"tencentcloud_security_group_rule":        40,
		"tencentcloud_subnet":                     41,
		"tencentcloud_vpc":                        42,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
