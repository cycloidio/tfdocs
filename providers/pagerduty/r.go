package pagerduty

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_addon",
			Category:         "Resources",
			ShortDescription: `Creates and manages an add-on in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"addon",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the add-on.`,
				},
				resource.Attribute{
					Name:        "src",
					Description: `(Required) The source URL to display in a frame in the PagerDuty UI. ` + "`" + `HTTPS` + "`" + ` is required. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the add-on. ## Import Add-ons can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_addon.example P3DH5M6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the add-on. ## Import Add-ons can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_addon.example P3DH5M6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_escalation_policy",
			Category:         "Resources",
			ShortDescription: `Creates and manages an escalation policy in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"escalation",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the escalation policy.`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `(Optional) Teams associated with the policy. Account must have the ` + "`" + `teams` + "`" + ` ability to use this parameter.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description of the escalation policy. If not set, a placeholder of "Managed by Terraform" will be set.`,
				},
				resource.Attribute{
					Name:        "num_loops",
					Description: `(Optional) The number of times the escalation policy will repeat after reaching the end of its escalation.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) An Escalation rule block. Escalation rules documented below. Escalation rules (` + "`" + `rule` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "escalation_delay_in_minutes",
					Description: `(Required) The number of minutes before an unacknowledged incident escalates away from this rule.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) A target block. Target blocks documented below. Targets (` + "`" + `target` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Can be ` + "`" + `user` + "`" + `, ` + "`" + `schedule` + "`" + `, ` + "`" + `user_reference` + "`" + ` or ` + "`" + `schedule_reference` + "`" + `. Defaults to ` + "`" + `user_reference` + "`" + ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) A target ID ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the escalation policy. ## Import Escalation policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_escalation_policy.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the escalation policy. ## Import Escalation policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_escalation_policy.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_extension",
			Category:         "Resources",
			ShortDescription: `Creates and manages a service extension in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"extension",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the service extension.`,
				},
				resource.Attribute{
					Name:        "endpoint_url",
					Description: `(Optional) The url of the extension.`,
				},
				resource.Attribute{
					Name:        "extension_schema",
					Description: `(Required) This is the schema for this extension.`,
				},
				resource.Attribute{
					Name:        "extension_objects",
					Description: `(Required) This is the objects for which the extension applies (An array of service ids).`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Optional) The configuration of the service extension as string containing plain JSON-encoded data.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the extension.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `a URL at which the entity is uniquely displayed in the Web app ## Import Extensions can be imported using the id.e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_extension.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the extension.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `a URL at which the entity is uniquely displayed in the Web app ## Import Extensions can be imported using the id.e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_extension.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_maintenance_window",
			Category:         "Resources",
			ShortDescription: `Creates and manages a maintenance window in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"maintenance",
				"window",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the maintenance window. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the maintenance window. ## Import Maintenance windows can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_maintenance_window.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the maintenance window. ## Import Maintenance windows can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_maintenance_window.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_schedule",
			Category:         "Resources",
			ShortDescription: `Creates and manages a schedule in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"schedule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the schedule.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Required) The time zone of the schedule (e.g Europe/Berlin).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the schedule`,
				},
				resource.Attribute{
					Name:        "layer",
					Description: `(Required) A schedule layer block. Schedule layers documented below.`,
				},
				resource.Attribute{
					Name:        "overflow",
					Description: `(Optional) Any on-call schedule entries that pass the date range bounds will be truncated at the bounds, unless the parameter ` + "`" + `overflow` + "`" + ` is passed. For instance, if your schedule is a rotation that changes daily at midnight UTC, and your date range is from ` + "`" + `2011-06-01T10:00:00Z` + "`" + ` to ` + "`" + `2011-06-01T14:00:00Z` + "`" + `: If you don't pass the overflow=true parameter, you will get one schedule entry returned with a start of ` + "`" + `2011-06-01T10:00:00Z` + "`" + ` and end of ` + "`" + `2011-06-01T14:00:00Z` + "`" + `. If you do pass the ` + "`" + `overflow` + "`" + ` parameter, you will get one schedule entry returned with a start of ` + "`" + `2011-06-01T00:00:00Z` + "`" + ` and end of ` + "`" + `2011-06-02T00:00:00Z` + "`" + `. Schedule layers (` + "`" + `layer` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the schedule layer.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) The start time of the schedule layer. This value will not be read back from the PagerDuty API because the API will always return a new ` + "`" + `start` + "`" + ` time, which represents the last updated time of the schedule layer.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Optional) The end time of the schedule layer. If not specified, the layer does not end.`,
				},
				resource.Attribute{
					Name:        "rotation_virtual_start",
					Description: `(Required) The effective start time of the schedule layer. This can be before the start time of the schedule.`,
				},
				resource.Attribute{
					Name:        "rotation_turn_length_seconds",
					Description: `(Required) The duration of each on-call shift in ` + "`" + `seconds` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Required) The ordered list of users on this layer. The position of the user on the list determines their order in the layer.`,
				},
				resource.Attribute{
					Name:        "restriction",
					Description: `(Optional) A schedule layer restriction block. Restriction blocks documented below. Restriction blocks (` + "`" + `restriction` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Can be ` + "`" + `daily_restriction` + "`" + ` or ` + "`" + `weekly_restriction` + "`" + ``,
				},
				resource.Attribute{
					Name:        "start_time_of_day",
					Description: `(Required) The start time in ` + "`" + `HH:mm:ss` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "duration_seconds",
					Description: `(Required) The duration of the restriction in ` + "`" + `seconds` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start_day_of_week",
					Description: `(Required for ` + "`" + `weekly_restriction` + "`" + `) Number of the day when restriction starts. From 1 to 7 where 1 is Monday and 7 is Sunday. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the schedule ## Import Schedules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_schedule.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the schedule ## Import Schedules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_schedule.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_service",
			Category:         "Resources",
			ShortDescription: `Creates and manages a service in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the service.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description of the service. If not set, a placeholder of "Managed by Terraform" will be set.`,
				},
				resource.Attribute{
					Name:        "auto_resolve_timeout",
					Description: `(Optional) Time in seconds that an incident is automatically resolved if left open for that long. Disabled if set to the ` + "`" + `"null"` + "`" + ` string.`,
				},
				resource.Attribute{
					Name:        "acknowledgement_timeout",
					Description: `(Optional) Time in seconds that an incident changes to the Triggered State after being Acknowledged. Disabled if set to the ` + "`" + `"null"` + "`" + ` string.`,
				},
				resource.Attribute{
					Name:        "escalation_policy",
					Description: `(Required) The escalation policy used by this service.`,
				},
				resource.Attribute{
					Name:        "alert_creation",
					Description: `(Optional) Must be one of two values. PagerDuty receives events from your monitoring systems and can then create incidents in different ways. Value "create_incidents" is default: events will create an incident that cannot be merged. Value "create_alerts_and_incidents" is the alternative: events will create an alert and then add it to a new incident, these incidents can be merged. You may specify one optional ` + "`" + `incident_urgency_rule` + "`" + ` block configuring what urgencies to use. Your PagerDuty account must have the ` + "`" + `urgencies` + "`" + ` ability to assign an incident urgency rule. The block contains the following arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of incident urgency: ` + "`" + `constant` + "`" + ` or ` + "`" + `use_support_hours` + "`" + ` (when depending on specific support hours; see ` + "`" + `support_hours` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "urgency",
					Description: `The urgency: ` + "`" + `low` + "`" + ` (does not escalate), or ` + "`" + `high` + "`" + ` (follows escalation rules).`,
				},
				resource.Attribute{
					Name:        "during_support_hours",
					Description: `(Optional) Incidents' urgency during support hours.`,
				},
				resource.Attribute{
					Name:        "outside_support_hours",
					Description: `(Optional) Incidents' urgency outside of support hours. When using ` + "`" + `type = "use_support_hours"` + "`" + ` in ` + "`" + `incident_urgency_rule` + "`" + ` you must specify exactly one (otherwise optional) ` + "`" + `support_hours` + "`" + ` block. Your PagerDuty account must have the ` + "`" + `service_support_hours` + "`" + ` ability to assign support hours. The block contains the following arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of support hours. Can be ` + "`" + `fixed_time_per_day` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `The time zone for the support hours.`,
				},
				resource.Attribute{
					Name:        "days_of_week",
					Description: `Array of days of week as integers. ` + "`" + `1` + "`" + ` to ` + "`" + `7` + "`" + `, ` + "`" + `1` + "`" + ` being Monday and ` + "`" + `7` + "`" + ` being Sunday.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The support hours' starting time of day.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The support hours' ending time of day. When using ` + "`" + `type = "use_support_hours"` + "`" + ` in ` + "`" + `incident_urgency_rule` + "`" + ` you must specify at least one (otherwise optional) ` + "`" + `scheduled_actions` + "`" + ` block. The block contains the following arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of scheduled action. Currently, this must be set to ` + "`" + `urgency_change` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "to_urgency",
					Description: `The urgency to change to: ` + "`" + `low` + "`" + ` (does not escalate), or ` + "`" + `high` + "`" + ` (follows escalation rules).`,
				},
				resource.Attribute{
					Name:        "at",
					Description: `A block representing when the scheduled action will occur. The ` + "`" + `at` + "`" + ` block contains the following arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of time specification. Currently, this must be set to ` + "`" + `named_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Designates either the start or the end of the scheduled action. Can be ` + "`" + `support_hours_start` + "`" + ` or ` + "`" + `support_hours_end` + "`" + `. Below is an example for a ` + "`" + `pagerduty_service` + "`" + ` resource with ` + "`" + `incident_urgency_rules` + "`" + ` with ` + "`" + `type = "use_support_hours"` + "`" + `, ` + "`" + `support_hours` + "`" + ` and a default ` + "`" + `scheduled_action` + "`" + ` as well. ` + "`" + `` + "`" + `` + "`" + `hcl resource "pagerduty_service" "foo" { name = "bar" description = "bar bar bar" auto_resolve_timeout = 3600 acknowledgement_timeout = 3600 escalation_policy = "${pagerduty_escalation_policy.foo.id}" incident_urgency_rule { type = "use_support_hours" during_support_hours { type = "constant" urgency = "high" } outside_support_hours { type = "constant" urgency = "low" } } support_hours { type = "fixed_time_per_day" time_zone = "America/Lima" start_time = "09:00:00" end_time = "17:00:00" days_of_week = [1, 2, 3, 4, 5] } scheduled_actions { type = "urgency_change" to_urgency = "high" at { type = "named_time" name = "support_hours_start" } } } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_service_integration",
			Category:         "Resources",
			ShortDescription: `Creates and manages a service integration in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"service",
				"integration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required) The ID of the service the integration should belong to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the service integration.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The service type. Can be: ` + "`" + `aws_cloudwatch_inbound_integration` + "`" + `, ` + "`" + `cloudkick_inbound_integration` + "`" + `, ` + "`" + `event_transformer_api_inbound_integration` + "`" + `, ` + "`" + `events_api_v2_inbound_integration` + "`" + ` (requires service ` + "`" + `alert_creation` + "`" + ` to be ` + "`" + `create_alerts_and_incidents` + "`" + `), ` + "`" + `generic_email_inbound_integration` + "`" + `, ` + "`" + `generic_events_api_inbound_integration` + "`" + `, ` + "`" + `keynote_inbound_integration` + "`" + `, ` + "`" + `nagios_inbound_integration` + "`" + `, ` + "`" + `pingdom_inbound_integration` + "`" + `or ` + "`" + `sql_monitor_inbound_integration` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `(Optional) The ID of the vendor the integration should integrate with (e.g Datadog or Amazon Cloudwatch).`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Optional) This is the unique key used to route events to this integration when received via the PagerDuty Events API.`,
				},
				resource.Attribute{
					Name:        "integration_email",
					Description: `(Optional) This is the unique fully-qualified email address used for routing emails to this integration for processing.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service integration.`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `This is the unique key used to route events to this integration when received via the PagerDuty Events API.`,
				},
				resource.Attribute{
					Name:        "integration_email",
					Description: `This is the unique fully-qualified email address used for routing emails to this integration for processing.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app To configure an event, please use the ` + "`" + `integration_key` + "`" + ` in the following interpolation: ` + "`" + `` + "`" + `` + "`" + `hcl https://events.pagerduty.com/integration/${pagerduty_service_integration.slack.integration_key}/enqueue ` + "`" + `` + "`" + `` + "`" + ` ## Import Services can be imported using their related ` + "`" + `service` + "`" + ` id and service integration ` + "`" + `id` + "`" + ` separated by a dot, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_service_integration.main PLSSSSS.PLIIIII ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service integration.`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `This is the unique key used to route events to this integration when received via the PagerDuty Events API.`,
				},
				resource.Attribute{
					Name:        "integration_email",
					Description: `This is the unique fully-qualified email address used for routing emails to this integration for processing.`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app To configure an event, please use the ` + "`" + `integration_key` + "`" + ` in the following interpolation: ` + "`" + `` + "`" + `` + "`" + `hcl https://events.pagerduty.com/integration/${pagerduty_service_integration.slack.integration_key}/enqueue ` + "`" + `` + "`" + `` + "`" + ` ## Import Services can be imported using their related ` + "`" + `service` + "`" + ` id and service integration ` + "`" + `id` + "`" + ` separated by a dot, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_service_integration.main PLSSSSS.PLIIIII ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_team",
			Category:         "Resources",
			ShortDescription: `Creates and manages a team in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"team",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description of the team. If not set, a placeholder of "Managed by Terraform" will be set. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team. ## Import Teams can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_team.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team. ## Import Teams can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_team.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_team_membership",
			Category:         "Resources",
			ShortDescription: `Creates and manages a team membership in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"team",
				"membership",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) The ID of the user to add to the team.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) The ID of the team in which the user will belong. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user belonging to the team.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The team ID the user belongs to. ## Import Team memberships can be imported using the ` + "`" + `user_id` + "`" + ` and ` + "`" + `team_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_team_membership.main PLBP09X:PLB09Z ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user belonging to the team.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The team ID the user belongs to. ## Import Team memberships can be imported using the ` + "`" + `user_id` + "`" + ` and ` + "`" + `team_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_team_membership.main PLBP09X:PLB09Z ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_user",
			Category:         "Resources",
			ShortDescription: `Creates and manages a user in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The user's email address.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) The schedule color for the user. Valid options are blue, brown, cayenne, chocolate, crimson, cyan, dark, darkolive, deep, firebrick, forest, goldenrod, gray, green, grey, grey20, indigo, lime, magenta, maroon, medium, midnight, olive, olivedrab, orange, orchid, pink, purple, red, royal, saddle, sea, slate, steel, teal, turquoise or violet.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The user role. Account must have the ` + "`" + `read_only_users` + "`" + ` ability to set a user as a ` + "`" + `read_only_user` + "`" + `. Can be ` + "`" + `admin` + "`" + `, ` + "`" + `limited_user` + "`" + `, ` + "`" + `owner` + "`" + `, ` + "`" + `read_only_user` + "`" + `, ` + "`" + `team_responder` + "`" + ` or ` + "`" + `user` + "`" + ``,
				},
				resource.Attribute{
					Name:        "job_title",
					Description: `(Optional) The user's title.`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `(Optional) A list of teams the user should belong to.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description of the user. If not set, a placeholder of "Managed by Terraform" will be set. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the user.`,
				},
				resource.Attribute{
					Name:        "avatar_url",
					Description: `The URL of the user's avatar.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `The timezone of the user`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app`,
				},
				resource.Attribute{
					Name:        "invitation_sent",
					Description: `If true, the user has an outstanding invitation. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_user.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the user.`,
				},
				resource.Attribute{
					Name:        "avatar_url",
					Description: `The URL of the user's avatar.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `The timezone of the user`,
				},
				resource.Attribute{
					Name:        "html_url",
					Description: `URL at which the entity is uniquely displayed in the Web app`,
				},
				resource.Attribute{
					Name:        "invitation_sent",
					Description: `If true, the user has an outstanding invitation. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_user.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_user_contact_method",
			Category:         "Resources",
			ShortDescription: `Creates and manages contact methods for a user in PagerDuty.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"contact",
				"method",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) The ID of the user.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The contact method type. May be (` + "`" + `email_contact_method` + "`" + `, ` + "`" + `phone_contact_method` + "`" + `, ` + "`" + `sms_contact_method` + "`" + `, ` + "`" + `push_notification_contact_method` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "send_short_email",
					Description: `(Optional) Send an abbreviated email message instead of the standard email output.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(Optional) The 1-to-3 digit country calling code. Required when using ` + "`" + `phone_contact_method` + "`" + ` or ` + "`" + `sms_contact_method` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The label (e.g., "Work", "Mobile", etc.).`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The "address" to deliver to: ` + "`" + `email` + "`" + `, ` + "`" + `phone number` + "`" + `, etc., depending on the type. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the contact method.`,
				},
				resource.Attribute{
					Name:        "blacklisted",
					Description: `If true, this phone has been blacklisted by PagerDuty and no messages will be sent to it.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If true, this phone is capable of receiving SMS messages. ## Import Contact methods can be imported using the ` + "`" + `user_id` + "`" + ` and the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_user_contact_method.main PLBP09X:PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the contact method.`,
				},
				resource.Attribute{
					Name:        "blacklisted",
					Description: `If true, this phone has been blacklisted by PagerDuty and no messages will be sent to it.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If true, this phone is capable of receiving SMS messages. ## Import Contact methods can be imported using the ` + "`" + `user_id` + "`" + ` and the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import pagerduty_user_contact_method.main PLBP09X:PLBP09X ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"pagerduty_addon":               0,
		"pagerduty_escalation_policy":   1,
		"pagerduty_extension":           2,
		"pagerduty_maintenance_window":  3,
		"pagerduty_schedule":            4,
		"pagerduty_service":             5,
		"pagerduty_service_integration": 6,
		"pagerduty_team":                7,
		"pagerduty_team_membership":     8,
		"pagerduty_user":                9,
		"pagerduty_user_contact_method": 10,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
