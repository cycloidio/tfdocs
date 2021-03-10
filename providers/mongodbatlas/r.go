package mongodbatlas

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_alert_configuration",
			Category:         "Resources",
			ShortDescription: `Provides an Alert Configuration resource.`,
			Description:      ``,
			Keywords: []string{
				"alert",
				"configuration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project where the alert configuration will create.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `It is not required, but If the attribute is omitted, by default will be false, and the configuration would be disabled. You must set true to enable the configuration.`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `(Required) The type of event that will trigger an alert. ->`,
				},
				resource.Attribute{
					Name:        "field_name",
					Description: `Name of the field in the target object to match on. | Host alerts | Replica set alerts | Sharded cluster alerts | |:---------- |:------------- |:------ | | ` + "`" + `TYPE_NAME` + "`" + ` | ` + "`" + `REPLICA_SET_NAME` + "`" + ` | ` + "`" + `CLUSTER_NAME` + "`" + ` | | ` + "`" + `HOSTNAME` + "`" + ` | ` + "`" + `SHARD_NAME` + "`" + ` | ` + "`" + `SHARD_NAME` + "`" + ` | | ` + "`" + `PORT` + "`" + ` | ` + "`" + `CLUSTER_NAME` + "`" + ` | | | ` + "`" + `HOSTNAME_AND_PORT` + "`" + ` | | | | ` + "`" + `REPLICA_SET_NAME` + "`" + ` | | | All other types of alerts do not support matchers.`,
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
					Description: `Optional. The following roles grant privileges within a project. Accepted values are: | Project roles | Organization roles | |:---------- |:----------- | | ` + "`" + `GROUP_CHARTS_ADMIN` + "`" + ` | ` + "`" + `ORG_OWNER` + "`" + ` | | ` + "`" + `GROUP_CLUSTER_MANAGER` + "`" + ` | ` + "`" + `ORG_MEMBER` + "`" + ` | | ` + "`" + `GROUP_DATA_ACCESS_ADMIN` + "`" + ` | ` + "`" + `ORG_GROUP_CREATOR` + "`" + ` | | ` + "`" + `GROUP_DATA_ACCESS_READ_ONLY` + "`" + ` | ` + "`" + `ORG_BILLING_ADMIN` + "`" + ` | | ` + "`" + `GROUP_DATA_ACCESS_READ_WRITE` + "`" + ` | ` + "`" + `ORG_READ_ONLY` + "`" + ` | | ` + "`" + `GROUP_OWNER` + "`" + ` | | | ` + "`" + `GROUP_READ_ONLY` + "`" + ` | | ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages and can be used to import.`,
				},
				resource.Attribute{
					Name:        "alert_configuration_id",
					Description: `Unique identifier for the alert configuration.`,
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
					Description: `Timestamp in ISO 8601 date and time format in UTC when this alert configuration was last updated. ## Import Alert Configuration can be imported using the ` + "`" + `project_id-alert_configuration_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_alert_configuration.test 5d0f1f74cf09a29120e123cd-5d0f1f74cf09a29120e1fscg ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/alert-configurations/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages and can be used to import.`,
				},
				resource.Attribute{
					Name:        "alert_configuration_id",
					Description: `Unique identifier for the alert configuration.`,
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
					Description: `Timestamp in ISO 8601 date and time format in UTC when this alert configuration was last updated. ## Import Alert Configuration can be imported using the ` + "`" + `project_id-alert_configuration_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_alert_configuration.test 5d0f1f74cf09a29120e123cd-5d0f1f74cf09a29120e1fscg ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/alert-configurations/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_auditing",
			Category:         "Resources",
			ShortDescription: `Provides a Auditing resource.`,
			Description:      ``,
			Keywords: []string{
				"auditing",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to configure auditing.`,
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
					Description: `Denotes whether or not the project associated with the {project_id} has database auditing enabled. ~>`,
				},
				resource.Attribute{
					Name:        "configuration_type",
					Description: `Denotes the configuration method for the audit filter. Possible values are:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "configuration_type",
					Description: `Denotes the configuration method for the audit filter. Possible values are:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot",
			Category:         "Resources",
			ShortDescription: `Provides an Cloud Backup Snapshot resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"provider",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Atlas cluster that contains the snapshots you want to retrieve.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the on-demand snapshot.`,
				},
				resource.Attribute{
					Name:        "retention_in_days",
					Description: `(Required) The number of days that Atlas should retain the on-demand snapshot. Must be at least 1. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Unique identifier of the snapshot.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas took the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the snapshot. Only present for on-demand snapshots.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas will delete the snapshot.`,
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
					Description: `Current status of the snapshot. One of the following values will be returned: queued, inProgress, completed, failed.`,
				},
				resource.Attribute{
					Name:        "storage_size_bytes",
					Description: `Specifies the size of the snapshot in bytes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. ## Import Cloud Backup Snapshot entries can be imported using project project_id, cluster_name and snapshot_id (Unique identifier of the snapshot), in the format ` + "`" + `PROJECTID-CLUSTERNAME-SNAPSHOTID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_provider_snapshot.test 5d0f1f73cf09a29120e173cf-MyClusterTest-5d116d82014b764445b2f9b5 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/backups/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Unique identifier of the snapshot.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas took the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the snapshot. Only present for on-demand snapshots.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas will delete the snapshot.`,
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
					Description: `Current status of the snapshot. One of the following values will be returned: queued, inProgress, completed, failed.`,
				},
				resource.Attribute{
					Name:        "storage_size_bytes",
					Description: `Specifies the size of the snapshot in bytes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. ## Import Cloud Backup Snapshot entries can be imported using project project_id, cluster_name and snapshot_id (Unique identifier of the snapshot), in the format ` + "`" + `PROJECTID-CLUSTERNAME-SNAPSHOTID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_provider_snapshot.test 5d0f1f73cf09a29120e173cf-MyClusterTest-5d116d82014b764445b2f9b5 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/backups/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot_backup_policy",
			Category:         "Resources",
			ShortDescription: `Provides a Cloud Backup Snapshot Policy resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"provider",
				"snapshot",
				"backup",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Atlas cluster that contains the snapshot backup policy you want to retrieve.`,
				},
				resource.Attribute{
					Name:        "reference_hour_of_day",
					Description: `(Optional) UTC Hour of day between 0 and 23, inclusive, representing which hour of the day that Atlas takes snapshots for backup policy items.`,
				},
				resource.Attribute{
					Name:        "reference_minute_of_hour",
					Description: `(Optional) UTC Minutes after referenceHourOfDay that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.`,
				},
				resource.Attribute{
					Name:        "restore_window_days",
					Description: `(Optional) Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.`,
				},
				resource.Attribute{
					Name:        "update_snapshots",
					Description: `(Optional) Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously. ### Policies`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) Contains a document for each backup policy item in the desired updated backup policy.`,
				},
				resource.Attribute{
					Name:        "policies.#.id",
					Description: `(Required) Unique identifier of the backup policy that you want to update. policies.#.id is a value obtained via the mongodbatlas_cluster resource. provider_backup_enabled of the mongodbatlas_cluster resource must be set to true. See the example above for how to refer to the mongodbatlas_cluster resource for policies.#.id #### Policy Item`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item",
					Description: `(Required) Array of backup policy items.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.id",
					Description: `(Required) Unique identifier of the backup policy item. ` + "`" + `policies.#.policy_item.#.id` + "`" + ` is a value obtained via the mongodbatlas_cluster resource. provider_backup_enabled of the mongodbatlas_cluster resource must be set to true. See the example above for how to refer to the mongodbatlas_cluster resource for policies.#.policy_item.#.id`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.frequency_interval",
					Description: `(Required) Desired frequency of the new backup policy item specified by frequencyType.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.frequency_type",
					Description: `(Required) Frequency associated with the backup policy item. One of the following values: hourly, daily, weekly or monthly.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.retention_unit",
					Description: `(Required) Scope of the backup policy item: days, weeks, or months.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.retention_value",
					Description: `(Required) Value to associate with retentionUnit. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Unique identifier of the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "next_snapshot",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch when Atlas takes the next snapshot. ## Import Cloud Backup Snapshot Policy entries can be imported using project project_id and cluster_name, in the format ` + "`" + `PROJECTID-CLUSTERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_provider_snapshot_backup_policy.test 5d0f1f73cf09a29120e173cf-MyClusterTest ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/schedule/modify-one-schedule/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Unique identifier of the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "next_snapshot",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch when Atlas takes the next snapshot. ## Import Cloud Backup Snapshot Policy entries can be imported using project project_id and cluster_name, in the format ` + "`" + `PROJECTID-CLUSTERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_provider_snapshot_backup_policy.test 5d0f1f73cf09a29120e173cf-MyClusterTest ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/schedule/modify-one-schedule/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot_restore_job",
			Category:         "Resources",
			ShortDescription: `Provides a Cloud Backup Snapshot Restore Job resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"provider",
				"snapshot",
				"restore",
				"job",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster whose snapshot you want to restore.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Atlas cluster whose snapshot you want to restore.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Required) Unique identifier of the snapshot to restore.`,
				},
				resource.Attribute{
					Name:        "delivery_type",
					Description: `(Required) Type of restore job to create. Possible values are:`,
				},
				resource.Attribute{
					Name:        "target_cluster_name",
					Description: `(Required) Name of the target Atlas cluster to which the restore job restores the snapshot. Only required if deliveryType is automated.`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `(Required) Unique ID of the target Atlas project for the specified targetClusterName. Only required if deliveryType is automated. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "snapshot_restore_job_id",
					Description: `The unique identifier of the restore job.`,
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
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `One or more links to sub-resources and/or related resources. The relations between URLs are explained in the Web Linking Specification.`,
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
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch from which to you want to restore this snapshot. Three conditions apply to this parameter:`,
				},
				resource.Attribute{
					Name:        "oplogInc",
					Description: `Oplog operation number from which to you want to restore this snapshot. This is the second part of an Oplog timestamp. Three conditions apply to this parameter:`,
				},
				resource.Attribute{
					Name:        "pointInTimeUTCSeconds",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch from which you want to restore this snapshot. Two conditions apply to this parameter:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_restore_job_id",
					Description: `The unique identifier of the restore job.`,
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
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `One or more links to sub-resources and/or related resources. The relations between URLs are explained in the Web Linking Specification.`,
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
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch from which to you want to restore this snapshot. Three conditions apply to this parameter:`,
				},
				resource.Attribute{
					Name:        "oplogInc",
					Description: `Oplog operation number from which to you want to restore this snapshot. This is the second part of an Oplog timestamp. Three conditions apply to this parameter:`,
				},
				resource.Attribute{
					Name:        "pointInTimeUTCSeconds",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch from which you want to restore this snapshot. Two conditions apply to this parameter:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cluster",
			Category:         "Resources",
			ShortDescription: `Provides a Cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) Cloud service provider on which the servers are provisioned. The possible values are: - ` + "`" + `AWS` + "`" + ` - Amazon AWS - ` + "`" + `GCP` + "`" + ` - Google Cloud Platform - ` + "`" + `AZURE` + "`" + ` - Microsoft Azure - ` + "`" + `TENANT` + "`" + ` - A multi-tenant deployment on one of the supported cloud service providers. Only valid when providerSettings.instanceSizeName is either M2 or M5.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the cluster as it appears in Atlas. Once the cluster is created, its name cannot be changed.`,
				},
				resource.Attribute{
					Name:        "provider_instance_size_name",
					Description: `(Required) Atlas provides different instance sizes, each with a default storage capacity and RAM size. The instance size you select is used for all the data-bearing servers in your cluster. See [Create a Cluster](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/) ` + "`" + `providerSettings.instanceSizeName` + "`" + ` for valid values and default resources.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_disk_gb_enabled",
					Description: `(Optional) Specifies whether disk auto-scaling is enabled. The default is true. - Set to ` + "`" + `true` + "`" + ` to enable disk auto-scaling. - Set to ` + "`" + `false` + "`" + ` to disable disk auto-scaling.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_enabled",
					Description: `(Optional) Specifies whether cluster tier auto-scaling is enabled. The default is false. - Set to ` + "`" + `true` + "`" + ` to enable cluster tier auto-scaling. If enabled, you must specify a value for ` + "`" + `providerSettings.autoScaling.compute.maxInstanceSize` + "`" + `. - Set to ` + "`" + `false` + "`" + ` to disable cluster tier auto-scaling. ~>`,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_scale_down_enabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable the cluster tier to scale down. This option is only available if ` + "`" + `autoScaling.compute.enabled` + "`" + ` is ` + "`" + `true` + "`" + `. - If this option is enabled, you must specify a value for ` + "`" + `providerSettings.autoScaling.compute.minInstanceSize` + "`" + ``,
				},
				resource.Attribute{
					Name:        "backup_enabled",
					Description: `(Optional) Legacy Backup - Set to true to enable Atlas legacy backups for the cluster.`,
				},
				resource.Attribute{
					Name:        "bi_connector",
					Description: `(Optional) Specifies BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) Specifies the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment. ->`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `(Optional - GCP/AWS Only) Capacity, in gigabytes, of the host’s root volume. Increase this number to add capacity, up to a maximum possible value of 4096 (i.e., 4 TB). This value must be a positive integer.`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `(Optional) Possible values are AWS, GCP, AZURE or NONE. Only needed if you desire to manage the keys, see [Encryption at Rest using Customer Key Management](https://docs.atlas.mongodb.com/security-aws-kms/) for complete documentation. You must configure encryption at rest for the Atlas project before enabling it on any cluster in the project. For complete documentation on configuring Encryption at Rest, see Encryption at Rest using Customer Key Management. Requires M10 or greater. and for legacy backups, backup_enabled, to be false or omitted.`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `(Optional) Version of the cluster to deploy. Atlas supports the following MongoDB versions for M10+ clusters: ` + "`" + `3.6` + "`" + `, ` + "`" + `4.0` + "`" + `, or ` + "`" + `4.2` + "`" + `. You must set this value to ` + "`" + `4.2` + "`" + ` if ` + "`" + `provider_instance_size_name` + "`" + ` is either M2 or M5.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `(Optional) Selects whether the cluster is a replica set or a sharded cluster. If you use the replicationSpecs parameter, you must set num_shards.`,
				},
				resource.Attribute{
					Name:        "pit_enabled",
					Description: `(Optional) - Flag that indicates if the cluster uses Continuous Cloud Backup. If set to true, provider_backup_enabled must also be set to true.`,
				},
				resource.Attribute{
					Name:        "provider_backup_enabled",
					Description: `(Optional) Flag indicating if the cluster uses Cloud Backup for backups. If true, the cluster uses Cloud Backup for backups. If provider_backup_enabled and backup_enabled are false, the cluster does not use Atlas backups. You cannot enable cloud backup if you have an existing cluster in the project with legacy backup enabled. ~>`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `(Optional) Cloud service provider on which the server for a multi-tenant cluster is provisioned. This setting is only valid when providerSetting.providerName is TENANT and providerSetting.instanceSizeName is M2 or M5. The possible values are: - AWS - Amazon AWS - GCP - Google Cloud Platform - AZURE - Microsoft Azure`,
				},
				resource.Attribute{
					Name:        "provider_disk_iops",
					Description: `(Optional) The maximum input/output operations per second (IOPS) the system can perform. The possible values depend on the selected ` + "`" + `provider_instance_size_name` + "`" + ` and ` + "`" + `disk_size_gb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provider_disk_type_name",
					Description: `(Optional - Azure Only) Azure disk type of the server’s root volume. If omitted, Atlas uses the default disk type for the selected providerSettings.instanceSizeName. Example disk types and associated storage sizes: P4 - 32GB, P6 - 64GB, P10 - 128GB, P15 - 256GB, P20 - 512GB, P30 - 1024GB, P40 - 2048GB, P50 - 4095GB. More information and the most update to date disk types/storage sizes can be located at https://docs.atlas.mongodb.com/reference/api/clusters-create-one/.`,
				},
				resource.Attribute{
					Name:        "provider_region_name",
					Description: `(Optional) Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. Requires the`,
				},
				resource.Attribute{
					Name:        "provider_volume_type",
					Description: `(AWS - Optional) The type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `. ` + "`" + `PROVISIONED` + "`" + ` required if setting IOPS higher than the default instance IOPS.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `(Deprecated) Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_min_instance_size",
					Description: `(Optional) Minimum instance size to which your cluster can automatically scale (e.g., M10). Required if ` + "`" + `autoScaling.compute.scaleDownEnabled` + "`" + ` is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_max_instance_size",
					Description: `(Optional) Maximum instance size to which your cluster can automatically scale (e.g., M40). Required if ` + "`" + `autoScaling.compute.enabled` + "`" + ` is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details. ### Multi-Region Cluster ` + "`" + `` + "`" + `` + "`" + `hcl //Example 3 Multi-Region block replication_specs { num_shards = 1 regions_config { region_name = "US_EAST_1" electable_nodes = 3 priority = 7 read_only_nodes = 0 } regions_config { region_name = "US_EAST_2" electable_nodes = 2 priority = 6 read_only_nodes = 0 } regions_config { region_name = "US_WEST_1" electable_nodes = 2 priority = 5 read_only_nodes = 2 } } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `(Required) Number of shards to deploy in the specified zone, minimum 1.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifer of the replication document for a zone in a Global Cluster.`,
				},
				resource.Attribute{
					Name:        "regions_config",
					Description: `(Optional) Physical location of the region. Each regionsConfig document describes the region’s priority in elections and the number and type of MongoDB nodes Atlas deploys to the region. You must order each regionsConfigs document by regionsConfig.priority, descending. See [Region Config](#region-config) below for more details.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Optional) Name for the zone in a Global Cluster.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Optional) Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. Requires the`,
				},
				resource.Attribute{
					Name:        "electable_nodes",
					Description: `(Optional) Number of electable nodes for Atlas to deploy to the region. Electable nodes can become the primary and can facilitate local reads.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Election priority of the region. For regions with only read-only nodes, set this value to 0.`,
				},
				resource.Attribute{
					Name:        "read_only_nodes",
					Description: `(Optional) Number of read-only nodes for Atlas to deploy to the region. Read-only nodes can never become the primary, but can facilitate local-reads. Specify 0 if you do not want any read-only nodes in the region.`,
				},
				resource.Attribute{
					Name:        "analytics_nodes",
					Description: `(Optional) The number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. If you do not specify this option, no analytics nodes are deployed to the region. ### BI Connector Specifies BI Connector for Atlas configuration. ` + "`" + `` + "`" + `` + "`" + `hcl bi_connector = { enabled = true read_preference = secondary } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Specifies whether or not BI Connector for Atlas is enabled on the cluster.l`,
				},
				resource.Attribute{
					Name:        "read_preference",
					Description: `(Optional) Specifies the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences). - Set to "primary" to have BI Connector for Atlas read from the primary. - Set to "secondary" to have BI Connector for Atlas read from a secondary member. Default if there are no analytics nodes in the cluster. - Set to "analytics" to have BI Connector for Atlas read from an analytics node. Default if the cluster contains analytics nodes. ### Advanced Configuration Options ->`,
				},
				resource.Attribute{
					Name:        "fail_index_key_too_long",
					Description: `(Optional) When true, documents can only be updated or inserted if, for all indexed fields on the target collection, the corresponding index entries do not exceed 1024 bytes. When false, mongod writes documents that exceed the limit but does not index them.`,
				},
				resource.Attribute{
					Name:        "javascript_enabled",
					Description: `(Optional) When true, the cluster allows execution of operations that perform server-side executions of JavaScript. When false, the cluster disables execution of those operations.`,
				},
				resource.Attribute{
					Name:        "minimum_enabled_tls_protocol",
					Description: `(Optional) Sets the minimum Transport Layer Security (TLS) version the cluster accepts for incoming connections.Valid values are: - TLS1_0 - TLS1_1 - TLS1_2`,
				},
				resource.Attribute{
					Name:        "no_table_scan",
					Description: `(Optional) When true, the cluster disables the execution of any query that requires a collection scan to return results. When false, the cluster allows the execution of those operations.`,
				},
				resource.Attribute{
					Name:        "oplog_size_mb",
					Description: `(Optional) The custom oplog size of the cluster. Without a value that indicates that the cluster uses the default oplog size calculated by Atlas.`,
				},
				resource.Attribute{
					Name:        "sample_size_bi_connector",
					Description: `(Optional) Number of documents per database to sample when gathering schema information. Defaults to 100. Available only for Atlas deployments in which BI Connector for Atlas is enabled.`,
				},
				resource.Attribute{
					Name:        "sample_refresh_interval_bi_connector",
					Description: `(Optional) Interval in seconds at which the mongosqld process re-samples data to create its relational schema. The default value is 300. The specified value must be a positive integer. Available only for Atlas deployments in which BI Connector for Atlas is enabled. ### Labels ` + "`" + `` + "`" + `` + "`" + `hcl labels { key = "Key 1" value = "Value 1" } labels { key = "Key 2" value = "Value 2" } ` + "`" + `` + "`" + `` + "`" + ` Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. You cannot set the key ` + "`" + `Infrastructure Tool` + "`" + `, it is used for internal purposes to track aggregate usage.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The cluster ID.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
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
					Description: `connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster.`,
				},
				resource.Attribute{
					Name:        "connection_strings",
					Description: `Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `The Network Peering Container ID. The id of the container either created programmatically by the user before any clusters existed in the project or when the first cluster in the region (AWS/Azure) or project (GCP) was created.`,
				},
				resource.Attribute{
					Name:        "paused",
					Description: `Flag that indicates whether the cluster is paused or not.`,
				},
				resource.Attribute{
					Name:        "srv_address",
					Description: `Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS/SSL. See the mongoURI for additional options.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING ### Cloud Backup Policy Cloud Backup Policy will be added if provider_backup_enabled is enabled because MongoDB Atlas automatically creates a default policy, if not, returned values will be empty.`,
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
					Description: `The number of days, weeks, or months the snapshot is retained. ## Import Clusters can be imported using project ID and cluster name, in the format ` + "`" + `PROJECTID-CLUSTERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cluster.my_cluster 1112222b3bf99403840e8934-Cluster0 ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The cluster ID.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
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
					Description: `connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster.`,
				},
				resource.Attribute{
					Name:        "connection_strings",
					Description: `Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `The Network Peering Container ID. The id of the container either created programmatically by the user before any clusters existed in the project or when the first cluster in the region (AWS/Azure) or project (GCP) was created.`,
				},
				resource.Attribute{
					Name:        "paused",
					Description: `Flag that indicates whether the cluster is paused or not.`,
				},
				resource.Attribute{
					Name:        "srv_address",
					Description: `Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS/SSL. See the mongoURI for additional options.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING ### Cloud Backup Policy Cloud Backup Policy will be added if provider_backup_enabled is enabled because MongoDB Atlas automatically creates a default policy, if not, returned values will be empty.`,
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
					Description: `The number of days, weeks, or months the snapshot is retained. ## Import Clusters can be imported using project ID and cluster name, in the format ` + "`" + `PROJECTID-CLUSTERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cluster.my_cluster 1112222b3bf99403840e8934-Cluster0 ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_custom_db_role",
			Category:         "Resources",
			ShortDescription: `Provides a Custom DB Role resource.`,
			Description:      ``,
			Keywords: []string{
				"custom",
				"db",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) Name of the custom role. ->`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Name of the privilege action. For a complete list of actions available in the Atlas API, see [Custom Role Actions](https://docs.atlas.mongodb.com/reference/api/custom-role-actions) ->`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Contains information on where the action is granted. Each object in the array either indicates a database and collection on which the action is granted, or indicates that the action is granted on the cluster resource.`,
				},
				resource.Attribute{
					Name:        "resources.#.collection",
					Description: `(Optional) Collection on which the action is granted. If this value is an empty string, the action is granted on all collections within the database specified in the actions.resources.db field. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages and can be used to import. ## Import Database users can be imported using project ID and username, in the format ` + "`" + `PROJECTID-ROLENAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_custom_db_role.my_role 1112222b3bf99403840e8934-MyCustomRole ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/custom-roles/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages and can be used to import. ## Import Database users can be imported using project ID and username, in the format ` + "`" + `PROJECTID-ROLENAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_custom_db_role.my_role 1112222b3bf99403840e8934-MyCustomRole ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/custom-roles/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_database_user",
			Category:         "Resources",
			ShortDescription: `Provides a Database User resource.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auth_database_name",
					Description: `(Required) Database against which Atlas authenticates the user. A user must provide both a username and authentication database to log into MongoDB. Accepted values include:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See [Roles](#roles) below for more details.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username for authenticating to MongoDB.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) User's initial password. A value is required to create the database user, however the argument but may be removed from your Terraform configuration after user creation without impacting the user, password or Terraform management. IMPORTANT --- Passwords may show up in Terraform related logs and it will be stored in the Terraform state file as plain-text. Password can be changed after creation using your preferred method, e.g. via the MongoDB Atlas UI, to ensure security. If you do change management of the password to outside of Terraform be sure to remove the argument from the Terraform configuration so it is not inadvertently updated to the original password.`,
				},
				resource.Attribute{
					Name:        "x509_type",
					Description: `(Optional) X.509 method by which the provided username is authenticated. If no value is given, Atlas uses the default value of NONE. The accepted types are:`,
				},
				resource.Attribute{
					Name:        "NONE",
					Description: `The user does not use X.509 authentication.`,
				},
				resource.Attribute{
					Name:        "MANAGED",
					Description: `The user is being created for use with Atlas-managed X.509.Externally authenticated users can only be created on the ` + "`" + `$external` + "`" + ` database.`,
				},
				resource.Attribute{
					Name:        "CUSTOMER",
					Description: `The user is being created for use with Self-Managed X.509. Users created with this x509Type require a Common Name (CN) in the username field. Externally authenticated users can only be created on the ` + "`" + `$external` + "`" + ` database.`,
				},
				resource.Attribute{
					Name:        "aws_iam_type",
					Description: `(Optional) If this value is set, the new database user authenticates with AWS IAM credentials. If no value is given, Atlas uses the default value of NONE. The accepted types are:`,
				},
				resource.Attribute{
					Name:        "NONE",
					Description: `The user does not use AWS IAM credentials.`,
				},
				resource.Attribute{
					Name:        "USER",
					Description: `New database user has AWS IAM user credentials.`,
				},
				resource.Attribute{
					Name:        "ROLE",
					Description: `New database user has credentials associated with an AWS IAM role.`,
				},
				resource.Attribute{
					Name:        "ldap_auth_type",
					Description: `(Optional) Method by which the provided ` + "`" + `username` + "`" + ` is authenticated. If no value is given, Atlas uses the default value of ` + "`" + `NONE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "NONE",
					Description: `Atlas authenticates this user through [SCRAM-SHA](https://docs.mongodb.com/manual/core/security-scram/), not LDAP.`,
				},
				resource.Attribute{
					Name:        "USER",
					Description: `LDAP server authenticates this user through the user's LDAP user. ` + "`" + `username` + "`" + ` must also be a fully qualified distinguished name, as defined in [RFC-2253](https://tools.ietf.org/html/rfc2253).`,
				},
				resource.Attribute{
					Name:        "GROUP",
					Description: `LDAP server authenticates this user using their LDAP user and authorizes this user using their LDAP group. To learn more about LDAP security, see [Set up User Authentication and Authorization with LDAP](https://docs.atlas.mongodb.com/security-ldaps). ` + "`" + `username` + "`" + ` must also be a fully qualified distinguished name, as defined in [RFC-2253](https://tools.ietf.org/html/rfc2253). ### Roles Block mapping a user's role to a database / collection. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. ->`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) Name of the role to grant. See [Create a Database User](https://docs.atlas.mongodb.com/reference/api/database-users-create-a-user/) ` + "`" + `roles.roleName` + "`" + ` for valid values and restrictions.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) Database on which the user has the specified role. A role on the ` + "`" + `admin` + "`" + ` database can include privileges that apply to the other databases.`,
				},
				resource.Attribute{
					Name:        "collection_name",
					Description: `(Optional) Collection for which the role applies. You can specify a collection for the ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + ` roles. If you do not specify a collection for ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + `, the role applies to all collections in the database (excluding some collections in the ` + "`" + `system` + "`" + `. database). ### Labels Containing key-value pairs that tag and categorize the database user. Each key and value has a maximum length of 255 characters.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. ### Scopes Array of clusters and Atlas Data Lakes that this user has access to. If omitted, Atlas grants the user access to all the clusters and Atlas Data Lakes in the project by default.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the cluster or Atlas Data Lake that the user has access to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of resource that the user has access to. Valid values are: ` + "`" + `CLUSTER` + "`" + ` and ` + "`" + `DATA_LAKE` + "`" + ` ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The database user's name. ## Import Database users can be imported using project ID and username, in the format ` + "`" + `project_id` + "`" + `-` + "`" + `username` + "`" + `-` + "`" + `auth_database_name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_database_user.my_user 1112222b3bf99403840e8934-my_user-admin ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The database user's name. ## Import Database users can be imported using project ID and username, in the format ` + "`" + `project_id` + "`" + `-` + "`" + `username` + "`" + `-` + "`" + `auth_database_name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_database_user.my_user 1112222b3bf99403840e8934-my_user-admin ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_encryption_at_rest",
			Category:         "Resources",
			ShortDescription: `Provides an Encryption At Rest resource.`,
			Description:      ``,
			Keywords: []string{
				"encryption",
				"at",
				"rest",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique identifier for the project.`,
				},
				resource.Attribute{
					Name:        "aws_kms",
					Description: `(Required) Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.`,
				},
				resource.Attribute{
					Name:        "azure_key_vault",
					Description: `(Required) Specifies Azure Key Vault configuration details and whether Encryption at Rest is enabled for an Atlas project.`,
				},
				resource.Attribute{
					Name:        "google_cloud_kms",
					Description: `(Required) Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project. ### aws_kms`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Specifies whether Encryption at Rest is enabled for an Atlas project, To disable Encryption at Rest, pass only this parameter with a value of false, When you disable Encryption at Rest, Atlas also removes the configuration details.`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `The IAM access key ID with permissions to access the customer master key specified by customerMasterKeyID.`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `The IAM secret access key with permissions to access the customer master key specified by customerMasterKeyID.`,
				},
				resource.Attribute{
					Name:        "customer_master_key_id",
					Description: `The AWS customer master key used to encrypt and decrypt the MongoDB master keys.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The AWS region in which the AWS customer master key exists: CA_CENTRAL_1, US_EAST_1, US_EAST_2, US_WEST_1, US_WEST_2, SA_EAST_1`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `ID of an AWS IAM role authorized to manage an AWS customer master key. To find the ID for an existing IAM role check the ` + "`" + `role_id` + "`" + ` attribute of the ` + "`" + `mongodbatlas_cloud_provider_access` + "`" + ` resource. ### azure_key_vault`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The client ID, also known as the application ID, for an Azure application associated with the Azure AD tenant.`,
				},
				resource.Attribute{
					Name:        "azure_environment",
					Description: `The Azure environment where the Azure account credentials reside. Valid values are the following: AZURE, AZURE_CHINA, AZURE_GERMANY`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `The unique identifier associated with an Azure subscription.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `The name of the Azure Resource group that contains an Azure Key Vault.`,
				},
				resource.Attribute{
					Name:        "key_vault_name",
					Description: `The name of an Azure Key Vault containing your key.`,
				},
				resource.Attribute{
					Name:        "key_identifier",
					Description: `The unique identifier of a key in an Azure Key Vault.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `The secret associated with the Azure Key Vault specified by azureKeyVault.tenantID.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The unique identifier for an Azure AD tenant within an Azure subscription. ### google_cloud_kms`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Specifies whether Encryption at Rest is enabled for an Atlas project. To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.`,
				},
				resource.Attribute{
					Name:        "service_account_key",
					Description: `String-formatted JSON object containing GCP KMS credentials from your GCP account.`,
				},
				resource.Attribute{
					Name:        "key_version_resource_id",
					Description: `The Key Version Resource ID from your GCP account. For more information see: [MongoDB Atlas API Reference for Encryption at Rest using Customer Key Management.](https://docs.atlas.mongodb.com/reference/api/encryption-at-rest/)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_global_cluster_config",
			Category:         "Resources",
			ShortDescription: `Provides a Global Cluster Configuration resource.`,
			Description:      ``,
			Keywords: []string{
				"global",
				"cluster",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "collection",
					Description: `(Required) The name of the collection associated with the managed namespace.`,
				},
				resource.Attribute{
					Name:        "custom_shard_key",
					Description: `(Required) The custom shard key for the collection. Global Clusters require a compound shard key consisting of a location field and a user-selected second key, the custom shard key.`,
				},
				resource.Attribute{
					Name:        "db",
					Description: `(Required) The name of the database containing the collection. ### Custom Zone Mapping`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The ISO location code to which you want to map a zone in your Global Cluster. You can find a list of all supported location codes [here](https://cloud.mongodb.com/static/atlas/country_iso_codes.txt).`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name of the zone in your Global Cluster that you want to map to location. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "custom_zone_mapping",
					Description: `A map of all custom zone mappings defined for the Global Cluster. Atlas automatically maps each location code to the closest geographical zone. Custom zone mappings allow administrators to override these automatic mappings. If your Global Cluster does not have any custom zone mappings, this document is empty. ## Import Database users can be imported using project ID and cluster name, in the format ` + "`" + `PROJECTID-CLUSTER_NAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_global_cluster_config.config 1112222b3bf99403840e8934-my-cluster ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Global Clusters](https://docs.atlas.mongodb.com/reference/api/global-clusters/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "custom_zone_mapping",
					Description: `A map of all custom zone mappings defined for the Global Cluster. Atlas automatically maps each location code to the closest geographical zone. Custom zone mappings allow administrators to override these automatic mappings. If your Global Cluster does not have any custom zone mappings, this document is empty. ## Import Database users can be imported using project ID and cluster name, in the format ` + "`" + `PROJECTID-CLUSTER_NAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_global_cluster_config.config 1112222b3bf99403840e8934-my-cluster ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Global Clusters](https://docs.atlas.mongodb.com/reference/api/global-clusters/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_maintenance_window",
			Category:         "Resources",
			ShortDescription: `Provides an Maintenance Window resource.`,
			Description:      ``,
			Keywords: []string{
				"maintenance",
				"window",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `The unique identifier of the project for the Maintenance Window.`,
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
					Description: `Number of times the current maintenance event for this project has been deferred, you can set a maximum of 2 deferrals.`,
				},
				resource.Attribute{
					Name:        "defer",
					Description: `Defer maintenance for the given project for one week. ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_container",
			Category:         "Resources",
			ShortDescription: `Provides a Network Peering resource.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"container",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique identifier for the Atlas project for this Network Peering Container.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `(Required) CIDR block that Atlas uses for the Network Peering containers in your project. Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3):`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Required AWS only) The Atlas AWS region name for where this container will exist, see the reference list for Atlas AWS region names [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required AZURE only) Atlas region where the container resides, see the reference list for Atlas Azure region names [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Optional GCP only) Atlas regions where the container resides. Provide this field only if you provide an ` + "`" + `atlas_cidr_block` + "`" + ` smaller than ` + "`" + `/18` + "`" + `. [GCP Regions values](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/#request-body-parameters). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `The Network Peering Container ID.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "provisioned",
					Description: `Indicates whether the project has Network Peering connections deployed in the container.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Atlas name for AWS region where the Atlas container resides.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Unique identifier of Atlas' AWS VPC.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `Unique identifier of the GCP project in which the network peer resides. Returns null. This value is populated once you create a new network peering connection with the network peering resource.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Unique identifier of the Network Peering connection in the Atlas project. Returns null. This value is populated once you create a new network peering connection with the network peering resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Azure region where the Atlas container resides.`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `Unique identifier of the Azure subscription in which the VNet resides.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `The name of the Azure VNet. Returns null. This value is populated once you create a new network peering connection with the network peering resource. ## Import Clusters can be imported using project ID and network peering container id, in the format ` + "`" + `PROJECTID-CONTAINER-ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_network_container.my_container 1112222b3bf99403840e8934-5cbf563d87d9d67253be590a ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "container_id",
					Description: `The Network Peering Container ID.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "provisioned",
					Description: `Indicates whether the project has Network Peering connections deployed in the container.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Atlas name for AWS region where the Atlas container resides.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Unique identifier of Atlas' AWS VPC.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `Unique identifier of the GCP project in which the network peer resides. Returns null. This value is populated once you create a new network peering connection with the network peering resource.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Unique identifier of the Network Peering connection in the Atlas project. Returns null. This value is populated once you create a new network peering connection with the network peering resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Azure region where the Atlas container resides.`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `Unique identifier of the Azure subscription in which the VNet resides.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `The name of the Azure VNet. Returns null. This value is populated once you create a new network peering connection with the network peering resource. ## Import Clusters can be imported using project ID and network peering container id, in the format ` + "`" + `PROJECTID-CONTAINER-ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_network_container.my_container 1112222b3bf99403840e8934-5cbf563d87d9d67253be590a ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_peering",
			Category:         "Resources",
			ShortDescription: `Provides a Network Peering resource.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"peering",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the MongoDB Atlas project to create the database user.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `(Required) Unique identifier of the MongoDB Atlas container for the provider (GCP) or provider/region (AWS, AZURE). You can create an MongoDB Atlas container using the network_container resource or it can be obtained from the cluster returned values if a cluster has been created before the first container.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) Cloud provider to whom the peering connection is being made. (Possible Values ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + `, ` + "`" + `GCP` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "accepter_region_name",
					Description: `(Required - AWS) Specifies the AWS region where the peer VPC resides. For complete lists of supported regions, see [Amazon Web Services](https://docs.atlas.mongodb.com/reference/amazon-aws/).`,
				},
				resource.Attribute{
					Name:        "aws_account_id",
					Description: `(Required - AWS) AWS Account ID of the owner of the peer VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Unique identifier of the AWS peer VPC (Note: this is`,
				},
				resource.Attribute{
					Name:        "route_table_cidr_block",
					Description: `(Required - AWS) AWS VPC CIDR block or subnet.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `(Required - GCP) GCP project ID of the owner of the network peer.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `(Required- AWS) Name of the network peer to which Atlas connects.`,
				},
				resource.Attribute{
					Name:        "azure_directory_id",
					Description: `(Required - AZURE) Unique identifier for an Azure AD directory.`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `(Required - AZURE) Unique identifier of the Azure subscription in which the VNet resides.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required - AZURE) Name of your Azure resource group.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `(Required - AZURE) Name of your Azure VNet. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "peer_id",
					Description: `Unique identifier of the Atlas network peer.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "connection_id",
					Description: `Unique identifier of the Atlas network peering container.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Cloud provider to whom the peering connection is being made. (Possible Values ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + `, ` + "`" + `GCP` + "`" + `).`,
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
					Description: `(AWS Only) The VPC peering connection status value can be one of the following: ` + "`" + `INITIATING` + "`" + `, ` + "`" + `PENDING_ACCEPTANCE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `FINALIZING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `TERMINATING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Atlas network peering connection. Azure/GCP: ` + "`" + `ADDING_PEER` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `DELETING` + "`" + ` GCP Only: ` + "`" + `WAITING_FOR_USER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `GCP project ID of the owner of the network peer.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `When ` + "`" + `"status" : "FAILED"` + "`" + `, Atlas provides a description of the error.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the network peer to which Atlas connects.`,
				},
				resource.Attribute{
					Name:        "atlas_gcp_project_id",
					Description: `The Atlas GCP Project ID for the GCP VPC used by your atlas cluster that it is need to set up the reciprocal connection.`,
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
					Name:        "error_state",
					Description: `Description of the Atlas error when ` + "`" + `status` + "`" + ` is ` + "`" + `Failed` + "`" + `, Otherwise, Atlas returns ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Name of your Azure resource group.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `Name of your Azure VNet. ## Import Clusters can be imported using project ID and network peering peering id, in the format ` + "`" + `PROJECTID-PEERID-PROVIDERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_network_peering.my_peering 1112222b3bf99403840e8934-5cbf563d87d9d67253be590a-AWS ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Network Peering Connection](https://docs.atlas.mongodb.com/reference/api/vpc-create-peering-connection/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "peer_id",
					Description: `Unique identifier of the Atlas network peer.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "connection_id",
					Description: `Unique identifier of the Atlas network peering container.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Cloud provider to whom the peering connection is being made. (Possible Values ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + `, ` + "`" + `GCP` + "`" + `).`,
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
					Description: `(AWS Only) The VPC peering connection status value can be one of the following: ` + "`" + `INITIATING` + "`" + `, ` + "`" + `PENDING_ACCEPTANCE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `FINALIZING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `TERMINATING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Atlas network peering connection. Azure/GCP: ` + "`" + `ADDING_PEER` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `DELETING` + "`" + ` GCP Only: ` + "`" + `WAITING_FOR_USER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `GCP project ID of the owner of the network peer.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `When ` + "`" + `"status" : "FAILED"` + "`" + `, Atlas provides a description of the error.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the network peer to which Atlas connects.`,
				},
				resource.Attribute{
					Name:        "atlas_gcp_project_id",
					Description: `The Atlas GCP Project ID for the GCP VPC used by your atlas cluster that it is need to set up the reciprocal connection.`,
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
					Name:        "error_state",
					Description: `Description of the Atlas error when ` + "`" + `status` + "`" + ` is ` + "`" + `Failed` + "`" + `, Otherwise, Atlas returns ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `Name of your Azure resource group.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `Name of your Azure VNet. ## Import Clusters can be imported using project ID and network peering peering id, in the format ` + "`" + `PROJECTID-PEERID-PROVIDERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_network_peering.my_peering 1112222b3bf99403840e8934-5cbf563d87d9d67253be590a-AWS ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Network Peering Connection](https://docs.atlas.mongodb.com/reference/api/vpc-create-peering-connection/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_private_endpoint",
			Category:         "Resources",
			ShortDescription: `Provides a Private Endpoint resource.`,
			Description:      ``,
			Keywords: []string{
				"private",
				"endpoint",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `Required Unique identifier for the project.`,
				},
				resource.Attribute{
					Name:        "providerName",
					Description: `(Required) Name of the cloud provider you want to create the private endpoint connection for. Must be AWS.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Cloud provider region in which you want to create the private endpoint connection. Accepted values are:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "private_link_id",
					Description: `Unique identifier of the AWS PrivateLink connection.`,
				},
				resource.Attribute{
					Name:        "endpoint_service_name",
					Description: `Name of the PrivateLink endpoint service in AWS. Returns null while the endpoint service is being created.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `Error message pertaining to the AWS PrivateLink connection. Returns null if there are no errors.`,
				},
				resource.Attribute{
					Name:        "interface_endpoints",
					Description: `Unique identifiers of the interface endpoints in your VPC that you added to the AWS PrivateLink connection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the AWS PrivateLink connection. Returns one of the following values:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "private_link_id",
					Description: `Unique identifier of the AWS PrivateLink connection.`,
				},
				resource.Attribute{
					Name:        "endpoint_service_name",
					Description: `Name of the PrivateLink endpoint service in AWS. Returns null while the endpoint service is being created.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `Error message pertaining to the AWS PrivateLink connection. Returns null if there are no errors.`,
				},
				resource.Attribute{
					Name:        "interface_endpoints",
					Description: `Unique identifiers of the interface endpoints in your VPC that you added to the AWS PrivateLink connection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the AWS PrivateLink connection. Returns one of the following values:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_private_endpoint_interface_link",
			Category:         "Resources",
			ShortDescription: `Provides a Private Endpoint Link resource.`,
			Description:      ``,
			Keywords: []string{
				"private",
				"endpoint",
				"interface",
				"link",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique identifier for the project.`,
				},
				resource.Attribute{
					Name:        "private_link_id",
					Description: `(Required) Unique identifier of the AWS PrivateLink connection which is created by ` + "`" + `mongodbatlas_private_endpoint` + "`" + ` resource.`,
				},
				resource.Attribute{
					Name:        "interface_endpoint_id",
					Description: `(Required) Unique identifier of the interface endpoint you created in your VPC with the AWS resource. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "delete_requested",
					Description: `Indicates if Atlas received a request to remove the interface endpoint from the private endpoint connection.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `Error message pertaining to the interface endpoint. Returns null if there are no errors.`,
				},
				resource.Attribute{
					Name:        "connection_status",
					Description: `Status of the interface endpoint. Returns one of the following values:`,
				},
				resource.Attribute{
					Name:        "NONE",
					Description: `Atlas created the network load balancer and VPC endpoint service, but AWS hasn’t yet created the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "PENDING_ACCEPTANCE",
					Description: `AWS has received the connection request from your VPC endpoint to the Atlas VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "PENDING",
					Description: `AWS is establishing the connection between your VPC endpoint and the Atlas VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "AVAILABLE",
					Description: `Atlas VPC resources are connected to the VPC endpoint in your VPC. You can connect to Atlas clusters in this region using AWS PrivateLink.`,
				},
				resource.Attribute{
					Name:        "REJECTED",
					Description: `AWS failed to establish a connection between Atlas VPC resources to the VPC endpoint in your VPC.`,
				},
				resource.Attribute{
					Name:        "DELETING",
					Description: `Atlas is removing the interface endpoint from the private endpoint connection. ## Import Private Endpoint Link Connection can be imported using project ID and username, in the format ` + "`" + `{project_id}-{private_link_id}-{interface_endpoint_id}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_private_endpoint_interface_link.test 1112222b3bf99403840e8934-3242342343112-vpce-4242342343 ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Private Endpoint Link Connection](https://docs.atlas.mongodb.com/reference/api/private-endpoint-create-one-interface-endpoint/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "delete_requested",
					Description: `Indicates if Atlas received a request to remove the interface endpoint from the private endpoint connection.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `Error message pertaining to the interface endpoint. Returns null if there are no errors.`,
				},
				resource.Attribute{
					Name:        "connection_status",
					Description: `Status of the interface endpoint. Returns one of the following values:`,
				},
				resource.Attribute{
					Name:        "NONE",
					Description: `Atlas created the network load balancer and VPC endpoint service, but AWS hasn’t yet created the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "PENDING_ACCEPTANCE",
					Description: `AWS has received the connection request from your VPC endpoint to the Atlas VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "PENDING",
					Description: `AWS is establishing the connection between your VPC endpoint and the Atlas VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "AVAILABLE",
					Description: `Atlas VPC resources are connected to the VPC endpoint in your VPC. You can connect to Atlas clusters in this region using AWS PrivateLink.`,
				},
				resource.Attribute{
					Name:        "REJECTED",
					Description: `AWS failed to establish a connection between Atlas VPC resources to the VPC endpoint in your VPC.`,
				},
				resource.Attribute{
					Name:        "DELETING",
					Description: `Atlas is removing the interface endpoint from the private endpoint connection. ## Import Private Endpoint Link Connection can be imported using project ID and username, in the format ` + "`" + `{project_id}-{private_link_id}-{interface_endpoint_id}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_private_endpoint_interface_link.test 1112222b3bf99403840e8934-3242342343112-vpce-4242342343 ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Private Endpoint Link Connection](https://docs.atlas.mongodb.com/reference/api/private-endpoint-create-one-interface-endpoint/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_private_ip_mode",
			Category:         "Resources",
			ShortDescription: `Provides a Private IP Mode resource.`,
			Description:      ``,
			Keywords: []string{
				"private",
				"ip",
				"mode",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to enable Only Private IP Mode.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Indicates whether Connect via Peering Only mode is enabled or disabled for an Atlas project ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The project id. ## Import Project must be imported using project ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_private_ip_mode.my_private_ip_mode 5d09d6a59ccf6445652a444a ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/get-private-ip-mode-for-project/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The project id. ## Import Project must be imported using project ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_private_ip_mode.my_private_ip_mode 5d09d6a59ccf6445652a444a ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/get-private-ip-mode-for-project/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_project",
			Category:         "Resources",
			ShortDescription: `Provides a Project resource.`,
			Description:      ``,
			Keywords: []string{
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project you want to create. (Cannot be changed via this Provider after creation.)`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) The ID of the organization you want to create the project within. ### Teams Teams attribute is optional ~>`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) The unique identifier of the team you want to associate with the project. The team and project must share the same parent organization.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `(Required) Each string in the array represents a project role you want to assign to the team. Every user associated with the team inherits these roles. You must specify an array even if you are only associating a single role with the team. The following are valid roles:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The project id.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The ISO-8601-formatted timestamp of when Atlas created the project..`,
				},
				resource.Attribute{
					Name:        "cluster_count",
					Description: `The number of Atlas clusters deployed in the project.. ## Import Project must be imported using project ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_project.my_project 5d09d6a59ccf6445652a444a ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/projects/) - [and MongoDB Atlas API - Teams](https://docs.atlas.mongodb.com/reference/api/teams/) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The project id.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The ISO-8601-formatted timestamp of when Atlas created the project..`,
				},
				resource.Attribute{
					Name:        "cluster_count",
					Description: `The number of Atlas clusters deployed in the project.. ## Import Project must be imported using project ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_project.my_project 5d09d6a59ccf6445652a444a ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/projects/) - [and MongoDB Atlas API - Teams](https://docs.atlas.mongodb.com/reference/api/teams/) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_project_ip_whitelist",
			Category:         "Resources",
			ShortDescription: `Provides an IP Whitelist resource.`,
			Description:      ``,
			Keywords: []string{
				"project",
				"ip",
				"whitelist",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project in which to add the whitelist entry.`,
				},
				resource.Attribute{
					Name:        "aws_security_group",
					Description: `(Optional) ID of the whitelisted AWS security group. Mutually exclusive with ` + "`" + `cidr_block` + "`" + ` and ` + "`" + `ip_address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) Whitelist entry in Classless Inter-Domain Routing (CIDR) notation. Mutually exclusive with ` + "`" + `aws_security_group` + "`" + ` and ` + "`" + `ip_address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Whitelisted IP address. Mutually exclusive with ` + "`" + `aws_security_group` + "`" + ` and ` + "`" + `cidr_block` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment to add to the whitelist entry. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages and can be used to import. ## Import IP Whitelist entries can be imported using the ` + "`" + `project_id` + "`" + ` and ` + "`" + `cidr_block` + "`" + ` or ` + "`" + `ip_address` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_project_ip_whitelist.test 5d0f1f74cf09a29120e123cd-10.242.88.0/21 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/whitelist/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages and can be used to import. ## Import IP Whitelist entries can be imported using the ` + "`" + `project_id` + "`" + ` and ` + "`" + `cidr_block` + "`" + ` or ` + "`" + `ip_address` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_project_ip_whitelist.test 5d0f1f74cf09a29120e123cd-10.242.88.0/21 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/whitelist/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_team",
			Category:         "Resources",
			ShortDescription: `Provides a Team resource.`,
			Description:      ``,
			Keywords: []string{
				"team",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) The unique identifier for the organization you want to associate the team with.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the team you want to create.`,
				},
				resource.Attribute{
					Name:        "usernames",
					Description: `(Required) The Atlas usernames (email address). You can only add Atlas users who are part of the organization. Users who have not accepted an invitation to join the organization cannot be added as team members. There is a maximum of 250 Atlas users per team. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The unique identifier for the team. ## Import Teams can be imported using the organization ID and team id, in the format ORGID-TEAMID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_teams.my_team 1112222b3bf99403840e8934-1112222b3bf99403840e8935 ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Teams](https://docs.atlas.mongodb.com/reference/api/teams-create-one/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The unique identifier for the team. ## Import Teams can be imported using the organization ID and team id, in the format ORGID-TEAMID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_teams.my_team 1112222b3bf99403840e8934-1112222b3bf99403840e8935 ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Teams](https://docs.atlas.mongodb.com/reference/api/teams-create-one/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_x509_authentication_database_user",
			Category:         "Resources",
			ShortDescription: `Provides a X509 Authentication Database User resource.`,
			Description:      ``,
			Keywords: []string{
				"x509",
				"authentication",
				"database",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Identifier for the Atlas project associated with the X.509 configuration.`,
				},
				resource.Attribute{
					Name:        "months_until_expiration",
					Description: `(Required) A number of months that the created certificate is valid for before expiry, up to 24 months. By default is 3.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Username of the database user to create a certificate for.`,
				},
				resource.Attribute{
					Name:        "customer_x509_cas",
					Description: `(Optional) PEM string containing one or more customer CAs for database user authentication. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `Fully distinguished name of the database user to which this certificate belongs. To learn more, see RFC 2253. ## Import X.509 Certificates for a User can be imported using project ID and username, in the format ` + "`" + `project_id-username` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_x509_authentication_database_user.test 1112222b3bf99403840e8934-myUsername ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/x509-configuration-get-certificates/) Current X.509 Configuration can be imported using project ID, in the format ` + "`" + `project_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_x509_authentication_database_user.test 1112222b3bf99403840e8934 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/x509-configuration-get-certificates/)`,
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
					Description: `Fully distinguished name of the database user to which this certificate belongs. To learn more, see RFC 2253. ## Import X.509 Certificates for a User can be imported using project ID and username, in the format ` + "`" + `project_id-username` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_x509_authentication_database_user.test 1112222b3bf99403840e8934-myUsername ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/x509-configuration-get-certificates/) Current X.509 Configuration can be imported using project ID, in the format ` + "`" + `project_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_x509_authentication_database_user.test 1112222b3bf99403840e8934 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/x509-configuration-get-certificates/)`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"mongodbatlas_alert_configuration":                   0,
		"mongodbatlas_auditing":                              1,
		"mongodbatlas_cloud_provider_snapshot":               2,
		"mongodbatlas_cloud_provider_snapshot_backup_policy": 3,
		"mongodbatlas_cloud_provider_snapshot_restore_job":   4,
		"mongodbatlas_cluster":                               5,
		"mongodbatlas_custom_db_role":                        6,
		"mongodbatlas_database_user":                         7,
		"mongodbatlas_encryption_at_rest":                    8,
		"mongodbatlas_global_cluster_config":                 9,
		"mongodbatlas_maintenance_window":                    10,
		"mongodbatlas_network_container":                     11,
		"mongodbatlas_network_peering":                       12,
		"mongodbatlas_private_endpoint":                      13,
		"mongodbatlas_private_endpoint_interface_link":       14,
		"mongodbatlas_private_ip_mode":                       15,
		"mongodbatlas_project":                               16,
		"mongodbatlas_project_ip_whitelist":                  17,
		"mongodbatlas_team":                                  18,
		"mongodbatlas_x509_authentication_database_user":     19,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
