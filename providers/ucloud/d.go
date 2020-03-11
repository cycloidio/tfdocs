package ucloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ucloud_db_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of database instance resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability zone where database instances are located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of database instance IDs, all the database instances belong to this region will be retrieved if the ID is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting database instances by name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "db_instances",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of database instances that satisfy the condition. - - - The attribute (` + "`" + `db_instances` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone where database instance is located.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of database instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of database instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Specifies the type of database instance.`,
				},
				resource.Attribute{
					Name:        "standby_zone",
					Description: `Availability zone where the standby database instance is located for the high availability database instance with multiple zone.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC linked to the database instances.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of subnet linked to the database instances.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The type of database instance engine.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `The database instance engine version.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port on which the database instance accepts connections.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address assigned to the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_storage",
					Description: `Specifies the allocated storage size in gigabytes (GB).`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of db instance.`,
				},
				resource.Attribute{
					Name:        "backup_count",
					Description: `Specifies the number of backup saved per week.`,
				},
				resource.Attribute{
					Name:        "backup_begin_time",
					Description: `Specifies when the backup starts, measured in hour.`,
				},
				resource.Attribute{
					Name:        "backup_date",
					Description: `Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.`,
				},
				resource.Attribute{
					Name:        "backup_black_list",
					Description: `The backup for database instance such as "test.%" or table such as "city.address" specified in the black lists are not supported.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to database instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the status of database instance , possible values are: ` + "`" + `Init` + "`" + `, ` + "`" + `Fail` + "`" + `, ` + "`" + `Starting` + "`" + `, ` + "`" + `Running` + "`" + `, ` + "`" + `Shutdown` + "`" + `, ` + "`" + `Shutoff` + "`" + `, ` + "`" + `Delete` + "`" + `, ` + "`" + `Upgrading` + "`" + `, ` + "`" + `Promoting` + "`" + `, ` + "`" + `Recovering` + "`" + ` and ` + "`" + `Recover fail` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of database instance , formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of database instance , formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The modification time of database instance , formatted by RFC3339 time string.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "db_instances",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of database instances that satisfy the condition. - - - The attribute (` + "`" + `db_instances` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone where database instance is located.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of database instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of database instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Specifies the type of database instance.`,
				},
				resource.Attribute{
					Name:        "standby_zone",
					Description: `Availability zone where the standby database instance is located for the high availability database instance with multiple zone.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC linked to the database instances.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of subnet linked to the database instances.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The type of database instance engine.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `The database instance engine version.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port on which the database instance accepts connections.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address assigned to the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_storage",
					Description: `Specifies the allocated storage size in gigabytes (GB).`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of db instance.`,
				},
				resource.Attribute{
					Name:        "backup_count",
					Description: `Specifies the number of backup saved per week.`,
				},
				resource.Attribute{
					Name:        "backup_begin_time",
					Description: `Specifies when the backup starts, measured in hour.`,
				},
				resource.Attribute{
					Name:        "backup_date",
					Description: `Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disabled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.`,
				},
				resource.Attribute{
					Name:        "backup_black_list",
					Description: `The backup for database instance such as "test.%" or table such as "city.address" specified in the black lists are not supported.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to database instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the status of database instance , possible values are: ` + "`" + `Init` + "`" + `, ` + "`" + `Fail` + "`" + `, ` + "`" + `Starting` + "`" + `, ` + "`" + `Running` + "`" + `, ` + "`" + `Shutdown` + "`" + `, ` + "`" + `Shutoff` + "`" + `, ` + "`" + `Delete` + "`" + `, ` + "`" + `Upgrading` + "`" + `, ` + "`" + `Promoting` + "`" + `, ` + "`" + `Recovering` + "`" + ` and ` + "`" + `Recover fail` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of database instance , formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of database instance , formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The modification time of database instance , formatted by RFC3339 time string.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_disks",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Disk resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability zone where Disk are located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Disk IDs, all the Disks belong to this region will be retrieved if the ID is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional) The type of disk. Possible values are: ` + "`" + `data_disk` + "`" + `as cloud disk, ` + "`" + `ssd_data_disk` + "`" + ` as SSD cloud disk, ` + "`" + `system_disk` + "`" + `as system disk, ` + "`" + `ssd_system_disk` + "`" + ` as SSD system disk, ` + "`" + `rssd_data_disk` + "`" + ` as RDMA-SSD cloud disk.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting disks by name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "disks",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Disks that satisfy the condition. - - - The attribute (` + "`" + `disks` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone where disk is located.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Disk.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Disk.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The size of disk. Purchase the size of disk in GB.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The type of disk.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of disk. Possible values are: ` + "`" + `year` + "`" + ` as pay by year, ` + "`" + `month` + "`" + ` as pay by month, ` + "`" + `dynamic` + "`" + ` as pay by hour.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to Disk.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of Disk, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of disk, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of disk. Possible values are: ` + "`" + `Available` + "`" + `, ` + "`" + `InUse` + "`" + `, ` + "`" + `Detaching` + "`" + `, ` + "`" + `Initializating` + "`" + `, ` + "`" + `Failed` + "`" + `, ` + "`" + `Cloning` + "`" + `, ` + "`" + `Restoring` + "`" + `, ` + "`" + `RestoreFailed` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "disks",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Disks that satisfy the condition. - - - The attribute (` + "`" + `disks` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone where disk is located.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Disk.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Disk.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The size of disk. Purchase the size of disk in GB.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The type of disk.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of disk. Possible values are: ` + "`" + `year` + "`" + ` as pay by year, ` + "`" + `month` + "`" + ` as pay by month, ` + "`" + `dynamic` + "`" + ` as pay by hour.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to Disk.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of Disk, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of disk, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of disk. Possible values are: ` + "`" + `Available` + "`" + `, ` + "`" + `InUse` + "`" + `, ` + "`" + `Detaching` + "`" + `, ` + "`" + `Initializating` + "`" + `, ` + "`" + `Failed` + "`" + `, ` + "`" + `Cloning` + "`" + `, ` + "`" + `Restoring` + "`" + `, ` + "`" + `RestoreFailed` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_eips",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of EIP resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Elastic IP IDs, all the EIPs belong to this region will be retrieved if the ID is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting eips by name.`,
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
					Description: `Total number of Elastic IPs that satisfy the condition. - - - The attribute (` + "`" + `eips` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Maximum bandwidth to the elastic public network, measured in Mbps.`,
				},
				resource.Attribute{
					Name:        "ip_set",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of Elastic IP, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time for Elastic IP, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "charge_mode",
					Description: `The charge mode of Elastic IP. Possible values are: ` + "`" + `traffic` + "`" + ` as pay by traffic, ` + "`" + `bandwidth` + "`" + ` as pay by bandwidth.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of Elastic IP. Possible values are: ` + "`" + `year` + "`" + ` as pay by year, ` + "`" + `month` + "`" + ` as pay by month, ` + "`" + `dynamic` + "`" + ` as pay by hour.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Elastic IP.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of Elastic IP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Elastic IP status. Possible values are: ` + "`" + `used` + "`" + ` as in use, ` + "`" + `free` + "`" + ` as available and ` + "`" + `freeze` + "`" + ` as associating.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to Elastic IP. The attribute (` + "`" + `ip_set` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "internet_type",
					Description: `Type of Elastic IP routes.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Elastic IP address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "eips",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Elastic IPs that satisfy the condition. - - - The attribute (` + "`" + `eips` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Maximum bandwidth to the elastic public network, measured in Mbps.`,
				},
				resource.Attribute{
					Name:        "ip_set",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of Elastic IP, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time for Elastic IP, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "charge_mode",
					Description: `The charge mode of Elastic IP. Possible values are: ` + "`" + `traffic` + "`" + ` as pay by traffic, ` + "`" + `bandwidth` + "`" + ` as pay by bandwidth.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of Elastic IP. Possible values are: ` + "`" + `year` + "`" + ` as pay by year, ` + "`" + `month` + "`" + ` as pay by month, ` + "`" + `dynamic` + "`" + ` as pay by hour.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Elastic IP.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of Elastic IP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Elastic IP status. Possible values are: ` + "`" + `used` + "`" + ` as in use, ` + "`" + `free` + "`" + ` as available and ` + "`" + `freeze` + "`" + ` as associating.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to Elastic IP. The attribute (` + "`" + `ip_set` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "internet_type",
					Description: `Type of Elastic IP routes.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Elastic IP address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_images",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available image resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability zone where images are located. such as: ` + "`" + `cn-bj2-02` + "`" + `. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist).`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting images by name. (Such as: ` + "`" + `^CentOS 7.[1-2] 64` + "`" + ` means CentOS 7.1 of 64-bit operating system or CentOS 7.2 of 64-bit operating system, "^Ubuntu 16.04 64" means Ubuntu 16.04 of 64-bit operating system).`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `(Optional) The type of image. Possible values are: ` + "`" + `base` + "`" + ` as standard image, ` + "`" + `business` + "`" + ` as owned by market place, and ` + "`" + `custom` + "`" + ` as custom-image, all the image types will be retrieved by default.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional) The type of OS. Possible values are: ` + "`" + `linux` + "`" + ` and ` + "`" + `windows` + "`" + `, all the OS types will be retrieved by default.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The ID of image. ~>`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of image IDs, all the images belong to this region will be retrieved if the ID is ` + "`" + `[]` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of images that satisfy the condition. - - - The attribute (` + "`" + `images` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone where image is located.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for image, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `To identify if any particular feature belongs to the instance, the value is ` + "`" + `NetEnhnced` + "`" + ` as I/O enhanced instance for now.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of image if any.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of image.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of image.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of image.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `The name of OS.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The type of OS.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of image. Possible values are ` + "`" + `Available` + "`" + `, ` + "`" + `Making` + "`" + ` and ` + "`" + `Unavailable` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "images",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of images that satisfy the condition. - - - The attribute (` + "`" + `images` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone where image is located.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for image, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `To identify if any particular feature belongs to the instance, the value is ` + "`" + `NetEnhnced` + "`" + ` as I/O enhanced instance for now.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of image if any.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of image.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of image.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of image.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `The name of OS.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The type of OS.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of image. Possible values are ` + "`" + `Available` + "`" + `, ` + "`" + `Making` + "`" + ` and ` + "`" + `Unavailable` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of UHost instance resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability zone where instances are located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance IDs, all the instances belongs to the defined region will be retrieved if this argument is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting instances by name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag assigned to instance. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `It is a nested type. instances documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of instances that satisfy the condition. - - - The attribute (` + "`" + `instances` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone where instances are located.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the instance.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `The number of cores of virtual CPU, measureed in core.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The size of memory, measured in MB (Megabyte).`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of instance, possible values are: ` + "`" + `year` + "`" + `, ` + "`" + `month` + "`" + ` and ` + "`" + `dynamic` + "`" + ` as pay by hour.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `Whether to renew an instance automatically or not.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of instance.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to the instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance current status. Possible values are ` + "`" + `Initializing` + "`" + `, ` + "`" + `Starting` + "`" + `, ` + "`" + `Running` + "`" + `, ` + "`" + `Stopping` + "`" + `, ` + "`" + `Stopped` + "`" + `, ` + "`" + `Install Fail` + "`" + ` and ` + "`" + `Rebooting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for instance, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time for instance, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address assigned to the instance.`,
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
					Name:        "ip_set",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "disk_set",
					Description: `It is a nested type which documented below. The attribute (` + "`" + `disk_set` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of disk, measured in GB (Gigabyte).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of disk.`,
				},
				resource.Attribute{
					Name:        "is_boot",
					Description: `Specifies whether boot disk or not. The attribute (` + "`" + `ip_set` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "internet_type",
					Description: `Type of Elastic IP routes.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Elastic IP address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instances",
					Description: `It is a nested type. instances documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of instances that satisfy the condition. - - - The attribute (` + "`" + `instances` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone where instances are located.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the instance.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `The number of cores of virtual CPU, measureed in core.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The size of memory, measured in MB (Megabyte).`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of instance, possible values are: ` + "`" + `year` + "`" + `, ` + "`" + `month` + "`" + ` and ` + "`" + `dynamic` + "`" + ` as pay by hour.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `Whether to renew an instance automatically or not.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of instance.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to the instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance current status. Possible values are ` + "`" + `Initializing` + "`" + `, ` + "`" + `Starting` + "`" + `, ` + "`" + `Running` + "`" + `, ` + "`" + `Stopping` + "`" + `, ` + "`" + `Stopped` + "`" + `, ` + "`" + `Install Fail` + "`" + ` and ` + "`" + `Rebooting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for instance, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time for instance, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address assigned to the instance.`,
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
					Name:        "ip_set",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "disk_set",
					Description: `It is a nested type which documented below. The attribute (` + "`" + `disk_set` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of disk, measured in GB (Gigabyte).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of disk.`,
				},
				resource.Attribute{
					Name:        "is_boot",
					Description: `Specifies whether boot disk or not. The attribute (` + "`" + `ip_set` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "internet_type",
					Description: `Type of Elastic IP routes.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Elastic IP address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_lb_attachments",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Load Balancer Attachment resources under the Load Balancer listener.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) The ID of a load balancer.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) The ID of a listener server.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of LB Attachment IDs, all the LB Attachments belong to the Load Balancer listener will be retrieved if the ID is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "lb_attachments",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of LB Attachments that satisfy the condition. - - - The attribute (` + "`" + `lb_attachments` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of LB Attachment.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The ID of a backend server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port opened on the backend server to receive requests, range: 1-65535.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private ip address for backend servers.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of backend servers. Possible values are: ` + "`" + `normalRunning` + "`" + `, ` + "`" + `exceptionRunning` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "lb_attachments",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of LB Attachments that satisfy the condition. - - - The attribute (` + "`" + `lb_attachments` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of LB Attachment.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The ID of a backend server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port opened on the backend server to receive requests, range: 1-65535.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private ip address for backend servers.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of backend servers. Possible values are: ` + "`" + `normalRunning` + "`" + `, ` + "`" + `exceptionRunning` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_lb_listeners",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Load Balancer Listener resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) The ID of a load balancer. - - -`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of LB Listener IDs, all the LB Listeners belong to this region will be retrieved if the ID is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting lb listeners by name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "lb_listeners",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of LB listeners that satisfy the condition. - - - The attribute (` + "`" + `lb_listeners` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of LB Listener.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of LB Listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `LB Listener protocol. Possible values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `request_proxy` + "`" + `, ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + ` if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `packets_transmit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "listen_type",
					Description: `The type of LB Listener. Possible values are ` + "`" + `request_proxy` + "`" + ` and ` + "`" + `packets_transmit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port opened on the LB Listener to receive requests, range: 1-65535.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Amount of time in seconds to wait for the response for in between two sessions if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `request_proxy` + "`" + `, range: 0-86400. Amount of time in seconds to wait for one session if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `packets_transmit` + "`" + `, range: 60-900. The session will be closed as soon as no response if it is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `The load balancer method in which the listener is. Possible values are: ` + "`" + `roundrobin` + "`" + `, ` + "`" + `source` + "`" + `, ` + "`" + `consistent_hash` + "`" + `, ` + "`" + `source_port` + "`" + ` , ` + "`" + `consistent_hash_port` + "`" + `, ` + "`" + `weight_roundrobin` + "`" + ` and ` + "`" + `leastconn` + "`" + `. - The ` + "`" + `consistent_hash` + "`" + `, ` + "`" + `source_port` + "`" + ` , ` + "`" + `consistent_hash_port` + "`" + `, ` + "`" + `roundrobin` + "`" + `, ` + "`" + `source` + "`" + ` and ` + "`" + `weight_roundrobin` + "`" + ` are valid if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `packets_transmit` + "`" + `. - The ` + "`" + `rundrobin` + "`" + `, ` + "`" + `source` + "`" + ` and ` + "`" + `weight_roundrobin` + "`" + ` and ` + "`" + `leastconn` + "`" + ` are vaild if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `request_proxy` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `Indicate whether the persistence session is enabled, it is invaild if ` + "`" + `persistence_type` + "`" + ` is ` + "`" + `none` + "`" + `, an auto-generated string will be exported if ` + "`" + `persistence_type` + "`" + ` is ` + "`" + `server_insert` + "`" + `, a custom string will be exported if ` + "`" + `persistence_type` + "`" + ` is ` + "`" + `user_defined` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "persistence_type",
					Description: `The type of session persistence of LB Listener. Possible values are: ` + "`" + `none` + "`" + ` as disabled, ` + "`" + `server_insert` + "`" + ` as auto-generated string and ` + "`" + `user_defined` + "`" + ` as cutom string. (Default: ` + "`" + `none` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `Health check method. Possible values are ` + "`" + `port` + "`" + ` as port checking and ` + "`" + `path` + "`" + ` as http checking.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Health check path checking.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Health check domain checking.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `LB Listener status. Possible values are: ` + "`" + `allNormal` + "`" + ` for all resource functioning well, ` + "`" + `partNormal` + "`" + ` for partial resource functioning well and ` + "`" + `allException` + "`" + ` for all resource functioning exceptional.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "lb_listeners",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of LB listeners that satisfy the condition. - - - The attribute (` + "`" + `lb_listeners` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of LB Listener.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of LB Listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `LB Listener protocol. Possible values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `request_proxy` + "`" + `, ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + ` if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `packets_transmit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "listen_type",
					Description: `The type of LB Listener. Possible values are ` + "`" + `request_proxy` + "`" + ` and ` + "`" + `packets_transmit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port opened on the LB Listener to receive requests, range: 1-65535.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Amount of time in seconds to wait for the response for in between two sessions if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `request_proxy` + "`" + `, range: 0-86400. Amount of time in seconds to wait for one session if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `packets_transmit` + "`" + `, range: 60-900. The session will be closed as soon as no response if it is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `The load balancer method in which the listener is. Possible values are: ` + "`" + `roundrobin` + "`" + `, ` + "`" + `source` + "`" + `, ` + "`" + `consistent_hash` + "`" + `, ` + "`" + `source_port` + "`" + ` , ` + "`" + `consistent_hash_port` + "`" + `, ` + "`" + `weight_roundrobin` + "`" + ` and ` + "`" + `leastconn` + "`" + `. - The ` + "`" + `consistent_hash` + "`" + `, ` + "`" + `source_port` + "`" + ` , ` + "`" + `consistent_hash_port` + "`" + `, ` + "`" + `roundrobin` + "`" + `, ` + "`" + `source` + "`" + ` and ` + "`" + `weight_roundrobin` + "`" + ` are valid if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `packets_transmit` + "`" + `. - The ` + "`" + `rundrobin` + "`" + `, ` + "`" + `source` + "`" + ` and ` + "`" + `weight_roundrobin` + "`" + ` and ` + "`" + `leastconn` + "`" + ` are vaild if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `request_proxy` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `Indicate whether the persistence session is enabled, it is invaild if ` + "`" + `persistence_type` + "`" + ` is ` + "`" + `none` + "`" + `, an auto-generated string will be exported if ` + "`" + `persistence_type` + "`" + ` is ` + "`" + `server_insert` + "`" + `, a custom string will be exported if ` + "`" + `persistence_type` + "`" + ` is ` + "`" + `user_defined` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "persistence_type",
					Description: `The type of session persistence of LB Listener. Possible values are: ` + "`" + `none` + "`" + ` as disabled, ` + "`" + `server_insert` + "`" + ` as auto-generated string and ` + "`" + `user_defined` + "`" + ` as cutom string. (Default: ` + "`" + `none` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `Health check method. Possible values are ` + "`" + `port` + "`" + ` as port checking and ` + "`" + `path` + "`" + ` as http checking.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Health check path checking.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Health check domain checking.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `LB Listener status. Possible values are: ` + "`" + `allNormal` + "`" + ` for all resource functioning well, ` + "`" + `partNormal` + "`" + ` for partial resource functioning well and ` + "`" + `allException` + "`" + ` for all resource functioning exceptional.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_lb_rules",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Load Balancer Rule resources belong to the Load Balancer listener.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) The ID of a load balancer.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) The ID of a listener server. - - -`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of LB Rule IDs, all the LB Rules belong to the Load Balancer listener will be retrieved if the ID is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "lb_rules",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of LB Rules that satisfy the condition. - - - The attribute (` + "`" + `lb_rules` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of LB Rule.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path of Content forward matching fields. ` + "`" + `path` + "`" + ` and ` + "`" + `domain` + "`" + ` cannot coexist.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain of content forward matching fields. ` + "`" + `path` + "`" + ` and ` + "`" + `domain` + "`" + ` cannot coexist.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "lb_rules",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of LB Rules that satisfy the condition. - - - The attribute (` + "`" + `lb_rules` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of LB Rule.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path of Content forward matching fields. ` + "`" + `path` + "`" + ` and ` + "`" + `domain` + "`" + ` cannot coexist.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain of content forward matching fields. ` + "`" + `path` + "`" + ` and ` + "`" + `domain` + "`" + ` cannot coexist.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_lb_ssls",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Load Balancer SSL certificate resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of LB SSL certificate resource IDs, all the LB SSL certificate resources in the current region will be retrieved if the ID is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting LB SSL by name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "lb_ssls",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of LB SSL certificate resources that satisfy the condition. - - - The attribute (` + "`" + `lb_ssls` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of LB SSL certificate resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of LB SSL certificate resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for lb ssl, formatted in RFC3339 time string.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "lb_ssls",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of LB SSL certificate resources that satisfy the condition. - - - The attribute (` + "`" + `lb_ssls` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of LB SSL certificate resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of LB SSL certificate resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for lb ssl, formatted in RFC3339 time string.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_lbs",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Load Balancer resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Load Balancer IDs, all the LBs belong to this region will be retrieved if the ID is ` + "`" + `[]` + "`" + `.`,
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
					Name:        "lbs",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Load Balancers that satisfy the condition. - - - The attribute (` + "`" + `lbs` + "`" + `) support the following:`,
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
					Name:        "internal",
					Description: `Indicate whether the load balancer is intranet.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to Load Balancer.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of Load Balancer.`,
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
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of Load Balancer, formatted in RFC3339 time string. The attribute (` + "`" + `ip_set` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "internet_type",
					Description: `Type of Load Balancer routes.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Load Balancer address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "lbs",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Load Balancers that satisfy the condition. - - - The attribute (` + "`" + `lbs` + "`" + `) support the following:`,
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
					Name:        "internal",
					Description: `Indicate whether the load balancer is intranet.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to Load Balancer.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of Load Balancer.`,
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
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of Load Balancer, formatted in RFC3339 time string. The attribute (` + "`" + `ip_set` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "internet_type",
					Description: `Type of Load Balancer routes.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Load Balancer address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_nat_gateways",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Nat Gateway resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Nat Gateway IDs, all the Nat Gateways belongs to the defined region will be retrieved if this argument is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting Nat Gateways by name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nat_gateways",
					Description: `It is a nested type. Nat Gateways documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Nat Gateways that satisfy the condition. - - - The attribute (` + "`" + `nat_gateways` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to the Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC linked to the Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `The list of subnet ID under the VPC.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for Nat Gateway, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "ip_set",
					Description: `It is a nested type which documented below. The attribute (` + "`" + `ip_set` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "internet_type",
					Description: `Type of Elastic IP routes.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Elastic IP address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nat_gateways",
					Description: `It is a nested type. Nat Gateways documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Nat Gateways that satisfy the condition. - - - The attribute (` + "`" + `nat_gateways` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to the Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC linked to the Nat Gateway.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `The list of subnet ID under the VPC.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for Nat Gateway, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "ip_set",
					Description: `It is a nested type which documented below. The attribute (` + "`" + `ip_set` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "internet_type",
					Description: `Type of Elastic IP routes.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Elastic IP address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_projects",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of projects owned by the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "is_finance",
					Description: `(Optional) To identify if the current account is granted with financial permission.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting projects by name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "projects",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of projects that satisfy the condition. - - - The attribute (` + "`" + `projects` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of project defined.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the defined project.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `The ID of the parent project where the sub project belongs to.`,
				},
				resource.Attribute{
					Name:        "parent_name",
					Description: `The name of the parent project where the sub project belongs to.`,
				},
				resource.Attribute{
					Name:        "member_count",
					Description: `The number of members belongs to the defined project.`,
				},
				resource.Attribute{
					Name:        "resource_count",
					Description: `The number of the resounce instance belong/s to the defined project.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for instance, formatted in RFC3339 time string.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "projects",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of projects that satisfy the condition. - - - The attribute (` + "`" + `projects` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of project defined.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the defined project.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `The ID of the parent project where the sub project belongs to.`,
				},
				resource.Attribute{
					Name:        "parent_name",
					Description: `The name of the parent project where the sub project belongs to.`,
				},
				resource.Attribute{
					Name:        "member_count",
					Description: `The number of members belongs to the defined project.`,
				},
				resource.Attribute{
					Name:        "resource_count",
					Description: `The number of the resounce instance belong/s to the defined project.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for instance, formatted in RFC3339 time string.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_security_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Security Group resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Security Group IDs, all the Security Group resources belong to this region will be retrieved if the ID is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting Security Group resources by name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of Security Group. Possible values are: ` + "`" + `recommend_web` + "`" + ` as the default Web security group that UCloud recommend to users, default opened port include 80, 443, 22, 3389, ` + "`" + `recommend_non_web` + "`" + ` as the default non Web security group that UCloud recommend to users, default opened port include 22, 3389, ` + "`" + `user_defined` + "`" + ` as the security groups defined by users. You may refer to [security group](https://docs.ucloud.cn/network/firewall/firewall.html).`,
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
					Description: `Total number of Security Group resources that satisfy the condition. - - - The attribute (` + "`" + `security_groups` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Security Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Security Group.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of Security Group.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of the security group.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to the security group.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for the security group, formatted in RFC3339 time string. The attribute (` + "`" + `rules` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The cidr block of source.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `Authorization policy. Can be either ` + "`" + `accept` + "`" + ` or ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `The range of port numbers, range: 1-65535. (eg: ` + "`" + `port` + "`" + ` or ` + "`" + `port1-port2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Rule priority. Can be ` + "`" + `high` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `low` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol. Can be ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `gre` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "security_groups",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Security Group resources that satisfy the condition. - - - The attribute (` + "`" + `security_groups` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Security Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Security Group.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of Security Group.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of the security group.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to the security group.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for the security group, formatted in RFC3339 time string. The attribute (` + "`" + `rules` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The cidr block of source.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `Authorization policy. Can be either ` + "`" + `accept` + "`" + ` or ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `The range of port numbers, range: 1-65535. (eg: ` + "`" + `port` + "`" + ` or ` + "`" + `port1-port2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Rule priority. Can be ` + "`" + `high` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `low` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol. Can be ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `gre` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_subnets",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Subnet resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Subnet IDs, all the Subnet resources belong to this region will be retrieved if the ID is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The id of the VPC that the desired Subnet belongs to.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting Subnet resources by name.`,
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
					Description: `Total number of Subnet resources that satisfy the condition. - - - The attribute (` + "`" + `subnets` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Subnet.`,
				},
				resource.Attribute{
					Name:        "name",
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
				resource.Attribute{
					Name:        "remark",
					Description: `The remark of the Subnet.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to Subnet.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subnets",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of Subnet resources that satisfy the condition. - - - The attribute (` + "`" + `subnets` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of Subnet.`,
				},
				resource.Attribute{
					Name:        "name",
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
				resource.Attribute{
					Name:        "remark",
					Description: `The remark of the Subnet.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to Subnet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_vpcs",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VPC resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of VPC IDs, all the VPC resources belong to this region will be retrieved if the ID is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting VPC resources by name.`,
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
					Description: `Total number of VPC resources that satisfy the condition. - - - The attribute (` + "`" + `vpcs` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of VPC.`,
				},
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `The CIDR blocks of VPC.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to VPC.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for VPC, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The time whenever there is a change made to VPC, formatted in RFC3339 time string.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpcs",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of VPC resources that satisfy the condition. - - - The attribute (` + "`" + `vpcs` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of VPC.`,
				},
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `The CIDR blocks of VPC.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to VPC.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for VPC, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The time whenever there is a change made to VPC, formatted in RFC3339 time string.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_vpn_connections",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VPN Connection resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of VPN Connection IDs, all the VPN Connections belongs to the defined region will be retrieved if this argument is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting VPN Connections by name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag assigned to VPN Connection.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The ID of VPC linked to the VPN Connection. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpn_connections",
					Description: `It is a nested type. VPN Connections documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of VPN Connections that satisfy the condition. - - - The attribute (` + "`" + `vpn_connections` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of VPN Connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN Connection.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of VPN Connection.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to the VPN Connection.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `The ID of VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `The ID of VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC linked to the VPN Connection.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for VPN Connection, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "ike_config",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "ipsec_config",
					Description: `It is a nested type which documented below. The attribute (` + "`" + `ike_config` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `The key used for authentication between the VPN gateway and the Customer gateway.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `The version of the IKE protocol.`,
				},
				resource.Attribute{
					Name:        "exchange_mode",
					Description: `The negotiation exchange mode of IKE V1 of VPN gateway.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `The encryption algorithm of IKE negotiation.`,
				},
				resource.Attribute{
					Name:        "authentication_algorithm",
					Description: `The authentication algorithm of IKE negotiation.`,
				},
				resource.Attribute{
					Name:        "local_id",
					Description: `The identification of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "remote_id",
					Description: `The identification of the Customer gateway.`,
				},
				resource.Attribute{
					Name:        "dh_group",
					Description: `The Diffie-Hellman group used by IKE negotiation.`,
				},
				resource.Attribute{
					Name:        "sa_life_time",
					Description: `The Security Association lifecycle as the result of IKE negotiation. The attribute (` + "`" + `ipsec_config` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "local_subnet_ids",
					Description: `The id list of Local subnet.`,
				},
				resource.Attribute{
					Name:        "remote_subnets",
					Description: `The ip address list of remote subnet.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The security protocol of IPSec negotiation.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `The encryption algorithm of IPSec negotiation.`,
				},
				resource.Attribute{
					Name:        "authentication_algorithm",
					Description: `The authentication algorithm of IPSec negotiation.`,
				},
				resource.Attribute{
					Name:        "pfs_dh_group",
					Description: `Whether the PFS of IPSec negotiation is on or off, ` + "`" + `disable` + "`" + ` as off, The Diffie-Hellman group as open.`,
				},
				resource.Attribute{
					Name:        "sa_life_time",
					Description: `The Security Association lifecycle as the result of IPSec negotiation.`,
				},
				resource.Attribute{
					Name:        "sa_life_time_bytes",
					Description: `The Security Association lifecycle in bytes as the result of IPSec negotiation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpn_connections",
					Description: `It is a nested type. VPN Connections documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of VPN Connections that satisfy the condition. - - - The attribute (` + "`" + `vpn_connections` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of VPN Connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN Connection.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of VPN Connection.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to the VPN Connection.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `The ID of VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `The ID of VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC linked to the VPN Connection.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for VPN Connection, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "ike_config",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "ipsec_config",
					Description: `It is a nested type which documented below. The attribute (` + "`" + `ike_config` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `The key used for authentication between the VPN gateway and the Customer gateway.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `The version of the IKE protocol.`,
				},
				resource.Attribute{
					Name:        "exchange_mode",
					Description: `The negotiation exchange mode of IKE V1 of VPN gateway.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `The encryption algorithm of IKE negotiation.`,
				},
				resource.Attribute{
					Name:        "authentication_algorithm",
					Description: `The authentication algorithm of IKE negotiation.`,
				},
				resource.Attribute{
					Name:        "local_id",
					Description: `The identification of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "remote_id",
					Description: `The identification of the Customer gateway.`,
				},
				resource.Attribute{
					Name:        "dh_group",
					Description: `The Diffie-Hellman group used by IKE negotiation.`,
				},
				resource.Attribute{
					Name:        "sa_life_time",
					Description: `The Security Association lifecycle as the result of IKE negotiation. The attribute (` + "`" + `ipsec_config` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "local_subnet_ids",
					Description: `The id list of Local subnet.`,
				},
				resource.Attribute{
					Name:        "remote_subnets",
					Description: `The ip address list of remote subnet.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The security protocol of IPSec negotiation.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `The encryption algorithm of IPSec negotiation.`,
				},
				resource.Attribute{
					Name:        "authentication_algorithm",
					Description: `The authentication algorithm of IPSec negotiation.`,
				},
				resource.Attribute{
					Name:        "pfs_dh_group",
					Description: `Whether the PFS of IPSec negotiation is on or off, ` + "`" + `disable` + "`" + ` as off, The Diffie-Hellman group as open.`,
				},
				resource.Attribute{
					Name:        "sa_life_time",
					Description: `The Security Association lifecycle as the result of IPSec negotiation.`,
				},
				resource.Attribute{
					Name:        "sa_life_time_bytes",
					Description: `The Security Association lifecycle in bytes as the result of IPSec negotiation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_vpn_customer_gateways",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VPN Gateway resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of VPN Customer Gateway IDs, all the VPN Customer Gateways belongs to the defined region will be retrieved if this argument is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting VPN Customer Gateways by name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag assigned to VPN Customer Gateway. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpn_customer_gateways",
					Description: `It is a nested type. VPN Customer Gateways documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of VPN Customer Gateways that satisfy the condition. - - - The attribute (` + "`" + `vpn_customer_gateways` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to the VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of the VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for VPN Customer Gateway, formatted in RFC3339 time string.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpn_customer_gateways",
					Description: `It is a nested type. VPN Customer Gateways documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of VPN Customer Gateways that satisfy the condition. - - - The attribute (` + "`" + `vpn_customer_gateways` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to the VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of the VPN Customer Gateway.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for VPN Customer Gateway, formatted in RFC3339 time string.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_vpn_gateways",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VPN Gateway resources in the current region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of VPN Gateway IDs, all the VPN Gateways belongs to the defined region will be retrieved if this argument is ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting VPN Gateways by name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag assigned to VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The ID of VPC linked to the VPN Gateway. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpn_gateways",
					Description: `It is a nested type. VPN Gateways documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of VPN Gateways that satisfy the condition. - - - The attribute (` + "`" + `vpn_gateways` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to the VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "grade",
					Description: `The type of the VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC linked to the VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `Whether to renew an VPN Gateway automatically or not.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for VPN Gateway, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time for VPN Gateway, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "ip_set",
					Description: `It is a nested type which documented below. The attribute (` + "`" + `ip_set` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "internet_type",
					Description: `Type of Elastic IP routes.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Elastic IP address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpn_gateways",
					Description: `It is a nested type. VPN Gateways documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of VPN Gateways that satisfy the condition. - - - The attribute (` + "`" + `vpn_gateways` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The remarks of VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `A tag assigned to the VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "grade",
					Description: `The type of the VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC linked to the VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The charge type of VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `Whether to renew an VPN Gateway automatically or not.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for VPN Gateway, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time for VPN Gateway, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "ip_set",
					Description: `It is a nested type which documented below. The attribute (` + "`" + `ip_set` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "internet_type",
					Description: `Type of Elastic IP routes.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Elastic IP address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_zones",
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
					Name:        "zones",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of zones that satisfy the condition. - - - The attribute (` + "`" + `zones` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of availability zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zones",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total number of zones that satisfy the condition. - - - The attribute (` + "`" + `zones` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of availability zone.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ucloud_db_instances":          0,
		"ucloud_disks":                 1,
		"ucloud_eips":                  2,
		"ucloud_images":                3,
		"ucloud_instances":             4,
		"ucloud_lb_attachments":        5,
		"ucloud_lb_listeners":          6,
		"ucloud_lb_rules":              7,
		"ucloud_lb_ssls":               8,
		"ucloud_lbs":                   9,
		"ucloud_nat_gateways":          10,
		"ucloud_projects":              11,
		"ucloud_security_groups":       12,
		"ucloud_subnets":               13,
		"ucloud_vpcs":                  14,
		"ucloud_vpn_connections":       15,
		"ucloud_vpn_customer_gateways": 16,
		"ucloud_vpn_gateways":          17,
		"ucloud_zones":                 18,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
