package mongodbatlas

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Provides an Cloud Provider Snapshot Datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Required) The unique identifier of the snapshot you want to retrieve.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Atlas cluster that contains the snapshot you want to retrieve.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the snapshot.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas took the snapshot.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas will delete the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `UDescription of the snapshot. Only present for on-demand snapshots.`,
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
					Description: `Current status of the snapshot. One of the following values: queued, inProgress, completed, failed.`,
				},
				resource.Attribute{
					Name:        "storage_size_bytes",
					Description: `Specifies the size of the snapshot in bytes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-provider-snapshot-get-one/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the snapshot.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas took the snapshot.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas will delete the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `UDescription of the snapshot. Only present for on-demand snapshots.`,
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
					Description: `Current status of the snapshot. One of the following values: queued, inProgress, completed, failed.`,
				},
				resource.Attribute{
					Name:        "storage_size_bytes",
					Description: `Specifies the size of the snapshot in bytes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-provider-snapshot-get-one/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot_restore_job",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Provider Snapshot Restore Job Datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Atlas cluster for which you want to retrieve the restore job.`,
				},
				resource.Attribute{
					Name:        "job_id",
					Description: `(Required) The unique identifier of the restore job to retrieve. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `The unique identifier of the restore job.`,
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
					Description: `Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-provider-snapshot-restore-jobs-get-one/)`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The unique identifier of the restore job.`,
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
					Description: `Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-provider-snapshot-restore-jobs-get-one/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot_restore_jobs",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Provider Snapshot Restore Jobs Datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Atlas cluster for which you want to retrieve restore jobs. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `Includes cloudProviderSnapshotRestoreJob object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### CloudProviderSnapshotRestoreJob`,
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
					Description: `The unique identifier of the restore job.`,
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
					Description: `Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-provider-snapshot-restore-jobs-get-all/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "results",
					Description: `Includes cloudProviderSnapshotRestoreJob object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### CloudProviderSnapshotRestoreJob`,
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
					Description: `The unique identifier of the restore job.`,
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
					Description: `Timestamp in ISO 8601 date and time format in UTC when the snapshot associated to snapshotId was taken. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-provider-snapshot-restore-jobs-get-all/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshots",
			Category:         "Data Sources",
			ShortDescription: `Provides an Cloud Provider Snapshot Datasource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Atlas cluster that contains the snapshot you want to retrieve.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `Includes cloudProviderSnapshot object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### CloudProviderSnapshot`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the snapshot.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas took the snapshot.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas will delete the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `UDescription of the snapshot. Only present for on-demand snapshots.`,
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
					Description: `Current status of the snapshot. One of the following values: queued, inProgress, completed, failed.`,
				},
				resource.Attribute{
					Name:        "storage_size_bytes",
					Description: `Specifies the size of the snapshot in bytes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-provider-snapshot-get-all/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "results",
					Description: `Includes cloudProviderSnapshot object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### CloudProviderSnapshot`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the snapshot.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas took the snapshot.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `UTC ISO 8601 formatted point in time when Atlas will delete the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `UDescription of the snapshot. Only present for on-demand snapshots.`,
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
					Description: `Current status of the snapshot. One of the following values: queued, inProgress, completed, failed.`,
				},
				resource.Attribute{
					Name:        "storage_size_bytes",
					Description: `Specifies the size of the snapshot in bytes.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-provider-snapshot-get-all/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cluster",
			Category:         "Data Sources",
			ShortDescription: `Describe a Cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the cluster as it appears in Atlas. Once the cluster is created, its name cannot be changed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
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
					Description: `Describes connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster. To review the connection string format, see the connection string format documentation. To add MongoDB users to a Atlas project, see Configure MongoDB Users. Atlas only displays this field after the cluster is operational, not while it builds the cluster.`,
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
					Description: `Indicates the current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING`,
				},
				resource.Attribute{
					Name:        "auto_scaling_disk_gb_enabled",
					Description: `Indicates whether disk auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "backup_enabled",
					Description: `Indicates whether Atlas continuous backups are enabled for the cluster.`,
				},
				resource.Attribute{
					Name:        "bi_connector",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Indicates the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Indicates the size in gigabytes of the server’s root volume.`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `Indicates whether Encryption at Rest is enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster as it appears in Atlas.`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `Indicates the version of the cluster to deploy.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Indicates whether the cluster is a replica set or a sharded cluster.`,
				},
				resource.Attribute{
					Name:        "provider_backup_enabled",
					Description: `Flag indicating if the cluster uses Cloud Provider Snapshots for backups.`,
				},
				resource.Attribute{
					Name:        "provider_instance_size_name",
					Description: `Atlas provides different instance sizes, each with a default storage capacity and RAM size.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Indicates the cloud service provider on which the servers are provisioned.`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `Indicates Cloud service provider on which the server for a multi-tenant cluster is provisioned.`,
				},
				resource.Attribute{
					Name:        "provider_disk_iops",
					Description: `Indicates the maximum input/output operations per second (IOPS) the system can perform. The possible values depend on the selected providerSettings.instanceSizeName and diskSizeGB.`,
				},
				resource.Attribute{
					Name:        "provider_disk_type_name",
					Description: `Describes Azure disk type of the server’s root volume.`,
				},
				resource.Attribute{
					Name:        "provider_encrypt_ebs_volume",
					Description: `Indicates whether the Amazon EBS encryption is enabled. This feature encrypts the server’s root volume for both data at rest within the volume and data moving between the volume and the instance.`,
				},
				resource.Attribute{
					Name:        "provider_region_name",
					Description: `Indicates Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases.`,
				},
				resource.Attribute{
					Name:        "provider_volume_type",
					Description: `Indicates the type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details. ### BI Connector Indicates BI Connector for Atlas configuration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates whether or not BI Connector for Atlas is enabled on the cluster.`,
				},
				resource.Attribute{
					Name:        "read_preference",
					Description: `Indicates the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences). ### Replication Spec Configuration for cluster regions.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifer of the replication document for a zone in a Global Cluster.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Number of shards to deploy in the specified zone.`,
				},
				resource.Attribute{
					Name:        "regions_config",
					Description: `Describes the physical location of the region. Each regionsConfig document describes the region’s priority in elections and the number and type of MongoDB nodes Atlas deploys to the region. You must order each regionsConfigs document by regionsConfig.priority, descending. See [Region Config](#region-config) below for more details.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Indicates the n ame for the zone in a Global Cluster. ### Region Config Physical location of the region.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Name for the region specified.`,
				},
				resource.Attribute{
					Name:        "electable_nodes",
					Description: `Number of electable nodes for Atlas to deploy to the region.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Election priority of the region. For regions with only read-only nodes, set this value to 0.`,
				},
				resource.Attribute{
					Name:        "read_only_nodes",
					Description: `Number of read-only nodes for Atlas to deploy to the region. Read-only nodes can never become the primary, but can facilitate local-reads. Specify 0 if you do not want any read-only nodes in the region.`,
				},
				resource.Attribute{
					Name:        "analytics_nodes",
					Description: `Indicates the number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
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
					Description: `Describes connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster. To review the connection string format, see the connection string format documentation. To add MongoDB users to a Atlas project, see Configure MongoDB Users. Atlas only displays this field after the cluster is operational, not while it builds the cluster.`,
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
					Description: `Indicates the current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING`,
				},
				resource.Attribute{
					Name:        "auto_scaling_disk_gb_enabled",
					Description: `Indicates whether disk auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "backup_enabled",
					Description: `Indicates whether Atlas continuous backups are enabled for the cluster.`,
				},
				resource.Attribute{
					Name:        "bi_connector",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Indicates the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Indicates the size in gigabytes of the server’s root volume.`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `Indicates whether Encryption at Rest is enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster as it appears in Atlas.`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `Indicates the version of the cluster to deploy.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Indicates whether the cluster is a replica set or a sharded cluster.`,
				},
				resource.Attribute{
					Name:        "provider_backup_enabled",
					Description: `Flag indicating if the cluster uses Cloud Provider Snapshots for backups.`,
				},
				resource.Attribute{
					Name:        "provider_instance_size_name",
					Description: `Atlas provides different instance sizes, each with a default storage capacity and RAM size.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Indicates the cloud service provider on which the servers are provisioned.`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `Indicates Cloud service provider on which the server for a multi-tenant cluster is provisioned.`,
				},
				resource.Attribute{
					Name:        "provider_disk_iops",
					Description: `Indicates the maximum input/output operations per second (IOPS) the system can perform. The possible values depend on the selected providerSettings.instanceSizeName and diskSizeGB.`,
				},
				resource.Attribute{
					Name:        "provider_disk_type_name",
					Description: `Describes Azure disk type of the server’s root volume.`,
				},
				resource.Attribute{
					Name:        "provider_encrypt_ebs_volume",
					Description: `Indicates whether the Amazon EBS encryption is enabled. This feature encrypts the server’s root volume for both data at rest within the volume and data moving between the volume and the instance.`,
				},
				resource.Attribute{
					Name:        "provider_region_name",
					Description: `Indicates Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases.`,
				},
				resource.Attribute{
					Name:        "provider_volume_type",
					Description: `Indicates the type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details. ### BI Connector Indicates BI Connector for Atlas configuration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates whether or not BI Connector for Atlas is enabled on the cluster.`,
				},
				resource.Attribute{
					Name:        "read_preference",
					Description: `Indicates the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences). ### Replication Spec Configuration for cluster regions.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifer of the replication document for a zone in a Global Cluster.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Number of shards to deploy in the specified zone.`,
				},
				resource.Attribute{
					Name:        "regions_config",
					Description: `Describes the physical location of the region. Each regionsConfig document describes the region’s priority in elections and the number and type of MongoDB nodes Atlas deploys to the region. You must order each regionsConfigs document by regionsConfig.priority, descending. See [Region Config](#region-config) below for more details.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Indicates the n ame for the zone in a Global Cluster. ### Region Config Physical location of the region.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Name for the region specified.`,
				},
				resource.Attribute{
					Name:        "electable_nodes",
					Description: `Number of electable nodes for Atlas to deploy to the region.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Election priority of the region. For regions with only read-only nodes, set this value to 0.`,
				},
				resource.Attribute{
					Name:        "read_only_nodes",
					Description: `Number of read-only nodes for Atlas to deploy to the region. Read-only nodes can never become the primary, but can facilitate local-reads. Specify 0 if you do not want any read-only nodes in the region.`,
				},
				resource.Attribute{
					Name:        "analytics_nodes",
					Description: `Indicates the number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_clusters",
			Category:         "Data Sources",
			ShortDescription: `Describe all Clusters in Project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to get the clusters. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Cluster. See [Cluster](#cluster) below for more details. ### Cluster`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster as it appears in Atlas.`,
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
					Description: `Describes connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster. To review the connection string format, see the connection string format documentation. To add MongoDB users to a Atlas project, see Configure MongoDB Users. Atlas only displays this field after the cluster is operational, not while it builds the cluster.`,
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
					Description: `Indicates the current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING`,
				},
				resource.Attribute{
					Name:        "auto_scaling_disk_gb_enabled",
					Description: `Indicates whether disk auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "backup_enabled",
					Description: `Indicates whether Atlas continuous backups are enabled for the cluster.`,
				},
				resource.Attribute{
					Name:        "bi_connector",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Indicates the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Indicates the size in gigabytes of the server’s root volume.`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `Indicates whether Encryption at Rest is enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `Indicates the version of the cluster to deploy.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Indicates whether the cluster is a replica set or a sharded cluster.`,
				},
				resource.Attribute{
					Name:        "provider_backup_enabled",
					Description: `Flag indicating if the cluster uses Cloud Provider Snapshots for backups.`,
				},
				resource.Attribute{
					Name:        "provider_instance_size_name",
					Description: `Atlas provides different instance sizes, each with a default storage capacity and RAM size.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Indicates the cloud service provider on which the servers are provisioned.`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `Indicates Cloud service provider on which the server for a multi-tenant cluster is provisioned.`,
				},
				resource.Attribute{
					Name:        "provider_disk_iops",
					Description: `Indicates the maximum input/output operations per second (IOPS) the system can perform. The possible values depend on the selected providerSettings.instanceSizeName and diskSizeGB.`,
				},
				resource.Attribute{
					Name:        "provider_disk_type_name",
					Description: `Describes Azure disk type of the server’s root volume.`,
				},
				resource.Attribute{
					Name:        "provider_encrypt_ebs_volume",
					Description: `Indicates whether the Amazon EBS encryption is enabled. This feature encrypts the server’s root volume for both data at rest within the volume and data moving between the volume and the instance.`,
				},
				resource.Attribute{
					Name:        "provider_region_name",
					Description: `Indicates Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases.`,
				},
				resource.Attribute{
					Name:        "provider_volume_type",
					Description: `Indicates the type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details. ### BI Connector Indicates BI Connector for Atlas configuration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates whether or not BI Connector for Atlas is enabled on the cluster.`,
				},
				resource.Attribute{
					Name:        "read_preference",
					Description: `Indicates the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences). ### Replication Spec Configuration for cluster regions.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifer of the replication document for a zone in a Global Cluster.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Number of shards to deploy in the specified zone.`,
				},
				resource.Attribute{
					Name:        "regions_config",
					Description: `Describes the physical location of the region. Each regionsConfig document describes the region’s priority in elections and the number and type of MongoDB nodes Atlas deploys to the region. You must order each regionsConfigs document by regionsConfig.priority, descending. See [Region Config](#region-config) below for more details.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Indicates the n ame for the zone in a Global Cluster. ### Region Config Physical location of the region.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Name for the region specified.`,
				},
				resource.Attribute{
					Name:        "electable_nodes",
					Description: `Number of electable nodes for Atlas to deploy to the region.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Election priority of the region. For regions with only read-only nodes, set this value to 0.`,
				},
				resource.Attribute{
					Name:        "read_only_nodes",
					Description: `Number of read-only nodes for Atlas to deploy to the region. Read-only nodes can never become the primary, but can facilitate local-reads. Specify 0 if you do not want any read-only nodes in the region.`,
				},
				resource.Attribute{
					Name:        "analytics_nodes",
					Description: `Indicates the number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Cluster. See [Cluster](#cluster) below for more details. ### Cluster`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster as it appears in Atlas.`,
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
					Description: `Describes connection string for connecting to the Atlas cluster. Includes the replicaSet, ssl, and authSource query parameters in the connection string with values appropriate for the cluster. To review the connection string format, see the connection string format documentation. To add MongoDB users to a Atlas project, see Configure MongoDB Users. Atlas only displays this field after the cluster is operational, not while it builds the cluster.`,
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
					Description: `Indicates the current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING`,
				},
				resource.Attribute{
					Name:        "auto_scaling_disk_gb_enabled",
					Description: `Indicates whether disk auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "backup_enabled",
					Description: `Indicates whether Atlas continuous backups are enabled for the cluster.`,
				},
				resource.Attribute{
					Name:        "bi_connector",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Indicates the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Indicates the size in gigabytes of the server’s root volume.`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `Indicates whether Encryption at Rest is enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `Indicates the version of the cluster to deploy.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Indicates whether the cluster is a replica set or a sharded cluster.`,
				},
				resource.Attribute{
					Name:        "provider_backup_enabled",
					Description: `Flag indicating if the cluster uses Cloud Provider Snapshots for backups.`,
				},
				resource.Attribute{
					Name:        "provider_instance_size_name",
					Description: `Atlas provides different instance sizes, each with a default storage capacity and RAM size.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Indicates the cloud service provider on which the servers are provisioned.`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `Indicates Cloud service provider on which the server for a multi-tenant cluster is provisioned.`,
				},
				resource.Attribute{
					Name:        "provider_disk_iops",
					Description: `Indicates the maximum input/output operations per second (IOPS) the system can perform. The possible values depend on the selected providerSettings.instanceSizeName and diskSizeGB.`,
				},
				resource.Attribute{
					Name:        "provider_disk_type_name",
					Description: `Describes Azure disk type of the server’s root volume.`,
				},
				resource.Attribute{
					Name:        "provider_encrypt_ebs_volume",
					Description: `Indicates whether the Amazon EBS encryption is enabled. This feature encrypts the server’s root volume for both data at rest within the volume and data moving between the volume and the instance.`,
				},
				resource.Attribute{
					Name:        "provider_region_name",
					Description: `Indicates Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases.`,
				},
				resource.Attribute{
					Name:        "provider_volume_type",
					Description: `Indicates the type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details. ### BI Connector Indicates BI Connector for Atlas configuration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates whether or not BI Connector for Atlas is enabled on the cluster.`,
				},
				resource.Attribute{
					Name:        "read_preference",
					Description: `Indicates the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences). ### Replication Spec Configuration for cluster regions.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifer of the replication document for a zone in a Global Cluster.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Number of shards to deploy in the specified zone.`,
				},
				resource.Attribute{
					Name:        "regions_config",
					Description: `Describes the physical location of the region. Each regionsConfig document describes the region’s priority in elections and the number and type of MongoDB nodes Atlas deploys to the region. You must order each regionsConfigs document by regionsConfig.priority, descending. See [Region Config](#region-config) below for more details.`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Indicates the n ame for the zone in a Global Cluster. ### Region Config Physical location of the region.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Name for the region specified.`,
				},
				resource.Attribute{
					Name:        "electable_nodes",
					Description: `Number of electable nodes for Atlas to deploy to the region.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Election priority of the region. For regions with only read-only nodes, set this value to 0.`,
				},
				resource.Attribute{
					Name:        "read_only_nodes",
					Description: `Number of read-only nodes for Atlas to deploy to the region. Read-only nodes can never become the primary, but can facilitate local-reads. Specify 0 if you do not want any read-only nodes in the region.`,
				},
				resource.Attribute{
					Name:        "analytics_nodes",
					Description: `Indicates the number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_database_user",
			Category:         "Data Sources",
			ShortDescription: `Describes a Database User.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username for authenticating to MongoDB.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The database user's name.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See [Roles](#roles) below for more details.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The user’s authentication database. A user must provide both a username and authentication database to log into MongoDB. In Atlas deployments of MongoDB, the authentication database is always the admin database. ### Roles Block mapping a user's role to a database / collection. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. ->`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the role to grant.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `Database on which the user has the specified role. A role on the ` + "`" + `admin` + "`" + ` database can include privileges that apply to the other databases.`,
				},
				resource.Attribute{
					Name:        "collection_name",
					Description: `Collection for which the role applies. You can specify a collection for the ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + ` roles. If you do not specify a collection for ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + `, the role applies to all collections in the database (excluding some collections in the ` + "`" + `system` + "`" + `. database). See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/database-users-get-single-user/) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The database user's name.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See [Roles](#roles) below for more details.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The user’s authentication database. A user must provide both a username and authentication database to log into MongoDB. In Atlas deployments of MongoDB, the authentication database is always the admin database. ### Roles Block mapping a user's role to a database / collection. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. ->`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the role to grant.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `Database on which the user has the specified role. A role on the ` + "`" + `admin` + "`" + ` database can include privileges that apply to the other databases.`,
				},
				resource.Attribute{
					Name:        "collection_name",
					Description: `Collection for which the role applies. You can specify a collection for the ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + ` roles. If you do not specify a collection for ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + `, the role applies to all collections in the database (excluding some collections in the ` + "`" + `system` + "`" + `. database). See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/database-users-get-single-user/) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_database_users",
			Category:         "Data Sources",
			ShortDescription: `Describes a Database Users.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to get all database users. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Database user. ### Database User`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the Atlas project the user belongs to.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for authenticating to MongoDB.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See [Roles](#roles) below for more details.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The user’s authentication database. A user must provide both a username and authentication database to log into MongoDB. In Atlas deployments of MongoDB, the authentication database is always the admin database. ### Roles Block mapping a user's role to a database / collection. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. ->`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the role to grant.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `Database on which the user has the specified role. A role on the ` + "`" + `admin` + "`" + ` database can include privileges that apply to the other databases.`,
				},
				resource.Attribute{
					Name:        "collection_name",
					Description: `Collection for which the role applies. You can specify a collection for the ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + ` roles. If you do not specify a collection for ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + `, the role applies to all collections in the database (excluding some collections in the ` + "`" + `system` + "`" + `. database). See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/database-users-get-single-user/) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Database user. ### Database User`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of the Atlas project the user belongs to.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username for authenticating to MongoDB.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See [Roles](#roles) below for more details.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The user’s authentication database. A user must provide both a username and authentication database to log into MongoDB. In Atlas deployments of MongoDB, the authentication database is always the admin database. ### Roles Block mapping a user's role to a database / collection. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. ->`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the role to grant.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `Database on which the user has the specified role. A role on the ` + "`" + `admin` + "`" + ` database can include privileges that apply to the other databases.`,
				},
				resource.Attribute{
					Name:        "collection_name",
					Description: `Collection for which the role applies. You can specify a collection for the ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + ` roles. If you do not specify a collection for ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + `, the role applies to all collections in the database (excluding some collections in the ` + "`" + `system` + "`" + `. database). See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/database-users-get-single-user/) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_container",
			Category:         "Data Sources",
			ShortDescription: `Describes a Cluster resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `(Required) The Network Peering Container ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Container ID.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `CIDR block that Atlas uses for your clusters. Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3).`,
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
					Description: `The name of the Azure VNet. This value is null until you provision an Azure VNet in the container. See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Container ID.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `CIDR block that Atlas uses for your clusters. Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3).`,
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
					Description: `The name of the Azure VNet. This value is null until you provision an Azure VNet in the container. See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_containers",
			Category:         "Data Sources",
			ShortDescription: `Describes all Network Peering Containers in the project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Network Peering Container. ### Network Peering Container`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Container ID.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `CIDR block that Atlas uses for your clusters. Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3).`,
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
					Description: `The name of the Azure VNet. This value is null until you provision an Azure VNet in the container. See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-get-containers-list/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Network Peering Container. ### Network Peering Container`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Container ID.`,
				},
				resource.Attribute{
					Name:        "atlas_cidr_block",
					Description: `CIDR block that Atlas uses for your clusters. Atlas uses the specified CIDR block for all other Network Peering connections created in the project. The Atlas CIDR block must be at least a /24 and at most a /21 in one of the following [private networks](https://tools.ietf.org/html/rfc1918.html#section-3).`,
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
					Description: `The name of the Azure VNet. This value is null until you provision an Azure VNet in the container. See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-get-containers-list/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_peering",
			Category:         "Data Sources",
			ShortDescription: `Describes a Network Peering.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "peering_id",
					Description: `(Required) Atlas assigned unique ID for the peering connection. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Connection ID.`,
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
					Description: `When ` + "`" + `"status" : "FAILED"` + "`" + `, Atlas provides a description of the error. See detailed information for arguments and attributes: [MongoDB API Network Peering Connection](https://docs.atlas.mongodb.com/reference/api/vpc-get-connection/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Connection ID.`,
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
					Description: `When ` + "`" + `"status" : "FAILED"` + "`" + `, Atlas provides a description of the error. See detailed information for arguments and attributes: [MongoDB API Network Peering Connection](https://docs.atlas.mongodb.com/reference/api/vpc-get-connection/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_peerings",
			Category:         "Data Sources",
			ShortDescription: `Describes all Network Peering Connections.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Connection ID.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Network Peering Connection. ### Network Peering Connection`,
				},
				resource.Attribute{
					Name:        "peering_id",
					Description: `Atlas assigned unique ID for the peering connection.`,
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
					Description: `When ` + "`" + `"status" : "FAILED"` + "`" + `, Atlas provides a description of the error. See detailed information for arguments and attributes: [MongoDB API Network Peering Connection](https://docs.atlas.mongodb.com/reference/api/vpc-get-connections-list/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Network Peering Connection ID.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Network Peering Connection. ### Network Peering Connection`,
				},
				resource.Attribute{
					Name:        "peering_id",
					Description: `Atlas assigned unique ID for the peering connection.`,
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
					Description: `When ` + "`" + `"status" : "FAILED"` + "`" + `, Atlas provides a description of the error. See detailed information for arguments and attributes: [MongoDB API Network Peering Connection](https://docs.atlas.mongodb.com/reference/api/vpc-get-connections-list/)`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"mongodbatlas_cloud_provider_snapshot":              0,
		"mongodbatlas_cloud_provider_snapshot_restore_job":  1,
		"mongodbatlas_cloud_provider_snapshot_restore_jobs": 2,
		"mongodbatlas_cloud_provider_snapshots":             3,
		"mongodbatlas_cluster":                              4,
		"mongodbatlas_clusters":                             5,
		"mongodbatlas_database_user":                        6,
		"mongodbatlas_database_users":                       7,
		"mongodbatlas_network_container":                    8,
		"mongodbatlas_network_containers":                   9,
		"mongodbatlas_network_peering":                      10,
		"mongodbatlas_network_peerings":                     11,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
