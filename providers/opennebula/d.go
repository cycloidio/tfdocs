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
					Name:        "name",
					Description: `(Required) The OpenNebula group to retrieve information for. ## Attribute Reference The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the group.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `List of User IDs part of the group.`,
				},
				resource.Attribute{
					Name:        "admins",
					Description: `List of Administrator user IDs part of the group.`,
				},
				resource.Attribute{
					Name:        "quotas",
					Description: `Quotas configured for the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the group.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `List of User IDs part of the group.`,
				},
				resource.Attribute{
					Name:        "admins",
					Description: `List of Administrator user IDs part of the group.`,
				},
				resource.Attribute{
					Name:        "quotas",
					Description: `Quotas configured for the group.`,
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
					Name:        "name",
					Description: `(Required) The OpenNebula image to retrieve information for. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags associated to the Image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "tags",
					Description: `Tags associated to the Image.`,
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
					Name:        "name",
					Description: `(Required) The OpenNebula security group to retrieve information for. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Security group tags.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "tags",
					Description: `Security group tags.`,
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
					Name:        "name",
					Description: `(Required) The OpenNebula template to retrieve information for.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `(Optional) Amount of CPU shares assigned to the VM.`,
				},
				resource.Attribute{
					Name:        "vpcu",
					Description: `(Optional) Number of CPU cores presented to the VM.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) Amount of RAM assigned to the VM in MB.`,
				},
				resource.Attribute{
					Name:        "context",
					Description: `(Optional) Array of free form key=value pairs, rendered and added to the CONTEXT variables for the VM. Recommended to include: ` + "`" + `NETWORK = "YES"` + "`" + ` and ` + "`" + `SET_HOSTNAME = "$NAME"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "graphics",
					Description: `(Optional) Graphics parameters.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `(Optional) OS parameters`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional) Disk parameters`,
				},
				resource.Attribute{
					Name:        "nic",
					Description: `(Optional) NIC parameters`,
				},
				resource.Attribute{
					Name:        "vmgroup",
					Description: `(Optional) VM group parameters`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Template tags (Key = Value).`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Deprecated) Text describing the OpenNebula template object, in Opennebula's XML string format.`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Name:        "name",
					Description: `(Required) The OpenNebula user to retrieve information for. ## Attribute Reference The following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password of the user (if set)`,
				},
				resource.Attribute{
					Name:        "auth_driver",
					Description: `Authentication Driver for User management`,
				},
				resource.Attribute{
					Name:        "primary_group",
					Description: `Primary group ID of the User.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `List of secondary groups ID of the user.`,
				},
				resource.Attribute{
					Name:        "quotas",
					Description: `User's quotas`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password of the user (if set)`,
				},
				resource.Attribute{
					Name:        "auth_driver",
					Description: `Authentication Driver for User management`,
				},
				resource.Attribute{
					Name:        "primary_group",
					Description: `Primary group ID of the User.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `List of secondary groups ID of the user.`,
				},
				resource.Attribute{
					Name:        "quotas",
					Description: `User's quotas`,
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
					Name:        "name",
					Description: `(Required) The OpenNebula virtual data center to retrieve information for.`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Name:        "name",
					Description: `(Required) The OpenNebula virtual machine group to retrieve information for.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Virtual Machine group tags.`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Name:        "name",
					Description: `(Required) The OpenNebula virtual network to retrieve information for.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Virtual Network.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) MTU of the Virtual Network.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Virtual Network tags.`,
				},
			},
			Attributes: []resource.Attribute{},
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
