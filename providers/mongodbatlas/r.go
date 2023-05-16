package mongodbatlas

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_access_list_api_key",
			Category:         "Resources",
			ShortDescription: `Provides an Access List API Key resource.`,
			Description: `

` + "`" + `mongodbatlas_access_list_api_key` + "`" + ` provides an IP Access List entry resource. The access list grants access from IPs, CIDRs or AWS Security Groups (if VPC Peering is enabled) to clusters within the Project.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

~> **IMPORTANT:**
When you remove an entry from the access list, existing connections from the removed address(es) may remain open for a variable amount of time. How much time passes before Atlas closes the connection depends on several factors, including how the connection was established, the particular behavior of the application or driver using the address, and the connection protocol (e.g., TCP or UDP). This is particularly important to consider when changing an existing IP address or CIDR block as they cannot be updated via the Provider (comments can however), hence a change will force the destruction and recreation of entries.  

~> **IMPORTANT WARNING:** Managing Atlas Programmatic API Keys (PAKs) with Terraform will expose sensitive organizational secrets in Terraform's state. We suggest following [Terraform's best practices](https://developer.hashicorp.com/terraform/language/state/sensitive-data). You may also want to consider managing your PAKs via a more secure method, such as the [HashiCorp Vault MongoDB Atlas Secrets Engine](https://developer.hashicorp.com/vault/docs/secrets/mongodbatlas).


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) Unique identifier for the organinzation to which you want to add one or more access list entries.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) Range of IP addresses in CIDR notation to be added to the access list. Your access list entry can include only one ` + "`" + `cidrBlock` + "`" + `, or one ` + "`" + `ipAddress` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Single IP address to be added to the access list.`,
				},
				resource.Attribute{
					Name:        "api_key_id",
					Description: `Unique identifier for the Organization API Key for which you want to create a new access list entry. ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_advanced_cluster",
			Category:         "Resources",
			ShortDescription: `Provides an Advanced Cluster resource.`,
			Description: `

` + "`" + `mongodbatlas_advanced_cluster` + "`" + ` provides an Advanced Cluster resource. The resource lets you create, edit and delete advanced clusters. The resource requires your Project ID.

More information on considerations for using advanced clusters please see [Considerations](https://docs.atlas.mongodb.com/reference/api/cluster-advanced/create-one-cluster-advanced/#considerations)

~> **IMPORTANT:**
<br> &#8226; The primary difference between [` + "`" + `mongodbatlas_cluster` + "`" + `](https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/resources/cluster) and [` + "`" + `mongodbatlas_advanced_cluster` + "`" + `](https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/resources/advanced_cluster) is that ` + "`" + `mongodbatlas_advanced_cluster` + "`" + ` supports multi-cloud clusters.  We recommend new users start with the ` + "`" + `mongodbatlas_advanced_cluster` + "`" + ` resource.  

-> **NOTE:** If Backup Compliance Policy is enabled for the project for which this backup schedule is defined, you cannot modify the backup schedule for an individual cluster below the minimum requirements set in the Backup Compliance Policy.  See [Backup Compliance Policy Prohibited Actions and Considerations](https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/#configure-a-backup-compliance-policy).


<br> &#8226; Upgrading the shared tier is supported. Any change from a shared tier cluster (a tenant) to a different instance size will be considered a tenant upgrade. When upgrading from the shared tier, change the ` + "`" + `provider_name` + "`" + ` from "TENANT" to your preferred provider (AWS, GCP or Azure) and remove the variable ` + "`" + `backing_provider_name` + "`" + `.  See the [Example Tenant Cluster Upgrade](#Example-Tenant-Cluster-Upgrade) below.   Note you can upgrade a shared tier cluster only to a single provider on an M10-tier cluster or greater.  
<br> &#8226; **IMPORTANT NOTE** When upgrading from the shared tier, *only* the upgrade changes will be applied. This helps avoid a corrupt state file in the event that the upgrade succeeds but subsequent updates fail within the same ` + "`" + `terraform apply` + "`" + `. To apply additional cluster changes, run a secondary ` + "`" + `terraform apply` + "`" + ` after the upgrade succeeds.
-> **NOTE:** Groups and projects are synonymous terms. You might find group_id in the official documentation.

-> **NOTE:** A network container is created for each provider/region combination on the advanced cluster. This can be referenced via a computed attribute for use with other resources. Refer to the ` + "`" + `replication_specs.#.container_id` + "`" + ` attribute in the [Attributes Reference](#attributes_reference) for more information.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the cluster as it appears in Atlas. Once the cluster is created, its name cannot be changed.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) Atlas provides different instance sizes, each with a default storage capacity and RAM size. The instance size you select is used for all the data-bearing servers in your cluster. See [Create a Cluster](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/) ` + "`" + `providerSettings.instanceSizeName` + "`" + ` for valid values and default resources.`,
				},
				resource.Attribute{
					Name:        "backup_enabled",
					Description: `(Optional) Flag that indicates whether the cluster can perform backups. If ` + "`" + `true` + "`" + `, the cluster can perform backups. You must set this value to ` + "`" + `true` + "`" + ` for NVMe clusters. Backup uses: [Cloud Backups](https://docs.atlas.mongodb.com/backup/cloud-backup/overview/#std-label-backup-cloud-provider) for dedicated clusters. [Shared Cluster Backups](https://docs.atlas.mongodb.com/backup/shared-tier/overview/#std-label-m2-m5-snapshots) for tenant clusters. If "` + "`" + `backup_enabled` + "`" + `" : ` + "`" + `false` + "`" + `, the cluster doesn't use Atlas backups. This parameter defaults to false.`,
				},
				resource.Attribute{
					Name:        "bi_connector_config",
					Description: `(Optional) Configuration settings applied to BI Connector for Atlas on this cluster. The MongoDB Connector for Business Intelligence for Atlas (BI Connector) is only available for M10 and larger clusters. The BI Connector is a powerful tool which provides users SQL-based access to their MongoDB databases. As a result, the BI Connector performs operations which may be CPU and memory intensive. Given the limited hardware resources on M10 and M20 cluster tiers, you may experience performance degradation of the cluster when enabling the BI Connector. If this occurs, upgrade to an M30 or larger cluster or disable the BI Connector. See [below](#bi_connector_config).`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required)Type of the cluster that you want to create. Accepted values include: - ` + "`" + `REPLICASET` + "`" + ` Replica set - ` + "`" + `SHARDED` + "`" + ` Sharded cluster - ` + "`" + `GEOSHARDED` + "`" + ` Global Cluster`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `(Optional) Capacity, in gigabytes, of the host's root volume. Increase this number to add capacity, up to a maximum possible value of 4096 (i.e., 4 TB). This value must be a positive number. You can't set this value with clusters with local [NVMe SSDs](https://docs.atlas.mongodb.com/cluster-tier/#std-label-nvme-storage). The minimum disk size for dedicated clusters is 10 GB for AWS and GCP. If you specify diskSizeGB with a lower disk size, Atlas defaults to the minimum disk size value. If your cluster includes Azure nodes, this value must correspond to an existing Azure disk type (8, 16, 32, 64, 128, 256, 512, 1024, 2048, or 4095)Atlas calculates storage charges differently depending on whether you choose the default value or a custom value. The maximum value for disk storage cannot exceed 50 times the maximum RAM for the selected cluster. If you require additional storage space beyond this limitation, consider [upgrading your cluster](https://docs.atlas.mongodb.com/scale-cluster/#std-label-scale-cluster-instance) to a higher tier. If your cluster spans cloud service providers, this value defaults to the minimum default of the providers involved.`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `(Optional) Possible values are AWS, GCP, AZURE or NONE. Only needed if you desire to manage the keys, see [Encryption at Rest using Customer Key Management](https://docs.atlas.mongodb.com/security-kms-encryption/) for complete documentation. You must configure encryption at rest for the Atlas project before enabling it on any cluster in the project. For Documentation, see [AWS](https://docs.atlas.mongodb.com/security-aws-kms/), [GCP](https://docs.atlas.mongodb.com/security-kms-encryption/) and [Azure](https://docs.atlas.mongodb.com/security-azure-kms/#std-label-security-azure-kms). Requirements are if ` + "`" + `replication_specs.#.region_configs.#.<type>Specs.instance_size` + "`" + ` is M10 or greater and ` + "`" + `backup_enabled` + "`" + ` is false or omitted.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Configuration for the collection of key-value pairs that tag and categorize the cluster. See [below](#labels).`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `(Optional) Version of the cluster to deploy. Atlas supports the following MongoDB versions for M10+ clusters: ` + "`" + `4.0` + "`" + `, ` + "`" + `4.2` + "`" + `, ` + "`" + `4.4` + "`" + `, or ` + "`" + `5.0` + "`" + `. If omitted, Atlas deploys a cluster that runs MongoDB 4.4. If ` + "`" + `replication_specs#.region_configs#.<type>Specs.instance_size` + "`" + `: ` + "`" + `M0` + "`" + `, ` + "`" + `M2` + "`" + ` or ` + "`" + `M5` + "`" + `, Atlas deploys MongoDB 4.4. Atlas always deploys the cluster with the latest stable release of the specified version. If you set a value to this parameter and set ` + "`" + `version_release_system` + "`" + ` ` + "`" + `CONTINUOUS` + "`" + `, the resource returns an error. Either clear this parameter or set ` + "`" + `version_release_system` + "`" + `: ` + "`" + `LTS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pit_enabled",
					Description: `(Optional) - Flag that indicates if the cluster uses Continuous Cloud Backup.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions and the hardware provisioned in them. See [below](#replication_specs)`,
				},
				resource.Attribute{
					Name:        "root_cert_type",
					Description: `(Optional) - Certificate Authority that MongoDB Atlas clusters use. You can specify ISRGROOTX1 (for ISRG Root X1).`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.`,
				},
				resource.Attribute{
					Name:        "version_release_system",
					Description: `(Optional) - Release cadence that Atlas uses for this cluster. This parameter defaults to ` + "`" + `LTS` + "`" + `. If you set this field to ` + "`" + `CONTINUOUS` + "`" + `, you must omit the ` + "`" + `mongo_db_major_version` + "`" + ` field. Atlas accepts: - ` + "`" + `CONTINUOUS` + "`" + `: Atlas creates your cluster using the most recent MongoDB release. Atlas automatically updates your cluster to the latest major and rapid MongoDB releases as they become available. - ` + "`" + `LTS` + "`" + `: Atlas creates your cluster using the latest patch release of the MongoDB version that you specify in the mongoDBMajorVersion field. Atlas automatically updates your cluster to subsequent patch releases of this MongoDB version. Atlas doesn't update your cluster to newer rapid or major MongoDB releases as they become available.`,
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
					Name:        "default_read_concern",
					Description: `(Optional) [Default level of acknowledgment requested from MongoDB for read operations](https://docs.mongodb.com/manual/reference/read-concern/) set for this cluster. MongoDB 4.4 clusters default to [available](https://docs.mongodb.com/manual/reference/read-concern-available/).`,
				},
				resource.Attribute{
					Name:        "default_write_concern",
					Description: `(Optional) [Default level of acknowledgment requested from MongoDB for write operations](https://docs.mongodb.com/manual/reference/write-concern/) set for this cluster. MongoDB 4.4 clusters default to [1](https://docs.mongodb.com/manual/reference/write-concern/).`,
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
					Name:        "oplog_min_retention_hours",
					Description: `(Optional) Minimum retention window for cluster's oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates.`,
				},
				resource.Attribute{
					Name:        "sample_size_bi_connector",
					Description: `(Optional) Number of documents per database to sample when gathering schema information. Defaults to 100. Available only for Atlas deployments in which BI Connector for Atlas is enabled.`,
				},
				resource.Attribute{
					Name:        "sample_refresh_interval_bi_connector",
					Description: `(Optional) Interval in seconds at which the mongosqld process re-samples data to create its relational schema. The default value is 300. The specified value must be a positive integer. Available only for Atlas deployments in which BI Connector for Atlas is enabled. ### labels ` + "`" + `` + "`" + `` + "`" + `terraform labels { key = "Key 1" value = "Value 1" } labels { key = "Key 2" value = "Value 2" } ` + "`" + `` + "`" + `` + "`" + ` Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. You cannot set the key ` + "`" + `Infrastructure Tool` + "`" + `, it is used for internal purposes to track aggregate usage.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. ### replication_specs ` + "`" + `` + "`" + `` + "`" + `terraform //Example Multicloud replication_specs { region_configs { electable_specs { instance_size = "M10" node_count = 3 } analytics_specs { instance_size = "M10" node_count = 1 } provider_name = "AWS" priority = 1 region_name = "US_EAST_1" } region_configs { electable_specs { instance_size = "M10" node_count = 2 } provider_name = "GCP" priority = 6 region_name = "NORTH_AMERICA_NORTHEAST_1" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `(Required) Provide this value if you set a ` + "`" + `cluster_type` + "`" + ` of SHARDED or GEOSHARDED. Omit this value if you selected a ` + "`" + `cluster_type` + "`" + ` of REPLICASET. This API resource accepts 1 through 50, inclusive. This parameter defaults to 1. If you specify a ` + "`" + `num_shards` + "`" + ` value of 1 and a ` + "`" + `cluster_type` + "`" + ` of SHARDED, Atlas deploys a single-shard [sharded cluster](https://docs.atlas.mongodb.com/reference/glossary/#std-term-sharded-cluster). Don't create a sharded cluster with a single shard for production environments. Single-shard sharded clusters don't provide the same benefits as multi-shard configurations.`,
				},
				resource.Attribute{
					Name:        "region_configs",
					Description: `(Optional) Configuration for the hardware specifications for nodes set for a given regionEach ` + "`" + `region_configs` + "`" + ` object describes the region's priority in elections and the number and type of MongoDB nodes that Atlas deploys to the region. Each ` + "`" + `region_configs` + "`" + ` object must have either an ` + "`" + `analytics_specs` + "`" + ` object, ` + "`" + `electable_specs` + "`" + ` object, or ` + "`" + `read_only_specs` + "`" + ` object. See [below](#region_configs)`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Optional) Name for the zone in a Global Cluster. ### region_configs`,
				},
				resource.Attribute{
					Name:        "analytics_specs",
					Description: `(Optional) Hardware specifications for [analytics nodes](https://docs.atlas.mongodb.com/reference/faq/deployment/#std-label-analytics-nodes-overview) needed in the region. Analytics nodes handle analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only and can never become the [primary](https://docs.atlas.mongodb.com/reference/glossary/#std-term-primary). If you don't specify this parameter, no analytics nodes deploy to this region. See [below](#specs)`,
				},
				resource.Attribute{
					Name:        "auto_scaling",
					Description: `(Optional) Configuration for the Collection of settings that configures auto-scaling information for the cluster. The values for the ` + "`" + `auto_scaling` + "`" + ` parameter must be the same for every item in the ` + "`" + `replication_specs` + "`" + ` array. See [below](#auto_scaling)`,
				},
				resource.Attribute{
					Name:        "analytics_auto_scaling",
					Description: `(Optional) Configuration for the Collection of settings that configures analytics-auto-scaling information for the cluster. The values for the ` + "`" + `analytics_auto_scaling` + "`" + ` parameter must be the same for every item in the ` + "`" + `replication_specs` + "`" + ` array. See [below](#analytics_auto_scaling)`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `(Optional) Cloud service provider on which you provision the host for a multi-tenant cluster. Use this only when a ` + "`" + `provider_name` + "`" + ` is ` + "`" + `TENANT` + "`" + ` and ` + "`" + `instance_size` + "`" + ` of a specs is ` + "`" + `M2` + "`" + ` or ` + "`" + `M5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "electable_specs",
					Description: `(Optional) Hardware specifications for electable nodes in the region. Electable nodes can become the [primary](https://docs.atlas.mongodb.com/reference/glossary/#std-term-primary) and can enable local reads. If you do not specify this option, no electable nodes are deployed to the region. See [below](#specs)`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Election priority of the region. For regions with only read-only nodes, set this value to 0.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Optional) Cloud service provider on which the servers are provisioned. The possible values are: - ` + "`" + `AWS` + "`" + ` - Amazon AWS - ` + "`" + `GCP` + "`" + ` - Google Cloud Platform - ` + "`" + `AZURE` + "`" + ` - Microsoft Azure - ` + "`" + `TENANT` + "`" + ` - M2 or M5 multi-tenant cluster. Use ` + "`" + `replication_specs.#.region_configs.#.backing_provider_name` + "`" + ` to set the cloud service provider.`,
				},
				resource.Attribute{
					Name:        "read_only_specs",
					Description: `(Optional) Hardware specifications for read-only nodes in the region. Read-only nodes can become the [primary](https://docs.atlas.mongodb.com/reference/glossary/#std-term-primary) and can enable local reads. If you don't specify this parameter, no read-only nodes are deployed to the region. See [below](#specs)`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Optional) Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. Requires the`,
				},
				resource.Attribute{
					Name:        "disk_iops",
					Description: `(Optional) Target throughput (IOPS) desired for AWS storage attached to your cluster. Set only if you selected AWS as your cloud service provider. You can't set this parameter for a multi-cloud cluster.`,
				},
				resource.Attribute{
					Name:        "ebs_volume_type",
					Description: `(Optional) Type of storage you want to attach to your AWS-provisioned cluster. Set only if you selected AWS as your cloud service provider. You can't set this parameter for a multi-cloud cluster. Valid values are:`,
				},
				resource.Attribute{
					Name:        "instance_size",
					Description: `(Optional) Hardware specification for the instance sizes in this region. Each instance size has a default storage and memory capacity. The instance size you select applies to all the data-bearing hosts in your instance size.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `(Optional) Number of read-only nodes for Atlas to deploy to the region. Read-only nodes can never become the [primary](https://docs.atlas.mongodb.com/reference/glossary/#std-term-primary), but can enable local reads. ### auto_scaling`,
				},
				resource.Attribute{
					Name:        "disk_gb_enabled",
					Description: `(Optional) Flag that indicates whether this cluster enables disk auto-scaling. This parameter defaults to true. - Set to ` + "`" + `true` + "`" + ` to enable disk auto-scaling. - Set to ` + "`" + `false` + "`" + ` to disable disk auto-scaling. ~>`,
				},
				resource.Attribute{
					Name:        "compute_enabled",
					Description: `(Optional) Flag that indicates whether instance size auto-scaling is enabled. This parameter defaults to false. ~>`,
				},
				resource.Attribute{
					Name:        "compute_scale_down_enabled",
					Description: `(Optional) Flag that indicates whether the instance size may scale down. Atlas requires this parameter if ` + "`" + `replication_specs.#.region_configs.#.auto_scaling.0.compute_enabled` + "`" + ` : true. If you enable this option, specify a value for ` + "`" + `replication_specs.#.region_configs.#.auto_scaling.0.compute_min_instance_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "compute_min_instance_size",
					Description: `(Optional) Minimum instance size to which your cluster can automatically scale (such as M10). Atlas requires this parameter if ` + "`" + `replication_specs.#.region_configs.#.auto_scaling.0.compute_scale_down_enabled` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "compute_max_instance_size",
					Description: `(Optional) Maximum instance size to which your cluster can automatically scale (such as M40). Atlas requires this parameter if ` + "`" + `replication_specs.#.region_configs.#.auto_scaling.0.compute_enabled` + "`" + ` is true. ### analytics_auto_scaling`,
				},
				resource.Attribute{
					Name:        "disk_gb_enabled",
					Description: `(Optional) Flag that indicates whether this cluster enables disk auto-scaling. This parameter defaults to true.`,
				},
				resource.Attribute{
					Name:        "compute_enabled",
					Description: `(Optional) Flag that indicates whether instance size auto-scaling is enabled. This parameter defaults to false. ~>`,
				},
				resource.Attribute{
					Name:        "compute_scale_down_enabled",
					Description: `(Optional) Flag that indicates whether the instance size may scale down. Atlas requires this parameter if ` + "`" + `replication_specs.#.region_configs.#.analytics_auto_scaling.0.compute_enabled` + "`" + ` : true. If you enable this option, specify a value for ` + "`" + `replication_specs.#.region_configs.#.analytics_auto_scaling.0.compute_min_instance_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "compute_min_instance_size",
					Description: `(Optional) Minimum instance size to which your cluster can automatically scale (such as M10). Atlas requires this parameter if ` + "`" + `replication_specs.#.region_configs.#.analytics_auto_scaling.0.compute_scale_down_enabled` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "compute_max_instance_size",
					Description: `(Optional) Maximum instance size to which your cluster can automatically scale (such as M40). Atlas requires this parameter if ` + "`" + `replication_specs.#.region_configs.#.analytics_auto_scaling.0.compute_enabled` + "`" + ` is true. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "connection_strings",
					Description: `Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Set of replication specifications for the cluster. Primary usage is covered under the [replication_specs argument reference](#replication_specs), though there are some computed attributes: - ` + "`" + `replication_specs.#.container_id` + "`" + ` - A key-value map of the Network Peering Container ID(s) for the configuration specified in ` + "`" + `region_configs` + "`" + `. The Container ID is the id of the container created when the first cluster in the region (AWS/Azure) or project (GCP) was created. The syntax is ` + "`" + `"providerName:regionName" = "containerId"` + "`" + `. Example ` + "`" + `AWS:US_EAST_1" = "61e0797dde08fb498ca11a71` + "`" + `. ## Import Clusters can be imported using project ID and cluster name, in the format ` + "`" + `PROJECTID-CLUSTERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_advanced_cluster.my_cluster 1112222b3bf99403840e8934-Cluster0 ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Advanced Clusters](https://docs.atlas.mongodb.com/reference/api/cluster-advanced/create-one-cluster-advanced/)`,
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
					Name:        "connection_strings",
					Description: `Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Set of replication specifications for the cluster. Primary usage is covered under the [replication_specs argument reference](#replication_specs), though there are some computed attributes: - ` + "`" + `replication_specs.#.container_id` + "`" + ` - A key-value map of the Network Peering Container ID(s) for the configuration specified in ` + "`" + `region_configs` + "`" + `. The Container ID is the id of the container created when the first cluster in the region (AWS/Azure) or project (GCP) was created. The syntax is ` + "`" + `"providerName:regionName" = "containerId"` + "`" + `. Example ` + "`" + `AWS:US_EAST_1" = "61e0797dde08fb498ca11a71` + "`" + `. ## Import Clusters can be imported using project ID and cluster name, in the format ` + "`" + `PROJECTID-CLUSTERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_advanced_cluster.my_cluster 1112222b3bf99403840e8934-Cluster0 ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Advanced Clusters](https://docs.atlas.mongodb.com/reference/api/cluster-advanced/create-one-cluster-advanced/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_alert_configuration",
			Category:         "Resources",
			ShortDescription: `Provides an Alert Configuration resource.`,
			Description: `

` + "`" + `mongodbatlas_alert_configuration` + "`" + ` provides an Alert Configuration resource to define the conditions that trigger an alert and the methods of notification within a MongoDB Atlas project.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
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
					Description: `Value to test with the specified operator. If ` + "`" + `field_name` + "`" + ` is set to TYPE_NAME, you can match on the following values: - ` + "`" + `PRIMARY` + "`" + ` - ` + "`" + `SECONDARY` + "`" + ` - ` + "`" + `STANDALONE` + "`" + ` - ` + "`" + `CONFIG` + "`" + ` - ` + "`" + `MONGOS` + "`" + ` ### Metric Threshold Config (` + "`" + `metric_threshold_config` + "`" + `) The threshold that causes an alert to be triggered. Required if ` + "`" + `event_type_name` + "`" + ` : ` + "`" + `OUTSIDE_METRIC_THRESHOLD` + "`" + ` or ` + "`" + `OUTSIDE_SERVERLESS_METRIC_THRESHOLD` + "`" + ``,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `Name of the metric to check. The full list being quite large, please refer to atlas docs [here for general metrics](https://docs.atlas.mongodb.com/reference/alert-host-metrics/#measurement-types) and [here for serverless metrics](https://www.mongodb.com/docs/atlas/reference/api/alert-configurations-create-config/#serverless-measurements)`,
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
					Description: `The units for the threshold value. Depends on the type of metric. Refer to the [MongoDB API Alert Configuration documentation](https://www.mongodb.com/docs/atlas/reference/api/alert-configurations-get-config/#request-body-parameters) for a list of accepted values.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `This must be set to AVERAGE. Atlas computes the current metric value as an average. ### Threshold Config (` + "`" + `threshold_config` + "`" + `)`,
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
					Description: `The units for the threshold value. Depends on the type of metric. Refer to the [MongoDB API Alert Configuration documentation](https://www.mongodb.com/docs/atlas/reference/api/alert-configurations-get-config/#request-body-parameters) for a list of accepted values. ### Notifications List of notifications to send when an alert condition is detected.`,
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
					Description: `Flag indicating email notifications should be sent. This flag is only valid if ` + "`" + `type_name` + "`" + ` is set to ` + "`" + `ORG` + "`" + `, ` + "`" + `GROUP` + "`" + `, or ` + "`" + `USER` + "`" + `.`,
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
					Description: `Flag indicating if text message notifications should be sent to this user's mobile phone. This flag is only valid if ` + "`" + `type_name` + "`" + ` is set to ` + "`" + `ORG` + "`" + `, ` + "`" + `GROUP` + "`" + `, or ` + "`" + `USER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `Unique identifier of a team.`,
				},
				resource.Attribute{
					Name:        "team_name",
					Description: `Label for the team that receives this notification.`,
				},
				resource.Attribute{
					Name:        "type_name",
					Description: `Type of alert notification. Accepted values are: - ` + "`" + `DATADOG` + "`" + ` - ` + "`" + `EMAIL` + "`" + ` - ` + "`" + `FLOWDOCK` + "`" + ` - ` + "`" + `GROUP` + "`" + ` (Project) - ` + "`" + `OPS_GENIE` + "`" + ` - ` + "`" + `ORG` + "`" + ` - ` + "`" + `PAGER_DUTY` + "`" + ` - ` + "`" + `SLACK` + "`" + ` - ` + "`" + `SMS` + "`" + ` - ` + "`" + `TEAM` + "`" + ` - ` + "`" + `USER` + "`" + ` - ` + "`" + `VICTOR_OPS` + "`" + ` - ` + "`" + `WEBHOOK` + "`" + ` - ` + "`" + `MICROSOFT_TEAMS` + "`" + ``,
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
					Name:        "webhook_secret",
					Description: `Optional authentication secret for the ` + "`" + `WEBHOOK` + "`" + ` notifications type.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `Target URL for the ` + "`" + `WEBHOOK` + "`" + ` notifications type.`,
				},
				resource.Attribute{
					Name:        "microsoft_teams_webhook_url",
					Description: `Microsoft Teams Webhook Uniform Resource Locator (URL) that MongoDB Cloud needs to send this notification via Microsoft Teams. Required if ` + "`" + `type_name` + "`" + ` is ` + "`" + `MICROSOFT_TEAMS` + "`" + `. If the URL later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Optional. One or more roles that receive the configured alert. If you include this field, Atlas sends alerts only to users assigned the roles you specify in the array. If you omit this field, Atlas sends alerts to users assigned any role. This parameter is only valid if ` + "`" + `type_name` + "`" + ` is set to ` + "`" + `ORG` + "`" + `, ` + "`" + `GROUP` + "`" + `, or ` + "`" + `USER` + "`" + `. Accepted values are: | Project roles | Organization roles | |:---------- |:----------- | | ` + "`" + `GROUP_CHARTS_ADMIN` + "`" + ` | ` + "`" + `ORG_OWNER` + "`" + ` | | ` + "`" + `GROUP_CLUSTER_MANAGER` + "`" + ` | ` + "`" + `ORG_MEMBER` + "`" + ` | | ` + "`" + `GROUP_DATA_ACCESS_ADMIN` + "`" + ` | ` + "`" + `ORG_GROUP_CREATOR` + "`" + ` | | ` + "`" + `GROUP_DATA_ACCESS_READ_ONLY` + "`" + ` | ` + "`" + `ORG_BILLING_ADMIN` + "`" + ` | | ` + "`" + `GROUP_DATA_ACCESS_READ_WRITE` + "`" + ` | ` + "`" + `ORG_READ_ONLY` + "`" + ` | | ` + "`" + `GROUP_OWNER` + "`" + ` | | | ` + "`" + `GROUP_READ_ONLY` + "`" + ` | | ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
			Type:             "mongodbatlas_api_key",
			Category:         "Resources",
			ShortDescription: `Provides a API Key resource.`,
			Description: `

` + "`" + `mongodbatlas_api_key` + "`" + ` provides a Organization API key resource. This allows an Organizational API key to be created.

~> **IMPORTANT WARNING:** Managing Atlas Programmatic API Keys (PAKs) with Terraform will expose sensitive organizational secrets in Terraform's state. We suggest following [Terraform's best practices](https://developer.hashicorp.com/terraform/language/state/sensitive-data). You may also want to consider managing your PAKs via a more secure method, such as the [HashiCorp Vault MongoDB Atlas Secrets Engine](https://developer.hashicorp.com/vault/docs/secrets/mongodbatlas).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `Unique identifier for the organization whose API keys you want to retrieve. Use the /orgs endpoint to retrieve all organizations to which the authenticated user has access.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this Organization API key.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `Name of the role. This resource returns all the roles the user has in Atlas. The following are valid roles:`,
				},
				resource.Attribute{
					Name:        "api_key_id",
					Description: `(Required) The unique identifier of the Programmatic API key you want to associate with the Project. The Programmatic API key and Project must share the same parent organization. Note: this is not the ` + "`" + `publicKey` + "`" + ` of the Programmatic API key but the ` + "`" + `id` + "`" + ` of the key. See [Programmatic API Keys](https://docs.atlas.mongodb.com/reference/api/apiKeys/) for more.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `(Required) List of Project roles that the Programmatic API key needs to have. Ensure you provide: at least one role and ensure all roles are valid for the Project. You must specify an array even if you are only associating a single role with the Programmatic API key. The following are valid roles:`,
				},
				resource.Attribute{
					Name:        "api_key_id",
					Description: `Unique identifier for this Organization API key. ## Import API Keys must be imported using org ID, API Key ID e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_api_key.test 5d09d6a59ccf6445652a444a-6576974933969669 ` + "`" + `` + "`" + `` + "`" + ` See [MongoDB Atlas API - API Key](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Programmatic-API-Keys/operation/createOneOrganizationApiKey) - Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_key_id",
					Description: `Unique identifier for this Organization API key. ## Import API Keys must be imported using org ID, API Key ID e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_api_key.test 5d09d6a59ccf6445652a444a-6576974933969669 ` + "`" + `` + "`" + `` + "`" + ` See [MongoDB Atlas API - API Key](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Programmatic-API-Keys/operation/createOneOrganizationApiKey) - Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_auditing",
			Category:         "Resources",
			ShortDescription: `Provides a Auditing resource.`,
			Description: `

` + "`" + `mongodbatlas_auditing` + "`" + ` provides an Auditing resource. This allows auditing to be created.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to configure auditing.`,
				},
				resource.Attribute{
					Name:        "audit_authorization_success",
					Description: `Indicates whether the auditing system captures successful authentication attempts for audit filters using the "atype" : "authCheck" auditing event. For more information, see [auditAuthorizationSuccess](https://docs.mongodb.com/manual/reference/parameters/#param.auditAuthorizationSuccess).`,
				},
				resource.Attribute{
					Name:        "audit_filter",
					Description: `JSON-formatted audit filter. For complete documentation on custom auditing filters, see [Configure Audit Filters](https://docs.mongodb.com/manual/tutorial/configure-audit-filters/).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Denotes whether or not the project associated with the {project_id} has database auditing enabled. Defaults to false. ~>`,
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
			Type:             "mongodbatlas_backup_compliance_policy",
			Category:         "Resources",
			ShortDescription: `Provides a Backup Compliance Policy resource.`,
			Description: `

` + "`" + `mongodbatlas_backup_compliance_policy` + "`" + ` provides a resource that enables you to set up a Backup Compliance Policy resource. [Backup Compliance Policy ](https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy) prevents any user, regardless of role, from modifying or deleting specific cluster settings, backups, and backup configurations. When enabled, the Backup Compliance Policy will be applied as the minimum policy for all clusters and backups in the project. It can only be disabled by contacting MongoDB support. This feature is only supported for cluster tiers M10+.

When enabled, the Backup Compliance Policy will be applied as the minimum backup policy to all clusters in a project and will protect all existing snapshots. This will prevent any user, regardless of role, from modifying or deleting existing snapshots prior to expiration. Changes made to existing backup policies will only apply to future snapshots.

-> **NOTE:** Groups and projects are synonymous terms. You might find ` + "`" + `groupId` + "`" + ` in the official documentation.

-> **IMPORTANT NOTE:** Once enable a Backup Compliance Policy, no user, regardless of role, can disable the Backup Compliance Policy via Terraform, or any other method, without contacting MongoDB support.   This means that once enabled some resources defined in Terraform will no longer be modifiable.   See the full list of [Backup Compliance Policy Prohibited Actions and Considerations](https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/#configure-a-backup-compliance-policy)

-> **NOTE:** With Backup Compliance Policy enabled, cluster backups are retained after a cluster is deleted and backups can be used normally until retention expiration. When the Backup Compliance Policy is not enabled, Atlas deletes the cluster's associated backup snapshots when a cluster is terminated. By default, a Backup Compliance Policy is not enabled. For more details see [Back Up, Restore, and Archive Data](https://www.mongodb.com/docs/atlas/backup-restore-cluster/). 

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies your project.`,
				},
				resource.Attribute{
					Name:        "authorized_email",
					Description: `Email address of a security or legal representative for the Backup Compliance Policy who is authorized to update the Backup Compliance Policy settings.`,
				},
				resource.Attribute{
					Name:        "copy_protection_enabled",
					Description: `Flag that indicates whether to enable additional backup copies for the cluster. If unspecified, this value defaults to false.`,
				},
				resource.Attribute{
					Name:        "pit_enabled",
					Description: `Flag that indicates whether the cluster uses Continuous Cloud Backups with a Backup Compliance Policy. If unspecified, this value defaults to false.`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_enabled",
					Description: `Flag that indicates whether Encryption at Rest using Customer Key Management is required for all clusters with a Backup Compliance Policy. If unspecified, this value defaults to false.`,
				},
				resource.Attribute{
					Name:        "restore_window_days",
					Description: `Number of previous days that you can restore back to with Continuous Cloud Backup with a Backup Compliance Policy. You must specify a positive, non-zero integer, and the maximum retention window can't exceed the hourly retention time. This parameter applies only to Continuous Cloud Backups with a Backup Compliance Policy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "updated_date",
					Description: `ISO 8601 timestamp format in UTC that indicates when the user updated the Data Protection Policy settings. MongoDB Cloud ignores this setting when you enable or update the Backup Compliance Policy settings.`,
				},
				resource.Attribute{
					Name:        "updated_user",
					Description: `Email address that identifies the user who updated the Backup Compliance Policy settings. MongoDB Cloud ignores this email setting when you enable or update the Backup Compliance Policy settings. ### On Demand Policy Item`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the backup policy item.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the backup policy item. For hourly policies, the frequency type is defined as ` + "`" + `ondemand` + "`" + `. Note that this is a read-only value and not required in plan files - its value is implied from the policy resource type.`,
				},
				resource.Attribute{
					Name:        "frequency_interval",
					Description: `Desired frequency of the new backup policy item specified by ` + "`" + `frequency_type` + "`" + ` (hourly in this case). The supported values for hourly policies are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `8` + "`" + ` or ` + "`" + `12` + "`" + ` hours. Note that ` + "`" + `12` + "`" + ` hours is the only accepted value for NVMe clusters.`,
				},
				resource.Attribute{
					Name:        "retention_unit",
					Description: `Scope of the backup policy item: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, or ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_value",
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. ### Policy Item Hourly`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the backup policy item.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the backup policy item. For hourly policies, the frequency type is defined as ` + "`" + `hourly` + "`" + `. Note that this is a read-only value and not required in plan files - its value is implied from the policy resource type.`,
				},
				resource.Attribute{
					Name:        "frequency_interval",
					Description: `Desired frequency of the new backup policy item specified by ` + "`" + `frequency_type` + "`" + ` (hourly in this case). The supported values for hourly policies are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `8` + "`" + ` or ` + "`" + `12` + "`" + ` hours. Note that ` + "`" + `12` + "`" + ` hours is the only accepted value for NVMe clusters.`,
				},
				resource.Attribute{
					Name:        "retention_unit",
					Description: `Scope of the backup policy item: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, or ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_value",
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. ### Policy Item Daily`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the backup policy item.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the backup policy item. For daily policies, the frequency type is defined as ` + "`" + `daily` + "`" + `. Note that this is a read-only value and not required in plan files - its value is implied from the policy resource type.`,
				},
				resource.Attribute{
					Name:        "frequency_interval",
					Description: `Desired frequency of the new backup policy item specified by ` + "`" + `frequency_type` + "`" + ` (daily in this case). The only supported value for daily policies is ` + "`" + `1` + "`" + ` day.`,
				},
				resource.Attribute{
					Name:        "retention_unit",
					Description: `Scope of the backup policy item: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, or ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_value",
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. Note that for less frequent policy items, Atlas requires that you specify a retention period greater than or equal to the retention period specified for more frequent policy items. For example: If the hourly policy item specifies a retention of two days, the daily retention policy must specify two days or greater. ### Policy Item Weekly`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the backup policy item.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the backup policy item. For weekly policies, the frequency type is defined as ` + "`" + `weekly` + "`" + `. Note that this is a read-only value and not required in plan files - its value is implied from the policy resource type.`,
				},
				resource.Attribute{
					Name:        "frequency_interval",
					Description: `Desired frequency of the new backup policy item specified by ` + "`" + `frequency_type` + "`" + ` (weekly in this case). The supported values for weekly policies are ` + "`" + `1` + "`" + ` through ` + "`" + `7` + "`" + `, where ` + "`" + `1` + "`" + ` represents Monday and ` + "`" + `7` + "`" + ` represents Sunday.`,
				},
				resource.Attribute{
					Name:        "retention_unit",
					Description: `Scope of the backup policy item: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, or ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_value",
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. Weekly policy must have retention of at least 7 days or 1 week. Note that for less frequent policy items, Atlas requires that you specify a retention period greater than or equal to the retention period specified for more frequent policy items. For example: If the daily policy item specifies a retention of two weeks, the weekly retention policy must specify two weeks or greater. ### Policy Item Monthly`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the backup policy item.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the backup policy item. For monthly policies, the frequency type is defined as ` + "`" + `monthly` + "`" + `. Note that this is a read-only value and not required in plan files - its value is implied from the policy resource type.`,
				},
				resource.Attribute{
					Name:        "frequency_interval",
					Description: `Desired frequency of the new backup policy item specified by ` + "`" + `frequency_type` + "`" + ` (monthly in this case). The supported values for weekly policies are`,
				},
				resource.Attribute{
					Name:        "retention_unit",
					Description: `Scope of the backup policy item: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, or ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_value",
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. Monthly policy must have retention days of at least 31 days or 5 weeks or 1 month. Note that for less frequent policy items, Atlas requires that you specify a retention period greater than or equal to the retention period specified for more frequent policy items. For example: If the weekly policy item specifies a retention of two weeks, the montly retention policy must specify two weeks or greater. ## Import Backup Compliance Policy entries can be imported using project project_id in the format ` + "`" + `project_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_backup_compliance_policy.backup_policy 5d0f1f73cf09a29120e173cf ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Cloud-Backups/operation/updateDataProtectionSettings) and [Backup Compliance Policy Prohibited Actions](https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/#prohibited-actions).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "updated_date",
					Description: `ISO 8601 timestamp format in UTC that indicates when the user updated the Data Protection Policy settings. MongoDB Cloud ignores this setting when you enable or update the Backup Compliance Policy settings.`,
				},
				resource.Attribute{
					Name:        "updated_user",
					Description: `Email address that identifies the user who updated the Backup Compliance Policy settings. MongoDB Cloud ignores this email setting when you enable or update the Backup Compliance Policy settings. ### On Demand Policy Item`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the backup policy item.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the backup policy item. For hourly policies, the frequency type is defined as ` + "`" + `ondemand` + "`" + `. Note that this is a read-only value and not required in plan files - its value is implied from the policy resource type.`,
				},
				resource.Attribute{
					Name:        "frequency_interval",
					Description: `Desired frequency of the new backup policy item specified by ` + "`" + `frequency_type` + "`" + ` (hourly in this case). The supported values for hourly policies are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `8` + "`" + ` or ` + "`" + `12` + "`" + ` hours. Note that ` + "`" + `12` + "`" + ` hours is the only accepted value for NVMe clusters.`,
				},
				resource.Attribute{
					Name:        "retention_unit",
					Description: `Scope of the backup policy item: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, or ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_value",
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. ### Policy Item Hourly`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the backup policy item.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the backup policy item. For hourly policies, the frequency type is defined as ` + "`" + `hourly` + "`" + `. Note that this is a read-only value and not required in plan files - its value is implied from the policy resource type.`,
				},
				resource.Attribute{
					Name:        "frequency_interval",
					Description: `Desired frequency of the new backup policy item specified by ` + "`" + `frequency_type` + "`" + ` (hourly in this case). The supported values for hourly policies are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `8` + "`" + ` or ` + "`" + `12` + "`" + ` hours. Note that ` + "`" + `12` + "`" + ` hours is the only accepted value for NVMe clusters.`,
				},
				resource.Attribute{
					Name:        "retention_unit",
					Description: `Scope of the backup policy item: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, or ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_value",
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. ### Policy Item Daily`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the backup policy item.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the backup policy item. For daily policies, the frequency type is defined as ` + "`" + `daily` + "`" + `. Note that this is a read-only value and not required in plan files - its value is implied from the policy resource type.`,
				},
				resource.Attribute{
					Name:        "frequency_interval",
					Description: `Desired frequency of the new backup policy item specified by ` + "`" + `frequency_type` + "`" + ` (daily in this case). The only supported value for daily policies is ` + "`" + `1` + "`" + ` day.`,
				},
				resource.Attribute{
					Name:        "retention_unit",
					Description: `Scope of the backup policy item: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, or ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_value",
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. Note that for less frequent policy items, Atlas requires that you specify a retention period greater than or equal to the retention period specified for more frequent policy items. For example: If the hourly policy item specifies a retention of two days, the daily retention policy must specify two days or greater. ### Policy Item Weekly`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the backup policy item.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the backup policy item. For weekly policies, the frequency type is defined as ` + "`" + `weekly` + "`" + `. Note that this is a read-only value and not required in plan files - its value is implied from the policy resource type.`,
				},
				resource.Attribute{
					Name:        "frequency_interval",
					Description: `Desired frequency of the new backup policy item specified by ` + "`" + `frequency_type` + "`" + ` (weekly in this case). The supported values for weekly policies are ` + "`" + `1` + "`" + ` through ` + "`" + `7` + "`" + `, where ` + "`" + `1` + "`" + ` represents Monday and ` + "`" + `7` + "`" + ` represents Sunday.`,
				},
				resource.Attribute{
					Name:        "retention_unit",
					Description: `Scope of the backup policy item: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, or ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_value",
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. Weekly policy must have retention of at least 7 days or 1 week. Note that for less frequent policy items, Atlas requires that you specify a retention period greater than or equal to the retention period specified for more frequent policy items. For example: If the daily policy item specifies a retention of two weeks, the weekly retention policy must specify two weeks or greater. ### Policy Item Monthly`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the backup policy item.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the backup policy item. For monthly policies, the frequency type is defined as ` + "`" + `monthly` + "`" + `. Note that this is a read-only value and not required in plan files - its value is implied from the policy resource type.`,
				},
				resource.Attribute{
					Name:        "frequency_interval",
					Description: `Desired frequency of the new backup policy item specified by ` + "`" + `frequency_type` + "`" + ` (monthly in this case). The supported values for weekly policies are`,
				},
				resource.Attribute{
					Name:        "retention_unit",
					Description: `Scope of the backup policy item: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, or ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_value",
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. Monthly policy must have retention days of at least 31 days or 5 weeks or 1 month. Note that for less frequent policy items, Atlas requires that you specify a retention period greater than or equal to the retention period specified for more frequent policy items. For example: If the weekly policy item specifies a retention of two weeks, the montly retention policy must specify two weeks or greater. ## Import Backup Compliance Policy entries can be imported using project project_id in the format ` + "`" + `project_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_backup_compliance_policy.backup_policy 5d0f1f73cf09a29120e173cf ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Cloud-Backups/operation/updateDataProtectionSettings) and [Backup Compliance Policy Prohibited Actions](https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/#prohibited-actions).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_backup_schedule",
			Category:         "Resources",
			ShortDescription: `Provides a Cloud Backup Schedule resource.`,
			Description: `

` + "`" + `mongodbatlas_cloud_backup_schedule` + "`" + ` provides a cloud backup schedule resource. The resource lets you create, read, update and delete a cloud backup schedule.

-> **NOTE** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

-> **API Key Access List** This resource requires an Atlas API Access Key List to utilize this feature. This means to manage this resources you must have the IP address or CIDR block that the Terraform connection is coming from added to the Atlas API Key Access List of the Atlas API key you are using. See [Resources that require API Key List](https://www.mongodb.com/docs/atlas/configure-api-access/#use-api-resources-that-require-an-access-list) for details.

-> **NOTE:** If Backup Compliance Policy is enabled for the project for which this backup schedule is defined, you cannot modify the backup schedule for an individual cluster below the minimum requirements set in the Backup Compliance Policy.  See [Backup Compliance Policy Prohibited Actions and Considerations](https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/#configure-a-backup-compliance-policy).

In the Terraform MongoDB Atlas Provider 1.0.0 we have re-architected the way in which Cloud Backup Policies are manged with Terraform to significantly reduce the complexity. Due to this change we've provided multiple examples below to help express how this new resource functions.


`,
			Keywords: []string{},
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
					Description: `(Optional) UTC Minutes after ` + "`" + `reference_hour_of_day` + "`" + ` that Atlas takes snapshots for backup policy items. Must be between 0 and 59, inclusive.`,
				},
				resource.Attribute{
					Name:        "restore_window_days",
					Description: `(Optional) Number of days back in time you can restore to with point-in-time accuracy. Must be a positive, non-zero integer.`,
				},
				resource.Attribute{
					Name:        "update_snapshots",
					Description: `(Optional) Specify true to apply the retention changes in the updated backup policy to snapshots that Atlas took previously.`,
				},
				resource.Attribute{
					Name:        "policy_item_hourly",
					Description: `(Optional) Hourly policy item`,
				},
				resource.Attribute{
					Name:        "policy_item_daily",
					Description: `(Optional) Daily policy item`,
				},
				resource.Attribute{
					Name:        "policy_item_weekly",
					Description: `(Optional) Weekly policy item`,
				},
				resource.Attribute{
					Name:        "policy_item_monthly",
					Description: `(Optional) Monthly policy item`,
				},
				resource.Attribute{
					Name:        "auto_export_enabled",
					Description: `Flag that indicates whether automatic export of cloud backup snapshots to the AWS bucket is enabled. Value can be one of the following: true - enables automatic export of cloud backup snapshots to the AWS bucket false - disables automatic export of cloud backup snapshots to the AWS bucket (default)`,
				},
				resource.Attribute{
					Name:        "use_org_and_group_names_in_export_prefix",
					Description: `Specify true to use organization and project names instead of organization and project UUIDs in the path for the metadata files that Atlas uploads to your S3 bucket after it finishes exporting the snapshots. To learn more about the metadata files that Atlas uploads, see [Export Cloud Backup Snapshot](https://www.mongodb.com/docs/atlas/backup/cloud-backup/export/#std-label-cloud-provider-snapshot-export). ### Export`,
				},
				resource.Attribute{
					Name:        "export_bucket_id",
					Description: `Unique identifier of the mongodbatlas_cloud_backup_snapshot_export_bucket export_bucket_id value.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the export snapshot item. ### Policy Item Hourly`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the backup policy item.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the backup policy item. For hourly policies, the frequency type is defined as ` + "`" + `hourly` + "`" + `. Note that this is a read-only value and not required in plan files - its value is implied from the policy resource type.`,
				},
				resource.Attribute{
					Name:        "frequency_interval",
					Description: `Desired frequency of the new backup policy item specified by ` + "`" + `frequency_type` + "`" + ` (hourly in this case). The supported values for hourly policies are ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `8` + "`" + ` or ` + "`" + `12` + "`" + ` hours. Note that ` + "`" + `12` + "`" + ` hours is the only accepted value for NVMe clusters.`,
				},
				resource.Attribute{
					Name:        "retention_unit",
					Description: `Scope of the backup policy item: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, or ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_value",
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. ### Policy Item Daily`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the backup policy item.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the backup policy item. For daily policies, the frequency type is defined as ` + "`" + `daily` + "`" + `. Note that this is a read-only value and not required in plan files - its value is implied from the policy resource type.`,
				},
				resource.Attribute{
					Name:        "frequency_interval",
					Description: `Desired frequency of the new backup policy item specified by ` + "`" + `frequency_type` + "`" + ` (daily in this case). The only supported value for daily policies is ` + "`" + `1` + "`" + ` day.`,
				},
				resource.Attribute{
					Name:        "retention_unit",
					Description: `Scope of the backup policy item: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, or ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_value",
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. Note that for less frequent policy items, Atlas requires that you specify a retention period greater than or equal to the retention period specified for more frequent policy items. For example: If the hourly policy item specifies a retention of two days, the daily retention policy must specify two days or greater. ### Policy Item Weekly`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the backup policy item.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the backup policy item. For weekly policies, the frequency type is defined as ` + "`" + `weekly` + "`" + `. Note that this is a read-only value and not required in plan files - its value is implied from the policy resource type.`,
				},
				resource.Attribute{
					Name:        "frequency_interval",
					Description: `Desired frequency of the new backup policy item specified by ` + "`" + `frequency_type` + "`" + ` (weekly in this case). The supported values for weekly policies are ` + "`" + `1` + "`" + ` through ` + "`" + `7` + "`" + `, where ` + "`" + `1` + "`" + ` represents Monday and ` + "`" + `7` + "`" + ` represents Sunday.`,
				},
				resource.Attribute{
					Name:        "retention_unit",
					Description: `Scope of the backup policy item: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, or ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_value",
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. Weekly policy must have retention of at least 7 days or 1 week. Note that for less frequent policy items, Atlas requires that you specify a retention period greater than or equal to the retention period specified for more frequent policy items. For example: If the daily policy item specifies a retention of two weeks, the weekly retention policy must specify two weeks or greater. ### Policy Item Monthly`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the backup policy item.`,
				},
				resource.Attribute{
					Name:        "frequency_type",
					Description: `Frequency associated with the backup policy item. For monthly policies, the frequency type is defined as ` + "`" + `monthly` + "`" + `. Note that this is a read-only value and not required in plan files - its value is implied from the policy resource type.`,
				},
				resource.Attribute{
					Name:        "frequency_interval",
					Description: `Desired frequency of the new backup policy item specified by ` + "`" + `frequency_type` + "`" + ` (monthly in this case). The supported values for weekly policies are`,
				},
				resource.Attribute{
					Name:        "retention_unit",
					Description: `Scope of the backup policy item: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, or ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_value",
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. Monthly policy must have retention days of at least 31 days or 5 weeks or 1 month. Note that for less frequent policy items, Atlas requires that you specify a retention period greater than or equal to the retention period specified for more frequent policy items. For example: If the weekly policy item specifies a retention of two weeks, the montly retention policy must specify two weeks or greater. ### Snapshot Distribution`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Required) Human-readable label that identifies the cloud provider that stores the snapshot copy. i.e. "AWS" "AZURE" "GCP"`,
				},
				resource.Attribute{
					Name:        "frequencies",
					Description: `(Required) List that describes which types of snapshots to copy. i.e. "HOURLY" "DAILY" "WEEKLY" "MONTHLY" "ON_DEMAND"`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Required) Target region to copy snapshots belonging to replicationSpecId to. Please supply the 'Atlas Region' which can be found under https://www.mongodb.com/docs/atlas/reference/cloud-providers/ 'regions' link`,
				},
				resource.Attribute{
					Name:        "should_copy_oplogs",
					Description: `(Required) Flag that indicates whether to copy the oplogs to the target region. You can use the oplogs to perform point-in-time restores. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Unique identifier of the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "next_snapshot",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch when Atlas takes the next snapshot.`,
				},
				resource.Attribute{
					Name:        "id_policy",
					Description: `Unique identifier of the backup policy. ## Import Cloud Backup Schedule entries can be imported using project_id and cluster_name, in the format ` + "`" + `PROJECTID-CLUSTERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_backup_schedule.test 5d0f1f73cf09a29120e173cf-MyClusterTest ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/schedule/modify-one-schedule/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Unique identifier of the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "next_snapshot",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch when Atlas takes the next snapshot.`,
				},
				resource.Attribute{
					Name:        "id_policy",
					Description: `Unique identifier of the backup policy. ## Import Cloud Backup Schedule entries can be imported using project_id and cluster_name, in the format ` + "`" + `PROJECTID-CLUSTERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_backup_schedule.test 5d0f1f73cf09a29120e173cf-MyClusterTest ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/schedule/modify-one-schedule/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_backup_snapshot",
			Category:         "Resources",
			ShortDescription: `Provides a Cloud Backup Snapshot resource.`,
			Description: `

` + "`" + `mongodbatlas_cloud_backup_snapshot` + "`" + ` provides a resource to take a cloud backup snapshot on demand.
On-demand snapshots happen immediately, unlike scheduled snapshots which occur at regular intervals. If there is already an on-demand snapshot with a status of queued or inProgress, you must wait until Atlas has completed the on-demand snapshot before taking another.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

-> **NOTE:** If Backup Compliance Policy is enabled for the project for which this backup schedule is defined, you cannot delete a backup snapshot or decrease the retention time for a snapshot after it's taken.  See [Backup Compliance Policy Prohibited Actions and Considerations](https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/#configure-a-backup-compliance-policy).

-> **API Key Access List**: This resource requires an Atlas API Access Key List to utilize this feature. This means to manage this resources you must have the IP address or CIDR block that the Terraform connection is coming from added to the Atlas API Key Access List of the Atlas API key you are using. See [Resources that require API Key List](https://www.mongodb.com/docs/atlas/configure-api-access/#use-api-resources-that-require-an-access-list) for details.

`,
			Keywords: []string{},
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
					Description: `Specifies the type of cluster: replicaSet or shardedCluster.`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `Cloud provider that stores this snapshot. Atlas returns this parameter when ` + "`" + `type` + "`" + ` is ` + "`" + `replicaSet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Block of List of snapshots and the cloud provider where the snapshots are stored. Atlas returns this parameter when ` + "`" + `type` + "`" + ` is ` + "`" + `shardedCluster` + "`" + `. See below`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `Label given to the replica set from which Atlas took this snapshot. Atlas returns this parameter when ` + "`" + `type` + "`" + ` is ` + "`" + `replicaSet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `Unique identifiers of the snapshots created for the shards and config server for a sharded cluster. Atlas returns this parameter when ` + "`" + `type` + "`" + ` is ` + "`" + `shardedCluster` + "`" + `. These identifiers should match those given in the ` + "`" + `members[n].id` + "`" + ` parameters. This allows you to map a snapshot to its shard or config server name. ### members`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `Cloud provider that stores this snapshot.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier for the sharded cluster snapshot.`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `Label given to a shard or config server from which Atlas took this snapshot. ## Import Cloud Backup Snapshot entries can be imported using project project_id, cluster_name and snapshot_id (Unique identifier of the snapshot), in the format ` + "`" + `PROJECTID-CLUSTERNAME-SNAPSHOTID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_backup_snapshot.test 5d0f1f73cf09a29120e173cf-MyClusterTest-5d116d82014b764445b2f9b5 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/backups/)`,
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
					Description: `Specifies the type of cluster: replicaSet or shardedCluster.`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `Cloud provider that stores this snapshot. Atlas returns this parameter when ` + "`" + `type` + "`" + ` is ` + "`" + `replicaSet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Block of List of snapshots and the cloud provider where the snapshots are stored. Atlas returns this parameter when ` + "`" + `type` + "`" + ` is ` + "`" + `shardedCluster` + "`" + `. See below`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `Label given to the replica set from which Atlas took this snapshot. Atlas returns this parameter when ` + "`" + `type` + "`" + ` is ` + "`" + `replicaSet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `Unique identifiers of the snapshots created for the shards and config server for a sharded cluster. Atlas returns this parameter when ` + "`" + `type` + "`" + ` is ` + "`" + `shardedCluster` + "`" + `. These identifiers should match those given in the ` + "`" + `members[n].id` + "`" + ` parameters. This allows you to map a snapshot to its shard or config server name. ### members`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `Cloud provider that stores this snapshot.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier for the sharded cluster snapshot.`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `Label given to a shard or config server from which Atlas took this snapshot. ## Import Cloud Backup Snapshot entries can be imported using project project_id, cluster_name and snapshot_id (Unique identifier of the snapshot), in the format ` + "`" + `PROJECTID-CLUSTERNAME-SNAPSHOTID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_backup_snapshot.test 5d0f1f73cf09a29120e173cf-MyClusterTest-5d116d82014b764445b2f9b5 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/backups/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_backup_snapshot_export_bucket",
			Category:         "Resources",
			ShortDescription: `Provides a Cloud Backup Snapshot Export Bucket resource.`,
			Description: `
` + "`" + `mongodbatlas_cloud_backup_snapshot_export_bucket` + "`" + ` resource allows you to create an export snapshot bucket for the specified project. 


-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

**API Key Access List**: Some Atlas API resources such as Cloud Backup Restores, Cloud Backup Snapshots, and Cloud Backup Schedules **require** an Atlas API Key Access List to utilize these feature.  Hence, if using Terraform, or any other programmatic control, to manage these resources you must have the IP address or CIDR block that the connection is coming from added to the Atlas API Key Access List of the Atlas API key you are using.   See [Resources that require API Key List](https://www.mongodb.com/docs/atlas/configure-api-access/#use-api-resources-that-require-an-access-list)
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "iam_role_id",
					Description: `(Required) Unique identifier of the role that Atlas can use to access the bucket. You must also specify the ` + "`" + `bucket_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) Name of the bucket that the provided role ID is authorized to access. You must also specify the ` + "`" + `iam_role_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Required) Name of the provider of the cloud service where Atlas can access the S3 bucket. Atlas only supports ` + "`" + `AWS` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "export_bucket_id",
					Description: `Unique identifier of the snapshot export bucket. ## Import Cloud Backup Snapshot Export Backup entries can be imported using project project_id, and bucket_id (Unique identifier of the snapshot export bucket), in the format ` + "`" + `PROJECTID-BUCKETID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_backup_snapshot_export_bucket.test 5d0f1f73cf09a29120e173cf-5d116d82014b764445b2f9b5 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/export/create-one-export-bucket/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "export_bucket_id",
					Description: `Unique identifier of the snapshot export bucket. ## Import Cloud Backup Snapshot Export Backup entries can be imported using project project_id, and bucket_id (Unique identifier of the snapshot export bucket), in the format ` + "`" + `PROJECTID-BUCKETID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_backup_snapshot_export_bucket.test 5d0f1f73cf09a29120e173cf-5d116d82014b764445b2f9b5 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/export/create-one-export-bucket/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_backup_snapshot_export_job",
			Category:         "Resources",
			ShortDescription: `Provides a Cloud Backup Snapshot Export Job resource.`,
			Description: `
` + "`" + `mongodbatlas_cloud_backup_snapshot_export_job` + "`" + ` resource allows you to create a cloud backup snapshot export job for the specified project. 


-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

**API Key Access List**: Some Atlas API resources such as Cloud Backup Restores, Cloud Backup Snapshots, and Cloud Backup Schedules **require** an Atlas API Key Access List to utilize these feature.  Hence, if using Terraform, or any other programmatic control, to manage these resources you must have the IP address or CIDR block that the connection is coming from added to the Atlas API Key Access List of the Atlas API key you are using.   See [Resources that require API Key List](https://www.mongodb.com/docs/atlas/configure-api-access/#use-api-resources-that-require-an-access-list)
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the project which contains the Atlas cluster whose snapshot you want to export.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) Name of the Atlas cluster whose snapshot you want to export.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Required) Unique identifier of the Cloud Backup snapshot to export. If necessary, use the [Get All Cloud Backups](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/get-all-backups/) API to retrieve the list of snapshot IDs for a cluster or use the data source [mongodbatlas_cloud_cloud_backup_snapshots](https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/data-sources/cloud_backup_snapshots)`,
				},
				resource.Attribute{
					Name:        "export_bucket_id",
					Description: `(Required) Unique identifier of the AWS bucket to export the Cloud Backup snapshot to. If necessary, use the [Get All Snapshot Export Buckets](https://docs.atlas.mongodb.com/reference/api/cloud-backup/export/get-all-export-buckets/) API to retrieve the IDs of all available export buckets for a project or use the data source [mongodbatlas_cloud_backup_snapshot_export_buckets](https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/data-sources/backup_snapshot_export_buckets)`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `(Optional) Custom data to include in the metadata file named ` + "`" + `.complete` + "`" + ` that Atlas uploads to the bucket when the export job finishes. Custom data can be specified as key and value pairs. ### Custom Data`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Required if you want to include custom data using ` + "`" + `custom_data` + "`" + ` in the metadata file uploaded to the bucket. Key to include in the metadata file that Atlas uploads to the bucket when the export job finishes.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Required if you specify ` + "`" + `key` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "components",
					Description: `_Returned for sharded clusters only._ Export job details for each replica set in the sharded cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp in ISO 8601 date and time format in UTC when the export job was created.`,
				},
				resource.Attribute{
					Name:        "err_msg",
					Description: `Error message, only if the export job failed.`,
				},
				resource.Attribute{
					Name:        "export_status",
					Description: `_Returned for replica set only._ Status of the export job.`,
				},
				resource.Attribute{
					Name:        "finished_at",
					Description: `Timestamp in ISO 8601 date and time format in UTC when the export job completes.`,
				},
				resource.Attribute{
					Name:        "export_job_id",
					Description: `Unique identifier of the export job.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `Full path on the cloud provider bucket to the folder where the snapshot is exported. The path is in the following format:` + "`" + `/exported_snapshots/{ORG-NAME}/{PROJECT-NAME}/{CLUSTER-NAME}/{SNAPSHOT-INITIATION-DATE}/{TIMESTAMP}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Status of the export job. Value can be one of the following:`,
				},
				resource.Attribute{
					Name:        "Queued",
					Description: `indicates that the export job is queued`,
				},
				resource.Attribute{
					Name:        "InProgress",
					Description: `indicates that the snapshot is being exported`,
				},
				resource.Attribute{
					Name:        "Successful",
					Description: `indicates that the export job has completed successfully`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `indicates that the export job has failed ### components`,
				},
				resource.Attribute{
					Name:        "export_id",
					Description: `_Returned for sharded clusters only._ Export job details for each replica set in the sharded cluster.`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `_Returned for sharded clusters only._ Unique identifier of the export job for the replica set. ### export_status`,
				},
				resource.Attribute{
					Name:        "exported_collections",
					Description: `_Returned for replica set only._ Number of collections that have been exported.`,
				},
				resource.Attribute{
					Name:        "total_collections",
					Description: `_Returned for replica set only._ Total number of collections to export. ## Import Cloud Backup Snapshot Export Backup entries can be imported using project project_id, cluster_name and export_job_id (Unique identifier of the snapshot export job), in the format ` + "`" + `PROJECTID-CLUSTERNAME-EXPORTJOBID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_backup_snapshot_export_job.test 5d0f1f73cf09a29120e173cf-5d116d82014b764445b2f9b5-5d116d82014b764445b2f9b5 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/export/create-one-export-job/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "components",
					Description: `_Returned for sharded clusters only._ Export job details for each replica set in the sharded cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp in ISO 8601 date and time format in UTC when the export job was created.`,
				},
				resource.Attribute{
					Name:        "err_msg",
					Description: `Error message, only if the export job failed.`,
				},
				resource.Attribute{
					Name:        "export_status",
					Description: `_Returned for replica set only._ Status of the export job.`,
				},
				resource.Attribute{
					Name:        "finished_at",
					Description: `Timestamp in ISO 8601 date and time format in UTC when the export job completes.`,
				},
				resource.Attribute{
					Name:        "export_job_id",
					Description: `Unique identifier of the export job.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `Full path on the cloud provider bucket to the folder where the snapshot is exported. The path is in the following format:` + "`" + `/exported_snapshots/{ORG-NAME}/{PROJECT-NAME}/{CLUSTER-NAME}/{SNAPSHOT-INITIATION-DATE}/{TIMESTAMP}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Status of the export job. Value can be one of the following:`,
				},
				resource.Attribute{
					Name:        "Queued",
					Description: `indicates that the export job is queued`,
				},
				resource.Attribute{
					Name:        "InProgress",
					Description: `indicates that the snapshot is being exported`,
				},
				resource.Attribute{
					Name:        "Successful",
					Description: `indicates that the export job has completed successfully`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `indicates that the export job has failed ### components`,
				},
				resource.Attribute{
					Name:        "export_id",
					Description: `_Returned for sharded clusters only._ Export job details for each replica set in the sharded cluster.`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `_Returned for sharded clusters only._ Unique identifier of the export job for the replica set. ### export_status`,
				},
				resource.Attribute{
					Name:        "exported_collections",
					Description: `_Returned for replica set only._ Number of collections that have been exported.`,
				},
				resource.Attribute{
					Name:        "total_collections",
					Description: `_Returned for replica set only._ Total number of collections to export. ## Import Cloud Backup Snapshot Export Backup entries can be imported using project project_id, cluster_name and export_job_id (Unique identifier of the snapshot export job), in the format ` + "`" + `PROJECTID-CLUSTERNAME-EXPORTJOBID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_cloud_backup_snapshot_export_job.test 5d0f1f73cf09a29120e173cf-5d116d82014b764445b2f9b5-5d116d82014b764445b2f9b5 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/export/create-one-export-job/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_backup_snapshot_restore_job",
			Category:         "Resources",
			ShortDescription: `Provides a Cloud Backup Snapshot Restore Job resource.`,
			Description: `

` + "`" + `mongodbatlas_cloud_backup_snapshot_restore_job` + "`" + ` provides a resource to create a new restore job from a cloud backup snapshot of a specified cluster. The restore job must define one of three delivery types: 
* **automated:** Atlas automatically restores the snapshot with snapshotId to the Atlas cluster with name targetClusterName in the Atlas project with targetGroupId.

* **download:** Atlas provides a URL to download a .tar.gz of the snapshot with snapshotId. The contents of the archive contain the data files for your Atlas cluster.

* **pointInTime:**  Atlas performs a Continuous Cloud Backup restore.

-> **Important:** If you specify ` + "`" + `deliveryType` + "`" + ` : ` + "`" + `automated` + "`" + ` or ` + "`" + `deliveryType` + "`" + ` : ` + "`" + `pointInTime` + "`" + ` in your request body to create an automated restore job, Atlas removes all existing data on the target cluster prior to the restore.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

-> **API Key Access List**: This resource requires an Atlas API Access Key List to utilize this feature. This means to manage this resources you must have the IP address or CIDR block that the Terraform connection is coming from added to the Atlas API Key Access List of the Atlas API key you are using. See [Resources that require API Key List](https://www.mongodb.com/docs/atlas/configure-api-access/#use-api-resources-that-require-an-access-list) for details.
`,
			Keywords: []string{},
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
					Name:        "delivery_type_config",
					Description: `(Required) Type of restore job to create. Possible configurations are:`,
				},
				resource.Attribute{
					Name:        "delivery_type_config.automated",
					Description: `Set to ` + "`" + `true` + "`" + ` to use the automated configuration.`,
				},
				resource.Attribute{
					Name:        "delivery_type_config.download",
					Description: `Set to ` + "`" + `true` + "`" + ` to use the download configuration.`,
				},
				resource.Attribute{
					Name:        "delivery_type_config.pointInTime",
					Description: `Set to ` + "`" + `true` + "`" + ` to use the pointInTime configuration. If using pointInTime configuration, you must also specify either ` + "`" + `oplog_ts` + "`" + ` and ` + "`" + `oplog_inc` + "`" + `, or ` + "`" + `point_in_time_utc_seconds` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delivery_type_config.target_cluster_name",
					Description: `Name of the target Atlas cluster to which the restore job restores the snapshot. Required for`,
				},
				resource.Attribute{
					Name:        "delivery_type_config.target_project_id",
					Description: `Name of the target Atlas cluster to which the restore job restores the snapshot. Required for`,
				},
				resource.Attribute{
					Name:        "delivery_type_config.oplog_ts",
					Description: `Optional setting for`,
				},
				resource.Attribute{
					Name:        "delivery_type_config.oplog_inc",
					Description: `Optional setting for`,
				},
				resource.Attribute{
					Name:        "delivery_type_config.point_in_time_utc_seconds",
					Description: `Optional setting for`,
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
					Name:        "delivery_type_config",
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
					Name:        "target_project_id",
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
					Name:        "delivery_type_config",
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
					Name:        "target_project_id",
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
			Type:             "mongodbatlas_cloud_provider_access",
			Category:         "Resources",
			ShortDescription: `Provides a Cloud Provider Access settings resource for registration, authorization, and deauthorization`,
			Description: ` Provider Access Configuration Paths

The Terraform MongoDB Atlas Provider offers two either-or/mutually exclusive paths to perform an authorization for a cloud provider role -

* A Single Resource path: using the ` + "`" + `mongodbatlas_cloud_provider_access` + "`" + ` that at provision time sets up all the required configuration for a given provider, then with a subsequent update it can perform the authorize of the role. Note this path requires two ` + "`" + `terraform apply` + "`" + ` commands, once for setup and once for auth.

* A Two Resource path: consisting of ` + "`" + `mongodbatlas_cloud_provider_access_setup` + "`" + ` and ` + "`" + `mongodbatlas_cloud_provider_access_authorization` + "`" + `. The first resource, ` + "`" + `mongodbatlas_cloud_provider_access_setup` + "`" + `, only generates
the initial configuration (create, delete operations). The second resource, ` + "`" + `mongodbatlas_cloud_provider_access_authorization` + "`" + `, helps to perform the authorization using the role_id of the first resource. This path is helpful in a multi-provider Terraform file, and allows for a single and decoupled apply. See example of this Two Resource path option with AWS Cloud [here](https://github.com/mongodb/terraform-provider-mongodbatlas/tree/master/examples/atlas-cloud-provider-access/aws). 

-> **IMPORTANT** If you want to move from the single resource path to the two resources path see the [migration guide](https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/guides/0.9.1-upgrade-guide#migration-to-cloud-provider-access-setup)

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) The cloud provider for which to create a new role. Currently only AWS is supported.`,
				},
				resource.Attribute{
					Name:        "iam_assumed_role_arn",
					Description: `(Optional) - ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access). ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used by terraform for internal management.`,
				},
				resource.Attribute{
					Name:        "atlas_assumed_role_external_id",
					Description: `Unique external ID Atlas uses when assuming the IAM role in your AWS account.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) The cloud provider for which to create a new role. Currently only AWS is supported. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used by terraform for internal management.`,
				},
				resource.Attribute{
					Name:        "aws",
					Description: `aws related arn roles`,
				},
				resource.Attribute{
					Name:        "atlas_assumed_role_external_id",
					Description: `Unique external ID Atlas uses when assuming the IAM role in your AWS account.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project`,
				},
				resource.Attribute{
					Name:        "iam_assumed_role_arn",
					Description: `(Required) ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access). ## Attributes Reference`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used by terraform for internal management.`,
				},
				resource.Attribute{
					Name:        "atlas_assumed_role_external_id",
					Description: `Unique external ID Atlas uses when assuming the IAM role in your AWS account.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) The cloud provider for which to create a new role. Currently only AWS is supported. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used by terraform for internal management.`,
				},
				resource.Attribute{
					Name:        "aws",
					Description: `aws related arn roles`,
				},
				resource.Attribute{
					Name:        "atlas_assumed_role_external_id",
					Description: `Unique external ID Atlas uses when assuming the IAM role in your AWS account.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project`,
				},
				resource.Attribute{
					Name:        "iam_assumed_role_arn",
					Description: `(Required) ARN of the IAM Role that Atlas assumes when accessing resources in your AWS account. This value is required after the creation (register of the role) as part of [Set Up Unified AWS Access](https://docs.atlas.mongodb.com/security/set-up-unified-aws-access/#set-up-unified-aws-access). ## Attributes Reference`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot",
			Category:         "Resources",
			ShortDescription: `Provides an Cloud Backup Snapshot resource.`,
			Description: `

` + "`" + `mongodbatlas_cloud_provider_snapshot` + "`" + ` provides a resource to take a cloud backup snapshot on demand.
On-demand snapshots happen immediately, unlike scheduled snapshots which occur at regular intervals. If there is already an on-demand snapshot with a status of queued or inProgress, you must wait until Atlas has completed the on-demand snapshot before taking another.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
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
					Description: `(Required) The number of days that Atlas should retain the on-demand snapshot. Must be at least 1.`,
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
			Description: `

` + "`" + `mongodbatlas_cloud_provider_snapshot_backup_policy` + "`" + ` provides a resource that enables you to view and modify the snapshot schedule and retention settings for an Atlas cluster with Cloud Backup enabled.  A default policy is created automatically when Cloud Backup is enabled for the cluster.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

# Examples - Modifying Polices
When Cloud Backup is enabled for a cluster MongoDB Atlas automatically creates a default Cloud Backup schedule for the cluster with four policy items; hourly, daily, weekly, and monthly. Because of this default creation this provider automatically saves the Cloud Backup Snapshot Policy into the Terraform state when a cluster is created/modified to use Cloud Backup. If the default works well for you then you do not need to do anything other than create a cluster with Cloud Backup enabled and your Terraform state will have this information if you need it. However, if you want the policy to be different than the default we've provided some examples to help below.

`,
			Keywords: []string{},
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
			Description: `

` + "`" + `mongodbatlas_cloud_provider_snapshot_restore_job` + "`" + ` provides a resource to create a new restore job from a cloud backup snapshot of a specified cluster. The restore job can be one of three types: 
* **automated:** Atlas automatically restores the snapshot with snapshotId to the Atlas cluster with name targetClusterName in the Atlas project with targetGroupId.

* **download:** Atlas provides a URL to download a .tar.gz of the snapshot with snapshotId. The contents of the archive contain the data files for your Atlas cluster.

* **pointInTime:**  Atlas performs a Continuous Cloud Backup restore.

-> **Important:** If you specify ` + "`" + `deliveryType` + "`" + ` : ` + "`" + `automated` + "`" + ` or ` + "`" + `deliveryType` + "`" + ` : ` + "`" + `pointInTime` + "`" + ` in your request body to create an automated restore job, Atlas removes all existing data on the target cluster prior to the restore.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
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
					Name:        "target_project_id",
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
					Name:        "delivery_type_config",
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
					Name:        "target_project_id",
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
					Name:        "delivery_type_config",
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
					Name:        "target_project_id",
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
			Description: `

` + "`" + `mongodbatlas_cluster` + "`" + ` provides a Cluster resource. The resource lets you create, edit and delete clusters. The resource requires your Project ID.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

-> **NOTE:** A network container is created for a cluster to reside in. To use this container with another resource, such as peering, reference the computed` + "`" + `container_id` + "`" + ` attribute on the cluster.

-> **NOTE:** If Backup Compliance Policy is enabled for the project for which this backup schedule is defined, you cannot modify the backup schedule for an individual cluster below the minimum requirements set in the Backup Compliance Policy.  See [Backup Compliance Policy Prohibited Actions and Considerations](https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/#configure-a-backup-compliance-policy).


~> **IMPORTANT:**
<br> &#8226; New Users: If you are not already using ` + "`" + `mongodbatlas_cluster` + "`" + ` for your deployment we recommend starting with the [` + "`" + `mongodbatlas_advanced_cluster` + "`" + `](https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/resources/advanced_cluster).  ` + "`" + `mongodbatlas_advanced_cluster` + "`" + ` has all the same functionality as ` + "`" + `mongodbatlas_cluster` + "`" + ` but also supports multi-cloud clusters.  
<br> &#8226; Free tier cluster creation (M0) is supported.
<br> &#8226; Shared tier clusters (M0, M2, M5) can be upgraded to dedicated tiers (M10+) via this provider. WARNING WHEN UPGRADING TENANT/SHARED CLUSTERS!!! Any change from shared tier to a different instance size will be considered a tenant upgrade. When upgrading from shared tier to dedicated simply change the ` + "`" + `provider_name` + "`" + ` from "TENANT"  to your preferred provider (AWS, GCP, AZURE) and remove the variable ` + "`" + `backing_provider_name` + "`" + `, for example if you have an existing tenant/shared cluster and want to upgrade your Terraform config should be changed from:
` + "`" + `` + "`" + `` + "`" + `
provider_instance_size_name = "M0"
provider_name               = "TENANT"
backing_provider_name       = "AWS"
` + "`" + `` + "`" + `` + "`" + `
To:
` + "`" + `` + "`" + `` + "`" + `
provider_instance_size_name = "M10"
provider_name               = "AWS"
` + "`" + `` + "`" + `` + "`" + `
<br> &#8226; Changes to cluster configurations can affect costs. Before making changes, please see [Billing](https://docs.atlas.mongodb.com/billing/).   
<br> &#8226; If your Atlas project contains a custom role that uses actions introduced in a specific MongoDB version, you cannot create a cluster with a MongoDB version less than that version unless you delete the custom role.

`,
			Keywords: []string{},
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
					Description: `(Optional) Specifies whether disk auto-scaling is enabled. The default is true. - Set to ` + "`" + `true` + "`" + ` to enable disk auto-scaling. - Set to ` + "`" + `false` + "`" + ` to disable disk auto-scaling. ~>`,
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
					Name:        "bi_connector_config",
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
					Description: `(Optional) Version of the cluster to deploy. Atlas supports the following MongoDB versions for M10+ clusters: ` + "`" + `4.2` + "`" + `, ` + "`" + `4.4` + "`" + `, ` + "`" + `5.0` + "`" + `, or ` + "`" + `6.0` + "`" + `. If omitted, Atlas deploys a cluster that runs MongoDB 5.0. If ` + "`" + `provider_instance_size_name` + "`" + `: ` + "`" + `M0` + "`" + `, ` + "`" + `M2` + "`" + ` or ` + "`" + `M5` + "`" + `, Atlas deploys MongoDB 5.0. Atlas always deploys the cluster with the latest stable release of the specified version. See [Release Notes](https://www.mongodb.com/docs/upcoming/release-notes/) for latest Current Stable Release.`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `(Optional) Selects whether the cluster is a replica set or a sharded cluster. If you use the replicationSpecs parameter, you must set num_shards.`,
				},
				resource.Attribute{
					Name:        "pit_enabled",
					Description: `(Optional) - Flag that indicates if the cluster uses Continuous Cloud Backup. If set to true, cloud_backup must also be set to true.`,
				},
				resource.Attribute{
					Name:        "provider_backup_enabled",
					Description: `(Optional) Flag indicating if the cluster uses Cloud Backup for backups.`,
				},
				resource.Attribute{
					Name:        "cloud_backup",
					Description: `(Optional) Flag indicating if the cluster uses Cloud Backup for backups. If true, the cluster uses Cloud Backup for backups. If cloud_backup and backup_enabled are false, the cluster does not use Atlas backups. You cannot enable cloud backup if you have an existing cluster in the project with legacy backup enabled. ~>`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `(Optional) Cloud service provider on which the server for a multi-tenant cluster is provisioned. This setting is only valid when providerSetting.providerName is TENANT and providerSetting.instanceSizeName is M2 or M5. The possible values are: - AWS - Amazon AWS - GCP - Google Cloud Platform - AZURE - Microsoft Azure`,
				},
				resource.Attribute{
					Name:        "provider_disk_iops",
					Description: `(Optional - AWS Only) The maximum input/output operations per second (IOPS) the system can perform. The possible values depend on the selected ` + "`" + `provider_instance_size_name` + "`" + ` and ` + "`" + `disk_size_gb` + "`" + `. This setting requires that ` + "`" + `provider_instance_size_name` + "`" + ` to be M30 or greater and cannot be used with clusters with local NVMe SSDs. The default value for ` + "`" + `provider_disk_iops` + "`" + ` is the same as the cluster tier's Standard IOPS value, as viewable in the Atlas console. It is used in cases where a higher number of IOPS is needed and possible. If a value is submitted that is lower or equal to the default IOPS value for the cluster tier Atlas ignores the requested value and uses the default. More details available under the providerSettings.diskIOPS parameter: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
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
					Description: `(AWS - Optional) The type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `. ` + "`" + `PROVISIONED` + "`" + ` is ONLY required if setting IOPS higher than the default instance IOPS.`,
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
					Description: `Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details.`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.`,
				},
				resource.Attribute{
					Name:        "version_release_system",
					Description: `(Optional) - Release cadence that Atlas uses for this cluster. This parameter defaults to ` + "`" + `LTS` + "`" + `. If you set this field to ` + "`" + `CONTINUOUS` + "`" + `, you must omit the ` + "`" + `mongo_db_major_version` + "`" + ` field. Atlas accepts: - ` + "`" + `CONTINUOUS` + "`" + `: Atlas creates your cluster using the most recent MongoDB release. Atlas automatically updates your cluster to the latest major and rapid MongoDB releases as they become available. - ` + "`" + `LTS` + "`" + `: Atlas creates your cluster using the latest patch release of the MongoDB version that you specify in the mongoDBMajorVersion field. Atlas automatically updates your cluster to subsequent patch releases of this MongoDB version. Atlas doesn't update your cluster to newer rapid or major MongoDB releases as they become available.`,
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
					Description: `(Optional) The number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. If you do not specify this option, no analytics nodes are deployed to the region. ### BI Connector Specifies BI Connector for Atlas configuration. ` + "`" + `` + "`" + `` + "`" + `terraform bi_connector_config { enabled = true read_preference = "secondary" } ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "default_read_concern",
					Description: `(Optional) [Default level of acknowledgment requested from MongoDB for read operations](https://docs.mongodb.com/manual/reference/read-concern/) set for this cluster. MongoDB 4.4 clusters default to [available](https://docs.mongodb.com/manual/reference/read-concern-available/).`,
				},
				resource.Attribute{
					Name:        "default_write_concern",
					Description: `(Optional) [Default level of acknowledgment requested from MongoDB for write operations](https://docs.mongodb.com/manual/reference/write-concern/) set for this cluster. MongoDB 4.4 clusters default to [1](https://docs.mongodb.com/manual/reference/write-concern/).`,
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
					Name:        "oplog_min_retention_hours",
					Description: `(Optional) Minimum retention window for cluster's oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates.`,
				},
				resource.Attribute{
					Name:        "sample_size_bi_connector",
					Description: `(Optional) Number of documents per database to sample when gathering schema information. Defaults to 100. Available only for Atlas deployments in which BI Connector for Atlas is enabled.`,
				},
				resource.Attribute{
					Name:        "sample_refresh_interval_bi_connector",
					Description: `(Optional) Interval in seconds at which the mongosqld process re-samples data to create its relational schema. The default value is 300. The specified value must be a positive integer. Available only for Atlas deployments in which BI Connector for Atlas is enabled. ### Labels ` + "`" + `` + "`" + `` + "`" + `terraform labels { key = "Key 1" value = "Value 1" } labels { key = "Key 2" value = "Value 2" } ` + "`" + `` + "`" + `` + "`" + ` Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. You cannot set the key ` + "`" + `Infrastructure Tool` + "`" + `, it is used for internal purposes to track aggregate usage.`,
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
					Description: `The Container ID is the id of the container created when the first cluster in the region (AWS/Azure) or project (GCP) was created.`,
				},
				resource.Attribute{
					Name:        "srv_address",
					Description: `Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS/SSL. See the mongoURI for additional options.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING ### Cloud Backup Policy`,
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
					Description: `The Container ID is the id of the container created when the first cluster in the region (AWS/Azure) or project (GCP) was created.`,
				},
				resource.Attribute{
					Name:        "srv_address",
					Description: `Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS/SSL. See the mongoURI for additional options.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Current state of the cluster. The possible states are: - IDLE - CREATING - UPDATING - DELETING - DELETED - REPAIRING ### Cloud Backup Policy`,
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
			Description: `

` + "`" + `mongodbatlas_custom_db_role` + "`" + ` provides a Custom DB Role resource. The customDBRoles resource lets you retrieve, create and modify the custom MongoDB roles in your cluster. Use custom MongoDB roles to specify custom sets of actions which cannot be described by the built-in Atlas database user privileges.

-> **IMPORTANT**  You define custom roles at the project level for all clusters in the project. The ` + "`" + `mongodbatlas_custom_db_role` + "`" + ` resource supports a subset of MongoDB privilege actions. For a complete list of [privilege actions](https://docs.mongodb.com/manual/reference/privilege-actions/) available for this resource, see [Custom Role actions](https://docs.atlas.mongodb.com/reference/api/custom-role-actions/). Custom roles must include actions that all project's clusters support, and that are compatible with each MongoDB version used by your project's clusters. For example, if your project has MongoDB 4.2 clusters, you can't create custom roles that use actions introduced in MongoDB 4.4.


-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

`,
			Keywords: []string{},
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
					Name:        "resources.#.collection_name",
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
			Type:             "mongodbatlas_custom_dns_configuration_cluster_aws",
			Category:         "Resources",
			ShortDescription: `Provides a Custom DNS Configuration for Atlas Clusters on AWS resource.`,
			Description: `

` + "`" + `mongodbatlas_custom_dns_configuration_cluster_aws` + "`" + ` provides a Custom DNS Configuration for Atlas Clusters on AWS resource. This represents a Custom DNS Configuration for Atlas Clusters on AWS that can be updated in an Atlas project.

~> **IMPORTANT:**You must have one of the following roles to successfully handle the resource:
  * Organization Owner
  * Project Owner

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `Required Unique identifier for the project.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Indicates whether the project's clusters deployed to AWS use custom DNS. If ` + "`" + `true` + "`" + `, the ` + "`" + `Get All Clusters` + "`" + ` and ` + "`" + `Get One Cluster` + "`" + ` endpoints return the ` + "`" + `connectionStrings.private` + "`" + ` and ` + "`" + `connectionStrings.privateSrv` + "`" + ` fields for clusters deployed to AWS . ## Import Custom DNS Configuration for Atlas Clusters on AWS must be imported using Project ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_custom_dns_configuration_cluster_aws.test 1112222b3bf99403840e8934 ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Custom DNS Configuration for Atlas Clusters on AWS](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Custom-DNS-for-Atlas-Clusters-Deployed-to-AWS)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_data_lake",
			Category:         "Resources",
			ShortDescription: `Provides a Data Lake resource.`,
			Description: `

` + "`" + `mongodbatlas_data_lake` + "`" + ` provides a Data Lake resource.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

~> **IMPORTANT:** All arguments including the password will be stored in the raw state as plain-text. [Read more about sensitive data in state.](https://www.terraform.io/docs/state/sensitive-data.html)

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create a data lake.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Atlas Data Lake.`,
				},
				resource.Attribute{
					Name:        "aws",
					Description: `(Required) AWS provider of the cloud service where Data Lake can access the S3 Bucket.`,
				},
				resource.Attribute{
					Name:        "aws.0.role_id",
					Description: `(Required) Unique identifier of the role that Data Lake can use to access the data stores. If necessary, use the Atlas [UI](https://docs.atlas.mongodb.com/security/manage-iam-roles/) or [API](https://docs.atlas.mongodb.com/reference/api/cloud-provider-access-get-roles/) to retrieve the role ID. You must also specify the ` + "`" + `aws.0.test_s3_bucket` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aws.0.test_s3_bucket",
					Description: `(Required) Name of the S3 data bucket that the provided role ID is authorized to access. You must also specify the ` + "`" + `aws.0.role_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "data_process_region",
					Description: `(Optional) The cloud provider region to which Atlas Data Lake routes client connections for data processing. Set to ` + "`" + `null` + "`" + ` to direct Atlas Data Lake to route client connections to the region nearest to the client based on DNS resolution.`,
				},
				resource.Attribute{
					Name:        "data_process_region.0.cloud_provider",
					Description: `(Required) Name of the cloud service provider. Atlas Data Lake only supports AWS.`,
				},
				resource.Attribute{
					Name:        "data_process_region.0.region",
					Description: `(Required). Name of the region to which Data Lake routes client connections for data processing. Atlas Data Lake only supports the following regions:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "aws.0.iam_assumed_role_arn",
					Description: `Amazon Resource Name (ARN) of the IAM Role that Data Lake assumes when accessing S3 Bucket data stores. The IAM Role must support the following actions against each S3 bucket:`,
				},
				resource.Attribute{
					Name:        "aws.0.iam_user_arn",
					Description: `Amazon Resource Name (ARN) of the user that Data Lake assumes when accessing S3 Bucket data stores.`,
				},
				resource.Attribute{
					Name:        "aws.0.external_id",
					Description: `Unique identifier associated with the IAM Role that Data Lake assumes when accessing the data stores.`,
				},
				resource.Attribute{
					Name:        "hostnames",
					Description: `The list of hostnames assigned to the Atlas Data Lake. Each string in the array is a hostname assigned to the Atlas Data Lake.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Atlas Data Lake:`,
				},
				resource.Attribute{
					Name:        "ACTIVE",
					Description: `The Data Lake is active and verified. You can query the data stores associated with the Atlas Data Lake.`,
				},
				resource.Attribute{
					Name:        "storage_databases",
					Description: `Configuration details for mapping each data store to queryable databases and collections. For complete documentation on this object and its nested fields, see [databases](https://docs.mongodb.com/datalake/reference/format/data-lake-configuration#std-label-datalake-databases-reference). An empty object indicates that the Data Lake has no mapping configuration for any data store.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.name",
					Description: `Name of the database to which Data Lake maps the data contained in the data store.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.collections",
					Description: `Array of objects where each object represents a collection and data sources that map to a [stores](https://docs.mongodb.com/datalake/reference/format/data-lake-configuration#mongodb-datalakeconf-datalakeconf.stores) data store.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.collections.#.name",
					Description: `Name of the collection.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.collections.#.data_sources",
					Description: `Array of objects where each object represents a stores data store to map with the collection.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.collections.#.data_sources.#.store_name",
					Description: `Name of a data store to map to the ` + "`" + `<collection>` + "`" + `. Must match the name of an object in the stores array.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.collections.#.data_sources.#.default_format",
					Description: `Default format that Data Lake assumes if it encounters a file without an extension while searching the storeName.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.collections.#.data_sources.#.path",
					Description: `Controls how Atlas Data Lake searches for and parses files in the storeName before mapping them to the ` + "`" + `<collection>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.views",
					Description: `Array of objects where each object represents an [aggregation pipeline](https://docs.mongodb.com/manual/core/aggregation-pipeline/#id1) on a collection. To learn more about views, see [Views](https://docs.mongodb.com/manual/core/views/).`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.views.#.name",
					Description: `Name of the view.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.views.#.source",
					Description: `Name of the source collection for the view.`,
				},
				resource.Attribute{
					Name:        "storage_stores",
					Description: `Each object in the array represents a data store. Data Lake uses the storage.databases configuration details to map data in each data store to queryable databases and collections. For complete documentation on this object and its nested fields, see [stores](https://docs.mongodb.com/datalake/reference/format/data-lake-configuration#std-label-datalake-stores-reference). An empty object indicates that the Data Lake has no configured data stores.`,
				},
				resource.Attribute{
					Name:        "storage_stores.#.name",
					Description: `Name of the data store.`,
				},
				resource.Attribute{
					Name:        "storage_stores.#.provider",
					Description: `Defines where the data is stored.`,
				},
				resource.Attribute{
					Name:        "storage_stores.#.region",
					Description: `Name of the AWS region in which the S3 bucket is hosted.`,
				},
				resource.Attribute{
					Name:        "storage_stores.#.bucket",
					Description: `Name of the AWS S3 bucket.`,
				},
				resource.Attribute{
					Name:        "storage_stores.#.prefix",
					Description: `Prefix Data Lake applies when searching for files in the S3 bucket .`,
				},
				resource.Attribute{
					Name:        "storage_stores.#.delimiter",
					Description: `The delimiter that separates ` + "`" + `storage_databases.#.collections.#.data_sources.#.path` + "`" + ` segments in the data store.`,
				},
				resource.Attribute{
					Name:        "storage_stores.#.include_tags",
					Description: `Determines whether or not to use S3 tags on the files in the given path as additional partition attributes. ## Import Data Lake can be imported using project ID, name of the data lake and name of the AWS s3 bucket, in the format ` + "`" + `project_id` + "`" + `--` + "`" + `name` + "`" + `--` + "`" + `aws_test_s3_bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_data_lake.example 1112222b3bf99403840e8934--test-data-lake--s3-test ` + "`" + `` + "`" + `` + "`" + ` See [MongoDB Atlas API](https://docs.mongodb.com/datalake/reference/api/dataLakes-create-one-tenant) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "aws.0.iam_assumed_role_arn",
					Description: `Amazon Resource Name (ARN) of the IAM Role that Data Lake assumes when accessing S3 Bucket data stores. The IAM Role must support the following actions against each S3 bucket:`,
				},
				resource.Attribute{
					Name:        "aws.0.iam_user_arn",
					Description: `Amazon Resource Name (ARN) of the user that Data Lake assumes when accessing S3 Bucket data stores.`,
				},
				resource.Attribute{
					Name:        "aws.0.external_id",
					Description: `Unique identifier associated with the IAM Role that Data Lake assumes when accessing the data stores.`,
				},
				resource.Attribute{
					Name:        "hostnames",
					Description: `The list of hostnames assigned to the Atlas Data Lake. Each string in the array is a hostname assigned to the Atlas Data Lake.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Atlas Data Lake:`,
				},
				resource.Attribute{
					Name:        "ACTIVE",
					Description: `The Data Lake is active and verified. You can query the data stores associated with the Atlas Data Lake.`,
				},
				resource.Attribute{
					Name:        "storage_databases",
					Description: `Configuration details for mapping each data store to queryable databases and collections. For complete documentation on this object and its nested fields, see [databases](https://docs.mongodb.com/datalake/reference/format/data-lake-configuration#std-label-datalake-databases-reference). An empty object indicates that the Data Lake has no mapping configuration for any data store.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.name",
					Description: `Name of the database to which Data Lake maps the data contained in the data store.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.collections",
					Description: `Array of objects where each object represents a collection and data sources that map to a [stores](https://docs.mongodb.com/datalake/reference/format/data-lake-configuration#mongodb-datalakeconf-datalakeconf.stores) data store.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.collections.#.name",
					Description: `Name of the collection.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.collections.#.data_sources",
					Description: `Array of objects where each object represents a stores data store to map with the collection.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.collections.#.data_sources.#.store_name",
					Description: `Name of a data store to map to the ` + "`" + `<collection>` + "`" + `. Must match the name of an object in the stores array.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.collections.#.data_sources.#.default_format",
					Description: `Default format that Data Lake assumes if it encounters a file without an extension while searching the storeName.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.collections.#.data_sources.#.path",
					Description: `Controls how Atlas Data Lake searches for and parses files in the storeName before mapping them to the ` + "`" + `<collection>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.views",
					Description: `Array of objects where each object represents an [aggregation pipeline](https://docs.mongodb.com/manual/core/aggregation-pipeline/#id1) on a collection. To learn more about views, see [Views](https://docs.mongodb.com/manual/core/views/).`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.views.#.name",
					Description: `Name of the view.`,
				},
				resource.Attribute{
					Name:        "storage_databases.#.views.#.source",
					Description: `Name of the source collection for the view.`,
				},
				resource.Attribute{
					Name:        "storage_stores",
					Description: `Each object in the array represents a data store. Data Lake uses the storage.databases configuration details to map data in each data store to queryable databases and collections. For complete documentation on this object and its nested fields, see [stores](https://docs.mongodb.com/datalake/reference/format/data-lake-configuration#std-label-datalake-stores-reference). An empty object indicates that the Data Lake has no configured data stores.`,
				},
				resource.Attribute{
					Name:        "storage_stores.#.name",
					Description: `Name of the data store.`,
				},
				resource.Attribute{
					Name:        "storage_stores.#.provider",
					Description: `Defines where the data is stored.`,
				},
				resource.Attribute{
					Name:        "storage_stores.#.region",
					Description: `Name of the AWS region in which the S3 bucket is hosted.`,
				},
				resource.Attribute{
					Name:        "storage_stores.#.bucket",
					Description: `Name of the AWS S3 bucket.`,
				},
				resource.Attribute{
					Name:        "storage_stores.#.prefix",
					Description: `Prefix Data Lake applies when searching for files in the S3 bucket .`,
				},
				resource.Attribute{
					Name:        "storage_stores.#.delimiter",
					Description: `The delimiter that separates ` + "`" + `storage_databases.#.collections.#.data_sources.#.path` + "`" + ` segments in the data store.`,
				},
				resource.Attribute{
					Name:        "storage_stores.#.include_tags",
					Description: `Determines whether or not to use S3 tags on the files in the given path as additional partition attributes. ## Import Data Lake can be imported using project ID, name of the data lake and name of the AWS s3 bucket, in the format ` + "`" + `project_id` + "`" + `--` + "`" + `name` + "`" + `--` + "`" + `aws_test_s3_bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_data_lake.example 1112222b3bf99403840e8934--test-data-lake--s3-test ` + "`" + `` + "`" + `` + "`" + ` See [MongoDB Atlas API](https://docs.mongodb.com/datalake/reference/api/dataLakes-create-one-tenant) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_database_user",
			Category:         "Resources",
			ShortDescription: `Provides a Database User resource.`,
			Description: `

` + "`" + `mongodbatlas_database_user` + "`" + ` provides a Database User resource. This represents a database user which will be applied to all clusters within the project.

Each user has a set of roles that provide access to the project’s databases. User's roles apply to all the clusters in the project: if two clusters have a ` + "`" + `products` + "`" + ` database and a user has a role granting ` + "`" + `read` + "`" + ` access on the products database, the user has that access on both clusters.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

~> **IMPORTANT:** All arguments including the password will be stored in the raw state as plain-text. [Read more about sensitive data in state.](https://www.terraform.io/docs/state/sensitive-data.html)

`,
			Keywords: []string{},
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
					Description: `(Required) Username for authenticating to MongoDB. USER_ARN or ROLE_ARN if ` + "`" + `aws_iam_type` + "`" + ` is USER or ROLE.`,
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
			Description: `

` + "`" + `mongodbatlas_encryption_at_rest` + "`" + ` Allows management of encryption at rest for an Atlas project with one of the following providers:

[Amazon Web Services Key Management Service](https://docs.atlas.mongodb.com/security-aws-kms/#security-aws-kms)
[Azure Key Vault](https://docs.atlas.mongodb.com/security-azure-kms/#security-azure-kms)
[Google Cloud KMS](https://docs.atlas.mongodb.com/security-gcp-kms/#security-gcp-kms)

After configuring at least one Encryption at Rest provider for the Atlas project, Project Owners can enable Encryption at Rest for each Atlas cluster for which they require encryption. The Encryption at Rest provider does not have to match the cluster cloud service provider.

Atlas does not automatically rotate user-managed encryption keys. Defer to your preferred Encryption at Rest provider’s documentation and guidance for best practices on key rotation. Atlas automatically creates a 90-day key rotation alert when you configure Encryption at Rest using your Key Management in an Atlas project.

See [Encryption at Rest](https://docs.atlas.mongodb.com/security-kms-encryption/index.html) for more information, including prerequisites and restrictions.

~> **IMPORTANT** Atlas encrypts all cluster storage and snapshot volumes, securing all cluster data on disk: a concept known as encryption at rest, by default.

~> **IMPORTANT** Atlas limits this feature to dedicated cluster tiers of M10 and greater. For more information see: https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Encryption-at-Rest-using-Customer-Key-Management

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.


-> **IMPORTANT NOTE** To disable the encryption at rest with customer key management for a project all existing clusters in the project must first either have encryption at rest for the provider set to none, e.g. ` + "`" + `encryption_at_rest_provider = "NONE"` + "`" + `, or be deleted.

`,
			Keywords: []string{},
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
					Description: `(Required) Specifies GCP KMS configuration details and whether Encryption at Rest is enabled for an Atlas project. ### aws_kms_config Refer to the example in the [official github repository](https://github.com/mongodb/terraform-provider-mongodbatlas/tree/master/examples) to implement Encryption at Rest`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Specifies whether Encryption at Rest is enabled for an Atlas project, To disable Encryption at Rest, pass only this parameter with a value of false, When you disable Encryption at Rest, Atlas also removes the configuration details.`,
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
					Description: `ID of an AWS IAM role authorized to manage an AWS customer master key. To find the ID for an existing IAM role check the ` + "`" + `role_id` + "`" + ` attribute of the ` + "`" + `mongodbatlas_cloud_provider_access` + "`" + ` resource. ### azure_key_vault_config`,
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
					Description: `The unique identifier for an Azure AD tenant within an Azure subscription. ### google_cloud_kms_config`,
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
					Description: `The Key Version Resource ID from your GCP account. ## Import Encryption at Rest Settings can be imported using project ID, in the format ` + "`" + `project_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_encryption_at_rest.example 1112222b3bf99403840e8934 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference for Encryption at Rest using Customer Key Management.](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Encryption-at-Rest-using-Customer-Key-Management)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_event_trigger",
			Category:         "Resources",
			ShortDescription: `Provides a Event Trigger resource.`,
			Description: `

` + "`" + `mongodbatlas_event_trigger` + "`" + ` provides a Event Trigger resource. 

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the trigger.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) The ObjectID of your application.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the trigger.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the trigger. Possible Values: ` + "`" + `DATABASE` + "`" + `, ` + "`" + `AUTHENTICATION` + "`" + `,` + "`" + `SCHEDULED` + "`" + ``,
				},
				resource.Attribute{
					Name:        "function_id",
					Description: `(Optional) The ID of the function associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Default: ` + "`" + `false` + "`" + ` If ` + "`" + `true` + "`" + `, the trigger is disabled.`,
				},
				resource.Attribute{
					Name:        "config_operation_types",
					Description: `(Optional) Required for ` + "`" + `DATABASE` + "`" + ` type. The [database event operation types](https://docs.mongodb.com/realm/triggers/database-triggers/#std-label-database-events) to listen for. This must contain at least one value. Possible Values: ` + "`" + `INSERT` + "`" + `, ` + "`" + `UPDATE` + "`" + `, ` + "`" + `REPLACE` + "`" + `, ` + "`" + `DELETE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "config_operation_type",
					Description: `(Optional) Required for ` + "`" + `AUTHENTICATION` + "`" + ` type. The [authentication operation type](https://docs.mongodb.com/realm/triggers/authentication-triggers/#std-label-authentication-event-operation-types) to listen for. Possible Values: ` + "`" + `LOGIN` + "`" + `, ` + "`" + `CREATE` + "`" + `, ` + "`" + `DELETE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "config_providers",
					Description: `(Optional) Required for ` + "`" + `AUTHENTICATION` + "`" + ` type. A list of one or more [authentication provider](https://docs.mongodb.com/realm/authentication/providers/) id values. The trigger will only listen for authentication events produced by these providers.`,
				},
				resource.Attribute{
					Name:        "config_database",
					Description: `(Optional) Required for ` + "`" + `DATABASE` + "`" + ` type. The name of the MongoDB database that contains the watched collection.`,
				},
				resource.Attribute{
					Name:        "config_collection",
					Description: `(Optional) Required for ` + "`" + `DATABASE` + "`" + ` type. The name of the MongoDB collection that the trigger watches for change events. The collection must be part of the specified database.`,
				},
				resource.Attribute{
					Name:        "config_service_id",
					Description: `(Optional) Required for ` + "`" + `DATABASE` + "`" + ` type. The ID of the MongoDB Service associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "config_match",
					Description: `(Optional) Optional for ` + "`" + `DATABASE` + "`" + ` type. A [$match](https://docs.mongodb.com/manual/reference/operator/aggregation/match/) expression document that MongoDB Realm includes in the underlying change stream pipeline for the trigger. This is useful when you want to filter change events beyond their operation type. The trigger will only fire if the expression evaluates to true for a given change event.`,
				},
				resource.Attribute{
					Name:        "config_project",
					Description: `(Optional) Optional for ` + "`" + `DATABASE` + "`" + ` type. A [$project](https://docs.mongodb.com/manual/reference/operator/aggregation/project/) expression document that Realm uses to filter the fields that appear in change event objects.`,
				},
				resource.Attribute{
					Name:        "config_full_document",
					Description: `(Optional) Optional for ` + "`" + `DATABASE` + "`" + ` type. If true, indicates that ` + "`" + `UPDATE` + "`" + ` change events should include the most current [majority-committed](https://docs.mongodb.com/manual/reference/read-concern-majority/) version of the modified document in the fullDocument field.`,
				},
				resource.Attribute{
					Name:        "unordered",
					Description: `Only Available for Database Triggers. If true, event ordering is disabled and this trigger can process events in parallel. If false, event ordering is enabled and the trigger executes serially.`,
				},
				resource.Attribute{
					Name:        "config_schedule",
					Description: `(Optional) Required for ` + "`" + `SCHEDULED` + "`" + ` type. A [cron expression](https://docs.mongodb.com/realm/triggers/cron-expressions/) that defines the trigger schedule.`,
				},
				resource.Attribute{
					Name:        "event_processors",
					Description: `(Optional) An object where each field name is an event processor ID and each value is an object that configures its corresponding event processor. The following event processors are supported: ` + "`" + `AWS_EVENTBRIDGE` + "`" + ` For an example configuration object, see [Send Trigger Events to AWS EventBridge](https://docs.mongodb.com/realm/triggers/eventbridge/#std-label-event_processor_example).`,
				},
				resource.Attribute{
					Name:        "event_processors.0.aws_eventbridge.config_account_id",
					Description: `(Optional) AWS Account ID.`,
				},
				resource.Attribute{
					Name:        "event_processors.0.aws_eventbridge.config_region",
					Description: `(Optional) Region of AWS Account. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "trigger_id",
					Description: `The unique ID of the trigger.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `The name of the function associated with the trigger. ## Import Event trigger can be imported using project ID, App ID and Trigger ID, in the format ` + "`" + `project_id` + "`" + `--` + "`" + `app_id` + "`" + `-` + "`" + `trigger_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_event_trigger.test 1112222b3bf99403840e8934--testing-example--1112222b3bf99403840e8934 ` + "`" + `` + "`" + `` + "`" + ` For more details on this resource see [Triggers resource](https://www.mongodb.com/docs/atlas/app-services/admin/api/v3/#tag/triggers) in Atlas App Services Documentation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "trigger_id",
					Description: `The unique ID of the trigger.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `The name of the function associated with the trigger. ## Import Event trigger can be imported using project ID, App ID and Trigger ID, in the format ` + "`" + `project_id` + "`" + `--` + "`" + `app_id` + "`" + `-` + "`" + `trigger_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_event_trigger.test 1112222b3bf99403840e8934--testing-example--1112222b3bf99403840e8934 ` + "`" + `` + "`" + `` + "`" + ` For more details on this resource see [Triggers resource](https://www.mongodb.com/docs/atlas/app-services/admin/api/v3/#tag/triggers) in Atlas App Services Documentation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_federated_settings_identity_provider",
			Category:         "Resources",
			ShortDescription: `Provides a federated settings Identity Provider resource.`,
			Description: `

` + "`" + `mongodbatlas_federated_settings_identity_provider` + "`" + ` provides an Atlas federated settings identity provider resource provides a subset of settings to be maintained post import of the existing resource.
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Human-readable label that identifies the identity provider.`,
				},
				resource.Attribute{
					Name:        "associated_domains",
					Description: `(Required) List that contains the domains associated with the identity provider.`,
				},
				resource.Attribute{
					Name:        "sso_debug_enabled",
					Description: `(Required) Flag that indicates whether the identity provider has SSO debug enabled.`,
				},
				resource.Attribute{
					Name:        "issuer_uri",
					Description: `(Required) Unique string that identifies the issuer of the SAML`,
				},
				resource.Attribute{
					Name:        "sso_url",
					Description: `(Required) Unique string that identifies the intended audience of the SAML assertion.`,
				},
				resource.Attribute{
					Name:        "request_binding",
					Description: `(Required) SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request. Atlas supports the following binding values: - HTTP POST - HTTP REDIRECT`,
				},
				resource.Attribute{
					Name:        "response_signature_algorithm",
					Description: `(Required) Signature algorithm that Federated Authentication uses to encrypt the identity provider signature. Valid values include SHA-1 and SHA-256. ## Attributes Reference In addition to all arguments above, the following attributes are exported: ### FederatedSettingsIdentityProvider`,
				},
				resource.Attribute{
					Name:        "okta_idp_id",
					Description: `Unique 20-hexadecimal digit string that identifies the IdP. ## Import Identity Provider`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "okta_idp_id",
					Description: `Unique 20-hexadecimal digit string that identifies the IdP. ## Import Identity Provider`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_federated_settings_org_config",
			Category:         "Resources",
			ShortDescription: `Provides a federated settings Organization Configuration.`,
			Description: `

` + "`" + `mongodbatlas_federated_settings_org_config` + "`" + ` provides an Federated Settings Identity Providers datasource. Atlas Cloud Federated Settings Identity Providers provides federated settings outputs for the configured Identity Providers.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the organization that contains your projects.`,
				},
				resource.Attribute{
					Name:        "domain_allow_list",
					Description: `List that contains the approved domains from which organization users can log in.`,
				},
				resource.Attribute{
					Name:        "post_auth_role_grants",
					Description: `(Optional) List that contains the default [roles](https://www.mongodb.com/docs/atlas/reference/user-roles/#std-label-organization-roles) granted to users who authenticate through the IdP in a connected organization.`,
				},
				resource.Attribute{
					Name:        "domain_restriction_enabled",
					Description: `(Required) Flag that indicates whether domain restriction is enabled for the connected organization.`,
				},
				resource.Attribute{
					Name:        "identity_provider_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the federated authentication configuration. ## Attributes Reference In addition to all arguments above, the following attributes are exported: ## Import FederatedSettingsOrgConfig must be imported using federation_settings_id-org_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_federated_settings_org_config.org_connection 627a9687f7f7f7f774de306f14-627a9683ea7ff7f74de306f14 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_federated_settings_org_role_mapping",
			Category:         "Resources",
			ShortDescription: `Provides a federated settings Role Mapping resource.`,
			Description: `

` + "`" + `mongodbatlas_federated_settings_org_role_mapping` + "`" + ` provides an Role Mapping resource. This allows organization role mapping to be created.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects.`,
				},
				resource.Attribute{
					Name:        "external_group_name",
					Description: `Unique human-readable label that identifies the identity provider group to which this role mapping applies.`,
				},
				resource.Attribute{
					Name:        "role_assignments",
					Description: `Atlas roles and the unique identifiers of the groups and organizations associated with each role.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Unique identifier of the project to which you want the role mapping to apply.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Specifies the Roles that are attached to the Role Mapping. Available role IDs can be found on [the User Roles Reference](https://www.mongodb.com/docs/atlas/reference/user-roles/). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this role mapping. ## Import FederatedSettingsOrgRoleMapping can be imported using federation_settings_id-org_id-role_mapping_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_federated_settings_org_role_mapping.org_group_role_mapping_import 6287a663c7f7f7f71c441c6c-627a96837f7f7f7e306f14-628ae97f7f7468ea3727 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this role mapping. ## Import FederatedSettingsOrgRoleMapping can be imported using federation_settings_id-org_id-role_mapping_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_federated_settings_org_role_mapping.org_group_role_mapping_import 6287a663c7f7f7f71c441c6c-627a96837f7f7f7e306f14-628ae97f7f7468ea3727 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_global_cluster_config",
			Category:         "Resources",
			ShortDescription: `Provides a Global Cluster Configuration resource.`,
			Description: `

` + "`" + `mongodbatlas_global_cluster_config` + "`" + ` provides a Global Cluster Configuration resource.


-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Global Cluster.`,
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
					Description: `(Required) The name of the database containing the collection.`,
				},
				resource.Attribute{
					Name:        "is_custom_shard_key_hashed",
					Description: `(Optional) Specifies whether the custom shard key for the collection is [hashed](https://docs.mongodb.com/manual/reference/method/sh.shardCollection/#hashed-shard-keys). If omitted, defaults to ` + "`" + `false` + "`" + `. If ` + "`" + `false` + "`" + `, Atlas uses [ranged sharding](https://docs.mongodb.com/manual/core/ranged-sharding/). This is only available for Atlas clusters with MongoDB v4.4 and later.`,
				},
				resource.Attribute{
					Name:        "is_shard_key_unique",
					Description: `(Optional) Specifies whether the underlying index enforces a unique constraint. If omitted, defaults to false. You cannot specify true when using [hashed shard keys](https://docs.mongodb.com/manual/core/hashed-sharding/#std-label-sharding-hashed). ### Custom Zone Mapping`,
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
					Description: `A map of all custom zone mappings defined for the Global Cluster. Atlas automatically maps each location code to the closest geographical zone. Custom zone mappings allow administrators to override these automatic mappings. If your Global Cluster does not have any custom zone mappings, this document is empty. ## Import Global Clusters can be imported using project ID and cluster name, in the format ` + "`" + `PROJECTID-CLUSTER_NAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_global_cluster_config.config 1112222b3bf99403840e8934-Cluster0 ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Global Clusters](https://docs.atlas.mongodb.com/reference/api/global-clusters/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "custom_zone_mapping",
					Description: `A map of all custom zone mappings defined for the Global Cluster. Atlas automatically maps each location code to the closest geographical zone. Custom zone mappings allow administrators to override these automatic mappings. If your Global Cluster does not have any custom zone mappings, this document is empty. ## Import Global Clusters can be imported using project ID and cluster name, in the format ` + "`" + `PROJECTID-CLUSTER_NAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_global_cluster_config.config 1112222b3bf99403840e8934-Cluster0 ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Global Clusters](https://docs.atlas.mongodb.com/reference/api/global-clusters/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_ldap_configuration",
			Category:         "Resources",
			ShortDescription: `Provides a LDAP Configuration resource.`,
			Description: `

` + "`" + `mongodbatlas_ldap_configuration` + "`" + ` provides an LDAP Configuration resource. This allows an LDAP configuration for an Atlas project to be crated and managed. This endpoint doesn’t verify connectivity using the provided LDAP over TLS configuration details. To verify a configuration before saving it, use the resource to [verify](https://github.com/mongodb/terraform-provider-mongodbatlas/blob/INTMDB-114/website/docs/r/ldap_verify.html.markdown) the LDAP configuration.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to configure LDAP.`,
				},
				resource.Attribute{
					Name:        "authentication_enabled",
					Description: `(Required) Specifies whether user authentication with LDAP is enabled.`,
				},
				resource.Attribute{
					Name:        "authorization_enabled",
					Description: `(Optional) Specifies whether user authorization with LDAP is enabled. You cannot enable user authorization with LDAP without first enabling user authentication with LDAP.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The hostname or IP address of the LDAP server. The server must be visible to the internet or connected to your Atlas cluster with VPC Peering.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port to which the LDAP server listens for client connections. Default: ` + "`" + `636` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bind_username",
					Description: `(Required) The user DN that Atlas uses to connect to the LDAP server. Must be the full DN, such as ` + "`" + `CN=BindUser,CN=Users,DC=myldapserver,DC=mycompany,DC=com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bind_password",
					Description: `(Required) The password used to authenticate the ` + "`" + `bind_username` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ca_certificate",
					Description: `(Optional) CA certificate used to verify the identify of the LDAP server. Self-signed certificates are allowed.`,
				},
				resource.Attribute{
					Name:        "authz_query_template",
					Description: `(Optional) An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs. Used only for user authorization. Use the {USER} placeholder in the URL to substitute the authenticated username. The query is relative to the host specified with hostname. The formatting for the query must conform to RFC4515 and RFC 4516. If you do not provide a query template, Atlas attempts to use the default value: ` + "`" + `{USER}?memberOf?base` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_to_dn_mapping",
					Description: `(Optional) Maps an LDAP username for authentication to an LDAP Distinguished Name (DN). Each document contains a ` + "`" + `match` + "`" + ` regular expression and either a ` + "`" + `substitution` + "`" + ` or ` + "`" + `ldap_query` + "`" + ` template used to transform the LDAP username extracted from the regular expression. Atlas steps through the each document in the array in the given order, checking the authentication username against the ` + "`" + `match` + "`" + ` filter. If a match is found, Atlas applies the transformation and uses the output to authenticate the user. Atlas does not check the remaining documents in the array. For more details and examples see the [MongoDB Atlas API Reference](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-save/).`,
				},
				resource.Attribute{
					Name:        "user_to_dn_mapping.0.match",
					Description: `(Optional) A regular expression to match against a provided LDAP username. Each parenthesis-enclosed section represents a regular expression capture group used by the ` + "`" + `substitution` + "`" + ` or ` + "`" + `ldap_query` + "`" + ` template.`,
				},
				resource.Attribute{
					Name:        "user_to_dn_mapping.0.substitution",
					Description: `(Optional) An LDAP Distinguished Name (DN) formatting template that converts the LDAP name matched by the ` + "`" + `match` + "`" + ` regular expression into an LDAP Distinguished Name. Each bracket-enclosed numeric value is replaced by the corresponding regular expression capture group extracted from the LDAP username that matched the ` + "`" + `match` + "`" + ` regular expression.`,
				},
				resource.Attribute{
					Name:        "user_to_dn_mapping.0.ldap_query",
					Description: `(Optional) An LDAP query formatting template that inserts the LDAP name matched by the ` + "`" + `match` + "`" + ` regular expression into an LDAP query URI as specified by RFC 4515 and RFC 4516. Each numeric value is replaced by the corresponding regular expression capture group extracted from the LDAP username that matched the ` + "`" + `match` + "`" + ` regular expression. ## Import LDAP Configuration must be imported using project ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_ldap_configuration.test 5d09d6a59ccf6445652a444a ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-save)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_ldap_verify",
			Category:         "Resources",
			ShortDescription: `Provides a LDAP Verify resource.`,
			Description: `

` + "`" + `mongodbatlas_ldap_verify` + "`" + ` provides an LDAP Verify resource. This allows a a verification of an LDAP configuration over TLS for an Atlas project. Atlas retains only the most recent request for each project.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to configure LDAP.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The hostname or IP address of the LDAP server. The server must be visible to the internet or connected to your Atlas cluster with VPC Peering.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port to which the LDAP server listens for client connections. Default: ` + "`" + `636` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bind_username",
					Description: `(Required) The user DN that Atlas uses to connect to the LDAP server. Must be the full DN, such as ` + "`" + `CN=BindUser,CN=Users,DC=myldapserver,DC=mycompany,DC=com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bind_password",
					Description: `(Required) The password used to authenticate the ` + "`" + `bind_username` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ca_certificate",
					Description: `(Optional) CA certificate used to verify the identify of the LDAP server. Self-signed certificates are allowed.`,
				},
				resource.Attribute{
					Name:        "authz_query_template",
					Description: `(Optional) An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs. Used only for user authorization. Use the {USER} placeholder in the URL to substitute the authenticated username. The query is relative to the host specified with hostname. The formatting for the query must conform to RFC4515 and RFC 4516. If you do not provide a query template, Atlas attempts to use the default value: ` + "`" + `{USER}?memberOf?base` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "request_id",
					Description: `The unique identifier for the request to verify the LDAP over TLS/SSL configuration.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the LDAP over TLS/SSL configuration. One of the following values: ` + "`" + `PENDING` + "`" + `, ` + "`" + `SUCCESS` + "`" + `, and ` + "`" + `FAILED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `One or more links to sub-resources. The relations in the URLs are explained in the Web Linking Specification.`,
				},
				resource.Attribute{
					Name:        "validations",
					Description: `Array of validation messages related to the verification of the provided LDAP over TLS/SSL configuration details. The array contains a document for each test that Atlas runs. Atlas stops running tests after the first failure. The following return values can be seen here: [Values](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-request-verification) ## Import LDAP Configuration must be imported using project ID and request ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_ldap_verify.test 5d09d6a59ccf6445652a444a-5d09d6a59ccf6445652a444a ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-request-verification)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "request_id",
					Description: `The unique identifier for the request to verify the LDAP over TLS/SSL configuration.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the LDAP over TLS/SSL configuration. One of the following values: ` + "`" + `PENDING` + "`" + `, ` + "`" + `SUCCESS` + "`" + `, and ` + "`" + `FAILED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `One or more links to sub-resources. The relations in the URLs are explained in the Web Linking Specification.`,
				},
				resource.Attribute{
					Name:        "validations",
					Description: `Array of validation messages related to the verification of the provided LDAP over TLS/SSL configuration details. The array contains a document for each test that Atlas runs. Atlas stops running tests after the first failure. The following return values can be seen here: [Values](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-request-verification) ## Import LDAP Configuration must be imported using project ID and request ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_ldap_verify.test 5d09d6a59ccf6445652a444a-5d09d6a59ccf6445652a444a ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-request-verification)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_maintenance_window",
			Category:         "Resources",
			ShortDescription: `Provides an Maintenance Window resource.`,
			Description: `

` + "`" + `mongodbatlas_maintenance_window` + "`" + ` provides a resource to schedule a maintenance window for your MongoDB Atlas Project and/or set to defer a scheduled maintenance up to two times.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
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
					Description: `Defer the next scheduled maintenance for the given project for one week.`,
				},
				resource.Attribute{
					Name:        "auto_defer",
					Description: `Defer any scheduled maintenance for the given project for one week.`,
				},
				resource.Attribute{
					Name:        "auto_defer_once_enabled",
					Description: `Flag that indicates whether you want to defer all maintenance windows one week they would be triggered. ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_container",
			Category:         "Resources",
			ShortDescription: `Provides a Network Peering resource.`,
			Description: `

` + "`" + `mongodbatlas_network_container` + "`" + ` provides a Network Peering Container resource. The resource lets you create, edit and delete network peering containers. The resource requires your Project ID.  Each cloud provider requires slightly different attributes so read the argument reference carefully.

 Network peering container is a general term used to describe any cloud providers' VPC/VNet concept.  Containers only need to be created if the peering connection to the cloud provider will be created before the first cluster that requires the container.  If the cluster has been/will be created first Atlas automatically creates the required container per the "containers per cloud provider" information that follows (in this case you can obtain the container id from the cluster resource attribute ` + "`" + `container_id` + "`" + `).

The following is the maximum number of Network Peering containers per cloud provider:
<br> &#8226;  GCP -  One container per project.
<br> &#8226;  AWS and Azure - One container per cloud provider region.

-> **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.


`,
			Keywords: []string{},
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
					Description: `The name of the Azure VNet. Returns null. This value is populated once you create a new network peering connection with the network peering resource. ## Import Clusters can be imported using project ID and network peering container id, in the format ` + "`" + `PROJECTID-CONTAINER-ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_network_container.my_container 1112222b3bf99403840e8934-5cbf563d87d9d67253be590a ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/) ->`,
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
					Description: `The name of the Azure VNet. Returns null. This value is populated once you create a new network peering connection with the network peering resource. ## Import Clusters can be imported using project ID and network peering container id, in the format ` + "`" + `PROJECTID-CONTAINER-ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_network_container.my_container 1112222b3bf99403840e8934-5cbf563d87d9d67253be590a ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/) ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_peering",
			Category:         "Resources",
			ShortDescription: `Provides a Network Peering resource.`,
			Description: `

` + "`" + `mongodbatlas_network_peering` + "`" + ` provides a Network Peering Connection resource. The resource lets you create, edit and delete network peering connections. The resource requires your Project ID.  

Ensure you have first created a network container if it is required for your configuration.  See the network_container resource documentation to determine if you need a network container first.  Examples for creating both container and peering resource are shown below as well as examples for creating the peering connection only.

~> **GCP AND AZURE ONLY:** Connect via Peering Only mode is deprecated, so no longer needed.  See [disable Peering Only mode](https://docs.atlas.mongodb.com/reference/faq/connection-changes/#disable-peering-mode) for details and [private_ip_mode](https://www.terraform.io/docs/providers/mongodbatlas/r/private_ip_mode.html) to disable.

~> **AZURE ONLY:** To create the peering request with an Azure VNET, you must grant Atlas the following permissions on the virtual network.
    Microsoft.Network/virtualNetworks/virtualNetworkPeerings/read
    Microsoft.Network/virtualNetworks/virtualNetworkPeerings/write
    Microsoft.Network/virtualNetworks/virtualNetworkPeerings/delete
    Microsoft.Network/virtualNetworks/peer/action
For more information see https://docs.atlas.mongodb.com/security-vpc-peering/ and https://docs.atlas.mongodb.com/reference/api/vpc-create-peering-connection/

-> **Create a Whitelist:** Ensure you whitelist the private IP ranges of the subnets in which your application is hosted in order to connect to your Atlas cluster.  See the project_ip_whitelist resource.

-> **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.


`,
			Keywords: []string{},
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
					Description: `(Required - GCP) Name of the network peer to which Atlas connects.`,
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
					Description: `Name of your Azure VNet. ## Import Clusters can be imported using project ID and network peering id, in the format ` + "`" + `PROJECTID-PEERID-PROVIDERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_network_peering.my_peering 1112222b3bf99403840e8934-5cbf563d87d9d67253be590a-AWS ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Network Peering Connection](https://docs.atlas.mongodb.com/reference/api/vpc-create-peering-connection/) ->`,
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
					Description: `Name of your Azure VNet. ## Import Clusters can be imported using project ID and network peering id, in the format ` + "`" + `PROJECTID-PEERID-PROVIDERNAME` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_network_peering.my_peering 1112222b3bf99403840e8934-5cbf563d87d9d67253be590a-AWS ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Network Peering Connection](https://docs.atlas.mongodb.com/reference/api/vpc-create-peering-connection/) ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_online_archive",
			Category:         "Resources",
			ShortDescription: `Provides a Online Archive resource for creation, update, and delete`,
			Description: `

` + "`" + `mongodbatlas_online_archive` + "`" + ` resource provides access to create, edit, pause and resume an online archive for a collection. 

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

~> **IMPORTANT:** The collection must exists before performing an online archive.

~> **IMPORTANT:** There are fields that are immutable after creation, i.e if ` + "`" + `date_field` + "`" + ` value does not exist in the collection, the online archive state will be pending forever, and this field cannot be updated, that means a destroy is required, known error ` + "`" + `ONLINE_ARCHIVE_CANNOT_MODIFY_FIELD` + "`" + `

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "partition_fields",
					Description: `(Recommended) Fields to use to partition data. You can specify up to two frequently queried fields to use for partitioning data. Note that queries that don’t contain the specified fields will require a full collection scan of all archived documents, which will take longer and increase your costs. To learn more about how partition improves query performance, see [Data Structure in S3](https://docs.mongodb.com/datalake/admin/optimize-query-performance/#data-structure-in-s3). The value of a partition field can be up to a maximum of 700 characters. Documents with values exceeding 700 characters are not archived.`,
				},
				resource.Attribute{
					Name:        "criteria.expire_after_days",
					Description: `Number of days that specifies the age limit for the data in the live Atlas cluster. Data is archived when the current date is greater than the value of the date field specified via the ` + "`" + `date_field` + "`" + ` parameter plus the number of days specified here. The only field required for criteria type ` + "`" + `CUSTOM` + "`" + ``,
				},
				resource.Attribute{
					Name:        "criteria.query",
					Description: `JSON query to use to select documents for archiving. Atlas uses the specified query with the db.collection.find(query) command. The empty document {} to return all documents is not supported. ### Partition fields details`,
				},
				resource.Attribute{
					Name:        "partition_fields.field_name",
					Description: `(Required) Name of the field. To specify a nested field, use the dot notation.`,
				},
				resource.Attribute{
					Name:        "partition_fields.order",
					Description: `(Required) Position of the field in the partition. Value can be: 0,1,2 By default, the date field specified in the criteria.dateField parameter is in the first position of the partition.`,
				},
				resource.Attribute{
					Name:        "partitio_fields.field_type",
					Description: `(Optional) type of the partition field ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "archive_id",
					Description: `ID of the online archive.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "archive_id",
					Description: `ID of the online archive.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_org_invitation",
			Category:         "Resources",
			ShortDescription: `Provides an Atlas Organization Invitation resource.`,
			Description: `

` + "`" + `mongodbatlas_org_invitation` + "`" + ` invites a user to join an Atlas organization.

Each invitation for an Atlas user includes roles that Atlas grants the user when they accept the invitation.

The [MongoDB Documentation](https://docs.atlas.mongodb.com/reference/user-roles/#organization-roles) describes the roles a user can have, which map to:

* ORG_OWNER
* ORG_GROUP_CREATOR
* ORG_BILLING_ADMIN
* ORG_READ_ONLY
* ORG_MEMBER

~> **IMPORTANT:** This resource is only for managing invitations, not for managing the Atlas User being invited. Possible provider behavior depending on the invitee's action:
* If the user has not yet accepted the invitation, the provider leaves the invitation as is.
* If the user has accepted the invitation and is now an organization member, the provider will remove the invitation from the Terraform state.  The invitation must then be removed from the Terraform resource configuration.
* If the user accepts the invitation and then leaves the organization, the provider will re-add the invitation if the resource definition is not removed from the Terraform configuration.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the organization to which you want to invite a user.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.`,
				},
				resource.Attribute{
					Name:        "teams_ids",
					Description: `(Optional) An array of unique 24-hexadecimal digit strings that identify the teams that the user was invited to join.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The following options are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated unique string that identifies this resource.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp in ISO 8601 date and time format in UTC when Atlas sent the invitation.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `Timestamp in ISO 8601 date and time format in UTC when the invitation expires. Users have 30 days to accept an invitation.`,
				},
				resource.Attribute{
					Name:        "invitation_id",
					Description: `Unique 24-hexadecimal digit string that identifies the invitation in Atlas.`,
				},
				resource.Attribute{
					Name:        "inviter_username",
					Description: `Atlas user who invited ` + "`" + `username` + "`" + ` to the organization. ## Import ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated unique string that identifies this resource.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp in ISO 8601 date and time format in UTC when Atlas sent the invitation.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `Timestamp in ISO 8601 date and time format in UTC when the invitation expires. Users have 30 days to accept an invitation.`,
				},
				resource.Attribute{
					Name:        "invitation_id",
					Description: `Unique 24-hexadecimal digit string that identifies the invitation in Atlas.`,
				},
				resource.Attribute{
					Name:        "inviter_username",
					Description: `Atlas user who invited ` + "`" + `username` + "`" + ` to the organization. ## Import ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_private_endpoint_regional_mode",
			Category:         "Resources",
			ShortDescription: `Provides a Private Endpoint Regional Mode resource`,
			Description: `

` + "`" + `mongodbatlas_private_endpoint_regional_mode` + "`" + ` provides a Private Endpoint Regional Mode resource. This represents a regionalized private endpoint setting for a Project. Enable it to allow region specific private endpoints.

~> **IMPORTANT:**You must have one of the following roles to successfully handle the resource:
  * Organization Owner
  * Project Owner

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

~> **WARNING:**Your [connection strings](https://www.mongodb.com/docs/atlas/reference/faq/connection-changes/#std-label-connstring-privatelink) to existing multi-region and global sharded clusters change when you enable this setting.  You must update your applications to use the new connection strings. This might cause downtime.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique identifier for the project.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Flag that indicates whether the regionalized private endpoint setting is enabled for the project. Set this value to true to create more than one private endpoint in a cloud provider region to connect to multi-region and global Atlas sharded clusters. You can enable this setting only if your Atlas project contains no replica sets. You can't disable this setting if you have:`,
				},
				resource.Attribute{
					Name:        "mongodbatlas_cluster.cluster-atlas.depends_on",
					Description: `Make your cluster dependent on the project's ` + "`" + `mongodbatlas_private_endpoint_regional_mode` + "`" + ` as well as any relevant ` + "`" + `mongodbatlas_privatelink_endpoint_service` + "`" + ` resources. See an [example](https://github.com/mongodb/terraform-provider-mongodbatlas/tree/master/examples/aws-privatelink-endpoint/cluster-geosharded).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_private_ip_mode",
			Category:         "Resources",
			ShortDescription: `Provides a Private IP Mode resource.`,
			Description: `

` + "`" + `mongodbatlas_private_ip_mode` + "`" + ` provides a Private IP Mode resource. This allows one to disable Connect via Peering Only mode for a MongoDB Atlas Project.

~> **Deprecated Feature**: <br> This feature has been deprecated. Use [Split Horizon connection strings](https://dochub.mongodb.org/core/atlas-horizon-faq) to connect to your cluster. These connection strings allow you to connect using both VPC/VNet Peering and whitelisted public IP addresses. To learn more about support for Split Horizon, see [this FAQ](https://dochub.mongodb.org/core/atlas-horizon-faq). You need this endpoint to [disable Peering Only](https://docs.atlas.mongodb.com/reference/faq/connection-changes/#disable-peering-mode).


`,
			Keywords: []string{},
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
			Type:             "mongodbatlas_privatelink_endpoint",
			Category:         "Resources",
			ShortDescription: `Provides a Private Endpoint resource.`,
			Description: `

` + "`" + `mongodbatlas_privatelink_endpoint` + "`" + ` provides a Private Endpoint resource. This represents a Private Endpoint Service that can be created in an Atlas project.

~> **IMPORTANT:**You must have one of the following roles to successfully handle the resource:
  * Organization Owner
  * Project Owner

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

-> **NOTE:** A network container is created for a private endpoint to reside in if one does not yet exist in the project.  


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `Required Unique identifier for the project.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) Name of the cloud provider for which you want to create the private endpoint service. Atlas accepts ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + ` or ` + "`" + `GCP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Cloud provider region in which you want to create the private endpoint connection. Accepted values are: [AWS regions](https://docs.atlas.mongodb.com/reference/amazon-aws/#amazon-aws), [AZURE regions](https://docs.atlas.mongodb.com/reference/microsoft-azure/#microsoft-azure) and [GCP regions](https://docs.atlas.mongodb.com/reference/google-gcp/#std-label-google-gcp)`,
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
					Name:        "error_message",
					Description: `Error message pertaining to the AWS PrivateLink connection. Returns null if there are no errors. AWS:`,
				},
				resource.Attribute{
					Name:        "endpoint_service_name",
					Description: `Name of the PrivateLink endpoint service in AWS. Returns null while the endpoint service is being created.`,
				},
				resource.Attribute{
					Name:        "interface_endpoints",
					Description: `Unique identifiers of the interface endpoints in your VPC that you added to the AWS PrivateLink connection. AZURE:`,
				},
				resource.Attribute{
					Name:        "private_endpoints",
					Description: `All private endpoints that you have added to this Azure Private Link Service.`,
				},
				resource.Attribute{
					Name:        "private_link_service_name",
					Description: `Name of the Azure Private Link Service that Atlas manages. GCP:`,
				},
				resource.Attribute{
					Name:        "endpoint_group_names",
					Description: `GCP network endpoint groups corresponding to the Private Service Connect endpoint service.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `GCP region for the Private Service Connect endpoint service.`,
				},
				resource.Attribute{
					Name:        "service_attachment_names",
					Description: `Unique alphanumeric and special character strings that identify the service attachments associated with the GCP Private Service Connect endpoint service. Returns an empty list while Atlas creates the service attachments.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the AWS PrivateLink connection or Status of the Azure Private Link Service. Atlas returns one of the following values: AWS:`,
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
					Name:        "error_message",
					Description: `Error message pertaining to the AWS PrivateLink connection. Returns null if there are no errors. AWS:`,
				},
				resource.Attribute{
					Name:        "endpoint_service_name",
					Description: `Name of the PrivateLink endpoint service in AWS. Returns null while the endpoint service is being created.`,
				},
				resource.Attribute{
					Name:        "interface_endpoints",
					Description: `Unique identifiers of the interface endpoints in your VPC that you added to the AWS PrivateLink connection. AZURE:`,
				},
				resource.Attribute{
					Name:        "private_endpoints",
					Description: `All private endpoints that you have added to this Azure Private Link Service.`,
				},
				resource.Attribute{
					Name:        "private_link_service_name",
					Description: `Name of the Azure Private Link Service that Atlas manages. GCP:`,
				},
				resource.Attribute{
					Name:        "endpoint_group_names",
					Description: `GCP network endpoint groups corresponding to the Private Service Connect endpoint service.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `GCP region for the Private Service Connect endpoint service.`,
				},
				resource.Attribute{
					Name:        "service_attachment_names",
					Description: `Unique alphanumeric and special character strings that identify the service attachments associated with the GCP Private Service Connect endpoint service. Returns an empty list while Atlas creates the service attachments.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the AWS PrivateLink connection or Status of the Azure Private Link Service. Atlas returns one of the following values: AWS:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_privatelink_endpoint_serverless",
			Category:         "Resources",
			ShortDescription: `Describes a Serverless PrivateLink Endpoint`,
			Description: `

` + "`" + `privatelink_endpoint_serverless` + "`" + ` Provides a Serverless PrivateLink Endpoint resource.
This is the first of two resources required to configure PrivateLink for Serverless, the second is [mongodbatlas_privatelink_endpoint_service_serverless](https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/resources/privatelink_endpoint_service_serverless).

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique 24-digit hexadecimal string that identifies the project.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) Human-readable label that identifies the serverless instance.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) Cloud provider name; AWS is currently supported ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "endpoint_id",
					Description: `Unique 24-hexadecimal digit string that identifies the private endpoint.`,
				},
				resource.Attribute{
					Name:        "endpoint_service_name",
					Description: `Unique string that identifies the PrivateLink endpoint service.`,
				},
				resource.Attribute{
					Name:        "private_link_service_resource_id",
					Description: `Root-relative path that identifies the Azure Private Link Service that MongoDB Cloud manages.`,
				},
				resource.Attribute{
					Name:        "cloud_provider_endpoint_id",
					Description: `Unique string that identifies the private endpoint's network interface.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Human-readable string to associate with this private endpoint.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Human-readable label that indicates the current operating status of the private endpoint. Values include: RESERVATION_REQUESTED, RESERVED, INITIATING, AVAILABLE, FAILED, DELETING.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_id",
					Description: `Unique 24-hexadecimal digit string that identifies the private endpoint.`,
				},
				resource.Attribute{
					Name:        "endpoint_service_name",
					Description: `Unique string that identifies the PrivateLink endpoint service.`,
				},
				resource.Attribute{
					Name:        "private_link_service_resource_id",
					Description: `Root-relative path that identifies the Azure Private Link Service that MongoDB Cloud manages.`,
				},
				resource.Attribute{
					Name:        "cloud_provider_endpoint_id",
					Description: `Unique string that identifies the private endpoint's network interface.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Human-readable string to associate with this private endpoint.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Human-readable label that indicates the current operating status of the private endpoint. Values include: RESERVATION_REQUESTED, RESERVED, INITIATING, AVAILABLE, FAILED, DELETING.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_privatelink_endpoint_service",
			Category:         "Resources",
			ShortDescription: `Provides a Private Endpoint Link resource.`,
			Description: `

` + "`" + `mongodbatlas_privatelink_endpoint_service` + "`" + ` provides a Private Endpoint Interface Link resource. This represents a Private Endpoint Interface Link, which adds one interface endpoint to a private endpoint connection in an Atlas project.

~> **IMPORTANT:**You must have one of the following roles to successfully handle the resource:
  * Organization Owner
  * Project Owner

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.
-> **NOTE:** Create and delete wait for all clusters on the project to IDLE in order for their operations to complete. This ensures the latest connection strings can be retrieved following creation or deletion of this resource. Default timeout is 2hrs.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique identifier for the project.`,
				},
				resource.Attribute{
					Name:        "private_link_id",
					Description: `(Required) Unique identifier of the ` + "`" + `AWS` + "`" + ` or ` + "`" + `AZURE` + "`" + ` PrivateLink connection which is created by ` + "`" + `mongodbatlas_privatelink_endpoint` + "`" + ` resource.`,
				},
				resource.Attribute{
					Name:        "endpoint_service_id",
					Description: `(Required) Unique identifier of the interface endpoint you created in your VPC with the ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + ` or ` + "`" + `GCP` + "`" + ` resource.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) Cloud provider for which you want to create a private endpoint. Atlas accepts ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + ` or ` + "`" + `GCP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_endpoint_ip_address",
					Description: `(Optional) Private IP address of the private endpoint network interface you created in your Azure VNet. Only for ` + "`" + `AZURE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `(Optional) Unique identifier of the GCP project in which you created your endpoints. Only for ` + "`" + `GCP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `(Optional) Collection of individual private endpoints that comprise your endpoint group. Only for ` + "`" + `GCP` + "`" + `. See below.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Private IP address of the endpoint you created in GCP.`,
				},
				resource.Attribute{
					Name:        "endpoint_name",
					Description: `(Optional) Forwarding rule that corresponds to the endpoint you created in GCP. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "interface_endpoint_id",
					Description: `Unique identifier of the interface endpoint.`,
				},
				resource.Attribute{
					Name:        "private_endpoint_connection_name",
					Description: `Name of the connection for this private endpoint that Atlas generates.`,
				},
				resource.Attribute{
					Name:        "private_endpoint_ip_address",
					Description: `Private IP address of the private endpoint network interface.`,
				},
				resource.Attribute{
					Name:        "private_endpoint_resource_id",
					Description: `Unique identifier of the private endpoint.`,
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
					Name:        "aws_connection_status",
					Description: `Status of the interface endpoint for AWS. Returns one of the following values:`,
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
					Description: `Atlas is removing the interface endpoint from the private endpoint connection.`,
				},
				resource.Attribute{
					Name:        "azure_status",
					Description: `Status of the interface endpoint for AZURE. Returns one of the following values:`,
				},
				resource.Attribute{
					Name:        "INITIATING",
					Description: `Atlas has not yet accepted the connection to your private endpoint.`,
				},
				resource.Attribute{
					Name:        "AVAILABLE",
					Description: `Atlas approved the connection to your private endpoint.`,
				},
				resource.Attribute{
					Name:        "FAILED",
					Description: `Atlas failed to accept the connection your private endpoint.`,
				},
				resource.Attribute{
					Name:        "DELETING",
					Description: `Atlas is removing the connection to your private endpoint from the Private Link service.`,
				},
				resource.Attribute{
					Name:        "gcp_status",
					Description: `Status of the interface endpoint for GCP. Returns one of the following values:`,
				},
				resource.Attribute{
					Name:        "INITIATING",
					Description: `Atlas has not yet accepted the connection to your private endpoint.`,
				},
				resource.Attribute{
					Name:        "AVAILABLE",
					Description: `Atlas approved the connection to your private endpoint.`,
				},
				resource.Attribute{
					Name:        "FAILED",
					Description: `Atlas failed to accept the connection your private endpoint.`,
				},
				resource.Attribute{
					Name:        "DELETING",
					Description: `Atlas is removing the connection to your private endpoint from the Private Link service.`,
				},
				resource.Attribute{
					Name:        "endpoint_group_name",
					Description: `(Optional) Unique identifier of the endpoint group. The endpoint group encompasses all of the endpoints that you created in GCP.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `Collection of individual private endpoints that comprise your network endpoint group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the endpoint. Atlas returns one of the [values shown above](https://docs.atlas.mongodb.com/reference/api/private-endpoints-endpoint-create-one/#std-label-ref-status-field).`,
				},
				resource.Attribute{
					Name:        "service_attachment_name",
					Description: `Unique alphanumeric and special character strings that identify the service attachment associated with the endpoint. ## Import Private Endpoint Link Connection can be imported using project ID and username, in the format ` + "`" + `{project_id}--{private_link_id}--{endpoint_service_id}--{provider_name}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_privatelink_endpoint_service.test 1112222b3bf99403840e8934--3242342343112--vpce-4242342343--AWS ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Private Endpoint Link Connection](https://docs.atlas.mongodb.com/reference/api/private-endpoints-endpoint-create-one/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "interface_endpoint_id",
					Description: `Unique identifier of the interface endpoint.`,
				},
				resource.Attribute{
					Name:        "private_endpoint_connection_name",
					Description: `Name of the connection for this private endpoint that Atlas generates.`,
				},
				resource.Attribute{
					Name:        "private_endpoint_ip_address",
					Description: `Private IP address of the private endpoint network interface.`,
				},
				resource.Attribute{
					Name:        "private_endpoint_resource_id",
					Description: `Unique identifier of the private endpoint.`,
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
					Name:        "aws_connection_status",
					Description: `Status of the interface endpoint for AWS. Returns one of the following values:`,
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
					Description: `Atlas is removing the interface endpoint from the private endpoint connection.`,
				},
				resource.Attribute{
					Name:        "azure_status",
					Description: `Status of the interface endpoint for AZURE. Returns one of the following values:`,
				},
				resource.Attribute{
					Name:        "INITIATING",
					Description: `Atlas has not yet accepted the connection to your private endpoint.`,
				},
				resource.Attribute{
					Name:        "AVAILABLE",
					Description: `Atlas approved the connection to your private endpoint.`,
				},
				resource.Attribute{
					Name:        "FAILED",
					Description: `Atlas failed to accept the connection your private endpoint.`,
				},
				resource.Attribute{
					Name:        "DELETING",
					Description: `Atlas is removing the connection to your private endpoint from the Private Link service.`,
				},
				resource.Attribute{
					Name:        "gcp_status",
					Description: `Status of the interface endpoint for GCP. Returns one of the following values:`,
				},
				resource.Attribute{
					Name:        "INITIATING",
					Description: `Atlas has not yet accepted the connection to your private endpoint.`,
				},
				resource.Attribute{
					Name:        "AVAILABLE",
					Description: `Atlas approved the connection to your private endpoint.`,
				},
				resource.Attribute{
					Name:        "FAILED",
					Description: `Atlas failed to accept the connection your private endpoint.`,
				},
				resource.Attribute{
					Name:        "DELETING",
					Description: `Atlas is removing the connection to your private endpoint from the Private Link service.`,
				},
				resource.Attribute{
					Name:        "endpoint_group_name",
					Description: `(Optional) Unique identifier of the endpoint group. The endpoint group encompasses all of the endpoints that you created in GCP.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `Collection of individual private endpoints that comprise your network endpoint group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the endpoint. Atlas returns one of the [values shown above](https://docs.atlas.mongodb.com/reference/api/private-endpoints-endpoint-create-one/#std-label-ref-status-field).`,
				},
				resource.Attribute{
					Name:        "service_attachment_name",
					Description: `Unique alphanumeric and special character strings that identify the service attachment associated with the endpoint. ## Import Private Endpoint Link Connection can be imported using project ID and username, in the format ` + "`" + `{project_id}--{private_link_id}--{endpoint_service_id}--{provider_name}` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_privatelink_endpoint_service.test 1112222b3bf99403840e8934--3242342343112--vpce-4242342343--AWS ` + "`" + `` + "`" + `` + "`" + ` See detailed information for arguments and attributes: [MongoDB API Private Endpoint Link Connection](https://docs.atlas.mongodb.com/reference/api/private-endpoints-endpoint-create-one/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_privatelink_endpoint_service_adl",
			Category:         "Resources",
			ShortDescription: `Provides an Atlas Data Lake and Online Archive PrivateLink endpoint.`,
			Description: `

` + "`" + `privatelink_endpoint_service_adl` + "`" + ` Provides an Atlas Data Lake (ADL) and Online Archive PrivateLink endpoint resource.   The same configuration will provide a PrivateLink connection for either Atlas Data Lake or Online Archive.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique 24-digit hexadecimal string that identifies the project.`,
				},
				resource.Attribute{
					Name:        "endpoint_id",
					Description: `(Required) Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [|aws| PrivateLink](https://aws.amazon.com/privatelink/) feature.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Human-readable label that identifies the type of resource to associate with this private endpoint. Atlas supports ` + "`" + `DATA_LAKE` + "`" + ` only. If empty, defaults to ` + "`" + `DATA_LAKE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) Human-readable label that identifies the cloud provider for this endpoint. Atlas supports AWS only. If empty, defaults to AWS.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Human-readable string to associate with this private endpoint. ## Import ADL privatelink endpoint can be imported using project ID and endpoint ID, in the format ` + "`" + `project_id` + "`" + `--` + "`" + `endpoint_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import privatelink_endpoint_service_adl.test 1112222b3bf99403840e8934--vpce-jjg5e24qp93513h03 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API - DataLake](https://docs.mongodb.com/datalake/reference/api/datalakes-api/) and [MongoDB Atlas API - Online Archive](https://docs.atlas.mongodb.com/reference/api/online-archive/) Documentation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_privatelink_endpoint_service_serverless",
			Category:         "Resources",
			ShortDescription: `Describes a Serverless PrivateLink Endpoint Service`,
			Description: `

` + "`" + `privatelink_endpoint_service_serverless` + "`" + ` Provides a Serverless PrivateLink Endpoint Service resource.
This is the second of two resources required to configure PrivateLink for Serverless, the first is [mongodbatlas_privatelink_endpoint_serverless](https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/resources/privatelink_endpoint_serverless).

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.
-> **NOTE:** Create waits for all serverless instances on the project to IDLE in order for their operations to complete. This ensures the latest connection strings can be retrieved following creation of this resource. Default timeout is 2hrs.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique 24-digit hexadecimal string that identifies the project.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) Human-readable label that identifies the serverless instance.`,
				},
				resource.Attribute{
					Name:        "endpoint_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the private endpoint.`,
				},
				resource.Attribute{
					Name:        "cloud_provider_endpoint_id",
					Description: `(Optional) Unique string that identifies the private endpoint's network interface.`,
				},
				resource.Attribute{
					Name:        "private_endpoint_ip_address",
					Description: `(Optional) IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) Cloud provider for which you want to create a private endpoint. Atlas accepts ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Human-readable string to associate with this private endpoint.`,
				},
				resource.Attribute{
					Name:        "endpoint_service_name",
					Description: `Unique string that identifies the PrivateLink endpoint service.`,
				},
				resource.Attribute{
					Name:        "private_link_service_resource_id",
					Description: `Root-relative path that identifies the Azure Private Link Service that MongoDB Cloud manages.`,
				},
				resource.Attribute{
					Name:        "private_endpoint_ip_address",
					Description: `IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service.`,
				},
				resource.Attribute{
					Name:        "cloud_provider_endpoint_id",
					Description: `Unique string that identifies the private endpoint's network interface.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Human-readable string to associate with this private endpoint.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `Human-readable error message that indicates the error condition associated with establishing the private endpoint connection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Human-readable label that indicates the current operating status of the private endpoint. Values include: RESERVATION_REQUESTED, RESERVED, INITIATING, AVAILABLE, FAILED, DELETING. ## Import Serverless privatelink endpoint can be imported using project ID and endpoint ID, in the format ` + "`" + `project_id` + "`" + `--` + "`" + `endpoint_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_privatelink_endpoint_service_serverless.test 1112222b3bf99403840e8934--serverless_name--vpce-jjg5e24qp93513h03 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API - Serverless Private Endpoints](https://www.mongodb.com/docs/atlas/reference/api/serverless-private-endpoints-get-one/).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_service_name",
					Description: `Unique string that identifies the PrivateLink endpoint service.`,
				},
				resource.Attribute{
					Name:        "private_link_service_resource_id",
					Description: `Root-relative path that identifies the Azure Private Link Service that MongoDB Cloud manages.`,
				},
				resource.Attribute{
					Name:        "private_endpoint_ip_address",
					Description: `IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service.`,
				},
				resource.Attribute{
					Name:        "cloud_provider_endpoint_id",
					Description: `Unique string that identifies the private endpoint's network interface.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Human-readable string to associate with this private endpoint.`,
				},
				resource.Attribute{
					Name:        "error_message",
					Description: `Human-readable error message that indicates the error condition associated with establishing the private endpoint connection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Human-readable label that indicates the current operating status of the private endpoint. Values include: RESERVATION_REQUESTED, RESERVED, INITIATING, AVAILABLE, FAILED, DELETING. ## Import Serverless privatelink endpoint can be imported using project ID and endpoint ID, in the format ` + "`" + `project_id` + "`" + `--` + "`" + `endpoint_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_privatelink_endpoint_service_serverless.test 1112222b3bf99403840e8934--serverless_name--vpce-jjg5e24qp93513h03 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API - Serverless Private Endpoints](https://www.mongodb.com/docs/atlas/reference/api/serverless-private-endpoints-get-one/).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_project",
			Category:         "Resources",
			ShortDescription: `Provides a Project resource.`,
			Description: `

` + "`" + `mongodbatlas_project` + "`" + ` provides a Project resource. This allows project to be created.

~> **IMPORTANT WARNING:**  Changing the name of an existing Project in your Terraform configuration will result the destruction of that Project and related resources (including Clusters) and the re-creation of those resources.  Terraform will inform you of the destroyed/created resources before applying so be sure to verify any change to your environment before applying.

-> **NOTE:** If Backup Compliance Policy is enabled for the project for which this backup schedule is defined, you cannot delete the Atlas project if any snapshots exist.  See [Backup Compliance Policy Prohibited Actions and Considerations](https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/#configure-a-backup-compliance-policy).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project you want to create. (Cannot be changed via this Provider after creation.)`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) The ID of the organization you want to create the project within.`,
				},
				resource.Attribute{
					Name:        "project_owner_id",
					Description: `(Optional) Unique 24-hexadecimal digit string that identifies the Atlas user account to be granted the [Project Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Project-Owner) role on the specified project. If you set this parameter, it overrides the default value of the oldest [Organization Owner](https://docs.atlas.mongodb.com/reference/user-roles/#mongodb-authrole-Organization-Owner).`,
				},
				resource.Attribute{
					Name:        "with_default_alerts_settings",
					Description: `(Optional) It allows users to disable the creation of the default alert settings. By default, this flag is set to true. ### Teams Teams attribute is optional ~>`,
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
					Name:        "api_key_id",
					Description: `(Required) The unique identifier of the Programmatic API key you want to associate with the Project. The Programmatic API key and Project must share the same parent organization. Note: this is not the ` + "`" + `publicKey` + "`" + ` of the Programmatic API key but the ` + "`" + `id` + "`" + ` of the key. See [Programmatic API Keys](https://docs.atlas.mongodb.com/reference/api/apiKeys/) for more.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `(Required) List of Project roles that the Programmatic API key needs to have. Ensure you provide: at least one role and ensure all roles are valid for the Project. You must specify an array even if you are only associating a single role with the Programmatic API key. The following are valid roles:`,
				},
				resource.Attribute{
					Name:        "is_collect_database_specifics_statistics_enabled",
					Description: `(Optional) Flag that indicates whether to enable statistics in [cluster metrics](https://www.mongodb.com/docs/atlas/monitor-cluster-metrics/) collection for the project.`,
				},
				resource.Attribute{
					Name:        "is_data_explorer_enabled",
					Description: `(Optional) Flag that indicates whether to enable Data Explorer for the project. If enabled, you can query your database with an easy to use interface. When Data Explorer is disabled, you cannot terminate slow operations from the [Real-Time Performance Panel](https://www.mongodb.com/docs/atlas/real-time-performance-panel/#std-label-real-time-metrics-status-tab) or create indexes from the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor). You can still view Performance Advisor recommendations, but you must create those indexes from [mongosh](https://www.mongodb.com/docs/mongodb-shell/#mongodb-binary-bin.mongosh).`,
				},
				resource.Attribute{
					Name:        "is_performance_advisor_enabled",
					Description: `(Optional) Flag that indicates whether to enable Performance Advisor and Profiler for the project. If enabled, you can analyze database logs to recommend performance improvements.`,
				},
				resource.Attribute{
					Name:        "is_realtime_performance_panel_enabled",
					Description: `(Optional) Flag that indicates whether to enable Real Time Performance Panel for the project. If enabled, you can see real time metrics from your MongoDB database.`,
				},
				resource.Attribute{
					Name:        "is_schema_advisor_enabled",
					Description: `(Optional) Flag that indicates whether to enable Schema Advisor for the project. If enabled, you receive customized recommendations to optimize your data model and enhance performance. Disable this setting to disable schema suggestions in the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor) and the [Data Explorer](https://www.mongodb.com/docs/atlas/atlas-ui/#std-label-atlas-ui).`,
				},
				resource.Attribute{
					Name:        "region_usage_restrictions",
					Description: `(Optional - set value to GOV_REGIONS_ONLY) Designates that this project can be used for government regions only. If not set the project will default to standard regions. You cannot deploy clusters across government and standard regions in the same project. AWS is the only cloud provider for AtlasGov. For more information see [MongoDB Atlas for Government](https://www.mongodb.com/docs/atlas/government/api/#creating-a-project). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `The number of Atlas clusters deployed in the project.. ## Import Project must be imported using project ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_project.my_project 5d09d6a59ccf6445652a444a ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas Admin API Projects](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Projects) and [MongoDB Atlas Admin API Teams](https://docs.atlas.mongodb.com/reference/api/teams/) Documentation for more information.`,
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
					Description: `The number of Atlas clusters deployed in the project.. ## Import Project must be imported using project ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_project.my_project 5d09d6a59ccf6445652a444a ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas Admin API Projects](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Projects) and [MongoDB Atlas Admin API Teams](https://docs.atlas.mongodb.com/reference/api/teams/) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_project_api_key",
			Category:         "Resources",
			ShortDescription: `Provides a Project API Key resource.`,
			Description: `

` + "`" + `mongodbatlas_project_api_key` + "`" + ` provides a Project API Key resource. This allows project API Key to be created.

~> **IMPORTANT WARNING:** Managing Atlas Programmatic API Keys (PAKs) with Terraform will expose sensitive organizational secrets in Terraform's state. We suggest following [Terraform's best practices](https://developer.hashicorp.com/terraform/language/state/sensitive-data). You may also want to consider managing your PAKs via a more secure method, such as the [HashiCorp Vault MongoDB Atlas Secrets Engine](https://developer.hashicorp.com/vault/docs/secrets/mongodbatlas).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project__id",
					Description: `Unique identifier for the project whose API keys you want to retrieve. Use the /orgs endpoint to retrieve all organizations to which the authenticated user has access.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this Organization API key.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `(Required) List of Project roles that the Programmatic API key needs to have. Ensure you provide: at least one role and ensure all roles are valid for the Project. You must specify an array even if you are only associating a single role with the Programmatic API key. The following are valid roles:`,
				},
				resource.Attribute{
					Name:        "api_key_id",
					Description: `(Required) The unique identifier of the Programmatic API key you want to associate with the Project. The Programmatic API key and Project must share the same parent organization. Note: this is not the ` + "`" + `publicKey` + "`" + ` of the Programmatic API key but the ` + "`" + `id` + "`" + ` of the key. See [Programmatic API Keys](https://docs.atlas.mongodb.com/reference/api/apiKeys/) for more. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "api_key_id",
					Description: `Unique identifier for this Project API key. ## Import API Keys must be imported using org ID, API Key ID e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_project_api_key.test 5d09d6a59ccf6445652a444a-6576974933969669 ` + "`" + `` + "`" + `` + "`" + ` See [MongoDB Atlas API - API Key](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Programmatic-API-Keys/operation/createAndAssignOneOrganizationApiKeyToOneProject) - Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_key_id",
					Description: `Unique identifier for this Project API key. ## Import API Keys must be imported using org ID, API Key ID e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_project_api_key.test 5d09d6a59ccf6445652a444a-6576974933969669 ` + "`" + `` + "`" + `` + "`" + ` See [MongoDB Atlas API - API Key](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Programmatic-API-Keys/operation/createAndAssignOneOrganizationApiKeyToOneProject) - Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_project_invitation",
			Category:         "Resources",
			ShortDescription: `Provides an Atlas Project Invitation resource.`,
			Description: `

` + "`" + `mongodbatlas_project_invitation` + "`" + ` invites a user to join an Atlas project.

Each invitation for an Atlas user includes roles that Atlas grants the user when they accept the invitation.

The [MongoDB Documentation](https://www.mongodb.com/docs/atlas/reference/user-roles/#project-roles) describes the roles which can be assigned to a user.

-> **NOTE:** Groups and projects are synonymous terms. You may find GROUP-ID in the official documentation.

~> **IMPORTANT:** This resource is only for managing invitations, not for managing the Atlas User being invited. Possible provider behavior depending on the invitee's action:
* If the user has not yet accepted the invitation, the provider leaves the invitation as is.
* If the user has accepted the invitation and is now a project member, the provider will remove the invitation from the Terraform state.  The invitation must then be removed from the Terraform resource configuration.
* If the user accepts the invitation and then leaves the project, the provider will re-add the invitation if the resource definition is not removed from the Terraform configuration.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the project to which you want to invite a user.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Email address to which Atlas sent the invitation. The user uses this email address as their Atlas username if they accept this invitation.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) List of Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. Refer to the [MongoDB Documentation](https://www.mongodb.com/docs/atlas/reference/user-roles/#project-roles) for information on valid roles. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this resource.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp in ISO 8601 date and time format in UTC when Atlas sent the invitation.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `Timestamp in ISO 8601 date and time format in UTC when the invitation expires. Users have 30 days to accept an invitation.`,
				},
				resource.Attribute{
					Name:        "invitation_id",
					Description: `Unique 24-hexadecimal digit string that identifies the invitation in Atlas.`,
				},
				resource.Attribute{
					Name:        "inviter_username",
					Description: `Atlas user who invited ` + "`" + `username` + "`" + ` to the project. See the [MongoDB Atlas Administration API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#operation/inviteOneMongoDBCloudUserToJoinOneProject) documentation for more information. ## Import ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this resource.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp in ISO 8601 date and time format in UTC when Atlas sent the invitation.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `Timestamp in ISO 8601 date and time format in UTC when the invitation expires. Users have 30 days to accept an invitation.`,
				},
				resource.Attribute{
					Name:        "invitation_id",
					Description: `Unique 24-hexadecimal digit string that identifies the invitation in Atlas.`,
				},
				resource.Attribute{
					Name:        "inviter_username",
					Description: `Atlas user who invited ` + "`" + `username` + "`" + ` to the project. See the [MongoDB Atlas Administration API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#operation/inviteOneMongoDBCloudUserToJoinOneProject) documentation for more information. ## Import ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_project_ip_access_list",
			Category:         "Resources",
			ShortDescription: `Provides an IP Access List resource.`,
			Description: `

` + "`" + `mongodbatlas_project_ip_access_list` + "`" + ` provides an IP Access List entry resource. The access list grants access from IPs, CIDRs or AWS Security Groups (if VPC Peering is enabled) to clusters within the Project.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

~> **IMPORTANT:**
When you remove an entry from the access list, existing connections from the removed address(es) may remain open for a variable amount of time. How much time passes before Atlas closes the connection depends on several factors, including how the connection was established, the particular behavior of the application or driver using the address, and the connection protocol (e.g., TCP or UDP). This is particularly important to consider when changing an existing IP address or CIDR block as they cannot be updated via the Provider (comments can however), hence a change will force the destruction and recreation of entries.   


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique identifier for the project to which you want to add one or more access list entries.`,
				},
				resource.Attribute{
					Name:        "aws_security_group",
					Description: `(Optional) Unique identifier of the AWS security group to add to the access list. Your access list entry can include only one ` + "`" + `awsSecurityGroup` + "`" + `, one ` + "`" + `cidrBlock` + "`" + `, or one ` + "`" + `ipAddress` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) Range of IP addresses in CIDR notation to be added to the access list. Your access list entry can include only one ` + "`" + `awsSecurityGroup` + "`" + `, one ` + "`" + `cidrBlock` + "`" + `, or one ` + "`" + `ipAddress` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Single IP address to be added to the access list. Mutually exclusive with ` + "`" + `awsSecurityGroup` + "`" + ` and ` + "`" + `cidrBlock` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment to add to the access list entry. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages and can be used to import. ## Import IP Access List entries can be imported using the ` + "`" + `project_id` + "`" + ` and ` + "`" + `cidr_block` + "`" + ` or ` + "`" + `ip_address` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_project_ip_access_list.test 5d0f1f74cf09a29120e123cd-10.242.88.0/21 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/access-lists/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages and can be used to import. ## Import IP Access List entries can be imported using the ` + "`" + `project_id` + "`" + ` and ` + "`" + `cidr_block` + "`" + ` or ` + "`" + `ip_address` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_project_ip_access_list.test 5d0f1f74cf09a29120e123cd-10.242.88.0/21 ` + "`" + `` + "`" + `` + "`" + ` For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/access-lists/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_search_index",
			Category:         "Resources",
			ShortDescription: `Provides a Search Index resource.`,
			Description: `

` + "`" + `mongodbatlas_search_index` + "`" + ` provides a Search Index resource. This allows indexes to be created.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the search index you want to create.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the organization or project you want to create the search index within.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the cluster where you want to create the search index within.`,
				},
				resource.Attribute{
					Name:        "wait_for_index_build_completion",
					Description: `(Optional) Wait for search index to achieve Active status before terraform considers resource built.`,
				},
				resource.Attribute{
					Name:        "analyzer",
					Description: `[Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)`,
				},
				resource.Attribute{
					Name:        "analyzers",
					Description: `[Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index. This is an array of JSON objects. ` + "`" + `` + "`" + `` + "`" + ` analyzers = <<-EOF [{ "name": "index_analyzer_test_name", "charFilters": { "type": "mapping", "mappings": {"\\" : "/"} }, "tokenizer": { "type": "nGram", "minGram": 2, "maxGram": 5 }, "tokenFilters": { "type": "length", "min": 20, "max": 33 } }] EOF ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "collection_name",
					Description: `(Required) Name of the collection the index is on.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required) Name of the database the collection is in.`,
				},
				resource.Attribute{
					Name:        "mappings_dynamic",
					Description: `Indicates whether the index uses dynamic or static mapping. For dynamic mapping, set the value to ` + "`" + `true` + "`" + `. For static mapping, specify the fields to index using ` + "`" + `mappings_fields` + "`" + ``,
				},
				resource.Attribute{
					Name:        "mappings_fields",
					Description: `attribute is required when ` + "`" + `mappings_dynamic` + "`" + ` is false. This field needs to be a JSON string in order to be decoded correctly. ` + "`" + `` + "`" + `` + "`" + `terraform mappings_fields = <<-EOF { "address": { "type": "document", "fields": { "city": { "type": "string", "analyzer": "lucene.simple", "ignoreAbove": 255 }, "state": { "type": "string", "analyzer": "lucene.english" } } }, "company": { "type": "string", "analyzer": "lucene.whitespace", "multi": { "mySecondaryAnalyzer": { "type": "string", "analyzer": "lucene.french" } } }, "employees": { "type": "string", "analyzer": "lucene.standard" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "search_analyzer",
					Description: `[Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index. Defaults to [lucene.standard](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/standard/#std-label-ref-standard-analyzer)`,
				},
				resource.Attribute{
					Name:        "synonyms",
					Description: `Synonyms mapping definition to use in this index. ### Analyzers An [Atlas Search analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/) prepares a set of documents to be indexed by performing a series of operations to transform, filter, and group sequences of characters. You can define a custom analyzer to suit your specific indexing needs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the custom analyzer. Names must be unique within an index, and may`,
				},
				resource.Attribute{
					Name:        "charFilters",
					Description: `Array containing zero or more character filters. Always require a ` + "`" + `type` + "`" + ` field, and some take additional options as well ` + "`" + `` + "`" + `` + "`" + `terraform "charFilters":{ "type": "<FILTER_TYPE>", "ADDITIONAL_OPTION": VALUE } ` + "`" + `` + "`" + `` + "`" + ` Atlas search supports four ` + "`" + `types` + "`" + ` of character filters:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Must be ` + "`" + `htmlStrip` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tokenizer",
					Description: `(Required) Tokenizer to use. Determines how Atlas Search splits up text into discrete chunks of indexing. Always require a type field, and some take additional options as well. ` + "`" + `` + "`" + `` + "`" + `terraform "tokenizer":{ "type": "<tokenizer-type>", "ADDITIONAL_OPTIONS": VALUE } ` + "`" + `` + "`" + `` + "`" + ` Atlas Search supports the following tokenizer options:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be ` + "`" + `standard` + "`" + ``,
				},
				resource.Attribute{
					Name:        "maxTokenLength",
					Description: `Maximum length for a single token. Tokens greater than this length are split at ` + "`" + `maxTokenLength` + "`" + ` into multiple tokens.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be ` + "`" + `whitespace` + "`" + ``,
				},
				resource.Attribute{
					Name:        "maxTokenLength",
					Description: `Maximum length for a single token. Tokens greater than this length are split at ` + "`" + `maxTokenLength` + "`" + ` into multiple tokens.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be ` + "`" + `nGram` + "`" + ``,
				},
				resource.Attribute{
					Name:        "minGram",
					Description: `(Required) Number of characters to include in the shortest token created.`,
				},
				resource.Attribute{
					Name:        "maxGram",
					Description: `(Required) Number of characters to include in the longest token created.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be ` + "`" + `edgeGram` + "`" + ``,
				},
				resource.Attribute{
					Name:        "minGram",
					Description: `(Required) Number of characters to include in the shortest token created.`,
				},
				resource.Attribute{
					Name:        "maxGram",
					Description: `(Required) Number of characters to include in the longest token created.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be ` + "`" + `regexCaptureGroup` + "`" + ``,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Required) A regular expression to match against.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Required) Index of the character group within the matching expression to extract into tokens. Use 0 to extract all character groups.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be ` + "`" + `regexSplit` + "`" + ``,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Required) A regular expression to match against.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be ` + "`" + `uaxUrlEmail` + "`" + ``,
				},
				resource.Attribute{
					Name:        "token_filters",
					Description: `Array containing zero or more token filters. Always require a type field, and some take additional options as well: ` + "`" + `` + "`" + `` + "`" + `terraform "tokenFilters":{ "type": "<FILTER_TYPE>", "ADDITIONAL-OPTIONS": VALUE } ` + "`" + `` + "`" + `` + "`" + ` Atlas Search supports the following token filters:`,
				},
				resource.Attribute{
					Name:        "originalTokens",
					Description: `Specifies whether to include or omit the original tokens in the output of the token filter. Value can be one of the following:`,
				},
				resource.Attribute{
					Name:        "include",
					Description: `to include the original tokens with the encoded tokens in the output of the token filter. We recommend this value if you want queries on both the original tokens as well as the encoded forms.`,
				},
				resource.Attribute{
					Name:        "omit",
					Description: `to omit the original tokens and include only the encoded tokens in the output of the token filter. Use this value if you want to only query on the encoded forms of the original tokens.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be ` + "`" + `length` + "`" + ``,
				},
				resource.Attribute{
					Name:        "min",
					Description: `The minimum length of a token. Must be less than or equal to ` + "`" + `max` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `The maximum length of a token. Must be greater than or equal to ` + "`" + `min` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be 'icuNormalizer'.`,
				},
				resource.Attribute{
					Name:        "normalizationForm",
					Description: `Normalization form to apply. Accepted values are:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be ` + "`" + `nGram` + "`" + ``,
				},
				resource.Attribute{
					Name:        "minGram",
					Description: `(Required) The minimum length of generated n-grams. Must be less than or equal to ` + "`" + `maxGram` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "maxGram",
					Description: `(Required) The maximum length of generated n-grams. Must be greater than or equal to ` + "`" + `minGram` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "termNotInBounds",
					Description: `Accepted values are:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be ` + "`" + `edgeGram` + "`" + ``,
				},
				resource.Attribute{
					Name:        "minGram",
					Description: `(Required) The minimum length of generated n-grams. Must be less than or equal to ` + "`" + `max_gram` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "maxGram",
					Description: `(Required) The maximum length of generated n-grams. Must be greater than or equal to ` + "`" + `min_gram` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "termsNotInBounds",
					Description: `Accepted values are:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be ` + "`" + `shingle` + "`" + ``,
				},
				resource.Attribute{
					Name:        "minShingleSize",
					Description: `(Required) Minimum number of tokens per shingle. Must be less than or equal to ` + "`" + `maxShingleSize` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "maxShingleSize",
					Description: `(Required) Maximum number of tokens per shingle. Must be greater than or equal to ` + "`" + `minShingleSize` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be ` + "`" + `regex` + "`" + ``,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Required) Regular expression pattern to apply to each token.`,
				},
				resource.Attribute{
					Name:        "replacement",
					Description: `(Required) Replacement string to substitute wherever a matching pattern occurs.`,
				},
				resource.Attribute{
					Name:        "matches",
					Description: `(Required) Acceptable values are:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be ` + "`" + `snowballstemming` + "`" + ``,
				},
				resource.Attribute{
					Name:        "stemmerName",
					Description: `(Required) The following values are valid:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Must be ` + "`" + `stopword` + "`" + ``,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required) The list of stop words that correspond to the tokens to remove. Value must be one or more stop words.`,
				},
				resource.Attribute{
					Name:        "ignoreCase",
					Description: `The flag that indicates whether or not to ignore case of stop words when filtering the tokens to remove. The value can be one of the following:`,
				},
				resource.Attribute{
					Name:        "true",
					Description: `to ignore case and remove all tokens that match the specified stop words`,
				},
				resource.Attribute{
					Name:        "false",
					Description: `to be case-sensitive and remove only tokens that exactly match the specified case If omitted, defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the [synonym mapping definition](https://docs.atlas.mongodb.com/reference/atlas-search/synonyms/#std-label-synonyms-ref). Name must be unique in this index definition and it can't be an empty string.`,
				},
				resource.Attribute{
					Name:        "source_collection",
					Description: `(Required) Name of the source MongoDB collection for the synonyms. Documents in this collection must be in the format described in the [Synonyms Source Collection Documents](https://docs.atlas.mongodb.com/reference/atlas-search/synonyms/#std-label-synonyms-coll-spec).`,
				},
				resource.Attribute{
					Name:        "analyzer",
					Description: `(Required) Name of the [analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use with this synonym mapping. Atlas Search doesn't support these [custom analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) tokenizers and token filters in [analyzers used in synonym mappings](https://docs.atlas.mongodb.com/reference/atlas-search/synonyms/#options):`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_serverless_instance",
			Category:         "Resources",
			ShortDescription: `Provides a Serverless Instance resource.`,
			Description: `

` + "`" + `mongodbatlas_serverless_instance` + "`" + ` provides a Serverless Instance resource. This allows serverless instances to be created.

> **NOTE:**  Serverless instances do not support some Atlas features at this time.
For a full list of unsupported features, see [Serverless Instance Limitations](https://docs.atlas.mongodb.com/reference/serverless-instance-limitations/).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Human-readable label that identifies the serverless instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the organization or project you want to create the serverless instance within.`,
				},
				resource.Attribute{
					Name:        "provider_settings_backing_provider_name",
					Description: `(Required) Cloud service provider on which MongoDB Cloud provisioned the serverless instance.`,
				},
				resource.Attribute{
					Name:        "provider_settings_provider_name",
					Description: `(Required) Cloud service provider that applies to the provisioned the serverless instance.`,
				},
				resource.Attribute{
					Name:        "provider_settings_region_name",
					Description: `(Required) Human-readable label that identifies the physical location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases.`,
				},
				resource.Attribute{
					Name:        "continuous_backup_enabled",
					Description: `(Optional) Flag that indicates whether the serverless instance uses [Serverless Continuous Backup](https://www.mongodb.com/docs/atlas/configure-serverless-backup). If this parameter is false or not used, the serverless instance uses [Basic Backup](https://www.mongodb.com/docs/atlas/configure-serverless-backup).`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies the serverless instance.`,
				},
				resource.Attribute{
					Name:        "connection_strings_standard_srv",
					Description: `Public ` + "`" + `mongodb+srv://` + "`" + ` connection string that you can use to connect to this serverless instance.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Timestamp that indicates when MongoDB Cloud created the serverless instance. The timestamp displays in the ISO 8601 date and time format in UTC.`,
				},
				resource.Attribute{
					Name:        "mongo_db_version",
					Description: `Version of MongoDB that the serverless instance runs, in ` + "`" + `<major version>` + "`" + `.` + "`" + `<minor version>` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Stage of deployment of this serverless instance when the resource made its request.`,
				},
				resource.Attribute{
					Name:        "connection_strings_private_endpoint_srv",
					Description: `List of Serverless Private Endpoint Connections`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies the serverless instance.`,
				},
				resource.Attribute{
					Name:        "connection_strings_standard_srv",
					Description: `Public ` + "`" + `mongodb+srv://` + "`" + ` connection string that you can use to connect to this serverless instance.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Timestamp that indicates when MongoDB Cloud created the serverless instance. The timestamp displays in the ISO 8601 date and time format in UTC.`,
				},
				resource.Attribute{
					Name:        "mongo_db_version",
					Description: `Version of MongoDB that the serverless instance runs, in ` + "`" + `<major version>` + "`" + `.` + "`" + `<minor version>` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Stage of deployment of this serverless instance when the resource made its request.`,
				},
				resource.Attribute{
					Name:        "connection_strings_private_endpoint_srv",
					Description: `List of Serverless Private Endpoint Connections`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_teams",
			Category:         "Resources",
			ShortDescription: `Provides a Team resource.`,
			Description: `

` + "`" + `mongodbatlas_teams` + "`" + ` provides a Team resource. The resource lets you create, edit and delete Teams. Also, Teams can be assigned to multiple projects, and team members’ access to the project is determined by the team’s project role.

> **IMPORTANT:** MongoDB Atlas Team limits: max 250 teams in an organization and max 100 teams per project.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

MongoDB Atlas Team limits: max 250 teams in an organization and max 100 teams per project.

`,
			Keywords: []string{},
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
			Type:             "mongodbatlas_third_party_integration",
			Category:         "Resources",
			ShortDescription: `Provides a Third-Party Integration Settings resource.`,
			Description: `

` + "`" + `mongodbatlas_third_party_integration` + "`" + ` Provides a Third-Party Integration Settings for the given type.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

-> **WARNING:** This field type has values (NEW_RELIC, FLOWDOCK) that are deprecated and will be removed in 1.9.0 release release

-> **NOTE:** Slack integrations now use the OAuth2 verification method and must be initially configured, or updated from a legacy integration, through the Atlas third-party service integrations page. Legacy tokens will soon no longer be supported.[Read more about slack setup](https://docs.atlas.mongodb.com/tutorial/third-party-service-integrations/)

~> **IMPORTANT** Each project can only have one configuration per {INTEGRATION-TYPE}.

~> **IMPORTANT:** All arguments including the secrets will be stored in the raw state as plain-text. [Read more about sensitive data in state.](https://www.terraform.io/docs/state/sensitive-data.html)


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to get all Third-Party service integrations`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `Your Service Key.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `Your API Key.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Indicates which API URL to use, either "US", "EU", "US3", or "US5". Datadog will use "US" by default.`,
				},
				resource.Attribute{
					Name:        "license_key",
					Description: `Your License Key.`,
				},
				resource.Attribute{
					Name:        "write_token",
					Description: `Your Insights Insert Key.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `Your API Key.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Indicates which API URL to use, either US or EU. Opsgenie will use US by default.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `Your API Key.`,
				},
				resource.Attribute{
					Name:        "routing_key",
					Description: `An optional field for your Routing Key.`,
				},
				resource.Attribute{
					Name:        "flow_name",
					Description: `Your Flowdock Flow name.`,
				},
				resource.Attribute{
					Name:        "api_token",
					Description: `Your API Token.`,
				},
				resource.Attribute{
					Name:        "org_name",
					Description: `Your Flowdock organization name.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Your webhook URL.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `An optional field for your webhook secret.`,
				},
				resource.Attribute{
					Name:        "microsoft_teams_webhook_url",
					Description: `Your Microsoft Teams incoming webhook URL.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Your Prometheus username.`,
				},
				resource.Attribute{
					Name:        "service_discovery",
					Description: `Indicates which service discovery method is used, either file or http.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `Your Prometheus protocol scheme configured for requests.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether your cluster has Prometheus enabled. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used by terraform for internal management, which can also be used to import. ## Import Third-Party Integration Settings can be imported using project ID and the integration type, in the format ` + "`" + `project_id` + "`" + `-` + "`" + `type` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_database_user.my_user 1112222b3bf99403840e8934-OPS_GENIE ` + "`" + `` + "`" + `` + "`" + ` See [MongoDB Atlas API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Third-Party-Integrations/operation/createThirdPartyIntegration) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used by terraform for internal management, which can also be used to import. ## Import Third-Party Integration Settings can be imported using project ID and the integration type, in the format ` + "`" + `project_id` + "`" + `-` + "`" + `type` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import mongodbatlas_database_user.my_user 1112222b3bf99403840e8934-OPS_GENIE ` + "`" + `` + "`" + `` + "`" + ` See [MongoDB Atlas API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Third-Party-Integrations/operation/createThirdPartyIntegration) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_x509_authentication_database_user",
			Category:         "Resources",
			ShortDescription: `Provides a X509 Authentication Database User resource.`,
			Description: `

` + "`" + `mongodbatlas_x509_authentication_database_user` + "`" + ` provides a X509 Authentication Database User resource. The mongodbatlas_x509_authentication_database_user resource lets you manage MongoDB users who authenticate using X.509 certificates. You can manage these X.509 certificates or let Atlas do it for you.

| Management  | Description  |
|---|---|
| Atlas  | Atlas manages your Certificate Authority and can generate certificates for your MongoDB users. No additional X.509 configuration is required.  |
| Customer  |  You must provide a Certificate Authority and generate certificates for your MongoDB users. |

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

`,
			Keywords: []string{},
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

		"mongodbatlas_access_list_api_key":                     0,
		"mongodbatlas_advanced_cluster":                        1,
		"mongodbatlas_alert_configuration":                     2,
		"mongodbatlas_api_key":                                 3,
		"mongodbatlas_auditing":                                4,
		"mongodbatlas_backup_compliance_policy":                5,
		"mongodbatlas_cloud_backup_schedule":                   6,
		"mongodbatlas_cloud_backup_snapshot":                   7,
		"mongodbatlas_cloud_backup_snapshot_export_bucket":     8,
		"mongodbatlas_cloud_backup_snapshot_export_job":        9,
		"mongodbatlas_cloud_backup_snapshot_restore_job":       10,
		"mongodbatlas_cloud_provider_access":                   11,
		"mongodbatlas_cloud_provider_snapshot":                 12,
		"mongodbatlas_cloud_provider_snapshot_backup_policy":   13,
		"mongodbatlas_cloud_provider_snapshot_restore_job":     14,
		"mongodbatlas_cluster":                                 15,
		"mongodbatlas_custom_db_role":                          16,
		"mongodbatlas_custom_dns_configuration_cluster_aws":    17,
		"mongodbatlas_data_lake":                               18,
		"mongodbatlas_database_user":                           19,
		"mongodbatlas_encryption_at_rest":                      20,
		"mongodbatlas_event_trigger":                           21,
		"mongodbatlas_federated_settings_identity_provider":    22,
		"mongodbatlas_federated_settings_org_config":           23,
		"mongodbatlas_federated_settings_org_role_mapping":     24,
		"mongodbatlas_global_cluster_config":                   25,
		"mongodbatlas_ldap_configuration":                      26,
		"mongodbatlas_ldap_verify":                             27,
		"mongodbatlas_maintenance_window":                      28,
		"mongodbatlas_network_container":                       29,
		"mongodbatlas_network_peering":                         30,
		"mongodbatlas_online_archive":                          31,
		"mongodbatlas_org_invitation":                          32,
		"mongodbatlas_private_endpoint_regional_mode":          33,
		"mongodbatlas_private_ip_mode":                         34,
		"mongodbatlas_privatelink_endpoint":                    35,
		"mongodbatlas_privatelink_endpoint_serverless":         36,
		"mongodbatlas_privatelink_endpoint_service":            37,
		"mongodbatlas_privatelink_endpoint_service_adl":        38,
		"mongodbatlas_privatelink_endpoint_service_serverless": 39,
		"mongodbatlas_project":                                 40,
		"mongodbatlas_project_api_key":                         41,
		"mongodbatlas_project_invitation":                      42,
		"mongodbatlas_project_ip_access_list":                  43,
		"mongodbatlas_search_index":                            44,
		"mongodbatlas_serverless_instance":                     45,
		"mongodbatlas_teams":                                   46,
		"mongodbatlas_third_party_integration":                 47,
		"mongodbatlas_x509_authentication_database_user":       48,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
