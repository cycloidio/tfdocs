package talos

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "talos_machine_bootstrap",
			Category:         "Resources",
			ShortDescription: `The machine bootstrap resource allows you to bootstrap a Talos node.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "talos_machine_configuration_apply",
			Category:         "Resources",
			ShortDescription: `The machine configuration apply resource allows to apply machine configuration to a node`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "talos_machine_secrets",
			Category:         "Resources",
			ShortDescription: `Generate machine secrets for Talos cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"talos_machine_bootstrap":           0,
		"talos_machine_configuration_apply": 1,
		"talos_machine_secrets":             2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
