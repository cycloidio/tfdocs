package lacework

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_aws_cloudwatch",
			Category:         "Resources",
			ShortDescription: `Create an manage AWS CloudWatch Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
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
					Name:        "min_alert_severity",
					Description: `(Optional) The minimum severity level to alert. Defaults to ` + "`" + `3` + "`" + `. The available levels are:`,
				},
				resource.Attribute{
					Name:        "group_issues_by",
					Description: `(Optional) Defines how Lacework compliance events get grouped. Must be one of ` + "`" + `EVENTS` + "`" + ` or ` + "`" + `RESOURCES` + "`" + `. Defaults to ` + "`" + `EVENTS` + "`" + `. The available options are:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ## Import A Lacework AWS CloudWatch Alert Channel integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_alert_channel_aws_cloudwatch.all_resources EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_alert_channel_pagerduty",
			Category:         "Resources",
			ShortDescription: `Create an manage PagerDuty Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
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
					Name:        "min_alert_severity",
					Description: `(Optional) The minimum severity level to alert. Defaults to ` + "`" + `3` + "`" + `. The available levels are:`,
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
			Type:             "lacework_alert_channel_slack",
			Category:         "Resources",
			ShortDescription: `Create an manage Slack Alert Channel integrations`,
			Description:      ``,
			Keywords: []string{
				"alert",
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
					Name:        "min_alert_severity",
					Description: `(Optional) The minimum severity level to alert. Defaults to ` + "`" + `3` + "`" + `. The available levels are:`,
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
			Type:             "lacework_integration_aws_cfg",
			Category:         "Resources",
			ShortDescription: `Create an manage AWS Config integrations`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"aws",
				"cfg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Lacework AWS Config integration name.`,
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
					Description: `(Required) The external ID for the IAM role. ## Import A Lacework AWS Config integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_aws_cfg.account_abc EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_aws_ct",
			Category:         "Resources",
			ShortDescription: `Create an manage AWS CloudTrail integrations`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"aws",
				"ct",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Lacework AWS CloudTrail integration name.`,
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
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_azure_al",
			Category:         "Resources",
			ShortDescription: `Create an manage Azure Cloud Activity Log integrations`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"azure",
				"al",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Lacework Azure Config integration name.`,
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
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
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
			Category:         "Resources",
			ShortDescription: `Create an manage Azure Cloud Config integrations`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"azure",
				"cfg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Lacework Azure Config integration name.`,
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
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
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
			Type:             "lacework_integration_gcp_at",
			Category:         "Resources",
			ShortDescription: `Create an manage Google Cloud Audit Trail integrations`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"gcp",
				"at",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Lacework GCP Audit Trail integration name.`,
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
					Description: `(Required) The service account private key. ## Import A Lacework GCP Audit Trail integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_gcp_at.account_abc EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_gcp_cfg",
			Category:         "Resources",
			ShortDescription: `Create an manage Google Cloud Config integrations`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"gcp",
				"cfg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Lacework GCP Config integration name.`,
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
					Description: `(Required) The service account private key. ## Import A Lacework GCP Config integration can be imported using a ` + "`" + `INT_GUID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lacework_integration_gcp_cfg.account_abc EXAMPLE_1234BAE1E42182964D23973F44CFEA3C4AB63B99E9A1EC5 ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"lacework_alert_channel_aws_cloudwatch": 0,
		"lacework_alert_channel_pagerduty":      1,
		"lacework_alert_channel_slack":          2,
		"lacework_integration_aws_cfg":          3,
		"lacework_integration_aws_ct":           4,
		"lacework_integration_azure_al":         5,
		"lacework_integration_azure_cfg":        6,
		"lacework_integration_gcp_at":           7,
		"lacework_integration_gcp_cfg":          8,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
