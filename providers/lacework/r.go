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
			ShortDescription: `Create and manage Amazon CloudWatch Alert Channel integrations`,
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
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Amazon CloudWatch Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_aws_cloudwatch.all_events EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_aws_s3",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage Amazon S3 Alert Channel integrations`,
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
					Description: `(Required) The external ID for the IAM role. ## Import A Lacework Amazon S3 Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_aws_s3.data_export EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
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
			Type:             "lacework_alert_channel_email",
			Category:         "Alert Channels",
			ShortDescription: `Create and manage Email Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"channels",
				"channel",
				"email",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Alert Channel integration name.`,
				},
				resource.Attribute{
					Name:        "recipients",
					Description: `(Required) The list of email addresses that you want to receive the alerts.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Email Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_email.auditors EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
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
			Type:             "lacework_alert_profile",
			Category:         "Alert Profiles",
			ShortDescription: `Create and manage Lacework Alert Profiles`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"profiles",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The alert profile name, uniquely identifies the profile. Cannot start 'LW_' which is reserved for Lacework profiles.`,
				},
				resource.Attribute{
					Name:        "extends",
					Description: `(Required) The name of existing alert profile from which this profile extends.`,
				},
				resource.Attribute{
					Name:        "alert",
					Description: `(Required) The list of alert templates. See [Alert](#alert) below for details. ### Alert ` + "`" + `alert` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name that policies can use to refer to this template when generating alerts.`,
				},
				resource.Attribute{
					Name:        "event_name",
					Description: `(Required) The name of the resulting alert.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The summary of the resulting alert.`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `(Required) A high-level observation of the resulting alert. ## Import A Lacework Alert Profile can be imported using it's ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_profile.example CUSTOM_PROFILE_TERRAFORM_TEST ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_rule",
			Category:         "Alert Rules",
			ShortDescription: `Create and manage Lacework Alert Rules`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"rules",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The alert rule name.`,
				},
				resource.Attribute{
					Name:        "alert_channels",
					Description: `(Required) The list of alert channels for the rule to use.`,
				},
				resource.Attribute{
					Name:        "severities",
					Description: `(Required) The list of the severities that the rule will apply. Valid severities include: ` + "`" + `Critical` + "`" + `, ` + "`" + `High` + "`" + `, ` + "`" + `Medium` + "`" + `, ` + "`" + `Low` + "`" + ` and ` + "`" + `Info` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the alert rule.`,
				},
				resource.Attribute{
					Name:        "event_categories",
					Description: `(Optional) The list of event categories the rule will apply to. Valid categories include: ` + "`" + `Compliance` + "`" + `, ` + "`" + `App` + "`" + `, ` + "`" + `Cloud` + "`" + `,` + "`" + `File` + "`" + `, ` + "`" + `Machine` + "`" + `, ` + "`" + `User` + "`" + ` and ` + "`" + `Platform` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_groups",
					Description: `(Optional) The list of resource groups the rule will apply to.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Alert Rule can be imported using a ` + "`" + `GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_rule.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ``,
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
			Type:             "lacework_integration_aws_eks_audit_log",
			Category:         "Cloud Account Integrations",
			ShortDescription: `Create and manage AWS EKS Audit Log integrations`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"account",
				"integrations",
				"integration",
				"aws",
				"eks",
				"audit",
				"log",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The AWS CloudTrail integration name.`,
				},
				resource.Attribute{
					Name:        "sns_arn",
					Description: `(Required) The SNS topic ARN to share with Lacework.`,
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
					Description: `(Optional) The number of attempts to create the cloud account integration. Defaults to ` + "`" + `5` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_aws_govcloud_cfg",
			Category:         "Cloud Account Integrations",
			ShortDescription: `Create and manage AWS GovCloud Config integrations`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"account",
				"integrations",
				"integration",
				"aws",
				"govcloud",
				"cfg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The AWS GovCloud Config integration name.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) The AWS account ID.`,
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
					Name:        "access_key_id",
					Description: `(Required) The AWS access key ID.`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Required) The AWS secret key for the specified AWS access key. ## Import A Lacework AWS Config integration for AWS GovCloud can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_aws_govcloud_cfg.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
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
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "limit_by_label",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "limit_by_repos",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "limit_num_imgs",
					Description: `(Optional) The maximum number of newest container images to assess per repository. Must be one of ` + "`" + `5` + "`" + `, ` + "`" + `10` + "`" + `, or ` + "`" + `15` + "`" + `. Defaults to ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_tags",
					Description: `(Optional) A list of image tags to limit the assessment of images with matching tags. If you specify ` + "`" + `limit_by_tags` + "`" + ` and ` + "`" + `limit_by_labels` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_labels",
					Description: `(Optional) A key based map of labels to limit the assessment of images with matching ` + "`" + `key:value` + "`" + ` labels. If you specify ` + "`" + `limit_by_tags` + "`" + ` and ` + "`" + `limit_by_labels` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_repositories",
					Description: `(Optional) A list of repositories to assess.`,
				},
				resource.Attribute{
					Name:        "non_os_package_support",
					Description: `(Optional) Enable [program language scanning](https://docs.lacework.com/container-image-support#language-libraries-support). Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Docker Hub container registry integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_docker_hub.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
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
					Name:        "notifications",
					Description: `(Optional) Subscribe to registry notifications. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_tag",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "limit_by_label",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_tags",
					Description: `(Optional) A list of image tags to limit the assessment of images with matching tags. If you specify ` + "`" + `limit_by_tags` + "`" + ` and ` + "`" + `limit_by_labels` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_labels",
					Description: `(Optional) A key based map of labels to limit the assessment of images with matching ` + "`" + `key:value` + "`" + ` labels. If you specify ` + "`" + `limit_by_tags` + "`" + ` and ` + "`" + `limit_by_labels` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "non_os_package_support",
					Description: `(Optional) Enable [program language scanning](https://docs.lacework.com/container-image-support#language-libraries-support). Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Docker V2 container registry integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_docker_v2.jfrog EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
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
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "limit_by_label",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "limit_by_repos",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "limit_num_imgs",
					Description: `(Optional) The maximum number of newest container images to assess per repository. Must be one of ` + "`" + `5` + "`" + `, ` + "`" + `10` + "`" + `, or ` + "`" + `15` + "`" + `. Defaults to ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_tags",
					Description: `(Optional) A list of image tags to limit the assessment of images with matching tags. If you specify ` + "`" + `limit_by_tags` + "`" + ` and ` + "`" + `limit_by_labels` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_labels",
					Description: `(Optional) A key based map of labels to limit the assessment of images with matching ` + "`" + `key:value` + "`" + ` labels. If you specify ` + "`" + `limit_by_tags` + "`" + ` and ` + "`" + `limit_by_labels` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_repositories",
					Description: `(Optional) A list of repositories to assess.`,
				},
				resource.Attribute{
					Name:        "non_os_package_support",
					Description: `(Optional) Enable [program language scanning](https://docs.lacework.com/container-image-support#language-libraries-support). Defaults to ` + "`" + `true` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the combination of the following arguments.`,
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
			Type:             "lacework_integration_gar",
			Category:         "Container Registry Integrations",
			ShortDescription: `Create and manage Google Artifact Registry (GAR) integrations`,
			Description:      ``,
			Keywords: []string{
				"container",
				"registry",
				"integrations",
				"integration",
				"gar",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The integration name.`,
				},
				resource.Attribute{
					Name:        "registry_domain",
					Description: `(Required) The GAR domain, which specifies the location where you store the images. For a list of supported domains, see [Supported Registry Domains](#supported-registry-domains) below.`,
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
					Name:        "limit_num_imgs",
					Description: `(Optional) The maximum number of newest container images to assess per repository. Must be one of ` + "`" + `5` + "`" + `, ` + "`" + `10` + "`" + `, or ` + "`" + `15` + "`" + `. Defaults to ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_tags",
					Description: `(Optional) A list of image tags to limit the assessment of images with matching tags. If you specify ` + "`" + `limit_by_tags` + "`" + ` and ` + "`" + `limit_by_label` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_label",
					Description: `(Optional) A list of key/value labels to limit the assessment of images. If you specify ` + "`" + `limit_by_tags` + "`" + ` and ` + "`" + `limit_by_label` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_repositories",
					Description: `(Optional) A list of repositories to assess.`,
				},
				resource.Attribute{
					Name:        "non_os_package_support",
					Description: `(Optional) Enable [program language scanning](https://docs.lacework.com/container-image-support#language-libraries-support). Defaults to ` + "`" + `true` + "`" + `. The ` + "`" + `limit_by_label` + "`" + ` block can be defined multiple times to define multiple label limits, it supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the label.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the label. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
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
			ShortDescription: `Create and manage Google Container Registry (GCR) integrations`,
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
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "limit_by_label",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "limit_by_repos",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "limit_num_imgs",
					Description: `(Optional) The maximum number of newest container images to assess per repository. Must be one of ` + "`" + `5` + "`" + `, ` + "`" + `10` + "`" + `, or ` + "`" + `15` + "`" + `. Defaults to ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_tags",
					Description: `(Optional) A list of image tags to limit the assessment of images with matching tags. If you specify ` + "`" + `limit_by_tags` + "`" + ` and ` + "`" + `limit_by_labels` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_labels",
					Description: `(Optional) A key based map of labels to limit the assessment of images with matching ` + "`" + `key:value` + "`" + ` labels. If you specify ` + "`" + `limit_by_tags` + "`" + ` and ` + "`" + `limit_by_labels` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_repositories",
					Description: `(Optional) A list of repositories to assess.`,
				},
				resource.Attribute{
					Name:        "non_os_package_support",
					Description: `(Optional) Enable [program language scanning](https://docs.lacework.com/container-image-support#language-libraries-support). Defaults to ` + "`" + `true` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
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
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_ghcr",
			Category:         "Container Registry Integrations",
			ShortDescription: `Create and manage Github Container Registry (GHCR) integrations`,
			Description:      ``,
			Keywords: []string{
				"container",
				"registry",
				"integrations",
				"integration",
				"ghcr",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The integration name.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The Github username.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The Github personal access token with permission ` + "`" + `read:packages` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `(Optional) Enable or disable SSL communication. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "registry_notifications",
					Description: `(Optional) Subscribe to registry notifications. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_num_imgs",
					Description: `(Optional) The maximum number of newest container images to assess per repository. Must be one of ` + "`" + `5` + "`" + `, ` + "`" + `10` + "`" + `, or ` + "`" + `15` + "`" + `. Defaults to ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_tags",
					Description: `(Optional) A list of image tags to limit the assessment of images with matching tags. If you specify ` + "`" + `limit_by_tags` + "`" + ` and ` + "`" + `limit_by_label` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_label",
					Description: `(Optional) A list of key/value labels to limit the assessment of images. If you specify ` + "`" + `limit_by_tags` + "`" + ` and ` + "`" + `limit_by_label` + "`" + ` limits, they function as an ` + "`" + `AND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit_by_repositories",
					Description: `(Optional) A list of repositories to assess.`,
				},
				resource.Attribute{
					Name:        "non_os_package_support",
					Description: `(Optional) Enable [program language scanning](https://docs.lacework.com/container-image-support#language-libraries-support). Defaults to ` + "`" + `true` + "`" + `. The ` + "`" + `limit_by_label` + "`" + ` block can be defined multiple times to define multiple label limits, it supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the label.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the label. ## Import A Lacework Github container registry integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_ghcr.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_policy",
			Category:         "LQL",
			ShortDescription: `Create and manage Lacework Policies`,
			Description:      ``,
			Keywords: []string{
				"lql",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "title",
					Description: `(Required) The policy title.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description of the policy.`,
				},
				resource.Attribute{
					Name:        "query_id",
					Description: `(Required) The query id.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Required) The list of the severities. Valid severities include: ` + "`" + `Critical` + "`" + `, ` + "`" + `High` + "`" + `, ` + "`" + `Medium` + "`" + `, ` + "`" + `Low` + "`" + ` and ` + "`" + `Info` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The policy type must be either ` + "`" + `Violation` + "`" + ` or ` + "`" + `Summary` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "evaluation",
					Description: `(Optional) The evaluation frequency at which the policy will be evaluated. Valid values are ` + "`" + `Hourly` + "`" + ` or ` + "`" + `Daily` + "`" + `. Defaults to ` + "`" + `Hourly` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "remediation",
					Description: `(Optional) The remediation message to display.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) Set the maximum number of records returned by the policy. Maximum value is ` + "`" + `5000` + "`" + `. Defaults to ` + "`" + `1000` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the policy is enabled or disabled. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_id_suffix",
					Description: `(Optional) The string appended to the end of the policy id.`,
				},
				resource.Attribute{
					Name:        "alerting",
					Description: `(Optional) Alerting. See [Alerting](#alerting) below for details. ### Alerting ` + "`" + `alerting` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(Required) The alerting profile.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the alerting profile is enabled or disabled. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework policy can be imported using a ` + "`" + `POLICY_ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_policy.example YourLQLPolicyID ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_query",
			Category:         "LQL",
			ShortDescription: `Create and manage Lacework Queries`,
			Description:      ``,
			Keywords: []string{
				"lql",
				"query",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "query_id",
					Description: `(Required) The query id.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) The query string. ## Import A Lacework query can be imported using a ` + "`" + `QUERY_ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_query.example YourLQLQueryID ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_report_rule",
			Category:         "Report Rules",
			ShortDescription: `Create and manage Lacework Report Rules`,
			Description:      ``,
			Keywords: []string{
				"report",
				"rules",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The report rule name.`,
				},
				resource.Attribute{
					Name:        "email_alert_channels",
					Description: `(Required) The list of email alert channels for the rule to use.`,
				},
				resource.Attribute{
					Name:        "severities",
					Description: `(Required) The list of the severities that the rule will apply. Valid severities include: ` + "`" + `Critical` + "`" + `, ` + "`" + `High` + "`" + `, ` + "`" + `Medium` + "`" + `, ` + "`" + `Low` + "`" + ` and ` + "`" + `Info` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the report rule.`,
				},
				resource.Attribute{
					Name:        "resource_groups",
					Description: `(Optional) The list of resource groups the rule will apply to.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aws_compliance_reports",
					Description: `(Optional) Compliance reports for Aws. See [Aws Compliance Reports](#aws-compliance-reports) below for details.`,
				},
				resource.Attribute{
					Name:        "azure_compliance_reports",
					Description: `(Optional) Compliance reports for Azure. See [Azure Compliance Reports](#azure-compliance-reports) below for details.`,
				},
				resource.Attribute{
					Name:        "gcp_compliance_reports",
					Description: `(Optional) Compliance reports for Gcp. See [Gcp Compliance Reports](#gcp-compliance-reports) below for details.`,
				},
				resource.Attribute{
					Name:        "daily_compliance_reports",
					Description: `(Optional) Daily event summary reports. See [Daily Compliance Reports](#faily-compliance-reports) below for details.`,
				},
				resource.Attribute{
					Name:        "weekly_snapshot",
					Description: `(Optional) A weekly compliance trend report for all monitored resources. Defaults to ` + "`" + `false` + "`" + `. ### Aws Compliance Reports ` + "`" + `aws_compliance_reports` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "cis_s3",
					Description: `(Optional) AWS CIS Benchmark and S3 Report. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hipaa",
					Description: `(Optional) AWS HIPAA Report. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iso_2700",
					Description: `(Optional) AWS ISO 27001:2013 Report. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "nist_800_53_rev4",
					Description: `(Optional) AWS NIST 800-53 Report. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "nist_800_171_rev2",
					Description: `(Optional) AWS NIST 800-171 Report. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pci",
					Description: `(Optional) AWS PCI DSS Report. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "soc",
					Description: `(Optional) AWS SOC 2 Report. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "soc_rev2",
					Description: `(Optional) AWS SOC 2 Report Rev2. Defaults to ` + "`" + `false` + "`" + `. ### Azure Compliance Reports ` + "`" + `azure_compliance_reports` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "cis",
					Description: `(Optional) Azure CIS Benchmark. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cis_131",
					Description: `(Optional) Azure CIS 1.3.1 Benchmark. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pci",
					Description: `(Optional) Azure PCI Benchmark. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "soc",
					Description: `(Optional) Azure SOC 2 Report. Defaults to ` + "`" + `false` + "`" + `. ### Gcp Compliance Reports ` + "`" + `gcp_compliance_reports` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "cis",
					Description: `(Optional) GCP CIS Benchmark. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hipaa",
					Description: `(Optional) GCP HIPAA Report. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hipaa_rev2",
					Description: `(Optional) GCP HIPAA Report Rev2. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iso_27001",
					Description: `(Optional) GCP ISO 27001 Report. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cis_12",
					Description: `(Optional) GCP CIS 1.2 Benchmark. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "k8s",
					Description: `(Optional) GCP K8S Benchmark. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pci",
					Description: `(Optional) GCP PCI Benchmark. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pci_rev2",
					Description: `(Optional) GCP PCI Benchmark Rev2. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "soc",
					Description: `(Optional) GCP SOC 2 Report. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "soc_rev2",
					Description: `(Optional) GCP SOC 2 Report Rev2. Defaults to ` + "`" + `false` + "`" + `. ### Daily Compliance Reports ` + "`" + `daily_compliance_reports` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "host_security",
					Description: `(Optional) Host Security. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `(Optional) Platform Events. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "openshift_compliance",
					Description: `Openshift Compliance (Optional) Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "openshift_compliance_events",
					Description: `Openshift Compliance Events (Optional) Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aws_cloudtrail",
					Description: `(Optional) AWS CloudTrail. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aws_compliance",
					Description: `(Optional) AWS Compliance. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "azure_activity_log",
					Description: `(Optional) Azure Activity Log. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gcp_audit_trail",
					Description: `(Optional) GCP Audit Trail. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gcp_compliance",
					Description: `(Optional) GCP Compliance. Defaults to ` + "`" + `false` + "`" + `. ## Import A Lacework Report Rule can be imported using a ` + "`" + `GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_report_rule.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_resource_group_account",
			Category:         "Resource Groups",
			ShortDescription: `Create and manage Lacework Account Resource Groups`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"groups",
				"group",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource group name.`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `(Required) The list of Lacework accounts to include in the resource group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the resource group.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Account Resource Group can be imported using a ` + "`" + `RESOURCE_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_resource_group_account.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_resource_group_aws",
			Category:         "Resource Groups",
			ShortDescription: `Create and manage AWS Resource Groups`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"groups",
				"group",
				"aws",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource group name.`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `(Required) The list of AWS account ids to include in the resource group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the resource group.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework AWS Resource Group can be imported using a ` + "`" + `RESOURCE_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_resource_group_aws.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_resource_group_azure",
			Category:         "Resource Groups",
			ShortDescription: `Create and manage Azure Resource Groups`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"groups",
				"group",
				"azure",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource group name.`,
				},
				resource.Attribute{
					Name:        "tenant",
					Description: `(Required) The Azure tenant id.`,
				},
				resource.Attribute{
					Name:        "subscriptions",
					Description: `(Required) The list of Azure subscription ids to include in the resource group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the resource group.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Azure Resource Group can be imported using a ` + "`" + `RESOURCE_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_resource_group_azure.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_resource_group_container",
			Category:         "Resource Groups",
			ShortDescription: `Create and manage Container Resource Groups`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"groups",
				"group",
				"container",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource group name.`,
				},
				resource.Attribute{
					Name:        "container_labels",
					Description: `(Required) The key value pairs of container labels to include in the resource group.`,
				},
				resource.Attribute{
					Name:        "container_tags",
					Description: `(Required) The list of container tags to include in the resource group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the resource group.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Container Resource Group can be imported using a ` + "`" + `RESOURCE_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_resource_group_container.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_resource_group_gcp",
			Category:         "Resource Groups",
			ShortDescription: `Create and manage GCP Resource Groups`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"groups",
				"group",
				"gcp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource group name.`,
				},
				resource.Attribute{
					Name:        "projects",
					Description: `(Required) The list of GCP project ids to include in the resource group.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) The GCP organization id. If your project is not part of an organization or if you are looking to group projects across multiple organizations, enter an asterisk ` + "`" + `"`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the resource group.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework GCP Resource Group can be imported using a ` + "`" + `RESOURCE_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_resource_group_gcp.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_resource_group_machine",
			Category:         "Resource Groups",
			ShortDescription: `Create and manage Machine Resource Groups`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"groups",
				"group",
				"machine",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource group name.`,
				},
				resource.Attribute{
					Name:        "machine_tags",
					Description: `(Required) The key value pairs of machine tags to include in the resource group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the resource group.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework Machine Resource Group can be imported using a ` + "`" + `RESOURCE_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_resource_group_machine.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_team_member",
			Category:         "Team Members",
			ShortDescription: `Create and manage Team Members`,
			Description:      ``,
			Keywords: []string{
				"team",
				"members",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "first_name",
					Description: `(Required) The team member first name.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `(Required) The team member last name.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The team member email address, which will also be used as the username.`,
				},
				resource.Attribute{
					Name:        "administrator",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to make the team member an administrator, otherwise the member will be a regular user. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Optional) Use this block to manage organization-level team members. See [Organization](#organization) below for details.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the team member. Defaults to ` + "`" + `true` + "`" + `. ### Organization ` + "`" + `organization` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) Whether the team member is an organization-level user. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "administrator",
					Description: `(Optional) Whether the team member is an organization-level administrator. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_accounts",
					Description: `(Optional) List of accounts the team member is a user.`,
				},
				resource.Attribute{
					Name:        "admin_accounts",
					Description: `(Optional) List of accounts the team member is an administrator. ## Import There are two ways to import a team member. ### Import Standalone Team Member A Lacework standalone team member can be imported using a ` + "`" + `USER_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_team_member.harry HOGWARTS_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ### Import Organizational Team Member A Lacework organization-level team member can be imported using the ` + "`" + `email` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_team_member.albus albus@hogwarts.io ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_vulnerability_exception_container",
			Category:         "Vulnerability Exceptions",
			ShortDescription: `Create and manage Lacework vulnerability exceptions for containers`,
			Description:      ``,
			Keywords: []string{
				"vulnerability",
				"exceptions",
				"exception",
				"container",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The vulnerability exception name.`,
				},
				resource.Attribute{
					Name:        "vulnerability_criteria",
					Description: `(Required) The vulnerability criteria. See [Vulnerability Criteria](#vulnerability-criteria) below for details.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the vulnerability exception.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `(Optional) The list of the severities that the rule will apply. Valid severities include: ` + "`" + `Accepted Risk` + "`" + `, ` + "`" + `False Positive` + "`" + `, ` + "`" + `Compensating Controls` + "`" + `, Fix Pending` + "`" + ` and ` + "`" + `Other` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_scope",
					Description: `(Optional) The resource scope. See [Resource Scope](#resource-scope) below for details. ### Vulnerability Criteria ` + "`" + `vulnerability criteria` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "severities",
					Description: `(Optional) The list of the severities. Valid severities include: ` + "`" + `Critical` + "`" + `, ` + "`" + `High` + "`" + `, ` + "`" + `Medium` + "`" + `, ` + "`" + `Low` + "`" + ` and ` + "`" + `Info` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cves",
					Description: `(Optional) The list of the cves.`,
				},
				resource.Attribute{
					Name:        "package",
					Description: `(Optional) The list of the packages.`,
				},
				resource.Attribute{
					Name:        "fixable",
					Description: `(Optional) Whether to filter on fixable. Defaults to ` + "`" + `false` + "`" + `. ### Resource Scope ` + "`" + `resource_scope` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "image_ids",
					Description: `(Optional) The list of image ids.`,
				},
				resource.Attribute{
					Name:        "image_tags",
					Description: `(Optional) The list of image tags.`,
				},
				resource.Attribute{
					Name:        "registries",
					Description: `(Optional) The list of registries.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `(Optional) The list of repositories.`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `(Optional) The list of namespaces. ## Import A Lacework Vulnerability Exception Container can be imported using a ` + "`" + `GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_vulnerability_exception_container.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_vulnerability_exception_host",
			Category:         "Vulnerability Exceptions",
			ShortDescription: `Create and manage Lacework vulnerability exceptions for hosts`,
			Description:      ``,
			Keywords: []string{
				"vulnerability",
				"exceptions",
				"exception",
				"host",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The vulnerability exception name.`,
				},
				resource.Attribute{
					Name:        "vulnerability_criteria",
					Description: `(Required) The vulnerability criteria. See [Vulnerability Criteria](#vulnerability-criteria) below for details.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the vulnerability exception.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `(Optional) The list of the severities that the rule will apply. Valid severities include: ` + "`" + `Accepted Risk` + "`" + `, ` + "`" + `False Positive` + "`" + `, ` + "`" + `Compensating Controls` + "`" + `, Fix Pending` + "`" + ` and ` + "`" + `Other` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_scope",
					Description: `(Optional) The resource scope. See [Resource Scope](#resource-scope) below for details. ### Vulnerability Criteria ` + "`" + `vulnerability criteria` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "severities",
					Description: `(Optional) The list of the severities. Valid severities include: ` + "`" + `Critical` + "`" + `, ` + "`" + `High` + "`" + `, ` + "`" + `Medium` + "`" + `, ` + "`" + `Low` + "`" + ` and ` + "`" + `Info` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cves",
					Description: `(Optional) The list of the cves.`,
				},
				resource.Attribute{
					Name:        "package",
					Description: `(Optional) The list of the cves. See [Resource Scope](#resource-scope) below for details.`,
				},
				resource.Attribute{
					Name:        "fixable",
					Description: `(Optional) Whether to filter on fixable. Defaults to ` + "`" + `false` + "`" + `. ### Resource Scope ` + "`" + `resource_scope` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "hostnames",
					Description: `(Optional) The list of hostnames.`,
				},
				resource.Attribute{
					Name:        "external_ips",
					Description: `(Optional) The list of external ips.`,
				},
				resource.Attribute{
					Name:        "cluster_names",
					Description: `(Optional) The list of cluster names.`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `(Optional) The list of namespaces. ## Import A Lacework vulnerability exception for hosts can be imported using a ` + "`" + `GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_vulnerability_exception_host.example EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"lacework_agent_access_token":                0,
		"lacework_alert_channel_aws_cloudwatch":      1,
		"lacework_alert_channel_aws_s3":              2,
		"lacework_alert_channel_cisco_webex":         3,
		"lacework_alert_channel_datadog":             4,
		"lacework_alert_channel_email":               5,
		"lacework_alert_channel_gcp_pub_sub":         6,
		"lacework_alert_channel_jira_cloud":          7,
		"lacework_alert_channel_jira_server":         8,
		"lacework_alert_channel_microsoft_teams":     9,
		"lacework_alert_channel_newrelic":            10,
		"lacework_alert_channel_pagerduty":           11,
		"lacework_alert_channel_qradar":              12,
		"lacework_alert_channel_service_now":         13,
		"lacework_alert_channel_slack":               14,
		"lacework_alert_channel_splunk":              15,
		"lacework_alert_channel_victorops":           16,
		"lacework_alert_channel_webhook":             17,
		"lacework_alert_profile":                     18,
		"lacework_alert_rule":                        19,
		"lacework_integration_aws_cfg":               20,
		"lacework_integration_aws_ct":                21,
		"lacework_integration_aws_eks_audit_log":     22,
		"lacework_integration_aws_govcloud_cfg":      23,
		"lacework_integration_azure_al":              24,
		"lacework_integration_azure_cfg":             25,
		"lacework_integration_docker_hub":            26,
		"lacework_integration_docker_v2":             27,
		"lacework_integration_ecr":                   28,
		"lacework_integration_gar":                   29,
		"lacework_integration_gcp_at":                30,
		"lacework_integration_gcp_cfg":               31,
		"lacework_integration_gcr":                   32,
		"lacework_integration_ghcr":                  33,
		"lacework_policy":                            34,
		"lacework_query":                             35,
		"lacework_report_rule":                       36,
		"lacework_resource_group_account":            37,
		"lacework_resource_group_aws":                38,
		"lacework_resource_group_azure":              39,
		"lacework_resource_group_container":          40,
		"lacework_resource_group_gcp":                41,
		"lacework_resource_group_machine":            42,
		"lacework_team_member":                       43,
		"lacework_vulnerability_exception_container": 44,
		"lacework_vulnerability_exception_host":      45,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
