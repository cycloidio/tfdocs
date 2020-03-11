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
			Description: `\_alert\_channel

Use this resource to create and manage New Relic alert policies.

`,
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
					Description: `(Required) The type of channel. One of: ` + "`" + `email` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `opsgenie` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `victorops` + "`" + `, or ` + "`" + `webhook` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Optional) A nested block that describes an alert channel configuration. Only one config block is permitted per alert channel definition. See [Nested config blocks](#nested-` + "`" + `config` + "`" + `-blocks) below for details.`,
				},
				resource.Attribute{
					Name:        "recipients",
					Description: `(Required) Comma delimited list of email addresses.`,
				},
				resource.Attribute{
					Name:        "include_json_attachment",
					Description: `(Optional) ` + "`" + `0` + "`" + ` or ` + "`" + `1` + "`" + `. Flag for whether or not to attach a JSON document containing information about the associated alert to the email that is sent to recipients.`,
				},
				resource.Attribute{
					Name:        "base_url",
					Description: `(Required) The base URL of the webhook destination.`,
				},
				resource.Attribute{
					Name:        "auth_password",
					Description: `(Optional) Specifies an authentication password for use with a channel. Supported by the ` + "`" + `webhook` + "`" + ` channel type.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) Specifies an authentication method for use with a channel. Supported by the ` + "`" + `webhook` + "`" + ` channel type. Only HTTP basic authentication is currently supported via the value ` + "`" + `BASIC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_username",
					Description: `(Optional) Specifies an authentication username for use with a channel. Supported by the ` + "`" + `webhook` + "`" + ` channel type.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `(Optional) A map of key/value pairs that represents extra HTTP headers to be sent along with the webhook payload.`,
				},
				resource.Attribute{
					Name:        "headers_string",
					Description: `(Optional) Use instead of ` + "`" + `headers` + "`" + ` if the desired payload is more complex than a list of key/value pairs (e.g. a set of headers that makes use of nested objects). The value provided should be a valid JSON string with escaped double quotes. Conflicts with ` + "`" + `headers` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "payload",
					Description: `(Optional) A map of key/value pairs that represents the webhook payload. Must provide ` + "`" + `payload_type` + "`" + ` if setting this argument.`,
				},
				resource.Attribute{
					Name:        "payload_string",
					Description: `(Optional) Use instead of ` + "`" + `payload` + "`" + ` if the desired payload is more complex than a list of key/value pairs (e.g. a payload that makes use of nested objects). The value provided should be a valid JSON string with escaped double quotes. Conflicts with ` + "`" + `payload` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "payload_type",
					Description: `(Optional) Can either be ` + "`" + `application/json` + "`" + ` or ` + "`" + `application/x-www-form-urlencoded` + "`" + `. The ` + "`" + `payload_type` + "`" + ` argument is _required_ if ` + "`" + `payload` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `(Required) Specifies the service key for integrating with Pagerduty.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key for integrating with VictorOps.`,
				},
				resource.Attribute{
					Name:        "route_key",
					Description: `(Required) The route key for integrating with VictorOps.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Your organization's Slack URL.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Optional) The Slack channel to send notifications to.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `(Required) The API key for integrating with OpsGenie.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The data center region to store your data. Valid values are ` + "`" + `US` + "`" + ` and ` + "`" + `EU` + "`" + `. Default is ` + "`" + `US` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `(Optional) A set of teams for targeting notifications. Multiple values are comma separated.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of tags for targeting notifications. Multiple values are comma separated.`,
				},
				resource.Attribute{
					Name:        "recipients",
					Description: `(Optional) A set of recipients for targeting notifications. Multiple values are comma separated. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the channel. ## Additional Examples ##### Slack ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_alert_channel" "foo" { name = "slack-example" type = "slack" config { url = "https://<YourOrganization>.slack.com" channel = "example-alerts-channel" } } ` + "`" + `` + "`" + `` + "`" + ` ##### OpsGenie ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_alert_channel" "foo" { name = "opsgenie-example" type = "opsgenie" config { api_key = "abc123" teams = "team1, team2" tags = "tag1, tag2" recipients = "user1@domain.com, user2@domain.com" } } ` + "`" + `` + "`" + `` + "`" + ` ##### PagerDuty ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_alert_channel" "foo" { name = "pagerduty-example" type = "pagerduty" config { service_key = "abc123" } } ` + "`" + `` + "`" + `` + "`" + ` ##### VictorOps ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_alert_channel" "foo" { name = "victorops-example" type = "victorops" config { key = "abc123" route_key = "/example" } } ` + "`" + `` + "`" + `` + "`" + ` ##### Webhook ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_alert_channel" "foo" { name = "webhook-example" type = "webhook" config { base_url = "http://www.test.com" payload_type = "application/json" payload = { condition_name = "$CONDITION_NAME" policy_name = "$POLICY_NAME" } headers = { header1 = value1 header2 = value2 } } } ` + "`" + `` + "`" + `` + "`" + ` ##### Webhook with complex payload ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_alert_channel" "foo" { name = "webhook-example" type = "webhook" config { base_url = "http://www.test.com" payload_type = "application/json" payload_string = <<EOF { "my_custom_values": { "condition_name": "$CONDITION_NAME", "policy_name": "$POLICY_NAME" } } EOF } } ` + "`" + `` + "`" + `` + "`" + ` ## Import Alert channels can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import newrelic_alert_channel.main <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the channel. ## Additional Examples ##### Slack ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_alert_channel" "foo" { name = "slack-example" type = "slack" config { url = "https://<YourOrganization>.slack.com" channel = "example-alerts-channel" } } ` + "`" + `` + "`" + `` + "`" + ` ##### OpsGenie ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_alert_channel" "foo" { name = "opsgenie-example" type = "opsgenie" config { api_key = "abc123" teams = "team1, team2" tags = "tag1, tag2" recipients = "user1@domain.com, user2@domain.com" } } ` + "`" + `` + "`" + `` + "`" + ` ##### PagerDuty ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_alert_channel" "foo" { name = "pagerduty-example" type = "pagerduty" config { service_key = "abc123" } } ` + "`" + `` + "`" + `` + "`" + ` ##### VictorOps ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_alert_channel" "foo" { name = "victorops-example" type = "victorops" config { key = "abc123" route_key = "/example" } } ` + "`" + `` + "`" + `` + "`" + ` ##### Webhook ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_alert_channel" "foo" { name = "webhook-example" type = "webhook" config { base_url = "http://www.test.com" payload_type = "application/json" payload = { condition_name = "$CONDITION_NAME" policy_name = "$POLICY_NAME" } headers = { header1 = value1 header2 = value2 } } } ` + "`" + `` + "`" + `` + "`" + ` ##### Webhook with complex payload ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_alert_channel" "foo" { name = "webhook-example" type = "webhook" config { base_url = "http://www.test.com" payload_type = "application/json" payload_string = <<EOF { "my_custom_values": { "condition_name": "$CONDITION_NAME", "policy_name": "$POLICY_NAME" } } EOF } } ` + "`" + `` + "`" + `` + "`" + ` ## Import Alert channels can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import newrelic_alert_channel.main <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_alert_condition",
			Category:         "Resources",
			ShortDescription: `Create and manage alert conditions for APM, Browser, and Mobile in New Relic.`,
			Description: `\_alert\_condition

Use this resource to create and manage alert conditions for APM, Browser, and Mobile in New Relic.

`,
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
					Description: `(Required) The metric field accepts parameters based on the ` + "`" + `type` + "`" + ` set. One of these metrics based on ` + "`" + `type` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "condition_scope",
					Description: `(Required for some types) ` + "`" + `application` + "`" + ` or ` + "`" + `instance` + "`" + `. Choose ` + "`" + `application` + "`" + ` for most scenarios. If you are using the JVM plugin in New Relic, the ` + "`" + `instance` + "`" + ` setting allows your condition to trigger [for specific app instances](https://docs.newrelic.com/docs/alerts/new-relic-alerts/defining-conditions/scope-alert-thresholds-specific-instances).`,
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
					Description: `(Required) In minutes, must be in the range of ` + "`" + `5` + "`" + ` to ` + "`" + `120` + "`" + `, inclusive.`,
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
					Description: `(Required) ` + "`" + `all` + "`" + ` or ` + "`" + `any` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert condition. ## Import Alert conditions can be imported using notation ` + "`" + `alert_policy_id:alert_condition_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_alert_condition.main 123456:6789012345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert condition. ## Import Alert conditions can be imported using notation ` + "`" + `alert_policy_id:alert_condition_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_alert_condition.main 123456:6789012345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_alert_policy",
			Category:         "Resources",
			ShortDescription: `Create and manage alert policies in New Relic.`,
			Description: `\_alert\_policy

Use this resource to create and manage New Relic alert policies.

`,
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
					Description: `(Optional) The rollup strategy for the policy. Options include: ` + "`" + `PER_POLICY` + "`" + `, ` + "`" + `PER_CONDITION` + "`" + `, or ` + "`" + `PER_CONDITION_AND_TARGET` + "`" + `. The default is ` + "`" + `PER_POLICY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "channel_ids",
					Description: `(Optional) An array of channel IDs (integers) to assign to the policy. Adding or removing channel IDs from this array will result in a new alert policy resource being created and the old one being destroyed. Also note that channel IDs _cannot_ be imported via ` + "`" + `terraform import` + "`" + ` (see [Import](#import) for info). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `The time the policy was last updated. ## Additional Examples ##### Provision multiple notification channels and add those channels to a policy ` + "`" + `` + "`" + `` + "`" + `hcl # Provision a Slack notification channel. resource "newrelic_alert_channel" "slack_channel" { name = "slack-example" type = "slack" config { url = "https://hooks.slack.com/services/<`,
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
					Description: `The time the policy was last updated. ## Additional Examples ##### Provision multiple notification channels and add those channels to a policy ` + "`" + `` + "`" + `` + "`" + `hcl # Provision a Slack notification channel. resource "newrelic_alert_channel" "slack_channel" { name = "slack-example" type = "slack" config { url = "https://hooks.slack.com/services/<`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_alert_policy_channel",
			Category:         "Resources",
			ShortDescription: `Map alert policies to alert channels in New Relic.`,
			Description: `

Use this resource to map alert policies to alert channels in New Relic.

`,
			Keywords: []string{
				"alert",
				"policy",
				"channel",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_application_label",
			Category:         "Resources",
			ShortDescription: `Create and manage an Application label in New Relic.`,
			Description:      ``,
			Keywords: []string{
				"application",
				"label",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "category",
					Description: `(Required) A string representing the label key/category.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A string that will be assigned to the label.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `(Required) The resources to which label should be assigned to. At least one of the following attributes must be set.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `An array of application IDs.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `An array of server IDs.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_dashboard",
			Category:         "Resources",
			ShortDescription: `Create and manage dashboards in New Relic.`,
			Description: `\_dashboard

Use this resource to create and manage New Relic dashboards.

`,
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
					Description: `(Optional) The icon for the dashboard. Valid values are ` + "`" + `adjust` + "`" + `, ` + "`" + `archive` + "`" + `, ` + "`" + `bar-chart` + "`" + `, ` + "`" + `bell` + "`" + `, ` + "`" + `bolt` + "`" + `, ` + "`" + `bug` + "`" + `, ` + "`" + `bullhorn` + "`" + `, ` + "`" + `bullseye` + "`" + `, ` + "`" + `clock-o` + "`" + `, ` + "`" + `cloud` + "`" + `, ` + "`" + `cog` + "`" + `, ` + "`" + `comments-o` + "`" + `, ` + "`" + `crosshairs` + "`" + `, ` + "`" + `dashboard` + "`" + `, ` + "`" + `envelope` + "`" + `, ` + "`" + `fire` + "`" + `, ` + "`" + `flag` + "`" + `, ` + "`" + `flask` + "`" + `, ` + "`" + `globe` + "`" + `, ` + "`" + `heart` + "`" + `, ` + "`" + `leaf` + "`" + `, ` + "`" + `legal` + "`" + `, ` + "`" + `life-ring` + "`" + `, ` + "`" + `line-chart` + "`" + `, ` + "`" + `magic` + "`" + `, ` + "`" + `mobile` + "`" + `, ` + "`" + `money` + "`" + `, ` + "`" + `none` + "`" + `, ` + "`" + `paper-plane` + "`" + `, ` + "`" + `pie-chart` + "`" + `, ` + "`" + `puzzle-piece` + "`" + `, ` + "`" + `road` + "`" + `, ` + "`" + `rocket` + "`" + `, ` + "`" + `shopping-cart` + "`" + `, ` + "`" + `sitemap` + "`" + `, ` + "`" + `sliders` + "`" + `, ` + "`" + `tablet` + "`" + `, ` + "`" + `thumbs-down` + "`" + `, ` + "`" + `thumbs-up` + "`" + `, ` + "`" + `trophy` + "`" + `, ` + "`" + `usd` + "`" + `, ` + "`" + `user` + "`" + `, and ` + "`" + `users` + "`" + `. Defaults to ` + "`" + `bar-chart` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) Determines who can see the dashboard in an account. Valid values are ` + "`" + `all` + "`" + ` or ` + "`" + `owner` + "`" + `. Defaults to ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "editable",
					Description: `(Optional) Determines who can edit the dashboard in an account. Valid values are ` + "`" + `all` + "`" + `, ` + "`" + `editable_by_all` + "`" + `, ` + "`" + `editable_by_owner` + "`" + `, or ` + "`" + `read_only` + "`" + `. Defaults to ` + "`" + `editable_by_all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "widget",
					Description: `(Optional) A nested block that describes a visualization. Up to 300 ` + "`" + `widget` + "`" + ` blocks are allowed in a dashboard definition. See [Nested widget blocks](#nested-` + "`" + `widget` + "`" + `-blocks) below for details.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A nested block that describes a dashboard filter. Exactly one nested ` + "`" + `filter` + "`" + ` block is allowed. See [Nested filter block](#nested-` + "`" + `filter` + "`" + `-block) below for details. ## Attribute Refence In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dashboard_url",
					Description: `The URL for viewing the dashboard. ### Nested ` + "`" + `widget` + "`" + ` blocks All nested ` + "`" + `widget` + "`" + ` blocks support the following common arguments:`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) A title for the widget.`,
				},
				resource.Attribute{
					Name:        "visualization",
					Description: `(Required) How the widget visualizes data. Valid values are ` + "`" + `billboard` + "`" + `, ` + "`" + `gauge` + "`" + `, ` + "`" + `billboard_comparison` + "`" + `, ` + "`" + `facet_bar_chart` + "`" + `, ` + "`" + `faceted_line_chart` + "`" + `, ` + "`" + `facet_pie_chart` + "`" + `, ` + "`" + `facet_table` + "`" + `, ` + "`" + `faceted_area_chart` + "`" + `, ` + "`" + `heatmap` + "`" + `, ` + "`" + `attribute_sheet` + "`" + `, ` + "`" + `single_event` + "`" + `, ` + "`" + `histogram` + "`" + `, ` + "`" + `funnel` + "`" + `, ` + "`" + `raw_json` + "`" + `, ` + "`" + `event_feed` + "`" + `, ` + "`" + `event_table` + "`" + `, ` + "`" + `uniques_list` + "`" + `, ` + "`" + `line_chart` + "`" + `, ` + "`" + `comparison_line_chart` + "`" + `, ` + "`" + `markdown` + "`" + `, and ` + "`" + `metric_line_chart` + "`" + `.`,
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
					Description: `(Optional) Width of the widget. Valid values are ` + "`" + `1` + "`" + ` to ` + "`" + `3` + "`" + ` inclusive. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `(Optional) Height of the widget. Valid values are ` + "`" + `1` + "`" + ` to ` + "`" + `3` + "`" + ` inclusive. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Description of the widget. Each ` + "`" + `visualization` + "`" + ` type supports an additional set of arguments:`,
				},
				resource.Attribute{
					Name:        "nrql",
					Description: `(Required) Valid NRQL query string. See [Writing NRQL Queries](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/using-nrql/introduction-nrql) for help.`,
				},
				resource.Attribute{
					Name:        "threshold_red",
					Description: `(Optional) Threshold above which the displayed value will be styled with a red color.`,
				},
				resource.Attribute{
					Name:        "threshold_yellow",
					Description: `(Optional) Threshold above which the displayed value will be styled with a yellow color.`,
				},
				resource.Attribute{
					Name:        "nrql",
					Description: `(Required) Valid NRQL query string. See [Writing NRQL Queries](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/using-nrql/introduction-nrql) for help.`,
				},
				resource.Attribute{
					Name:        "threshold_red",
					Description: `(Required) Threshold above which the displayed value will be styled with a red color.`,
				},
				resource.Attribute{
					Name:        "threshold_yellow",
					Description: `(Optional) Threshold above which the displayed value will be styled with a yellow color.`,
				},
				resource.Attribute{
					Name:        "nrql",
					Description: `(Required) Valid NRQL query string. See [Writing NRQL Queries](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/using-nrql/introduction-nrql) for help.`,
				},
				resource.Attribute{
					Name:        "drilldown_dashboard_id",
					Description: `(Optional) The ID of a dashboard to link to from the widget's facets.`,
				},
				resource.Attribute{
					Name:        "nrql",
					Description: `(Required) Valid NRQL query string. See [Writing NRQL Queries](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/using-nrql/introduction-nrql) for help.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The markdown source to be rendered in the widget.`,
				},
				resource.Attribute{
					Name:        "entity_ids",
					Description: `(Required) A collection of entity ids to display data for. These are typically application IDs.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Required) A nested block that describes a metric. Nested ` + "`" + `metric` + "`" + ` blocks support the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The metric name to display.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) The metric values to display.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) The duration, in ms, of the time window represented in the chart.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) The end time of the time window represented in the chart in epoch time. When not set, the time window will end at the current time.`,
				},
				resource.Attribute{
					Name:        "facet",
					Description: `(Optional) Can be set to "host" to facet the metric data by host.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) The limit of distinct data series to display.`,
				},
				resource.Attribute{
					Name:        "entity_ids",
					Description: `(Required) A collection of entity IDs to display data. These are typically application IDs. ### Nested ` + "`" + `filter` + "`" + ` block The optional filter block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "event_types",
					Description: `(Optional) A list of event types to enable filtering for.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Optional) A list of attributes belonging to the specified event types to enable filtering for. ## Import New Relic dashboards can be imported using their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_dashboard.my_dashboard 8675309 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_infra_alert_condition",
			Category:         "Resources",
			ShortDescription: `Create and manage an Infrastructure alert condition for a policy in New Relic.`,
			Description: `\_infra_alert\_condition

Use this resource to create and manage Infrastructure alert conditions in New Relic.

`,
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
					Name:        "type",
					Description: `(Required) The type of Infrastructure alert condition. Valid values are ` + "`" + `infra_process_running` + "`" + `, ` + "`" + `infra_metric` + "`" + `, and ` + "`" + `infra_host_not_reporting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event",
					Description: `(Required) The metric event; for example, ` + "`" + `SystemSample` + "`" + ` or ` + "`" + `StorageSample` + "`" + `. Supported by the ` + "`" + `infra_metric` + "`" + ` condition type.`,
				},
				resource.Attribute{
					Name:        "select",
					Description: `(Required) The attribute name to identify the metric being targeted; for example, ` + "`" + `cpuPercent` + "`" + `, ` + "`" + `diskFreePercent` + "`" + `, or ` + "`" + `memoryResidentSizeBytes` + "`" + `. The underlying API will automatically populate this value for Infrastructure integrations (for example ` + "`" + `diskFreePercent` + "`" + `), so make sure to explicitly include this value to avoid diff issues. Supported by the ` + "`" + `infra_metric` + "`" + ` condition type.`,
				},
				resource.Attribute{
					Name:        "comparison",
					Description: `(Required) The operator used to evaluate the threshold value. Valid values are ` + "`" + `above` + "`" + `, ` + "`" + `below` + "`" + `, and ` + "`" + `equal` + "`" + `. Supported by the ` + "`" + `infra_metric` + "`" + ` and ` + "`" + `infra_process_running` + "`" + ` condition types.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Required) Identifies the threshold parameters for opening a critial alert violation. See [Thresholds](#thresholds) below for details.`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Optional) Identifies the threshold parameters for opening a warning alert violation. See [Thresholds](#thresholds) below for details.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the condition is turned on or off. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "where",
					Description: `(Optional) If applicable, this identifies any Infrastructure host filters used; for example: ` + "`" + `hostname LIKE '%cassandra%'` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "process_where",
					Description: `(Optional) Any filters applied to processes; for example: ` + "`" + `commandName = 'java'` + "`" + `. Supported by the ` + "`" + `infra_process_running` + "`" + ` condition type.`,
				},
				resource.Attribute{
					Name:        "integration_provider",
					Description: `(Optional) For alerts on integrations, use this instead of ` + "`" + `event` + "`" + `. Supported by the ` + "`" + `infra_metric` + "`" + ` condition type.`,
				},
				resource.Attribute{
					Name:        "runbook_url",
					Description: `(Optional) Runbook URL to display in notifications. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Infrastructure alert condition. ## Thresholds The ` + "`" + `critical` + "`" + ` and ` + "`" + `warning` + "`" + ` threshold mapping supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) Identifies the number of minutes the threshold must be passed or met for the alert to trigger. Threshold durations must be between 1 and 60 minutes (inclusive).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Threshold value, computed against the ` + "`" + `comparison` + "`" + ` operator. Supported by ` + "`" + `infra_metric` + "`" + ` and ` + "`" + `infra_process_running` + "`" + ` alert condition types.`,
				},
				resource.Attribute{
					Name:        "time_function",
					Description: `(Optional) Indicates if the condition needs to be sustained or to just break the threshold once; ` + "`" + `all` + "`" + ` or ` + "`" + `any` + "`" + `. Supported by the ` + "`" + `infra_metric` + "`" + ` alert condition type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Infrastructure alert condition. ## Thresholds The ` + "`" + `critical` + "`" + ` and ` + "`" + `warning` + "`" + ` threshold mapping supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) Identifies the number of minutes the threshold must be passed or met for the alert to trigger. Threshold durations must be between 1 and 60 minutes (inclusive).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Threshold value, computed against the ` + "`" + `comparison` + "`" + ` operator. Supported by ` + "`" + `infra_metric` + "`" + ` and ` + "`" + `infra_process_running` + "`" + ` alert condition types.`,
				},
				resource.Attribute{
					Name:        "time_function",
					Description: `(Optional) Indicates if the condition needs to be sustained or to just break the threshold once; ` + "`" + `all` + "`" + ` or ` + "`" + `any` + "`" + `. Supported by the ` + "`" + `infra_metric` + "`" + ` alert condition type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_insights_event",
			Category:         "Resources",
			ShortDescription: `Create one or more Insights events.`,
			Description: `\_insights\_event

Use this resource to create one or more Insights events during a terraform run.

`,
			Keywords: []string{
				"insights",
				"event",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "event",
					Description: `(Required) An event to insert into Insights. Multiple event blocks can be defined. See [Events](#events) below for details. ## Events The ` + "`" + `event` + "`" + ` mapping supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The event's name. Can be a combination of alphanumeric characters, underscores, and colons.`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `(Optional) Must be a Unix epoch timestamp. You can define timestamps either in seconds or in milliseconds.`,
				},
				resource.Attribute{
					Name:        "attribute",
					Description: `(Required) An attribute to include in your event payload. Multiple attribute blocks can be defined for an event. See [Attributes](#attributes) below for details. ### Attributes The ` + "`" + `attribute` + "`" + ` mapping supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The name of the attribute.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the attribute.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Specify the type for the attribute value. This is useful when passing integer or float values to Insights. Allowed values are ` + "`" + `string` + "`" + `, ` + "`" + `int` + "`" + `, or ` + "`" + `float` + "`" + `. Defaults to ` + "`" + `string` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_nrql_alert_condition",
			Category:         "Resources",
			ShortDescription: `Create and manage a NRQL alert condition for a policy in New Relic.`,
			Description: `

Use this resource to create and manage NRQL alert conditions in New Relic.

`,
			Keywords: []string{
				"nrql",
				"alert",
				"condition",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_plugins_alert_condition",
			Category:         "Resources",
			ShortDescription: `Create and manage a Plugins alert condition for a policy in New Relic.`,
			Description: `\_plugins\_alert\_condition

Use this resource to create and manage plugins alert conditions in New Relic.

`,
			Keywords: []string{
				"plugins",
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
					Name:        "metric",
					Description: `(Required) The metric field accepts parameters based on the ` + "`" + `type` + "`" + ` set.`,
				},
				resource.Attribute{
					Name:        "plugin_id",
					Description: `(Required) The ID of the installed plugin instance which produces the metric.`,
				},
				resource.Attribute{
					Name:        "plugin_guid",
					Description: `(Required) The GUID of the plugin which produces the metric.`,
				},
				resource.Attribute{
					Name:        "runbook_url",
					Description: `(Optional) Runbook URL to display in notifications.`,
				},
				resource.Attribute{
					Name:        "term",
					Description: `(Required) A list of terms for this condition. See [Terms](#terms) below for details. ## Terms The ` + "`" + `term` + "`" + ` mapping supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) In minutes, must be in the range of ` + "`" + `5` + "`" + ` to ` + "`" + `120` + "`" + `, inclusive.`,
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
					Description: `(Required) ` + "`" + `all` + "`" + ` or ` + "`" + `any` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert condition. ## Import Alert conditions can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_plugins_alert_condition.main 12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert condition. ## Import Alert conditions can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_plugins_alert_condition.main 12345 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_synthetics_alert_condition",
			Category:         "Resources",
			ShortDescription: `Create and manage a Synthetics alert condition for a policy in New Relic.`,
			Description: `\_synthetics\_alert\_condition

Use this resource to create and manage synthetics alert conditions in New Relic.

`,
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
					Description: `(Optional) Set whether to enable the alert condition. Defaults to ` + "`" + `true` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "newrelic_synthetics_label",
			Category:         "Resources",
			ShortDescription: `Create and manage a Synthetics label in New Relic.`,
			Description:      ``,
			Keywords: []string{
				"synthetics",
				"label",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "monitor_id",
					Description: `(Required) The ID of the monitor that will be assigned the label.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) A string representing the label key/category.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) A string representing the label value. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The URL of the Synthetics label. ## Import Synthetics labels can be imported using an ID in the format ` + "`" + `<monitor_id>:<type>:<value>` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_synthetics_labels.foo 1a272364-f204-4cd3-ae2a-2d15a2bedadd:MyCategory:MyValue ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "href",
					Description: `The URL of the Synthetics label. ## Import Synthetics labels can be imported using an ID in the format ` + "`" + `<monitor_id>:<type>:<value>` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_synthetics_labels.foo 1a272364-f204-4cd3-ae2a-2d15a2bedadd:MyCategory:MyValue ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) The monitor type. Valid values are ` + "`" + `SIMPLE` + "`" + `, ` + "`" + `BROWSER` + "`" + `, ` + "`" + `SCRIPT_BROWSER` + "`" + `, and ` + "`" + `SCRIPT_API` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `(Required) The interval (in minutes) at which this monitor should run.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Required) The monitor status (i.e. ` + "`" + `ENABLED` + "`" + `, ` + "`" + `MUTED` + "`" + `, ` + "`" + `DISABLED` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `(Required) The locations in which this monitor should be run.`,
				},
				resource.Attribute{
					Name:        "sla_threshold",
					Description: `(Optional) The base threshold for the SLA report. The ` + "`" + `SIMPLE` + "`" + ` monitor type supports the following additional arguments:`,
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
					Name:        "treat_redirect_as_failure",
					Description: `(Optional) Fail the monitor check if redirected. The ` + "`" + `BROWSER` + "`" + ` monitor type supports the following additional arguments:`,
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
					Description: `(Optional) Verify SSL. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Synthetics monitor. ## Additional Examples Type: ` + "`" + `BROWSER` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_synthetics_monitor" "foo" { name = "foo" type = "BROWSER" frequency = 5 status = "ENABLED" locations = ["AWS_US_EAST_1"] uri = "https://example.com" # required for type "SIMPLE" and "BROWSER" validation_string = "add example validation check here" # optional for type "SIMPLE" and "BROWSER" verify_ssl = true # optional for type "SIMPLE" and "BROWSER" bypass_head_request = true # Note: optional for type "BROWSER" only treat_redirect_as_failure = true # Note: optional for type "BROWSER" only } ` + "`" + `` + "`" + `` + "`" + ` Type: ` + "`" + `SCRIPT_BROWSER` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_synthetics_monitor" "foo" { name = "foo" type = "SCRIPT_BROWSER" frequency = 5 status = "ENABLED" locations = ["AWS_US_EAST_1"] } ` + "`" + `` + "`" + `` + "`" + ` Type: ` + "`" + `SCRIPT_API` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_synthetics_monitor" "foo" { name = "foo" type = "SCRIPT_API" frequency = 5 status = "ENABLED" locations = ["AWS_US_EAST_1"] } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Synthetics monitor. ## Additional Examples Type: ` + "`" + `BROWSER` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_synthetics_monitor" "foo" { name = "foo" type = "BROWSER" frequency = 5 status = "ENABLED" locations = ["AWS_US_EAST_1"] uri = "https://example.com" # required for type "SIMPLE" and "BROWSER" validation_string = "add example validation check here" # optional for type "SIMPLE" and "BROWSER" verify_ssl = true # optional for type "SIMPLE" and "BROWSER" bypass_head_request = true # Note: optional for type "BROWSER" only treat_redirect_as_failure = true # Note: optional for type "BROWSER" only } ` + "`" + `` + "`" + `` + "`" + ` Type: ` + "`" + `SCRIPT_BROWSER` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_synthetics_monitor" "foo" { name = "foo" type = "SCRIPT_BROWSER" frequency = 5 status = "ENABLED" locations = ["AWS_US_EAST_1"] } ` + "`" + `` + "`" + `` + "`" + ` Type: ` + "`" + `SCRIPT_API` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `hcl resource "newrelic_synthetics_monitor" "foo" { name = "foo" type = "SCRIPT_API" frequency = 5 status = "ENABLED" locations = ["AWS_US_EAST_1"] } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_synthetics_monitor_script",
			Category:         "Resources",
			ShortDescription: `Update and manage a Synthetics monitor script in New Relic.`,
			Description: `\_synthetics\_monitor\_script

Use this resource to update a synthetics monitor script in New Relic.

`,
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
					Description: `(Required) plaintext of the monitor script. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_synthetics_secure_credential",
			Category:         "Resources",
			ShortDescription: `Create and manage Synthetics secure credentials in New Relic.`,
			Description: `\_synthetics\_secure\_credential

Use this resource to create and manage New Relic Synthetic secure credentials.

`,
			Keywords: []string{
				"synthetics",
				"secure",
				"credential",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The secure credential's key name. Regardless of the case used in the configuration, the provider will provide an upcased key to the underlying API.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The secure credential's value.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The secure credential's description. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time the secure credential was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time the secure credential was last updated. ## Import A Synthetics secure credential can be imported using its ` + "`" + `key` + "`" + `: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_synthetics_secure_credential.foo MY_KEY ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `The time the secure credential was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time the secure credential was last updated. ## Import A Synthetics secure credential can be imported using its ` + "`" + `key` + "`" + `: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import newrelic_synthetics_secure_credential.foo MY_KEY ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"newrelic_alert_channel":                0,
		"newrelic_alert_condition":              1,
		"newrelic_alert_policy":                 2,
		"newrelic_alert_policy_channel":         3,
		"newrelic_application_label":            4,
		"newrelic_dashboard":                    5,
		"newrelic_infra_alert_condition":        6,
		"newrelic_insights_event":               7,
		"newrelic_nrql_alert_condition":         8,
		"newrelic_plugins_alert_condition":      9,
		"newrelic_synthetics_alert_condition":   10,
		"newrelic_synthetics_label":             11,
		"newrelic_synthetics_monitor":           12,
		"newrelic_synthetics_monitor_script":    13,
		"newrelic_synthetics_secure_credential": 14,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
