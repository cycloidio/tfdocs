package lacework

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "lacework_agent_access_token",
			Category:         "Agents",
			ShortDescription: `Manage agent access tokens`,
			Description:      ``,
			Keywords: []string{
				"agents",
				"agent",
				"access",
				"token",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The agent access token name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The agent access token description. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The agent access token. ## Import A Lacework agent access token can be imported using the token itself, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_agent_access_token.k8s YourAgentToken ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "token",
					Description: `The agent access token. ## Import A Lacework agent access token can be imported using the token itself, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_agent_access_token.k8s YourAgentToken ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_aws_cloudwatch",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage AWS CloudWatch Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"aws",
				"cloudwatch",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "event_bus_arn",
					Description: `(Required) The ARN of your AWS CloudWatch event bus.`,
				},
				resource.Attribute{
					Name:        "group_issues_by",
					Description: `(Optional) Defines how Lacework compliance events get grouped. Must be one of ` + "`" + `Events` + "`" + ` or ` + "`" + `Resources` + "`" + `. Defaults to ` + "`" + `Events` + "`" + `. The available options are:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework AWS CloudWatch Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_aws_cloudwatch.all_events EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_aws_s3",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage AWS S3 Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"aws",
				"s3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "bucket_arn",
					Description: `(Required) The ARN of the S3 bucket.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Required) The ARN of the IAM role.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Required) The external ID for the IAM role. ## Import A Lacework AWS S3 Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_aws_s3.data_export EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_cisco_webex",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage Cisco Webex Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"cisco",
				"webex",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `(Required) The Cisco Webex webhook URL.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Cisco Webex Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_cisco_webex.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_datadog",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage Datadog Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"datadog",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `(Required) The Datadog api key required to submit metrics and events to Datadog`,
				},
				resource.Attribute{
					Name:        "datadog_site",
					Description: `(Optional) Where to store your logs, either the US or Europe. Must be one of ` + "`" + `com` + "`" + ` or ` + "`" + `eu` + "`" + `. Defaults to ` + "`" + `com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datadog_service",
					Description: `(Optional) The level of detail of logs or event stream. ` + "`" + `Logs Detail` + "`" + `, ` + "`" + `Logs Summary` + "`" + `, or ` + "`" + `Events Summary` + "`" + `. Defaults to ` + "`" + `Logs Detail` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Datadog Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_datadog.ops_critical EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_gcp_pub_sub",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage GCP Pub Sub Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"gcp",
				"pub",
				"sub",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The name of the Google Cloud Project.`,
				},
				resource.Attribute{
					Name:        "topic_id",
					Description: `(Required) The ID of the Google Cloud Pub/Sub topic.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "issue_grouping",
					Description: `(Optional) Defines how Lacework compliance events get grouped. Must be one of ` + "`" + `Events` + "`" + ` or ` + "`" + `Resources` + "`" + `. Defaults to ` + "`" + `Events` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) The service account client ID.`,
				},
				resource.Attribute{
					Name:        "client_email",
					Description: `(Required) The service account client email.`,
				},
				resource.Attribute{
					Name:        "private_key_id",
					Description: `(Required) The service account private key ID.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The service account private key. ## Import A Lacework GCP Pub Sub Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_gcp_pub_sub.data_export EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_jira_cloud",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage Jira Cloud Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"jira",
				"cloud",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "jira_url",
					Description: `(Required) The URL of your Jira implementation without https protocol (` + "`" + `https://` + "`" + `). For example, ` + "`" + `mycompany.atlassian.net` + "`" + ` or ` + "`" + `mycompany.jira.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "issue_type",
					Description: `(Required) The Jira Issue type (such as a ` + "`" + `Bug` + "`" + `) to create when a new Jira issue is created.`,
				},
				resource.Attribute{
					Name:        "project_key",
					Description: `(Required) The project key for the Jira project where the new Jira issues should be created.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The Jira user name. Lacework recommends a dedicated Jira user. See above for more information.`,
				},
				resource.Attribute{
					Name:        "api_token",
					Description: `(Required) The Jira API Token. For more information, see [how to create a Jira API Token](https://confluence.atlassian.com/cloud/api-tokens-938839638.html).`,
				},
				resource.Attribute{
					Name:        "group_issues_by",
					Description: `(Optional) Defines how Lacework compliance events get grouped. Must be one of ` + "`" + `Events` + "`" + ` or ` + "`" + `Resources` + "`" + `. Defaults to ` + "`" + `Events` + "`" + `. The available options are:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_template_file",
					Description: `(Optional) A Custom Template JSON file to populate fields in the new Jira issues. ## Import A Lacework Jira Cloud Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_jira_cloud.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_jira_server",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage Jira Server Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"jira",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "jira_url",
					Description: `(Required) The URL of your Jira implementation without https protocol (` + "`" + `https://` + "`" + `). For example, ` + "`" + `mycompany.atlassian.net` + "`" + ` or ` + "`" + `mycompany.jira.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "issue_type",
					Description: `(Required) The Jira Issue type (such as a ` + "`" + `Bug` + "`" + `) to create when a new Jira issue is created.`,
				},
				resource.Attribute{
					Name:        "project_key",
					Description: `(Required) The project key for the Jira project where the new Jira issues should be created.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The Jira user name. Lacework recommends a dedicated Jira user. See above for more information.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password to the Jira user.`,
				},
				resource.Attribute{
					Name:        "group_issues_by",
					Description: `(Optional) Defines how Lacework compliance events get grouped. Must be one of ` + "`" + `Events` + "`" + ` or ` + "`" + `Resources` + "`" + `. Defaults to ` + "`" + `Events` + "`" + `. The available options are:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_template_file",
					Description: `(Optional) A Custom Template JSON file to populate fields in the new Jira issues. ## Import A Lacework Jira Server Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_jira_server.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_microsoft_teams",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage Microsoft Teams Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"microsoft",
				"teams",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `(Required) The URL of your Microsoft Teams incoming webhook.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Webhook Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_microsoft_teams.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_newrelic",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage New Relic Insights Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"newrelic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) The New Relic account ID.`,
				},
				resource.Attribute{
					Name:        "insert_key",
					Description: `(Required) The New Relic Insert API key.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework New Relic Insights Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_newrelic.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_pagerduty",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage PagerDuty Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"pagerduty",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Required) The PagerDuty service integration key.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework PagerDuty Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_pagerduty.critical EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_qradar",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage IBM QRadar Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"qradar",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "host_url",
					Description: `(Required) The domain name or IP address of QRadar.`,
				},
				resource.Attribute{
					Name:        "host_port",
					Description: `(Required) The listen port defined in QRadar.`,
				},
				resource.Attribute{
					Name:        "communication_type",
					Description: `(Required) The communication protocol used. Must be one of ` + "`" + `HTTPS` + "`" + ` or ` + "`" + `HTTPS Self Signed Cert` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework IBM QRadar Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_qradar.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_service_now",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage Service Now Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"service",
				"now",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "instance_url",
					Description: `(Required) The ServiceNow instance URL.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The ServiceNow user name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The ServiceNow password.`,
				},
				resource.Attribute{
					Name:        "custom_template_file",
					Description: `(Optional) Populate fields in the ServiceNow incident with values from a custom template JSON file.`,
				},
				resource.Attribute{
					Name:        "issue_grouping",
					Description: `(Optional) Defines how Lacework compliance events get grouped. Must be one of ` + "`" + `Events` + "`" + ` or ` + "`" + `Resources` + "`" + `. Defaults to ` + "`" + `Events` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Service Now Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_service_now.ops_critical EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_slack",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage Slack Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"slack",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "slack_url",
					Description: `(Required) The URL of the incoming Slack webhook.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Slack Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_slack.ops_critical EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_splunk",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage Splunk Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"splunk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "hec_token",
					Description: `(Required) The token you generate when you create a new HEC input.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) The hostname of the client from which you're sending data.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The destination port for forwarding events.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `(Optional) Enable or Disable SSL.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Optional) The Splunk channel name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ### Event Data ` + "`" + `event_data` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The Splunk source.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Required) Index to store generated events. ## Import A Lacework Splunk Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_splunk.ops_critical EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_victorops",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage VictorOps Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"victorops",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `(Required) The URL of your VictorOps webhook that will receive the HTTP POST.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework VictorOps Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_victorops.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_webhook",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage Webhook Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"webhook",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `(Required) The URL of your webhook that will receive the HTTP POST.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Webhook Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_webhook.ops_critical EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_aws_cfg",
			Category:         "Cloud Account Integrations",
			ShortDescription: `Create and manage AWS Config integrations`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"account",
				"integrations",
				"integration",
				"aws",
				"cfg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The AWS Config integration name.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) The number of attempts to create the external integration. Defaults to ` + "`" + `5` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Required) The ARN of the IAM role.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Required) The external ID for the IAM role. ## Import A Lacework AWS Config integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_aws_cfg.account_abc EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_aws_ct",
			Category:         "Cloud Account Integrations",
			ShortDescription: `Create and manage AWS CloudTrail integrations`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"account",
				"integrations",
				"integration",
				"aws",
				"ct",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The AWS CloudTrail integration name.`,
				},
				resource.Attribute{
					Name:        "queue_url",
					Description: `(Required) The SQS Queue URL.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) The number of attempts to create the external integration. Defaults to ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "org_account_mappings",
					Description: `(Optional) Mapping of AWS accounts to Lacework accounts within a Lacework organization. See [Account Mappings](#organization-account-mappings) below for details. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_azure_al",
			Category:         "Cloud Account Integrations",
			ShortDescription: `Create and manage Azure Cloud Activity Log integrations`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"account",
				"integrations",
				"integration",
				"azure",
				"al",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Azure Activity Log integration name.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) The directory tenant ID.`,
				},
				resource.Attribute{
					Name:        "queue_url",
					Description: `(Required) The storage queue URL.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) The number of attempts to create the external integration. Defaults to ` + "`" + `5` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) The application client ID.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Required) The client secret. ## Import A Lacework Azure Activity Log integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_azure_at.account_abc EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_azure_cfg",
			Category:         "Cloud Account Integrations",
			ShortDescription: `Create and manage Azure Cloud Config integrations`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"account",
				"integrations",
				"integration",
				"azure",
				"cfg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Azure Config integration name.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) The directory tenant ID.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) The number of attempts to create the external integration. Defaults to ` + "`" + `5` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) The application client ID.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Required) The client secret. ## Import A Lacework Azure Config integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_azure_cfg.account_abc EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_docker_hub",
			Category:         "Container Registry Integrations",
			ShortDescription: `Create and manage Docker Hub container registry integrations`,
			Description:      ``,
			Keywords: []string{
				"container",
				"registry",
				"integrations",
				"integration",
				"docker",
				"hub",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Container Registry integration name.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The Docker user that has at least read-only permissions to the Docker Hub container repositories.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password for the specified Docker Hub user.`,
				},
				resource.Attribute{
					Name:        "limit_by_tag",
					Description: `(Optional) An image tag to limit the assessment of images with matching tag. If you specify ` + "`" + `limit_by_tag` + "`" + ` and ` + "`" + `limit_by_label` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `. Supported field input are ` + "`" + `mytext`,
				},
				resource.Attribute{
					Name:        "limit_by_label",
					Description: `(Optional) An image label to limit the assessment of images with matching label. If you specify ` + "`" + `limit_by_tag` + "`" + ` and ` + "`" + `limit_by_label` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `. Supported field input are ` + "`" + `mytext`,
				},
				resource.Attribute{
					Name:        "limit_by_repos",
					Description: `(Optional) A comma-separated list of repositories to assess. (without spaces recommended)`,
				},
				resource.Attribute{
					Name:        "limit_num_imgs",
					Description: `(Optional) The maximum number of newest container images to assess per repository. Must be one of ` + "`" + `5` + "`" + `, ` + "`" + `10` + "`" + `, or ` + "`" + `15` + "`" + `. Defaults to ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Docker Hub container registry integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_docker_hub.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_docker_v2",
			Category:         "Container Registry Integrations",
			ShortDescription: `Create and manage Docker V2 container registry integrations`,
			Description:      ``,
			Keywords: []string{
				"container",
				"registry",
				"integrations",
				"integration",
				"docker",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Docker V2 Registry integration name.`,
				},
				resource.Attribute{
					Name:        "registry_domain",
					Description: `(Required) The registry domain. Allowed formats are ` + "`" + `YourIP:YourPort` + "`" + ` or ` + "`" + `YourDomain:YourPort` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The user that has at permissions to pull from the container registry the images to be assessed.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password for the specified user.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `(Optional) Enable or disable SSL communication. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_tag",
					Description: `(Optional) An image tag to limit the assessment of images with matching tag. If you specify ` + "`" + `limit_by_tag` + "`" + ` and ` + "`" + `limit_by_label` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `. Supported field input are ` + "`" + `mytext`,
				},
				resource.Attribute{
					Name:        "limit_by_label",
					Description: `(Optional) An image label to limit the assessment of images with matching label. If you specify ` + "`" + `limit_by_tag` + "`" + ` and ` + "`" + `limit_by_label` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `. Supported field input are ` + "`" + `mytext`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Docker V2 container registry integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_docker_v2.jfrog EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_ecr",
			Category:         "Container Registry Integrations",
			ShortDescription: `Create and manage ECR integrations`,
			Description:      ``,
			Keywords: []string{
				"container",
				"registry",
				"integrations",
				"integration",
				"ecr",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The ECR integration name.`,
				},
				resource.Attribute{
					Name:        "registry_domain",
					Description: `(Required) The Amazon Container Registry (ECR) domain in the format ` + "`" + `YourAWSAccount.dkr.ecr.YourRegion.amazonaws.com` + "`" + `, where ` + "`" + `YourAWSAcount` + "`" + ` is the AWS account number for the AWS IAM user that has a role with permissions to access the ECR and ` + "`" + `YourRegion` + "`" + ` is your AWS region such as ` + "`" + `us-west-2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "limit_by_tag",
					Description: `(Optional) An image tag to limit the assessment of images with matching tag. If you specify ` + "`" + `limit_by_tag` + "`" + ` and ` + "`" + `limit_by_label` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `. Supported field input are ` + "`" + `mytext`,
				},
				resource.Attribute{
					Name:        "limit_by_label",
					Description: `(Optional) An image label to limit the assessment of images with matching label. If you specify ` + "`" + `limit_by_tag` + "`" + ` and ` + "`" + `limit_by_label` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `. Supported field input are ` + "`" + `mytext`,
				},
				resource.Attribute{
					Name:        "limit_by_repos",
					Description: `(Optional) A comma-separated list of repositories to assess. (without spaces recommended)`,
				},
				resource.Attribute{
					Name:        "limit_num_imgs",
					Description: `(Optional) The maximum number of newest container images to assess per repository. Must be one of ` + "`" + `5` + "`" + `, ` + "`" + `10` + "`" + `, or ` + "`" + `15` + "`" + `. Defaults to ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the combination of the following arguments.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `The ARN of the IAM role with permissions to access the Amazon Container Registry (ECR).`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `The external ID for the IAM role.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `The AWS access key ID for an AWS IAM user that has a role with permissions to access the Amazon Container Registry (ECR).`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `The AWS secret key for the specified AWS access key. ## Import A Lacework ECR integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_ecr.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_gcp_at",
			Category:         "Cloud Account Integrations",
			ShortDescription: `Create and manage Google Cloud Audit Trail integrations`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"account",
				"integrations",
				"integration",
				"gcp",
				"at",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The GCP Audit Trail integration name.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The organization or project ID.`,
				},
				resource.Attribute{
					Name:        "subscription",
					Description: `(Required) The subscription queue name.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "resource_level",
					Description: `(Optional) The integration level. Must be one of ` + "`" + `PROJECT` + "`" + ` or ` + "`" + `ORGANIZATION` + "`" + `. Defaults to ` + "`" + `PROJECT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) The number of attempts to create the external integration. Defaults to ` + "`" + `5` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) The service account client ID.`,
				},
				resource.Attribute{
					Name:        "client_email",
					Description: `(Required) The service account client email.`,
				},
				resource.Attribute{
					Name:        "private_key_id",
					Description: `(Required) The service account private key ID.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The service account private key. ## Import A Lacework GCP Audit Trail integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_gcp_at.account_abc EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_gcp_cfg",
			Category:         "Cloud Account Integrations",
			ShortDescription: `Create and manage Google Cloud Config integrations`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"account",
				"integrations",
				"integration",
				"gcp",
				"cfg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The GCP Config integration name.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The organization or project ID.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "resource_level",
					Description: `(Optional) The integration level. Must be one of ` + "`" + `PROJECT` + "`" + ` or ` + "`" + `ORGANIZATION` + "`" + `. Defaults to ` + "`" + `PROJECT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) The number of attempts to create the external integration. Defaults to ` + "`" + `5` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) The service account client ID.`,
				},
				resource.Attribute{
					Name:        "client_email",
					Description: `(Required) The service account client email.`,
				},
				resource.Attribute{
					Name:        "private_key_id",
					Description: `(Required) The service account private key ID.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The service account private key. ## Import A Lacework GCP Config integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_gcp_cfg.account_abc EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_gcr",
			Category:         "Container Registry Integrations",
			ShortDescription: `Create and manage GCR integrations`,
			Description:      ``,
			Keywords: []string{
				"container",
				"registry",
				"integrations",
				"integration",
				"gcr",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The GCR integration name.`,
				},
				resource.Attribute{
					Name:        "registry_domain",
					Description: `(Required) The GCR domain, which specifies the location where you store the images. Supported domains are ` + "`" + `gcr.io` + "`" + `, ` + "`" + `us.gcr.io` + "`" + `, ` + "`" + `eu.gcr.io` + "`" + `, or ` + "`" + `asia.gcr.io` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "limit_by_tag",
					Description: `(Optional) An image tag to limit the assessment of images with matching tag. If you specify ` + "`" + `limit_by_tag` + "`" + ` and ` + "`" + `limit_by_label` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `. Supported field input are ` + "`" + `mytext`,
				},
				resource.Attribute{
					Name:        "limit_by_label",
					Description: `(Optional) An image label to limit the assessment of images with matching label. If you specify ` + "`" + `limit_by_tag` + "`" + ` and ` + "`" + `limit_by_label` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `. Supported field input are ` + "`" + `mytext`,
				},
				resource.Attribute{
					Name:        "limit_by_repos",
					Description: `(Optional) A comma-separated list of repositories to assess. (without spaces recommended)`,
				},
				resource.Attribute{
					Name:        "limit_num_imgs",
					Description: `(Optional) The maximum number of newest container images to assess per repository. Must be one of ` + "`" + `5` + "`" + `, ` + "`" + `10` + "`" + `, or ` + "`" + `15` + "`" + `. Defaults to ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) The service account client ID.`,
				},
				resource.Attribute{
					Name:        "client_email",
					Description: `(Required) The service account client email.`,
				},
				resource.Attribute{
					Name:        "private_key_id",
					Description: `(Required) The service account private key ID.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The service account private key. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"lacework_agent_access_token":            0,
		"lacework_alert_channel_aws_cloudwatch":  1,
		"lacework_alert_channel_aws_s3":          2,
		"lacework_alert_channel_cisco_webex":     3,
		"lacework_alert_channel_datadog":         4,
		"lacework_alert_channel_gcp_pub_sub":     5,
		"lacework_alert_channel_jira_cloud":      6,
		"lacework_alert_channel_jira_server":     7,
		"lacework_alert_channel_microsoft_teams": 8,
		"lacework_alert_channel_newrelic":        9,
		"lacework_alert_channel_pagerduty":       10,
		"lacework_alert_channel_qradar":          11,
		"lacework_alert_channel_service_now":     12,
		"lacework_alert_channel_slack":           13,
		"lacework_alert_channel_splunk":          14,
		"lacework_alert_channel_victorops":       15,
		"lacework_alert_channel_webhook":         16,
		"lacework_integration_aws_cfg":           17,
		"lacework_integration_aws_ct":            18,
		"lacework_integration_azure_al":          19,
		"lacework_integration_azure_cfg":         20,
		"lacework_integration_docker_hub":        21,
		"lacework_integration_docker_v2":         22,
		"lacework_integration_ecr":               23,
		"lacework_integration_gcp_at":            24,
		"lacework_integration_gcp_cfg":           25,
		"lacework_integration_gcr":               26,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
