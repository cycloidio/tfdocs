package zenlayercloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_ddos_ips",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query DDoS IP instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "associated_instance_id",
					Description: `(Optional, String) The ID of instance to bind with DDoS IPs to be queried.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, String) The ID of zone that the DDoS IPs locates at.`,
				},
				resource.Attribute{
					Name:        "ip_ids",
					Description: `(Optional, Set: [` + "`" + `String` + "`" + `]) IDs of the DDoS IP to be queried.`,
				},
				resource.Attribute{
					Name:        "ip_status",
					Description: `(Optional, String) The status of elastic ip to be queried.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional, String) The address of elastic ip to be queried.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, String) The ID of resource group grouped instances to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, String) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ip_list",
					Description: `An information list of DDoS IP. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of zone that the DDoS IP locates at.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS IP.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the DDoS IP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance id to bind with the DDoS IP.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The instance name to bind with the DDoS IP.`,
				},
				resource.Attribute{
					Name:        "ip_charge_type",
					Description: `The charge type of DDoS IP.`,
				},
				resource.Attribute{
					Name:        "ip_id",
					Description: `ID of the DDoS IP.`,
				},
				resource.Attribute{
					Name:        "ip_status",
					Description: `Current status of the DDoS IP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The elastic ip address.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of resource group grouped instances to be queried.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of resource group grouped instances to be queried.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_list",
					Description: `An information list of DDoS IP. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of zone that the DDoS IP locates at.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the DDoS IP.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the DDoS IP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance id to bind with the DDoS IP.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The instance name to bind with the DDoS IP.`,
				},
				resource.Attribute{
					Name:        "ip_charge_type",
					Description: `The charge type of DDoS IP.`,
				},
				resource.Attribute{
					Name:        "ip_id",
					Description: `ID of the DDoS IP.`,
				},
				resource.Attribute{
					Name:        "ip_status",
					Description: `Current status of the DDoS IP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The elastic ip address.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of resource group grouped instances to be queried.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of resource group grouped instances to be queried.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_eips",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query eip instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "associated_instance_id",
					Description: `(Optional, String) The ID of instance to bind with EIPs to be queried.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, String) The ID of zone that the EIPs locates at.`,
				},
				resource.Attribute{
					Name:        "eip_ids",
					Description: `(Optional, Set: [` + "`" + `String` + "`" + `]) IDs of the EIP to be queried.`,
				},
				resource.Attribute{
					Name:        "eip_status",
					Description: `(Optional, String) The status of elastic ip to be queried.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional, String) The address of elastic ip to be queried.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, String) The ID of resource group grouped instances to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, String) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "eip_list",
					Description: `An information list of EIP. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of zone that the EIP locates at.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the EIP.`,
				},
				resource.Attribute{
					Name:        "eip_charge_type",
					Description: `The charge type of EIP.`,
				},
				resource.Attribute{
					Name:        "eip_id",
					Description: `ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "eip_status",
					Description: `Current status of the EIP.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the EIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance id to bind with the EIP.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The instance name to bind with the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The elastic ip address.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of resource group grouped instances to be queried.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of resource group grouped instances to be queried.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "eip_list",
					Description: `An information list of EIP. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of zone that the EIP locates at.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the EIP.`,
				},
				resource.Attribute{
					Name:        "eip_charge_type",
					Description: `The charge type of EIP.`,
				},
				resource.Attribute{
					Name:        "eip_id",
					Description: `ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "eip_status",
					Description: `Current status of the EIP.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the EIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance id to bind with the EIP.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The instance name to bind with the EIP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The elastic ip address.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of resource group grouped instances to be queried.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of resource group grouped instances to be queried.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_images",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query images.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "catalog",
					Description: `(Optional, String) The catalog which the image belongs to. Valid values: 'centos', 'windows', 'ubuntu', 'debian', 'esxi'.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional, String) ID of the image.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional, String) Name of the image, such as ` + "`" + `CentOS7.4-x86_64` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `(Optional, String) The image type. Valid values: 'PUBLIC_IMAGE', 'CUSTOM_IMAGE'.`,
				},
				resource.Attribute{
					Name:        "instance_type_id",
					Description: `(Optional, String) Filter images which are supported to install on specified instance type, such as ` + "`" + `M6C` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional, String) os type of the image. Valid values: 'windows', 'linux'.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, String) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `An information list of image. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `Created time of the image.`,
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
					Name:        "image_type",
					Description: `Type of the image. with value: ` + "`" + `PUBLIC_IMAGE` + "`" + ` and ` + "`" + `CUSTOM_IMAGE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `Type of the image, windows or linux.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "images",
					Description: `An information list of image. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `Created time of the image.`,
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
					Name:        "image_type",
					Description: `Type of the image. with value: ` + "`" + `PUBLIC_IMAGE` + "`" + ` and ` + "`" + `CUSTOM_IMAGE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `Type of the image, windows or linux.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_instance_types",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query instances types.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, String) The available zone that the BMC instance locates at.`,
				},
				resource.Attribute{
					Name:        "exclude_sold_out",
					Description: `(Optional, Bool) Indicate to filter instances types that is sold out or not, default is false.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, String) The charge type of instance. Valid values are ` + "`" + `POSTPAID` + "`" + `, ` + "`" + `PREPAID` + "`" + `. The default is ` + "`" + `POSTPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_type_id",
					Description: `(Optional, String) The instance type id of the instance.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, String) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `An information list of available bmc instance types. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The zone id that the bmc instance locates at.`,
				},
				resource.Attribute{
					Name:        "default_traffic_package_size",
					Description: `The default value of traffic package size.`,
				},
				resource.Attribute{
					Name:        "instance_type_id",
					Description: `Type ID of the instance.`,
				},
				resource.Attribute{
					Name:        "internet_charge_types",
					Description: `The supported internet charge types of the instance at specified zone.`,
				},
				resource.Attribute{
					Name:        "maximum_bandwidth_out",
					Description: `The maximum public bandwidth of the instance type.`,
				},
				resource.Attribute{
					Name:        "sell_status",
					Description: `Sell status of the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_types",
					Description: `An information list of available bmc instance types. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The zone id that the bmc instance locates at.`,
				},
				resource.Attribute{
					Name:        "default_traffic_package_size",
					Description: `The default value of traffic package size.`,
				},
				resource.Attribute{
					Name:        "instance_type_id",
					Description: `Type ID of the instance.`,
				},
				resource.Attribute{
					Name:        "internet_charge_types",
					Description: `The supported internet charge types of the instance at specified zone.`,
				},
				resource.Attribute{
					Name:        "maximum_bandwidth_out",
					Description: `The maximum public bandwidth of the instance type.`,
				},
				resource.Attribute{
					Name:        "sell_status",
					Description: `Sell status of the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_instances",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query bmc instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, String) The ID of zone that the bmc instance locates at.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional, String) The hostname of the instance to be queried.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional, String) The image of the instance to be queried.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `(Optional, Set: [` + "`" + `String` + "`" + `]) IDs of the instances to be queried.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional, String) Name of the instances to be queried.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `(Optional, String) Status of the instances to be queried.`,
				},
				resource.Attribute{
					Name:        "instance_type_id",
					Description: `(Optional, String) Instance type, such as ` + "`" + `M6C` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_ipv4",
					Description: `(Optional, String) The private ip of the instances to be queried.`,
				},
				resource.Attribute{
					Name:        "public_ipv4",
					Description: `(Optional, String) The public ipv4 of the instances to be queried.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, String) The ID of resource group that the instance grouped by.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, String) Used to save results.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, String) The ID of vpc subnetwork. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_list",
					Description: `An information list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of zone that the bmc instance locates at.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the instance.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname of the instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of image to use for the instance.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The image name to use for the instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_prepaid_period",
					Description: `The tenancy (time unit is month) of the prepaid instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `The charge type of instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instances.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The name of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Current status of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type_id",
					Description: `The type of the instance.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second).`,
				},
				resource.Attribute{
					Name:        "nic_lan_name",
					Description: `The lan name of the nic. The lan name should be a combination of 4 to 10 characters comprised of letters (case insensitive), numbers. The lan name must start with letter. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "nic_wan_name",
					Description: `The wan name of the nic. The wan name should be a combination of 4 to 10 characters comprised of letters (case insensitive), numbers. The wan name must start with letter. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "partitions",
					Description: `Partition for the instance. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "fs_path",
					Description: `The drive letter(windows) or device name(linux) for the partition.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `The type of the partitioned file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the partitioned disk.`,
				},
				resource.Attribute{
					Name:        "private_ipv4_addresses",
					Description: `Private Ipv4 addresses of the instance.`,
				},
				resource.Attribute{
					Name:        "public_ipv4_addresses",
					Description: `Public Ipv4 addresses of the instance.`,
				},
				resource.Attribute{
					Name:        "public_ipv6_addresses",
					Description: `Public Ipv6 addresses of the instance.`,
				},
				resource.Attribute{
					Name:        "raid_config_custom",
					Description: `Custom config for instance raid. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "disk_sequence",
					Description: `The sequence of disk to make raid.`,
				},
				resource.Attribute{
					Name:        "raid_type",
					Description: `Simple config for raid.`,
				},
				resource.Attribute{
					Name:        "raid_config_type",
					Description: `Simple config for instance raid. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of resource group that the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of resource group that the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "traffic_package_size",
					Description: `Traffic package size.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_list",
					Description: `An information list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of zone that the bmc instance locates at.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the instance.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Expired time of the instance.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname of the instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of image to use for the instance.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The image name to use for the instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_prepaid_period",
					Description: `The tenancy (time unit is month) of the prepaid instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `The charge type of instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instances.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The name of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Current status of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type_id",
					Description: `The type of the instance.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second).`,
				},
				resource.Attribute{
					Name:        "nic_lan_name",
					Description: `The lan name of the nic. The lan name should be a combination of 4 to 10 characters comprised of letters (case insensitive), numbers. The lan name must start with letter. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "nic_wan_name",
					Description: `The wan name of the nic. The wan name should be a combination of 4 to 10 characters comprised of letters (case insensitive), numbers. The wan name must start with letter. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "partitions",
					Description: `Partition for the instance. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "fs_path",
					Description: `The drive letter(windows) or device name(linux) for the partition.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `The type of the partitioned file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the partitioned disk.`,
				},
				resource.Attribute{
					Name:        "private_ipv4_addresses",
					Description: `Private Ipv4 addresses of the instance.`,
				},
				resource.Attribute{
					Name:        "public_ipv4_addresses",
					Description: `Public Ipv4 addresses of the instance.`,
				},
				resource.Attribute{
					Name:        "public_ipv6_addresses",
					Description: `Public Ipv6 addresses of the instance.`,
				},
				resource.Attribute{
					Name:        "raid_config_custom",
					Description: `Custom config for instance raid. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "disk_sequence",
					Description: `The sequence of disk to make raid.`,
				},
				resource.Attribute{
					Name:        "raid_type",
					Description: `Simple config for raid.`,
				},
				resource.Attribute{
					Name:        "raid_config_type",
					Description: `Simple config for instance raid. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of resource group that the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of resource group that the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "traffic_package_size",
					Description: `Traffic package size.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_subnets",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query vpc subnets information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, String) Zone of the subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional, String) Filter subnet with this CIDR.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, String) The ID of resource group grouped subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, String) Used to save results.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, String) ID of the subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `(Optional, String) Name of the subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, String) ID of the VPC to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "subnet_list",
					Description: `An information list of subnet. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone of the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of the subnet.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the subnet.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of resource group grouped subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of resource group grouped subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `Name of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_status",
					Description: `Status of the subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Name of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "subnet_list",
					Description: `An information list of subnet. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone of the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of the subnet.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the subnet.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of resource group grouped subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of resource group grouped subnet to be queried.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `Name of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_status",
					Description: `Status of the subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Name of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_vpc_regions",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get the available regions for vpc.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, String) The zone that the vpc region contains.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) The region that the vpc locates at.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, String) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `An information list of vpc regions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `The zones that the vpc region contains.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the region.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `An information list of vpc regions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `The zones that the vpc region contains.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the region.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_vpcs",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to query vpc information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional, String) Filter VPC with this CIDR.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) region of the VPC to be queried.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, String) The ID of resource group grouped VPC to be queried.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, String) Used to save results.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, String) ID of the VPC to be queried. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpc_list",
					Description: `An information list of VPC. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of the VPC.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPC.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the VPC located.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of resource group grouped VPC to be queried.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of resource group grouped VPC to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_status",
					Description: `VPC of the subnet.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_list",
					Description: `An information list of VPC. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `A network address block of the VPC.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VPC.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the VPC located.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of resource group grouped VPC to be queried.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The Name of resource group grouped VPC to be queried.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_status",
					Description: `VPC of the subnet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_zones",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get all bmc available zones.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, String) A regex string to apply to the zone list returned.`,
				},
				resource.Attribute{
					Name:        "result_output_file",
					Description: `(Optional, String) Used to save results. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `An information list of availability zone. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The name of the zone, like ` + "`" + `SEL Zone A` + "`" + `, usually not used in api parameter.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `ID of the zone, such as ` + "`" + `SEL-A` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "zones",
					Description: `An information list of availability zone. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The name of the zone, like ` + "`" + `SEL Zone A` + "`" + `, usually not used in api parameter.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `ID of the zone, such as ` + "`" + `SEL-A` + "`" + `.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"zenlayercloud_bmc_ddos_ips":       0,
		"zenlayercloud_bmc_eips":           1,
		"zenlayercloud_bmc_images":         2,
		"zenlayercloud_bmc_instance_types": 3,
		"zenlayercloud_bmc_instances":      4,
		"zenlayercloud_bmc_subnets":        5,
		"zenlayercloud_bmc_vpc_regions":    6,
		"zenlayercloud_bmc_vpcs":           7,
		"zenlayercloud_bmc_zones":          8,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
