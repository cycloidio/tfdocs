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
			Category:         "CLB",
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
					Name:        "backends",
					Description: `(Required) list of backend server.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) listener ID.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required, ForceNew) loadbalancer ID.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `(Optional, ForceNew) location ID, only support for layer 7 loadbalancer. The ` + "`" + `backends` + "`" + ` object supports the following:`,
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
					Description: `(Optional) Weight of the backend server. Valid value range: [0-100]. Default to 10. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `The protocol type, http or tcp.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `The protocol type, http or tcp.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_attachment",
			Category:         "Auto Scaling(AS)",
			ShortDescription: `Provides a resource to attach or detach CVM instances to a specified scaling group.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
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
					Description: `(Required, ForceNew) ID of a scaling group. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_lifecycle_hook",
			Category:         "Auto Scaling(AS)",
			ShortDescription: `Provides a resource for an AS (Auto scaling) lifecycle hook.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
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
					Description: `(Optional) For CMQ_TOPIC type, a name of topic must be set. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_notification",
			Category:         "Auto Scaling(AS)",
			ShortDescription: `Provides a resource for an AS (Auto scaling) notification.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
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
					Description: `(Required, ForceNew) ID of a scaling group. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_config",
			Category:         "Auto Scaling(AS)",
			ShortDescription: `Provides a resource to create a configuration for an AS (Auto scaling) instance.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
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
					Description: `(Optional) Charge types for network traffic. Available values include ` + "`" + `BANDWIDTH_PREPAID` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `.`,
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
					Description: `(Optional) Type of a CVM disk, and available values include CLOUD_PREMIUM and CLOUD_SSD. Default is CLOUD_PREMIUM.`,
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
					Description: `(Optional) Types of disk, available values: CLOUD_PREMIUM and CLOUD_SSD.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) Data disk snapshot ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the launch configuration was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current statues of a launch configuration. ## Import AutoScaling Configuration can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_as_scaling_config.scaling_config asc-n32ymck2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the launch configuration was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current statues of a launch configuration. ## Import AutoScaling Configuration can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_as_scaling_config.scaling_config asc-n32ymck2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_group",
			Category:         "Auto Scaling(AS)",
			ShortDescription: `Provides a resource to create a group of AS (Auto scaling) instances.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
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
					Description: `(Optional) Specifies to which project the scaling group belongs.`,
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
					Name:        "tags",
					Description: `(Optional) Tags of a scaling group.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the AS group was created.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Instance number of a scaling group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of a scaling group. ## Import AutoScaling Groups can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_as_scaling_group.scaling_group asg-n32ymck2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the AS group was created.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Instance number of a scaling group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of a scaling group. ## Import AutoScaling Groups can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_as_scaling_group.scaling_group asg-n32ymck2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_policy",
			Category:         "Auto Scaling(AS)",
			ShortDescription: `Provides a resource for an AS (Auto scaling) policy.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
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
					Description: `(Optional) Statistic types, include AVERAGE, MAXIMUM and MINIMUM. Default is AVERAGE. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_schedule",
			Category:         "Auto Scaling(AS)",
			ShortDescription: `Provides a resource for an AS (Auto scaling) schedule.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
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
					Description: `(Optional) The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format. And this argument should be set with end_time together. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_group",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM group.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of CAM group.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Description of the CAM group. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM group. ## Import CAM group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_group.foo 90496 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM group. ## Import CAM group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_group.foo 90496 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_group_membership",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM group membership.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"group",
				"membership",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Id of CAM group.`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `(Required) Id set of the CAM group members. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CAM group membership can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_group_membership.foo 12515263 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CAM group membership can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_group_membership.foo 12515263 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_group_policy_attachment",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM group policy attachment.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"group",
				"policy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required, ForceNew) Id of the attached CAM group.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, ForceNew) Id of the policy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM group policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM group policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'Group' means customer strategy and 'QCS' means preset strategy. ## Import CAM group policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_group_policy_attachment.foo 12515263#26800353 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM group policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM group policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'Group' means customer strategy and 'QCS' means preset strategy. ## Import CAM group policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_group_policy_attachment.foo 12515263#26800353 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_policy",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM policy.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "document",
					Description: `(Required) Document of the CAM policy. The syntax refers to https://intl.cloud.tencent.com/document/product/598/10604. There are some notes when using this para in terraform: 1. The elements in JSON claimed supporting two types as ` + "`" + `string` + "`" + ` and ` + "`" + `array` + "`" + ` only support type ` + "`" + `array` + "`" + `; 2. Terraform does not support the ` + "`" + `root` + "`" + ` syntax, when it appears, it must be replaced with the uin it stands for.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of CAM policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the CAM policy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the policy strategy. 1 means customer strategy and 2 means preset strategy.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM policy. ## Import CAM policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_policy.foo 26655801 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the policy strategy. 1 means customer strategy and 2 means preset strategy.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM policy. ## Import CAM policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_policy.foo 26655801 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_role",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM role.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "document",
					Description: `(Required) Document of the CAM role. The syntax refers to https://intl.cloud.tencent.com/document/product/598/10604. There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as ` + "`" + `string` + "`" + ` and ` + "`" + `array` + "`" + ` only support type ` + "`" + `array` + "`" + `; 2. Terraform does not support the ` + "`" + `root` + "`" + ` syntax, when appears, it must be replaced with the uin it stands for.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of CAM role.`,
				},
				resource.Attribute{
					Name:        "console_login",
					Description: `(Optional, ForceNew) Indicade whether the CAM role can login or not.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the CAM role. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM role.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM role. ## Import CAM role can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_role.foo 4611686018427733635 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM role.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM role. ## Import CAM role can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_role.foo 4611686018427733635 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_role_policy_attachment",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM role policy attachment.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"role",
				"policy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, ForceNew) Id of the policy.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required, ForceNew) Id of the attached CAM role. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM role policy attachment. 1 means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM role policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy. ## Import CAM role policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_role_policy_attachment.foo 4611686018427922725#26800353 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM role policy attachment. 1 means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM role policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy. ## Import CAM role policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_role_policy_attachment.foo 4611686018427922725#26800353 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_saml_provider",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM SAML provider.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"saml",
				"provider",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "meta_data",
					Description: `(Required) The meta data document of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of CAM SAML provider. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "provider_arn",
					Description: `The arn of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM SAML provider. ## Import CAM SAML provider can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_saml_provider.foo cam-SAML-provider-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "provider_arn",
					Description: `The arn of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM SAML provider. ## Import CAM SAML provider can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_saml_provider.foo cam-SAML-provider-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_user",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to manage CAM user.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the CAM user.`,
				},
				resource.Attribute{
					Name:        "console_login",
					Description: `(Optional) Indicate whether the CAM user can login to the web console or not.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(Optional) Country code of the phone number, for example: '86'.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Email of the CAM user.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to force deletes the CAM user. If set false, the API secret key will be checked and failed when exists; otherwise the user will be deleted directly. Default is false.`,
				},
				resource.Attribute{
					Name:        "need_reset_password",
					Description: `(Optional) Indicate whether the CAM user need to reset the password when first logins.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password of the CAM user. Password should be at least 8 characters and no more than 32 characters, includes uppercase letters, lowercase letters, numbers and special characters. Only required when ` + "`" + `console_login` + "`" + ` is true. If not set, a random password will be automatically generated.`,
				},
				resource.Attribute{
					Name:        "phone_num",
					Description: `(Optional) Phone number of the CAM user.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Remark of the CAM user.`,
				},
				resource.Attribute{
					Name:        "use_api",
					Description: `(Optional) Indicate whether to generate the API secret key or not. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "secret_id",
					Description: `Secret ID of the CAM user.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Secret key of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `ID of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uin",
					Description: `Uin of the CAM User. ## Import CAM user can be imported using the user name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_user.foo cam-user-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "secret_id",
					Description: `Secret ID of the CAM user.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Secret key of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `ID of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uin",
					Description: `Uin of the CAM User. ## Import CAM user can be imported using the user name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_user.foo cam-user-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_user_policy_attachment",
			Category:         "Cloud Access Management(CAM)",
			ShortDescription: `Provides a resource to create a CAM user policy attachment.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"access",
				"management",
				"cam",
				"user",
				"policy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, ForceNew) Id of the policy.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required, ForceNew) Id of the attached CAM user. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM user policy attachment. 1 means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM user policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy. ## Import CAM user policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_user_policy_attachment.foo cam-test#26800353 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM user policy attachment. 1 means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM user policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy. ## Import CAM user policy attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cam_user_policy_attachment.foo cam-test#26800353 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_snapshot",
			Category:         "Cloud Block Storage(CBS)",
			ShortDescription: `Provides a resource to create a CBS snapshot.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"block",
				"storage",
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
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_snapshot_policy",
			Category:         "Cloud Block Storage(CBS)",
			ShortDescription: `Provides a snapshot policy resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"block",
				"storage",
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
					Description: `(Optional) Retention days of the snapshot, and the default value is 7. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CBS snapshot policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_snapshot_policy.snapshot_policy asp-jliex1tn ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CBS snapshot policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cbs_snapshot_policy.snapshot_policy asp-jliex1tn ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_snapshot_policy_attachment",
			Category:         "Cloud Block Storage(CBS)",
			ShortDescription: `Provides a CBS snapshot policy attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"block",
				"storage",
				"cbs",
				"snapshot",
				"policy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_policy_id",
					Description: `(Required, ForceNew) ID of CBS snapshot policy.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `(Required, ForceNew) ID of CBS. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_storage",
			Category:         "Cloud Block Storage(CBS)",
			ShortDescription: `Provides a resource to create a CBS.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"block",
				"storage",
				"cbs",
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
					Description: `(Required) Volume of CBS, and unit is GB. If storage type is ` + "`" + `CLOUD_SSD` + "`" + `, the size range is [100, 16000], and the others are [10-16000].`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Required, ForceNew) Type of CBS medium, and available values include CLOUD_BASIC, CLOUD_PREMIUM and CLOUD_SSD.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) The charge type of CBS instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `(Optional, ForceNew) Indicates whether CBS is encrypted.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to delete CBS instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "prepaid_period",
					Description: `(Optional) The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `(Optional) When enabled, the CBS instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + ` and ` + "`" + `DISABLE_NOTIFY_AND_MANUAL_RENEW` + "`" + `. NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_storage_attachment",
			Category:         "Cloud Block Storage(CBS)",
			ShortDescription: `Provides a CBS storage attachment resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"block",
				"storage",
				"cbs",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the CVM instance.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `(Required, ForceNew) ID of the mounted CBS. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ccn",
			Category:         "Cloud Connect Network(CCN)",
			ShortDescription: `Provides a resource to create a CCN instance.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"connect",
				"network",
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
					Name:        "id",
					Description: `ID of the resource.`,
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
					Description: `States of instance. The available value include 'ISOLATED'(arrears) and 'AVAILABLE'. ## Import Ccn instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ccn.test ccn-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Description: `States of instance. The available value include 'ISOLATED'(arrears) and 'AVAILABLE'. ## Import Ccn instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ccn.test ccn-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ccn_attachment",
			Category:         "Cloud Connect Network(CCN)",
			ShortDescription: `Provides a CCN attaching resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"connect",
				"network",
				"ccn",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ccn_id",
					Description: `(Required, ForceNew) ID of the CCN.`,
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
					Description: `(Required, ForceNew) Type of attached instance network, and available values include VPC, DIRECTCONNECT, BMVPC and VPNGW. Note: VPNGW type is only for whitelist customer now.`,
				},
				resource.Attribute{
					Name:        "ccn_uin",
					Description: `(Optional, ForceNew) Uin of the ccn attached. Default is ` + "`" + `` + "`" + `, which means the uin of this account. This parameter is used with case when attaching ccn of other account to the instance of this account. For now only support instance type ` + "`" + `VPC` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ccn_bandwidth_limit",
			Category:         "Cloud Connect Network(CCN)",
			ShortDescription: `Provides a resource to limit CCN bandwidth.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"connect",
				"network",
				"ccn",
				"bandwidth",
				"limit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ccn_id",
					Description: `(Required, ForceNew) ID of the CCN.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required, ForceNew) Limitation of region.`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit",
					Description: `(Optional) Limitation of bandwidth. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cdn_domain",
			Category:         "CDN",
			ShortDescription: `Provides a resource to create a CDN domain.`,
			Description:      ``,
			Keywords: []string{
				"cdn",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) Name of the acceleration domain.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Required) Origin server configuration. It's a list and consist of at most one item.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Required, ForceNew) Service type of Acceleration domain name. Valid values are ` + "`" + `web` + "`" + `, ` + "`" + `download` + "`" + ` and ` + "`" + `media` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "area",
					Description: `(Optional) Domain name acceleration region. Valid values are ` + "`" + `mainland` + "`" + `, ` + "`" + `overseas` + "`" + ` and ` + "`" + `global` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "full_url_cache",
					Description: `(Optional) Whether to enable full-path cache. Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "https_config",
					Description: `(Optional) HTTPS acceleration configuration. It's a list and consist of at most one item.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The project CDN belongs to, default to 0.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of cdn domain. The ` + "`" + `client_certificate_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "certificate_content",
					Description: `(Required) Client Certificate PEM format, requires Base64 encoding. The ` + "`" + `https_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "https_switch",
					Description: `(Required) HTTPS configuration switch. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "client_certificate_config",
					Description: `(Optional) Client certificate configuration information.`,
				},
				resource.Attribute{
					Name:        "http2_switch",
					Description: `(Optional) HTTP2 configuration switch. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `, and default value is ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ocsp_stapling_switch",
					Description: `(Optional) OCSP configuration switch. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `, and default value is ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server_certificate_config",
					Description: `(Optional) Server certificate configuration information.`,
				},
				resource.Attribute{
					Name:        "spdy_switch",
					Description: `(Optional) Spdy configuration switch. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `, and default value is ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verify_client",
					Description: `(Optional) Client certificate authentication feature. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `, and default value is ` + "`" + `off` + "`" + `. The ` + "`" + `origin` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "origin_list",
					Description: `(Required) Master origin server list. Valid values can be ip or domain name. When modifying the origin server, you need to enter the corresponding ` + "`" + `origin_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "origin_type",
					Description: `(Required) Master origin server type. Valid values are ` + "`" + `domain` + "`" + `, ` + "`" + `cos` + "`" + `, ` + "`" + `ip` + "`" + `, ` + "`" + `ipv6` + "`" + ` and ` + "`" + `ip_ipv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_origin_list",
					Description: `(Optional) Backup origin server list. Valid values can be ip or domain name. When modifying the backup origin server, you need to enter the corresponding ` + "`" + `backup_origin_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_origin_type",
					Description: `(Optional) Backup origin server type. Valid values are ` + "`" + `domain` + "`" + ` and ` + "`" + `ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_server_name",
					Description: `(Optional) Host header used when accessing the backup origin server. If left empty, the ServerName of master origin server will be used by default.`,
				},
				resource.Attribute{
					Name:        "cos_private_access",
					Description: `(Optional) When OriginType is COS, you can specify if access to private buckets is allowed. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `, and default value is ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "origin_pull_protocol",
					Description: `(Optional) Origin-pull protocol configuration. Valid values are ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` and ` + "`" + `follow` + "`" + `, and default value is ` + "`" + `http` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `(Optional) Host header used when accessing the master origin server. If left empty, the acceleration domain name will be used by default. The ` + "`" + `server_certificate_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "certificate_content",
					Description: `(Optional) Server certificate information. This is required when uploading an external certificate, which should contain the complete certificate chain.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) Server certificate ID.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `(Optional) Certificate remarks.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional) Server key information. This is required when uploading an external certificate. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `CNAME address of domain name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of domain name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Acceleration service status. ## Import CDN domain can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cdn_domain.foo xxxx.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `CNAME address of domain name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of domain name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Acceleration service status. ## Import CDN domain can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cdn_domain.foo xxxx.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cfs_access_group",
			Category:         "Cloud File Storage(CFS)",
			ShortDescription: `Provides a resource to create a CFS access group.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"file",
				"storage",
				"cfs",
				"access",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the access group, and max length is 64.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the access group, and max length is 255. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the access group. ## Import CFS access group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cfs_access_group.foo pgroup-7nx89k7l ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the access group. ## Import CFS access group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cfs_access_group.foo pgroup-7nx89k7l ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cfs_access_rule",
			Category:         "Cloud File Storage(CFS)",
			ShortDescription: `Provides a resource to create a CFS access rule.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"file",
				"storage",
				"cfs",
				"access",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_group_id",
					Description: `(Required, ForceNew) ID of a access group.`,
				},
				resource.Attribute{
					Name:        "auth_client_ip",
					Description: `(Required) A single IP or a single IP address range such as 10.1.10.11 or 10.10.1.0/24 indicates that all IPs are allowed. Please note that the IP entered should be CVM's private IP.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) The priority level of rule. The range is 1-100, and 1 indicates the highest priority.`,
				},
				resource.Attribute{
					Name:        "rw_permission",
					Description: `(Optional) Read and write permissions. Valid values are ` + "`" + `RO` + "`" + ` and ` + "`" + `RW` + "`" + `, and default is ` + "`" + `RO` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_permission",
					Description: `(Optional) The permissions of accessing users. Valid values are ` + "`" + `all_squash` + "`" + `, ` + "`" + `no_all_squash` + "`" + `, ` + "`" + `root_squash` + "`" + ` and ` + "`" + `no_root_squash` + "`" + `, and default is ` + "`" + `root_squash` + "`" + `. ` + "`" + `all_squash` + "`" + ` indicates that all access users are mapped as anonymous users or user groups; ` + "`" + `no_all_squash` + "`" + ` indicates that access users will match local users first and be mapped to anonymous users or user groups after matching failed; ` + "`" + `root_squash` + "`" + ` indicates that map access root users to anonymous users or user groups; ` + "`" + `no_root_squash` + "`" + ` indicates that access root users keep root account permission. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cfs_file_system",
			Category:         "Cloud File Storage(CFS)",
			ShortDescription: `Provides a resource to create a cloud file system(CFS).`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"file",
				"storage",
				"cfs",
				"system",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_group_id",
					Description: `(Required) ID of a access group.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) The available zone that the file system locates at.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) ID of a subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of a VPC network.`,
				},
				resource.Attribute{
					Name:        "mount_ip",
					Description: `(Optional, ForceNew) IP of mount point.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of a file system.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional, ForceNew) File service protocol. Valid values are ` + "`" + `NFS` + "`" + ` and ` + "`" + `CIFS` + "`" + `, and the default is ` + "`" + `NFS` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the file system. ## Import Cloud file system can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cfs_file_system.foo cfs-6hgquxmj ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the file system. ## Import Cloud file system can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cfs_file_system.foo cfs-6hgquxmj ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_attachment",
			Category:         "CLB",
			ShortDescription: `Provides a resource to create a CLB attachment.`,
			Description:      ``,
			Keywords: []string{
				"clb",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required, ForceNew) Id of the CLB.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) Id of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) Information of the backends to be attached.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional, ForceNew) Id of the CLB listener rule. Only supports listeners of 'HTTPS' and 'HTTP' protocol. The ` + "`" + `targets` + "`" + ` object supports the following:`,
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
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `Type of protocol within the listener. ## Import CLB attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_attachment.foo loc-4xxr2cy7#lbl-hh141sn9#lb-7a0t6zqb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `Type of protocol within the listener. ## Import CLB attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_attachment.foo loc-4xxr2cy7#lbl-hh141sn9#lb-7a0t6zqb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_instance",
			Category:         "CLB",
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
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "clb_vips",
					Description: `The virtual service address table of the CLB. ## Import CLB instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_instance.foo lb-7a0t6zqb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "clb_vips",
					Description: `The virtual service address table of the CLB. ## Import CLB instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_instance.foo lb-7a0t6zqb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_listener",
			Category:         "CLB",
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
					Description: `(Required, ForceNew) Type of protocol within the listener, and available values are 'TCP', 'UDP', 'HTTP', 'HTTPS' and 'TCP_SSL'.`,
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
					Description: `(Optional) Type of certificate, and available values are 'UNIDIRECTIONAL', 'MUTUAL'. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol and must be set when it is available.`,
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
					Description: `(Optional) Unhealthy threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, the CVM is identified as unhealthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, ForceNew) Port of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling method of the CLB listener, and available values are 'WRR' and 'LEAST_CONN'. The default is 'WRR'. NOTES: The listener of HTTP and 'HTTPS' protocol additionally supports the 'IP Hash' method. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "session_expire_time",
					Description: `(Optional) Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as 'WRR', and not available when listener protocol is 'TCP_SSL'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "sni_switch",
					Description: `(Optional, ForceNew) Indicates whether SNI is enabled, and only supported with protocol 'HTTPS'. If enabled, you can set a certificate for each rule in ` + "`" + `tencentcloud_clb_listener_rule` + "`" + `, otherwise all rules have a certificate. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_listener_rule",
			Category:         "CLB",
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
					Description: `(Optional) HTTP Status Code. The default is 31 and value range is 1-31. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_http_domain",
					Description: `(Optional) Domain name of health check. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_http_method",
					Description: `(Optional) Methods of health check. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol. The default is 'HEAD', the available value are 'HEAD' and 'GET'.`,
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
					Description: `(Optional) Unhealthy threshold of health check, and the default is 3. If the unhealth result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling method of the CLB listener rules, and available values are 'WRR', 'IP HASH' and 'LEAST_CONN'. The default is 'WRR'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "session_expire_time",
					Description: `(Optional) Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as 'WRR', and not available when listener protocol is 'TCP_SSL'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_redirection",
			Category:         "CLB",
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
					Description: `(Optional, ForceNew) Rule id of source listener. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CLB redirection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_redirection.foo loc-ft8fmngv#loc-4xxr2cy7#lbl-jc1dx6ju#lbl-asj1hzuo#lb-p7olt9e5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import CLB redirection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_clb_redirection.foo loc-ft8fmngv#loc-4xxr2cy7#lbl-jc1dx6ju#lbl-asj1hzuo#lb-p7olt9e5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_container_cluster",
			Category:         "Container Cluster",
			ShortDescription: `Provides a TencentCloud Container Cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bandwidth_type",
					Description: `(Required) The network type of the node.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The network bandwidth of the node.`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `(Required) The CIDR which the cluster is going to use.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the cluster.`,
				},
				resource.Attribute{
					Name:        "goods_num",
					Description: `(Required) The node number is going to create in the cluster.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) The instance type of the node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "is_vpc_gateway",
					Description: `(Required) Describe whether the node enable the gateway capability.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `(Required) The system os name of the node.`,
				},
				resource.Attribute{
					Name:        "root_size",
					Description: `(Required) The size of the root volume.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `(Required) The size of the data volume.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The subnet id which the node stays in.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specify vpc which the node(s) stay in.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The zone which the node stays in.`,
				},
				resource.Attribute{
					Name:        "cluster_desc",
					Description: `(Optional) The description of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `(Optional) The kubernetes version of the cluster.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "cvm_type",
					Description: `(Optional) The type of node needed by cvm.`,
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
					Name:        "key_id",
					Description: `(Optional) The key_id of each node(if using key pair to access).`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional) The path which volume is going to be mounted.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password of each node.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The puchase duration of the node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "require_wan_ip",
					Description: `(Optional) Indicate whether wan ip is needed.`,
				},
				resource.Attribute{
					Name:        "root_type",
					Description: `(Optional) The type of the root volume. see more from CVM.`,
				},
				resource.Attribute{
					Name:        "sg_id",
					Description: `(Optional) The security group id.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) The type of the data volume. see more from CVM.`,
				},
				resource.Attribute{
					Name:        "unschedulable",
					Description: `(Optional) Determine whether the node will be schedulable. 0 is the default meaning node will be schedulable. 1 for unschedulable.`,
				},
				resource.Attribute{
					Name:        "user_script",
					Description: `(Optional) User defined script in a base64-format. The script runs after the kubernetes component is ready on node. see more from CCS api documents. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `The kubernetes version of the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_num",
					Description: `The node number of the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_status",
					Description: `The node status of the cluster.`,
				},
				resource.Attribute{
					Name:        "total_cpu",
					Description: `The total cpu of the cluster.`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `The total memory of the cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `The kubernetes version of the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_num",
					Description: `The node number of the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_status",
					Description: `The node status of the cluster.`,
				},
				resource.Attribute{
					Name:        "total_cpu",
					Description: `The total cpu of the cluster.`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `The total memory of the cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_container_cluster_instance",
			Category:         "Container Cluster",
			ShortDescription: `Provides a TencentCloud Container Cluster Instance resource.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"cluster",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bandwidth_type",
					Description: `(Required) The network type of the node.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The network bandwidth of the node.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The id of the cluster.`,
				},
				resource.Attribute{
					Name:        "is_vpc_gateway",
					Description: `(Required) Describe whether the node enable the gateway capability.`,
				},
				resource.Attribute{
					Name:        "root_size",
					Description: `(Required) The size of the root volume.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `(Required) The size of the data volume.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The subnet id which the node stays in.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The zone which the node stays in.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "cvm_type",
					Description: `(Optional) The type of node needed by cvm.`,
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
					Name:        "instance_type",
					Description: `(Optional) The instance type of the node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) The key_id of each node(if using key pair to access).`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "mount_target",
					Description: `(Optional) The path which volume is going to be mounted.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password of each node.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The puchase duration of the node needed by cvm.`,
				},
				resource.Attribute{
					Name:        "require_wan_ip",
					Description: `(Optional) Indicate whether wan ip is needed.`,
				},
				resource.Attribute{
					Name:        "root_type",
					Description: `(Optional) The type of the root volume. see more from CVM.`,
				},
				resource.Attribute{
					Name:        "sg_id",
					Description: `(Optional) The security group id.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) The type of the data volume. see more from CVM.`,
				},
				resource.Attribute{
					Name:        "unschedulable",
					Description: `(Optional) Determine whether the node will be schedulable. 0 is the default meaning node will be schedulable. 1 for unschedulable.`,
				},
				resource.Attribute{
					Name:        "user_script",
					Description: `(Optional) User defined script in a base64-format. The script runs after the kubernetes component is ready on node. see more from CCS api documents. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "lan_ip",
					Description: `Describe the lan ip of the node.`,
				},
				resource.Attribute{
					Name:        "wan_ip",
					Description: `Describe the wan ip of the node.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "lan_ip",
					Description: `Describe the lan ip of the node.`,
				},
				resource.Attribute{
					Name:        "wan_ip",
					Description: `Describe the wan ip of the node.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cos_bucket",
			Category:         "COS",
			ShortDescription: `Provides a COS resource to create a COS bucket and set its attributes.`,
			Description:      ``,
			Keywords: []string{
				"cos",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required, ForceNew) The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example ` + "`" + `mycos-1258798060` + "`" + `.`,
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
					Name:        "encryption_algorithm",
					Description: `(Optional) The server-side encryption algorithm to use. Valid value is ` + "`" + `AES256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifecycle_rules",
					Description: `(Optional) A configuration of object lifecycle management (documented below).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of a bucket.`,
				},
				resource.Attribute{
					Name:        "versioning_enable",
					Description: `(Optional) Enable bucket versioning.`,
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
					Description: `(Optional) Specifies time in seconds that browser can cache the response for a preflight request. The ` + "`" + `expiration` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) Specifies the date after which you want the corresponding action to take effect.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Optional) Specifies the number of days after object creation when the specific rule action takes effect. The ` + "`" + `lifecycle_rules` + "`" + ` object supports the following:`,
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
					Description: `(Optional) Specifies the number of days after object creation when the specific rule action takes effect. The ` + "`" + `website` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `(Optional) An absolute path to the document to return in case of a 4XX error.`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `(Optional) COS returns this index document when requests are made to the root domain or any of the subfolders. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import COS bucket can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cos_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import COS bucket can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_cos_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cos_bucket_object",
			Category:         "COS",
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
					Description: `(Required, ForceNew) The name of a bucket to use. Bucket format should be [custom name]-[appid], for example ` + "`" + `mycos-1258798060` + "`" + `.`,
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
					Description: `(Optional) Specifies caching behavior along the request/reply chain. For further details, RFC2616 can be referred.`,
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
					Description: `(Optional) Object storage type, Available values include STANDARD, STANDARD_IA and ARCHIVE. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_cc_http_policy",
			Category:         "Anti-DDoS(Dayu)",
			ShortDescription: `Use this resource to create a dayu CC self-define http policy`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayu",
				"cc",
				"http",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the CC self-define http policy. Length should between 1 and 20.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required, ForceNew) ID of the resource that the CC self-define http policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the CC self-define http policy works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action mode, only valid when ` + "`" + `smode` + "`" + ` is ` + "`" + `matching` + "`" + `. Valid values are ` + "`" + `alg` + "`" + ` and ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `(Optional) Max frequency per minute, only valid when ` + "`" + `smode` + "`" + ` is ` + "`" + `speedlimit` + "`" + `, the valid value ranges from 1 to 10000.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) Ip of the CC self-define http policy, only valid when ` + "`" + `resource_type` + "`" + ` is ` + "`" + `bgp-multip` + "`" + `. The num of list items can only be set one.`,
				},
				resource.Attribute{
					Name:        "rule_list",
					Description: `(Optional) Rule list of the CC self-define http policy, only valid when ` + "`" + `smode` + "`" + ` is ` + "`" + `matching` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "smode",
					Description: `(Optional) Match mode, and valid values are ` + "`" + `matching` + "`" + `, ` + "`" + `speedlimit` + "`" + `. Note: the speed limit type CC self-define policy can only set one.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `(Optional) Indicate the CC self-define http policy takes effect or not. The ` + "`" + `rule_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) Operator of the rule, valid values are ` + "`" + `include` + "`" + `, ` + "`" + `not_include` + "`" + `, ` + "`" + `equal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "skey",
					Description: `(Optional) Key of the rule, valid values are ` + "`" + `host` + "`" + `, ` + "`" + `cgi` + "`" + `, ` + "`" + `ua` + "`" + `, ` + "`" + `referer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Rule value, then length should be less than 31 bytes. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the CC self-define http policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the CC self-define http policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_cc_https_policy",
			Category:         "Anti-DDoS(Dayu)",
			ShortDescription: `Use this resource to create a dayu CC self-define https policy`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayu",
				"cc",
				"https",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) Domain that the CC self-define https policy works for, only valid when ` + "`" + `protocol` + "`" + ` is ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the CC self-define https policy. Length should between 1 and 20.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required, ForceNew) ID of the resource that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the CC self-define https policy works for, valid value is ` + "`" + `bgpip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Required, ForceNew) Rule id of the domain that the CC self-define https policy works for, only valid when ` + "`" + `protocol` + "`" + ` is ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule_list",
					Description: `(Required) Rule list of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action mode. Valid values are ` + "`" + `alg` + "`" + ` and ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `(Optional) Indicate the CC self-define https policy takes effect or not. The ` + "`" + `rule_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) Operator of the rule, valid values are ` + "`" + `include` + "`" + ` and ` + "`" + `equal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "skey",
					Description: `(Required) Key of the rule, valid values are ` + "`" + `cgi` + "`" + `, ` + "`" + `ua` + "`" + ` and ` + "`" + `referer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Rule value, then length should be less than 31 bytes. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `Ip of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the CC self-define https policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `Ip of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the CC self-define https policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_ddos_policy",
			Category:         "Anti-DDoS(Dayu)",
			ShortDescription: `Use this resource to create dayu DDoS policy`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayu",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "drop_options",
					Description: `(Required) Option list of abnormal check of the DDos policy, should set at least one policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the DDoS policy. Length should between 1 and 32.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the DDoS policy works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "black_ips",
					Description: `(Optional) Black ip list.`,
				},
				resource.Attribute{
					Name:        "packet_filters",
					Description: `(Optional) Message filter options list.`,
				},
				resource.Attribute{
					Name:        "port_filters",
					Description: `(Optional) Port limits of abnormal check of the DDos policy.`,
				},
				resource.Attribute{
					Name:        "watermark_filters",
					Description: `(Optional) Watermark policy options, and only support one watermark policy at most.`,
				},
				resource.Attribute{
					Name:        "white_ips",
					Description: `(Optional) White ip list. The ` + "`" + `drop_options` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "bad_conn_threshold",
					Description: `(Required) The number of new connections based on destination IP that trigger suppression of connections, and valid value is range from 0 to 4294967295.`,
				},
				resource.Attribute{
					Name:        "check_sync_conn",
					Description: `(Required) Indicate whether to check null connection or not.`,
				},
				resource.Attribute{
					Name:        "conn_timeout",
					Description: `(Required) Connection timeout of abnormal connection check, and valid value is range from 0 to 65535.`,
				},
				resource.Attribute{
					Name:        "d_conn_limit",
					Description: `(Required) The limit of concurrent connections based on destination IP, and valid value is range from 0 to 4294967295.`,
				},
				resource.Attribute{
					Name:        "d_new_limit",
					Description: `(Required) The limit of new connections based on destination IP, and valid value is range from 0 to 4294967295.`,
				},
				resource.Attribute{
					Name:        "drop_abroad",
					Description: `(Required) Indicate whether to drop abroad traffic or not.`,
				},
				resource.Attribute{
					Name:        "drop_icmp",
					Description: `(Required) Indicate whether to drop ICMP protocol or not.`,
				},
				resource.Attribute{
					Name:        "drop_other",
					Description: `(Required) Indicate whether to drop other protocols(exclude TCP/UDP/ICMP) or not.`,
				},
				resource.Attribute{
					Name:        "drop_tcp",
					Description: `(Required) Indicate whether to drop TCP protocol or not.`,
				},
				resource.Attribute{
					Name:        "drop_udp",
					Description: `(Required) Indicate to drop UDP protocol or not.`,
				},
				resource.Attribute{
					Name:        "icmp_mbps_limit",
					Description: `(Required) The limit of ICMP traffic rate, and valid value is range from 0 to 4294967295(Mbps).`,
				},
				resource.Attribute{
					Name:        "null_conn_enable",
					Description: `(Required) Indicate to enable null connection or not.`,
				},
				resource.Attribute{
					Name:        "other_mbps_limit",
					Description: `(Required) The limit of other protocols(exclude TCP/UDP/ICMP) traffic rate, and valid value is range from 0 to 4294967295(Mbps).`,
				},
				resource.Attribute{
					Name:        "s_conn_limit",
					Description: `(Required) The limit of concurrent connections based on source IP, and valid value is range from 0 to 4294967295.`,
				},
				resource.Attribute{
					Name:        "s_new_limit",
					Description: `(Required) The limit of new connections based on source IP, and valid value is range from 0 to 4294967295.`,
				},
				resource.Attribute{
					Name:        "syn_limit",
					Description: `(Required) The limit of syn of abnormal connection check, and valid value is range from 0 to 100.`,
				},
				resource.Attribute{
					Name:        "tcp_mbps_limit",
					Description: `(Required) The limit of TCP traffic, and valid value is range from 0 to 4294967295(Mbps).`,
				},
				resource.Attribute{
					Name:        "udp_mbps_limit",
					Description: `(Required) The limit of UDP traffic rate, and valid value is range from 0 to 4294967295(Mbps).`,
				},
				resource.Attribute{
					Name:        "syn_rate",
					Description: `(Optional) The percentage of syn in ack of abnormal connection check, and valid value is range from 0 to 100. The ` + "`" + `packet_filters` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action of port to take, valid values area ` + "`" + `drop` + "`" + `(drop the packet), ` + "`" + `drop_black` + "`" + `(drop the packet and black the ip),` + "`" + `drop_rst` + "`" + `(drop the packet and disconnect),` + "`" + `drop_black_rst` + "`" + `(drop the packet, black the ip and disconnect),` + "`" + `transmit` + "`" + `(transmit the packet).`,
				},
				resource.Attribute{
					Name:        "d_end_port",
					Description: `(Optional) End port of the destination, valid value is range from 0 to 65535. It must be greater than ` + "`" + `d_start_port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "d_start_port",
					Description: `(Optional) Start port of the destination, valid value is range from 0 to 65535.`,
				},
				resource.Attribute{
					Name:        "depth",
					Description: `(Optional) The depth of match, and valid value is range from 0 to 1500.`,
				},
				resource.Attribute{
					Name:        "is_include",
					Description: `(Optional) Indicate whether to include the key word/regular expression or not.`,
				},
				resource.Attribute{
					Name:        "match_begin",
					Description: `(Optional) Indicate whether to check load or not, ` + "`" + `begin_l5` + "`" + ` means to match and ` + "`" + `no_match` + "`" + ` means not.`,
				},
				resource.Attribute{
					Name:        "match_str",
					Description: `(Optional) The key word or regular expression.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Optional) Match type, valid values are ` + "`" + `sunday` + "`" + ` and ` + "`" + `pcre` + "`" + `, ` + "`" + `sunday` + "`" + ` means key word match while ` + "`" + `pcre` + "`" + ` means regular match.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) The offset of match, and valid value is range from 0 to 1500.`,
				},
				resource.Attribute{
					Name:        "pkt_length_max",
					Description: `(Optional) The max length of the packet, and valid value is range from 0 to 1500(Mbps). It must be greater than ` + "`" + `pkt_length_min` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pkt_length_min",
					Description: `(Optional) The minimum length of the packet, and valid value is range from 0 to 1500(Mbps).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol, valid values are ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "s_end_port",
					Description: `(Optional) End port of the source, valid value is range from 0 to 65535. It must be greater than ` + "`" + `s_start_port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "s_start_port",
					Description: `(Optional) Start port of the source, valid value is range from 0 to 65535. The ` + "`" + `port_filters` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action of port to take, valid values area ` + "`" + `drop` + "`" + `, ` + "`" + `transmit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "end_port",
					Description: `(Optional) End port, valid value is range from 0 to 65535. It must be greater than ` + "`" + `start_port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Optional) The type of forbidden port, and valid values are 0, 1, 2. 0 for destination ports make effect, 1 for source ports make effect. 2 for both destination and source ports.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol, valid values are ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start_port",
					Description: `(Optional) Start port, valid value is range from 0 to 65535. The ` + "`" + `watermark_filters` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "auto_remove",
					Description: `(Optional) Indicate whether to auto-remove the watermark or not.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) The offset of watermark, and valid value is range from 0 to 100.`,
				},
				resource.Attribute{
					Name:        "open_switch",
					Description: `(Optional) Indicate whether to open watermark or not. It muse be set ` + "`" + `true` + "`" + ` when any field of watermark was set.`,
				},
				resource.Attribute{
					Name:        "tcp_port_list",
					Description: `(Optional) Port range of TCP, the format is like ` + "`" + `2000-3000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "udp_port_list",
					Description: `(Optional) Port range of TCP, the format is like ` + "`" + `2000-3000` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of policy.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `Id of policy case that the DDoS policy works for.`,
				},
				resource.Attribute{
					Name:        "watermark_key",
					Description: `Watermark content.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Content of the watermark.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the watermark.`,
				},
				resource.Attribute{
					Name:        "open_switch",
					Description: `Indicate whether to auto-remove the watermark or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of policy.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `Id of policy case that the DDoS policy works for.`,
				},
				resource.Attribute{
					Name:        "watermark_key",
					Description: `Watermark content.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Content of the watermark.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the watermark.`,
				},
				resource.Attribute{
					Name:        "open_switch",
					Description: `Indicate whether to auto-remove the watermark or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_ddos_policy_attachment",
			Category:         "Anti-DDoS(Dayu)",
			ShortDescription: `Provides a resource to create a dayu DDoS policy attachment.`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayu",
				"policy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required, ForceNew) Id of the policy.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required, ForceNew) Id of the attached resource.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the DDoS policy works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + `, ` + "`" + `net` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_ddos_policy_case",
			Category:         "Anti-DDoS(Dayu)",
			ShortDescription: `Use this resource to create dayu DDoS policy case`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayu",
				"policy",
				"case",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_protocols",
					Description: `(Required) App protocol set of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "app_type",
					Description: `(Required) App type of the DDoS policy case, and valid values are ` + "`" + `WEB` + "`" + `, ` + "`" + `GAME` + "`" + `, ` + "`" + `APP` + "`" + ` and ` + "`" + `OTHER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "has_abroad",
					Description: `(Required) Indicate whether the service involves overseas or not, valid values are ` + "`" + `no` + "`" + ` and ` + "`" + `yes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "has_initiate_tcp",
					Description: `(Required) Indicate whether the service actively initiates TCP requests or not, valid values are ` + "`" + `no` + "`" + ` and ` + "`" + `yes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the DDoS policy case. Length should between 1 and 64.`,
				},
				resource.Attribute{
					Name:        "platform_types",
					Description: `(Required) Platform set of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the DDoS policy case works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + ` and ` + "`" + `bgp-multip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tcp_end_port",
					Description: `(Required) End port of the TCP service, valid value is range from 0 to 65535. It must be greater than ` + "`" + `tcp_start_port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tcp_start_port",
					Description: `(Required) Start port of the TCP service, valid value is range from 0 to 65535.`,
				},
				resource.Attribute{
					Name:        "udp_end_port",
					Description: `(Required) End port of the UDP service, valid value is range from 0 to 65535. It must be greater than ` + "`" + `udp_start_port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "udp_start_port",
					Description: `(Required) Start port of the UDP service, valid value is range from 0 to 65535.`,
				},
				resource.Attribute{
					Name:        "web_api_urls",
					Description: `(Required) Web API url set.`,
				},
				resource.Attribute{
					Name:        "has_initiate_udp",
					Description: `(Optional) Indicate whether the actively initiate UDP requests or not, valid values are ` + "`" + `no` + "`" + ` and ` + "`" + `yes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "has_vpn",
					Description: `(Optional) Indicate whether the service involves VPN service or not, valid values are ` + "`" + `no` + "`" + ` and ` + "`" + `yes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_tcp_package_len",
					Description: `(Optional) The max length of TCP message package, valid value length should be greater than 0 and less than 1500. It should be greater than ` + "`" + `min_tcp_package_len` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_udp_package_len",
					Description: `(Optional) The max length of UDP message package, valid value length should be greater than 0 and less than 1500. It should be greater than ` + "`" + `min_udp_package_len` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_tcp_package_len",
					Description: `(Optional) The minimum length of TCP message package, valid value length should be greater than 0 and less than 1500.`,
				},
				resource.Attribute{
					Name:        "min_udp_package_len",
					Description: `(Optional) The minimum length of UDP message package, valid value length should be greater than 0 and less than 1500.`,
				},
				resource.Attribute{
					Name:        "peer_tcp_port",
					Description: `(Optional) The port that actively initiates TCP requests, valid value is range from 1 to 65535.`,
				},
				resource.Attribute{
					Name:        "peer_udp_port",
					Description: `(Optional) The port that actively initiates UDP requests, valid value is range from 1 to 65535.`,
				},
				resource.Attribute{
					Name:        "tcp_footprint",
					Description: `(Optional) The fixed signature of TCP protocol load, valid value length is range from 1 to 512.`,
				},
				resource.Attribute{
					Name:        "udp_footprint",
					Description: `(Optional) The fixed signature of TCP protocol load, valid value length is range from 1 to 512. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `Id of the DDoS policy case.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `Id of the DDoS policy case.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_l4_rule",
			Category:         "Anti-DDoS(Dayu)",
			ShortDescription: `Use this resource to create dayu layer 4 rule`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayu",
				"l4",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "d_port",
					Description: `(Required) The destination port of the L4 rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the rule. When the ` + "`" + `resource_type` + "`" + ` is ` + "`" + `net` + "`" + `, this field should be set with valid domain.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol of the rule, valid values are ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `. When ` + "`" + `source_type` + "`" + ` is 1(host source), the value of this field can only set with ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required, ForceNew) ID of the resource that the layer 4 rule works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the layer 4 rule works for, valid values are ` + "`" + `bgpip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "s_port",
					Description: `(Required) The source port of the L4 rule.`,
				},
				resource.Attribute{
					Name:        "source_list",
					Description: `(Required) Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 20.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required, ForceNew) Source type, 1 for source of host, 2 for source of ip.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `(Optional) Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is 2-10.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `(Optional) Interval time of health check. The value range is 10-60 sec, and the default is 15 sec.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `(Optional) Indicates whether health check is enabled. The default is ` + "`" + `false` + "`" + `. Only valid when source list has more than one source item.`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `(Optional) HTTP Status Code. The default is 26 and value range is 2-60.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `(Optional) Unhealthy threshold of health check, and the default is 3. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is 2-10.`,
				},
				resource.Attribute{
					Name:        "session_switch",
					Description: `(Optional) Indicate that the session will keep or not, and default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "session_time",
					Description: `(Optional) Session keep time, only valid when ` + "`" + `session_switch` + "`" + ` is true, the available value ranges from 1 to 300 and unit is second. The ` + "`" + `source_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) Source ip or domain, valid format of ip is like ` + "`" + `1.1.1.1` + "`" + ` and valid format of host source is like ` + "`" + `abc.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Weight of the source, the valid value ranges from 0 to 100. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "lb_type",
					Description: `LB type of the rule, 1 for weight cycling and 2 for IP hash.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of the layer 4 rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "lb_type",
					Description: `LB type of the rule, 1 for weight cycling and 2 for IP hash.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of the layer 4 rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_l7_rule",
			Category:         "Anti-DDoS(Dayu)",
			ShortDescription: `Use this resource to create dayu layer 7 rule`,
			Description:      ``,
			Keywords: []string{
				"anti",
				"ddos",
				"dayu",
				"l7",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol of the rule, valid values are ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required, ForceNew) ID of the resource that the layer 7 rule works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required, ForceNew) Type of the resource that the layer 7 rule works for, valid value is ` + "`" + `bgpip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_list",
					Description: `(Required) Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 16.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Source type, 1 for source of host, 2 for source of ip.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `(Required) Indicate the rule will take effect or not.`,
				},
				resource.Attribute{
					Name:        "health_check_code",
					Description: `(Optional) HTTP Status Code. The default is 26 and value range is 1-31. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `(Optional) Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is 2-10.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `(Optional) Interval time of health check. The value range is 10-60 sec, and the default is 15 sec.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `(Optional) Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `(Optional) Path of health check. The default is ` + "`" + `/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `(Optional) Indicates whether health check is enabled. The default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `(Optional) Unhealthy threshold of health check, and the default is 3. If the unhealth result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is 2-10.`,
				},
				resource.Attribute{
					Name:        "ssl_id",
					Description: `(Optional) SSL id, when the ` + "`" + `protocol` + "`" + ` is ` + "`" + `https` + "`" + `, the field should be set with valid SSL id. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of the layer 7 rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the rule. 0 for create/modify success, 2 for create/modify fail, 3 for delete success, 5 for delete failed, 6 for waiting to be created/modified, 7 for waiting to be deleted and 8 for waiting to get SSL id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of the layer 7 rule.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the rule. 0 for create/modify success, 2 for create/modify fail, 3 for delete success, 5 for delete failed, 6 for waiting to be created/modified, 7 for waiting to be deleted and 8 for waiting to get SSL id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dc_gateway",
			Category:         "Direct Connect Gateway(DCG)",
			ShortDescription: `Provides a resource to creating direct connect gateway instance.`,
			Description:      ``,
			Keywords: []string{
				"direct",
				"connect",
				"gateway",
				"dcg",
				"dc",
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
					Description: `(Optional, ForceNew) Type of the gateway, the available value include 'NORMAL' and 'NAT'. Default is 'NORMAL'. NOTES: CCN only supports 'NORMAL' and a vpc can create two DCGs, the one is NAT type and the other is non-NAT type. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dc_gateway_ccn_route",
			Category:         "Direct Connect Gateway(DCG)",
			ShortDescription: `Provides a resource to creating direct connect gateway route entry.`,
			Description:      ``,
			Keywords: []string{
				"direct",
				"connect",
				"gateway",
				"dcg",
				"dc",
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
					Description: `(Required, ForceNew) ID of the DCG. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "as_path",
					Description: `As_Path list of the BGP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "as_path",
					Description: `As_Path list of the BGP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dcx",
			Category:         "Direct Connect(DC)",
			ShortDescription: `Provides a resource to creating dedicated tunnels instances.`,
			Description:      ``,
			Keywords: []string{
				"direct",
				"connect",
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
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dnat",
			Category:         "VPC",
			ShortDescription: `Provides a resource to create a NAT forwarding.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"dnat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "elastic_ip",
					Description: `(Required, ForceNew) Network address of the EIP.`,
				},
				resource.Attribute{
					Name:        "elastic_port",
					Description: `(Required, ForceNew) Port of the EIP.`,
				},
				resource.Attribute{
					Name:        "nat_id",
					Description: `(Required, ForceNew) Id of the NAT gateway.`,
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
					Description: `(Required, ForceNew) Type of the network protocol, the available values are: ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) Id of the VPC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the NAT forward. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import NAT forwarding can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_dnat.foo tcp://vpc-asg3sfa3:nat-1asg3t63@127.15.2.3:8080 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import NAT forwarding can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_dnat.foo tcp://vpc-asg3sfa3:nat-1asg3t63@127.15.2.3:8080 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eip",
			Category:         "CVM",
			ShortDescription: `Provides an EIP resource.`,
			Description:      ``,
			Keywords: []string{
				"cvm",
				"eip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "anycast_zone",
					Description: `(Optional, ForceNew) The zone of anycast, and available values include ` + "`" + `ANYCAST_ZONE_GLOBAL` + "`" + ` and ` + "`" + `ANYCAST_ZONE_OVERSEAS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "applicable_for_clb",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) The charge type of eip, and available values include ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `, ` + "`" + `BANDWIDTH_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional, ForceNew) The bandwidth limit of eip, unit is Mbps.`,
				},
				resource.Attribute{
					Name:        "internet_service_provider",
					Description: `(Optional, ForceNew) Internet service provider of eip, and available values include ` + "`" + `BGP` + "`" + `, ` + "`" + `CMCC` + "`" + `, ` + "`" + `CTCC` + "`" + ` and ` + "`" + `CUCC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of eip.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of eip.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, ForceNew) The type of eip, and available values include ` + "`" + `EIP` + "`" + ` and ` + "`" + `AnycastEIP` + "`" + `. Default is ` + "`" + `EIP` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The elastic ip address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The eip current status. ## Import EIP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_eip.foo eip-nyvf60va ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The elastic ip address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The eip current status. ## Import EIP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_eip.foo eip-nyvf60va ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eip_association",
			Category:         "CVM",
			ShortDescription: `Provides an eip resource associated with other resource like CVM, ENI and CLB.`,
			Description:      ``,
			Keywords: []string{
				"cvm",
				"eip",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "eip_id",
					Description: `(Required, ForceNew) The id of eip.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional, ForceNew) The CVM or CLB instance id going to bind with the eip. This field is conflict with ` + "`" + `network_interface_id` + "`" + ` and ` + "`" + `private_ip fields` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `(Optional, ForceNew) Indicates the network interface id like ` + "`" + `eni-xxxxxx` + "`" + `. This field is conflict with ` + "`" + `instance_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional, ForceNew) Indicates an IP belongs to the ` + "`" + `network_interface_id` + "`" + `. This field is conflict with ` + "`" + `instance_id` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_elasticsearch_instance",
			Category:         "Elasticsearch",
			ShortDescription: `Provides an elasticsearch instance resource.`,
			Description:      ``,
			Keywords: []string{
				"elasticsearch",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) Availability zone.`,
				},
				resource.Attribute{
					Name:        "node_info_list",
					Description: `(Required, ForceNew) Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password to an instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) The id of a VPC subnetwork.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Version of the instance. Valid values are ` + "`" + `5.6.4` + "`" + `, ` + "`" + `6.4.3` + "`" + `, ` + "`" + `6.8.2` + "`" + ` and ` + "`" + `7.5.1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) The id of a VPC network.`,
				},
				resource.Attribute{
					Name:        "basic_security_type",
					Description: `(Optional) Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are ` + "`" + `1` + "`" + ` and ` + "`" + `2` + "`" + `, ` + "`" + `1` + "`" + ` is disabled, ` + "`" + `2` + "`" + ` is enabled, and default value is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "charge_period",
					Description: `(Optional, ForceNew) The tenancy of the prepaid instance, and uint is month. NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "deploy_mode",
					Description: `(Optional, ForceNew) Cluster deployment mode. Valid values are ` + "`" + `0` + "`" + ` and ` + "`" + `1` + "`" + `, ` + "`" + `0` + "`" + ` is single-AZ deployment, and ` + "`" + `1` + "`" + ` is multi-AZ deployment. Default value is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `(Optional) License type. Valid values are ` + "`" + `oss` + "`" + `, ` + "`" + `basic` + "`" + ` and ` + "`" + `platinum` + "`" + `, and default value is ` + "`" + `platinum` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "multi_zone_infos",
					Description: `(Optional, ForceNew) Details of AZs in multi-AZ deployment mode (which is required when deploy_mode is ` + "`" + `1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "renew_flag",
					Description: `(Optional, ForceNew) When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are ` + "`" + `RENEW_FLAG_AUTO` + "`" + ` and ` + "`" + `RENEW_FLAG_MANUAL` + "`" + `. NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354). The ` + "`" + `multi_zone_infos` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Availability zone.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The id of a VPC subnetwork. The ` + "`" + `node_info_list` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `(Required) Number of nodes.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Required) Node specification, and valid values refer to [document of tencentcloud](https://intl.cloud.tencent.com/document/product/845/18376).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) Node disk size. Unit is GB, and default value is ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional) Node disk type. Valid values are ` + "`" + `CLOUD_SSD` + "`" + ` and ` + "`" + `CLOUD_PREMIUM` + "`" + `, and default value is ` + "`" + `CLOUD_SSD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Node type. Valid values are ` + "`" + `hotData` + "`" + `, ` + "`" + `warmData` + "`" + ` and ` + "`" + `dedicatedMaster` + "`" + `, and default value is 'hotData` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Instance creation time.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_domain",
					Description: `Elasticsearch domain name.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_port",
					Description: `Elasticsearch port.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_vip",
					Description: `Elasticsearch VIP.`,
				},
				resource.Attribute{
					Name:        "kibana_url",
					Description: `Kibana access URL. ## Import Elasticsearch instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_elasticsearch_instance.foo es-17634f05 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Instance creation time.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_domain",
					Description: `Elasticsearch domain name.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_port",
					Description: `Elasticsearch port.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_vip",
					Description: `Elasticsearch VIP.`,
				},
				resource.Attribute{
					Name:        "kibana_url",
					Description: `Kibana access URL. ## Import Elasticsearch instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_elasticsearch_instance.foo es-17634f05 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eni",
			Category:         "VPC",
			ShortDescription: `Provides a resource to create an ENI.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"eni",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the ENI, maximum length 60.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) ID of the subnet within this vpc.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the vpc.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the ENI, maximum length 60.`,
				},
				resource.Attribute{
					Name:        "ipv4_count",
					Description: `(Optional) The number of intranet IPv4s. When it is greater than 1, there is only one primary intranet IP. The others are auxiliary intranet IPs, which conflict with ` + "`" + `ipv4s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv4s",
					Description: `(Optional) Applying for intranet IPv4s collection, conflict with ` + "`" + `ipv4_count` + "`" + `. When there are multiple ipv4s, can only be one primary IP, and the maximum length of the array is 30. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) A set of security group IDs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the ENI. The ` + "`" + `ipv4s` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) Intranet IP.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Required) Indicates whether the IP is primary.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the IP, maximum length 25. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the ENI.`,
				},
				resource.Attribute{
					Name:        "ipv4_info",
					Description: `An information list of IPv4s. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the IP.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Intranet IP.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Indicates whether the IP is primary.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Indicates whether the IP is primary.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the ENI. ## Import ENI can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_eni.foo eni-qka182br ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the ENI.`,
				},
				resource.Attribute{
					Name:        "ipv4_info",
					Description: `An information list of IPv4s. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the IP.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Intranet IP.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Indicates whether the IP is primary.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Indicates whether the IP is primary.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the ENI. ## Import ENI can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_eni.foo eni-qka182br ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eni_attachment",
			Category:         "VPC",
			ShortDescription: `Provides a resource to detailed information of attached backend server to an ENI.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"eni",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "eni_id",
					Description: `(Required, ForceNew) ID of the ENI.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the instance which bind the ENI. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import ENI attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_eni_attachment.foo eni-gtlvkjvz+ins-0h3a5new ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import ENI attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_eni_attachment.foo eni-gtlvkjvz+ins-0h3a5new ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_certificate",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a certificate of GAAP.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
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
					Description: `(Required, ForceNew) Type of the certificate, the available values include ` + "`" + `BASIC` + "`" + `, ` + "`" + `CLIENT` + "`" + `, ` + "`" + `SERVER` + "`" + `, ` + "`" + `REALSERVER` + "`" + ` and ` + "`" + `PROXY` + "`" + `; ` + "`" + `BASIC` + "`" + ` means basic certificate; ` + "`" + `CLIENT` + "`" + ` means client CA certificate; ` + "`" + `SERVER` + "`" + ` means server SSL certificate; ` + "`" + `REALSERVER` + "`" + ` means realserver CA certificate; ` + "`" + `PROXY` + "`" + ` means proxy SSL certificate.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional, ForceNew) Key of the ` + "`" + `SSL` + "`" + ` certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the certificate. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_domain_error_page",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provide a resource to custom error page info for a GAAP HTTP domain.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"domain",
				"error",
				"page",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "body",
					Description: `(Required, ForceNew) New response body.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) HTTP domain.`,
				},
				resource.Attribute{
					Name:        "error_codes",
					Description: `(Required, ForceNew) Original error codes.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "clear_headers",
					Description: `(Optional, ForceNew) Response headers to be removed.`,
				},
				resource.Attribute{
					Name:        "new_error_code",
					Description: `(Optional, ForceNew) New error code.`,
				},
				resource.Attribute{
					Name:        "set_headers",
					Description: `(Optional, ForceNew) Response headers to be set. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_http_domain",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a forward domain of layer7 listener.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
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
					Description: `(Optional) Indicates whether basic authentication is enable, default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) ID of the server certificate, default value is ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "client_certificate_ids",
					Description: `(Optional) ID list of the poly client certificate.`,
				},
				resource.Attribute{
					Name:        "gaap_auth_id",
					Description: `(Optional) ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "gaap_auth",
					Description: `(Optional) Indicates whether SSL certificate authentication is enable, default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "realserver_auth",
					Description: `(Optional) Indicates whether realserver authentication is enable, default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_domain",
					Description: `(Optional) CA certificate domain of the realserver.`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_id",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_ids",
					Description: `(Optional) CA certificate ID list of the realserver. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP http domain can be imported using the id, e.g. ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP http domain can be imported using the id, e.g. ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_http_rule",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a forward rule of layer7 listener.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"http",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required, ForceNew) Forward domain of the forward rule.`,
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
					Description: `(Required, ForceNew) Type of the realserver, the available values include ` + "`" + `IP` + "`" + ` and ` + "`" + `DOMAIN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "realservers",
					Description: `(Required) An information list of GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `(Optional) Timeout of the health check response, default value is 2s.`,
				},
				resource.Attribute{
					Name:        "forward_host",
					Description: `(Optional) The default value of requested host which is forwarded to the realserver by the listener is ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `(Optional) Method of the health check, the available values includes ` + "`" + `GET` + "`" + ` and ` + "`" + `HEAD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `(Optional) Path of health check. Maximum length is 80.`,
				},
				resource.Attribute{
					Name:        "health_check_status_codes",
					Description: `(Optional) Return code of confirmed normal, the available values include ` + "`" + `100` + "`" + `, ` + "`" + `200` + "`" + `, ` + "`" + `300` + "`" + `, ` + "`" + `400` + "`" + ` and ` + "`" + `500` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Interval of the health check, default value is 5s.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling policy of the forward rule, default value is ` + "`" + `rr` + "`" + `, the available values include ` + "`" + `rr` + "`" + `, ` + "`" + `wrr` + "`" + ` and ` + "`" + `lc` + "`" + `. The ` + "`" + `realservers` + "`" + ` object supports the following:`,
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
					Description: `(Optional) Scheduling weight, default value is ` + "`" + `1` + "`" + `. The range of values is [1,100]. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP http rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_http_rule.foo rule-3bsuu01r ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP http rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_http_rule.foo rule-3bsuu01r ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_layer4_listener",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a layer4 listener of GAAP.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
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
					Description: `(Required, ForceNew) Protocol of the layer4 listener, the available values include ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `(Required, ForceNew) ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "realserver_type",
					Description: `(Required, ForceNew) Type of the realserver, the available values include ` + "`" + `IP` + "`" + ` and ` + "`" + `DOMAIN` + "`" + `. NOTES: when the ` + "`" + `protocol` + "`" + ` is specified as ` + "`" + `TCP` + "`" + ` and the ` + "`" + `scheduler` + "`" + ` is specified as ` + "`" + `wrr` + "`" + `, the item can only be set to ` + "`" + `IP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `(Optional) Timeout of the health check response, should less than interval, default value is 2s. NOTES: Only supports listeners of ` + "`" + `TCP` + "`" + ` protocol and require less than ` + "`" + `interval` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Optional) Indicates whether health check is enable, default value is ` + "`" + `false` + "`" + `. NOTES: Only supports listeners of ` + "`" + `TCP` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Interval of the health check, default value is 5s. NOTES: Only supports listeners of ` + "`" + `TCP` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "realserver_bind_set",
					Description: `(Optional) An information list of GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling policy of the layer4 listener, default value is ` + "`" + `rr` + "`" + `, the available values include ` + "`" + `rr` + "`" + `, ` + "`" + `wrr` + "`" + ` and ` + "`" + `lc` + "`" + `. The ` + "`" + `realserver_bind_set` + "`" + ` object supports the following:`,
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
					Description: `(Optional) Scheduling weight, default value is ` + "`" + `1` + "`" + `. The range of values is [1,100]. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer4 listener. ## Import GAAP layer4 listener can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_layer4_listener.foo listener-11112222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer4 listener. ## Import GAAP layer4 listener can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_layer4_listener.foo listener-11112222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_layer7_listener",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a layer7 listener of GAAP.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
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
					Description: `(Required, ForceNew) Protocol of the layer7 listener, the available values include ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + `.`,
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
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "client_certificate_ids",
					Description: `(Optional) ID list of the client certificate. Set only when ` + "`" + `auth_type` + "`" + ` is specified as mutual authentication. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "forward_protocol",
					Description: `(Optional, ForceNew) Protocol type of the forwarding, the available values include ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + `. NOTES: Only supports listeners of ` + "`" + `HTTPS` + "`" + ` protocol. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer7 listener. ## Import GAAP layer7 listener can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_layer7_listener.foo listener-11112222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer7 listener. ## Import GAAP layer7 listener can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_layer7_listener.foo listener-11112222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_proxy",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a GAAP proxy.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_region",
					Description: `(Required, ForceNew) Access region of the GAAP proxy, the available values include ` + "`" + `NorthChina` + "`" + `, ` + "`" + `EastChina` + "`" + `, ` + "`" + `SouthChina` + "`" + `, ` + "`" + `SouthwestChina` + "`" + `, ` + "`" + `Hongkong` + "`" + `, ` + "`" + `SL_TAIWAN` + "`" + `, ` + "`" + `SoutheastAsia` + "`" + `, ` + "`" + `Korea` + "`" + `, ` + "`" + `SL_India` + "`" + `, ` + "`" + `SL_Australia` + "`" + `, ` + "`" + `Europe` + "`" + `, ` + "`" + `SL_UK` + "`" + `, ` + "`" + `SL_SouthAmerica` + "`" + `, ` + "`" + `NorthAmerica` + "`" + `, ` + "`" + `SL_MiddleUSA` + "`" + `, ` + "`" + `Canada` + "`" + `, ` + "`" + `SL_VIET` + "`" + `, ` + "`" + `WestIndia` + "`" + `, ` + "`" + `Thailand` + "`" + `, ` + "`" + `Virginia` + "`" + `, ` + "`" + `Russia` + "`" + `, ` + "`" + `Japan` + "`" + ` and ` + "`" + `SL_Indonesia` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) Maximum bandwidth of the GAAP proxy, unit is Mbps, the available values include ` + "`" + `10` + "`" + `, ` + "`" + `20` + "`" + `, ` + "`" + `50` + "`" + `, ` + "`" + `100` + "`" + `, ` + "`" + `200` + "`" + `, ` + "`" + `500` + "`" + ` and ` + "`" + `1000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "concurrent",
					Description: `(Required) Maximum concurrency of the GAAP proxy, unit is 10k, the available values include ` + "`" + `2` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `20` + "`" + `, ` + "`" + `30` + "`" + `, ` + "`" + `40` + "`" + `, ` + "`" + `50` + "`" + `, ` + "`" + `60` + "`" + `, ` + "`" + `70` + "`" + `, ` + "`" + `80` + "`" + `, ` + "`" + `90` + "`" + ` and ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the GAAP proxy, the maximum length is 30.`,
				},
				resource.Attribute{
					Name:        "realserver_region",
					Description: `(Required, ForceNew) Region of the GAAP realserver, the available values include ` + "`" + `NorthChina` + "`" + `, ` + "`" + `EastChina` + "`" + `, ` + "`" + `SouthChina` + "`" + `, ` + "`" + `SouthwestChina` + "`" + `, ` + "`" + `Hongkong` + "`" + `, ` + "`" + `SL_TAIWAN` + "`" + `, ` + "`" + `SoutheastAsia` + "`" + `, ` + "`" + `Korea` + "`" + `, ` + "`" + `SL_India` + "`" + `, ` + "`" + `SL_Australia` + "`" + `, ` + "`" + `Europe` + "`" + `, ` + "`" + `SL_UK` + "`" + `, ` + "`" + `SL_SouthAmerica` + "`" + `, ` + "`" + `NorthAmerica` + "`" + `, ` + "`" + `SL_MiddleUSA` + "`" + `, ` + "`" + `Canada` + "`" + `, ` + "`" + `SL_VIET` + "`" + `, ` + "`" + `WestIndia` + "`" + `, ` + "`" + `Thailand` + "`" + `, ` + "`" + `Virginia` + "`" + `, ` + "`" + `Russia` + "`" + `, ` + "`" + `Japan` + "`" + ` and ` + "`" + `SL_Indonesia` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Indicates whether GAAP proxy is enabled, default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project within the GAAP proxy, '0' means is default project.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the GAAP proxy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_realserver",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a GAAP realserver.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
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
					Description: `(Optional, ForceNew) ID of the project within the GAAP realserver, '0' means is default project.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the GAAP realserver. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP realserver can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_realserver.foo rs-4ftghy6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP realserver can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_realserver.foo rs-4ftghy6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_security_policy",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a security policy of GAAP proxy.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"security",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(Required, ForceNew) Default policy, the available values include ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `(Required, ForceNew) ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Indicates whether policy is enable, default value is ` + "`" + `true` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP security policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_security_policy.foo pl-xxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP security policy can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_security_policy.foo pl-xxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_security_rule",
			Category:         "Global Application Acceleration(GAAP)",
			ShortDescription: `Provides a resource to create a security policy rule.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"application",
				"acceleration",
				"gaap",
				"security",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(Required, ForceNew) Policy of the rule, the available values include ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `.`,
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
					Description: `(Optional, ForceNew) Target port. Default value is ` + "`" + `ALL` + "`" + `, the available values include ` + "`" + `80` + "`" + `, ` + "`" + `80,443` + "`" + ` and ` + "`" + `3306-20000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional, ForceNew) Protocol of the security policy rule. Default value is ` + "`" + `ALL` + "`" + `, the available values include ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + ` and ` + "`" + `ALL` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP security rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_security_rule.foo sr-xxxxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import GAAP security rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_gaap_security_rule.foo sr-xxxxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ha_vip",
			Category:         "VPC",
			ShortDescription: `Provides a resource to create a HA VIP.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"ha",
				"vip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the HA VIP. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) Subnet id.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) VPC id.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `(Optional, ForceNew) Virtual IP address, it must not be occupied and in this VPC network segment. If not set, it will be assigned after resource created automatically. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "address_ip",
					Description: `EIP that is associated.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the HA VIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Instance id that is associated.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `Network interface id that is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the HA VIP, values are ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `UNBIND` + "`" + `. ## Import HA VIP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ha_vip.foo havip-kjqwe4ba ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "address_ip",
					Description: `EIP that is associated.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the HA VIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Instance id that is associated.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `Network interface id that is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the HA VIP, values are ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `UNBIND` + "`" + `. ## Import HA VIP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ha_vip.foo havip-kjqwe4ba ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ha_vip_eip_attachment",
			Category:         "VPC",
			ShortDescription: `Provides a resource to create a HA VIP EIP attachment.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"ha",
				"vip",
				"eip",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address_ip",
					Description: `(Required, ForceNew) Public address of the EIP.`,
				},
				resource.Attribute{
					Name:        "havip_id",
					Description: `(Required, ForceNew) Id of the attached HA VIP. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import HA VIP EIP attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ha_vip_eip_attachment.foo havip-kjqwe4ba#1.1.1.1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import HA VIP EIP attachment can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_ha_vip_eip_attachment.foo havip-kjqwe4ba#1.1.1.1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_instance",
			Category:         "CVM",
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
					Description: `(Optional, ForceNew) Settings for data disk.`,
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
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to delete instance directly or not. Default is false. If set true, the instance will be permanently deleted instead of staying in recycle bin. Note: only works for ` + "`" + `PREPAID` + "`" + ` instance.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional, ForceNew) The hostname of CVM. Windows instance: The name should be a combination of 2 to 15 characters comprised of letters (case insensitive), numbers, and hyphens (-). Period (.) is not supported, and the name cannot be a string of pure numbers. Other types (such as Linux) of instances: The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-).`,
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
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `SPOTPAID` + "`" + `, The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Note: TencentCloud International only supports ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. ` + "`" + `PREPAID` + "`" + ` instance may not allow to delete before expired. ` + "`" + `SPOTPAID` + "`" + ` instance must set ` + "`" + `spot_instance_type` + "`" + ` and ` + "`" + `spot_max_price` + "`" + ` at the same time.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name of the CVM. The max length of instance_name is 60, and default value is ` + "`" + `Terraform-CVM-Instance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The type of instance to start.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) Internet charge type of the instance, Valid values are ` + "`" + `BANDWIDTH_PREPAID` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `BANDWIDTH_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `. The default is ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional, ForceNew) Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bit per second). If this value is not specified, then automatically sets it to 0 Mbps.`,
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
					Name:        "placement_group_id",
					Description: `(Optional, ForceNew) The id of a placement group.`,
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
					Name:        "spot_instance_type",
					Description: `(Optional) Type of spot instance, only support ` + "`" + `ONE-TIME` + "`" + ` now. Note: it only works when instance_charge_type is set to ` + "`" + `SPOTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_max_price",
					Description: `(Optional, ForceNew) Max price of spot instance, is the format of decimal string, for example "0.50". Note: it only works when instance_charge_type is set to ` + "`" + `SPOTPAID` + "`" + `.`,
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
					Description: `(Optional, ForceNew) Size of the system disk. Value range: [50, 1000], and unit is GB. Default is 50GB.`,
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
					Description: `(Required, ForceNew) Size of the data disk, and unit is GB. If disk type is ` + "`" + `CLOUD_SSD` + "`" + `, the size range is [100, 16000], and the others are [10-16000].`,
				},
				resource.Attribute{
					Name:        "data_disk_type",
					Description: `(Required, ForceNew) Type of the data disk. Valid values are ` + "`" + `LOCAL_BASIC` + "`" + `, ` + "`" + `LOCAL_SSD` + "`" + `, ` + "`" + `CLOUD_BASIC` + "`" + `, ` + "`" + `CLOUD_SSD` + "`" + ` and ` + "`" + `CLOUD_PREMIUM` + "`" + `. NOTE: ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + ` are deprecated.`,
				},
				resource.Attribute{
					Name:        "data_disk_id",
					Description: `(Optional) Data disk snapshot ID used to initialize the data disk. When data disk type is ` + "`" + `LOCAL_BASIC` + "`" + ` and ` + "`" + `LOCAL_SSD` + "`" + `, disk id is not supported.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `(Optional, ForceNew) Decides whether the disk is deleted with instance(only applied to ` + "`" + `CLOUD_BASIC` + "`" + `, ` + "`" + `CLOUD_SSD` + "`" + ` and ` + "`" + `CLOUD_PREMIUM` + "`" + ` disk with ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` instance), default is true. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_key_pair",
			Category:         "CVM",
			ShortDescription: `Provides a key pair resource.`,
			Description:      ``,
			Keywords: []string{
				"cvm",
				"key",
				"pair",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required) The key pair's name. It is the only in one TencentCloud account.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required, ForceNew) You can import an existing public key and using TencentCloud key pair to manage it.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) Specifys to which project the key pair belongs. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Key pair can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_key_pair.foo skey-17634f05 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Key pair can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_key_pair.foo skey-17634f05 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_kubernetes_as_scaling_group",
			Category:         "Kubernetes",
			ShortDescription: `Provide a resource to create an auto scaling group for kubernetes cluster.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"as",
				"scaling",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auto_scaling_config",
					Description: `(Required, ForceNew) Auto scaling config parameters.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group",
					Description: `(Required, ForceNew) Auto scaling group parameters.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, ForceNew) Labels of kubernetes AS Group created nodes. The ` + "`" + `auto_scaling_config` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "configuration_name",
					Description: `(Required, ForceNew) Name of a launch configuration.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required, ForceNew) Specified types of CVM instance.`,
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
					Name:        "instance_tags",
					Description: `(Optional, ForceNew) A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Optional, ForceNew) Charge types for network traffic. Available values include ` + "`" + `BANDWIDTH_PREPAID` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `TRAFFIC_POSTPAID_BY_HOUR` + "`" + ` and ` + "`" + `BANDWIDTH_PACKAGE` + "`" + `.`,
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
					Name:        "project_id",
					Description: `(Optional, ForceNew) Specifys to which project the configuration belongs.`,
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
					Description: `(Optional, ForceNew) Type of a CVM disk, and available values include CLOUD_PREMIUM and CLOUD_SSD. Default is CLOUD_PREMIUM. The ` + "`" + `auto_scaling_group` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required, ForceNew) Maximum number of CVM instances (0~2000).`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Required, ForceNew) Minimum number of CVM instances (0~2000).`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `(Required, ForceNew) Name of a scaling group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of VPC network.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `(Optional, ForceNew) Default cooldown time in second, and default value is 300.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Optional, ForceNew) Desired volume of CVM instances, which is between max_size and min_size.`,
				},
				resource.Attribute{
					Name:        "forward_balancer_ids",
					Description: `(Optional, ForceNew) List of application load balancers, which can't be specified with load_balancer_ids together.`,
				},
				resource.Attribute{
					Name:        "load_balancer_ids",
					Description: `(Optional, ForceNew) ID list of traditional load balancers.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) Specifys to which project the scaling group belongs.`,
				},
				resource.Attribute{
					Name:        "retry_policy",
					Description: `(Optional, ForceNew) Available values for retry policies include IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional, ForceNew) ID list of subnet, and for VPC it is required.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, ForceNew) Tags of a scaling group.`,
				},
				resource.Attribute{
					Name:        "termination_policies",
					Description: `(Optional, ForceNew) Available values for termination policies include OLDEST_INSTANCE and NEWEST_INSTANCE.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `(Optional, ForceNew) List of available zones, for Basic network it is required. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is 0.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk, available values: CLOUD_PREMIUM and CLOUD_SSD.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) Data disk snapshot ID. The ` + "`" + `forward_balancer_ids` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required, ForceNew) Listener ID for application load balancers.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required, ForceNew) ID of available load balancers.`,
				},
				resource.Attribute{
					Name:        "target_attribute",
					Description: `(Required, ForceNew) Attribute list of target rules.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional, ForceNew) ID of forwarding rules. The ` + "`" + `target_attribute` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required, ForceNew) Port number.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required, ForceNew) Weight. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_kubernetes_cluster",
			Category:         "Kubernetes",
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
					Name:        "cluster_internet",
					Description: `(Optional) Open internet access or not.`,
				},
				resource.Attribute{
					Name:        "cluster_intranet_subnet_id",
					Description: `(Optional) Subnet id who can access this independent cluster, this field must and can only set when ` + "`" + `cluster_intranet` + "`" + ` is true. ` + "`" + `cluster_intranet_subnet_id` + "`" + ` can not modify once be set.`,
				},
				resource.Attribute{
					Name:        "cluster_intranet",
					Description: `(Optional) Open intranet access or not.`,
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
					Name:        "cluster_os_type",
					Description: `(Optional, ForceNew) Image type of the cluster os, the available values include: 'DOCKER_CUSTOMIZE','GENERAL'. Default is 'GENERAL'. 'DOCKER_CUSTOMIZE' means 'TKE-Optimized'. Only 'centos7.6x86_64' or 'ubuntu18.04.1 LTSx86_64' support 'DOCKER_CUSTOMIZE' now.`,
				},
				resource.Attribute{
					Name:        "cluster_os",
					Description: `(Optional, ForceNew) Operating system of the cluster, the available values include: 'centos7.2x86_64','centos7.6x86_64','ubuntu16.04.1 LTSx86_64','ubuntu18.04.1 LTSx86_64'. Default is 'ubuntu16.04.1 LTSx86_64'.`,
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
					Name:        "labels",
					Description: `(Optional, ForceNew) Labels of tke cluster nodes.`,
				},
				resource.Attribute{
					Name:        "managed_cluster_internet_security_policies",
					Description: `(Optional) Security policies for managed cluster internet, like:'192.168.1.0/24' or '113.116.51.27', '0.0.0.0/0' means all. This field can only set when field ` + "`" + `cluster_deploy_type` + "`" + ` is 'MANAGED_CLUSTER' and ` + "`" + `cluster_internet` + "`" + ` is true. ` + "`" + `managed_cluster_internet_security_policies` + "`" + ` can not delete or empty once be set.`,
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
					Name:        "tags",
					Description: `(Optional) The tags of the cluster.`,
				},
				resource.Attribute{
					Name:        "worker_config",
					Description: `(Optional, ForceNew) Deploy the machine configuration information of the 'WORKER' service, and create <=20 units for common users. The other 'WORK' service are added by 'tencentcloud_kubernetes_worker'. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is 0.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk, available values: CLOUD_PREMIUM and CLOUD_SSD.`,
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
					Name:        "instance_charge_type_prepaid_period",
					Description: `(Optional, ForceNew) The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `(Optional, ForceNew) When enabled, the CVM instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + ` and ` + "`" + `DISABLE_NOTIFY_AND_MANUAL_RENEW` + "`" + `. NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Note: TencentCloud International only supports ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `PREPAID` + "`" + ` instance will not terminated after cluster deleted, and may not allow to delete before expired.`,
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
					Description: `(Optional, ForceNew) ID list of keys, should be set if ` + "`" + `password` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) Password to access, should be set if ` + "`" + `key_ids` + "`" + ` not set.`,
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
					Description: `(Optional, ForceNew) ase64-encoded User Data text, the length limit is 16KB. The ` + "`" + `worker_config` + "`" + ` object supports the following:`,
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
					Name:        "instance_charge_type_prepaid_period",
					Description: `(Optional, ForceNew) The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `(Optional, ForceNew) When enabled, the CVM instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + ` and ` + "`" + `DISABLE_NOTIFY_AND_MANUAL_RENEW` + "`" + `. NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Note: TencentCloud International only supports ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `PREPAID` + "`" + ` instance will not terminated after cluster deleted, and may not allow to delete before expired.`,
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
					Description: `(Optional, ForceNew) ID list of keys, should be set if ` + "`" + `password` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) Password to access, should be set if ` + "`" + `key_ids` + "`" + ` not set.`,
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
					Description: `(Optional, ForceNew) ase64-encoded User Data text, the length limit is 16KB. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "certification_authority",
					Description: `The certificate used for access.`,
				},
				resource.Attribute{
					Name:        "cluster_external_endpoint",
					Description: `External network address to access.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "certification_authority",
					Description: `The certificate used for access.`,
				},
				resource.Attribute{
					Name:        "cluster_external_endpoint",
					Description: `External network address to access.`,
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
			Type:             "tencentcloud_kubernetes_cluster_attachment",
			Category:         "Kubernetes",
			ShortDescription: `Provide a resource to attach an existing cvm to kubernetes cluster.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"cluster",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, ForceNew) ID of the CVM instance, this cvm will reinstall the system.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `(Optional, ForceNew) The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if ` + "`" + `password` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, ForceNew) Labels of tke attachment exits cvm.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) Password to access, should be set if ` + "`" + `key_ids` + "`" + ` not set. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `A list of security group ids after attach to cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `A list of security group ids after attach to cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_kubernetes_scale_worker",
			Category:         "Kubernetes",
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
					Description: `(Required, ForceNew) Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users. The ` + "`" + `data_disk` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, ForceNew) Volume of disk in GB. Default is 0.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, ForceNew) Types of disk, available values: CLOUD_PREMIUM and CLOUD_SSD.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional, ForceNew) Data disk snapshot ID. The ` + "`" + `worker_config` + "`" + ` object supports the following:`,
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
					Name:        "instance_charge_type_prepaid_period",
					Description: `(Optional, ForceNew) The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `(Optional, ForceNew) When enabled, the CVM instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, ` + "`" + `NOTIFY_AND_MANUAL_RENEW` + "`" + ` and ` + "`" + `DISABLE_NOTIFY_AND_MANUAL_RENEW` + "`" + `. NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, The default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `. Note: TencentCloud International only supports ` + "`" + `POSTPAID_BY_HOUR` + "`" + `, ` + "`" + `PREPAID` + "`" + ` instance will not terminated after cluster deleted, and may not allow to delete before expired.`,
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
					Description: `(Optional, ForceNew) ID list of keys, should be set if ` + "`" + `password` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, ForceNew) Password to access, should be set if ` + "`" + `key_ids` + "`" + ` not set.`,
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
					Description: `(Optional, ForceNew) ase64-encoded User Data text, the length limit is 16KB. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "worker_instances_list",
					Description: `An information list of kubernetes cluster 'WORKER'. Each element contains the following attributes:`,
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
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "worker_instances_list",
					Description: `An information list of kubernetes cluster 'WORKER'. Each element contains the following attributes:`,
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
			Type:             "tencentcloud_lb",
			Category:         "CLB",
			ShortDescription: `Provides a Load Balancer resource.`,
			Description:      ``,
			Keywords: []string{
				"clb",
				"lb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) The network type of the LB, valid choices: 'OPEN', 'INTERNAL'.`,
				},
				resource.Attribute{
					Name:        "forward",
					Description: `(Optional, ForceNew) The type of the LB, valid choices: 'CLASSIC', 'APPLICATION'.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the LB.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) The project id of the LB, unspecified or 0 stands for default project.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) The VPC ID of the LB, unspecified or 0 stands for CVM basic network. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the LB.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the LB.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mongodb_instance",
			Category:         "MongoDB",
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
					Description: `(Required, ForceNew) Version of the Mongodb, and available values include ` + "`" + `MONGO_3_WT` + "`" + `, ` + "`" + `MONGO_3_ROCKS` + "`" + ` and ` + "`" + `MONGO_36_WT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) Name of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Required, ForceNew) Type of Mongodb instance, and available values include ` + "`" + `GIO` + "`" + `(or ` + "`" + `HIO` + "`" + `) and ` + "`" + `TGIO` + "`" + `(or ` + "`" + `HIO10G` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size. The minimum value is 2, and unit is GB.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of this Mongodb account.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Disk size. The minimum value is 25, and unit is GB.`,
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
					Name:        "tags",
					Description: `(Optional) The tags of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of the VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb instance, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).`,
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
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb instance, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).`,
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
			Category:         "MongoDB",
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
					Description: `(Required, ForceNew) Version of the Mongodb, and available values include ` + "`" + `MONGO_3_WT` + "`" + `, ` + "`" + `MONGO_3_ROCKS` + "`" + ` and ` + "`" + `MONGO_36_WT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) Name of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Required, ForceNew) Type of Mongodb instance, and available values include ` + "`" + `GIO` + "`" + `(or ` + "`" + `HIO` + "`" + `) and ` + "`" + `TGIO` + "`" + `(or ` + "`" + `HIO10G` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size. The minimum value is 2, and unit is GB.`,
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
					Description: `(Required) Disk size. The minimum value is 25, and unit is GB.`,
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
					Name:        "tags",
					Description: `(Optional) The tags of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of the VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb instance, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).`,
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
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb instance, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).`,
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
			Type:             "tencentcloud_monitor_binding_object",
			Category:         "Monitor",
			ShortDescription: `Provides a resource for bind objects to a policy group resource.`,
			Description:      ``,
			Keywords: []string{
				"monitor",
				"binding",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Required, ForceNew) A list objects. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required, ForceNew) Policy group id for binding objects. The ` + "`" + `dimensions` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "dimensions_json",
					Description: `(Required, ForceNew) Represents a collection of dimensions of an object instance, json format.eg:'{"unInstanceId":"ins-ot3cq4bi"}'. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_binding_receiver",
			Category:         "Monitor",
			ShortDescription: `Provides a resource for bind receivers to a policy group resource.`,
			Description:      ``,
			Keywords: []string{
				"monitor",
				"binding",
				"receiver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required, ForceNew) Policy group id for binding receivers.`,
				},
				resource.Attribute{
					Name:        "receivers",
					Description: `(Optional) A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the following attributes: The ` + "`" + `receivers` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "notify_way",
					Description: `(Required) Method of warning notification.Optional ` + "`" + `CALL` + "`" + `,` + "`" + `EMAIL` + "`" + `,` + "`" + `SITE` + "`" + `,` + "`" + `SMS` + "`" + `,` + "`" + `WECHAT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "receiver_type",
					Description: `(Required) Receive type. Optional ` + "`" + `group` + "`" + `,` + "`" + `user` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) End of alarm period. Meaning with ` + "`" + `start_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "receive_language",
					Description: `(Optional) Alert sending language. Optional ` + "`" + `en-US` + "`" + `,` + "`" + `zh-CN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "receiver_group_list",
					Description: `(Optional) Alarm receive group id list.`,
				},
				resource.Attribute{
					Name:        "receiver_user_list",
					Description: `(Optional) Alarm receiver id list.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) Alarm period start time.Range [0,86399], which removes the date after it is converted to Beijing time as a Unix timestamp, for example 7200 means '10:0:0'. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_policy_group",
			Category:         "Monitor",
			ShortDescription: `Provides a policy group resource for monitor.`,
			Description:      ``,
			Keywords: []string{
				"monitor",
				"policy",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) Policy group name, length should between 1 and 20.`,
				},
				resource.Attribute{
					Name:        "policy_view_name",
					Description: `(Required, ForceNew) Policy view name, eg:` + "`" + `cvm_device` + "`" + `,` + "`" + `BANDWIDTHPACKAGE` + "`" + `, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(policy_view_name)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Required, ForceNew) Policy group's remark information.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional) A list of threshold rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "event_conditions",
					Description: `(Optional) A list of event rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "is_union_rule",
					Description: `(Optional) The and or relation of indicator alarm rule, 0 represents or rule (if any rule is met, the alarm will be raised), 1 represents and rule (if all rules are met, the alarm will be raised).The default is 0.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) The project id to which the policy group belongs, default is 0. The ` + "`" + `conditions` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "alarm_notify_period",
					Description: `(Required) Alarm sending cycle per second. <0 does not fire, 0 only fires once, and >0 fires every triggerTime second.`,
				},
				resource.Attribute{
					Name:        "alarm_notify_type",
					Description: `(Required) Alarm sending convergence type. 0 continuous alarm, 1 index alarm.`,
				},
				resource.Attribute{
					Name:        "metric_id",
					Description: `(Required) Id of the metric, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(metric_id)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "calc_period",
					Description: `(Optional) Data aggregation cycle (unit of second), if the metric has a default value can not be filled, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(period_keys)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "calc_type",
					Description: `(Optional) Compare type, 1 means more than, 2 means greater than or equal, 3 means less than, 4 means less than or equal to, 5 means equal, 6 means not equal, 7 means days rose, 8 means days fell, 9 means weeks rose, 10 means weeks fell, 11 means period rise, 12 means period fell, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(calc_type_keys)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "calc_value",
					Description: `(Optional) Threshold value, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(calc_value_`,
				},
				resource.Attribute{
					Name:        "continue_period",
					Description: `(Optional) The rule triggers an alert that lasts for several detection cycles, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(period_num_keys)` + "`" + `. The ` + "`" + `event_conditions` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "alarm_notify_period",
					Description: `(Required) Alarm sending cycle per second.<0 does not fire, 0 only fires once, and >0 fires every triggerTime second.`,
				},
				resource.Attribute{
					Name:        "alarm_notify_type",
					Description: `(Required) Alarm sending convergence type. 0 continuous alarm, 1 index alarm.`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `(Required) The id of this event metric, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(event_id). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "binding_objects",
					Description: `A list binding objects(list only those in the ` + "`" + `provider.region` + "`" + `). Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "dimensions_json",
					Description: `Represents a collection of dimensions of an object instance, json format.`,
				},
				resource.Attribute{
					Name:        "is_shielded",
					Description: `Whether the object is shielded or not, 0 means unshielded and 1 means shielded.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the object is located.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `Object unique id.`,
				},
				resource.Attribute{
					Name:        "dimension_group",
					Description: `A list of dimensions for this policy group.`,
				},
				resource.Attribute{
					Name:        "last_edit_uin",
					Description: `Recently edited user uin.`,
				},
				resource.Attribute{
					Name:        "receivers",
					Description: `A list of receivers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `End of alarm period. Meaning with ` + "`" + `start_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "need_send_notice",
					Description: `Do need a telephone alarm contact prompt.You don't need 0, you need 1.`,
				},
				resource.Attribute{
					Name:        "notify_way",
					Description: `Method of warning notification.Optional "SMS", "SITE", "EMAIL", "CALL", "WECHAT".`,
				},
				resource.Attribute{
					Name:        "person_interval",
					Description: `Telephone warning to individual interval (seconds).`,
				},
				resource.Attribute{
					Name:        "receive_language",
					Description: `Alert sending language.`,
				},
				resource.Attribute{
					Name:        "receiver_group_list",
					Description: `Alarm receive group id list.`,
				},
				resource.Attribute{
					Name:        "receiver_type",
					Description: `Receive type, 'group' (receiving group) or 'user' (receiver).`,
				},
				resource.Attribute{
					Name:        "receiver_user_list",
					Description: `Alarm receiver id list.`,
				},
				resource.Attribute{
					Name:        "recover_notify",
					Description: `Restore notification mode. Optional "SMS".`,
				},
				resource.Attribute{
					Name:        "round_interval",
					Description: `Telephone alarm interval per round (seconds).`,
				},
				resource.Attribute{
					Name:        "round_number",
					Description: `Telephone alarm number.`,
				},
				resource.Attribute{
					Name:        "send_for",
					Description: `Telephone warning time.Option "OCCUR","RECOVER".`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Alarm period start time.Range [0,86400], which removes the date after it is converted to Beijing time as a Unix timestamp, for example 7200 means '10:0:0'.`,
				},
				resource.Attribute{
					Name:        "uid_list",
					Description: `The phone alerts the receiver uid.`,
				},
				resource.Attribute{
					Name:        "support_regions",
					Description: `Support regions this policy group.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The policy group update time. ## Import Policy group instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_monitor_policy_group.group group-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "binding_objects",
					Description: `A list binding objects(list only those in the ` + "`" + `provider.region` + "`" + `). Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "dimensions_json",
					Description: `Represents a collection of dimensions of an object instance, json format.`,
				},
				resource.Attribute{
					Name:        "is_shielded",
					Description: `Whether the object is shielded or not, 0 means unshielded and 1 means shielded.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the object is located.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `Object unique id.`,
				},
				resource.Attribute{
					Name:        "dimension_group",
					Description: `A list of dimensions for this policy group.`,
				},
				resource.Attribute{
					Name:        "last_edit_uin",
					Description: `Recently edited user uin.`,
				},
				resource.Attribute{
					Name:        "receivers",
					Description: `A list of receivers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `End of alarm period. Meaning with ` + "`" + `start_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "need_send_notice",
					Description: `Do need a telephone alarm contact prompt.You don't need 0, you need 1.`,
				},
				resource.Attribute{
					Name:        "notify_way",
					Description: `Method of warning notification.Optional "SMS", "SITE", "EMAIL", "CALL", "WECHAT".`,
				},
				resource.Attribute{
					Name:        "person_interval",
					Description: `Telephone warning to individual interval (seconds).`,
				},
				resource.Attribute{
					Name:        "receive_language",
					Description: `Alert sending language.`,
				},
				resource.Attribute{
					Name:        "receiver_group_list",
					Description: `Alarm receive group id list.`,
				},
				resource.Attribute{
					Name:        "receiver_type",
					Description: `Receive type, 'group' (receiving group) or 'user' (receiver).`,
				},
				resource.Attribute{
					Name:        "receiver_user_list",
					Description: `Alarm receiver id list.`,
				},
				resource.Attribute{
					Name:        "recover_notify",
					Description: `Restore notification mode. Optional "SMS".`,
				},
				resource.Attribute{
					Name:        "round_interval",
					Description: `Telephone alarm interval per round (seconds).`,
				},
				resource.Attribute{
					Name:        "round_number",
					Description: `Telephone alarm number.`,
				},
				resource.Attribute{
					Name:        "send_for",
					Description: `Telephone warning time.Option "OCCUR","RECOVER".`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Alarm period start time.Range [0,86400], which removes the date after it is converted to Beijing time as a Unix timestamp, for example 7200 means '10:0:0'.`,
				},
				resource.Attribute{
					Name:        "uid_list",
					Description: `The phone alerts the receiver uid.`,
				},
				resource.Attribute{
					Name:        "support_regions",
					Description: `Support regions this policy group.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The policy group update time. ## Import Policy group instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_monitor_policy_group.group group-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_account",
			Category:         "MySQL",
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
				resource.Attribute{
					Name:        "host",
					Description: `(Optional, ForceNew) Account host, default is ` + "`" + `%` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_account_privilege",
			Category:         "MySQL",
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
					Name:        "account_host",
					Description: `(Optional, ForceNew) Account host, default is ` + "`" + `%` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "privileges",
					Description: `(Optional) Database permissions. Available values for Privileges: "SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "REFERENCES", "INDEX", "ALTER", "CREATE TEMPORARY TABLES", "LOCK TABLES","EXECUTE", "CREATE VIEW", "SHOW VIEW", "CREATE ROUTINE", "ALTER ROUTINE", "EVENT", and "TRIGGER". ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_backup_policy",
			Category:         "MySQL",
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
					Description: `(Optional) Backup method. Supported values include: 'physical' - physical backup.`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `(Optional) Instance backup time, in the format of "HH:mm-HH:mm". Time setting interval is four hours. Default to "02:00-06:00". The following value can be supported: 02:00-06:00, 06:00-10:00, 10:00-14:00, 14:00-18:00, 18:00-22:00, and 22:00-02:00.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `(Optional) Instance backup retention days. Valid values: [7-730]. And default value is 7. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "binlog_period",
					Description: `Retention period for binlog in days.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "binlog_period",
					Description: `Retention period for binlog in days.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_instance",
			Category:         "MySQL",
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
					Name:        "auto_renew_flag",
					Description: `(Optional) Auto renew flag. NOTES: Only supported prepaid instance.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Indicates which availability zone will be used.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) Pay type of instance, valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID` + "`" + `. Default is ` + "`" + `POSTPAID` + "`" + `.`,
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
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to delete instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for ` + "`" + `PREPAID` + "`" + ` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.`,
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
					Name:        "pay_type",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "prepaid_period",
					Description: `(Optional) Period of instance. NOTES: Only supported prepaid instance.`,
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
					Description: `(Optional) ID of VPC, which can be modified once every 24 hours and can't be removed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Description: `Instance status. Available values: 0 - Creating; 1 - Running; 4 - Isolating; 5 - Isolated.`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Indicates which kind of operations is being executed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Description: `Instance status. Available values: 0 - Creating; 1 - Running; 4 - Isolating; 5 - Isolated.`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Indicates which kind of operations is being executed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_privilege",
			Category:         "MySQL",
			ShortDescription: `Provides a mysql account privilege resource to grant different access privilege to different database. A database can be granted by multiple account.`,
			Description:      ``,
			Keywords: []string{
				"mysql",
				"privilege",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required, ForceNew) Account name.the forbidden value is:root,mysql.sys,tencentroot.`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Required) Global privileges. available values for Privileges:SELECT,INSERT,UPDATE,DELETE,CREATE,PROCESS,DROP,REFERENCES,INDEX,ALTER,SHOW DATABASES,CREATE TEMPORARY TABLES,LOCK TABLES,EXECUTE,CREATE VIEW,SHOW VIEW,CREATE ROUTINE,ALTER ROUTINE,EVENT,TRIGGER.`,
				},
				resource.Attribute{
					Name:        "mysql_id",
					Description: `(Required, ForceNew) Instance ID.`,
				},
				resource.Attribute{
					Name:        "account_host",
					Description: `(Optional, ForceNew) Account host, default is ` + "`" + `%` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "column",
					Description: `(Optional) Column privileges list.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Optional) Database privileges list.`,
				},
				resource.Attribute{
					Name:        "table",
					Description: `(Optional) Table privileges list. The ` + "`" + `column` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "column_name",
					Description: `(Required) Column name.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) Database name.`,
				},
				resource.Attribute{
					Name:        "privileges",
					Description: `(Required) Column privilege.available values for Privileges:SELECT,INSERT,UPDATE,REFERENCES.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `(Required) Table name. The ` + "`" + `database` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) Database name.`,
				},
				resource.Attribute{
					Name:        "privileges",
					Description: `(Required) Database privilege.available values for Privileges:SELECT,INSERT,UPDATE,DELETE,CREATE,DROP,REFERENCES,INDEX,ALTER,CREATE TEMPORARY TABLES,LOCK TABLES,EXECUTE,CREATE VIEW,SHOW VIEW,CREATE ROUTINE,ALTER ROUTINE,EVENT,TRIGGER. The ` + "`" + `table` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) Database name.`,
				},
				resource.Attribute{
					Name:        "privileges",
					Description: `(Required) Table privilege.available values for Privileges:SELECT,INSERT,UPDATE,DELETE,CREATE,DROP,REFERENCES,INDEX,ALTER,CREATE VIEW,SHOW VIEW,TRIGGER.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `(Required) Table name. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_readonly_instance",
			Category:         "MySQL",
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
					Name:        "auto_renew_flag",
					Description: `(Optional) Auto renew flag. NOTES: Only supported prepaid instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) Pay type of instance, valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID` + "`" + `. Default is ` + "`" + `POSTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to delete instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for ` + "`" + `PREPAID` + "`" + ` instance. When the main mysql instance set true, this para of the readonly mysql instance will not take effect.`,
				},
				resource.Attribute{
					Name:        "intranet_port",
					Description: `(Optional) Public access port, rang form 1024 to 65535 and default value is 3306.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "prepaid_period",
					Description: `(Optional) Period of instance. NOTES: Only supported prepaid instance.`,
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
					Description: `(Optional) ID of VPC, which can be modified once every 24 hours and can't be removed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Description: `Instance status. Available values: 0 - Creating; 1 - Running; 4 - Isolating; 5 - Isolated.`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Indicates which kind of operations is being executed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Description: `Instance status. Available values: 0 - Creating; 1 - Running; 4 - Isolating; 5 - Isolated.`,
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
			Category:         "VPC",
			ShortDescription: `Provides a resource to create a NAT gateway.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"nat",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `(Required) EIP IP address set bound to the gateway. The value of at least 1 and at most 10.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the vpc.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The maximum public network output bandwidth of NAT gateway (unit: Mbps), the available values include: 20,50,100,200,500,1000,2000,5000. Default is 100.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `(Optional) The upper limit of concurrent connection of NAT gateway, the available values include: 1000000,3000000,10000000. Default is 1000000. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Create time of the NAT gateway. ## Import NAT gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_nat_gateway.foo nat-1asg3t63 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Create time of the NAT gateway. ## Import NAT gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_nat_gateway.foo nat-1asg3t63 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_placement_group",
			Category:         "CVM",
			ShortDescription: `Provide a resource to create a placement group.`,
			Description:      ``,
			Keywords: []string{
				"cvm",
				"placement",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the placement group, 1-60 characters in length.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, ForceNew) Type of the placement group, the available values include ` + "`" + `HOST` + "`" + `,` + "`" + `SW` + "`" + ` and ` + "`" + `RACK` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the placement group.`,
				},
				resource.Attribute{
					Name:        "current_num",
					Description: `Number of hosts in the placement group.`,
				},
				resource.Attribute{
					Name:        "cvm_quota_total",
					Description: `Maximum number of hosts in the placement group. ## Import Placement group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_placement_group.foo ps-ilan8vjf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the placement group.`,
				},
				resource.Attribute{
					Name:        "current_num",
					Description: `Number of hosts in the placement group.`,
				},
				resource.Attribute{
					Name:        "cvm_quota_total",
					Description: `Maximum number of hosts in the placement group. ## Import Placement group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_placement_group.foo ps-ilan8vjf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_postgresql_instance",
			Category:         "PostgreSQL",
			ShortDescription: `Use this resource to create postgresql instance.`,
			Description:      ``,
			Keywords: []string{
				"postgresql",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size(in GB). Allowed value must be larger than ` + "`" + `memory` + "`" + ` that data source ` + "`" + `tencentcloud_postgresql_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `(Required) Password of root account. This parameter can be specified when you purchase master instances, but it should be ignored when you purchase read-only instances or disaster recovery instances.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Required) Volume size(in GB). Allowed value must be a multiple of 10. The storage must be set with the limit of ` + "`" + `storage_min` + "`" + ` and ` + "`" + `storage_max` + "`" + ` which data source ` + "`" + `tencentcloud_postgresql_specinfos` + "`" + ` provides.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, ForceNew) Availability zone.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) Pay type of the postgresql instance. For now, only ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` is valid.`,
				},
				resource.Attribute{
					Name:        "charset",
					Description: `(Optional, ForceNew) Charset of the root account. Valid values are ` + "`" + `UTF8` + "`" + `,` + "`" + `LATIN1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional, ForceNew) Version of the postgresql database engine. Allowed values are ` + "`" + `9.3.5` + "`" + `, ` + "`" + `9.5.4` + "`" + `, ` + "`" + `10.4` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project id, default value is 0.`,
				},
				resource.Attribute{
					Name:        "public_access_switch",
					Description: `(Optional) Indicates whether to enable the access to an instance from public network or not.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) ID of subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "private_access_ip",
					Description: `IP for private access.`,
				},
				resource.Attribute{
					Name:        "private_access_port",
					Description: `Port for private access.`,
				},
				resource.Attribute{
					Name:        "public_access_host",
					Description: `Host for public access.`,
				},
				resource.Attribute{
					Name:        "public_access_port",
					Description: `Port for public access. ## Import postgresql instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_postgresql_instance.foo postgres-cda1iex1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the postgresql instance.`,
				},
				resource.Attribute{
					Name:        "private_access_ip",
					Description: `IP for private access.`,
				},
				resource.Attribute{
					Name:        "private_access_port",
					Description: `Port for private access.`,
				},
				resource.Attribute{
					Name:        "public_access_host",
					Description: `Host for public access.`,
				},
				resource.Attribute{
					Name:        "public_access_port",
					Description: `Port for public access. ## Import postgresql instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_postgresql_instance.foo postgres-cda1iex1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_redis_backup_config",
			Category:         "Redis",
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
					Description: `(Required) Specifys which day the backup action should take place. Supported values include: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday and Sunday.`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `(Required) Specifys what time the backup action should take place.`,
				},
				resource.Attribute{
					Name:        "redis_id",
					Description: `(Required, ForceNew) ID of a Redis instance to which the policy will be applied. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Redis backup config can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_redis_backup_config.redisconfig redis-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Redis backup config can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_redis_backup_config.redisconfig redis-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_redis_instance",
			Category:         "Redis",
			ShortDescription: `Provides a resource to create a Redis instance and set its attributes.`,
			Description:      ``,
			Keywords: []string{
				"redis",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, ForceNew) The available zone ID of an instance to be created, please refer to tencentcloud_redis_zone_config.list.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `(Required) The memory volume of an available instance(in MB), please refer to tencentcloud_redis_zone_config.list[zone].mem_sizes.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password for a Redis user, which should be 8 to 16 characters.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID` + "`" + `. Default value is ` + "`" + `POSTPAID` + "`" + `. Note: TencentCloud International only supports ` + "`" + `POSTPAID` + "`" + `. Caution that update operation on this field will delete old instances and create new with new charge type.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for PREPAID instance.`,
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
					Name:        "prepaid_period",
					Description: `(Optional) The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to ` + "`" + `PREPAID` + "`" + `. Valid values are 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Specifies which project the instance should belong to.`,
				},
				resource.Attribute{
					Name:        "redis_replicas_num",
					Description: `(Optional, ForceNew) The number of instance copies. This is not required for standalone and master slave versions.`,
				},
				resource.Attribute{
					Name:        "redis_shard_num",
					Description: `(Optional, ForceNew) The number of instance shard. This is not required for standalone and master slave versions.`,
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
					Name:        "tags",
					Description: `(Optional) Instance tags.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `(Optional, ForceNew) Instance type. Refer to ` + "`" + `data.tencentcloud_redis_zone_config.list.type_id` + "`" + ` get available values.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, ForceNew,`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of the vpc with which the instance is to be associated. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Description: `Current status of an instance, maybe: init, processing, online, isolate and todelete. ## Import Redis instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_redis_instance.redislab redis-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Description: `Current status of an instance, maybe: init, processing, online, isolate and todelete. ## Import Redis instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_redis_instance.redislab redis-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_reserved_instance",
			Category:         "CVM",
			ShortDescription: `Provides a reserved instance resource.`,
			Description:      ``,
			Keywords: []string{
				"cvm",
				"reserved",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) Configuration id of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `(Required) Number of reserved instances to be purchased. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Expiry time of the RI.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Start time of the RI.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the RI at the time of purchase. ## Import Reserved instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_reserved_instance.foo 6cc16e7c-47d7-4fae-9b44-ce5c0f59a920 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Expiry time of the RI.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Start time of the RI.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the RI at the time of purchase. ## Import Reserved instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_reserved_instance.foo 6cc16e7c-47d7-4fae-9b44-ce5c0f59a920 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_route_entry",
			Category:         "VPC",
			ShortDescription: `Provides a resource to create a routing entry in a VPC routing table.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"route",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, ForceNew) The RouteEntry's target network segment.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `(Required, ForceNew) The route entry's next hub. CVM instance ID or VPC router interface ID.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `(Required, ForceNew) The next hop type. Available value is ` + "`" + `public_gateway` + "`" + `,` + "`" + `vpn_gateway` + "`" + `,` + "`" + `sslvpn_gateway` + "`" + `,` + "`" + `dc_gateway` + "`" + `,` + "`" + `peering_connection` + "`" + `,` + "`" + `nat_gateway` + "`" + ` and ` + "`" + `instance` + "`" + `. ` + "`" + `instance` + "`" + ` points to CVM Instance.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required, ForceNew) The ID of the route table.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) The VPC ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_route_table",
			Category:         "VPC",
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
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_route_table_entry",
			Category:         "VPC",
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
					Description: `(Optional, ForceNew) Description of the routing table entry. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_scf_function",
			Category:         "Serverless Cloud Function(SCF)",
			ShortDescription: `Provide a resource to create a SCF function.`,
			Description:      ``,
			Keywords: []string{
				"serverless",
				"cloud",
				"function",
				"scf",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "handler",
					Description: `(Required) Handler of the SCF function. The format of name is ` + "`" + `<filename>.<method_name>` + "`" + `, and it supports 26 English letters, numbers, connectors, and underscores, it should start with a letter. The last character cannot be ` + "`" + `-` + "`" + ` or ` + "`" + `_` + "`" + `. Available length is 2-60.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, ForceNew) Name of the SCF function. Name supports 26 English letters, numbers, connectors, and underscores, it should start with a letter. The last character cannot be ` + "`" + `-` + "`" + ` or ` + "`" + `_` + "`" + `. Available length is 2-60.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `(Required) Runtime of the SCF function, only supports ` + "`" + `Python2.7` + "`" + `, ` + "`" + `Python3.6` + "`" + `, ` + "`" + `Nodejs6.10` + "`" + `, ` + "`" + `Nodejs8.9` + "`" + `, ` + "`" + `Nodejs10.15` + "`" + `, ` + "`" + `PHP5` + "`" + `, ` + "`" + `PHP7` + "`" + `, ` + "`" + `Golang1` + "`" + `, and ` + "`" + `Java8` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cls_logset_id",
					Description: `(Optional) cls logset id of the SCF function.`,
				},
				resource.Attribute{
					Name:        "cls_topic_id",
					Description: `(Optional) cls topic id of the SCF function.`,
				},
				resource.Attribute{
					Name:        "cos_bucket_name",
					Description: `(Optional) Cos bucket name of the SCF function, such as ` + "`" + `cos-1234567890` + "`" + `, conflict with ` + "`" + `zip_file` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cos_bucket_region",
					Description: `(Optional) Cos bucket region of the SCF function, conflict with ` + "`" + `zip_file` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cos_object_name",
					Description: `(Optional) Cos object name of the SCF function, should have suffix ` + "`" + `.zip` + "`" + ` or ` + "`" + `.jar` + "`" + `, conflict with ` + "`" + `zip_file` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the SCF function. Description supports English letters, numbers, spaces, commas, newlines, periods and Chinese, the maximum length is 1000.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) Environment of the SCF function.`,
				},
				resource.Attribute{
					Name:        "l5_enable",
					Description: `(Optional) Enable L5 for SCF function, default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `(Optional) Memory size of the SCF function, unit is MB. The default is ` + "`" + `128` + "`" + `MB. The range is 128M-1536M, and the ladder is 128M.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional, ForceNew) Namespace of the SCF function, default is ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) Role of the SCF function.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Subnet id of the SCF function.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the SCF function.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout of the SCF function, unit is second. Default ` + "`" + `3` + "`" + `. Available value is 1-900.`,
				},
				resource.Attribute{
					Name:        "triggers",
					Description: `(Optional) Trigger list of the SCF function, note that if you modify the trigger list, all existing triggers will be deleted, and then create triggers in the new list. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) VPC id of the SCF function.`,
				},
				resource.Attribute{
					Name:        "zip_file",
					Description: `(Optional) Zip file of the SCF function, conflict with ` + "`" + `cos_bucket_name` + "`" + `, ` + "`" + `cos_object_name` + "`" + `, ` + "`" + `cos_bucket_region` + "`" + `. The ` + "`" + `triggers` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the SCF function trigger, if ` + "`" + `type` + "`" + ` is ` + "`" + `ckafka` + "`" + `, the format of name must be ` + "`" + `<ckafkaInstanceId>-<topicId>` + "`" + `; if ` + "`" + `type` + "`" + ` is ` + "`" + `cos` + "`" + `, the name is cos bucket id, other In any case, it can be combined arbitrarily. It can only contain English letters, numbers, connectors and underscores. The maximum length is 100.`,
				},
				resource.Attribute{
					Name:        "trigger_desc",
					Description: `(Required) TriggerDesc of the SCF function trigger, parameter format of ` + "`" + `timer` + "`" + ` is linux cron expression; parameter of ` + "`" + `cos` + "`" + ` type is json string ` + "`" + `{"event":"cos:ObjectCreated:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of the SCF function trigger, support ` + "`" + `cos` + "`" + `, ` + "`" + `cmq` + "`" + `, ` + "`" + `timer` + "`" + `, ` + "`" + `ckafka` + "`" + `, ` + "`" + `apigw` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cos_region",
					Description: `(Optional) Region of cos bucket. if ` + "`" + `type` + "`" + ` is ` + "`" + `cos` + "`" + `, ` + "`" + `cos_region` + "`" + ` is required. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "code_error",
					Description: `SCF function code error message.`,
				},
				resource.Attribute{
					Name:        "code_result",
					Description: `SCF function code is correct.`,
				},
				resource.Attribute{
					Name:        "code_size",
					Description: `SCF function code size, unit is M.`,
				},
				resource.Attribute{
					Name:        "eip_fixed",
					Description: `Whether EIP is a fixed IP.`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `SCF function EIP list.`,
				},
				resource.Attribute{
					Name:        "err_no",
					Description: `SCF function code error code.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `SCF function domain name.`,
				},
				resource.Attribute{
					Name:        "install_dependency",
					Description: `Whether to automatically install dependencies.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `SCF function last modified time.`,
				},
				resource.Attribute{
					Name:        "status_desc",
					Description: `SCF status description.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `SCF function status.`,
				},
				resource.Attribute{
					Name:        "trigger_info",
					Description: `SCF trigger details list. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "custom_argument",
					Description: `User-defined parameters of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Whether SCF function trigger is enable.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "trigger_desc",
					Description: `TriggerDesc of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `SCF function vip. ## Import SCF function can be imported, e.g. ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "code_error",
					Description: `SCF function code error message.`,
				},
				resource.Attribute{
					Name:        "code_result",
					Description: `SCF function code is correct.`,
				},
				resource.Attribute{
					Name:        "code_size",
					Description: `SCF function code size, unit is M.`,
				},
				resource.Attribute{
					Name:        "eip_fixed",
					Description: `Whether EIP is a fixed IP.`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `SCF function EIP list.`,
				},
				resource.Attribute{
					Name:        "err_no",
					Description: `SCF function code error code.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `SCF function domain name.`,
				},
				resource.Attribute{
					Name:        "install_dependency",
					Description: `Whether to automatically install dependencies.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `SCF function last modified time.`,
				},
				resource.Attribute{
					Name:        "status_desc",
					Description: `SCF status description.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `SCF function status.`,
				},
				resource.Attribute{
					Name:        "trigger_info",
					Description: `SCF trigger details list. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "custom_argument",
					Description: `User-defined parameters of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Whether SCF function trigger is enable.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "trigger_desc",
					Description: `TriggerDesc of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `SCF function vip. ## Import SCF function can be imported, e.g. ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_scf_namespace",
			Category:         "Serverless Cloud Function(SCF)",
			ShortDescription: `Provide a resource to create a SCF namespace.`,
			Description:      ``,
			Keywords: []string{
				"serverless",
				"cloud",
				"function",
				"scf",
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required, ForceNew) Name of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the SCF namespace. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `SCF namespace creation time.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `SCF namespace last modified time.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `SCF namespace type. ## Import SCF namespace can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_scf_function.test default ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `SCF namespace creation time.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `SCF namespace last modified time.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `SCF namespace type. ## Import SCF namespace can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_scf_function.test default ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_security_group",
			Category:         "VPC",
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
					Description: `(Optional, ForceNew) Project ID of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the security group. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Security group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_security_group.sglab sg-ey3wmiz1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Security group can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_security_group.sglab sg-ey3wmiz1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_security_group_lite_rule",
			Category:         "VPC",
			ShortDescription: `Provide a resource to create security group some lite rules quickly.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"security",
				"group",
				"lite",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required, ForceNew) ID of the security group.`,
				},
				resource.Attribute{
					Name:        "egress",
					Description: `(Optional) Egress rules set. A rule must match the following format: [action]#[cidr_ip]#[port]#[protocol]. The available value of 'action' is ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `. The 'cidr_ip' must be an IP address network or segment. The 'port' valid format is ` + "`" + `80` + "`" + `, ` + "`" + `80,443` + "`" + `, ` + "`" + `80-90` + "`" + ` or ` + "`" + `ALL` + "`" + `. The available value of 'protocol' is ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + ` and ` + "`" + `ALL` + "`" + `. When 'protocol' is ` + "`" + `ICMP` + "`" + ` or ` + "`" + `ALL` + "`" + `, the 'port' must be ` + "`" + `ALL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `(Optional) Ingress rules set. A rule must match the following format: [action]#[cidr_ip]#[port]#[protocol]. The available value of 'action' is ` + "`" + `ACCEPT` + "`" + ` and ` + "`" + `DROP` + "`" + `. The 'cidr_ip' must be an IP address network or segment. The 'port' valid format is ` + "`" + `80` + "`" + `, ` + "`" + `80,443` + "`" + `, ` + "`" + `80-90` + "`" + ` or ` + "`" + `ALL` + "`" + `. The available value of 'protocol' is ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `ICMP` + "`" + ` and ` + "`" + `ALL` + "`" + `. When 'protocol' is ` + "`" + `ICMP` + "`" + ` or ` + "`" + `ALL` + "`" + `, the 'port' must be ` + "`" + `ALL` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Security group lite rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_security_group_lite_rule.foo sg-ey3wmiz1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Security group lite rule can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_security_group_lite_rule.foo sg-ey3wmiz1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_security_group_rule",
			Category:         "VPC",
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
					Description: `(Optional, ForceNew) ID of the nested security group, and conflict with ` + "`" + `cidr_ip` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ssl_certificate",
			Category:         "SSL Certificates",
			ShortDescription: `Provides a resource to create a SSL certificate.`,
			Description:      ``,
			Keywords: []string{
				"ssl",
				"certificates",
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
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_subnet",
			Category:         "VPC",
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
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_cluster",
			Category:         "TcaplusDB",
			ShortDescription: `Use this resource to create TcaplusDB cluster.`,
			Description:      ``,
			Keywords: []string{
				"tcaplusdb",
				"tcaplus",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) Name of the TcaplusDB cluster. Name length should be between 1 and 30.`,
				},
				resource.Attribute{
					Name:        "idl_type",
					Description: `(Required, ForceNew) IDL type of the TcaplusDB cluster. Valid values are PROTO and TDR.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of the TcaplusDB cluster. Password length should be between 12 and 16. The password must be a`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required, ForceNew) Subnet id of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) VPC id of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "old_password_expire_last",
					Description: `(Optional) Expiration time of old password after password update, unit: second. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "api_access_id",
					Description: `Access id of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "api_access_ip",
					Description: `Access ip of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "api_access_port",
					Description: `Access port of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Network type of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "old_password_expire_time",
					Description: `Expiration time of the old password. If ` + "`" + `password_status` + "`" + ` is ` + "`" + `unmodifiable` + "`" + `, it means the old password has not yet expired.`,
				},
				resource.Attribute{
					Name:        "password_status",
					Description: `Password status of the TcaplusDB cluster. Valid values: ` + "`" + `unmodifiable` + "`" + `, which means the password can not be changed in this moment; ` + "`" + `modifiable` + "`" + `, which means the password can be changed in this moment. ## Import tcaplus cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcaplus_cluster.test 26655801 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "api_access_id",
					Description: `Access id of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "api_access_ip",
					Description: `Access ip of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "api_access_port",
					Description: `Access port of the TcaplusDB cluster.For TcaplusDB SDK connect.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Network type of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "old_password_expire_time",
					Description: `Expiration time of the old password. If ` + "`" + `password_status` + "`" + ` is ` + "`" + `unmodifiable` + "`" + `, it means the old password has not yet expired.`,
				},
				resource.Attribute{
					Name:        "password_status",
					Description: `Password status of the TcaplusDB cluster. Valid values: ` + "`" + `unmodifiable` + "`" + `, which means the password can not be changed in this moment; ` + "`" + `modifiable` + "`" + `, which means the password can be changed in this moment. ## Import tcaplus cluster can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_tcaplus_cluster.test 26655801 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_idl",
			Category:         "TcaplusDB",
			ShortDescription: `Use this resource to create TcaplusDB IDL file.`,
			Description:      ``,
			Keywords: []string{
				"tcaplusdb",
				"tcaplus",
				"idl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) Id of the TcaplusDB cluster to which the table group belongs.`,
				},
				resource.Attribute{
					Name:        "file_content",
					Description: `(Required, ForceNew) IDL file content of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "file_ext_type",
					Description: `(Required, ForceNew) File ext type of the IDL file. If ` + "`" + `file_type` + "`" + ` is ` + "`" + `PROTO` + "`" + `, ` + "`" + `file_ext_type` + "`" + ` must be 'proto'; If ` + "`" + `file_type` + "`" + ` is ` + "`" + `TDR` + "`" + `, ` + "`" + `file_ext_type` + "`" + ` must be 'xml'.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `(Required, ForceNew) Name of the IDL file.`,
				},
				resource.Attribute{
					Name:        "file_type",
					Description: `(Required, ForceNew) Type of the IDL file. Valid values are PROTO and TDR.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `(Required, ForceNew) Id of the table group to which the IDL file belongs. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "table_infos",
					Description: `Table info of the IDL.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Error messages for creating IDL file.`,
				},
				resource.Attribute{
					Name:        "index_key_set",
					Description: `Index key set of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "key_fields",
					Description: `Primary key fields of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "sum_key_field_size",
					Description: `Total size of primary key field of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "sum_value_field_size",
					Description: `Total size of non-primary key fields of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `Name of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "value_fields",
					Description: `Non-primary key fields of the TcaplusDB table.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "table_infos",
					Description: `Table info of the IDL.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Error messages for creating IDL file.`,
				},
				resource.Attribute{
					Name:        "index_key_set",
					Description: `Index key set of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "key_fields",
					Description: `Primary key fields of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "sum_key_field_size",
					Description: `Total size of primary key field of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "sum_value_field_size",
					Description: `Total size of non-primary key fields of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `Name of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "value_fields",
					Description: `Non-primary key fields of the TcaplusDB table.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_table",
			Category:         "TcaplusDB",
			ShortDescription: `Use this resource to create TcaplusDB table.`,
			Description:      ``,
			Keywords: []string{
				"tcaplusdb",
				"tcaplus",
				"table",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) ID of the TcaplusDB cluster to which the table belongs.`,
				},
				resource.Attribute{
					Name:        "idl_id",
					Description: `(Required) ID of the IDL File.`,
				},
				resource.Attribute{
					Name:        "reserved_read_cu",
					Description: `(Required, ForceNew) Reserved read capacity units of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "reserved_volume",
					Description: `(Required, ForceNew) Reserved storage capacity of the TcaplusDB table (unit: GB).`,
				},
				resource.Attribute{
					Name:        "reserved_write_cu",
					Description: `(Required, ForceNew) Reserved write capacity units of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_idl_type",
					Description: `(Required) IDL type of the TcaplusDB table. Valid values are PROTO and TDR.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `(Required, ForceNew) Name of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_type",
					Description: `(Required, ForceNew) Type of the TcaplusDB table. Valid values are GENERIC and LIST.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `(Required, ForceNew) ID of the table group to which the table belongs.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the TcaplusDB table. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Error messages for creating TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_size",
					Description: `Size of the TcaplusDB table.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Error messages for creating TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_size",
					Description: `Size of the TcaplusDB table.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_tablegroup",
			Category:         "TcaplusDB",
			ShortDescription: `Use this resource to create TcaplusDB table group.`,
			Description:      ``,
			Keywords: []string{
				"tcaplusdb",
				"tcaplus",
				"tablegroup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, ForceNew) Id of the TcaplusDB cluster to which the table group belongs.`,
				},
				resource.Attribute{
					Name:        "tablegroup_name",
					Description: `(Required) Name of the TcaplusDB table group. Name length should be between 1 and 30. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB table group.`,
				},
				resource.Attribute{
					Name:        "table_count",
					Description: `Number of tables.`,
				},
				resource.Attribute{
					Name:        "total_size",
					Description: `Total storage size (MB).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB table group.`,
				},
				resource.Attribute{
					Name:        "table_count",
					Description: `Number of tables.`,
				},
				resource.Attribute{
					Name:        "total_size",
					Description: `Total storage size (MB).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpc",
			Category:         "VPC",
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
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "id",
					Description: `ID of the resource.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_connection",
			Category:         "VPN",
			ShortDescription: `Provides a resource to create a VPN connection.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `(Required, ForceNew) ID of the customer gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the VPN connection. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "pre_share_key",
					Description: `(Required) Pre-shared key of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "security_group_policy",
					Description: `(Required) Security group policy of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `(Required, ForceNew) ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "ike_dh_group_name",
					Description: `(Optional) DH group name of the IKE operation specification, valid values are ` + "`" + `GROUP1` + "`" + `, ` + "`" + `GROUP2` + "`" + `, ` + "`" + `GROUP5` + "`" + `, ` + "`" + `GROUP14` + "`" + `, ` + "`" + `GROUP24` + "`" + `. Default value is ` + "`" + `GROUP1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_exchange_mode",
					Description: `(Optional) Exchange mode of the IKE operation specification, valid values are ` + "`" + `AGGRESSIVE` + "`" + `, ` + "`" + `MAIN` + "`" + `. Default value is ` + "`" + `MAIN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_local_address",
					Description: `(Optional) Local address of IKE operation specification, valid when ike_local_identity is ` + "`" + `ADDRESS` + "`" + `, generally the value is public_ip_address of the related VPN gateway.`,
				},
				resource.Attribute{
					Name:        "ike_local_fqdn_name",
					Description: `(Optional) Local FQDN name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_local_identity",
					Description: `(Optional) Local identity way of IKE operation specification, valid values are ` + "`" + `ADDRESS` + "`" + `, ` + "`" + `FQDN` + "`" + `. Default value is ` + "`" + `ADDRESS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_proto_authen_algorithm",
					Description: `(Optional) Proto authenticate algorithm of the IKE operation specification, valid values are ` + "`" + `MD5` + "`" + `, ` + "`" + `SHA` + "`" + `. Default Value is ` + "`" + `MD5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_proto_encry_algorithm",
					Description: `(Optional) Proto encrypt algorithm of the IKE operation specification, valid values are ` + "`" + `3DES-CBC` + "`" + `, ` + "`" + `AES-CBC-128` + "`" + `, ` + "`" + `AES-CBC-128` + "`" + `, ` + "`" + `AES-CBC-256` + "`" + `, ` + "`" + `DES-CBC` + "`" + `. Default value is ` + "`" + `3DES-CBC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_remote_address",
					Description: `(Optional) Remote address of IKE operation specification, valid when ike_remote_identity is ` + "`" + `ADDRESS` + "`" + `, generally the value is public_ip_address of the related customer gateway.`,
				},
				resource.Attribute{
					Name:        "ike_remote_fqdn_name",
					Description: `(Optional) Remote FQDN name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_remote_identity",
					Description: `(Optional) Remote identity way of IKE operation specification, valid values are ` + "`" + `ADDRESS` + "`" + `, ` + "`" + `FQDN` + "`" + `. Default value is ` + "`" + `ADDRESS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_sa_lifetime_seconds",
					Description: `(Optional) SA lifetime of the IKE operation specification, unit is ` + "`" + `second` + "`" + `. The value ranges from 60 to 604800. Default value is 86400 seconds.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `(Optional) Version of the IKE operation specification. Default value is ` + "`" + `IKEV1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_encrypt_algorithm",
					Description: `(Optional) Encrypt algorithm of the IPSEC operation specification, valid values are ` + "`" + `3DES-CBC` + "`" + `, ` + "`" + `AES-CBC-128` + "`" + `, ` + "`" + `AES-CBC-128` + "`" + `, ` + "`" + `AES-CBC-256` + "`" + `, ` + "`" + `DES-CBC` + "`" + `. Default value is ` + "`" + `3DES-CBC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_integrity_algorithm",
					Description: `(Optional) Integrity algorithm of the IPSEC operation specification, valid values are ` + "`" + `SHA1` + "`" + `, ` + "`" + `MD5` + "`" + `. Default value is ` + "`" + `MD5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_pfs_dh_group",
					Description: `(Optional) PFS DH group, valid values are ` + "`" + `GROUP1` + "`" + `, ` + "`" + `GROUP2` + "`" + `, ` + "`" + `GROUP5` + "`" + `, ` + "`" + `GROUP14` + "`" + `, ` + "`" + `GROUP24` + "`" + `, ` + "`" + `NULL` + "`" + `. Default value is ` + "`" + `NULL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_sa_lifetime_seconds",
					Description: `(Optional) SA lifetime of the IPSEC operation specification, unit is ` + "`" + `second` + "`" + `. The value ranges from 180 to 604800. Default value is 3600 seconds.`,
				},
				resource.Attribute{
					Name:        "ipsec_sa_lifetime_traffic",
					Description: `(Optional) SA lifetime of the IPSEC operation specification, unit is ` + "`" + `KB` + "`" + `. The value should not be less then 2560. Default value is 1843200.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags used to associate different resources. The ` + "`" + `security_group_policy` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "local_cidr_block",
					Description: `(Required) Local cidr block.`,
				},
				resource.Attribute{
					Name:        "remote_cidr_block",
					Description: `(Required) Remote cidr block list. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "encrypt_proto",
					Description: `Encrypt proto of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "net_status",
					Description: `Net status of the VPN connection, values are ` + "`" + `AVAILABLE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `Route type of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the connection, values are ` + "`" + `PENDING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `DELETING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpn_proto",
					Description: `Vpn proto of the VPN connection. ## Import VPN connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_connection.foo vpnx-nadifg3s ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "encrypt_proto",
					Description: `Encrypt proto of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "net_status",
					Description: `Net status of the VPN connection, values are ` + "`" + `AVAILABLE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `Route type of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the connection, values are ` + "`" + `PENDING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `DELETING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpn_proto",
					Description: `Vpn proto of the VPN connection. ## Import VPN connection can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_connection.foo vpnx-nadifg3s ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_customer_gateway",
			Category:         "VPN",
			ShortDescription: `Provides a resource to create a VPN customer gateway.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"customer",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the customer gateway. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `(Required, ForceNew) Public ip of the customer gateway.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags used to associate different resources. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the customer gateway. ## Import VPN customer gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_customer_gateway.foo cgw-xfqag ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the customer gateway. ## Import VPN customer gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_customer_gateway.foo cgw-xfqag ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_gateway",
			Category:         "VPN",
			ShortDescription: `Provides a resource to create a VPN gateway.`,
			Description:      ``,
			Keywords: []string{
				"vpn",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the VPN gateway. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required, ForceNew) ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required, ForceNew) Zone of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The maximum public network output bandwidth of VPN gateway (unit: Mbps), the available values include: 5,10,20,50,100,200,500,1000. Default is 5. When charge type is ` + "`" + `PREPAID` + "`" + `, bandwidth degradation operation is unsupported.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) Charge Type of the VPN gateway, valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` and default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "prepaid_period",
					Description: `(Optional) Period of instance to be prepaid. Valid values are 1, 2, 3, 4, 6, 7, 8, 9, 12, 24, 36 and unit is month. Caution: when this para and renew_flag para are valid, the request means to renew several months more pre-paid period. This para can only be set to take effect in create operation.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `(Optional) Flag indicates whether to renew or not, valid values are ` + "`" + `NOTIFY_AND_RENEW` + "`" + `, ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, ` + "`" + `NOT_NOTIFY_AND_NOT_RENEW` + "`" + `. This para can only be set to take effect in create operation.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of gateway instance, valid values are ` + "`" + `IPSEC` + "`" + `, ` + "`" + `SSL` + "`" + ` and ` + "`" + `CCN` + "`" + `. Note: CCN type is only for whitelist customer now. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the VPN gateway when charge type is ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_address_blocked",
					Description: `Indicates whether ip address is blocked.`,
				},
				resource.Attribute{
					Name:        "new_purchase_plan",
					Description: `The plan of new purchase, valid value is ` + "`" + `PREPAID_TO_POSTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `Public ip of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "restrict_state",
					Description: `Restrict state of gateway, valid values are ` + "`" + `PRETECIVELY_ISOLATED` + "`" + `, ` + "`" + `NORMAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the VPN gateway, valid values are ` + "`" + `PENDING` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `. ## Import VPN gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_gateway.foo vpngw-8ccsnclt ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the VPN gateway when charge type is ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_address_blocked",
					Description: `Indicates whether ip address is blocked.`,
				},
				resource.Attribute{
					Name:        "new_purchase_plan",
					Description: `The plan of new purchase, valid value is ` + "`" + `PREPAID_TO_POSTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `Public ip of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "restrict_state",
					Description: `Restrict state of gateway, valid values are ` + "`" + `PRETECIVELY_ISOLATED` + "`" + `, ` + "`" + `NORMAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the VPN gateway, valid values are ` + "`" + `PENDING` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `. ## Import VPN gateway can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import tencentcloud_vpn_gateway.foo vpngw-8ccsnclt ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"tencentcloud_alb_server_attachment":          0,
		"tencentcloud_as_attachment":                  1,
		"tencentcloud_as_lifecycle_hook":              2,
		"tencentcloud_as_notification":                3,
		"tencentcloud_as_scaling_config":              4,
		"tencentcloud_as_scaling_group":               5,
		"tencentcloud_as_scaling_policy":              6,
		"tencentcloud_as_schedule":                    7,
		"tencentcloud_cam_group":                      8,
		"tencentcloud_cam_group_membership":           9,
		"tencentcloud_cam_group_policy_attachment":    10,
		"tencentcloud_cam_policy":                     11,
		"tencentcloud_cam_role":                       12,
		"tencentcloud_cam_role_policy_attachment":     13,
		"tencentcloud_cam_saml_provider":              14,
		"tencentcloud_cam_user":                       15,
		"tencentcloud_cam_user_policy_attachment":     16,
		"tencentcloud_cbs_snapshot":                   17,
		"tencentcloud_cbs_snapshot_policy":            18,
		"tencentcloud_cbs_snapshot_policy_attachment": 19,
		"tencentcloud_cbs_storage":                    20,
		"tencentcloud_cbs_storage_attachment":         21,
		"tencentcloud_ccn":                            22,
		"tencentcloud_ccn_attachment":                 23,
		"tencentcloud_ccn_bandwidth_limit":            24,
		"tencentcloud_cdn_domain":                     25,
		"tencentcloud_cfs_access_group":               26,
		"tencentcloud_cfs_access_rule":                27,
		"tencentcloud_cfs_file_system":                28,
		"tencentcloud_clb_attachment":                 29,
		"tencentcloud_clb_instance":                   30,
		"tencentcloud_clb_listener":                   31,
		"tencentcloud_clb_listener_rule":              32,
		"tencentcloud_clb_redirection":                33,
		"tencentcloud_container_cluster":              34,
		"tencentcloud_container_cluster_instance":     35,
		"tencentcloud_cos_bucket":                     36,
		"tencentcloud_cos_bucket_object":              37,
		"tencentcloud_dayu_cc_http_policy":            38,
		"tencentcloud_dayu_cc_https_policy":           39,
		"tencentcloud_dayu_ddos_policy":               40,
		"tencentcloud_dayu_ddos_policy_attachment":    41,
		"tencentcloud_dayu_ddos_policy_case":          42,
		"tencentcloud_dayu_l4_rule":                   43,
		"tencentcloud_dayu_l7_rule":                   44,
		"tencentcloud_dc_gateway":                     45,
		"tencentcloud_dc_gateway_ccn_route":           46,
		"tencentcloud_dcx":                            47,
		"tencentcloud_dnat":                           48,
		"tencentcloud_eip":                            49,
		"tencentcloud_eip_association":                50,
		"tencentcloud_elasticsearch_instance":         51,
		"tencentcloud_eni":                            52,
		"tencentcloud_eni_attachment":                 53,
		"tencentcloud_gaap_certificate":               54,
		"tencentcloud_gaap_domain_error_page":         55,
		"tencentcloud_gaap_http_domain":               56,
		"tencentcloud_gaap_http_rule":                 57,
		"tencentcloud_gaap_layer4_listener":           58,
		"tencentcloud_gaap_layer7_listener":           59,
		"tencentcloud_gaap_proxy":                     60,
		"tencentcloud_gaap_realserver":                61,
		"tencentcloud_gaap_security_policy":           62,
		"tencentcloud_gaap_security_rule":             63,
		"tencentcloud_ha_vip":                         64,
		"tencentcloud_ha_vip_eip_attachment":          65,
		"tencentcloud_instance":                       66,
		"tencentcloud_key_pair":                       67,
		"tencentcloud_kubernetes_as_scaling_group":    68,
		"tencentcloud_kubernetes_cluster":             69,
		"tencentcloud_kubernetes_cluster_attachment":  70,
		"tencentcloud_kubernetes_scale_worker":        71,
		"tencentcloud_lb":                             72,
		"tencentcloud_mongodb_instance":               73,
		"tencentcloud_mongodb_sharding_instance":      74,
		"tencentcloud_monitor_binding_object":         75,
		"tencentcloud_monitor_binding_receiver":       76,
		"tencentcloud_monitor_policy_group":           77,
		"tencentcloud_mysql_account":                  78,
		"tencentcloud_mysql_account_privilege":        79,
		"tencentcloud_mysql_backup_policy":            80,
		"tencentcloud_mysql_instance":                 81,
		"tencentcloud_mysql_privilege":                82,
		"tencentcloud_mysql_readonly_instance":        83,
		"tencentcloud_nat_gateway":                    84,
		"tencentcloud_placement_group":                85,
		"tencentcloud_postgresql_instance":            86,
		"tencentcloud_redis_backup_config":            87,
		"tencentcloud_redis_instance":                 88,
		"tencentcloud_reserved_instance":              89,
		"tencentcloud_route_entry":                    90,
		"tencentcloud_route_table":                    91,
		"tencentcloud_route_table_entry":              92,
		"tencentcloud_scf_function":                   93,
		"tencentcloud_scf_namespace":                  94,
		"tencentcloud_security_group":                 95,
		"tencentcloud_security_group_lite_rule":       96,
		"tencentcloud_security_group_rule":            97,
		"tencentcloud_ssl_certificate":                98,
		"tencentcloud_subnet":                         99,
		"tencentcloud_tcaplus_cluster":                100,
		"tencentcloud_tcaplus_idl":                    101,
		"tencentcloud_tcaplus_table":                  102,
		"tencentcloud_tcaplus_tablegroup":             103,
		"tencentcloud_vpc":                            104,
		"tencentcloud_vpn_connection":                 105,
		"tencentcloud_vpn_customer_gateway":           106,
		"tencentcloud_vpn_gateway":                    107,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
