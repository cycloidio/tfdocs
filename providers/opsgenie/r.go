package opsgenie

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_api_integration",
			Category:         "Resources",
			ShortDescription: `Manages an API Integration within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"integration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the integration. Name must be unique for each integration.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of the integration (API, Marid, Prometheus, etc). The full list of options can be found [here](https://docs.opsgenie.com/docs/integration-types-to-use-with-api).`,
				},
				resource.Attribute{
					Name:        "allow_write_access",
					Description: `(Optional) This parameter is for configuring the write access of integration. If write access is restricted, the integration will not be authorized to write within any domain. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) This parameter is for specifying whether the integration will be enabled or not. Defaults to true`,
				},
				resource.Attribute{
					Name:        "ignore_responders_from_payload",
					Description: `(Optional) If enabled, the integration will ignore recipients sent in request payloads. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "suppress_notifications",
					Description: `(Optional) If enabled, notifications that come from alerts will be suppressed. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "owner_team_id",
					Description: `(Optional) Owner team id of the integration.`,
				},
				resource.Attribute{
					Name:        "responders",
					Description: `(Optional) User, schedule, teams or escalation names to calculate which users will receive the notifications of the alert. ` + "`" + `responders` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The responder type.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The id of the responder. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie API Integration.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `(Computed) API key of the created integration ## Import API Integrations can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_team.team1 812be1a1-32c8-4666-a7fb-03ecc385106c` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie API Integration.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `(Computed) API key of the created integration ## Import API Integrations can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_team.team1 812be1a1-32c8-4666-a7fb-03ecc385106c` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_email_integration",
			Category:         "Resources",
			ShortDescription: `Manages an Email Integration within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"email",
				"integration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the integration. Name must be unique for each integration.`,
				},
				resource.Attribute{
					Name:        "email_username",
					Description: `(Required) The username part of the email address. It must be unique for each integration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) A Member block as documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_responders_from_payload",
					Description: `(Optional) If enabled, the integration will ignore recipients sent in request payloads. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "suppress_notifications",
					Description: `(Optional) If enabled, notifications that come from alerts will be suppressed. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "owner_team_id",
					Description: `(Optional) Owner team id of the integration.`,
				},
				resource.Attribute{
					Name:        "responder",
					Description: `(Optional) User, schedule, teams or escalation names to calculate which users will receive the notifications of the alert. ` + "`" + `responder` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The responder type.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The id of the responder. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Email based Integration. ## Import Email Integrations can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_email_integration.test 812be1a1-32c8-4666-a7fb-03ecc385106c` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Email based Integration. ## Import Email Integrations can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_email_integration.test 812be1a1-32c8-4666-a7fb-03ecc385106c` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_escalation",
			Category:         "Resources",
			ShortDescription: `Manages an Escalation within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"escalation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the escalation.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Required) List of the escalation rules.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the escalation.`,
				},
				resource.Attribute{
					Name:        "owner_team_id",
					Description: `(Optional) Owner team id of the escalation.`,
				},
				resource.Attribute{
					Name:        "repeat",
					Description: `(Optional) Repeat preferences of the escalation including repeat interval, count, reverting acknowledge and seen states back and closing an alert automatically as soon as repeats are completed ` + "`" + `rules` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Required) The condition for notifying the recipient of escalation rule that is based on the alert state. Possible values are: if-not-acked and if-not-closed. If not given, if-not-acked is used.`,
				},
				resource.Attribute{
					Name:        "notify_type",
					Description: `(Required) Recipient calculation logic for schedules. Possible values are: ` + "`" + `` + "`" + `` + "`" + ` default: on call users next: next users in rotation previous: previous users on rotation users: users of the team admins: admins of the team all: all members of the team ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "recipient",
					Description: `(Required) Object of schedule, team, or users which will be notified in escalation. The possible values for participants are: user, schedule, team.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Required) Time delay of the escalation rule. This parameter takes an object that consists timeAmount field that takes time amount in minutes. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Escalation. ## Import Escalations can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_escalation.test 812be1a1-32c8-4666-a7fb-03ecc385106c` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Escalation. ## Import Escalations can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_escalation.test 812be1a1-32c8-4666-a7fb-03ecc385106c` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_heartbeat",
			Category:         "Resources",
			ShortDescription: `Manages Heartbeat within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"heartbeat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the heartbeat`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the heartbeat`,
				},
				resource.Attribute{
					Name:        "interval_unit",
					Description: `(Required) Interval specified as minutes, hours or days.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Required) Specifies how often a heartbeat message should be expected.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(True) Enable/disable heartbeat monitoring.`,
				},
				resource.Attribute{
					Name:        "owner_team_id",
					Description: `(Optional) Owner team of the heartbeat.`,
				},
				resource.Attribute{
					Name:        "alert_message",
					Description: `(Optional) Specifies the alert message for heartbeat expiration alert. If this is not provided, default alert message is "HeartbeatName is expired".`,
				},
				resource.Attribute{
					Name:        "alert_priority",
					Description: `(Optional) Specifies the alert priority for heartbeat expiration alert. If this is not provided, default priority is P3.`,
				},
				resource.Attribute{
					Name:        "alert_tags",
					Description: `(Optional) Specifies the alert tags for heartbeat expiration alert. ## Attributes Reference Only the arguments listed above are exposed as attributes. ## Import Heartbeat Integrations can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_heartbeat.test 812be1a1-32c8-4666-a7fb-03ecc385106c` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_maintenance",
			Category:         "Resources",
			ShortDescription: `Manages a Maintenance within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"maintenance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "time",
					Description: `(Required) Time configuration of maintenance. It takes a time object which has type, startDate and endDate fields`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Required) Rules of maintenance, which takes a list of rule objects and defines the maintenance rules over integrations and policies.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the maintenance. ` + "`" + `times` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) This parameter defines when the maintenance will be active. It can take one of for-5-minutes, for-30-minutes, for-1-hour, indefinitely or schedule.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Required) This parameter takes a date format as (yyyy-MM-dd'T'HH:mm:ssZ) (e.g. 2019-06-11T08:00:00+02:00).`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Required) This parameter takes a date format as (yyyy-MM-dd'T'HH:mm:ssZ) (e.g. 2019-06-11T08:00:00+02:00). ` + "`" + `rules` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "entity",
					Description: `(Required) This field represents the entity that maintenance will be applied. Entity field takes two mandatory fields as id and type. - ` + "`" + `id` + "`" + ` - (Required) The id of the entity that maintenance will be applied. - ` + "`" + `type` + "`" + ` - (Required) The type of the entity that maintenance will be applied. It can be either integration or policy.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Required) State of rule that will be defined in maintenance and can take either enabled or disabled for policy type rules. This field has to be disabled for integration type entity rules. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Maintenance Policy. ## Import Maintenance policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_maintenance.test 812be1a1-32c8-4666-a7fb-03ecc385106c` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Maintenance Policy. ## Import Maintenance policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_maintenance.test 812be1a1-32c8-4666-a7fb-03ecc385106c` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_notification_policy",
			Category:         "Resources",
			ShortDescription: `Manages a Notification Policy within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"notification",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the notification policy`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) Id of team that this policy belons to.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) If policy should be enabled. Default: true`,
				},
				resource.Attribute{
					Name:        "policy_description",
					Description: `(Optional) Description of the policy. This can be max 512 characters.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) A notification filter which will be applied. This filter can be empty: filter {} - this means 'match-all'. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "time_restriction",
					Description: `(Optional) Time restrictions specified in this field must be met for this policy to work. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "auto_close_action",
					Description: `(Optional) Auto Restart Action of the policy. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "auto_restart_action",
					Description: `(Optional) Auto Restart Action of the policy. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "de_duplication_action",
					Description: `(Optional) Deduplication Action of the policy. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "delay_action",
					Description: `(Optional) Delay notifications. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "suppress",
					Description: `(Optional) Suppress value of the policy. Values are: true, false. Default: false The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `(Required) Specifies which alert field will be used in condition. Possible values are "message", "alias", "description", "source", "entity", "tags", "actions", "details", "extra-properties", "recipients", "teams", "priority"`,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `(Required) It is the operation that will be executed for the given field and key. Possible operations are "matches", "contains", "starts-with", "ends-with", "equals", "contains-key", "contains-value", "greater-than", "less-than", "is-empty", "equals-ignore-whitespace".`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) If ` + "`" + `field` + "`" + ` is set as extra-properties, key could be used for key-value pair`,
				},
				resource.Attribute{
					Name:        "not",
					Description: `(Optional) Indicates behaviour of the given operation. Default: false`,
				},
				resource.Attribute{
					Name:        "expected_value",
					Description: `(Optional) User defined value that will be compared with alert field according to the operation. Default: empty string`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Order of the condition in conditions list The ` + "`" + `time_restriction` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Defines if restriction should apply daily on given hours or on certain days and hours. Possible values are: "time-of-day", "weekday-and-time-of-day"`,
				},
				resource.Attribute{
					Name:        "restrictions",
					Description: `(Optional) List of days and hours definitions for field type = "weekday-and-time-of-day". This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "restriction",
					Description: `(Optional) A definition of hourly definition applied daily, this has to be used with combination: type = "time-of-day". This is a block, structure is documented below. The ` + "`" + `restrictions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start_day",
					Description: `(Required) Starting day of restriction (eg. "monday")`,
				},
				resource.Attribute{
					Name:        "end_day",
					Description: `(Required) Ending day of restriction (eg. "wednesday)`,
				},
				resource.Attribute{
					Name:        "start_hour",
					Description: `(Required) Starting hour of restriction on defined ` + "`" + `start_day` + "`" + ``,
				},
				resource.Attribute{
					Name:        "end_hour",
					Description: `(Required) Ending hour of restriction on defined ` + "`" + `end_day` + "`" + ``,
				},
				resource.Attribute{
					Name:        "start_minute",
					Description: `(Required) Staring minute of restriction on defined ` + "`" + `start_hour` + "`" + ``,
				},
				resource.Attribute{
					Name:        "end_minute",
					Description: `(Required) Ending minute of restriction on defined ` + "`" + `end_hour` + "`" + ` The ` + "`" + `restriction` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start_hour",
					Description: `(Required) Starting hour of restriction.`,
				},
				resource.Attribute{
					Name:        "end_hour",
					Description: `(Required) Ending hour of restriction.`,
				},
				resource.Attribute{
					Name:        "start_minute",
					Description: `(Required) Staring minute of restriction on defined ` + "`" + `start_hour` + "`" + ``,
				},
				resource.Attribute{
					Name:        "end_minute",
					Description: `(Required) Ending minute of restriction on defined ` + "`" + `end_hour` + "`" + ` The ` + "`" + `auto_close_action` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) Duration of this action. This is a block, structure is documented below. The ` + "`" + `auto_restart_action` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) Duration of this action. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "max_repeat_count",
					Description: `(Required) How many times to repeat. This is a integer attribute. The ` + "`" + `de_duplication_action` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) Duration of this action. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "de_duplication_action_type",
					Description: `(Required) Deduplication type. Possible values are: "value-based", "frequency-based"`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Required) - Count The ` + "`" + `delay_action` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "delay_option",
					Description: `(Required) Defines until what day to delay or for what duration. Possible values are: "for-duration", "next-time", "next-weekday", "next-monday", "next-tuesday", "next-wednesday", "next-thursday", "next-friday", "next-saturday", "next-sunday"`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) Duration of this action. If ` + "`" + `delay_option` + "`" + ` = "for-duration" this has to be set. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "until_hour",
					Description: `(Optional) Until what hour notifications will be delayed. If ` + "`" + `delay_option` + "`" + ` is set to antyhing else then "for-duration" this has to be set.`,
				},
				resource.Attribute{
					Name:        "until_minute",
					Description: `(Optional) Until what minute on ` + "`" + `until_hour` + "`" + ` notifications will be delayed. If ` + "`" + `delay_option` + "`" + ` is set to antyhing else then "for-duration" this has to be set. The ` + "`" + `duration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "time_unit",
					Description: `(Optional) Valid time units are: "minutes", "hours", "days". Default: minutes`,
				},
				resource.Attribute{
					Name:        "time_amount",
					Description: `(Required) A amount of time in ` + "`" + `time_units` + "`" + `. This is a integer attribute. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Notification Policy. ## Import Notification policies can be imported using the ` + "`" + `team id` + "`" + ` and ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_notification_policy.test teamId/Id` + "`" + ` For this example: - Team Id = ` + "`" + `c827c472-31f2-497b-9ec6-8ec855d7d94c` + "`" + ` - Notification Policy Id = ` + "`" + `2d1a78d0-c13e-47d3-af0a-8b6d0cc2b7b1` + "`" + ` ` + "`" + `$ terraform import opsgenie_notification_policy.test c827c472-31f2-497b-9ec6-8ec855d7d94c/2d1a78d0-c13e-47d3-af0a-8b6d0cc2b7b1` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Notification Policy. ## Import Notification policies can be imported using the ` + "`" + `team id` + "`" + ` and ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_notification_policy.test teamId/Id` + "`" + ` For this example: - Team Id = ` + "`" + `c827c472-31f2-497b-9ec6-8ec855d7d94c` + "`" + ` - Notification Policy Id = ` + "`" + `2d1a78d0-c13e-47d3-af0a-8b6d0cc2b7b1` + "`" + ` ` + "`" + `$ terraform import opsgenie_notification_policy.test c827c472-31f2-497b-9ec6-8ec855d7d94c/2d1a78d0-c13e-47d3-af0a-8b6d0cc2b7b1` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_schedule",
			Category:         "Resources",
			ShortDescription: `Manages a Schedule within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"schedule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the schedule.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Required) A Member block as documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Timezone of schedule. Please look at [Supported Timezone Ids](https://docs.opsgenie.com/docs/supported-timezone-ids) for available timezones - Defaults to "America/New_York".`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) The description of schedule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable/disable state of schedule`,
				},
				resource.Attribute{
					Name:        "owner_team_id",
					Description: `(Optional) Owner team id of the schedule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Schedule. ## Import Schedule can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_schedule.test 812be1a1-32c8-4666-a7fb-03ecc385106c` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Schedule. ## Import Schedule can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_schedule.test 812be1a1-32c8-4666-a7fb-03ecc385106c` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_schedule_rotation",
			Category:         "Resources",
			ShortDescription: `Manages a Schedule Rotation within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"schedule",
				"rotation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "schedule_id",
					Description: `(Required) Identifier of the schedule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of rotation.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Required) This parameter takes a date format as (yyyy-MM-dd'T'HH:mm:ssZ) (e.g. 2019-06-11T08:00:00+02:00). Minutes may take 0 or 30 as value. Otherwise they will be converted to nearest 0 or 30 automatically`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) This parameter takes a date format as (yyyy-MM-dd'T'HH:mm:ssZ) (e.g. 2019-06-11T08:00:00+02:00). Minutes may take 0 or 30 as value. Otherwise they will be converted to nearest 0 or 30 automatically`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of rotation. May be one of daily, weekly and hourly.`,
				},
				resource.Attribute{
					Name:        "length",
					Description: `(Required) Length of the rotation with default value 1.`,
				},
				resource.Attribute{
					Name:        "participant",
					Description: `(Required) List of escalations, teams, users or the reserved word none which will be used in schedule. Each of them can be used multiple times and will be rotated in the order they given. "user,escalation,team,none"`,
				},
				resource.Attribute{
					Name:        "time_restriction",
					Description: `(Required) ` + "`" + `participant` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The responder type.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The id of the responder. ` + "`" + `time_restriction` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) This parameter should be set time-of-day`,
				},
				resource.Attribute{
					Name:        "restriction",
					Description: `(Required) It is a restriction object which is described below. In this case startDay/endDay fields are not supported. ` + "`" + `restriction` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "start_hour",
					Description: `(Required) Value of the hour that frame will start`,
				},
				resource.Attribute{
					Name:        "start_min",
					Description: `(Required) Value of the minute that frame will start. Minutes may take 0 or 30 as value. Otherwise they will be converted to nearest 0 or 30 automatically.`,
				},
				resource.Attribute{
					Name:        "end_hour",
					Description: `(Required) Value of the hour that frame will end.`,
				},
				resource.Attribute{
					Name:        "end_min",
					Description: `(Required) Value of the minute that frame will end. Minutes may take 0 or 30 as value. Otherwise they will be converted to nearest 0 or 30 automatically. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Schedule Rotation ## Import Schedule Rotations can be imported using the ` + "`" + `id` + "`" + ` and ` + "`" + `schedule_id` + "`" + `, e.g.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Schedule Rotation ## Import Schedule Rotations can be imported using the ` + "`" + `id` + "`" + ` and ` + "`" + `schedule_id` + "`" + `, e.g.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_team",
			Category:         "Resources",
			ShortDescription: `Manages a Team within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"team",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name associated with this team. Opsgenie defines that this must not be longer than 100 characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for this team.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Optional) A Member block as documented below. ` + "`" + `member` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The UUID for the member to add to this Team.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The role for the user within the Team - can be either 'admin' or 'user', defaults to 'user' if not set. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Team. ## Import Teams can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_team.team1 812be1a1-32c8-4666-a7fb-03ecc385106c` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Team. ## Import Teams can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_team.team1 812be1a1-32c8-4666-a7fb-03ecc385106c` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_team_routing_rule",
			Category:         "Resources",
			ShortDescription: `Manages a Team Routing Rule within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"team",
				"routing",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the team routing rule`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) The order of the team routing rule within the rules. order value is actually the index of the team routing rule whose minimum value is 0 and whose maximum value is n-1 (number of team routing rules is n)`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) Timezone of team routing rule. If timezone field is not given, account timezone is used as default.You can refer to Supported Locale IDs for available timezones`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `(Optional) You can refer Criteria for detailed information about criteria and its fields`,
				},
				resource.Attribute{
					Name:        "timeRestriction",
					Description: `(Optional) You can refer Time Restriction for detailed information about time restriction and its fields.`,
				},
				resource.Attribute{
					Name:        "notify",
					Description: `(Optional) Target entity of schedule, escalation, or the reserved word none which will be notified in routing rule. The possible values for notify type : schedule, escalation, none ` + "`" + `notify` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ` + "`" + `criteria` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of the operation will be applied on conditions. Should be one of match-all, match-any-condition or match-all-conditions.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional) List of conditions will be checked before applying team routing rule. ` + "`" + `conditions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `(Required) Specifies which alert field will be used in condition. Possible values are message, alias, description, source, entity, tags, actions, extra-properties, recipients, teams or priority.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) If field is set as extra-properties, key could be used for key-value pair.`,
				},
				resource.Attribute{
					Name:        "not",
					Description: `(Optional) Indicates behaviour of the given operation. Default value is false.`,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `(true) It is the operation that will be executed for the given field and key. Possible operations are matches, contains, starts-with, ends-with, equals, contains-key, contains-value, greater-than, less-than, is-empty and equals-ignore-whitespace.`,
				},
				resource.Attribute{
					Name:        "expectedValue",
					Description: `(Optional) User defined value that will be compared with alert field according to the operation. Default value is empty string.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Order of the condition in conditions list. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Team Routing Rule. ## Import Teams can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_team_routing_rule.ruletest teamId/routingRuleId` + "`" + ` For this example: - Team Id = ` + "`" + `c827c472-31f2-497b-9ec6-8ec855d7d94c` + "`" + ` - Routing Rule Id = ` + "`" + `2d1a78d0-c13e-47d3-af0a-8b6d0cc2b7b1` + "`" + ` ` + "`" + `$ terraform import opsgenie_team_routing_rule.ruletest c827c472-31f2-497b-9ec6-8ec855d7d94c/2d1a78d0-c13e-47d3-af0a-8b6d0cc2b7b1` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Team Routing Rule. ## Import Teams can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_team_routing_rule.ruletest teamId/routingRuleId` + "`" + ` For this example: - Team Id = ` + "`" + `c827c472-31f2-497b-9ec6-8ec855d7d94c` + "`" + ` - Routing Rule Id = ` + "`" + `2d1a78d0-c13e-47d3-af0a-8b6d0cc2b7b1` + "`" + ` ` + "`" + `$ terraform import opsgenie_team_routing_rule.ruletest c827c472-31f2-497b-9ec6-8ec855d7d94c/2d1a78d0-c13e-47d3-af0a-8b6d0cc2b7b1` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_user",
			Category:         "Resources",
			ShortDescription: `Manages a User within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The email address associated with this user. Opsgenie defines that this must not be longer than 100 characters.`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `(Required) The Full Name of the User.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The Role assigned to the User. Either a built-in such as 'Admin' or 'User' - or the name of a custom role.`,
				},
				resource.Attribute{
					Name:        "locale",
					Description: `(Optional) Location information for the user. Please look at [Supported Locale Ids](https://docs.opsgenie.com/docs/supported-locales) for available locales.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) Timezone information of the user. Please look at [Supported Timezone Ids](https://docs.opsgenie.com/docs/supported-timezone-ids) for available timezones. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie User. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_user.user da4faf16-5546-41e4-8330-4d0002b74048s` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie User. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_user.user da4faf16-5546-41e4-8330-4d0002b74048s` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_user_contact",
			Category:         "Resources",
			ShortDescription: `Manages a User Contact.`,
			Description:      ``,
			Keywords: []string{
				"user",
				"contact",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username for contact.(reference)`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `(Required) to field is the address of given method.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) This parameter is the contact method of user and should be one of email, sms or voice. Please note that adding mobile is not supported from API.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable contact of the user in OpsGenie. Default value is true. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Contact. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_user_contact.testcontact username/contactId` + "`" + ` For this example: - Username = ` + "`" + `genie@awesometeam.com` + "`" + ` - Contact Id = ` + "`" + `2d1a78d0-c13e-47d3-af0a-8b6d0cc2b7b1` + "`" + ` ` + "`" + `$ terraform import opsgenie_user_contact.testcontact genie@awesometeam.com/2d1a78d0-c13e-47d3-af0a-8b6d0cc2b7b1` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Contact. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_user_contact.testcontact username/contactId` + "`" + ` For this example: - Username = ` + "`" + `genie@awesometeam.com` + "`" + ` - Contact Id = ` + "`" + `2d1a78d0-c13e-47d3-af0a-8b6d0cc2b7b1` + "`" + ` ` + "`" + `$ terraform import opsgenie_user_contact.testcontact genie@awesometeam.com/2d1a78d0-c13e-47d3-af0a-8b6d0cc2b7b1` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"opsgenie_api_integration":     0,
		"opsgenie_email_integration":   1,
		"opsgenie_escalation":          2,
		"opsgenie_heartbeat":           3,
		"opsgenie_maintenance":         4,
		"opsgenie_notification_policy": 5,
		"opsgenie_schedule":            6,
		"opsgenie_schedule_rotation":   7,
		"opsgenie_team":                8,
		"opsgenie_team_routing_rule":   9,
		"opsgenie_user":                10,
		"opsgenie_user_contact":        11,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
