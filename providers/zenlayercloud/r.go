package zenlayercloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_ddos_ip",
			Category:         "Bare Metal Cloud(BMC)",
			ShortDescription: `Provides an DDoS IP resource.`,
			Description:      ``,
			Keywords: []string{
				"bare",
				"metal",
				"cloud",
				"bmc",
				"ddos",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, String, ForceNew) The ID of zone that the DDoS IP locates at.`,
				},
				resource.Attribute{
					Name:        "charge_prepaid_period",
					Description: `(Optional, Int, ForceNew) The tenancy (time unit is month) of the prepaid DDoS IP, NOTE: it only works when DDoS charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Optional, String, ForceNew) The charge type of DDoS IP. Valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID` + "`" + `. The default is ` + "`" + `POSTPAID` + "`" + `. Note: ` + "`" + `PREPAID` + "`" + ` DDoS IP may not allow to delete before expired.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional, Bool) Indicate whether to force delete the DDoS IP. Default is ` + "`" + `false` + "`" + `. If set true, the DDoS IP will be permanently deleted instead of being moved into the recycle bin.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, String) The resource group id the DDoS IP belongs to, default to Default Resource Group. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "ip_status",
					Description: `Current status of the DDoS IP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The DDoS IP address.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The resource group name the DDoS IP belongs to, default to Default Resource Group. ## Import EIP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import zenlayercloud_bmc_ddos_ip.foo 123123xxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "ip_status",
					Description: `Current status of the DDoS IP.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The DDoS IP address.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The resource group name the DDoS IP belongs to, default to Default Resource Group. ## Import EIP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import zenlayercloud_bmc_ddos_ip.foo 123123xxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_ddos_ip_association",
			Category:         "Bare Metal Cloud(BMC)",
			ShortDescription: `Provides an DDoS IP resource associated with BMC instance.`,
			Description:      ``,
			Keywords: []string{
				"bare",
				"metal",
				"cloud",
				"bmc",
				"ddos",
				"ip",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ddos_ip_id",
					Description: `(Required, String, ForceNew) The ID of DDoS IP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, String, ForceNew) The instance id going to bind with the DDoS IP. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import DDoS IP association can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import zenlayercloud_bmc_ddos_ip_association.bar ddosIpIdxxxxxx:instanceIdxxxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import DDoS IP association can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import zenlayercloud_bmc_ddos_ip_association.bar ddosIpIdxxxxxx:instanceIdxxxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_eip",
			Category:         "Bare Metal Cloud(BMC)",
			ShortDescription: `Provides an EIP resource.`,
			Description:      ``,
			Keywords: []string{
				"bare",
				"metal",
				"cloud",
				"bmc",
				"eip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, String, ForceNew) The ID of zone that the EIP locates at.`,
				},
				resource.Attribute{
					Name:        "eip_charge_prepaid_period",
					Description: `(Optional, Int, ForceNew) The tenancy (time unit is month) of the prepaid EIP, NOTE: it only works when eip_charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eip_charge_type",
					Description: `(Optional, String, ForceNew) The charge type of EIP. Valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID` + "`" + `. The default is ` + "`" + `POSTPAID` + "`" + `. Note: ` + "`" + `PREPAID` + "`" + ` EIP may not allow to delete before expired.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional, Bool) Indicate whether to force delete the EIP. Default is ` + "`" + `false` + "`" + `. If set true, the EIP will be permanently deleted instead of being moved into the recycle bin.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, String) The resource group id the EIP belongs to, default to Default Resource Group. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the EIP.`,
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
					Name:        "public_ip",
					Description: `The EIP address.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The resource group name the EIP belongs to, default to Default Resource Group. ## Import EIP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import zenlayercloud_bmc_eip.foo 123123xxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the EIP.`,
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
					Name:        "public_ip",
					Description: `The EIP address.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The resource group name the EIP belongs to, default to Default Resource Group. ## Import EIP can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import zenlayercloud_bmc_eip.foo 123123xxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_eip_association",
			Category:         "Bare Metal Cloud(BMC)",
			ShortDescription: `Provides an eip resource associated with BMC instance.`,
			Description:      ``,
			Keywords: []string{
				"bare",
				"metal",
				"cloud",
				"bmc",
				"eip",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "eip_id",
					Description: `(Required, String, ForceNew) The ID of EIP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, String, ForceNew) The instance id going to bind with the EIP. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Eip association can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import zenlayercloud_bmc_eip_association.bar eipIdxxxxxx:instanceIdxxxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource. ## Import Eip association can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import zenlayercloud_bmc_eip_association.bar eipIdxxxxxx:instanceIdxxxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_instance",
			Category:         "Bare Metal Cloud(BMC)",
			ShortDescription: `Provides a BMC instance resource.`,
			Description:      ``,
			Keywords: []string{
				"bare",
				"metal",
				"cloud",
				"bmc",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, String, ForceNew) The ID of zone that the bmc instance locates at.`,
				},
				resource.Attribute{
					Name:        "instance_type_id",
					Description: `(Required, String, ForceNew) The type of the instance.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `(Required, String, ForceNew) Internet charge type of the instance, Valid values are ` + "`" + `ByBandwidth` + "`" + `, ` + "`" + `ByTrafficPackage` + "`" + `, ` + "`" + `ByInstanceBandwidth95` + "`" + ` and ` + "`" + `ByClusterBandwidth95` + "`" + `. This value currently not support to change.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional, Bool) Indicate whether to force delete the instance. Default is ` + "`" + `false` + "`" + `. If set true, the instance will be permanently deleted instead of being moved into the recycle bin.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional, String) The hostname of the instance. The name should be a combination of 2 to 64 characters comprised of letters (case insensitive), numbers, hyphens (-) and Period (.), and the name must be start with letter. The default value is ` + "`" + `Terraform-Instance` + "`" + `. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional, String) The image to use for the instance. Changing ` + "`" + `image_id` + "`" + ` will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "instance_charge_prepaid_period",
					Description: `(Optional, Int) The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to ` + "`" + `PREPAID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, String, ForceNew) The charge type of instance. Valid values are ` + "`" + `PREPAID` + "`" + `, ` + "`" + `POSTPAID` + "`" + `. The default is ` + "`" + `POSTPAID` + "`" + `. Note: ` + "`" + `PREPAID` + "`" + ` instance may not allow to delete before expired.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional, String) The name of the instance. The max length of instance_name is 64, and default value is ` + "`" + `Terraform-Instance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `(Optional, Int) Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second).`,
				},
				resource.Attribute{
					Name:        "nic_lan_name",
					Description: `(Optional, String) The lan name of the nic. The lan name should be a combination of 4 to 10 characters comprised of letters (case insensitive), numbers. The lan name must start with letter. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "nic_wan_name",
					Description: `(Optional, String) The wan name of the nic. The wan name should be a combination of 4 to 10 characters comprised of letters (case insensitive), numbers. The wan name must start with letter. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "partitions",
					Description: `(Optional, List) Partition for the instance. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, String) Password for the instance. The max length of password is 16. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "raid_config_custom",
					Description: `(Optional, List) Custom config for instance raid. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "raid_config_type",
					Description: `(Optional, String) Simple config for instance raid. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, String) The resource group id the instance belongs to, default to Default Resource Group.`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `(Optional, Set: [` + "`" + `String` + "`" + `]) The ssh keys to use for the instance. The max number of ssh keys is 5. Modifying will cause the instance reset.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, String) The ID of a VPC subnet. If you want to create instances in a VPC network, this parameter must be set.`,
				},
				resource.Attribute{
					Name:        "traffic_package_size",
					Description: `(Optional, Float64) Traffic package size. Only valid when the charge type of instance is ` + "`" + `ByTrafficPackage` + "`" + `. The ` + "`" + `partitions` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "fs_path",
					Description: `(Required, String) The drive letter(windows) or device name(linux) for the partition.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Required, String) The type of the partitioned file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required, Int) The size of the partitioned disk. The ` + "`" + `raid_config_custom` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "disk_sequence",
					Description: `(Required, List) The sequence of disk to make raid.`,
				},
				resource.Attribute{
					Name:        "raid_type",
					Description: `(Required, String) Simple config for raid. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "image_name",
					Description: `The image name to use for the instance.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Current status of the instance.`,
				},
				resource.Attribute{
					Name:        "primary_ipv4_address",
					Description: `Primary Ipv4 address of the instance.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `Private Ip addresses of the instance.`,
				},
				resource.Attribute{
					Name:        "public_ipv4_addresses",
					Description: `Public Ipv4 addresses bind to the instance.`,
				},
				resource.Attribute{
					Name:        "public_ipv6_addresses",
					Description: `Public Ipv6 addresses of the instance.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The resource group name the instance belongs to, default to Default Resource Group. ## Import BMC instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import zenlayercloud_bmc_instance.foo 123123xxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
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
					Name:        "image_name",
					Description: `The image name to use for the instance.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `Current status of the instance.`,
				},
				resource.Attribute{
					Name:        "primary_ipv4_address",
					Description: `Primary Ipv4 address of the instance.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `Private Ip addresses of the instance.`,
				},
				resource.Attribute{
					Name:        "public_ipv4_addresses",
					Description: `Public Ipv4 addresses bind to the instance.`,
				},
				resource.Attribute{
					Name:        "public_ipv6_addresses",
					Description: `Public Ipv6 addresses of the instance.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The resource group name the instance belongs to, default to Default Resource Group. ## Import BMC instance can be imported using the id, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import zenlayercloud_bmc_instance.foo 123123xxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_subnet",
			Category:         "Bare Metal Cloud(BMC)",
			ShortDescription: `Provide a resource to create a VPC subnet.`,
			Description:      ``,
			Keywords: []string{
				"bare",
				"metal",
				"cloud",
				"bmc",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required, String, ForceNew) The ID of zone that the bmc subnet locates at.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, String, ForceNew) A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) The name of the bmc subnet.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, String) The resource group id the subnet belongs to, default to Default Resource Group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, String, ForceNew) ID of the VPC to be associated. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the subnet.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The resource group name the subnet belongs to, default to Default Resource Group.`,
				},
				resource.Attribute{
					Name:        "subnet_status",
					Description: `Current status of the subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Name of the VPC to be associated. ## Import Vpc subnet instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import zenlayercloud_bmc_subnet.subnet subnet_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the subnet.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The resource group name the subnet belongs to, default to Default Resource Group.`,
				},
				resource.Attribute{
					Name:        "subnet_status",
					Description: `Current status of the subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Name of the VPC to be associated. ## Import Vpc subnet instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import zenlayercloud_bmc_subnet.subnet subnet_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zenlayercloud_bmc_vpc",
			Category:         "Bare Metal Cloud(BMC)",
			ShortDescription: `Provide a resource to create a VPC.`,
			Description:      ``,
			Keywords: []string{
				"bare",
				"metal",
				"cloud",
				"bmc",
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required, String, ForceNew) A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required, String, ForceNew) The ID of region that the vpc locates at.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) The name of the vpc.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, String) The resource group id the vpc belongs to, default to ID of Default Resource Group. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the vpc.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The resource group name the vpc belongs to, default to Default Resource Group.`,
				},
				resource.Attribute{
					Name:        "vpc_status",
					Description: `Current status of the vpc. ## Import Vpc instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import zenlayercloud_bmc_vpc.test vpc-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Create time of the vpc.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The resource group name the vpc belongs to, default to Default Resource Group.`,
				},
				resource.Attribute{
					Name:        "vpc_status",
					Description: `Current status of the vpc. ## Import Vpc instance can be imported, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import zenlayercloud_bmc_vpc.test vpc-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"zenlayercloud_bmc_ddos_ip":             0,
		"zenlayercloud_bmc_ddos_ip_association": 1,
		"zenlayercloud_bmc_eip":                 2,
		"zenlayercloud_bmc_eip_association":     3,
		"zenlayercloud_bmc_instance":            4,
		"zenlayercloud_bmc_subnet":              5,
		"zenlayercloud_bmc_vpc":                 6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
