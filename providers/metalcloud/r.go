package metalcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "drive_array",
			Category:         "Resources",
			ShortDescription: `Controls a Metalcloud DriveArray.`,
			Description:      ``,
			Keywords: []string{
				"drive",
				"array",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firewall_rule",
			Category:         "Resources",
			ShortDescription: `Represents a firewall rule that is applied on an instance array.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"rule",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "instance_array",
			Category:         "Resources",
			ShortDescription: `Controls a Metalcloud InstanceArray (collection of servers)`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"array",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_array_label",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "instance_array_instance_count",
					Description: `(Required)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "interface",
			Category:         "Resources",
			ShortDescription: `Controls where an InstanceArray's Instances network interface will be connected.`,
			Description:      ``,
			Keywords: []string{
				"interface",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "network",
			Category:         "Resources",
			ShortDescription: `Controls a Metalcloud network.`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "shared_drive",
			Category:         "Resources",
			ShortDescription: `Controls a Metalcloud DriveArray.`,
			Description:      ``,
			Keywords: []string{
				"shared",
				"drive",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"drive_array":    0,
		"firewall_rule":  1,
		"instance_array": 2,
		"interface":      3,
		"network":        4,
		"shared_drive":   5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
