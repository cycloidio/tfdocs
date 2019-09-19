package datadog

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "datadog_dashboard",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog dashboard resource. This can be used to create and manage dashboards.`,
			Description:      ``,
			Keywords: []string{
				"dashboard",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_dashboard_list",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog dashboard list resource. This can be used to create and manage dashboard lists and their sub elements.`,
			Description:      ``,
			Keywords: []string{
				"dashboard",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of this Dashbaord List.`,
				},
				resource.Attribute{
					Name:        "dash_item",
					Description: `(Optional) An individual dashboard object to add to this Dashboard List. If present, must contain the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of this dashboard. Available options are: ` + "`" + `custom_timeboard` + "`" + `, ` + "`" + `custom_screenboard` + "`" + `, ` + "`" + `integration_screenboard` + "`" + `, ` + "`" + `integration_timeboard` + "`" + `, and ` + "`" + `host_timeboard` + "`" + ``,
				},
				resource.Attribute{
					Name:        "dash_id",
					Description: `(Required) The ID of this dashboard. ## Import dashboard lists can be imported using their id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import datadog_dashboard_list.new_list 123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_downtime",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog downtime resource. This can be used to create and manage downtimes.`,
			Description:      ``,
			Keywords: []string{
				"downtime",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scope",
					Description: `(Required) A list of items to apply the downtime to, e.g. host:X`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) A flag indicating if the downtime is active now.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) A flag indicating if the downtime was disabled.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Optional) POSIX timestamp to start the downtime.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `(Optional) String representing date and time to start the downtime in RFC3339 format.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Optional) POSIX timestamp to end the downtime.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) String representing date and time to end the downtime in RFC3339 format.`,
				},
				resource.Attribute{
					Name:        "recurrence",
					Description: `(Optional) A dictionary to configure the downtime to be recurring.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `days, weeks, months, or years`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `How often to repeat as an integer. For example to repeat every 3 days, select a type of days and a period of 3.`,
				},
				resource.Attribute{
					Name:        "week_days",
					Description: `(Optional) A list of week days to repeat on. Choose from: Mon, Tue, Wed, Thu, Fri, Sat or Sun. Only applicable when type is weeks. First letter must be capitalized.`,
				},
				resource.Attribute{
					Name:        "until_occurrences",
					Description: `(Optional) How many times the downtime will be rescheduled. ` + "`" + `until_occurrences` + "`" + ` and ` + "`" + `until_date` + "`" + ` are mutually exclusive.`,
				},
				resource.Attribute{
					Name:        "until_date",
					Description: `(Optional) The date at which the recurrence should end as a POSIX timestamp. ` + "`" + `until_occurrences` + "`" + ` and ` + "`" + `until_date` + "`" + ` are mutually exclusive.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `(Optional) A message to include with notifications for this downtime.`,
				},
				resource.Attribute{
					Name:        "monitor_tags",
					Description: `(Optional) A list of monitor tags to match. The resulting downtime applies to monitors that match`,
				},
				resource.Attribute{
					Name:        "monitor_id",
					Description: `(Optional) Reference to which monitor this downtime is applied. When scheduling downtime for a given monitor, datadog changes ` + "`" + `silenced` + "`" + ` property of the monitor to match the ` + "`" + `end` + "`" + ` POSIX timestamp.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Datadog downtime. On updates this can sometime change based on API logic. For recurring downtimes it would be recommended to ` + "`" + `ignore_changes` + "`" + ` on this field.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `If true this indicates the downtime is currently active.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `If true this indicates the downtime is currently disabled. ## Import Downtimes can be imported using their numeric ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import datadog_downtime.bytes_received_localhost 2081 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Datadog downtime. On updates this can sometime change based on API logic. For recurring downtimes it would be recommended to ` + "`" + `ignore_changes` + "`" + ` on this field.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `If true this indicates the downtime is currently active.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `If true this indicates the downtime is currently disabled. ## Import Downtimes can be imported using their numeric ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import datadog_downtime.bytes_received_localhost 2081 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_aws",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog - Amazon Web Services integration resource. This can be used to create and manage the integrations.`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"aws",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) Your AWS Account ID without dashes.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) Your Datadog role delegation name.`,
				},
				resource.Attribute{
					Name:        "filter_tags",
					Description: `(Optional) Array of EC2 tags (in the form ` + "`" + `key:value` + "`" + `) defines a filter that Datadog use when collecting metrics from EC2. Wildcards, such as ` + "`" + `?` + "`" + ` (for single characters) and ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "host_tags",
					Description: `(Optional) Array of tags (in the form key:value) to add to all hosts and metrics reporting through this integration.`,
				},
				resource.Attribute{
					Name:        "account_specific_namespace_rules",
					Description: `(Optional) Enables or disables metric collection for specific AWS namespaces for this AWS account only. A list of namespaces can be found at the [available namespace rules API endpoint](https://api.datadoghq.com/api/v1/integration/aws/available_namespace_rules). ### See also`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `AWS External ID ## Import Amazon Web Services integrations can be imported using their ` + "`" + `account ID` + "`" + ` and ` + "`" + `role name` + "`" + ` separated with a colon (` + "`" + `:` + "`" + `), while the ` + "`" + `external_id` + "`" + ` should be passed by setting an environment variable called ` + "`" + `EXTERNAL_ID` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ EXTERNAL_ID=${external_id} terraform import datadog_integration_aws.test ${account_id}:${role_name} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "external_id",
					Description: `AWS External ID ## Import Amazon Web Services integrations can be imported using their ` + "`" + `account ID` + "`" + ` and ` + "`" + `role name` + "`" + ` separated with a colon (` + "`" + `:` + "`" + `), while the ` + "`" + `external_id` + "`" + ` should be passed by setting an environment variable called ` + "`" + `EXTERNAL_ID` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ EXTERNAL_ID=${external_id} terraform import datadog_integration_aws.test ${account_id}:${role_name} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_gcp",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog - Google Cloud Platform integration resource. This can be used to create and manage the integrations.`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"gcp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Your Google Cloud project ID found in your JSON service account key.`,
				},
				resource.Attribute{
					Name:        "private_key_id",
					Description: `(Required) Your private key ID found in your JSON service account key.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) Your private key name found in your JSON service account key.`,
				},
				resource.Attribute{
					Name:        "client_email",
					Description: `(Required) Your email found in your JSON service account key.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) Your ID found in your JSON service account key.`,
				},
				resource.Attribute{
					Name:        "host_filters",
					Description: `(Optional) Limit the GCE instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog. ### See also`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Google Cloud project ID`,
				},
				resource.Attribute{
					Name:        "client_email",
					Description: `Google Cloud project service account email`,
				},
				resource.Attribute{
					Name:        "host_filters",
					Description: `Host filters ## Import Google Cloud Platform integrations can be imported using their project ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import datadog_integration_gcp.awesome_gcp_project_integration project_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `Google Cloud project ID`,
				},
				resource.Attribute{
					Name:        "client_email",
					Description: `Google Cloud project service account email`,
				},
				resource.Attribute{
					Name:        "host_filters",
					Description: `Host filters ## Import Google Cloud Platform integrations can be imported using their project ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import datadog_integration_gcp.awesome_gcp_project_integration project_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_pagerduty",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog - PagerDuty integration resource. This can be used to create and manage the integration.`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"pagerduty",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "individual_services",
					Description: `(Optional) Boolean to specify whether or not individual service objects specified by [datadog_integration_pagerduty_service_object](/docs/providers/datadog/r/integration_pagerduty_service_object.html) resource are to be used. Mutually exclusive with ` + "`" + `services` + "`" + ` key.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional) Array of PagerDuty service objects.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) Your Service name in PagerDuty.`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `(Required) Your Service name associated service key in Pagerduty.`,
				},
				resource.Attribute{
					Name:        "schedules",
					Description: `(Optional) Array of your schedule URLs.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Required) Your PagerDuty account’s personalized subdomain name.`,
				},
				resource.Attribute{
					Name:        "api_token",
					Description: `(Optional) Your PagerDuty API token. ### See also`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_pagerduty_service_object",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog - PagerDuty integration resource. This can be used to create and manage the integration.`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"pagerduty",
				"service",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) Your Service name in PagerDuty.`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `(Required) Your Service name associated service key in PagerDuty. Note: Since the Datadog API never returns service keys, it is impossible to detect [drifts](https://www.hashicorp.com/blog/detecting-and-managing-drift-with-terraform). The best way to solve a drift is to manually mark the Service Object resource with [terraform taint](https://www.terraform.io/docs/commands/taint.html) to have it destroyed and recreated.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_metric_metadata",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog metric metadata resource. This can be used to manage a metric's metadata.`,
			Description:      ``,
			Keywords: []string{
				"metric",
				"metadata",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metric",
					Description: `(Required) The name of the metric.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the metric.`,
				},
				resource.Attribute{
					Name:        "short_name",
					Description: `(Optional) A short name of the metric.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Optional) Primary unit of the metric such as 'byte' or 'operation'.`,
				},
				resource.Attribute{
					Name:        "per_unit",
					Description: `(Optional) 'Per' unit of the metric such as 'second' in 'bytes per second'.`,
				},
				resource.Attribute{
					Name:        "statsd_interval",
					Description: `(Optional) If applicable, stasd flush interval in seconds for the metric.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_monitor",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog monitor resource. This can be used to create and manage monitors.`,
			Description:      ``,
			Keywords: []string{
				"monitor",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the monitor. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API [documentation](https://docs.datadoghq.com/api/?lang=python#create-a-monitor) page. Available options to choose from are:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Datadog monitor`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) The monitor query to notify on. Note this is not the same query you see in the UI and the syntax is different depending on the monitor ` + "`" + `type` + "`" + `, please see the [API Reference](https://docs.datadoghq.com/api/?lang=python#create-a-monitor) for details.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `(Required) A message to include with notifications for this monitor. Email notifications can be sent to specific users by using the same '@username' notation as events.`,
				},
				resource.Attribute{
					Name:        "escalation_message",
					Description: `(Optional) A message to include with a re-notification. Supports the '@username' notification allowed elsewhere.`,
				},
				resource.Attribute{
					Name:        "thresholds",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Datadog monitor ## Import Monitors can be imported using their numeric ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import datadog_monitor.bytes_received_localhost 2081 ` + "`" + `` + "`" + `` + "`" + ` ## Composite Monitors You can compose monitors of all types in order to define more specific alert conditions (see the [doc](https://docs.datadoghq.com/monitors/monitor_types/composite/)). You just need to reuse the ID of your ` + "`" + `datadog_monitor` + "`" + ` resources. You can also compose any monitor with a ` + "`" + `datadog_synthetics_test` + "`" + ` by passing the computed ` + "`" + `monitor_id` + "`" + ` attribute in the query. ` + "`" + `` + "`" + `` + "`" + `tf resource "datadog_monitor" "bar" { name = "Composite Monitor" type = "composite" message = "This is a message" query = "${datadog_monitor.foo.id} || ${datadog_synthetics_test.foo.monitor_id}" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Datadog monitor ## Import Monitors can be imported using their numeric ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import datadog_monitor.bytes_received_localhost 2081 ` + "`" + `` + "`" + `` + "`" + ` ## Composite Monitors You can compose monitors of all types in order to define more specific alert conditions (see the [doc](https://docs.datadoghq.com/monitors/monitor_types/composite/)). You just need to reuse the ID of your ` + "`" + `datadog_monitor` + "`" + ` resources. You can also compose any monitor with a ` + "`" + `datadog_synthetics_test` + "`" + ` by passing the computed ` + "`" + `monitor_id` + "`" + ` attribute in the query. ` + "`" + `` + "`" + `` + "`" + `tf resource "datadog_monitor" "bar" { name = "Composite Monitor" type = "composite" message = "This is a message" query = "${datadog_monitor.foo.id} || ${datadog_synthetics_test.foo.monitor_id}" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_screenboard",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog screenboard resource. This can be used to create and manage screenboards.`,
			Description:      ``,
			Keywords: []string{
				"screenboard",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata_json",
					Description: `(Optional) A JSON blob (preferrably created using [jsonencode](https://www.terraform.io/docs/configuration/functions/jsonencode.html)) representing mapping of query expressions to alias names. Note that the query expressions in ` + "`" + `metadata_json` + "`" + ` will be ignored if they're not present in the query. For example, this is how you define ` + "`" + `metadata_json` + "`" + ` with Terraform >= 0.12: ` + "`" + `` + "`" + `` + "`" + ` metadata_json = jsonencode({ "avg:redis.info.latency_ms{$host}": { "alias": "Redis latency" } }) ` + "`" + `` + "`" + `` + "`" + ` And here's how you define ` + "`" + `metadata_json` + "`" + ` with Terraform < 0.12: ` + "`" + `` + "`" + `` + "`" + ` variable "my_metadata" { default = { "avg:redis.info.latency_ms{$host}" = { "alias": "Redis latency" } } } resource "datadog_screenboard" "SomeScreenboard" { ... metadata_json = "${jsonencode(var.my_metadata)}" } ` + "`" + `` + "`" + `` + "`" + ` Note that this has to be a JSON blob because of [limitations](https://github.com/hashicorp/terraform/issues/6215) of Terraform's handling complex nested structures. This is also why the key is called ` + "`" + `metadata_json` + "`" + ` even though it sets ` + "`" + `metadata` + "`" + ` attribute on the API call. ### Nested ` + "`" + `widget` + "`" + ` ` + "`" + `tile_def` + "`" + ` ` + "`" + `request` + "`" + ` ` + "`" + `style` + "`" + ` block Only for widgets of type "timeseries", "query_value", "toplist", "process". The nested ` + "`" + `style` + "`" + ` blocks has the following structure: - ` + "`" + `palette` + "`" + ` - (Optional) Color of the line drawn. For widgets of type "timeseries", "query_value", "toplist", one of: "classic", "cool", "warm", "purple", "orange" or "gray". For widgets of type "process", one of: "dog_classic_area", "YlOrRd", "GnBu", "Reds", "Oranges", "Greens", "Blues", "Purples". - ` + "`" + `width` + "`" + ` - (Optional) Line width. Possible values: "thin", "normal", "thick". Default: "normal". - ` + "`" + `type` + "`" + ` - (Optional) Type of line drawn. Possible values: "dashed", "solid", "dotted". Default: "solid". ### Nested ` + "`" + `widget` + "`" + ` ` + "`" + `tile_def` + "`" + ` ` + "`" + `request` + "`" + ` ` + "`" + `apm_query` + "`" + ` and ` + "`" + `log_query` + "`" + ` blocks Nested ` + "`" + `apm_query` + "`" + ` and ` + "`" + `log_query` + "`" + ` blocks have the following structure (Visit the [ Graph Primer](https://docs.datadoghq.com/graphing/) for more information about these values): - ` + "`" + `index` + "`" + ` - (Required) - ` + "`" + `compute` + "`" + ` - (Required). Exactly one nested block is required with the following structure: - ` + "`" + `aggregation` + "`" + ` - (Required) - ` + "`" + `facet` + "`" + ` - (Optional) - ` + "`" + `interval` + "`" + ` - (Optional) - ` + "`" + `search` + "`" + ` - (Optional). One nested block is allowed with the following structure: - ` + "`" + `query` + "`" + ` - (Optional) - ` + "`" + `group_by` + "`" + ` - (Optional). Multiple nested blocks are allowed with the following structure: - ` + "`" + `facet` + "`" + ` - (Optional) - ` + "`" + `limit` + "`" + ` - (Optional) - ` + "`" + `sort` + "`" + ` - (Optional). One nested block is allowed with the following structure: - ` + "`" + `aggregation` + "`" + ` - (Optional) - ` + "`" + `order` + "`" + ` - (Optional) - ` + "`" + `facet` + "`" + ` - (Optional) ### Nested ` + "`" + `widget` + "`" + ` ` + "`" + `tile_def` + "`" + ` ` + "`" + `request` + "`" + ` ` + "`" + `process_query` + "`" + ` blocks Nested ` + "`" + `process_query` + "`" + ` blocks have the following structure (Visit the [ Graph Primer](https://docs.datadoghq.com/graphing/) for more information about these values): - ` + "`" + `metric` + "`" + ` - (Required) - ` + "`" + `search_by` + "`" + ` - (Required) - ` + "`" + `filter_by` + "`" + ` - (Required) - ` + "`" + `limit` + "`" + ` - (Required) ### Nested ` + "`" + `widget` + "`" + ` ` + "`" + `tile_def` + "`" + ` ` + "`" + `request` + "`" + ` ` + "`" + `conditional_format` + "`" + ` block The nested ` + "`" + `conditional_format` + "`" + ` blocks has the following structure: - ` + "`" + `palette` + "`" + ` - (Optional) Color scheme to be used if the condition is met. One of: "red_on_white", "white_on_red", "yellow_on_white", "white_on_yellow", "green_on_white", "white_on_green", "gray_on_white", "white_on_gray", "custom_text", "custom_bg", "custom_image". - ` + "`" + `comparator` + "`" + ` - (Required) Comparison operator. Example: ">", "<". - ` + "`" + `value` + "`" + ` - (Optional) Value that is the threshold for the conditional format. - ` + "`" + `color` + "`" + ` - (Optional) Custom color (e.g., #205081). - ` + "`" + `invert` + "`" + ` - (Optional) Boolean indicating whether to invert color scheme. ### Nested ` + "`" + `template_variable` + "`" + ` blocks Nested ` + "`" + `template_variable` + "`" + ` blocks have the following structure: - ` + "`" + `name` + "`" + ` - (Required) The variable name. Can be referenced as \$name in ` + "`" + `graph` + "`" + ` ` + "`" + `request` + "`" + ` ` + "`" + `q` + "`" + ` query strings. - ` + "`" + `prefix` + "`" + ` - (Optional) The tag group. Default: no tag group. - ` + "`" + `default` + "`" + ` - (Optional) The default tag. Default: "\`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_service_level_objective",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog service level objective resource. This can be used to create and manage service level objectives.`,
			Description:      ``,
			Keywords: []string{
				"service",
				"level",
				"objective",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the service level objective. The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API [documentation](https://docs.datadoghq.com/api/?lang=python#create-a-service-level-objective) page. Available options to choose from are:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Datadog service level objective`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of this service level objective.`,
				},
				resource.Attribute{
					Name:        "thresholds",
					Description: `(Required) - A list of thresholds and targets that define the service level objectives from the provided SLIs.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) the objective's target ` + "`" + `[0,100]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "target_display",
					Description: `(Optional) the string version to specify additional digits in the case of ` + "`" + `99` + "`" + ` but want 3 digits like ` + "`" + `99.000` + "`" + ` to display.`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Optional) the objective's warning value ` + "`" + `[0,100]` + "`" + `. This must be ` + "`" + `> target` + "`" + ` value.`,
				},
				resource.Attribute{
					Name:        "warning_display",
					Description: `(Optional) the string version to specify additional digits in the case of ` + "`" + `99` + "`" + ` but want 3 digits like ` + "`" + `99.000` + "`" + ` to display. The following options are specific to the ` + "`" + `type` + "`" + ` of service level objective:`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) The metric query configuration to use for the SLI. This is a dictionary and requires both the ` + "`" + `numerator` + "`" + ` and ` + "`" + `denominator` + "`" + ` fields which should be ` + "`" + `count` + "`" + ` metrics using the ` + "`" + `sum` + "`" + ` aggregator.`,
				},
				resource.Attribute{
					Name:        "numerator",
					Description: `(Required) the sum of all the ` + "`" + `good` + "`" + ` events`,
				},
				resource.Attribute{
					Name:        "denominator",
					Description: `(Required) the sum of the ` + "`" + `total` + "`" + ` events`,
				},
				resource.Attribute{
					Name:        "monitor_ids",
					Description: `(Optional) A list of numeric monitor IDs for which to use as SLIs. Their tags will be auto-imported into ` + "`" + `monitor_tags` + "`" + ` field in the API resource. At least 1 of ` + "`" + `monitor_ids` + "`" + ` or ` + "`" + `monitor_search` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "monitor_search",
					Description: `(Optional) The monitor query search used on the monitor search API to add monitor_ids by searching. Their tags will be auto-imported into ` + "`" + `monitor_tags` + "`" + ` field in the API resource. At least 1 of ` + "`" + `monitor_ids` + "`" + ` or ` + "`" + `monitor_search` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) A custom set of groups from the monitor(s) for which to use as the SLI instead of all the groups. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Datadog service level objective ## Import Service Level Objectives can be imported using their string ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import datadog_service_level_objective.12345678901234567890123456789012 "baz" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Datadog service level objective ## Import Service Level Objectives can be imported using their string ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import datadog_service_level_objective.12345678901234567890123456789012 "baz" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_synthetics",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog synthetics resource. This can be used to create and manage synthetics.`,
			Description:      ``,
			Keywords: []string{
				"synthetics",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_timeboard",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog timeboard resource. This can be used to create and manage timeboards.`,
			Description:      ``,
			Keywords: []string{
				"timeboard",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_user",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog user resource. This can be used to create and manage users.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_role",
					Description: `(Optional) Role description for user. Can be ` + "`" + `st` + "`" + ` (standard user), ` + "`" + `adm` + "`" + ` (admin user) or ` + "`" + `ro` + "`" + ` (read-only user). Default is ` + "`" + `st` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Whether the user is disabled`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email address for user`,
				},
				resource.Attribute{
					Name:        "handle",
					Description: `(Required) The user handle, must be a valid email.`,
				},
				resource.Attribute{
					Name:        "is_admin",
					Description: `(Deprecated) (Optional) Whether the user is an administrator.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for user`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Deprecated) Role description for user.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Returns true if Datadog user is disabled (NOTE: Datadog does not actually delete users so this will be true for those as well)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Datadog user`,
				},
				resource.Attribute{
					Name:        "verified",
					Description: `Returns true if Datadog user is verified ## Import users can be imported using their handle, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import datadog_user.example_user existing@example.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "disabled",
					Description: `Returns true if Datadog user is disabled (NOTE: Datadog does not actually delete users so this will be true for those as well)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Datadog user`,
				},
				resource.Attribute{
					Name:        "verified",
					Description: `Returns true if Datadog user is verified ## Import users can be imported using their handle, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import datadog_user.example_user existing@example.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"datadog_dashboard":                            0,
		"datadog_dashboard_list":                       1,
		"datadog_downtime":                             2,
		"datadog_integration_aws":                      3,
		"datadog_integration_gcp":                      4,
		"datadog_integration_pagerduty":                5,
		"datadog_integration_pagerduty_service_object": 6,
		"datadog_metric_metadata":                      7,
		"datadog_monitor":                              8,
		"datadog_screenboard":                          9,
		"datadog_service_level_objective":              10,
		"datadog_synthetics":                           11,
		"datadog_timeboard":                            12,
		"datadog_user":                                 13,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
