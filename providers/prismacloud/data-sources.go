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
					Description: `List of TargetTag objects, as defined [below](#tags)`,
				},
				resource.Attribute{
					Name:        "resource_list",
					Description: `Model for holding the resource list for compute access groups [below](#compute-access-group-ids) ### Tags`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Resource tag target`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `List of values for resource tag key ### Compute Access Group IDs`,
				},
				resource.Attribute{
					Name:        "compute_access_group_ids",
					Description: `List of compute access group IDs. ### Notification Config`,
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
					Description: `List of TargetTag objects, as defined [below](#tags)`,
				},
				resource.Attribute{
					Name:        "resource_list",
					Description: `Model for holding the resource list for compute access groups [below](#compute-access-group-ids) ### Tags`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Resource tag target`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `List of values for resource tag key ### Compute Access Group IDs`,
				},
				resource.Attribute{
					Name:        "compute_access_group_ids",
					Description: `List of compute access group IDs. ### Notification Config`,
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
			Type:             "prismacloud_anomaly_setting",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) Policy ID ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "alert_disposition",
					Description: `Alert disposition`,
				},
				resource.Attribute{
					Name:        "alert_disposition_description",
					Description: `Alert disposition information [below](#alert-disposition-description)`,
				},
				resource.Attribute{
					Name:        "policy_description",
					Description: `Policy description`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Policy name`,
				},
				resource.Attribute{
					Name:        "training_model_description",
					Description: `Training model information [below](#training-model-description)`,
				},
				resource.Attribute{
					Name:        "training_model_threshold",
					Description: `Training model threshold information ### Alert Disposition Description`,
				},
				resource.Attribute{
					Name:        "aggressive",
					Description: `Aggressive`,
				},
				resource.Attribute{
					Name:        "moderate",
					Description: `Moderate`,
				},
				resource.Attribute{
					Name:        "conservative",
					Description: `Conservative ### Training Model Description`,
				},
				resource.Attribute{
					Name:        "low",
					Description: `Low`,
				},
				resource.Attribute{
					Name:        "medium",
					Description: `Medium`,
				},
				resource.Attribute{
					Name:        "high",
					Description: `High`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "alert_disposition",
					Description: `Alert disposition`,
				},
				resource.Attribute{
					Name:        "alert_disposition_description",
					Description: `Alert disposition information [below](#alert-disposition-description)`,
				},
				resource.Attribute{
					Name:        "policy_description",
					Description: `Policy description`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Policy name`,
				},
				resource.Attribute{
					Name:        "training_model_description",
					Description: `Training model information [below](#training-model-description)`,
				},
				resource.Attribute{
					Name:        "training_model_threshold",
					Description: `Training model threshold information ### Alert Disposition Description`,
				},
				resource.Attribute{
					Name:        "aggressive",
					Description: `Aggressive`,
				},
				resource.Attribute{
					Name:        "moderate",
					Description: `Moderate`,
				},
				resource.Attribute{
					Name:        "conservative",
					Description: `Conservative ### Training Model Description`,
				},
				resource.Attribute{
					Name:        "low",
					Description: `Low`,
				},
				resource.Attribute{
					Name:        "medium",
					Description: `Medium`,
				},
				resource.Attribute{
					Name:        "high",
					Description: `High`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_anomaly_settings",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type. Valid values are ` + "`" + `Network` + "`" + `, ` + "`" + `UEBA` + "`" + ` or ` + "`" + `DNS` + "`" + `. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of anomaly settings.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of anomaly settings, as defined [below](#listing). ## Listing`,
				},
				resource.Attribute{
					Name:        "alert_disposition",
					Description: `Alert disposition`,
				},
				resource.Attribute{
					Name:        "alert_disposition_description",
					Description: `Alert disposition information [below](#alert-disposition-description)`,
				},
				resource.Attribute{
					Name:        "policy_description",
					Description: `Policy description`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Policy ID`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Policy name`,
				},
				resource.Attribute{
					Name:        "training_model_description",
					Description: `Training model information [below](#training-model-description)`,
				},
				resource.Attribute{
					Name:        "training_model_threshold",
					Description: `Training model threshold information ### Alert Disposition Description`,
				},
				resource.Attribute{
					Name:        "aggressive",
					Description: `Aggressive`,
				},
				resource.Attribute{
					Name:        "moderate",
					Description: `Moderate`,
				},
				resource.Attribute{
					Name:        "conservative",
					Description: `Conservative ### Training Model Description`,
				},
				resource.Attribute{
					Name:        "low",
					Description: `Low`,
				},
				resource.Attribute{
					Name:        "medium",
					Description: `Medium`,
				},
				resource.Attribute{
					Name:        "high",
					Description: `High`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of anomaly settings.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of anomaly settings, as defined [below](#listing). ## Listing`,
				},
				resource.Attribute{
					Name:        "alert_disposition",
					Description: `Alert disposition`,
				},
				resource.Attribute{
					Name:        "alert_disposition_description",
					Description: `Alert disposition information [below](#alert-disposition-description)`,
				},
				resource.Attribute{
					Name:        "policy_description",
					Description: `Policy description`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Policy ID`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Policy name`,
				},
				resource.Attribute{
					Name:        "training_model_description",
					Description: `Training model information [below](#training-model-description)`,
				},
				resource.Attribute{
					Name:        "training_model_threshold",
					Description: `Training model threshold information ### Alert Disposition Description`,
				},
				resource.Attribute{
					Name:        "aggressive",
					Description: `Aggressive`,
				},
				resource.Attribute{
					Name:        "moderate",
					Description: `Moderate`,
				},
				resource.Attribute{
					Name:        "conservative",
					Description: `Conservative ### Training Model Description`,
				},
				resource.Attribute{
					Name:        "low",
					Description: `Low`,
				},
				resource.Attribute{
					Name:        "medium",
					Description: `Medium`,
				},
				resource.Attribute{
					Name:        "high",
					Description: `High`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_anomaly_trusted_list",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "atl_id",
					Description: `(int) Anomaly Trusted List ID ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Anomaly Trusted List name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Reason for trusted listing`,
				},
				resource.Attribute{
					Name:        "trusted_list_type",
					Description: `Anomaly Trusted List type`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Anomaly Trusted List account id`,
				},
				resource.Attribute{
					Name:        "applicable_policies",
					Description: `Applicable Policies`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `VPC`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Created on`,
				},
				resource.Attribute{
					Name:        "trusted_list_entries",
					Description: `List of network anomalies in the trusted list [below](#trusted-list-entries). ### Trusted List Entries`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Image ID`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `Tag key`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `Tag value`,
				},
				resource.Attribute{
					Name:        "ip_cidr",
					Description: `IP CIDR`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `Resource ID`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `Subject`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Anomaly Trusted List name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Reason for trusted listing`,
				},
				resource.Attribute{
					Name:        "trusted_list_type",
					Description: `Anomaly Trusted List type`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Anomaly Trusted List account id`,
				},
				resource.Attribute{
					Name:        "applicable_policies",
					Description: `Applicable Policies`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `VPC`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Created on`,
				},
				resource.Attribute{
					Name:        "trusted_list_entries",
					Description: `List of network anomalies in the trusted list [below](#trusted-list-entries). ### Trusted List Entries`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Image ID`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `Tag key`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `Tag value`,
				},
				resource.Attribute{
					Name:        "ip_cidr",
					Description: `IP CIDR`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `Resource ID`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `Subject`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_anomaly_trusted_lists",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of anomaly trusted lists`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of anomaly trusted list, as defined [below](#listing). ### Listing`,
				},
				resource.Attribute{
					Name:        "atl_id",
					Description: `Anomaly Trusted List ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Anomaly Trusted List name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Reason for trusted listing`,
				},
				resource.Attribute{
					Name:        "trusted_list_type",
					Description: `Anomaly Trusted List type`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Anomaly Trusted List account id`,
				},
				resource.Attribute{
					Name:        "applicable_policies",
					Description: `Applicable Policies`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `VPC`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Created by`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Created on`,
				},
				resource.Attribute{
					Name:        "trusted_list_entries",
					Description: `List of network anomalies in the trusted list [below](#trusted-list-entries). ### Trusted List Entries`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Image ID`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `Tag key`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `Tag value`,
				},
				resource.Attribute{
					Name:        "ip_cidr",
					Description: `IP CIDR`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `Resource ID`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Service`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `Subject`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_aws_cft_generator_external_id",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_type",
					Description: `(Required) AWS account type. ` + "`" + `account` + "`" + ` or ` + "`" + `organization` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) AWS account ID.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `(Optional) List of features. If features key/field is not passed, then the default features will be applicable. Refer :`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `AWS account external ID.`,
				},
				resource.Attribute{
					Name:        "create_stack_link_with_s3_presigned_url",
					Description: `AWS account create stack link.`,
				},
				resource.Attribute{
					Name:        "event_bridge_rule_name_prefix",
					Description: `AWS account event bridge rule name prefix.`,
				},
				resource.Attribute{
					Name:        "s3_presigned_cft_url",
					Description: `AWS CFT S3 Presigned Unencoded URL.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "external_id",
					Description: `AWS account external ID.`,
				},
				resource.Attribute{
					Name:        "create_stack_link_with_s3_presigned_url",
					Description: `AWS account create stack link.`,
				},
				resource.Attribute{
					Name:        "event_bridge_rule_name_prefix",
					Description: `AWS account event bridge rule name prefix.`,
				},
				resource.Attribute{
					Name:        "s3_presigned_cft_url",
					Description: `AWS CFT S3 Presigned Unencoded URL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_aws_storage_uuid",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) AWS account ID.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Required) Unique identifier for an AWS resource (ARN).`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Required) External id for aws account. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `Storage UUID for aws account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `Storage UUID for aws account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_azure_template",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_type",
					Description: `(Required) Azure account type.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Azure tenant ID.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `(Required) File name to store azure template (Provide filename along with path to store azure template).`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Optional) Azure subscription ID.`,
				},
				resource.Attribute{
					Name:        "root_sync_enabled",
					Description: `(Optional) Azure tenant has children. Must be set to true if ` + "`" + `account_type` + "`" + ` is ` + "`" + `tenant` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `(Optional) Deployment type.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `(Optional) List of features. If features key/field is not passed, then the default features will be applicable. Refer :`,
				},
			},
			Attributes: []resource.Attribute{},
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
			Type:             "prismacloud_cloud_account_supported_features",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) Cloud type. ` + "`" + `aws` + "`" + `, ` + "`" + `azure` + "`" + `, or ` + "`" + `gcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `(Required) Cloud account type. ` + "`" + `account` + "`" + `, ` + "`" + `organization` + "`" + `, ` + "`" + `masterServiceAccount` + "`" + `, or ` + "`" + `tenant` + "`" + `. Supported values based on cloud_type are given below. <br /> - account, organization - cloud_type:`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "aws_partition",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "root_sync_enabled",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud type.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `Cloud Account deployment type.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Cloud account type.`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `Customer License type.`,
				},
				resource.Attribute{
					Name:        "supported_features",
					Description: `List of supported feature names.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud type.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `Cloud Account deployment type.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Cloud account type.`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `Customer License type.`,
				},
				resource.Attribute{
					Name:        "supported_features",
					Description: `List of supported feature names.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_cloud_account_v2",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) The cloud type. Valid value is ` + "`" + `aws` + "`" + `, ` + "`" + `azure` + "`" + `, ` + "`" + `gcp` + "`" + ` or ` + "`" + `ibm` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, computed) Cloud account name; computed if this is not supplied. Applicable only for ` + "`" + `aws` + "`" + `, ` + "`" + `azure` + "`" + ` and ` + "`" + `ibm` + "`" + `.`,
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
					Description: `Gcp account type spec, defined [below](#gcp).`,
				},
				resource.Attribute{
					Name:        "ibm",
					Description: `IBM account type spec, defined [below](#ibm). ### AWS`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `AWS account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the account is enabled.`,
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
					Description: `Unique identifier for an AWS resource (ARN).`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `Account created epoch time.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Prisma customer name.`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `Whether the account is deleted or not.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `aws` + "`" + ` for aws account.`,
				},
				resource.Attribute{
					Name:        "eventbridge_rule_name_prefix",
					Description: `Eventbridge rule name prefix.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External id for aws account.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features applicable for aws account, defined [below](#features).`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_epoch_millis",
					Description: `Last modified at epoch millis.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent id.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection mode of account.`,
				},
				resource.Attribute{
					Name:        "storage_scan_config",
					Description: `Storage scan config, defined [below](#storage_scan_config)`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `StorageUUID. ### Azure`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Azure account ID.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `Application ID registered with Active Directory.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the account is enabled.`,
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
					Description: `` + "`" + `account` + "`" + ` for azure subscription account.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection mode of account.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features applicable for azure account, defined [below](#features).`,
				},
				resource.Attribute{
					Name:        "environment_type",
					Description: `Environment type.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent id.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Prisma customer name.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `Account created epoch time.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_epoch_millis",
					Description: `Last modified at epoch millis.`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(bool) Whether the account is deleted or not.`,
				},
				resource.Attribute{
					Name:        "template_url",
					Description: `Template URL.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `az` + "`" + ` for azure account.`,
				},
				resource.Attribute{
					Name:        "deployment_type_description",
					Description: `Deployment type description. ### Gcp`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Gcp account ID.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `` + "`" + `account` + "`" + ` for gcp project account and ` + "`" + `masterServiceAccount` + "`" + ` for gcp master service account.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the account is enabled.`,
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
					Name:        "compression_enabled",
					Description: `(bool) Enable or disable compressed network flow log generation.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `Content of the JSON credentials file.`,
				},
				resource.Attribute{
					Name:        "data_flow_enabled_project",
					Description: `Project ID where the Dataflow API is enabled .`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features applicable for gcp account, defined [below](#features).`,
				},
				resource.Attribute{
					Name:        "flow_log_storage_bucket",
					Description: `Cloud Storage Bucket name that is used store the flow logs.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection mode of account.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent ID.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Prisma customer name.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `Account created epoch time.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_epoch_millis",
					Description: `Last modified at epoch millis.`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(bool) Whether the account is deleted or not.`,
				},
				resource.Attribute{
					Name:        "storage_scan_enabled",
					Description: `(bool) Whether the storage scan is enabled.`,
				},
				resource.Attribute{
					Name:        "added_on_ts",
					Description: `Added on time stamp.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `gcp` + "`" + ` for gcp account.`,
				},
				resource.Attribute{
					Name:        "deployment_type_description",
					Description: `Deployment type description.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Gcp Project ID.`,
				},
				resource.Attribute{
					Name:        "service_account_email",
					Description: `Service account email of gcp account.`,
				},
				resource.Attribute{
					Name:        "authentication_type",
					Description: `Authentication type of gcp account.`,
				},
				resource.Attribute{
					Name:        "account_group_creation_mode",
					Description: `Account group creation mode.`,
				},
				resource.Attribute{
					Name:        "default_account_group_id",
					Description: `Account group id to which you are assigning this account. ### IBM`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `IBM account ID.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `` + "`" + `account` + "`" + ` for IBM account.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `IBM service API key.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the account is enabled.`,
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
					Name:        "svc_id_iam_id",
					Description: `IBM service ID.`,
				},
				resource.Attribute{
					Name:        "added_on_ts",
					Description: `Added on time stamp.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `Account created epoch time.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Prisma customer name.`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(bool) Whether the account is deleted or not.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `ibm` + "`" + ` for IBM account.`,
				},
				resource.Attribute{
					Name:        "deployment_type_description",
					Description: `Deployment type description.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features applicable for IBM account, defined [below](#features).`,
				},
				resource.Attribute{
					Name:        "last_modified_epoch_millis",
					Description: `Last modified at epoch millis.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent id.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection mode of account.`,
				},
				resource.Attribute{
					Name:        "storage_scan_enabled",
					Description: `(bool) Whether the storage scan is enabled. #### FEATURES`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Feature name.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Feature state. ` + "`" + `enabled` + "`" + ` or ` + "`" + `disabled` + "`" + `. #### STORAGE_SCAN_CONFIG`,
				},
				resource.Attribute{
					Name:        "scan_option",
					Description: `Scan option.`,
				},
				resource.Attribute{
					Name:        "sns_topic_arn",
					Description: `SNS topic arn.`,
				},
				resource.Attribute{
					Name:        "buckets",
					Description: `List of buckets, defined [below](#buckets). #### BUCKETS`,
				},
				resource.Attribute{
					Name:        "backward",
					Description: `List of backward buckets.`,
				},
				resource.Attribute{
					Name:        "forward",
					Description: `List of forward buckets.`,
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
					Description: `Gcp account type spec, defined [below](#gcp).`,
				},
				resource.Attribute{
					Name:        "ibm",
					Description: `IBM account type spec, defined [below](#ibm). ### AWS`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `AWS account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the account is enabled.`,
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
					Description: `Unique identifier for an AWS resource (ARN).`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `Account created epoch time.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Prisma customer name.`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `Whether the account is deleted or not.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `aws` + "`" + ` for aws account.`,
				},
				resource.Attribute{
					Name:        "eventbridge_rule_name_prefix",
					Description: `Eventbridge rule name prefix.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External id for aws account.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features applicable for aws account, defined [below](#features).`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_epoch_millis",
					Description: `Last modified at epoch millis.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent id.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection mode of account.`,
				},
				resource.Attribute{
					Name:        "storage_scan_config",
					Description: `Storage scan config, defined [below](#storage_scan_config)`,
				},
				resource.Attribute{
					Name:        "storage_uuid",
					Description: `StorageUUID. ### Azure`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Azure account ID.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `Application ID registered with Active Directory.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the account is enabled.`,
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
					Description: `` + "`" + `account` + "`" + ` for azure subscription account.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection mode of account.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features applicable for azure account, defined [below](#features).`,
				},
				resource.Attribute{
					Name:        "environment_type",
					Description: `Environment type.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent id.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Prisma customer name.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `Account created epoch time.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_epoch_millis",
					Description: `Last modified at epoch millis.`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(bool) Whether the account is deleted or not.`,
				},
				resource.Attribute{
					Name:        "template_url",
					Description: `Template URL.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `az` + "`" + ` for azure account.`,
				},
				resource.Attribute{
					Name:        "deployment_type_description",
					Description: `Deployment type description. ### Gcp`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Gcp account ID.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `` + "`" + `account` + "`" + ` for gcp project account and ` + "`" + `masterServiceAccount` + "`" + ` for gcp master service account.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the account is enabled.`,
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
					Name:        "compression_enabled",
					Description: `(bool) Enable or disable compressed network flow log generation.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `Content of the JSON credentials file.`,
				},
				resource.Attribute{
					Name:        "data_flow_enabled_project",
					Description: `Project ID where the Dataflow API is enabled .`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features applicable for gcp account, defined [below](#features).`,
				},
				resource.Attribute{
					Name:        "flow_log_storage_bucket",
					Description: `Cloud Storage Bucket name that is used store the flow logs.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection mode of account.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent ID.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Prisma customer name.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `Account created epoch time.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_epoch_millis",
					Description: `Last modified at epoch millis.`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(bool) Whether the account is deleted or not.`,
				},
				resource.Attribute{
					Name:        "storage_scan_enabled",
					Description: `(bool) Whether the storage scan is enabled.`,
				},
				resource.Attribute{
					Name:        "added_on_ts",
					Description: `Added on time stamp.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `gcp` + "`" + ` for gcp account.`,
				},
				resource.Attribute{
					Name:        "deployment_type_description",
					Description: `Deployment type description.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Gcp Project ID.`,
				},
				resource.Attribute{
					Name:        "service_account_email",
					Description: `Service account email of gcp account.`,
				},
				resource.Attribute{
					Name:        "authentication_type",
					Description: `Authentication type of gcp account.`,
				},
				resource.Attribute{
					Name:        "account_group_creation_mode",
					Description: `Account group creation mode.`,
				},
				resource.Attribute{
					Name:        "default_account_group_id",
					Description: `Account group id to which you are assigning this account. ### IBM`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `IBM account ID.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `` + "`" + `account` + "`" + ` for IBM account.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `IBM service API key.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the account is enabled.`,
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
					Name:        "svc_id_iam_id",
					Description: `IBM service ID.`,
				},
				resource.Attribute{
					Name:        "added_on_ts",
					Description: `Added on time stamp.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `Account created epoch time.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Prisma customer name.`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(bool) Whether the account is deleted or not.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `ibm` + "`" + ` for IBM account.`,
				},
				resource.Attribute{
					Name:        "deployment_type_description",
					Description: `Deployment type description.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features applicable for IBM account, defined [below](#features).`,
				},
				resource.Attribute{
					Name:        "last_modified_epoch_millis",
					Description: `Last modified at epoch millis.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent id.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection mode of account.`,
				},
				resource.Attribute{
					Name:        "storage_scan_enabled",
					Description: `(bool) Whether the storage scan is enabled. #### FEATURES`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Feature name.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Feature state. ` + "`" + `enabled` + "`" + ` or ` + "`" + `disabled` + "`" + `. #### STORAGE_SCAN_CONFIG`,
				},
				resource.Attribute{
					Name:        "scan_option",
					Description: `Scan option.`,
				},
				resource.Attribute{
					Name:        "sns_topic_arn",
					Description: `SNS topic arn.`,
				},
				resource.Attribute{
					Name:        "buckets",
					Description: `List of buckets, defined [below](#buckets). #### BUCKETS`,
				},
				resource.Attribute{
					Name:        "backward",
					Description: `List of backward buckets.`,
				},
				resource.Attribute{
					Name:        "forward",
					Description: `List of forward buckets.`,
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
				resource.Attribute{
					Name:        "named_users_access_keys_expiry_notifications_enabled",
					Description: `(bool) Named users access keys expiry notifications enabled.`,
				},
				resource.Attribute{
					Name:        "service_users_access_keys_expiry_notifications_enabled",
					Description: `(bool) Service users access keys expiry notifications enabled.`,
				},
				resource.Attribute{
					Name:        "notification_threshold_access_keys_expiry",
					Description: `(int) Notification threshold access keys expiry.`,
				},
				resource.Attribute{
					Name:        "audit_log_siem_intgr_ids",
					Description: `List of integration ids.`,
				},
				resource.Attribute{
					Name:        "audit_logs_enabled",
					Description: `(bool) Enable audit logs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_gcp_template",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform (must be unique).`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `(Required) Gcp account type.`,
				},
				resource.Attribute{
					Name:        "authentication_type",
					Description: `(Required) Authentication type of gcp account.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Gcp Project ID. Must be provided for account type ` + "`" + `account` + "`" + ` and ` + "`" + `masterServiceAccount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `(Optional) Gcp organization ID. Must be provided for account type ` + "`" + `organization` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `(Required) File name to store gcp template (Provide filename along with path to store gcp template).`,
				},
				resource.Attribute{
					Name:        "flow_log_storage_bucket",
					Description: `(Optional) Cloud Storage Bucket name that is used store the flow logs.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `(Optional) List of features. If features key/field is not passed, then the default features will be applicable. Refer :`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_ibm_template",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_type",
					Description: `(Required) IBM account type.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `(Required) File name to store ibm template (Provide filename along with path to store ibm template).`,
				},
			},
			Attributes: []resource.Attribute{},
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
			Type:             "prismacloud_notification_template",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Notification template ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "integration_id",
					Description: `Integration ID.`,
				},
				resource.Attribute{
					Name:        "created_ts",
					Description: `(int) Creation Unix timestamp in milliseconds.`,
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
					Name:        "integration_type",
					Description: `Integration type.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `User who created the notification template.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the template on the Prisma Cloud platform.`,
				},
				resource.Attribute{
					Name:        "integration_name",
					Description: `Integration name.`,
				},
				resource.Attribute{
					Name:        "customer_id",
					Description: `(int) Prisma customer ID.`,
				},
				resource.Attribute{
					Name:        "module",
					Description: `Module.`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `Type of notification template.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the template is enabled.`,
				},
				resource.Attribute{
					Name:        "template_config",
					Description: `Template config spec, as defined [below](#template_config). ## Template Config`,
				},
				resource.Attribute{
					Name:        "basic_config",
					Description: `This field includes additional attributes that can be used to customize the notification, as defined [below](#config_params).`,
				},
				resource.Attribute{
					Name:        "open",
					Description: `Provide config to map the ` + "`" + `open` + "`" + ` alert state to ` + "`" + `jira` + "`" + `/` + "`" + `service_now` + "`" + ` state and configure the ` + "`" + `jira` + "`" + `/` + "`" + `service_now` + "`" + ` fields. This field includes additional attributes, as defined [below](#config_params).`,
				},
				resource.Attribute{
					Name:        "resolved",
					Description: `Provide config to map the ` + "`" + `resolved` + "`" + ` alert state to ` + "`" + `jira` + "`" + `/` + "`" + `service_now` + "`" + ` state and configure the ` + "`" + `jira` + "`" + `/` + "`" + `service_now` + "`" + ` fields. This field includes additional attributes, as defined [below](#config_params).`,
				},
				resource.Attribute{
					Name:        "dismissed",
					Description: `Provide config to map the ` + "`" + `dismissed` + "`" + ` alert state to ` + "`" + `service_now` + "`" + ` state and configure the ` + "`" + `service_now` + "`" + ` fields. This field includes additional attributes, as defined [below](#config_params).`,
				},
				resource.Attribute{
					Name:        "snoozed",
					Description: `This field represents the notification status when the user has chosen to temporarily delay or "snooze" the notification. This field includes additional attributes, as defined [below](#config_params). ### Config Params`,
				},
				resource.Attribute{
					Name:        "field_name",
					Description: `Field name.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name.`,
				},
				resource.Attribute{
					Name:        "redlock_mapping",
					Description: `(bool) Prisma Cloud will provide the field value for notification.`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(bool) Required.`,
				},
				resource.Attribute{
					Name:        "type_ahead_uri",
					Description: `URL used to query suggestions for field value.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of field.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Field value.`,
				},
				resource.Attribute{
					Name:        "alias_field",
					Description: `Alias field.`,
				},
				resource.Attribute{
					Name:        "max_length",
					Description: `(int) Maximum length.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `Options, as defined [below](#options). #### Options`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Field option name.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Field option key.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Field option ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "integration_id",
					Description: `Integration ID.`,
				},
				resource.Attribute{
					Name:        "created_ts",
					Description: `(int) Creation Unix timestamp in milliseconds.`,
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
					Name:        "integration_type",
					Description: `Integration type.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `User who created the notification template.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the template on the Prisma Cloud platform.`,
				},
				resource.Attribute{
					Name:        "integration_name",
					Description: `Integration name.`,
				},
				resource.Attribute{
					Name:        "customer_id",
					Description: `(int) Prisma customer ID.`,
				},
				resource.Attribute{
					Name:        "module",
					Description: `Module.`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `Type of notification template.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the template is enabled.`,
				},
				resource.Attribute{
					Name:        "template_config",
					Description: `Template config spec, as defined [below](#template_config). ## Template Config`,
				},
				resource.Attribute{
					Name:        "basic_config",
					Description: `This field includes additional attributes that can be used to customize the notification, as defined [below](#config_params).`,
				},
				resource.Attribute{
					Name:        "open",
					Description: `Provide config to map the ` + "`" + `open` + "`" + ` alert state to ` + "`" + `jira` + "`" + `/` + "`" + `service_now` + "`" + ` state and configure the ` + "`" + `jira` + "`" + `/` + "`" + `service_now` + "`" + ` fields. This field includes additional attributes, as defined [below](#config_params).`,
				},
				resource.Attribute{
					Name:        "resolved",
					Description: `Provide config to map the ` + "`" + `resolved` + "`" + ` alert state to ` + "`" + `jira` + "`" + `/` + "`" + `service_now` + "`" + ` state and configure the ` + "`" + `jira` + "`" + `/` + "`" + `service_now` + "`" + ` fields. This field includes additional attributes, as defined [below](#config_params).`,
				},
				resource.Attribute{
					Name:        "dismissed",
					Description: `Provide config to map the ` + "`" + `dismissed` + "`" + ` alert state to ` + "`" + `service_now` + "`" + ` state and configure the ` + "`" + `service_now` + "`" + ` fields. This field includes additional attributes, as defined [below](#config_params).`,
				},
				resource.Attribute{
					Name:        "snoozed",
					Description: `This field represents the notification status when the user has chosen to temporarily delay or "snooze" the notification. This field includes additional attributes, as defined [below](#config_params). ### Config Params`,
				},
				resource.Attribute{
					Name:        "field_name",
					Description: `Field name.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name.`,
				},
				resource.Attribute{
					Name:        "redlock_mapping",
					Description: `(bool) Prisma Cloud will provide the field value for notification.`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(bool) Required.`,
				},
				resource.Attribute{
					Name:        "type_ahead_uri",
					Description: `URL used to query suggestions for field value.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of field.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Field value.`,
				},
				resource.Attribute{
					Name:        "alias_field",
					Description: `Alias field.`,
				},
				resource.Attribute{
					Name:        "max_length",
					Description: `(int) Maximum length.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `Options, as defined [below](#options). #### Options`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Field option name.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Field option key.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Field option ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_notification_templates",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of notification templates.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of notification templates returned, as defined [below](#listing). ### Listing Each notification template has the following attributes:`,
				},
				resource.Attribute{
					Name:        "integration_id",
					Description: `Integration ID.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Notification template id.`,
				},
				resource.Attribute{
					Name:        "created_ts",
					Description: `(int) Creation Unix timestamp in milliseconds.`,
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
					Name:        "created_by",
					Description: `User who created the notification template.`,
				},
				resource.Attribute{
					Name:        "integration_type",
					Description: `Integration type.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the template on the Prisma Cloud platform.`,
				},
				resource.Attribute{
					Name:        "integration_name",
					Description: `Integration name.`,
				},
				resource.Attribute{
					Name:        "customer_id",
					Description: `(int) Prisma customer id.`,
				},
				resource.Attribute{
					Name:        "module",
					Description: `Module.`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `Type of notification template.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the template is enabled.`,
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
			Type:             "prismacloud_org_cloud_account_v2",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) The cloud type. Valid value is ` + "`" + `aws` + "`" + `, ` + "`" + `azure` + "`" + ` or ` + "`" + `gcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, computed) Cloud account name; computed if this is not supplied. Applicable only for ` + "`" + `aws` + "`" + `, ` + "`" + `azure` + "`" + ` and ` + "`" + `ibm` + "`" + `.`,
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
					Description: `Gcp account type spec, defined [below](#gcp). ### AWS`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `AWS account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the account is enabled.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you have assigned this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `Unique identifier for an AWS resource (ARN).`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `Account created epoch time.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Prisma customer name.`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `Whether the account is deleted or not.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `aws` + "`" + ` for aws account.`,
				},
				resource.Attribute{
					Name:        "eventbridge_rule_name_prefix",
					Description: `Eventbridge rule name prefix.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External id for aws account.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features applicable for aws account, defined [below](#features).`,
				},
				resource.Attribute{
					Name:        "has_member_role",
					Description: `Whether account has member role.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_epoch_millis",
					Description: `Last modified at epoch millis.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent id.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection mode of account.`,
				},
				resource.Attribute{
					Name:        "hierarchy_selection",
					Description: `List of hierarchy selection. Each item has resource ID, display name, node type and selection type, as defined [below](#hierarchy-selection). ### Azure`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Azure tenant account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the account is enabled.`,
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
					Description: `(bool) Automatically ingest flow logs. Must be set to false when azure tenant is onboarded without children i.e., for ` + "`" + `Active Directory Tenant` + "`" + `.`,
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
					Description: `` + "`" + `tenant` + "`" + ` for azure account.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection mode of account.`,
				},
				resource.Attribute{
					Name:        "default_account_group_id",
					Description: `Account group id to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "root_sync_enabled",
					Description: `(bool) Azure tenant has children. Must be set to true when azure tenant is onboarded with children i.e., for ` + "`" + `Tenant` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hierarchy_selection",
					Description: `List of hierarchy selection. Each item has resource ID, display name, node type and selection type, as defined [below](#hierarchy-selection).`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features applicable for azure account, defined [below](#features).`,
				},
				resource.Attribute{
					Name:        "environment_type",
					Description: `Environment type.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent id.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Prisma customer name.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `Account created epoch time.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_epoch_millis",
					Description: `Last modified at epoch millis.`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(bool) Whether the account is deleted or not.`,
				},
				resource.Attribute{
					Name:        "template_url",
					Description: `Template URL.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `az` + "`" + ` for azure account.`,
				},
				resource.Attribute{
					Name:        "deployment_type_description",
					Description: `Deployment type description.`,
				},
				resource.Attribute{
					Name:        "member_sync_enabled",
					Description: `(bool) Azure tenant has children. Must be set to true when azure tenant is onboarded with children i.e., for ` + "`" + `Tenant` + "`" + `. ### Gcp`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Gcp account ID.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `` + "`" + `organization` + "`" + ` for gcp organization account.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the account is enabled.`,
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
					Name:        "compression_enabled",
					Description: `(bool) Enable or disable compressed network flow log generation.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `Content of the JSON credentials file.`,
				},
				resource.Attribute{
					Name:        "data_flow_enabled_project",
					Description: `Project ID where the Dataflow API is enabled.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features applicable for gcp account, defined [below](#features).`,
				},
				resource.Attribute{
					Name:        "flow_log_storage_bucket",
					Description: `Cloud Storage Bucket name that is used store the flow logs.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection mode of account.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent ID.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Prisma customer name.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `Account created epoch time.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_epoch_millis",
					Description: `Last modified at epoch millis.`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(bool) Whether the account is deleted or not.`,
				},
				resource.Attribute{
					Name:        "storage_scan_enabled",
					Description: `(bool) Whether the storage scan is enabled.`,
				},
				resource.Attribute{
					Name:        "added_on_ts",
					Description: `Added on time stamp.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `gcp` + "`" + ` for gcp account.`,
				},
				resource.Attribute{
					Name:        "deployment_type_description",
					Description: `Deployment type description.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Gcp Project ID.`,
				},
				resource.Attribute{
					Name:        "service_account_email",
					Description: `Service account email of gcp account.`,
				},
				resource.Attribute{
					Name:        "authentication_type",
					Description: `Authentication type of gcp account.`,
				},
				resource.Attribute{
					Name:        "account_group_creation_mode",
					Description: `Account group creation mode.`,
				},
				resource.Attribute{
					Name:        "default_account_group_id",
					Description: `Account group id to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "hierarchy_selection",
					Description: `List of hierarchy selection. Each item has resource ID, display name, node type and selection type, as defined [below](#hierarchy-selection).`,
				},
				resource.Attribute{
					Name:        "organization_name",
					Description: `Gcp organization name. #### Hierarchy Selection`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `Resource ID. For ACCOUNT, OU, ROOT, TENANT, SUBSCRIPTION, PROJECT, FOLDER or ORG. Example : ` + "`" + `root` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for ACCOUNT, OU, ROOT, TENANT, SUBSCRIPTION, PROJECT, FOLDER or ORG. Example : ` + "`" + `Root` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `Node type - ORG, OU, ACCOUNT, SUBSCRIPTION, TENANT, MANAGEMENT_GROUP, PROJECT, FOLDER or ORG.`,
				},
				resource.Attribute{
					Name:        "selection_type",
					Description: `Selection type. Valid values: INCLUDE to include the specified resource to onboard, EXCLUDE to exclude the specified resource and onboard the rest, ALL to onboard all resources in the organization. #### FEATURES`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Feature name.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Feature state. ` + "`" + `enabled` + "`" + ` or ` + "`" + `disabled` + "`" + `.`,
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
					Description: `Gcp account type spec, defined [below](#gcp). ### AWS`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `AWS account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the account is enabled.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `List of account IDs to which you have assigned this account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name to be used for the account on the Prisma Cloud platform.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `Unique identifier for an AWS resource (ARN).`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `Account created epoch time.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Prisma customer name.`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `Whether the account is deleted or not.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `aws` + "`" + ` for aws account.`,
				},
				resource.Attribute{
					Name:        "eventbridge_rule_name_prefix",
					Description: `Eventbridge rule name prefix.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External id for aws account.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features applicable for aws account, defined [below](#features).`,
				},
				resource.Attribute{
					Name:        "has_member_role",
					Description: `Whether account has member role.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_epoch_millis",
					Description: `Last modified at epoch millis.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent id.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection mode of account.`,
				},
				resource.Attribute{
					Name:        "hierarchy_selection",
					Description: `List of hierarchy selection. Each item has resource ID, display name, node type and selection type, as defined [below](#hierarchy-selection). ### Azure`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Azure tenant account ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the account is enabled.`,
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
					Description: `(bool) Automatically ingest flow logs. Must be set to false when azure tenant is onboarded without children i.e., for ` + "`" + `Active Directory Tenant` + "`" + `.`,
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
					Description: `` + "`" + `tenant` + "`" + ` for azure account.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection mode of account.`,
				},
				resource.Attribute{
					Name:        "default_account_group_id",
					Description: `Account group id to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "root_sync_enabled",
					Description: `(bool) Azure tenant has children. Must be set to true when azure tenant is onboarded with children i.e., for ` + "`" + `Tenant` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hierarchy_selection",
					Description: `List of hierarchy selection. Each item has resource ID, display name, node type and selection type, as defined [below](#hierarchy-selection).`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features applicable for azure account, defined [below](#features).`,
				},
				resource.Attribute{
					Name:        "environment_type",
					Description: `Environment type.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent id.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Prisma customer name.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `Account created epoch time.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_epoch_millis",
					Description: `Last modified at epoch millis.`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(bool) Whether the account is deleted or not.`,
				},
				resource.Attribute{
					Name:        "template_url",
					Description: `Template URL.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `az` + "`" + ` for azure account.`,
				},
				resource.Attribute{
					Name:        "deployment_type_description",
					Description: `Deployment type description.`,
				},
				resource.Attribute{
					Name:        "member_sync_enabled",
					Description: `(bool) Azure tenant has children. Must be set to true when azure tenant is onboarded with children i.e., for ` + "`" + `Tenant` + "`" + `. ### Gcp`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Gcp account ID.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `` + "`" + `organization` + "`" + ` for gcp organization account.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(bool) Whether the account is enabled.`,
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
					Name:        "compression_enabled",
					Description: `(bool) Enable or disable compressed network flow log generation.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `Content of the JSON credentials file.`,
				},
				resource.Attribute{
					Name:        "data_flow_enabled_project",
					Description: `Project ID where the Dataflow API is enabled.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features applicable for gcp account, defined [below](#features).`,
				},
				resource.Attribute{
					Name:        "flow_log_storage_bucket",
					Description: `Cloud Storage Bucket name that is used store the flow logs.`,
				},
				resource.Attribute{
					Name:        "protection_mode",
					Description: `Protection mode of account.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent ID.`,
				},
				resource.Attribute{
					Name:        "customer_name",
					Description: `Prisma customer name.`,
				},
				resource.Attribute{
					Name:        "created_epoch_millis",
					Description: `Account created epoch time.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_epoch_millis",
					Description: `Last modified at epoch millis.`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(bool) Whether the account is deleted or not.`,
				},
				resource.Attribute{
					Name:        "storage_scan_enabled",
					Description: `(bool) Whether the storage scan is enabled.`,
				},
				resource.Attribute{
					Name:        "added_on_ts",
					Description: `Added on time stamp.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `` + "`" + `gcp` + "`" + ` for gcp account.`,
				},
				resource.Attribute{
					Name:        "deployment_type_description",
					Description: `Deployment type description.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Gcp Project ID.`,
				},
				resource.Attribute{
					Name:        "service_account_email",
					Description: `Service account email of gcp account.`,
				},
				resource.Attribute{
					Name:        "authentication_type",
					Description: `Authentication type of gcp account.`,
				},
				resource.Attribute{
					Name:        "account_group_creation_mode",
					Description: `Account group creation mode.`,
				},
				resource.Attribute{
					Name:        "default_account_group_id",
					Description: `Account group id to which you are assigning this account.`,
				},
				resource.Attribute{
					Name:        "hierarchy_selection",
					Description: `List of hierarchy selection. Each item has resource ID, display name, node type and selection type, as defined [below](#hierarchy-selection).`,
				},
				resource.Attribute{
					Name:        "organization_name",
					Description: `Gcp organization name. #### Hierarchy Selection`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `Resource ID. For ACCOUNT, OU, ROOT, TENANT, SUBSCRIPTION, PROJECT, FOLDER or ORG. Example : ` + "`" + `root` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for ACCOUNT, OU, ROOT, TENANT, SUBSCRIPTION, PROJECT, FOLDER or ORG. Example : ` + "`" + `Root` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `Node type - ORG, OU, ACCOUNT, SUBSCRIPTION, TENANT, MANAGEMENT_GROUP, PROJECT, FOLDER or ORG.`,
				},
				resource.Attribute{
					Name:        "selection_type",
					Description: `Selection type. Valid values: INCLUDE to include the specified resource to onboard, EXCLUDE to exclude the specified resource and onboard the rest, ALL to onboard all resources in the organization. #### FEATURES`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Feature name.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Feature state. ` + "`" + `enabled` + "`" + ` or ` + "`" + `disabled` + "`" + `.`,
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
			Type:             "prismacloud_permission_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Permission group id`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the permission group. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "permission_group_type",
					Description: `Permission group type.`,
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
					Name:        "associated_roles",
					Description: `List of associated user roles which cannot exist in the system without the permission group.`,
				},
				resource.Attribute{
					Name:        "accept_account_groups",
					Description: `(bool) Accept account groups.`,
				},
				resource.Attribute{
					Name:        "accept_resource_lists",
					Description: `(bool) Accept resource lists.`,
				},
				resource.Attribute{
					Name:        "accept_code_repositories",
					Description: `(bool) Accept code repositories.`,
				},
				resource.Attribute{
					Name:        "custom",
					Description: `(bool) Boolean value signifying whether this is a custom (i.e. user-defined) permission group.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Collection of permitted features associated with the role, as defined [below](#features). ### Features`,
				},
				resource.Attribute{
					Name:        "feature_name",
					Description: `Prisma Cloud Feature Name.`,
				},
				resource.Attribute{
					Name:        "operations",
					Description: `A mapping of operations and a boolean value representing whether the privilege to perform the operation needs to be granted, as defined [below](#operations). #### Operations`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "permission_group_type",
					Description: `Permission group type.`,
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
					Name:        "associated_roles",
					Description: `List of associated user roles which cannot exist in the system without the permission group.`,
				},
				resource.Attribute{
					Name:        "accept_account_groups",
					Description: `(bool) Accept account groups.`,
				},
				resource.Attribute{
					Name:        "accept_resource_lists",
					Description: `(bool) Accept resource lists.`,
				},
				resource.Attribute{
					Name:        "accept_code_repositories",
					Description: `(bool) Accept code repositories.`,
				},
				resource.Attribute{
					Name:        "custom",
					Description: `(bool) Boolean value signifying whether this is a custom (i.e. user-defined) permission group.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Collection of permitted features associated with the role, as defined [below](#features). ### Features`,
				},
				resource.Attribute{
					Name:        "feature_name",
					Description: `Prisma Cloud Feature Name.`,
				},
				resource.Attribute{
					Name:        "operations",
					Description: `A mapping of operations and a boolean value representing whether the privilege to perform the operation needs to be granted, as defined [below](#operations). #### Operations`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_permission_groups",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of permission groups.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of permission groups returned, as defined [below](#listing). ### Listing Each permission group has the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Permission group id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the permission group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of permission group.`,
				},
				resource.Attribute{
					Name:        "permission_group_type",
					Description: `Permission group type.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `Last modified by.`,
				},
				resource.Attribute{
					Name:        "last_modified_ts",
					Description: `Last modified timestamp.`,
				},
				resource.Attribute{
					Name:        "accept_account_groups",
					Description: `(bool) Accept account groups.`,
				},
				resource.Attribute{
					Name:        "accept_resource_lists",
					Description: `(bool) Accept resource lists.`,
				},
				resource.Attribute{
					Name:        "accept_code_repositories",
					Description: `(bool) Accept code repositories.`,
				},
				resource.Attribute{
					Name:        "custom",
					Description: `(bool) Boolean value signifying whether this is a custom (i.e. user-defined) permission group.`,
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
					Description: `(Optional, map) Filters to limit policies returned. Filter options can be found [here](https://prisma.pan.dev/api/cloud/cspm/policy#operation/get-policies). ## Attribute Reference`,
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
					Name:        "children",
					Description: `Children description for build policy, as defined [below](#children)`,
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
					Description: `List of file extensions #### Children`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `Criteria for build policy.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `YAML string for code build policy.`,
				},
				resource.Attribute{
					Name:        "recommendation",
					Description: `Recommendation.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of policy.`,
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
					Name:        "children",
					Description: `Children description for build policy, as defined [below](#children)`,
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
					Description: `List of file extensions #### Children`,
				},
				resource.Attribute{
					Name:        "criteria",
					Description: `Criteria for build policy.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `YAML string for code build policy.`,
				},
				resource.Attribute{
					Name:        "recommendation",
					Description: `Recommendation.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of policy.`,
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
			Type:             "prismacloud_trusted_alert_ip",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the trusted alert ip.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "cidrs",
					Description: `List of CIDRs, as defined [below](#CIDR).`,
				},
				resource.Attribute{
					Name:        "cidr_count",
					Description: `CIDR count. ### CIDR`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(string) Ip address.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID for cidr.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `(int) Created on.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidrs",
					Description: `List of CIDRs, as defined [below](#CIDR).`,
				},
				resource.Attribute{
					Name:        "cidr_count",
					Description: `CIDR count. ### CIDR`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(string) Ip address.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID for cidr.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `(int) Created on.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_trusted_alert_ips",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of trusted alert ips.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of trusted alert ips returned, as defined [below](#listing). ### Listing Each trusted alert ip has the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the trusted alert ip.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID.`,
				},
				resource.Attribute{
					Name:        "cidrs",
					Description: `List of CIDRs, as defined [below](#CIDR).`,
				},
				resource.Attribute{
					Name:        "cidr_count",
					Description: `CIDR count. ### CIDR`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(string) Ip address.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID for cidr.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `(int) Created on.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_trusted_login_ip",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the trusted login ip Allow List.`,
				},
				resource.Attribute{
					Name:        "trusted_login_ip_id",
					Description: `Trusted login ip allow List ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the list of CIDR blocks that are in allow list for access.`,
				},
				resource.Attribute{
					Name:        "trusted_login_ip_id",
					Description: `Login IP allow list ID`,
				},
				resource.Attribute{
					Name:        "last_modified_ts",
					Description: `Timestamp for last modification of CIDR block list.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `List of CIDR blocks (IP addresses) from which access is allowed when Login IP Allow List is enabled.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the list of CIDR blocks that are in allow list for access.`,
				},
				resource.Attribute{
					Name:        "trusted_login_ip_id",
					Description: `Login IP allow list ID`,
				},
				resource.Attribute{
					Name:        "last_modified_ts",
					Description: `Timestamp for last modification of CIDR block list.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `List of CIDR blocks (IP addresses) from which access is allowed when Login IP Allow List is enabled.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "prismacloud_trusted_login_ips",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of trusted login ips.`,
				},
				resource.Attribute{
					Name:        "listing",
					Description: `List of trusted login Ips, as defined [below](#listing). ### Listing`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the list of CIDR blocks that are in allow list for access.`,
				},
				resource.Attribute{
					Name:        "trusted_login_ip_id",
					Description: `Login IP allow list ID`,
				},
				resource.Attribute{
					Name:        "last_modified_ts",
					Description: `Timestamp for last modification of CIDR block list.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `List of CIDR blocks (IP addresses) from which access is allowed when Login IP Allow List is enabled.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
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
		"prismacloud_anomaly_setting":                          6,
		"prismacloud_anomaly_settings":                         7,
		"prismacloud_anomaly_trusted_list":                     8,
		"prismacloud_anomaly_trusted_lists":                    9,
		"prismacloud_aws_cft_generator_external_id":            10,
		"prismacloud_aws_storage_uuid":                         11,
		"prismacloud_azure_template":                           12,
		"prismacloud_cloud_account":                            13,
		"prismacloud_cloud_account_supported_features":         14,
		"prismacloud_cloud_account_v2":                         15,
		"prismacloud_cloud_accounts":                           16,
		"prismacloud_compliance_standard":                      17,
		"prismacloud_compliance_standard_requirement":          18,
		"prismacloud_compliance_standard_requirement_section":  19,
		"prismacloud_compliance_standard_requirement_sections": 20,
		"prismacloud_compliance_standard_requirements":         21,
		"prismacloud_compliance_standards":                     22,
		"prismacloud_datapattern":                              23,
		"prismacloud_datapatterns":                             24,
		"prismacloud_dataprofile":                              25,
		"prismacloud_dataprofiles":                             26,
		"prismacloud_enterprise_settings":                      27,
		"prismacloud_gcp_template":                             28,
		"prismacloud_ibm_template":                             29,
		"prismacloud_integration":                              30,
		"prismacloud_integrations":                             31,
		"prismacloud_notification_template":                    32,
		"prismacloud_notification_templates":                   33,
		"prismacloud_org_cloud_account":                        34,
		"prismacloud_org_cloud_account_v2":                     35,
		"prismacloud_org_cloud_accounts":                       36,
		"prismacloud_permission_group":                         37,
		"prismacloud_permission_groups":                        38,
		"prismacloud_policies":                                 39,
		"prismacloud_policy":                                   40,
		"prismacloud_report":                                   41,
		"prismacloud_reports":                                  42,
		"prismacloud_rql_historic_search":                      43,
		"prismacloud_rql_historic_searches":                    44,
		"prismacloud_trusted_alert_ip":                         45,
		"prismacloud_trusted_alert_ips":                        46,
		"prismacloud_trusted_login_ip":                         47,
		"prismacloud_trusted_login_ips":                        48,
		"prismacloud_user_profile":                             49,
		"prismacloud_user_profiles":                            50,
		"prismacloud_user_role":                                51,
		"prismacloud_user_roles":                               52,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
