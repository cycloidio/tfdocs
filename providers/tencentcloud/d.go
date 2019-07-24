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
					Description: `ID list of login keys`,
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
					Description: `ID list of login keys`,
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
					Description: `(Optional) A scaling group name used to query. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `A lsit of application clb ids.`,
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
					Description: `A lsit of traditional clb ids which the CVM instances attached to.`,
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
					Description: `A lsit of application clb ids.`,
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
					Description: `A lsit of traditional clb ids which the CVM instances attached to.`,
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
			Type:             "tencentcloud_availability_zones",
			Category:         "Data Sources",
			ShortDescription: `Get available zones in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "include_unavailable",
					Description: `(Optional) A bool variable Indicates that the query will include ` + "`" + `UNAVAILABLE` + "`" + ` zones.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) When specified, only the zone with the exactly name match will return. ## Attributes Reference A list of zones will be exported and its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An internal id for the zone, like ` + "`" + `200003` + "`" + `, usually not so useful for end user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The english name for the zone, like ` + "`" + `ap-guangzhou-3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the zone, unfortunately only Chinese characters at this stage.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state for the zone, indicate availability using ` + "`" + `AVAILABLE` + "`" + ` and ` + "`" + `UNAVAILABLE` + "`" + ` values.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An internal id for the zone, like ` + "`" + `200003` + "`" + `, usually not so useful for end user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The english name for the zone, like ` + "`" + `ap-guangzhou-3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the zone, unfortunately only Chinese characters at this stage.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state for the zone, indicate availability using ` + "`" + `AVAILABLE` + "`" + ` and ` + "`" + `UNAVAILABLE` + "`" + ` values.`,
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
					Description: `(Required, ForceNew) ID of the CCN to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, ForceNew) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "limits",
					Description: `The bandwidth limits of regions`,
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
					Description: `The bandwidth limits of regions`,
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
					Description: `(Optional, ForceNew) ID of the CCN to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) Name of the CCN to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, ForceNew) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `Type of attached instance network, and available values include VPC, DIRECTCONNECT and BMVPC.`,
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
					Description: `Type of attached instance network, and available values include VPC, DIRECTCONNECT and BMVPC.`,
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
					Description: `(Optional) An int variable describe how many instances in return at most. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Describe how many nodes in the cluster. A list of nodes will be exported and its every element contains the following attributes:`,
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
					Name:        "mem",
					Description: `Describe the memory of the node.`,
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
					Name:        "total_count",
					Description: `Describe how many nodes in the cluster. A list of nodes will be exported and its every element contains the following attributes:`,
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
					Name:        "mem",
					Description: `Describe the memory of the node.`,
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
					Description: `(Optional) An int variable describe how many cluster in return at most . ## Attributes Reference A list of clusters will be exported and its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `An id identify the cluster, like ` + "`" + `cls-xxxxxx` + "`" + `.`,
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
					Name:        "security_username",
					Description: `Describe the username needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "security_password",
					Description: `Describe the password needed for using kubectl to access to kubernetes.`,
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
					Name:        "total_cpu",
					Description: `Describe the total cpu of each instance in the cluster.`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `Describe the total memory of each instance in the cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `An id identify the cluster, like ` + "`" + `cls-xxxxxx` + "`" + `.`,
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
					Name:        "security_username",
					Description: `Describe the username needed for using kubectl to access to kubernetes.`,
				},
				resource.Attribute{
					Name:        "security_password",
					Description: `Describe the password needed for using kubectl to access to kubernetes.`,
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
					Name:        "total_cpu",
					Description: `Describe the total cpu of each instance in the cluster.`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `Describe the total memory of each instance in the cluster.`,
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
					Description: `ETag generated for the object，which is may not equal to MD5 value.`,
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
					Description: `ETag generated for the object，which is may not equal to MD5 value.`,
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
					Description: `(Optional) A prefix string to filter results by bucket name`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `ID of the DCG`,
				},
				resource.Attribute{
					Name:        "dcg_ip",
					Description: `IP of the DCG`,
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
					Description: `Name of the DCG`,
				},
				resource.Attribute{
					Name:        "network_instance_id",
					Description: `Type of associated network, the available value include 'VPC' and 'CCN'.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `IP of the DCG`,
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
					Description: `ID of the DCG`,
				},
				resource.Attribute{
					Name:        "dcg_ip",
					Description: `IP of the DCG`,
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
					Description: `Name of the DCG`,
				},
				resource.Attribute{
					Name:        "network_instance_id",
					Description: `Type of associated network, the available value include 'VPC' and 'CCN'.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `IP of the DCG`,
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
					Description: `(Optional, ForceNew) ID of the DC to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) Name of the DC to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, ForceNew) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `(Optional, ForceNew) ID of the dedicated tunnels to be queried.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) Name of the dedicated tunnels to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, ForceNew) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_eip",
			Category:         "Data Sources",
			ShortDescription: `Provides an available EIP for the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to filter off of. There are several valid keys: ` + "`" + `address-id` + "`" + `,` + "`" + `address-name` + "`" + `,` + "`" + `address-ip` + "`" + `. For a full reference, check out [DescribeImages in the TencentCloud API reference](https://intl.cloud.tencent.com/document/api/213/9451#filter). ## Attributes Reference`,
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
					Description: `The status of the EIP, there are several status like ` + "`" + `BIND` + "`" + `, ` + "`" + `UNBIND` + "`" + `, and ` + "`" + `BIND_ENI` + "`" + `. For a full reference, check out [DescribeImages in the TencentCloud API reference](https://intl.cloud.tencent.com/document/api/213/9452#eip_state).`,
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
					Description: `The status of the EIP, there are several status like ` + "`" + `BIND` + "`" + `, ` + "`" + `UNBIND` + "`" + `, and ` + "`" + `BIND_ENI` + "`" + `. For a full reference, check out [DescribeImages in the TencentCloud API reference](https://intl.cloud.tencent.com/document/api/213/9452#eip_state).`,
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
					Name:        "image_name_regex",
					Description: `(Optional) A regex string to apply to the image list returned by TencentCloud.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `(Optional) A string to apply with fuzzy match to the os_name atrribute on the image list returned by TencentCloud.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to filter off of. There are several valid keys: ` + "`" + `image-id` + "`" + `,` + "`" + `image-type` + "`" + `,` + "`" + `image-name` + "`" + `. For a full reference, check out [DescribeImages in the TencentCloud API reference](https://intl.cloud.tencent.com/document/api/213/9451#filter). ## Attributes Reference`,
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
			Type:             "tencentcloud_instance_types",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of instance types available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) One or more name/value pairs to filter off of. There are several valid keys: ` + "`" + `zone` + "`" + `,` + "`" + `instance-family` + "`" + `. For a full reference, check out [DescribeInstanceTypeConfigs in the TencentCloud API reference](https://intl.cloud.tencent.com/document/api/213/9391).`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `(Optional) Limit search to specific cpu core count.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `(Optional) Limit search to specific memory size. ## Attributes Reference The following attributes are exported`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Indicate the availability zone for this instance type.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `TencentCloud instance type of the cvm instance.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `Number of CPU cores.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Size of memory, measured in GB.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The instance type family.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Indicate the availability zone for this instance type.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `TencentCloud instance type of the cvm instance.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `Number of CPU cores.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Size of memory, measured in GB.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The instance type family.`,
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
					Description: `(Required, ForceNew) Instance ID, such as cdb-c1nl9rpv. It is identical to the instance ID displayed in the database console page.`,
				},
				resource.Attribute{
					Name:        "max_number",
					Description: `(Optional, ForceNew) The latest files to list, rang from 1 to 10000. And the default value is 10.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, ForceNew) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "engine_version",
					Description: `(Optional) The version number of the database engine to use. Supported versions include 5.5/5.6/5.7.`,
				},
				resource.Attribute{
					Name:        "init_flag",
					Description: `(Optional) Initialization mark. Available values: 0 - Uninitialized; 1 – Initialized.`,
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
					Name:        "result_output_file",
					Description: `(Optional) Used to store results.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) Security groups ID of instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Instance status. Available values: 0 - Creating; 1 - Running; 4 - Isolating; 5 – Isolated.`,
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
					Name:        "cpu_core_count",
					Description: `CPU count.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time at which a instance is created.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Supported instance model.HA - high available version; Basic - basic version.`,
				},
				resource.Attribute{
					Name:        "dr_instance_ids",
					Description: `ID list of disaster-recovory type associated with the current instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `The version number of the database engine to use. Supported versions include 5.5/5.6/5.7.`,
				},
				resource.Attribute{
					Name:        "init_flag",
					Description: `Initialization mark. Available values: 0 - Uninitialized; 1 – Initialized.`,
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
					Name:        "memory_size",
					Description: `Memory size (in MB).`,
				},
				resource.Attribute{
					Name:        "mysql_id",
					Description: `Instance ID, such as cdb-c1nl9rpv. It is identical to the instance ID displayed in the database console page.`,
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
					Description: `Instance status. Available values: 0 - Creating; 1 - Running; 4 - Isolating; 5 – Isolated.`,
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
					Name:        "cpu_core_count",
					Description: `CPU count.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time at which a instance is created.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Supported instance model.HA - high available version; Basic - basic version.`,
				},
				resource.Attribute{
					Name:        "dr_instance_ids",
					Description: `ID list of disaster-recovory type associated with the current instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `The version number of the database engine to use. Supported versions include 5.5/5.6/5.7.`,
				},
				resource.Attribute{
					Name:        "init_flag",
					Description: `Initialization mark. Available values: 0 - Uninitialized; 1 – Initialized.`,
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
					Name:        "memory_size",
					Description: `Memory size (in MB).`,
				},
				resource.Attribute{
					Name:        "mysql_id",
					Description: `Instance ID, such as cdb-c1nl9rpv. It is identical to the instance ID displayed in the database console page.`,
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
					Description: `Instance status. Available values: 0 - Creating; 1 - Running; 4 - Isolating; 5 – Isolated.`,
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
					Description: `(Optional, ForceNew) Region parameter, which is used to identify the region to which the data you want to work with belongs.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, ForceNew) Used to store results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "tencentcloud_nats",
			Category:         "Data Sources",
			ShortDescription: `The NATs data source lists a number of NATs resource information owned by an TencentCloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The VPC ID for NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "max_concurrent",
					Description: `(Optional) The upper limit of concurrent connection of NAT gateway, for example: 1000000, 3000000, 10000000. To learn more, please refer to [Virtual Private Cloud Gateway Description](https://intl.cloud.tencent.com/doc/product/215/1682).`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The maximum public network output bandwidth of the gateway (unit: Mbps), for example: 10, 20, 50, 100, 200, 500, 1000, 2000, 5000. For more information, please refer to [Virtual Private Cloud Gateway Description](https://intl.cloud.tencent.com/doc/product/215/1682).`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `(Optional) Elastic IP arrays bound to the gateway, For more information on elastic IP, please refer to [Elastic IP](eip.html).`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) NAT gateway status, 0: Running, 1: Unavailable, 2: Be in arrears and out of service ## Attributes Reference The following attributes are exported:`,
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
					Description: `The upper limit of concurrent connection of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of the NAT gateway (unit: Mbps).`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `Elastic IP arrays bound to the NAT gateway`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `NAT gateway status, 0: Running, 1: Unavailable, 2: Be in arrears and out of service`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the NAT gateway`,
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
					Description: `The upper limit of concurrent connection of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The maximum public network output bandwidth of the NAT gateway (unit: Mbps).`,
				},
				resource.Attribute{
					Name:        "assigned_eip_set",
					Description: `Elastic IP arrays bound to the NAT gateway`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `NAT gateway status, 0: Running, 1: Unavailable, 2: Be in arrears and out of service`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the NAT gateway`,
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
					Description: `(Optional, ForceNew) The number limitation of results for a query.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional, ForceNew) ID of the project to which redis instance belongs.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, ForceNew) Used to save results.`,
				},
				resource.Attribute{
					Name:        "search_key",
					Description: `(Optional, ForceNew) Key words used to match the results, and the key words can be: instance ID, instance name and IP address.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional, ForceNew) ID of an available zone. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `A list of redis instance. Each element contains the following attributes:`,
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
					Description: `Memory size in MB`,
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
					Name:        "status",
					Description: `Current status of an instance，maybe: init, processing, online, isolate and todelete.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the vpc subnet.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Instance type. Available values: master_slave_redis, master_slave_ckv, cluster_ckv, cluster_redis and standalone_redis.`,
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
					Name:        "create_time",
					Description: `The time when the instance is created.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of an instance.`,
				},
				resource.Attribute{
					Name:        "mem_size",
					Description: `Memory size in MB`,
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
					Name:        "status",
					Description: `Current status of an instance，maybe: init, processing, online, isolate and todelete.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the vpc subnet.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Instance type. Available values: master_slave_redis, master_slave_ckv, cluster_ckv, cluster_redis and standalone_redis.`,
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
					Description: `(Optional, ForceNew) Name of a region. If this value is not set, the current region getting from provider's configuration will be used.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, ForceNew) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `A list of zone. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "mem_sizes",
					Description: `The memory volume of an available instance（in MB）.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Instance type. Available values: master_slave_redis, master_slave_ckv, cluster_ckv, cluster_redis and standalone_redis.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version description of an available instance. Possible values: Redis 3.2, Redis 4.0 .`,
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
					Description: `The memory volume of an available instance（in MB）.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Instance type. Available values: master_slave_redis, master_slave_ckv, cluster_ckv, cluster_redis and standalone_redis.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version description of an available instance. Possible values: Redis 3.2, Redis 4.0 .`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `ID of available zone.`,
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
					Description: `(Required) The Route Table ID. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for Route Table.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `routes are also exported with the following attributes, when there are relevants: Each route supports the following:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The RouteEntry's target network segment.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `The ` + "`" + `next_hub` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `The RouteEntry's next hub.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The RouteEntry's description.`,
				},
				resource.Attribute{
					Name:        "subnet_num",
					Description: `Number of associated subnets.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of routing table, for example: 2018-01-22 17:50:21.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name for Route Table.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `routes are also exported with the following attributes, when there are relevants: Each route supports the following:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The RouteEntry's target network segment.`,
				},
				resource.Attribute{
					Name:        "next_type",
					Description: `The ` + "`" + `next_hub` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "next_hub",
					Description: `The RouteEntry's next hub.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The RouteEntry's description.`,
				},
				resource.Attribute{
					Name:        "subnet_num",
					Description: `Number of associated subnets.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of routing table, for example: 2018-01-22 17:50:21.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tencentcloud_security_group",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Security Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The ID of the security group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "be_associate_count",
					Description: `Number of associated instances.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of security group, for example: 2018-01-22 17:50:21.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "be_associate_count",
					Description: `Number of associated instances.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of security group, for example: 2018-01-22 17:50:21.`,
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
					Name:        "vpc_id",
					Description: `(Required) The VPC ID.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the Subnet. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for the Subnet.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the Subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The Route Table ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name for the Subnet.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the Subnet.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The Route Table ID.`,
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
					Description: `(Optional) VPC name. Fuzzy search is supported, as defined by [the underlying TencentCloud API](https://intl.cloud.tencent.com/document/product/215/1372). ## Attributes Reference All of the argument attributes except ` + "`" + `filter` + "`" + ` blocks are also exported as result attributes. This data source will complete the data by populating any fields that are not included in the configuration with the data for the selected VPC. The following attribute is additionally exported:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the VPC.`,
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
					Name:        "name",
					Description: `(Optional, ForceNew) Name of the VPC to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, ForceNew) Used to save results.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) ID of the VPC to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "name",
					Description: `(Optional, ForceNew) Name of the routing table to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, ForceNew) Used to save results.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Optional, ForceNew) ID of the routing table to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `The information list of the VPC.`,
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
					Name:        "name",
					Description: `(Optional, ForceNew) Name of the subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, ForceNew) Used to save results.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, ForceNew) ID of the subnet to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"tencentcloud_as_scaling_configs":          0,
		"tencentcloud_as_scaling_groups":           1,
		"tencentcloud_as_scaling_policies":         2,
		"tencentcloud_availability_zones":          3,
		"tencentcloud_cbs_snapshots":               4,
		"tencentcloud_cbs_storages":                5,
		"tencentcloud_ccn_bandwidth_limits":        6,
		"tencentcloud_ccn_instances":               7,
		"tencentcloud_container_cluster_instances": 8,
		"tencentcloud_container_clusters":          9,
		"tencentcloud_cos_bucket_object":           10,
		"tencentcloud_cos_buckets":                 11,
		"tencentcloud_dc_gateway_ccn_routes":       12,
		"tencentcloud_dc_gateway_instances":        13,
		"tencentcloud_dc_instances":                14,
		"tencentcloud_dcx_instances":               15,
		"tencentcloud_eip":                         16,
		"tencentcloud_image":                       17,
		"tencentcloud_instance_types":              18,
		"tencentcloud_mysql_backup_list":           19,
		"tencentcloud_mysql_instance":              20,
		"tencentcloud_mysql_parameter_list":        21,
		"tencentcloud_mysql_zone_config":           22,
		"tencentcloud_nats":                        23,
		"tencentcloud_redis_instances":             24,
		"tencentcloud_redis_zone_config":           25,
		"tencentcloud_route_table":                 26,
		"tencentcloud_security_group":              27,
		"tencentcloud_subnet":                      28,
		"tencentcloud_vpc":                         29,
		"tencentcloud_vpc_instances":               30,
		"tencentcloud_vpc_route_tables":            31,
		"tencentcloud_vpc_subnets":                 32,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
