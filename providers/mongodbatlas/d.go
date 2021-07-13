package mongodbatlas

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_alert_configuration",
			Category:         "Data Sources",
			ShortDescription: `Describes a Alert Configuration.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project where the alert configuration will create.`,
				},
				resource.Attribute{
					Name:        "alert_configuration_id",
					Description: `(Required) Unique identifier for the alert configuration. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Unique identifier of the project that owns this alert configuration.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Timestamp in ISO 8601 date and time format in UTC when this alert configuration was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Timestamp in ISO 8601 date and time format in UTC when this alert configuration was last updated.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If set to true, the alert configuration is enabled. If enabled is not exported it is set to false.`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `The type of event that will trigger an alert. ->`,
				},
				resource.Attribute{
					Name:        "field_name",
					Description: `Name of the field in the target object to match on. | Host alerts | Replica set alerts | Sharded cluster alerts | |:-------------- |:------------- |:------ | | ` + "`" + `TYPE_NAME` + "`" + ` | ` + "`" + `REPLICA_SET_NAME` + "`" + ` | ` + "`" + `CLUSTER_NAME` + "`" + ` | | ` + "`" + `HOSTNAME` + "`" + ` | ` + "`" + `SHARD_NAME` + "`" + ` | ` + "`" + `SHARD_NAME` + "`" + ` | | ` + "`" + `PORT` + "`" + ` | ` + "`" + `CLUSTER_NAME` + "`" + ` | | | ` + "`" + `HOSTNAME_AND_PORT` + "`" + ` | | | | ` + "`" + `REPLICA_SET_NAME` + "`" + ` | | | All other types of alerts do not support matchers.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `If omitted, the configuration is disabled.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `If omitted, the configuration is disabled.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `The operator to test the field’s value. Accepted values are: - ` + "`" + `EQUALS` + "`" + ` - ` + "`" + `NOT_EQUALS` + "`" + ` - ` + "`" + `CONTAINS` + "`" + ` - ` + "`" + `NOT_CONTAINS` + "`" + ` - ` + "`" + `STARTS_WITH` + "`" + ` - ` + "`" + `ENDS_WITH` + "`" + ` - ` + "`" + `REGEX` + "`" + ``,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value to test with the specified operator. If ` + "`" + `field_name` + "`" + ` is set to TYPE_NAME, you can match on the following values: - ` + "`" + `PRIMARY` + "`" + ` - ` + "`" + `SECONDARY` + "`" + ` - ` + "`" + `STANDALONE` + "`" + ` - ` + "`" + `CONFIG` + "`" + ` - ` + "`" + `MONGOS` + "`" + ` ### Metric Threshold The threshold that causes an alert to be triggered. Required if ` + "`" + `event_type_name` + "`" + ` : "OUTSIDE_METRIC_THRESHOLD".`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `Name of the metric to check. The full list of current options is available [here](https://docs.atlas.mongodb.com/reference/alert-host-metrics/#measurement-types)`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Operator to apply when checking the current metric value against the threshold value. Accepted values are: - ` + "`" + `GREATER_THAN` + "`" + ` - ` + "`" + `LESS_THAN` + "`" + ``,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Threshold value outside of which an alert will be triggered.`,
				},
				resource.Attribute{
					Name:        "units",
					Description: `The units for the threshold value. Depends on the type of metric. Accepted values are: - ` + "`" + `RAW` + "`" + ` - ` + "`" + `BITS` + "`" + ` - ` + "`" + `BYTES` + "`" + ` - ` + "`" + `KILOBITS` + "`" + ` - ` + "`" + `KILOBYTES` + "`" + ` - ` + "`" + `MEGABITS` + "`" + ` - ` + "`" + `MEGABYTES` + "`" + ` - ` + "`" + `GIGABITS` + "`" + ` - ` + "`" + `GIGABYTES` + "`" + ` - ` + "`" + `TERABYTES` + "`" + ` - ` + "`" + `PETABYTES` + "`" + ` - ` + "`" + `MILLISECONDS` + "`" + ` - ` + "`" + `SECONDS` + "`" + ` - ` + "`" + `MINUTES` + "`" + ` - ` + "`" + `HOURS` + "`" + ` - ` + "`" + `DAYS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `This must be set to AVERAGE. Atlas computes the current metric value as an average. ### Threshold`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Operator to apply when checking the current metric value against the threshold value. Accepted values are: - ` + "`" + `GREATER_THAN` + "`" + ` - ` + "`" + `LESS_THAN` + "`" + ``,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Threshold value outside of which an alert will be triggered.`,
				},
				resource.Attribute{
					Name:        "units",
					Description: `The units for the threshold value. Depends on the type of metric. Accepted values are: - ` + "`" + `RAW` + "`" + ` - ` + "`" + `BITS` + "`" + ` - ` + "`" + `BYTES` + "`" + ` - ` + "`" + `KILOBITS` + "`" + ` - ` + "`" + `KILOBYTES` + "`" + ` - ` + "`" + `MEGABITS` + "`" + ` - ` + "`" + `MEGABYTES` + "`" + ` - ` + "`" + `GIGABITS` + "`" + ` - ` + "`" + `GIGABYTES` + "`" + ` - ` + "`" + `TERABYTES` + "`" + ` - ` + "`" + `PETABYTES` + "`" + ` - ` + "`" + `MILLISECONDS` + "`" + ` - ` + "`" + `SECONDS` + "`" + ` - ` + "`" + `MINUTES` + "`" + ` - ` + "`" + `HOURS` + "`" + ` - ` + "`" + `DAYS` + "`" + ` ### Notifications Notifications to send when an alert condition is detected.`,
				},
				resource.Attribute{
					Name:        "api_token",
					Description: `Slack API token. Required for the SLACK notifications type. If the token later becomes invalid, Atlas sends an email to the project owner and eventually removes the token.`,
				},
				resource.Attribute{
					Name:        "channel_name",
					Description: `Slack channel name. Required for the SLACK notifications type.`,
				},
				resource.Attribute{
					Name:        "datadog_api_key",
					Description: `Datadog API Key. Found in the Datadog dashboard. Required for the DATADOG notifications type.`,
				},
				resource.Attribute{
					Name:        "datadog_region",
					Description: `Region that indicates which API URL to use. Accepted regions are: ` + "`" + `US` + "`" + `, ` + "`" + `EU` + "`" + `. The default Datadog region is US.`,
				},
				resource.Attribute{
					Name:        "delay_min",
					Description: `Number of minutes to wait after an alert condition is detected before sending out the first notification.`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `Email address to which alert notifications are sent. Required for the EMAIL notifications type.`,
				},
				resource.Attribute{
					Name:        "email_enabled",
					Description: `Flag indicating if email notifications should be sent. Configurable for ` + "`" + `ORG` + "`" + `, ` + "`" + `GROUP` + "`" + `, and ` + "`" + `USER` + "`" + ` notifications types.`,
				},
				resource.Attribute{
					Name:        "flowdock_api_token",
					Description: `The Flowdock personal API token. Required for the ` + "`" + `FLOWDOCK` + "`" + ` notifications type. If the token later becomes invalid, Atlas sends an email to the project owner and eventually removes the token.`,
				},
				resource.Attribute{
					Name:        "flow_name",
					Description: `Flowdock flow name in lower-case letters. Required for the ` + "`" + `FLOWDOCK` + "`" + ` notifications type`,
				},
				resource.Attribute{
					Name:        "interval_min",
					Description: `Number of minutes to wait between successive notifications for unacknowledged alerts that are not resolved. The minimum value is 5.`,
				},
				resource.Attribute{
					Name:        "mobile_number",
					Description: `Mobile number to which alert notifications are sent. Required for the SMS notifications type.`,
				},
				resource.Attribute{
					Name:        "ops_genie_api_key",
					Description: `Opsgenie API Key. Required for the ` + "`" + `OPS_GENIE` + "`" + ` notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the token.`,
				},
				resource.Attribute{
					Name:        "ops_genie_region",
					Description: `Region that indicates which API URL to use. Accepted regions are: ` + "`" + `US` + "`" + ` ,` + "`" + `EU` + "`" + `. The default Opsgenie region is US.`,
				},
				resource.Attribute{
					Name:        "org_name",
					Description: `Flowdock organization name in lower-case letters. This is the name that appears after www.flowdock.com/app/ in the URL string. Required for the FLOWDOCK notifications type.`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `PagerDuty service key. Required for the PAGER_DUTY notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the key.`,
				},
				resource.Attribute{
					Name:        "sms_enabled",
					Description: `Flag indicating if text message notifications should be sent. Configurable for ` + "`" + `ORG` + "`" + `, ` + "`" + `GROUP` + "`" + `, and ` + "`" + `USER` + "`" + ` notifications types.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `Unique identifier of a team.`,
				},
				resource.Attribute{
					Name:        "type_name",
					Description: `Type of alert notification. Accepted values are: - ` + "`" + `DATADOG` + "`" + ` - ` + "`" + `EMAIL` + "`" + ` - ` + "`" + `FLOWDOCK` + "`" + ` - ` + "`" + `GROUP` + "`" + ` (Project) - ` + "`" + `OPS_GENIE` + "`" + ` - ` + "`" + `ORG` + "`" + ` - ` + "`" + `PAGER_DUTY` + "`" + ` - ` + "`" + `SLACK` + "`" + ` - ` + "`" + `SMS` + "`" + ` - ` + "`" + `TEAM` + "`" + ` - ` + "`" + `USER` + "`" + ` - ` + "`" + `VICTOR_OPS` + "`" + ` - ` + "`" + `WEBHOOK` + "`" + ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Name of the Atlas user to which to send notifications. Only a user in the project that owns the alert configuration is allowed here. Required for the ` + "`" + `USER` + "`" + ` notifications type.`,
				},
				resource.Attribute{
					Name:        "victor_ops_api_key",
					Description: `VictorOps API key. Required for the ` + "`" + `VICTOR_OPS` + "`" + ` notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the key.`,
				},
				resource.Attribute{
					Name:        "victor_ops_routing_key",
					Description: `VictorOps routing key. Optional for the ` + "`" + `VICTOR_OPS` + "`" + ` notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the key.`,
				},
				resource.Attribute{
					Name:        "Roles",
					Description: `The following roles grant privileges within a project. | Project roles | Organization roles | |:---------- |:----------- | | ` + "`" + `GROUP_CHARTS_ADMIN` + "`" + ` | ` + "`" + `ORG_OWNER` + "`" + ` | | ` + "`" + `GROUP_CLUSTER_MANAGER` + "`" + ` | ` + "`" + `ORG_MEMBER` + "`" + ` | | ` + "`" + `GROUP_DATA_ACCESS_ADMIN` + "`" + ` | ` + "`" + `ORG_GROUP_CREATOR` + "`" + ` | | ` + "`" + `GROUP_DATA_ACCESS_READ_ONLY` + "`" + ` | ` + "`" + `ORG_BILLING_ADMIN` + "`" + ` | | ` + "`" + `GROUP_DATA_ACCESS_READ_WRITE` + "`" + ` | ` + "`" + `ORG_READ_ONLY` + "`" + ` | | ` + "`" + `GROUP_OWNER` + "`" + ` | | | ` + "`" + `GROUP_READ_ONLY` + "`" + ` | | See detailed information for arguments and attributes: [MongoDB API Alert Configuration](https://docs.atlas.mongodb.com/reference/api/alert-configurations-get-config/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `Unique identifier of the project that owns this alert configuration.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Timestamp in ISO 8601 date and time format in UTC when this alert configuration was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Timestamp in ISO 8601 date and time format in UTC when this alert configuration was last updated.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If set to true, the alert configuration is enabled. If enabled is not exported it is set to false.`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `The type of event that will trigger an alert. ->`,
				},
				resource.Attribute{
					Name:        "field_name",
					Description: `Name of the field in the target object to match on. | Host alerts | Replica set alerts | Sharded cluster alerts | |:-------------- |:------------- |:------ | | ` + "`" + `TYPE_NAME` + "`" + ` | ` + "`" + `REPLICA_SET_NAME` + "`" + ` | ` + "`" + `CLUSTER_NAME` + "`" + ` | | ` + "`" + `HOSTNAME` + "`" + ` | ` + "`" + `SHARD_NAME` + "`" + ` | ` + "`" + `SHARD_NAME` + "`" + ` | | ` + "`" + `PORT` + "`" + ` | ` + "`" + `CLUSTER_NAME` + "`" + ` | | | ` + "`" + `HOSTNAME_AND_PORT` + "`" + ` | | | | ` + "`" + `REPLICA_SET_NAME` + "`" + ` | | | All other types of alerts do not support matchers.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `If omitted, the configuration is disabled.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `If omitted, the configuration is disabled.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `The operator to test the field’s value. Accepted values are: - ` + "`" + `EQUALS` + "`" + ` - ` + "`" + `NOT_EQUALS` + "`" + ` - ` + "`" + `CONTAINS` + "`" + ` - ` + "`" + `NOT_CONTAINS` + "`" + ` - ` + "`" + `STARTS_WITH` + "`" + ` - ` + "`" + `ENDS_WITH` + "`" + ` - ` + "`" + `REGEX` + "`" + ``,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value to test with the specified operator. If ` + "`" + `field_name` + "`" + ` is set to TYPE_NAME, you can match on the following values: - ` + "`" + `PRIMARY` + "`" + ` - ` + "`" + `SECONDARY` + "`" + ` - ` + "`" + `STANDALONE` + "`" + ` - ` + "`" + `CONFIG` + "`" + ` - ` + "`" + `MONGOS` + "`" + ` ### Metric Threshold The threshold that causes an alert to be triggered. Required if ` + "`" + `event_type_name` + "`" + ` : "OUTSIDE_METRIC_THRESHOLD".`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `Name of the metric to check. The full list of current options is available [here](https://docs.atlas.mongodb.com/reference/alert-host-metrics/#measurement-types)`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Operator to apply when checking the current metric value against the threshold value. Accepted values are: - ` + "`" + `GREATER_THAN` + "`" + ` - ` + "`" + `LESS_THAN` + "`" + ``,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Threshold value outside of which an alert will be triggered.`,
				},
				resource.Attribute{
					Name:        "units",
					Description: `The units for the threshold value. Depends on the type of metric. Accepted values are: - ` + "`" + `RAW` + "`" + ` - ` + "`" + `BITS` + "`" + ` - ` + "`" + `BYTES` + "`" + ` - ` + "`" + `KILOBITS` + "`" + ` - ` + "`" + `KILOBYTES` + "`" + ` - ` + "`" + `MEGABITS` + "`" + ` - ` + "`" + `MEGABYTES` + "`" + ` - ` + "`" + `GIGABITS` + "`" + ` - ` + "`" + `GIGABYTES` + "`" + ` - ` + "`" + `TERABYTES` + "`" + ` - ` + "`" + `PETABYTES` + "`" + ` - ` + "`" + `MILLISECONDS` + "`" + ` - ` + "`" + `SECONDS` + "`" + ` - ` + "`" + `MINUTES` + "`" + ` - ` + "`" + `HOURS` + "`" + ` - ` + "`" + `DAYS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `This must be set to AVERAGE. Atlas computes the current metric value as an average. ### Threshold`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `Operator to apply when checking the current metric value against the threshold value. Accepted values are: - ` + "`" + `GREATER_THAN` + "`" + ` - ` + "`" + `LESS_THAN` + "`" + ``,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `Threshold value outside of which an alert will be triggered.`,
				},
				resource.Attribute{
					Name:        "units",
					Description: `The units for the threshold value. Depends on the type of metric. Accepted values are: - ` + "`" + `RAW` + "`" + ` - ` + "`" + `BITS` + "`" + ` - ` + "`" + `BYTES` + "`" + ` - ` + "`" + `KILOBITS` + "`" + ` - ` + "`" + `KILOBYTES` + "`" + ` - ` + "`" + `MEGABITS` + "`" + ` - ` + "`" + `MEGABYTES` + "`" + ` - ` + "`" + `GIGABITS` + "`" + ` - ` + "`" + `GIGABYTES` + "`" + ` - ` + "`" + `TERABYTES` + "`" + ` - ` + "`" + `PETABYTES` + "`" + ` - ` + "`" + `MILLISECONDS` + "`" + ` - ` + "`" + `SECONDS` + "`" + ` - ` + "`" + `MINUTES` + "`" + ` - ` + "`" + `HOURS` + "`" + ` - ` + "`" + `DAYS` + "`" + ` ### Notifications Notifications to send when an alert condition is detected.`,
				},
				resource.Attribute{
					Name:        "api_token",
					Description: `Slack API token. Required for the SLACK notifications type. If the token later becomes invalid, Atlas sends an email to the project owner and eventually removes the token.`,
				},
				resource.Attribute{
					Name:        "channel_name",
					Description: `Slack channel name. Required for the SLACK notifications type.`,
				},
				resource.Attribute{
					Name:        "datadog_api_key",
					Description: `Datadog API Key. Found in the Datadog dashboard. Required for the DATADOG notifications type.`,
				},
				resource.Attribute{
					Name:        "datadog_region",
					Description: `Region that indicates which API URL to use. Accepted regions are: ` + "`" + `US` + "`" + `, ` + "`" + `EU` + "`" + `. The default Datadog region is US.`,
				},
				resource.Attribute{
					Name:        "delay_min",
					Description: `Number of minutes to wait after an alert condition is detected before sending out the first notification.`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `Email address to which alert notifications are sent. Required for the EMAIL notifications type.`,
				},
				resource.Attribute{
					Name:        "email_enabled",
					Description: `Flag indicating if email notifications should be sent. Configurable for ` + "`" + `ORG` + "`" + `, ` + "`" + `GROUP` + "`" + `, and ` + "`" + `USER` + "`" + ` notifications types.`,
				},
				resource.Attribute{
					Name:        "flowdock_api_token",
					Description: `The Flowdock personal API token. Required for the ` + "`" + `FLOWDOCK` + "`" + ` notifications type. If the token later becomes invalid, Atlas sends an email to the project owner and eventually removes the token.`,
				},
				resource.Attribute{
					Name:        "flow_name",
					Description: `Flowdock flow name in lower-case letters. Required for the ` + "`" + `FLOWDOCK` + "`" + ` notifications type`,
				},
				resource.Attribute{
					Name:        "interval_min",
					Description: `Number of minutes to wait between successive notifications for unacknowledged alerts that are not resolved. The minimum value is 5.`,
				},
				resource.Attribute{
					Name:        "mobile_number",
					Description: `Mobile number to which alert notifications are sent. Required for the SMS notifications type.`,
				},
				resource.Attribute{
					Name:        "ops_genie_api_key",
					Description: `Opsgenie API Key. Required for the ` + "`" + `OPS_GENIE` + "`" + ` notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the token.`,
				},
				resource.Attribute{
					Name:        "ops_genie_region",
					Description: `Region that indicates which API URL to use. Accepted regions are: ` + "`" + `US` + "`" + ` ,` + "`" + `EU` + "`" + `. The default Opsgenie region is US.`,
				},
				resource.Attribute{
					Name:        "org_name",
					Description: `Flowdock organization name in lower-case letters. This is the name that appears after www.flowdock.com/app/ in the URL string. Required for the FLOWDOCK notifications type.`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `PagerDuty service key. Required for the PAGER_DUTY notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the key.`,
				},
				resource.Attribute{
					Name:        "sms_enabled",
					Description: `Flag indicating if text message notifications should be sent. Configurable for ` + "`" + `ORG` + "`" + `, ` + "`" + `GROUP` + "`" + `, and ` + "`" + `USER` + "`" + ` notifications types.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `Unique identifier of a team.`,
				},
				resource.Attribute{
					Name:        "type_name",
					Description: `Type of alert notification. Accepted values are: - ` + "`" + `DATADOG` + "`" + ` - ` + "`" + `EMAIL` + "`" + ` - ` + "`" + `FLOWDOCK` + "`" + ` - ` + "`" + `GROUP` + "`" + ` (Project) - ` + "`" + `OPS_GENIE` + "`" + ` - ` + "`" + `ORG` + "`" + ` - ` + "`" + `PAGER_DUTY` + "`" + ` - ` + "`" + `SLACK` + "`" + ` - ` + "`" + `SMS` + "`" + ` - ` + "`" + `TEAM` + "`" + ` - ` + "`" + `USER` + "`" + ` - ` + "`" + `VICTOR_OPS` + "`" + ` - ` + "`" + `WEBHOOK` + "`" + ``,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Name of the Atlas user to which to send notifications. Only a user in the project that owns the alert configuration is allowed here. Required for the ` + "`" + `USER` + "`" + ` notifications type.`,
				},
				resource.Attribute{
					Name:        "victor_ops_api_key",
					Description: `VictorOps API key. Required for the ` + "`" + `VICTOR_OPS` + "`" + ` notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the key.`,
				},
				resource.Attribute{
					Name:        "victor_ops_routing_key",
					Description: `VictorOps routing key. Optional for the ` + "`" + `VICTOR_OPS` + "`" + ` notifications type. If the key later becomes invalid, Atlas sends an email to the project owner and eventually removes the key.`,
				},
				resource.Attribute{
					Name:        "Roles",
					Description: `The following roles grant privileges within a project. | Project roles | Organization roles | |:---------- |:----------- | | ` + "`" + `GROUP_CHARTS_ADMIN` + "`" + ` | ` + "`" + `ORG_OWNER` + "`" + ` | | ` + "`" + `GROUP_CLUSTER_MANAGER` + "`" + ` | ` + "`" + `ORG_MEMBER` + "`" + ` | | ` + "`" + `GROUP_DATA_ACCESS_ADMIN` + "`" + ` | ` + "`" + `ORG_GROUP_CREATOR` + "`" + ` | | ` + "`" + `GROUP_DATA_ACCESS_READ_ONLY` + "`" + ` | ` + "`" + `ORG_BILLING_ADMIN` + "`" + ` | | ` + "`" + `GROUP_DATA_ACCESS_READ_WRITE` + "`" + ` | ` + "`" + `ORG_READ_ONLY` + "`" + ` | | ` + "`" + `GROUP_OWNER` + "`" + ` | | | ` + "`" + `GROUP_READ_ONLY` + "`" + ` | | See detailed information for arguments and attributes: [MongoDB API Alert Configuration](https://docs.atlas.mongodb.com/reference/api/alert-configurations-get-config/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_auditing",
			Category:         "Data Sources",
			ShortDescription: `Describes a Auditing.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "configuration_type",
					Description: `Denotes the configuration method for the audit filter. Possible values are: NONE - auditing not configured for the project.m FILTER_BUILDER - auditing configured via Atlas UI filter builderm FILTER_JSON - auditing configured via Atlas custom filter or API.`,
				},
				resource.Attribute{
					Name:        "audit_filter",
					Description: `Indicates whether the auditing system captures successful authentication attempts for audit filters using the "atype" : "authCheck" auditing event. For more information, see auditAuthorizationSuccess`,
				},
				resource.Attribute{
					Name:        "audit_authorization_success",
					Description: `JSON-formatted audit filter used by the project`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Denotes whether or not the project associated with the {GROUP-ID} has database auditing enabled. See detailed information for arguments and attributes: [MongoDB API Auditing](https://docs.atlas.mongodb.com/reference/api/auditing/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "configuration_type",
					Description: `Denotes the configuration method for the audit filter. Possible values are: NONE - auditing not configured for the project.m FILTER_BUILDER - auditing configured via Atlas UI filter builderm FILTER_JSON - auditing configured via Atlas custom filter or API.`,
				},
				resource.Attribute{
					Name:        "audit_filter",
					Description: `Indicates whether the auditing system captures successful authentication attempts for audit filters using the "atype" : "authCheck" auditing event. For more information, see auditAuthorizationSuccess`,
				},
				resource.Attribute{
					Name:        "audit_authorization_success",
					Description: `JSON-formatted audit filter used by the project`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Denotes whether or not the project associated with the {GROUP-ID} has database auditing enabled. See detailed information for arguments and attributes: [MongoDB API Auditing](https://docs.atlas.mongodb.com/reference/api/auditing/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Provides an Cloud Backup Snapshot Datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Required) The unique identifier of the snapshot you want to retrieve.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Atlas cluster that contains the snapshot you want to retrieve.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the snapshot.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas took the snapshot.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas will delete the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `UDescription of the snapshot. Only present for on-demand snapshots.`,
				},
				resource.Attribute{
					Name:        "master_key_uuid",
					Description: `Unique ID of the AWS KMS Customer Master Key used to encrypt the snapshot. Only visible for clusters using Encryption at Rest via Customer KMS.`,
				},
				resource.Attribute{
					Name:        "mongod_version",
					Description: `Version of the MongoDB server.`,
				},
				resource.Attribute{
					Name:        "snapshot_type",
					Description: `Specified the type of snapshot. Valid values are onDemand and scheduled.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the snapshot. One of the following values: queued, inProgress, completed, failed.`,
				},
				resource.Attribute{
					Name:        "storage_size_bytes",
					Description: `Specifies the size of the snapshot in bytes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/get-one-backup/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the snapshot.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas took the snapshot.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas will delete the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `UDescription of the snapshot. Only present for on-demand snapshots.`,
				},
				resource.Attribute{
					Name:        "master_key_uuid",
					Description: `Unique ID of the AWS KMS Customer Master Key used to encrypt the snapshot. Only visible for clusters using Encryption at Rest via Customer KMS.`,
				},
				resource.Attribute{
					Name:        "mongod_version",
					Description: `Version of the MongoDB server.`,
				},
				resource.Attribute{
					Name:        "snapshot_type",
					Description: `Specified the type of snapshot. Valid values are onDemand and scheduled.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the snapshot. One of the following values: queued, inProgress, completed, failed.`,
				},
				resource.Attribute{
					Name:        "storage_size_bytes",
					Description: `Specifies the size of the snapshot in bytes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/get-one-backup/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot_backup_policy",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Backup Snapshot Policy Datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Atlas cluster that contains the snapshots backup policy you want to retrieve. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Unique identifier of the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "next_snapshot",
					Description: `UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.`,
				},
				resource.Attribute{
					Name:        "reference_hour_of_day",
					Description: `UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.`,
				},
				resource.Attribute{
					Name:        "reference_minute_of_hour",
					Description: `UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.`,
				},
				resource.Attribute{
					Name:        "restore_window_days",
					Description: `Specifies a restore window in days for cloud backup to maintain. ### Policies`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policy definitions for the cluster.`,
				},
				resource.Attribute{
					Name:        "policies.#.id",
					Description: `Unique identifier of the backup policy. #### Policy Item`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item",
					Description: `A list of specifications for a policy.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.id",
					Description: `Unique identifier for this policy item.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.frequency_interval",
					Description: `The frequency interval for a set of snapshots.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.frequency_type",
					Description: `A type of frequency (hourly, daily, weekly, monthly).`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.retention_unit",
					Description: `The unit of time in which snapshot retention is measured (days, weeks, months).`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.retention_value",
					Description: `The number of days, weeks, or months the snapshot is retained. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/schedule/get-all-schedules/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Unique identifier of the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "next_snapshot",
					Description: `UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.`,
				},
				resource.Attribute{
					Name:        "reference_hour_of_day",
					Description: `UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.`,
				},
				resource.Attribute{
					Name:        "reference_minute_of_hour",
					Description: `UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.`,
				},
				resource.Attribute{
					Name:        "restore_window_days",
					Description: `Specifies a restore window in days for cloud backup to maintain. ### Policies`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policy definitions for the cluster.`,
				},
				resource.Attribute{
					Name:        "policies.#.id",
					Description: `Unique identifier of the backup policy. #### Policy Item`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item",
					Description: `A list of specifications for a policy.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.id",
					Description: `Unique identifier for this policy item.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.frequency_interval",
					Description: `The frequency interval for a set of snapshots.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.frequency_type",
					Description: `A type of frequency (hourly, daily, weekly, monthly).`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.retention_unit",
					Description: `The unit of time in which snapshot retention is measured (days, weeks, months).`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.retention_value",
					Description: `The number of days, weeks, or months the snapshot is retained. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/schedule/get-all-schedules/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot_restore_job",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Backup Snapshot Restore Job Datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Atlas cluster for which you want to retrieve the restore job.`,
				},
				resource.Attribute{
					Name:        "job_id",
					Description: `(Required) The unique identifier of the restore job to retrieve. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cancelled",
					Description: `Indicates whether the restore job was canceled.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas created the restore job.`,
				},
				resource.Attribute{
					Name:        "delivery_type",
					Description: `Type of restore job to create. Possible values are: automated and download.`,
				},
				resource.Attribute{
					Name:        "delivery_url",
					Description: `One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `Indicates whether the restore job expired.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `UTC ISO 8601 formatted point in time when the restore job expires.`,
				},
				resource.Attribute{
					Name:        "finished_at",
					Description: `UTC ISO 8601 formatted point in time when the restore job completed.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the restore job.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Unique identifier of the source snapshot ID of the restore job.`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `Name of the target Atlas project of the restore job. Only visible if deliveryType is automated.`,
				},
				resource.Attribute{
					Name:        "target_cluster_name",
					Description: `Name of the target Atlas cluster to which the restore job restores the snapshot. Only visible if deliveryType is automated.`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.`,
				},
				resource.Attribute{
					Name:        "oplogTs",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
				},
				resource.Attribute{
					Name:        "oplogInc",
					Description: `Oplog operation number from which to you want to restore this snapshot.`,
				},
				resource.Attribute{
					Name:        "pointInTimeUTCSeconds",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/get-one-restore-job/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cancelled",
					Description: `Indicates whether the restore job was canceled.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas created the restore job.`,
				},
				resource.Attribute{
					Name:        "delivery_type",
					Description: `Type of restore job to create. Possible values are: automated and download.`,
				},
				resource.Attribute{
					Name:        "delivery_url",
					Description: `One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `Indicates whether the restore job expired.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `UTC ISO 8601 formatted point in time when the restore job expires.`,
				},
				resource.Attribute{
					Name:        "finished_at",
					Description: `UTC ISO 8601 formatted point in time when the restore job completed.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the restore job.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Unique identifier of the source snapshot ID of the restore job.`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `Name of the target Atlas project of the restore job. Only visible if deliveryType is automated.`,
				},
				resource.Attribute{
					Name:        "target_cluster_name",
					Description: `Name of the target Atlas cluster to which the restore job restores the snapshot. Only visible if deliveryType is automated.`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.`,
				},
				resource.Attribute{
					Name:        "oplogTs",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
				},
				resource.Attribute{
					Name:        "oplogInc",
					Description: `Oplog operation number from which to you want to restore this snapshot.`,
				},
				resource.Attribute{
					Name:        "pointInTimeUTCSeconds",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/get-one-restore-job/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot_restore_jobs",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Backup Snapshot Restore Jobs Datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Atlas cluster for which you want to retrieve restore jobs.`,
				},
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `Includes cloudProviderSnapshotRestoreJob object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### CloudProviderSnapshotRestoreJob`,
				},
				resource.Attribute{
					Name:        "cancelled",
					Description: `Indicates whether the restore job was canceled.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas created the restore job.`,
				},
				resource.Attribute{
					Name:        "delivery_type",
					Description: `Type of restore job to create. Possible values are: automated and download.`,
				},
				resource.Attribute{
					Name:        "delivery_url",
					Description: `One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `Indicates whether the restore job expired.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `UTC ISO 8601 formatted point in time when the restore job expires.`,
				},
				resource.Attribute{
					Name:        "finished_at",
					Description: `UTC ISO 8601 formatted point in time when the restore job completed.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the restore job.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Unique identifier of the source snapshot ID of the restore job.`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `Name of the target Atlas project of the restore job. Only visible if deliveryType is automated.`,
				},
				resource.Attribute{
					Name:        "target_cluster_name",
					Description: `Name of the target Atlas cluster to which the restore job restores the snapshot. Only visible if deliveryType is automated.`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.`,
				},
				resource.Attribute{
					Name:        "oplogTs",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
				},
				resource.Attribute{
					Name:        "oplogInc",
					Description: `Oplog operation number from which to you want to restore this snapshot.`,
				},
				resource.Attribute{
					Name:        "pointInTimeUTCSeconds",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/get-all-restore-jobs/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "results",
					Description: `Includes cloudProviderSnapshotRestoreJob object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### CloudProviderSnapshotRestoreJob`,
				},
				resource.Attribute{
					Name:        "cancelled",
					Description: `Indicates whether the restore job was canceled.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas created the restore job.`,
				},
				resource.Attribute{
					Name:        "delivery_type",
					Description: `Type of restore job to create. Possible values are: automated and download.`,
				},
				resource.Attribute{
					Name:        "delivery_url",
					Description: `One or more URLs for the compressed snapshot files for manual download. Only visible if deliveryType is download.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `Indicates whether the restore job expired.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `UTC ISO 8601 formatted point in time when the restore job expires.`,
				},
				resource.Attribute{
					Name:        "finished_at",
					Description: `UTC ISO 8601 formatted point in time when the restore job completed.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the restore job.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Unique identifier of the source snapshot ID of the restore job.`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `Name of the target Atlas project of the restore job. Only visible if deliveryType is automated.`,
				},
				resource.Attribute{
					Name:        "target_cluster_name",
					Description: `Name of the target Atlas cluster to which the restore job restores the snapshot. Only visible if deliveryType is automated.`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken.`,
				},
				resource.Attribute{
					Name:        "oplogTs",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
				},
				resource.Attribute{
					Name:        "oplogInc",
					Description: `Oplog operation number from which to you want to restore this snapshot.`,
				},
				resource.Attribute{
					Name:        "pointInTimeUTCSeconds",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/get-all-restore-jobs/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshots",
			Category:         "Data Sources",
			ShortDescription: `Provides an Cloud Backup Snapshot Datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Atlas cluster that contains the snapshot you want to retrieve.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `Includes cloudProviderSnapshot object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### CloudProviderSnapshot`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the snapshot.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas took the snapshot.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas will delete the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `UDescription of the snapshot. Only present for on-demand snapshots.`,
				},
				resource.Attribute{
					Name:        "master_key_uuid",
					Description: `Unique ID of the AWS KMS Customer Master Key used to encrypt the snapshot. Only visible for clusters using Encryption at Rest via Customer KMS.`,
				},
				resource.Attribute{
					Name:        "mongod_version",
					Description: `Version of the MongoDB server.`,
				},
				resource.Attribute{
					Name:        "snapshot_type",
					Description: `Specified the type of snapshot. Valid values are onDemand and scheduled.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the snapshot. One of the following values: queued, inProgress, completed, failed.`,
				},
				resource.Attribute{
					Name:        "storage_size_bytes",
					Description: `Specifies the size of the snapshot in bytes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/get-all-backups/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "results",
					Description: `Includes cloudProviderSnapshot object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### CloudProviderSnapshot`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the snapshot.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas took the snapshot.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas will delete the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `UDescription of the snapshot. Only present for on-demand snapshots.`,
				},
				resource.Attribute{
					Name:        "master_key_uuid",
					Description: `Unique ID of the AWS KMS Customer Master Key used to encrypt the snapshot. Only visible for clusters using Encryption at Rest via Customer KMS.`,
				},
				resource.Attribute{
					Name:        "mongod_version",
					Description: `Version of the MongoDB server.`,
				},
				resource.Attribute{
					Name:        "snapshot_type",
					Description: `Specified the type of snapshot. Valid values are onDemand and scheduled.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the snapshot. One of the following values: queued, inProgress, completed, failed.`,
				},
				resource.Attribute{
					Name:        "storage_size_bytes",
					Description: `Specifies the size of the snapshot in bytes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/get-all-backups/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cluster",
			Category:         "Data Sources",
			ShortDescription: `Describe a Cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the cluster as it appears in Atlas. Once the cluster is created, its name cannot be changed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
				},
				resource.Attribute{
					Name:        "mongo_uri",
					Description: `Base connection string for the cluster. Atlas only displays this field after the cluster is operational, not while it builds the cluster.`,
				},
				resource.Attribute{
					Name:        "mongo_uri_updated",
					Description: `Lists when the connection string was last updated. The connection string changes, for example, if you change a replica set to a sharded cluster.`,
				},
				resource.Attribute{
					Name:        "mongo_uri_with_options",
					Description: `Describes connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster. To review the connection string format, see the connection string format documentation. To add MongoDB users to a Atlas project, see Configure MongoDB Users. Atlas only displays this field after the cluster is operational, not while it builds the cluster.`,
				},
				resource.Attribute{
					Name:        "paused",
					Description: `Flag that indicates whether the cluster is paused or not.`,
				},
				resource.Attribute{
					Name:        "pit_enabled",
					Description: `Flag that indicates if the cluster uses Continuous Cloud Backup.`,
				},
				resource.Attribute{
					Name:        "srv_address",
					Description: `Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS/SSL. See the mongoURI for additional options.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Indicates the current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING`,
				},
				resource.Attribute{
					Name:        "auto_scaling_disk_gb_enabled",
					Description: `Indicates whether disk auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_enabled",
					Description: `(Optional) Specifies whether cluster tier auto-scaling is enabled. The default is false.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_scale_down_enabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable the cluster tier to scale down.`,
				},
				resource.Attribute{
					Name:        "backup_enabled",
					Description: `Legacy Option, Indicates whether Atlas continuous backups are enabled for the cluster.`,
				},
				resource.Attribute{
					Name:        "bi_connector",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "bi_connector_config",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Indicates the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.`,
				},
				resource.Attribute{
					Name:        "connection_strings",
					Description: `Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster. - ` + "`" + `connection_strings.standard` + "`" + ` - Public mongodb:// connection string for this cluster. - ` + "`" + `connection_strings.standard_srv` + "`" + ` - Public mongodb+srv:// connection string for this cluster. The mongodb+srv protocol tells the driver to look up the seed list of hosts in DNS. Atlas synchronizes this list with the nodes in a cluster. If the connection string uses this URI format, you don’t need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn’t, use connectionStrings.standard. - ` + "`" + `connection_strings.aws_private_link` + "`" + ` - [Private-endpoint-aware](https://docs.atlas.mongodb.com/security-private-endpoint/#private-endpoint-connection-strings) mongodb://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a AWS PrivateLink connection to this cluster.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Indicates the size in gigabytes of the server’s root volume (AWS/GCP Only).`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `Indicates whether Encryption at Rest is enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster as it appears in Atlas.`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `Indicates the version of the cluster to deploy.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Indicates whether the cluster is a replica set or a sharded cluster.`,
				},
				resource.Attribute{
					Name:        "provider_backup_enabled",
					Description: `Flag indicating if the cluster uses Cloud Backup Snapshots for backups.`,
				},
				resource.Attribute{
					Name:        "provider_instance_size_name",
					Description: `Atlas provides different instance sizes, each with a default storage capacity and RAM size.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Indicates the cloud service provider on which the servers are provisioned.`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `Indicates Cloud service provider on which the server for a multi-tenant cluster is provisioned.`,
				},
				resource.Attribute{
					Name:        "provider_disk_iops",
					Description: `Indicates the maximum input/output operations per second (IOPS) the system can perform. The possible values depend on the selected providerSettings.instanceSizeName and diskSizeGB.`,
				},
				resource.Attribute{
					Name:        "provider_disk_type_name",
					Description: `Describes Azure disk type of the server’s root volume (Azure Only).`,
				},
				resource.Attribute{
					Name:        "provider_region_name",
					Description: `Indicates Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. Requires the Atlas Region name, see the reference list for [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).`,
				},
				resource.Attribute{
					Name:        "provider_volume_type",
					Description: `Indicates the type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `(Deprecated) Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_min_instance_size",
					Description: `(Optional) Minimum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_max_instance_size",
					Description: `(Optional) Maximum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `The Network Peering Container ID. ### BI Connector Indicates BI Connector for Atlas configuration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates whether or not BI Connector for Atlas is enabled on the cluster.`,
				},
				resource.Attribute{
					Name:        "read_preference",
					Description: `Indicates the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences). ### Replication Spec Configuration for cluster regions.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifer of the replication document for a zone in a Global Cluster.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Number of shards to deploy in the specified zone.`,
				},
				resource.Attribute{
					Name:        "regions_config",
					Description: `Describes the physical location of the region. Each regionsConfig document describes the region’s priority in elections and the number and type of MongoDB nodes Atlas deploys to the region. You must order each regionsConfigs document by regionsConfig.priority, descending. See [Region Config](#region-config) below for more details.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Indicates the n ame for the zone in a Global Cluster. ### Region Config Physical location of the region.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Name for the region specified.`,
				},
				resource.Attribute{
					Name:        "electable_nodes",
					Description: `Number of electable nodes for Atlas to deploy to the region.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Election priority of the region. For regions with only read-only nodes, set this value to 0.`,
				},
				resource.Attribute{
					Name:        "read_only_nodes",
					Description: `Number of read-only nodes for Atlas to deploy to the region. Read-only nodes can never become the primary, but can facilitate local-reads. Specify 0 if you do not want any read-only nodes in the region.`,
				},
				resource.Attribute{
					Name:        "analytics_nodes",
					Description: `Indicates the number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. ### Labels Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. Note: the key ` + "`" + `Infrastructure Tool` + "`" + `, is used for internal purposes to track aggregate usage.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that was set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that represents the key. ### Plugin Contains a key-value pair that tags that the cluster was created by a Terraform Provider and notes the version.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the current plugin`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of the plugin. ### Cloud Provider Snapshot Backup Policy`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy",
					Description: `current snapshot schedule and retention settings for the cluster.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.cluster_id",
					Description: `Unique identifier of the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.cluster_name",
					Description: `Name of the Atlas cluster that contains the snapshot backup policy.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.next_snapshot",
					Description: `UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.reference_hour_of_day",
					Description: `UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.reference_minute_of_hour",
					Description: `UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.restore_window_days",
					Description: `Specifies a restore window in days for the cloud provider backup to maintain.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.update_snapshots",
					Description: `Specifies it's true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously. ### Policies`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies",
					Description: `A list of policy definitions for the cluster.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.id",
					Description: `Unique identifier of the backup policy. #### Policy Item`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item",
					Description: `A list of specifications for a policy.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.id",
					Description: `Unique identifier for this policy item.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.frequency_interval",
					Description: `The frequency interval for a set of snapshots.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.frequency_type",
					Description: `A type of frequency (hourly, daily, weekly, monthly).`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.retention_unit",
					Description: `The unit of time in which snapshot retention is measured (days, weeks, months).`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.retention_value",
					Description: `The number of days, weeks, or months the snapshot is retained. See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
				},
				resource.Attribute{
					Name:        "mongo_uri",
					Description: `Base connection string for the cluster. Atlas only displays this field after the cluster is operational, not while it builds the cluster.`,
				},
				resource.Attribute{
					Name:        "mongo_uri_updated",
					Description: `Lists when the connection string was last updated. The connection string changes, for example, if you change a replica set to a sharded cluster.`,
				},
				resource.Attribute{
					Name:        "mongo_uri_with_options",
					Description: `Describes connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster. To review the connection string format, see the connection string format documentation. To add MongoDB users to a Atlas project, see Configure MongoDB Users. Atlas only displays this field after the cluster is operational, not while it builds the cluster.`,
				},
				resource.Attribute{
					Name:        "paused",
					Description: `Flag that indicates whether the cluster is paused or not.`,
				},
				resource.Attribute{
					Name:        "pit_enabled",
					Description: `Flag that indicates if the cluster uses Continuous Cloud Backup.`,
				},
				resource.Attribute{
					Name:        "srv_address",
					Description: `Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS/SSL. See the mongoURI for additional options.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Indicates the current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING`,
				},
				resource.Attribute{
					Name:        "auto_scaling_disk_gb_enabled",
					Description: `Indicates whether disk auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_enabled",
					Description: `(Optional) Specifies whether cluster tier auto-scaling is enabled. The default is false.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_scale_down_enabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable the cluster tier to scale down.`,
				},
				resource.Attribute{
					Name:        "backup_enabled",
					Description: `Legacy Option, Indicates whether Atlas continuous backups are enabled for the cluster.`,
				},
				resource.Attribute{
					Name:        "bi_connector",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "bi_connector_config",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Indicates the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.`,
				},
				resource.Attribute{
					Name:        "connection_strings",
					Description: `Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster. - ` + "`" + `connection_strings.standard` + "`" + ` - Public mongodb:// connection string for this cluster. - ` + "`" + `connection_strings.standard_srv` + "`" + ` - Public mongodb+srv:// connection string for this cluster. The mongodb+srv protocol tells the driver to look up the seed list of hosts in DNS. Atlas synchronizes this list with the nodes in a cluster. If the connection string uses this URI format, you don’t need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn’t, use connectionStrings.standard. - ` + "`" + `connection_strings.aws_private_link` + "`" + ` - [Private-endpoint-aware](https://docs.atlas.mongodb.com/security-private-endpoint/#private-endpoint-connection-strings) mongodb://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a AWS PrivateLink connection to this cluster.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Indicates the size in gigabytes of the server’s root volume (AWS/GCP Only).`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `Indicates whether Encryption at Rest is enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster as it appears in Atlas.`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `Indicates the version of the cluster to deploy.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Indicates whether the cluster is a replica set or a sharded cluster.`,
				},
				resource.Attribute{
					Name:        "provider_backup_enabled",
					Description: `Flag indicating if the cluster uses Cloud Backup Snapshots for backups.`,
				},
				resource.Attribute{
					Name:        "provider_instance_size_name",
					Description: `Atlas provides different instance sizes, each with a default storage capacity and RAM size.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Indicates the cloud service provider on which the servers are provisioned.`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `Indicates Cloud service provider on which the server for a multi-tenant cluster is provisioned.`,
				},
				resource.Attribute{
					Name:        "provider_disk_iops",
					Description: `Indicates the maximum input/output operations per second (IOPS) the system can perform. The possible values depend on the selected providerSettings.instanceSizeName and diskSizeGB.`,
				},
				resource.Attribute{
					Name:        "provider_disk_type_name",
					Description: `Describes Azure disk type of the server’s root volume (Azure Only).`,
				},
				resource.Attribute{
					Name:        "provider_region_name",
					Description: `Indicates Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. Requires the Atlas Region name, see the reference list for [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).`,
				},
				resource.Attribute{
					Name:        "provider_volume_type",
					Description: `Indicates the type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `(Deprecated) Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_min_instance_size",
					Description: `(Optional) Minimum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_max_instance_size",
					Description: `(Optional) Maximum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `The Network Peering Container ID. ### BI Connector Indicates BI Connector for Atlas configuration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates whether or not BI Connector for Atlas is enabled on the cluster.`,
				},
				resource.Attribute{
					Name:        "read_preference",
					Description: `Indicates the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences). ### Replication Spec Configuration for cluster regions.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifer of the replication document for a zone in a Global Cluster.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Number of shards to deploy in the specified zone.`,
				},
				resource.Attribute{
					Name:        "regions_config",
					Description: `Describes the physical location of the region. Each regionsConfig document describes the region’s priority in elections and the number and type of MongoDB nodes Atlas deploys to the region. You must order each regionsConfigs document by regionsConfig.priority, descending. See [Region Config](#region-config) below for more details.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Indicates the n ame for the zone in a Global Cluster. ### Region Config Physical location of the region.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Name for the region specified.`,
				},
				resource.Attribute{
					Name:        "electable_nodes",
					Description: `Number of electable nodes for Atlas to deploy to the region.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Election priority of the region. For regions with only read-only nodes, set this value to 0.`,
				},
				resource.Attribute{
					Name:        "read_only_nodes",
					Description: `Number of read-only nodes for Atlas to deploy to the region. Read-only nodes can never become the primary, but can facilitate local-reads. Specify 0 if you do not want any read-only nodes in the region.`,
				},
				resource.Attribute{
					Name:        "analytics_nodes",
					Description: `Indicates the number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. ### Labels Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. Note: the key ` + "`" + `Infrastructure Tool` + "`" + `, is used for internal purposes to track aggregate usage.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that was set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that represents the key. ### Plugin Contains a key-value pair that tags that the cluster was created by a Terraform Provider and notes the version.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the current plugin`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of the plugin. ### Cloud Provider Snapshot Backup Policy`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy",
					Description: `current snapshot schedule and retention settings for the cluster.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.cluster_id",
					Description: `Unique identifier of the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.cluster_name",
					Description: `Name of the Atlas cluster that contains the snapshot backup policy.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.next_snapshot",
					Description: `UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.reference_hour_of_day",
					Description: `UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.reference_minute_of_hour",
					Description: `UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.restore_window_days",
					Description: `Specifies a restore window in days for the cloud provider backup to maintain.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.update_snapshots",
					Description: `Specifies it's true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously. ### Policies`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies",
					Description: `A list of policy definitions for the cluster.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.id",
					Description: `Unique identifier of the backup policy. #### Policy Item`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item",
					Description: `A list of specifications for a policy.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.id",
					Description: `Unique identifier for this policy item.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.frequency_interval",
					Description: `The frequency interval for a set of snapshots.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.frequency_type",
					Description: `A type of frequency (hourly, daily, weekly, monthly).`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.retention_unit",
					Description: `The unit of time in which snapshot retention is measured (days, weeks, months).`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.retention_value",
					Description: `The number of days, weeks, or months the snapshot is retained. See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_clusters",
			Category:         "Data Sources",
			ShortDescription: `Describe all Clusters in Project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to get the clusters. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Cluster. See [Cluster](#cluster) below for more details. ### Cluster`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster as it appears in Atlas.`,
				},
				resource.Attribute{
					Name:        "mongo_uri",
					Description: `Base connection string for the cluster. Atlas only displays this field after the cluster is operational, not while it builds the cluster.`,
				},
				resource.Attribute{
					Name:        "mongo_uri_updated",
					Description: `Lists when the connection string was last updated. The connection string changes, for example, if you change a replica set to a sharded cluster.`,
				},
				resource.Attribute{
					Name:        "mongo_uri_with_options",
					Description: `Describes connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster. To review the connection string format, see the connection string format documentation. To add MongoDB users to a Atlas project, see Configure MongoDB Users. Atlas only displays this field after the cluster is operational, not while it builds the cluster.`,
				},
				resource.Attribute{
					Name:        "paused",
					Description: `Flag that indicates whether the cluster is paused or not.`,
				},
				resource.Attribute{
					Name:        "pit_enabled",
					Description: `Flag that indicates if the cluster uses Continuous Cloud Backup.`,
				},
				resource.Attribute{
					Name:        "srv_address",
					Description: `Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS/SSL. See the mongoURI for additional options.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Indicates the current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING`,
				},
				resource.Attribute{
					Name:        "auto_scaling_disk_gb_enabled",
					Description: `Indicates whether disk auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_enabled",
					Description: `(Optional) Specifies whether cluster tier auto-scaling is enabled. The default is false.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_scale_down_enabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable the cluster tier to scale down.`,
				},
				resource.Attribute{
					Name:        "backup_enabled",
					Description: `Legacy Option, Indicates whether Atlas continuous backups are enabled for the cluster.`,
				},
				resource.Attribute{
					Name:        "bi_connector",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "bi_connector_config",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Indicates the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.`,
				},
				resource.Attribute{
					Name:        "connection_strings",
					Description: `Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster. - ` + "`" + `connection_strings.standard` + "`" + ` - Public mongodb:// connection string for this cluster. - ` + "`" + `connection_strings.standard_srv` + "`" + ` - Public mongodb+srv:// connection string for this cluster. The mongodb+srv protocol tells the driver to look up the seed list of hosts in DNS. Atlas synchronizes this list with the nodes in a cluster. If the connection string uses this URI format, you don’t need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn’t, use connectionStrings.standard. - ` + "`" + `connection_strings.aws_private_link` + "`" + ` - [Private-endpoint-aware](https://docs.atlas.mongodb.com/security-private-endpoint/#private-endpoint-connection-strings) mongodb://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a AWS PrivateLink connection to this cluster.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Indicates the size in gigabytes of the server’s root volume (AWS/GCP Only).`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `Indicates whether Encryption at Rest is enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `Indicates the version of the cluster to deploy.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Indicates whether the cluster is a replica set or a sharded cluster.`,
				},
				resource.Attribute{
					Name:        "provider_backup_enabled",
					Description: `Flag indicating if the cluster uses Cloud Backup Snapshots for backups.`,
				},
				resource.Attribute{
					Name:        "provider_instance_size_name",
					Description: `Atlas provides different instance sizes, each with a default storage capacity and RAM size.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Indicates the cloud service provider on which the servers are provisioned.`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `Indicates Cloud service provider on which the server for a multi-tenant cluster is provisioned.`,
				},
				resource.Attribute{
					Name:        "provider_disk_iops",
					Description: `Indicates the maximum input/output operations per second (IOPS) the system can perform. The possible values depend on the selected providerSettings.instanceSizeName and diskSizeGB.`,
				},
				resource.Attribute{
					Name:        "provider_disk_type_name",
					Description: `Describes Azure disk type of the server’s root volume (Azure Only).`,
				},
				resource.Attribute{
					Name:        "provider_region_name",
					Description: `Indicates Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. Requires the Atlas Region name, see the reference list for [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).`,
				},
				resource.Attribute{
					Name:        "provider_volume_type",
					Description: `Indicates the type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_min_instance_size",
					Description: `(Optional) Minimum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_max_instance_size",
					Description: `(Optional) Maximum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `(Deprecated) Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `The Network Peering Container ID. ### BI Connector Indicates BI Connector for Atlas configuration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates whether or not BI Connector for Atlas is enabled on the cluster.`,
				},
				resource.Attribute{
					Name:        "read_preference",
					Description: `Indicates the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences). ### Replication Spec Configuration for cluster regions.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifer of the replication document for a zone in a Global Cluster.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Number of shards to deploy in the specified zone.`,
				},
				resource.Attribute{
					Name:        "regions_config",
					Description: `Describes the physical location of the region. Each regionsConfig document describes the region’s priority in elections and the number and type of MongoDB nodes Atlas deploys to the region. You must order each regionsConfigs document by regionsConfig.priority, descending. See [Region Config](#region-config) below for more details.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Indicates the n ame for the zone in a Global Cluster. ### Region Config Physical location of the region.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Name for the region specified.`,
				},
				resource.Attribute{
					Name:        "electable_nodes",
					Description: `Number of electable nodes for Atlas to deploy to the region.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Election priority of the region. For regions with only read-only nodes, set this value to 0.`,
				},
				resource.Attribute{
					Name:        "read_only_nodes",
					Description: `Number of read-only nodes for Atlas to deploy to the region. Read-only nodes can never become the primary, but can facilitate local-reads. Specify 0 if you do not want any read-only nodes in the region.`,
				},
				resource.Attribute{
					Name:        "analytics_nodes",
					Description: `Indicates the number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. ### Labels Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. Note: the key ` + "`" + `Infrastructure Tool` + "`" + `, is used for internal purposes to track aggregate usage.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that was set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that represents the key. ### Plugin Contains a key-value pair that tags that the cluster was created by a Terraform Provider and notes the version.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the current plugin`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of the plugin. ### Cloud Provider Snapshot Backup Policy`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy",
					Description: `current snapshot schedule and retention settings for the cluster.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.cluster_id",
					Description: `Unique identifier of the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.cluster_name",
					Description: `Name of the Atlas cluster that contains the snapshot backup policy.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.next_snapshot",
					Description: `UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.reference_hour_of_day",
					Description: `UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.reference_minute_of_hour",
					Description: `UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.restore_window_days",
					Description: `Specifies a restore window in days for the cloud provider backup to maintain.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.update_snapshots",
					Description: `Specifies it's true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously. ### Policies`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies",
					Description: `A list of policy definitions for the cluster.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.id",
					Description: `Unique identifier of the backup policy. #### Policy Item`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item",
					Description: `A list of specifications for a policy.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.id",
					Description: `Unique identifier for this policy item.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.frequency_interval",
					Description: `The frequency interval for a set of snapshots.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.frequency_type",
					Description: `A type of frequency (hourly, daily, weekly, monthly).`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.retention_unit",
					Description: `The unit of time in which snapshot retention is measured (days, weeks, months).`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.retention_value",
					Description: `The number of days, weeks, or months the snapshot is retained. See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Cluster. See [Cluster](#cluster) below for more details. ### Cluster`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster as it appears in Atlas.`,
				},
				resource.Attribute{
					Name:        "mongo_uri",
					Description: `Base connection string for the cluster. Atlas only displays this field after the cluster is operational, not while it builds the cluster.`,
				},
				resource.Attribute{
					Name:        "mongo_uri_updated",
					Description: `Lists when the connection string was last updated. The connection string changes, for example, if you change a replica set to a sharded cluster.`,
				},
				resource.Attribute{
					Name:        "mongo_uri_with_options",
					Description: `Describes connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster. To review the connection string format, see the connection string format documentation. To add MongoDB users to a Atlas project, see Configure MongoDB Users. Atlas only displays this field after the cluster is operational, not while it builds the cluster.`,
				},
				resource.Attribute{
					Name:        "paused",
					Description: `Flag that indicates whether the cluster is paused or not.`,
				},
				resource.Attribute{
					Name:        "pit_enabled",
					Description: `Flag that indicates if the cluster uses Continuous Cloud Backup.`,
				},
				resource.Attribute{
					Name:        "srv_address",
					Description: `Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS/SSL. See the mongoURI for additional options.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Indicates the current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING`,
				},
				resource.Attribute{
					Name:        "auto_scaling_disk_gb_enabled",
					Description: `Indicates whether disk auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_enabled",
					Description: `(Optional) Specifies whether cluster tier auto-scaling is enabled. The default is false.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_scale_down_enabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable the cluster tier to scale down.`,
				},
				resource.Attribute{
					Name:        "backup_enabled",
					Description: `Legacy Option, Indicates whether Atlas continuous backups are enabled for the cluster.`,
				},
				resource.Attribute{
					Name:        "bi_connector",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "bi_connector_config",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Indicates the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.`,
				},
				resource.Attribute{
					Name:        "connection_strings",
					Description: `Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster. - ` + "`" + `connection_strings.standard` + "`" + ` - Public mongodb:// connection string for this cluster. - ` + "`" + `connection_strings.standard_srv` + "`" + ` - Public mongodb+srv:// connection string for this cluster. The mongodb+srv protocol tells the driver to look up the seed list of hosts in DNS. Atlas synchronizes this list with the nodes in a cluster. If the connection string uses this URI format, you don’t need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn’t, use connectionStrings.standard. - ` + "`" + `connection_strings.aws_private_link` + "`" + ` - [Private-endpoint-aware](https://docs.atlas.mongodb.com/security-private-endpoint/#private-endpoint-connection-strings) mongodb://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a AWS PrivateLink connection to this cluster.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Indicates the size in gigabytes of the server’s root volume (AWS/GCP Only).`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `Indicates whether Encryption at Rest is enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `Indicates the version of the cluster to deploy.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Indicates whether the cluster is a replica set or a sharded cluster.`,
				},
				resource.Attribute{
					Name:        "provider_backup_enabled",
					Description: `Flag indicating if the cluster uses Cloud Backup Snapshots for backups.`,
				},
				resource.Attribute{
					Name:        "provider_instance_size_name",
					Description: `Atlas provides different instance sizes, each with a default storage capacity and RAM size.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Indicates the cloud service provider on which the servers are provisioned.`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `Indicates Cloud service provider on which the server for a multi-tenant cluster is provisioned.`,
				},
				resource.Attribute{
					Name:        "provider_disk_iops",
					Description: `Indicates the maximum input/output operations per second (IOPS) the system can perform. The possible values depend on the selected providerSettings.instanceSizeName and diskSizeGB.`,
				},
				resource.Attribute{
					Name:        "provider_disk_type_name",
					Description: `Describes Azure disk type of the server’s root volume (Azure Only).`,
				},
				resource.Attribute{
					Name:        "provider_region_name",
					Description: `Indicates Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. Requires the Atlas Region name, see the reference list for [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).`,
				},
				resource.Attribute{
					Name:        "provider_volume_type",
					Description: `Indicates the type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_min_instance_size",
					Description: `(Optional) Minimum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_max_instance_size",
					Description: `(Optional) Maximum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `(Deprecated) Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `The Network Peering Container ID. ### BI Connector Indicates BI Connector for Atlas configuration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates whether or not BI Connector for Atlas is enabled on the cluster.`,
				},
				resource.Attribute{
					Name:        "read_preference",
					Description: `Indicates the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences). ### Replication Spec Configuration for cluster regions.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifer of the replication document for a zone in a Global Cluster.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Number of shards to deploy in the specified zone.`,
				},
				resource.Attribute{
					Name:        "regions_config",
					Description: `Describes the physical location of the region. Each regionsConfig document describes the region’s priority in elections and the number and type of MongoDB nodes Atlas deploys to the region. You must order each regionsConfigs document by regionsConfig.priority, descending. See [Region Config](#region-config) below for more details.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Indicates the n ame for the zone in a Global Cluster. ### Region Config Physical location of the region.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Name for the region specified.`,
				},
				resource.Attribute{
					Name:        "electable_nodes",
					Description: `Number of electable nodes for Atlas to deploy to the region.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Election priority of the region. For regions with only read-only nodes, set this value to 0.`,
				},
				resource.Attribute{
					Name:        "read_only_nodes",
					Description: `Number of read-only nodes for Atlas to deploy to the region. Read-only nodes can never become the primary, but can facilitate local-reads. Specify 0 if you do not want any read-only nodes in the region.`,
				},
				resource.Attribute{
					Name:        "analytics_nodes",
					Description: `Indicates the number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. ### Labels Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. Note: the key ` + "`" + `Infrastructure Tool` + "`" + `, is used for internal purposes to track aggregate usage.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that was set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that represents the key. ### Plugin Contains a key-value pair that tags that the cluster was created by a Terraform Provider and notes the version.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the current plugin`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of the plugin. ### Cloud Provider Snapshot Backup Policy`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy",
					Description: `current snapshot schedule and retention settings for the cluster.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.cluster_id",
					Description: `Unique identifier of the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.cluster_name",
					Description: `Name of the Atlas cluster that contains the snapshot backup policy.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.next_snapshot",
					Description: `UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.reference_hour_of_day",
					Description: `UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.reference_minute_of_hour",
					Description: `UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.restore_window_days",
					Description: `Specifies a restore window in days for the cloud provider backup to maintain.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.update_snapshots",
					Description: `Specifies it's true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously. ### Policies`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies",
					Description: `A list of policy definitions for the cluster.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.id",
					Description: `Unique identifier of the backup policy. #### Policy Item`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item",
					Description: `A list of specifications for a policy.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.id",
					Description: `Unique identifier for this policy item.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.frequency_interval",
					Description: `The frequency interval for a set of snapshots.`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.frequency_type",
					Description: `A type of frequency (hourly, daily, weekly, monthly).`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.retention_unit",
					Description: `The unit of time in which snapshot retention is measured (days, weeks, months).`,
				},
				resource.Attribute{
					Name:        "snapshot_backup_policy.#.policies.#.policy_item.#.retention_value",
					Description: `The number of days, weeks, or months the snapshot is retained. See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_custom_db_roles",
			Category:         "Data Sources",
			ShortDescription: `Describes a Custom DB Roles.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to get all custom db roles. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a custom db roles. ### Actions Each object in the actions array represents an individual privilege action granted by the role. It is an required field.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Name of the privilege action. For a complete list of actions available in the Atlas API, see Custom Role Actions.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Contains information on where the action is granted. Each object in the array either indicates a database and collection on which the action is granted, or indicates that the action is granted on the cluster resource.`,
				},
				resource.Attribute{
					Name:        "resources.#.collection",
					Description: `(Optional) Collection on which the action is granted. If this value is an empty string, the action is granted on all collections within the database specified in the actions.resources.db field.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a custom db roles. ### Actions Each object in the actions array represents an individual privilege action granted by the role. It is an required field.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Name of the privilege action. For a complete list of actions available in the Atlas API, see Custom Role Actions.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Contains information on where the action is granted. Each object in the array either indicates a database and collection on which the action is granted, or indicates that the action is granted on the cluster resource.`,
				},
				resource.Attribute{
					Name:        "resources.#.collection",
					Description: `(Optional) Collection on which the action is granted. If this value is an empty string, the action is granted on all collections within the database specified in the actions.resources.db field.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_database_user",
			Category:         "Data Sources",
			ShortDescription: `Describes a Database User.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username for authenticating to MongoDB.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "auth_database_name",
					Description: `(Required) The user’s authentication database. A user must provide both a username and authentication database to log into MongoDB. In Atlas deployments of MongoDB, the authentication database is almost always the admin database, for X509 it is $external. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See [Roles](#roles) below for more details.`,
				},
				resource.Attribute{
					Name:        "x509_type",
					Description: `X.509 method by which the provided username is authenticated.`,
				},
				resource.Attribute{
					Name:        "aws_iam_type",
					Description: `The new database user authenticates with AWS IAM credentials. Default is ` + "`" + `NONE` + "`" + `, ` + "`" + `USER` + "`" + ` means user has AWS IAM user credentials, ` + "`" + `ROLE` + "`" + ` - means user has credentials associated with an AWS IAM role.`,
				},
				resource.Attribute{
					Name:        "ldap_auth_type",
					Description: `Method by which the provided username is authenticated. Default is ` + "`" + `NONE` + "`" + `. Other valid values are: ` + "`" + `USER` + "`" + `, ` + "`" + `GROUP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `Array of clusters and Atlas Data Lakes that this user has access to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster or Atlas Data Lake that the user has access to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of resource that the user has access to. Valid values are: ` + "`" + `CLUSTER` + "`" + ` and ` + "`" + `DATA_LAKE` + "`" + ` ### Roles Block mapping a user's role to a database / collection. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. ->`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the role to grant.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `Database on which the user has the specified role. A role on the ` + "`" + `admin` + "`" + ` database can include privileges that apply to the other databases.`,
				},
				resource.Attribute{
					Name:        "collection_name",
					Description: `Collection for which the role applies. You can specify a collection for the ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + ` roles. If you do not specify a collection for ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + `, the role applies to all collections in the database (excluding some collections in the ` + "`" + `system` + "`" + `. database). ### Labels Containing key-value pairs that tag and categorize the database user. Each key and value has a maximum length of 255 characters.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/database-users-get-single-user/) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See [Roles](#roles) below for more details.`,
				},
				resource.Attribute{
					Name:        "x509_type",
					Description: `X.509 method by which the provided username is authenticated.`,
				},
				resource.Attribute{
					Name:        "aws_iam_type",
					Description: `The new database user authenticates with AWS IAM credentials. Default is ` + "`" + `NONE` + "`" + `, ` + "`" + `USER` + "`" + ` means user has AWS IAM user credentials, ` + "`" + `ROLE` + "`" + ` - means user has credentials associated with an AWS IAM role.`,
				},
				resource.Attribute{
					Name:        "ldap_auth_type",
					Description: `Method by which the provided username is authenticated. Default is ` + "`" + `NONE` + "`" + `. Other valid values are: ` + "`" + `USER` + "`" + `, ` + "`" + `GROUP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `Array of clusters and Atlas Data Lakes that this user has access to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster or Atlas Data Lake that the user has access to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of resource that the user has access to. Valid values are: ` + "`" + `CLUSTER` + "`" + ` and ` + "`" + `DATA_LAKE` + "`" + ` ### Roles Block mapping a user's role to a database / collection. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. ->`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the role to grant.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `Database on which the user has the specified role. A role on the ` + "`" + `admin` + "`" + ` database can include privileges that apply to the other databases.`,
				},
				resource.Attribute{
					Name:        "collection_name",
					Description: `Collection for which the role applies. You can specify a collection for the ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + ` roles. If you do not specify a collection for ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + `, the role applies to all collections in the database (excluding some collections in the ` + "`" + `system` + "`" + `. database). ### Labels Containing key-value pairs that tag and categorize the database user. Each key and value has a maximum length of 255 characters.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/database-users-get-single-user/) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_database_users",
			Category:         "Data Sources",
			ShortDescription: `Describes a Database Users.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to get all database users. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Database user. ### Database User`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the Atlas project the user belongs to.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for authenticating to MongoDB.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See [Roles](#roles) below for more details.`,
				},
				resource.Attribute{
					Name:        "auth_database_name",
					Description: `(Required) Database against which Atlas authenticates the user. A user must provide both a username and authentication database to log into MongoDB. Possible values include:`,
				},
				resource.Attribute{
					Name:        "ldap_auth_type",
					Description: `is USER or GROUP.`,
				},
				resource.Attribute{
					Name:        "x509_type",
					Description: `X.509 method by which the provided username is authenticated.`,
				},
				resource.Attribute{
					Name:        "aws_iam_type",
					Description: `The new database user authenticates with AWS IAM credentials. Default is ` + "`" + `NONE` + "`" + `, ` + "`" + `USER` + "`" + ` means user has AWS IAM user credentials, ` + "`" + `ROLE` + "`" + ` - means user has credentials associated with an AWS IAM role.`,
				},
				resource.Attribute{
					Name:        "ldap_auth_type",
					Description: `Method by which the provided username is authenticated. Default is ` + "`" + `NONE` + "`" + `. Other valid values are: ` + "`" + `USER` + "`" + `, ` + "`" + `GROUP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `Array of clusters and Atlas Data Lakes that this user has access to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster or Atlas Data Lake that the user has access to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of resource that the user has access to. Valid values are: ` + "`" + `CLUSTER` + "`" + ` and ` + "`" + `DATA_LAKE` + "`" + ` ### Roles Block mapping a user's role to a database / collection. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. ->`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the role to grant.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `Database on which the user has the specified role. A role on the ` + "`" + `admin` + "`" + ` database can include privileges that apply to the other databases.`,
				},
				resource.Attribute{
					Name:        "collection_name",
					Description: `Collection for which the role applies. You can specify a collection for the ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + ` roles. If you do not specify a collection for ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + `, the role applies to all collections in the database (excluding some collections in the ` + "`" + `system` + "`" + `. database). ### Labels Containing key-value pairs that tag and categorize the database user. Each key and value has a maximum length of 255 characters.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/database-users-get-single-user/) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Database user. ### Database User`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the Atlas project the user belongs to.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for authenticating to MongoDB.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See [Roles](#roles) below for more details.`,
				},
				resource.Attribute{
					Name:        "auth_database_name",
					Description: `(Required) Database against which Atlas authenticates the user. A user must provide both a username and authentication database to log into MongoDB. Possible values include:`,
				},
				resource.Attribute{
					Name:        "ldap_auth_type",
					Description: `is USER or GROUP.`,
				},
				resource.Attribute{
					Name:        "x509_type",
					Description: `X.509 method by which the provided username is authenticated.`,
				},
				resource.Attribute{
					Name:        "aws_iam_type",
					Description: `The new database user authenticates with AWS IAM credentials. Default is ` + "`" + `NONE` + "`" + `, ` + "`" + `USER` + "`" + ` means user has AWS IAM user credentials, ` + "`" + `ROLE` + "`" + ` - means user has credentials associated with an AWS IAM role.`,
				},
				resource.Attribute{
					Name:        "ldap_auth_type",
					Description: `Method by which the provided username is authenticated. Default is ` + "`" + `NONE` + "`" + `. Other valid values are: ` + "`" + `USER` + "`" + `, ` + "`" + `GROUP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `Array of clusters and Atlas Data Lakes that this user has access to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster or Atlas Data Lake that the user has access to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of resource that the user has access to. Valid values are: ` + "`" + `CLUSTER` + "`" + ` and ` + "`" + `DATA_LAKE` + "`" + ` ### Roles Block mapping a user's role to a database / collection. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. ->`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the role to grant.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `Database on which the user has the specified role. A role on the ` + "`" + `admin` + "`" + ` database can include privileges that apply to the other databases.`,
				},
				resource.Attribute{
					Name:        "collection_name",
					Description: `Collection for which the role applies. You can specify a collection for the ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + ` roles. If you do not specify a collection for ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + `, the role applies to all collections in the database (excluding some collections in the ` + "`" + `system` + "`" + `. database). ### Labels Containing key-value pairs that tag and categorize the database user. Each key and value has a maximum length of 255 characters.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/database-users-get-single-user/) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_maintenance_window",
			Category:         "Data Sources",
			ShortDescription: `Provides a Maintenance Window Datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `The unique identifier of the project for the Maintenance Window. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "day_of_week",
					Description: `Day of the week when you would like the maintenance window to start as a 1-based integer: S=1, M=2, T=3, W=4, T=5, F=6, S=7.`,
				},
				resource.Attribute{
					Name:        "hour_of_day",
					Description: `Hour of the day when you would like the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0, noon is 12 (Time zone is UTC).`,
				},
				resource.Attribute{
					Name:        "start_asap",
					Description: `Flag indicating whether project maintenance has been directed to start immediately. If you request that maintenance begin immediately, this field returns true from the time the request was made until the time the maintenance event completes.`,
				},
				resource.Attribute{
					Name:        "number_of_deferrals",
					Description: `Number of times the current maintenance event for this project has been deferred, you can set a maximum of 2 deferrals. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/maintenance-windows/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "day_of_week",
					Description: `Day of the week when you would like the maintenance window to start as a 1-based integer: S=1, M=2, T=3, W=4, T=5, F=6, S=7.`,
				},
				resource.Attribute{
					Name:        "hour_of_day",
					Description: `Hour of the day when you would like the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0, noon is 12 (Time zone is UTC).`,
				},
				resource.Attribute{
					Name:        "start_asap",
					Description: `Flag indicating whether project maintenance has been directed to start immediately. If you request that maintenance begin immediately, this field returns true from the time the request was made until the time the maintenance event completes.`,
				},
				resource.Attribute{
					Name:        "number_of_deferrals",
					Description: `Number of times the current maintenance event for this project has been deferred, you can set a maximum of 2 deferrals. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/maintenance-windows/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_container",
			Category:         "Data Sources",
			ShortDescription: `Describes a Cluster resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `(Required) The Network Peering Container ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Container ID.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `CIDR block that Atlas uses for your clusters. Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3).`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `The Atlas AWS region name for where this container will exist.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The Atlas Azure region name for where this container will exist.`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `Unique identifer of the Azure subscription in which the VNet resides.`,
				},
				resource.Attribute{
					Name:        "provisioned",
					Description: `Indicates whether the project has Network Peering connections deployed in the container.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `Unique identifier of the GCP project in which the Network Peering connection resides.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the Network Peering connection in the Atlas project.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Unique identifier of the project’s VPC.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `The name of the Azure VNet. This value is null until you provision an Azure VNet in the container.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `Atlas GCP regions where the container resides. See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Container ID.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `CIDR block that Atlas uses for your clusters. Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3).`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `The Atlas AWS region name for where this container will exist.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The Atlas Azure region name for where this container will exist.`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `Unique identifer of the Azure subscription in which the VNet resides.`,
				},
				resource.Attribute{
					Name:        "provisioned",
					Description: `Indicates whether the project has Network Peering connections deployed in the container.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `Unique identifier of the GCP project in which the Network Peering connection resides.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the Network Peering connection in the Atlas project.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Unique identifier of the project’s VPC.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `The name of the Azure VNet. This value is null until you provision an Azure VNet in the container.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `Atlas GCP regions where the container resides. See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_containers",
			Category:         "Data Sources",
			ShortDescription: `Describes all Network Peering Containers in the project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) Cloud provider for this Network peering container. Accepted values are AWS, GCP, and Azure. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Network Peering Container. ### Network Peering Container`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Container ID.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `CIDR block that Atlas uses for your clusters. Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3).`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `The Atlas AWS region name for where this container exists.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The Atlas Azure region name for where this container exists.`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `Unique identifer of the Azure subscription in which the VNet resides.`,
				},
				resource.Attribute{
					Name:        "provisioned",
					Description: `Indicates whether the project has Network Peering connections deployed in the container.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `Unique identifier of the GCP project in which the Network Peering connection resides.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the Network Peering connection in the Atlas project.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Unique identifier of the project’s VPC.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `The name of the Azure VNet. This value is null until you provision an Azure VNet in the container.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `Atlas GCP regions where the container resides. See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-get-containers-list/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Network Peering Container. ### Network Peering Container`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Container ID.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `CIDR block that Atlas uses for your clusters. Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3).`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `The Atlas AWS region name for where this container exists.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The Atlas Azure region name for where this container exists.`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `Unique identifer of the Azure subscription in which the VNet resides.`,
				},
				resource.Attribute{
					Name:        "provisioned",
					Description: `Indicates whether the project has Network Peering connections deployed in the container.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `Unique identifier of the GCP project in which the Network Peering connection resides.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the Network Peering connection in the Atlas project.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Unique identifier of the project’s VPC.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `The name of the Azure VNet. This value is null until you provision an Azure VNet in the container.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `Atlas GCP regions where the container resides. See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-get-containers-list/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_peering",
			Category:         "Data Sources",
			ShortDescription: `Describes a Network Peering.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "peering_id",
					Description: `(Required) Atlas assigned unique ID for the peering connection. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Connection ID.`,
				},
				resource.Attribute{
					Name:        "connection_id",
					Description: `Unique identifier for the peering connection.`,
				},
				resource.Attribute{
					Name:        "accepter_region_name",
					Description: `Specifies the region where the peer VPC resides. For complete lists of supported regions, see [Amazon Web Services](https://docs.atlas.mongodb.com/reference/amazon-aws/).`,
				},
				resource.Attribute{
					Name:        "aws_account_id",
					Description: `Account ID of the owner of the peer VPC.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Cloud provider for this VPC peering connection. If omitted, Atlas sets this parameter to AWS. (Possible Values ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + `, ` + "`" + `GCP` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_table_cidr_block",
					Description: `Peer VPC CIDR block or subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Unique identifier of the peer VPC.`,
				},
				resource.Attribute{
					Name:        "error_state_name",
					Description: `Error state, if any. The VPC peering connection error state value can be one of the following: ` + "`" + `REJECTED` + "`" + `, ` + "`" + `EXPIRED` + "`" + `, ` + "`" + `INVALID_ARGUMENT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status_name",
					Description: `The VPC peering connection status value can be one of the following: ` + "`" + `INITIATING` + "`" + `, ` + "`" + `PENDING_ACCEPTANCE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `FINALIZING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `TERMINATING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `Unique identifier for an Azure AD directory.`,
				},
				resource.Attribute{
					Name:        "azure_directory_id",
					Description: `Unique identifier for an Azure AD directory.`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `Unique identifer of the Azure subscription in which the VNet resides.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Name of your Azure resource group.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `Name of your Azure VNet.`,
				},
				resource.Attribute{
					Name:        "error_state",
					Description: `Description of the Atlas error when ` + "`" + `status` + "`" + ` is ` + "`" + `Failed` + "`" + `, Otherwise, Atlas returns ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Atlas network peering connection: ` + "`" + `ADDING_PEER` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `WAITING_FOR_USER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `GCP project ID of the owner of the network peer.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the network peer to which Atlas connects.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `When ` + "`" + `"status" : "FAILED"` + "`" + `, Atlas provides a description of the error. See detailed information for arguments and attributes: [MongoDB API Network Peering Connection](https://docs.atlas.mongodb.com/reference/api/vpc-get-connection/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Connection ID.`,
				},
				resource.Attribute{
					Name:        "connection_id",
					Description: `Unique identifier for the peering connection.`,
				},
				resource.Attribute{
					Name:        "accepter_region_name",
					Description: `Specifies the region where the peer VPC resides. For complete lists of supported regions, see [Amazon Web Services](https://docs.atlas.mongodb.com/reference/amazon-aws/).`,
				},
				resource.Attribute{
					Name:        "aws_account_id",
					Description: `Account ID of the owner of the peer VPC.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Cloud provider for this VPC peering connection. If omitted, Atlas sets this parameter to AWS. (Possible Values ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + `, ` + "`" + `GCP` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_table_cidr_block",
					Description: `Peer VPC CIDR block or subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Unique identifier of the peer VPC.`,
				},
				resource.Attribute{
					Name:        "error_state_name",
					Description: `Error state, if any. The VPC peering connection error state value can be one of the following: ` + "`" + `REJECTED` + "`" + `, ` + "`" + `EXPIRED` + "`" + `, ` + "`" + `INVALID_ARGUMENT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status_name",
					Description: `The VPC peering connection status value can be one of the following: ` + "`" + `INITIATING` + "`" + `, ` + "`" + `PENDING_ACCEPTANCE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `FINALIZING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `TERMINATING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `Unique identifier for an Azure AD directory.`,
				},
				resource.Attribute{
					Name:        "azure_directory_id",
					Description: `Unique identifier for an Azure AD directory.`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `Unique identifer of the Azure subscription in which the VNet resides.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Name of your Azure resource group.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `Name of your Azure VNet.`,
				},
				resource.Attribute{
					Name:        "error_state",
					Description: `Description of the Atlas error when ` + "`" + `status` + "`" + ` is ` + "`" + `Failed` + "`" + `, Otherwise, Atlas returns ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Atlas network peering connection: ` + "`" + `ADDING_PEER` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `WAITING_FOR_USER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `GCP project ID of the owner of the network peer.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the network peer to which Atlas connects.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `When ` + "`" + `"status" : "FAILED"` + "`" + `, Atlas provides a description of the error. See detailed information for arguments and attributes: [MongoDB API Network Peering Connection](https://docs.atlas.mongodb.com/reference/api/vpc-get-connection/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_peerings",
			Category:         "Data Sources",
			ShortDescription: `Describes all Network Peering Connections.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Connection ID.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Network Peering Connection. ### Network Peering Connection`,
				},
				resource.Attribute{
					Name:        "peering_id",
					Description: `Atlas assigned unique ID for the peering connection.`,
				},
				resource.Attribute{
					Name:        "connection_id",
					Description: `Unique identifier for the peering connection.`,
				},
				resource.Attribute{
					Name:        "accepter_region_name",
					Description: `Specifies the region where the peer VPC resides. For complete lists of supported regions, see [Amazon Web Services](https://docs.atlas.mongodb.com/reference/amazon-aws/).`,
				},
				resource.Attribute{
					Name:        "aws_account_id",
					Description: `Account ID of the owner of the peer VPC.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Cloud provider for this VPC peering connection. If omitted, Atlas sets this parameter to AWS. (Possible Values ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + `, ` + "`" + `GCP` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_table_cidr_block",
					Description: `Peer VPC CIDR block or subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Unique identifier of the peer VPC.`,
				},
				resource.Attribute{
					Name:        "error_state_name",
					Description: `Error state, if any. The VPC peering connection error state value can be one of the following: ` + "`" + `REJECTED` + "`" + `, ` + "`" + `EXPIRED` + "`" + `, ` + "`" + `INVALID_ARGUMENT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status_name",
					Description: `The VPC peering connection status value can be one of the following: ` + "`" + `INITIATING` + "`" + `, ` + "`" + `PENDING_ACCEPTANCE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `FINALIZING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `TERMINATING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `Unique identifier for an Azure AD directory.`,
				},
				resource.Attribute{
					Name:        "azure_directory_id",
					Description: `Unique identifier for an Azure AD directory.`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `Unique identifer of the Azure subscription in which the VNet resides.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Name of your Azure resource group.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `Name of your Azure VNet.`,
				},
				resource.Attribute{
					Name:        "error_state",
					Description: `Description of the Atlas error when ` + "`" + `status` + "`" + ` is ` + "`" + `Failed` + "`" + `, Otherwise, Atlas returns ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Atlas network peering connection: ` + "`" + `ADDING_PEER` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `WAITING_FOR_USER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `GCP project ID of the owner of the network peer.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the network peer to which Atlas connects.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `When ` + "`" + `"status" : "FAILED"` + "`" + `, Atlas provides a description of the error. See detailed information for arguments and attributes: [MongoDB API Network Peering Connection](https://docs.atlas.mongodb.com/reference/api/vpc-get-connections-list/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Connection ID.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Network Peering Connection. ### Network Peering Connection`,
				},
				resource.Attribute{
					Name:        "peering_id",
					Description: `Atlas assigned unique ID for the peering connection.`,
				},
				resource.Attribute{
					Name:        "connection_id",
					Description: `Unique identifier for the peering connection.`,
				},
				resource.Attribute{
					Name:        "accepter_region_name",
					Description: `Specifies the region where the peer VPC resides. For complete lists of supported regions, see [Amazon Web Services](https://docs.atlas.mongodb.com/reference/amazon-aws/).`,
				},
				resource.Attribute{
					Name:        "aws_account_id",
					Description: `Account ID of the owner of the peer VPC.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Cloud provider for this VPC peering connection. If omitted, Atlas sets this parameter to AWS. (Possible Values ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + `, ` + "`" + `GCP` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_table_cidr_block",
					Description: `Peer VPC CIDR block or subnet.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Unique identifier of the peer VPC.`,
				},
				resource.Attribute{
					Name:        "error_state_name",
					Description: `Error state, if any. The VPC peering connection error state value can be one of the following: ` + "`" + `REJECTED` + "`" + `, ` + "`" + `EXPIRED` + "`" + `, ` + "`" + `INVALID_ARGUMENT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status_name",
					Description: `The VPC peering connection status value can be one of the following: ` + "`" + `INITIATING` + "`" + `, ` + "`" + `PENDING_ACCEPTANCE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `FINALIZING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `TERMINATING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `Unique identifier for an Azure AD directory.`,
				},
				resource.Attribute{
					Name:        "azure_directory_id",
					Description: `Unique identifier for an Azure AD directory.`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `Unique identifer of the Azure subscription in which the VNet resides.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Name of your Azure resource group.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `Name of your Azure VNet.`,
				},
				resource.Attribute{
					Name:        "error_state",
					Description: `Description of the Atlas error when ` + "`" + `status` + "`" + ` is ` + "`" + `Failed` + "`" + `, Otherwise, Atlas returns ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Atlas network peering connection: ` + "`" + `ADDING_PEER` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `WAITING_FOR_USER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `GCP project ID of the owner of the network peer.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the network peer to which Atlas connects.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `When ` + "`" + `"status" : "FAILED"` + "`" + `, Atlas provides a description of the error. See detailed information for arguments and attributes: [MongoDB API Network Peering Connection](https://docs.atlas.mongodb.com/reference/api/vpc-get-connections-list/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_project",
			Category:         "Data Sources",
			ShortDescription: `Describes a Project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The unique ID for the project.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The unique ID for the project. ~>`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the project you want to create. (Cannot be changed via this Provider after creation.)`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The ID of the organization you want to create the project within.`,
				},
				resource.Attribute{
					Name:        "teams.#.team_id",
					Description: `The unique identifier of the team you want to associate with the project. The team and project must share the same parent organization.`,
				},
				resource.Attribute{
					Name:        "teams.#.role_names",
					Description: `Each string in the array represents a project role assigned to the team. Every user associated with the team inherits these roles. The following are valid roles:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the project you want to create. (Cannot be changed via this Provider after creation.)`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The ID of the organization you want to create the project within.`,
				},
				resource.Attribute{
					Name:        "teams.#.team_id",
					Description: `The unique identifier of the team you want to associate with the project. The team and project must share the same parent organization.`,
				},
				resource.Attribute{
					Name:        "teams.#.role_names",
					Description: `Each string in the array represents a project role assigned to the team. Every user associated with the team inherits these roles. The following are valid roles:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_projects",
			Category:         "Data Sources",
			ShortDescription: `Describes a Projects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Represents the total of the projects`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Projects. ### Project`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the project you want to create. (Cannot be changed via this Provider after creation.)`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The ID of the organization you want to create the project within.`,
				},
				resource.Attribute{
					Name:        "teams.#.team_id",
					Description: `The unique identifier of the team you want to associate with the project. The team and project must share the same parent organization.`,
				},
				resource.Attribute{
					Name:        "teams.#.role_names",
					Description: `Each string in the array represents a project role assigned to the team. Every user associated with the team inherits these roles. The following are valid roles:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_team",
			Category:         "Data Sources",
			ShortDescription: `Describes a Team.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) The unique identifier for the organization you want to associate the team with.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `(Optional) The unique identifier for the team.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The team name. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The unique identifier for the team.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the team you want to create.`,
				},
				resource.Attribute{
					Name:        "usernames",
					Description: `The users who are part of the organization. See detailed information for arguments and attributes: [MongoDB API Teams](https://docs.atlas.mongodb.com/reference/api/teams-create-one/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The unique identifier for the team.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the team you want to create.`,
				},
				resource.Attribute{
					Name:        "usernames",
					Description: `The users who are part of the organization. See detailed information for arguments and attributes: [MongoDB API Teams](https://docs.atlas.mongodb.com/reference/api/teams-create-one/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_x509_authentication_database_user",
			Category:         "Data Sources",
			ShortDescription: `Describes a Custom DB Role.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Identifier for the Atlas project associated with the X.509 configuration.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Username of the database user to create a certificate for. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "current_certificate",
					Description: `Contains the last X.509 certificate and private key created for a database user. #### Certificates`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `Array of objects where each details one unexpired database user certificate.`,
				},
				resource.Attribute{
					Name:        "certificates.#.id",
					Description: `Serial number of this certificate.`,
				},
				resource.Attribute{
					Name:        "certificates.#.created_at",
					Description: `Timestamp in ISO 8601 date and time format in UTC when Atlas created this X.509 certificate.`,
				},
				resource.Attribute{
					Name:        "certificates.#.group_id",
					Description: `Unique identifier of the Atlas project to which this certificate belongs.`,
				},
				resource.Attribute{
					Name:        "certificates.#.not_after",
					Description: `Timestamp in ISO 8601 date and time format in UTC when this certificate expires.`,
				},
				resource.Attribute{
					Name:        "certificates.#.subject",
					Description: `Fully distinguished name of the database user to which this certificate belongs. To learn more, see RFC 2253. See [MongoDB Atlas - X509 User Certificates](https://docs.atlas.mongodb.com/reference/api/x509-configuration-get-certificates/) and [MongoDB Atlas - Current X509 Configuratuion](https://docs.atlas.mongodb.com/reference/api/x509-configuration-get-current/) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "current_certificate",
					Description: `Contains the last X.509 certificate and private key created for a database user. #### Certificates`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `Array of objects where each details one unexpired database user certificate.`,
				},
				resource.Attribute{
					Name:        "certificates.#.id",
					Description: `Serial number of this certificate.`,
				},
				resource.Attribute{
					Name:        "certificates.#.created_at",
					Description: `Timestamp in ISO 8601 date and time format in UTC when Atlas created this X.509 certificate.`,
				},
				resource.Attribute{
					Name:        "certificates.#.group_id",
					Description: `Unique identifier of the Atlas project to which this certificate belongs.`,
				},
				resource.Attribute{
					Name:        "certificates.#.not_after",
					Description: `Timestamp in ISO 8601 date and time format in UTC when this certificate expires.`,
				},
				resource.Attribute{
					Name:        "certificates.#.subject",
					Description: `Fully distinguished name of the database user to which this certificate belongs. To learn more, see RFC 2253. See [MongoDB Atlas - X509 User Certificates](https://docs.atlas.mongodb.com/reference/api/x509-configuration-get-certificates/) and [MongoDB Atlas - Current X509 Configuratuion](https://docs.atlas.mongodb.com/reference/api/x509-configuration-get-current/) Documentation for more information.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"mongodbatlas_alert_configuration":                   0,
		"mongodbatlas_auditing":                              1,
		"mongodbatlas_cloud_provider_snapshot":               2,
		"mongodbatlas_cloud_provider_snapshot_backup_policy": 3,
		"mongodbatlas_cloud_provider_snapshot_restore_job":   4,
		"mongodbatlas_cloud_provider_snapshot_restore_jobs":  5,
		"mongodbatlas_cloud_provider_snapshots":              6,
		"mongodbatlas_cluster":                               7,
		"mongodbatlas_clusters":                              8,
		"mongodbatlas_custom_db_roles":                       9,
		"mongodbatlas_database_user":                         10,
		"mongodbatlas_database_users":                        11,
		"mongodbatlas_maintenance_window":                    12,
		"mongodbatlas_network_container":                     13,
		"mongodbatlas_network_containers":                    14,
		"mongodbatlas_network_peering":                       15,
		"mongodbatlas_network_peerings":                      16,
		"mongodbatlas_project":                               17,
		"mongodbatlas_projects":                              18,
		"mongodbatlas_team":                                  19,
		"mongodbatlas_x509_authentication_database_user":     20,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
