package zerotier

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "zerotier_identity",
			Category:         "Resources",
			ShortDescription: `Identity generator for ZeroTier members. Use this provider with others to authenticate and join users to your networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zerotier_member",
			Category:         "Resources",
			ShortDescription: `Manage ZeroTier members and join them to networks`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zerotier_network",
			Category:         "Resources",
			ShortDescription: `Network provider for ZeroTier, allows you to create ZeroTier networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"zerotier_identity": 0,
		"zerotier_member":   1,
		"zerotier_network":  2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
