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
			Type:             "datadog_cloud_workload_security_agent_rules",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about existing Cloud Workload Security Agent Rules for use in other resources.`,
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
			Type:             "datadog_hosts",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about your live hosts in Datadog.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_aws_logs_services",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve all AWS log ready services.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_aws_namespace_rules",
			Category:         "Data Sources",
			ShortDescription: `Provides a Datadog AWS Integration Namespace Rules data source. This can be used to retrieve all available namespace rules for a Datadog-AWS integration.`,
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
			Type:             "datadog_logs_archives_order",
			Category:         "Data Sources",
			ShortDescription: `Get the current order of your logs archives.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_logs_indexes",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to list several existing logs indexes for use in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_logs_indexes_order",
			Category:         "Data Sources",
			ShortDescription: `Get the current order of your log indexes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_logs_pipelines",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to list all existing logs pipelines for use in other resources.`,
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
			Type:             "datadog_monitor_config_policies",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to list existing monitor config policies for use in other resources.`,
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
			Type:             "datadog_roles",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about multiple roles for use in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_rum_application",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve a Datadog RUM Application.`,
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
			Type:             "datadog_sensitive_data_scanner_group_order",
			Category:         "Data Sources",
			ShortDescription: `Provides a Datadog Sensitive Data Scanner Group Order API data source. This can be used to retrieve the order of Datadog Sensitive Data Scanner Groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_sensitive_data_scanner_standard_pattern",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve information about an existing sensitive data scanner standard pattern.`,
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
			Type:             "datadog_synthetics_test",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to retrieve a Datadog Synthetic Test.`,
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

		"datadog_api_key":                                 0,
		"datadog_application_key":                         1,
		"datadog_cloud_workload_security_agent_rules":     2,
		"datadog_dashboard":                               3,
		"datadog_dashboard_list":                          4,
		"datadog_hosts":                                   5,
		"datadog_integration_aws_logs_services":           6,
		"datadog_integration_aws_namespace_rules":         7,
		"datadog_ip_ranges":                               8,
		"datadog_logs_archives_order":                     9,
		"datadog_logs_indexes":                            10,
		"datadog_logs_indexes_order":                      11,
		"datadog_logs_pipelines":                          12,
		"datadog_monitor":                                 13,
		"datadog_monitor_config_policies":                 14,
		"datadog_monitors":                                15,
		"datadog_permissions":                             16,
		"datadog_role":                                    17,
		"datadog_roles":                                   18,
		"datadog_rum_application":                         19,
		"datadog_security_monitoring_filters":             20,
		"datadog_security_monitoring_rules":               21,
		"datadog_sensitive_data_scanner_group_order":      22,
		"datadog_sensitive_data_scanner_standard_pattern": 23,
		"datadog_service_level_objective":                 24,
		"datadog_service_level_objectives":                25,
		"datadog_synthetics_global_variable":              26,
		"datadog_synthetics_locations":                    27,
		"datadog_synthetics_test":                         28,
		"datadog_user":                                    29,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
