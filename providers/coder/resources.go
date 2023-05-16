package coder

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "coder_agent",
			Category:         "Resources",
			ShortDescription: `Use this resource to associate an agent.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coder_agent_instance",
			Category:         "Resources",
			ShortDescription: `Use this resource to associate an instance ID with an agent for zero-trust authentication. This association is done automatically for "googlecomputeinstance", "awsinstance", "azurermlinuxvirtualmachine", and "azurermwindowsvirtual_machine" resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coder_app",
			Category:         "Resources",
			ShortDescription: `Use this resource to define shortcuts to access applications in a workspace.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coder_metadata",
			Category:         "Resources",
			ShortDescription: `Use this resource to attach metadata to a resource. They will be displayed in the Coder dashboard.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"coder_agent":          0,
		"coder_agent_instance": 1,
		"coder_app":            2,
		"coder_metadata":       3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
