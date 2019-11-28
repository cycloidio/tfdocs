package dome9

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dome9_cloudaccount_aws",
			Category:         "Resources",
			ShortDescription: `Onboard AWS cloud account`,
			Description:      ``,
			Keywords: []string{
				"cloudaccount",
				"aws",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of AWS account in Dome9`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The information needed for Dome9 System in order to connect to the AWS cloud account ### Credentials ` + "`" + `credentials` + "`" + ` has the following arguments:`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `(Required) AWS Role ARN (to be assumed by Dome9)`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Required) The AWS role External ID (Dome9 will have to use this secret in order to assume the role)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The cloud account onboarding method. Set to "RoleBased". ### Network security configuration ` + "`" + `net_sec` + "`" + ` has the these arguments:`,
				},
				resource.Attribute{
					Name:        "Regions",
					Description: `(Required) list of the supported regions, and their configuration:`,
				},
				resource.Attribute{
					Name:        "new_group_behavior",
					Description: `(Required) The network security configuration. Select "ReadOnly", "FullManage", or "Reset".`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) AWS region, in AWS format (e.g., "us-east-1") ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider ("AWS").`,
				},
				resource.Attribute{
					Name:        "external_account_number",
					Description: `The AWS account number.`,
				},
				resource.Attribute{
					Name:        "is_fetching_suspended",
					Description: `Fetching suspending status.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Date the account was onboarded to Dome9.`,
				},
				resource.Attribute{
					Name:        "full_protection",
					Description: `The protection mode for existing security groups in the account.`,
				},
				resource.Attribute{
					Name:        "allow_read_only",
					Description: `The AWS cloud account operation mode. true for "Full-Manage", false for "Readonly".`,
				},
				resource.Attribute{
					Name:        "net_sec",
					Description: `The network security configuration for the AWS cloud account. If not given, sets to default value. ## Import AWS cloud account can be imported; use ` + "`" + `<AWS CLOUD ACCOUNT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_cloudaccount_aws.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider ("AWS").`,
				},
				resource.Attribute{
					Name:        "external_account_number",
					Description: `The AWS account number.`,
				},
				resource.Attribute{
					Name:        "is_fetching_suspended",
					Description: `Fetching suspending status.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Date the account was onboarded to Dome9.`,
				},
				resource.Attribute{
					Name:        "full_protection",
					Description: `The protection mode for existing security groups in the account.`,
				},
				resource.Attribute{
					Name:        "allow_read_only",
					Description: `The AWS cloud account operation mode. true for "Full-Manage", false for "Readonly".`,
				},
				resource.Attribute{
					Name:        "net_sec",
					Description: `The network security configuration for the AWS cloud account. If not given, sets to default value. ## Import AWS cloud account can be imported; use ` + "`" + `<AWS CLOUD ACCOUNT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_cloudaccount_aws.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_cloudaccount_azure",
			Category:         "Resources",
			ShortDescription: `Onboard Azure cloud account`,
			Description:      ``,
			Keywords: []string{
				"cloudaccount",
				"azure",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Azure account in Dome9`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required) The Azure subscription id for account`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) The Azure tenant id`,
				},
				resource.Attribute{
					Name:        "operation_mode",
					Description: `(Required) Dome9 operation mode for the Azure account ("Read-Only" or "Managed")`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) Azure account credentials`,
				},
				resource.Attribute{
					Name:        "organizational_unit_id",
					Description: `Organizational unit id, will apply on update state ### Credentials The ` + "`" + `credentials` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) Azure account id`,
				},
				resource.Attribute{
					Name:        "client_password",
					Description: `(Required) Password for account ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Azure cloud account`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider ("Azure")`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Date the account was onboarded to Dome9`,
				},
				resource.Attribute{
					Name:        "organizational_unit_path",
					Description: `Organizational unit path`,
				},
				resource.Attribute{
					Name:        "organizational_unit_name",
					Description: `Organizational unit name ## Import Azure cloud account can be imported; use ` + "`" + `<Azure CLOUD ACCOUNT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_cloudaccount_Azure.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Azure cloud account`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider ("Azure")`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Date the account was onboarded to Dome9`,
				},
				resource.Attribute{
					Name:        "organizational_unit_path",
					Description: `Organizational unit path`,
				},
				resource.Attribute{
					Name:        "organizational_unit_name",
					Description: `Organizational unit name ## Import Azure cloud account can be imported; use ` + "`" + `<Azure CLOUD ACCOUNT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_cloudaccount_Azure.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_cloudaccount_gcp",
			Category:         "Resources",
			ShortDescription: `Onboard GCP cloud account`,
			Description:      ``,
			Keywords: []string{
				"cloudaccount",
				"gcp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Google account name in Dome9`,
				},
				resource.Attribute{
					Name:        "gsuite_user",
					Description: `(Optional) The gsuite user`,
				},
				resource.Attribute{
					Name:        "service_account_credentials",
					Description: `(Required) The service account JSON block (from the GCP console)`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional) The domain name`,
				},
				resource.Attribute{
					Name:        "organizational_unit_id",
					Description: `(Optional) Organizational unit id, Will apply on update state ### Service Account Credentials The ` + "`" + `service_account_credentials` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) type. i.e "service_account"`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Project ID`,
				},
				resource.Attribute{
					Name:        "private_key_id",
					Description: `(Required) Private key ID`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) Private key`,
				},
				resource.Attribute{
					Name:        "client_email",
					Description: `(Required) GCP client email`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) Client id`,
				},
				resource.Attribute{
					Name:        "auth_uri",
					Description: `(Required) Auth URI. i.e "https://accounts.google.com/o/oauth2/auth"`,
				},
				resource.Attribute{
					Name:        "token_uri",
					Description: `(Required) Token URI. i.e "https://oauth2.googleapis.com/token"`,
				},
				resource.Attribute{
					Name:        "auth_provider_x509_cert_url",
					Description: `(Required) auth_provider_x509_cert_url. i.e "https://www.googleapis.com/oauth2/v1/certs"`,
				},
				resource.Attribute{
					Name:        "client_x509_cert_url",
					Description: `(Required) client_x509_cert_url ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the GCP cloud account`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `creation date for project in Google.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider (gcp).`,
				},
				resource.Attribute{
					Name:        "organizational_unit_path",
					Description: `Organizational unit path.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_name",
					Description: `Organizational unit name. ## Import GCP cloud account can be imported; use ` + "`" + `<GCP CLOUD ACCOUNT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_cloudaccount_gcp.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the GCP cloud account`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `creation date for project in Google.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider (gcp).`,
				},
				resource.Attribute{
					Name:        "organizational_unit_path",
					Description: `Organizational unit path.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_name",
					Description: `Organizational unit name. ## Import GCP cloud account can be imported; use ` + "`" + `<GCP CLOUD ACCOUNT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_cloudaccount_gcp.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_continuous_compliance_notification",
			Category:         "Resources",
			ShortDescription: `Creates continuous compliance notification in Dome9`,
			Description:      ``,
			Keywords: []string{
				"continuous",
				"compliance",
				"notification",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The cloud account id in Dome9.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the notification. at least one of ` + "`" + `alerts_console` + "`" + `, ` + "`" + `scheduled_report` + "`" + `, or ` + "`" + `change_detection` + "`" + ` must be included`,
				},
				resource.Attribute{
					Name:        "alerts_console",
					Description: `(Optional) send findings (also) to the Dome9 web app alerts console (Boolean); default is False.`,
				},
				resource.Attribute{
					Name:        "scheduled_report",
					Description: `Scheduled email report notification block:`,
				},
				resource.Attribute{
					Name:        "email_sending_state",
					Description: `send schedule report of findings by email; can be "Enabled" or "Disabled". if ` + "`" + `email_sending_state` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "schedule_data",
					Description: `Schedule details:`,
				},
				resource.Attribute{
					Name:        "cron_expression",
					Description: `the schedule to issue the email report (in cron expression format)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `type of report; can be "Detailed", "Summary", "FullCsv" or "FullCsvZip"`,
				},
				resource.Attribute{
					Name:        "recipients",
					Description: `comma-separated list of email recipients`,
				},
				resource.Attribute{
					Name:        "change_detection",
					Description: `Send changes in findings options:`,
				},
				resource.Attribute{
					Name:        "email_sending_stat",
					Description: `send email report of changes in findings; can be "Enabled" or "Disabled". if ` + "`" + `email_sending_stat` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "email_data",
					Description: `Email notification details:`,
				},
				resource.Attribute{
					Name:        "recipients",
					Description: `comma-separated list of email recipients`,
				},
				resource.Attribute{
					Name:        "email_per_finding_sending_state",
					Description: `send separate email notification for each finding; can be "Enabled" or "Disabled" if ` + "`" + `email_per_finding_sending_state` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "email_per_finding_data",
					Description: `Email per finding notification details:`,
				},
				resource.Attribute{
					Name:        "recipients",
					Description: `comma-separated list of email recipients`,
				},
				resource.Attribute{
					Name:        "notification_output_format",
					Description: `(Required) format of JSON block for finding; can be "JsonWithFullEntity", "JsonWithBasicEntity", or "PlainText".`,
				},
				resource.Attribute{
					Name:        "sns_sending_state",
					Description: `send by AWS SNS for each new finding; can be "Enabled" or "Disabled". if ` + "`" + `sns_sending_state` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "sns_data",
					Description: `SNS notification details:`,
				},
				resource.Attribute{
					Name:        "sns_topic_arn",
					Description: `SNS topic ARN`,
				},
				resource.Attribute{
					Name:        "sns_output_format",
					Description: `SNS output format; can be "JsonWithFullEntity", "JsonWithBasicEntity", or "PlainText".`,
				},
				resource.Attribute{
					Name:        "external_ticket_creating_state",
					Description: `send each finding to an external ticketing system; can be "Enabled" or "Disabled". if ` + "`" + `external_ticket_creating_state` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "ticketing_system_data",
					Description: `Ticketing system details:`,
				},
				resource.Attribute{
					Name:        "system_type",
					Description: `system type; can be "ServiceOne", "Jira", or "PagerDuty"`,
				},
				resource.Attribute{
					Name:        "should_close_tickets",
					Description: `ticketing system should close tickets when resolved (bool)`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `ServiceNow domain name (ServiceNow only)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `user name (ServiceNow only)`,
				},
				resource.Attribute{
					Name:        "pass",
					Description: `password (ServiceNow only)`,
				},
				resource.Attribute{
					Name:        "project_key",
					Description: `project key (Jira) or API Key (PagerDuty)`,
				},
				resource.Attribute{
					Name:        "issue_type",
					Description: `issue type (Jira)`,
				},
				resource.Attribute{
					Name:        "webhook_integration_state",
					Description: `send findings to an HTTP endpoint (webhook); can be "Enabled" or "Disabled". if ` + "`" + `webhook_integration_state` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "webhook_data",
					Description: `Webhook data block supports:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `HTTP endpoint URL`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `HTTP method, "Post" by default.`,
				},
				resource.Attribute{
					Name:        "auth_method",
					Description: `authentication method; "NoAuth" by default`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `username in endpoint system`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `password in endpoint system`,
				},
				resource.Attribute{
					Name:        "format_type",
					Description: `format for JSON block for finding; can be "Basic" or "ServiceNow"`,
				},
				resource.Attribute{
					Name:        "aws_security_hub_integration_state",
					Description: `send findings to AWS Secure Hub; can be "Enabled" or "Disabled". if ` + "`" + `aws_security_hub_integration_state` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "aws_security_hub_integration",
					Description: `AWS security hub integration block supports:`,
				},
				resource.Attribute{
					Name:        "external_account_id",
					Description: `(Required) external account id`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) AWS region ` + "`" + `gcp_security_command_center_integration` + "`" + ` is a change_detection option`,
				},
				resource.Attribute{
					Name:        "gcp_security_command_center_integration",
					Description: `GCP security command center details`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `send findings to the GCP Security Command Center; can be "Enabled" or "Disabled" if ` + "`" + `state` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `GCP Project id`,
				},
				resource.Attribute{
					Name:        "source_id",
					Description: `GCP Source id ## Import The notification can be imported; use ` + "`" + `<NOTIFICATION ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_continuouscompliance_notification.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_continuous_compliance_policy",
			Category:         "Resources",
			ShortDescription: `Creates continuous compliance policies in Dome9`,
			Description:      ``,
			Keywords: []string{
				"continuous",
				"compliance",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `(Required) The cloud account id.`,
				},
				resource.Attribute{
					Name:        "external_account_id",
					Description: `(Required) The account number.`,
				},
				resource.Attribute{
					Name:        "bundle_id",
					Description: `(Required) The bundle id for the bundle that will be used in the policy.`,
				},
				resource.Attribute{
					Name:        "cloud_account_type",
					Description: `(Required) The cloud account provider ("AWS", "Azure", "Google").`,
				},
				resource.Attribute{
					Name:        "notification_ids",
					Description: `(Required) The notification policy id's for the policy [list]. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the compliance policy. ## Import The policy can be imported; use ` + "`" + `<POLICY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_continuouscompliance_policy.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Id of the compliance policy. ## Import The policy can be imported; use ` + "`" + `<POLICY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_continuouscompliance_policy.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_iplist",
			Category:         "Resources",
			ShortDescription: `Create IP lists in Dome9`,
			Description:      ``,
			Keywords: []string{
				"iplist",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the IP list in Dome9`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the list (what it represents); defaults to empty string`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `(Optional) the individual IP addresses for the list; defaults to empty list ### Items The ` + "`" + `items` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) IP address`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `IP list Id ## Import IP list can be imported; use ` + "`" + `<IP LIST ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_iplist.test 00000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `IP list Id ## Import IP list can be imported; use ` + "`" + `<IP LIST ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_iplist.test 00000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"dome9_cloudaccount_aws":                   0,
		"dome9_cloudaccount_azure":                 1,
		"dome9_cloudaccount_gcp":                   2,
		"dome9_continuous_compliance_notification": 3,
		"dome9_continuous_compliance_policy":       4,
		"dome9_iplist":                             5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
