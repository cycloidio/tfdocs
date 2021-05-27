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
			Type:             "grafana_dashboard",
			Category:         "Resources",
			ShortDescription: `Official documentation https://grafana.com/docs/grafana/latest/dashboards/HTTP API https://grafana.com/docs/grafana/latest/http_api/dashboard/`,
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

		"grafana_alert_notification":   0,
		"grafana_dashboard":            1,
		"grafana_dashboard_permission": 2,
		"grafana_data_source":          3,
		"grafana_folder":               4,
		"grafana_folder_permission":    5,
		"grafana_organization":         6,
		"grafana_team":                 7,
		"grafana_team_preferences":     8,
		"grafana_user":                 9,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
