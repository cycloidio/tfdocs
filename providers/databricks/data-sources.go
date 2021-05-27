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
			Category:         "AWS",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"aws",
				"assume",
				"role",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `AWS IAM Policy JSON document`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `AWS IAM Policy JSON document`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_aws_bucket_policy",
			Category:         "AWS",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
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
					Description: `(Optional) Data access role that can have full access for this bucket ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `(Read-only) AWS IAM Policy JSON document to grant Databricks full access to bucket.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `(Read-only) AWS IAM Policy JSON document to grant Databricks full access to bucket.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "databricks_aws_crossaccount_policy",
			Category:         "AWS",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"aws",
				"crossaccount",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `AWS IAM Policy JSON document`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "json",
					Description: `AWS IAM Policy JSON document`,
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
					Description: `(Required) Path on DBFS for the file to get content of`,
				},
				resource.Attribute{
					Name:        "limit_file_size",
					Description: `(Required) Do lot load content for files smaller than this in bytes ## Attribute Reference This data source exports the following attributes:`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `base64-encoded file contents`,
				},
				resource.Attribute{
					Name:        "file_size",
					Description: `size of the file in bytes`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "content",
					Description: `base64-encoded file contents`,
				},
				resource.Attribute{
					Name:        "file_size",
					Description: `size of the file in bytes`,
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
					Description: `returns list of objects with ` + "`" + `path` + "`" + ` and ` + "`" + `file_size` + "`" + ` attributes in each`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "path_list",
					Description: `returns list of objects with ` + "`" + `path` + "`" + ` and ` + "`" + `file_size` + "`" + ` attributes in each`,
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
					Name:        "members",
					Description: `Set of [user](../resources/user.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.`,
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
					Description: `True if group members can create [instance pools](../resources/instance_pool.md)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id for the group object.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `Set of [user](../resources/user.md) identifiers, that can be modified with [databricks_group_member](../resources/group_member.md) resource.`,
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
					Description: `True if group members can create [instance pools](../resources/instance_pool.md)`,
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
					Name:        "category",
					Description: `(Optional) Node category, which can be one of:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `node type, that can be used for [databricks_job](../resources/job.md), [databricks_cluster](../resources/cluster.md), or [databricks_instance_pool](../resources/instance_pool.md).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `node type, that can be used for [databricks_job](../resources/job.md), [databricks_cluster](../resources/cluster.md), or [databricks_instance_pool](../resources/instance_pool.md).`,
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
					Description: `(boolean, optional) if we should return only the latest version if there is more than one result. Default to ` + "`" + `true` + "`" + `. If set to ` + "`" + `false` + "`" + ` and multiple versions are matching, throws an error`,
				},
				resource.Attribute{
					Name:        "long_term_support",
					Description: `(boolean, optional) if we should limit the search only to LTS (long term support) versions. Default to ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ml",
					Description: `(boolean, optional) if we should limit the search only to ML runtimes. Default to ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "genomics",
					Description: `(boolean, optional) if we should limit the search only to Genomics (HLS) runtimes. Default to ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "gpu",
					Description: `(boolean, optional) if we should limit the search only to runtimes that support GPUs. Default to ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "beta",
					Description: `(boolean, optional) if we should limit the search only to runtimes that are in Beta stage. Default to ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "scala",
					Description: `(string, optional) if we should limit the search only to runtimes that are based on specific Scala version. Default to ` + "`" + `2.12` + "`" + ``,
				},
				resource.Attribute{
					Name:        "spark_version",
					Description: `(string, optional) if we should limit the search only to runtimes that are based on specific Spark version. Default to empty string. It could be specified as ` + "`" + `3` + "`" + `, or ` + "`" + `3.0` + "`" + `, or full version, like, ` + "`" + `3.0.1` + "`" + ` ## Attribute Reference Data source exposes the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Databricks Runtime version, that can be used as ` + "`" + `spark_version` + "`" + ` field in [databricks_job](../resources/job.md), [databricks_cluster](../resources/cluster.md), or [databricks_instance_pool](../resources/instance_pool.md).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Databricks Runtime version, that can be used as ` + "`" + `spark_version` + "`" + ` field in [databricks_job](../resources/job.md), [databricks_cluster](../resources/cluster.md), or [databricks_instance_pool](../resources/instance_pool.md).`,
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
			Type:             "databricks_zones",
			Category:         "AWS",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"aws",
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
		"databricks_current_user":            3,
		"databricks_dbfs_file":               4,
		"databricks_dbfs_file_paths":         5,
		"databricks_group":                   6,
		"databricks_node_type":               7,
		"databricks_notebook":                8,
		"databricks_notebook_paths":          9,
		"databricks_spark_version":           10,
		"databricks_user":                    11,
		"databricks_zones":                   12,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
