package googleworkspace

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_chrome_policy",
			Category:         "Resources",
			ShortDescription: `Chrome Policy resource in the Terraform Googleworkspace provider. Currently only supports policies not requiring additionalTargetKeys.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_domain",
			Category:         "Resources",
			ShortDescription: `Domain resource manages Google Workspace Domains.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_domain_alias",
			Category:         "Resources",
			ShortDescription: `Domain Alias resource manages Google Workspace Domain Aliases.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_group",
			Category:         "Resources",
			ShortDescription: `Group resource manages Google Workspace Groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_group_member",
			Category:         "Resources",
			ShortDescription: `Group resource manages Google Workspace Groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_group_settings",
			Category:         "Resources",
			ShortDescription: `Group Settings resource manages Google Workspace Groups Setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_org_unit",
			Category:         "Resources",
			ShortDescription: `OrgUnit resource manages Google Workspace OrgUnits.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_role",
			Category:         "Resources",
			ShortDescription: `Role resource in the Terraform Googleworkspace provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_role_assignment",
			Category:         "Resources",
			ShortDescription: `RoleAssignment resource in the Terraform Googleworkspace provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_schema",
			Category:         "Resources",
			ShortDescription: `Schema resource manages Google Workspace Schemas.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_user",
			Category:         "Resources",
			ShortDescription: `User resource manages Google Workspace Users.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"googleworkspace_chrome_policy":   0,
		"googleworkspace_domain":          1,
		"googleworkspace_domain_alias":    2,
		"googleworkspace_group":           3,
		"googleworkspace_group_member":    4,
		"googleworkspace_group_settings":  5,
		"googleworkspace_org_unit":        6,
		"googleworkspace_role":            7,
		"googleworkspace_role_assignment": 8,
		"googleworkspace_schema":          9,
		"googleworkspace_user":            10,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
