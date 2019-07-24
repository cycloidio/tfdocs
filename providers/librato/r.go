package librato

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "librato_alert",
			Category:         "Resources",
			ShortDescription: `Provides a Librato Alert resource. This can be used to create and manage alerts on Librato.`,
			Description:      ``,
			Keywords: []string{
				"alert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the alert.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the alert.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `whether the alert is active (can be triggered). Defaults to true.`,
				},
				resource.Attribute{
					Name:        "rearm_seconds",
					Description: `minimum amount of time between sending alert notifications, in seconds.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `list of notification service IDs.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `A trigger condition for the alert. Conditions documented below.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `A hash of additional attribtues for the alert. Attributes documented below. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the alert.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the alert.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `whether the alert is active (can be triggered). Defaults to true.`,
				},
				resource.Attribute{
					Name:        "rearm_seconds",
					Description: `minimum amount of time between sending alert notifications, in seconds.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `list of notification service IDs.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `A trigger condition for the alert. Conditions documented below. Conditions (` + "`" + `condition` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of condition. Must be one of ` + "`" + `above` + "`" + `, ` + "`" + `below` + "`" + ` or ` + "`" + `absent` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "detect_reset",
					Description: `boolean: toggles the method used to calculate the delta from the previous sample when the summary_function is ` + "`" + `derivative` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `number of seconds condition must be true to fire the alert (required for type ` + "`" + `absent` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `float: measurements over this number will fire the alert (only for ` + "`" + `above` + "`" + ` or ` + "`" + `below` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "summary_function",
					Description: `Indicates which statistic of an aggregated measurement to alert on. ((only for ` + "`" + `above` + "`" + ` or ` + "`" + `below` + "`" + `). Attributes (` + "`" + `attributes` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "runbook_url",
					Description: `a URL for the runbook to be followed when this alert is firing. Used in the Librato UI if set.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the alert.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the alert.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `whether the alert is active (can be triggered). Defaults to true.`,
				},
				resource.Attribute{
					Name:        "rearm_seconds",
					Description: `minimum amount of time between sending alert notifications, in seconds.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `list of notification service IDs.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `A trigger condition for the alert. Conditions documented below. Conditions (` + "`" + `condition` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of condition. Must be one of ` + "`" + `above` + "`" + `, ` + "`" + `below` + "`" + ` or ` + "`" + `absent` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "detect_reset",
					Description: `boolean: toggles the method used to calculate the delta from the previous sample when the summary_function is ` + "`" + `derivative` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `number of seconds condition must be true to fire the alert (required for type ` + "`" + `absent` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `float: measurements over this number will fire the alert (only for ` + "`" + `above` + "`" + ` or ` + "`" + `below` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "summary_function",
					Description: `Indicates which statistic of an aggregated measurement to alert on. ((only for ` + "`" + `above` + "`" + ` or ` + "`" + `below` + "`" + `). Attributes (` + "`" + `attributes` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "runbook_url",
					Description: `a URL for the runbook to be followed when this alert is firing. Used in the Librato UI if set.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "librato_metric",
			Category:         "Resources",
			ShortDescription: `Provides a Librato Metric resource. This can be used to create and manage metrics on Librato.`,
			Description:      ``,
			Keywords: []string{
				"metric",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of metric to create (gauge, counter, or composite).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique identifier of the metric.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The name which will be used for the metric when viewing the Metrics website.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Text that can be used to explain precisely what the metric is measuring.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Number of seconds that is the standard reporting period of the metric.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `The attributes hash configures specific components of a metric’s visualization.`,
				},
				resource.Attribute{
					Name:        "composite",
					Description: `The definition of the composite metric. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The identifier for the metric.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The name which will be used for the metric when viewing the Metrics website.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of metric to create (gauge, counter, or composite).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Text that describes precisely what the metric is measuring.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Number of seconds that is the standard reporting period of the metric. Setting the period enables Metrics to detect abnormal interruptions in reporting and aids in analytics. For gauge metrics that have service-side aggregation enabled, this option will define the period that aggregation occurs on.`,
				},
				resource.Attribute{
					Name:        "composite",
					Description: `The composite definition. Only used when type is composite. Attributes (` + "`" + `attributes` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Sets a default color to prefer when visually rendering the metric. Must be a seven character string that represents the hex code of the color e.g. #52D74C.`,
				},
				resource.Attribute{
					Name:        "display_max",
					Description: `If a metric has a known theoretical maximum value, set display_max so that visualizations can provide perspective of the current values relative to the maximum value.`,
				},
				resource.Attribute{
					Name:        "display_min",
					Description: `If a metric has a known theoretical minimum value, set display_min so that visualizations can provide perspective of the current values relative to the minimum value.`,
				},
				resource.Attribute{
					Name:        "display_units_long",
					Description: `A string that identifies the unit of measurement e.g. Microseconds. Typically the long form of display_units_short and used in visualizations e.g. the Y-axis label on a graph.`,
				},
				resource.Attribute{
					Name:        "display_units_short",
					Description: `A terse (usually abbreviated) string that identifies the unit of measurement e.g. uS (Microseconds). Typically the short form of display_units_long and used in visualizations e.g. the tooltip for a point on a graph.`,
				},
				resource.Attribute{
					Name:        "display_stacked",
					Description: `A boolean value indicating whether or not multiple metric streams should be aggregated in a visualization (e.g. stacked graphs). By default counters have display_stacked enabled while gauges have it disabled.`,
				},
				resource.Attribute{
					Name:        "summarize_function",
					Description: `Determines how to calculate values when rolling up from raw values to higher resolution intervals. Must be one of: ‘average’, 'sum’, 'count’, 'min’, 'max’. If summarize_function is not set the behavior defaults to average. If the values of the measurements to be rolled up are: 2, 10, 5:`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `Enable service-side aggregation for this metric. When enabled, measurements sent using the same tag set will be aggregated into single measurements on an interval defined by the period of the metric. If there is no period defined for the metric then all measurements will be aggregated on a 60-second interval. This option takes a value of true or false. If this option is not set for a metric it will default to false.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The identifier for the metric.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The name which will be used for the metric when viewing the Metrics website.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of metric to create (gauge, counter, or composite).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Text that describes precisely what the metric is measuring.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Number of seconds that is the standard reporting period of the metric. Setting the period enables Metrics to detect abnormal interruptions in reporting and aids in analytics. For gauge metrics that have service-side aggregation enabled, this option will define the period that aggregation occurs on.`,
				},
				resource.Attribute{
					Name:        "composite",
					Description: `The composite definition. Only used when type is composite. Attributes (` + "`" + `attributes` + "`" + `) support the following:`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `Sets a default color to prefer when visually rendering the metric. Must be a seven character string that represents the hex code of the color e.g. #52D74C.`,
				},
				resource.Attribute{
					Name:        "display_max",
					Description: `If a metric has a known theoretical maximum value, set display_max so that visualizations can provide perspective of the current values relative to the maximum value.`,
				},
				resource.Attribute{
					Name:        "display_min",
					Description: `If a metric has a known theoretical minimum value, set display_min so that visualizations can provide perspective of the current values relative to the minimum value.`,
				},
				resource.Attribute{
					Name:        "display_units_long",
					Description: `A string that identifies the unit of measurement e.g. Microseconds. Typically the long form of display_units_short and used in visualizations e.g. the Y-axis label on a graph.`,
				},
				resource.Attribute{
					Name:        "display_units_short",
					Description: `A terse (usually abbreviated) string that identifies the unit of measurement e.g. uS (Microseconds). Typically the short form of display_units_long and used in visualizations e.g. the tooltip for a point on a graph.`,
				},
				resource.Attribute{
					Name:        "display_stacked",
					Description: `A boolean value indicating whether or not multiple metric streams should be aggregated in a visualization (e.g. stacked graphs). By default counters have display_stacked enabled while gauges have it disabled.`,
				},
				resource.Attribute{
					Name:        "summarize_function",
					Description: `Determines how to calculate values when rolling up from raw values to higher resolution intervals. Must be one of: ‘average’, 'sum’, 'count’, 'min’, 'max’. If summarize_function is not set the behavior defaults to average. If the values of the measurements to be rolled up are: 2, 10, 5:`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `Enable service-side aggregation for this metric. When enabled, measurements sent using the same tag set will be aggregated into single measurements on an interval defined by the period of the metric. If there is no period defined for the metric then all measurements will be aggregated on a 60-second interval. This option takes a value of true or false. If this option is not set for a metric it will default to false.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "librato_service",
			Category:         "Resources",
			ShortDescription: `Provides a Librato service resource. This can be used to create and manage notification services on Librato.`,
			Description:      ``,
			Keywords: []string{
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of notificaion.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) The alert title.`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `(Required) a JSON hash of settings specific to the alert type. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of notificaion.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The alert title.`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `a JSON hash of settings specific to the alert type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of notificaion.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The alert title.`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `a JSON hash of settings specific to the alert type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "librato_space",
			Category:         "Resources",
			ShortDescription: `Provides a Librato Space resource. This can be used to create and manage spaces on Librato.`,
			Description:      ``,
			Keywords: []string{
				"space",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the space. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the space.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the space.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the space.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the space.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "librato_space_chart",
			Category:         "Resources",
			ShortDescription: `Provides a Librato Space Chart resource. This can be used to create and manage charts in Librato Spaces.`,
			Description:      ``,
			Keywords: []string{
				"space",
				"chart",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "space_id",
					Description: `(Required) The ID of the space this chart should be in.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The title of the chart when it is displayed.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Indicates the type of chart. Must be one of line or stacked (default to line).`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `(Optional) The minimum display value of the chart's Y-axis.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `(Optional) The maximum display value of the chart's Y-axis.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The Y-axis label.`,
				},
				resource.Attribute{
					Name:        "related_space",
					Description: `(Optional) The ID of another space to which this chart is related.`,
				},
				resource.Attribute{
					Name:        "stream",
					Description: `(Optional) Nested block describing a metric to use for data in the chart. The structure of this block is described below. The ` + "`" + `stream` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Required) The name of the metric. May not be specified if ` + "`" + `composite` + "`" + ` is specified.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The name of a source, or ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "group_function",
					Description: `(Required) How to process the results when multiple sources will be returned. Value must be one of average, sum, breakout. If average or sum, a single line will be drawn representing the average or sum (respectively) of all sources. If the group_function is breakout, a separate line will be drawn for each source. If this property is not supplied, the behavior will default to average. May not be specified if ` + "`" + `composite` + "`" + ` is specified.`,
				},
				resource.Attribute{
					Name:        "composite",
					Description: `(Required) A composite metric query string to execute when this stream is displayed. May not be specified if ` + "`" + `metric` + "`" + `, ` + "`" + `source` + "`" + ` or ` + "`" + `group_function` + "`" + ` is specified.`,
				},
				resource.Attribute{
					Name:        "summary_function",
					Description: `(Optional) When visualizing complex measurements or a rolled-up measurement, this allows you to choose which statistic to use. Defaults to "average". Valid options are: "max", "min", "average", "sum" or "count".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A display name to use for the stream when generating the tooltip.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Sets a color to use when rendering the stream. Must be a seven character string that represents the hex code of the color e.g. "#52D74C".`,
				},
				resource.Attribute{
					Name:        "units_short",
					Description: `(Optional) Unit value string to use as the tooltip label.`,
				},
				resource.Attribute{
					Name:        "units_long",
					Description: `(Optional) String value to set as they Y-axis label. All streams that share the same units_long value will be plotted on the same Y-axis.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `(Optional) Theoretical minimum Y-axis value.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `(Optional) Theoretical maximum Y-axis value.`,
				},
				resource.Attribute{
					Name:        "transform_function",
					Description: `(Optional) Linear formula to run on each measurement prior to visualization.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) An integer value of seconds that defines the period this stream reports at. This aids in the display of the stream and allows the period to be used in stream display transforms. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the chart.`,
				},
				resource.Attribute{
					Name:        "space_id",
					Description: `The ID of the space this chart should be in.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The title of the chart when it is displayed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the chart.`,
				},
				resource.Attribute{
					Name:        "space_id",
					Description: `The ID of the space this chart should be in.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The title of the chart when it is displayed.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"librato_alert":       0,
		"librato_metric":      1,
		"librato_service":     2,
		"librato_space":       3,
		"librato_space_chart": 4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
