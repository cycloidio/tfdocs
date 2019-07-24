package newrelic

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "newrelic_alert_channel",
			Category:         "Resources",
			ShortDescription: `Create and manage a notification channel for alerts in New Relic.`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the channel.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of channel. One of: ` + "`" + `campfire` + "`" + `, ` + "`" + `email` + "`" + `, ` + "`" + `hipchat` + "`" + `, ` + "`" + `opsgenie` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `victorops` + "`" + `, or ` + "`" + `webhook` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Required) A map of key / value pairs with channel type specific values. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the channel. ## Import Alert channels can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_alert_channel.main 12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the channel. ## Import Alert channels can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_alert_channel.main 12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_alert_condition",
			Category:         "Resources",
			ShortDescription: `Create and manage an alert condition for a policy in New Relic.`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"condition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The ID of the policy where this condition should be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The title of the condition. Must be between 1 and 64 characters, inclusive.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of condition. One of: ` + "`" + `apm_app_metric` + "`" + `, ` + "`" + `apm_jvm_metric` + "`" + `, ` + "`" + `apm_kt_metric` + "`" + `, ` + "`" + `servers_metric` + "`" + `, ` + "`" + `browser_metric` + "`" + `, ` + "`" + `mobile_metric` + "`" + ``,
				},
				resource.Attribute{
					Name:        "entities",
					Description: `(Required) The instance IDS associated with this condition.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Required) The metric field accepts parameters based on the ` + "`" + `type` + "`" + ` set.`,
				},
				resource.Attribute{
					Name:        "gc_metric",
					Description: `(Optional) A valid Garbage Collection metric e.g. ` + "`" + `GC/G1 Young Generation` + "`" + `. This is required if you are using ` + "`" + `apm_jvm_metric` + "`" + ` with ` + "`" + `gc_cpu_time` + "`" + ` condition type.`,
				},
				resource.Attribute{
					Name:        "violation_close_timer",
					Description: `(Optional) Automatically close instance-based violations, including JVM health metric violations, after the number of hours specified. Must be: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `12` + "`" + ` or ` + "`" + `24` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "runbook_url",
					Description: `(Optional) Runbook URL to display in notifications.`,
				},
				resource.Attribute{
					Name:        "condition_scope",
					Description: `(Optional) ` + "`" + `instance` + "`" + ` or ` + "`" + `application` + "`" + `. This is required if you are using the JVM plugin in New Relic.`,
				},
				resource.Attribute{
					Name:        "term",
					Description: `(Required) A list of terms for this condition. See [Terms](#terms) below for details.`,
				},
				resource.Attribute{
					Name:        "user_defined_metric",
					Description: `(Optional) A custom metric to be evaluated.`,
				},
				resource.Attribute{
					Name:        "user_defined_value_function",
					Description: `(Optional) One of: ` + "`" + `average` + "`" + `, ` + "`" + `min` + "`" + `, ` + "`" + `max` + "`" + `, ` + "`" + `total` + "`" + `, or ` + "`" + `sample_size` + "`" + `. ## Terms The ` + "`" + `term` + "`" + ` mapping supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) In minutes, must be: ` + "`" + `5` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `15` + "`" + `, ` + "`" + `30` + "`" + `, ` + "`" + `60` + "`" + `, or ` + "`" + `120` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) ` + "`" + `above` + "`" + `, ` + "`" + `below` + "`" + `, or ` + "`" + `equal` + "`" + `. Defaults to ` + "`" + `equal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) ` + "`" + `critical` + "`" + ` or ` + "`" + `warning` + "`" + `. Defaults to ` + "`" + `critical` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) Must be 0 or greater.`,
				},
				resource.Attribute{
					Name:        "time_function",
					Description: `(Required) ` + "`" + `all` + "`" + ` or ` + "`" + `any` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert condition. ## Import Alert conditions can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_alert_condition.main 12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert condition. ## Import Alert conditions can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_alert_condition.main 12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_alert_policy",
			Category:         "Resources",
			ShortDescription: `Create and manage alert policies in New Relic.`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the policy.`,
				},
				resource.Attribute{
					Name:        "incident_preference",
					Description: `(Optional) The rollup strategy for the policy. Options include: ` + "`" + `PER_POLICY` + "`" + `, ` + "`" + `PER_CONDITION` + "`" + `, or ` + "`" + `PER_CONDITION_AND_TARGET` + "`" + `. The default is ` + "`" + `PER_POLICY` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the policy.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time the policy was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time the policy was last updated. ## Import Alert policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_alert_policy.main 12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the policy.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time the policy was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time the policy was last updated. ## Import Alert policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_alert_policy.main 12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_alert_policy_channel",
			Category:         "Resources",
			ShortDescription: `Map alert policies to alert channels in New Relic.`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"policy",
				"channel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The ID of the policy.`,
				},
				resource.Attribute{
					Name:        "channel_id",
					Description: `(Required) The ID of the channel.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_dashboard",
			Category:         "Resources",
			ShortDescription: `Create and manage dashboards in New Relic.`,
			Description:      ``,
			Keywords: []string{
				"dashboard",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "title",
					Description: `(Required) The title of the dashboard.`,
				},
				resource.Attribute{
					Name:        "icon",
					Description: `(Optional) The icon for the dashboard. Defaults to ` + "`" + `bar-chart` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) Who can see the dashboard in an account. Must be ` + "`" + `owner` + "`" + ` or ` + "`" + `all` + "`" + `. Defaults to ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "widget",
					Description: `(Optional) A widget that describes a visualization. See [Widgets](#widgets) below for details.`,
				},
				resource.Attribute{
					Name:        "editable",
					Description: `(Optional) Who can edit the dashboard in an account. Must be ` + "`" + `read_only` + "`" + `, ` + "`" + `editable_by_owner` + "`" + `, ` + "`" + `editable_by_all` + "`" + `, or ` + "`" + `all` + "`" + `. Defaults to ` + "`" + `editable_by_all` + "`" + `. ## Widgets The ` + "`" + `widget` + "`" + ` mapping supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) A title for the widget.`,
				},
				resource.Attribute{
					Name:        "visualization",
					Description: `(Required) How the widget visualizes data.`,
				},
				resource.Attribute{
					Name:        "row",
					Description: `(Required) Row position of widget from top left, starting at ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "column",
					Description: `(Required) Column position of widget from top left, starting at ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `(Optional) Width of the widget. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `(Optional) Height of the widget. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Description of the widget.`,
				},
				resource.Attribute{
					Name:        "nrql",
					Description: `(Optional) Valid NRQL query string. See [Writing NRQL Queries](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/using-nrql/introduction-nrql) for help. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the dashboard.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the dashboard.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_infra_alert_condition",
			Category:         "Resources",
			ShortDescription: `Create and manage an Infrastructure alert condition for a policy in New Relic.`,
			Description:      ``,
			Keywords: []string{
				"infra",
				"alert",
				"condition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The ID of the alert policy where this condition should be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Infrastructure alert condition's name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Set whether to enable the alert condition. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of Infrastructure alert condition: "infra_process_running", "infra_metric", or "infra_host_not_reporting".`,
				},
				resource.Attribute{
					Name:        "event",
					Description: `(Required) The metric event; for example, system metrics, process metrics, storage metrics, or network metrics.`,
				},
				resource.Attribute{
					Name:        "select",
					Description: `(Required) The attribute name to identify the type of metric condition; for example, "network", "process", "system", or "storage".`,
				},
				resource.Attribute{
					Name:        "comparison",
					Description: `(Required) The operator used to evaluate the threshold value; "above", "below", "equal".`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Required) Identifies the critical threshold parameters for triggering an alert notification. See [Thresholds](#thresholds) below for details.`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Optional) Identifies the warning threshold parameters. See [Thresholds](#thresholds) below for details.`,
				},
				resource.Attribute{
					Name:        "where",
					Description: `(Optional) Infrastructure host filter for the alert condition.`,
				},
				resource.Attribute{
					Name:        "process_where",
					Description: `(Optional) Any filters applied to processes; for example: ` + "`" + `"commandName = 'java'"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "integration_provider",
					Description: `(Optional) For alerts on integrations, use this instead of ` + "`" + `event` + "`" + `. ## Thresholds The ` + "`" + `critical` + "`" + ` and ` + "`" + `warning` + "`" + ` threshold mapping supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) Identifies the number of minutes the threshold must be passed or met for the alert to trigger. Threshold durations must be between 1 and 60 minutes (inclusive).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Threshold value, computed against the ` + "`" + `comparison` + "`" + ` operator. Supported by "infra_metric" and "infra_process_running" alert condition types.`,
				},
				resource.Attribute{
					Name:        "time_function",
					Description: `(Optional) Indicates if the condition needs to be sustained or to just break the threshold once; ` + "`" + `all` + "`" + ` or ` + "`" + `any` + "`" + `. Supported by the "infra_metric" alert condition type. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Infrastructure alert condition.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Infrastructure alert condition.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_nrql_alert_condition",
			Category:         "Resources",
			ShortDescription: `Create and manage a NRQL alert condition for a policy in New Relic.`,
			Description:      ``,
			Keywords: []string{
				"nrql",
				"alert",
				"condition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The ID of the policy where this condition should be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The title of the condition`,
				},
				resource.Attribute{
					Name:        "runbook_url",
					Description: `(Optional) Runbook URL to display in notifications.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Set whether to enable the alert condition. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "term",
					Description: `(Required) A list of terms for this condition. See [Terms](#terms) below for details.`,
				},
				resource.Attribute{
					Name:        "nrql",
					Description: `(Required) A NRQL query. See [NRQL](#nrql) below for details.`,
				},
				resource.Attribute{
					Name:        "value_function",
					Description: `(Optional) Possible values are ` + "`" + `single_value` + "`" + `, ` + "`" + `sum` + "`" + `. ## Terms The ` + "`" + `term` + "`" + ` mapping supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) In minutes, must be: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `15` + "`" + `, ` + "`" + `30` + "`" + `, ` + "`" + `60` + "`" + `, or ` + "`" + `120` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) ` + "`" + `above` + "`" + `, ` + "`" + `below` + "`" + `, or ` + "`" + `equal` + "`" + `. Defaults to ` + "`" + `equal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) ` + "`" + `critical` + "`" + ` or ` + "`" + `warning` + "`" + `. Defaults to ` + "`" + `critical` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) Must be 0 or greater.`,
				},
				resource.Attribute{
					Name:        "time_function",
					Description: `(Required) ` + "`" + `all` + "`" + ` or ` + "`" + `any` + "`" + `. ## NRQL The ` + "`" + `nrql` + "`" + ` attribute supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) The NRQL query to execute for the condition.`,
				},
				resource.Attribute{
					Name:        "since_value",
					Description: `(Required) The value to be used in the ` + "`" + `SINCE <X> MINUTES AGO` + "`" + ` clause for the NRQL query. Must be between ` + "`" + `1` + "`" + ` and ` + "`" + `20` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NRQL alert condition. ## Import Alert conditions can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_nrql_alert_condition.main 12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NRQL alert condition. ## Import Alert conditions can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_nrql_alert_condition.main 12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_synthetics_alert_condition",
			Category:         "Resources",
			ShortDescription: `Create and manage a Synthetics alert condition for a policy in New Relic.`,
			Description:      ``,
			Keywords: []string{
				"synthetics",
				"alert",
				"condition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The ID of the policy where this condition should be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The title of this condition.`,
				},
				resource.Attribute{
					Name:        "monitor_id",
					Description: `(Required) The ID of the Synthetics monitor to be referenced in the alert condition.`,
				},
				resource.Attribute{
					Name:        "runbook_url",
					Description: `(Optional) Runbook URL to display in notifications.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Set whether to enable the alert condition. Defaults to ` + "`" + `true` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Synthetics alert condition.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Synthetics alert condition.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_synthetics_monitor",
			Category:         "Resources",
			ShortDescription: `Create and manage a Synthetics monitor in New Relic.`,
			Description:      ``,
			Keywords: []string{
				"synthetics",
				"monitor",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The title of this monitor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The monitor type.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `(Required) The interval (in minutes) at which this monitor should run.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Required) The monitor status (i.e. ENABLED, MUTED, DISABLED)`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(Required) The locations in which this monitor should be run.`,
				},
				resource.Attribute{
					Name:        "sla_threshold",
					Description: `(Optional) The base threshold for the SLA report. For SIMPLE and BROWSER monitor types, the following arguments are also supported:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Required) The URI for the monitor to hit.`,
				},
				resource.Attribute{
					Name:        "validation_string",
					Description: `(Optional) The string to validate against in the response.`,
				},
				resource.Attribute{
					Name:        "verify_ssl",
					Description: `(Optional) Verify SSL.`,
				},
				resource.Attribute{
					Name:        "bypass_head_request",
					Description: `(Optional) Bypass HEAD request.`,
				},
				resource.Attribute{
					Name:        "treat_redirect_as_faiure",
					Description: `(Optional) Fail the monitor check if redirected. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Synthetics monitor.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Synthetics monitor.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_synthetics_monitor_script",
			Category:         "Resources",
			ShortDescription: `Update and manage a Synthetics monitor script in New Relic.`,
			Description:      ``,
			Keywords: []string{
				"synthetics",
				"monitor",
				"script",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "monitor_id",
					Description: `(Required) The ID of the monitor to attach the script to.`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `(Required) plaintext of the monitor script. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Synthetics monitor that the script is attached to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Synthetics monitor that the script is attached to.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"newrelic_alert_channel":              0,
		"newrelic_alert_condition":            1,
		"newrelic_alert_policy":               2,
		"newrelic_alert_policy_channel":       3,
		"newrelic_dashboard":                  4,
		"newrelic_infra_alert_condition":      5,
		"newrelic_nrql_alert_condition":       6,
		"newrelic_synthetics_alert_condition": 7,
		"newrelic_synthetics_monitor":         8,
		"newrelic_synthetics_monitor_script":  9,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
