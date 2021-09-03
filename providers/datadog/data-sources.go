package datadog

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "datadog_api_key",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about an existing api key.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_application_key",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about an existing application key.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_dashboard",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about an existing dashboard, for use in other resources. In particular, it can be used in a monitor message to link to a specific dashboard.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_dashboard_list",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about an existing dashboard list, for use in other resources. In particular, it can be used in a dashboard to register it in the list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about Datadog's IP addresses.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_monitor",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about an existing monitor for use in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_monitors",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to list several existing monitors for use in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_permissions",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve the list of Datadog permissions by name and their corresponding ID, for use in the role resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_role",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about an existing role for use in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_security_monitoring_filters",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about existing security monitoring filters for use in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_security_monitoring_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about existing security monitoring rules for use in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_service_level_objective",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about an existing SLO for use in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_service_level_objectives",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about multiple SLOs for use in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_synthetics_global_variable",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve a Datadog Synthetics global variable (to be used in Synthetics tests).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_synthetics_locations",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve Datadog's Synthetics Locations (to be used in Synthetics tests).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_user",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about an existing user to use it in an other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"datadog_api_key":                     0,
		"datadog_application_key":             1,
		"datadog_dashboard":                   2,
		"datadog_dashboard_list":              3,
		"datadog_ip_ranges":                   4,
		"datadog_monitor":                     5,
		"datadog_monitors":                    6,
		"datadog_permissions":                 7,
		"datadog_role":                        8,
		"datadog_security_monitoring_filters": 9,
		"datadog_security_monitoring_rules":   10,
		"datadog_service_level_objective":     11,
		"datadog_service_level_objectives":    12,
		"datadog_synthetics_global_variable":  13,
		"datadog_synthetics_locations":        14,
		"datadog_user":                        15,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
