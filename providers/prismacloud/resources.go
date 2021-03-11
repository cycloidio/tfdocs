package prismacloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

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
					Description: `(Optional) Description. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `List of cloud account IDs.`,
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
					Name:        "account_ids",
					Description: `List of cloud account IDs.`,
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
					Name:        "config_id",
					Description: `Alert rule notification config ID`,
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
					Description: `List of unique email addresses to notify`,
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
					Description: `Config type. Valid values are ` + "`" + `email` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `splunk` + "`" + `, ` + "`" + `amazon_sqs` + "`" + `, ` + "`" + `jira` + "`" + `, ` + "`" + `microsoft_teams` + "`" + `, ` + "`" + `webhook` + "`" + `, ` + "`" + `aws_security_hub` + "`" + `, ` + "`" + `google_cscc` + "`" + `, ` + "`" + `service_now` + "`" + `, ` + "`" + `pager_duty` + "`" + `, or ` + "`" + `demisto` + "`" + ``,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `Template ID`,
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
					Name:        "r_rule_schedule",
					Description: `R rule schedule`,
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
					Description: `(int) Offset ## Attribute Reference`,
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
					Name:        "last_updated",
					Description: `(int) Last updated`,
				},
				resource.Attribute{
					Name:        "last_sent_ts",
					Description: `(int) Time of last notification in miliseconds ## Import Resources can be imported using the policy scan config ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_alert_rule.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "last_updated",
					Description: `(int) Last updated`,
				},
				resource.Attribute{
					Name:        "last_sent_ts",
					Description: `(int) Time of last notification in miliseconds ## Import Resources can be imported using the policy scan config ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_alert_rule.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Alibaba account type spec, defined [below](#alibaba-cloud). ### AWS -> AWS Org support is not available yet`,
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
					Description: `(Optional) Defaults to "MONITOR" ### Azure`,
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
					Description: `(Optional) Defaults to "MONITOR" ### GCP ->GCP Org support is not available yet`,
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
					Description: `(Optional) Defaults to "MONITOR" ### Alibaba Cloud`,
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
					Name:        "anomaly_training_model_threshold",
					Description: `Anomaly training model threshold (` + "`" + `low` + "`" + `, ` + "`" + `medium` + "`" + `, or ` + "`" + `high` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "anomaly_alert_disposition",
					Description: `Anomaly alert disposition (` + "`" + `low` + "`" + `, ` + "`" + `medium` + "`" + `, or ` + "`" + `high` + "`" + `).`,
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
					Description: `(Required) Integration type (valid values can be found [here](https://api.docs.prismacloud.io/reference#integrations).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Enabled.`,
				},
				resource.Attribute{
					Name:        "integration_config",
					Description: `(Required) Integration configuration, the values depend on the integration type, as defined [below](#integration-config). ### Integration Config Refer to the [Prisma Cloud integration documentation](https://api.docs.prismacloud.io/reference#integration-configuration) if you need more information on a specific integration.`,
				},
				resource.Attribute{
					Name:        "queue_url",
					Description: `The Queue URL you used when you configured Prisma Cloud in Amazon SQS`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `(Qualys/ServiceNow) Login.`,
				},
				resource.Attribute{
					Name:        "base_url",
					Description: `Qualys Security Operations Center server API URL (without "http(s)")`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Qualys/ServiceNow) Password`,
				},
				resource.Attribute{
					Name:        "host_url",
					Description: `ServiceNow URL.`,
				},
				resource.Attribute{
					Name:        "tables",
					Description: `(Map of bools) Key/value pairs that identify the ServiceNow module tables with which to integrate (e.g. - incident, sn_si_incident, or em_event).`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `ServiceNow release version.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Webhook URL.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `List of webhook headers, as defined [below](#headers).`,
				},
				resource.Attribute{
					Name:        "auth_token",
					Description: `PagerDuty authentication token for the event collector.`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `PagerDuty integration key. ### Headers`,
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
					Description: `(bool) Read-only. ## Attribute Reference`,
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
					Description: `Internationalization key.`,
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
					Description: `Internationalization key.`,
				},
			},
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
					Description: `(Required) Policy type. Valid values are ` + "`" + `config` + "`" + `, ` + "`" + `audit_event` + "`" + `, or ` + "`" + `network` + "`" + ``,
				},
				resource.Attribute{
					Name:        "system_default",
					Description: `(bool) If policy is a system default policy or not`,
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
					Description: `Cloud type`,
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
					Name:        "remediable",
					Description: `(bool) Is remediable or not`,
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
					Description: `(Required) Saved search ID that defines the rule criteria`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Required, map of strings) Parameters. Valid keys are ` + "`" + `withIac` + "`" + ` and ` + "`" + `savedSearch` + "`" + ` and value is ` + "`" + `"true"` + "`" + ` or ` + "`" + `"false"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `(Required) Type of rule or RQL query. Valid values are ` + "`" + `Config` + "`" + `, ` + "`" + `AuditEvent` + "`" + `, or ` + "`" + `Network` + "`" + ` ### Remediation This section may be present or may be ommitted:`,
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
					Name:        "standard_name",
					Description: `Compliance standard name`,
				},
				resource.Attribute{
					Name:        "standard_description",
					Description: `Compliance standard description`,
				},
				resource.Attribute{
					Name:        "requirement_id",
					Description: `Requirement ID`,
				},
				resource.Attribute{
					Name:        "requirement_name",
					Description: `Requirement name`,
				},
				resource.Attribute{
					Name:        "requirement_description",
					Description: `Requirement description`,
				},
				resource.Attribute{
					Name:        "section_id",
					Description: `Section ID`,
				},
				resource.Attribute{
					Name:        "section_description",
					Description: `Section description`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Policy ID`,
				},
				resource.Attribute{
					Name:        "compliance_id",
					Description: `Compliance ID`,
				},
				resource.Attribute{
					Name:        "section_label",
					Description: `Section label`,
				},
				resource.Attribute{
					Name:        "custom_assigned",
					Description: `(bool) Custom assigned ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Policy ID`,
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
					Description: `Policy mode ## Import Resources can be imported using the poilcy ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_policy.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `Policy ID`,
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
					Description: `Policy mode ## Import Resources can be imported using the poilcy ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import prismacloud_policy.example 11111111-2222-3333-4444-555555555555 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) The search type. Valid values are ` + "`" + `config` + "`" + ` (default) and ` + "`" + `event` + "`" + `.`,
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
					Description: `(For ` + "`" + `search_type="event"` + "`" + `, list) List of event_data specs, as defined below. ` + "`" + `config_data` + "`" + ` supports the following attributes:`,
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
					Description: `Region API identifier.`,
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
					Description: `(For ` + "`" + `search_type="event"` + "`" + `, list) List of event_data specs, as defined below. ` + "`" + `config_data` + "`" + ` supports the following attributes:`,
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
					Description: `Region API identifier.`,
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
					Description: `(bool) This is set to ` + "`" + `true` + "`" + ` when the saved search is created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "saved",
					Description: `(bool) This is set to ` + "`" + `true` + "`" + ` when the saved search is created.`,
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
					Description: `(Required) User role type. Valid valies are ` + "`" + `System Admin` + "`" + `, ` + "`" + `Account Group Admin` + "`" + `, ` + "`" + `Account Group Read Only` + "`" + `, ` + "`" + `Cloud Provisioning Admin` + "`" + `, or ` + "`" + `Account and Cloud Provisioning Admin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "account_group_ids",
					Description: `(Optional) List of accessible account group IDs.`,
				},
				resource.Attribute{
					Name:        "associated_users",
					Description: `(Optional) List of associated application users which cannot exist in the system without the user role.`,
				},
				resource.Attribute{
					Name:        "restrict_dismissal_access",
					Description: `(Optional, bool) Restrict dismissal access. ## Attribute Reference`,
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

		"prismacloud_account_group":                           0,
		"prismacloud_alert_rule":                              1,
		"prismacloud_cloud_account":                           2,
		"prismacloud_compliance_standard":                     3,
		"prismacloud_compliance_standard_requirement":         4,
		"prismacloud_compliance_standard_requirement_section": 5,
		"prismacloud_enterprise_settings":                     6,
		"prismacloud_integration":                             7,
		"prismacloud_policy":                                  8,
		"prismacloud_rql_search":                              9,
		"prismacloud_saved_search":                            10,
		"prismacloud_user_role":                               11,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
