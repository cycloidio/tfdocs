package logzio

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

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
					Description: `Runs Elasticsearch Bool Query ` + "`" + `must` + "`" + ` filters on the data (before the search query is applied). The most efficient way to grab the logs you are looking for.`,
				},
				resource.Attribute{
					Name:        "sub_components.filter_must_not",
					Description: `Runs Elasticsearch Bool Query ` + "`" + `must_not` + "`" + ` filters on the data (before the search query is applied). The most efficient way to grab the logs you are looking for.`,
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
					Description: `Runs Elasticsearch Bool Query ` + "`" + `must` + "`" + ` filters on the data (before the search query is applied). The most efficient way to grab the logs you are looking for.`,
				},
				resource.Attribute{
					Name:        "sub_components.filter_must_not",
					Description: `Runs Elasticsearch Bool Query ` + "`" + `must_not` + "`" + ` filters on the data (before the search query is applied). The most efficient way to grab the logs you are looking for.`,
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
					Description: `(String) Specifies the storage provider.`,
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
					Name:        "aws_credentials_type",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `S3` + "`" + `. Specifies which credentials will be used for authentication.`,
				},
				resource.Attribute{
					Name:        "aws_s3_path",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `S3` + "`" + `. Specify a path to the`,
				},
				resource.Attribute{
					Name:        "aws_s3_iam_credentials_arn",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `S3` + "`" + `. Amazon Resource Name (ARN) to uniquely identify the S3 bucket.`,
				},
				resource.Attribute{
					Name:        "aws_access_key",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `S3` + "`" + `. AWS access key.`,
				},
				resource.Attribute{
					Name:        "azure_tenant_id",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `BLOB` + "`" + `. Azure Directory (tenant) ID. The Tenant ID of the AD app.`,
				},
				resource.Attribute{
					Name:        "azure_client_id",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `BLOB` + "`" + `. Azure application (client) ID. The Client ID of the AD app.`,
				},
				resource.Attribute{
					Name:        "azure_account_name",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `BLOB` + "`" + `. Azure Storage account name.`,
				},
				resource.Attribute{
					Name:        "azure_container_name",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `BLOB` + "`" + `. Name of the container in the Storage account.`,
				},
				resource.Attribute{
					Name:        "azure_blob_path",
					Description: `(String) Optional virtual sub-folder specifying a path within the container.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "storage_type",
					Description: `(String) Specifies the storage provider.`,
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
					Name:        "aws_credentials_type",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `S3` + "`" + `. Specifies which credentials will be used for authentication.`,
				},
				resource.Attribute{
					Name:        "aws_s3_path",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `S3` + "`" + `. Specify a path to the`,
				},
				resource.Attribute{
					Name:        "aws_s3_iam_credentials_arn",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `S3` + "`" + `. Amazon Resource Name (ARN) to uniquely identify the S3 bucket.`,
				},
				resource.Attribute{
					Name:        "aws_access_key",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `S3` + "`" + `. AWS access key.`,
				},
				resource.Attribute{
					Name:        "azure_tenant_id",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `BLOB` + "`" + `. Azure Directory (tenant) ID. The Tenant ID of the AD app.`,
				},
				resource.Attribute{
					Name:        "azure_client_id",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `BLOB` + "`" + `. Azure application (client) ID. The Client ID of the AD app.`,
				},
				resource.Attribute{
					Name:        "azure_account_name",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `BLOB` + "`" + `. Azure Storage account name.`,
				},
				resource.Attribute{
					Name:        "azure_container_name",
					Description: `(String) Applicable when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `BLOB` + "`" + `. Name of the container in the Storage account.`,
				},
				resource.Attribute{
					Name:        "azure_blob_path",
					Description: `(String) Optional virtual sub-folder specifying a path within the container.`,
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
					Description: `Detailed description of the endpoint.`,
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
					Description: `Detailed description of the endpoint.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_grafana_dashboard",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dashboard_uid",
					Description: `The unique identifier (uid) of the dashboard. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Dashboard url.`,
				},
				resource.Attribute{
					Name:        "folder_uid",
					Description: `The unique identifier (uid) of a folder to store your dashboard.`,
				},
				resource.Attribute{
					Name:        "dashboard_json",
					Description: `The complete dashboard model, to create a new dashboard, in a JSON format.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_kibana_object",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "object_id",
					Description: `(String) The id of the Kibana Object.`,
				},
				resource.Attribute{
					Name:        "object_type",
					Description: `(String) The type of the Kibana Object. Can be one of the following: ` + "`" + `search` + "`" + `, ` + "`" + `dashboard` + "`" + `, ` + "`" + `visualization` + "`" + `. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "kibana_version",
					Description: `(String) The version of Kibana used at the time of export.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(String) Exported Kibana objects.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "kibana_version",
					Description: `(String) The version of Kibana used at the time of export.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(String) Exported Kibana objects.`,
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
					Description: `(String) Descriptive name for this log shipping token.`,
				},
				resource.Attribute{
					Name:        "token_id",
					Description: `(Integer) The log shipping token's ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Boolean) To enable this log shipping token, true. To disable, false.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(String) The log shipping token itself.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `(Integer) Unix timestamp of when this log shipping token was last updated.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `(String) Email address of the last user to update this log shipping token.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Integer) Unix timestamp of when this log shipping token was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `(String) Email address of the user who created this log shipping token.`,
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
			Type:             "logzio_s3_fetcher",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fetcher_id",
					Description: `ID of the S3 Fetcher. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "aws_access_key",
					Description: `(String) AWS S3 bucket access key. Not applicable if you chose to authenticate with ` + "`" + `aws_arn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aws_arn",
					Description: `(String) Amazon Resource Name (ARN) to uniquely identify the S3 bucket. Not applicable if you choose to authenticate with AWS keys (access key & secret key).`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(String) AWS S3 bucket name.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Boolean) If true, the S3 bucket connector is active and logs are being fetched to Logz.io. If false, the connector is disabled.`,
				},
				resource.Attribute{
					Name:        "aws_region",
					Description: `(String) Bucket's region. Allowed values: ` + "`" + `US_EAST_1` + "`" + `, ` + "`" + `US_EAST_2` + "`" + `, ` + "`" + `US_WEST_1` + "`" + `, ` + "`" + `US_WEST_2` + "`" + `, ` + "`" + `EU_WEST_1` + "`" + `, ` + "`" + `EU_WEST_2` + "`" + `, ` + "`" + `EU_WEST_3` + "`" + `, ` + "`" + `EU_CENTRAL_1` + "`" + `, ` + "`" + `AP_NORTHEAST_1` + "`" + `, ` + "`" + `AP_NORTHEAST_2` + "`" + `, ` + "`" + `AP_SOUTHEAST_1` + "`" + `, ` + "`" + `AP_SOUTHEAST_2` + "`" + `, ` + "`" + `SA_EAST_1` + "`" + `, ` + "`" + `AP_SOUTH_1` + "`" + `, ` + "`" + `CA_CENTRAL_1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logs_type",
					Description: `(String) Specifies the log type being sent to Logz.io. Determines the parsing pipeline used to parse and map the logs. [Learn more about parsing options supported by Logz.io](https://docs.logz.io/user-guide/log-shipping/built-in-log-types.html). Allowed values: ` + "`" + `elb` + "`" + `, ` + "`" + `vpcflow` + "`" + `, ` + "`" + `S3Access` + "`" + `, ` + "`" + `cloudfront` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(String) Prefix of the AWS S3 bucket.`,
				},
				resource.Attribute{
					Name:        "add_s3_object_key_as_log_field",
					Description: `(Boolean) If ` + "`" + `true` + "`" + `, enriches logs with a new field detailing the S3 object key.`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Float) The maximum volume of data that an account can index per calendar day. Depends on ` + "`" + `flexible` + "`" + `. For further info see [the docs](https://docs.logz.io/api/#operation/createTimeBasedAccount).`,
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
					Name:        "role",
					Description: `User role. Can be ` + "`" + `USER_ROLE_READONLY` + "`" + `, ` + "`" + `USER_ROLE_REGULAR` + "`" + ` or ` + "`" + `USER_ROLE_ACCOUNT_ADMIN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `If ` + "`" + `true` + "`" + `, the user is active, if ` + "`" + `false` + "`" + `, the user is suspended.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Logz.io account ID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"logzio_alert_v2":              0,
		"logzio_archive_logs":          1,
		"logzio_authentication_groups": 2,
		"logzio_drop_filter":           3,
		"logzio_endpoint":              4,
		"logzio_grafana_dashboard":     5,
		"logzio_kibana_object":         6,
		"logzio_log_shipping_token":    7,
		"logzio_restore_logs":          8,
		"logzio_s3_fetcher":            9,
		"logzio_subaccount":            10,
		"logzio_user":                  11,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
