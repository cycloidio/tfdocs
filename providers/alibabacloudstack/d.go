package alibabacloudstack

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_adb_db_clusters",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Adb DBClusters to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) The description of DBCluster.`,
				},
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional, ForceNew) A regex string to filter results by DBCluster description.`,
				},
				resource.Attribute{
					Name:        "enable_details",
					Description: `(Optional) Default to ` + "`" + `false` + "`" + `. Set it to ` + "`" + `true` + "`" + ` can output more details about resource attributes.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, ForceNew, Computed) A list of DBCluster IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew) The ID of the resource group.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, ForceNew) The status of the resource. ## Argument Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of DBCluster descriptions.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of Adb Db Clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "auto_renew_period",
					Description: `Auto-renewal period of an cluster, in the unit of the month.`,
				},
				resource.Attribute{
					Name:        "commodity_code",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "compute_resource",
					Description: `The specifications of computing resources in elastic mode. The increase of resources can speed up queries. AnalyticDB for MySQL automatically scales computing resources. For more information, see [Specifications](https://www.alibabacloud.com/help/en/doc-detail/144851.htm).`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `The endpoint of the cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The CreateTime of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "db_cluster_category",
					Description: `The db cluster category.`,
				},
				resource.Attribute{
					Name:        "db_cluster_id",
					Description: `The db cluster id.`,
				},
				resource.Attribute{
					Name:        "db_cluster_network_type",
					Description: `The db cluster network type.`,
				},
				resource.Attribute{
					Name:        "db_cluster_type",
					Description: `The db cluster type.`,
				},
				resource.Attribute{
					Name:        "db_cluster_version",
					Description: `The db cluster version.`,
				},
				resource.Attribute{
					Name:        "db_node_class",
					Description: `The db node class.`,
				},
				resource.Attribute{
					Name:        "db_node_count",
					Description: `The db node count.`,
				},
				resource.Attribute{
					Name:        "db_node_storage",
					Description: `The db node storage.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of DBCluster.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The type of the disk.`,
				},
				resource.Attribute{
					Name:        "dts_job_id",
					Description: `The ID of the data synchronization task in Data Transmission Service (DTS). This parameter is valid only for analytic instances.`,
				},
				resource.Attribute{
					Name:        "elastic_io_resource",
					Description: `The elastic io resource.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The engine of the database.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `The engine version of the database..`,
				},
				resource.Attribute{
					Name:        "executor_count",
					Description: `The number of nodes. The node resources are used for data computing in elastic mode.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The time when the cluster expires.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `Indicates whether the cluster has expired.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the DBCluster.`,
				},
				resource.Attribute{
					Name:        "lock_mode",
					Description: `The lock mode of the cluster.`,
				},
				resource.Attribute{
					Name:        "lock_reason",
					Description: `The reason why the cluster is locked.`,
				},
				resource.Attribute{
					Name:        "maintain_time",
					Description: `The maintenance window of the cluster.`,
				},
				resource.Attribute{
					Name:        "payment_type",
					Description: `The payment type of the resource.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `The payment type of the resource.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port that is used to access the cluster.`,
				},
				resource.Attribute{
					Name:        "rds_instance_id",
					Description: `The ID of the ApsaraDB RDS instance from which data is synchronized to the cluster. This parameter is valid only for analytic instances.`,
				},
				resource.Attribute{
					Name:        "renewal_status",
					Description: `The status of renewal.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of the resource group.`,
				},
				resource.Attribute{
					Name:        "security_ips",
					Description: `List of IP addresses allowed to access all databases of an cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource.`,
				},
				resource.Attribute{
					Name:        "storage_resource",
					Description: `The specifications of storage resources in elastic mode. The resources are used for data read and write operations. The increase of resources can improve the read and write performance of your cluster. For more information, see [Specifications](https://www.alibabacloud.com/help/en/doc-detail/144851.htm).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tag of the resource.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `The key of the tags.`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `The value of the tags.`,
				},
				resource.Attribute{
					Name:        "vpc_cloud_instance_id",
					Description: `The vpc cloud instance id.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The vpc id.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The vswitch id.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The zone ID of the resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_adb_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for ADB that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "multi",
					Description: `(Optional) Indicate whether the zones can be used in a multi AZ configuration. Default to ` + "`" + `false` + "`" + `. Multi AZ is usually used to launch ADB instances.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_api_gateway_apis",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of apis to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_id",
					Description: `(Deprecated, Optional) (It has been deprecated from version 1.52.2, and use field 'ids' to replace.) ID of the specified API.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) ID of the specified group.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter api gateway apis by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.52.2+) A list of api IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of api IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of api names.`,
				},
				resource.Attribute{
					Name:        "apis",
					Description: `A list of apis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `API ID, which is generated by the system and globally unique.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `API name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `API description.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The ID of the region where the API is located.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The group id that the apis belong to.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The group name that the apis belong to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of api IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of api names.`,
				},
				resource.Attribute{
					Name:        "apis",
					Description: `A list of apis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `API ID, which is generated by the system and globally unique.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `API name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `API description.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The ID of the region where the API is located.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The group id that the apis belong to.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The group name that the apis belong to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_api_gateway_apps",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of apps to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter apps by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available in 1.52.2+) A list of app IDs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.55.3+) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of app IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of app names.`,
				},
				resource.Attribute{
					Name:        "apps",
					Description: `A list of apps. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `App ID, which is generated by the system and globally unique.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `App name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `App description.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Creation time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `Last modification time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "app_code",
					Description: `App code.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of app IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of app names.`,
				},
				resource.Attribute{
					Name:        "apps",
					Description: `A list of apps. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `App ID, which is generated by the system and globally unique.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `App name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `App description.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Creation time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `Last modification time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "app_code",
					Description: `App code.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_api_gateway_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of api groups to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter api gateway groups by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.52.1+) A list of api group IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of api group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of api group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of api groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `API group ID, which is generated by the system and globally unique.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `API group name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `API group description.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The ID of the region where the API group is located.`,
				},
				resource.Attribute{
					Name:        "sub_domain",
					Description: `Second-level domain name automatically assigned to the API group.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Creation time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `Last modification time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "traffic_limit",
					Description: `Upper QPS limit of the API group; default value: 500, which can be increased by submitting an application.`,
				},
				resource.Attribute{
					Name:        "billing_status",
					Description: `Billing status. - NORMAL: The API group is normal. - LOCKED: Locked due to outstanding payment.`,
				},
				resource.Attribute{
					Name:        "illegal_status",
					Description: `Locking in invalid state. - NORMAL: The API group is normal. - LOCKED: Locked due to illegality.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of api group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of api group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of api groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `API group ID, which is generated by the system and globally unique.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `API group name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `API group description.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The ID of the region where the API group is located.`,
				},
				resource.Attribute{
					Name:        "sub_domain",
					Description: `Second-level domain name automatically assigned to the API group.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Creation time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `Last modification time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "traffic_limit",
					Description: `Upper QPS limit of the API group; default value: 500, which can be increased by submitting an application.`,
				},
				resource.Attribute{
					Name:        "billing_status",
					Description: `Billing status. - NORMAL: The API group is normal. - LOCKED: Locked due to outstanding payment.`,
				},
				resource.Attribute{
					Name:        "illegal_status",
					Description: `Locking in invalid state. - NORMAL: The API group is normal. - LOCKED: Locked due to illegality.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_api_gateway_service",
			Category:         "Data Sources",
			ShortDescription: `Provides a datasource to open the API gateway service automatically.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Setting the value to ` + "`" + `On` + "`" + ` to enable the service. If has been enabled, return the result. Valid values: "On" or "Off". Default to "Off". ->`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current service enable status.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The current service enable status.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_ecs_instance_families",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ecs instance families to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ecs instance family IDs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Required) Filter the results by specifying the status of ecs instance families.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "families",
					Description: `A list of ecs instance families. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_type_family_id",
					Description: `ID of the ecs instance families.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `generation of ecs instance families.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "families",
					Description: `A list of ecs instance families. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_type_family_id",
					Description: `ID of the ecs instance families.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `generation of ecs instance families.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_environment_services_by_product",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of environment services to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of environment service IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `A list of environment services. Each element contains the following attributes:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `A list of environment services. Each element contains the following attributes:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_instance_families",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of instance families to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance family IDs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Specify Status to filter the resulting instance families by their availability.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Optional) Filter the results by the specified resource type.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "families",
					Description: `A list of instance families. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance families.`,
				},
				resource.Attribute{
					Name:        "series_name",
					Description: `Series name for instance families.`,
				},
				resource.Attribute{
					Name:        "modifier",
					Description: `Modifier name.`,
				},
				resource.Attribute{
					Name:        "series_name",
					Description: `Series name for instance families.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Specified resource type.`,
				},
				resource.Attribute{
					Name:        "is_deleted",
					Description: `Specify the state in "Y" or "N" form.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "families",
					Description: `A list of instance families. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance families.`,
				},
				resource.Attribute{
					Name:        "series_name",
					Description: `Series name for instance families.`,
				},
				resource.Attribute{
					Name:        "modifier",
					Description: `Modifier name.`,
				},
				resource.Attribute{
					Name:        "series_name",
					Description: `Series name for instance families.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Specified resource type.`,
				},
				resource.Attribute{
					Name:        "is_deleted",
					Description: `Specify the state in "Y" or "N" form.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_logon_policies",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Logon Policies.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) The ids of the Logon Policies.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter Logon Policies by name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The Logon Policies description.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) The Rule for the Logon Policies.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) The IP Ranges for the Logon Policies.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Logon Policies.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `The list of the Logon Policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The rule of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The ip range of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The end time of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The start time of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "login_policy_id",
					Description: `The login policy id of the Logon Policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Logon Policies.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `The list of the Logon Policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `The rule of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `The ip range of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The end time of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The start time of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "login_policy_id",
					Description: `The login policy id of the Logon Policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_organizations",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of organizations to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of organizations IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by organization name.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Optional) Filter the results by the specified organization parent ID.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "organizations",
					Description: `A list of organizations. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the organization.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `organization name.`,
				},
				resource.Attribute{
					Name:        "cuser_id",
					Description: `Id of a Cuser.`,
				},
				resource.Attribute{
					Name:        "muser_id",
					Description: `Id of a Muser.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `alias for the Organization.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent id of an Organization.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "organizations",
					Description: `A list of organizations. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the organization.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `organization name.`,
				},
				resource.Attribute{
					Name:        "cuser_id",
					Description: `Id of a Cuser.`,
				},
				resource.Attribute{
					Name:        "muser_id",
					Description: `Id of a Muser.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `alias for the Organization.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `Parent id of an Organization.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_password_policies",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of password policies to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance family IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `A list of password policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the password policies.`,
				},
				resource.Attribute{
					Name:        "hard_expiry",
					Description: `Specifies whether to disable logon after the password expires.`,
				},
				resource.Attribute{
					Name:        "require_numbers",
					Description: `Specifies whether digits are required.`,
				},
				resource.Attribute{
					Name:        "require_symbols",
					Description: `Specifies whether special characters are required.`,
				},
				resource.Attribute{
					Name:        "require_lowercase_characters",
					Description: `Specifies whether lowercase letters are required.`,
				},
				resource.Attribute{
					Name:        "require_uppercase_characters",
					Description: `Specifies whether uppercase letters are required.`,
				},
				resource.Attribute{
					Name:        "max_login_attempts",
					Description: `The maximum number of allowed logon attempts.`,
				},
				resource.Attribute{
					Name:        "max_password_age",
					Description: `The validity period of the password. Unit: days.`,
				},
				resource.Attribute{
					Name:        "minimum_password_length",
					Description: `The minimum length of the password.`,
				},
				resource.Attribute{
					Name:        "password_reuse_prevention",
					Description: `The maximum number of allowed password reuse attempts.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policies",
					Description: `A list of password policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the password policies.`,
				},
				resource.Attribute{
					Name:        "hard_expiry",
					Description: `Specifies whether to disable logon after the password expires.`,
				},
				resource.Attribute{
					Name:        "require_numbers",
					Description: `Specifies whether digits are required.`,
				},
				resource.Attribute{
					Name:        "require_symbols",
					Description: `Specifies whether special characters are required.`,
				},
				resource.Attribute{
					Name:        "require_lowercase_characters",
					Description: `Specifies whether lowercase letters are required.`,
				},
				resource.Attribute{
					Name:        "require_uppercase_characters",
					Description: `Specifies whether uppercase letters are required.`,
				},
				resource.Attribute{
					Name:        "max_login_attempts",
					Description: `The maximum number of allowed logon attempts.`,
				},
				resource.Attribute{
					Name:        "max_password_age",
					Description: `The validity period of the password. Unit: days.`,
				},
				resource.Attribute{
					Name:        "minimum_password_length",
					Description: `The minimum length of the password.`,
				},
				resource.Attribute{
					Name:        "password_reuse_prevention",
					Description: `The maximum number of allowed password reuse attempts.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_quotas",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of quota to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "product_name",
					Description: `(Required) The name of the service. Valid values: ECS, OSS, VPC, RDS, SLB, ODPS, GPDB, DDS, R-KVSTORE, and EIP.`,
				},
				resource.Attribute{
					Name:        "quota_type",
					Description: `(Required) The type of the quota. Valid values: organization and resourceGroup.`,
				},
				resource.Attribute{
					Name:        "quota_type_id",
					Description: `(Required) The ID of the quota type. Specify an organization ID when the QuotaType parameter is set to organization. Specify a resource set ID when the QuotaType parameter is set to resourceGroup.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Optional) The name of the cluster. This reserved parameter is optional and can be left empty.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `(Optional) This reserved parameter is optional and can be left empty. It will be used only for some products. Products where target_type are required with their values - RDS ("MySql"), R-KVSTORE ("redis") and DDS ("mongodb").`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "quotas",
					Description: `A list of Quota. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the quota.`,
				},
				resource.Attribute{
					Name:        "quota_type",
					Description: `Name of an organization, or a Resource Group.`,
				},
				resource.Attribute{
					Name:        "quota_type_id",
					Description: `ID of an organization, or a Resource Group.`,
				},
				resource.Attribute{
					Name:        "total_vip_internal",
					Description: `Total vip internal.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `name of the region where product belong.`,
				},
				resource.Attribute{
					Name:        "total_vip_public",
					Description: `Total vip public.`,
				},
				resource.Attribute{
					Name:        "total_vpc",
					Description: `Total Vpc.`,
				},
				resource.Attribute{
					Name:        "total_cpu",
					Description: `Total Cpu.`,
				},
				resource.Attribute{
					Name:        "total_cu",
					Description: `Total Cu.`,
				},
				resource.Attribute{
					Name:        "total_disk",
					Description: `Total Disk.`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `Total Mem.`,
				},
				resource.Attribute{
					Name:        "used_mem",
					Description: `Consumed Mem.`,
				},
				resource.Attribute{
					Name:        "total_gpu",
					Description: `Total Gpu.`,
				},
				resource.Attribute{
					Name:        "total_amount",
					Description: `Total Amount.`,
				},
				resource.Attribute{
					Name:        "total_disk_cloud_ssd",
					Description: `Total disk cloud ssd.`,
				},
				resource.Attribute{
					Name:        "used_disk",
					Description: `Consumed Disk.`,
				},
				resource.Attribute{
					Name:        "allocate_disk",
					Description: `Allocated Disk.`,
				},
				resource.Attribute{
					Name:        "allocate_cpu",
					Description: `Allocated Cpu.`,
				},
				resource.Attribute{
					Name:        "total_eip",
					Description: `Total Eip.`,
				},
				resource.Attribute{
					Name:        "total_disk_cloud_efficiency",
					Description: `Total disk cloud efficiency.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "quotas",
					Description: `A list of Quota. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the quota.`,
				},
				resource.Attribute{
					Name:        "quota_type",
					Description: `Name of an organization, or a Resource Group.`,
				},
				resource.Attribute{
					Name:        "quota_type_id",
					Description: `ID of an organization, or a Resource Group.`,
				},
				resource.Attribute{
					Name:        "total_vip_internal",
					Description: `Total vip internal.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `name of the region where product belong.`,
				},
				resource.Attribute{
					Name:        "total_vip_public",
					Description: `Total vip public.`,
				},
				resource.Attribute{
					Name:        "total_vpc",
					Description: `Total Vpc.`,
				},
				resource.Attribute{
					Name:        "total_cpu",
					Description: `Total Cpu.`,
				},
				resource.Attribute{
					Name:        "total_cu",
					Description: `Total Cu.`,
				},
				resource.Attribute{
					Name:        "total_disk",
					Description: `Total Disk.`,
				},
				resource.Attribute{
					Name:        "total_mem",
					Description: `Total Mem.`,
				},
				resource.Attribute{
					Name:        "used_mem",
					Description: `Consumed Mem.`,
				},
				resource.Attribute{
					Name:        "total_gpu",
					Description: `Total Gpu.`,
				},
				resource.Attribute{
					Name:        "total_amount",
					Description: `Total Amount.`,
				},
				resource.Attribute{
					Name:        "total_disk_cloud_ssd",
					Description: `Total disk cloud ssd.`,
				},
				resource.Attribute{
					Name:        "used_disk",
					Description: `Consumed Disk.`,
				},
				resource.Attribute{
					Name:        "allocate_disk",
					Description: `Allocated Disk.`,
				},
				resource.Attribute{
					Name:        "allocate_cpu",
					Description: `Allocated Cpu.`,
				},
				resource.Attribute{
					Name:        "total_eip",
					Description: `Total Eip.`,
				},
				resource.Attribute{
					Name:        "total_disk_cloud_efficiency",
					Description: `Total disk cloud efficiency.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_ram_policies",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ram policies to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ram policy IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by ram policy name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Policy name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description about the policy.`,
				},
				resource.Attribute{
					Name:        "ctime",
					Description: `Creation time of ram policy.`,
				},
				resource.Attribute{
					Name:        "cuser_id",
					Description: `ID of the policy creator.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Name of the region where policy belongs.`,
				},
				resource.Attribute{
					Name:        "policy_document",
					Description: `Policy Document.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Policy name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description about the policy.`,
				},
				resource.Attribute{
					Name:        "ctime",
					Description: `Creation time of ram policy.`,
				},
				resource.Attribute{
					Name:        "cuser_id",
					Description: `ID of the policy creator.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Name of the region where policy belongs.`,
				},
				resource.Attribute{
					Name:        "policy_document",
					Description: `Policy Document.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_ram_policies_for_user",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ram policy of the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ram policy IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by login name of the user.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Policy name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description about the policy.`,
				},
				resource.Attribute{
					Name:        "attach_date",
					Description: `Creation Date of ram policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy.`,
				},
				resource.Attribute{
					Name:        "default_version",
					Description: `Default version.`,
				},
				resource.Attribute{
					Name:        "policy_document",
					Description: `Policy Document.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `Policy name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description about the policy.`,
				},
				resource.Attribute{
					Name:        "attach_date",
					Description: `Creation Date of ram policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `Type of the policy.`,
				},
				resource.Attribute{
					Name:        "default_version",
					Description: `Default version.`,
				},
				resource.Attribute{
					Name:        "policy_document",
					Description: `Policy Document.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_ram_service_roles",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of RAM Service roles to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ram roles IDs.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `(Optional) A regex string to filter results by their product. valid values - "ECS".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description about the ram role.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `A list of roles. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `role name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description about the role.`,
				},
				resource.Attribute{
					Name:        "role_type",
					Description: `types of role.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `types of role.`,
				},
				resource.Attribute{
					Name:        "organization_name",
					Description: `Name of an Organization.`,
				},
				resource.Attribute{
					Name:        "aliyun_user_id",
					Description: `Aliyun User Id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "roles",
					Description: `A list of roles. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `role name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description about the role.`,
				},
				resource.Attribute{
					Name:        "role_type",
					Description: `types of role.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `types of role.`,
				},
				resource.Attribute{
					Name:        "organization_name",
					Description: `Name of an Organization.`,
				},
				resource.Attribute{
					Name:        "aliyun_user_id",
					Description: `Aliyun User Id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_regions_by_product",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of regions to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "product_name",
					Description: `(Required) Filter the results by specified The name of the service.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Optional) Filter the results by the specified name of the organization.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "region_list",
					Description: `A list of regions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `ID of the region.`,
				},
				resource.Attribute{
					Name:        "region_type",
					Description: `type of region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region_list",
					Description: `A list of regions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `ID of the region.`,
				},
				resource.Attribute{
					Name:        "region_type",
					Description: `type of region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_resource_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Resource Groups owned by an Alibabacloudstack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Resource Groups IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by name of Resource Group.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `(Optional) Organization ID Alibabacloudstack Cloud account.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Resource Groups IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Resource Groups names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of Resource Groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Resource Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Resource Group.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `Organization ID for Alibabacloudstack Cloud account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Resource Groups IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Resource Groups names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of Resource Groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Resource Group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Resource Group.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `Organization ID for Alibabacloudstack Cloud account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_roles",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of roles to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) It is used to filter results by role ID.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by role name.`,
				},
				resource.Attribute{
					Name:        "role_type",
					Description: `(Optional) It is used to filter results by Role Type. Valid Values - "ROLETYPE_RAM", "ROLETYPE_ASCM".`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `A list of roles. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `role name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description about the role.`,
				},
				resource.Attribute{
					Name:        "role_level",
					Description: `role level.`,
				},
				resource.Attribute{
					Name:        "role_type",
					Description: `types of role.`,
				},
				resource.Attribute{
					Name:        "ram_role",
					Description: `ram authorized role.`,
				},
				resource.Attribute{
					Name:        "role_range",
					Description: `specific range for a role.`,
				},
				resource.Attribute{
					Name:        "user_count",
					Description: `user count.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `enable.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `default role.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `Role status.`,
				},
				resource.Attribute{
					Name:        "owner_organization_id",
					Description: `ID of the owner organization where role belongs.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `role code.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "roles",
					Description: `A list of roles. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `role name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description about the role.`,
				},
				resource.Attribute{
					Name:        "role_level",
					Description: `role level.`,
				},
				resource.Attribute{
					Name:        "role_type",
					Description: `types of role.`,
				},
				resource.Attribute{
					Name:        "ram_role",
					Description: `ram authorized role.`,
				},
				resource.Attribute{
					Name:        "role_range",
					Description: `specific range for a role.`,
				},
				resource.Attribute{
					Name:        "user_count",
					Description: `user count.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `enable.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `default role.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `Role status.`,
				},
				resource.Attribute{
					Name:        "owner_organization_id",
					Description: `ID of the owner organization where role belongs.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `role code.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_service_clusters_by_product",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of service cluster to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance family IDs.`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `(Required) Filter the results by specifying name of the service.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "cluster_list",
					Description: `A list of instance families. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cluster_by_region",
					Description: `cluster by a region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_list",
					Description: `A list of instance families. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cluster_by_region",
					Description: `cluster by a region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_specific_fields",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of specific fields to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of specific fields IDs.`,
				},
				resource.Attribute{
					Name:        "group_filed",
					Description: `(Required) The field for which to query valid values.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Filter the results by the specified resource type. Valid values: OSS, ADB, DRDS, SLB, NAT, MAXCOMPUTE, POSTGRESQL, ECS, RDS, IPSIX, REDIS, MONGODB, and HITSDB.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) Specifies whether to internationalize the field. Valid values: true and false.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "specific_fields",
					Description: `A list of specific fields.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "specific_fields",
					Description: `A list of specific fields.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ascm_users",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of users to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of users IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by user login name.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `(Optional) Filter the results by the specified user Organization ID.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `A list of users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User login name.`,
				},
				resource.Attribute{
					Name:        "cell_phone_number",
					Description: `Cellphone Number of a user.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name of a user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email ID of a user.`,
				},
				resource.Attribute{
					Name:        "mobile_nation_code",
					Description: `Mobile Nation Code of a user, where user belongs to.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `User Organization ID.`,
				},
				resource.Attribute{
					Name:        "login_policy_id",
					Description: `User login policy ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "users",
					Description: `A list of users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User login name.`,
				},
				resource.Attribute{
					Name:        "cell_phone_number",
					Description: `Cellphone Number of a user.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name of a user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email ID of a user.`,
				},
				resource.Attribute{
					Name:        "mobile_nation_code",
					Description: `Mobile Nation Code of a user, where user belongs to.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `User Organization ID.`,
				},
				resource.Attribute{
					Name:        "login_policy_id",
					Description: `User login policy ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cms_alarm_contact_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of CMS Groups to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, ForceNew) A list of Alarm Contact Group IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, ForceNew) A regex string to filter results by Alarm Contact Group name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Argument Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of CMS Alarm Contact Group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of CMS groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the CMS.`,
				},
				resource.Attribute{
					Name:        "alarm_contact_group_name",
					Description: `The name of Alarm Contact Group.`,
				},
				resource.Attribute{
					Name:        "contacts",
					Description: `The alarm contacts in the alarm group.`,
				},
				resource.Attribute{
					Name:        "describe",
					Description: `The description of the Alarm Group.`,
				},
				resource.Attribute{
					Name:        "enable_subscribed",
					Description: `Indicates whether the alarm group subscribes to weekly reports.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cms_alarms",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Cms alarms.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter Alarms list by Rule name.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Optional) A regex string to filter id of the Alarm. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "alarms",
					Description: `The list of the Alarms. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the application group.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `The name of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "no_effective_interval",
					Description: `TThe time period during which the alert rule is ineffective.`,
				},
				resource.Attribute{
					Name:        "silence_time",
					Description: `Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400`,
				},
				resource.Attribute{
					Name:        "contact_groups",
					Description: `List contact groups of the alarm rule, which must have been created on the console.`,
				},
				resource.Attribute{
					Name:        "mail_subject",
					Description: `The subject of the alert notification email.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `The type of the alert rule. Valid values: METRIC: the alert rule for time series metrics. EVENT: the alert rule for event-type metrics. This type was used in earlier versions and has been discarded.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `The rule id of the Alarm.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).`,
				},
				resource.Attribute{
					Name:        "effective_interval",
					Description: `The interval of effecting alarm rule. It foramt as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `The namespace of the monitored service.`,
				},
				resource.Attribute{
					Name:        "enable_state",
					Description: `Indicates whether the alert rule was enabled. Valid values: true: indicates that the alert rule was enabled. false: indicates that the alert rule was disabled. This parameter is not specified in the request by default. In this case, both enabled and disabled rules are returned.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `The resources associated with the alert rule.`,
				},
				resource.Attribute{
					Name:        "rule_name",
					Description: `The alarm rule name.`,
				},
				resource.Attribute{
					Name:        "critical_comparison_operator",
					Description: `Critical level alarm comparison operator for critical alarm. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".`,
				},
				resource.Attribute{
					Name:        "critical_times",
					Description: `Critical level alarm retry times. Default to 3.`,
				},
				resource.Attribute{
					Name:        "critical_statistics",
					Description: `Critical level alarm statistics method for critical alarm. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".`,
				},
				resource.Attribute{
					Name:        "critical_threshold",
					Description: `Critical level alarm threshold value for critical alarm, which must be a numeric value currently.`,
				},
				resource.Attribute{
					Name:        "info_comparison_operator",
					Description: `Critical level alarm comparison operator for info alarm. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".`,
				},
				resource.Attribute{
					Name:        "info_times",
					Description: `Critical level alarm retry times for info alarm. Default to 3.`,
				},
				resource.Attribute{
					Name:        "info_statistics",
					Description: `Critical level alarm statistics method for info alarm. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".`,
				},
				resource.Attribute{
					Name:        "info_threshold",
					Description: `Critical level alarm threshold value for info alarm, which must be a numeric value currently.`,
				},
				resource.Attribute{
					Name:        "warn_comparison_operator",
					Description: `Critical level alarm comparison operator for warn alarm. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".`,
				},
				resource.Attribute{
					Name:        "warn_times",
					Description: `Critical level alarm retry times for warn alarm. Default to 3.`,
				},
				resource.Attribute{
					Name:        "warn_statistics",
					Description: `Critical level alarm statistics method for warn alarm. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".`,
				},
				resource.Attribute{
					Name:        "warn_threshold",
					Description: `Critical level alarm threshold value for warn alarm, which must be a numeric value currently.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "alarms",
					Description: `The list of the Alarms. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the application group.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `The name of the Logon Policy.`,
				},
				resource.Attribute{
					Name:        "no_effective_interval",
					Description: `TThe time period during which the alert rule is ineffective.`,
				},
				resource.Attribute{
					Name:        "silence_time",
					Description: `Notification silence period in the alarm state, in seconds. Valid value range: [300, 86400]. Default to 86400`,
				},
				resource.Attribute{
					Name:        "contact_groups",
					Description: `List contact groups of the alarm rule, which must have been created on the console.`,
				},
				resource.Attribute{
					Name:        "mail_subject",
					Description: `The subject of the alert notification email.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `The type of the alert rule. Valid values: METRIC: the alert rule for time series metrics. EVENT: the alert rule for event-type metrics. This type was used in earlier versions and has been discarded.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `The rule id of the Alarm.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `Index query cycle, which must be consistent with that defined for metrics. Default to 300, in seconds.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `Map of the resources associated with the alarm rule, such as "instanceId", "device" and "port". Each key's value is a string and it uses comma to split multiple items. For more information, see [Metrics Reference](https://www.alibabacloud.com/help/doc-detail/28619.htm).`,
				},
				resource.Attribute{
					Name:        "effective_interval",
					Description: `The interval of effecting alarm rule. It foramt as "hh:mm-hh:mm", like "0:00-4:00". Default to "00:00-23:59".`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `The namespace of the monitored service.`,
				},
				resource.Attribute{
					Name:        "enable_state",
					Description: `Indicates whether the alert rule was enabled. Valid values: true: indicates that the alert rule was enabled. false: indicates that the alert rule was disabled. This parameter is not specified in the request by default. In this case, both enabled and disabled rules are returned.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `The resources associated with the alert rule.`,
				},
				resource.Attribute{
					Name:        "rule_name",
					Description: `The alarm rule name.`,
				},
				resource.Attribute{
					Name:        "critical_comparison_operator",
					Description: `Critical level alarm comparison operator for critical alarm. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".`,
				},
				resource.Attribute{
					Name:        "critical_times",
					Description: `Critical level alarm retry times. Default to 3.`,
				},
				resource.Attribute{
					Name:        "critical_statistics",
					Description: `Critical level alarm statistics method for critical alarm. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".`,
				},
				resource.Attribute{
					Name:        "critical_threshold",
					Description: `Critical level alarm threshold value for critical alarm, which must be a numeric value currently.`,
				},
				resource.Attribute{
					Name:        "info_comparison_operator",
					Description: `Critical level alarm comparison operator for info alarm. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".`,
				},
				resource.Attribute{
					Name:        "info_times",
					Description: `Critical level alarm retry times for info alarm. Default to 3.`,
				},
				resource.Attribute{
					Name:        "info_statistics",
					Description: `Critical level alarm statistics method for info alarm. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".`,
				},
				resource.Attribute{
					Name:        "info_threshold",
					Description: `Critical level alarm threshold value for info alarm, which must be a numeric value currently.`,
				},
				resource.Attribute{
					Name:        "warn_comparison_operator",
					Description: `Critical level alarm comparison operator for warn alarm. Valid values: ["<=", "<", ">", ">=", "==", "!="]. Default to "==".`,
				},
				resource.Attribute{
					Name:        "warn_times",
					Description: `Critical level alarm retry times for warn alarm. Default to 3.`,
				},
				resource.Attribute{
					Name:        "warn_statistics",
					Description: `Critical level alarm statistics method for warn alarm. It must be consistent with that defined for metrics. Valid values: ["Average", "Minimum", "Maximum"]. Default to "Average".`,
				},
				resource.Attribute{
					Name:        "warn_threshold",
					Description: `Critical level alarm threshold value for warn alarm, which must be a numeric value currently.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cms_metric_metalist",
			Category:         "Data Sources",
			ShortDescription: `Provides a Metalist owned by an Alibabacloudstack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required, ForceNew) The namespace of the service. You can call the operation to obtain namespaces. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `A list of cms metriclist. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `The name of the metric.`,
				},
				resource.Attribute{
					Name:        "periods",
					Description: `The statistical period of the metric.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the metric.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `The dimensions of the metric. Multiple dimensions are separated with commas.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The tags of the metric. The value is a JSON array string. The array can include repeated tag names. Sample value: [{"name":"Tag name","value":"Tag value"}] . The available tag names are as follows: metricCategory: the category of the metrics. alertEnable: indicates whether the alert is enabled. alertUnit: the unit of the metric in the alert. unitFactor: the factor for metric unit conversion. minAlertPeriod: the minimum time interval to raise a new alert. productCategory: the category of the service.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `The unit of the metric.`,
				},
				resource.Attribute{
					Name:        "statistics",
					Description: `The statistical method of the metric. Multiple statistical methods are separated with commas (,), for example, Average,Minimum,Maximum.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `The namespace of the monitored service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `A list of cms metriclist. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `The name of the metric.`,
				},
				resource.Attribute{
					Name:        "periods",
					Description: `The statistical period of the metric.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the metric.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `The dimensions of the metric. Multiple dimensions are separated with commas.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The tags of the metric. The value is a JSON array string. The array can include repeated tag names. Sample value: [{"name":"Tag name","value":"Tag value"}] . The available tag names are as follows: metricCategory: the category of the metrics. alertEnable: indicates whether the alert is enabled. alertUnit: the unit of the metric in the alert. unitFactor: the factor for metric unit conversion. minAlertPeriod: the minimum time interval to raise a new alert. productCategory: the category of the service.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `The unit of the metric.`,
				},
				resource.Attribute{
					Name:        "statistics",
					Description: `The statistical method of the metric. Multiple statistical methods are separated with commas (,), for example, Average,Minimum,Maximum.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `The namespace of the monitored service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cms_project_meta",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of project meta owned by an Alibabacloudstack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, ForceNew) A regex string to filter results by project meta description. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `A list of cms project meta. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for a project meta.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels for a cms project meta. A tag of a metric is used as a special mark of alerts triggered by the metric. The format is ` + "`" + `[{"name":"Tag name","value":"Tag value"}, {"name":"Tag name","value":"Tag value"}]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `The namespace of the service, which is used to distinguish between services. Generally, the value is in the format acs_Abbreviation of the service name .`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `A list of cms project meta. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for a project meta.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels for a cms project meta. A tag of a metric is used as a special mark of alerts triggered by the metric. The format is ` + "`" + `[{"name":"Tag name","value":"Tag value"}, {"name":"Tag name","value":"Tag value"}]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `The namespace of the service, which is used to distinguish between services. Generally, the value is in the format acs_Abbreviation of the service name .`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_common_bandwidth_packages",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Common Bandwidth Packages owned by an Alibabacloudstack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Common Bandwidth Packages IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Common Bandwidth Packages IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Common Bandwidth Packages names.`,
				},
				resource.Attribute{
					Name:        "packages",
					Description: `A list of Common Bandwidth Packages. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The peak bandwidth of the Internet Shared Bandwidth instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Common Bandwidth Package instance.`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `The business status of the Common Bandwidth Package instance.`,
				},
				resource.Attribute{
					Name:        "isp",
					Description: `ISP of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `Public ip addresses that in the Common Bandwidth Package. ## Public ip addresses Block The public ip addresses mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "allocation_id",
					Description: `The ID of the EIP instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Common Bandwidth Packages IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Common Bandwidth Packages names.`,
				},
				resource.Attribute{
					Name:        "packages",
					Description: `A list of Common Bandwidth Packages. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The peak bandwidth of the Internet Shared Bandwidth instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Common Bandwidth Package instance.`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `The business status of the Common Bandwidth Package instance.`,
				},
				resource.Attribute{
					Name:        "isp",
					Description: `ISP of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `Public ip addresses that in the Common Bandwidth Package. ## Public ip addresses Block The public ip addresses mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "allocation_id",
					Description: `The ID of the EIP instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cr_ee_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Registry Enterprise Edition instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ids to filter results by instance id.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by instance name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_details",
					Description: `(Optional, Available in 1.132.0+) Default to ` + "`" + `true` + "`" + `. Set it to true can output instance authorization token. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition instances. Its element is an instance uuid.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of matched Container Registry Enterprise Editioninstances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `Specification of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "namespace_quota",
					Description: `The max number of namespaces that an instance can create.`,
				},
				resource.Attribute{
					Name:        "namespace_usage",
					Description: `The number of namespaces already created.`,
				},
				resource.Attribute{
					Name:        "repo_quota",
					Description: `The max number of repos that an instance can create.`,
				},
				resource.Attribute{
					Name:        "repo_usage",
					Description: `The number of repos already created.`,
				},
				resource.Attribute{
					Name:        "vpc_endpoints",
					Description: `A list of domains for access on vpc network.`,
				},
				resource.Attribute{
					Name:        "public_endpoints",
					Description: `A list of domains for access on internet network.`,
				},
				resource.Attribute{
					Name:        "authorization_token",
					Description: `The password that was used to log on to the registry.`,
				},
				resource.Attribute{
					Name:        "temp_username",
					Description: `The username that was used to log on to the registry.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition instances. Its element is an instance uuid.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of matched Container Registry Enterprise Editioninstances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `Specification of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "namespace_quota",
					Description: `The max number of namespaces that an instance can create.`,
				},
				resource.Attribute{
					Name:        "namespace_usage",
					Description: `The number of namespaces already created.`,
				},
				resource.Attribute{
					Name:        "repo_quota",
					Description: `The max number of repos that an instance can create.`,
				},
				resource.Attribute{
					Name:        "repo_usage",
					Description: `The number of repos already created.`,
				},
				resource.Attribute{
					Name:        "vpc_endpoints",
					Description: `A list of domains for access on vpc network.`,
				},
				resource.Attribute{
					Name:        "public_endpoints",
					Description: `A list of domains for access on internet network.`,
				},
				resource.Attribute{
					Name:        "authorization_token",
					Description: `The password that was used to log on to the registry.`,
				},
				resource.Attribute{
					Name:        "temp_username",
					Description: `The username that was used to log on to the registry.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cr_ee_namespaces",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Registry Enterprise Edition namespaces.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ids to filter results by namespace id.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by namespace name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition namespaces. Its element is a namespace uuid.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of namespace names.`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `A list of matched Container Registry Enterprise Edition namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition namespace.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition namespace.`,
				},
				resource.Attribute{
					Name:        "auto_create",
					Description: `Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.`,
				},
				resource.Attribute{
					Name:        "default_visibility",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, default repository visibility in this namespace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition namespaces. Its element is a namespace uuid.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of namespace names.`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `A list of matched Container Registry Enterprise Edition namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition namespace.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition namespace.`,
				},
				resource.Attribute{
					Name:        "auto_create",
					Description: `Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.`,
				},
				resource.Attribute{
					Name:        "default_visibility",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, default repository visibility in this namespace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cr_ee_repos",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Registry Enterprise Edition repositories.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Name of Container Registry Enterprise Edition namespace where the repositories are located in.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ids to filter results by repository id.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by repository name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_details",
					Description: `(Optional) Boolean, false by default, only repository attributes are exported. Set to true if tags belong to this repository are needed. See ` + "`" + `tags` + "`" + ` in attributes. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition repositories. Its element is a repository id.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of repository names.`,
				},
				resource.Attribute{
					Name:        "repos",
					Description: `A list of matched Container Registry Enterprise Edition namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Name of Container Registry Enterprise Edition namespace where repo is located.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition repository.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition repository.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `The repository general information.`,
				},
				resource.Attribute{
					Name:        "repo_type",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, repository's visibility.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of image tags belong to this repository. Each contains several attributes, see ` + "`" + `Block Tag` + "`" + `. ### Block Tag`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Tag of this image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Id of this image.`,
				},
				resource.Attribute{
					Name:        "digest",
					Description: `Digest of this image.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of this image.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Status of this image, in bytes.`,
				},
				resource.Attribute{
					Name:        "image_update",
					Description: `Last update time of this image, unix time in nanoseconds.`,
				},
				resource.Attribute{
					Name:        "image_create",
					Description: `Create time of this image, unix time in nanoseconds.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition repositories. Its element is a repository id.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of repository names.`,
				},
				resource.Attribute{
					Name:        "repos",
					Description: `A list of matched Container Registry Enterprise Edition namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Name of Container Registry Enterprise Edition namespace where repo is located.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition repository.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition repository.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `The repository general information.`,
				},
				resource.Attribute{
					Name:        "repo_type",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, repository's visibility.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of image tags belong to this repository. Each contains several attributes, see ` + "`" + `Block Tag` + "`" + `. ### Block Tag`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Tag of this image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Id of this image.`,
				},
				resource.Attribute{
					Name:        "digest",
					Description: `Digest of this image.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of this image.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Status of this image, in bytes.`,
				},
				resource.Attribute{
					Name:        "image_update",
					Description: `Last update time of this image, unix time in nanoseconds.`,
				},
				resource.Attribute{
					Name:        "image_create",
					Description: `Create time of this image, unix time in nanoseconds.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cr_ee_sync_rules",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Registry Enterprise Edition sync rules.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of Container Registry Enterprise Edition local instance.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `(Optional) Name of Container Registry Enterprise Edition local namespace.`,
				},
				resource.Attribute{
					Name:        "repo_name",
					Description: `(Optional) Name of Container Registry Enterprise Edition local repo.`,
				},
				resource.Attribute{
					Name:        "target_instance_id",
					Description: `(Optional) ID of Container Registry Enterprise Edition target instance.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by sync rule name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ids to filter results by sync rule id.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition sync rules. Its element is a sync rule uuid.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of sync rule names.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of matched Container Registry Enterprise Edition sync rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition sync rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition sync rule.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region of Container Registry Enterprise Edition local instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of Container Registry Enterprise Edition local instance.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `Name of Container Registry Enterprise Edition local namespace.`,
				},
				resource.Attribute{
					Name:        "repo_name",
					Description: `Name of Container Registry Enterprise Edition local repo.`,
				},
				resource.Attribute{
					Name:        "target_region_id",
					Description: `Region of Container Registry Enterprise Edition target instance.`,
				},
				resource.Attribute{
					Name:        "target_instance_id",
					Description: `ID of Container Registry Enterprise Edition target instance.`,
				},
				resource.Attribute{
					Name:        "target_namespace_name",
					Description: `Name of Container Registry Enterprise Edition target namespace.`,
				},
				resource.Attribute{
					Name:        "target_repo_name",
					Description: `Name of Container Registry Enterprise Edition target repo.`,
				},
				resource.Attribute{
					Name:        "tag_filter",
					Description: `The regular expression used to filter image tags for synchronization in the source repository.`,
				},
				resource.Attribute{
					Name:        "sync_direction",
					Description: `` + "`" + `FROM` + "`" + ` or ` + "`" + `TO` + "`" + `, the direction of synchronization. ` + "`" + `FROM` + "`" + ` indicates that the local instance is the source instance. ` + "`" + `TO` + "`" + ` indicates that the local instance is the target instance to be synchronized.`,
				},
				resource.Attribute{
					Name:        "sync_scope",
					Description: `` + "`" + `REPO` + "`" + ` or ` + "`" + `NAMESPACE` + "`" + `,the scope that the synchronization rule applies.`,
				},
				resource.Attribute{
					Name:        "sync_trigger",
					Description: `` + "`" + `PASSIVE` + "`" + ` or ` + "`" + `INITIATIVE` + "`" + `, the policy configured to trigger the synchronization rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition sync rules. Its element is a sync rule uuid.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of sync rule names.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of matched Container Registry Enterprise Edition sync rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition sync rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition sync rule.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region of Container Registry Enterprise Edition local instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of Container Registry Enterprise Edition local instance.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `Name of Container Registry Enterprise Edition local namespace.`,
				},
				resource.Attribute{
					Name:        "repo_name",
					Description: `Name of Container Registry Enterprise Edition local repo.`,
				},
				resource.Attribute{
					Name:        "target_region_id",
					Description: `Region of Container Registry Enterprise Edition target instance.`,
				},
				resource.Attribute{
					Name:        "target_instance_id",
					Description: `ID of Container Registry Enterprise Edition target instance.`,
				},
				resource.Attribute{
					Name:        "target_namespace_name",
					Description: `Name of Container Registry Enterprise Edition target namespace.`,
				},
				resource.Attribute{
					Name:        "target_repo_name",
					Description: `Name of Container Registry Enterprise Edition target repo.`,
				},
				resource.Attribute{
					Name:        "tag_filter",
					Description: `The regular expression used to filter image tags for synchronization in the source repository.`,
				},
				resource.Attribute{
					Name:        "sync_direction",
					Description: `` + "`" + `FROM` + "`" + ` or ` + "`" + `TO` + "`" + `, the direction of synchronization. ` + "`" + `FROM` + "`" + ` indicates that the local instance is the source instance. ` + "`" + `TO` + "`" + ` indicates that the local instance is the target instance to be synchronized.`,
				},
				resource.Attribute{
					Name:        "sync_scope",
					Description: `` + "`" + `REPO` + "`" + ` or ` + "`" + `NAMESPACE` + "`" + `,the scope that the synchronization rule applies.`,
				},
				resource.Attribute{
					Name:        "sync_trigger",
					Description: `` + "`" + `PASSIVE` + "`" + ` or ` + "`" + `INITIATIVE` + "`" + `, the policy configured to trigger the synchronization rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cr_namespaces",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Registry namespaces.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by namespace name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry namespaces. Its element is a namespace name.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of namespace names.`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `A list of matched Container Registry namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry namespace.`,
				},
				resource.Attribute{
					Name:        "auto_create",
					Description: `Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.`,
				},
				resource.Attribute{
					Name:        "default_visibility",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, default repository visibility in this namespace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry namespaces. Its element is a namespace name.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of namespace names.`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `A list of matched Container Registry namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry namespace.`,
				},
				resource.Attribute{
					Name:        "auto_create",
					Description: `Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.`,
				},
				resource.Attribute{
					Name:        "default_visibility",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, default repository visibility in this namespace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cr_repos",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Registry repositories.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Name of container registry namespace where the repositories are located in.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by repository name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_details",
					Description: `(Optional) Boolean, false by default, only repository attributes are exported. Set to true if domain list and tags belong to this repository are needed. See ` + "`" + `tags` + "`" + ` in attributes. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Repositories. Its element is set to ` + "`" + `names` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of repository names.`,
				},
				resource.Attribute{
					Name:        "repos",
					Description: `A list of matched Container Registry Namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Name of container registry namespace where repo is located.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of container registry namespace.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `The repository general information.`,
				},
				resource.Attribute{
					Name:        "repo_type",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, repository's visibility.`,
				},
				resource.Attribute{
					Name:        "domain_list",
					Description: `The repository domain list.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Domain of public endpoint.`,
				},
				resource.Attribute{
					Name:        "internal",
					Description: `Domain of internal endpoint, only in some regions.`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `Domain of vpc endpoint.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of image tags belong to this repository. Each contains several attributes, see ` + "`" + `Block Tag` + "`" + `. ### Block Tag`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Tag of this image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Id of this image.`,
				},
				resource.Attribute{
					Name:        "digest",
					Description: `Digest of this image.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of this image.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Status of this image, in bytes.`,
				},
				resource.Attribute{
					Name:        "image_update",
					Description: `Last update time of this image, unix time in nanoseconds.`,
				},
				resource.Attribute{
					Name:        "image_create",
					Description: `Create time of this image, unix time in nanoseconds.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Repositories. Its element is set to ` + "`" + `names` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of repository names.`,
				},
				resource.Attribute{
					Name:        "repos",
					Description: `A list of matched Container Registry Namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Name of container registry namespace where repo is located.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of container registry namespace.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `The repository general information.`,
				},
				resource.Attribute{
					Name:        "repo_type",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, repository's visibility.`,
				},
				resource.Attribute{
					Name:        "domain_list",
					Description: `The repository domain list.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Domain of public endpoint.`,
				},
				resource.Attribute{
					Name:        "internal",
					Description: `Domain of internal endpoint, only in some regions.`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `Domain of vpc endpoint.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of image tags belong to this repository. Each contains several attributes, see ` + "`" + `Block Tag` + "`" + `. ### Block Tag`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Tag of this image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Id of this image.`,
				},
				resource.Attribute{
					Name:        "digest",
					Description: `Digest of this image.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of this image.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Status of this image, in bytes.`,
				},
				resource.Attribute{
					Name:        "image_update",
					Description: `Last update time of this image, unix time in nanoseconds.`,
				},
				resource.Attribute{
					Name:        "image_create",
					Description: `Create time of this image, unix time in nanoseconds.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_cs_kubernetes_clusters",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Service Kubernetes Clusters to be used by the alibabacloudstack_cs_kubernetes_cluster resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) Cluster IDs to filter.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by cluster name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enabled_details",
					Description: `(Optional) Boolean, false by default, only ` + "`" + `id` + "`" + ` and ` + "`" + `name` + "`" + ` are exported. Set to true if more details are needed, e.g., ` + "`" + `master_disk_category` + "`" + `, ` + "`" + `slb_internet_enabled` + "`" + `, ` + "`" + `connections` + "`" + `. See full list in attributes.`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `(Optional) Boolean, Set to true to obtain kubeconfig for a cluster and add clusterId in ` + "`" + `ids` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Kubernetes clusters' ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of matched Kubernetes clusters' names.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of matched Kubernetes clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the container cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the container cluster.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of availability zone.`,
				},
				resource.Attribute{
					Name:        "worker_numbers",
					Description: `The ECS instance node number in the current container cluster.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `The ID of VSwitches where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "slb_internet_enabled",
					Description: `Whether internet load balancer for API Server is created`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of node image.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `The ID of nat gateway used to launch kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "master_instance_types",
					Description: `The instance type of master node.`,
				},
				resource.Attribute{
					Name:        "worker_instance_types",
					Description: `The instance type of worker node.`,
				},
				resource.Attribute{
					Name:        "master_disk_category",
					Description: `The system disk category of master node.`,
				},
				resource.Attribute{
					Name:        "master_disk_size",
					Description: `The system disk size of master node.`,
				},
				resource.Attribute{
					Name:        "worker_disk_category",
					Description: `The system disk category of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_disk_size",
					Description: `The system disk size of worker node.`,
				},
				resource.Attribute{
					Name:        "master_nodes",
					Description: `List of cluster master nodes. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "worker_nodes",
					Description: `List of cluster worker nodes. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `Map of kubernetes cluster connection information. It contains several attributes to ` + "`" + `Block Connections` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "node_cidr_mask",
					Description: `The network mask used on pods for each node.`,
				},
				resource.Attribute{
					Name:        "log_config",
					Description: `A list of one element containing information about the associated log store. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of collecting logs.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `Log Service project name. ### Block Nodes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the node.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Node name.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of node. ### Block Connections`,
				},
				resource.Attribute{
					Name:        "api_server_internet",
					Description: `API Server Internet endpoint.`,
				},
				resource.Attribute{
					Name:        "api_server_intranet",
					Description: `API Server Intranet endpoint.`,
				},
				resource.Attribute{
					Name:        "master_public_ip",
					Description: `Master node SSH IP address.`,
				},
				resource.Attribute{
					Name:        "service_domain",
					Description: `Service Access Domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Kubernetes clusters' ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of matched Kubernetes clusters' names.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of matched Kubernetes clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the container cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the container cluster.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of availability zone.`,
				},
				resource.Attribute{
					Name:        "worker_numbers",
					Description: `The ECS instance node number in the current container cluster.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `The ID of VSwitches where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "slb_internet_enabled",
					Description: `Whether internet load balancer for API Server is created`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of node image.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `The ID of nat gateway used to launch kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "master_instance_types",
					Description: `The instance type of master node.`,
				},
				resource.Attribute{
					Name:        "worker_instance_types",
					Description: `The instance type of worker node.`,
				},
				resource.Attribute{
					Name:        "master_disk_category",
					Description: `The system disk category of master node.`,
				},
				resource.Attribute{
					Name:        "master_disk_size",
					Description: `The system disk size of master node.`,
				},
				resource.Attribute{
					Name:        "worker_disk_category",
					Description: `The system disk category of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_disk_size",
					Description: `The system disk size of worker node.`,
				},
				resource.Attribute{
					Name:        "master_nodes",
					Description: `List of cluster master nodes. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "worker_nodes",
					Description: `List of cluster worker nodes. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `Map of kubernetes cluster connection information. It contains several attributes to ` + "`" + `Block Connections` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "node_cidr_mask",
					Description: `The network mask used on pods for each node.`,
				},
				resource.Attribute{
					Name:        "log_config",
					Description: `A list of one element containing information about the associated log store. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of collecting logs.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `Log Service project name. ### Block Nodes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the node.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Node name.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of node. ### Block Connections`,
				},
				resource.Attribute{
					Name:        "api_server_internet",
					Description: `API Server Internet endpoint.`,
				},
				resource.Attribute{
					Name:        "api_server_intranet",
					Description: `API Server Intranet endpoint.`,
				},
				resource.Attribute{
					Name:        "master_public_ip",
					Description: `Master node SSH IP address.`,
				},
				resource.Attribute{
					Name:        "service_domain",
					Description: `Service Access Domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_datahub_service",
			Category:         "Data Sources",
			ShortDescription: `Provides a datasource to open the DataHub service automatically.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Setting the value to ` + "`" + `On` + "`" + ` to enable the service. If has been enabled, return the result. Valid values: ` + "`" + `On` + "`" + ` or ` + "`" + `Off` + "`" + `. Default to ` + "`" + `Off` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current service enable status.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The current service enable status.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_db_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of RDS instances according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by instance name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of RDS instance IDs.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Optional) Database type. Options are ` + "`" + `MySQL` + "`" + `, ` + "`" + `SQLServer` + "`" + `, ` + "`" + `PostgreSQL` + "`" + ` and ` + "`" + `PPAS` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of the instance.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `(Optional) ` + "`" + `Primary` + "`" + ` for primary instance, ` + "`" + `Readonly` + "`" + ` for read-only instance, ` + "`" + `Guard` + "`" + ` for disaster recovery instance, and ` + "`" + `Temp` + "`" + ` for temporary instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Used to retrieve instances belong to specified VPC.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) Used to retrieve instances belong to specified ` + "`" + `vswitch` + "`" + ` resources.`,
				},
				resource.Attribute{
					Name:        "connection_mode",
					Description: `(Optional) ` + "`" + `Standard` + "`" + ` for standard access mode and ` + "`" + `Safe` + "`" + ` for high security access mode.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A map of tags assigned to the DB instances.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of RDS instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of RDS instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of RDS instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `` + "`" + `Primary` + "`" + ` for primary instance, ` + "`" + `Readonly` + "`" + ` for read-only instance, ` + "`" + `Guard` + "`" + ` for disaster recovery instance, and ` + "`" + `Temp` + "`" + ` for temporary instance.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expiration time. Pay-As-You-Go instances never expire.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database type. Options are ` + "`" + `MySQL` + "`" + `, ` + "`" + `SQLServer` + "`" + `, ` + "`" + `PostgreSQL` + "`" + ` and ` + "`" + `PPAS` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Database version.`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `` + "`" + `Internet` + "`" + ` for public network or ` + "`" + `Intranet` + "`" + ` for private network.`,
				},
				resource.Attribute{
					Name:        "connection_mode",
					Description: `` + "`" + `Standard` + "`" + ` for standard access mode and ` + "`" + `Safe` + "`" + ` for high security access mode.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Sizing of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "master_instance_id",
					Description: `ID of the primary instance. If this parameter is not returned, the current instance is a primary instance.`,
				},
				resource.Attribute{
					Name:        "guard_instance_id",
					Description: `If a disaster recovery instance is attached to the current instance, the ID of the disaster recovery instance applies.`,
				},
				resource.Attribute{
					Name:        "temp_instance_id",
					Description: `If a temporary instance is attached to the current instance, the ID of the temporary instance applies.`,
				},
				resource.Attribute{
					Name:        "readonly_instance_ids",
					Description: `A list of IDs of read-only instances attached to the primary instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `() RDS database connection port.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `RDS database connection string.`,
				},
				resource.Attribute{
					Name:        "instance_storage",
					Description: `User-defined DB instance storage space.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of RDS instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of RDS instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of RDS instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `` + "`" + `Primary` + "`" + ` for primary instance, ` + "`" + `Readonly` + "`" + ` for read-only instance, ` + "`" + `Guard` + "`" + ` for disaster recovery instance, and ` + "`" + `Temp` + "`" + ` for temporary instance.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expiration time. Pay-As-You-Go instances never expire.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database type. Options are ` + "`" + `MySQL` + "`" + `, ` + "`" + `SQLServer` + "`" + `, ` + "`" + `PostgreSQL` + "`" + ` and ` + "`" + `PPAS` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Database version.`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `` + "`" + `Internet` + "`" + ` for public network or ` + "`" + `Intranet` + "`" + ` for private network.`,
				},
				resource.Attribute{
					Name:        "connection_mode",
					Description: `` + "`" + `Standard` + "`" + ` for standard access mode and ` + "`" + `Safe` + "`" + ` for high security access mode.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Sizing of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "master_instance_id",
					Description: `ID of the primary instance. If this parameter is not returned, the current instance is a primary instance.`,
				},
				resource.Attribute{
					Name:        "guard_instance_id",
					Description: `If a disaster recovery instance is attached to the current instance, the ID of the disaster recovery instance applies.`,
				},
				resource.Attribute{
					Name:        "temp_instance_id",
					Description: `If a temporary instance is attached to the current instance, the ID of the temporary instance applies.`,
				},
				resource.Attribute{
					Name:        "readonly_instance_ids",
					Description: `A list of IDs of read-only instances attached to the primary instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `() RDS database connection port.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `RDS database connection string.`,
				},
				resource.Attribute{
					Name:        "instance_storage",
					Description: `User-defined DB instance storage space.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_db_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for RDS that can be used by an Alibabacloudstack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "multi",
					Description: `(Optional) Indicate whether the zones can be used in a multi AZ configuration. Default to ` + "`" + `false` + "`" + `. Multi AZ is usually used to launch RDS instances.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_disks",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of disks to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of disks IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by disk name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Disk type. Possible values: ` + "`" + `system` + "`" + ` and ` + "`" + `data` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) Disk category. Possible values: ` + "`" + `cloud` + "`" + ` (basic cloud disk), ` + "`" + `cloud_efficiency` + "`" + ` (ultra cloud disk), ` + "`" + `ephemeral_ssd` + "`" + ` (local SSD cloud disk), ` + "`" + `cloud_ssd` + "`" + ` (SSD cloud disk), and ` + "`" + `cloud_essd` + "`" + ` (ESSD cloud disk).`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) Filter the results by the specified ECS instance ID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A map of tags assigned to the disks. It must be in the format: ` + "`" + `` + "`" + `` + "`" + ` data "alibabacloudstack_disks" "disks_ds" { tags = { tagKey1 = "tagValue1", tagKey2 = "tagValue2" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "disks",
					Description: `A list of disks. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the disk.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Disk name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Disk description.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the disk belongs to.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone of the disk.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status. Possible values: ` + "`" + `In_use` + "`" + `, ` + "`" + `Available` + "`" + `, ` + "`" + `Attaching` + "`" + `, ` + "`" + `Detaching` + "`" + `, ` + "`" + `Creating` + "`" + ` and ` + "`" + `ReIniting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Disk type. Possible values: ` + "`" + `system` + "`" + ` and ` + "`" + `data` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Disk category. Possible values: ` + "`" + `cloud` + "`" + ` (basic cloud disk), ` + "`" + `cloud_efficiency` + "`" + ` (ultra cloud disk), ` + "`" + `ephemeral_ssd` + "`" + ` (local SSD cloud disk), ` + "`" + `cloud_ssd` + "`" + ` (SSD cloud disk), and ` + "`" + `cloud_essd` + "`" + ` (ESSD cloud disk).`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Indicate whether the disk is encrypted or not. Possible values: ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Disk size in GiB.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the image from which the disk is created. It is null unless the disk is created using an image.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Snapshot used to create the disk. It is null if no snapshot is used to create the disk.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the related instance. It is ` + "`" + `null` + "`" + ` unless the ` + "`" + `status` + "`" + ` is ` + "`" + `In_use` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The ID of the KMS key corresponding to the data disk.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Disk creation time.`,
				},
				resource.Attribute{
					Name:        "attached_time",
					Description: `Disk attachment time.`,
				},
				resource.Attribute{
					Name:        "detached_time",
					Description: `Disk detachment time.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `Disk expiration time.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the disk.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "disks",
					Description: `A list of disks. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the disk.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Disk name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Disk description.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the disk belongs to.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone of the disk.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status. Possible values: ` + "`" + `In_use` + "`" + `, ` + "`" + `Available` + "`" + `, ` + "`" + `Attaching` + "`" + `, ` + "`" + `Detaching` + "`" + `, ` + "`" + `Creating` + "`" + ` and ` + "`" + `ReIniting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Disk type. Possible values: ` + "`" + `system` + "`" + ` and ` + "`" + `data` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Disk category. Possible values: ` + "`" + `cloud` + "`" + ` (basic cloud disk), ` + "`" + `cloud_efficiency` + "`" + ` (ultra cloud disk), ` + "`" + `ephemeral_ssd` + "`" + ` (local SSD cloud disk), ` + "`" + `cloud_ssd` + "`" + ` (SSD cloud disk), and ` + "`" + `cloud_essd` + "`" + ` (ESSD cloud disk).`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Indicate whether the disk is encrypted or not. Possible values: ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Disk size in GiB.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the image from which the disk is created. It is null unless the disk is created using an image.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Snapshot used to create the disk. It is null if no snapshot is used to create the disk.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the related instance. It is ` + "`" + `null` + "`" + ` unless the ` + "`" + `status` + "`" + ` is ` + "`" + `In_use` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `The ID of the KMS key corresponding to the data disk.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Disk creation time.`,
				},
				resource.Attribute{
					Name:        "attached_time",
					Description: `Disk attachment time.`,
				},
				resource.Attribute{
					Name:        "detached_time",
					Description: `Disk detachment time.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `Disk expiration time.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the disk.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dms_enterprise_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available DMS Enterprise Instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Filter the results by status of the DMS Enterprise Instances. Valid values: ` + "`" + `NORMAL` + "`" + `, ` + "`" + `UNAVAILABLE` + "`" + `, ` + "`" + `UNKNOWN` + "`" + `, ` + "`" + `DELETED` + "`" + `, ` + "`" + `DISABLE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "env_type",
					Description: `(Optional) The type of the environment to which the database instance belongs.`,
				},
				resource.Attribute{
					Name:        "instance_source",
					Description: `(Optional) The source of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `(Optional) The status of the database instance.`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `(Optional) The network type of the database instance. Valid values: CLASSIC and VPC. For more information about the valid values, see the description of the RegisterInstance operation.`,
				},
				resource.Attribute{
					Name:        "search_key",
					Description: `(Optional) The keyword used to query database instances.`,
				},
				resource.Attribute{
					Name:        "tid",
					Description: `(Optional) The ID of the tenant in Data Management (DMS) Enterprise.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter the results by the DMS Enterprise Instance instance_alias.`,
				},
				resource.Attribute{
					Name:        "instance_alias_regex",
					Description: `(Optional) A regex string to filter the results by the DMS Enterprise Instance instance_alias.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of DMS Enterprise IDs (Each of them consists of host:port).`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of DMS Enterprise names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of KMS keys. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "data_link_name",
					Description: `The name of the data link for the database instance.`,
				},
				resource.Attribute{
					Name:        "database_password",
					Description: `The logon password of the database instance.`,
				},
				resource.Attribute{
					Name:        "database_user",
					Description: `The logon username of the database instance.`,
				},
				resource.Attribute{
					Name:        "dba_id",
					Description: `The ID of the database administrator (DBA) of the database instance.`,
				},
				resource.Attribute{
					Name:        "dba_nick_name",
					Description: `The nickname of the DBA.`,
				},
				resource.Attribute{
					Name:        "ddl_online",
					Description: `Indicates whether the online data description language (DDL) service was enabled for the database instance.`,
				},
				resource.Attribute{
					Name:        "ecs_instance_id",
					Description: `The ID of the Elastic Compute Service (ECS) instance to which the database instance belongs.`,
				},
				resource.Attribute{
					Name:        "ecs_region",
					Description: `The region where the database instance resides.`,
				},
				resource.Attribute{
					Name:        "env_type",
					Description: `The type of the environment to which the database instance belongs..`,
				},
				resource.Attribute{
					Name:        "export_timeout",
					Description: `The timeout period for exporting the database instance.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The endpoint of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_alias",
					Description: `The alias of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_source",
					Description: `The ID of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The ID of the database instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The connection port of the database instance.`,
				},
				resource.Attribute{
					Name:        "query_timeout",
					Description: `The timeout period for querying the database instance.`,
				},
				resource.Attribute{
					Name:        "safe_rule_id",
					Description: `The ID of the security rule for the database instance.`,
				},
				resource.Attribute{
					Name:        "sid",
					Description: `The system ID (SID) of the database instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the database instance.`,
				},
				resource.Attribute{
					Name:        "use_dsql",
					Description: `Indicates whether cross-database query was enabled for the database instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the Virtual Private Cloud (VPC) to which the database instance belongs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of DMS Enterprise IDs (Each of them consists of host:port).`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of DMS Enterprise names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of KMS keys. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "data_link_name",
					Description: `The name of the data link for the database instance.`,
				},
				resource.Attribute{
					Name:        "database_password",
					Description: `The logon password of the database instance.`,
				},
				resource.Attribute{
					Name:        "database_user",
					Description: `The logon username of the database instance.`,
				},
				resource.Attribute{
					Name:        "dba_id",
					Description: `The ID of the database administrator (DBA) of the database instance.`,
				},
				resource.Attribute{
					Name:        "dba_nick_name",
					Description: `The nickname of the DBA.`,
				},
				resource.Attribute{
					Name:        "ddl_online",
					Description: `Indicates whether the online data description language (DDL) service was enabled for the database instance.`,
				},
				resource.Attribute{
					Name:        "ecs_instance_id",
					Description: `The ID of the Elastic Compute Service (ECS) instance to which the database instance belongs.`,
				},
				resource.Attribute{
					Name:        "ecs_region",
					Description: `The region where the database instance resides.`,
				},
				resource.Attribute{
					Name:        "env_type",
					Description: `The type of the environment to which the database instance belongs..`,
				},
				resource.Attribute{
					Name:        "export_timeout",
					Description: `The timeout period for exporting the database instance.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The endpoint of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_alias",
					Description: `The alias of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_source",
					Description: `The ID of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The ID of the database instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The connection port of the database instance.`,
				},
				resource.Attribute{
					Name:        "query_timeout",
					Description: `The timeout period for querying the database instance.`,
				},
				resource.Attribute{
					Name:        "safe_rule_id",
					Description: `The ID of the security rule for the database instance.`,
				},
				resource.Attribute{
					Name:        "sid",
					Description: `The system ID (SID) of the database instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the database instance.`,
				},
				resource.Attribute{
					Name:        "use_dsql",
					Description: `Indicates whether cross-database query was enabled for the database instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the Virtual Private Cloud (VPC) to which the database instance belongs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dms_enterprise_users",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available DMS Enterprise Users.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The role of the user to query.`,
				},
				resource.Attribute{
					Name:        "search_key",
					Description: `(Optional) The keyword used to query users.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the user.`,
				},
				resource.Attribute{
					Name:        "tid",
					Description: `(Optional) The ID of the tenant in DMS Enterprise.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of DMS Enterprise User IDs (UID).`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter the results by the DMS Enterprise User nick_name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of DMS Enterprise User IDs (UID).`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of DMS Enterprise User names.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `A list of DMS Enterprise Users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "mobile",
					Description: `The DingTalk number or mobile number of the user.`,
				},
				resource.Attribute{
					Name:        "nick_name",
					Description: `The nickname of the user.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The nickname of the user.`,
				},
				resource.Attribute{
					Name:        "parent_uid",
					Description: `The Alibaba Cloud unique ID (UID) of the parent account if the user corresponds to a Resource Access Management (RAM) user.`,
				},
				resource.Attribute{
					Name:        "role_ids",
					Description: `The list ids of the role that the user plays.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `The list names of the role that he user plays.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Alibaba Cloud unique ID (UID) of the user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of DMS Enterprise User IDs (UID).`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of DMS Enterprise User names.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `A list of DMS Enterprise Users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "mobile",
					Description: `The DingTalk number or mobile number of the user.`,
				},
				resource.Attribute{
					Name:        "nick_name",
					Description: `The nickname of the user.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The nickname of the user.`,
				},
				resource.Attribute{
					Name:        "parent_uid",
					Description: `The Alibaba Cloud unique ID (UID) of the parent account if the user corresponds to a Resource Access Management (RAM) user.`,
				},
				resource.Attribute{
					Name:        "role_ids",
					Description: `The list ids of the role that the user plays.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `The list names of the role that he user plays.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Alibaba Cloud unique ID (UID) of the user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dns_domains",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of domains available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional) A regex string to filter results by the domain name.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew) The ID of resource group which the dns belongs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of domain IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of domain names.`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `A list of domains. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `ID of the domain.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Name of the domain.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `DNS list of the domain in the analysis system.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of resource group which the dns belongs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of domain IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of domain names.`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `A list of domains. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `ID of the domain.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Name of the domain.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `DNS list of the domain in the analysis system.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of resource group which the dns belongs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dns_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of groups available to the dns.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by group name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of group IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Id of the group.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `Name of the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Id of the group.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `Name of the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_dns_records",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of records available to the dns.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The domain Id associated to the records.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Host record regex.`,
				},
				resource.Attribute{
					Name:        "value_regex",
					Description: `(Optional) Host record value regex.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type. Valid items are ` + "`" + `A` + "`" + `, ` + "`" + `NS` + "`" + `, ` + "`" + `MX` + "`" + `, ` + "`" + `TXT` + "`" + `, ` + "`" + `CNAME` + "`" + `, ` + "`" + `SRV` + "`" + `, ` + "`" + `AAAA` + "`" + `, ` + "`" + `REDIRECT_URL` + "`" + `, ` + "`" + `FORWORD_URL` + "`" + ` .`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of record IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of record IDs.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `A list of records. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "record_id",
					Description: `ID of the record.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `ID of the domain the record belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Host record of the domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `TTL of the record.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Description of the record.`,
				},
				resource.Attribute{
					Name:        "rr_set",
					Description: `RrSet for the record.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of record IDs.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `A list of records. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "record_id",
					Description: `ID of the record.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `ID of the domain the record belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Host record of the domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `TTL of the record.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Description of the record.`,
				},
				resource.Attribute{
					Name:        "rr_set",
					Description: `RrSet for the record.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_drds_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of DRDS instances according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, Deprecated) A regex string to filter results by instance description. It is deprecated since v1.91.0 and will be removed in a future release, please use 'description_regex' instead.`,
				},
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional) A regex string to filter results by instance description.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of DRDS instance IDs. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of DRDS instance IDs.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of DRDS descriptions.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of DRDS instances.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the DRDS instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The DRDS instance description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The DRDS Instance type.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `` + "`" + `Classic` + "`" + ` for public classic network or ` + "`" + `VPC` + "`" + ` for private network.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `Zone ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The DRDS Instance version.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of DRDS instance IDs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of DRDS instance IDs.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of DRDS descriptions.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of DRDS instances.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the DRDS instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The DRDS instance description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The DRDS Instance type.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `` + "`" + `Classic` + "`" + ` for public classic network or ` + "`" + `VPC` + "`" + ` for private network.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `Zone ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The DRDS Instance version.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of DRDS instance IDs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ecs_auto_snapshot_policies",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Ecs Auto Snapshot Policies to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, ForceNew, Computed) A list of Auto Snapshot Policy IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, ForceNew) A regex string to filter results by Auto Snapshot Policy name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, ForceNew) The status of Auto Snapshot Policy. Valid Values: ` + "`" + `Expire` + "`" + `, ` + "`" + `Normal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Argument Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Auto Snapshot Policy names.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `A list of Ecs Auto Snapshot Policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "auto_snapshot_policy_id",
					Description: `The ID of the Auto Snapshot Policy.`,
				},
				resource.Attribute{
					Name:        "copied_snapshots_retention_days",
					Description: `The retention period of the snapshot copied across regions.`,
				},
				resource.Attribute{
					Name:        "disk_nums",
					Description: `The number of disks to which the automatic snapshot policy is applied.`,
				},
				resource.Attribute{
					Name:        "enable_cross_region_copy",
					Description: `Specifies whether to enable the system to automatically copy snapshots across regions.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Auto Snapshot Policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The snapshot policy name..`,
				},
				resource.Attribute{
					Name:        "repeat_weekdays",
					Description: `The automatic snapshot repetition dates.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `The snapshot retention time, and the unit of measurement is day.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of Auto Snapshot Policy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "target_copy_regions",
					Description: `The destination region to which the snapshot is copied.`,
				},
				resource.Attribute{
					Name:        "time_points",
					Description: `The automatic snapshot creation schedule, and the unit of measurement is hour.`,
				},
				resource.Attribute{
					Name:        "volume_nums",
					Description: `The number of extended volumes on which this policy is enabled.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ecs_commands",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Ecs Commands to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Optional, ForceNew) The Base64-encoded content of the command.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, ForceNew) The description of command.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, ForceNew, Computed) A list of Command IDs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, ForceNew) The name of the command.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, ForceNew) A regex string to filter results by Command name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "command_provider",
					Description: `(Optional, ForceNew) Public order provider.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, ForceNew) The command type. Valid Values: ` + "`" + `RunBatScript` + "`" + `, ` + "`" + `RunPowerShellScript` + "`" + ` and ` + "`" + `RunShellScript` + "`" + `. ## Argument Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Command names.`,
				},
				resource.Attribute{
					Name:        "commands",
					Description: `A list of Ecs Commands. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "command_content",
					Description: `The Base64-encoded content of the command.`,
				},
				resource.Attribute{
					Name:        "command_id",
					Description: `The ID of the Command.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of command.`,
				},
				resource.Attribute{
					Name:        "enable_parameter",
					Description: `Specifies whether to use custom parameters in the command to be created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Command.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the command`,
				},
				resource.Attribute{
					Name:        "parameter_names",
					Description: `A list of custom parameter names which are parsed from the command content specified when the command was being created.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `The timeout period that is specified for the command to be run on ECS instances.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The command type.`,
				},
				resource.Attribute{
					Name:        "working_dir",
					Description: `The execution path of the command in the ECS instance.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ecs_dedicated_hosts",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available ECS Dedicated Hosts.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ECS Dedicated Host ids.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the ECS Dedicated Host name.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_id",
					Description: `(Optional) The ID of ECS Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_name",
					Description: `(Optional) The name of ECS Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_type",
					Description: `(Optional) The type of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "operation_locks",
					Description: `(Optional, Available in 1.123.1+) The reason why the dedicated host resource is locked.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional) The ID of the resource group to which the ECS Dedicated Host belongs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the ECS Dedicated Host. validate value: ` + "`" + `Available` + "`" + `, ` + "`" + `Creating` + "`" + `, ` + "`" + `PermanentFailure` + "`" + `, ` + "`" + `Released` + "`" + `, ` + "`" + `UnderAssessment` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The zone ID of the ECS Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) Save the result to the file. #### Block operation_locks The operation_locks supports the following:`,
				},
				resource.Attribute{
					Name:        "lock_reason",
					Description: `(Optional, ForceNew) The reason why the dedicated host resource is locked. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of ECS Dedicated Host ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of ECS Dedicated Host names.`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `A list of ECS Dedicated Hosts. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ECS Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "action_on_maintenance",
					Description: `The policy used to migrate the instances from the dedicated host when the dedicated host fails or needs to be repaired online.`,
				},
				resource.Attribute{
					Name:        "auto_placement",
					Description: `Specifies whether to add the dedicated host to the resource pool for automatic deployment.`,
				},
				resource.Attribute{
					Name:        "auto_release_time",
					Description: `The automatic release time of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_id",
					Description: `ID of the ECS Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_name",
					Description: `The name of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_type",
					Description: `The type of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `The expiration time of the subscription dedicated host.`,
				},
				resource.Attribute{
					Name:        "gpu_spec",
					Description: `The GPU model.`,
				},
				resource.Attribute{
					Name:        "machine_id",
					Description: `The machine code of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "payment_type",
					Description: `The billing method of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "physical_gpus",
					Description: `The number of physical GPUs.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of the resource group to which the dedicated host belongs.`,
				},
				resource.Attribute{
					Name:        "sale_cycle",
					Description: `The unit of the subscription billing method.`,
				},
				resource.Attribute{
					Name:        "sockets",
					Description: `The number of physical CPUs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The service status of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "supported_instance_types_list",
					Description: `The list of ECS instance`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Available in 1.123.1+) A collection of proprietary host performance indicators.`,
				},
				resource.Attribute{
					Name:        "available_local_storage",
					Description: `The remaining local disk capacity. Unit: GiB.`,
				},
				resource.Attribute{
					Name:        "available_memory",
					Description: `The remaining memory capacity, unit: GiB.`,
				},
				resource.Attribute{
					Name:        "available_vcpus",
					Description: `The number of remaining vCPU cores.`,
				},
				resource.Attribute{
					Name:        "available_vgpus",
					Description: `The number of available virtual GPUs.`,
				},
				resource.Attribute{
					Name:        "local_storage_category",
					Description: `Local disk type.`,
				},
				resource.Attribute{
					Name:        "total_local_storage",
					Description: `The total capacity of the local disk, in GiB.`,
				},
				resource.Attribute{
					Name:        "total_memory",
					Description: `The total memory capacity, unit: GiB.`,
				},
				resource.Attribute{
					Name:        "total_vcpus",
					Description: `The total number of vCPU cores.`,
				},
				resource.Attribute{
					Name:        "total_vgpus",
					Description: `The total number of virtual GPUs.`,
				},
				resource.Attribute{
					Name:        "cpu_over_commit_ratio",
					Description: `(Available in 1.123.1+) CPU oversold ratio.`,
				},
				resource.Attribute{
					Name:        "network_attributes",
					Description: `dedicated host network parameters. contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "slb_udp_timeout",
					Description: `The timeout period for a UDP session between Server Load Balancer (SLB) and the dedicated host. Unit: seconds.`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `(Available in 1.123.1+) The timeout period for a UDP session between a user and an Apsara Stack Cloud service on the dedicated host. Unit: seconds.`,
				},
				resource.Attribute{
					Name:        "operation_locks",
					Description: `(Available in 1.123.1+) The operation_locks. contains the following attribute:`,
				},
				resource.Attribute{
					Name:        "lock_reason",
					Description: `The reason why the dedicated host resource is locked.`,
				},
				resource.Attribute{
					Name:        "supported_instance_type_families",
					Description: `(Available in 1.123.1+) ECS instance type family supported by the dedicated host.`,
				},
				resource.Attribute{
					Name:        "supported_custom_instance_type_families",
					Description: `(Available in 1.123.1+) A custom instance type family supported by a dedicated host.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of ECS Dedicated Host ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of ECS Dedicated Host names.`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `A list of ECS Dedicated Hosts. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ECS Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "action_on_maintenance",
					Description: `The policy used to migrate the instances from the dedicated host when the dedicated host fails or needs to be repaired online.`,
				},
				resource.Attribute{
					Name:        "auto_placement",
					Description: `Specifies whether to add the dedicated host to the resource pool for automatic deployment.`,
				},
				resource.Attribute{
					Name:        "auto_release_time",
					Description: `The automatic release time of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_id",
					Description: `ID of the ECS Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_name",
					Description: `The name of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_type",
					Description: `The type of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `The expiration time of the subscription dedicated host.`,
				},
				resource.Attribute{
					Name:        "gpu_spec",
					Description: `The GPU model.`,
				},
				resource.Attribute{
					Name:        "machine_id",
					Description: `The machine code of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "payment_type",
					Description: `The billing method of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "physical_gpus",
					Description: `The number of physical GPUs.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of the resource group to which the dedicated host belongs.`,
				},
				resource.Attribute{
					Name:        "sale_cycle",
					Description: `The unit of the subscription billing method.`,
				},
				resource.Attribute{
					Name:        "sockets",
					Description: `The number of physical CPUs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The service status of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "supported_instance_types_list",
					Description: `The list of ECS instance`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Available in 1.123.1+) A collection of proprietary host performance indicators.`,
				},
				resource.Attribute{
					Name:        "available_local_storage",
					Description: `The remaining local disk capacity. Unit: GiB.`,
				},
				resource.Attribute{
					Name:        "available_memory",
					Description: `The remaining memory capacity, unit: GiB.`,
				},
				resource.Attribute{
					Name:        "available_vcpus",
					Description: `The number of remaining vCPU cores.`,
				},
				resource.Attribute{
					Name:        "available_vgpus",
					Description: `The number of available virtual GPUs.`,
				},
				resource.Attribute{
					Name:        "local_storage_category",
					Description: `Local disk type.`,
				},
				resource.Attribute{
					Name:        "total_local_storage",
					Description: `The total capacity of the local disk, in GiB.`,
				},
				resource.Attribute{
					Name:        "total_memory",
					Description: `The total memory capacity, unit: GiB.`,
				},
				resource.Attribute{
					Name:        "total_vcpus",
					Description: `The total number of vCPU cores.`,
				},
				resource.Attribute{
					Name:        "total_vgpus",
					Description: `The total number of virtual GPUs.`,
				},
				resource.Attribute{
					Name:        "cpu_over_commit_ratio",
					Description: `(Available in 1.123.1+) CPU oversold ratio.`,
				},
				resource.Attribute{
					Name:        "network_attributes",
					Description: `dedicated host network parameters. contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "slb_udp_timeout",
					Description: `The timeout period for a UDP session between Server Load Balancer (SLB) and the dedicated host. Unit: seconds.`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `(Available in 1.123.1+) The timeout period for a UDP session between a user and an Apsara Stack Cloud service on the dedicated host. Unit: seconds.`,
				},
				resource.Attribute{
					Name:        "operation_locks",
					Description: `(Available in 1.123.1+) The operation_locks. contains the following attribute:`,
				},
				resource.Attribute{
					Name:        "lock_reason",
					Description: `The reason why the dedicated host resource is locked.`,
				},
				resource.Attribute{
					Name:        "supported_instance_type_families",
					Description: `(Available in 1.123.1+) ECS instance type family supported by the dedicated host.`,
				},
				resource.Attribute{
					Name:        "supported_custom_instance_type_families",
					Description: `(Available in 1.123.1+) A custom instance type family supported by a dedicated host.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ehpc_job_templates",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Ehpc Job Templates to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, ForceNew, Computed) A list of Job Template IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Argument Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "templates",
					Description: `A list of Ehpc Job Templates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "array_request",
					Description: `Queue Jobs, Is of the Form: 1-10:2.`,
				},
				resource.Attribute{
					Name:        "clock_time",
					Description: `Job Maximum Run Time.`,
				},
				resource.Attribute{
					Name:        "command_line",
					Description: `Job Commands.`,
				},
				resource.Attribute{
					Name:        "gpu",
					Description: `A Single Compute Node Using the GPU Number.Possible Values: 1~20000.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Job Template.`,
				},
				resource.Attribute{
					Name:        "job_template_id",
					Description: `The first ID of the resource.`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `A Single Compute Node Maximum Memory.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `Submit a Task Is Required for Computing the Number of Data Nodes to Be. Possible Values: 1~5000 .`,
				},
				resource.Attribute{
					Name:        "package_path",
					Description: `Job Commands the Directory.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The Job Priority.Possible Values: 0~9.`,
				},
				resource.Attribute{
					Name:        "queue",
					Description: `The Job Queue.`,
				},
				resource.Attribute{
					Name:        "re_runable",
					Description: `If the Job Is Support for the Re-Run.`,
				},
				resource.Attribute{
					Name:        "runas_user",
					Description: `The name of the user who performed the job.`,
				},
				resource.Attribute{
					Name:        "stderr_redirect_path",
					Description: `Error Output Path.`,
				},
				resource.Attribute{
					Name:        "stdout_redirect_path",
					Description: `Standard Output Path and.`,
				},
				resource.Attribute{
					Name:        "task",
					Description: `A Single Compute Node Required Number of Tasks. Possible Values: 1~20000 .`,
				},
				resource.Attribute{
					Name:        "thread",
					Description: `A Single Task and the Number of Required Threads.Possible Values: 1~20000.`,
				},
				resource.Attribute{
					Name:        "variables",
					Description: `The Job of the Environment Variable.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_eips",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of EIP owned by an Alibabacloudstack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of EIP IDs.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(Optional) A list of EIP public IP addresses.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of EIP IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(Optional) A list of EIP names.`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `A list of EIPs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `EIP status. Possible values are: ` + "`" + `Associating` + "`" + `, ` + "`" + `Unassociating` + "`" + `, ` + "`" + `InUse` + "`" + ` and ` + "`" + `Available` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Public IP Address of the the EIP.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `EIP internet max bandwidth in Mbps.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance that is being bound.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The instance type of that the EIP is bound.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of EIP IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(Optional) A list of EIP names.`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `A list of EIPs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `EIP status. Possible values are: ` + "`" + `Associating` + "`" + `, ` + "`" + `Unassociating` + "`" + `, ` + "`" + `InUse` + "`" + ` and ` + "`" + `Available` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Public IP Address of the the EIP.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `EIP internet max bandwidth in Mbps.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance that is being bound.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The instance type of that the EIP is bound.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_elasticsearch",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of Elasticsearch instances according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional) A regex string to apply to the instance description.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.52.1+) A list of Elasticsearch instance IDs.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Elasticsearch version. Options are ` + "`" + `5.5.3_with_X-Pack` + "`" + `, ` + "`" + `6.3.2_with_X-Pack` + "`" + ` and ` + "`" + `6.7.0_with_X-Pack` + "`" + `. If no value is specified, all versions are returned.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available 1.74.0+) A map of tags assigned to instances.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Elasticsearch instance IDs.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of Elasticsearch instance descriptions.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of Elasticsearch instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `Billing method. Value options: ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for subscription.`,
				},
				resource.Attribute{
					Name:        "data_node_amount",
					Description: `The Elasticsearch cluster's data node quantity, between 2 and 50.`,
				},
				resource.Attribute{
					Name:        "data_node_spec",
					Description: `The data node specifications of the elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "data_node_disk_size",
					Description: `The single data node storage space. Unit: GB.`,
				},
				resource.Attribute{
					Name:        "data_node_disk_type",
					Description: `The data node disk type. Included values: ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VSwitch ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Elasticsearch version includes ` + "`" + `5.5.3_with_X-Pack` + "`" + `, ` + "`" + `6.3.2_with_X-Pack` + "`" + ` and ` + "`" + `6.7.0_with_X-Pack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cerated_at",
					Description: `The creation time of the instance. It's a GTM format, such as: "2019-01-08T15:50:50.623Z".`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The last modified time of the instance. It's a GMT format, such as: "2019-01-08T15:50:50.623Z".`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance. It includes ` + "`" + `active` + "`" + `, ` + "`" + `activating` + "`" + `, ` + "`" + `inactive` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Elasticsearch instance IDs.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of Elasticsearch instance descriptions.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of Elasticsearch instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `Billing method. Value options: ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for subscription.`,
				},
				resource.Attribute{
					Name:        "data_node_amount",
					Description: `The Elasticsearch cluster's data node quantity, between 2 and 50.`,
				},
				resource.Attribute{
					Name:        "data_node_spec",
					Description: `The data node specifications of the elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "data_node_disk_size",
					Description: `The single data node storage space. Unit: GB.`,
				},
				resource.Attribute{
					Name:        "data_node_disk_type",
					Description: `The data node disk type. Included values: ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VSwitch ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Elasticsearch version includes ` + "`" + `5.5.3_with_X-Pack` + "`" + `, ` + "`" + `6.3.2_with_X-Pack` + "`" + ` and ` + "`" + `6.7.0_with_X-Pack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cerated_at",
					Description: `The creation time of the instance. It's a GTM format, such as: "2019-01-08T15:50:50.623Z".`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The last modified time of the instance. It's a GMT format, such as: "2019-01-08T15:50:50.623Z".`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance. It includes ` + "`" + `active` + "`" + `, ` + "`" + `activating` + "`" + `, ` + "`" + `inactive` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_elasticsearch_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for Elasticsearch that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "multi",
					Description: `(Optional) Indicate whether the zones can be used in a multi AZ configuration. Default to ` + "`" + `false` + "`" + `. Multi AZ is usually used to launch Elasticsearch instances.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_lifecycle_hooks",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of lifecycle hooks available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Optional) Scaling group id the lifecycle hooks belong to.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting lifecycle hook by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of lifecycle hook IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of lifecycle hook ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of lifecycle hook names.`,
				},
				resource.Attribute{
					Name:        "hooks",
					Description: `A list of lifecycle hooks. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "default_result",
					Description: `Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses.`,
				},
				resource.Attribute{
					Name:        "heartbeat_timeout",
					Description: `Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter.`,
				},
				resource.Attribute{
					Name:        "lifecycle_transition",
					Description: `Type of Scaling activity attached to lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "notification_arn",
					Description: `The Arn of notification target.`,
				},
				resource.Attribute{
					Name:        "notification_metadata",
					Description: `Additional information that you want to include when Auto Scaling sends a message to the notification target.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of lifecycle hook ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of lifecycle hook names.`,
				},
				resource.Attribute{
					Name:        "hooks",
					Description: `A list of lifecycle hooks. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "default_result",
					Description: `Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses.`,
				},
				resource.Attribute{
					Name:        "heartbeat_timeout",
					Description: `Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter.`,
				},
				resource.Attribute{
					Name:        "lifecycle_transition",
					Description: `Type of Scaling activity attached to lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "notification_arn",
					Description: `The Arn of notification target.`,
				},
				resource.Attribute{
					Name:        "notification_metadata",
					Description: `Additional information that you want to include when Auto Scaling sends a message to the notification target.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_notifications",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of notifications available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required) Scaling group id the notifications belong to.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional)A list of notification ids.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of notification ids.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `A list of notifications. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the notification.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "notification_arn",
					Description: `The Alibaba Cloud Resource Name (ARN) for the notification object.`,
				},
				resource.Attribute{
					Name:        "notification_types",
					Description: `The notification types of Auto Scaling events and resource changes.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of notification ids.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `A list of notifications. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the notification.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "notification_arn",
					Description: `The Alibaba Cloud Resource Name (ARN) for the notification object.`,
				},
				resource.Attribute{
					Name:        "notification_types",
					Description: `The notification types of Auto Scaling events and resource changes.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_scaling_configurations",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of scaling configurations available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Optional) Scaling group id the scaling configurations belong to.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting scaling configurations by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of scaling configuration IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scaling configuration ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scaling configuration names.`,
				},
				resource.Attribute{
					Name:        "configurations",
					Description: `A list of scaling rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Image ID of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group ID of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_in",
					Description: `Internet max bandwidth in of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Internet max bandwidth of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "system_disk_category",
					Description: `System disk category of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `System disk size of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `Data disks of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of data disk.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Category of data disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Size of data disk.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device attribute of data disk.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `Delete_with_instance attribute of data disk.`,
				},
				resource.Attribute{
					Name:        "lifecycle_state",
					Description: `Lifecycle state of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the scaling configuration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scaling configuration ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scaling configuration names.`,
				},
				resource.Attribute{
					Name:        "configurations",
					Description: `A list of scaling rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Image ID of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group ID of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_in",
					Description: `Internet max bandwidth in of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Internet max bandwidth of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "system_disk_category",
					Description: `System disk category of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `System disk size of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `Data disks of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of data disk.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Category of data disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Size of data disk.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device attribute of data disk.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `Delete_with_instance attribute of data disk.`,
				},
				resource.Attribute{
					Name:        "lifecycle_state",
					Description: `Lifecycle state of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the scaling configuration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_scaling_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of scaling groups available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting scaling groups by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of scaling group IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scaling group ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scaling group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of scaling groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scaling group.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the scaling group belongs to.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `The minimum number of ECS instances.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum number of ECS instances.`,
				},
				resource.Attribute{
					Name:        "cooldown_time",
					Description: `Default cooldown time of scaling group.`,
				},
				resource.Attribute{
					Name:        "removal_policies",
					Description: `Removal policy used to select the ECS instance to remove from the scaling group.`,
				},
				resource.Attribute{
					Name:        "load_balancer_ids",
					Description: `Slb instances id which the ECS instance attached to.`,
				},
				resource.Attribute{
					Name:        "db_instance_ids",
					Description: `Db instances id which the ECS instance attached to.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `Vswitches id in which the ECS instance launched.`,
				},
				resource.Attribute{
					Name:        "lifecycle_state",
					Description: `Lifecycle state of scaling group.`,
				},
				resource.Attribute{
					Name:        "total_capacity",
					Description: `Number of instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "active_capacity",
					Description: `Number of active instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "pending_capacity",
					Description: `Number of pending instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "removing_capacity",
					Description: `Number of removing instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of scaling group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scaling group ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scaling group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of scaling groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scaling group.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the scaling group belongs to.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `The minimum number of ECS instances.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum number of ECS instances.`,
				},
				resource.Attribute{
					Name:        "cooldown_time",
					Description: `Default cooldown time of scaling group.`,
				},
				resource.Attribute{
					Name:        "removal_policies",
					Description: `Removal policy used to select the ECS instance to remove from the scaling group.`,
				},
				resource.Attribute{
					Name:        "load_balancer_ids",
					Description: `Slb instances id which the ECS instance attached to.`,
				},
				resource.Attribute{
					Name:        "db_instance_ids",
					Description: `Db instances id which the ECS instance attached to.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `Vswitches id in which the ECS instance launched.`,
				},
				resource.Attribute{
					Name:        "lifecycle_state",
					Description: `Lifecycle state of scaling group.`,
				},
				resource.Attribute{
					Name:        "total_capacity",
					Description: `Number of instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "active_capacity",
					Description: `Number of active instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "pending_capacity",
					Description: `Number of pending instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "removing_capacity",
					Description: `Number of removing instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of scaling group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_scaling_rules",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of scaling rules available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Optional) Scaling group id the scaling rules belong to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of scaling rule.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting scaling rules by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of scaling rule IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scaling rule ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scaling rule names.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of scaling rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `Cooldown time of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "adjustment_type",
					Description: `Adjustment type of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "adjustment_value",
					Description: `Adjustment value of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "min_adjustment_magnitude",
					Description: `Min adjustment magnitude of scaling rule.`,
				},
				resource.Attribute{
					Name:        "scaling_rule_ari",
					Description: `Ari of scaling rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scaling rule ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scaling rule names.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of scaling rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `Cooldown time of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "adjustment_type",
					Description: `Adjustment type of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "adjustment_value",
					Description: `Adjustment value of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "min_adjustment_magnitude",
					Description: `Min adjustment magnitude of scaling rule.`,
				},
				resource.Attribute{
					Name:        "scaling_rule_ari",
					Description: `Ari of scaling rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ess_scheduled_tasks",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of scheduled tasks available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scheduled_task_id",
					Description: `(Optional) The id of the scheduled task.`,
				},
				resource.Attribute{
					Name:        "scheduled_action",
					Description: `(Optional) The operation to be performed when a scheduled task is triggered.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting scheduled tasks by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of scheduled task IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scheduled task ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scheduled task names.`,
				},
				resource.Attribute{
					Name:        "tasks",
					Description: `A list of scheduled tasks. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scheduled task id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scheduled task name.`,
				},
				resource.Attribute{
					Name:        "scheduled_action",
					Description: `The operation to be performed when a scheduled task is triggered.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the scheduled task.`,
				},
				resource.Attribute{
					Name:        "launch_expiration_time",
					Description: `The time period during which a failed scheduled task is retried.`,
				},
				resource.Attribute{
					Name:        "launch_time",
					Description: `The time at which the scheduled task is triggered.`,
				},
				resource.Attribute{
					Name:        "recurrence_type",
					Description: `Specifies the recurrence type of the scheduled task.`,
				},
				resource.Attribute{
					Name:        "recurrence_value",
					Description: `Specifies how often a scheduled task recurs.`,
				},
				resource.Attribute{
					Name:        "recurrence_end_time",
					Description: `Specifies the end time after which the scheduled task is no longer repeated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scheduled task ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scheduled task names.`,
				},
				resource.Attribute{
					Name:        "tasks",
					Description: `A list of scheduled tasks. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scheduled task id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scheduled task name.`,
				},
				resource.Attribute{
					Name:        "scheduled_action",
					Description: `The operation to be performed when a scheduled task is triggered.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the scheduled task.`,
				},
				resource.Attribute{
					Name:        "launch_expiration_time",
					Description: `The time period during which a failed scheduled task is retried.`,
				},
				resource.Attribute{
					Name:        "launch_time",
					Description: `The time at which the scheduled task is triggered.`,
				},
				resource.Attribute{
					Name:        "recurrence_type",
					Description: `Specifies the recurrence type of the scheduled task.`,
				},
				resource.Attribute{
					Name:        "recurrence_value",
					Description: `Specifies how often a scheduled task recurs.`,
				},
				resource.Attribute{
					Name:        "recurrence_end_time",
					Description: `Specifies the end time after which the scheduled task is no longer repeated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_express_connect_access_points",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Express Connect Access Points to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, ForceNew, Computed) A list of Access Point IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, ForceNew) A regex string to filter results by Access Point name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, ForceNew) The Physical Connection to Which the Access Point State. Valid values: ` + "`" + `disabled` + "`" + `, ` + "`" + `full` + "`" + `, ` + "`" + `hot` + "`" + `, ` + "`" + `recommended` + "`" + `. ## Argument Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Access Point names.`,
				},
				resource.Attribute{
					Name:        "points",
					Description: `A list of Express Connect Access Points. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_point_id",
					Description: `The Access Point ID.`,
				},
				resource.Attribute{
					Name:        "access_point_name",
					Description: `Access Point Name.`,
				},
				resource.Attribute{
					Name:        "attached_region_no",
					Description: `The Access Point Is Located an ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The Access Point Description.`,
				},
				resource.Attribute{
					Name:        "host_operator",
					Description: `The Access Point Belongs to the Operator.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Access Point.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Location of the Access Point.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The Physical Connection to Which the Access Point State.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Physical Connection to Which the Network Type.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_express_connect_physical_connections",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Express Connect Physical Connections to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, ForceNew, Computed) A list of Physical Connection IDs.`,
				},
				resource.Attribute{
					Name:        "include_reservation_data",
					Description: `(Optional, ForceNew) The include reservation data.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, ForceNew) A regex string to filter results by Physical Connection name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, ForceNew) Resources on Behalf of a State of the Resource Attribute Field. Valid values: ` + "`" + `Canceled` + "`" + `, ` + "`" + `Enabled` + "`" + `, ` + "`" + `Terminated` + "`" + `. ## Argument Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Physical Connection names.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `A list of Express Connect Physical Connections. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_point_id",
					Description: `The Physical Leased Line Access Point ID.`,
				},
				resource.Attribute{
					Name:        "ad_location",
					Description: `To Connect a Device Physical Location.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `On the Bandwidth of the ECC Service and Physical Connection.`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `The Physical Connection to Which the Payment Status: Normal, financiallocked, securitylocked.`,
				},
				resource.Attribute{
					Name:        "circuit_code",
					Description: `Operators for Physical Connection Circuit Provided Coding.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The Representative of the Creation Time Resources Attribute Field.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The Physical Connection to Which the Description.`,
				},
				resource.Attribute{
					Name:        "enabled_time",
					Description: `The Physical Connection to Which the Activation Time.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The Expiration Time.`,
				},
				resource.Attribute{
					Name:        "has_reservation_data",
					Description: `HasReservationData.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Physical Connection.`,
				},
				resource.Attribute{
					Name:        "line_operator",
					Description: `Provides Access to the Physical Line Operator Value CT: China Telecom, CU: China Unicom, CM: china Mobile, CO: Other Chinese, Equinix:Equinix, Other: Other Overseas.`,
				},
				resource.Attribute{
					Name:        "loa_status",
					Description: `Loa State.`,
				},
				resource.Attribute{
					Name:        "payment_type",
					Description: `on Behalf of the Pay-as-You-Type of Resource Attribute Field.`,
				},
				resource.Attribute{
					Name:        "peer_location",
					Description: `and an on-Premises Data Center Location.`,
				},
				resource.Attribute{
					Name:        "physical_connection_id",
					Description: `on Behalf of the Resource Level Id of the Resources Property Fields.`,
				},
				resource.Attribute{
					Name:        "physical_connection_name",
					Description: `on Behalf of the Resource Name of the Resources-Attribute Field.`,
				},
				resource.Attribute{
					Name:        "port_number",
					Description: `To Connect a Device Port: The Port Number of.`,
				},
				resource.Attribute{
					Name:        "port_type",
					Description: `The Physical Leased Line Access Port Type Value 100Base-T: Fast Electrical Ports, 1000Base-T (the Default): gigabit Electrical Ports, 1000Base-LX: Gigabit Singlemode Optical Ports (10Km), 10GBase-T: Gigabit Electrical Port, 10GBase-LR: Gigabit Singlemode Optical Ports (10Km).`,
				},
				resource.Attribute{
					Name:        "redundant_physical_connection_id",
					Description: `Redundant Physical Connection to Which the ID.`,
				},
				resource.Attribute{
					Name:        "reservation_active_time",
					Description: `The Renewal of the Entry into Force of the Time.`,
				},
				resource.Attribute{
					Name:        "reservation_internet_charge_type",
					Description: `Renewal Type.`,
				},
				resource.Attribute{
					Name:        "reservation_order_type",
					Description: `Renewal Order Type.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `The Physical Connection to Which the Specifications.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Resources on Behalf of a State of the Resource Attribute Field.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Physical Private Line of Type. Default Value: VPC.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_express_connect_virtual_border_routers",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Express Connect Virtual Border Routers to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, ForceNew) Custom filter block as described below.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, ForceNew, Computed) A list of Virtual Border Router IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, ForceNew) A regex string to filter results by Virtual Border Router name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, ForceNew) The instance state with. Valid values: ` + "`" + `active` + "`" + `, ` + "`" + `deleting` + "`" + `, ` + "`" + `recovering` + "`" + `, ` + "`" + `terminated` + "`" + `, ` + "`" + `terminating` + "`" + `, ` + "`" + `unconfirmed` + "`" + `. ### Block filter More complex filters can be expressed using one or more ` + "`" + `filter` + "`" + ` sub-blocks, which take the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the field to filter by, as defined by`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Set of values that are accepted for the given field. ## Argument Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Virtual Border Router names.`,
				},
				resource.Attribute{
					Name:        "routers",
					Description: `A list of Express Connect Virtual Border Routers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "access_point_id",
					Description: `The physical leased line access point ID.`,
				},
				resource.Attribute{
					Name:        "activation_time",
					Description: `The first activation time of VBR.`,
				},
				resource.Attribute{
					Name:        "circuit_code",
					Description: `Operators for physical connection circuit provided coding.`,
				},
				resource.Attribute{
					Name:        "cloud_box_instance_id",
					Description: `Box Instance Id.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The representative of the creation time resources attribute field.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of VBR. Length is from 2 to 256 characters, must start with a letter or the Chinese at the beginning, but not at the http:// Or https:// at the beginning.`,
				},
				resource.Attribute{
					Name:        "detect_multiplier",
					Description: `Detection time multiplier that recipient allows the sender to send a message of the maximum allowable connections for the number of packets, used to detect whether the link normal. Value: 3~10.`,
				},
				resource.Attribute{
					Name:        "ecc_id",
					Description: `High Speed Migration Service Instance Id.`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `Whether to Enable IPv6.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Virtual Border Router.`,
				},
				resource.Attribute{
					Name:        "local_gateway_ip",
					Description: `Alibaba Cloud-Connected IPv4 address.`,
				},
				resource.Attribute{
					Name:        "local_ipv6_gateway_ip",
					Description: `Alibaba Cloud-Connected IPv6 Address.`,
				},
				resource.Attribute{
					Name:        "min_rx_interval",
					Description: `Configure BFD packet reception interval of values include: 200~1000, unit: ms.`,
				},
				resource.Attribute{
					Name:        "min_tx_interval",
					Description: `Configure BFD packet transmission interval maximum value: 200~1000, unit: ms.`,
				},
				resource.Attribute{
					Name:        "payment_vbr_expire_time",
					Description: `The Billing of the Extended Time.`,
				},
				resource.Attribute{
					Name:        "peer_gateway_ip",
					Description: `The Client-Side Interconnection IPv4 Address.`,
				},
				resource.Attribute{
					Name:        "peer_ipv6_gateway_ip",
					Description: `The Client-Side Interconnection IPv6 Address.`,
				},
				resource.Attribute{
					Name:        "peering_ipv6_subnet_mask",
					Description: `Alibaba Cloud-Connected IPv6 with Client-Side Interconnection IPv6 of Subnet Mask.`,
				},
				resource.Attribute{
					Name:        "peering_subnet_mask",
					Description: `Alibaba Cloud-Connected IPv4 and Client-Side Interconnection IPv4 of Subnet Mask.`,
				},
				resource.Attribute{
					Name:        "physical_connection_business_status",
					Description: `Physical Private Line Service Status Value Normal: Normal, financiallocked: If You Lock.`,
				},
				resource.Attribute{
					Name:        "physical_connection_id",
					Description: `The ID of the Physical Connection to Which the ID.`,
				},
				resource.Attribute{
					Name:        "physical_connection_owner_uid",
					Description: `Physical Private Line Where the Account ID.`,
				},
				resource.Attribute{
					Name:        "physical_connection_status",
					Description: `Physical Private Line State.`,
				},
				resource.Attribute{
					Name:        "recovery_time",
					Description: `The Last from a Terminated State to the Active State of the Time.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `Route Table ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The VBR state.`,
				},
				resource.Attribute{
					Name:        "termination_time",
					Description: `The Most Recent Was Aborted by the Time.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `VBR Type.`,
				},
				resource.Attribute{
					Name:        "virtual_border_router_id",
					Description: `The VBR ID.`,
				},
				resource.Attribute{
					Name:        "virtual_border_router_name",
					Description: `The name of VBR. Length is from 2 to 128 characters, must start with a letter or the Chinese at the beginning can contain numbers, the underscore character (_) and dash (-). But do not start with http:// or https:// at the beginning.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `The VLAN ID of the VBR. Value range: 0~2999.`,
				},
				resource.Attribute{
					Name:        "vlan_interface_id",
					Description: `The ID of the Router Interface.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_forward_entries",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Forward Entries owned by an Apsara Stack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Forward Entries IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by forward entry name.`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `(Optional) The public IP address.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `(Optional) The private IP address.`,
				},
				resource.Attribute{
					Name:        "forward_table_id",
					Description: `(Required) The ID of the Forward table.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Forward Entries IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Forward Entries names.`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `A list of Forward Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Forward Entry.`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `The public IP address.`,
				},
				resource.Attribute{
					Name:        "external_port",
					Description: `The public port.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The protocol type.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `The private IP address.`,
				},
				resource.Attribute{
					Name:        "internal_port",
					Description: `The private port.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Forward Entry.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Forward Entries IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Forward Entries names.`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `A list of Forward Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Forward Entry.`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `The public IP address.`,
				},
				resource.Attribute{
					Name:        "external_port",
					Description: `The public port.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The protocol type.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `The private IP address.`,
				},
				resource.Attribute{
					Name:        "internal_port",
					Description: `The private port.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Forward Entry.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_gpdb_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of AnalyticDB for PostgreSQL instances according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the instance name.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Instance availability zone.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) Used to retrieve instances belong to specified ` + "`" + `vswitch` + "`" + ` resources.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save the collection of instances after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `The ids list of AnalyticDB for PostgreSQL instances.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names list of AnalyticDB for PostgreSQL instance.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of AnalyticDB for PostgreSQL instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance id.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of an instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options are ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for yearly or monthly subscription.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Instance availability zone.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `The time when you create an instance. The format is YYYY-MM-DDThh:mm:ssZ, such as 2011-05-30T12:11:4Z.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database engine type. Supported option is ` + "`" + `gpdb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Database engine version.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Classic network or VPC.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `The group type.`,
				},
				resource.Attribute{
					Name:        "instance_group_count",
					Description: `The number of groups.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `The ids list of AnalyticDB for PostgreSQL instances.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names list of AnalyticDB for PostgreSQL instance.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of AnalyticDB for PostgreSQL instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance id.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of an instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options are ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for yearly or monthly subscription.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Instance availability zone.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `The time when you create an instance. The format is YYYY-MM-DDThh:mm:ssZ, such as 2011-05-30T12:11:4Z.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database engine type. Supported option is ` + "`" + `gpdb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Database engine version.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Classic network or VPC.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `The group type.`,
				},
				resource.Attribute{
					Name:        "instance_group_count",
					Description: `The number of groups.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_images",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of images available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting images by name.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional, type: bool) If more than one result are returned, select the most recent one.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) Filter results by a specific image owner. Valid items are ` + "`" + `system` + "`" + `, ` + "`" + `self` + "`" + `, ` + "`" + `others` + "`" + `, ` + "`" + `marketplace` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ->`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of image IDs.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `A list of images. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `Platform type of the image system: i386 or x86_64.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the image.`,
				},
				resource.Attribute{
					Name:        "image_owner_alias",
					Description: `Alias of the image owner.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `Display Chinese name of the OS.`,
				},
				resource.Attribute{
					Name:        "os_name_en",
					Description: `Display English name of the OS.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the image. Possible values: ` + "`" + `UnAvailable` + "`" + `, ` + "`" + `Available` + "`" + `, ` + "`" + `Creating` + "`" + ` and ` + "`" + `CreateFailed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the image.`,
				},
				resource.Attribute{
					Name:        "disk_device_mappings",
					Description: `Description of the system with disks and snapshots under the image.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device information of the created disk: such as /dev/xvdb.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the created disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Snapshot ID.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `Product code of the image on the image market.`,
				},
				resource.Attribute{
					Name:        "is_subscribed",
					Description: `Whether the user has subscribed to the terms of service for the image product corresponding to the ProductCode.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `Version of the image.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `Progress of image creation, presented in percentages.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of image IDs.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `A list of images. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `Platform type of the image system: i386 or x86_64.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the image.`,
				},
				resource.Attribute{
					Name:        "image_owner_alias",
					Description: `Alias of the image owner.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `Display Chinese name of the OS.`,
				},
				resource.Attribute{
					Name:        "os_name_en",
					Description: `Display English name of the OS.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the image. Possible values: ` + "`" + `UnAvailable` + "`" + `, ` + "`" + `Available` + "`" + `, ` + "`" + `Creating` + "`" + ` and ` + "`" + `CreateFailed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the image.`,
				},
				resource.Attribute{
					Name:        "disk_device_mappings",
					Description: `Description of the system with disks and snapshots under the image.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device information of the created disk: such as /dev/xvdb.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the created disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Snapshot ID.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `Product code of the image on the image market.`,
				},
				resource.Attribute{
					Name:        "is_subscribed",
					Description: `Whether the user has subscribed to the terms of service for the image product corresponding to the ProductCode.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `Version of the image.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `Progress of image creation, presented in percentages.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_instance_type_families",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ECS Instance Type Families to be used by the alibabacloudstack_instance resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, ForceNew) The Zone to launch the instance.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `(Optional) The generation of the instance type family,`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance type family IDs.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance type family.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The generation of the instance type family.`,
				},
				resource.Attribute{
					Name:        "zone_ids",
					Description: `A list of Zone to launch the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance type family IDs.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance type family.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The generation of the instance type family.`,
				},
				resource.Attribute{
					Name:        "zone_ids",
					Description: `A list of Zone to launch the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_instance_types",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ECS Instance Types to be used by the alibabacloudstack_instance resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The zone where instance types are supported.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `(Optional) Filter the results to a specific number of cpu cores.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `(Optional) Filter the results to a specific memory size in GB.`,
				},
				resource.Attribute{
					Name:        "gpu_amount",
					Description: `(Optional) The GPU amount of an instance type.`,
				},
				resource.Attribute{
					Name:        "gpu_spec",
					Description: `(Optional) The GPU spec of an instance type.`,
				},
				resource.Attribute{
					Name:        "instance_type_family",
					Description: `(Optional) Filter the results based on their family name. For example: 'ecs.n4'.`,
				},
				resource.Attribute{
					Name:        "eni_amount",
					Description: `(Optional) Filter the result whose network interface number is no more than ` + "`" + `eni_amount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kubernetes_node_role",
					Description: `(Optional) Filter the result which is used to create a kubernetes cluster Optional Values: ` + "`" + `Master` + "`" + ` and ` + "`" + `Worker` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance type IDs.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `A list of image types. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance type.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `Number of CPU cores.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Size of memory, measured in GB.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The instance type family.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `List of availability zones that support the instance type.`,
				},
				resource.Attribute{
					Name:        "gpu",
					Description: `The GPU attribution of an instance type:`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `The amount of GPU of an instance type.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The category of GPU of an instance type.`,
				},
				resource.Attribute{
					Name:        "burstable_instance",
					Description: `The burstable instance attribution:`,
				},
				resource.Attribute{
					Name:        "initial_credit",
					Description: `The initial CPU credit of a burstable instance.`,
				},
				resource.Attribute{
					Name:        "baseline_credit",
					Description: `The compute performance benchmark CPU credit of a burstable instance.`,
				},
				resource.Attribute{
					Name:        "eni_amount",
					Description: `The maximum number of network interfaces that an instance type can be attached to.`,
				},
				resource.Attribute{
					Name:        "local_storage",
					Description: `Local storage of an instance type:`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of a local storage in GB.`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `The number of local storage devices that an instance has been attached to.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The category of local storage that an instance has been attached to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance type IDs.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `A list of image types. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance type.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `Number of CPU cores.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Size of memory, measured in GB.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The instance type family.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `List of availability zones that support the instance type.`,
				},
				resource.Attribute{
					Name:        "gpu",
					Description: `The GPU attribution of an instance type:`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `The amount of GPU of an instance type.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The category of GPU of an instance type.`,
				},
				resource.Attribute{
					Name:        "burstable_instance",
					Description: `The burstable instance attribution:`,
				},
				resource.Attribute{
					Name:        "initial_credit",
					Description: `The initial CPU credit of a burstable instance.`,
				},
				resource.Attribute{
					Name:        "baseline_credit",
					Description: `The compute performance benchmark CPU credit of a burstable instance.`,
				},
				resource.Attribute{
					Name:        "eni_amount",
					Description: `The maximum number of network interfaces that an instance type can be attached to.`,
				},
				resource.Attribute{
					Name:        "local_storage",
					Description: `Local storage of an instance type:`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of a local storage in GB.`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `The number of local storage devices that an instance has been attached to.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The category of local storage that an instance has been attached to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ECS instances to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ECS instance IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by instance name.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The image ID of some ECS instance used.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Instance status. Valid values: "Creating", "Starting", "Running", "Stopping" and "Stopped". If undefined, all statuses are considered.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC linked to the instances.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) ID of the VSwitch linked to the instances.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability zone where instances are located.`,
				},
				resource.Attribute{
					Name:        "ram_role_name",
					Description: `(Optional, ForceNew) The RAM role name which the instance attaches.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of ECS instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instances names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance current status.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Instance name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Instance description.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Image ID the instance is using.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Instance private IP address.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `EIP address the VPC instance is using.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `List of security group IDs the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `Key pair the instance is using.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Instance creation time.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Max output bandwidth for internet.`,
				},
				resource.Attribute{
					Name:        "disk_device_mappings",
					Description: `Description of the attached disks.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device information of the created disk: such as /dev/xvdb.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the created disk.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Cloud disk category.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Cloud disk type: system disk or data disk.`,
				},
				resource.Attribute{
					Name:        "ram_role_name",
					Description: `The Ram role name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of ECS instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instances names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance current status.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Instance name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Instance description.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Image ID the instance is using.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Instance private IP address.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `EIP address the VPC instance is using.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `List of security group IDs the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `Key pair the instance is using.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Instance creation time.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Max output bandwidth for internet.`,
				},
				resource.Attribute{
					Name:        "disk_device_mappings",
					Description: `Description of the attached disks.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device information of the created disk: such as /dev/xvdb.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the created disk.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Cloud disk category.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Cloud disk type: system disk or data disk.`,
				},
				resource.Attribute{
					Name:        "ram_role_name",
					Description: `The Ram role name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_key_pairs",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available key pairs that can be used by an Alibabacloudstack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the resulting key pairs.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of key pair IDs.`,
				},
				resource.Attribute{
					Name:        "finger_print",
					Description: `(Optional) A finger print used to retrieve specified key pair.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of key pair names.`,
				},
				resource.Attribute{
					Name:        "key_pairs",
					Description: `A list of key pairs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the key pair.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `Name of the key pair.`,
				},
				resource.Attribute{
					Name:        "finger_print",
					Description: `Finger print of the key pair.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of ECS instances that has been bound this key pair.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of the availability zone where the ECS instance is located.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The name of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of the VSwitch attached to the ECS instance.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP address or EIP of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of key pair names.`,
				},
				resource.Attribute{
					Name:        "key_pairs",
					Description: `A list of key pairs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the key pair.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `Name of the key pair.`,
				},
				resource.Attribute{
					Name:        "finger_print",
					Description: `Finger print of the key pair.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of ECS instances that has been bound this key pair.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of the availability zone where the ECS instance is located.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The name of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of the VSwitch attached to the ECS instance.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP address or EIP of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kms_aliases",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available KMS Aliases.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of KMS aliases IDs. The value is same as KMS alias_name.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter the results by the KMS alias name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of kms aliases IDs. The value is same as KMS alias_name.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of KMS alias name.`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `A list of KMS User alias. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the alias. The value is same as KMS alias_name.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `ID of the key.`,
				},
				resource.Attribute{
					Name:        "alias_name",
					Description: `The unique identifier of the alias.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of kms aliases IDs. The value is same as KMS alias_name.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of KMS alias name.`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `A list of KMS User alias. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the alias. The value is same as KMS alias_name.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `ID of the key.`,
				},
				resource.Attribute{
					Name:        "alias_name",
					Description: `The unique identifier of the alias.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kms_ciphertext",
			Category:         "Data Sources",
			ShortDescription: `Encrypt data with KMS.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "plaintext",
					Description: `The plaintext to be encrypted which must be encoded in Base64.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `The globally unique ID of the CMK.`,
				},
				resource.Attribute{
					Name:        "encryption_context",
					Description: `(Optional) The Encryption context. If you specify this parameter here, it is also required when you call the Decrypt API operation. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ciphertext_blob",
					Description: `The ciphertext of the data key encrypted with the primary CMK version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ciphertext_blob",
					Description: `The ciphertext of the data key encrypted with the primary CMK version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kms_keys",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available KMS Keys.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of KMS key IDs.`,
				},
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional) A regex string to filter the results by the KMS key description.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Filter the results by status of the KMS keys. Valid values: ` + "`" + `Enabled` + "`" + `, ` + "`" + `Disabled` + "`" + `, ` + "`" + `PendingDeletion` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of KMS key IDs.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `A list of KMS keys. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the key.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The alibabacloudstack Cloud Resource Name (ARN) of the key.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the key.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the key. Possible values: ` + "`" + `Enabled` + "`" + `, ` + "`" + `Disabled` + "`" + ` and ` + "`" + `PendingDeletion` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of key.`,
				},
				resource.Attribute{
					Name:        "delete_date",
					Description: `Deletion date of key.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `The owner of the key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of KMS key IDs.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `A list of KMS keys. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the key.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The alibabacloudstack Cloud Resource Name (ARN) of the key.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the key.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the key. Possible values: ` + "`" + `Enabled` + "`" + `, ` + "`" + `Disabled` + "`" + ` and ` + "`" + `PendingDeletion` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of key.`,
				},
				resource.Attribute{
					Name:        "delete_date",
					Description: `Deletion date of key.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `The owner of the key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kms_secrets",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available KMS Secrets.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fetch_tags",
					Description: `(Optional) Whether to include the predetermined resource tag in the return value. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of KMS Secret ids. The value is same as KMS secret_name.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter the results by the KMS secret_name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Kms Secret ids. The value is same as KMS secret_name.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of KMS Secret names.`,
				},
				resource.Attribute{
					Name:        "secrets",
					Description: `A list of KMS Secrets. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Kms Secret. The value is same as KMS secret_name.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `Name of the KMS Secret.`,
				},
				resource.Attribute{
					Name:        "planned_delete_time",
					Description: `Schedule deletion time.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Kms Secret ids. The value is same as KMS secret_name.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of KMS Secret names.`,
				},
				resource.Attribute{
					Name:        "secrets",
					Description: `A list of KMS Secrets. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Kms Secret. The value is same as KMS secret_name.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `Name of the KMS Secret.`,
				},
				resource.Attribute{
					Name:        "planned_delete_time",
					Description: `Schedule deletion time.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kvstore_instance_classes",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of KVStore instacne classes info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The Zone to launch the KVStore instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Filter the results by charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PrePaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Optional) Database type. Options are ` + "`" + `Redis` + "`" + `, ` + "`" + `Memcache` + "`" + `. Default to ` + "`" + `Redis` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional) Database version required by the user. Value options of Redis can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/60873.htm) ` + "`" + `EngineVersion` + "`" + `. Value of Memcache should be empty.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `(Optional) The KVStore instance system architecture required by the user. Valid values: ` + "`" + `standard` + "`" + `, ` + "`" + `cluster` + "`" + ` and ` + "`" + `rwsplit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Optional) The KVStore instance node type required by the user. Valid values: ` + "`" + `double` + "`" + `, ` + "`" + `single` + "`" + `, ` + "`" + `readone` + "`" + `, ` + "`" + `readthree` + "`" + ` and ` + "`" + `readfive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform apply` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "edition_type",
					Description: `(Optional) The KVStore instance edition type required by the user. Valid values: ` + "`" + `Community` + "`" + ` and ` + "`" + `Enterprise` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "series_type",
					Description: `(Optional) The KVStore instance series type required by the user. Valid values: ` + "`" + `enhanced_performance_type` + "`" + ` and ` + "`" + `hybrid_storage` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "shard_number",
					Description: `(Optional) The number of shard.Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `16` + "`" + `, ` + "`" + `32` + "`" + `, ` + "`" + `64` + "`" + `, ` + "`" + `128` + "`" + `, ` + "`" + `256` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "instance_classes",
					Description: `A list of KVStore available instance classes.`,
				},
				resource.Attribute{
					Name:        "classes",
					Description: `A list of KVStore available instance classes when the ` + "`" + `sorted_by` + "`" + ` is "Price". include:`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `KVStore available instance class.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_classes",
					Description: `A list of KVStore available instance classes.`,
				},
				resource.Attribute{
					Name:        "classes",
					Description: `A list of KVStore available instance classes when the ` + "`" + `sorted_by` + "`" + ` is "Price". include:`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `KVStore available instance class.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kvstore_instance_engines",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of KVStore instacne engines info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The Zone to launch the KVStore instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Filter the results by charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PrePaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Optional) Database type. Options are ` + "`" + `Redis` + "`" + `, ` + "`" + `Memcache` + "`" + `. Default to ` + "`" + `Redis` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional) Database version required by the user. Value options of Redis can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/60873.htm) ` + "`" + `EngineVersion` + "`" + `. Value of Memcache should be empty.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform apply` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "instance_engines",
					Description: `A list of KVStore available instance engines. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The Zone to launch the KVStore instance.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database type.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `KVStore Instance version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_engines",
					Description: `A list of KVStore available instance engines. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The Zone to launch the KVStore instance.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database type.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `KVStore Instance version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kvstore_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of kvstore instances according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the instance name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of RKV instance IDs.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) Database type. Options are ` + "`" + `Memcache` + "`" + `, and ` + "`" + `Redis` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of the instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Query the instance bound to the tag. The format of the incoming value is ` + "`" + `json` + "`" + ` string, including ` + "`" + `TagKey` + "`" + ` and ` + "`" + `TagValue` + "`" + `. ` + "`" + `TagKey` + "`" + ` cannot be null, and ` + "`" + `TagValue` + "`" + ` can be empty. Format example ` + "`" + `{"key1":"value1"}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save the collection of instances after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of RKV instance IDs.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of RKV instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the RKV instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the RKV instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options: ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for subscription.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expiration time. Pay-As-You-Go instances are never expire.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) Database type. Options are ` + "`" + `Memcache` + "`" + `, and ` + "`" + `Redis` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VSwitch ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the instance.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Capacity of the applied ApsaraDB for Redis instance. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Instance bandwidth limit. Unit: Mbit/s.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `Instance connection quantity limit. Unit: count.`,
				},
				resource.Attribute{
					Name:        "connections_domain",
					Description: `Instance connection domain (only Intranet access supported).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Connection port of the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of RKV instance IDs.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of RKV instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the RKV instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the RKV instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options: ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for subscription.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expiration time. Pay-As-You-Go instances are never expire.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) Database type. Options are ` + "`" + `Memcache` + "`" + `, and ` + "`" + `Redis` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VSwitch ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the instance.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Capacity of the applied ApsaraDB for Redis instance. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Instance bandwidth limit. Unit: Mbit/s.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `Instance connection quantity limit. Unit: count.`,
				},
				resource.Attribute{
					Name:        "connections_domain",
					Description: `Instance connection domain (only Intranet access supported).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Connection port of the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_kvstore_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for KVStore that can be used by an Apsara Stack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "multi",
					Description: `(Optional) Indicate whether the zones can be used in a multi AZ configuration. Default to ` + "`" + `false` + "`" + `. Multi AZ is usually used to launch KVStore instances.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Filter the results by a specific instance charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_nat_gateways",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Nat Gateways owned by an Alibabacloudstack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of NAT gateways IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter nat gateways by name.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Nat gateways IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Nat gateways names.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `A list of Nat gateways. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `The specification of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "snat_table_id",
					Description: `The snat table id.`,
				},
				resource.Attribute{
					Name:        "forward_table_id",
					Description: `The forward table id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Nat gateways IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Nat gateways names.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `A list of Nat gateways. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `The specification of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "snat_table_id",
					Description: `The snat table id.`,
				},
				resource.Attribute{
					Name:        "forward_table_id",
					Description: `The forward table id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_network_acls",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Network Acls to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, ForceNew) A list of Network Acl ID.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, ForceNew) A regex string to filter results by Network Acl name.`,
				},
				resource.Attribute{
					Name:        "network_acl_name",
					Description: `(Optional, ForceNew) The name of the network ACL.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, ForceNew) The state of the network ACL. Valid values: ` + "`" + `Available` + "`" + ` and ` + "`" + `Modifying` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) The ID of the associated VPC.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Optional, ForceNew) The type of the associated resource. Valid values ` + "`" + `VSwitch` + "`" + `. ` + "`" + `resource_type` + "`" + ` and ` + "`" + `resource_id` + "`" + ` need to be specified at the same time to take effect.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Optional, ForceNew) The ID of the associated resource. ## Argument Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Network Acl names.`,
				},
				resource.Attribute{
					Name:        "acls",
					Description: `A list of Network Acls. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of network ACL information.`,
				},
				resource.Attribute{
					Name:        "egress_acl_entries",
					Description: `Output direction rule information.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Give the description information of the direction rule.`,
				},
				resource.Attribute{
					Name:        "destination_cidr_ip",
					Description: `The destination address segment.`,
				},
				resource.Attribute{
					Name:        "network_acl_entry_name",
					Description: `The name of the entry for the direction rule.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The authorization policy.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Destination port range.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Transport layer protocol.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Network Acl.`,
				},
				resource.Attribute{
					Name:        "ingress_acl_entries",
					Description: `Entry direction rule information.`,
				},
				resource.Attribute{
					Name:        "source_cidr_ip",
					Description: `The source address field.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the entry direction rule.`,
				},
				resource.Attribute{
					Name:        "network_acl_entry_name",
					Description: `The name of the entry direction rule entry.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The authorization policy.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Source port range.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Transport layer protocol.`,
				},
				resource.Attribute{
					Name:        "network_acl_id",
					Description: `The first ID of the resource.`,
				},
				resource.Attribute{
					Name:        "network_acl_name",
					Description: `The name of the network ACL.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `The associated resource.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The ID of the associated resource.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `The type of the associated resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The state of the associated resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The state of the network ACL.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the associated VPC.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_network_interfaces",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to get a list of elastic network interfaces according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "interfaces",
					Description: `A list of ENIs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ENI.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the ENI.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch that the ENI is linked to.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `ID of the availability zone that the ENI belongs to.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP of the ENI.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Primary private IP of the ENI.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `A list of secondary private IP address that is assigned to the ENI.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `A list of security group that the ENI belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the ENI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the ENI.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instance that the ENI is attached to.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the ENI.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the ENI.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ons_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ons groups available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the ONS Instance that owns the groups.`,
				},
				resource.Attribute{
					Name:        "group_id_regex",
					Description: `(Optional) A regex string to filter results by the group name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the group.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The ID of the group owner, which is the Apsara Stack Cloud UID.`,
				},
				resource.Attribute{
					Name:        "independent_naming",
					Description: `Indicates whether namespaces are available.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the group.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The ID of the group owner, which is the Apsara Stack Cloud UID.`,
				},
				resource.Attribute{
					Name:        "independent_naming",
					Description: `Indicates whether namespaces are available.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ons_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ons instances available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the instance name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `The status of the instance.`,
				},
				resource.Attribute{
					Name:        "release_time",
					Description: `The automatic release time of an Enterprise Platinum Edition instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `The status of the instance.`,
				},
				resource.Attribute{
					Name:        "release_time",
					Description: `The automatic release time of an Enterprise Platinum Edition instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_ons_topics",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ons topics available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the ONS Instance that owns the topics.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the topic name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of topic names.`,
				},
				resource.Attribute{
					Name:        "topics",
					Description: `A list of topics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `The name of the topic.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The ID of the topic owner, which is the Apsara Stack Cloud UID.`,
				},
				resource.Attribute{
					Name:        "relation",
					Description: `The relation ID.`,
				},
				resource.Attribute{
					Name:        "relation_name",
					Description: `The name of the relation, for example, owner, publishable, subscribable, and publishable and subscribable.`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `The type of the message.`,
				},
				resource.Attribute{
					Name:        "independent_naming",
					Description: `Indicates whether namespaces are available.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the topic.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of topic names.`,
				},
				resource.Attribute{
					Name:        "topics",
					Description: `A list of topics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `The name of the topic.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The ID of the topic owner, which is the Apsara Stack Cloud UID.`,
				},
				resource.Attribute{
					Name:        "relation",
					Description: `The relation ID.`,
				},
				resource.Attribute{
					Name:        "relation_name",
					Description: `The name of the relation, for example, owner, publishable, subscribable, and publishable and subscribable.`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `The type of the message.`,
				},
				resource.Attribute{
					Name:        "independent_naming",
					Description: `Indicates whether namespaces are available.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the topic.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_oos_executions",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of OOS Executions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of OOS Execution ids.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) The category of template. Valid: ` + "`" + `AlarmTrigger` + "`" + `, ` + "`" + `EventTrigger` + "`" + `, ` + "`" + `Other` + "`" + ` and ` + "`" + `TimerTrigger` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `(Optional) The time when the execution was ended.`,
				},
				resource.Attribute{
					Name:        "end_date_after",
					Description: `(Optional) Execution whose end time is less than or equal to the specified time.`,
				},
				resource.Attribute{
					Name:        "executed_by",
					Description: `(Optional) The user who execute the template.`,
				},
				resource.Attribute{
					Name:        "include_child_execution",
					Description: `(Optional) Whether to include sub-execution.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The mode of OOS Execution. Valid: ` + "`" + `Automatic` + "`" + `, ` + "`" + `Debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "parent_execution_id",
					Description: `(Optional) The id of parent OOS Execution.`,
				},
				resource.Attribute{
					Name:        "ram_role",
					Description: `(Optional) The role that executes the current template.`,
				},
				resource.Attribute{
					Name:        "sort_field",
					Description: `(Optional) The sort field.`,
				},
				resource.Attribute{
					Name:        "sort_order",
					Description: `(Optional) The sort order.`,
				},
				resource.Attribute{
					Name:        "start_date_after",
					Description: `(Optional) The execution whose start time is greater than or equal to the specified time.`,
				},
				resource.Attribute{
					Name:        "start_date_before",
					Description: `(Optional) The execution with start time less than or equal to the specified time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The Status of OOS Execution. Valid: ` + "`" + `Cancelled` + "`" + `, ` + "`" + `Failed` + "`" + `, ` + "`" + `Queued` + "`" + `, ` + "`" + `Running` + "`" + `, ` + "`" + `Started` + "`" + `, ` + "`" + `Success` + "`" + `, ` + "`" + `Waiting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Optional) The name of execution template.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of OOS Execution ids.`,
				},
				resource.Attribute{
					Name:        "executions",
					Description: `A list of OOS Executions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the OOS Executions.`,
				},
				resource.Attribute{
					Name:        "counters",
					Description: `The counters of OOS Execution.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The time when the execution was created.`,
				},
				resource.Attribute{
					Name:        "execution_id",
					Description: `ID of the OOS Executions.`,
				},
				resource.Attribute{
					Name:        "is_parent",
					Description: `Whether to include subtasks.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `The outputs of OOS Executions.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `The parameters required by the template`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `The time when the template was started.`,
				},
				resource.Attribute{
					Name:        "status_message",
					Description: `The message of status.`,
				},
				resource.Attribute{
					Name:        "status_reason",
					Description: `The reason of status.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The id of execution template.`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `The version of execution template.`,
				},
				resource.Attribute{
					Name:        "update_date",
					Description: `The time when the template was updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of OOS Execution ids.`,
				},
				resource.Attribute{
					Name:        "executions",
					Description: `A list of OOS Executions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the OOS Executions.`,
				},
				resource.Attribute{
					Name:        "counters",
					Description: `The counters of OOS Execution.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The time when the execution was created.`,
				},
				resource.Attribute{
					Name:        "execution_id",
					Description: `ID of the OOS Executions.`,
				},
				resource.Attribute{
					Name:        "is_parent",
					Description: `Whether to include subtasks.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `The outputs of OOS Executions.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `The parameters required by the template`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `The time when the template was started.`,
				},
				resource.Attribute{
					Name:        "status_message",
					Description: `The message of status.`,
				},
				resource.Attribute{
					Name:        "status_reason",
					Description: `The reason of status.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The id of execution template.`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `The version of execution template.`,
				},
				resource.Attribute{
					Name:        "update_date",
					Description: `The time when the template was updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_oos_templates",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of OOS Templates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of OOS Template ids. Each element in the list is same as template_name.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter the results by the template_name.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) The category of template.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `(Optional) The creator of the template.`,
				},
				resource.Attribute{
					Name:        "created_date",
					Description: `(Optional) The template whose creation time is less than or equal to the specified time. The format is: YYYY-MM-DDThh:mm::ssZ.`,
				},
				resource.Attribute{
					Name:        "has_trigger",
					Description: `(Optional) Is it triggered successfully.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `(Optional) The sharing type of the template. Valid values: ` + "`" + `Private` + "`" + `, ` + "`" + `Public` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "template_format",
					Description: `(Optional) The format of the template. Valid values: ` + "`" + `JSON` + "`" + `, ` + "`" + `YAML` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `(Optional) The type of OOS Template.`,
				},
				resource.Attribute{
					Name:        "sort_field",
					Description: `(Optional) Sort field. Valid values: ` + "`" + `TotalExecutionCount` + "`" + `, ` + "`" + `Popularity` + "`" + `, ` + "`" + `TemplateName` + "`" + ` and ` + "`" + `CreatedDate` + "`" + `. Default to ` + "`" + `TotalExecutionCount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sort_order",
					Description: `(Optional) Sort order. Valid values: ` + "`" + `Ascending` + "`" + `, ` + "`" + `Descending` + "`" + `. Default to ` + "`" + `Descending` + "`" + ``,
				},
				resource.Attribute{
					Name:        "created_date_after",
					Description: `(Optional) Create a template whose time is greater than or equal to the specified time. The format is: YYYY-MM-DDThh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of OOS Template ids. Each element in the list is same as template_name.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(Available in v1.114.0+) A list of OOS Template names.`,
				},
				resource.Attribute{
					Name:        "templates",
					Description: `A list of OOS Templates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the OOS Template. The value is same as template_name.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `Name of the OOS Template.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the OOS Template.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `ID of the OOS Template resource.`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `Version of the OOS Template.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `The user who updated the template.`,
				},
				resource.Attribute{
					Name:        "updated_date",
					Description: `The time when the template was updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of OOS Template ids. Each element in the list is same as template_name.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(Available in v1.114.0+) A list of OOS Template names.`,
				},
				resource.Attribute{
					Name:        "templates",
					Description: `A list of OOS Templates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the OOS Template. The value is same as template_name.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `Name of the OOS Template.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the OOS Template.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `ID of the OOS Template resource.`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `Version of the OOS Template.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `The user who updated the template.`,
				},
				resource.Attribute{
					Name:        "updated_date",
					Description: `The time when the template was updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_oss_bucket_objects",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of bucket objects to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket_name",
					Description: `Name of the bucket that contains the objects to find.`,
				},
				resource.Attribute{
					Name:        "key_regex",
					Description: `(Optional) A regex string to filter results by key.`,
				},
				resource.Attribute{
					Name:        "key_prefix",
					Description: `(Optional) Filter results by the given key prefix (such as "path/to/folder/logs-").`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "objects",
					Description: `A list of bucket objects. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Object key.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `Object access control list. Possible values: ` + "`" + `default` + "`" + `, ` + "`" + `private` + "`" + `, ` + "`" + `public-read` + "`" + ` and ` + "`" + `public-read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `Standard MIME type describing the format of the object data, e.g. "application/octet-stream".`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Caching behavior along the request/reply chain. Read [RFC2616 Cache-Control](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Presentational information for the object. Read [RFC2616 Content-Disposition](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Content encodings that have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [RFC2616 Content-Encoding](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_md5",
					Description: `MD5 value of the content. Read [MD5](https://www.alibabacloud.com/help/doc-detail/31978.htm) for computing method.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `Expiration date for the the request/response. Read [RFC2616 Expires](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `Server-side encryption of the object in OSS. It can be empty or ` + "`" + `AES256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `If present, specifies the ID of the Key Management Service(KMS) master encryption key that was used for the object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Object storage type. Possible values: ` + "`" + `Standard` + "`" + `, ` + "`" + `IA` + "`" + ` and ` + "`" + `Archive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_modification_time",
					Description: `Last modification time of the object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "objects",
					Description: `A list of bucket objects. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Object key.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `Object access control list. Possible values: ` + "`" + `default` + "`" + `, ` + "`" + `private` + "`" + `, ` + "`" + `public-read` + "`" + ` and ` + "`" + `public-read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `Standard MIME type describing the format of the object data, e.g. "application/octet-stream".`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Caching behavior along the request/reply chain. Read [RFC2616 Cache-Control](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Presentational information for the object. Read [RFC2616 Content-Disposition](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Content encodings that have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [RFC2616 Content-Encoding](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_md5",
					Description: `MD5 value of the content. Read [MD5](https://www.alibabacloud.com/help/doc-detail/31978.htm) for computing method.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `Expiration date for the the request/response. Read [RFC2616 Expires](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `Server-side encryption of the object in OSS. It can be empty or ` + "`" + `AES256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `If present, specifies the ID of the Key Management Service(KMS) master encryption key that was used for the object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Object storage type. Possible values: ` + "`" + `Standard` + "`" + `, ` + "`" + `IA` + "`" + ` and ` + "`" + `Archive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_modification_time",
					Description: `Last modification time of the object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_oss_buckets",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of OSS buckets to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by bucket name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of bucket names.`,
				},
				resource.Attribute{
					Name:        "buckets",
					Description: `A list of buckets. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Bucket name.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `Bucket access control list. Possible values: ` + "`" + `private` + "`" + `, ` + "`" + `public-read` + "`" + ` and ` + "`" + `public-read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extranet_endpoint",
					Description: `Internet domain name for accessing the bucket from outside.`,
				},
				resource.Attribute{
					Name:        "intranet_endpoint",
					Description: `Intranet domain name for accessing the bucket from an ECS instance in the same region.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Region of the data center where the bucket is located.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Bucket owner.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Object storage type. Possible values: ` + "`" + `Standard` + "`" + `, ` + "`" + `IA` + "`" + ` and ` + "`" + `Archive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Bucket creation date.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of bucket names.`,
				},
				resource.Attribute{
					Name:        "buckets",
					Description: `A list of buckets. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Bucket name.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `Bucket access control list. Possible values: ` + "`" + `private` + "`" + `, ` + "`" + `public-read` + "`" + ` and ` + "`" + `public-read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extranet_endpoint",
					Description: `Internet domain name for accessing the bucket from outside.`,
				},
				resource.Attribute{
					Name:        "intranet_endpoint",
					Description: `Intranet domain name for accessing the bucket from an ECS instance in the same region.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Region of the data center where the bucket is located.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Bucket owner.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Object storage type. Possible values: ` + "`" + `Standard` + "`" + `, ` + "`" + `IA` + "`" + ` and ` + "`" + `Archive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Bucket creation date.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_route_entries",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Route Entries owned by an Alibabacloudstack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required, ForceNew) The ID of the router table to which the route entry belongs.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) The instance ID of the next hop.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the route entry.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) The destination CIDR block of the route entry.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `A list of Route Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the route entry.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `The type of the next hop.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the route entry.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance ID of the next hop.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the router table to which the route entry belongs.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The destination CIDR block of the route entry.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "entries",
					Description: `A list of Route Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the route entry.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `The type of the next hop.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the route entry.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance ID of the next hop.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the router table to which the route entry belongs.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The destination CIDR block of the route entry.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_route_tables",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Route Tables owned by an Alibabacloudstack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Route Tables IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter route tables by name.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Vpc id of the route table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Route Tables IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Route Tables names.`,
				},
				resource.Attribute{
					Name:        "tables",
					Description: `A list of Route Tables. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Route Table.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `Router Id of the route table.`,
				},
				resource.Attribute{
					Name:        "route_table_type",
					Description: `The type of route table.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the route table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the route table instance.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Route Tables IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Route Tables names.`,
				},
				resource.Attribute{
					Name:        "tables",
					Description: `A list of Route Tables. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Route Table.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `Router Id of the route table.`,
				},
				resource.Attribute{
					Name:        "route_table_type",
					Description: `The type of route table.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the route table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the route table instance.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_router_interfaces",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of router interfaces to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string used to filter by router interface name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Expected status. Valid values are ` + "`" + `Active` + "`" + `, ` + "`" + `Inactive` + "`" + ` and ` + "`" + `Idle` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `(Optional) Specification of the link, such as ` + "`" + `Small.1` + "`" + ` (10Mb), ` + "`" + `Middle.1` + "`" + ` (100Mb), ` + "`" + `Large.2` + "`" + ` (2Gb), ...etc.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Optional) ID of the VRouter located in the local region.`,
				},
				resource.Attribute{
					Name:        "router_type",
					Description: `(Optional) Router type in the local region. Valid values are ` + "`" + `VRouter` + "`" + ` and ` + "`" + `VBR` + "`" + ` (physical connection).`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) Role of the router interface. Valid values are ` + "`" + `InitiatingSide` + "`" + ` (connection initiator) and ` + "`" + `AcceptingSide` + "`" + ` (connection receiver). The value of this parameter must be ` + "`" + `InitiatingSide` + "`" + ` if the ` + "`" + `router_type` + "`" + ` is set to ` + "`" + `VBR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_id",
					Description: `(Optional) ID of the peer router interface.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_owner_id",
					Description: `(Optional) Account ID of the owner of the peer router interface.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of router interface IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of router interface IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of router interface names.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `A list of router interfaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Router interface ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Router interface status. Possible values: ` + "`" + `Active` + "`" + `, ` + "`" + `Inactive` + "`" + ` and ` + "`" + `Idle` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Router interface name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Router interface description.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Router interface role. Possible values: ` + "`" + `InitiatingSide` + "`" + ` and ` + "`" + `AcceptingSide` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `Router interface specification. Possible values: ` + "`" + `Small.1` + "`" + `, ` + "`" + `Middle.1` + "`" + `, ` + "`" + `Large.2` + "`" + `, ...etc.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `ID of the VRouter located in the local region.`,
				},
				resource.Attribute{
					Name:        "router_type",
					Description: `Router type in the local region. Possible values: ` + "`" + `VRouter` + "`" + ` and ` + "`" + `VBR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC that owns the router in the local region.`,
				},
				resource.Attribute{
					Name:        "access_point_id",
					Description: `ID of the access point used by the VBR.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Router interface creation time.`,
				},
				resource.Attribute{
					Name:        "opposite_region_id",
					Description: `Peer router region ID.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_id",
					Description: `Peer router interface ID.`,
				},
				resource.Attribute{
					Name:        "opposite_router_id",
					Description: `Peer router ID.`,
				},
				resource.Attribute{
					Name:        "opposite_router_type",
					Description: `Router type in the peer region. Possible values: ` + "`" + `VRouter` + "`" + ` and ` + "`" + `VBR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_owner_id",
					Description: `Account ID of the owner of the peer router interface.`,
				},
				resource.Attribute{
					Name:        "health_check_source_ip",
					Description: `Source IP address used to perform health check on the physical connection.`,
				},
				resource.Attribute{
					Name:        "health_check_target_ip",
					Description: `Destination IP address used to perform health check on the physical connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of router interface IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of router interface names.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `A list of router interfaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Router interface ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Router interface status. Possible values: ` + "`" + `Active` + "`" + `, ` + "`" + `Inactive` + "`" + ` and ` + "`" + `Idle` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Router interface name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Router interface description.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Router interface role. Possible values: ` + "`" + `InitiatingSide` + "`" + ` and ` + "`" + `AcceptingSide` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `Router interface specification. Possible values: ` + "`" + `Small.1` + "`" + `, ` + "`" + `Middle.1` + "`" + `, ` + "`" + `Large.2` + "`" + `, ...etc.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `ID of the VRouter located in the local region.`,
				},
				resource.Attribute{
					Name:        "router_type",
					Description: `Router type in the local region. Possible values: ` + "`" + `VRouter` + "`" + ` and ` + "`" + `VBR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC that owns the router in the local region.`,
				},
				resource.Attribute{
					Name:        "access_point_id",
					Description: `ID of the access point used by the VBR.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Router interface creation time.`,
				},
				resource.Attribute{
					Name:        "opposite_region_id",
					Description: `Peer router region ID.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_id",
					Description: `Peer router interface ID.`,
				},
				resource.Attribute{
					Name:        "opposite_router_id",
					Description: `Peer router ID.`,
				},
				resource.Attribute{
					Name:        "opposite_router_type",
					Description: `Router type in the peer region. Possible values: ` + "`" + `VRouter` + "`" + ` and ` + "`" + `VBR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_owner_id",
					Description: `Account ID of the owner of the peer router interface.`,
				},
				resource.Attribute{
					Name:        "health_check_source_ip",
					Description: `Source IP address used to perform health check on the physical connection.`,
				},
				resource.Attribute{
					Name:        "health_check_target_ip",
					Description: `Destination IP address used to perform health check on the physical connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_security_group_rules",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of Security Group Rules available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The ID of the security group that owns the rules.`,
				},
				resource.Attribute{
					Name:        "nic_type",
					Description: `(Optional) Refers to the network type. Can be either ` + "`" + `internet` + "`" + ` or ` + "`" + `intranet` + "`" + `. The default value is ` + "`" + `internet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Authorization direction. Valid values are: ` + "`" + `ingress` + "`" + ` or ` + "`" + `egress` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) The IP protocol. Valid values are: ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `gre` + "`" + ` and ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) Authorization policy. Can be either ` + "`" + `accept` + "`" + ` or ` + "`" + `drop` + "`" + `. The default value is ` + "`" + `accept` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the security group that owns the rules.`,
				},
				resource.Attribute{
					Name:        "group_desc",
					Description: `The description of the security group that owns the rules.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of security group rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The protocol. Can be ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `gre` + "`" + ` or ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `The range of port numbers.`,
				},
				resource.Attribute{
					Name:        "source_cidr_ip",
					Description: `Source IP address segment for ingress authorization.`,
				},
				resource.Attribute{
					Name:        "source_group_owner_account",
					Description: `Alibabacloudstack Cloud account of the source security group.`,
				},
				resource.Attribute{
					Name:        "dest_cidr_ip",
					Description: `Target IP address segment for egress authorization.`,
				},
				resource.Attribute{
					Name:        "dest_group_owner_account",
					Description: `Alibabacloudstack Cloud account of the target security group.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `Authorization policy. Can be either ` + "`" + `accept` + "`" + ` or ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "nic_type",
					Description: `Network type, ` + "`" + `internet` + "`" + ` or ` + "`" + `intranet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Rule priority.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `Authorization direction, ` + "`" + `ingress` + "`" + ` or ` + "`" + `egress` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rules",
					Description: `A list of rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the security group that owns the rules.`,
				},
				resource.Attribute{
					Name:        "group_desc",
					Description: `The description of the security group that owns the rules.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of security group rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The protocol. Can be ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `gre` + "`" + ` or ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `The range of port numbers.`,
				},
				resource.Attribute{
					Name:        "source_cidr_ip",
					Description: `Source IP address segment for ingress authorization.`,
				},
				resource.Attribute{
					Name:        "source_group_owner_account",
					Description: `Alibabacloudstack Cloud account of the source security group.`,
				},
				resource.Attribute{
					Name:        "dest_cidr_ip",
					Description: `Target IP address segment for egress authorization.`,
				},
				resource.Attribute{
					Name:        "dest_group_owner_account",
					Description: `Alibabacloudstack Cloud account of the target security group.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `Authorization policy. Can be either ` + "`" + `accept` + "`" + ` or ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "nic_type",
					Description: `Network type, ` + "`" + `internet` + "`" + ` or ` + "`" + `intranet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Rule priority.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `Authorization direction, ` + "`" + `ingress` + "`" + ` or ` + "`" + `egress` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_security_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Security Groups available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Security Group IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter the resulting security groups by their names.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Used to retrieve security groups that belong to the specified VPC ID.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A map of tags assigned to the ECS instances. It must be in the format: ` + "`" + `` + "`" + `` + "`" + ` data "alibabacloudstack_security_groups" "taggedSecurityGroups" { tags = { tagKey1 = "tagValue1", tagKey2 = "tagValue2" } } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Security Group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Security Group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of Security Groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC that owns the security group.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the ECS instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Security Group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Security Group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of Security Groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC that owns the security group.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the ECS instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_acls",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of server load balancer acls (access control lists) to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of acls IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by acl name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew) The Id of resource group which acl belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB acls IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB acls names.`,
				},
				resource.Attribute{
					Name:        "acls",
					Description: `A list of SLB acls. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Acl ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Acl name.`,
				},
				resource.Attribute{
					Name:        "entry_list",
					Description: `A list of entry (IP addresses or CIDR blocks). Each entry contains two sub-fields as ` + "`" + `Entry Block` + "`" + ` follows.`,
				},
				resource.Attribute{
					Name:        "related_listeners",
					Description: `A list of listener are attached by the acl. Each listener contains four sub-fields as ` + "`" + `Listener Block` + "`" + ` follows.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `Resource group ID. ## Entry Block The entry mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `the comment of the entry. ## Listener Block The Listener mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `the id of load balancer instance, the listener belongs to.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `the listener port.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB acls IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB acls names.`,
				},
				resource.Attribute{
					Name:        "acls",
					Description: `A list of SLB acls. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Acl ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Acl name.`,
				},
				resource.Attribute{
					Name:        "entry_list",
					Description: `A list of entry (IP addresses or CIDR blocks). Each entry contains two sub-fields as ` + "`" + `Entry Block` + "`" + ` follows.`,
				},
				resource.Attribute{
					Name:        "related_listeners",
					Description: `A list of listener are attached by the acl. Each listener contains four sub-fields as ` + "`" + `Listener Block` + "`" + ` follows.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `Resource group ID. ## Entry Block The entry mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `the comment of the entry. ## Listener Block The Listener mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `the id of load balancer instance, the listener belongs to.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `the listener port.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_backend_servers",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of server load balancer backend servers to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) ID of the SLB with attachments.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) List of attached ECS instance IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "backend_servers",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `backend server ID.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight associated to the ECS instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "backend_servers",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `backend server ID.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight associated to the ECS instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_ca_certificates",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of slb CA certificates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ca certificates IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by ca certificate name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB ca certificates IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB ca certificates names.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A list of SLB ca certificates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `CA certificate ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `CA certificate name.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `CA certificate fingerprint.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `CA certificate created time.`,
				},
				resource.Attribute{
					Name:        "created_timestamp",
					Description: `CA certificate created timestamp.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The region Id of CA certificate.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB ca certificates IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB ca certificates names.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A list of SLB ca certificates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `CA certificate ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `CA certificate name.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `CA certificate fingerprint.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `CA certificate created time.`,
				},
				resource.Attribute{
					Name:        "created_timestamp",
					Description: `CA certificate created timestamp.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The region Id of CA certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_domain_extensions",
			Category:         "Data Sources",
			ShortDescription: `Provides a Load Banlancer domain extension Resource and add it to one Listener.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) IDs of the SLB domain extensions.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) The ID of the SLB instance.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `(Required) The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `A list of SLB domain extension. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the domain extension.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The ID of the certificate used by the domain name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "extensions",
					Description: `A list of SLB domain extension. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the domain extension.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The ID of the certificate used by the domain name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_listeners",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of server load balancer listeners to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) ID of the SLB with listeners.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Filter listeners by the specified protocol. Valid values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `(Optional) Filter listeners by the specified frontend port.`,
				},
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional) A regex string to filter results by SLB listener description.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "slb_listeners",
					Description: `A list of SLB listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `Frontend port used to receive incoming traffic and distribute it to the backend servers.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `Port opened on the backend server to receive requests.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Listener protocol. Possible values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Listener status.`,
				},
				resource.Attribute{
					Name:        "security_status",
					Description: `Security status. Only available when the protocol is ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Peak bandwidth. If the value is set to -1, the listener is not limited by bandwidth.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Algorithm used to distribute traffic. Possible values: ` + "`" + `wrr` + "`" + ` (weighted round robin), ` + "`" + `wlc` + "`" + ` (weighted least connection) and ` + "`" + `rr` + "`" + ` (round robin).`,
				},
				resource.Attribute{
					Name:        "server_group_id",
					Description: `ID of the linked VServer group.`,
				},
				resource.Attribute{
					Name:        "master_slave_server_group_id",
					Description: `ID of the active/standby server group.`,
				},
				resource.Attribute{
					Name:        "persistence_timeout",
					Description: `Timeout value of the TCP connection in seconds. If the value is 0, the session persistence function is disabled. Only available when the protocol is ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "established_timeout",
					Description: `Connection timeout in seconds for the Layer 4 TCP listener. Only available when the protocol is ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sticky_session",
					Description: `Indicate whether session persistence is enabled or not. If enabled, all session requests from the same client are sent to the same backend server. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sticky_session_type",
					Description: `Method used to handle the cookie. Possible values are ` + "`" + `insert` + "`" + ` (cookie added to the response) and ` + "`" + `server` + "`" + ` (cookie set by the backend server). Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + ` and sticky_session is ` + "`" + `on` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `Cookie timeout in seconds. Only available when the sticky_session_type is ` + "`" + `insert` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `Cookie configured by the backend server. Only available when the sticky_session_type is ` + "`" + `server` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Indicate whether health check is enabled of not. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `Health check method. Possible values are ` + "`" + `tcp` + "`" + ` and ` + "`" + `http` + "`" + `. Only available when the protocol is ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_domain",
					Description: `Domain name used for health check. The SLB sends HTTP head requests to the backend server, the domain is useful when the backend server verifies the host field in the requests. Only available when the protocol is ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `tcp` + "`" + ` (in this case health_check_type must be ` + "`" + `http` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "health_check_uri",
					Description: `URI used for health check. Only available when the protocol is ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `tcp` + "`" + ` (in this case health_check_type must be ` + "`" + `http` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "health_check_connect_port",
					Description: `Port used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_connect_timeout",
					Description: `Amount of time in seconds to wait for the response for a health check.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `Number of consecutive successes of health check performed on the same ECS instance (from failure to success).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `Number of consecutive failures of health check performed on the same ECS instance (from success to failure).`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `Amount of time in seconds to wait for the response from a health check. If an ECS instance sends no response within the specified timeout period, the health check fails. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `Time interval between two consecutive health checks.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `HTTP status codes indicating that the health check is normal. It can contain several comma-separated values such as "http_2xx,http_3xx". Only available when the protocol is ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `tcp` + "`" + ` (in this case health_check_type must be ` + "`" + `http` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "gzip",
					Description: `Indicate whether Gzip compression is enabled or not. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For" is added or not; it allows the backend server to know about the user's IP address. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for_slb_ip",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For_SLBIP" is added or not; it allows the backend server to know about the SLB IP address. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for_slb_id",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For_SLBID" is added or not; it allows the backend server to know about the SLB ID. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for_slb_proto",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For_proto" is added or not; it allows the backend server to know about the user's protocol. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of slb listener.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "slb_listeners",
					Description: `A list of SLB listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `Frontend port used to receive incoming traffic and distribute it to the backend servers.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `Port opened on the backend server to receive requests.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Listener protocol. Possible values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Listener status.`,
				},
				resource.Attribute{
					Name:        "security_status",
					Description: `Security status. Only available when the protocol is ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Peak bandwidth. If the value is set to -1, the listener is not limited by bandwidth.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Algorithm used to distribute traffic. Possible values: ` + "`" + `wrr` + "`" + ` (weighted round robin), ` + "`" + `wlc` + "`" + ` (weighted least connection) and ` + "`" + `rr` + "`" + ` (round robin).`,
				},
				resource.Attribute{
					Name:        "server_group_id",
					Description: `ID of the linked VServer group.`,
				},
				resource.Attribute{
					Name:        "master_slave_server_group_id",
					Description: `ID of the active/standby server group.`,
				},
				resource.Attribute{
					Name:        "persistence_timeout",
					Description: `Timeout value of the TCP connection in seconds. If the value is 0, the session persistence function is disabled. Only available when the protocol is ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "established_timeout",
					Description: `Connection timeout in seconds for the Layer 4 TCP listener. Only available when the protocol is ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sticky_session",
					Description: `Indicate whether session persistence is enabled or not. If enabled, all session requests from the same client are sent to the same backend server. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sticky_session_type",
					Description: `Method used to handle the cookie. Possible values are ` + "`" + `insert` + "`" + ` (cookie added to the response) and ` + "`" + `server` + "`" + ` (cookie set by the backend server). Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + ` and sticky_session is ` + "`" + `on` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `Cookie timeout in seconds. Only available when the sticky_session_type is ` + "`" + `insert` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `Cookie configured by the backend server. Only available when the sticky_session_type is ` + "`" + `server` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Indicate whether health check is enabled of not. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `Health check method. Possible values are ` + "`" + `tcp` + "`" + ` and ` + "`" + `http` + "`" + `. Only available when the protocol is ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_domain",
					Description: `Domain name used for health check. The SLB sends HTTP head requests to the backend server, the domain is useful when the backend server verifies the host field in the requests. Only available when the protocol is ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `tcp` + "`" + ` (in this case health_check_type must be ` + "`" + `http` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "health_check_uri",
					Description: `URI used for health check. Only available when the protocol is ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `tcp` + "`" + ` (in this case health_check_type must be ` + "`" + `http` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "health_check_connect_port",
					Description: `Port used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_connect_timeout",
					Description: `Amount of time in seconds to wait for the response for a health check.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `Number of consecutive successes of health check performed on the same ECS instance (from failure to success).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `Number of consecutive failures of health check performed on the same ECS instance (from success to failure).`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `Amount of time in seconds to wait for the response from a health check. If an ECS instance sends no response within the specified timeout period, the health check fails. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `Time interval between two consecutive health checks.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `HTTP status codes indicating that the health check is normal. It can contain several comma-separated values such as "http_2xx,http_3xx". Only available when the protocol is ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `tcp` + "`" + ` (in this case health_check_type must be ` + "`" + `http` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "gzip",
					Description: `Indicate whether Gzip compression is enabled or not. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For" is added or not; it allows the backend server to know about the user's IP address. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for_slb_ip",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For_SLBIP" is added or not; it allows the backend server to know about the SLB IP address. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for_slb_id",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For_SLBID" is added or not; it allows the backend server to know about the SLB ID. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for_slb_proto",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For_proto" is added or not; it allows the backend server to know about the user's protocol. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of slb listener.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_master_slave_server_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of master slave server groups related to a server load balancer to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `ID of the SLB.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of master slave server group IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by master slave server group name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB master slave server groups IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB master slave server groups names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of SLB master slave server groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `master slave server group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `master slave server group name.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `ECS instances associated to the group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the attached ECS instance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight associated to the ECS instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port used by the master slave server group.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `The server type of the attached ECS instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB master slave server groups IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB master slave server groups names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of SLB master slave server groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `master slave server group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `master slave server group name.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `ECS instances associated to the group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the attached ECS instance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight associated to the ECS instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port used by the master slave server group.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `The server type of the attached ECS instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_rules",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of server load balancer rules to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) ID of the SLB with listener rules.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `(Required) SLB listener port.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of rules IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by rule name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB listener rules IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB listener rules names.`,
				},
				resource.Attribute{
					Name:        "slb_rules",
					Description: `A list of SLB listener rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Rule ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Rule name.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name in the HTTP request where the rule applies (e.g. "`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Path in the HTTP request where the rule applies (e.g. "/image").`,
				},
				resource.Attribute{
					Name:        "server_group_id",
					Description: `ID of the linked VServer group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB listener rules IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB listener rules names.`,
				},
				resource.Attribute{
					Name:        "slb_rules",
					Description: `A list of SLB listener rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Rule ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Rule name.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name in the HTTP request where the rule applies (e.g. "`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Path in the HTTP request where the rule applies (e.g. "/image").`,
				},
				resource.Attribute{
					Name:        "server_group_id",
					Description: `ID of the linked VServer group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_server_certificates",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of slb server certificates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of server certificates IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by server certificate name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB server certificates IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB server certificates names.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A list of SLB server certificates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Server certificate ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Server certificate name.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Server certificate fingerprint.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Server certificate created time.`,
				},
				resource.Attribute{
					Name:        "created_timestamp",
					Description: `Server certificate created timestamp.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB server certificates IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB server certificates names.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A list of SLB server certificates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Server certificate ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Server certificate name.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Server certificate fingerprint.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Server certificate created time.`,
				},
				resource.Attribute{
					Name:        "created_timestamp",
					Description: `Server certificate created timestamp.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_server_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VServer groups related to a server load balancer to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) ID of the SLB.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of VServer group IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by VServer group name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB VServer groups IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB VServer groups names.`,
				},
				resource.Attribute{
					Name:        "slb_server_groups",
					Description: `A list of SLB VServer groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `VServer group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VServer group name.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `ECS instances associated to the group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the attached ECS instance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight associated to the ECS instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB VServer groups IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB VServer groups names.`,
				},
				resource.Attribute{
					Name:        "slb_server_groups",
					Description: `A list of SLB VServer groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `VServer group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VServer group name.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `ECS instances associated to the group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the attached ECS instance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight associated to the ECS instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slb_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for SLB that can be used by an AlibabacloudStack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_details",
					Description: `(Optional) Default to false and only output ` + "`" + `id` + "`" + ` in the ` + "`" + `zones` + "`" + ` block. Set it to true can output more details. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "slb_slave_zone_ids",
					Description: `A list of slb slave zone ids in which the slb master zone.`,
				},
				resource.Attribute{
					Name:        "local_name",
					Description: `The name of the secondary zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "slb_slave_zone_ids",
					Description: `A list of slb slave zone ids in which the slb master zone.`,
				},
				resource.Attribute{
					Name:        "local_name",
					Description: `The name of the secondary zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_slbs",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of server load balancers to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of SLBs IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by SLB name.`,
				},
				resource.Attribute{
					Name:        "master_availability_zone",
					Description: `(Optional) Master availability zone of the SLBs.`,
				},
				resource.Attribute{
					Name:        "slave_availability_zone",
					Description: `(Optional) Slave availability zone of the SLBs.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Optional) Network type of the SLBs. Valid values: ` + "`" + `vpc` + "`" + ` and ` + "`" + `classic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC linked to the SLBs.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) ID of the VSwitch linked to the SLBs.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) Service address of the SLBs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of slb IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of slb names.`,
				},
				resource.Attribute{
					Name:        "slbs",
					Description: `A list of SLBs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SLB.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the SLB belongs to.`,
				},
				resource.Attribute{
					Name:        "master_availability_zone",
					Description: `Master availability zone of the SLBs.`,
				},
				resource.Attribute{
					Name:        "slave_availability_zone",
					Description: `Slave availability zone of the SLBs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `SLB name.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Network type of the SLB. Possible values: ` + "`" + `vpc` + "`" + ` and ` + "`" + `classic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the SLB belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch the SLB belongs to.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Service address of the SLB.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `SLB creation time.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of slb IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of slb names.`,
				},
				resource.Attribute{
					Name:        "slbs",
					Description: `A list of SLBs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SLB.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the SLB belongs to.`,
				},
				resource.Attribute{
					Name:        "master_availability_zone",
					Description: `Master availability zone of the SLBs.`,
				},
				resource.Attribute{
					Name:        "slave_availability_zone",
					Description: `Slave availability zone of the SLBs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `SLB name.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Network type of the SLB. Possible values: ` + "`" + `vpc` + "`" + ` and ` + "`" + `classic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the SLB belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch the SLB belongs to.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Service address of the SLB.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `SLB creation time.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_snapshots",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to get a list of snapshot according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of snapshot IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of snapshots names.`,
				},
				resource.Attribute{
					Name:        "snapshots",
					Description: `A list of snapshots. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `Progress of snapshot creation, presented in percentage.`,
				},
				resource.Attribute{
					Name:        "source_disk_id",
					Description: `Source disk ID, which is retained after the source disk of the snapshot is deleted.`,
				},
				resource.Attribute{
					Name:        "source_disk_size",
					Description: `Size of the source disk, measured in GB.`,
				},
				resource.Attribute{
					Name:        "source_disk_type",
					Description: `Source disk attribute. Value range:`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `Product code on the image market place.`,
				},
				resource.Attribute{
					Name:        "remain_time",
					Description: `The remaining time of a snapshot creation task, in seconds.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time. Time of creation. It is represented according to ISO8601, and UTC time is used. Format: YYYY-MM-DDThh:mmZ.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The snapshot status. Value range:`,
				},
				resource.Attribute{
					Name:        "usage",
					Description: `Whether the snapshots are used to create resources or not. Value range:`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_snat_entries",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Snat Entries owned by an Apsara Stack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Snat Entries IDs.`,
				},
				resource.Attribute{
					Name:        "source_cidr",
					Description: `(Optional) The source CIDR block of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "snat_table_id",
					Description: `(Required) The ID of the Snat table.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Snat Entries IDs.`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `A list of Snat Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "snat_ip",
					Description: `The public IP of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "source_cidr",
					Description: `The source CIDR block of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Snat Entry.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Snat Entries IDs.`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `A list of Snat Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "snat_ip",
					Description: `The public IP of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "source_cidr",
					Description: `The source CIDR block of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Snat Entry.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpc_ipv6_addresses",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Vpc Ipv6 Addresses to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "associated_instance_id",
					Description: `(Optional, ForceNew) The ID of the instance that is assigned the IPv6 address.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, ForceNew) The status of the IPv6 address. Valid values:` + "`" + `Pending` + "`" + ` or ` + "`" + `Available` + "`" + `. - ` + "`" + `Pending` + "`" + `: The IPv6 address is being configured. - ` + "`" + `Available` + "`" + `: The IPv6 address is available.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) The ID of the VPC to which the IPv6 address belongs.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional, ForceNew) The ID of the vSwitch to which the IPv6 address belongs. ## Argument Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Ipv6 Address names.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `A list of Vpc Ipv6 Addresses. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "associated_instance_id",
					Description: `The ID of the instance that is assigned the IPv6 address.`,
				},
				resource.Attribute{
					Name:        "associated_instance_type",
					Description: `The type of the instance that is assigned the IPv6 address.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the IPv6 address was created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Ipv6 Address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `The address of the Ipv6 Address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_id",
					Description: `The ID of the IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_name",
					Description: `The name of the IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway_id",
					Description: `The ID of the IPv6 gateway to which the IPv6 address belongs.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of communication supported by the IPv6 address. Valid values:` + "`" + `Private` + "`" + ` or ` + "`" + `Public` + "`" + `. ` + "`" + `Private` + "`" + `: communication within the private network. ` + "`" + `Public` + "`" + `: communication over the public network`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the IPv6 address. Valid values:` + "`" + `Pending` + "`" + ` or ` + "`" + `Available` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC to which the IPv6 address belongs.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of the vSwitch to which the IPv6 address belongs.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpc_ipv6_egress_rules",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Vpc Ipv6 Egress Rules to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, ForceNew, Computed) A list of Ipv6 Egress Rule IDs.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional, ForceNew) The ID of the instance that is associated with the IPv6 address to which the egress-only rule is applied.`,
				},
				resource.Attribute{
					Name:        "ipv6_egress_rule_name",
					Description: `(Optional, ForceNew) The name of the resource.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway_id",
					Description: `(Required, ForceNew) The ID of the IPv6 gateway.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, ForceNew) A regex string to filter results by Ipv6 Egress Rule name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, ForceNew) The status of the resource. Valid values: ` + "`" + `Available` + "`" + `, ` + "`" + `Deleting` + "`" + `, ` + "`" + `Pending` + "`" + `. ## Argument Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Ipv6 Egress Rule names.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of Vpc Ipv6 Egress Rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the egress-only rule.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Ipv6 Egress Rule. The value formats as ` + "`" + `<ipv6_gateway_id>:<ipv6_egress_rule_id>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance to which the egress-only rule is applied.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of the instance to which the egress-only rule is applied.`,
				},
				resource.Attribute{
					Name:        "ipv6_egress_rule_id",
					Description: `The first ID of the resource.`,
				},
				resource.Attribute{
					Name:        "ipv6_egress_rule_name",
					Description: `The name of the resource.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway_id",
					Description: `The ID of the IPv6 gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource. Valid values: ` + "`" + `Available` + "`" + `, ` + "`" + `Pending` + "`" + ` and ` + "`" + `Deleting` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpc_ipv6_gateways",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Vpc Ipv6 Gateways to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, ForceNew, Computed) A list of Ipv6 Gateway IDs.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway_name",
					Description: `(Optional, ForceNew) The name of the IPv6 gateway.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, ForceNew) A regex string to filter results by Ipv6 Gateway name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, ForceNew) The status of the resource. Valid values: ` + "`" + `Available` + "`" + `, ` + "`" + `Deleting` + "`" + `, ` + "`" + `Pending` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, ForceNew) The ID of the virtual private cloud (VPC) to which the IPv6 gateway belongs. ## Argument Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Ipv6 Gateway names.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `A list of Vpc Ipv6 Gateways. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `The status of the IPv6 gateway. Valid values:` + "`" + `Normal` + "`" + `, ` + "`" + `FinancialLocked` + "`" + ` and ` + "`" + `SecurityLocked` + "`" + `. ` + "`" + `Normal` + "`" + `: working as expected. ` + "`" + `FinancialLocked` + "`" + `: locked due to overdue payments. ` + "`" + `SecurityLocked` + "`" + `: locked due to security reasons.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the IPv6 gateway.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `The time when the IPv6 gateway expires.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Ipv6 Gateway.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `The metering method of the IPv6 gateway. Valid values: ` + "`" + `PayAsYouGo` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway_id",
					Description: `The first ID of the resource.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway_name",
					Description: `The name of the IPv6 gateway.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `The specification of the IPv6 gateway. Valid values: ` + "`" + `Large` + "`" + `, ` + "`" + `Medium` + "`" + ` and ` + "`" + `Small` + "`" + `. ` + "`" + `Small` + "`" + ` (default): Free Edition. ` + "`" + `Medium` + "`" + `: Enterprise Edition . ` + "`" + `Large` + "`" + `: Enhanced Enterprise Edition. The throughput capacity of an IPv6 gateway varies based on the edition. For more information, see [Editions of IPv6 gateways](https://www.alibabacloud.com/help/doc-detail/98926.htm).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the IPv6 gateway. Valid values: ` + "`" + `Available` + "`" + `, ` + "`" + `Deleting` + "`" + `, ` + "`" + `Pending` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the virtual private cloud (VPC) to which the IPv6 gateway belongs.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpc_ipv6_internet_bandwidths",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Vpc Ipv6 Internet Bandwidths to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv6_internet_bandwidth_id",
					Description: `(Optional, ForceNew) The ID of the Ipv6 Internet Bandwidth.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_id",
					Description: `(Optional, ForceNew) The ID of the IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, ForceNew, Computed) A list of Ipv6 Internet Bandwidth IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, ForceNew) The status of the resource. Valid values: ` + "`" + `Normal` + "`" + `, ` + "`" + `FinancialLocked` + "`" + ` and ` + "`" + `SecurityLocked` + "`" + `. ## Argument Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Ipv6 Internet Bandwidth names.`,
				},
				resource.Attribute{
					Name:        "bandwidths",
					Description: `A list of Vpc Ipv6 Internet Bandwidths. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The amount of Internet bandwidth resources of the IPv6 address, Unit: ` + "`" + `Mbit/s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Ipv6 Internet Bandwidth.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `The metering method of the Internet bandwidth resources of the IPv6 gateway.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_id",
					Description: `The ID of the IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway_id",
					Description: `The ID of the IPv6 gateway.`,
				},
				resource.Attribute{
					Name:        "ipv6_internet_bandwidth_id",
					Description: `The ID of the Ipv6 Internet Bandwidth.`,
				},
				resource.Attribute{
					Name:        "payment_type",
					Description: `The payment type of the resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource. Valid values: ` + "`" + `Normal` + "`" + `, ` + "`" + `FinancialLocked` + "`" + ` and ` + "`" + `SecurityLocked` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpcs",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VPCs owned by an AlibabacloudStack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) Filter results by a specific CIDR block. For example: "172.16.0.0/12".`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Filter results by a specific status. Valid value are ` + "`" + `Pending` + "`" + ` and ` + "`" + `Available` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter VPCs by name.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional, type: bool) Indicate whether the VPC is the default one in the specified region.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) Filter results by the specified VSwitch.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of VPC IDs.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional) The Id of resource group which VPC belongs.`,
				},
				resource.Attribute{
					Name:        "dhcp_options_set_id",
					Description: `(Optional, ForceNew) The ID of dhcp options set.`,
				},
				resource.Attribute{
					Name:        "dry_run",
					Description: `(Optional, ForceNew) Indicates whether to check this request only. Valid values: ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `(Optional, ForceNew) The name of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_owner_id",
					Description: `(Optional, ForceNew) The owner ID of VPC.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of VPC IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of VPC names.`,
				},
				resource.Attribute{
					Name:        "vpcs",
					Description: `A list of VPCs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `ID of the region where the VPC is located.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Name of the VPC.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `List of VSwitch IDs in the specified VPC`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `CIDR block of the VPC.`,
				},
				resource.Attribute{
					Name:        "vrouter_id",
					Description: `ID of the VRouter.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `Route table ID of the VRouter.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the VPC`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether the VPC is the default VPC in the region.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the VPC.`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr_block",
					Description: `The IPv6 CIDR block of the VPC.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `The ID of the VRouter.`,
				},
				resource.Attribute{
					Name:        "secondary_cidr_blocks",
					Description: `A list of secondary IPv4 CIDR blocks of the VPC.`,
				},
				resource.Attribute{
					Name:        "user_cidrs",
					Description: `A list of user CIDRs.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of VPC IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of VPC names.`,
				},
				resource.Attribute{
					Name:        "vpcs",
					Description: `A list of VPCs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `ID of the region where the VPC is located.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Name of the VPC.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `List of VSwitch IDs in the specified VPC`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `CIDR block of the VPC.`,
				},
				resource.Attribute{
					Name:        "vrouter_id",
					Description: `ID of the VRouter.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `Route table ID of the VRouter.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the VPC`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether the VPC is the default VPC in the region.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the VPC.`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr_block",
					Description: `The IPv6 CIDR block of the VPC.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `The ID of the VRouter.`,
				},
				resource.Attribute{
					Name:        "secondary_cidr_blocks",
					Description: `A list of secondary IPv4 CIDR blocks of the VPC.`,
				},
				resource.Attribute{
					Name:        "user_cidrs",
					Description: `A list of user CIDRs.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpn_connections",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VPN connections which owned by an Alibabacloudstack account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) IDs of the VPN connections.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `(Optional) Use the VPN gateway ID as the search key.`,
				},
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `(Optional)Use the VPN customer gateway ID as the search key.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string of VPN connection name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) Save the result to the file. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) IDs of the VPN connections.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(Optional) names of the VPN connections.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `A list of VPN connections. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `ID of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "local_subnet",
					Description: `The local subnet of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "remote_subnet",
					Description: `The remote subnet of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the VPN connection, valid value:ike_sa_not_established, ike_sa_established, ipsec_sa_not_established, ipsec_sa_established.`,
				},
				resource.Attribute{
					Name:        "ike_config",
					Description: `The configurations of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_config",
					Description: `The configurations of phase-two negotiation. ### Block ike_config The ike_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "psk",
					Description: `Used for authentication between the IPsec VPN gateway and the customer gateway.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `The version of the IKE protocol.`,
				},
				resource.Attribute{
					Name:        "ike_mode",
					Description: `The negotiation mode of IKE phase-one.`,
				},
				resource.Attribute{
					Name:        "ike_enc_alg",
					Description: `The encryption algorithm of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_auth_alg",
					Description: `The authentication algorithm of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_pfs",
					Description: `The Diffie-Hellman key exchange algorithm used by phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_lifetime",
					Description: `The SA lifecycle as the result of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_local_id",
					Description: `The identification of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "ike_remote_id",
					Description: `The identification of the customer gateway. ### Block ipsec_config The ipsec_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "ipsec_enc_alg",
					Description: `The encryption algorithm of phase-two negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_auth_alg",
					Description: `The authentication algorithm of phase-two negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_pfs",
					Description: `The Diffie-Hellman key exchange algorithm used by phase-two negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_lifetime",
					Description: `The SA lifecycle as the result of phase-two negotiation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) IDs of the VPN connections.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(Optional) names of the VPN connections.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `A list of VPN connections. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `ID of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "local_subnet",
					Description: `The local subnet of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "remote_subnet",
					Description: `The remote subnet of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the VPN connection, valid value:ike_sa_not_established, ike_sa_established, ipsec_sa_not_established, ipsec_sa_established.`,
				},
				resource.Attribute{
					Name:        "ike_config",
					Description: `The configurations of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_config",
					Description: `The configurations of phase-two negotiation. ### Block ike_config The ike_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "psk",
					Description: `Used for authentication between the IPsec VPN gateway and the customer gateway.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `The version of the IKE protocol.`,
				},
				resource.Attribute{
					Name:        "ike_mode",
					Description: `The negotiation mode of IKE phase-one.`,
				},
				resource.Attribute{
					Name:        "ike_enc_alg",
					Description: `The encryption algorithm of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_auth_alg",
					Description: `The authentication algorithm of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_pfs",
					Description: `The Diffie-Hellman key exchange algorithm used by phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_lifetime",
					Description: `The SA lifecycle as the result of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_local_id",
					Description: `The identification of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "ike_remote_id",
					Description: `The identification of the customer gateway. ### Block ipsec_config The ipsec_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "ipsec_enc_alg",
					Description: `The encryption algorithm of phase-two negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_auth_alg",
					Description: `The authentication algorithm of phase-two negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_pfs",
					Description: `The Diffie-Hellman key exchange algorithm used by phase-two negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_lifetime",
					Description: `The SA lifecycle as the result of phase-two negotiation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpn_customer_gateways",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VPN customer gateways which owned by an Alibabacloudstack account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) ID of the VPN customer gateways.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string of VPN customer gateways name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) Save the result to the file. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `A list of VPN customer gateways. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN customer gateway .`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of the VPN customer gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gateways",
					Description: `A list of VPN customer gateways. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN customer gateway .`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of the VPN customer gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vpn_gateways",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VPNs which owned by an Alibabacloudstack account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Use the VPC ID as the search key.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) IDs of the VPN.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Limit search to specific status - valid value is "Init", "Provisioning", "Active", "Updating", "Deleting".`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `(Optional) Limit search to specific business status - valid value is "Normal", "FinancialLocked".`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string of VPN name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) Save the result to the file. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `IDs of the VPN.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `names of the VPN.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `A list of VPN gateways. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC that the VPN belongs.`,
				},
				resource.Attribute{
					Name:        "internet_ip",
					Description: `The internet ip of the VPN.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The expiration time of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `The Specification of the VPN`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the VPN`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the VPN`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `The business status of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `The charge type of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "enable_ipsec",
					Description: `Whether the ipsec function is enabled.`,
				},
				resource.Attribute{
					Name:        "enable_ssl",
					Description: `Whether the ssl function is enabled.`,
				},
				resource.Attribute{
					Name:        "ssl_connections",
					Description: `Total count of ssl vpn connections.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `IDs of the VPN.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `names of the VPN.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `A list of VPN gateways. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC that the VPN belongs.`,
				},
				resource.Attribute{
					Name:        "internet_ip",
					Description: `The internet ip of the VPN.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The expiration time of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `The Specification of the VPN`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the VPN`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the VPN`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `The business status of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `The charge type of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "enable_ipsec",
					Description: `Whether the ipsec function is enabled.`,
				},
				resource.Attribute{
					Name:        "enable_ssl",
					Description: `Whether the ssl function is enabled.`,
				},
				resource.Attribute{
					Name:        "ssl_connections",
					Description: `Total count of ssl vpn connections.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_vswitches",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VSwitch owned by an Alibabacloudstack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) Filter results by a specific CIDR block. For example: "172.16.0.0/12".`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The availability zone of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by name.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional, type: bool) Indicate whether the VSwitch is created by the system.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC that owns the VSwitch.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of VSwitch IDs. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of VSwitch IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of VSwitch names.`,
				},
				resource.Attribute{
					Name:        "vswitches",
					Description: `A list of VSwitches. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `ID of the availability zone where the VSwitch is located.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC that owns the VSwitch.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `List of ECS instance IDs in the specified VSwitch.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `CIDR block of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether the VSwitch is the default one in the region.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of VSwitch IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of VSwitch names.`,
				},
				resource.Attribute{
					Name:        "vswitches",
					Description: `A list of VSwitches. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `ID of the availability zone where the VSwitch is located.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC that owns the VSwitch.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `List of ECS instance IDs in the specified VSwitch.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `CIDR block of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether the VSwitch is the default one in the region.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alibabacloudstack_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones that can be used by an Alibabacloudstack Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "available_instance_type",
					Description: `(Optional) Filter the results by a specific instance type.`,
				},
				resource.Attribute{
					Name:        "available_resource_creation",
					Description: `(Optional) Filter the results by a specific resource type. Valid values: ` + "`" + `Instance` + "`" + `, ` + "`" + `Disk` + "`" + `, ` + "`" + `VSwitch` + "`" + `, ` + "`" + `Rds` + "`" + `, ` + "`" + `KVStore` + "`" + `, ` + "`" + `Slb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "available_disk_category",
					Description: `(Optional) Filter the results by a specific disk category. Can be either ` + "`" + `cloud` + "`" + `, ` + "`" + `cloud_efficiency` + "`" + `, ` + "`" + `cloud_ssd` + "`" + `, ` + "`" + `ephemeral_ssd` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "multi",
					Description: `(Optional, type: bool) Indicate whether the zones can be used in a multi AZ configuration. Default to ` + "`" + `false` + "`" + `. Multi AZ is usually used to launch RDS instances.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Filter the results by a specific ECS instance charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Optional) Filter the results by a specific network type. Valid values: ` + "`" + `Classic` + "`" + ` and ` + "`" + `Vpc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_strategy",
					Description: `- (Optional) Filter the results by a specific ECS spot type. Valid values: ` + "`" + `NoSpot` + "`" + `, ` + "`" + `SpotWithPriceLimit` + "`" + ` and ` + "`" + `SpotAsPriceGo` + "`" + `. Default to ` + "`" + `NoSpot` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_details",
					Description: `(Optional) Default to false and only output ` + "`" + `id` + "`" + ` in the ` + "`" + `zones` + "`" + ` block. Set it to true can output more details.`,
				},
				resource.Attribute{
					Name:        "available_slb_address_type",
					Description: `Filter the results by a slb instance address type. Can be either ` + "`" + `Vpc` + "`" + `, ` + "`" + `classic_internet` + "`" + ` or ` + "`" + `classic_intranet` + "`" + ``,
				},
				resource.Attribute{
					Name:        "available_slb_address_ip_version",
					Description: `Filter the results by a slb instance address version. Can be either ` + "`" + `ipv4` + "`" + `, or ` + "`" + `ipv6` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "local_name",
					Description: `Name of the zone in the local language.`,
				},
				resource.Attribute{
					Name:        "available_instance_types",
					Description: `Allowed instance types.`,
				},
				resource.Attribute{
					Name:        "available_resource_creation",
					Description: `Type of resources that can be created.`,
				},
				resource.Attribute{
					Name:        "available_disk_categories",
					Description: `Set of supported disk categories.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
				resource.Attribute{
					Name:        "slb_slave_zone_ids",
					Description: `A list of slb slave zone ids in which the slb master zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "local_name",
					Description: `Name of the zone in the local language.`,
				},
				resource.Attribute{
					Name:        "available_instance_types",
					Description: `Allowed instance types.`,
				},
				resource.Attribute{
					Name:        "available_resource_creation",
					Description: `Type of resources that can be created.`,
				},
				resource.Attribute{
					Name:        "available_disk_categories",
					Description: `Set of supported disk categories.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
				resource.Attribute{
					Name:        "slb_slave_zone_ids",
					Description: `A list of slb slave zone ids in which the slb master zone.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"alibabacloudstack_adb_db_clusters":                        0,
		"alibabacloudstack_adb_zones":                              1,
		"alibabacloudstack_api_gateway_apis":                       2,
		"alibabacloudstack_api_gateway_apps":                       3,
		"alibabacloudstack_api_gateway_groups":                     4,
		"alibabacloudstack_api_gateway_service":                    5,
		"alibabacloudstack_ascm_ecs_instance_families":             6,
		"alibabacloudstack_ascm_environment_services_by_product":   7,
		"alibabacloudstack_ascm_instance_families":                 8,
		"alibabacloudstack_ascm_logon_policies":                    9,
		"alibabacloudstack_ascm_organizations":                     10,
		"alibabacloudstack_ascm_password_policies":                 11,
		"alibabacloudstack_ascm_quotas":                            12,
		"alibabacloudstack_ascm_ram_policies":                      13,
		"alibabacloudstack_ascm_ram_policies_for_user":             14,
		"alibabacloudstack_ascm_ram_service_roles":                 15,
		"alibabacloudstack_ascm_regions_by_product":                16,
		"alibabacloudstack_ascm_resource_groups":                   17,
		"alibabacloudstack_ascm_roles":                             18,
		"alibabacloudstack_ascm_service_clusters_by_product":       19,
		"alibabacloudstack_ascm_specific_fields":                   20,
		"alibabacloudstack_ascm_users":                             21,
		"alibabacloudstack_cms_alarm_contact_groups":               22,
		"alibabacloudstack_cms_alarms":                             23,
		"alibabacloudstack_cms_metric_metalist":                    24,
		"alibabacloudstack_cms_project_meta":                       25,
		"alibabacloudstack_common_bandwidth_packages":              26,
		"alibabacloudstack_cr_ee_instances":                        27,
		"alibabacloudstack_cr_ee_namespaces":                       28,
		"alibabacloudstack_cr_ee_repos":                            29,
		"alibabacloudstack_cr_ee_sync_rules":                       30,
		"alibabacloudstack_cr_namespaces":                          31,
		"alibabacloudstack_cr_repos":                               32,
		"alibabacloudstack_cs_kubernetes_clusters":                 33,
		"alibabacloudstack_datahub_service":                        34,
		"alibabacloudstack_db_instances":                           35,
		"alibabacloudstack_db_zones":                               36,
		"alibabacloudstack_disks":                                  37,
		"alibabacloudstack_dms_enterprise_instances":               38,
		"alibabacloudstack_dms_enterprise_users":                   39,
		"alibabacloudstack_dns_domains":                            40,
		"alibabacloudstack_dns_groups":                             41,
		"alibabacloudstack_dns_records":                            42,
		"alibabacloudstack_drds_instances":                         43,
		"alibabacloudstack_ecs_auto_snapshot_policies":             44,
		"alibabacloudstack_ecs_commands":                           45,
		"alibabacloudstack_ecs_dedicated_hosts":                    46,
		"alibabacloudstack_ehpc_job_templates":                     47,
		"alibabacloudstack_eips":                                   48,
		"alibabacloudstack_elasticsearch":                          49,
		"alibabacloudstack_elasticsearch_zones":                    50,
		"alibabacloudstack_ess_lifecycle_hooks":                    51,
		"alibabacloudstack_ess_notifications":                      52,
		"alibabacloudstack_ess_scaling_configurations":             53,
		"alibabacloudstack_ess_scaling_groups":                     54,
		"alibabacloudstack_ess_scaling_rules":                      55,
		"alibabacloudstack_ess_scheduled_tasks":                    56,
		"alibabacloudstack_express_connect_access_points":          57,
		"alibabacloudstack_express_connect_physical_connections":   58,
		"alibabacloudstack_express_connect_virtual_border_routers": 59,
		"alibabacloudstack_forward_entries":                        60,
		"alibabacloudstack_gpdb_instances":                         61,
		"alibabacloudstack_images":                                 62,
		"alibabacloudstack_instance_type_families":                 63,
		"alibabacloudstack_instance_types":                         64,
		"alibabacloudstack_instances":                              65,
		"alibabacloudstack_key_pairs":                              66,
		"alibabacloudstack_kms_aliases":                            67,
		"alibabacloudstack_kms_ciphertext":                         68,
		"alibabacloudstack_kms_keys":                               69,
		"alibabacloudstack_kms_secrets":                            70,
		"alibabacloudstack_kvstore_instance_classes":               71,
		"alibabacloudstack_kvstore_instance_engines":               72,
		"alibabacloudstack_kvstore_instances":                      73,
		"alibabacloudstack_kvstore_zones":                          74,
		"alibabacloudstack_nat_gateways":                           75,
		"alibabacloudstack_network_acls":                           76,
		"alibabacloudstack_network_interfaces":                     77,
		"alibabacloudstack_ons_groups":                             78,
		"alibabacloudstack_ons_instances":                          79,
		"alibabacloudstack_ons_topics":                             80,
		"alibabacloudstack_oos_executions":                         81,
		"alibabacloudstack_oos_templates":                          82,
		"alibabacloudstack_oss_bucket_objects":                     83,
		"alibabacloudstack_oss_buckets":                            84,
		"alibabacloudstack_route_entries":                          85,
		"alibabacloudstack_route_tables":                           86,
		"alibabacloudstack_router_interfaces":                      87,
		"alibabacloudstack_security_group_rules":                   88,
		"alibabacloudstack_security_groups":                        89,
		"alibabacloudstack_slb_acls":                               90,
		"alibabacloudstack_slb_backend_servers":                    91,
		"alibabacloudstack_slb_ca_certificates":                    92,
		"alibabacloudstack_slb_domain_extensions":                  93,
		"alibabacloudstack_slb_listeners":                          94,
		"alibabacloudstack_slb_master_slave_server_groups":         95,
		"alibabacloudstack_slb_rules":                              96,
		"alibabacloudstack_slb_server_certificates":                97,
		"alibabacloudstack_slb_server_groups":                      98,
		"alibabacloudstack_slb_zones":                              99,
		"alibabacloudstack_slbs":                                   100,
		"alibabacloudstack_snapshots":                              101,
		"alibabacloudstack_snat_entries":                           102,
		"alibabacloudstack_vpc_ipv6_addresses":                     103,
		"alibabacloudstack_vpc_ipv6_egress_rules":                  104,
		"alibabacloudstack_vpc_ipv6_gateways":                      105,
		"alibabacloudstack_vpc_ipv6_internet_bandwidths":           106,
		"alibabacloudstack_vpcs":                                   107,
		"alibabacloudstack_vpn_connections":                        108,
		"alibabacloudstack_vpn_customer_gateways":                  109,
		"alibabacloudstack_vpn_gateways":                           110,
		"alibabacloudstack_vswitches":                              111,
		"alibabacloudstack_zones":                                  112,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
