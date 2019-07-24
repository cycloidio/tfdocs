package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "datadog_downtime",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog downtime resource. This can be used to create and manage downtimes.`,
			Description:      ``,
			Keywords: []string{
				"downtime",
			},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Type:             "datadog_metric_metadata",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog metric metadata resource. This can be used to manage a metric's metadata.`,
			Description:      ``,
			Keywords: []string{
				"metric",
				"metadata",
			},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{},
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "metadata_json",
					Description: `(Optional) A JSON blob (preferrably created using [jsonencode](https://www.terraform.io/docs/configuration/functions/jsonencode.html)) representing mapping of query expressions to alias names. Note that the query expressions in ` + "`" + `metadata_json` + "`" + ` will be ignored if they're not present in the query. For example, this is how you define ` + "`" + `metadata_json` + "`" + ` with Terraform >= 0.12: ` + "`" + `` + "`" + `` + "`" + ` metadata_json = jsonencode({ "avg:redis.info.latency_ms{$host}": { "alias": "Redis latency" } }) ` + "`" + `` + "`" + `` + "`" + ` And here's how you define ` + "`" + `metadata_json` + "`" + ` with Terraform < 0.12: ` + "`" + `` + "`" + `` + "`" + ` variable "my_metadata" { default = { "avg:redis.info.latency_ms{$host}" = { "alias": "Redis latency" } } } resource "datadog_screenboard" "SomeScreenboard" { ... metadata_json = "${jsonencode(var.my_metadata)}" } ` + "`" + `` + "`" + `` + "`" + ` Note that this has to be a JSON blob because of [limitations](https://github.com/hashicorp/terraform/issues/6215) of Terraform's handling complex nested structures. This is also why the key is called ` + "`" + `metadata_json` + "`" + ` even though it sets ` + "`" + `metadata` + "`" + ` attribute on the API call. ### Nested ` + "`" + `widget` + "`" + ` ` + "`" + `tile_def` + "`" + ` ` + "`" + `request` + "`" + ` ` + "`" + `style` + "`" + ` block Only for widgets of type "timeseries", "query_value", "toplist", "process". The nested ` + "`" + `style` + "`" + ` blocks has the following structure: - ` + "`" + `palette` + "`" + ` - (Optional) Color of the line drawn. For widgets of type "timeseries", "query_value", "toplist", one of: "classic", "cool", "warm", "purple", "orange" or "gray". For widgets of type "process", one of: "dog_classic_area", "YlOrRd", "GnBu", "Reds", "Oranges", "Greens", "Blues", "Purples". - ` + "`" + `width` + "`" + ` - (Optional) Line width. Possible values: "thin", "normal", "thick". Default: "normal". - ` + "`" + `type` + "`" + ` - (Optional) Type of line drawn. Possible values: "dashed", "solid", "dotted". Default: "solid". ### Nested ` + "`" + `widget` + "`" + ` ` + "`" + `tile_def` + "`" + ` ` + "`" + `request` + "`" + ` ` + "`" + `conditional_format` + "`" + ` block The nested ` + "`" + `conditional_format` + "`" + ` blocks has the following structure: - ` + "`" + `palette` + "`" + ` - (Optional) Color scheme to be used if the condition is met. One of: "red_on_white", "white_on_red", "yellow_on_white", "white_on_yellow", "green_on_white", "white_on_green", "gray_on_white", "white_on_gray", "custom_text", "custom_bg", "custom_image". - ` + "`" + `comparator` + "`" + ` - (Required) Comparison operator. Example: ">", "<". - ` + "`" + `value` + "`" + ` - (Optional) Value that is the threshold for the conditional format. - ` + "`" + `color` + "`" + ` - (Optional) Custom color (e.g., #205081). - ` + "`" + `invert` + "`" + ` - (Optional) Boolean indicating whether to invert color scheme. ### Nested ` + "`" + `template_variable` + "`" + ` blocks Nested ` + "`" + `template_variable` + "`" + ` blocks have the following structure: - ` + "`" + `name` + "`" + ` - (Required) The variable name. Can be referenced as $name in ` + "`" + `graph` + "`" + ` ` + "`" + `request` + "`" + ` ` + "`" + `q` + "`" + ` query strings. - ` + "`" + `prefix` + "`" + ` - (Optional) The tag group. Default: no tag group. - ` + "`" + `default` + "`" + ` - (Optional) The default tag. Default: "\`,
				},
			},
			Attributes: []resource.Argument{},
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
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
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
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "title",
					Description: `(Required) The name of the dashboard.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description of the dashboard's content.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) The read-only status of the timeboard. Default is false.`,
				},
				resource.Attribute{
					Name:        "graph",
					Description: `(Required) Nested block describing a graph definition. The structure of this block is described below. Multiple graph blocks are allowed within a datadog_timeboard resource.`,
				},
				resource.Attribute{
					Name:        "template_variable",
					Description: `(Optional) Nested block describing a template variable. The structure of this block is described below. Multiple template_variable blocks are allowed within a datadog_timeboard resource. ### Nested ` + "`" + `graph` + "`" + ` blocks Nested ` + "`" + `graph` + "`" + ` blocks have the following structure:`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) The name of the graph.`,
				},
				resource.Attribute{
					Name:        "viz",
					Description: `(Required) The type of visualization to use for the graph. Valid choices are "change", "distribution", "heatmap", "hostmap", "query_value", timeseries", and "toplist".`,
				},
				resource.Attribute{
					Name:        "request",
					Description: `Nested block describing a graph definition request (a metric query to plot on the graph). The structure of this block is described below. Multiple request blocks are allowed within a graph block.`,
				},
				resource.Attribute{
					Name:        "events",
					Description: `(Optional) A list of event filter strings. Note that, while supported by the Datadog API, the Datadog UI does not (currently) support multiple event filters very well, so use at your own risk.`,
				},
				resource.Attribute{
					Name:        "autoscale",
					Description: `(Optional) Boolean that determines whether to autoscale graphs.`,
				},
				resource.Attribute{
					Name:        "precision",
					Description: `(Optional) Number of digits displayed, use ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "custom_unit",
					Description: `(Optional) Display a custom unit on the graph (such as 'hertz')`,
				},
				resource.Attribute{
					Name:        "text_align",
					Description: `(Optional) How to align text in the graph, can be one of 'left', 'center', or 'right'.`,
				},
				resource.Attribute{
					Name:        "style",
					Description: `(Optional) Nested block describing hostmaps. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) List of groups for hostmaps (shown as 'group by' in the UI).`,
				},
				resource.Attribute{
					Name:        "include_no_metric_hosts",
					Description: `(Optional) If set to true, will display hosts on hostmap that have no reported metrics.`,
				},
				resource.Attribute{
					Name:        "include_ungrouped_hosts",
					Description: `(Optional) If set to true, will display hosts without groups on hostmaps.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Optional) What nodes to display in a hostmap. Can be one of 'host' (default) or 'container'.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) List of scopes for hostmaps (shown as 'filter by' in the UI).`,
				},
				resource.Attribute{
					Name:        "yaxis",
					Description: `(Optional) Nested block describing modifications to the yaxis rendering. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "marker",
					Description: `(Optional) Nested block describing lines / ranges added to graph for formatting. The structure of this block is described below. Multiple marker blocks are allowed within a graph block. ### Nested ` + "`" + `graph` + "`" + ` ` + "`" + `marker` + "`" + ` blocks Nested ` + "`" + `graph` + "`" + ` ` + "`" + `marker` + "`" + ` blocks have the following structure:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) How the marker lines will look. Possible values are {"error", "warning", "info", "ok"} {"dashed", "solid", "bold"}. Example: "error dashed".`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Mathematical expression describing the marker. Examples: "y > 1", "-5 < y < 0", "y = 19".`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) A label for the line or range.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `(Optional) Minimum bound for the graph's yaxis, a string.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `(Optional) Maximum bound for the graph's yaxis, a string.`,
				},
				resource.Attribute{
					Name:        "scale",
					Description: `(Optional) How to scale the yaxis. Possible values are: "linear", "log", "sqrt", "pow##" (eg. pow2, pow0.5, 2 is used if only "pow" was provided). Default: "linear". ### Nested ` + "`" + `graph` + "`" + ` ` + "`" + `request` + "`" + ` blocks Nested ` + "`" + `graph` + "`" + ` ` + "`" + `request` + "`" + ` blocks have the following structure:`,
				},
				resource.Attribute{
					Name:        "q",
					Description: `(Required) The query of the request. Pro tip: Use the JSON tab inside the Datadog UI to help build you query strings.`,
				},
				resource.Attribute{
					Name:        "aggregator",
					Description: `(Optional) The aggregation method used when the number of data points outnumbers the max that can be shown.`,
				},
				resource.Attribute{
					Name:        "stacked",
					Description: `(Optional) Boolean value to determine if this is this a stacked area graph. Default: false (line chart).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Choose how to draw the graph. For example: "line", "bars" or "area". Default: "line".`,
				},
				resource.Attribute{
					Name:        "style",
					Description: `(Optional) Nested block to customize the graph style.`,
				},
				resource.Attribute{
					Name:        "conditional_format",
					Description: `(Optional) Nested block to customize the graph style if certain conditions are met. Currently only applies to ` + "`" + `Query Value` + "`" + ` and ` + "`" + `Top List` + "`" + ` type graphs.`,
				},
				resource.Attribute{
					Name:        "extra_col",
					Description: `(Optional, only for graphs of visualization "change") If set to "present", displays current value. Can be left empty otherwise.`,
				},
				resource.Attribute{
					Name:        "metadata_json",
					Description: `(Optional) A JSON blob (preferrably created using [jsonencode](https://www.terraform.io/docs/configuration/functions/jsonencode.html)) representing mapping of query expressions to alias names. Note that the query expressions in ` + "`" + `metadata_json` + "`" + ` will be ignored if they're not present in the query. For example, this is how you define ` + "`" + `metadata_json` + "`" + ` with Terraform >= 0.12: ` + "`" + `` + "`" + `` + "`" + ` metadata_json = jsonencode({ "avg:redis.info.latency_ms{$host}": { "alias": "Redis latency" } }) ` + "`" + `` + "`" + `` + "`" + ` And here's how you define ` + "`" + `metadata_json` + "`" + ` with Terraform < 0.12: ` + "`" + `` + "`" + `` + "`" + ` variable "my_metadata" { default = { "avg:redis.info.latency_ms{$host}" = { "alias": "Redis latency" } } } resource "datadog_timeboard" "SomeTimeboard" { ... metadata_json = "${jsonencode(var.my_metadata)}" } ` + "`" + `` + "`" + `` + "`" + ` Note that this has to be a JSON blob because of [limitations](https://github.com/hashicorp/terraform/issues/6215) of Terraform's handling complex nested structures. This is also why the key is called ` + "`" + `metadata_json` + "`" + ` even though it sets ` + "`" + `metadata` + "`" + ` attribute on the API call. ### Nested ` + "`" + `graph` + "`" + ` ` + "`" + `style` + "`" + ` block The nested ` + "`" + `style` + "`" + ` block is used specifically for styling ` + "`" + `hostmap` + "`" + ` graphs, and has the following structure:`,
				},
				resource.Attribute{
					Name:        "fill_max",
					Description: `(Optional) Maximum value for the hostmap fill query.`,
				},
				resource.Attribute{
					Name:        "fill_min",
					Description: `(Optional) Minimum value for the hostmap fill query.`,
				},
				resource.Attribute{
					Name:        "palette",
					Description: `(Optional) Spectrum of colors to use when styling a hostmap. For example: "green_to_orange", "yellow_to_green", "YlOrRd", or "hostmap_blues". Default: "green_to_orange".`,
				},
				resource.Attribute{
					Name:        "palette_flip",
					Description: `(Optional) Flip how the hostmap is rendered. For example, with the default palette, low values are represented as green, with high values as orange. If palette_flip is "true", then low values will be orange, and high values will be green. ### Nested ` + "`" + `graph` + "`" + ` ` + "`" + `request` + "`" + ` ` + "`" + `style` + "`" + ` block The nested ` + "`" + `style` + "`" + ` blocks has the following structure:`,
				},
				resource.Attribute{
					Name:        "palette",
					Description: `(Optional) Color of the line drawn. For example: "classic", "cool", "warm", "purple", "orange" or "gray". Default: "classic".`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `(Optional) Line width. Possible values: "thin", "normal", "thick". Default: "normal".`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of line drawn. Possible values: "dashed", "solid", "dotted". Default: "solid". ### Nested ` + "`" + `graph` + "`" + ` ` + "`" + `request` + "`" + ` ` + "`" + `conditional_format` + "`" + ` block The nested ` + "`" + `conditional_format` + "`" + ` blocks has the following structure:`,
				},
				resource.Attribute{
					Name:        "palette",
					Description: `(Optional) Color scheme to be used if the condition is met. For example: "red_on_white", "white_on_red", "yellow_on_white", "white_on_yellow", "green_on_white", "white_on_green", "gray_on_white", "white_on_gray", "custom_text", "custom_bg", "custom_image".`,
				},
				resource.Attribute{
					Name:        "comparator",
					Description: `(Required) Comparison operator. Example: ">", "<".`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value that is the threshold for the conditional format.`,
				},
				resource.Attribute{
					Name:        "custom_fg_color",
					Description: `(Optional) Used when ` + "`" + `palette` + "`" + ` is set to ` + "`" + `custom_text` + "`" + `. Set the color of the text to a custom web color, such as "#205081".`,
				},
				resource.Attribute{
					Name:        "custom_bg_color",
					Description: `(Optional) Used when ` + "`" + `palette` + "`" + ` is set to ` + "`" + `custom_bg` + "`" + `. Set the color of the background to a custom web color, such as "#205081". ### Nested ` + "`" + `template_variable` + "`" + ` blocks Nested ` + "`" + `template_variable` + "`" + ` blocks have the following structure:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The variable name. Can be referenced as $name in ` + "`" + `graph` + "`" + ` ` + "`" + `request` + "`" + ` ` + "`" + `q` + "`" + ` query strings.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) The tag group. Default: no tag group.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) The default tag. Default: "`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of this timeboard in your Datadog account. The web interface URL to this timeboard can be generated by appending this ID to ` + "`" + `` + "`" + `https://app.datadoghq.com/dash/` + "`" + `` + "`" + ` ## Import Timeboards can be imported using their numeric ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import datadog_timeboard.my_service_timeboard 2081 ` + "`" + `` + "`" + `` + "`" + ` ## Dynamic Timeboards Since Terraform 0.12, it's possible to create timeboard graphs dynamically based on contents of a list/map variable. This can be achieved by using the [dynamic blocks](https://www.terraform.io/docs/configuration/expressions.html#dynamic-blocks) feature. For example: ` + "`" + `` + "`" + `` + "`" + ` variable "my_list" { default = ["First", "Second", "Third"] } variable "my_map" { default = { "First" = "value1" "Second" = "value2" } } # Create a timeboard with "First", "Second" and "Third" timeseries graphs resource "datadog_timeboard" "my_timeboard" { title = "My Timeboard" description = "My Description" read_only = true dynamic "graph" { for_each = var.my_list content { title = "${graph.value}" viz = "timeseries" request { q = "anomalies(sum:mycount{adapter:${graph.value}}.as_count().rollup(sum, 3600), 'robust', 4, direction='below')" } } } } # Create a timeboard with "First" and "Second" timeseries graphs, use map keys as titles and map values as adapter names resource "datadog_timeboard" "my_timeboard_map" { title = "My Timeboard From Map" description = "My Description" read_only = true dynamic "graph" { for_each = var.my_map content { title = "${graph.key}" viz = "timeseries" request { q = "anomalies(sum:mycount{adapter:${graph.value}}.as_count().rollup(sum, 3600), 'robust', 4, direction='below')" } } } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of this timeboard in your Datadog account. The web interface URL to this timeboard can be generated by appending this ID to ` + "`" + `` + "`" + `https://app.datadoghq.com/dash/` + "`" + `` + "`" + ` ## Import Timeboards can be imported using their numeric ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import datadog_timeboard.my_service_timeboard 2081 ` + "`" + `` + "`" + `` + "`" + ` ## Dynamic Timeboards Since Terraform 0.12, it's possible to create timeboard graphs dynamically based on contents of a list/map variable. This can be achieved by using the [dynamic blocks](https://www.terraform.io/docs/configuration/expressions.html#dynamic-blocks) feature. For example: ` + "`" + `` + "`" + `` + "`" + ` variable "my_list" { default = ["First", "Second", "Third"] } variable "my_map" { default = { "First" = "value1" "Second" = "value2" } } # Create a timeboard with "First", "Second" and "Third" timeseries graphs resource "datadog_timeboard" "my_timeboard" { title = "My Timeboard" description = "My Description" read_only = true dynamic "graph" { for_each = var.my_list content { title = "${graph.value}" viz = "timeseries" request { q = "anomalies(sum:mycount{adapter:${graph.value}}.as_count().rollup(sum, 3600), 'robust', 4, direction='below')" } } } } # Create a timeboard with "First" and "Second" timeseries graphs, use map keys as titles and map values as adapter names resource "datadog_timeboard" "my_timeboard_map" { title = "My Timeboard From Map" description = "My Description" read_only = true dynamic "graph" { for_each = var.my_map content { title = "${graph.key}" viz = "timeseries" request { q = "anomalies(sum:mycount{adapter:${graph.value}}.as_count().rollup(sum, 3600), 'robust', 4, direction='below')" } } } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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

		"datadog_downtime":        0,
		"datadog_integration_aws": 1,
		"datadog_integration_gcp": 2,
		"datadog_metric_metadata": 3,
		"datadog_monitor":         4,
		"datadog_screenboard":     5,
		"datadog_synthetics":      6,
		"datadog_timeboard":       7,
		"datadog_user":            8,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
