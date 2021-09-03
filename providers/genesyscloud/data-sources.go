package genesyscloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_auth_division",
			Category:         "Data Sources",
			ShortDescription: `Data source for Genesys Cloud Divisions. Select a division by name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_auth_role",
			Category:         "Data Sources",
			ShortDescription: `Data source for Genesys Cloud Roles. Select a role by name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_flow",
			Category:         "Data Sources",
			ShortDescription: `Data source for Genesys Cloud Flows. Select a flow by name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_routing_email_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for Genesys Cloud Email Domains. Select an email domain by name`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_routing_language",
			Category:         "Data Sources",
			ShortDescription: `Data source for Genesys Cloud Routing Languages. Select a language by name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_routing_skill",
			Category:         "Data Sources",
			ShortDescription: `Data source for Genesys Cloud Routing Skills. Select a skill by name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_script",
			Category:         "Data Sources",
			ShortDescription: `Data source for Genesys Cloud Scripts. Select a script by name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "genesyscloud_user",
			Category:         "Data Sources",
			ShortDescription: `Data source for Genesys Cloud Users. Select a user by email or name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"genesyscloud_auth_division":        0,
		"genesyscloud_auth_role":            1,
		"genesyscloud_flow":                 2,
		"genesyscloud_routing_email_domain": 3,
		"genesyscloud_routing_language":     4,
		"genesyscloud_routing_skill":        5,
		"genesyscloud_script":               6,
		"genesyscloud_user":                 7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
