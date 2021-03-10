package okta

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "okta_app",
			Category:         "Data Sources",
			ShortDescription: `Get an application of any kind from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_metadata_saml",
			Category:         "Data Sources",
			ShortDescription: `Get a SAML application's metadata from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_app_saml",
			Category:         "Data Sources",
			ShortDescription: `Get a SAML application from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_auth_server",
			Category:         "Data Sources",
			ShortDescription: `Get an auth server from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_default_policy",
			Category:         "Data Sources",
			ShortDescription: `Get a Default policy from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_everyone_group",
			Category:         "Data Sources",
			ShortDescription: `Get the ` + "`" + `Everyone` + "`" + ` group from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_group",
			Category:         "Data Sources",
			ShortDescription: `Get a group from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_idp_metadata_saml",
			Category:         "Data Sources",
			ShortDescription: `Get SAML IdP metadata from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_idp_saml",
			Category:         "Data Sources",
			ShortDescription: `Get a SAML IdP from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_policy",
			Category:         "Data Sources",
			ShortDescription: `Get a policy from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_user",
			Category:         "Data Sources",
			ShortDescription: `Get a single users from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_user_profile_mapping_source",
			Category:         "Data Sources",
			ShortDescription: `Get the base user Profile Mapping source or target from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "okta_users",
			Category:         "Data Sources",
			ShortDescription: `Get a list of users from Okta.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"okta_app":                         0,
		"okta_app_metadata_saml":           1,
		"okta_app_saml":                    2,
		"okta_auth_server":                 3,
		"okta_default_policy":              4,
		"okta_everyone_group":              5,
		"okta_group":                       6,
		"okta_idp_metadata_saml":           7,
		"okta_idp_saml":                    8,
		"okta_policy":                      9,
		"okta_user":                        10,
		"okta_user_profile_mapping_source": 11,
		"okta_users":                       12,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
