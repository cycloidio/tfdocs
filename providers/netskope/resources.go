package netskope

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "netskope_ipsec_tunnels",
			Category:         "Resources",
			ShortDescription: `IPSec tunnels are not yet supported.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "netskope_privateapps",
			Category:         "Resources",
			ShortDescription: `The privateapps resource is used to create, delete and update Private Applications in a Netskope tenant.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "netskope_publishers",
			Category:         "Resources",
			ShortDescription: `The publishers resource is used to create, delete and update Publishers in a Netskope tenant.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"netskope_ipsec_tunnels": 0,
		"netskope_privateapps":   1,
		"netskope_publishers":    2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
