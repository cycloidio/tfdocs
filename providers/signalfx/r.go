package signalfx

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "signalfx_alert_muting_rule",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx Alert Muting Rules`,
			Description: `

Provides a SignalFx resource for managing alert muting rules. See [Mute Notifications](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html) for more information.

~> **WARNING** SignalFx does not allow the start time of a **currently active** muting rule to be modified. As such, attempting to modify a currently active rule will destroy the existing rule and create a new rule. This may result in the emission of notifications.

`,
			Keywords: []string{
				"alert",
				"muting",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description for this muting rule`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) Starting time of an alert muting rule as a Unit time stamp in seconds.`,
				},
				resource.Attribute{
					Name:        "stop_time",
					Description: `(Optional) Starting time of an alert muting rule as a Unix time stamp in seconds.`,
				},
				resource.Attribute{
					Name:        "detectors",
					Description: `(Optional) A convenience attribute that associated this muting rule with specific detector ids.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filters for this rule. See [Creating muting rules from scratch](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html#rule-from-scratch) for more information.`,
				},
				resource.Attribute{
					Name:        "property",
					Description: `(Required) The property to filter.`,
				},
				resource.Attribute{
					Name:        "property_value",
					Description: `(Required) The property value to filter.`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional) Determines if this is a "not" filter. Defaults to ` + "`" + `false` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_aws_external_integration",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx AWS External ID Integrations`,
			Description: `

SignalFx AWS CloudWatch integrations using Role ARNs. For help with this integration see [Connect to AWS CloudWatch](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#connect-to-aws).

~> **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider.

~> **WARNING** This resource implements a part of a workflow. You must use it with ` + "`" + `signalfx_aws_integration` + "`" + `. Check with SignalFx support for your realm's AWS account id.

`,
			Keywords: []string{
				"aws",
				"external",
				"integration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of this integration ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this integration, used with ` + "`" + `signalfx_aws_integration` + "`" + ``,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `The external ID to use with your IAM role and with ` + "`" + `signalfx_aws_integration` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "signalfx_aws_account",
					Description: `The AWS Account ARN to use with your policies/roles, provided by SignalFx.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this integration, used with ` + "`" + `signalfx_aws_integration` + "`" + ``,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `The external ID to use with your IAM role and with ` + "`" + `signalfx_aws_integration` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "signalfx_aws_account",
					Description: `The AWS Account ARN to use with your policies/roles, provided by SignalFx.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_aws_integration",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx AWS Integrations`,
			Description: `

SignalFx AWS CloudWatch integrations. For help with this integration see [Monitoring Amazon Web Services](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#monitor-amazon-web-services).

~> **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider.

~> **WARNING** This resource implements a part of a workflow. You must use it with one of either ` + "`" + `signalfx_aws_external_integration` + "`" + ` or ` + "`" + `signalfx_aws_token_integration` + "`" + `.

`,
			Keywords: []string{
				"aws",
				"integration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Whether the integration is enabled.`,
				},
				resource.Attribute{
					Name:        "integration_id",
					Description: `(Required) The id of one of a ` + "`" + `signalfx_aws_external_integration` + "`" + ` or ` + "`" + `signalfx_aws_token_integration` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Required) The ` + "`" + `external_id` + "`" + ` property from one of a ` + "`" + `signalfx_aws_external_integration` + "`" + ` or ` + "`" + `signalfx_aws_token_integration` + "`" + ``,
				},
				resource.Attribute{
					Name:        "custom_cloudwatch_namespaces",
					Description: `(Optional) List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; SignalFx imports the metrics so you can monitor them.`,
				},
				resource.Attribute{
					Name:        "custom_namespace_sync_rule",
					Description: `(Optional) Each element controls the data collected by SignalFx for the specified namespace. Conflicts with the ` + "`" + `custom_cloudwatch_namespaces` + "`" + ` property.`,
				},
				resource.Attribute{
					Name:        "default_action",
					Description: `(Optional) Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of ` + "`" + `"Include"` + "`" + ` or ` + "`" + `"Exclude"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter_action",
					Description: `(Optional) Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of ` + "`" + `"Include"` + "`" + ` or ` + "`" + `"Exclude"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter_source",
					Description: `(Optional) Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow ` + "`" + `filter()` + "`" + ` function; it can be any valid SignalFlow filter expression.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.`,
				},
				resource.Attribute{
					Name:        "namespace_sync_rule",
					Description: `(Optional) Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the ` + "`" + `services` + "`" + ` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.`,
				},
				resource.Attribute{
					Name:        "default_action",
					Description: `(Optional) Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of ` + "`" + `"Include"` + "`" + ` or ` + "`" + `"Exclude"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter_action",
					Description: `(Optional) Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of ` + "`" + `"Include"` + "`" + ` or ` + "`" + `"Exclude"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "filter_source",
					Description: `(Optional) Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow ` + "`" + `filter()` + "`" + ` function; it can be any valid SignalFlow filter expression.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.`,
				},
				resource.Attribute{
					Name:        "enable_aws_usage",
					Description: `(Optional) Flag that controls how SignalFx imports usage metrics from AWS to use with AWS Cost Optimizer. If ` + "`" + `true` + "`" + `, SignalFx imports the metrics.`,
				},
				resource.Attribute{
					Name:        "import_cloud_watch",
					Description: `(Optional) Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) If you specify ` + "`" + `auth_method = \"SecurityToken\"` + "`" + ` in your request to create an AWS integration object, use this property to specify the key.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Optional) List of AWS regions that SignalFx should monitor.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Optional) Role ARN that you add to an existing AWS integration object.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional) List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with ` + "`" + `namespace_sync_rule` + "`" + `. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.`,
				},
				resource.Attribute{
					Name:        "poll_rate",
					Description: `(Optional) AWS poll rate (in seconds). One of ` + "`" + `60` + "`" + ` or ` + "`" + `300` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_get_metric_data_method",
					Description: `(Optional) Enable the use of Amazon's ` + "`" + `GetMetricData` + "`" + ` for collecting metrics. Note that this requires the inclusion of the ` + "`" + `"cloudwatch:GetMetricData"` + "`" + ` permission.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_aws_token_integration",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx AWS Security Token Integrations`,
			Description: `

SignalFx AWS CloudWatch integrations using security tokens. For help with this integration see [Connect to AWS CloudWatch](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#connect-to-aws).

~> **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider.

~> **WARNING** This resource implements a part of a workflow. You must use it with ` + "`" + `signalfx_aws_integration` + "`" + `.

`,
			Keywords: []string{
				"aws",
				"token",
				"integration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of this integration ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the integration to use with ` + "`" + `signalfx_aws_integration` + "`" + ``,
				},
				resource.Attribute{
					Name:        "signalfx_aws_account",
					Description: `The AWS Account ARN to use with your policies/roles, provided by SignalFx.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the integration to use with ` + "`" + `signalfx_aws_integration` + "`" + ``,
				},
				resource.Attribute{
					Name:        "signalfx_aws_account",
					Description: `The AWS Account ARN to use with your policies/roles, provided by SignalFx.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_azure_integration",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx Azure Integrations`,
			Description: `

SignalFx Azure integrations. For help with this integration see [Monitoring Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure).

~> **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.

`,
			Keywords: []string{
				"azure",
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
					Name:        "app_id",
					Description: `(Required) Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.`,
				},
				resource.Attribute{
					Name:        "poll_rate",
					Description: `(Optional) AWS poll rate (in seconds). One of ` + "`" + `60` + "`" + ` or ` + "`" + `300` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required) Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.`,
				},
				resource.Attribute{
					Name:        "subscriptions",
					Description: `(Required) List of Azure subscriptions that SignalFx should monitor.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_dashboard",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx Dashboards`,
			Description: `

A dashboard is a curated collection of specific charts and supports dimensional [filters](http://docs.signalfx.com/en/latest/dashboards/dashboard-filter-dynamic.html#filter-dashboard-charts), [dashboard variables](http://docs.signalfx.com/en/latest/dashboards/dashboard-filter-dynamic.html#dashboard-variables) and [time range](http://docs.signalfx.com/en/latest/_sidebars-and-includes/using-time-range-selector.html#time-range-selector) options. These options are applied to all charts in the dashboard, providing a consistent view of the data displayed in that dashboard. This also means that when you open a chart to drill down for more details, you are viewing the same data that is visible in the dashboard view.

~> **NOTE** Since every dashboard is included in a [dashboard group](dashboard_group.html) (SignalFx collection of dashboards), you need to create that first and reference it as shown in the example.

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

In the SignalFx web UI, a [dashboard group](https://developers.signalfx.com/dashboard_groups_reference.html) is a collection of dashboards.

~> **NOTE** Dashboard groups cannot be accessed directly, but just via a dashboard contained in them. This is the reason why make show won't show any of yours dashboard groups.

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
					Name:        "authorized_writer_teams",
					Description: `(Optional) Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in ` + "`" + `authorized_writer_teams` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "authorized_writer_users",
					Description: `(Optional) User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in ` + "`" + `authorized_writer_teams` + "`" + `).`,
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
			Type:             "signalfx_data_link",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx data link`,
			Description: `

Manage SignalFx [Data Links](https://docs.signalfx.com/en/latest/managing/data-links.html).

`,
			Keywords: []string{
				"data",
				"link",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "property_name",
					Description: `(Optional) Name (key) of the metadata that's the trigger of a data link. If you specify ` + "`" + `property_value` + "`" + `, you must specify ` + "`" + `property_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "property_value",
					Description: `(Optional) Value of the metadata that's the trigger of a data link. If you specify this property, you must also specify ` + "`" + `property_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "context_dashboard_id",
					Description: `(Optional) If provided, scopes this data link to the supplied dashboard id. If omitted then the link will be global.`,
				},
				resource.Attribute{
					Name:        "target_external_url",
					Description: `(Optional) Link to an external URL`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Flag that designates a target as the default for a data link object. ` + "`" + `true` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "time_format",
					Description: `(Optional) [Designates the format](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) of ` + "`" + `minimum_time_window` + "`" + ` in the same data link target object. Must be one of ` + "`" + `"ISO8601"` + "`" + `, ` + "`" + `"EpochSeconds"` + "`" + ` or ` + "`" + `"Epoch"` + "`" + ` (which is milliseconds). Defaults to ` + "`" + `"ISO8601"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minimum_time_window",
					Description: `(Optional) The [minimum time window](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) for a search sent to an external site. Defaults to ` + "`" + `6000` + "`" + ``,
				},
				resource.Attribute{
					Name:        "property_key_mapping",
					Description: `Describes the relationship between SignalFx metadata keys and external system properties when the key names are different.`,
				},
				resource.Attribute{
					Name:        "target_signalfx_dashboard",
					Description: `(Optional) Link to a SignalFx dashboard`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Flag that designates a target as the default for a data link object. ` + "`" + `true` + "`" + ` by default`,
				},
				resource.Attribute{
					Name:        "dashboard_id",
					Description: `(Required) SignalFx-assigned ID of the dashboard link target`,
				},
				resource.Attribute{
					Name:        "dashboard_group_id",
					Description: `(Required) SignalFx-assigned ID of the dashboard link target's dashboard group`,
				},
				resource.Attribute{
					Name:        "target_splunk",
					Description: `(Optional) Link to an external URL`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Flag that designates a target as the default for a data link object. ` + "`" + `true` + "`" + ` by default`,
				},
				resource.Attribute{
					Name:        "property_key_mapping",
					Description: `Describes the relationship between SignalFx metadata keys and external system properties when the key names are different.`,
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

~> **NOTE** If you're interested in using SignalFx detector features such as Historical Anomaly, Resource Running Out, or others then consider building them in the UI first then using the "Show SignalFlow" feature to extract the value for ` + "`" + `program_text` + "`" + `. You may also consult the [documentation for detector functions in signalflow-library](https://github.com/signalfx/signalflow-library/tree/master/library/signalfx/detectors).

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
					Description: `(Required) Signalflow program text for the detector. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the detector.`,
				},
				resource.Attribute{
					Name:        "authorized_writer_teams",
					Description: `(Optional) Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in ` + "`" + `authorized_writer_users` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "authorized_writer_users",
					Description: `(Optional) User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in ` + "`" + `authorized_writer_teams` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "max_delay",
					Description: `(Optional) How long (in seconds) to wait for late datapoints. See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info. Max value is ` + "`" + `900` + "`" + ` seconds (15 minutes). ` + "`" + `Auto` + "`" + ` (as little as possible) by default.`,
				},
				resource.Attribute{
					Name:        "show_data_markers",
					Description: `(Optional) When ` + "`" + `true` + "`" + `, markers will be drawn for each datapoint within the visualization. ` + "`" + `true` + "`" + ` by default.`,
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
					Description: `(Optional) Seconds to display in the visualization. This is a rolling range from the current time. Example: ` + "`" + `3600` + "`" + ` corresponds to ` + "`" + `-1h` + "`" + ` in web UI. ` + "`" + `3600` + "`" + ` by default.`,
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
					Description: `(Optional) Team IDs to associate the detector to.`,
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
					Name:        "description",
					Description: `(Optional) Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) When true, notifications and events will not be generated for the detect label. ` + "`" + `false` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `(Optional) List of strings specifying where notifications will be sent when an incident occurs. See [Create A Single Detector](https://developers.signalfx.com/detectors_reference.html#operation/Create%20Single%20Detector) for more info.`,
				},
				resource.Attribute{
					Name:        "parameterized_body",
					Description: `(Optional) Custom notification message body when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.signalfx.com/en/latest/detect-alert/set-up-detectors.html#about-detectors#alert-settings) for more info.`,
				},
				resource.Attribute{
					Name:        "parameterized_subject",
					Description: `(Optional) Custom notification message subject when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.signalfx.com/en/latest/detect-alert/set-up-detectors.html#about-detectors#alert-settings) for more info.`,
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
					Name:        "id",
					Description: `ID of the SignalFx detector ## Import Detectors can be imported using their string ID (recoverable from URL: ` + "`" + `/#/detector/v2/abc123/edit` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import signalfx_detector.application_delay abc123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SignalFx detector ## Import Detectors can be imported using their string ID (recoverable from URL: ` + "`" + `/#/detector/v2/abc123/edit` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import signalfx_detector.application_delay abc123 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).`,
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
			Type:             "signalfx_gcp_integration",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx GCP Integrations`,
			Description: `

SignalFx GCP Integration

~> **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.

`,
			Keywords: []string{
				"gcp",
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
					Name:        "poll_rate",
					Description: `(Required) GCP integration poll rate in seconds. Can be set to either 60 or 300 (1 minute or 5 minutes).`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional) GCP service metrics to import. Can be an empty list, or not included, to import 'All services'. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valid values.`,
				},
				resource.Attribute{
					Name:        "project_service_keys",
					Description: `(Required) GCP projects to add.`,
				},
				resource.Attribute{
					Name:        "whilelist",
					Description: `(Optional) [Compute Metadata Whitelist](https://docs.signalfx.com/en/latest/integrations/google-cloud-platform.html#compute-engine-instance).`,
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
					Description: `(Optional, Default) Values and color for the color range. Example: ` + "`" + `color_range : { min : 0, max : 100, color : "#0000ff" }` + "`" + `. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).`,
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
					Description: `(Required) The color range to use. The starting hex color value for data values in a heatmap chart. Specify the value as a 6-character hexadecimal value preceded by the '#' character, for example "#ea1849" (grass green).`,
				},
				resource.Attribute{
					Name:        "color_scale",
					Description: `(Optional. Conflicts with ` + "`" + `color_range` + "`" + `) One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: ` + "`" + `color_scale { gt = 60, color = "blue" } color_scale { lte = 60, color = "yellow" }` + "`" + `. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).`,
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
					Description: `(Optional) Indicates the upper threshold non-inclusive value for this range.`,
				},
				resource.Attribute{
					Name:        "lte",
					Description: `(Optional) Indicates the upper threshold inclusive value for this range.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Required) The color range to use. Hex values are not supported here. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_jira_integration",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx Jira Integrations`,
			Description: `

SignalFx Jira integrations. For help with this integration see [Integration with Jira](https://docs.signalfx.com/en/latest/admin-guide/integrate-notifications.html#integrate-with-jira).

~> **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.

`,
			Keywords: []string{
				"jira",
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
					Name:        "auth_method",
					Description: `(Required) Authentication method used when creating the Jira integration. One of ` + "`" + `EmailAndToken` + "`" + ` (using ` + "`" + `user_email` + "`" + ` and ` + "`" + `api_token` + "`" + `) or ` + "`" + `UsernameAndPassword` + "`" + ` (using ` + "`" + `username` + "`" + ` and ` + "`" + `password` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "api_token",
					Description: `(Required if ` + "`" + `auth_method` + "`" + ` is ` + "`" + `EmailAndToken` + "`" + `) The API token for the user email`,
				},
				resource.Attribute{
					Name:        "user_email",
					Description: `(Required if ` + "`" + `auth_method` + "`" + ` is ` + "`" + `EmailAndToken` + "`" + `) Email address used to authenticate the Jira integration.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required if ` + "`" + `auth_method` + "`" + ` is ` + "`" + `UsernameAndPassword` + "`" + `) User name used to authenticate the Jira integration.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required if ` + "`" + `auth_method` + "`" + ` is ` + "`" + `UsernameAndPassword` + "`" + `) Password used to authenticate the Jira integration.`,
				},
				resource.Attribute{
					Name:        "base_url",
					Description: `(Required) Base URL of the Jira instance that's integrated with SignalFx.`,
				},
				resource.Attribute{
					Name:        "issue_type",
					Description: `(Required) Issue type (for example, Story) for tickets that Jira creates for detector notifications. SignalFx validates issue types, so you must specify a type that's valid for the Jira project specified in ` + "`" + `projectKey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_key",
					Description: `(Required) Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.`,
				},
				resource.Attribute{
					Name:        "assignee_name",
					Description: `(Required) Jira user name for the assignee.`,
				},
				resource.Attribute{
					Name:        "assignee_display_name",
					Description: `(Optional) Jira display name for the assignee.`,
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
					Description: `(Required) Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).`,
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
					Description: `(Optional) Must be one of ` + "`" + `"Scale"` + "`" + `, ` + "`" + `"Dimension"` + "`" + ` or ` + "`" + `"Metric"` + "`" + `. ` + "`" + `"Dimension"` + "`" + ` by default.`,
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
					Description: `(Optional) The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.`,
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
					Description: `(Required) The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.`,
				},
				resource.Attribute{
					Name:        "sort_by",
					Description: `(Optional) The property to use when sorting the elements. Use ` + "`" + `value` + "`" + ` if you want to sort by value. Must be prepended with ` + "`" + `+` + "`" + ` for ascending or ` + "`" + `-` + "`" + ` for descending (e.g. ` + "`" + `-foo` + "`" + `). Note there are some special values for some of the options provided in the UX: ` + "`" + `"value"` + "`" + ` for Value, ` + "`" + `"sf_originatingMetric"` + "`" + ` for Metric, and ` + "`" + `"sf_metric"` + "`" + ` for plot.`,
				},
				resource.Attribute{
					Name:        "time_range",
					Description: `(Optional) How many seconds ago from which to display data. For example, the last hour would be ` + "`" + `3600` + "`" + `, etc. Conflicts with ` + "`" + `start_time` + "`" + ` and ` + "`" + `end_time` + "`" + `.`,
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
			Type:             "signalfx_opsgenie_integration",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx Opsgenie Integrations`,
			Description: `

SignalFx Opsgenie integration.

~> **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.

`,
			Keywords: []string{
				"opsgenie",
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
					Name:        "api_key",
					Description: `(Required) The API key`,
				},
				resource.Attribute{
					Name:        "api_url",
					Description: `(Optional) Opsgenie API URL. Will default to ` + "`" + `https://api.opsgenie.com` + "`" + `. You might also want ` + "`" + `https://api.eu.opsgenie.com` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_org_token",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx text notes`,
			Description: `

Manage SignalFx org tokens.

`,
			Keywords: []string{
				"org",
				"token",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the token.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Flag that controls enabling the token. If set to ` + "`" + `true` + "`" + `, the token is disabled, and you can't use it for authentication. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `The secret token created by the API. You cannot set this value.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `(Optional) Where to send notifications about this token's limits. Please consult the [Notification Format](https://www.terraform.io/docs/providers/signalfx/r/detector.html#notification-format) laid out in detectors.`,
				},
				resource.Attribute{
					Name:        "host_or_usage_limits",
					Description: `(Optional) Specify Usage-based limits for this token.`,
				},
				resource.Attribute{
					Name:        "host_limit",
					Description: `(Optional) Max number of hosts that can use this token`,
				},
				resource.Attribute{
					Name:        "host_notification_threshold",
					Description: `(Optional) Notification threshold for hosts`,
				},
				resource.Attribute{
					Name:        "container_limit",
					Description: `(Optional) Max number of Docker containers that can use this token`,
				},
				resource.Attribute{
					Name:        "container_notification_threshold",
					Description: `(Optional) Notification threshold for Docker containers`,
				},
				resource.Attribute{
					Name:        "custom_metrics_limit",
					Description: `(Optional) Max number of custom metrics that can be sent with this token`,
				},
				resource.Attribute{
					Name:        "custom_metrics_notification_threshold",
					Description: `(Optional) Notification threshold for custom metrics`,
				},
				resource.Attribute{
					Name:        "high_res_metrics_limit",
					Description: `(Optional) Max number of hi-res metrics that can be sent with this toke`,
				},
				resource.Attribute{
					Name:        "high_res_metrics_notification_threshold",
					Description: `(Optional) Notification threshold for hi-res metrics`,
				},
				resource.Attribute{
					Name:        "dpm_notification_threshold",
					Description: `(Optional) DPM level at which SignalFx sends the notification for this token. If you don't specify a notification, SignalFx sends the generic notification.`,
				},
				resource.Attribute{
					Name:        "dpm_limit",
					Description: `(Required) The datapoints per minute (dpm) limit for this token. If you exceed this limit, SignalFx sends out an alert.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_pagerduty_integration",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx PagerDuty Integrations`,
			Description: `

SignalFx PagerDuty integrations

~> **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.

`,
			Keywords: []string{
				"pagerduty",
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
					Name:        "api_key",
					Description: `(Required) PagerDuty API key.`,
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
					Description: `(Required) Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).`,
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
					Description: `(Required) The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.`,
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
					Description: `(Optional) The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.`,
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
			Type:             "signalfx_slack_integration",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx Slack Integrations`,
			Description: `

SignalFx Slack integration.

~> **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.

`,
			Keywords: []string{
				"slack",
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
					Name:        "webhook_url",
					Description: `(Required) Slack incoming webhook URL.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_team",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx teams`,
			Description: `

Handles management of SignalFx teams.

You can configure [team notification policies](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html) using this resource and the various ` + "`" + `notifications_*` + "`" + ` properties.

`,
			Keywords: []string{
				"team",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the team.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the team.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) List of user IDs to include in the team.`,
				},
				resource.Attribute{
					Name:        "notifications_critical",
					Description: `(Optional) Where to send notifications for critical alerts`,
				},
				resource.Attribute{
					Name:        "notifications_default",
					Description: `(Optional) Where to send notifications for default alerts`,
				},
				resource.Attribute{
					Name:        "notifications_info",
					Description: `(Optional) Where to send notifications for info alerts`,
				},
				resource.Attribute{
					Name:        "notifications_major",
					Description: `(Optional) Where to send notifications for major alerts`,
				},
				resource.Attribute{
					Name:        "notifications_minor",
					Description: `(Optional) Where to send notifications for minor alerts`,
				},
				resource.Attribute{
					Name:        "notifications_warning",
					Description: `(Optional) Where to send notifications for warning alerts`,
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
					Description: `(Required) Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).`,
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
					Description: `(Optional) How many seconds ago from which to display data. For example, the last hour would be ` + "`" + `3600` + "`" + `, etc. Conflicts with ` + "`" + `start_time` + "`" + ` and ` + "`" + `end_time` + "`" + `.`,
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
					Name:        "event_options",
					Description: `(Optional) Event customization options, associated with a publish statement. You will need to use this to change settings for any ` + "`" + `events(…)` + "`" + ` statements you use.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Label used in the publish statement that displays the event query you want to customize.`,
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
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_victor_ops_integration",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx VictorOps Integrations`,
			Description: `

SignalFx VictorOps integration.

~> **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.

`,
			Keywords: []string{
				"victor",
				"ops",
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
					Name:        "post_url",
					Description: `(Optional) Victor Ops REST API URL.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_webhook_integration",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create and manage SignalFx Webhook Integrations`,
			Description: `

SignalFx Webhook integration.

~> **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.

`,
			Keywords: []string{
				"webhook",
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
					Name:        "url",
					Description: `(Required) The URL to request`,
				},
				resource.Attribute{
					Name:        "shared_secret",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `(Optional) A header to send with the request`,
				},
				resource.Attribute{
					Name:        "header_key",
					Description: `(Required) The key of the header to send`,
				},
				resource.Attribute{
					Name:        "header_value",
					Description: `(Required) The value of the header to send`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"signalfx_alert_muting_rule":        0,
		"signalfx_aws_external_integration": 1,
		"signalfx_aws_integration":          2,
		"signalfx_aws_token_integration":    3,
		"signalfx_azure_integration":        4,
		"signalfx_dashboard":                5,
		"signalfx_dashboard_group":          6,
		"signalfx_data_link":                7,
		"signalfx_detector":                 8,
		"signalfx_event_feed_chart":         9,
		"signalfx_gcp_integration":          10,
		"signalfx_heatmap_chart":            11,
		"signalfx_jira_integration":         12,
		"signalfx_list_chart":               13,
		"signalfx_opsgenie_integration":     14,
		"signalfx_org_token":                15,
		"signalfx_pagerduty_integration":    16,
		"signalfx_single_value_chart":       17,
		"signalfx_slack_integration":        18,
		"signalfx_team":                     19,
		"signalfx_text_chart":               20,
		"signalfx_time_chart":               21,
		"signalfx_victor_ops_integration":   22,
		"signalfx_webhook_integration":      23,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
