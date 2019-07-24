package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ncloud_access_control_group",
			Category:         "Data Sources",
			ShortDescription: `Get access control group`,
			Description: `

When creating a server instance (VM), you can add an access control group (ACG) that you specified to set firewalls. ` + "`" + `ncloud_access_control_group` + "`" + ` provides details about a specific access control group (ACG) information.


`,
			Keywords: []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "configuration_no",
					Description: `(Optional) List of ACG configuration numbers you want to get`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the ACG you want to get`,
				},
				resource.Attribute{
					Name:        "is_default_group",
					Description: `(Optional) Indicates whether to get default group only Conditional: Requires ` + "`" + `configuration_no` + "`" + ` or` + "`" + ` name` + "`" + ` or ` + "`" + `is_default_group` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `ACG description`,
				},
			},
			Attributes: []resource.Argument{
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
			ShortDescription: `Get access control group list`,
			Description: `

When creating a server instance (VM), you can add an access control group (ACG) that you specified to set firewalls. This data source gets a list of access control groups necessary to set firewalls.

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "configuration_no_list",
					Description: `(Optional) List of ACG configuration numbers you want to get`,
				},
				resource.Attribute{
					Name:        "is_default_group",
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
			Attributes: []resource.Argument{
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
			ShortDescription: `Get access control rule`,
			Description: `

Access configuration rule you want to get

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			ShortDescription: `Get access control rule list`,
			Description: `

List of access configuration rules you want to get

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "access_control_rules",
					Description: `A list of access control rules configuration no`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_member_server_image",
			Category:         "Data Sources",
			ShortDescription: `Get member server image`,
			Description: `

Gets a member server image.

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			ShortDescription: `Get member server image list`,
			Description: `

Gets a list of member server images.

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			ShortDescription: `Get NAS volume instance`,
			Description: `

Get NAS volume instance

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
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
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: ` + "`" + `KR` + "`" + ` region.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `NAS volume instance number`,
				},
				resource.Attribute{
					Name:        "volume_name",
					Description: `Volume name`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `NAS volume instance status code`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the NAS Volume instance`,
				},
				resource.Attribute{
					Name:        "volume_total_size",
					Description: `Volume total size`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `Volume size`,
				},
				resource.Attribute{
					Name:        "volume_use_size",
					Description: `Volume use size`,
				},
				resource.Attribute{
					Name:        "volume_use_ratio",
					Description: `Volume use ratio`,
				},
				resource.Attribute{
					Name:        "snapshot_volume_size",
					Description: `Snapshot volume size`,
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
					Name:        "instance_custom_ip_list",
					Description: `NAS volume instance custom IP list`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `NAS volume description`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "instance_no",
					Description: `NAS volume instance number`,
				},
				resource.Attribute{
					Name:        "volume_name",
					Description: `Volume name`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `NAS volume instance status code`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the NAS Volume instance`,
				},
				resource.Attribute{
					Name:        "volume_total_size",
					Description: `Volume total size`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `Volume size`,
				},
				resource.Attribute{
					Name:        "volume_use_size",
					Description: `Volume use size`,
				},
				resource.Attribute{
					Name:        "volume_use_ratio",
					Description: `Volume use ratio`,
				},
				resource.Attribute{
					Name:        "snapshot_volume_size",
					Description: `Snapshot volume size`,
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
					Name:        "instance_custom_ip_list",
					Description: `NAS volume instance custom IP list`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `NAS volume description`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_nas_volumes",
			Category:         "Data Sources",
			ShortDescription: `Get NAS volume instance list`,
			Description: `

Gets a list of NAS volume instances.

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
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
					Description: `(Optional) The name of file that can save data source after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "nas_volumes",
					Description: `A list of NAS Volume Instance no`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "nas_volumes",
					Description: `A list of NAS Volume Instance no`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_port_forwarding_rule",
			Category:         "Data Sources",
			ShortDescription: `Get port forwarding rule`,
			Description: `

Get a port forwarding rule.
When a server is created for the first time, a public IP address for port forwarding is given per account.

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "internet_line_type_code",
					Description: `(Optional) Internet line code. PUBLC(Public), GLBL(Global)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: KR region.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. You can decide a zone where servers are created. You can decide which zone the product list will be requested at. default : Select the first Zone in the specific region Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `.`,
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
			Attributes: []resource.Argument{
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
			ShortDescription: `Get port forwarding rule list`,
			Description: `

Gets a list of port forwarding rules.
When a server is created for the first time, a public IP address for port forwarding is given per account.

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "internet_line_type_code",
					Description: `(Optional) Internet line code. PUBLC(Public), GLBL(Global)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: KR region.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. You can decide a zone where servers are created. You can decide which zone the product list will be requested at. default : Select the first Zone in the specific region Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `.`,
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
			Attributes: []resource.Argument{
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
			ShortDescription: `Get public IP`,
			Description: `

Get public IP instance.


`,
			Keywords: []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "internet_line_type",
					Description: `(Optional) Internet line type code. ` + "`" + `PUBLC` + "`" + ` (Public), ` + "`" + `GLBL` + "`" + ` (Global)`,
				},
				resource.Attribute{
					Name:        "is_associated",
					Description: `(Optional) Indicates whether the public IP address is associated or not.`,
				},
				resource.Attribute{
					Name:        "instance_no_list",
					Description: `(Optional) List of public IP instance numbers to get.`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `(Optional) List of public IP addresses to get.`,
				},
				resource.Attribute{
					Name:        "search_filter_name",
					Description: `(Optional) ` + "`" + `publicIp` + "`" + ` (Public IP) | ` + "`" + `associatedServerName` + "`" + ` (Associated server name)`,
				},
				resource.Attribute{
					Name:        "search_filter_value",
					Description: `(Optional) Filter value to search`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: KR region.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. You can filter the list of public IP instances by zones. All the public IP addresses in the zone of the region will be selected if the filter is not specified. Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sorted_by",
					Description: `(Optional) The column based on which you want to sort the list.`,
				},
				resource.Attribute{
					Name:        "sorting_order",
					Description: `(Optional) Sorting order of the list. ` + "`" + `ascending` + "`" + ` (Ascending) | ` + "`" + `descending` + "`" + ` (Descending) [case insensitive]. Default: ` + "`" + `ascending` + "`" + ` Ascending ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "instance_no",
					Description: `Public IP instance number`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Public IP description`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the public ip`,
				},
				resource.Attribute{
					Name:        "internet_line_type",
					Description: `Internet line type`,
				},
				resource.Attribute{
					Name:        "instance_status_name",
					Description: `Public IP instance status name`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Public IP instance status`,
				},
				resource.Attribute{
					Name:        "instance_operation",
					Description: `Public IP instance operation`,
				},
				resource.Attribute{
					Name:        "kind_type",
					Description: `Public IP kind type`,
				},
				resource.Attribute{
					Name:        "server_instance",
					Description: `Associated server instance`,
				},
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `Associated server instance number`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Associated server name`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the server instance`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "instance_no",
					Description: `Public IP instance number`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Public IP description`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the public ip`,
				},
				resource.Attribute{
					Name:        "internet_line_type",
					Description: `Internet line type`,
				},
				resource.Attribute{
					Name:        "instance_status_name",
					Description: `Public IP instance status name`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Public IP instance status`,
				},
				resource.Attribute{
					Name:        "instance_operation",
					Description: `Public IP instance operation`,
				},
				resource.Attribute{
					Name:        "kind_type",
					Description: `Public IP kind type`,
				},
				resource.Attribute{
					Name:        "server_instance",
					Description: `Associated server instance`,
				},
				resource.Attribute{
					Name:        "server_instance_no",
					Description: `Associated server instance number`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Associated server name`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the server instance`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_regions",
			Category:         "Data Sources",
			ShortDescription: `Get region list`,
			Description: `

Gets a list of available regions.

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			ShortDescription: `Get root password`,
			Description: `

Gets the password of a root account with the server's login key.

~> **Note:** All arguments including the private key will be stored in the raw state as plain-text.
[Read more about sensitive data in state](/docs/state/sensitive-data.html).

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "root_password",
					Description: `password of a root account`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_server_image",
			Category:         "Data Sources",
			ShortDescription: `Get server image product`,
			Description: `

To create a server instance (VM), you should select a server image. This data source get a server image.

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "product_name_regex",
					Description: `(Optional) A regex string to apply to the server image list returned by ncloud.`,
				},
				resource.Attribute{
					Name:        "exclusion_product_code",
					Description: `(Optional) Product code you want to exclude from the list.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `(Optional) Product code you want to view on the list. Use this when searching for 1 product.`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `(Optional) Product type code`,
				},
				resource.Attribute{
					Name:        "platform_type_code_list",
					Description: `(Optional) Values required for identifying platforms in list-type. The available values are as follows: Linux 32Bit(LNX32) | Linux 64Bit(LNX64) | Windows 32Bit(WND32) | Windows 64Bit(WND64) | Ubuntu Desktop 64Bit(UBD64) | Ubuntu Server 64Bit(UBS64)`,
				},
				resource.Attribute{
					Name:        "block_storage_size",
					Description: `(Optional) Block storage size.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: KR region.`,
				},
				resource.Attribute{
					Name:        "infra_resource_detail_type_code",
					Description: `(Optional) infra resource detail type code. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Product name`,
				},
				resource.Attribute{
					Name:        "product_description",
					Description: `Product description`,
				},
				resource.Attribute{
					Name:        "infra_resource_type",
					Description: `Infra resource type code`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `CPU count`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `Base block storage size`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `Platform type code`,
				},
				resource.Attribute{
					Name:        "os_information",
					Description: `OS Information`,
				},
				resource.Attribute{
					Name:        "add_block_storage_size",
					Description: `Additional block storage size`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "product_name",
					Description: `Product name`,
				},
				resource.Attribute{
					Name:        "product_description",
					Description: `Product description`,
				},
				resource.Attribute{
					Name:        "infra_resource_type",
					Description: `Infra resource type code`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `CPU count`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `Base block storage size`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `Platform type code`,
				},
				resource.Attribute{
					Name:        "os_information",
					Description: `OS Information`,
				},
				resource.Attribute{
					Name:        "add_block_storage_size",
					Description: `Additional block storage size`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_server_images",
			Category:         "Data Sources",
			ShortDescription: `Get server image product list`,
			Description: `

To create a server instance (VM), you should select a server image. This data source gets a list of server images.

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "product_name_regex",
					Description: `(Optional) A regex string to apply to the server image list returned by ncloud.`,
				},
				resource.Attribute{
					Name:        "exclusion_product_code",
					Description: `(Optional) Product code you want to exclude from the list.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `(Optional) Product code you want to view on the list. Use this when searching for 1 product.`,
				},
				resource.Attribute{
					Name:        "platform_type_code_list",
					Description: `(Optional) Values required for identifying platforms in list-type. The available values are as follows: Linux 32Bit(LNX32) | Linux 64Bit(LNX64) | Windows 32Bit(WND32) | Windows 64Bit(WND64) | Ubuntu Desktop 64Bit(UBD64) | Ubuntu Server 64Bit(UBS64)`,
				},
				resource.Attribute{
					Name:        "block_storage_size",
					Description: `(Optional) Block storage size.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: KR region.`,
				},
				resource.Attribute{
					Name:        "infra_resource_detail_type_code",
					Description: `(Optional) infra resource detail type code.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save data source after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "server_images",
					Description: `A List of server image product code`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "server_images",
					Description: `A List of server image product code`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_server_product",
			Category:         "Data Sources",
			ShortDescription: `Get server product`,
			Description: `

ou should select a server product (server specification) to create a server instance (VM).
To this end, we provide data source by which you can search a server product.

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "server_image_product_code",
					Description: `(Required) You can get one from ` + "`" + `data ncloud_server_images` + "`" + `. This is a required value, and each available server's specification varies depending on the server image product.`,
				},
				resource.Attribute{
					Name:        "product_name_regex",
					Description: `(Optional) A regex string to apply to the Server Product list returned.`,
				},
				resource.Attribute{
					Name:        "exclusion_product_code",
					Description: `(Optional) Enter a product code to exclude from the list.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `(Optional) Enter a product code to search from the list. Use it for a single search.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: KR region.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. You can decide a zone where servers are created. You can decide which zone the product list will be requested at. default : Select the first Zone in the specific region Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_line_type_code",
					Description: `(Optional) Internet line code. PUBLC(Public), GLBL(Global) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Product name`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type code`,
				},
				resource.Attribute{
					Name:        "product_description",
					Description: `Product description`,
				},
				resource.Attribute{
					Name:        "infra_resource_type",
					Description: `Infra resource type code`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `CPU count`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `Base block storage size`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `Platform type code`,
				},
				resource.Attribute{
					Name:        "os_information",
					Description: `OS Information`,
				},
				resource.Attribute{
					Name:        "add_block_storage_size",
					Description: `Additional block storage size`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "product_name",
					Description: `Product name`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type code`,
				},
				resource.Attribute{
					Name:        "product_description",
					Description: `Product description`,
				},
				resource.Attribute{
					Name:        "infra_resource_type",
					Description: `Infra resource type code`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `CPU count`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Memory size`,
				},
				resource.Attribute{
					Name:        "base_block_storage_size",
					Description: `Base block storage size`,
				},
				resource.Attribute{
					Name:        "platform_type",
					Description: `Platform type code`,
				},
				resource.Attribute{
					Name:        "os_information",
					Description: `OS Information`,
				},
				resource.Attribute{
					Name:        "add_block_storage_size",
					Description: `Additional block storage size`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_server_products",
			Category:         "Data Sources",
			ShortDescription: `Searching a server product list`,
			Description: `

You should select a server product (server specification) to create a server instance (VM).
To this end, we provide data source by which you can search a server product.

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "server_image_product_code",
					Description: `(Required) You can get one from ` + "`" + `data ncloud_server_images` + "`" + `. This is a required value, and each available server's specification varies depending on the server image product.`,
				},
				resource.Attribute{
					Name:        "product_name_regex",
					Description: `(Optional) A regex string to apply to the Server Product list returned.`,
				},
				resource.Attribute{
					Name:        "exclusion_product_code",
					Description: `(Optional) Enter a product code to exclude from the list.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `(Optional) Enter a product code to search from the list. Use it for a single search.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region code. Get available values using the data source ` + "`" + `ncloud_regions` + "`" + `. Default: KR region.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone code. You can decide a zone where servers are created. You can decide which zone the product list will be requested at. default : Select the first Zone in the specific region Get available values using the data source ` + "`" + `ncloud_zones` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_line_type_code",
					Description: `(Optional) Internet line code. PUBLC(Public), GLBL(Global) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "server_products",
					Description: `A List of server product code`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "server_products",
					Description: `A List of server product code`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ncloud_zones",
			Category:         "Data Sources",
			ShortDescription: `Get region list`,
			Description: `

Gets a list of available zones.

`,
			Keywords: []string{},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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

	dataSourcesMap = map[string]Resource{

		"ncloud_access_control_group":  0,
		"ncloud_access_control_groups": 1,
		"ncloud_access_control_rule":   2,
		"ncloud_access_control_rules":  3,
		"ncloud_member_server_image":   4,
		"ncloud_member_server_images":  5,
		"ncloud_nas_volume":            6,
		"ncloud_nas_volumes":           7,
		"ncloud_port_forwarding_rule":  8,
		"ncloud_port_forwarding_rules": 9,
		"ncloud_public_ip":             10,
		"ncloud_regions":               11,
		"ncloud_root_password":         12,
		"ncloud_server_image":          13,
		"ncloud_server_images":         14,
		"ncloud_server_product":        15,
		"ncloud_server_products":       16,
		"ncloud_zones":                 17,
	}
)

func GetDataSource(r string) (*resouce.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs]
}
