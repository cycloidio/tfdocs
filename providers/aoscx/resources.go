package aoscx

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "aoscx_interface",
			Category:         "Resources",
			ShortDescription: `Resource to configure interfaces physical attributes on AOS-CX switches.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aoscx_l2_interface",
			Category:         "Resources",
			ShortDescription: `Resource to configure interface Layer2 attributes on AOS-CX switches.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aoscx_vlan",
			Category:         "Resources",
			ShortDescription: `Resource to configure VLANs on AOS-CX switches.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"aoscx_interface":    0,
		"aoscx_l2_interface": 1,
		"aoscx_vlan":         2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
