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
			Type:             "logzio_endpoint",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_type",
					Description: `(Required) Specifies the endpoint resource type: ` + "`" + `custom` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `pager_duty` + "`" + `, ` + "`" + `big_panda` + "`" + `, ` + "`" + `data_dog` + "`" + `, ` + "`" + `victorops` + "`" + `. Use the appropriate parameters for your selected endpoint type.`,
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
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `pager_duty` + "`" + `. Manages a webhook to PagerDuty.`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `API key generated from PagerDuty for the purpose of the integration.`,
				},
				resource.Attribute{
					Name:        "big_panda",
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `big_panda` + "`" + `. Manages a webhook to BigPanda.`,
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
					Description: `(Optional) Relevant when ` + "`" + `endpoint_type` + "`" + ` is ` + "`" + `data_dog` + "`" + `. Manages a webhook to Datadog.`,
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
					Description: `(Required) IDs of accounts that can access the account's Kibana objects.`,
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
					Description: `Log shipping token for the subaccount. [Learn more](https://docs.logz.io/user-guide/tokens/log-shipping-tokens/) ## Endpoints used`,
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

		"logzio_alert":      0,
		"logzio_endpoint":   1,
		"logzio_subaccount": 2,
		"logzio_user":       3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
