package ad

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ad_computer",
			Category:         "Resources",
			ShortDescription: `ad_computer manages computer objects in an AD tree.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ad_gplink",
			Category:         "Resources",
			ShortDescription: `ad_gplink manages links between GPOs and container objects such as OUs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ad_gpo",
			Category:         "Resources",
			ShortDescription: `ad_gpo manages Group Policy Objects (GPOs).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ad_gpo_security",
			Category:         "Resources",
			ShortDescription: `ad_gpo_security manages the security settings portion of a Group Policy Object (GPO).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ad_group",
			Category:         "Resources",
			ShortDescription: `ad_group manages Group objects in an Active Directory tree.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ad_group_membership",
			Category:         "Resources",
			ShortDescription: `ad_group_membership manages the members of a given Active Directory group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ad_ou",
			Category:         "Resources",
			ShortDescription: `ad_ou manages OU objects in an AD tree.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ad_user",
			Category:         "Resources",
			ShortDescription: `ad_user manages User objects in an Active Directory tree.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"ad_computer":         0,
		"ad_gplink":           1,
		"ad_gpo":              2,
		"ad_gpo_security":     3,
		"ad_group":            4,
		"ad_group_membership": 5,
		"ad_ou":               6,
		"ad_user":             7,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
