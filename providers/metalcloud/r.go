package metalcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

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
			Type:             "instance_array_custom_variables",
			Category:         "Resources",
			ShortDescription: `Represents a set of custom variables that is applied on an instance array.`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"array",
				"custom",
				"variables",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "instance_array_interface",
			Category:         "Resources",
			ShortDescription: `Controls where an InstanceArray's Instances network interface will be connected.`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"array",
				"interface",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "instance_custom_variables",
			Category:         "Resources",
			ShortDescription: `Represents a set of custom variables that is applied on an specific instance of an instance array.`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"custom",
				"variables",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"firewall_rule":                   0,
		"instance_array_custom_variables": 1,
		"instance_array_interface":        2,
		"instance_custom_variables":       3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
