package opsgenie

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_alert_policy",
			Category:         "Resources",
			ShortDescription: `Manages a Alert Policy within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the alert policy`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `(Optional) Id of team that this policy belongs to.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) If policy should be enabled. Default: ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "policy_description",
					Description: `(Optional) Description of the policy. This can be max 512 characters.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) A alert filter which will be applied. This filter can be empty: ` + "`" + `filter {}` + "`" + ` - this means ` + "`" + `match-all` + "`" + `. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "time_restriction",
					Description: `(Optional) Time restrictions specified in this field must be met for this policy to work. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `(Required) Message of the alerts`,
				},
				resource.Attribute{
					Name:        "continue_policy",
					Description: `(Optional) It will trigger other modify policies if set to ` + "`" + `true` + "`" + `. Default: ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `(Optional) Alias of the alert. You can use ` + "`" + `{{alias}}` + "`" + ` to refer to the original alias. Default: ` + "`" + `{{alias}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "alert_description",
					Description: `(Optional) Description of the alert. You can use ` + "`" + `{{description}}` + "`" + ` to refer to the original alert description. Default: ` + "`" + `{{description}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "entity",
					Description: `(Optional) Entity field of the alert. You can use ` + "`" + `{{entity}}` + "`" + ` to refer to the original entity. Default: ` + "`" + `{{entity}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Source field of the alert. You can use ` + "`" + `{{source}}` + "`" + ` to refer to the original source. Default: ` + "`" + `{{source}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ignore_original_actions",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, policy will ignore the original actions of the alert. Default: ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ignore_original_details",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, policy will ignore the original details of the alert. Default: ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ignore_original_responders",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, policy will ignore the original responders of the alert. Default: ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "responders",
					Description: `(Optional) Responders to add to the alerts original responders value as a list of teams, users or the reserved word none or all. If ` + "`" + `ignore_original_responders` + "`" + ` field is set to ` + "`" + `true` + "`" + `, this will replace the original responders. The possible values for responders are: ` + "`" + `user` + "`" + `, ` + "`" + `team` + "`" + `. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "ignore_original_tags",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, policy will ignore the original tags of the alert. Default: ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `(Optional) Actions to add to the alerts original actions value as a list of strings. If ` + "`" + `ignore_original_actions` + "`" + ` field is set to ` + "`" + `true` + "`" + `, this will replace the original actions.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags to add to the alerts original tags value as a list of strings. If ` + "`" + `ignore_original_responders` + "`" + ` field is set to ` + "`" + `true` + "`" + `, this will replace the original responders.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority of the alert. Should be one of ` + "`" + `P1` + "`" + `, ` + "`" + `P2` + "`" + `, ` + "`" + `P3` + "`" + `, ` + "`" + `P4` + "`" + `, or ` + "`" + `P5` + "`" + ` The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `(Required) Specifies which alert field will be used in condition. Possible values are ` + "`" + `message` + "`" + `, ` + "`" + `alias` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `source` + "`" + `, ` + "`" + `entity` + "`" + `, ` + "`" + `tags` + "`" + `, ` + "`" + `actions` + "`" + `, ` + "`" + `details` + "`" + `, ` + "`" + `extra-properties` + "`" + `, ` + "`" + `recipients` + "`" + `, ` + "`" + `teams` + "`" + `, ` + "`" + `priority` + "`" + ``,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `(Required) It is the operation that will be executed for the given field and key. Possible operations are ` + "`" + `matches` + "`" + `, ` + "`" + `contains` + "`" + `, ` + "`" + `starts-with` + "`" + `, ` + "`" + `ends-with` + "`" + `, ` + "`" + `equals` + "`" + `, ` + "`" + `contains-key` + "`" + `, ` + "`" + `contains-value` + "`" + `, ` + "`" + `greater-than` + "`" + `, ` + "`" + `less-than` + "`" + `, ` + "`" + `is-empty` + "`" + `, ` + "`" + `equals-ignore-whitespace` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) If ` + "`" + `field` + "`" + ` is set as extra-properties, key could be used for key-value pair`,
				},
				resource.Attribute{
					Name:        "not",
					Description: `(Optional) Indicates behaviour of the given operation. Default: ` + "`" + `false` + "`" + ``,
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
					Description: `(Required) Defines if restriction should apply daily on given hours or on certain days and hours. Possible values are: ` + "`" + `time-of-day` + "`" + `, ` + "`" + `weekday-and-time-of-day` + "`" + ``,
				},
				resource.Attribute{
					Name:        "restrictions",
					Description: `(Optional) List of days and hours definitions for field type = ` + "`" + `weekday-and-time-of-day` + "`" + `. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "restriction",
					Description: `(Optional) A definition of hourly definition applied daily, this has to be used with combination: type = ` + "`" + `time-of-day` + "`" + `. This is a block, structure is documented below. The ` + "`" + `restrictions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start_day",
					Description: `(Required) Starting day of restriction (eg. ` + "`" + `monday` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "end_day",
					Description: `(Required) Ending day of restriction (eg. ` + "`" + `wednesday` + "`" + `)`,
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
					Name:        "start_min",
					Description: `(Required) Staring minute of restriction on defined ` + "`" + `start_hour` + "`" + ``,
				},
				resource.Attribute{
					Name:        "end_min",
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
					Name:        "start_min",
					Description: `(Required) Staring minute of restriction on defined ` + "`" + `start_hour` + "`" + ``,
				},
				resource.Attribute{
					Name:        "end_min",
					Description: `(Required) Ending minute of restriction on defined ` + "`" + `end_hour` + "`" + ` The ` + "`" + `responders` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of responder. Acceptable values are: ` + "`" + `user` + "`" + ` or ` + "`" + `team` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the responder`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ID of the responder`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Username of the responder ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Alert Policy. ## Import Alert policies can be imported using the ` + "`" + `team_id/policy_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_notification_policy.test team_id/policy_id` + "`" + ` You can import global polices using only policy identifier ` + "`" + `$ terraform import opsgenie_alert_policy.test policy_id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Alert Policy. ## Import Alert policies can be imported using the ` + "`" + `team_id/policy_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_notification_policy.test team_id/policy_id` + "`" + ` You can import global polices using only policy identifier ` + "`" + `$ terraform import opsgenie_alert_policy.test policy_id` + "`" + ``,
				},
			},
		},
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
					Description: `(Optional) This parameter is for configuring the write access of integration. If write access is restricted, the integration will not be authorized to write within any domain. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) This parameter is for specifying whether the integration will be enabled or not. Default: ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ignore_responders_from_payload",
					Description: `(Optional) If enabled, the integration will ignore recipients sent in request payloads. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suppress_notifications",
					Description: `(Optional) If enabled, notifications that come from alerts will be suppressed. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "owner_team_id",
					Description: `(Optional) Owner team id of the integration.`,
				},
				resource.Attribute{
					Name:        "responders",
					Description: `(Optional) User, schedule, teams or escalation names to calculate which users will receive the notifications of the alert.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `(Optional) It is required if type is ` + "`" + `Webhook` + "`" + `. This is the url Opsgenie will be sending request to. ` + "`" + `responders` + "`" + ` supports the following:`,
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
					Description: `(Computed) API key of the created integration ## Import API Integrations can be imported using the ` + "`" + `integration_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_api_integration.this integration_id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie API Integration.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `(Computed) API key of the created integration ## Import API Integrations can be imported using the ` + "`" + `integration_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_api_integration.this integration_id` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_custom_role",
			Category:         "Resources",
			ShortDescription: `Manages custom user roles within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"custom",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) Name of the custom role.`,
				},
				resource.Attribute{
					Name:        "extended_role",
					Description: `(Required) The role from which this role has been derived. Allowed Values: "user", "observer", "stakeholder".`,
				},
				resource.Attribute{
					Name:        "granted_rights",
					Description: `(Optional) The rights granted to this role. For allowed values please refer [User Right Prerequisites](https://docs.opsgenie.com/docs/custom-user-role-api#section-user-right-prerequisites)`,
				},
				resource.Attribute{
					Name:        "disallowed_rights",
					Description: `(Optional) The rights this role cannot have. For allowed values please refer [User Right Prerequisites](https://docs.opsgenie.com/docs/custom-user-role-api#section-user-right-prerequisites) ## Attributes Reference Only the arguments listed above are exposed as attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Optional) If enabled, the integration will ignore recipients sent in request payloads. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suppress_notifications",
					Description: `(Optional) If enabled, notifications that come from alerts will be suppressed. Default: ` + "`" + `false` + "`" + `.`,
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
					Description: `The ID of the Opsgenie Email based Integration. ## Import Email Integrations can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_email_integration.test id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Email based Integration. ## Import Email Integrations can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_email_integration.test id` + "`" + ``,
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
					Description: `(Required) The condition for notifying the recipient of escalation rule that is based on the alert state. Possible values are: ` + "`" + `if-not-acked` + "`" + ` and ` + "`" + `if-not-closed` + "`" + `. Default: ` + "`" + `if-not-acked` + "`" + ``,
				},
				resource.Attribute{
					Name:        "notify_type",
					Description: `(Required) Recipient calculation logic for schedules. Possible values are: - ` + "`" + `default` + "`" + `: on call users - ` + "`" + `next` + "`" + `: next users in rotation - ` + "`" + `previous` + "`" + `: previous users on rotation - ` + "`" + `users` + "`" + `: users of the team - ` + "`" + `admins` + "`" + `: admins of the team - ` + "`" + `all` + "`" + `: all members of the team`,
				},
				resource.Attribute{
					Name:        "recipient",
					Description: `(Required) Object of schedule, team, or users which will be notified in escalation. The possible values for participants are: ` + "`" + `user` + "`" + `, ` + "`" + `schedule` + "`" + `, ` + "`" + `team` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Required) Time delay of the escalation rule, in minutes. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Escalation. ## Import Escalations can be imported using the ` + "`" + `escalation_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_escalation.test escalation_id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Escalation. ## Import Escalations can be imported using the ` + "`" + `escalation_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_escalation.test escalation_id` + "`" + ``,
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
					Description: `(Optional) Specifies the alert tags for heartbeat expiration alert. ## Attributes Reference Only the arguments listed above are exposed as attributes. ## Import Heartbeat Integrations can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_heartbeat.test name` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_incident_template",
			Category:         "Resources",
			ShortDescription: `Manages an Incident Template within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"incident",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Incident Template. ## Import Service can be imported using the ` + "`" + `template_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_incident_template.test template_id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Incident Template. ## Import Service can be imported using the ` + "`" + `template_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_incident_template.test template_id` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_integration_action",
			Category:         "Resources",
			ShortDescription: `Manages advanced actions for integrations within Opsgenie`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"action",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "integration_id",
					Description: `(Required) ID of the parent integration resource to bind to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the integration action.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `(Optional) An identifier that is used for alert deduplication. Default: ` + "`" + `{{alias}}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Integer value that defines in which order the action will be performed. Default: ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) Owner of the execution for integration action.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Additional alert action note.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Used to specify rules for matching alerts and the filter type. Please note that depending on the integration type the field names in the filter conditions are:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Detailed description of the alert, anything that may not have fit in the ` + "`" + `message` + "`" + ` field.`,
				},
				resource.Attribute{
					Name:        "entity",
					Description: `(Optional) The entity the alert is related to.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Alert priority.`,
				},
				resource.Attribute{
					Name:        "custom_priority",
					Description: `(Optional) Custom alert priority. e.g. ` + "`" + `` + "`" + `{{message.substring(0,2)}}` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "extra_properties",
					Description: `(Optional) Set of user defined properties specified as a map.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `(Optional) Alert text limited to 130 characters.`,
				},
				resource.Attribute{
					Name:        "responders",
					Description: `(Optional) User, schedule, teams or escalation names to calculate which users will receive notifications of the alert.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) User defined field to specify source of action.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Comma separated list of labels to be attached to the alert.`,
				},
				resource.Attribute{
					Name:        "ignore_responders_from_payload",
					Description: `(Optional) If enabled, the integration will ignore responders sent in request payloads.`,
				},
				resource.Attribute{
					Name:        "ignore_teams_from_payload",
					Description: `(Optional) If enabled, the integration will ignore teams sent in request payloads. ` + "`" + `responders` + "`" + ` is supported only in create action and supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The id of the responder.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The responder type - can be ` + "`" + `escalation` + "`" + `, ` + "`" + `team` + "`" + ` or ` + "`" + `user` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie API Integration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie API Integration.`,
				},
			},
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
					Description: `The ID of the Opsgenie Maintenance Policy. ## Import Maintenance policies can be imported using the ` + "`" + `policy_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_maintenance.test policy_id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Maintenance Policy. ## Import Maintenance policies can be imported using the ` + "`" + `policy_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_maintenance.test policy_id` + "`" + ``,
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
					Description: `(Optional) If policy should be enabled. Default: ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "policy_description",
					Description: `(Optional) Description of the policy. This can be max 512 characters.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) A notification filter which will be applied. This filter can be empty: ` + "`" + `filter {}` + "`" + ` - this means ` + "`" + `match-all` + "`" + `. This is a block, structure is documented below.`,
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
					Description: `(Optional) Suppress value of the policy. Values are: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `. Default: ` + "`" + `false` + "`" + ` The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `(Required) Specifies which alert field will be used in condition. Possible values are ` + "`" + `message` + "`" + `, ` + "`" + `alias` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `source` + "`" + `, ` + "`" + `entity` + "`" + `, ` + "`" + `tags` + "`" + `, ` + "`" + `actions` + "`" + `, ` + "`" + `details` + "`" + `, ` + "`" + `extra-properties` + "`" + `, ` + "`" + `recipients` + "`" + `, ` + "`" + `teams` + "`" + `, ` + "`" + `priority` + "`" + ``,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `(Required) It is the operation that will be executed for the given field and key. Possible operations are ` + "`" + `matches` + "`" + `, ` + "`" + `contains` + "`" + `, ` + "`" + `starts-with` + "`" + `, ` + "`" + `ends-with` + "`" + `, ` + "`" + `equals` + "`" + `, ` + "`" + `contains-key` + "`" + `, ` + "`" + `contains-value` + "`" + `, ` + "`" + `greater-than` + "`" + `, ` + "`" + `less-than` + "`" + `, ` + "`" + `is-empty` + "`" + `, ` + "`" + `equals-ignore-whitespace` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) If ` + "`" + `field` + "`" + ` is set as extra-properties, key could be used for key-value pair`,
				},
				resource.Attribute{
					Name:        "not",
					Description: `(Optional) Indicates behaviour of the given operation. Default: ` + "`" + `false` + "`" + ``,
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
					Description: `(Required) Defines if restriction should apply daily on given hours or on certain days and hours. Possible values are: ` + "`" + `time-of-day` + "`" + `, ` + "`" + `weekday-and-time-of-day` + "`" + ``,
				},
				resource.Attribute{
					Name:        "restrictions",
					Description: `(Optional) List of days and hours definitions for field type = ` + "`" + `weekday-and-time-of-day` + "`" + `. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "restriction",
					Description: `(Optional) A definition of hourly definition applied daily, this has to be used with combination: type = ` + "`" + `time-of-day` + "`" + `. This is a block, structure is documented below. The ` + "`" + `restrictions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start_day",
					Description: `(Required) Starting day of restriction (eg. ` + "`" + `monday` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "end_day",
					Description: `(Required) Ending day of restriction (eg. ` + "`" + `wednesday` + "`" + `)`,
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
					Name:        "start_min",
					Description: `(Required) Staring minute of restriction on defined ` + "`" + `start_hour` + "`" + ``,
				},
				resource.Attribute{
					Name:        "end_min",
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
					Name:        "start_min",
					Description: `(Required) Staring minute of restriction on defined ` + "`" + `start_hour` + "`" + ``,
				},
				resource.Attribute{
					Name:        "end_min",
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
					Name:        "de_duplication_action_type",
					Description: `(Required) Deduplication type. Possible values are: "value-based", "frequency-based"`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Required) - Count`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) Duration of this action (only required for "frequency-based" de-duplication action). This is a block, structure is documented below. The ` + "`" + `delay_action` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "delay_option",
					Description: `(Required) Defines until what day to delay or for what duration. Possible values are: ` + "`" + `for-duration` + "`" + `, ` + "`" + `next-time` + "`" + `, ` + "`" + `next-weekday` + "`" + `, ` + "`" + `next-monday` + "`" + `, ` + "`" + `next-tuesday` + "`" + `, ` + "`" + `next-wednesday` + "`" + `, ` + "`" + `next-thursday` + "`" + `, ` + "`" + `next-friday` + "`" + `, ` + "`" + `next-saturday` + "`" + `, ` + "`" + `next-sunday` + "`" + ``,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) Duration of this action. If ` + "`" + `delay_option` + "`" + ` = ` + "`" + `for-duration` + "`" + ` this has to be set. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "until_hour",
					Description: `(Optional) Until what hour notifications will be delayed. If ` + "`" + `delay_option` + "`" + ` is set to antyhing else then ` + "`" + `for-duration` + "`" + ` this has to be set.`,
				},
				resource.Attribute{
					Name:        "until_minute",
					Description: `(Optional) Until what minute on ` + "`" + `until_hour` + "`" + ` notifications will be delayed. If ` + "`" + `delay_option` + "`" + ` is set to antyhing else then ` + "`" + `for-duration` + "`" + ` this has to be set. The ` + "`" + `duration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "time_unit",
					Description: `(Optional) Valid time units are: ` + "`" + `minutes` + "`" + `, ` + "`" + `hours` + "`" + `, ` + "`" + `days` + "`" + `. Default: ` + "`" + `minutes` + "`" + ``,
				},
				resource.Attribute{
					Name:        "time_amount",
					Description: `(Required) A amount of time in ` + "`" + `time_units` + "`" + `. This is a integer attribute. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Notification Policy. ## Import Notification policies can be imported using the ` + "`" + `team_id` + "`" + ` and ` + "`" + `notification_policy_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_notification_policy.test team_id/notification_policy_id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Notification Policy. ## Import Notification policies can be imported using the ` + "`" + `team_id` + "`" + ` and ` + "`" + `notification_policy_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_notification_policy.test team_id/notification_policy_id` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_notification_rule",
			Category:         "Resources",
			ShortDescription: `Manages a Notification Rule within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"notification",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the notification policy`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username of user to which this notification rule belongs to.`,
				},
				resource.Attribute{
					Name:        "action_type",
					Description: `(Required) Type of the action that notification rule will have. Allowed values: ` + "`" + `create-alert` + "`" + `, ` + "`" + `acknowledged-alert` + "`" + `, ` + "`" + `closed-alert` + "`" + `, ` + "`" + `assigned-alert` + "`" + `, ` + "`" + `add-note` + "`" + `, ` + "`" + `schedule-start` + "`" + `, ` + "`" + `schedule-end` + "`" + `, ` + "`" + `incoming-call-routing` + "`" + ``,
				},
				resource.Attribute{
					Name:        "notification_time",
					Description: `(Optional) List of Time Periods that notification for schedule start/end will be sent. Allowed values: ` + "`" + `just-before` + "`" + `, ` + "`" + `15-minutes-ago` + "`" + `, ` + "`" + `1-hour-ago` + "`" + `, ` + "`" + `1-day-ago` + "`" + `. If ` + "`" + `action_type` + "`" + ` is ` + "`" + `schedule-start` + "`" + ` or ` + "`" + `schedule-end` + "`" + ` then it is required.`,
				},
				resource.Attribute{
					Name:        "steps",
					Description: `(Optional) Notification rule steps to take (eg. SMS or email message). This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) If policy should be enabled. Default: ` + "`" + `true` + "`" + ` The ` + "`" + `steps` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Defined if this step is enabled. Default: ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "send_after",
					Description: `(Optional) Time period, in minutes, notification will be sent after.`,
				},
				resource.Attribute{
					Name:        "contact",
					Description: `(Required) Defines the contact that notification will be sent to. This is a block, structure is documented below. The ` + "`" + `contact` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) Contact method. Possible values: ` + "`" + `email` + "`" + `, ` + "`" + `sms` + "`" + `, ` + "`" + `voice` + "`" + `, ` + "`" + `mobile` + "`" + ``,
				},
				resource.Attribute{
					Name:        "to",
					Description: `(Required) Address of a given method (eg. email address for ` + "`" + `email` + "`" + `, phone number for ` + "`" + `sms` + "`" + `/` + "`" + `voice` + "`" + ` or mobile application name for ` + "`" + `mobile` + "`" + `) The ` + "`" + `criteria` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Kind of matching filter. Possible values: ` + "`" + `match-all` + "`" + `, ` + "`" + `match-any-condition` + "`" + `, ` + "`" + `match-all-conditions` + "`" + ``,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional) Defines the fields and values when the condition applies The ` + "`" + `conditions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `(Required) Possible values: ` + "`" + `message` + "`" + `, ` + "`" + `alias` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `source` + "`" + `, ` + "`" + `entity` + "`" + `, ` + "`" + `tags` + "`" + `, ` + "`" + `actions` + "`" + `, ` + "`" + `details` + "`" + `, ` + "`" + `extra-properties` + "`" + `, ` + "`" + `recipients` + "`" + `, ` + "`" + `teams` + "`" + `, ` + "`" + `priority` + "`" + ``,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `(Required) Possible values: ` + "`" + `matches` + "`" + `, ` + "`" + `contains` + "`" + `, ` + "`" + `starts-with` + "`" + `, ` + "`" + `ends-with` + "`" + `, ` + "`" + `equals` + "`" + `, ` + "`" + `contains-key` + "`" + `, ` + "`" + `contains-value` + "`" + `, ` + "`" + `greater-than` + "`" + `, ` + "`" + `less-than` + "`" + `, ` + "`" + `is-empty` + "`" + `, ` + "`" + `equals-ignore-whitespace` + "`" + ``,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) If 'field' is set as 'extra-properties', key could be used for key-value pair`,
				},
				resource.Attribute{
					Name:        "not",
					Description: `(Optional) Indicates behaviour of the given operation. Default: ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "expected_value",
					Description: `(Optional) User defined value that will be compared with alert field according to the operation. Default: empty string`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Order of the condition in conditions list ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Notification Rule. ## Import Notification policies can be imported using the ` + "`" + `user_id/notification_rule_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_notification_rule.test user_id/notification_rule_id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Notification Rule. ## Import Notification policies can be imported using the ` + "`" + `user_id/notification_rule_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_notification_rule.test user_id/notification_rule_id` + "`" + ``,
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
					Description: `(Optional) The description of schedule.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) Timezone of schedule. Please look at [Supported Timezone Ids](https://docs.opsgenie.com/docs/supported-timezone-ids) for available timezones - Default: ` + "`" + `America/New_York` + "`" + `.`,
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
					Description: `The ID of the Opsgenie Schedule. ## Import Schedule can be imported using the ` + "`" + `schedule_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_schedule.test schedule_id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Schedule. ## Import Schedule can be imported using the ` + "`" + `schedule_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_schedule.test schedule_id` + "`" + ``,
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
					Description: `(Optional) Length of the rotation with default value 1.`,
				},
				resource.Attribute{
					Name:        "participant",
					Description: `(Required) List of escalations, teams, users or the reserved word none which will be used in schedule. Each of them can be used multiple times and will be rotated in the order they given. "user,escalation,team,none"`,
				},
				resource.Attribute{
					Name:        "time_restriction",
					Description: `(Optional) ` + "`" + `participant` + "`" + ` supports the following:`,
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
					Description: `(Required) This parameter should be set to ` + "`" + `time-of-day` + "`" + ` or ` + "`" + `weekday-and-time-of-day` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "restriction",
					Description: `(Optional) It is a restriction object which is described below. In this case startDay/endDay fields are not supported. This can be used only if time restriction type is ` + "`" + `time-of-day` + "`" + `. ` + "`" + `restriction` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "start_hour",
					Description: `(Required) Value of the hour that frame will start.`,
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
					Description: `(Required) Value of the minute that frame will end. Minutes may take 0 or 30 as value. Otherwise they will be converted to nearest 0 or 30 automatically.`,
				},
				resource.Attribute{
					Name:        "restrictions",
					Description: `(Optional) It is a restriction object which is described below. This can be used only if time restriction type is ` + "`" + `weekday-and-time-of-day` + "`" + `. ` + "`" + `restrictions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "start_day",
					Description: `(Required) Value of the day that frame will start.`,
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
					Name:        "end_day",
					Description: `(Required) Value of the day that frame will end.`,
				},
				resource.Attribute{
					Name:        "end_hour",
					Description: `(Required) Value of the hour that frame will end.`,
				},
				resource.Attribute{
					Name:        "end_min",
					Description: `(Required) Value of the minute that frame will end. Minutes may take 0 or 30 as value. Otherwise they will be converted to nearest 0 or 30 automatically. Both ` + "`" + `start_day` + "`" + ` and ` + "`" + `end_day` + "`" + ` can assume only ` + "`" + `monday` + "`" + `, ` + "`" + `tuesday` + "`" + `, ` + "`" + `wednesday` + "`" + `, ` + "`" + `thursday` + "`" + `, ` + "`" + `friday` + "`" + `, ` + "`" + `saturday` + "`" + `, or ` + "`" + `sunday` + "`" + ` values. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Schedule Rotation ## Import Schedule Rotations can be imported using the ` + "`" + `schedule_id/rotation_id` + "`" + `, e.g.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Schedule Rotation ## Import Schedule Rotations can be imported using the ` + "`" + `schedule_id/rotation_id` + "`" + `, e.g.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_service",
			Category:         "Resources",
			ShortDescription: `Manages a Service within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the service. This field must not be longer than 100 characters.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) Team id of the service. This field must not be longer than 512 characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description field of the service that is generally used to provide a detailed information about the service. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Service. ## Import Teams can be imported using the ` + "`" + `service_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_service.this service_id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Service. ## Import Teams can be imported using the ` + "`" + `service_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_service.this service_id` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_service_incident_rule",
			Category:         "Resources",
			ShortDescription: `Manages a Service Incident Rule within Opsgenie.`,
			Description:      ``,
			Keywords: []string{
				"service",
				"incident",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_id",
					Description: `(Required) ID of the service associated`,
				},
				resource.Attribute{
					Name:        "incident_rule",
					Description: `(Required) This is the rule configuration for this incident rule. This is a block, structure is documented below. The ` + "`" + `incident_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "condition_match_type",
					Description: `(Optional) A Condition type, supported types are: ` + "`" + `match-all` + "`" + `, ` + "`" + `match-any-condition` + "`" + `, ` + "`" + `match-all-conditions` + "`" + `. Default: ` + "`" + `match-all` + "`" + ``,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional) Conditions applied to incident. This is a block, structure is documented below.`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `(Required) Specifies which alert field will be used in condition. Possible values are ` + "`" + `message` + "`" + `, ` + "`" + `alias` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `source` + "`" + `, ` + "`" + `entity` + "`" + `, ` + "`" + `tags` + "`" + `, ` + "`" + `actions` + "`" + `, ` + "`" + `details` + "`" + `, ` + "`" + `extra-properties` + "`" + `, ` + "`" + `recipients` + "`" + `, ` + "`" + `teams` + "`" + `, ` + "`" + `priority` + "`" + ``,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `(Required) It is the operation that will be executed for the given field and key. Possible operations are ` + "`" + `matches` + "`" + `, ` + "`" + `contains` + "`" + `, ` + "`" + `starts-with` + "`" + `, ` + "`" + `ends-with` + "`" + `, ` + "`" + `equals` + "`" + `, ` + "`" + `contains-key` + "`" + `, ` + "`" + `contains-value` + "`" + `, ` + "`" + `greater-than` + "`" + `, ` + "`" + `less-than` + "`" + `, ` + "`" + `is-empty` + "`" + `, ` + "`" + `equals-ignore-whitespace` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "not",
					Description: `(Optional) Indicates behaviour of the given operation. Default: false`,
				},
				resource.Attribute{
					Name:        "expected_value",
					Description: `(Optional) User defined value that will be compared with alert field according to the operation. Default: empty string The ` + "`" + `incident_properties` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `(Required) Message of the related incident rule.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags of the alert.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `(Optional) Map of key-value pairs to use as custom properties of the alert.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description field of the incident rule.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) Priority level of the alert. Possible values are ` + "`" + `P1` + "`" + `, ` + "`" + `P2` + "`" + `, ` + "`" + `P3` + "`" + `, ` + "`" + `P4` + "`" + ` and ` + "`" + `P5` + "`" + ``,
				},
				resource.Attribute{
					Name:        "stakeholder_properties",
					Description: `(Required) DEtails about stakeholders for this rule. This is a block, structure is documented below. The ` + "`" + `stakeholder_properties` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Option to enable stakeholder notifications.Default value is true.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `(Required) Message that is to be passed to audience that is generally used to provide a content information about the alert.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description that is generally used to provide a detailed information about the alert. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Service Incident Policy. ## Import Service Incident Rule can be imported using the ` + "`" + `service_id/service_incident_rule_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_service_incident_rule.this service_id/service_incident_rule_id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Service Incident Policy. ## Import Service Incident Rule can be imported using the ` + "`" + `service_id/service_incident_rule_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_service_incident_rule.this service_id/service_incident_rule_id` + "`" + ``,
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
					Name:        "ignore_members",
					Description: `(Optional) Set to true to ignore any configured member blocks and any team member added/updated/removed via OpsGenie web UI. Use this option e.g. to maintain membership via web UI only and use it only for new teams. Changing the value for existing teams might lead to strange behaviour. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delete_default_resources",
					Description: `(Optional) Set to true to remove default escalation and schedule for newly created team.`,
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
					Description: `(Optional) The role for the user within the Team - can be either ` + "`" + `admin` + "`" + ` or ` + "`" + `user` + "`" + `. Default: ` + "`" + `user` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Team. ## Import Teams can be imported using the ` + "`" + `team_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_team.team1 team_id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Team. ## Import Teams can be imported using the ` + "`" + `team_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_team.team1 team_id` + "`" + ``,
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
					Name:        "team_id",
					Description: `(Required) Id of the team owning the routing rule`,
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
					Description: `(Required) Target entity of schedule, escalation, or the reserved word none which will be notified in routing rule. The possible values are: ` + "`" + `schedule` + "`" + `, ` + "`" + `escalation` + "`" + `, ` + "`" + `none` + "`" + ` ` + "`" + `notify` + "`" + ` supports the following:`,
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
					Description: `(Required) Type of the operation will be applied on conditions. Should be one of ` + "`" + `match-all` + "`" + `, ` + "`" + `match-any-condition` + "`" + ` or ` + "`" + `match-all-conditions` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional) List of conditions will be checked before applying team routing rule. This field declaration should be omitted if the criteria type is set to match-all. ` + "`" + `conditions` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `(Required) Specifies which alert field will be used in condition. Possible values are ` + "`" + `message` + "`" + `, ` + "`" + `alias` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `source` + "`" + `, ` + "`" + `entity` + "`" + `, ` + "`" + `tags` + "`" + `, ` + "`" + `actions` + "`" + `, ` + "`" + `extra-properties` + "`" + `, ` + "`" + `recipients` + "`" + `, ` + "`" + `teams` + "`" + ` or ` + "`" + `priority` + "`" + `.`,
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
					Description: `(true) It is the operation that will be executed for the given field and key. Possible operations are ` + "`" + `matches` + "`" + `, ` + "`" + `contains` + "`" + `, ` + "`" + `starts-with` + "`" + `, ` + "`" + `ends-with` + "`" + `, ` + "`" + `equals` + "`" + `, ` + "`" + `contains-key` + "`" + `, ` + "`" + `contains-value` + "`" + `, ` + "`" + `greater-than` + "`" + `, ` + "`" + `less-than` + "`" + `, ` + "`" + `is-empty` + "`" + ` and ` + "`" + `equals-ignore-whitespace` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expectedValue",
					Description: `(Optional) User defined value that will be compared with alert field according to the operation. Default: empty string.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Order of the condition in conditions list. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Team Routing Rule. ## Import Team Routing Rules can be imported using the ` + "`" + `team_id/routing_rule_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_team_routing_rule.ruletest team_id/routing_rule_id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Team Routing Rule. ## Import Team Routing Rules can be imported using the ` + "`" + `team_id/routing_rule_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_team_routing_rule.ruletest team_id/routing_rule_id` + "`" + ``,
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
					Description: `(Required) The email address associated with this user. Opsgenie defines that this must not be longer than 100 characters and must contain lowercase characters only.`,
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
					Description: `(Optional) Timezone information of the user. Please look at [Supported Timezone Ids](https://docs.opsgenie.com/docs/supported-timezone-ids) for available timezones.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags to be associated with the user.`,
				},
				resource.Attribute{
					Name:        "skype_username",
					Description: `(Optional) Skype username of the user.`,
				},
				resource.Attribute{
					Name:        "user_details",
					Description: `(Optional) Details about the user in form of key and list. of values.`,
				},
				resource.Attribute{
					Name:        "user_address",
					Description: `(Optional) Address of the user. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie User. ## Import Users can be imported using the ` + "`" + `user_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_user.user user_id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie User. ## Import Users can be imported using the ` + "`" + `user_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_user.user user_id` + "`" + ``,
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
					Description: `The ID of the Opsgenie Contact. ## Import Users can be imported using the ` + "`" + `username/contact_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_user_contact.testcontact username/contact_id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Contact. ## Import Users can be imported using the ` + "`" + `username/contact_id` + "`" + `, e.g. ` + "`" + `$ terraform import opsgenie_user_contact.testcontact username/contact_id` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"opsgenie_alert_policy":          0,
		"opsgenie_api_integration":       1,
		"opsgenie_custom_role":           2,
		"opsgenie_email_integration":     3,
		"opsgenie_escalation":            4,
		"opsgenie_heartbeat":             5,
		"opsgenie_incident_template":     6,
		"opsgenie_integration_action":    7,
		"opsgenie_maintenance":           8,
		"opsgenie_notification_policy":   9,
		"opsgenie_notification_rule":     10,
		"opsgenie_schedule":              11,
		"opsgenie_schedule_rotation":     12,
		"opsgenie_service":               13,
		"opsgenie_service_incident_rule": 14,
		"opsgenie_team":                  15,
		"opsgenie_team_routing_rule":     16,
		"opsgenie_user":                  17,
		"opsgenie_user_contact":          18,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
