package mongodbatlas

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot",
			Category:         "Resources",
			ShortDescription: `Provides an Cloud Provider Snapshot resource.`,
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
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. ## Import Cloud Provider Snapshot entries can be imported using project project_id, cluster_name and snapshot_id (Unique identifier of the snapshot), in the format ` + "`" + `PROJECTID-CLUSTERNAME-SNAPSHOTID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_provider_snapshot.test 5d0f1f73cf09a29120e173cf-MyClusterTest-5d116d82014b764445b2f9b5 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-provider-snapshot/)`,
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
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. ## Import Cloud Provider Snapshot entries can be imported using project project_id, cluster_name and snapshot_id (Unique identifier of the snapshot), in the format ` + "`" + `PROJECTID-CLUSTERNAME-SNAPSHOTID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_provider_snapshot.test 5d0f1f73cf09a29120e173cf-MyClusterTest-5d116d82014b764445b2f9b5 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-provider-snapshot/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot_restore_job",
			Category:         "Resources",
			ShortDescription: `Provides a Cloud Provider Snapshot Restore Job resource.`,
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
					Description: `Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken. ## Import Cloud Provider Snapshot Restore Job entries can be imported using project project_id, cluster_name and snapshot_id (Unique identifier of the snapshot), in the format ` + "`" + `PROJECTID-CLUSTERNAME-JOBID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_provider_snapshot_restore_job.test 5cf5a45a9ccf6400e60981b6-MyCluster-5d1b654ecf09a24b888f4c79 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-provider-snapshot-restore-jobs/)`,
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
					Description: `Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken. ## Import Cloud Provider Snapshot Restore Job entries can be imported using project project_id, cluster_name and snapshot_id (Unique identifier of the snapshot), in the format ` + "`" + `PROJECTID-CLUSTERNAME-JOBID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_provider_snapshot_restore_job.test 5cf5a45a9ccf6400e60981b6-MyCluster-5d1b654ecf09a24b888f4c79 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-provider-snapshot-restore-jobs/)`,
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
					Description: `(Required) Cloud service provider on which the servers are provisioned.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the cluster as it appears in Atlas. Once the cluster is created, its name cannot be changed. The possible values are: - ` + "`" + `AWS` + "`" + ` - Amazon AWS - ` + "`" + `GCP` + "`" + ` - Google Cloud Platform - ` + "`" + `AZURE` + "`" + ` - Microsoft Azure - ` + "`" + `TENANT` + "`" + ` - A multi-tenant deployment on one of the supported cloud service providers. Only valid when providerSettings.instanceSizeName is either M2 or M5.`,
				},
				resource.Attribute{
					Name:        "provider_instance_size_name",
					Description: `(Required) Atlas provides different instance sizes, each with a default storage capacity and RAM size. The instance size you select is used for all the data-bearing servers in your cluster. See [Create a Cluster](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/) ` + "`" + `providerSettings.instanceSizeName` + "`" + ` for valid values and default resources.`,
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
					Name:        "backup_enabled",
					Description: `(Optional) Set to true to enable Atlas continuous backups for the cluster. Set to false to disable continuous backups for the cluster. Atlas deletes any stored snapshots. See the continuous backup Snapshot Schedule for more information. You cannot enable continuous backups if you have an existing cluster in the project with Cloud Provider Snapshots enabled. The default value is false.`,
				},
				resource.Attribute{
					Name:        "bi_connector",
					Description: `(Optional) Specifies BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Optional) Specifies the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment. ->`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `(Optional) The size in gigabytes of the server’s root volume. You can add capacity by increasing this number, up to a maximum possible value of 4096 (i.e., 4 TB). This value must be a positive integer. The minimum disk size for dedicated clusters is 10GB for AWS and GCP, and 32GB for Azure. If you specify diskSizeGB with a lower disk size, Atlas defaults to the minimum disk size value.`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `(Optional) Set the Encryption at Rest parameter.`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `(Optional) Version of the cluster to deploy. Atlas supports the following MongoDB versions for M10+ clusters: ` + "`" + `3.4` + "`" + `, ` + "`" + `3.6` + "`" + ` or ` + "`" + `4.0` + "`" + `. You must set this value to ` + "`" + `4.0` + "`" + ` if ` + "`" + `provider_instance_size_name` + "`" + ` is either M2 or M5.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `(Optional) Selects whether the cluster is a replica set or a sharded cluster. If you use the replicationSpecs parameter, you must set num_shards.`,
				},
				resource.Attribute{
					Name:        "provider_backup_enabled",
					Description: `(Optional) Flag indicating if the cluster uses Cloud Provider Snapshots for backups. If true, the cluster uses Cloud Provider Snapshots for backups. If providerBackupEnabled and backupEnabled are false, the cluster does not use Atlas backups. You cannot enable cloud provider snapshots if you have an existing cluster in the project with Continuous Backups enabled.`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `(Optional) Cloud service provider on which the server for a multi-tenant cluster is provisioned. This setting is only valid when providerSetting.providerName is TENANT and providerSetting.instanceSizeName is M2 or M5. The possible values are: - AWS - Amazon AWS - GCP - Google Cloud Platform - AZURE - Microsoft Azure`,
				},
				resource.Attribute{
					Name:        "provider_disk_iops",
					Description: `(Optional) The maximum input/output operations per second (IOPS) the system can perform. The possible values depend on the selected providerSettings.instanceSizeName and diskSizeGB.`,
				},
				resource.Attribute{
					Name:        "provider_disk_type_name",
					Description: `(Optional) Azure disk type of the server’s root volume. If omitted, Atlas uses the default disk type for the selected providerSettings.instanceSizeName.`,
				},
				resource.Attribute{
					Name:        "provider_encrypt_ebs_volume",
					Description: `(Optional) If enabled, the Amazon EBS encryption feature encrypts the server’s root volume for both data at rest within the volume and for data moving between the volume and the instance.`,
				},
				resource.Attribute{
					Name:        "provider_region_name",
					Description: `(Optional) Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. Do not specify this field when creating a multi-region cluster using the replicationSpec document or a Global Cluster with the replicationSpecs array.`,
				},
				resource.Attribute{
					Name:        "provider_volume_type",
					Description: `(Optional) The type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `(Optional) Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `(Optional) Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details. ### BI Connector Specifies BI Connector for Atlas configuration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Specifies whether or not BI Connector for Atlas is enabled on the cluster. - Set to ` + "`" + `true` + "`" + ` to enable BI Connector for Atlas. - Set to ` + "`" + `false` + "`" + ` to disable BI Connector for Atlas.`,
				},
				resource.Attribute{
					Name:        "read_preference",
					Description: `(Optional) Specifies the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences). - Set to "primary" to have BI Connector for Atlas read from the primary. - Set to "secondary" to have BI Connector for Atlas read from a secondary member. Default if there are no analytics nodes in the cluster. - Set to "analytics" to have BI Connector for Atlas read from an analytics node. Default if the cluster contains analytics nodes. ### Replication Spec Configuration for cluster regions.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `(Required) Number of shards to deploy in the specified zone.`,
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
					Description: `(Optional) Name for the zone in a Global Cluster. ### Region Config Physical location of the region.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Optional) Name for the region specified.`,
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
					Description: `(Optional) The number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. If you do not specify this option, no analytics nodes are deployed to the region. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster. To review the connection string format, see the connection string format documentation. To add MongoDB users to a Atlas project, see Configure MongoDB Users. Atlas only displays this field after the cluster is operational, not while it builds the cluster.`,
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
					Description: `Current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING ## Import Clusters can be imported using project ID and cluster name, in the format ` + "`" + `PROJECTID-CLUSTERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cluster.my_cluster 1112222b3bf99403840e8934-Cluster0 ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
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
					Description: `connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster. To review the connection string format, see the connection string format documentation. To add MongoDB users to a Atlas project, see Configure MongoDB Users. Atlas only displays this field after the cluster is operational, not while it builds the cluster.`,
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
					Description: `Current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING ## Import Clusters can be imported using project ID and cluster name, in the format ` + "`" + `PROJECTID-CLUSTERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cluster.my_cluster 1112222b3bf99403840e8934-Cluster0 ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
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
					Name:        "database_name",
					Description: `(Required) The user’s authentication database. A user must provide both a username and authentication database to log into MongoDB. In Atlas deployments of MongoDB, the authentication database is always the admin database.`,
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
					Description: `(Optional) User's initial password. This is required to create the user but may be removed after. Password may show up in logs, and it will be stored in the state file as plain-text. Password can be changed in the web interface to increase security. ### Roles Block mapping a user's role to a database / collection. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. ->`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the role to grant. See [Create a Database User](https://docs.atlas.mongodb.com/reference/api/database-users-create-a-user/) ` + "`" + `roles.roleName` + "`" + ` for valid values and restrictions.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) Database on which the user has the specified role. A role on the ` + "`" + `admin` + "`" + ` database can include privileges that apply to the other databases.`,
				},
				resource.Attribute{
					Name:        "collection_name",
					Description: `(Optional) Collection for which the role applies. You can specify a collection for the ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + ` roles. If you do not specify a collection for ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + `, the role applies to all collections in the database (excluding some collections in the ` + "`" + `system` + "`" + `. database). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The database user's name. ## Import Database users can be imported using project ID and username, in the format ` + "`" + `PROJECTID-USERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_database_user.my_user 1112222b3bf99403840e8934-my_user ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The database user's name. ## Import Database users can be imported using project ID and username, in the format ` + "`" + `PROJECTID-USERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_database_user.my_user 1112222b3bf99403840e8934-my_user ` + "`" + `` + "`" + `` + "`" + ` ~>`,
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
					Description: `The AWS region in which the AWS customer master key exists: CA_CENTRAL_1, US_EAST_1, US_EAST_2, US_WEST_1, US_WEST_2, SA_EAST_1 ### azure_key_vault`,
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
					Description: `The Key Version Resource ID from your GCP account. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/encryption-at-rest/)`,
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
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `(Required) CIDR block that Atlas uses for your clusters. Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3).`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Optional | AWS provider only) AWS region.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional | AZURE provider only) Azure region where the container resides. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `The Network Peering Container ID.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `AWS region.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Azure region where the container resides.`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `Unique identifer of the Azure subscription in which the VNet resides.`,
				},
				resource.Attribute{
					Name:        "provisioned",
					Description: `Indicates whether the project has Network Peering connections deployed in the container.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `Unique identifier of the GCP project in which the Network Peering connection resides.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the Network Peering connection in the Atlas project.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Unique identifier of the project’s VPC.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `The name of the Azure VNet. This value is null until you provision an Azure VNet in the container. ## Import Clusters can be imported using project ID and network peering container id, in the format ` + "`" + `PROJECTID-CONTAINER-ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_network_container.my_container 1112222b3bf99403840e8934-5cbf563d87d9d67253be590a ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "container_id",
					Description: `The Network Peering Container ID.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `AWS region.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Azure region where the container resides.`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `Unique identifer of the Azure subscription in which the VNet resides.`,
				},
				resource.Attribute{
					Name:        "provisioned",
					Description: `Indicates whether the project has Network Peering connections deployed in the container.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `Unique identifier of the GCP project in which the Network Peering connection resides.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the Network Peering connection in the Atlas project.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Unique identifier of the project’s VPC.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `The name of the Azure VNet. This value is null until you provision an Azure VNet in the container. ## Import Clusters can be imported using project ID and network peering container id, in the format ` + "`" + `PROJECTID-CONTAINER-ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_network_container.my_container 1112222b3bf99403840e8934-5cbf563d87d9d67253be590a ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/)`,
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
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `(Required) Unique identifier of the Atlas VPC container for the region. You can create an Atlas VPC container using the Create Container endpoint. You cannot create more than one container per region. To retrieve a list of container IDs, use the Get list of VPC containers endpoint.`,
				},
				resource.Attribute{
					Name:        "accepter_region_name",
					Description: `(Optional |`,
				},
				resource.Attribute{
					Name:        "aws_account_id",
					Description: `(Optional |`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Optional) Cloud provider for this VPC peering connection. If omitted, Atlas sets this parameter to AWS. (Possible Values ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + `, ` + "`" + `GCP` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "route_table_cidr_block",
					Description: `(Optional |`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional |`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `(Optional |`,
				},
				resource.Attribute{
					Name:        "azure_directory_id",
					Description: `(Optional |`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `(Optional |`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Optional |`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `(Optional |`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `(Optinal |`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `(Optional |`,
				},
				resource.Attribute{
					Name:        "peer_id",
					Description: `The Network Peering Container ID.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "connection_id",
					Description: `Unique identifier for the peering connection.`,
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
					Name:        "provider_name",
					Description: `Cloud provider for this VPC peering connection. If omitted, Atlas sets this parameter to AWS. (Possible Values ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + `, ` + "`" + `GCP` + "`" + `).`,
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
					Description: `The VPC peering connection status value can be one of the following: ` + "`" + `INITIATING` + "`" + `, ` + "`" + `PENDING_ACCEPTANCE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `FINALIZING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `TERMINATING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `Unique identifier for an Azure AD directory.`,
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
					Name:        "resource_group_name",
					Description: `Name of your Azure resource group.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `Name of your Azure VNet.`,
				},
				resource.Attribute{
					Name:        "error_state",
					Description: `Description of the Atlas error when ` + "`" + `status` + "`" + ` is ` + "`" + `Failed` + "`" + `, Otherwise, Atlas returns ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Atlas network peering connection: ` + "`" + `ADDING_PEER` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `WAITING_FOR_USER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `GCP project ID of the owner of the network peer.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the network peer to which Atlas connects.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `When ` + "`" + `"status" : "FAILED"` + "`" + `, Atlas provides a description of the error. ## Import Clusters can be imported using project ID and network peering peering id, in the format ` + "`" + `PROJECTID-PEER-ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_network_peering.my_peering 1112222b3bf99403840e8934-5cbf563d87d9d67253be590a ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Network Peering Connection](https://docs.atlas.mongodb.com/reference/api/vpc-create-peering-connection/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "peer_id",
					Description: `The Network Peering Container ID.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "connection_id",
					Description: `Unique identifier for the peering connection.`,
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
					Name:        "provider_name",
					Description: `Cloud provider for this VPC peering connection. If omitted, Atlas sets this parameter to AWS. (Possible Values ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + `, ` + "`" + `GCP` + "`" + `).`,
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
					Description: `The VPC peering connection status value can be one of the following: ` + "`" + `INITIATING` + "`" + `, ` + "`" + `PENDING_ACCEPTANCE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `FINALIZING` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `TERMINATING` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `Unique identifier for an Azure AD directory.`,
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
					Name:        "resource_group_name",
					Description: `Name of your Azure resource group.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `Name of your Azure VNet.`,
				},
				resource.Attribute{
					Name:        "error_state",
					Description: `Description of the Atlas error when ` + "`" + `status` + "`" + ` is ` + "`" + `Failed` + "`" + `, Otherwise, Atlas returns ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Atlas network peering connection: ` + "`" + `ADDING_PEER` + "`" + `, ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `DELETING` + "`" + `, ` + "`" + `WAITING_FOR_USER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `GCP project ID of the owner of the network peer.`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `Name of the network peer to which Atlas connects.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `When ` + "`" + `"status" : "FAILED"` + "`" + `, Atlas provides a description of the error. ## Import Clusters can be imported using project ID and network peering peering id, in the format ` + "`" + `PROJECTID-PEER-ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_network_peering.my_peering 1112222b3bf99403840e8934-5cbf563d87d9d67253be590a ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Network Peering Connection](https://docs.atlas.mongodb.com/reference/api/vpc-create-peering-connection/)`,
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
					Description: `(Required) The name of the project you want to create.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) The ID of the organization you want to create the project within. ~>`,
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
					Description: `The number of Atlas clusters deployed in the project.. ## Import Project must be imported using project ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_project.my_project 5d09d6a59ccf6445652a444a ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/projects/)`,
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
					Description: `The number of Atlas clusters deployed in the project.. ## Import Project must be imported using project ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_project.my_project 5d09d6a59ccf6445652a444a ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/projects/)`,
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
					Name:        "cidr_block",
					Description: `(Optional) The whitelist entry in Classless Inter-Domain Routing (CIDR) notation. Mutually exclusive with ` + "`" + `ip_address` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The whitelisted IP address. Mutually exclusive with ` + "`" + `cidr_block` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment to add to the whitelist entry. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages and can be used to import. ## Import IP Whitelist entries can be imported using the ` + "`" + `project_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_project_ip_whitelist.test 5d0f1f74cf09a29120e123cd ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/whitelist/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages and can be used to import. ## Import IP Whitelist entries can be imported using the ` + "`" + `project_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_project_ip_whitelist.test 5d0f1f74cf09a29120e123cd ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/whitelist/)`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"mongodbatlas_cloud_provider_snapshot":             0,
		"mongodbatlas_cloud_provider_snapshot_restore_job": 1,
		"mongodbatlas_cluster":                             2,
		"mongodbatlas_database_user":                       3,
		"mongodbatlas_encryption_at_rest":                  4,
		"mongodbatlas_network_container":                   5,
		"mongodbatlas_network_peering":                     6,
		"mongodbatlas_project":                             7,
		"mongodbatlas_project_ip_whitelist":                8,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
