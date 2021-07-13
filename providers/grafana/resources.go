package grafana

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "grafana_alert_notification",
			Category:         "Resources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana/latest/alerting/notifications/HTTP API https://grafana.com/docs/grafana/latest/http_api/alerting_notification_channels/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_builtin_role_assignment",
			Category:         "Resources",
			ShortDescription: `Note: This resource is available only with Grafana Enterprise 8.+. Official documentation https://grafana.com/docs/grafana/latest/enterprise/access-control/HTTP API https://grafana.com/docs/grafana/latest/http_api/access_control/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_dashboard",
			Category:         "Resources",
			ShortDescription: `Manages Grafana dashboards. Official documentation https://grafana.com/docs/grafana/latest/dashboards/HTTP API https://grafana.com/docs/grafana/latest/http_api/dashboard/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_dashboard_permission",
			Category:         "Resources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana/latest/permissions/dashboard_folder_permissions/HTTP API https://grafana.com/docs/grafana/latest/http_api/dashboard_permissions/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_data_source",
			Category:         "Resources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana/latest/datasources/HTTP API https://grafana.com/docs/grafana/latest/http_api/data_source/ The required arguments for this resource vary depending on the type of data source selected (via the 'type' argument).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_folder",
			Category:         "Resources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana/latest/dashboards/dashboard_folders/HTTP API https://grafana.com/docs/grafana/latest/http_api/folder/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_folder_permission",
			Category:         "Resources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana/latest/permissions/dashboard_folder_permissions/HTTP API https://grafana.com/docs/grafana/latest/http_api/folder_permissions/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_organization",
			Category:         "Resources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana/latest/manage-users/server-admin/server-admin-manage-orgs/HTTP API https://grafana.com/docs/grafana/latest/http_api/org/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_role",
			Category:         "Resources",
			ShortDescription: `Note: This resource is available only with Grafana Enterprise 8.+. Official documentation https://grafana.com/docs/grafana/latest/enterprise/access-control/HTTP API https://grafana.com/docs/grafana/latest/http_api/access_control/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_synthetic_monitoring_check",
			Category:         "Resources",
			ShortDescription: `Synthetic Monitoring checks are tests that run on selected probes at defined intervals and report metrics and logs back to your Grafana Cloud account. The target for checks can be a domain name, a server, or a website, depending on what information you would like to gather about your endpoint. You can define multiple checks for a single endpoint to check different capabilities. Official documentation https://grafana.com/docs/grafana-cloud/synthetic-monitoring/checks/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_synthetic_monitoring_probe",
			Category:         "Resources",
			ShortDescription: `Besides the public probes run by Grafana Labs, you can also install your own private probes. These are only accessible to you and only write data to your Grafana Cloud account. Private probes are instances of the open source Grafana Synthetic Monitoring Agent. Official documentation https://grafana.com/docs/grafana-cloud/synthetic-monitoring/private-probes/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_team",
			Category:         "Resources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana/latest/manage-users/manage-teams/HTTP API https://grafana.com/docs/grafana/latest/http_api/team/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_team_external_group",
			Category:         "Resources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana/latest/enterprise/team-sync/HTTP API https://grafana.com/docs/grafana/latest/http_api/external_group_sync/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_team_preferences",
			Category:         "Resources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana/latest/administration/preferences/HTTP API https://grafana.com/docs/grafana/latest/http_api/team/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "grafana_user",
			Category:         "Resources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana/latest/manage-users/server-admin/server-admin-manage-users/HTTP API https://grafana.com/docs/grafana/latest/http_api/user/ This resource uses Grafana's admin APIs for creating and updating users which does not currently work with API Tokens. You must use basic auth.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"grafana_alert_notification":         0,
		"grafana_builtin_role_assignment":    1,
		"grafana_dashboard":                  2,
		"grafana_dashboard_permission":       3,
		"grafana_data_source":                4,
		"grafana_folder":                     5,
		"grafana_folder_permission":          6,
		"grafana_organization":               7,
		"grafana_role":                       8,
		"grafana_synthetic_monitoring_check": 9,
		"grafana_synthetic_monitoring_probe": 10,
		"grafana_team":                       11,
		"grafana_team_external_group":        12,
		"grafana_team_preferences":           13,
		"grafana_user":                       14,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
