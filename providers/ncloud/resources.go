package ncloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ncloud_access_control_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides an ACG(Access Control Group) resource.

~> **NOTE:** This resource only supports VPC environment.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Required) The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name to create. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Indicates whether to get default group only. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of ACG(Access Control Group)`,
				},
				resource.Attribute{
					Name:        "access_control_group_no",
					Description: `The ID of ACG(Access Control Group) (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether is default or not by VPC creation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of ACG(Access Control Group)`,
				},
				resource.Attribute{
					Name:        "access_control_group_no",
					Description: `The ID of ACG(Access Control Group) (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether is default or not by VPC creation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_access_control_group_rule",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides an rule of ACG(Access Control Group) resource.

~> **NOTE:** This resource only supports VPC environment.

~> **NOTE:** To performance, we recommend using one resource per ACG(Access Control Group). If you use multiple resources in a single ACG(Access Control Group), then can cause a slowdown.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_control_group_no",
					Description: `(Required) The ID of the ACG.`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `(Optional) Specifies an Inbound(ingress) rules. Parameters defined below. This argument is processed in [attriutbe-as-blocks](https://www.terraform.io/docs/configuration/attr-as-blocks.html) mode.`,
				},
				resource.Attribute{
					Name:        "outbound",
					Description: `(Optional) Specifies an Outbound(egress) rules. Parameters defined below. This argument is processed in [attriutbe-as-blocks](https://www.terraform.io/docs/configuration/attr-as-blocks.html) mode. ### Access Control Group Rule Reference Both ` + "`" + `inbound` + "`" + ` and ` + "`" + `outbound` + "`" + ` support following attributes:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Select between TCP, UDP, and ICMP. Accepted values: ` + "`" + `TCP` + "`" + ` | ` + "`" + `UDP` + "`" + ` | ` + "`" + `ICMP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ip_block",
					Description: `(Optional) The CIDR block to match. This must be a valid network mask. Cannot be specified with ` + "`" + `source_access_control_group_no` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_access_control_group_no",
					Description: `(Optional) The ID of specific ACG to apply this rule to. Cannot be specified with ` + "`" + `ip_block` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `(Optional) Range of ports to apply. You can enter from ` + "`" + `1` + "`" + ` to ` + "`" + `65535` + "`" + `. e.g. set single port: ` + "`" + `22` + "`" + ` or set range port : ` + "`" + `8000-9000` + "`" + ` ~>`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description to create. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of ACG(Access Control Group) rule`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of ACG(Access Control Group) rule`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_auto_scaling_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a ncloud auto scaling group resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Auto Scaling Group name to create. default : Ncloud assigns default values.`,
				},
				resource.Attribute{
					Name:        "launch_configuration_no",
					Description: `(Required) Launch Configuration Number for creating Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Optional) The number of servers is adjusted according to the desired capacity value. This value must be more than the minimum capacity and less than the maximum capacity. If this value is not specified, initially create a server with a minimum capacity. valid from ` + "`" + `0` + "`" + ` to ` + "`" + `30` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Required) The minimum size of the Auto Scaling Group. valid from ` + "`" + `0` + "`" + ` to ` + "`" + `30` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required) The maximum size of the Auto Scaling Group. valid from ` + "`" + `0` + "`" + ` to ` + "`" + `30` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_cooldown",
					Description: `(Optional) The cooldown time is the period set to ignore even if the monitoring event alarm occurs after the actual scaling is being performed or is completed.`,
				},
				resource.Attribute{
					Name:        "health_check_type_code",
					Description: `(Optional) ` + "`" + `SVR` + "`" + ` or ` + "`" + `LOADB` + "`" + `. Controls how health checking is done.`,
				},
				resource.Attribute{
					Name:        "wait_for_capacity_timeout",
					Description: `(Optional) The maximum amount of time Terraform should wait for an ASG instance to become healthy. Setting this to "0" causes Terraform to skip all Capacity Waiting behavior.`,
				},
				resource.Attribute{
					Name:        "health_check_grace_period",
					Description: `(Optional) Set the time to hold health check after the server instance is put into the service with the health check hold period. ~>`,
				},
				resource.Attribute{
					Name:        "zone_no_list",
					Description: `(Required) the list of zone numbers where server instances belonging to this group will exist. ~>`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `(Required) The ID of the associated Subnet.`,
				},
				resource.Attribute{
					Name:        "access_control_group_no_list",
					Description: `(Required) The ID of the ACG.`,
				},
				resource.Attribute{
					Name:        "target_group_list",
					Description: `(Optional) - Target Group number list of Load Balancer. ~>`,
				},
				resource.Attribute{
					Name:        "server_name_prefix",
					Description: `(Optional) Create name beginning with the specified prefix. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group_no",
					Description: `The ID of Auto Scaling Group (It is the same result as id)`,
				},
				resource.Attribute{
					Name:        "server_instance_no_list",
					Description: `List of server instances belonging to Auto Scaling Group. ~>`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group_no",
					Description: `The ID of Auto Scaling Group (It is the same result as id)`,
				},
				resource.Attribute{
					Name:        "server_instance_no_list",
					Description: `List of server instances belonging to Auto Scaling Group. ~>`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_auto_scaling_policy",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a ncloud auto scaling policy resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Auto Scaling Policy name to create.`,
				},
				resource.Attribute{
					Name:        "adjustment_type_code",
					Description: `(Required) Determines how the number of servers is scaled when the scaling policy is performed. Valid values are ` + "`" + `CHANG` + "`" + `, ` + "`" + `EXACT` + "`" + `, and ` + "`" + `PRCNT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scaling_adjustment",
					Description: `(Required) Specify the adjustment value for the adjustment type. Enter a negative value to decrease when adjustTypeCode is ` + "`" + `CHANG` + "`" + ` or ` + "`" + `PRCNT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `(Optional) The cooldown time is the period set to ignore even if the monitoring event alarm occurs after the actual scaling is being performed or is completed.`,
				},
				resource.Attribute{
					Name:        "min_adjustment_step",
					Description: `(Optional) Change the number of server instances by the minimum adjustment width.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group_no",
					Description: `(Required) The number of the auto scaling group. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Auto Scaling Policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The ID of Auto Scaling Policy (It is the same result as id).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Auto Scaling Policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The ID of Auto Scaling Policy (It is the same result as id).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_auto_scaling_schedule",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a ncloud auto scaling schedule resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Auto Scaling Schedule name to create.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Required) The number of servers is adjusted according to the desired capacity value.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Required) The minimum size of the Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required) The maximum size of the Auto Scaling Group.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) You can determine the date and time when the schedule first starts. If you don't enter ` + "`" + `recurrence` + "`" + `, be sure to enter startTime. It cannot be duplicated with the startTime of another schedule and must be later than the current time, before endTime. Format : ` + "`" + `yyyy-MM-ddTHH:mm:ssZ` + "`" + ` format in UTC/KST only (for example, 2021-02-02T15:00:00+0900).`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) You can determine the date and time when the schedule end. If you don't enter ` + "`" + `recurrence` + "`" + `, be sure to enter startTime. It must be a time later than the current time and a time later than the startTime. Format : ` + "`" + `yyyy-MM-ddTHH:mm:ssZ` + "`" + ` format in UTC/KST only (for example, 2021-02-02T18:00:00+0900).`,
				},
				resource.Attribute{
					Name:        "recurrence",
					Description: `(Optional) Repeat Settings. You can specify a recurring schedule in crontab format.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_group_no",
					Description: `(Required) The number of the auto scaling group. ~>`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Optional) the time band for the repeat settings. Valid values are ` + "`" + `KST` + "`" + ` and ` + "`" + `UTC` + "`" + `. Default : ` + "`" + `KST` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Auto Scaling Schedule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The ID of Auto Scaling Schedule (It is the same result as id).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Auto Scaling Schedule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The ID of Auto Scaling Schedule (It is the same result as id).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_block_storage",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Block Storage resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the block storage to create. It is automatically set when you take a snapshot.`,
				},
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `(Required) Server instance ID to which you want to assign the block storage.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name to create. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description to create.`,
				},
				resource.Attribute{
					Name:        "disk_detail_type",
					Description: `(Optional) Type of block storage disk detail to create. Default ` + "`" + `SSD` + "`" + `. Accepted values: ` + "`" + `SSD` + "`" + ` | ` + "`" + `HDD` + "`" + ` ~>`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The availability zone in which the block storage instance will be created.`,
				},
				resource.Attribute{
					Name:        "snapshot_no",
					Description: `(Optional) Create the block storage from the snapshots you take. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Block storage instance.`,
				},
				resource.Attribute{
					Name:        "block_storage_no",
					Description: `The ID of Block storage instance. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Server name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Block storage type code.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `Device name.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `Block storage product code.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Block storage instance status code.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Disk type code.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Block storage instance.`,
				},
				resource.Attribute{
					Name:        "block_storage_no",
					Description: `The ID of Block storage instance. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Server name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Block storage type code.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `Device name.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `Block storage product code.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Block storage instance status code.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Disk type code.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_block_storage_snapshot",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a ncloud block storage snapshot resource.

~> **NOTE:** This resource only supports Classic environment.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "block_storage_instance_no",
					Description: `(Required) Block storage instance Number for creating snapshot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Block storage snapshot name to create. default : Ncloud assigns default values.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Descriptions on a snapshot to create. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `Block Storage Snapshot Instance Number`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `Block Storage Snapshot Volume Size`,
				},
				resource.Attribute{
					Name:        "original_block_storage_instance_no",
					Description: `Original Block Storage Instance Number`,
				},
				resource.Attribute{
					Name:        "original_block_storage_name",
					Description: `Original Block Storage Name`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Block Storage Snapshot Instance Status code`,
				},
				resource.Attribute{
					Name:        "instance_status_name",
					Description: `Block Storage Snapshot Instance Status Name`,
				},
				resource.Attribute{
					Name:        "instance_operation",
					Description: `Block Storage Snapshot Instance Operation code`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the block storage snapshot instance`,
				},
				resource.Attribute{
					Name:        "server_image_product_code",
					Description: `Server Image Product Code`,
				},
				resource.Attribute{
					Name:        "os_information",
					Description: `OS Information`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_no",
					Description: `Block Storage Snapshot Instance Number`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `Block Storage Snapshot Volume Size`,
				},
				resource.Attribute{
					Name:        "original_block_storage_instance_no",
					Description: `Original Block Storage Instance Number`,
				},
				resource.Attribute{
					Name:        "original_block_storage_name",
					Description: `Original Block Storage Name`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Block Storage Snapshot Instance Status code`,
				},
				resource.Attribute{
					Name:        "instance_status_name",
					Description: `Block Storage Snapshot Instance Status Name`,
				},
				resource.Attribute{
					Name:        "instance_operation",
					Description: `Block Storage Snapshot Instance Operation code`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the block storage snapshot instance`,
				},
				resource.Attribute{
					Name:        "server_image_product_code",
					Description: `Server Image Product Code`,
				},
				resource.Attribute{
					Name:        "os_information",
					Description: `OS Information`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_init_script",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides an Init script resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "content",
					Description: `(Required) Initialization script content. Scripts such as Python, Perl, Shell are available for Linux environments. However, on the first line, you must specify the script path you want to run in the form of ` + "`" + `#!/usr/bin/env` + "`" + ` python, ` + "`" + `#!/bin/perl` + "`" + `, ` + "`" + `#!/bin/bash` + "`" + `, etc. Windows environments can only write Visual Basic scripts.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name to create. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description to create.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional) Type of O/S to apply server instance. Default ` + "`" + `LNX` + "`" + `. Accepted values: ` + "`" + `LNX` + "`" + ` (LINUX) | ` + "`" + `WND` + "`" + ` (WINDOWS) ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Init script.`,
				},
				resource.Attribute{
					Name:        "init_script_no",
					Description: `The ID of the Init script. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Init script.`,
				},
				resource.Attribute{
					Name:        "init_script_no",
					Description: `The ID of the Init script. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_launch_configuration",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a ncloud launch configuration resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Launch Configuration name to create. default : Ncloud assigns default values.`,
				},
				resource.Attribute{
					Name:        "server_image_product_code",
					Description: `(Optional) Server image product code to determine which server image to create. It can be obtained through data ncloud_server_images. You are required to select one between two parameters: server image product code (server_image_product_code) and member server image number member_server_image_no)`,
				},
				resource.Attribute{
					Name:        "server_product_code",
					Description: `(Optional) Server product code to determine the server specification to create. It can be obtained through the getServerProductList action. Default : Selected as minimum specification. The minimum standards are 1. memory 2. CPU 3. basic block storage size 4. disk type (NET,LOCAL)`,
				},
				resource.Attribute{
					Name:        "member_server_image_no",
					Description: `(Optional) Required value when creating a server from a manually created server image. It can be obtained through the getMemberServerImageList action.`,
				},
				resource.Attribute{
					Name:        "login_key_name",
					Description: `(Optional) The login key name to encrypt with the public key. Default : Uses the login key name most recently created.`,
				},
				resource.Attribute{
					Name:        "init_script_no",
					Description: `(Optional) Set init script ID, The server can run a user-set initialization script at first boot. ~>`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The server will execute the user data script set by the user at first boot. To view the column, it is returned only when viewing the server instance.`,
				},
				resource.Attribute{
					Name:        "access_control_group_no_list",
					Description: `(Optional) You can set the ACG created when creating the server. ACG setting number can be obtained through the getAccessControlGroupList action. Default : Default ACG number. ~>`,
				},
				resource.Attribute{
					Name:        "is_encrypted_volume",
					Description: `(Optional) you can set whether to encrypt basic block storage if server image is RHV. Default false. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Launch Configuration.`,
				},
				resource.Attribute{
					Name:        "launch_configuration_no",
					Description: `The ID of Launch Configuration (It is the same result as id)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Launch Configuration.`,
				},
				resource.Attribute{
					Name:        "launch_configuration_no",
					Description: `The ID of Launch Configuration (It is the same result as id)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_lb",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Load Balancer resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the load balancer.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Optional) The network type of load balancer to create. Accepted values: ` + "`" + `PUBLIC` + "`" + ` | ` + "`" + `PRIVATE` + "`" + `. Default: ` + "`" + `PUBLIC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional) The time in seconds that the idle timeout. Valid only if the load balancer type is not ` + "`" + `NETWORK` + "`" + `. Default: 60.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of load balancer to create. Accepted values: ` + "`" + `APPLICATION` + "`" + ` | ` + "`" + `NETWORK` + "`" + ` | ` + "`" + `NETWORK_PROXY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "throughput_type",
					Description: `(Optional) The performance type code of load balancer. Accepted values: ` + "`" + `SMALL` + "`" + ` | ` + "`" + `MEDIUM` + "`" + ` | ` + "`" + `LARGE` + "`" + `. If the load balancer type is ` + "`" + `NETWORK` + "`" + ` and the load balancer network type is ` + "`" + `PRIVATE` + "`" + `, only ` + "`" + `SMALL` + "`" + ` can be selected. Default: ` + "`" + `SMALL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_no_list",
					Description: `(Required) A list of IDs in the associated Subnets. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_no",
					Description: `The ID of load balancer (It is the same result as id).`,
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
					Name:        "id",
					Description: `The ID of load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_no",
					Description: `The ID of load balancer (It is the same result as id).`,
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
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Load Balancer Listener resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_no",
					Description: `(Required) The ID of the load balancer.`,
				},
				resource.Attribute{
					Name:        "target_group_no",
					Description: `(Required) The ID of the target group.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port on which the load balancer is listening.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol type for the listener. The types of protocols available are limited by the type of load balancer. ` + "`" + `APPLICATION` + "`" + ` Load Balancer Accepted values: ` + "`" + `HTTP` + "`" + ` | ` + "`" + `HTTPS` + "`" + `, ` + "`" + `NETWORK` + "`" + ` Load Balancer Accepted values : ` + "`" + `TCP` + "`" + `, ` + "`" + `NETWORK_PROXY` + "`" + ` Load Balancer Accepted values : ` + "`" + `TCP` + "`" + ` | ` + "`" + `TLS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tls_min_version_type",
					Description: `(Optional) The TLS minimum supported version type code. Valid only if the listener protocol type is ` + "`" + `HTTPS` + "`" + ` or ` + "`" + `TLS` + "`" + `. Accepted values : ` + "`" + `TLSV10` + "`" + `(TLSv1.0) | ` + "`" + `TLSV11` + "`" + `(TLSv1.1) | ` + "`" + `TLSV12` + "`" + `(TLSv1.2). Default: ` + "`" + `TLSV10` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_http2",
					Description: `(Optional) Whether to use HTTP/2 protocol. Valid only if the listener protocol type is ` + "`" + `HTTPS` + "`" + `. Accepted values : ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl_certificate_no",
					Description: `(Optional) The ID of the SSL certificate. If the listener protocol type is ` + "`" + `HTTPS` + "`" + ` or ` + "`" + `TLS` + "`" + `, an SSL certificate must be set. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of listener.`,
				},
				resource.Attribute{
					Name:        "listener_no",
					Description: `The ID of listener (It is the same result as id).`,
				},
				resource.Attribute{
					Name:        "rule_no_list",
					Description: `The list of listener rule number.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of listener.`,
				},
				resource.Attribute{
					Name:        "listener_no",
					Description: `The ID of listener (It is the same result as id).`,
				},
				resource.Attribute{
					Name:        "rule_no_list",
					Description: `The list of listener rule number.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_lb_target_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Target Group resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the target group.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port on which targets receive traffic. Default: 80.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol to use for routing traffic to the targets. Accepted values: ` + "`" + `TCP` + "`" + ` | ` + "`" + `PROXY_TCP` + "`" + ` | ` + "`" + `HTTP` + "`" + ` | ` + "`" + `HTTPS` + "`" + `. The protocol you use determines which type of load balancer is applicable. ` + "`" + `APPLICATION` + "`" + ` Load Balancer Accepted values: ` + "`" + `HTTP` + "`" + ` | ` + "`" + `HTTPS` + "`" + `, ` + "`" + `NETWORK` + "`" + ` Load Balancer Accepted values : ` + "`" + `TCP` + "`" + `, ` + "`" + `NETWORK_PROXY` + "`" + ` Load Balancer Accepted values : ` + "`" + `PROXY_TCP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the target group.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Optional) The health check to check the health of the target.`,
				},
				resource.Attribute{
					Name:        "cycle",
					Description: `(Optional) The number of health check cycle.`,
				},
				resource.Attribute{
					Name:        "down_threshold",
					Description: `(Optional) The number of health check failure threshold. You can determine the number of consecutive health check failures that are required before a health check is considered a failed state.`,
				},
				resource.Attribute{
					Name:        "up_threshold",
					Description: `(Optional) The number of health check normal threshold. You can determine the number of consecutive health checks that are required before health checks are considered success state.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `(Optional) The HTTP method for the health check. You can determine which HTTP method to use for health checks. If the health check protocol type is ` + "`" + `HTTP` + "`" + ` or ` + "`" + `HTTPS` + "`" + `, be sure to enter it. Accepted values: ` + "`" + `HEAD` + "`" + ` | ` + "`" + `GET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port to use for health checks. Default: 80.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The type of protocol to use for health checks. If the target group protocol type is ` + "`" + `TCP` + "`" + ` or ` + "`" + `PROXY_TCP` + "`" + `, Heal Check Protocol is only valid for ` + "`" + `TCP` + "`" + `. If the target group protocol type is ` + "`" + `HTTP` + "`" + ` or ` + "`" + `HTTPS` + "`" + `, Heal Check Protocol is valid only for ` + "`" + `HTTP` + "`" + ` and ` + "`" + `HTTPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `(Optional) The URL path of the health check. Valid only if Health Check protocol type is ` + "`" + `HTTP` + "`" + ` or ` + "`" + `HTTPS` + "`" + `. URL path must begin with ` + "`" + `/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `(Optional) The type of target to be added to the target group.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Required) The ID of the VPC in to create the target group.`,
				},
				resource.Attribute{
					Name:        "use_sticky_session",
					Description: `(Optional) Whether to use session specific access.`,
				},
				resource.Attribute{
					Name:        "use_proxy_protocol",
					Description: `(Optional) Whether to use a proxy protocol. Valid only available if the target group type selected is ` + "`" + `TCP` + "`" + ` | ` + "`" + `HTTP` + "`" + ` | ` + "`" + `HTTPS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "algorithm_type",
					Description: `(Optional) The type of algorithm to use for load balancing. Accepted values: ` + "`" + `RR` + "`" + `(Round Robin) | ` + "`" + `SIPHS` + "`" + `(Source IP Hash) | ` + "`" + `LC` + "`" + `(Least Connection) | ` + "`" + `MH` + "`" + `(Maglev Hash). ` + "`" + `RR` + "`" + ` | ` + "`" + `SIPHS` + "`" + ` | ` + "`" + `LC` + "`" + ` are valid only if the target group type is ` + "`" + `PROXY_TCP` + "`" + `, ` + "`" + `HTTP` + "`" + ` or ` + "`" + `HTTPS` + "`" + `. ` + "`" + `MH` + "`" + ` | ` + "`" + `RR` + "`" + ` are valid only if the target group type is ` + "`" + `TCP` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of target group.`,
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
					Name:        "target_no_list",
					Description: `The list of target number to bind to the target group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of target group.`,
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
					Name:        "target_no_list",
					Description: `The list of target number to bind to the target group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_lb_target_group_attachment",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Target Group Attachment resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "target_group_no",
					Description: `(Required) The ID of target group.`,
				},
				resource.Attribute{
					Name:        "target_no_list",
					Description: `(Required) The List of server instance ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of target.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of target.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_load_balancer",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Provides a ncloud load balancer instance resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "rule_list",
					Description: `(Required) Load balancer rules.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(Required) Protocol type code of load balancer rules. The following codes are available. [HTTP | HTTPS | TCP | SSL]`,
				},
				resource.Attribute{
					Name:        "load_balancer_port",
					Description: `(Required) Load balancer port of load balancer rules`,
				},
				resource.Attribute{
					Name:        "server_port",
					Description: `(Required) Server port of load balancer rules`,
				},
				resource.Attribute{
					Name:        "l7_health_check_path",
					Description: `Health check path of load balancer rules. Required when the ` + "`" + `protocol_type` + "`" + ` is HTTP/HTTPS.`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `Load balancer SSL certificate name. Required when the ` + "`" + `protocol_type` + "`" + ` value is SSL/HTTPS.`,
				},
				resource.Attribute{
					Name:        "proxy_protocol_use_yn",
					Description: `(Optional) Use 'Y' if you want to check client IP addresses by enabling the proxy protocol while you select TCP or SSL.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of a load balancer instance. Default: Automatically specified by Ncloud.`,
				},
				resource.Attribute{
					Name:        "algorithm_type",
					Description: `(Optional) Load balancer algorithm type code. The available algorithms are as follows: [ROUND ROBIN (RR) | LEAST_CONNECTION (LC)]. Default: ROUND ROBIN (RR)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of a load balancer instance.`,
				},
				resource.Attribute{
					Name:        "server_instance_no_list",
					Description: `(Optional) List of server instance numbers to be bound to the load balancer`,
				},
				resource.Attribute{
					Name:        "network_usage_type",
					Description: `(Optional) Network usage identification code. PBLIP(PublicIP), PRVT(PrivateIP). default : PBLIP(PublicIP)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: KR region.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. Zone in which you want to create a NAS volume. Default: The first zone of the region. Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `Load balancer instance No`,
				},
				resource.Attribute{
					Name:        "virtual_ip",
					Description: `Virtual IP address`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the load balancer instance`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Domain name`,
				},
				resource.Attribute{
					Name:        "instance_status_name",
					Description: `Load balancer instance status name`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Load balancer instance status code`,
				},
				resource.Attribute{
					Name:        "instance_operation",
					Description: `Load balancer instance operation code`,
				},
				resource.Attribute{
					Name:        "is_http_keep_alive",
					Description: `Http keep alive value [true | false]`,
				},
				resource.Attribute{
					Name:        "connection_timeout",
					Description: `Connection timeout`,
				},
				resource.Attribute{
					Name:        "load_balanced_server_instance_list",
					Description: `Load balanced server instance list`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_no",
					Description: `Load balancer instance No`,
				},
				resource.Attribute{
					Name:        "virtual_ip",
					Description: `Virtual IP address`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the load balancer instance`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Domain name`,
				},
				resource.Attribute{
					Name:        "instance_status_name",
					Description: `Load balancer instance status name`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Load balancer instance status code`,
				},
				resource.Attribute{
					Name:        "instance_operation",
					Description: `Load balancer instance operation code`,
				},
				resource.Attribute{
					Name:        "is_http_keep_alive",
					Description: `Http keep alive value [true | false]`,
				},
				resource.Attribute{
					Name:        "connection_timeout",
					Description: `Connection timeout`,
				},
				resource.Attribute{
					Name:        "load_balanced_server_instance_list",
					Description: `Load balanced server instance list`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_load_balancer_ssl_certificate",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a ncloud load balancer ssl certificate resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "certificate_name",
					Description: `(Required) Name of a certificate`,
				},
				resource.Attribute{
					Name:        "privatekey",
					Description: `(Required) Private key for a certificate`,
				},
				resource.Attribute{
					Name:        "publickey_certificate",
					Description: `(Required) Public key for a certificate`,
				},
				resource.Attribute{
					Name:        "certificate_chain",
					Description: `(Optional) Chainca certificate (Required if the certificate is issued with a chainca)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_login_key",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Login key resource.

~> **Note:** All arguments including the private key will be stored in the raw state as plain-text.
[Read more about sensitive data in state](/docs/state/sensitive-data.html).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required) Key name to generate. If the generated key name exists, an error occurs. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `Generated private key`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of the login key`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the login key`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "private_key",
					Description: `Generated private key`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of the login key`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the login key`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_nas_volume",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a NAS volume.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "volume_name_postfix",
					Description: `(Required) Name of a NAS volume to create. Enter a volume name that is 3-20 characters in length after entering the name for user identification.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Required) Enter the nas volume size to be created. You can enter in GiB.`,
				},
				resource.Attribute{
					Name:        "volume_allotment_protocol_type",
					Description: `(Required) Volume allotment protocol type code. ` + "`" + `NFS` + "`" + ` | ` + "`" + `CIFS` + "`" + ` ` + "`" + `NFS` + "`" + `: You can mount the volume in a Linux server such as CentOS and Ubuntu. ` + "`" + `CIFS` + "`" + `: You can mount the volume in a Windows server.`,
				},
				resource.Attribute{
					Name:        "server_instance_no_list",
					Description: `(Optional) List of server instance numbers for which access to NFS is to be controlled.`,
				},
				resource.Attribute{
					Name:        "cifs_user_name",
					Description: `(Optional) CIFS user name. The ID must contain a combination of English alphabet and numbers, which can be 6-20 characters in length.`,
				},
				resource.Attribute{
					Name:        "cifs_user_password",
					Description: `(Optional) CIFS user password. The password must contain a combination of at least 2 English letters, numbers and special characters, which can be 8-14 characters in length.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) NAS volume description`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. Zone in which you want to create a NAS volume. Default: The first zone of the region. Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "custom_ip_list",
					Description: `(Optional) To add a server of another account to the NAS volume, enter a private IP address of the server. ~>`,
				},
				resource.Attribute{
					Name:        "is_encrypted_volume",
					Description: `(Optional) Volume encryption. Default ` + "`" + `false` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of NAS Volume.`,
				},
				resource.Attribute{
					Name:        "nas_volume_no",
					Description: `The ID of NAS Volume. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `NAS volume name.`,
				},
				resource.Attribute{
					Name:        "volume_total_size",
					Description: `Volume total size, in GiB`,
				},
				resource.Attribute{
					Name:        "snapshot_volume_size",
					Description: `Snapshot volume size, in GiB`,
				},
				resource.Attribute{
					Name:        "is_snapshot_configuration",
					Description: `Indicates whether a snapshot volume is set.`,
				},
				resource.Attribute{
					Name:        "is_event_configuration",
					Description: `Indicates whether the event is set.`,
				},
				resource.Attribute{
					Name:        "mount_information",
					Description: `Mount information for NAS volume.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of NAS Volume.`,
				},
				resource.Attribute{
					Name:        "nas_volume_no",
					Description: `The ID of NAS Volume. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `NAS volume name.`,
				},
				resource.Attribute{
					Name:        "volume_total_size",
					Description: `Volume total size, in GiB`,
				},
				resource.Attribute{
					Name:        "snapshot_volume_size",
					Description: `Snapshot volume size, in GiB`,
				},
				resource.Attribute{
					Name:        "is_snapshot_configuration",
					Description: `Indicates whether a snapshot volume is set.`,
				},
				resource.Attribute{
					Name:        "is_event_configuration",
					Description: `Indicates whether the event is set.`,
				},
				resource.Attribute{
					Name:        "mount_information",
					Description: `Mount information for NAS volume.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_nat_gateway",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a NAT gateway resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Required) The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) Available zone where the subnet will be placed physically.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name to create. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description to create. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_no",
					Description: `The ID of the NAT Gateway. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP on created NAT Gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NAT Gateway.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_no",
					Description: `The ID of the NAT Gateway. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP on created NAT Gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_network_acl",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a rule of Network ACL resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Required) The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name to create. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description to create ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Network ACL.`,
				},
				resource.Attribute{
					Name:        "network_acl_no",
					Description: `The ID of the Network ACL. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether is default or not by VPC creation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Network ACL.`,
				},
				resource.Attribute{
					Name:        "network_acl_no",
					Description: `The ID of the Network ACL. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether is default or not by VPC creation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_network_acl_deny_allow_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a rule of Network ACL Deny-Allow Group resource. You can manage list of IP using this resource, \
Network ACL Deny-Allow Group can be added to the Network ACL Rule(` + "`" + `ncloud_network_acl_rule` + "`" + `) using ` + "`" + `deny_allow_group_no` + "`" + `.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Required) The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `(Required) Enter the IP addresses as list to be registered in the Deny-Allow Group. Up to 100 IPs can be registered. Duplicate IP addresses are not allowed.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name to create. If omitted, terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description to create ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Deny-Allow Group.`,
				},
				resource.Attribute{
					Name:        "network_acl_deny_allow_group_no",
					Description: `The ID of the Deny-Allow Group. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Deny-Allow Group.`,
				},
				resource.Attribute{
					Name:        "network_acl_deny_allow_group_no",
					Description: `The ID of the Deny-Allow Group. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_network_acl_rule",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a rule of Network ACL  resource.

~> **NOTE:** To performance, we recommend using one resource per Network ACL. If you use multiple resources in a single Network ACL, then can cause a slowdown.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_acl_no",
					Description: `(Required) The ID of the Network ACL.`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `(Optional) Specifies an Inbound(ingress) rules. Parameters defined below. This argument is processed in [attriutbe-as-blocks](https://www.terraform.io/docs/configuration/attr-as-blocks.html) mode.`,
				},
				resource.Attribute{
					Name:        "outbound",
					Description: `(Optional) Specifies an Outbound(egress) rules. Parameters defined below. This argument is processed in [attriutbe-as-blocks](https://www.terraform.io/docs/configuration/attr-as-blocks.html) mode. ### Network ACL Rule Reference Both ` + "`" + `inbound` + "`" + ` and ` + "`" + `outbound` + "`" + ` support following attributes:`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) Priority for rules, Used for ordering. Can be an integer from ` + "`" + `1` + "`" + ` to ` + "`" + `199` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Select between TCP, UDP, and ICMP. Accepted values: ` + "`" + `TCP` + "`" + ` | ` + "`" + `UDP` + "`" + ` | ` + "`" + `ICMP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "rule_action",
					Description: `(Required) The action to take. Accepted values: ` + "`" + `ALLOW` + "`" + ` | ` + "`" + `DROP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ip_block",
					Description: `(Optional, Required if ` + "`" + `deny_allow_group_no` + "`" + ` is not provided) The CIDR block to match. This must be a valid network mask.`,
				},
				resource.Attribute{
					Name:        "deny_allow_group_no",
					Description: `(Optional, Required if ` + "`" + `ip_block` + "`" + ` is not provided) The access source Deny-Allow Group number of network ACL rules. You can specify a Deny-Allow Group instead of an IP address block as the access source. ` + "`" + `deny_allow_group_no` + "`" + ` can be obtained through the Data source ` + "`" + `ncloud_network_acl_deny_allow_group` + "`" + ``,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `(Optional) Range of ports to apply. You can enter from ` + "`" + `1` + "`" + ` to ` + "`" + `65535` + "`" + `. e.g. set single port: ` + "`" + `22` + "`" + ` or set range port : ` + "`" + `8000-9000` + "`" + ` ~>`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description to create.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_network_interface",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Network Interface resource.

~> **NOTE:** This resource only supports VPC environment.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "subnet_no",
					Description: `(Required) The ID of the associated Subnet.`,
				},
				resource.Attribute{
					Name:        "access_control_groups",
					Description: `(Required) List of ACG ID to apply to network interfaces. A maximum of three ACGs can be applied.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name to create. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description to create.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) Set the IP addresses that you want to assign to the network interface. Must be in the IP address range of the subnet where the network interface is created. The last ` + "`" + `0` + "`" + ` to ` + "`" + `5' IP address of the Subnet is not available and duplicate IP addresses are not available at the Subnet scope.`,
				},
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `(Optional) The ID of server instance to assign network interface. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Network Interface.`,
				},
				resource.Attribute{
					Name:        "network_interface_no",
					Description: `The ID of Network Interface. (It is the same result as ` + "`" + `id` + "`" + `)`,
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
					Description: `Whether is default or not by Server instance creation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Network Interface.`,
				},
				resource.Attribute{
					Name:        "network_interface_no",
					Description: `The ID of Network Interface. (It is the same result as ` + "`" + `id` + "`" + `)`,
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
					Description: `Whether is default or not by Server instance creation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_nks_cluster",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Kubernetes Service cluster resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cluster name.`,
				},
				resource.Attribute{
					Name:        "login_key_name",
					Description: `(Required) Login key name.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) zone Code.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Required) VPC No.`,
				},
				resource.Attribute{
					Name:        "subnet_no_list",
					Description: `(Required) Subnet No. list. (For now, ` + "`" + `public subnet` + "`" + ` is not supported. Will be supported soon)`,
				},
				resource.Attribute{
					Name:        "lb_private_subnet_no",
					Description: `(Required) Subnet No. for private loadbalancer only.`,
				},
				resource.Attribute{
					Name:        "lb_public_subnet_no",
					Description: `(Optional) Subnet No. for public loadbalancer only. (Available only ` + "`" + `SGN` + "`" + ` region)`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "audit",
					Description: `(Required) Audit log availability. (` + "`" + `boolean` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `(Optional) Kubenretes version . ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster uuid.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Cluster uuid. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `Control Plane API address. ## Import Kubernetes Service Cluster can be imported using the name, e.g., $ terraform import ncloud_nks_cluster.my_cluster uuid`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Cluster uuid.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Cluster uuid. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `Control Plane API address. ## Import Kubernetes Service Cluster can be imported using the name, e.g., $ terraform import ncloud_nks_cluster.my_cluster uuid`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_nks_node_pool",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Kubernetes Service nodepool resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "node_pool_name",
					Description: `(Required) Nodepool name.`,
				},
				resource.Attribute{
					Name:        "cluster_uuid",
					Description: `(Required) Cluster uuid.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `(Required) Number of nodes.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `(Required) Product code.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Required) Auto scaling availability.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `(Required) Maximum number of nodes available for auto scaling.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `(Required) Minimum number of nodes available for auto scaling.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `(Optional) Subnet No. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of nodepool.` + "`" + `CusterUuid:NodePoolName` + "`" + ``,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `Instance No. ## Import NKS Node Pools can be imported using the cluster_name and node_pool_name separated by a colon (:), e.g., $ terraform import ncloud_nks_node_pool.my_node_pool uuid:my_node_pool`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of nodepool.` + "`" + `CusterUuid:NodePoolName` + "`" + ``,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `Instance No. ## Import NKS Node Pools can be imported using the cluster_name and node_pool_name separated by a colon (:), e.g., $ terraform import ncloud_nks_node_pool.my_node_pool uuid:my_node_pool`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_placement_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Placement group resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name to create. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "placement_group_type",
					Description: `(Optional) Type of placement group. Default ` + "`" + `AA` + "`" + `. Accepted values: ` + "`" + `AA` + "`" + ` ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Placement group.`,
				},
				resource.Attribute{
					Name:        "placement_group_no",
					Description: `The ID of the Placement group. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Placement group.`,
				},
				resource.Attribute{
					Name:        "placement_group_no",
					Description: `The ID of the Placement group. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_port_forwarding_rule",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `(Required) Server instance number for which port forwarding is set`,
				},
				resource.Attribute{
					Name:        "port_forwarding_external_port",
					Description: `(Required) External port for port forwarding`,
				},
				resource.Attribute{
					Name:        "port_forwarding_internal_port",
					Description: `(Required) Internal port for port forwarding. Only the following ports are available. [Linux: ` + "`" + `22` + "`" + ` | Windows: ` + "`" + `3389` + "`" + `]`,
				},
				resource.Attribute{
					Name:        "port_forwarding_configuration_no",
					Description: `(Optional) Port forwarding configuration number. You can get by calling ` + "`" + `data ncloud_port_forwarding_rules` + "`" + ` ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "port_forwarding_public_ip",
					Description: `Port forwarding Public IP`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Zone code`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "port_forwarding_public_ip",
					Description: `Port forwarding Public IP`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Zone code`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_public_ip",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Public IP instance resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `(Optional) Server instance number to assign after creating a public IP. You can get one by calling getPublicIpTargetServerInstanceList.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Public IP description. ~>`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. You can decide a zone where servers are created. You can decide in which zone the product list will be requested. default : Select the first Zone in the specific region Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_no",
					Description: `The ID of Public IP. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP Address.`,
				},
				resource.Attribute{
					Name:        "kind_type",
					Description: `Public IP kind type`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Public IP.`,
				},
				resource.Attribute{
					Name:        "public_ip_no",
					Description: `The ID of Public IP. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP Address.`,
				},
				resource.Attribute{
					Name:        "kind_type",
					Description: `Public IP kind type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_route",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Route resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table_no",
					Description: `(Required) The ID of the Route table.`,
				},
				resource.Attribute{
					Name:        "destination_cidr_block",
					Description: `(Required) Destination CIDR block, Set the destination IP address range for the route you want to add. (e.g. 0.0.0.0/0, 100.10.20.0/24)`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `(Required) Destination target type, Select the destination type of the route to add. Accepted values: ` + "`" + `NATGW` + "`" + ` (NAT Gateway) | ` + "`" + `VPCPEERING` + "`" + ` (VPC Peering) | ` + "`" + `VGW` + "`" + ` (Virtual Private Gateway).`,
				},
				resource.Attribute{
					Name:        "target_no",
					Description: `(Required) Set the destination identification number for the destination type.`,
				},
				resource.Attribute{
					Name:        "target_name",
					Description: `(Required) Set the destination name for the destination type. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether is default or not by Route table creation.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether is default or not by Route table creation.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the associated VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_route_table",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Route Table resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Required) The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "supported_subnet_type",
					Description: `(Required) Subnet type. Accepted values : ` + "`" + `PUBLIC` + "`" + ` (Public) | ` + "`" + `PRIVATE` + "`" + ` (Private).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name to create. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description to create. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Route table.`,
				},
				resource.Attribute{
					Name:        "route_table_no",
					Description: `The ID of the Route table. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether is default or not by VPC creation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Route table.`,
				},
				resource.Attribute{
					Name:        "route_table_no",
					Description: `The ID of the Route table. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether is default or not by VPC creation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_route_table_association",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provide a resource to create association between route table and a subnet.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table_no",
					Description: `(Required) The ID of the Route Table.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `(Required) The ID of Subnet to create association. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the route table association (` + "`" + `route_table_no` + "`" + `:` + "`" + `subnet_no` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the route table association (` + "`" + `route_table_no` + "`" + `:` + "`" + `subnet_no` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_server",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Server instance resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server_image_product_code",
					Description: `(Optional, Required if ` + "`" + `member_server_image_no` + "`" + ` is not provided) Server image product code to determine which server image to create. It can be obtained through ` + "`" + `data.ncloud_server_image(s)` + "`" + `. - [Docs server Image Products](https://github.com/NaverCloudPlatform/terraform-ncloud-docs/blob/main/docs/server_image_product.md) - [` + "`" + `ncloud_server_image` + "`" + ` data source](../data-sources/server_image.md) - [` + "`" + `ncloud_server_images` + "`" + ` data source](../data-sources/server_images.md)`,
				},
				resource.Attribute{
					Name:        "server_product_code",
					Description: `(Optional) Server product code to determine the server specification to create. It can be obtained through the ` + "`" + `data.ncloud_server_product(s)` + "`" + ` action. Default : Selected as minimum specification. The minimum standards are 1. memory 2. CPU 3. basic block storage size 4. disk type (NET,LOCAL) - [Docs server Image Products](https://github.com/NaverCloudPlatform/terraform-ncloud-docs/blob/main/docs/server_image_product.md) - [` + "`" + `ncloud_server_product` + "`" + ` data source](../data-sources/server_product.md) - [` + "`" + `ncloud_server_products` + "`" + ` data source](../data-sources/server_products.md)`,
				},
				resource.Attribute{
					Name:        "member_server_image_no",
					Description: `(Optional, Required if ` + "`" + `server_image_product_code` + "`" + ` is not provided) Required value when creating a server from a manually created server image. It can be obtained through the ` + "`" + `data.ncloud_member_server_image(s)` + "`" + ` action. - [` + "`" + `ncloud_member_server_image` + "`" + ` data source](../data-sources/member_server_image.md) - [` + "`" + `ncloud_member_server_images` + "`" + ` data source](../data-sources/member_server_images.md)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Server name to create. default: Assigned by ncloud`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Server description to create.`,
				},
				resource.Attribute{
					Name:        "login_key_name",
					Description: `(Optional) The login key name to encrypt with the public key. Default : Uses the login key name most recently created.`,
				},
				resource.Attribute{
					Name:        "is_protect_server_termination",
					Description: `(Optional) You can set whether or not to protect return when creating. default :false`,
				},
				resource.Attribute{
					Name:        "fee_system_type_code",
					Description: `(Optional) A rate system identification code. There are time plan(MTRAT) and flat rate (FXSUM). Default : Time plan(MTRAT)`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. You can determine the ZONE where the server will be created. Default : Assigned by NAVER Cloud Platform. Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "access_control_group_configuration_no_list",
					Description: `(Optional) You can set the ACG created when creating the server. ACG setting number can be obtained through the getAccessControlGroupList action. Default : Default ACG number`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The server will execute the user data script set by the user at first boot. To view the column, it is returned only when viewing the server instance.`,
				},
				resource.Attribute{
					Name:        "raid_type_name",
					Description: `(Optional) Raid Type Name.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `(Optional) Server instance tag list.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `(Required) Instance tag key`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `(Required) Instance tag value ~>`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `(Required) The ID of the associated Subnet.`,
				},
				resource.Attribute{
					Name:        "init_script_no",
					Description: `(Optional) Set init script ID, The server can run a user-set initialization script at first boot.`,
				},
				resource.Attribute{
					Name:        "placement_group_no",
					Description: `(Optional) Physical placement group that belongs to the server instance.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `(Optional) List of Network Interface. You can assign up to three network interfaces.`,
				},
				resource.Attribute{
					Name:        "network_interface_no",
					Description: `(Required) If you want to add a network interface that you created yourself, set the network interface ID.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Required) Sets the order of network interfaces to be assigned to the server to create. The unit name (eth0, eth1, etc.) is determined in that order. There must be one primary network interface. If you set ` + "`" + `0` + "`" + `, network interface is set by default. You can assign up to three network interfaces.`,
				},
				resource.Attribute{
					Name:        "is_encrypted_base_block_storage_volume",
					Description: `(Optional) you can set whether to encrypt basic block storage if server image is RHV. Default ` + "`" + `false` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of server instance.`,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `The ID of server instance.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `number of CPUs.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `The size of the memory in bytes.`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `The size of base block storage in bytes.`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `Platform type code.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP.`,
				},
				resource.Attribute{
					Name:        "server_image_name",
					Description: `Server image name.`,
				},
				resource.Attribute{
					Name:        "port_forwarding_public_ip",
					Description: `Port forwarding public ip.`,
				},
				resource.Attribute{
					Name:        "port_forwarding_external_port",
					Description: `Port forwarding external port.`,
				},
				resource.Attribute{
					Name:        "port_forwarding_internal_port",
					Description: `Port forwarding internal port.`,
				},
				resource.Attribute{
					Name:        "base_block_storage_disk_type",
					Description: `Base block storage disk type code.`,
				},
				resource.Attribute{
					Name:        "base_block_storage_disk_detail_type",
					Description: `Base block storage disk detail type code. ~>`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the VPC where you want to place the Server Instance.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `List of Network Interface.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet ID of the network interface.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `IP address of the network interface.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of server instance.`,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `The ID of server instance.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `number of CPUs.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `The size of the memory in bytes.`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `The size of base block storage in bytes.`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `Platform type code.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP.`,
				},
				resource.Attribute{
					Name:        "server_image_name",
					Description: `Server image name.`,
				},
				resource.Attribute{
					Name:        "port_forwarding_public_ip",
					Description: `Port forwarding public ip.`,
				},
				resource.Attribute{
					Name:        "port_forwarding_external_port",
					Description: `Port forwarding external port.`,
				},
				resource.Attribute{
					Name:        "port_forwarding_internal_port",
					Description: `Port forwarding internal port.`,
				},
				resource.Attribute{
					Name:        "base_block_storage_disk_type",
					Description: `Base block storage disk type code.`,
				},
				resource.Attribute{
					Name:        "base_block_storage_disk_detail_type",
					Description: `Base block storage disk detail type code. ~>`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of the VPC where you want to place the Server Instance.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `List of Network Interface.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `Subnet ID of the network interface.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `IP address of the network interface.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_subnet",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Subnet resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_no",
					Description: `(Required) The ID of the VPC where you want to place the Subnet.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) assign some subnet address ranges within the range of VPC addresses, must be between /16 and/28 within the private band (10.0.0/8,172.16.0.0/12,192.168.0.0/16).`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) Available zone where the subnet will be placed physically.`,
				},
				resource.Attribute{
					Name:        "network_acl_no",
					Description: `(Required) The ID of Network ACL.`,
				},
				resource.Attribute{
					Name:        "subnet_type",
					Description: `(Required) Internet connectivity. If you use ` + "`" + `PUBLIC` + "`" + ` all VMs created within Subnet will be assigned a certified IP by default and will be able to communicate directly over the Internet. Considering the characteristics of Subnet, you can choose Subnet for the purpose of use. Accepted values: ` + "`" + `PUBLIC` + "`" + ` (Public) | ` + "`" + `PRIVATE` + "`" + ` (Private).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name to create. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "usage_type",
					Description: `(Optional) Usage type, Default ` + "`" + `GEN` + "`" + `. Accepted values: ` + "`" + `GEN` + "`" + ` (General) | ` + "`" + `LOADB` + "`" + ` (For load balancer). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `The ID of the Subnet. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_no",
					Description: `The ID of the Subnet. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_vpc",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a VPC resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name to create. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "ipv4_cidr_block",
					Description: `(Required) The CIDR block of the VPC. The range must be between /16 and/28 within the private band (10.0.0/8,172.16.0.0/12,192.168.0.0/16). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "id",
					Description: `The ID of VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_no",
					Description: `The ID of VPC. (It is the same result as ` + "`" + `id` + "`" + `)`,
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
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a VPC Peering resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_vpc_no",
					Description: `(Required) The ID of VPC from which the request is sent.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name to create. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) description to create. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of VPC peering.`,
				},
				resource.Attribute{
					Name:        "vpc_peering_no",
					Description: `The ID of VPC peering. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "has_reverse_vpc_peering",
					Description: `Reverse VPC Peering exists.`,
				},
				resource.Attribute{
					Name:        "is_between_accounts",
					Description: `VPC Peering Between Accounts.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of VPC peering.`,
				},
				resource.Attribute{
					Name:        "vpc_peering_no",
					Description: `The ID of VPC peering. (It is the same result as ` + "`" + `id` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "has_reverse_vpc_peering",
					Description: `Reverse VPC Peering exists.`,
				},
				resource.Attribute{
					Name:        "is_between_accounts",
					Description: `VPC Peering Between Accounts.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"ncloud_access_control_group":          0,
		"ncloud_access_control_group_rule":     1,
		"ncloud_auto_scaling_group":            2,
		"ncloud_auto_scaling_policy":           3,
		"ncloud_auto_scaling_schedule":         4,
		"ncloud_block_storage":                 5,
		"ncloud_block_storage_snapshot":        6,
		"ncloud_init_script":                   7,
		"ncloud_launch_configuration":          8,
		"ncloud_lb":                            9,
		"ncloud_lb_listener":                   10,
		"ncloud_lb_target_group":               11,
		"ncloud_lb_target_group_attachment":    12,
		"ncloud_load_balancer":                 13,
		"ncloud_load_balancer_ssl_certificate": 14,
		"ncloud_login_key":                     15,
		"ncloud_nas_volume":                    16,
		"ncloud_nat_gateway":                   17,
		"ncloud_network_acl":                   18,
		"ncloud_network_acl_deny_allow_group":  19,
		"ncloud_network_acl_rule":              20,
		"ncloud_network_interface":             21,
		"ncloud_nks_cluster":                   22,
		"ncloud_nks_node_pool":                 23,
		"ncloud_placement_group":               24,
		"ncloud_port_forwarding_rule":          25,
		"ncloud_public_ip":                     26,
		"ncloud_route":                         27,
		"ncloud_route_table":                   28,
		"ncloud_route_table_association":       29,
		"ncloud_server":                        30,
		"ncloud_subnet":                        31,
		"ncloud_vpc":                           32,
		"ncloud_vpc_peering":                   33,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
