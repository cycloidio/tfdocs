package upcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "upcloud_firewall_rules",
			Category:         "Resources",
			ShortDescription: `Manages UpCloud firewall rules list`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_floating_ip_address",
			Category:         "Resources",
			ShortDescription: `Manages UpCloud Floating IP Addresses`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_network",
			Category:         "Resources",
			ShortDescription: `Manages UpCloud networks`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_object_storage",
			Category:         "Resources",
			ShortDescription: `Manages Object Storage bucket resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_router",
			Category:         "Resources",
			ShortDescription: `Manages Upcloud router resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_server",
			Category:         "Resources",
			ShortDescription: `Manages UpCloud Servers`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_storage",
			Category:         "Resources",
			ShortDescription: `Manages UpCloud Storage devices`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_tag",
			Category:         "Resources",
			ShortDescription: `Manages Tags (deprecated)`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"upcloud_firewall_rules":      0,
		"upcloud_floating_ip_address": 1,
		"upcloud_network":             2,
		"upcloud_object_storage":      3,
		"upcloud_router":              4,
		"upcloud_server":              5,
		"upcloud_storage":             6,
		"upcloud_tag":                 7,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
