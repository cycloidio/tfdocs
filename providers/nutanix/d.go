package nutanix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nutanix_cluster",
			Category:         "Data Sources",
			ShortDescription: `Describes a Cluster`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The API version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The API version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_clusters",
			Category:         "Data Sources",
			ShortDescription: `Describes a Clusters`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The API version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The API version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_image",
			Category:         "Data Sources",
			ShortDescription: `Describes a Image`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_network_security_rule",
			Category:         "Data Sources",
			ShortDescription: `Describes a Network security rule`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_subnet",
			Category:         "Data Sources",
			ShortDescription: `This operation retrieves a subnet based on the input parameters. A subnet is a block of IP addresses.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_virtual_machine",
			Category:         "Data Sources",
			ShortDescription: `Describes a Virtual Machine`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
				resource.Attribute{
					Name:        "ngt_enabled_capability_list",
					Description: `Application names that are enabled.`,
				},
				resource.Attribute{
					Name:        "guest_customization_cloud_init_meta_data",
					Description: `The contents of the meta_data configuration for cloud-init. This can be formatted as YAML or JSON. The value must be base64 encoded.`,
				},
				resource.Attribute{
					Name:        "disk_size_bytes",
					Description: `Size of the disk in Bytes.`,
				},
				resource.Attribute{
					Name:        "disk_size_mib",
					Description: `Size of the disk in MiB. Must match the size specified in 'disk_size_bytes' - rounded up to the nearest MiB - when that field is present.`,
				},
				resource.Attribute{
					Name:        "device_properties",
					Description: `Properties to a device.`,
				},
				resource.Attribute{
					Name:        "data_source_reference",
					Description: `Reference to a data source.`,
				},
				resource.Attribute{
					Name:        "volume_group_reference",
					Description: `Reference to a volume group. ### Device Properties The device_properties attribute supports the following.`,
				},
				resource.Attribute{
					Name:        "pci_address",
					Description: `GPU {segment:bus:device:function} (sbdf) address if assigned.`,
				},
				resource.Attribute{
					Name:        "fraction",
					Description: `Fraction of the physical GPU assigned.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
				resource.Attribute{
					Name:        "ngt_enabled_capability_list",
					Description: `Application names that are enabled.`,
				},
				resource.Attribute{
					Name:        "guest_customization_cloud_init_meta_data",
					Description: `The contents of the meta_data configuration for cloud-init. This can be formatted as YAML or JSON. The value must be base64 encoded.`,
				},
				resource.Attribute{
					Name:        "disk_size_bytes",
					Description: `Size of the disk in Bytes.`,
				},
				resource.Attribute{
					Name:        "disk_size_mib",
					Description: `Size of the disk in MiB. Must match the size specified in 'disk_size_bytes' - rounded up to the nearest MiB - when that field is present.`,
				},
				resource.Attribute{
					Name:        "device_properties",
					Description: `Properties to a device.`,
				},
				resource.Attribute{
					Name:        "data_source_reference",
					Description: `Reference to a data source.`,
				},
				resource.Attribute{
					Name:        "volume_group_reference",
					Description: `Reference to a volume group. ### Device Properties The device_properties attribute supports the following.`,
				},
				resource.Attribute{
					Name:        "pci_address",
					Description: `GPU {segment:bus:device:function} (sbdf) address if assigned.`,
				},
				resource.Attribute{
					Name:        "fraction",
					Description: `Fraction of the physical GPU assigned.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_volume_group",
			Category:         "Data Sources",
			ShortDescription: `Describes a Image`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) Version of the API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) Version of the API.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"nutanix_cluster":               0,
		"nutanix_clusters":              1,
		"nutanix_image":                 2,
		"nutanix_network_security_rule": 3,
		"nutanix_subnet":                4,
		"nutanix_virtual_machine":       5,
		"nutanix_volume_group":          6,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
