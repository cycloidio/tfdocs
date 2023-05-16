package devcycle

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "devcycle_environment",
			Category:         "Resources",
			ShortDescription: `DevCycle Environment resource. This resource is used to create and manage DevCycle environments.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "devcycle_feature",
			Category:         "Resources",
			ShortDescription: `DevCycle Feature resource. It's recommended to use the variable resource instead of this resource to manage variables. Variations currently have to be managed in this resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "devcycle_project",
			Category:         "Resources",
			ShortDescription: `DevCycle project resource. Allows for creation/modification of a project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "devcycle_variable",
			Category:         "Resources",
			ShortDescription: `DevCycle Variable resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"devcycle_environment": 0,
		"devcycle_feature":     1,
		"devcycle_project":     2,
		"devcycle_variable":    3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
