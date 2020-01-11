package huaweicloudstack

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_images_image_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on a HuaweiCloudStack Image.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Glance client. A Glance client is needed to create an Image that can be used with a compute instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the image.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) The owner (UUID) of the image.`,
				},
				resource.Attribute{
					Name:        "size_min",
					Description: `(Optional) The minimum size (in bytes) of the image to return.`,
				},
				resource.Attribute{
					Name:        "size_max",
					Description: `(Optional) The maximum size (in bytes) of the image to return.`,
				},
				resource.Attribute{
					Name:        "sort_direction",
					Description: `(Optional) Order the results in either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sort_key",
					Description: `(Optional) Sort images based on a certain key. Must be one of "name", "container_format", "disk_format", "status", "id" or "size". Defaults to ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Search for images with a specific tag.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) The visibility of the image. Must be one of "public", "private", "community", or "shared". Defaults to ` + "`" + `private` + "`" + `. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found image. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `The checksum of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `the trailing path after the glance endpoint that represent the location of the image or the path to retrieve it.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `The minimum amount of disk space required to use the image.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `The minimum amount of ram required to use the image.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `Freeform information about the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `Whether or not the image is protected.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `The path to the JSON-schema that represent the image or image`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `The size of the image (in bytes).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `The date the image was last updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "checksum",
					Description: `The checksum of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `the trailing path after the glance endpoint that represent the location of the image or the path to retrieve it.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `The minimum amount of disk space required to use the image.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `The minimum amount of ram required to use the image.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `Freeform information about the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `Whether or not the image is protected.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `The path to the JSON-schema that represent the image or image`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `The size of the image (in bytes).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `The date the image was last updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_network_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an HuaweiCloudStack Network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve networks ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the network.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the network.`,
				},
				resource.Attribute{
					Name:        "matching_subnet_cidr",
					Description: `(Optional) The CIDR of a subnet within the network.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the network. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found network. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Specifies whether the network resource can be accessed by any tenant or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Specifies whether the network resource can be accessed by any tenant or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_port_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information of an HuaweiCloudStack Port.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve port ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The owner of the port.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) The ID of the port.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the port.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the port.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network the port belongs to.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `(Optional) The device owner of the port.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) The MAC address of the port.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Optional) The ID of the device the port belongs to.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) The port IP address filter.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the port.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) The list of port security group IDs to filter. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found port. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The set of security group IDs applied on the port.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The set of security group IDs applied on the port.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_secgroup_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an HuaweiCloudStack Security Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve security groups ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "secgroup_id",
					Description: `(Optional) The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the security group. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found security group. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_subnet_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an HuaweiCloudStack Subnet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve subnet ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the subnet.`,
				},
				resource.Attribute{
					Name:        "dhcp_enabled",
					Description: `(Optional) If the subnet has DHCP enabled.`,
				},
				resource.Attribute{
					Name:        "dhcp_disabled",
					Description: `(Optional) If the subnet has DHCP disabled.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) The IP version of the subnet (either 4 or 6).`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) The IP of the subnet's gateway.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) The CIDR of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network the subnet belongs to.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the subnet. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found subnet. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `Allocation pools of the subnet.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `Whether the subnet has DHCP enabled or not.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `DNS Nameservers of the subnet.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `Host Routes of the subnet.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `Allocation pools of the subnet.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `Whether the subnet has DHCP enabled or not.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `DNS Nameservers of the subnet.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `Host Routes of the subnet.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"huaweicloudstack_images_image_v2":        0,
		"huaweicloudstack_networking_network_v2":  1,
		"huaweicloudstack_networking_port_v2":     2,
		"huaweicloudstack_networking_secgroup_v2": 3,
		"huaweicloudstack_networking_subnet_v2":   4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
