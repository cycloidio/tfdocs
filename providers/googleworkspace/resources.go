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
			ShortDescription: `Chrome Policy resource in the Terraform Googleworkspace provider. Currently only supports policies not requiring additionalTargetKeys. Chrome Policy Schema resides under the https://www.googleapis.com/auth/chrome.management.policy client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_domain",
			Category:         "Resources",
			ShortDescription: `Domain resource manages Google Workspace Domains. Domain resides under the https://www.googleapis.com/auth/admin.directory.domain client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_domain_alias",
			Category:         "Resources",
			ShortDescription: `Domain Alias resource manages Google Workspace Domain Aliases. Domain Alias resides under the https://www.googleapis.com/auth/admin.directory.domain client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_gmail_send_as_alias",
			Category:         "Resources",
			ShortDescription: `Gmail Send As Alias resource in the Terraform Googleworkspace provider. Please ensure the Gmail API is enabled for your workspace and that the user being configured has a Gmail license. Gmail Send As Alias resides under the https://www.googleapis.com/auth/gmail.settings client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_group",
			Category:         "Resources",
			ShortDescription: `Group resource manages Google Workspace Groups. Group resides under the https://www.googleapis.com/auth/admin.directory.group client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_group_member",
			Category:         "Resources",
			ShortDescription: `Group Member resource manages Google Workspace Groups Members. Group Member resides under the https://www.googleapis.com/auth/admin.directory.group client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_group_members",
			Category:         "Resources",
			ShortDescription: `Group Members resource manages Google Workspace Groups Members. Group Members resides under the https://www.googleapis.com/auth/admin.directory.group client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_group_settings",
			Category:         "Resources",
			ShortDescription: `Group Settings resource manages Google Workspace Groups Setting. Group Settings requires the https://www.googleapis.com/auth/apps.groups.settings client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_org_unit",
			Category:         "Resources",
			ShortDescription: `OrgUnit resource manages Google Workspace OrgUnits. Org Unit resides under the https://www.googleapis.com/auth/admin.directory.orgunit client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_role",
			Category:         "Resources",
			ShortDescription: `Role resource in the Terraform Googleworkspace provider. Role resides under the https://www.googleapis.com/auth/admin.directory.rolemanagement client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_role_assignment",
			Category:         "Resources",
			ShortDescription: `Role Assignment resource in the Terraform Googleworkspace provider. Role Assignment resides under the https://www.googleapis.com/auth/admin.directory.rolemanagement client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_schema",
			Category:         "Resources",
			ShortDescription: `Schema resource manages Google Workspace Schemas. Schema resides under the https://www.googleapis.com/auth/admin.directory.userschema client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_user",
			Category:         "Resources",
			ShortDescription: `User resource manages Google Workspace Users. User resides under the https://www.googleapis.com/auth/admin.directory.user client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"googleworkspace_chrome_policy":       0,
		"googleworkspace_domain":              1,
		"googleworkspace_domain_alias":        2,
		"googleworkspace_gmail_send_as_alias": 3,
		"googleworkspace_group":               4,
		"googleworkspace_group_member":        5,
		"googleworkspace_group_members":       6,
		"googleworkspace_group_settings":      7,
		"googleworkspace_org_unit":            8,
		"googleworkspace_role":                9,
		"googleworkspace_role_assignment":     10,
		"googleworkspace_schema":              11,
		"googleworkspace_user":                12,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
