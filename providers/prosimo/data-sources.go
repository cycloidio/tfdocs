package prosimo

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "prosimo_app_onboarding",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on onboarded applications.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prosimo_certificates",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on existing certificates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prosimo_cloud_creds",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on existing cloud credentials.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prosimo_discovered_networks",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on discovered networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prosimo_edge",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on existing edges.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prosimo_idp",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on available identity providers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prosimo_ip_address",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on available ip ranges.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prosimo_network_onboarding",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on onboarded networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prosimo_policy_access",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on existing access policies.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prosimo_policy_transit",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on existing Transit Policies.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prosimo_s3bucket",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on existing s3 buckets.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"prosimo_app_onboarding":      0,
		"prosimo_certificates":        1,
		"prosimo_cloud_creds":         2,
		"prosimo_discovered_networks": 3,
		"prosimo_edge":                4,
		"prosimo_idp":                 5,
		"prosimo_ip_address":          6,
		"prosimo_network_onboarding":  7,
		"prosimo_policy_access":       8,
		"prosimo_policy_transit":      9,
		"prosimo_s3bucket":            10,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
