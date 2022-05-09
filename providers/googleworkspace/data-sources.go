package googleworkspace

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_chrome_policy_schema",
			Category:         "Data Sources",
			ShortDescription: `Chrome Policy Schema data source in the Terraform Googleworkspace provider. Chrome Policy Schema resides under the https://www.googleapis.com/auth/chrome.management.policy client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_domain",
			Category:         "Data Sources",
			ShortDescription: `Domain data source in the Terraform Googleworkspace provider. Domain resides under the https://www.googleapis.com/auth/admin.directory.domain client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_domain_alias",
			Category:         "Data Sources",
			ShortDescription: `Domain Alias data source in the Terraform Googleworkspace provider. Domain Alias resides under the https://www.googleapis.com/auth/admin.directory.domain client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_group",
			Category:         "Data Sources",
			ShortDescription: `Group data source in the Terraform Googleworkspace provider. Group resides under the https://www.googleapis.com/auth/admin.directory.group client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_group_member",
			Category:         "Data Sources",
			ShortDescription: `Group Member data source in the Terraform Googleworkspace provider. Group Member resides under the https://www.googleapis.com/auth/admin.directory.group client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_group_members",
			Category:         "Data Sources",
			ShortDescription: `Group Members data source in the Terraform Googleworkspace provider. Group Members resides under the https://www.googleapis.com/auth/admin.directory.group client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_group_settings",
			Category:         "Data Sources",
			ShortDescription: `Group Settings data source in the Terraform Googleworkspace provider. Group Settings resides under the https://www.googleapis.com/auth/apps.groups.settings client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_org_unit",
			Category:         "Data Sources",
			ShortDescription: `Org Unit data source in the Terraform Googleworkspace provider. Org Unit resides under the https://www.googleapis.com/auth/admin.directory.orgunit client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_privileges",
			Category:         "Data Sources",
			ShortDescription: `Privileges data source in the Terraform Googleworkspace provider. Privileges resides under the https://www.googleapis.com/auth/admin.directory.rolemanagement client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_role",
			Category:         "Data Sources",
			ShortDescription: `Role data source in the Terraform Googleworkspace provider. Role resides under the https://www.googleapis.com/auth/admin.directory.rolemanagement client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_schema",
			Category:         "Data Sources",
			ShortDescription: `Schema data source in the Terraform Googleworkspace provider. Schema resides under the https://www.googleapis.com/auth/admin.directory.userschema client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_user",
			Category:         "Data Sources",
			ShortDescription: `User data source in the Terraform Googleworkspace provider. User resides under the https://www.googleapis.com/auth/admin.directory.user client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_users",
			Category:         "Data Sources",
			ShortDescription: `Users data source in the Terraform Googleworkspace provider. Users resides under the https://www.googleapis.com/auth/admin.directory.user client scope.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"googleworkspace_chrome_policy_schema": 0,
		"googleworkspace_domain":               1,
		"googleworkspace_domain_alias":         2,
		"googleworkspace_group":                3,
		"googleworkspace_group_member":         4,
		"googleworkspace_group_members":        5,
		"googleworkspace_group_settings":       6,
		"googleworkspace_org_unit":             7,
		"googleworkspace_privileges":           8,
		"googleworkspace_role":                 9,
		"googleworkspace_schema":               10,
		"googleworkspace_user":                 11,
		"googleworkspace_users":                12,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
