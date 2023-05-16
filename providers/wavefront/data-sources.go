package wavefront

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "wavefront_alert",
			Category:         "Data Sources",
			ShortDescription: `Get the information about a specific Wavefront alert.`,
			Description: `

Use this data source to get information about a Wavefront alert by its ID.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID associated with the alert data to be fetched. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get the information about the alert. data "wavefront_alert" "example" { id = "alert-id" } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the alert as it is displayed in Wavefront.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert in Wavefront.`,
				},
				resource.Attribute{
					Name:        "additional_information",
					Description: `User-supplied additional explanatory information about this alert.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `An email address or integration endpoint (such as PagerDuty or webhook) to notify when the alert status changes.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `A comma-separated list of the email addresses or integration endpoints (such as PagerDuty or webhook) to notify when the alert status changes. Multiple target types can be in the list.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `A Wavefront query that is evaluated at regular intervals (default is 1 minute). The alert fires and notifications are triggered when a data series matching this query evaluates to a non-zero value for a set number of consecutive minutes.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `A map of severity to condition for which this alert will trigger.`,
				},
				resource.Attribute{
					Name:        "display_expression",
					Description: `A second query the results of which are displayed in the alert user interface instead of the condition query.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `The number of consecutive minutes that a series matching the condition query must evaluate to "true" (non-zero value) before the alert fires.`,
				},
				resource.Attribute{
					Name:        "resolve_after_minutes",
					Description: `The number of consecutive minutes that a firing series matching the condition query must evaluate to "false" (zero value) before the alert resolves.`,
				},
				resource.Attribute{
					Name:        "notification_resend_frequency_minutes",
					Description: `How often to re-trigger a continually failing alert.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `The severity of the alert.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the alert.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags assigned to the alert.`,
				},
				resource.Attribute{
					Name:        "can_view",
					Description: `A list of users or groups that can view the alert.`,
				},
				resource.Attribute{
					Name:        "can_modify",
					Description: `A list of users or groups that can modify the alert.`,
				},
				resource.Attribute{
					Name:        "process_rate_minutes",
					Description: `The specified query is executed every ` + "`" + `process_rate_minutes` + "`" + ` minutes.`,
				},
				resource.Attribute{
					Name:        "evaluate_realtime_data",
					Description: `A Boolean flag to enable real-time evaluation.`,
				},
				resource.Attribute{
					Name:        "include_obsolete_metrics",
					Description: `A Boolean flag indicating whether to include obsolete metrics or not.`,
				},
				resource.Attribute{
					Name:        "failing_host_label_pairs",
					Description: `A list of failing host label pairs.`,
				},
				resource.Attribute{
					Name:        "in_maintenance_host_label_pairs",
					Description: `A list of in maintenance host label pairs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the alert as it is displayed in Wavefront.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert in Wavefront.`,
				},
				resource.Attribute{
					Name:        "additional_information",
					Description: `User-supplied additional explanatory information about this alert.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `An email address or integration endpoint (such as PagerDuty or webhook) to notify when the alert status changes.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `A comma-separated list of the email addresses or integration endpoints (such as PagerDuty or webhook) to notify when the alert status changes. Multiple target types can be in the list.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `A Wavefront query that is evaluated at regular intervals (default is 1 minute). The alert fires and notifications are triggered when a data series matching this query evaluates to a non-zero value for a set number of consecutive minutes.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `A map of severity to condition for which this alert will trigger.`,
				},
				resource.Attribute{
					Name:        "display_expression",
					Description: `A second query the results of which are displayed in the alert user interface instead of the condition query.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `The number of consecutive minutes that a series matching the condition query must evaluate to "true" (non-zero value) before the alert fires.`,
				},
				resource.Attribute{
					Name:        "resolve_after_minutes",
					Description: `The number of consecutive minutes that a firing series matching the condition query must evaluate to "false" (zero value) before the alert resolves.`,
				},
				resource.Attribute{
					Name:        "notification_resend_frequency_minutes",
					Description: `How often to re-trigger a continually failing alert.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `The severity of the alert.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the alert.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags assigned to the alert.`,
				},
				resource.Attribute{
					Name:        "can_view",
					Description: `A list of users or groups that can view the alert.`,
				},
				resource.Attribute{
					Name:        "can_modify",
					Description: `A list of users or groups that can modify the alert.`,
				},
				resource.Attribute{
					Name:        "process_rate_minutes",
					Description: `The specified query is executed every ` + "`" + `process_rate_minutes` + "`" + ` minutes.`,
				},
				resource.Attribute{
					Name:        "evaluate_realtime_data",
					Description: `A Boolean flag to enable real-time evaluation.`,
				},
				resource.Attribute{
					Name:        "include_obsolete_metrics",
					Description: `A Boolean flag indicating whether to include obsolete metrics or not.`,
				},
				resource.Attribute{
					Name:        "failing_host_label_pairs",
					Description: `A list of failing host label pairs.`,
				},
				resource.Attribute{
					Name:        "in_maintenance_host_label_pairs",
					Description: `A list of in maintenance host label pairs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_alerts",
			Category:         "Data Sources",
			ShortDescription: `Get the information about all Wavefront alerts.`,
			Description: `

Use this data source to get information about all Wavefront alerts.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Limit is the maximum number of results to be returned. Defaults to 100.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Offset is the offset from the first result to be returned. Defaults to 0. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get the information about all alerts. data "wavefront_alerts" "example" { limit = 10 offset = 0 } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "alerts",
					Description: `List of all alerts in Wavefront. For each alert you will see a list of attributes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the alert as it is displayed in Wavefront.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert in Wavefront.`,
				},
				resource.Attribute{
					Name:        "additional_information",
					Description: `User-supplied additional explanatory information about this alert.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `An email address or integration endpoint (such as PagerDuty or webhook) to notify when the alert status changes.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `A comma-separated list of the email addresses or integration endpoints (such as PagerDuty or webhook) to notify when the alert status changes. Multiple target types can be in the list.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `A Wavefront query that is evaluated at regular intervals (default is 1 minute). The alert fires and notifications are triggered when a data series matching this query evaluates to a non-zero value for a set number of consecutive minutes.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `A map of severity to condition for which this alert will trigger.`,
				},
				resource.Attribute{
					Name:        "display_expression",
					Description: `A second query the results of which are displayed in the alert user interface instead of the condition query.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `The number of consecutive minutes that a series matching the condition query must evaluate to "true" (non-zero value) before the alert fires.`,
				},
				resource.Attribute{
					Name:        "resolve_after_minutes",
					Description: `The number of consecutive minutes that a firing series matching the condition query must evaluate to "false" (zero value) before the alert resolves.`,
				},
				resource.Attribute{
					Name:        "notification_resend_frequency_minutes",
					Description: `How often to re-trigger a continually failing alert.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `The severity of the alert.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the alert.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags assigned to the alert.`,
				},
				resource.Attribute{
					Name:        "can_view",
					Description: `A list of users or groups that can view the alert.`,
				},
				resource.Attribute{
					Name:        "can_modify",
					Description: `A list of users or groups that can modify the alert.`,
				},
				resource.Attribute{
					Name:        "process_rate_minutes",
					Description: `The specified query is executed every ` + "`" + `process_rate_minutes` + "`" + ` minutes.`,
				},
				resource.Attribute{
					Name:        "evaluate_realtime_data",
					Description: `A Boolean flag to enable real-time evaluation.`,
				},
				resource.Attribute{
					Name:        "include_obsolete_metrics",
					Description: `A Boolean flag indicating whether to include obsolete metrics or not.`,
				},
				resource.Attribute{
					Name:        "failing_host_label_pairs",
					Description: `A list of failing host label pairs.`,
				},
				resource.Attribute{
					Name:        "in_maintenance_host_label_pairs",
					Description: `A list of in maintenance host label pairs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "alerts",
					Description: `List of all alerts in Wavefront. For each alert you will see a list of attributes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the alert as it is displayed in Wavefront.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert in Wavefront.`,
				},
				resource.Attribute{
					Name:        "additional_information",
					Description: `User-supplied additional explanatory information about this alert.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `An email address or integration endpoint (such as PagerDuty or webhook) to notify when the alert status changes.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `A comma-separated list of the email addresses or integration endpoints (such as PagerDuty or webhook) to notify when the alert status changes. Multiple target types can be in the list.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `A Wavefront query that is evaluated at regular intervals (default is 1 minute). The alert fires and notifications are triggered when a data series matching this query evaluates to a non-zero value for a set number of consecutive minutes.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `A map of severity to condition for which this alert will trigger.`,
				},
				resource.Attribute{
					Name:        "display_expression",
					Description: `A second query the results of which are displayed in the alert user interface instead of the condition query.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `The number of consecutive minutes that a series matching the condition query must evaluate to "true" (non-zero value) before the alert fires.`,
				},
				resource.Attribute{
					Name:        "resolve_after_minutes",
					Description: `The number of consecutive minutes that a firing series matching the condition query must evaluate to "false" (zero value) before the alert resolves.`,
				},
				resource.Attribute{
					Name:        "notification_resend_frequency_minutes",
					Description: `How often to re-trigger a continually failing alert.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `The severity of the alert.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the alert.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags assigned to the alert.`,
				},
				resource.Attribute{
					Name:        "can_view",
					Description: `A list of users or groups that can view the alert.`,
				},
				resource.Attribute{
					Name:        "can_modify",
					Description: `A list of users or groups that can modify the alert.`,
				},
				resource.Attribute{
					Name:        "process_rate_minutes",
					Description: `The specified query is executed every ` + "`" + `process_rate_minutes` + "`" + ` minutes.`,
				},
				resource.Attribute{
					Name:        "evaluate_realtime_data",
					Description: `A Boolean flag to enable real-time evaluation.`,
				},
				resource.Attribute{
					Name:        "include_obsolete_metrics",
					Description: `A Boolean flag indicating whether to include obsolete metrics or not.`,
				},
				resource.Attribute{
					Name:        "failing_host_label_pairs",
					Description: `A list of failing host label pairs.`,
				},
				resource.Attribute{
					Name:        "in_maintenance_host_label_pairs",
					Description: `A list of in maintenance host label pairs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_dashboard",
			Category:         "Data Sources",
			ShortDescription: `Get the information about a specific Wavefront dashboard.`,
			Description: `

Use this data source to get information about a certain Wavefront dashboard by its ID.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID associated with the dashboard data to be fetched. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get the information about a dashboard. data "wavefront_dashboard" "example" { id = "dashboard-id" } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the dashboard.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags to assign to this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description of the dashboard.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Unique identifier, also a URL slug of the dashboard.`,
				},
				resource.Attribute{
					Name:        "section",
					Description: `Dashboard chart sections. See [dashboard sections](#dashboard-sections).`,
				},
				resource.Attribute{
					Name:        "display_query_parameters",
					Description: `Whether the dashboard parameters section is opened by default when the dashboard is shown.`,
				},
				resource.Attribute{
					Name:        "parameter_details",
					Description: `The current JSON representation of dashboard parameters. See [parameter details](#parameter-details).`,
				},
				resource.Attribute{
					Name:        "display_section_table_of_contents",
					Description: `Whether the "pills" quick-linked the sections of the dashboard are displayed by default when the dashboard is shown.`,
				},
				resource.Attribute{
					Name:        "can_modify",
					Description: `A list of users that have modify ACL access to the dashboard.`,
				},
				resource.Attribute{
					Name:        "can_view",
					Description: `A list of users that have view ACL access to the dashboard.`,
				},
				resource.Attribute{
					Name:        "event_filter_type",
					Description: `How charts belonging to this dashboard should display events. ` + "`" + `BYCHART` + "`" + ` is default if unspecified. Valid options are: ` + "`" + `BYCHART` + "`" + `, ` + "`" + `AUTOMATIC` + "`" + `, ` + "`" + `ALL` + "`" + `, ` + "`" + `NONE` + "`" + `, ` + "`" + `BYDASHBOARD` + "`" + `, and ` + "`" + `BYCHARTANDDASHBOARD` + "`" + `. ### Dashboard Sections The ` + "`" + `section` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of this section.`,
				},
				resource.Attribute{
					Name:        "row",
					Description: `See [dashboard section rows](#dashboard-section-rows). ### Dashboard Section Rows The ` + "`" + `row` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "chart",
					Description: `Charts in this section. See [dashboard chart](#dashboard-chart). ### Dashboard Chart The ` + "`" + `chart` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Query expression to plot on the chart. See [chart source queries](#chart-source-queries).`,
				},
				resource.Attribute{
					Name:        "chart_setting",
					Description: `Chart settings. See [chart settings](#chart-settings).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the source.`,
				},
				resource.Attribute{
					Name:        "units",
					Description: `String to label the units of the chart on the Y-Axis.`,
				},
				resource.Attribute{
					Name:        "summarization",
					Description: `Summarization strategy for the chart. MEAN is default.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the chart.`,
				},
				resource.Attribute{
					Name:        "base",
					Description: `The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale. ### Chart Source Queries The ` + "`" + `source` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the source.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `Query expression to plot on the chart.`,
				},
				resource.Attribute{
					Name:        "source_description",
					Description: `A description for the purpose of this source.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Whether the source is disabled.`,
				},
				resource.Attribute{
					Name:        "scatter_plot_source",
					Description: `For scatter plots, does this query source the X-axis or the Y-axis, ` + "`" + `X` + "`" + `, or ` + "`" + `Y` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "query_builder_enabled",
					Description: `Whether or not this source line should have the query builder enabled. ### Chart Settings The ` + "`" + `chart_setting` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Chart Type. ` + "`" + `line` + "`" + ` refers to the Line Plot, ` + "`" + `scatter` + "`" + ` to the Point Plot, ` + "`" + `stacked-area` + "`" + ` to the Stacked Area plot, ` + "`" + `table` + "`" + ` to the Tabular View, ` + "`" + `scatterplot-xy` + "`" + ` to Scatter Plot, ` + "`" + `markdown-widget` + "`" + ` to the Markdown display, and ` + "`" + `sparkline` + "`" + ` to the Single Stat view. Valid options are ` + "`" + `line` + "`" + `, ` + "`" + `scatterplot` + "`" + `, ` + "`" + `stacked-area` + "`" + `, ` + "`" + `stacked-column` + "`" + `, ` + "`" + `table` + "`" + `, ` + "`" + `scatterplot-xy` + "`" + `, ` + "`" + `markdown-widget` + "`" + `, ` + "`" + `sparkline` + "`" + `, ` + "`" + `globe` + "`" + `, ` + "`" + `nodemap` + "`" + `, ` + "`" + `top-k` + "`" + `, ` + "`" + `status-list` + "`" + `, and ` + "`" + `histogram` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `Max value of the Y-axis. Set to null or leave blank for auto.`,
				},
				resource.Attribute{
					Name:        "line_type",
					Description: `Plot interpolation type. ` + "`" + `linear` + "`" + ` is default. Valid options are ` + "`" + `linear` + "`" + `, ` + "`" + `step-before` + "`" + `, ` + "`" + `step-after` + "`" + `, ` + "`" + `basis` + "`" + `, ` + "`" + `cardinal` + "`" + `, and ` + "`" + `monotone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "stack_type",
					Description: `Type of stacked chart (applicable only if chart type is ` + "`" + `stacked` + "`" + `). ` + "`" + `zero` + "`" + ` (default) means stacked from y=0. ` + "`" + `expand` + "`" + ` means normalized from 0 to 1. ` + "`" + `wiggle` + "`" + ` means minimize weighted changes. ` + "`" + `silhouette` + "`" + ` means to center the stream. Valid options are ` + "`" + `zero` + "`" + `, ` + "`" + `expand` + "`" + `, ` + "`" + `wiggle` + "`" + `, ` + "`" + `silhouette` + "`" + `, and ` + "`" + `bars` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "windowing",
					Description: `For the tabular view, whether to use the full time window for the query or the last X minutes. Valid options are ` + "`" + `full` + "`" + ` or ` + "`" + `last` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "window_size",
					Description: `Width, in minutes, of the time window to use for ` + "`" + `last` + "`" + ` windowing.`,
				},
				resource.Attribute{
					Name:        "show_hosts",
					Description: `For the tabular view, whether to display sources. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_labels",
					Description: `For the tabular view, whether to display labels. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_raw_values",
					Description: `For the tabular view, whether to display raw values. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_column_tags",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "column_tags",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "tag_mode",
					Description: `For the tabular view, which mode to use to determine which point tags to display. Valid options are ` + "`" + `all` + "`" + `, ` + "`" + `top` + "`" + `, or ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "num_tags",
					Description: `For the tabular view defines how many point tags to display.`,
				},
				resource.Attribute{
					Name:        "custom_tags",
					Description: `For the tabular view, a list of point tags to display when using the ` + "`" + `custom` + "`" + ` tag display mode.`,
				},
				resource.Attribute{
					Name:        "group_by_source",
					Description: `For the tabular view, whether to group multi metrics into a single row by a common source. If set to ` + "`" + `false` + "`" + `, each source is displayed in its own row. If set to ` + "`" + `true` + "`" + `, multiple metrics for the same host are displayed as different columns in the same row.`,
				},
				resource.Attribute{
					Name:        "sort_values_descending",
					Description: `For the tabular view, whether to display values in descending order. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "y1max",
					Description: `For plots with multiple Y-axes, max value for the right side Y-axis. Set null for auto.`,
				},
				resource.Attribute{
					Name:        "y1min",
					Description: `For plots with multiple Y-axes, min value for the right side Y-axis. Set null for auto.`,
				},
				resource.Attribute{
					Name:        "y1_units",
					Description: `For plots with multiple Y-axes, units for right side Y-axis.`,
				},
				resource.Attribute{
					Name:        "y0_scale_si_by_1024",
					Description: `(Optional) Whether to scale numerical magnitude labels for left Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).`,
				},
				resource.Attribute{
					Name:        "y1_scale_si_by_1024",
					Description: `(Optional) Whether to scale numerical magnitude labels for right Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).`,
				},
				resource.Attribute{
					Name:        "y0_unit_autoscaling",
					Description: `(Optional) Whether to automatically adjust magnitude labels and units for the left Y-axis to favor smaller magnitudes and larger units.`,
				},
				resource.Attribute{
					Name:        "y1_unit_autoscaling",
					Description: `(Optional) Whether to automatically adjust magnitude labels and units for the right Y-axis to favor smaller magnitudes and larger units.`,
				},
				resource.Attribute{
					Name:        "invert_dynamic_legend_hover_control",
					Description: `(Optional) Whether to disable the display of the floating legend (but reenable it when the ctrl-key is pressed).`,
				},
				resource.Attribute{
					Name:        "fixed_legend_enabled",
					Description: `(Optional) Whether to enable a fixed tabular legend adjacent to the chart.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_use_raw_stats",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, the legend uses non-summarized stats instead of summarized.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_position",
					Description: `(Optional) Where the fixed legend should be displayed with respect to the chart. Valid options are ` + "`" + `RIGHT` + "`" + `, ` + "`" + `TOP` + "`" + `, ` + "`" + `LEFT` + "`" + `, ` + "`" + `BOTTOM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_display_stats",
					Description: `(Optional) For a chart with a fixed legend, a list of statistics to display in the legend.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_sort",
					Description: `(Optional) Whether to display ` + "`" + `TOP` + "`" + ` or ` + "`" + `BOTTOM` + "`" + ` ranked series in a fixed legend. Valid options are ` + "`" + `TOP` + "`" + `, and ` + "`" + `BOTTOM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_limit",
					Description: `(Optional) Number of series to include in the fixed legend.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_field",
					Description: `(Optional) Statistic to use for determining whether a series is displayed on the fixed legend. Valid options are ` + "`" + `CURRENT` + "`" + `, ` + "`" + `MEAN` + "`" + `, ` + "`" + `MEDIAN` + "`" + `, ` + "`" + `SUM` + "`" + `, ` + "`" + `MIN` + "`" + `, ` + "`" + `MAX` + "`" + `, and ` + "`" + `COUNT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_hide_label",
					Description: `(Optional) This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "xmax",
					Description: `For x-y scatterplots, max value for the X-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "ymax",
					Description: `For x-y scatterplots, max value for the Y-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "xmin",
					Description: `For x-y scatterplots, min value for the X-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "ymin",
					Description: `For x-y scatterplots, min value for the Y-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "time_based_coloring",
					Description: `For x-y scatterplots, whether to color more recent points as darker than older points.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_value_type",
					Description: `For the single stat view, where to display the name of the query or the value of the query. Valid options are ` + "`" + `VALUE` + "`" + ` or ` + "`" + `LABEL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_color",
					Description: `For the single stat view, the color of the displayed text (when not dynamically determined). Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_vertical_position",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_horizontal_position",
					Description: `For the single stat view, the horizontal position of the displayed text. Valid options are ` + "`" + `MIDDLE` + "`" + `, ` + "`" + `LEFT` + "`" + `, ` + "`" + `RIGHT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_font_size",
					Description: `For the single stat view, the font size of the displayed text, in percent.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_prefix",
					Description: `For the single stat view, a string to add before the displayed text.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_postfix",
					Description: `For the single stat view, a string to append to the displayed text.`,
				},
				resource.Attribute{
					Name:        "sparkline_size",
					Description: `For the single stat view, this determines whether the sparkline of the statistic is displayed in the chart. Valid options are ` + "`" + `BACKGROUND` + "`" + `, ` + "`" + `BOTTOM` + "`" + `, ` + "`" + `NONE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_line_color",
					Description: `For the single stat view, the color of the line. Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Min value of the Y-axis. Set to null or leave blank for auto.`,
				},
				resource.Attribute{
					Name:        "plain_markdown_content",
					Description: `The markdown content for a Markdown display, in plain text.`,
				},
				resource.Attribute{
					Name:        "sparkline_fill_color",
					Description: `For the single stat view, the color of the background fill. Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_colors",
					Description: `For the single stat view, a list of colors that differing query values map to. Must contain one more element than ` + "`" + `sparkline_value_color_map_values_v2` + "`" + `. Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_values_v2",
					Description: `For the single stat view, a list of boundaries for mapping different query values to colors. Must contain one element less than ` + "`" + `sparkline_value_color_map_colors` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_values",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_apply_to",
					Description: `For the single stat view, whether to apply dynamic color settings to the displayed ` + "`" + `TEXT` + "`" + ` or ` + "`" + `BACKGROUND` + "`" + `. Valid options are ` + "`" + `TEXT` + "`" + ` or ` + "`" + `BACKGROUND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_decimal_precision",
					Description: `For the single stat view, the decimal precision of the displayed number.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_text_map_text",
					Description: `For the single stat view, a list of display text values that different query values map to. Must contain one more element than ` + "`" + `sparkline_value_text_map_thresholds` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_text_map_thresholds",
					Description: `For the single stat view, a list of threshold boundaries for mapping different query values to display text. Must contain one element less than ` + "`" + `sparkline_value_text_map_text` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expected_data_spacing",
					Description: `Threshold (in seconds) for time delta between consecutive points in a series above which a dotted line will replace a solid in line plots. Default is 60. ### Parameter Details The ` + "`" + `parameter_details` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label for the parameter.`,
				},
				resource.Attribute{
					Name:        "values_to_readable_strings",
					Description: `A string to string map. At least one of the keys must match the value of ` + "`" + `default_value` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the parameters.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `The default value of the parameter.`,
				},
				resource.Attribute{
					Name:        "hide_from_view",
					Description: `If ` + "`" + `true` + "`" + ` the parameter will only be shown on the edit view of the dashboard.`,
				},
				resource.Attribute{
					Name:        "parameter_type",
					Description: `The type of the parameter. ` + "`" + `SIMPLE` + "`" + `, ` + "`" + `LIST` + "`" + `, or ` + "`" + `DYNAMIC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `For ` + "`" + `TAG_KEY` + "`" + ` dynamic field types, the tag key to return.`,
				},
				resource.Attribute{
					Name:        "query_value",
					Description: `For ` + "`" + `DYNAMIC` + "`" + ` parameter types, the query to execute to return values.`,
				},
				resource.Attribute{
					Name:        "dynamic_field_type",
					Description: `For ` + "`" + `DYNAMIC` + "`" + ` parameter types, the type of the field. Valid options are ` + "`" + `SOURCE` + "`" + `, ` + "`" + `SOURCE_TAG` + "`" + `, ` + "`" + `METRIC_NAME` + "`" + `, ` + "`" + `TAG_KEY` + "`" + `, and ` + "`" + `MATCHING_SOURCE_TAG` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the dashboard.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags to assign to this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description of the dashboard.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Unique identifier, also a URL slug of the dashboard.`,
				},
				resource.Attribute{
					Name:        "section",
					Description: `Dashboard chart sections. See [dashboard sections](#dashboard-sections).`,
				},
				resource.Attribute{
					Name:        "display_query_parameters",
					Description: `Whether the dashboard parameters section is opened by default when the dashboard is shown.`,
				},
				resource.Attribute{
					Name:        "parameter_details",
					Description: `The current JSON representation of dashboard parameters. See [parameter details](#parameter-details).`,
				},
				resource.Attribute{
					Name:        "display_section_table_of_contents",
					Description: `Whether the "pills" quick-linked the sections of the dashboard are displayed by default when the dashboard is shown.`,
				},
				resource.Attribute{
					Name:        "can_modify",
					Description: `A list of users that have modify ACL access to the dashboard.`,
				},
				resource.Attribute{
					Name:        "can_view",
					Description: `A list of users that have view ACL access to the dashboard.`,
				},
				resource.Attribute{
					Name:        "event_filter_type",
					Description: `How charts belonging to this dashboard should display events. ` + "`" + `BYCHART` + "`" + ` is default if unspecified. Valid options are: ` + "`" + `BYCHART` + "`" + `, ` + "`" + `AUTOMATIC` + "`" + `, ` + "`" + `ALL` + "`" + `, ` + "`" + `NONE` + "`" + `, ` + "`" + `BYDASHBOARD` + "`" + `, and ` + "`" + `BYCHARTANDDASHBOARD` + "`" + `. ### Dashboard Sections The ` + "`" + `section` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of this section.`,
				},
				resource.Attribute{
					Name:        "row",
					Description: `See [dashboard section rows](#dashboard-section-rows). ### Dashboard Section Rows The ` + "`" + `row` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "chart",
					Description: `Charts in this section. See [dashboard chart](#dashboard-chart). ### Dashboard Chart The ` + "`" + `chart` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Query expression to plot on the chart. See [chart source queries](#chart-source-queries).`,
				},
				resource.Attribute{
					Name:        "chart_setting",
					Description: `Chart settings. See [chart settings](#chart-settings).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the source.`,
				},
				resource.Attribute{
					Name:        "units",
					Description: `String to label the units of the chart on the Y-Axis.`,
				},
				resource.Attribute{
					Name:        "summarization",
					Description: `Summarization strategy for the chart. MEAN is default.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the chart.`,
				},
				resource.Attribute{
					Name:        "base",
					Description: `The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale. ### Chart Source Queries The ` + "`" + `source` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the source.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `Query expression to plot on the chart.`,
				},
				resource.Attribute{
					Name:        "source_description",
					Description: `A description for the purpose of this source.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Whether the source is disabled.`,
				},
				resource.Attribute{
					Name:        "scatter_plot_source",
					Description: `For scatter plots, does this query source the X-axis or the Y-axis, ` + "`" + `X` + "`" + `, or ` + "`" + `Y` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "query_builder_enabled",
					Description: `Whether or not this source line should have the query builder enabled. ### Chart Settings The ` + "`" + `chart_setting` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Chart Type. ` + "`" + `line` + "`" + ` refers to the Line Plot, ` + "`" + `scatter` + "`" + ` to the Point Plot, ` + "`" + `stacked-area` + "`" + ` to the Stacked Area plot, ` + "`" + `table` + "`" + ` to the Tabular View, ` + "`" + `scatterplot-xy` + "`" + ` to Scatter Plot, ` + "`" + `markdown-widget` + "`" + ` to the Markdown display, and ` + "`" + `sparkline` + "`" + ` to the Single Stat view. Valid options are ` + "`" + `line` + "`" + `, ` + "`" + `scatterplot` + "`" + `, ` + "`" + `stacked-area` + "`" + `, ` + "`" + `stacked-column` + "`" + `, ` + "`" + `table` + "`" + `, ` + "`" + `scatterplot-xy` + "`" + `, ` + "`" + `markdown-widget` + "`" + `, ` + "`" + `sparkline` + "`" + `, ` + "`" + `globe` + "`" + `, ` + "`" + `nodemap` + "`" + `, ` + "`" + `top-k` + "`" + `, ` + "`" + `status-list` + "`" + `, and ` + "`" + `histogram` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `Max value of the Y-axis. Set to null or leave blank for auto.`,
				},
				resource.Attribute{
					Name:        "line_type",
					Description: `Plot interpolation type. ` + "`" + `linear` + "`" + ` is default. Valid options are ` + "`" + `linear` + "`" + `, ` + "`" + `step-before` + "`" + `, ` + "`" + `step-after` + "`" + `, ` + "`" + `basis` + "`" + `, ` + "`" + `cardinal` + "`" + `, and ` + "`" + `monotone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "stack_type",
					Description: `Type of stacked chart (applicable only if chart type is ` + "`" + `stacked` + "`" + `). ` + "`" + `zero` + "`" + ` (default) means stacked from y=0. ` + "`" + `expand` + "`" + ` means normalized from 0 to 1. ` + "`" + `wiggle` + "`" + ` means minimize weighted changes. ` + "`" + `silhouette` + "`" + ` means to center the stream. Valid options are ` + "`" + `zero` + "`" + `, ` + "`" + `expand` + "`" + `, ` + "`" + `wiggle` + "`" + `, ` + "`" + `silhouette` + "`" + `, and ` + "`" + `bars` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "windowing",
					Description: `For the tabular view, whether to use the full time window for the query or the last X minutes. Valid options are ` + "`" + `full` + "`" + ` or ` + "`" + `last` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "window_size",
					Description: `Width, in minutes, of the time window to use for ` + "`" + `last` + "`" + ` windowing.`,
				},
				resource.Attribute{
					Name:        "show_hosts",
					Description: `For the tabular view, whether to display sources. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_labels",
					Description: `For the tabular view, whether to display labels. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_raw_values",
					Description: `For the tabular view, whether to display raw values. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_column_tags",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "column_tags",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "tag_mode",
					Description: `For the tabular view, which mode to use to determine which point tags to display. Valid options are ` + "`" + `all` + "`" + `, ` + "`" + `top` + "`" + `, or ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "num_tags",
					Description: `For the tabular view defines how many point tags to display.`,
				},
				resource.Attribute{
					Name:        "custom_tags",
					Description: `For the tabular view, a list of point tags to display when using the ` + "`" + `custom` + "`" + ` tag display mode.`,
				},
				resource.Attribute{
					Name:        "group_by_source",
					Description: `For the tabular view, whether to group multi metrics into a single row by a common source. If set to ` + "`" + `false` + "`" + `, each source is displayed in its own row. If set to ` + "`" + `true` + "`" + `, multiple metrics for the same host are displayed as different columns in the same row.`,
				},
				resource.Attribute{
					Name:        "sort_values_descending",
					Description: `For the tabular view, whether to display values in descending order. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "y1max",
					Description: `For plots with multiple Y-axes, max value for the right side Y-axis. Set null for auto.`,
				},
				resource.Attribute{
					Name:        "y1min",
					Description: `For plots with multiple Y-axes, min value for the right side Y-axis. Set null for auto.`,
				},
				resource.Attribute{
					Name:        "y1_units",
					Description: `For plots with multiple Y-axes, units for right side Y-axis.`,
				},
				resource.Attribute{
					Name:        "y0_scale_si_by_1024",
					Description: `(Optional) Whether to scale numerical magnitude labels for left Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).`,
				},
				resource.Attribute{
					Name:        "y1_scale_si_by_1024",
					Description: `(Optional) Whether to scale numerical magnitude labels for right Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).`,
				},
				resource.Attribute{
					Name:        "y0_unit_autoscaling",
					Description: `(Optional) Whether to automatically adjust magnitude labels and units for the left Y-axis to favor smaller magnitudes and larger units.`,
				},
				resource.Attribute{
					Name:        "y1_unit_autoscaling",
					Description: `(Optional) Whether to automatically adjust magnitude labels and units for the right Y-axis to favor smaller magnitudes and larger units.`,
				},
				resource.Attribute{
					Name:        "invert_dynamic_legend_hover_control",
					Description: `(Optional) Whether to disable the display of the floating legend (but reenable it when the ctrl-key is pressed).`,
				},
				resource.Attribute{
					Name:        "fixed_legend_enabled",
					Description: `(Optional) Whether to enable a fixed tabular legend adjacent to the chart.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_use_raw_stats",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, the legend uses non-summarized stats instead of summarized.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_position",
					Description: `(Optional) Where the fixed legend should be displayed with respect to the chart. Valid options are ` + "`" + `RIGHT` + "`" + `, ` + "`" + `TOP` + "`" + `, ` + "`" + `LEFT` + "`" + `, ` + "`" + `BOTTOM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_display_stats",
					Description: `(Optional) For a chart with a fixed legend, a list of statistics to display in the legend.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_sort",
					Description: `(Optional) Whether to display ` + "`" + `TOP` + "`" + ` or ` + "`" + `BOTTOM` + "`" + ` ranked series in a fixed legend. Valid options are ` + "`" + `TOP` + "`" + `, and ` + "`" + `BOTTOM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_limit",
					Description: `(Optional) Number of series to include in the fixed legend.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_field",
					Description: `(Optional) Statistic to use for determining whether a series is displayed on the fixed legend. Valid options are ` + "`" + `CURRENT` + "`" + `, ` + "`" + `MEAN` + "`" + `, ` + "`" + `MEDIAN` + "`" + `, ` + "`" + `SUM` + "`" + `, ` + "`" + `MIN` + "`" + `, ` + "`" + `MAX` + "`" + `, and ` + "`" + `COUNT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_hide_label",
					Description: `(Optional) This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "xmax",
					Description: `For x-y scatterplots, max value for the X-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "ymax",
					Description: `For x-y scatterplots, max value for the Y-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "xmin",
					Description: `For x-y scatterplots, min value for the X-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "ymin",
					Description: `For x-y scatterplots, min value for the Y-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "time_based_coloring",
					Description: `For x-y scatterplots, whether to color more recent points as darker than older points.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_value_type",
					Description: `For the single stat view, where to display the name of the query or the value of the query. Valid options are ` + "`" + `VALUE` + "`" + ` or ` + "`" + `LABEL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_color",
					Description: `For the single stat view, the color of the displayed text (when not dynamically determined). Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_vertical_position",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_horizontal_position",
					Description: `For the single stat view, the horizontal position of the displayed text. Valid options are ` + "`" + `MIDDLE` + "`" + `, ` + "`" + `LEFT` + "`" + `, ` + "`" + `RIGHT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_font_size",
					Description: `For the single stat view, the font size of the displayed text, in percent.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_prefix",
					Description: `For the single stat view, a string to add before the displayed text.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_postfix",
					Description: `For the single stat view, a string to append to the displayed text.`,
				},
				resource.Attribute{
					Name:        "sparkline_size",
					Description: `For the single stat view, this determines whether the sparkline of the statistic is displayed in the chart. Valid options are ` + "`" + `BACKGROUND` + "`" + `, ` + "`" + `BOTTOM` + "`" + `, ` + "`" + `NONE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_line_color",
					Description: `For the single stat view, the color of the line. Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Min value of the Y-axis. Set to null or leave blank for auto.`,
				},
				resource.Attribute{
					Name:        "plain_markdown_content",
					Description: `The markdown content for a Markdown display, in plain text.`,
				},
				resource.Attribute{
					Name:        "sparkline_fill_color",
					Description: `For the single stat view, the color of the background fill. Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_colors",
					Description: `For the single stat view, a list of colors that differing query values map to. Must contain one more element than ` + "`" + `sparkline_value_color_map_values_v2` + "`" + `. Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_values_v2",
					Description: `For the single stat view, a list of boundaries for mapping different query values to colors. Must contain one element less than ` + "`" + `sparkline_value_color_map_colors` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_values",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_apply_to",
					Description: `For the single stat view, whether to apply dynamic color settings to the displayed ` + "`" + `TEXT` + "`" + ` or ` + "`" + `BACKGROUND` + "`" + `. Valid options are ` + "`" + `TEXT` + "`" + ` or ` + "`" + `BACKGROUND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_decimal_precision",
					Description: `For the single stat view, the decimal precision of the displayed number.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_text_map_text",
					Description: `For the single stat view, a list of display text values that different query values map to. Must contain one more element than ` + "`" + `sparkline_value_text_map_thresholds` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_text_map_thresholds",
					Description: `For the single stat view, a list of threshold boundaries for mapping different query values to display text. Must contain one element less than ` + "`" + `sparkline_value_text_map_text` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expected_data_spacing",
					Description: `Threshold (in seconds) for time delta between consecutive points in a series above which a dotted line will replace a solid in line plots. Default is 60. ### Parameter Details The ` + "`" + `parameter_details` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label for the parameter.`,
				},
				resource.Attribute{
					Name:        "values_to_readable_strings",
					Description: `A string to string map. At least one of the keys must match the value of ` + "`" + `default_value` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the parameters.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `The default value of the parameter.`,
				},
				resource.Attribute{
					Name:        "hide_from_view",
					Description: `If ` + "`" + `true` + "`" + ` the parameter will only be shown on the edit view of the dashboard.`,
				},
				resource.Attribute{
					Name:        "parameter_type",
					Description: `The type of the parameter. ` + "`" + `SIMPLE` + "`" + `, ` + "`" + `LIST` + "`" + `, or ` + "`" + `DYNAMIC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `For ` + "`" + `TAG_KEY` + "`" + ` dynamic field types, the tag key to return.`,
				},
				resource.Attribute{
					Name:        "query_value",
					Description: `For ` + "`" + `DYNAMIC` + "`" + ` parameter types, the query to execute to return values.`,
				},
				resource.Attribute{
					Name:        "dynamic_field_type",
					Description: `For ` + "`" + `DYNAMIC` + "`" + ` parameter types, the type of the field. Valid options are ` + "`" + `SOURCE` + "`" + `, ` + "`" + `SOURCE_TAG` + "`" + `, ` + "`" + `METRIC_NAME` + "`" + `, ` + "`" + `TAG_KEY` + "`" + `, and ` + "`" + `MATCHING_SOURCE_TAG` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_dashboards",
			Category:         "Data Sources",
			ShortDescription: `Get the information about all Wavefront dashboards.`,
			Description: `

Use this data source to get information about all Wavefront dashboards.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Limit is the maximum number of results to be returned. Defaults to 100.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Offset is the offset from the first result to be returned. Defaults to 0. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get the information about all dashboards. data "wavefront_dashboards" "example" { limit = 10 offset = 0 } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "dashboards",
					Description: `List of all Wavefront dashboards. For each dashboard you will see a list of attributes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the dashboard.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags to assign to this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description of the dashboard.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Unique identifier, also a URL slug of the dashboard.`,
				},
				resource.Attribute{
					Name:        "section",
					Description: `Dashboard chart sections. See [dashboard sections](#dashboard-sections).`,
				},
				resource.Attribute{
					Name:        "display_query_parameters",
					Description: `Whether the dashboard parameters section is opened by default when the dashboard is shown.`,
				},
				resource.Attribute{
					Name:        "parameter_details",
					Description: `The current JSON representation of dashboard parameters. See [parameter details](#parameter-details).`,
				},
				resource.Attribute{
					Name:        "display_section_table_of_contents",
					Description: `Whether the "pills" quick-linked the sections of the dashboard are displayed by default when the dashboard is shown.`,
				},
				resource.Attribute{
					Name:        "can_modify",
					Description: `A list of users that have modify ACL access to the dashboard.`,
				},
				resource.Attribute{
					Name:        "can_view",
					Description: `A list of users that have view ACL access to the dashboard.`,
				},
				resource.Attribute{
					Name:        "event_filter_type",
					Description: `How charts belonging to this dashboard should display events. ` + "`" + `BYCHART` + "`" + ` is default if unspecified. Valid options are: ` + "`" + `BYCHART` + "`" + `, ` + "`" + `AUTOMATIC` + "`" + `, ` + "`" + `ALL` + "`" + `, ` + "`" + `NONE` + "`" + `, ` + "`" + `BYDASHBOARD` + "`" + `, and ` + "`" + `BYCHARTANDDASHBOARD` + "`" + `. ### Dashboard Sections The ` + "`" + `section` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of this section.`,
				},
				resource.Attribute{
					Name:        "row",
					Description: `See [dashboard section rows](#dashboard-section-rows). ### Dashboard Section Rows The ` + "`" + `row` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "chart",
					Description: `Charts in this section. See [dashboard chart](#dashboard-chart). ### Dashboard Chart The ` + "`" + `chart` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Query expression to plot on the chart. See [chart source queries](#chart-source-queries).`,
				},
				resource.Attribute{
					Name:        "chart_setting",
					Description: `Chart settings. See [chart settings](#chart-settings).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the source.`,
				},
				resource.Attribute{
					Name:        "units",
					Description: `String to label the units of the chart on the Y-Axis.`,
				},
				resource.Attribute{
					Name:        "summarization",
					Description: `Summarization strategy for the chart. MEAN is default.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the chart.`,
				},
				resource.Attribute{
					Name:        "base",
					Description: `The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale. ### Chart Source Queries The ` + "`" + `source` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the source.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `Query expression to plot on the chart.`,
				},
				resource.Attribute{
					Name:        "source_description",
					Description: `A description for the purpose of this source.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Whether the source is disabled.`,
				},
				resource.Attribute{
					Name:        "scatter_plot_source",
					Description: `For scatter plots, does this query source the X-axis or the Y-axis, ` + "`" + `X` + "`" + `, or ` + "`" + `Y` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "query_builder_enabled",
					Description: `Whether or not this source line should have the query builder enabled. ### Chart Settings The ` + "`" + `chart_setting` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Chart Type. ` + "`" + `line` + "`" + ` refers to the Line Plot, ` + "`" + `scatter` + "`" + ` to the Point Plot, ` + "`" + `stacked-area` + "`" + ` to the Stacked Area plot, ` + "`" + `table` + "`" + ` to the Tabular View, ` + "`" + `scatterplot-xy` + "`" + ` to Scatter Plot, ` + "`" + `markdown-widget` + "`" + ` to the Markdown display, and ` + "`" + `sparkline` + "`" + ` to the Single Stat view. Valid options are ` + "`" + `line` + "`" + `, ` + "`" + `scatterplot` + "`" + `, ` + "`" + `stacked-area` + "`" + `, ` + "`" + `stacked-column` + "`" + `, ` + "`" + `table` + "`" + `, ` + "`" + `scatterplot-xy` + "`" + `, ` + "`" + `markdown-widget` + "`" + `, ` + "`" + `sparkline` + "`" + `, ` + "`" + `globe` + "`" + `, ` + "`" + `nodemap` + "`" + `, ` + "`" + `top-k` + "`" + `, ` + "`" + `status-list` + "`" + `, and ` + "`" + `histogram` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `Max value of the Y-axis. Set to null or leave blank for auto.`,
				},
				resource.Attribute{
					Name:        "line_type",
					Description: `Plot interpolation type. ` + "`" + `linear` + "`" + ` is default. Valid options are ` + "`" + `linear` + "`" + `, ` + "`" + `step-before` + "`" + `, ` + "`" + `step-after` + "`" + `, ` + "`" + `basis` + "`" + `, ` + "`" + `cardinal` + "`" + `, and ` + "`" + `monotone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "stack_type",
					Description: `Type of stacked chart (applicable only if chart type is ` + "`" + `stacked` + "`" + `). ` + "`" + `zero` + "`" + ` (default) means stacked from y=0. ` + "`" + `expand` + "`" + ` means normalized from 0 to 1. ` + "`" + `wiggle` + "`" + ` means minimize weighted changes. ` + "`" + `silhouette` + "`" + ` means to center the stream. Valid options are ` + "`" + `zero` + "`" + `, ` + "`" + `expand` + "`" + `, ` + "`" + `wiggle` + "`" + `, ` + "`" + `silhouette` + "`" + `, and ` + "`" + `bars` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "windowing",
					Description: `For the tabular view, whether to use the full time window for the query or the last X minutes. Valid options are ` + "`" + `full` + "`" + ` or ` + "`" + `last` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "window_size",
					Description: `Width, in minutes, of the time window to use for ` + "`" + `last` + "`" + ` windowing.`,
				},
				resource.Attribute{
					Name:        "show_hosts",
					Description: `For the tabular view, whether to display sources. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_labels",
					Description: `For the tabular view, whether to display labels. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_raw_values",
					Description: `For the tabular view, whether to display raw values. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_column_tags",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "column_tags",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "tag_mode",
					Description: `For the tabular view, which mode to use to determine which point tags to display. Valid options are ` + "`" + `all` + "`" + `, ` + "`" + `top` + "`" + `, or ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "num_tags",
					Description: `For the tabular view defines how many point tags to display.`,
				},
				resource.Attribute{
					Name:        "custom_tags",
					Description: `For the tabular view, a list of point tags to display when using the ` + "`" + `custom` + "`" + ` tag display mode.`,
				},
				resource.Attribute{
					Name:        "group_by_source",
					Description: `For the tabular view, whether to group multi metrics into a single row by a common source. If set to ` + "`" + `false` + "`" + `, each source is displayed in its own row. If set to ` + "`" + `true` + "`" + `, multiple metrics for the same host are displayed as different columns in the same row.`,
				},
				resource.Attribute{
					Name:        "sort_values_descending",
					Description: `For the tabular view, whether to display values in descending order. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "y1max",
					Description: `For plots with multiple Y-axes, max value for the right side Y-axis. Set null for auto.`,
				},
				resource.Attribute{
					Name:        "y1min",
					Description: `For plots with multiple Y-axes, min value for the right side Y-axis. Set null for auto.`,
				},
				resource.Attribute{
					Name:        "y1_units",
					Description: `For plots with multiple Y-axes, units for right side Y-axis.`,
				},
				resource.Attribute{
					Name:        "y0_scale_si_by_1024",
					Description: `(Optional) Whether to scale numerical magnitude labels for left Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).`,
				},
				resource.Attribute{
					Name:        "y1_scale_si_by_1024",
					Description: `(Optional) Whether to scale numerical magnitude labels for right Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).`,
				},
				resource.Attribute{
					Name:        "y0_unit_autoscaling",
					Description: `(Optional) Whether to automatically adjust magnitude labels and units for the left Y-axis to favor smaller magnitudes and larger units.`,
				},
				resource.Attribute{
					Name:        "y1_unit_autoscaling",
					Description: `(Optional) Whether to automatically adjust magnitude labels and units for the right Y-axis to favor smaller magnitudes and larger units.`,
				},
				resource.Attribute{
					Name:        "invert_dynamic_legend_hover_control",
					Description: `(Optional) Whether to disable the display of the floating legend (but reenable it when the ctrl-key is pressed).`,
				},
				resource.Attribute{
					Name:        "fixed_legend_enabled",
					Description: `(Optional) Whether to enable a fixed tabular legend adjacent to the chart.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_use_raw_stats",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, the legend uses non-summarized stats instead of summarized.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_position",
					Description: `(Optional) Where the fixed legend should be displayed with respect to the chart. Valid options are ` + "`" + `RIGHT` + "`" + `, ` + "`" + `TOP` + "`" + `, ` + "`" + `LEFT` + "`" + `, ` + "`" + `BOTTOM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_display_stats",
					Description: `(Optional) For a chart with a fixed legend, a list of statistics to display in the legend.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_sort",
					Description: `(Optional) Whether to display ` + "`" + `TOP` + "`" + ` or ` + "`" + `BOTTOM` + "`" + ` ranked series in a fixed legend. Valid options are ` + "`" + `TOP` + "`" + `, and ` + "`" + `BOTTOM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_limit",
					Description: `(Optional) Number of series to include in the fixed legend.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_field",
					Description: `(Optional) Statistic to use for determining whether a series is displayed on the fixed legend. Valid options are ` + "`" + `CURRENT` + "`" + `, ` + "`" + `MEAN` + "`" + `, ` + "`" + `MEDIAN` + "`" + `, ` + "`" + `SUM` + "`" + `, ` + "`" + `MIN` + "`" + `, ` + "`" + `MAX` + "`" + `, and ` + "`" + `COUNT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_hide_label",
					Description: `(Optional) This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "xmax",
					Description: `For x-y scatterplots, max value for the X-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "ymax",
					Description: `For x-y scatterplots, max value for the Y-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "xmin",
					Description: `For x-y scatterplots, min value for the X-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "ymin",
					Description: `For x-y scatterplots, min value for the Y-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "time_based_coloring",
					Description: `For x-y scatterplots, whether to color more recent points as darker than older points.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_value_type",
					Description: `For the single stat view, where to display the name of the query or the value of the query. Valid options are ` + "`" + `VALUE` + "`" + ` or ` + "`" + `LABEL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_color",
					Description: `For the single stat view, the color of the displayed text (when not dynamically determined). Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_vertical_position",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_horizontal_position",
					Description: `For the single stat view, the horizontal position of the displayed text. Valid options are ` + "`" + `MIDDLE` + "`" + `, ` + "`" + `LEFT` + "`" + `, ` + "`" + `RIGHT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_font_size",
					Description: `For the single stat view, the font size of the displayed text, in percent.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_prefix",
					Description: `For the single stat view, a string to add before the displayed text.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_postfix",
					Description: `For the single stat view, a string to append to the displayed text.`,
				},
				resource.Attribute{
					Name:        "sparkline_size",
					Description: `For the single stat view, this determines whether the sparkline of the statistic is displayed in the chart. Valid options are ` + "`" + `BACKGROUND` + "`" + `, ` + "`" + `BOTTOM` + "`" + `, ` + "`" + `NONE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_line_color",
					Description: `For the single stat view, the color of the line. Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Min value of the Y-axis. Set to null or leave blank for auto.`,
				},
				resource.Attribute{
					Name:        "plain_markdown_content",
					Description: `The markdown content for a Markdown display, in plain text.`,
				},
				resource.Attribute{
					Name:        "sparkline_fill_color",
					Description: `For the single stat view, the color of the background fill. Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_colors",
					Description: `For the single stat view, a list of colors that differing query values map to. Must contain one more element than ` + "`" + `sparkline_value_color_map_values_v2` + "`" + `. Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_values_v2",
					Description: `For the single stat view, a list of boundaries for mapping different query values to colors. Must contain one element less than ` + "`" + `sparkline_value_color_map_colors` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_values",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_apply_to",
					Description: `For the single stat view, whether to apply dynamic color settings to the displayed ` + "`" + `TEXT` + "`" + ` or ` + "`" + `BACKGROUND` + "`" + `. Valid options are ` + "`" + `TEXT` + "`" + ` or ` + "`" + `BACKGROUND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_decimal_precision",
					Description: `For the single stat view, the decimal precision of the displayed number.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_text_map_text",
					Description: `For the single stat view, a list of display text values that different query values map to. Must contain one more element than ` + "`" + `sparkline_value_text_map_thresholds` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_text_map_thresholds",
					Description: `For the single stat view, a list of threshold boundaries for mapping different query values to display text. Must contain one element less than ` + "`" + `sparkline_value_text_map_text` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expected_data_spacing",
					Description: `Threshold (in seconds) for time delta between consecutive points in a series above which a dotted line will replace a solid in line plots. Default is 60. ### Parameter Details The ` + "`" + `parameter_details` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label for the parameter.`,
				},
				resource.Attribute{
					Name:        "values_to_readable_strings",
					Description: `A string to string map. At least one of the keys must match the value of ` + "`" + `default_value` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the parameters.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `The default value of the parameter.`,
				},
				resource.Attribute{
					Name:        "hide_from_view",
					Description: `If ` + "`" + `true` + "`" + ` the parameter will only be shown on the edit view of the dashboard.`,
				},
				resource.Attribute{
					Name:        "parameter_type",
					Description: `The type of the parameter. ` + "`" + `SIMPLE` + "`" + `, ` + "`" + `LIST` + "`" + `, or ` + "`" + `DYNAMIC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `For ` + "`" + `TAG_KEY` + "`" + ` dynamic field types, the tag key to return.`,
				},
				resource.Attribute{
					Name:        "query_value",
					Description: `For ` + "`" + `DYNAMIC` + "`" + ` parameter types, the query to execute to return values.`,
				},
				resource.Attribute{
					Name:        "dynamic_field_type",
					Description: `For ` + "`" + `DYNAMIC` + "`" + ` parameter types, the type of the field. Valid options are ` + "`" + `SOURCE` + "`" + `, ` + "`" + `SOURCE_TAG` + "`" + `, ` + "`" + `METRIC_NAME` + "`" + `, ` + "`" + `TAG_KEY` + "`" + `, and ` + "`" + `MATCHING_SOURCE_TAG` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dashboards",
					Description: `List of all Wavefront dashboards. For each dashboard you will see a list of attributes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the dashboard.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags to assign to this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description of the dashboard.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Unique identifier, also a URL slug of the dashboard.`,
				},
				resource.Attribute{
					Name:        "section",
					Description: `Dashboard chart sections. See [dashboard sections](#dashboard-sections).`,
				},
				resource.Attribute{
					Name:        "display_query_parameters",
					Description: `Whether the dashboard parameters section is opened by default when the dashboard is shown.`,
				},
				resource.Attribute{
					Name:        "parameter_details",
					Description: `The current JSON representation of dashboard parameters. See [parameter details](#parameter-details).`,
				},
				resource.Attribute{
					Name:        "display_section_table_of_contents",
					Description: `Whether the "pills" quick-linked the sections of the dashboard are displayed by default when the dashboard is shown.`,
				},
				resource.Attribute{
					Name:        "can_modify",
					Description: `A list of users that have modify ACL access to the dashboard.`,
				},
				resource.Attribute{
					Name:        "can_view",
					Description: `A list of users that have view ACL access to the dashboard.`,
				},
				resource.Attribute{
					Name:        "event_filter_type",
					Description: `How charts belonging to this dashboard should display events. ` + "`" + `BYCHART` + "`" + ` is default if unspecified. Valid options are: ` + "`" + `BYCHART` + "`" + `, ` + "`" + `AUTOMATIC` + "`" + `, ` + "`" + `ALL` + "`" + `, ` + "`" + `NONE` + "`" + `, ` + "`" + `BYDASHBOARD` + "`" + `, and ` + "`" + `BYCHARTANDDASHBOARD` + "`" + `. ### Dashboard Sections The ` + "`" + `section` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of this section.`,
				},
				resource.Attribute{
					Name:        "row",
					Description: `See [dashboard section rows](#dashboard-section-rows). ### Dashboard Section Rows The ` + "`" + `row` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "chart",
					Description: `Charts in this section. See [dashboard chart](#dashboard-chart). ### Dashboard Chart The ` + "`" + `chart` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Query expression to plot on the chart. See [chart source queries](#chart-source-queries).`,
				},
				resource.Attribute{
					Name:        "chart_setting",
					Description: `Chart settings. See [chart settings](#chart-settings).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the source.`,
				},
				resource.Attribute{
					Name:        "units",
					Description: `String to label the units of the chart on the Y-Axis.`,
				},
				resource.Attribute{
					Name:        "summarization",
					Description: `Summarization strategy for the chart. MEAN is default.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the chart.`,
				},
				resource.Attribute{
					Name:        "base",
					Description: `The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale. ### Chart Source Queries The ` + "`" + `source` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the source.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `Query expression to plot on the chart.`,
				},
				resource.Attribute{
					Name:        "source_description",
					Description: `A description for the purpose of this source.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Whether the source is disabled.`,
				},
				resource.Attribute{
					Name:        "scatter_plot_source",
					Description: `For scatter plots, does this query source the X-axis or the Y-axis, ` + "`" + `X` + "`" + `, or ` + "`" + `Y` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "query_builder_enabled",
					Description: `Whether or not this source line should have the query builder enabled. ### Chart Settings The ` + "`" + `chart_setting` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Chart Type. ` + "`" + `line` + "`" + ` refers to the Line Plot, ` + "`" + `scatter` + "`" + ` to the Point Plot, ` + "`" + `stacked-area` + "`" + ` to the Stacked Area plot, ` + "`" + `table` + "`" + ` to the Tabular View, ` + "`" + `scatterplot-xy` + "`" + ` to Scatter Plot, ` + "`" + `markdown-widget` + "`" + ` to the Markdown display, and ` + "`" + `sparkline` + "`" + ` to the Single Stat view. Valid options are ` + "`" + `line` + "`" + `, ` + "`" + `scatterplot` + "`" + `, ` + "`" + `stacked-area` + "`" + `, ` + "`" + `stacked-column` + "`" + `, ` + "`" + `table` + "`" + `, ` + "`" + `scatterplot-xy` + "`" + `, ` + "`" + `markdown-widget` + "`" + `, ` + "`" + `sparkline` + "`" + `, ` + "`" + `globe` + "`" + `, ` + "`" + `nodemap` + "`" + `, ` + "`" + `top-k` + "`" + `, ` + "`" + `status-list` + "`" + `, and ` + "`" + `histogram` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `Max value of the Y-axis. Set to null or leave blank for auto.`,
				},
				resource.Attribute{
					Name:        "line_type",
					Description: `Plot interpolation type. ` + "`" + `linear` + "`" + ` is default. Valid options are ` + "`" + `linear` + "`" + `, ` + "`" + `step-before` + "`" + `, ` + "`" + `step-after` + "`" + `, ` + "`" + `basis` + "`" + `, ` + "`" + `cardinal` + "`" + `, and ` + "`" + `monotone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "stack_type",
					Description: `Type of stacked chart (applicable only if chart type is ` + "`" + `stacked` + "`" + `). ` + "`" + `zero` + "`" + ` (default) means stacked from y=0. ` + "`" + `expand` + "`" + ` means normalized from 0 to 1. ` + "`" + `wiggle` + "`" + ` means minimize weighted changes. ` + "`" + `silhouette` + "`" + ` means to center the stream. Valid options are ` + "`" + `zero` + "`" + `, ` + "`" + `expand` + "`" + `, ` + "`" + `wiggle` + "`" + `, ` + "`" + `silhouette` + "`" + `, and ` + "`" + `bars` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "windowing",
					Description: `For the tabular view, whether to use the full time window for the query or the last X minutes. Valid options are ` + "`" + `full` + "`" + ` or ` + "`" + `last` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "window_size",
					Description: `Width, in minutes, of the time window to use for ` + "`" + `last` + "`" + ` windowing.`,
				},
				resource.Attribute{
					Name:        "show_hosts",
					Description: `For the tabular view, whether to display sources. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_labels",
					Description: `For the tabular view, whether to display labels. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_raw_values",
					Description: `For the tabular view, whether to display raw values. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_column_tags",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "column_tags",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "tag_mode",
					Description: `For the tabular view, which mode to use to determine which point tags to display. Valid options are ` + "`" + `all` + "`" + `, ` + "`" + `top` + "`" + `, or ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "num_tags",
					Description: `For the tabular view defines how many point tags to display.`,
				},
				resource.Attribute{
					Name:        "custom_tags",
					Description: `For the tabular view, a list of point tags to display when using the ` + "`" + `custom` + "`" + ` tag display mode.`,
				},
				resource.Attribute{
					Name:        "group_by_source",
					Description: `For the tabular view, whether to group multi metrics into a single row by a common source. If set to ` + "`" + `false` + "`" + `, each source is displayed in its own row. If set to ` + "`" + `true` + "`" + `, multiple metrics for the same host are displayed as different columns in the same row.`,
				},
				resource.Attribute{
					Name:        "sort_values_descending",
					Description: `For the tabular view, whether to display values in descending order. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "y1max",
					Description: `For plots with multiple Y-axes, max value for the right side Y-axis. Set null for auto.`,
				},
				resource.Attribute{
					Name:        "y1min",
					Description: `For plots with multiple Y-axes, min value for the right side Y-axis. Set null for auto.`,
				},
				resource.Attribute{
					Name:        "y1_units",
					Description: `For plots with multiple Y-axes, units for right side Y-axis.`,
				},
				resource.Attribute{
					Name:        "y0_scale_si_by_1024",
					Description: `(Optional) Whether to scale numerical magnitude labels for left Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).`,
				},
				resource.Attribute{
					Name:        "y1_scale_si_by_1024",
					Description: `(Optional) Whether to scale numerical magnitude labels for right Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).`,
				},
				resource.Attribute{
					Name:        "y0_unit_autoscaling",
					Description: `(Optional) Whether to automatically adjust magnitude labels and units for the left Y-axis to favor smaller magnitudes and larger units.`,
				},
				resource.Attribute{
					Name:        "y1_unit_autoscaling",
					Description: `(Optional) Whether to automatically adjust magnitude labels and units for the right Y-axis to favor smaller magnitudes and larger units.`,
				},
				resource.Attribute{
					Name:        "invert_dynamic_legend_hover_control",
					Description: `(Optional) Whether to disable the display of the floating legend (but reenable it when the ctrl-key is pressed).`,
				},
				resource.Attribute{
					Name:        "fixed_legend_enabled",
					Description: `(Optional) Whether to enable a fixed tabular legend adjacent to the chart.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_use_raw_stats",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, the legend uses non-summarized stats instead of summarized.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_position",
					Description: `(Optional) Where the fixed legend should be displayed with respect to the chart. Valid options are ` + "`" + `RIGHT` + "`" + `, ` + "`" + `TOP` + "`" + `, ` + "`" + `LEFT` + "`" + `, ` + "`" + `BOTTOM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_display_stats",
					Description: `(Optional) For a chart with a fixed legend, a list of statistics to display in the legend.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_sort",
					Description: `(Optional) Whether to display ` + "`" + `TOP` + "`" + ` or ` + "`" + `BOTTOM` + "`" + ` ranked series in a fixed legend. Valid options are ` + "`" + `TOP` + "`" + `, and ` + "`" + `BOTTOM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_limit",
					Description: `(Optional) Number of series to include in the fixed legend.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_field",
					Description: `(Optional) Statistic to use for determining whether a series is displayed on the fixed legend. Valid options are ` + "`" + `CURRENT` + "`" + `, ` + "`" + `MEAN` + "`" + `, ` + "`" + `MEDIAN` + "`" + `, ` + "`" + `SUM` + "`" + `, ` + "`" + `MIN` + "`" + `, ` + "`" + `MAX` + "`" + `, and ` + "`" + `COUNT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_hide_label",
					Description: `(Optional) This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "xmax",
					Description: `For x-y scatterplots, max value for the X-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "ymax",
					Description: `For x-y scatterplots, max value for the Y-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "xmin",
					Description: `For x-y scatterplots, min value for the X-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "ymin",
					Description: `For x-y scatterplots, min value for the Y-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "time_based_coloring",
					Description: `For x-y scatterplots, whether to color more recent points as darker than older points.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_value_type",
					Description: `For the single stat view, where to display the name of the query or the value of the query. Valid options are ` + "`" + `VALUE` + "`" + ` or ` + "`" + `LABEL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_color",
					Description: `For the single stat view, the color of the displayed text (when not dynamically determined). Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_vertical_position",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_horizontal_position",
					Description: `For the single stat view, the horizontal position of the displayed text. Valid options are ` + "`" + `MIDDLE` + "`" + `, ` + "`" + `LEFT` + "`" + `, ` + "`" + `RIGHT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_font_size",
					Description: `For the single stat view, the font size of the displayed text, in percent.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_prefix",
					Description: `For the single stat view, a string to add before the displayed text.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_postfix",
					Description: `For the single stat view, a string to append to the displayed text.`,
				},
				resource.Attribute{
					Name:        "sparkline_size",
					Description: `For the single stat view, this determines whether the sparkline of the statistic is displayed in the chart. Valid options are ` + "`" + `BACKGROUND` + "`" + `, ` + "`" + `BOTTOM` + "`" + `, ` + "`" + `NONE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_line_color",
					Description: `For the single stat view, the color of the line. Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Min value of the Y-axis. Set to null or leave blank for auto.`,
				},
				resource.Attribute{
					Name:        "plain_markdown_content",
					Description: `The markdown content for a Markdown display, in plain text.`,
				},
				resource.Attribute{
					Name:        "sparkline_fill_color",
					Description: `For the single stat view, the color of the background fill. Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_colors",
					Description: `For the single stat view, a list of colors that differing query values map to. Must contain one more element than ` + "`" + `sparkline_value_color_map_values_v2` + "`" + `. Values should be in RGBA format.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_values_v2",
					Description: `For the single stat view, a list of boundaries for mapping different query values to colors. Must contain one element less than ` + "`" + `sparkline_value_color_map_colors` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_values",
					Description: `This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_apply_to",
					Description: `For the single stat view, whether to apply dynamic color settings to the displayed ` + "`" + `TEXT` + "`" + ` or ` + "`" + `BACKGROUND` + "`" + `. Valid options are ` + "`" + `TEXT` + "`" + ` or ` + "`" + `BACKGROUND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_decimal_precision",
					Description: `For the single stat view, the decimal precision of the displayed number.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_text_map_text",
					Description: `For the single stat view, a list of display text values that different query values map to. Must contain one more element than ` + "`" + `sparkline_value_text_map_thresholds` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_text_map_thresholds",
					Description: `For the single stat view, a list of threshold boundaries for mapping different query values to display text. Must contain one element less than ` + "`" + `sparkline_value_text_map_text` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expected_data_spacing",
					Description: `Threshold (in seconds) for time delta between consecutive points in a series above which a dotted line will replace a solid in line plots. Default is 60. ### Parameter Details The ` + "`" + `parameter_details` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label for the parameter.`,
				},
				resource.Attribute{
					Name:        "values_to_readable_strings",
					Description: `A string to string map. At least one of the keys must match the value of ` + "`" + `default_value` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the parameters.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `The default value of the parameter.`,
				},
				resource.Attribute{
					Name:        "hide_from_view",
					Description: `If ` + "`" + `true` + "`" + ` the parameter will only be shown on the edit view of the dashboard.`,
				},
				resource.Attribute{
					Name:        "parameter_type",
					Description: `The type of the parameter. ` + "`" + `SIMPLE` + "`" + `, ` + "`" + `LIST` + "`" + `, or ` + "`" + `DYNAMIC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `For ` + "`" + `TAG_KEY` + "`" + ` dynamic field types, the tag key to return.`,
				},
				resource.Attribute{
					Name:        "query_value",
					Description: `For ` + "`" + `DYNAMIC` + "`" + ` parameter types, the query to execute to return values.`,
				},
				resource.Attribute{
					Name:        "dynamic_field_type",
					Description: `For ` + "`" + `DYNAMIC` + "`" + ` parameter types, the type of the field. Valid options are ` + "`" + `SOURCE` + "`" + `, ` + "`" + `SOURCE_TAG` + "`" + `, ` + "`" + `METRIC_NAME` + "`" + `, ` + "`" + `TAG_KEY` + "`" + `, and ` + "`" + `MATCHING_SOURCE_TAG` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_default_user_group",
			Category:         "Data Sources",
			ShortDescription: `Get the default user group ` + "`" + `Everyone` + "`" + ` from Wavefront`,
			Description: `

Use this data source to get the Group ID of the ` + "`" + `Everyone` + "`" + ` group in Wavefront. 

`,
			Keywords:  []string{},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `Set to the Group ID of the ` + "`" + `Everyone` + "`" + ` group, suitable for referencing in other resources that support group memberships.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_derived_metric",
			Category:         "Data Sources",
			ShortDescription: `Get the information about a certain Wavefront derived metric by its ID.`,
			Description: `

Use this data source to get information about a certain Wavefront derived metric by its ID.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID associated with the derived metric data to be fetched. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl #Get the information about a derived metric. data "wavefront_derived_metric" "example" { id = "derived_metric_id" } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the derived metric in Wavefront.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the derived metric in Wavefront.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `How frequently the query generating the derived metric is run.`,
				},
				resource.Attribute{
					Name:        "in_trash",
					Description: `A Boolean variable indicating trash status.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags assigned to the derived metric.`,
				},
				resource.Attribute{
					Name:        "query_failing",
					Description: `A Boolean variable indicating whether query is failing for the derived metric.`,
				},
				resource.Attribute{
					Name:        "last_error_message",
					Description: `Last error message occurred.`,
				},
				resource.Attribute{
					Name:        "last_failed_time",
					Description: `Timestamp of the last failed derived metric.`,
				},
				resource.Attribute{
					Name:        "additional_information",
					Description: `User-supplied additional explanatory information about the derived metric.`,
				},
				resource.Attribute{
					Name:        "create_user_id",
					Description: `The ID of the user who created the derived metric.`,
				},
				resource.Attribute{
					Name:        "update_user_id",
					Description: `The ID of the user who updated the derived metric.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the derived metric.`,
				},
				resource.Attribute{
					Name:        "hosts_used",
					Description: `A list of hosts used in the derived metric.`,
				},
				resource.Attribute{
					Name:        "last_processed_millis",
					Description: `The last processed timestamp.`,
				},
				resource.Attribute{
					Name:        "process_rate_minutes",
					Description: `The specified query is executed every ` + "`" + `process_rate_minutes` + "`" + ` minutes.`,
				},
				resource.Attribute{
					Name:        "points_scanned_at_last_query",
					Description: `The number of points scanned when last query was executed.`,
				},
				resource.Attribute{
					Name:        "last_query_time",
					Description: `The timestamp indicating the last time the query was executed.`,
				},
				resource.Attribute{
					Name:        "metrics_used",
					Description: `A list of metrics used in the derived metric.`,
				},
				resource.Attribute{
					Name:        "query_qb_enabled",
					Description: `A Boolean flag for enabling ` + "`" + `query_qb` + "`" + ``,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `A Boolean flag indicating whether the derived metric is deleted or not.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the derived metric is created.`,
				},
				resource.Attribute{
					Name:        "updated_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the derived metric is updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the derived metric in Wavefront.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the derived metric in Wavefront.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `How frequently the query generating the derived metric is run.`,
				},
				resource.Attribute{
					Name:        "in_trash",
					Description: `A Boolean variable indicating trash status.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags assigned to the derived metric.`,
				},
				resource.Attribute{
					Name:        "query_failing",
					Description: `A Boolean variable indicating whether query is failing for the derived metric.`,
				},
				resource.Attribute{
					Name:        "last_error_message",
					Description: `Last error message occurred.`,
				},
				resource.Attribute{
					Name:        "last_failed_time",
					Description: `Timestamp of the last failed derived metric.`,
				},
				resource.Attribute{
					Name:        "additional_information",
					Description: `User-supplied additional explanatory information about the derived metric.`,
				},
				resource.Attribute{
					Name:        "create_user_id",
					Description: `The ID of the user who created the derived metric.`,
				},
				resource.Attribute{
					Name:        "update_user_id",
					Description: `The ID of the user who updated the derived metric.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the derived metric.`,
				},
				resource.Attribute{
					Name:        "hosts_used",
					Description: `A list of hosts used in the derived metric.`,
				},
				resource.Attribute{
					Name:        "last_processed_millis",
					Description: `The last processed timestamp.`,
				},
				resource.Attribute{
					Name:        "process_rate_minutes",
					Description: `The specified query is executed every ` + "`" + `process_rate_minutes` + "`" + ` minutes.`,
				},
				resource.Attribute{
					Name:        "points_scanned_at_last_query",
					Description: `The number of points scanned when last query was executed.`,
				},
				resource.Attribute{
					Name:        "last_query_time",
					Description: `The timestamp indicating the last time the query was executed.`,
				},
				resource.Attribute{
					Name:        "metrics_used",
					Description: `A list of metrics used in the derived metric.`,
				},
				resource.Attribute{
					Name:        "query_qb_enabled",
					Description: `A Boolean flag for enabling ` + "`" + `query_qb` + "`" + ``,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `A Boolean flag indicating whether the derived metric is deleted or not.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the derived metric is created.`,
				},
				resource.Attribute{
					Name:        "updated_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the derived metric is updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_derived_metrics",
			Category:         "Data Sources",
			ShortDescription: `Get the information about all Wavefront derived metrics.`,
			Description: `
Use this data source to get information about all Wavefront derived metrics.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Limit is the maximum number of results to be returned. Defaults to 100.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Offset is the offset from the first result to be returned. Defaults to 0. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get the information about all derived metrics. data "wavefront_derived_metrics" "example" { limit = 10 offset = 0 } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "derived_metrics",
					Description: `List of all derived metrics in Wavefront. For each derived metric you will see a list of attributes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the derived metric in Wavefront.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the derived metric in Wavefront.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `How frequently the query generating the derived metric is run.`,
				},
				resource.Attribute{
					Name:        "in_trash",
					Description: `A Boolean variable indicating trash status.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags assigned to the derived metric.`,
				},
				resource.Attribute{
					Name:        "query_failing",
					Description: `A Boolean variable indicating whether query is failing for the derived metric.`,
				},
				resource.Attribute{
					Name:        "last_error_message",
					Description: `Last error message occurred.`,
				},
				resource.Attribute{
					Name:        "last_failed_time",
					Description: `Timestamp of the last failed derived metric.`,
				},
				resource.Attribute{
					Name:        "additional_information",
					Description: `User-supplied additional explanatory information about the derived metric.`,
				},
				resource.Attribute{
					Name:        "create_user_id",
					Description: `The ID of the user who created the derived metric.`,
				},
				resource.Attribute{
					Name:        "update_user_id",
					Description: `The ID of the user who updated the derived metric.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the derived metric.`,
				},
				resource.Attribute{
					Name:        "hosts_used",
					Description: `A list of hosts used in the derived metric.`,
				},
				resource.Attribute{
					Name:        "last_processed_millis",
					Description: `The last processed timestamp.`,
				},
				resource.Attribute{
					Name:        "process_rate_minutes",
					Description: `The specified query is executed every ` + "`" + `process_rate_minutes` + "`" + ` minutes.`,
				},
				resource.Attribute{
					Name:        "points_scanned_at_last_query",
					Description: `The number of points scanned when the last query was executed.`,
				},
				resource.Attribute{
					Name:        "include_obsolete_metrics",
					Description: `A Boolean flag indicating whether to include obsolete metrics or not.`,
				},
				resource.Attribute{
					Name:        "last_query_time",
					Description: `The timestamp indicating the last time the query was executed.`,
				},
				resource.Attribute{
					Name:        "query_qb_enabled",
					Description: `A Boolean flag for enabling ` + "`" + `query_qb` + "`" + ``,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `A Boolean flag indicating whether the derived metric is deleted or not.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the derived metric is created.`,
				},
				resource.Attribute{
					Name:        "updated_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the derived metric is updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "derived_metrics",
					Description: `List of all derived metrics in Wavefront. For each derived metric you will see a list of attributes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the derived metric in Wavefront.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the derived metric in Wavefront.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `How frequently the query generating the derived metric is run.`,
				},
				resource.Attribute{
					Name:        "in_trash",
					Description: `A Boolean variable indicating trash status.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags assigned to the derived metric.`,
				},
				resource.Attribute{
					Name:        "query_failing",
					Description: `A Boolean variable indicating whether query is failing for the derived metric.`,
				},
				resource.Attribute{
					Name:        "last_error_message",
					Description: `Last error message occurred.`,
				},
				resource.Attribute{
					Name:        "last_failed_time",
					Description: `Timestamp of the last failed derived metric.`,
				},
				resource.Attribute{
					Name:        "additional_information",
					Description: `User-supplied additional explanatory information about the derived metric.`,
				},
				resource.Attribute{
					Name:        "create_user_id",
					Description: `The ID of the user who created the derived metric.`,
				},
				resource.Attribute{
					Name:        "update_user_id",
					Description: `The ID of the user who updated the derived metric.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the derived metric.`,
				},
				resource.Attribute{
					Name:        "hosts_used",
					Description: `A list of hosts used in the derived metric.`,
				},
				resource.Attribute{
					Name:        "last_processed_millis",
					Description: `The last processed timestamp.`,
				},
				resource.Attribute{
					Name:        "process_rate_minutes",
					Description: `The specified query is executed every ` + "`" + `process_rate_minutes` + "`" + ` minutes.`,
				},
				resource.Attribute{
					Name:        "points_scanned_at_last_query",
					Description: `The number of points scanned when the last query was executed.`,
				},
				resource.Attribute{
					Name:        "include_obsolete_metrics",
					Description: `A Boolean flag indicating whether to include obsolete metrics or not.`,
				},
				resource.Attribute{
					Name:        "last_query_time",
					Description: `The timestamp indicating the last time the query was executed.`,
				},
				resource.Attribute{
					Name:        "query_qb_enabled",
					Description: `A Boolean flag for enabling ` + "`" + `query_qb` + "`" + ``,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `A Boolean flag indicating whether the derived metric is deleted or not.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the derived metric is created.`,
				},
				resource.Attribute{
					Name:        "updated_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the derived metric is updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_event",
			Category:         "Data Sources",
			ShortDescription: `Get the information about a certain Wavefront event.`,
			Description: `

Use this data source to get information about a certain Wavefront event.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID associated with the event data to be fetched. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get the information about a Wavefront event by its ID. data "wavefront_event" "example" { id = "sample-event-id" } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the event in Wavefront.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the event in Wavefront.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The end time of the event in epoch milliseconds.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `The severity category of the event.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the event.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `The description of the event.`,
				},
				resource.Attribute{
					Name:        "is_ephemeral",
					Description: `A Boolean flag. If set to ` + "`" + `true` + "`" + `, creates a point-in-time event (i.e. with no duration).`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `Annotations associated with the event.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags assigned to the event.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the event in Wavefront.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the event in Wavefront.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The end time of the event in epoch milliseconds.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `The severity category of the event.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the event.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `The description of the event.`,
				},
				resource.Attribute{
					Name:        "is_ephemeral",
					Description: `A Boolean flag. If set to ` + "`" + `true` + "`" + `, creates a point-in-time event (i.e. with no duration).`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `Annotations associated with the event.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags assigned to the event.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_events",
			Category:         "Data Sources",
			ShortDescription: `Get the information about all Wavefront events.`,
			Description: `

Use this data source to get information about all Wavefront events.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "latest_start_time_epoch_millis",
					Description: `(Required) The latest start time in epoch milliseconds.`,
				},
				resource.Attribute{
					Name:        "earliest_start_time_epoch_millis",
					Description: `(Required) The earliest start time in epoch milliseconds.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Limit is the maximum number of results to be returned. Defaults to 100.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Offset is the offset from the first result to be returned. Defaults to 0. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get the information about all events data "wavefront_events" "example" { limit = 10 offset = 0 latest_start_time_epoch_millis = 1665427195 earliest_start_time_epoch_millis = 1665427195 } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "events",
					Description: `List of all events in Wavefront. For each event you will see a list of attributes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the event in Wavefront.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the event in Wavefront.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The end time of the event in epoch milliseconds.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `The severity category of the event.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the event.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `The description of the event.`,
				},
				resource.Attribute{
					Name:        "is_ephemeral",
					Description: `A Boolean flag. If set to ` + "`" + `true` + "`" + `, creates a point-in-time event (i.e. with no duration).`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `Annotations associated with the event.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags assigned to the event.`,
				},
				resource.Attribute{
					Name:        "latest_start_time_epoch_millis",
					Description: `Latest start time in epoch milliseconds.`,
				},
				resource.Attribute{
					Name:        "earliest_start_time_epoch_millis",
					Description: `Earliest start time in epoch milliseconds.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "events",
					Description: `List of all events in Wavefront. For each event you will see a list of attributes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the event in Wavefront.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the event in Wavefront.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The end time of the event in epoch milliseconds.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `The severity category of the event.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the event.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `The description of the event.`,
				},
				resource.Attribute{
					Name:        "is_ephemeral",
					Description: `A Boolean flag. If set to ` + "`" + `true` + "`" + `, creates a point-in-time event (i.e. with no duration).`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `Annotations associated with the event.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tags assigned to the event.`,
				},
				resource.Attribute{
					Name:        "latest_start_time_epoch_millis",
					Description: `Latest start time in epoch milliseconds.`,
				},
				resource.Attribute{
					Name:        "earliest_start_time_epoch_millis",
					Description: `Earliest start time in epoch milliseconds.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_external_link",
			Category:         "Data Sources",
			ShortDescription: `Get the information about a specific Wavefront external link.`,
			Description: `

Use this data source to get information about a Wavefront external link by its ID.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the external link. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get the information about a specific external links. data "wavefront_external_link" "example" { id = "sample-external-link-id" } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the external link.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the external link.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description of this link.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `The mustache template for the link. The template must expand to a full URL, including scheme, origin, etc.`,
				},
				resource.Attribute{
					Name:        "metric_filter_regex",
					Description: `Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "source_filter_regex",
					Description: `Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "point_tag_filter_regexes",
					Description: `(Optional) Controls whether a link is displayed in the context menu of a highlighted series. This is a map from string to regular expression. The highlighted series must contain point tags whose keys are present in the keys of this map and whose values match the regular expressions associated with those keys in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "is_log_integration",
					Description: `Whether this is a "Log Integration" subType of external link.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the external link is created.`,
				},
				resource.Attribute{
					Name:        "updated_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the external link is updated.`,
				},
				resource.Attribute{
					Name:        "creator_id",
					Description: `The ID of the user who created the external link.`,
				},
				resource.Attribute{
					Name:        "updater_id",
					Description: `The ID of the user who updated the external link.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the external link.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the external link.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description of this link.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `The mustache template for the link. The template must expand to a full URL, including scheme, origin, etc.`,
				},
				resource.Attribute{
					Name:        "metric_filter_regex",
					Description: `Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "source_filter_regex",
					Description: `Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "point_tag_filter_regexes",
					Description: `(Optional) Controls whether a link is displayed in the context menu of a highlighted series. This is a map from string to regular expression. The highlighted series must contain point tags whose keys are present in the keys of this map and whose values match the regular expressions associated with those keys in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "is_log_integration",
					Description: `Whether this is a "Log Integration" subType of external link.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the external link is created.`,
				},
				resource.Attribute{
					Name:        "updated_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the external link is updated.`,
				},
				resource.Attribute{
					Name:        "creator_id",
					Description: `The ID of the user who created the external link.`,
				},
				resource.Attribute{
					Name:        "updater_id",
					Description: `The ID of the user who updated the external link.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_external_links",
			Category:         "Data Sources",
			ShortDescription: `Get the information about all Wavefront external links.`,
			Description: `

Use this data source to get information about all Wavefront external links.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Limit is the maximum number of results to be returned. Defaults to 100.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Offset is the offset from the first result to be returned. Defaults to 0. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get the information about all external links. data "wavefront_external_links" "example" { limit = 10 offset = 0 } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "external_links",
					Description: `List of all external links in Wavefront. For each external link you will see a list of attributes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the external link.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the external link.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description of the link.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `The mustache template for the link. The template must expand to a full URL, including scheme, origin, etc.`,
				},
				resource.Attribute{
					Name:        "metric_filter_regex",
					Description: `Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "source_filter_regex",
					Description: `Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "point_tag_filter_regexes",
					Description: `(Optional) Controls whether a link is displayed in the context menu of a highlighted series. This is a map from string to regular expression. The highlighted series must contain point tags whose keys are present in the keys of this map and whose values match the regular expressions associated with those keys in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "is_log_integration",
					Description: `Whether this is a "Log Integration" subType of external link.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the external link is created.`,
				},
				resource.Attribute{
					Name:        "updated_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the external link is updated.`,
				},
				resource.Attribute{
					Name:        "creator_id",
					Description: `The ID of the user who created the external link.`,
				},
				resource.Attribute{
					Name:        "updater_id",
					Description: `The ID of the user who updated the external link.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "external_links",
					Description: `List of all external links in Wavefront. For each external link you will see a list of attributes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the external link.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the external link.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description of the link.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `The mustache template for the link. The template must expand to a full URL, including scheme, origin, etc.`,
				},
				resource.Attribute{
					Name:        "metric_filter_regex",
					Description: `Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "source_filter_regex",
					Description: `Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "point_tag_filter_regexes",
					Description: `(Optional) Controls whether a link is displayed in the context menu of a highlighted series. This is a map from string to regular expression. The highlighted series must contain point tags whose keys are present in the keys of this map and whose values match the regular expressions associated with those keys in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "is_log_integration",
					Description: `Whether this is a "Log Integration" subType of external link.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the external link is created.`,
				},
				resource.Attribute{
					Name:        "updated_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the external link is updated.`,
				},
				resource.Attribute{
					Name:        "creator_id",
					Description: `The ID of the user who created the external link.`,
				},
				resource.Attribute{
					Name:        "updater_id",
					Description: `The ID of the user who updated the external link.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_maintenance_window",
			Category:         "Data Sources",
			ShortDescription: `Get the information about a specific Wavefront maintenance window.`,
			Description: `

Use this data source to get information about a Wavefront maintenance window by its ID.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the maintenance window. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get the information about specific maintenance window. data "wavefront_maintenance_window" "example" { id = "sample-maintenance-window-id" } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "event_name",
					Description: `The event name of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `The reason for the maintenance window.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The title of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "customer_id",
					Description: `The ID of the customer in Wavefront.`,
				},
				resource.Attribute{
					Name:        "running_state",
					Description: `The running state of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "start_time_in_seconds",
					Description: `The start time in seconds after 1 Jan 1970 GMT.`,
				},
				resource.Attribute{
					Name:        "end_time_in_seconds",
					Description: `The end time in seconds after 1 Jan 1970 GMT.`,
				},
				resource.Attribute{
					Name:        "relevant_customer_tags",
					Description: `The list of alert tags whose matching alerts will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_tags",
					Description: `The list of source or host tags whose matching sources or hosts will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_names",
					Description: `The list of source or host names that will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_tags_anded",
					Description: `Whether to AND source or host tags listed in ` + "`" + `relevant_host_tags` + "`" + `. If set to ` + "`" + `true` + "`" + `, the source or host must contain all tags for the maintenance window to apply. If set to ` + "`" + `false` + "`" + `, the tags are OR'ed, and the source or host must contain one of the tags. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host_tag_group_host_names_group_anded",
					Description: `If set to ` + "`" + `true` + "`" + `, the source or host must be in ` + "`" + `relevant_host_names` + "`" + ` and must have tags matching the specification formed by ` + "`" + `relevant_host_tags` + "`" + ` and ` + "`" + `relevant_host_tags_anded` + "`" + ` in for this maintenance window to apply. If set to false, the source or host must either be in ` + "`" + `relevant_host_names` + "`" + ` or match ` + "`" + `relevant_host_tags` + "`" + ` and ` + "`" + `relevant_host_tags_anded` + "`" + `. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the maintenance window is created.`,
				},
				resource.Attribute{
					Name:        "updated_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the maintenance window is updated.`,
				},
				resource.Attribute{
					Name:        "creator_id",
					Description: `The ID of the user who created the maintenance window.`,
				},
				resource.Attribute{
					Name:        "updater_id",
					Description: `The ID of the user who updated the maintenance window.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "event_name",
					Description: `The event name of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `The reason for the maintenance window.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The title of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "customer_id",
					Description: `The ID of the customer in Wavefront.`,
				},
				resource.Attribute{
					Name:        "running_state",
					Description: `The running state of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "start_time_in_seconds",
					Description: `The start time in seconds after 1 Jan 1970 GMT.`,
				},
				resource.Attribute{
					Name:        "end_time_in_seconds",
					Description: `The end time in seconds after 1 Jan 1970 GMT.`,
				},
				resource.Attribute{
					Name:        "relevant_customer_tags",
					Description: `The list of alert tags whose matching alerts will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_tags",
					Description: `The list of source or host tags whose matching sources or hosts will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_names",
					Description: `The list of source or host names that will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_tags_anded",
					Description: `Whether to AND source or host tags listed in ` + "`" + `relevant_host_tags` + "`" + `. If set to ` + "`" + `true` + "`" + `, the source or host must contain all tags for the maintenance window to apply. If set to ` + "`" + `false` + "`" + `, the tags are OR'ed, and the source or host must contain one of the tags. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host_tag_group_host_names_group_anded",
					Description: `If set to ` + "`" + `true` + "`" + `, the source or host must be in ` + "`" + `relevant_host_names` + "`" + ` and must have tags matching the specification formed by ` + "`" + `relevant_host_tags` + "`" + ` and ` + "`" + `relevant_host_tags_anded` + "`" + ` in for this maintenance window to apply. If set to false, the source or host must either be in ` + "`" + `relevant_host_names` + "`" + ` or match ` + "`" + `relevant_host_tags` + "`" + ` and ` + "`" + `relevant_host_tags_anded` + "`" + `. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the maintenance window is created.`,
				},
				resource.Attribute{
					Name:        "updated_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the maintenance window is updated.`,
				},
				resource.Attribute{
					Name:        "creator_id",
					Description: `The ID of the user who created the maintenance window.`,
				},
				resource.Attribute{
					Name:        "updater_id",
					Description: `The ID of the user who updated the maintenance window.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_maintenance_windows",
			Category:         "Data Sources",
			ShortDescription: `Get the information about all Wavefront maintenance windows.`,
			Description: `

Use this data source to get information about all Wavefront maintenance windows.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Limit is the maximum number of results to be returned. Defaults to 100.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Offset is the offset from the first result to be returned. Defaults to 0. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get the information about all maintenance windows. data "wavefront_maintenance_window_all" "example" { limit = 10 offset = 0 } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "maintenance_windows",
					Description: `List of all maintenance windows in Wavefront. For each maintenance window you will see a list of attributes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "event_name",
					Description: `The event name of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `The reason for the maintenance window.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The title of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "customer_id",
					Description: `The ID of the customer in Wavefront.`,
				},
				resource.Attribute{
					Name:        "running_state",
					Description: `The running state of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "start_time_in_seconds",
					Description: `The start time in seconds after 1 Jan 1970 GMT.`,
				},
				resource.Attribute{
					Name:        "end_time_in_seconds",
					Description: `The end time in seconds after 1 Jan 1970 GMT.`,
				},
				resource.Attribute{
					Name:        "relevant_customer_tags",
					Description: `The list of alert tags whose matching alerts will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_tags",
					Description: `The list of source or host tags whose matching sources or hosts will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_names",
					Description: `The list of source or host names that will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_tags_anded",
					Description: `Whether to AND source or host tags listed in ` + "`" + `relevant_host_tags` + "`" + `. If set to ` + "`" + `true` + "`" + `, the source or host must contain all tags for the maintenance window to apply. If set to ` + "`" + `false` + "`" + `, the tags are OR'ed, and the source or host must contain one of the tags. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host_tag_group_host_names_group_anded",
					Description: `If set to ` + "`" + `true` + "`" + `, the source or host must be in ` + "`" + `relevant_host_names` + "`" + ` and must have tags matching the specification formed by ` + "`" + `relevant_host_tags` + "`" + ` and ` + "`" + `relevant_host_tags_anded` + "`" + ` in for this maintenance window to apply. If set to false, the source or host must either be in ` + "`" + `relevant_host_names` + "`" + ` or match ` + "`" + `relevant_host_tags` + "`" + ` and ` + "`" + `relevant_host_tags_anded` + "`" + `. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the maintenance window is created.`,
				},
				resource.Attribute{
					Name:        "updated_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the maintenance window is updated.`,
				},
				resource.Attribute{
					Name:        "creator_id",
					Description: `The ID of the user who created the maintenance window.`,
				},
				resource.Attribute{
					Name:        "updater_id",
					Description: `The ID of the user who updated the maintenance window.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "maintenance_windows",
					Description: `List of all maintenance windows in Wavefront. For each maintenance window you will see a list of attributes.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "event_name",
					Description: `The event name of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `The reason for the maintenance window.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The title of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "customer_id",
					Description: `The ID of the customer in Wavefront.`,
				},
				resource.Attribute{
					Name:        "running_state",
					Description: `The running state of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "start_time_in_seconds",
					Description: `The start time in seconds after 1 Jan 1970 GMT.`,
				},
				resource.Attribute{
					Name:        "end_time_in_seconds",
					Description: `The end time in seconds after 1 Jan 1970 GMT.`,
				},
				resource.Attribute{
					Name:        "relevant_customer_tags",
					Description: `The list of alert tags whose matching alerts will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_tags",
					Description: `The list of source or host tags whose matching sources or hosts will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_names",
					Description: `The list of source or host names that will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_tags_anded",
					Description: `Whether to AND source or host tags listed in ` + "`" + `relevant_host_tags` + "`" + `. If set to ` + "`" + `true` + "`" + `, the source or host must contain all tags for the maintenance window to apply. If set to ` + "`" + `false` + "`" + `, the tags are OR'ed, and the source or host must contain one of the tags. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host_tag_group_host_names_group_anded",
					Description: `If set to ` + "`" + `true` + "`" + `, the source or host must be in ` + "`" + `relevant_host_names` + "`" + ` and must have tags matching the specification formed by ` + "`" + `relevant_host_tags` + "`" + ` and ` + "`" + `relevant_host_tags_anded` + "`" + ` in for this maintenance window to apply. If set to false, the source or host must either be in ` + "`" + `relevant_host_names` + "`" + ` or match ` + "`" + `relevant_host_tags` + "`" + ` and ` + "`" + `relevant_host_tags_anded` + "`" + `. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the maintenance window is created.`,
				},
				resource.Attribute{
					Name:        "updated_epoch_millis",
					Description: `The timestamp in epoch milliseconds indicating when the maintenance window is updated.`,
				},
				resource.Attribute{
					Name:        "creator_id",
					Description: `The ID of the user who created the maintenance window.`,
				},
				resource.Attribute{
					Name:        "updater_id",
					Description: `The ID of the user who updated the maintenance window.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_role",
			Category:         "Data Sources",
			ShortDescription: `Get the information about a specific Wavefront role.`,
			Description: `

Use this data source to get information about a Wavefront role by its ID.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID associated with the role data to be fetched. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get the information about the role. data "wavefront_role" "example" { id = "role-id" } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the role in Wavefront.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the role in Wavefront.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description of the role.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `The list of permissions associated with role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the role in Wavefront.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the role in Wavefront.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description of the role.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `The list of permissions associated with role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_roles",
			Category:         "Data Sources",
			ShortDescription: `Get all Roles from Wavefront`,
			Description: `

Use this data source to get all Roles in Wavefront. 

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Limit is the maximum number of results to be returned. Defaults to 100.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Offset is the offset from the first result to be returned. Defaults to 0. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get all Roles data "wavefront_roles" "roles" { limit = 10 offset = 0 } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of Wavefront Roles.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Role ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Role Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The Role's description.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `List of Permissions (Strings) associated with Role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "roles",
					Description: `List of Wavefront Roles.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Role ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The Role Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The Role's description.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `List of Permissions (Strings) associated with Role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_user",
			Category:         "Data Sources",
			ShortDescription: `Get the information for a given user from Wavefront`,
			Description: `

Use this data source to get information for a given user by email from Wavefront. 

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `The email associated with the user data to be fetched. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get the info for user "example.user@example.com" data "wavefront_user" "example" { email = "example.user@example.com" } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `List of permissions granted to a user.`,
				},
				resource.Attribute{
					Name:        "user_group_ids",
					Description: `List of User Group Ids the user is a member of.`,
				},
				resource.Attribute{
					Name:        "last_successful_login",
					Description: `When the user last logged in to Wavefront.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "permissions",
					Description: `List of permissions granted to a user.`,
				},
				resource.Attribute{
					Name:        "user_group_ids",
					Description: `List of User Group Ids the user is a member of.`,
				},
				resource.Attribute{
					Name:        "last_successful_login",
					Description: `When the user last logged in to Wavefront.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_user_group",
			Category:         "Data Sources",
			ShortDescription: `Get the information about a specific Wavefront user group.`,
			Description: `

Use this data source to get information about a Wavefront user group by its ID.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID associated with the user group data to be fetched. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get the information about the user group. data "wavefront_user_group" "example" { id = "user-group-id" } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the group in Wavefront.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the group in Wavefront.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description of the group.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `The list of roles associated with the group.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `The list of users assigned to the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the group in Wavefront.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the group in Wavefront.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description of the group.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `The list of roles associated with the group.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `The list of users assigned to the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_user_groups",
			Category:         "Data Sources",
			ShortDescription: `Get all User Groups from Wavefront`,
			Description: `

Use this data source to get all User Groups in Wavefront. 

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Limit is the maximum number of results to be returned. Defaults to 100.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Offset is the offset from the first result to be returned. Defaults to 0. ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Get all user groups data "wavefront_user_groups" "groups" { limit = 10 offset = 0 } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "user_groups",
					Description: `List of user groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The group name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The group description.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of roles associated with the group.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `List of users assigned to the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_groups",
					Description: `List of user groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The group name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The group description.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of roles associated with the group.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `List of users assigned to the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_users",
			Category:         "Data Sources",
			ShortDescription: `Get all users from Wavefront`,
			Description: `

Use this data source to get all users in Wavefront. 



`,
			Keywords:  []string{},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "users",
					Description: `List of all users in Wavefront.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `List of permissions granted to a user.`,
				},
				resource.Attribute{
					Name:        "user_group_ids",
					Description: `List of User Group Ids the user is a member of.`,
				},
				resource.Attribute{
					Name:        "last_successful_login",
					Description: `When the user last logged in to Wavefront.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"wavefront_alert":               0,
		"wavefront_alerts":              1,
		"wavefront_dashboard":           2,
		"wavefront_dashboards":          3,
		"wavefront_default_user_group":  4,
		"wavefront_derived_metric":      5,
		"wavefront_derived_metrics":     6,
		"wavefront_event":               7,
		"wavefront_events":              8,
		"wavefront_external_link":       9,
		"wavefront_external_links":      10,
		"wavefront_maintenance_window":  11,
		"wavefront_maintenance_windows": 12,
		"wavefront_role":                13,
		"wavefront_roles":               14,
		"wavefront_user":                15,
		"wavefront_user_group":          16,
		"wavefront_user_groups":         17,
		"wavefront_users":               18,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
