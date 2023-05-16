package cloudngfwaws

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_app_id_version",
			Category:         "Data Sources",
			ShortDescription: `Data source to retrieve information on a given AppId version.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_app_id_versions",
			Category:         "Data Sources",
			ShortDescription: `Data source get a list of AppId versions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_certificate",
			Category:         "Data Sources",
			ShortDescription: `Data source for retrieving certificate information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_country",
			Category:         "Data Sources",
			ShortDescription: `Data source get a list of countries and their country codes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_custom_url_category",
			Category:         "Data Sources",
			ShortDescription: `Data source for retrieving custom url category information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_fqdn_list",
			Category:         "Data Sources",
			ShortDescription: `Data source for retrieving fqdn list information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_intelligent_feed",
			Category:         "Data Sources",
			ShortDescription: `Data source for retrieving intelligent feed information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_ngfw",
			Category:         "Data Sources",
			ShortDescription: `Data source for retrieving NGFW information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_ngfw_log_profile",
			Category:         "Data Sources",
			ShortDescription: `Data source for retrieving log profile information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_ngfws",
			Category:         "Data Sources",
			ShortDescription: `Data source get a list of NGFWs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_predefined_url_categories",
			Category:         "Data Sources",
			ShortDescription: `Data source for retrieving the predefined URL categories.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_predefined_url_category_override",
			Category:         "Data Sources",
			ShortDescription: `Data source for retrieving a predefined URL category override.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_prefix_list",
			Category:         "Data Sources",
			ShortDescription: `Data source for retrieving prefix list information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_rulestack",
			Category:         "Data Sources",
			ShortDescription: `Data source for retrieving rulestack information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_security_rule",
			Category:         "Data Sources",
			ShortDescription: `Data source for retrieving security rule information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudngfwaws_validate_rulestack",
			Category:         "Data Sources",
			ShortDescription: `Data source to validate the rulestack config.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"cloudngfwaws_app_id_version":                   0,
		"cloudngfwaws_app_id_versions":                  1,
		"cloudngfwaws_certificate":                      2,
		"cloudngfwaws_country":                          3,
		"cloudngfwaws_custom_url_category":              4,
		"cloudngfwaws_fqdn_list":                        5,
		"cloudngfwaws_intelligent_feed":                 6,
		"cloudngfwaws_ngfw":                             7,
		"cloudngfwaws_ngfw_log_profile":                 8,
		"cloudngfwaws_ngfws":                            9,
		"cloudngfwaws_predefined_url_categories":        10,
		"cloudngfwaws_predefined_url_category_override": 11,
		"cloudngfwaws_prefix_list":                      12,
		"cloudngfwaws_rulestack":                        13,
		"cloudngfwaws_security_rule":                    14,
		"cloudngfwaws_validate_rulestack":               15,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
