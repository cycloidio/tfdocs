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
			ShortDescription: `Chrome Policy Schema data source in the Terraform Googleworkspace provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_domain",
			Category:         "Data Sources",
			ShortDescription: `Domain data source in the Terraform Googleworkspace provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_domain_alias",
			Category:         "Data Sources",
			ShortDescription: `Domain Alias data source in the Terraform Googleworkspace provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_group",
			Category:         "Data Sources",
			ShortDescription: `Group data source in the Terraform Googleworkspace provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_group_member",
			Category:         "Data Sources",
			ShortDescription: `Group Member data source in the Terraform Googleworkspace provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_group_settings",
			Category:         "Data Sources",
			ShortDescription: `Group Settings data source in the Terraform Googleworkspace provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_org_unit",
			Category:         "Data Sources",
			ShortDescription: `OrgUnit data source in the Terraform Googleworkspace provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_privileges",
			Category:         "Data Sources",
			ShortDescription: `Privileges data source in the Terraform Googleworkspace provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_role",
			Category:         "Data Sources",
			ShortDescription: `Role data source in the Terraform Googleworkspace provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_schema",
			Category:         "Data Sources",
			ShortDescription: `Schema data source in the Terraform Googleworkspace provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "googleworkspace_user",
			Category:         "Data Sources",
			ShortDescription: `User data source in the Terraform Googleworkspace provider.`,
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
		"googleworkspace_group_settings":       5,
		"googleworkspace_org_unit":             6,
		"googleworkspace_privileges":           7,
		"googleworkspace_role":                 8,
		"googleworkspace_schema":               9,
		"googleworkspace_user":                 10,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
