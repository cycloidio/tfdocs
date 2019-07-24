package ucloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ucloud_db_instance",
			Category:         "DB Resources",
			ShortDescription: `Provides a Database instance resource.`,
			Description:      ``,
			Keywords: []string{
				"db",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Availability zone where database instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)`,
				},
				resource.Attribute{
					Name:        "standby_zone",
					Description: `(Optional) Availability zone where the standby database instance is located for the high availability database instance with multiple zone; The disaster recovery of data center can be activated by switching to the standby database instance for the high availability database instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password for the database instance which should have 8-30 characters. It must contain at least 3 items of Capital letters, small letter, numbers and special characters. The special characters include ` + "`" + `-_` + "`" + `. If not specified, terraform will autogenerate a password.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) The type of database engine, possible values are: "mysql", "percona".`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required) The database engine version, possible values are: "5.5", "5.6", "5.7". - 5.5/5.6/5.7 for mysql and percona engine.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of database instance, which contains 6-63 characters and only support Chinese, English, numbers, '-', '_', '.', ',', '[', ']', ':'. If not specified, terraform will autogenerate a name beginning with ` + "`" + `tf-db-instance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_storage",
					Description: `(Required) Specifies the allocated storage size in gigabytes (GB), range from 20 to 3000GB. The volume adjustment must be a multiple of 10 GB. The maximum disk volume for SSD type are： - 500GB if the memory chosen is equal or less than 8GB; - 1000GB if the memory chosen is from 12 to 24GB; - 2000GB if the memory chosen is 32GB; - 3000GB if the memory chosen is equal or more than 48GB.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) Specifies the type of database instance with format "engine-type-memory", Possible values are: - "mysql" and "percona"for engine; - "ha" as high availability version for type of database, high availability version use the dual main hot standby structure which can thoroughly solved the issue of unavailable database caused by the system downtime or hardware failure, the "ha" version only supports "mysql" and "percona" engine. - possible values for memory are: 1, 2, 4, 6, 8, 12, 16, 24, 32, 48, 64GB. You may refer to [list of instance type](https://docs.ucloud.cn/compute/terraform/specification/db_instance)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port on which the database accepts connections, the default port is 3306 for mysql and percona.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) The charge type of db instance, possible values are: ` + "`" + `year` + "`" + `, ` + "`" + `month` + "`" + ` and ` + "`" + `dynamic` + "`" + ` as pay by hour (specific permission required). (Default: ` + "`" + `month` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) The duration that you will buy the db instance (Default: ` + "`" + `1` + "`" + `). The value is ` + "`" + `0` + "`" + ` when pay by month and the instance will be vaild till the last day of that month. It is not required when ` + "`" + `dynamic` + "`" + ` (pay by hour).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The ID of VPC linked to the database instances.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of subnet.`,
				},
				resource.Attribute{
					Name:        "backup_count",
					Description: `(Optional) Specifies the number of backup saved per week, it is 7 backups saved per week by default.`,
				},
				resource.Attribute{
					Name:        "backup_begin_time",
					Description: `(Optional) Specifies when the backup starts, measured in hour, it starts at one o'clock of 1, 2, 3, 4 in the morning by default.`,
				},
				resource.Attribute{
					Name:        "backup_date",
					Description: `(Optional) Specifies whether the backup took place from Sunday to Saturday by displaying 7 digits. 0 stands for backup disbaled and 1 stands for backup enabled. The rightmost digit specifies whether the backup took place on Sunday, and the digits from right to left specify whether the backup took place from Monday to Saturday, it's mandatory required to backup twice per week at least. such as: digits "1100000" stands for the backup took place on Saturday and Friday.`,
				},
				resource.Attribute{
					Name:        "backup_black_list",
					Description: `(Optional) The backup for database such as "test.%" or table such as "city.address" specified in the black lists are not supprted.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag assigned to database instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: ` + "`" + `Default` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the status of database, possible values are: ` + "`" + `Init` + "`" + `, ` + "`" + `Fail` + "`" + `, ` + "`" + `Starting` + "`" + `, ` + "`" + `Running` + "`" + `, ` + "`" + `Shutdown` + "`" + `, ` + "`" + `Shutoff` + "`" + `, ` + "`" + `Delete` + "`" + `, ` + "`" + `Upgrading` + "`" + `, ` + "`" + `Promoting` + "`" + `, ` + "`" + `Recovering` + "`" + ` and ` + "`" + `Recover fail` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address assigned to the database instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of database, formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of database, formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The modification time of database, formatted by RFC3339 time string. ## Import DB Instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_db_instance.example udbha-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the status of database, possible values are: ` + "`" + `Init` + "`" + `, ` + "`" + `Fail` + "`" + `, ` + "`" + `Starting` + "`" + `, ` + "`" + `Running` + "`" + `, ` + "`" + `Shutdown` + "`" + `, ` + "`" + `Shutoff` + "`" + `, ` + "`" + `Delete` + "`" + `, ` + "`" + `Upgrading` + "`" + `, ` + "`" + `Promoting` + "`" + `, ` + "`" + `Recovering` + "`" + ` and ` + "`" + `Recover fail` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address assigned to the database instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of database, formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of database, formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The modification time of database, formatted by RFC3339 time string. ## Import DB Instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_db_instance.example udbha-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_disk",
			Category:         "UHost Resources",
			ShortDescription: `Provides a Cloud Disk resource.`,
			Description:      ``,
			Keywords: []string{
				"uhost",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Availability zone where cloud disk is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) The size of disk. Purchase the size of disk in GB. 1-8000 for a cloud disk, 1-4000 for SSD cloud disk.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of disk, should have 6-63 characters and only support Chinese, English, numbers, '-', '_'. If not specified, terraform will autogenerate a name beginning with ` + "`" + `tf-disk` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional) The type of disk. Possible values are: ` + "`" + `data_disk` + "`" + `as cloud disk, ` + "`" + `ssd_data_disk` + "`" + ` as ssd cloud disk, ` + "`" + `rssd_data_disk` + "`" + ` as RDMA-SSD cloud disk (the ` + "`" + `rssd_data_disk` + "`" + ` only be supported in ` + "`" + `cn-bj2-05` + "`" + `).(Default: ` + "`" + `data_disk` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) Charge type of disk. Possible values are: ` + "`" + `year` + "`" + ` as pay by year, ` + "`" + `month` + "`" + ` as pay by month, ` + "`" + `dynamic` + "`" + ` as pay by hour. (Default: ` + "`" + `month` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) The duration that you will buy the resource. (Default: ` + "`" + `1` + "`" + `). It is not required when ` + "`" + `dynamic` + "`" + ` (pay by hour), the value is ` + "`" + `0` + "`" + ` when ` + "`" + `month` + "`" + `(pay by month) and the disk will be vaild till the last day of that month.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag assigned to VPC, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: ` + "`" + `Default` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation of disk, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of disk, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of disk. Possible values are: ` + "`" + `Available` + "`" + `, ` + "`" + `InUse` + "`" + `, ` + "`" + `Detaching` + "`" + `, ` + "`" + `Initializating` + "`" + `, ` + "`" + `Failed` + "`" + `, ` + "`" + `Cloning` + "`" + `, ` + "`" + `Restoring` + "`" + `, ` + "`" + `RestoreFailed` + "`" + `. ## Import Disk can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_disk.example bsm-abcdefg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation of disk, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of disk, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of disk. Possible values are: ` + "`" + `Available` + "`" + `, ` + "`" + `InUse` + "`" + `, ` + "`" + `Detaching` + "`" + `, ` + "`" + `Initializating` + "`" + `, ` + "`" + `Failed` + "`" + `, ` + "`" + `Cloning` + "`" + `, ` + "`" + `Restoring` + "`" + `, ` + "`" + `RestoreFailed` + "`" + `. ## Import Disk can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_disk.example bsm-abcdefg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_disk_attachment",
			Category:         "UHost Resources",
			ShortDescription: `Provides a Cloud Disk Attachment resource for attaching Cloud Disk to UHost Instance.`,
			Description:      ``,
			Keywords: []string{
				"uhost",
				"disk",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) The Zone to attach the disk in.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The ID of host instance.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `(Required) The ID of disk that needs to be attached`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_eip",
			Category:         "UNet Resources",
			ShortDescription: `Provides an Elastic IP resource.`,
			Description:      ``,
			Keywords: []string{
				"unet",
				"eip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "internet_type",
					Description: `(Required) Type of Elastic IP routes. Possible values are: ` + "`" + `international` + "`" + ` as internaltional BGP IP and ` + "`" + `bgp` + "`" + ` as china BGP IP.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) Maximum bandwidth to the elastic public network, measured in Mbps (Mega bit per second). the ranges for bandwidth are: 1-200 for pay by traffic, 1-800 for pay by bandwith. (Default: ` + "`" + `1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) The duration that you will buy the resource. (Default: ` + "`" + `1` + "`" + `). It is not required when ` + "`" + `dynamic` + "`" + ` (pay by hour), the value is ` + "`" + `0` + "`" + ` when ` + "`" + `month` + "`" + `(pay by month) and the instance will be vaild till the last day of that month.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) Elastic IP charge type. Possible values are: ` + "`" + `year` + "`" + ` as pay by year, ` + "`" + `month` + "`" + ` as pay by month, ` + "`" + `dynamic` + "`" + ` as pay by hour (specific permission required). (Default: ` + "`" + `month` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the EIP, which contains 1-63 characters and only support Chinese, English, numbers, '-', '_', '.'. If not specified, terraform will autogenerate a name beginning with ` + "`" + `tf-eip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) The remarks of the EIP. (Default: ` + "`" + `""` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag assigned to Elastic IP, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: ` + "`" + `Default` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for EIP, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time for EIP, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "ip_set",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `EIP status. Possible values are: ` + "`" + `used` + "`" + ` as in use, ` + "`" + `free` + "`" + ` as available and ` + "`" + `freeze` + "`" + ` as associating.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of Elastic IP. The attribute (` + "`" + `ip_set` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "internet_type",
					Description: `Type of Elastic IP routes. The attribute (` + "`" + `resource` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource with EIP attached.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of resource with EIP attached. Possible values are ` + "`" + `instance` + "`" + ` as instance, ` + "`" + `vrouter` + "`" + ` as visual router, ` + "`" + `lb` + "`" + ` as load balancer. ## Import EIP can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_eip.example eip-abcdefg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for EIP, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time for EIP, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "ip_set",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `EIP status. Possible values are: ` + "`" + `used` + "`" + ` as in use, ` + "`" + `free` + "`" + ` as available and ` + "`" + `freeze` + "`" + ` as associating.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of Elastic IP. The attribute (` + "`" + `ip_set` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "internet_type",
					Description: `Type of Elastic IP routes. The attribute (` + "`" + `resource` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource with EIP attached.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of resource with EIP attached. Possible values are ` + "`" + `instance` + "`" + ` as instance, ` + "`" + `vrouter` + "`" + ` as visual router, ` + "`" + `lb` + "`" + ` as load balancer. ## Import EIP can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_eip.example eip-abcdefg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_eip_association",
			Category:         "UNet Resources",
			ShortDescription: `Provides an EIP Association resource for associating Elastic IP to UHost Instance, Load Balancer, etc..`,
			Description:      ``,
			Keywords: []string{
				"unet",
				"eip",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "eip_id",
					Description: `(Required) The ID of EIP.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The ID of resource with EIP attached.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_instance",
			Category:         "UHost Resources",
			ShortDescription: `Provides an UHost Instance resource.`,
			Description:      ``,
			Keywords: []string{
				"uhost",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Availability zone where instance is located. such as: ` + "`" + `cn-bj2-02` + "`" + `. You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) The ID for the image to use for the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) The type of instance. You may refer to [list of instance type](https://docs.ucloud.cn/compute/terraform/specification/instance) ~>`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `(Optional) The password for the instance, which contains 8-30 characters, and at least 2 items of capital letters, lower case letters, numbers and special characters. The special characters include <code>` + "`" + `()~!@#$%^&`,
				},
				resource.Attribute{
					Name:        "boot_disk_size",
					Description: `(Optional) The size of the boot disk, measured in GB (GigaByte). Range: 20-100. The value set of disk size must be larger or equal to ` + "`" + `20` + "`" + `(default: ` + "`" + `20` + "`" + `) for Linux and ` + "`" + `40` + "`" + ` (default: ` + "`" + `40` + "`" + `) for Windows. The responsive time is a bit longer if the value set is larger than default for local boot disk, and further settings may be required on host instance if the value set is larger than default for cloud boot disk. The disk volume adjustment must be a multiple of 10 GB. In addition, any reduction of boot disk size is not supported. ~>`,
				},
				resource.Attribute{
					Name:        "boot_disk_type",
					Description: `(Optional) The type of boot disk. Possible values are: ` + "`" + `local_normal` + "`" + ` and ` + "`" + `local_ssd` + "`" + ` for local boot disk, ` + "`" + `cloud_normal` + "`" + ` and ` + "`" + `cloud_ssd` + "`" + ` for cloud boot disk. (Default: ` + "`" + `local_normal` + "`" + `). The ` + "`" + `local_ssd` + "`" + `, ` + "`" + `cloud_normal` + "`" + ` and ` + "`" + `cloud_ssd` + "`" + ` are not fully support by all regions as boot disk type, please proceed to UCloud console for more details.`,
				},
				resource.Attribute{
					Name:        "data_disk_type",
					Description: `(Optional) The type of local data disk. Possible values are: ` + "`" + `local_normal` + "`" + ` and ` + "`" + `local_ssd` + "`" + ` for local data disk. (Default: ` + "`" + `local_normal` + "`" + `). The ` + "`" + `local_ssd` + "`" + ` is not supported in all regions as data disk type, please proceed to UCloud console for more details.`,
				},
				resource.Attribute{
					Name:        "data_disk_size",
					Description: `(Optional) The size of local data disk, measured in GB (GigaByte), range: 0-8000 (Default: ` + "`" + `20` + "`" + `), 0-8000 for cloud disk, 0-2000 for local sata disk and 100-1000 for local ssd disk (all the GPU type instances are included). The volume adjustment must be a multiple of 10 GB. In addition, any reduction of data disk size is not supported. ~>`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) The charge type of instance, possible values are: ` + "`" + `year` + "`" + `, ` + "`" + `month` + "`" + ` and ` + "`" + `dynamic` + "`" + ` as pay by hour (specific permission required). (Default: ` + "`" + `month` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) The duration that you will buy the instance (Default: ` + "`" + `1` + "`" + `). The value is ` + "`" + `0` + "`" + ` when pay by month and the instance will be vaild till the last day of that month. It is not required when ` + "`" + `dynamic` + "`" + ` (pay by hour).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of instance, which contains 1-63 characters and only support Chinese, English, numbers, '-', '_', '.'. If not specified, terraform will autogenerate a name beginning with ` + "`" + `tf-instance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) The remarks of instance. (Default: ` + "`" + `""` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `(Optional) The ID of the associated security group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The ID of VPC linked to the instance. If not defined ` + "`" + `vpc_id` + "`" + `, the instance will use the default VPC in the current region.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of subnet. If defined ` + "`" + `vpc_id` + "`" + `, the ` + "`" + `subnet_id` + "`" + ` is Required. If not defined ` + "`" + `vpc_id` + "`" + ` and ` + "`" + `subnet_id` + "`" + `, the instance will use the default subnet in the current region.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag assigned to instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: ` + "`" + `Default` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "isolation_group",
					Description: `(Optional) The ID of the associated isolation group. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `Whether to renew an instance automatically or not.`,
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
					Name:        "create_time",
					Description: `The time of creation for instance, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time for instance, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance current status. Possible values are ` + "`" + `Initializing` + "`" + `, ` + "`" + `Starting` + "`" + `, ` + "`" + `Running` + "`" + `, ` + "`" + `Stopping` + "`" + `, ` + "`" + `Stopped` + "`" + `, ` + "`" + `Install Fail` + "`" + `, ` + "`" + `ResizeFail` + "`" + ` and ` + "`" + `Rebooting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address assigned to the instance.`,
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
					Description: `Type of Elastic IP routes. Possible values are: ` + "`" + `International` + "`" + ` as internaltional BGP IP, ` + "`" + `BGP` + "`" + ` as china BGP IP and ` + "`" + `Private` + "`" + ` as private IP.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Elastic IP address. ## Import Instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_instance.example uhost-abcdefg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "auto_renew",
					Description: `Whether to renew an instance automatically or not.`,
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
					Name:        "create_time",
					Description: `The time of creation for instance, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time for instance, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance current status. Possible values are ` + "`" + `Initializing` + "`" + `, ` + "`" + `Starting` + "`" + `, ` + "`" + `Running` + "`" + `, ` + "`" + `Stopping` + "`" + `, ` + "`" + `Stopped` + "`" + `, ` + "`" + `Install Fail` + "`" + `, ` + "`" + `ResizeFail` + "`" + ` and ` + "`" + `Rebooting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address assigned to the instance.`,
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
					Description: `Type of Elastic IP routes. Possible values are: ` + "`" + `International` + "`" + ` as internaltional BGP IP, ` + "`" + `BGP` + "`" + ` as china BGP IP and ` + "`" + `Private` + "`" + ` as private IP.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Elastic IP address. ## Import Instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_instance.example uhost-abcdefg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_isolation_group",
			Category:         "UHost Resources",
			ShortDescription: `Provides an Isolation Group resource.`,
			Description:      ``,
			Keywords: []string{
				"uhost",
				"isolation",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the isolation group information which contains 1-63 characters and only support Chinese, English, numbers, '-', '_', '.', ',', '[', ']', ':'. If not specified, terraform will autogenerate a name beginning with ` + "`" + `tf-isolation-group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) The remarks of the isolation group. (Default: ` + "`" + `""` + "`" + `). ## Import Isolation Group can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_isolation_group.example ig-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_lb",
			Category:         "ULB Resources",
			ShortDescription: `Provides a Load Balancer resource.`,
			Description:      ``,
			Keywords: []string{
				"ulb",
				"lb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "internal",
					Description: `(Optional) Indicate whether the load balancer is intranet mode.(Default: ` + "`" + `"false"` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the load balancer. If not specified, terraform will autogenerate a name beginning with ` + "`" + `tf-lb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The ID of the VPC linked to the Load Balancers, This argumnet is not required if default VPC.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of subnet that intrant load balancer belongs to. This argumnet is not required if default subnet.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag assigned to load balancer, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: ` + "`" + `Default` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) The remarks of the load balancer. (Default: ` + "`" + `""` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for load balancer, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "ip_set",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The IP address of intranet IP. It is ` + "`" + `""` + "`" + ` if ` + "`" + `internal` + "`" + ` is ` + "`" + `false` + "`" + `. The attribute (` + "`" + `ip_set` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "internet_type",
					Description: `Type of Elastic IP routes.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Elastic IP address. ## Import LB can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_lb.example ulb-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for load balancer, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "ip_set",
					Description: `It is a nested type which documented below.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The IP address of intranet IP. It is ` + "`" + `""` + "`" + ` if ` + "`" + `internal` + "`" + ` is ` + "`" + `false` + "`" + `. The attribute (` + "`" + `ip_set` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "internet_type",
					Description: `Type of Elastic IP routes.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Elastic IP address. ## Import LB can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_lb.example ulb-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_lb_attachment",
			Category:         "ULB Resources",
			ShortDescription: `Provides a Load Balancer Attachment resource for attaching Load Balancer to UHost Instance, etc.`,
			Description:      ``,
			Keywords: []string{
				"ulb",
				"lb",
				"attachment",
			},
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
					Name:        "resource_id",
					Description: `(Required) The ID of a backend server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port opened on the backend server to receive requests, range: 1-65535, (Default: ` + "`" + `80` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "ucloud_lb_listener",
			Category:         "ULB Resources",
			ShortDescription: `Provides a Load Balancer Listener resource.`,
			Description:      ``,
			Keywords: []string{
				"ulb",
				"lb",
				"listener",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) The ID of load balancer instance.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Listener protocol. Possible values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `tcp` + "`" + ` if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `request_proxy` + "`" + `, ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + ` if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `packets_transmit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the listener. If not specified, terraform will autogenerate a name beginning with ` + "`" + `tf-lb-listener` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "listen_type",
					Description: `(Optional) The type of listener. Possible values are ` + "`" + `request_proxy` + "`" + ` and ` + "`" + `packets_transmit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port opened on the listeners to receive requests, range: 1-65535. The default value: ` + "`" + `80` + "`" + ` as ` + "`" + `protocol` + "`" + ` is ` + "`" + `http` + "`" + `, ` + "`" + `443` + "`" + ` as ` + "`" + `protocol` + "`" + ` is ` + "`" + `https` + "`" + `, ` + "`" + `1024` + "`" + ` as ` + "`" + `protocol` + "`" + ` is ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional) Amount of time in seconds to wait for the response for in between two sessions if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `request_proxy` + "`" + `, range: 0-86400. (Default: ` + "`" + `60` + "`" + `). Amount of time in seconds to wait for one session if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `packets_transmit` + "`" + `, range: 60-900. The session will be closed as soon as no response if it is ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) The load balancer method in which the listener is. Possible values are: ` + "`" + `roundrobin` + "`" + `, ` + "`" + `source` + "`" + `, ` + "`" + `consistent_hash` + "`" + `, ` + "`" + `source_port` + "`" + ` , ` + "`" + `consistent_hash_port` + "`" + `, ` + "`" + `weight_roundrobin` + "`" + ` and ` + "`" + `leastconn` + "`" + `. (Default: ` + "`" + `roundrobin` + "`" + `). - The ` + "`" + `consistent_hash` + "`" + `, ` + "`" + `source_port` + "`" + ` , ` + "`" + `consistent_hash_port` + "`" + `, ` + "`" + `roundrobin` + "`" + `, ` + "`" + `source` + "`" + ` and ` + "`" + `weight_roundrobin` + "`" + ` are valid if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `packets_transmit` + "`" + `. - The ` + "`" + `roundrobin` + "`" + `, ` + "`" + `source` + "`" + ` and ` + "`" + `weight_roundrobin` + "`" + ` and ` + "`" + `leastconn` + "`" + ` are vaild if ` + "`" + `listen_type` + "`" + ` is ` + "`" + `request_proxy` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `(Optional) Indicate whether the persistence session is enabled, it is invaild if ` + "`" + `persistence_type` + "`" + ` is ` + "`" + `none` + "`" + `, an auto-generated string will be exported if ` + "`" + `persistence_type` + "`" + ` is ` + "`" + `server_insert` + "`" + `, a custom string will be exported if ` + "`" + `persistence_type` + "`" + ` is ` + "`" + `user_defined` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "persistence_type",
					Description: `(Optional) The type of session persistence of listener. Possible values are: ` + "`" + `none` + "`" + ` as disabled, ` + "`" + `server_insert` + "`" + ` as auto-generated string and ` + "`" + `user_defined` + "`" + ` as cutom string. (Default: ` + "`" + `none` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `(Optional) Health check method. Possible values are ` + "`" + `port` + "`" + ` as port checking and ` + "`" + `path` + "`" + ` as http checking.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Health check path checking.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Health check domain checking. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Listener status. Possible values are: ` + "`" + `allNormal` + "`" + ` for all resource functioning well, ` + "`" + `partNormal` + "`" + ` for partial resource functioning well and ` + "`" + `allException` + "`" + ` for all resource functioning exceptional. ## Import LB Listener can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_lb_listener.example vserver-abcdefg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Listener status. Possible values are: ` + "`" + `allNormal` + "`" + ` for all resource functioning well, ` + "`" + `partNormal` + "`" + ` for partial resource functioning well and ` + "`" + `allException` + "`" + ` for all resource functioning exceptional. ## Import LB Listener can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_lb_listener.example vserver-abcdefg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_lb_rule",
			Category:         "ULB Resources",
			ShortDescription: `Provides a Load Balancer Rule resource to add content forwarding policies for Load Balancer backend resource.`,
			Description:      ``,
			Keywords: []string{
				"ulb",
				"lb",
				"rule",
			},
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
					Name:        "backend_ids",
					Description: `(Required) The IDs of the backend servers where rule applies, this argument is populated base on the ` + "`" + `backend_id` + "`" + ` responed from ` + "`" + `lb_attachment` + "`" + ` create.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path of Content forward matching fields. ` + "`" + `path` + "`" + ` and ` + "`" + `domain` + "`" + ` cannot coexist. ` + "`" + `path` + "`" + ` and ` + "`" + `domain` + "`" + ` must be filled in one.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain of content forward matching fields. ` + "`" + `path` + "`" + ` and ` + "`" + `domain` + "`" + ` cannot coexist. ` + "`" + `path` + "`" + ` and ` + "`" + `domain` + "`" + ` must be filled in one.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_lb_ssl",
			Category:         "ULB Resources",
			ShortDescription: `Provides a Load Balancer SSL certificate resource.`,
			Description:      ``,
			Keywords: []string{
				"ulb",
				"lb",
				"ssl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the LB ssl, which contains 1-63 characters and only support Chinese, English, numbers, '-', '_', '.'. If not specified, terraform will autogenerate a name beginning with ` + "`" + `tf-lb-ssl` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The content of the private key about ssl certificate.`,
				},
				resource.Attribute{
					Name:        "user_cert",
					Description: `(Required) The content of the user certificate about ssl certificate.`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `(Optional) The content of the CA certificate about ssl certificate. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for lb ssl, formatted in RFC3339 time string.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for lb ssl, formatted in RFC3339 time string.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_lb_ssl_attachment",
			Category:         "ULB Resources",
			ShortDescription: `Provides a Load Balancer SSL attachment resource for attaching SSL certificate to Load Balancer Listener.`,
			Description:      ``,
			Keywords: []string{
				"ulb",
				"lb",
				"ssl",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ssl_id",
					Description: `(Required) The ID of SSL certificate.`,
				},
				resource.Attribute{
					Name:        "load_balance_id",
					Description: `(Required) The ID of load balancer instance.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) The ID of listener servers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_memcache_instance",
			Category:         "UMem Resources DB Resources DB Resources",
			ShortDescription: `Provides a Memcache instance resource.`,
			Description:      ``,
			Keywords: []string{
				"umem",
				"db",
				"memcache",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Availability zone where Memcache instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of Memcache instance, which contains 6-63 characters and only support English, numbers, '-', '_'. If not specified, terraform will autogenerate a name beginning with ` + "`" + `tf-memcache-instance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) The type of Memcache instance, please visit the [instance type table](https://docs.ucloud.cn/compute/terraform/specification/umem_instance) for more details. Specifies the type of Memcache instance with format ` + "`" + `engine-type-memory` + "`" + `, Possible values are: - ` + "`" + `engine` + "`" + ` can only be ` + "`" + `memcache` + "`" + `. - ` + "`" + `type` + "`" + ` can only be ` + "`" + `master` + "`" + ` as active-standby Memcache instance. - ` + "`" + `memory` + "`" + ` can be one of ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `16` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `32` + "`" + ` for active-standby Memcache instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) The charge type of Memcache instance, possible values are: ` + "`" + `year` + "`" + `, ` + "`" + `month` + "`" + ` and ` + "`" + `dynamic` + "`" + ` as pay by hour (specific permission required). (Default: ` + "`" + `month` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) The duration that you will buy the Memcache instance (Default: ` + "`" + `1` + "`" + `). The value is ` + "`" + `0` + "`" + ` when pay by month and the instance will be vaild till the last day of that month. It is not required when ` + "`" + `dynamic` + "`" + ` (pay by hour).`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag assigned to Memcache instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: ` + "`" + `Default` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The ID of VPC linked to the Memcache instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of subnet linked to the Memcache instance. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ip_set",
					Description: `ip_set is a nested type. ip_set documented below.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of Memcache instance, formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of Memcache instance, formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of KV Memcache instance. The attribute (` + "`" + `ip_set` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The virtual ip of Memcache instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port on which Memcache instance accepts connections, it is 6379 by default.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_set",
					Description: `ip_set is a nested type. ip_set documented below.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of Memcache instance, formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of Memcache instance, formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of KV Memcache instance. The attribute (` + "`" + `ip_set` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The virtual ip of Memcache instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port on which Memcache instance accepts connections, it is 6379 by default.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_redis_instance",
			Category:         "UMem Resources DB Resources",
			ShortDescription: `Provides a Redis instance resource.`,
			Description:      ``,
			Keywords: []string{
				"umem",
				"db",
				"redis",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Availability zone where Redis instance is located. Such as: "cn-bj2-02". You may refer to [list of availability zone](https://docs.ucloud.cn/api/summary/regionlist)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of Redis instance, which contains 6-63 characters and only support English, numbers, '-', '_'. If not specified, terraform will autogenerate a name beginning with ` + "`" + `tf-redis-instance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) The type of Redis instance, please visit the [instance type table](https://docs.ucloud.cn/compute/terraform/specification/umem_instance) for more details. Specifies the type of Redis instance with format ` + "`" + `engine-type-memory` + "`" + `, Possible values are: - ` + "`" + `engine` + "`" + ` can only be ` + "`" + `redis` + "`" + `. - ` + "`" + `type` + "`" + ` can be ` + "`" + `master` + "`" + ` and ` + "`" + `distributed` + "`" + `, ` + "`" + `master` + "`" + ` as active-standby Redis instance, ` + "`" + `distributed` + "`" + ` as distributed Redis instance. - ` + "`" + `memory` + "`" + ` can be one of ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `16` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `32` + "`" + ` for active-standby Redis instance. ` + "`" + `memory` + "`" + ` range of ` + "`" + `16` + "`" + `-` + "`" + `1000` + "`" + ` for distributed Redis instance (memory must be divisible by 4).`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional) The charge type of Redis instance, possible values are: ` + "`" + `year` + "`" + `, ` + "`" + `month` + "`" + ` and ` + "`" + `dynamic` + "`" + ` as pay by hour (specific permission required). (Default: ` + "`" + `month` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) The duration that you will buy the Redis instance (Default: ` + "`" + `1` + "`" + `). The value is ` + "`" + `0` + "`" + ` when pay by month and the instance will be vaild till the last day of that month. It is not required when ` + "`" + `dynamic` + "`" + ` (pay by hour).`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag assigned to Redis instance, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: ` + "`" + `Default` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The ID of VPC linked to the Redis instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of subnet linked to the Redis instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(active-standby Redis Required) The version of engine of active-standby Redis. Possible values are: 3.0, 3.2 and 4.0.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password for active-standby Redis instance which should have 6-36 characters. It must contain at least 3 items of Capital letters, small letter, numbers and special characters. The special characters include ` + "`" + `-_` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "ip_set",
					Description: `ip_set is a nested type. ip_set documented below.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of Redis instance, formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of Redis instance, formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of KV Redis instance. The attribute (` + "`" + `ip_set` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The virtual ip of Redis instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port on which Redis instance accepts connections, it is 6379 by default.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_set",
					Description: `ip_set is a nested type. ip_set documented below.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of Redis instance, formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of Redis instance, formatted by RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of KV Redis instance. The attribute (` + "`" + `ip_set` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The virtual ip of Redis instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port on which Redis instance accepts connections, it is 6379 by default.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_security_group",
			Category:         "UNet Resources",
			ShortDescription: `Provides a Security Group resource.`,
			Description:      ``,
			Keywords: []string{
				"unet",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "rules",
					Description: `(Required) A list of security group rules. Can be specified multiple times for each rules. Each rules supports fields documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the security group which contains 1-63 characters and only support Chinese, English, numbers, '-', '_' and '.'. If not specified, terraform will autogenerate a name beginning with ` + "`" + `tf-security-group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) The remarks of the security group. (Default: ` + "`" + `""` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag assigned to security group, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: ` + "`" + `Default` + "`" + `). The ` + "`" + `rules` + "`" + ` supports:`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `(Optional) The range of port numbers, range: 1-65535. (eg: ` + "`" + `port` + "`" + ` or ` + "`" + `port1-port2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) The cidr block of source.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) Authorization policy. Possible values are: ` + "`" + `accept` + "`" + `, ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Rule priority. Possible values are: ` + "`" + `high` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `low` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol. Possible values are: ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `gre` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation of security group, formatted in RFC3339 time string. ## Import Security Group can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_security_group.example firewall-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation of security group, formatted in RFC3339 time string. ## Import Security Group can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_security_group.example firewall-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_subnet",
			Category:         "VPC Resources",
			ShortDescription: `Provides a Subnet resource under VPC resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required) The cidr block of the desired subnet, format in "0.0.0.0/0", such as: ` + "`" + `192.168.0.0/24` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The id of the VPC that the desired subnet belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the desired subnet. If not specified, terraform will autogenerate a name beginning with ` + "`" + `tf-subnet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) The remarks of the subnet. (Default: ` + "`" + `""` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag assigned to subnet, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: ` + "`" + `Default` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation of subnet, formatted in RFC3339 time string. ## Import Subnet can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_subnet.example subnet-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation of subnet, formatted in RFC3339 time string. ## Import Subnet can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_subnet.example subnet-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_vpc",
			Category:         "VPC Resources",
			ShortDescription: `Provides a VPC resource.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `(Required) The CIDR blocks of VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of VPC. If not specified, terraform will autogenerate a name beginning with ` + "`" + `tf-vpc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A tag assigned to VPC, which contains at most 63 characters and only support Chinese, English, numbers, '-', '_', and '.'. If it is not filled in or a empty string is filled in, then default tag will be assigned. (Default: ` + "`" + `Default` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) The remarks of the VPC. (Default: ` + "`" + `""` + "`" + `). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for VPC, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The time whenever there is a change made to VPC, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "network_info",
					Description: `It is a nested type which documented below. The attribute (` + "`" + `network_info` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the VPC. ## Import VPC can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_vpc.example uvnet-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for VPC, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The time whenever there is a change made to VPC, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "network_info",
					Description: `It is a nested type which documented below. The attribute (` + "`" + `network_info` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the VPC. ## Import VPC can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ucloud_vpc.example uvnet-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ucloud_vpc_peering_connection",
			Category:         "VPC Resources",
			ShortDescription: `Provides an VPC Peering Connection for establishing a connection between multiple VPC.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"peering",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The short of ID of the requester VPC of the specific VPC Peering Connection to retrieve.`,
				},
				resource.Attribute{
					Name:        "peer_vpc_id",
					Description: `(Required) The short ID of accepter VPC of the specific VPC Peering Connection to retrieve.`,
				},
				resource.Attribute{
					Name:        "peer_project_id",
					Description: `(Optional) The ID of accepter project of the specific VPC Peering Connection to retrieve.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"ucloud_db_instance":            0,
		"ucloud_disk":                   1,
		"ucloud_disk_attachment":        2,
		"ucloud_eip":                    3,
		"ucloud_eip_association":        4,
		"ucloud_instance":               5,
		"ucloud_isolation_group":        6,
		"ucloud_lb":                     7,
		"ucloud_lb_attachment":          8,
		"ucloud_lb_listener":            9,
		"ucloud_lb_rule":                10,
		"ucloud_lb_ssl":                 11,
		"ucloud_lb_ssl_attachment":      12,
		"ucloud_memcache_instance":      13,
		"ucloud_redis_instance":         14,
		"ucloud_security_group":         15,
		"ucloud_subnet":                 16,
		"ucloud_vpc":                    17,
		"ucloud_vpc_peering_connection": 18,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
