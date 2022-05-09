package xray

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "xray_license_policy",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"license",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "xray_security_policy",
			Category:         "Policies",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policies",
				"security",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "xray_settings",
			Category:         "Settings",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"settings",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "xray_watch",
			Category:         "Watch",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"watch",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"xray_license_policy":  0,
		"xray_security_policy": 1,
		"xray_settings":        2,
		"xray_watch":           3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
