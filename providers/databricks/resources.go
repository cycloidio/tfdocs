package databricks

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "databricks_aws_s3_mount",
			Category:         "AWS",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"aws",
				"s3",
				"mount",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) (String) [Cluster](cluster.md) to use for mounting. If no cluster is specified, a new cluster will be created and will mount the bucket for all of the clusters in this workspace. If a cluster is specified, mount will be visible for all clusters with the same [instance profile](./instance_profile.md). If the cluster is not running - it's going to be started, so be aware to set auto-termination rules on it.`,
				},
				resource.Attribute{
					Name:        "instance_profile",
					Description: `(Optional) (String) ARN of registered [instance profile](instance_profile.md) for data access.`,
				},
				resource.Attribute{
					Name:        "mount_name",
					Description: `(Required) (String) Name, under which mount will be accessible in ` + "`" + `dbfs:/mnt/<MOUNT_NAME>` + "`" + ` or locally on each instance through FUSE mount ` + "`" + `/dbfs/mnt/<MOUNT_NAME>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "s3_bucket_name",
					Description: `(Required) (String) S3 bucket name to be mounted. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `mount name`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(String) HDFS-compatible S3 bucket url ` + "`" + `s3a://<s3_bucket_name>` + "`" + ` ## Import The resource aws s3 mount can be imported using it's mount name ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_aws_s3_mount.this <mount_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `mount name`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(String) HDFS-compatible S3 bucket url ` + "`" + `s3a://<s3_bucket_name>` + "`" + ` ## Import The resource aws s3 mount can be imported using it's mount name ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_aws_s3_mount.this <mount_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_azure_adls_gen1_mount",
			Category:         "Azure",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"azure",
				"adls",
				"gen1",
				"mount",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) (String) This is the client_id for the enterprise application for the service principal.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) (String) This is your azure directory tenant id. This is required for creating the mount.`,
				},
				resource.Attribute{
					Name:        "client_secret_key",
					Description: `(Required) (String) This is the secret key in which your service principal/enterprise app client secret will be stored.`,
				},
				resource.Attribute{
					Name:        "client_secret_scope",
					Description: `(Required) (String) This is the secret scope in which your service principal/enterprise app client secret will be stored.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) (String) Cluster to use for mounting. If no cluster is specified, a new cluster will be created and will mount the bucket for all of the clusters in this workspace. If the cluster is not running - it's going to be started, so be aware to set auto-termination rules on it.`,
				},
				resource.Attribute{
					Name:        "mount_name",
					Description: `(Required) (String) Name, under which mount will be accessible in ` + "`" + `dbfs:/mnt/<MOUNT_NAME>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_resource_name",
					Description: `(Required) (String) The name of the storage resource in which the data is for ADLS gen 1. This is what you are trying to mount.`,
				},
				resource.Attribute{
					Name:        "spark_conf_prefix",
					Description: `(Optional) (String) This is the spark configuration prefix for adls gen 1 mount. The options are ` + "`" + `fs.adl` + "`" + `, ` + "`" + `dfs.adls` + "`" + `. Use ` + "`" + `fs.adl` + "`" + ` for runtime 6.0 and above for the clusters. Otherwise use ` + "`" + `dfs.adls` + "`" + `. The default value is: ` + "`" + `fs.adl` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "directory",
					Description: `(Computed) (String) This is optional if you want to add an additional directory that you wish to mount. This must start with a "/". ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `mount name`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(String) HDFS-compatible url ` + "`" + `adl://<adlsv1-account>` + "`" + ` ## Import The resource can be imported using it's mount name ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_azure_adls_gen1_mount.this <mount_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `mount name`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(String) HDFS-compatible url ` + "`" + `adl://<adlsv1-account>` + "`" + ` ## Import The resource can be imported using it's mount name ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_azure_adls_gen1_mount.this <mount_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_azure_adls_gen2_mount",
			Category:         "Azure",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"azure",
				"adls",
				"gen2",
				"mount",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) (String) This is the client_id (Application Object ID) for the enterprise application for the service principal.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) (String) This is your azure directory tenant id. This is required for creating the mount.`,
				},
				resource.Attribute{
					Name:        "client_secret_key",
					Description: `(Required) (String) This is the secret key in which your service principal/enterprise app client secret will be stored.`,
				},
				resource.Attribute{
					Name:        "client_secret_scope",
					Description: `(Required) (String) This is the secret scope in which your service principal/enterprise app client secret will be stored.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) (String) Cluster to use for mounting. If no cluster is specified, a new cluster will be created and will mount the bucket for all of the clusters in this workspace. If the cluster is not running - it's going to be started, so be aware to set auto-termination rules on it.`,
				},
				resource.Attribute{
					Name:        "container_name",
					Description: `(Required) (String) ADLS gen2 container name`,
				},
				resource.Attribute{
					Name:        "storage_account_name",
					Description: `(Required) (String) The name of the storage resource in which the data is.`,
				},
				resource.Attribute{
					Name:        "mount_name",
					Description: `(Required) (String) Name, under which mount will be accessible in ` + "`" + `dbfs:/mnt/<MOUNT_NAME>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "directory",
					Description: `(Computed) (String) This is optional if you want to add an additional directory that you wish to mount. This must start with a "/".`,
				},
				resource.Attribute{
					Name:        "initialize_file_system",
					Description: `(Required) (Bool) either or not initialize FS for the first use ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `mount name`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(String) HDFS-compatible url ` + "`" + `abfss://<adlsv2-account>` + "`" + ` ## Import The resource can be imported using it's mount name ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_azure_adls_gen2_mount.this <mount_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `mount name`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(String) HDFS-compatible url ` + "`" + `abfss://<adlsv2-account>` + "`" + ` ## Import The resource can be imported using it's mount name ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_azure_adls_gen2_mount.this <mount_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_azure_blob_mount",
			Category:         "Azure",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"azure",
				"blob",
				"mount",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Required) (String) This is the auth type for blob storage. This can either be SAS tokens or account access keys.`,
				},
				resource.Attribute{
					Name:        "token_secret_scope",
					Description: `(Required) (String) This is the secret scope in which your auth type token is stored.`,
				},
				resource.Attribute{
					Name:        "token_secret_key",
					Description: `(Required) (String) This is the secret key in which your auth type token is stored.`,
				},
				resource.Attribute{
					Name:        "container_name",
					Description: `(Required) (String) The container in which the data is. This is what you are trying to mount.`,
				},
				resource.Attribute{
					Name:        "storage_account_name",
					Description: `(Required) (String) The name of the storage resource in which the data is.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) (String) Cluster to use for mounting. If no cluster is specified, a new cluster will be created and will mount the bucket for all of the clusters in this workspace. If the cluster is not running - it's going to be started, so be aware to set auto-termination rules on it.`,
				},
				resource.Attribute{
					Name:        "mount_name",
					Description: `(Required) (String) Name, under which mount will be accessible in ` + "`" + `dbfs:/mnt/<MOUNT_NAME>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "directory",
					Description: `(Computed) (String) This is optional if you want to add an additional directory that you wish to mount. This must start with a "/". ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `mount name`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(String) HDFS-compatible url ` + "`" + `wasbs://<adlsv2-account>` + "`" + ` ## Import The resource can be imported using it's mount name ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_azure_blob_mount.this <mount_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `mount name`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(String) HDFS-compatible url ` + "`" + `wasbs://<adlsv2-account>` + "`" + ` ## Import The resource can be imported using it's mount name ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_azure_blob_mount.this <mount_name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
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
					Name:        "cluster_name",
					Description: `(Optional) Cluster name, which doesn’t have to be unique. If not specified at creation, the cluster name will be an empty string.`,
				},
				resource.Attribute{
					Name:        "spark_version",
					Description: `(Required) [Runtime version](https://docs.databricks.com/runtime/index.html) of the cluster. Any supported [databricks_spark_version](../data-sources/spark_version.md) id. We advise using [Cluster Policies](cluster_policy.md) to restrict the list of versions for simplicity while maintaining enough control.`,
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
					Name:        "autotermination_minutes",
					Description: `(Optional) Automatically terminate the cluster after being inactive for this time in minutes. If not set, Databricks won't automatically terminate an inactive cluster. If specified, the threshold must be between 10 and 10000 minutes. You can also set this value to 0 to explicitly disable automatic termination. _We highly recommend having this setting present for Interactive/BI clusters._`,
				},
				resource.Attribute{
					Name:        "enable_elastic_disk",
					Description: `(Optional) If you don’t want to allocate a fixed number of EBS volumes at cluster creation time, use autoscaling local storage. With autoscaling local storage, Databricks monitors the amount of free disk space available on your cluster’s Spark workers. If a worker begins to run too low on disk, Databricks automatically attaches a new EBS volume to the worker before it runs out of disk space. EBS volumes are attached up to a limit of 5 TB of total disk space per instance (including the instance’s local storage). To scale down EBS usage, make sure you have ` + "`" + `autotermination_minutes` + "`" + ` and ` + "`" + `autoscale` + "`" + ` attributes set. More documentation available at [cluster configuration page](https://docs.databricks.com/clusters/configure.html#autoscaling-local-storage-1).`,
				},
				resource.Attribute{
					Name:        "enable_local_disk_encryption",
					Description: `(Optional) Some instance types you use to run clusters may have locally attached disks. Databricks may store shuffle data or temporary data on these locally attached disks. To ensure that all data at rest is encrypted for all storage types, including shuffle data stored temporarily on your cluster’s local disks, you can enable local disk encryption. When local disk encryption is enabled, Databricks generates an encryption key locally unique to each cluster node and encrypting all data stored on local disks. The scope of the key is local to each cluster node and is destroyed along with the cluster node itself. During its lifetime, the key resides in memory for encryption and decryption and is stored encrypted on the disk. _Your workloads may run more slowly because of the performance impact of reading and writing encrypted data to and from local volumes. This feature is not available for all Azure Databricks subscriptions. Contact your Microsoft or Databricks account representative to request access._`,
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
					Description: `(Optional) Additional tags for cluster resources. Databricks will tag all cluster resources (e.g., AWS EC2 instances and EBS volumes) with these tags in addition to ` + "`" + `default_tags` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spark_conf",
					Description: `(Optional) Map with key-value pairs to fine-tune Spark clusters, where you can provide custom [Spark configuration properties](https://spark.apache.org/docs/latest/configuration.html) in a cluster configuration.`,
				},
				resource.Attribute{
					Name:        "is_pinned",
					Description: `(Optional) boolean value specifying if cluster is pinned (not pinned by default). You must be a Databricks administrator to use this. The pinned clusters' maximum number is [limited to 20](https://docs.databricks.com/clusters/clusters-manage.html#pin-a-cluster), so ` + "`" + `apply` + "`" + ` may fail if you have more than that. The following example demonstrates how to create an autoscaling cluster with [Delta Cache](https://docs.databricks.com/delta/optimizations/delta-cache.html) enabled: ` + "`" + `` + "`" + `` + "`" + `hcl data "databricks_node_type" "smallest" { local_disk = true } data "databricks_spark_version" "latest_lts" { long_term_support = true } resource "databricks_cluster" "shared_autoscaling" { cluster_name = "Shared Autoscaling" spark_version = data.databricks_spark_version.latest_lts.id node_type_id = data.databricks_node_type.smallest.id autotermination_minutes = 20 autoscale { min_workers = 1 max_workers = 50 } spark_conf = { "spark.databricks.io.cache.enabled": true, "spark.databricks.io.cache.maxDiskUsage": "50g", "spark.databricks.io.cache.maxMetaDataCache": "1g" } } ` + "`" + `` + "`" + `` + "`" + ` ## Fixed size or autoscaling cluster When you [create a Databricks cluster](https://docs.databricks.com/clusters/configure.html#cluster-size-and-autoscaling), you can either provide a ` + "`" + `num_workers` + "`" + ` for the fixed-size cluster or provide ` + "`" + `min_workers` + "`" + ` and/or ` + "`" + `max_workers` + "`" + ` for the cluster within the ` + "`" + `autoscale` + "`" + ` group. When you give a fixed-sized cluster, Databricks ensures that your cluster has a specified number of workers. When you provide a range for the number of workers, Databricks chooses the appropriate number of workers required to run your job - also known as "autoscaling." With autoscaling, Databricks dynamically reallocates workers to account for the characteristics of your job. Certain parts of your pipeline may be more computationally demanding than others, and Databricks automatically adds additional workers during these phases of your job (and removes them when they’re no longer needed). ` + "`" + `autoscale` + "`" + ` optional configuration block supports the following:`,
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
					Description: `(Optional) S3 endpoint, e.g. https://s3-us-west-2.amazonaws.com. Either ` + "`" + `region` + "`" + ` or ` + "`" + `endpoint` + "`" + ` needs to be set. If both are set, the endpoint is used.`,
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
					Description: `(Optional) Set canned access control list, e.g. ` + "`" + `bucket-owner-full-control` + "`" + `. If ` + "`" + `canned_cal` + "`" + ` is set, the cluster instance profile must have ` + "`" + `s3:PutObjectAcl` + "`" + ` permission on the destination bucket and prefix. The full list of possible canned ACLs can be found [here](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl). By default, only the object owner gets full control. If you are using a cross-account role for writing data, you may want to set ` + "`" + `bucket-owner-full-control` + "`" + ` to make bucket owners able to read the logs. ## init_scripts You can specify up to 10 different init scripts for the specific cluster. If you want a shell script to run on all clusters and jobs within the same workspace, you should consider [databricks_global_init_script](global_init_script.md). Example of taking init script from DBFS: ` + "`" + `` + "`" + `` + "`" + `hcl init_scripts { dbfs { destination = "dbfs://init-scripts/install-elk.sh" } } ` + "`" + `` + "`" + `` + "`" + ` Example of taking init script from S3: ` + "`" + `` + "`" + `` + "`" + `hcl init_scripts { s3 { destination = "s3a://acmecorp-main/init-scripts/install-elk.sh" region = "us-east-1" } } ` + "`" + `` + "`" + `` + "`" + ` Like the ` + "`" + `cluster_log_conf` + "`" + ` configuration block, init scripts support S3 and DBFS locations. In addition, you can also specify a local file as follows: ` + "`" + `` + "`" + `` + "`" + `hcl init_scripts { file { destination = "file:/my/local/file.sh" } } ` + "`" + `` + "`" + `` + "`" + ` ## aws_attributes ` + "`" + `aws_attributes` + "`" + ` optional configuration block contains attributes related to [clusters running on Amazon Web Services](https://docs.databricks.com/clusters/configure.html#aws-configurations). Here is the example of shared autoscaling cluster with some of AWS options set: ` + "`" + `` + "`" + `` + "`" + `hcl resource "databricks_cluster" "this" { cluster_name = "Shared Autoscaling" spark_version = "6.6.x-scala2.11" node_type_id = "i3.xlarge" autotermination_minutes = 20 autoscale { min_workers = 1 max_workers = 50 } aws_attributes { availability = "SPOT" zone_id = "us-east-1" first_on_demand = 1 spot_bid_price_percent = 100 } } ` + "`" + `` + "`" + `` + "`" + ` The following options are available:`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) Identifier for the availability zone/datacenter in which the cluster resides. This string will be of a form like “us-west-2a”. The provided availability zone must be in the same region as the Databricks deployment. For example, “us-west-2a” is not a valid zone ID if the Databricks deployment resides in the “us-east-1” region.`,
				},
				resource.Attribute{
					Name:        "availability",
					Description: `(Optional) Availability type used for all subsequent nodes past the ` + "`" + `first_on_demand` + "`" + ` ones. Valid values are ` + "`" + `SPOT` + "`" + `, ` + "`" + `SPOT_WITH_FALLBACK` + "`" + ` and ` + "`" + `ON_DEMAND` + "`" + `. Note: If ` + "`" + `first_on_demand` + "`" + ` is zero, this availability type will be used for the entire cluster.`,
				},
				resource.Attribute{
					Name:        "first_on_demand",
					Description: `(Optional) The first ` + "`" + `first_on_demand` + "`" + ` nodes of the cluster will be placed on on-demand instances. If this value is greater than 0, the cluster driver node will be placed on an on-demand instance. If this value is greater than or equal to the current cluster size, all nodes will be placed on on-demand instances. If this value is less than the current cluster size, ` + "`" + `first_on_demand` + "`" + ` nodes will be placed on on-demand instances, and the remainder will be placed on availability instances. This value does not affect cluster size and cannot be mutated over the lifetime of a cluster.`,
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
					Description: `(Optional) The type of EBS volumes that will be launched with this cluster. Valid values are ` + "`" + `GENERAL_PURPOSE_SSD` + "`" + ` or ` + "`" + `THROUGHPUT_OPTIMIZED_HDD` + "`" + `. Use this option only if you're not picking _Delta Optimized ` + "`" + `i3.`,
				},
				resource.Attribute{
					Name:        "ebs_volume_count",
					Description: `(Optional) The number of volumes launched for each instance. You can choose up to 10 volumes. This feature is only enabled for supported node types. Legacy node types cannot specify custom EBS volumes. For node types with no instance store, at least one EBS volume needs to be specified; otherwise, cluster creation will fail. These EBS volumes will be mounted at /ebs0, /ebs1, and etc. Instance store volumes will be mounted at /local_disk0, /local_disk1, and etc. If EBS volumes are attached, Databricks will configure Spark to use only the EBS volumes for scratch storage because heterogeneously sized scratch devices can lead to inefficient disk utilization. If no EBS volumes are attached, Databricks will configure Spark to use instance store volumes. If EBS volumes are specified, then the Spark configuration spark.local.dir will be overridden.`,
				},
				resource.Attribute{
					Name:        "ebs_volume_size",
					Description: `(Optional) The size of each EBS volume (in GiB) launched for each instance. For general purpose SSD, this value must be within the range 100 - 4096. For throughput optimized HDD, this value must be within the range 500 - 4096. Custom EBS volumes cannot be specified for the legacy node types (memory-optimized and compute-optimized). ## azure_attributes ` + "`" + `azure_attributes` + "`" + ` optional configuration block contains attributes related to [clusters running on Azure](https://docs.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/clusters#--azureattributes). Here is the example of shared autoscaling cluster with some of AWS options set: ` + "`" + `` + "`" + `` + "`" + `hcl resource "databricks_cluster" "this" { cluster_name = "Shared Autoscaling" spark_version = "6.6.x-scala2.11" node_type_id = "Standard_DS3_v2" autotermination_minutes = 20 autoscale { min_workers = 1 max_workers = 50 } azure_attributes { availability = "SPOT_AZURE" first_on_demand = 1 spot_bid_max_price = 100 } } ` + "`" + `` + "`" + `` + "`" + ` The following options are [available](https://docs.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/clusters#--azureattributes):`,
				},
				resource.Attribute{
					Name:        "availability",
					Description: `(Optional) Availability type used for all subsequent nodes past the ` + "`" + `first_on_demand` + "`" + ` ones. Valid values are ` + "`" + `SPOT_AZURE` + "`" + `, ` + "`" + `SPOT_WITH_FALLBACK` + "`" + `, and ` + "`" + `ON_DEMAND_AZURE` + "`" + `. Note: If ` + "`" + `first_on_demand` + "`" + ` is zero, this availability type will be used for the entire cluster.`,
				},
				resource.Attribute{
					Name:        "first_on_demand",
					Description: `(Optional) The first ` + "`" + `first_on_demand` + "`" + ` nodes of the cluster will be placed on on-demand instances. If this value is greater than 0, the cluster driver node will be placed on an on-demand instance. If this value is greater than or equal to the current cluster size, all nodes will be placed on on-demand instances. If this value is less than the current cluster size, ` + "`" + `first_on_demand` + "`" + ` nodes will be placed on on-demand instances, and the remainder will be placed on availability instances. This value does not affect cluster size and cannot be mutated over the lifetime of a cluster.`,
				},
				resource.Attribute{
					Name:        "spot_bid_max_price",
					Description: `(Optional) The max price for Azure spot instances. Use ` + "`" + `-1` + "`" + ` to specify lowest price. ## gcp_attributes ` + "`" + `gcp_attributes` + "`" + ` optional configuration block contains attributes related to [clusters running on GCP](https://docs.gcp.databricks.com/dev-tools/api/latest/clusters.html#clustergcpattributes). The following options are available:`,
				},
				resource.Attribute{
					Name:        "use_preemptible_executors",
					Description: `(Optional, bool) if we should use preemptible executors ([GCP documentation](https://cloud.google.com/compute/docs/instances/preemptible))`,
				},
				resource.Attribute{
					Name:        "google_service_account",
					Description: `(Optional, string) Google Service Account email address that the cluster uses to authenticate with Google Identity. This field is used for authentication with the GCS and BigQuery data sources. ## docker_image [Databricks Container Services](https://docs.databricks.com/clusters/custom-containers.html) lets you specify a Docker image when you create a cluster. You need to enable Container Services in`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL for the Docker image`,
				},
				resource.Attribute{
					Name:        "basic_auth",
					Description: `(Optional) ` + "`" + `basic_auth.username` + "`" + ` and ` + "`" + `basic_auth.password` + "`" + ` for Docker repository. Docker registry credentials are encrypted when they are stored in Databricks internal storage and when they are passed to a registry upon fetching Docker images at cluster launch. However, other authenticated and authorized API users of this workspace can access the username and password. Example usage with [azurerm_container_registry](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/container_registry) and [docker_registry_image](https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs/resources/registry_image), that you can adapt to your specific use-case: ` + "`" + `` + "`" + `` + "`" + `hcl resource "docker_registry_image" "this" { name = "${azurerm_container_registry.this.login_server}/sample:latest" build { # ... } } resource "databricks_cluster" "this" { # ... docker_image { url = docker_registry_image.this.name basic_auth { username = azurerm_container_registry.this.admin_username password = azurerm_container_registry.this.admin_password } } } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the cluster.`,
				},
				resource.Attribute{
					Name:        "default_tags",
					Description: `(map) Tags that are added by Databricks by default, regardless of any custom_tags that may have been added. These include: Vendor: Databricks, Creator: <username_of_creator>, ClusterName: <name_of_cluster>, ClusterId: <id_of_cluster>, Name: <Databricks internal use>`,
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
					Description: `(map) Tags that are added by Databricks by default, regardless of any custom_tags that may have been added. These include: Vendor: Databricks, Creator: <username_of_creator>, ClusterName: <name_of_cluster>, ClusterId: <id_of_cluster>, Name: <Databricks internal use>`,
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
					Name:        "definition",
					Description: `(Required) Policy definition JSON document expressed in [Databricks Policy Definition Language](https://docs.databricks.com/administration-guide/clusters/policies.html#cluster-policy-definition). ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the cluster policy. This is equal to policy_id.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Canonical unique identifier for the cluster policy. ## Import The resource cluster policy can be imported using the policy id: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_cluster_policy.this <cluster-policy-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the cluster policy. This is equal to policy_id.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Canonical unique identifier for the cluster policy. ## Import The resource cluster policy can be imported using the policy id: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_cluster_policy.this <cluster-policy-id> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Path, but with ` + "`" + `dbfs:` + "`" + ` prefix ## Import The resource dbfs file can be imported using the path of the file ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_dbfs_file.this <path> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Path, but with ` + "`" + `dbfs:` + "`" + ` prefix ## Import The resource dbfs file can be imported using the path of the file ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_dbfs_file.this <path> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `ID assigned to a global init script by API ## Access Control Global init scripts are available only for administrators, so you can't change permissions for it. ## Import The resource global init script can be imported using script ID: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_global_init_script.this script_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID assigned to a global init script by API ## Access Control Global init scripts are available only for administrators, so you can't change permissions for it. ## Import The resource global init script can be imported using script ID: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_global_init_script.this script_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
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
					Name:        "allow_cluster_create",
					Description: `(Optional) This is a field to allow the group to have [cluster](cluster.md) create privileges. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Cluster-usage) and [cluster_id](permissions.md#cluster_id) argument. Everyone without ` + "`" + `allow_cluster_create` + "`" + ` argument set, but with [permission to use](permissions.md#Cluster-Policy-usage) Cluster Policy would be able to create clusters, but within boundaries of that specific policy.`,
				},
				resource.Attribute{
					Name:        "allow_instance_pool_create",
					Description: `(Optional) This is a field to allow the group to have [instance pool](instance_pool.md) create privileges. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Instance-Pool-usage) and [instance_pool_id](permissions.md#instance_pool_id) argument.`,
				},
				resource.Attribute{
					Name:        "allow_sql_analytics_access",
					Description: `(Optional) This is a field to allow the group to have access to [SQL Analytics](https://databricks.com/product/sql-analytics) feature through [databricks_sql_endpoint](sql_endpoint.md). ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
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
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Required) This is the id of the [group](group.md) or [user](user.md). ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `(Optional) (Integer) The maximum number of instances the pool can contain, including both idle instances and ones in use by clusters. Once the maximum capacity is reached, you cannot create new clusters from the pool and existing clusters cannot autoscale up until some instances are made idle in the pool via [cluster](cluster.md) termination or down-scaling.`,
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
					Description: `(Optional) (Map) Additional tags for instance pool resources. Databricks tags all pool resources (e.g. AWS & Azure instances and Disk volumes).`,
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
					Description: `(Required) (String) Identifier for the availability zone/datacenter in which the instance pool resides. This string is of a form like ` + "`" + `"us-west-2a"` + "`" + `. The provided availability zone must be in the same region as the Databricks deployment. For example, ` + "`" + `"us-west-2a"` + "`" + ` is not a valid zone ID if the Databricks deployment resides in the ` + "`" + `"us-east-1"` + "`" + ` region. This is an optional field. If not specified, a default zone is used. You can find the list of available zones as well as the default value by using the [List Zones API](https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterclusterservicelistavailablezones).`,
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
					Description: `(Optional) Availability type used for all subsequent nodes past the ` + "`" + `first_on_demand` + "`" + ` ones. Valid values are ` + "`" + `SPOT_AZURE` + "`" + ` and ` + "`" + `ON_DEMAND_AZURE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_bid_max_price",
					Description: `(Optional) The max price for Azure spot instances. Use ` + "`" + `-1` + "`" + ` to specify lowest price. ### disk_spec Configuration Block For disk_spec make sure to use`,
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
			Category:         "AWS",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"aws",
				"instance",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_profile_arn",
					Description: `(Required) ` + "`" + `ARN` + "`" + ` attribute of ` + "`" + `aws_iam_instance_profile` + "`" + ` output, the EC2 instance profile association to AWS IAM role. This ARN would be validated upon resource creation and it's not possible to skip validation. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ARN for EC2 Instance Profile, that is registered with Databricks. ## Import The resource instance profile can be imported using the ARN of it ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_instance_profile.this <instance-profile-arn> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ARN for EC2 Instance Profile, that is registered with Databricks. ## Import The resource instance profile can be imported using the ARN of it ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_instance_profile.this <instance-profile-arn> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Can only be "ALLOW" or "BLOCK"`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `This is a field to allow the group to have instance pool create priviliges.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) This is the display name for the given IP ACL List.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Boolean ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + ` indicating whether this list should be active. Defaults to ` + "`" + `true` + "`" + ` ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "list_id",
					Description: `Canonical unique identifier for the IP Access List. ## Import The databricks_ip_access_list can be imported using id: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_ip_access_list.this <list-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list_id",
					Description: `Canonical unique identifier for the IP Access List. ## Import The databricks_ip_access_list can be imported using id: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_ip_access_list.this <list-id> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "new_cluster",
					Description: `(Optional) Same set of parameters as for [databricks_cluster](cluster.md) resource.`,
				},
				resource.Attribute{
					Name:        "existing_cluster_id",
					Description: `(Optional) If existing_cluster_id, the ID of an existing [cluster](cluster.md) that will be used for all runs of this job. When running jobs on an existing cluster, you may need to manually restart the cluster if it stops responding. We strongly suggest to use ` + "`" + `new_cluster` + "`" + ` for greater reliability.`,
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
					Description: `(Optional) (Integer) An optional maximum number of times to retry an unsuccessful run. A run is considered to be unsuccessful if it completes with a FAILED result_state or INTERNAL_ERROR life_cycle_state. The value -1 means to retry indefinitely and the value 0 means to never retry. The default behavior is to never retry.`,
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
					Description: `(Optional) (Integer) An optional maximum allowed number of concurrent runs of the job.`,
				},
				resource.Attribute{
					Name:        "email_notifications",
					Description: `(Optional) (List) An optional set of email addresses notified when runs of this job begin and complete and when this job is deleted. The default behavior is to not send any emails. This field is a block and is documented below.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Optional) (List) An optional periodic schedule for this job. The default behavior is that the job runs when triggered by clicking Run Now in the Jobs UI or sending an API request to runNow. This field is a block and is documented below. ### schedule Configuration Block`,
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
					Description: `(Optional) Indicate whether this schedule is paused or not. Either “PAUSED” or “UNPAUSED”. When the pause_status field is omitted and a schedule is provided, the server will default to using "UNPAUSED" as a value for pause_status. ### spark_jar_task Configuration Block`,
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
					Description: `(Required) The URI of the Python file to be executed. [databricks_dbfs_file](dbfs_file.md#path) and S3 paths are supported. This field is required.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) (List) Command line parameters passed to the Python file. ### notebook_task Configuration Block`,
				},
				resource.Attribute{
					Name:        "base_parameters",
					Description: `(Optional) (Map) Base parameters to be used for each run of this job. If the run is initiated by a call to run-now with parameters specified, the two parameters maps will be merged. If the same key is specified in base_parameters and in run-now, the value from run-now will be used. If the notebook takes a parameter that is not specified in the job’s base_parameters or the run-now override parameters, the default value from the notebook will be used. Retrieve these parameters in a notebook using ` + "`" + `dbutils.widgets.get` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notebook_path",
					Description: `(Required) The absolute path of the [databricks_notebook](notebook.md#path) to be run in the Databricks workspace. This path must begin with a slash. This field is required. ### email_notifications Configuration Block`,
				},
				resource.Attribute{
					Name:        "on_failure",
					Description: `(Optional) (List) list of emails to notify on failure`,
				},
				resource.Attribute{
					Name:        "no_alert_for_skipped_runs",
					Description: `(Optional) (Bool) don't send alert for skipped runs`,
				},
				resource.Attribute{
					Name:        "on_start",
					Description: `(Optional) (List) list of emails to notify on failure`,
				},
				resource.Attribute{
					Name:        "on_success",
					Description: `(Optional) (List) list of emails to notify on failure ## Access Control By default, all users can create and modify jobs unless an administrator [enables jobs access control](https://docs.databricks.com/administration-guide/access-control/jobs-acl.html). With jobs access control, individual permissions determine a user’s abilities.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_credentials",
			Category:         "AWS",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"aws",
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
					Description: `(String) identifier of credentials`,
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
					Description: `(String) identifier of credentials`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_customer_managed_keys",
			Category:         "AWS",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"aws",
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
					Description: `(String) ID of the notebook encryption key configuration object.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(Integer) Time in epoch milliseconds when the customer key was created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the mws customer managed keys.`,
				},
				resource.Attribute{
					Name:        "customer_managed_key_id",
					Description: `(String) ID of the notebook encryption key configuration object.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `(Integer) Time in epoch milliseconds when the customer key was created.`,
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
			Category:         "AWS",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"aws",
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
					Description: `name under which this network is regisstered`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `[aws_vpc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc) id`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `ids of [aws_subnet](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/subnet)`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `ids of [aws_security_group](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the mws networks.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(String) id of network to be used for [databricks_mws_workspace](mws_workspaces.md) resource.`,
				},
				resource.Attribute{
					Name:        "vpc_status",
					Description: `(String) VPC attachment status`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `(Integer) id of associated workspace`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the mws networks.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(String) id of network to be used for [databricks_mws_workspace](mws_workspaces.md) resource.`,
				},
				resource.Attribute{
					Name:        "vpc_status",
					Description: `(String) VPC attachment status`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `(Integer) id of associated workspace`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_private_access_settings",
			Category:         "AWS",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"aws",
				"mws",
				"private",
				"access",
				"settings",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/)`,
				},
				resource.Attribute{
					Name:        "private_access_settings_name",
					Description: `Name of Private Access Settings in Databricks Account`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of AWS VPC ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "private_access_settings_id",
					Description: `Canonical unique identifier of Private Access Settings in Databricks Account`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of Private Access Settings`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "private_access_settings_id",
					Description: `Canonical unique identifier of Private Access Settings in Databricks Account`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of Private Access Settings`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_storage_configurations",
			Category:         "AWS",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"aws",
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
					Description: `(String) id of storage config to be used for ` + "`" + `databricks_mws_workspace` + "`" + ` resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the mws storage configurations.`,
				},
				resource.Attribute{
					Name:        "storage_configuration_id",
					Description: `(String) id of storage config to be used for ` + "`" + `databricks_mws_workspace` + "`" + ` resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_vpc_endpoint",
			Category:         "AWS",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"aws",
				"mws",
				"vpc",
				"endpoint",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/)`,
				},
				resource.Attribute{
					Name:        "aws_vpc_endpoint_id",
					Description: `ID of configured [aws_vpc_endpoint](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_endpoint)`,
				},
				resource.Attribute{
					Name:        "vpc_endpoint_name",
					Description: `Name of VPC Endpoint in Databricks Account`,
				},
				resource.Attribute{
					Name:        "aws_endpoint_service_id",
					Description: `ID of Databricks VPC endpoint service to connect to. Please contact your Databricks representative to request mapping`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of AWS VPC ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpc_endpoint_id",
					Description: `Canonical unique identifier of VPC Endpoint in Databricks Account`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of VPC Endpoint`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_endpoint_id",
					Description: `Canonical unique identifier of VPC Endpoint in Databricks Account`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of VPC Endpoint`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_mws_workspaces",
			Category:         "AWS",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"aws",
				"mws",
				"workspaces",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) ` + "`" + `network_id` + "`" + ` from [networks](mws_networks.md)`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/).`,
				},
				resource.Attribute{
					Name:        "credentials_id",
					Description: `` + "`" + `credentials_id` + "`" + ` from [credentials](mws_credentials.md)`,
				},
				resource.Attribute{
					Name:        "customer_managed_key_id",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "managed_services_customer_managed_key_id",
					Description: `(Optional) ` + "`" + `customer_managed_key_id` + "`" + ` from [customer managed keys](mws_customer_managed_keys.md) with ` + "`" + `use_cases` + "`" + ` set to ` + "`" + `MANAGED_SERVICES` + "`" + `. This is used to encrypt the workspace's notebook and secret data in the control plane.`,
				},
				resource.Attribute{
					Name:        "storage_customer_managed_key_id",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "deployment_name",
					Description: `part of URL: ` + "`" + `https://<deployment-name>.cloud.databricks.com` + "`" + ``,
				},
				resource.Attribute{
					Name:        "workspace_name",
					Description: `name of the workspace, will appear on UI`,
				},
				resource.Attribute{
					Name:        "aws_region",
					Description: `AWS region of VPC`,
				},
				resource.Attribute{
					Name:        "storage_configuration_id",
					Description: `` + "`" + `storage_configuration_id` + "`" + ` from [storage configuration](mws_storage_configurations.md)`,
				},
				resource.Attribute{
					Name:        "private_access_settings_id",
					Description: `(Optional) Canonical unique identifier of [databricks_mws_private_access_settings](mws_private_access_settings.md) in Databricks Account ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the workspace.`,
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
					Description: `(String) URL of the workspace ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify ` + "`" + `create` + "`" + `, ` + "`" + `read` + "`" + ` and ` + "`" + `update` + "`" + ` timeouts. It usually takes 5-7 minutes to provision Databricks E2 Workspace and another couple of minutes for your local DNS caches to resolve. Please launch ` + "`" + `TF_LOG=DEBUG terraform apply` + "`" + ` whenever you observe timeout issues. ` + "`" + `` + "`" + `` + "`" + `hcl timeouts { create = "30m" read = "10m" update = "20m } ` + "`" + `` + "`" + `` + "`" + ` You can reset local DNS caches before provisioning new workspaces with one of the following commands:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the workspace.`,
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
					Description: `(String) URL of the workspace ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify ` + "`" + `create` + "`" + `, ` + "`" + `read` + "`" + ` and ` + "`" + `update` + "`" + ` timeouts. It usually takes 5-7 minutes to provision Databricks E2 Workspace and another couple of minutes for your local DNS caches to resolve. Please launch ` + "`" + `TF_LOG=DEBUG terraform apply` + "`" + ` whenever you observe timeout issues. ` + "`" + `` + "`" + `` + "`" + `hcl timeouts { create = "30m" read = "10m" update = "20m } ` + "`" + `` + "`" + `` + "`" + ` You can reset local DNS caches before provisioning new workspaces with one of the following commands:`,
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
			Type:             "databricks_permissions",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `[cluster](cluster.md) id`,
				},
				resource.Attribute{
					Name:        "job_id",
					Description: `[job](job.md) id`,
				},
				resource.Attribute{
					Name:        "directory_id",
					Description: `[directory](notebook.md) id`,
				},
				resource.Attribute{
					Name:        "directory_path",
					Description: `path of directory`,
				},
				resource.Attribute{
					Name:        "notebook_id",
					Description: `ID of [notebook](notebook.md) within workspace`,
				},
				resource.Attribute{
					Name:        "notebook_path",
					Description: `path of notebook`,
				},
				resource.Attribute{
					Name:        "cluster_policy_id",
					Description: `[cluster policy](cluster_policy.md) id`,
				},
				resource.Attribute{
					Name:        "instance_pool_id",
					Description: `[instance pool](instance_pool.md) id`,
				},
				resource.Attribute{
					Name:        "authorization",
					Description: `either [` + "`" + `tokens` + "`" + `](https://docs.databricks.com/administration-guide/access-control/tokens.html) or [` + "`" + `passwords` + "`" + `](https://docs.databricks.com/administration-guide/users-groups/single-sign-on/index.html#configure-password-permission). One or more ` + "`" + `access_control` + "`" + ` blocks are required to actually set the permission levels: ` + "`" + `` + "`" + `` + "`" + `hcl access_control { group_name = databricks_group.datascience.display_name permission_level = "CAN_USE" } ` + "`" + `` + "`" + `` + "`" + ` Attributes are:`,
				},
				resource.Attribute{
					Name:        "permission_level",
					Description: `(Required) permission level according to specific resource. See examples above for the reference.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional) name of the [user](user.md), which should be used if group name is not used`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Optional) name of the [group](group.md), which should be used if the user name is not used. We recommend setting permissions on groups. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the permissions.`,
				},
				resource.Attribute{
					Name:        "object_type",
					Description: `type of permissions. ## Import The resource permissions can be imported using the object id ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_permissions.this /<object type>/<object id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the permissions.`,
				},
				resource.Attribute{
					Name:        "object_type",
					Description: `type of permissions. ## Import The resource permissions can be imported using the object id ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_permissions.this /<object type>/<object id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
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
					Description: `(Integer) time secret was updated ## Import The resource secret can be imported using ` + "`" + `scopeName|||secretKey` + "`" + ` combination.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the secret.`,
				},
				resource.Attribute{
					Name:        "last_updated_timestamp",
					Description: `(Integer) time secret was updated ## Import The resource secret can be imported using ` + "`" + `scopeName|||secretKey` + "`" + ` combination.`,
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
					Description: `(Required) ` + "`" + `READ` + "`" + `, ` + "`" + `WRITE` + "`" + ` or ` + "`" + `MANAGE` + "`" + `. ## Import The resource secret acl can be imported using ` + "`" + `scopeName|||principalName` + "`" + ` combination.`,
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
					Description: `(Optional) The principal with the only possible value ` + "`" + `users` + "`" + ` that is initially granted ` + "`" + `MANAGE` + "`" + ` permission to the created scope. If it's omitted, then the [databricks_secret_acl](secret_acl.md) with ` + "`" + `MANAGE` + "`" + ` permission applied to the scope is assigned to the API request issuer's user identity (see [documentation](https://docs.databricks.com/dev-tools/api/latest/secrets.html#create-secret-scope)). This part of the state cannot be imported. ## keyvault_metadata On Azure it's possible to create and manage secrets in Azure Key Vault and have use Azure Databricks secret redaction & access control functionality for reading them. There has to be a single Key Vault per single secret scope. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id for the secret scope object.`,
				},
				resource.Attribute{
					Name:        "backend_type",
					Description: `Either ` + "`" + `DATABRICKS` + "`" + ` or ` + "`" + `AZURE_KEYVAULT` + "`" + ` ## Import The secret resource scope can be imported using the scope name. ` + "`" + `initial_manage_principal` + "`" + ` state won't be imported, because the underlying API doesn't include it in the response. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_secret_scope.object <scopeName> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id for the secret scope object.`,
				},
				resource.Attribute{
					Name:        "backend_type",
					Description: `Either ` + "`" + `DATABRICKS` + "`" + ` or ` + "`" + `AZURE_KEYVAULT` + "`" + ` ## Import The secret resource scope can be imported using the scope name. ` + "`" + `initial_manage_principal` + "`" + ` state won't be imported, because the underlying API doesn't include it in the response. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_secret_scope.object <scopeName> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) This is the application id of the given service principal and will be their form of access and identity.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) This is an alias for the service principal can be the full name of the service principal.`,
				},
				resource.Attribute{
					Name:        "allow_cluster_create",
					Description: `(Optional) Allow the service principal to have [cluster](cluster.md) create priviliges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Cluster-usage) and ` + "`" + `cluster_id` + "`" + ` argument. Everyone without ` + "`" + `allow_cluster_create` + "`" + ` arugment set, but with [permission to use](permissions.md#Cluster-Policy-usage) Cluster Policy would be able to create clusters, but within boundaries of that specific policy.`,
				},
				resource.Attribute{
					Name:        "allow_instance_pool_create",
					Description: `(Optional) Allow the service principal to have [instance pool](instance_pool.md) create priviliges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Instance-Pool-usage) and [instance_pool_id](permissions.md#instance_pool_id) argument.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Either service principal is active or not. True by default, but can be set to false in case of service principal deactivation with preserving service principal assets. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the service principal. ## Import The resource scim service principal can be imported using id: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_service_principal.me <service-principal-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the service principal. ## Import The resource scim service principal can be imported using id: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_service_principal.me <service-principal-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_dashboard",
			Category:         "SQL Analytics",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"sql",
				"analytics",
				"dashboard",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_endpoint",
			Category:         "SQL Analytics",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"sql",
				"analytics",
				"endpoint",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the SQL endpoint. Must be unique.`,
				},
				resource.Attribute{
					Name:        "cluster_size",
					Description: `(Required) The size of the clusters allocated to the endpoint: "2X-Small", "X-Small", "Small", "Medium", "Large", "X-Large", "2X-Large", "3X-Large", "4X-Large".`,
				},
				resource.Attribute{
					Name:        "min_num_clusters",
					Description: `Minimum number of clusters available when a SQL endpoint is running. The default is 1.`,
				},
				resource.Attribute{
					Name:        "max_num_clusters",
					Description: `Maximum number of clusters available when a SQL endpoint is running. This field is required. If multi-cluster load balancing is not enabled, this is default to 1.`,
				},
				resource.Attribute{
					Name:        "auto_stop_mins",
					Description: `Time in minutes until an idle SQL endpoint terminates all clusters and stops. This field is optional. The default is 0, which means auto stop is disabled.`,
				},
				resource.Attribute{
					Name:        "instance_profile_arn",
					Description: `[databricks_instance_profile](instance_profile.md) used to access storage from the SQL endpoint. This field is optional.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Databricks tags all endpoint resources with these tags.`,
				},
				resource.Attribute{
					Name:        "spot_instance_policy",
					Description: `The spot policy to use for allocating instances to clusters: ` + "`" + `COST_OPTIMIZED` + "`" + ` or ` + "`" + `RELIABILITY_OPTIMIZED` + "`" + `. This field is optional.`,
				},
				resource.Attribute{
					Name:        "enable_photon",
					Description: `Whether to enable [Photon](https://databricks.com/product/delta-engine). This field is optional. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "jdbc_url",
					Description: `JDBC connection string.`,
				},
				resource.Attribute{
					Name:        "odbc_params",
					Description: `ODBC connection params: ` + "`" + `odbc_params.host` + "`" + `, ` + "`" + `odbc_params.path` + "`" + `, ` + "`" + `odbc_params.protocol` + "`" + `, and ` + "`" + `odbc_params.port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "data_source_id",
					Description: `ID of the data source for this endpoint. This is used to bind an SQLA query to an endpoint. ## Access Control`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "jdbc_url",
					Description: `JDBC connection string.`,
				},
				resource.Attribute{
					Name:        "odbc_params",
					Description: `ODBC connection params: ` + "`" + `odbc_params.host` + "`" + `, ` + "`" + `odbc_params.path` + "`" + `, ` + "`" + `odbc_params.protocol` + "`" + `, and ` + "`" + `odbc_params.port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "data_source_id",
					Description: `ID of the data source for this endpoint. This is used to bind an SQLA query to an endpoint. ## Access Control`,
				},
			},
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
					Description: `` + "`" + `display_name` + "`" + ` of [databricks_group](group.md#display_name) or [databricks_user](user.md#display_name).`,
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
					Description: `anonymous function. ` + "`" + `/` + "`" + ` suffix is mandatory. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_sql_permissions.foo /<object-type>/<object-name> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_query",
			Category:         "SQL Analytics",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"sql",
				"analytics",
				"query",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_visualization",
			Category:         "SQL Analytics",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"sql",
				"analytics",
				"visualization",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_widget",
			Category:         "SQL Analytics",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"sql",
				"analytics",
				"widget",
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
					Name:        "allow_cluster_create",
					Description: `(Optional) Allow the user to have [cluster](cluster.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Cluster-usage) and ` + "`" + `cluster_id` + "`" + ` argument. Everyone without ` + "`" + `allow_cluster_create` + "`" + ` argument set, but with [permission to use](permissions.md#Cluster-Policy-usage) Cluster Policy would be able to create clusters, but within boundaries of that specific policy.`,
				},
				resource.Attribute{
					Name:        "allow_instance_pool_create",
					Description: `(Optional) Allow the user to have [instance pool](instance_pool.md) create privileges. Defaults to false. More fine grained permissions could be assigned with [databricks_permissions](permissions.md#Instance-Pool-usage) and [instance_pool_id](permissions.md#instance_pool_id) argument.`,
				},
				resource.Attribute{
					Name:        "allow_sql_analytics_access",
					Description: `(Optional) This is a field to allow the group to have access to [SQL Analytics](https://databricks.com/product/sql-analytics) feature through [databricks_sql_endpoint](sql_endpoint.md).`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Either user is active or not. True by default, but can be set to false in case of user deactivation with preserving user assets. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the user. ## Import The resource scim user can be imported using id: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_user.me <user-id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Canonical unique identifier for the user. ## Import The resource scim user can be imported using id: ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import databricks_user.me <user-id> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) Key-value map of strings, that represent workspace configuration. Upon resource deletion, properties that start with ` + "`" + `enable` + "`" + ` or ` + "`" + `enforce` + "`" + ` will be reset to ` + "`" + `false` + "`" + ` value, regardless of initial default one. ## Import This resource doesn't support import.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"databricks_aws_s3_mount":                0,
		"databricks_azure_adls_gen1_mount":       1,
		"databricks_azure_adls_gen2_mount":       2,
		"databricks_azure_blob_mount":            3,
		"databricks_cluster":                     4,
		"databricks_cluster_policy":              5,
		"databricks_dbfs_file":                   6,
		"databricks_global_init_script":          7,
		"databricks_group":                       8,
		"databricks_group_instance_profile":      9,
		"databricks_group_member":                10,
		"databricks_instance_pool":               11,
		"databricks_instance_profile":            12,
		"databricks_ip_access_list":              13,
		"databricks_job":                         14,
		"databricks_mws_credentials":             15,
		"databricks_mws_customer_managed_keys":   16,
		"databricks_mws_log_delivery":            17,
		"databricks_mws_networks":                18,
		"databricks_mws_private_access_settings": 19,
		"databricks_mws_storage_configurations":  20,
		"databricks_mws_vpc_endpoint":            21,
		"databricks_mws_workspaces":              22,
		"databricks_notebook":                    23,
		"databricks_permissions":                 24,
		"databricks_secret":                      25,
		"databricks_secret_acl":                  26,
		"databricks_secret_scope":                27,
		"databricks_service_principal":           28,
		"databricks_sql_dashboard":               29,
		"databricks_sql_endpoint":                30,
		"databricks_sql_permissions":             31,
		"databricks_sql_query":                   32,
		"databricks_sql_visualization":           33,
		"databricks_sql_widget":                  34,
		"databricks_token":                       35,
		"databricks_user":                        36,
		"databricks_user_instance_profile":       37,
		"databricks_workspace_conf":              38,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
