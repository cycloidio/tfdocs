package boundary

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "boundary_account",
			Category:         "Resources",
			ShortDescription: `The account resource allows you to configure a Boundary account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "boundary_auth_method",
			Category:         "Resources",
			ShortDescription: `The auth method resource allows you to configure a Boundary auth_method.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "boundary_group",
			Category:         "Resources",
			ShortDescription: `The group resource allows you to configure a Boundary group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "boundary_host",
			Category:         "Resources",
			ShortDescription: `The host resource allows you to configure a Boundary static host. Hosts are always part of a project, so a project resource should be used inline or you should have the project ID in hand to successfully configure a host.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "boundary_host_catalog",
			Category:         "Resources",
			ShortDescription: `The host catalog resource allows you to configure a Boundary host catalog. Host catalogs are always part of a project, so a project resource should be used inline or you should have the project ID in hand to successfully configure a host catalog.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "boundary_host_set",
			Category:         "Resources",
			ShortDescription: `The host_set resource allows you to configure a Boundary host set. Host sets are always part of a host catalog, so a host catalog resource should be used inline or you should have the host catalog ID in hand to successfully configure a host set.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "boundary_role",
			Category:         "Resources",
			ShortDescription: `The role resource allows you to configure a Boundary role.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "boundary_scope",
			Category:         "Resources",
			ShortDescription: `The scope resource allows you to configure a Boundary scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "boundary_target",
			Category:         "Resources",
			ShortDescription: `The target resource allows you to configure a Boundary target.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "boundary_user",
			Category:         "Resources",
			ShortDescription: `The user resource allows you to configure a Boundary user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"boundary_account":      0,
		"boundary_auth_method":  1,
		"boundary_group":        2,
		"boundary_host":         3,
		"boundary_host_catalog": 4,
		"boundary_host_set":     5,
		"boundary_role":         6,
		"boundary_scope":        7,
		"boundary_target":       8,
		"boundary_user":         9,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
