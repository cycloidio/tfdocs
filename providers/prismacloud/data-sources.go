package prismacloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_ExampleUsage",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_account_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the account group.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Account group ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
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
					Name:        "account_ids",
					Description: `List of cloud account IDs.`,
				},
				resource.Attribute{
					Name:        "child_group_ids",
					Description: `List of child account group IDs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
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
					Name:        "account_ids",
					Description: `List of cloud account IDs.`,
				},
				resource.Attribute{
					Name:        "child_group_ids",
					Description: `List of child account group IDs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_account_groups",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of account groups.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of accounts, as defined [below](#listing). ### Listing`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the account group.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Account group ID.`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `Associated cloud accounts spec, as defined [below](#accounts).`,
				},
				resource.Attribute{
					Name:        "alert_rules",
					Description: `Singly associated rules which cannot exist in the system without the account group spec, as defined [below](#alert-rules).`,
				},
				resource.Attribute{
					Name:        "parent_info",
					Description: `Parent account group info spec, as defined [below](#parent-info). ### Accounts Each account has the following attributes.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Associated cloud account ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Associated cloud account name.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Associated cloud account type. ### Alert Rules Each alert rule has the following attributes.`,
				},
				resource.Attribute{
					Name:        "alert_id",
					Description: `The alert ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Alert name. ### Parent Info Each parent info has the following attributes.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Parent account group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Parent account group name.`,
				},
				resource.Attribute{
					Name:        "auto_created",
					Description: `(bool) Boolean to indicate if account group is automatically created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_alert_rule",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_scan_config_id",
					Description: `Policy scan config ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Rule/Scan name ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Enabled`,
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
					Name:        "last_modified_on",
					Description: `(int) Last modified on`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
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
					Description: `(bool) Include open alerts in notification`,
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
					Description: `(bool) Read only`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Model for the target filter, as defined [below](#target)`,
				},
				resource.Attribute{
					Name:        "notification_config",
					Description: `List of data for notifications to third-party tools, as defined [below](#notification-config) ### Target`,
				},
				resource.Attribute{
					Name:        "account_groups",
					Description: `List of account groups`,
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
					Description: `List of TargetTag objects, as defined [below](#tags) ### Tags`,
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
					Description: `Frequency`,
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
					Name:        "last_updated",
					Description: `(int) Last updated`,
				},
				resource.Attribute{
					Name:        "last_sent_ts",
					Description: `(int) Time of last notification in miliseconds`,
				},
				resource.Attribute{
					Name:        "config_type",
					Description: `Config type`,
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
					Description: `(int) Offset`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Enabled`,
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
					Name:        "last_modified_on",
					Description: `(int) Last modified on`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
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
					Description: `(bool) Include open alerts in notification`,
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
					Description: `(bool) Read only`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Model for the target filter, as defined [below](#target)`,
				},
				resource.Attribute{
					Name:        "notification_config",
					Description: `List of data for notifications to third-party tools, as defined [below](#notification-config) ### Target`,
				},
				resource.Attribute{
					Name:        "account_groups",
					Description: `List of account groups`,
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
					Description: `List of TargetTag objects, as defined [below](#tags) ### Tags`,
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
					Description: `Frequency`,
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
					Name:        "last_updated",
					Description: `(int) Last updated`,
				},
				resource.Attribute{
					Name:        "last_sent_ts",
					Description: `(int) Time of last notification in miliseconds`,
				},
				resource.Attribute{
					Name:        "config_type",
					Description: `Config type`,
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
					Description: `(int) Offset`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_alert_rules",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of alert rules.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of alerts returned, as defined [below](#listing). ### Listing Each alert rule has the following attributes:`,
				},
				resource.Attribute{
					Name:        "policy_scan_config_id",
					Description: `Policy scan config ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Rule/Scan name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Rule/Scan is enabled or not`,
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
					Name:        "owner",
					Description: `Customer`,
				},
				resource.Attribute{
					Name:        "open_alerts_count",
					Description: `(int) Open alerts count`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(bool) Model is read-only`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(bool) Deleted`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_alerts",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "time_range",
					Description: `(Required) The time range spec, as defined [below](#time-range).`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional, int) Max number of alerts to return (default: ` + "`" + `10000` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "filters",
					Description: `(Optional) Filtering parameters spec, as defined [below](#filters).`,
				},
				resource.Attribute{
					Name:        "sort_by",
					Description: `(Optional) Array of sort properties. Append :asc or :desc to the key to sort by ascending or descending order respectively. ### Time Range The ` + "`" + `time_range` + "`" + ` block allows you to specify one of multiple supported time ranges. Only one time range can be specified.`,
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
					Description: `A to-now time range spec, as defined [below](#to-now-time-range). ### Absolute Time Range`,
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
					Description: `(Required) The time unit. Valid values are ` + "`" + `hour` + "`" + `, ` + "`" + `day` + "`" + `, ` + "`" + `week` + "`" + `, ` + "`" + `month` + "`" + `, or ` + "`" + `year` + "`" + `. ### To Now Time Range From some time in the past until now.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Required) The time unit. Valid values are ` + "`" + `login` + "`" + `, ` + "`" + `epoch` + "`" + `, ` + "`" + `day` + "`" + `, ` + "`" + `week` + "`" + `, ` + "`" + `month` + "`" + `, or ` + "`" + `year` + "`" + `. ### Filters Filtering parameters. This block can be specified multiple times to add more filters.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Param name to filter on.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) Operator between the name and value params (default: ` + "`" + `=` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Param value for the filter. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "page_token",
					Description: `The next page token returned.`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of alerts returned.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `Alert listing, as defined [below](#listing). ### Listing The alert information.`,
				},
				resource.Attribute{
					Name:        "alert_id",
					Description: `Alert ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Alert status.`,
				},
				resource.Attribute{
					Name:        "first_seen",
					Description: `(int) First seen.`,
				},
				resource.Attribute{
					Name:        "last_seen",
					Description: `(int) Last seen.`,
				},
				resource.Attribute{
					Name:        "alert_time",
					Description: `(int) Alert time.`,
				},
				resource.Attribute{
					Name:        "event_occurred",
					Description: `(int) Event occurred.`,
				},
				resource.Attribute{
					Name:        "triggered_by",
					Description: `Triggered by.`,
				},
				resource.Attribute{
					Name:        "alert_count",
					Description: `(int) Alert count.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "page_token",
					Description: `The next page token returned.`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of alerts returned.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `Alert listing, as defined [below](#listing). ### Listing The alert information.`,
				},
				resource.Attribute{
					Name:        "alert_id",
					Description: `Alert ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Alert status.`,
				},
				resource.Attribute{
					Name:        "first_seen",
					Description: `(int) First seen.`,
				},
				resource.Attribute{
					Name:        "last_seen",
					Description: `(int) Last seen.`,
				},
				resource.Attribute{
					Name:        "alert_time",
					Description: `(int) Alert time.`,
				},
				resource.Attribute{
					Name:        "event_occurred",
					Description: `(int) Event occurred.`,
				},
				resource.Attribute{
					Name:        "triggered_by",
					Description: `Triggered by.`,
				},
				resource.Attribute{
					Name:        "alert_count",
					Description: `(int) Alert count.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_cloud_account",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) The cloud type. Valid values are ` + "`" + `aws` + "`" + `, ` + "`" + `azure` + "`" + `, ` + "`" + `gcp` + "`" + `, or ` + "`" + `alibaba_cloud` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, computed) Cloud account name; computed if this is not supplied.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional, computed) Account ID; computed if this is not supplied. ## Attribute Reference The cloud type given above determines which of the attributes are populated:`,
				},
				resource.Attribute{
					Name:        "disable_on_destroy",
					Description: `(bool) To disable cloud account instead of deleting when calling Terraform destroy.`,
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
					Description: `AWS account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `AWS account external ID.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `Unique identifier for an AWS resource (ARN). ### Azure`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Azure account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `Application ID registered with Active Directory.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Application ID key.`,
				},
				resource.Attribute{
					Name:        "monitor_flow_logs",
					Description: `(bool) Automatically ingest flow logs.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Active Directory ID associated with Azure.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `Unique ID of the service principal object associated with the Prisma Cloud application that you create. ### GCP`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `GCP account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform`,
				},
				resource.Attribute{
					Name:        "compression_enabled",
					Description: `(bool) Enable flow log compression.`,
				},
				resource.Attribute{
					Name:        "data_flow_enabled_project",
					Description: `GCP project for flow log compression.`,
				},
				resource.Attribute{
					Name:        "flow_log_storage_bucket",
					Description: `GCP Flow logs storage bucket.`,
				},
				resource.Attribute{
					Name:        "credentials_json",
					Description: `Content of the JSON credentials file. ### Alibaba Cloud`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Alibaba account ID.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform.`,
				},
				resource.Attribute{
					Name:        "ram_arn",
					Description: `Unique identifier for an Alibaba RAM role resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "disable_on_destroy",
					Description: `(bool) To disable cloud account instead of deleting when calling Terraform destroy.`,
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
					Description: `AWS account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `AWS account external ID.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `Unique identifier for an AWS resource (ARN). ### Azure`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Azure account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `Application ID registered with Active Directory.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Application ID key.`,
				},
				resource.Attribute{
					Name:        "monitor_flow_logs",
					Description: `(bool) Automatically ingest flow logs.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Active Directory ID associated with Azure.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `Unique ID of the service principal object associated with the Prisma Cloud application that you create. ### GCP`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `GCP account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform`,
				},
				resource.Attribute{
					Name:        "compression_enabled",
					Description: `(bool) Enable flow log compression.`,
				},
				resource.Attribute{
					Name:        "data_flow_enabled_project",
					Description: `GCP project for flow log compression.`,
				},
				resource.Attribute{
					Name:        "flow_log_storage_bucket",
					Description: `GCP Flow logs storage bucket.`,
				},
				resource.Attribute{
					Name:        "credentials_json",
					Description: `Content of the JSON credentials file. ### Alibaba Cloud`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Alibaba account ID.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform.`,
				},
				resource.Attribute{
					Name:        "ram_arn",
					Description: `Unique identifier for an Alibaba RAM role resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_cloud_accounts",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of cloud accounts.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of cloud accounts, defined [below](#listing). ### Listing Each cloud account has the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Account name`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud type`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Account ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_compliance_standard",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cs_id",
					Description: `Compliance standard ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Compliance standard name ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
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
					Description: `List of cloud types (determined based on policies assigned)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
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
					Description: `List of cloud types (determined based on policies assigned)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_compliance_standard_requirement",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "csr_id",
					Description: `Compliance standard requirement ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Compliance standard requirement name Additional arguments:`,
				},
				resource.Attribute{
					Name:        "cs_id",
					Description: `(Required) Compliance standard ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
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
					Name:        "requirement_id",
					Description: `Compliance requirement number`,
				},
				resource.Attribute{
					Name:        "view_order",
					Description: `(int) View order`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
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
					Name:        "requirement_id",
					Description: `Compliance requirement number`,
				},
				resource.Attribute{
					Name:        "view_order",
					Description: `(int) View order`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_compliance_standard_requirement_section",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "csrs_id",
					Description: `Compliance standard requirement section ID`,
				},
				resource.Attribute{
					Name:        "section_id",
					Description: `Compliance section ID Additional arguments:`,
				},
				resource.Attribute{
					Name:        "csr_id",
					Description: `(Required) Compliance standard ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
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
					Name:        "view_order",
					Description: `(int) View order`,
				},
				resource.Attribute{
					Name:        "associated_policy_ids",
					Description: `List of associated policy IDs`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
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
					Name:        "view_order",
					Description: `(int) View order`,
				},
				resource.Attribute{
					Name:        "associated_policy_ids",
					Description: `List of associated policy IDs`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_compliance_standard_requirement_sections",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "csr_id",
					Description: `(Required) Compliance standard ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of system supported and custom compliance standard requirement sections.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of compliance requirement sections, as defined [below](#listing) ### Listing Each requirement section has the following attributes:`,
				},
				resource.Attribute{
					Name:        "csrs_id",
					Description: `Compliance standard requirement section ID`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
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
					Name:        "section_id",
					Description: `Compliance section ID`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Section label`,
				},
				resource.Attribute{
					Name:        "view_order",
					Description: `(int) View order`,
				},
				resource.Attribute{
					Name:        "associated_policy_ids",
					Description: `List of associated policy IDs`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of system supported and custom compliance standard requirement sections.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of compliance requirement sections, as defined [below](#listing) ### Listing Each requirement section has the following attributes:`,
				},
				resource.Attribute{
					Name:        "csrs_id",
					Description: `Compliance standard requirement section ID`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
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
					Name:        "section_id",
					Description: `Compliance section ID`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `Section label`,
				},
				resource.Attribute{
					Name:        "view_order",
					Description: `(int) View order`,
				},
				resource.Attribute{
					Name:        "associated_policy_ids",
					Description: `List of associated policy IDs`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_compliance_standard_requirements",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cs_id",
					Description: `(Required) Compliance standard ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of system supported and custom compliance standard requirements`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of compliance requirements, as defined [below](#listing) ### Listing Each requirement has the following attributes:`,
				},
				resource.Attribute{
					Name:        "csr_id",
					Description: `Compliance standard requirement ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Compliance standard requirement name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
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
					Name:        "requirement_id",
					Description: `Compliance requirement number`,
				},
				resource.Attribute{
					Name:        "view_order",
					Description: `(int) View order`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of system supported and custom compliance standard requirements`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of compliance requirements, as defined [below](#listing) ### Listing Each requirement has the following attributes:`,
				},
				resource.Attribute{
					Name:        "csr_id",
					Description: `Compliance standard requirement ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Compliance standard requirement name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
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
					Name:        "requirement_id",
					Description: `Compliance requirement number`,
				},
				resource.Attribute{
					Name:        "view_order",
					Description: `(int) View order`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_compliance_standards",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Number of system supported and custom compliance standards`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of system supported and custom compliance standards, as defined [below](#listing) ### Listing Each standard has the following attributes:`,
				},
				resource.Attribute{
					Name:        "cs_id",
					Description: `Compliance standard ID`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
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
					Name:        "name",
					Description: `Compliance standard name`,
				},
				resource.Attribute{
					Name:        "cloud_types",
					Description: `List of cloud types (determined based on policies assigned)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_datapattern",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pattern_id",
					Description: `Pattern ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Pattern name. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Pattern description.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Pattern mode (predefined or custom).`,
				},
				resource.Attribute{
					Name:        "detection_technique",
					Description: `Detection technique.`,
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
					Name:        "proximity_keywords",
					Description: `List of proximity keywords.`,
				},
				resource.Attribute{
					Name:        "regexes",
					Description: `List of regexes, as defined [below](#regexes).`,
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
					Description: `(bool) Is editable. ### Regexes`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `Regular expression (match criteria for the data you want to find within your assets).`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(int) Weight to assign a score to a text entry (pattern match occurs when the score threshold is exceeded).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Pattern description.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Pattern mode (predefined or custom).`,
				},
				resource.Attribute{
					Name:        "detection_technique",
					Description: `Detection technique.`,
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
					Name:        "proximity_keywords",
					Description: `List of proximity keywords.`,
				},
				resource.Attribute{
					Name:        "regexes",
					Description: `List of regexes, as defined [below](#regexes).`,
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
					Description: `(bool) Is editable. ### Regexes`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `Regular expression (match criteria for the data you want to find within your assets).`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(int) Weight to assign a score to a text entry (pattern match occurs when the score threshold is exceeded).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_datapatterns",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of data patterns.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of data patterns returned, as defined [below](#listing). ### Listing Each data pattern has the following attributes:`,
				},
				resource.Attribute{
					Name:        "pattern_id",
					Description: `Pattern ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Pattern Name.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Pattern mode (predefined or custom).`,
				},
				resource.Attribute{
					Name:        "detection_technique",
					Description: `Detection technique.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `(int) Last updated at.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `Updated by.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by.`,
				},
				resource.Attribute{
					Name:        "is_editable",
					Description: `(bool) Is editable.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_dataprofile",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `Profile ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Profile Name. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Profile description.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `Type (basic or advance).`,
				},
				resource.Attribute{
					Name:        "profile_type",
					Description: `Profile type (custom or system).`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status (hidden or non_hidden).`,
				},
				resource.Attribute{
					Name:        "profile_status",
					Description: `Profile status (active or disabled).`,
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
					Description: `Updated at (unix time).`,
				},
				resource.Attribute{
					Name:        "data_patterns_rule_1",
					Description: `Model for DataProfile Rules, as defined [below](#data-patterns-rule-1). ### Data Patterns Rule 1`,
				},
				resource.Attribute{
					Name:        "operator_type",
					Description: `Pattern operator type.`,
				},
				resource.Attribute{
					Name:        "data_pattern_rules",
					Description: `List of DataPattern Rules. Each item has data-pattern information, as defined [below](#data-pattern-rules). #### Data Pattern Rules`,
				},
				resource.Attribute{
					Name:        "pattern_id",
					Description: `Pattern ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Pattern name.`,
				},
				resource.Attribute{
					Name:        "detection_technique",
					Description: `Detection technique.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `Match type.`,
				},
				resource.Attribute{
					Name:        "occurrence_operator_type",
					Description: `Occurrence operator type.`,
				},
				resource.Attribute{
					Name:        "occurrence_count",
					Description: `Occurrence count.`,
				},
				resource.Attribute{
					Name:        "confidence_level",
					Description: `Confidence level.`,
				},
				resource.Attribute{
					Name:        "supported_confidence_levels",
					Description: `List of supported confidence levels.`,
				},
				resource.Attribute{
					Name:        "occurrence_high",
					Description: `High occurrence value.`,
				},
				resource.Attribute{
					Name:        "occurrence_low",
					Description: `Low occurrence value.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Profile description.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `Type (basic or advance).`,
				},
				resource.Attribute{
					Name:        "profile_type",
					Description: `Profile type (custom or system).`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Tenant ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status (hidden or non_hidden).`,
				},
				resource.Attribute{
					Name:        "profile_status",
					Description: `Profile status (active or disabled).`,
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
					Description: `Updated at (unix time).`,
				},
				resource.Attribute{
					Name:        "data_patterns_rule_1",
					Description: `Model for DataProfile Rules, as defined [below](#data-patterns-rule-1). ### Data Patterns Rule 1`,
				},
				resource.Attribute{
					Name:        "operator_type",
					Description: `Pattern operator type.`,
				},
				resource.Attribute{
					Name:        "data_pattern_rules",
					Description: `List of DataPattern Rules. Each item has data-pattern information, as defined [below](#data-pattern-rules). #### Data Pattern Rules`,
				},
				resource.Attribute{
					Name:        "pattern_id",
					Description: `Pattern ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Pattern name.`,
				},
				resource.Attribute{
					Name:        "detection_technique",
					Description: `Detection technique.`,
				},
				resource.Attribute{
					Name:        "match_type",
					Description: `Match type.`,
				},
				resource.Attribute{
					Name:        "occurrence_operator_type",
					Description: `Occurrence operator type.`,
				},
				resource.Attribute{
					Name:        "occurrence_count",
					Description: `Occurrence count.`,
				},
				resource.Attribute{
					Name:        "confidence_level",
					Description: `Confidence level.`,
				},
				resource.Attribute{
					Name:        "supported_confidence_levels",
					Description: `List of supported confidence levels.`,
				},
				resource.Attribute{
					Name:        "occurrence_high",
					Description: `High occurrence value.`,
				},
				resource.Attribute{
					Name:        "occurrence_low",
					Description: `Low occurrence value.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_dataprofiles",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of data profiles.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of data profiles returned, as defined [below](#listing). ### Listing Each data profile has the following attributes:`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `Profile ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Profile Name.`,
				},
				resource.Attribute{
					Name:        "profile_status",
					Description: `Profile status (active or disabled).`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Updated at (unix time).`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `Updated by.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_enterprise_settings",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
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
					Description: `(bool) Alarms enabled. Alarms are Prisma Cloud Platform health notifications which are generated to notify users of system level issues/errors.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_integration",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "integration_type",
					Description: `(Required) Integration type. One of the following must be specified:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the integration.`,
				},
				resource.Attribute{
					Name:        "integration_id",
					Description: `Integration ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Enabled.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by.`,
				},
				resource.Attribute{
					Name:        "created_ts",
					Description: `(int) Created on.`,
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
					Name:        "status",
					Description: `Status.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `(bool) Valid.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `Model for the integration status details, as defined [below](#reason).`,
				},
				resource.Attribute{
					Name:        "integration_config",
					Description: `Integration configuration, the values depend on the integration type, as defined [below](#integration-config). ### Reason`,
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
					Description: `Internationalization key. ### Integration Config`,
				},
				resource.Attribute{
					Name:        "queue_url",
					Description: `The URL configured in the Azure Service Bus queue where Prisma cloud sends alerts.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Azure account ID with service principal to which the Azure Service Bus queue belongs.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Azure Shared Access Signature connection string.`,
				},
				resource.Attribute{
					Name:        "queue_url",
					Description: `The Queue URL you used when you configured Prisma Cloud in Amazon SQS.`,
				},
				resource.Attribute{
					Name:        "more_info",
					Description: `(bool) Whether specific IAM credentials are specified for SQS queue access.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `AWS access key belonging to AWS IAM credentials meant for SQS queue access.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `AWS secret key for the given access key.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `Role ARN associated with the IAM role on Prisma Cloud`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External ID associated with the IAM role on Prisma Cloud.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `Qualys Login Username.`,
				},
				resource.Attribute{
					Name:        "base_url",
					Description: `Qualys Security Operations Center server API URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Qualys Password.`,
				},
				resource.Attribute{
					Name:        "host_url",
					Description: `ServiceNow URL.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `ServiceNow Login Username.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `ServiceNow password for login.`,
				},
				resource.Attribute{
					Name:        "tables",
					Description: `(Map of bools) Key/value pairs that identify the ServiceNow module tables with which to integrate.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Webhook URL.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `Webhook headers, as defined [below](#headers).`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `PagerDuty integration key.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `Slack webhook URL starting with ` + "`" + `https://hooks.slack.com/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_token",
					Description: `Splunk authentication token for the event collector.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Splunk HTTP event collector URL.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `Splunk source type.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Webhook URL.`,
				},
				resource.Attribute{
					Name:        "host_url",
					Description: `The Cortex XSOAR instance FQDN/IP — either the name or the IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `The consumer key you configured when you created the Prisma Cloud application access in your Cortex XSOAR environment.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Cortex release version.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Secret key from Tenable.io.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `Access key from Tenable.io.`,
				},
				resource.Attribute{
					Name:        "source_id",
					Description: `GCP source ID for the service account you used to onboard your GCP organization to Prisma Cloud.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `GCP organization ID.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Okta domain name.`,
				},
				resource.Attribute{
					Name:        "api_token",
					Description: `The authentication API token for Okta.`,
				},
				resource.Attribute{
					Name:        "s3_uri",
					Description: `Amazon S3 bucket URI.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `AWS region where the S3 bucket resides.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `Role ARN associated with the IAM role on Prisma Cloud.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External ID associated with the IAM role on Prisma Cloud.`,
				},
				resource.Attribute{
					Name:        "roll_up_interval",
					Description: `(int) Time in minutes at which batching of Prisma Cloud alerts would roll up.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `AWS account ID to which you assigned AWS Security Hub read-only access.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `List of AWS regions, as defined [below](#regions).`,
				},
				resource.Attribute{
					Name:        "host_url",
					Description: `Snowflake Account URL.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Snowflake Username.`,
				},
				resource.Attribute{
					Name:        "staging_integration_id",
					Description: `Existing Amazon S3 integration ID.`,
				},
				resource.Attribute{
					Name:        "pipe_name",
					Description: `Snowpipe Name.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `Private Key.`,
				},
				resource.Attribute{
					Name:        "pass_phrase",
					Description: `PassPhrase for private key.`,
				},
				resource.Attribute{
					Name:        "roll_up_interval",
					Description: `(int) Time in minutes at which batching of Prisma Cloud alerts would roll up. #### Headers`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Header name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Header value.`,
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
					Description: `AWS region name.`,
				},
				resource.Attribute{
					Name:        "api_identifier",
					Description: `AWS region code.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud Type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Enabled.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by.`,
				},
				resource.Attribute{
					Name:        "created_ts",
					Description: `(int) Created on.`,
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
					Name:        "status",
					Description: `Status.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `(bool) Valid.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `Model for the integration status details, as defined [below](#reason).`,
				},
				resource.Attribute{
					Name:        "integration_config",
					Description: `Integration configuration, the values depend on the integration type, as defined [below](#integration-config). ### Reason`,
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
					Description: `Internationalization key. ### Integration Config`,
				},
				resource.Attribute{
					Name:        "queue_url",
					Description: `The URL configured in the Azure Service Bus queue where Prisma cloud sends alerts.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Azure account ID with service principal to which the Azure Service Bus queue belongs.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Azure Shared Access Signature connection string.`,
				},
				resource.Attribute{
					Name:        "queue_url",
					Description: `The Queue URL you used when you configured Prisma Cloud in Amazon SQS.`,
				},
				resource.Attribute{
					Name:        "more_info",
					Description: `(bool) Whether specific IAM credentials are specified for SQS queue access.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `AWS access key belonging to AWS IAM credentials meant for SQS queue access.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `AWS secret key for the given access key.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `Role ARN associated with the IAM role on Prisma Cloud`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External ID associated with the IAM role on Prisma Cloud.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `Qualys Login Username.`,
				},
				resource.Attribute{
					Name:        "base_url",
					Description: `Qualys Security Operations Center server API URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Qualys Password.`,
				},
				resource.Attribute{
					Name:        "host_url",
					Description: `ServiceNow URL.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `ServiceNow Login Username.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `ServiceNow password for login.`,
				},
				resource.Attribute{
					Name:        "tables",
					Description: `(Map of bools) Key/value pairs that identify the ServiceNow module tables with which to integrate.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Webhook URL.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `Webhook headers, as defined [below](#headers).`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `PagerDuty integration key.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `Slack webhook URL starting with ` + "`" + `https://hooks.slack.com/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_token",
					Description: `Splunk authentication token for the event collector.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Splunk HTTP event collector URL.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `Splunk source type.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Webhook URL.`,
				},
				resource.Attribute{
					Name:        "host_url",
					Description: `The Cortex XSOAR instance FQDN/IP — either the name or the IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `The consumer key you configured when you created the Prisma Cloud application access in your Cortex XSOAR environment.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Cortex release version.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Secret key from Tenable.io.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `Access key from Tenable.io.`,
				},
				resource.Attribute{
					Name:        "source_id",
					Description: `GCP source ID for the service account you used to onboard your GCP organization to Prisma Cloud.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `GCP organization ID.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Okta domain name.`,
				},
				resource.Attribute{
					Name:        "api_token",
					Description: `The authentication API token for Okta.`,
				},
				resource.Attribute{
					Name:        "s3_uri",
					Description: `Amazon S3 bucket URI.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `AWS region where the S3 bucket resides.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `Role ARN associated with the IAM role on Prisma Cloud.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External ID associated with the IAM role on Prisma Cloud.`,
				},
				resource.Attribute{
					Name:        "roll_up_interval",
					Description: `(int) Time in minutes at which batching of Prisma Cloud alerts would roll up.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `AWS account ID to which you assigned AWS Security Hub read-only access.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `List of AWS regions, as defined [below](#regions).`,
				},
				resource.Attribute{
					Name:        "host_url",
					Description: `Snowflake Account URL.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Snowflake Username.`,
				},
				resource.Attribute{
					Name:        "staging_integration_id",
					Description: `Existing Amazon S3 integration ID.`,
				},
				resource.Attribute{
					Name:        "pipe_name",
					Description: `Snowpipe Name.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `Private Key.`,
				},
				resource.Attribute{
					Name:        "pass_phrase",
					Description: `PassPhrase for private key.`,
				},
				resource.Attribute{
					Name:        "roll_up_interval",
					Description: `(int) Time in minutes at which batching of Prisma Cloud alerts would roll up. #### Headers`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Header name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Header value.`,
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
					Description: `AWS region name.`,
				},
				resource.Attribute{
					Name:        "api_identifier",
					Description: `AWS region code.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud Type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_integrations",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of all integrations.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of integrations, as defined [below](#listing). ### Listing`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the integration.`,
				},
				resource.Attribute{
					Name:        "integration_id",
					Description: `Integration ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "integration_type",
					Description: `Integration type.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Enabled.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `(bool) Valid.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_org_cloud_account",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) The cloud type. Valid values are ` + "`" + `aws` + "`" + `, ` + "`" + `azure` + "`" + `, ` + "`" + `gcp` + "`" + `, or ` + "`" + `oci` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, computed) Cloud account name; computed if this is not supplied.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional, computed) Account ID; computed if this is not supplied. ## Attribute Reference The cloud type given above determines which of the attributes are populated:`,
				},
				resource.Attribute{
					Name:        "disable_on_destroy",
					Description: `(bool) To disable cloud account instead of deleting when calling Terraform destroy (default: ` + "`" + `false` + "`" + `).`,
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
					Description: `Oci account type spec, defined [below](#OCI). ### AWS`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `AWS account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `AWS account external ID.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `Unique identifier for an AWS resource (ARN).`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Defaults to "account" if not specified.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Defaults to "MONITOR". ### Azure`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Azure org account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `Application ID registered with Active Directory.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Application ID key.`,
				},
				resource.Attribute{
					Name:        "monitor_flow_logs",
					Description: `(bool) Automatically ingest flow logs.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Active Directory ID associated with Azure.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `Unique ID of the service principal object associated with the Prisma Cloud application that you create.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Defaults to "tenant" if not specified.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Defaults to "MONITOR". ### GCP`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `GCP org project ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform.`,
				},
				resource.Attribute{
					Name:        "compression_enabled",
					Description: `(bool) Enable flow log compression.`,
				},
				resource.Attribute{
					Name:        "data_flow_enabled_project",
					Description: `GCP project for flow log compression.`,
				},
				resource.Attribute{
					Name:        "flow_log_storage_bucket",
					Description: `GCP Flow logs storage bucket.`,
				},
				resource.Attribute{
					Name:        "credentials_json",
					Description: `Content of the JSON credentials file.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Account type - organization, or account.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection Mode - Monitor, or Monitor and Protect.`,
				},
				resource.Attribute{
					Name:        "organization_name",
					Description: `GCP org organization name.`,
				},
				resource.Attribute{
					Name:        "account_group_creation_mode",
					Description: `Cloud account group creation mode - MANUAL, AUTO, or RECURSIVE.`,
				},
				resource.Attribute{
					Name:        "hierarchy_selection",
					Description: `List of hierarchy selection. Each item has resource ID, display name, node type and selection type, as defined [below](#hierarchy-selection). #### Hierarchy Selection`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `Resource ID. For folders, format is folders/{folder ID}. For projects, format is {project number}. For orgs, format is organizations/{org ID}.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for folder, project, or organization.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `Node type - FOLDER, PROJECT, or ORG.`,
				},
				resource.Attribute{
					Name:        "selection_type",
					Description: `Selection type - INCLUDE, EXCLUDE, or ALL. ### OCI`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Oci account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) OCI identity group name that you define. Can be an existing group.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `account ID to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `OCI identity group name that you define. Can be an existing group`,
				},
				resource.Attribute{
					Name:        "ram_arn",
					Description: `Unique identifier for an Alibaba RAM role resource.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Account type - account or tenant`,
				},
				resource.Attribute{
					Name:        "default_account_group_id",
					Description: `(Required) account ID to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "home_region",
					Description: `OCI tenancy home region`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `OCI identity policy name that you define. Can be an existing policy that has the right policy statements`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `OCI identity user name that you define. Can be an existing user that has the right privileges`,
				},
				resource.Attribute{
					Name:        "user_ocid",
					Description: `OCI identity user Ocid that you define. Can be an existing user that has the right privileges`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "disable_on_destroy",
					Description: `(bool) To disable cloud account instead of deleting when calling Terraform destroy (default: ` + "`" + `false` + "`" + `).`,
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
					Description: `Oci account type spec, defined [below](#OCI). ### AWS`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `AWS account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `AWS account external ID.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `Unique identifier for an AWS resource (ARN).`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Defaults to "account" if not specified.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Defaults to "MONITOR". ### Azure`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Azure org account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `Application ID registered with Active Directory.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Application ID key.`,
				},
				resource.Attribute{
					Name:        "monitor_flow_logs",
					Description: `(bool) Automatically ingest flow logs.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Active Directory ID associated with Azure.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `Unique ID of the service principal object associated with the Prisma Cloud application that you create.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Defaults to "tenant" if not specified.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Defaults to "MONITOR". ### GCP`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `GCP org project ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform.`,
				},
				resource.Attribute{
					Name:        "compression_enabled",
					Description: `(bool) Enable flow log compression.`,
				},
				resource.Attribute{
					Name:        "data_flow_enabled_project",
					Description: `GCP project for flow log compression.`,
				},
				resource.Attribute{
					Name:        "flow_log_storage_bucket",
					Description: `GCP Flow logs storage bucket.`,
				},
				resource.Attribute{
					Name:        "credentials_json",
					Description: `Content of the JSON credentials file.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Account type - organization, or account.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection Mode - Monitor, or Monitor and Protect.`,
				},
				resource.Attribute{
					Name:        "organization_name",
					Description: `GCP org organization name.`,
				},
				resource.Attribute{
					Name:        "account_group_creation_mode",
					Description: `Cloud account group creation mode - MANUAL, AUTO, or RECURSIVE.`,
				},
				resource.Attribute{
					Name:        "hierarchy_selection",
					Description: `List of hierarchy selection. Each item has resource ID, display name, node type and selection type, as defined [below](#hierarchy-selection). #### Hierarchy Selection`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `Resource ID. For folders, format is folders/{folder ID}. For projects, format is {project number}. For orgs, format is organizations/{org ID}.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for folder, project, or organization.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `Node type - FOLDER, PROJECT, or ORG.`,
				},
				resource.Attribute{
					Name:        "selection_type",
					Description: `Selection type - INCLUDE, EXCLUDE, or ALL. ### OCI`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Oci account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether or not the account is enabled.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) OCI identity group name that you define. Can be an existing group.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `account ID to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `OCI identity group name that you define. Can be an existing group`,
				},
				resource.Attribute{
					Name:        "ram_arn",
					Description: `Unique identifier for an Alibaba RAM role resource.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Account type - account or tenant`,
				},
				resource.Attribute{
					Name:        "default_account_group_id",
					Description: `(Required) account ID to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "home_region",
					Description: `OCI tenancy home region`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `OCI identity policy name that you define. Can be an existing policy that has the right policy statements`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `OCI identity user name that you define. Can be an existing user that has the right privileges`,
				},
				resource.Attribute{
					Name:        "user_ocid",
					Description: `OCI identity user Ocid that you define. Can be an existing user that has the right privileges`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_org_cloud_accounts",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of cloud accounts.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of cloud accounts, defined [below](#listing). ### Listing Each cloud account has the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Account name`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud type`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Account ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_policies",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filters",
					Description: `(Optional, map) Filters to limit policies returned. Filter options can be found [here](https://api.docs.prismacloud.io/reference#get-policies-v2). ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of policies.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of policies returned, as defined [below](#listing). ### Listing Each policy has the following attributes:`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Policy ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Policy name`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Policy type`,
				},
				resource.Attribute{
					Name:        "system_default",
					Description: `(bool) If the policy is a system default for Prisma Cloud`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Severity`,
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
					Name:        "open_alerts_count",
					Description: `(int) Open alerts count`,
				},
				resource.Attribute{
					Name:        "policy_mode",
					Description: `Policy mode`,
				},
				resource.Attribute{
					Name:        "remediable",
					Description: `(bool) Remediable`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of policies.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of policies returned, as defined [below](#listing). ### Listing Each policy has the following attributes:`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Policy ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Policy name`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Policy type`,
				},
				resource.Attribute{
					Name:        "system_default",
					Description: `(bool) If the policy is a system default for Prisma Cloud`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Severity`,
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
					Name:        "open_alerts_count",
					Description: `(int) Open alerts count`,
				},
				resource.Attribute{
					Name:        "policy_mode",
					Description: `Policy mode`,
				},
				resource.Attribute{
					Name:        "remediable",
					Description: `(bool) Remediable`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_policy",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `Policy ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Policy name ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Policy type`,
				},
				resource.Attribute{
					Name:        "policy_subtypes",
					Description: `Policy subtypes`,
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
					Description: `Severity`,
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
					Name:        "remediable",
					Description: `(bool) Is remediable or not`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `Model for the rule, as defined [below](#rule)`,
				},
				resource.Attribute{
					Name:        "remediation",
					Description: `Model for remediation, as defined [below](#remediation)`,
				},
				resource.Attribute{
					Name:        "compliance_metadata",
					Description: `List of compliance data. Each item has compliance standard, requirement, and/or section information, as defined [below](#compliance-metadata) ### Rule`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name`,
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
					Description: `Saved search ID that defines the rule criteria`,
				},
				resource.Attribute{
					Name:        "data_criteria",
					Description: `Criteria for DLP Rule, as defined [below](#data-criteria)`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(map of strings) Parameters`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `Type of rule or RQL query ### Remediation`,
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
					Description: `Compliance Section UUID`,
				},
				resource.Attribute{
					Name:        "section_label",
					Description: `Section label`,
				},
				resource.Attribute{
					Name:        "custom_assigned",
					Description: `(bool) Custom assigned #### Data Criteria`,
				},
				resource.Attribute{
					Name:        "classification_result",
					Description: `Data Profile name required for DLP rule criteria`,
				},
				resource.Attribute{
					Name:        "exposure",
					Description: `File exposure`,
				},
				resource.Attribute{
					Name:        "extension",
					Description: `List of file extensions`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_type",
					Description: `Policy type`,
				},
				resource.Attribute{
					Name:        "policy_subtypes",
					Description: `Policy subtypes`,
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
					Description: `Severity`,
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
					Name:        "remediable",
					Description: `(bool) Is remediable or not`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `Model for the rule, as defined [below](#rule)`,
				},
				resource.Attribute{
					Name:        "remediation",
					Description: `Model for remediation, as defined [below](#remediation)`,
				},
				resource.Attribute{
					Name:        "compliance_metadata",
					Description: `List of compliance data. Each item has compliance standard, requirement, and/or section information, as defined [below](#compliance-metadata) ### Rule`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name`,
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
					Description: `Saved search ID that defines the rule criteria`,
				},
				resource.Attribute{
					Name:        "data_criteria",
					Description: `Criteria for DLP Rule, as defined [below](#data-criteria)`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(map of strings) Parameters`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `Type of rule or RQL query ### Remediation`,
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
					Description: `Compliance Section UUID`,
				},
				resource.Attribute{
					Name:        "section_label",
					Description: `Section label`,
				},
				resource.Attribute{
					Name:        "custom_assigned",
					Description: `(bool) Custom assigned #### Data Criteria`,
				},
				resource.Attribute{
					Name:        "classification_result",
					Description: `Data Profile name required for DLP rule criteria`,
				},
				resource.Attribute{
					Name:        "exposure",
					Description: `File exposure`,
				},
				resource.Attribute{
					Name:        "extension",
					Description: `List of file extensions`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_report",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "report_id",
					Description: `Report ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Report name ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "report_type",
					Description: `Report type`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud type`,
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
					Name:        "target",
					Description: `Model for report target, as defined [below](#target)`,
				},
				resource.Attribute{
					Name:        "counts",
					Description: `Model for compliance aggregate count, as defined [below](#counts). ### Target`,
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
					Description: `List of compliance standard IDs`,
				},
				resource.Attribute{
					Name:        "resource_groups",
					Description: `List of resource groups`,
				},
				resource.Attribute{
					Name:        "notify_to",
					Description: `List of email addresses to receive notification`,
				},
				resource.Attribute{
					Name:        "compression_enabled",
					Description: `(bool) Business unit detailed report compression enabled`,
				},
				resource.Attribute{
					Name:        "download_now",
					Description: `(bool) True = download now`,
				},
				resource.Attribute{
					Name:        "schedule_enabled",
					Description: `(bool) Report scheduling enabled`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `Recurring report schedule in RRULE format`,
				},
				resource.Attribute{
					Name:        "notification_template_id",
					Description: `Notification template id`,
				},
				resource.Attribute{
					Name:        "time_range",
					Description: `(Required) The time range spec, as defined [below](#time-range). ### Time Range`,
				},
				resource.Attribute{
					Name:        "absolute",
					Description: `An absolute time range spec, as defined [below](#absolute-time-range)`,
				},
				resource.Attribute{
					Name:        "relative",
					Description: `A relative time range spec, as defined [below](#relative-time-range)`,
				},
				resource.Attribute{
					Name:        "to_now",
					Description: `A "To Now" time range spec, as defined [below](#to-now-time-range) ### Absolute Time Range`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(int) Start time`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(int) End time ### Relative Time Range`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `(int) The time number`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `The time unit ### To Now Time Range`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `The time unit ### Counts`,
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
					Description: `(int) Total`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "report_type",
					Description: `Report type`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud type`,
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
					Name:        "target",
					Description: `Model for report target, as defined [below](#target)`,
				},
				resource.Attribute{
					Name:        "counts",
					Description: `Model for compliance aggregate count, as defined [below](#counts). ### Target`,
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
					Description: `List of compliance standard IDs`,
				},
				resource.Attribute{
					Name:        "resource_groups",
					Description: `List of resource groups`,
				},
				resource.Attribute{
					Name:        "notify_to",
					Description: `List of email addresses to receive notification`,
				},
				resource.Attribute{
					Name:        "compression_enabled",
					Description: `(bool) Business unit detailed report compression enabled`,
				},
				resource.Attribute{
					Name:        "download_now",
					Description: `(bool) True = download now`,
				},
				resource.Attribute{
					Name:        "schedule_enabled",
					Description: `(bool) Report scheduling enabled`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `Recurring report schedule in RRULE format`,
				},
				resource.Attribute{
					Name:        "notification_template_id",
					Description: `Notification template id`,
				},
				resource.Attribute{
					Name:        "time_range",
					Description: `(Required) The time range spec, as defined [below](#time-range). ### Time Range`,
				},
				resource.Attribute{
					Name:        "absolute",
					Description: `An absolute time range spec, as defined [below](#absolute-time-range)`,
				},
				resource.Attribute{
					Name:        "relative",
					Description: `A relative time range spec, as defined [below](#relative-time-range)`,
				},
				resource.Attribute{
					Name:        "to_now",
					Description: `A "To Now" time range spec, as defined [below](#to-now-time-range) ### Absolute Time Range`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(int) Start time`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(int) End time ### Relative Time Range`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `(int) The time number`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `The time unit ### To Now Time Range`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `The time unit ### Counts`,
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
					Description: `(int) Total`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_reports",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of reports.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of reports returned, as defined [below](#listing). ### Listing Each report has the following attributes:`,
				},
				resource.Attribute{
					Name:        "report_id",
					Description: `Report ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Report name`,
				},
				resource.Attribute{
					Name:        "report_type",
					Description: `Report type`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud type`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Report status`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_rql_historic_search",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "search_id",
					Description: `Historic RQL search ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Historic RQL search name ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "search_type",
					Description: `Search type`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud type`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `RQL query`,
				},
				resource.Attribute{
					Name:        "saved",
					Description: `(bool) If this is a saved search`,
				},
				resource.Attribute{
					Name:        "time_range",
					Description: `The RQL time range spec, as defined [below](#time-range) ### Time Range Only one of these will be defined:`,
				},
				resource.Attribute{
					Name:        "absolute",
					Description: `An absolute time range spec, as defined [below](#absolute-time-range)`,
				},
				resource.Attribute{
					Name:        "relative",
					Description: `A relative time range spec, as defined [below](#relative-time-range)`,
				},
				resource.Attribute{
					Name:        "to_now",
					Description: `A "To Now" time range spec, as defined [below](#to-now-time-range) ### Absolute Time Range`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(int) Start time`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(int) End time ### Relative Time Range`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `(int) The time number`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `The time unit ### To Now Time Range`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `The time unit`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "search_type",
					Description: `Search type`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud type`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `RQL query`,
				},
				resource.Attribute{
					Name:        "saved",
					Description: `(bool) If this is a saved search`,
				},
				resource.Attribute{
					Name:        "time_range",
					Description: `The RQL time range spec, as defined [below](#time-range) ### Time Range Only one of these will be defined:`,
				},
				resource.Attribute{
					Name:        "absolute",
					Description: `An absolute time range spec, as defined [below](#absolute-time-range)`,
				},
				resource.Attribute{
					Name:        "relative",
					Description: `A relative time range spec, as defined [below](#relative-time-range)`,
				},
				resource.Attribute{
					Name:        "to_now",
					Description: `A "To Now" time range spec, as defined [below](#to-now-time-range) ### Absolute Time Range`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(int) Start time`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(int) End time ### Relative Time Range`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `(int) The time number`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `The time unit ### To Now Time Range`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `The time unit`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_rql_historic_searches",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter for historic RQL searches. Valid values are ` + "`" + `saved` + "`" + ` (default) or ` + "`" + `recent` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional, int) Max number of historic RQL searches to return (default: ` + "`" + `1000` + "`" + `). ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of RQL historic searches.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of historic RQL searches, as defined [below](#listing). ### Listing Each result in the listing has the following attributes:`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "search_id",
					Description: `Historic RQL search ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name`,
				},
				resource.Attribute{
					Name:        "search_type",
					Description: `Search type`,
				},
				resource.Attribute{
					Name:        "saved",
					Description: `(bool) If this is a saved search`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of RQL historic searches.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of historic RQL searches, as defined [below](#listing). ### Listing Each result in the listing has the following attributes:`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by`,
				},
				resource.Attribute{
					Name:        "search_id",
					Description: `Historic RQL search ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name`,
				},
				resource.Attribute{
					Name:        "search_type",
					Description: `Search type`,
				},
				resource.Attribute{
					Name:        "saved",
					Description: `(bool) If this is a saved search`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_user_profile",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Required) Profile ID (email or username). ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Account Type (USER_ACCOUNT or SERVICE_ACCOUNT).`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `User email or service account name.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `First name.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `Last name.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email ID.`,
				},
				resource.Attribute{
					Name:        "access_keys_allowed",
					Description: `(bool) Access keys allowed.`,
				},
				resource.Attribute{
					Name:        "default_role_id",
					Description: `Default User Role ID.`,
				},
				resource.Attribute{
					Name:        "role_ids",
					Description: `List of Role IDs.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `Time zone (e.g. America/Los_Angeles).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Enabled.`,
				},
				resource.Attribute{
					Name:        "last_login_ts",
					Description: `(int) Last login timestamp.`,
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
					Name:        "access_keys_count",
					Description: `(int) Access key count.`,
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
					Description: `User Role Type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_type",
					Description: `Account Type (USER_ACCOUNT or SERVICE_ACCOUNT).`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `User email or service account name.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `First name.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `Last name.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email ID.`,
				},
				resource.Attribute{
					Name:        "access_keys_allowed",
					Description: `(bool) Access keys allowed.`,
				},
				resource.Attribute{
					Name:        "default_role_id",
					Description: `Default User Role ID.`,
				},
				resource.Attribute{
					Name:        "role_ids",
					Description: `List of Role IDs.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `Time zone (e.g. America/Los_Angeles).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Enabled.`,
				},
				resource.Attribute{
					Name:        "last_login_ts",
					Description: `(int) Last login timestamp.`,
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
					Name:        "access_keys_count",
					Description: `(int) Access key count.`,
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
					Description: `User Role Type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_user_profiles",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of user profiles.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of user profiles returned, as defined [below](#listing). ### Listing Each user profile has the following attributes:`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `Profile ID (email or username).`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Account Type (USER_ACCOUNT or SERVICE_ACCOUNT).`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `User email or service account name.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name.`,
				},
				resource.Attribute{
					Name:        "default_role_id",
					Description: `Default User Role ID.`,
				},
				resource.Attribute{
					Name:        "role_ids",
					Description: `List of Role IDs.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `Time zone (e.g. America/Los_Angeles).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Enabled.`,
				},
				resource.Attribute{
					Name:        "last_login_ts",
					Description: `(int) Last login timestamp.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_ts",
					Description: `(int) Last modified timestamp.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_user_role",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role_id",
					Description: `Role Id`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the role. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "role_type",
					Description: `User role type.`,
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
					Name:        "account_group_ids",
					Description: `List of accessible account group IDs.`,
				},
				resource.Attribute{
					Name:        "resource_list_ids",
					Description: `List of resource list IDs.`,
				},
				resource.Attribute{
					Name:        "code_repository_ids",
					Description: `List of code repository IDs.`,
				},
				resource.Attribute{
					Name:        "restrict_dismissal_access",
					Description: `(bool) Restrict dismissal access.`,
				},
				resource.Attribute{
					Name:        "associated_users",
					Description: `List of associated application users which cannot exist in the system without the user role.`,
				},
				resource.Attribute{
					Name:        "additional_attributes",
					Description: `An Additional attributes spec, as defined [below](#additional-attributes). ## Additional Attributes`,
				},
				resource.Attribute{
					Name:        "only_allow_ci_access",
					Description: `(bool) - Allows only CI Access.`,
				},
				resource.Attribute{
					Name:        "only_allow_read_access",
					Description: `(bool) - Allow read only access.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "role_type",
					Description: `User role type.`,
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
					Name:        "account_group_ids",
					Description: `List of accessible account group IDs.`,
				},
				resource.Attribute{
					Name:        "resource_list_ids",
					Description: `List of resource list IDs.`,
				},
				resource.Attribute{
					Name:        "code_repository_ids",
					Description: `List of code repository IDs.`,
				},
				resource.Attribute{
					Name:        "restrict_dismissal_access",
					Description: `(bool) Restrict dismissal access.`,
				},
				resource.Attribute{
					Name:        "associated_users",
					Description: `List of associated application users which cannot exist in the system without the user role.`,
				},
				resource.Attribute{
					Name:        "additional_attributes",
					Description: `An Additional attributes spec, as defined [below](#additional-attributes). ## Additional Attributes`,
				},
				resource.Attribute{
					Name:        "only_allow_ci_access",
					Description: `(bool) - Allows only CI Access.`,
				},
				resource.Attribute{
					Name:        "only_allow_read_access",
					Description: `(bool) - Allow read only access.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_user_roles",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of user roles.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of user roles returned, as defined [below](#listing). ### Listing Each user role has the following attributes:`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `Role Id`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the role.`,
				},
				resource.Attribute{
					Name:        "role_type",
					Description: `User role type.`,
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
					Name:        "associated_users",
					Description: `List of associated application users which cannot exist in the system without the user role.`,
				},
				resource.Attribute{
					Name:        "restrict_dismissal_access",
					Description: `(bool) Restrict dismissal access.`,
				},
				resource.Attribute{
					Name:        "account_groups",
					Description: `List of associated account groups, as defined [below](#account-groups).`,
				},
				resource.Attribute{
					Name:        "additional_attributes",
					Description: `An Additional attributes spec, as defined [below](#additional-attributes). #### Account Groups Each account group has the following attributes.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Group name. #### Additional Attributes`,
				},
				resource.Attribute{
					Name:        "only_allow_ci_access",
					Description: `(bool) - Allows only CI Access.`,
				},
				resource.Attribute{
					Name:        "only_allow_read_access",
					Description: `(bool) - Allow read only access.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"prismacloud_ExampleUsage":                             0,
		"prismacloud_account_group":                            1,
		"prismacloud_account_groups":                           2,
		"prismacloud_alert_rule":                               3,
		"prismacloud_alert_rules":                              4,
		"prismacloud_alerts":                                   5,
		"prismacloud_cloud_account":                            6,
		"prismacloud_cloud_accounts":                           7,
		"prismacloud_compliance_standard":                      8,
		"prismacloud_compliance_standard_requirement":          9,
		"prismacloud_compliance_standard_requirement_section":  10,
		"prismacloud_compliance_standard_requirement_sections": 11,
		"prismacloud_compliance_standard_requirements":         12,
		"prismacloud_compliance_standards":                     13,
		"prismacloud_datapattern":                              14,
		"prismacloud_datapatterns":                             15,
		"prismacloud_dataprofile":                              16,
		"prismacloud_dataprofiles":                             17,
		"prismacloud_enterprise_settings":                      18,
		"prismacloud_integration":                              19,
		"prismacloud_integrations":                             20,
		"prismacloud_org_cloud_account":                        21,
		"prismacloud_org_cloud_accounts":                       22,
		"prismacloud_policies":                                 23,
		"prismacloud_policy":                                   24,
		"prismacloud_report":                                   25,
		"prismacloud_reports":                                  26,
		"prismacloud_rql_historic_search":                      27,
		"prismacloud_rql_historic_searches":                    28,
		"prismacloud_user_profile":                             29,
		"prismacloud_user_profiles":                            30,
		"prismacloud_user_role":                                31,
		"prismacloud_user_roles":                               32,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
