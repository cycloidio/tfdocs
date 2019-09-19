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
			Category:         "CLB Resources",
			ShortDescription: `Provides an tencentcloud application load balancer servers attachment as a resource, to attach and detach instances from load balancer.`,
			Description:      ``,
			Keywords: []string{
				"clb",
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
					Name:        "rule_id",
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
			ShortDescription: `Provides a resource to create a CBS snapshot.`,
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
			ShortDescription: `Provides a resource to create a CBS.`,
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
			Type:             "tencentcloud_clb_attachment",
			Category:         "CLB Resources",
			ShortDescription: `Provides a resource to create a CLB attachment.`,
			Description:      ``,
			Keywords: []string{
				"clb",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required, ForceNew) Id of the clb.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) Id of the clb listener.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) Information of the backends to be attached.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional, ForceNew) Id of the clb listener rule. Only supports listeners of 'HTTPS' and 'HTTP' protocol. The ` + "`" + `targets` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) Id of the backend server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port of the backend server.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Forwarding weight of the backend service, the range of [0, 100], defaults to 10. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `Type of protocol within the listener. ## Import CLB attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_attachment.foo loc-4xxr2cy7#lbl-hh141sn9#lb-7a0t6zqb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "protocol_type",
					Description: `Type of protocol within the listener. ## Import CLB attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_attachment.foo loc-4xxr2cy7#lbl-hh141sn9#lb-7a0t6zqb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_instance",
			Category:         "CLB Resources",
			ShortDescription: `Provides a resource to create a CLB instance.`,
			Description:      ``,
			Keywords: []string{
				"clb",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_name",
					Description: `(Required) Name of the CLB. The name can only contain Chinese characters, English letters, numbers, underscore and hyphen '-'.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Required, ForceNew) Type of CLB instance, and available values include 'OPEN' and 'INTERNAL'.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) Id of the project within the CLB instance, '0' - Default Project.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) Security groups of the CLB instance. Only supports 'OPEN' CLBs.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) Subnet id of the CLB. Effective only for CLB within the VPC. Only supports 'INTERNAL' CLBs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, ForceNew) The available tags within this CLB.`,
				},
				resource.Attribute{
					Name:        "target_region_info_region",
					Description: `(Optional) Region information of backend services are attached the CLB instance. Only supports 'OPEN' CLBs.`,
				},
				resource.Attribute{
					Name:        "target_region_info_vpc_id",
					Description: `(Optional) Vpc information of backend services are attached the CLB instance. Only supports 'OPEN' CLBs.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) VPC id of the CLB. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "clb_vips",
					Description: `The virtual service address table of the CLB. ## Import CLB instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_instance.foo lb-7a0t6zqb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_vips",
					Description: `The virtual service address table of the CLB. ## Import CLB instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_instance.foo lb-7a0t6zqb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_listener",
			Category:         "CLB Resources",
			ShortDescription: `Provides a resource to create a CLB listener.`,
			Description:      ``,
			Keywords: []string{
				"clb",
				"listener",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required, ForceNew) Id of the CLB.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `(Required) Name of the CLB listener, and available values can only be Chinese characters, English letters, numbers, underscore and hyphen '-'.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, ForceNew) Type of protocol within the listener, and available values include 'TCP', 'UDP', 'HTTP', 'HTTPS' and 'TCP_SSL'.`,
				},
				resource.Attribute{
					Name:        "certificate_ca_id",
					Description: `(Optional) Id of the client certificate. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol and must be set when the ssl mode is 'MUTUAL'.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) Id of the server certificate. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol and must be set when it is available.`,
				},
				resource.Attribute{
					Name:        "certificate_ssl_mode",
					Description: `(Optional) Type of certificate, and available values inclue 'UNIDIRECTIONAL', 'MUTUAL'. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol and must be set when it is available.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `(Optional) Health threshold of health check, and the default is 3. If a success result is returned for the health check for 3 consecutive times, the backend CVM is identified as healthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_interval_time",
					Description: `(Optional) Interval time of health check. The value range is 5-300 sec, and the default is 5 sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `(Optional) Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_time_out",
					Description: `(Optional) Response timeout of health check. The value range is 2-60 sec, and the default is 2 sec. Response timeout needs to be less than check interval. NOTES: Only supports listeners of 'TCP','UDP','TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `(Optional) Unhealth threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, the CVM is identified as unhealthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, ForceNew) Port of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling method of the CLB listener, and available values include 'WRR' and 'LEAST_CONN'. The default is 'WRR'. NOTES: The listener of HTTP and 'HTTPS' protocol additionally supports the 'IP Hash' method. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "session_expire_time",
					Description: `(Optional) Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as 'WRR', and not available when listener protocol is 'TCP_SSL'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "sni_switch",
					Description: `(Optional, ForceNew) Indicates whether SNI is enabled, and only supported with protocol 'HTTPS'. If enabled, you can set a certificate for each rule in ` + "`" + `tencentcloud_clb_listener_rule` + "`" + `, otherwise all rules have a certificate.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_listener_rule",
			Category:         "CLB Resources",
			ShortDescription: `Provides a resource to create a CLB listener rule.`,
			Description:      ``,
			Keywords: []string{
				"clb",
				"listener",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required) Id of CLB instance.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) Domain name of the listener rule.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) Id of CLB listener.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required, ForceNew) Url of the listener rule.`,
				},
				resource.Attribute{
					Name:        "certificate_ca_id",
					Description: `(Optional, ForceNew) Id of the client certificate. NOTES: Only supports listeners of 'HTTPS' protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional, ForceNew) Id of the server certificate. NOTES: Only supports listeners of 'HTTPS' protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_ssl_mode",
					Description: `(Optional, ForceNew) Type of certificate, and available values inclue 'UNIDIRECTIONAL', 'MUTUAL'. NOTES: Only supports listeners of 'HTTPS' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `(Optional) Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `(Optional) HTTP Status Code. The default is 31 and value range is 1-31. '0b0001' means the return value '1xx' is health. '0b0010' means the return value '2xx' is health. '0b0100' means the return value '3xx' is health. '0b1000' means the return value '4xx' is health. 0b10000 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_http_domain",
					Description: `(Optional) Domain name of health check. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_http_method",
					Description: `(Optional) Methods of health check. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol. The default is 'HEAD', the available value include 'HEAD' and 'GET'.`,
				},
				resource.Attribute{
					Name:        "health_check_http_path",
					Description: `(Optional) Path of health check. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_interval_time",
					Description: `(Optional) Interval time of health check. The value range is 5-300 sec, and the default is 5 sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `(Optional) Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `(Optional) Unhealth threshold of health check, and the default is 3. If the unhealth result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling method of the CLB listener rules, and available values include 'WRR', 'IP HASH' and 'LEAST_CONN'. The default is 'WRR'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "session_expire_time",
					Description: `(Optional) Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as 'WRR', and not available when listener protocol is 'TCP_SSL'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_redirection",
			Category:         "CLB Resources",
			ShortDescription: `Provides a resource to create a CLB redirection.`,
			Description:      ``,
			Keywords: []string{
				"clb",
				"redirection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required, ForceNew) Id of CLB instance.`,
				},
				resource.Attribute{
					Name:        "target_listener_id",
					Description: `(Required, ForceNew) Id of source listener.`,
				},
				resource.Attribute{
					Name:        "target_rule_id",
					Description: `(Required, ForceNew) Rule id of target listener.`,
				},
				resource.Attribute{
					Name:        "is_auto_rewrite",
					Description: `(Optional, ForceNew) Indicates whether automatic forwarding is enable, default is false. If enabled, the source listener and location should be empty, the target listener must be https protocol and port is 443.`,
				},
				resource.Attribute{
					Name:        "source_listener_id",
					Description: `(Optional, ForceNew) Id of source listener.`,
				},
				resource.Attribute{
					Name:        "source_rule_id",
					Description: `(Optional, ForceNew) Rule id of source listener. ## Import CLB redirection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_redirection.foo loc-ft8fmngv#loc-4xxr2cy7#lbl-jc1dx6ju#lbl-asj1hzuo#lb-p7olt9e5 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) A website object(documented below). The ` + "`" + `cors_rules` + "`" + ` object supports the following:`,
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
					Description: `(Optional) Specifies time in seconds that browser can cache the response for a preflight request. The ` + "`" + `lifecycle_rules` + "`" + ` object supports the following:`,
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
					Description: `(Optional) COS returns this index document when requests are made to the root domain or any of the subfolders. ## Import COS bucket can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cos_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
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
			ShortDescription: `Provides a resource to create a NAT forwarding.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"dnat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "elastic_ip",
					Description: `(Required, ForceNew) Network address of the eip.`,
				},
				resource.Attribute{
					Name:        "elastic_port",
					Description: `(Required, ForceNew) Port of the eip.`,
				},
				resource.Attribute{
					Name:        "nat_id",
					Description: `(Required, ForceNew) ID of the nat.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Required, ForceNew) Network address of the backend service.`,
				},
				resource.Attribute{
					Name:        "private_port",
					Description: `(Required, ForceNew) Port of intranet.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, ForceNew) Type of the network protocol, the available values include： TCP and UDP.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the vpc.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the nat forward. ## Import NAT forwarding can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_dnat.foo tcp://vpc-asg3sfa3:nat-1asg3t63@127.15.2.3:8080 ` + "`" + `` + "`" + `` + "`" + ``,
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
			Type:             "tencentcloud_gaap_certificate",
			Category:         "GAAP Resources",
			ShortDescription: `Provides a resource to create a certificate of GAAP.`,
			Description:      ``,
			Keywords: []string{
				"gaap",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "content",
					Description: `(Required, ForceNew) Content of the certificate, and URL encoding. When the certificate is basic authentication, use the ` + "`" + `user:xxx password:xxx` + "`" + ` format, where the password is encrypted with ` + "`" + `htpasswd` + "`" + ` or ` + "`" + `openssl` + "`" + `; When the certificate is ` + "`" + `CA` + "`" + ` or ` + "`" + `SSL` + "`" + `, the format is ` + "`" + `pem` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) Type of the certificate. Available values include: ` + "`" + `BASIC` + "`" + `,` + "`" + `CLIENT` + "`" + `,` + "`" + `SERVER` + "`" + `,` + "`" + `REALSERVER` + "`" + ` and ` + "`" + `PROXY` + "`" + `; ` + "`" + `BASIC` + "`" + ` means basic certificate; ` + "`" + `CLIENT` + "`" + ` means client CA certificate; ` + "`" + `SERVER` + "`" + ` means server SSL certificate; ` + "`" + `REALSERVER` + "`" + ` means realserver CA certificate; ` + "`" + `PROXY` + "`" + ` means proxy SSL certificate.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional, ForceNew) Key of the ` + "`" + `CA` + "`" + ` or ` + "`" + `SSL` + "`" + ` certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the certificate. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "begin_time",
					Description: `Beginning time of the certificate.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the certificate.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Ending time of the certificate.`,
				},
				resource.Attribute{
					Name:        "issuer_cn",
					Description: `Issuer name of the certificate.`,
				},
				resource.Attribute{
					Name:        "subject_cn",
					Description: `Subject name of the certificate. ## Import GAAP certificate can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_certificate.foo cert-d5y6ei3b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "begin_time",
					Description: `Beginning time of the certificate.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the certificate.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Ending time of the certificate.`,
				},
				resource.Attribute{
					Name:        "issuer_cn",
					Description: `Issuer name of the certificate.`,
				},
				resource.Attribute{
					Name:        "subject_cn",
					Description: `Subject name of the certificate. ## Import GAAP certificate can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_certificate.foo cert-d5y6ei3b ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_http_domain",
			Category:         "GAAP Resources",
			ShortDescription: `Provides a resource to create a forward domain of layer7 listener.`,
			Description:      ``,
			Keywords: []string{
				"gaap",
				"http",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) Forward domain of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "basic_auth_id",
					Description: `(Optional) ID of the basic authentication.`,
				},
				resource.Attribute{
					Name:        "basic_auth",
					Description: `(Optional) Indicates whether basic authentication is enable, default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) ID of the server certificate, default value is ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(Optional) ID of the client certificate, default value is ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gaap_auth_id",
					Description: `(Optional) ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "gaap_auth",
					Description: `(Optional) Indicates whether SSL certificate authentication is enable, default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "realserver_auth",
					Description: `(Optional) Indicates whether realserver authentication is enable, default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_domain",
					Description: `(Optional) CA certificate domain of the realserver.`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_id",
					Description: `(Optional) CA certificate ID of the realserver.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_http_rule",
			Category:         "GAAP Resources",
			ShortDescription: `Provides a resource to create a forward rule of layer7 listener.`,
			Description:      ``,
			Keywords: []string{
				"gaap",
				"http",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) Forward rule domain of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Required) Indicates whether health check is enable.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path of the forward rule. Maximum length is 80.`,
				},
				resource.Attribute{
					Name:        "realserver_type",
					Description: `(Required, ForceNew) Type of the realserver, and the available values include ` + "`" + `IP` + "`" + `,` + "`" + `DOMAIN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "realservers",
					Description: `(Required) An information list of GAAP realserver. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `(Optional) Timeout of the health check response, default is 2s.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `(Optional) Method of the health check. Available values includes ` + "`" + `GET` + "`" + ` and ` + "`" + `HEAD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `(Optional) Path of health check. Maximum length is 80.`,
				},
				resource.Attribute{
					Name:        "health_check_status_codes",
					Description: `(Optional) Return code of confirmed normal. Available values includes ` + "`" + `100` + "`" + `,` + "`" + `200` + "`" + `,` + "`" + `300` + "`" + `,` + "`" + `400` + "`" + ` and ` + "`" + `500` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Interval of the health check, default is 5s.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling policy of the layer4 listener, default is ` + "`" + `rr` + "`" + `. Available values include ` + "`" + `rr` + "`" + `,` + "`" + `wrr` + "`" + ` and ` + "`" + `lc` + "`" + `. The ` + "`" + `realservers` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) IP of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Scheduling weight, default is 1. The range of values is [1,100].`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_layer4_listener",
			Category:         "GAAP Resources",
			ShortDescription: `Provides a resource to create a layer4 listener of GAAP.`,
			Description:      ``,
			Keywords: []string{
				"gaap",
				"layer4",
				"listener",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the layer4 listener, the maximum length is 30.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required, ForceNew) Port of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, ForceNew) Protocol of the layer4 listener, and the available values include ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `(Required, ForceNew) ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "realserver_type",
					Description: `(Required, ForceNew) Type of the realserver, and the available values include ` + "`" + `IP` + "`" + `,` + "`" + `DOMAIN` + "`" + `. NOTES: when the ` + "`" + `protocol` + "`" + ` is specified as ` + "`" + `TCP` + "`" + ` and the ` + "`" + `scheduler` + "`" + ` is specified as ` + "`" + `wrr` + "`" + `, the item can only be set to ` + "`" + `IP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `(Optional) Timeout of the health check response, should less than interval, default is 2s. NOTES: Only supports listeners of ` + "`" + `TCP` + "`" + ` protocol and require less than ` + "`" + `interval` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Optional) Indicates whether health check is enable, default is false. NOTES: Only supports listeners of ` + "`" + `TCP` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Interval of the health check, default is 5s. NOTES: Only supports listeners of ` + "`" + `TCP` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "realserver_bind_set",
					Description: `(Optional) An information list of GAAP realserver. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling policy of the layer4 listener, default is ` + "`" + `rr` + "`" + `. Available values include ` + "`" + `rr` + "`" + `,` + "`" + `wrr` + "`" + ` and ` + "`" + `lc` + "`" + `. The ` + "`" + `realserver_bind_set` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) IP of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Scheduling weight, default is 1. The range of values is [1,100]. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer4 listener.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer4 listener.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_layer7_listener",
			Category:         "GAAP Resources",
			ShortDescription: `Provides a resource to create a layer7 listener of GAAP.`,
			Description:      ``,
			Keywords: []string{
				"gaap",
				"layer7",
				"listener",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the layer7 listener, the maximum length is 30.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required, ForceNew) Port of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required, ForceNew) Protocol of the layer7 listener, and the available values include ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `(Required, ForceNew) ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional, ForceNew) Authentication type of the layer7 listener. ` + "`" + `0` + "`" + ` is one-way authentication and ` + "`" + `1` + "`" + ` is mutual authentication. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) Certificate ID of the layer7 listener. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(Optional) ID of the client certificate. Set only when ` + "`" + `auth_type` + "`" + ` is specified as mutual authentication. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "forward_protocol",
					Description: `(Optional, ForceNew) Protocol type of the forwarding, the available values include ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + `. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` protocol. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer7 listener.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer7 listener.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_proxy",
			Category:         "GAAP Resources",
			ShortDescription: `Provides a resource to create a GAAP proxy.`,
			Description:      ``,
			Keywords: []string{
				"gaap",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_region",
					Description: `(Required, ForceNew) Access region of the GAAP proxy. The available values include ` + "`" + `NorthChina` + "`" + `, ` + "`" + `EastChina` + "`" + `, ` + "`" + `SouthChina` + "`" + `, ` + "`" + `SouthwestChina` + "`" + `, ` + "`" + `Hongkong` + "`" + `, ` + "`" + `SL_TAIWAN` + "`" + `, ` + "`" + `SoutheastAsia` + "`" + `, ` + "`" + `Korea` + "`" + `, ` + "`" + `SL_India` + "`" + `, ` + "`" + `SL_Australia` + "`" + `, ` + "`" + `Europe` + "`" + `, ` + "`" + `SL_UK` + "`" + `, ` + "`" + `SL_SouthAmerica` + "`" + `, ` + "`" + `NorthAmerica` + "`" + `, ` + "`" + `SL_MiddleUSA` + "`" + `, ` + "`" + `Canada` + "`" + `, ` + "`" + `SL_VIET` + "`" + `, ` + "`" + `WestIndia` + "`" + `, ` + "`" + `Thailand` + "`" + `, ` + "`" + `Virginia` + "`" + `, ` + "`" + `Russia` + "`" + `, ` + "`" + `Japan` + "`" + `, ` + "`" + `SL_Indonesia` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) Maximum bandwidth of the GAAP proxy, unit is Mbps. The available values include ` + "`" + `10` + "`" + `,` + "`" + `20` + "`" + `,` + "`" + `50` + "`" + `,` + "`" + `100` + "`" + `,` + "`" + `200` + "`" + `,` + "`" + `500` + "`" + `,` + "`" + `1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "concurrent",
					Description: `(Required) Maximum concurrency of the GAAP proxy, unit is 10k. The available values include ` + "`" + `2` + "`" + `,` + "`" + `5` + "`" + `,` + "`" + `10` + "`" + `,` + "`" + `20` + "`" + `,` + "`" + `30` + "`" + `,` + "`" + `40` + "`" + `,` + "`" + `50` + "`" + `,` + "`" + `60` + "`" + `,` + "`" + `70` + "`" + `,` + "`" + `80` + "`" + `,` + "`" + `90` + "`" + `,` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the GAAP proxy, the maximum length is 30.`,
				},
				resource.Attribute{
					Name:        "realserver_region",
					Description: `(Required, ForceNew) Region of the GAAP realserver. The available values include ` + "`" + `NorthChina` + "`" + `, ` + "`" + `EastChina` + "`" + `, ` + "`" + `SouthChina` + "`" + `, ` + "`" + `SouthwestChina` + "`" + `, ` + "`" + `Hongkong` + "`" + `, ` + "`" + `SL_TAIWAN` + "`" + `, ` + "`" + `SoutheastAsia` + "`" + `, ` + "`" + `Korea` + "`" + `, ` + "`" + `SL_India` + "`" + `, ` + "`" + `SL_Australia` + "`" + `, ` + "`" + `Europe` + "`" + `, ` + "`" + `SL_UK` + "`" + `, ` + "`" + `SL_SouthAmerica` + "`" + `, ` + "`" + `NorthAmerica` + "`" + `, ` + "`" + `SL_MiddleUSA` + "`" + `, ` + "`" + `Canada` + "`" + `, ` + "`" + `SL_VIET` + "`" + `, ` + "`" + `WestIndia` + "`" + `, ` + "`" + `Thailand` + "`" + `, ` + "`" + `Virginia` + "`" + `, ` + "`" + `Russia` + "`" + `, ` + "`" + `Japan` + "`" + `, ` + "`" + `SL_Indonesia` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Indicates whether GAAP proxy is enabled, default is true.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project within the GAAP proxy, '0' means is Default Project.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the GAAP proxy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Access domain of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "forward_ip",
					Description: `Forwarding IP of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Access IP of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "scalable",
					Description: `Indicates whether GAAP proxy can scalable.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "support_protocols",
					Description: `Supported protocols of the GAAP proxy. ## Import GAAP proxy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_proxy.foo link-11112222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Access domain of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "forward_ip",
					Description: `Forwarding IP of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Access IP of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "scalable",
					Description: `Indicates whether GAAP proxy can scalable.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "support_protocols",
					Description: `Supported protocols of the GAAP proxy. ## Import GAAP proxy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_proxy.foo link-11112222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_realserver",
			Category:         "GAAP Resources",
			ShortDescription: `Provides a resource to create a GAAP realserver.`,
			Description:      ``,
			Keywords: []string{
				"gaap",
				"realserver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the GAAP realserver, the maximum length is 30.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional, ForceNew) Domain of the GAAP realserver, conflict with ` + "`" + `ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional, ForceNew) IP of the GAAP realserver, conflict with ` + "`" + `domain` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) ID of the project within the GAAP realserver, '0' means is Default Project.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the GAAP realserver. ## Import GAAP realserver can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_realserver.foo rs-4ftghy6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_security_policy",
			Category:         "GAAP Resources",
			ShortDescription: `Provides a resource to create a security policy of GAAP proxy.`,
			Description:      ``,
			Keywords: []string{
				"gaap",
				"security",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(Required, ForceNew) Default policy, the available values includes ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `(Required, ForceNew) ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Indicates whether policy is enable, default is true. ## Import GAAP security policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_security_policy.foo pl-xxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_security_rule",
			Category:         "GAAP Resources",
			ShortDescription: `Provides a resource to create a security policy rule.`,
			Description:      ``,
			Keywords: []string{
				"gaap",
				"security",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(Required, ForceNew) Policy of the rule, the available values includes ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `(Required, ForceNew) A network address block of the request source.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, ForceNew) ID of the security policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the security policy rule. Maximum length is 30.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, ForceNew) Target port. Available values includes ` + "`" + `80` + "`" + `,` + "`" + `80,443` + "`" + `,` + "`" + `3306-20000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional, ForceNew) Protocol of the security policy rule. Default is ` + "`" + `ALL` + "`" + `, the available values includes ` + "`" + `TCP` + "`" + `,` + "`" + `UDP` + "`" + ` and ` + "`" + `ALL` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Name:        "availability_zone",
					Description: `(Required, ForceNew) The available zone that the CVM instance locates at.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required, ForceNew) The Image to use for the instance. Change 'image_id' will case instance destroy and re-created.`,
				},
				resource.Attribute{
					Name:        "allocate_public_ip",
					Description: `(Optional, ForceNew) Associate a public ip address with an instance in a VPC or Classic. Boolean value, Default is false.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `(Optional) Settings for data disk.`,
				},
				resource.Attribute{
					Name:        "disable_monitor_service",
					Description: `(Optional) Disable enhance service for monitor, it is enabled by default. When this options is set, monitor agent won't be installed.`,
				},
				resource.Attribute{
					Name:        "disable_security_service",
					Description: `(Optional) Disable enhance service for security, it is enabled by default. When this options is set, security agent won't be installed.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional, ForceNew) The hostname of CVM.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_period",
					Description: `(Optional) The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `(Optional) When enabled, the CVM instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + ` and ` + "`" + `DISABLE_NOTIFY_AND_MANUAL_RENEW` + "`" + `. NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `SPOTPAID` + "`" + `, The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name of the CVM. The max length of instance_name is 60, and default value is ` + "`" + `Terrafrom-CVM-Instance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional, ForceNew) The type of instance to start.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) Internet charge type of the instance, Valid values are ` + "`" + `BANDWIDTH_PREPAID` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `BANDWIDTH_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `. The default is ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional, ForceNew) Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bit per second). Value range: [0, 100], If this value is not specified, then automatically sets it to 0 Mbps.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional) The key pair to use for the instance, it looks like skey-16jig7tx.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password to an instance. In order to take effect new password, the instance will be restarted after modifying the password.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) The private ip to be assigned to this instance, must be in the provided subnet and available.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The project CVM belongs to, default to 0.`,
				},
				resource.Attribute{
					Name:        "running_flag",
					Description: `(Optional) Set instance to running or stop. Default value is true, the instance will shutdown when flag is false.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) A list of security group ids to associate with.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The id of a VPC subnetwork. If you want to create instances in VPC network, this parameter must be set.`,
				},
				resource.Attribute{
					Name:        "system_disk_id",
					Description: `(Optional) System disk snapshot ID used to initialize the system disk. When system disk type is ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + `, disk id is not supported.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional) Size of the system disk. Value range: [50, 1000], and unit is GB. Default is 50GB.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional, ForceNew) Type of the system disk. Valid values are ` + "`" + `LOCAL_BASIC` + "`" + `, ` + "`" + `LOCAL_SSD` + "`" + `, ` + "`" + `CLOUD_BASIC` + "`" + `, ` + "`" + `CLOUD_SSD` + "`" + ` and ` + "`" + `CLOUD_PREMIUM` + "`" + `, default value is ` + "`" + `CLOUD_BASIC` + "`" + `. NOTE: ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + ` are deprecated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).`,
				},
				resource.Attribute{
					Name:        "user_data_raw",
					Description: `(Optional, ForceNew) The user data to be specified into this instance, plain text. Conflicts with ` + "`" + `user_data` + "`" + `. Limited in 16 KB after encrypted in base64 format.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, ForceNew) The user data to be specified into this instance. Must be encrypted in base64 format and limited in 16 KB.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The id of a VPC network. If you want to create instances in VPC network, this parameter must be set. The ` + "`" + `data_disks` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "data_disk_size",
					Description: `(Required) Size of the system disk. Value range: [50, 16000], and unit is GB.`,
				},
				resource.Attribute{
					Name:        "data_disk_type",
					Description: `(Required) Type of the data disk. Valid values are ` + "`" + `LOCAL_BASIC` + "`" + `, ` + "`" + `LOCAL_SSD` + "`" + `, ` + "`" + `CLOUD_BASIC` + "`" + `, ` + "`" + `CLOUD_SSD` + "`" + ` and ` + "`" + `CLOUD_PREMIUM` + "`" + `. NOTE: ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + ` are deprecated.`,
				},
				resource.Attribute{
					Name:        "data_disk_id",
					Description: `(Optional) Data disk snapshot ID used to initialize the data disk. When data disk type is ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + `, disk id is not supported.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `(Optional) Decides whether the disk is deleted with instance(only applied to cloud disk), default to true. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Current status of the instance.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public ip of the instance. ## Import CVM instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_instance.foo ins-2qol3a80 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Current status of the instance.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public ip of the instance. ## Import CVM instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import tencentcloud_instance.foo ins-2qol3a80 ` + "`" + `` + "`" + `` + "`" + ``,
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
			Type:             "tencentcloud_kubernetes_cluster",
			Category:         "Kubernetes Resources",
			ShortDescription: `Provide a resource to create a kubernetes cluster.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `(Required, ForceNew) A network address block of the cluster. Different from vpc cidr and cidr of other clusters within this vpc. Must be in 10./192.168/172.[16-31] segments.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) Vpc Id of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_deploy_type",
					Description: `(Optional, ForceNew) Deployment type of the cluster, the available values include: 'MANAGED_CLUSTER' and 'INDEPENDENT_CLUSTER', Default is 'MANAGED_CLUSTER'.`,
				},
				resource.Attribute{
					Name:        "cluster_desc",
					Description: `(Optional, ForceNew) Description of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_ipvs",
					Description: `(Optional, ForceNew) Indicates whether ipvs is enabled. Default is true.`,
				},
				resource.Attribute{
					Name:        "cluster_max_pod_num",
					Description: `(Optional, ForceNew) The maximum number of Pods per node in the cluster. Default is 256. Must be a multiple of 16 and large than 32.`,
				},
				resource.Attribute{
					Name:        "cluster_max_service_num",
					Description: `(Optional, ForceNew) The maximum number of services in the cluster. Default is 256. Must be a multiple of 16.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Optional, ForceNew) Name of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_os",
					Description: `(Optional, ForceNew) Operating system of the cluster, the available values include: 'centos7.2x86_64' and 'ubuntu16.04.1 LTSx86_64'. Default is 'ubuntu16.04.1 LTSx86_64'.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `(Optional, ForceNew) Version of the cluster, Default is '1.10.5'.`,
				},
				resource.Attribute{
					Name:        "container_runtime",
					Description: `(Optional, ForceNew) Runtime type of the cluster, the available values include: 'docker' and 'containerd'. Default is 'docker'.`,
				},
				resource.Attribute{
					Name:        "ignore_cluster_cidr_conflict",
					Description: `(Optional, ForceNew) Indicates whether to ignore the cluster cidr conflict error. Default is false.`,
				},
				resource.Attribute{
					Name:        "master_config",
					Description: `(Optional, ForceNew) Deploy the machine configuration information of the 'MASTER_ETCD' service, and create <=7 units for common users.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) Project ID, default value is 0.`,
				},
				resource.Attribute{
					Name:        "worker_config",
					Description: `(Optional, ForceNew) Deploy the machine configuration information of the 'WORKER' service, and create <=20 units for common users. The other 'WORK' service are added by 'tencentcloud_kubernetes_worker'. The ` + "`" + `worker_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) Specified types of CVM instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) Private network ID.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Indicates which availability zone will be used.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Optional, ForceNew) Number of cvm.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional, ForceNew) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "enhanced_monitor_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud monitor service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "enhanced_security_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud security service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional, ForceNew) Name of the CVMs.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) Charge types for network traffic. Available values include TRAFFIC_POSTPAID_BY_HOUR.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional, ForceNew) Max bandwidth of Internet access in Mbps. Default is 0.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `(Optional, ForceNew) ID list of keys.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) Password to access.`,
				},
				resource.Attribute{
					Name:        "public_ip_assigned",
					Description: `(Optional, ForceNew) Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional, ForceNew) Security groups to which a CVM instance belongs.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional, ForceNew) Volume of system disk in GB. Default is 50.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional, ForceNew) Type of a CVM disk, and available values include CLOUD_PREMIUM and CLOUD_SSD. Default is CLOUD_PREMIUM.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, ForceNew) ase64-encoded User Data text, the length limit is 16KB. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is 0.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk，available values: CLOUD_PREMIUM and CLOUD_SSD.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) Data disk snapshot ID. The ` + "`" + `master_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) Specified types of CVM instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) Private network ID.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Indicates which availability zone will be used.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Optional, ForceNew) Number of cvm.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional, ForceNew) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "enhanced_monitor_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud monitor service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "enhanced_security_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud security service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional, ForceNew) Name of the CVMs.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) Charge types for network traffic. Available values include TRAFFIC_POSTPAID_BY_HOUR.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional, ForceNew) Max bandwidth of Internet access in Mbps. Default is 0.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `(Optional, ForceNew) ID list of keys.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) Password to access.`,
				},
				resource.Attribute{
					Name:        "public_ip_assigned",
					Description: `(Optional, ForceNew) Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional, ForceNew) Security groups to which a CVM instance belongs.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional, ForceNew) Volume of system disk in GB. Default is 50.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional, ForceNew) Type of a CVM disk, and available values include CLOUD_PREMIUM and CLOUD_SSD. Default is CLOUD_PREMIUM.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, ForceNew) ase64-encoded User Data text, the length limit is 16KB. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is 0.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk，available values: CLOUD_PREMIUM and CLOUD_SSD.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) Data disk snapshot ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certification_authority",
					Description: `The certificate used for access.`,
				},
				resource.Attribute{
					Name:        "cluster_external_endpoint",
					Description: `External network address to access`,
				},
				resource.Attribute{
					Name:        "cluster_node_num",
					Description: `Number of nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name for access.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password of account.`,
				},
				resource.Attribute{
					Name:        "pgw_endpoint",
					Description: `The Intranet address used for access.`,
				},
				resource.Attribute{
					Name:        "security_policy",
					Description: `Access policy.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `User name of account.`,
				},
				resource.Attribute{
					Name:        "worker_instances_list",
					Description: `An information list of cvm within the 'WORKER' clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "failed_reason",
					Description: `Information of the cvm when it is failed.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the cvm.`,
				},
				resource.Attribute{
					Name:        "instance_role",
					Description: `Role of the cvm.`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `State of the cvm.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "certification_authority",
					Description: `The certificate used for access.`,
				},
				resource.Attribute{
					Name:        "cluster_external_endpoint",
					Description: `External network address to access`,
				},
				resource.Attribute{
					Name:        "cluster_node_num",
					Description: `Number of nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name for access.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password of account.`,
				},
				resource.Attribute{
					Name:        "pgw_endpoint",
					Description: `The Intranet address used for access.`,
				},
				resource.Attribute{
					Name:        "security_policy",
					Description: `Access policy.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `User name of account.`,
				},
				resource.Attribute{
					Name:        "worker_instances_list",
					Description: `An information list of cvm within the 'WORKER' clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "failed_reason",
					Description: `Information of the cvm when it is failed.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the cvm.`,
				},
				resource.Attribute{
					Name:        "instance_role",
					Description: `Role of the cvm.`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `State of the cvm.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_kubernetes_scale_worker",
			Category:         "Kubernetes Resources",
			ShortDescription: `Provide a resource to increase instance to cluster`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"scale",
				"worker",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "worker_config",
					Description: `(Required, ForceNew) Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users. The ` + "`" + `worker_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) Specified types of CVM instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) Private network ID.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Indicates which availability zone will be used.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Optional, ForceNew) Number of cvm.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional, ForceNew) Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "enhanced_monitor_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud monitor service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "enhanced_security_service",
					Description: `(Optional, ForceNew) To specify whether to enable cloud security service. Default is TRUE.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional, ForceNew) Name of the CVMs.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) Charge types for network traffic. Available values include TRAFFIC_POSTPAID_BY_HOUR.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional, ForceNew) Max bandwidth of Internet access in Mbps. Default is 0.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `(Optional, ForceNew) ID list of keys.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) Password to access.`,
				},
				resource.Attribute{
					Name:        "public_ip_assigned",
					Description: `(Optional, ForceNew) Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional, ForceNew) Security groups to which a CVM instance belongs.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `(Optional, ForceNew) Volume of system disk in GB. Default is 50.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `(Optional, ForceNew) Type of a CVM disk, and available values include CLOUD_PREMIUM and CLOUD_SSD. Default is CLOUD_PREMIUM`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional, ForceNew) ase64-encoded User Data text, the length limit is 16KB. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is 0.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk，available values: CLOUD_PREMIUM and CLOUD_SSD.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) Data disk snapshot ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "worker_instances_list",
					Description: `An information list of kubernetes cluster 'WORKER' . Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "failed_reason",
					Description: `Information of the cvm when it is failed.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the cvm`,
				},
				resource.Attribute{
					Name:        "instance_role",
					Description: `Role of the cvm`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `State of the cvm`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "worker_instances_list",
					Description: `An information list of kubernetes cluster 'WORKER' . Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "failed_reason",
					Description: `Information of the cvm when it is failed.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the cvm`,
				},
				resource.Attribute{
					Name:        "instance_role",
					Description: `Role of the cvm`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `State of the cvm`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_lb",
			Category:         "CLB Resources",
			ShortDescription: `Provides a Load Balancer resource.`,
			Description:      ``,
			Keywords: []string{
				"clb",
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
			Type:             "tencentcloud_mongodb_instance",
			Category:         "MongoDB Resources",
			ShortDescription: `Provide a resource to create a Mongodb instance.`,
			Description:      ``,
			Keywords: []string{
				"mongodb",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Required, ForceNew) The available zone of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required, ForceNew) Version of the Mongodb, and available values include MONGO_3_WT, MONGO_3_ROCKS and MONGO_36_WT.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) Name of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Required, ForceNew) Type of Mongodb instance, and available values include GIO and TGIO.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of this Mongodb account.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Disk size.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) ID of the security group.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) ID of the subnet within this VPC. The vaule is required if VpcId is set.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of the VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb instance, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2)`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `IP port of the Mongodb instance. ## Import Mongodb instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_mongodb_instance.mongodb cmgo-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb instance, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2)`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `IP port of the Mongodb instance. ## Import Mongodb instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_mongodb_instance.mongodb cmgo-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mongodb_sharding_instance",
			Category:         "MongoDB Resources",
			ShortDescription: `Provide a resource to create a Mongodb sharding instance.`,
			Description:      ``,
			Keywords: []string{
				"mongodb",
				"sharding",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Required, ForceNew) The available zone of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required, ForceNew) Version of the Mongodb, and available values include MONGO_3_WT, MONGO_3_ROCKS and MONGO_36_WT.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) Name of the Mongodb instance`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Required, ForceNew) Type of Mongodb instance, and available values include GIO and TGIO.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size.`,
				},
				resource.Attribute{
					Name:        "nodes_per_shard",
					Description: `(Required, ForceNew) Number of nodes per shard, at least 3(one master and two slaves).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of this Mongodb account.`,
				},
				resource.Attribute{
					Name:        "shard_quantity",
					Description: `(Required, ForceNew) Number of sharding.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Disk size.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) ID of the security group.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) ID of the subnet within this VPC. The vaule is required if VpcId is set.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of the VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb instance, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2)`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `IP port of the Mongodb instance. ## Import Mongodb sharding instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_mongodb_sharding_instance.mongodb cmgo-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb instance, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2)`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `IP port of the Mongodb instance. ## Import Mongodb sharding instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_mongodb_sharding_instance.mongodb cmgo-41s6jwy4 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) Backup method. Supported values include: 'physical' - physical backup`,
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
			ShortDescription: `Provides a resource to create a NAT gateway.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"nat",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the nat gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the vpc.`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `(Optional) EIP arrays bound to the gateway. The value of at least 1.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The maximum public network output bandwidth of nat gateway (unit: Mbps), the available values include： 20,50,100,200,500,1000,2000,5000. Default is 100.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `(Optional) The upper limit of concurrent connection of nat gateway, the available values include : 1000000,3000000,10000000, Default is 1000000. ## Import NAT gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_nat_gateway.foo nat-1asg3t63 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Required, ForceNew) ID of VPC to which the route table should be associated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of routing table. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `ID list of the subnets associated with this route table. ## Import Vpc routetable instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_route_table.test route_table_id ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `ID list of the subnets associated with this route table. ## Import Vpc routetable instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_route_table.test route_table_id ` + "`" + `` + "`" + `` + "`" + ``,
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
			ShortDescription: `Provides a resource to create security group.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the security group to be queried.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the security group.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) Project ID of the security group. ## Import Security group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_security_group.sglab sg-ey3wmiz1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_security_group_rule",
			Category:         "VPC Resources",
			ShortDescription: `Provides a resource to create security group rule.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"security",
				"group",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy",
					Description: `(Required, ForceNew) Rule policy of security group, the available value include ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required, ForceNew) ID of the security group to be queried.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) Type of the security group rule, the available value include ` + "`" + `ingress` + "`" + ` and ` + "`" + `egress` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `(Optional, ForceNew) An IP address network or segment, and conflict with ` + "`" + `source_sgid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) Description of the security group rule.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional, ForceNew) Type of ip protocol, the available value include ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + ` and ` + "`" + `ICMP` + "`" + `. Default to all types protocol.`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `(Optional, ForceNew) Range of the port. The available value can be one, multiple or one segment. E.g. ` + "`" + `80` + "`" + `, ` + "`" + `80,90` + "`" + ` and ` + "`" + `80-90` + "`" + `. Default to all ports.`,
				},
				resource.Attribute{
					Name:        "source_sgid",
					Description: `(Optional, ForceNew) ID of the nested security group, and conflict with ` + "`" + `cidr_ip` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ssl_certificate",
			Category:         "SSL Resources",
			ShortDescription: `Provides a resource to create a SSL certificate.`,
			Description:      ``,
			Keywords: []string{
				"ssl",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cert",
					Description: `(Required, ForceNew) Content of the SSL certificate. Not allowed newline at the start and end.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) Type of the SSL certificate. Available values includes: ` + "`" + `CA` + "`" + ` and ` + "`" + `SVR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional, ForceNew) Key of the SSL certificate and required when certificate type is ` + "`" + `SVR` + "`" + `. Not allowed newline at the start and end.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) Name of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) Project ID of the SSL certificate. Default is ` + "`" + `0` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "begin_time",
					Description: `Beginning time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Primary domain of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Ending time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "product_zh_name",
					Description: `Certificate authority.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "subject_names",
					Description: `ALL domains included in the SSL certificate. Including the primary domain name. ## Import ssl certificate can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ssl_certificate.cert GjTNRoK7 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "begin_time",
					Description: `Beginning time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Primary domain of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Ending time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "product_zh_name",
					Description: `Certificate authority.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "subject_names",
					Description: `ALL domains included in the SSL certificate. Including the primary domain name. ## Import ssl certificate can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ssl_certificate.cert GjTNRoK7 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) ID of a routing table to which the subnet should be associated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the subnet. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `Indicates whether it is the default VPC for this region. ## Import Vpc subnet instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_subnet.test subnet_id ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Indicates whether it is the default VPC for this region. ## Import Vpc subnet instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_subnet.test subnet_id ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) Indicates whether VPC multicast is enabled. The default value is 'true'.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default VPC for this region. ## Import Vpc instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpc.test vpc-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default VPC for this region. ## Import Vpc instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpc.test vpc-id ` + "`" + `` + "`" + `` + "`" + ``,
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
		"tencentcloud_clb_attachment":             15,
		"tencentcloud_clb_instance":               16,
		"tencentcloud_clb_listener":               17,
		"tencentcloud_clb_listener_rule":          18,
		"tencentcloud_clb_redirection":            19,
		"tencentcloud_container_cluster":          20,
		"tencentcloud_container_cluster_instance": 21,
		"tencentcloud_cos_bucket":                 22,
		"tencentcloud_cos_bucket_object":          23,
		"tencentcloud_dc_gateway":                 24,
		"tencentcloud_dc_gateway_ccn_route":       25,
		"tencentcloud_dcx":                        26,
		"tencentcloud_dnat":                       27,
		"tencentcloud_eip":                        28,
		"tencentcloud_eip_association":            29,
		"tencentcloud_gaap_certificate":           30,
		"tencentcloud_gaap_http_domain":           31,
		"tencentcloud_gaap_http_rule":             32,
		"tencentcloud_gaap_layer4_listener":       33,
		"tencentcloud_gaap_layer7_listener":       34,
		"tencentcloud_gaap_proxy":                 35,
		"tencentcloud_gaap_realserver":            36,
		"tencentcloud_gaap_security_policy":       37,
		"tencentcloud_gaap_security_rule":         38,
		"tencentcloud_instance":                   39,
		"tencentcloud_key_pair":                   40,
		"tencentcloud_kubernetes_cluster":         41,
		"tencentcloud_kubernetes_scale_worker":    42,
		"tencentcloud_lb":                         43,
		"tencentcloud_mongodb_instance":           44,
		"tencentcloud_mongodb_sharding_instance":  45,
		"tencentcloud_mysql_account":              46,
		"tencentcloud_mysql_account_privilege":    47,
		"tencentcloud_mysql_backup_policy":        48,
		"tencentcloud_mysql_instance":             49,
		"tencentcloud_mysql_readonly_instance":    50,
		"tencentcloud_nat_gateway":                51,
		"tencentcloud_redis_backup_config":        52,
		"tencentcloud_redis_instance":             53,
		"tencentcloud_route_entry":                54,
		"tencentcloud_route_table":                55,
		"tencentcloud_route_table_entry":          56,
		"tencentcloud_security_group":             57,
		"tencentcloud_security_group_rule":        58,
		"tencentcloud_ssl_certificate":            59,
		"tencentcloud_subnet":                     60,
		"tencentcloud_vpc":                        61,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
