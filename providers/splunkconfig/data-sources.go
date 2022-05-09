package splunkconfig

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "splunkconfig_app_attributes",
			Category:         "Data Sources",
			ShortDescription: `Get attributes for a specific app`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunkconfig_app_ids",
			Category:         "Data Sources",
			ShortDescription: `Return App IDs from the Splunk Configuration`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunkconfig_app_package",
			Category:         "Data Sources",
			ShortDescription: `Create a tarball for an app.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunkconfig_index_attributes",
			Category:         "Data Sources",
			ShortDescription: `Get attributes for a specific index`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunkconfig_index_names",
			Category:         "Data Sources",
			ShortDescription: `Return Index Names from the Splunk Configuration`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunkconfig_lookup_attributes",
			Category:         "Data Sources",
			ShortDescription: `Get fields and rows for a specific lookup`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunkconfig_role_attributes",
			Category:         "Data Sources",
			ShortDescription: `Get attributes for a specific role`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunkconfig_role_names",
			Category:         "Data Sources",
			ShortDescription: `Return Role Names from the Splunk Configuration`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunkconfig_saml_group_attributes",
			Category:         "Data Sources",
			ShortDescription: `Get attributes for a specific SAML group`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunkconfig_saml_group_names",
			Category:         "Data Sources",
			ShortDescription: `Return SAML Group Names from the Splunk Configuration`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunkconfig_user_attributes",
			Category:         "Data Sources",
			ShortDescription: `Get attributes for a specific user`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "splunkconfig_user_names",
			Category:         "Data Sources",
			ShortDescription: `Return User Names from the Splunk Configuration`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"splunkconfig_app_attributes":        0,
		"splunkconfig_app_ids":               1,
		"splunkconfig_app_package":           2,
		"splunkconfig_index_attributes":      3,
		"splunkconfig_index_names":           4,
		"splunkconfig_lookup_attributes":     5,
		"splunkconfig_role_attributes":       6,
		"splunkconfig_role_names":            7,
		"splunkconfig_saml_group_attributes": 8,
		"splunkconfig_saml_group_names":      9,
		"splunkconfig_user_attributes":       10,
		"splunkconfig_user_names":            11,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
