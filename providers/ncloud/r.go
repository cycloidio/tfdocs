package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ncloud_block_storage",
			Category:         "Resources",
			ShortDescription: `Provides a ncloud block storage resource.`,
			Description:      ``,
			Keywords: []string{
				"block",
				"storage",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Enter a block storage size to ceate. You can enter by the unit of GB. Up to 1000GB you can enter.`,
				},
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `(Required) Server instance No. to attach. It is required and you can get a server instance No. by calling getServerInstanceList.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Block storage name to create default : Ncloud configures it by itself.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Block storage descriptions`,
				},
				resource.Attribute{
					Name:        "disk_detail_type",
					Description: `(Optional) You can choose a disk detail type code of HDD and SSD. default : HDD ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `Block storage instance no`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Server name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Block storage type code`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `Device name`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `Block storage product code`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Block storage instance status code`,
				},
				resource.Attribute{
					Name:        "instance_operation",
					Description: `Block storage instance operation`,
				},
				resource.Attribute{
					Name:        "instance_status_name",
					Description: `Block storage instance status name`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the block storage`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Disk type code`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "instance_no",
					Description: `Block storage instance no`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Server name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Block storage type code`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `Device name`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `Block storage product code`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Block storage instance status code`,
				},
				resource.Attribute{
					Name:        "instance_operation",
					Description: `Block storage instance operation`,
				},
				resource.Attribute{
					Name:        "instance_status_name",
					Description: `Block storage instance status name`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the block storage`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Disk type code`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_block_storage_snapshot",
			Category:         "Resources",
			ShortDescription: `Provides a ncloud block storage snapshot resource.`,
			Description:      ``,
			Keywords: []string{
				"block",
				"storage",
				"snapshot",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "block_storage_instance_no",
					Description: `(Required) Block storage instance No for creating snapshot.`,
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
			Attributes: []resource.Argument{
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
			Type:             "ncloud_load_balancer",
			Category:         "Resources",
			ShortDescription: `Provides a ncloud load balancer instance resource.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
			},
			Arguments: []resource.Argument{
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
					Name:        "internet_line_type",
					Description: `(Optional) Internet line identification code. PUBLC(Public), GLBL(Global). default : PUBLC(Public)`,
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
			Attributes: []resource.Argument{
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
			ShortDescription: `Provides a ncloud load balancer ssl certificate resource.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"ssl",
				"certificate",
			},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_login_key",
			Category:         "Resources",
			ShortDescription: `Provides a ncloud login key resource.`,
			Description:      ``,
			Keywords: []string{
				"login",
				"key",
			},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			ShortDescription: `Provides a ncloud NAS volume.`,
			Description:      ``,
			Keywords: []string{
				"nas",
				"volume",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "volume_name_postfix",
					Description: `(Required) Name of a NAS volume to create. Enter a volume name that can be 3-20 characters in length after the name already entered for user identification.`,
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
					Description: `(Optional) List of server instance numbers for which access to NFS is to be controlled`,
				},
				resource.Attribute{
					Name:        "custom_ip_list",
					Description: `(Optional) To add a server of another account to the NAS volume, enter a private IP address of the server.`,
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
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: KR region.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. Zone in which you want to create a NAS volume. Default: The first zone of the region. Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "volume_name",
					Description: `NAS volume name.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `NAS Volume instance status code`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the NAS volume`,
				},
				resource.Attribute{
					Name:        "volume_total_size",
					Description: `Volume total size, in GiB`,
				},
				resource.Attribute{
					Name:        "volume_use_size",
					Description: `Volume use size, in GiB`,
				},
				resource.Attribute{
					Name:        "volume_use_ratio",
					Description: `Volume use ratio`,
				},
				resource.Attribute{
					Name:        "snapshot_volume_size",
					Description: `Snapshot volume size, in GiB`,
				},
				resource.Attribute{
					Name:        "snapshot_volume_use_size",
					Description: `Snapshot volume use size`,
				},
				resource.Attribute{
					Name:        "snapshot_volume_use_ratio",
					Description: `Snapshot volume use ratio`,
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
					Name:        "instance_custom_ip_list",
					Description: `NAS volume instance custom IP list`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "volume_name",
					Description: `NAS volume name.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `NAS Volume instance status code`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the NAS volume`,
				},
				resource.Attribute{
					Name:        "volume_total_size",
					Description: `Volume total size, in GiB`,
				},
				resource.Attribute{
					Name:        "volume_use_size",
					Description: `Volume use size, in GiB`,
				},
				resource.Attribute{
					Name:        "volume_use_ratio",
					Description: `Volume use ratio`,
				},
				resource.Attribute{
					Name:        "snapshot_volume_size",
					Description: `Snapshot volume size, in GiB`,
				},
				resource.Attribute{
					Name:        "snapshot_volume_use_size",
					Description: `Snapshot volume use size`,
				},
				resource.Attribute{
					Name:        "snapshot_volume_use_ratio",
					Description: `Snapshot volume use ratio`,
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
					Name:        "instance_custom_ip_list",
					Description: `NAS volume instance custom IP list`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_port_forwarding_rule",
			Category:         "Resources",
			ShortDescription: `Provides a ncloud port forwarding rule resource.`,
			Description:      ``,
			Keywords: []string{
				"port",
				"forwarding",
				"rule",
			},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			ShortDescription: `Provides a ncloud public IP instance resource.`,
			Description:      ``,
			Keywords: []string{
				"public",
				"ip",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `(Optional) Server instance No. to assign after creating a public IP. You can get one by calling getPublicIpTargetServerInstanceList.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Public IP description.`,
				},
				resource.Attribute{
					Name:        "internet_line_type",
					Description: `(Optional) Internet line code. PUBLC(Public), GLBL(Global)`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. You can decide a zone where servers are created. You can decide which zone the product list will be requested at. default : Select the first Zone in the specific region Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `Public IP instance No.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP Address.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the public IP instance`,
				},
				resource.Attribute{
					Name:        "instance_status_name",
					Description: `Public IP instance status name`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Public IP instance status code`,
				},
				resource.Attribute{
					Name:        "instance_operation",
					Description: `Public IP instance operation code`,
				},
				resource.Attribute{
					Name:        "kind_type",
					Description: `Public IP kind type`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "instance_no",
					Description: `Public IP instance No.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP Address.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the public IP instance`,
				},
				resource.Attribute{
					Name:        "instance_status_name",
					Description: `Public IP instance status name`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Public IP instance status code`,
				},
				resource.Attribute{
					Name:        "instance_operation",
					Description: `Public IP instance operation code`,
				},
				resource.Attribute{
					Name:        "kind_type",
					Description: `Public IP kind type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_server",
			Category:         "Resources",
			ShortDescription: `Provides a ncloud server instance resource.`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "server_image_product_code",
					Description: `(Optional) Server image product code to determine which server image to create. It can be obtained through ` + "`" + `data ncloud_server_images` + "`" + `. You are required to select one among two parameters: server image product code (server_image_product_code) and member server image number(member_server_image_no).`,
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
					Name:        "name",
					Description: `(Optional) Server name to create. default: Assigned by ncloud`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Server description to create`,
				},
				resource.Attribute{
					Name:        "login_key_name",
					Description: `(Optional) The login key name to encrypt with the public key. Default : Uses the most recently created login key name`,
				},
				resource.Attribute{
					Name:        "is_protect_server_termination",
					Description: `(Optional) You can set whether or not to protect return when creating. default : false`,
				},
				resource.Attribute{
					Name:        "internet_line_type",
					Description: `(Optional) Internet line identification code. PUBLC(Public), GLBL(Global). default : PUBLC(Public)`,
				},
				resource.Attribute{
					Name:        "fee_system_type_code",
					Description: `(Optional) A rate system identification code. There are time plan(MTRAT) and flat rate (FXSUM). Default : Time plan(MTRAT)`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. You can determine the ZONE where the server will be created. Default : Assigned by NAVER Cloud Platform. Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `.`,
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
					Description: `(Required) Instance tag value ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID.`,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `Server instance number`,
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
					Name:        "base_block_storage_size",
					Description: `The size of base block storage in bytes`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `Platform type code`,
				},
				resource.Attribute{
					Name:        "is_fee_charging_monitoring",
					Description: `Fee charging monitoring`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP`,
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
					Name:        "instance_status",
					Description: `Server instance status code`,
				},
				resource.Attribute{
					Name:        "instance_status_name",
					Description: `Server instance status name`,
				},
				resource.Attribute{
					Name:        "instance_operation",
					Description: `Server instance operation code`,
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
					Description: `Port forwarding internal port`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region code`,
				},
				resource.Attribute{
					Name:        "base_block_storage_disk_type",
					Description: `Base block storage disk type code`,
				},
				resource.Attribute{
					Name:        "base_block_storage_disk_detail_type",
					Description: `Base block storage disk detail type code`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID.`,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `Server instance number`,
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
					Name:        "base_block_storage_size",
					Description: `The size of base block storage in bytes`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `Platform type code`,
				},
				resource.Attribute{
					Name:        "is_fee_charging_monitoring",
					Description: `Fee charging monitoring`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP`,
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
					Name:        "instance_status",
					Description: `Server instance status code`,
				},
				resource.Attribute{
					Name:        "instance_status_name",
					Description: `Server instance status name`,
				},
				resource.Attribute{
					Name:        "instance_operation",
					Description: `Server instance operation code`,
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
					Description: `Port forwarding internal port`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region code`,
				},
				resource.Attribute{
					Name:        "base_block_storage_disk_type",
					Description: `Base block storage disk type code`,
				},
				resource.Attribute{
					Name:        "base_block_storage_disk_detail_type",
					Description: `Base block storage disk detail type code`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"ncloud_block_storage":                 0,
		"ncloud_block_storage_snapshot":        1,
		"ncloud_load_balancer":                 2,
		"ncloud_load_balancer_ssl_certificate": 3,
		"ncloud_login_key":                     4,
		"ncloud_nas_volume":                    5,
		"ncloud_port_forwarding_rule":          6,
		"ncloud_public_ip":                     7,
		"ncloud_server":                        8,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
