package logzio

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "logzio_alert",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "title",
					Description: `(Required) Alert title.`,
				},
				resource.Attribute{
					Name:        "search_timeframe_minutes",
					Description: `(Required) The time frame for evaluating the log data is a sliding window, with 1 minute granularity.`,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `Defaults to ` + "`" + `GREATER_THAN` + "`" + `. Specifies the operator for evaluating the ` + "`" + `severity_threshold_tiers` + "`" + ` results. Enum: ` + "`" + `LESS_THAN` + "`" + `, ` + "`" + `GREATER_THAN` + "`" + `, ` + "`" + `LESS_THAN_OR_EQUALS` + "`" + `, ` + "`" + `GREATER_THAN_OR_EQUALS` + "`" + `, ` + "`" + `EQUALS` + "`" + `, ` + "`" + `NOT_EQUALS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity_threshold_tiers",
					Description: `(Required) Set as many as 5 thresholds, each with its own severity level.`,
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
					Description: `(Required but can be sent empty) Add IDs of endpoint channels to automatically receive notifications with sample data when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "notification_emails",
					Description: `(Required but can be sent empty) Add email addresses to automatically receive notifications with sample data when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the event, its significance, and suggested next steps or instructions for the team.`,
				},
				resource.Attribute{
					Name:        "query_string",
					Description: `(Required) Search query in Lucene syntax. You can combine filters and a search query to specify the logs you are looking for. You can combine filters and a search query to specify the logs you are looking for.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) You can use ` + "`" + `must` + "`" + ` and ` + "`" + `must_not` + "`" + ` filters. Filters are more efficient compared to a query, so it's recommended to opt for a filter over a ` + "`" + `query_string` + "`" + `, where possible.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags for filtering alerts and triggered alerts. Can be used in Kibana Discover, Kibana dashboards, and more.`,
				},
				resource.Attribute{
					Name:        "group_by_aggregation_fields",
					Description: `(Optional) Specify 1-3 fields by which to group the results and count them. If you apply a group by operation, the alert returns a count of the results aggregated by unique values.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) True by default. If ` + "`" + `true` + "`" + `, the alert is currently active.`,
				},
				resource.Attribute{
					Name:        "last_triggered_at",
					Description: `(Optional) Date and time in UTC when the alert last triggered.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `(Optional) Date and time in UTC when the alert was last updated.`,
				},
				resource.Attribute{
					Name:        "suppress_notifications_minutes",
					Description: `(Optional) Add a waiting period in minutes to space out notifications. (The alert will still trigger but will not send out notifications during the waiting period.)`,
				},
				resource.Attribute{
					Name:        "value_aggregation_field",
					Description: `(Optional) Specify the field on which to run the aggregation for the trigger condition.`,
				},
				resource.Attribute{
					Name:        "value_aggregation_type",
					Description: `(Required) Specifies the aggregation operator. Can be: ` + "`" + `SUM` + "`" + `, ` + "`" + `MIN` + "`" + `, ` + "`" + `MAX` + "`" + `, ` + "`" + `AVG` + "`" + `, ` + "`" + `COUNT` + "`" + `, ` + "`" + `UNIQUE_COUNT` + "`" + `, ` + "`" + `NONE` + "`" + `. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Logz.io alert ID.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date and time in UTC when the alert was first created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Email of the user who first created the alert.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Logz.io alert ID.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date and time in UTC when the alert was first created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Email of the user who first created the alert.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_alert_v2",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "title",
					Description: `(String) Alert title.`,
				},
				resource.Attribute{
					Name:        "search_timeframe_minutes",
					Description: `(Integer) The time frame for evaluating the log data is a sliding window, with 1 minute granularity.`,
				},
				resource.Attribute{
					Name:        "sub_components",
					Description: `(Block list) Sets the search criteria using a search query, filters, group by aggregations, accounts to search, and trigger conditions. See below for`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(String) A description of the event, its significance, and suggested next steps or instructions for the team.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(String list) Tags for filtering alerts and triggered alerts. Can be used in Kibana Discover, dashboards, and more.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Boolean) True by default. If ` + "`" + `true` + "`" + `, the alert is currently active.`,
				},
				resource.Attribute{
					Name:        "notification_emails",
					Description: `(String list) Array of email addresses to be notified when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "alert_notification_endpoints",
					Description: `(Integer list) Array of IDs of pre-configured endpoint channels to notify when the alert triggers.`,
				},
				resource.Attribute{
					Name:        "suppress_notifications_minutes",
					Description: `(Integer) Defaults to 5. Add a waiting period in minutes to space out notifications. (The alert will still trigger but will not send out notifications during the waiting period.)`,
				},
				resource.Attribute{
					Name:        "output_type",
					Description: `(String) Selects the output format for the alert notification. Can be: ` + "`" + `"JSON"` + "`" + ` or ` + "`" + `"TABLE""` + "`" + ` If the alert has no aggregations/group by fields, JSON offers the option to send full sample logs without selecting specific fields.`,
				},
				resource.Attribute{
					Name:        "correlation_operator",
					Description: `(String) Comma separated string of supported operators. Only applicable when multiple sub-components are in use. Selects a logic for correlating the alert’s sub-components. ` + "`" + `AND` + "`" + ` is currently the only supported operator. When AND is the correlation_operator, both sub-components must meet their triggering criteria for the alert to trigger.`,
				},
				resource.Attribute{
					Name:        "joins",
					Description: `(Map list) Specifies which group by fields must have the same values to trigger the alert. Joins the group by fields from the first and second sub-components. The key represents the index of the sub component in the array. The fields must be ordered pairs of the group by fields already in use in the ` + "`" + `sub_components.query_string` + "`" + `. #### Nested schema for ` + "`" + `sub_components` + "`" + `: ##### Required:`,
				},
				resource.Attribute{
					Name:        "query_string",
					Description: `(String) Provide a Kibana search query written in Lucene syntax. The search query together with the filters select for the relevant logs. Cannot be null - send an asterisk wildcard ` + "`" + `"`,
				},
				resource.Attribute{
					Name:        "value_aggregation_type",
					Description: `(String) Specifies the aggregation operator. Can be: ` + "`" + `"SUM"` + "`" + `, ` + "`" + `"MIN"` + "`" + `, ` + "`" + `"MAX"` + "`" + `, ` + "`" + `"AVG"` + "`" + `, ` + "`" + `"COUNT"` + "`" + `, ` + "`" + `"UNIQUE_COUNT"` + "`" + `, ` + "`" + `"NONE"` + "`" + `. If ` + "`" + `"COUNT"` + "`" + ` or ` + "`" + `"NONE"` + "`" + `, ` + "`" + `value_aggregation_field` + "`" + ` must be null, and ` + "`" + `group_by_aggregation_fields` + "`" + ` fields must not be empty. If any other operator type (other than ` + "`" + `"NONE"` + "`" + ` or ` + "`" + `"COUNT"` + "`" + `), ` + "`" + `value_aggregation_field` + "`" + ` must not be null.`,
				},
				resource.Attribute{
					Name:        "severity_threshold_tiers",
					Description: `(Block) Sets a severity label per trigger threshold. If using more than one sub-component, only 1 severityThresholdTiers is allowed. Otherwise, 1 per enum are allowed (for a total of 5 thresholds of increasing severities). Increasing severity must adhere to the logic of the operator. See below for`,
				},
				resource.Attribute{
					Name:        "filter_must_not",
					Description: `(Map) Runs Elasticsearch Bool Query filters on the data (before the search query is applied). The most efficient way to grab the logs you are looking for.`,
				},
				resource.Attribute{
					Name:        "group_by_aggregation_fields",
					Description: `(String list) Specify 1-3 fields by which to group the results and count them. If you apply a group by operation, the alert returns a count of the results aggregated by unique values.`,
				},
				resource.Attribute{
					Name:        "value_aggregation_field",
					Description: `(String) Selects the field on which to run the aggregation for the trigger condition. Cannot be a field already in use for ` + "`" + `group_by_aggregation_fields` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "should_query_on_all_accounts",
					Description: `(Boolean) Defaults to true. Only applicable when the alert is run from the main account. If true, the alert runs on the main account and all associated searchable sub accounts. If false, specify relevant account IDs for the alert to monitor using the ` + "`" + `account_ids_to_query_on` + "`" + ` field.`,
				},
				resource.Attribute{
					Name:        "account_ids_to_query_on",
					Description: `(Integer list) Specify Account IDs to select which accounts the alert should monitor. The alert will be checked only on these accounts.`,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `(String) Specifies the operator for evaluating the results. Can be: ` + "`" + `"LESS_THAN"` + "`" + `, ` + "`" + `"GREATER_THAN"` + "`" + `, ` + "`" + `"LESS_THAN_OR_EQUALS"` + "`" + `, ` + "`" + `"GREATER_THAN_OR_EQUALS"` + "`" + `, ` + "`" + `"EQUALS"` + "`" + `, ` + "`" + `"NOT_EQUALS"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "columns",
					Description: `(Block list) See below for`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(String) Labels the event with a severity tag. Available severity tags are: ` + "`" + `"INFO"` + "`" + `, ` + "`" + `"LOW"` + "`" + `, ` + "`" + `"MEDIUM"` + "`" + `, ` + "`" + `"HIGH"` + "`" + `, ` + "`" + `"SEVERE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Integer) Number of logs per search timeframe. #### Nested schema for ` + "`" + `sub_components.columns` + "`" + `: ##### Optional:`,
				},
				resource.Attribute{
					Name:        "field_name",
					Description: `(String) Specify the fields to be included in the notification.`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(String) Trims the data using regex filters. [Learn more](https://docs.logz.io/user-guide/alerts/regex-filters.html).`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(String) Specify a single field to sort by. The field cannot be an analyzed field (a field that supports free text search or searching by part of a message, such as the 'message' field). Should be: ` + "`" + `"DESC"` + "`" + `, ` + "`" + `"ASC"` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_drop_filter",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "field_conditions",
					Description: `(Block list) Filters for an exact match of a field:value pair.`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `(String) Filters for the [log type](https://docs.logz.io/user-guide/log-shipping/built-in-log-types.html). Emit or leave empty if you want this filter to apply to all types.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Boolean) If true, the drop filter is active and logs that match the filter are dropped before indexing. If false, the drop filter is disabled.`,
				},
				resource.Attribute{
					Name:        "field_name",
					Description: `(String) Exact field name in your Kibana mapping for the selected ` + "`" + `log_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Object) Exact field value. The filter looks for an exact value match of the entire object.`,
				},
				resource.Attribute{
					Name:        "drop_filter_id",
					Description: `(String) Drop filter ID in the Logz.io database.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_endpoint",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_type",
					Description: `(Required) Specifies the endpoint resource type: ` + "`" + `custom` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `bigpanda` + "`" + `, ` + "`" + `datadog` + "`" + `, ` + "`" + `victorops` + "`" + `. Use the appropriate parameters for your selected endpoint type.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) Name of the endpoint.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Detailed description of the endpoint.`,
				},
				resource.Attribute{
					Name:        "slack",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `slack` + "`" + `. Manages a webhook to a specific Slack channel.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Slack webhook URL to a specific Slack channel.`,
				},
				resource.Attribute{
					Name:        "pager_duty",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `pagerduty` + "`" + `. Manages a webhook to PagerDuty.`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `API key generated from PagerDuty for the purpose of the integration.`,
				},
				resource.Attribute{
					Name:        "big_panda",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `bigpanda` + "`" + `. Manages a webhook to BigPanda.`,
				},
				resource.Attribute{
					Name:        "api_token",
					Description: `API authentication token from BigPanda.`,
				},
				resource.Attribute{
					Name:        "app_key",
					Description: `Application key from BigPanda.`,
				},
				resource.Attribute{
					Name:        "data_dog",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `datadog` + "`" + `. Manages a webhook to Datadog.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `API key from Datadog.`,
				},
				resource.Attribute{
					Name:        "victorops",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `victorops` + "`" + `. Manages a webhook to VictorOps.`,
				},
				resource.Attribute{
					Name:        "routing_key",
					Description: `Alert routing key from VictorOps.`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `VictorOps REST API ` + "`" + `message_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_api_key",
					Description: `API key from VictorOps.`,
				},
				resource.Attribute{
					Name:        "custom",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `custom` + "`" + `. Manages a custom webhook for your integration of choice.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Specifies the URL destination.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `Selects the HTTP request method.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `Header parameters for the request. Sent as comma-separated key-value pairs.`,
				},
				resource.Attribute{
					Name:        "body_template",
					Description: `JSON object that serves as the template for the message body. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the notification endpoint. ## Endpoints used Logz.io integrates with:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the notification endpoint. ## Endpoints used Logz.io integrates with:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_log_shipping_token",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(String) Descriptive name for this token. ### Optional:`,
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
			Type:             "logzio_subaccount",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email address of an existing admin user on the main account which will also become the admin of the subaccount being created.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Name of the subaccount.`,
				},
				resource.Attribute{
					Name:        "max_daily_gb",
					Description: `(Required) Maximum daily log volume that the subaccount can index, in GB.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `(Required) Number of days that log data is retained.`,
				},
				resource.Attribute{
					Name:        "searchable",
					Description: `(Optional) False by default. Determines if other accounts can search logs indexed by the subaccount.`,
				},
				resource.Attribute{
					Name:        "accessible",
					Description: `(Optional) False by default. Determines if users of main account can access the subaccount.`,
				},
				resource.Attribute{
					Name:        "doc_size_setting",
					Description: `(Optional) False by default. If enabled, Logz.io adds a ` + "`" + `LogSize` + "`" + ` field to record the size of the log line in bytes, for the purpose of managing account utilization. [Learn more about managing account usage](https://docs.logz.io/user-guide/accounts/manage-account-usage.html#enabling-account-utilization-metrics-and-log-size)`,
				},
				resource.Attribute{
					Name:        "sharing_objects_accounts",
					Description: `(Required) IDs of accounts that can access the account's Kibana objects. Can be an empty array.`,
				},
				resource.Attribute{
					Name:        "utilization_settings",
					Description: `(Optional) If enabled, account utilization metrics and expected utilization at the current indexing rate are measured at a pre-defined sampling rate. Useful for managing account utilization and avoiding running out of capacity. [Learn more about managing account capacity](https://docs.logz.io/user-guide/accounts/manage-account-usage.html)`,
				},
				resource.Attribute{
					Name:        "frequencyMinutes",
					Description: `Determines the sampling rate in minutes.`,
				},
				resource.Attribute{
					Name:        "utilizationEnabled",
					Description: `Enables the feature. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `ID of the subaccount.`,
				},
				resource.Attribute{
					Name:        "account_token",
					Description: `Log shipping token for the subaccount. [Learn more](https://docs.logz.io/user-guide/tokens/log-shipping-tokens/)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logzio_user",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fullname",
					Description: `(Required) First and last name of the user.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username credential.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) For User access, ` + "`" + `2` + "`" + `. For Admin access, ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Required) If ` + "`" + `true` + "`" + `, the user is active, if ` + "`" + `false` + "`" + `, the user is suspended.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) Logz.io account ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user in the Logz.io platform. ## Endpoints used`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"logzio_alert":              0,
		"logzio_alert_v2":           1,
		"logzio_drop_filter":        2,
		"logzio_endpoint":           3,
		"logzio_log_shipping_token": 4,
		"logzio_subaccount":         5,
		"logzio_user":               6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
