package databricks

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "databricks_aws_assume_role_policy",
			Category:         "Deployment",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"aws",
				"assume",
				"role",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `AWS IAM Policy JSON document ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `AWS IAM Policy JSON document ## Related Resources The following resources are used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_aws_bucket_policy",
			Category:         "Deployment",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"aws",
				"bucket",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) AWS S3 Bucket name for which to generate the policy document.`,
				},
				resource.Attribute{
					Name:        "full_access_role",
					Description: `(Optional) Data access role that can have full access for this bucket`,
				},
				resource.Attribute{
					Name:        "databricks_e2_account_id",
					Description: `(Optional) Your Databricks E2 account ID. Used to generate restrictive IAM policies that will increase the security of your root bucket ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `(Read-only) AWS IAM Policy JSON document to grant Databricks full access to bucket. ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `(Read-only) AWS IAM Policy JSON document to grant Databricks full access to bucket. ## Related Resources The following resources are used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_aws_crossaccount_policy",
			Category:         "Deployment",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"aws",
				"crossaccount",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `AWS IAM Policy JSON document ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `AWS IAM Policy JSON document ## Related Resources The following resources are used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_catalogs",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"catalogs",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `set of [databricks_catalog](../resources/catalog.md) names ## Related Resources The following resources are used in the same context:`,
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
					Name:        "cluster_id",
					Description: `(Required if ` + "`" + `cluster_name` + "`" + ` isn't specified) The id of the cluster`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required if ` + "`" + `cluster_id` + "`" + ` isn't specified) The exact name of the cluster to search ## Attribute Reference This data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `cluster ID`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Cluster name, which doesn’t have to be unique.`,
				},
				resource.Attribute{
					Name:        "spark_version",
					Description: `[Runtime version](https://docs.databricks.com/runtime/index.html) of the cluster.`,
				},
				resource.Attribute{
					Name:        "runtime_engine",
					Description: `The type of runtime of the cluster`,
				},
				resource.Attribute{
					Name:        "driver_node_type_id",
					Description: `The node type of the Spark driver.`,
				},
				resource.Attribute{
					Name:        "node_type_id",
					Description: `Any supported [databricks_node_type](../data-sources/node_type.md) id.`,
				},
				resource.Attribute{
					Name:        "driver_instance_pool_id",
					Description: `similar to ` + "`" + `instance_pool_id` + "`" + `, but for driver node.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Identifier of [Cluster Policy](cluster_policy.md) to validate cluster and preset certain defaults.`,
				},
				resource.Attribute{
					Name:        "autotermination_minutes",
					Description: `Automatically terminate the cluster after being inactive for this time in minutes. If specified, the threshold must be between 10 and 10000 minutes. You can also set this value to 0 to explicitly disable automatic termination.`,
				},
				resource.Attribute{
					Name:        "enable_elastic_disk",
					Description: `Use autoscaling local storage.`,
				},
				resource.Attribute{
					Name:        "enable_local_disk_encryption",
					Description: `Enable local disk encryption.`,
				},
				resource.Attribute{
					Name:        "data_security_mode",
					Description: `Security features of the cluster. Unity Catalog requires ` + "`" + `SINGLE_USER` + "`" + ` or ` + "`" + `USER_ISOLATION` + "`" + ` mode. ` + "`" + `LEGACY_PASSTHROUGH` + "`" + ` for passthrough cluster and ` + "`" + `LEGACY_TABLE_ACL` + "`" + ` for Table ACL cluster. Default to ` + "`" + `NONE` + "`" + `, i.e. no security feature enabled.`,
				},
				resource.Attribute{
					Name:        "single_user_name",
					Description: `The optional user name of the user to assign to an interactive cluster. This field is required when using standard AAD Passthrough for Azure Data Lake Storage (ADLS) with a single-user cluster (i.e., not high-concurrency clusters).`,
				},
				resource.Attribute{
					Name:        "idempotency_token",
					Description: `An optional token to guarantee the idempotency of cluster creation requests.`,
				},
				resource.Attribute{
					Name:        "ssh_public_keys",
					Description: `SSH public key contents that will be added to each Spark node in this cluster.`,
				},
				resource.Attribute{
					Name:        "spark_env_vars",
					Description: `Map with environment variable key-value pairs to fine-tune Spark clusters. Key-value pairs of the form (X,Y) are exported (i.e., X='Y') while launching the driver and workers.`,
				},
				resource.Attribute{
					Name:        "custom_tags",
					Description: `Additional tags for cluster resources.`,
				},
				resource.Attribute{
					Name:        "spark_conf",
					Description: `Map with key-value pairs to fine-tune Spark clusters. ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `cluster ID`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Cluster name, which doesn’t have to be unique.`,
				},
				resource.Attribute{
					Name:        "spark_version",
					Description: `[Runtime version](https://docs.databricks.com/runtime/index.html) of the cluster.`,
				},
				resource.Attribute{
					Name:        "runtime_engine",
					Description: `The type of runtime of the cluster`,
				},
				resource.Attribute{
					Name:        "driver_node_type_id",
					Description: `The node type of the Spark driver.`,
				},
				resource.Attribute{
					Name:        "node_type_id",
					Description: `Any supported [databricks_node_type](../data-sources/node_type.md) id.`,
				},
				resource.Attribute{
					Name:        "driver_instance_pool_id",
					Description: `similar to ` + "`" + `instance_pool_id` + "`" + `, but for driver node.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `Identifier of [Cluster Policy](cluster_policy.md) to validate cluster and preset certain defaults.`,
				},
				resource.Attribute{
					Name:        "autotermination_minutes",
					Description: `Automatically terminate the cluster after being inactive for this time in minutes. If specified, the threshold must be between 10 and 10000 minutes. You can also set this value to 0 to explicitly disable automatic termination.`,
				},
				resource.Attribute{
					Name:        "enable_elastic_disk",
					Description: `Use autoscaling local storage.`,
				},
				resource.Attribute{
					Name:        "enable_local_disk_encryption",
					Description: `Enable local disk encryption.`,
				},
				resource.Attribute{
					Name:        "data_security_mode",
					Description: `Security features of the cluster. Unity Catalog requires ` + "`" + `SINGLE_USER` + "`" + ` or ` + "`" + `USER_ISOLATION` + "`" + ` mode. ` + "`" + `LEGACY_PASSTHROUGH` + "`" + ` for passthrough cluster and ` + "`" + `LEGACY_TABLE_ACL` + "`" + ` for Table ACL cluster. Default to ` + "`" + `NONE` + "`" + `, i.e. no security feature enabled.`,
				},
				resource.Attribute{
					Name:        "single_user_name",
					Description: `The optional user name of the user to assign to an interactive cluster. This field is required when using standard AAD Passthrough for Azure Data Lake Storage (ADLS) with a single-user cluster (i.e., not high-concurrency clusters).`,
				},
				resource.Attribute{
					Name:        "idempotency_token",
					Description: `An optional token to guarantee the idempotency of cluster creation requests.`,
				},
				resource.Attribute{
					Name:        "ssh_public_keys",
					Description: `SSH public key contents that will be added to each Spark node in this cluster.`,
				},
				resource.Attribute{
					Name:        "spark_env_vars",
					Description: `Map with environment variable key-value pairs to fine-tune Spark clusters. Key-value pairs of the form (X,Y) are exported (i.e., X='Y') while launching the driver and workers.`,
				},
				resource.Attribute{
					Name:        "custom_tags",
					Description: `Additional tags for cluster resources.`,
				},
				resource.Attribute{
					Name:        "spark_conf",
					Description: `Map with key-value pairs to fine-tune Spark clusters. ## Related Resources The following resources are often used in the same context:`,
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_clusters",
			Category:         "Compute",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"compute",
				"clusters",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name_contains",
					Description: `(Optional) Only return [databricks_cluster](../resources/cluster.md#cluster_id) ids that match the given name string. ## Attribute Reference This data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `list of [databricks_cluster](../resources/cluster.md#cluster_id) ids ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `list of [databricks_cluster](../resources/cluster.md#cluster_id) ids ## Related Resources The following resources are used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_current_user",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"current",
				"user",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
					Name:        "path",
					Description: `(Required) Path on DBFS for the file from which to get content.`,
				},
				resource.Attribute{
					Name:        "limit_file_size",
					Description: `(Required - boolean) Do not load content for files larger than 4MB. ## Attribute Reference This data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `base64-encoded file contents`,
				},
				resource.Attribute{
					Name:        "file_size",
					Description: `size of the file in bytes ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "content",
					Description: `base64-encoded file contents`,
				},
				resource.Attribute{
					Name:        "file_size",
					Description: `size of the file in bytes ## Related Resources The following resources are used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_dbfs_file_paths",
			Category:         "Storage",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"storage",
				"dbfs",
				"file",
				"paths",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path on DBFS for the file to perform listing`,
				},
				resource.Attribute{
					Name:        "recursive",
					Description: `(Required) Either or not recursively list all files ## Attribute Reference This data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "path_list",
					Description: `returns list of objects with ` + "`" + `path` + "`" + ` and ` + "`" + `file_size` + "`" + ` attributes in each ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "path_list",
					Description: `returns list of objects with ` + "`" + `path` + "`" + ` and ` + "`" + `file_size` + "`" + ` attributes in each ## Related Resources The following resources are used in the same context:`,
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path to a directory in the workspace ## Attribute Reference This data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `directory object ID`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "object_id",
					Description: `directory object ID`,
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
					Description: `(Required) Display name of the group. The group must exist before this resource can be planned.`,
				},
				resource.Attribute{
					Name:        "recursive",
					Description: `(Optional) Collect information for all nested groups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id for the group object.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `ID of the group in an external identity provider.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Set of [databricks_user](../resources/user.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.`,
				},
				resource.Attribute{
					Name:        "service_principals",
					Description: `Set of [databricks_service_principal](../resources/service_principal.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.`,
				},
				resource.Attribute{
					Name:        "child_groups",
					Description: `Set of [databricks_group](../resources/group.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Set of [group](../resources/group.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.`,
				},
				resource.Attribute{
					Name:        "instance_profiles",
					Description: `Set of [instance profile](../resources/instance_profile.md) ARNs, that can be modified by [databricks_group_instance_profile](../resources/group_instance_profile.md) resource.`,
				},
				resource.Attribute{
					Name:        "allow_cluster_create",
					Description: `True if group members can create [clusters](../resources/cluster.md)`,
				},
				resource.Attribute{
					Name:        "allow_instance_pool_create",
					Description: `True if group members can create [instance pools](../resources/instance_pool.md) ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id for the group object.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `ID of the group in an external identity provider.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Set of [databricks_user](../resources/user.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.`,
				},
				resource.Attribute{
					Name:        "service_principals",
					Description: `Set of [databricks_service_principal](../resources/service_principal.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.`,
				},
				resource.Attribute{
					Name:        "child_groups",
					Description: `Set of [databricks_group](../resources/group.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Set of [group](../resources/group.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.`,
				},
				resource.Attribute{
					Name:        "instance_profiles",
					Description: `Set of [instance profile](../resources/instance_profile.md) ARNs, that can be modified by [databricks_group_instance_profile](../resources/group_instance_profile.md) resource.`,
				},
				resource.Attribute{
					Name:        "allow_cluster_create",
					Description: `True if group members can create [clusters](../resources/cluster.md)`,
				},
				resource.Attribute{
					Name:        "allow_instance_pool_create",
					Description: `True if group members can create [instance pools](../resources/instance_pool.md) ## Related Resources The following resources are used in the same context:`,
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
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
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `the id of [databricks_job](../resources/job.md) if the resource was matched by name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the job name of [databricks_job](../resources/job.md) if the resource was matched by id.`,
				},
				resource.Attribute{
					Name:        "job_settings",
					Description: `the same fields as in [databricks_job](../resources/job.md). ## Related Resources The following resources are used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_jobs",
			Category:         "Compute",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"compute",
				"jobs",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `map of [databricks_job](../resources/job.md) names to ids ## Related Resources The following resources are used in the same context:`,
				},
			},
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
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `name-to-id map for all of the credentials in the account ## Related Resources The following resources are used in the same context:`,
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
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `name-to-id map for all of the workspaces in the account ## Related Resources The following resources are used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_node_type",
			Category:         "Compute",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"compute",
				"node",
				"type",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "min_memory_gb",
					Description: `(Optional) Minimum amount of memory per node in gigabytes. Defaults to`,
				},
				resource.Attribute{
					Name:        "gb_per_core",
					Description: `(Optional) Number of gigabytes per core available on instance. Conflicts with ` + "`" + `min_memory_gb` + "`" + `. Defaults to`,
				},
				resource.Attribute{
					Name:        "min_cores",
					Description: `(Optional) Minimum number of CPU cores available on instance. Defaults to`,
				},
				resource.Attribute{
					Name:        "min_gpus",
					Description: `(Optional) Minimum number of GPU's attached to instance. Defaults to`,
				},
				resource.Attribute{
					Name:        "local_disk",
					Description: `(Optional) Pick only nodes with local storage. Defaults to`,
				},
				resource.Attribute{
					Name:        "local_disk_min_size",
					Description: `(Optional) Pick only nodes that have size local storage greater or equal to given value. Defaults to`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional, case insensitive string) Node category, which can be one of (depending on the cloud environment, could be checked with ` + "`" + `databricks clusters list-node-types|jq '.node_types[]|.category'|sort |uniq` + "`" + `):`,
				},
				resource.Attribute{
					Name:        "photon_worker_capable",
					Description: `(Optional) Pick only nodes that can run Photon workers. Defaults to`,
				},
				resource.Attribute{
					Name:        "photon_driver_capable",
					Description: `(Optional) Pick only nodes that can run Photon driver. Defaults to`,
				},
				resource.Attribute{
					Name:        "graviton",
					Description: `(boolean, optional) if we should limit the search only to nodes with AWS Graviton CPUs. Default to`,
				},
				resource.Attribute{
					Name:        "fleet",
					Description: `(boolean, optional) if we should limit the search only to [AWS fleet instance types](https://docs.databricks.com/compute/aws-fleet-instances.html). Default to`,
				},
				resource.Attribute{
					Name:        "is_io_cache_enabled",
					Description: `(Optional) . Pick only nodes that have IO Cache. Defaults to`,
				},
				resource.Attribute{
					Name:        "support_port_forwarding",
					Description: `(Optional) Pick only nodes that support port forwarding. Defaults to`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `node type, that can be used for [databricks_job](../resources/job.md), [databricks_cluster](../resources/cluster.md), or [databricks_instance_pool](../resources/instance_pool.md). ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `node type, that can be used for [databricks_job](../resources/job.md), [databricks_cluster](../resources/cluster.md), or [databricks_instance_pool](../resources/instance_pool.md). ## Related Resources The following resources are used in the same context:`,
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
					Description: `(Required) Notebook path on the workspace`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Required) Notebook format to export. Either ` + "`" + `SOURCE` + "`" + `, ` + "`" + `HTML` + "`" + `, ` + "`" + `JUPYTER` + "`" + `, or ` + "`" + `DBC` + "`" + `. ## Attribute Reference This data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `notebook content in selected format`,
				},
				resource.Attribute{
					Name:        "language",
					Description: `notebook language`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `notebook object ID`,
				},
				resource.Attribute{
					Name:        "object_type",
					Description: `notebook object type`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "content",
					Description: `notebook content in selected format`,
				},
				resource.Attribute{
					Name:        "language",
					Description: `notebook language`,
				},
				resource.Attribute{
					Name:        "object_id",
					Description: `notebook object ID`,
				},
				resource.Attribute{
					Name:        "object_type",
					Description: `notebook object type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_notebook_paths",
			Category:         "Workspace",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"workspace",
				"notebook",
				"paths",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path to workspace directory`,
				},
				resource.Attribute{
					Name:        "recursive",
					Description: `(Required) Either or recursively walk given path ## Attribute Reference This data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "notebook_path_list",
					Description: `list of objects with ` + "`" + `path` + "`" + ` and ` + "`" + `language` + "`" + ` attributes`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "notebook_path_list",
					Description: `list of objects with ` + "`" + `path` + "`" + ` and ` + "`" + `language` + "`" + ` attributes`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_pipelines",
			Category:         "Compute",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"compute",
				"pipelines",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pipeline_name",
					Description: `(Optional) Filter Delta Live Tables pipelines by name for a given search term. ` + "`" + `%` + "`" + ` is the supported wildcard operator. ## Attribute Reference This data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `List of ids for [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html) pipelines matching the provided search criteria. ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `List of ids for [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html) pipelines matching the provided search criteria. ## Related Resources The following resources are used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_schemas",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"schemas",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "catalog_name",
					Description: `(Required) Name of [databricks_catalog](../resources/catalog.md) ## Attribute Reference This data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `set of [databricks_schema](../resources/schema.md) full names:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `set of [databricks_schema](../resources/schema.md) full names:`,
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_service_principals",
			Category:         "Security",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"security",
				"service",
				"principals",
			},
			Arguments:  []resource.Attribute{},
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
					Description: `(Required) The name of the share ## Attribute Reference This data source exports the following attributes:`,
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
					Name:        "object",
					Description: `arrays containing details of each object in the share.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Full name of the object being shared.`,
				},
				resource.Attribute{
					Name:        "data_object_type",
					Description: `Type of the object.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Description about the object. ## Related Resources The following resources are used in the same context:`,
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
					Name:        "object",
					Description: `arrays containing details of each object in the share.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Full name of the object being shared.`,
				},
				resource.Attribute{
					Name:        "data_object_type",
					Description: `Type of the object.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Description about the object. ## Related Resources The following resources are used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_shares",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"shares",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "shares",
					Description: `list of [databricks_share](../resources/share.md) names. ## Related Resources The following resources are used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_spark_version",
			Category:         "Compute",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"compute",
				"spark",
				"version",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "latest",
					Description: `(boolean, optional) if we should return only the latest version if there is more than one result. Default to ` + "`" + `true` + "`" + `. If set to ` + "`" + `false` + "`" + ` and multiple versions are matching, throws an error.`,
				},
				resource.Attribute{
					Name:        "long_term_support",
					Description: `(boolean, optional) if we should limit the search only to LTS (long term support) & ESR (extended support) versions. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ml",
					Description: `(boolean, optional) if we should limit the search only to ML runtimes. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "genomics",
					Description: `(boolean, optional) if we should limit the search only to Genomics (HLS) runtimes. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gpu",
					Description: `(boolean, optional) if we should limit the search only to runtimes that support GPUs. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "photon",
					Description: `(boolean, optional) if we should limit the search only to Photon runtimes. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "graviton",
					Description: `(boolean, optional) if we should limit the search only to runtimes supporting AWS Graviton CPUs. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "beta",
					Description: `(boolean, optional) if we should limit the search only to runtimes that are in Beta stage. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scala",
					Description: `(string, optional) if we should limit the search only to runtimes that are based on specific Scala version. Default to ` + "`" + `2.12` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spark_version",
					Description: `(string, optional) if we should limit the search only to runtimes that are based on specific Spark version. Default to empty string. It could be specified as ` + "`" + `3` + "`" + `, or ` + "`" + `3.0` + "`" + `, or full version, like, ` + "`" + `3.0.1` + "`" + `. ## Attribute Reference Data source exposes the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Databricks Runtime version, that can be used as ` + "`" + `spark_version` + "`" + ` field in [databricks_job](../resources/job.md), [databricks_cluster](../resources/cluster.md), or [databricks_instance_pool](../resources/instance_pool.md). ## Related Resources The following resources are used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Databricks Runtime version, that can be used as ` + "`" + `spark_version` + "`" + ` field in [databricks_job](../resources/job.md), [databricks_cluster](../resources/cluster.md), or [databricks_instance_pool](../resources/instance_pool.md). ## Related Resources The following resources are used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_warehouse",
			Category:         "Databricks SQL",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"sql",
				"warehouse",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_sql_warehouses",
			Category:         "Databricks SQL",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"sql",
				"warehouses",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "warehouse_name_contains",
					Description: `(Optional) Only return [databricks_sql_endpoint](../resources/sql_endpoint.md#id) ids that match the given name string. ## Attribute Reference This data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `list of [databricks_sql_endpoint](../resources/sql_endpoint.md#id) ids ## Related Resources The following resources are often used in the same context:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `list of [databricks_sql_endpoint](../resources/sql_endpoint.md#id) ids ## Related Resources The following resources are often used in the same context:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_tables",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"tables",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "catalog_name",
					Description: `(Required) Name of [databricks_catalog](../resources/catalog.md)`,
				},
				resource.Attribute{
					Name:        "schema_name",
					Description: `(Required) Name of [databricks_schema](../resources/schema.md) ## Attribute Reference This data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `set of databricks_table full names:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `set of databricks_table full names:`,
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
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_views",
			Category:         "Unity Catalog",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"unity",
				"catalog",
				"views",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "catalog_name",
					Description: `(Required) Name of [databricks_catalog](../resources/catalog.md)`,
				},
				resource.Attribute{
					Name:        "schema_name",
					Description: `(Required) Name of [databricks_schema](../resources/schema.md) ## Attribute Reference This data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `set of databricks_view full names:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `set of databricks_view full names:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_zones",
			Category:         "Deployment",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deployment",
				"zones",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id for the zone object.`,
				},
				resource.Attribute{
					Name:        "default_zone",
					Description: `This is the default zone that gets assigned to your workspace. This is the zone used by default for clusters and instance pools.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `This is a list of all the zones available for your subnets in your Databricks workspace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id for the zone object.`,
				},
				resource.Attribute{
					Name:        "default_zone",
					Description: `This is the default zone that gets assigned to your workspace. This is the zone used by default for clusters and instance pools.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `This is a list of all the zones available for your subnets in your Databricks workspace.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"databricks_aws_assume_role_policy":  0,
		"databricks_aws_bucket_policy":       1,
		"databricks_aws_crossaccount_policy": 2,
		"databricks_catalogs":                3,
		"databricks_cluster":                 4,
		"databricks_cluster_policy":          5,
		"databricks_clusters":                6,
		"databricks_current_user":            7,
		"databricks_dbfs_file":               8,
		"databricks_dbfs_file_paths":         9,
		"databricks_directory":               10,
		"databricks_group":                   11,
		"databricks_instance_pool":           12,
		"databricks_job":                     13,
		"databricks_jobs":                    14,
		"databricks_mws_credentials":         15,
		"databricks_mws_workspaces":          16,
		"databricks_node_type":               17,
		"databricks_notebook":                18,
		"databricks_notebook_paths":          19,
		"databricks_pipelines":               20,
		"databricks_schemas":                 21,
		"databricks_service_principal":       22,
		"databricks_service_principals":      23,
		"databricks_share":                   24,
		"databricks_shares":                  25,
		"databricks_spark_version":           26,
		"databricks_sql_warehouse":           27,
		"databricks_sql_warehouses":          28,
		"databricks_tables":                  29,
		"databricks_user":                    30,
		"databricks_views":                   31,
		"databricks_zones":                   32,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
