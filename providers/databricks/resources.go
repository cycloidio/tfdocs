package databricks

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "databricks_catalog",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of Catalog relative to parent metastore. Change forces creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "storage_root",
					Description: `(Optional) Managed location of the catalog. Location in cloud storage where data for managed tables will be stored. If not specified, the location will default to the metastore root location. Change forces creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Optional) For Delta Sharing Catalogs: the name of the delta sharing provider. Change forces creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "share_name",
					Description: `(Optional) For Delta Sharing Catalogs: the name of the share under the share provider. Change forces creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) Username/groupname/sp application_id of the catalog owner.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) User-supplied free-form text.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) Extensible Catalog properties.`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional) Delete catalog regardless of its contents. ## Import This resource can be imported by name: ` + "`" + `` + "`" + `` + "`" + `bash terraform import databricks_catalog.this <name> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_cluster",
			Category:         "Compute",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"compute",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "num_workers",
					Description: `(Optional) Number of worker nodes that this cluster should have. A cluster has one Spark driver and ` + "`" + `num_workers` + "`" + ` executors for a total of ` + "`" + `num_workers` + "`" + ` + 1 Spark nodes.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Optional) Cluster name, which doesn’t have to be unique. If not specified at creation, the cluster name will be an empty string.`,
				},
				resource.Attribute{
					Name:        "spark_version",
					Description: `(Required) [Runtime version](https://docs.databricks.com/runtime/index.html) of the cluster. Any supported [databricks_spark_version](../data-sources/spark_version.md) id. We advise using [Cluster Policies](cluster_policy.md) to restrict the list of versions for simplicity while maintaining enough control.`,
				},
				resource.Attribute{
					Name:        "runtime_engine",
					Description: `(Optional) The type of runtime engine to use. If not specified, the runtime engine type is inferred based on the spark_version value. Allowed values include: ` + "`" + `PHOTON` + "`" + `, ` + "`" + `STANDARD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "driver_node_type_id",
					Description: `(Optional) The node type of the Spark driver. This field is optional; if unset, API will set the driver node type to the same value as ` + "`" + `node_type_id` + "`" + ` defined above.`,
				},
				resource.Attribute{
					Name:        "node_type_id",
					Description: `(Required - optional if ` + "`" + `instance_pool_id` + "`" + ` is given) Any supported [databricks_node_type](../data-sources/node_type.md) id. If ` + "`" + `instance_pool_id` + "`" + ` is specified, this field is not needed.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) Identifier of [Cluster Policy](cluster_policy.md) to validate cluster and preset certain defaults.`,
				},
				resource.Attribute{
					Name:        "apply_policy_default_values",
					Description: `(Optional) Whether to use policy default values for missing cluster attributes.`,
				},
				resource.Attribute{
					Name:        "autotermination_minutes",
					Description: `(Optional) Automatically terminate the cluster after being inactive for this time in minutes. If specified, the threshold must be between 10 and 10000 minutes. You can also set this value to 0 to explicitly disable automatic termination. Defaults to ` + "`" + `60` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_elastic_disk",
					Description: `(Optional) If you don’t want to allocate a fixed number of EBS volumes at cluster creation time, use autoscaling local storage. With autoscaling local storage, Databricks monitors the amount of free disk space available on your cluster’s Spark workers. If a worker begins to run too low on disk, Databricks automatically attaches a new EBS volume to the worker before it runs out of disk space. EBS volumes are attached up to a limit of 5 TB of total disk space per instance (including the instance’s local storage). To scale down EBS usage, make sure you have ` + "`" + `autotermination_minutes` + "`" + ` and ` + "`" + `autoscale` + "`" + ` attributes set. More documentation available at [cluster configuration page](https://docs.databricks.com/clusters/configure.html#autoscaling-local-storage-1).`,
				},
				resource.Attribute{
					Name:        "enable_local_disk_encryption",
					Description: `(Optional) Some instance types you use to run clusters may have locally attached disks. Databricks may store shuffle data or temporary data on these locally attached disks. To ensure that all data at rest is encrypted for all storage types, including shuffle data stored temporarily on your cluster’s local disks, you can enable local disk encryption. When local disk encryption is enabled, Databricks generates an encryption key locally unique to each cluster node and uses it to encrypt all data stored on local disks. The scope of the key is local to each cluster node and is destroyed along with the cluster node itself. During its lifetime, the key resides in memory for encryption and decryption and is stored encrypted on the disk.`,
				},
				resource.Attribute{
					Name:        "data_security_mode",
					Description: `(Optional) Select the security features of the cluster. [Unity Catalog requires](https://docs.databricks.com/data-governance/unity-catalog/compute.html#create-clusters--sql-warehouses-with-unity-catalog-access) ` + "`" + `SINGLE_USER` + "`" + ` or ` + "`" + `USER_ISOLATION` + "`" + ` mode. ` + "`" + `LEGACY_PASSTHROUGH` + "`" + ` for passthrough cluster and ` + "`" + `LEGACY_TABLE_ACL` + "`" + ` for Table ACL cluster. Default to ` + "`" + `NONE` + "`" + `, i.e. no security feature enabled. In the Databricks UI, this has been recently been renamed`,
				},
				resource.Attribute{
					Name:        "single_user_name",
					Description: `(Optional) The optional user name of the user to assign to an interactive cluster. This field is required when using standard AAD Passthrough for Azure Data Lake Storage (ADLS) with a single-user cluster (i.e., not high-concurrency clusters).`,
				},
				resource.Attribute{
					Name:        "idempotency_token",
					Description: `(Optional) An optional token to guarantee the idempotency of cluster creation requests. If an active cluster with the provided token already exists, the request will not create a new cluster, but it will return the existing running cluster's ID instead. If you specify the idempotency token, upon failure, you can retry until the request succeeds. Databricks platform guarantees to launch exactly one cluster with that idempotency token. This token should have at most 64 characters.`,
				},
				resource.Attribute{
					Name:        "ssh_public_keys",
					Description: `(Optional) SSH public key contents that will be added to each Spark node in this cluster. The corresponding private keys can be used to login with the user name ubuntu on port 2200. You can specify up to 10 keys.`,
				},
				resource.Attribute{
					Name:        "spark_env_vars",
					Description: `(Optional) Map with environment variable key-value pairs to fine-tune Spark clusters. Key-value pairs of the form (X,Y) are exported (i.e., X='Y') while launching the driver and workers.`,
				},
				resource.Attribute{
					Name:        "custom_tags",
					Description: `(Optional) Additional tags for cluster resources. Databricks will tag all cluster resources (e.g., AWS EC2 instances and EBS volumes) with these tags in addition to ` + "`" + `default_tags` + "`" + `. If a custom cluster tag has the same name as a default cluster tag, the custom tag is prefixed with an ` + "`" + `x_` + "`" + ` when it is propagated.`,
				},
				resource.Attribute{
					Name:        "spark_conf",
					Description: `(Optional) Map with key-value pairs to fine-tune Spark clusters, where you can provide custom [Spark configuration properties](https://spark.apache.org/docs/latest/configuration.html) in a cluster configuration.`,
				},
				resource.Attribute{
					Name:        "is_pinned",
					Description: `(Optional) boolean value specifying if the cluster is pinned (not pinned by default). You must be a Databricks administrator to use this. The pinned clusters' maximum number is [limited to 70](https://docs.databricks.com/clusters/clusters-manage.html#pin-a-cluster), so ` + "`" + `apply` + "`" + ` may fail if you have more than that. The following example demonstrates how to create an autoscaling cluster with [Delta Cache](https://docs.databricks.com/delta/optimizations/delta-cache.html) enabled: ` + "`" + `` + "`" + `` + "`" + `hcl data "databricks_node_type" "smallest" { local_disk = true } data "databricks_spark_version" "latest_lts" { long_term_support = true } resource "databricks_cluster" "shared_autoscaling" { cluster_name = "Shared Autoscaling" spark_version = data.databricks_spark_version.latest_lts.id node_type_id = data.databricks_node_type.smallest.id autotermination_minutes = 20 autoscale { min_workers = 1 max_workers = 50 } spark_conf = { "spark.databricks.io.cache.enabled" : true, "spark.databricks.io.cache.maxDiskUsage" : "50g", "spark.databricks.io.cache.maxMetaDataCache" : "1g" } } ` + "`" + `` + "`" + `` + "`" + ` ## Fixed size or autoscaling cluster When you [create a Databricks cluster](https://docs.databricks.com/clusters/configure.html#cluster-size-and-autoscaling), you can either provide a ` + "`" + `num_workers` + "`" + ` for the fixed-size cluster or provide ` + "`" + `min_workers` + "`" + ` and/or ` + "`" + `max_workers` + "`" + ` for the cluster within the ` + "`" + `autoscale` + "`" + ` group. When you give a fixed-sized cluster, Databricks ensures that your cluster has a specified number of workers. When you provide a range for the number of workers, Databricks chooses the appropriate number of workers required to run your job - also known as "autoscaling." With autoscaling, Databricks dynamically reallocates workers to account for the characteristics of your job. Certain parts of your pipeline may be more computationally demanding than others, and Databricks automatically adds additional workers during these phases of your job (and removes them when they’re no longer needed). ` + "`" + `autoscale` + "`" + ` optional configuration block supports the following:`,
				},
				resource.Attribute{
					Name:        "min_workers",
					Description: `(Optional) The minimum number of workers to which the cluster can scale down when underutilized. It is also the initial number of workers the cluster will have after creation.`,
				},
				resource.Attribute{
					Name:        "max_workers",
					Description: `(Optional) The maximum number of workers to which the cluster can scale up when overloaded. max_workers must be strictly greater than min_workers. When using a [Single Node cluster](https://docs.databricks.com/clusters/single-node.html), ` + "`" + `num_workers` + "`" + ` needs to be ` + "`" + `0` + "`" + `. It can be set to ` + "`" + `0` + "`" + ` explicitly, or simply not specified, as it defaults to ` + "`" + `0` + "`" + `. When ` + "`" + `num_workers` + "`" + ` is ` + "`" + `0` + "`" + `, provider checks for presence of the required Spark configurations:`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `S3 destination, e.g., ` + "`" + `s3://my-bucket/some-prefix` + "`" + ` You must configure the cluster with an instance profile, and the instance profile must have write access to the destination. You cannot use AWS keys.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) S3 region, e.g. ` + "`" + `us-west-2` + "`" + `. Either ` + "`" + `region` + "`" + ` or ` + "`" + `endpoint` + "`" + ` must be set. If both are set, the endpoint is used.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) S3 endpoint, e.g. <https://s3-us-west-2.amazonaws.com>. Either ` + "`" + `region` + "`" + ` or ` + "`" + `endpoint` + "`" + ` needs to be set. If both are set, the endpoint is used.`,
				},
				resource.Attribute{
					Name:        "enable_encryption",
					Description: `(Optional) Enable server-side encryption, false by default.`,
				},
				resource.Attribute{
					Name:        "encryption_type",
					Description: `(Optional) The encryption type, it could be ` + "`" + `sse-s3` + "`" + ` or ` + "`" + `sse-kms` + "`" + `. It is used only when encryption is enabled, and the default type is ` + "`" + `sse-s3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kms_key",
					Description: `(Optional) KMS key used if encryption is enabled and encryption type is set to ` + "`" + `sse-kms` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "canned_acl",
					Description: `(Optional) Set canned access control list, e.g. ` + "`" + `bucket-owner-full-control` + "`" + `. If ` + "`" + `canned_cal` + "`" + ` is set, the cluster instance profile must have ` + "`" + `s3:PutObjectAcl` + "`" + ` permission on the destination bucket and prefix. The full list of possible canned ACLs can be found [here](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl). By default, only the object owner gets full control. If you are using a cross-account role for writing data, you may want to set ` + "`" + `bucket-owner-full-control` + "`" + ` to make bucket owners able to read the logs. ## init_scripts To run a particular init script on all clusters within the same workspace, both automated/job and interactive/all-purpose cluster types, please consider the [databricks_global_init_script](global_init_script.md) resource. It is possible to specify up to 10 different cluster-scoped init scripts per cluster. Init scripts support DBFS, cloud storage locations, and workspace files. Example of using a Databricks workspace file as init script: ` + "`" + `` + "`" + `` + "`" + `hcl init_scripts { workspace { destination = "/Users/user@domain/install-elk.sh" } } ` + "`" + `` + "`" + `` + "`" + ` Example of taking init script from DBFS: ` + "`" + `` + "`" + `` + "`" + `hcl init_scripts { dbfs { destination = "dbfs:/init-scripts/install-elk.sh" } } ` + "`" + `` + "`" + `` + "`" + ` Example of taking init script from S3: ` + "`" + `` + "`" + `` + "`" + `hcl init_scripts { s3 { destination = "s3://acmecorp-main/init-scripts/install-elk.sh" region = "us-east-1" } } ` + "`" + `` + "`" + `` + "`" + ` Similarly, for an init script stored in GCS: ` + "`" + `` + "`" + `` + "`" + `hcl init_scripts { gcs { destination = "gs://init-scripts/install-elk.sh" } } ` + "`" + `` + "`" + `` + "`" + ` Similarly, for an init script stored in ADLS: ` + "`" + `` + "`" + `` + "`" + `hcl init_scripts { abfss { destination = "abfss://container@storage.dfs.core.windows.net/install-elk.sh" } } ` + "`" + `` + "`" + `` + "`" + ` Please note that you need to provide Spark Hadoop configuration (` + "`" + `spark.hadoop.fs.azure...` + "`" + `) to authenticate to ADLS to get access to the init script. Clusters with [custom Docker containers](https://docs.databricks.com/clusters/custom-containers.html) also allow a local file location for init scripts as follows: ` + "`" + `` + "`" + `` + "`" + `hcl init_scripts { file { destination = "file:/my/local/file.sh" } } ` + "`" + `` + "`" + `` + "`" + ` ## aws_attributes ` + "`" + `aws_attributes` + "`" + ` optional configuration block contains attributes related to [clusters running on Amazon Web Services](https://docs.databricks.com/clusters/configure.html#aws-configurations). Here is the example of shared autoscaling cluster with some of AWS options set: ` + "`" + `` + "`" + `` + "`" + `hcl data "databricks_spark_version" "latest" {} data "databricks_node_type" "smallest" { local_disk = true } resource "databricks_cluster" "this" { cluster_name = "Shared Autoscaling" spark_version = data.databricks_spark_version.latest.id node_type_id = data.databricks_node_type.smallest.id autotermination_minutes = 20 autoscale { min_workers = 1 max_workers = 50 } aws_attributes { availability = "SPOT" zone_id = "us-east-1" first_on_demand = 1 spot_bid_price_percent = 100 } } ` + "`" + `` + "`" + `` + "`" + ` The following options are available:`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) Identifier for the availability zone/datacenter in which the cluster resides. This string will be of a form like ` + "`" + `us-west-2a` + "`" + `. The provided availability zone must be in the same region as the Databricks deployment. For example, ` + "`" + `us-west-2a` + "`" + ` is not a valid zone ID if the Databricks deployment resides in the ` + "`" + `us-east-1` + "`" + ` region. Enable automatic availability zone selection ("Auto-AZ"), by setting the value ` + "`" + `auto` + "`" + `. Databricks selects the AZ based on available IPs in the workspace subnets and retries in other availability zones if AWS returns insufficient capacity errors.`,
				},
				resource.Attribute{
					Name:        "availability",
					Description: `(Optional) Availability type used for all subsequent nodes past the ` + "`" + `first_on_demand` + "`" + ` ones. Valid values are ` + "`" + `SPOT` + "`" + `, ` + "`" + `SPOT_WITH_FALLBACK` + "`" + ` and ` + "`" + `ON_DEMAND` + "`" + `. Note: If ` + "`" + `first_on_demand` + "`" + ` is zero, this availability type will be used for the entire cluster. Backend default value is ` + "`" + `SPOT_WITH_FALLBACK` + "`" + ` and could change in the future`,
				},
				resource.Attribute{
					Name:        "first_on_demand",
					Description: `(Optional) The first ` + "`" + `first_on_demand` + "`" + ` nodes of the cluster will be placed on on-demand instances. If this value is greater than 0, the cluster driver node will be placed on an on-demand instance. If this value is greater than or equal to the current cluster size, all nodes will be placed on on-demand instances. If this value is less than the current cluster size, ` + "`" + `first_on_demand` + "`" + ` nodes will be placed on on-demand instances, and the remainder will be placed on availability instances. This value does not affect cluster size and cannot be mutated over the lifetime of a cluster. Backend default value is ` + "`" + `1` + "`" + ` and could change in the future`,
				},
				resource.Attribute{
					Name:        "spot_bid_price_percent",
					Description: `(Optional) The max price for AWS spot instances, as a percentage of the corresponding instance type’s on-demand price. For example, if this field is set to 50, and the cluster needs a new ` + "`" + `i3.xlarge` + "`" + ` spot instance, then the max price is half of the price of on-demand ` + "`" + `i3.xlarge` + "`" + ` instances. Similarly, if this field is set to 200, the max price is twice the price of on-demand ` + "`" + `i3.xlarge` + "`" + ` instances. If not specified, the default value is ` + "`" + `100` + "`" + `. When spot instances are requested for this cluster, only spot instances whose max price percentage matches this field will be considered. For safety, we enforce this field to be no more than ` + "`" + `10000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_profile_arn",
					Description: `(Optional) Nodes for this cluster will only be placed on AWS instances with this instance profile. Please see [databricks_instance_profile](instance_profile.md) resource documentation for extended examples on adding a valid instance profile using Terraform.`,
				},
				resource.Attribute{
					Name:        "ebs_volume_type",
					Description: `(Optional) The type of EBS volumes that will be launched with this cluster. Valid values are ` + "`" + `GENERAL_PURPOSE_SSD` + "`" + ` or ` + "`" + `THROUGHPUT_OPTIMIZED_HDD` + "`" + `. Use this option only if you're not picking`,
				},
				resource.Attribute{
					Name:        "ebs_volume_count",
					Description: `(Optional) The number of volumes launched for each instance. You can choose up to 10 volumes. This feature is only enabled for supported node types. Legacy node types cannot specify custom EBS volumes. For node types with no instance store, at least one EBS volume needs to be specified; otherwise, cluster creation will fail. These EBS volumes will be mounted at /ebs0, /ebs1, and etc. Instance store volumes will be mounted at /local_disk0, /local_disk1, and etc. If EBS volumes are attached, Databricks will configure Spark to use only the EBS volumes for scratch storage because heterogeneously sized scratch devices can lead to inefficient disk utilization. If no EBS volumes are attached, Databricks will configure Spark to use instance store volumes. If EBS volumes are specified, then the Spark configuration spark.local.dir will be overridden.`,
				},
				resource.Attribute{
					Name:        "ebs_volume_size",
					Description: `(Optional) The size of each EBS volume (in GiB) launched for each instance. For general purpose SSD, this value must be within the range 100 - 4096. For throughput optimized HDD, this value must be within the range 500 - 4096. Custom EBS volumes cannot be specified for the legacy node types (memory-optimized and compute-optimized). ## azure_attributes ` + "`" + `azure_attributes` + "`" + ` optional configuration block contains attributes related to [clusters running on Azure](https://docs.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/clusters#--azureattributes). Here is the example of shared autoscaling cluster with some of Azure options set: ` + "`" + `` + "`" + `` + "`" + `hcl data "databricks_spark_version" "latest" {} data "databricks_node_type" "smallest" { local_disk = true } resource "databricks_cluster" "this" { cluster_name = "Shared Autoscaling" spark_version = data.databricks_spark_version.latest.id node_type_id = data.databricks_node_type.smallest.id autotermination_minutes = 20 autoscale { min_workers = 1 max_workers = 50 } azure_attributes { availability = "SPOT_WITH_FALLBACK_AZURE" first_on_demand = 1 spot_bid_max_price = 100 } } ` + "`" + `` + "`" + `` + "`" + ` The following options are [available](https://docs.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/clusters#--azureattributes):`,
				},
				resource.Attribute{
					Name:        "availability",
					Description: `(Optional) Availability type used for all subsequent nodes past the ` + "`" + `first_on_demand` + "`" + ` ones. Valid values are ` + "`" + `SPOT_AZURE` + "`" + `, ` + "`" + `SPOT_WITH_FALLBACK_AZURE` + "`" + `, and ` + "`" + `ON_DEMAND_AZURE` + "`" + `. Note: If ` + "`" + `first_on_demand` + "`" + ` is zero, this availability type will be used for the entire cluster.`,
				},
				resource.Attribute{
					Name:        "first_on_demand",
					Description: `(Optional) The first ` + "`" + `first_on_demand` + "`" + ` nodes of the cluster will be placed on on-demand instances. If this value is greater than 0, the cluster driver node will be placed on an on-demand instance. If this value is greater than or equal to the current cluster size, all nodes will be placed on on-demand instances. If this value is less than the current cluster size, ` + "`" + `first_on_demand` + "`" + ` nodes will be placed on on-demand instances, and the remainder will be placed on availability instances. This value does not affect cluster size and cannot be mutated over the lifetime of a cluster.`,
				},
				resource.Attribute{
					Name:        "spot_bid_max_price",
					Description: `(Optional) The max price for Azure spot instances. Use ` + "`" + `-1` + "`" + ` to specify the lowest price. ## gcp_attributes ` + "`" + `gcp_attributes` + "`" + ` optional configuration block contains attributes related to [clusters running on GCP](https://docs.gcp.databricks.com/dev-tools/api/latest/clusters.html#clustergcpattributes). Here is the example of shared autoscaling cluster with some of GCP options set: ` + "`" + `` + "`" + `` + "`" + `hcl resource "databricks_cluster" "this" { cluster_name = "Shared Autoscaling" spark_version = data.databricks_spark_version.latest.id node_type_id = data.databricks_node_type.smallest.id autotermination_minutes = 20 autoscale { min_workers = 1 max_workers = 50 } gcp_attributes { availability = "PREEMPTIBLE_WITH_FALLBACK_GCP" zone_id = "AUTO" } } ` + "`" + `` + "`" + `` + "`" + ` The following options are available:`,
				},
				resource.Attribute{
					Name:        "use_preemptible_executors",
					Description: `(Optional, bool) if we should use preemptible executors ([GCP documentation](https://cloud.google.com/compute/docs/instances/preemptible)).`,
				},
				resource.Attribute{
					Name:        "google_service_account",
					Description: `(Optional, string) Google Service Account email address that the cluster uses to authenticate with Google Identity. This field is used for authentication with the GCS and BigQuery data sources.`,
				},
				resource.Attribute{
					Name:        "availability",
					Description: `(Optional) Availability type used for all nodes. Valid values are ` + "`" + `PREEMPTIBLE_GCP` + "`" + `, ` + "`" + `PREEMPTIBLE_WITH_FALLBACK_GCP` + "`" + ` and ` + "`" + `ON_DEMAND_GCP` + "`" + `, default: ` + "`" + `ON_DEMAND_GCP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL for the Docker image`,
				},
				resource.Attribute{
					Name:        "basic_auth",
					Description: `(Optional) ` + "`" + `basic_auth.username` + "`" + ` and ` + "`" + `basic_auth.password` + "`" + ` for Docker repository. Docker registry credentials are encrypted when they are stored in Databricks internal storage and when they are passed to a registry upon fetching Docker images at cluster launch. However, other authenticated and authorized API users of this workspace can access the username and password. Example usage with [azurerm_container_registry](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/container_registry) and [docker_registry_image](https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs/resources/registry_image), that you can adapt to your specific use-case: ` + "`" + `` + "`" + `` + "`" + `hcl resource "docker_registry_image" "this" { name = "${azurerm_container_registry.this.login_server}/sample:latest" build { # ... } } resource "databricks_cluster" "this" { # ... docker_image { url = docker_registry_image.this.name basic_auth { username = azurerm_container_registry.this.admin_username password = azurerm_container_registry.this.admin_password } } } ` + "`" + `` + "`" + `` + "`" + ` ## cluster_mount_info blocks It's possible to mount NFS (Network File System) resources into the Spark containers inside the cluster. You can specify one or more ` + "`" + `cluster_mount_info` + "`" + ` blocks describing the mount. This block has following attributes:`,
				},
				resource.Attribute{
					Name:        "network_filesystem_info",
					Description: `block specifying connection. It consists of:`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Required) host name.`,
				},
				resource.Attribute{
					Name:        "mount_options",
					Description: `(Optional) string that will be passed as options passed to the ` + "`" + `mount` + "`" + ` command.`,
				},
				resource.Attribute{
					Name:        "remote_mount_dir_path",
					Description: `(Optional) string specifying path to mount on the remote service.`,
				},
				resource.Attribute{
					Name:        "local_mount_dir_path",
					Description: `(Required) path inside the Spark container. For example, you can mount Azure Data Lake Storage container using the following code: ` + "`" + `` + "`" + `` + "`" + `hcl locals { storage_account = "ewfw3ggwegwg" storage_container = "test" } resource "databricks_cluster" "with_nfs" { # ... cluster_mount_info { network_filesystem_info { server_address = "${local.storage_account}.blob.core.windows.net" mount_options = "sec=sys,vers=3,nolock,proto=tcp" } remote_mount_dir_path = "${local.storage_account}/${local.storage_container}" local_mount_dir_path = "/mnt/nfs-test" } } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the cluster.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `(map) Tags that are added by Databricks by default, regardless of any ` + "`" + `custom_tags` + "`" + ` that may have been added. These include: Vendor: Databricks, Creator: <username_of_creator>, ClusterName: <name_of_cluster>, ClusterId: <id_of_cluster>, Name: <Databricks internal use>, and any workspace and pool tags.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(string) State of the cluster. ## Access Control`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the cluster.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `(map) Tags that are added by Databricks by default, regardless of any ` + "`" + `custom_tags` + "`" + ` that may have been added. These include: Vendor: Databricks, Creator: <username_of_creator>, ClusterName: <name_of_cluster>, ClusterId: <id_of_cluster>, Name: <Databricks internal use>, and any workspace and pool tags.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(string) State of the cluster. ## Access Control`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_cluster_policy",
			Category:         "Compute",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"compute",
				"cluster",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cluster policy name. This must be unique. Length must be between 1 and 100 characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Additional human-readable description of the cluster policy.`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Policy definition: JSON document expressed in [Databricks Policy Definition Language](https://docs.databricks.com/administration-guide/clusters/policies.html#cluster-policy-definition). Cannot be used with ` + "`" + `policy_family_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "max_clusters_per_user",
					Description: `(Optional, integer) Maximum number of clusters allowed per user. When omitted, there is no limit. If specified, value must be greater than zero.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the cluster policy. This is equal to policy_id.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Canonical unique identifier for the cluster policy. ## Import The resource cluster policy can be imported using the policy id: ` + "`" + `` + "`" + `` + "`" + `bash terraform import databricks_cluster_policy.this <cluster-policy-id> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the cluster policy. This is equal to policy_id.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Canonical unique identifier for the cluster policy. ## Import The resource cluster policy can be imported using the policy id: ` + "`" + `` + "`" + `` + "`" + `bash terraform import databricks_cluster_policy.this <cluster-policy-id> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_dbfs_file",
			Category:         "Storage",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"storage",
				"dbfs",
				"file",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source",
					Description: `The full absolute path to the file. Conflicts with ` + "`" + `content_base64` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `Encoded file contents. Conflicts with ` + "`" + `source` + "`" + `. Use of ` + "`" + `content_base64` + "`" + ` is discouraged, as it's increasing memory footprint of Terraform state and should only be used in exceptional circumstances, like creating a data pipeline configuration file.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path of the file in which you wish to save. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Same as ` + "`" + `path` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_size",
					Description: `The file size of the file that is being tracked by this resource in bytes.`,
				},
				resource.Attribute{
					Name:        "dbfs_path",
					Description: `Path, but with ` + "`" + `dbfs:` + "`" + ` prefix. ## Import The resource dbfs file can be imported using the path of the file: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_dbfs_file.this <path> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Same as ` + "`" + `path` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_size",
					Description: `The file size of the file that is being tracked by this resource in bytes.`,
				},
				resource.Attribute{
					Name:        "dbfs_path",
					Description: `Path, but with ` + "`" + `dbfs:` + "`" + ` prefix. ## Import The resource dbfs file can be imported using the path of the file: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_dbfs_file.this <path> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_directory",
			Category:         "Workspace",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"workspace",
				"directory",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_entitlements",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"entitlements",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `Canonical unique identifier for the user.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Canonical unique identifier for the group.`,
				},
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `Canonical unique identifier for the service principal. The following entitlements are available.`,
				},
				resource.Attribute{
					Name:        "allow_cluster_create",
					Description: `(Optional) Allow the principal to have [cluster](cluster.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Cluster-usage) and ` + "`" + `cluster_id` + "`" + ` argument. Everyone without ` + "`" + `allow_cluster_create` + "`" + ` argument set, but with [permission to use](permissions.md#Cluster-Policy-usage) Cluster Policy would be able to create clusters, but within boundaries of that specific policy.`,
				},
				resource.Attribute{
					Name:        "allow_instance_pool_create",
					Description: `(Optional) Allow the principal to have [instance pool](instance_pool.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Instance-Pool-usage) and [instance_pool_id](permissions.md#instance_pool_id) argument.`,
				},
				resource.Attribute{
					Name:        "databricks_sql_access",
					Description: `(Optional) This is a field to allow the principal to have access to [Databricks SQL](https://databricks.com/product/databricks-sql) feature in User Interface and through [databricks_sql_endpoint](sql_endpoint.md).`,
				},
				resource.Attribute{
					Name:        "workspace_access",
					Description: `(Optional) This is a field to allow the principal to have access to Databricks Workspace. ## Import The resource can be imported using a synthetic identifier. Examples of valid synthetic identifiers are:`,
				},
				resource.Attribute{
					Name:        "user/user_id",
					Description: `user ` + "`" + `user_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group/group_id",
					Description: `group ` + "`" + `group_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spn/spn_id",
					Description: `service principal ` + "`" + `spn_id` + "`" + `. ` + "`" + `` + "`" + `` + "`" + `bash terraform import databricks_entitlements.me user/<user-id> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_external_location",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"external",
				"location",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_git_credential",
			Category:         "Workspace",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"workspace",
				"git",
				"credential",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "personal_access_token",
					Description: `(Required) The personal access token used to authenticate to the corresponding Git provider. If value is not provided, it's sourced from the first environment variable of [` + "`" + `GITHUB_TOKEN` + "`" + `](https://registry.terraform.io/providers/integrations/github/latest/docs#oauth--personal-access-token), [` + "`" + `GITLAB_TOKEN` + "`" + `](https://registry.terraform.io/providers/gitlabhq/gitlab/latest/docs#required), or [` + "`" + `AZDO_PERSONAL_ACCESS_TOKEN` + "`" + `](https://registry.terraform.io/providers/microsoft/azuredevops/latest/docs#argument-reference), that has a non-empty value.`,
				},
				resource.Attribute{
					Name:        "git_username",
					Description: `(Required) user name at Git provider.`,
				},
				resource.Attribute{
					Name:        "git_provider",
					Description: `(Required) case insensitive name of the Git provider. Following values are supported right now (could be a subject for a change, consult [Git Credentials API documentation](https://docs.databricks.com/dev-tools/api/latest/gitcredentials.html)): ` + "`" + `gitHub` + "`" + `, ` + "`" + `gitHubEnterprise` + "`" + `, ` + "`" + `bitbucketCloud` + "`" + `, ` + "`" + `bitbucketServer` + "`" + `, ` + "`" + `azureDevOpsServices` + "`" + `, ` + "`" + `gitLab` + "`" + `, ` + "`" + `gitLabEnterpriseEdition` + "`" + `, ` + "`" + `awsCodeCommit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) specify if settings need to be enforced - right now, Databricks allows only single Git credential, so if it's already configured, the apply operation will fail. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `identifier of specific Git credential ## Import The resource cluster can be imported using ID of Git credential that could be obtained via REST API: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_git_credential.this <git-credential-id> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `identifier of specific Git credential ## Import The resource cluster can be imported using ID of Git credential that could be obtained via REST API: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_git_credential.this <git-credential-id> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_global_init_script",
			Category:         "Workspace",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"workspace",
				"global",
				"init",
				"script",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source",
					Description: `Path to script's source code on local filesystem. Conflicts with ` + "`" + `content_base64` + "`" + ``,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `The base64-encoded source code global init script. Conflicts with ` + "`" + `source` + "`" + `. Use of ` + "`" + `content_base64` + "`" + ` is discouraged, as it's increasing memory footprint of Terraform state and should only be used in exceptional circumstances`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID assigned to a global init script by API ## Access Control Global init scripts are available only for administrators, so you can't change permissions for it. ## Import The resource global init script can be imported using script ID: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_global_init_script.this script_id ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID assigned to a global init script by API ## Access Control Global init scripts are available only for administrators, so you can't change permissions for it. ## Import The resource global init script can be imported using script ID: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_global_init_script.this script_id ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_grants",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"grants",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_group",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) This is the display name for the given group.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Optional) ID of the group in an external identity provider.`,
				},
				resource.Attribute{
					Name:        "allow_cluster_create",
					Description: `(Optional) This is a field to allow the group to have [cluster](cluster.md) create privileges. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Cluster-usage) and [cluster_id](permissions.md#cluster_id) argument. Everyone without ` + "`" + `allow_cluster_create` + "`" + ` argument set, but with [permission to use](permissions.md#Cluster-Policy-usage) Cluster Policy would be able to create clusters, but within boundaries of that specific policy.`,
				},
				resource.Attribute{
					Name:        "allow_instance_pool_create",
					Description: `(Optional) This is a field to allow the group to have [instance pool](instance_pool.md) create privileges. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Instance-Pool-usage) and [instance_pool_id](permissions.md#instance_pool_id) argument.`,
				},
				resource.Attribute{
					Name:        "databricks_sql_access",
					Description: `(Optional) This is a field to allow the group to have access to [Databricks SQL](https://databricks.com/product/databricks-sql) feature in User Interface and through [databricks_sql_endpoint](sql_endpoint.md).`,
				},
				resource.Attribute{
					Name:        "workspace_access",
					Description: `(Optional) This is a field to allow the group to have access to Databricks Workspace.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) Ignore ` + "`" + `cannot create group: Group with name X already exists.` + "`" + ` errors and implicitly import the specific group into Terraform state, enforcing entitlements defined in the instance of resource. _This functionality is experimental_ and is designed to simplify corner cases, like Azure Active Directory synchronisation. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id for the group object. ## Import You can import a ` + "`" + `databricks_group` + "`" + ` resource with the name ` + "`" + `my_group` + "`" + ` like the following: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_group.my_group <group_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id for the group object. ## Import You can import a ` + "`" + `databricks_group` + "`" + ` resource with the name ` + "`" + `my_group` + "`" + ` like the following: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_group.my_group <group_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_group_instance_profile",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
				"instance",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) This is the id of the [group](group.md) resource.`,
				},
				resource.Attribute{
					Name:        "instance_profile_id",
					Description: `(Required) This is the id of the [instance profile](instance_profile.md) resource. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id in the format ` + "`" + `<group_id>|<instance_profile_id>` + "`" + `. ## Import ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id in the format ` + "`" + `<group_id>|<instance_profile_id>` + "`" + `. ## Import ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_group_member",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) This is the id of the [group](group.md) resource.`,
				},
				resource.Attribute{
					Name:        "member_id",
					Description: `(Required) This is the id of the [group](group.md), [service principal](service_principal.md), or [user](user.md). ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id for the ` + "`" + `databricks_group_member` + "`" + ` object which is in the format ` + "`" + `<group_id>|<member_id>` + "`" + `. ## Import ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id for the ` + "`" + `databricks_group_member` + "`" + ` object which is in the format ` + "`" + `<group_id>|<member_id>` + "`" + `. ## Import ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_group_role",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) This is the id of the [group](group.md) resource.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) Either a role name or the ARN/ID of the [instance profile](instance_profile.md) resource. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id for the ` + "`" + `databricks_group_role` + "`" + ` object which is in the format ` + "`" + `<group_id>|<role>` + "`" + `. ## Import ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id for the ` + "`" + `databricks_group_role` + "`" + ` object which is in the format ` + "`" + `<group_id>|<role>` + "`" + `. ## Import ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_instance_pool",
			Category:         "Compute",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"compute",
				"instance",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_pool_name",
					Description: `(Required) (String) The name of the instance pool. This is required for create and edit operations. It must be unique, non-empty, and less than 100 characters.`,
				},
				resource.Attribute{
					Name:        "min_idle_instances",
					Description: `(Optional) (Integer) The minimum number of idle instances maintained by the pool. This is in addition to any instances in use by active clusters.`,
				},
				resource.Attribute{
					Name:        "max_capacity",
					Description: `(Optional) (Integer) The maximum number of instances the pool can contain, including both idle instances and ones in use by clusters. Once the maximum capacity is reached, you cannot create new clusters from the pool and existing clusters cannot autoscale up until some instances are made idle in the pool via [cluster](cluster.md) termination or down-scaling. There is no default limit, but as a [best practice](https://docs.databricks.com/clusters/instance-pools/pool-best-practices.html#configure-pools-to-control-cost), this should be set based on anticipated usage.`,
				},
				resource.Attribute{
					Name:        "idle_instance_autotermination_minutes",
					Description: `(Required) (Integer) The number of minutes that idle instances in excess of the min_idle_instances are maintained by the pool before being terminated. If not specified, excess idle instances are terminated automatically after a default timeout period. If specified, the time must be between 0 and 10000 minutes. If you specify 0, excess idle instances are removed as soon as possible.`,
				},
				resource.Attribute{
					Name:        "node_type_id",
					Description: `(Required) (String) The node type for the instances in the pool. All clusters attached to the pool inherit this node type and the pool’s idle instances are allocated based on this type. You can retrieve a list of available node types by using the [List Node Types API](https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterclusterservicelistnodetypes) call.`,
				},
				resource.Attribute{
					Name:        "custom_tags",
					Description: `(Optional) (Map) Additional tags for instance pool resources. Databricks tags all pool resources (e.g. AWS & Azure instances and Disk volumes). The tags of the instance pool will propagate to the clusters using the pool (see the [official documentation](https://docs.databricks.com/administration-guide/account-settings/usage-detail-tags-aws.html#tag-propagation)). Attempting to set the same tags in both cluster and instance pool will raise an error.`,
				},
				resource.Attribute{
					Name:        "enable_elastic_disk",
					Description: `(Optional) (Bool) Autoscaling Local Storage: when enabled, the instances in the pool dynamically acquire additional disk space when they are running low on disk space.`,
				},
				resource.Attribute{
					Name:        "preloaded_spark_versions",
					Description: `(Optional) (List) A list with at most one runtime version the pool installs on each instance. Pool clusters that use a preloaded runtime version start faster as they do not have to wait for the image to download. You can retrieve them via [databricks_spark_version](../data-sources/spark-version.md) data source or via [Runtime Versions API](https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterclusterservicelistsparkversions) call. ### aws_attributes Configuration Block The following options are [available](https://docs.databricks.com/dev-tools/api/latest/instance-pools.html#clusterinstancepoolawsattributes):`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) (String) Identifier for the availability zone/datacenter in which the instance pool resides. This string is of the form like ` + "`" + `"us-west-2a"` + "`" + `. The provided availability zone must be in the same region as the Databricks deployment. For example, ` + "`" + `"us-west-2a"` + "`" + ` is not a valid zone ID if the Databricks deployment resides in the ` + "`" + `"us-east-1"` + "`" + ` region. If not specified, a default zone is used. You can find the list of available zones as well as the default value by using the [List Zones API](https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterclusterservicelistavailablezones).`,
				},
				resource.Attribute{
					Name:        "spot_bid_price_percent",
					Description: `(Optional) (Integer) The max price for AWS spot instances, as a percentage of the corresponding instance type’s on-demand price. For example, if this field is set to 50, and the instance pool needs a new i3.xlarge spot instance, then the max price is half of the price of on-demand i3.xlarge instances. Similarly, if this field is set to 200, the max price is twice the price of on-demand i3.xlarge instances. If not specified, the`,
				},
				resource.Attribute{
					Name:        "availability",
					Description: `(Optional) (String) Availability type used for all instances in the pool. Only ` + "`" + `ON_DEMAND` + "`" + ` and ` + "`" + `SPOT` + "`" + ` are supported. ## azure_attributes Configuration Block ` + "`" + `azure_attributes` + "`" + ` optional configuration block contains attributes related to [instance pools on Azure](https://docs.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/instance-pools#--instancepoolazureattributes). The following options are [available](https://docs.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/clusters#--azureattributes):`,
				},
				resource.Attribute{
					Name:        "availability",
					Description: `(Optional) Availability type used for all nodes. Valid values are ` + "`" + `SPOT_AZURE` + "`" + ` and ` + "`" + `ON_DEMAND_AZURE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_bid_max_price",
					Description: `(Optional) The max price for Azure spot instances. Use ` + "`" + `-1` + "`" + ` to specify the lowest price. ## gcp_attributes Configuration Block ` + "`" + `gcp_attributes` + "`" + ` optional configuration block contains attributes related to [instance pools on GCP](https://docs.gcp.databricks.com/dev-tools/api/latest/instance-pools.html#clusterinstancepoolgcpattributes). The following options are [available](https://docs.gcp.databricks.com/dev-tools/api/latest/clusters.html#gcpavailability):`,
				},
				resource.Attribute{
					Name:        "availability",
					Description: `(Optional) Availability type used for all nodes. Valid values are ` + "`" + `PREEMPTIBLE_GCP` + "`" + `, ` + "`" + `PREEMPTIBLE_WITH_FALLBACK_GCP` + "`" + ` and ` + "`" + `ON_DEMAND_GCP` + "`" + `, default: ` + "`" + `ON_DEMAND_GCP` + "`" + `. ### disk_spec Configuration Block For disk_spec make sure to use`,
				},
				resource.Attribute{
					Name:        "disk_count",
					Description: `(Optional) (Integer) The number of disks to attach to each instance. This feature is only enabled for supported node types. Users can choose up to the limit of the disks supported by the node type. For node types with no local disk, at least one disk needs to be specified.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) (Integer) The size of each disk (in GiB) to attach. #### disk_type sub-block ` + "`" + `ebs_volume_type` + "`" + ` - (Optional) (String) The EBS volume type to use. Options are: ` + "`" + `GENERAL_PURPOSE_SSD` + "`" + ` (Provision extra storage using AWS gp2 EBS volumes) or ` + "`" + `THROUGHPUT_OPTIMIZED_HDD` + "`" + ` (Provision extra storage using AWS st1 volumes)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL for the Docker image`,
				},
				resource.Attribute{
					Name:        "basic_auth",
					Description: `(Optional) ` + "`" + `basic_auth.username` + "`" + ` and ` + "`" + `basic_auth.password` + "`" + ` for Docker repository. Docker registry credentials are encrypted when they are stored in Databricks internal storage and when they are passed to a registry upon fetching Docker images at cluster launch. However, other authenticated and authorized API users of this workspace can access the username and password. Example usage with [azurerm_container_registry](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/container_registry) and [docker_registry_image](https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs/resources/registry_image), that you can adapt to your specific use-case: ` + "`" + `` + "`" + `` + "`" + `hcl resource "docker_registry_image" "this" { name = "${azurerm_container_registry.this.login_server}/sample:latest" build { # ... } } resource "databricks_instance_pool" "this" { # ... preloaded_docker_image { url = docker_registry_image.this.name basic_auth { username = azurerm_container_registry.this.admin_username password = azurerm_container_registry.this.admin_password } } } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the instance pool. ## Access Control`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the instance pool. ## Access Control`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_instance_profile",
			Category:         "Deployment",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"instance",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_profile_arn",
					Description: `(Required) ` + "`" + `ARN` + "`" + ` attribute of ` + "`" + `aws_iam_instance_profile` + "`" + ` output, the EC2 instance profile association to AWS IAM role. This ARN would be validated upon resource creation.`,
				},
				resource.Attribute{
					Name:        "iam_role_arn",
					Description: `(Optional) The AWS IAM role ARN of the role associated with the instance profile. It must have the form ` + "`" + `arn:aws:iam::<account-id>:role/<name>` + "`" + `. This field is required if your role name and instance profile name do not match and you want to use the instance profile with Databricks SQL Serverless.`,
				},
				resource.Attribute{
					Name:        "is_meta_instance_profile",
					Description: `(Optional) Whether the instance profile is a meta instance profile. Used only in [IAM credential passthrough](https://docs.databricks.com/security/credential-passthrough/iam-passthrough.html).`,
				},
				resource.Attribute{
					Name:        "skip_validation",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ARN for EC2 Instance Profile, that is registered with Databricks. ## Import The resource instance profile can be imported using the ARN of it ` + "`" + `` + "`" + `` + "`" + `bash terraform import databricks_instance_profile.this <instance-profile-arn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ARN for EC2 Instance Profile, that is registered with Databricks. ## Import The resource instance profile can be imported using the ARN of it ` + "`" + `` + "`" + `` + "`" + `bash terraform import databricks_instance_profile.this <instance-profile-arn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_ip_access_list",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"ip",
				"access",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "list_type",
					Description: `Can only be "ALLOW" or "BLOCK".`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `A string list of IP addresses and CIDR ranges.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `This is the display name for the given IP ACL List.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Boolean ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + ` indicating whether this list should be active. Defaults to ` + "`" + `true` + "`" + ` ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list_id",
					Description: `Canonical unique identifier for the IP Access List. ## Import The databricks_ip_access_list can be imported using id: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_ip_access_list.this <list-id> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list_id",
					Description: `Canonical unique identifier for the IP Access List. ## Import The databricks_ip_access_list can be imported using id: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_ip_access_list.this <list-id> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_job",
			Category:         "Compute",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"compute",
				"job",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) An optional name for the job. The default value is Untitled.`,
				},
				resource.Attribute{
					Name:        "job_cluster",
					Description: `(Optional) A list of job [databricks_cluster](cluster.md) specifications that can be shared and reused by tasks of this job. Libraries cannot be declared in a shared job cluster. You must declare dependent libraries in task settings.`,
				},
				resource.Attribute{
					Name:        "always_running",
					Description: `(Optional) (Bool) Whenever the job is always running, like a Spark Streaming application, on every update restart the current active run or start it again, if nothing it is not running. False by default. Any job runs are started with ` + "`" + `parameters` + "`" + ` specified in ` + "`" + `spark_jar_task` + "`" + ` or ` + "`" + `spark_submit_task` + "`" + ` or ` + "`" + `spark_python_task` + "`" + ` or ` + "`" + `notebook_task` + "`" + ` blocks.`,
				},
				resource.Attribute{
					Name:        "library",
					Description: `(Optional) (Set) An optional list of libraries to be installed on the cluster that will execute the job. Please consult [libraries section](cluster.md#libraries) for [databricks_cluster](cluster.md) resource.`,
				},
				resource.Attribute{
					Name:        "retry_on_timeout",
					Description: `(Optional) (Bool) An optional policy to specify whether to retry a job when it times out. The default behavior is to not retry on timeout.`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `(Optional) (Integer) An optional maximum number of times to retry an unsuccessful run. A run is considered to be unsuccessful if it completes with a FAILED or INTERNAL_ERROR lifecycle state. The value -1 means to retry indefinitely and the value 0 means to never retry. The default behavior is to never retry. A run can have the following lifecycle state: PENDING, RUNNING, TERMINATING, TERMINATED, SKIPPED or INTERNAL_ERROR`,
				},
				resource.Attribute{
					Name:        "timeout_seconds",
					Description: `(Optional) (Integer) An optional timeout applied to each run of this job. The default behavior is to have no timeout.`,
				},
				resource.Attribute{
					Name:        "min_retry_interval_millis",
					Description: `(Optional) (Integer) An optional minimal interval in milliseconds between the start of the failed run and the subsequent retry run. The default behavior is that unsuccessful runs are immediately retried.`,
				},
				resource.Attribute{
					Name:        "max_concurrent_runs",
					Description: `(Optional) (Integer) An optional maximum allowed number of concurrent runs of the job. Defaults to`,
				},
				resource.Attribute{
					Name:        "email_notifications",
					Description: `(Optional) (List) An optional set of email addresses notified when runs of this job begins, completes and fails. The default behavior is to not send any emails. This field is a block and is documented below.`,
				},
				resource.Attribute{
					Name:        "webhook_notifications",
					Description: `(Optional) (List) An optional set of system destinations (for example, webhook destinations or Slack) to be notified when runs of this job begins, completes and fails. The default behavior is to not send any notifications. This field is a block and is documented below.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Optional) (List) An optional periodic schedule for this job. The default behavior is that the job runs when triggered by clicking Run Now in the Jobs UI or sending an API request to runNow. This field is a block and is documented below. ### tags Configuration Map ` + "`" + `tags` + "`" + ` - (Optional) (Map) An optional map of the tags associated with the job. Specified tags will be used as cluster tags for job clusters. Example ` + "`" + `` + "`" + `` + "`" + `hcl tags = { environment = "dev" owner = "dream-team" } ` + "`" + `` + "`" + `` + "`" + ` ### job_cluster Configuration Block [Shared job cluster](https://docs.databricks.com/jobs.html#use-shared-job-clusters) specification. Allows multiple tasks in the same job run to reuse the cluster.`,
				},
				resource.Attribute{
					Name:        "job_cluster_key",
					Description: `(Required) Identifier that can be referenced in ` + "`" + `task` + "`" + ` block, so that cluster is shared between tasks`,
				},
				resource.Attribute{
					Name:        "new_cluster",
					Description: `Same set of parameters as for [databricks_cluster](cluster.md) resource. ### schedule Configuration Block`,
				},
				resource.Attribute{
					Name:        "quartz_cron_expression",
					Description: `(Required) A [Cron expression using Quartz syntax](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) that describes the schedule for a job. This field is required.`,
				},
				resource.Attribute{
					Name:        "timezone_id",
					Description: `(Required) A Java timezone ID. The schedule for a job will be resolved with respect to this timezone. See Java TimeZone for details. This field is required.`,
				},
				resource.Attribute{
					Name:        "pause_status",
					Description: `(Optional) Indicate whether this schedule is paused or not. Either ` + "`" + `PAUSED` + "`" + ` or ` + "`" + `UNPAUSED` + "`" + `. When the ` + "`" + `pause_status` + "`" + ` field is omitted and a schedule is provided, the server will default to using ` + "`" + `UNPAUSED` + "`" + ` as a value for ` + "`" + `pause_status` + "`" + `. ### continuous Configuration Block`,
				},
				resource.Attribute{
					Name:        "pause_status",
					Description: `(Optional) Indicate whether this continuous job is paused or not. Either ` + "`" + `PAUSED` + "`" + ` or ` + "`" + `UNPAUSED` + "`" + `. When the ` + "`" + `pause_status` + "`" + ` field is omitted in the block, the server will default to using ` + "`" + `UNPAUSED` + "`" + ` as a value for ` + "`" + `pause_status` + "`" + `. ### trigger Configuration Block`,
				},
				resource.Attribute{
					Name:        "pause_status",
					Description: `(Optional) Indicate whether this trigger is paused or not. Either ` + "`" + `PAUSED` + "`" + ` or ` + "`" + `UNPAUSED` + "`" + `. When the ` + "`" + `pause_status` + "`" + ` field is omitted in the block, the server will default to using ` + "`" + `UNPAUSED` + "`" + ` as a value for ` + "`" + `pause_status` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_arrival",
					Description: `(Required) configuration block to define a trigger for [File Arrival events](https://learn.microsoft.com/en-us/azure/databricks/workflows/jobs/file-arrival-triggers) consisting of following attributes:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) string with URL under the Unity Catalog external location that will be monitored for new files. Please note that have a trailing slash character (` + "`" + `/` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "min_time_between_trigger_seconds",
					Description: `(Optional) If set, the trigger starts a run only after the specified amount of time passed since the last time the trigger fired. The minimum allowed value is 60 seconds.`,
				},
				resource.Attribute{
					Name:        "wait_after_last_change_seconds",
					Description: `(Optional) If set, the trigger starts a run only after no file activity has occurred for the specified amount of time. This makes it possible to wait for a batch of incoming files to arrive before triggering a run. The minimum allowed value is 60 seconds. ### git_source Configuration Block This block is used to specify Git repository information & branch/tag/commit that will be used to pull source code from to execute a job. Supported options are:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) URL of the Git repository to use.`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional, if it's possible to detect Git provider by host name) case insensitive name of the Git provider. Following values are supported right now (could be a subject for change, consult [Repos API documentation](https://docs.databricks.com/dev-tools/api/latest/repos.html)): ` + "`" + `gitHub` + "`" + `, ` + "`" + `gitHubEnterprise` + "`" + `, ` + "`" + `bitbucketCloud` + "`" + `, ` + "`" + `bitbucketServer` + "`" + `, ` + "`" + `azureDevOpsServices` + "`" + `, ` + "`" + `gitLab` + "`" + `, ` + "`" + `gitLabEnterpriseEdition` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `name of the Git branch to use. Conflicts with ` + "`" + `tag` + "`" + ` and ` + "`" + `commit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `name of the Git branch to use. Conflicts with ` + "`" + `branch` + "`" + ` and ` + "`" + `commit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "commit",
					Description: `hash of Git commit to use. Conflicts with ` + "`" + `branch` + "`" + ` and ` + "`" + `tag` + "`" + `. ### email_notifications Configuration Block`,
				},
				resource.Attribute{
					Name:        "on_start",
					Description: `(Optional) (List) list of emails to notify when the run starts.`,
				},
				resource.Attribute{
					Name:        "on_success",
					Description: `(Optional) (List) list of emails to notify when the run completes successfully.`,
				},
				resource.Attribute{
					Name:        "on_failure",
					Description: `(Optional) (List) list of emails to notify when the run fails.`,
				},
				resource.Attribute{
					Name:        "no_alert_for_skipped_runs",
					Description: `(Optional) (Bool) don't send alert for skipped runs. (It's recommended to use the corresponding setting in the ` + "`" + `notification_settings` + "`" + ` configuration block). ### webhook_notifications Configuration Block Each entry in ` + "`" + `webhook_notification` + "`" + ` block takes a list ` + "`" + `webhook` + "`" + ` blocks. The field is documented below.`,
				},
				resource.Attribute{
					Name:        "on_start",
					Description: `(Optional) (List) list of notification IDs to call when the run starts. A maximum of 3 destinations can be specified.`,
				},
				resource.Attribute{
					Name:        "on_success",
					Description: `(Optional) (List) list of notification IDs to call when the run completes successfully. A maximum of 3 destinations can be specified.`,
				},
				resource.Attribute{
					Name:        "on_failure",
					Description: `(Optional) (List) list of notification IDs to call when the run fails. A maximum of 3 destinations can be specified. Note that the ` + "`" + `id` + "`" + ` is not to be confused with the name of the alert destination. The ` + "`" + `id` + "`" + ` can be retrieved through the API or the URL of Databricks UI ` + "`" + `https://<workspace host>/sql/destinations/<notification id>?o=<workspace id>` + "`" + ` Example ` + "`" + `` + "`" + `` + "`" + `hcl webhook_notifications { on_failure { id = "fb99f3dc-a0a0-11ed-a8fc-0242ac120002" } } ` + "`" + `` + "`" + `` + "`" + ` ### webhook Configuration Block`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the system notification that is notified when an event defined in ` + "`" + `webhook_notifications` + "`" + ` is triggered. ->`,
				},
				resource.Attribute{
					Name:        "no_alert_for_skipped_runs",
					Description: `(Optional) (Bool) don't send alert for skipped runs.`,
				},
				resource.Attribute{
					Name:        "no_alert_for_canceled_runs",
					Description: `(Optional) (Bool) don't send alert for cancelled runs. ### spark_jar_task Configuration Block`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) (List) Parameters passed to the main method.`,
				},
				resource.Attribute{
					Name:        "main_class_name",
					Description: `(Optional) The full name of the class containing the main method to be executed. This class must be contained in a JAR provided as a library. The code should use ` + "`" + `SparkContext.getOrCreate` + "`" + ` to obtain a Spark context; otherwise, runs of the job will fail. ### spark_submit_task Configuration Block You can invoke Spark submit tasks only on new clusters.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) (List) Command-line parameters passed to spark submit. ### spark_python_task Configuration Block`,
				},
				resource.Attribute{
					Name:        "python_file",
					Description: `(Required) The URI of the Python file to be executed. [databricks_dbfs_file](dbfs_file.md#path), cloud file URIs (e.g. ` + "`" + `s3:/` + "`" + `, ` + "`" + `abfss:/` + "`" + `, ` + "`" + `gs:/` + "`" + `), workspace paths and remote repository are supported. For Python files stored in the Databricks workspace, the path must be absolute and begin with ` + "`" + `/Repos` + "`" + `. For files stored in a remote repository, the path must be relative. This field is required.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Location type of the Python file, can only be ` + "`" + `GIT` + "`" + `. When set to ` + "`" + `GIT` + "`" + `, the Python file will be retrieved from a Git repository defined in ` + "`" + `git_source` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) (List) Command line parameters passed to the Python file. ### notebook_task Configuration Block`,
				},
				resource.Attribute{
					Name:        "notebook_path",
					Description: `(Required) The path of the [databricks_notebook](notebook.md#path) to be run in the Databricks workspace or remote repository. For notebooks stored in the Databricks workspace, the path must be absolute and begin with a slash. For notebooks stored in a remote repository, the path must be relative. This field is required.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Location type of the notebook, can only be ` + "`" + `WORKSPACE` + "`" + ` or ` + "`" + `GIT` + "`" + `. When set to ` + "`" + `WORKSPACE` + "`" + `, the notebook will be retrieved from the local Databricks workspace. When set to ` + "`" + `GIT` + "`" + `, the notebook will be retrieved from a Git repository defined in ` + "`" + `git_source` + "`" + `. If the value is empty, the task will use ` + "`" + `GIT` + "`" + ` if ` + "`" + `git_source` + "`" + ` is defined and ` + "`" + `WORKSPACE` + "`" + ` otherwise.`,
				},
				resource.Attribute{
					Name:        "base_parameters",
					Description: `(Optional) (Map) Base parameters to be used for each run of this job. If the run is initiated by a call to run-now with parameters specified, the two parameters maps will be merged. If the same key is specified in base_parameters and in run-now, the value from run-now will be used. If the notebook takes a parameter that is not specified in the job’s base_parameters or the run-now override parameters, the default value from the notebook will be used. Retrieve these parameters in a notebook using ` + "`" + `dbutils.widgets.get` + "`" + `. ### pipeline_task Configuration Block`,
				},
				resource.Attribute{
					Name:        "pipeline_id",
					Description: `(Required) The pipeline's unique ID. ->`,
				},
				resource.Attribute{
					Name:        "entry_point",
					Description: `(Optional) Python function as entry point for the task`,
				},
				resource.Attribute{
					Name:        "package_name",
					Description: `(Optional) Name of Python package`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) Parameters for the task`,
				},
				resource.Attribute{
					Name:        "named_parameters",
					Description: `(Optional) Named parameters for the task ### dbt_task Configuration Block`,
				},
				resource.Attribute{
					Name:        "commands",
					Description: `(Required) (Array) Series of dbt commands to execute in sequence. Every command must start with "dbt".`,
				},
				resource.Attribute{
					Name:        "project_directory",
					Description: `(Optional) The relative path to the directory in the repository specified in ` + "`" + `git_source` + "`" + ` where dbt should look in for the ` + "`" + `dbt_project.yml` + "`" + ` file. If not specified, defaults to the repository's root directory. Equivalent to passing ` + "`" + `--project-dir` + "`" + ` to a dbt command.`,
				},
				resource.Attribute{
					Name:        "profiles_directory",
					Description: `(Optional) The relative path to the directory in the repository specified by ` + "`" + `git_source` + "`" + ` where dbt should look in for the ` + "`" + `profiles.yml` + "`" + ` file. If not specified, defaults to the repository's root directory. Equivalent to passing ` + "`" + `--profile-dir` + "`" + ` to a dbt command.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Optional) The name of the catalog to use inside Unity Catalog.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `(Optional) The name of the schema dbt should run in. Defaults to ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "warehouse_id",
					Description: `(Optional) The ID of the SQL warehouse that dbt should execute against. You also need to include a ` + "`" + `git_source` + "`" + ` block to configure the repository that contains the dbt project. ### sql_task Configuration Block One of the ` + "`" + `query` + "`" + `, ` + "`" + `dashboard` + "`" + ` or ` + "`" + `alert` + "`" + ` needs to be provided.`,
				},
				resource.Attribute{
					Name:        "warehouse_id",
					Description: `(Required) ID of the (the [databricks_sql_endpoint](sql_endpoint.md)) that will be used to execute the task. Only Serverless & Pro warehouses are supported right now.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) (Map) parameters to be used for each run of this task. The SQL alert task does not support custom parameters.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Optional) block consisting of single string field: ` + "`" + `query_id` + "`" + ` - identifier of the Databricks SQL Query ([databricks_sql_query](sql_query.md)).`,
				},
				resource.Attribute{
					Name:        "dashboard",
					Description: `(Optional) block consisting of single string field: ` + "`" + `dashboard_id` + "`" + ` - identifier of the Databricks SQL Dashboard [databricks_sql_dashboard](sql_dashboard.md).`,
				},
				resource.Attribute{
					Name:        "alert",
					Description: `(Optional) block consisting of single string field: ` + "`" + `alert_id` + "`" + ` - identifier of the Databricks SQL Alert.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `(Optional) block consisting of single string field: ` + "`" + `path` + "`" + ` - a relative path to the file (inside the Git repository) with SQL commands to execute.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the job`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL of the job on the given workspace ## Access Control By default, all users can create and modify jobs unless an administrator [enables jobs access control](https://docs.databricks.com/administration-guide/access-control/jobs-acl.html). With jobs access control, individual permissions determine a user’s abilities.`,
				},
				resource.Attribute{
					Name:        "new_cluster",
					Description: `(Optional) Same set of parameters as for [databricks_cluster](cluster.md) resource.`,
				},
				resource.Attribute{
					Name:        "existing_cluster_id",
					Description: `(Optional) If existing_cluster_id, the ID of an existing [cluster](cluster.md) that will be used for all runs of this job. When running jobs on an existing cluster, you may need to manually restart the cluster if it stops responding. We strongly suggest to use ` + "`" + `new_cluster` + "`" + ` for greater reliability. ` + "`" + `` + "`" + `` + "`" + `hcl data "databricks_current_user" "me" {} data "databricks_spark_version" "latest" {} data "databricks_node_type" "smallest" { local_disk = true } resource "databricks_notebook" "this" { path = "${data.databricks_current_user.me.home}/Terraform" language = "PYTHON" content_base64 = base64encode(<<-EOT # created from ${abspath(path.module)} display(spark.range(10)) EOT ) } resource "databricks_job" "this" { name = "Terraform Demo (${data.databricks_current_user.me.alphanumeric})" new_cluster { num_workers = 1 spark_version = data.databricks_spark_version.latest.id node_type_id = data.databricks_node_type.smallest.id } notebook_task { notebook_path = databricks_notebook.this.path } } output "notebook_url" { value = databricks_notebook.this.url } output "job_url" { value = databricks_job.this.url } ` + "`" + `` + "`" + `` + "`" + ` ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify ` + "`" + `create` + "`" + ` and ` + "`" + `update` + "`" + ` timeouts if you have an ` + "`" + `always_running` + "`" + ` job. Please launch ` + "`" + `TF_LOG=DEBUG terraform apply` + "`" + ` whenever you observe timeout issues. ` + "`" + `` + "`" + `` + "`" + `hcl timeouts { create = "20m" update = "20m" } ` + "`" + `` + "`" + `` + "`" + ` ## Import The resource job can be imported using the id of the job ` + "`" + `` + "`" + `` + "`" + `bash terraform import databricks_job.this <job-id> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_library",
			Category:         "Compute",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"compute",
				"library",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_metastore",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"metastore",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of metastore.`,
				},
				resource.Attribute{
					Name:        "storage_root",
					Description: `Path on cloud storage account, where managed ` + "`" + `databricks_table` + "`" + ` are stored. Change forces creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) Username/groupname/sp application_id of the metastore owner.`,
				},
				resource.Attribute{
					Name:        "delta_sharing_scope",
					Description: `(Optional) Required along with ` + "`" + `delta_sharing_recipient_token_lifetime_in_seconds` + "`" + `. Used to enable delta sharing on the metastore. Valid values: INTERNAL, INTERNAL_AND_EXTERNAL.`,
				},
				resource.Attribute{
					Name:        "delta_sharing_recipient_token_lifetime_in_seconds",
					Description: `(Optional) Required along with ` + "`" + `delta_sharing_scope` + "`" + `. Used to set expiration duration in seconds on recipient data access tokens. Set to 0 for unlimited duration.`,
				},
				resource.Attribute{
					Name:        "delta_sharing_organization_name",
					Description: `(Optional) The organization name of a Delta Sharing entity. This field is used for Databricks to Databricks sharing. Once this is set it cannot be removed and can only be modified to another valid value. To delete this value please taint and recreate the resource.`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional) Destroy metastore regardless of its contents. ## Import This resource can be imported by ID: ` + "`" + `` + "`" + `` + "`" + `bash terraform import databricks_metastore.this <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_metastore_assignment",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"metastore",
				"assignment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metastore_id",
					Description: `Unique identifier of the parent Metastore`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `id of the workspace for the assignment`,
				},
				resource.Attribute{
					Name:        "default_catalog_name",
					Description: `(Optional) Default catalog used for this assignment, default to ` + "`" + `hive_metastore` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_metastore_data_access",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"metastore",
				"data",
				"access",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of Data Access Configuration, which must be unique within the [databricks_metastore](metastore.md). Change forces creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "metastore_id",
					Description: `Unique identifier of the parent Metastore ` + "`" + `aws_iam_role` + "`" + ` optional configuration block for credential details for AWS:`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `The Amazon Resource Name (ARN) of the AWS IAM role for S3 data access, of the form ` + "`" + `arn:aws:iam::1234567890:role/MyRole-AJJHDSKSDF` + "`" + ` ` + "`" + `azure_service_principal` + "`" + ` optional configuration block for credential details for Azure:`,
				},
				resource.Attribute{
					Name:        "directory_id",
					Description: `The directory ID corresponding to the Azure Active Directory (AAD) tenant of the application`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `The application ID of the application registration within the referenced AAD tenant`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `The client secret generated for the above app ID in AAD.`,
				},
				resource.Attribute{
					Name:        "access_connector_id",
					Description: `The Resource ID of the Azure Databricks Access Connector resource, of the form ` + "`" + `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-name/providers/Microsoft.Databricks/accessConnectors/connector-name` + "`" + ` ` + "`" + `databricks_gcp_service_account` + "`" + ` optional configuration block for creating a Databricks-managed GCP Service Account:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mlflow_experiment",
			Category:         "MLflow",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"mlflow",
				"experiment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of MLflow experiment. It must be an absolute path within the Databricks workspace, e.g. ` + "`" + `/Users/<some-username>/my-experiment` + "`" + `. For more information about changes to experiment naming conventions, see [mlflow docs](https://docs.databricks.com/applications/mlflow/experiments.html#experiment-migration).`,
				},
				resource.Attribute{
					Name:        "artifact_location",
					Description: `Path to dbfs:/ or s3:// artifact location of the MLflow experiment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the MLflow experiment. ## Access Control`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mlflow_model",
			Category:         "MLflow",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"mlflow",
				"model",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of MLflow model. Change of name triggers new resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the MLflow model.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags for the MLflow model. ## Import The model resource can be imported using the name ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_mlflow_model.this <name> ` + "`" + `` + "`" + `` + "`" + ` ## Access Control`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mlflow_webhook",
			Category:         "MLflow",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"mlflow",
				"webhook",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "model_name",
					Description: `(Optional) Name of MLflow model for which webhook will be created. If the model name is not specified, a registry-wide webhook is created that listens for the specified events across all versions of all registered models.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Optional description of the MLflow webhook.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Optional status of webhook. Possible values are ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `TEST_MODE` + "`" + `, ` + "`" + `DISABLED` + "`" + `. Default is ` + "`" + `ACTIVE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "events",
					Description: `(Required) The list of events that will trigger execution of Databricks job or POSTing to an URL, for example, ` + "`" + `MODEL_VERSION_CREATED` + "`" + `, ` + "`" + `MODEL_VERSION_TRANSITIONED_STAGE` + "`" + `, ` + "`" + `TRANSITION_REQUEST_CREATED` + "`" + `, etc. Refer to the [Webhooks API documentation](https://docs.databricks.com/dev-tools/api/latest/mlflow.html#operation/create-registry-webhook) for a full list of supported events. Configuration must include one of ` + "`" + `http_url_spec` + "`" + ` or ` + "`" + `job_spec` + "`" + ` blocks, but not both. ### job_spec`,
				},
				resource.Attribute{
					Name:        "access_token",
					Description: `(Required) The personal access token used to authorize webhook's job runs.`,
				},
				resource.Attribute{
					Name:        "job_id",
					Description: `(Required) ID of the Databricks job that the webhook runs.`,
				},
				resource.Attribute{
					Name:        "workspace_url",
					Description: `(Optional) URL of the workspace containing the job that this webhook runs. If not specified, the job’s workspace URL is assumed to be the same as the workspace where the webhook is created. ### http_url_spec`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) External HTTPS URL called on event trigger (by using a POST request). Structure of payload depends on the event type, refer to [documentation](https://docs.databricks.com/applications/mlflow/model-registry-webhooks.html) for more details.`,
				},
				resource.Attribute{
					Name:        "authorization",
					Description: `(Optional) Value of the authorization header that should be sent in the request sent by the wehbook. It should be of the form ` + "`" + `<auth type> <credentials>` + "`" + `, e.g. ` + "`" + `Bearer <access_token>` + "`" + `. If set to an empty string, no authorization header will be included in the request.`,
				},
				resource.Attribute{
					Name:        "enable_ssl_verification",
					Description: `(Optional) Enable/disable SSL certificate validation. Default is ` + "`" + `true` + "`" + `. For self-signed certificates, this field must be ` + "`" + `false` + "`" + ` AND the destination server must disable certificate validation as well. For security purposes, it is encouraged to perform secret validation with the HMAC-encoded portion of the payload and acknowledge the risk associated with disabling hostname validation whereby it becomes more likely that requests can be maliciously routed to an unintended host.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Shared secret required for HMAC encoding payload. The HMAC-encoded payload will be sent in the header as ` + "`" + `X-Databricks-Signature: encoded_payload` + "`" + `. ## Import ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_model_serving",
			Category:         "Serving",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"serving",
				"model",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the model serving endpoint. This field is required and must be unique across a workspace. An endpoint name can consist of alphanumeric characters, dashes, and underscores. NOTE: Changing this name will delete the existing endpoint and create a new endpoint with the update name.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Required) The model serving endpoint configuration. ### config Configuration Block`,
				},
				resource.Attribute{
					Name:        "served_models",
					Description: `(Required) Each block represents a served model for the endpoint to serve. A model serving endpoint can have up to 10 served models.`,
				},
				resource.Attribute{
					Name:        "traffic_config",
					Description: `A single block represents the traffic split configuration amongst the served models. ### served_models Configuration Block`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of a served model. It must be unique across an endpoint. If not specified, this field will default to ` + "`" + `modelname-modelversion` + "`" + `. A served model name can consist of alphanumeric characters, dashes, and underscores.`,
				},
				resource.Attribute{
					Name:        "model_name",
					Description: `(Required) The name of the model in Databricks Model Registry to be served.`,
				},
				resource.Attribute{
					Name:        "model_version",
					Description: `(Required) The version of the model in Databricks Model Registry to be served.`,
				},
				resource.Attribute{
					Name:        "workload_size",
					Description: `(Required) The workload size of the served model. The workload size corresponds to a range of provisioned concurrency that the compute will autoscale between. A single unit of provisioned concurrency can process one request at a time. Valid workload sizes are "Small" (4 - 4 provisioned concurrency), "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64 provisioned concurrency).`,
				},
				resource.Attribute{
					Name:        "scale_to_zero_enabled",
					Description: `Whether the compute resources for the served model should scale down to zero. If scale-to-zero is enabled, the lower bound of the provisioned concurrency for each workload size will be 0. The default value is ` + "`" + `true` + "`" + `. ### traffic_config Configuration Block`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `(Required) Each block represents a route that defines traffic to each served model. Each ` + "`" + `served_models` + "`" + ` block needs to have a corresponding ` + "`" + `routes` + "`" + ` block ### routes Configuration Block`,
				},
				resource.Attribute{
					Name:        "served_model_name",
					Description: `(Required) The name of the served model this route configures traffic for. This needs to match the name of a ` + "`" + `served_models` + "`" + ` block`,
				},
				resource.Attribute{
					Name:        "traffic_percentage",
					Description: `(Required) The percentage of endpoint traffic to send to this route. It must be an integer between 0 and 100 inclusive. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify ` + "`" + `create` + "`" + ` and ` + "`" + `update` + "`" + ` timeouts. The default right now is 45 minutes for both operations. ` + "`" + `` + "`" + `` + "`" + `hcl timeouts { create = "30m" } ` + "`" + `` + "`" + `` + "`" + ` ## Import The model serving resource can be imported using the name of the endpoint. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_model_serving.this <model-serving-endpoint-name> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mount",
			Category:         "Storage",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"storage",
				"mount",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `mount name`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(String) HDFS-compatible url ## Import ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_credentials",
			Category:         "Deployment",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"mws",
				"credentials",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/)`,
				},
				resource.Attribute{
					Name:        "credentials_name",
					Description: `(Required) name of credentials to register`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Required) ARN of cross-account role ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the mws credentials.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(Integer) time of credentials registration`,
				},
				resource.Attribute{
					Name:        "credentials_id",
					Description: `(String) identifier of credentials ## Import ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the mws credentials.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(Integer) time of credentials registration`,
				},
				resource.Attribute{
					Name:        "credentials_id",
					Description: `(String) identifier of credentials ## Import ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_customer_managed_keys",
			Category:         "Deployment",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"mws",
				"customer",
				"managed",
				"keys",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aws_key_info",
					Description: `This field is a block and is documented below.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/)`,
				},
				resource.Attribute{
					Name:        "MANAGED_SERVICES",
					Description: `for encryption of the workspace objects (notebooks, secrets) that are stored in the control plane`,
				},
				resource.Attribute{
					Name:        "STORAGE",
					Description: `for encryption of the DBFS Storage & Cluster EBS Volumes ### aws_key_info Configuration Block`,
				},
				resource.Attribute{
					Name:        "key_arn",
					Description: `The AWS KMS key's Amazon Resource Name (ARN).`,
				},
				resource.Attribute{
					Name:        "key_alias",
					Description: `The AWS KMS key alias.`,
				},
				resource.Attribute{
					Name:        "key_region",
					Description: `(Optional) (Computed) The AWS region in which KMS key is deployed to. This is not required. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the mws customer managed keys.`,
				},
				resource.Attribute{
					Name:        "customer_managed_key_id",
					Description: `(String) ID of the encryption key configuration object.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(Integer) Time in epoch milliseconds when the customer key was created. ## Import ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the mws customer managed keys.`,
				},
				resource.Attribute{
					Name:        "customer_managed_key_id",
					Description: `(String) ID of the encryption key configuration object.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(Integer) Time in epoch milliseconds when the customer key was created. ## Import ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_log_delivery",
			Category:         "Log Delivery",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"log",
				"delivery",
				"mws",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_networks",
			Category:         "Deployment",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"mws",
				"networks",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/)`,
				},
				resource.Attribute{
					Name:        "network_name",
					Description: `name under which this network is registered`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(AWS only) [aws_vpc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc) id`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(AWS only) ids of [aws_subnet](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet)`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(AWS only) ids of [aws_security_group](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group)`,
				},
				resource.Attribute{
					Name:        "vpc_endpoints",
					Description: `(Optional) mapping of [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md) for PrivateLink or Private Service Connect connections`,
				},
				resource.Attribute{
					Name:        "gcp_network_info",
					Description: `(GCP only) a block consists of Google Cloud specific information for this network, for example the VPC ID, subnet ID, and secondary IP ranges. It has the following fields:`,
				},
				resource.Attribute{
					Name:        "network_project_id",
					Description: `The Google Cloud project ID of the VPC network.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC associated with this network. VPC IDs can be used in multiple network configurations.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet associated with this network.`,
				},
				resource.Attribute{
					Name:        "subnet_region",
					Description: `The Google Cloud region of the workspace data plane. For example, ` + "`" + `us-east4` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pod_ip_range_name",
					Description: `The name of the secondary IP range for pods. A Databricks-managed GKE cluster uses this IP range for its pods. This secondary IP range can only be used by one workspace.`,
				},
				resource.Attribute{
					Name:        "service_ip_range_name",
					Description: `The name of the secondary IP range for services. A Databricks-managed GKE cluster uses this IP range for its services. This secondary IP range can only be used by one workspace. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the mws networks.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(String) id of network to be used for [databricks_mws_workspaces](mws_workspaces.md) resource.`,
				},
				resource.Attribute{
					Name:        "vpc_status",
					Description: `(String) VPC attachment status`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `(Integer) id of associated workspace ## Import ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the mws networks.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(String) id of network to be used for [databricks_mws_workspaces](mws_workspaces.md) resource.`,
				},
				resource.Attribute{
					Name:        "vpc_status",
					Description: `(String) VPC attachment status`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `(Integer) id of associated workspace ## Import ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_permission_assignment",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"mws",
				"permission",
				"assignment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "workspace_id",
					Description: `Databricks workspace ID.`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `Databricks ID of the user, service principal, or group. The principal ID can be retrieved using the SCIM API, or using [databricks_user](../data-sources/user.md), [databricks_service_principal](../data-sources/service_principal.md) or [databricks_group](../data-sources/group.md) data sources.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `The list of workspace permissions to assign to the principal:`,
				},
				resource.Attribute{
					Name:        "USER",
					Description: `Can access the workspace with basic privileges.`,
				},
				resource.Attribute{
					Name:        "ADMIN",
					Description: `Can access the workspace and has workspace admin privileges to manage users and groups, workspace configurations, and more. ## Import The resource ` + "`" + `databricks_mws_permission_assignment` + "`" + ` can be imported using the workspace id and principal id ` + "`" + `` + "`" + `` + "`" + `bash terraform import databricks_mws_permission_assignment.this "workspace_id|principal_id" ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_private_access_settings",
			Category:         "Deployment",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"mws",
				"private",
				"access",
				"settings",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `Account Id that could be found in the Accounts Console for [AWS](https://accounts.cloud.databricks.com/) or [GCP](https://accounts.gcp.databricks.com/)`,
				},
				resource.Attribute{
					Name:        "private_access_settings_name",
					Description: `Name of Private Access Settings in Databricks Account`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of AWS VPC or the Google Cloud VPC network`,
				},
				resource.Attribute{
					Name:        "private_access_level",
					Description: `(Optional) The private access level controls which VPC endpoints can connect to the UI or API of any workspace that attaches this private access settings object. ` + "`" + `ACCOUNT` + "`" + ` level access _(default)_ lets only [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md) that are registered in your Databricks account connect to your [databricks_mws_workspaces](mws_workspaces.md). ` + "`" + `ENDPOINT` + "`" + ` level access lets only specified [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md) connect to your workspace. Please see the ` + "`" + `allowed_vpc_endpoint_ids` + "`" + ` documentation for more details.`,
				},
				resource.Attribute{
					Name:        "allowed_vpc_endpoint_ids",
					Description: `(Optional) An array of [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md#vpc_endpoint_id) ` + "`" + `vpc_endpoint_id` + "`" + ` (not ` + "`" + `id` + "`" + `). Only used when ` + "`" + `private_access_level` + "`" + ` is set to ` + "`" + `ENDPOINT` + "`" + `. This is an allow list of [databricks_mws_vpc_endpoint](mws_vpc_endpoint.md) that in your account that can connect to your [databricks_mws_workspaces](mws_workspaces.md) over AWS PrivateLink. If hybrid access to your workspace is enabled by setting ` + "`" + `public_access_enabled` + "`" + ` to true, then this control only works for PrivateLink connections. To control how your workspace is accessed via public internet, see the article for [databricks_ip_access_list](ip_access_list.md). ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "private_access_settings_id",
					Description: `Canonical unique identifier of Private Access Settings in Databricks Account`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(AWS only) Status of Private Access Settings ## Import ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "private_access_settings_id",
					Description: `Canonical unique identifier of Private Access Settings in Databricks Account`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(AWS only) Status of Private Access Settings ## Import ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_storage_configurations",
			Category:         "Deployment",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"mws",
				"storage",
				"configurations",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket_name",
					Description: `name of AWS S3 bucket`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/)`,
				},
				resource.Attribute{
					Name:        "storage_configuration_name",
					Description: `name under which this storage configuration is stored ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the mws storage configurations.`,
				},
				resource.Attribute{
					Name:        "storage_configuration_id",
					Description: `(String) id of storage config to be used for ` + "`" + `databricks_mws_workspace` + "`" + ` resource. ## Import ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the mws storage configurations.`,
				},
				resource.Attribute{
					Name:        "storage_configuration_id",
					Description: `(String) id of storage config to be used for ` + "`" + `databricks_mws_workspace` + "`" + ` resource. ## Import ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_vpc_endpoint",
			Category:         "Deployment",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"mws",
				"vpc",
				"endpoint",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `Account Id that could be found in the Accounts Console for [AWS](https://accounts.cloud.databricks.com/) or [GCP](https://accounts.gcp.databricks.com/)`,
				},
				resource.Attribute{
					Name:        "aws_vpc_endpoint_id",
					Description: `(AWS only) ID of configured [aws_vpc_endpoint](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint)`,
				},
				resource.Attribute{
					Name:        "vpc_endpoint_name",
					Description: `Name of VPC Endpoint in Databricks Account`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(AWS only) Region of AWS VPC`,
				},
				resource.Attribute{
					Name:        "gcp_vpc_endpoint_info",
					Description: `(GCP only) a block consists of Google Cloud specific information for this PSC endpoint. It has the following fields:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The Google Cloud project ID of the VPC network where the PSC connection resides.`,
				},
				resource.Attribute{
					Name:        "psc_endpoint_name",
					Description: `The name of the PSC endpoint in the Google Cloud project.`,
				},
				resource.Attribute{
					Name:        "endpoint_region",
					Description: `Region of the PSC endpoint. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpc_endpoint_id",
					Description: `Canonical unique identifier of VPC Endpoint in Databricks Account`,
				},
				resource.Attribute{
					Name:        "aws_endpoint_service_id",
					Description: `(AWS Only) The ID of the Databricks endpoint service that this VPC endpoint is connected to. Please find the list of endpoint service IDs for each supported region in the [Databricks PrivateLink documentation](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html)`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(AWS Only) State of VPC Endpoint`,
				},
				resource.Attribute{
					Name:        "psc_connection_id",
					Description: `The unique ID of this PSC connection.`,
				},
				resource.Attribute{
					Name:        "service_attachment_id",
					Description: `The service attachment this PSC connection connects to. ## Import ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_endpoint_id",
					Description: `Canonical unique identifier of VPC Endpoint in Databricks Account`,
				},
				resource.Attribute{
					Name:        "aws_endpoint_service_id",
					Description: `(AWS Only) The ID of the Databricks endpoint service that this VPC endpoint is connected to. Please find the list of endpoint service IDs for each supported region in the [Databricks PrivateLink documentation](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html)`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(AWS Only) State of VPC Endpoint`,
				},
				resource.Attribute{
					Name:        "psc_connection_id",
					Description: `The unique ID of this PSC connection.`,
				},
				resource.Attribute{
					Name:        "service_attachment_id",
					Description: `The service attachment this PSC connection connects to. ## Import ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_workspaces",
			Category:         "Deployment",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"mws",
				"workspaces",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/).`,
				},
				resource.Attribute{
					Name:        "deployment_name",
					Description: `(Optional) part of URL as in ` + "`" + `https://<prefix>-<deployment-name>.cloud.databricks.com` + "`" + `. Deployment name cannot be used until a deployment name prefix is defined. Please contact your Databricks representative. Once a new deployment prefix is added/updated, it only will affect the new workspaces created.`,
				},
				resource.Attribute{
					Name:        "workspace_name",
					Description: `name of the workspace, will appear on UI.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) ` + "`" + `network_id` + "`" + ` from [networks](mws_networks.md).`,
				},
				resource.Attribute{
					Name:        "aws_region",
					Description: `(AWS only) region of VPC.`,
				},
				resource.Attribute{
					Name:        "storage_configuration_id",
					Description: `(AWS only)` + "`" + `storage_configuration_id` + "`" + ` from [storage configuration](mws_storage_configurations.md).`,
				},
				resource.Attribute{
					Name:        "managed_services_customer_managed_key_id",
					Description: `(Optional, AWS only) ` + "`" + `customer_managed_key_id` + "`" + ` from [customer managed keys](mws_customer_managed_keys.md) with ` + "`" + `use_cases` + "`" + ` set to ` + "`" + `MANAGED_SERVICES` + "`" + `. This is used to encrypt the workspace's notebook and secret data in the control plane.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(GCP only) region of the subnet.`,
				},
				resource.Attribute{
					Name:        "cloud_resource_container",
					Description: `(GCP only) A block that specifies GCP workspace configurations, consisting of following blocks:`,
				},
				resource.Attribute{
					Name:        "gcp",
					Description: `A block that consists of the following field:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The Google Cloud project ID, which the workspace uses to instantiate cloud resources for your workspace.`,
				},
				resource.Attribute{
					Name:        "gke_config",
					Description: `(GCP only) A block that specifies GKE configuration for the Databricks workspace:`,
				},
				resource.Attribute{
					Name:        "private_access_settings_id",
					Description: `(Optional) Canonical unique identifier of [databricks_mws_private_access_settings](mws_private_access_settings.md) in Databricks Account. ## token block You can specify a ` + "`" + `token` + "`" + ` block in the body of the workspace resource, so that Terraform manages the refresh of the PAT token for the deployment user. The other option is to create [databricks_obo_token](obo_token.md), though it requires Premium or Enterprise plan enabled as well as more complex setup. Token block exposes ` + "`" + `token_value` + "`" + `, that holds sensitive PAT token and optionally it can accept two arguments: ->`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment, that will appear in "User Settings / Access Tokens" page on Workspace UI. By default it's "Terraform PAT".`,
				},
				resource.Attribute{
					Name:        "lifetime_seconds",
					Description: `(Optional) Token expiry lifetime. By default its 2592000 (30 days). On AWS, the following arguments could be modified after the workspace is running:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional, AWS only) ` + "`" + `network_id` + "`" + ` from [networks](mws_networks.md). Modifying [networks on running workspaces](mws_networks.md#modifying-networks-on-running-workspaces) would require three separate ` + "`" + `terraform apply` + "`" + ` steps.`,
				},
				resource.Attribute{
					Name:        "credentials_id",
					Description: `(AWS only) ` + "`" + `credentials_id` + "`" + ` from [credentials](mws_credentials.md)`,
				},
				resource.Attribute{
					Name:        "storage_customer_managed_key_id",
					Description: `(Optional, AWS only) ` + "`" + `customer_managed_key_id` + "`" + ` from [customer managed keys](mws_customer_managed_keys.md) with ` + "`" + `use_cases` + "`" + ` set to ` + "`" + `STORAGE` + "`" + `. This is used to encrypt the DBFS Storage & Cluster EBS Volumes.`,
				},
				resource.Attribute{
					Name:        "private_access_settings_id",
					Description: `(Optional, AWS only) Canonical unique identifier of [databricks_mws_private_access_settings](mws_private_access_settings.md) in Databricks Account ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(String) Canonical unique identifier for the workspace, of the format ` + "`" + `<account-id>/<workspace-id>` + "`" + ``,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `(String) workspace id`,
				},
				resource.Attribute{
					Name:        "workspace_status_message",
					Description: `(String) updates on workspace status`,
				},
				resource.Attribute{
					Name:        "workspace_status",
					Description: `(String) workspace status`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(Integer) time when workspace was created`,
				},
				resource.Attribute{
					Name:        "workspace_url",
					Description: `(String) URL of the workspace ## Import ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(String) Canonical unique identifier for the workspace, of the format ` + "`" + `<account-id>/<workspace-id>` + "`" + ``,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `(String) workspace id`,
				},
				resource.Attribute{
					Name:        "workspace_status_message",
					Description: `(String) updates on workspace status`,
				},
				resource.Attribute{
					Name:        "workspace_status",
					Description: `(String) workspace status`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(Integer) time when workspace was created`,
				},
				resource.Attribute{
					Name:        "workspace_url",
					Description: `(String) URL of the workspace ## Import ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_notebook",
			Category:         "Workspace",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"workspace",
				"notebook",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The absolute path of the notebook or directory, beginning with "/", e.g. "/Demo".`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Path to notebook in source code format on local filesystem. Conflicts with ` + "`" + `content_base64` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `The base64-encoded notebook source code. Conflicts with ` + "`" + `source` + "`" + `. Use of ` + "`" + `content_base64` + "`" + ` is discouraged, as it's increasing memory footprint of Terraform state and should only be used in exceptional circumstances, like creating a notebook with configuration properties for a data pipeline.`,
				},
				resource.Attribute{
					Name:        "language",
					Description: `(required with ` + "`" + `content_base64` + "`" + `) One of ` + "`" + `SCALA` + "`" + `, ` + "`" + `PYTHON` + "`" + `, ` + "`" + `SQL` + "`" + `, ` + "`" + `R` + "`" + `. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Path of notebook on workspace`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Routable URL of the notebook`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `Unique identifier for a NOTEBOOK ## Access Control`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Path of notebook on workspace`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Routable URL of the notebook`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `Unique identifier for a NOTEBOOK ## Access Control`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_obo_token",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"obo",
				"token",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_id",
					Description: `Application ID of [databricks_service_principal](service_principal.md#application_id) to create a PAT token for.`,
				},
				resource.Attribute{
					Name:        "lifetime_seconds",
					Description: `(Integer, Optional) The number of seconds before the token expires. Token resource is re-created when it expires. If no lifetime is specified, the token remains valid indefinitely.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(String, Optional) Comment that describes the purpose of the token. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the token.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the token.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_permission_assignment",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"permission",
				"assignment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "principal_id",
					Description: `Databricks ID of the user, service principal, or group. The principal ID can be retrieved using the account-level SCIM API, or using [databricks_user](../data-sources/user.md), [databricks_service_principal](../data-sources/service_principal.md) or [databricks_group](../data-sources/group.md) data sources with account API (and has to be an account admin). A more sensible approach is to retrieve the list of ` + "`" + `principal_id` + "`" + ` as outputs from another Terraform stack.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `The list of workspace permissions to assign to the principal:`,
				},
				resource.Attribute{
					Name:        "USER",
					Description: `Can access the workspace with basic privileges.`,
				},
				resource.Attribute{
					Name:        "ADMIN",
					Description: `Can access the workspace and has workspace admin privileges to manage users and groups, workspace configurations, and more. ## Import The resource ` + "`" + `databricks_permission_assignment` + "`" + ` can be imported using the principal id ` + "`" + `` + "`" + `` + "`" + `bash terraform import databricks_permission_assignment.this principal_id ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_permissions",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"permissions",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_pipeline",
			Category:         "Compute",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"compute",
				"pipeline",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `A user-friendly name for this pipeline. The name can be used to identify pipeline jobs in the UI.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `A location on DBFS or cloud storage where output data and metadata required for pipeline execution are stored. By default, tables are stored in a subdirectory of this location.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `An optional list of values to apply to the entire pipeline. Elements must be formatted as key:value pairs.`,
				},
				resource.Attribute{
					Name:        "continuous",
					Description: `A flag indicating whether to run the pipeline continuously. The default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "development",
					Description: `A flag indicating whether to run the pipeline in development mode. The default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "photon",
					Description: `A flag indicating whether to use Photon engine. The default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `The name of a database for persisting pipeline output data. Configuring the target setting allows you to view and query the pipeline output data from the Databricks UI.`,
				},
				resource.Attribute{
					Name:        "edition",
					Description: `optional name of the [product edition](https://docs.databricks.com/data-engineering/delta-live-tables/delta-live-tables-concepts.html#editions). Supported values are: ` + "`" + `core` + "`" + `, ` + "`" + `pro` + "`" + `, ` + "`" + `advanced` + "`" + ` (default).`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `optional name of the release channel for Spark version used by DLT pipeline. Supported values are: ` + "`" + `current` + "`" + ` (default) and ` + "`" + `preview` + "`" + `. ### notification block DLT allows to specify one or more notification blocks to get notifications about pipeline's execution. This block consists of following attributes:`,
				},
				resource.Attribute{
					Name:        "on-update-success",
					Description: `a pipeline update completes successfully.`,
				},
				resource.Attribute{
					Name:        "on-update-failure",
					Description: `a pipeline update fails with a retryable error.`,
				},
				resource.Attribute{
					Name:        "on-update-fatal-failure",
					Description: `a pipeline update fails with a non-retryable (fatal) error.`,
				},
				resource.Attribute{
					Name:        "on-flow-failure",
					Description: `a single data flow fails. ## Import The resource job can be imported using the id of the pipeline ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_pipeline.this <pipeline-id> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_provider",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"provider",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of provider. Change forces creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Description about the provider.`,
				},
				resource.Attribute{
					Name:        "authentication_type",
					Description: `(Optional) The delta sharing authentication type. Valid values are ` + "`" + `TOKEN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "recipient_profile_str",
					Description: `(Optional) This is the json file that is created from a recipient url. ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_recipient",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"recipient",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of recipient. Change forces creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Description about the recipient.`,
				},
				resource.Attribute{
					Name:        "sharing_code",
					Description: `(Optional) The one-time sharing code provided by the data recipient.`,
				},
				resource.Attribute{
					Name:        "authentication_type",
					Description: `(Optional) The delta sharing authentication type. Valid values are ` + "`" + `TOKEN` + "`" + ` and ` + "`" + `DATABRICKS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "data_recipient_global_metastore_id",
					Description: `Required when authentication_type is DATABRICKS.`,
				},
				resource.Attribute{
					Name:        "ip_access_list",
					Description: `(Optional) The one-time sharing code provided by the data recipient. ### Ip Access List Argument Only one ` + "`" + `ip_access_list` + "`" + ` blocks is allowed in a recipient. It conflicts with authentication type DATABRICKS. ` + "`" + `` + "`" + `` + "`" + `hcl ip_access_list { allowed_ip_addresses = ["0.0.0.0/0"] } ` + "`" + `` + "`" + `` + "`" + ` Arguments for the ` + "`" + `ip_access_list` + "`" + ` block are: Exactly one of the below arguments is required:`,
				},
				resource.Attribute{
					Name:        "allowed_ip_addresses",
					Description: `Allowed IP Addresses in CIDR notation. Limit of 100. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "tokens",
					Description: `List of Recipient Tokens. ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "tokens",
					Description: `List of Recipient Tokens. ## Related Resources The following resources are often used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_repo",
			Category:         "Workspace",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"workspace",
				"repo",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL of the Git Repository to clone from. If the value changes, repo is re-created.`,
				},
				resource.Attribute{
					Name:        "git_provider",
					Description: `(Optional, if it's possible to detect Git provider by host name) case insensitive name of the Git provider. Following values are supported right now (could be a subject for a change, consult [Repos API documentation](https://docs.databricks.com/dev-tools/api/latest/repos.html)): ` + "`" + `gitHub` + "`" + `, ` + "`" + `gitHubEnterprise` + "`" + `, ` + "`" + `bitbucketCloud` + "`" + `, ` + "`" + `bitbucketServer` + "`" + `, ` + "`" + `azureDevOpsServices` + "`" + `, ` + "`" + `gitLab` + "`" + `, ` + "`" + `gitLabEnterpriseEdition` + "`" + `, ` + "`" + `awsCodeCommit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) path to put the checked out Repo. If not specified, then repo will be created in the user's repo directory (` + "`" + `/Repos/<username>/...` + "`" + `). If the value changes, repo is re-created.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `(Optional) name of the branch for initial checkout. If not specified, the default branch of the repository will be used. Conflicts with ` + "`" + `tag` + "`" + `. If ` + "`" + `branch` + "`" + ` is removed, and ` + "`" + `tag` + "`" + ` isn't specified, then the repository will stay at the previously checked out state.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) name of the tag for initial checkout. Conflicts with ` + "`" + `branch` + "`" + `. ## sparse_checkout Optional ` + "`" + `sparse_checkout` + "`" + ` configuration block contains attributes related to [sparse checkout feature](https://docs.databricks.com/repos/git-operations-with-repos.html#configure-sparse-checkout-mode) in Databricks Repos. It supports following attributes:`,
				},
				resource.Attribute{
					Name:        "patterns",
					Description: `array of paths (directories) that will be used for sparse checkout. List of patterns could be updated in-place. Addition or removal of the ` + "`" + `sparse_checkout` + "`" + ` configuration block will lead to recreation of the repo. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Repo identifier`,
				},
				resource.Attribute{
					Name:        "commit_hash",
					Description: `Hash of the HEAD commit at time of the last executed operation. It won't change if you manually perform pull operation via UI or API ## Access Control`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Repo identifier`,
				},
				resource.Attribute{
					Name:        "commit_hash",
					Description: `Hash of the HEAD commit at time of the last executed operation. It won't change if you manually perform pull operation via UI or API ## Access Control`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_schema",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"schema",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of Schema relative to parent catalog. Change forces creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Name of parent catalog`,
				},
				resource.Attribute{
					Name:        "storage_root",
					Description: `(Optional) Managed location of the schema. Location in cloud storage where data for managed tables will be stored. If not specified, the location will default to the metastore root location. Change forces creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) Username/groupname/sp application_id of the schema owner.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) User-supplied free-form text.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) Extensible Schema properties.`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional) Delete schema regardless of its contents. ## Import This resource can be imported by name: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_schema.this <name> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_secret",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"secret",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "string_value",
					Description: `(Required) (String) super secret sensitive value.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Required) (String) name of databricks secret scope. Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) (String) key within secret scope. Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the secret.`,
				},
				resource.Attribute{
					Name:        "last_updated_timestamp",
					Description: `(Integer) time secret was updated`,
				},
				resource.Attribute{
					Name:        "config_reference",
					Description: `(String) value to use as a secret reference in [Spark configuration and environment variables](https://docs.databricks.com/security/secrets/secrets.html#use-a-secret-in-a-spark-configuration-property-or-environment-variable): ` + "`" + `{{secrets/scope/key}}` + "`" + `. ## Import The resource secret can be imported using ` + "`" + `scopeName|||secretKey` + "`" + ` combination.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the secret.`,
				},
				resource.Attribute{
					Name:        "last_updated_timestamp",
					Description: `(Integer) time secret was updated`,
				},
				resource.Attribute{
					Name:        "config_reference",
					Description: `(String) value to use as a secret reference in [Spark configuration and environment variables](https://docs.databricks.com/security/secrets/secrets.html#use-a-secret-in-a-spark-configuration-property-or-environment-variable): ` + "`" + `{{secrets/scope/key}}` + "`" + `. ## Import The resource secret can be imported using ` + "`" + `scopeName|||secretKey` + "`" + ` combination.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_secret_acl",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"secret",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scope",
					Description: `(Required) name of the scope`,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `(Required) name of the principals. It can be ` + "`" + `users` + "`" + ` for all users or name or ` + "`" + `display_name` + "`" + ` of [databricks_group](group.md)`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Required) ` + "`" + `READ` + "`" + `, ` + "`" + `WRITE` + "`" + ` or ` + "`" + `MANAGE` + "`" + `. ## Import The resource secret acl can be imported using ` + "`" + `scopeName|||principalName` + "`" + ` combination. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_secret_acl.object ` + "`" + `scopeName|||principalName` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_secret_scope",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"secret",
				"scope",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Scope name requested by the user. Must be unique within a workspace. Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters.`,
				},
				resource.Attribute{
					Name:        "initial_manage_principal",
					Description: `(Optional) The principal with the only possible value ` + "`" + `users` + "`" + ` that is initially granted ` + "`" + `MANAGE` + "`" + ` permission to the created scope. If it's omitted, then the [databricks_secret_acl](secret_acl.md) with ` + "`" + `MANAGE` + "`" + ` permission applied to the scope is assigned to the API request issuer's user identity (see [documentation](https://docs.databricks.com/dev-tools/api/latest/secrets.html#create-secret-scope)). This part of the state cannot be imported. ## keyvault_metadata On Azure it's possible to create and manage secrets in Azure Key Vault and have use Azure Databricks secret redaction & access control functionality for reading them. There has to be a single Key Vault per single secret scope. To define AKV access policies, you must use [azurerm_key_vault_access_policy](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/key_vault_access_policy) instead of [access_policy](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/key_vault#access_policy) blocks on ` + "`" + `azurerm_key_vault` + "`" + `, otherwise Terraform will remove access policies needed to access the Key Vault and the secret scope won't be in a usable state anymore. ` + "`" + `` + "`" + `` + "`" + `hcl data "azurerm_client_config" "current" { } resource "azurerm_key_vault" "this" { name = "${var.prefix}-kv" location = azurerm_resource_group.example.location resource_group_name = azurerm_resource_group.example.name tenant_id = data.azurerm_client_config.current.tenant_id soft_delete_enabled = false purge_protection_enabled = false sku_name = "standard" tags = var.tags } resource "azurerm_key_vault_access_policy" "this" { key_vault_id = azurerm_key_vault.this.id tenant_id = data.azurerm_client_config.current.tenant_id object_id = data.azurerm_client_config.current.object_id secret_permissions = ["delete", "get", "list", "set"] } resource "databricks_secret_scope" "kv" { name = "keyvault-managed" keyvault_metadata { resource_id = azurerm_key_vault.this.id dns_name = azurerm_key_vault.this.vault_uri } } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id for the secret scope object.`,
				},
				resource.Attribute{
					Name:        "backend_type",
					Description: `Either ` + "`" + `DATABRICKS` + "`" + ` or ` + "`" + `AZURE_KEYVAULT` + "`" + ` ## Import The secret resource scope can be imported using the scope name. ` + "`" + `initial_manage_principal` + "`" + ` state won't be imported, because the underlying API doesn't include it in the response. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_secret_scope.object <scopeName> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id for the secret scope object.`,
				},
				resource.Attribute{
					Name:        "backend_type",
					Description: `Either ` + "`" + `DATABRICKS` + "`" + ` or ` + "`" + `AZURE_KEYVAULT` + "`" + ` ## Import The secret resource scope can be imported using the scope name. ` + "`" + `initial_manage_principal` + "`" + ` state won't be imported, because the underlying API doesn't include it in the response. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_secret_scope.object <scopeName> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_service_principal",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"service",
				"principal",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_id",
					Description: `This is the Azure Application ID of the given Azure service principal and will be their form of access and identity. On other clouds than Azure this value is auto-generated.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) This is an alias for the service principal and can be the full name of the service principal.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Optional) ID of the service principal in an external identity provider.`,
				},
				resource.Attribute{
					Name:        "allow_cluster_create",
					Description: `(Optional) Allow the service principal to have [cluster](cluster.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Cluster-usage) and ` + "`" + `cluster_id` + "`" + ` argument. Everyone without ` + "`" + `allow_cluster_create` + "`" + ` argument set, but with [permission to use](permissions.md#Cluster-Policy-usage) Cluster Policy would be able to create clusters, but within the boundaries of that specific policy.`,
				},
				resource.Attribute{
					Name:        "allow_instance_pool_create",
					Description: `(Optional) Allow the service principal to have [instance pool](instance_pool.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Instance-Pool-usage) and [instance_pool_id](permissions.md#instance_pool_id) argument.`,
				},
				resource.Attribute{
					Name:        "databricks_sql_access",
					Description: `(Optional) This is a field to allow the group to have access to [Databricks SQL](https://databricks.com/product/databricks-sql) feature through [databricks_sql_endpoint](sql_endpoint.md).`,
				},
				resource.Attribute{
					Name:        "workspace_access",
					Description: `(Optional) This is a field to allow the group to have access to Databricks Workspace.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Either service principal is active or not. True by default, but can be set to false in case of service principal deactivation with preserving service principal assets.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) Ignore ` + "`" + `cannot create service principal: Service principal with application ID X already exists` + "`" + ` errors and implicitly import the specified service principal into Terraform state, enforcing entitlements defined in the instance of resource. _This functionality is experimental_ and is designed to simplify corner cases, like Azure Active Directory synchronisation.`,
				},
				resource.Attribute{
					Name:        "force_delete_repos",
					Description: `(Optional) This flag determines whether the service principal's repo directory is deleted when the user is deleted. It will have no impact when in the accounts SCIM API. False by default.`,
				},
				resource.Attribute{
					Name:        "force_delete_home_dir",
					Description: `(Optional) This flag determines whether the service principal's home directory is deleted when the user is deleted. It will have no impact when in the accounts SCIM API. False by default. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the service principal. - ` + "`" + `home` + "`" + ` - Home folder of the service principal, e.g. ` + "`" + `/Users/00000000-0000-0000-0000-000000000000` + "`" + `. - ` + "`" + `repos` + "`" + ` - Personal Repos location of the service principal, e.g. ` + "`" + `/Repos/00000000-0000-0000-0000-000000000000` + "`" + `. ## Import The resource scim service principal can be imported using its id, for example ` + "`" + `2345678901234567` + "`" + `. To get the service principal ID, call [Get service principals](https://docs.databricks.com/dev-tools/api/latest/scim/scim-sp.html#get-service-principals). ` + "`" + `` + "`" + `` + "`" + `bash terraform import databricks_service_principal.me <service-principal-id> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the service principal. - ` + "`" + `home` + "`" + ` - Home folder of the service principal, e.g. ` + "`" + `/Users/00000000-0000-0000-0000-000000000000` + "`" + `. - ` + "`" + `repos` + "`" + ` - Personal Repos location of the service principal, e.g. ` + "`" + `/Repos/00000000-0000-0000-0000-000000000000` + "`" + `. ## Import The resource scim service principal can be imported using its id, for example ` + "`" + `2345678901234567` + "`" + `. To get the service principal ID, call [Get service principals](https://docs.databricks.com/dev-tools/api/latest/scim/scim-sp.html#get-service-principals). ` + "`" + `` + "`" + `` + "`" + `bash terraform import databricks_service_principal.me <service-principal-id> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_service_principal_role",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"service",
				"principal",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `(Required) This is the id of the [service principal](service_principal.md) resource.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) This is the id of the role or [instance profile](instance_profile.md) resource. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_service_principal_secret",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"service",
				"principal",
				"secret",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_principal_id",
					Description: `ID of the [databricks_service_principal](service_principal.md) ## Attribute Reference In addition to all arguments above, the following attributes are exported: - ` + "`" + `id` + "`" + ` - ID of the secret - ` + "`" + `secret` + "`" + ` - Generated secret for the service principal ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_share",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"share",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the partition column.`,
				},
				resource.Attribute{
					Name:        "op",
					Description: `The operator to apply for the value, one of: ` + "`" + `EQUAL` + "`" + `, ` + "`" + `LIKE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Time when the share was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The principal that created the share.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the object, one of: ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `PERMISSION_DENIED` + "`" + `. ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Time when the share was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The principal that created the share.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the object, one of: ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `PERMISSION_DENIED` + "`" + `. ## Related Resources The following resources are often used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_alert",
			Category:         "Databricks SQL",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"sql",
				"alert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "query_id",
					Description: `(Required, String) ID of the query evaluated by the alert.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Name of the alert.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Required) Alert configuration options.`,
				},
				resource.Attribute{
					Name:        "column",
					Description: `(Required, String) Name of column in the query result to compare in alert evaluation.`,
				},
				resource.Attribute{
					Name:        "op",
					Description: `(Required, String Enum) Operator used to compare in alert evaluation. (Enum: ` + "`" + `>` + "`" + `, ` + "`" + `>=` + "`" + `, ` + "`" + `<` + "`" + `, ` + "`" + `<=` + "`" + `, ` + "`" + `==` + "`" + `, ` + "`" + `!=` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required, String) Value used to compare in alert evaluation.`,
				},
				resource.Attribute{
					Name:        "muted",
					Description: `(Optional, bool) Whether or not the alert is muted. If an alert is muted, it will not notify users and alert destinations when triggered.`,
				},
				resource.Attribute{
					Name:        "custom_subject",
					Description: `(Optional, String) Custom subject of alert notification, if it exists. This includes email subject, Slack notification header, etc. See [Alerts API reference](https://docs.databricks.com/sql/user/alerts/index.html) for custom templating instructions.`,
				},
				resource.Attribute{
					Name:        "custom_body",
					Description: `(Optional, String) Custom body of alert notification, if it exists. See [Alerts API reference](https://docs.databricks.com/sql/user/alerts/index.html) for custom templating instructions.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `(Optional, String) The identifier of the workspace folder containing the alert. The default is ther user's home folder. The folder identifier is formatted as ` + "`" + `folder/<folder_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rearm",
					Description: `(Optional, Integer) Number of seconds after being triggered before the alert rearms itself and can be triggered again. If not defined, alert will never be triggered again. ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_dashboard",
			Category:         "Databricks SQL",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"sql",
				"dashboard",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_endpoint",
			Category:         "Databricks SQL",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"sql",
				"endpoint",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_global_config",
			Category:         "Databricks SQL",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"sql",
				"global",
				"config",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_permissions",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"sql",
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "database",
					Description: `Name of the database. Has default value of ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "table",
					Description: `Name of the table. Can be combined with ` + "`" + `database` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "view",
					Description: `Name of the view. Can be combined with ` + "`" + `database` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Boolean) If this access control for the entire catalog. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "any_file",
					Description: `(Boolean) If this access control for reading any file. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "anonymous_function",
					Description: `(Boolean) If this access control for using anonymous function. Defaults to ` + "`" + `false` + "`" + `. ### ` + "`" + `privilege_assignments` + "`" + ` blocks You must specify one or many ` + "`" + `privilege_assignments` + "`" + ` configuration blocks to declare ` + "`" + `privileges` + "`" + ` to a ` + "`" + `principal` + "`" + `, which corresponds to ` + "`" + `display_name` + "`" + ` of [databricks_group](group.md#display_name) or [databricks_user](user.md#display_name). Terraform would ensure that only those principals and privileges defined in the resource are applied for the data object and would remove anything else. It would not remove any transitive privileges. ` + "`" + `DENY` + "`" + ` statements are intentionally not supported. Every ` + "`" + `privilege_assignments` + "`" + ` has the following required arguments:`,
				},
				resource.Attribute{
					Name:        "principal",
					Description: `` + "`" + `display_name` + "`" + ` for a [databricks_group](group.md#display_name) or [databricks_user](user.md#display_name), ` + "`" + `application_id` + "`" + ` for a [databricks_service_principal](service_principal.md).`,
				},
				resource.Attribute{
					Name:        "privileges",
					Description: `set of available privilege names in upper case. [Available](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html) privilege names are:`,
				},
				resource.Attribute{
					Name:        "SELECT",
					Description: `gives read access to an object.`,
				},
				resource.Attribute{
					Name:        "CREATE",
					Description: `gives the ability to create an object (for example, a table in a database).`,
				},
				resource.Attribute{
					Name:        "MODIFY",
					Description: `gives the ability to add, delete, and modify data to or from an object.`,
				},
				resource.Attribute{
					Name:        "USAGE",
					Description: `do not give any abilities, but is an additional requirement to perform any action on a database object.`,
				},
				resource.Attribute{
					Name:        "READ_METADATA",
					Description: `gives the ability to view an object and its metadata.`,
				},
				resource.Attribute{
					Name:        "CREATE_NAMED_FUNCTION",
					Description: `gives the ability to create a named UDF in an existing catalog or database.`,
				},
				resource.Attribute{
					Name:        "MODIFY_CLASSPATH",
					Description: `gives the ability to add files to the Spark class path. -> Even though the value ` + "`" + `ALL PRIVILEGES` + "`" + ` is mentioned in Table ACL documentation, it's not recommended to use it from terraform, as it may result in unnecessary state updates. ## Import The resource can be imported using a synthetic identifier. Examples of valid synthetic identifiers are:`,
				},
				resource.Attribute{
					Name:        "table/default.foo",
					Description: `table ` + "`" + `foo` + "`" + ` in a ` + "`" + `default` + "`" + ` database. Database is always mandatory.`,
				},
				resource.Attribute{
					Name:        "view/bar.foo",
					Description: `view ` + "`" + `foo` + "`" + ` in ` + "`" + `bar` + "`" + ` database.`,
				},
				resource.Attribute{
					Name:        "database/bar",
					Description: `` + "`" + `bar` + "`" + ` database.`,
				},
				resource.Attribute{
					Name:        "catalog/",
					Description: `entire catalog. ` + "`" + `/` + "`" + ` suffix is mandatory.`,
				},
				resource.Attribute{
					Name:        "any file/",
					Description: `direct access to any file. ` + "`" + `/` + "`" + ` suffix is mandatory.`,
				},
				resource.Attribute{
					Name:        "anonymous function/",
					Description: `anonymous function. ` + "`" + `/` + "`" + ` suffix is mandatory. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_sql_permissions.foo /<object-type>/<object-name> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_query",
			Category:         "Databricks SQL",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"sql",
				"query",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_table",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"sql",
				"table",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of table relative to parent catalog and schema. Change forces creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Name of parent catalog`,
				},
				resource.Attribute{
					Name:        "schema_name",
					Description: `Name of parent Schema relative to parent Catalog`,
				},
				resource.Attribute{
					Name:        "table_type",
					Description: `Distinguishes a view vs. managed/external Table. ` + "`" + `MANAGED` + "`" + `, ` + "`" + `EXTERNAL` + "`" + ` or ` + "`" + `VIEW` + "`" + `. Change forces creation of a new resource.`,
				},
				resource.Attribute{
					Name:        "storage_location",
					Description: `(Optional) URL of storage location for Table data (required for EXTERNAL Tables). Not supported for ` + "`" + `VIEW` + "`" + ` or ` + "`" + `MANAGED` + "`" + ` table_type.`,
				},
				resource.Attribute{
					Name:        "data_source_format",
					Description: `(Optional) External tables are supported in multiple data source formats. The string constants identifying these formats are ` + "`" + `DELTA` + "`" + `, ` + "`" + `CSV` + "`" + `, ` + "`" + `JSON` + "`" + `, ` + "`" + `AVRO` + "`" + `, ` + "`" + `PARQUET` + "`" + `, ` + "`" + `ORC` + "`" + `, ` + "`" + `TEXT` + "`" + `. Change forces creation of a new resource. Not supported for ` + "`" + `MANAGED` + "`" + ` tables or ` + "`" + `VIEW` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "view_definition",
					Description: `(Optional) SQL text defining the view (for ` + "`" + `table_type == "VIEW"` + "`" + `). Not supported for ` + "`" + `MANAGED` + "`" + ` or ` + "`" + `EXTERNAL` + "`" + ` table_type.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) All table CRUD operations must be executed on a running cluster. If a cluster_id is specified, it will be used to execute SQL commands to manage this table. If empty, a cluster will be created automatically with the name ` + "`" + `terraform-sql-table` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_credential_name",
					Description: `(Optional) For EXTERNAL Tables only: the name of storage credential to use. This cannot be updated`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) User-supplied free-form text. Changing comment is not currently supported on ` + "`" + `VIEW` + "`" + ` table_type.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) Extensible Table properties. ### ` + "`" + `column` + "`" + ` configuration block For table columns Currently, changing the column definitions for a table will require dropping and re-creating the table`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User-visible name of column`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Column type spec (with metadata) as SQL text`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) User-supplied free-form text.`,
				},
				resource.Attribute{
					Name:        "nullable",
					Description: `(Optional) Whether field is nullable (Default: ` + "`" + `true` + "`" + `) ## Import This resource can be imported by name: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_sql_table.this <name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_visualization",
			Category:         "Databricks SQL",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"sql",
				"visualization",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_widget",
			Category:         "Databricks SQL",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"sql",
				"widget",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_storage_credential",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"storage",
				"credential",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_token",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"token",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "lifetime_seconds",
					Description: `(Optional) (Integer) The lifetime of the token, in seconds. If no lifetime is specified, the token remains valid indefinitely.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) (String) Comment that will appear on the user’s settings page for this token. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the token.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the token.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_user",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) This is the username of the given user and will be their form of access and identity.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) This is an alias for the username that can be the full name of the user.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Optional) ID of the user in an external identity provider.`,
				},
				resource.Attribute{
					Name:        "allow_cluster_create",
					Description: `(Optional) Allow the user to have [cluster](cluster.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Cluster-usage) and ` + "`" + `cluster_id` + "`" + ` argument. Everyone without ` + "`" + `allow_cluster_create` + "`" + ` argument set, but with [permission to use](permissions.md#Cluster-Policy-usage) Cluster Policy would be able to create clusters, but within boundaries of that specific policy.`,
				},
				resource.Attribute{
					Name:        "allow_instance_pool_create",
					Description: `(Optional) Allow the user to have [instance pool](instance_pool.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Instance-Pool-usage) and [instance_pool_id](permissions.md#instance_pool_id) argument.`,
				},
				resource.Attribute{
					Name:        "databricks_sql_access",
					Description: `(Optional) This is a field to allow the group to have access to [Databricks SQL](https://databricks.com/product/databricks-sql) feature in User Interface and through [databricks_sql_endpoint](sql_endpoint.md).`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Either user is active or not. True by default, but can be set to false in case of user deactivation with preserving user assets.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) Ignore ` + "`" + `cannot create user: User with username X already exists` + "`" + ` errors and implicitly import the specific user into Terraform state, enforcing entitlements defined in the instance of resource. _This functionality is experimental_ and is designed to simplify corner cases, like Azure Active Directory synchronisation.`,
				},
				resource.Attribute{
					Name:        "force_delete_repos",
					Description: `(Optional) This flag determines whether the user's repo directory is deleted when the user is deleted. It will have no impact when in the accounts SCIM API. False by default.`,
				},
				resource.Attribute{
					Name:        "force_delete_home_dir",
					Description: `(Optional) This flag determines whether the user's home directory is deleted when the user is deleted. It will have not impact when in the accounts SCIM API. False by default. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the user. - ` + "`" + `home` + "`" + ` - Home folder of the user, e.g. ` + "`" + `/Users/mr.foo@example.com` + "`" + `. - ` + "`" + `repos` + "`" + ` - Personal Repos location of the user, e.g. ` + "`" + `/Repos/mr.foo@example.com` + "`" + `. ## Import The resource scim user can be imported using id: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_user.me <user-id> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the user. - ` + "`" + `home` + "`" + ` - Home folder of the user, e.g. ` + "`" + `/Users/mr.foo@example.com` + "`" + `. - ` + "`" + `repos` + "`" + ` - Personal Repos location of the user, e.g. ` + "`" + `/Repos/mr.foo@example.com` + "`" + `. ## Import The resource scim user can be imported using id: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_user.me <user-id> ` + "`" + `` + "`" + `` + "`" + ` ## Related Resources The following resources are often used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_user_instance_profile",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"user",
				"instance",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) This is the id of the [user](user.md) resource.`,
				},
				resource.Attribute{
					Name:        "instance_profile_id",
					Description: `(Required) This is the id of the [instance profile](instance_profile.md) resource. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_user_role",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"user",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) This is the id of the [user](user.md) resource.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) Either a role name or the ARN/ID of the [instance profile](instance_profile.md) resource. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id in the format ` + "`" + `<user_id>|<role>` + "`" + `. ## Import ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id in the format ` + "`" + `<user_id>|<role>` + "`" + `. ## Import ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_workspace_conf",
			Category:         "Workspace",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"workspace",
				"conf",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "custom_config",
					Description: `(Required) Key-value map of strings that represent workspace configuration. Upon resource deletion, properties that start with ` + "`" + `enable` + "`" + ` or ` + "`" + `enforce` + "`" + ` will be reset to ` + "`" + `false` + "`" + ` value, regardless of initial default one. ## Import ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_workspace_file",
			Category:         "Workspace",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"workspace",
				"file",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The absolute path of the workspace file, beginning with "/", e.g. "/Demo".`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `Path to file on local filesystem. Conflicts with ` + "`" + `content_base64` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `The base64-encoded file content. Conflicts with ` + "`" + `source` + "`" + `. Use of ` + "`" + `content_base64` + "`" + ` is discouraged, as it's increasing memory footprint of Terraform state and should only be used in exceptional circumstances, like creating a workspace file with configuration properties for a data pipeline. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Path of workspace file`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Routable URL of the workspace file`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `Unique identifier for a workspace file ## Access Control`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Path of workspace file`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Routable URL of the workspace file`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `Unique identifier for a workspace file ## Access Control`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"databricks_catalog":                     0,
		"databricks_cluster":                     1,
		"databricks_cluster_policy":              2,
		"databricks_dbfs_file":                   3,
		"databricks_directory":                   4,
		"databricks_entitlements":                5,
		"databricks_external_location":           6,
		"databricks_git_credential":              7,
		"databricks_global_init_script":          8,
		"databricks_grants":                      9,
		"databricks_group":                       10,
		"databricks_group_instance_profile":      11,
		"databricks_group_member":                12,
		"databricks_group_role":                  13,
		"databricks_instance_pool":               14,
		"databricks_instance_profile":            15,
		"databricks_ip_access_list":              16,
		"databricks_job":                         17,
		"databricks_library":                     18,
		"databricks_metastore":                   19,
		"databricks_metastore_assignment":        20,
		"databricks_metastore_data_access":       21,
		"databricks_mlflow_experiment":           22,
		"databricks_mlflow_model":                23,
		"databricks_mlflow_webhook":              24,
		"databricks_model_serving":               25,
		"databricks_mount":                       26,
		"databricks_mws_credentials":             27,
		"databricks_mws_customer_managed_keys":   28,
		"databricks_mws_log_delivery":            29,
		"databricks_mws_networks":                30,
		"databricks_mws_permission_assignment":   31,
		"databricks_mws_private_access_settings": 32,
		"databricks_mws_storage_configurations":  33,
		"databricks_mws_vpc_endpoint":            34,
		"databricks_mws_workspaces":              35,
		"databricks_notebook":                    36,
		"databricks_obo_token":                   37,
		"databricks_permission_assignment":       38,
		"databricks_permissions":                 39,
		"databricks_pipeline":                    40,
		"databricks_provider":                    41,
		"databricks_recipient":                   42,
		"databricks_repo":                        43,
		"databricks_schema":                      44,
		"databricks_secret":                      45,
		"databricks_secret_acl":                  46,
		"databricks_secret_scope":                47,
		"databricks_service_principal":           48,
		"databricks_service_principal_role":      49,
		"databricks_service_principal_secret":    50,
		"databricks_share":                       51,
		"databricks_sql_alert":                   52,
		"databricks_sql_dashboard":               53,
		"databricks_sql_endpoint":                54,
		"databricks_sql_global_config":           55,
		"databricks_sql_permissions":             56,
		"databricks_sql_query":                   57,
		"databricks_sql_table":                   58,
		"databricks_sql_visualization":           59,
		"databricks_sql_widget":                  60,
		"databricks_storage_credential":          61,
		"databricks_token":                       62,
		"databricks_user":                        63,
		"databricks_user_instance_profile":       64,
		"databricks_user_role":                   65,
		"databricks_workspace_conf":              66,
		"databricks_workspace_file":              67,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
