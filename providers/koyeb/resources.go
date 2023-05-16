package koyeb

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "koyeb_app",
			Category:         "Resources",
			ShortDescription: `App resource in the Koyeb Terraform provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "koyeb_domain",
			Category:         "Resources",
			ShortDescription: `Domain resource in the Koyeb Terraform provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "koyeb_secret",
			Category:         "Resources",
			ShortDescription: `Secret resource in the Koyeb Terraform provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "koyeb_service",
			Category:         "Resources",
			ShortDescription: `Service resource in the Koyeb Terraform provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"koyeb_app":     0,
		"koyeb_domain":  1,
		"koyeb_secret":  2,
		"koyeb_service": 3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
