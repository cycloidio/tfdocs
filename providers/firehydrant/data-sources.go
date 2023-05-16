package firehydrant

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_environment",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "environment_id",
					Description: `(Required) The ID of the environment. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the environment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the environment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the environment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the environment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the environment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the environment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_functionality",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "functionality_id",
					Description: `(Required) The ID of the functionality. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the functionality.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the functionality.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the functionality.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `A set of IDs of the services this functionality is associated with.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the functionality.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the functionality.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the functionality.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `A set of IDs of the services this functionality is associated with.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_incident_role",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the incident role. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the incident role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the incident role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the incident role.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `A summary of the incident role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the incident role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the incident role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the incident role.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `A summary of the incident role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_priority",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The slug representing the priority. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the priority. This is the same as the slug.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Indicates whether the priority should be the default priority for incidents.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the priority.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the priority. This is the same as the slug.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Indicates whether the priority should be the default priority for incidents.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the priority.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_runbook",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the runbook. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the runbook.`,
				},
				resource.Attribute{
					Name:        "attachment_rule",
					Description: `JSON string representing the attachment rule configuration for the runbook.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the runbook.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the runbook.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the team that owns this runbook.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the runbook.`,
				},
				resource.Attribute{
					Name:        "attachment_rule",
					Description: `JSON string representing the attachment rule configuration for the runbook.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the runbook.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the runbook.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the team that owns this runbook.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_runbook_action",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "integration_slug",
					Description: `(Required) The slug of the integration associated with the runbook action.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The slug of the runbook action.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of runbook supported for the runbook action. Valid values are ` + "`" + `incident` + "`" + `, ` + "`" + `infrastructure` + "`" + `, and ` + "`" + `incident_role` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the runbook action.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the runbook action.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the runbook action.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the runbook action.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_schedule",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the oncall schedule. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the schedule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the schedule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_service",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the service. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "alert_on_add",
					Description: `Indicates if FireHydrant should automatically create an alert based on the integrations set up for this service, if this service is added to an active incident.`,
				},
				resource.Attribute{
					Name:        "auto_add_responding_team",
					Description: `Indicates if FireHydrant should automatically add the responding team if this service is added to an active incident.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the service.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Key-value pairs associated with the service. Useful for supporting searching and filtering of the service catalog.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Links associated with the service`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the team that owns this service.`,
				},
				resource.Attribute{
					Name:        "service_tier",
					Description: `The service tier of this resource - between 1 - 5. Lower values represent higher criticality.`,
				},
				resource.Attribute{
					Name:        "team_ids",
					Description: `A set of IDs of the teams responsible for this service's incident response. The ` + "`" + `links` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "href_url",
					Description: `The URL for the link.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the link.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "alert_on_add",
					Description: `Indicates if FireHydrant should automatically create an alert based on the integrations set up for this service, if this service is added to an active incident.`,
				},
				resource.Attribute{
					Name:        "auto_add_responding_team",
					Description: `Indicates if FireHydrant should automatically add the responding team if this service is added to an active incident.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the service.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Key-value pairs associated with the service. Useful for supporting searching and filtering of the service catalog.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Links associated with the service`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the team that owns this service.`,
				},
				resource.Attribute{
					Name:        "service_tier",
					Description: `The service tier of this resource - between 1 - 5. Lower values represent higher criticality.`,
				},
				resource.Attribute{
					Name:        "team_ids",
					Description: `A set of IDs of the teams responsible for this service's incident response. The ` + "`" + `links` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "href_url",
					Description: `The URL for the link.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the link.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_services",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels on the runbooks being searched for.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Optional) A query to search for runbooks by their name. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `All the services matching the criteria specified by ` + "`" + `labels` + "`" + ` and ` + "`" + `query` + "`" + `. The ` + "`" + `services` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "alert_on_add",
					Description: `Indicates if FireHydrant should automatically create an alert based on the integrations set up for this service, if this service is added to an active incident.`,
				},
				resource.Attribute{
					Name:        "auto_add_responding_team",
					Description: `Indicates if FireHydrant should automatically add the responding team if this service is added to an active incident.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the service.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Key-value pairs associated with the service. Useful for supporting searching and filtering of the service catalog.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Links associated with the service`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the team that owns this service.`,
				},
				resource.Attribute{
					Name:        "service_tier",
					Description: `The service tier of this resource - between 1 - 5. Lower values represent higher criticality.`,
				},
				resource.Attribute{
					Name:        "team_ids",
					Description: `A set of IDs of the teams responsible for this service's incident response. The ` + "`" + `links` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "href_url",
					Description: `The URL for the link.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the link.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "services",
					Description: `All the services matching the criteria specified by ` + "`" + `labels` + "`" + ` and ` + "`" + `query` + "`" + `. The ` + "`" + `services` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "alert_on_add",
					Description: `Indicates if FireHydrant should automatically create an alert based on the integrations set up for this service, if this service is added to an active incident.`,
				},
				resource.Attribute{
					Name:        "auto_add_responding_team",
					Description: `Indicates if FireHydrant should automatically add the responding team if this service is added to an active incident.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the service.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Key-value pairs associated with the service. Useful for supporting searching and filtering of the service catalog.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Links associated with the service`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of the team that owns this service.`,
				},
				resource.Attribute{
					Name:        "service_tier",
					Description: `The service tier of this resource - between 1 - 5. Lower values represent higher criticality.`,
				},
				resource.Attribute{
					Name:        "team_ids",
					Description: `A set of IDs of the teams responsible for this service's incident response. The ` + "`" + `links` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "href_url",
					Description: `The URL for the link.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the link.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_severity",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The slug representing the severity. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the severity. This is the same as the slug.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the severity.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the severity. Possible values are ` + "`" + `gameday` + "`" + `, ` + "`" + `maintenance` + "`" + `, and ` + "`" + `unexpected_downtime` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the severity. This is the same as the slug.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the severity.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the severity. Possible values are ` + "`" + `gameday` + "`" + `, ` + "`" + `maintenance` + "`" + `, and ` + "`" + `unexpected_downtime` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_task_list",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the task list. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the task list.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the task list.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the task list.`,
				},
				resource.Attribute{
					Name:        "task_list_items",
					Description: `A list of tasks in the task list. The ` + "`" + `task_list_items` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `(Required) A summary of the task.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the task list.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the task list.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the task list.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the task list.`,
				},
				resource.Attribute{
					Name:        "task_list_items",
					Description: `A list of tasks in the task list. The ` + "`" + `task_list_items` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `(Required) A summary of the task.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the task list.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The user's email address. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the user.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"firehydrant_environment":    0,
		"firehydrant_functionality":  1,
		"firehydrant_incident_role":  2,
		"firehydrant_priority":       3,
		"firehydrant_runbook":        4,
		"firehydrant_runbook_action": 5,
		"firehydrant_schedule":       6,
		"firehydrant_service":        7,
		"firehydrant_services":       8,
		"firehydrant_severity":       9,
		"firehydrant_task_list":      10,
		"firehydrant_user":           11,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
