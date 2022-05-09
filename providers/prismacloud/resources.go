package prismacloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_ExampleUsage",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_account_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) List of cloud account IDs.`,
				},
				resource.Attribute{
					Name:        "child_group_ids",
					Description: `(Optional) List of child account group IDs. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Account group ID.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_ts",
					Description: `(int) Last modified timestamp. ## Import Resources can be imported using the group ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_account_group.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `Account group ID.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_ts",
					Description: `(int) Last modified timestamp. ## Import Resources can be imported using the group ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_account_group.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_alert_rule",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Rule/Scan name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Enabled (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "scan_all",
					Description: `(bool) Scan all policies`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `List of specific policies to scan`,
				},
				resource.Attribute{
					Name:        "policy_labels",
					Description: `List of policy labels`,
				},
				resource.Attribute{
					Name:        "excluded_policies",
					Description: `List of policies to exclude from scan`,
				},
				resource.Attribute{
					Name:        "allow_auto_remediate",
					Description: `(bool) Allow auto-remediation`,
				},
				resource.Attribute{
					Name:        "delay_notification_ms",
					Description: `(int) Delay notifications by the specified miliseconds`,
				},
				resource.Attribute{
					Name:        "notify_on_open",
					Description: `(bool) Include open alerts in notification (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "notify_on_snoozed",
					Description: `(bool) Include snoozed alerts in notification`,
				},
				resource.Attribute{
					Name:        "notify_on_dismissed",
					Description: `(bool) Include dismissed alerts in notification`,
				},
				resource.Attribute{
					Name:        "notify_on_resolved",
					Description: `(bool) Include resolved alerts in notification`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) Model for the target filter, as defined [below](#target)`,
				},
				resource.Attribute{
					Name:        "notification_config",
					Description: `List of data for notifications to third-party tools, as defined [below](#notification-config) ### Target There should be one and only one target block:`,
				},
				resource.Attribute{
					Name:        "account_groups",
					Description: `(Required) List of account groups`,
				},
				resource.Attribute{
					Name:        "excluded_accounts",
					Description: `List of excluded accounts`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `List of regions`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `List of tag models, as defined [below](#tags) ### Tags`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Resource tag target`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `List of values for resource tag key ### Notification Config`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Frequency. Valid values are ` + "`" + `as_it_happens` + "`" + `, ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `, or ` + "`" + `monthly` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Scan enabled`,
				},
				resource.Attribute{
					Name:        "recipients",
					Description: `List of unique email addresses to notify (For email notifications), List of integration ids (For integrations without notification templates), or List of notification template ids (For integrations with notification templates)`,
				},
				resource.Attribute{
					Name:        "detailed_report",
					Description: `(bool) Provide CSV detailed report`,
				},
				resource.Attribute{
					Name:        "with_compression",
					Description: `(bool) Compress detailed report`,
				},
				resource.Attribute{
					Name:        "include_remediation",
					Description: `(bool) Include remediation in detailed report`,
				},
				resource.Attribute{
					Name:        "config_type",
					Description: `Config type. Valid values are ` + "`" + `email` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `splunk` + "`" + `, ` + "`" + `amazon_sqs` + "`" + `, ` + "`" + `microsoft_teams` + "`" + `, ` + "`" + `jira` + "`" + `, ` + "`" + `webhook` + "`" + `, ` + "`" + `aws_security_hub` + "`" + `, ` + "`" + `google_cscc` + "`" + `, ` + "`" + `service_now` + "`" + `, ` + "`" + `pager_duty` + "`" + `, ` + "`" + `aws_s3` + "`" + `, ` + "`" + `snowflake` + "`" + ` or ` + "`" + `demisto` + "`" + ``,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Template ID`,
				},
				resource.Attribute{
					Name:        "r_rule_schedule",
					Description: `R rule schedule ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "policy_scan_config_id",
					Description: `Policy scan config ID`,
				},
				resource.Attribute{
					Name:        "last_modified_on",
					Description: `(int) Last modified on`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Owner`,
				},
				resource.Attribute{
					Name:        "notification_channels",
					Description: `List of notification channels`,
				},
				resource.Attribute{
					Name:        "open_alerts_count",
					Description: `(int) Open alerts count`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(bool) Read only In each ` + "`" + `notification_config` + "`" + ` section, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `Alert rule notification config ID`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `(int) Last updated`,
				},
				resource.Attribute{
					Name:        "last_sent_ts",
					Description: `(int) Time of last notification in miliseconds`,
				},
				resource.Attribute{
					Name:        "timezone_id",
					Description: `Timezone ID`,
				},
				resource.Attribute{
					Name:        "day_of_month",
					Description: `(int) Day of month`,
				},
				resource.Attribute{
					Name:        "frequency_from_r_rule",
					Description: `Frequency from R rule`,
				},
				resource.Attribute{
					Name:        "hour_of_day",
					Description: `(int) Hour of day`,
				},
				resource.Attribute{
					Name:        "days_of_week",
					Description: `List of days of week, as defined [below](#days-of-week) ### Days Of Week`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `Day`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(int) Offset ## Import Resources can be imported using the policy scan config ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_alert_rule.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_scan_config_id",
					Description: `Policy scan config ID`,
				},
				resource.Attribute{
					Name:        "last_modified_on",
					Description: `(int) Last modified on`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Owner`,
				},
				resource.Attribute{
					Name:        "notification_channels",
					Description: `List of notification channels`,
				},
				resource.Attribute{
					Name:        "open_alerts_count",
					Description: `(int) Open alerts count`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(bool) Read only In each ` + "`" + `notification_config` + "`" + ` section, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `Alert rule notification config ID`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `(int) Last updated`,
				},
				resource.Attribute{
					Name:        "last_sent_ts",
					Description: `(int) Time of last notification in miliseconds`,
				},
				resource.Attribute{
					Name:        "timezone_id",
					Description: `Timezone ID`,
				},
				resource.Attribute{
					Name:        "day_of_month",
					Description: `(int) Day of month`,
				},
				resource.Attribute{
					Name:        "frequency_from_r_rule",
					Description: `Frequency from R rule`,
				},
				resource.Attribute{
					Name:        "hour_of_day",
					Description: `(int) Hour of day`,
				},
				resource.Attribute{
					Name:        "days_of_week",
					Description: `List of days of week, as defined [below](#days-of-week) ### Days Of Week`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `Day`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(int) Offset ## Import Resources can be imported using the policy scan config ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_alert_rule.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_cloud_account",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "disable_on_destroy",
					Description: `(Optional, bool) To disable cloud account instead of deleting when calling Terraform destroy (default: ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "aws",
					Description: `AWS account type spec, defined [below](#aws).`,
				},
				resource.Attribute{
					Name:        "azure",
					Description: `Azure account type spec, defined [below](#azure).`,
				},
				resource.Attribute{
					Name:        "gcp",
					Description: `GCP account type spec, defined [below](#gcp).`,
				},
				resource.Attribute{
					Name:        "alibaba_cloud",
					Description: `Alibaba account type spec, defined [below](#alibaba-cloud). ### AWS`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) AWS account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, bool) Whether or not the account is enabled (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Required) AWS account external ID.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `(Required) List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Required) Unique identifier for an AWS resource (ARN).`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `(Optional) Defaults to "account" if not specified`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `(Optional) Defaults to "MONITOR".Valid values : ` + "`" + `MONITOR` + "`" + ` or ` + "`" + `MONITOR_AND_PROTECT` + "`" + ` ### Azure`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) Azure account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, bool) Whether or not the account is enabled (defualt: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `(Required) List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) Application ID registered with Active Directory.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Application ID key.`,
				},
				resource.Attribute{
					Name:        "monitor_flow_logs",
					Description: `(Optional, bool) Automatically ingest flow logs.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Active Directory ID associated with Azure.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `(Required) Unique ID of the service principal object associated with the Prisma Cloud application that you create.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `(Optional) Defaults to "account" if not specified`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `(Optional) Defaults to "MONITOR". Valid values : ` + "`" + `MONITOR` + "`" + ` or ` + "`" + `MONITOR_AND_PROTECT` + "`" + ` ### GCP !> The Prisma Cloud API returns a series of asterisks for the private key of the ` + "`" + `credentials_json` + "`" + ` field instead of the configured value. Because of this, the provider cannot detect configuration drift on the private key within the ` + "`" + `credentials_json` + "`" + ` param.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) GCP project ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, bool) Whether or not the account is enabled (defualt: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `(Required) List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "compression_enabled",
					Description: `(Optional, bool) Enable flow log compression.`,
				},
				resource.Attribute{
					Name:        "data_flow_enabled_project",
					Description: `(Optional) GCP project for flow log compression.`,
				},
				resource.Attribute{
					Name:        "flow_log_storage_bucket",
					Description: `(Optional) GCP Flow logs storage bucket.`,
				},
				resource.Attribute{
					Name:        "credentials_json",
					Description: `(Required) Content of the JSON credentials file (read in using ` + "`" + `file()` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `(Optional) Defaults to "account" if not specified`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `(Optional) Defaults to "MONITOR". Valid values : ` + "`" + `MONITOR` + "`" + ` or ` + "`" + `MONITOR_AND_PROTECT` + "`" + ` ### Alibaba Cloud`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) Alibaba account ID.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `(Required) List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "ram_arn",
					Description: `(Required) Unique identifier for an Alibaba RAM role resource. ## Import Resources can be imported using the cloud type (` + "`" + `aws` + "`" + `, ` + "`" + `azure` + "`" + `, ` + "`" + `gcp` + "`" + `, or ` + "`" + `alibaba_cloud` + "`" + `) and the ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_cloud_account.aws_example aws:accountIdHere ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_compliance_standard",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Compliance standard name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "cs_id",
					Description: `Compliance standard ID`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `(int) Created on`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "last_modified_on",
					Description: `(int) Last modified on`,
				},
				resource.Attribute{
					Name:        "system_default",
					Description: `(bool) System default`,
				},
				resource.Attribute{
					Name:        "policies_assigned_count",
					Description: `(int) Number of assigned policies`,
				},
				resource.Attribute{
					Name:        "cloud_types",
					Description: `List of cloud types (determined based on policies assigned) ## Import Resources can be imported using the compliance standard ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_compliance_standard.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cs_id",
					Description: `Compliance standard ID`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `(int) Created on`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "last_modified_on",
					Description: `(int) Last modified on`,
				},
				resource.Attribute{
					Name:        "system_default",
					Description: `(bool) System default`,
				},
				resource.Attribute{
					Name:        "policies_assigned_count",
					Description: `(int) Number of assigned policies`,
				},
				resource.Attribute{
					Name:        "cloud_types",
					Description: `List of cloud types (determined based on policies assigned) ## Import Resources can be imported using the compliance standard ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_compliance_standard.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_compliance_standard_requirement",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cs_id",
					Description: `(Required) Compliance standard ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Compliance standard requirement name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "requirement_id",
					Description: `(Required) Compliance requirement number`,
				},
				resource.Attribute{
					Name:        "view_order",
					Description: `(Computed, int) View order ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "csr_id",
					Description: `Compliance standard requirement ID`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `(int) Created on`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "last_modified_on",
					Description: `(int) Last modified on`,
				},
				resource.Attribute{
					Name:        "system_default",
					Description: `(bool) System default`,
				},
				resource.Attribute{
					Name:        "policies_assigned_count",
					Description: `(int) Number of assigned policies`,
				},
				resource.Attribute{
					Name:        "standard_name",
					Description: `Compliance standard name ## Import Resources can be imported using the cs_id and the csr_id: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_compliance_standard_requirement.example 11111111-2222-3333-4444-555555555555:11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "csr_id",
					Description: `Compliance standard requirement ID`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `(int) Created on`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "last_modified_on",
					Description: `(int) Last modified on`,
				},
				resource.Attribute{
					Name:        "system_default",
					Description: `(bool) System default`,
				},
				resource.Attribute{
					Name:        "policies_assigned_count",
					Description: `(int) Number of assigned policies`,
				},
				resource.Attribute{
					Name:        "standard_name",
					Description: `Compliance standard name ## Import Resources can be imported using the cs_id and the csr_id: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_compliance_standard_requirement.example 11111111-2222-3333-4444-555555555555:11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_compliance_standard_requirement_section",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "csr_id",
					Description: `(Required) Compliance standard ID.`,
				},
				resource.Attribute{
					Name:        "section_id",
					Description: `(Required) Compliance section ID`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "view_order",
					Description: `(Computed, int) View order ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "csrs_id",
					Description: `Compliance standard requirement section ID`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `(int) Created on`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "last_modified_on",
					Description: `(int) Last modified on`,
				},
				resource.Attribute{
					Name:        "system_default",
					Description: `(bool) System default`,
				},
				resource.Attribute{
					Name:        "policies_assigned_count",
					Description: `(int) Number of assigned policies`,
				},
				resource.Attribute{
					Name:        "standard_name",
					Description: `Compliance standard name`,
				},
				resource.Attribute{
					Name:        "requirement_name",
					Description: `Compliance requirement name`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Section label`,
				},
				resource.Attribute{
					Name:        "associated_policy_ids",
					Description: `List of associated policy IDs ## Import Resources can be imported using the csr_id and the csrs_id: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_compliance_standard_requirement_section.example 11111111-2222-3333-4444-555555555555:11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "csrs_id",
					Description: `Compliance standard requirement section ID`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `(int) Created on`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "last_modified_on",
					Description: `(int) Last modified on`,
				},
				resource.Attribute{
					Name:        "system_default",
					Description: `(bool) System default`,
				},
				resource.Attribute{
					Name:        "policies_assigned_count",
					Description: `(int) Number of assigned policies`,
				},
				resource.Attribute{
					Name:        "standard_name",
					Description: `Compliance standard name`,
				},
				resource.Attribute{
					Name:        "requirement_name",
					Description: `Compliance requirement name`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Section label`,
				},
				resource.Attribute{
					Name:        "associated_policy_ids",
					Description: `List of associated policy IDs ## Import Resources can be imported using the csr_id and the csrs_id: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_compliance_standard_requirement_section.example 11111111-2222-3333-4444-555555555555:11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_datapattern",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Pattern name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Pattern description.`,
				},
				resource.Attribute{
					Name:        "detection_technique",
					Description: `Detection technique (default: ` + "`" + `regex` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "proximity_keywords",
					Description: `List of proximity keywords.`,
				},
				resource.Attribute{
					Name:        "regexes",
					Description: `(Required) List of regexes, as defined [below](#regexes). ### Regexes`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Required) Regular expression (match criteria for the data you want to find within your assets).`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(int) Weight to assign a score to a text entry (pattern match occurs when the score threshold is exceeded). Default: ` + "`" + `1` + "`" + `. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "pattern_id",
					Description: `Pattern ID.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Pattern mode (predefined or custom).`,
				},
				resource.Attribute{
					Name:        "entity",
					Description: `Entity value.`,
				},
				resource.Attribute{
					Name:        "grammar",
					Description: `Grammar value.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent ID for cloned data pattern.`,
				},
				resource.Attribute{
					Name:        "root_type",
					Description: `Root type (predefined or custom) for cloned data pattern.`,
				},
				resource.Attribute{
					Name:        "s3_path",
					Description: `S3 Path to the grammar.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `Updated by.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `(int) Last updated at.`,
				},
				resource.Attribute{
					Name:        "is_editable",
					Description: `(bool) Is editable. ## Import Resources can be imported using the data pattern ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_datapattern.example 111111111111111111111111 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "pattern_id",
					Description: `Pattern ID.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Pattern mode (predefined or custom).`,
				},
				resource.Attribute{
					Name:        "entity",
					Description: `Entity value.`,
				},
				resource.Attribute{
					Name:        "grammar",
					Description: `Grammar value.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent ID for cloned data pattern.`,
				},
				resource.Attribute{
					Name:        "root_type",
					Description: `Root type (predefined or custom) for cloned data pattern.`,
				},
				resource.Attribute{
					Name:        "s3_path",
					Description: `S3 Path to the grammar.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `Updated by.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `(int) Last updated at.`,
				},
				resource.Attribute{
					Name:        "is_editable",
					Description: `(bool) Is editable. ## Import Resources can be imported using the data pattern ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_datapattern.example 111111111111111111111111 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_dataprofile",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Profile Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Profile description.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `Type. Valid values are ` + "`" + `basic` + "`" + ` (default), or ` + "`" + `advance` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "profile_type",
					Description: `Profile Type. Valid values are ` + "`" + `custom` + "`" + ` (default), or ` + "`" + `system` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status. Valid values are ` + "`" + `non_hidden` + "`" + ` (default), or ` + "`" + `hidden` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "profile_status",
					Description: `Profile status. Valid values are ` + "`" + `active` + "`" + ` (default), or ` + "`" + `disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "data_patterns_rule_1",
					Description: `(Required) Model for DataProfile Rules, as defined [below](#data-patterns-rule-1). ### Data Patterns Rule 1`,
				},
				resource.Attribute{
					Name:        "operator_type",
					Description: `Pattern operator type. Default: ` + "`" + `or` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "data_pattern_rules",
					Description: `(Required) List of DataPattern Rules. Each item has data-pattern information, as defined [below](#data-pattern-rules). #### Data Pattern Rules`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Pattern name.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `(Required) Match type. Valid values are ` + "`" + `include` + "`" + `, or ` + "`" + `exclude` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "occurrence_operator_type",
					Description: `(Required) Occurrence operator type. Valid values are ` + "`" + `any` + "`" + `, ` + "`" + `more_than_equal_to` + "`" + `, ` + "`" + `less_than_equal_to` + "`" + `, or ` + "`" + `between` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "occurrence_count",
					Description: `(Required if value of ` + "`" + `occurrence_operator_type` + "`" + ` is ` + "`" + `more_than_equal_to` + "`" + ` or ` + "`" + `less_than_equal_to` + "`" + `) Occurrence count. Value must be a number between ` + "`" + `1` + "`" + ` and ` + "`" + `250` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "confidence_level",
					Description: `(Required) Confidence level.`,
				},
				resource.Attribute{
					Name:        "occurrence_high",
					Description: `(Required if value of ` + "`" + `occurrence_operator_type` + "`" + ` is ` + "`" + `between` + "`" + `) High occurrence value. Value must be a number between ` + "`" + `1` + "`" + ` and ` + "`" + `250` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "occurrence_low",
					Description: `(Required if value of ` + "`" + `occurrence_operator_type` + "`" + ` is ` + "`" + `between` + "`" + `) Low occurrence value. Value must be a number between ` + "`" + `1` + "`" + ` and ` + "`" + `250` + "`" + `. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `Profile ID.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `Updated by.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Created at (unix time).`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Updated at (unix time). In each ` + "`" + `Data Pattern Rules` + "`" + ` section, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "pattern_id",
					Description: `Pattern ID.`,
				},
				resource.Attribute{
					Name:        "detection_technique",
					Description: `Detection technique.`,
				},
				resource.Attribute{
					Name:        "supported_confidence_levels",
					Description: `List of supported confidence levels. ## Import Resources can be imported using the data profile ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_dataprofile.example 11111111 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `Profile ID.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `Updated by.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Created at (unix time).`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Updated at (unix time). In each ` + "`" + `Data Pattern Rules` + "`" + ` section, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "pattern_id",
					Description: `Pattern ID.`,
				},
				resource.Attribute{
					Name:        "detection_technique",
					Description: `Detection technique.`,
				},
				resource.Attribute{
					Name:        "supported_confidence_levels",
					Description: `List of supported confidence levels. ## Import Resources can be imported using the data profile ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_dataprofile.example 11111111 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_enterprise_settings",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "session_timeout",
					Description: `(int) Browser session timeout.`,
				},
				resource.Attribute{
					Name:        "access_key_max_validity",
					Description: `(int) Access Keys maximum validity in days.`,
				},
				resource.Attribute{
					Name:        "user_attribution_in_notification",
					Description: `(bool) User attribution in notification.`,
				},
				resource.Attribute{
					Name:        "require_alert_dismissal_note",
					Description: `(bool) Require alert dismissal note.`,
				},
				resource.Attribute{
					Name:        "default_policies_enabled",
					Description: `(Map of bools) Default policies enabled.`,
				},
				resource.Attribute{
					Name:        "apply_default_policies_enabled",
					Description: `(bool) Apply default policies enabled.`,
				},
				resource.Attribute{
					Name:        "alarm_enabled",
					Description: `(bool) Alarms enabled (Default : ` + "`" + `true` + "`" + `). Alarms are Prisma Cloud Platform health notifications which are generated to notify users of system level issues/errors. Disabling alarms will delete all existing alarms which were previously generated.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_integration",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the integration.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "integration_type",
					Description: `(Required) Integration type. Valid values are : ` + "`" + `okta_idp` + "`" + `, ` + "`" + `qualys` + "`" + `, ` + "`" + `tenable` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `splunk` + "`" + `, ` + "`" + `amazon_sqs` + "`" + `, ` + "`" + `webhook` + "`" + `, ` + "`" + `microsoft_teams` + "`" + `, ` + "`" + `azure_service_bus_queue` + "`" + `, ` + "`" + `service_now` + "`" + `, ` + "`" + `pager_duty` + "`" + `, ` + "`" + `demisto` + "`" + `, ` + "`" + `google_cscc` + "`" + `, ` + "`" + `aws_security_hub` + "`" + `, ` + "`" + `aws_s3` + "`" + `, ` + "`" + `snowflake` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Enabled. Default: ` + "`" + `true` + "`" + ` (For outbound integrations (i.e. all integrations except ` + "`" + `okta_idp` + "`" + `, ` + "`" + `qualys` + "`" + `, ` + "`" + `tenable` + "`" + `) this will always be ` + "`" + `true` + "`" + ` while creating, can be changed to ` + "`" + `false` + "`" + ` only while updating).`,
				},
				resource.Attribute{
					Name:        "integration_config",
					Description: `(Required) Integration configuration, the values depend on the integration type, as defined [below](#integration-config). ### Integration Config Refer to the [Prisma Cloud integration documentation](https://prisma.pan.dev/api/cloud/api-integration-config/) if you need more information on a specific integration.`,
				},
				resource.Attribute{
					Name:        "queue_url",
					Description: `(Required) The URL configured in the Azure Service Bus queue where Prisma cloud sends alerts.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required if you want to use the service principal-based access provided when the Azure cloud account was onboarded to Prisma Cloud) Azure account ID with service principal to which the Azure Service Bus queue belongs.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `(Required if you want to use a role with limited permissions) Azure Shared Access Signature connection string.`,
				},
				resource.Attribute{
					Name:        "queue_url",
					Description: `(Required) The Queue URL you used when you configured Prisma Cloud in Amazon SQS.`,
				},
				resource.Attribute{
					Name:        "more_info",
					Description: `(Optional, bool) Whether specific IAM credentials are specified for SQS queue access. Set it to ` + "`" + `true` + "`" + ` while configuring additional IAM information like ` + "`" + `role_arn` + "`" + ` and ` + "`" + `external_id` + "`" + ` or ` + "`" + `secret_key` + "`" + ` and ` + "`" + `access_key` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required if you want to use IAM access keys) AWS access key belonging to AWS IAM credentials meant for SQS queue access.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required when you are using IAM access keys) AWS secret key for the given access key.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Required if you want to use the IAM Role associated with Prisma Cloud) Role ARN associated with the IAM role on Prisma Cloud`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Required when you are using the IAM Role associated with Prisma Cloud) External ID associated with the IAM role on Prisma Cloud. New or updated value must be a unique 128-bit UUID.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `(Required) Qualys Login Username.`,
				},
				resource.Attribute{
					Name:        "base_url",
					Description: `(Required) Qualys Security Operations Center server API URL (without http(s)).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Qualys Password.`,
				},
				resource.Attribute{
					Name:        "host_url",
					Description: `(Required) ServiceNow URL.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `(Required) ServiceNow Login Username.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) ServiceNow password for login.`,
				},
				resource.Attribute{
					Name:        "tables",
					Description: `(Required, Map of bools) Key/value pairs that identify the ServiceNow module tables with which to integrate. The possible keys are: ` + "`" + `incident` + "`" + `, ` + "`" + `sn_si_incident` + "`" + `, ` + "`" + `em_event` + "`" + `. The possible values for each key are: ` + "`" + `true` + "`" + `, ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Webhook URL.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `(Optional) Webhook headers, as defined [below](#headers).`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Required) PagerDuty integration key.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `(Required) Slack webhook URL starting with ` + "`" + `https://hooks.slack.com/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_token",
					Description: `(Required) Splunk authentication token for the event collector.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Splunk HTTP event collector URL.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Optional) Splunk source type.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Webhook URL.`,
				},
				resource.Attribute{
					Name:        "host_url",
					Description: `(Required) The Cortex XSOAR instance FQDN/IP — either the name or the IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `(Required) The consumer key you configured when you created the Prisma Cloud application access in your Cortex XSOAR environment.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required) Secret key from Tenable.io.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required) Access key from Tenable.io.`,
				},
				resource.Attribute{
					Name:        "source_id",
					Description: `(Required) GCP source ID for the service account you used to onboard your GCP organization to Prisma Cloud.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) GCP organization ID.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Okta domain name.`,
				},
				resource.Attribute{
					Name:        "api_token",
					Description: `(Required) The authentication API token for Okta. The token must be of type Read-Only Admin.`,
				},
				resource.Attribute{
					Name:        "s3_uri",
					Description: `(Required) Amazon S3 bucket URI.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) AWS region where the S3 bucket resides.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Required) Role ARN associated with the IAM role on Prisma Cloud.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Required) External ID associated with the IAM role on Prisma Cloud. Any new or updated value must be a unique 128-bit UUID.`,
				},
				resource.Attribute{
					Name:        "roll_up_interval",
					Description: `(Required, int) Time in minutes at which batching of Prisma Cloud alerts would roll up. Valid values are ` + "`" + `15` + "`" + `, ` + "`" + `30` + "`" + `, ` + "`" + `60` + "`" + `, or ` + "`" + `180` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) AWS account ID to which you assigned AWS Security Hub read-only access.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Required) List of AWS regions, as defined [below](#regions).`,
				},
				resource.Attribute{
					Name:        "host_url",
					Description: `(Required) Snowflake Account URL. Format should be 'YOURACCOUNTNAME.snowflakecomputing.com'.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) Snowflake Username.`,
				},
				resource.Attribute{
					Name:        "staging_integration_id",
					Description: `(Required) Existing Amazon S3 integration ID.`,
				},
				resource.Attribute{
					Name:        "pipe_name",
					Description: `(Required) Snowpipe Name. Format should be '<db_name>.<schema_name>.<pipe_name>'.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) Private Key.`,
				},
				resource.Attribute{
					Name:        "pass_phrase",
					Description: `(Required if you are using encrypted key) PassPhrase for private key.`,
				},
				resource.Attribute{
					Name:        "roll_up_interval",
					Description: `(Required, int) Time in minutes at which batching of Prisma Cloud alerts would roll up. Valid values are ` + "`" + `15` + "`" + `, ` + "`" + `30` + "`" + `, ` + "`" + `60` + "`" + `, or ` + "`" + `180` + "`" + `. #### Headers`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Header name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Header value.`,
				},
				resource.Attribute{
					Name:        "secure",
					Description: `(bool) Secure.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(bool) Read-only. #### Regions`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `AWS region name e.g. ` + "`" + `AWS California` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "api_identifier",
					Description: `AWS region code e.g. ` + "`" + `us-west-1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud Type (default: ` + "`" + `aws` + "`" + `). ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "integration_id",
					Description: `Integration ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `(bool) Valid.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by.`,
				},
				resource.Attribute{
					Name:        "created_ts",
					Description: `(int) Created timestamp.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_ts",
					Description: `(int) Last modified timestamp.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `Model for the integration status details, as defined [below](#reason). ### Reason`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `(int) Last updated.`,
				},
				resource.Attribute{
					Name:        "error_type",
					Description: `Error type.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Message.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Model for message details, as defined [below](#details). ### Details`,
				},
				resource.Attribute{
					Name:        "status_code",
					Description: `(int) Status code.`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `Subject.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Internationalization key. In ` + "`" + `integration_config` + "`" + ` section, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Cortex release version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "integration_id",
					Description: `Integration ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `(bool) Valid.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by.`,
				},
				resource.Attribute{
					Name:        "created_ts",
					Description: `(int) Created timestamp.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_ts",
					Description: `(int) Last modified timestamp.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `Model for the integration status details, as defined [below](#reason). ### Reason`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `(int) Last updated.`,
				},
				resource.Attribute{
					Name:        "error_type",
					Description: `Error type.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Message.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Model for message details, as defined [below](#details). ### Details`,
				},
				resource.Attribute{
					Name:        "status_code",
					Description: `(int) Status code.`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `Subject.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Internationalization key. In ` + "`" + `integration_config` + "`" + ` section, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Cortex release version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_org_cloud_account",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "disable_on_destroy",
					Description: `(Optional,bool) To disable cloud account instead of deleting when calling Terraform destroy (default: ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "aws",
					Description: `AWS org account type spec, defined [below](#aws).`,
				},
				resource.Attribute{
					Name:        "azure",
					Description: `Azure org account type spec, defined [below](#azure).`,
				},
				resource.Attribute{
					Name:        "gcp",
					Description: `GCP org account type spec, defined [below](#gcp).`,
				},
				resource.Attribute{
					Name:        "oci",
					Description: `Oci account type spec, defined [below](#oci). ### AWS`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) AWS Org account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, bool) Whether or not the account is enabled (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Required) AWS org account external ID.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `(Required) List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Required) Unique identifier for an AWS org resource (ARN).`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `(Optional) Defaults to "organization" if not specified.`,
				},
				resource.Attribute{
					Name:        "member_external_id",
					Description: `(Required) AWS org Member account role's external ID.`,
				},
				resource.Attribute{
					Name:        "member_role_status",
					Description: `(Optional, bool) - True = The member role created using stack set exists in all the member accounts. All the Org accounts will be added. false = Only the master account will be added(Default = False).`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `(Optional) Defaults to "MONITOR". Valid values : ` + "`" + `MONITOR` + "`" + ` or ` + "`" + `MONITOR_AND_PROTECT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "hierarchy_selection",
					Description: `(Optional) List of AWS Organization Units (OU), AWS accounts, and AWS Organizations to onboard under this organization. [below](#For-AWS) ### Azure`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) Azure org account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, bool) Whether or not the account is enabled (defualt: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `(Required) List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) Application ID registered with Active Directory.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Application ID key.`,
				},
				resource.Attribute{
					Name:        "monitor_flow_logs",
					Description: `(Required, bool) Automatically ingest flow logs.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Active Directory ID associated with Azure.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `(Required) Unique ID of the service principal object associated with the Prisma Cloud application that you create.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `(Optional) Defaults to "tenant" if not specified.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `(Optional) Defaults to "MONITOR". Valid values : ` + "`" + `MONITOR` + "`" + ` ### GCP`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) GCP org project ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, bool) Whether or not the account is enabled (defualt: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "compression_enabled",
					Description: `(Optional, bool) Enable flow log compression.`,
				},
				resource.Attribute{
					Name:        "data_flow_enabled_project",
					Description: `(Optional) GCP project for flow log compression.`,
				},
				resource.Attribute{
					Name:        "flow_log_storage_bucket",
					Description: `(Optional) GCP Flow logs storage bucket.`,
				},
				resource.Attribute{
					Name:        "credentials_json",
					Description: `(Required) Content of the JSON credentials file (read in using ` + "`" + `file()` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `(Optional) Account type. Defaults to ` + "`" + `organization` + "`" + ` if not specified.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `(Optional) Protection Mode. Valid values : ` + "`" + `MONITOR` + "`" + ` or ` + "`" + `MONITOR_AND_PROTECT` + "`" + `. Defaults to ` + "`" + `MONITOR` + "`" + ` if not specified.`,
				},
				resource.Attribute{
					Name:        "organization_name",
					Description: `(Required) GCP org organization name.`,
				},
				resource.Attribute{
					Name:        "account_group_creation_mode",
					Description: `(Optional) Cloud account group creation mode. Valid values : ` + "`" + `MANUAL` + "`" + `: Create account groups manually, ` + "`" + `AUTO` + "`" + `: Create high-level account groups based on folders identified, or ` + "`" + `RECURSIVE` + "`" + `: Drill down in folder tree to create account groups (default : ` + "`" + `MANUAL` + "`" + `). ` + "`" + `AUTO` + "`" + ` can't be used if ` + "`" + `selection_type` + "`" + ` in ` + "`" + `hierarchy_selection` + "`" + ` is ` + "`" + `EXCLUDE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hierarchy_selection",
					Description: `(Optional) List of hierarchy selection. Each item has resource ID, display name, node type and selection type, as defined [below](#For-GCP). #### Hierarchy Selection ##### For AWS`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) Resource ID. Valid values are AWS OU ID, AWS account ID, or AWS Organization ID. Note you must escape any double quotes in the resource ID with a backslash.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name for AWS OU, AWS account, or AWS organization.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Required) Valid values: ` + "`" + `OU` + "`" + `, ` + "`" + `ACCOUNT` + "`" + `, ` + "`" + `ORG` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "selection_type",
					Description: `(Required) Valid values: ` + "`" + `INCLUDE` + "`" + ` to include the specified resource to onboard, ` + "`" + `EXCLUDE` + "`" + ` to exclude the specified resource and onboard the rest, ` + "`" + `ALL` + "`" + ` to onboard all resources in the organization. ##### For GCP`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) Resource ID. For folders, format is folders/{folder ID}. For projects, format is {project number}. For orgs, format is organizations/{org ID}`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Display name for folder, project, or organization.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Required) Node type. Valid values : ` + "`" + `FOLDER` + "`" + `, ` + "`" + `PROJECT` + "`" + `, or ` + "`" + `ORG` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "selection_type",
					Description: `(Required) Selection type. Valid values: ` + "`" + `INCLUDE` + "`" + `, ` + "`" + `EXCLUDE` + "`" + `, or ` + "`" + `ALL` + "`" + `. If ` + "`" + `node_type` + "`" + ` is ` + "`" + `PROJECT` + "`" + ` or ` + "`" + `FOLDER` + "`" + `, then a valid value is either ` + "`" + `INCLUDE` + "`" + ` or ` + "`" + `EXCLUDE` + "`" + `. ### Oci`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) OCI account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, bool) Whether or not the account is enabled (default: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) OCI identity group name that you define. Can be an existing group.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `(Required) account ID to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) OCI identity group name that you define. Can be an existing group.`,
				},
				resource.Attribute{
					Name:        "ram_arn",
					Description: `(Required) Unique identifier for an Alibaba RAM role resource.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `(Required) Account type - account or tenant.`,
				},
				resource.Attribute{
					Name:        "default_account_group_id",
					Description: `(Required) account ID to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "home_region",
					Description: `(Required) OCI tenancy home region.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required) OCI identity policy name that you define. Can be an existing policy that has the right policy statements.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) OCI identity user name that you define. Can be an existing user that has the right privileges.`,
				},
				resource.Attribute{
					Name:        "user_ocid",
					Description: `(Required) OCI identity user Ocid that you define. Can be an existing user that has the right privileges. ## Import Resources can be imported using the cloud type (` + "`" + `aws` + "`" + `, ` + "`" + `azure` + "`" + `, ` + "`" + `gcp` + "`" + `, or ` + "`" + `oci` + "`" + `) and the ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_org_cloud_account.aws_example aws:accountIdHere ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_policy",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Policy name`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Required) Policy type. Valid values are ` + "`" + `config` + "`" + `, ` + "`" + `audit_event` + "`" + `, ` + "`" + `iam` + "`" + `, ` + "`" + `network` + "`" + `, ` + "`" + `data` + "`" + `, or ` + "`" + `anomaly` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Severity. Valid values are ` + "`" + `low` + "`" + ` (default), ` + "`" + `medium` + "`" + `, or ` + "`" + `high` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "recommendation",
					Description: `Remediation recommendation`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud type (Optional for policies having RQL query with multiway joins, otherwise required) - valid values are ` + "`" + `aws` + "`" + `,` + "`" + `azure` + "`" + `,` + "`" + `gcp` + "`" + `,` + "`" + `alibaba_cloud` + "`" + ` and ` + "`" + `all` + "`" + ``,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `List of labels`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Enabled`,
				},
				resource.Attribute{
					Name:        "overridden",
					Description: `(bool) Overridden`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(bool) Deleted`,
				},
				resource.Attribute{
					Name:        "restrict_alert_dismissal",
					Description: `(bool) Restrict alert dismissal`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) Model for the rule, as defined [below](#rule)`,
				},
				resource.Attribute{
					Name:        "remediation",
					Description: `Model for remediation, as defined [below](#remediation)`,
				},
				resource.Attribute{
					Name:        "compliance_metadata",
					Description: `List of compliance data. Each item has compliance standard, requirement, and/or section information, as defined [below](#compliance-metadata) ### Rule One and only one of these must be present:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud type`,
				},
				resource.Attribute{
					Name:        "cloud_account",
					Description: `Cloud account`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Resource type`,
				},
				resource.Attribute{
					Name:        "api_name",
					Description: `API name`,
				},
				resource.Attribute{
					Name:        "resource_id_path",
					Description: `Resource ID path`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `(Required for Config, Audit Event, IAM and Network policies) Saved search ID that defines the rule criteria`,
				},
				resource.Attribute{
					Name:        "data_criteria",
					Description: `(Required for Data policy) Criteria for DLP Rule, as defined [below](#data-criteria)`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Required for Config, Audit Event, IAM and Network policies, map of strings) Parameters. Valid keys are ` + "`" + `withIac` + "`" + ` and ` + "`" + `savedSearch` + "`" + ` and value is ` + "`" + `"true"` + "`" + `or ` + "`" + `"false"` + "`" + ` (` + "`" + `SavedSearch` + "`" + ` is true when we are using savedsearch and it is false when we directly give search query and ` + "`" + `withIac` + "`" + ` is true for build policies otherwise false)`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `(Required) Type of rule or RQL query. Valid values are ` + "`" + `Config` + "`" + `, ` + "`" + `AuditEvent` + "`" + `, ` + "`" + `IAM` + "`" + `, ` + "`" + `Network` + "`" + `, ` + "`" + `DLP` + "`" + `, or ` + "`" + `Anomaly` + "`" + ` ### Remediation This section may be present or may be omitted:`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `Template type`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "cli_script_template",
					Description: `CLI script template`,
				},
				resource.Attribute{
					Name:        "cli_script_json_schema_string",
					Description: `CLI script JSON schema ### Compliance Metadata`,
				},
				resource.Attribute{
					Name:        "compliance_id",
					Description: `(Required) Compliance Section UUID #### Data Criteria`,
				},
				resource.Attribute{
					Name:        "classification_result",
					Description: `(Required) Data Profile name required for DLP rule criteria`,
				},
				resource.Attribute{
					Name:        "exposure",
					Description: `File exposure. Valid values are ` + "`" + `private` + "`" + `, ` + "`" + `public` + "`" + `, or ` + "`" + `conditional` + "`" + ``,
				},
				resource.Attribute{
					Name:        "extension",
					Description: `List of file extensions ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Policy ID`,
				},
				resource.Attribute{
					Name:        "policy_subtypes",
					Description: `Policy subtypes`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `(int) Created on`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "last_modified_on",
					Description: `(int) Last modified on`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "rule_last_modified_on",
					Description: `(int) Rule last modified on`,
				},
				resource.Attribute{
					Name:        "open_alerts_count",
					Description: `(int) Open alerts count`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Owner`,
				},
				resource.Attribute{
					Name:        "policy_mode",
					Description: `Policy mode`,
				},
				resource.Attribute{
					Name:        "policy_category",
					Description: `Policy category`,
				},
				resource.Attribute{
					Name:        "policy_class",
					Description: `Policy class`,
				},
				resource.Attribute{
					Name:        "system_default",
					Description: `(bool) If policy is a system default policy or not`,
				},
				resource.Attribute{
					Name:        "remediable",
					Description: `(bool) Is remediable or not In each ` + "`" + `Compliance Metadata` + "`" + ` section, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "standard_name",
					Description: `Compliance standard name`,
				},
				resource.Attribute{
					Name:        "standard_description",
					Description: `Compliance standard description`,
				},
				resource.Attribute{
					Name:        "requirement_name",
					Description: `Requirement name`,
				},
				resource.Attribute{
					Name:        "requirement_id",
					Description: `Requirement ID`,
				},
				resource.Attribute{
					Name:        "requirement_description",
					Description: `Requirement description`,
				},
				resource.Attribute{
					Name:        "section_description",
					Description: `Section description`,
				},
				resource.Attribute{
					Name:        "section_id",
					Description: `Section ID`,
				},
				resource.Attribute{
					Name:        "section_label",
					Description: `Section label`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Policy ID`,
				},
				resource.Attribute{
					Name:        "custom_assigned",
					Description: `(bool) Custom assigned ## Import Resources can be imported using the policy ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_policy.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `Policy ID`,
				},
				resource.Attribute{
					Name:        "policy_subtypes",
					Description: `Policy subtypes`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `(int) Created on`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "last_modified_on",
					Description: `(int) Last modified on`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "rule_last_modified_on",
					Description: `(int) Rule last modified on`,
				},
				resource.Attribute{
					Name:        "open_alerts_count",
					Description: `(int) Open alerts count`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Owner`,
				},
				resource.Attribute{
					Name:        "policy_mode",
					Description: `Policy mode`,
				},
				resource.Attribute{
					Name:        "policy_category",
					Description: `Policy category`,
				},
				resource.Attribute{
					Name:        "policy_class",
					Description: `Policy class`,
				},
				resource.Attribute{
					Name:        "system_default",
					Description: `(bool) If policy is a system default policy or not`,
				},
				resource.Attribute{
					Name:        "remediable",
					Description: `(bool) Is remediable or not In each ` + "`" + `Compliance Metadata` + "`" + ` section, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "standard_name",
					Description: `Compliance standard name`,
				},
				resource.Attribute{
					Name:        "standard_description",
					Description: `Compliance standard description`,
				},
				resource.Attribute{
					Name:        "requirement_name",
					Description: `Requirement name`,
				},
				resource.Attribute{
					Name:        "requirement_id",
					Description: `Requirement ID`,
				},
				resource.Attribute{
					Name:        "requirement_description",
					Description: `Requirement description`,
				},
				resource.Attribute{
					Name:        "section_description",
					Description: `Section description`,
				},
				resource.Attribute{
					Name:        "section_id",
					Description: `Section ID`,
				},
				resource.Attribute{
					Name:        "section_label",
					Description: `Section label`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Policy ID`,
				},
				resource.Attribute{
					Name:        "custom_assigned",
					Description: `(bool) Custom assigned ## Import Resources can be imported using the policy ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_policy.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_report",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Report Name`,
				},
				resource.Attribute{
					Name:        "report_type",
					Description: `(Required) Report type. Valid values are ` + "`" + `RIS` + "`" + ` (for Cloud Security Assessment report) , ` + "`" + `INVENTORY_OVERVIEW` + "`" + ` (for Business Unit report), ` + "`" + `INVENTORY_DETAIL` + "`" + ` (for Detailed Business Unit report), or name of a compliance standard (for Compliance report)`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud type`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) Model for report target, as defined [below](#target) ### Target There should be one and only one target block:`,
				},
				resource.Attribute{
					Name:        "account_groups",
					Description: `List of cloud account groups`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `List of cloud accounts`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `List of regions`,
				},
				resource.Attribute{
					Name:        "compliance_standard_ids",
					Description: `List of compliance standard IDs (For Business Unit Report and Detailed Business Unit Report)`,
				},
				resource.Attribute{
					Name:        "resource_groups",
					Description: `List of resource groups`,
				},
				resource.Attribute{
					Name:        "notify_to",
					Description: `List of email addresses to receive notification (not supported for Cloud Security Assessment Report)`,
				},
				resource.Attribute{
					Name:        "compression_enabled",
					Description: `(bool) Business unit detailed report compression enabled (For Detailed Business Unit Report)`,
				},
				resource.Attribute{
					Name:        "download_now",
					Description: `(bool) True = download now`,
				},
				resource.Attribute{
					Name:        "schedule_enabled",
					Description: `(bool) Report scheduling enabled (not supported for Cloud Security Assessment Report)`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `Recurring report schedule in RRULE format (not supported for Cloud Security Assessment Report)`,
				},
				resource.Attribute{
					Name:        "notification_template_id",
					Description: `Notification template id (not supported for Cloud Security Assessment Report)`,
				},
				resource.Attribute{
					Name:        "time_range",
					Description: `(Required) The time range spec, as defined [below](#time-range). ### Time Range Only one of these can be defined:`,
				},
				resource.Attribute{
					Name:        "absolute",
					Description: `An absolute time range spec, as defined [below](#absolute-time-range).`,
				},
				resource.Attribute{
					Name:        "relative",
					Description: `A relative time range spec, as defined [below](#relative-time-range).`,
				},
				resource.Attribute{
					Name:        "to_now",
					Description: `A "To Now" time range spec, as defined [below](#to-now-time-range). ### Absolute Time Range`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required, int) Start time.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required, int) End time. ### Relative Time Range`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `(Required, int) The time number.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Required) The time unit. ### To Now Time Range`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Required) The time unit. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "report_id",
					Description: `Report ID`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `(int) Created on`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "last_modified_on",
					Description: `(int) Last modified on`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "compliance_standard_id",
					Description: `Compliance Standard ID`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Report status`,
				},
				resource.Attribute{
					Name:        "next_schedule",
					Description: `(int) Next schedule`,
				},
				resource.Attribute{
					Name:        "last_scheduled",
					Description: `(int) Last scheduled`,
				},
				resource.Attribute{
					Name:        "total_instance_count",
					Description: `(int) Total instance count`,
				},
				resource.Attribute{
					Name:        "counts",
					Description: `Model for compliance aggregate count, as defined [below](#counts). ### Counts`,
				},
				resource.Attribute{
					Name:        "failed",
					Description: `(int) Failed`,
				},
				resource.Attribute{
					Name:        "high_severity_failed",
					Description: `(int) Number of high-severity failures`,
				},
				resource.Attribute{
					Name:        "low_severity_failed",
					Description: `(int) Number of low-severity failures`,
				},
				resource.Attribute{
					Name:        "medium_severity_failed",
					Description: `(int) Number of medium-severity failures`,
				},
				resource.Attribute{
					Name:        "passed",
					Description: `(int) Passed`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total ## Import Resources can be imported using the report ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_report.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "report_id",
					Description: `Report ID`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `(int) Created on`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "last_modified_on",
					Description: `(int) Last modified on`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "compliance_standard_id",
					Description: `Compliance Standard ID`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Report status`,
				},
				resource.Attribute{
					Name:        "next_schedule",
					Description: `(int) Next schedule`,
				},
				resource.Attribute{
					Name:        "last_scheduled",
					Description: `(int) Last scheduled`,
				},
				resource.Attribute{
					Name:        "total_instance_count",
					Description: `(int) Total instance count`,
				},
				resource.Attribute{
					Name:        "counts",
					Description: `Model for compliance aggregate count, as defined [below](#counts). ### Counts`,
				},
				resource.Attribute{
					Name:        "failed",
					Description: `(int) Failed`,
				},
				resource.Attribute{
					Name:        "high_severity_failed",
					Description: `(int) Number of high-severity failures`,
				},
				resource.Attribute{
					Name:        "low_severity_failed",
					Description: `(int) Number of low-severity failures`,
				},
				resource.Attribute{
					Name:        "medium_severity_failed",
					Description: `(int) Number of medium-severity failures`,
				},
				resource.Attribute{
					Name:        "passed",
					Description: `(int) Passed`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total ## Import Resources can be imported using the report ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_report.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_rql_search",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "search_type",
					Description: `(Required) The search type. Valid values are ` + "`" + `config` + "`" + ` (default) ` + "`" + `event` + "`" + `, ` + "`" + `network` + "`" + ` and ` + "`" + `iam` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) The RQL query.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(int) Limit rules (default: ` + "`" + `10` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "time_range",
					Description: `(Required for config, event and network RQL search) The RQL time range spec, as defined [below](#time-range). ### Time Range Only one of these can be defined:`,
				},
				resource.Attribute{
					Name:        "absolute",
					Description: `An absolute time range spec, as defined [below](#absolute-time-range).`,
				},
				resource.Attribute{
					Name:        "relative",
					Description: `A relative time range spec, as defined [below](#relative-time-range).`,
				},
				resource.Attribute{
					Name:        "to_now",
					Description: `A "To Now" time range spec, as defined [below](#to-now-time-range). ### Absolute Time Range`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required, int) Start time.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required, int) End time. ### Relative Time Range`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `(Required, int) The time number.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Required) The time unit. ### To Now Time Range`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Required) The time unit. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "search_id",
					Description: `The search ID returned from a successful RQL query.`,
				},
				resource.Attribute{
					Name:        "group_by",
					Description: `(list) Group by.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `The cloud type.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "config_data",
					Description: `(for ` + "`" + `search_type="config"` + "`" + `, list) List of config_data specs, as defined below.`,
				},
				resource.Attribute{
					Name:        "event_data",
					Description: `(For ` + "`" + `search_type="event"` + "`" + `, list) List of event_data specs, as defined below.`,
				},
				resource.Attribute{
					Name:        "network_data",
					Description: `(For ` + "`" + `search_type="network"` + "`" + `, list) List of network_data specs, as defined below.`,
				},
				resource.Attribute{
					Name:        "iam_data",
					Description: `(For ` + "`" + `search_type="iam"` + "`" + `, list) List of iam_data specs, as defined below. ` + "`" + `config_data` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "state_id",
					Description: `The state ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL. ` + "`" + `event_data` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "account",
					Description: `Account.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(int) Region ID.`,
				},
				resource.Attribute{
					Name:        "region_api_identifier",
					Description: `Region API identifier. ` + "`" + `network_data` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "account",
					Description: `Account.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(int) Region ID.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Account Name. ` + "`" + `iam_data` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "accessed_resources_count",
					Description: `(int) Accessed resource count.`,
				},
				resource.Attribute{
					Name:        "dest_cloud_account",
					Description: `Destination cloud account.`,
				},
				resource.Attribute{
					Name:        "dest_cloud_region",
					Description: `Destination cloud region.`,
				},
				resource.Attribute{
					Name:        "dest_cloud_resource_rrn",
					Description: `Destination cloud resource RRN.`,
				},
				resource.Attribute{
					Name:        "dest_cloud_service_name",
					Description: `Destination cloud service name.`,
				},
				resource.Attribute{
					Name:        "dest_cloud_type",
					Description: `Destination cloud type.`,
				},
				resource.Attribute{
					Name:        "dest_resource_id",
					Description: `Destination cloud resource id.`,
				},
				resource.Attribute{
					Name:        "dest_resource_name",
					Description: `Destination cloud resource name.`,
				},
				resource.Attribute{
					Name:        "dest_resource_type",
					Description: `Destination cloud resource type.`,
				},
				resource.Attribute{
					Name:        "effective_action_name",
					Description: `Effective action name.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_entity_id",
					Description: `Granted by cloud entity id.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_entity_name",
					Description: `Granted by cloud entity name.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_entity_rrn",
					Description: `Granted by cloud entity rrn.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_entity_type",
					Description: `Granted by cloud entity type.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_policy_id",
					Description: `Granted by cloud policy id.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_policy_name",
					Description: `Granted by cloud policy name.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_policy_rrn",
					Description: `Granted by cloud policy rrn.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_policy_type",
					Description: `Granted by cloud policy type.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_type",
					Description: `Granted by cloud type.`,
				},
				resource.Attribute{
					Name:        "message_id",
					Description: `Message id.`,
				},
				resource.Attribute{
					Name:        "is_wild_card_dest_cloud_resource_name",
					Description: `(bool) Is destination cloud resource name a wildcard.`,
				},
				resource.Attribute{
					Name:        "last_access_date",
					Description: `Last access date.`,
				},
				resource.Attribute{
					Name:        "source_cloud_account",
					Description: `Source cloud account.`,
				},
				resource.Attribute{
					Name:        "source_cloud_region",
					Description: `Source cloud region.`,
				},
				resource.Attribute{
					Name:        "source_cloud_resource_rrn",
					Description: `Source cloud resource rrn.`,
				},
				resource.Attribute{
					Name:        "source_cloud_service_name",
					Description: `Source cloud service name.`,
				},
				resource.Attribute{
					Name:        "source_cloud_type",
					Description: `Source cloud type.`,
				},
				resource.Attribute{
					Name:        "source_idp_domain",
					Description: `Source IDP domain.`,
				},
				resource.Attribute{
					Name:        "source_idp_email",
					Description: `Source IDP email.`,
				},
				resource.Attribute{
					Name:        "source_idp_group",
					Description: `Source IDP group.`,
				},
				resource.Attribute{
					Name:        "source_idp_rrn",
					Description: `Source IDP rrn.`,
				},
				resource.Attribute{
					Name:        "source_idp_service",
					Description: `Source IDP service.`,
				},
				resource.Attribute{
					Name:        "source_idp_user_name",
					Description: `Source IDP user name.`,
				},
				resource.Attribute{
					Name:        "source_public",
					Description: `(bool) Is source public.`,
				},
				resource.Attribute{
					Name:        "source_resource_id",
					Description: `Source cloud resource id.`,
				},
				resource.Attribute{
					Name:        "source_resource_type",
					Description: `Source cloud resource type.`,
				},
				resource.Attribute{
					Name:        "exceptions",
					Description: `(list) Permission exception list, as defined below. ` + "`" + `exceptions` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "message_code",
					Description: `Message code.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "search_id",
					Description: `The search ID returned from a successful RQL query.`,
				},
				resource.Attribute{
					Name:        "group_by",
					Description: `(list) Group by.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `The cloud type.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "config_data",
					Description: `(for ` + "`" + `search_type="config"` + "`" + `, list) List of config_data specs, as defined below.`,
				},
				resource.Attribute{
					Name:        "event_data",
					Description: `(For ` + "`" + `search_type="event"` + "`" + `, list) List of event_data specs, as defined below.`,
				},
				resource.Attribute{
					Name:        "network_data",
					Description: `(For ` + "`" + `search_type="network"` + "`" + `, list) List of network_data specs, as defined below.`,
				},
				resource.Attribute{
					Name:        "iam_data",
					Description: `(For ` + "`" + `search_type="iam"` + "`" + `, list) List of iam_data specs, as defined below. ` + "`" + `config_data` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "state_id",
					Description: `The state ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL. ` + "`" + `event_data` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "account",
					Description: `Account.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(int) Region ID.`,
				},
				resource.Attribute{
					Name:        "region_api_identifier",
					Description: `Region API identifier. ` + "`" + `network_data` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "account",
					Description: `Account.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(int) Region ID.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Account Name. ` + "`" + `iam_data` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "accessed_resources_count",
					Description: `(int) Accessed resource count.`,
				},
				resource.Attribute{
					Name:        "dest_cloud_account",
					Description: `Destination cloud account.`,
				},
				resource.Attribute{
					Name:        "dest_cloud_region",
					Description: `Destination cloud region.`,
				},
				resource.Attribute{
					Name:        "dest_cloud_resource_rrn",
					Description: `Destination cloud resource RRN.`,
				},
				resource.Attribute{
					Name:        "dest_cloud_service_name",
					Description: `Destination cloud service name.`,
				},
				resource.Attribute{
					Name:        "dest_cloud_type",
					Description: `Destination cloud type.`,
				},
				resource.Attribute{
					Name:        "dest_resource_id",
					Description: `Destination cloud resource id.`,
				},
				resource.Attribute{
					Name:        "dest_resource_name",
					Description: `Destination cloud resource name.`,
				},
				resource.Attribute{
					Name:        "dest_resource_type",
					Description: `Destination cloud resource type.`,
				},
				resource.Attribute{
					Name:        "effective_action_name",
					Description: `Effective action name.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_entity_id",
					Description: `Granted by cloud entity id.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_entity_name",
					Description: `Granted by cloud entity name.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_entity_rrn",
					Description: `Granted by cloud entity rrn.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_entity_type",
					Description: `Granted by cloud entity type.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_policy_id",
					Description: `Granted by cloud policy id.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_policy_name",
					Description: `Granted by cloud policy name.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_policy_rrn",
					Description: `Granted by cloud policy rrn.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_policy_type",
					Description: `Granted by cloud policy type.`,
				},
				resource.Attribute{
					Name:        "granted_by_cloud_type",
					Description: `Granted by cloud type.`,
				},
				resource.Attribute{
					Name:        "message_id",
					Description: `Message id.`,
				},
				resource.Attribute{
					Name:        "is_wild_card_dest_cloud_resource_name",
					Description: `(bool) Is destination cloud resource name a wildcard.`,
				},
				resource.Attribute{
					Name:        "last_access_date",
					Description: `Last access date.`,
				},
				resource.Attribute{
					Name:        "source_cloud_account",
					Description: `Source cloud account.`,
				},
				resource.Attribute{
					Name:        "source_cloud_region",
					Description: `Source cloud region.`,
				},
				resource.Attribute{
					Name:        "source_cloud_resource_rrn",
					Description: `Source cloud resource rrn.`,
				},
				resource.Attribute{
					Name:        "source_cloud_service_name",
					Description: `Source cloud service name.`,
				},
				resource.Attribute{
					Name:        "source_cloud_type",
					Description: `Source cloud type.`,
				},
				resource.Attribute{
					Name:        "source_idp_domain",
					Description: `Source IDP domain.`,
				},
				resource.Attribute{
					Name:        "source_idp_email",
					Description: `Source IDP email.`,
				},
				resource.Attribute{
					Name:        "source_idp_group",
					Description: `Source IDP group.`,
				},
				resource.Attribute{
					Name:        "source_idp_rrn",
					Description: `Source IDP rrn.`,
				},
				resource.Attribute{
					Name:        "source_idp_service",
					Description: `Source IDP service.`,
				},
				resource.Attribute{
					Name:        "source_idp_user_name",
					Description: `Source IDP user name.`,
				},
				resource.Attribute{
					Name:        "source_public",
					Description: `(bool) Is source public.`,
				},
				resource.Attribute{
					Name:        "source_resource_id",
					Description: `Source cloud resource id.`,
				},
				resource.Attribute{
					Name:        "source_resource_type",
					Description: `Source cloud resource type.`,
				},
				resource.Attribute{
					Name:        "exceptions",
					Description: `(list) Permission exception list, as defined below. ` + "`" + `exceptions` + "`" + ` supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "message_code",
					Description: `Message code.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_saved_search",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "query",
					Description: `(Required) The RQL query.`,
				},
				resource.Attribute{
					Name:        "search_id",
					Description: `(Required) The search ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "time_range",
					Description: `(Required) The RQL time range spec, as defined [below](#time-range). ### Time Range Only one of these can be defined:`,
				},
				resource.Attribute{
					Name:        "absolute",
					Description: `An absolute time range spec, as defined [below](#absolute-time-range).`,
				},
				resource.Attribute{
					Name:        "relative",
					Description: `A relative time range spec, as defined [below](#relative-time-range).`,
				},
				resource.Attribute{
					Name:        "to_now",
					Description: `A "To Now" time range spec, as defined [below](#to-now-time-range). ### Absolute Time Range`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required, int) Start time.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required, int) End time. ### Relative Time Range`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `(Required, int) The time number.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Required) The time unit. ### To Now Time Range`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Required) The time unit. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "saved",
					Description: `(bool) This is set to ` + "`" + `true` + "`" + ` when the saved search is created. ## Import Resources can be imported using the saved-search ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_saved_search.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "saved",
					Description: `(bool) This is set to ` + "`" + `true` + "`" + ` when the saved search is created. ## Import Resources can be imported using the saved-search ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_saved_search.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_user_profile",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_type",
					Description: `(Optional) Account Type. Valid values are ` + "`" + `USER_ACCOUNT` + "`" + `, or ` + "`" + `SERVICE_ACCOUNT` + "`" + `. (default: ` + "`" + `USER_ACCOUNT` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) User email or service account name.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `(Required if ` + "`" + `account_type` + "`" + ` is ` + "`" + `USER_ACCOUNT` + "`" + `) First name.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `(Required if ` + "`" + `account_type` + "`" + ` is ` + "`" + `USER_ACCOUNT` + "`" + `) Last name.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required if ` + "`" + `account_type` + "`" + ` is ` + "`" + `USER_ACCOUNT` + "`" + `) Email ID.`,
				},
				resource.Attribute{
					Name:        "access_key_expiration",
					Description: `(Optional, int) Access key expiration timestamp in milliseconds for ` + "`" + `SERVICE_ACCOUNT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "access_key_name",
					Description: `(Required if ` + "`" + `account_type` + "`" + ` is ` + "`" + `SERVICE_ACCOUNT` + "`" + `) Access key name.`,
				},
				resource.Attribute{
					Name:        "enable_key_expiration",
					Description: `(Optional, bool) Enable access key expiration. (default: ` + "`" + `false` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "role_ids",
					Description: `(Required) List of Role IDs. (default: ` + "`" + `false` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "default_role_id",
					Description: `(Required) Default Role ID, must be present in ` + "`" + `role_ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Required) Time zone (e.g. America/Los_Angeles).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, bool) Is account enabled. (default: ` + "`" + `true` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "access_keys_allowed",
					Description: `(Optional, bool) Access keys allowed. (For ` + "`" + `USER_ACCOUNT` + "`" + ` default value is ` + "`" + `true` + "`" + ` if ` + "`" + `role_ids` + "`" + ` contain ` + "`" + `System Admin` + "`" + ` role) ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `Profile ID (email or username).`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `Access key ID generated for ` + "`" + `SERVICE_ACCOUNT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Access key secret generated for ` + "`" + `SERVICE_ACCOUNT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_login_ts",
					Description: `(int) Last login timestamp.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "last_modified_ts",
					Description: `(int) Last modified timestamp.`,
				},
				resource.Attribute{
					Name:        "access_keys_count",
					Description: `(int) Access keys count.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of User Profile Roles Details. Each item has role information, as defined [below](#roles). ### Roles`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `User Role ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User Role Name.`,
				},
				resource.Attribute{
					Name:        "only_allow_ci_access",
					Description: `(bool) Allow only CI Access for Build and Deploy security roles.`,
				},
				resource.Attribute{
					Name:        "only_allow_compute_access",
					Description: `(bool) Allow only Compute Access for reduced system admin roles.`,
				},
				resource.Attribute{
					Name:        "only_allow_read_access",
					Description: `(bool) Allow only read access.`,
				},
				resource.Attribute{
					Name:        "role_type",
					Description: `User Role Type. ## Import Resources can be imported using the username/email: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_user_profile.example user@email.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `Profile ID (email or username).`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `Access key ID generated for ` + "`" + `SERVICE_ACCOUNT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Access key secret generated for ` + "`" + `SERVICE_ACCOUNT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_login_ts",
					Description: `(int) Last login timestamp.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "last_modified_ts",
					Description: `(int) Last modified timestamp.`,
				},
				resource.Attribute{
					Name:        "access_keys_count",
					Description: `(int) Access keys count.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of User Profile Roles Details. Each item has role information, as defined [below](#roles). ### Roles`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `User Role ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User Role Name.`,
				},
				resource.Attribute{
					Name:        "only_allow_ci_access",
					Description: `(bool) Allow only CI Access for Build and Deploy security roles.`,
				},
				resource.Attribute{
					Name:        "only_allow_compute_access",
					Description: `(bool) Allow only Compute Access for reduced system admin roles.`,
				},
				resource.Attribute{
					Name:        "only_allow_read_access",
					Description: `(bool) Allow only read access.`,
				},
				resource.Attribute{
					Name:        "role_type",
					Description: `User Role Type. ## Import Resources can be imported using the username/email: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_user_profile.example user@email.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_user_role",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description.`,
				},
				resource.Attribute{
					Name:        "role_type",
					Description: `(Required) User role type. Valid values are ` + "`" + `System Admin` + "`" + `, ` + "`" + `Account Group Admin` + "`" + `, ` + "`" + `Account Group Read Only` + "`" + `, ` + "`" + `Cloud Provisioning Admin` + "`" + `, ` + "`" + `Account and Cloud Provisioning Admin` + "`" + `, or ` + "`" + `Build and Deploy Security` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "account_group_ids",
					Description: `(Optional) List of accessible account group IDs. (Can't be set if ` + "`" + `role_type` + "`" + ` is ` + "`" + `System Admin` + "`" + ` or ` + "`" + `Build and Deploy Security` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "resource_list_ids",
					Description: `(Optional) List of resource list IDs.`,
				},
				resource.Attribute{
					Name:        "code_repository_ids",
					Description: `(Optional) List of code repository IDs.`,
				},
				resource.Attribute{
					Name:        "restrict_dismissal_access",
					Description: `(Optional, bool) Restrict dismissal access.`,
				},
				resource.Attribute{
					Name:        "additional_attributes",
					Description: `(Optional) An Additional attributes spec, as defined [below](#additional-attributes). ## Additional Attributes`,
				},
				resource.Attribute{
					Name:        "only_allow_ci_access",
					Description: `(Optional, bool) - Allows only CI Access`,
				},
				resource.Attribute{
					Name:        "only_allow_read_access",
					Description: `(Optional, bool) - Allow read only access (True for Account Group Read Only user role)`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `Role UUID.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "last_modified_ts",
					Description: `(int) Last modified timestamp.`,
				},
				resource.Attribute{
					Name:        "associated_users",
					Description: `List of associated application users which cannot exist in the system without the user role.`,
				},
				resource.Attribute{
					Name:        "account_groups",
					Description: `List of account groups, as defined [below](#account-groups). ### Account Groups Each account group has the following attributes.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Group name. ## Import Resources can be imported using the role ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_user_role.example 11111-22-33 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "role_id",
					Description: `Role UUID.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "last_modified_ts",
					Description: `(int) Last modified timestamp.`,
				},
				resource.Attribute{
					Name:        "associated_users",
					Description: `List of associated application users which cannot exist in the system without the user role.`,
				},
				resource.Attribute{
					Name:        "account_groups",
					Description: `List of account groups, as defined [below](#account-groups). ### Account Groups Each account group has the following attributes.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Group name. ## Import Resources can be imported using the role ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_user_role.example 11111-22-33 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"prismacloud_ExampleUsage":                            0,
		"prismacloud_account_group":                           1,
		"prismacloud_alert_rule":                              2,
		"prismacloud_cloud_account":                           3,
		"prismacloud_compliance_standard":                     4,
		"prismacloud_compliance_standard_requirement":         5,
		"prismacloud_compliance_standard_requirement_section": 6,
		"prismacloud_datapattern":                             7,
		"prismacloud_dataprofile":                             8,
		"prismacloud_enterprise_settings":                     9,
		"prismacloud_integration":                             10,
		"prismacloud_org_cloud_account":                       11,
		"prismacloud_policy":                                  12,
		"prismacloud_report":                                  13,
		"prismacloud_rql_search":                              14,
		"prismacloud_saved_search":                            15,
		"prismacloud_user_profile":                            16,
		"prismacloud_user_role":                               17,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
