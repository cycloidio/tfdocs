package tencentcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_configs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query scaling configuration information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "configuration_id",
					Description: `(Optional) Launch configuration ID.`,
				},
				resource.Attribute{
					Name:        "configuration_name",
					Description: `(Optional) Launch configuration name.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "configuration_list",
					Description: `A list of configuration. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "configuration_id",
					Description: `Launch configuration ID.`,
				},
				resource.Attribute{
					Name:        "configuration_name",
					Description: `Launch configuration name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the launch configuration was created.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of disk in GB. Default is 0.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Type of disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Data disk snapshot ID.`,
				},
				resource.Attribute{
					Name:        "enhanced_monitor_service",
					Description: `Whether to activate cloud monitor service.`,
				},
				resource.Attribute{
					Name:        "enhanced_security_service",
					Description: `Whether to activate cloud security service.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of available image, for example img-8toqc6s3.`,
				},
				resource.Attribute{
					Name:        "instance_tags",
					Description: `A tag list associates with an instance.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `Instance type list of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `Charge types for network traffic.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Max bandwidth of Internet access in Mbps.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `ID list of login keys.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project to which the configuration belongs. Default value is 0.`,
				},
				resource.Attribute{
					Name:        "public_ip_assigned",
					Description: `Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `Security groups to which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current statues of a launch configuration.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `System disk size of the scaling configuration in GB.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `System disk category of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `Base64-encoded User Data text.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "configuration_list",
					Description: `A list of configuration. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "configuration_id",
					Description: `Launch configuration ID.`,
				},
				resource.Attribute{
					Name:        "configuration_name",
					Description: `Launch configuration name.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the launch configuration was created.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `Configurations of data disk.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Volume of disk in GB. Default is 0.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Type of disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Data disk snapshot ID.`,
				},
				resource.Attribute{
					Name:        "enhanced_monitor_service",
					Description: `Whether to activate cloud monitor service.`,
				},
				resource.Attribute{
					Name:        "enhanced_security_service",
					Description: `Whether to activate cloud security service.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of available image, for example img-8toqc6s3.`,
				},
				resource.Attribute{
					Name:        "instance_tags",
					Description: `A tag list associates with an instance.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `Instance type list of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `Charge types for network traffic.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Max bandwidth of Internet access in Mbps.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `ID list of login keys.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project to which the configuration belongs. Default value is 0.`,
				},
				resource.Attribute{
					Name:        "public_ip_assigned",
					Description: `Specify whether to assign an Internet IP address.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `Security groups to which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current statues of a launch configuration.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `System disk size of the scaling configuration in GB.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `System disk category of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `Base64-encoded User Data text.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the detail information of an existing autoscaling group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "configuration_id",
					Description: `(Optional) Filter results by launch configuration ID.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Optional) A specified scaling group ID used to query.`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `(Optional) A scaling group name used to query.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags used to query. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "scaling_group_list",
					Description: `A list of scaling group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "configuration_id",
					Description: `Launch configuration ID.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the AS group was created.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `Default cooldown time of scaling group.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `The desired number of CVM instances.`,
				},
				resource.Attribute{
					Name:        "forward_balancer_ids",
					Description: `A list of application clb ids.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `Listener ID for application load balancers.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `ID of available load balancers.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `ID of forwarding rules.`,
				},
				resource.Attribute{
					Name:        "target_attribute",
					Description: `Attribute list of target rules.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of instance.`,
				},
				resource.Attribute{
					Name:        "load_balancer_ids",
					Description: `A list of traditional clb ids which the CVM instances attached to.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum number of CVM instances.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `The minimum number of CVM instances.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project to which the scaling group belongs. Default value is 0.`,
				},
				resource.Attribute{
					Name:        "retry_policy",
					Description: `A retry policy can be used when a creation fails.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `Auto scaling group ID.`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `Auto scaling group name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of a scaling group.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `A list of subnet IDs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the scaling group.`,
				},
				resource.Attribute{
					Name:        "termination_policies",
					Description: `A policy used to select a CVM instance to be terminated from the scaling group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc with which the instance is associated.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of available zones.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_list",
					Description: `A list of scaling group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "configuration_id",
					Description: `Launch configuration ID.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the AS group was created.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `Default cooldown time of scaling group.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `The desired number of CVM instances.`,
				},
				resource.Attribute{
					Name:        "forward_balancer_ids",
					Description: `A list of application clb ids.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `Listener ID for application load balancers.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `ID of available load balancers.`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `ID of forwarding rules.`,
				},
				resource.Attribute{
					Name:        "target_attribute",
					Description: `Attribute list of target rules.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of instance.`,
				},
				resource.Attribute{
					Name:        "load_balancer_ids",
					Description: `A list of traditional clb ids which the CVM instances attached to.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum number of CVM instances.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `The minimum number of CVM instances.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project to which the scaling group belongs. Default value is 0.`,
				},
				resource.Attribute{
					Name:        "retry_policy",
					Description: `A retry policy can be used when a creation fails.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `Auto scaling group ID.`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `Auto scaling group name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of a scaling group.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `A list of subnet IDs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the scaling group.`,
				},
				resource.Attribute{
					Name:        "termination_policies",
					Description: `A policy used to select a CVM instance to be terminated from the scaling group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc with which the instance is associated.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of available zones.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_as_scaling_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of scaling policy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Optional) Scaling policy name.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Optional) Scaling group ID.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_id",
					Description: `(Optional) Scaling policy ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "scaling_policy_list",
					Description: `A list of scaling policy. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "adjustment_type",
					Description: `Adjustment type of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "adjustment_value",
					Description: `Adjustment value of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "comparison_operator",
					Description: `Comparison operator.`,
				},
				resource.Attribute{
					Name:        "continuous_time",
					Description: `Retry times.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `Cooldown time of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `Name of an indicator.`,
				},
				resource.Attribute{
					Name:        "notification_user_group_ids",
					Description: `Users need to be notified when an alarm is triggered.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Time period in second.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Scaling policy name.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `Scaling policy ID.`,
				},
				resource.Attribute{
					Name:        "statistic",
					Description: `Statistic types.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Alarm threshold.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_policy_list",
					Description: `A list of scaling policy. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "adjustment_type",
					Description: `Adjustment type of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "adjustment_value",
					Description: `Adjustment value of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "comparison_operator",
					Description: `Comparison operator.`,
				},
				resource.Attribute{
					Name:        "continuous_time",
					Description: `Retry times.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `Cooldown time of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `Name of an indicator.`,
				},
				resource.Attribute{
					Name:        "notification_user_group_ids",
					Description: `Users need to be notified when an alarm is triggered.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Time period in second.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Scaling policy name.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `Scaling policy ID.`,
				},
				resource.Attribute{
					Name:        "statistic",
					Description: `Statistic types.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Alarm threshold.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_availability_regions",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get the available regions. By default only ` + "`" + `AVAILABLE` + "`" + ` regions will be returned, but ` + "`" + `UNAVAILABLE` + "`" + ` regions can also be fetched when ` + "`" + `include_unavailable` + "`" + ` is specified.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "include_unavailable",
					Description: `(Optional) A bool variable indicates that the query will include ` + "`" + `UNAVAILABLE` + "`" + ` regions.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) When specified, only the region with the exactly name match will be returned.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of regions will be exported and its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the region, like ` + "`" + `Guangzhou Region` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the region, like ` + "`" + `ap-guangzhou` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the region, indicate availability using ` + "`" + `AVAILABLE` + "`" + ` and ` + "`" + `UNAVAILABLE` + "`" + ` values.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `A list of regions will be exported and its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the region, like ` + "`" + `Guangzhou Region` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the region, like ` + "`" + `ap-guangzhou` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the region, indicate availability using ` + "`" + `AVAILABLE` + "`" + ` and ` + "`" + `UNAVAILABLE` + "`" + ` values.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_availability_zones",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get the available zones in current region. By default only ` + "`" + `AVAILABLE` + "`" + ` zones will be returned, but ` + "`" + `UNAVAILABLE` + "`" + ` zones can also be fetched when ` + "`" + `include_unavailable` + "`" + ` is specified.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "include_unavailable",
					Description: `(Optional) A bool variable indicates that the query will include ` + "`" + `UNAVAILABLE` + "`" + ` zones.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) When specified, only the zone with the exactly name match will be returned.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of zones will be exported and its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the zone, like ` + "`" + `Guangzhou Zone 3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An internal id for the zone, like ` + "`" + `200003` + "`" + `, usually not so useful.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the zone, like ` + "`" + `ap-guangzhou-3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the zone, indicate availability using ` + "`" + `AVAILABLE` + "`" + ` and ` + "`" + `UNAVAILABLE` + "`" + ` values.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zones",
					Description: `A list of zones will be exported and its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the zone, like ` + "`" + `Guangzhou Zone 3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An internal id for the zone, like ` + "`" + `200003` + "`" + `, usually not so useful.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the zone, like ` + "`" + `ap-guangzhou-3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the zone, indicate availability using ` + "`" + `AVAILABLE` + "`" + ` and ` + "`" + `UNAVAILABLE` + "`" + ` values.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_group_memberships",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM group memberships`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) Id of CAM group to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "membership_list",
					Description: `A list of CAM group membership. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Id of CAM group.`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `Id set of the CAM group members.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "membership_list",
					Description: `A list of CAM group membership. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Id of CAM group.`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `Id set of the CAM group members.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_group_policy_attachments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM group policy attachments`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Id of the attached CAM group to be queried.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `(Optional) Mode of Creation of the CAM user policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) Id of CAM policy to be queried.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Optional) Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "group_policy_attachment_list",
					Description: `A list of CAM group policy attachments. Each element contains the following attributes:`,
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
					Name:        "group_id",
					Description: `Id of CAM group.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Name of CAM group.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group_policy_attachment_list",
					Description: `A list of CAM group policy attachments. Each element contains the following attributes:`,
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
					Name:        "group_id",
					Description: `Id of CAM group.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Name of CAM group.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM groups`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) Id of CAM group to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CAM group to be queried.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Description of the cam group to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "group_list",
					Description: `A list of CAM groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM group.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Id of the CAM group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM group.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Description of CAM group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group_list",
					Description: `A list of CAM groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM group.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Id of the CAM group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM group.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Description of CAM group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM policies`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "create_mode",
					Description: `(Optional) Mode of creation of policy strategy. 1 means policy was created with console, and 2 means it was created by strategies.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the CAM policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CAM policy to be queried.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) Id of CAM policy to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of the policy strategy. 1 means customer strategy and 2 means preset strategy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "policy_list",
					Description: `A list of CAM policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `Number of attached users.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of creation of policy strategy. 1 means policy was created with console, and 2 means it was created by strategies.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of CAM policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the policy strategy.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `Name of attached products.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the policy strategy. 1 means customer strategy and 2 means preset strategy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_list",
					Description: `A list of CAM policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `Number of attached users.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of creation of policy strategy. 1 means policy was created with console, and 2 means it was created by strategies.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of CAM policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the policy strategy.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `Name of attached products.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the policy strategy. 1 means customer strategy and 2 means preset strategy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_role_policy_attachments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM role policy attachments`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) Id of the attached CAM role to be queried.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `(Optional) Mode of Creation of the CAM user policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) Id of CAM policy to be queried.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Optional) Type of the policy strategy. Valid values are 'User', 'QCS', '', 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "role_policy_attachment_list",
					Description: `A list of CAM role policy attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM role policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM role policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Name of CAM role.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `Id of CAM role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "role_policy_attachment_list",
					Description: `A list of CAM role policy attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM role policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM role policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Name of CAM role.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `Id of CAM role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_roles",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM roles`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the CAM role to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CAM policy to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Optional) Id of the CAM role to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "role_list",
					Description: `A list of CAM roles. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "console_login",
					Description: `Indicate whether the CAM role can be login or not.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of CAM role.`,
				},
				resource.Attribute{
					Name:        "document",
					Description: `Policy document of CAM role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM role.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `Id of CAM role.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "role_list",
					Description: `A list of CAM roles. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "console_login",
					Description: `Indicate whether the CAM role can be login or not.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of CAM role.`,
				},
				resource.Attribute{
					Name:        "document",
					Description: `Policy document of CAM role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM role.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `Id of CAM role.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The last update time of the CAM role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_saml_providers",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM SAML providers`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CAM SAML provider to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "provider_list",
					Description: `A list of CAM SAML providers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The last modify time of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM SAML provider.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "provider_list",
					Description: `A list of CAM SAML providers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The last modify time of the CAM SAML provider.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM SAML provider.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_user_policy_attachments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM user policy attachments`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) Id of the attached CAM user to be queried.`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `(Optional) Mode of Creation of the CAM user policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) Id of CAM policy to be queried.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Optional) Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "user_policy_attachment_list",
					Description: `A list of CAM user policy attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM user policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM user policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Name of CAM user.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Id of CAM user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_policy_attachment_list",
					Description: `A list of CAM user policy attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_mode",
					Description: `Mode of Creation of the CAM user policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the CAM user policy attachment.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Name of CAM user.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `The name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Id of CAM user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cam_users",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CAM users`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "console_login",
					Description: `(Optional) Indicate whether the user can login in.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(Optional) Country code of the CAM user to be queried.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Email of the CAM user to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of CAM user to be queried.`,
				},
				resource.Attribute{
					Name:        "phone_num",
					Description: `(Optional) Phone num of the CAM user to be queried.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Remark of the CAM user to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Uid of the CAM user to be queried.`,
				},
				resource.Attribute{
					Name:        "uin",
					Description: `(Optional) Uin of the CAM user to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "user_list",
					Description: `A list of CAM users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `Country code of the CAM user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email of the CAM user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM user.`,
				},
				resource.Attribute{
					Name:        "phone_num",
					Description: `Phone num of the CAM user.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Uid of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uin",
					Description: `Uin of the CAM user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Id of CAM user. Its value equals to ` + "`" + `name` + "`" + ` argument.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_list",
					Description: `A list of CAM users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `Country code of the CAM user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email of the CAM user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of CAM user.`,
				},
				resource.Attribute{
					Name:        "phone_num",
					Description: `Phone num of the CAM user.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `Uid of the CAM user.`,
				},
				resource.Attribute{
					Name:        "uin",
					Description: `Uin of the CAM user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Id of CAM user. Its value equals to ` + "`" + `name` + "`" + ` argument.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_snapshot_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CBS snapshot policies.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_id",
					Description: `(Optional) ID of the snapshot policy to be queried.`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_name",
					Description: `(Optional) Name of the snapshot policy to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_list",
					Description: `A list of snapshot policy. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "attached_storage_ids",
					Description: `Storage ids that the snapshot policy attached.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the snapshot policy.`,
				},
				resource.Attribute{
					Name:        "repeat_hours",
					Description: `Trigger hours of periodic snapshot.`,
				},
				resource.Attribute{
					Name:        "repeat_weekdays",
					Description: `Trigger days of periodic snapshot.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `Retention days of the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_id",
					Description: `ID of the snapshot policy.`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_name",
					Description: `Name of the snapshot policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the snapshot policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_policy_list",
					Description: `A list of snapshot policy. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "attached_storage_ids",
					Description: `Storage ids that the snapshot policy attached.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the snapshot policy.`,
				},
				resource.Attribute{
					Name:        "repeat_hours",
					Description: `Trigger hours of periodic snapshot.`,
				},
				resource.Attribute{
					Name:        "repeat_weekdays",
					Description: `Trigger days of periodic snapshot.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `Retention days of the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_id",
					Description: `ID of the snapshot policy.`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_name",
					Description: `Name of the snapshot policy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the snapshot policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_snapshots",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CBS snapshots.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the CBS instance locates at.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project within the snapshot.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) ID of the snapshot to be queried.`,
				},
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `(Optional) Name of the snapshot to be queried.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `(Optional) ID of the the CBS which this snapshot created from.`,
				},
				resource.Attribute{
					Name:        "storage_usage",
					Description: `(Optional) Types of CBS which this snapshot created from, and available values include SYSTEM_DISK and DATA_DISK. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "snapshot_list",
					Description: `A list of snapshot. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the CBS instance locates at.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of snapshot.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `Indicates whether the snapshot is encrypted.`,
				},
				resource.Attribute{
					Name:        "percent",
					Description: `Snapshot creation progress percentage.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project within the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `Name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `ID of the the CBS which this snapshot created from.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Volume of storage which this snapshot created from.`,
				},
				resource.Attribute{
					Name:        "storage_usage",
					Description: `Types of CBS which this snapshot created from.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_list",
					Description: `A list of snapshot. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the CBS instance locates at.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of snapshot.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `Indicates whether the snapshot is encrypted.`,
				},
				resource.Attribute{
					Name:        "percent",
					Description: `Snapshot creation progress percentage.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project within the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `Name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `ID of the the CBS which this snapshot created from.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Volume of storage which this snapshot created from.`,
				},
				resource.Attribute{
					Name:        "storage_usage",
					Description: `Types of CBS which this snapshot created from.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cbs_storages",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CBS storages.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the CBS instance locates at.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project with which the CBS is associated.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `(Optional) ID of the CBS to be queried.`,
				},
				resource.Attribute{
					Name:        "storage_name",
					Description: `(Optional) Name of the CBS to be queried.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) Types of storage medium, and available values include CLOUD_BASIC, CLOUD_PREMIUM and CLOUD_SSD.`,
				},
				resource.Attribute{
					Name:        "storage_usage",
					Description: `(Optional) Types of CBS, and available values include SYSTEM_DISK and DATA_DISK. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "storage_list",
					Description: `A list of storage. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "attached",
					Description: `Indicates whether the CBS is mounted the CVM.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The zone of CBS.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Pay type of the CBS instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of CBS.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `Indicates whether CBS is encrypted.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the CVM instance that be mounted by this CBS.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `ID of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_name",
					Description: `Name of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Volume of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Types of storage medium.`,
				},
				resource.Attribute{
					Name:        "storage_usage",
					Description: `Types of CBS.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The available tags within this CBS.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "storage_list",
					Description: `A list of storage. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "attached",
					Description: `Indicates whether the CBS is mounted the CVM.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The zone of CBS.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Pay type of the CBS instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of CBS.`,
				},
				resource.Attribute{
					Name:        "encrypt",
					Description: `Indicates whether CBS is encrypted.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the CVM instance that be mounted by this CBS.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_id",
					Description: `ID of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_name",
					Description: `Name of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Volume of CBS.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Types of storage medium.`,
				},
				resource.Attribute{
					Name:        "storage_usage",
					Description: `Types of CBS.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The available tags within this CBS.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ccn_bandwidth_limits",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CCN bandwidth limits.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ccn_id",
					Description: `(Required) ID of the CCN to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "limits",
					Description: `The bandwidth limits of regions:`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit",
					Description: `Limitation of bandwidth.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Limitation of region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "limits",
					Description: `The bandwidth limits of regions:`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit",
					Description: `Limitation of bandwidth.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Limitation of region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ccn_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CCN instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ccn_id",
					Description: `(Optional) ID of the CCN to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CCN to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of CCN.`,
				},
				resource.Attribute{
					Name:        "attachment_list",
					Description: `Information list of instance is attached.`,
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
					Name:        "instance_id",
					Description: `ID of instance is attached.`,
				},
				resource.Attribute{
					Name:        "instance_region",
					Description: `The region that the instance locates at.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Type of attached instance network, and available values include VPC, DIRECTCONNECT, BMVPC and VPNGW.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance is attached, and available values include PENDING, ACTIVE, EXPIRED, REJECTED, DELETED, FAILED(asynchronous forced disassociation after 2 hours), ATTACHING, DETACHING and DETACHFAILED(asynchronous forced disassociation after 2 hours).`,
				},
				resource.Attribute{
					Name:        "ccn_id",
					Description: `ID of the CCN.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the CCN.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CCN.`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `Service quality of CCN, and the available value include 'PT', 'AU', 'AG'. The default is 'AU'.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance. The available value include 'ISOLATED'(arrears) and 'AVAILABLE'.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of CCN.`,
				},
				resource.Attribute{
					Name:        "attachment_list",
					Description: `Information list of instance is attached.`,
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
					Name:        "instance_id",
					Description: `ID of instance is attached.`,
				},
				resource.Attribute{
					Name:        "instance_region",
					Description: `The region that the instance locates at.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Type of attached instance network, and available values include VPC, DIRECTCONNECT, BMVPC and VPNGW.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance is attached, and available values include PENDING, ACTIVE, EXPIRED, REJECTED, DELETED, FAILED(asynchronous forced disassociation after 2 hours), ATTACHING, DETACHING and DETACHFAILED(asynchronous forced disassociation after 2 hours).`,
				},
				resource.Attribute{
					Name:        "ccn_id",
					Description: `ID of the CCN.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the CCN.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CCN.`,
				},
				resource.Attribute{
					Name:        "qos",
					Description: `Service quality of CCN, and the available value include 'PT', 'AU', 'AG'. The default is 'AU'.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of instance. The available value include 'ISOLATED'(arrears) and 'AVAILABLE'.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cfs_access_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the detail information of CFS access group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_group_id",
					Description: `(Optional) A specified access group ID used to query.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A access group Name used to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_group_list",
					Description: `An information list of CFS access group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_group_id",
					Description: `ID of the access group.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the access group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the access group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the access group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_group_list",
					Description: `An information list of CFS access group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_group_id",
					Description: `ID of the access group.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the access group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the access group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the access group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cfs_access_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the detail information of CFS access rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_group_id",
					Description: `(Required) A specified access group ID used to query.`,
				},
				resource.Attribute{
					Name:        "access_rule_id",
					Description: `(Optional) A specified access rule ID used to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_rule_list",
					Description: `An information list of CFS access rule. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_rule_id",
					Description: `ID of the access rule.`,
				},
				resource.Attribute{
					Name:        "auth_client_ip",
					Description: `Allowed IP of the access rule.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of access rule.`,
				},
				resource.Attribute{
					Name:        "rw_permission",
					Description: `Read and write permissions.`,
				},
				resource.Attribute{
					Name:        "user_permission",
					Description: `The permissions of accessing users.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_rule_list",
					Description: `An information list of CFS access rule. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_rule_id",
					Description: `ID of the access rule.`,
				},
				resource.Attribute{
					Name:        "auth_client_ip",
					Description: `Allowed IP of the access rule.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority level of access rule.`,
				},
				resource.Attribute{
					Name:        "rw_permission",
					Description: `Read and write permissions.`,
				},
				resource.Attribute{
					Name:        "user_permission",
					Description: `The permissions of accessing users.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cfs_file_systems",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the detail information of cloud file systems(CFS).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the file system locates at.`,
				},
				resource.Attribute{
					Name:        "file_system_id",
					Description: `(Optional) A specified file system ID used to query.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A file system name used to query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of a vpc subnetwork.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the vpc to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "file_system_list",
					Description: `An information list of cloud file system. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_group_id",
					Description: `ID of the access group.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the file system locates at.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the file system.`,
				},
				resource.Attribute{
					Name:        "file_system_id",
					Description: `ID of the file system.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the file system.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the file system.`,
				},
				resource.Attribute{
					Name:        "size_limit",
					Description: `Size limit of the file system.`,
				},
				resource.Attribute{
					Name:        "size_used",
					Description: `Size used of the file system.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the file system.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Storage type of the file system.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "file_system_list",
					Description: `An information list of cloud file system. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_group_id",
					Description: `ID of the access group.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the file system locates at.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the file system.`,
				},
				resource.Attribute{
					Name:        "file_system_id",
					Description: `ID of the file system.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the file system.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the file system.`,
				},
				resource.Attribute{
					Name:        "size_limit",
					Description: `Size limit of the file system.`,
				},
				resource.Attribute{
					Name:        "size_used",
					Description: `Size used of the file system.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the file system.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `Storage type of the file system.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_attachments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CLB attachments`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required) Id of the CLB to be queried.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) Id of the CLB listener to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional) Id of the CLB listener rule. If the protocol of listener is HTTP/HTTPS, this para is required. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "attachment_list",
					Description: `A list of cloud load balancer attachment configurations. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `Id of the CLB.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `Id of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `Type of protocol within the listener, and available values include 'TCP', 'UDP', 'HTTP', 'HTTPS' and 'TCP_SSL'.NOTES: TCP_SSL is testing internally, please apply if you need to use.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of the CLB listener rule.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `Information of the backends to be attached.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Id of the backend server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the backend server.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Forwarding weight of the backend service, the range of [0, 100], defaults to 10.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "attachment_list",
					Description: `A list of cloud load balancer attachment configurations. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `Id of the CLB.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `Id of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `Type of protocol within the listener, and available values include 'TCP', 'UDP', 'HTTP', 'HTTPS' and 'TCP_SSL'.NOTES: TCP_SSL is testing internally, please apply if you need to use.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of the CLB listener rule.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `Information of the backends to be attached.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Id of the backend server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the backend server.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Forwarding weight of the backend service, the range of [0, 100], defaults to 10.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CLB`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Optional) Id of the CLB to be queried.`,
				},
				resource.Attribute{
					Name:        "clb_name",
					Description: `(Optional) Name of the CLB to be queried.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Optional) Type of CLB instance, and available values include 'OPEN' and 'INTERNAL'.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project id of the CLB.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "clb_list",
					Description: `A list of cloud load balancers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `Id of CLB.`,
				},
				resource.Attribute{
					Name:        "clb_name",
					Description: `Name of CLB.`,
				},
				resource.Attribute{
					Name:        "clb_vips",
					Description: `The virtual service address table of the CLB.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the CLB.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Types of CLB.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Id of the project.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `Id set of the security groups.`,
				},
				resource.Attribute{
					Name:        "status_time",
					Description: `Latest state transition time of CLB.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of CLB.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Id of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The available tags within this CLB.`,
				},
				resource.Attribute{
					Name:        "target_region_info_region",
					Description: `Region information of backend service are attached the CLB.`,
				},
				resource.Attribute{
					Name:        "target_region_info_vpc_id",
					Description: `VpcId information of backend service are attached the CLB.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Id of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_list",
					Description: `A list of cloud load balancers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `Id of CLB.`,
				},
				resource.Attribute{
					Name:        "clb_name",
					Description: `Name of CLB.`,
				},
				resource.Attribute{
					Name:        "clb_vips",
					Description: `The virtual service address table of the CLB.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the CLB.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Types of CLB.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Id of the project.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `Id set of the security groups.`,
				},
				resource.Attribute{
					Name:        "status_time",
					Description: `Latest state transition time of CLB.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of CLB.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Id of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The available tags within this CLB.`,
				},
				resource.Attribute{
					Name:        "target_region_info_region",
					Description: `Region information of backend service are attached the CLB.`,
				},
				resource.Attribute{
					Name:        "target_region_info_vpc_id",
					Description: `VpcId information of backend service are attached the CLB.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Id of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_listener_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CLB listener rule`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required) Id of the CLB to be queried.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) Id of the CLB listener to be queried.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Domain name of the forwarding rule to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional) Id of the forwarding rule to be queried.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) Scheduling method of the forwarding rule of thr CLB listener, and available values include 'WRR', 'IP HASH' and 'LEAST_CONN'. The default is 'WRR'.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) Url of the forwarding rule to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rule_list",
					Description: `A list of forward rules of listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "certificate_ca_id",
					Description: `Id of the client certificate. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `Id of the server certificate. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_ssl_mode",
					Description: `Type of SSL Mode, and available values inclue 'UNIDIRECTIONAL', 'MUTUAL'.NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `Id of the CLB.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check, and the default is 3. If a success result is returned for the health check three consecutive times, the CVM is identified as healthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `HTTP Status Code. The default is 31 and value range is 1-31. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value 4xx is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_http_domain",
					Description: `Domain name of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_http_method",
					Description: `Methods of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol. The default is 'HEAD', the available value include 'HEAD' and 'GET'.`,
				},
				resource.Attribute{
					Name:        "health_check_http_path",
					Description: `Path of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_interval_time",
					Description: `Interval time of health check. The value range is 5-300 sec, and the default is 5 sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealth threshold of health check, and the default is 3. If a success result is returned for the health check three consecutive times, the CVM is identified as unhealthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `Id of the listener.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of the rule.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling method of the CLB listener, and available values include 'WRR', 'IP_HASH' and 'LEAST_CONN'. The default is 'WRR'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "session_expire_time",
					Description: `Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as 'WRR'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rule_list",
					Description: `A list of forward rules of listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "certificate_ca_id",
					Description: `Id of the client certificate. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `Id of the server certificate. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_ssl_mode",
					Description: `Type of SSL Mode, and available values inclue 'UNIDIRECTIONAL', 'MUTUAL'.NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `Id of the CLB.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check, and the default is 3. If a success result is returned for the health check three consecutive times, the CVM is identified as healthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `HTTP Status Code. The default is 31 and value range is 1-31. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value 4xx is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values. NOTES: The 'HTTP' health check of the 'TCP' listener only supports specifying one health check status code. NOTES: Only supports listeners of 'HTTP' and 'HTTPS' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_http_domain",
					Description: `Domain name of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_http_method",
					Description: `Methods of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol. The default is 'HEAD', the available value include 'HEAD' and 'GET'.`,
				},
				resource.Attribute{
					Name:        "health_check_http_path",
					Description: `Path of health check. NOTES: Only supports listeners of 'HTTPS' and 'HTTP' protocol.`,
				},
				resource.Attribute{
					Name:        "health_check_interval_time",
					Description: `Interval time of health check. The value range is 5-300 sec, and the default is 5 sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealth threshold of health check, and the default is 3. If a success result is returned for the health check three consecutive times, the CVM is identified as unhealthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `Id of the listener.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of the rule.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling method of the CLB listener, and available values include 'WRR', 'IP_HASH' and 'LEAST_CONN'. The default is 'WRR'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "session_expire_time",
					Description: `Time of session persistence within the CLB listener. NOTES: Available when scheduler is specified as 'WRR'. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_listeners",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CLB listener`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required) Id of the CLB to be queried.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Optional) Id of the listener to be queried.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Type of protocol within the listener, and available values are 'TCP', 'UDP', 'HTTP', 'HTTPS' and 'TCP_SSL'.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listener_list",
					Description: `A list of listeners of cloud load balancers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "certificate_ca_id",
					Description: `Id of the client certificate. It must be set when SSLMode is 'mutual'. NOTES: only supported by listeners of 'HTTPS' and 'TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `Id of the server certificate. It must be set when protocol is 'HTTPS' or 'TCP_SSL'. NOTES: only supported by listeners of 'HTTPS' and 'TCP_SSL' protocol and must be set when it is available.`,
				},
				resource.Attribute{
					Name:        "certificate_ssl_mode",
					Description: `Type of certificate, and available values inclue 'UNIDIRECTIONAL', 'MUTUAL'. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol and must be set when it is available.`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `Id of the CLB.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check, and the default is 3. If a success result is returned for the health check three consecutive times, the CVM is identified as healthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_interval_time",
					Description: `Interval time of health check. The value range is 5-300 sec, and the default is 5 sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_time_out",
					Description: `Response timeout of health check. The value range is 2-60 sec, and the default is 2 sec. Response timeout needs to be less than check interval. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealthy threshold of health check, and the default is 3. If a success result is returned for the health check three consecutive times, the CVM is identified as unhealthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `Id of the listener.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `Name of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the listener. Available values are 'HTTP', 'HTTPS', 'TCP', 'UDP', 'TCP_SSL'.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling method of the CLB listener, and available values are 'WRR' and 'LEAST_CONN'. The default is 'WRR'. NOTES: The listener of 'HTTP' and 'HTTPS' protocol additionally supports the 'IP HASH' method. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "session_expire_time",
					Description: `Time of session persistence within the CLB listener. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "sni_switch",
					Description: `Indicates whether SNI is enabled. NOTES: Only supported by 'HTTPS' protocol.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_list",
					Description: `A list of listeners of cloud load balancers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "certificate_ca_id",
					Description: `Id of the client certificate. It must be set when SSLMode is 'mutual'. NOTES: only supported by listeners of 'HTTPS' and 'TCP_SSL' protocol.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `Id of the server certificate. It must be set when protocol is 'HTTPS' or 'TCP_SSL'. NOTES: only supported by listeners of 'HTTPS' and 'TCP_SSL' protocol and must be set when it is available.`,
				},
				resource.Attribute{
					Name:        "certificate_ssl_mode",
					Description: `Type of certificate, and available values inclue 'UNIDIRECTIONAL', 'MUTUAL'. NOTES: Only supports listeners of 'HTTPS' and 'TCP_SSL' protocol and must be set when it is available.`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `Id of the CLB.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check, and the default is 3. If a success result is returned for the health check three consecutive times, the CVM is identified as healthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_interval_time",
					Description: `Interval time of health check. The value range is 5-300 sec, and the default is 5 sec. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_time_out",
					Description: `Response timeout of health check. The value range is 2-60 sec, and the default is 2 sec. Response timeout needs to be less than check interval. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealthy threshold of health check, and the default is 3. If a success result is returned for the health check three consecutive times, the CVM is identified as unhealthy. The value range is 2-10. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `Id of the listener.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `Name of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the CLB listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the listener. Available values are 'HTTP', 'HTTPS', 'TCP', 'UDP', 'TCP_SSL'.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling method of the CLB listener, and available values are 'WRR' and 'LEAST_CONN'. The default is 'WRR'. NOTES: The listener of 'HTTP' and 'HTTPS' protocol additionally supports the 'IP HASH' method. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "session_expire_time",
					Description: `Time of session persistence within the CLB listener. NOTES: TCP/UDP/TCP_SSL listener allows direct configuration, HTTP/HTTPS listener needs to be configured in tencentcloud_clb_listener_rule.`,
				},
				resource.Attribute{
					Name:        "sni_switch",
					Description: `Indicates whether SNI is enabled. NOTES: Only supported by 'HTTPS' protocol.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_clb_redirections",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of CLB redirections`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clb_id",
					Description: `(Required) Id of the CLB to be queried.`,
				},
				resource.Attribute{
					Name:        "source_listener_id",
					Description: `(Required) Id of source listener to be queried.`,
				},
				resource.Attribute{
					Name:        "source_rule_id",
					Description: `(Required) Rule id of source listener to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "target_listener_id",
					Description: `(Optional) Id of target listener to be queried.`,
				},
				resource.Attribute{
					Name:        "target_rule_id",
					Description: `(Optional) Rule id of target listener to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "redirection_list",
					Description: `A list of cloud load balancer redirection configurations. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `Id of the CLB.`,
				},
				resource.Attribute{
					Name:        "source_listener_id",
					Description: `Id of source listener.`,
				},
				resource.Attribute{
					Name:        "source_rule_id",
					Description: `Rule id of source listener.`,
				},
				resource.Attribute{
					Name:        "target_listener_id",
					Description: `Id of target listener.`,
				},
				resource.Attribute{
					Name:        "target_rule_id",
					Description: `Rule id of target listener.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "redirection_list",
					Description: `A list of cloud load balancer redirection configurations. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "clb_id",
					Description: `Id of the CLB.`,
				},
				resource.Attribute{
					Name:        "source_listener_id",
					Description: `Id of source listener.`,
				},
				resource.Attribute{
					Name:        "source_rule_id",
					Description: `Rule id of source listener.`,
				},
				resource.Attribute{
					Name:        "target_listener_id",
					Description: `Id of target listener.`,
				},
				resource.Attribute{
					Name:        "target_rule_id",
					Description: `Rule id of target listener.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_container_cluster_instances",
			Category:         "Data Sources",
			ShortDescription: `Get all instances of the specific cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) An id identify the cluster, like cls-xxxxxx.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) An int variable describe how many instances in return at most. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `An information list of kubernetes instances.`,
				},
				resource.Attribute{
					Name:        "abnormal_reason",
					Description: `Describe the reason when node is in abnormal state(if it was).`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Describe the cpu of the node.`,
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
					Name:        "mem",
					Description: `Describe the memory of the node.`,
				},
				resource.Attribute{
					Name:        "wan_ip",
					Description: `Describe the wan ip of the node.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Number of instances.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nodes",
					Description: `An information list of kubernetes instances.`,
				},
				resource.Attribute{
					Name:        "abnormal_reason",
					Description: `Describe the reason when node is in abnormal state(if it was).`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Describe the cpu of the node.`,
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
					Name:        "mem",
					Description: `Describe the memory of the node.`,
				},
				resource.Attribute{
					Name:        "wan_ip",
					Description: `Describe the wan ip of the node.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Number of instances.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_container_clusters",
			Category:         "Data Sources",
			ShortDescription: `Get container clusters in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) An id identify the cluster, like ` + "`" + `cls-xxxxxx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) An int variable describe how many cluster in return at most. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `An information list of kubernetes clusters.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `An id identify the cluster, like ` + "`" + `cls-xxxxxx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name the cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the cluster.`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `Describe the running kubernetes version on the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_num",
					Description: `Describe how many cluster instances in the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_status",
					Description: `Describe the current status of the instances in the cluster.`,
				},
				resource.Attribute{
					Name:        "security_certification_authority",
					Description: `Describe the certificate string needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "security_cluster_external_endpoint",
					Description: `Describe the address needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "security_password",
					Description: `Describe the password needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "security_username",
					Description: `Describe the username needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "total_cpu",
					Description: `Describe the total cpu of each instance in the cluster.`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `Describe the total memory of each instance in the cluster.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Number of clusters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "clusters",
					Description: `An information list of kubernetes clusters.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `An id identify the cluster, like ` + "`" + `cls-xxxxxx` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name the cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the cluster.`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `Describe the running kubernetes version on the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_num",
					Description: `Describe how many cluster instances in the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes_status",
					Description: `Describe the current status of the instances in the cluster.`,
				},
				resource.Attribute{
					Name:        "security_certification_authority",
					Description: `Describe the certificate string needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "security_cluster_external_endpoint",
					Description: `Describe the address needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "security_password",
					Description: `Describe the password needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "security_username",
					Description: `Describe the username needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "total_cpu",
					Description: `Describe the total cpu of each instance in the cluster.`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `Describe the total memory of each instance in the cluster.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Number of clusters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cos_bucket_object",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the metadata of an object stored inside a bucket.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Name of the bucket that contains the objects to query.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The full path to the object inside the bucket.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `ETag generated for the object, which is may not equal to MD5 value.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Object storage type such as STANDARD.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `ETag generated for the object, which is may not equal to MD5 value.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Object storage type such as STANDARD.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_cos_buckets",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the COS buckets of the current Tencent Cloud user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket_prefix",
					Description: `(Optional) A prefix string to filter results by bucket name.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags to filter bucket. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "bucket_list",
					Description: `A list of bucket. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Bucket name, the format likes ` + "`" + `<bucket>-<appid>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cors_rules",
					Description: `A list of CORS rule configurations.`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `Specifies which headers are allowed.`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `Specifies which origins are allowed.`,
				},
				resource.Attribute{
					Name:        "expose_headers",
					Description: `Specifies expose header in the response.`,
				},
				resource.Attribute{
					Name:        "max_age_seconds",
					Description: `Specifies time in seconds that browser can cache the response for a preflight request.`,
				},
				resource.Attribute{
					Name:        "lifecycle_rules",
					Description: `The lifecycle configuration of a bucket.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `Specifies a period in the object's expire.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Specifies the date after which you want the corresponding action to take effect.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `Specifies the number of days after object creation when the specific rule action takes effect.`,
				},
				resource.Attribute{
					Name:        "filter_prefix",
					Description: `Object key prefix identifying one or more objects to which the rule applies.`,
				},
				resource.Attribute{
					Name:        "transition",
					Description: `Specifies a period in the object's transitions.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Specifies the date after which you want the corresponding action to take effect.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `Specifies the number of days after object creation when the specific rule action takes effect.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Specifies the storage class to which you want the object to transition. Available values include STANDARD, STANDARD_IA and ARCHIVE.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of a bucket.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `A list of one element containing configuration parameters used when the bucket is used as a website.`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `An absolute path to the document to return in case of a 4XX error.`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `COS returns this index document when requests are made to the root domain or any of the subfolders.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket_list",
					Description: `A list of bucket. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Bucket name, the format likes ` + "`" + `<bucket>-<appid>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cors_rules",
					Description: `A list of CORS rule configurations.`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `Specifies which headers are allowed.`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `Specifies which origins are allowed.`,
				},
				resource.Attribute{
					Name:        "expose_headers",
					Description: `Specifies expose header in the response.`,
				},
				resource.Attribute{
					Name:        "max_age_seconds",
					Description: `Specifies time in seconds that browser can cache the response for a preflight request.`,
				},
				resource.Attribute{
					Name:        "lifecycle_rules",
					Description: `The lifecycle configuration of a bucket.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `Specifies a period in the object's expire.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Specifies the date after which you want the corresponding action to take effect.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `Specifies the number of days after object creation when the specific rule action takes effect.`,
				},
				resource.Attribute{
					Name:        "filter_prefix",
					Description: `Object key prefix identifying one or more objects to which the rule applies.`,
				},
				resource.Attribute{
					Name:        "transition",
					Description: `Specifies a period in the object's transitions.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Specifies the date after which you want the corresponding action to take effect.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `Specifies the number of days after object creation when the specific rule action takes effect.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Specifies the storage class to which you want the object to transition. Available values include STANDARD, STANDARD_IA and ARCHIVE.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of a bucket.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `A list of one element containing configuration parameters used when the bucket is used as a website.`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `An absolute path to the document to return in case of a 4XX error.`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `COS returns this index document when requests are made to the root domain or any of the subfolders.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_cc_http_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query dayu CC http policies`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) Id of the resource that the CC http policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Type of the resource that the CC http policy works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CC http policy to be queried.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) Id of the CC http policy to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of CC http policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action mode.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Max frequency per minute.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `Ip of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `ID of the resource that the CC self-define http policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the CC self-define http policy works for.`,
				},
				resource.Attribute{
					Name:        "smode",
					Description: `Match mode.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Indicate the CC self-define http policy takes effect or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of CC http policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action mode.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Max frequency per minute.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `Ip of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the CC self-define http policy.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `ID of the resource that the CC self-define http policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the CC self-define http policy works for.`,
				},
				resource.Attribute{
					Name:        "smode",
					Description: `Match mode.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Indicate the CC self-define http policy takes effect or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_cc_https_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query dayu CC https policies`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) Id of the resource that the CC https policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Type of the resource that the CC https policy works for, valid value is ` + "`" + `bgpip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the CC https policy to be queried.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) Id of the CC https policy to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of CC https policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action mode.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `Ip of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `ID of the resource that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Rule id of the domain that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Indicate the CC self-define https policy takes effect or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of CC https policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action mode.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `Ip of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the CC self-define https policy.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `ID of the resource that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Rule id of the domain that the CC self-define https policy works for.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Indicate the CC self-define https policy takes effect or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_ddos_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query dayu DDoS policies`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Type of the resource that the DDoS policy works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) Id of the DDoS policy to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of DDoS policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "drop_options",
					Description: `Option list of abnormal check of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "bad_conn_threshold",
					Description: `The number of new connections based on destination IP that trigger suppression of connections.`,
				},
				resource.Attribute{
					Name:        "check_sync_conn",
					Description: `Indicate whether to check null connection or not.`,
				},
				resource.Attribute{
					Name:        "conn_timeout",
					Description: `Connection timeout of abnormal connection check.`,
				},
				resource.Attribute{
					Name:        "d_conn_limit",
					Description: `The limit of concurrent connections based on destination IP.`,
				},
				resource.Attribute{
					Name:        "d_new_limit",
					Description: `The limit of new connections based on destination IP.`,
				},
				resource.Attribute{
					Name:        "drop_icmp",
					Description: `Indicate whether to drop ICMP protocol or not.`,
				},
				resource.Attribute{
					Name:        "drop_other",
					Description: `Indicate whether to drop other protocols(exclude TCP/UDP/ICMP) or not.`,
				},
				resource.Attribute{
					Name:        "drop_tcp",
					Description: `Indicate whether to drop TCP protocol or not.`,
				},
				resource.Attribute{
					Name:        "drop_udp",
					Description: `Indicate to drop UDP protocol or not.`,
				},
				resource.Attribute{
					Name:        "icmp_mbps_limit",
					Description: `The limit of ICMP traffic rate.`,
				},
				resource.Attribute{
					Name:        "null_conn_enable",
					Description: `Indicate to enable null connection or not.`,
				},
				resource.Attribute{
					Name:        "other_mbps_limit",
					Description: `The limit of other protocols(exclude TCP/UDP/ICMP) traffic rate.`,
				},
				resource.Attribute{
					Name:        "s_conn_limit",
					Description: `The limit of concurrent connections based on source IP.`,
				},
				resource.Attribute{
					Name:        "s_new_limit",
					Description: `The limit of new connections based on source IP.`,
				},
				resource.Attribute{
					Name:        "syn_limit",
					Description: `The limit of syn of abnormal connection check.`,
				},
				resource.Attribute{
					Name:        "syn_rate",
					Description: `The percentage of syn in ack of abnormal connection check.`,
				},
				resource.Attribute{
					Name:        "tcp_mbps_limit",
					Description: `The limit of TCP traffic.`,
				},
				resource.Attribute{
					Name:        "udp_mbps_limit",
					Description: `The limit of UDP traffic rate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "packet_filters",
					Description: `Message filter options list.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action of port to take.`,
				},
				resource.Attribute{
					Name:        "d_end_port",
					Description: `End port of the destination.`,
				},
				resource.Attribute{
					Name:        "d_start_port",
					Description: `Start port of the destination.`,
				},
				resource.Attribute{
					Name:        "depth",
					Description: `The depth of match.`,
				},
				resource.Attribute{
					Name:        "is_include",
					Description: `Indicate whether to include the key word/regular expression or not.`,
				},
				resource.Attribute{
					Name:        "match_begin",
					Description: `Indicate whether to check load or not.`,
				},
				resource.Attribute{
					Name:        "match_str",
					Description: `The key word or regular expression.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `Match type.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `The offset of match.`,
				},
				resource.Attribute{
					Name:        "pkt_length_max",
					Description: `The max length of the packet.`,
				},
				resource.Attribute{
					Name:        "pkt_length_min",
					Description: `The minimum length of the packet.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol.`,
				},
				resource.Attribute{
					Name:        "s_end_port",
					Description: `End port of the source.`,
				},
				resource.Attribute{
					Name:        "s_start_port",
					Description: `Start port of the source.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of policy.`,
				},
				resource.Attribute{
					Name:        "port_filters",
					Description: `Port limits of abnormal check of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action of port to take.`,
				},
				resource.Attribute{
					Name:        "end_port",
					Description: `End port.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The type of forbidden port, and valid values are 0, 1, 2. 0 for destination port, 1 for source port and 2 for both destination and source posts.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol.`,
				},
				resource.Attribute{
					Name:        "start_port",
					Description: `Start port.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `Id of policy case that the DDoS policy works for.`,
				},
				resource.Attribute{
					Name:        "watermark_filters",
					Description: `Watermark policy options, and only support one watermark policy at most.`,
				},
				resource.Attribute{
					Name:        "auto_remove",
					Description: `Indicate whether to auto-remove the watermark or not.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `The offset of watermark.`,
				},
				resource.Attribute{
					Name:        "open_switch",
					Description: `Indicate whether to open watermark or not.`,
				},
				resource.Attribute{
					Name:        "tcp_port_list",
					Description: `Port range of TCP.`,
				},
				resource.Attribute{
					Name:        "udp_port_list",
					Description: `Port range of TCP.`,
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
					Name:        "list",
					Description: `A list of DDoS policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "drop_options",
					Description: `Option list of abnormal check of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "bad_conn_threshold",
					Description: `The number of new connections based on destination IP that trigger suppression of connections.`,
				},
				resource.Attribute{
					Name:        "check_sync_conn",
					Description: `Indicate whether to check null connection or not.`,
				},
				resource.Attribute{
					Name:        "conn_timeout",
					Description: `Connection timeout of abnormal connection check.`,
				},
				resource.Attribute{
					Name:        "d_conn_limit",
					Description: `The limit of concurrent connections based on destination IP.`,
				},
				resource.Attribute{
					Name:        "d_new_limit",
					Description: `The limit of new connections based on destination IP.`,
				},
				resource.Attribute{
					Name:        "drop_icmp",
					Description: `Indicate whether to drop ICMP protocol or not.`,
				},
				resource.Attribute{
					Name:        "drop_other",
					Description: `Indicate whether to drop other protocols(exclude TCP/UDP/ICMP) or not.`,
				},
				resource.Attribute{
					Name:        "drop_tcp",
					Description: `Indicate whether to drop TCP protocol or not.`,
				},
				resource.Attribute{
					Name:        "drop_udp",
					Description: `Indicate to drop UDP protocol or not.`,
				},
				resource.Attribute{
					Name:        "icmp_mbps_limit",
					Description: `The limit of ICMP traffic rate.`,
				},
				resource.Attribute{
					Name:        "null_conn_enable",
					Description: `Indicate to enable null connection or not.`,
				},
				resource.Attribute{
					Name:        "other_mbps_limit",
					Description: `The limit of other protocols(exclude TCP/UDP/ICMP) traffic rate.`,
				},
				resource.Attribute{
					Name:        "s_conn_limit",
					Description: `The limit of concurrent connections based on source IP.`,
				},
				resource.Attribute{
					Name:        "s_new_limit",
					Description: `The limit of new connections based on source IP.`,
				},
				resource.Attribute{
					Name:        "syn_limit",
					Description: `The limit of syn of abnormal connection check.`,
				},
				resource.Attribute{
					Name:        "syn_rate",
					Description: `The percentage of syn in ack of abnormal connection check.`,
				},
				resource.Attribute{
					Name:        "tcp_mbps_limit",
					Description: `The limit of TCP traffic.`,
				},
				resource.Attribute{
					Name:        "udp_mbps_limit",
					Description: `The limit of UDP traffic rate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "packet_filters",
					Description: `Message filter options list.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action of port to take.`,
				},
				resource.Attribute{
					Name:        "d_end_port",
					Description: `End port of the destination.`,
				},
				resource.Attribute{
					Name:        "d_start_port",
					Description: `Start port of the destination.`,
				},
				resource.Attribute{
					Name:        "depth",
					Description: `The depth of match.`,
				},
				resource.Attribute{
					Name:        "is_include",
					Description: `Indicate whether to include the key word/regular expression or not.`,
				},
				resource.Attribute{
					Name:        "match_begin",
					Description: `Indicate whether to check load or not.`,
				},
				resource.Attribute{
					Name:        "match_str",
					Description: `The key word or regular expression.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `Match type.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `The offset of match.`,
				},
				resource.Attribute{
					Name:        "pkt_length_max",
					Description: `The max length of the packet.`,
				},
				resource.Attribute{
					Name:        "pkt_length_min",
					Description: `The minimum length of the packet.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol.`,
				},
				resource.Attribute{
					Name:        "s_end_port",
					Description: `End port of the source.`,
				},
				resource.Attribute{
					Name:        "s_start_port",
					Description: `Start port of the source.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of policy.`,
				},
				resource.Attribute{
					Name:        "port_filters",
					Description: `Port limits of abnormal check of the DDoS policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action of port to take.`,
				},
				resource.Attribute{
					Name:        "end_port",
					Description: `End port.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The type of forbidden port, and valid values are 0, 1, 2. 0 for destination port, 1 for source port and 2 for both destination and source posts.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol.`,
				},
				resource.Attribute{
					Name:        "start_port",
					Description: `Start port.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `Id of policy case that the DDoS policy works for.`,
				},
				resource.Attribute{
					Name:        "watermark_filters",
					Description: `Watermark policy options, and only support one watermark policy at most.`,
				},
				resource.Attribute{
					Name:        "auto_remove",
					Description: `Indicate whether to auto-remove the watermark or not.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `The offset of watermark.`,
				},
				resource.Attribute{
					Name:        "open_switch",
					Description: `Indicate whether to open watermark or not.`,
				},
				resource.Attribute{
					Name:        "tcp_port_list",
					Description: `Port range of TCP.`,
				},
				resource.Attribute{
					Name:        "udp_port_list",
					Description: `Port range of TCP.`,
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
			Type:             "tencentcloud_dayu_ddos_policy_attachments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of dayu DDoS policy attachments`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Type of the resource that the DDoS policy works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) Id of the policy to be queried.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Optional) Id of the attached resource to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dayu_ddos_policy_attachment_list",
					Description: `A list of dayu DDoS policy attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the policy.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `Id of the attached resource.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the DDoS policy works for.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dayu_ddos_policy_attachment_list",
					Description: `A list of dayu DDoS policy attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Id of the policy.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `Id of the attached resource.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the DDoS policy works for.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_ddos_policy_cases",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query dayu DDoS policy cases`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Type of the resource that the DDoS policy case works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `(Required) Id of the DDoS policy case to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of DDoS policy cases. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "app_protocols",
					Description: `App protocol set of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "app_type",
					Description: `App type of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "has_abroad",
					Description: `Indicate whether the service involves overseas or not.`,
				},
				resource.Attribute{
					Name:        "has_initiate_tcp",
					Description: `Indicate whether the service actively initiates TCP requests or not.`,
				},
				resource.Attribute{
					Name:        "has_initiate_udp",
					Description: `Indicate whether the actively initiate UDP requests or not.`,
				},
				resource.Attribute{
					Name:        "has_vpn",
					Description: `Indicate whether the service involves VPN service or not.`,
				},
				resource.Attribute{
					Name:        "max_tcp_package_len",
					Description: `The max length of TCP message package.`,
				},
				resource.Attribute{
					Name:        "max_udp_package_len",
					Description: `The max length of UDP message package.`,
				},
				resource.Attribute{
					Name:        "min_tcp_package_len",
					Description: `The minimum length of TCP message package.`,
				},
				resource.Attribute{
					Name:        "min_udp_package_len",
					Description: `The minimum length of UDP message package.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "peer_tcp_port",
					Description: `The port that actively initiates TCP requests.`,
				},
				resource.Attribute{
					Name:        "peer_udp_port",
					Description: `The port that actively initiates UDP requests.`,
				},
				resource.Attribute{
					Name:        "platform_types",
					Description: `Platform set of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the DDoS policy case works for.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `Id of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "tcp_end_port",
					Description: `End port of the TCP service.`,
				},
				resource.Attribute{
					Name:        "tcp_footprint",
					Description: `The fixed signature of TCP protocol load.`,
				},
				resource.Attribute{
					Name:        "tcp_start_port",
					Description: `Start port of the TCP service.`,
				},
				resource.Attribute{
					Name:        "udp_end_port",
					Description: `End port of the UDP service.`,
				},
				resource.Attribute{
					Name:        "udp_footprint",
					Description: `The fixed signature of TCP protocol load.`,
				},
				resource.Attribute{
					Name:        "udp_start_port",
					Description: `Start port of the UDP service.`,
				},
				resource.Attribute{
					Name:        "web_api_urls",
					Description: `Web API url set.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of DDoS policy cases. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "app_protocols",
					Description: `App protocol set of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "app_type",
					Description: `App type of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "has_abroad",
					Description: `Indicate whether the service involves overseas or not.`,
				},
				resource.Attribute{
					Name:        "has_initiate_tcp",
					Description: `Indicate whether the service actively initiates TCP requests or not.`,
				},
				resource.Attribute{
					Name:        "has_initiate_udp",
					Description: `Indicate whether the actively initiate UDP requests or not.`,
				},
				resource.Attribute{
					Name:        "has_vpn",
					Description: `Indicate whether the service involves VPN service or not.`,
				},
				resource.Attribute{
					Name:        "max_tcp_package_len",
					Description: `The max length of TCP message package.`,
				},
				resource.Attribute{
					Name:        "max_udp_package_len",
					Description: `The max length of UDP message package.`,
				},
				resource.Attribute{
					Name:        "min_tcp_package_len",
					Description: `The minimum length of TCP message package.`,
				},
				resource.Attribute{
					Name:        "min_udp_package_len",
					Description: `The minimum length of UDP message package.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "peer_tcp_port",
					Description: `The port that actively initiates TCP requests.`,
				},
				resource.Attribute{
					Name:        "peer_udp_port",
					Description: `The port that actively initiates UDP requests.`,
				},
				resource.Attribute{
					Name:        "platform_types",
					Description: `Platform set of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of the resource that the DDoS policy case works for.`,
				},
				resource.Attribute{
					Name:        "scene_id",
					Description: `Id of the DDoS policy case.`,
				},
				resource.Attribute{
					Name:        "tcp_end_port",
					Description: `End port of the TCP service.`,
				},
				resource.Attribute{
					Name:        "tcp_footprint",
					Description: `The fixed signature of TCP protocol load.`,
				},
				resource.Attribute{
					Name:        "tcp_start_port",
					Description: `Start port of the TCP service.`,
				},
				resource.Attribute{
					Name:        "udp_end_port",
					Description: `End port of the UDP service.`,
				},
				resource.Attribute{
					Name:        "udp_footprint",
					Description: `The fixed signature of TCP protocol load.`,
				},
				resource.Attribute{
					Name:        "udp_start_port",
					Description: `Start port of the UDP service.`,
				},
				resource.Attribute{
					Name:        "web_api_urls",
					Description: `Web API url set.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_l4_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query dayu layer 4 rules`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) Id of the resource that the layer 4 rule works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Type of the resource that the layer 4 rule works for, valid values are ` + "`" + `bgpip` + "`" + `, ` + "`" + `bgp` + "`" + `, ` + "`" + `bgp-multip` + "`" + ` and ` + "`" + `net` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the layer 4 rule to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional) Id of the layer 4 rule to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of layer 4 rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "d_port",
					Description: `The destination port of the layer 4 rule.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `Interval time of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `HTTP Status Code. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealthy threshold of health check.`,
				},
				resource.Attribute{
					Name:        "lb_type",
					Description: `LB type of the rule, 1 for weight cycling and 2 for IP hash.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the rule.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of the 4 layer rule.`,
				},
				resource.Attribute{
					Name:        "s_port",
					Description: `The source port of the layer 4 rule.`,
				},
				resource.Attribute{
					Name:        "session_switch",
					Description: `Indicate that the session will keep or not.`,
				},
				resource.Attribute{
					Name:        "session_time",
					Description: `Session keep time, only valid when ` + "`" + `session_switch` + "`" + ` is true, the available value ranges from 1 to 300 and unit is second.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `Source type, 1 for source of host, 2 for source of ip.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of layer 4 rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "d_port",
					Description: `The destination port of the layer 4 rule.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `Interval time of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `HTTP Status Code. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealthy threshold of health check.`,
				},
				resource.Attribute{
					Name:        "lb_type",
					Description: `LB type of the rule, 1 for weight cycling and 2 for IP hash.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the rule.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of the 4 layer rule.`,
				},
				resource.Attribute{
					Name:        "s_port",
					Description: `The source port of the layer 4 rule.`,
				},
				resource.Attribute{
					Name:        "session_switch",
					Description: `Indicate that the session will keep or not.`,
				},
				resource.Attribute{
					Name:        "session_time",
					Description: `Session keep time, only valid when ` + "`" + `session_switch` + "`" + ` is true, the available value ranges from 1 to 300 and unit is second.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `Source type, 1 for source of host, 2 for source of ip.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dayu_l7_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query dayu layer 7 rules`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) Id of the resource that the layer 7 rule works for.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Type of the resource that the layer 7 rule works for, valid value is ` + "`" + `bgpip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Domain of the layer 7 rule to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional) Id of the layer 7 rule to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of layer 7 rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain that the 7 layer rule works for.`,
				},
				resource.Attribute{
					Name:        "health_check_code",
					Description: `HTTP Status Code. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `Interval time of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `Methods of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `Path of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealthy threshold of health check.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the rule.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of the 7 layer rule.`,
				},
				resource.Attribute{
					Name:        "source_list",
					Description: `Source list of the rule.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `Source type, 1 for source of host, 2 for source of ip.`,
				},
				resource.Attribute{
					Name:        "ssl_id",
					Description: `SSL id.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the rule. 0 for create/modify success, 2 for create/modify fail, 3 for delete success, 5 for waiting to be created/modified, 7 for waiting to be deleted and 8 for waiting to get SSL id.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Indicate the rule will take effect or not.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Threshold of the rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of layer 7 rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain that the 7 layer rule works for.`,
				},
				resource.Attribute{
					Name:        "health_check_code",
					Description: `HTTP Status Code. 1 means the return value '1xx' is health. 2 means the return value '2xx' is health. 4 means the return value '3xx' is health. 8 means the return value '4xx' is health. 16 means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add the corresponding values.`,
				},
				resource.Attribute{
					Name:        "health_check_health_num",
					Description: `Health threshold of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `Interval time of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `Methods of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `Path of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_switch",
					Description: `Indicates whether health check is enabled.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealth_num",
					Description: `Unhealthy threshold of health check.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the rule.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of the 7 layer rule.`,
				},
				resource.Attribute{
					Name:        "source_list",
					Description: `Source list of the rule.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `Source type, 1 for source of host, 2 for source of ip.`,
				},
				resource.Attribute{
					Name:        "ssl_id",
					Description: `SSL id.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the rule. 0 for create/modify success, 2 for create/modify fail, 3 for delete success, 5 for waiting to be created/modified, 7 for waiting to be deleted and 8 for waiting to get SSL id.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `Indicate the rule will take effect or not.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Threshold of the rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dc_gateway_ccn_routes",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of direct connect gateway route entries.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dcg_id",
					Description: `(Required) ID of the DCG to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the DCG route entries.`,
				},
				resource.Attribute{
					Name:        "as_path",
					Description: `As_Path list of the BGP.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address segment of IDC.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `ID of the DCG.`,
				},
				resource.Attribute{
					Name:        "route_id",
					Description: `ID of the DCG route.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the DCG route entries.`,
				},
				resource.Attribute{
					Name:        "as_path",
					Description: `As_Path list of the BGP.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address segment of IDC.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `ID of the DCG.`,
				},
				resource.Attribute{
					Name:        "route_id",
					Description: `ID of the DCG route.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dc_gateway_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of direct connect gateway instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dcg_id",
					Description: `(Optional) ID of the DCG to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the DCG to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the DCG.`,
				},
				resource.Attribute{
					Name:        "cnn_route_type",
					Description: `Type of CCN route, the available value include 'BGP' and 'STATIC'.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `ID of the DCG.`,
				},
				resource.Attribute{
					Name:        "dcg_ip",
					Description: `IP of the DCG.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `Indicates whether the BGP is enabled.`,
				},
				resource.Attribute{
					Name:        "gateway_type",
					Description: `Type of the gateway, the available value include 'NORMAL' and 'NAT'. Default is 'NORMAL'.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DCG.`,
				},
				resource.Attribute{
					Name:        "network_instance_id",
					Description: `Type of associated network, the available value include 'VPC' and 'CCN'.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `IP of the DCG.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the DCG.`,
				},
				resource.Attribute{
					Name:        "cnn_route_type",
					Description: `Type of CCN route, the available value include 'BGP' and 'STATIC'.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `ID of the DCG.`,
				},
				resource.Attribute{
					Name:        "dcg_ip",
					Description: `IP of the DCG.`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `Indicates whether the BGP is enabled.`,
				},
				resource.Attribute{
					Name:        "gateway_type",
					Description: `Type of the gateway, the available value include 'NORMAL' and 'NAT'. Default is 'NORMAL'.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DCG.`,
				},
				resource.Attribute{
					Name:        "network_instance_id",
					Description: `Type of associated network, the available value include 'VPC' and 'CCN'.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `IP of the DCG.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dc_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of DC instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dc_id",
					Description: `(Optional) ID of the DC to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the DC to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the DC.`,
				},
				resource.Attribute{
					Name:        "access_point_id",
					Description: `Access point ID of tne DC.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Bandwidth of the DC.`,
				},
				resource.Attribute{
					Name:        "circuit_code",
					Description: `The circuit code provided by the operator for the DC.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "customer_address",
					Description: `Interconnect IP of the DC within client. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "customer_email",
					Description: `Applicant email of the DC, the default is obtained from the account. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Applicant name of the DC, the default is obtained from the account. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "customer_phone",
					Description: `Applicant phone number of the DC, the default is obtained from the account. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `ID of the DC.`,
				},
				resource.Attribute{
					Name:        "enabled_time",
					Description: `Enable time of resource.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expire date of resource.`,
				},
				resource.Attribute{
					Name:        "fault_report_contact_person",
					Description: `Contact of reporting a faulty. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "fault_report_contact_phone",
					Description: `Phone number of reporting a faulty. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "line_operator",
					Description: `Operator of the DC, and available values include ChinaTelecom, ChinaMobile, ChinaUnicom, In-houseWiring, ChinaOther and InternationalOperator.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The DC location where the connection is located.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DC.`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `Port type of the DC in client, and available values include 100Base-T, 1000Base-T, 1000Base-LX, 10GBase-T and 10GBase-LR. The default value is 1000Base-LX.`,
				},
				resource.Attribute{
					Name:        "redundant_dc_id",
					Description: `ID of the redundant DC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the DC, and available values include REJECTED, TOPAY, PAID, ALLOCATED, AVAILABLE, DELETING and DELETED.`,
				},
				resource.Attribute{
					Name:        "tencent_address",
					Description: `Interconnect IP of the DC within Tencent. Note: This field may return null, indicating that no valid values are taken.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the DC.`,
				},
				resource.Attribute{
					Name:        "access_point_id",
					Description: `Access point ID of tne DC.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Bandwidth of the DC.`,
				},
				resource.Attribute{
					Name:        "circuit_code",
					Description: `The circuit code provided by the operator for the DC.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "customer_address",
					Description: `Interconnect IP of the DC within client. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "customer_email",
					Description: `Applicant email of the DC, the default is obtained from the account. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Applicant name of the DC, the default is obtained from the account. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "customer_phone",
					Description: `Applicant phone number of the DC, the default is obtained from the account. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `ID of the DC.`,
				},
				resource.Attribute{
					Name:        "enabled_time",
					Description: `Enable time of resource.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expire date of resource.`,
				},
				resource.Attribute{
					Name:        "fault_report_contact_person",
					Description: `Contact of reporting a faulty. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "fault_report_contact_phone",
					Description: `Phone number of reporting a faulty. Note: This field may return null, indicating that no valid values are taken.`,
				},
				resource.Attribute{
					Name:        "line_operator",
					Description: `Operator of the DC, and available values include ChinaTelecom, ChinaMobile, ChinaUnicom, In-houseWiring, ChinaOther and InternationalOperator.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The DC location where the connection is located.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the DC.`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `Port type of the DC in client, and available values include 100Base-T, 1000Base-T, 1000Base-LX, 10GBase-T and 10GBase-LR. The default value is 1000Base-LX.`,
				},
				resource.Attribute{
					Name:        "redundant_dc_id",
					Description: `ID of the redundant DC.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the DC, and available values include REJECTED, TOPAY, PAID, ALLOCATED, AVAILABLE, DELETING and DELETED.`,
				},
				resource.Attribute{
					Name:        "tencent_address",
					Description: `Interconnect IP of the DC within Tencent. Note: This field may return null, indicating that no valid values are taken.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dcx_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of dedicated tunnels instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dcx_id",
					Description: `(Optional) ID of the dedicated tunnels to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the dedicated tunnels to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the dedicated tunnels.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Bandwidth of the DC.`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `BGP ASN of the user.`,
				},
				resource.Attribute{
					Name:        "bgp_auth_key",
					Description: `BGP key of the user.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "customer_address",
					Description: `Interconnect IP of the DC within client.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `ID of the DC.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `ID of the DC Gateway. Currently only new in the console.`,
				},
				resource.Attribute{
					Name:        "dcx_id",
					Description: `ID of the dedicated tunnel.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the dedicated tunnel.`,
				},
				resource.Attribute{
					Name:        "network_region",
					Description: `The region of the dedicated tunnel.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Type of the network, and available values include VPC, BMVPC and CCN. The default value is VPC.`,
				},
				resource.Attribute{
					Name:        "route_filter_prefixes",
					Description: `Static route, the network address of the user IDC.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `Type of the route, and available values include BGP and STATIC. The default value is BGP.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the dedicated tunnels, and available values include PENDING, ALLOCATING, ALLOCATED, ALTERING, DELETING, DELETED, COMFIRMING and REJECTED.`,
				},
				resource.Attribute{
					Name:        "tencent_address",
					Description: `Interconnect IP of the DC within Tencent.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `Vlan of the dedicated tunnels, and the range of values is [0-3000]. '0' means that only one tunnel can be created for the physical connect.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC or BMVPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `Information list of the dedicated tunnels.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Bandwidth of the DC.`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `BGP ASN of the user.`,
				},
				resource.Attribute{
					Name:        "bgp_auth_key",
					Description: `BGP key of the user.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of resource.`,
				},
				resource.Attribute{
					Name:        "customer_address",
					Description: `Interconnect IP of the DC within client.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `ID of the DC.`,
				},
				resource.Attribute{
					Name:        "dcg_id",
					Description: `ID of the DC Gateway. Currently only new in the console.`,
				},
				resource.Attribute{
					Name:        "dcx_id",
					Description: `ID of the dedicated tunnel.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the dedicated tunnel.`,
				},
				resource.Attribute{
					Name:        "network_region",
					Description: `The region of the dedicated tunnel.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Type of the network, and available values include VPC, BMVPC and CCN. The default value is VPC.`,
				},
				resource.Attribute{
					Name:        "route_filter_prefixes",
					Description: `Static route, the network address of the user IDC.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `Type of the route, and available values include BGP and STATIC. The default value is BGP.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the dedicated tunnels, and available values include PENDING, ALLOCATING, ALLOCATED, ALTERING, DELETING, DELETED, COMFIRMING and REJECTED.`,
				},
				resource.Attribute{
					Name:        "tencent_address",
					Description: `Interconnect IP of the DC within Tencent.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `Vlan of the dedicated tunnels, and the range of values is [0-3000]. '0' means that only one tunnel can be created for the physical connect.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC or BMVPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_dnats",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of DNATs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the NAT forward.`,
				},
				resource.Attribute{
					Name:        "elastic_ip",
					Description: `(Optional) Network address of the EIP.`,
				},
				resource.Attribute{
					Name:        "elastic_port",
					Description: `(Optional) Port of the EIP.`,
				},
				resource.Attribute{
					Name:        "nat_id",
					Description: `(Optional) Id of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) Network address of the backend service.`,
				},
				resource.Attribute{
					Name:        "private_port",
					Description: `(Optional) Port of intranet.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Id of the VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dnat_list",
					Description: `Information list of the DNATs.`,
				},
				resource.Attribute{
					Name:        "elastic_ip",
					Description: `Network address of the EIP.`,
				},
				resource.Attribute{
					Name:        "elastic_port",
					Description: `Port of the EIP.`,
				},
				resource.Attribute{
					Name:        "nat_id",
					Description: `Id of the NAT.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Network address of the backend service.`,
				},
				resource.Attribute{
					Name:        "private_port",
					Description: `Port of intranet.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Type of the network protocol, the available values include: ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Id of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dnat_list",
					Description: `Information list of the DNATs.`,
				},
				resource.Attribute{
					Name:        "elastic_ip",
					Description: `Network address of the EIP.`,
				},
				resource.Attribute{
					Name:        "elastic_port",
					Description: `Port of the EIP.`,
				},
				resource.Attribute{
					Name:        "nat_id",
					Description: `Id of the NAT.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Network address of the backend service.`,
				},
				resource.Attribute{
					Name:        "private_port",
					Description: `Port of intranet.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Type of the network protocol, the available values include: ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Id of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eip",
			Category:         "Data Sources",
			ShortDescription: `Provides an available EIP for the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to filter.`,
				},
				resource.Attribute{
					Name:        "include_arrears",
					Description: `(Optional) Whether the IP is arrears.`,
				},
				resource.Attribute{
					Name:        "include_blocked",
					Description: `(Optional) Whether the IP is blocked. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Key of the filter, valid keys: ` + "`" + `address-id` + "`" + `,` + "`" + `address-name` + "`" + `,` + "`" + `address-ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Value of the filter. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An EIP id indicate the uniqueness of a certain EIP, which can be used for instance binding or network interface binding.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `An public IP address for the EIP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the EIP, there are several status like ` + "`" + `BIND` + "`" + `, ` + "`" + `UNBIND` + "`" + `, and ` + "`" + `BIND_ENI` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An EIP id indicate the uniqueness of a certain EIP, which can be used for instance binding or network interface binding.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `An public IP address for the EIP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the EIP, there are several status like ` + "`" + `BIND` + "`" + `, ` + "`" + `UNBIND` + "`" + `, and ` + "`" + `BIND_ENI` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_eips",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query eip instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "eip_id",
					Description: `(Optional) ID of the eip to be queried.`,
				},
				resource.Attribute{
					Name:        "eip_name",
					Description: `(Optional) Name of the eip to be queried.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) The elastic ip address.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of eip. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "eip_list",
					Description: `An information list of eip. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the eip.`,
				},
				resource.Attribute{
					Name:        "eip_id",
					Description: `ID of the eip.`,
				},
				resource.Attribute{
					Name:        "eip_name",
					Description: `Name of the eip.`,
				},
				resource.Attribute{
					Name:        "eip_type",
					Description: `Type of the eip.`,
				},
				resource.Attribute{
					Name:        "eni_id",
					Description: `The eni id to bind with the eip.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance id to bind with the eip.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The elastic ip address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The eip current status.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the eip.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "eip_list",
					Description: `An information list of eip. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the eip.`,
				},
				resource.Attribute{
					Name:        "eip_id",
					Description: `ID of the eip.`,
				},
				resource.Attribute{
					Name:        "eip_name",
					Description: `Name of the eip.`,
				},
				resource.Attribute{
					Name:        "eip_type",
					Description: `Type of the eip.`,
				},
				resource.Attribute{
					Name:        "eni_id",
					Description: `The eni id to bind with the eip.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance id to bind with the eip.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The elastic ip address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The eip current status.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the eip.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_elasticsearch_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query elasticsearch instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) ID of the instance to be queried.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) Name of the instance to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tag of the instance to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `An information list of elasticsearch instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "basic_security_type",
					Description: `Whether to enable X-Pack security authentication in Basic Edition 6.8 and above.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Instance creation time.`,
				},
				resource.Attribute{
					Name:        "deploy_mode",
					Description: `Cluster deployment mode.`,
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
					Name:        "instance_id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the instance.`,
				},
				resource.Attribute{
					Name:        "kibana_url",
					Description: `Kibana access URL.`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `License type.`,
				},
				resource.Attribute{
					Name:        "multi_zone_infos",
					Description: `Details of AZs in multi-AZ deployment mode.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The id of a VPC subnetwork.`,
				},
				resource.Attribute{
					Name:        "node_info_list",
					Description: `Node information list, which describe the specification information of various types of nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Node disk size.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Node disk type.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `Number of nodes.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `Node specification.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Node type.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The id of a VPC subnetwork.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the instance.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The id of a VPC network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `An information list of elasticsearch instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "basic_security_type",
					Description: `Whether to enable X-Pack security authentication in Basic Edition 6.8 and above.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Instance creation time.`,
				},
				resource.Attribute{
					Name:        "deploy_mode",
					Description: `Cluster deployment mode.`,
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
					Name:        "instance_id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the instance.`,
				},
				resource.Attribute{
					Name:        "kibana_url",
					Description: `Kibana access URL.`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `License type.`,
				},
				resource.Attribute{
					Name:        "multi_zone_infos",
					Description: `Details of AZs in multi-AZ deployment mode.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The id of a VPC subnetwork.`,
				},
				resource.Attribute{
					Name:        "node_info_list",
					Description: `Node information list, which describe the specification information of various types of nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Node disk size.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Node disk type.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `Number of nodes.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `Node specification.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Node type.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The id of a VPC subnetwork.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the instance.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The id of a VPC network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_enis",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query query ENIs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the ENI. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) ID of the ENIs to be queried. Conflict with ` + "`" + `vpc_id` + "`" + `,` + "`" + `subnet_id` + "`" + `,` + "`" + `instance_id` + "`" + `,` + "`" + `security_group` + "`" + `,` + "`" + `name` + "`" + `,` + "`" + `ipv4` + "`" + ` and ` + "`" + `tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) ID of the instance which bind the ENI. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) Intranet IP of the ENI. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the ENI to be queried. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `(Optional) A set of security group IDs which bind the ENI. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of the subnet within this vpc to be queried. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the ENI. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the vpc to be queried. Conflict with ` + "`" + `ids` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "enis",
					Description: `An information list of ENIs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the ENI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the ENI.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ENI.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instance which bind the ENI.`,
				},
				resource.Attribute{
					Name:        "ipv4s",
					Description: `A set of intranet IPv4s.`,
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
					Name:        "name",
					Description: `Name of the ENI.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Indicates whether the IP is primary.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `A set of security group IDs which bind the ENI.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of the ENI.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet within this vpc.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the ENI.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "enis",
					Description: `An information list of ENIs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the ENI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the ENI.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ENI.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instance which bind the ENI.`,
				},
				resource.Attribute{
					Name:        "ipv4s",
					Description: `A set of intranet IPv4s.`,
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
					Name:        "name",
					Description: `Name of the ENI.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Indicates whether the IP is primary.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `A set of security group IDs which bind the ENI.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `States of the ENI.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet within this vpc.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the ENI.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_certificates",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query GAAP certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the certificate to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the certificate to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of the certificate to be queried, the available values include ` + "`" + `BASIC` + "`" + `, ` + "`" + `CLIENT` + "`" + `, ` + "`" + `SERVER` + "`" + `, ` + "`" + `REALSERVER` + "`" + ` and ` + "`" + `PROXY` + "`" + `; ` + "`" + `BASIC` + "`" + ` means basic certificate; ` + "`" + `CLIENT` + "`" + ` means client CA certificate; ` + "`" + `SERVER` + "`" + ` means server SSL certificate; ` + "`" + `REALSERVER` + "`" + ` means realserver CA certificate; ` + "`" + `PROXY` + "`" + ` means proxy SSL certificate. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `An information list of certificate. Each element contains the following attributes:`,
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
					Name:        "id",
					Description: `ID of the certificate.`,
				},
				resource.Attribute{
					Name:        "issuer_cn",
					Description: `Issuer name of the certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the certificate.`,
				},
				resource.Attribute{
					Name:        "subject_cn",
					Description: `Subject name of the certificate.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the certificate.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "certificates",
					Description: `An information list of certificate. Each element contains the following attributes:`,
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
					Name:        "id",
					Description: `ID of the certificate.`,
				},
				resource.Attribute{
					Name:        "issuer_cn",
					Description: `Issuer name of the certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the certificate.`,
				},
				resource.Attribute{
					Name:        "subject_cn",
					Description: `Subject name of the certificate.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_domain_error_pages",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query custom GAAP HTTP domain error page info list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) HTTP domain to be queried.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) ID of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) List of the error page info ID to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "error_page_info_list",
					Description: `An information list of error page info detail. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `New response body.`,
				},
				resource.Attribute{
					Name:        "clear_headers",
					Description: `Response headers to be removed.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `HTTP domain.`,
				},
				resource.Attribute{
					Name:        "error_codes",
					Description: `Original error codes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the error page info.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "new_error_codes",
					Description: `New error code.`,
				},
				resource.Attribute{
					Name:        "set_headers",
					Description: `Response headers to be set.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "error_page_info_list",
					Description: `An information list of error page info detail. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `New response body.`,
				},
				resource.Attribute{
					Name:        "clear_headers",
					Description: `Response headers to be removed.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `HTTP domain.`,
				},
				resource.Attribute{
					Name:        "error_codes",
					Description: `Original error codes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the error page info.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "new_error_codes",
					Description: `New error code.`,
				},
				resource.Attribute{
					Name:        "set_headers",
					Description: `Response headers to be set.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_http_domains",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query forward domain of layer7 listeners.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Forward domain of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) ID of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `An information list of forward domain of the layer7 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "basic_auth_id",
					Description: `ID of the basic authentication.`,
				},
				resource.Attribute{
					Name:        "basic_auth",
					Description: `Indicates whether basic authentication is enable.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `ID of the server certificate.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "client_certificate_ids",
					Description: `ID list of the client certificate.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Forward domain of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "gaap_auth_id",
					Description: `ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "gaap_auth",
					Description: `Indicates whether SSL certificate authentication is enable.`,
				},
				resource.Attribute{
					Name:        "realserver_auth",
					Description: `Indicates whether realserver authentication is enable.`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_domain",
					Description: `CA certificate domain of the realserver.`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_ids",
					Description: `CA certificate ID list of the realserver.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domains",
					Description: `An information list of forward domain of the layer7 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "basic_auth_id",
					Description: `ID of the basic authentication.`,
				},
				resource.Attribute{
					Name:        "basic_auth",
					Description: `Indicates whether basic authentication is enable.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `ID of the server certificate.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "client_certificate_ids",
					Description: `ID list of the client certificate.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Forward domain of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "gaap_auth_id",
					Description: `ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "gaap_auth",
					Description: `Indicates whether SSL certificate authentication is enable.`,
				},
				resource.Attribute{
					Name:        "realserver_auth",
					Description: `Indicates whether realserver authentication is enable.`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_domain",
					Description: `CA certificate domain of the realserver.`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "realserver_certificate_ids",
					Description: `CA certificate ID list of the realserver.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_http_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query forward rule of layer7 listeners.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) ID of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Forward domain of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "forward_host",
					Description: `(Optional) Requested host which is forwarded to the realserver by the listener to be queried.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path of the forward rule to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `An information list of forward rule of the layer7 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `Timeout of the health check response.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Forward domain of the forward rule.`,
				},
				resource.Attribute{
					Name:        "forward_host",
					Description: `Requested host which is forwarded to the realserver by the listener.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `Method of the health check.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `Path of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_status_codes",
					Description: `Return code of confirmed normal.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Indicates whether health check is enable.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the forward rule.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of the health check.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Path of the forward rule.`,
				},
				resource.Attribute{
					Name:        "realserver_type",
					Description: `Type of the realserver.`,
				},
				resource.Attribute{
					Name:        "realservers",
					Description: `An information list of GAAP realserver. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Scheduling weight.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling policy of the forward rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rules",
					Description: `An information list of forward rule of the layer7 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `Timeout of the health check response.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Forward domain of the forward rule.`,
				},
				resource.Attribute{
					Name:        "forward_host",
					Description: `Requested host which is forwarded to the realserver by the listener.`,
				},
				resource.Attribute{
					Name:        "health_check_method",
					Description: `Method of the health check.`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `Path of health check.`,
				},
				resource.Attribute{
					Name:        "health_check_status_codes",
					Description: `Return code of confirmed normal.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Indicates whether health check is enable.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the forward rule.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of the health check.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Path of the forward rule.`,
				},
				resource.Attribute{
					Name:        "realserver_type",
					Description: `Type of the realserver.`,
				},
				resource.Attribute{
					Name:        "realservers",
					Description: `An information list of GAAP realserver. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Scheduling weight.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling policy of the forward rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_layer4_listeners",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query gaap layer4 listeners.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol of the layer4 listener to be queried, the available values include ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Optional) ID of the layer4 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `(Optional) Name of the layer4 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port of the layer4 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `(Optional) ID of the GAAP proxy to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `An information list of layer4 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `Timeout of the health check response.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Indicates whether health check is enable.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of the health check.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "realserver_type",
					Description: `Type of the realserver.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling policy of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer4 listener.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listeners",
					Description: `An information list of layer4 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `Timeout of the health check response.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Indicates whether health check is enable.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of the health check.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "realserver_type",
					Description: `Type of the realserver.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Scheduling policy of the layer4 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer4 listener.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_layer7_listeners",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query gaap layer7 listeners.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol of the layer7 listener to be queried, the available values include ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Optional) ID of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `(Optional) Name of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port of the layer7 listener to be queried.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `(Optional) ID of the GAAP proxy to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `An information list of layer7 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `Authentication type of the layer7 listener. ` + "`" + `0` + "`" + ` is one-way authentication and ` + "`" + `1` + "`" + ` is mutual authentication.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `Certificate ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "client_certificate_ids",
					Description: `ID list of the client certificate.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "forward_protocol",
					Description: `Protocol type of the forwarding.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer7 listener.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listeners",
					Description: `An information list of layer7 listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `Authentication type of the layer7 listener. ` + "`" + `0` + "`" + ` is one-way authentication and ` + "`" + `1` + "`" + ` is mutual authentication.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `Certificate ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "client_certificate_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "client_certificate_ids",
					Description: `ID list of the client certificate.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "forward_protocol",
					Description: `Protocol type of the forwarding.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the layer7 listener.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the layer7 listener.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_proxies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query gaap proxies.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_region",
					Description: `(Optional) Access region of the GAAP proxy to be queried. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) ID of the GAAP proxy to be queried. Conflict with ` + "`" + `project_id` + "`" + `, ` + "`" + `access_region` + "`" + ` and ` + "`" + `realserver_region` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID of the GAAP proxy to be queried. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "realserver_region",
					Description: `(Optional) Region of the GAAP realserver to be queried. Conflict with ` + "`" + `ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the GAAP proxy to be queried. Support up to 5, display the information as long as it matches one. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "proxies",
					Description: `An information list of GAAP proxy. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_region",
					Description: `Access region of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Maximum bandwidth of the GAAP proxy, unit is Mbps.`,
				},
				resource.Attribute{
					Name:        "concurrent",
					Description: `Maximum concurrency of the GAAP proxy, unit is 10k.`,
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
					Name:        "id",
					Description: `ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Access domain of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Security policy ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project within the GAAP proxy, '0' means is default project.`,
				},
				resource.Attribute{
					Name:        "realserver_region",
					Description: `Region of the GAAP realserver.`,
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
					Description: `Supported protocols of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the GAAP proxy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "proxies",
					Description: `An information list of GAAP proxy. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_region",
					Description: `Access region of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Maximum bandwidth of the GAAP proxy, unit is Mbps.`,
				},
				resource.Attribute{
					Name:        "concurrent",
					Description: `Maximum concurrency of the GAAP proxy, unit is 10k.`,
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
					Name:        "id",
					Description: `ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Access domain of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Security policy ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project within the GAAP proxy, '0' means is default project.`,
				},
				resource.Attribute{
					Name:        "realserver_region",
					Description: `Region of the GAAP realserver.`,
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
					Description: `Supported protocols of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the GAAP proxy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_realservers",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query gaap realservers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Domain of the GAAP realserver to be queried, conflict with ` + "`" + `ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) IP of the GAAP realserver to be queried, conflict with ` + "`" + `domain` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the GAAP realserver to be queried, the maximum length is 30.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project within the GAAP realserver to be queried, default value is ` + "`" + `-1` + "`" + `, no set means all projects.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the GAAP proxy to be queried. Support up to 5, display the information as long as it matches one. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "realservers",
					Description: `An information list of GAAP realserver. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project within the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the GAAP realserver.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "realservers",
					Description: `An information list of GAAP realserver. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project within the GAAP realserver.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the GAAP realserver.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_security_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query security policies of GAAP proxy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID of the security policy to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Default policy.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the security policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `Default policy.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `ID of the GAAP proxy.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the security policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_gaap_security_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query security policy rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) ID of the security policy to be queried.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Policy of the rule to be queried.`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `(Optional) A network address block of the request source to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the security policy rule to be queried.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port of the security policy rule to be queried.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol of the security policy rule to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional) ID of the security policy rules to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `An information list of security policy rule. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Policy of the rule.`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `A network address block of the request source.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the security policy rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security policy rule.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the security policy rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the security policy rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rules",
					Description: `An information list of security policy rule. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Policy of the rule.`,
				},
				resource.Attribute{
					Name:        "cidr_ip",
					Description: `A network address block of the request source.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the security policy rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security policy rule.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the security policy rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the security policy rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ha_vip_eip_attachments",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of HA VIP EIP attachments`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "havip_id",
					Description: `(Required) Id of the attached HA VIP to be queried.`,
				},
				resource.Attribute{
					Name:        "address_ip",
					Description: `(Optional) Public IP address of EIP to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ha_vip_eip_attachment_list",
					Description: `A list of HA VIP EIP attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "address_ip",
					Description: `Public IP address of EIP.`,
				},
				resource.Attribute{
					Name:        "havip_id",
					Description: `Id of the attached HA VIP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ha_vip_eip_attachment_list",
					Description: `A list of HA VIP EIP attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "address_ip",
					Description: `Public IP address of EIP.`,
				},
				resource.Attribute{
					Name:        "havip_id",
					Description: `Id of the attached HA VIP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ha_vips",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of HA VIPs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address_ip",
					Description: `(Optional) EIP of the HA VIP to be queried.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Id of the HA VIP to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the HA VIP. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Subnet id of the HA VIP to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) VPC id of the HA VIP to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ha_vip_list",
					Description: `Information list of the dedicated HA VIPs.`,
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
					Name:        "id",
					Description: `Id of the HA VIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Instance id that is associated.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the HA VIP.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `Network interface id that is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the HA VIP, values are ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `UNBIND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet id.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `Virtual IP address, it must not be occupied and in this VPC network segment. If not set, it will be assigned after resource created automatically.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ha_vip_list",
					Description: `Information list of the dedicated HA VIPs.`,
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
					Name:        "id",
					Description: `Id of the HA VIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Instance id that is associated.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the HA VIP.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `Network interface id that is associated.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the HA VIP, values are ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `UNBIND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet id.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `Virtual IP address, it must not be occupied and in this VPC network segment. If not set, it will be assigned after resource created automatically.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_image",
			Category:         "Data Sources",
			ShortDescription: `Provides an available image for the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to filter.`,
				},
				resource.Attribute{
					Name:        "image_name_regex",
					Description: `(Optional) A regex string to apply to the image list returned by TencentCloud.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `(Optional) A string to apply with fuzzy match to the os_name attribute on the image list returned by TencentCloud.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Key of the filter, valid keys: ` + "`" + `image-id` + "`" + `, ` + "`" + `image-type` + "`" + `, ` + "`" + `image-name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Values of the filter. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `An image id indicate the uniqueness of a certain image, which can be used for instance creation or resetting.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Name of this image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "image_id",
					Description: `An image id indicate the uniqueness of a certain image, which can be used for instance creation or resetting.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Name of this image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_images",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query images.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) ID of the image to be queried.`,
				},
				resource.Attribute{
					Name:        "image_name_regex",
					Description: `(Optional) A regex string to apply to the image list returned by TencentCloud, conflict with 'os_name'.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `(Optional) A list of the image type to be queried. Available values include: 'PUBLIC_IMAGE', 'PRIVATE_IMAGE', 'SHARED_IMAGE', 'MARKET_IMAGE'.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `(Optional) A string to apply with fuzzy match to the os_name attribute on the image list returned by TencentCloud, conflict with 'image_name_regex'.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `An information list of image. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `Architecture of the image.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Created time of the image.`,
				},
				resource.Attribute{
					Name:        "image_creator",
					Description: `Image creator of the image.`,
				},
				resource.Attribute{
					Name:        "image_description",
					Description: `Description of the image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Name of the image.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Size of the image.`,
				},
				resource.Attribute{
					Name:        "image_source",
					Description: `Image source of the image.`,
				},
				resource.Attribute{
					Name:        "image_state",
					Description: `State of the image.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `Type of the image.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `OS name of the image.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `Platform of the image.`,
				},
				resource.Attribute{
					Name:        "support_cloud_init",
					Description: `Whether support cloud-init.`,
				},
				resource.Attribute{
					Name:        "sync_percent",
					Description: `Sync percent of the image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "images",
					Description: `An information list of image. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `Architecture of the image.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Created time of the image.`,
				},
				resource.Attribute{
					Name:        "image_creator",
					Description: `Image creator of the image.`,
				},
				resource.Attribute{
					Name:        "image_description",
					Description: `Description of the image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Name of the image.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Size of the image.`,
				},
				resource.Attribute{
					Name:        "image_source",
					Description: `Image source of the image.`,
				},
				resource.Attribute{
					Name:        "image_state",
					Description: `State of the image.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `Type of the image.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `OS name of the image.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `Platform of the image.`,
				},
				resource.Attribute{
					Name:        "support_cloud_init",
					Description: `Whether support cloud-init.`,
				},
				resource.Attribute{
					Name:        "sync_percent",
					Description: `Sync percent of the image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_instance_types",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query instances type.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the CVM instance locates at. This field is conflict with ` + "`" + `filter` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `(Optional) The number of CPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to filter. This field is conflict with ` + "`" + `availability_zone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gpu_core_count",
					Description: `(Optional) The number of GPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `(Optional) Instance memory capacity, unit in GB.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. The ` + "`" + `filter` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The filter name, the available values include ` + "`" + `zone` + "`" + ` and ` + "`" + `instance-family` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) The filter values. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `An information list of cvm instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the CVM instance locates at.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `The number of CPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `Type series of the instance.`,
				},
				resource.Attribute{
					Name:        "gpu_core_count",
					Description: `The number of GPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Type of the instance.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Instance memory capacity, unit in GB.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_types",
					Description: `An information list of cvm instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the CVM instance locates at.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `The number of CPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `Type series of the instance.`,
				},
				resource.Attribute{
					Name:        "gpu_core_count",
					Description: `The number of GPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Type of the instance.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Instance memory capacity, unit in GB.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query cvm instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the CVM instance locates at.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) ID of the instances to be queried.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) Name of the instances to be queried.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The project CVM belongs to.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of a vpc subnetwork.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the vpc to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `An information list of cvm instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "allocate_public_ip",
					Description: `Indicates whether public ip is assigned.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the CVM instance locates at.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `The number of CPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `An information list of data disk. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "data_disk_id",
					Description: `Image ID of the data disk.`,
				},
				resource.Attribute{
					Name:        "data_disk_size",
					Description: `Size of the data disk.`,
				},
				resource.Attribute{
					Name:        "data_disk_type",
					Description: `Type of the data disk.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `Indicates whether the data disk is destroyed with the instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `The way that CVM instance will be renew automatically or not when it reach the end of the prepaid tenancy.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `The charge type of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instances.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the instances.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Type of the instance.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `The charge type of the instance.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Public network maximum output bandwidth of the instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Instance memory capacity, unit in GB.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private ip of the instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project CVM belongs to.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public ip of the instance.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `Security groups of the instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of a vpc subnetwork.`,
				},
				resource.Attribute{
					Name:        "system_disk_id",
					Description: `Image ID of the system disk.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `Size of the system disk.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `Type of the system disk.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `An information list of cvm instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "allocate_public_ip",
					Description: `Indicates whether public ip is assigned.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The available zone that the CVM instance locates at.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `The number of CPU cores of the instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `An information list of data disk. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "data_disk_id",
					Description: `Image ID of the data disk.`,
				},
				resource.Attribute{
					Name:        "data_disk_size",
					Description: `Size of the data disk.`,
				},
				resource.Attribute{
					Name:        "data_disk_type",
					Description: `Type of the data disk.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `Indicates whether the data disk is destroyed with the instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type_prepaid_renew_flag",
					Description: `The way that CVM instance will be renew automatically or not when it reach the end of the prepaid tenancy.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `The charge type of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instances.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the instances.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Type of the instance.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `The charge type of the instance.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Public network maximum output bandwidth of the instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Instance memory capacity, unit in GB.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private ip of the instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project CVM belongs to.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public ip of the instance.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `Security groups of the instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of a vpc subnetwork.`,
				},
				resource.Attribute{
					Name:        "system_disk_id",
					Description: `Image ID of the system disk.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `Size of the system disk.`,
				},
				resource.Attribute{
					Name:        "system_disk_type",
					Description: `Type of the system disk.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_key_pairs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query key pairs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) ID of the key pair to be queried.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional) Name of the key pair to be queried. Support regular expression search, only ` + "`" + `^` + "`" + ` and ` + "`" + `$` + "`" + ` are supported.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project id of the key pair to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "key_pair_list",
					Description: `An information list of key pair. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the key pair.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `ID of the key pair.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `Name of the key pair.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project id of the key pair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `public key of the key pair.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_pair_list",
					Description: `An information list of key pair. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the key pair.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `ID of the key pair.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `Name of the key pair.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project id of the key pair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `public key of the key pair.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_kubernetes_clusters",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of kubernetes clusters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) ID of the cluster. Conflict with cluster_name, can not be set at the same time.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Optional) Name of the cluster. Conflict with cluster_id, can not be set at the same time.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the cluster. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `An information list of kubernetes clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "certification_authority",
					Description: `The certificate used for access.`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `A network address block of the cluster. Different from vpc cidr and cidr of other clusters within this vpc.`,
				},
				resource.Attribute{
					Name:        "cluster_deploy_type",
					Description: `Deployment type of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_desc",
					Description: `Description of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_external_endpoint",
					Description: `External network address to access.`,
				},
				resource.Attribute{
					Name:        "cluster_ipvs",
					Description: `Indicates whether ipvs is enabled.`,
				},
				resource.Attribute{
					Name:        "cluster_max_pod_num",
					Description: `The maximum number of Pods per node in the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_max_service_num",
					Description: `The maximum number of services in the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_node_num",
					Description: `Number of nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_os",
					Description: `Operating system of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `Version of the cluster.`,
				},
				resource.Attribute{
					Name:        "container_runtime",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name for access.`,
				},
				resource.Attribute{
					Name:        "ignore_cluster_cidr_conflict",
					Description: `Indicates whether to ignore the cluster cidr conflict error.`,
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
					Name:        "project_id",
					Description: `Project Id of the cluster.`,
				},
				resource.Attribute{
					Name:        "security_policy",
					Description: `Access policy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the cluster.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `User name of account.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Vpc Id of the cluster.`,
				},
				resource.Attribute{
					Name:        "worker_instances_list",
					Description: `An information list of cvm within the WORKER clusters. Each element contains the following attributes.`,
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
					Name:        "list",
					Description: `An information list of kubernetes clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "certification_authority",
					Description: `The certificate used for access.`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `A network address block of the cluster. Different from vpc cidr and cidr of other clusters within this vpc.`,
				},
				resource.Attribute{
					Name:        "cluster_deploy_type",
					Description: `Deployment type of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_desc",
					Description: `Description of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_external_endpoint",
					Description: `External network address to access.`,
				},
				resource.Attribute{
					Name:        "cluster_ipvs",
					Description: `Indicates whether ipvs is enabled.`,
				},
				resource.Attribute{
					Name:        "cluster_max_pod_num",
					Description: `The maximum number of Pods per node in the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_max_service_num",
					Description: `The maximum number of services in the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_node_num",
					Description: `Number of nodes in the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_os",
					Description: `Operating system of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `Version of the cluster.`,
				},
				resource.Attribute{
					Name:        "container_runtime",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name for access.`,
				},
				resource.Attribute{
					Name:        "ignore_cluster_cidr_conflict",
					Description: `Indicates whether to ignore the cluster cidr conflict error.`,
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
					Name:        "project_id",
					Description: `Project Id of the cluster.`,
				},
				resource.Attribute{
					Name:        "security_policy",
					Description: `Access policy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the cluster.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `User name of account.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Vpc Id of the cluster.`,
				},
				resource.Attribute{
					Name:        "worker_instances_list",
					Description: `An information list of cvm within the WORKER clusters. Each element contains the following attributes.`,
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
			Type:             "tencentcloud_mongodb_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of Mongodb instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Optional) Type of Mongodb cluster, and available values include replica set cluster(expressed with ` + "`" + `REPLSET` + "`" + `), sharding cluster(expressed with ` + "`" + `SHARD` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) ID of the Mongodb instance to be queried.`,
				},
				resource.Attribute{
					Name:        "instance_name_prefix",
					Description: `(Optional) Name prefix of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the Mongodb instance to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `The available zone of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Type of Mongodb cluster.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Number of cpu's core.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the Mongodb engine.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `Type of Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "shard_quantity",
					Description: `Number of sharding.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `Disk size.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `IP port of the Mongodb instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `The available zone of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Type of Mongodb cluster.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Number of cpu's core.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the Mongodb engine.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `Type of Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "shard_quantity",
					Description: `Number of sharding.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Mongodb, and available values include pending initialization(expressed with 0), processing(expressed with 1), running(expressed with 2) and expired(expressed with -2).`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `IP of the Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `Disk size.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vport",
					Description: `IP port of the Mongodb instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mongodb_zone_config",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the available mongodb specifications for different zone.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Optional) The available zone of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of zone config. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `The available zone of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Type of Mongodb cluster.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Number of cpu's core.`,
				},
				resource.Attribute{
					Name:        "default_storage",
					Description: `Default disk size.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the Mongodb version.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `Type of Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "max_storage",
					Description: `Maximum size of the disk.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "min_storage",
					Description: `Minimum sie of the disk.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of zone config. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `The available zone of the Mongodb.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Type of Mongodb cluster.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Number of cpu's core.`,
				},
				resource.Attribute{
					Name:        "default_storage",
					Description: `Default disk size.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Version of the Mongodb version.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `Type of Mongodb instance.`,
				},
				resource.Attribute{
					Name:        "max_storage",
					Description: `Maximum size of the disk.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "min_storage",
					Description: `Minimum sie of the disk.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_binding_objects",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query policy group binding objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Policy group id for query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list objects. Each element contains the following attributes:`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list objects. Each element contains the following attributes:`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_data",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query monitor data. for complex queries, use (https://github.com/tencentyun/tencentcloud-exporter)`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Required) Dimensional composition of instance objects.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Required) End time for this query, eg:` + "`" + `2018-09-22T20:00:00+08:00` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) Metric name, please refer to the documentation of monitor interface of each product.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Namespace of each cloud product in monitor system, refer to ` + "`" + `data.tencentcloud_monitor_product_namespace` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) Start time for this query, eg:` + "`" + `2018-09-22T19:51:23+08:00` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) Statistical period.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. The ` + "`" + `dimensions` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Instance dimension name, eg: ` + "`" + `InstanceId` + "`" + ` for cvm.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Instance dimension value, eg: ` + "`" + `ins-j0hk02zo` + "`" + ` for cvm. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list data point. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `Statistical timestamp.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Statistical value.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list data point. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `Statistical timestamp.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Statistical value.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_policy_conditions",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query monitor policy conditions(There is a lot of data and it is recommended to output to a file)`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the policy name, support partial matching, eg:` + "`" + `Cloud Virtual Machine` + "`" + `,` + "`" + `Virtual` + "`" + `,` + "`" + `Cloud Load Banlancer-Private CLB Listener` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list policy condition. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "event_metrics",
					Description: `A list of event condition metrics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `The id of this event metric.`,
				},
				resource.Attribute{
					Name:        "event_show_name",
					Description: `The name of this event metric.`,
				},
				resource.Attribute{
					Name:        "need_recovered",
					Description: `Whether to recover.`,
				},
				resource.Attribute{
					Name:        "is_support_multi_region",
					Description: `Whether to support multi region.`,
				},
				resource.Attribute{
					Name:        "metrics",
					Description: `A list of event condition metrics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "calc_type_keys",
					Description: `Calculate type of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_type_need",
					Description: `Whether ` + "`" + `calc_type` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "calc_value_default",
					Description: `The default calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_fixed",
					Description: `The fixed calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_max",
					Description: `The max calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_min",
					Description: `The min calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_need",
					Description: `Whether ` + "`" + `calc_value` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "continue_time_default",
					Description: `The default continue time(seconds) config for this metric.`,
				},
				resource.Attribute{
					Name:        "continue_time_keys",
					Description: `The continue time(seconds) keys for this metric.`,
				},
				resource.Attribute{
					Name:        "continue_time_need",
					Description: `Whether ` + "`" + `continue_time` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "metric_id",
					Description: `The id of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_show_name",
					Description: `The name of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_unit",
					Description: `The unit of this metric.`,
				},
				resource.Attribute{
					Name:        "period_default",
					Description: `The default data time(seconds) config for this metric.`,
				},
				resource.Attribute{
					Name:        "period_keys",
					Description: `The data time(seconds) keys for this metric.`,
				},
				resource.Attribute{
					Name:        "period_need",
					Description: `Whether ` + "`" + `period` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "period_num_default",
					Description: `The default period number config for this metric.`,
				},
				resource.Attribute{
					Name:        "period_num_keys",
					Description: `The period number keys for this metric.`,
				},
				resource.Attribute{
					Name:        "period_num_need",
					Description: `Whether ` + "`" + `period_num` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "stat_type_p10",
					Description: `Data aggregation mode, cycle of 10 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p1800",
					Description: `Data aggregation mode, cycle of 1800 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p300",
					Description: `Data aggregation mode, cycle of 300 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p3600",
					Description: `Data aggregation mode, cycle of 3600 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p5",
					Description: `Data aggregation mode, cycle of 5 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p600",
					Description: `Data aggregation mode, cycle of 600 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p60",
					Description: `Data aggregation mode, cycle of 60 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p86400",
					Description: `Data aggregation mode, cycle of 86400 seconds.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of this policy name.`,
				},
				resource.Attribute{
					Name:        "policy_view_name",
					Description: `Policy view name, eg:` + "`" + `cvm_device` + "`" + `,` + "`" + `BANDWIDTHPACKAGE` + "`" + `, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(policy_view_name)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "support_regions",
					Description: `Support regions of this policy view.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list policy condition. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "event_metrics",
					Description: `A list of event condition metrics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `The id of this event metric.`,
				},
				resource.Attribute{
					Name:        "event_show_name",
					Description: `The name of this event metric.`,
				},
				resource.Attribute{
					Name:        "need_recovered",
					Description: `Whether to recover.`,
				},
				resource.Attribute{
					Name:        "is_support_multi_region",
					Description: `Whether to support multi region.`,
				},
				resource.Attribute{
					Name:        "metrics",
					Description: `A list of event condition metrics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "calc_type_keys",
					Description: `Calculate type of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_type_need",
					Description: `Whether ` + "`" + `calc_type` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "calc_value_default",
					Description: `The default calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_fixed",
					Description: `The fixed calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_max",
					Description: `The max calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_min",
					Description: `The min calculate value of this metric.`,
				},
				resource.Attribute{
					Name:        "calc_value_need",
					Description: `Whether ` + "`" + `calc_value` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "continue_time_default",
					Description: `The default continue time(seconds) config for this metric.`,
				},
				resource.Attribute{
					Name:        "continue_time_keys",
					Description: `The continue time(seconds) keys for this metric.`,
				},
				resource.Attribute{
					Name:        "continue_time_need",
					Description: `Whether ` + "`" + `continue_time` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "metric_id",
					Description: `The id of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_show_name",
					Description: `The name of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_unit",
					Description: `The unit of this metric.`,
				},
				resource.Attribute{
					Name:        "period_default",
					Description: `The default data time(seconds) config for this metric.`,
				},
				resource.Attribute{
					Name:        "period_keys",
					Description: `The data time(seconds) keys for this metric.`,
				},
				resource.Attribute{
					Name:        "period_need",
					Description: `Whether ` + "`" + `period` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "period_num_default",
					Description: `The default period number config for this metric.`,
				},
				resource.Attribute{
					Name:        "period_num_keys",
					Description: `The period number keys for this metric.`,
				},
				resource.Attribute{
					Name:        "period_num_need",
					Description: `Whether ` + "`" + `period_num` + "`" + ` required in the configuration.`,
				},
				resource.Attribute{
					Name:        "stat_type_p10",
					Description: `Data aggregation mode, cycle of 10 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p1800",
					Description: `Data aggregation mode, cycle of 1800 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p300",
					Description: `Data aggregation mode, cycle of 300 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p3600",
					Description: `Data aggregation mode, cycle of 3600 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p5",
					Description: `Data aggregation mode, cycle of 5 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p600",
					Description: `Data aggregation mode, cycle of 600 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p60",
					Description: `Data aggregation mode, cycle of 60 seconds.`,
				},
				resource.Attribute{
					Name:        "stat_type_p86400",
					Description: `Data aggregation mode, cycle of 86400 seconds.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of this policy name.`,
				},
				resource.Attribute{
					Name:        "policy_view_name",
					Description: `Policy view name, eg:` + "`" + `cvm_device` + "`" + `,` + "`" + `BANDWIDTHPACKAGE` + "`" + `, refer to ` + "`" + `data.tencentcloud_monitor_policy_conditions(policy_view_name)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "support_regions",
					Description: `Support regions of this policy view.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_policy_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query monitor policy groups (There is a lot of data and it is recommended to output to a file)`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Policy group name for query.`,
				},
				resource.Attribute{
					Name:        "policy_view_names",
					Description: `(Optional) The policy view for query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list policy groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "can_set_default",
					Description: `Whether it can be set as the default policy.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `A list of threshold rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "alarm_notify_period",
					Description: `Alarm sending cycle per second.<0 does not fire, 0 only fires once, and >0 fires every triggerTime second.`,
				},
				resource.Attribute{
					Name:        "alarm_notify_type",
					Description: `Alarm sending convergence type. 0 continuous alarm, 1 index alarm.`,
				},
				resource.Attribute{
					Name:        "calc_type",
					Description: `Compare type, 1 means more than, 2 means greater than or equal, 3 means less than, 4 means less than or equal to, 5 means equal, 6 means not equal, 7 means days rose, 8 means days fell, 9 means weeks rose, 10 means weeks fell, 11 means period rise, 12 means period fell.`,
				},
				resource.Attribute{
					Name:        "calc_value",
					Description: `Threshold value.`,
				},
				resource.Attribute{
					Name:        "continue_time",
					Description: `How long does the triggering rule last (per second).`,
				},
				resource.Attribute{
					Name:        "metric_id",
					Description: `The id of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_show_name",
					Description: `The name of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_unit",
					Description: `The unit of this metric.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Data aggregation cycle (unit second).`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Threshold rule id.`,
				},
				resource.Attribute{
					Name:        "event_conditions",
					Description: `A list of event rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "alarm_notify_period",
					Description: `Alarm sending cycle per second.<0 does not fire, 0 only fires once, and >0 fires every triggerTime second.`,
				},
				resource.Attribute{
					Name:        "alarm_notify_type",
					Description: `Alarm sending convergence type. 0 continuous alarm, 1 index alarm.`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `The id of this event metric.`,
				},
				resource.Attribute{
					Name:        "event_show_name",
					Description: `The name of this event metric.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Threshold rule id.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The policy group id.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The policy group name.`,
				},
				resource.Attribute{
					Name:        "insert_time",
					Description: `The policy group create timestamp.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `If is default policy group or not,0 represents the non-default policy, and 1 represents the default policy.`,
				},
				resource.Attribute{
					Name:        "is_open",
					Description: `Whether open or not.`,
				},
				resource.Attribute{
					Name:        "last_edit_uin",
					Description: `Recently edited user uin.`,
				},
				resource.Attribute{
					Name:        "no_shielded_sum",
					Description: `Number of unmasked instances of policy group bindings.`,
				},
				resource.Attribute{
					Name:        "parent_group_id",
					Description: `Parent policy group id.`,
				},
				resource.Attribute{
					Name:        "policy_view_name",
					Description: `The policy group view name.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project id to which the policy group belongs.`,
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
					Description: `Method of warning notification.Optional ` + "`" + `CALL` + "`" + `,` + "`" + `EMAIL` + "`" + `,` + "`" + `SITE` + "`" + `,` + "`" + `SMS` + "`" + `,` + "`" + `WECHAT` + "`" + `.`,
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
					Description: `Receive type. Optional 'group' or 'user'.`,
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
					Description: `Alarm period start time.Range [0,86399], which removes the date after it is converted to Beijing time as a Unix timestamp, for example 7200 means '10:0:0'.`,
				},
				resource.Attribute{
					Name:        "uid_list",
					Description: `The phone alerts the receiver uid.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Policy group remarks.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The policy group update timestamp.`,
				},
				resource.Attribute{
					Name:        "use_sum",
					Description: `Number of instances of policy group bindings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list policy groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "can_set_default",
					Description: `Whether it can be set as the default policy.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `A list of threshold rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "alarm_notify_period",
					Description: `Alarm sending cycle per second.<0 does not fire, 0 only fires once, and >0 fires every triggerTime second.`,
				},
				resource.Attribute{
					Name:        "alarm_notify_type",
					Description: `Alarm sending convergence type. 0 continuous alarm, 1 index alarm.`,
				},
				resource.Attribute{
					Name:        "calc_type",
					Description: `Compare type, 1 means more than, 2 means greater than or equal, 3 means less than, 4 means less than or equal to, 5 means equal, 6 means not equal, 7 means days rose, 8 means days fell, 9 means weeks rose, 10 means weeks fell, 11 means period rise, 12 means period fell.`,
				},
				resource.Attribute{
					Name:        "calc_value",
					Description: `Threshold value.`,
				},
				resource.Attribute{
					Name:        "continue_time",
					Description: `How long does the triggering rule last (per second).`,
				},
				resource.Attribute{
					Name:        "metric_id",
					Description: `The id of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_show_name",
					Description: `The name of this metric.`,
				},
				resource.Attribute{
					Name:        "metric_unit",
					Description: `The unit of this metric.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Data aggregation cycle (unit second).`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Threshold rule id.`,
				},
				resource.Attribute{
					Name:        "event_conditions",
					Description: `A list of event rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "alarm_notify_period",
					Description: `Alarm sending cycle per second.<0 does not fire, 0 only fires once, and >0 fires every triggerTime second.`,
				},
				resource.Attribute{
					Name:        "alarm_notify_type",
					Description: `Alarm sending convergence type. 0 continuous alarm, 1 index alarm.`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `The id of this event metric.`,
				},
				resource.Attribute{
					Name:        "event_show_name",
					Description: `The name of this event metric.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Threshold rule id.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The policy group id.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The policy group name.`,
				},
				resource.Attribute{
					Name:        "insert_time",
					Description: `The policy group create timestamp.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `If is default policy group or not,0 represents the non-default policy, and 1 represents the default policy.`,
				},
				resource.Attribute{
					Name:        "is_open",
					Description: `Whether open or not.`,
				},
				resource.Attribute{
					Name:        "last_edit_uin",
					Description: `Recently edited user uin.`,
				},
				resource.Attribute{
					Name:        "no_shielded_sum",
					Description: `Number of unmasked instances of policy group bindings.`,
				},
				resource.Attribute{
					Name:        "parent_group_id",
					Description: `Parent policy group id.`,
				},
				resource.Attribute{
					Name:        "policy_view_name",
					Description: `The policy group view name.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project id to which the policy group belongs.`,
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
					Description: `Method of warning notification.Optional ` + "`" + `CALL` + "`" + `,` + "`" + `EMAIL` + "`" + `,` + "`" + `SITE` + "`" + `,` + "`" + `SMS` + "`" + `,` + "`" + `WECHAT` + "`" + `.`,
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
					Description: `Receive type. Optional 'group' or 'user'.`,
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
					Description: `Alarm period start time.Range [0,86399], which removes the date after it is converted to Beijing time as a Unix timestamp, for example 7200 means '10:0:0'.`,
				},
				resource.Attribute{
					Name:        "uid_list",
					Description: `The phone alerts the receiver uid.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Policy group remarks.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The policy group update timestamp.`,
				},
				resource.Attribute{
					Name:        "use_sum",
					Description: `Number of instances of policy group bindings.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_product_event",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query monitor events(There is a lot of data and it is recommended to output to a file)`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Optional) Dimensional composition of instance objects.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) End timestamp for this query, eg:` + "`" + `1588232111` + "`" + `. Default start time is ` + "`" + `now-3000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event_name",
					Description: `(Optional) Event name filtering, such as ` + "`" + `guest_reboot` + "`" + ` indicates that the machine restart.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) Affect objects, such as ` + "`" + `ins-19708ino` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_alarm_config",
					Description: `(Optional) Alarm status configuration filter, 1means configured, 0(default) means not configured.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `(Optional) Product type filtering, such as ` + "`" + `cvm` + "`" + ` for cloud server.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID filter.`,
				},
				resource.Attribute{
					Name:        "region_list",
					Description: `(Optional) Region filter, such as ` + "`" + `gz` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) Start timestamp for this query, eg:` + "`" + `1588230000` + "`" + `. Default start time is ` + "`" + `now-3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Event status filter, value range ` + "`" + `-` + "`" + `,` + "`" + `alarm` + "`" + `,` + "`" + `recover` + "`" + `, indicating recovered, unrecovered and stateless.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Event type filtering, with value range ` + "`" + `abnormal` + "`" + `,` + "`" + `status_change` + "`" + `, indicating state change and abnormal events. The ` + "`" + `dimensions` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Instance dimension name, eg: ` + "`" + `deviceWanIp` + "`" + ` for internet ip.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Instance dimension value, eg: ` + "`" + `119.119.119.119` + "`" + ` for internet ip. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list events. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "addition_msg",
					Description: `A list of addition message. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of this addition message.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this addition message.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of this addition message.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `A list of event dimensions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of this dimension.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this dimension.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of this dimension.`,
				},
				resource.Attribute{
					Name:        "event_cname",
					Description: `Event chinese name.`,
				},
				resource.Attribute{
					Name:        "event_ename",
					Description: `Event english name.`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `Event id.`,
				},
				resource.Attribute{
					Name:        "event_name",
					Description: `Event short name.`,
				},
				resource.Attribute{
					Name:        "group_info",
					Description: `A list of group info. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Policy group id.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `Policy group name.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance id of this event.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The name of this instance.`,
				},
				resource.Attribute{
					Name:        "is_alarm_config",
					Description: `Whether to configure alarm.`,
				},
				resource.Attribute{
					Name:        "product_cname",
					Description: `Product chinese name.`,
				},
				resource.Attribute{
					Name:        "product_ename",
					Description: `Product english name.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Product short name.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project id of this instance.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of this instance.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The start timestamp of this event.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this event.`,
				},
				resource.Attribute{
					Name:        "support_alarm",
					Description: `Whether to support alarm.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of this event.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The update timestamp of this event.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list events. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "addition_msg",
					Description: `A list of addition message. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of this addition message.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this addition message.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of this addition message.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `A list of event dimensions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of this dimension.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this dimension.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of this dimension.`,
				},
				resource.Attribute{
					Name:        "event_cname",
					Description: `Event chinese name.`,
				},
				resource.Attribute{
					Name:        "event_ename",
					Description: `Event english name.`,
				},
				resource.Attribute{
					Name:        "event_id",
					Description: `Event id.`,
				},
				resource.Attribute{
					Name:        "event_name",
					Description: `Event short name.`,
				},
				resource.Attribute{
					Name:        "group_info",
					Description: `A list of group info. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Policy group id.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `Policy group name.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance id of this event.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The name of this instance.`,
				},
				resource.Attribute{
					Name:        "is_alarm_config",
					Description: `Whether to configure alarm.`,
				},
				resource.Attribute{
					Name:        "product_cname",
					Description: `Product chinese name.`,
				},
				resource.Attribute{
					Name:        "product_ename",
					Description: `Product english name.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Product short name.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project id of this instance.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of this instance.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The start timestamp of this event.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this event.`,
				},
				resource.Attribute{
					Name:        "support_alarm",
					Description: `Whether to support alarm.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of this event.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The update timestamp of this event.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_monitor_product_namespace",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query product namespace in monitor)`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for filter, eg:` + "`" + `Load Banlancer` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list product namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace of each cloud product in monitor system.`,
				},
				resource.Attribute{
					Name:        "product_chinese_name",
					Description: `Chinese name of this product.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `English name of this product.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list product namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace of each cloud product in monitor system.`,
				},
				resource.Attribute{
					Name:        "product_chinese_name",
					Description: `Chinese name of this product.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `English name of this product.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_backup_list",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the list of backup databases.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mysql_id",
					Description: `(Required) Instance ID, such as cdb-c1nl9rpv. It is identical to the instance ID displayed in the database console page.`,
				},
				resource.Attribute{
					Name:        "max_number",
					Description: `(Optional) The latest files to list, rang from 1 to 10000. And the default value is 10.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of MySQL backup. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "backup_id",
					Description: `ID of Backup task.`,
				},
				resource.Attribute{
					Name:        "backup_model",
					Description: `Backup method. Supported values include: physical - physical backup, and logical - logical backup.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `The owner of the backup files.`,
				},
				resource.Attribute{
					Name:        "finish_time",
					Description: `The time at which the backup finishes.`,
				},
				resource.Attribute{
					Name:        "internet_url",
					Description: `URL for downloads externally.`,
				},
				resource.Attribute{
					Name:        "intranet_url",
					Description: `URL for downloads internally.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `the size of backup file.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `The earliest time at which the backup starts. For example, 2 indicates 2:00 am.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of MySQL backup. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "backup_id",
					Description: `ID of Backup task.`,
				},
				resource.Attribute{
					Name:        "backup_model",
					Description: `Backup method. Supported values include: physical - physical backup, and logical - logical backup.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `The owner of the backup files.`,
				},
				resource.Attribute{
					Name:        "finish_time",
					Description: `The time at which the backup finishes.`,
				},
				resource.Attribute{
					Name:        "internet_url",
					Description: `URL for downloads externally.`,
				},
				resource.Attribute{
					Name:        "intranet_url",
					Description: `URL for downloads internally.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `the size of backup file.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `The earliest time at which the backup starts. For example, 2 indicates 2:00 am.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_instance",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information about a MySQL instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) Pay type of instance, valid values are ` + "`" + `PREPAID` + "`" + ` and ` + "`" + `POSTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional) The version number of the database engine to use. Supported versions include 5.5/5.6/5.7.`,
				},
				resource.Attribute{
					Name:        "init_flag",
					Description: `(Optional) Initialization mark. Available values: 0 - Uninitialized; 1 - Initialized.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) Name of mysql instance.`,
				},
				resource.Attribute{
					Name:        "instance_role",
					Description: `(Optional) Instance type. Supported values include: master - master instance, dr - disaster recovery instance, and ro - read-only instance.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Number of results returned for a single request. Default is 20, and maximum is 2000.`,
				},
				resource.Attribute{
					Name:        "mysql_id",
					Description: `(Optional) Instance ID, such as cdb-c1nl9rpv. It is identical to the instance ID displayed in the database console page.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Record offset. Default is 0.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) Security groups ID of instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Instance status. Available values: 0 - Creating; 1 - Running; 4 - Isolating; 5 - Isolated.`,
				},
				resource.Attribute{
					Name:        "with_dr",
					Description: `(Optional) Indicates whether to query disaster recovery instances.`,
				},
				resource.Attribute{
					Name:        "with_master",
					Description: `(Optional) Indicates whether to query master instances.`,
				},
				resource.Attribute{
					Name:        "with_ro",
					Description: `(Optional) Indicates whether to query read-only instances. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `Auto renew flag. NOTES: Only supported prepay instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Pay type of instance.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `CPU count.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time at which a instance is created.`,
				},
				resource.Attribute{
					Name:        "dead_line_time",
					Description: `Expire date of instance. NOTES: Only supported prepay instance.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Supported instance model.HA - high available version; Basic - basic version.`,
				},
				resource.Attribute{
					Name:        "dr_instance_ids",
					Description: `ID list of disaster-recovery type associated with the current instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `The version number of the database engine to use. Supported versions include 5.5/5.6/5.7.`,
				},
				resource.Attribute{
					Name:        "init_flag",
					Description: `Initialization mark. Available values: 0 - Uninitialized; 1 - Initialized.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of mysql instance.`,
				},
				resource.Attribute{
					Name:        "instance_role",
					Description: `Instance type. Supported values include: master - master instance, dr - disaster recovery instance, and ro - read-only instance.`,
				},
				resource.Attribute{
					Name:        "internet_host",
					Description: `Public network domain name.`,
				},
				resource.Attribute{
					Name:        "internet_port",
					Description: `Public network port.`,
				},
				resource.Attribute{
					Name:        "internet_status",
					Description: `Status of public network.`,
				},
				resource.Attribute{
					Name:        "intranet_ip",
					Description: `Instance IP for internal access.`,
				},
				resource.Attribute{
					Name:        "intranet_port",
					Description: `Transport layer port number for internal purpose.`,
				},
				resource.Attribute{
					Name:        "master_instance_id",
					Description: `Indicates the master instance ID of recovery instances.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size (in MB).`,
				},
				resource.Attribute{
					Name:        "mysql_id",
					Description: `Instance ID, such as cdb-c1nl9rpv. It is identical to the instance ID displayed in the database console page.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `Pay type of instance, 0: prepaid, 1: postpaid.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID to which the current instance belongs.`,
				},
				resource.Attribute{
					Name:        "ro_instance_ids",
					Description: `ID list of read-only type associated with the current instance.`,
				},
				resource.Attribute{
					Name:        "slave_sync_mode",
					Description: `Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Available values: 0 - Creating; 1 - Running; 4 - Isolating; 5 - Isolated.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of subnet to which the current instance belongs.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `Disk capacity (in GB).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of Virtual Private Cloud.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Information of available zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "auto_renew_flag",
					Description: `Auto renew flag. NOTES: Only supported prepay instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Pay type of instance.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `CPU count.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time at which a instance is created.`,
				},
				resource.Attribute{
					Name:        "dead_line_time",
					Description: `Expire date of instance. NOTES: Only supported prepay instance.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Supported instance model.HA - high available version; Basic - basic version.`,
				},
				resource.Attribute{
					Name:        "dr_instance_ids",
					Description: `ID list of disaster-recovery type associated with the current instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `The version number of the database engine to use. Supported versions include 5.5/5.6/5.7.`,
				},
				resource.Attribute{
					Name:        "init_flag",
					Description: `Initialization mark. Available values: 0 - Uninitialized; 1 - Initialized.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of mysql instance.`,
				},
				resource.Attribute{
					Name:        "instance_role",
					Description: `Instance type. Supported values include: master - master instance, dr - disaster recovery instance, and ro - read-only instance.`,
				},
				resource.Attribute{
					Name:        "internet_host",
					Description: `Public network domain name.`,
				},
				resource.Attribute{
					Name:        "internet_port",
					Description: `Public network port.`,
				},
				resource.Attribute{
					Name:        "internet_status",
					Description: `Status of public network.`,
				},
				resource.Attribute{
					Name:        "intranet_ip",
					Description: `Instance IP for internal access.`,
				},
				resource.Attribute{
					Name:        "intranet_port",
					Description: `Transport layer port number for internal purpose.`,
				},
				resource.Attribute{
					Name:        "master_instance_id",
					Description: `Indicates the master instance ID of recovery instances.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size (in MB).`,
				},
				resource.Attribute{
					Name:        "mysql_id",
					Description: `Instance ID, such as cdb-c1nl9rpv. It is identical to the instance ID displayed in the database console page.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `Pay type of instance, 0: prepaid, 1: postpaid.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID to which the current instance belongs.`,
				},
				resource.Attribute{
					Name:        "ro_instance_ids",
					Description: `ID list of read-only type associated with the current instance.`,
				},
				resource.Attribute{
					Name:        "slave_sync_mode",
					Description: `Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Available values: 0 - Creating; 1 - Running; 4 - Isolating; 5 - Isolated.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of subnet to which the current instance belongs.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `Disk capacity (in GB).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of Virtual Private Cloud.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Information of available zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_parameter_list",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information about a parameter group of a database instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional) The version number of the database engine to use. Supported versions include 5.5/5.6/5.7.`,
				},
				resource.Attribute{
					Name:        "mysql_id",
					Description: `(Optional) Instance ID.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "parameter_list",
					Description: `A list of parameters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "current_value",
					Description: `Current value.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `Default value.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Parameter specification description.`,
				},
				resource.Attribute{
					Name:        "enum_value",
					Description: `Enumerated value.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `Maximum value for the parameter.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Minimum value for the parameter.`,
				},
				resource.Attribute{
					Name:        "need_reboot",
					Description: `Indicates whether reboot is needed to enable the new parameters.`,
				},
				resource.Attribute{
					Name:        "parameter_name",
					Description: `Parameter name.`,
				},
				resource.Attribute{
					Name:        "parameter_type",
					Description: `Parameter type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "parameter_list",
					Description: `A list of parameters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "current_value",
					Description: `Current value.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `Default value.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Parameter specification description.`,
				},
				resource.Attribute{
					Name:        "enum_value",
					Description: `Enumerated value.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `Maximum value for the parameter.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Minimum value for the parameter.`,
				},
				resource.Attribute{
					Name:        "need_reboot",
					Description: `Indicates whether reboot is needed to enable the new parameters.`,
				},
				resource.Attribute{
					Name:        "parameter_name",
					Description: `Parameter name.`,
				},
				resource.Attribute{
					Name:        "parameter_type",
					Description: `Parameter type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_mysql_zone_config",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the available database specifications for different regions. And a maximum of 20 requests can be initiated per second for this query.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region parameter, which is used to identify the region to which the data you want to work with belongs.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of zone config. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "disaster_recovery_zones",
					Description: `Information about available zones of recovery.`,
				},
				resource.Attribute{
					Name:        "engine_versions",
					Description: `The version number of the database engine to use. Supported versions include 5.5/5.6/5.7.`,
				},
				resource.Attribute{
					Name:        "first_slave_zones",
					Description: `Zone information about first slave instance.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether the current DC is the default DC for the region. Possible returned values: 0 - No; 1 - Yes.`,
				},
				resource.Attribute{
					Name:        "is_support_disaster_recovery",
					Description: `Indicates whether recovery is supported: 0 - No; 1 - Yes.`,
				},
				resource.Attribute{
					Name:        "is_support_vpc",
					Description: `Indicates whether VPC is supported: 0 - No; 1 - Yes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of available zone which is equal to a specific datacenter.`,
				},
				resource.Attribute{
					Name:        "second_slave_zones",
					Description: `Zone information about second slave instance.`,
				},
				resource.Attribute{
					Name:        "sells",
					Description: `A list of supported instance types for sell:`,
				},
				resource.Attribute{
					Name:        "max_volume_size",
					Description: `Maximum disk size (in GB).`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `Memory size (in MB).`,
				},
				resource.Attribute{
					Name:        "min_volume_size",
					Description: `Minimum disk size (in GB).`,
				},
				resource.Attribute{
					Name:        "qps",
					Description: `Queries per second.`,
				},
				resource.Attribute{
					Name:        "volume_step",
					Description: `Disk increment (in GB).`,
				},
				resource.Attribute{
					Name:        "slave_deploy_modes",
					Description: `Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.`,
				},
				resource.Attribute{
					Name:        "support_slave_sync_modes",
					Description: `Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of zone config. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "disaster_recovery_zones",
					Description: `Information about available zones of recovery.`,
				},
				resource.Attribute{
					Name:        "engine_versions",
					Description: `The version number of the database engine to use. Supported versions include 5.5/5.6/5.7.`,
				},
				resource.Attribute{
					Name:        "first_slave_zones",
					Description: `Zone information about first slave instance.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether the current DC is the default DC for the region. Possible returned values: 0 - No; 1 - Yes.`,
				},
				resource.Attribute{
					Name:        "is_support_disaster_recovery",
					Description: `Indicates whether recovery is supported: 0 - No; 1 - Yes.`,
				},
				resource.Attribute{
					Name:        "is_support_vpc",
					Description: `Indicates whether VPC is supported: 0 - No; 1 - Yes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of available zone which is equal to a specific datacenter.`,
				},
				resource.Attribute{
					Name:        "second_slave_zones",
					Description: `Zone information about second slave instance.`,
				},
				resource.Attribute{
					Name:        "sells",
					Description: `A list of supported instance types for sell:`,
				},
				resource.Attribute{
					Name:        "max_volume_size",
					Description: `Maximum disk size (in GB).`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `Memory size (in MB).`,
				},
				resource.Attribute{
					Name:        "min_volume_size",
					Description: `Minimum disk size (in GB).`,
				},
				resource.Attribute{
					Name:        "qps",
					Description: `Queries per second.`,
				},
				resource.Attribute{
					Name:        "volume_step",
					Description: `Disk increment (in GB).`,
				},
				resource.Attribute{
					Name:        "slave_deploy_modes",
					Description: `Availability zone deployment method. Available values: 0 - Single availability zone; 1 - Multiple availability zones.`,
				},
				resource.Attribute{
					Name:        "support_slave_sync_modes",
					Description: `Data replication mode. 0 - Async replication; 1 - Semisync replication; 2 - Strongsync replication.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_nat_gateways",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of NAT gateways.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Id of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Id of the VPC. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nats",
					Description: `Information list of the dedicated NATs.`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `EIP IP address set bound to the gateway. The value of at least 1.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of NAT gateway (unit: Mbps), the available values include: 20,50,100,200,500,1000,2000,5000. Default is 100.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `The upper limit of concurrent connection of NAT gateway, the available values include: 1000000,3000000,10000000. Default is 1000000.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Id of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nats",
					Description: `Information list of the dedicated NATs.`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `EIP IP address set bound to the gateway. The value of at least 1.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of NAT gateway (unit: Mbps), the available values include: 20,50,100,200,500,1000,2000,5000. Default is 100.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `The upper limit of concurrent connection of NAT gateway, the available values include: 1000000,3000000,10000000. Default is 1000000.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Id of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_nats",
			Category:         "Data Sources",
			ShortDescription: `The NATs data source lists a number of NATs resource information owned by an TencentCloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The maximum public network output bandwidth of the gateway (unit: Mbps), for example: 10, 20, 50, 100, 200, 500, 1000, 2000, 5000.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `(Optional) The upper limit of concurrent connection of NAT gateway, for example: 1000000, 3000000, 10000000.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) NAT gateway status, 0: Running, 1: Unavailable, 2: Be in arrears and out of service.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The VPC ID for NAT Gateway. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nats",
					Description: `Information list of the dedicated tunnels.`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `Elastic IP arrays bound to the gateway.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of the gateway (unit: Mbps), for example: 10, 20, 50, 100, 200, 500, 1000, 2000, 5000.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `The upper limit of concurrent connection of NAT gateway, for example: 1000000, 3000000, 10000000.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `NAT gateway status, 0: Running, 1: Unavailable, 2: Be in arrears and out of service.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID for NAT Gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nats",
					Description: `Information list of the dedicated tunnels.`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `Elastic IP arrays bound to the gateway.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of the gateway (unit: Mbps), for example: 10, 20, 50, 100, 200, 500, 1000, 2000, 5000.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `The upper limit of concurrent connection of NAT gateway, for example: 1000000, 3000000, 10000000.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `NAT gateway status, 0: Running, 1: Unavailable, 2: Be in arrears and out of service.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID for NAT Gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_placement_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query placement groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the placement group to be queried.`,
				},
				resource.Attribute{
					Name:        "placement_group_id",
					Description: `(Optional) ID of the placement group to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "placement_group_list",
					Description: `An information list of placement group. Each element contains the following attributes:`,
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
					Description: `Maximum number of hosts in the placement group.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `Host IDs in the placement group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the placement group.`,
				},
				resource.Attribute{
					Name:        "placement_group_id",
					Description: `ID of the placement group.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the placement group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "placement_group_list",
					Description: `An information list of placement group. Each element contains the following attributes:`,
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
					Description: `Maximum number of hosts in the placement group.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `Host IDs in the placement group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the placement group.`,
				},
				resource.Attribute{
					Name:        "placement_group_id",
					Description: `ID of the placement group.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the placement group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_redis_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query the detail information of redis instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) The number limitation of results for a query.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project to which redis instance belongs.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "search_key",
					Description: `(Optional) Key words used to match the results, and the key words can be: instance ID, instance name and IP address.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of redis instance.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) ID of an available zone. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of redis instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of instance. Valid values are ` + "`" + `POSTPAID` + "`" + ` and ` + "`" + `PREPAID` + "`" + `. Default value is ` + "`" + `POSTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the instance is created.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of an instance.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `Memory size in MB.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a redis instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port used to access a redis instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project to which a redis instance belongs.`,
				},
				resource.Attribute{
					Name:        "redis_id",
					Description: `ID of a redis instance.`,
				},
				resource.Attribute{
					Name:        "redis_replicas_num",
					Description: `The number of instance copies.`,
				},
				resource.Attribute{
					Name:        "redis_shard_num",
					Description: `The number of instance shard.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of an instance, maybe: init, processing, online, isolate and todelete.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the vpc subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of an instance.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `Instance type. Refer to ` + "`" + `data.tencentcloud_redis_zone_config.list.type_id` + "`" + ` get available values.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc with which the instance is associated.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Available zone to which a redis instance belongs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of redis instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of instance. Valid values are ` + "`" + `POSTPAID` + "`" + ` and ` + "`" + `PREPAID` + "`" + `. Default value is ` + "`" + `POSTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the instance is created.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of an instance.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `Memory size in MB.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a redis instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port used to access a redis instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the project to which a redis instance belongs.`,
				},
				resource.Attribute{
					Name:        "redis_id",
					Description: `ID of a redis instance.`,
				},
				resource.Attribute{
					Name:        "redis_replicas_num",
					Description: `The number of instance copies.`,
				},
				resource.Attribute{
					Name:        "redis_shard_num",
					Description: `The number of instance shard.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of an instance, maybe: init, processing, online, isolate and todelete.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the vpc subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of an instance.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `Instance type. Refer to ` + "`" + `data.tencentcloud_redis_zone_config.list.type_id` + "`" + ` get available values.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the vpc with which the instance is associated.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Available zone to which a redis instance belongs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_redis_zone_config",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query which instance types of Redis are available in a specific region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Name of a region. If this value is not set, the current region getting from provider's configuration will be used.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `(Optional) Instance type id. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of zone. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "mem_sizes",
					Description: `The memory volume of an available instance(in MB).`,
				},
				resource.Attribute{
					Name:        "redis_replicas_nums",
					Description: `The support numbers of instance copies.`,
				},
				resource.Attribute{
					Name:        "redis_shard_nums",
					Description: `The support numbers of instance shard.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `Instance type. Which redis type supports in this zone.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version description of an available instance. Possible values: Redis 3.2, Redis 4.0.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `ID of available zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of zone. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "mem_sizes",
					Description: `The memory volume of an available instance(in MB).`,
				},
				resource.Attribute{
					Name:        "redis_replicas_nums",
					Description: `The support numbers of instance copies.`,
				},
				resource.Attribute{
					Name:        "redis_shard_nums",
					Description: `The support numbers of instance shard.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `Instance type. Which redis type supports in this zone.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version description of an available instance. Possible values: Redis 3.2, Redis 4.0.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `ID of available zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_reserved_instance_configs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query reserved instances configuration.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the reserved instance locates at.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) Validity period of the reserved instance. Valid values are 31536000(1 year) and 94608000(3 years).`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The type of reserved instance.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "config_list",
					Description: `An information list of reserved instance configuration. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone of the purchasable reserved instance.`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `Configuration ID of the purchasable reserved instance.`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Settlement currency of the reserved instance, which is a standard currency code as listed in ISO 4217.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Validity period of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `Platform of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Purchase price of the reserved instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "config_list",
					Description: `An information list of reserved instance configuration. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone of the purchasable reserved instance.`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `Configuration ID of the purchasable reserved instance.`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Settlement currency of the reserved instance, which is a standard currency code as listed in ISO 4217.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Validity period of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `Platform of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Purchase price of the reserved instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_reserved_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query reserved instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The available zone that the reserved instance locates at.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The type of reserved instance.`,
				},
				resource.Attribute{
					Name:        "reserved_instance_id",
					Description: `(Optional) ID of the reserved instance to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "reserved_instance_list",
					Description: `An information list of reserved instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Expiry time of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of reserved instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of reserved instance.`,
				},
				resource.Attribute{
					Name:        "reserved_instance_id",
					Description: `ID of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Start time of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the reserved instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "reserved_instance_list",
					Description: `An information list of reserved instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `Expiry time of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of reserved instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of reserved instance.`,
				},
				resource.Attribute{
					Name:        "reserved_instance_id",
					Description: `ID of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Start time of the reserved instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the reserved instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_route_table",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Route Table.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required) The Route Table ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Route Table name. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of routing table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `The information list of the VPC route table.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The RouteEntry's target network segment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The RouteEntry's description.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `The RouteEntry's next hub.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `The ` + "`" + `next_hub` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "subnet_num",
					Description: `Number of associated subnets.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of routing table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `The information list of the VPC route table.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The RouteEntry's target network segment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The RouteEntry's description.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `The RouteEntry's next hub.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `The ` + "`" + `next_hub` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "subnet_num",
					Description: `Number of associated subnets.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_scf_functions",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query SCF functions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the SCF function to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the SCF function to be queried.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace of the SCF function to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the SCF function to be queried, can use up to 10 tags. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "functions",
					Description: `An information list of functions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cls_logset_id",
					Description: `CLS logset ID of the SCF function.`,
				},
				resource.Attribute{
					Name:        "cls_topic_id",
					Description: `CLS topic ID of the SCF function.`,
				},
				resource.Attribute{
					Name:        "code_error",
					Description: `Code error of the SCF function.`,
				},
				resource.Attribute{
					Name:        "code_result",
					Description: `Code result of the SCF function.`,
				},
				resource.Attribute{
					Name:        "code_size",
					Description: `Code size of the SCF function.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SCF function.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the SCF function.`,
				},
				resource.Attribute{
					Name:        "eip_fixed",
					Description: `Whether EIP is a fixed IP.`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `EIP list of the SCF function.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Environment variable of the SCF function.`,
				},
				resource.Attribute{
					Name:        "err_no",
					Description: `Errno of the SCF function.`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `Handler of the SCF function.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host of the SCF function.`,
				},
				resource.Attribute{
					Name:        "install_dependency",
					Description: `Whether to automatically install dependencies.`,
				},
				resource.Attribute{
					Name:        "l5_enable",
					Description: `Whether to enable L5.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `Memory size of the SCF function runtime, unit is M.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of the SCF function.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SCF function.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace of the SCF function.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `CAM role of the SCF function.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `Runtime of the SCF function.`,
				},
				resource.Attribute{
					Name:        "status_desc",
					Description: `Status description of the SCF function.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SCF function.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet ID of the SCF function.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the SCF function.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout of the SCF function maximum execution time, unit is second.`,
				},
				resource.Attribute{
					Name:        "trigger_info",
					Description: `Trigger details list the SCF function. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "custom_argument",
					Description: `user-defined parameter of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Whether to enable SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "trigger_desc",
					Description: `TriggerDesc of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `Vip of the SCF function.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID of the SCF function.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "functions",
					Description: `An information list of functions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cls_logset_id",
					Description: `CLS logset ID of the SCF function.`,
				},
				resource.Attribute{
					Name:        "cls_topic_id",
					Description: `CLS topic ID of the SCF function.`,
				},
				resource.Attribute{
					Name:        "code_error",
					Description: `Code error of the SCF function.`,
				},
				resource.Attribute{
					Name:        "code_result",
					Description: `Code result of the SCF function.`,
				},
				resource.Attribute{
					Name:        "code_size",
					Description: `Code size of the SCF function.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SCF function.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the SCF function.`,
				},
				resource.Attribute{
					Name:        "eip_fixed",
					Description: `Whether EIP is a fixed IP.`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `EIP list of the SCF function.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `Environment variable of the SCF function.`,
				},
				resource.Attribute{
					Name:        "err_no",
					Description: `Errno of the SCF function.`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `Handler of the SCF function.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host of the SCF function.`,
				},
				resource.Attribute{
					Name:        "install_dependency",
					Description: `Whether to automatically install dependencies.`,
				},
				resource.Attribute{
					Name:        "l5_enable",
					Description: `Whether to enable L5.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `Memory size of the SCF function runtime, unit is M.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of the SCF function.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SCF function.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Namespace of the SCF function.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `CAM role of the SCF function.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `Runtime of the SCF function.`,
				},
				resource.Attribute{
					Name:        "status_desc",
					Description: `Status description of the SCF function.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SCF function.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet ID of the SCF function.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the SCF function.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Timeout of the SCF function maximum execution time, unit is second.`,
				},
				resource.Attribute{
					Name:        "trigger_info",
					Description: `Trigger details list the SCF function. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "custom_argument",
					Description: `user-defined parameter of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Whether to enable SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "trigger_desc",
					Description: `TriggerDesc of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the SCF function trigger.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `Vip of the SCF function.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID of the SCF function.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_scf_logs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query SCF function logs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "function_name",
					Description: `(Required) Name of the SCF function to be queried.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) The end time of the query, the format is ` + "`" + `2017-05-16 20:00:00` + "`" + `, which can only be within one day from ` + "`" + `start_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "invoke_request_id",
					Description: `(Optional) Corresponding requestId when executing function.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Number of logs, the default is ` + "`" + `10000` + "`" + `, offset+limit cannot be greater than 10000.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace of the SCF function to be queried.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Log offset, default is ` + "`" + `0` + "`" + `, offset+limit cannot be greater than 10000.`,
				},
				resource.Attribute{
					Name:        "order_by",
					Description: `(Optional) Sort the logs according to the following fields: ` + "`" + `function_name` + "`" + `, ` + "`" + `duration` + "`" + `, ` + "`" + `mem_usage` + "`" + `, ` + "`" + `start_time` + "`" + `, default ` + "`" + `start_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Order to sort the log, optional values ` + "`" + `desc` + "`" + ` and ` + "`" + `asc` + "`" + `, default ` + "`" + `desc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "ret_code",
					Description: `(Optional) Use to filter log, optional value: ` + "`" + `not0` + "`" + ` only returns the error log. ` + "`" + `is0` + "`" + ` only returns the correct log. ` + "`" + `TimeLimitExceeded` + "`" + ` returns the log of the function call timeout. ` + "`" + `ResourceLimitExceeded` + "`" + ` returns the function call generation resource overrun log. ` + "`" + `UserCodeException` + "`" + ` returns logs of the user code error that occurred in the function call. Not passing the parameter means returning all logs.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) The start time of the query, the format is ` + "`" + `2017-05-16 20:00:00` + "`" + `, which can only be within one day from ` + "`" + `end_time` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "logs",
					Description: `An information list of logs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "bill_duration",
					Description: `Function billing time, according to duration up to the last 100ms, unit is ms.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Function execution time-consuming, unit is ms.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `Name of the SCF function.`,
				},
				resource.Attribute{
					Name:        "invoke_finished",
					Description: `Whether the function call ends, ` + "`" + `1` + "`" + ` means the execution ends, other values indicate the call exception.`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `Log level.`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `Log output during function execution.`,
				},
				resource.Attribute{
					Name:        "mem_usage",
					Description: `The actual memory size consumed in the execution of the function, unit is Byte.`,
				},
				resource.Attribute{
					Name:        "request_id",
					Description: `Execute the requestId corresponding to the function.`,
				},
				resource.Attribute{
					Name:        "ret_code",
					Description: `Execution result of function, ` + "`" + `0` + "`" + ` means the execution is successful, other values indicate failure.`,
				},
				resource.Attribute{
					Name:        "ret_msg",
					Description: `Return value after function execution is completed.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Log source.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Point in time at which the function begins execution.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "logs",
					Description: `An information list of logs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "bill_duration",
					Description: `Function billing time, according to duration up to the last 100ms, unit is ms.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Function execution time-consuming, unit is ms.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `Name of the SCF function.`,
				},
				resource.Attribute{
					Name:        "invoke_finished",
					Description: `Whether the function call ends, ` + "`" + `1` + "`" + ` means the execution ends, other values indicate the call exception.`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `Log level.`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `Log output during function execution.`,
				},
				resource.Attribute{
					Name:        "mem_usage",
					Description: `The actual memory size consumed in the execution of the function, unit is Byte.`,
				},
				resource.Attribute{
					Name:        "request_id",
					Description: `Execute the requestId corresponding to the function.`,
				},
				resource.Attribute{
					Name:        "ret_code",
					Description: `Execution result of function, ` + "`" + `0` + "`" + ` means the execution is successful, other values indicate failure.`,
				},
				resource.Attribute{
					Name:        "ret_msg",
					Description: `Return value after function execution is completed.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Log source.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Point in time at which the function begins execution.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_scf_namespaces",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query SCF namespaces.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the SCF namespace to be queried.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Name of the SCF namespace to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `An information list of namespace. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Name of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the SCF namespace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "namespaces",
					Description: `An information list of namespace. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `Modify time of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Name of the SCF namespace.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the SCF namespace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_security_group",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of security group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the security group to be queried. Conflict with ` + "`" + `security_group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) ID of the security group to be queried. Conflict with ` + "`" + `name` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "be_associate_count",
					Description: `Number of security group binding resources.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the security group.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of the security group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "be_associate_count",
					Description: `Number of security group binding resources.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the security group.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of the security group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_security_groups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of security groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the security group to be queried. Conflict with ` + "`" + `security_group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Project ID of the security group to be queried. Conflict with ` + "`" + `security_group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) ID of the security group to be queried. Conflict with ` + "`" + `name` + "`" + ` and ` + "`" + `project_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the security group to be queried. Conflict with ` + "`" + `security_group_id` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `Information list of security group.`,
				},
				resource.Attribute{
					Name:        "be_associate_count",
					Description: `Number of security group binding resources.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the security group.`,
				},
				resource.Attribute{
					Name:        "egress",
					Description: `Egress rules set. For items like ` + "`" + `[action]#[cidr_ip]#[port]#[protocol]` + "`" + `, it means a regular rule; for items like ` + "`" + `sg-XXXX` + "`" + `, it means a nested security group.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `Ingress rules set. For items like ` + "`" + `[action]#[cidr_ip]#[port]#[protocol]` + "`" + `, it means a regular rule; for items like ` + "`" + `sg-XXXX` + "`" + `, it means a nested security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security group.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `ID of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the security group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "security_groups",
					Description: `Information list of security group.`,
				},
				resource.Attribute{
					Name:        "be_associate_count",
					Description: `Number of security group binding resources.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the security group.`,
				},
				resource.Attribute{
					Name:        "egress",
					Description: `Egress rules set. For items like ` + "`" + `[action]#[cidr_ip]#[port]#[protocol]` + "`" + `, it means a regular rule; for items like ` + "`" + `sg-XXXX` + "`" + `, it means a nested security group.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `Ingress rules set. For items like ` + "`" + `[action]#[cidr_ip]#[port]#[protocol]` + "`" + `, it means a regular rule; for items like ` + "`" + `sg-XXXX` + "`" + `, it means a nested security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security group.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of the security group.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `ID of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the security group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_ssl_certificates",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query SSL certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the SSL certificate to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the SSL certificate to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of the SSL certificate to be queried. Available values includes: ` + "`" + `CA` + "`" + ` and ` + "`" + `SVR` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `An information list of certificate. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "begin_time",
					Description: `Beginning time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `Content of the SSL certificate.`,
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
					Name:        "id",
					Description: `ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "product_zh_name",
					Description: `Certificate authority.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "subject_names",
					Description: `ALL domains included in the SSL certificate. Including the primary domain name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the SSL certificate.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "certificates",
					Description: `An information list of certificate. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "begin_time",
					Description: `Beginning time of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `Content of the SSL certificate.`,
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
					Name:        "id",
					Description: `ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "product_zh_name",
					Description: `Certificate authority.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "subject_names",
					Description: `ALL domains included in the SSL certificate. Including the primary domain name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the SSL certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_subnet",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific VPC subnet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the Subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The VPC ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The AZ for the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the Subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for the Subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The Route Table ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The AZ for the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the Subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for the Subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The Route Table ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_clusters",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query TcaplusDB clusters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) Id of the TcaplusDB cluster to be query.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Optional) Name of the TcaplusDB cluster to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) File for saving results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of TcaplusDB cluster. Each element contains the following attributes.`,
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
					Name:        "cluster_id",
					Description: `Id of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "idl_type",
					Description: `IDL type of the TcaplusDB cluster.`,
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
					Description: `Password status of the TcaplusDB cluster. Valid values: ` + "`" + `unmodifiable` + "`" + `, which means the password can not be changed in this moment; ` + "`" + `modifiable` + "`" + `, which means the password can be changed in this moment.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Access password of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet id of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC id of the TcaplusDB cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of TcaplusDB cluster. Each element contains the following attributes.`,
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
					Name:        "cluster_id",
					Description: `Id of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "idl_type",
					Description: `IDL type of the TcaplusDB cluster.`,
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
					Description: `Password status of the TcaplusDB cluster. Valid values: ` + "`" + `unmodifiable` + "`" + `, which means the password can not be changed in this moment; ` + "`" + `modifiable` + "`" + `, which means the password can be changed in this moment.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Access password of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Subnet id of the TcaplusDB cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC id of the TcaplusDB cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_idls",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query IDL information of the TcaplusDB table.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Id of the TcaplusDB cluster to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) File for saving results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of TcaplusDB table IDL. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "idl_id",
					Description: `Id of the IDL.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of TcaplusDB table IDL. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "idl_id",
					Description: `Id of the IDL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_tablegroups",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query table groups of the TcaplusDB cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Id of the TcaplusDB cluster to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) File for saving results.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `(Optional) Id of the table group to be query.`,
				},
				resource.Attribute{
					Name:        "tablegroup_name",
					Description: `(Optional) Name of the table group to be query. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of table group. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the table group..`,
				},
				resource.Attribute{
					Name:        "table_count",
					Description: `Number of tables.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `Id of the table group.`,
				},
				resource.Attribute{
					Name:        "tablegroup_name",
					Description: `Name of the table group.`,
				},
				resource.Attribute{
					Name:        "total_size",
					Description: `Total storage size (MB).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of table group. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the table group..`,
				},
				resource.Attribute{
					Name:        "table_count",
					Description: `Number of tables.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `Id of the table group.`,
				},
				resource.Attribute{
					Name:        "tablegroup_name",
					Description: `Name of the table group.`,
				},
				resource.Attribute{
					Name:        "total_size",
					Description: `Total storage size (MB).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_tcaplus_tables",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query TcaplusDB tables.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Id of the TcaplusDB cluster to be query.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) File for saving results.`,
				},
				resource.Attribute{
					Name:        "table_id",
					Description: `(Optional) Table id to be query.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `(Optional) Table name to be query.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `(Optional) Id of the table group to be query. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of TcaplusDB tables. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Error message for creating TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "idl_id",
					Description: `IDL file id of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "reserved_read_cu",
					Description: `Reserved read capacity units of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "reserved_volume",
					Description: `Reserved storage capacity of the TcaplusDB table (unit:GB).`,
				},
				resource.Attribute{
					Name:        "reserved_write_cu",
					Description: `Reserved write capacity units of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_id",
					Description: `Id of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_idl_type",
					Description: `IDL type of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `Name of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_size",
					Description: `Size of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_type",
					Description: `Type of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `Table group id of the TcaplusDB table.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `A list of TcaplusDB tables. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Error message for creating TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "idl_id",
					Description: `IDL file id of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "reserved_read_cu",
					Description: `Reserved read capacity units of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "reserved_volume",
					Description: `Reserved storage capacity of the TcaplusDB table (unit:GB).`,
				},
				resource.Attribute{
					Name:        "reserved_write_cu",
					Description: `Reserved write capacity units of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_id",
					Description: `Id of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_idl_type",
					Description: `IDL type of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `Name of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_size",
					Description: `Size of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "table_type",
					Description: `Type of the TcaplusDB table.`,
				},
				resource.Attribute{
					Name:        "tablegroup_id",
					Description: `Table group id of the TcaplusDB table.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpc",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific VPC.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the specific VPC to retrieve. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether or not the default VPC.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `Whether or not the VPC has Multicast support.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether or not the default VPC.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `Whether or not the VPC has Multicast support.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpc_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query vpc instances' information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) Filter VPC with this CIDR.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Filter default or no default VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the VPC to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `(Optional) Filter if VPC has this tag.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the VPC to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `The information list of the VPC.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of a VPC CIDR.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of VPC.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `A list of DNS servers which can be used within the VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default VPC for this region.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `Indicates whether VPC multicast is enabled.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPC.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `A ID list of subnets within this VPC.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `The information list of the VPC.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of a VPC CIDR.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of VPC.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `A list of DNS servers which can be used within the VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default VPC for this region.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `Indicates whether VPC multicast is enabled.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPC.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `A ID list of subnets within this VPC.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpc_route_tables",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query vpc route tables information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "association_main",
					Description: `(Optional) Filter the main routing table.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the routing table to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Optional) ID of the routing table to be queried.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `(Optional) Filter if routing table has this tag.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the routing table to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `The information list of the VPC route table.`,
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
					Name:        "name",
					Description: `Name of the routing table.`,
				},
				resource.Attribute{
					Name:        "route_entry_infos",
					Description: `Detailed information of each entry of the route table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description information user defined for a route table rule.`,
				},
				resource.Attribute{
					Name:        "destination_cidr_block",
					Description: `The destination address block.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `ID of next-hop gateway. Note: when 'next_type' is EIP, GatewayId will fix the value '0'.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `Type of next-hop, and available values include CVM, VPN, DIRECTCONNECT, PEERCONNECTION, SSLVPN, NAT, NORMAL_CVM, EIP and CCN.`,
				},
				resource.Attribute{
					Name:        "route_entry_id",
					Description: `ID of a route table entry.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `ID of the routing table.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `List of subnet IDs bound to the route table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the routing table.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `The information list of the VPC route table.`,
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
					Name:        "name",
					Description: `Name of the routing table.`,
				},
				resource.Attribute{
					Name:        "route_entry_infos",
					Description: `Detailed information of each entry of the route table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description information user defined for a route table rule.`,
				},
				resource.Attribute{
					Name:        "destination_cidr_block",
					Description: `The destination address block.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `ID of next-hop gateway. Note: when 'next_type' is EIP, GatewayId will fix the value '0'.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `Type of next-hop, and available values include CVM, VPN, DIRECTCONNECT, PEERCONNECTION, SSLVPN, NAT, NORMAL_CVM, EIP and CCN.`,
				},
				resource.Attribute{
					Name:        "route_entry_id",
					Description: `ID of a route table entry.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `ID of the routing table.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `List of subnet IDs bound to the route table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the routing table.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpc_subnets",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query vpc subnets information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Zone of the subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) Filter subnet with this CIDR.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Filter default or no default subnets.`,
				},
				resource.Attribute{
					Name:        "is_remote_vpc_snat",
					Description: `(Optional) Filter the VPC SNAT address pool subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of the subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `(Optional) Filter if subnet has this tag.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `List of subnets.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone of the subnet.`,
				},
				resource.Attribute{
					Name:        "available_ip_count",
					Description: `The number of available IPs.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of the subnet.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the subnet resource.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default subnet of the VPC for this region.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `Indicates whether multicast is enabled.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `ID of the routing table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the subnet resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `List of subnets.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone of the subnet.`,
				},
				resource.Attribute{
					Name:        "available_ip_count",
					Description: `The number of available IPs.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of the subnet.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the subnet resource.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether it is the default subnet of the VPC for this region.`,
				},
				resource.Attribute{
					Name:        "is_multicast",
					Description: `Indicates whether multicast is enabled.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `ID of the routing table.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the subnet resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_connections",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of VPN connections.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `(Optional) Customer gateway ID of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the VPN connection. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the VPN connection to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `(Optional) VPN gateway ID of the VPN connection. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "connection_list",
					Description: `Information list of the dedicated connections.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `ID of the customer gateway.`,
				},
				resource.Attribute{
					Name:        "encrypt_proto",
					Description: `Encrypt proto of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "ike_dh_group_name",
					Description: `DH group name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_exchange_mode",
					Description: `Exchange mode of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_local_address",
					Description: `Local address of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_local_fqdn_name",
					Description: `Local FQDN name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_local_identity",
					Description: `Local identity of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_proto_authen_algorithm",
					Description: `Proto authenticate algorithm of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_proto_encry_algorithm",
					Description: `Proto encrypt algorithm of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_remote_address",
					Description: `Remote address of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_remote_fqdn_name",
					Description: `Remote FQDN name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_remote_identity",
					Description: `Remote identity of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_sa_lifetime_seconds",
					Description: `SA lifetime of the IKE operation specification, unit is ` + "`" + `second` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `Version of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_encrypt_algorithm",
					Description: `Encrypt algorithm of the IPSEC operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_integrity_algorithm",
					Description: `Integrity algorithm of the IPSEC operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_pfs_dh_group",
					Description: `PFS DH group name of the IPSEC operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_sa_lifetime_seconds",
					Description: `SA lifetime of the IPSEC operation specification, unit is ` + "`" + `second` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_sa_lifetime_traffic",
					Description: `SA lifetime traffic of the IPSEC operation specification, unit is ` + "`" + `KB` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "net_status",
					Description: `Net status of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "pre_share_key",
					Description: `Pre-shared key of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `Route type of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "security_group_policy",
					Description: `Security group policy of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "local_cidr_block",
					Description: `Local cidr block.`,
				},
				resource.Attribute{
					Name:        "remote_cidr_block",
					Description: `Remote cidr block list.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_proto",
					Description: `Vpn proto of the VPN connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_list",
					Description: `Information list of the dedicated connections.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `ID of the customer gateway.`,
				},
				resource.Attribute{
					Name:        "encrypt_proto",
					Description: `Encrypt proto of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "ike_dh_group_name",
					Description: `DH group name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_exchange_mode",
					Description: `Exchange mode of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_local_address",
					Description: `Local address of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_local_fqdn_name",
					Description: `Local FQDN name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_local_identity",
					Description: `Local identity of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_proto_authen_algorithm",
					Description: `Proto authenticate algorithm of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_proto_encry_algorithm",
					Description: `Proto encrypt algorithm of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_remote_address",
					Description: `Remote address of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_remote_fqdn_name",
					Description: `Remote FQDN name of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_remote_identity",
					Description: `Remote identity of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ike_sa_lifetime_seconds",
					Description: `SA lifetime of the IKE operation specification, unit is ` + "`" + `second` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `Version of the IKE operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_encrypt_algorithm",
					Description: `Encrypt algorithm of the IPSEC operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_integrity_algorithm",
					Description: `Integrity algorithm of the IPSEC operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_pfs_dh_group",
					Description: `PFS DH group name of the IPSEC operation specification.`,
				},
				resource.Attribute{
					Name:        "ipsec_sa_lifetime_seconds",
					Description: `SA lifetime of the IPSEC operation specification, unit is ` + "`" + `second` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipsec_sa_lifetime_traffic",
					Description: `SA lifetime traffic of the IPSEC operation specification, unit is ` + "`" + `KB` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "net_status",
					Description: `Net status of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "pre_share_key",
					Description: `Pre-shared key of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `Route type of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "security_group_policy",
					Description: `Security group policy of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "local_cidr_block",
					Description: `Local cidr block.`,
				},
				resource.Attribute{
					Name:        "remote_cidr_block",
					Description: `Remote cidr block list.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_proto",
					Description: `Vpn proto of the VPN connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_customer_gateways",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of VPN customer gateways.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the customer gateway. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `(Optional) Public ip address of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the VPN customer gateway to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "gateway_list",
					Description: `Information list of the dedicated gateways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `Public ip address of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the VPN customer gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gateway_list",
					Description: `Information list of the dedicated gateways.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `Public ip address of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the VPN customer gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_vpn_gateways",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query detailed information of VPN gateways.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the VPN gateway. The length of character is limited to 1-60.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `(Optional) Public ip address of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the VPN gateway to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone of the VPN gateway. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "gateway_list",
					Description: `Information list of the dedicated gateways.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of VPN gateway (unit: Mbps), the available values include: 5,10,20,50,100. Default is 5.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Charge Type of the VPN gateway, valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` and default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
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
					Name:        "id",
					Description: `ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "is_address_blocked",
					Description: `Indicates whether ip address is blocked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "new_purchase_plan",
					Description: `The plan of new purchase, valid value is ` + "`" + `PREPAID_TO_POSTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `Flag indicates whether to renew or not, valid values are ` + "`" + `NOTIFY_AND_RENEW` + "`" + `, ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, ` + "`" + `NOT_NOTIFY_AND_NOT_RENEW` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `Public ip of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "restrict_state",
					Description: `Restrict state of VPN gateway, valid values are ` + "`" + `PRETECIVELY_ISOLATED` + "`" + `, ` + "`" + `NORMAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the VPN gateway, valid values are ` + "`" + `PENDING` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of gateway instance, valid values are ` + "`" + `IPSEC` + "`" + `, ` + "`" + `SSL` + "`" + ` and ` + "`" + `CCN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Zone of the VPN gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gateway_list",
					Description: `Information list of the dedicated gateways.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of VPN gateway (unit: Mbps), the available values include: 5,10,20,50,100. Default is 5.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Charge Type of the VPN gateway, valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID_BY_HOUR` + "`" + ` and default is ` + "`" + `POSTPAID_BY_HOUR` + "`" + `.`,
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
					Name:        "id",
					Description: `ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "is_address_blocked",
					Description: `Indicates whether ip address is blocked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "new_purchase_plan",
					Description: `The plan of new purchase, valid value is ` + "`" + `PREPAID_TO_POSTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "prepaid_renew_flag",
					Description: `Flag indicates whether to renew or not, valid values are ` + "`" + `NOTIFY_AND_RENEW` + "`" + `, ` + "`" + `NOTIFY_AND_AUTO_RENEW` + "`" + `, ` + "`" + `NOT_NOTIFY_AND_NOT_RENEW` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `Public ip of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "restrict_state",
					Description: `Restrict state of VPN gateway, valid values are ` + "`" + `PRETECIVELY_ISOLATED` + "`" + `, ` + "`" + `NORMAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the VPN gateway, valid values are ` + "`" + `PENDING` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags used to associate different resources.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of gateway instance, valid values are ` + "`" + `IPSEC` + "`" + `, ` + "`" + `SSL` + "`" + ` and ` + "`" + `CCN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Zone of the VPN gateway.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"tencentcloud_as_scaling_configs":           0,
		"tencentcloud_as_scaling_groups":            1,
		"tencentcloud_as_scaling_policies":          2,
		"tencentcloud_availability_regions":         3,
		"tencentcloud_availability_zones":           4,
		"tencentcloud_cam_group_memberships":        5,
		"tencentcloud_cam_group_policy_attachments": 6,
		"tencentcloud_cam_groups":                   7,
		"tencentcloud_cam_policies":                 8,
		"tencentcloud_cam_role_policy_attachments":  9,
		"tencentcloud_cam_roles":                    10,
		"tencentcloud_cam_saml_providers":           11,
		"tencentcloud_cam_user_policy_attachments":  12,
		"tencentcloud_cam_users":                    13,
		"tencentcloud_cbs_snapshot_policies":        14,
		"tencentcloud_cbs_snapshots":                15,
		"tencentcloud_cbs_storages":                 16,
		"tencentcloud_ccn_bandwidth_limits":         17,
		"tencentcloud_ccn_instances":                18,
		"tencentcloud_cfs_access_groups":            19,
		"tencentcloud_cfs_access_rules":             20,
		"tencentcloud_cfs_file_systems":             21,
		"tencentcloud_clb_attachments":              22,
		"tencentcloud_clb_instances":                23,
		"tencentcloud_clb_listener_rules":           24,
		"tencentcloud_clb_listeners":                25,
		"tencentcloud_clb_redirections":             26,
		"tencentcloud_container_cluster_instances":  27,
		"tencentcloud_container_clusters":           28,
		"tencentcloud_cos_bucket_object":            29,
		"tencentcloud_cos_buckets":                  30,
		"tencentcloud_dayu_cc_http_policies":        31,
		"tencentcloud_dayu_cc_https_policies":       32,
		"tencentcloud_dayu_ddos_policies":           33,
		"tencentcloud_dayu_ddos_policy_attachments": 34,
		"tencentcloud_dayu_ddos_policy_cases":       35,
		"tencentcloud_dayu_l4_rules":                36,
		"tencentcloud_dayu_l7_rules":                37,
		"tencentcloud_dc_gateway_ccn_routes":        38,
		"tencentcloud_dc_gateway_instances":         39,
		"tencentcloud_dc_instances":                 40,
		"tencentcloud_dcx_instances":                41,
		"tencentcloud_dnats":                        42,
		"tencentcloud_eip":                          43,
		"tencentcloud_eips":                         44,
		"tencentcloud_elasticsearch_instances":      45,
		"tencentcloud_enis":                         46,
		"tencentcloud_gaap_certificates":            47,
		"tencentcloud_gaap_domain_error_pages":      48,
		"tencentcloud_gaap_http_domains":            49,
		"tencentcloud_gaap_http_rules":              50,
		"tencentcloud_gaap_layer4_listeners":        51,
		"tencentcloud_gaap_layer7_listeners":        52,
		"tencentcloud_gaap_proxies":                 53,
		"tencentcloud_gaap_realservers":             54,
		"tencentcloud_gaap_security_policies":       55,
		"tencentcloud_gaap_security_rules":          56,
		"tencentcloud_ha_vip_eip_attachments":       57,
		"tencentcloud_ha_vips":                      58,
		"tencentcloud_image":                        59,
		"tencentcloud_images":                       60,
		"tencentcloud_instance_types":               61,
		"tencentcloud_instances":                    62,
		"tencentcloud_key_pairs":                    63,
		"tencentcloud_kubernetes_clusters":          64,
		"tencentcloud_mongodb_instances":            65,
		"tencentcloud_mongodb_zone_config":          66,
		"tencentcloud_monitor_binding_objects":      67,
		"tencentcloud_monitor_data":                 68,
		"tencentcloud_monitor_policy_conditions":    69,
		"tencentcloud_monitor_policy_groups":        70,
		"tencentcloud_monitor_product_event":        71,
		"tencentcloud_monitor_product_namespace":    72,
		"tencentcloud_mysql_backup_list":            73,
		"tencentcloud_mysql_instance":               74,
		"tencentcloud_mysql_parameter_list":         75,
		"tencentcloud_mysql_zone_config":            76,
		"tencentcloud_nat_gateways":                 77,
		"tencentcloud_nats":                         78,
		"tencentcloud_placement_groups":             79,
		"tencentcloud_redis_instances":              80,
		"tencentcloud_redis_zone_config":            81,
		"tencentcloud_reserved_instance_configs":    82,
		"tencentcloud_reserved_instances":           83,
		"tencentcloud_route_table":                  84,
		"tencentcloud_scf_functions":                85,
		"tencentcloud_scf_logs":                     86,
		"tencentcloud_scf_namespaces":               87,
		"tencentcloud_security_group":               88,
		"tencentcloud_security_groups":              89,
		"tencentcloud_ssl_certificates":             90,
		"tencentcloud_subnet":                       91,
		"tencentcloud_tcaplus_clusters":             92,
		"tencentcloud_tcaplus_idls":                 93,
		"tencentcloud_tcaplus_tablegroups":          94,
		"tencentcloud_tcaplus_tables":               95,
		"tencentcloud_vpc":                          96,
		"tencentcloud_vpc_instances":                97,
		"tencentcloud_vpc_route_tables":             98,
		"tencentcloud_vpc_subnets":                  99,
		"tencentcloud_vpn_connections":              100,
		"tencentcloud_vpn_customer_gateways":        101,
		"tencentcloud_vpn_gateways":                 102,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
