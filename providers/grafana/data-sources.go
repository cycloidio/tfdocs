package grafana

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "grafana_cloud_stack",
			Category:         "Data Sources",
			ShortDescription: `Data source for Grafana Stack`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_dashboard",
			Category:         "Data Sources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana/latest/dashboards/Folder/Dashboard Search HTTP API https://grafana.com/docs/grafana/latest/http_api/folder_dashboard_search/Dashboard HTTP API https://grafana.com/docs/grafana/latest/http_api/dashboard/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_dashboards",
			Category:         "Data Sources",
			ShortDescription: `Datasource for retrieving all dashboards. Specify list of folder IDs to search in for dashboards. Official documentation https://grafana.com/docs/grafana/latest/dashboards/Folder/Dashboard Search HTTP API https://grafana.com/docs/grafana/latest/http_api/folder_dashboard_search/Dashboard HTTP API https://grafana.com/docs/grafana/latest/http_api/dashboard/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_folder",
			Category:         "Data Sources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana/latest/dashboards/dashboard_folders/HTTP API https://grafana.com/docs/grafana/latest/http_api/folder/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_library_panel",
			Category:         "Data Sources",
			ShortDescription: `Data source for retrieving a single library panel by name or uid.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_oncall_action",
			Category:         "Data Sources",
			ShortDescription: `HTTP API https://grafana.com/docs/grafana-cloud/oncall/oncall-api-reference/outgoing_webhooks/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_oncall_escalation_chain",
			Category:         "Data Sources",
			ShortDescription: `HTTP API https://grafana.com/docs/grafana-cloud/oncall/oncall-api-reference/escalation_chains/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_oncall_schedule",
			Category:         "Data Sources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana-cloud/oncall/calendar-schedules/HTTP API https://grafana.com/docs/grafana-cloud/oncall/oncall-api-reference/schedules/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_oncall_slack_channel",
			Category:         "Data Sources",
			ShortDescription: `HTTP API https://grafana.com/docs/grafana-cloud/oncall/oncall-api-reference/slack_channels/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_oncall_team",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_oncall_user",
			Category:         "Data Sources",
			ShortDescription: `HTTP API https://grafana.com/docs/grafana-cloud/oncall/oncall-api-reference/users/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_oncall_user_group",
			Category:         "Data Sources",
			ShortDescription: `HTTP API https://grafana.com/docs/grafana-cloud/oncall/oncall-api-reference/user_groups/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_synthetic_monitoring_probe",
			Category:         "Data Sources",
			ShortDescription: `Data source for retrieving a single probe by name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_synthetic_monitoring_probes",
			Category:         "Data Sources",
			ShortDescription: `Data source for retrieving all probes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_user",
			Category:         "Data Sources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana/latest/administration/manage-users-and-permissions/manage-server-users/HTTP API https://grafana.com/docs/grafana/latest/http_api/user/ This data source uses Grafana's admin APIs for reading users which does not currently work with API Tokens. You must use basic auth.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"grafana_cloud_stack":                 0,
		"grafana_dashboard":                   1,
		"grafana_dashboards":                  2,
		"grafana_folder":                      3,
		"grafana_library_panel":               4,
		"grafana_oncall_action":               5,
		"grafana_oncall_escalation_chain":     6,
		"grafana_oncall_schedule":             7,
		"grafana_oncall_slack_channel":        8,
		"grafana_oncall_team":                 9,
		"grafana_oncall_user":                 10,
		"grafana_oncall_user_group":           11,
		"grafana_synthetic_monitoring_probe":  12,
		"grafana_synthetic_monitoring_probes": 13,
		"grafana_user":                        14,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
