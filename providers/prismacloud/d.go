package prismacloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_account_group",
			Category:         "Data Sources",
			ShortDescription: `Retrieves account group information.`,
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
					Name:        "account_ids",
					Description: `List of cloud account IDs.`,
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `List of cloud account IDs.`,
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
			Type:             "prismacloud_account_groups",
			Category:         "Data Sources",
			ShortDescription: `Lists account groups.`,
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
					Description: `Singly associated rules which cannot exist in the system without the account group spec, as defined [below](#alert-rules). ### Accounts Each account has the following attributes.`,
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
					Description: `Alert name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_alert_rule",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information on a specific alert rule.`,
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
			ShortDescription: `Retrieve a list of alert rules.`,
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
			ShortDescription: `Retrieve alert information.`,
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
			ShortDescription: `Retrieve information on a specific cloud account.`,
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
			ShortDescription: `Retrieve a list of cloud accounts onboarded onto the Prisma Cloud platform.`,
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
			ShortDescription: `Retrieve info for a compliance standard.`,
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
			ShortDescription: `Retrieve info on a compliance standard requirement.`,
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
			ShortDescription: `Retrieve information on a compliance standard requirement section.`,
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
			ShortDescription: `Retrieve a list of compliance standard requirement sections.`,
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
			ShortDescription: `Retrieve a list of compliance standard requirements.`,
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
			ShortDescription: `Retrieve a list of compliance standards.`,
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
			Type:             "prismacloud_enterprise_settings",
			Category:         "Data Sources",
			ShortDescription: `Retrieves enterprise settings information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "session_timeout",
					Description: `(int) Browser session timeout.`,
				},
				resource.Attribute{
					Name:        "anomaly_training_model_threshold",
					Description: `Anomaly training model threshold.`,
				},
				resource.Attribute{
					Name:        "anomaly_alert_disposition",
					Description: `Anomaly alert disposition.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_integration",
			Category:         "Data Sources",
			ShortDescription: `Retrieves integration information.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
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
					Name:        "integration_type",
					Description: `Integration type.`,
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
					Description: `Webhook headers, as defined [below](#headers).`,
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
					Description: `(bool) Read-only.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `Webhook headers, as defined [below](#headers).`,
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
					Description: `(bool) Read-only.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_integrations",
			Category:         "Data Sources",
			ShortDescription: `Retrieves an integration listing.`,
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
			Type:             "prismacloud_policies",
			Category:         "Data Sources",
			ShortDescription: `Retrieve a list of policies.`,
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
			ShortDescription: `Retrieve information on a specific policy.`,
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
					Description: `Compliance ID`,
				},
				resource.Attribute{
					Name:        "section_label",
					Description: `Section label`,
				},
				resource.Attribute{
					Name:        "custom_assigned",
					Description: `(bool) Custom assigned`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_type",
					Description: `Policy type`,
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
					Description: `Compliance ID`,
				},
				resource.Attribute{
					Name:        "section_label",
					Description: `Section label`,
				},
				resource.Attribute{
					Name:        "custom_assigned",
					Description: `(bool) Custom assigned`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_rql_historic_search",
			Category:         "Data Sources",
			ShortDescription: `Retrieve a specific historic RQL search.`,
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
			ShortDescription: `Retrieve a list of historic RQL searches.`,
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
	}

	dataSourcesMap = map[string]int{

		"prismacloud_account_group":                            0,
		"prismacloud_account_groups":                           1,
		"prismacloud_alert_rule":                               2,
		"prismacloud_alert_rules":                              3,
		"prismacloud_alerts":                                   4,
		"prismacloud_cloud_account":                            5,
		"prismacloud_cloud_accounts":                           6,
		"prismacloud_compliance_standard":                      7,
		"prismacloud_compliance_standard_requirement":          8,
		"prismacloud_compliance_standard_requirement_section":  9,
		"prismacloud_compliance_standard_requirement_sections": 10,
		"prismacloud_compliance_standard_requirements":         11,
		"prismacloud_compliance_standards":                     12,
		"prismacloud_enterprise_settings":                      13,
		"prismacloud_integration":                              14,
		"prismacloud_integrations":                             15,
		"prismacloud_policies":                                 16,
		"prismacloud_policy":                                   17,
		"prismacloud_rql_historic_search":                      18,
		"prismacloud_rql_historic_searches":                    19,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
