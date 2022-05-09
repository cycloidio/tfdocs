package logzio

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "logzio_alert",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "title",
					Description: `Alert title.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Logz.io alert ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date and time in UTC when the alert was first created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Email of the user who first created the alert.`,
				},
				resource.Attribute{
					Name:        "search_timeframe_minutes",
					Description: `The time frame for evaluating the log data is a sliding window, with 1 minute granularity.`,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `Specifies the operator for evaluating the results. Enum: ` + "`" + `LESS_THAN` + "`" + `, ` + "`" + `GREATER_THAN` + "`" + `, ` + "`" + `LESS_THAN_OR_EQUALS` + "`" + `, ` + "`" + `GREATER_THAN_OR_EQUALS` + "`" + `, ` + "`" + `EQUALS` + "`" + `, ` + "`" + `NOT_EQUALS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity_threshold_tiers",
					Description: `Set as many as 5 thresholds, each with its own severity level.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Defaults to ` + "`" + `MEDIUM` + "`" + `. Labels the event with a severity tag. Available severity tags are: ` + "`" + `INFO` + "`" + `, ` + "`" + `LOW` + "`" + `, ` + "`" + `MEDIUM` + "`" + `, ` + "`" + `HIGH` + "`" + `, ` + "`" + `SEVERE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Number of logs per search timeframe.`,
				},
				resource.Attribute{
					Name:        "alert_notification_endpoints",
					Description: `Add email addresses and/or endpoint channels to automatically receive notifications with sample data when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "notification_emails",
					Description: `Add email addresses to automatically receive notifications with sample data when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the event, its significance, and suggested next steps or instructions for the team.`,
				},
				resource.Attribute{
					Name:        "query_string",
					Description: `Search query in Lucene syntax. You can combine filters and a search query to specify the logs you are looking for. You can combine filters and a search query to specify the logs you are looking for.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `You can use ` + "`" + `must` + "`" + ` and ` + "`" + `must_not` + "`" + ` filters. Filters are more efficient compared to a query, so it's recommended to opt for a filter over a ` + "`" + `query_string` + "`" + `, where possible.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags for filtering alerts and triggered alerts. Can be used in Kibana Discover, Kibana dashboards, and more.`,
				},
				resource.Attribute{
					Name:        "group_by_aggregation_fields",
					Description: `Specify 1-3 fields by which to group the results and count them. If you apply a group by operation, the alert returns a count of the results aggregated by unique values.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `True by default. If ` + "`" + `true` + "`" + `, the alert is currently active.`,
				},
				resource.Attribute{
					Name:        "last_triggered_at",
					Description: `Date and time in UTC when the alert last triggered.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `Date and time in UTC when the alert was last updated.`,
				},
				resource.Attribute{
					Name:        "notification_emails",
					Description: `Array of email addresses to be notified when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "suppress_notifications_minutes",
					Description: `Add a waiting period in minutes to space out notifications. (The alert will still trigger but will not send out notifications during the waiting period.)`,
				},
				resource.Attribute{
					Name:        "value_aggregation_field",
					Description: `Specify the field on which to run the aggregation for the trigger condition.`,
				},
				resource.Attribute{
					Name:        "value_aggregation_type",
					Description: `Specifies the aggregation operator. Can be: ` + "`" + `SUM` + "`" + `, ` + "`" + `MIN` + "`" + `, ` + "`" + `MAX` + "`" + `, ` + "`" + `AVG` + "`" + `, ` + "`" + `COUNT` + "`" + `, ` + "`" + `UNIQUE_COUNT` + "`" + `, ` + "`" + `NONE` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date and time in UTC when the alert was first created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Email of the user who first created the alert.`,
				},
				resource.Attribute{
					Name:        "search_timeframe_minutes",
					Description: `The time frame for evaluating the log data is a sliding window, with 1 minute granularity.`,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `Specifies the operator for evaluating the results. Enum: ` + "`" + `LESS_THAN` + "`" + `, ` + "`" + `GREATER_THAN` + "`" + `, ` + "`" + `LESS_THAN_OR_EQUALS` + "`" + `, ` + "`" + `GREATER_THAN_OR_EQUALS` + "`" + `, ` + "`" + `EQUALS` + "`" + `, ` + "`" + `NOT_EQUALS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity_threshold_tiers",
					Description: `Set as many as 5 thresholds, each with its own severity level.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Defaults to ` + "`" + `MEDIUM` + "`" + `. Labels the event with a severity tag. Available severity tags are: ` + "`" + `INFO` + "`" + `, ` + "`" + `LOW` + "`" + `, ` + "`" + `MEDIUM` + "`" + `, ` + "`" + `HIGH` + "`" + `, ` + "`" + `SEVERE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Number of logs per search timeframe.`,
				},
				resource.Attribute{
					Name:        "alert_notification_endpoints",
					Description: `Add email addresses and/or endpoint channels to automatically receive notifications with sample data when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "notification_emails",
					Description: `Add email addresses to automatically receive notifications with sample data when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the event, its significance, and suggested next steps or instructions for the team.`,
				},
				resource.Attribute{
					Name:        "query_string",
					Description: `Search query in Lucene syntax. You can combine filters and a search query to specify the logs you are looking for. You can combine filters and a search query to specify the logs you are looking for.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `You can use ` + "`" + `must` + "`" + ` and ` + "`" + `must_not` + "`" + ` filters. Filters are more efficient compared to a query, so it's recommended to opt for a filter over a ` + "`" + `query_string` + "`" + `, where possible.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags for filtering alerts and triggered alerts. Can be used in Kibana Discover, Kibana dashboards, and more.`,
				},
				resource.Attribute{
					Name:        "group_by_aggregation_fields",
					Description: `Specify 1-3 fields by which to group the results and count them. If you apply a group by operation, the alert returns a count of the results aggregated by unique values.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `True by default. If ` + "`" + `true` + "`" + `, the alert is currently active.`,
				},
				resource.Attribute{
					Name:        "last_triggered_at",
					Description: `Date and time in UTC when the alert last triggered.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `Date and time in UTC when the alert was last updated.`,
				},
				resource.Attribute{
					Name:        "notification_emails",
					Description: `Array of email addresses to be notified when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "suppress_notifications_minutes",
					Description: `Add a waiting period in minutes to space out notifications. (The alert will still trigger but will not send out notifications during the waiting period.)`,
				},
				resource.Attribute{
					Name:        "value_aggregation_field",
					Description: `Specify the field on which to run the aggregation for the trigger condition.`,
				},
				resource.Attribute{
					Name:        "value_aggregation_type",
					Description: `Specifies the aggregation operator. Can be: ` + "`" + `SUM` + "`" + `, ` + "`" + `MIN` + "`" + `, ` + "`" + `MAX` + "`" + `, ` + "`" + `AVG` + "`" + `, ` + "`" + `COUNT` + "`" + `, ` + "`" + `UNIQUE_COUNT` + "`" + `, ` + "`" + `NONE` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_alert_v2",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "title",
					Description: `Alert title.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Logz.io alert ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date and time in UTC when the alert was first created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Email of the user who first created the alert.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date and time in UTC when the alert was last updated.`,
				},
				resource.Attribute{
					Name:        "updated by",
					Description: `Email of the user who last updated the alert.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the event, its significance, and suggested next steps or instructions for the team.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags for filtering alerts and triggered alerts. Can be used in Kibana Discover, dashboards, and more.`,
				},
				resource.Attribute{
					Name:        "search_timeframe_minutes",
					Description: `The time frame for evaluating the log data is a sliding window, with 1 minute granularity.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If ` + "`" + `true` + "`" + `, the alert is currently active.`,
				},
				resource.Attribute{
					Name:        "notification_emails",
					Description: `Array of email addresses to be notified when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "alert_notification_endpoints",
					Description: `Array of IDs of pre-configured endpoint channels to notify when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "suppress_notifications_minutes",
					Description: `Add a waiting period in minutes to space out notifications. (The alert will still trigger but will not send out notifications during the waiting period.)`,
				},
				resource.Attribute{
					Name:        "output_type",
					Description: `Selects the output format for the alert notification. Can be: ` + "`" + `"JSON"` + "`" + ` or ` + "`" + `"TABLE""` + "`" + ` If the alert has no aggregations/group by fields, JSON offers the option to send full sample logs without selecting specific fields.`,
				},
				resource.Attribute{
					Name:        "correlation_operator",
					Description: `Comma separated string of supported operators. Only applicable when multiple sub-components are in use. Selects a logic for correlating the alert’s sub-components. ` + "`" + `AND` + "`" + ` is currently the only supported operator. When AND is the correlation_operator, both sub-components must meet their triggering criteria for the alert to trigger.`,
				},
				resource.Attribute{
					Name:        "joins",
					Description: `Specifies which group by fields must have the same values to trigger the alert. Joins the group by fields from the first and second sub-components. The key represents the index of the sub component in the array. The fields must be ordered pairs of the group by fields already in use in the ` + "`" + `sub_components.query_string` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_components",
					Description: `Sets the search criteria using a search query, filters, group by aggregations, accounts to search, and trigger conditions.`,
				},
				resource.Attribute{
					Name:        "sub_components.query_string",
					Description: `Provide a Kibana search query written in Lucene syntax. The search query together with the filters select for the relevant logs. Cannot be null - send an asterisk wildcard ` + "`" + `"`,
				},
				resource.Attribute{
					Name:        "sub_components.filter_must",
					Description: `Runs Elasticsearch Bool Query filters on the data (before the search query is applied). The most efficient way to grab the logs you are looking for.`,
				},
				resource.Attribute{
					Name:        "sub_components.filter_must_not",
					Description: `Runs Elasticsearch Bool Query filters on the data (before the search query is applied). The most efficient way to grab the logs you are looking for.`,
				},
				resource.Attribute{
					Name:        "sub_components.group_by_aggregation_fields",
					Description: `Specify 1-3 fields by which to group the results and count them. If you apply a group by operation, the alert returns a count of the results aggregated by unique values.`,
				},
				resource.Attribute{
					Name:        "sub_components.value_aggregation_type",
					Description: `Specifies the aggregation operator. Can be: ` + "`" + `"SUM"` + "`" + `, ` + "`" + `"MIN"` + "`" + `, ` + "`" + `"MAX"` + "`" + `, ` + "`" + `"AVG"` + "`" + `, ` + "`" + `"COUNT"` + "`" + `, ` + "`" + `"UNIQUE_COUNT"` + "`" + `, ` + "`" + `"NONE"` + "`" + `. If ` + "`" + `"COUNT"` + "`" + ` or ` + "`" + `"NONE"` + "`" + `, ` + "`" + `value_aggregation_field` + "`" + ` must be null, and ` + "`" + `group_by_aggregation_fields` + "`" + ` fields must not be empty. If any other operator type (other than ` + "`" + `"NONE"` + "`" + ` or ` + "`" + `"COUNT"` + "`" + `), ` + "`" + `value_aggregation_field` + "`" + ` must not be null.`,
				},
				resource.Attribute{
					Name:        "sub_components.value_aggregation_field",
					Description: `Selects the field on which to run the aggregation for the trigger condition. Cannot be a field already in use for ` + "`" + `group_by_aggregation_fields` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_components.should_query_on_all_accounts",
					Description: `Defaults to true. Only applicable when the alert is run from the main account. If true, the alert runs on the main account and all associated searchable sub accounts. If false, specify relevant account IDs for the alert to monitor using the ` + "`" + `account_ids_to_query_on` + "`" + ` field.`,
				},
				resource.Attribute{
					Name:        "sub_components.account_ids_to_query_on",
					Description: `Specify Account IDs to select which accounts the alert should monitor. The alert will be checked only on these accounts.`,
				},
				resource.Attribute{
					Name:        "sub_components.operation",
					Description: `Specifies the operator for evaluating the results. Can be: ` + "`" + `"LESS_THAN"` + "`" + `, ` + "`" + `"GREATER_THAN"` + "`" + `, ` + "`" + `"LESS_THAN_OR_EQUALS"` + "`" + `, ` + "`" + `"GREATER_THAN_OR_EQUALS"` + "`" + `, ` + "`" + `"EQUALS"` + "`" + `, ` + "`" + `"NOT_EQUALS"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_components.severity_threshold_tiers",
					Description: `Sets a severity label per trigger threshold. If using more than one sub-component, only 1 severityThresholdTiers is allowed. Otherwise, 1 per enum are allowed (for a total of 5 thresholds of increasing severities). Increasing severity must adhere to the logic of the operator.`,
				},
				resource.Attribute{
					Name:        "sub_components.severity_threshold_tiers.severity",
					Description: `Labels the event with a severity tag. Available severity tags are: ` + "`" + `"INFO"` + "`" + `, ` + "`" + `"LOW"` + "`" + `, ` + "`" + `"MEDIUM"` + "`" + `, ` + "`" + `"HIGH"` + "`" + `, ` + "`" + `"SEVERE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_components.severity_threshold_tiers.threshold",
					Description: `Number of logs per search timeframe.`,
				},
				resource.Attribute{
					Name:        "sub_components.columns.field_name",
					Description: `Specify the fields to be included in the notification.`,
				},
				resource.Attribute{
					Name:        "sub_components.columns.regex",
					Description: `Trims the data using regex filters. [Learn more](https://docs.logz.io/user-guide/alerts/regex-filters.html).`,
				},
				resource.Attribute{
					Name:        "sub_components.columns.sort",
					Description: `Specify a single field to sort by. The field cannot be an analyzed field (a field that supports free text search or searching by part of a message, such as the 'message' field). Should be: ` + "`" + `"DESC"` + "`" + `, ` + "`" + `"ASC"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_components.output_should_use_all_fields",
					Description: `If true, the notification output will include entire logs with all of their fields in the sample data.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date and time in UTC when the alert was first created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Email of the user who first created the alert.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date and time in UTC when the alert was last updated.`,
				},
				resource.Attribute{
					Name:        "updated by",
					Description: `Email of the user who last updated the alert.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the event, its significance, and suggested next steps or instructions for the team.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags for filtering alerts and triggered alerts. Can be used in Kibana Discover, dashboards, and more.`,
				},
				resource.Attribute{
					Name:        "search_timeframe_minutes",
					Description: `The time frame for evaluating the log data is a sliding window, with 1 minute granularity.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `If ` + "`" + `true` + "`" + `, the alert is currently active.`,
				},
				resource.Attribute{
					Name:        "notification_emails",
					Description: `Array of email addresses to be notified when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "alert_notification_endpoints",
					Description: `Array of IDs of pre-configured endpoint channels to notify when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "suppress_notifications_minutes",
					Description: `Add a waiting period in minutes to space out notifications. (The alert will still trigger but will not send out notifications during the waiting period.)`,
				},
				resource.Attribute{
					Name:        "output_type",
					Description: `Selects the output format for the alert notification. Can be: ` + "`" + `"JSON"` + "`" + ` or ` + "`" + `"TABLE""` + "`" + ` If the alert has no aggregations/group by fields, JSON offers the option to send full sample logs without selecting specific fields.`,
				},
				resource.Attribute{
					Name:        "correlation_operator",
					Description: `Comma separated string of supported operators. Only applicable when multiple sub-components are in use. Selects a logic for correlating the alert’s sub-components. ` + "`" + `AND` + "`" + ` is currently the only supported operator. When AND is the correlation_operator, both sub-components must meet their triggering criteria for the alert to trigger.`,
				},
				resource.Attribute{
					Name:        "joins",
					Description: `Specifies which group by fields must have the same values to trigger the alert. Joins the group by fields from the first and second sub-components. The key represents the index of the sub component in the array. The fields must be ordered pairs of the group by fields already in use in the ` + "`" + `sub_components.query_string` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_components",
					Description: `Sets the search criteria using a search query, filters, group by aggregations, accounts to search, and trigger conditions.`,
				},
				resource.Attribute{
					Name:        "sub_components.query_string",
					Description: `Provide a Kibana search query written in Lucene syntax. The search query together with the filters select for the relevant logs. Cannot be null - send an asterisk wildcard ` + "`" + `"`,
				},
				resource.Attribute{
					Name:        "sub_components.filter_must",
					Description: `Runs Elasticsearch Bool Query filters on the data (before the search query is applied). The most efficient way to grab the logs you are looking for.`,
				},
				resource.Attribute{
					Name:        "sub_components.filter_must_not",
					Description: `Runs Elasticsearch Bool Query filters on the data (before the search query is applied). The most efficient way to grab the logs you are looking for.`,
				},
				resource.Attribute{
					Name:        "sub_components.group_by_aggregation_fields",
					Description: `Specify 1-3 fields by which to group the results and count them. If you apply a group by operation, the alert returns a count of the results aggregated by unique values.`,
				},
				resource.Attribute{
					Name:        "sub_components.value_aggregation_type",
					Description: `Specifies the aggregation operator. Can be: ` + "`" + `"SUM"` + "`" + `, ` + "`" + `"MIN"` + "`" + `, ` + "`" + `"MAX"` + "`" + `, ` + "`" + `"AVG"` + "`" + `, ` + "`" + `"COUNT"` + "`" + `, ` + "`" + `"UNIQUE_COUNT"` + "`" + `, ` + "`" + `"NONE"` + "`" + `. If ` + "`" + `"COUNT"` + "`" + ` or ` + "`" + `"NONE"` + "`" + `, ` + "`" + `value_aggregation_field` + "`" + ` must be null, and ` + "`" + `group_by_aggregation_fields` + "`" + ` fields must not be empty. If any other operator type (other than ` + "`" + `"NONE"` + "`" + ` or ` + "`" + `"COUNT"` + "`" + `), ` + "`" + `value_aggregation_field` + "`" + ` must not be null.`,
				},
				resource.Attribute{
					Name:        "sub_components.value_aggregation_field",
					Description: `Selects the field on which to run the aggregation for the trigger condition. Cannot be a field already in use for ` + "`" + `group_by_aggregation_fields` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_components.should_query_on_all_accounts",
					Description: `Defaults to true. Only applicable when the alert is run from the main account. If true, the alert runs on the main account and all associated searchable sub accounts. If false, specify relevant account IDs for the alert to monitor using the ` + "`" + `account_ids_to_query_on` + "`" + ` field.`,
				},
				resource.Attribute{
					Name:        "sub_components.account_ids_to_query_on",
					Description: `Specify Account IDs to select which accounts the alert should monitor. The alert will be checked only on these accounts.`,
				},
				resource.Attribute{
					Name:        "sub_components.operation",
					Description: `Specifies the operator for evaluating the results. Can be: ` + "`" + `"LESS_THAN"` + "`" + `, ` + "`" + `"GREATER_THAN"` + "`" + `, ` + "`" + `"LESS_THAN_OR_EQUALS"` + "`" + `, ` + "`" + `"GREATER_THAN_OR_EQUALS"` + "`" + `, ` + "`" + `"EQUALS"` + "`" + `, ` + "`" + `"NOT_EQUALS"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_components.severity_threshold_tiers",
					Description: `Sets a severity label per trigger threshold. If using more than one sub-component, only 1 severityThresholdTiers is allowed. Otherwise, 1 per enum are allowed (for a total of 5 thresholds of increasing severities). Increasing severity must adhere to the logic of the operator.`,
				},
				resource.Attribute{
					Name:        "sub_components.severity_threshold_tiers.severity",
					Description: `Labels the event with a severity tag. Available severity tags are: ` + "`" + `"INFO"` + "`" + `, ` + "`" + `"LOW"` + "`" + `, ` + "`" + `"MEDIUM"` + "`" + `, ` + "`" + `"HIGH"` + "`" + `, ` + "`" + `"SEVERE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_components.severity_threshold_tiers.threshold",
					Description: `Number of logs per search timeframe.`,
				},
				resource.Attribute{
					Name:        "sub_components.columns.field_name",
					Description: `Specify the fields to be included in the notification.`,
				},
				resource.Attribute{
					Name:        "sub_components.columns.regex",
					Description: `Trims the data using regex filters. [Learn more](https://docs.logz.io/user-guide/alerts/regex-filters.html).`,
				},
				resource.Attribute{
					Name:        "sub_components.columns.sort",
					Description: `Specify a single field to sort by. The field cannot be an analyzed field (a field that supports free text search or searching by part of a message, such as the 'message' field). Should be: ` + "`" + `"DESC"` + "`" + `, ` + "`" + `"ASC"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sub_components.output_should_use_all_fields",
					Description: `If true, the notification output will include entire logs with all of their fields in the sample data.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_archive_logs",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "archive_id",
					Description: `(String) Archive ID in the Logz.io database. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(String) Specifies the storage provider. If ` + "`" + `S3` + "`" + `, the ` + "`" + `amazon_s3_storage_settings` + "`" + ` are relevant. If ` + "`" + `BLOB` + "`" + `, the ` + "`" + `azure_blob_storage_settings` + "`" + ` are relevant.`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings",
					Description: `(Object) Applicable settings when the ` + "`" + `storage_type` + "`" + ` is ` + "`" + `S3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "azure_blob_storage_settings",
					Description: `(Object) Applicable settings when the ` + "`" + `storage_type` + "`" + ` is ` + "`" + `BLOB` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Boolean) If ` + "`" + `true` + "`" + `, archiving is currently enabled.`,
				},
				resource.Attribute{
					Name:        "compressed",
					Description: `(Boolean) If ` + "`" + `true` + "`" + `, logs are compressed before they are archived.`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings.credentials_type",
					Description: `(String) Specifies which credentials will be used for authentication.`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings.path",
					Description: `(String) Specify a path to the`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings.s3_secret_credentials",
					Description: `(Object) Applicable settings when the ` + "`" + `credentials_type` + "`" + ` is ` + "`" + `KEYS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings.s3_iam_credentials_arn",
					Description: `(String) Amazon Resource Name (ARN) to uniquely identify the S3 bucket.`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings.s3_external_id",
					Description: `(String) The external id that gives Logz.io access to your S3 bucket.`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings.s3_secret_credentials.access_key",
					Description: `(String) AWS access key.`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings.s3_secret_credentials.secret_key",
					Description: `(String) AWS secret key.`,
				},
				resource.Attribute{
					Name:        "azure_blob_storage_settings.tenant_id",
					Description: `(String) Azure Directory (tenant) ID. The Tenant ID of the AD app.`,
				},
				resource.Attribute{
					Name:        "azure_blob_storage_settings.client_id",
					Description: `(String) Azure application (client) ID. The Client ID of the AD app.`,
				},
				resource.Attribute{
					Name:        "azure_blob_storage_settings.client_secret",
					Description: `(String) Azure client secret.`,
				},
				resource.Attribute{
					Name:        "azure_blob_storage_settings.account_name",
					Description: `(String) Azure Storage account name.`,
				},
				resource.Attribute{
					Name:        "azure_blob_storage_settings.container_name",
					Description: `(String) Name of the container in the Storage account.`,
				},
				resource.Attribute{
					Name:        "azure_blob_storage_settingspath",
					Description: `(String) Optional virtual sub-folder specifiying a path within the container.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "storage_type",
					Description: `(String) Specifies the storage provider. If ` + "`" + `S3` + "`" + `, the ` + "`" + `amazon_s3_storage_settings` + "`" + ` are relevant. If ` + "`" + `BLOB` + "`" + `, the ` + "`" + `azure_blob_storage_settings` + "`" + ` are relevant.`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings",
					Description: `(Object) Applicable settings when the ` + "`" + `storage_type` + "`" + ` is ` + "`" + `S3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "azure_blob_storage_settings",
					Description: `(Object) Applicable settings when the ` + "`" + `storage_type` + "`" + ` is ` + "`" + `BLOB` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Boolean) If ` + "`" + `true` + "`" + `, archiving is currently enabled.`,
				},
				resource.Attribute{
					Name:        "compressed",
					Description: `(Boolean) If ` + "`" + `true` + "`" + `, logs are compressed before they are archived.`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings.credentials_type",
					Description: `(String) Specifies which credentials will be used for authentication.`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings.path",
					Description: `(String) Specify a path to the`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings.s3_secret_credentials",
					Description: `(Object) Applicable settings when the ` + "`" + `credentials_type` + "`" + ` is ` + "`" + `KEYS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings.s3_iam_credentials_arn",
					Description: `(String) Amazon Resource Name (ARN) to uniquely identify the S3 bucket.`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings.s3_external_id",
					Description: `(String) The external id that gives Logz.io access to your S3 bucket.`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings.s3_secret_credentials.access_key",
					Description: `(String) AWS access key.`,
				},
				resource.Attribute{
					Name:        "amazon_s3_storage_settings.s3_secret_credentials.secret_key",
					Description: `(String) AWS secret key.`,
				},
				resource.Attribute{
					Name:        "azure_blob_storage_settings.tenant_id",
					Description: `(String) Azure Directory (tenant) ID. The Tenant ID of the AD app.`,
				},
				resource.Attribute{
					Name:        "azure_blob_storage_settings.client_id",
					Description: `(String) Azure application (client) ID. The Client ID of the AD app.`,
				},
				resource.Attribute{
					Name:        "azure_blob_storage_settings.client_secret",
					Description: `(String) Azure client secret.`,
				},
				resource.Attribute{
					Name:        "azure_blob_storage_settings.account_name",
					Description: `(String) Azure Storage account name.`,
				},
				resource.Attribute{
					Name:        "azure_blob_storage_settings.container_name",
					Description: `(String) Name of the container in the Storage account.`,
				},
				resource.Attribute{
					Name:        "azure_blob_storage_settingspath",
					Description: `(String) Optional virtual sub-folder specifiying a path within the container.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_authentication_groups",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "authentication_group",
					Description: `(Block List) Details for the authentication groups. #### Nested schema for ` + "`" + `authentication_group` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(String) Name of authentication group.`,
				},
				resource.Attribute{
					Name:        "user_role",
					Description: `(String) User role for that group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_drop_filter",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "field_conditions",
					Description: `Filters for an exact match of a field:value pair.`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `Filters for the [log type](https://docs.logz.io/user-guide/log-shipping/built-in-log-types.html). Emit or leave empty if you want this filter to apply to all types.`,
				},
				resource.Attribute{
					Name:        "drop_filter_id",
					Description: `Drop filter ID in the Logz.io database. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `If true, the drop filter is active and logs that match the filter are dropped before indexing. If false, the drop filter is disabled.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "active",
					Description: `If true, the drop filter is active and logs that match the filter are dropped before indexing. If false, the drop filter is disabled.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_endpoint",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the notification endpoint. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "endpoint_type",
					Description: `Specifies the endpoint resource type: ` + "`" + `custom` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `bigpanda` + "`" + `, ` + "`" + `datadog` + "`" + `, ` + "`" + `victorops` + "`" + `, ` + "`" + `opsgenie` + "`" + `, ` + "`" + `servicenow` + "`" + `, ` + "`" + `microsoftteams` + "`" + `. Use the appropriate parameters for your selected endpoint type.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Name of the endpoint.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Detailed description of the endpoint. ## Endpoints used`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_type",
					Description: `Specifies the endpoint resource type: ` + "`" + `custom` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `bigpanda` + "`" + `, ` + "`" + `datadog` + "`" + `, ` + "`" + `victorops` + "`" + `, ` + "`" + `opsgenie` + "`" + `, ` + "`" + `servicenow` + "`" + `, ` + "`" + `microsoftteams` + "`" + `. Use the appropriate parameters for your selected endpoint type.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Name of the endpoint.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Detailed description of the endpoint. ## Endpoints used`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_log_shipping_token",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(String) Descriptive name for this token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Boolean) To enable this token, true. To disable, false.`,
				},
				resource.Attribute{
					Name:        "token_id",
					Description: `(Integer) The token's ID.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(String) The token itself.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `(Integer) Unix timestamp of when this token was last updated.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `(String) Email address of the last user to update this token.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Integer) Unix timestamp of when this token was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `(String) Email address of the user who created this token.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_restore_logs",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "restore_operation_id",
					Description: `(Integer) ID of the restore operation in Logz.io. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(String) Name of the restored account.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Integer) UNIX timestamp in milliseconds specifying the earliest logs to be restored.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Integer) UNIX timestamp in milliseconds specifying the latest logs to be restored.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Integer) ID of the restored account in Logz.io.`,
				},
				resource.Attribute{
					Name:        "restored_volume_gb",
					Description: `(Float) Volume of data restored so far. If the restore operation is still in progress, this will be continuously updated.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(String) Returns the current status of the restored account. See [documentation](https://docs.logz.io/api/#operation/getRestoreRequestByIdApi) for more info about the possible statuses and their meaning.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Integer) Timestamp when the restore process was created and entered the queue. Since only one account can be restored at a time, the process may not initiate immediately.`,
				},
				resource.Attribute{
					Name:        "started_at",
					Description: `(Integer) UNIX timestamp in milliseconds when the restore process initiated.`,
				},
				resource.Attribute{
					Name:        "finished_at",
					Description: `(Integer) UNIX timestamp in milliseconds when the restore process completed.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `(Integer) UNIX timestamp in milliseconds specifying when the account is due to expire. Restored accounts expire automatically after a number of days, as specified in the account's terms.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(String) Name of the restored account.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Integer) UNIX timestamp in milliseconds specifying the earliest logs to be restored.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Integer) UNIX timestamp in milliseconds specifying the latest logs to be restored.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Integer) ID of the restored account in Logz.io.`,
				},
				resource.Attribute{
					Name:        "restored_volume_gb",
					Description: `(Float) Volume of data restored so far. If the restore operation is still in progress, this will be continuously updated.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(String) Returns the current status of the restored account. See [documentation](https://docs.logz.io/api/#operation/getRestoreRequestByIdApi) for more info about the possible statuses and their meaning.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Integer) Timestamp when the restore process was created and entered the queue. Since only one account can be restored at a time, the process may not initiate immediately.`,
				},
				resource.Attribute{
					Name:        "started_at",
					Description: `(Integer) UNIX timestamp in milliseconds when the restore process initiated.`,
				},
				resource.Attribute{
					Name:        "finished_at",
					Description: `(Integer) UNIX timestamp in milliseconds when the restore process completed.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `(Integer) UNIX timestamp in milliseconds specifying when the account is due to expire. Restored accounts expire automatically after a number of days, as specified in the account's terms.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_subaccount",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `ID of the subaccount. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(String) Email address of an existing admin user on the main account which will also become the admin of the subaccount being created.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(String) Name of the subaccount.`,
				},
				resource.Attribute{
					Name:        "max_daily_gb",
					Description: `(Float) Maximum daily log volume that the subaccount can index, in GB.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `(Integer) Number of days that log data is retained.`,
				},
				resource.Attribute{
					Name:        "sharing_objects_accounts",
					Description: `(List) IDs of accounts that can access the account's Kibana objects. Can be an empty array.`,
				},
				resource.Attribute{
					Name:        "searchable",
					Description: `(Boolean) False by default. Determines if other accounts can search logs indexed by the subaccount.`,
				},
				resource.Attribute{
					Name:        "accessible",
					Description: `(Boolean) False by default. Determines if users of main account can access the subaccount.`,
				},
				resource.Attribute{
					Name:        "doc_size_setting",
					Description: `(Boolean) False by default. If enabled, Logz.io adds a ` + "`" + `LogSize` + "`" + ` field to record the size of the log line in bytes, for the purpose of managing account utilization. [Learn more about managing account usage](https://docs.logz.io/user-guide/accounts/manage-account-usage.html#enabling-account-utilization-metrics-and-log-size)`,
				},
				resource.Attribute{
					Name:        "utilization_enabled",
					Description: `(Boolean) If enabled, account utilization metrics and expected utilization at the current indexing rate are measured at a pre-defined sampling rate. Useful for managing account utilization and avoiding running out of capacity. [Learn more about managing account capacity](https://docs.logz.io/user-guide/accounts/manage-account-usage.html).`,
				},
				resource.Attribute{
					Name:        "frequency_minutes",
					Description: `(Int) Determines the sampling rate in minutes of the utilization.`,
				},
				resource.Attribute{
					Name:        "flexible",
					Description: `(Boolean) Defaults to false. Whether the sub account that created is flexible or not. Can be set to flexible only if the main account is flexible.`,
				},
				resource.Attribute{
					Name:        "reserved_daily_gb",
					Description: `(Float) The maximum volume of data that an account can index per calendar day. Depends on ` + "`" + `flexible` + "`" + `. For further info see [the docs](https://docs.logz.io/api/#operation/createTimeBasedAccount). ## Endpoints used`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user in the Logz.io platform.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username credential. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "fullname",
					Description: `First and last name of the user.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `For User access, ` + "`" + `2` + "`" + `. For Admin access, ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `If ` + "`" + `true` + "`" + `, the user is active, if ` + "`" + `false` + "`" + `, the user is suspended.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Logz.io account ID. ## Endpoints used`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"logzio_alert":                 0,
		"logzio_alert_v2":              1,
		"logzio_archive_logs":          2,
		"logzio_authentication_groups": 3,
		"logzio_drop_filter":           4,
		"logzio_endpoint":              5,
		"logzio_log_shipping_token":    6,
		"logzio_restore_logs":          7,
		"logzio_subaccount":            8,
		"logzio_user":                  9,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
