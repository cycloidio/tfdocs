package ncloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ncloud_access_control_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

When creating a server instance (VM), you can add an ACG(Access Control Group) that you specified to set firewalls. ` + "`" + `ncloud_access_control_group` + "`" + ` provides details about a specific ACG(Access Control Group) information.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) List of ACG ID you want to get`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the ACG you want to get`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Optional) The ID of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Indicates whether to get default group only`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "access_control_group_no",
					Description: `The ID of ACG. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `ACG description`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_control_group_no",
					Description: `The ID of ACG. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `ACG description`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_access_control_groups",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

When creating a server instance (VM), you can add an ACG(Access Control Group) that you specified to set firewalls. This data source gets a list of access control groups necessary to set firewalls.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "configuration_no_list",
					Description: `(Optional) List of ACG configuration numbers you want to get`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Indicates whether to get default groups only`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the ACG you want to get`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save data source after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "access_control_groups",
					Description: `A List of access control group configuration_no.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_control_groups",
					Description: `A List of access control group configuration_no.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_access_control_rule",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Access configuration rule you want to get.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_control_group_configuration_no",
					Description: `(Optional) Access control group number to search`,
				},
				resource.Attribute{
					Name:        "access_control_group_name",
					Description: `(Optional) Access control group name to search`,
				},
				resource.Attribute{
					Name:        "is_default_group",
					Description: `(Optional) Whether default group`,
				},
				resource.Attribute{
					Name:        "source_name_regex",
					Description: `(Optional) A regex string to apply to the source access control rule list returned by ncloud ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IP`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `Destination Port`,
				},
				resource.Attribute{
					Name:        "protocol_type_code",
					Description: `Protocol type code`,
				},
				resource.Attribute{
					Name:        "configuration_no",
					Description: `Access control rule configuration no`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `Protocol type code`,
				},
				resource.Attribute{
					Name:        "source_configuration_no",
					Description: `Source access control rule configuration no`,
				},
				resource.Attribute{
					Name:        "source_name",
					Description: `Source access control rule name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Access control rule description`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "source_ip",
					Description: `Source IP`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `Destination Port`,
				},
				resource.Attribute{
					Name:        "protocol_type_code",
					Description: `Protocol type code`,
				},
				resource.Attribute{
					Name:        "configuration_no",
					Description: `Access control rule configuration no`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `Protocol type code`,
				},
				resource.Attribute{
					Name:        "source_configuration_no",
					Description: `Source access control rule configuration no`,
				},
				resource.Attribute{
					Name:        "source_name",
					Description: `Source access control rule name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Access control rule description`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_access_control_rules",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

List of access configuration rules you want to get

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_control_group_configuration_no",
					Description: `(Required) Access control group configuration number to search`,
				},
				resource.Attribute{
					Name:        "source_name_regex",
					Description: `(Optional) A regex string to apply to the ACG rule list returned by ncloud`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save data source after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "access_control_rules",
					Description: `A list of access control rules configuration no`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_control_rules",
					Description: `A list of access control rules configuration no`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_auto_scaling_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Auto Scaling Group created before.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific auto scaling group to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group_no",
					Description: `The ID of Auto Scaling Group. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "launch_configuration_no",
					Description: `The number of the associated launch configuration.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `The number of servers is adjusted according to the desired capacity value.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `The minimum size of the Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum size of the Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `The cooldown time is the period set to ignore even if the monitoring event alarm occurs after the actual scaling is being performed or is completed.`,
				},
				resource.Attribute{
					Name:        "health_check_type_code",
					Description: `` + "`" + `SVR` + "`" + ` or ` + "`" + `LOADB` + "`" + `. Controls how health checking is done.`,
				},
				resource.Attribute{
					Name:        "wait_for_capacity_timeout",
					Description: `The maximum amount of time Terraform should wait for an ASG instance to become healthy.`,
				},
				resource.Attribute{
					Name:        "health_check_grace_period",
					Description: `time to hold health check after the server instance is put into the service with the health check hold period.`,
				},
				resource.Attribute{
					Name:        "server_instance_no_list",
					Description: `List of server instances belonging to Auto Scaling Group. ~>`,
				},
				resource.Attribute{
					Name:        "zone_no_list",
					Description: `the list of zone numbers where server instances belonging to this group will exist. ~>`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `The ID of the associated Subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "access_control_group_no_list",
					Description: `The ID of the ACG.`,
				},
				resource.Attribute{
					Name:        "target_group_list",
					Description: `Target Group number list of Load Balancer.`,
				},
				resource.Attribute{
					Name:        "server_name_prefix",
					Description: `Create name beginning with the specified prefix.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "auto_scaling_group_no",
					Description: `The ID of Auto Scaling Group. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "launch_configuration_no",
					Description: `The number of the associated launch configuration.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `The number of servers is adjusted according to the desired capacity value.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `The minimum size of the Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum size of the Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `The cooldown time is the period set to ignore even if the monitoring event alarm occurs after the actual scaling is being performed or is completed.`,
				},
				resource.Attribute{
					Name:        "health_check_type_code",
					Description: `` + "`" + `SVR` + "`" + ` or ` + "`" + `LOADB` + "`" + `. Controls how health checking is done.`,
				},
				resource.Attribute{
					Name:        "wait_for_capacity_timeout",
					Description: `The maximum amount of time Terraform should wait for an ASG instance to become healthy.`,
				},
				resource.Attribute{
					Name:        "health_check_grace_period",
					Description: `time to hold health check after the server instance is put into the service with the health check hold period.`,
				},
				resource.Attribute{
					Name:        "server_instance_no_list",
					Description: `List of server instances belonging to Auto Scaling Group. ~>`,
				},
				resource.Attribute{
					Name:        "zone_no_list",
					Description: `the list of zone numbers where server instances belonging to this group will exist. ~>`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `The ID of the associated Subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "access_control_group_no_list",
					Description: `The ID of the ACG.`,
				},
				resource.Attribute{
					Name:        "target_group_list",
					Description: `Target Group number list of Load Balancer.`,
				},
				resource.Attribute{
					Name:        "server_name_prefix",
					Description: `Create name beginning with the specified prefix.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_auto_scaling_policy",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Auto Scaling Policy created before.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific auto scaling policy to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Auto Scaling Policy.`,
				},
				resource.Attribute{
					Name:        "adjustment_type_code",
					Description: `how the number of servers is scaled when the scaling policy is performed.`,
				},
				resource.Attribute{
					Name:        "scaling_adjustment",
					Description: `Specify the adjustment value for the adjustment type.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `The cooldown time is the period set to ignore even if the monitoring event alarm occurs after the actual scaling is being performed or is completed.`,
				},
				resource.Attribute{
					Name:        "min_adjustment_step",
					Description: `Change the number of server instances by the minimum adjustment width.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of Auto Scaling Policy.`,
				},
				resource.Attribute{
					Name:        "adjustment_type_code",
					Description: `how the number of servers is scaled when the scaling policy is performed.`,
				},
				resource.Attribute{
					Name:        "scaling_adjustment",
					Description: `Specify the adjustment value for the adjustment type.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `The cooldown time is the period set to ignore even if the monitoring event alarm occurs after the actual scaling is being performed or is completed.`,
				},
				resource.Attribute{
					Name:        "min_adjustment_step",
					Description: `Change the number of server instances by the minimum adjustment width.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_auto_scaling_schedule",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Auto Scaling Schedule created before.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific auto scaling schedule to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Auto Scaling Schedule.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `The number of servers is adjusted according to the desired capacity value.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `The minimum size of the Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum size of the Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `the date and time when the schedule first starts.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `the date and time when the schedule end.`,
				},
				resource.Attribute{
					Name:        "recurrence",
					Description: `Repeat Settings.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group_no",
					Description: `The number of the auto scaling group. ~>`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `the time band for the repeat settings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of Auto Scaling Schedule.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `The number of servers is adjusted according to the desired capacity value.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `The minimum size of the Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum size of the Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `the date and time when the schedule first starts.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `the date and time when the schedule end.`,
				},
				resource.Attribute{
					Name:        "recurrence",
					Description: `Repeat Settings.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group_no",
					Description: `The number of the auto scaling group. ~>`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `the time band for the repeat settings.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_block_storage",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Block Storage created before.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific block storage to retrieve.`,
				},
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `(Optional) The ID of the server instance associated with block storage to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "block_storage_no",
					Description: `The ID of Block Storage. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the Block Storage.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Block Storage.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Block Storage`,
				},
				resource.Attribute{
					Name:        "disk_detail_type",
					Description: `Type of Block Storage disk detail.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Server name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Block Storage type code.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `Device name.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `Block Storage product code.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Block Storage instance status code.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Disk type code. ~>`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Available zone where the Block Storage placed.`,
				},
				resource.Attribute{
					Name:        "snapshot_no",
					Description: `The ID of Block Storage Snapshot.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "block_storage_no",
					Description: `The ID of Block Storage. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the Block Storage.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Block Storage.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Block Storage`,
				},
				resource.Attribute{
					Name:        "disk_detail_type",
					Description: `Type of Block Storage disk detail.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Server name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Block Storage type code.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `Device name.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `Block Storage product code.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Block Storage instance status code.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Disk type code. ~>`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Available zone where the Block Storage placed.`,
				},
				resource.Attribute{
					Name:        "snapshot_no",
					Description: `The ID of Block Storage Snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_block_storage_snapshot",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Snapshot (Block Storage) created before.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific Snapshot to retrieve.`,
				},
				resource.Attribute{
					Name:        "block_storage_no",
					Description: `(Optional) The ID of the specific Block storage to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "snapshot_no",
					Description: `The ID of Snapshot. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of snapshot volume.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of snapshot.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_no",
					Description: `The ID of Snapshot. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of snapshot.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of snapshot volume.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_cdss_cluster",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attribute Reference In addition to all arguments above, the following attributes are exported`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster id.`,
				},
				resource.Attribute{
					Name:        "service_group_instance_no",
					Description: `Service Group Instance number.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Cluster name.`,
				},
				resource.Attribute{
					Name:        "kafka_version_code",
					Description: `Cloud Data Streaming Service version to be used.`,
				},
				resource.Attribute{
					Name:        "config_group_no",
					Description: `ConfigGroup number to be used.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `VPC number to be used.`,
				},
				resource.Attribute{
					Name:        "os_image",
					Description: `OS type to be used.`,
				},
				resource.Attribute{
					Name:        "cmak",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `CMAK access ID. Only lowercase alphanumeric characters and non-consecutive hyphens (-) allowed First character must be a letter, but the last character may be a letter or a number.`,
				},
				resource.Attribute{
					Name:        "user_password",
					Description: `CMAK access password. Must be at least 8 characters and contain at least one of each: English uppercase letter, lowercase letter, special character, and number.`,
				},
				resource.Attribute{
					Name:        "manager_node",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "node_product_code",
					Description: `HW specifications of the manager node.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet number where the manager node is to be located.`,
				},
				resource.Attribute{
					Name:        "broker_nodes",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "node_product_code",
					Description: `HW specifications of the broker node.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet number where the broker node is to be located.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of broker nodes. At least 3 units, up to 10 units allowed. (Can only be increased)`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Broker node storage capacity. At least 100 GB, up to 2000 GB. Must be in units of 10 GB.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "plaintext",
					Description: `List of broker nodes (Port 9092).`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `List of broker nodes (Port 9093).`,
				},
				resource.Attribute{
					Name:        "public_endpoint_plaintext",
					Description: `List of public endpoint of broker nodes.`,
				},
				resource.Attribute{
					Name:        "public_endpoint_plaintext_listener_port",
					Description: `List of listener port for public endpoint of broker nodes.`,
				},
				resource.Attribute{
					Name:        "public_endpoint_tls",
					Description: `List of public endpoint of broker nodes (TLS).`,
				},
				resource.Attribute{
					Name:        "public_endpoint_tls_listener_port",
					Description: `List of listener port for public endpoint of broker nodes (TLS).`,
				},
				resource.Attribute{
					Name:        "hosts_private_endpoint_tls",
					Description: `Editing details of the hosts file (Private IP hostname format).`,
				},
				resource.Attribute{
					Name:        "hosts_public_endpoint_tls",
					Description: `Editing details of the hosts file (Public IP hostname format).`,
				},
				resource.Attribute{
					Name:        "zookeeper",
					Description: `List of ZooKeeper nodes (Port 2181).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Cluster id.`,
				},
				resource.Attribute{
					Name:        "service_group_instance_no",
					Description: `Service Group Instance number.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Cluster name.`,
				},
				resource.Attribute{
					Name:        "kafka_version_code",
					Description: `Cloud Data Streaming Service version to be used.`,
				},
				resource.Attribute{
					Name:        "config_group_no",
					Description: `ConfigGroup number to be used.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `VPC number to be used.`,
				},
				resource.Attribute{
					Name:        "os_image",
					Description: `OS type to be used.`,
				},
				resource.Attribute{
					Name:        "cmak",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `CMAK access ID. Only lowercase alphanumeric characters and non-consecutive hyphens (-) allowed First character must be a letter, but the last character may be a letter or a number.`,
				},
				resource.Attribute{
					Name:        "user_password",
					Description: `CMAK access password. Must be at least 8 characters and contain at least one of each: English uppercase letter, lowercase letter, special character, and number.`,
				},
				resource.Attribute{
					Name:        "manager_node",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "node_product_code",
					Description: `HW specifications of the manager node.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet number where the manager node is to be located.`,
				},
				resource.Attribute{
					Name:        "broker_nodes",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "node_product_code",
					Description: `HW specifications of the broker node.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet number where the broker node is to be located.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of broker nodes. At least 3 units, up to 10 units allowed. (Can only be increased)`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Broker node storage capacity. At least 100 GB, up to 2000 GB. Must be in units of 10 GB.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "plaintext",
					Description: `List of broker nodes (Port 9092).`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `List of broker nodes (Port 9093).`,
				},
				resource.Attribute{
					Name:        "public_endpoint_plaintext",
					Description: `List of public endpoint of broker nodes.`,
				},
				resource.Attribute{
					Name:        "public_endpoint_plaintext_listener_port",
					Description: `List of listener port for public endpoint of broker nodes.`,
				},
				resource.Attribute{
					Name:        "public_endpoint_tls",
					Description: `List of public endpoint of broker nodes (TLS).`,
				},
				resource.Attribute{
					Name:        "public_endpoint_tls_listener_port",
					Description: `List of listener port for public endpoint of broker nodes (TLS).`,
				},
				resource.Attribute{
					Name:        "hosts_private_endpoint_tls",
					Description: `Editing details of the hosts file (Private IP hostname format).`,
				},
				resource.Attribute{
					Name:        "hosts_public_endpoint_tls",
					Description: `Editing details of the hosts file (Public IP hostname format).`,
				},
				resource.Attribute{
					Name:        "zookeeper",
					Description: `List of ZooKeeper nodes (Port 2181).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_cdss_config_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "kafka_version_code",
					Description: `(Required) Cloud Data Streaming Service version to be used.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attribute Reference In addition to all arguments above, the following attributes are exported`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ConfigGroup id.`,
				},
				resource.Attribute{
					Name:        "config_group_no",
					Description: `Config group number.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `ConfigGroup name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `ConfigGroup description.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ConfigGroup id.`,
				},
				resource.Attribute{
					Name:        "config_group_no",
					Description: `Config group number.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `ConfigGroup name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `ConfigGroup description.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_cdss_kafka_version",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of kafka version.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Kafka version name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of kafka version.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Kafka version name`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_cdss_kafka_versions",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "kafka_versions",
					Description: `A list of Kafka version`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "kafka_versions",
					Description: `A list of Kafka version`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_cdss_node_product",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "os_image",
					Description: `(Required) OS type to be used.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `(Required) Subnet number where the node will be located.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of server product.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `CPU count.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type code.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of server product.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `CPU count.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type code.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_cdss_node_products",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "os_image",
					Description: `(Required) OS type to be used.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `(Required) Subnet number where the node will be located. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "node_products",
					Description: `A list of Server product`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "node_products",
					Description: `A list of Server product`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_cdss_os_image",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of server image product.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Os image name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of server image product.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Os image name`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_cdss_os_images",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "os_images",
					Description: `A list of Os image`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "os_images",
					Description: `A list of Os image`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_init_script",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Init script created before.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific Init script to retrieve.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the specific Init script to retrieve.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional) Type of O/S to apply server instance. Accepted values: ` + "`" + `LNX` + "`" + ` (LINUX) | ` + "`" + `WND` + "`" + ` (WINDOWS)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "init_script_no",
					Description: `The ID of Init script. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Initialization script content.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Init script`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "init_script_no",
					Description: `The ID of Init script. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Initialization script content.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Init script`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_launch_configuration",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Launch Configuration created before.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific launch configuration to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "launch_configuration_no",
					Description: `The ID of Launch Configuration. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Launch Configuration.`,
				},
				resource.Attribute{
					Name:        "server_image_product_code",
					Description: `Server image product code.`,
				},
				resource.Attribute{
					Name:        "server_product_code",
					Description: `Server product code.`,
				},
				resource.Attribute{
					Name:        "member_server_image_no",
					Description: `The ID of Member server image.`,
				},
				resource.Attribute{
					Name:        "login_key_name",
					Description: `The login key name to encrypt with the public key.`,
				},
				resource.Attribute{
					Name:        "init_script_no",
					Description: `Set init script ID, The server can run a user-set initialization script at first boot. ~>`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `defined actionable scripts. ~>`,
				},
				resource.Attribute{
					Name:        "is_encrypted_volume",
					Description: `Whether to encrypt basic block storage if server image is RHV.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "launch_configuration_no",
					Description: `The ID of Launch Configuration. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Launch Configuration.`,
				},
				resource.Attribute{
					Name:        "server_image_product_code",
					Description: `Server image product code.`,
				},
				resource.Attribute{
					Name:        "server_product_code",
					Description: `Server product code.`,
				},
				resource.Attribute{
					Name:        "member_server_image_no",
					Description: `The ID of Member server image.`,
				},
				resource.Attribute{
					Name:        "login_key_name",
					Description: `The login key name to encrypt with the public key.`,
				},
				resource.Attribute{
					Name:        "init_script_no",
					Description: `Set init script ID, The server can run a user-set initialization script at first boot. ~>`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `defined actionable scripts. ~>`,
				},
				resource.Attribute{
					Name:        "is_encrypted_volume",
					Description: `Whether to encrypt basic block storage if server image is RHV.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_lb",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Load Balancer created before.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific load balancer to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "load_balancer_no",
					Description: `The ID of load balancer (It is the same result as id).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the load balancer.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The network type of load balancer.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `The time in seconds that the idle timeout.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of load balancer.`,
				},
				resource.Attribute{
					Name:        "throughput_type",
					Description: `The performance type code of load balancer.`,
				},
				resource.Attribute{
					Name:        "subnet_no_list",
					Description: `A list of IDs in the associated Subnets.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name of load balancer.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `A list of IP address of load balancer.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_no",
					Description: `The ID of load balancer (It is the same result as id).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the load balancer.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The network type of load balancer.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `The time in seconds that the idle timeout.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of load balancer.`,
				},
				resource.Attribute{
					Name:        "throughput_type",
					Description: `The performance type code of load balancer.`,
				},
				resource.Attribute{
					Name:        "subnet_no_list",
					Description: `A list of IDs in the associated Subnets.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name of load balancer.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `A list of IP address of load balancer.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_lb_listener",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Load Balancer Listener created before.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific listener to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listener_no",
					Description: `The ID of listener (It is the same result as id).`,
				},
				resource.Attribute{
					Name:        "rule_no_list",
					Description: `The list of listener rule number`,
				},
				resource.Attribute{
					Name:        "load_balancer_no",
					Description: `The ID of the load balancer.`,
				},
				resource.Attribute{
					Name:        "target_group_no",
					Description: `The ID of the target group.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port on which the load balancer is listening.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol type for the listener.`,
				},
				resource.Attribute{
					Name:        "tls_min_version_type",
					Description: `The TLS minimum supported version type code.`,
				},
				resource.Attribute{
					Name:        "use_http2",
					Description: `Whether to use HTTP/2 protocol.`,
				},
				resource.Attribute{
					Name:        "ssl_certificate_no",
					Description: `The ID of the SSL certificate.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_no",
					Description: `The ID of listener (It is the same result as id).`,
				},
				resource.Attribute{
					Name:        "rule_no_list",
					Description: `The list of listener rule number`,
				},
				resource.Attribute{
					Name:        "load_balancer_no",
					Description: `The ID of the load balancer.`,
				},
				resource.Attribute{
					Name:        "target_group_no",
					Description: `The ID of the target group.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port on which the load balancer is listening.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol type for the listener.`,
				},
				resource.Attribute{
					Name:        "tls_min_version_type",
					Description: `The TLS minimum supported version type code.`,
				},
				resource.Attribute{
					Name:        "use_http2",
					Description: `Whether to use HTTP/2 protocol.`,
				},
				resource.Attribute{
					Name:        "ssl_certificate_no",
					Description: `The ID of the SSL certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_lb_target_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Load Balancer Target Group created before.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific target group to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "target_group_no",
					Description: `The ID of target group (It is the same result as id).`,
				},
				resource.Attribute{
					Name:        "load_balancer_instance_no",
					Description: `The ID of the Load Balancer associated with the Target Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the target group.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port on which targets receive traffic.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol to use for routing traffic to the targets.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the target group.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `The health check to check the health of the target.`,
				},
				resource.Attribute{
					Name:        "cycle",
					Description: `The number of health check cycle.`,
				},
				resource.Attribute{
					Name:        "down_threshold",
					Description: `The number of health check failure threshold. You can determine the number of consecutive health check failures that are required before a health check is considered a failed state.`,
				},
				resource.Attribute{
					Name:        "up_threshold",
					Description: `The number of health check normal threshold. You can determine the number of consecutive health checks that are required before health checks are considered success state.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `The HTTP method for the health check. You can determine which HTTP method to use for health checks.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port to use for health checks.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The type of protocol to use for health checks.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `The URL path of the health check.`,
				},
				resource.Attribute{
					Name:        "target_no_list",
					Description: `The list of target number to bind to the target group.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `The type of target to be added to the target group.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the VPC in to create the target group.`,
				},
				resource.Attribute{
					Name:        "use_sticky_session",
					Description: `Whether to use session specific access.`,
				},
				resource.Attribute{
					Name:        "use_proxy_protocol",
					Description: `Whether to use a proxy protocol.`,
				},
				resource.Attribute{
					Name:        "algorithm_type",
					Description: `The type of algorithm to use for load balancing.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "target_group_no",
					Description: `The ID of target group (It is the same result as id).`,
				},
				resource.Attribute{
					Name:        "load_balancer_instance_no",
					Description: `The ID of the Load Balancer associated with the Target Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the target group.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port on which targets receive traffic.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol to use for routing traffic to the targets.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the target group.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `The health check to check the health of the target.`,
				},
				resource.Attribute{
					Name:        "cycle",
					Description: `The number of health check cycle.`,
				},
				resource.Attribute{
					Name:        "down_threshold",
					Description: `The number of health check failure threshold. You can determine the number of consecutive health check failures that are required before a health check is considered a failed state.`,
				},
				resource.Attribute{
					Name:        "up_threshold",
					Description: `The number of health check normal threshold. You can determine the number of consecutive health checks that are required before health checks are considered success state.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `The HTTP method for the health check. You can determine which HTTP method to use for health checks.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port to use for health checks.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The type of protocol to use for health checks.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `The URL path of the health check.`,
				},
				resource.Attribute{
					Name:        "target_no_list",
					Description: `The list of target number to bind to the target group.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `The type of target to be added to the target group.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the VPC in to create the target group.`,
				},
				resource.Attribute{
					Name:        "use_sticky_session",
					Description: `Whether to use session specific access.`,
				},
				resource.Attribute{
					Name:        "use_proxy_protocol",
					Description: `Whether to use a proxy protocol.`,
				},
				resource.Attribute{
					Name:        "algorithm_type",
					Description: `The type of algorithm to use for load balancing.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_member_server_image",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Gets a member server image.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the member server image list returned by ncloud`,
				},
				resource.Attribute{
					Name:        "no_list",
					Description: `(Optional) List of member server images to view`,
				},
				resource.Attribute{
					Name:        "platform_type_code_list",
					Description: `(Optional) List of platform codes of server images to view. Linux 32Bit (` + "`" + `LNX32` + "`" + `) | Linux 64Bit (` + "`" + `LNX64` + "`" + `) | Windows 32Bit (` + "`" + `WND32` + "`" + `) | Windows 64Bit (` + "`" + `WND64` + "`" + `) | Ubuntu Desktop 64Bit (` + "`" + `UBD64` + "`" + `) | Ubuntu Server 64Bit (` + "`" + `UBS64` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: ` + "`" + `KR` + "`" + ` region. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "no",
					Description: `Member server image no`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Member server image name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Member server image description`,
				},
				resource.Attribute{
					Name:        "original_server_instance_no",
					Description: `Original server instance no`,
				},
				resource.Attribute{
					Name:        "original_server_product_code",
					Description: `Original server product code`,
				},
				resource.Attribute{
					Name:        "original_server_name",
					Description: `Original server name`,
				},
				resource.Attribute{
					Name:        "original_base_block_storage_disk_type",
					Description: `Original base block storage disk type`,
				},
				resource.Attribute{
					Name:        "original_server_image_product_code",
					Description: `Original server image product code`,
				},
				resource.Attribute{
					Name:        "original_os_information",
					Description: `Original os information`,
				},
				resource.Attribute{
					Name:        "original_server_image_name",
					Description: `Original server image name`,
				},
				resource.Attribute{
					Name:        "status_name",
					Description: `Member server image status name`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Member server image status`,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `Member server image operation`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `Member server image platform type`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region info`,
				},
				resource.Attribute{
					Name:        "block_storage_total_rows",
					Description: `Member server image block storage total rows`,
				},
				resource.Attribute{
					Name:        "block_storage_total_size",
					Description: `Member server image block storage total size`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "no",
					Description: `Member server image no`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Member server image name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Member server image description`,
				},
				resource.Attribute{
					Name:        "original_server_instance_no",
					Description: `Original server instance no`,
				},
				resource.Attribute{
					Name:        "original_server_product_code",
					Description: `Original server product code`,
				},
				resource.Attribute{
					Name:        "original_server_name",
					Description: `Original server name`,
				},
				resource.Attribute{
					Name:        "original_base_block_storage_disk_type",
					Description: `Original base block storage disk type`,
				},
				resource.Attribute{
					Name:        "original_server_image_product_code",
					Description: `Original server image product code`,
				},
				resource.Attribute{
					Name:        "original_os_information",
					Description: `Original os information`,
				},
				resource.Attribute{
					Name:        "original_server_image_name",
					Description: `Original server image name`,
				},
				resource.Attribute{
					Name:        "status_name",
					Description: `Member server image status name`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Member server image status`,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `Member server image operation`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `Member server image platform type`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region info`,
				},
				resource.Attribute{
					Name:        "block_storage_total_rows",
					Description: `Member server image block storage total rows`,
				},
				resource.Attribute{
					Name:        "block_storage_total_size",
					Description: `Member server image block storage total size`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_member_server_images",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Gets a list of member server images.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the member server image list returned by ncloud`,
				},
				resource.Attribute{
					Name:        "no_list",
					Description: `(Optional) List of member server images to view`,
				},
				resource.Attribute{
					Name:        "platform_type_code_list",
					Description: `(Optional) List of platform codes of server images to view. Linux 32Bit (LNX32) | Linux 64Bit (LNX64) | Windows 32Bit (WND32) | Windows 64Bit (WND64) | Ubuntu Desktop 64Bit (UBD64) | Ubuntu Server 64Bit (UBS64)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: KR region.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save data source after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "member_server_images",
					Description: `A list of Member server image no`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member_server_images",
					Description: `A list of Member server image no`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_nas_volume",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Get NAS volume instance

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of specific NAS Volume to retrieve.`,
				},
				resource.Attribute{
					Name:        "volume_allotment_protocol_type_code",
					Description: `(Optional) Volume allotment protocol type code. All volume instances will be selected if the filter is not specified. (` + "`" + `NFS` + "`" + ` | ` + "`" + `CIFS` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "is_event_configuration",
					Description: `(Optional) Indicates whether the event is set. All volume instances will be selected if the filter is not specified. (` + "`" + `true` + "`" + ` | ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "is_snapshot_configuration",
					Description: `(Optional) Indicates whether a snapshot volume is set. All volume instances will be selected if the filter is not specified. (` + "`" + `true` + "`" + ` | ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "nas_volume_no",
					Description: `The ID of NAS Volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Volume name.`,
				},
				resource.Attribute{
					Name:        "volume_total_size",
					Description: `Volume total size.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `Volume size.`,
				},
				resource.Attribute{
					Name:        "snapshot_volume_size",
					Description: `Snapshot volume size.`,
				},
				resource.Attribute{
					Name:        "custom_ip_list",
					Description: `NAS volume instance custom IP list.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `NAS volume description.`,
				},
				resource.Attribute{
					Name:        "is_encrypted_volume",
					Description: `Volume encryption.`,
				},
				resource.Attribute{
					Name:        "mount_information",
					Description: `Mount information for NAS volume.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nas_volume_no",
					Description: `The ID of NAS Volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Volume name.`,
				},
				resource.Attribute{
					Name:        "volume_total_size",
					Description: `Volume total size.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `Volume size.`,
				},
				resource.Attribute{
					Name:        "snapshot_volume_size",
					Description: `Snapshot volume size.`,
				},
				resource.Attribute{
					Name:        "custom_ip_list",
					Description: `NAS volume instance custom IP list.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `NAS volume description.`,
				},
				resource.Attribute{
					Name:        "is_encrypted_volume",
					Description: `Volume encryption.`,
				},
				resource.Attribute{
					Name:        "mount_information",
					Description: `Mount information for NAS volume.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_nas_volumes",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Gets a list of NAS volume instances.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "volume_allotment_protocol_type_code",
					Description: `(Optional) Volume allotment protocol type code. All volume instances will be selected if the filter is not specified. (` + "`" + `NFS` + "`" + ` | ` + "`" + `CIFS` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "is_event_configuration",
					Description: `(Optional) Indicates whether the event is set. All volume instances will be selected if the filter is not specified. (` + "`" + `true` + "`" + ` | ` + "`" + `false` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "is_snapshot_configuration",
					Description: `(Optional) Indicates whether a snapshot volume is set. All volume instances will be selected if the filter is not specified. (` + "`" + `true` + "`" + ` | ` + "`" + `false` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "no_list",
					Description: `(Optional) List of nas volume instance numbers.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: KR region.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save data source after running ` + "`" + `terraform plan` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of NAS Volume ID.`,
				},
				resource.Attribute{
					Name:        "nas_volumes",
					Description: `A list of NAS Volume Instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of NAS Volume ID.`,
				},
				resource.Attribute{
					Name:        "nas_volumes",
					Description: `A list of NAS Volume Instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_nat_gateway",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can provide useful for get detail of NAT Gateway created before.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific NAT gateway to retrieve.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the specific NAT gateway to retrieve.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `(Optional) name of the specific associated VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nat_gateway_no",
					Description: `The ID of NAT gateway. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Available zone where the NAT gateway placed.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP on NAT Gateway created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of NAT Gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nat_gateway_no",
					Description: `The ID of NAT gateway. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Available zone where the NAT gateway placed.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP on NAT Gateway created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of NAT Gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_network_acl_deny_allow_groups",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This resource is useful for look up the list of Network ACL Deny-Allow Group in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_acl_deny_allow_group_no_list",
					Description: `(Optional) List of Deny-Allow Group ID to retrieve.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Optional) The ID of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) name of the specific Deny-Allow Group to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "network_acl_deny_allow_groups",
					Description: `The list of Deny-Allow Group ### Network ACL Deny Allow Group Reference ` + "`" + `network_acl_deny_allow_groups` + "`" + ` are also exported with the following attributes, where are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Deny-Allow Group.`,
				},
				resource.Attribute{
					Name:        "network_acl_deny_allow_group_no",
					Description: `The ID of Deny-Allow Group. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `list of IP address that registered in the Deny-Allow Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Deny-Allow Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Deny-Allow Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_acl_deny_allow_groups",
					Description: `The list of Deny-Allow Group ### Network ACL Deny Allow Group Reference ` + "`" + `network_acl_deny_allow_groups` + "`" + ` are also exported with the following attributes, where are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Deny-Allow Group.`,
				},
				resource.Attribute{
					Name:        "network_acl_deny_allow_group_no",
					Description: `The ID of Deny-Allow Group. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `list of IP address that registered in the Deny-Allow Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Deny-Allow Group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Deny-Allow Group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_network_acls",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This resource is useful for look up the list of Network ACL in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_acl_no_list",
					Description: `(Optional) List of Network ACL ID to retrieve.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Optional) The ID of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) name of the specific Network ACLs to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "network_acls",
					Description: `The list of Network ACL ### Network ACL Reference ` + "`" + `network_acls` + "`" + ` are also exported with the following attributes, where are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Network ACL.`,
				},
				resource.Attribute{
					Name:        "network_acl_no",
					Description: `The ID of Network ACL. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether default or not by VPC creation.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Network ACL.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Network ACL.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_acls",
					Description: `The list of Network ACL ### Network ACL Reference ` + "`" + `network_acls` + "`" + ` are also exported with the following attributes, where are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Network ACL.`,
				},
				resource.Attribute{
					Name:        "network_acl_no",
					Description: `The ID of Network ACL. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether default or not by VPC creation.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Network ACL.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Network ACL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_network_interface",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Network Interface created before.

~> **NOTE:** This resource only support VPC environment.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific Network Interface to retrieve.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the specific Network Interface to retrieve.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) The Private IP of the specific Network Interface to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "network_interface_no",
					Description: `The ID of Network Interface. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `The ID of the associated Subnet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Network Interface.`,
				},
				resource.Attribute{
					Name:        "access_control_groups",
					Description: `List of ACG ID applied to network interfaces.`,
				},
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `The ID of server instance assigned to network interface.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of Network Interface.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Type of server instance.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether default or not by Server instance creation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_interface_no",
					Description: `The ID of Network Interface. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `The ID of the associated Subnet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Network Interface.`,
				},
				resource.Attribute{
					Name:        "access_control_groups",
					Description: `List of ACG ID applied to network interfaces.`,
				},
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `The ID of server instance assigned to network interface.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of Network Interface.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Type of server instance.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether default or not by Server instance creation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_nks_cluster",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Provides a Kubernetes Service cluster data.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required) Cluster uuid. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Cluster name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster uuid.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `Control Plane API address.`,
				},
				resource.Attribute{
					Name:        "lb_private_subnet_no",
					Description: `Subnet No. for private loadbalancer only.`,
				},
				resource.Attribute{
					Name:        "lb_public_subnet_no",
					Description: `Subnet No. for public loadbalancer only. (Available only ` + "`" + `SGN` + "`" + ` region` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "subnet_no_list",
					Description: `Subnet No. list.`,
				},
				resource.Attribute{
					Name:        "public_network",
					Description: `Public Subnet Network`,
				},
				resource.Attribute{
					Name:        "kube_network_plugin",
					Description: `Kubernetes network plugin.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Cluster type. ` + "`" + `Maximum number of nodes` + "`" + ``,
				},
				resource.Attribute{
					Name:        "login_key_name",
					Description: `Login key name.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `zone Code.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `VPC No.`,
				},
				resource.Attribute{
					Name:        "audit",
					Description: `Audit log availability.`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `Kubenretes version.`,
				},
				resource.Attribute{
					Name:        "acg_no",
					Description: `The ID of cluster ACG.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Cluster name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster uuid.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `Control Plane API address.`,
				},
				resource.Attribute{
					Name:        "lb_private_subnet_no",
					Description: `Subnet No. for private loadbalancer only.`,
				},
				resource.Attribute{
					Name:        "lb_public_subnet_no",
					Description: `Subnet No. for public loadbalancer only. (Available only ` + "`" + `SGN` + "`" + ` region` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "subnet_no_list",
					Description: `Subnet No. list.`,
				},
				resource.Attribute{
					Name:        "public_network",
					Description: `Public Subnet Network`,
				},
				resource.Attribute{
					Name:        "kube_network_plugin",
					Description: `Kubernetes network plugin.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Cluster type. ` + "`" + `Maximum number of nodes` + "`" + ``,
				},
				resource.Attribute{
					Name:        "login_key_name",
					Description: `Login key name.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `zone Code.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `VPC No.`,
				},
				resource.Attribute{
					Name:        "audit",
					Description: `Audit log availability.`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `Kubenretes version.`,
				},
				resource.Attribute{
					Name:        "acg_no",
					Description: `The ID of cluster ACG.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_nks_clusters",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Provides list of Kubernetes Service cluster uuid.

`,
			Keywords:  []string{},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Ncloud Region.`,
				},
				resource.Attribute{
					Name:        "cluster_uuids",
					Description: `Set of NKS Clusters uuids`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_nks_kube_config",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Provides a kubeconfig from Kubernetes Service cluster.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_uuid",
					Description: `(Required) Cluster uuid. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster uuid.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host on kubeconfig.`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `Client certificate on kubeconfig.`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `Client key on kubeconfig.`,
				},
				resource.Attribute{
					Name:        "cluster_ca_certificate",
					Description: `Cluster CA certificate on kubeconfig.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Cluster uuid.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host on kubeconfig.`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `Client certificate on kubeconfig.`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `Client key on kubeconfig.`,
				},
				resource.Attribute{
					Name:        "cluster_ca_certificate",
					Description: `Cluster CA certificate on kubeconfig.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_nks_node_pool",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Provides a Kubernetes Service nodepool data.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "node_pool_name",
					Description: `(Required) The name of nodepool.`,
				},
				resource.Attribute{
					Name:        "cluster_uuid",
					Description: `(Required) Cluster uuid. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of nodepool.` + "`" + `CusterUuid:NodePoolName` + "`" + ``,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of nodes.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `Product code.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Auto scaling availability.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `Maximum number of nodes available for auto scaling.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Minimum number of nodes available for auto scaling.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet No.`,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `Nodepool instance No.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Server instance.`,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `The ID of server instance.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `Server spec.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP.`,
				},
				resource.Attribute{
					Name:        "node_status",
					Description: `Node Status.`,
				},
				resource.Attribute{
					Name:        "container_version",
					Description: `Container version of node.`,
				},
				resource.Attribute{
					Name:        "kernel_version",
					Description: `kernel version of node.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of nodepool.` + "`" + `CusterUuid:NodePoolName` + "`" + ``,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of nodes.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `Product code.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Auto scaling availability.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `Maximum number of nodes available for auto scaling.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Minimum number of nodes available for auto scaling.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet No.`,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `Nodepool instance No.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Server instance.`,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `The ID of server instance.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `Server spec.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP.`,
				},
				resource.Attribute{
					Name:        "node_status",
					Description: `Node Status.`,
				},
				resource.Attribute{
					Name:        "container_version",
					Description: `Container version of node.`,
				},
				resource.Attribute{
					Name:        "kernel_version",
					Description: `kernel version of node.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_nks_node_pools",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Provides list of Kubernetes Service nodepool name.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_uuid",
					Description: `(Required) Cluster uuid. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster uuid.`,
				},
				resource.Attribute{
					Name:        "node_pool_names",
					Description: `Set of all node pool names in NKS Clusters.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Cluster uuid.`,
				},
				resource.Attribute{
					Name:        "node_pool_names",
					Description: `Set of all node pool names in NKS Clusters.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_nks_versions",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Provides list of available Kubernetes Service versions.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `A list of verions`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Version label`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Version value`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "versions",
					Description: `A list of verions`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Version label`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Version value`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_placement_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Placement group created before.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of specific Placement group to retrieve.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of specific Placement group to retrieve.`,
				},
				resource.Attribute{
					Name:        "placement_group_type",
					Description: `(Optional) Type of placement group to retrieve. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "placement_group_no",
					Description: `The ID of Placement group. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "placement_group_no",
					Description: `The ID of Placement group. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_port_forwarding_rule",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Get a port forwarding rule.
When a server is created for the first time, a public IP address for port forwarding is given per account.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: KR region.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. You can decide a zone where servers are created. You can decide which zone the product list will be requested in. Default : Select the first Zone in the specific region Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `Filter by server instance number`,
				},
				resource.Attribute{
					Name:        "port_forwarding_internal_port",
					Description: `(Optional) Port forwarding internal port.`,
				},
				resource.Attribute{
					Name:        "port_forwarding_external_port",
					Description: `Port forwarding external port. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "port_forwarding_configuration_no",
					Description: `Port forwarding configuration number`,
				},
				resource.Attribute{
					Name:        "port_forwarding_public_ip",
					Description: `Port forwarding public ip`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "port_forwarding_configuration_no",
					Description: `Port forwarding configuration number`,
				},
				resource.Attribute{
					Name:        "port_forwarding_public_ip",
					Description: `Port forwarding public ip`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_port_forwarding_rules",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Gets a list of port forwarding rules.
When a server is created for the first time, a public IP address for port forwarding is given per account.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: KR region.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. You can decide a zone where servers are created. You can decide which zone the product list will be requested in. Default : Select the first Zone in the specific region Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port_forwarding_internal_port",
					Description: `(Optional) Port forwarding internal port.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save data source after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "port_forwarding_configuration_no",
					Description: `Port forwarding configuration number`,
				},
				resource.Attribute{
					Name:        "port_forwarding_public_ip",
					Description: `Port forwarding public ip`,
				},
				resource.Attribute{
					Name:        "port_forwarding_rule_list",
					Description: `Port forwarding rule list`,
				},
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `Server instance number`,
				},
				resource.Attribute{
					Name:        "port_forwarding_external_port",
					Description: `Port forwarding external port.`,
				},
				resource.Attribute{
					Name:        "port_forwarding_internal_port",
					Description: `Port forwarding internal port.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "port_forwarding_configuration_no",
					Description: `Port forwarding configuration number`,
				},
				resource.Attribute{
					Name:        "port_forwarding_public_ip",
					Description: `Port forwarding public ip`,
				},
				resource.Attribute{
					Name:        "port_forwarding_rule_list",
					Description: `Port forwarding rule list`,
				},
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `Server instance number`,
				},
				resource.Attribute{
					Name:        "port_forwarding_external_port",
					Description: `Port forwarding external port.`,
				},
				resource.Attribute{
					Name:        "port_forwarding_internal_port",
					Description: `Port forwarding internal port.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_public_ip",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Get public IP instance.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific Public IP instance to retrieve.`,
				},
				resource.Attribute{
					Name:        "is_associated",
					Description: `(Optional) Indicates whether the public IP address is associated or not. ~>`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. You can filter the list of public IP instances by zone. All the public IP addresses in the zone of the region will be selected if the filter is not specified. Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "public_ip_no",
					Description: `The ID of Public IP. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Public IP description`,
				},
				resource.Attribute{
					Name:        "kind_type",
					Description: `Public IP kind type`,
				},
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `Associated server instance number`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Associated server name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "public_ip_no",
					Description: `The ID of Public IP. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Public IP description`,
				},
				resource.Attribute{
					Name:        "kind_type",
					Description: `Public IP kind type`,
				},
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `Associated server instance number`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Associated server name`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_regions",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Gets a list of available regions.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "code",
					Description: `(Optional) region code for filtering`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save data source after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A List of region`,
				},
				resource.Attribute{
					Name:        "region_no",
					Description: `Region number`,
				},
				resource.Attribute{
					Name:        "region_code",
					Description: `Region code`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Region name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `A List of region`,
				},
				resource.Attribute{
					Name:        "region_no",
					Description: `Region number`,
				},
				resource.Attribute{
					Name:        "region_code",
					Description: `Region code`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Region name`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_root_password",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Gets the password of a root account with the server's login key.

~> **Note:** All arguments including the private key will be stored in the raw state as plain-text.
[Read more about sensitive data in state](/docs/state/sensitive-data.html).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `(Required) Server instance number`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) Server’s login key (auth key) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `password of a root account`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "root_password",
					Description: `password of a root account`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_route_table",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Route Table created before.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific Route Table to retrieve.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Optional) The ID of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "supported_subnet_type",
					Description: `(Optional) Subnet type. Accepted values : ` + "`" + `PUBLIC` + "`" + ` (Public) | ` + "`" + `PRIVATE` + "`" + ` (Private).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) name of the specific Route Table to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "route_table_no",
					Description: `The ID of Route Table. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Route Table.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether is default or not by VPC creation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table_no",
					Description: `The ID of Route Table. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Route Table.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether is default or not by VPC creation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_route_tables",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This resource is useful for look up the list of Route table in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Optional) The ID of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "supported_subnet_type",
					Description: `(Optional) Subnet type. Accepted values : ` + "`" + `PUBLIC` + "`" + ` (Public) | ` + "`" + `PRIVATE` + "`" + ` (Private).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the specific Route Table to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "route_tables",
					Description: `The list of Route Tables ### Route Table Reference ` + "`" + `route_tables` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Route Table.`,
				},
				resource.Attribute{
					Name:        "route_table_no",
					Description: `The ID of Route Table. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "supported_subnet_type",
					Description: `Subnet type. Accepted values : ` + "`" + `PUBLIC` + "`" + ` (Public) | ` + "`" + `PRIVATE` + "`" + ` (Private).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Route Table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Route Table.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "route_tables",
					Description: `The list of Route Tables ### Route Table Reference ` + "`" + `route_tables` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Route Table.`,
				},
				resource.Attribute{
					Name:        "route_table_no",
					Description: `The ID of Route Table. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "supported_subnet_type",
					Description: `Subnet type. Accepted values : ` + "`" + `PUBLIC` + "`" + ` (Public) | ` + "`" + `PRIVATE` + "`" + ` (Private).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Route Table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of Route Table.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_server",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Server instance created before.

`,
			Keywords:  []string{},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_no",
					Description: `The ID of Server instance. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Server instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the server.`,
				},
				resource.Attribute{
					Name:        "server_image_product_code",
					Description: `Server image product code.`,
				},
				resource.Attribute{
					Name:        "server_product_code",
					Description: `Server product code.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `number of CPUs`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `The size of the memory in bytes.`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `Platform type code`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP`,
				},
				resource.Attribute{
					Name:        "base_block_storage_disk_type",
					Description: `Base block storage disk type code`,
				},
				resource.Attribute{
					Name:        "base_block_storage_disk_detail_type",
					Description: `Base block storage disk detail type code`,
				},
				resource.Attribute{
					Name:        "member_server_image_no",
					Description: `The ID of Member server image.`,
				},
				resource.Attribute{
					Name:        "login_key_name",
					Description: `The login key name to encrypt with the public key.`,
				},
				resource.Attribute{
					Name:        "is_protect_server_termination",
					Description: `Whether is protect return when creating.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Available zone where the Server instance placed. ~>`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP`,
				},
				resource.Attribute{
					Name:        "server_image_name",
					Description: `Server image name`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `The size of base block storage in bytes`,
				},
				resource.Attribute{
					Name:        "port_forwarding_public_ip",
					Description: `Port forwarding public ip`,
				},
				resource.Attribute{
					Name:        "port_forwarding_external_port",
					Description: `Port forwarding external port`,
				},
				resource.Attribute{
					Name:        "port_forwarding_internal_port",
					Description: `Port forwarding internal port ~>`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `The ID of the associated Subnet.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `List of Network Interface.`,
				},
				resource.Attribute{
					Name:        "network_interface_no",
					Description: `The ID of Network interface.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Order of network interfaces to be assigned to the server to create.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet ID of the network interface.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `IP address of the network interface.`,
				},
				resource.Attribute{
					Name:        "init_script_no",
					Description: `The ID of Init script.`,
				},
				resource.Attribute{
					Name:        "placement_group_no",
					Description: `The ID of Physical placement group.`,
				},
				resource.Attribute{
					Name:        "is_encrypted_base_block_storage_volume",
					Description: `Whether to encrypt basic block storage if server image is RHV.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_server_image",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

To create a server instance (VM), you should select a server image. This data source gets a server image.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "product_code",
					Description: `(Optional) Product code you want to view on the list. Use this when searching for 1 product.`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `(Optional) Values required for identifying platform. The available values are as follows: Linux 32Bit(LNX32) | Linux 64Bit(LNX64) | Windows 32Bit(WND32) | Windows 64Bit(WND64) | Ubuntu Desktop 64Bit(UBD64) | Ubuntu Server 64Bit(UBS64)`,
				},
				resource.Attribute{
					Name:        "infra_resource_detail_type_code",
					Description: `(Optional) infra resource detail type code.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of server image product.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Product name`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type code.`,
				},
				resource.Attribute{
					Name:        "product_description",
					Description: `Product description.`,
				},
				resource.Attribute{
					Name:        "infra_resource_type",
					Description: `Infra resource type code.`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `Base block storage size.`,
				},
				resource.Attribute{
					Name:        "os_information",
					Description: `OS Information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of server image product.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Product name`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type code.`,
				},
				resource.Attribute{
					Name:        "product_description",
					Description: `Product description.`,
				},
				resource.Attribute{
					Name:        "infra_resource_type",
					Description: `Infra resource type code.`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `Base block storage size.`,
				},
				resource.Attribute{
					Name:        "os_information",
					Description: `OS Information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_server_images",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

To create a server instance (VM), you should select a server image. This data source gets a list of server images.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "product_code",
					Description: `(Optional) Product code you want to view on the list. Use this when searching for 1 product.`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `(Optional) Values required for identifying platform. The available values are as follows: Linux 32Bit(LNX32) | Linux 64Bit(LNX64) | Windows 32Bit(WND32) | Windows 64Bit(WND64) | Ubuntu Desktop 64Bit(UBD64) | Ubuntu Server 64Bit(UBS64)`,
				},
				resource.Attribute{
					Name:        "infra_resource_detail_type_code",
					Description: `(Optional) infra resource detail type code.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save data source after running ` + "`" + `terraform plan` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A List of server image product code.`,
				},
				resource.Attribute{
					Name:        "server_images",
					Description: `A List of server image product. ### Server Image Product Reference ` + "`" + `server_images` + "`" + ` are also exported with the following attributes, when there are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of server image product.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `The ID of server image product. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Product name.`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type code.`,
				},
				resource.Attribute{
					Name:        "product_description",
					Description: `Product description.`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `Platform type code. The available values are as follows: Linux 32Bit(LNX32) | Linux 64Bit(LNX64) | Windows 32Bit(WND32) | Windows 64Bit(WND64) | Ubuntu Desktop 64Bit(UBD64) | Ubuntu Server 64Bit(UBS64).`,
				},
				resource.Attribute{
					Name:        "infra_resource_detail_type_code",
					Description: `infra resource detail type code.`,
				},
				resource.Attribute{
					Name:        "infra_resource_type",
					Description: `Infra resource type code.`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `Base block storage size.`,
				},
				resource.Attribute{
					Name:        "os_information",
					Description: `OS Information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A List of server image product code.`,
				},
				resource.Attribute{
					Name:        "server_images",
					Description: `A List of server image product. ### Server Image Product Reference ` + "`" + `server_images` + "`" + ` are also exported with the following attributes, when there are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of server image product.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `The ID of server image product. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Product name.`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type code.`,
				},
				resource.Attribute{
					Name:        "product_description",
					Description: `Product description.`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `Platform type code. The available values are as follows: Linux 32Bit(LNX32) | Linux 64Bit(LNX64) | Windows 32Bit(WND32) | Windows 64Bit(WND64) | Ubuntu Desktop 64Bit(UBD64) | Ubuntu Server 64Bit(UBS64).`,
				},
				resource.Attribute{
					Name:        "infra_resource_detail_type_code",
					Description: `infra resource detail type code.`,
				},
				resource.Attribute{
					Name:        "infra_resource_type",
					Description: `Infra resource type code.`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `Base block storage size.`,
				},
				resource.Attribute{
					Name:        "os_information",
					Description: `OS Information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_server_product",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

You should select a server product (server specification) to create a server instance (VM).
To this end, we provide data source by which you can search a server product.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server_image_product_code",
					Description: `(Required) You can get one from ` + "`" + `data ncloud_server_images` + "`" + `. This is a required value, and each available server's specification varies depending on the server image product.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `(Optional) Enter a product code to search from the list. Use it for a single search.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. You can decide a zone where servers are created. You can decide which zone the product list will be requested in. default : Select the first Zone in the specific region. Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of server product.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Product name.`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type code.`,
				},
				resource.Attribute{
					Name:        "product_description",
					Description: `Product description.`,
				},
				resource.Attribute{
					Name:        "infra_resource_type",
					Description: `Infra resource type code.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `CPU count.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Disk type.`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `Base block storage size.`,
				},
				resource.Attribute{
					Name:        "generation_code",
					Description: `Generation Code.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of server product.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Product name.`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type code.`,
				},
				resource.Attribute{
					Name:        "product_description",
					Description: `Product description.`,
				},
				resource.Attribute{
					Name:        "infra_resource_type",
					Description: `Infra resource type code.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `CPU count.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Disk type.`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `Base block storage size.`,
				},
				resource.Attribute{
					Name:        "generation_code",
					Description: `Generation Code.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_server_products",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

You should select a server product (server specification) to create a server instance (VM).
To this end, we provide data source by which you can search a server product.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server_image_product_code",
					Description: `(Required) You can get one from ` + "`" + `data ncloud_server_images` + "`" + `. This is a required value, and each available server's specification varies depending on the server image product.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `(Optional) Enter a product code to search from the list. Use it for a single search.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. You can decide a zone where servers are created. You can decide which zone the product list will be requested in. default : Select the first Zone in the specific region. Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A List of server product code.`,
				},
				resource.Attribute{
					Name:        "server_products",
					Description: `A List of server product. ### Server Product Reference ` + "`" + `server_products` + "`" + ` are also exported with the following attributes, when there are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of server product.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `The ID of server product. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Product name.`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type code.`,
				},
				resource.Attribute{
					Name:        "product_description",
					Description: `Product description.`,
				},
				resource.Attribute{
					Name:        "infra_resource_type",
					Description: `Infra resource type code.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `CPU count.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Disk type.`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `Base block storage size.`,
				},
				resource.Attribute{
					Name:        "generation_code",
					Description: `Generation Code.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A List of server product code.`,
				},
				resource.Attribute{
					Name:        "server_products",
					Description: `A List of server product. ### Server Product Reference ` + "`" + `server_products` + "`" + ` are also exported with the following attributes, when there are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of server product.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `The ID of server product. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Product name.`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type code.`,
				},
				resource.Attribute{
					Name:        "product_description",
					Description: `Product description.`,
				},
				resource.Attribute{
					Name:        "infra_resource_type",
					Description: `Infra resource type code.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `CPU count.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Disk type.`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `Base block storage size.`,
				},
				resource.Attribute{
					Name:        "generation_code",
					Description: `Generation Code.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_servers",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Use this data source to get multiple ` + "`" + `ncloud_server` + "`" + ` ids 

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_ses_cluster",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Provides a Search Engine Service cluster data.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Cluster Instance No. ## Attribute Reference In addition to all arguments above, the following attributes are exported`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Cluster name.`,
				},
				resource.Attribute{
					Name:        "service_group_instance_no",
					Description: `Cluster Instance No.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster Instance No.`,
				},
				resource.Attribute{
					Name:        "os_image_code",
					Description: `OS type to be used.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `VPC number to be used.`,
				},
				resource.Attribute{
					Name:        "search_engine",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "version_code",
					Description: `Search Engine Service version to be used.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Search Engine UserName. Only lowercase alphanumeric characters and non-consecutive hyphens (-) allowed First character must be a letter, but the last character may be a letter or a number.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Search Engine Port.`,
				},
				resource.Attribute{
					Name:        "dashboard_port",
					Description: `Search Engine Dashboard Port.`,
				},
				resource.Attribute{
					Name:        "manager_node",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "is_dual_manager",
					Description: `Redundancy of manager node`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `HW specifications of the manager node.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet number where the manager node is to be located.`,
				},
				resource.Attribute{
					Name:        "acg_id",
					Description: `The ID of manager node ACG.`,
				},
				resource.Attribute{
					Name:        "acg_name",
					Description: `The name of manager node ACG.`,
				},
				resource.Attribute{
					Name:        "data_node",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `HW specifications of the data node.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet number where the data node is to be located.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of data nodes. At least 3 units, up to 10 units allowed.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Data node storage capacity.`,
				},
				resource.Attribute{
					Name:        "acg_id",
					Description: `The ID of data node ACG.`,
				},
				resource.Attribute{
					Name:        "acg_name",
					Description: `The name of data node ACG.`,
				},
				resource.Attribute{
					Name:        "master_node(Optional)",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `HW specifications of the master node.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet number where the master node is to be located.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of master nodes.`,
				},
				resource.Attribute{
					Name:        "acg_id",
					Description: `The ID of master node ACG.`,
				},
				resource.Attribute{
					Name:        "acg_name",
					Description: `The name of master node ACG.`,
				},
				resource.Attribute{
					Name:        "manager_node_instance_no_list",
					Description: `List of Manager node's instance number`,
				},
				resource.Attribute{
					Name:        "cluster_node_list",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "compute_instance_name",
					Description: `The name of Server instance.`,
				},
				resource.Attribute{
					Name:        "login_key_name",
					Description: `Required Login key to access Manager node server`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Cluster name.`,
				},
				resource.Attribute{
					Name:        "service_group_instance_no",
					Description: `Cluster Instance No.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster Instance No.`,
				},
				resource.Attribute{
					Name:        "os_image_code",
					Description: `OS type to be used.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `VPC number to be used.`,
				},
				resource.Attribute{
					Name:        "search_engine",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "version_code",
					Description: `Search Engine Service version to be used.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Search Engine UserName. Only lowercase alphanumeric characters and non-consecutive hyphens (-) allowed First character must be a letter, but the last character may be a letter or a number.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Search Engine Port.`,
				},
				resource.Attribute{
					Name:        "dashboard_port",
					Description: `Search Engine Dashboard Port.`,
				},
				resource.Attribute{
					Name:        "manager_node",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "is_dual_manager",
					Description: `Redundancy of manager node`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `HW specifications of the manager node.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet number where the manager node is to be located.`,
				},
				resource.Attribute{
					Name:        "acg_id",
					Description: `The ID of manager node ACG.`,
				},
				resource.Attribute{
					Name:        "acg_name",
					Description: `The name of manager node ACG.`,
				},
				resource.Attribute{
					Name:        "data_node",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `HW specifications of the data node.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet number where the data node is to be located.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of data nodes. At least 3 units, up to 10 units allowed.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Data node storage capacity.`,
				},
				resource.Attribute{
					Name:        "acg_id",
					Description: `The ID of data node ACG.`,
				},
				resource.Attribute{
					Name:        "acg_name",
					Description: `The name of data node ACG.`,
				},
				resource.Attribute{
					Name:        "master_node(Optional)",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `HW specifications of the master node.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet number where the master node is to be located.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of master nodes.`,
				},
				resource.Attribute{
					Name:        "acg_id",
					Description: `The ID of master node ACG.`,
				},
				resource.Attribute{
					Name:        "acg_name",
					Description: `The name of master node ACG.`,
				},
				resource.Attribute{
					Name:        "manager_node_instance_no_list",
					Description: `List of Manager node's instance number`,
				},
				resource.Attribute{
					Name:        "cluster_node_list",
					Description: `.`,
				},
				resource.Attribute{
					Name:        "compute_instance_name",
					Description: `The name of Server instance.`,
				},
				resource.Attribute{
					Name:        "login_key_name",
					Description: `Required Login key to access Manager node server`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_ses_clusters",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Provides list of Search Engine Service cluster uuid.

`,
			Keywords:  []string{},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "clusters",
					Description: `A List of Search Engine Service cluster. ### Search Engine Service Cluster Reference ` + "`" + `clusters` + "`" + ` are also exported with the following attributes, when there are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster Instance No.`,
				},
				resource.Attribute{
					Name:        "service_group_instance_no",
					Description: `Cluster Instance No(Same as Cluster Id).`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Cluster name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_ses_node_os_images",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Provides list of available Server OS images.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `A List of OS image product. ### OS Image Product Reference ` + "`" + `images` + "`" + ` are also exported with the following attributes, when there are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of OS image product.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `OS image name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "images",
					Description: `A List of OS image product. ### OS Image Product Reference ` + "`" + `images` + "`" + ` are also exported with the following attributes, when there are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of OS image product.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `OS image name`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_ses_node_products",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Provides list of available Server product.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "os_image_code",
					Description: `(Required) OS type to be used.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `(Required) Subnet number where the node will be located.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "codes",
					Description: `A List of server product. ### Node Product Reference ` + "`" + `codes` + "`" + ` are also exported with the following attributes, when there are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The value of server product code.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `CPU count.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Product name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "codes",
					Description: `A List of server product. ### Node Product Reference ` + "`" + `codes` + "`" + ` are also exported with the following attributes, when there are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The value of server product code.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `CPU count.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Product name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_ses_versions",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Provides list of available Search Engine Service versions.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `A List of SES Version. ### Search Engine Service Version Reference ` + "`" + `versions` + "`" + ` are also exported with the following attributes, when there are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Code of SES Version`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `SES version name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `SES version type`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `SES version`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "versions",
					Description: `A List of SES Version. ### Search Engine Service Version Reference ` + "`" + `versions` + "`" + ` are also exported with the following attributes, when there are relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Code of SES Version`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `SES version name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `SES version type`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `SES version`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcebuild_project",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note** This data source only supports 'public' site.

~> **Note:** This data source is a beta release. Some features may change in the future.

This data source is useful for look up Sourcebuild project detail in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Sourcebuild Project ID. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "project_no",
					Description: `Sourcebuild Project ID. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Sourcebuild Project.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Sourcebuild project description.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Build target's type and config.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Build target type.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `Build target config.`,
				},
				resource.Attribute{
					Name:        "repository_name",
					Description: `Repository name to build.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `Branch to build.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Build environment.`,
				},
				resource.Attribute{
					Name:        "compute",
					Description: `Computing environment to build.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Computing type id.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `CPU of build environment.`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `Memory of build environment.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `Information about the build environment image.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Build environment image type.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `Build environment image config.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `OS config.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `OS id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `OS name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `OS version.`,
				},
				resource.Attribute{
					Name:        "archi",
					Description: `OS architecture.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `Runtime config.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `runtime id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `runtime name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `runtime version.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `runtime version id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `runtime version name.`,
				},
				resource.Attribute{
					Name:        "registry",
					Description: `Registry name of NCP Container Registry where the image to build is located.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `Container image name to build.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Container image tag to build.`,
				},
				resource.Attribute{
					Name:        "docker_engine",
					Description: `Docker engine to use when building docker image.`,
				},
				resource.Attribute{
					Name:        "use",
					Description: `Whether or not to use of docker engine.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Docker engine id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Docker engine name.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Build timeout (in Minutes).`,
				},
				resource.Attribute{
					Name:        "env_var",
					Description: `Environment variable to use for build.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key of environment variable.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of environment variable.`,
				},
				resource.Attribute{
					Name:        "build_command",
					Description: `Commands to execute in build.`,
				},
				resource.Attribute{
					Name:        "pre_build",
					Description: `Commands before build.`,
				},
				resource.Attribute{
					Name:        "in_build",
					Description: `Commands during build.`,
				},
				resource.Attribute{
					Name:        "post_build",
					Description: `Commands after build.`,
				},
				resource.Attribute{
					Name:        "docker_image_build",
					Description: `Docker image build config.`,
				},
				resource.Attribute{
					Name:        "use",
					Description: `Whether or not to use of dockerbuild.`,
				},
				resource.Attribute{
					Name:        "dockerfile",
					Description: `Dockerfile path in build source folder.`,
				},
				resource.Attribute{
					Name:        "registry",
					Description: `Registry name of NCP Container Registry to store the image.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `Image name to upload to registry.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Image tag to upload to registry.`,
				},
				resource.Attribute{
					Name:        "latest",
					Description: `Save status of the latest tag.`,
				},
				resource.Attribute{
					Name:        "artifact",
					Description: `Artifact to save build results.`,
				},
				resource.Attribute{
					Name:        "use",
					Description: `Whether or not to save build results.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Location to save build results.`,
				},
				resource.Attribute{
					Name:        "object_storage_to_upload",
					Description: `Object Storage to save build results.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Bucket name of NCP Object Storage to save build results.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `path in the NCP Object Storage bucket to save build results.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `File name to save build results.`,
				},
				resource.Attribute{
					Name:        "backup",
					Description: `Whether or not to backup build results.`,
				},
				resource.Attribute{
					Name:        "build_image_upload",
					Description: `Save build environment after completing this build.`,
				},
				resource.Attribute{
					Name:        "use",
					Description: `Whether or not to save build environment.`,
				},
				resource.Attribute{
					Name:        "container_registry_name",
					Description: `Registry name of NCP Container Registry to store the image of the build environment after completing the build.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Image name to upload to registry.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Image tag to upload to registry.`,
				},
				resource.Attribute{
					Name:        "latest",
					Description: `Save status of the latest tag.`,
				},
				resource.Attribute{
					Name:        "linked",
					Description: `Set up linkage with other services related this build.`,
				},
				resource.Attribute{
					Name:        "cloud_log_analytics",
					Description: `Whether or not to save build log in the NCP Cloud Log Analytics.`,
				},
				resource.Attribute{
					Name:        "cloud_log_analytics",
					Description: `Whether or not to check safety using NCP File Safer.`,
				},
				resource.Attribute{
					Name:        "lastbuild",
					Description: `Information of last build.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of last build.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of last build.`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `Time of last build.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Information about creating a Sourcebuild project.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Created user`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `Created time`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "project_no",
					Description: `Sourcebuild Project ID. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Sourcebuild Project.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Sourcebuild project description.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Build target's type and config.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Build target type.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `Build target config.`,
				},
				resource.Attribute{
					Name:        "repository_name",
					Description: `Repository name to build.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `Branch to build.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Build environment.`,
				},
				resource.Attribute{
					Name:        "compute",
					Description: `Computing environment to build.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Computing type id.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `CPU of build environment.`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `Memory of build environment.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `Information about the build environment image.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Build environment image type.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `Build environment image config.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `OS config.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `OS id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `OS name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `OS version.`,
				},
				resource.Attribute{
					Name:        "archi",
					Description: `OS architecture.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `Runtime config.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `runtime id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `runtime name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `runtime version.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `runtime version id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `runtime version name.`,
				},
				resource.Attribute{
					Name:        "registry",
					Description: `Registry name of NCP Container Registry where the image to build is located.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `Container image name to build.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Container image tag to build.`,
				},
				resource.Attribute{
					Name:        "docker_engine",
					Description: `Docker engine to use when building docker image.`,
				},
				resource.Attribute{
					Name:        "use",
					Description: `Whether or not to use of docker engine.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Docker engine id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Docker engine name.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Build timeout (in Minutes).`,
				},
				resource.Attribute{
					Name:        "env_var",
					Description: `Environment variable to use for build.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key of environment variable.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of environment variable.`,
				},
				resource.Attribute{
					Name:        "build_command",
					Description: `Commands to execute in build.`,
				},
				resource.Attribute{
					Name:        "pre_build",
					Description: `Commands before build.`,
				},
				resource.Attribute{
					Name:        "in_build",
					Description: `Commands during build.`,
				},
				resource.Attribute{
					Name:        "post_build",
					Description: `Commands after build.`,
				},
				resource.Attribute{
					Name:        "docker_image_build",
					Description: `Docker image build config.`,
				},
				resource.Attribute{
					Name:        "use",
					Description: `Whether or not to use of dockerbuild.`,
				},
				resource.Attribute{
					Name:        "dockerfile",
					Description: `Dockerfile path in build source folder.`,
				},
				resource.Attribute{
					Name:        "registry",
					Description: `Registry name of NCP Container Registry to store the image.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `Image name to upload to registry.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Image tag to upload to registry.`,
				},
				resource.Attribute{
					Name:        "latest",
					Description: `Save status of the latest tag.`,
				},
				resource.Attribute{
					Name:        "artifact",
					Description: `Artifact to save build results.`,
				},
				resource.Attribute{
					Name:        "use",
					Description: `Whether or not to save build results.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Location to save build results.`,
				},
				resource.Attribute{
					Name:        "object_storage_to_upload",
					Description: `Object Storage to save build results.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `Bucket name of NCP Object Storage to save build results.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `path in the NCP Object Storage bucket to save build results.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `File name to save build results.`,
				},
				resource.Attribute{
					Name:        "backup",
					Description: `Whether or not to backup build results.`,
				},
				resource.Attribute{
					Name:        "build_image_upload",
					Description: `Save build environment after completing this build.`,
				},
				resource.Attribute{
					Name:        "use",
					Description: `Whether or not to save build environment.`,
				},
				resource.Attribute{
					Name:        "container_registry_name",
					Description: `Registry name of NCP Container Registry to store the image of the build environment after completing the build.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Image name to upload to registry.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Image tag to upload to registry.`,
				},
				resource.Attribute{
					Name:        "latest",
					Description: `Save status of the latest tag.`,
				},
				resource.Attribute{
					Name:        "linked",
					Description: `Set up linkage with other services related this build.`,
				},
				resource.Attribute{
					Name:        "cloud_log_analytics",
					Description: `Whether or not to save build log in the NCP Cloud Log Analytics.`,
				},
				resource.Attribute{
					Name:        "cloud_log_analytics",
					Description: `Whether or not to check safety using NCP File Safer.`,
				},
				resource.Attribute{
					Name:        "lastbuild",
					Description: `Information of last build.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of last build.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of last build.`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `Time of last build.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Information about creating a Sourcebuild project.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Created user`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `Created time`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcebuild_project_computes",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note** This data source only supports 'public' site.

~> **Note:** This data source is a beta release. Some features may change in the future.

This data source is useful for look up the list of Sourcebuild compute in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "computes",
					Description: `Computing environments available at Sourcebuild. ### Computes Reference ` + "`" + `computes` + "`" + ` is also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Compute ID.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `CPU of build environment.`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `Memory of build environment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "computes",
					Description: `Computing environments available at Sourcebuild. ### Computes Reference ` + "`" + `computes` + "`" + ` is also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Compute ID.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `CPU of build environment.`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `Memory of build environment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcebuild_project_docker_engines",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note** This data source only supports 'public' site.

~> **Note:** This data source is a beta release. Some features may change in the future.

This data source is useful for look up the list of Sourcebuild docker engine in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "docker_engines",
					Description: `Docker engines available at Sourcebuild. ### Docker Engines Reference ` + "`" + `docker_engines` + "`" + ` is also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Docker engine ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Docker engine name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "docker_engines",
					Description: `Docker engines available at Sourcebuild. ### Docker Engines Reference ` + "`" + `docker_engines` + "`" + ` is also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Docker engine ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Docker engine name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcebuild_project_os",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note** This data source only supports 'public' site.

~> **Note:** This data source is a beta release. Some features may change in the future.

This data source is useful for look up the list of Sourcebuild os in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `OS available at Sourcebuild. ### OS Reference ` + "`" + `os` + "`" + ` is also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `OS ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `OS name.`,
				},
				resource.Attribute{
					Name:        "archi",
					Description: `OS architecture.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `OS version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "os",
					Description: `OS available at Sourcebuild. ### OS Reference ` + "`" + `os` + "`" + ` is also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `OS ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `OS name.`,
				},
				resource.Attribute{
					Name:        "archi",
					Description: `OS architecture.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `OS version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcebuild_project_os_runtime_versions",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note** This data source only supports 'public' site.

~> **Note:** This data source is a beta release. Some features may change in the future.

This data source is useful for look up the list of Sourcebuild runtime version environment in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "os_id",
					Description: `(Required) OS ID which runtime belongs.`,
				},
				resource.Attribute{
					Name:        "runtime_id",
					Description: `(Required) Runtime ID which runtime version belongs.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "runtime_versions",
					Description: `Runtime versions available at Sourcebuild. ### Runtime Versions Reference ` + "`" + `runtime_versions` + "`" + ` is also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Runtime version ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Runtime version name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "runtime_versions",
					Description: `Runtime versions available at Sourcebuild. ### Runtime Versions Reference ` + "`" + `runtime_versions` + "`" + ` is also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Runtime version ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Runtime version name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcebuild_project_os_runtimes",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note** This data source only supports 'public' site.

~> **Note:** This data source is a beta release. Some features may change in the future.

This data source is useful for look up the list of Sourcebuild runtime environment in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "os_id",
					Description: `(Required) OS ID which runtime belongs.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "runtimes",
					Description: `Runtimes available at Sourcebuild. ### Runtime Reference ` + "`" + `runtimes` + "`" + ` are is exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Runtime ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Runtime name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "runtimes",
					Description: `Runtimes available at Sourcebuild. ### Runtime Reference ` + "`" + `runtimes` + "`" + ` are is exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Runtime ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Runtime name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcebuild_projects",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note** This data source only supports 'public' site.

~> **Note:** This data source is a beta release. Some features may change in the future.

This data source is useful for look up the list of Sourcebuild project in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_name",
					Description: `(Optional) Search by project name (project including string).`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "projects",
					Description: `Sourcebuild projects. ### Project Reference ` + "`" + `projects` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Sourcebuild project ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Sourcebuild project Name.`,
				},
				resource.Attribute{
					Name:        "action_name",
					Description: `Permission status for searching details.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Permission name for searching details. (` + "`" + `Allow` + "`" + ` or ` + "`" + `Deny` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "projects",
					Description: `Sourcebuild projects. ### Project Reference ` + "`" + `projects` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Sourcebuild project ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Sourcebuild project Name.`,
				},
				resource.Attribute{
					Name:        "action_name",
					Description: `Permission status for searching details.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Permission name for searching details. (` + "`" + `Allow` + "`" + ` or ` + "`" + `Deny` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcecommit_repositories",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note:** This data source only supports 'public' site.

~> **Note:** This data source is a beta release. Some features may change in the future.

This data source is useful for look up the list of Sourcecommit repository in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `The list of repositories. ### Repository Reference ` + "`" + `repositories` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Sourcecommit repository ID.`,
				},
				resource.Attribute{
					Name:        "repository_no",
					Description: `Sourcecommit repository ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Sourcecommit repository Name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "repositories",
					Description: `The list of repositories. ### Repository Reference ` + "`" + `repositories` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Sourcecommit repository ID.`,
				},
				resource.Attribute{
					Name:        "repository_no",
					Description: `Sourcecommit repository ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Sourcecommit repository Name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcecommit_repository",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note:** This data source only supports 'public' site.

~> **Note:** This data source is a beta release. Some features may change in the future.

This data source is  useful for getting detail of Sourcecommit repository which is already created, such as getting git address of Sourcecommit repository.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the specific Repository to retrieve. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Sourcecommit repository ID.`,
				},
				resource.Attribute{
					Name:        "repository_no",
					Description: `Sourcecommit repository ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of reposiory.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `The creator of repository.`,
				},
				resource.Attribute{
					Name:        "git_https_url",
					Description: `The https git address of repository.`,
				},
				resource.Attribute{
					Name:        "git_ssh_url",
					Description: `The ssh git address of repository.`,
				},
				resource.Attribute{
					Name:        "file_safer",
					Description: `whether to use the [File Safer](https://www.ncloud.com/product/security/fileSafer) service`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Sourcecommit repository ID.`,
				},
				resource.Attribute{
					Name:        "repository_no",
					Description: `Sourcecommit repository ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of reposiory.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `The creator of repository.`,
				},
				resource.Attribute{
					Name:        "git_https_url",
					Description: `The https git address of repository.`,
				},
				resource.Attribute{
					Name:        "git_ssh_url",
					Description: `The ssh git address of repository.`,
				},
				resource.Attribute{
					Name:        "file_safer",
					Description: `whether to use the [File Safer](https://www.ncloud.com/product/security/fileSafer) service`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcedeploy_project_stage",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note** This data source only supports 'public' site.

-> **Note:** This data source is a beta release. Some features may change in the future.

This resource is useful for look up the list of Sourcedeploy stage detail in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of Sourcedeploy project.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of Sourcedeploy stage. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of stage.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `The type of deploy target.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `The configuration of deploy target.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `server`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of server.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group_no",
					Description: `The ID of Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group_name",
					Description: `The name of Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "cluster_uuid",
					Description: `The uuid of Kubernetes Service Cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `The name of Kubernetes Service Cluster.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `The name of ObjectStorage bucket.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of stage.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `The type of deploy target.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `The configuration of deploy target.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `server`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of server.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group_no",
					Description: `The ID of Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group_name",
					Description: `The name of Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "cluster_uuid",
					Description: `The uuid of Kubernetes Service Cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `The name of Kubernetes Service Cluster.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `The name of ObjectStorage bucket.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcedeploy_project_stage_scenario",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note** This data source only supports 'public' site.

-> **Note:** This data source is a beta release. Some features may change in the future.

This resource is useful for look up the list of Sourcedeploy scenario in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of Sourcedeploy project.`,
				},
				resource.Attribute{
					Name:        "stage_id",
					Description: `(Required) The ID of Sourcedeploy stage.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of Sourcedeploy scenario. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of scenario.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Sourcedeploy project description.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `scenario config.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `Deployment strategy.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `Deployment file.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `File type.`,
				},
				resource.Attribute{
					Name:        "object_storage",
					Description: `Objectstorage config.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `The Name of ObjectStorage bucket.`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `ObjectStorage object .`,
				},
				resource.Attribute{
					Name:        "source_build",
					Description: `Sourcebuild config.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of SourceBiuld project.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of SourceBuild project.`,
				},
				resource.Attribute{
					Name:        "rollboack",
					Description: `Rollback on deployment failure.`,
				},
				resource.Attribute{
					Name:        "deploy_command",
					Description: `Commands to execute in deploy.`,
				},
				resource.Attribute{
					Name:        "pre_deploy",
					Description: `Commands before deploy.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Running Account.`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `Run Command.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Deploy file.`,
				},
				resource.Attribute{
					Name:        "source_path",
					Description: `Source file path.`,
				},
				resource.Attribute{
					Name:        "deploy_path",
					Description: `Deploy Path.`,
				},
				resource.Attribute{
					Name:        "post_deploy",
					Description: `Commands after deploy.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Running Account.`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `Run Command.`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `Loadbalancer target group for blue-green deployment.`,
				},
				resource.Attribute{
					Name:        "load_balancer_target_group_no",
					Description: `Loadbalancer Target Group no.`,
				},
				resource.Attribute{
					Name:        "load_balancer_target_group_name",
					Description: `The name of Loadbalancer Target Group.`,
				},
				resource.Attribute{
					Name:        "delete_server",
					Description: `Whether to delete Servers in the auto scaling group.`,
				},
				resource.Attribute{
					Name:        "manifest",
					Description: `Manifest file for Kubernetesservice deployment.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Repository type.`,
				},
				resource.Attribute{
					Name:        "repository_name",
					Description: `The name of repository.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `The name of repository branch.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `File path.`,
				},
				resource.Attribute{
					Name:        "canary_config",
					Description: `config when deploying Kubernetesservice canary.`,
				},
				resource.Attribute{
					Name:        "canary_count",
					Description: `Number of baseline and canary pod.`,
				},
				resource.Attribute{
					Name:        "analysis_type",
					Description: `Canary analysis method.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Maximum time of deployment/cancellation.`,
				},
				resource.Attribute{
					Name:        "prometheus",
					Description: `Prometheus Url.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Analysis environment.`,
				},
				resource.Attribute{
					Name:        "baseline",
					Description: `Analysis environment variable > baseline.`,
				},
				resource.Attribute{
					Name:        "canary",
					Description: `Analysis environment variable > canary.`,
				},
				resource.Attribute{
					Name:        "metrics",
					Description: `Metric.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Metric name.`,
				},
				resource.Attribute{
					Name:        "success_criteria",
					Description: `Success criteria.`,
				},
				resource.Attribute{
					Name:        "query_type",
					Description: `Query type.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `Metric.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `Filter.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `Query.`,
				},
				resource.Attribute{
					Name:        "analysis_config",
					Description: `Analysis config.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Analysis time.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `Analysis delay time.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Analysis cycle.`,
				},
				resource.Attribute{
					Name:        "step",
					Description: `Metric collection cycle.`,
				},
				resource.Attribute{
					Name:        "pass_score",
					Description: `Analysis success score.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Deploy file.`,
				},
				resource.Attribute{
					Name:        "source_path",
					Description: `Source file path.`,
				},
				resource.Attribute{
					Name:        "deploy_path",
					Description: `Deploy Path.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of scenario.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Sourcedeploy project description.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `scenario config.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `Deployment strategy.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `Deployment file.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `File type.`,
				},
				resource.Attribute{
					Name:        "object_storage",
					Description: `Objectstorage config.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `The Name of ObjectStorage bucket.`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `ObjectStorage object .`,
				},
				resource.Attribute{
					Name:        "source_build",
					Description: `Sourcebuild config.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of SourceBiuld project.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of SourceBuild project.`,
				},
				resource.Attribute{
					Name:        "rollboack",
					Description: `Rollback on deployment failure.`,
				},
				resource.Attribute{
					Name:        "deploy_command",
					Description: `Commands to execute in deploy.`,
				},
				resource.Attribute{
					Name:        "pre_deploy",
					Description: `Commands before deploy.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Running Account.`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `Run Command.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Deploy file.`,
				},
				resource.Attribute{
					Name:        "source_path",
					Description: `Source file path.`,
				},
				resource.Attribute{
					Name:        "deploy_path",
					Description: `Deploy Path.`,
				},
				resource.Attribute{
					Name:        "post_deploy",
					Description: `Commands after deploy.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Running Account.`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `Run Command.`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `Loadbalancer target group for blue-green deployment.`,
				},
				resource.Attribute{
					Name:        "load_balancer_target_group_no",
					Description: `Loadbalancer Target Group no.`,
				},
				resource.Attribute{
					Name:        "load_balancer_target_group_name",
					Description: `The name of Loadbalancer Target Group.`,
				},
				resource.Attribute{
					Name:        "delete_server",
					Description: `Whether to delete Servers in the auto scaling group.`,
				},
				resource.Attribute{
					Name:        "manifest",
					Description: `Manifest file for Kubernetesservice deployment.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Repository type.`,
				},
				resource.Attribute{
					Name:        "repository_name",
					Description: `The name of repository.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `The name of repository branch.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `File path.`,
				},
				resource.Attribute{
					Name:        "canary_config",
					Description: `config when deploying Kubernetesservice canary.`,
				},
				resource.Attribute{
					Name:        "canary_count",
					Description: `Number of baseline and canary pod.`,
				},
				resource.Attribute{
					Name:        "analysis_type",
					Description: `Canary analysis method.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Maximum time of deployment/cancellation.`,
				},
				resource.Attribute{
					Name:        "prometheus",
					Description: `Prometheus Url.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Analysis environment.`,
				},
				resource.Attribute{
					Name:        "baseline",
					Description: `Analysis environment variable > baseline.`,
				},
				resource.Attribute{
					Name:        "canary",
					Description: `Analysis environment variable > canary.`,
				},
				resource.Attribute{
					Name:        "metrics",
					Description: `Metric.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Metric name.`,
				},
				resource.Attribute{
					Name:        "success_criteria",
					Description: `Success criteria.`,
				},
				resource.Attribute{
					Name:        "query_type",
					Description: `Query type.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `Metric.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `Filter.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `Query.`,
				},
				resource.Attribute{
					Name:        "analysis_config",
					Description: `Analysis config.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Analysis time.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `Analysis delay time.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Analysis cycle.`,
				},
				resource.Attribute{
					Name:        "step",
					Description: `Metric collection cycle.`,
				},
				resource.Attribute{
					Name:        "pass_score",
					Description: `Analysis success score.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Deploy file.`,
				},
				resource.Attribute{
					Name:        "source_path",
					Description: `Source file path.`,
				},
				resource.Attribute{
					Name:        "deploy_path",
					Description: `Deploy Path.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcedeploy_project_stage_scenarios",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note** This data source only supports 'public' site.

-> **Note:** This data source is a beta release. Some features may change in the future.

This resource is useful for look up the list of Sourcedeploy scenario in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of Sourcedeploy project.`,
				},
				resource.Attribute{
					Name:        "stage_id",
					Description: `(Required) The ID of Sourcedeploy stage.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "scenarios",
					Description: `The list of Sourcedeploy scenario. ### Scenarios Reference ` + "`" + `scenarios` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Sourcedeploy scenario.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Sourcedeploy scenario.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "scenarios",
					Description: `The list of Sourcedeploy scenario. ### Scenarios Reference ` + "`" + `scenarios` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Sourcedeploy scenario.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Sourcedeploy scenario.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcedeploy_project_stages",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note** This data source only supports 'public' site.

-> **Note:** This data source is a beta release. Some features may change in the future.

This resource is useful for look up the list of Sourcedeploy stage in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of Sourcedeploy project.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "stages",
					Description: `The list of Sourcedeploy stage. ### Stage Reference ` + "`" + `stages` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Sourcedeploy stage.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Sourcedeploy stage.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "stages",
					Description: `The list of Sourcedeploy stage. ### Stage Reference ` + "`" + `stages` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Sourcedeploy stage.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Sourcedeploy stage.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcedeploy_projects",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note** This data source only supports 'public' site.

-> **Note:** This data source is a beta release. Some features may change in the future.

This data source is useful for look up the list of Sourcedeploy project in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "projects",
					Description: `The list of Sourcedeploy project. ### Project Reference ` + "`" + `projects` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Sourcedeploy project.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Sourcedeploy project.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "projects",
					Description: `The list of Sourcedeploy project. ### Project Reference ` + "`" + `projects` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Sourcedeploy project.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Sourcedeploy project.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcepipeline_project",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note** This data source only supports 'public' site.

~> **Note:** This data source is a beta release. Some features may change in the future.

This module can be useful for getting detail of Sourcepipeline project created before.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcepipeline_projects",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

~> **Note** This data source only supports 'public' site.

~> **Note:** This data source is a beta release. Some features may change in the future.

This data source is useful for look up the list of Sourcepipeline projects in the region.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_sourcepipeline_trigger_timezone",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This data source is useful for look up the list of Sourcepipeline trigger time zone.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_subnet",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of Subnet created before. for example, determine the CIDR block of that Subnet

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific Subnet to retrieve.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Optional) The ID of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional) The CIDR block of Subnet to retrieve.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Available zone where the subnet will be placed physically.`,
				},
				resource.Attribute{
					Name:        "network_acl_no",
					Description: `(Optional) The ID of Network ACL.`,
				},
				resource.Attribute{
					Name:        "subnet_type",
					Description: `(Optional) Internet connectivity. If you use ` + "`" + `PUBLIC` + "`" + `, all VMs created within Subnet will be assigned a certified IP by default and will be able to communicate directly over the Internet. Considering the characteristics of Subnet, you can choose Subnet for the purpose of use. Accepted values: ` + "`" + `PUBLIC` + "`" + ` (Public) | ` + "`" + `PRIVATE` + "`" + ` (Private).`,
				},
				resource.Attribute{
					Name:        "usage_type",
					Description: `(Optional) Usage type, Accepted values: ` + "`" + `GEN` + "`" + ` (General) | ` + "`" + `LOADB` + "`" + ` (For load balancer) | ` + "`" + `NATGW` + "`" + ` (for NAT Gateway. Only pub env)..`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `The ID of Subnet. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Subnet.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subnet_no",
					Description: `The ID of Subnet. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Subnet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_subnets",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This resource is useful for look up the list of Subnet in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "subnet_no",
					Description: `(Optional) The ID of the specific Subnet to retrieve.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Optional) The ID of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional) The CIDR block of subnet to retrieve.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Available zone where the subnet will be placed physically.`,
				},
				resource.Attribute{
					Name:        "network_acl_no",
					Description: `(Optional) The ID of Network ACL.`,
				},
				resource.Attribute{
					Name:        "subnet_type",
					Description: `(Optional) Internet connectivity. If you use ` + "`" + `PUBLIC` + "`" + `, all VMs created within Subnet will be assigned a certified IP by default and will be able to communicate directly over the Internet. Considering the characteristics of Subnet, you can choose Subnet for the purpose of use. Accepted values: ` + "`" + `PUBLIC` + "`" + ` (Public) | ` + "`" + `PRIVATE` + "`" + ` (Private).`,
				},
				resource.Attribute{
					Name:        "usage_type",
					Description: `(Optional) Usage type, Accepted values: ` + "`" + `GEN` + "`" + ` (General) | ` + "`" + `LOADB` + "`" + ` (For load balancer) | ` + "`" + `NATGW` + "`" + ` (for NAT Gateway. Only pub env).`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `The list of Subnet ### Subnet Reference ` + "`" + `subnets` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `The ID of Subnet. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of subnet.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `The CIDR block of subnet.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Available zone where the Subnet is placed.`,
				},
				resource.Attribute{
					Name:        "network_acl_no",
					Description: `The ID of Network ACL.`,
				},
				resource.Attribute{
					Name:        "subnet_type",
					Description: `Internet connectivity.`,
				},
				resource.Attribute{
					Name:        "usage_type",
					Description: `Usage type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subnets",
					Description: `The list of Subnet ### Subnet Reference ` + "`" + `subnets` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `The ID of Subnet. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of subnet.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `The CIDR block of subnet.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Available zone where the Subnet is placed.`,
				},
				resource.Attribute{
					Name:        "network_acl_no",
					Description: `The ID of Network ACL.`,
				},
				resource.Attribute{
					Name:        "subnet_type",
					Description: `Internet connectivity.`,
				},
				resource.Attribute{
					Name:        "usage_type",
					Description: `Usage type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_vpc",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of VPC created before, such as determining the CIDR block of that VPC.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) name of the specific VPC to retrieve`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of VPC. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "ipv4_cidr_block",
					Description: `The CIDR block for the association.`,
				},
				resource.Attribute{
					Name:        "default_network_acl_no",
					Description: `The ID of the network ACL created by default on VPC creation.`,
				},
				resource.Attribute{
					Name:        "default_access_control_group_no",
					Description: `The ID of the ACG created by default on VPC creation.`,
				},
				resource.Attribute{
					Name:        "default_public_route_table_no",
					Description: `The ID of the Public Route Table created by default on VPC creation.`,
				},
				resource.Attribute{
					Name:        "default_private_route_table_no",
					Description: `The ID of the Private Route Table created by default on VPC creation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of VPC. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "ipv4_cidr_block",
					Description: `The CIDR block for the association.`,
				},
				resource.Attribute{
					Name:        "default_network_acl_no",
					Description: `The ID of the network ACL created by default on VPC creation.`,
				},
				resource.Attribute{
					Name:        "default_access_control_group_no",
					Description: `The ID of the ACG created by default on VPC creation.`,
				},
				resource.Attribute{
					Name:        "default_public_route_table_no",
					Description: `The ID of the Public Route Table created by default on VPC creation.`,
				},
				resource.Attribute{
					Name:        "default_private_route_table_no",
					Description: `The ID of the Private Route Table created by default on VPC creation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_vpc_peering",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This module can be useful for getting detail of VPC peering created before.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the specific VPC peering.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the specific VPC Peering to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpc_peering_no",
					Description: `The ID of VPC peering. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "source_vpc_no",
					Description: `The ID of VPC to which the request.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of VPC Peering`,
				},
				resource.Attribute{
					Name:        "has_reverse_vpc_peering",
					Description: `Is existing Reverse VPC Peering.`,
				},
				resource.Attribute{
					Name:        "is_between_accounts",
					Description: `VPC Peering Between Accounts.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_peering_no",
					Description: `The ID of VPC peering. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "source_vpc_no",
					Description: `The ID of VPC to which the request.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of VPC Peering`,
				},
				resource.Attribute{
					Name:        "has_reverse_vpc_peering",
					Description: `Is existing Reverse VPC Peering.`,
				},
				resource.Attribute{
					Name:        "is_between_accounts",
					Description: `VPC Peering Between Accounts.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_vpcs",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

This resource is useful for look up the list of VPC in the region.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) name of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the field to filter by.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) is ` + "`" + `values` + "`" + ` treated as a regular expression. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpcs",
					Description: `The list of vpcs ### VPC Reference ` + "`" + `vpcs` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of VPC. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of VPC.`,
				},
				resource.Attribute{
					Name:        "ipv4_cidr_block",
					Description: `The CIDR block for the association.`,
				},
				resource.Attribute{
					Name:        "default_network_acl_no",
					Description: `The ID of the network ACL created by default on VPC creation.`,
				},
				resource.Attribute{
					Name:        "default_access_control_group_no",
					Description: `The ID of the ACG created by default on VPC creation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpcs",
					Description: `The list of vpcs ### VPC Reference ` + "`" + `vpcs` + "`" + ` are also exported with the following attributes, where relevant: Each element supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of VPC. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of VPC.`,
				},
				resource.Attribute{
					Name:        "ipv4_cidr_block",
					Description: `The CIDR block for the association.`,
				},
				resource.Attribute{
					Name:        "default_network_acl_no",
					Description: `The ID of the network ACL created by default on VPC creation.`,
				},
				resource.Attribute{
					Name:        "default_access_control_group_no",
					Description: `The ID of the ACG created by default on VPC creation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_zones",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Gets a list of available zones.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: KR region.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save data source after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A List of region`,
				},
				resource.Attribute{
					Name:        "zone_no",
					Description: `Zone number`,
				},
				resource.Attribute{
					Name:        "zone_code",
					Description: `Zone code`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Zone name`,
				},
				resource.Attribute{
					Name:        "zone_description",
					Description: `Zone description`,
				},
				resource.Attribute{
					Name:        "region_no",
					Description: `Region number`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zones",
					Description: `A List of region`,
				},
				resource.Attribute{
					Name:        "zone_no",
					Description: `Zone number`,
				},
				resource.Attribute{
					Name:        "zone_code",
					Description: `Zone code`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Zone name`,
				},
				resource.Attribute{
					Name:        "zone_description",
					Description: `Zone description`,
				},
				resource.Attribute{
					Name:        "region_no",
					Description: `Region number`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ncloud_access_control_group":                    0,
		"ncloud_access_control_groups":                   1,
		"ncloud_access_control_rule":                     2,
		"ncloud_access_control_rules":                    3,
		"ncloud_auto_scaling_group":                      4,
		"ncloud_auto_scaling_policy":                     5,
		"ncloud_auto_scaling_schedule":                   6,
		"ncloud_block_storage":                           7,
		"ncloud_block_storage_snapshot":                  8,
		"ncloud_cdss_cluster":                            9,
		"ncloud_cdss_config_group":                       10,
		"ncloud_cdss_kafka_version":                      11,
		"ncloud_cdss_kafka_versions":                     12,
		"ncloud_cdss_node_product":                       13,
		"ncloud_cdss_node_products":                      14,
		"ncloud_cdss_os_image":                           15,
		"ncloud_cdss_os_images":                          16,
		"ncloud_init_script":                             17,
		"ncloud_launch_configuration":                    18,
		"ncloud_lb":                                      19,
		"ncloud_lb_listener":                             20,
		"ncloud_lb_target_group":                         21,
		"ncloud_member_server_image":                     22,
		"ncloud_member_server_images":                    23,
		"ncloud_nas_volume":                              24,
		"ncloud_nas_volumes":                             25,
		"ncloud_nat_gateway":                             26,
		"ncloud_network_acl_deny_allow_groups":           27,
		"ncloud_network_acls":                            28,
		"ncloud_network_interface":                       29,
		"ncloud_nks_cluster":                             30,
		"ncloud_nks_clusters":                            31,
		"ncloud_nks_kube_config":                         32,
		"ncloud_nks_node_pool":                           33,
		"ncloud_nks_node_pools":                          34,
		"ncloud_nks_versions":                            35,
		"ncloud_placement_group":                         36,
		"ncloud_port_forwarding_rule":                    37,
		"ncloud_port_forwarding_rules":                   38,
		"ncloud_public_ip":                               39,
		"ncloud_regions":                                 40,
		"ncloud_root_password":                           41,
		"ncloud_route_table":                             42,
		"ncloud_route_tables":                            43,
		"ncloud_server":                                  44,
		"ncloud_server_image":                            45,
		"ncloud_server_images":                           46,
		"ncloud_server_product":                          47,
		"ncloud_server_products":                         48,
		"ncloud_servers":                                 49,
		"ncloud_ses_cluster":                             50,
		"ncloud_ses_clusters":                            51,
		"ncloud_ses_node_os_images":                      52,
		"ncloud_ses_node_products":                       53,
		"ncloud_ses_versions":                            54,
		"ncloud_sourcebuild_project":                     55,
		"ncloud_sourcebuild_project_computes":            56,
		"ncloud_sourcebuild_project_docker_engines":      57,
		"ncloud_sourcebuild_project_os":                  58,
		"ncloud_sourcebuild_project_os_runtime_versions": 59,
		"ncloud_sourcebuild_project_os_runtimes":         60,
		"ncloud_sourcebuild_projects":                    61,
		"ncloud_sourcecommit_repositories":               62,
		"ncloud_sourcecommit_repository":                 63,
		"ncloud_sourcedeploy_project_stage":              64,
		"ncloud_sourcedeploy_project_stage_scenario":     65,
		"ncloud_sourcedeploy_project_stage_scenarios":    66,
		"ncloud_sourcedeploy_project_stages":             67,
		"ncloud_sourcedeploy_projects":                   68,
		"ncloud_sourcepipeline_project":                  69,
		"ncloud_sourcepipeline_projects":                 70,
		"ncloud_sourcepipeline_trigger_timezone":         71,
		"ncloud_subnet":                                  72,
		"ncloud_subnets":                                 73,
		"ncloud_vpc":                                     74,
		"ncloud_vpc_peering":                             75,
		"ncloud_vpcs":                                    76,
		"ncloud_zones":                                   77,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
