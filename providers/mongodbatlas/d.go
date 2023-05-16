package mongodbatlas

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_access_list_api_key",
			Category:         "Data Sources",
			ShortDescription: `Provides an Access List API Key resource.`,
			Description: `

` + "`" + `mongodbatlas_access_list_api_key` + "`" + ` describes an Access List API Key entry resource. The access list grants access from IPs, CIDRs) to clusters within the Project.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

~> **IMPORTANT:**
When you remove an entry from the access list, existing connections from the removed address(es) may remain open for a variable amount of time. How much time passes before Atlas closes the connection depends on several factors, including how the connection was established, the particular behavior of the application or driver using the address, and the connection protocol (e.g., TCP or UDP). This is particularly important to consider when changing an existing IP address or CIDR block as they cannot be updated via the Provider (comments can however), hence a change will force the destruction and recreation of entries.   

~> **IMPORTANT WARNING:** Managing Atlas Programmatic API Keys (PAKs) with Terraform will expose sensitive organizational secrets in Terraform's state. We suggest following [Terraform's best practices](https://developer.hashicorp.com/terraform/language/state/sensitive-data). You may also want to consider managing your PAKs via a more secure method, such as the [HashiCorp Vault MongoDB Atlas Secrets Engine](https://developer.hashicorp.com/vault/docs/secrets/mongodbatlas).


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) Unique identifier for the Organization to which you want to retrieve one or more access list entries.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) Range of IP addresses in CIDR notation to be added to the access list.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Single IP address to be added to the access list.`,
				},
				resource.Attribute{
					Name:        "api_key_id",
					Description: `(Required) Unique identifier for the Organization API Key for which you want to retrieve an access list entry.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used by Terraform for internal management and can be used to import.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment to add to the access list entry. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/access-lists/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used by Terraform for internal management and can be used to import.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment to add to the access list entry. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/access-lists/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_access_list_api_keys",
			Category:         "Data Sources",
			ShortDescription: `Provides an Access List API Keys resource.`,
			Description: `

` + "`" + `mongodbatlas_access_list_api_keys` + "`" + ` describes an Access List API Key entry resource. The access list grants access from IPs, CIDRs) to clusters within the Project.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

~> **IMPORTANT:**
When you remove an entry from the access list, existing connections from the removed address(es) may remain open for a variable amount of time. How much time passes before Atlas closes the connection depends on several factors, including how the connection was established, the particular behavior of the application or driver using the address, and the connection protocol (e.g., TCP or UDP). This is particularly important to consider when changing an existing IP address or CIDR block as they cannot be updated via the Provider (comments can however), hence a change will force the destruction and recreation of entries.   

~> **IMPORTANT WARNING:** Managing Atlas Programmatic API Keys (PAKs) with Terraform will expose sensitive organizational secrets in Terraform's state. We suggest following [Terraform's best practices](https://developer.hashicorp.com/terraform/language/state/sensitive-data). You may also want to consider managing your PAKs via a more secure method, such as the [HashiCorp Vault MongoDB Atlas Secrets Engine](https://developer.hashicorp.com/vault/docs/secrets/mongodbatlas).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Projects. ### API Keys`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) Unique identifier for the Organization to which you want to retrieve one or more access list entries.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) Range of IP addresses in CIDR notation to be added to the access list.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Single IP address to be added to the access list.`,
				},
				resource.Attribute{
					Name:        "api_key_id",
					Description: `(Required) Unique identifier for the Organization API Key for which you want to retrieve an access list entry.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used by Terraform for internal management and can be used to import.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment to add to the access list entry. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/access-lists/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used by Terraform for internal management and can be used to import.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment to add to the access list entry. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/access-lists/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_advanced_cluster",
			Category:         "Data Sources",
			ShortDescription: `Describe an Advanced Cluster.`,
			Description: `

` + "`" + `mongodbatlas_advanced_cluster` + "`" + ` describes an Advanced Cluster. The data source requires your Project ID.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

~> **IMPORTANT:**
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
					Name:        "name",
					Description: `(Required) Name of the cluster as it appears in Atlas. Once the cluster is created, its name cannot be changed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
				},
				resource.Attribute{
					Name:        "bi_connector_config",
					Description: `Configuration settings applied to BI Connector for Atlas on this cluster. See [below](#bi_connector_config).`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Type of the cluster that you want to create.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Capacity, in gigabytes, of the host's root volume.`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `Possible values are AWS, GCP, AZURE or NONE.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Configuration for the collection of key-value pairs that tag and categorize the cluster. See [below](#labels).`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `Version of the cluster to deploy.`,
				},
				resource.Attribute{
					Name:        "pit_enabled",
					Description: `Flag that indicates if the cluster uses Continuous Cloud Backup.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions and the hardware provisioned in them. See [below](#replication_specs)`,
				},
				resource.Attribute{
					Name:        "root_cert_type",
					Description: `Certificate Authority that MongoDB Atlas clusters use.`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.`,
				},
				resource.Attribute{
					Name:        "version_release_system",
					Description: `Release cadence that Atlas uses for this cluster.`,
				},
				resource.Attribute{
					Name:        "advanced_configuration",
					Description: `Get the advanced configuration options. See [Advanced Configuration](#advanced-configuration) below for more details. ### bi_connector_config Specifies BI Connector for Atlas configuration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Specifies whether or not BI Connector for Atlas is enabled on the cluster.l`,
				},
				resource.Attribute{
					Name:        "read_preference",
					Description: `Specifies the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences). ### labels Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. You cannot set the key ` + "`" + `Infrastructure Tool` + "`" + `, it is used for internal purposes to track aggregate usage.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. ### replication_specs`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Provide this value if you set a ` + "`" + `cluster_type` + "`" + ` of SHARDED or GEOSHARDED.`,
				},
				resource.Attribute{
					Name:        "region_configs",
					Description: `Configuration for the hardware specifications for nodes set for a given regionEach ` + "`" + `region_configs` + "`" + ` object describes the region's priority in elections and the number and type of MongoDB nodes that Atlas deploys to the region. Each ` + "`" + `region_configs` + "`" + ` object must have either an ` + "`" + `analytics_specs` + "`" + ` object, ` + "`" + `electable_specs` + "`" + ` object, or ` + "`" + `read_only_specs` + "`" + ` object. See [below](#region_configs)`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Name for the zone in a Global Cluster. ### region_configs`,
				},
				resource.Attribute{
					Name:        "analytics_specs",
					Description: `Hardware specifications for [analytics nodes](https://docs.atlas.mongodb.com/reference/faq/deployment/#std-label-analytics-nodes-overview) needed in the region. See [below](#specs)`,
				},
				resource.Attribute{
					Name:        "auto_scaling",
					Description: `Configuration for the Collection of settings that configures auto-scaling information for the cluster. See [below](#auto_scaling)`,
				},
				resource.Attribute{
					Name:        "analytics_auto_scaling",
					Description: `Configuration for the Collection of settings that configures analytics-auto-scaling information for the cluster. See [below](#analytics_auto_scaling)`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `Cloud service provider on which you provision the host for a multi-tenant cluster.`,
				},
				resource.Attribute{
					Name:        "electable_specs",
					Description: `Hardware specifications for electable nodes in the region.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Election priority of the region.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Cloud service provider on which the servers are provisioned.`,
				},
				resource.Attribute{
					Name:        "read_only_specs",
					Description: `Hardware specifications for read-only nodes in the region. See [below](#specs)`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Physical location of your MongoDB cluster. ### specs`,
				},
				resource.Attribute{
					Name:        "disk_iops",
					Description: `Target throughput (IOPS) desired for AWS storage attached to your cluster.`,
				},
				resource.Attribute{
					Name:        "ebs_volume_type",
					Description: `Type of storage you want to attach to your AWS-provisioned cluster.`,
				},
				resource.Attribute{
					Name:        "instance_size",
					Description: `Hardware specification for the instance sizes in this region.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of read-only nodes for Atlas to deploy to the region. ### auto_scaling`,
				},
				resource.Attribute{
					Name:        "disk_gb_enabled",
					Description: `Flag that indicates whether this cluster enables disk auto-scaling.`,
				},
				resource.Attribute{
					Name:        "compute_enabled",
					Description: `Flag that indicates whether instance size auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "compute_scale_down_enabled",
					Description: `Flag that indicates whether the instance size may scale down.`,
				},
				resource.Attribute{
					Name:        "compute_min_instance_size",
					Description: `Minimum instance size to which your cluster can automatically scale (such as M10).`,
				},
				resource.Attribute{
					Name:        "compute_max_instance_size",
					Description: `Maximum instance size to which your cluster can automatically scale (such as M40). ### analytics_auto_scaling`,
				},
				resource.Attribute{
					Name:        "disk_gb_enabled",
					Description: `Flag that indicates whether this cluster enables disk auto-scaling.`,
				},
				resource.Attribute{
					Name:        "compute_enabled",
					Description: `Flag that indicates whether instance size auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "compute_scale_down_enabled",
					Description: `Flag that indicates whether the instance size may scale down.`,
				},
				resource.Attribute{
					Name:        "compute_min_instance_size",
					Description: `Minimum instance size to which your cluster can automatically scale (such as M10).`,
				},
				resource.Attribute{
					Name:        "compute_max_instance_size",
					Description: `Maximum instance size to which your cluster can automatically scale (such as M40). #### Advanced Configuration`,
				},
				resource.Attribute{
					Name:        "default_read_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for read operations](https://docs.mongodb.com/manual/reference/read-concern/) set for this cluster. MongoDB 4.4 clusters default to [available](https://docs.mongodb.com/manual/reference/read-concern-available/).`,
				},
				resource.Attribute{
					Name:        "default_write_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for write operations](https://docs.mongodb.com/manual/reference/write-concern/) set for this cluster. MongoDB 4.4 clusters default to [1](https://docs.mongodb.com/manual/reference/write-concern/).`,
				},
				resource.Attribute{
					Name:        "fail_index_key_too_long",
					Description: `When true, documents can only be updated or inserted if, for all indexed fields on the target collection, the corresponding index entries do not exceed 1024 bytes. When false, mongod writes documents that exceed the limit but does not index them.`,
				},
				resource.Attribute{
					Name:        "javascript_enabled",
					Description: `When true, the cluster allows execution of operations that perform server-side executions of JavaScript. When false, the cluster disables execution of those operations.`,
				},
				resource.Attribute{
					Name:        "minimum_enabled_tls_protocol",
					Description: `Sets the minimum Transport Layer Security (TLS) version the cluster accepts for incoming connections.Valid values are: - TLS1_0 - TLS1_1 - TLS1_2`,
				},
				resource.Attribute{
					Name:        "no_table_scan",
					Description: `When true, the cluster disables the execution of any query that requires a collection scan to return results. When false, the cluster allows the execution of those operations.`,
				},
				resource.Attribute{
					Name:        "oplog_size_mb",
					Description: `The custom oplog size of the cluster. Without a value that indicates that the cluster uses the default oplog size calculated by Atlas.`,
				},
				resource.Attribute{
					Name:        "oplog_min_retention_hours",
					Description: `Minimum retention window for cluster's oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates.`,
				},
				resource.Attribute{
					Name:        "sample_size_bi_connector",
					Description: `Number of documents per database to sample when gathering schema information. Defaults to 100. Available only for Atlas deployments in which BI Connector for Atlas is enabled.`,
				},
				resource.Attribute{
					Name:        "sample_refresh_interval_bi_connector",
					Description: `Interval in seconds at which the mongosqld process re-samples data to create its relational schema. The default value is 300. The specified value must be a positive integer. Available only for Atlas deployments in which BI Connector for Atlas is enabled. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "paused",
					Description: `Flag that indicates whether the cluster is paused or not.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Current state of the cluster. The possible states are: See detailed information for arguments and attributes: [MongoDB API Advanced Cluster](https://docs.atlas.mongodb.com/reference/api/cluster-advanced/get-one-cluster-advanced/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
				},
				resource.Attribute{
					Name:        "bi_connector_config",
					Description: `Configuration settings applied to BI Connector for Atlas on this cluster. See [below](#bi_connector_config).`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Type of the cluster that you want to create.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Capacity, in gigabytes, of the host's root volume.`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `Possible values are AWS, GCP, AZURE or NONE.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Configuration for the collection of key-value pairs that tag and categorize the cluster. See [below](#labels).`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `Version of the cluster to deploy.`,
				},
				resource.Attribute{
					Name:        "pit_enabled",
					Description: `Flag that indicates if the cluster uses Continuous Cloud Backup.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions and the hardware provisioned in them. See [below](#replication_specs)`,
				},
				resource.Attribute{
					Name:        "root_cert_type",
					Description: `Certificate Authority that MongoDB Atlas clusters use.`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.`,
				},
				resource.Attribute{
					Name:        "version_release_system",
					Description: `Release cadence that Atlas uses for this cluster.`,
				},
				resource.Attribute{
					Name:        "advanced_configuration",
					Description: `Get the advanced configuration options. See [Advanced Configuration](#advanced-configuration) below for more details. ### bi_connector_config Specifies BI Connector for Atlas configuration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Specifies whether or not BI Connector for Atlas is enabled on the cluster.l`,
				},
				resource.Attribute{
					Name:        "read_preference",
					Description: `Specifies the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences). ### labels Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. You cannot set the key ` + "`" + `Infrastructure Tool` + "`" + `, it is used for internal purposes to track aggregate usage.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. ### replication_specs`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Provide this value if you set a ` + "`" + `cluster_type` + "`" + ` of SHARDED or GEOSHARDED.`,
				},
				resource.Attribute{
					Name:        "region_configs",
					Description: `Configuration for the hardware specifications for nodes set for a given regionEach ` + "`" + `region_configs` + "`" + ` object describes the region's priority in elections and the number and type of MongoDB nodes that Atlas deploys to the region. Each ` + "`" + `region_configs` + "`" + ` object must have either an ` + "`" + `analytics_specs` + "`" + ` object, ` + "`" + `electable_specs` + "`" + ` object, or ` + "`" + `read_only_specs` + "`" + ` object. See [below](#region_configs)`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Name for the zone in a Global Cluster. ### region_configs`,
				},
				resource.Attribute{
					Name:        "analytics_specs",
					Description: `Hardware specifications for [analytics nodes](https://docs.atlas.mongodb.com/reference/faq/deployment/#std-label-analytics-nodes-overview) needed in the region. See [below](#specs)`,
				},
				resource.Attribute{
					Name:        "auto_scaling",
					Description: `Configuration for the Collection of settings that configures auto-scaling information for the cluster. See [below](#auto_scaling)`,
				},
				resource.Attribute{
					Name:        "analytics_auto_scaling",
					Description: `Configuration for the Collection of settings that configures analytics-auto-scaling information for the cluster. See [below](#analytics_auto_scaling)`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `Cloud service provider on which you provision the host for a multi-tenant cluster.`,
				},
				resource.Attribute{
					Name:        "electable_specs",
					Description: `Hardware specifications for electable nodes in the region.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Election priority of the region.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Cloud service provider on which the servers are provisioned.`,
				},
				resource.Attribute{
					Name:        "read_only_specs",
					Description: `Hardware specifications for read-only nodes in the region. See [below](#specs)`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Physical location of your MongoDB cluster. ### specs`,
				},
				resource.Attribute{
					Name:        "disk_iops",
					Description: `Target throughput (IOPS) desired for AWS storage attached to your cluster.`,
				},
				resource.Attribute{
					Name:        "ebs_volume_type",
					Description: `Type of storage you want to attach to your AWS-provisioned cluster.`,
				},
				resource.Attribute{
					Name:        "instance_size",
					Description: `Hardware specification for the instance sizes in this region.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of read-only nodes for Atlas to deploy to the region. ### auto_scaling`,
				},
				resource.Attribute{
					Name:        "disk_gb_enabled",
					Description: `Flag that indicates whether this cluster enables disk auto-scaling.`,
				},
				resource.Attribute{
					Name:        "compute_enabled",
					Description: `Flag that indicates whether instance size auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "compute_scale_down_enabled",
					Description: `Flag that indicates whether the instance size may scale down.`,
				},
				resource.Attribute{
					Name:        "compute_min_instance_size",
					Description: `Minimum instance size to which your cluster can automatically scale (such as M10).`,
				},
				resource.Attribute{
					Name:        "compute_max_instance_size",
					Description: `Maximum instance size to which your cluster can automatically scale (such as M40). ### analytics_auto_scaling`,
				},
				resource.Attribute{
					Name:        "disk_gb_enabled",
					Description: `Flag that indicates whether this cluster enables disk auto-scaling.`,
				},
				resource.Attribute{
					Name:        "compute_enabled",
					Description: `Flag that indicates whether instance size auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "compute_scale_down_enabled",
					Description: `Flag that indicates whether the instance size may scale down.`,
				},
				resource.Attribute{
					Name:        "compute_min_instance_size",
					Description: `Minimum instance size to which your cluster can automatically scale (such as M10).`,
				},
				resource.Attribute{
					Name:        "compute_max_instance_size",
					Description: `Maximum instance size to which your cluster can automatically scale (such as M40). #### Advanced Configuration`,
				},
				resource.Attribute{
					Name:        "default_read_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for read operations](https://docs.mongodb.com/manual/reference/read-concern/) set for this cluster. MongoDB 4.4 clusters default to [available](https://docs.mongodb.com/manual/reference/read-concern-available/).`,
				},
				resource.Attribute{
					Name:        "default_write_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for write operations](https://docs.mongodb.com/manual/reference/write-concern/) set for this cluster. MongoDB 4.4 clusters default to [1](https://docs.mongodb.com/manual/reference/write-concern/).`,
				},
				resource.Attribute{
					Name:        "fail_index_key_too_long",
					Description: `When true, documents can only be updated or inserted if, for all indexed fields on the target collection, the corresponding index entries do not exceed 1024 bytes. When false, mongod writes documents that exceed the limit but does not index them.`,
				},
				resource.Attribute{
					Name:        "javascript_enabled",
					Description: `When true, the cluster allows execution of operations that perform server-side executions of JavaScript. When false, the cluster disables execution of those operations.`,
				},
				resource.Attribute{
					Name:        "minimum_enabled_tls_protocol",
					Description: `Sets the minimum Transport Layer Security (TLS) version the cluster accepts for incoming connections.Valid values are: - TLS1_0 - TLS1_1 - TLS1_2`,
				},
				resource.Attribute{
					Name:        "no_table_scan",
					Description: `When true, the cluster disables the execution of any query that requires a collection scan to return results. When false, the cluster allows the execution of those operations.`,
				},
				resource.Attribute{
					Name:        "oplog_size_mb",
					Description: `The custom oplog size of the cluster. Without a value that indicates that the cluster uses the default oplog size calculated by Atlas.`,
				},
				resource.Attribute{
					Name:        "oplog_min_retention_hours",
					Description: `Minimum retention window for cluster's oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates.`,
				},
				resource.Attribute{
					Name:        "sample_size_bi_connector",
					Description: `Number of documents per database to sample when gathering schema information. Defaults to 100. Available only for Atlas deployments in which BI Connector for Atlas is enabled.`,
				},
				resource.Attribute{
					Name:        "sample_refresh_interval_bi_connector",
					Description: `Interval in seconds at which the mongosqld process re-samples data to create its relational schema. The default value is 300. The specified value must be a positive integer. Available only for Atlas deployments in which BI Connector for Atlas is enabled. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "paused",
					Description: `Flag that indicates whether the cluster is paused or not.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Current state of the cluster. The possible states are: See detailed information for arguments and attributes: [MongoDB API Advanced Cluster](https://docs.atlas.mongodb.com/reference/api/cluster-advanced/get-one-cluster-advanced/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_advanced_clusters",
			Category:         "Data Sources",
			ShortDescription: `Describe all Advanced Clusters in Project.`,
			Description: `

` + "`" + `mongodbatlas_cluster` + "`" + ` describes all Advanced Clusters by the provided project_id. The data source requires your Project ID.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

~> **IMPORTANT:**
<br> &#8226; Changes to cluster configurations can affect costs. Before making changes, please see [Billing](https://docs.atlas.mongodb.com/billing/).
<br> &#8226; If your Atlas project contains a custom role that uses actions introduced in a specific MongoDB version, you cannot create a cluster with a MongoDB version less than that version unless you delete the custom role.

`,
			Keywords: []string{},
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
					Description: `A list where each represents a Cluster. See below for more details. ### Advanced Cluster`,
				},
				resource.Attribute{
					Name:        "bi_connector_config",
					Description: `Configuration settings applied to BI Connector for Atlas on this cluster. See [below](#bi_connector_config).`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Type of the cluster that you want to create.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Capacity, in gigabytes, of the host's root volume.`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `Possible values are AWS, GCP, AZURE or NONE.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Configuration for the collection of key-value pairs that tag and categorize the cluster. See [below](#labels).`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `Version of the cluster to deploy.`,
				},
				resource.Attribute{
					Name:        "pit_enabled",
					Description: `Flag that indicates if the cluster uses Continuous Cloud Backup.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions and the hardware provisioned in them. See [below](#replication_specs)`,
				},
				resource.Attribute{
					Name:        "root_cert_type",
					Description: `Certificate Authority that MongoDB Atlas clusters use.`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.`,
				},
				resource.Attribute{
					Name:        "version_release_system",
					Description: `Release cadence that Atlas uses for this cluster.`,
				},
				resource.Attribute{
					Name:        "advanced_configuration",
					Description: `Get the advanced configuration options. See [Advanced Configuration](#advanced-configuration) below for more details. ### bi_connector_config Specifies BI Connector for Atlas configuration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Specifies whether or not BI Connector for Atlas is enabled on the cluster.l`,
				},
				resource.Attribute{
					Name:        "read_preference",
					Description: `Specifies the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences). ### labels Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. You cannot set the key ` + "`" + `Infrastructure Tool` + "`" + `, it is used for internal purposes to track aggregate usage.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. ### replication_specs`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Provide this value if you set a ` + "`" + `cluster_type` + "`" + ` of SHARDED or GEOSHARDED.`,
				},
				resource.Attribute{
					Name:        "region_configs",
					Description: `Configuration for the hardware specifications for nodes set for a given regionEach ` + "`" + `region_configs` + "`" + ` object describes the region's priority in elections and the number and type of MongoDB nodes that Atlas deploys to the region. Each ` + "`" + `region_configs` + "`" + ` object must have either an ` + "`" + `analytics_specs` + "`" + ` object, ` + "`" + `electable_specs` + "`" + ` object, or ` + "`" + `read_only_specs` + "`" + ` object. See [below](#region_configs)`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Name for the zone in a Global Cluster. ### region_configs`,
				},
				resource.Attribute{
					Name:        "analytics_specs",
					Description: `Hardware specifications for [analytics nodes](https://docs.atlas.mongodb.com/reference/faq/deployment/#std-label-analytics-nodes-overview) needed in the region. See [below](#specs)`,
				},
				resource.Attribute{
					Name:        "auto_scaling",
					Description: `Configuration for the Collection of settings that configures auto-scaling information for the cluster. See [below](#auto_scaling)`,
				},
				resource.Attribute{
					Name:        "analytics_auto_scaling",
					Description: `Configuration for the Collection of settings that configures analytis-auto-scaling information for the cluster. See [below](#analytics_auto_scaling)`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `Cloud service provider on which you provision the host for a multi-tenant cluster.`,
				},
				resource.Attribute{
					Name:        "electable_specs",
					Description: `Hardware specifications for electable nodes in the region.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Election priority of the region.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Cloud service provider on which the servers are provisioned.`,
				},
				resource.Attribute{
					Name:        "read_only_specs",
					Description: `Hardware specifications for read-only nodes in the region. See [below](#specs)`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Physical location of your MongoDB cluster. ### specs`,
				},
				resource.Attribute{
					Name:        "disk_iops",
					Description: `Target throughput (IOPS) desired for AWS storage attached to your cluster.`,
				},
				resource.Attribute{
					Name:        "ebs_volume_type",
					Description: `Type of storage you want to attach to your AWS-provisioned cluster.`,
				},
				resource.Attribute{
					Name:        "instance_size",
					Description: `Hardware specification for the instance sizes in this region.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of read-only nodes for Atlas to deploy to the region. ### auto_scaling`,
				},
				resource.Attribute{
					Name:        "disk_gb_enabled",
					Description: `Flag that indicates whether this cluster enables disk auto-scaling.`,
				},
				resource.Attribute{
					Name:        "compute_enabled",
					Description: `Flag that indicates whether instance size auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "compute_scale_down_enabled",
					Description: `Flag that indicates whether the instance size may scale down.`,
				},
				resource.Attribute{
					Name:        "compute_min_instance_size",
					Description: `Minimum instance size to which your cluster can automatically scale (such as M10).`,
				},
				resource.Attribute{
					Name:        "compute_max_instance_size",
					Description: `Maximum instance size to which your cluster can automatically scale (such as M40). ### analytics_auto_scaling`,
				},
				resource.Attribute{
					Name:        "disk_gb_enabled",
					Description: `Flag that indicates whether this cluster enables disk auto-scaling.`,
				},
				resource.Attribute{
					Name:        "compute_enabled",
					Description: `Flag that indicates whether instance size auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "compute_scale_down_enabled",
					Description: `Flag that indicates whether the instance size may scale down.`,
				},
				resource.Attribute{
					Name:        "compute_min_instance_size",
					Description: `Minimum instance size to which your cluster can automatically scale (such as M10).`,
				},
				resource.Attribute{
					Name:        "compute_max_instance_size",
					Description: `Maximum instance size to which your cluster can automatically scale (such as M40). #### Advanced Configuration`,
				},
				resource.Attribute{
					Name:        "default_read_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for read operations](https://docs.mongodb.com/manual/reference/read-concern/) set for this cluster. MongoDB 4.4 clusters default to [available](https://docs.mongodb.com/manual/reference/read-concern-available/).`,
				},
				resource.Attribute{
					Name:        "default_write_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for write operations](https://docs.mongodb.com/manual/reference/write-concern/) set for this cluster. MongoDB 4.4 clusters default to [1](https://docs.mongodb.com/manual/reference/write-concern/).`,
				},
				resource.Attribute{
					Name:        "fail_index_key_too_long",
					Description: `When true, documents can only be updated or inserted if, for all indexed fields on the target collection, the corresponding index entries do not exceed 1024 bytes. When false, mongod writes documents that exceed the limit but does not index them.`,
				},
				resource.Attribute{
					Name:        "javascript_enabled",
					Description: `When true, the cluster allows execution of operations that perform server-side executions of JavaScript. When false, the cluster disables execution of those operations.`,
				},
				resource.Attribute{
					Name:        "minimum_enabled_tls_protocol",
					Description: `Sets the minimum Transport Layer Security (TLS) version the cluster accepts for incoming connections.Valid values are: - TLS1_0 - TLS1_1 - TLS1_2`,
				},
				resource.Attribute{
					Name:        "no_table_scan",
					Description: `When true, the cluster disables the execution of any query that requires a collection scan to return results. When false, the cluster allows the execution of those operations.`,
				},
				resource.Attribute{
					Name:        "oplog_size_mb",
					Description: `The custom oplog size of the cluster. Without a value that indicates that the cluster uses the default oplog size calculated by Atlas.`,
				},
				resource.Attribute{
					Name:        "oplog_min_retention_hours",
					Description: `Minimum retention window for cluster's oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates.`,
				},
				resource.Attribute{
					Name:        "sample_size_bi_connector",
					Description: `Number of documents per database to sample when gathering schema information. Defaults to 100. Available only for Atlas deployments in which BI Connector for Atlas is enabled.`,
				},
				resource.Attribute{
					Name:        "sample_refresh_interval_bi_connector",
					Description: `Interval in seconds at which the mongosqld process re-samples data to create its relational schema. The default value is 300. The specified value must be a positive integer. Available only for Atlas deployments in which BI Connector for Atlas is enabled. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "paused",
					Description: `Flag that indicates whether the cluster is paused or not.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Current state of the cluster. The possible states are: See detailed information for arguments and attributes: [MongoDB API Advanced Clusters](https://docs.atlas.mongodb.com/reference/api/cluster-advanced/get-all-cluster-advanced/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Cluster. See below for more details. ### Advanced Cluster`,
				},
				resource.Attribute{
					Name:        "bi_connector_config",
					Description: `Configuration settings applied to BI Connector for Atlas on this cluster. See [below](#bi_connector_config).`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Type of the cluster that you want to create.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Capacity, in gigabytes, of the host's root volume.`,
				},
				resource.Attribute{
					Name:        "encryption_at_rest_provider",
					Description: `Possible values are AWS, GCP, AZURE or NONE.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Configuration for the collection of key-value pairs that tag and categorize the cluster. See [below](#labels).`,
				},
				resource.Attribute{
					Name:        "mongo_db_major_version",
					Description: `Version of the cluster to deploy.`,
				},
				resource.Attribute{
					Name:        "pit_enabled",
					Description: `Flag that indicates if the cluster uses Continuous Cloud Backup.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions and the hardware provisioned in them. See [below](#replication_specs)`,
				},
				resource.Attribute{
					Name:        "root_cert_type",
					Description: `Certificate Authority that MongoDB Atlas clusters use.`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.`,
				},
				resource.Attribute{
					Name:        "version_release_system",
					Description: `Release cadence that Atlas uses for this cluster.`,
				},
				resource.Attribute{
					Name:        "advanced_configuration",
					Description: `Get the advanced configuration options. See [Advanced Configuration](#advanced-configuration) below for more details. ### bi_connector_config Specifies BI Connector for Atlas configuration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Specifies whether or not BI Connector for Atlas is enabled on the cluster.l`,
				},
				resource.Attribute{
					Name:        "read_preference",
					Description: `Specifies the read preference to be used by BI Connector for Atlas on the cluster. Each BI Connector for Atlas read preference contains a distinct combination of [readPreference](https://docs.mongodb.com/manual/core/read-preference/) and [readPreferenceTags](https://docs.mongodb.com/manual/core/read-preference/#tag-sets) options. For details on BI Connector for Atlas read preferences, refer to the [BI Connector Read Preferences Table](https://docs.atlas.mongodb.com/tutorial/create-global-writes-cluster/#bic-read-preferences). ### labels Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. You cannot set the key ` + "`" + `Infrastructure Tool` + "`" + `, it is used for internal purposes to track aggregate usage.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. ### replication_specs`,
				},
				resource.Attribute{
					Name:        "num_shards",
					Description: `Provide this value if you set a ` + "`" + `cluster_type` + "`" + ` of SHARDED or GEOSHARDED.`,
				},
				resource.Attribute{
					Name:        "region_configs",
					Description: `Configuration for the hardware specifications for nodes set for a given regionEach ` + "`" + `region_configs` + "`" + ` object describes the region's priority in elections and the number and type of MongoDB nodes that Atlas deploys to the region. Each ` + "`" + `region_configs` + "`" + ` object must have either an ` + "`" + `analytics_specs` + "`" + ` object, ` + "`" + `electable_specs` + "`" + ` object, or ` + "`" + `read_only_specs` + "`" + ` object. See [below](#region_configs)`,
				},
				resource.Attribute{
					Name:        "zone_name",
					Description: `Name for the zone in a Global Cluster. ### region_configs`,
				},
				resource.Attribute{
					Name:        "analytics_specs",
					Description: `Hardware specifications for [analytics nodes](https://docs.atlas.mongodb.com/reference/faq/deployment/#std-label-analytics-nodes-overview) needed in the region. See [below](#specs)`,
				},
				resource.Attribute{
					Name:        "auto_scaling",
					Description: `Configuration for the Collection of settings that configures auto-scaling information for the cluster. See [below](#auto_scaling)`,
				},
				resource.Attribute{
					Name:        "analytics_auto_scaling",
					Description: `Configuration for the Collection of settings that configures analytis-auto-scaling information for the cluster. See [below](#analytics_auto_scaling)`,
				},
				resource.Attribute{
					Name:        "backing_provider_name",
					Description: `Cloud service provider on which you provision the host for a multi-tenant cluster.`,
				},
				resource.Attribute{
					Name:        "electable_specs",
					Description: `Hardware specifications for electable nodes in the region.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Election priority of the region.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Cloud service provider on which the servers are provisioned.`,
				},
				resource.Attribute{
					Name:        "read_only_specs",
					Description: `Hardware specifications for read-only nodes in the region. See [below](#specs)`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Physical location of your MongoDB cluster. ### specs`,
				},
				resource.Attribute{
					Name:        "disk_iops",
					Description: `Target throughput (IOPS) desired for AWS storage attached to your cluster.`,
				},
				resource.Attribute{
					Name:        "ebs_volume_type",
					Description: `Type of storage you want to attach to your AWS-provisioned cluster.`,
				},
				resource.Attribute{
					Name:        "instance_size",
					Description: `Hardware specification for the instance sizes in this region.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of read-only nodes for Atlas to deploy to the region. ### auto_scaling`,
				},
				resource.Attribute{
					Name:        "disk_gb_enabled",
					Description: `Flag that indicates whether this cluster enables disk auto-scaling.`,
				},
				resource.Attribute{
					Name:        "compute_enabled",
					Description: `Flag that indicates whether instance size auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "compute_scale_down_enabled",
					Description: `Flag that indicates whether the instance size may scale down.`,
				},
				resource.Attribute{
					Name:        "compute_min_instance_size",
					Description: `Minimum instance size to which your cluster can automatically scale (such as M10).`,
				},
				resource.Attribute{
					Name:        "compute_max_instance_size",
					Description: `Maximum instance size to which your cluster can automatically scale (such as M40). ### analytics_auto_scaling`,
				},
				resource.Attribute{
					Name:        "disk_gb_enabled",
					Description: `Flag that indicates whether this cluster enables disk auto-scaling.`,
				},
				resource.Attribute{
					Name:        "compute_enabled",
					Description: `Flag that indicates whether instance size auto-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "compute_scale_down_enabled",
					Description: `Flag that indicates whether the instance size may scale down.`,
				},
				resource.Attribute{
					Name:        "compute_min_instance_size",
					Description: `Minimum instance size to which your cluster can automatically scale (such as M10).`,
				},
				resource.Attribute{
					Name:        "compute_max_instance_size",
					Description: `Maximum instance size to which your cluster can automatically scale (such as M40). #### Advanced Configuration`,
				},
				resource.Attribute{
					Name:        "default_read_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for read operations](https://docs.mongodb.com/manual/reference/read-concern/) set for this cluster. MongoDB 4.4 clusters default to [available](https://docs.mongodb.com/manual/reference/read-concern-available/).`,
				},
				resource.Attribute{
					Name:        "default_write_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for write operations](https://docs.mongodb.com/manual/reference/write-concern/) set for this cluster. MongoDB 4.4 clusters default to [1](https://docs.mongodb.com/manual/reference/write-concern/).`,
				},
				resource.Attribute{
					Name:        "fail_index_key_too_long",
					Description: `When true, documents can only be updated or inserted if, for all indexed fields on the target collection, the corresponding index entries do not exceed 1024 bytes. When false, mongod writes documents that exceed the limit but does not index them.`,
				},
				resource.Attribute{
					Name:        "javascript_enabled",
					Description: `When true, the cluster allows execution of operations that perform server-side executions of JavaScript. When false, the cluster disables execution of those operations.`,
				},
				resource.Attribute{
					Name:        "minimum_enabled_tls_protocol",
					Description: `Sets the minimum Transport Layer Security (TLS) version the cluster accepts for incoming connections.Valid values are: - TLS1_0 - TLS1_1 - TLS1_2`,
				},
				resource.Attribute{
					Name:        "no_table_scan",
					Description: `When true, the cluster disables the execution of any query that requires a collection scan to return results. When false, the cluster allows the execution of those operations.`,
				},
				resource.Attribute{
					Name:        "oplog_size_mb",
					Description: `The custom oplog size of the cluster. Without a value that indicates that the cluster uses the default oplog size calculated by Atlas.`,
				},
				resource.Attribute{
					Name:        "oplog_min_retention_hours",
					Description: `Minimum retention window for cluster's oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates.`,
				},
				resource.Attribute{
					Name:        "sample_size_bi_connector",
					Description: `Number of documents per database to sample when gathering schema information. Defaults to 100. Available only for Atlas deployments in which BI Connector for Atlas is enabled.`,
				},
				resource.Attribute{
					Name:        "sample_refresh_interval_bi_connector",
					Description: `Interval in seconds at which the mongosqld process re-samples data to create its relational schema. The default value is 300. The specified value must be a positive integer. Available only for Atlas deployments in which BI Connector for Atlas is enabled. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "paused",
					Description: `Flag that indicates whether the cluster is paused or not.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Current state of the cluster. The possible states are: See detailed information for arguments and attributes: [MongoDB API Advanced Clusters](https://docs.atlas.mongodb.com/reference/api/cluster-advanced/get-all-cluster-advanced/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_alert_configuration",
			Category:         "Data Sources",
			ShortDescription: `Describes a Alert Configuration.`,
			Description: `

` + "`" + `mongodbatlas_alert_configuration` + "`" + ` describes an Alert Configuration.

-> **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project where the alert configuration will create.`,
				},
				resource.Attribute{
					Name:        "alert_configuration_id",
					Description: `(Required) Unique identifier for the alert configuration.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `(Optional) List of formatted output requested for this alert configuration`,
				},
				resource.Attribute{
					Name:        "output.#.type",
					Description: `(Required) If the output is requested, you must specify its type. The format is computed as ` + "`" + `output.#.value` + "`" + `, the following are the supported types: - ` + "`" + `resource_hcl` + "`" + `: This string is used to define the resource as it exists in MongoDB Atlas. - ` + "`" + `resource_import` + "`" + `: This string is used to import the existing resource into the state file. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Timestamp in ISO 8601 date and time format in UTC when this alert configuration was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Timestamp in ISO 8601 date and time format in UTC when this alert configuration was last updated.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If set to true, the alert configuration is enabled. If enabled is not exported it is set to false.`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `The type of event that will trigger an alert. ->`,
				},
				resource.Attribute{
					Name:        "field_name",
					Description: `Name of the field in the target object to match on. | Host alerts | Replica set alerts | Sharded cluster alerts | |:-------------- |:------------- |:------ | | ` + "`" + `TYPE_NAME` + "`" + ` | ` + "`" + `REPLICA_SET_NAME` + "`" + ` | ` + "`" + `CLUSTER_NAME` + "`" + ` | | ` + "`" + `HOSTNAME` + "`" + ` | ` + "`" + `SHARD_NAME` + "`" + ` | ` + "`" + `SHARD_NAME` + "`" + ` | | ` + "`" + `PORT` + "`" + ` | ` + "`" + `CLUSTER_NAME` + "`" + ` | | | ` + "`" + `HOSTNAME_AND_PORT` + "`" + ` | | | | ` + "`" + `REPLICA_SET_NAME` + "`" + ` | | | All other types of alerts do not support matchers.`,
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
					Description: `Value to test with the specified operator. If ` + "`" + `field_name` + "`" + ` is set to TYPE_NAME, you can match on the following values: - ` + "`" + `PRIMARY` + "`" + ` - ` + "`" + `SECONDARY` + "`" + ` - ` + "`" + `STANDALONE` + "`" + ` - ` + "`" + `CONFIG` + "`" + ` - ` + "`" + `MONGOS` + "`" + ` ### Metric Threshold Config (` + "`" + `metric_threshold_config` + "`" + `) The threshold that causes an alert to be triggered. Required if ` + "`" + `event_type_name` + "`" + ` : ` + "`" + `OUTSIDE_METRIC_THRESHOLD` + "`" + ` or ` + "`" + `OUTSIDE_SERVERLESS_METRIC_THRESHOLD` + "`" + `.`,
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
					Description: `The units for the threshold value. Depends on the type of metric. Refer to the [MongoDB API Alert Configuration documentation](https://www.mongodb.com/docs/atlas/reference/api/alert-configurations-get-config/#request-body-parameters) for a list of accepted values. ### Notifications Notifications to send when an alert condition is detected.`,
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
					Description: `Flag indicating email notifications should be sent. Atlas returns this value if ` + "`" + `type_name` + "`" + ` is set to ` + "`" + `ORG` + "`" + `, ` + "`" + `GROUP` + "`" + `, or ` + "`" + `USER` + "`" + `.`,
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
					Description: `Flag indicating text notifications should be sent. Atlas returns this value if ` + "`" + `type_name` + "`" + ` is set to ` + "`" + `ORG` + "`" + `, ` + "`" + `GROUP` + "`" + `, or ` + "`" + `USER` + "`" + `.`,
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
					Description: `Authentication secret for the ` + "`" + `WEBHOOK` + "`" + ` notifications type.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `Target URL for the ` + "`" + `WEBHOOK` + "`" + ` notifications type.`,
				},
				resource.Attribute{
					Name:        "microsoft_teams_webhook_url",
					Description: `Microsoft Teams channel incoming webhook URL. Required for the ` + "`" + `MICROSOFT_TEAMS` + "`" + ` notifications type.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Atlas role in current Project or Organization. Atlas returns this value if you set ` + "`" + `type_name` + "`" + ` to ` + "`" + `ORG` + "`" + ` or ` + "`" + `GROUP` + "`" + `. See detailed information for arguments and attributes: [MongoDB API Alert Configuration](https://docs.atlas.mongodb.com/reference/api/alert-configurations-get-config/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `Timestamp in ISO 8601 date and time format in UTC when this alert configuration was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Timestamp in ISO 8601 date and time format in UTC when this alert configuration was last updated.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If set to true, the alert configuration is enabled. If enabled is not exported it is set to false.`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `The type of event that will trigger an alert. ->`,
				},
				resource.Attribute{
					Name:        "field_name",
					Description: `Name of the field in the target object to match on. | Host alerts | Replica set alerts | Sharded cluster alerts | |:-------------- |:------------- |:------ | | ` + "`" + `TYPE_NAME` + "`" + ` | ` + "`" + `REPLICA_SET_NAME` + "`" + ` | ` + "`" + `CLUSTER_NAME` + "`" + ` | | ` + "`" + `HOSTNAME` + "`" + ` | ` + "`" + `SHARD_NAME` + "`" + ` | ` + "`" + `SHARD_NAME` + "`" + ` | | ` + "`" + `PORT` + "`" + ` | ` + "`" + `CLUSTER_NAME` + "`" + ` | | | ` + "`" + `HOSTNAME_AND_PORT` + "`" + ` | | | | ` + "`" + `REPLICA_SET_NAME` + "`" + ` | | | All other types of alerts do not support matchers.`,
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
					Description: `Value to test with the specified operator. If ` + "`" + `field_name` + "`" + ` is set to TYPE_NAME, you can match on the following values: - ` + "`" + `PRIMARY` + "`" + ` - ` + "`" + `SECONDARY` + "`" + ` - ` + "`" + `STANDALONE` + "`" + ` - ` + "`" + `CONFIG` + "`" + ` - ` + "`" + `MONGOS` + "`" + ` ### Metric Threshold Config (` + "`" + `metric_threshold_config` + "`" + `) The threshold that causes an alert to be triggered. Required if ` + "`" + `event_type_name` + "`" + ` : ` + "`" + `OUTSIDE_METRIC_THRESHOLD` + "`" + ` or ` + "`" + `OUTSIDE_SERVERLESS_METRIC_THRESHOLD` + "`" + `.`,
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
					Description: `The units for the threshold value. Depends on the type of metric. Refer to the [MongoDB API Alert Configuration documentation](https://www.mongodb.com/docs/atlas/reference/api/alert-configurations-get-config/#request-body-parameters) for a list of accepted values. ### Notifications Notifications to send when an alert condition is detected.`,
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
					Description: `Flag indicating email notifications should be sent. Atlas returns this value if ` + "`" + `type_name` + "`" + ` is set to ` + "`" + `ORG` + "`" + `, ` + "`" + `GROUP` + "`" + `, or ` + "`" + `USER` + "`" + `.`,
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
					Description: `Flag indicating text notifications should be sent. Atlas returns this value if ` + "`" + `type_name` + "`" + ` is set to ` + "`" + `ORG` + "`" + `, ` + "`" + `GROUP` + "`" + `, or ` + "`" + `USER` + "`" + `.`,
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
					Description: `Authentication secret for the ` + "`" + `WEBHOOK` + "`" + ` notifications type.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `Target URL for the ` + "`" + `WEBHOOK` + "`" + ` notifications type.`,
				},
				resource.Attribute{
					Name:        "microsoft_teams_webhook_url",
					Description: `Microsoft Teams channel incoming webhook URL. Required for the ` + "`" + `MICROSOFT_TEAMS` + "`" + ` notifications type.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Atlas role in current Project or Organization. Atlas returns this value if you set ` + "`" + `type_name` + "`" + ` to ` + "`" + `ORG` + "`" + ` or ` + "`" + `GROUP` + "`" + `. See detailed information for arguments and attributes: [MongoDB API Alert Configuration](https://docs.atlas.mongodb.com/reference/api/alert-configurations-get-config/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_alert_configurations",
			Category:         "Data Sources",
			ShortDescription: `Describe all Alert Configurations in Project.`,
			Description: `

` + "`" + `mongodbatlas_alert_configurations` + "`" + ` describes all Alert Configurations by the provided project_id. The data source requires your Project ID.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to get the alert configurations.`,
				},
				resource.Attribute{
					Name:        "list_options",
					Description: `(Optional) Arguments that dictate how many and which results are returned by the data source`,
				},
				resource.Attribute{
					Name:        "list_options.page_num",
					Description: `Which page of results to retrieve (default to first page)`,
				},
				resource.Attribute{
					Name:        "list_options.items_per_page",
					Description: `How many alerts to retrieve per page (default 100)`,
				},
				resource.Attribute{
					Name:        "list_options.include_count",
					Description: `Whether to include total count of results in the response (default false)`,
				},
				resource.Attribute{
					Name:        "output_type",
					Description: `(Optional) List of requested string formatted output to be included on each individual result. Options are ` + "`" + `resource_hcl` + "`" + ` and ` + "`" + `resource_import` + "`" + `. Available to make it easy to gather resource statements for existing alert configurations, and corresponding import statements to import said resource state into the statefile. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Total count of results`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list of alert configurations for the project_id, constrained by the ` + "`" + `list_options` + "`" + `. ### Alert Configuration`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project where the alert configuration exists`,
				},
				resource.Attribute{
					Name:        "alert_configuration_id",
					Description: `The ID of the alert configuration`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Timestamp in ISO 8601 date and time format in UTC when this alert configuration was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Timestamp in ISO 8601 date and time format in UTC when this alert configuration was last updated.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If set to true, the alert configuration is enabled. If enabled is not exported it is set to false.`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `The type of event that will trigger an alert.`,
				},
				resource.Attribute{
					Name:        "matcher",
					Description: `Rules to apply when matching an object against this alert configuration`,
				},
				resource.Attribute{
					Name:        "metric_threshold_config",
					Description: `The threshold that causes an alert to be triggered. Required if ` + "`" + `event_type_name` + "`" + ` : ` + "`" + `OUTSIDE_METRIC_THRESHOLD` + "`" + ` or ` + "`" + `OUTSIDE_SERVERLESS_METRIC_THRESHOLD` + "`" + ``,
				},
				resource.Attribute{
					Name:        "threshold_config",
					Description: `Threshold that triggers an alert. Required if ` + "`" + `event_type_name` + "`" + ` is any value other than ` + "`" + `OUTSIDE_METRIC_THRESHOLD` + "`" + ` or ` + "`" + `OUTSIDE_SERVERLESS_METRIC_THRESHOLD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `List of notifications to send when an alert condition is detected.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `Requested output string format for the alert configuration For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/alert-configurations/) Or refer to the individual resource or data_source documentation on alert configuration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total_count",
					Description: `Total count of results`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list of alert configurations for the project_id, constrained by the ` + "`" + `list_options` + "`" + `. ### Alert Configuration`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project where the alert configuration exists`,
				},
				resource.Attribute{
					Name:        "alert_configuration_id",
					Description: `The ID of the alert configuration`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Timestamp in ISO 8601 date and time format in UTC when this alert configuration was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Timestamp in ISO 8601 date and time format in UTC when this alert configuration was last updated.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If set to true, the alert configuration is enabled. If enabled is not exported it is set to false.`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `The type of event that will trigger an alert.`,
				},
				resource.Attribute{
					Name:        "matcher",
					Description: `Rules to apply when matching an object against this alert configuration`,
				},
				resource.Attribute{
					Name:        "metric_threshold_config",
					Description: `The threshold that causes an alert to be triggered. Required if ` + "`" + `event_type_name` + "`" + ` : ` + "`" + `OUTSIDE_METRIC_THRESHOLD` + "`" + ` or ` + "`" + `OUTSIDE_SERVERLESS_METRIC_THRESHOLD` + "`" + ``,
				},
				resource.Attribute{
					Name:        "threshold_config",
					Description: `Threshold that triggers an alert. Required if ` + "`" + `event_type_name` + "`" + ` is any value other than ` + "`" + `OUTSIDE_METRIC_THRESHOLD` + "`" + ` or ` + "`" + `OUTSIDE_SERVERLESS_METRIC_THRESHOLD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `List of notifications to send when an alert condition is detected.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `Requested output string format for the alert configuration For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/alert-configurations/) Or refer to the individual resource or data_source documentation on alert configuration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_api_key",
			Category:         "Data Sources",
			ShortDescription: `Describes a API Key.`,
			Description: `

` + "`" + `mongodbatlas_api_key` + "`" + ` describes a MongoDB Atlas API Key. This represents a API Key that has been created.

~> **IMPORTANT WARNING:** Managing Atlas Programmatic API Keys (PAKs) with Terraform will expose sensitive organizational secrets in Terraform's state. We suggest following [Terraform's best practices](https://developer.hashicorp.com/terraform/language/state/sensitive-data). You may also want to consider managing your PAKs via a more secure method, such as the [HashiCorp Vault MongoDB Atlas Secrets Engine](https://developer.hashicorp.com/vault/docs/secrets/mongodbatlas).

-> **NOTE:** You may find org_id in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) The unique ID for the project. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `Unique identifier for the organization whose API keys you want to retrieve. Use the /orgs endpoint to retrieve all organizations to which the authenticated user has access.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this Organization API key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `Public key for this Organization API key.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `Private key for this Organization API key.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `Name of the role. This resource returns all the roles the user has in Atlas. The following are valid roles:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `Unique identifier for the organization whose API keys you want to retrieve. Use the /orgs endpoint to retrieve all organizations to which the authenticated user has access.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this Organization API key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `Public key for this Organization API key.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `Private key for this Organization API key.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `Name of the role. This resource returns all the roles the user has in Atlas. The following are valid roles:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_api_keys",
			Category:         "Data Sources",
			ShortDescription: `Describes a API Keys.`,
			Description: `

` + "`" + `mongodbatlas_api_keys` + "`" + ` describe all API Keys. This represents API Keys that have been created.

~> **IMPORTANT WARNING:** Managing Atlas Programmatic API Keys (PAKs) with Terraform will expose sensitive organizational secrets in Terraform's state. We suggest following [Terraform's best practices](https://developer.hashicorp.com/terraform/language/state/sensitive-data). You may also want to consider managing your PAKs via a more secure method, such as the [HashiCorp Vault MongoDB Atlas Secrets Engine](https://developer.hashicorp.com/vault/docs/secrets/mongodbatlas).

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a List of API Keys. ### API Keys`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `Unique identifier for the organization whose API keys you want to retrieve. Use the /orgs endpoint to retrieve all organizations to which the authenticated user has access.`,
				},
				resource.Attribute{
					Name:        "api_key_id",
					Description: `Unique identifier for the API key you want to update. Use the /orgs/{ORG-ID}/apiKeys endpoint to retrieve all API keys to which the authenticated user has access for the specified organization.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this Organization API key.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `Name of the role. This resource returns all the roles the user has in Atlas. The following are valid roles:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_auditing",
			Category:         "Data Sources",
			ShortDescription: `Describes a Auditing.`,
			Description: `

` + "`" + `mongodbatlas_auditing` + "`" + ` describes a Auditing.

-> **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "configuration_type",
					Description: `Denotes the configuration method for the audit filter. Possible values are: NONE - auditing not configured for the project.m FILTER_BUILDER - auditing configured via Atlas UI filter builderm FILTER_JSON - auditing configured via Atlas custom filter or API.`,
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
					Description: `Denotes whether or not the project associated with the {GROUP-ID} has database auditing enabled. See detailed information for arguments and attributes: [MongoDB API Auditing](https://docs.atlas.mongodb.com/reference/api/auditing/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "configuration_type",
					Description: `Denotes the configuration method for the audit filter. Possible values are: NONE - auditing not configured for the project.m FILTER_BUILDER - auditing configured via Atlas UI filter builderm FILTER_JSON - auditing configured via Atlas custom filter or API.`,
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
					Description: `Denotes whether or not the project associated with the {GROUP-ID} has database auditing enabled. See detailed information for arguments and attributes: [MongoDB API Auditing](https://docs.atlas.mongodb.com/reference/api/auditing/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_backup_compliance_policy",
			Category:         "Data Sources",
			ShortDescription: `Provides a Backup Compliance Policy Datasource.`,
			Description: `

` + "`" + `mongodbatlas_backup_compliance_policy` + "`" + ` provides an Atlas Backup Compliance Policy. An Atlas Backup Compliance Policy contains the current protection policy settings for a project. A compliance policy prevents any user, regardless of role, from modifying or deleting specific cluster configurations and backups. To disable a Backup Compliance Policy, you must contact MongoDB support. Backup Compliance Policies are only supported for clusters M10 and higher and are applied as the minimum policy for all clusters.

-> **IMPORTANT NOTE:** Once you enable a Backup Compliance Policy, no user, regardless of role, can disable the Backup Compliance Policy via Terraform, or any other method, without contacting MongoDB support. This means that, once enabled, some resources defined in Terraform can not be modified. To learn more, see the full list of [Backup Compliance Policy Prohibited Actions and Considerations](https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/#configure-a-backup-compliance-policy).

-> **NOTE:** Groups and projects are synonymous terms. You might find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies your project ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "authorized_email",
					Description: `Email address of the user who is authorized to update the Backup Compliance Policy settings.`,
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
					Name:        "reference_minute_of_hour",
					Description: `Integer between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.`,
				},
				resource.Attribute{
					Name:        "restore_window_days",
					Description: `Number of previous days that you can restore back to with Continuous Cloud Backup with a Backup Compliance Policy. You must specify a positive, non-zero integer, and the maximum retention window can't exceed the hourly retention time. This parameter applies only to Continuous Cloud Backups with a Backup Compliance Policy.`,
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
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `.`,
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
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. Monthly policy must have retention days of at least 31 days or 5 weeks or 1 month. Note that for less frequent policy items, Atlas requires that you specify a retention period greater than or equal to the retention period specified for more frequent policy items. For example: If the weekly policy item specifies a retention of two weeks, the montly retention policy must specify two weeks or greater. For more information, see [MongoDB Atlas API Reference](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Cloud-Backups/operation/getDataProtectionSettings) and [Backup Compliance Policy Prohibited Actions](https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/#prohibited-actions)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "authorized_email",
					Description: `Email address of the user who is authorized to update the Backup Compliance Policy settings.`,
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
					Name:        "reference_minute_of_hour",
					Description: `Integer between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.`,
				},
				resource.Attribute{
					Name:        "restore_window_days",
					Description: `Number of previous days that you can restore back to with Continuous Cloud Backup with a Backup Compliance Policy. You must specify a positive, non-zero integer, and the maximum retention window can't exceed the hourly retention time. This parameter applies only to Continuous Cloud Backups with a Backup Compliance Policy.`,
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
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `.`,
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
					Description: `Value to associate with ` + "`" + `retention_unit` + "`" + `. Monthly policy must have retention days of at least 31 days or 5 weeks or 1 month. Note that for less frequent policy items, Atlas requires that you specify a retention period greater than or equal to the retention period specified for more frequent policy items. For example: If the weekly policy item specifies a retention of two weeks, the montly retention policy must specify two weeks or greater. For more information, see [MongoDB Atlas API Reference](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Cloud-Backups/operation/getDataProtectionSettings) and [Backup Compliance Policy Prohibited Actions](https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/#prohibited-actions)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_backup_schedule",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Backup Schedule Datasource.`,
			Description: `

` + "`" + `mongodbatlas_cloud_backup_schedule` + "`" + ` provides a Cloud Backup Schedule datasource. An Atlas Cloud Backup Schedule provides the current cloud backup schedule for the cluster. 

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
					Description: `(Required) The name of the Atlas cluster that contains the snapshots backup policy you want to retrieve. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Unique identifier of the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "next_snapshot",
					Description: `UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.`,
				},
				resource.Attribute{
					Name:        "reference_hour_of_day",
					Description: `UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.`,
				},
				resource.Attribute{
					Name:        "reference_minute_of_hour",
					Description: `UTC Minute of day between 0 and 59 representing which minute of the ` + "`" + `reference_hour_of_day` + "`" + ` that Atlas takes the snapshot.`,
				},
				resource.Attribute{
					Name:        "restore_window_days",
					Description: `Specifies a restore window in days for cloud backup to maintain.`,
				},
				resource.Attribute{
					Name:        "id_policy",
					Description: `Unique identifier of the backup policy.`,
				},
				resource.Attribute{
					Name:        "policy_item_hourly",
					Description: `Hourly policy item`,
				},
				resource.Attribute{
					Name:        "policy_item_daily",
					Description: `Daily policy item`,
				},
				resource.Attribute{
					Name:        "policy_item_weekly",
					Description: `Weekly policy item`,
				},
				resource.Attribute{
					Name:        "policy_item_monthly",
					Description: `Monthly policy item`,
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
					Description: `Human-readable label that identifies the cloud provider that stores the snapshot copy. i.e. "AWS" "AZURE" "GCP"`,
				},
				resource.Attribute{
					Name:        "frequencies",
					Description: `List that describes which types of snapshots to copy. i.e. "HOURLY" "DAILY" "WEEKLY" "MONTHLY" "ON_DEMAND"`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Target region to copy snapshots belonging to replicationSpecId to. Please supply the 'Atlas Region' which can be found under https://www.mongodb.com/docs/atlas/reference/cloud-providers/ 'regions' link`,
				},
				resource.Attribute{
					Name:        "replication_spec_id",
					Description: `Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica set clusters, there is only one zone in the cluster. To find the Replication Spec Id, do a GET request to Return One Cluster in One Project and consult the replicationSpecs array https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#operation/returnOneCluster`,
				},
				resource.Attribute{
					Name:        "should_copy_oplogs",
					Description: `Flag that indicates whether to copy the oplogs to the target region. You can use the oplogs to perform point-in-time restores.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Unique identifier of the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "next_snapshot",
					Description: `UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.`,
				},
				resource.Attribute{
					Name:        "reference_hour_of_day",
					Description: `UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.`,
				},
				resource.Attribute{
					Name:        "reference_minute_of_hour",
					Description: `UTC Minute of day between 0 and 59 representing which minute of the ` + "`" + `reference_hour_of_day` + "`" + ` that Atlas takes the snapshot.`,
				},
				resource.Attribute{
					Name:        "restore_window_days",
					Description: `Specifies a restore window in days for cloud backup to maintain.`,
				},
				resource.Attribute{
					Name:        "id_policy",
					Description: `Unique identifier of the backup policy.`,
				},
				resource.Attribute{
					Name:        "policy_item_hourly",
					Description: `Hourly policy item`,
				},
				resource.Attribute{
					Name:        "policy_item_daily",
					Description: `Daily policy item`,
				},
				resource.Attribute{
					Name:        "policy_item_weekly",
					Description: `Weekly policy item`,
				},
				resource.Attribute{
					Name:        "policy_item_monthly",
					Description: `Monthly policy item`,
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
					Description: `Human-readable label that identifies the cloud provider that stores the snapshot copy. i.e. "AWS" "AZURE" "GCP"`,
				},
				resource.Attribute{
					Name:        "frequencies",
					Description: `List that describes which types of snapshots to copy. i.e. "HOURLY" "DAILY" "WEEKLY" "MONTHLY" "ON_DEMAND"`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `Target region to copy snapshots belonging to replicationSpecId to. Please supply the 'Atlas Region' which can be found under https://www.mongodb.com/docs/atlas/reference/cloud-providers/ 'regions' link`,
				},
				resource.Attribute{
					Name:        "replication_spec_id",
					Description: `Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica set clusters, there is only one zone in the cluster. To find the Replication Spec Id, do a GET request to Return One Cluster in One Project and consult the replicationSpecs array https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#operation/returnOneCluster`,
				},
				resource.Attribute{
					Name:        "should_copy_oplogs",
					Description: `Flag that indicates whether to copy the oplogs to the target region. You can use the oplogs to perform point-in-time restores.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_backup_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Backup Snapshot Datasource.`,
			Description: `

` + "`" + `mongodbatlas_cloud_backup_snapshot` + "`" + ` provides an Cloud Backup Snapshot datasource. Atlas Cloud Backup Snapshots provide localized backup storage using the native snapshot functionality of the cluster’s cloud service.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
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
					Description: `Specifies the type of cluster: replicaSet or shardedCluster.`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `Cloud provider that stores this snapshot.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Block of List of snapshots and the cloud provider where the snapshots are stored. See below`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `Label given to the replica set from which Atlas took this snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `Unique identifiers of the snapshots created for the shards and config server for a sharded cluster. ### members`,
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
					Description: `Label given to a shard or config server from which Atlas took this snapshot. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/get-one-backup/)`,
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
					Description: `Specifies the type of cluster: replicaSet or shardedCluster.`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `Cloud provider that stores this snapshot.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Block of List of snapshots and the cloud provider where the snapshots are stored. See below`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `Label given to the replica set from which Atlas took this snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `Unique identifiers of the snapshots created for the shards and config server for a sharded cluster. ### members`,
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
					Description: `Label given to a shard or config server from which Atlas took this snapshot. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/get-one-backup/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_backup_snapshot_export_bucket",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Backup Snapshot Export Bucket resource.`,
			Description: `
` + "`" + `mongodbatlas_cloud_backup_snapshot_export_bucket` + "`" + ` datasource allows you to retrieve all the buckets for the specified project.


-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "export_bucket_id",
					Description: `(Required) Unique identifier of the snapshot export bucket. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "iam_role_id",
					Description: `Unique identifier of the role that Atlas can use to access the bucket. You must also specify the ` + "`" + `bucket_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `Name of the bucket that the provided role ID is authorized to access. You must also specify the ` + "`" + `iam_role_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `Name of the provider of the cloud service where Atlas can access the S3 bucket. Atlas only supports ` + "`" + `AWS` + "`" + `. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/export/create-one-export-bucket/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "iam_role_id",
					Description: `Unique identifier of the role that Atlas can use to access the bucket. You must also specify the ` + "`" + `bucket_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `Name of the bucket that the provided role ID is authorized to access. You must also specify the ` + "`" + `iam_role_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `Name of the provider of the cloud service where Atlas can access the S3 bucket. Atlas only supports ` + "`" + `AWS` + "`" + `. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/export/create-one-export-bucket/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_backup_snapshot_export_buckets",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Backup Snapshot Export Bucket resource.`,
			Description: `
` + "`" + `mongodbatlas_cloud_backup_snapshot_export_buckets` + "`" + ` datasource allows you to retrieve all the buckets for the specified project.


-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `One or more links to sub-resources and/or related resources.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `Includes CloudProviderSnapshotExportBucket object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### CloudProviderSnapshotExportBucket`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "export_bucket_id",
					Description: `Unique identifier of the snapshot bucket id.`,
				},
				resource.Attribute{
					Name:        "iam_role_id",
					Description: `Unique identifier of the role that Atlas can use to access the bucket. You must also specify the ` + "`" + `bucket_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `Name of the bucket that the provided role ID is authorized to access. You must also specify the ` + "`" + `iam_role_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `Name of the provider of the cloud service where Atlas can access the S3 bucket. Atlas only supports ` + "`" + `AWS` + "`" + `. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/export/create-one-export-bucket/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "links",
					Description: `One or more links to sub-resources and/or related resources.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `Includes CloudProviderSnapshotExportBucket object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### CloudProviderSnapshotExportBucket`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "export_bucket_id",
					Description: `Unique identifier of the snapshot bucket id.`,
				},
				resource.Attribute{
					Name:        "iam_role_id",
					Description: `Unique identifier of the role that Atlas can use to access the bucket. You must also specify the ` + "`" + `bucket_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `Name of the bucket that the provided role ID is authorized to access. You must also specify the ` + "`" + `iam_role_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `Name of the provider of the cloud service where Atlas can access the S3 bucket. Atlas only supports ` + "`" + `AWS` + "`" + `. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/export/create-one-export-bucket/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_backup_snapshot_export_job",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Backup Snapshot Export Job resource.`,
			Description: `
` + "`" + `mongodbatlas_cloud_backup_snapshot_export_job` + "`" + ` datasource allows you to retrieve a snapshot export job for the specified project and cluster.


-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the project which contains the Atlas cluster whose snapshot you want to retrieve.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) Name of the Atlas cluster whose export job you want to retrieve.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Unique identifier of the Cloud Backup snapshot to export.`,
				},
				resource.Attribute{
					Name:        "export_bucket_id",
					Description: `Unique identifier of the AWS bucket to export the Cloud Backup snapshot to.`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `Custom data to include in the metadata file named ` + "`" + `.complete` + "`" + ` that Atlas uploads to the bucket when the export job finishes. Custom data can be specified as key and value pairs.`,
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
					Description: `indicates that the export job has failed ### Custom Data`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Custom data specified as key in the key and value pair.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value for the key specified using ` + "`" + `key` + "`" + `. ### components`,
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
					Description: `_Returned for replica set only._ Total number of collections to export. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/export/get-one-export-job/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Unique identifier of the Cloud Backup snapshot to export.`,
				},
				resource.Attribute{
					Name:        "export_bucket_id",
					Description: `Unique identifier of the AWS bucket to export the Cloud Backup snapshot to.`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `Custom data to include in the metadata file named ` + "`" + `.complete` + "`" + ` that Atlas uploads to the bucket when the export job finishes. Custom data can be specified as key and value pairs.`,
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
					Description: `indicates that the export job has failed ### Custom Data`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Custom data specified as key in the key and value pair.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value for the key specified using ` + "`" + `key` + "`" + `. ### components`,
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
					Description: `_Returned for replica set only._ Total number of collections to export. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/export/get-one-export-job/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_backup_snapshot_export_jobs",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Backup Snapshot Export Jobs resource.`,
			Description: `
` + "`" + `mongodbatlas_cloud_backup_snapshot_export_jobs` + "`" + ` datasource allows you to retrieve all the buckets for the specified project.


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
					Description: `(Required) Name of the Atlas cluster whose export job you want to retrieve.`,
				},
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `One or more links to sub-resources and/or related resources.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `Includes CloudProviderSnapshotExportJob object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### CloudProviderSnapshotExportJob`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "export_job_id",
					Description: `Unique identifier of the S3 bucket.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Unique identifier of the Cloud Backup snapshot to export.`,
				},
				resource.Attribute{
					Name:        "export_bucket_id",
					Description: `Unique identifier of the AWS bucket to export the Cloud Backup snapshot to.`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `Custom data to include in the metadata file named ` + "`" + `.complete` + "`" + ` that Atlas uploads to the bucket when the export job finishes. Custom data can be specified as key and value pairs.`,
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
					Description: `indicates that the export job has failed #### Custom Data`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Custom data specified as key in the key and value pair.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value for the key specified using ` + "`" + `key` + "`" + `. #### components`,
				},
				resource.Attribute{
					Name:        "export_id",
					Description: `_Returned for sharded clusters only._ Export job details for each replica set in the sharded cluster.`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `_Returned for sharded clusters only._ Unique identifier of the export job for the replica set. #### export_status`,
				},
				resource.Attribute{
					Name:        "exported_collections",
					Description: `_Returned for replica set only._ Number of collections that have been exported.`,
				},
				resource.Attribute{
					Name:        "total_collections",
					Description: `_Returned for replica set only._ Total number of collections to export. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/export/get-all-export-jobs/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "links",
					Description: `One or more links to sub-resources and/or related resources.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `Includes CloudProviderSnapshotExportJob object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### CloudProviderSnapshotExportJob`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "export_job_id",
					Description: `Unique identifier of the S3 bucket.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Unique identifier of the Cloud Backup snapshot to export.`,
				},
				resource.Attribute{
					Name:        "export_bucket_id",
					Description: `Unique identifier of the AWS bucket to export the Cloud Backup snapshot to.`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `Custom data to include in the metadata file named ` + "`" + `.complete` + "`" + ` that Atlas uploads to the bucket when the export job finishes. Custom data can be specified as key and value pairs.`,
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
					Description: `indicates that the export job has failed #### Custom Data`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Custom data specified as key in the key and value pair.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value for the key specified using ` + "`" + `key` + "`" + `. #### components`,
				},
				resource.Attribute{
					Name:        "export_id",
					Description: `_Returned for sharded clusters only._ Export job details for each replica set in the sharded cluster.`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `_Returned for sharded clusters only._ Unique identifier of the export job for the replica set. #### export_status`,
				},
				resource.Attribute{
					Name:        "exported_collections",
					Description: `_Returned for replica set only._ Number of collections that have been exported.`,
				},
				resource.Attribute{
					Name:        "total_collections",
					Description: `_Returned for replica set only._ Total number of collections to export. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/export/get-all-export-jobs/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_backup_snapshot_restore_job",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Backup Snapshot Restore Job Datasource.`,
			Description: `

` + "`" + `mongodbatlas_cloud_backup_snapshot_restore_job` + "`" + ` provides a Cloud Backup Snapshot Restore Job datasource. Gets all the cloud backup snapshot restore jobs for the specified cluster.

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
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
				},
				resource.Attribute{
					Name:        "oplogInc",
					Description: `Oplog operation number from which to you want to restore this snapshot.`,
				},
				resource.Attribute{
					Name:        "pointInTimeUTCSeconds",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/get-one-restore-job/)`,
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
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
				},
				resource.Attribute{
					Name:        "oplogInc",
					Description: `Oplog operation number from which to you want to restore this snapshot.`,
				},
				resource.Attribute{
					Name:        "pointInTimeUTCSeconds",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/get-one-restore-job/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_backup_snapshot_restore_jobs",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Backup Snapshot Restore Jobs Datasource.`,
			Description: `

` + "`" + `mongodbatlas_cloud_backup_snapshot_restore_jobs` + "`" + ` provides a Cloud Backup Snapshot Restore Jobs datasource. Gets all the cloud backup snapshot restore jobs for the specified cluster.

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
					Description: `(Required) The name of the Atlas cluster for which you want to retrieve restore jobs.`,
				},
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
				},
				resource.Attribute{
					Name:        "oplogInc",
					Description: `Oplog operation number from which to you want to restore this snapshot.`,
				},
				resource.Attribute{
					Name:        "pointInTimeUTCSeconds",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/get-all-restore-jobs/)`,
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
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
				},
				resource.Attribute{
					Name:        "oplogInc",
					Description: `Oplog operation number from which to you want to restore this snapshot.`,
				},
				resource.Attribute{
					Name:        "pointInTimeUTCSeconds",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/get-all-restore-jobs/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_backup_snapshots",
			Category:         "Data Sources",
			ShortDescription: `Provides an Cloud Backup Snapshot Datasource.`,
			Description: `

` + "`" + `mongodbatlas_cloud_backup_snapshots` + "`" + ` provides an Cloud Backup Snapshot datasource. Atlas Cloud Backup Snapshots provide localized backup storage using the native snapshot functionality of the cluster’s cloud service.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Atlas cluster that contains the snapshot you want to retrieve.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `Specifies the type of cluster: replicaSet or shardedCluster.`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `Cloud provider that stores this snapshot.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Block of List of snapshots and the cloud provider where the snapshots are stored. See below`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `Label given to the replica set from which Atlas took this snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `Unique identifiers of the snapshots created for the shards and config server for a sharded cluster. ### members`,
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
					Description: `Label given to a shard or config server from which Atlas took this snapshot. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/get-all-backups/)`,
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
					Description: `Specifies the type of cluster: replicaSet or shardedCluster.`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `Cloud provider that stores this snapshot.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Block of List of snapshots and the cloud provider where the snapshots are stored. See below`,
				},
				resource.Attribute{
					Name:        "replica_set_name",
					Description: `Label given to the replica set from which Atlas took this snapshot.`,
				},
				resource.Attribute{
					Name:        "snapshot_ids",
					Description: `Unique identifiers of the snapshots created for the shards and config server for a sharded cluster. ### members`,
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
					Description: `Label given to a shard or config server from which Atlas took this snapshot. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/get-all-backups/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_access",
			Category:         "Data Sources",
			ShortDescription: `Allows you to get the list of cloud provider access roles`,
			Description: `

` + "`" + `mongodbatlas_cloud_provider_access` + "`" + ` allows you to get the list of cloud provider access roles, currently only AWS is supported.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to get all Cloud Provider Access ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "aws_iam_roles",
					Description: `A list where each represents a Cloud Provider Access Role. ### Each element in the aws_iam_roles array consists in an object with the following elements`,
				},
				resource.Attribute{
					Name:        "atlas_assumed_role_external_id",
					Description: `Unique external ID Atlas uses when assuming the IAM role in your AWS account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "aws_iam_roles",
					Description: `A list where each represents a Cloud Provider Access Role. ### Each element in the aws_iam_roles array consists in an object with the following elements`,
				},
				resource.Attribute{
					Name:        "atlas_assumed_role_external_id",
					Description: `Unique external ID Atlas uses when assuming the IAM role in your AWS account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_access_setup",
			Category:         "Data Sources",
			ShortDescription: `Allows you to get the a single role for cloud provider access setup`,
			Description: `

` + "`" + `mongodbatlas_cloud_provider_access` + "`" + ` allows you to get a single role for a provider access role setup, currently only AWS is supported.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to get all Cloud Provider Access`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) cloud provider name, currently only AWS is supported`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) unique role id among all the aws roles provided by mongodb atlas ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "atlas_assumed_role_external_id",
					Description: `Unique external ID Atlas uses when assuming the IAM role in your AWS account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "atlas_assumed_role_external_id",
					Description: `Unique external ID Atlas uses when assuming the IAM role in your AWS account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Provides an Cloud Backup Snapshot Datasource.`,
			Description: `

` + "`" + `mongodbatlas_cloud_provider_snapshot` + "`" + ` provides an Cloud Backup Snapshot datasource. Atlas Cloud Backup Snapshots provide localized backup storage using the native snapshot functionality of the cluster’s cloud service.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
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
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/get-one-backup/)`,
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
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/get-one-backup/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot_backup_policy",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Backup Snapshot Policy Datasource.`,
			Description: `

` + "`" + `mongodbatlas_cloud_provider_snapshot_backup_policy` + "`" + ` provides a Cloud Backup Snapshot Backup Policy datasource. An Atlas Cloud Backup Snapshot Policy provides the current snapshot schedule and retention settings for the cluster. 

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
					Description: `(Required) The name of the Atlas cluster that contains the snapshots backup policy you want to retrieve. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Unique identifier of the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "next_snapshot",
					Description: `UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.`,
				},
				resource.Attribute{
					Name:        "reference_hour_of_day",
					Description: `UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.`,
				},
				resource.Attribute{
					Name:        "reference_minute_of_hour",
					Description: `UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.`,
				},
				resource.Attribute{
					Name:        "restore_window_days",
					Description: `Specifies a restore window in days for cloud backup to maintain. ### Policies`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policy definitions for the cluster.`,
				},
				resource.Attribute{
					Name:        "policies.#.id",
					Description: `Unique identifier of the backup policy. #### Policy Item`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item",
					Description: `A list of specifications for a policy.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.id",
					Description: `Unique identifier for this policy item.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.frequency_interval",
					Description: `The frequency interval for a set of snapshots.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.frequency_type",
					Description: `A type of frequency (hourly, daily, weekly, monthly).`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.retention_unit",
					Description: `The unit of time in which snapshot retention is measured (days, weeks, months).`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.retention_value",
					Description: `The number of days, weeks, or months the snapshot is retained. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/schedule/get-all-schedules/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Unique identifier of the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "next_snapshot",
					Description: `UTC ISO 8601 formatted point in time when Atlas will take the next snapshot.`,
				},
				resource.Attribute{
					Name:        "reference_hour_of_day",
					Description: `UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot.`,
				},
				resource.Attribute{
					Name:        "reference_minute_of_hour",
					Description: `UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.`,
				},
				resource.Attribute{
					Name:        "restore_window_days",
					Description: `Specifies a restore window in days for cloud backup to maintain. ### Policies`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policy definitions for the cluster.`,
				},
				resource.Attribute{
					Name:        "policies.#.id",
					Description: `Unique identifier of the backup policy. #### Policy Item`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item",
					Description: `A list of specifications for a policy.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.id",
					Description: `Unique identifier for this policy item.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.frequency_interval",
					Description: `The frequency interval for a set of snapshots.`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.frequency_type",
					Description: `A type of frequency (hourly, daily, weekly, monthly).`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.retention_unit",
					Description: `The unit of time in which snapshot retention is measured (days, weeks, months).`,
				},
				resource.Attribute{
					Name:        "policies.#.policy_item.#.retention_value",
					Description: `The number of days, weeks, or months the snapshot is retained. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/schedule/get-all-schedules/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot_restore_job",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Backup Snapshot Restore Job Datasource.`,
			Description: `

` + "`" + `mongodbatlas_cloud_provider_snapshot_restore_job` + "`" + ` provides a Cloud Backup Snapshot Restore Job datasource. Gets all the cloud backup snapshot restore jobs for the specified cluster.

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
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
				},
				resource.Attribute{
					Name:        "oplogInc",
					Description: `Oplog operation number from which to you want to restore this snapshot.`,
				},
				resource.Attribute{
					Name:        "pointInTimeUTCSeconds",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/get-one-restore-job/)`,
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
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
				},
				resource.Attribute{
					Name:        "oplogInc",
					Description: `Oplog operation number from which to you want to restore this snapshot.`,
				},
				resource.Attribute{
					Name:        "pointInTimeUTCSeconds",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/get-one-restore-job/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshot_restore_jobs",
			Category:         "Data Sources",
			ShortDescription: `Provides a Cloud Backup Snapshot Restore Jobs Datasource.`,
			Description: `

` + "`" + `mongodbatlas_cloud_provider_snapshot_restore_jobs` + "`" + ` provides a Cloud Backup Snapshot Restore Jobs datasource. Gets all the cloud backup snapshot restore jobs for the specified cluster.

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
					Description: `(Required) The name of the Atlas cluster for which you want to retrieve restore jobs.`,
				},
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
				},
				resource.Attribute{
					Name:        "oplogInc",
					Description: `Oplog operation number from which to you want to restore this snapshot.`,
				},
				resource.Attribute{
					Name:        "pointInTimeUTCSeconds",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/get-all-restore-jobs/)`,
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
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch.`,
				},
				resource.Attribute{
					Name:        "oplogInc",
					Description: `Oplog operation number from which to you want to restore this snapshot.`,
				},
				resource.Attribute{
					Name:        "pointInTimeUTCSeconds",
					Description: `Timestamp in the number of seconds that have elapsed since the UNIX epoch. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/restore/get-all-restore-jobs/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cloud_provider_snapshots",
			Category:         "Data Sources",
			ShortDescription: `Provides an Cloud Backup Snapshot Datasource.`,
			Description: `

` + "`" + `mongodbatlas_cloud_provider_snapshots` + "`" + ` provides an Cloud Backup Snapshot datasource. Atlas Cloud Backup Snapshots provide localized backup storage using the native snapshot functionality of the cluster’s cloud service.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the Atlas cluster that contains the snapshot you want to retrieve.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The unique identifier of the project for the Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/get-all-backups/)`,
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
					Description: `Specifies the type of cluster: replicaSet or shardedCluster. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/cloud-backup/backup/get-all-backups/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_cluster",
			Category:         "Data Sources",
			ShortDescription: `Describe a Cluster.`,
			Description: `

` + "`" + `mongodbatlas_cluster` + "`" + ` describes a Cluster. The data source requires your Project ID.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

~> **IMPORTANT:**
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
					Name:        "name",
					Description: `(Required) Name of the cluster as it appears in Atlas. Once the cluster is created, its name cannot be changed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
				},
				resource.Attribute{
					Name:        "mongo_db_version",
					Description: `Version of MongoDB the cluster runs, in ` + "`" + `major-version` + "`" + `.` + "`" + `minor-version` + "`" + ` format.`,
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
					Name:        "pit_enabled",
					Description: `Flag that indicates if the cluster uses Continuous Cloud Backup.`,
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
					Name:        "auto_scaling_compute_enabled",
					Description: `Specifies whether cluster tier auto-scaling is enabled. The default is false.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_scale_down_enabled",
					Description: `Specifies whether cluster tier auto-down-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "backup_enabled",
					Description: `Legacy Option, Indicates whether Atlas continuous backups are enabled for the cluster.`,
				},
				resource.Attribute{
					Name:        "bi_connector",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "bi_connector_config",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Indicates the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.`,
				},
				resource.Attribute{
					Name:        "connection_strings",
					Description: `Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster. Ensure ` + "`" + `connection_strings` + "`" + ` are available following ` + "`" + `terraform apply` + "`" + ` by adding a ` + "`" + `depends_on` + "`" + ` relationship to the ` + "`" + `cluster` + "`" + `, ex: ` + "`" + `` + "`" + `` + "`" + ` data "mongodbatlas_cluster" "example_cluster" { project_id = var.project_id name = var.cluster_name depends_on = [mongodbatlas_privatelink_endpoint_service.example_endpoint] } ` + "`" + `` + "`" + `` + "`" + ` - ` + "`" + `connection_strings.standard` + "`" + ` - Public mongodb:// connection string for this cluster. - ` + "`" + `connection_strings.standard_srv` + "`" + ` - Public mongodb+srv:// connection string for this cluster. The mongodb+srv protocol tells the driver to look up the seed list of hosts in DNS. Atlas synchronizes this list with the nodes in a cluster. If the connection string uses this URI format, you don’t need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn’t, use connectionStrings.standard. - ` + "`" + `connection_strings.aws_private_link` + "`" + ` - [Private-endpoint-aware](https://docs.atlas.mongodb.com/security-private-endpoint/#private-endpoint-connection-strings) mongodb://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a AWS PrivateLink connection to this cluster.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Indicates the size in gigabytes of the server’s root volume (AWS/GCP Only).`,
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
					Name:        "cloud_backup",
					Description: `Flag indicating if the cluster uses Cloud Backup Snapshots for backups.`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.`,
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
					Description: `Describes Azure disk type of the server’s root volume (Azure Only).`,
				},
				resource.Attribute{
					Name:        "provider_region_name",
					Description: `Indicates Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. Requires the Atlas Region name, see the reference list for [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).`,
				},
				resource.Attribute{
					Name:        "provider_volume_type",
					Description: `Indicates the type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `(Deprecated) Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_min_instance_size",
					Description: `Minimum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_max_instance_size",
					Description: `Maximum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `The Network Peering Container ID. ->`,
				},
				resource.Attribute{
					Name:        "version_release_system",
					Description: `Release cadence that Atlas uses for this cluster.`,
				},
				resource.Attribute{
					Name:        "advanced_configuration",
					Description: `Get the advanced configuration options. See [Advanced Configuration](#advanced-configuration) below for more details. ### BI Connector Indicates BI Connector for Atlas configuration.`,
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
					Description: `Indicates the number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. ### Labels Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. Note: the key ` + "`" + `Infrastructure Tool` + "`" + `, is used for internal purposes to track aggregate usage.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that was set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that represents the key. ### Plugin Contains a key-value pair that tags that the cluster was created by a Terraform Provider and notes the version.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the current plugin`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of the plugin. ### Cloud Provider Snapshot Backup Policy`,
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
					Description: `The number of days, weeks, or months the snapshot is retained. #### Advanced Configuration`,
				},
				resource.Attribute{
					Name:        "default_read_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for read operations](https://docs.mongodb.com/manual/reference/read-concern/) set for this cluster. MongoDB 4.4 clusters default to [available](https://docs.mongodb.com/manual/reference/read-concern-available/).`,
				},
				resource.Attribute{
					Name:        "default_write_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for write operations](https://docs.mongodb.com/manual/reference/write-concern/) set for this cluster. MongoDB 4.4 clusters default to [1](https://docs.mongodb.com/manual/reference/write-concern/).`,
				},
				resource.Attribute{
					Name:        "fail_index_key_too_long",
					Description: `When true, documents can only be updated or inserted if, for all indexed fields on the target collection, the corresponding index entries do not exceed 1024 bytes. When false, mongod writes documents that exceed the limit but does not index them.`,
				},
				resource.Attribute{
					Name:        "javascript_enabled",
					Description: `When true, the cluster allows execution of operations that perform server-side executions of JavaScript. When false, the cluster disables execution of those operations.`,
				},
				resource.Attribute{
					Name:        "minimum_enabled_tls_protocol",
					Description: `Sets the minimum Transport Layer Security (TLS) version the cluster accepts for incoming connections.Valid values are: - TLS1_0 - TLS1_1 - TLS1_2`,
				},
				resource.Attribute{
					Name:        "no_table_scan",
					Description: `When true, the cluster disables the execution of any query that requires a collection scan to return results. When false, the cluster allows the execution of those operations.`,
				},
				resource.Attribute{
					Name:        "oplog_size_mb",
					Description: `The custom oplog size of the cluster. Without a value that indicates that the cluster uses the default oplog size calculated by Atlas.`,
				},
				resource.Attribute{
					Name:        "oplog_min_retention_hours",
					Description: `Minimum retention window for cluster's oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates.`,
				},
				resource.Attribute{
					Name:        "sample_size_bi_connector",
					Description: `Number of documents per database to sample when gathering schema information. Defaults to 100. Available only for Atlas deployments in which BI Connector for Atlas is enabled.`,
				},
				resource.Attribute{
					Name:        "sample_refresh_interval_bi_connector",
					Description: `Interval in seconds at which the mongosqld process re-samples data to create its relational schema. The default value is 300. The specified value must be a positive integer. Available only for Atlas deployments in which BI Connector for Atlas is enabled. See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The cluster ID.`,
				},
				resource.Attribute{
					Name:        "mongo_db_version",
					Description: `Version of MongoDB the cluster runs, in ` + "`" + `major-version` + "`" + `.` + "`" + `minor-version` + "`" + ` format.`,
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
					Name:        "pit_enabled",
					Description: `Flag that indicates if the cluster uses Continuous Cloud Backup.`,
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
					Name:        "auto_scaling_compute_enabled",
					Description: `Specifies whether cluster tier auto-scaling is enabled. The default is false.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_scale_down_enabled",
					Description: `Specifies whether cluster tier auto-down-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "backup_enabled",
					Description: `Legacy Option, Indicates whether Atlas continuous backups are enabled for the cluster.`,
				},
				resource.Attribute{
					Name:        "bi_connector",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "bi_connector_config",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Indicates the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.`,
				},
				resource.Attribute{
					Name:        "connection_strings",
					Description: `Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster. Ensure ` + "`" + `connection_strings` + "`" + ` are available following ` + "`" + `terraform apply` + "`" + ` by adding a ` + "`" + `depends_on` + "`" + ` relationship to the ` + "`" + `cluster` + "`" + `, ex: ` + "`" + `` + "`" + `` + "`" + ` data "mongodbatlas_cluster" "example_cluster" { project_id = var.project_id name = var.cluster_name depends_on = [mongodbatlas_privatelink_endpoint_service.example_endpoint] } ` + "`" + `` + "`" + `` + "`" + ` - ` + "`" + `connection_strings.standard` + "`" + ` - Public mongodb:// connection string for this cluster. - ` + "`" + `connection_strings.standard_srv` + "`" + ` - Public mongodb+srv:// connection string for this cluster. The mongodb+srv protocol tells the driver to look up the seed list of hosts in DNS. Atlas synchronizes this list with the nodes in a cluster. If the connection string uses this URI format, you don’t need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn’t, use connectionStrings.standard. - ` + "`" + `connection_strings.aws_private_link` + "`" + ` - [Private-endpoint-aware](https://docs.atlas.mongodb.com/security-private-endpoint/#private-endpoint-connection-strings) mongodb://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a AWS PrivateLink connection to this cluster.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Indicates the size in gigabytes of the server’s root volume (AWS/GCP Only).`,
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
					Name:        "cloud_backup",
					Description: `Flag indicating if the cluster uses Cloud Backup Snapshots for backups.`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.`,
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
					Description: `Describes Azure disk type of the server’s root volume (Azure Only).`,
				},
				resource.Attribute{
					Name:        "provider_region_name",
					Description: `Indicates Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. Requires the Atlas Region name, see the reference list for [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).`,
				},
				resource.Attribute{
					Name:        "provider_volume_type",
					Description: `Indicates the type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `(Deprecated) Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_min_instance_size",
					Description: `Minimum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_max_instance_size",
					Description: `Maximum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `The Network Peering Container ID. ->`,
				},
				resource.Attribute{
					Name:        "version_release_system",
					Description: `Release cadence that Atlas uses for this cluster.`,
				},
				resource.Attribute{
					Name:        "advanced_configuration",
					Description: `Get the advanced configuration options. See [Advanced Configuration](#advanced-configuration) below for more details. ### BI Connector Indicates BI Connector for Atlas configuration.`,
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
					Description: `Indicates the number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. ### Labels Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. Note: the key ` + "`" + `Infrastructure Tool` + "`" + `, is used for internal purposes to track aggregate usage.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that was set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that represents the key. ### Plugin Contains a key-value pair that tags that the cluster was created by a Terraform Provider and notes the version.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the current plugin`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of the plugin. ### Cloud Provider Snapshot Backup Policy`,
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
					Description: `The number of days, weeks, or months the snapshot is retained. #### Advanced Configuration`,
				},
				resource.Attribute{
					Name:        "default_read_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for read operations](https://docs.mongodb.com/manual/reference/read-concern/) set for this cluster. MongoDB 4.4 clusters default to [available](https://docs.mongodb.com/manual/reference/read-concern-available/).`,
				},
				resource.Attribute{
					Name:        "default_write_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for write operations](https://docs.mongodb.com/manual/reference/write-concern/) set for this cluster. MongoDB 4.4 clusters default to [1](https://docs.mongodb.com/manual/reference/write-concern/).`,
				},
				resource.Attribute{
					Name:        "fail_index_key_too_long",
					Description: `When true, documents can only be updated or inserted if, for all indexed fields on the target collection, the corresponding index entries do not exceed 1024 bytes. When false, mongod writes documents that exceed the limit but does not index them.`,
				},
				resource.Attribute{
					Name:        "javascript_enabled",
					Description: `When true, the cluster allows execution of operations that perform server-side executions of JavaScript. When false, the cluster disables execution of those operations.`,
				},
				resource.Attribute{
					Name:        "minimum_enabled_tls_protocol",
					Description: `Sets the minimum Transport Layer Security (TLS) version the cluster accepts for incoming connections.Valid values are: - TLS1_0 - TLS1_1 - TLS1_2`,
				},
				resource.Attribute{
					Name:        "no_table_scan",
					Description: `When true, the cluster disables the execution of any query that requires a collection scan to return results. When false, the cluster allows the execution of those operations.`,
				},
				resource.Attribute{
					Name:        "oplog_size_mb",
					Description: `The custom oplog size of the cluster. Without a value that indicates that the cluster uses the default oplog size calculated by Atlas.`,
				},
				resource.Attribute{
					Name:        "oplog_min_retention_hours",
					Description: `Minimum retention window for cluster's oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates.`,
				},
				resource.Attribute{
					Name:        "sample_size_bi_connector",
					Description: `Number of documents per database to sample when gathering schema information. Defaults to 100. Available only for Atlas deployments in which BI Connector for Atlas is enabled.`,
				},
				resource.Attribute{
					Name:        "sample_refresh_interval_bi_connector",
					Description: `Interval in seconds at which the mongosqld process re-samples data to create its relational schema. The default value is 300. The specified value must be a positive integer. Available only for Atlas deployments in which BI Connector for Atlas is enabled. See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_clusters",
			Category:         "Data Sources",
			ShortDescription: `Describe all Clusters in Project.`,
			Description: `

` + "`" + `mongodbatlas_cluster` + "`" + ` describes all Clusters by the provided project_id. The data source requires your Project ID.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

~> **IMPORTANT:**
<br> &#8226; Changes to cluster configurations can affect costs. Before making changes, please see [Billing](https://docs.atlas.mongodb.com/billing/).
<br> &#8226; If your Atlas project contains a custom role that uses actions introduced in a specific MongoDB version, you cannot create a cluster with a MongoDB version less than that version unless you delete the custom role.

`,
			Keywords: []string{},
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
					Name:        "mongo_db_version",
					Description: `Version of MongoDB the cluster runs, in ` + "`" + `major-version` + "`" + `.` + "`" + `minor-version` + "`" + ` format.`,
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
					Name:        "pit_enabled",
					Description: `Flag that indicates if the cluster uses Continuous Cloud Backup.`,
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
					Name:        "auto_scaling_compute_enabled",
					Description: `Specifies whether cluster tier auto-scaling is enabled. The default is false.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_scale_down_enabled",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_scale_down_enabled",
					Description: `Specifies whether cluster tier auto-down-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "backup_enabled",
					Description: `Legacy Option, Indicates whether Atlas continuous backups are enabled for the cluster.`,
				},
				resource.Attribute{
					Name:        "bi_connector",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "bi_connector_config",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Indicates the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.`,
				},
				resource.Attribute{
					Name:        "connection_strings",
					Description: `Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster. - ` + "`" + `connection_strings.standard` + "`" + ` - Public mongodb:// connection string for this cluster. - ` + "`" + `connection_strings.standard_srv` + "`" + ` - Public mongodb+srv:// connection string for this cluster. The mongodb+srv protocol tells the driver to look up the seed list of hosts in DNS. Atlas synchronizes this list with the nodes in a cluster. If the connection string uses this URI format, you don’t need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn’t, use connectionStrings.standard. - ` + "`" + `connection_strings.aws_private_link` + "`" + ` - [Private-endpoint-aware](https://docs.atlas.mongodb.com/security-private-endpoint/#private-endpoint-connection-strings) mongodb://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a AWS PrivateLink connection to this cluster.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Indicates the size in gigabytes of the server’s root volume (AWS/GCP Only).`,
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
					Description: `Flag indicating if the cluster uses Cloud Backup Snapshots for backups.`,
				},
				resource.Attribute{
					Name:        "cloud_backup",
					Description: `Flag indicating if the cluster uses Cloud Backup Snapshots for backups.`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.`,
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
					Description: `Describes Azure disk type of the server’s root volume (Azure Only).`,
				},
				resource.Attribute{
					Name:        "provider_region_name",
					Description: `Indicates Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. Requires the Atlas Region name, see the reference list for [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).`,
				},
				resource.Attribute{
					Name:        "provider_volume_type",
					Description: `Indicates the type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_min_instance_size",
					Description: `Minimum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_max_instance_size",
					Description: `Maximum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `(Deprecated) Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `The Network Peering Container ID. ->`,
				},
				resource.Attribute{
					Name:        "version_release_system",
					Description: `Release cadence that Atlas uses for this cluster.`,
				},
				resource.Attribute{
					Name:        "advanced_configuration",
					Description: `Get the advanced configuration options. See [Advanced Configuration](#advanced-configuration) below for more details. ### BI Connector Indicates BI Connector for Atlas configuration.`,
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
					Description: `Indicates the number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. ### Labels Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. Note: the key ` + "`" + `Infrastructure Tool` + "`" + `, is used for internal purposes to track aggregate usage.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that was set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that represents the key. ### Plugin Contains a key-value pair that tags that the cluster was created by a Terraform Provider and notes the version.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the current plugin`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of the plugin. ### Cloud Provider Snapshot Backup Policy`,
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
					Description: `The number of days, weeks, or months the snapshot is retained. #### Advanced Configuration`,
				},
				resource.Attribute{
					Name:        "default_read_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for read operations](https://docs.mongodb.com/manual/reference/read-concern/) set for this cluster. MongoDB 4.4 clusters default to [available](https://docs.mongodb.com/manual/reference/read-concern-available/).`,
				},
				resource.Attribute{
					Name:        "default_write_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for write operations](https://docs.mongodb.com/manual/reference/write-concern/) set for this cluster. MongoDB 4.4 clusters default to [1](https://docs.mongodb.com/manual/reference/write-concern/).`,
				},
				resource.Attribute{
					Name:        "fail_index_key_too_long",
					Description: `When true, documents can only be updated or inserted if, for all indexed fields on the target collection, the corresponding index entries do not exceed 1024 bytes. When false, mongod writes documents that exceed the limit but does not index them.`,
				},
				resource.Attribute{
					Name:        "javascript_enabled",
					Description: `When true, the cluster allows execution of operations that perform server-side executions of JavaScript. When false, the cluster disables execution of those operations.`,
				},
				resource.Attribute{
					Name:        "minimum_enabled_tls_protocol",
					Description: `Sets the minimum Transport Layer Security (TLS) version the cluster accepts for incoming connections.Valid values are: - TLS1_0 - TLS1_1 - TLS1_2`,
				},
				resource.Attribute{
					Name:        "no_table_scan",
					Description: `When true, the cluster disables the execution of any query that requires a collection scan to return results. When false, the cluster allows the execution of those operations.`,
				},
				resource.Attribute{
					Name:        "oplog_size_mb",
					Description: `The custom oplog size of the cluster. Without a value that indicates that the cluster uses the default oplog size calculated by Atlas.`,
				},
				resource.Attribute{
					Name:        "oplog_min_retention_hours",
					Description: `Minimum retention window for cluster's oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates.`,
				},
				resource.Attribute{
					Name:        "sample_size_bi_connector",
					Description: `Number of documents per database to sample when gathering schema information. Defaults to 100. Available only for Atlas deployments in which BI Connector for Atlas is enabled.`,
				},
				resource.Attribute{
					Name:        "sample_refresh_interval_bi_connector",
					Description: `Interval in seconds at which the mongosqld process re-samples data to create its relational schema. The default value is 300. The specified value must be a positive integer. Available only for Atlas deployments in which BI Connector for Atlas is enabled. See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
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
					Name:        "mongo_db_version",
					Description: `Version of MongoDB the cluster runs, in ` + "`" + `major-version` + "`" + `.` + "`" + `minor-version` + "`" + ` format.`,
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
					Name:        "pit_enabled",
					Description: `Flag that indicates if the cluster uses Continuous Cloud Backup.`,
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
					Name:        "auto_scaling_compute_enabled",
					Description: `Specifies whether cluster tier auto-scaling is enabled. The default is false.`,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_scale_down_enabled",
					Description: ``,
				},
				resource.Attribute{
					Name:        "auto_scaling_compute_scale_down_enabled",
					Description: `Specifies whether cluster tier auto-down-scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "backup_enabled",
					Description: `Legacy Option, Indicates whether Atlas continuous backups are enabled for the cluster.`,
				},
				resource.Attribute{
					Name:        "bi_connector",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "bi_connector_config",
					Description: `Indicates BI Connector for Atlas configuration on this cluster. BI Connector for Atlas is only available for M10+ clusters. See [BI Connector](#bi-connector) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `Indicates the type of the cluster that you want to modify. You cannot convert a sharded cluster deployment to a replica set deployment.`,
				},
				resource.Attribute{
					Name:        "connection_strings",
					Description: `Set of connection strings that your applications use to connect to this cluster. More info in [Connection-strings](https://docs.mongodb.com/manual/reference/connection-string/). Use the parameters in this object to connect your applications to this cluster. To learn more about the formats of connection strings, see [Connection String Options](https://docs.atlas.mongodb.com/reference/faq/connection-changes/). NOTE: Atlas returns the contents of this object after the cluster is operational, not while it builds the cluster. - ` + "`" + `connection_strings.standard` + "`" + ` - Public mongodb:// connection string for this cluster. - ` + "`" + `connection_strings.standard_srv` + "`" + ` - Public mongodb+srv:// connection string for this cluster. The mongodb+srv protocol tells the driver to look up the seed list of hosts in DNS. Atlas synchronizes this list with the nodes in a cluster. If the connection string uses this URI format, you don’t need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn’t, use connectionStrings.standard. - ` + "`" + `connection_strings.aws_private_link` + "`" + ` - [Private-endpoint-aware](https://docs.atlas.mongodb.com/security-private-endpoint/#private-endpoint-connection-strings) mongodb://connection strings for each interface VPC endpoint you configured to connect to this cluster. Returned only if you created a AWS PrivateLink connection to this cluster.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Indicates the size in gigabytes of the server’s root volume (AWS/GCP Only).`,
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
					Description: `Flag indicating if the cluster uses Cloud Backup Snapshots for backups.`,
				},
				resource.Attribute{
					Name:        "cloud_backup",
					Description: `Flag indicating if the cluster uses Cloud Backup Snapshots for backups.`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster.`,
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
					Description: `Describes Azure disk type of the server’s root volume (Azure Only).`,
				},
				resource.Attribute{
					Name:        "provider_region_name",
					Description: `Indicates Physical location of your MongoDB cluster. The region you choose can affect network latency for clients accessing your databases. Requires the Atlas Region name, see the reference list for [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).`,
				},
				resource.Attribute{
					Name:        "provider_volume_type",
					Description: `Indicates the type of the volume. The possible values are: ` + "`" + `STANDARD` + "`" + ` and ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_min_instance_size",
					Description: `Minimum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "provider_auto_scaling_compute_max_instance_size",
					Description: `Maximum instance size to which your cluster can automatically scale.`,
				},
				resource.Attribute{
					Name:        "replication_factor",
					Description: `(Deprecated) Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy. The possible values are 3, 5, or 7. The default value is 3.`,
				},
				resource.Attribute{
					Name:        "replication_specs",
					Description: `Configuration for cluster regions. See [Replication Spec](#replication-spec) below for more details.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `The Network Peering Container ID. ->`,
				},
				resource.Attribute{
					Name:        "version_release_system",
					Description: `Release cadence that Atlas uses for this cluster.`,
				},
				resource.Attribute{
					Name:        "advanced_configuration",
					Description: `Get the advanced configuration options. See [Advanced Configuration](#advanced-configuration) below for more details. ### BI Connector Indicates BI Connector for Atlas configuration.`,
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
					Description: `Indicates the number of analytics nodes for Atlas to deploy to the region. Analytics nodes are useful for handling analytic data such as reporting queries from BI Connector for Atlas. Analytics nodes are read-only, and can never become the primary. ### Labels Key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters. Note: the key ` + "`" + `Infrastructure Tool` + "`" + `, is used for internal purposes to track aggregate usage.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that was set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that represents the key. ### Plugin Contains a key-value pair that tags that the cluster was created by a Terraform Provider and notes the version.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the current plugin`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The current version of the plugin. ### Cloud Provider Snapshot Backup Policy`,
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
					Description: `The number of days, weeks, or months the snapshot is retained. #### Advanced Configuration`,
				},
				resource.Attribute{
					Name:        "default_read_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for read operations](https://docs.mongodb.com/manual/reference/read-concern/) set for this cluster. MongoDB 4.4 clusters default to [available](https://docs.mongodb.com/manual/reference/read-concern-available/).`,
				},
				resource.Attribute{
					Name:        "default_write_concern",
					Description: `[Default level of acknowledgment requested from MongoDB for write operations](https://docs.mongodb.com/manual/reference/write-concern/) set for this cluster. MongoDB 4.4 clusters default to [1](https://docs.mongodb.com/manual/reference/write-concern/).`,
				},
				resource.Attribute{
					Name:        "fail_index_key_too_long",
					Description: `When true, documents can only be updated or inserted if, for all indexed fields on the target collection, the corresponding index entries do not exceed 1024 bytes. When false, mongod writes documents that exceed the limit but does not index them.`,
				},
				resource.Attribute{
					Name:        "javascript_enabled",
					Description: `When true, the cluster allows execution of operations that perform server-side executions of JavaScript. When false, the cluster disables execution of those operations.`,
				},
				resource.Attribute{
					Name:        "minimum_enabled_tls_protocol",
					Description: `Sets the minimum Transport Layer Security (TLS) version the cluster accepts for incoming connections.Valid values are: - TLS1_0 - TLS1_1 - TLS1_2`,
				},
				resource.Attribute{
					Name:        "no_table_scan",
					Description: `When true, the cluster disables the execution of any query that requires a collection scan to return results. When false, the cluster allows the execution of those operations.`,
				},
				resource.Attribute{
					Name:        "oplog_size_mb",
					Description: `The custom oplog size of the cluster. Without a value that indicates that the cluster uses the default oplog size calculated by Atlas.`,
				},
				resource.Attribute{
					Name:        "oplog_min_retention_hours",
					Description: `Minimum retention window for cluster's oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates.`,
				},
				resource.Attribute{
					Name:        "sample_size_bi_connector",
					Description: `Number of documents per database to sample when gathering schema information. Defaults to 100. Available only for Atlas deployments in which BI Connector for Atlas is enabled.`,
				},
				resource.Attribute{
					Name:        "sample_refresh_interval_bi_connector",
					Description: `Interval in seconds at which the mongosqld process re-samples data to create its relational schema. The default value is 300. The specified value must be a positive integer. Available only for Atlas deployments in which BI Connector for Atlas is enabled. See detailed information for arguments and attributes: [MongoDB API Clusters](https://docs.atlas.mongodb.com/reference/api/clusters-create-one/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_custom_db_role",
			Category:         "Data Sources",
			ShortDescription: `Describes a Custom DB Role.`,
			Description: `

` + "`" + `mongodbatlas_custom_db_role` + "`" + ` describe a Custom DB Role. This represents a custom db role.

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
					Description: `(Required) Name of the custom role. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages and can be used to import. ### Actions Each object in the actions array represents an individual privilege action granted by the role. It is an required field.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Name of the privilege action. For a complete list of actions available in the Atlas API, see Custom Role Actions.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Contains information on where the action is granted. Each object in the array either indicates a database and collection on which the action is granted, or indicates that the action is granted on the cluster resource.`,
				},
				resource.Attribute{
					Name:        "resources.#.collection_name",
					Description: `(Optional) Collection on which the action is granted. If this value is an empty string, the action is granted on all collections within the database specified in the actions.resources.db field.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages and can be used to import. ### Actions Each object in the actions array represents an individual privilege action granted by the role. It is an required field.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Name of the privilege action. For a complete list of actions available in the Atlas API, see Custom Role Actions.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Contains information on where the action is granted. Each object in the array either indicates a database and collection on which the action is granted, or indicates that the action is granted on the cluster resource.`,
				},
				resource.Attribute{
					Name:        "resources.#.collection_name",
					Description: `(Optional) Collection on which the action is granted. If this value is an empty string, the action is granted on all collections within the database specified in the actions.resources.db field.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_custom_db_roles",
			Category:         "Data Sources",
			ShortDescription: `Describes a Custom DB Roles.`,
			Description: `

` + "`" + `mongodbatlas_custom_db_roles` + "`" + ` describe all Custom DB Roles. This represents a custom db roles.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to get all custom db roles. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a custom db roles. ### Actions Each object in the actions array represents an individual privilege action granted by the role. It is an required field.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Name of the privilege action. For a complete list of actions available in the Atlas API, see Custom Role Actions.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Contains information on where the action is granted. Each object in the array either indicates a database and collection on which the action is granted, or indicates that the action is granted on the cluster resource.`,
				},
				resource.Attribute{
					Name:        "resources.#.collection_name",
					Description: `(Optional) Collection on which the action is granted. If this value is an empty string, the action is granted on all collections within the database specified in the actions.resources.db field.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a custom db roles. ### Actions Each object in the actions array represents an individual privilege action granted by the role. It is an required field.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Name of the privilege action. For a complete list of actions available in the Atlas API, see Custom Role Actions.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Contains information on where the action is granted. Each object in the array either indicates a database and collection on which the action is granted, or indicates that the action is granted on the cluster resource.`,
				},
				resource.Attribute{
					Name:        "resources.#.collection_name",
					Description: `(Optional) Collection on which the action is granted. If this value is an empty string, the action is granted on all collections within the database specified in the actions.resources.db field.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_custom_dns_configuration_cluster_aws",
			Category:         "Data Sources",
			ShortDescription: `Describes a Custom DNS Configuration for Atlas Clusters on AWS.`,
			Description: `

` + "`" + `mongodbatlas_custom_dns_configuration_cluster_aws` + "`" + ` describes a Custom DNS Configuration for Atlas Clusters on AWS.

-> **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique identifier for the project. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The project ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates whether the project's clusters deployed to AWS use custom DNS. See detailed information for arguments and attributes: [MongoDB API Custom DNS Configuration for Atlas Clusters on AWS](https://docs.atlas.mongodb.com/reference/api/aws-custom-dns-get)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The project ID.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates whether the project's clusters deployed to AWS use custom DNS. See detailed information for arguments and attributes: [MongoDB API Custom DNS Configuration for Atlas Clusters on AWS](https://docs.atlas.mongodb.com/reference/api/aws-custom-dns-get)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_data_lake",
			Category:         "Data Sources",
			ShortDescription: `Describes a Data Lake.`,
			Description: `

` + "`" + `mongodbatlas_data_lake` + "`" + ` describe a Data Lake.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the data lake.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create a data lake. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "aws",
					Description: `AWS provider of the cloud service where Data Lake can access the S3 Bucket.`,
				},
				resource.Attribute{
					Name:        "aws.0.role_id",
					Description: `Unique identifier of the role that Data Lake can use to access the data stores.`,
				},
				resource.Attribute{
					Name:        "aws.0.test_s3_bucket",
					Description: `Name of the S3 data bucket that the provided role ID is authorized to access.`,
				},
				resource.Attribute{
					Name:        "aws.0.role_id",
					Description: `Unique identifier of the role that Data Lake can use to access the data stores.`,
				},
				resource.Attribute{
					Name:        "aws.0.test_s3_bucket",
					Description: `Name of the S3 data bucket that the provided role ID is authorized to access.`,
				},
				resource.Attribute{
					Name:        "aws.0.iam_assumed_role_arn",
					Description: `Amazon Resource Name (ARN) of the IAM Role that Data Lake assumes when accessing S3 Bucket data stores. For more information on S3 actions, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazons3.html).`,
				},
				resource.Attribute{
					Name:        "aws_iam_user_arn",
					Description: `Amazon Resource Name (ARN) of the user that Data Lake assumes when accessing S3 Bucket data stores.`,
				},
				resource.Attribute{
					Name:        "aws_external_id",
					Description: `Unique identifier associated with the IAM Role that Data Lake assumes when accessing the data stores.`,
				},
				resource.Attribute{
					Name:        "data_process_region",
					Description: `The cloud provider region to which Atlas Data Lake routes client connections for data processing.`,
				},
				resource.Attribute{
					Name:        "data_process_region.0.cloud_provider",
					Description: `Name of the cloud service provider.`,
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
					Description: `Configuration details for mapping each data store to queryable databases and collections.`,
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
					Description: `Name of a data store to map to the ` + "`" + `<collection>` + "`" + `.`,
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
					Description: `Array of objects where each object represents an [aggregation pipeline](https://docs.mongodb.com/manual/core/aggregation-pipeline/#id1) on a collection.`,
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
					Description: `Each object in the array represents a data store. Data Lake uses the storage.databases configuration details to map data in each data store to queryable databases and collections.`,
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
					Description: `Determines whether or not to use S3 tags on the files in the given path as additional partition attributes. See [MongoDB Atlas API](https://docs.mongodb.com/datalake/reference/api/dataLakes-get-one-tenant) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "aws",
					Description: `AWS provider of the cloud service where Data Lake can access the S3 Bucket.`,
				},
				resource.Attribute{
					Name:        "aws.0.role_id",
					Description: `Unique identifier of the role that Data Lake can use to access the data stores.`,
				},
				resource.Attribute{
					Name:        "aws.0.test_s3_bucket",
					Description: `Name of the S3 data bucket that the provided role ID is authorized to access.`,
				},
				resource.Attribute{
					Name:        "aws.0.role_id",
					Description: `Unique identifier of the role that Data Lake can use to access the data stores.`,
				},
				resource.Attribute{
					Name:        "aws.0.test_s3_bucket",
					Description: `Name of the S3 data bucket that the provided role ID is authorized to access.`,
				},
				resource.Attribute{
					Name:        "aws.0.iam_assumed_role_arn",
					Description: `Amazon Resource Name (ARN) of the IAM Role that Data Lake assumes when accessing S3 Bucket data stores. For more information on S3 actions, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazons3.html).`,
				},
				resource.Attribute{
					Name:        "aws_iam_user_arn",
					Description: `Amazon Resource Name (ARN) of the user that Data Lake assumes when accessing S3 Bucket data stores.`,
				},
				resource.Attribute{
					Name:        "aws_external_id",
					Description: `Unique identifier associated with the IAM Role that Data Lake assumes when accessing the data stores.`,
				},
				resource.Attribute{
					Name:        "data_process_region",
					Description: `The cloud provider region to which Atlas Data Lake routes client connections for data processing.`,
				},
				resource.Attribute{
					Name:        "data_process_region.0.cloud_provider",
					Description: `Name of the cloud service provider.`,
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
					Description: `Configuration details for mapping each data store to queryable databases and collections.`,
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
					Description: `Name of a data store to map to the ` + "`" + `<collection>` + "`" + `.`,
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
					Description: `Array of objects where each object represents an [aggregation pipeline](https://docs.mongodb.com/manual/core/aggregation-pipeline/#id1) on a collection.`,
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
					Description: `Each object in the array represents a data store. Data Lake uses the storage.databases configuration details to map data in each data store to queryable databases and collections.`,
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
					Description: `Determines whether or not to use S3 tags on the files in the given path as additional partition attributes. See [MongoDB Atlas API](https://docs.mongodb.com/datalake/reference/api/dataLakes-get-one-tenant) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_data_lakes",
			Category:         "Data Sources",
			ShortDescription: `Describes a Data Lakes.`,
			Description: `

` + "`" + `mongodbatlas_data_lakes` + "`" + ` describe all Data Lakes.


-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to get all data lakes. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Data lake. ### Data Lake`,
				},
				resource.Attribute{
					Name:        "aws_role_id",
					Description: `Unique identifier of the role that Data Lake can use to access the data stores.`,
				},
				resource.Attribute{
					Name:        "aws_test_s3_bucket",
					Description: `Name of the S3 data bucket that the provided role ID is authorized to access.`,
				},
				resource.Attribute{
					Name:        "data_process_region",
					Description: `The cloud provider region to which Atlas Data Lake routes client connections for data processing.`,
				},
				resource.Attribute{
					Name:        "data_process_region.0.cloud_provider",
					Description: `Name of the cloud service provider.`,
				},
				resource.Attribute{
					Name:        "aws_iam_assumed_role_arn",
					Description: `Amazon Resource Name (ARN) of the IAM Role that Data Lake assumes when accessing S3 Bucket data stores. For more information on S3 actions, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazons3.html).`,
				},
				resource.Attribute{
					Name:        "aws_iam_user_arn",
					Description: `Amazon Resource Name (ARN) of the user that Data Lake assumes when accessing S3 Bucket data stores.`,
				},
				resource.Attribute{
					Name:        "aws_external_id",
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
					Description: `Configuration details for mapping each data store to queryable databases and collections.`,
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
					Description: `Name of a data store to map to the ` + "`" + `<collection>` + "`" + `.`,
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
					Description: `Array of objects where each object represents an [aggregation pipeline](https://docs.mongodb.com/manual/core/aggregation-pipeline/#id1) on a collection.`,
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
					Description: `Each object in the array represents a data store. Data Lake uses the storage.databases configuration details to map data in each data store to queryable databases and collections.`,
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
					Description: `Determines whether or not to use S3 tags on the files in the given path as additional partition attributes. See [MongoDB Atlas API](https://docs.mongodb.com/datalake/reference/api/dataLakes-get-all-tenants) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Data lake. ### Data Lake`,
				},
				resource.Attribute{
					Name:        "aws_role_id",
					Description: `Unique identifier of the role that Data Lake can use to access the data stores.`,
				},
				resource.Attribute{
					Name:        "aws_test_s3_bucket",
					Description: `Name of the S3 data bucket that the provided role ID is authorized to access.`,
				},
				resource.Attribute{
					Name:        "data_process_region",
					Description: `The cloud provider region to which Atlas Data Lake routes client connections for data processing.`,
				},
				resource.Attribute{
					Name:        "data_process_region.0.cloud_provider",
					Description: `Name of the cloud service provider.`,
				},
				resource.Attribute{
					Name:        "aws_iam_assumed_role_arn",
					Description: `Amazon Resource Name (ARN) of the IAM Role that Data Lake assumes when accessing S3 Bucket data stores. For more information on S3 actions, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazons3.html).`,
				},
				resource.Attribute{
					Name:        "aws_iam_user_arn",
					Description: `Amazon Resource Name (ARN) of the user that Data Lake assumes when accessing S3 Bucket data stores.`,
				},
				resource.Attribute{
					Name:        "aws_external_id",
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
					Description: `Configuration details for mapping each data store to queryable databases and collections.`,
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
					Description: `Name of a data store to map to the ` + "`" + `<collection>` + "`" + `.`,
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
					Description: `Array of objects where each object represents an [aggregation pipeline](https://docs.mongodb.com/manual/core/aggregation-pipeline/#id1) on a collection.`,
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
					Description: `Each object in the array represents a data store. Data Lake uses the storage.databases configuration details to map data in each data store to queryable databases and collections.`,
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
					Description: `Determines whether or not to use S3 tags on the files in the given path as additional partition attributes. See [MongoDB Atlas API](https://docs.mongodb.com/datalake/reference/api/dataLakes-get-all-tenants) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_database_user",
			Category:         "Data Sources",
			ShortDescription: `Describes a Database User.`,
			Description: `

` + "`" + `mongodbatlas_database_user` + "`" + ` describe a Database User. This represents a database user which will be applied to all clusters within the project.

Each user has a set of roles that provide access to the project’s databases. User's roles apply to all the clusters in the project: if two clusters have a ` + "`" + `products` + "`" + ` database and a user has a role granting ` + "`" + `read` + "`" + ` access on the products database, the user has that access on both clusters.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username for authenticating to MongoDB.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "auth_database_name",
					Description: `(Required) The user’s authentication database. A user must provide both a username and authentication database to log into MongoDB. In Atlas deployments of MongoDB, the authentication database is almost always the admin database, for X509 it is $external. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See [Roles](#roles) below for more details.`,
				},
				resource.Attribute{
					Name:        "x509_type",
					Description: `X.509 method by which the provided username is authenticated.`,
				},
				resource.Attribute{
					Name:        "aws_iam_type",
					Description: `The new database user authenticates with AWS IAM credentials. Default is ` + "`" + `NONE` + "`" + `, ` + "`" + `USER` + "`" + ` means user has AWS IAM user credentials, ` + "`" + `ROLE` + "`" + ` - means user has credentials associated with an AWS IAM role.`,
				},
				resource.Attribute{
					Name:        "ldap_auth_type",
					Description: `Method by which the provided username is authenticated. Default is ` + "`" + `NONE` + "`" + `. Other valid values are: ` + "`" + `USER` + "`" + `, ` + "`" + `GROUP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `Array of clusters and Atlas Data Lakes that this user has access to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster or Atlas Data Lake that the user has access to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of resource that the user has access to. Valid values are: ` + "`" + `CLUSTER` + "`" + ` and ` + "`" + `DATA_LAKE` + "`" + ` ### Roles Block mapping a user's role to a database / collection. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. ->`,
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
					Description: `Collection for which the role applies. You can specify a collection for the ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + ` roles. If you do not specify a collection for ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + `, the role applies to all collections in the database (excluding some collections in the ` + "`" + `system` + "`" + `. database). ### Labels Containing key-value pairs that tag and categorize the database user. Each key and value has a maximum length of 255 characters.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/database-users-get-single-user/) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `List of user’s roles and the databases / collections on which the roles apply. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. See [Roles](#roles) below for more details.`,
				},
				resource.Attribute{
					Name:        "x509_type",
					Description: `X.509 method by which the provided username is authenticated.`,
				},
				resource.Attribute{
					Name:        "aws_iam_type",
					Description: `The new database user authenticates with AWS IAM credentials. Default is ` + "`" + `NONE` + "`" + `, ` + "`" + `USER` + "`" + ` means user has AWS IAM user credentials, ` + "`" + `ROLE` + "`" + ` - means user has credentials associated with an AWS IAM role.`,
				},
				resource.Attribute{
					Name:        "ldap_auth_type",
					Description: `Method by which the provided username is authenticated. Default is ` + "`" + `NONE` + "`" + `. Other valid values are: ` + "`" + `USER` + "`" + `, ` + "`" + `GROUP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `Array of clusters and Atlas Data Lakes that this user has access to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster or Atlas Data Lake that the user has access to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of resource that the user has access to. Valid values are: ` + "`" + `CLUSTER` + "`" + ` and ` + "`" + `DATA_LAKE` + "`" + ` ### Roles Block mapping a user's role to a database / collection. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. ->`,
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
					Description: `Collection for which the role applies. You can specify a collection for the ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + ` roles. If you do not specify a collection for ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + `, the role applies to all collections in the database (excluding some collections in the ` + "`" + `system` + "`" + `. database). ### Labels Containing key-value pairs that tag and categorize the database user. Each key and value has a maximum length of 255 characters.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/database-users-get-single-user/) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_database_users",
			Category:         "Data Sources",
			ShortDescription: `Describes a Database Users.`,
			Description: `

` + "`" + `mongodbatlas_database_users` + "`" + ` describe all Database Users. This represents a database user which will be applied to all clusters within the project.

Each user has a set of roles that provide access to the project’s databases. User's roles apply to all the clusters in the project: if two clusters have a ` + "`" + `products` + "`" + ` database and a user has a role granting ` + "`" + `read` + "`" + ` access on the products database, the user has that access on both clusters.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
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
					Name:        "auth_database_name",
					Description: `(Required) Database against which Atlas authenticates the user. A user must provide both a username and authentication database to log into MongoDB. Possible values include:`,
				},
				resource.Attribute{
					Name:        "ldap_auth_type",
					Description: `is USER or GROUP.`,
				},
				resource.Attribute{
					Name:        "x509_type",
					Description: `X.509 method by which the provided username is authenticated.`,
				},
				resource.Attribute{
					Name:        "aws_iam_type",
					Description: `The new database user authenticates with AWS IAM credentials. Default is ` + "`" + `NONE` + "`" + `, ` + "`" + `USER` + "`" + ` means user has AWS IAM user credentials, ` + "`" + `ROLE` + "`" + ` - means user has credentials associated with an AWS IAM role.`,
				},
				resource.Attribute{
					Name:        "ldap_auth_type",
					Description: `Method by which the provided username is authenticated. Default is ` + "`" + `NONE` + "`" + `. Other valid values are: ` + "`" + `USER` + "`" + `, ` + "`" + `GROUP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `Array of clusters and Atlas Data Lakes that this user has access to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster or Atlas Data Lake that the user has access to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of resource that the user has access to. Valid values are: ` + "`" + `CLUSTER` + "`" + ` and ` + "`" + `DATA_LAKE` + "`" + ` ### Roles Block mapping a user's role to a database / collection. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. ->`,
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
					Description: `Collection for which the role applies. You can specify a collection for the ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + ` roles. If you do not specify a collection for ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + `, the role applies to all collections in the database (excluding some collections in the ` + "`" + `system` + "`" + `. database). ### Labels Containing key-value pairs that tag and categorize the database user. Each key and value has a maximum length of 255 characters.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/database-users-get-single-user/) Documentation for more information.`,
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
					Name:        "auth_database_name",
					Description: `(Required) Database against which Atlas authenticates the user. A user must provide both a username and authentication database to log into MongoDB. Possible values include:`,
				},
				resource.Attribute{
					Name:        "ldap_auth_type",
					Description: `is USER or GROUP.`,
				},
				resource.Attribute{
					Name:        "x509_type",
					Description: `X.509 method by which the provided username is authenticated.`,
				},
				resource.Attribute{
					Name:        "aws_iam_type",
					Description: `The new database user authenticates with AWS IAM credentials. Default is ` + "`" + `NONE` + "`" + `, ` + "`" + `USER` + "`" + ` means user has AWS IAM user credentials, ` + "`" + `ROLE` + "`" + ` - means user has credentials associated with an AWS IAM role.`,
				},
				resource.Attribute{
					Name:        "ldap_auth_type",
					Description: `Method by which the provided username is authenticated. Default is ` + "`" + `NONE` + "`" + `. Other valid values are: ` + "`" + `USER` + "`" + `, ` + "`" + `GROUP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `Array of clusters and Atlas Data Lakes that this user has access to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cluster or Atlas Data Lake that the user has access to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of resource that the user has access to. Valid values are: ` + "`" + `CLUSTER` + "`" + ` and ` + "`" + `DATA_LAKE` + "`" + ` ### Roles Block mapping a user's role to a database / collection. A role allows the user to perform particular actions on the specified database. A role on the admin database can include privileges that apply to the other databases as well. ->`,
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
					Description: `Collection for which the role applies. You can specify a collection for the ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + ` roles. If you do not specify a collection for ` + "`" + `read` + "`" + ` and ` + "`" + `readWrite` + "`" + `, the role applies to all collections in the database (excluding some collections in the ` + "`" + `system` + "`" + `. database). ### Labels Containing key-value pairs that tag and categorize the database user. Each key and value has a maximum length of 255 characters.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key that you want to write.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that you want to write. See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/database-users-get-single-user/) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_event_trigger",
			Category:         "Data Sources",
			ShortDescription: `Describes an Event Trigger.`,
			Description: `

` + "`" + `mongodbatlas_event_trigger` + "`" + ` describe an Event Trigger. 

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
					Name:        "trigger_id",
					Description: `(Required) The unique ID of the trigger. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the trigger.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the trigger.`,
				},
				resource.Attribute{
					Name:        "function_id",
					Description: `The ID of the function associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `The name of the function associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "function_id",
					Description: `The ID of the function associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Status of a trigger.`,
				},
				resource.Attribute{
					Name:        "config_operation_types",
					Description: `The [database event operation types](https://docs.mongodb.com/realm/triggers/database-triggers/#std-label-database-events) to listen for.`,
				},
				resource.Attribute{
					Name:        "config_operation_type",
					Description: `The [authentication operation type](https://docs.mongodb.com/realm/triggers/authentication-triggers/#std-label-authentication-event-operation-types) to listen for.`,
				},
				resource.Attribute{
					Name:        "config_providers",
					Description: `A list of one or more [authentication provider](https://docs.mongodb.com/realm/authentication/providers/) id values. The trigger will only listen for authentication events produced by these providers.`,
				},
				resource.Attribute{
					Name:        "config_database",
					Description: `The name of the MongoDB database that contains the watched collection.`,
				},
				resource.Attribute{
					Name:        "config_collection",
					Description: `The name of the MongoDB collection that the trigger watches for change events.`,
				},
				resource.Attribute{
					Name:        "config_service_id",
					Description: `The ID of the MongoDB Service associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "config_match",
					Description: `A [$match](https://docs.mongodb.com/manual/reference/operator/aggregation/match/) expression document that MongoDB Realm includes in the underlying change stream pipeline for the trigger.`,
				},
				resource.Attribute{
					Name:        "config_project",
					Description: `A [$project](https://docs.mongodb.com/manual/reference/operator/aggregation/project/) expression document that Realm uses to filter the fields that appear in change event objects.`,
				},
				resource.Attribute{
					Name:        "config_full_document",
					Description: `If true, indicates that ` + "`" + `UPDATE` + "`" + ` change events should include the most current [majority-committed](https://docs.mongodb.com/manual/reference/read-concern-majority/) version of the modified document in the fullDocument field.`,
				},
				resource.Attribute{
					Name:        "unordered",
					Description: `Only Available for Database Triggers. If true, event ordering is disabled and this trigger can process events in parallel. If false, event ordering is enabled and the trigger executes serially.`,
				},
				resource.Attribute{
					Name:        "config_schedule",
					Description: `A [cron expression](https://docs.mongodb.com/realm/triggers/cron-expressions/) that defines the trigger schedule.`,
				},
				resource.Attribute{
					Name:        "event_processors",
					Description: `An object where each field name is an event processor ID and each value is an object that configures its corresponding event processor.`,
				},
				resource.Attribute{
					Name:        "event_processors.0.aws_eventbridge.config_account_id",
					Description: `AWS Account ID.`,
				},
				resource.Attribute{
					Name:        "event_processors.0.aws_eventbridge.config_region",
					Description: `Region of AWS Account. See [MongoDB Realm API](https://docs.mongodb.com/realm/admin/api/v3/#get-/groups/%7Bgroupid%7D/apps/%7Bappid%7D/triggers/%7Btriggerid%7D) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the trigger.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the trigger.`,
				},
				resource.Attribute{
					Name:        "function_id",
					Description: `The ID of the function associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `The name of the function associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "function_id",
					Description: `The ID of the function associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Status of a trigger.`,
				},
				resource.Attribute{
					Name:        "config_operation_types",
					Description: `The [database event operation types](https://docs.mongodb.com/realm/triggers/database-triggers/#std-label-database-events) to listen for.`,
				},
				resource.Attribute{
					Name:        "config_operation_type",
					Description: `The [authentication operation type](https://docs.mongodb.com/realm/triggers/authentication-triggers/#std-label-authentication-event-operation-types) to listen for.`,
				},
				resource.Attribute{
					Name:        "config_providers",
					Description: `A list of one or more [authentication provider](https://docs.mongodb.com/realm/authentication/providers/) id values. The trigger will only listen for authentication events produced by these providers.`,
				},
				resource.Attribute{
					Name:        "config_database",
					Description: `The name of the MongoDB database that contains the watched collection.`,
				},
				resource.Attribute{
					Name:        "config_collection",
					Description: `The name of the MongoDB collection that the trigger watches for change events.`,
				},
				resource.Attribute{
					Name:        "config_service_id",
					Description: `The ID of the MongoDB Service associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "config_match",
					Description: `A [$match](https://docs.mongodb.com/manual/reference/operator/aggregation/match/) expression document that MongoDB Realm includes in the underlying change stream pipeline for the trigger.`,
				},
				resource.Attribute{
					Name:        "config_project",
					Description: `A [$project](https://docs.mongodb.com/manual/reference/operator/aggregation/project/) expression document that Realm uses to filter the fields that appear in change event objects.`,
				},
				resource.Attribute{
					Name:        "config_full_document",
					Description: `If true, indicates that ` + "`" + `UPDATE` + "`" + ` change events should include the most current [majority-committed](https://docs.mongodb.com/manual/reference/read-concern-majority/) version of the modified document in the fullDocument field.`,
				},
				resource.Attribute{
					Name:        "unordered",
					Description: `Only Available for Database Triggers. If true, event ordering is disabled and this trigger can process events in parallel. If false, event ordering is enabled and the trigger executes serially.`,
				},
				resource.Attribute{
					Name:        "config_schedule",
					Description: `A [cron expression](https://docs.mongodb.com/realm/triggers/cron-expressions/) that defines the trigger schedule.`,
				},
				resource.Attribute{
					Name:        "event_processors",
					Description: `An object where each field name is an event processor ID and each value is an object that configures its corresponding event processor.`,
				},
				resource.Attribute{
					Name:        "event_processors.0.aws_eventbridge.config_account_id",
					Description: `AWS Account ID.`,
				},
				resource.Attribute{
					Name:        "event_processors.0.aws_eventbridge.config_region",
					Description: `Region of AWS Account. See [MongoDB Realm API](https://docs.mongodb.com/realm/admin/api/v3/#get-/groups/%7Bgroupid%7D/apps/%7Bappid%7D/triggers/%7Btriggerid%7D) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_event_triggers",
			Category:         "Data Sources",
			ShortDescription: `Describes an Event Triggers.`,
			Description: `

` + "`" + `mongodbatlas_event_triggers` + "`" + ` describe all Event Triggers.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to get all event triggers.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) The ObjectID of your application.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Event Trigger. ### Event Trigger`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the trigger.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the trigger. Possible Values: ` + "`" + `DATABASE` + "`" + `, ` + "`" + `AUTHENTICATION` + "`" + ``,
				},
				resource.Attribute{
					Name:        "function_id",
					Description: `The ID of the function associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `The name of the function associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Status of a trigger.`,
				},
				resource.Attribute{
					Name:        "config_operation_types",
					Description: `The [database event operation types](https://docs.mongodb.com/realm/triggers/database-triggers/#std-label-database-events) to listen for.`,
				},
				resource.Attribute{
					Name:        "config_operation_type",
					Description: `The [authentication operation type](https://docs.mongodb.com/realm/triggers/authentication-triggers/#std-label-authentication-event-operation-types) to listen for.`,
				},
				resource.Attribute{
					Name:        "config_providers",
					Description: `A list of one or more [authentication provider](https://docs.mongodb.com/realm/authentication/providers/) id values. The trigger will only listen for authentication events produced by these providers.`,
				},
				resource.Attribute{
					Name:        "config_database",
					Description: `The name of the MongoDB database that contains the watched collection.`,
				},
				resource.Attribute{
					Name:        "config_collection",
					Description: `The name of the MongoDB collection that the trigger watches for change events.`,
				},
				resource.Attribute{
					Name:        "config_service_id",
					Description: `The ID of the MongoDB Service associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "config_match",
					Description: `A [$match](https://docs.mongodb.com/manual/reference/operator/aggregation/match/) expression document that MongoDB Realm includes in the underlying change stream pipeline for the trigger.`,
				},
				resource.Attribute{
					Name:        "config_project",
					Description: `A [$project](https://docs.mongodb.com/manual/reference/operator/aggregation/project/) expression document that Realm uses to filter the fields that appear in change event objects.`,
				},
				resource.Attribute{
					Name:        "config_full_document",
					Description: `If true, indicates that ` + "`" + `UPDATE` + "`" + ` change events should include the most current [majority-committed](https://docs.mongodb.com/manual/reference/read-concern-majority/) version of the modified document in the fullDocument field.`,
				},
				resource.Attribute{
					Name:        "unordered",
					Description: `Sort order for ` + "`" + `DATABASE` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "config_schedule",
					Description: `A [cron expression](https://docs.mongodb.com/realm/triggers/cron-expressions/) that defines the trigger schedule.`,
				},
				resource.Attribute{
					Name:        "event_processors",
					Description: `An object where each field name is an event processor ID and each value is an object that configures its corresponding event processor.`,
				},
				resource.Attribute{
					Name:        "event_processors.0.aws_eventbridge.config_account_id",
					Description: `AWS Account ID.`,
				},
				resource.Attribute{
					Name:        "event_processors.0.aws_eventbridge.config_region",
					Description: `Region of AWS Account. See [MongoDB Realm API](https://docs.mongodb.com/realm/admin/api/v3/#get-/groups/%7Bgroupid%7D/apps/%7Bappid%7D/triggers) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Event Trigger. ### Event Trigger`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the trigger.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the trigger. Possible Values: ` + "`" + `DATABASE` + "`" + `, ` + "`" + `AUTHENTICATION` + "`" + ``,
				},
				resource.Attribute{
					Name:        "function_id",
					Description: `The ID of the function associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `The name of the function associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Status of a trigger.`,
				},
				resource.Attribute{
					Name:        "config_operation_types",
					Description: `The [database event operation types](https://docs.mongodb.com/realm/triggers/database-triggers/#std-label-database-events) to listen for.`,
				},
				resource.Attribute{
					Name:        "config_operation_type",
					Description: `The [authentication operation type](https://docs.mongodb.com/realm/triggers/authentication-triggers/#std-label-authentication-event-operation-types) to listen for.`,
				},
				resource.Attribute{
					Name:        "config_providers",
					Description: `A list of one or more [authentication provider](https://docs.mongodb.com/realm/authentication/providers/) id values. The trigger will only listen for authentication events produced by these providers.`,
				},
				resource.Attribute{
					Name:        "config_database",
					Description: `The name of the MongoDB database that contains the watched collection.`,
				},
				resource.Attribute{
					Name:        "config_collection",
					Description: `The name of the MongoDB collection that the trigger watches for change events.`,
				},
				resource.Attribute{
					Name:        "config_service_id",
					Description: `The ID of the MongoDB Service associated with the trigger.`,
				},
				resource.Attribute{
					Name:        "config_match",
					Description: `A [$match](https://docs.mongodb.com/manual/reference/operator/aggregation/match/) expression document that MongoDB Realm includes in the underlying change stream pipeline for the trigger.`,
				},
				resource.Attribute{
					Name:        "config_project",
					Description: `A [$project](https://docs.mongodb.com/manual/reference/operator/aggregation/project/) expression document that Realm uses to filter the fields that appear in change event objects.`,
				},
				resource.Attribute{
					Name:        "config_full_document",
					Description: `If true, indicates that ` + "`" + `UPDATE` + "`" + ` change events should include the most current [majority-committed](https://docs.mongodb.com/manual/reference/read-concern-majority/) version of the modified document in the fullDocument field.`,
				},
				resource.Attribute{
					Name:        "unordered",
					Description: `Sort order for ` + "`" + `DATABASE` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "config_schedule",
					Description: `A [cron expression](https://docs.mongodb.com/realm/triggers/cron-expressions/) that defines the trigger schedule.`,
				},
				resource.Attribute{
					Name:        "event_processors",
					Description: `An object where each field name is an event processor ID and each value is an object that configures its corresponding event processor.`,
				},
				resource.Attribute{
					Name:        "event_processors.0.aws_eventbridge.config_account_id",
					Description: `AWS Account ID.`,
				},
				resource.Attribute{
					Name:        "event_processors.0.aws_eventbridge.config_region",
					Description: `Region of AWS Account. See [MongoDB Realm API](https://docs.mongodb.com/realm/admin/api/v3/#get-/groups/%7Bgroupid%7D/apps/%7Bappid%7D/triggers) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_federated_settings",
			Category:         "Data Sources",
			ShortDescription: `Provides a federated settings data source.`,
			Description: `

` + "`" + `mongodbatlas_federated_settings` + "`" + ` provides a federated settings data source. Atlas Cloud federated settings provides federated settings outputs.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. ## Attributes Reference In addition to all arguments above, the following attributes are exported: ### FederatedSettings`,
				},
				resource.Attribute{
					Name:        "federated_domains",
					Description: `List that contains the domains associated with the organization's identity provider.`,
				},
				resource.Attribute{
					Name:        "has_role_mappings",
					Description: `Flag that indicates whether this organization has role mappings configured.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this federation.`,
				},
				resource.Attribute{
					Name:        "identity_provider_id",
					Description: `Unique 20-hexadecimal digit string that identifies the identity provider connected to this organization.`,
				},
				resource.Attribute{
					Name:        "identity_provider_status",
					Description: `Value that indicates whether the identity provider is active. Atlas returns ACTIVE if the identity provider is active and INACTIVE if the identity provider is inactive. For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "federated_domains",
					Description: `List that contains the domains associated with the organization's identity provider.`,
				},
				resource.Attribute{
					Name:        "has_role_mappings",
					Description: `Flag that indicates whether this organization has role mappings configured.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this federation.`,
				},
				resource.Attribute{
					Name:        "identity_provider_id",
					Description: `Unique 20-hexadecimal digit string that identifies the identity provider connected to this organization.`,
				},
				resource.Attribute{
					Name:        "identity_provider_status",
					Description: `Value that indicates whether the identity provider is active. Atlas returns ACTIVE if the identity provider is active and INACTIVE if the identity provider is inactive. For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_federated_settings_identity_provider",
			Category:         "Data Sources",
			ShortDescription: `Provides a federated settings Organization identity provider data source.`,
			Description: `

` + "`" + `mongodbatlas_federated_settings_identity_provider` + "`" + ` provides a federated settings identity provider data source. Atlas federated settings identity provider provides federated settings outputs for the configured identity provider.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "identity_provider_id",
					Description: `(Required) Unique 20-hexadecimal digit string that identifies the IdP. ## Attributes Reference In addition to all arguments above, the following attributes are exported: ### FederatedSettingsIdentityProvider`,
				},
				resource.Attribute{
					Name:        "acs_url",
					Description: `Assertion consumer service URL to which the IdP sends the SAML response.`,
				},
				resource.Attribute{
					Name:        "associated_domains",
					Description: `List that contains the configured domains from which users can log in for this IdP.`,
				},
				resource.Attribute{
					Name:        "associated_orgs",
					Description: `List that contains the organizations from which users can log in for this IdP.`,
				},
				resource.Attribute{
					Name:        "domain_allow_list",
					Description: `List that contains the approved domains from which organization users can log in.`,
				},
				resource.Attribute{
					Name:        "domain_restriction_enabled",
					Description: `Flag that indicates whether domain restriction is enabled for the connected organization.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects.`,
				},
				resource.Attribute{
					Name:        "post_auth_role_grants",
					Description: `List that contains the default roles granted to users who authenticate through the IdP in a connected organization. If you provide a postAuthRoleGrants field in the request, the array that you provide replaces the current postAuthRoleGrants. ### Role_mappings`,
				},
				resource.Attribute{
					Name:        "external_group_name",
					Description: `Unique human-readable label that identifies the identity provider group to which this role mapping applies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this role mapping.`,
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
					Name:        "role",
					Description: `Specifies the Role that is attached to the Role Mapping. ### User Conflicts`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `Email address of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `First name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `Last name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Name of the Atlas user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "audience_uri",
					Description: `Identifier for the intended audience of the SAML Assertion.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Human-readable label that identifies the IdP.`,
				},
				resource.Attribute{
					Name:        "issuer_uri",
					Description: `Identifier for the issuer of the SAML Assertion.`,
				},
				resource.Attribute{
					Name:        "okta_idp_id",
					Description: `Unique 20-hexadecimal digit string that identifies the IdP. ### Pem File Info - List that contains the file information, including: start date, and expiration date for the identity provider's PEM-encoded public key certificate.`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `Expiration Date.`,
				},
				resource.Attribute{
					Name:        "not_before",
					Description: `Start Date.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `Filename of certificate`,
				},
				resource.Attribute{
					Name:        "request_binding",
					Description: `SAML Authentication Request Protocol binding used to send the AuthNRequest. Atlas supports the following binding values: - HTTP POST - HTTP REDIRECT`,
				},
				resource.Attribute{
					Name:        "response_signature_algorithm",
					Description: `Algorithm used to encrypt the IdP signature. Atlas supports the following signature algorithm values: - SHA-1 - SHA-256`,
				},
				resource.Attribute{
					Name:        "sso_debug_enabled",
					Description: `Flag that indicates whether the IdP has enabled Bypass SAML Mode. Enabling this mode generates a URL that allows you bypass SAML and login to your organizations at any point. You can authenticate with this special URL only when Bypass Mode is enabled. Set this parameter to true during testing. This keeps you from getting locked out of MongoDB.`,
				},
				resource.Attribute{
					Name:        "sso_url",
					Description: `URL of the receiver of the SAML AuthNRequest.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Label that indicates whether the identity provider is active. The IdP is Inactive until you map at least one domain to the IdP. For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "acs_url",
					Description: `Assertion consumer service URL to which the IdP sends the SAML response.`,
				},
				resource.Attribute{
					Name:        "associated_domains",
					Description: `List that contains the configured domains from which users can log in for this IdP.`,
				},
				resource.Attribute{
					Name:        "associated_orgs",
					Description: `List that contains the organizations from which users can log in for this IdP.`,
				},
				resource.Attribute{
					Name:        "domain_allow_list",
					Description: `List that contains the approved domains from which organization users can log in.`,
				},
				resource.Attribute{
					Name:        "domain_restriction_enabled",
					Description: `Flag that indicates whether domain restriction is enabled for the connected organization.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects.`,
				},
				resource.Attribute{
					Name:        "post_auth_role_grants",
					Description: `List that contains the default roles granted to users who authenticate through the IdP in a connected organization. If you provide a postAuthRoleGrants field in the request, the array that you provide replaces the current postAuthRoleGrants. ### Role_mappings`,
				},
				resource.Attribute{
					Name:        "external_group_name",
					Description: `Unique human-readable label that identifies the identity provider group to which this role mapping applies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this role mapping.`,
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
					Name:        "role",
					Description: `Specifies the Role that is attached to the Role Mapping. ### User Conflicts`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `Email address of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `First name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `Last name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Name of the Atlas user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "audience_uri",
					Description: `Identifier for the intended audience of the SAML Assertion.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Human-readable label that identifies the IdP.`,
				},
				resource.Attribute{
					Name:        "issuer_uri",
					Description: `Identifier for the issuer of the SAML Assertion.`,
				},
				resource.Attribute{
					Name:        "okta_idp_id",
					Description: `Unique 20-hexadecimal digit string that identifies the IdP. ### Pem File Info - List that contains the file information, including: start date, and expiration date for the identity provider's PEM-encoded public key certificate.`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `Expiration Date.`,
				},
				resource.Attribute{
					Name:        "not_before",
					Description: `Start Date.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `Filename of certificate`,
				},
				resource.Attribute{
					Name:        "request_binding",
					Description: `SAML Authentication Request Protocol binding used to send the AuthNRequest. Atlas supports the following binding values: - HTTP POST - HTTP REDIRECT`,
				},
				resource.Attribute{
					Name:        "response_signature_algorithm",
					Description: `Algorithm used to encrypt the IdP signature. Atlas supports the following signature algorithm values: - SHA-1 - SHA-256`,
				},
				resource.Attribute{
					Name:        "sso_debug_enabled",
					Description: `Flag that indicates whether the IdP has enabled Bypass SAML Mode. Enabling this mode generates a URL that allows you bypass SAML and login to your organizations at any point. You can authenticate with this special URL only when Bypass Mode is enabled. Set this parameter to true during testing. This keeps you from getting locked out of MongoDB.`,
				},
				resource.Attribute{
					Name:        "sso_url",
					Description: `URL of the receiver of the SAML AuthNRequest.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Label that indicates whether the identity provider is active. The IdP is Inactive until you map at least one domain to the IdP. For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_federated_settings_identity_providers",
			Category:         "Data Sources",
			ShortDescription: `Provides a federated settings Organization Identity Provider datasource.`,
			Description: `

` + "`" + `mongodbatlas_federated_settings_identity_providers` + "`" + ` provides an Federated Settings Identity Providers datasource. Atlas Cloud Federated Settings Identity Providers provides federated settings outputs for the configured Identity Providers.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `Includes cloudProviderSnapshot object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### FederatedSettingsIdentityProvider`,
				},
				resource.Attribute{
					Name:        "identity_provider_id",
					Description: `Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "acs_url",
					Description: `Assertion consumer service URL to which the IdP sends the SAML response.`,
				},
				resource.Attribute{
					Name:        "associated_domains",
					Description: `List that contains the configured domains from which users can log in for this IdP.`,
				},
				resource.Attribute{
					Name:        "associated_orgs",
					Description: `List that contains the configured domains from which users can log in for this IdP.`,
				},
				resource.Attribute{
					Name:        "domain_allow_list",
					Description: `List that contains the approved domains from which organization users can log in.`,
				},
				resource.Attribute{
					Name:        "domain_restriction_enabled",
					Description: `Flag that indicates whether domain restriction is enabled for the connected organization.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects.`,
				},
				resource.Attribute{
					Name:        "post_auth_role_grants",
					Description: `List that contains the default roles granted to users who authenticate through the IdP in a connected organization. If you provide a postAuthRoleGrants field in the request, the array that you provide replaces the current postAuthRoleGrants. ### Role_mappings`,
				},
				resource.Attribute{
					Name:        "external_group_name",
					Description: `Unique human-readable label that identifies the identity provider group to which this role mapping applies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this role mapping.`,
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
					Name:        "role",
					Description: `Specifies the Role that is attached to the Role Mapping. ### User Conflicts`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `Email address of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `First name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `Last name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Name of the Atlas user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "audience_uri",
					Description: `Identifier for the intended audience of the SAML Assertion.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Human-readable label that identifies the IdP.`,
				},
				resource.Attribute{
					Name:        "issuer_uri",
					Description: `Identifier for the issuer of the SAML Assertion.`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `Unique 20-hexadecimal digit string that identifies the IdP. ### Pem File Info - List that contains the file information, including: start date, and expiration date for the identity provider's PEM-encoded public key certificate.`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `Expiration Date.`,
				},
				resource.Attribute{
					Name:        "not_before",
					Description: `Start Date.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `Filename of certificate`,
				},
				resource.Attribute{
					Name:        "request_binding",
					Description: `SAML Authentication Request Protocol binding used to send the AuthNRequest. Atlas supports the following binding values: - HTTP POST - HTTP REDIRECT`,
				},
				resource.Attribute{
					Name:        "response_signature_algorithm",
					Description: `Algorithm used to encrypt the IdP signature. Atlas supports the following signature algorithm values: - SHA-1 - SHA-256`,
				},
				resource.Attribute{
					Name:        "sso_debug_enabled",
					Description: `Flag that indicates whether the IdP has enabled Bypass SAML Mode. Enabling this mode generates a URL that allows you bypass SAML and login to your organizations at any point. You can authenticate with this special URL only when Bypass Mode is enabled. Set this parameter to true during testing. This keeps you from getting locked out of MongoDB.`,
				},
				resource.Attribute{
					Name:        "sso_url",
					Description: `URL of the receiver of the SAML AuthNRequest.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Label that indicates whether the identity provider is active. The IdP is Inactive until you map at least one domain to the IdP. For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "results",
					Description: `Includes cloudProviderSnapshot object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### FederatedSettingsIdentityProvider`,
				},
				resource.Attribute{
					Name:        "identity_provider_id",
					Description: `Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "acs_url",
					Description: `Assertion consumer service URL to which the IdP sends the SAML response.`,
				},
				resource.Attribute{
					Name:        "associated_domains",
					Description: `List that contains the configured domains from which users can log in for this IdP.`,
				},
				resource.Attribute{
					Name:        "associated_orgs",
					Description: `List that contains the configured domains from which users can log in for this IdP.`,
				},
				resource.Attribute{
					Name:        "domain_allow_list",
					Description: `List that contains the approved domains from which organization users can log in.`,
				},
				resource.Attribute{
					Name:        "domain_restriction_enabled",
					Description: `Flag that indicates whether domain restriction is enabled for the connected organization.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects.`,
				},
				resource.Attribute{
					Name:        "post_auth_role_grants",
					Description: `List that contains the default roles granted to users who authenticate through the IdP in a connected organization. If you provide a postAuthRoleGrants field in the request, the array that you provide replaces the current postAuthRoleGrants. ### Role_mappings`,
				},
				resource.Attribute{
					Name:        "external_group_name",
					Description: `Unique human-readable label that identifies the identity provider group to which this role mapping applies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this role mapping.`,
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
					Name:        "role",
					Description: `Specifies the Role that is attached to the Role Mapping. ### User Conflicts`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `Email address of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `First name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `Last name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Name of the Atlas user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "audience_uri",
					Description: `Identifier for the intended audience of the SAML Assertion.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Human-readable label that identifies the IdP.`,
				},
				resource.Attribute{
					Name:        "issuer_uri",
					Description: `Identifier for the issuer of the SAML Assertion.`,
				},
				resource.Attribute{
					Name:        "idp_id",
					Description: `Unique 20-hexadecimal digit string that identifies the IdP. ### Pem File Info - List that contains the file information, including: start date, and expiration date for the identity provider's PEM-encoded public key certificate.`,
				},
				resource.Attribute{
					Name:        "not_after",
					Description: `Expiration Date.`,
				},
				resource.Attribute{
					Name:        "not_before",
					Description: `Start Date.`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `Filename of certificate`,
				},
				resource.Attribute{
					Name:        "request_binding",
					Description: `SAML Authentication Request Protocol binding used to send the AuthNRequest. Atlas supports the following binding values: - HTTP POST - HTTP REDIRECT`,
				},
				resource.Attribute{
					Name:        "response_signature_algorithm",
					Description: `Algorithm used to encrypt the IdP signature. Atlas supports the following signature algorithm values: - SHA-1 - SHA-256`,
				},
				resource.Attribute{
					Name:        "sso_debug_enabled",
					Description: `Flag that indicates whether the IdP has enabled Bypass SAML Mode. Enabling this mode generates a URL that allows you bypass SAML and login to your organizations at any point. You can authenticate with this special URL only when Bypass Mode is enabled. Set this parameter to true during testing. This keeps you from getting locked out of MongoDB.`,
				},
				resource.Attribute{
					Name:        "sso_url",
					Description: `URL of the receiver of the SAML AuthNRequest.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Label that indicates whether the identity provider is active. The IdP is Inactive until you map at least one domain to the IdP. For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_federated_settings_org_config",
			Category:         "Data Sources",
			ShortDescription: `Provides a federated settings Organization Configuration.`,
			Description: `

` + "`" + `mongodbatlas_federated_settings_org_config` + "`" + ` provides an Federated Settings Identity Providers datasource. Atlas Cloud Federated Settings Organizational configuration provides federated settings outputs for the configured Organizational configuration.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. ## Attributes Reference In addition to all arguments above, the following attributes are exported: ### FederatedSettingsOrgConfig`,
				},
				resource.Attribute{
					Name:        "domain_allow_list",
					Description: `List that contains the approved domains from which organization users can log in. Note: If the organization uses an identity provider, ` + "`" + `domain_allow_list` + "`" + ` includes: any SSO domains associated with organization's identity provider and any custom domains associated with the specific organization.`,
				},
				resource.Attribute{
					Name:        "domain_restriction_enabled",
					Description: `Flag that indicates whether domain restriction is enabled for the connected organization. User Conflicts returns null when ` + "`" + `domain_restriction_enabled` + "`" + ` is false.`,
				},
				resource.Attribute{
					Name:        "identity_provider_id",
					Description: `Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "post_auth_role_grants",
					Description: `List that contains the default [roles](https://www.mongodb.com/docs/atlas/reference/user-roles/#std-label-organization-roles) granted to users who authenticate through the IdP in a connected organization. ### Role_mappings`,
				},
				resource.Attribute{
					Name:        "external_group_name",
					Description: `Unique human-readable label that identifies the identity provider group to which this role mapping applies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this role mapping.`,
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
					Name:        "role",
					Description: `Specifies the Role that is attached to the Role Mapping. ### User Conflicts`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `Email address of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `First name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `Last name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Name of the Atlas user that conflicts with selected domains. For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_allow_list",
					Description: `List that contains the approved domains from which organization users can log in. Note: If the organization uses an identity provider, ` + "`" + `domain_allow_list` + "`" + ` includes: any SSO domains associated with organization's identity provider and any custom domains associated with the specific organization.`,
				},
				resource.Attribute{
					Name:        "domain_restriction_enabled",
					Description: `Flag that indicates whether domain restriction is enabled for the connected organization. User Conflicts returns null when ` + "`" + `domain_restriction_enabled` + "`" + ` is false.`,
				},
				resource.Attribute{
					Name:        "identity_provider_id",
					Description: `Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "post_auth_role_grants",
					Description: `List that contains the default [roles](https://www.mongodb.com/docs/atlas/reference/user-roles/#std-label-organization-roles) granted to users who authenticate through the IdP in a connected organization. ### Role_mappings`,
				},
				resource.Attribute{
					Name:        "external_group_name",
					Description: `Unique human-readable label that identifies the identity provider group to which this role mapping applies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this role mapping.`,
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
					Name:        "role",
					Description: `Specifies the Role that is attached to the Role Mapping. ### User Conflicts`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `Email address of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `First name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `Last name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Name of the Atlas user that conflicts with selected domains. For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_federated_settings_org_configs",
			Category:         "Data Sources",
			ShortDescription: `Provides a federated settings Organization Configurations.`,
			Description: `

` + "`" + `mongodbatlas_federated_settings_org_configs` + "`" + ` provides an Federated Settings Identity Providers datasource. Atlas Cloud Federated Settings Identity Providers provides federated settings outputs for the configured Identity Providers.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `Includes cloudProviderSnapshot object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### FederatedSettingsOrgConfigs`,
				},
				resource.Attribute{
					Name:        "domain_allow_list",
					Description: `List that contains the approved domains from which organization users can log in.`,
				},
				resource.Attribute{
					Name:        "domain_restriction_enabled",
					Description: `Flag that indicates whether domain restriction is enabled for the connected organization.`,
				},
				resource.Attribute{
					Name:        "identity_provider_id",
					Description: `Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects.`,
				},
				resource.Attribute{
					Name:        "post_auth_role_grants",
					Description: `List that contains the default roles granted to users who authenticate through the IdP in a connected organization. ### Role_mappings`,
				},
				resource.Attribute{
					Name:        "external_group_name",
					Description: `Unique human-readable label that identifies the identity provider group to which this role mapping applies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this role mapping.`,
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
					Name:        "role",
					Description: `Specifies the Role that is attached to the Role Mapping. ### User Conflicts`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `Email address of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `First name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `Last name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Name of the Atlas user that conflicts with selected domains. For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "results",
					Description: `Includes cloudProviderSnapshot object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### FederatedSettingsOrgConfigs`,
				},
				resource.Attribute{
					Name:        "domain_allow_list",
					Description: `List that contains the approved domains from which organization users can log in.`,
				},
				resource.Attribute{
					Name:        "domain_restriction_enabled",
					Description: `Flag that indicates whether domain restriction is enabled for the connected organization.`,
				},
				resource.Attribute{
					Name:        "identity_provider_id",
					Description: `Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects.`,
				},
				resource.Attribute{
					Name:        "post_auth_role_grants",
					Description: `List that contains the default roles granted to users who authenticate through the IdP in a connected organization. ### Role_mappings`,
				},
				resource.Attribute{
					Name:        "external_group_name",
					Description: `Unique human-readable label that identifies the identity provider group to which this role mapping applies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this role mapping.`,
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
					Name:        "role",
					Description: `Specifies the Role that is attached to the Role Mapping. ### User Conflicts`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `Email address of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `First name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `Last name of the the user that conflicts with selected domains.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Name of the Atlas user that conflicts with selected domains. For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_federated_settings_org_role_mapping",
			Category:         "Data Sources",
			ShortDescription: `Provides a federated settings Role Mapping datasource.`,
			Description: `

` + "`" + `mongodbatlas_federated_settings_org_role_mapping` + "`" + ` provides an Federated Settings Org Role Mapping datasource. Atlas Cloud Federated Settings Org Role Mapping provides federated settings outputs for the configured Org Role Mapping.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "federation_settings_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the federated authentication configuration.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. ## Attributes Reference In addition to all arguments above, the following attributes are exported: ### FederatedSettingsOrgRoleMappings`,
				},
				resource.Attribute{
					Name:        "external_group_name",
					Description: `Unique human-readable label that identifies the identity provider group to which this role mapping applies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this role mapping.`,
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
					Name:        "role",
					Description: `Specifies the Role that is attached to the Role Mapping. For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "external_group_name",
					Description: `Unique human-readable label that identifies the identity provider group to which this role mapping applies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this role mapping.`,
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
					Name:        "role",
					Description: `Specifies the Role that is attached to the Role Mapping. For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_federated_settings_org_role_mappings",
			Category:         "Data Sources",
			ShortDescription: `Provides a federated settings Role Mapping datasource.`,
			Description: `

` + "`" + `mongodbatlas_federated_settings_org_role_mappings` + "`" + ` provides an Federated Settings Org Role Mapping datasource. Atlas Cloud Federated Settings Org Role Mapping provides federated settings outputs for the configured Org Role Mapping.


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
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `Includes cloudProviderSnapshot object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### FederatedSettingsOrgRoleMappings`,
				},
				resource.Attribute{
					Name:        "external_group_name",
					Description: `Unique human-readable label that identifies the identity provider group to which this role mapping applies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this role mapping.`,
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
					Name:        "role",
					Description: `Specifies the Role that is attached to the Role Mapping. For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "results",
					Description: `Includes cloudProviderSnapshot object for each item detailed in the results array section.`,
				},
				resource.Attribute{
					Name:        "totalCount",
					Description: `Count of the total number of items in the result set. It may be greater than the number of objects in the results array if the entire result set is paginated. ### FederatedSettingsOrgRoleMappings`,
				},
				resource.Attribute{
					Name:        "external_group_name",
					Description: `Unique human-readable label that identifies the identity provider group to which this role mapping applies.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies this role mapping.`,
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
					Name:        "role",
					Description: `Specifies the Role that is attached to the Role Mapping. For more information see: [MongoDB Atlas API Reference.](https://www.mongodb.com/docs/atlas/reference/api/federation-configuration/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_global_cluster_config",
			Category:         "Data Sources",
			ShortDescription: `Describes the Global Cluster Configuration.`,
			Description: `

` + "`" + `mongodbatlas_global_cluster_config` + "`" + ` describes all managed namespaces and custom zone mappings associated with the specified Global Cluster.


-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "custom_zone_mapping",
					Description: `A map of all custom zone mappings defined for the Global Cluster. Atlas automatically maps each location code to the closest geographical zone. Custom zone mappings allow administrators to override these automatic mappings. If your Global Cluster does not have any custom zone mappings, this document is empty.`,
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
					Description: `Specifies whether the custom shard key for the collection is [hashed](https://docs.mongodb.com/manual/reference/method/sh.shardCollection/#hashed-shard-keys). If omitted, defaults to ` + "`" + `false` + "`" + `. If ` + "`" + `false` + "`" + `, Atlas uses [ranged sharding](https://docs.mongodb.com/manual/core/ranged-sharding/). This is only available for Atlas clusters with MongoDB v4.4 and later.`,
				},
				resource.Attribute{
					Name:        "is_shard_key_unique",
					Description: `Specifies whether the underlying index enforces a unique constraint. If omitted, defaults to false. You cannot specify true when using [hashed shard keys](https://docs.mongodb.com/manual/core/hashed-sharding/#std-label-sharding-hashed). See detailed information for arguments and attributes: [MongoDB API Global Clusters](https://docs.atlas.mongodb.com/reference/api/global-clusters/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "custom_zone_mapping",
					Description: `A map of all custom zone mappings defined for the Global Cluster. Atlas automatically maps each location code to the closest geographical zone. Custom zone mappings allow administrators to override these automatic mappings. If your Global Cluster does not have any custom zone mappings, this document is empty.`,
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
					Description: `Specifies whether the custom shard key for the collection is [hashed](https://docs.mongodb.com/manual/reference/method/sh.shardCollection/#hashed-shard-keys). If omitted, defaults to ` + "`" + `false` + "`" + `. If ` + "`" + `false` + "`" + `, Atlas uses [ranged sharding](https://docs.mongodb.com/manual/core/ranged-sharding/). This is only available for Atlas clusters with MongoDB v4.4 and later.`,
				},
				resource.Attribute{
					Name:        "is_shard_key_unique",
					Description: `Specifies whether the underlying index enforces a unique constraint. If omitted, defaults to false. You cannot specify true when using [hashed shard keys](https://docs.mongodb.com/manual/core/hashed-sharding/#std-label-sharding-hashed). See detailed information for arguments and attributes: [MongoDB API Global Clusters](https://docs.atlas.mongodb.com/reference/api/global-clusters/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_ldap_configuration",
			Category:         "Data Sources",
			ShortDescription: `Describes a LDAP Configuration.`,
			Description: `

` + "`" + `mongodbatlas_ldap_configuration` + "`" + ` describes a LDAP Configuration.

-> **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Identifier for the Atlas project associated with the LDAP over TLS/SSL configuration. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "authentication_enabled",
					Description: `Specifies whether user authentication with LDAP is enabled.`,
				},
				resource.Attribute{
					Name:        "authorization_enabled",
					Description: `Specifies whether user authorization with LDAP is enabled.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The hostname or IP address of the LDAP server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port to which the LDAP server listens for client connections.`,
				},
				resource.Attribute{
					Name:        "bind_username",
					Description: `The user DN that Atlas uses to connect to the LDAP server.`,
				},
				resource.Attribute{
					Name:        "bind_password",
					Description: `The password used to authenticate the ` + "`" + `bind_username` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ca_certificate",
					Description: `CA certificate used to verify the identify of the LDAP server.`,
				},
				resource.Attribute{
					Name:        "authz_query_template",
					Description: `An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs.`,
				},
				resource.Attribute{
					Name:        "user_to_dn_mapping",
					Description: `Maps an LDAP username for authentication to an LDAP Distinguished Name (DN).`,
				},
				resource.Attribute{
					Name:        "user_to_dn_mapping.0.match",
					Description: `A regular expression to match against a provided LDAP username.`,
				},
				resource.Attribute{
					Name:        "user_to_dn_mapping.0.substitution",
					Description: `An LDAP Distinguished Name (DN) formatting template that converts the LDAP name matched by the ` + "`" + `match` + "`" + ` regular expression into an LDAP Distinguished Name.`,
				},
				resource.Attribute{
					Name:        "user_to_dn_mapping.0.ldap_query",
					Description: `An LDAP query formatting template that inserts the LDAP name matched by the ` + "`" + `match` + "`" + ` regular expression into an LDAP query URI as specified by RFC 4515 and RFC 4516. See detailed information for arguments and attributes: [MongoDB API LDAP Configuration](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-get-current)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "authentication_enabled",
					Description: `Specifies whether user authentication with LDAP is enabled.`,
				},
				resource.Attribute{
					Name:        "authorization_enabled",
					Description: `Specifies whether user authorization with LDAP is enabled.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The hostname or IP address of the LDAP server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port to which the LDAP server listens for client connections.`,
				},
				resource.Attribute{
					Name:        "bind_username",
					Description: `The user DN that Atlas uses to connect to the LDAP server.`,
				},
				resource.Attribute{
					Name:        "bind_password",
					Description: `The password used to authenticate the ` + "`" + `bind_username` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ca_certificate",
					Description: `CA certificate used to verify the identify of the LDAP server.`,
				},
				resource.Attribute{
					Name:        "authz_query_template",
					Description: `An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs.`,
				},
				resource.Attribute{
					Name:        "user_to_dn_mapping",
					Description: `Maps an LDAP username for authentication to an LDAP Distinguished Name (DN).`,
				},
				resource.Attribute{
					Name:        "user_to_dn_mapping.0.match",
					Description: `A regular expression to match against a provided LDAP username.`,
				},
				resource.Attribute{
					Name:        "user_to_dn_mapping.0.substitution",
					Description: `An LDAP Distinguished Name (DN) formatting template that converts the LDAP name matched by the ` + "`" + `match` + "`" + ` regular expression into an LDAP Distinguished Name.`,
				},
				resource.Attribute{
					Name:        "user_to_dn_mapping.0.ldap_query",
					Description: `An LDAP query formatting template that inserts the LDAP name matched by the ` + "`" + `match` + "`" + ` regular expression into an LDAP query URI as specified by RFC 4515 and RFC 4516. See detailed information for arguments and attributes: [MongoDB API LDAP Configuration](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-get-current)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_ldap_verify",
			Category:         "Data Sources",
			ShortDescription: `Describes a LDAP Verify.`,
			Description: `

` + "`" + `mongodbatlas_ldap_verify` + "`" + ` describes a LDAP Verify.

-> **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique identifier for the Atlas project associated with the verification request.`,
				},
				resource.Attribute{
					Name:        "request_id",
					Description: `(Required) Unique identifier of a request to verify an LDAP configuration. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The hostname or IP address of the LDAP server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `LDAP ConfigurationThe port to which the LDAP server listens for client connections.`,
				},
				resource.Attribute{
					Name:        "bind_username",
					Description: `The user DN that Atlas uses to connect to the LDAP server.`,
				},
				resource.Attribute{
					Name:        "bind_password",
					Description: `The password used to authenticate the ` + "`" + `bind_username` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ca_certificate",
					Description: `LDAP ConfigurationCA certificate used to verify the identify of the LDAP server.`,
				},
				resource.Attribute{
					Name:        "authz_query_template",
					Description: `LDAP ConfigurationAn LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs.`,
				},
				resource.Attribute{
					Name:        "request_id",
					Description: `The unique identifier for the request to verify the LDAP over TLS/SSL configuration.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the LDAP over TLS/SSL configuration.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `One or more links to sub-resources. The relations in the URLs are explained in the Web Linking Specification.`,
				},
				resource.Attribute{
					Name:        "validations",
					Description: `Array of validation messages related to the verification of the provided LDAP over TLS/SSL configuration details. See detailed information for arguments and attributes: [MongoDB API LDAP Verify](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-verification-status)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The hostname or IP address of the LDAP server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `LDAP ConfigurationThe port to which the LDAP server listens for client connections.`,
				},
				resource.Attribute{
					Name:        "bind_username",
					Description: `The user DN that Atlas uses to connect to the LDAP server.`,
				},
				resource.Attribute{
					Name:        "bind_password",
					Description: `The password used to authenticate the ` + "`" + `bind_username` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ca_certificate",
					Description: `LDAP ConfigurationCA certificate used to verify the identify of the LDAP server.`,
				},
				resource.Attribute{
					Name:        "authz_query_template",
					Description: `LDAP ConfigurationAn LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs.`,
				},
				resource.Attribute{
					Name:        "request_id",
					Description: `The unique identifier for the request to verify the LDAP over TLS/SSL configuration.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the LDAP over TLS/SSL configuration.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `One or more links to sub-resources. The relations in the URLs are explained in the Web Linking Specification.`,
				},
				resource.Attribute{
					Name:        "validations",
					Description: `Array of validation messages related to the verification of the provided LDAP over TLS/SSL configuration details. See detailed information for arguments and attributes: [MongoDB API LDAP Verify](https://docs.atlas.mongodb.com/reference/api/ldaps-configuration-verification-status)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_maintenance_window",
			Category:         "Data Sources",
			ShortDescription: `Provides a Maintenance Window Datasource.`,
			Description: `

` + "`" + `mongodbatlas_maintenance_window` + "`" + ` provides a Maintenance Window entry datasource. Gets information regarding the configured maintenance window for a MongoDB Atlas project.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `The unique identifier of the project for the Maintenance Window. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "auto_defer_once_enabled",
					Description: `Flag that indicates whether you want to defer all maintenance windows one week they would be triggered. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/maintenance-windows/)`,
				},
			},
			Attributes: []resource.Attribute{
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
					Name:        "auto_defer_once_enabled",
					Description: `Flag that indicates whether you want to defer all maintenance windows one week they would be triggered. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/maintenance-windows/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_container",
			Category:         "Data Sources",
			ShortDescription: `Describes a Cluster resource.`,
			Description: `

` + "`" + `mongodbatlas_network_container` + "`" + ` describes a Network Peering Container. The resource requires your Project ID and container ID.

~> **IMPORTANT:** This resource creates one Network Peering container into which Atlas can deploy Network Peering connections. An Atlas project can have a maximum of one container for each cloud provider. You must have either the Project Owner or Organization Owner role to successfully call this endpoint.

-> **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.


`,
			Keywords: []string{},
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
					Description: `The Atlas AWS region name for where this container will exist.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The Atlas Azure region name for where this container will exist.`,
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
					Description: `The name of the Azure VNet. This value is null until you provision an Azure VNet in the container.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `Atlas GCP regions where the container resides. See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/) ->`,
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
					Description: `The Atlas AWS region name for where this container will exist.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The Atlas Azure region name for where this container will exist.`,
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
					Description: `The name of the Azure VNet. This value is null until you provision an Azure VNet in the container.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `Atlas GCP regions where the container resides. See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-create-container/) ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_containers",
			Category:         "Data Sources",
			ShortDescription: `Describes all Network Peering Containers in the project.`,
			Description: `

` + "`" + `mongodbatlas_network_containers` + "`" + ` describes all Network Peering Containers. The data source requires your Project ID.

-> **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to create the database user.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) Cloud provider for this Network peering container. Accepted values are AWS, GCP, and Azure. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `The Atlas AWS region name for where this container exists.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The Atlas Azure region name for where this container exists.`,
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
					Description: `The name of the Azure VNet. This value is null until you provision an Azure VNet in the container.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `Atlas GCP regions where the container resides. See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-get-containers-list/) ->`,
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
					Description: `The Atlas AWS region name for where this container exists.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The Atlas Azure region name for where this container exists.`,
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
					Description: `The name of the Azure VNet. This value is null until you provision an Azure VNet in the container.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `Atlas GCP regions where the container resides. See detailed information for arguments and attributes: [MongoDB API Network Peering Container](https://docs.atlas.mongodb.com/reference/api/vpc-get-containers-list/) ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_network_peering",
			Category:         "Data Sources",
			ShortDescription: `Describes a Network Peering.`,
			Description: `

` + "`" + `mongodbatlas_network_peering` + "`" + ` describes a Network Peering Connection.

-> **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.

-> **NOTE:** If you need to get an existing container ID see the [How-To Guide](https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/guides/howto-guide.html).


`,
			Keywords: []string{},
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
			Description: `

` + "`" + `mongodbatlas_network_peerings` + "`" + ` describes all Network Peering Connections.

-> **NOTE:** Groups and projects are synonymous terms. You may find **group_id** in the official documentation.


`,
			Keywords: []string{},
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
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_online_archive",
			Category:         "Data Sources",
			ShortDescription: `Describes an Online Archive`,
			Description: `

` + "`" + `mongodbatlas_online_archive` + "`" + ` describes an Online Archive

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "criteria.expire_after_days",
					Description: `Number of days that specifies the age limit for the data in the live Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "criteria.query",
					Description: `JSON query to use to select documents for archiving. Only for ` + "`" + `CUSTOM` + "`" + ` type`,
				},
				resource.Attribute{
					Name:        "partition_fields",
					Description: `Fields to use to partition data.`,
				},
				resource.Attribute{
					Name:        "partition_fields.field_name",
					Description: `Name of the field. To specify a nested field, use the dot notation.`,
				},
				resource.Attribute{
					Name:        "partition_fields.order",
					Description: `Position of the field in the partition. Value can be: 0,1,2 By default, the date field specified in the criteria.dateField parameter is in the first position of the partition.`,
				},
				resource.Attribute{
					Name:        "partitio_fields.field_type",
					Description: `Type of the partition field`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_online_archives",
			Category:         "Data Sources",
			ShortDescription: `Describes the list of all the online archives for a cluster`,
			Description: `

` + "`" + `mongodbatlas_online_archive` + "`" + ` Describes the list of all the online archives for a cluster

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents an online archive ## Attributes reference Each object in the results array represents an online archive with the following attributes:`,
				},
				resource.Attribute{
					Name:        "criteria.expire_after_days",
					Description: `Number of days that specifies the age limit for the data in the live Atlas cluster.`,
				},
				resource.Attribute{
					Name:        "criteria.query",
					Description: `JSON query to use to select documents for archiving. Only for ` + "`" + `CUSTOM` + "`" + ` type`,
				},
				resource.Attribute{
					Name:        "partition_fields",
					Description: `Fields to use to partition data.`,
				},
				resource.Attribute{
					Name:        "partition_fields.field_name",
					Description: `Name of the field. To specify a nested field, use the dot notation.`,
				},
				resource.Attribute{
					Name:        "partition_fields.order",
					Description: `Position of the field in the partition. Value can be: 0,1,2 By default, the date field specified in the criteria.dateField parameter is in the first position of the partition.`,
				},
				resource.Attribute{
					Name:        "partitio_fields.field_type",
					Description: `Type of the partition field`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_org_invitation",
			Category:         "Data Sources",
			ShortDescription: `Provides an Atlas Organization Invitation.`,
			Description: `

` + "`" + `mongodbatlas_org_invitation` + "`" + ` describes an invitation for a user to join an Atlas organization.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the organization to which you invited the user.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.`,
				},
				resource.Attribute{
					Name:        "invitation_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the invitation in Atlas. ## Attributes Reference In addition to the arguments, this data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated unique string that identifies this data source.`,
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
					Name:        "inviter_username",
					Description: `Atlas user who invited ` + "`" + `username` + "`" + ` to the organization.`,
				},
				resource.Attribute{
					Name:        "teams_ids",
					Description: `An array of unique 24-hexadecimal digit strings that identify the teams that the user was invited to join.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The following options are available:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated unique string that identifies this data source.`,
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
					Name:        "inviter_username",
					Description: `Atlas user who invited ` + "`" + `username` + "`" + ` to the organization.`,
				},
				resource.Attribute{
					Name:        "teams_ids",
					Description: `An array of unique 24-hexadecimal digit strings that identify the teams that the user was invited to join.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. The following options are available:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_private_endpoint_regional_mode",
			Category:         "Data Sources",
			ShortDescription: `Describes a Private Endpoint Regional Mode`,
			Description: `

` + "`" + `private_endpoint_regional_mode` + "`" + ` describe a Private Endpoint Regional Mode. This represents a Private Endpoint Regional Mode Connection that wants to retrieve settings of an Atlas project.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique identifier for the project.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Flag that indicates whether the regionalized private endpoitn setting is enabled for the project. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management. See detailed information for arguments and attributes:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management. See detailed information for arguments and attributes:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_privatelink_endpoint",
			Category:         "Data Sources",
			ShortDescription: `Describes a Private Endpoint.`,
			Description: `

` + "`" + `mongodbatlas_privatelink_endpoint` + "`" + ` describe a Private Endpoint. This represents a Private Endpoint Connection to retrieve details regarding a private endpoint by id in an Atlas project

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique identifier for the project.`,
				},
				resource.Attribute{
					Name:        "private_link_id",
					Description: `(Required) Unique identifier of the private endpoint service that you want to retrieve.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) Cloud provider for which you want to retrieve a private endpoint service. Atlas accepts ` + "`" + `AWS` + "`" + `, ` + "`" + `AZURE` + "`" + ` or ` + "`" + `GCP` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
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
				resource.Attribute{
					Name:        "private_endpoints",
					Description: `All private endpoints that you have added to this Azure Private Link Service.`,
				},
				resource.Attribute{
					Name:        "private_link_service_name",
					Description: `Name of the Azure Private Link Service that Atlas manages.`,
				},
				resource.Attribute{
					Name:        "private_link_service_resource_id",
					Description: `Resource ID of the Azure Private Link Service that Atlas manages.`,
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
					Description: `Unique alphanumeric and special character strings that identify the service attachments associated with the GCP Private Service Connect endpoint service. See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/private-endpoints-service-get-one/) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
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
				resource.Attribute{
					Name:        "private_endpoints",
					Description: `All private endpoints that you have added to this Azure Private Link Service.`,
				},
				resource.Attribute{
					Name:        "private_link_service_name",
					Description: `Name of the Azure Private Link Service that Atlas manages.`,
				},
				resource.Attribute{
					Name:        "private_link_service_resource_id",
					Description: `Resource ID of the Azure Private Link Service that Atlas manages.`,
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
					Description: `Unique alphanumeric and special character strings that identify the service attachments associated with the GCP Private Service Connect endpoint service. See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/private-endpoints-service-get-one/) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_privatelink_endpoint_service",
			Category:         "Data Sources",
			ShortDescription: `Describes a Private Endpoint Link.`,
			Description: `

` + "`" + `mongodbatlas_privatelink_endpoint_service` + "`" + ` describe a Private Endpoint Link. This represents a Private Endpoint Link Connection that wants to retrieve details in an Atlas project.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique identifier for the project.`,
				},
				resource.Attribute{
					Name:        "private_link_id",
					Description: `(Required) Unique identifier of the private endpoint service for which you want to retrieve a private endpoint.`,
				},
				resource.Attribute{
					Name:        "endpoint_service_id",
					Description: `(Required) Unique identifier of the ` + "`" + `AWS` + "`" + ` or ` + "`" + `AZURE` + "`" + ` resource.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Required) Cloud provider for which you want to create a private endpoint. Atlas accepts ` + "`" + `AWS` + "`" + ` or ` + "`" + `AZURE` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "endpoints",
					Description: `Collection of individual private endpoints that comprise your network endpoint group.`,
				},
				resource.Attribute{
					Name:        "endpoint_name",
					Description: `Forwarding rule that corresponds to the endpoint you created in GCP.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Private IP address of the network endpoint group you created in GCP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the endpoint. Atlas returns one of the [values shown above](https://docs.atlas.mongodb.com/reference/api/private-endpoints-endpoint-create-one/#std-label-ref-status-field).`,
				},
				resource.Attribute{
					Name:        "service_attachment_name",
					Description: `Unique alphanumeric and special character strings that identify the service attachment associated with the endpoint. See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/private-endpoints-endpoint-get-one/) Documentation for more information.`,
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
					Name:        "endpoints",
					Description: `Collection of individual private endpoints that comprise your network endpoint group.`,
				},
				resource.Attribute{
					Name:        "endpoint_name",
					Description: `Forwarding rule that corresponds to the endpoint you created in GCP.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Private IP address of the network endpoint group you created in GCP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the endpoint. Atlas returns one of the [values shown above](https://docs.atlas.mongodb.com/reference/api/private-endpoints-endpoint-create-one/#std-label-ref-status-field).`,
				},
				resource.Attribute{
					Name:        "service_attachment_name",
					Description: `Unique alphanumeric and special character strings that identify the service attachment associated with the endpoint. See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/private-endpoints-endpoint-get-one/) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_privatelink_endpoint_service_adl",
			Category:         "Data Sources",
			ShortDescription: `Describes an Atlas Data Lake and Online Archive PrivateLink endpoint`,
			Description: `

` + "`" + `privatelink_endpoint_service_adl` + "`" + ` Provides an Atlas Data Lake (ADL) and Online Archive PrivateLink endpoint resource.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique 24-digit hexadecimal string that identifies the project.`,
				},
				resource.Attribute{
					Name:        "endpoint_id",
					Description: `(Required) Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [|aws| PrivateLink](https://aws.amazon.com/privatelink/) feature. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Human-readable label that identifies the type of resource to associate with this private endpoint.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Human-readable label that identifies the cloud provider for this endpoint.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Human-readable string to associate with this private endpoint. For more information see: [MongoDB Atlas API - DataLake](https://docs.mongodb.com/datalake/reference/api/datalakes-api/) and [MongoDB Atlas API - Online Archive](https://docs.atlas.mongodb.com/reference/api/online-archive/) Documentation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `Human-readable label that identifies the type of resource to associate with this private endpoint.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Human-readable label that identifies the cloud provider for this endpoint.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Human-readable string to associate with this private endpoint. For more information see: [MongoDB Atlas API - DataLake](https://docs.mongodb.com/datalake/reference/api/datalakes-api/) and [MongoDB Atlas API - Online Archive](https://docs.atlas.mongodb.com/reference/api/online-archive/) Documentation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_privatelink_endpoint_service_serverless",
			Category:         "Data Sources",
			ShortDescription: `Describes a Serverless PrivateLink Endpoint Service`,
			Description: `

` + "`" + `privatelink_endpoint_service_serverless` + "`" + ` Provides a Serverless PrivateLink Endpoint Service resource.

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
					Description: `(Required) Human-readable label that identifies the serverless instance`,
				},
				resource.Attribute{
					Name:        "endpoint_id",
					Description: `(Required) Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [AWS PrivateLink](https://aws.amazon.com/privatelink/) feature.`,
				},
				resource.Attribute{
					Name:        "cloud_provider_endpoint_id",
					Description: `Unique string that identifies the private endpoint's network interface.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Human-readable string to associate with this private endpoint. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "endpoint_service_name",
					Description: `Unique string that identifies the PrivateLink endpoint service. MongoDB Cloud returns null while it creates the endpoint service.`,
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
					Name:        "status",
					Description: `Human-readable label that indicates the current operating status of the private endpoint. Values include: RESERVATION_REQUESTED, RESERVED, INITIATING, AVAILABLE, FAILED, DELETING. For more information see: [MongoDB Atlas API - Serverless Private Endpoints](https://www.mongodb.com/docs/atlas/reference/api/serverless-private-endpoints-get-one/).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_service_name",
					Description: `Unique string that identifies the PrivateLink endpoint service. MongoDB Cloud returns null while it creates the endpoint service.`,
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
					Name:        "status",
					Description: `Human-readable label that indicates the current operating status of the private endpoint. Values include: RESERVATION_REQUESTED, RESERVED, INITIATING, AVAILABLE, FAILED, DELETING. For more information see: [MongoDB Atlas API - Serverless Private Endpoints](https://www.mongodb.com/docs/atlas/reference/api/serverless-private-endpoints-get-one/).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_privatelink_endpoints_service_adl",
			Category:         "Data Sources",
			ShortDescription: `Describes the list of all Atlas Data Lake and Online Archive PrivateLink endpoints.`,
			Description: `

` + "`" + `privatelink_endpoints_service_adl` + "`" + ` Describes the list of all Atlas Data Lake (ADL) and Online Archive PrivateLink endpoints resource.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "links",
					Description: `The links array includes one or more links to sub-resources or related resources. The relations between URLs are explained in the [Web Linking Specification](http://tools.ietf.org/html/rfc5988).`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `Each element in the ` + "`" + `result` + "`" + ` array is one private endpoint.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `This value is the count of the total number of items in the result set. ` + "`" + `total_count may be greater than the number of objects in the results array if the entire result set is paginated. ### links Each object in the ` + "`" + `links` + "`" + ` array represents an online archive with the following attributes:`,
				},
				resource.Attribute{
					Name:        "self",
					Description: `The URL endpoint for this resource. ### results Each object in the ` + "`" + `results` + "`" + ` array represents an online archive with the following attributes:`,
				},
				resource.Attribute{
					Name:        "endpoint_id",
					Description: `(Required) Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [|aws| PrivateLink](https://aws.amazon.com/privatelink/) feature.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Human-readable label that identifies the type of resource to associate with this private endpoint.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `Human-readable label that identifies the cloud provider for this endpoint.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Human-readable string to associate with this private endpoint. See [MongoDB Atlas API](https://docs.atlas.mongodb.com/reference/api/online-archive-get-all-for-cluster/) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_privatelink_endpoints_service_serverless",
			Category:         "Data Sources",
			ShortDescription: `Describes the list of all Serverless PrivateLink Endpoint Service`,
			Description: `

` + "`" + `privatelink_endpoints_service_serverless` + "`" + ` Describes the list of all Serverless PrivateLink Endpoint Service resource.

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
					Description: `Human-readable label that identifies the serverless instance ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `Each element in the ` + "`" + `result` + "`" + ` array is one private serverless endpoint. ### results Each object in the ` + "`" + `results` + "`" + ` array represents an online archive with the following attributes:`,
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
					Name:        "endpoint_id",
					Description: `(Required) Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [AWS PrivateLink](https://aws.amazon.com/privatelink/) feature.`,
				},
				resource.Attribute{
					Name:        "endpoint_service_name",
					Description: `Unique string that identifies the PrivateLink endpoint service. MongoDB Cloud returns null while it creates the endpoint service.`,
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
					Name:        "status",
					Description: `Human-readable label that indicates the current operating status of the private endpoint. Values include: RESERVATION_REQUESTED, RESERVED, INITIATING, AVAILABLE, FAILED, DELETING. For more information see: [MongoDB Atlas API - Serverless Private Endpoints](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Serverless-Private-Endpoints/operation/createOnePrivateEndpointForOneServerlessInstance/).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "results",
					Description: `Each element in the ` + "`" + `result` + "`" + ` array is one private serverless endpoint. ### results Each object in the ` + "`" + `results` + "`" + ` array represents an online archive with the following attributes:`,
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
					Name:        "endpoint_id",
					Description: `(Required) Unique 22-character alphanumeric string that identifies the private endpoint. Atlas supports AWS private endpoints using the [AWS PrivateLink](https://aws.amazon.com/privatelink/) feature.`,
				},
				resource.Attribute{
					Name:        "endpoint_service_name",
					Description: `Unique string that identifies the PrivateLink endpoint service. MongoDB Cloud returns null while it creates the endpoint service.`,
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
					Name:        "status",
					Description: `Human-readable label that indicates the current operating status of the private endpoint. Values include: RESERVATION_REQUESTED, RESERVED, INITIATING, AVAILABLE, FAILED, DELETING. For more information see: [MongoDB Atlas API - Serverless Private Endpoints](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Serverless-Private-Endpoints/operation/createOnePrivateEndpointForOneServerlessInstance/).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_project",
			Category:         "Data Sources",
			ShortDescription: `Describes a Project.`,
			Description: `

` + "`" + `mongodbatlas_project` + "`" + ` describes a MongoDB Atlas Project. This represents a project that has been created.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The unique ID for the project.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The unique ID for the project. ~>`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the project you want to create. (Cannot be changed via this Provider after creation.)`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The ID of the organization you want to create the project within.`,
				},
				resource.Attribute{
					Name:        "cluster_count",
					Description: `The number of Atlas clusters deployed in the project.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The ISO-8601-formatted timestamp of when Atlas created the project.`,
				},
				resource.Attribute{
					Name:        "teams.#.team_id",
					Description: `The unique identifier of the team you want to associate with the project. The team and project must share the same parent organization.`,
				},
				resource.Attribute{
					Name:        "teams.#.role_names",
					Description: `Each string in the array represents a project role assigned to the team. Every user associated with the team inherits these roles. The following are valid roles:`,
				},
				resource.Attribute{
					Name:        "api_keys.#.api_key_id",
					Description: `The unique identifier of the programmatic API key you want to associate with the project. The programmatic API key and project must share the same parent organization.`,
				},
				resource.Attribute{
					Name:        "api_keys.#.role_names",
					Description: `Each string in the array represents a project role assigned to the programmatic API key. The following are valid roles:`,
				},
				resource.Attribute{
					Name:        "is_collect_database_specifics_statistics_enabled",
					Description: `Flag that indicates whether to enable statistics in [cluster metrics](https://www.mongodb.com/docs/atlas/monitor-cluster-metrics/) collection for the project.`,
				},
				resource.Attribute{
					Name:        "is_data_explorer_enabled",
					Description: `Flag that indicates whether to enable Data Explorer for the project. If enabled, you can query your database with an easy to use interface.`,
				},
				resource.Attribute{
					Name:        "is_performance_advisor_enabled",
					Description: `Flag that indicates whether to enable Performance Advisor and Profiler for the project. If enabled, you can analyze database logs to recommend performance improvements.`,
				},
				resource.Attribute{
					Name:        "is_realtime_performance_panel_enabled",
					Description: `Flag that indicates whether to enable Real Time Performance Panel for the project. If enabled, you can see real time metrics from your MongoDB database.`,
				},
				resource.Attribute{
					Name:        "is_schema_advisor_enabled",
					Description: `Flag that indicates whether to enable Schema Advisor for the project. If enabled, you receive customized recommendations to optimize your data model and enhance performance. Disable this setting to disable schema suggestions in the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor) and the [Data Explorer](https://www.mongodb.com/docs/atlas/atlas-ui/#std-label-atlas-ui).`,
				},
				resource.Attribute{
					Name:        "region_usage_restrictions",
					Description: `If GOV_REGIONS_ONLY the project can be used for government regions only, otherwise defaults to standard regions. For more information see [MongoDB Atlas for Government](https://www.mongodb.com/docs/atlas/government/api/#creating-a-project). See [MongoDB Atlas API - Project](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Projects) - [and MongoDB Atlas API - Teams](https://docs.atlas.mongodb.com/reference/api/project-get-teams/) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the project you want to create. (Cannot be changed via this Provider after creation.)`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The ID of the organization you want to create the project within.`,
				},
				resource.Attribute{
					Name:        "cluster_count",
					Description: `The number of Atlas clusters deployed in the project.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The ISO-8601-formatted timestamp of when Atlas created the project.`,
				},
				resource.Attribute{
					Name:        "teams.#.team_id",
					Description: `The unique identifier of the team you want to associate with the project. The team and project must share the same parent organization.`,
				},
				resource.Attribute{
					Name:        "teams.#.role_names",
					Description: `Each string in the array represents a project role assigned to the team. Every user associated with the team inherits these roles. The following are valid roles:`,
				},
				resource.Attribute{
					Name:        "api_keys.#.api_key_id",
					Description: `The unique identifier of the programmatic API key you want to associate with the project. The programmatic API key and project must share the same parent organization.`,
				},
				resource.Attribute{
					Name:        "api_keys.#.role_names",
					Description: `Each string in the array represents a project role assigned to the programmatic API key. The following are valid roles:`,
				},
				resource.Attribute{
					Name:        "is_collect_database_specifics_statistics_enabled",
					Description: `Flag that indicates whether to enable statistics in [cluster metrics](https://www.mongodb.com/docs/atlas/monitor-cluster-metrics/) collection for the project.`,
				},
				resource.Attribute{
					Name:        "is_data_explorer_enabled",
					Description: `Flag that indicates whether to enable Data Explorer for the project. If enabled, you can query your database with an easy to use interface.`,
				},
				resource.Attribute{
					Name:        "is_performance_advisor_enabled",
					Description: `Flag that indicates whether to enable Performance Advisor and Profiler for the project. If enabled, you can analyze database logs to recommend performance improvements.`,
				},
				resource.Attribute{
					Name:        "is_realtime_performance_panel_enabled",
					Description: `Flag that indicates whether to enable Real Time Performance Panel for the project. If enabled, you can see real time metrics from your MongoDB database.`,
				},
				resource.Attribute{
					Name:        "is_schema_advisor_enabled",
					Description: `Flag that indicates whether to enable Schema Advisor for the project. If enabled, you receive customized recommendations to optimize your data model and enhance performance. Disable this setting to disable schema suggestions in the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor) and the [Data Explorer](https://www.mongodb.com/docs/atlas/atlas-ui/#std-label-atlas-ui).`,
				},
				resource.Attribute{
					Name:        "region_usage_restrictions",
					Description: `If GOV_REGIONS_ONLY the project can be used for government regions only, otherwise defaults to standard regions. For more information see [MongoDB Atlas for Government](https://www.mongodb.com/docs/atlas/government/api/#creating-a-project). See [MongoDB Atlas API - Project](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Projects) - [and MongoDB Atlas API - Teams](https://docs.atlas.mongodb.com/reference/api/project-get-teams/) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_project_api_key",
			Category:         "Data Sources",
			ShortDescription: `Describes a Project API Key.`,
			Description: `

` + "`" + `mongodbatlas_project_api_key` + "`" + ` describes a MongoDB Atlas Project API Key. This represents a Project API Key that has been created.

~> **IMPORTANT WARNING:** Managing Atlas Programmatic API Keys (PAKs) with Terraform will expose sensitive organizational secrets in Terraform's state. We suggest following [Terraform's best practices](https://developer.hashicorp.com/terraform/language/state/sensitive-data). You may also want to consider managing your PAKs via a more secure method, such as the [HashiCorp Vault MongoDB Atlas Secrets Engine](https://developer.hashicorp.com/vault/docs/secrets/mongodbatlas).

-> **NOTE:** You may find project_id in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Unique identifier for the project whose API keys you want to retrieve. Use the /groups endpoint to retrieve all projects to which the authenticated user has access.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this Project API key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `Public key for this Organization API key.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `Private key for this Organization API key.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `Name of the role. This resource returns all the roles the user has in Atlas. The following are valid roles:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `Unique identifier for the project whose API keys you want to retrieve. Use the /groups endpoint to retrieve all projects to which the authenticated user has access.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this Project API key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `Public key for this Organization API key.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `Private key for this Organization API key.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `Name of the role. This resource returns all the roles the user has in Atlas. The following are valid roles:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_project_api_keys",
			Category:         "Data Sources",
			ShortDescription: `Describes a API Keys.`,
			Description: `

` + "`" + `mongodbatlas_api_keys` + "`" + ` describe all API Keys. This represents API Keys that have been created.

~> **IMPORTANT WARNING:** Managing Atlas Programmatic API Keys (PAKs) with Terraform will expose sensitive organizational secrets in Terraform's state. We suggest following [Terraform's best practices](https://developer.hashicorp.com/terraform/language/state/sensitive-data). You may also want to consider managing your PAKs via a more secure method, such as the [HashiCorp Vault MongoDB Atlas Secrets Engine](https://developer.hashicorp.com/vault/docs/secrets/mongodbatlas).

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Projects. ### API Keys`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Unique identifier for the project whose API keys you want to retrieve. Use the /groups endpoint to retrieve all projects to which the authenticated user has access.`,
				},
				resource.Attribute{
					Name:        "api_key_id",
					Description: `Unique identifier for the API key you want to update. Use the /orgs/{ORG-ID}/apiKeys endpoint to retrieve all API keys to which the authenticated user has access for the specified organization.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of this Project API key.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `Name of the role. This resource returns all the roles the user has in Atlas. The following are valid roles:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_project_invitation",
			Category:         "Data Sources",
			ShortDescription: `Provides an Atlas project invitation.`,
			Description: `

` + "`" + `mongodbatlas_project_invitation` + "`" + ` describes an invitation to a user to join an Atlas project.

-> **NOTE:** Groups and projects are synonymous terms. You may find GROUP-ID in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the project to which you invited the user.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Email address of the invited user. This is the address to which Atlas sends the invite. If the user accepts the invitation, they log in to Atlas with this username.`,
				},
				resource.Attribute{
					Name:        "invitation_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the invitation in Atlas. ## Attributes Reference In addition to the arguments, this data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated unique string that identifies this data source.`,
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
					Name:        "inviter_username",
					Description: `Atlas user who invited ` + "`" + `username` + "`" + ` to the project.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. Refer to the [MongoDB Documentation](https://www.mongodb.com/docs/atlas/reference/user-roles/#project-roles) for information on valid roles. See the [MongoDB Atlas Administration API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#operation/inviteOneMongoDBCloudUserToJoinOneProject) documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated unique string that identifies this data source.`,
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
					Name:        "inviter_username",
					Description: `Atlas user who invited ` + "`" + `username` + "`" + ` to the project.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Atlas roles to assign to the invited user. If the user accepts the invitation, Atlas assigns these roles to them. Refer to the [MongoDB Documentation](https://www.mongodb.com/docs/atlas/reference/user-roles/#project-roles) for information on valid roles. See the [MongoDB Atlas Administration API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#operation/inviteOneMongoDBCloudUserToJoinOneProject) documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_project_ip_access_list",
			Category:         "Data Sources",
			ShortDescription: `Provides an IP Access List resource.`,
			Description: `

` + "`" + `mongodbatlas_project_ip_access_list` + "`" + ` describes an IP Access List entry resource. The access list grants access from IPs, CIDRs or AWS Security Groups (if VPC Peering is enabled) to clusters within the Project.

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
					Description: `(Optional) Unique identifier of the AWS security group to add to the access list.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) Range of IP addresses in CIDR notation to be added to the access list.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Single IP address to be added to the access list. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used by Terraform for internal management and can be used to import.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment to add to the access list entry. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/access-lists/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used by Terraform for internal management and can be used to import.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment to add to the access list entry. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/access-lists/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_projects",
			Category:         "Data Sources",
			ShortDescription: `Describes a Projects.`,
			Description: `

` + "`" + `mongodbatlas_projects` + "`" + ` describe all Projects. This represents projects that have been created.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "page_num",
					Description: `(Optional) The page to return. Defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `(Optional) Number of items to return per page, up to a maximum of 500. Defaults to ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Represents the total of the projects`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Projects. ### Project`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the project you want to create. (Cannot be changed via this Provider after creation.)`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The ID of the organization you want to create the project within.`,
				},
				resource.Attribute{
					Name:        "cluster_count",
					Description: `The number of Atlas clusters deployed in the project.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The ISO-8601-formatted timestamp of when Atlas created the project.`,
				},
				resource.Attribute{
					Name:        "teams.#.team_id",
					Description: `The unique identifier of the team you want to associate with the project. The team and project must share the same parent organization.`,
				},
				resource.Attribute{
					Name:        "teams.#.role_names",
					Description: `Each string in the array represents a project role assigned to the team. Every user associated with the team inherits these roles. The following are valid roles:`,
				},
				resource.Attribute{
					Name:        "api_keys.#.api_key_id",
					Description: `The unique identifier of the Organization Programmatic API key assigned to the Project.`,
				},
				resource.Attribute{
					Name:        "api_keys.#.role_names",
					Description: `List of roles that the Organization Programmatic API key has been assigned. The following are valid roles:`,
				},
				resource.Attribute{
					Name:        "is_collect_database_specifics_statistics_enabled",
					Description: `Flag that indicates whether to enable statistics in [cluster metrics](https://www.mongodb.com/docs/atlas/monitor-cluster-metrics/) collection for the project.`,
				},
				resource.Attribute{
					Name:        "is_data_explorer_enabled",
					Description: `Flag that indicates whether to enable Data Explorer for the project. If enabled, you can query your database with an easy to use interface.`,
				},
				resource.Attribute{
					Name:        "is_performance_advisor_enabled",
					Description: `Flag that indicates whether to enable Performance Advisor and Profiler for the project. If enabled, you can analyze database logs to recommend performance improvements.`,
				},
				resource.Attribute{
					Name:        "is_realtime_performance_panel_enabled",
					Description: `Flag that indicates whether to enable Real Time Performance Panel for the project. If enabled, you can see real time metrics from your MongoDB database.`,
				},
				resource.Attribute{
					Name:        "is_schema_advisor_enabled",
					Description: `Flag that indicates whether to enable Schema Advisor for the project. If enabled, you receive customized recommendations to optimize your data model and enhance performance. Disable this setting to disable schema suggestions in the [Performance Advisor](https://www.mongodb.com/docs/atlas/performance-advisor/#std-label-performance-advisor) and the [Data Explorer](https://www.mongodb.com/docs/atlas/atlas-ui/#std-label-atlas-ui).`,
				},
				resource.Attribute{
					Name:        "region_usage_restrictions",
					Description: `If GOV_REGIONS_ONLY the project can be used for government regions only, otherwise defaults to standard regions. For more information see [MongoDB Atlas for Government](https://www.mongodb.com/docs/atlas/government/api/#creating-a-project). See [MongoDB Atlas API - Projects](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Projects) - [and MongoDB Atlas API - Teams](https://docs.atlas.mongodb.com/reference/api/project-get-teams/) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_roles_org_id",
			Category:         "Data Sources",
			ShortDescription: `Describes a Roles Org ID.`,
			Description: `

` + "`" + `mongodbatlas_project` + "`" + ` describes a MongoDB Atlas Roles Org ID. This represents a Roles Org ID.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `The ID of the organization you want to retrieve associated to an API Key. See [MongoDB Atlas API - Role Org ID](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Root/operation/getSystemStatus) - Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `The ID of the organization you want to retrieve associated to an API Key. See [MongoDB Atlas API - Role Org ID](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Root/operation/getSystemStatus) - Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_search_index",
			Category:         "Data Sources",
			ShortDescription: `Describes a Search Index.`,
			Description: `

` + "`" + `mongodbatlas_search_index` + "`" + ` describe a single search indexes. This represents a single search index that have been created.

> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "index_id",
					Description: `(Required) The unique identifier of the Atlas Search index. Use the ` + "`" + `mongodbatlas_search_indexes` + "`" + `datasource to find the IDs of all Atlas Search indexes.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique identifier for the [project](https://docs.atlas.mongodb.com/organizations-projects/#std-label-projects) that contains the specified cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the cluster containing the collection with one or more Atlas Search indexes. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the index.`,
				},
				resource.Attribute{
					Name:        "analyzer",
					Description: `[Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index.`,
				},
				resource.Attribute{
					Name:        "analyzers",
					Description: `[Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index (this is an array of objects).`,
				},
				resource.Attribute{
					Name:        "collection_name",
					Description: `Name of the collection the index is on.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `Name of the database the collection is in.`,
				},
				resource.Attribute{
					Name:        "mappings_dynamic",
					Description: `Flag indicating whether the index uses dynamic or static mappings.`,
				},
				resource.Attribute{
					Name:        "mappings_fields",
					Description: `Object containing one or more field specifications.`,
				},
				resource.Attribute{
					Name:        "search_analyzer",
					Description: `[Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index.`,
				},
				resource.Attribute{
					Name:        "synonyms",
					Description: `Synonyms mapping definition to use in this index.`,
				},
				resource.Attribute{
					Name:        "synonyms.#.name",
					Description: `Name of the [synonym mapping definition](https://docs.atlas.mongodb.com/reference/atlas-search/synonyms/#std-label-synonyms-ref).`,
				},
				resource.Attribute{
					Name:        "synonyms.#.source_collection",
					Description: `Name of the source MongoDB collection for the synonyms.`,
				},
				resource.Attribute{
					Name:        "synonyms.#.analyzer",
					Description: `Name of the [analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use with this synonym mapping. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/atlas-search/) - [and MongoDB Atlas API - Search](https://docs.atlas.mongodb.com/reference/api/atlas-search/) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the index.`,
				},
				resource.Attribute{
					Name:        "analyzer",
					Description: `[Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index.`,
				},
				resource.Attribute{
					Name:        "analyzers",
					Description: `[Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index (this is an array of objects).`,
				},
				resource.Attribute{
					Name:        "collection_name",
					Description: `Name of the collection the index is on.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `Name of the database the collection is in.`,
				},
				resource.Attribute{
					Name:        "mappings_dynamic",
					Description: `Flag indicating whether the index uses dynamic or static mappings.`,
				},
				resource.Attribute{
					Name:        "mappings_fields",
					Description: `Object containing one or more field specifications.`,
				},
				resource.Attribute{
					Name:        "search_analyzer",
					Description: `[Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index.`,
				},
				resource.Attribute{
					Name:        "synonyms",
					Description: `Synonyms mapping definition to use in this index.`,
				},
				resource.Attribute{
					Name:        "synonyms.#.name",
					Description: `Name of the [synonym mapping definition](https://docs.atlas.mongodb.com/reference/atlas-search/synonyms/#std-label-synonyms-ref).`,
				},
				resource.Attribute{
					Name:        "synonyms.#.source_collection",
					Description: `Name of the source MongoDB collection for the synonyms.`,
				},
				resource.Attribute{
					Name:        "synonyms.#.analyzer",
					Description: `Name of the [analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use with this synonym mapping. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/atlas-search/) - [and MongoDB Atlas API - Search](https://docs.atlas.mongodb.com/reference/api/atlas-search/) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_search_indexes",
			Category:         "Data Sources",
			ShortDescription: `Describes a Search Indexes.`,
			Description: `

` + "`" + `mongodbatlas_search_indexes` + "`" + ` describe all search indexes. This represents search indexes that have been created.

> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique identifier for the [project](https://docs.atlas.mongodb.com/organizations-projects/#std-label-projects) that contains the specified cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) Name of the cluster containing the collection with one or more Atlas Search indexes.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) Name of the database containing the collection with one or more Atlas Search indexes.`,
				},
				resource.Attribute{
					Name:        "collection_name",
					Description: `(Required) Name of the collection with one or more Atlas Search indexes.`,
				},
				resource.Attribute{
					Name:        "page_num",
					Description: `Page number, starting with one, that Atlas returns of the total number of objects.`,
				},
				resource.Attribute{
					Name:        "items_per_page",
					Description: `Number of items that Atlas returns per page, up to a maximum of 500. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "total_count",
					Description: `Represents the total of the search indexes`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a search index. ### Results`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the index.`,
				},
				resource.Attribute{
					Name:        "analyzer",
					Description: `[Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index.`,
				},
				resource.Attribute{
					Name:        "analyzers",
					Description: `[Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index (this is an array of objects).`,
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
					Description: `Flag indicating whether the index uses dynamic or static mappings.`,
				},
				resource.Attribute{
					Name:        "mappings_fields",
					Description: `Object containing one or more field specifications.`,
				},
				resource.Attribute{
					Name:        "search_analyzer",
					Description: `[Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index.`,
				},
				resource.Attribute{
					Name:        "synonyms",
					Description: `Synonyms mapping definition to use in this index.`,
				},
				resource.Attribute{
					Name:        "synonyms.#.name",
					Description: `Name of the [synonym mapping definition](https://docs.atlas.mongodb.com/reference/atlas-search/synonyms/#std-label-synonyms-ref).`,
				},
				resource.Attribute{
					Name:        "synonyms.#.source_collection",
					Description: `Name of the source MongoDB collection for the synonyms.`,
				},
				resource.Attribute{
					Name:        "synonyms.#.analyzer",
					Description: `Name of the [analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use with this synonym mapping. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/atlas-search/) - [and MongoDB Atlas API - Search](https://docs.atlas.mongodb.com/reference/api/atlas-search/) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "total_count",
					Description: `Represents the total of the search indexes`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a search index. ### Results`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the index.`,
				},
				resource.Attribute{
					Name:        "analyzer",
					Description: `[Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when creating the index.`,
				},
				resource.Attribute{
					Name:        "analyzers",
					Description: `[Custom analyzers](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/custom/#std-label-custom-analyzers) to use in this index (this is an array of objects).`,
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
					Description: `Flag indicating whether the index uses dynamic or static mappings.`,
				},
				resource.Attribute{
					Name:        "mappings_fields",
					Description: `Object containing one or more field specifications.`,
				},
				resource.Attribute{
					Name:        "search_analyzer",
					Description: `[Analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use when searching the index.`,
				},
				resource.Attribute{
					Name:        "synonyms",
					Description: `Synonyms mapping definition to use in this index.`,
				},
				resource.Attribute{
					Name:        "synonyms.#.name",
					Description: `Name of the [synonym mapping definition](https://docs.atlas.mongodb.com/reference/atlas-search/synonyms/#std-label-synonyms-ref).`,
				},
				resource.Attribute{
					Name:        "synonyms.#.source_collection",
					Description: `Name of the source MongoDB collection for the synonyms.`,
				},
				resource.Attribute{
					Name:        "synonyms.#.analyzer",
					Description: `Name of the [analyzer](https://docs.atlas.mongodb.com/reference/atlas-search/analyzers/#std-label-analyzers-ref) to use with this synonym mapping. For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/atlas-search/) - [and MongoDB Atlas API - Search](https://docs.atlas.mongodb.com/reference/api/atlas-search/) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_serverless_instance",
			Category:         "Data Sources",
			ShortDescription: `Provides a Serverless Instance.`,
			Description: `

` + "`" + `mongodbatlas_serverless_instance` + "`" + ` describe a single serverless instance. This represents a single serverless instance that have been created.
> **NOTE:**  Serverless instances do not support some Atlas features at this time.
For a full list of unsupported features, see [Serverless Instance Limitations](https://docs.atlas.mongodb.com/reference/serverless-instance-limitations/).
 

> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique 24-hexadecimal digit string that identifies the project that contains your serverless instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Human-readable label that identifies your serverless instance. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "connection_strings_standard_srv",
					Description: `Public ` + "`" + `mongodb+srv://` + "`" + ` connection string that you can use to connect to this serverless instance.`,
				},
				resource.Attribute{
					Name:        "connection_strings_private_endpoint_srv",
					Description: `List of Serverless Private Endpoint Connections`,
				},
				resource.Attribute{
					Name:        "created_date",
					Description: `Timestamp that indicates when MongoDB Cloud created the serverless instance. The timestamp displays in the ISO 8601 date and time format in UTC.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies the serverless instance.`,
				},
				resource.Attribute{
					Name:        "mongo_db_version",
					Description: `Version of MongoDB that the serverless instance runs, in ` + "`" + `<major version>` + "`" + `.` + "`" + `<minor version>` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "provider_settings_backing_provider_name",
					Description: `Cloud service provider on which MongoDB Cloud provisioned the serverless instance.`,
				},
				resource.Attribute{
					Name:        "provider_settings_provider_name",
					Description: `Cloud service provider that applies to the provisioned the serverless instance.`,
				},
				resource.Attribute{
					Name:        "provider_settings_region_name",
					Description: `Human-readable label that identifies the physical location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Stage of deployment of this serverless instance when the resource made its request.`,
				},
				resource.Attribute{
					Name:        "continuous_backup_enabled",
					Description: `Flag that indicates whether the serverless instance uses Serverless Continuous Backup.`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster. For more information see: [MongoDB Atlas API - Serverless Instance](https://docs.atlas.mongodb.com/reference/api/serverless-instances/) Documentation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_strings_standard_srv",
					Description: `Public ` + "`" + `mongodb+srv://` + "`" + ` connection string that you can use to connect to this serverless instance.`,
				},
				resource.Attribute{
					Name:        "connection_strings_private_endpoint_srv",
					Description: `List of Serverless Private Endpoint Connections`,
				},
				resource.Attribute{
					Name:        "created_date",
					Description: `Timestamp that indicates when MongoDB Cloud created the serverless instance. The timestamp displays in the ISO 8601 date and time format in UTC.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies the serverless instance.`,
				},
				resource.Attribute{
					Name:        "mongo_db_version",
					Description: `Version of MongoDB that the serverless instance runs, in ` + "`" + `<major version>` + "`" + `.` + "`" + `<minor version>` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "provider_settings_backing_provider_name",
					Description: `Cloud service provider on which MongoDB Cloud provisioned the serverless instance.`,
				},
				resource.Attribute{
					Name:        "provider_settings_provider_name",
					Description: `Cloud service provider that applies to the provisioned the serverless instance.`,
				},
				resource.Attribute{
					Name:        "provider_settings_region_name",
					Description: `Human-readable label that identifies the physical location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Stage of deployment of this serverless instance when the resource made its request.`,
				},
				resource.Attribute{
					Name:        "continuous_backup_enabled",
					Description: `Flag that indicates whether the serverless instance uses Serverless Continuous Backup.`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster. For more information see: [MongoDB Atlas API - Serverless Instance](https://docs.atlas.mongodb.com/reference/api/serverless-instances/) Documentation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_serverless_instances",
			Category:         "Data Sources",
			ShortDescription: `Describes a Serverless Instances.`,
			Description: `

` + "`" + `mongodbatlas_serverless_instances` + "`" + ` describe all serverless instances. This represents serverless instances that have been created for the specified group id.

> **NOTE:**  Serverless instances do not support some Atlas features at this time.
For a full list of unsupported features, see [Serverless Instance Limitations](https://docs.atlas.mongodb.com/reference/serverless-instance-limitations/).

> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Unique identifier for the [project](https://docs.atlas.mongodb.com/organizations-projects/#std-label-projects) that contains the specified cluster. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a search index. ### Results`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Human-readable label that identifies your serverless instance.`,
				},
				resource.Attribute{
					Name:        "connection_strings_standard_srv",
					Description: `Public ` + "`" + `mongodb+srv://` + "`" + ` connection string that you can use to connect to this serverless instance.`,
				},
				resource.Attribute{
					Name:        "created_date",
					Description: `Timestamp that indicates when MongoDB Cloud created the serverless instance. The timestamp displays in the ISO 8601 date and time format in UTC.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies the serverless instance.`,
				},
				resource.Attribute{
					Name:        "mongo_db_version",
					Description: `Version of MongoDB that the serverless instance runs, in ` + "`" + `<major version>` + "`" + `.` + "`" + `<minor version>` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "provider_settings_backing_provider_name",
					Description: `Cloud service provider on which MongoDB Cloud provisioned the serverless instance.`,
				},
				resource.Attribute{
					Name:        "provider_settings_provider_name",
					Description: `Cloud service provider that applies to the provisioned the serverless instance.`,
				},
				resource.Attribute{
					Name:        "provider_settings_region_name",
					Description: `Human-readable label that identifies the physical location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Stage of deployment of this serverless instance when the resource made its request.`,
				},
				resource.Attribute{
					Name:        "continuous_backup_enabled",
					Description: `Flag that indicates whether the serverless instance uses Serverless Continuous Backup.`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster. For more information see: [MongoDB Atlas API - Serverless Instance](https://docs.atlas.mongodb.com/reference/api/serverless-instances/) Documentation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a search index. ### Results`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Human-readable label that identifies your serverless instance.`,
				},
				resource.Attribute{
					Name:        "connection_strings_standard_srv",
					Description: `Public ` + "`" + `mongodb+srv://` + "`" + ` connection string that you can use to connect to this serverless instance.`,
				},
				resource.Attribute{
					Name:        "created_date",
					Description: `Timestamp that indicates when MongoDB Cloud created the serverless instance. The timestamp displays in the ISO 8601 date and time format in UTC.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique 24-hexadecimal digit string that identifies the serverless instance.`,
				},
				resource.Attribute{
					Name:        "mongo_db_version",
					Description: `Version of MongoDB that the serverless instance runs, in ` + "`" + `<major version>` + "`" + `.` + "`" + `<minor version>` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "provider_settings_backing_provider_name",
					Description: `Cloud service provider on which MongoDB Cloud provisioned the serverless instance.`,
				},
				resource.Attribute{
					Name:        "provider_settings_provider_name",
					Description: `Cloud service provider that applies to the provisioned the serverless instance.`,
				},
				resource.Attribute{
					Name:        "provider_settings_region_name",
					Description: `Human-readable label that identifies the physical location of your MongoDB serverless instance. The region you choose can affect network latency for clients accessing your databases.`,
				},
				resource.Attribute{
					Name:        "state_name",
					Description: `Stage of deployment of this serverless instance when the resource made its request.`,
				},
				resource.Attribute{
					Name:        "continuous_backup_enabled",
					Description: `Flag that indicates whether the serverless instance uses Serverless Continuous Backup.`,
				},
				resource.Attribute{
					Name:        "termination_protection_enabled",
					Description: `Flag that indicates whether termination protection is enabled on the cluster. If set to true, MongoDB Cloud won't delete the cluster. If set to false, MongoDB Cloud will delete the cluster. For more information see: [MongoDB Atlas API - Serverless Instance](https://docs.atlas.mongodb.com/reference/api/serverless-instances/) Documentation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_teams",
			Category:         "Data Sources",
			ShortDescription: `Describes a Team.`,
			Description: `

` + "`" + `mongodbatlas_teams` + "`" + ` describes a Team. The resource requires your Organization ID, Project ID and Team ID.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) The unique identifier for the organization you want to associate the team with.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `(Optional) The unique identifier for the team.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The team name. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The unique identifier for the team.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the team you want to create.`,
				},
				resource.Attribute{
					Name:        "usernames",
					Description: `The users who are part of the organization. See detailed information for arguments and attributes: [MongoDB API Teams](https://docs.atlas.mongodb.com/reference/api/teams-create-one/)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Terraform's unique identifier used internally for state management.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The unique identifier for the team.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the team you want to create.`,
				},
				resource.Attribute{
					Name:        "usernames",
					Description: `The users who are part of the organization. See detailed information for arguments and attributes: [MongoDB API Teams](https://docs.atlas.mongodb.com/reference/api/teams-create-one/)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_third_party_integration",
			Category:         "Data Sources",
			ShortDescription: `Describes all Third-Party Integration Settings in the project.`,
			Description: `

` + "`" + `mongodbatlas_third_party_integration` + "`" + ` describe a Third-Party Integration Settings for the given type.

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to get all Third-Party service integrations`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages and can be used to import.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Property equal to its own integration type Additional values based on Type`,
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
					Name:        "password",
					Description: `Your Prometheus password.`,
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
					Description: `Whether your cluster has Prometheus enabled. See [MongoDB Atlas API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Third-Party-Integrations/operation/createThirdPartyIntegration) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier used for terraform for internal manages and can be used to import.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Property equal to its own integration type Additional values based on Type`,
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
					Name:        "password",
					Description: `Your Prometheus password.`,
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
					Description: `Whether your cluster has Prometheus enabled. See [MongoDB Atlas API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Third-Party-Integrations/operation/createThirdPartyIntegration) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_third_party_integrations",
			Category:         "Data Sources",
			ShortDescription: `Describes all Third-Party Integration Settings in the project.`,
			Description: `

` + "`" + `mongodbatlas_third_party_integrations` + "`" + ` describe all Third-Party Integration Settings. This represents two Third-Party services ` + "`" + `PAGER_DUTY` + "`" + ` and ` + "`" + `FLOWDOCK` + "`" + `
applied across the project. 

-> **NOTE:** Groups and projects are synonymous terms. You may find ` + "`" + `groupId` + "`" + ` in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The unique ID for the project to get all Third-Party service integrations ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Third-Party service integration. ### Third-Party Service Integration`,
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
					Name:        "name",
					Description: `Your Microsoft Teams incoming webhook name.`,
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
					Description: `Whether your cluster has Prometheus enabled. See [MongoDB Atlas API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Third-Party-Integrations/operation/createThirdPartyIntegration) Documentation for more information.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Autogenerated Unique ID for this data source.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list where each represents a Third-Party service integration. ### Third-Party Service Integration`,
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
					Name:        "name",
					Description: `Your Microsoft Teams incoming webhook name.`,
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
					Description: `Whether your cluster has Prometheus enabled. See [MongoDB Atlas API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Third-Party-Integrations/operation/createThirdPartyIntegration) Documentation for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mongodbatlas_x509_authentication_database_user",
			Category:         "Data Sources",
			ShortDescription: `Describes a Custom DB Role.`,
			Description: `

` + "`" + `mongodbatlas_x509_authentication_database_user` + "`" + ` describe a X509 Authentication Database User. This represents a X509 Authentication Database User.

-> **NOTE:** Groups and projects are synonymous terms. You may find group_id in the official documentation.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Identifier for the Atlas project associated with the X.509 configuration.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Username of the database user to create a certificate for. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `Fully distinguished name of the database user to which this certificate belongs. To learn more, see RFC 2253. See [MongoDB Atlas - X509 User Certificates](https://docs.atlas.mongodb.com/reference/api/x509-configuration-get-certificates/) and [MongoDB Atlas - Current X509 Configuratuion](https://docs.atlas.mongodb.com/reference/api/x509-configuration-get-current/) Documentation for more information.`,
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
					Description: `Fully distinguished name of the database user to which this certificate belongs. To learn more, see RFC 2253. See [MongoDB Atlas - X509 User Certificates](https://docs.atlas.mongodb.com/reference/api/x509-configuration-get-certificates/) and [MongoDB Atlas - Current X509 Configuratuion](https://docs.atlas.mongodb.com/reference/api/x509-configuration-get-current/) Documentation for more information.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"mongodbatlas_access_list_api_key":                      0,
		"mongodbatlas_access_list_api_keys":                     1,
		"mongodbatlas_advanced_cluster":                         2,
		"mongodbatlas_advanced_clusters":                        3,
		"mongodbatlas_alert_configuration":                      4,
		"mongodbatlas_alert_configurations":                     5,
		"mongodbatlas_api_key":                                  6,
		"mongodbatlas_api_keys":                                 7,
		"mongodbatlas_auditing":                                 8,
		"mongodbatlas_backup_compliance_policy":                 9,
		"mongodbatlas_cloud_backup_schedule":                    10,
		"mongodbatlas_cloud_backup_snapshot":                    11,
		"mongodbatlas_cloud_backup_snapshot_export_bucket":      12,
		"mongodbatlas_cloud_backup_snapshot_export_buckets":     13,
		"mongodbatlas_cloud_backup_snapshot_export_job":         14,
		"mongodbatlas_cloud_backup_snapshot_export_jobs":        15,
		"mongodbatlas_cloud_backup_snapshot_restore_job":        16,
		"mongodbatlas_cloud_backup_snapshot_restore_jobs":       17,
		"mongodbatlas_cloud_backup_snapshots":                   18,
		"mongodbatlas_cloud_provider_access":                    19,
		"mongodbatlas_cloud_provider_access_setup":              20,
		"mongodbatlas_cloud_provider_snapshot":                  21,
		"mongodbatlas_cloud_provider_snapshot_backup_policy":    22,
		"mongodbatlas_cloud_provider_snapshot_restore_job":      23,
		"mongodbatlas_cloud_provider_snapshot_restore_jobs":     24,
		"mongodbatlas_cloud_provider_snapshots":                 25,
		"mongodbatlas_cluster":                                  26,
		"mongodbatlas_clusters":                                 27,
		"mongodbatlas_custom_db_role":                           28,
		"mongodbatlas_custom_db_roles":                          29,
		"mongodbatlas_custom_dns_configuration_cluster_aws":     30,
		"mongodbatlas_data_lake":                                31,
		"mongodbatlas_data_lakes":                               32,
		"mongodbatlas_database_user":                            33,
		"mongodbatlas_database_users":                           34,
		"mongodbatlas_event_trigger":                            35,
		"mongodbatlas_event_triggers":                           36,
		"mongodbatlas_federated_settings":                       37,
		"mongodbatlas_federated_settings_identity_provider":     38,
		"mongodbatlas_federated_settings_identity_providers":    39,
		"mongodbatlas_federated_settings_org_config":            40,
		"mongodbatlas_federated_settings_org_configs":           41,
		"mongodbatlas_federated_settings_org_role_mapping":      42,
		"mongodbatlas_federated_settings_org_role_mappings":     43,
		"mongodbatlas_global_cluster_config":                    44,
		"mongodbatlas_ldap_configuration":                       45,
		"mongodbatlas_ldap_verify":                              46,
		"mongodbatlas_maintenance_window":                       47,
		"mongodbatlas_network_container":                        48,
		"mongodbatlas_network_containers":                       49,
		"mongodbatlas_network_peering":                          50,
		"mongodbatlas_network_peerings":                         51,
		"mongodbatlas_online_archive":                           52,
		"mongodbatlas_online_archives":                          53,
		"mongodbatlas_org_invitation":                           54,
		"mongodbatlas_private_endpoint_regional_mode":           55,
		"mongodbatlas_privatelink_endpoint":                     56,
		"mongodbatlas_privatelink_endpoint_service":             57,
		"mongodbatlas_privatelink_endpoint_service_adl":         58,
		"mongodbatlas_privatelink_endpoint_service_serverless":  59,
		"mongodbatlas_privatelink_endpoints_service_adl":        60,
		"mongodbatlas_privatelink_endpoints_service_serverless": 61,
		"mongodbatlas_project":                                  62,
		"mongodbatlas_project_api_key":                          63,
		"mongodbatlas_project_api_keys":                         64,
		"mongodbatlas_project_invitation":                       65,
		"mongodbatlas_project_ip_access_list":                   66,
		"mongodbatlas_projects":                                 67,
		"mongodbatlas_roles_org_id":                             68,
		"mongodbatlas_search_index":                             69,
		"mongodbatlas_search_indexes":                           70,
		"mongodbatlas_serverless_instance":                      71,
		"mongodbatlas_serverless_instances":                     72,
		"mongodbatlas_teams":                                    73,
		"mongodbatlas_third_party_integration":                  74,
		"mongodbatlas_third_party_integrations":                 75,
		"mongodbatlas_x509_authentication_database_user":        76,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
