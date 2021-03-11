package anxcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_ip_address",
			Category:         "Resources",
			ShortDescription: `The IP address resource allows you to create IP address at Anexia Cloud.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_network_prefix",
			Category:         "Resources",
			ShortDescription: `The network prefix resource allows you to create network prefix at Anexia Cloud.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_tag",
			Category:         "Resources",
			ShortDescription: `The tag resource allows you to create a tag at Anexia Cloud.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_virtual_server",
			Category:         "Resources",
			ShortDescription: `The Virual Server resource allows you to create virtual machines at Anexia Cloud.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_vlan",
			Category:         "Resources",
			ShortDescription: `The VLAN resource allows you to create VLAN at Anexia Cloud.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"anxcloud_ip_address":     0,
		"anxcloud_network_prefix": 1,
		"anxcloud_tag":            2,
		"anxcloud_virtual_server": 3,
		"anxcloud_vlan":           4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
