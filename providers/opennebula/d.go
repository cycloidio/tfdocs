package opennebula

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "opennebula_group",
			Category:         "Data Sources",
			ShortDescription: `Get the group information for a given name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The OpenNebula group to retrieve information for.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags associated to the Image. ## Attribute Reference The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the group.`,
				},
				resource.Attribute{
					Name:        "admins",
					Description: `List of Administrator user IDs part of the group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the group (Key = Value).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the group.`,
				},
				resource.Attribute{
					Name:        "admins",
					Description: `List of Administrator user IDs part of the group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the group (Key = Value).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_image",
			Category:         "Data Sources",
			ShortDescription: `Get the image information for a given name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The OpenNebula image to retrieve information for.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags associated to the image. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the image.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the image (Key = Value).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the image.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the image (Key = Value).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_security group",
			Category:         "Data Sources",
			ShortDescription: `Get the security group information for a given name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The OpenNebula security group to retrieve information for.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Security group tags. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security group`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the security group (Key = Value).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security group`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the security group (Key = Value).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_template",
			Category:         "Data Sources",
			ShortDescription: `Get the template information for a given name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the template.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The OpenNebula template to retrieve information for.`,
				},
				resource.Attribute{
					Name:        "has_cpu",
					Description: `(Optional) Indicate if a CPU value has been defined.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Optional) Amount of CPU shares assigned to the VM.`,
				},
				resource.Attribute{
					Name:        "has_vcpu",
					Description: `(Optional) Indicate if a VCPU value has been defined.`,
				},
				resource.Attribute{
					Name:        "vcpu",
					Description: `(Optional) Number of CPU cores presented to the VM.`,
				},
				resource.Attribute{
					Name:        "has_memory",
					Description: `(Optional) Indicate if a memory value has been defined.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) Amount of RAM assigned to the VM in MB.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Template tags (Key = Value). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the template.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the template.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Amount of CPU shares assigned to the VM.`,
				},
				resource.Attribute{
					Name:        "vcpu",
					Description: `Number of CPU cores presented to the VM.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Amount of RAM assigned to the VM in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Disk parameters`,
				},
				resource.Attribute{
					Name:        "nic",
					Description: `NIC parameters`,
				},
				resource.Attribute{
					Name:        "vmgroup",
					Description: `VM group parameters`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the template (Key = Value).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the template.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the template.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Amount of CPU shares assigned to the VM.`,
				},
				resource.Attribute{
					Name:        "vcpu",
					Description: `Number of CPU cores presented to the VM.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Amount of RAM assigned to the VM in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Disk parameters`,
				},
				resource.Attribute{
					Name:        "nic",
					Description: `NIC parameters`,
				},
				resource.Attribute{
					Name:        "vmgroup",
					Description: `VM group parameters`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the template (Key = Value).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_user",
			Category:         "Data Sources",
			ShortDescription: `Get the user information for a given name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The OpenNebula user to retrieve information for.`,
				},
				resource.Attribute{
					Name:        "primary_group",
					Description: `(Optional) Primary group ID of the user.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) List of secondary groups ID of the user. ## Attribute Reference The following attribute are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the user.`,
				},
				resource.Attribute{
					Name:        "primary_group",
					Description: `Primary group ID of the user.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `List of secondary groups ID of the user.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the user (Key = Value).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the user.`,
				},
				resource.Attribute{
					Name:        "primary_group",
					Description: `Primary group ID of the user.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `List of secondary groups ID of the user.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the user (Key = Value).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_virtual data center",
			Category:         "Data Sources",
			ShortDescription: `Get the virtual data center information for a given name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the virtual data center.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The OpenNebula virtual data center to retrieve information for.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Virtual data center tags (Key = Value). ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual data center.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the virtual data center.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the virtual data center (Key = Value).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual data center.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the virtual data center.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the virtual data center (Key = Value).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_virtual machine group",
			Category:         "Data Sources",
			ShortDescription: `Get the virtual machine group information for a given name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the virtual machine group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The OpenNebula virtual machine group to retrieve information for.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Virtual Machine group tags (Key = Value). ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual machine group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the virtual machine group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the virtual machine group (Key = Value).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual machine group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the virtual machine group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the virtual machine group (Key = Value).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opennebula_virtual network",
			Category:         "Data Sources",
			ShortDescription: `Get the virtual network information for a given name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the virtual network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The OpenNebula virtual network to retrieve information for.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Virtual network tags (Key = Value). ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the virtual network.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `MTU of the virtual network.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the virtual network (Key = Value).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the virtual network.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `MTU of the virtual network.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags of the virtual network (Key = Value).`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"opennebula_group":                 0,
		"opennebula_image":                 1,
		"opennebula_security group":        2,
		"opennebula_template":              3,
		"opennebula_user":                  4,
		"opennebula_virtual data center":   5,
		"opennebula_virtual machine group": 6,
		"opennebula_virtual network":       7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
