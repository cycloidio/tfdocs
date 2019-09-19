package nutanix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nutanix_category_key",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Category key resource to Create a category key name.`,
			Description:      ``,
			Keywords: []string{
				"category",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) The version of the API. See detailed information in [Nutanix Image](https://nutanix.github.io/Automation/experimental/swagger-redoc-sandbox/#tag/category/paths/~1categories~1{name}/get).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) The version of the API. See detailed information in [Nutanix Image](https://nutanix.github.io/Automation/experimental/swagger-redoc-sandbox/#tag/category/paths/~1categories~1{name}/get).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_category_value",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Category value resource to Create a category value.`,
			Description:      ``,
			Keywords: []string{
				"category",
				"value",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value for the category value.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) The version of the API. See detailed information in [Nutanix Image](http://developer.nutanix.com/reference/prism_central/v3/#category).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) The version of the API. See detailed information in [Nutanix Image](http://developer.nutanix.com/reference/prism_central/v3/#category).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_image",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Image resource to Create a Image.`,
			Description:      ``,
			Keywords: []string{
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_network_security_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Network Security Rule resource to Create a Network Security Rule .`,
			Description:      ``,
			Keywords: []string{
				"network",
				"security",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_subnet",
			Category:         "Resources",
			ShortDescription: `This operation submits a request to create a subnet based on the input parameters. A subnet is a block of IP addresses.`,
			Description:      ``,
			Keywords: []string{
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API. ### Metadata The metadata attribute exports the following:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_virtual_machine",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Virtual Machine resource to Create a virtual machine.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"machine",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ngt_enabled_capability_list",
					Description: `(Optional) Application names that are enabled.`,
				},
				resource.Attribute{
					Name:        "guest_customization_cloud_init_meta_data",
					Description: `(Optional) The contents of the meta_data configuration for cloud-init. This can be formatted as YAML or JSON. The value must be base64 encoded.`,
				},
				resource.Attribute{
					Name:        "disk_size_bytes",
					Description: `(Optional) Size of the disk in Bytes.`,
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
					Description: `Reference to a volume group. The disk_size (the disk size_mib and the disk_size_bytes attributes) is only honored by creating an empty disk. When you are creating from an image, the size is ignored and the disk becomes the size of the image from which it was cloned. In VM creation, you can't set either disk size_mib or disk_size_bytes when you set data_source_reference but, you can update the disk_size after creation (second apply). ### Device Properties The device_properties attribute supports the following.`,
				},
				resource.Attribute{
					Name:        "pci_address",
					Description: `(ReadOnly) GPU {segment:bus:device:function} (sbdf) address if assigned.`,
				},
				resource.Attribute{
					Name:        "fraction",
					Description: `(ReadOnly) Fraction of the physical GPU assigned.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `The version of the API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanix_volume_group",
			Category:         "Resources",
			ShortDescription: `Provides a Nutanix Virtual Machine resource to Create a volume group.`,
			Description:      ``,
			Keywords: []string{
				"volume",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) Version of the API.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"nutanix_category_key":          0,
		"nutanix_category_value":        1,
		"nutanix_image":                 2,
		"nutanix_network_security_rule": 3,
		"nutanix_subnet":                4,
		"nutanix_virtual_machine":       5,
		"nutanix_volume_group":          6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
