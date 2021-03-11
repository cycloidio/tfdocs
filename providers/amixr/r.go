package amixr

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "amixr_escalation",
			Category:         "Resources",
			ShortDescription: `Configures escalation policies in Amixr.`,
			Description:      ``,
			Keywords: []string{
				"escalation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route_id",
					Description: `(Required) The ID of the route.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Required) The position of the escalation step (starts from 0)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of escalation policy. Can be: - ` + "`" + `wait` + "`" + ` - just wait - ` + "`" + `notify_persons` + "`" + ` - notify multiple users at the same time - ` + "`" + `notify_person_next_each_time` + "`" + ` - notify one user from queue - ` + "`" + `notify_on_call_from_schedule` + "`" + ` - notify by schedule - ` + "`" + `notify_user_group` + "`" + ` - notify User Group (available for teams with Slack integration) - ` + "`" + `trigger_action` + "`" + ` - trigger action (outgoing webhook) - ` + "`" + `notify_if_time_from_to` + "`" + ` - continue escalation only if the time is within the selected interval - ` + "`" + `notify_whole_channel` + "`" + ` - notify a channel in Slack (available for teams with Slack integration) - ` + "`" + `resolve` + "`" + ` - resolve incident - ` + "`" + `null` + "`" + ` - do nothing`,
				},
				resource.Attribute{
					Name:        "important",
					Description: `(Optional) Will activate "important" personal notification rules. Can be ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Actual for steps: ` + "`" + `notify_persons` + "`" + `, ` + "`" + `notify_on_call_from_schedule` + "`" + ` and ` + "`" + `notify_user_group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) The duration of delay for ` + "`" + `wait` + "`" + ` type step.`,
				},
				resource.Attribute{
					Name:        "persons_to_notify",
					Description: `(Optional) The list of ID's of users for ` + "`" + `notify_persons` + "`" + ` type step.`,
				},
				resource.Attribute{
					Name:        "persons_to_notify_next_each_time",
					Description: `(Optional) The list of ID's of users for ` + "`" + `notify_person_next_each_time` + "`" + ` type step.`,
				},
				resource.Attribute{
					Name:        "notify_on_call_from_schedule",
					Description: `(Optional) ID of a Schedule for ` + "`" + `notify_on_call_from_schedule` + "`" + ` type step.`,
				},
				resource.Attribute{
					Name:        "action_to_trigger",
					Description: `(Optional) ID of an Action for ` + "`" + `trigger_action` + "`" + ` type step.`,
				},
				resource.Attribute{
					Name:        "group_to_notify",
					Description: `(Optional) ID of a User Group for ` + "`" + `notify_user_group` + "`" + ` type step.`,
				},
				resource.Attribute{
					Name:        "notify_if_time_from",
					Description: `(Optional) The beginning of the time interval for ` + "`" + `notify_if_time_from_to` + "`" + ` type step in UTC (for example 08:00:00Z).`,
				},
				resource.Attribute{
					Name:        "notify_if_time_to",
					Description: `(Optional) The end of the time interval for ` + "`" + `notify_if_time_from_to` + "`" + ` type step in UTC (for example 18:00:00Z). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the escalation policy. ## Import Existing escalations can be imported using the escalation ID: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import amixr_escalation.example_escalation {escalation ID} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the escalation policy. ## Import Existing escalations can be imported using the escalation ID: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import amixr_escalation.example_escalation {escalation ID} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "amixr_integration",
			Category:         "Resources",
			ShortDescription: `Creates and manages an integration in Amixr.`,
			Description:      ``,
			Keywords: []string{
				"integration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the service integration.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of integration. Can be: ` + "`" + `grafana` + "`" + `, ` + "`" + `webhook` + "`" + `, ` + "`" + `alertmanager` + "`" + `, ` + "`" + `kapacitor` + "`" + ` , ` + "`" + `fabric` + "`" + `, ` + "`" + `newrelic` + "`" + `, ` + "`" + `datadog` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `pingdom` + "`" + `, ` + "`" + `elastalert` + "`" + `, ` + "`" + `amazon_sns` + "`" + `, ` + "`" + `curler` + "`" + `, ` + "`" + `sentry` + "`" + `, ` + "`" + `formatted_webhook` + "`" + `, ` + "`" + `heartbeat` + "`" + `, ` + "`" + `demo` + "`" + `, ` + "`" + `stackdriver` + "`" + `, ` + "`" + `uptimerobot` + "`" + `, ` + "`" + `sentry_platform` + "`" + `, ` + "`" + `zabbix` + "`" + `, ` + "`" + `prtg` + "`" + ` or ` + "`" + `inbound_email` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "templates",
					Description: `(Optional) Jinja2 templates for Alert payload. Includes: - ` + "`" + `grouping_key` + "`" + `- (Optional) Template for the key by which alerts are grouped. - ` + "`" + `resolve_signal` + "`" + `- (Optional) Template for sending a signal to resolve the Incident. This template should output one of the following values: ok, true, 1 (case insensitive) - ` + "`" + `slack` + "`" + `- (Optional) Templates for Slack: - ` + "`" + `title` + "`" + `- (Optional) Template for Alert title. - ` + "`" + `message` + "`" + `- (Optional) Template for Alert message. - ` + "`" + `image_url` + "`" + `- (Optional) Template for Alert image url. To set a parameter to default value just remove it or set the value to ` + "`" + `null` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the integration.`,
				},
				resource.Attribute{
					Name:        "link",
					Description: `Link for using in an integrated tool. ## Import Existing integrations can be imported using the integration ID: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import amixr_integration.example_integration {integration ID} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the integration.`,
				},
				resource.Attribute{
					Name:        "link",
					Description: `Link for using in an integrated tool. ## Import Existing integrations can be imported using the integration ID: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import amixr_integration.example_integration {integration ID} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "amixr_on_call_shift",
			Category:         "Resources",
			ShortDescription: `Creates and manages schedule on-call shifts in Amixr.`,
			Description:      ``,
			Keywords: []string{
				"on",
				"call",
				"shift",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The shift's name`,
				},
				resource.Attribute{
					Name:        "schedule_id",
					Description: `(Required) The ID of the schedule`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The shift's type. Can be: - ` + "`" + `single_event` + "`" + ` - the event will be triggered once and does not repeat - ` + "`" + `recurrent_event` + "`" + ` - the event will be repeated in accordance with the recurrence rules - ` + "`" + `rolling_users` + "`" + ` - the event will be repeated in accordance with the recurrence rules, rolling users for each shift`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `(Optional) - The priority level. The higher the value, the higher the priority. Within one schedule if two events have an overlap in time Amixr will choose the event with higher level.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) The start time of the on-call shift. This parameter takes a date format as ` + "`" + `yyyy-MM-dd'T'HH:mm:ss` + "`" + ` (for example "2020-09-05T08:00:00")`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) The duration of the event`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `(Required for recurrent events) The frequency of the event. Can be: ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `, ` + "`" + `monthly` + "`" + ``,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) This parameter takes a positive integer representing at which intervals the recurrence rule repeats`,
				},
				resource.Attribute{
					Name:        "week_start",
					Description: `(Optional) Start day of the week in iCal format. Can be: ` + "`" + `SU` + "`" + ` (Sunday), ` + "`" + `MO` + "`" + ` (Monday), ` + "`" + `TU` + "`" + ` (Tuesday), ` + "`" + `WE` + "`" + ` (Wednesday), ` + "`" + `TH` + "`" + ` (Thursday), ` + "`" + `FR` + "`" + ` (Friday), ` + "`" + `SA` + "`" + ` (Saturday). Default: ` + "`" + `SU` + "`" + ``,
				},
				resource.Attribute{
					Name:        "by_day",
					Description: `(Optional) This parameter takes a list of days in iCal format (` + "`" + `SU` + "`" + `, ` + "`" + `MO` + "`" + `, ` + "`" + `TU` + "`" + `, ` + "`" + `WE` + "`" + `, ` + "`" + `TH` + "`" + `, ` + "`" + `FR` + "`" + `, ` + "`" + `SA` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "by_month",
					Description: `(Optional) This parameter takes a list of months. Valid values are 1 to 12`,
				},
				resource.Attribute{
					Name:        "by_monthday",
					Description: `(Optional) This parameter takes a list of days of the month. Valid values are 1 to 31 or -31 to -1`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) The list of on-call users (for ` + "`" + `single_event` + "`" + ` and ` + "`" + `recurrent_event` + "`" + ` event type)`,
				},
				resource.Attribute{
					Name:        "rolling_users",
					Description: `(Optional) The list of lists with on-call users (for ` + "`" + `rolling_users` + "`" + ` event type) Please check [RFC 5545](https://tools.ietf.org/html/rfc5545#section-3.3.10) for more information about the recurrence rules. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the on-call shift. ## Import Existing on-call shift can be imported using the shift ID: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import amixr_on_call_shift.example_shift {on-call shift ID} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the on-call shift. ## Import Existing on-call shift can be imported using the shift ID: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import amixr_on_call_shift.example_shift {on-call shift ID} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "amixr_route",
			Category:         "Resources",
			ShortDescription: `Creates and manages integration routes in Amixr.`,
			Description:      ``,
			Keywords: []string{
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "integration_id",
					Description: `(Required) The ID of the integration.`,
				},
				resource.Attribute{
					Name:        "routing_regex",
					Description: `(Required) Python Regex query. Amixr choose the route for an alert in case there is a match inside the whole alert payload.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Required) The position of the route (starts from 0)`,
				},
				resource.Attribute{
					Name:        "slack",
					Description: `(Optional) Dictionary with slack-specific settings for a route. Includes: - ` + "`" + `channel_id` + "`" + ` - Slack channel id. Alerts will be directed to this channel in Slack. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the route. ## Import Existing routes can be imported using the route ID: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import amixr_route.example_route {route ID} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the route. ## Import Existing routes can be imported using the route ID: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import amixr_route.example_route {route ID} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "amixr_schedule",
			Category:         "Resources",
			ShortDescription: `Creates and manages schedules in Amixr.`,
			Description:      ``,
			Keywords: []string{
				"schedule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The schedule's name`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Optional) The schedule's time zone. More information about [time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).`,
				},
				resource.Attribute{
					Name:        "ical_url",
					Description: `(Optional) The URL of the external calendar iCal file`,
				},
				resource.Attribute{
					Name:        "slack",
					Description: `(Optional) Dictionary with slack-specific settings for a schedule. Includes: - ` + "`" + `channel_id` + "`" + ` - Slack channel id. Reminder about schedule shifts will be directed to this channel in Slack ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the schedule. ## Import Existing schedule can be imported using the schedule ID: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import amixr_schedule.example_schedule {schedule ID} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the schedule. ## Import Existing schedule can be imported using the schedule ID: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import amixr_schedule.example_schedule {schedule ID} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"amixr_escalation":    0,
		"amixr_integration":   1,
		"amixr_on_call_shift": 2,
		"amixr_route":         3,
		"amixr_schedule":      4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
