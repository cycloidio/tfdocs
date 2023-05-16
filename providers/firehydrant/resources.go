package firehydrant

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_environment",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name the environment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the environment. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the environment. ## Import Environments can be imported; use ` + "`" + `<ENVIRONMENT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_environment.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the environment. ## Import Environments can be imported; use ` + "`" + `<ENVIRONMENT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_environment.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_functionality",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the functionality.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the functionality.`,
				},
				resource.Attribute{
					Name:        "service_ids",
					Description: `(Optional) A set of IDs of the services this functionality is associated with.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the service. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the functionality. ## Import Functionalities can be imported; use ` + "`" + `<FUNCTIONALITY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_functionality.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the functionality. ## Import Functionalities can be imported; use ` + "`" + `<FUNCTIONALITY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_functionality.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_incident_role",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the incident role.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `(Required) A summary of the incident role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the incident role. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the incident role. ## Import Incident roles can be imported; use ` + "`" + `<INCIDENT ROLE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_incident_role.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the incident role. ## Import Incident roles can be imported; use ` + "`" + `<INCIDENT ROLE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_incident_role.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_priority",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The slug representing the priority. It must be unique and only contain alphanumeric characters. The slug cannot be longer than 23 characters.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) Indicates whether the priority should be the default priority for incidents. At most one resource can have default set to ` + "`" + `true` + "`" + `. Setting default to ` + "`" + `true` + "`" + ` for multiple priority resources will result in inconsistent plans in Terraform. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the priority. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the priority. This is the same as the slug. ## Import Priorities can be imported; use ` + "`" + `<PRIORITY SLUG>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_priority.test P1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the priority. This is the same as the slug. ## Import Priorities can be imported; use ` + "`" + `<PRIORITY SLUG>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_priority.test P1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_runbook",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the runbook.`,
				},
				resource.Attribute{
					Name:        "steps",
					Description: `(Required) Steps to add to the runbook.`,
				},
				resource.Attribute{
					Name:        "attachment_rule",
					Description: `(Optional) JSON string representing the attachment rule configuration for the runbook. Use [Terraform's jsonencode](https://www.terraform.io/language/functions/jsonencode) function so that [Terraform can guarantee valid JSON syntax](https://www.terraform.io/language/expressions/strings#generating-json-or-yaml). For more information on the conditional logic used in ` + "`" + `attachment_rule` + "`" + `, see the [Runbooks - Conditional Logic](../guides/runbooks_conditional_logic.md) documentation. Defaults to attaching manually: ` + "`" + `` + "`" + `` + "`" + `hcl attachment_rule = jsonencode({ logic = { manually = [ { var = "when_invoked" } ] } user_data = {} }) ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the runbook.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `(Optional) The ID of the team that owns this runbook. The ` + "`" + `steps` + "`" + ` block supports: Available attributes and whether they are available and required varies depending on the specific runbook step in question. See [Runbook Steps Configuration documentation](../guides/runbooks_steps.md) for more detailed documentation on each step.`,
				},
				resource.Attribute{
					Name:        "action_id",
					Description: `(Required) The ID of the runbook action for the step.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the step.`,
				},
				resource.Attribute{
					Name:        "automatic",
					Description: `(Optional) Whether this step should be automatically execute. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Optional/Required) JSON string representing the configuration settings for the step. Use [Terraform's jsonencode](https://www.terraform.io/language/functions/jsonencode) function so that [Terraform can guarantee valid JSON syntax](https://www.terraform.io/language/expressions/strings#generating-json-or-yaml).`,
				},
				resource.Attribute{
					Name:        "repeats",
					Description: `(Optional) Whether this step should repeat. Defaults to ` + "`" + `false` + "`" + `. When this value is ` + "`" + `true` + "`" + `, ` + "`" + `repeats_duration` + "`" + ` _must_ be provided.`,
				},
				resource.Attribute{
					Name:        "repeats_duration",
					Description: `(Optional) How often this step should repeat in ISO8601. Example: PT10M [Format Spec](https://www.digi.com/resources/documentation/digidocs/90001437-13/reference/r_iso_8601_duration_format.htm) This value _must_ be provided if ` + "`" + `repeats` + "`" + ` is ` + "`" + `true` + "`" + `. This value _must not_ be provided if ` + "`" + `repeats` + "`" + ` is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) JSON string representing the rule configuration for the runbook step. For more information on the conditional logic used in ` + "`" + `rule` + "`" + `, see the [Runbooks - Conditional Logic](../guides/runbooks_conditional_logic.md) documentation. The step will default to running manually if ` + "`" + `rule` + "`" + ` is not specified and ` + "`" + `automatic` + "`" + ` and ` + "`" + `repeats` + "`" + ` are both ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the runbook. The ` + "`" + `steps` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "step_id",
					Description: `The ID of the step. ## Import Runbooks can be imported; use ` + "`" + `<RUNBOOK ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_runbook.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the runbook. The ` + "`" + `steps` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "step_id",
					Description: `The ID of the step. ## Import Runbooks can be imported; use ` + "`" + `<RUNBOOK ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_runbook.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_service",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the service.`,
				},
				resource.Attribute{
					Name:        "alert_on_add",
					Description: `(Optional) Indicates if FireHydrant should automatically create an alert based on the integrations set up for this service, if this service is added to an active incident. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_add_responding_team",
					Description: `(Optional) Indicates if FireHydrant should automatically add the responding team if this service is added to an active incident. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the service.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Key-value pairs associated with the service. Useful for supporting searching and filtering of the service catalog.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `(Optional) Links associated with the service`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `(Optional) The ID of the team that owns this service.`,
				},
				resource.Attribute{
					Name:        "service_tier",
					Description: `(Optional) The service tier of this resource - between 1 - 5. Lower values represent higher criticality. Defaults to ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "team_ids",
					Description: `(Optional) A set of IDs of the teams responsible for this service's incident response. The ` + "`" + `links` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "href_url",
					Description: `(Required) The URL to use for the link.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the link. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service. ## Import Services can be imported; use ` + "`" + `<SERVICE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_service.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service. ## Import Services can be imported; use ` + "`" + `<SERVICE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_service.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_service_dependency",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connected_service_id",
					Description: `(Required) The ID of the service to define a dependency for. The ` + "`" + `connected_service_id` + "`" + ` represents a service that is a downstream dependency of the service represented by ` + "`" + `service_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `(Required) The ID of the service to define a dependency for. The ` + "`" + `service_id` + "`" + ` represents a service that is an upstream dependency of the service represented by ` + "`" + `connected_service_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Any notes to add to the service dependency. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service dependency. ## Import Service dependencies can be imported; use ` + "`" + `<SERVICE DEPENDENCY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_service_dependency.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service dependency. ## Import Service dependencies can be imported; use ` + "`" + `<SERVICE DEPENDENCY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_service_dependency.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_severity",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The slug representing the severity. It must be unique and only contain alphanumeric characters. The slug cannot be longer than 23 characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the severity.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the severity. Valid values are ` + "`" + `gameday` + "`" + `, ` + "`" + `maintenance` + "`" + `, and ` + "`" + `unexpected_downtime` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the severity. This is the same as the slug. ## Import Severities can be imported; use ` + "`" + `<SEVERITY SLUG>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_severity.test SEV3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the severity. This is the same as the slug. ## Import Severities can be imported; use ` + "`" + `<SEVERITY SLUG>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_severity.test SEV3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_task_list",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the task list.`,
				},
				resource.Attribute{
					Name:        "task_list_items",
					Description: `(Required) A list of tasks to include in the task list.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the task list. The ` + "`" + `task_list_items` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `(Required) A summary of the task.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the task list. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the task list. ## Import Task Lists can be imported; use ` + "`" + `<TASK LIST ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_task_list.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the task list. ## Import Task Lists can be imported; use ` + "`" + `<TASK LIST ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_task_list.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "firehydrant_team",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the team.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the team.`,
				},
				resource.Attribute{
					Name:        "memberships",
					Description: `(Optional) A resource to tie a schedule or user to a team via a incident role. The ` + "`" + `memberships` + "`" + ` block supports: Either the user_id or schedule_id is required for this block.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) Id of the user.`,
				},
				resource.Attribute{
					Name:        "schedule_id",
					Description: `(Required) Id of the schedule.`,
				},
				resource.Attribute{
					Name:        "default_incident_role_id",
					Description: `(Optional) Incident role to assign the user or schedule. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team. ## Import Teams can be imported; use ` + "`" + `<TEAM ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_team.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team. ## Import Teams can be imported; use ` + "`" + `<TEAM ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import firehydrant_team.test 3638b647-b99c-5051-b715-eda2c912c42e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"firehydrant_environment":        0,
		"firehydrant_functionality":      1,
		"firehydrant_incident_role":      2,
		"firehydrant_priority":           3,
		"firehydrant_runbook":            4,
		"firehydrant_service":            5,
		"firehydrant_service_dependency": 6,
		"firehydrant_severity":           7,
		"firehydrant_task_list":          8,
		"firehydrant_team":               9,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
