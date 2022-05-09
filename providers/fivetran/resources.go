package fivetran

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fivetran_connector",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

This resource allows you to create, update, and delete connectors.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fivetran_destination",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

This resource allows you to create, update, and delete destinations.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fivetran_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

This resource allows you to create, update, and delete groups.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fivetran_user",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

This resource allows you to create, update, and delete users.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"fivetran_connector":   0,
		"fivetran_destination": 1,
		"fivetran_group":       2,
		"fivetran_user":        3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
