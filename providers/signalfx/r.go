package signalfx

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "signalfx_dashboard",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx Dashboards`,
			Description: `

A dashboard is a curated collection of specific charts and supports dimensional [filters](http://docs.signalfx.com/en/latest/dashboards/dashboard-filter-dynamic.html#filter-dashboard-charts), [dashboard variables](http://docs.signalfx.com/en/latest/dashboards/dashboard-filter-dynamic.html#dashboard-variables) and [time range](http://docs.signalfx.com/en/latest/_sidebars-and-includes/using-time-range-selector.html#time-range-selector) options. These options are applied to all charts in the dashboard, providing a consistent view of the data displayed in that dashboard. This also means that when you open a chart to drill down for more details, you are viewing the same data that is visible in the dashboard view.

**NOTE:** Since every dashboard is included in a [dashboard group](dashboard_group.html) (SignalFx collection of dashboards), you need to create that first and reference it as shown in the example.

`,
			Keywords: []string{
				"dashboard",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the dashboard.`,
				},
				resource.Attribute{
					Name:        "dashboard_group",
					Description: `(Required) The ID of the dashboard group that contains the dashboard.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the dashboard.`,
				},
				resource.Attribute{
					Name:        "charts_resolution",
					Description: `(Optional) Specifies the chart data display resolution for charts in this dashboard. Value can be one of ` + "`" + `"default"` + "`" + `, ` + "`" + `"low"` + "`" + `, ` + "`" + `"high"` + "`" + `, or ` + "`" + `"highest"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "time_range",
					Description: `(Optional) The time range prior to now to visualize. SignalFx time syntax (e.g. ` + "`" + `"-5m"` + "`" + `, ` + "`" + `"-1h"` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) Seconds since epoch. Used for visualization. You must specify time_span_type = ` + "`" + `"absolute"` + "`" + ` too.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) Seconds since epoch. Used for visualization. You must specify time_span_type = ` + "`" + `"absolute"` + "`" + ` too.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter to apply to the charts when displaying the dashboard.`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `(Required) A metric time series dimension or property name.`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) Whether this filter should be a not filter. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) List of of strings (which will be treated as an OR filter on the property).`,
				},
				resource.Attribute{
					Name:        "apply_if_exist",
					Description: `(Optional) If true, this filter will also match data that doesn't have this property at all.`,
				},
				resource.Attribute{
					Name:        "variable",
					Description: `(Optional) Dashboard variable to apply to each chart in the dashboard.`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `(Required) A metric time series dimension or property name.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `(Required) An alias for the dashboard variable. This text will appear as the label for the dropdown field on the dashboard.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Variable description.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) List of of strings (which will be treated as an OR filter on the property).`,
				},
				resource.Attribute{
					Name:        "value_required",
					Description: `(Optional) Determines whether a value is required for this variable (and therefore whether it will be possible to view this dashboard without this filter applied). ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "values_suggested",
					Description: `(Optional) A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.`,
				},
				resource.Attribute{
					Name:        "restricted_suggestions",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, this variable may only be set to the values listed in ` + "`" + `values_suggested` + "`" + ` and only these values will appear in autosuggestion menus. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "replace_only",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, this variable will only apply to charts that have a filter for the property.`,
				},
				resource.Attribute{
					Name:        "apply_if_exist",
					Description: `(Optional) If true, this variable will also match data that doesn't have this property at all.`,
				},
				resource.Attribute{
					Name:        "chart",
					Description: `(Optional) Chart ID and layout information for the charts in the dashboard.`,
				},
				resource.Attribute{
					Name:        "chart_id",
					Description: `(Required) ID of the chart to display.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `(Optional) How many columns (out of a total of 12) the chart should take up (between ` + "`" + `1` + "`" + ` and ` + "`" + `12` + "`" + `). ` + "`" + `12` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `(Optional) How many rows the chart should take up (greater than or equal to ` + "`" + `1` + "`" + `). ` + "`" + `1` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "row",
					Description: `(Optional) The row to show the chart in (zero-based); if ` + "`" + `height > 1` + "`" + `, this value represents the topmost row of the chart (greater than or equal to ` + "`" + `0` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "column",
					Description: `(Optional) The column to show the chart in (zero-based); this value always represents the leftmost column of the chart (between ` + "`" + `0` + "`" + ` and ` + "`" + `11` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "grid",
					Description: `(Optional) Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.`,
				},
				resource.Attribute{
					Name:        "chart_ids",
					Description: `(Required) List of IDs of the charts to display.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `(Optional) How many columns (out of a total of 12) every chart should take up (between ` + "`" + `1` + "`" + ` and ` + "`" + `12` + "`" + `). ` + "`" + `12` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `(Optional) How many rows every chart should take up (greater than or equal to ` + "`" + `1` + "`" + `). ` + "`" + `1` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "column",
					Description: `(Optional) Column layout. Charts listed will be placed in a single column with the same width and height.`,
				},
				resource.Attribute{
					Name:        "chart_ids",
					Description: `(Required) List of IDs of the charts to display.`,
				},
				resource.Attribute{
					Name:        "column",
					Description: `(Optional) Column number for the layout.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `(Optional) How many columns (out of a total of ` + "`" + `12` + "`" + `) every chart should take up (between ` + "`" + `1` + "`" + ` and ` + "`" + `12` + "`" + `). ` + "`" + `12` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "height",
					Description: `(Optional) How many rows every chart should take up (greater than or equal to 1). 1 by default.`,
				},
				resource.Attribute{
					Name:        "event_overlay",
					Description: `(Optional) Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the`,
				},
				resource.Attribute{
					Name:        "line",
					Description: `(Optional) Show a vertical line for the event. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) Text shown in the dropdown when selecting this overlay from the menu.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.`,
				},
				resource.Attribute{
					Name:        "signal",
					Description: `Search term used to choose the events shown in the overlay.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Can be set to ` + "`" + `eventTimeSeries` + "`" + ` (the default) to refer to externally reported events, or ` + "`" + `detectorEvents` + "`" + ` to refer to events from detector triggers.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Each element specifies a filter to use against the signal specified in the ` + "`" + `signal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `The name of a dimension to filter against.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `A list of values to be used with the ` + "`" + `property` + "`" + `, they will be combined via ` + "`" + `OR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) If true, only data that does not match the specified value of the specified property appear in the event overlay. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "selected_event_overlay",
					Description: `(Optional) Defines event overlays which are enabled by`,
				},
				resource.Attribute{
					Name:        "signal",
					Description: `Search term used to choose the events shown in the overlay.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Can be set to ` + "`" + `eventTimeSeries` + "`" + ` (the default) to refer to externally reported events, or ` + "`" + `detectorEvents` + "`" + ` to refer to events from detector triggers.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Each element specifies a filter to use against the signal specified in the ` + "`" + `signal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `The name of a dimension to filter against.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `A list of values to be used with the ` + "`" + `property` + "`" + `, they will be combined via ` + "`" + `OR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) If true, only data that does not match the specified value of the specified property appear in the event overlay. Defaults to ` + "`" + `false` + "`" + `. ## Dashboard Layout Information`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_dashboard_group",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx Dashboard Groups`,
			Description: `

In the SignalFx web UI, a [dashboard group](https://developers.signalfx.com/v2/docs/dashboard-group-model) is a collection of dashboards.

**NOTE:** Dashboard groups cannot be accessed directly, but just via a dashboard contained in them. This is the reason why make show won't show any of yours dashboard groups.

`,
			Keywords: []string{
				"dashboard",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the dashboard group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the dashboard group.`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `(Optional) Team IDs to associate the dashboard group to.`,
				},
				resource.Attribute{
					Name:        "dashboard",
					Description: `(Optional) [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group.`,
				},
				resource.Attribute{
					Name:        "dashboard_id",
					Description: `(Required) The dashboard id to mirror`,
				},
				resource.Attribute{
					Name:        "name_override",
					Description: `(Optional) The name that will override the original dashboards's name.`,
				},
				resource.Attribute{
					Name:        "description_override",
					Description: `(Optional) The description that will override the original dashboards's description.`,
				},
				resource.Attribute{
					Name:        "filter_override",
					Description: `(Optional) The description that will override the original dashboards's description.`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `(Required) The name of a dimension to filter against.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to be used with the ` + "`" + `property` + "`" + `, they will be combined via ` + "`" + `OR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) If true, only data that does not match the specified value of the specified property appear in the event overlay. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter_override",
					Description: `(Optional) The description that will override the original dashboards's description.`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `(Required) A metric time series dimension or property name.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) (Optional) List of of strings (which will be treated as an OR filter on the property).`,
				},
				resource.Attribute{
					Name:        "values_suggested",
					Description: `(Optional) A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_detector",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx Dashboards`,
			Description: `

Provides a SignalFx detector resource. This can be used to create and manage detectors.

`,
			Keywords: []string{
				"detector",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the detector.`,
				},
				resource.Attribute{
					Name:        "program_text",
					Description: `(Required) Signalflow program text for the detector. More info at <https://developers.signalfx.com/docs/signalflow-overview>.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the detector.`,
				},
				resource.Attribute{
					Name:        "max_delay",
					Description: `(Optional) How long (in seconds) to wait for late datapoints. See <https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints> for more info. Max value is ` + "`" + `900` + "`" + ` seconds (15 minutes).`,
				},
				resource.Attribute{
					Name:        "show_data_markers",
					Description: `(Optional) When ` + "`" + `true` + "`" + `, markers will be drawn for each datapoint within the visualization. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "show_event_lines",
					Description: `(Optional) When ` + "`" + `true` + "`" + `, the visualization will display a vertical line for each event trigger. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "disable_sampling",
					Description: `(Optional) When ` + "`" + `false` + "`" + `, the visualization may sample the output timeseries rather than displaying them all. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "time_range",
					Description: `(Optional) From when to display data. SignalFx time syntax (e.g. ` + "`" + `"-5m"` + "`" + `, ` + "`" + `"-1h"` + "`" + `). Conflicts with ` + "`" + `start_time` + "`" + ` and ` + "`" + `end_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) Seconds since epoch. Used for visualization. Conflicts with ` + "`" + `time_range` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) Seconds since epoch. Used for visualization. Conflicts with ` + "`" + `time_range` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `(Optional) Team IDs to associcate the detector to.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) Set of rules used for alerting.`,
				},
				resource.Attribute{
					Name:        "detect_label",
					Description: `(Required) A detect label which matches a detect label within ` + "`" + `program_text` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Required) The severity of the rule, must be one of: ` + "`" + `"Critical"` + "`" + `, ` + "`" + `"Major"` + "`" + `, ` + "`" + `"Minor"` + "`" + `, ` + "`" + `"Warning"` + "`" + `, ` + "`" + `"Info"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) When true, notifications and events will not be generated for the detect label. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `(Optional) List of strings specifying where notifications will be sent when an incident occurs. See <https://developers.signalfx.com/detectors_reference.html#operation/Create%20Single%20Detector> for more info.`,
				},
				resource.Attribute{
					Name:        "parameterized_body",
					Description: `(Optional) Custom notification message body when an alert is triggered. See <https://docs.signalfx.com/en/latest/detect-alert/set-up-detectors.html#about-detectors#alert-settings> for more info.`,
				},
				resource.Attribute{
					Name:        "parameterized_subject",
					Description: `(Optional) Custom notification message subject when an alert is triggered. See <https://docs.signalfx.com/en/latest/detect-alert/set-up-detectors.html#about-detectors#alert-settings> for more info.`,
				},
				resource.Attribute{
					Name:        "runbook_url",
					Description: `(Optional) URL of page to consult when an alert is triggered. This can be used with custom notification messages.`,
				},
				resource.Attribute{
					Name:        "tip",
					Description: `(Optional) Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SignalFx detector ## Import Downtimes can be imported using their string ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import signalfx_detector.application_delay abc123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SignalFx detector ## Import Downtimes can be imported using their string ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import signalfx_detector.application_delay abc123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_event_feed_chart",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx time charts`,
			Description: `

Displays a listing of events as a widget in a dashboard.

`,
			Keywords: []string{
				"event",
				"feed",
				"chart",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the text note.`,
				},
				resource.Attribute{
					Name:        "program_text",
					Description: `(Required) Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the text note.`,
				},
				resource.Attribute{
					Name:        "time_range",
					Description: `(Optional) From when to display data. SignalFx time syntax (e.g. ` + "`" + `"-5m"` + "`" + `, ` + "`" + `"-1h"` + "`" + `). Conflicts with ` + "`" + `start_time` + "`" + ` and ` + "`" + `end_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) Seconds since epoch. Used for visualization. Conflicts with ` + "`" + `time_range` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) Seconds since epoch. Used for visualization. Conflicts with ` + "`" + `time_range` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_heatmap_chart",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx heat map chart`,
			Description: `

This chart type displays the specified plot in a heatmap fashion. This format is similar to the [Infrastructure Navigator](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/built-in-content/infra-nav.html#infra), with squares representing each source for the selected metric, and the color of each square representing the value range of the metric.

`,
			Keywords: []string{
				"heatmap",
				"chart",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the chart.`,
				},
				resource.Attribute{
					Name:        "program_text",
					Description: `(Required) Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the chart.`,
				},
				resource.Attribute{
					Name:        "unit_prefix",
					Description: `(Optional) Must be ` + "`" + `"Metric"` + "`" + ` or ` + "`" + `"Binary` + "`" + `". ` + "`" + `"Metric"` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "minimum_resolution",
					Description: `(Optional) The minimum resolution (in seconds) to use for computing the underlying program.`,
				},
				resource.Attribute{
					Name:        "max_delay",
					Description: `(Optional) How long (in seconds) to wait for late datapoints.`,
				},
				resource.Attribute{
					Name:        "refresh_interval",
					Description: `(Optional) How often (in seconds) to refresh the values of the heatmap.`,
				},
				resource.Attribute{
					Name:        "disable_sampling",
					Description: `(Optional) If ` + "`" + `false` + "`" + `, samples a subset of the output MTS, which improves UI performance. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "group_by",
					Description: `(Optional) Properties to group by in the heatmap (in nesting order).`,
				},
				resource.Attribute{
					Name:        "sort_by",
					Description: `(Optional) The property to use when sorting the elements. Must be prepended with ` + "`" + `+` + "`" + ` for ascending or ` + "`" + `-` + "`" + ` for descending (e.g. ` + "`" + `-foo` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "hide_timestamp",
					Description: `(Optional) Whether to show the timestamp in the chart. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "color_range",
					Description: `(Optional. Conflict with color_scale) Values and color for the color range. Example: ` + "`" + `color_range : { min : 0, max : 100, color : blue }` + "`" + `. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).`,
				},
				resource.Attribute{
					Name:        "min_value",
					Description: `(Optional) The minimum value within the coloring range.`,
				},
				resource.Attribute{
					Name:        "max_value",
					Description: `(Optional) The maximum value within the coloring range.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Required) The color range to use. Must be either gray, blue, navy, orange, yellow, magenta, purple, violet, lilac, green, aquamarine.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_integration",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx Integrations`,
			Description: `

SignalFx supports integrations to ingest metrics from other monitoring systems, connect to Single Sign-On providers, and to report notifications for messaging and incident management. Note that your API key must have admin permissions to use the SignalFx integration API.

**Note:** This resource is deprecated.

In a future release ` + "`" + `signalfx_integration` + "`" + ` will be replaced with specific resources for each integration. Please see the specific ` + "`" + `signalfx_integration_*` + "`" + ` resources in the sidebar.


`,
			Keywords: []string{
				"integration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the integration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Whether the integration is enabled.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of the integration. See [the full list here](https://developers.signalfx.com/integrations_reference.html).`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `(Required for ` + "`" + `PagerDuty` + "`" + `) PagerDuty API key.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `(Required for ` + "`" + `Slack` + "`" + `) Slack incoming webhook URL.`,
				},
				resource.Attribute{
					Name:        "poll_rate",
					Description: `(Required for ` + "`" + `GCP` + "`" + `) GCP integration poll rate in milliseconds. Can be set to either 60000 or 300000 (1 minute or 5 minutes).`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional for ` + "`" + `GCP` + "`" + `) GCP service metrics to import. Can be an empty list, or not included, to import 'All services'.`,
				},
				resource.Attribute{
					Name:        "project_service_keys",
					Description: `(Required for ` + "`" + `GCP` + "`" + `) GCP projects to add.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_integration_gcp",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx GCP Integrations`,
			Description: `

SignalFx GCP Integration

`,
			Keywords: []string{
				"integration",
				"gcp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the integration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Whether the integration is enabled.`,
				},
				resource.Attribute{
					Name:        "poll_rate",
					Description: `(Required) GCP integration poll rate in milliseconds. Can be set to either 60000 or 300000 (1 minute or 5 minutes).`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional) GCP service metrics to import. Can be an empty list, or not included, to import 'All services'.`,
				},
				resource.Attribute{
					Name:        "project_service_keys",
					Description: `(Required) GCP projects to add.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_integration_pagerduty",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx PagerDuty Integrations`,
			Description: `

SignalFx PagerDuty integrations

`,
			Keywords: []string{
				"integration",
				"pagerduty",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the integration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Whether the integration is enabled.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `(Required) PagerDuty API key.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_integration_slack",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx Slack Integrations`,
			Description: `

SignalFx Slack integration.

`,
			Keywords: []string{
				"integration",
				"slack",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the integration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Whether the integration is enabled.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `(Required) Slack incoming webhook URL.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_list_chart",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx list charts`,
			Description: `

This chart type displays current data values in a list format.

The name of each value in the chart reflects the name of the plot and any associated dimensions. We recommend you click the Pencil icon and give the plot a meaningful name, as in plot B below. Otherwise, just the raw metric name will be displayed on the chart, as in plot A below.


`,
			Keywords: []string{
				"list",
				"chart",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the chart.`,
				},
				resource.Attribute{
					Name:        "program_text",
					Description: `(Required) Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the chart.`,
				},
				resource.Attribute{
					Name:        "unit_prefix",
					Description: `(Optional) Must be ` + "`" + `"Metric"` + "`" + ` or ` + "`" + `"Binary` + "`" + `". ` + "`" + `"Metric"` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "color_by",
					Description: `(Optional) Must be ` + "`" + `"Dimension"` + "`" + ` or ` + "`" + `"Metric"` + "`" + `. ` + "`" + `"Dimension"` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "max_delay",
					Description: `(Optional) How long (in seconds) to wait for late datapoints.`,
				},
				resource.Attribute{
					Name:        "disable_sampling",
					Description: `(Optional) If ` + "`" + `false` + "`" + `, samples a subset of the output MTS, which improves UI performance. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "refresh_interval",
					Description: `(Optional) How often (in seconds) to refresh the values of the list.`,
				},
				resource.Attribute{
					Name:        "viz_options",
					Description: `(Optional) Plot-level customization options, associated with a publish statement.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Label used in the publish statement that displays the plot (metric time series data) you want to customize.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.`,
				},
				resource.Attribute{
					Name:        "value_unit",
					Description: `(Optional) A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are ` + "`" + `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "legend_fields_to_hide",
					Description: `(Optional) List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use ` + "`" + `legend_options_fields` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "legend_options_fields",
					Description: `(Optional) List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with ` + "`" + `legend_fields_to_hide` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_precision",
					Description: `(Optional) Maximum number of digits to display when rounding values up or down.`,
				},
				resource.Attribute{
					Name:        "secondary_visualization",
					Description: `(Optional) The type of secondary visualization. Can be ` + "`" + `None` + "`" + `, ` + "`" + `Radial` + "`" + `, ` + "`" + `Linear` + "`" + `, or ` + "`" + `Sparkline` + "`" + `. If unset, the SignalFx default is used (` + "`" + `Sparkline` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "color_scale",
					Description: `(Optional. ` + "`" + `color_by` + "`" + ` must be ` + "`" + `"Scale"` + "`" + `) Single color range including both the color to display for that range and the borders of the range. Example: ` + "`" + `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]` + "`" + `. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).`,
				},
				resource.Attribute{
					Name:        "gt",
					Description: `(Optional) Indicates the lower threshold non-inclusive value for this range.`,
				},
				resource.Attribute{
					Name:        "gte",
					Description: `(Optional) Indicates the lower threshold inclusive value for this range.`,
				},
				resource.Attribute{
					Name:        "lt",
					Description: `(Optional) Indicates the upper threshold non-inculsive value for this range.`,
				},
				resource.Attribute{
					Name:        "lte",
					Description: `(Optional) Indicates the upper threshold inclusive value for this range.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Required) The color range to use. Must be either gray, blue, navy, orange, yellow, magenta, purple, violet, lilac, green, aquamarine.`,
				},
				resource.Attribute{
					Name:        "sort_by",
					Description: `(Optional) The property to use when sorting the elements. Use ` + "`" + `value` + "`" + ` if you want to sort by value. Must be prepended with ` + "`" + `+` + "`" + ` for ascending or ` + "`" + `-` + "`" + ` for descending (e.g. ` + "`" + `-foo` + "`" + `). Note there are some special values for some of the options provided in the UX: ` + "`" + `"value"` + "`" + ` for Value, ` + "`" + `"sf_originatingMetric"` + "`" + ` for Metric, and ` + "`" + `"sf_metric"` + "`" + ` for plot.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_single_value_chart",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx single value charts`,
			Description: `

This chart type displays a single number in a large font, representing the current value of a single metric on a plot line.

If the time period is in the past, the number represents the value of the metric near the end of the time period.

`,
			Keywords: []string{
				"single",
				"value",
				"chart",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the chart.`,
				},
				resource.Attribute{
					Name:        "program_text",
					Description: `(Required) Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the chart.`,
				},
				resource.Attribute{
					Name:        "color_by",
					Description: `(Optional) Must be ` + "`" + `"Dimension"` + "`" + `, ` + "`" + `"Scale"` + "`" + ` or ` + "`" + `"Metric"` + "`" + `. ` + "`" + `"Dimension"` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "color_scale",
					Description: `(Optional. ` + "`" + `color_by` + "`" + ` must be ` + "`" + `"Scale"` + "`" + `) Single color range including both the color to display for that range and the borders of the range. Example: ` + "`" + `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]` + "`" + `. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).`,
				},
				resource.Attribute{
					Name:        "gt",
					Description: `(Optional) Indicates the lower threshold non-inclusive value for this range.`,
				},
				resource.Attribute{
					Name:        "gte",
					Description: `(Optional) Indicates the lower threshold inclusive value for this range.`,
				},
				resource.Attribute{
					Name:        "lt",
					Description: `(Optional) Indicates the upper threshold non-inculsive value for this range.`,
				},
				resource.Attribute{
					Name:        "lte",
					Description: `(Optional) Indicates the upper threshold inclusive value for this range.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Required) The color range to use. Must be either gray, blue, navy, orange, yellow, magenta, purple, violet, lilac, green, aquamarine.`,
				},
				resource.Attribute{
					Name:        "viz_options",
					Description: `(Optional) Plot-level customization options, associated with a publish statement.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Label used in the publish statement that displays the plot (metric time series data) you want to customize.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.`,
				},
				resource.Attribute{
					Name:        "value_unit",
					Description: `(Optional) A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are ` + "`" + `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "unit_prefix",
					Description: `(Optional) Must be ` + "`" + `"Metric"` + "`" + ` or ` + "`" + `"Binary"` + "`" + `. ` + "`" + `"Metric"` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "max_delay",
					Description: `(Optional) How long (in seconds) to wait for late datapoints`,
				},
				resource.Attribute{
					Name:        "refresh_interval",
					Description: `(Optional) How often (in seconds) to refresh the value.`,
				},
				resource.Attribute{
					Name:        "max_precision",
					Description: `(Optional) The maximum precision to for value displayed.`,
				},
				resource.Attribute{
					Name:        "is_timestamp_hidden",
					Description: `(Optional) Whether to hide the timestamp in the chart. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "secondary_visualization",
					Description: `(Optional) The type of secondary visualization. Can be ` + "`" + `None` + "`" + `, ` + "`" + `Radial` + "`" + `, ` + "`" + `Linear` + "`" + `, or ` + "`" + `Sparkline` + "`" + `. If unset, the SignalFx default is used (` + "`" + `None` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "show_spark_line",
					Description: `(Optional) Whether to show a trend line below the current value. ` + "`" + `false` + "`" + ` by default.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_text_chart",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx text notes`,
			Description: `

This special type of chart doesn’t display any metric data. Rather, it lets you place a text note on the dashboard.

`,
			Keywords: []string{
				"text",
				"chart",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the text note.`,
				},
				resource.Attribute{
					Name:        "markdown",
					Description: `(Required) Markdown text to display.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the text note.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_time_chart",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx time charts`,
			Description: `

Provides a SignalFx time chart resource. This can be used to create and manage the different types of time charts.

Time charts display data points over a period of time.

`,
			Keywords: []string{
				"time",
				"chart",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the chart.`,
				},
				resource.Attribute{
					Name:        "program_text",
					Description: `(Required) Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.`,
				},
				resource.Attribute{
					Name:        "plot_type",
					Description: `(Optional) The default plot display style for the visualization. Must be ` + "`" + `"LineChart"` + "`" + `, ` + "`" + `"AreaChart"` + "`" + `, ` + "`" + `"ColumnChart"` + "`" + `, or ` + "`" + `"Histogram"` + "`" + `. Default: ` + "`" + `"LineChart"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the chart.`,
				},
				resource.Attribute{
					Name:        "axes_precision",
					Description: `(Optional) Specifies the digits SignalFx displays for values plotted on the chart. Defaults to ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "unit_prefix",
					Description: `(Optional) Must be ` + "`" + `"Metric"` + "`" + ` or ` + "`" + `"Binary` + "`" + `". ` + "`" + `"Metric"` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "color_by",
					Description: `(Optional) Must be ` + "`" + `"Dimension"` + "`" + ` or ` + "`" + `"Metric"` + "`" + `. ` + "`" + `"Dimension"` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "minimum_resolution",
					Description: `(Optional) The minimum resolution (in seconds) to use for computing the underlying program.`,
				},
				resource.Attribute{
					Name:        "max_delay",
					Description: `(Optional) How long (in seconds) to wait for late datapoints.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) A string denotes the geographic region associated with the time zone.`,
				},
				resource.Attribute{
					Name:        "disable_sampling",
					Description: `(Optional) If ` + "`" + `false` + "`" + `, samples a subset of the output MTS, which improves UI performance. ` + "`" + `false` + "`" + ` by default`,
				},
				resource.Attribute{
					Name:        "time_range",
					Description: `(Optional) From when to display data. SignalFx time syntax (e.g. ` + "`" + `"-5m"` + "`" + `, ` + "`" + `"-1h"` + "`" + `). Conflicts with ` + "`" + `start_time` + "`" + ` and ` + "`" + `end_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) Seconds since epoch. Used for visualization. Conflicts with ` + "`" + `time_range` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) Seconds since epoch. Used for visualization. Conflicts with ` + "`" + `time_range` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "axes_include_zero",
					Description: `(Optional) Force the chart to display zero on the y-axes, even if none of the data is near zero.`,
				},
				resource.Attribute{
					Name:        "axis_left",
					Description: `(Optional) Set of axis options.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) Label of the left axis.`,
				},
				resource.Attribute{
					Name:        "min_value",
					Description: `(Optional) The minimum value for the left axis.`,
				},
				resource.Attribute{
					Name:        "max_value",
					Description: `(Optional) The maximum value for the left axis.`,
				},
				resource.Attribute{
					Name:        "high_watermark",
					Description: `(Optional) A line to draw as a high watermark.`,
				},
				resource.Attribute{
					Name:        "high_watermark_label",
					Description: `(Optional) A label to attach to the high watermark line.`,
				},
				resource.Attribute{
					Name:        "low_watermark_label",
					Description: `(Optional) A label to attach to the low watermark line.`,
				},
				resource.Attribute{
					Name:        "axis_right",
					Description: `(Optional) Set of axis options.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) Label of the right axis.`,
				},
				resource.Attribute{
					Name:        "min_value",
					Description: `(Optional) The minimum value for the right axis.`,
				},
				resource.Attribute{
					Name:        "max_value",
					Description: `(Optional) The maximum value for the right axis.`,
				},
				resource.Attribute{
					Name:        "high_watermark",
					Description: `(Optional) A line to draw as a high watermark.`,
				},
				resource.Attribute{
					Name:        "high_watermark_label",
					Description: `(Optional) A label to attach to the high watermark line.`,
				},
				resource.Attribute{
					Name:        "low_watermark_label",
					Description: `(Optional) A label to attach to the low watermark line.`,
				},
				resource.Attribute{
					Name:        "viz_options",
					Description: `(Optional) Plot-level customization options, associated with a publish statement.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Label used in the publish statement that displays the plot (metric time series data) you want to customize.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.`,
				},
				resource.Attribute{
					Name:        "axis",
					Description: `(Optional) Y-axis associated with values for this plot. Must be either ` + "`" + `right` + "`" + ` or ` + "`" + `left` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "plot_type",
					Description: `(Optional) The visualization style to use. Must be ` + "`" + `"LineChart"` + "`" + `, ` + "`" + `"AreaChart"` + "`" + `, ` + "`" + `"ColumnChart"` + "`" + `, or ` + "`" + `"Histogram"` + "`" + `. Chart level ` + "`" + `plot_type` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "value_unit",
					Description: `(Optional) A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are ` + "`" + `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "histogram_options",
					Description: `(Optional) Only used when ` + "`" + `plot_type` + "`" + ` is ` + "`" + `"Histogram"` + "`" + `. Histogram specific options.`,
				},
				resource.Attribute{
					Name:        "color_theme",
					Description: `(Optional) Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine, red, gold, greenyellow, chartreuse, jade`,
				},
				resource.Attribute{
					Name:        "legend_fields_to_hide",
					Description: `(Optional) List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use ` + "`" + `legend_options_fields` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "legend_options_fields",
					Description: `(Optional) List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with ` + "`" + `legend_fields_to_hide` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "on_chart_legend_dimension",
					Description: `(Optional) Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: ` + "`" + `"metric"` + "`" + `, ` + "`" + `"plot_label"` + "`" + ` and any dimension.`,
				},
				resource.Attribute{
					Name:        "show_event_lines",
					Description: `(Optional) Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "show_data_markers",
					Description: `(Optional) Show markers (circles) for each datapoint used to draw line or area charts. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "stacked",
					Description: `(Optional) Whether area and bar charts in the visualization should be stacked. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). ` + "`" + `"UTC"` + "`" + ` by default.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"signalfx_dashboard":             0,
		"signalfx_dashboard_group":       1,
		"signalfx_detector":              2,
		"signalfx_event_feed_chart":      3,
		"signalfx_heatmap_chart":         4,
		"signalfx_integration":           5,
		"signalfx_integration_gcp":       6,
		"signalfx_integration_pagerduty": 7,
		"signalfx_integration_slack":     8,
		"signalfx_list_chart":            9,
		"signalfx_single_value_chart":    10,
		"signalfx_text_chart":            11,
		"signalfx_time_chart":            12,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
