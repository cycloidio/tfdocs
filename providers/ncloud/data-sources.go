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
					Name:        "description",
					Description: `(Optional) description to create`,
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
					Description: `Instance No.`,
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
					Description: `Instance No.`,
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
					Description: `(Optional) Usage type, Accepted values: ` + "`" + `GEN` + "`" + ` (General) | ` + "`" + `LOADB` + "`" + ` (For load balancer).`,
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
					Description: `(Optional) Usage type, Accepted values: ` + "`" + `GEN` + "`" + ` (General) | ` + "`" + `LOADB` + "`" + ` (For load balancer).`,
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

		"ncloud_access_control_group":          0,
		"ncloud_access_control_groups":         1,
		"ncloud_access_control_rule":           2,
		"ncloud_access_control_rules":          3,
		"ncloud_auto_scaling_group":            4,
		"ncloud_auto_scaling_policy":           5,
		"ncloud_auto_scaling_schedule":         6,
		"ncloud_block_storage":                 7,
		"ncloud_block_storage_snapshot":        8,
		"ncloud_init_script":                   9,
		"ncloud_launch_configuration":          10,
		"ncloud_lb":                            11,
		"ncloud_lb_listener":                   12,
		"ncloud_lb_target_group":               13,
		"ncloud_member_server_image":           14,
		"ncloud_member_server_images":          15,
		"ncloud_nas_volume":                    16,
		"ncloud_nas_volumes":                   17,
		"ncloud_nat_gateway":                   18,
		"ncloud_network_acl_deny_allow_groups": 19,
		"ncloud_network_acls":                  20,
		"ncloud_network_interface":             21,
		"ncloud_nks_cluster":                   22,
		"ncloud_nks_clusters":                  23,
		"ncloud_nks_kube_config":               24,
		"ncloud_nks_node_pool":                 25,
		"ncloud_nks_node_pools":                26,
		"ncloud_nks_versions":                  27,
		"ncloud_placement_group":               28,
		"ncloud_port_forwarding_rule":          29,
		"ncloud_port_forwarding_rules":         30,
		"ncloud_public_ip":                     31,
		"ncloud_regions":                       32,
		"ncloud_root_password":                 33,
		"ncloud_route_table":                   34,
		"ncloud_route_tables":                  35,
		"ncloud_server":                        36,
		"ncloud_server_image":                  37,
		"ncloud_server_images":                 38,
		"ncloud_server_product":                39,
		"ncloud_server_products":               40,
		"ncloud_subnet":                        41,
		"ncloud_subnets":                       42,
		"ncloud_vpc":                           43,
		"ncloud_vpc_peering":                   44,
		"ncloud_vpcs":                          45,
		"ncloud_zones":                         46,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
