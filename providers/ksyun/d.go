package ksyun

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ksyun_availability_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available zones in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "availability_zone_name",
					Description: `Name of the zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone_name",
					Description: `Name of the zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_ebs_volumes",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of volume resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of volume IDs.Without this parameter, query all volumes' information.`,
				},
				resource.Attribute{
					Name:        "volume_category",
					Description: `(Optional) Volume classification, system disk "system" or data disk "data".`,
				},
				resource.Attribute{
					Name:        "volume_status",
					Description: `(Optional) The status of volumes, “creating|available|attaching|in-use|detaching|extending|deleting|error|recycling”.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) The type of volumes. "SSD" or "SATA".`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_eips",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of EIP resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Elastic IP IDs, all the EIPs belong to this region will be retrieved if the ID is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) One or more project IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Elastic IPs that satisfy the condition.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "eips",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Elastic IPs that satisfy the condition.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_healthchecks",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of healthcheck resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of health check IDs, all the healthcheck belong to this region will be retrieved if the ID is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Optional) A list of listener IDs, all the healthcheck belong to this region will be retrieved if the ID is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "health_checks",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Elastic IPs that satisfy the condition.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "health_checks",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Elastic IPs that satisfy the condition.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_images",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available image resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of image IDs`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting images by name. (Such as: ` + "`" + `^CentOS 7.[1-2] 64` + "`" + ` means CentOS 7.1 of 64-bit operating system or CentOS 7.2 of 64-bit operating system, "^Ubuntu 16.04 64" means Ubuntu 16.04 of 64-bit operating system).`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional, type: bool) If more than one result are returned, select the most recent one. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of image.`,
				},
				resource.Attribute{
					Name:        "image_source",
					Description: `Valid values are import, copy, share, extend, system.`,
				},
				resource.Attribute{
					Name:        "image_state",
					Description: `Status of the image.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `If ksyun provide the image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Display name of the image.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `Platform type of the image system.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `Progress of image creation.`,
				},
				resource.Attribute{
					Name:        "sys_disk",
					Description: `Size of the created disk.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_date",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of image.`,
				},
				resource.Attribute{
					Name:        "image_source",
					Description: `Valid values are import, copy, share, extend, system.`,
				},
				resource.Attribute{
					Name:        "image_state",
					Description: `Status of the image.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `If ksyun provide the image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Display name of the image.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `Platform type of the image system.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `Progress of image creation.`,
				},
				resource.Attribute{
					Name:        "sys_disk",
					Description: `Size of the created disk.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of instance resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by instance name.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The image ID of some instance used.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of subnet. the instance will use the subnet in the current region.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) Security Group to associate with.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `instances documented below. The attribute (` + "`" + `instances` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of instance.`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `The state of instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC linked to the instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of subnet linked to the instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID for the image to use for the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of instance.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security Group to associate with.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The name of instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project instance belongs to.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The user data to be specified into this instance.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Instance charge type.`,
				},
				resource.Attribute{
					Name:        "availability_zone_name",
					Description: `Name of the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `Instance private IP address.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of project_id the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The size of systemdisk.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The type of systemdisk.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `instances documented below. The attribute (` + "`" + `instances` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of instance.`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `The state of instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC linked to the instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of subnet linked to the instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID for the image to use for the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of instance.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security Group to associate with.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The name of instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project instance belongs to.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The user data to be specified into this instance.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Instance charge type.`,
				},
				resource.Attribute{
					Name:        "availability_zone_name",
					Description: `Name of the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `Instance private IP address.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of project_id the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The size of systemdisk.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The type of systemdisk.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_krds",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of krds resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "output_file",
					Description: `(Required) will return the file name of the content store`,
				},
				resource.Attribute{
					Name:        "db_instance_identifier",
					Description: `(Optional) instance ID (passed in the instance ID to get the details of the instance, otherwise get the list)`,
				},
				resource.Attribute{
					Name:        "db_instance_type",
					Description: `(Optional) hrds (highly available), RR (read-only), trds (temporary)`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) case sensitive, value range: default (default sorting method), group (sorting by replication group, will rank read-only instances after their primary instances)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) the default value is all projects`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_krds_security_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of krds_security_groups resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_lb_acls",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Load Balancer Rule resources belong to the Load Balancer listener.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of LB Rule IDs, all the LB Rules belong to the Load Balancer listener will be retrieved if the ID is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "lb_acls",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of LB Rules that satisfy the condition.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "lb_acls",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of LB Rules that satisfy the condition.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_lb_listeners",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Load Balancer Listener resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) The ID of a load balancer.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of LB Listener IDs, all the LB Listeners belong to this region will be retrieved if the ID is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of LB listeners that satisfy the condition. The attribute (` + "`" + `lb_listeners` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of LB Listener.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `The name of LB Listener.`,
				},
				resource.Attribute{
					Name:        "listener_protocol",
					Description: `LB Listener protocol. Possible values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` , ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + ` .`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `Port opened on the LB Listener to receive requests, range: 1-65535.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `The load balancer method in which the listener is.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listeners",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of LB listeners that satisfy the condition. The attribute (` + "`" + `lb_listeners` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of LB Listener.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `The name of LB Listener.`,
				},
				resource.Attribute{
					Name:        "listener_protocol",
					Description: `LB Listener protocol. Possible values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` , ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + ` .`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `Port opened on the LB Listener to receive requests, range: 1-65535.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `The load balancer method in which the listener is.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_lines",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of line resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "lines",
					Description: `All the lines accourding the argument.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of lines that satisfy the condition.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "lines",
					Description: `All the lines accourding the argument.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of lines that satisfy the condition.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_listener_servers",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Load Balancer Listener Server resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of LB Listener Server IDs, all the LB Listener Servers belong to this region will be retrieved if the ID is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of LB Listener.`,
				},
				resource.Attribute{
					Name:        "real_server_ip",
					Description: `The name of LB Listener Server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of LB Listener.`,
				},
				resource.Attribute{
					Name:        "real_server_ip",
					Description: `The name of LB Listener Server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_mongodbs",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of MongoDB resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of MongoDB, all the MongoDBs belong to this region will be retrieved if the name is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) The id of MongoDB, all the MongoDBs belong to this region will be retrieved if the instance_id is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iam_project_id",
					Description: `(Optional) The project instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Used to retrieve instances belong to specified VPC .`,
				},
				resource.Attribute{
					Name:        "vnet_id",
					Description: `(Optional) The ID of subnet. the instance will use the subnet in the current region.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `(Optional) The vip of instances.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of MongoDBs that satisfy the condition.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instances",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of MongoDBs that satisfy the condition.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_network_interface",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Network Interface in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Network Interface IDs, all the Network Interfaces belong to this region will be retrieved if the ID is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "network_interfaces",
					Description: `It is a nested type which documented below. The attribute (` + "`" + `network_interfaces` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the network interface.`,
				},
				resource.Attribute{
					Name:        "securitygroup_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of network interface.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of instance.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The private IP address assigned to the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_interfaces",
					Description: `It is a nested type which documented below. The attribute (` + "`" + `network_interfaces` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the network interface.`,
				},
				resource.Attribute{
					Name:        "securitygroup_id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of network interface.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of instance.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The private IP address assigned to the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_redis_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Redis resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of redis instance, all the Redis instances belong to this region will be retrieved if the name is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iam_project_id",
					Description: `(Optional) The project instance belongs to.`,
				},
				resource.Attribute{
					Name:        "cache_id",
					Description: `(Optional) The ID of the intance .`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Used to retrieve instances belong to specified VPC .`,
				},
				resource.Attribute{
					Name:        "vnet_id",
					Description: `(Optional) The ID of subnet. the instance will use the subnet in the current region.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `(Optional) Private IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Redis instances that satisfy the condition.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instances",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Redis instances that satisfy the condition.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_security_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Security Group resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Security Group IDs, all the Security Group resources belong to this region will be retrieved if the ID is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Security Group resources that satisfy the condition. The attribute (` + "`" + `security_groups` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Security Group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of Security Group.`,
				},
				resource.Attribute{
					Name:        "security_group_entry",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "security_group_type",
					Description: `The type of Security Group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of vpc .`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for the security group, formatted in RFC3339 time string. The attribute (` + "`" + `security_group_entry` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The cidr block of source.`,
				},
				resource.Attribute{
					Name:        "port_range_from",
					Description: `The start of port numbers .`,
				},
				resource.Attribute{
					Name:        "port_range_to",
					Description: `The end of port numbers.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol. Can be ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `ip` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "security_groups",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Security Group resources that satisfy the condition. The attribute (` + "`" + `security_groups` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Security Group.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `The name of Security Group.`,
				},
				resource.Attribute{
					Name:        "security_group_entry",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "security_group_type",
					Description: `The type of Security Group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of vpc .`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for the security group, formatted in RFC3339 time string. The attribute (` + "`" + `security_group_entry` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The cidr block of source.`,
				},
				resource.Attribute{
					Name:        "port_range_from",
					Description: `The start of port numbers .`,
				},
				resource.Attribute{
					Name:        "port_range_to",
					Description: `The end of port numbers.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol. Can be ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `ip` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_slbs",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Load Balancer resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Load Balancer IDs, all the LBs belong to this region will be retrieved if the ID is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting lbs by name.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The ID of the VPC linked to the Load Balancers.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of subnet that intrant load balancer belongs to.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "slbs",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Load Balancers that satisfy the condition. The attribute (` + "`" + `slbs` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Load Balancer.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Load Balancer.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC linked to the Load Balancers.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of subnet that intrant load balancer belongs to.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The IP address of intranet IP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "slbs",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Load Balancers that satisfy the condition. The attribute (` + "`" + `slbs` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Load Balancer.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Load Balancer.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC linked to the Load Balancers.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of subnet that intrant load balancer belongs to.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The IP address of intranet IP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_sqlservers",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of sqlservers resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_subnets",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Subnet resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Subnet IDs, all the Subnet resources belong to this region will be retrieved if the ID is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The id of the VPC that the desired Subnet belongs to.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Subnet resources that satisfy the condition. The attribute (` + "`" + `subnets` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `The name of Subnet.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The cidr block of the desired Subnet.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation of Subnet, formatted in RFC3339 time string.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subnets",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Subnet resources that satisfy the condition. The attribute (` + "`" + `subnets` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `The name of Subnet.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The cidr block of the desired Subnet.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation of Subnet, formatted in RFC3339 time string.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_vpcs",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VPC resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of VPC IDs, all the VPC resources belong to this region will be retrieved if the ID is ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpcs",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of VPC resources that satisfy the condition. The attribute (` + "`" + `vpcs` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `The name of VPC.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR blocks of VPC.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for VPC, formatted in RFC3339 time string.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpcs",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of VPC resources that satisfy the condition. The attribute (` + "`" + `vpcs` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `The name of VPC.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR blocks of VPC.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for VPC, formatted in RFC3339 time string.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ksyun_availability_zones":   0,
		"ksyun_ebs_volumes":          1,
		"ksyun_eips":                 2,
		"ksyun_healthchecks":         3,
		"ksyun_images":               4,
		"ksyun_instances":            5,
		"ksyun_krds":                 6,
		"ksyun_krds_security_groups": 7,
		"ksyun_lb_acls":              8,
		"ksyun_lb_listeners":         9,
		"ksyun_lines":                10,
		"ksyun_listener_servers":     11,
		"ksyun_mongodbs":             12,
		"ksyun_network_interface":    13,
		"ksyun_redis_instances":      14,
		"ksyun_security_groups":      15,
		"ksyun_slbs":                 16,
		"ksyun_sqlservers":           17,
		"ksyun_subnets":              18,
		"ksyun_vpcs":                 19,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
