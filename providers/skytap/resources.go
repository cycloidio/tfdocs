package skytap

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "skytap_environment",
			Category:         "Resources",
			ShortDescription: `Provides a Skytap Environment resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skytap_icnr_tunnel",
			Category:         "Resources",
			ShortDescription: `Provides ICNR Tunnel resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skytap_label_category",
			Category:         "Resources",
			ShortDescription: `Provides a Skytap Label Category resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skytap_network",
			Category:         "Resources",
			ShortDescription: `Provides a Skytap Network resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skytap_project",
			Category:         "Resources",
			ShortDescription: `Provides a Skytap Project resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "skytap_vm",
			Category:         "Resources",
			ShortDescription: `Provides a Skytap Virtual Machine (VM) resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"skytap_environment":    0,
		"skytap_icnr_tunnel":    1,
		"skytap_label_category": 2,
		"skytap_network":        3,
		"skytap_project":        4,
		"skytap_vm":             5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
